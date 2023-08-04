package vegeMysql

import (
	"fmt"
	"testing"
)

func TestMysqlDSN(t *testing.T) {
	dsn := "user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	parseDSN := ParseMysqlDSN(dsn)
	fmt.Printf("%#v", parseDSN)
	newDns := NewDefaultMysqlDsn(
		DefaultParams,
		WithAuth("user", "password"),
		WithAddress("127.0.0.1:3306"),
		WithDatabase("sql_test"),
	)
	fmt.Println()
	fmt.Printf("%s", newDns.String())
}
