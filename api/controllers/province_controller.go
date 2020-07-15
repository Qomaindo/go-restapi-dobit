package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hdyro/go-restapi-bit/api/models"
	"github.com/hdyro/go-restapi-bit/api/responses"
)

func (server *Server) GetProvinces(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseProvinces

	province := models.Province{}

	provinces, err := province.FindAllProvinces(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = provinces

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetProvince(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseProvince

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	province := models.Province{}

	provinceReceived, err := province.FindProvinceByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data.ProvinceID = provinceReceived.ProvinceID
	response.Data.CountryID = provinceReceived.CountryID
	response.Data.ProvinceName = provinceReceived.ProvinceName
	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetProvinceCountry(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseProvinces

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	province := models.Province{}

	provinceReceived, err := province.FindProvinceCountryByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = provinceReceived

	responses.JSON(w, http.StatusOK, response)
}
