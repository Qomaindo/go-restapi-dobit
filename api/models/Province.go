package models

import (
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

type Province struct {
	ProvinceID   uint64 `gorm:"column:province_id;primary_key;" json:"province_id"`
	CountryID    string `gorm:"column:country_id;size:255;not null;unique" json:"country_id"`
	ProvinceName string `gorm:"column:province_name;size:255;not null;unique" json:"province_name"`
}

type ProvinceData struct {
	ProvinceName string `gorm:"json:"province_name"`
}

type ResponseProvinces struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    *[]Province `json:"data"`
}

type ResponseProvince struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    Province `json:"data"`
}

func (Province) TableName() string {
	return "m_province"
}

func (p *Province) Prepare() {
	p.ProvinceID = 0
	p.CountryID = p.CountryID
	p.ProvinceName = html.EscapeString(strings.TrimSpace(p.ProvinceName))
}

func (p *Province) FindAllProvinces(db *gorm.DB) (*[]Province, error) {
	var err error
	countries := []Province{}
	err = db.Debug().Model(&Province{}).Limit(100).Find(&countries).Error
	if err != nil {
		return &[]Province{}, err
	}
	return &countries, nil
}

func (p *Province) FindProvinceByID(db *gorm.DB, pid uint64) (*Province, error) {
	var err error
	err = db.Debug().Model(&Province{}).Where("province_id = ?", pid).Take(&p).Error
	if err != nil {
		return &Province{}, err
	}
	return p, nil
}

func (p *Province) FindProvinceCountryByID(db *gorm.DB, pid uint64) (*[]Province, error) {
	var err error
	countries := []Province{}
	err = db.Debug().Model(&Province{}).Where("country_id = ?", pid).Limit(100).Find(&countries).Error
	if err != nil {
		return &[]Province{}, err
	}
	return &countries, nil
}
