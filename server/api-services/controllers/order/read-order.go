package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorvk/anything-commerce/server/api-services/common"
	"github.com/gorvk/anything-commerce/server/api-services/common/constants"
	customTypes "github.com/gorvk/anything-commerce/server/api-services/common/types"
	models "github.com/gorvk/anything-commerce/server/api-services/models/order"
)

func GetAllOrders(w http.ResponseWriter, r *http.Request) {
	common.CheckHttpResponseType(w, r, http.MethodGet)
	var err error

	rows, err := models.GetAllOrders()
	common.HandleDbError(err, w, constants.ERROR_DB_UNABLE_TO_GET_RECORD, http.StatusInternalServerError)

	var orders []customTypes.Order

	defer rows.Close()
	for rows.Next() {
		orderRow := customTypes.Order{}
		rows.Scan(
			&orderRow.Id,
			&orderRow.FromMapLocation,
			&orderRow.ToMapLocation,
			&orderRow.LastStopMapLocation,
			&orderRow.OrderStatus,
			&orderRow.PaymentStatus,
			&orderRow.ProductId,
			&orderRow.BuyerId,
			&orderRow.ShopId)
		orders = append(orders, orderRow)
	}

	var Response customTypes.RESPONSE_PARAMETERS
	Response.Result = orders

	d, err := json.Marshal(Response)
	common.HandleHttpError(err, w, constants.ERROR_HTTP_UNABLE_TO_PARSE_RESPONSE, http.StatusInternalServerError)

	_, err = w.Write(d)
	common.HandleHttpError(err, w, constants.ERROR_HTTP_UNABLE_TO_WRITE_RESPONSE, http.StatusInternalServerError)
}

func GetOrderById(w http.ResponseWriter, r *http.Request) {
	common.CheckHttpResponseType(w, r, http.MethodGet)
	var payload customTypes.GET_ORDER_BY_ID_INPUT
	var err error

	d, err := io.ReadAll(r.Body)
	common.HandleHttpError(err, w, constants.ERROR_HTTP_INVALID_REQUEST_BODY, http.StatusBadRequest)

	err = json.Unmarshal(d, &payload)
	common.HandleHttpError(err, w, constants.ERROR_HTTP_UNABLE_TO_PARSE_REQUEST, http.StatusBadRequest)

	rows, err := models.GetOrderById(payload.Id)
	common.HandleDbError(err, w, constants.ERROR_DB_UNABLE_TO_GET_RECORD, http.StatusInternalServerError)

	var orders []customTypes.Order

	defer rows.Close()
	for rows.Next() {
		orderRow := customTypes.Order{}
		rows.Scan(
			&orderRow.Id,
			&orderRow.FromMapLocation,
			&orderRow.ToMapLocation,
			&orderRow.LastStopMapLocation,
			&orderRow.OrderStatus,
			&orderRow.PaymentStatus,
			&orderRow.ProductId,
			&orderRow.BuyerId,
			&orderRow.ShopId)
		orders = append(orders, orderRow)
	}

	var Response customTypes.RESPONSE_PARAMETERS
	Response.Result = orders

	d, err = json.Marshal(Response)
	common.HandleHttpError(err, w, constants.ERROR_HTTP_UNABLE_TO_PARSE_RESPONSE, http.StatusInternalServerError)

	_, err = w.Write(d)
	common.HandleHttpError(err, w, constants.ERROR_HTTP_UNABLE_TO_WRITE_RESPONSE, http.StatusInternalServerError)
}
