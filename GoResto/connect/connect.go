package connect

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //github call to every controller
)

func ConnectHandler() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/restoran")
	if err != nil {
		return nil, err
	}
	return db, nil
}
