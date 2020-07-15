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

func (server *Server) GetReferals(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseReferals

	referal := models.Referal{}

	referals, err := referal.FindAllReferal(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = referals

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetReferal(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseReferal

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	referal := models.Referal{}

	referalReceived, err := referal.FindReferalByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = referalReceived

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetReferalsStatus(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseReferals

	vars := mux.Vars(r)
	status, err := strconv.ParseUint(vars["status"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	referal := models.Referal{}

	referalReceived, err := referal.FindAllReferalStatus(server.DB, status)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = referalReceived

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetReferalStatus(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseReferal

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	status, err := strconv.ParseUint(vars["status"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	referal := models.Referal{}

	referalReceived, err := referal.FindReferalStatusByID(server.DB, pid, status)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = referalReceived

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetReferalTemps(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseReferalTempsPend

	referal := models.ReferalTemp{}

	referals, err := referal.FindAllReferalTemp(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = referals

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetReferalTemp(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseReferalTemp

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	referal := models.ReferalTemp{}

	referalReceivedTemp, err := referal.FindReferalTempByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = referalReceivedTemp

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetReferalTempsStatus(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseReferalTemps

	vars := mux.Vars(r)
	status, err := strconv.ParseUint(vars["status"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	referal := models.ReferalTemp{}

	referalReceived, err := referal.FindAllReferalTempStatus(server.DB, status)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = referalReceived

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetReferalTempStatus(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseReferalTemp

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	status, err := strconv.ParseUint(vars["status"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	referal := models.ReferalTemp{}

	referalReceived, err := referal.FindReferalTempStatusByID(server.DB, pid, status)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = referalReceived

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetReferalTempsPendingActive(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseReferalTempsPend

	referal := models.ReferalTemp{}

	referalReceivedTemp, err := referal.FindAllReferalTempPendingActive(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = referalReceivedTemp

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetReferalTempPendingActive(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseReferalTempPend

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	referal := models.ReferalTemp{}

	referalReceived, err := referal.FindReferalTempStatusByIDPendingActive(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = referalReceived

	responses.JSON(w, http.StatusOK, response)
}

//====================================================================================================================================

func (server *Server) CreateReferalNew(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseReferalTempCUD
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

	uploadedFile, handler, err := r.FormFile("referal_temp_image")
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

	referaltemppend := models.ReferalTempPend{}
	referaltemppend.Prepare()
	// err = referaltemp.Validate("insertupdate")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	referaltemppend.ReferalTempName = r.FormValue("referal_temp_name")
	referaltemppend.ReferalTempDesc = r.FormValue("referal_temp_desc")
	referaltemppend.ReferalTempProgram = r.FormValue("referal_temp_program")
	referaltemppend.ReferalTempImage = image
	referaltempquota, _ := strconv.ParseUint(r.FormValue("referal_temp_quota"), 10, 0)
	referaltemppend.ReferalTempQuota = referaltempquota
	referaltemppend.ReferalTempDateActiveStart = r.FormValue("referal_temp_date_active_start")
	referaltemppend.ReferalTempDateActiveFinish = r.FormValue("referal_temp_date_active_finish")
	referaltemppend.ReferalTempReason = r.FormValue("referal_temp_reason")
	referaltemppend.ReferalTempStatus = 11
	referaltemppend.ReferalTempCreatedBy = r.FormValue("referal_temp_created_by")
	referaltemppendCreated, err := referaltemppend.SaveReferalTempPend(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Pending Create Referal With ReferalTempID = %d", referaltemppendCreated.ReferalTempID)

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

	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, referaltemppendCreated.ReferalTempID))

	response.Status = http.StatusCreated
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusCreated, response)
}

func (server *Server) CreateReferalConfirm(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseReferalCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	referal := models.Referal{}
	referaltemp := models.ReferalTemp{}

	referal.Prepare()
	// err = referal.Validate("insert")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	referaltemp.Prepare()
	// err = referal.Validate("insert")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readreferalReceivedTemp, err := referaltemp.FindReferalTempByIDPendingActive(server.DB, pid)
	referal.ReferalCode = RandomString(6)
	referal.ReferalName = readreferalReceivedTemp.ReferalTempName
	referal.ReferalDesc = readreferalReceivedTemp.ReferalTempDesc
	referal.ReferalProgram = readreferalReceivedTemp.ReferalTempProgram
	referal.ReferalImage = readreferalReceivedTemp.ReferalTempImage
	referal.ReferalQuota = readreferalReceivedTemp.ReferalTempQuota
	referal.ReferalDateActiveStart = readreferalReceivedTemp.ReferalTempDateActiveStart
	referal.ReferalDateActiveFinish = readreferalReceivedTemp.ReferalTempDateActiveFinish
	referal.ReferalCreatedBy = readreferalReceivedTemp.ReferalTempCreatedBy
	referal.ReferalStatus = 1
	referalCreated, err := referal.SaveReferal(server.DB)

	_, err = referaltemp.DeleteReferalTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Confirm Create Referal From ReferalTempID = %d To With ReferalID = %d", pid, referalCreated.ReferalID)

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

	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, referalCreated.ReferalID))

	response.Status = http.StatusCreated
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusCreated, response)
}

func (server *Server) CreateReferalReject(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseReferalTempCUD

	vars := mux.Vars(r)

	referal := models.ReferalTemp{}

	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = referal.DeleteReferalTemp(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Reject Create Referal From ReferalTempID = %d", pid)

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

func (server *Server) UpdateReferalNew(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseReferalCUD
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

	referal := models.Referal{}
	referaltemp := models.ReferalTemp{}
	referaltemp.Prepare()

	alias := r.FormValue("alias")
	value := r.FormValue(("referal_temp_image"))
	if value != "" {
		referaltemp.ReferalTempImage = r.FormValue(("referal_temp_image"))
	} else {

		uploadedFile, handler, err := r.FormFile("referal_temp_image")
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
		referaltemp.ReferalTempImage = image
	}

	readreferalReceived, err := referal.FindReferalByID(server.DB, pid)

	referal.ReferalStatus = 13
	_, err = referal.UpdateReferalStatusOnly(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	referaltemp.ReferalID = pid
	referaltemp.ReferalTempName = r.FormValue("referal_temp_name")
	referaltemp.ReferalTempDesc = r.FormValue("referal_temp_desc")
	referaltemp.ReferalTempProgram = r.FormValue("referal_temp_program")
	referaltempquota, _ := strconv.ParseUint(r.FormValue("referal_temp_quota"), 10, 0)
	referaltemp.ReferalTempQuota = referaltempquota
	referaltemp.ReferalTempDateActiveStart = r.FormValue("referal_temp_date_active_start")
	referaltemp.ReferalTempDateActiveFinish = r.FormValue("referal_temp_date_active_finish")
	referaltemp.ReferalTempReason = r.FormValue("referal_temp_reason")
	referaltemp.ReferalTempStatus = 13
	referaltemp.ReferalTempActionBefore = readreferalReceived.ReferalStatus
	referaltemp.ReferalTempCreatedAt = time.Now()
	referaltemp.ReferalTempCreatedBy = r.FormValue("referal_temp_created_by")
	referaltempCreated, err := referaltemp.SaveReferalTemp(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Pending Update Referal From ReferalID = %d To With ReferalTempID = %d", pid, referaltempCreated.ReferalTempID)

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

func (server *Server) UpdateReferalConfirm(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseReferalCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	referal := models.Referal{}
	referaltemp := models.ReferalTemp{}

	referaltemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readreferalReceivedTemp, err := referaltemp.FindReferalTempByID(server.DB, pid)
	referal.ReferalName = readreferalReceivedTemp.ReferalTempName
	referal.ReferalDesc = readreferalReceivedTemp.ReferalTempDesc
	referal.ReferalProgram = readreferalReceivedTemp.ReferalTempProgram
	referal.ReferalImage = readreferalReceivedTemp.ReferalTempImage
	referal.ReferalQuota = readreferalReceivedTemp.ReferalTempQuota
	referal.ReferalDateActiveStart = readreferalReceivedTemp.ReferalTempDateActiveStart
	referal.ReferalDateActiveFinish = readreferalReceivedTemp.ReferalTempDateActiveFinish
	referal.ReferalUpdatedBy = readreferalReceivedTemp.ReferalTempCreatedBy
	referal.ReferalStatus = 1
	referalUpdated, err := referal.UpdateReferal(server.DB, readreferalReceivedTemp.ReferalID)

	_, err = referaltemp.DeleteReferalTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Confirm Update Referal From ReferalTempID = %d To With ReferalID = %d", pid, readreferalReceivedTemp.ReferalID)

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

	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, referalUpdated.ReferalID))

	response.Status = http.StatusCreated
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusCreated, response)
}

func (server *Server) UpdateReferalReject(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseReferalCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	referal := models.Referal{}
	referaltemp := models.ReferalTemp{}

	referal.Prepare()
	referaltemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readreferalReceivedTemp, err := referaltemp.FindReferalTempByID(server.DB, pid)
	referal.ReferalStatus = readreferalReceivedTemp.ReferalTempActionBefore
	referalUpdated, err := referal.UpdateReferalStatusOnly(server.DB, readreferalReceivedTemp.ReferalID)

	_, err = referaltemp.DeleteReferalTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Reject Update Referal From ReferalTempID = %d To With ReferalID = %d", pid, readreferalReceivedTemp.ReferalID)

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

	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, referalUpdated.ReferalID))

	response.Status = http.StatusCreated
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusCreated, response)
}

//====================================================================================================================================

func (server *Server) InactiveReferalNew(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseReferalCUD

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

	referal := models.Referal{}
	referaltemp := models.ReferalTemp{}

	err = json.Unmarshal(body, &referaltemp)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	referaltemp.Prepare()

	readreferalReceived, err := referal.FindReferalByID(server.DB, pid)

	referal.ReferalStatus = 12
	_, err = referal.UpdateReferalStatusOnly(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	referaltemp.ReferalID = pid
	referaltemp.ReferalTempName = readreferalReceived.ReferalName
	referaltemp.ReferalTempDesc = readreferalReceived.ReferalDesc
	referaltemp.ReferalTempProgram = readreferalReceived.ReferalProgram
	referaltemp.ReferalTempImage = readreferalReceived.ReferalImage
	referaltemp.ReferalTempQuota = readreferalReceived.ReferalQuota
	referaltemp.ReferalTempDateActiveStart = readreferalReceived.ReferalDateActiveStart
	referaltemp.ReferalTempDateActiveFinish = readreferalReceived.ReferalDateActiveFinish
	referaltemp.ReferalTempStatus = 12
	referaltemp.ReferalTempActionBefore = readreferalReceived.ReferalStatus
	referaltemp.ReferalTempCreatedAt = time.Now()
	referaltempCreated, err := referaltemp.SaveReferalTemp(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Pending Inactive Referal From ReferalTempID = %d To With ReferalID = %d", pid, referaltempCreated.ReferalTempID)

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

func (server *Server) InactiveReferalConfirm(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseReferalCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	referal := models.Referal{}
	referaltemp := models.ReferalTemp{}

	referal.Prepare()
	referaltemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readreferalReceivedTemp, err := referaltemp.FindReferalTempByID(server.DB, pid)
	referal.ReferalUpdatedBy = readreferalReceivedTemp.ReferalTempCreatedBy
	referal.ReferalStatus = 2
	_, err = referal.UpdateReferalStatus(server.DB, readreferalReceivedTemp.ReferalID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	_, err = referaltemp.DeleteReferalTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Confirm Inactive Referal From ReferalTempID = %d To With ReferalID = %d", pid, readreferalReceivedTemp.ReferalID)

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

func (server *Server) InactiveReferalReject(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseReferalCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	referal := models.Referal{}
	referaltemp := models.ReferalTemp{}

	referal.Prepare()
	referaltemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readreferalReceivedTemp, err := referaltemp.FindReferalTempByID(server.DB, pid)
	referal.ReferalStatus = readreferalReceivedTemp.ReferalTempActionBefore
	_, err = referal.UpdateReferalStatusOnly(server.DB, readreferalReceivedTemp.ReferalID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	_, err = referaltemp.DeleteReferalTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Reject Inactive Referal From ReferalTempID = %d To With ReferalID = %d", pid, readreferalReceivedTemp.ReferalID)

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

func (server *Server) ActiveReferalNew(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseReferalCUD

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

	referal := models.Referal{}
	referaltemp := models.ReferalTemp{}

	err = json.Unmarshal(body, &referaltemp)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	referaltemp.Prepare()

	readreferalReceived, err := referal.FindReferalByID(server.DB, pid)

	referal.ReferalStatus = 11
	_, err = referal.UpdateReferalStatusOnly(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	referaltemp.ReferalID = pid
	referaltemp.ReferalTempName = readreferalReceived.ReferalName
	referaltemp.ReferalTempDesc = readreferalReceived.ReferalDesc
	referaltemp.ReferalTempProgram = readreferalReceived.ReferalProgram
	referaltemp.ReferalTempImage = readreferalReceived.ReferalImage
	referaltemp.ReferalTempQuota = readreferalReceived.ReferalQuota
	referaltemp.ReferalTempDateActiveStart = readreferalReceived.ReferalDateActiveStart
	referaltemp.ReferalTempDateActiveFinish = readreferalReceived.ReferalDateActiveFinish
	referaltemp.ReferalTempStatus = 11
	referaltemp.ReferalTempActionBefore = readreferalReceived.ReferalStatus
	referaltemp.ReferalTempCreatedAt = time.Now()
	referaltempCreated, err := referaltemp.SaveReferalTemp(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Pending Active Referal From ReferalTempID = %d To With ReferalID = %d", pid, referaltempCreated.ReferalTempID)

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

func (server *Server) ActiveReferalConfirm(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseReferalCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	referal := models.Referal{}
	referaltemp := models.ReferalTemp{}

	referal.Prepare()
	referaltemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readreferalReceivedTemp, err := referaltemp.FindReferalTempByID(server.DB, pid)
	referal.ReferalUpdatedBy = readreferalReceivedTemp.ReferalTempCreatedBy
	referal.ReferalStatus = 1
	_, err = referal.UpdateReferalStatus(server.DB, readreferalReceivedTemp.ReferalID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	_, err = referaltemp.DeleteReferalTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Confirm Active Referal From ReferalTempID = %d To With ReferalID = %d", pid, readreferalReceivedTemp.ReferalID)

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

func (server *Server) ActiveReferalReject(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseReferalCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	referal := models.Referal{}
	referaltemp := models.ReferalTemp{}

	referal.Prepare()
	referaltemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readreferalReceivedTemp, err := referaltemp.FindReferalTempByID(server.DB, pid)
	referal.ReferalStatus = readreferalReceivedTemp.ReferalTempActionBefore
	_, err = referal.UpdateReferalStatusOnly(server.DB, readreferalReceivedTemp.ReferalID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	_, err = referaltemp.DeleteReferalTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Reject Active Referal From ReferalTempID = %d To With ReferalID = %d", pid, readreferalReceivedTemp.ReferalID)

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

func (server *Server) DeleteReferalNew(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseReferalCUD

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

	referal := models.Referal{}
	referaltemp := models.ReferalTemp{}

	err = json.Unmarshal(body, &referaltemp)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	referaltemp.Prepare()

	readreferalReceived, err := referal.FindReferalByID(server.DB, pid)

	referal.ReferalStatus = 14
	_, err = referal.UpdateReferalStatusOnly(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	referaltemp.ReferalID = pid
	referaltemp.ReferalTempName = readreferalReceived.ReferalName
	referaltemp.ReferalTempDesc = readreferalReceived.ReferalDesc
	referaltemp.ReferalTempProgram = readreferalReceived.ReferalProgram
	referaltemp.ReferalTempImage = readreferalReceived.ReferalImage
	referaltemp.ReferalTempQuota = readreferalReceived.ReferalQuota
	referaltemp.ReferalTempDateActiveStart = readreferalReceived.ReferalDateActiveStart
	referaltemp.ReferalTempDateActiveFinish = readreferalReceived.ReferalDateActiveFinish
	referaltemp.ReferalTempStatus = 14
	referaltemp.ReferalTempActionBefore = readreferalReceived.ReferalStatus
	referaltemp.ReferalTempCreatedAt = time.Now()
	referaltempCreated, err := referaltemp.SaveReferalTemp(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Pending Delete Referal From ReferalTempID = %d To With ReferalID = %d", pid, referaltempCreated.ReferalTempID)

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

func (server *Server) DeleteReferalConfirm(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseReferalCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	referal := models.Referal{}
	referaltemp := models.ReferalTemp{}

	referal.Prepare()
	referaltemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readreferalReceivedTemp, err := referaltemp.FindReferalTempByID(server.DB, pid)
	referal.ReferalDeletedBy = readreferalReceivedTemp.ReferalTempCreatedBy
	referal.ReferalStatus = 3
	_, err = referal.UpdateReferalStatusDelete(server.DB, readreferalReceivedTemp.ReferalID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	_, err = referaltemp.DeleteReferalTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Confirm Delete Referal From ReferalTempID = %d To With ReferalID = %d", pid, readreferalReceivedTemp.ReferalID)

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

func (server *Server) DeleteReferalReject(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseReferalCUD

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	referal := models.Referal{}
	referaltemp := models.ReferalTemp{}

	referal.Prepare()
	referaltemp.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	readreferalReceivedTemp, err := referaltemp.FindReferalTempByID(server.DB, pid)
	referal.ReferalStatus = readreferalReceivedTemp.ReferalTempActionBefore
	_, err = referal.UpdateReferalStatusOnly(server.DB, readreferalReceivedTemp.ReferalID)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	_, err = referaltemp.DeleteReferalTemp(server.DB, pid)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Reject Delete Referal From ReferalTempID = %d To With ReferalID = %d", pid, readreferalReceivedTemp.ReferalID)

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

func (server *Server) DeleteReferal(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseReferalCUD

	vars := mux.Vars(r)

	referal := models.Referal{}

	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = referal.DeleteReferal(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Delete Referal From ReferalID = %d", pid)

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

func (server *Server) DeleteReferalTemp(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseReferalTempCUD

	vars := mux.Vars(r)

	referal := models.ReferalTemp{}

	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = referal.DeleteReferalTemp(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	bibiteuserid, err := auth.ExtractTokenUserIDBibite(r)
	bibiteusertype, err := auth.ExtractTokenUserTypeBibite(r)
	LogDescription := fmt.Sprintf("Delete Referal Temp From ReferalTempID = %d", pid)

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

func (server *Server) GetReferalCheck(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseReferalCUD

	vars := mux.Vars(r)
	referalcode := vars["referalcode"]
	lengthreferalcode := len(referalcode)

	referal := models.Referal{}
	mitra := models.Mitra{}
	company := models.Company{}
	mitracompany := models.MitraCompany{}
	mitrauser := models.MitraUser{}
	customeruser := models.CustomerUser{}

	if lengthreferalcode == 6 {
		_, err := referal.FindAllReferalCode(server.DB, referalcode)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
	} else if lengthreferalcode == 8 {
		_, err := mitra.FindAllMitraReferalCode(server.DB, referalcode)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
	} else if lengthreferalcode == 10 {
		_, err := customeruser.FindAllCustomerReferalCode(server.DB, referalcode)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
	} else if lengthreferalcode == 12 {
		_, err := mitrauser.FindAllMitraUserReferalCode(server.DB, referalcode)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
	} else if lengthreferalcode == 13 {
		_, err := company.FindAllCompanyReferalCode(server.DB, referalcode)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
	} else if lengthreferalcode == 15 {
		_, err := mitracompany.FindAllMitraCompanyReferalCode(server.DB, referalcode)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusOK, response)
}
