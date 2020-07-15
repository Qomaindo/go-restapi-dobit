package models

import (
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

type Regency struct {
	RegencyID   uint64 `gorm:"column:regency_id;primary_key;" json:"regency_id"`
	ProvinceID  string `gorm:"column:province_id;size:255;not null;unique" json:"province_id"`
	RegencyName string `gorm:"column:regency_name;size:255;not null;unique" json:"regency_name"`
}

type RegencyData struct {
	RegencyName string `json:"regency_name"`
	RegencyBSNI string `json:"regency_bsni"`
}

type ResponseRegencies struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *[]Regency `json:"data"`
}

type ResponseRegency struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    Regency `json:"data"`
}

func (Regency) TableName() string {
	return "m_regency"
}

func (p *Regency) Prepare() {
	p.RegencyID = 0
	p.ProvinceID = p.ProvinceID
	p.RegencyName = html.EscapeString(strings.TrimSpace(p.RegencyName))
}

func (p *Regency) SaveRegency(db *gorm.DB) (*Regency, error) {
	var err error
	err = db.Debug().Model(&Regency{}).Create(&p).Error
	if err != nil {
		return &Regency{}, err
	}
	return p, nil
}

func (p *Regency) FindAllRegencies(db *gorm.DB) (*[]Regency, error) {
	var err error
	countries := []Regency{}
	err = db.Debug().Model(&Regency{}).Limit(100).Find(&countries).Error
	if err != nil {
		return &[]Regency{}, err
	}
	return &countries, nil
}

func (p *Regency) FindRegencyByID(db *gorm.DB, pid uint64) (*Regency, error) {
	var err error
	err = db.Debug().Model(&Regency{}).Where("regency_id = ?", pid).Take(&p).Error
	if err != nil {
		return &Regency{}, err
	}
	return p, nil
}

func (p *Regency) FindRegencyProvinceByID(db *gorm.DB, pid uint64) (*[]Regency, error) {
	var err error
	countries := []Regency{}
	err = db.Debug().Model(&Regency{}).Where("province_id = ?", pid).Find(&countries).Error
	if err != nil {
		return &[]Regency{}, err
	}
	return &countries, nil
}
