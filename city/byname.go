package city

import (
	"api-lokasi-indonesia/data"
	"api-lokasi-indonesia/models"
	"api-lokasi-indonesia/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type byName struct {
}

//WARNING!!! value is pointer MUST BE COMPARE WITH DEPOINTER FIRST
func (b *byName) FromProvince(value *string, ginctx *gin.Context) {

	provinceid := data.GetProvinceIDByName(value)

	cities, err := data.UnmarshallCity(func(c interface{}) (interface{}, bool, bool) {
		if c.(models.City).ProvinceID == provinceid {
			return c.(interface{}), true, false
		}
		return c.(interface{}), false, false
	})
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	respond, err := json.Marshal(cities)
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	ginctx.Writer.Write(respond)
}
