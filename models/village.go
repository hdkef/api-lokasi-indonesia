package models

type Village struct {
	ID         string `json:"id" csv:"villageid"`
	DistrictID string `json:"districtid" csv:"districtid"`
	Name       string `json:"name" csv:"villagename"`
}
