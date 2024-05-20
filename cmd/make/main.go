package main

import (
	"flag"
	"fmt"

	"github.com/CarlosEduardoNop/apphook/internal/migration"
)

func main() {
	makeType := flag.String("type", "", "Type of make")
	migrationName := flag.String("name", "", "Name of migration")
	flag.Parse()

	if *makeType == "" {
		fmt.Println("Type is required")
		return
	}

	if *makeType == "migration" {
		migration.Make(migrationName)
	}
}
