package handlers

import (
	"encoding/json"
	"net/http"
	"server/db"
	"server/models"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

//funcion auxiliar
func GetEventByID(r *http.Request) (models.Event, *gorm.DB) {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	event := models.Event{}

	if err := db.Database().First(&event, userId); err.Error != nil {
		return event, err
	} else {
		return event, nil
	}
}

//handlers
func GetEvent(rw http.ResponseWriter, r *http.Request) {
	if event, err := GetEventByID(r); err != nil {
		SendError(rw, http.StatusNotFound)
	} else {
		SendData(rw, event, http.StatusOK)
	}
}

func GetAllEvents(rw http.ResponseWriter, r *http.Request) {
	events := models.AllEvents{}
	db.Database().Find(&events)
	SendData(rw, events, http.StatusOK)
}

func CreateEvent(rw http.ResponseWriter, r *http.Request) {
	event := models.Event{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&event); err != nil {
		SendError(rw, http.StatusUnprocessableEntity)
	} else {
		db.Database().Save(&event)
		SendData(rw, event, http.StatusOK)
	}
}

func UpdateEvent(rw http.ResponseWriter, r *http.Request) {
	if eventAnt, err := GetEventByID(r); err != nil {
		SendError(rw, http.StatusUnprocessableEntity)
	} else {
		event := models.Event{}
		eventId := eventAnt.ID
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&event); err != nil {
			SendError(rw, http.StatusUnprocessableEntity)
		} else {
			event.ID = eventId
			db.Database().Save(&event)
			SendData(rw, event, http.StatusOK)
		}
	}
}

func DeleteEvent(rw http.ResponseWriter, r *http.Request) {
	if event, err := GetEventByID(r); err != nil {
		SendError(rw, http.StatusNotFound)
	} else {
		db.Database().Delete(&event)
		SendData(rw, event, http.StatusOK)
	}
}
