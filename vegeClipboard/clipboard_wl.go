//go:build wl

package vegeClipboard

import (
	"os/exec"
	"bytes"
)

func Copy(text string) error {
	cmd := exec.Command("wl-copy")
	cmd.Stdin = bytes.NewBufferString(text)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func Paste() (string, error) {
	cmd := exec.Command("wl-paste", "-n")
	cmd.Stdin = bytes.NewBufferString("-n")
	p, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(p), nil
}
