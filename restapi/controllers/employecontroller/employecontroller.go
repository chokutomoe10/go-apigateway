package employecontroller

import (
	"encoding/json"
	"go-api-mux/helper"
	"go-api-mux/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var ResponseJson = helper.ResponseJson
var ResponseError = helper.ResponseError

func Index(w http.ResponseWriter, r *http.Request) {
	var employes []models.Employe

	if err := models.DB.Find(&employes).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJson(w, http.StatusOK, employes)
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var employe models.Employe
	if err := models.DB.First(&employe, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, "Employe Not Found")
			return
		default:
			ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	ResponseJson(w, http.StatusOK, employe)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var employe models.Employe

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&employe); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if err := models.DB.Create(&employe).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJson(w, http.StatusCreated, employe)
}

func Update(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var employe models.Employe

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&employe); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if models.DB.Where("id = ?", id).Updates(&employe).RowsAffected == 0 {
		ResponseError(w, http.StatusBadRequest, "Cannot update employe")
		return
	}

	employe.Id = id

	ResponseJson(w, http.StatusOK, employe)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	input := map[string]string{"id": ""}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	var employe models.Employe
	if models.DB.Delete(&employe, input["id"]).RowsAffected == 0 {
		ResponseError(w, http.StatusBadRequest, "Cannot delete employe")
		return
	}

	response := map[string]string{"message": "Employe Deleted"}
	ResponseJson(w, http.StatusOK, response)
}
