package common

import (
	"fmt"
	"net/http"

	"github.com/gorvk/anything-commerce/server/api-services/common/constants"
	"github.com/lib/pq"
)

func HandleDbError(err error, w http.ResponseWriter, friendlyMessage string, statusCode int) {
	pqErr, isPqError := err.(*pq.Error)
	if isPqError {
		msg := fmt.Sprintf("PostgreSQL Error: %q - %q", friendlyMessage, pqErr.Message)
		http.Error(w, msg, statusCode)
	}
}

func CheckHttpResponseType(w http.ResponseWriter, r *http.Request, methodType string) {
	if r.Method == methodType {
		return
	}
	HandleHttpError(http.ErrNotSupported, w, constants.ERROR_HTTP_METHOD_NOT_ALLOWED, http.StatusMethodNotAllowed)
}

func HandleHttpError(err error, w http.ResponseWriter, friendlyMessage string, statusCode int) {
	if err != nil {
		msg := fmt.Sprintf("Http Error: %q", friendlyMessage)
		http.Error(w, msg, statusCode)
	}
}
