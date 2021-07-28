package city

import (
	"api-lokasi-indonesia/data"
	"api-lokasi-indonesia/models"
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
)

type byID struct {
}

func (b *byID) FromProvince(value *string, ginctx *gin.Context) {

	dec, err := data.UnmarshallCity()

	var cities []models.City

	for {
		var city models.City
		err = dec.Decode(&city)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err.Error())
		}

		if city.ProvinceID == *value {
			cities = append(cities, city)
		}
	}

	if len(cities) == 0 {
		ginctx.Writer.Write([]byte("not found"))
		return
	}

	respond, _ := json.Marshal(cities)

	ginctx.Writer.Write(respond)
}
