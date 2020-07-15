package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hdyro/go-restapi-bit/api/models"
	"github.com/hdyro/go-restapi-bit/api/responses"
	"github.com/hdyro/go-restapi-bit/api/utils/formaterror"
)

// BASE CRUD
// =====================================================================================

func (server *Server) GetCustomerUserLogs(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerLogs

	customerlogs := models.CustomerLog{}

	customerLog, err := customerlogs.FindAllCustomerLogs(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = customerLog

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCustomerUserLog(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerLog

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	customerlogs := models.CustomerLog{}

	customerlogsReceived, err := customerlogs.FindCustomerLogByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = customerlogsReceived

	responses.JSON(w, http.StatusOK, response)
}

// ==============================================================================================

func (server *Server) CreateCustomerUserLog(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerLogCUD

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customerlogs := models.CustomerLog{}
	err = json.Unmarshal(body, &customerlogs)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customerlogs.Prepare()
	err = customerlogs.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customerLogCreated, err := customerlogs.SaveCustomerLog(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, customerLogCreated.CustomerLogID))

	response.Status = http.StatusCreated
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusCreated, response)
}

// ==============================================================================================

func (server *Server) UpdateCustomerUserLog(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerLogCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customerlogs := models.CustomerLog{}
	err = json.Unmarshal(body, &customerlogs)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customerlogs.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	_, err = customerlogs.UpdateCustomerLog(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusOK, response)
}

// ==============================================================================================

func (server *Server) DeleteCustomerUserLog(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerLogCUD

	vars := mux.Vars(r)

	customerlogs := models.CustomerLog{}

	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = customerlogs.DeleteCustomerLog(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))

	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusOK, response)
}
