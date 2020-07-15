package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hdyro/go-restapi-bit/api/models"
	"github.com/hdyro/go-restapi-bit/api/responses"
)

func (server *Server) GetVillages(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseVillages

	village := models.Village{}

	villages, err := village.FindAllVillages(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = villages

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetVillage(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseVillage

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	village := models.Village{}

	villageReceived, err := village.FindVillageByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data.VillageID = villageReceived.VillageID
	response.Data.DistrictID = villageReceived.DistrictID
	response.Data.VillageName = villageReceived.VillageName

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetVillageDistrict(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseVillages

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	village := models.Village{}

	villageReceived, err := village.FindVillageDistrictByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = villageReceived

	responses.JSON(w, http.StatusOK, response)
}
