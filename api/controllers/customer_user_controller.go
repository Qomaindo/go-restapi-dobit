package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hdyro/go-restapi-bit/api/auth"
	"github.com/hdyro/go-restapi-bit/api/models"
	"github.com/hdyro/go-restapi-bit/api/responses"
	"github.com/hdyro/go-restapi-bit/api/utils/formaterror"
)

func (server *Server) CreateCustomerUser(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerUserIU

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	customeruser := models.CustomerUser{}
	err = json.Unmarshal(body, &customeruser)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customeruser.Prepare()
	err = customeruser.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customeruser.CustomerUserReferalCode = RandomString(10)
	customeruserCreated, err := customeruser.SaveCustomerUser(server.DB)

	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, customeruserCreated.CustomerUserID))

	response.Status = http.StatusCreated
	response.Message = "SUCCESS"
	response.Data = customeruserCreated

	responses.JSON(w, http.StatusCreated, response)
}

func (server *Server) GetMobileUser(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseMobileUser

	customeruser := models.CustomerUser{}

	customerusers, err := customeruser.FindMobileUser(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = customerusers

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetDetailUserMobile(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseDetailMobileUser

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	customeruser := models.CustomerUser{}
	customeruserGotten, err := customeruser.FindMobileUserByID(server.DB, uint64(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = customeruserGotten

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCustomerUsers(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerUsers

	customeruser := models.CustomerUser{}

	customerusers, err := customeruser.FindAllCustomerUser(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = customerusers

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCustomerUser(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerUser

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	customeruser := models.CustomerUser{}
	customeruserGotten, err := customeruser.FindCustomerUserByID(server.DB, uint64(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = customeruserGotten

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) UpdateCustomerUser(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerUserIU

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customeruser := models.CustomerUser{}
	err = json.Unmarshal(body, &customeruser)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customeruser.Prepare()
	err = customeruser.Validate("update")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedCustomerUser, err := customeruser.UpdateCustomerUser(server.DB, uint64(uid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	response.Status = http.StatusCreated
	response.Message = "SUCCESS"
	response.Data = updatedCustomerUser

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) SuspendMobileUser(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerUserIU

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customeruser := models.CustomerUser{}
	err = json.Unmarshal(body, &customeruser)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customeruser.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customeruser.CustomerUserIsLocked = true
	customeruser.CustomerUserStatus = 4
	_, err = customeruser.SuspendCustomerUser(server.DB, uint64(uid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Suspend Customer User From CustomerUserID = %d", uid)

	bibiteuserlog := models.BibiteLog{}
	bibiteuserlog.Prepare()
	bibiteuserlog.BibiteUserID = bibiteuserid
	bibiteuserlog.BibiteLogType = bibiteusertype
	bibiteuserlog.BibiteLogAction = "SUSPEND USER"
	bibiteuserlog.BibiteLogDescription = LogDescription
	_, err = bibiteuserlog.SaveBibiteLog(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	response.Status = http.StatusCreated
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) DeleteCustomerUser(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerUserDel

	vars := mux.Vars(r)

	customeruser := models.CustomerUser{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = customeruser.DeleteCustomerUser(server.DB, uint64(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))

	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusOK, response)
}
