package models

import (
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

type District struct {
	DistrictID   uint64 `gorm:"column:district_id;primary_key;" json:"district_id"`
	RegencyID    string `gorm:"column:regency_id;size:255;not null;unique" json:"regency_id"`
	DistrictName string `gorm:"column:district_name;size:255;not null;unique" json:"district_name"`
}

type DistrictData struct {
	DistrictName string `json:"disctrict_name"`
}

type ResponseDistricts struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    *[]District `json:"data"`
}

type ResponseDistrict struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    District `json:"data"`
}

func (District) TableName() string {
	return "m_district"
}

func (p *District) Prepare() {
	p.DistrictID = 0
	p.RegencyID = p.RegencyID
	p.DistrictName = html.EscapeString(strings.TrimSpace(p.DistrictName))
}

func (p *District) FindAllDistricts(db *gorm.DB) (*[]District, error) {
	var err error
	districts := []District{}
	err = db.Debug().Model(&District{}).Limit(100).Find(&districts).Error
	if err != nil {
		return &[]District{}, err
	}
	return &districts, nil
}

func (p *District) FindDistrictByID(db *gorm.DB, pid uint64) (*District, error) {
	var err error
	err = db.Debug().Model(&District{}).Where("district_id = ?", pid).Take(&p).Error
	if err != nil {
		return &District{}, err
	}
	return p, nil
}

func (p *District) FindDistrictRegencyByID(db *gorm.DB, pid uint64) (*[]District, error) {
	var err error
	districts := []District{}
	err = db.Debug().Model(&District{}).Where("regency_id = ?", pid).Limit(100).Find(&districts).Error
	if err != nil {
		return &[]District{}, err
	}
	return &districts, nil
}
