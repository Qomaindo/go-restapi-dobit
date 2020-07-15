package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hdyro/go-restapi-bit/api/auth"
	"github.com/hdyro/go-restapi-bit/api/models"
	"github.com/hdyro/go-restapi-bit/api/responses"
	"github.com/hdyro/go-restapi-bit/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) CustomerSetPin(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerUserDel

	// vars := mux.Vars(r)
	// uid, err := strconv.ParseUint(vars["id"], 10, 64)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusBadRequest, err)
	// 	return
	// }
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customeruser := models.CustomerUser{}
	customerlog := models.CustomerLog{}
	err = json.Unmarshal(body, &customeruser)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customeruser.Prepare()
	err = customeruser.Validate("setpin")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customerid, err := auth.ExtractTokenUserIDCustomer(r)

	_, err = customeruser.SetPinCustomerUser(server.DB, customerid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	LogDescription := fmt.Sprintf("Customer User Set PIN")
	customerlog.CustomerID = customerid
	customerlog.CustomerLogAction = "Set PIN"
	customerlog.CustomerLogDescription = LogDescription
	_, err = customerlog.SaveCustomerLog(server.DB)
	response.Status = http.StatusCreated
	response.Message = "SUCCESS"
	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) CustomerVerifyPin(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerPin
	var err error

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customeruser := models.CustomerUser{}
	customerlog := models.CustomerLog{}

	err = json.Unmarshal(body, &customeruser)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customeruser.Prepare()
	err = customeruser.Validate("verifypin")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customerid, err := auth.ExtractTokenUserIDCustomer(r)

	pin := customeruser.CustomerUserPin
	customerpin, err := customeruser.FindCustomerUserByCustomerID(server.DB, customerid)
	if err != nil {
		err = errors.New("PIN Salah")
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = models.VerifyPinMobile(customerpin.CustomerUserPin, pin)
	if err != nil || err == bcrypt.ErrMismatchedHashAndPassword {
		messageError := errors.New("PIN Salah")
		responses.ERROR(w, http.StatusUnprocessableEntity, messageError)
		return
	}

	LogDescription := fmt.Sprintf("Customer User Verify PIN")
	customerlog.CustomerID = customerid
	customerlog.CustomerLogAction = "Verify PIN"
	customerlog.CustomerLogDescription = LogDescription
	_, err = customerlog.SaveCustomerLog(server.DB)

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data.Status = customerpin.CustomerUserStatus
	response.Data.Name = customerpin.CustomerName
	response.Data.Photo = customerpin.CustomerImage

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) CustomerUpdatePin(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerPin

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customeruser := models.CustomerUser{}
	customer := models.Customer{}
	customerlog := models.CustomerLog{}

	err = json.Unmarshal(body, &customeruser)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customeruser.Prepare()
	err = customeruser.Validate("updatepin")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customerid, err := auth.ExtractTokenUserIDCustomer(r)

	pin := customeruser.CustomerUserPinOld
	customerpin, err := customeruser.FindCustomerUserByCustomerID(server.DB, customerid)
	if err != nil {
		messageError := errors.New("PIN Salah")
		responses.ERROR(w, http.StatusUnprocessableEntity, messageError)
		return
	}

	err = models.VerifyPinMobile(customerpin.CustomerUserPin, pin)
	if err != nil || err == bcrypt.ErrMismatchedHashAndPassword {
		messageError := errors.New("PIN Salah")
		responses.ERROR(w, http.StatusUnprocessableEntity, messageError)
		return
	}

	updatedCustomerUser, err := customeruser.UpdatePinCustomerUser(server.DB, customerid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	customerdata, err := customer.FindCustomerByID(server.DB, uint64(updatedCustomerUser.CustomerID))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	LogDescription := fmt.Sprintf("Customer User Update PIN")
	customerlog.CustomerID = customerid
	customerlog.CustomerLogAction = "Update PIN"
	customerlog.CustomerLogDescription = LogDescription
	_, err = customerlog.SaveCustomerLog(server.DB)

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data.Status = customerpin.CustomerUserStatus
	response.Data.Name = customerdata.CustomerName
	response.Data.Photo = customerdata.CustomerImage

	responses.JSON(w, http.StatusOK, response)
}
