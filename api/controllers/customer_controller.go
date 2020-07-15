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

func (server *Server) CreateCustomer(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerIU

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customer := models.Customer{}
	err = json.Unmarshal(body, &customer)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customer.Prepare()
	err = customer.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customerCreated, err := customer.SaveCustomer(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, customerCreated.CustomerID))

	response.Status = http.StatusCreated
	response.Message = "SUCCESS"
	response.Data = customerCreated

	responses.JSON(w, http.StatusCreated, response)
}

func (server *Server) GetCustomers(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomers

	customer := models.Customer{}

	customers, err := customer.FindAllCustomer(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = customers

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCustomer(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomer

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	customer := models.Customer{}

	customerReceived, err := customer.FindCustomerByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = customerReceived

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) UpdateCustomer(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerIU

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
	customer := models.Customer{}
	err = json.Unmarshal(body, &customer)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customer.Prepare()
	err = customer.Validate("setpin")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customerUpdated, err := customer.UpdateCustomer(server.DB, pid)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = customerUpdated

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) DeleteCustomer(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerDel

	vars := mux.Vars(r)

	customer := models.Customer{}

	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = customer.DeleteCustomer(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))

	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusOK, response)
}
