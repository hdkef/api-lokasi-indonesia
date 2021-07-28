package models

type Village struct {
	ID          string `json:"id" csv:"villageid"`
	DISTRICT_ID string `json:"districtid" csv:"districtid"`
	Name        string `json:"name" csv:"villagename"`
}
