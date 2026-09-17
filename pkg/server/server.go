package server

import (
	"net/http"

	"github.com/Detaleryn/final/pkg/api"
)

// Init регистрирует обработчики API
func Init() {

	api.Init()

	http.Handle("/", http.FileServer(http.Dir("./web")))

}
