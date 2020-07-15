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
	"github.com/pusher/pusher-http-go"
)

// =========================
// Status Notif :
// 0 -> Belum dibaca
// 1 -> Sudah dibaca
// 2 -> Status didelete
// =========================

func (server *Server) GetNotificationMitraUser(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseGetNotification
	notifMasuk := models.Notification{}
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	notifMasuks, err := notifMasuk.FindNotificationMitraUser(server.DB, pid, "mitrauser")
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = notifMasuks
	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetNotification(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseGetNotification
	notifMasuk := models.Notification{}
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	notifMasuks, err := notifMasuk.FindNotificationCustomer(server.DB, pid, "customer")
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = notifMasuks
	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetNotificationCustomer(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseGetNotification
	notifMasuk := models.Notification{}

	customerid, err := auth.ExtractTokenUserIDCustomer(r)

	notifMasuks, err := notifMasuk.FindNotificationCustomerByCustomerID(server.DB, customerid, "customer")
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	customerlog := models.CustomerLog{}
	LogDescription := fmt.Sprintf("Customer Get Notification")
	customerlog.CustomerID = customerid
	customerlog.CustomerLogAction = "Customer Get Notification"
	customerlog.CustomerLogDescription = LogDescription
	_, err = customerlog.SaveCustomerLog(server.DB)

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = notifMasuks
	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetNotificationMitraUserByMitraUserID(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseGetNotification
	notifMasuk := models.Notification{}

	mitrauserid, err := auth.ExtractTokenUserIDMitra(r)

	notifMasuks, err := notifMasuk.FindNotificationMitraUserByMitraUserID(server.DB, mitrauserid, "mitrauser")
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	mitrausertype, err := auth.ExtractTokenUserTypeMitra(r)
	LogDescription := fmt.Sprintf("Get Notification")

	mitrauserlog := models.MitraLog{}
	mitrauserlog.Prepare()
	mitrauserlog.MitraUserID = mitrauserid
	mitrauserlog.MitraLogType = mitrausertype
	mitrauserlog.MitraLogAction = "Get Notification"
	mitrauserlog.MitraLogDescription = LogDescription
	_, err = mitrauserlog.SaveMitraLog(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = notifMasuks
	responses.JSON(w, http.StatusOK, response)
}

func handleConnections(server *Server, MitraUserID uint64, CustomerUserID uint64, NotificationStatus uint64, NotificationTitle string, NotificationReceiver string, NotificationMessage string) string {
	pusherClient := pusher.Client{
		AppID:   "955394",
		Key:     "e27bde0d743c43d81e3b",
		Secret:  "97f3fa7cd052ecc5662d",
		Cluster: "ap1",
		Secure:  true,
	}

	// timeNow := fmt.Sprintf(time.Now().Format("2006-01-08 15:04:05"))
	// var timeWIB string
	// timeWIB = fmt.Sprintf(timeNow + "+00")

	notifMasuk := models.Notification{}
	notifMasuk.NotificationTitle = NotificationTitle
	notifMasuk.NotificationReceiver = NotificationReceiver
	notifMasuk.CustomerUserID = CustomerUserID
	notifMasuk.MitraUserID = MitraUserID
	notifMasuk.NotificationReceiver = NotificationReceiver
	notifMasuk.NotificationMessage = NotificationMessage
	// notifMasuk.NotificationStatus = 2
	notifMasuk.NotificationIncoming = 0
	notifMasuk.NotificationStatus = NotificationStatus
	notifMasuk.NotificationCreatedAt = time.Now()
	notifMasuk.SaveNotification(server.DB)
	if notifMasuk.NotificationReceiver == "mitrauser" {
		data := map[string]string{"message": notifMasuk.NotificationMessage}
		pusherClient.Trigger(notifMasuk.NotificationReceiver, fmt.Sprint(notifMasuk.MitraUserID), data)
	} else {
		data := map[string]string{"message": notifMasuk.NotificationMessage}
		fmt.Print(notifMasuk.NotificationReceiver)
		pusherClient.Trigger(notifMasuk.NotificationReceiver, fmt.Sprint(notifMasuk.CustomerUserID), data)
	}
	return "berhasil"
}
func (server *Server) PushNotif(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseGetNotificationDetail

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	// notifMasuk := models.Notification{}
	pusherClient := pusher.Client{
		AppID:   "955394",
		Key:     "e27bde0d743c43d81e3b",
		Secret:  "97f3fa7cd052ecc5662d",
		Cluster: "ap1",
		Secure:  true,
	}

	// timeNow := fmt.Sprintf(time.Now().Format("2006-01-08 15:04:05"))
	// var timeWIB string
	// timeWIB = fmt.Sprintf(timeNow + "+00")

	customeruser := models.CustomerUser{}
	mitrauser := models.MitraUser{}
	notifMasuk := models.Notification{}

	err = json.Unmarshal(body, &notifMasuk)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = json.Unmarshal(body, &customeruser)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = json.Unmarshal(body, &mitrauser)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	mitrauserid, err := auth.ExtractTokenUserIDMitra(r)

	customerid, err := auth.ExtractTokenUserIDCustomer(r)

	customerData, err := customeruser.FindCustomerUserByCustomerID(server.DB, customerid)

	mitrauserData, err := mitrauser.FindMitraUserByID(server.DB, mitrauserid)

	notifMasuk.NotificationTitle = notifMasuk.NotificationTitle
	notifMasuk.NotificationSender = notifMasuk.NotificationSender
	notifMasuk.NotificationReceiver = notifMasuk.NotificationReceiver
	notifMasuk.NotificationLink = notifMasuk.NotificationLink

	if notifMasuk.NotificationReceiver == "allmitrauser" {

		notifMasuk.CustomerUserID = customerData.CustomerUserID

		customerlog := models.CustomerLog{}
		LogDescription := fmt.Sprintf("Customer Create Notification Where CustomerID = %d", customerid)
		customerlog.CustomerID = customerid
		customerlog.CustomerLogAction = "Customer Create Notification"
		customerlog.CustomerLogDescription = LogDescription
		_, err = customerlog.SaveCustomerLog(server.DB)

	} else if notifMasuk.NotificationReceiver == "allcustomer" {

		notifMasuk.MitraUserID = mitrauserData.MitraUserID

		mitrauserid, err := auth.ExtractTokenUserIDMitra(r)
		mitrausertype, err := auth.ExtractTokenUserTypeMitra(r)
		LogDescription := fmt.Sprintf("Create Notification Where MitraUserID = %d", mitrauserid)

		mitrauserlog := models.MitraLog{}
		mitrauserlog.Prepare()
		mitrauserlog.MitraUserID = mitrauserid
		mitrauserlog.MitraLogType = mitrausertype
		mitrauserlog.MitraLogAction = "Create Notification"
		mitrauserlog.MitraLogDescription = LogDescription
		_, err = mitrauserlog.SaveMitraLog(server.DB)
		if err != nil {
			formattedError := formaterror.FormatError(err.Error())
			responses.ERROR(w, http.StatusInternalServerError, formattedError)
			return
		}

	} else if notifMasuk.NotificationReceiver == "customer" {
		customerID, err := customeruser.CheckCustomerUserByPhone(server.DB, customeruser.CustomerUserPhone)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
		}
		notifMasuk.CustomerUserID = customerID.CustomerUserID
		notifMasuk.MitraUserID = mitrauserData.MitraUserID

		mitrauserid, err := auth.ExtractTokenUserIDMitra(r)
		mitrausertype, err := auth.ExtractTokenUserTypeMitra(r)
		LogDescription := fmt.Sprintf("Create Notification Where MitraUserID = %d", mitrauserid)

		mitrauserlog := models.MitraLog{}
		mitrauserlog.Prepare()
		mitrauserlog.MitraUserID = mitrauserid
		mitrauserlog.MitraLogType = mitrausertype
		mitrauserlog.MitraLogAction = "Create Notification"
		mitrauserlog.MitraLogDescription = LogDescription
		_, err = mitrauserlog.SaveMitraLog(server.DB)
		if err != nil {
			formattedError := formaterror.FormatError(err.Error())
			responses.ERROR(w, http.StatusInternalServerError, formattedError)
			return
		}

	} else {
		mitrauserID, err := mitrauser.CheckMitraUserByPhone(server.DB, mitrauser.MitraUserPhone)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
		}
		notifMasuk.CustomerUserID = customerData.CustomerUserID
		notifMasuk.MitraUserID = mitrauserID.MitraUserID
	}

	notifMasuk.NotificationReceiver = notifMasuk.NotificationReceiver
	notifMasuk.NotificationMessage = notifMasuk.NotificationMessage
	// notifMasuk.NotificationStatus = 2
	notifMasuk.NotificationIncoming = 0
	notifMasuk.NotificationStatus = notifMasuk.NotificationStatus
	notifMasuk.NotificationCreatedAt = time.Now()
	saveNotif, err := notifMasuk.SaveNotification(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	if notifMasuk.NotificationReceiver == "mitrauser" {
		data := map[string]string{"message": notifMasuk.NotificationMessage}
		pusherClient.Trigger(notifMasuk.NotificationReceiver, fmt.Sprint(notifMasuk.MitraUserID), data)
	} else if notifMasuk.NotificationReceiver == "customer" {
		data := map[string]string{"message": notifMasuk.NotificationMessage}
		fmt.Print(notifMasuk.NotificationReceiver)
		pusherClient.Trigger(notifMasuk.NotificationReceiver, fmt.Sprint(notifMasuk.CustomerUserID), data)
	} else {
		data := map[string]string{"message": notifMasuk.NotificationMessage}
		fmt.Print(notifMasuk.NotificationReceiver)
		pusherClient.Trigger(notifMasuk.NotificationReceiver, fmt.Sprint(nil), data)
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = saveNotif

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) ReadNotification(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseGetNotificationIU

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	notifMasuk := models.Notification{}

	notifMasuk.NotificationStatus = 1

	_, err = notifMasuk.UpdateNotificationStatus(server.DB, pid)

	mitrauserid, err := auth.ExtractTokenUserIDMitra(r)
	customerid, err := auth.ExtractTokenUserIDCustomer(r)

	if mitrauserid != 0 {

		mitrausertype, err := auth.ExtractTokenUserTypeMitra(r)
		LogDescription := fmt.Sprintf("Read Notification Where MitraUserID = %d", mitrauserid)

		mitrauserlog := models.MitraLog{}
		mitrauserlog.Prepare()
		mitrauserlog.MitraUserID = mitrauserid
		mitrauserlog.MitraLogType = mitrausertype
		mitrauserlog.MitraLogAction = "Read Notification"
		mitrauserlog.MitraLogDescription = LogDescription
		_, err = mitrauserlog.SaveMitraLog(server.DB)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
	} else if customerid != 0 {

		customerlog := models.CustomerLog{}
		LogDescription := fmt.Sprintf("Read Notification Where CustomerID = %d", customerid)
		customerlog.CustomerID = customerid
		customerlog.CustomerLogAction = "Read Notification"
		customerlog.CustomerLogDescription = LogDescription
		_, err = customerlog.SaveCustomerLog(server.DB)

	} else {
		fmt.Println("Data Null")
		fmt.Println("DATA ID USER MITRA ATAU USER BIBITE KOSONG")

	}

	fmt.Println("DATA CUSTOMER ID :")
	fmt.Println(customerid)
	fmt.Println("DATA MITRA USER ID :")
	fmt.Println(mitrauserid)

	response.Status = http.StatusCreated
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusCreated, response)
}

func (server *Server) UpdateNotificationIncoming(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseGetNotificationIU

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	notifMasuk := models.Notification{}

	_, err = notifMasuk.UpdateNotificationIncoming(server.DB, pid)

	mitrauserid, err := auth.ExtractTokenUserIDMitra(r)
	customerid, err := auth.ExtractTokenUserIDCustomer(r)

	if mitrauserid != 0 {

		mitrausertype, err := auth.ExtractTokenUserTypeMitra(r)
		LogDescription := fmt.Sprintf("Update Incoming Notification Where MitraUserID = %d", mitrauserid)

		mitrauserlog := models.MitraLog{}
		mitrauserlog.Prepare()
		mitrauserlog.MitraUserID = mitrauserid
		mitrauserlog.MitraLogType = mitrausertype
		mitrauserlog.MitraLogAction = "Update Incoming Notification"
		mitrauserlog.MitraLogDescription = LogDescription
		_, err = mitrauserlog.SaveMitraLog(server.DB)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
	} else if customerid != 0 {

		customerlog := models.CustomerLog{}
		LogDescription := fmt.Sprintf("Update Incoming Notification Where CustomerID = %d", customerid)
		customerlog.CustomerID = customerid
		customerlog.CustomerLogAction = "Update Incoming Notification"
		customerlog.CustomerLogDescription = LogDescription
		_, err = customerlog.SaveCustomerLog(server.DB)

	} else {
		fmt.Println("Data Null")
		fmt.Println("DATA ID USER MITRA ATAU USER BIBITE KOSONG")

	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusCreated, response)
}

func (server *Server) DeleteNotification(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseGetNotificationIU

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	notifMasuk := models.Notification{}

	notifMasuk.NotificationStatus = 2

	_, err = notifMasuk.UpdateNotificationStatus(server.DB, pid)

	mitrauserid, err := auth.ExtractTokenUserIDMitra(r)
	customerid, err := auth.ExtractTokenUserIDCustomer(r)

	if mitrauserid != 0 {

		mitrausertype, err := auth.ExtractTokenUserTypeMitra(r)
		LogDescription := fmt.Sprintf("Delete Notification Where MitraUserID = %d", mitrauserid)

		mitrauserlog := models.MitraLog{}
		mitrauserlog.Prepare()
		mitrauserlog.MitraUserID = mitrauserid
		mitrauserlog.MitraLogType = mitrausertype
		mitrauserlog.MitraLogAction = "Delete Notification"
		mitrauserlog.MitraLogDescription = LogDescription
		_, err = mitrauserlog.SaveMitraLog(server.DB)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
	} else if customerid != 0 {

		customerlog := models.CustomerLog{}
		LogDescription := fmt.Sprintf("Delete Notification Where CustomerID = %d", customerid)
		customerlog.CustomerID = customerid
		customerlog.CustomerLogAction = "Delete Notification"
		customerlog.CustomerLogDescription = LogDescription
		_, err = customerlog.SaveCustomerLog(server.DB)

	} else {
		fmt.Println("Data Null")
		fmt.Println("DATA ID USER MITRA ATAU USER BIBITE KOSONG")

	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusCreated, response)
}
