package controllers

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hdyro/go-restapi-bit/api/models"
	"github.com/hdyro/go-restapi-bit/api/responses"
)

func RandomString(n int) string {
	var letter = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func (server *Server) GetCountries(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCountries

	country := models.Country{}

	countries, err := country.FindAllCountries(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = countries

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCountry(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCountry

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	country := models.Country{}

	countryReceived, err := country.FindCountryByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data.CountryID = countryReceived.CountryID
	response.Data.CountryName = countryReceived.CountryName

	responses.JSON(w, http.StatusOK, response)
}
