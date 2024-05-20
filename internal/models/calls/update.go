package calls

import "github.com/CarlosEduardoNop/apphook/internal/repository"

func Update(data map[string]interface{}, columnSearch map[string]interface{}) error {
	err := repository.Update("calls", data, columnSearch)

	if err != nil {
		return err
	}

	return nil
}
