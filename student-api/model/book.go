package model

import "net/url"

type Book struct {
	Base
	Name       string `gorm:"type:varchar(150)" json:"name"`
	TotalStock *int   `gorm:"type:int" json:"totalStock"`
}

func (s *Book) Validate() url.Values {
	errs := url.Values{}

	//username is required
	if s.Name == "" {
		errs.Add("name", "name is required")
	}

	//Password is required
	if s.TotalStock == nil {
		errs.Add("totalStock", "TotalStock is required")
	}
	return errs
}