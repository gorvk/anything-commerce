package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorvk/anything-commerce/server/api-services/common"
	"github.com/gorvk/anything-commerce/server/api-services/common/constants"
	customTypes "github.com/gorvk/anything-commerce/server/api-services/common/types"
	models "github.com/gorvk/anything-commerce/server/api-services/models/product"
)

func UpdateProductById(w http.ResponseWriter, r *http.Request) {
	common.CheckHttpResponseType(w, r, http.MethodGet)
	var err error

	d, err := io.ReadAll(r.Body)
	common.HandleHttpError(err, w, constants.ERROR_HTTP_INVALID_REQUEST_BODY, http.StatusBadRequest)

	var payload customTypes.UPDATE_PRODUCT_BY_ID_INPUT

	err = json.Unmarshal(d, &payload)
	common.HandleHttpError(err, w, constants.ERROR_HTTP_UNABLE_TO_PARSE_REQUEST, http.StatusBadRequest)

	_, err = models.UpdateProductById(payload.Id, payload.Field, payload.Value)
	common.HandleDbError(err, w, constants.ERROR_DB_UNABLE_TO_UPDATE_RECORD, http.StatusInternalServerError)
}
