package vegeTools

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	CertName = "cert.pem"
	KeyName  = "key.pem"
)
var (
	host       = flag.String("host", "", "Comma-separated hostnames and IPs to generate a certificate for")
	validFrom  = flag.String("start-date", "", "Creation date formatted as Jan 1 15:04:05 2011")
	validFor   = flag.Duration("duration", 365*24*time.Hour, "Duration that certificate is valid for")
	isCA       = flag.Bool("ca", false, "whether this cert should be its own Certificate Authority")
	rsaBits    = flag.Int("rsa-bits", 2048, "Size of RSA key to generate. Ignored if --ecdsa-curve is set")
	ecdsaCurve = flag.String("ecdsa-curve", "", "ECDSA curve to use to generate a key. Valid values are P224, P256 (recommended), P384, P521")
	ed25519Key = flag.Bool("ed25519", false, "Generate an Ed25519 key")
)

func publicKey(priv any) any {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &k.PublicKey
	case *ecdsa.PrivateKey:
		return &k.PublicKey
	case ed25519.PrivateKey:
		return k.Public().(ed25519.PublicKey)
	default:
		return nil
	}
}

// TlsCertGenerateToFile 生成自签tls证书 cert.pem key.pem 到指定文件夹中
func TlsCertGenerateToFile(path string) error {
	flag.Parse()

	var priv any
	var err error
	switch *ecdsaCurve {
	case "":
		if *ed25519Key {
			_, priv, err = ed25519.GenerateKey(rand.Reader)
		} else {
			priv, err = rsa.GenerateKey(rand.Reader, *rsaBits)
		}
	case "P224":
		priv, err = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	case "P256":
		priv, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case "P384":
		priv, err = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case "P521":
		priv, err = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	default:
		err = errors.New(fmt.Sprintf("Unrecognized elliptic curve: %q", *ecdsaCurve))
	}
	if err != nil {
		err = errors.New(fmt.Sprintf("Failed to generate private key: %v", err))
		return err
	}

	// ECDSA, ED25519 and RSA subject keys should have the DigitalSignature
	// KeyUsage bits set in the x509.Certificate template
	keyUsage := x509.KeyUsageDigitalSignature
	// Only RSA subject keys should have the KeyEncipherment KeyUsage bits set. In
	// the context of TLS this KeyUsage is particular to RSA key exchange and
	// authentication.
	if _, isRSA := priv.(*rsa.PrivateKey); isRSA {
		keyUsage |= x509.KeyUsageKeyEncipherment
	}

	var notBefore time.Time
	if len(*validFrom) == 0 {
		notBefore = time.Now()
	} else {
		notBefore, err = time.Parse("Jan 2 15:04:05 2006", *validFrom)
		if err != nil {
			err = errors.New(fmt.Sprintf("Failed to parse creation date: %v", err))
			return err
		}
	}

	notAfter := notBefore.Add(*validFor)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		err = errors.New(fmt.Sprintf("Failed to generate serial number: %v", err))
		return err
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Acme Co"},
		},
		NotBefore: notBefore,
		NotAfter:  notAfter,

		KeyUsage:              keyUsage,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	hosts := make([]string, 0)
	if len(*host) != 0 {
		hosts = strings.Split(*host, ",")
	} else {
		hosts, err = GetLocalIpv4List()
		if err != nil {
			err = errors.New(fmt.Sprintf("Failed to GetLocalIpv4List: %v", err))
			return err
		}
	}

	for _, h := range hosts {
		if ip := net.ParseIP(h); ip != nil {
			template.IPAddresses = append(template.IPAddresses, ip)
		} else {
			template.DNSNames = append(template.DNSNames, h)
		}
	}

	if *isCA {
		template.IsCA = true
		template.KeyUsage |= x509.KeyUsageCertSign
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, publicKey(priv), priv)
	if err != nil {
		err = errors.New(fmt.Sprintf("Failed to create certificate: %v", err))
		return err
	}

	c := filepath.Join(path, CertName)
	certOut, err := os.Create(c)
	if err != nil {
		err = errors.New(fmt.Sprintf("Failed to open cert.pem for writing: %v", err))
		return err
	}
	if err = pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		err = errors.New(fmt.Sprintf("Failed to write data to cert.pem: %v", err))
		return err
	}
	if err = certOut.Close(); err != nil {
		err = errors.New(fmt.Sprintf("Error closing cert.pem: %v", err))
		return err
	}

	log.Println("KdIbIlLT wrote " + c)

	k := filepath.Join(path, KeyName)
	keyOut, err := os.OpenFile(k, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		err = errors.New(fmt.Sprintf("Failed to open key.pem for writing: %v", err))
		return err
	}
	privBytes, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		err = errors.New(fmt.Sprintf("Unable to marshal private key: %v", err))
		return err
	}
	if err = pem.Encode(keyOut, &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes}); err != nil {
		err = errors.New(fmt.Sprintf("Failed to write data to key.pem: %v", err))
		return err
	}
	if err = keyOut.Close(); err != nil {
		err = errors.New(fmt.Sprintf("Error closing key.pem: %v", err))
		return err
	}
	log.Println("ShyxpJet wrote " + k)
	return nil
}

// TlsCertGenerateToMap 生成自签tls证书 cert.pem key.pem 到map
func TlsCertGenerateToMap() (cert map[string]string, err error) {
	flag.Parse()

	var priv any
	switch *ecdsaCurve {
	case "":
		if *ed25519Key {
			_, priv, err = ed25519.GenerateKey(rand.Reader)
		} else {
			priv, err = rsa.GenerateKey(rand.Reader, *rsaBits)
		}
	case "P224":
		priv, err = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	case "P256":
		priv, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case "P384":
		priv, err = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case "P521":
		priv, err = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	default:
		err = errors.New(fmt.Sprintf("Unrecognized elliptic curve: %q", *ecdsaCurve))
	}
	if err != nil {
		err = errors.New(fmt.Sprintf("Failed to generate private key: %v", err))
		return nil, err
	}

	// ECDSA, ED25519 and RSA subject keys should have the DigitalSignature
	// KeyUsage bits set in the x509.Certificate template
	keyUsage := x509.KeyUsageDigitalSignature
	// Only RSA subject keys should have the KeyEncipherment KeyUsage bits set. In
	// the context of TLS this KeyUsage is particular to RSA key exchange and
	// authentication.
	if _, isRSA := priv.(*rsa.PrivateKey); isRSA {
		keyUsage |= x509.KeyUsageKeyEncipherment
	}

	var notBefore time.Time
	if len(*validFrom) == 0 {
		notBefore = time.Now()
	} else {
		notBefore, err = time.Parse("Jan 2 15:04:05 2006", *validFrom)
		if err != nil {
			err = errors.New(fmt.Sprintf("Failed to parse creation date: %v", err))
			return nil, err
		}
	}

	notAfter := notBefore.Add(*validFor)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		err = errors.New(fmt.Sprintf("Failed to generate serial number: %v", err))
		return nil, err
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Acme Co"},
		},
		NotBefore: notBefore,
		NotAfter:  notAfter,

		KeyUsage:              keyUsage,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	hosts := make([]string, 0)
	if len(*host) != 0 {
		hosts = strings.Split(*host, ",")
	} else {
		hosts, err = GetLocalIpv4List()
		if err != nil {
			err = errors.New(fmt.Sprintf("Failed to GetLocalIpv4List: %v", err))
			return nil, err
		}
	}

	for _, h := range hosts {
		if ip := net.ParseIP(h); ip != nil {
			template.IPAddresses = append(template.IPAddresses, ip)
		} else {
			template.DNSNames = append(template.DNSNames, h)
		}
	}

	if *isCA {
		template.IsCA = true
		template.KeyUsage |= x509.KeyUsageCertSign
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, publicKey(priv), priv)
	if err != nil {
		err = errors.New(fmt.Sprintf("Failed to create certificate: %v", err))
		return nil, err
	}

	//writer
	buf := bytes.NewBuffer(make([]byte, 0))
	cert = make(map[string]string, 2)

	certOut := buf
	if err = pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		err = errors.New(fmt.Sprintf("Failed to write data to cert.pem: %v", err))
		return nil, err
	}
	cert[CertName] = certOut.String()
	log.Println("KdIbIlLT wrote " + CertName)
	buf.Reset()

	keyOut := buf
	privBytes, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		err = errors.New(fmt.Sprintf("Unable to marshal private key: %v", err))
		return nil, err
	}
	if err = pem.Encode(keyOut, &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes}); err != nil {
		err = errors.New(fmt.Sprintf("Failed to write data to key.pem: %v", err))
		return nil, err
	}
	cert[KeyName] = keyOut.String()
	buf.Reset()
	log.Println("ShyxpJet wrote " + KeyName)
	return cert, nil
}
