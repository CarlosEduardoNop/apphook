package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/CarlosEduardoNop/apphook/pkg/db"
)

type Migration struct {
	ID   string `bson:"_id,omiempty" json:"id"`
	Name string `bson:"name" json:"name"`
}

func main() {
	conn, err := db.OpenConnection()

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	CreateTable(conn)

	res, err := conn.Query("SELECT * FROM migrations;")

	if err != nil {
		fmt.Println(err)
	}

	defer res.Close()

	migrations := []Migration{}

	for res.Next() {
		var migration Migration

		err = res.Scan(&migration.ID, &migration.Name)

		if err != nil {
			fmt.Println(err)
		}

		migrations = append(migrations, migration)
	}

	filepath.Walk("./migrations", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if !strings.HasSuffix(info.Name(), ".sql") {
			return nil
		}

		sqlFile, err := ioutil.ReadFile(path)

		if err != nil {
			return err
		}

		reg := regexp.MustCompile(`^(\d+)_`)

		match := reg.FindStringSubmatch(info.Name())[1]

		if inSlice(match, migrations) {
			return nil
		}

		_, err = conn.Exec(string(sqlFile))

		if err != nil {
			fmt.Println(err)
			return err
		}

		sql := fmt.Sprintf("INSERT INTO migrations (id, name) VALUES ('%s', '%s');", match, info.Name())

		_, err = conn.Exec(sql)

		if err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Printf("%s migrated", info.Name())

		return nil
	})
}

func inSlice(val interface{}, slice []Migration) bool {
	for _, item := range slice {
		if item.ID == val {
			return true
		}
	}
	return false
}

func CreateTable(conn *sql.DB) error {
	_, err := conn.Query("CREATE TABLE IF NOT EXISTS migrations (id VARCHAR(255) PRIMARY KEY, name VARCHAR(255));")

	if err != nil {
		fmt.Println(err)
	}

	return nil
}
