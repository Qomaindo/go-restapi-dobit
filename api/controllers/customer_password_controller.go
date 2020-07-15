package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hdyro/go-restapi-bit/api/auth"
	"github.com/hdyro/go-restapi-bit/api/models"
	"github.com/hdyro/go-restapi-bit/api/responses"
	"github.com/hdyro/go-restapi-bit/api/utils/formaterror"
)

func (server *Server) CustomerVerifyPassword(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerPswd

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
	err = customeruser.Validate("verifypswd")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customerpswd, err := customeruser.LoginCustomerUserByPswd(server.DB, string(customeruser.CustomerUserPhone), string(customeruser.CustomerUserPassword))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data.Status = customerpswd.CustomerUserStatus
	response.Data.Name = customerpswd.CustomerName
	response.Data.Photo = customerpswd.CustomerImage

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) CustomerUpdatePassword(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerPswd

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customeruser := models.CustomerUser{}
	customer := models.Customer{}

	err = json.Unmarshal(body, &customeruser)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customeruser.Prepare()
	err = customeruser.Validate("updatepswd")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customerid, err := auth.ExtractTokenUserIDCustomer(r)

	customerpswd, err := customeruser.LoginCustomerUserByPswd(server.DB, string(customeruser.CustomerUserPhone), string(customeruser.CustomerUserPswdOld))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	fmt.Println(customerpswd)
	updatedCustomerUser, err := customeruser.UpdatePswdCustomerUser(server.DB, customerid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	fmt.Println(updatedCustomerUser)
	customerdata, err := customer.FindCustomerByID(server.DB, uint64(updatedCustomerUser.CustomerID))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data.Status = customerpswd.CustomerUserStatus
	response.Data.Name = customerdata.CustomerName
	response.Data.Photo = customerdata.CustomerImage

	responses.JSON(w, http.StatusOK, response)
}
