package vegeMysql

import (
	"net/url"
	"reflect"
	"regexp"
	"strings"
)

type MysqlDSN struct {
	Dsn      string
	Type     string //deprecated
	Username string
	Password string
	Protocol string
	Address  string
	Database string
	Params   map[string]string
}

var DefaultParams = map[string]string{
	"charset":   "utf8mb4",
	"parseTime": "True",
	"loc":       "Local",
}

type MysqlDsnOptionFunc func(dsn *MysqlDSN)

func NewDefaultMysqlDsn(params map[string]string, ops ...MysqlDsnOptionFunc) *MysqlDSN {
	obj := new(MysqlDSN)
	obj.Params = params
	obj.Protocol = "tcp"
	for _, op := range ops {
		op(obj)
	}
	return obj
}

func WithAuth(u string, pswd string) MysqlDsnOptionFunc {
	return func(o *MysqlDSN) {
		o.Username = u
		o.Password = pswd
	}
}

func WithAddress(addr string) MysqlDsnOptionFunc {
	return func(o *MysqlDSN) {
		o.Address = addr
	}
}

func WithDatabase(dbname string) MysqlDsnOptionFunc {
	return func(o *MysqlDSN) {
		o.Database = dbname
	}
}

// Converts a MysqlDSN struct into its string representation.
func (d *MysqlDSN) String() string {
	if len(d.Dsn) != 0 {
		return d.Dsn
	}
	str := ""
	if d.Type != "" {
		str += d.Type + "://"
	}
	if d.Username != "" {
		str += d.Username
	}
	if d.Password != "" {
		str += ":" + d.Password
	}
	if d.Username != "" && d.Password != "" {
		str += "@"
	}
	if d.Protocol != "" {
		str += d.Protocol
	}
	if d.Address != "" {
		str += "(" + d.Address + ")"
	}
	str += "/"
	if d.Database != "" {
		str += d.Database
	}
	if d.Params != nil && len(d.Params) > 0 {
		str += "?"
		i := 0
		for key, value := range d.Params {
			str += key + "=" + value
			if i < len(d.Params)-1 {
				str += "&"
			}
			i++
		}
	}
	d.Dsn = str
	return d.Dsn
}

var (
	// Regex testing: http://regoio.herokuapp.com/
	regex = regexp.MustCompile(
		`^(?:(?P<Type>.*?)?://)?` + // [type://]
			`(?:(?P<Username>.*?)(?::(?P<Password>.*))?@)?` + // [username[:password]@]
			`(?:(?P<Protocol>[^\(]*)(?:\((?P<Address>[^\)]*)\))?)?` + // [protocol[(address)]]
			`\/(?P<Database>.*?)` + // /database
			`(?:\?(?P<Params>[^\?]*))?$`) // [?param1=value1]
)

// Turns a MysqlDSN string into a parsed MysqlDSN struct.
func ParseMysqlDSN(s string) *MysqlDSN {
	dsn := &MysqlDSN{Dsn: s}
	matches := regex.FindStringSubmatch(s)
	names := regex.SubexpNames()
	vof := reflect.ValueOf(dsn).Elem()
	if len(matches) > 0 {
		for n, match := range matches[1:] {
			name := names[n+1]
			if name == "Params" {
				values, err := url.ParseQuery(match)
				if err != nil {
					panic(err)
				}
				dsn.Params = make(map[string]string)
				for key, vals := range values {
					dsn.Params[key] = strings.Join(vals, ",")
				}
			} else {
				vof.FieldByName(name).SetString(match)
			}
		}
	}
	return dsn
}
