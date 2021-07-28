package city

import (
	"api-lokasi-indonesia/data"
	"api-lokasi-indonesia/models"
	"api-lokasi-indonesia/utils"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type byName struct {
}

func (b *byName) FromProvince(value *string, ginctx *gin.Context) {

	provinceid := data.GetProvinceIDByName(value)

	dec, err := data.UnmarshallCity()

	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	var cities []models.City

	for {
		var city models.City
		err = dec.Decode(&city)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err.Error())
		}

		if city.ProvinceID == provinceid {
			cities = append(cities, city)
		}
	}

	if len(cities) == 0 {
		ginctx.Writer.Write([]byte("not found"))
		return
	}

	respond, err := json.Marshal(cities)
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	ginctx.Writer.Write(respond)
}
