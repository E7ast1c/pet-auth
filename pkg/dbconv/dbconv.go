package dbconv

import "fmt"

func PGStandardConv(port uint32, host, user, dbname, password string) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}

// PGURLConv The URI scheme designator can be either postgresql:// or postgres://. Each of the remaining URI parts is optional.
// https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING
func PGURLConv(port uint32, host, user, dbname, password string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		user, password, host, port, dbname)
}
