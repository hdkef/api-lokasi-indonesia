package district

import (
	"api-lokasi-indonesia/data"
	"api-lokasi-indonesia/models"
	"api-lokasi-indonesia/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type byID struct {
}

//WARNING!!! value is pointer MUST BE COMPARE WITH DEPOINTER FIRST
func (b *byID) FromCity(value *string, ginctx *gin.Context) {
	districts, err := data.UnmarshallDistrict(func(c interface{}) (interface{}, bool, bool) {
		if c.(models.District).CityID == *value {
			return c.(interface{}), true, false
		}
		return c.(interface{}), false, false
	})
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	respond, err := json.Marshal(districts)
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	ginctx.Writer.Write(respond)
}
