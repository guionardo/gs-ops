package routes

import (
	"net/http"

	"github.com/guionardo/gs-ops/src/host/handlers"
)

func SetupAPI(routes *http.ServeMux) {
	routes.Handle("/version", handlers.GetVersionHandler())
}
