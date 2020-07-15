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

type MitraCoverageRegencyTemp struct {
	MitraCoverageRegencyTempID           uint64    `gorm:"column:mitra_cov_rgcy_temp_id;primary_key;" json:"mitra_cov_rgcy_temp_id"`
	MitraCoverageRegencyID               uint64    `gorm:"column:mitra_cov_rgcy_id" json:"mitra_cov_rgcy_id"`
	MitraTempID                          uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	RegencyTempID                        uint64    `gorm:"column:regency_temp_id" json:"regency_temp_id"`
	MitraCoverageRegencyTempReason       string    `gorm:"column:mitra_cov_rgcy_temp_reason" json:"mitra_cov_rgcy_temp_reason"`
	MitraCoverageRegencyTempStatus       uint64    `gorm:"column:mitra_cov_rgcy_temp_status;size:2" json:"mitra_cov_rgcy_temp_status"`
	MitraCoverageRegencyTempActionBefore uint64    `gorm:"column:mitra_cov_rgcy_temp_action_before;size:2" json:"mitra_cov_rgcy_temp_action_before"`
	MitraCoverageRegencyTempCreatedBy    string    `gorm:"column:mitra_cov_rgcy_temp_created_by;size:125" json:"mitra_cov_rgcy_temp_created_by"`
	MitraCoverageRegencyTempCreatedAt    time.Time `gorm:"column:mitra_cov_rgcy_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_cov_rgcy_temp_created_at"`
}

type MitraCoverageRegencyTempPend struct {
	MitraCoverageRegencyTempID           *uint64   `gorm:"column:mitra_cov_rgcy_temp_id;primary_key;" json:"mitra_cov_rgcy_temp_id"`
	MitraCoverageRegencyID               *uint64   `gorm:"column:mitra_cov_rgcy_id" json:"mitra_cov_rgcy_id"`
	MitraTempID                          uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	RegencyTempID                        uint64    `gorm:"column:regency_temp_id" json:"regency_temp_id"`
	MitraCoverageRegencyTempReason       string    `gorm:"column:mitra_cov_rgcy_temp_reason" json:"mitra_cov_rgcy_temp_reason"`
	MitraCoverageRegencyTempStatus       uint64    `gorm:"column:mitra_cov_rgcy_temp_status;size:2" json:"mitra_cov_rgcy_temp_status"`
	MitraCoverageRegencyTempActionBefore uint64    `gorm:"column:mitra_cov_rgcy_temp_action_before;size:2" json:"mitra_cov_rgcy_temp_action_before"`
	MitraCoverageRegencyTempCreatedBy    string    `gorm:"column:mitra_cov_rgcy_temp_created_by;size:125" json:"mitra_cov_rgcy_temp_created_by"`
	MitraCoverageRegencyTempCreatedAt    time.Time `gorm:"column:mitra_cov_rgcy_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_cov_rgcy_temp_created_at"`
}

type MitraCoverageRegencyTempPendData struct {
	MitraCoverageRegencyTempID           uint64    `gorm:"column:mitra_cov_rgcy_temp_id;primary_key;" json:"mitra_cov_rgcy_temp_id"`
	MitraCoverageRegencyID               uint64    `gorm:"column:mitra_cov_rgcy_id" json:"mitra_cov_rgcy_id"`
	MitraTempID                          uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	RegencyTempID                        uint64    `gorm:"column:regency_temp_id" json:"regency_temp_id"`
	RegencyTempName                      string    `gorm:"column:regency_temp_name" json:"regency_temp_name"`
	MitraCoverageRegencyTempReason       string    `gorm:"column:mitra_cov_rgcy_temp_reason" json:"mitra_cov_rgcy_temp_reason"`
	MitraCoverageRegencyTempStatus       uint64    `gorm:"column:mitra_cov_rgcy_temp_status;size:2" json:"mitra_cov_rgcy_temp_status"`
	MitraCoverageRegencyTempActionBefore uint64    `gorm:"column:mitra_cov_rgcy_temp_action_before;size:2" json:"mitra_cov_rgcy_temp_action_before"`
	MitraCoverageRegencyTempCreatedBy    string    `gorm:"column:mitra_cov_rgcy_temp_created_by;size:125" json:"mitra_cov_rgcy_temp_created_by"`
	MitraCoverageRegencyTempCreatedAt    time.Time `gorm:"column:mitra_cov_rgcy_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_cov_rgcy_temp_created_at"`
}

type MitraCoverageRegencyTempData struct {
	MitraCoverageRegencyTempID           uint64    `gorm:"column:mitra_cov_rgcy_temp_id" json:"mitra_cov_rgcy_temp_id"`
	MitraTempID                          uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	MitraTempName                        string    `gorm:"column:mitra_temp_name" json:"mitra_temp_name"`
	RegencyTempID                        uint64    `gorm:"column:regency_temp_id" json:"regency_temp_id"`
	RegencyTempName                      string    `gorm:"column:regency_temp_name" json:"regency_temp_name"`
	MitraID                              uint64    `gorm:"column:mitra_id" json:"mitra_id"`
	MitraName                            string    `gorm:"column:mitra_name" json:"mitra_name"`
	RegencyID                            uint64    `gorm:"column:regency_id" json:"regency_id"`
	RegencyName                          string    `gorm:"column:regency_name" json:"regency_name"`
	MitraCoverageRegencyTempReason       string    `gorm:"column:mitra_cov_rgcy_temp_reason" json:"mitra_cov_rgcy_temp_reason"`
	MitraCoverageRegencyTempStatus       uint64    `gorm:"column:mitra_cov_rgcy_temp_status;size:2" json:"mitra_cov_rgcy_temp_status"`
	MitraCoverageRegencyTempActionBefore uint64    `gorm:"column:mitra_cov_rgcy_temp_action_before;size:2" json:"mitra_cov_rgcy_temp_action_before"`
	MitraCoverageRegencyTempCreatedBy    string    `gorm:"column:mitra_cov_rgcy_temp_created_by;size:125" json:"mitra_cov_rgcy_temp_created_by"`
	MitraCoverageRegencyTempCreatedAt    time.Time `gorm:"column:mitra_cov_rgcy_temp_created_at" json:"mitra_cov_rgcy_temp_created_at"`
	MitraCoverageRegencyID               uint64    `gorm:"column:mitra_cov_rgcy_id" json:"mitra_cov_rgcy_id"`
}

type ResponseMitraCoverageRegencyTemps struct {
	Status  int                             `json:"status"`
	Message string                          `json:"message"`
	Data    *[]MitraCoverageRegencyTempData `json:"data"`
}

type ResponseMitraCoverageRegencyTemp struct {
	Status  int                           `json:"status"`
	Message string                        `json:"message"`
	Data    *MitraCoverageRegencyTempData `json:"data"`
}

type ResponseMitraCoverageRegencyTempsPend struct {
	Status  int                             `json:"status"`
	Message string                          `json:"message"`
	Data    *[]MitraCoverageRegencyTempPend `json:"data"`
}

type ResponseMitraCoverageRegencyTempPend struct {
	Status  int                           `json:"status"`
	Message string                        `json:"message"`
	Data    *MitraCoverageRegencyTempPend `json:"data"`
}

type ResponseMitraCoverageRegencyTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraCoverageRegencyTemp) TableName() string {
	return "m_mitra_coverage_regency_temp"
}

func (MitraCoverageRegencyTempPend) TableName() string {
	return "m_mitra_coverage_regency_temp"
}

func (MitraCoverageRegencyTempPendData) TableName() string {
	return "m_mitra_coverage_regency_temp"
}

func (MitraCoverageRegencyTempData) TableName() string {
	return "m_mitra_coverage_regency_temp"
}

func (p *MitraCoverageRegencyTemp) Prepare() {
	p.MitraCoverageRegencyID = p.MitraCoverageRegencyID
	p.MitraTempID = p.MitraTempID
	p.RegencyTempID = p.RegencyTempID
	p.MitraCoverageRegencyTempReason = p.MitraCoverageRegencyTempReason
	p.MitraCoverageRegencyTempStatus = p.MitraCoverageRegencyTempStatus
	p.MitraCoverageRegencyTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraCoverageRegencyTempCreatedBy))
	p.MitraCoverageRegencyTempCreatedAt = time.Now()
}

func (p *MitraCoverageRegencyTempPend) Prepare() {
	p.MitraCoverageRegencyTempID = nil
	p.MitraCoverageRegencyID = nil
	p.MitraTempID = p.MitraTempID
	p.RegencyTempID = p.RegencyTempID
	p.MitraCoverageRegencyTempReason = p.MitraCoverageRegencyTempReason
	p.MitraCoverageRegencyTempStatus = p.MitraCoverageRegencyTempStatus
	p.MitraCoverageRegencyTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraCoverageRegencyTempCreatedBy))
	p.MitraCoverageRegencyTempCreatedAt = time.Now()
}

func (p *MitraCoverageRegencyTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	case "insertupdate":
		if p.MitraTempID == 0 {
			return errors.New("Negara Wajib diisi")
		}
		if p.RegencyTempID == 0 {
			return errors.New("Provinsi Wajib diisi")
		}
		if p.MitraCoverageRegencyTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.MitraCoverageRegencyTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	case "reason":
		if p.MitraCoverageRegencyTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.MitraCoverageRegencyTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	default:
		if p.MitraCoverageRegencyTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil
	}
}

func (p *MitraCoverageRegencyTemp) SaveMitraCoverageRegencyTemp(db *gorm.DB) (*MitraCoverageRegencyTemp, error) {
	var err error
	err = db.Debug().Model(&MitraCoverageRegencyTemp{}).Create(&p).Error
	if err != nil {
		return &MitraCoverageRegencyTemp{}, err
	}
	return p, nil
}

func (p *MitraCoverageRegencyTempPend) SaveMitraCoverageRegencyTempPend(db *gorm.DB) (*MitraCoverageRegencyTempPend, error) {
	var err error
	err = db.Debug().Model(&MitraCoverageRegencyTempPend{}).Create(&p).Error
	if err != nil {
		return &MitraCoverageRegencyTempPend{}, err
	}
	return p, nil
}

func (p *MitraCoverageRegencyTemp) FindAllMitraCoverageRegencyTemp(db *gorm.DB) (*[]MitraCoverageRegencyTempPend, error) {
	var err error
	mitracoverageregency := []MitraCoverageRegencyTempPend{}
	err = db.Debug().Model(&MitraCoverageRegencyTempPend{}).Limit(100).
		Select(`m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_id,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_id,
				m_mitra_coverage_regency_temp.mitra_temp_id,
				m_mitra_coverage_regency_temp.regency_temp_id,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_reason,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_status,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_action_before,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_created_by,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_created_at`).
		Find(&mitracoverageregency).Error
	if err != nil {
		return &[]MitraCoverageRegencyTempPend{}, err
	}
	return &mitracoverageregency, nil
}

func (p *MitraCoverageRegencyTemp) FindAllMitraCoverageRegencyTempPendingActive(db *gorm.DB) (*[]MitraCoverageRegencyTempPend, error) {
	var err error
	mitracoverageregency := []MitraCoverageRegencyTempPend{}
	err = db.Debug().Model(&MitraCoverageRegencyTempPend{}).Limit(100).
		Select(`m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_id,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_id,
				m_mitra_coverage_regency_temp.mitra_temp_id,
				m_mitra_coverage_regency_temp.regency_temp_id,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_reason,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_status,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_action_before,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_created_by,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_created_at`).
		Where("mitra_cov_rgcy_temp_status = ?", 11).
		Find(&mitracoverageregency).Error
	if err != nil {
		return &[]MitraCoverageRegencyTempPend{}, err
	}
	return &mitracoverageregency, nil
}

func (p *MitraCoverageRegencyTemp) FindAllMitraCoverageRegencyTempStatus(db *gorm.DB, status uint64) (*[]MitraCoverageRegencyTempData, error) {
	var err error
	mitracoverageregency := []MitraCoverageRegencyTempData{}
	err = db.Debug().Model(&MitraCoverageRegencyTempData{}).Limit(100).
		Select(`m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_regency_temp.regency_id as regency_temp_id,
				m_regency_temp.regency_name as regency_temp_name,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_reason,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_status,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_action_before,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_created_by,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_created_at,
				m_mitra_coverage_regency.mitra_cov_rgcy_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_mitra_coverage_regency.mitra_cov_rgcy_status,
				m_mitra_coverage_regency.mitra_cov_rgcy_created_by,
				m_mitra_coverage_regency.mitra_cov_rgcy_created_at,
				m_mitra_coverage_regency.mitra_cov_rgcy_updated_by,
				m_mitra_coverage_regency.mitra_cov_rgcy_updated_at,
				m_mitra_coverage_regency.mitra_cov_rgcy_deleted_by,
				m_mitra_coverage_regency.mitra_cov_rgcy_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_coverage_regency_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_regency m_regency_temp on m_mitra_coverage_regency_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_mitra_coverage_regency on m_mitra_coverage_regency_temp.mitra_cov_rgcy_id=m_mitra_coverage_regency.mitra_cov_rgcy_id").
		Joins("JOIN m_mitra on m_mitra_coverage_regency.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_regency on m_mitra_coverage_regency.regency_id=m_regency.regency_id").
		Where("mitra_cov_rgcy_temp_status = ?", status).
		Find(&mitracoverageregency).Error
	if err != nil {
		return &[]MitraCoverageRegencyTempData{}, err
	}
	return &mitracoverageregency, nil
}

func (p *MitraCoverageRegencyTemp) FindMitraCoverageRegencyTempDataByID(db *gorm.DB, pid uint64) (*MitraCoverageRegencyTemp, error) {
	var err error
	err = db.Debug().Model(&MitraCoverageRegencyTemp{}).Where("mitra_cov_rgcy_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraCoverageRegencyTemp{}, err
	}
	return p, nil
}

func (p *MitraCoverageRegencyTemp) FindMitraCoverageRegencyTempByIDPendingActive(db *gorm.DB, pid uint64) (*MitraCoverageRegencyTempPend, error) {
	var err error
	mitracoverageregency := MitraCoverageRegencyTempPend{}
	err = db.Debug().Model(&MitraCoverageRegencyTempPend{}).Limit(100).
		Select(`m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_id,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_id,
				m_mitra_coverage_regency_temp.mitra_temp_id,
				m_mitra_coverage_regency_temp.regency_temp_id,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_reason,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_status,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_action_before,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_created_by,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_created_at`).
		Where("mitra_cov_rgcy_temp_id = ? AND mitra_cov_rgcy_temp_status = ?", pid, 11).
		Find(&mitracoverageregency).Error
	if err != nil {
		return &MitraCoverageRegencyTempPend{}, err
	}
	return &mitracoverageregency, nil
}

func (p *MitraCoverageRegencyTemp) FindMitraCoverageRegencyTempByID(db *gorm.DB, pid uint64) (*MitraCoverageRegencyTempData, error) {
	var err error
	mitracoverageregency := MitraCoverageRegencyTempData{}
	err = db.Debug().Model(&MitraCoverageRegencyTempData{}).Limit(100).
		Select(`m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_regency_temp.regency_id as regency_temp_id,
				m_regency_temp.regency_name as regency_temp_name,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_reason,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_status,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_action_before,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_created_by,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_created_at,
				m_mitra_coverage_regency.mitra_cov_rgcy_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_mitra_coverage_regency.mitra_cov_rgcy_status,
				m_mitra_coverage_regency.mitra_cov_rgcy_created_by,
				m_mitra_coverage_regency.mitra_cov_rgcy_created_at,
				m_mitra_coverage_regency.mitra_cov_rgcy_updated_by,
				m_mitra_coverage_regency.mitra_cov_rgcy_updated_at,
				m_mitra_coverage_regency.mitra_cov_rgcy_deleted_by,
				m_mitra_coverage_regency.mitra_cov_rgcy_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_coverage_regency_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_regency m_regency_temp on m_mitra_coverage_regency_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_mitra_coverage_regency on m_mitra_coverage_regency_temp.mitra_cov_rgcy_id=m_mitra_coverage_regency.mitra_cov_rgcy_id").
		Joins("JOIN m_mitra on m_mitra_coverage_regency.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_regency on m_mitra_coverage_regency.regency_id=m_regency.regency_id").
		Where("mitra_cov_rgcy_temp_id = ?", pid).
		Find(&mitracoverageregency).Error
	if err != nil {
		return &MitraCoverageRegencyTempData{}, err
	}
	return &mitracoverageregency, nil
}

func (p *MitraCoverageRegencyTemp) FindMitraCoverageRegencyTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraCoverageRegencyTempData, error) {
	var err error
	mitracoverageregency := MitraCoverageRegencyTempData{}
	err = db.Debug().Model(&MitraCoverageRegencyTempData{}).Limit(100).
		Select(`m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_regency_temp.regency_id as regency_temp_id,
				m_regency_temp.regency_name as regency_temp_name,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_reason,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_status,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_action_before,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_created_by,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_created_at,
				m_mitra_coverage_regency.mitra_cov_rgcy_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_mitra_coverage_regency.mitra_cov_rgcy_status,
				m_mitra_coverage_regency.mitra_cov_rgcy_created_by,
				m_mitra_coverage_regency.mitra_cov_rgcy_created_at,
				m_mitra_coverage_regency.mitra_cov_rgcy_updated_by,
				m_mitra_coverage_regency.mitra_cov_rgcy_updated_at,
				m_mitra_coverage_regency.mitra_cov_rgcy_deleted_by,
				m_mitra_coverage_regency.mitra_cov_rgcy_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_coverage_regency_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_regency m_regency_temp on m_mitra_coverage_regency_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_mitra_coverage_regency on m_mitra_coverage_regency_temp.mitra_cov_rgcy_id=m_mitra_coverage_regency.mitra_cov_rgcy_id").
		Joins("JOIN m_mitra on m_mitra_coverage_regency.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_regency on m_mitra_coverage_regency.regency_id=m_regency.regency_id").
		Where("mitra_cov_rgcy_temp_id = ? AND mitra_cov_rgcy_temp_status = ?", pid, status).
		Find(&mitracoverageregency).Error
	if err != nil {
		return &MitraCoverageRegencyTempData{}, err
	}
	return &mitracoverageregency, nil
}

func (p *MitraCoverageRegencyTemp) UpdateMitraCoverageRegencyTemp(db *gorm.DB, pid uint64) (*MitraCoverageRegencyTemp, error) {
	var err error
	db = db.Debug().Model(&MitraCoverageRegencyTemp{}).Where("mitra_cov_rgcy_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_temp_id":              p.MitraTempID,
			"regency_temp_id":            p.RegencyTempID,
			"mitra_cov_rgcy_temp_reason": p.MitraCoverageRegencyTempReason,
		},
	)
	err = db.Debug().Model(&MitraCoverageRegencyTemp{}).Where("mitra_cov_rgcy_temp_id = ?", pid).Error
	if err != nil {
		return &MitraCoverageRegencyTemp{}, err
	}
	return p, nil
}

func (p *MitraCoverageRegencyTemp) UpdateMitraCoverageRegencyTempStatus(db *gorm.DB, pid uint64) (*MitraCoverageRegencyTemp, error) {
	var err error
	db = db.Debug().Model(&MitraCoverageRegencyTemp{}).Where("mitra_cov_rgcy_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_cov_rgcy_temp_reason": p.MitraCoverageRegencyTempReason,
			"mitra_cov_rgcy_temp_status": p.MitraCoverageRegencyTempStatus,
		},
	)
	err = db.Debug().Model(&MitraCoverageRegencyTemp{}).Where("mitra_cov_rgcy_temp_id = ?", pid).Error
	if err != nil {
		return &MitraCoverageRegencyTemp{}, err
	}
	return p, nil
}

func (p *MitraCoverageRegencyTemp) DeleteMitraCoverageRegencyTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraCoverageRegencyTemp{}).Where("mitra_cov_rgcy_temp_id = ?", pid).Take(&MitraCoverageRegencyTemp{}).Delete(&MitraCoverageRegencyTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraCoverageRegencyTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

func (p *MitraCoverageRegencyTemp) FindMitraCoverageRegencyTempPendByMitraTempID(db *gorm.DB, pid uint64) ([]MitraCoverageRegencyTempPendData, error) {
	var err error
	mitracoverageregency := []MitraCoverageRegencyTempPendData{}
	err = db.Debug().Model(&MitraCoverageRegencyTempPendData{}).Limit(100).
		Select(`m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_id,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_id,
				m_mitra_coverage_regency_temp.mitra_temp_id,
				m_mitra_temp.mitra_temp_name,
				m_mitra_coverage_regency_temp.regency_temp_id,
				m_regency_temp.regency_name as regency_temp_name,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_reason,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_status,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_action_before,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_created_by,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_created_at`).
		Joins("JOIN m_regency m_regency_temp on m_mitra_coverage_regency_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_mitra_temp on m_mitra_coverage_regency_temp.mitra_temp_id=m_mitra_temp.mitra_temp_id").
		Where("m_mitra_coverage_regency_temp.mitra_temp_id = ?", pid).
		Find(&mitracoverageregency).Error
	if err != nil {
		return []MitraCoverageRegencyTempPendData{}, err
	}
	return mitracoverageregency, nil
}

func (p *MitraCoverageRegencyTemp) FindMitraCoverageRegencyTempPendStatusByMitraTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraCoverageRegencyTempPendData, error) {
	var err error
	mitracoverageregency := []MitraCoverageRegencyTempPendData{}
	err = db.Debug().Model(&MitraCoverageRegencyTempPendData{}).Limit(100).
		Select(`m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_id,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_id,
				m_mitra_coverage_regency_temp.mitra_temp_id,
				m_mitra_temp.mitra_temp_name,
				m_mitra_coverage_regency_temp.regency_temp_id,
				m_regency_temp.regency_name as regency_temp_name,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_reason,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_status,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_action_before,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_created_by,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_created_at`).
		Joins("JOIN m_regency m_regency_temp on m_mitra_coverage_regency_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_mitra_temp on m_mitra_coverage_regency_temp.mitra_temp_id=m_mitra_temp.mitra_temp_id").
		Where("m_mitra_coverage_regency_temp.mitra_temp_id = ? AND mitra_cov_rgcy_temp_status = ?", pid, status).
		Find(&mitracoverageregency).Error
	if err != nil {
		return []MitraCoverageRegencyTempPendData{}, err
	}
	return mitracoverageregency, nil
}

func (p *MitraCoverageRegencyTemp) FindMitraCoverageRegencyTempByMitraTempID(db *gorm.DB, pid uint64) ([]MitraCoverageRegencyTempData, error) {
	var err error
	mitracoverageregency := []MitraCoverageRegencyTempData{}
	err = db.Debug().Model(&MitraCoverageRegencyTempData{}).Limit(100).
		Select(`m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_id,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_id,
				m_mitra_coverage_regency_temp.mitra_temp_id,
				m_mitra_temp.mitra_temp_name,
				m_mitra_coverage_regency_temp.regency_temp_id,
				m_regency_temp.regency_name as regency_temp_name,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_reason,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_status,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_action_before,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_created_by,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_created_at`).
		Joins("JOIN m_regency m_regency_temp on m_mitra_coverage_regency_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_mitra_temp on m_mitra_coverage_regency_temp.mitra_temp_id=m_mitra_temp.mitra_temp_id").
		Where("m_mitra_coverage_regency_temp.mitra_temp_id = ?", pid).
		Find(&mitracoverageregency).Error
	if err != nil {
		return []MitraCoverageRegencyTempData{}, err
	}
	return mitracoverageregency, nil
}

func (p *MitraCoverageRegencyTemp) FindMitraCoverageRegencyTempStatusByMitraTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraCoverageRegencyTempData, error) {
	var err error
	mitracoverageregency := []MitraCoverageRegencyTempData{}
	err = db.Debug().Model(&MitraCoverageRegencyTempData{}).Limit(100).
		Select(`m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_id,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_id,
				m_mitra_coverage_regency_temp.mitra_temp_id,
				m_mitra_temp.mitra_temp_name,
				m_mitra_coverage_regency_temp.regency_temp_id,
				m_regency_temp.regency_name as regency_temp_name,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_reason,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_status,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_action_before,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_created_by,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_created_at`).
		Joins("JOIN m_regency m_regency_temp on m_mitra_coverage_regency_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_mitra_temp on m_mitra_coverage_regency_temp.mitra_temp_id=m_mitra_temp.mitra_temp_id").
		Where("m_mitra_coverage_regency_temp.mitra_temp_id = ? AND mitra_cov_rgcy_temp_status = ?", pid, status).
		Find(&mitracoverageregency).Error
	if err != nil {
		return []MitraCoverageRegencyTempData{}, err
	}
	return mitracoverageregency, nil
}

func (p *MitraCoverageRegencyTemp) FindMitraCoverageRegencyTempDataStatusByMitraTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraCoverageRegencyTempData, error) {
	var err error
	mitracoverageregency := []MitraCoverageRegencyTempData{}
	err = db.Debug().Model(&MitraCoverageRegencyTempData{}).Limit(100).
		Select(`m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_id,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_id,
				m_mitra_temp.mitra_temp_id,
				m_mitra_temp.mitra_temp_name,
				m_mitra_coverage_regency_temp.regency_temp_id,
				m_regency_temp.regency_name as regency_temp_name,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_reason,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_status,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_action_before,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_created_by,
				m_mitra_coverage_regency_temp.mitra_cov_rgcy_temp_created_at`).
		Joins("JOIN m_regency m_regency_temp on m_mitra_coverage_regency_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_mitra_temp on m_mitra_coverage_regency_temp.mitra_temp_id=m_mitra_temp.mitra_temp_id").
		Joins("JOIN m_mitra_coverage_regency on m_mitra_coverage_regency_temp.mitra_cov_rgcy_id=m_mitra_coverage_regency.mitra_cov_rgcy_id").
		Joins("JOIN m_regency on m_mitra_coverage_regency.regency_id=m_regency.regency_id").
		Joins("JOIN m_mitra on m_mitra_coverage_regency.mitra_id=m_mitra.mitra_id").
		Where("m_mitra_coverage_regency_temp.mitra_temp_id = ? AND mitra_cov_rgcy_temp_status = ?", pid, status).
		Find(&mitracoverageregency).Error
	if err != nil {
		return []MitraCoverageRegencyTempData{}, err
	}
	return mitracoverageregency, nil
}
