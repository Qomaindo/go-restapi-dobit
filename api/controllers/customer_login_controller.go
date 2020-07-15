package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/hdyro/go-restapi-bit/api/auth"
	"github.com/hdyro/go-restapi-bit/api/models"
	"github.com/hdyro/go-restapi-bit/api/responses"
)

func (server *Server) CustomerLoginCheckPhone(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerPhone

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

	year, month, day := time.Now().Date()
	var sMonth, thn, hr, dtNow string
	sMonth = fmt.Sprintf("%02d", month)
	thn = strconv.Itoa(year)
	hr = strconv.Itoa(day)

	dtNow = fmt.Sprintf(hr + "-" + sMonth + "-" + thn)

	customeruser.Prepare()
	err = customeruser.Validate("phone")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customerphone, err := customeruser.LoginCustomerUserByPhone(server.DB, string(customeruser.CustomerUserPhone))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	custPhone, err := strconv.ParseUint(customerphone.CustomerUserPhone, 10, 64)

	token, err := auth.CreateTokenCustomer(customerphone.CustomerID, custPhone)

	LogDescription := fmt.Sprintf("Customer User Login At = %d", dtNow)
	customerlog.CustomerID = customerphone.CustomerID
	customerlog.CustomerLogAction = "Login Customer"
	customerlog.CustomerLogDescription = LogDescription
	_, err = customerlog.SaveCustomerLog(server.DB)

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data.Status = customerphone.CustomerUserStatus
	response.Data.Token = token
	response.Data.Name = customerphone.CustomerName
	response.Data.Photo = customerphone.CustomerImage

	responses.JSON(w, http.StatusOK, response)
}
