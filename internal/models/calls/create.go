package calls

import (
	"database/sql"
	"fmt"

	"github.com/CarlosEduardoNop/apphook/internal/repository"
)

func Create(data map[string]interface{}) (sql.Result, error) {
	res, err := repository.Create("calls", data)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return res, nil
}
