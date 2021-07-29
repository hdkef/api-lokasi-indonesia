package main

import (
	"api-lokasi-indonesia/forwarder"
	"api-lokasi-indonesia/konstant"
	"api-lokasi-indonesia/province"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

func main() {

	k := konstant.GetKonstantObject()

	r := gin.Default()

	r.GET(k.GetPath(), forwarder.Forward()) //this route path will be forwarded by forwarder
	r.GET(k.GetAllProvincePath(), province.GetAllProvince())

	r.Run()
}
