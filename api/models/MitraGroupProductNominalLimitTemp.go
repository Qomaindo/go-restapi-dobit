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

type MitraGroupProductNominalLimitTemp struct {
	MitraGroupProductNominalLimitTempID           *uint64   `gorm:"column:mitra_group_prod_nom_lim_temp_id;primary_key;" json:"mitra_group_prod_nom_lim_temp_id"`
	MitraGroupProductNominalLimitID               uint64    `gorm:"column:mitra_group_prod_nom_lim_id" json:"mitra_group_prod_nom_lim_id"`
	MitraGroupProductLimitTempID                  uint64    `gorm:"column:mitra_group_prod_lim_temp_id" json:"mitra_group_prod_lim_temp_id"`
	MitraProductTempID                            uint64    `gorm:"column:mitra_prod_temp_id" json:"mitra_prod_temp_id"`
	MitraGroupProductNominalLimitTempRole         string    `gorm:"column:mitra_group_prod_nom_lim_temp_role" json:"mitra_group_prod_nom_lim_temp_role"`
	MitraGroupProductNominalLimitTempMin          uint64    `gorm:"column:mitra_group_prod_nom_lim_temp_min" json:"mitra_group_prod_nom_lim_temp_min"`
	MitraGroupProductNominalLimitTempMax          uint64    `gorm:"column:mitra_group_prod_nom_lim_temp_max" json:"mitra_group_prod_nom_lim_temp_max"`
	MitraGroupProductNominalLimitTempReason       string    `gorm:"column:mitra_group_prod_nom_lim_temp_reason" json:"mitra_group_prod_nom_lim_temp_reason"`
	MitraGroupProductNominalLimitTempStatus       uint64    `gorm:"column:mitra_group_prod_nom_lim_temp_status;size:2" json:"mitra_group_prod_nom_lim_temp_status"`
	MitraGroupProductNominalLimitTempActionBefore uint64    `gorm:"column:mitra_group_prod_nom_lim_temp_action_before;size:2" json:"mitra_group_prod_nom_lim_temp_action_before"`
	MitraGroupProductNominalLimitTempCreatedBy    string    `gorm:"column:mitra_group_prod_nom_lim_temp_created_by;size:125" json:"mitra_group_prod_nom_lim_temp_created_by"`
	MitraGroupProductNominalLimitTempCreatedAt    time.Time `gorm:"column:mitra_group_prod_nom_lim_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_group_prod_nom_lim_temp_created_at"`
}

type MitraGroupProductNominalLimitTempPend struct {
	MitraGroupProductNominalLimitTempID           *uint64   `gorm:"column:mitra_group_prod_nom_lim_temp_id;primary_key;" json:"mitra_group_prod_nom_lim_temp_id"`
	MitraGroupProductNominalLimitID               *uint64   `gorm:"column:mitra_group_prod_nom_lim_id" json:"mitra_group_prod_nom_lim_id"`
	MitraGroupProductLimitTempID                  uint64    `gorm:"column:mitra_group_prod_lim_temp_id" json:"mitra_group_prod_lim_temp_id"`
	MitraProductTempID                            uint64    `gorm:"column:mitra_prod_temp_id" json:"mitra_prod_temp_id"`
	MitraGroupProductNominalLimitTempRole         string    `gorm:"column:mitra_group_prod_nom_lim_temp_role" json:"mitra_group_prod_nom_lim_temp_role"`
	MitraGroupProductNominalLimitTempMin          uint64    `gorm:"column:mitra_group_prod_nom_lim_temp_min" json:"mitra_group_prod_nom_lim_temp_min"`
	MitraGroupProductNominalLimitTempMax          uint64    `gorm:"column:mitra_group_prod_nom_lim_temp_max" json:"mitra_group_prod_nom_lim_temp_max"`
	MitraGroupProductNominalLimitTempReason       string    `gorm:"column:mitra_group_prod_nom_lim_temp_reason" json:"mitra_group_prod_nom_lim_temp_reason"`
	MitraGroupProductNominalLimitTempStatus       uint64    `gorm:"column:mitra_group_prod_nom_lim_temp_status;size:2" json:"mitra_group_prod_nom_lim_temp_status"`
	MitraGroupProductNominalLimitTempActionBefore uint64    `gorm:"column:mitra_group_prod_nom_lim_temp_action_before;size:2" json:"mitra_group_prod_nom_lim_temp_action_before"`
	MitraGroupProductNominalLimitTempCreatedBy    string    `gorm:"column:mitra_group_prod_nom_lim_temp_created_by;size:125" json:"mitra_group_prod_nom_lim_temp_created_by"`
	MitraGroupProductNominalLimitTempCreatedAt    time.Time `gorm:"column:mitra_group_prod_nom_lim_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_group_prod_nom_lim_temp_created_at"`
}

type MitraGroupProductNominalLimitTempPendData struct {
	MitraGroupProductNominalLimitTempID           uint64    `gorm:"column:mitra_group_prod_nom_lim_temp_id;primary_key;" json:"mitra_group_prod_nom_lim_temp_id"`
	MitraGroupProductNominalLimitID               uint64    `gorm:"column:mitra_group_prod_nom_lim_id" json:"mitra_group_prod_nom_lim_id"`
	MitraGroupProductLimitTempID                  uint64    `gorm:"column:mitra_group_prod_lim_temp_id" json:"mitra_group_prod_lim_temp_id"`
	MitraProductTempID                            uint64    `gorm:"column:mitra_prod_temp_id" json:"mitra_prod_temp_id"`
	MitraProductTempName                          string    `gorm:"column:mitra_product_temp_name" json:"mitra_product_temp_name"`
	MitraGroupProductNominalLimitTempRole         string    `gorm:"column:mitra_group_prod_nom_lim_temp_role" json:"mitra_group_prod_nom_lim_temp_role"`
	MitraGroupProductNominalLimitTempMin          uint64    `gorm:"column:mitra_group_prod_nom_lim_temp_min" json:"mitra_group_prod_nom_lim_temp_min"`
	MitraGroupProductNominalLimitTempMax          uint64    `gorm:"column:mitra_group_prod_nom_lim_temp_max" json:"mitra_group_prod_nom_lim_temp_max"`
	MitraGroupProductNominalLimitTempReason       string    `gorm:"column:mitra_group_prod_nom_lim_temp_reason" json:"mitra_group_prod_nom_lim_temp_reason"`
	MitraGroupProductNominalLimitTempStatus       uint64    `gorm:"column:mitra_group_prod_nom_lim_temp_status;size:2" json:"mitra_group_prod_nom_lim_temp_status"`
	MitraGroupProductNominalLimitTempActionBefore uint64    `gorm:"column:mitra_group_prod_nom_lim_temp_action_before;size:2" json:"mitra_group_prod_nom_lim_temp_action_before"`
	MitraGroupProductNominalLimitTempCreatedBy    string    `gorm:"column:mitra_group_prod_nom_lim_temp_created_by;size:125" json:"mitra_group_prod_nom_lim_temp_created_by"`
	MitraGroupProductNominalLimitTempCreatedAt    time.Time `gorm:"column:mitra_group_prod_nom_lim_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_group_prod_nom_lim_temp_created_at"`
}

type MitraGroupProductNominalLimitTempData struct {
	MitraGroupProductNominalLimitTempID           uint64    `gorm:"column:mitra_group_prod_nom_lim_temp_id" json:"mitra_group_prod_nom_lim_temp_id"`
	MitraGroupProductLimitTempID                  uint64    `gorm:"column:mitra_group_prod_lim_temp_id" json:"mitra_group_prod_lim_temp_id"`
	MitraGroupProductLimitTempName                string    `gorm:"column:mitra_group_prod_lim_temp_name" json:"mitra_group_prod_lim_temp_name"`
	MitraProductTempID                            uint64    `gorm:"column:mitra_prod_temp_id" json:"mitra_prod_temp_id"`
	MitraProductTempName                          string    `gorm:"column:mitra_product_temp_name" json:"mitra_product_temp_name"`
	MitraGroupProductLimitID                      uint64    `gorm:"column:mitra_group_prod_lim_id" json:"mitra_group_prod_lim_id"`
	MitraGroupProductLimitName                    string    `gorm:"column:mitra_group_prod_lim_name" json:"mitra_group_prod_lim_name"`
	MitraProductID                                uint64    `gorm:"column:mitra_prod_id" json:"mitra_prod_id"`
	MitraProductName                              string    `gorm:"column:mitra_prod_name" json:"mitra_prod_name"`
	MitraGroupProductNominalLimitTempRole         string    `gorm:"column:mitra_group_prod_nom_lim_temp_role" json:"mitra_group_prod_nom_lim_temp_role"`
	MitraGroupProductNominalLimitTempMin          uint64    `gorm:"column:mitra_group_prod_nom_lim_temp_min" json:"mitra_group_prod_nom_lim_temp_min"`
	MitraGroupProductNominalLimitTempMax          uint64    `gorm:"column:mitra_group_prod_nom_lim_temp_max" json:"mitra_group_prod_nom_lim_temp_max"`
	MitraGroupProductNominalLimitTempReason       string    `gorm:"column:mitra_group_prod_nom_lim_temp_reason" json:"mitra_group_prod_nom_lim_temp_reason"`
	MitraGroupProductNominalLimitTempStatus       uint64    `gorm:"column:mitra_group_prod_nom_lim_temp_status;size:2" json:"mitra_group_prod_nom_lim_temp_status"`
	MitraGroupProductNominalLimitTempActionBefore uint64    `gorm:"column:mitra_group_prod_nom_lim_temp_action_before;size:2" json:"mitra_group_prod_nom_lim_temp_action_before"`
	MitraGroupProductNominalLimitTempCreatedBy    string    `gorm:"column:mitra_group_prod_nom_lim_temp_created_by;size:125" json:"mitra_group_prod_nom_lim_temp_created_by"`
	MitraGroupProductNominalLimitTempCreatedAt    time.Time `gorm:"column:mitra_group_prod_nom_lim_temp_created_at" json:"mitra_group_prod_nom_lim_temp_created_at"`
	MitraGroupProductNominalLimitID               uint64    `gorm:"column:mitra_group_prod_nom_lim_id" json:"mitra_group_prod_nom_lim_id"`
}

type MitraGroupProductNominalLimitTempPendInput struct {
	MitraProductTempID                    uint64 `gorm:"column:mitra_prod_temp_id" json:"mitra_prod_temp_id"`
	MitraGroupProductNominalLimitTempRole string `gorm:"column:mitra_group_prod_nom_lim_temp_role" json:"mitra_group_prod_nom_lim_temp_role"`
	MitraGroupProductNominalLimitTempMin  uint64 `gorm:"column:mitra_group_prod_nom_lim_temp_min" json:"mitra_group_prod_nom_lim_temp_min"`
	MitraGroupProductNominalLimitTempMax  uint64 `gorm:"column:mitra_group_prod_nom_lim_temp_max" json:"mitra_group_prod_nom_lim_temp_max"`
}

type ResponseMitraGroupProductNominalLimitTemps struct {
	Status  int                                      `json:"status"`
	Message string                                   `json:"message"`
	Data    *[]MitraGroupProductNominalLimitTempData `json:"data"`
}

type ResponseMitraGroupProductNominalLimitTemp struct {
	Status  int                                    `json:"status"`
	Message string                                 `json:"message"`
	Data    *MitraGroupProductNominalLimitTempData `json:"data"`
}

type ResponseMitraGroupProductNominalLimitTempsPend struct {
	Status  int                                      `json:"status"`
	Message string                                   `json:"message"`
	Data    *[]MitraGroupProductNominalLimitTempPend `json:"data"`
}

type ResponseMitraGroupProductNominalLimitTempPend struct {
	Status  int                                    `json:"status"`
	Message string                                 `json:"message"`
	Data    *MitraGroupProductNominalLimitTempPend `json:"data"`
}

type ResponseMitraGroupProductNominalLimitTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraGroupProductNominalLimitTemp) TableName() string {
	return "m_mitra_group_product_nominal_limit_temp"
}

func (MitraGroupProductNominalLimitTempPend) TableName() string {
	return "m_mitra_group_product_nominal_limit_temp"
}

func (MitraGroupProductNominalLimitTempPendData) TableName() string {
	return "m_mitra_group_product_nominal_limit_temp"
}

func (MitraGroupProductNominalLimitTempData) TableName() string {
	return "m_mitra_group_product_nominal_limit_temp"
}

func (p *MitraGroupProductNominalLimitTemp) Prepare() {
	p.MitraGroupProductNominalLimitTempID = nil
	p.MitraGroupProductNominalLimitID = p.MitraGroupProductNominalLimitID
	p.MitraGroupProductLimitTempID = p.MitraGroupProductLimitTempID
	p.MitraProductTempID = p.MitraProductTempID
	p.MitraGroupProductNominalLimitTempRole = p.MitraGroupProductNominalLimitTempRole
	p.MitraGroupProductNominalLimitTempMin = p.MitraGroupProductNominalLimitTempMin
	p.MitraGroupProductNominalLimitTempMax = p.MitraGroupProductNominalLimitTempMax
	p.MitraGroupProductNominalLimitTempReason = p.MitraGroupProductNominalLimitTempReason
	p.MitraGroupProductNominalLimitTempStatus = p.MitraGroupProductNominalLimitTempStatus
	p.MitraGroupProductNominalLimitTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraGroupProductNominalLimitTempCreatedBy))
	p.MitraGroupProductNominalLimitTempCreatedAt = time.Now()
}

func (p *MitraGroupProductNominalLimitTempPend) Prepare() {
	p.MitraGroupProductNominalLimitTempID = nil
	p.MitraGroupProductNominalLimitID = nil
	p.MitraGroupProductLimitTempID = p.MitraGroupProductLimitTempID
	p.MitraProductTempID = p.MitraProductTempID
	p.MitraGroupProductNominalLimitTempRole = p.MitraGroupProductNominalLimitTempRole
	p.MitraGroupProductNominalLimitTempMin = p.MitraGroupProductNominalLimitTempMin
	p.MitraGroupProductNominalLimitTempMax = p.MitraGroupProductNominalLimitTempMax
	p.MitraGroupProductNominalLimitTempReason = p.MitraGroupProductNominalLimitTempReason
	p.MitraGroupProductNominalLimitTempStatus = p.MitraGroupProductNominalLimitTempStatus
	p.MitraGroupProductNominalLimitTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraGroupProductNominalLimitTempCreatedBy))
	p.MitraGroupProductNominalLimitTempCreatedAt = time.Now()
}

func (p *MitraGroupProductNominalLimitTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	case "insertupdate":
		if p.MitraGroupProductLimitTempID == 0 {
			return errors.New("Negara Wajib diisi")
		}
		if p.MitraProductTempID == 0 {
			return errors.New("Provinsi Wajib diisi")
		}
		if p.MitraGroupProductNominalLimitTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.MitraGroupProductNominalLimitTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	case "reason":
		if p.MitraGroupProductNominalLimitTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.MitraGroupProductNominalLimitTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	default:
		if p.MitraGroupProductNominalLimitTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil
	}
}

func (p *MitraGroupProductNominalLimitTemp) SaveMitraGroupProductNominalLimitTemp(db *gorm.DB) (*MitraGroupProductNominalLimitTemp, error) {
	var err error
	err = db.Debug().Model(&MitraGroupProductNominalLimitTemp{}).Create(&p).Error
	if err != nil {
		return &MitraGroupProductNominalLimitTemp{}, err
	}
	return p, nil
}

func (p *MitraGroupProductNominalLimitTempPend) SaveMitraGroupProductNominalLimitTempPend(db *gorm.DB) (*MitraGroupProductNominalLimitTempPend, error) {
	var err error
	err = db.Debug().Model(&MitraGroupProductNominalLimitTempPend{}).Create(&p).Error
	if err != nil {
		return &MitraGroupProductNominalLimitTempPend{}, err
	}
	return p, nil
}

func (p *MitraGroupProductNominalLimitTemp) FindAllMitraGroupProductNominalLimitTemp(db *gorm.DB) (*[]MitraGroupProductNominalLimitTempPend, error) {
	var err error
	mitragroupproductnominallimit := []MitraGroupProductNominalLimitTempPend{}
	err = db.Debug().Model(&MitraGroupProductNominalLimitTempPend{}).Limit(100).
		Select(`m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_id,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_id,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_lim_temp_id,
				m_mitra_group_product_nominal_limit_temp.mitra_prod_temp_id,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_role,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_min,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_max,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_reason,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_status,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_action_before,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_created_by,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_created_at`).
		Find(&mitragroupproductnominallimit).Error
	if err != nil {
		return &[]MitraGroupProductNominalLimitTempPend{}, err
	}
	return &mitragroupproductnominallimit, nil
}

func (p *MitraGroupProductNominalLimitTemp) FindAllMitraGroupProductNominalLimitTempPendingActive(db *gorm.DB) (*[]MitraGroupProductNominalLimitTempPend, error) {
	var err error
	mitragroupproductnominallimit := []MitraGroupProductNominalLimitTempPend{}
	err = db.Debug().Model(&MitraGroupProductNominalLimitTempPend{}).Limit(100).
		Select(`m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_id,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_id,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_lim_temp_id,
				m_mitra_group_product_nominal_limit_temp.mitra_prod_temp_id,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_role,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_min,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_max,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_reason,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_status,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_action_before,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_created_by,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_created_at`).
		Where("mitra_group_prod_nom_lim_temp_status = ?", 11).
		Find(&mitragroupproductnominallimit).Error
	if err != nil {
		return &[]MitraGroupProductNominalLimitTempPend{}, err
	}
	return &mitragroupproductnominallimit, nil
}

func (p *MitraGroupProductNominalLimitTemp) FindAllMitraGroupProductNominalLimitTempStatus(db *gorm.DB, status uint64) (*[]MitraGroupProductNominalLimitTempData, error) {
	var err error
	mitragroupproductnominallimit := []MitraGroupProductNominalLimitTempData{}
	err = db.Debug().Model(&MitraGroupProductNominalLimitTempData{}).Limit(100).
		Select(`m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_id,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_id as mitra_group_prod_lim_temp_id,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_name as mitra_group_prod_lim_temp_name,
				m_mitra_product_temp.mitra_prod_id as mitra_prod_temp_id,
				m_mitra_product_temp.mitra_prod_name as mitra_product_temp_name,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_role,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_min,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_max,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_reason,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_status,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_action_before,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_created_by,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_created_at,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_id,
				m_mitra_group_product_limit.mitra_group_prod_lim_id,
				m_mitra_group_product_limit.mitra_group_prod_lim_name,
				m_mitra_product.mitra_prod_id,
				m_mitra_product.mitra_prod_name,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_role,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_min,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_max,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_status,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_created_by,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_created_at,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_updated_by,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_updated_at,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_deleted_by,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_deleted_at`).
		Joins("JOIN m_mitra_group_product_limit m_mitra_group_product_limit_temp on m_mitra_group_product_nominal_limit_temp.mitra_group_prod_lim_temp_id=m_mitra_group_product_limit_temp.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_product m_mitra_product_temp on m_mitra_group_product_nominal_limit_temp.mitra_prod_temp_id=m_mitra_product_temp.mitra_prod_id").
		Joins("JOIN m_mitra_group_product_nominal_limit on m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_id=m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_id").
		Joins("JOIN m_mitra_group_product_limit on m_mitra_group_product_nominal_limit.mitra_group_prod_lim_id=m_mitra_group_product_limit.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_product on m_mitra_group_product_nominal_limit.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Where("mitra_group_prod_nom_lim_temp_status = ?", status).
		Find(&mitragroupproductnominallimit).Error
	if err != nil {
		return &[]MitraGroupProductNominalLimitTempData{}, err
	}
	return &mitragroupproductnominallimit, nil
}

func (p *MitraGroupProductNominalLimitTemp) FindMitraGroupProductNominalLimitTempDataByID(db *gorm.DB, pid uint64) (*MitraGroupProductNominalLimitTemp, error) {
	var err error
	err = db.Debug().Model(&MitraGroupProductNominalLimitTemp{}).Where("mitra_group_prod_nom_lim_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraGroupProductNominalLimitTemp{}, err
	}
	return p, nil
}

func (p *MitraGroupProductNominalLimitTemp) FindMitraGroupProductNominalLimitTempByIDPendingActive(db *gorm.DB, pid uint64) (*MitraGroupProductNominalLimitTempPend, error) {
	var err error
	mitragroupproductnominallimit := MitraGroupProductNominalLimitTempPend{}
	err = db.Debug().Model(&MitraGroupProductNominalLimitTempPend{}).Limit(100).
		Select(`m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_id,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_id,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_lim_temp_id,
				m_mitra_group_product_nominal_limit_temp.mitra_prod_temp_id,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_role,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_min,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_max,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_reason,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_status,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_action_before,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_created_by,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_created_at`).
		Where("mitra_group_prod_nom_lim_temp_id = ? AND mitra_group_prod_nom_lim_temp_status = ?", pid, 11).
		Find(&mitragroupproductnominallimit).Error
	if err != nil {
		return &MitraGroupProductNominalLimitTempPend{}, err
	}
	return &mitragroupproductnominallimit, nil
}

func (p *MitraGroupProductNominalLimitTemp) FindMitraGroupProductNominalLimitTempByID(db *gorm.DB, pid uint64) (*MitraGroupProductNominalLimitTempData, error) {
	var err error
	mitragroupproductnominallimit := MitraGroupProductNominalLimitTempData{}
	err = db.Debug().Model(&MitraGroupProductNominalLimitTempData{}).Limit(100).
		Select(`m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_id,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_id as mitra_group_prod_lim_temp_id,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_name as mitra_group_prod_lim_temp_name,
				m_mitra_product_temp.mitra_prod_id as mitra_prod_temp_id,
				m_mitra_product_temp.mitra_prod_name as mitra_product_temp_name,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_min,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_max,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_reason,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_status,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_action_before,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_created_by,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_created_at,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_id,
				m_mitra_group_product_limit.mitra_group_prod_lim_id,
				m_mitra_group_product_limit.mitra_group_prod_lim_name,
				m_mitra_product.mitra_prod_id,
				m_mitra_product.mitra_prod_name,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_role,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_min,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_max,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_status,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_created_by,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_created_at,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_updated_by,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_updated_at,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_deleted_by,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_deleted_at`).
		Joins("JOIN m_mitra_group_product_limit m_mitra_group_product_limit_temp on m_mitra_group_product_nominal_limit_temp.mitra_group_prod_lim_temp_id=m_mitra_group_product_limit_temp.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_product m_mitra_product_temp on m_mitra_group_product_nominal_limit_temp.mitra_prod_temp_id=m_mitra_product_temp.mitra_prod_id").
		Joins("JOIN m_mitra_group_product_nominal_limit on m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_id=m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_id").
		Joins("JOIN m_mitra_group_product_limit on m_mitra_group_product_nominal_limit.mitra_group_prod_lim_id=m_mitra_group_product_limit.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_product on m_mitra_group_product_nominal_limit.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Where("mitra_group_prod_nom_lim_temp_id = ?", pid).
		Find(&mitragroupproductnominallimit).Error
	if err != nil {
		return &MitraGroupProductNominalLimitTempData{}, err
	}
	return &mitragroupproductnominallimit, nil
}

func (p *MitraGroupProductNominalLimitTemp) FindMitraGroupProductNominalLimitTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraGroupProductNominalLimitTempData, error) {
	var err error
	mitragroupproductnominallimit := MitraGroupProductNominalLimitTempData{}
	err = db.Debug().Model(&MitraGroupProductNominalLimitTempData{}).Limit(100).
		Select(`m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_id,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_id as mitra_group_prod_lim_temp_id,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_name as mitra_group_prod_lim_temp_name,
				m_mitra_product_temp.mitra_prod_id as mitra_prod_temp_id,
				m_mitra_product_temp.mitra_prod_name as mitra_product_temp_name,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_min,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_max,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_reason,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_status,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_action_before,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_created_by,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_created_at,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_id,
				m_mitra_group_product_limit.mitra_group_prod_lim_id,
				m_mitra_group_product_limit.mitra_group_prod_lim_name,
				m_mitra_product.mitra_prod_id,
				m_mitra_product.mitra_prod_name,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_role,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_min,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_max,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_status,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_created_by,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_created_at,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_updated_by,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_updated_at,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_deleted_by,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_deleted_at`).
		Joins("JOIN m_mitra_group_product_limit m_mitra_group_product_limit_temp on m_mitra_group_product_nominal_limit_temp.mitra_group_prod_lim_temp_id=m_mitra_group_product_limit_temp.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_product m_mitra_product_temp on m_mitra_group_product_nominal_limit_temp.mitra_prod_temp_id=m_mitra_product_temp.mitra_prod_id").
		Joins("JOIN m_mitra_group_product_nominal_limit on m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_id=m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_id").
		Joins("JOIN m_mitra_group_product_limit on m_mitra_group_product_nominal_limit.mitra_group_prod_lim_id=m_mitra_group_product_limit.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_product on m_mitra_group_product_nominal_limit.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Where("mitra_group_prod_nom_lim_temp_id = ? AND mitra_group_prod_nom_lim_temp_status = ?", pid, status).
		Find(&mitragroupproductnominallimit).Error
	if err != nil {
		return &MitraGroupProductNominalLimitTempData{}, err
	}
	return &mitragroupproductnominallimit, nil
}

func (p *MitraGroupProductNominalLimitTemp) UpdateMitraGroupProductNominalLimitTemp(db *gorm.DB, pid uint64) (*MitraGroupProductNominalLimitTemp, error) {
	var err error
	db = db.Debug().Model(&MitraGroupProductNominalLimitTemp{}).Where("mitra_group_prod_nom_lim_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_group_prod_lim_temp_id":         p.MitraGroupProductLimitTempID,
			"mitra_prod_temp_id":                   p.MitraProductTempID,
			"mitra_group_prod_nom_lim_temp_role":   p.MitraGroupProductNominalLimitTempRole,
			"mitra_group_prod_nom_lim_temp_min":    p.MitraGroupProductNominalLimitTempMin,
			"mitra_group_prod_nom_lim_temp_max":    p.MitraGroupProductNominalLimitTempMax,
			"mitra_group_prod_nom_lim_temp_reason": p.MitraGroupProductNominalLimitTempReason,
		},
	)
	err = db.Debug().Model(&MitraGroupProductNominalLimitTemp{}).Where("mitra_group_prod_nom_lim_temp_id = ?", pid).Error
	if err != nil {
		return &MitraGroupProductNominalLimitTemp{}, err
	}
	return p, nil
}

func (p *MitraGroupProductNominalLimitTemp) UpdateMitraGroupProductNominalLimitTempStatus(db *gorm.DB, pid uint64) (*MitraGroupProductNominalLimitTemp, error) {
	var err error
	db = db.Debug().Model(&MitraGroupProductNominalLimitTemp{}).Where("mitra_group_prod_nom_lim_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_group_prod_nom_lim_temp_reason": p.MitraGroupProductNominalLimitTempReason,
			"mitra_group_prod_nom_lim_temp_status": p.MitraGroupProductNominalLimitTempStatus,
		},
	)
	err = db.Debug().Model(&MitraGroupProductNominalLimitTemp{}).Where("mitra_group_prod_nom_lim_temp_id = ?", pid).Error
	if err != nil {
		return &MitraGroupProductNominalLimitTemp{}, err
	}
	return p, nil
}

func (p *MitraGroupProductNominalLimitTemp) DeleteMitraGroupProductNominalLimitTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraGroupProductNominalLimitTemp{}).Where("mitra_group_prod_nom_lim_temp_id = ?", pid).Take(&MitraGroupProductNominalLimitTemp{}).Delete(&MitraGroupProductNominalLimitTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraGroupProductNominalLimitTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

func (p *MitraGroupProductNominalLimitTemp) FindMitraGroupProductNominalLimitTempPendactByMitraGroupProductLimitTempID(db *gorm.DB, pid uint64) ([]MitraGroupProductNominalLimitTempPendData, error) {
	var err error
	mitragroupproductnominallimit := []MitraGroupProductNominalLimitTempPendData{}
	err = db.Debug().Model(&MitraGroupProductNominalLimitTempPendData{}).Limit(100).
		Select(`m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_id,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_lim_temp_id,
				m_mitra_product_temp.mitra_prod_id as mitra_prod_temp_id,
				m_mitra_product_temp.mitra_prod_name as mitra_product_temp_name,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_role,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_min,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_max,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_reason,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_status,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_action_before,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_created_by,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_created_at`).
		Joins("JOIN m_mitra_product m_mitra_product_temp on m_mitra_group_product_nominal_limit_temp.mitra_prod_temp_id=m_mitra_product_temp.mitra_prod_id").
		Where("mitra_group_prod_lim_temp_id = ?", pid).
		Find(&mitragroupproductnominallimit).Error
	if err != nil {
		return []MitraGroupProductNominalLimitTempPendData{}, err
	}
	return mitragroupproductnominallimit, nil
}

func (p *MitraGroupProductNominalLimitTemp) FindMitraGroupProductNominalLimitTempByMitraGroupProductLimitTempID(db *gorm.DB, pid uint64) ([]MitraGroupProductNominalLimitTempData, error) {
	var err error
	mitragroupproductnominallimit := []MitraGroupProductNominalLimitTempData{}
	err = db.Debug().Model(&MitraGroupProductNominalLimitTempData{}).Limit(100).
		Select(`m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_id,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_id,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_name,
				m_mitra_product_temp.mitra_prod_id as mitra_prod_temp_id,
				m_mitra_product_temp.mitra_prod_name as mitra_product_temp_name,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_role,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_min,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_max,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_reason,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_status,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_action_before,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_created_by,
				m_mitra_group_product_nominal_limit_temp.mitra_group_prod_nom_lim_temp_created_at`).
		Joins("JOIN m_mitra_group_product_limit_temp on m_mitra_group_product_nominal_limit_temp.mitra_group_prod_lim_temp_id=m_mitra_group_product_limit_temp.mitra_group_prod_lim_temp_id").
		Joins("JOIN m_mitra_product m_mitra_product_temp on m_mitra_group_product_nominal_limit_temp.mitra_prod_temp_id=m_mitra_product_temp.mitra_prod_id").
		Where("m_mitra_group_product_nominal_limit_temp.mitra_group_prod_lim_temp_id = ?", pid).
		Find(&mitragroupproductnominallimit).Error
	if err != nil {
		return []MitraGroupProductNominalLimitTempData{}, err
	}
	return mitragroupproductnominallimit, nil
}
