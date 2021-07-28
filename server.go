package main

import (
	"api-lokasi-indonesia/data"
	"api-lokasi-indonesia/forwarder"
	"api-lokasi-indonesia/konstant"
	"api-lokasi-indonesia/models"
	"api-lokasi-indonesia/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

var province []models.Province

func init() {
	province = data.GetProvince()
}

func getAllProvince() gin.HandlerFunc {
	return func(ginctx *gin.Context) {
		respond, err := json.Marshal(province)
		if err != nil {
			utils.ResErr(ginctx, http.StatusInternalServerError, err)
			return
		}

		ginctx.Writer.Write(respond)
	}
}

func main() {

	k := konstant.GetKonstantObject()

	r := gin.Default()

	r.GET(k.GetPath(), forwarder.Forward())
	r.GET(k.GetAllProvincePath(), getAllProvince())

	r.Run()
}
