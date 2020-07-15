package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hdyro/go-restapi-bit/api/auth"
	"github.com/hdyro/go-restapi-bit/api/models"
	"github.com/hdyro/go-restapi-bit/api/responses"
	"github.com/hdyro/go-restapi-bit/api/utils/formaterror"
)

func (server *Server) CustomerRegisterSaveData(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerRegister

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customer := models.Customer{}
	customeruser := models.CustomerUser{}
	customeraddress := models.CustomerAddress{}
	address := models.Address{}
	customerlog := models.CustomerLog{}

	err = json.Unmarshal(body, &customer)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customer.Prepare()
	err = customer.Validate("register")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = json.Unmarshal(body, &customeruser)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customeruser.Prepare()
	err = customeruser.Validate("register")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = json.Unmarshal(body, &customeraddress)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customeraddress.Prepare()
	// err = customeraddress.Validate("register")
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }
	err = json.Unmarshal(body, &address)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	address.Prepare()
	// err = address.Validate("register")
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }

	match, _ := regexp.MatchString("([a-z]+)", customeruser.CustomerUserPhone)
	if match {
		fmt.Println("INI BUKAN ANGKA !!!! ")
		response.Status = http.StatusUnprocessableEntity
		response.Message = "ERROR, Isi nomor telfon dengan angka !"
	} else {

		// Find Customer User
		checkPhone, _ := customeruser.CheckPhoneCustomerUser(server.DB, customeruser.CustomerUserPhone)

		checkEmail, _ := customeruser.CheckEmailCustomerUser(server.DB, customeruser.CustomerUserEmail)

		fmt.Println("Count Data By Phone :")
		fmt.Println(checkPhone.CountPhoneCustomer)

		if checkPhone.CountPhoneCustomer > 0 {
			response.Status = http.StatusUnprocessableEntity
			response.Message = "ERROR, Nomer Handphone Sudah Terdaftar !"
		} else {
			if checkEmail.CountEmailCustomer > 0 {
				response.Status = http.StatusUnprocessableEntity
				response.Message = "ERROR, Email Sudah Terdaftar !"
			} else {
				// Customer
				customer.CustomerStatus = 1
				customerCreated, err := customer.SaveCustomer(server.DB)

				// Customer User
				customeruser.CustomerUserStatus = 1
				customeruser.CustomerID = customerCreated.CustomerID
				customeruser.CustomerUserReferalCode = RandomString(10)

				// err = customeruser.Validation()
				// err := validate.Struct(models.CustomerUser)
				// if err != nil {
				// 	response.Status = http.StatusUnprocessableEntity
				// 	response.Message = "Nomer HP terlalu pendek"
				// }
				customeruserCreated, err := customeruser.SaveCustomerUser(server.DB)

				// Customer Address
				address.CountryID = 0
				address.ProvinceID = 0
				address.RegencyID = 0
				address.DistrictID = 0
				address.VillageID = 0
				address.AddressStatus = 1
				addressCreated, err := address.SaveAddress(server.DB)
				customeraddress.CustomerAddressStatus = 1
				customeraddress.CustomerID = customerCreated.CustomerID
				customeraddress.AddressID = *addressCreated.AddressID
				_, err = customeraddress.SaveCustomerAddress(server.DB)

				LogDescription := fmt.Sprintf("Registration Customer User, Where CustomerID = %d", customeruserCreated.CustomerID)
				customerlog.CustomerID = customeruserCreated.CustomerID
				customerlog.CustomerLogAction = "Regiser Customer"
				customerlog.CustomerLogDescription = LogDescription
				_, err = customerlog.SaveCustomerLog(server.DB)

				if err != nil {
					formattedError := formaterror.FormatError(err.Error())
					responses.ERROR(w, http.StatusInternalServerError, formattedError)
					return
				}
				w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, customeruserCreated.CustomerUserID))

				response.Status = http.StatusOK
				response.Message = "SUCCESS"
			}
		}
	}
	responses.JSON(w, http.StatusCreated, response)
}

func (server *Server) CustomerRegisterSetPin(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerUserIU

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 64)
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
	updatedCustomerUser, err := customeruser.SetPinCustomerUser(server.DB, uint64(uid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	customerid, err := auth.ExtractTokenUserIDCustomer(r)

	LogDescription := fmt.Sprintf("Registration Customer User, Where CustomerID = %d", customerid)
	customerlog.CustomerID = customerid
	customerlog.CustomerLogAction = "Set PIN"
	customerlog.CustomerLogDescription = LogDescription
	_, err = customerlog.SaveCustomerLog(server.DB)

	response.Status = http.StatusCreated
	response.Message = "SUCCESS"
	response.Data = updatedCustomerUser

	responses.JSON(w, http.StatusOK, response)
}
