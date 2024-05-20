package repository

import (
	"fmt"

	"github.com/CarlosEduardoNop/apphook/pkg/db"
)

func Update(table string, data map[string]interface{}, columnSearch map[string]interface{}) error {
	conn, err := db.OpenConnection()

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer conn.Close()

	search := ""

	for key, value := range columnSearch {
		search = fmt.Sprintf("%s = %v", key, value)
	}

	values := ""
	count := 0

	for key, value := range data {
		count++

		if count == len(data) {
			if stringValue, ok := value.(string); ok {
				values += fmt.Sprintf("%s = '%v'", key, stringValue)
				continue
			}

			values += fmt.Sprintf("%s = %v", key, value)
			continue
		}

		if stringValue, ok := value.(string); ok {
			values += fmt.Sprintf("%s = '%v', ", key, stringValue)
			continue
		}

		values += fmt.Sprintf("%s = %v, ", key, value)
	}

	sql := fmt.Sprintf("UPDATE %s SET %s WHERE %s", table, values, search)

	_, err = conn.Exec(sql)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
