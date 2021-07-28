package main

import (
	"api-lokasi-indonesia/data"
	"api-lokasi-indonesia/forwarder"
	"api-lokasi-indonesia/konstant"

	"github.com/gin-gonic/gin"
)

func main() {

	k := konstant.GetKonstantObject()

	r := gin.Default()

	r.GET(k.GetPath(), forwarder.Forward())
	r.GET(k.GetAllProvincePath(), data.GetAllProvince())

	r.Run()
}
