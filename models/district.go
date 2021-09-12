package models

type District struct {
	ID     string `json:"id" csv:"districtid"`
	CityID string `json:"cityid" csv:"cityid"`
	Name   string `json:"name" csv:"districtname"`
}
