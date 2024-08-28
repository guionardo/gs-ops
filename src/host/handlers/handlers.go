package handlers

import (
	"net/http"

	"github.com/guionardo/gs-ops/internal/commons"
)

func GetVersionHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = WriteJson(w, commons.Version, http.StatusOK)
	})
}
