package models

type Province struct {
	ID   string `json:"id" csv:"provinceid"`
	Name string `json:"name" csv:"provincename"`
}
