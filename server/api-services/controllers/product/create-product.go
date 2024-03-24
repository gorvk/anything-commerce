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

func CreateNewProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct customTypes.Product
	var err error

	d, err := io.ReadAll(r.Body)
	common.HandleHttpError(err, w, constants.ERROR_HTTP_INVALID_REQUEST_BODY, http.StatusBadRequest)

	err = json.Unmarshal(d, &newProduct)
	common.HandleHttpError(err, w, constants.ERROR_HTTP_INVALID_REQUEST_BODY, http.StatusBadRequest)

	_, err = models.CreateNewProduct(newProduct)
	common.HandleDbError(err, w, constants.ERROR_DB_UNABLE_TO_CREATE_RECORD, http.StatusInternalServerError)
}
