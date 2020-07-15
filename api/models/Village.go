package models

import (
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

type Village struct {
	VillageID   uint64 `gorm:"column:village_id;primary_key;" json:"village_id"`
	DistrictID  string `gorm:"column:district_id;size:255;not null;unique" json:"district_id"`
	VillageName string `gorm:"column:village_name;size:255;not null;unique" json:"village_name"`
}

type VillageData struct {
	VillageName    string `json:"village_name"`
	VillageZipCode string `json:"village_zipcode"`
}

type ResponseVillages struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *[]Village `json:"data"`
}

type ResponseVillage struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    Village `json:"data"`
}

func (Village) TableName() string {
	return "m_village"
}

func (p *Village) Prepare() {
	p.VillageID = 0
	p.DistrictID = p.DistrictID
	p.VillageName = html.EscapeString(strings.TrimSpace(p.VillageName))
}

func (p *Village) FindAllVillages(db *gorm.DB) (*[]Village, error) {
	var err error
	villages := []Village{}
	err = db.Debug().Model(&Village{}).Limit(100).Find(&villages).Error
	if err != nil {
		return &[]Village{}, err
	}
	return &villages, nil
}

func (p *Village) FindVillageByID(db *gorm.DB, pid uint64) (*Village, error) {
	var err error
	err = db.Debug().Model(&Village{}).Where("village_id = ?", pid).Take(&p).Error
	if err != nil {
		return &Village{}, err
	}
	return p, nil
}

func (p *Village) FindVillageDistrictByID(db *gorm.DB, pid uint64) (*[]Village, error) {
	var err error
	villages := []Village{}
	err = db.Debug().Model(&Village{}).Where("district_id = ?", pid).Limit(100).Find(&villages).Error
	if err != nil {
		return &[]Village{}, err
	}
	return &villages, nil
}
