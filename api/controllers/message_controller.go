package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/hdyro/go-restapi-bit/api/models"
	"github.com/hdyro/go-restapi-bit/api/responses"
)

func (server *Server) GetMessages(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseMessages

	vars := mux.Vars(r)
	pid := vars["id"]

	message := models.Message{}

	messageReceived, err := message.FindAllMessage(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = messageReceived

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) GetMessageByID(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseMessage

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	message := models.Message{}

	messageReceived, err := message.FindMessageByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = messageReceived

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseMessageIU

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	message := models.Message{}

	err = json.Unmarshal(body, &message)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// timeNow := fmt.Sprintf(time.Now().Format("2006-01-08 15:04:05"))
	// var timeWIB string
	// timeWIB = fmt.Sprintf(timeNow + "+00")

	// Save Message
	message.MessageStatus = 0
	message.MessageCreatedAt = time.Now()
	saveMessave, err := message.SaveMessage(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	response.Status = http.StatusOK
	response.Message = "SUCCESS"
	response.Data = saveMessave

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) ReadMessage(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseMessageDel

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	message := models.Message{}

	// Save Message
	message.MessageStatus = 1
	_, err = message.UpdateMessageStatus(server.DB, pid)

	err = json.Unmarshal(body, &message)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusOK, response)
}

func (server *Server) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	var response models.ResponseMessageDel

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	message := models.Message{}

	// Save Message
	_, err = message.UpdateMessageStatusDelete(server.DB, pid)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	response.Status = http.StatusOK
	response.Message = "SUCCESS"

	responses.JSON(w, http.StatusOK, response)
}
