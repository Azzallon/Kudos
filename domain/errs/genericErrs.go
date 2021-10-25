package errs

import (
	"database/sql"
	"fmt"

	"github.com/labstack/gommon/log"
)

func GetGenericErrorMessage(errorType error) *ApiError {
	if errorType == nil {
		log.Error("Nil reference error while handling error!")
		unexpectedMsg := fmt.Sprintf("Unexpected Error: %v", errorType)
		return NewUnexpectedError(unexpectedMsg)
	}

	if errorType == sql.ErrNoRows {
		return NewBadRequestError("Dados solicitados não encontrados")
	}

	unexpectedMsg := fmt.Sprintf("Ocorreu um erro inesperado: %v", errorType)
	return NewUnexpectedError(unexpectedMsg)
}
