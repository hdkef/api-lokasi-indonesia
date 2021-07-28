package models

type District struct {
	ID         string `json:"id" csv:"districtid"`
	REGENCY_ID string `json:"regencyid" csv:"regencyid"`
	Name       string `json:"name" csv:"districtname"`
}
