package vegeTools

import (
	"fmt"
	"testing"
)

func TestMysqlDSN(t *testing.T) {
	dsn := "user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	parseDSN := ParseMysqlDSN(dsn)
	fmt.Printf("%#v", parseDSN)
	newDns := NewDefaultMysqlDsn()
	newDns.Username = "user"
	newDns.Password = "password"
	newDns.Address = "127.0.0.1:3306"
	newDns.Database = "sql_test"
	fmt.Println()
	fmt.Printf("%s", newDns.String())
}
