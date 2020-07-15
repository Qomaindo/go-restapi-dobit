package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/hdyro/go-restapi-bit/api/auth"
	"github.com/hdyro/go-restapi-bit/api/models"
	"github.com/hdyro/go-restapi-bit/api/responses"
	"github.com/hdyro/go-restapi-bit/api/utils/formaterror"
	"github.com/xlzd/gotp"
)

func (server *Server) CustomerSendOTP(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerCodeOTP
	// var otpResponse models.ResponseUserOTP

	secretLength := 16
	totp := gotp.NewTOTP(gotp.RandomSecret(secretLength), 6, 350, nil)

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

	err = customeruser.Validate("sendotp")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customerid, err := auth.ExtractTokenUserIDCustomer(r)

	customeruserReceived, err := customeruser.FindCustomerUserByCustomerID(server.DB, customerid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	customeruser.CustomerUserID = customeruserReceived.CustomerUserID
	customeruser.CustomerUserOTP = totp.Now()

	_, err = customeruser.UpdateCustomerUserOTP(server.DB, customeruserReceived.CustomerUserID)
	// customeruserUpdated, err := customeruser.UpdateCustomerUserOTP(server.DB, customeruserReceived.CustomerUserID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	// apikey := "b2db4ec014343189f197ff1918ae947d"
	// from := "Bibite.ID"
	// to := "62" + customeruserUpdated.CustomerUserPhone
	// msg := "%20Hati-hati%20penipuan!%20Kode%20OTP%20ini%20hanya%20untuk%20kamu,%20jangan%20berikan%20ke%20siapapun.%20Pihak%20Bibite%20tidak%20pernah%20meminta%20kode%20ini" + "%20Kode%20OTP:%20" + totp.Now() + "."
	// apisms := fmt.Sprint("https://api.ayosms.com/mconnect/gw/sendsms.php?api_key=", apikey, "&from=", from, "&to=", to, "&msg=", msg, "%21&trx_id=mytrxid12345&priority=high")
	// resp, err := http.Get(apisms)

	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer resp.Body.Close()

	LogDescription := fmt.Sprintf("Customer User Send OTP Code")
	customerlog.CustomerID = customerid
	customerlog.CustomerLogAction = "Send OTP"
	customerlog.CustomerLogDescription = LogDescription
	_, err = customerlog.SaveCustomerLog(server.DB)

	response.Message = "SUCCESS"
	response.Code = totp.Now()
	response.TimeSend = time.Now()

	responses.JSON(w, http.StatusOK, response)
}
func (server *Server) CustomerVerifyOTP(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerOTPData

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

	err = customeruser.Validate("verifyotp")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customerid, err := auth.ExtractTokenUserIDCustomer(r)

	customeruserReceived, err := customeruser.FindCustomerUserByCustomerID(server.DB, customerid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	customeruser.CustomerUserOTPSecretNew = time.Now()

	// a := true
	// if customeruserReceived.CustomerUserSecret >= 3 {
	// 	_, err = customeruser.UpdateCustomerUserOTPBlank(server.DB, customeruserReceived.CustomerUserID)

	// 	response.Status = http.StatusOK
	// 	response.Condition = "SUSPEND"

	// 	responses.JSON(w, http.StatusOK, response)
	// 	return
	// }

	// if a != (customeruserReceived.CustomerUserOTP == (customeruser.CustomerUserOTPNew)) {
	// 	customeruser.CustomerUserSecret = customeruserReceived.CustomerUserSecret + 1
	// 	_, err = customeruser.UpdateCustomerUserSecret(server.DB, customeruserReceived.CustomerUserID)
	// 	if err != nil {
	// 		formattedError := formaterror.FormatError(err.Error())
	// 		responses.ERROR(w, http.StatusInternalServerError, formattedError)
	// 		return
	// 	}
	// 	response.Status = http.StatusOK
	// 	response.Condition = "WRONG OTP"

	// 	responses.JSON(w, http.StatusOK, response)
	// 	return
	// }
	// if a != (customeruser.CustomerUserOTPSecretNew.Unix() <= (customeruserReceived.CustomerUserOTPSecret.Unix() + 180)) {
	// 	response.Status = http.StatusOK
	// 	response.Condition = "TIME OUT OTP"

	// 	responses.JSON(w, http.StatusOK, response)
	// 	return
	// }

	// _, err = customeruser.UpdateCustomerUserOTPBlank(server.DB, customeruserReceived.CustomerUserID)
	// customeruser.CustomerUserSecret = 0
	// _, err = customeruser.UpdateCustomerUserSecret(server.DB, customeruserReceived.CustomerUserID)

	bibitelogtype := "USR"

	LogDescription := fmt.Sprintf("Customer User Verify OTP")
	customerlog.CustomerID = customerid
	customerlog.CustomerLogAction = "Verify OTP"
	customerlog.CustomerLogDescription = LogDescription
	_, err = customerlog.SaveCustomerLog(server.DB)

	token, err := auth.CreateTokenBibite(customeruserReceived.CustomerUserID, bibitelogtype)
	response.Status = http.StatusOK
	response.Condition = "SUCCESS"
	response.Token = token
	response.Name = customeruserReceived.CustomerName
	response.Phone = customeruserReceived.CustomerUserPhone
	response.Photo = customeruserReceived.CustomerImage

	responses.JSON(w, http.StatusOK, response)
}
