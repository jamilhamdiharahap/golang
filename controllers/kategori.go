package controllers

import (
	"encoding/json"
	"net/http"
	"pemesanan/config"
	"pemesanan/models"

	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = config.Connect()
}

func PostKategoris(w http.ResponseWriter, r *http.Request) {
	var kategoris models.Kategori
	err := json.NewDecoder(r.Body).Decode(&kategoris)
	if err != nil {
		Respon(w, 400, nil, "Bad Request")
		return
	}

	DB.Create(&kategoris)
}
