package domain

import (
	"fmt"
	"github.com/gpankaj/storage-soolutions/application/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		123 : {Id: 123, FirstName: "Rama", LastName: "Shyma", Email: "rama.shyma@gmail.com"},
	}
)
func GetUser(userId int64)(*User, *utils.StorageApplicationError)  {
	user := users[userId]
	if user == nil {
		//return nil, errors.New(fmt.Sprintf("user %v was not found ", userId))

		return nil, &utils.StorageApplicationError{
			Message: fmt.Sprintf("user %v was not found ", userId),
			StatusCode: http.StatusNotFound,
			Code: "not_found",
		}
	}
	return user, nil
}
