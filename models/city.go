package models

type City struct {
	ID         string `json:"id" csv:"cityid"`
	ProvinceID string `json:"provinceid" csv:"provinceid"`
	Name       string `json:"name" csv:"cityname"`
}
