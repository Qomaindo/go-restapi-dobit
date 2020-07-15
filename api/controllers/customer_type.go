package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/hdyro/go-restapi-bit/api/models"
	"github.com/hdyro/go-restapi-bit/api/responses"
	"github.com/hdyro/go-restapi-bit/api/utils/formaterror"
)

//BASE CRUD
//====================================================================================================================================

func (server *Server) GetCustomerTypes(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerTypes

	customertype := models.CustomerType{}

	customertypes, err := customertype.FindAllCustomerType(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = customertypes

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCustomerType(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerType

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	customertype := models.CustomerType{}

	customertypeReceived, err := customertype.FindCustomerTypeByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = customertypeReceived

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCustomerTypesStatus(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerTypes

	vars := mux.Vars(r)
	status, err := strconv.ParseUint(vars["status"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	customertype := models.CustomerType{}

	customertypeReceived, err := customertype.FindAllCustomerTypeStatus(server.DB, status)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = customertypeReceived

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCustomerTypeStatus(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerType

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	status, err := strconv.ParseUint(vars["status"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	customertype := models.CustomerType{}

	customertypeReceived, err := customertype.FindCustomerTypeStatusByID(server.DB, pid, status)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = customertypeReceived

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCustomerTypeTemps(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerTypeTempsPend

	customertype := models.CustomerTypeTemp{}

	customertypes, err := customertype.FindAllCustomerTypeTemp(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = customertypes

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCustomerTypeTemp(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerTypeTemp

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	customertype := models.CustomerTypeTemp{}

	customertypeReceivedTemp, err := customertype.FindCustomerTypeTempByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = customertypeReceivedTemp

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCustomerTypeTempsStatus(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerTypeTemps

	vars := mux.Vars(r)
	status, err := strconv.ParseUint(vars["status"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	customertype := models.CustomerTypeTemp{}

	customertypeReceived, err := customertype.FindAllCustomerTypeTempStatus(server.DB, status)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = customertypeReceived

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCustomerTypeTempStatus(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerTypeTemp

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	status, err := strconv.ParseUint(vars["status"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	customertype := models.CustomerTypeTemp{}

	customertypeReceived, err := customertype.FindCustomerTypeTempStatusByID(server.DB, pid, status)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = customertypeReceived

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCustomerTypeTempsPendingActive(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerTypeTempsPend

	customertype := models.CustomerTypeTemp{}

	customertypeReceivedTemp, err := customertype.FindAllCustomerTypeTempPendingActive(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = customertypeReceivedTemp

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCustomerTypeTempPendingActive(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerTypeTempPend

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	customertype := models.CustomerTypeTemp{}

	customertypeReceived, err := customertype.FindCustomerTypeTempByIDPendingActive(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = customertypeReceived

	responses.JSON(w, http.StatusOK, response)
}

//====================================================================================================================================

func (server *Server) CreateCustomerTypeNew(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerTypeTempCUD

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customertypetemp := models.CustomerTypeTemp{}
	customertypetemppend := models.CustomerTypeTempPend{}

	err = json.Unmarshal(body, &customertypetemppend)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = json.Unmarshal(body, &customertypetemp)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customertypetemppend.Prepare()
	// err = customertypetemp.Validate("insertupdate")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customertypetemppend.CustomerTypeTempStatus = 11
	customertypetemppendCreated, err := customertypetemppend.SaveCustomerTypeTempPend(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, customertypetemppendCreated.CustomerTypeTempID))

	response.Status = http.StatusCreated
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusCreated, response)
}

func (server *Server) CreateCustomerTypeConfirm(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerTypeCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	customertype := models.CustomerType{}
	customertypetemp := models.CustomerTypeTemp{}

	customertype.Prepare()
	// err = customertype.Validate("insert")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customertypetemp.Prepare()
	// err = customertype.Validate("insert")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readcustomertypeReceivedTemp, err := customertypetemp.FindCustomerTypeTempByIDPendingActive(server.DB, pid)
	customertype.CustomerTypeCode = readcustomertypeReceivedTemp.CustomerTypeTempCode
	customertype.CustomerTypeName = readcustomertypeReceivedTemp.CustomerTypeTempName
	customertype.CustomerTypeCreatedBy = readcustomertypeReceivedTemp.CustomerTypeTempCreatedBy
	customertype.CustomerTypeStatus = 1
	customertypeCreated, err := customertype.SaveCustomerType(server.DB)

	_, err = customertypetemp.DeleteCustomerTypeTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, customertypeCreated.CustomerTypeID))

	response.Status = http.StatusCreated
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusCreated, response)
}

func (server *Server) CreateCustomerTypeReject(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerTypeTempCUD

	vars := mux.Vars(r)

	customertype := models.CustomerTypeTemp{}

	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = customertype.DeleteCustomerTypeTemp(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))

	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusOK, response)
}

//====================================================================================================================================

func (server *Server) UpdateCustomerTypeNew(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerTypeCUD

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

	customertype := models.CustomerType{}
	customertypetemp := models.CustomerTypeTemp{}

	err = json.Unmarshal(body, &customertypetemp)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customertypetemp.Prepare()
	// err = customertypetemp.Validate("insertupdate")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readcustomertypeReceived, err := customertype.FindCustomerTypeByID(server.DB, pid)

	customertype.CustomerTypeStatus = 13
	_, err = customertype.UpdateCustomerTypeStatusOnly(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	customertypetemp.CustomerTypeID = pid
	customertypetemp.CustomerTypeTempStatus = 13
	customertypetemp.CustomerTypeTempActionBefore = readcustomertypeReceived.CustomerTypeStatus
	customertypetemp.CustomerTypeTempCreatedAt = time.Now()
	_, err = customertypetemp.SaveCustomerTypeTemp(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) UpdateCustomerTypeConfirm(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerTypeCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	customertype := models.CustomerType{}
	customertypetemp := models.CustomerTypeTemp{}

	customertypetemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readcustomertypeReceivedTemp, err := customertypetemp.FindCustomerTypeTempByID(server.DB, pid)
	customertype.CustomerTypeCode = readcustomertypeReceivedTemp.CustomerTypeTempCode
	customertype.CustomerTypeName = readcustomertypeReceivedTemp.CustomerTypeTempName
	customertype.CustomerTypeUpdatedBy = readcustomertypeReceivedTemp.CustomerTypeTempCreatedBy
	customertype.CustomerTypeStatus = 1
	customertypeUpdated, err := customertype.UpdateCustomerType(server.DB, readcustomertypeReceivedTemp.CustomerTypeID)

	_, err = customertypetemp.DeleteCustomerTypeTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, customertypeUpdated.CustomerTypeID))

	response.Status = http.StatusCreated
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusCreated, response)
}

func (server *Server) UpdateCustomerTypeReject(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerTypeCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	customertype := models.CustomerType{}
	customertypetemp := models.CustomerTypeTemp{}

	customertype.Prepare()
	customertypetemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readcustomertypeReceivedTemp, err := customertypetemp.FindCustomerTypeTempByID(server.DB, pid)
	customertype.CustomerTypeStatus = readcustomertypeReceivedTemp.CustomerTypeTempActionBefore
	customertypeUpdated, err := customertype.UpdateCustomerTypeStatusOnly(server.DB, readcustomertypeReceivedTemp.CustomerTypeID)

	_, err = customertypetemp.DeleteCustomerTypeTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, customertypeUpdated.CustomerTypeID))

	response.Status = http.StatusCreated
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusCreated, response)
}

//====================================================================================================================================

func (server *Server) InactiveCustomerTypeNew(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerTypeCUD

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

	customertype := models.CustomerType{}
	customertypetemp := models.CustomerTypeTemp{}

	err = json.Unmarshal(body, &customertypetemp)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customertypetemp.Prepare()
	// err = customertypetemp.Validate("reason")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readcustomertypeReceived, err := customertype.FindCustomerTypeByID(server.DB, pid)

	customertype.CustomerTypeStatus = 12
	_, err = customertype.UpdateCustomerTypeStatusOnly(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	customertypetemp.CustomerTypeID = pid
	customertypetemp.CustomerTypeTempCode = readcustomertypeReceived.CustomerTypeCode
	customertypetemp.CustomerTypeTempName = readcustomertypeReceived.CustomerTypeName
	customertypetemp.CustomerTypeTempStatus = 12
	customertypetemp.CustomerTypeTempActionBefore = readcustomertypeReceived.CustomerTypeStatus
	customertypetemp.CustomerTypeTempCreatedAt = time.Now()
	_, err = customertypetemp.SaveCustomerTypeTemp(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) InactiveCustomerTypeConfirm(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerTypeCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	customertype := models.CustomerType{}
	customertypetemp := models.CustomerTypeTemp{}

	customertype.Prepare()
	customertypetemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readcustomertypeReceivedTemp, err := customertypetemp.FindCustomerTypeTempByID(server.DB, pid)
	customertype.CustomerTypeUpdatedBy = readcustomertypeReceivedTemp.CustomerTypeTempCreatedBy
	customertype.CustomerTypeStatus = 2
	_, err = customertype.UpdateCustomerTypeStatus(server.DB, readcustomertypeReceivedTemp.CustomerTypeID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	_, err = customertypetemp.DeleteCustomerTypeTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) InactiveCustomerTypeReject(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerTypeCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	customertype := models.CustomerType{}
	customertypetemp := models.CustomerTypeTemp{}

	customertype.Prepare()
	customertypetemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readcustomertypeReceivedTemp, err := customertypetemp.FindCustomerTypeTempByID(server.DB, pid)
	customertype.CustomerTypeStatus = readcustomertypeReceivedTemp.CustomerTypeTempActionBefore
	_, err = customertype.UpdateCustomerTypeStatusOnly(server.DB, readcustomertypeReceivedTemp.CustomerTypeID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	_, err = customertypetemp.DeleteCustomerTypeTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusOK, response)
}

//====================================================================================================================================

func (server *Server) ActiveCustomerTypeNew(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerTypeCUD

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

	customertype := models.CustomerType{}
	customertypetemp := models.CustomerTypeTemp{}

	err = json.Unmarshal(body, &customertypetemp)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customertypetemp.Prepare()
	// err = customertypetemp.Validate("reason")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readcustomertypeReceived, err := customertype.FindCustomerTypeByID(server.DB, pid)

	customertype.CustomerTypeStatus = 11
	_, err = customertype.UpdateCustomerTypeStatusOnly(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	customertypetemp.CustomerTypeID = pid
	customertypetemp.CustomerTypeTempCode = readcustomertypeReceived.CustomerTypeCode
	customertypetemp.CustomerTypeTempName = readcustomertypeReceived.CustomerTypeName
	customertypetemp.CustomerTypeTempStatus = 11
	customertypetemp.CustomerTypeTempActionBefore = readcustomertypeReceived.CustomerTypeStatus
	customertypetemp.CustomerTypeTempCreatedAt = time.Now()
	_, err = customertypetemp.SaveCustomerTypeTemp(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) ActiveCustomerTypeConfirm(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerTypeCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	customertype := models.CustomerType{}
	customertypetemp := models.CustomerTypeTemp{}

	customertype.Prepare()
	customertypetemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readcustomertypeReceivedTemp, err := customertypetemp.FindCustomerTypeTempByID(server.DB, pid)
	customertype.CustomerTypeUpdatedBy = readcustomertypeReceivedTemp.CustomerTypeTempCreatedBy
	customertype.CustomerTypeStatus = 1
	_, err = customertype.UpdateCustomerTypeStatus(server.DB, readcustomertypeReceivedTemp.CustomerTypeID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	_, err = customertypetemp.DeleteCustomerTypeTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) ActiveCustomerTypeReject(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerTypeCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	customertype := models.CustomerType{}
	customertypetemp := models.CustomerTypeTemp{}

	customertype.Prepare()
	customertypetemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readcustomertypeReceivedTemp, err := customertypetemp.FindCustomerTypeTempByID(server.DB, pid)
	customertype.CustomerTypeStatus = readcustomertypeReceivedTemp.CustomerTypeTempActionBefore
	_, err = customertype.UpdateCustomerTypeStatusOnly(server.DB, readcustomertypeReceivedTemp.CustomerTypeID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	_, err = customertypetemp.DeleteCustomerTypeTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusOK, response)
}

//====================================================================================================================================

func (server *Server) DeleteCustomerTypeNew(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerTypeCUD

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

	customertype := models.CustomerType{}
	customertypetemp := models.CustomerTypeTemp{}

	err = json.Unmarshal(body, &customertypetemp)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customertypetemp.Prepare()
	// err = customertypetemp.Validate("reason")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readcustomertypeReceived, err := customertype.FindCustomerTypeByID(server.DB, pid)

	customertype.CustomerTypeStatus = 14
	_, err = customertype.UpdateCustomerTypeStatusOnly(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	customertypetemp.CustomerTypeID = pid
	customertypetemp.CustomerTypeTempCode = readcustomertypeReceived.CustomerTypeCode
	customertypetemp.CustomerTypeTempName = readcustomertypeReceived.CustomerTypeName
	customertypetemp.CustomerTypeTempStatus = 14
	customertypetemp.CustomerTypeTempActionBefore = readcustomertypeReceived.CustomerTypeStatus
	customertypetemp.CustomerTypeTempCreatedAt = time.Now()
	_, err = customertypetemp.SaveCustomerTypeTemp(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) DeleteCustomerTypeConfirm(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerTypeCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	customertype := models.CustomerType{}
	customertypetemp := models.CustomerTypeTemp{}

	customertype.Prepare()
	customertypetemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readcustomertypeReceivedTemp, err := customertypetemp.FindCustomerTypeTempByID(server.DB, pid)
	customertype.CustomerTypeDeletedBy = readcustomertypeReceivedTemp.CustomerTypeTempCreatedBy
	customertype.CustomerTypeStatus = 3
	_, err = customertype.UpdateCustomerTypeStatusDelete(server.DB, readcustomertypeReceivedTemp.CustomerTypeID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	_, err = customertypetemp.DeleteCustomerTypeTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) DeleteCustomerTypeReject(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerTypeCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	customertype := models.CustomerType{}
	customertypetemp := models.CustomerTypeTemp{}

	customertype.Prepare()
	customertypetemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readcustomertypeReceivedTemp, err := customertypetemp.FindCustomerTypeTempByID(server.DB, pid)
	customertype.CustomerTypeStatus = readcustomertypeReceivedTemp.CustomerTypeTempActionBefore
	_, err = customertype.UpdateCustomerTypeStatusOnly(server.DB, readcustomertypeReceivedTemp.CustomerTypeID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	_, err = customertypetemp.DeleteCustomerTypeTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusOK, response)
}

//====================================================================================================================================

func (server *Server) DeleteCustomerType(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerTypeCUD

	vars := mux.Vars(r)

	customertype := models.CustomerType{}

	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = customertype.DeleteCustomerType(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))

	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) DeleteCustomerTypeTemp(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerTypeTempCUD

	vars := mux.Vars(r)

	customertype := models.CustomerTypeTemp{}

	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = customertype.DeleteCustomerTypeTemp(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))

	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusOK, response)
}

//ADDITIONAL
//====================================================================================================================================
