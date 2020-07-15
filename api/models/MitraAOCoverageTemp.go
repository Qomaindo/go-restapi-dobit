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

type MitraAOCoverageTemp struct {
	MitraAOCoverageTempID           uint64    `gorm:"column:mitra_ao_cov_temp_id;primary_key;" json:"mitra_ao_cov_temp_id"`
	MitraAOCoverageID               uint64    `gorm:"column:mitra_ao_cov_id" json:"mitra_ao_cov_id"`
	MitraTempID                     uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	MitraAOCoverageTempCode         string    `gorm:"column:mitra_ao_cov_temp_code" json:"mitra_ao_cov_temp_code"`
	MitraAOCoverageTempName         string    `gorm:"column:mitra_ao_cov_temp_name" json:"mitra_ao_cov_temp_name"`
	MitraAOCoverageTempReason       string    `gorm:"column:mitra_ao_cov_temp_reason" json:"mitra_ao_cov_temp_reason"`
	MitraAOCoverageTempStatus       uint64    `gorm:"column:mitra_ao_cov_temp_status;size:2" json:"mitra_ao_cov_temp_status"`
	MitraAOCoverageTempActionBefore uint64    `gorm:"column:mitra_ao_cov_temp_action_before;size:2" json:"mitra_ao_cov_temp_action_before"`
	MitraAOCoverageTempCreatedBy    string    `gorm:"column:mitra_ao_cov_temp_created_by;size:125" json:"mitra_ao_cov_temp_created_by"`
	MitraAOCoverageTempCreatedAt    time.Time `gorm:"column:mitra_ao_cov_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_ao_cov_temp_created_at"`

	MitraAOCoverageRegencyTemp []MitraAOCoverageRegencyTempPendData `json:"mitra_ao_cov_temp_regency"`
}

type MitraAOCoverageTempPend struct {
	MitraAOCoverageTempID           uint64    `gorm:"column:mitra_ao_cov_temp_id;primary_key;" json:"mitra_ao_cov_temp_id"`
	MitraAOCoverageID               *uint64   `gorm:"column:mitra_ao_cov_id" json:"mitra_ao_cov_id"`
	MitraTempID                     uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	MitraAOCoverageTempCode         string    `gorm:"column:mitra_ao_cov_temp_code" json:"mitra_ao_cov_temp_code"`
	MitraAOCoverageTempName         string    `gorm:"column:mitra_ao_cov_temp_name" json:"mitra_ao_cov_temp_name"`
	MitraAOCoverageTempReason       string    `gorm:"column:mitra_ao_cov_temp_reason" json:"mitra_ao_cov_temp_reason"`
	MitraAOCoverageTempStatus       uint64    `gorm:"column:mitra_ao_cov_temp_status;size:2" json:"mitra_ao_cov_temp_status"`
	MitraAOCoverageTempActionBefore uint64    `gorm:"column:mitra_ao_cov_temp_action_before;size:2" json:"mitra_ao_cov_temp_action_before"`
	MitraAOCoverageTempCreatedBy    string    `gorm:"column:mitra_ao_cov_temp_created_by;size:125" json:"mitra_ao_cov_temp_created_by"`
	MitraAOCoverageTempCreatedAt    time.Time `gorm:"column:mitra_ao_cov_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_ao_cov_temp_created_at"`

	MitraAOCoverageRegencyTemp []MitraAOCoverageRegencyTempPendData `json:"mitra_ao_cov_temp_regency"`
}

type MitraAOCoverageTempPendData struct {
	MitraAOCoverageTempID           uint64    `gorm:"column:mitra_ao_cov_temp_id;primary_key;" json:"mitra_ao_cov_temp_id"`
	MitraAOCoverageID               uint64    `gorm:"column:mitra_ao_cov_id" json:"mitra_ao_cov_id"`
	MitraTempID                     uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	MitraAOCoverageTempCode         string    `gorm:"column:mitra_ao_cov_temp_code" json:"mitra_ao_cov_temp_code"`
	MitraAOCoverageTempName         string    `gorm:"column:mitra_ao_cov_temp_name" json:"mitra_ao_cov_temp_name"`
	MitraAOCoverageTempReason       string    `gorm:"column:mitra_ao_cov_temp_reason" json:"mitra_ao_cov_temp_reason"`
	MitraAOCoverageTempStatus       uint64    `gorm:"column:mitra_ao_cov_temp_status;size:2" json:"mitra_ao_cov_temp_status"`
	MitraAOCoverageTempActionBefore uint64    `gorm:"column:mitra_ao_cov_temp_action_before;size:2" json:"mitra_ao_cov_temp_action_before"`
	MitraAOCoverageTempCreatedBy    string    `gorm:"column:mitra_ao_cov_temp_created_by;size:125" json:"mitra_ao_cov_temp_created_by"`
	MitraAOCoverageTempCreatedAt    time.Time `gorm:"column:mitra_ao_cov_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_ao_cov_temp_created_at"`
}

type MitraAOCoverageTempData struct {
	MitraAOCoverageTempID           uint64    `gorm:"column:mitra_ao_cov_temp_id" json:"mitra_ao_cov_temp_id"`
	MitraTempID                     uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	MitraAOCoverageTempCode         string    `gorm:"column:mitra_ao_cov_temp_code" json:"mitra_ao_cov_temp_code"`
	MitraAOCoverageTempName         string    `gorm:"column:mitra_ao_cov_temp_name" json:"mitra_ao_cov_temp_name"`
	MitraAOCoverageTempReason       string    `gorm:"column:mitra_ao_cov_temp_reason" json:"mitra_ao_cov_temp_reason"`
	MitraAOCoverageTempStatus       uint64    `gorm:"column:mitra_ao_cov_temp_status;size:2" json:"mitra_ao_cov_temp_status"`
	MitraAOCoverageTempActionBefore uint64    `gorm:"column:mitra_ao_cov_temp_action_before;size:2" json:"mitra_ao_cov_temp_action_before"`
	MitraAOCoverageTempCreatedBy    string    `gorm:"column:mitra_ao_cov_temp_created_by;size:125" json:"mitra_ao_cov_temp_created_by"`
	MitraAOCoverageTempCreatedAt    time.Time `gorm:"column:mitra_ao_cov_temp_created_at" json:"mitra_ao_cov_temp_created_at"`

	MitraAOCoverageRegencyTemp []MitraAOCoverageRegencyTempData `json:"mitra_ao_cov_temp_regency"`

	MitraAOCoverageID        uint64    `gorm:"column:mitra_ao_cov_id" json:"mitra_ao_cov_id"`
	MitraAOCoverageCode      string    `gorm:"column:mitra_ao_cov_code" json:"mitra_ao_cov_code"`
	MitraAOCoverageName      string    `gorm:"column:mitra_ao_cov_name" json:"mitra_ao_cov_name"`
	MitraAOCoverageStatus    uint64    `gorm:"column:mitra_ao_cov_status" json:"mitra_ao_cov_status"`
	MitraAOCoverageCreatedBy string    `gorm:"column:mitra_ao_cov_created_by" json:"mitra_ao_cov_created_by"`
	MitraAOCoverageCreatedAt time.Time `gorm:"column:mitra_ao_cov_created_at" json:"mitra_ao_cov_created_at"`
	MitraAOCoverageUpdatedBy string    `gorm:"column:mitra_ao_cov_updated_by" json:"mitra_ao_cov_updated_by"`
	MitraAOCoverageUpdatedAt time.Time `gorm:"column:mitra_ao_cov_updated_at" json:"mitra_ao_cov_updated_at"`
	MitraAOCoverageDeletedBy string    `gorm:"column:mitra_ao_cov_deleted_by" json:"mitra_ao_cov_deleted_by"`
	MitraAOCoverageDeletedAt time.Time `gorm:"column:mitra_ao_cov_deleted_at" json:"mitra_ao_cov_deleted_at"`

	MitraAOCoverageRegency []MitraAOCoverageRegencyData `json:"mitra_ao_cov_regency"`
}

type ResponseMitraAOCoverageTemps struct {
	Status  int                        `json:"status"`
	Message string                     `json:"message"`
	Data    *[]MitraAOCoverageTempData `json:"data"`
}

type ResponseMitraAOCoverageTemp struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    *MitraAOCoverageTempData `json:"data"`
}

type ResponseMitraAOCoverageTempsPend struct {
	Status  int                        `json:"status"`
	Message string                     `json:"message"`
	Data    *[]MitraAOCoverageTempPend `json:"data"`
}

type ResponseMitraAOCoverageTempPend struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    *MitraAOCoverageTempPend `json:"data"`
}

type ResponseMitraAOCoverageTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraAOCoverageTemp) TableName() string {
	return "m_mitra_ao_coverage_temp"
}

func (MitraAOCoverageTempPend) TableName() string {
	return "m_mitra_ao_coverage_temp"
}

func (MitraAOCoverageTempPendData) TableName() string {
	return "m_mitra_ao_coverage_temp"
}

func (MitraAOCoverageTempData) TableName() string {
	return "m_mitra_ao_coverage_temp"
}

func (p *MitraAOCoverageTemp) Prepare() {
	p.MitraAOCoverageID = p.MitraAOCoverageID
	p.MitraAOCoverageTempCode = p.MitraAOCoverageTempCode
	p.MitraAOCoverageTempName = p.MitraAOCoverageTempName
	p.MitraAOCoverageTempReason = p.MitraAOCoverageTempReason
	p.MitraAOCoverageTempStatus = p.MitraAOCoverageTempStatus
	p.MitraAOCoverageTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraAOCoverageTempCreatedBy))
	p.MitraAOCoverageTempCreatedAt = time.Now()
}

func (p *MitraAOCoverageTempPend) Prepare() {
	p.MitraAOCoverageID = nil
	p.MitraAOCoverageTempCode = p.MitraAOCoverageTempCode
	p.MitraAOCoverageTempName = p.MitraAOCoverageTempName
	p.MitraAOCoverageTempReason = p.MitraAOCoverageTempReason
	p.MitraAOCoverageTempStatus = p.MitraAOCoverageTempStatus
	p.MitraAOCoverageTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraAOCoverageTempCreatedBy))
	p.MitraAOCoverageTempCreatedAt = time.Now()
}

func (p *MitraAOCoverageTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	case "insertupdate":
		if p.MitraAOCoverageTempName == "" {
			return errors.New("Negara Wajib diisi")
		}
		if p.MitraAOCoverageTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.MitraAOCoverageTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	case "reason":
		if p.MitraAOCoverageTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.MitraAOCoverageTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	default:
		if p.MitraAOCoverageTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil
	}
}

func (p *MitraAOCoverageTemp) SaveMitraAOCoverageTemp(db *gorm.DB) (*MitraAOCoverageTemp, error) {
	var err error
	err = db.Debug().Model(&MitraAOCoverageTemp{}).Create(&p).Error
	if err != nil {
		return &MitraAOCoverageTemp{}, err
	}
	return p, nil
}

func (p *MitraAOCoverageTempPend) SaveMitraAOCoverageTempPend(db *gorm.DB) (*MitraAOCoverageTempPend, error) {
	var err error
	err = db.Debug().Model(&MitraAOCoverageTempPend{}).Create(&p).Error
	if err != nil {
		return &MitraAOCoverageTempPend{}, err
	}
	return p, nil
}

func (p *MitraAOCoverageTemp) FindAllMitraAOCoverageTemp(db *gorm.DB) (*[]MitraAOCoverageTempPend, error) {
	var err error
	mitraaocoverage := []MitraAOCoverageTempPend{}
	err = db.Debug().Model(&MitraAOCoverageTempPend{}).Limit(100).
		Select(`m_mitra_ao_coverage_temp.mitra_ao_cov_temp_id,
				m_mitra_ao_coverage_temp.mitra_ao_cov_id,
				m_mitra_ao_coverage_temp.mitra_temp_id,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_code,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_name,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_reason,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_status,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_action_before,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_created_by,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_created_at`).
		Find(&mitraaocoverage).Error
	if err != nil {
		return &[]MitraAOCoverageTempPend{}, err
	}
	return &mitraaocoverage, nil
}

func (p *MitraAOCoverageTemp) FindAllMitraAOCoverageTempPendingActive(db *gorm.DB) (*[]MitraAOCoverageTempPend, error) {
	var err error
	mitraaocoverage := []MitraAOCoverageTempPend{}
	err = db.Debug().Model(&MitraAOCoverageTempPend{}).Limit(100).
		Select(`m_mitra_ao_coverage_temp.mitra_ao_cov_temp_id,
				m_mitra_ao_coverage_temp.mitra_ao_cov_id,
				m_mitra_ao_coverage_temp.mitra_temp_id,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_code,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_name,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_reason,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_status,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_action_before,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_created_by,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_created_at`).
		Where("mitra_ao_cov_temp_status = ?", 11).
		Find(&mitraaocoverage).Error
	if err != nil {
		return &[]MitraAOCoverageTempPend{}, err
	}
	return &mitraaocoverage, nil
}

func (p *MitraAOCoverageTemp) FindAllMitraAOCoverageTempStatus(db *gorm.DB, status uint64) (*[]MitraAOCoverageTempData, error) {
	var err error
	mitraaocoverage := []MitraAOCoverageTempData{}
	err = db.Debug().Model(&MitraAOCoverageTempData{}).Limit(100).
		Select(`m_mitra_ao_coverage_temp.mitra_ao_cov_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_code,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_name,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_reason,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_status,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_action_before,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_created_by,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_created_at,
				m_mitra_ao_coverage.mitra_ao_cov_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_mitra_ao_coverage.mitra_ao_cov_code,
				m_mitra_ao_coverage.mitra_ao_cov_name,
				m_mitra_ao_coverage.mitra_ao_cov_status,
				m_mitra_ao_coverage.mitra_ao_cov_created_by,
				m_mitra_ao_coverage.mitra_ao_cov_created_at,
				m_mitra_ao_coverage.mitra_ao_cov_updated_by,
				m_mitra_ao_coverage.mitra_ao_cov_updated_at,
				m_mitra_ao_coverage.mitra_ao_cov_deleted_by,
				m_mitra_ao_coverage.mitra_ao_cov_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_ao_coverage_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_ao_coverage on m_mitra_ao_coverage_temp.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Joins("JOIN m_mitra on m_mitra_ao_coverage.mitra_id=m_mitra.mitra_id").
		Where("mitra_ao_cov_temp_status = ?", status).
		Find(&mitraaocoverage).Error
	if err != nil {
		return &[]MitraAOCoverageTempData{}, err
	}
	return &mitraaocoverage, nil
}

func (p *MitraAOCoverageTemp) FindMitraAOCoverageTempDataByID(db *gorm.DB, pid uint64) (*MitraAOCoverageTemp, error) {
	var err error
	err = db.Debug().Model(&MitraAOCoverageTemp{}).Where("mitra_ao_cov_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraAOCoverageTemp{}, err
	}
	return p, nil
}

func (p *MitraAOCoverageTemp) FindMitraAOCoverageTempByIDPendingActive(db *gorm.DB, pid uint64) (*MitraAOCoverageTempPend, error) {
	var err error
	mitraaocoverage := MitraAOCoverageTempPend{}
	err = db.Debug().Model(&MitraAOCoverageTempPend{}).Limit(100).
		Select(`m_mitra_ao_coverage_temp.mitra_ao_cov_temp_id,
				m_mitra_ao_coverage_temp.mitra_ao_cov_id,
				m_mitra_ao_coverage_temp.mitra_temp_id,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_code,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_name,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_reason,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_status,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_action_before,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_created_by,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_created_at`).
		Where("mitra_ao_cov_temp_id = ? AND mitra_ao_cov_temp_status = ?", pid, 11).
		Find(&mitraaocoverage).Error
	if err != nil {
		return &MitraAOCoverageTempPend{}, err
	}
	return &mitraaocoverage, nil
}

func (p *MitraAOCoverageTemp) FindMitraAOCoverageTempByID(db *gorm.DB, pid uint64) (*MitraAOCoverageTempData, error) {
	var err error
	mitraaocoverage := MitraAOCoverageTempData{}
	err = db.Debug().Model(&MitraAOCoverageTempData{}).Limit(100).
		Select(`m_mitra_ao_coverage_temp.mitra_ao_cov_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_code,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_name,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_reason,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_status,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_action_before,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_created_by,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_created_at,
				m_mitra_ao_coverage.mitra_ao_cov_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_mitra_ao_coverage.mitra_ao_cov_code,
				m_mitra_ao_coverage.mitra_ao_cov_name,
				m_mitra_ao_coverage.mitra_ao_cov_status,
				m_mitra_ao_coverage.mitra_ao_cov_created_by,
				m_mitra_ao_coverage.mitra_ao_cov_created_at,
				m_mitra_ao_coverage.mitra_ao_cov_updated_by,
				m_mitra_ao_coverage.mitra_ao_cov_updated_at,
				m_mitra_ao_coverage.mitra_ao_cov_deleted_by,
				m_mitra_ao_coverage.mitra_ao_cov_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_ao_coverage_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_ao_coverage on m_mitra_ao_coverage_temp.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Joins("JOIN m_mitra on m_mitra_ao_coverage.mitra_id=m_mitra.mitra_id").
		Where("mitra_ao_cov_temp_id = ?", pid).
		Find(&mitraaocoverage).Error
	if err != nil {
		return &MitraAOCoverageTempData{}, err
	}
	return &mitraaocoverage, nil
}

func (p *MitraAOCoverageTemp) FindMitraAOCoverageTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraAOCoverageTempData, error) {
	var err error
	mitraaocoverage := MitraAOCoverageTempData{}
	err = db.Debug().Model(&MitraAOCoverageTempData{}).Limit(100).
		Select(`m_mitra_ao_coverage_temp.mitra_ao_cov_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_code,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_name,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_reason,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_status,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_action_before,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_created_by,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_created_at,
				m_mitra_ao_coverage.mitra_ao_cov_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_mitra_ao_coverage.mitra_ao_cov_code,
				m_mitra_ao_coverage.mitra_ao_cov_name,
				m_mitra_ao_coverage.mitra_ao_cov_status,
				m_mitra_ao_coverage.mitra_ao_cov_created_by,
				m_mitra_ao_coverage.mitra_ao_cov_created_at,
				m_mitra_ao_coverage.mitra_ao_cov_updated_by,
				m_mitra_ao_coverage.mitra_ao_cov_updated_at,
				m_mitra_ao_coverage.mitra_ao_cov_deleted_by,
				m_mitra_ao_coverage.mitra_ao_cov_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_ao_coverage_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_ao_coverage on m_mitra_ao_coverage_temp.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Joins("JOIN m_mitra on m_mitra_ao_coverage.mitra_id=m_mitra.mitra_id").
		Where("mitra_ao_cov_temp_id = ? AND mitra_ao_cov_temp_status = ?", pid, status).
		Find(&mitraaocoverage).Error
	if err != nil {
		return &MitraAOCoverageTempData{}, err
	}
	return &mitraaocoverage, nil
}

func (p *MitraAOCoverageTemp) UpdateMitraAOCoverageTemp(db *gorm.DB, pid uint64) (*MitraAOCoverageTemp, error) {
	var err error
	db = db.Debug().Model(&MitraAOCoverageTemp{}).Where("mitra_ao_cov_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_ao_cov_temp_code":   p.MitraAOCoverageTempCode,
			"mitra_ao_cov_temp_name":   p.MitraAOCoverageTempName,
			"mitra_ao_cov_temp_reason": p.MitraAOCoverageTempReason,
		},
	)
	err = db.Debug().Model(&MitraAOCoverageTemp{}).Where("mitra_ao_cov_temp_id = ?", pid).Error
	if err != nil {
		return &MitraAOCoverageTemp{}, err
	}
	return p, nil
}

func (p *MitraAOCoverageTemp) UpdateMitraAOCoverageTempStatus(db *gorm.DB, pid uint64) (*MitraAOCoverageTemp, error) {
	var err error
	db = db.Debug().Model(&MitraAOCoverageTemp{}).Where("mitra_ao_cov_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_ao_cov_temp_reason": p.MitraAOCoverageTempReason,
			"mitra_ao_cov_temp_status": p.MitraAOCoverageTempStatus,
		},
	)
	err = db.Debug().Model(&MitraAOCoverageTemp{}).Where("mitra_ao_cov_temp_id = ?", pid).Error
	if err != nil {
		return &MitraAOCoverageTemp{}, err
	}
	return p, nil
}

func (p *MitraAOCoverageTemp) DeleteMitraAOCoverageTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraAOCoverageTemp{}).Where("mitra_ao_cov_temp_id = ?", pid).Take(&MitraAOCoverageTemp{}).Delete(&MitraAOCoverageTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraAOCoverageTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//BY MITRA ID
//===================================================================================================================================

func (p *MitraAOCoverageTemp) FindAllMitraAOCoverageTempMitra(db *gorm.DB, mitra uint64) (*[]MitraAOCoverageTempPend, error) {
	var err error
	mitraaocoverage := []MitraAOCoverageTempPend{}
	err = db.Debug().Model(&MitraAOCoverageTempPend{}).Limit(100).
		Select(`m_mitra_ao_coverage_temp.mitra_ao_cov_temp_id,
				m_mitra_ao_coverage_temp.mitra_ao_cov_id,
				m_mitra_ao_coverage_temp.mitra_temp_id,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_code,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_name,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_reason,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_status,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_action_before,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_created_by,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_created_at`).
		Where("m_mitra_ao_coverage_temp.mitra_temp_id = ?", mitra).
		Find(&mitraaocoverage).Error
	if err != nil {
		return &[]MitraAOCoverageTempPend{}, err
	}
	return &mitraaocoverage, nil
}

func (p *MitraAOCoverageTemp) FindAllMitraAOCoverageTempPendingActiveMitra(db *gorm.DB, mitra uint64) (*[]MitraAOCoverageTempPend, error) {
	var err error
	mitraaocoverage := []MitraAOCoverageTempPend{}
	err = db.Debug().Model(&MitraAOCoverageTempPend{}).Limit(100).
		Select(`m_mitra_ao_coverage_temp.mitra_ao_cov_temp_id,
				m_mitra_ao_coverage_temp.mitra_ao_cov_id,
				m_mitra_ao_coverage_temp.mitra_temp_id,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_code,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_name,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_reason,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_status,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_action_before,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_created_by,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_created_at`).
		Where("mitra_ao_cov_temp_status = ? AND m_mitra_ao_coverage_temp.mitra_temp_id = ?", 11, mitra).
		Find(&mitraaocoverage).Error
	if err != nil {
		return &[]MitraAOCoverageTempPend{}, err
	}
	return &mitraaocoverage, nil
}

func (p *MitraAOCoverageTemp) FindAllMitraAOCoverageTempStatusMitra(db *gorm.DB, status uint64, mitra uint64) (*[]MitraAOCoverageTempData, error) {
	var err error
	mitraaocoverage := []MitraAOCoverageTempData{}
	err = db.Debug().Model(&MitraAOCoverageTempData{}).Limit(100).
		Select(`m_mitra_ao_coverage_temp.mitra_ao_cov_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_code,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_name,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_reason,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_status,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_action_before,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_created_by,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_created_at,
				m_mitra_ao_coverage.mitra_ao_cov_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_mitra_ao_coverage.mitra_ao_cov_code,
				m_mitra_ao_coverage.mitra_ao_cov_name,
				m_mitra_ao_coverage.mitra_ao_cov_status,
				m_mitra_ao_coverage.mitra_ao_cov_created_by,
				m_mitra_ao_coverage.mitra_ao_cov_created_at,
				m_mitra_ao_coverage.mitra_ao_cov_updated_by,
				m_mitra_ao_coverage.mitra_ao_cov_updated_at,
				m_mitra_ao_coverage.mitra_ao_cov_deleted_by,
				m_mitra_ao_coverage.mitra_ao_cov_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_ao_coverage_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_ao_coverage on m_mitra_ao_coverage_temp.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Joins("JOIN m_mitra on m_mitra_ao_coverage.mitra_id=m_mitra.mitra_id").
		Where("mitra_ao_cov_temp_status = ? AND m_mitra_ao_coverage_temp.mitra_temp_id = ?", status, mitra).
		Find(&mitraaocoverage).Error
	if err != nil {
		return &[]MitraAOCoverageTempData{}, err
	}
	return &mitraaocoverage, nil
}

func (p *MitraAOCoverageTemp) FindMitraAOCoverageTempByIDPendingActiveMitra(db *gorm.DB, pid uint64, mitra uint64) (*MitraAOCoverageTempPend, error) {
	var err error
	mitraaocoverage := MitraAOCoverageTempPend{}
	err = db.Debug().Model(&MitraAOCoverageTempPend{}).Limit(100).
		Select(`m_mitra_ao_coverage_temp.mitra_ao_cov_temp_id,
				m_mitra_ao_coverage_temp.mitra_ao_cov_id,
				m_mitra_ao_coverage_temp.mitra_temp_id,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_code,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_name,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_reason,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_status,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_action_before,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_created_by,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_created_at`).
		Where("mitra_ao_cov_temp_id = ? AND mitra_ao_cov_temp_status = ? AND m_mitra_ao_coverage_temp.mitra_temp_id = ?", pid, 11, mitra).
		Find(&mitraaocoverage).Error
	if err != nil {
		return &MitraAOCoverageTempPend{}, err
	}
	return &mitraaocoverage, nil
}

func (p *MitraAOCoverageTemp) FindMitraAOCoverageTempByIDMitra(db *gorm.DB, pid uint64, mitra uint64) (*MitraAOCoverageTempData, error) {
	var err error
	mitraaocoverage := MitraAOCoverageTempData{}
	err = db.Debug().Model(&MitraAOCoverageTempData{}).Limit(100).
		Select(`m_mitra_ao_coverage_temp.mitra_ao_cov_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_code,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_name,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_reason,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_status,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_action_before,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_created_by,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_created_at,
				m_mitra_ao_coverage.mitra_ao_cov_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_mitra_ao_coverage.mitra_ao_cov_code,
				m_mitra_ao_coverage.mitra_ao_cov_name,
				m_mitra_ao_coverage.mitra_ao_cov_status,
				m_mitra_ao_coverage.mitra_ao_cov_created_by,
				m_mitra_ao_coverage.mitra_ao_cov_created_at,
				m_mitra_ao_coverage.mitra_ao_cov_updated_by,
				m_mitra_ao_coverage.mitra_ao_cov_updated_at,
				m_mitra_ao_coverage.mitra_ao_cov_deleted_by,
				m_mitra_ao_coverage.mitra_ao_cov_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_ao_coverage_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_ao_coverage on m_mitra_ao_coverage_temp.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Joins("JOIN m_mitra on m_mitra_ao_coverage.mitra_id=m_mitra.mitra_id").
		Where("mitra_ao_cov_temp_id = ? AND m_mitra_ao_coverage_temp.mitra_temp_id = ?", pid, mitra).
		Find(&mitraaocoverage).Error
	if err != nil {
		return &MitraAOCoverageTempData{}, err
	}
	return &mitraaocoverage, nil
}

func (p *MitraAOCoverageTemp) FindMitraAOCoverageTempStatusByIDMitra(db *gorm.DB, pid uint64, status uint64, mitra uint64) (*MitraAOCoverageTempData, error) {
	var err error
	mitraaocoverage := MitraAOCoverageTempData{}
	err = db.Debug().Model(&MitraAOCoverageTempData{}).Limit(100).
		Select(`m_mitra_ao_coverage_temp.mitra_ao_cov_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_code,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_name,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_reason,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_status,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_action_before,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_created_by,
				m_mitra_ao_coverage_temp.mitra_ao_cov_temp_created_at,
				m_mitra_ao_coverage.mitra_ao_cov_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_mitra_ao_coverage.mitra_ao_cov_code,
				m_mitra_ao_coverage.mitra_ao_cov_name,
				m_mitra_ao_coverage.mitra_ao_cov_status,
				m_mitra_ao_coverage.mitra_ao_cov_created_by,
				m_mitra_ao_coverage.mitra_ao_cov_created_at,
				m_mitra_ao_coverage.mitra_ao_cov_updated_by,
				m_mitra_ao_coverage.mitra_ao_cov_updated_at,
				m_mitra_ao_coverage.mitra_ao_cov_deleted_by,
				m_mitra_ao_coverage.mitra_ao_cov_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_ao_coverage_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_ao_coverage on m_mitra_ao_coverage_temp.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Joins("JOIN m_mitra on m_mitra_ao_coverage.mitra_id=m_mitra.mitra_id").
		Where("mitra_ao_cov_temp_id = ? AND mitra_ao_cov_temp_status = ? AND m_mitra_ao_coverage_temp.mitra_temp_id = ?", pid, status, mitra).
		Find(&mitraaocoverage).Error
	if err != nil {
		return &MitraAOCoverageTempData{}, err
	}
	return &mitraaocoverage, nil
}

//ADDITIONAL
//====================================================================================================================================
