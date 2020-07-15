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

type MitraAOCoverageRegencyTemp struct {
	MitraAOCoverageRegencyTempID           *uint64   `gorm:"column:mitra_ao_cov_rgcy_temp_id;primary_key;" json:"mitra_ao_cov_rgcy_temp_id"`
	MitraAOCoverageRegencyID               uint64    `gorm:"column:mitra_ao_cov_rgcy_id" json:"mitra_ao_cov_rgcy_id"`
	MitraAOCoverageTempID                  uint64    `gorm:"column:mitra_ao_cov_temp_id" json:"mitra_ao_cov_temp_id"`
	RegencyTempID                          uint64    `gorm:"column:regency_temp_id" json:"regency_temp_id"`
	MitraAOCoverageRegencyTempReason       string    `gorm:"column:mitra_ao_cov_rgcy_temp_reason" json:"mitra_ao_cov_rgcy_temp_reason"`
	MitraAOCoverageRegencyTempStatus       uint64    `gorm:"column:mitra_ao_cov_rgcy_temp_status;size:2" json:"mitra_ao_cov_rgcy_temp_status"`
	MitraAOCoverageRegencyTempActionBefore uint64    `gorm:"column:mitra_ao_cov_rgcy_temp_action_before;size:2" json:"mitra_ao_cov_rgcy_temp_action_before"`
	MitraAOCoverageRegencyTempCreatedBy    string    `gorm:"column:mitra_ao_cov_rgcy_temp_created_by;size:125" json:"mitra_ao_cov_rgcy_temp_created_by"`
	MitraAOCoverageRegencyTempCreatedAt    time.Time `gorm:"column:mitra_ao_cov_rgcy_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_ao_cov_rgcy_temp_created_at"`
}

type MitraAOCoverageRegencyTempPend struct {
	MitraAOCoverageRegencyTempID           *uint64   `gorm:"column:mitra_ao_cov_rgcy_temp_id;primary_key;" json:"mitra_ao_cov_rgcy_temp_id"`
	MitraAOCoverageRegencyID               *uint64   `gorm:"column:mitra_ao_cov_rgcy_id" json:"mitra_ao_cov_rgcy_id"`
	MitraAOCoverageTempID                  uint64    `gorm:"column:mitra_ao_cov_temp_id" json:"mitra_ao_cov_temp_id"`
	RegencyTempID                          uint64    `gorm:"column:regency_temp_id" json:"regency_temp_id"`
	MitraAOCoverageRegencyTempReason       string    `gorm:"column:mitra_ao_cov_rgcy_temp_reason" json:"mitra_ao_cov_rgcy_temp_reason"`
	MitraAOCoverageRegencyTempStatus       uint64    `gorm:"column:mitra_ao_cov_rgcy_temp_status;size:2" json:"mitra_ao_cov_rgcy_temp_status"`
	MitraAOCoverageRegencyTempActionBefore uint64    `gorm:"column:mitra_ao_cov_rgcy_temp_action_before;size:2" json:"mitra_ao_cov_rgcy_temp_action_before"`
	MitraAOCoverageRegencyTempCreatedBy    string    `gorm:"column:mitra_ao_cov_rgcy_temp_created_by;size:125" json:"mitra_ao_cov_rgcy_temp_created_by"`
	MitraAOCoverageRegencyTempCreatedAt    time.Time `gorm:"column:mitra_ao_cov_rgcy_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_ao_cov_rgcy_temp_created_at"`
}

type MitraAOCoverageRegencyTempPendData struct {
	MitraAOCoverageRegencyTempID           uint64    `gorm:"column:mitra_ao_cov_rgcy_temp_id;primary_key;" json:"mitra_ao_cov_rgcy_temp_id"`
	MitraAOCoverageRegencyID               uint64    `gorm:"column:mitra_ao_cov_rgcy_id" json:"mitra_ao_cov_rgcy_id"`
	MitraAOCoverageTempID                  uint64    `gorm:"column:mitra_ao_cov_temp_id" json:"mitra_ao_cov_temp_id"`
	RegencyTempID                          uint64    `gorm:"column:regency_temp_id" json:"regency_temp_id"`
	RegencyTempName                        string    `gorm:"column:regency_temp_name" json:"regency_temp_name"`
	MitraAOCoverageRegencyTempReason       string    `gorm:"column:mitra_ao_cov_rgcy_temp_reason" json:"mitra_ao_cov_rgcy_temp_reason"`
	MitraAOCoverageRegencyTempStatus       uint64    `gorm:"column:mitra_ao_cov_rgcy_temp_status;size:2" json:"mitra_ao_cov_rgcy_temp_status"`
	MitraAOCoverageRegencyTempActionBefore uint64    `gorm:"column:mitra_ao_cov_rgcy_temp_action_before;size:2" json:"mitra_ao_cov_rgcy_temp_action_before"`
	MitraAOCoverageRegencyTempCreatedBy    string    `gorm:"column:mitra_ao_cov_rgcy_temp_created_by;size:125" json:"mitra_ao_cov_rgcy_temp_created_by"`
	MitraAOCoverageRegencyTempCreatedAt    time.Time `gorm:"column:mitra_ao_cov_rgcy_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_ao_cov_rgcy_temp_created_at"`
}

type MitraAOCoverageRegencyTempData struct {
	MitraAOCoverageRegencyTempID           uint64    `gorm:"column:mitra_ao_cov_rgcy_temp_id" json:"mitra_ao_cov_rgcy_temp_id"`
	MitraAOCoverageRegencyID               uint64    `gorm:"column:mitra_ao_cov_rgcy_id" json:"mitra_ao_cov_rgcy_id"`
	MitraAOCoverageTempID                  uint64    `gorm:"column:mitra_ao_cov_temp_id" json:"mitra_ao_cov_temp_id"`
	MitraAOCoverageTempName                string    `gorm:"column:mitra_ao_cov_name" json:"mitra_ao_cov_name"`
	RegencyTempID                          uint64    `gorm:"column:regency_temp_id" json:"regency_temp_id"`
	RegencyTempName                        string    `gorm:"column:regency_temp_name" json:"regency_temp_name"`
	MitraAOCoverageID                      uint64    `gorm:"column:mitra_ao_cov_id" json:"mitra_ao_cov_id"`
	MitraAOCoverageName                    string    `gorm:"column:mitra_ao_cov_name" json:"mitra_ao_cov_name"`
	RegencyID                              uint64    `gorm:"column:regency_id" json:"regency_id"`
	RegencyName                            string    `gorm:"column:regency_name" json:"regency_name"`
	MitraAOCoverageRegencyTempReason       string    `gorm:"column:mitra_ao_cov_rgcy_temp_reason" json:"mitra_ao_cov_rgcy_temp_reason"`
	MitraAOCoverageRegencyTempStatus       uint64    `gorm:"column:mitra_ao_cov_rgcy_temp_status;size:2" json:"mitra_ao_cov_rgcy_temp_status"`
	MitraAOCoverageRegencyTempActionBefore uint64    `gorm:"column:mitra_ao_cov_rgcy_temp_action_before;size:2" json:"mitra_ao_cov_rgcy_temp_action_before"`
	MitraAOCoverageRegencyTempCreatedBy    string    `gorm:"column:mitra_ao_cov_rgcy_temp_created_by;size:125" json:"mitra_ao_cov_rgcy_temp_created_by"`
	MitraAOCoverageRegencyTempCreatedAt    time.Time `gorm:"column:mitra_ao_cov_rgcy_temp_created_at" json:"mitra_ao_cov_rgcy_temp_created_at"`
}

type ResponseMitraAOCoverageRegencyTemps struct {
	Status  int                               `json:"status"`
	Message string                            `json:"message"`
	Data    *[]MitraAOCoverageRegencyTempData `json:"data"`
}

type ResponseMitraAOCoverageRegencyTemp struct {
	Status  int                             `json:"status"`
	Message string                          `json:"message"`
	Data    *MitraAOCoverageRegencyTempData `json:"data"`
}

type ResponseMitraAOCoverageRegencyTempsPend struct {
	Status  int                               `json:"status"`
	Message string                            `json:"message"`
	Data    *[]MitraAOCoverageRegencyTempPend `json:"data"`
}

type ResponseMitraAOCoverageRegencyTempPend struct {
	Status  int                             `json:"status"`
	Message string                          `json:"message"`
	Data    *MitraAOCoverageRegencyTempPend `json:"data"`
}

type ResponseMitraAOCoverageRegencyTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraAOCoverageRegencyTemp) TableName() string {
	return "m_mitra_ao_coverage_regency_temp"
}

func (MitraAOCoverageRegencyTempPend) TableName() string {
	return "m_mitra_ao_coverage_regency_temp"
}

func (MitraAOCoverageRegencyTempPendData) TableName() string {
	return "m_mitra_ao_coverage_regency_temp"
}

func (MitraAOCoverageRegencyTempData) TableName() string {
	return "m_mitra_ao_coverage_regency_temp"
}

func (p *MitraAOCoverageRegencyTemp) Prepare() {
	p.MitraAOCoverageRegencyTempID = nil
	p.MitraAOCoverageRegencyID = p.MitraAOCoverageRegencyID
	p.MitraAOCoverageTempID = p.MitraAOCoverageTempID
	p.RegencyTempID = p.RegencyTempID
	p.MitraAOCoverageRegencyTempReason = p.MitraAOCoverageRegencyTempReason
	p.MitraAOCoverageRegencyTempStatus = p.MitraAOCoverageRegencyTempStatus
	p.MitraAOCoverageRegencyTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraAOCoverageRegencyTempCreatedBy))
	p.MitraAOCoverageRegencyTempCreatedAt = time.Now()
}

func (p *MitraAOCoverageRegencyTempPend) Prepare() {
	p.MitraAOCoverageRegencyTempID = nil
	p.MitraAOCoverageRegencyID = nil
	p.MitraAOCoverageTempID = p.MitraAOCoverageTempID
	p.RegencyTempID = p.RegencyTempID
	p.MitraAOCoverageRegencyTempReason = p.MitraAOCoverageRegencyTempReason
	p.MitraAOCoverageRegencyTempStatus = p.MitraAOCoverageRegencyTempStatus
	p.MitraAOCoverageRegencyTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraAOCoverageRegencyTempCreatedBy))
	p.MitraAOCoverageRegencyTempCreatedAt = time.Now()
}

func (p *MitraAOCoverageRegencyTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	case "insertupdate":
		if p.MitraAOCoverageTempID == 0 {
			return errors.New("Negara Wajib diisi")
		}
		if p.RegencyTempID == 0 {
			return errors.New("Provinsi Wajib diisi")
		}
		if p.MitraAOCoverageRegencyTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.MitraAOCoverageRegencyTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	case "reason":
		if p.MitraAOCoverageRegencyTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.MitraAOCoverageRegencyTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	default:
		if p.MitraAOCoverageRegencyTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil
	}
}

func (p *MitraAOCoverageRegencyTemp) SaveMitraAOCoverageRegencyTemp(db *gorm.DB) (*MitraAOCoverageRegencyTemp, error) {
	var err error
	err = db.Debug().Model(&MitraAOCoverageRegencyTemp{}).Create(&p).Error
	if err != nil {
		return &MitraAOCoverageRegencyTemp{}, err
	}
	return p, nil
}

func (p *MitraAOCoverageRegencyTempPend) SaveMitraAOCoverageRegencyTempPend(db *gorm.DB) (*MitraAOCoverageRegencyTempPend, error) {
	var err error
	err = db.Debug().Model(&MitraAOCoverageRegencyTempPend{}).Create(&p).Error
	if err != nil {
		return &MitraAOCoverageRegencyTempPend{}, err
	}
	return p, nil
}

func (p *MitraAOCoverageRegencyTemp) FindAllMitraAOCoverageRegencyTemp(db *gorm.DB) (*[]MitraAOCoverageRegencyTempPend, error) {
	var err error
	mitraaocoverageregency := []MitraAOCoverageRegencyTempPend{}
	err = db.Debug().Model(&MitraAOCoverageRegencyTempPend{}).Limit(100).
		Select(`m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_id,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_id,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_temp_id,
				m_mitra_ao_coverage_regency_temp.regency_temp_id,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_reason,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_status,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_action_before,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_created_by,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_created_at`).
		Find(&mitraaocoverageregency).Error
	if err != nil {
		return &[]MitraAOCoverageRegencyTempPend{}, err
	}
	return &mitraaocoverageregency, nil
}

func (p *MitraAOCoverageRegencyTemp) FindAllMitraAOCoverageRegencyTempPendingActive(db *gorm.DB) (*[]MitraAOCoverageRegencyTempPend, error) {
	var err error
	mitraaocoverageregency := []MitraAOCoverageRegencyTempPend{}
	err = db.Debug().Model(&MitraAOCoverageRegencyTempPend{}).Limit(100).
		Select(`m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_id,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_id,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_temp_id,
				m_mitra_ao_coverage_regency_temp.regency_temp_id,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_reason,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_status,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_action_before,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_created_by,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_created_at`).
		Where("mitra_ao_cov_rgcy_temp_status = ?", 11).
		Find(&mitraaocoverageregency).Error
	if err != nil {
		return &[]MitraAOCoverageRegencyTempPend{}, err
	}
	return &mitraaocoverageregency, nil
}

func (p *MitraAOCoverageRegencyTemp) FindAllMitraAOCoverageRegencyTempStatus(db *gorm.DB, status uint64) (*[]MitraAOCoverageRegencyTempData, error) {
	var err error
	mitraaocoverageregency := []MitraAOCoverageRegencyTempData{}
	err = db.Debug().Model(&MitraAOCoverageRegencyTempData{}).Limit(100).
		Select(`m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_id,
				m_mitra_ao_coverage_temp.mitra_ao_cov_id as mitra_ao_cov_id,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_name as mitra_ao_cov_name,
				m_regency_temp.regency_id as regency_temp_id,
				m_regency_temp.regency_name as regency_temp_name,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_reason,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_status,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_action_before,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_created_by,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_created_at,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_id,
				m_mitra_ao_coverage.mitra_ao_cov_id,
				m_mitra_ao_coverage.mitra_ao_cov_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_status,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_created_by,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_created_at,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_updated_by,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_updated_at,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_deleted_by,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_deleted_at`).
		Joins("JOIN m_mitra_ao_coverage m_mitra_ao_coverage_temp on m_mitra_ao_coverage_regency_temp.mitra_ao_cov_temp_id=m_mitra_ao_coverage_temp.mitra_ao_cov_temp_id").
		Joins("JOIN m_regency m_regency_temp on m_mitra_ao_coverage_regency_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_mitra_ao_coverage_regency on m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_id=m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_id").
		Joins("JOIN m_mitra_ao_coverage on m_mitra_ao_coverage_regency.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Joins("JOIN m_regency on m_mitra_ao_coverage_regency.regency_id=m_regency.regency_id").
		Where("mitra_ao_cov_rgcy_temp_status = ?", status).
		Find(&mitraaocoverageregency).Error
	if err != nil {
		return &[]MitraAOCoverageRegencyTempData{}, err
	}
	return &mitraaocoverageregency, nil
}

func (p *MitraAOCoverageRegencyTemp) FindMitraAOCoverageRegencyTempDataByID(db *gorm.DB, pid uint64) (*MitraAOCoverageRegencyTemp, error) {
	var err error
	err = db.Debug().Model(&MitraAOCoverageRegencyTemp{}).Where("mitra_ao_cov_rgcy_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraAOCoverageRegencyTemp{}, err
	}
	return p, nil
}

func (p *MitraAOCoverageRegencyTemp) FindMitraAOCoverageRegencyTempByIDPendingActive(db *gorm.DB, pid uint64) (*MitraAOCoverageRegencyTempPend, error) {
	var err error
	mitraaocoverageregency := MitraAOCoverageRegencyTempPend{}
	err = db.Debug().Model(&MitraAOCoverageRegencyTempPend{}).Limit(100).
		Select(`m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_id,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_id,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_temp_id,
				m_mitra_ao_coverage_regency_temp.regency_temp_id,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_reason,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_status,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_action_before,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_created_by,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_created_at`).
		Where("mitra_ao_cov_rgcy_temp_id = ? AND mitra_ao_cov_rgcy_temp_status = ?", pid, 11).
		Find(&mitraaocoverageregency).Error
	if err != nil {
		return &MitraAOCoverageRegencyTempPend{}, err
	}
	return &mitraaocoverageregency, nil
}

func (p *MitraAOCoverageRegencyTemp) FindMitraAOCoverageRegencyTempByID(db *gorm.DB, pid uint64) (*MitraAOCoverageRegencyTempData, error) {
	var err error
	mitraaocoverageregency := MitraAOCoverageRegencyTempData{}
	err = db.Debug().Model(&MitraAOCoverageRegencyTempData{}).Limit(100).
		Select(`m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_id,
				m_mitra_ao_coverage_temp.mitra_ao_cov_id as mitra_ao_cov_id,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_name as mitra_ao_cov_name,
				m_regency_temp.regency_id as regency_temp_id,
				m_regency_temp.regency_name as regency_temp_name,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_reason,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_status,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_action_before,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_created_by,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_created_at,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_id,
				m_mitra_ao_coverage.mitra_ao_cov_id,
				m_mitra_ao_coverage.mitra_ao_cov_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_status,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_created_by,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_created_at,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_updated_by,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_updated_at,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_deleted_by,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_deleted_at`).
		Joins("JOIN m_mitra_ao_coverage m_mitra_ao_coverage_temp on m_mitra_ao_coverage_regency_temp.mitra_ao_cov_temp_id=m_mitra_ao_coverage_temp.mitra_ao_cov_temp_id").
		Joins("JOIN m_regency m_regency_temp on m_mitra_ao_coverage_regency_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_mitra_ao_coverage_regency on m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_id=m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_id").
		Joins("JOIN m_mitra_ao_coverage on m_mitra_ao_coverage_regency.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Joins("JOIN m_regency on m_mitra_ao_coverage_regency.regency_id=m_regency.regency_id").
		Where("mitra_ao_cov_rgcy_temp_id = ?", pid).
		Find(&mitraaocoverageregency).Error
	if err != nil {
		return &MitraAOCoverageRegencyTempData{}, err
	}
	return &mitraaocoverageregency, nil
}

func (p *MitraAOCoverageRegencyTemp) FindMitraAOCoverageRegencyTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraAOCoverageRegencyTempData, error) {
	var err error
	mitraaocoverageregency := MitraAOCoverageRegencyTempData{}
	err = db.Debug().Model(&MitraAOCoverageRegencyTempData{}).Limit(100).
		Select(`m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_id,
				m_mitra_ao_coverage_temp.mitra_ao_cov_id as mitra_ao_cov_id,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_name as mitra_ao_cov_name,
				m_regency_temp.regency_id as regency_temp_id,
				m_regency_temp.regency_name as regency_temp_name,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_reason,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_status,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_action_before,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_created_by,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_created_at,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_id,
				m_mitra_ao_coverage.mitra_ao_cov_id,
				m_mitra_ao_coverage.mitra_ao_cov_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_status,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_created_by,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_created_at,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_updated_by,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_updated_at,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_deleted_by,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_deleted_at`).
		Joins("JOIN m_mitra_ao_coverage m_mitra_ao_coverage_temp on m_mitra_ao_coverage_regency_temp.mitra_ao_cov_temp_id=m_mitra_ao_coverage_temp.mitra_ao_cov_temp_id").
		Joins("JOIN m_regency m_regency_temp on m_mitra_ao_coverage_regency_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_mitra_ao_coverage_regency on m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_id=m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_id").
		Joins("JOIN m_mitra_ao_coverage on m_mitra_ao_coverage_regency.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Joins("JOIN m_regency on m_mitra_ao_coverage_regency.regency_id=m_regency.regency_id").
		Where("mitra_ao_cov_rgcy_temp_id = ? AND mitra_ao_cov_rgcy_temp_status = ?", pid, status).
		Find(&mitraaocoverageregency).Error
	if err != nil {
		return &MitraAOCoverageRegencyTempData{}, err
	}
	return &mitraaocoverageregency, nil
}

func (p *MitraAOCoverageRegencyTemp) UpdateMitraAOCoverageRegencyTemp(db *gorm.DB, pid uint64) (*MitraAOCoverageRegencyTemp, error) {
	var err error
	db = db.Debug().Model(&MitraAOCoverageRegencyTemp{}).Where("mitra_ao_cov_rgcy_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_ao_cov_id":               p.MitraAOCoverageTempID,
			"regency_temp_id":               p.RegencyTempID,
			"mitra_ao_cov_rgcy_temp_reason": p.MitraAOCoverageRegencyTempReason,
		},
	)
	err = db.Debug().Model(&MitraAOCoverageRegencyTemp{}).Where("mitra_ao_cov_rgcy_temp_id = ?", pid).Error
	if err != nil {
		return &MitraAOCoverageRegencyTemp{}, err
	}
	return p, nil
}

func (p *MitraAOCoverageRegencyTemp) UpdateMitraAOCoverageRegencyTempStatus(db *gorm.DB, pid uint64) (*MitraAOCoverageRegencyTemp, error) {
	var err error
	db = db.Debug().Model(&MitraAOCoverageRegencyTemp{}).Where("mitra_ao_cov_rgcy_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_ao_cov_rgcy_temp_reason": p.MitraAOCoverageRegencyTempReason,
			"mitra_ao_cov_rgcy_temp_status": p.MitraAOCoverageRegencyTempStatus,
		},
	)
	err = db.Debug().Model(&MitraAOCoverageRegencyTemp{}).Where("mitra_ao_cov_rgcy_temp_id = ?", pid).Error
	if err != nil {
		return &MitraAOCoverageRegencyTemp{}, err
	}
	return p, nil
}

func (p *MitraAOCoverageRegencyTemp) DeleteMitraAOCoverageRegencyTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraAOCoverageRegencyTemp{}).Where("mitra_ao_cov_rgcy_temp_id = ?", pid).Take(&MitraAOCoverageRegencyTemp{}).Delete(&MitraAOCoverageRegencyTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraAOCoverageRegencyTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

func (p *MitraAOCoverageRegencyTemp) FindMitraAOCoverageRegencyTempPendByMitraAOCoverageTempID(db *gorm.DB, pid uint64) ([]MitraAOCoverageRegencyTempPendData, error) {
	var err error
	mitraaocoverageregency := []MitraAOCoverageRegencyTempPendData{}
	err = db.Debug().Model(&MitraAOCoverageRegencyTempPendData{}).Limit(100).
		Select(`m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_id,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_id,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_temp_id,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_name,
				m_mitra_ao_coverage_regency_temp.regency_temp_id,
				m_regency_temp.regency_name as regency_temp_name,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_reason,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_status,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_action_before,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_created_by,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_created_at`).
		Joins("JOIN m_regency m_regency_temp on m_mitra_ao_coverage_regency_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_mitra_ao_coverage_temp on m_mitra_ao_coverage_regency_temp.mitra_ao_cov_temp_id=m_mitra_ao_coverage_temp.mitra_ao_cov_temp_id").
		Where("m_mitra_ao_coverage_regency_temp.mitra_ao_cov_temp_id = ?", pid).
		Find(&mitraaocoverageregency).Error
	if err != nil {
		return []MitraAOCoverageRegencyTempPendData{}, err
	}
	return mitraaocoverageregency, nil
}

func (p *MitraAOCoverageRegencyTemp) FindMitraAOCoverageRegencyTempPendStatusByMitraAOCoverageTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraAOCoverageRegencyTempPendData, error) {
	var err error
	mitraaocoverageregency := []MitraAOCoverageRegencyTempPendData{}
	err = db.Debug().Model(&MitraAOCoverageRegencyTempPendData{}).Limit(100).
		Select(`m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_id,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_id,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_temp_id,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_name,
				m_mitra_ao_coverage_regency_temp.regency_temp_id,
				m_regency_temp.regency_name as regency_temp_name,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_reason,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_status,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_action_before,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_created_by,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_created_at`).
		Joins("JOIN m_regency m_regency_temp on m_mitra_ao_coverage_regency_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_mitra_ao_coverage_temp on m_mitra_ao_coverage_regency_temp.mitra_ao_cov_temp_id=m_mitra_ao_coverage_temp.mitra_ao_cov_temp_id").
		Where("m_mitra_ao_coverage_regency_temp.mitra_ao_cov_temp_id = ? AND mitra_ao_cov_rgcy_temp_status = ?", pid, status).
		Find(&mitraaocoverageregency).Error
	if err != nil {
		return []MitraAOCoverageRegencyTempPendData{}, err
	}
	return mitraaocoverageregency, nil
}

func (p *MitraAOCoverageRegencyTemp) FindMitraAOCoverageRegencyTempByMitraAOCoverageTempID(db *gorm.DB, pid uint64) ([]MitraAOCoverageRegencyTempData, error) {
	var err error
	mitraaocoverageregency := []MitraAOCoverageRegencyTempData{}
	err = db.Debug().Model(&MitraAOCoverageRegencyTempData{}).Limit(100).
		Select(`m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_id,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_id,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_temp_id,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_name,
				m_mitra_ao_coverage_regency_temp.regency_temp_id,
				m_regency_temp.regency_name as regency_temp_name,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_reason,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_status,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_action_before,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_created_by,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_created_at`).
		Joins("JOIN m_regency m_regency_temp on m_mitra_ao_coverage_regency_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_mitra_ao_coverage_temp on m_mitra_ao_coverage_regency_temp.mitra_ao_cov_temp_id=m_mitra_ao_coverage_temp.mitra_ao_cov_temp_id").
		Where("m_mitra_ao_coverage_regency_temp.mitra_ao_cov_temp_id = ?", pid).
		Find(&mitraaocoverageregency).Error
	if err != nil {
		return []MitraAOCoverageRegencyTempData{}, err
	}
	return mitraaocoverageregency, nil
}

func (p *MitraAOCoverageRegencyTemp) FindMitraAOCoverageRegencyTempStatusByMitraAOCoverageTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraAOCoverageRegencyTempData, error) {
	var err error
	mitraaocoverageregency := []MitraAOCoverageRegencyTempData{}
	err = db.Debug().Model(&MitraAOCoverageRegencyTempData{}).Limit(100).
		Select(`m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_id,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_id,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_temp_id,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_name,
				m_mitra_ao_coverage_regency_temp.regency_temp_id,
				m_regency_temp.regency_name as regency_temp_name,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_reason,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_status,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_action_before,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_created_by,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_created_at`).
		Joins("JOIN m_regency m_regency_temp on m_mitra_ao_coverage_regency_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_mitra_ao_coverage_temp on m_mitra_ao_coverage_regency_temp.mitra_ao_cov_temp_id=m_mitra_ao_coverage_temp.mitra_ao_cov_temp_id").
		Where("m_mitra_ao_coverage_regency_temp.mitra_ao_cov_temp_id = ? AND mitra_ao_cov_rgcy_temp_status = ?", pid, status).
		Find(&mitraaocoverageregency).Error
	if err != nil {
		return []MitraAOCoverageRegencyTempData{}, err
	}
	return mitraaocoverageregency, nil
}

func (p *MitraAOCoverageRegencyTemp) FindMitraAOCoverageRegencyTempDataStatusByMitraAOCoverageTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraAOCoverageRegencyTempData, error) {
	var err error
	mitraaocoverageregency := []MitraAOCoverageRegencyTempData{}
	err = db.Debug().Model(&MitraAOCoverageRegencyTempData{}).Limit(100).
		Select(`m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_id,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_id,
				m_mitra_ao_coverage_temp.mitra_ao_cov_id,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_name,
				m_mitra_ao_coverage_regency_temp.regency_temp_id,
				m_regency_temp.regency_name as regency_temp_name,
				m_mitra_ao_coverage.mitra_ao_cov_id,
				m_mitra_ao_coverage.mitra_ao_cov_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_reason,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_status,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_action_before,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_created_by,
				m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_temp_created_at`).
		Joins("JOIN m_regency m_regency_temp on m_mitra_ao_coverage_regency_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_mitra_ao_coverage_temp on m_mitra_ao_coverage_regency_temp.mitra_ao_cov_temp_id=m_mitra_ao_coverage_temp.mitra_ao_cov_temp_id").
		Joins("JOIN m_mitra_ao_coverage_regency on m_mitra_ao_coverage_regency_temp.mitra_ao_cov_rgcy_id=m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_id").
		Joins("JOIN m_regency on m_mitra_ao_coverage_regency.regency_id=m_regency.regency_id").
		Joins("JOIN m_mitra_ao_coverage on m_mitra_ao_coverage_regency.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Where("m_mitra_ao_coverage_regency_temp.mitra_ao_cov_temp_id = ? AND mitra_ao_cov_rgcy_temp_status = ?", pid, status).
		Find(&mitraaocoverageregency).Error
	if err != nil {
		return []MitraAOCoverageRegencyTempData{}, err
	}
	return mitraaocoverageregency, nil
}
