package village

import (
	"api-lokasi-indonesia/data"
	"api-lokasi-indonesia/models"
	"api-lokasi-indonesia/province"
	"api-lokasi-indonesia/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type byName struct {
}

//WARNING!!! value is pointer MUST BE COMPARE WITH DEPOINTER FIRST
func (b *byName) FromDistrict(value *string, ginctx *gin.Context) {
	districtname := *value

	//get one district from districtname
	district, err := data.UnmarshallDistrict(func(c interface{}) (interface{}, bool, bool) {
		if c.(models.District).Name == districtname {
			return c.(interface{}), true, true //AFTER FOUND 1 RESULT, IMMEDIATELY LOOP MUST BREAK
		}
		return c.(interface{}), false, false
	})
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	//get all village from districtid
	villages, err := data.UnmarshallVillage(func(c interface{}) (interface{}, bool, bool) {
		if c.(models.Village).DistrictID == district[0].ID { //First item of slices
			return c.(interface{}), true, false
		}
		return c.(interface{}), false, false
	})
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	respond, err := json.Marshal(villages)
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	ginctx.Writer.Write(respond)
}

//WARNING!! THIS OPERATION TAKES O^n time complexity
//WARNING!!! value is pointer MUST BE COMPARE WITH DEPOINTER FIRST
func (b *byName) FromCity(value *string, ginctx *gin.Context) {

	var allvillages []models.Village

	cityname := *value

	city, err := data.UnmarshallCity(func(c interface{}) (interface{}, bool, bool) {
		if c.(models.City).Name == cityname {
			return c.(interface{}), true, true //AFTER FOUND 1 RESULT, IMMEDIATELY LOOP MUST BREAK
		}
		return c.(interface{}), false, false
	})
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	//get all districts from city id
	districts, err := data.UnmarshallDistrict(func(c interface{}) (interface{}, bool, bool) {
		if c.(models.District).CityID == city[0].ID { //First item of slices
			return c.(interface{}), true, false
		}
		return c.(interface{}), false, false
	})
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	for _, v := range districts {
		//get all village from districtid
		villages, err := data.UnmarshallVillage(func(c interface{}) (interface{}, bool, bool) {
			if c.(models.Village).DistrictID == v.ID {
				return c.(interface{}), true, false
			}
			return c.(interface{}), false, false
		})
		if err != nil {
			utils.ResErr(ginctx, http.StatusInternalServerError, err)
			return
		}
		allvillages = append(allvillages, villages...)
	}

	respond, err := json.Marshal(allvillages)
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	ginctx.Writer.Write(respond)
}

//WARNING!! THIS OPERATION TAKES O^n*n time complexity
//WARNING!!! value is pointer MUST BE COMPARE WITH DEPOINTER FIRST
//THE USE OF THIS ENDPOINT IS NOT RECOMMENDED
//IT TAKES ABOUT 40 SECOND (CPU AMD A10 RAM 8 GB)
func (b *byName) FromProvince(value *string, ginctx *gin.Context) {

	var allvillages []models.Village
	var alldistricts []models.District

	province := province.GetProvinceByName(value)

	cities, err := data.UnmarshallCity(func(c interface{}) (interface{}, bool, bool) {
		if c.(models.City).ProvinceID == province.ID {
			return c.(interface{}), true, false
		}
		return c.(interface{}), false, false
	})
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	for _, v := range cities {
		districts, err := data.UnmarshallDistrict(func(c interface{}) (interface{}, bool, bool) {
			if c.(models.District).CityID == v.ID {
				return c.(interface{}), true, false
			}
			return c.(interface{}), false, false
		})
		if err != nil {
			utils.ResErr(ginctx, http.StatusInternalServerError, err)
			return
		}
		alldistricts = append(alldistricts, districts...)
	}

	for _, v := range alldistricts {
		villages, err := data.UnmarshallVillage(func(c interface{}) (interface{}, bool, bool) {
			if c.(models.Village).DistrictID == v.ID {
				return c.(interface{}), true, false
			}
			return c.(interface{}), false, false
		})
		if err != nil {
			utils.ResErr(ginctx, http.StatusInternalServerError, err)
			return
		}
		allvillages = append(allvillages, villages...)
	}

	respond, err := json.Marshal(allvillages)
	if err != nil {
		utils.ResErr(ginctx, http.StatusInternalServerError, err)
		return
	}

	ginctx.Writer.Write(respond)

}
