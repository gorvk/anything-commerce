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

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	common.CheckHttpResponseType(w, r, http.MethodGet)
	var err error

	rows, err := models.GetAllUsers()
	common.HandleDbError(err, w, constants.ERROR_DB_UNABLE_TO_GET_RECORD, http.StatusInternalServerError)

	var users []customTypes.User

	defer rows.Close()
	for rows.Next() {
		userRow := customTypes.User{}
		rows.Scan(
			&userRow.Id,
			&userRow.FirstName,
			&userRow.LastName,
			&userRow.Email,
			&userRow.PhoneNumber,
			&userRow.UserAddress,
			&userRow.IsSeller,
		)
		users = append(users, userRow)
	}

	var Response customTypes.RESPONSE_PARAMETERS
	Response.Result = users

	d, err := json.Marshal(Response)
	common.HandleHttpError(err, w, constants.ERROR_HTTP_UNABLE_TO_PARSE_RESPONSE, http.StatusInternalServerError)

	_, err = w.Write(d)
	common.HandleHttpError(err, w, constants.ERROR_HTTP_UNABLE_TO_WRITE_RESPONSE, http.StatusInternalServerError)
}

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	common.CheckHttpResponseType(w, r, http.MethodGet)
	var payload customTypes.GET_USER_BY_EMAIL_INPUT
	var err error

	d, err := io.ReadAll(r.Body)
	common.HandleHttpError(err, w, constants.ERROR_HTTP_INVALID_REQUEST_BODY, http.StatusBadRequest)

	err = json.Unmarshal(d, &payload)
	common.HandleHttpError(err, w, constants.ERROR_HTTP_UNABLE_TO_PARSE_REQUEST, http.StatusBadRequest)

	rows, err := models.GetUserByEmail(payload.Email)
	common.HandleDbError(err, w, constants.ERROR_DB_UNABLE_TO_GET_RECORD, http.StatusInternalServerError)

	var users []customTypes.User

	defer rows.Close()
	for rows.Next() {
		userRow := customTypes.User{}
		rows.Scan(
			&userRow.Id,
			&userRow.FirstName,
			&userRow.LastName,
			&userRow.Email,
			&userRow.PhoneNumber,
			&userRow.UserAddress,
			&userRow.IsSeller,
		)
		users = append(users, userRow)
	}

	var Response customTypes.RESPONSE_PARAMETERS
	Response.Result = users

	d, err = json.Marshal(Response)
	if err != nil {
		common.HandleHttpError(err, w, constants.ERROR_HTTP_UNABLE_TO_PARSE_RESPONSE, http.StatusInternalServerError)
	}

	_, err = w.Write(d)
	if err != nil {
		common.HandleHttpError(err, w, constants.ERROR_HTTP_UNABLE_TO_WRITE_RESPONSE, http.StatusInternalServerError)
	}
}
