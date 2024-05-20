package repository

import (
	"database/sql"
	"fmt"

	"github.com/CarlosEduardoNop/apphook/pkg/db"
)

func FindById(table string, id int) (*sql.Row, error) {
	conn, err := db.OpenConnection()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer conn.Close()

	sql := fmt.Sprintf("SELECT * FROM %s WHERE id = ? LIMIT 1", table)

	res := conn.QueryRow(sql, id)

	return res, nil
}
