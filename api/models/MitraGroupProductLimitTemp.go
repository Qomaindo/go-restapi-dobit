package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

//BASE CRUD
//====================================================================================================================================

type MitraGroupProductLimitTemp struct {
	MitraGroupProductLimitTempID           uint64                                  `gorm:"column:mitra_group_prod_lim_temp_id;primary_key;" json:"mitra_group_prod_lim_temp_id"`
	MitraGroupProductLimitID               uint64                                  `gorm:"column:mitra_group_prod_lim_id" json:"mitra_group_prod_lim_id"`
	MitraTempID                            uint64                                  `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	MitraGroupProductLimitTempCode         string                                  `gorm:"column:mitra_group_prod_lim_temp_code;size:25" json:"mitra_group_prod_lim_temp_code"`
	MitraGroupProductLimitTempName         string                                  `gorm:"column:mitra_group_prod_lim_temp_name;size:255" json:"mitra_group_prod_lim_temp_name"`
	MitraGroupProductLimitTempStatus       uint64                                  `gorm:"column:mitra_group_prod_lim_temp_status;size:2" json:"mitra_group_prod_lim_temp_status"`
	MitraGroupProductLimitTempActionBefore uint64                                  `gorm:"column:mitra_group_prod_lim_temp_action_before;size:2" json:"mitra_group_prod_lim_temp_action_before"`
	MitraGroupProductLimitTempReason       string                                  `gorm:"column:mitra_group_prod_lim_temp_reason" json:"mitra_group_prod_lim_temp_reason"`
	MitraGroupProductLimitTempCreatedBy    string                                  `gorm:"column:mitra_group_prod_lim_temp_created_by;size:125" json:"mitra_group_prod_lim_temp_created_by"`
	MitraGroupProductLimitTempCreatedAt    time.Time                               `gorm:"column:mitra_group_prod_lim_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_group_prod_lim_temp_created_at"`
	MitraGroupProductNominalLimitTemp      []MitraGroupProductNominalLimitTempData `gorm:"-" json:"mitra_group_prod_lim_temp_nominal"`
}

type MitraGroupProductLimitTempPend struct {
	MitraGroupProductLimitTempID           uint64                                  `gorm:"column:mitra_group_prod_lim_temp_id;primary_key;" json:"mitra_group_prod_lim_temp_id"`
	MitraGroupProductLimitID               *uint64                                 `gorm:"column:mitra_group_prod_lim_id" json:"mitra_group_prod_lim_id"`
	MitraTempID                            uint64                                  `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	MitraGroupProductLimitTempCode         string                                  `gorm:"column:mitra_group_prod_lim_temp_code;size:25" json:"mitra_group_prod_lim_temp_code"`
	MitraGroupProductLimitTempName         string                                  `gorm:"column:mitra_group_prod_lim_temp_name;size:255" json:"mitra_group_prod_lim_temp_name"`
	MitraGroupProductLimitTempStatus       uint64                                  `gorm:"column:mitra_group_prod_lim_temp_status;size:2" json:"mitra_group_prod_lim_temp_status"`
	MitraGroupProductLimitTempActionBefore uint64                                  `gorm:"column:mitra_group_prod_lim_temp_action_before;size:2" json:"mitra_group_prod_lim_temp_action_before"`
	MitraGroupProductLimitTempReason       string                                  `gorm:"column:mitra_group_prod_lim_temp_reason" json:"mitra_group_prod_lim_temp_reason"`
	MitraGroupProductLimitTempCreatedBy    string                                  `gorm:"column:mitra_group_prod_lim_temp_created_by;size:125" json:"mitra_group_prod_lim_temp_created_by"`
	MitraGroupProductLimitTempCreatedAt    time.Time                               `gorm:"column:mitra_group_prod_lim_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_group_prod_lim_temp_created_at"`
	MitraGroupProductNominalLimitTemp      []MitraGroupProductNominalLimitTempPend `gorm:"-" json:"mitra_group_prod_lim_temp_nominal"`
}

type MitraGroupProductLimitTempPendData struct {
	MitraGroupProductLimitTempID           uint64                                      `gorm:"column:mitra_group_prod_lim_temp_id;primary_key;" json:"mitra_group_prod_lim_temp_id"`
	MitraGroupProductLimitID               uint64                                      `gorm:"column:mitra_group_prod_lim_id" json:"mitra_group_prod_lim_id"`
	MitraTempID                            uint64                                      `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	MitraGroupProductLimitTempCode         string                                      `gorm:"column:mitra_group_prod_lim_temp_code;size:25" json:"mitra_group_prod_lim_temp_code"`
	MitraGroupProductLimitTempName         string                                      `gorm:"column:mitra_group_prod_lim_temp_name;size:255" json:"mitra_group_prod_lim_temp_name"`
	MitraGroupProductLimitTempStatus       uint64                                      `gorm:"column:mitra_group_prod_lim_temp_status;size:2" json:"mitra_group_prod_lim_temp_status"`
	MitraGroupProductLimitTempActionBefore uint64                                      `gorm:"column:mitra_group_prod_lim_temp_action_before;size:2" json:"mitra_group_prod_lim_temp_action_before"`
	MitraGroupProductLimitTempReason       string                                      `gorm:"column:mitra_group_prod_lim_temp_reason" json:"mitra_group_prod_lim_temp_reason"`
	MitraGroupProductLimitTempCreatedBy    string                                      `gorm:"column:mitra_group_prod_lim_temp_created_by;size:125" json:"mitra_group_prod_lim_temp_created_by"`
	MitraGroupProductLimitTempCreatedAt    time.Time                                   `gorm:"column:mitra_group_prod_lim_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_group_prod_lim_temp_created_at"`
	MitraGroupProductNominalLimitTemp      []MitraGroupProductNominalLimitTempPendData `gorm:"-" json:"mitra_group_prod_lim_temp_nominal"`
}

type MitraGroupProductLimitTempDatas struct {
	MitraGroupProductLimitTempID           uint64    `gorm:"column:mitra_group_prod_lim_temp_id" json:"mitra_group_prod_lim_temp_id"`
	MitraTempID                            uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	MitraGroupProductLimitTempCode         string    `gorm:"column:mitra_group_prod_lim_temp_code" json:"mitra_group_prod_lim_temp_code"`
	MitraGroupProductLimitTempName         string    `gorm:"column:mitra_group_prod_lim_temp_name" json:"mitra_group_prod_lim_temp_name"`
	MitraGroupProductLimitTempReason       string    `gorm:"column:mitra_group_prod_lim_temp_reason" json:"mitra_group_prod_lim_temp_reason"`
	MitraGroupProductLimitTempStatus       uint64    `gorm:"column:mitra_group_prod_lim_temp_status;size:2" json:"mitra_group_prod_lim_temp_status"`
	MitraGroupProductLimitTempActionBefore uint64    `gorm:"column:mitra_group_prod_lim_temp_action_before;size:2" json:"mitra_group_prod_lim_temp_action_before"`
	MitraGroupProductLimitTempCreatedBy    string    `gorm:"column:mitra_group_prod_lim_temp_created_by;size:125" json:"mitra_group_prod_lim_temp_created_by"`
	MitraGroupProductLimitTempCreatedAt    time.Time `gorm:"column:mitra_group_prod_lim_temp_created_at" json:"mitra_group_prod_lim_temp_created_at"`
	MitraGroupProductLimitID               uint64    `gorm:"column:mitra_group_prod_lim_id" json:"mitra_group_prod_lim_id"`
	MitraGroupProductLimitCode             string    `gorm:"column:mitra_group_prod_lim_code" json:"mitra_group_prod_lim_code"`
	MitraGroupProductLimitName             string    `gorm:"column:mitra_group_prod_lim_name" json:"mitra_group_prod_lim_name"`
	MitraGroupProductLimitStatus           uint64    `gorm:"column:mitra_group_prod_lim_status" json:"mitra_group_prod_lim_status"`
	MitraGroupProductLimitCreatedBy        string    `gorm:"column:mitra_group_prod_lim_created_by" json:"mitra_group_prod_lim_created_by"`
	MitraGroupProductLimitCreatedAt        time.Time `gorm:"column:mitra_group_prod_lim_created_at" json:"mitra_group_prod_lim_created_at"`
	MitraGroupProductLimitUpdatedBy        string    `gorm:"column:mitra_group_prod_lim_updated_by" json:"mitra_group_prod_lim_updated_by"`
	MitraGroupProductLimitUpdatedAt        time.Time `gorm:"column:mitra_group_prod_lim_updated_at" json:"mitra_group_prod_lim_updated_at"`
	MitraGroupProductLimitDeletedBy        string    `gorm:"column:mitra_group_prod_lim_deleted_by" json:"mitra_group_prod_lim_deleted_by"`
	MitraGroupProductLimitDeletedAt        time.Time `gorm:"column:mitra_group_prod_lim_deleted_at" json:"mitra_group_prod_lim_deleted_at"`
}

type MitraGroupProductLimitTempData struct {
	MitraGroupProductLimitTempID           uint64                                  `gorm:"column:mitra_group_prod_lim_temp_id" json:"mitra_group_prod_lim_temp_id"`
	MitraTempID                            uint64                                  `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	MitraTempCode                          string                                  `gorm:"column:mitra_temp_code;size:25" json:"mitra_temp_code"`
	MitraTempName                          string                                  `gorm:"column:mitra_temp_name;size:255" json:"mitra_temp_name"`
	MitraTempWebsite                       string                                  `gorm:"column:mitra_temp_website;size:255" json:"mitra_temp_website"`
	MitraTempEmail                         string                                  `gorm:"column:mitra_temp_email;size:255" json:"mitra_temp_email"`
	MitraTempLogo                          string                                  `gorm:"column:mitra_temp_logo;size:255" json:"mitra_temp_logo"`
	MitraGroupProductLimitID               uint64                                  `gorm:"column:mitra_group_prod_lim_id" json:"mitra_group_prod_lim_id"`
	MitraGroupProductLimitCode             string                                  `gorm:"column:mitra_group_prod_lim_code;size:25" json:"mitra_group_prod_lim_code"`
	MitraGroupProductLimitName             string                                  `gorm:"column:mitra_group_prod_lim_name;size:255" json:"mitra_group_prod_lim_name"`
	MitraGroupProductLimitStatus           uint64                                  `gorm:"column:mitra_group_prod_lim_status;size:2" json:"mitra_group_prod_lim_status"`
	MitraGroupProductLimitCreatedAt        time.Time                               `gorm:"column:mitra_group_prod_lim_created_at;default:CURRENT_TIMESTAMP" json:"mitra_group_prod_lim_created_at"`
	MitraGroupProductNominalLimit          []MitraGroupProductNominalLimitData     `json:"mitra_group_prod_lim_privilege"`
	MitraGroupProductLimitTempCode         string                                  `gorm:"column:mitra_group_prod_lim_temp_code;size:25" json:"mitra_group_prod_lim_temp_code"`
	MitraGroupProductLimitTempName         string                                  `gorm:"column:mitra_group_prod_lim_temp_name;size:255" json:"mitra_group_prod_lim_temp_name"`
	MitraGroupProductLimitTempStatus       uint64                                  `gorm:"column:mitra_group_prod_lim_temp_status;size:2" json:"mitra_group_prod_lim_temp_status"`
	MitraGroupProductLimitTempReason       string                                  `gorm:"column:mitra_group_prod_lim_temp_reason" json:"mitra_group_prod_lim_temp_reason"`
	MitraGroupProductLimitTempActionBefore uint64                                  `gorm:"column:mitra_group_prod_lim_temp_action_before;size:2" json:"mitra_group_prod_lim_temp_action_before"`
	MitraGroupProductLimitTempCreatedBy    string                                  `gorm:"column:mitra_group_prod_lim_temp_created_by;size:125" json:"mitra_group_prod_lim_temp_created_by"`
	MitraGroupProductLimitTempCreatedAt    time.Time                               `gorm:"column:mitra_group_prod_lim_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_group_prod_lim_temp_created_at"`
	MitraGroupProductNominalLimitTemp      []MitraGroupProductNominalLimitTempData `json:"mitra_group_prod_lim_temp_nominal"`
}

type ResponseMitraGroupProductLimitTemps struct {
	Status  int                                `json:"status"`
	Message string                             `json:"message"`
	Data    *[]MitraGroupProductLimitTempDatas `json:"data"`
}

type ResponseMitraGroupProductLimitTemp struct {
	Status  int                             `json:"status"`
	Message string                          `json:"message"`
	Data    *MitraGroupProductLimitTempData `json:"data"`
}

type ResponseMitraGroupProductLimitTempsPend struct {
	Status  int                                   `json:"status"`
	Message string                                `json:"message"`
	Data    *[]MitraGroupProductLimitTempPendData `json:"data"`
}

type ResponseMitraGroupProductLimitTempPend struct {
	Status  int                                 `json:"status"`
	Message string                              `json:"message"`
	Data    *MitraGroupProductLimitTempPendData `json:"data"`
}

type ResponseMitraGroupProductLimitTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraGroupProductLimitTemp) TableName() string {
	return "m_mitra_group_product_limit_temp"
}

func (MitraGroupProductLimitTempPend) TableName() string {
	return "m_mitra_group_product_limit_temp"
}

func (MitraGroupProductLimitTempPendData) TableName() string {
	return "m_mitra_group_product_limit_temp"
}

func (MitraGroupProductLimitTempData) TableName() string {
	return "m_mitra_group_product_limit_temp"
}

func (MitraGroupProductLimitTempDatas) TableName() string {
	return "m_mitra_group_product_limit_temp"
}

func (p *MitraGroupProductLimitTemp) Prepare() {
	p.MitraGroupProductLimitID = p.MitraGroupProductLimitID
	p.MitraTempID = p.MitraTempID
	p.MitraGroupProductLimitTempCode = html.EscapeString(strings.TrimSpace(p.MitraGroupProductLimitTempCode))
	p.MitraGroupProductLimitTempName = html.EscapeString(strings.TrimSpace(p.MitraGroupProductLimitTempName))
	p.MitraGroupProductLimitTempReason = html.EscapeString(strings.TrimSpace(p.MitraGroupProductLimitTempReason))
	p.MitraGroupProductLimitTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraGroupProductLimitTempCreatedBy))
	p.MitraGroupProductLimitTempStatus = p.MitraGroupProductLimitTempStatus
	p.MitraGroupProductLimitTempCreatedAt = time.Now()
}

func (p *MitraGroupProductLimitTempPend) Prepare() {
	p.MitraGroupProductLimitID = nil
	p.MitraTempID = p.MitraTempID
	p.MitraGroupProductLimitTempCode = html.EscapeString(strings.TrimSpace(p.MitraGroupProductLimitTempCode))
	p.MitraGroupProductLimitTempName = html.EscapeString(strings.TrimSpace(p.MitraGroupProductLimitTempName))
	p.MitraGroupProductLimitTempReason = html.EscapeString(strings.TrimSpace(p.MitraGroupProductLimitTempReason))
	p.MitraGroupProductLimitTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraGroupProductLimitTempCreatedBy))
	p.MitraGroupProductLimitTempStatus = p.MitraGroupProductLimitTempStatus
	p.MitraGroupProductLimitTempCreatedAt = time.Now()
}

func (p *MitraGroupProductLimitTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.MitraTempID == 0 {
			return errors.New("Required MitraGroupProductLimitTemp Code")
		}
		if p.MitraGroupProductLimitTempCode == "" {
			return errors.New("Required MitraGroupProductLimitTemp Code")
		}
		if p.MitraGroupProductLimitTempName == "" {
			return errors.New("Required MitraGroupProductLimitTemp Name")
		}
		return nil
	}
}

func (p *MitraGroupProductLimitTemp) SaveMitraGroupProductLimitTemp(db *gorm.DB) (*MitraGroupProductLimitTemp, error) {
	var err error
	err = db.Debug().Model(&MitraGroupProductLimitTemp{}).Create(&p).Error
	if err != nil {
		return &MitraGroupProductLimitTemp{}, err
	}
	return p, nil
}

func (p *MitraGroupProductLimitTempPend) SaveMitraGroupProductLimitTempPend(db *gorm.DB) (*MitraGroupProductLimitTempPend, error) {
	var err error
	err = db.Debug().Model(&MitraGroupProductLimitTempPend{}).Create(&p).Error
	if err != nil {
		return &MitraGroupProductLimitTempPend{}, err
	}
	return p, nil
}

func (p *MitraGroupProductLimitTemp) FindAllMitraGroupProductLimitTemp(db *gorm.DB, mitra uint64) (*[]MitraGroupProductLimitTempPendData, error) {
	var err error
	mitragroupproductlimit := []MitraGroupProductLimitTempPendData{}
	err = db.Debug().Model(&MitraGroupProductLimitTempPendData{}).Limit(100).
		Select(`m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_id,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_id,
				m_mitra_group_product_limit_temp.mitra_temp_id,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_code,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_name,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_status,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_action_before,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_reason,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_created_by,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_created_at`).
		Where("mitra_temp_id = ?", mitra).
		Find(&mitragroupproductlimit).Error
	if err != nil {
		return &[]MitraGroupProductLimitTempPendData{}, err
	}
	return &mitragroupproductlimit, nil
}

func (p *MitraGroupProductLimitTemp) FindAllMitraGroupProductLimitTempStatusPendingActive(db *gorm.DB, mitra uint64) (*[]MitraGroupProductLimitTempPendData, error) {
	var err error
	mitragroupproductlimit := []MitraGroupProductLimitTempPendData{}
	err = db.Debug().Model(&MitraGroupProductLimitTempPendData{}).Limit(100).
		Select(`m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_id,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_id,
				m_mitra_group_product_limit_temp.mitra_temp_id,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_code,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_name,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_status,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_action_before,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_reason,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_created_by,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_created_at`).
		Where("mitra_group_prod_lim_temp_status = ? AND mitra_temp_id = ?", 11, mitra).
		Find(&mitragroupproductlimit).Error
	if err != nil {
		return &[]MitraGroupProductLimitTempPendData{}, err
	}
	return &mitragroupproductlimit, nil
}

func (p *MitraGroupProductLimitTemp) FindAllMitraGroupProductLimitTempStatus(db *gorm.DB, status uint64, mitra uint64) (*[]MitraGroupProductLimitTempDatas, error) {
	var err error
	mitragroupproductlimit := []MitraGroupProductLimitTempDatas{}
	err = db.Debug().Model(&MitraGroupProductLimitTempDatas{}).Limit(100).
		Select(`m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_id,
				m_mitra_group_product_limit_temp.mitra_temp_id,
				m_mitra_temp.mitra_code as mitra_temp_code,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_temp.mitra_website as mitra_temp_website,
				m_mitra_temp.mitra_email as mitra_temp_email,
				m_mitra_temp.mitra_logo as mitra_temp_logo,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_id,
				m_mitra_group_product_limit.mitra_group_prod_lim_code,
				m_mitra_group_product_limit.mitra_group_prod_lim_name,
				m_mitra_group_product_limit.mitra_group_prod_lim_status,
				m_mitra_group_product_limit.mitra_group_prod_lim_created_at,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_code,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_name,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_status,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_action_before,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_reason,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_created_by,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_created_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_group_product_limit_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_group_product_limit on m_mitra_group_product_limit_temp.mitra_group_prod_lim_id=m_mitra_group_product_limit.mitra_group_prod_lim_id").
		Where("mitra_group_prod_lim_temp_status = ? AND mitra_temp_id = ?", status, mitra).
		Find(&mitragroupproductlimit).Error
	if err != nil {
		return &[]MitraGroupProductLimitTempDatas{}, err
	}
	return &mitragroupproductlimit, nil
}

func (p *MitraGroupProductLimitTemp) FindMitraGroupProductLimitTempDataByID(db *gorm.DB, pid uint64) (*MitraGroupProductLimitTemp, error) {
	var err error
	err = db.Debug().Model(&MitraGroupProductLimitTemp{}).Where("mitra_group_prod_lim_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraGroupProductLimitTemp{}, err
	}
	return p, nil
}

func (p *MitraGroupProductLimitTemp) FindMitraGroupProductLimitTempByIDPendingActive(db *gorm.DB, pid uint64) (*MitraGroupProductLimitTempPendData, error) {
	var err error
	mitragroupproductlimit := MitraGroupProductLimitTempPendData{}
	err = db.Debug().Model(&MitraGroupProductLimitTempPendData{}).Limit(100).
		Select(`m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_id,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_id,
				m_mitra_group_product_limit_temp.mitra_temp_id,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_code,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_name,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_status,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_action_before,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_reason,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_created_by,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_created_at`).
		Where("mitra_group_prod_lim_temp_id = ?", pid).
		Find(&mitragroupproductlimit).Error
	if err != nil {
		return &MitraGroupProductLimitTempPendData{}, err
	}
	return &mitragroupproductlimit, nil
}

func (p *MitraGroupProductLimitTemp) FindMitraGroupProductLimitTempStatusByIDPendingActive(db *gorm.DB, pid uint64, mitra uint64) (*MitraGroupProductLimitTempPendData, error) {
	var err error
	mitragroupproductlimit := MitraGroupProductLimitTempPendData{}
	err = db.Debug().Model(&MitraGroupProductLimitTempPendData{}).Limit(100).
		Select(`m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_id,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_id,
				m_mitra_group_product_limit_temp.mitra_temp_id,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_code,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_name,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_status,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_action_before,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_reason,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_created_by,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_created_at`).
		Where("mitra_group_prod_lim_temp_id = ? AND mitra_group_prod_lim_temp_status = ? AND mitra_temp_id = ?", pid, 11, mitra).
		Find(&mitragroupproductlimit).Error
	if err != nil {
		return &MitraGroupProductLimitTempPendData{}, err
	}
	return &mitragroupproductlimit, nil
}

func (p *MitraGroupProductLimitTemp) FindMitraGroupProductLimitTempByID(db *gorm.DB, pid uint64) (*MitraGroupProductLimitTempData, error) {
	var err error
	mitragroupproductlimit := MitraGroupProductLimitTempData{}
	err = db.Debug().Model(&MitraGroupProductLimitTempData{}).
		Select(`m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_id,
				m_mitra_group_product_limit_temp.mitra_temp_id,
				m_mitra_temp.mitra_code as mitra_temp_code,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_temp.mitra_website as mitra_temp_website,
				m_mitra_temp.mitra_email as mitra_temp_email,
				m_mitra_temp.mitra_logo as mitra_temp_logo,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_id,
				m_mitra_group_product_limit.mitra_group_prod_lim_code,
				m_mitra_group_product_limit.mitra_group_prod_lim_name,
				m_mitra_group_product_limit.mitra_group_prod_lim_status,
				m_mitra_group_product_limit.mitra_group_prod_lim_created_at,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_code,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_name,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_status,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_reason,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_action_before,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_created_by,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_created_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_group_product_limit_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_group_product_limit on m_mitra_group_product_limit_temp.mitra_group_prod_lim_id=m_mitra_group_product_limit.mitra_group_prod_lim_id").
		Where("mitra_group_prod_lim_temp_id = ?", pid).
		Take(&mitragroupproductlimit).Error
	if err != nil {
		return &MitraGroupProductLimitTempData{}, err
	}
	return &mitragroupproductlimit, nil
}

func (p *MitraGroupProductLimitTemp) FindMitraGroupProductLimitTempByIDMitra(db *gorm.DB, pid uint64, mitra uint64) (*MitraGroupProductLimitTempData, error) {
	var err error
	mitragroupproductlimit := MitraGroupProductLimitTempData{}
	err = db.Debug().Model(&MitraGroupProductLimitTempData{}).
		Select(`m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_id,
				m_mitra_group_product_limit_temp.mitra_temp_id,
				m_mitra_temp.mitra_code as mitra_temp_code,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_temp.mitra_website as mitra_temp_website,
				m_mitra_temp.mitra_email as mitra_temp_email,
				m_mitra_temp.mitra_logo as mitra_temp_logo,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_id,
				m_mitra_group_product_limit.mitra_group_prod_lim_code,
				m_mitra_group_product_limit.mitra_group_prod_lim_name,
				m_mitra_group_product_limit.mitra_group_prod_lim_status,
				m_mitra_group_product_limit.mitra_group_prod_lim_created_at,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_code,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_name,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_status,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_action_before,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_reason,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_created_by,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_created_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_group_product_limit_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_group_product_limit on m_mitra_group_product_limit_temp.mitra_group_prod_lim_id=m_mitra_group_product_limit.mitra_group_prod_lim_id").
		Where("mitra_group_prod_lim_temp_id = ? AND m_mitra_group_product_limit_temp.mitra_temp_id = ?", pid, mitra).
		Take(&mitragroupproductlimit).Error
	if err != nil {
		return &MitraGroupProductLimitTempData{}, err
	}
	return &mitragroupproductlimit, nil
}

func (p *MitraGroupProductLimitTemp) FindMitraGroupProductLimitTempStatusByID(db *gorm.DB, pid uint64, status uint64, mitra uint64) (*MitraGroupProductLimitTempData, error) {
	var err error
	mitragroupproductlimit := MitraGroupProductLimitTempData{}
	err = db.Debug().Model(&MitraGroupProductLimitTempData{}).
		Select(`m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_id,
				m_mitra_group_product_limit_temp.mitra_temp_id,
				m_mitra_temp.mitra_code as mitra_temp_code,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_temp.mitra_website as mitra_temp_website,
				m_mitra_temp.mitra_email as mitra_temp_email,
				m_mitra_temp.mitra_logo as mitra_temp_logo,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_id,
				m_mitra_group_product_limit.mitra_group_prod_lim_code,
				m_mitra_group_product_limit.mitra_group_prod_lim_name,
				m_mitra_group_product_limit.mitra_group_prod_lim_status,
				m_mitra_group_product_limit.mitra_group_prod_lim_created_at,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_code,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_name,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_status,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_action_before,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_reason,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_created_by,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_created_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_group_product_limit_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_group_product_limit on m_mitra_group_product_limit_temp.mitra_group_prod_lim_id=m_mitra_group_product_limit.mitra_group_prod_lim_id").
		Where("mitra_group_prod_lim_temp_id = ? AND mitra_group_prod_lim_temp_status = ? AND mitra_temp_id = ?", pid, status, mitra).
		Take(&mitragroupproductlimit).Error
	if err != nil {
		return &MitraGroupProductLimitTempData{}, err
	}
	return &mitragroupproductlimit, nil
}

func (p *MitraGroupProductLimitTemp) UpdateMitraGroupProductLimitTemp(db *gorm.DB, pid uint64) (*MitraGroupProductLimitTemp, error) {

	var err error
	db = db.Debug().Model(&MitraGroupProductLimitTemp{}).Where("mitra_group_prod_lim_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_temp_id":                        p.MitraTempID,
			"mitra_group_prod_lim_temp_code":       p.MitraGroupProductLimitTempCode,
			"mitra_group_prod_lim_temp_name":       p.MitraGroupProductLimitTempName,
			"mitra_group_prod_lim_temp_status":     p.MitraGroupProductLimitTempStatus,
			"mitra_group_prod_lim_temp_reason":     p.MitraGroupProductLimitTempReason,
			"mitra_group_prod_lim_temp_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraGroupProductLimitTemp{}).Where("mitra_group_prod_lim_temp_id = ?", pid).Error
	if err != nil {
		return &MitraGroupProductLimitTemp{}, err
	}
	return p, nil
}

func (p *MitraGroupProductLimitTemp) UpdateMitraGroupProductLimitTempStatus(db *gorm.DB, pid uint64) (*MitraGroupProductLimitTemp, error) {

	var err error
	db = db.Debug().Model(&MitraGroupProductLimitTemp{}).Where("mitra_group_prod_lim_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_group_prod_lim_temp_status":     p.MitraGroupProductLimitTempStatus,
			"mitra_group_prod_lim_temp_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraGroupProductLimitTemp{}).Where("mitra_group_prod_lim_temp_id = ?", pid).Error
	if err != nil {
		return &MitraGroupProductLimitTemp{}, err
	}
	return p, nil
}

func (p *MitraGroupProductLimitTemp) DeleteMitraGroupProductLimitTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraGroupProductLimitTemp{}).Where("mitra_group_prod_lim_temp_id = ?", pid).Take(&MitraGroupProductLimitTemp{}).Delete(&MitraGroupProductLimitTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraGroupProductLimitTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================
