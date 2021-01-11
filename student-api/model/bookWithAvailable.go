package model

type BookWithAvailable struct {
	Base
	Name       string `json:"name"`
	TotalStock *int   `json:"totalStock"`
	Available  int    `json:"available"`
}
