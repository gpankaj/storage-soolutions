package controllers

import (
	"encoding/json"
	"github.com/gpankaj/storage-soolutions/application/services"
	"github.com/gpankaj/storage-soolutions/application/utils"
	"log"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {

	userId := req.URL.Query().Get("user_id")
	intUserId , err := strconv.ParseInt(userId, 10, 64);
	if err != nil {
		userErr := &utils.StorageApplicationError{
			Message: 	"User id must be a number! " + err.Error(),
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		jsonValue, _ := json.Marshal(userErr);
		resp.WriteHeader(userErr.StatusCode);
		resp.Write(jsonValue);
		return;

	}
	log.Println("About to process user id ", intUserId);
	user, appErr :=services.GetUser(intUserId)
	if appErr != nil {
		//Handle the error response.
		jsonValue, _ := json.Marshal(appErr);
		resp.WriteHeader(appErr.StatusCode)

		resp.Write(jsonValue);
		return
	}
	jsonValue, _ := json.Marshal(user);
	resp.Write(jsonValue)

}
