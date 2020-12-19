package app

import (
	"github.com/gpankaj/storage-soolutions/application/controllers"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser);

	if err:=http.ListenAndServe("0.0.0.0:8080", nil); err!=nil {
		panic(err)
	}
}