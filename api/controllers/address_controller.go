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

func (server *Server) GetAddresss(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseAddresss

	address := models.Address{}

	addresss, err := address.FindAllAddress(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = addresss

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetAddress(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseAddress

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	address := models.Address{}

	addressReceived, err := address.FindAddressByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = addressReceived

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetAddresssStatus(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseAddresss

	vars := mux.Vars(r)
	status, err := strconv.ParseUint(vars["status"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	address := models.Address{}

	addressReceived, err := address.FindAllAddressStatus(server.DB, status)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = addressReceived

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetAddressStatus(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseAddress

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	status, err := strconv.ParseUint(vars["status"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	address := models.Address{}

	addressReceived, err := address.FindAddressStatusByID(server.DB, pid, status)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = addressReceived

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetAddressTemps(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseAddressTempsPend

	address := models.AddressTemp{}

	addresss, err := address.FindAllAddressTemp(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = addresss

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetAddressTemp(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseAddressTemp

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	address := models.AddressTemp{}

	addressReceivedTemp, err := address.FindAddressTempByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = addressReceivedTemp

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetAddressTempsStatus(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseAddressTemps

	vars := mux.Vars(r)
	status, err := strconv.ParseUint(vars["status"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	address := models.AddressTemp{}

	addressReceived, err := address.FindAllAddressTempStatus(server.DB, status)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = addressReceived

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetAddressTempStatus(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseAddressTemp

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	status, err := strconv.ParseUint(vars["status"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	address := models.AddressTemp{}

	addressReceived, err := address.FindAddressTempStatusByID(server.DB, pid, status)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = addressReceived

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetAddressTempsPendingActive(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseAddressTempsPend

	address := models.AddressTemp{}

	addressReceived, err := address.FindAllAddressTempPendingActive(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = addressReceived

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetAddressTempPendingActive(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseAddressTempPend

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	address := models.AddressTemp{}

	addressReceived, err := address.FindAddressTempByIDPendingActive(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = addressReceived

	responses.JSON(w, http.StatusOK, response)
}

//====================================================================================================================================

func (server *Server) CreateAddressNew(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseAddressTempCUD

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	addresstemp := models.AddressTemp{}
	addresstemppend := models.AddressTempPend{}

	err = json.Unmarshal(body, &addresstemppend)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = json.Unmarshal(body, &addresstemp)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	addresstemppend.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	addresstemppend.AddressTempStatus = 11
	addresstemppendCreated, err := addresstemppend.SaveAddressTempPend(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Pending Create Address With AddressTempID = %d", *addresstemppendCreated.AddressTempID)

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

	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, addresstemppendCreated.AddressTempID))

	response.Status = http.StatusCreated
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusCreated, response)
}

func (server *Server) CreateAddressConfirm(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseAddressCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	address := models.Address{}
	addresstemp := models.AddressTemp{}

	address.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	addresstemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readaddressReceivedTemp, err := addresstemp.FindAddressTempByIDPendingActive(server.DB, pid)
	address.CountryID = readaddressReceivedTemp.CountryTempID
	address.ProvinceID = readaddressReceivedTemp.ProvinceTempID
	address.RegencyID = readaddressReceivedTemp.RegencyTempID
	address.DistrictID = readaddressReceivedTemp.DistrictTempID
	address.VillageID = readaddressReceivedTemp.VillageTempID
	address.AddressCreatedBy = readaddressReceivedTemp.AddressTempCreatedBy
	address.AddressStatus = 1
	addressCreated, err := address.SaveAddress(server.DB)

	_, err = addresstemp.DeleteAddressTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Confirm Create Address From Address Temp = %d To With AddressID = %d", pid, *addressCreated.AddressID)

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

	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, addressCreated.AddressID))

	response.Status = http.StatusCreated
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusCreated, response)
}

func (server *Server) CreateAddressReject(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseAddressTempCUD

	vars := mux.Vars(r)

	address := models.AddressTemp{}

	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = address.DeleteAddressTemp(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Reject Create Address With AddressTempID = %d", pid)

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

func (server *Server) UpdateAddressNew(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseAddressCUD

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

	address := models.Address{}
	addresstemp := models.AddressTemp{}

	err = json.Unmarshal(body, &addresstemp)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	addresstemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readaddressReceived, err := address.FindAddressByID(server.DB, pid)

	address.AddressStatus = 13
	_, err = address.UpdateAddressStatusOnly(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	addresstemp.AddressID = pid
	addresstemp.AddressTempStatus = 13
	addresstemp.AddressTempActionBefore = readaddressReceived.AddressStatus
	addresstemp.AddressTempCreatedAt = time.Now()
	addresstempCreated, err := addresstemp.SaveAddressTemp(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Pending Update Address From AddressID = %d to AddressTempID = %d", pid, addresstempCreated.AddressTempID)

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

func (server *Server) UpdateAddressConfirm(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseAddressCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	address := models.Address{}
	addresstemp := models.AddressTemp{}

	addresstemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readaddressReceivedTemp, err := addresstemp.FindAddressTempByID(server.DB, pid)
	address.CountryID = readaddressReceivedTemp.CountryTempID
	address.ProvinceID = readaddressReceivedTemp.ProvinceTempID
	address.RegencyID = readaddressReceivedTemp.RegencyTempID
	address.DistrictID = readaddressReceivedTemp.DistrictTempID
	address.VillageID = readaddressReceivedTemp.VillageTempID
	address.AddressUpdatedBy = readaddressReceivedTemp.AddressTempCreatedBy
	address.AddressStatus = 1
	addressUpdated, err := address.UpdateAddress(server.DB, readaddressReceivedTemp.AddressID)

	_, err = addresstemp.DeleteAddressTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Confirm Update Address From Address Temp = %d To With AddressID = %d", pid, readaddressReceivedTemp.AddressID)

	bibiteuserlog := models.BibiteLog{}
	bibiteuserlog.Prepare()
	bibiteuserlog.BibiteUserID = bibiteuserid
	bibiteuserlog.BibiteLogType = bibiteusertype
	bibiteuserlog.BibiteLogAction = "CONFIRM UPDATE"
	bibiteuserlog.BibiteLogDescription = LogDescription
	_, err = bibiteuserlog.SaveBibiteLog(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, addressUpdated.AddressID))

	response.Status = http.StatusCreated
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusCreated, response)
}

func (server *Server) UpdateAddressReject(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseAddressCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	address := models.Address{}
	addresstemp := models.AddressTemp{}

	address.Prepare()
	addresstemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readaddressReceivedTemp, err := addresstemp.FindAddressTempByID(server.DB, pid)
	address.AddressStatus = readaddressReceivedTemp.AddressTempActionBefore
	addressUpdated, err := address.UpdateAddressStatusOnly(server.DB, readaddressReceivedTemp.AddressID)

	_, err = addresstemp.DeleteAddressTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Reject Update Address With AddressTempID = %d", pid)

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

	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, addressUpdated.AddressID))

	response.Status = http.StatusCreated
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusCreated, response)
}

//====================================================================================================================================

func (server *Server) InactiveAddressNew(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseAddressCUD

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

	address := models.Address{}
	addresstemp := models.AddressTemp{}

	err = json.Unmarshal(body, &addresstemp)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	addresstemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readaddressReceived, err := address.FindAddressByID(server.DB, pid)

	address.AddressStatus = 12
	_, err = address.UpdateAddressStatusOnly(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	addresstemp.AddressID = pid
	addresstemp.CountryTempID = readaddressReceived.CountryID
	addresstemp.ProvinceTempID = readaddressReceived.ProvinceID
	addresstemp.RegencyTempID = readaddressReceived.RegencyID
	addresstemp.DistrictTempID = readaddressReceived.DistrictID
	addresstemp.VillageTempID = readaddressReceived.VillageID
	addresstemp.AddressTempStatus = 12
	addresstemp.AddressTempActionBefore = readaddressReceived.AddressStatus
	addresstemp.AddressTempCreatedAt = time.Now()
	addresstempCreated, err := addresstemp.SaveAddressTemp(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Pending Inactive Address From AddressID = %d to AddressTempID = %d", pid, addresstempCreated.AddressTempID)

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

func (server *Server) InactiveAddressConfirm(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseAddressCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	address := models.Address{}
	addresstemp := models.AddressTemp{}

	address.Prepare()
	addresstemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readaddressReceivedTemp, err := addresstemp.FindAddressTempByID(server.DB, pid)
	address.AddressUpdatedBy = readaddressReceivedTemp.AddressTempCreatedBy
	address.AddressStatus = 2
	_, err = address.UpdateAddressStatus(server.DB, readaddressReceivedTemp.AddressID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	_, err = addresstemp.DeleteAddressTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Confirm Inactive Address From Address Temp = %d To With AddressID = %d", pid, readaddressReceivedTemp.AddressID)

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

func (server *Server) InactiveAddressReject(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseAddressCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	address := models.Address{}
	addresstemp := models.AddressTemp{}

	address.Prepare()
	addresstemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readaddressReceivedTemp, err := addresstemp.FindAddressTempByID(server.DB, pid)
	address.AddressStatus = readaddressReceivedTemp.AddressTempActionBefore
	_, err = address.UpdateAddressStatusOnly(server.DB, readaddressReceivedTemp.AddressID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	_, err = addresstemp.DeleteAddressTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Reject Inactive Address With AddressTempID = %d", pid)

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

func (server *Server) ActiveAddressNew(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseAddressCUD

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

	address := models.Address{}
	addresstemp := models.AddressTemp{}

	err = json.Unmarshal(body, &addresstemp)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	addresstemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readaddressReceived, err := address.FindAddressByID(server.DB, pid)

	address.AddressStatus = 11
	_, err = address.UpdateAddressStatusOnly(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	addresstemp.AddressID = pid
	addresstemp.CountryTempID = readaddressReceived.CountryID
	addresstemp.ProvinceTempID = readaddressReceived.ProvinceID
	addresstemp.RegencyTempID = readaddressReceived.RegencyID
	addresstemp.DistrictTempID = readaddressReceived.DistrictID
	addresstemp.VillageTempID = readaddressReceived.VillageID
	addresstemp.AddressTempStatus = 11
	addresstemp.AddressTempActionBefore = readaddressReceived.AddressStatus
	addresstemp.AddressTempCreatedAt = time.Now()
	addresstempCreated, err := addresstemp.SaveAddressTemp(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Pending Active Address From AddressID = %d to AddressTempID = %d", pid, addresstempCreated.AddressTempID)

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

func (server *Server) ActiveAddressConfirm(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseAddressCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	address := models.Address{}
	addresstemp := models.AddressTemp{}

	address.Prepare()
	addresstemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readaddressReceivedTemp, err := addresstemp.FindAddressTempByID(server.DB, pid)
	address.AddressUpdatedBy = readaddressReceivedTemp.AddressTempCreatedBy
	address.AddressStatus = 1
	_, err = address.UpdateAddressStatus(server.DB, readaddressReceivedTemp.AddressID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	_, err = addresstemp.DeleteAddressTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Confirm Active Address From Address Temp = %d To With AddressID = %d", pid, readaddressReceivedTemp.AddressID)

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

func (server *Server) ActiveAddressReject(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseAddressCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	address := models.Address{}
	addresstemp := models.AddressTemp{}

	address.Prepare()
	addresstemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readaddressReceivedTemp, err := addresstemp.FindAddressTempByID(server.DB, pid)
	address.AddressStatus = readaddressReceivedTemp.AddressTempActionBefore
	_, err = address.UpdateAddressStatusOnly(server.DB, readaddressReceivedTemp.AddressID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	_, err = addresstemp.DeleteAddressTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Reject Active Address With AddressTempID = %d", pid)

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

func (server *Server) DeleteAddressNew(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseAddressCUD

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

	address := models.Address{}
	addresstemp := models.AddressTemp{}

	err = json.Unmarshal(body, &addresstemp)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	addresstemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readaddressReceived, err := address.FindAddressByID(server.DB, pid)

	address.AddressStatus = 14
	_, err = address.UpdateAddressStatusOnly(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	addresstemp.AddressID = pid
	addresstemp.CountryTempID = readaddressReceived.CountryID
	addresstemp.ProvinceTempID = readaddressReceived.ProvinceID
	addresstemp.RegencyTempID = readaddressReceived.RegencyID
	addresstemp.DistrictTempID = readaddressReceived.DistrictID
	addresstemp.VillageTempID = readaddressReceived.VillageID
	addresstemp.AddressTempStatus = 14
	addresstemp.AddressTempActionBefore = readaddressReceived.AddressStatus
	addresstemp.AddressTempCreatedAt = time.Now()
	addresstempCreated, err := addresstemp.SaveAddressTemp(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Pending Delete Address From AddressID = %d to AddressTempID = %d", pid, addresstempCreated.AddressTempID)

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

func (server *Server) DeleteAddressConfirm(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseAddressCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	address := models.Address{}
	addresstemp := models.AddressTemp{}

	address.Prepare()
	addresstemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readaddressReceivedTemp, err := addresstemp.FindAddressTempByID(server.DB, pid)
	address.AddressDeletedBy = readaddressReceivedTemp.AddressTempCreatedBy
	address.AddressStatus = 3
	_, err = address.UpdateAddressStatusDelete(server.DB, readaddressReceivedTemp.AddressID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	_, err = addresstemp.DeleteAddressTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Confirm Delete Address From Address Temp = %d To With AddressID = %d", pid, readaddressReceivedTemp.AddressID)

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

func (server *Server) DeleteAddressReject(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseAddressCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	address := models.Address{}
	addresstemp := models.AddressTemp{}

	address.Prepare()
	addresstemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readaddressReceivedTemp, err := addresstemp.FindAddressTempByID(server.DB, pid)
	address.AddressStatus = readaddressReceivedTemp.AddressTempActionBefore
	_, err = address.UpdateAddressStatusOnly(server.DB, readaddressReceivedTemp.AddressID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	_, err = addresstemp.DeleteAddressTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Reject Delete Address With AddressTempID = %d", pid)

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

func (server *Server) DeleteAddress(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseAddressCUD

	vars := mux.Vars(r)

	address := models.Address{}

	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = address.DeleteAddress(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Permanent Delete Address With AddressTempID = %d", pid)

	bibiteuserlog := models.BibiteLog{}
	bibiteuserlog.Prepare()
	bibiteuserlog.BibiteUserID = bibiteuserid
	bibiteuserlog.BibiteLogType = bibiteusertype
	bibiteuserlog.BibiteLogAction = "PERMANENT DELETE"
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

func (server *Server) DeleteAddressTemp(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseAddressTempCUD

	vars := mux.Vars(r)

	address := models.AddressTemp{}

	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = address.DeleteAddressTemp(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Permanent Delete Address With AddressID = %d", pid)

	bibiteuserlog := models.BibiteLog{}
	bibiteuserlog.Prepare()
	bibiteuserlog.BibiteUserID = bibiteuserid
	bibiteuserlog.BibiteLogType = bibiteusertype
	bibiteuserlog.BibiteLogAction = "PERMANENT DELETE"
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
