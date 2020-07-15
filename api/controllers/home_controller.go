package controllers

import (
	"net/http"

	"github.com/hdyro/go-restapi-bit/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API Bibite International Technology Version 0.6")
}
