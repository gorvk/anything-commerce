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

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	common.CheckHttpResponseType(w, r, http.MethodGet)
	var err error

	rows, err := models.GetAllProducts()
	common.HandleDbError(err, w, constants.ERROR_DB_UNABLE_TO_GET_RECORD, http.StatusInternalServerError)

	var products []customTypes.Product

	defer rows.Close()
	for rows.Next() {
		productRow := customTypes.Product{}
		rows.Scan(
			&productRow.Id,
			&productRow.ProductName,
			&productRow.SellerId,
			&productRow.ShopId,
			&productRow.ProductType,
			&productRow.ProductCondition,
			&productRow.Price,
			&productRow.OriginalPurchasedDate,
			&productRow.OriginalPurchaisingRecieptNo,
			&productRow.ProductDescription)
		products = append(products, productRow)
	}

	var Response customTypes.RESPONSE_PARAMETERS
	Response.Result = products

	d, err := json.Marshal(Response)
	common.HandleHttpError(err, w, constants.ERROR_HTTP_UNABLE_TO_PARSE_RESPONSE, http.StatusInternalServerError)

	_, err = w.Write(d)
	common.HandleHttpError(err, w, constants.ERROR_HTTP_UNABLE_TO_WRITE_RESPONSE, http.StatusInternalServerError)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	common.CheckHttpResponseType(w, r, http.MethodGet)
	var payload customTypes.GET_PRODUCT_BY_ID_INPUT
	var err error

	d, err := io.ReadAll(r.Body)
	common.HandleHttpError(err, w, constants.ERROR_HTTP_INVALID_REQUEST_BODY, http.StatusBadRequest)

	err = json.Unmarshal(d, &payload)
	common.HandleHttpError(err, w, constants.ERROR_HTTP_UNABLE_TO_PARSE_REQUEST, http.StatusBadRequest)

	rows, err := models.GetProductById(payload.Id)
	common.HandleDbError(err, w, constants.ERROR_DB_UNABLE_TO_GET_RECORD, http.StatusInternalServerError)

	defer rows.Close()
	var products []customTypes.Product
	for rows.Next() {
		productRow := customTypes.Product{}
		rows.Scan(
			&productRow.ProductName,
			&productRow.SellerId,
			&productRow.ShopId,
			&productRow.ProductType,
			&productRow.ProductCondition,
			&productRow.Price,
			&productRow.OriginalPurchasedDate,
			&productRow.OriginalPurchaisingRecieptNo,
			&productRow.ProductDescription)
		products = append(products, productRow)
	}

	var Response customTypes.RESPONSE_PARAMETERS
	Response.Result = products

	d, err = json.Marshal(Response)
	common.HandleHttpError(err, w, constants.ERROR_HTTP_UNABLE_TO_PARSE_RESPONSE, http.StatusInternalServerError)

	_, err = w.Write(d)
	common.HandleHttpError(err, w, constants.ERROR_HTTP_UNABLE_TO_WRITE_RESPONSE, http.StatusInternalServerError)
}
