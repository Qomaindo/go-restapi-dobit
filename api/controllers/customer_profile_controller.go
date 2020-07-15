package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"cloud.google.com/go/storage"
	"github.com/hdyro/go-restapi-bit/api/auth"
	"github.com/hdyro/go-restapi-bit/api/models"
	"github.com/hdyro/go-restapi-bit/api/responses"
	"github.com/hdyro/go-restapi-bit/api/utils/formaterror"
	"github.com/xlzd/gotp"
)

func (server *Server) CustomerDashboard(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerDashboard

	customeruser := models.CustomerUser{}
	customerlog := models.CustomerLog{}

	customeruser.Prepare()

	customerid, err := auth.ExtractTokenUserIDCustomer(r)

	customerphone, err := customeruser.FindCustomerUserByCustomerID(server.DB, customerid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	LogDescription := fmt.Sprintf("Customer User Get Data Dashboard")
	customerlog.CustomerID = customerid
	customerlog.CustomerLogAction = "Dashboard Customer"
	customerlog.CustomerLogDescription = LogDescription
	_, err = customerlog.SaveCustomerLog(server.DB)

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data.Name = customerphone.CustomerName
	response.Data.Phone = customerphone.CustomerUserPhone
	response.Data.Email = customerphone.CustomerUserEmail
	response.Data.RegisterDate = customerphone.CustomerUserCreatedAt
	response.Data.Status = customerphone.CustomerUserStatus

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) CustomerProfile(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerProfile
	var AddressStatus string

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customeruser := models.CustomerUser{}
	customeraddress := models.CustomerAddress{}
	customerlog := models.CustomerLog{}

	err = json.Unmarshal(body, &customeruser)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customerid, err := auth.ExtractTokenUserIDCustomer(r)

	customeruser.Prepare()
	err = customeruser.Validate("phone-profile")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customerphone, err := customeruser.LoginCustomerUserByPhone(server.DB, string(customeruser.CustomerUserPhone))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	countcustomeraddress, err := customeraddress.FindCustomerAddressByCustomerID(server.DB, uint64(customerphone.CustomerID))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	if countcustomeraddress.VillageID == 0 {
		AddressStatus = "empty"
	} else {
		AddressStatus = "filled"
	}

	LogDescription := fmt.Sprintf("Customer Profile")
	customerlog.CustomerID = customerid
	customerlog.CustomerLogAction = "Customer Profile"
	customerlog.CustomerLogDescription = LogDescription
	_, err = customerlog.SaveCustomerLog(server.DB)

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data.ReferalCode = customerphone.CustomerUserReferalCode
	response.Data.Name = customerphone.CustomerName
	response.Data.Phone = customerphone.CustomerUserPhone
	response.Data.Email = customerphone.CustomerUserEmail
	response.Data.Sex = customerphone.CustomerSex
	response.Data.BirthPlace = customerphone.CustomerBirthPlace
	response.Data.BirthDate = customerphone.CustomerBirthDate
	response.Data.NIK = customerphone.CustomerKTP
	response.Data.NPWP = customerphone.CustomerNPWP
	response.Data.Photo = customerphone.CustomerImage
	response.Data.Passport = customerphone.CustomerPassport
	response.Data.Age = customerphone.CustomerAge
	response.Data.LastEdcuation = customerphone.CustomerLastEducation
	response.Data.MaritalStatus = customerphone.CustomerMaritalStatus
	response.Data.Photo = customerphone.CustomerImage
	response.Data.RegisterDate = customerphone.CustomerUserCreatedAt
	response.Data.Address = AddressStatus
	response.Data.Status = customerphone.CustomerUserStatus

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) CustomerProfileAddress(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCustomerProfileAddress

	customeruser := models.CustomerUser{}
	customeraddress := models.CustomerAddress{}
	customerlog := models.CustomerLog{}

	customerid, err := auth.ExtractTokenUserIDCustomer(r)

	customerphone, err := customeruser.FindCustomerUserByCustomerID(server.DB, customerid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	customeradrs, err := customeraddress.FindCustomerAddressByCustomerID(server.DB, customerid)
	if err != nil {
		errs := `record customer address not found`
		responses.JSON(w, http.StatusInternalServerError, errs)
		return
	}

	LogDescription := fmt.Sprintf("Customer User Get Profile Address")
	customerlog.CustomerID = customerid
	customerlog.CustomerLogAction = "Get Profile Address"
	customerlog.CustomerLogDescription = LogDescription
	_, err = customerlog.SaveCustomerLog(server.DB)

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data.ReferalCodeFrom = customerphone.CustomerUserReferalCodeFrom
	response.Data.ReferalCode = customerphone.CustomerUserReferalCode
	response.Data.Name = customerphone.CustomerName
	response.Data.Phone = customerphone.CustomerUserPhone
	response.Data.Email = customerphone.CustomerUserEmail
	response.Data.Sex = customerphone.CustomerSex
	response.Data.BirthPlace = customerphone.CustomerBirthPlace
	response.Data.BirthDate = customerphone.CustomerBirthDate
	response.Data.NIK = customerphone.CustomerKTP
	response.Data.NPWP = customerphone.CustomerNPWP
	response.Data.Photo = customerphone.CustomerImage
	response.Data.Passport = customerphone.CustomerPassport
	response.Data.Age = customerphone.CustomerAge
	response.Data.LastEdcuation = customerphone.CustomerLastEducation
	response.Data.MaritalStatus = customerphone.CustomerMaritalStatus
	response.Data.Photo = customerphone.CustomerImage
	response.Data.Address = customeradrs.CustomerAddressStreet
	response.Data.CountryID = customeradrs.CountryID
	response.Data.Country = customeradrs.CountryName
	response.Data.ProvinceID = customeradrs.ProvinceID
	response.Data.Province = customeradrs.ProvinceName
	response.Data.RegencyID = customeradrs.RegencyID
	response.Data.Regency = customeradrs.RegencyName
	response.Data.DistrictID = customeradrs.DistrictID
	response.Data.District = customeradrs.DistrictName
	response.Data.VillageID = customeradrs.VillageID
	response.Data.Village = customeradrs.VillageName
	response.Data.ZipCode = customeradrs.CustomerAddressZipCode
	response.Data.RegisterDate = customerphone.CustomerUserCreatedAt
	response.Data.Status = customerphone.CustomerUserStatus

	responses.JSON(w, http.StatusOK, response)
}

// func (server *Server) FindProfileAddressImage(w http.ResponseWriter, r *http.Request) {

// 	var response models.ResponseCustomerAddressImageData
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}

// 	customeruser := models.CustomerUser{}
// 	customeraddress := models.CustomerAddress{}
// 	customeraddressimage := models.CustomerAddressImage{}
// 	err = json.Unmarshal(body, &customeruser)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}

// 	customeruser.Prepare()
// 	err = customeruser.Validate("phone-profile")
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}

// 	customerphone, err := customeruser.LoginCustomerUserByPhone(server.DB, string(customeruser.CustomerUserPhone))
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	customeradrs, err := customeraddress.FindCustomerAddressByCustomerID(server.DB, uint64(customerphone.CustomerID))
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	customeradrsimg, err := customeraddressimage.FindCustomerAddressImageByIDAddress(server.DB, uint64(customeradrs.CustomerAddressID))
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	fmt.Println(customeradrsimg)

// 	response.Status = http.StatusOK
// 	response.Message = "SUCCESS"
// 	response.Data = customeradrsimg

// 	responses.JSON(w, http.StatusOK, response)
// }

// func (server *Server) UploadCustomerProfileAddressImage(w http.ResponseWriter, r *http.Request) {

// 	var response models.ResponseCustomerAddressImage
// 	vars := mux.Vars(r)
// 	uid, err := strconv.ParseUint(vars["id"], 10, 64)
// 	fmt.Println(uid)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusBadRequest, err)
// 		return
// 	}
// 	if err := r.ParseMultipartForm(1024); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	secretLength := 30
// 	random := gotp.RandomSecret(secretLength)
// 	if r.Method != "POST" {
// 		http.Error(w, "", http.StatusBadRequest)
// 		return
// 	}
// 	if err := r.ParseMultipartForm(1024); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	alias := r.FormValue("alias")

// 	uploadedFile, handler, err := r.FormFile("cust_adrs_img_name")

// 	if err := r.ParseMultipartForm(1024); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	defer uploadedFile.Close()

// 	dir, err := os.Getwd()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	filename := handler.Filename + random
// 	if alias != "" {
// 		filename = fmt.Sprintf("%s%s", alias, filepath.Ext(handler.Filename))
// 	}

// 	fileLocation := filepath.Join(dir, "files", filename)
// 	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer targetFile.Close()

// 	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	log.SetFlags(0)

// 	var projectID, bucket, source, name string
// 	var public bool

// 	bucket = mustGetEnv("CLOUD_BUCKET", bucket)
// 	projectID = mustGetEnv("GCLOUD_PROJECT", projectID)
// 	name = filename
// 	source = fileLocation
// 	public = true
// 	var read io.Reader
// 	if source == "" {
// 		read = os.Stdin
// 		log.Printf("Reading from stdin...")
// 	} else {
// 		f, err := os.Open(source)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		defer f.Close()
// 		read = f
// 	}

// 	if name == "" {
// 		if source != "" {
// 			name = filepath.Base(source)
// 		} else {
// 			name = "test-sample"
// 		}
// 	}

// 	ctx := context.Background()
// 	_, objAttrs, err := upload(ctx, read, projectID, bucket, name, public)
// 	if err != nil {
// 		switch err {
// 		case storage.ErrBucketNotExist:
// 			log.Fatal("Please create the bucket first e.g. with `gsutil mb`")
// 		default:
// 			log.Fatal(err)
// 		}
// 	}
// 	image := objectURL(objAttrs)
// 	log.Printf("URL: %s", objectURL(objAttrs))
// 	os.Remove(source)

// 	customeruser := models.CustomerUser{}
// 	customeraddress := models.CustomerAddress{}
// 	customeraddressimage := models.CustomerAddressImage{}

// 	customeruser.Prepare()

// 	struid := fmt.Sprint(uid)
// 	customerphone, err := customeruser.LoginCustomerUserByPhone(server.DB, struid)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	customeradrs, err := customeraddress.FindCustomerAddressByCustomerID(server.DB, uint64(customerphone.CustomerID))
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	customeraddressimage.Prepare()
// 	customeraddressimage.CustomerAddressID = customeradrs.CustomerAddressID
// 	customeraddressimage.CustomerAddressImageName = image
// 	customeraddressimage.CustomerAddressImageDesc = r.FormValue("cust_adrs_img_desc")
// 	customeraddressimageCreated, err := customeraddressimage.SaveCustomerAddressImage(server.DB)
// 	if err != nil {
// 		formattedError := formaterror.FormatError(err.Error())
// 		responses.ERROR(w, http.StatusInternalServerError, formattedError)
// 		return
// 	}
// 	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, customeraddressimageCreated.CustomerAddressImageID))

// 	response.Status = http.StatusOK
// 	response.Message = "SUCCESS"
// 	response.Data = customeraddressimageCreated

// 	responses.JSON(w, http.StatusOK, response)
// }

func (server *Server) uploadProfile(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseProfileImage

	customer := models.Customer{}
	customeruser := models.CustomerUser{}
	customerlog := models.CustomerLog{}
	customer.Prepare()
	customeruser.Prepare()

	secretLength := 30
	random := gotp.RandomSecret(secretLength)
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	alias := r.FormValue("alias")

	uploadedFile, handler, err := r.FormFile("uploadImage")

	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer uploadedFile.Close()

	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	filename := handler.Filename + random
	if alias != "" {
		filename = fmt.Sprintf("%s%s", alias, filepath.Ext(handler.Filename))
	}

	fileLocation := filepath.Join(dir, "files", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.SetFlags(0)

	var projectID, bucket, source, name string
	var public bool

	bucket = mustGetEnv("CLOUD_BUCKET", bucket)
	projectID = mustGetEnv("GCLOUD_PROJECT", projectID)
	name = filename
	source = fileLocation
	public = true
	var read io.Reader
	if source == "" {
		read = os.Stdin
		log.Printf("Reading from stdin...")
	} else {
		f, err := os.Open(source)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		read = f
	}

	if name == "" {
		if source != "" {
			name = filepath.Base(source)
		} else {
			name = "test-sample"
		}
	}

	ctx := context.Background()
	_, objAttrs, err := upload(ctx, read, projectID, bucket, name, public)
	if err != nil {
		switch err {
		case storage.ErrBucketNotExist:
			log.Fatal("Please create the bucket first e.g. with `gsutil mb`")
		default:
			log.Fatal(err)
		}
	}
	image := objectURL(objAttrs)
	log.Printf("URL: %s", objectURL(objAttrs))
	os.Remove(source)

	customerid, err := auth.ExtractTokenUserIDCustomer(r)

	UpdateCustomerImage, err := customer.UpdateCustomerImage(server.DB, customerid, image)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	LogDescription := fmt.Sprintf("Customer User Upload Foto Profile")
	customerlog.CustomerID = customerid
	customerlog.CustomerLogAction = "Upload Foto Profile"
	customerlog.CustomerLogDescription = LogDescription
	_, err = customerlog.SaveCustomerLog(server.DB)

	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, UpdateCustomerImage.CustomerID))

	response.Status = http.StatusCreated
	response.Message = "SUCCESS"
	response.Data = image

	responses.JSON(w, http.StatusCreated, response)
}

func (server *Server) CustomerUpdateProfile(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseUpdateCustomerProfile

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
	address.Prepare()

	customeraddress.Prepare()

	customeruser.Prepare()

	customer.Prepare()

	err = json.Unmarshal(body, &customeruser)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = json.Unmarshal(body, &customer)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	customerid, err := auth.ExtractTokenUserIDCustomer(r)

	_, err = customeruser.UpdateProfileCustomer(server.DB, customerid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	err = json.Unmarshal(body, &customer)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	updatedCustomer, err := customer.UpdateProfileCustomer(server.DB, customerid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	err = json.Unmarshal(body, &customeraddress)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedCustomerAddress, err := customeraddress.UpdateProfileCustomer(server.DB, customerid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	err = json.Unmarshal(body, &address)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	_, err = address.UpdateAddress(server.DB, uint64(updatedCustomerAddress.AddressID))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	LogDescription := fmt.Sprintf("Customer User Update Profile")
	customerlog.CustomerID = customerid
	customerlog.CustomerLogAction = "Update Profile"
	customerlog.CustomerLogDescription = LogDescription
	_, err = customerlog.SaveCustomerLog(server.DB)

	response.Status = http.StatusCreated
	response.Message = "SUCCESS"
	response.Data.UserCustomer = customerid
	response.Data.Customer = updatedCustomer.CustomerID

	responses.JSON(w, http.StatusOK, response)
}
