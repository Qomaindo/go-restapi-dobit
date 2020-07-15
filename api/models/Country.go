package models

import (
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

type Country struct {
	CountryID   uint64 `gorm:"column:country_id;primary_key;" json:"country_id"`
	CountryName string `gorm:"column:country_name;size:255;not null;unique" json:"country_name"`
}

type CountryData struct {
	CountryName string `json:"country_name"`
}

type ResponseCountries struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *[]Country `json:"data"`
}

type ResponseCountry struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    Country `json:"data"`
}

func (Country) TableName() string {
	return "m_country"
}

func (p *Country) Prepare() {
	p.CountryID = 0
	p.CountryName = html.EscapeString(strings.TrimSpace(p.CountryName))
}

func (p *Country) FindAllCountries(db *gorm.DB) (*[]Country, error) {
	var err error
	countries := []Country{}
	err = db.Debug().Model(&Country{}).Limit(100).Find(&countries).Error
	if err != nil {
		return &[]Country{}, err
	}
	return &countries, nil
}

func (p *Country) FindCountryByID(db *gorm.DB, pid uint64) (*Country, error) {
	var err error
	err = db.Debug().Model(&Country{}).Where("country_id = ?", pid).Take(&p).Error
	if err != nil {
		return &Country{}, err
	}
	return p, nil
}
