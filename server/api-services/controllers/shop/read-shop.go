package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorvk/anything-commerce/server/api-services/common"
	"github.com/gorvk/anything-commerce/server/api-services/common/constants"
	customTypes "github.com/gorvk/anything-commerce/server/api-services/common/types"
	models "github.com/gorvk/anything-commerce/server/api-services/models/shop"
)

func GetAllShops(w http.ResponseWriter, r *http.Request) {
	var shops []customTypes.Shop
	var err error

	rows, err := models.GetAllShops()
	common.HandleDbError(err, w, constants.ERROR_DB_UNABLE_TO_GET_RECORD, http.StatusInternalServerError)

	defer rows.Close()
	for rows.Next() {
		shopRow := customTypes.Shop{}
		rows.Scan(
			&shopRow.Id,
			&shopRow.ShopName,
			&shopRow.Email,
			&shopRow.PhoneNumber,
			&shopRow.MapLocation,
			&shopRow.ShopType,
			&shopRow.ShopDescription,
			&shopRow.OwnerId)
		shops = append(shops, shopRow)
	}

	var Response customTypes.RESPONSE_PARAMETERS
	Response.Result = shops

	d, err := json.Marshal(Response)
	common.HandleHttpError(err, w, constants.ERROR_HTTP_UNABLE_TO_PARSE_RESPONSE, http.StatusInternalServerError)

	_, err = w.Write(d)
	common.HandleHttpError(err, w, constants.ERROR_HTTP_UNABLE_TO_WRITE_RESPONSE, http.StatusInternalServerError)
}

func GetShopByEmail(w http.ResponseWriter, r *http.Request) {

	common.CheckHttpResponseType(w, r, http.MethodGet)
	var payload customTypes.GET_SHOP_BY_EMAIL_INPUT
	var err error

	d, err := io.ReadAll(r.Body)
	common.HandleHttpError(err, w, constants.ERROR_HTTP_INVALID_REQUEST_BODY, http.StatusBadRequest)

	err = json.Unmarshal(d, &payload)
	common.HandleHttpError(err, w, constants.ERROR_HTTP_UNABLE_TO_PARSE_REQUEST, http.StatusBadRequest)

	rows, err := models.GetShopByEmail(payload.Email)
	common.HandleDbError(err, w, constants.ERROR_DB_UNABLE_TO_GET_RECORD, http.StatusInternalServerError)

	var shops []customTypes.Shop

	defer rows.Close()
	for rows.Next() {
		shopRow := customTypes.Shop{}
		rows.Scan(
			&shopRow.Id,
			&shopRow.ShopName,
			&shopRow.Email,
			&shopRow.PhoneNumber,
			&shopRow.MapLocation,
			&shopRow.ShopType,
			&shopRow.ShopDescription,
			&shopRow.OwnerId)
		shops = append(shops, shopRow)
	}

	var Response customTypes.RESPONSE_PARAMETERS
	Response.Result = shops

	d, err = json.Marshal(Response)
	if err != nil {
		common.HandleHttpError(err, w, constants.ERROR_HTTP_UNABLE_TO_PARSE_RESPONSE, http.StatusInternalServerError)
	}

	_, err = w.Write(d)
	if err != nil {
		common.HandleHttpError(err, w, constants.ERROR_HTTP_UNABLE_TO_WRITE_RESPONSE, http.StatusInternalServerError)
	}
}
