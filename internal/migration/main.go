package migration

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Make(migrationName *string) {
	files, err := filepath.Glob("migrations/*.sql")

	if err != nil {
		fmt.Println("Error to list migrations:", err)
		return
	}

	sequencialName := fmt.Sprintf("migrations/%d_%s.sql", GetLastMigrationId(files), *migrationName)

	_, err = os.Create(sequencialName)
	if err != nil {
		fmt.Println("Error trying create migration", err)
		return
	}

	fmt.Println("Migration successfully created:", sequencialName)
}

func GetLastMigrationId(files []string) int {
	if len(files) == 0 {
		return 1
	}

	lastFile := files[len(files)-1]

	split := strings.Split(lastFile, "migrations\\")

	id := split[1]

	indiceUnderscore := strings.Index(id, "_")
	if indiceUnderscore != -1 {
		id = id[:indiceUnderscore]
	}

	idInt, _ := strconv.Atoi(id)

	return idInt + 1
}


