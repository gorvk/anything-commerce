package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorvk/anything-commerce/server/api-services/common"
	"github.com/gorvk/anything-commerce/server/api-services/common/constants"
	customTypes "github.com/gorvk/anything-commerce/server/api-services/common/types"
	models "github.com/gorvk/anything-commerce/server/api-services/models/user"
)

func UpdateUserByEmail(w http.ResponseWriter, r *http.Request) {
	common.CheckHttpResponseType(w, r, http.MethodGet)
	var err error

	d, err := io.ReadAll(r.Body)
	common.HandleHttpError(err, w, constants.ERROR_HTTP_INVALID_REQUEST_BODY, http.StatusBadRequest)

	var payload customTypes.UPDATE_USER_BY_EMAIL_INPUT

	err = json.Unmarshal(d, &payload)
	common.HandleHttpError(err, w, constants.ERROR_HTTP_UNABLE_TO_PARSE_REQUEST, http.StatusBadRequest)

	_, err = models.UpdateUserByEmail(payload.Email, payload.Field, payload.Value)
	common.HandleDbError(err, w, constants.ERROR_DB_UNABLE_TO_UPDATE_RECORD, http.StatusInternalServerError)
}
