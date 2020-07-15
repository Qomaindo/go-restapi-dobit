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
	"strconv"
	"time"

	"cloud.google.com/go/storage"
	"github.com/gorilla/mux"
	"github.com/hdyro/go-restapi-bit/api/auth"
	"github.com/hdyro/go-restapi-bit/api/models"
	"github.com/hdyro/go-restapi-bit/api/responses"
	"github.com/hdyro/go-restapi-bit/api/utils/formaterror"
	"github.com/xlzd/gotp"
)

//BASE CRUD
//====================================================================================================================================

func (server *Server) GetCompanys(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCompanys

	company := models.Company{}

	companys, err := company.FindAllCompany(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = companys

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCompany(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCompany

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	company := models.Company{}

	companyReceived, err := company.FindCompanyByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = companyReceived

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCompanysStatus(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCompanys

	vars := mux.Vars(r)
	status, err := strconv.ParseUint(vars["status"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	company := models.Company{}

	companyReceived, err := company.FindAllCompanyStatus(server.DB, status)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = companyReceived

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCompanyStatus(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCompany

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	status, err := strconv.ParseUint(vars["status"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	company := models.Company{}

	companyReceived, err := company.FindCompanyStatusByID(server.DB, pid, status)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = companyReceived

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCompanyTemps(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCompanyTempsPend

	company := models.CompanyTemp{}

	companys, err := company.FindAllCompanyTemp(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = companys

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCompanyTemp(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCompanyTemp

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	company := models.CompanyTemp{}

	companyReceivedTemp, err := company.FindCompanyTempByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = companyReceivedTemp

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCompanyTempsStatus(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCompanyTemps

	vars := mux.Vars(r)
	status, err := strconv.ParseUint(vars["status"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	company := models.CompanyTemp{}

	companyReceived, err := company.FindAllCompanyTempStatus(server.DB, status)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = companyReceived

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCompanyTempStatus(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCompanyTemp

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	status, err := strconv.ParseUint(vars["status"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	company := models.CompanyTemp{}

	companyReceived, err := company.FindCompanyTempStatusByID(server.DB, pid, status)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = companyReceived

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCompanyTempsPendingActive(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCompanyTempsPend

	company := models.CompanyTemp{}

	companyReceivedTemp, err := company.FindAllCompanyTempStatusPendingActive(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = companyReceivedTemp

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetCompanyTempPendingActive(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCompanyTempPend

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	company := models.CompanyTemp{}

	companyReceived, err := company.FindCompanyTempStatusByIDPendingActive(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = companyReceived

	responses.JSON(w, http.StatusOK, response)
}

//====================================================================================================================================

func (server *Server) CreateCompanyNew(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCompanyTempCUD
	var projectID, bucket, source, name string
	var public bool

	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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

	uploadedFile, handler, err := r.FormFile("company_temp_logo")
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
	if err != nil {
		switch err {
		case storage.ErrBucketNotExist:
			log.Fatal("Please create the bucket first e.g. with `gsutil mb`")
		default:
			log.Fatal(err)
		}
	}
	image := objectURL(objAttrs)
	os.Remove(source)

	companytemppend := models.CompanyTempPend{}

	companytemppend.CompanyTempCode = r.FormValue("company_temp_code")
	companytemppend.CompanyTempName = r.FormValue("company_temp_name")
	companytemppend.CompanyTempWebsite = r.FormValue("company_temp_website")
	companytemppend.CompanyTempEmail = r.FormValue("company_temp_email")
	companytemppend.CompanyTempLogo = image
	companytemppend.CompanyTempReason = r.FormValue("company_temp_reason")
	companytemppend.CompanyTempCreatedBy = r.FormValue("company_temp_created_by")
	companytemppend.CompanyTempStatus = 11
	companytemppendCreated, err := companytemppend.SaveCompanyTempPend(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Pending Create Company With CompanyTempID = %d", companytemppendCreated.CompanyTempID)

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

	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, companytemppendCreated.CompanyTempID))

	response.Status = http.StatusCreated
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusCreated, response)
}

func (server *Server) CreateCompanyConfirm(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCompanyCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	company := models.Company{}
	companytemp := models.CompanyTemp{}

	company.Prepare()
	companytemp.Prepare()

	readcompanyReceivedTemp, err := companytemp.FindCompanyTempByIDPendingActive(server.DB, pid)
	company.CompanyCode = readcompanyReceivedTemp.CompanyTempCode
	company.CompanyReferalCode = RandomString(13)
	company.CompanyName = readcompanyReceivedTemp.CompanyTempName
	company.CompanyWebsite = readcompanyReceivedTemp.CompanyTempWebsite
	company.CompanyEmail = readcompanyReceivedTemp.CompanyTempEmail
	company.CompanyLogo = readcompanyReceivedTemp.CompanyTempLogo
	company.CompanyCreatedBy = readcompanyReceivedTemp.CompanyTempCreatedBy
	company.CompanyStatus = 1
	companyCreated, err := company.SaveCompany(server.DB)

	_, err = companytemp.DeleteCompanyTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Confirm Create Company From CompanyTempID = %d To With CompanyID = %d", pid, companyCreated.CompanyID)

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

	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, companyCreated.CompanyID))

	response.Status = http.StatusCreated
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusCreated, response)
}

func (server *Server) CreateCompanyReject(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCompanyTempCUD

	vars := mux.Vars(r)

	company := models.CompanyTemp{}

	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = company.DeleteCompanyTemp(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Reject Create Company From CompanyTempID = %d", pid)

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

func (server *Server) UpdateCompanyNew(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCompanyCUD
	var projectID, bucket, source, name string
	var public bool

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	secretLength := 30
	random := gotp.RandomSecret(secretLength)

	if r.Method != "PUT" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	company := models.Company{}
	companytemp := models.CompanyTemp{}
	companytemp.Prepare()

	alias := r.FormValue("alias")
	value := r.FormValue(("company_temp_logo"))
	if value != "" {
		companytemp.CompanyTempLogo = r.FormValue(("company_temp_logo"))
	} else {

		uploadedFile, handler, err := r.FormFile("company_temp_logo")
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
		if err != nil {
			switch err {
			case storage.ErrBucketNotExist:
				log.Fatal("Please create the bucket first e.g. with `gsutil mb`")
			default:
				log.Fatal(err)
			}
		}
		image := objectURL(objAttrs)
		os.Remove(source)
		companytemp.CompanyTempLogo = image
	}

	readcompanyReceived, err := company.FindCompanyByID(server.DB, pid)

	company.CompanyStatus = 13
	_, err = company.UpdateCompanyStatusOnly(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	companytemp.CompanyID = pid
	companytemp.CompanyTempCode = r.FormValue("company_temp_code")
	companytemp.CompanyTempName = r.FormValue("company_temp_name")
	companytemp.CompanyTempWebsite = r.FormValue("company_temp_website")
	companytemp.CompanyTempEmail = r.FormValue("company_temp_email")
	companytemp.CompanyTempReason = r.FormValue("company_temp_reason")
	companytemp.CompanyTempCreatedBy = r.FormValue("company_temp_created_by")
	companytemp.CompanyTempStatus = 13
	companytemp.CompanyTempActionBefore = readcompanyReceived.CompanyStatus
	companytemp.CompanyTempCreatedAt = time.Now()
	companytempCreated, err := companytemp.SaveCompanyTemp(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Pending Update Company From CompanyID = %d To With CompanyTempID = %d", pid, companytempCreated.CompanyTempID)

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

func (server *Server) UpdateCompanyConfirm(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCompanyCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	company := models.Company{}
	companytemp := models.CompanyTemp{}

	companytemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readcompanyReceivedTemp, err := companytemp.FindCompanyTempByID(server.DB, pid)
	company.CompanyCode = readcompanyReceivedTemp.CompanyTempCode
	company.CompanyName = readcompanyReceivedTemp.CompanyTempName
	company.CompanyWebsite = readcompanyReceivedTemp.CompanyTempWebsite
	company.CompanyEmail = readcompanyReceivedTemp.CompanyTempEmail
	company.CompanyLogo = readcompanyReceivedTemp.CompanyTempLogo
	company.CompanyUpdatedBy = readcompanyReceivedTemp.CompanyTempCreatedBy
	company.CompanyStatus = 1
	companyUpdated, err := company.UpdateCompany(server.DB, readcompanyReceivedTemp.CompanyID)

	_, err = companytemp.DeleteCompanyTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Confirm Update Company From CompanyTempID = %d To With CompanyID = %d", pid, readcompanyReceivedTemp.CompanyID)

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

	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, companyUpdated.CompanyID))

	response.Status = http.StatusCreated
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusCreated, response)
}

func (server *Server) UpdateCompanyReject(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCompanyCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	company := models.Company{}
	companytemp := models.CompanyTemp{}

	company.Prepare()
	companytemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readcompanyReceivedTemp, err := companytemp.FindCompanyTempByID(server.DB, pid)
	company.CompanyStatus = readcompanyReceivedTemp.CompanyTempActionBefore
	companyUpdated, err := company.UpdateCompanyStatusOnly(server.DB, readcompanyReceivedTemp.CompanyID)

	_, err = companytemp.DeleteCompanyTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Reject Update Company From CompanyTempID = %d To With CompanyID = %d", pid, readcompanyReceivedTemp.CompanyID)

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

	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, companyUpdated.CompanyID))

	response.Status = http.StatusCreated
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusCreated, response)
}

//====================================================================================================================================

func (server *Server) InactiveCompanyNew(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCompanyCUD

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

	company := models.Company{}
	companytemp := models.CompanyTemp{}

	err = json.Unmarshal(body, &companytemp)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	companytemp.Prepare()

	readcompanyReceived, err := company.FindCompanyByID(server.DB, pid)

	company.CompanyStatus = 12
	_, err = company.UpdateCompanyStatusOnly(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	companytemp.CompanyID = pid
	companytemp.CompanyTempCode = readcompanyReceived.CompanyCode
	companytemp.CompanyTempName = readcompanyReceived.CompanyName
	companytemp.CompanyTempWebsite = readcompanyReceived.CompanyWebsite
	companytemp.CompanyTempEmail = readcompanyReceived.CompanyEmail
	companytemp.CompanyTempLogo = readcompanyReceived.CompanyLogo
	companytemp.CompanyTempStatus = 12
	companytemp.CompanyTempActionBefore = readcompanyReceived.CompanyStatus
	companytemp.CompanyTempCreatedAt = time.Now()
	companytempCreated, err := companytemp.SaveCompanyTemp(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Pending Inactive Company From CompanyTempID = %d To With CompanyID = %d", pid, companytempCreated.CompanyTempID)

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

func (server *Server) InactiveCompanyConfirm(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCompanyCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	company := models.Company{}
	companytemp := models.CompanyTemp{}

	company.Prepare()
	companytemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readcompanyReceivedTemp, err := companytemp.FindCompanyTempByID(server.DB, pid)
	company.CompanyUpdatedBy = readcompanyReceivedTemp.CompanyTempCreatedBy
	company.CompanyStatus = 2
	_, err = company.UpdateCompanyStatus(server.DB, readcompanyReceivedTemp.CompanyID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	_, err = companytemp.DeleteCompanyTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Confirm Inactive Company From CompanyTempID = %d To With CompanyID = %d", pid, readcompanyReceivedTemp.CompanyID)

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

func (server *Server) InactiveCompanyReject(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCompanyCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	company := models.Company{}
	companytemp := models.CompanyTemp{}

	company.Prepare()
	companytemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readcompanyReceivedTemp, err := companytemp.FindCompanyTempByID(server.DB, pid)
	company.CompanyStatus = readcompanyReceivedTemp.CompanyTempActionBefore
	_, err = company.UpdateCompanyStatusOnly(server.DB, readcompanyReceivedTemp.CompanyID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	_, err = companytemp.DeleteCompanyTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Reject Inactive Company From CompanyTempID = %d To With CompanyID = %d", pid, readcompanyReceivedTemp.CompanyID)

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

func (server *Server) ActiveCompanyNew(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCompanyCUD

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

	company := models.Company{}
	companytemp := models.CompanyTemp{}

	err = json.Unmarshal(body, &companytemp)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	companytemp.Prepare()

	readcompanyReceived, err := company.FindCompanyByID(server.DB, pid)

	company.CompanyStatus = 11
	_, err = company.UpdateCompanyStatusOnly(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	companytemp.CompanyID = pid
	companytemp.CompanyTempCode = readcompanyReceived.CompanyCode
	companytemp.CompanyTempName = readcompanyReceived.CompanyName
	companytemp.CompanyTempWebsite = readcompanyReceived.CompanyWebsite
	companytemp.CompanyTempEmail = readcompanyReceived.CompanyEmail
	companytemp.CompanyTempLogo = readcompanyReceived.CompanyLogo
	companytemp.CompanyTempStatus = 11
	companytemp.CompanyTempActionBefore = readcompanyReceived.CompanyStatus
	companytemp.CompanyTempCreatedAt = time.Now()
	companytempCreated, err := companytemp.SaveCompanyTemp(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Pending Active Company From CompanyTempID = %d To With CompanyID = %d", pid, companytempCreated.CompanyTempID)

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

func (server *Server) ActiveCompanyConfirm(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCompanyCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	company := models.Company{}
	companytemp := models.CompanyTemp{}

	company.Prepare()
	companytemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readcompanyReceivedTemp, err := companytemp.FindCompanyTempByID(server.DB, pid)
	company.CompanyUpdatedBy = readcompanyReceivedTemp.CompanyTempCreatedBy
	company.CompanyStatus = 1
	_, err = company.UpdateCompanyStatus(server.DB, readcompanyReceivedTemp.CompanyID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	_, err = companytemp.DeleteCompanyTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Confirm Active Company From CompanyTempID = %d To With CompanyID = %d", pid, readcompanyReceivedTemp.CompanyID)

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

func (server *Server) ActiveCompanyReject(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCompanyCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	company := models.Company{}
	companytemp := models.CompanyTemp{}

	company.Prepare()
	companytemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readcompanyReceivedTemp, err := companytemp.FindCompanyTempByID(server.DB, pid)
	company.CompanyStatus = readcompanyReceivedTemp.CompanyTempActionBefore
	_, err = company.UpdateCompanyStatusOnly(server.DB, readcompanyReceivedTemp.CompanyID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	_, err = companytemp.DeleteCompanyTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Reject Active Company From CompanyTempID = %d To With CompanyID = %d", pid, readcompanyReceivedTemp.CompanyID)

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

func (server *Server) DeleteCompanyNew(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCompanyCUD

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

	company := models.Company{}
	companytemp := models.CompanyTemp{}

	err = json.Unmarshal(body, &companytemp)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	companytemp.Prepare()

	readcompanyReceived, err := company.FindCompanyByID(server.DB, pid)

	company.CompanyStatus = 14
	_, err = company.UpdateCompanyStatusOnly(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	companytemp.CompanyID = pid
	companytemp.CompanyTempCode = readcompanyReceived.CompanyCode
	companytemp.CompanyTempName = readcompanyReceived.CompanyName
	companytemp.CompanyTempWebsite = readcompanyReceived.CompanyWebsite
	companytemp.CompanyTempEmail = readcompanyReceived.CompanyEmail
	companytemp.CompanyTempLogo = readcompanyReceived.CompanyLogo
	companytemp.CompanyTempStatus = 14
	companytemp.CompanyTempActionBefore = readcompanyReceived.CompanyStatus
	companytemp.CompanyTempCreatedAt = time.Now()
	companytempCreated, err := companytemp.SaveCompanyTemp(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Pending Delete Company From CompanyTempID = %d To With CompanyID = %d", pid, companytempCreated.CompanyTempID)

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

func (server *Server) DeleteCompanyConfirm(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCompanyCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	company := models.Company{}
	companytemp := models.CompanyTemp{}

	company.Prepare()
	companytemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readcompanyReceivedTemp, err := companytemp.FindCompanyTempByID(server.DB, pid)
	company.CompanyDeletedBy = readcompanyReceivedTemp.CompanyTempCreatedBy
	company.CompanyStatus = 3
	_, err = company.UpdateCompanyStatusDelete(server.DB, readcompanyReceivedTemp.CompanyID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	_, err = companytemp.DeleteCompanyTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Confirm Delete Company From CompanyTempID = %d To With CompanyID = %d", pid, readcompanyReceivedTemp.CompanyID)

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

func (server *Server) DeleteCompanyReject(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCompanyCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	company := models.Company{}
	companytemp := models.CompanyTemp{}

	company.Prepare()
	companytemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readcompanyReceivedTemp, err := companytemp.FindCompanyTempByID(server.DB, pid)
	company.CompanyStatus = readcompanyReceivedTemp.CompanyTempActionBefore
	_, err = company.UpdateCompanyStatusOnly(server.DB, readcompanyReceivedTemp.CompanyID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	_, err = companytemp.DeleteCompanyTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Reject Delete Company From CompanyTempID = %d To With CompanyID = %d", pid, readcompanyReceivedTemp.CompanyID)

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

func (server *Server) DeleteCompany(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCompanyCUD

	vars := mux.Vars(r)

	company := models.Company{}

	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = company.DeleteCompany(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Delete Company From CompanyID = %d", pid)

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

func (server *Server) DeleteCompanyTemp(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseCompanyTempCUD

	vars := mux.Vars(r)

	company := models.CompanyTemp{}

	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = company.DeleteCompanyTemp(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Delete Company Temp From CompanyTempID = %d", pid)

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
