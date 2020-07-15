package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/hdyro/go-restapi-bit/api/auth"
	"github.com/hdyro/go-restapi-bit/api/models"
	"github.com/hdyro/go-restapi-bit/api/responses"
	"github.com/hdyro/go-restapi-bit/api/utils/formaterror"
)

//BASE CRUD
//====================================================================================================================================

func (server *Server) GetCustomerSubTypes(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerSubTypes

	customersubtype := models.CustomerSubType{}

	customersubtypes, err := customersubtype.FindAllCustomerSubType(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = customersubtypes

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCustomerSubType(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerSubType

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	customersubtype := models.CustomerSubType{}

	customersubtypeReceived, err := customersubtype.FindCustomerSubTypeByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = customersubtypeReceived

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCustomerSubTypesStatus(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerSubTypes

	vars := mux.Vars(r)
	status, err := strconv.ParseUint(vars["status"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	customersubtype := models.CustomerSubType{}

	customersubtypeReceived, err := customersubtype.FindAllCustomerSubTypeStatus(server.DB, status)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = customersubtypeReceived

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCustomerSubTypeStatus(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerSubType

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	status, err := strconv.ParseUint(vars["status"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	customersubtype := models.CustomerSubType{}

	customersubtypeReceived, err := customersubtype.FindCustomerSubTypeStatusByID(server.DB, pid, status)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = customersubtypeReceived

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCustomerSubTypeTemps(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerSubTypeTempsPend

	customersubtype := models.CustomerSubTypeTemp{}

	customersubtypes, err := customersubtype.FindAllCustomerSubTypeTemp(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = customersubtypes

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCustomerSubTypeTemp(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerSubTypeTemp

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	customersubtype := models.CustomerSubTypeTemp{}

	customersubtypeReceivedTemp, err := customersubtype.FindCustomerSubTypeTempByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = customersubtypeReceivedTemp

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCustomerSubTypeTempsStatus(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerSubTypeTemps

	vars := mux.Vars(r)
	status, err := strconv.ParseUint(vars["status"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	customersubtype := models.CustomerSubTypeTemp{}

	customersubtypeReceived, err := customersubtype.FindAllCustomerSubTypeTempStatus(server.DB, status)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = customersubtypeReceived

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCustomerSubTypeTempStatus(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerSubTypeTemp

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	status, err := strconv.ParseUint(vars["status"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	customersubtype := models.CustomerSubTypeTemp{}

	customersubtypeReceived, err := customersubtype.FindCustomerSubTypeTempStatusByID(server.DB, pid, status)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = customersubtypeReceived

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCustomerSubTypeTempsPendingActive(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerSubTypeTempsPend

	customersubtype := models.CustomerSubTypeTemp{}

	customersubtypeReceivedTemp, err := customersubtype.FindAllCustomerSubTypeTempStatusPendingActive(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = customersubtypeReceivedTemp

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCustomerSubTypeTempPendingActive(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerSubTypeTempPend

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	customersubtype := models.CustomerSubTypeTemp{}

	customersubtypeReceived, err := customersubtype.FindCustomerSubTypeTempStatusByIDPendingActive(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = customersubtypeReceived

	responses.JSON(w, http.StatusOK, response)
}

//====================================================================================================================================

func (server *Server) CreateCustomerSubTypeNew(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerSubTypeTempCUD

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customersubtypetemp := models.CustomerSubTypeTemp{}
	customersubtypetemppend := models.CustomerSubTypeTempPend{}

	err = json.Unmarshal(body, &customersubtypetemppend)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = json.Unmarshal(body, &customersubtypetemp)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customersubtypetemppend.Prepare()

	customersubtypetemppend.CustomerSubTypeTempStatus = 11
	customersubtypetemppendCreated, err := customersubtypetemppend.SaveCustomerSubTypeTempPend(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Pending Create Customer Sub Type With CustomerSubTypeTempID = %d", customersubtypetemppendCreated.CustomerSubTypeTempID)

	bibiteuserlog := models.BibiteLog{}
	bibiteuserlog.Prepare()
	bibiteuserlog.BibiteUserID = bibiteuserid
	bibiteuserlog.BibiteLogType = bibiteusertype
	bibiteuserlog.BibiteLogAction = "PENDING CREATE"
	bibiteuserlog.BibiteLogDescription = LogDescription
	_, err = bibiteuserlog.SaveBibiteLog(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, customersubtypetemppendCreated.CustomerSubTypeTempID))

	response.Status = http.StatusCreated
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusCreated, response)
}

func (server *Server) CreateCustomerSubTypeConfirm(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerSubTypeCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	customersubtype := models.CustomerSubType{}
	customersubtypetemp := models.CustomerSubTypeTemp{}

	customersubtype.Prepare()
	// err = customersubtype.Validate("insert")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customersubtypetemp.Prepare()
	// err = customersubtype.Validate("insert")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readcustomersubtypeReceivedTemp, err := customersubtypetemp.FindCustomerSubTypeTempByIDPendingActive(server.DB, pid)
	customersubtype.CustomerTypeID = readcustomersubtypeReceivedTemp.CustomerTypeTempID
	customersubtype.CustomerSubTypeName = readcustomersubtypeReceivedTemp.CustomerSubTypeTempName
	customersubtype.CustomerSubTypeCode = readcustomersubtypeReceivedTemp.CustomerSubTypeTempCode
	customersubtype.CustomerSubTypeCreatedBy = readcustomersubtypeReceivedTemp.CustomerSubTypeTempCreatedBy
	customersubtype.CustomerSubTypeStatus = 1
	customersubtypeCreated, err := customersubtype.SaveCustomerSubType(server.DB)

	_, err = customersubtypetemp.DeleteCustomerSubTypeTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Confirm Create Customer Sub Type From CustomerSubTypeTempID = %d To With CustomerSubTypeID = %d", pid, customersubtypeCreated.CustomerSubTypeID)

	bibiteuserlog := models.BibiteLog{}
	bibiteuserlog.Prepare()
	bibiteuserlog.BibiteUserID = bibiteuserid
	bibiteuserlog.BibiteLogType = bibiteusertype
	bibiteuserlog.BibiteLogAction = "CONFIRM CREATE"
	bibiteuserlog.BibiteLogDescription = LogDescription
	_, err = bibiteuserlog.SaveBibiteLog(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, customersubtypeCreated.CustomerSubTypeID))

	response.Status = http.StatusCreated
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusCreated, response)
}

func (server *Server) CreateCustomerSubTypeReject(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerSubTypeTempCUD

	vars := mux.Vars(r)

	customersubtype := models.CustomerSubTypeTemp{}

	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = customersubtype.DeleteCustomerSubTypeTemp(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Reject Create Customer Sub Type From CustomerSubTypeTempID = %d", pid)

	bibiteuserlog := models.BibiteLog{}
	bibiteuserlog.Prepare()
	bibiteuserlog.BibiteUserID = bibiteuserid
	bibiteuserlog.BibiteLogType = bibiteusertype
	bibiteuserlog.BibiteLogAction = "REJECT CREATE"
	bibiteuserlog.BibiteLogDescription = LogDescription
	_, err = bibiteuserlog.SaveBibiteLog(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%d", pid))

	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusOK, response)
}

//====================================================================================================================================

func (server *Server) UpdateCustomerSubTypeNew(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerSubTypeCUD

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

	customersubtype := models.CustomerSubType{}
	customersubtypetemp := models.CustomerSubTypeTemp{}

	err = json.Unmarshal(body, &customersubtypetemp)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customersubtypetemp.Prepare()

	readcustomersubtypeReceived, err := customersubtype.FindCustomerSubTypeByID(server.DB, pid)

	customersubtype.CustomerSubTypeStatus = 13
	_, err = customersubtype.UpdateCustomerSubTypeStatusOnly(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	customersubtypetemp.CustomerSubTypeID = pid
	customersubtypetemp.CustomerSubTypeTempStatus = 13
	customersubtypetemp.CustomerSubTypeTempActionBefore = readcustomersubtypeReceived.CustomerSubTypeStatus
	customersubtypetemp.CustomerSubTypeTempCreatedAt = time.Now()
	customersubtypetempCreated, err := customersubtypetemp.SaveCustomerSubTypeTemp(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Pending Update Customer Sub Type From CustomerSubTypeID = %d To With CustomerSubTypeTempID = %d", pid, customersubtypetempCreated.CustomerSubTypeTempID)

	bibiteuserlog := models.BibiteLog{}
	bibiteuserlog.Prepare()
	bibiteuserlog.BibiteUserID = bibiteuserid
	bibiteuserlog.BibiteLogType = bibiteusertype
	bibiteuserlog.BibiteLogAction = "PENDING UPDATE"
	bibiteuserlog.BibiteLogDescription = LogDescription
	_, err = bibiteuserlog.SaveBibiteLog(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) UpdateCustomerSubTypeConfirm(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerSubTypeCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	customersubtype := models.CustomerSubType{}
	customersubtypetemp := models.CustomerSubTypeTemp{}

	customersubtypetemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readcustomersubtypeReceivedTemp, err := customersubtypetemp.FindCustomerSubTypeTempByID(server.DB, pid)
	customersubtype.CustomerTypeID = readcustomersubtypeReceivedTemp.CustomerTypeTempID
	customersubtype.CustomerSubTypeName = readcustomersubtypeReceivedTemp.CustomerSubTypeTempName
	customersubtype.CustomerSubTypeCode = readcustomersubtypeReceivedTemp.CustomerSubTypeTempCode
	customersubtype.CustomerSubTypeUpdatedBy = readcustomersubtypeReceivedTemp.CustomerSubTypeTempCreatedBy
	customersubtype.CustomerSubTypeStatus = 1
	customersubtypeUpdated, err := customersubtype.UpdateCustomerSubType(server.DB, readcustomersubtypeReceivedTemp.CustomerSubTypeID)

	_, err = customersubtypetemp.DeleteCustomerSubTypeTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Confirm Update Customer Sub Type From CustomerSubTypeTempID = %d To With CustomerSubTypeID = %d", pid, readcustomersubtypeReceivedTemp.CustomerSubTypeID)

	bibiteuserlog := models.BibiteLog{}
	bibiteuserlog.Prepare()
	bibiteuserlog.BibiteUserID = bibiteuserid
	bibiteuserlog.BibiteLogType = bibiteusertype
	bibiteuserlog.BibiteLogAction = "CONFRIM UPDATE"
	bibiteuserlog.BibiteLogDescription = LogDescription
	_, err = bibiteuserlog.SaveBibiteLog(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, customersubtypeUpdated.CustomerSubTypeID))

	response.Status = http.StatusCreated
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusCreated, response)
}

func (server *Server) UpdateCustomerSubTypeReject(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerSubTypeCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	customersubtype := models.CustomerSubType{}
	customersubtypetemp := models.CustomerSubTypeTemp{}

	customersubtype.Prepare()
	customersubtypetemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readcustomersubtypeReceivedTemp, err := customersubtypetemp.FindCustomerSubTypeTempByID(server.DB, pid)
	customersubtype.CustomerSubTypeStatus = readcustomersubtypeReceivedTemp.CustomerSubTypeTempActionBefore
	customersubtypeUpdated, err := customersubtype.UpdateCustomerSubTypeStatusOnly(server.DB, readcustomersubtypeReceivedTemp.CustomerSubTypeID)

	_, err = customersubtypetemp.DeleteCustomerSubTypeTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Reject Update Customer Sub Type From CustomerSubTypeTempID = %d To With CustomerSubTypeID = %d", pid, readcustomersubtypeReceivedTemp.CustomerSubTypeID)

	bibiteuserlog := models.BibiteLog{}
	bibiteuserlog.Prepare()
	bibiteuserlog.BibiteUserID = bibiteuserid
	bibiteuserlog.BibiteLogType = bibiteusertype
	bibiteuserlog.BibiteLogAction = "REJECT UPDATE"
	bibiteuserlog.BibiteLogDescription = LogDescription
	_, err = bibiteuserlog.SaveBibiteLog(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, customersubtypeUpdated.CustomerSubTypeID))

	response.Status = http.StatusCreated
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusCreated, response)
}

//====================================================================================================================================

func (server *Server) InactiveCustomerSubTypeNew(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerSubTypeCUD

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

	customersubtype := models.CustomerSubType{}
	customersubtypetemp := models.CustomerSubTypeTemp{}

	err = json.Unmarshal(body, &customersubtypetemp)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customersubtypetemp.Prepare()

	readcustomersubtypeReceived, err := customersubtype.FindCustomerSubTypeByID(server.DB, pid)

	customersubtype.CustomerSubTypeStatus = 12
	_, err = customersubtype.UpdateCustomerSubTypeStatusOnly(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	customersubtypetemp.CustomerSubTypeID = pid
	customersubtypetemp.CustomerTypeTempID = readcustomersubtypeReceived.CustomerTypeID
	customersubtypetemp.CustomerSubTypeTempCode = readcustomersubtypeReceived.CustomerSubTypeCode
	customersubtypetemp.CustomerSubTypeTempName = readcustomersubtypeReceived.CustomerTypeName
	customersubtypetemp.CustomerSubTypeTempStatus = 12
	customersubtypetemp.CustomerSubTypeTempActionBefore = readcustomersubtypeReceived.CustomerSubTypeStatus
	customersubtypetemp.CustomerSubTypeTempCreatedAt = time.Now()
	customersubtypetempCreated, err := customersubtypetemp.SaveCustomerSubTypeTemp(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Pending Inactive Customer Sub Type From CustomerSubTypeTempID = %d To With CustomerSubTypeID = %d", pid, customersubtypetempCreated.CustomerSubTypeTempID)

	bibiteuserlog := models.BibiteLog{}
	bibiteuserlog.Prepare()
	bibiteuserlog.BibiteUserID = bibiteuserid
	bibiteuserlog.BibiteLogType = bibiteusertype
	bibiteuserlog.BibiteLogAction = "PENDING INACTIVE"
	bibiteuserlog.BibiteLogDescription = LogDescription
	_, err = bibiteuserlog.SaveBibiteLog(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) InactiveCustomerSubTypeConfirm(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerSubTypeCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	customersubtype := models.CustomerSubType{}
	customersubtypetemp := models.CustomerSubTypeTemp{}

	customersubtype.Prepare()
	customersubtypetemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readcustomersubtypeReceivedTemp, err := customersubtypetemp.FindCustomerSubTypeTempByID(server.DB, pid)
	customersubtype.CustomerSubTypeUpdatedBy = readcustomersubtypeReceivedTemp.CustomerSubTypeTempCreatedBy
	customersubtype.CustomerSubTypeStatus = 2
	_, err = customersubtype.UpdateCustomerSubTypeStatus(server.DB, readcustomersubtypeReceivedTemp.CustomerSubTypeID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	_, err = customersubtypetemp.DeleteCustomerSubTypeTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Confirm Inactive Customer Sub Type From CustomerSubTypeTempID = %d To With CustomerSubTypeID = %d", pid, readcustomersubtypeReceivedTemp.CustomerSubTypeID)

	bibiteuserlog := models.BibiteLog{}
	bibiteuserlog.Prepare()
	bibiteuserlog.BibiteUserID = bibiteuserid
	bibiteuserlog.BibiteLogType = bibiteusertype
	bibiteuserlog.BibiteLogAction = "CONFIRM INACTIVE"
	bibiteuserlog.BibiteLogDescription = LogDescription
	_, err = bibiteuserlog.SaveBibiteLog(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) InactiveCustomerSubTypeReject(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerSubTypeCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	customersubtype := models.CustomerSubType{}
	customersubtypetemp := models.CustomerSubTypeTemp{}

	customersubtype.Prepare()
	customersubtypetemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readcustomersubtypeReceivedTemp, err := customersubtypetemp.FindCustomerSubTypeTempByID(server.DB, pid)
	customersubtype.CustomerSubTypeStatus = readcustomersubtypeReceivedTemp.CustomerSubTypeTempActionBefore
	_, err = customersubtype.UpdateCustomerSubTypeStatusOnly(server.DB, readcustomersubtypeReceivedTemp.CustomerSubTypeID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	_, err = customersubtypetemp.DeleteCustomerSubTypeTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Reject Inactive Customer Sub Type From CustomerSubTypeTempID = %d To With CustomerSubTypeID = %d", pid, readcustomersubtypeReceivedTemp.CustomerSubTypeID)

	bibiteuserlog := models.BibiteLog{}
	bibiteuserlog.Prepare()
	bibiteuserlog.BibiteUserID = bibiteuserid
	bibiteuserlog.BibiteLogType = bibiteusertype
	bibiteuserlog.BibiteLogAction = "REJECT INACTIVE"
	bibiteuserlog.BibiteLogDescription = LogDescription
	_, err = bibiteuserlog.SaveBibiteLog(server.DB)
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

func (server *Server) ActiveCustomerSubTypeNew(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerSubTypeCUD

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

	customersubtype := models.CustomerSubType{}
	customersubtypetemp := models.CustomerSubTypeTemp{}

	err = json.Unmarshal(body, &customersubtypetemp)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customersubtypetemp.Prepare()

	readcustomersubtypeReceived, err := customersubtype.FindCustomerSubTypeByID(server.DB, pid)

	customersubtype.CustomerSubTypeStatus = 11
	_, err = customersubtype.UpdateCustomerSubTypeStatusOnly(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	customersubtypetemp.CustomerSubTypeID = pid
	customersubtypetemp.CustomerTypeTempID = readcustomersubtypeReceived.CustomerTypeID
	customersubtypetemp.CustomerSubTypeTempCode = readcustomersubtypeReceived.CustomerSubTypeCode
	customersubtypetemp.CustomerSubTypeTempName = readcustomersubtypeReceived.CustomerTypeName
	customersubtypetemp.CustomerSubTypeTempStatus = 11
	customersubtypetemp.CustomerSubTypeTempActionBefore = readcustomersubtypeReceived.CustomerSubTypeStatus
	customersubtypetemp.CustomerSubTypeTempCreatedAt = time.Now()
	customersubtypetempCreated, err := customersubtypetemp.SaveCustomerSubTypeTemp(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Pending Active Customer Sub Type From CustomerSubTypeTempID = %d To With CustomerSubTypeID = %d", pid, customersubtypetempCreated.CustomerSubTypeTempID)

	bibiteuserlog := models.BibiteLog{}
	bibiteuserlog.Prepare()
	bibiteuserlog.BibiteUserID = bibiteuserid
	bibiteuserlog.BibiteLogType = bibiteusertype
	bibiteuserlog.BibiteLogAction = "PENDING ACTIVE"
	bibiteuserlog.BibiteLogDescription = LogDescription
	_, err = bibiteuserlog.SaveBibiteLog(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) ActiveCustomerSubTypeConfirm(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerSubTypeCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	customersubtype := models.CustomerSubType{}
	customersubtypetemp := models.CustomerSubTypeTemp{}

	customersubtype.Prepare()
	customersubtypetemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readcustomersubtypeReceivedTemp, err := customersubtypetemp.FindCustomerSubTypeTempByID(server.DB, pid)
	customersubtype.CustomerSubTypeUpdatedBy = readcustomersubtypeReceivedTemp.CustomerSubTypeTempCreatedBy
	customersubtype.CustomerSubTypeStatus = 1
	_, err = customersubtype.UpdateCustomerSubTypeStatus(server.DB, readcustomersubtypeReceivedTemp.CustomerSubTypeID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	_, err = customersubtypetemp.DeleteCustomerSubTypeTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Confirm Active Customer Sub Type From CustomerSubTypeTempID = %d To With CustomerSubTypeID = %d", pid, readcustomersubtypeReceivedTemp.CustomerSubTypeID)

	bibiteuserlog := models.BibiteLog{}
	bibiteuserlog.Prepare()
	bibiteuserlog.BibiteUserID = bibiteuserid
	bibiteuserlog.BibiteLogType = bibiteusertype
	bibiteuserlog.BibiteLogAction = "CONFIRM ACTIVE"
	bibiteuserlog.BibiteLogDescription = LogDescription
	_, err = bibiteuserlog.SaveBibiteLog(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) ActiveCustomerSubTypeReject(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerSubTypeCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	customersubtype := models.CustomerSubType{}
	customersubtypetemp := models.CustomerSubTypeTemp{}

	customersubtype.Prepare()
	customersubtypetemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readcustomersubtypeReceivedTemp, err := customersubtypetemp.FindCustomerSubTypeTempByID(server.DB, pid)
	customersubtype.CustomerSubTypeStatus = readcustomersubtypeReceivedTemp.CustomerSubTypeTempActionBefore
	_, err = customersubtype.UpdateCustomerSubTypeStatusOnly(server.DB, readcustomersubtypeReceivedTemp.CustomerSubTypeID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	_, err = customersubtypetemp.DeleteCustomerSubTypeTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Reject Active Customer Sub Type From CustomerSubTypeTempID = %d To With CustomerSubTypeID = %d", pid, readcustomersubtypeReceivedTemp.CustomerSubTypeID)

	bibiteuserlog := models.BibiteLog{}
	bibiteuserlog.Prepare()
	bibiteuserlog.BibiteUserID = bibiteuserid
	bibiteuserlog.BibiteLogType = bibiteusertype
	bibiteuserlog.BibiteLogAction = "REJECT ACTIVE"
	bibiteuserlog.BibiteLogDescription = LogDescription
	_, err = bibiteuserlog.SaveBibiteLog(server.DB)
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

func (server *Server) DeleteCustomerSubTypeNew(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerSubTypeCUD

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

	customersubtype := models.CustomerSubType{}
	customersubtypetemp := models.CustomerSubTypeTemp{}

	err = json.Unmarshal(body, &customersubtypetemp)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customersubtypetemp.Prepare()

	readcustomersubtypeReceived, err := customersubtype.FindCustomerSubTypeByID(server.DB, pid)

	customersubtype.CustomerSubTypeStatus = 14
	_, err = customersubtype.UpdateCustomerSubTypeStatusOnly(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	customersubtypetemp.CustomerSubTypeID = pid
	customersubtypetemp.CustomerTypeTempID = readcustomersubtypeReceived.CustomerTypeID
	customersubtypetemp.CustomerSubTypeTempCode = readcustomersubtypeReceived.CustomerSubTypeCode
	customersubtypetemp.CustomerSubTypeTempName = readcustomersubtypeReceived.CustomerTypeName
	customersubtypetemp.CustomerSubTypeTempStatus = 14
	customersubtypetemp.CustomerSubTypeTempActionBefore = readcustomersubtypeReceived.CustomerSubTypeStatus
	customersubtypetemp.CustomerSubTypeTempCreatedAt = time.Now()
	customersubtypetempCreated, err := customersubtypetemp.SaveCustomerSubTypeTemp(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Pending Delete Customer Sub Type From CustomerSubTypeTempID = %d To With CustomerSubTypeID = %d", pid, customersubtypetempCreated.CustomerSubTypeTempID)

	bibiteuserlog := models.BibiteLog{}
	bibiteuserlog.Prepare()
	bibiteuserlog.BibiteUserID = bibiteuserid
	bibiteuserlog.BibiteLogType = bibiteusertype
	bibiteuserlog.BibiteLogAction = "PENDING DELETE"
	bibiteuserlog.BibiteLogDescription = LogDescription
	_, err = bibiteuserlog.SaveBibiteLog(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) DeleteCustomerSubTypeConfirm(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerSubTypeCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	customersubtype := models.CustomerSubType{}
	customersubtypetemp := models.CustomerSubTypeTemp{}

	customersubtype.Prepare()
	customersubtypetemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readcustomersubtypeReceivedTemp, err := customersubtypetemp.FindCustomerSubTypeTempByID(server.DB, pid)
	customersubtype.CustomerSubTypeDeletedBy = readcustomersubtypeReceivedTemp.CustomerSubTypeTempCreatedBy
	customersubtype.CustomerSubTypeStatus = 3
	_, err = customersubtype.UpdateCustomerSubTypeStatusDelete(server.DB, readcustomersubtypeReceivedTemp.CustomerSubTypeID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	_, err = customersubtypetemp.DeleteCustomerSubTypeTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Confirm Delete Customer Sub Type From CustomerSubTypeTempID = %d To With CustomerSubTypeID = %d", pid, readcustomersubtypeReceivedTemp.CustomerSubTypeID)

	bibiteuserlog := models.BibiteLog{}
	bibiteuserlog.Prepare()
	bibiteuserlog.BibiteUserID = bibiteuserid
	bibiteuserlog.BibiteLogType = bibiteusertype
	bibiteuserlog.BibiteLogAction = "CONFIRM DELETE"
	bibiteuserlog.BibiteLogDescription = LogDescription
	_, err = bibiteuserlog.SaveBibiteLog(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) DeleteCustomerSubTypeReject(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerSubTypeCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	customersubtype := models.CustomerSubType{}
	customersubtypetemp := models.CustomerSubTypeTemp{}

	customersubtype.Prepare()
	customersubtypetemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readcustomersubtypeReceivedTemp, err := customersubtypetemp.FindCustomerSubTypeTempByID(server.DB, pid)
	customersubtype.CustomerSubTypeStatus = readcustomersubtypeReceivedTemp.CustomerSubTypeTempActionBefore
	_, err = customersubtype.UpdateCustomerSubTypeStatusOnly(server.DB, readcustomersubtypeReceivedTemp.CustomerSubTypeID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	_, err = customersubtypetemp.DeleteCustomerSubTypeTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Reject Delete Customer Sub Type From CustomerSubTypeTempID = %d To With CustomerSubTypeID = %d", pid, readcustomersubtypeReceivedTemp.CustomerSubTypeID)

	bibiteuserlog := models.BibiteLog{}
	bibiteuserlog.Prepare()
	bibiteuserlog.BibiteUserID = bibiteuserid
	bibiteuserlog.BibiteLogType = bibiteusertype
	bibiteuserlog.BibiteLogAction = "REJECT DELETE"
	bibiteuserlog.BibiteLogDescription = LogDescription
	_, err = bibiteuserlog.SaveBibiteLog(server.DB)
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

func (server *Server) DeleteCustomerSubType(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerSubTypeCUD

	vars := mux.Vars(r)

	customersubtype := models.CustomerSubType{}

	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = customersubtype.DeleteCustomerSubType(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Delete Customer Sub Type From CustomerSubTypeID = %d", pid)

	bibiteuserlog := models.BibiteLog{}
	bibiteuserlog.Prepare()
	bibiteuserlog.BibiteUserID = bibiteuserid
	bibiteuserlog.BibiteLogType = bibiteusertype
	bibiteuserlog.BibiteLogAction = "DELETE"
	bibiteuserlog.BibiteLogDescription = LogDescription
	_, err = bibiteuserlog.SaveBibiteLog(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%d", pid))

	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) DeleteCustomerSubTypeTemp(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerSubTypeTempCUD

	vars := mux.Vars(r)

	customersubtype := models.CustomerSubTypeTemp{}

	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = customersubtype.DeleteCustomerSubTypeTemp(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Delete Customer Sub Type Temp From CustomerSubTypeTempID = %d", pid)

	bibiteuserlog := models.BibiteLog{}
	bibiteuserlog.Prepare()
	bibiteuserlog.BibiteUserID = bibiteuserid
	bibiteuserlog.BibiteLogType = bibiteusertype
	bibiteuserlog.BibiteLogAction = "DELETE"
	bibiteuserlog.BibiteLogDescription = LogDescription
	_, err = bibiteuserlog.SaveBibiteLog(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%d", pid))

	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusOK, response)
}

//ADDITIONAL
//====================================================================================================================================

func (server *Server) GetCustomerSubTypesByCustomerType(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerSubTypes

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["custtype"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	customersubtype := models.CustomerSubType{}

	customersubtypeReceived, err := customersubtype.FindCustomerSubTypeByCustomerTypeID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = customersubtypeReceived

	responses.JSON(w, http.StatusOK, response)
}
