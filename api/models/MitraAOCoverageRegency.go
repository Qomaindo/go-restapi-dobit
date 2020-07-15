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

type MitraAOCoverageRegency struct {
	MitraAOCoverageRegencyID        *uint64    `gorm:"column:mitra_ao_cov_rgcy_id;primary_key;" json:"mitra_ao_cov_rgcy_id"`
	MitraAOCoverageID               *uint64    `gorm:"column:mitra_ao_cov_id" json:"mitra_ao_cov_id"`
	RegencyID                       uint64     `gorm:"column:regency_id" json:"regency_id"`
	MitraAOCoverageRegencyStatus    uint64     `gorm:"column:mitra_ao_cov_rgcy_status;size:2" json:"mitra_ao_cov_rgcy_status"`
	MitraAOCoverageRegencyCreatedBy string     `gorm:"column:mitra_ao_cov_rgcy_created_by;size:125" json:"mitra_ao_cov_rgcy_created_by"`
	MitraAOCoverageRegencyCreatedAt time.Time  `gorm:"column:mitra_ao_cov_rgcy_created_at;default:CURRENT_TIMESTAMP" json:"mitra_ao_cov_rgcy_created_at"`
	MitraAOCoverageRegencyUpdatedBy string     `gorm:"column:mitra_ao_cov_rgcy_updated_by;size:125" json:"mitra_ao_cov_rgcy_updated_by"`
	MitraAOCoverageRegencyUpdatedAt *time.Time `gorm:"column:mitra_ao_cov_rgcy_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_ao_cov_rgcy_created_at"`
	MitraAOCoverageRegencyDeletedBy string     `gorm:"column:mitra_ao_cov_rgcy_deleted_by;size:125" json:"mitra_ao_cov_rgcy_deleted_by"`
	MitraAOCoverageRegencyDeletedAt *time.Time `gorm:"column:mitra_ao_cov_rgcy_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_ao_cov_rgcy_deleted_at"`
}

type MitraAOCoverageRegencyData struct {
	MitraAOCoverageRegencyID        uint64     `gorm:"column:mitra_ao_cov_rgcy_id" json:"mitra_ao_cov_rgcy_id"`
	MitraAOCoverageID               uint64     `gorm:"column:mitra_ao_cov_id" json:"mitra_ao_cov_id"`
	MitraAOCoverageName             string     `gorm:"column:mitra_ao_cov_name" json:"mitra_ao_cov_name"`
	RegencyID                       uint64     `gorm:"column:regency_id" json:"regency_id"`
	RegencyName                     string     `gorm:"column:regency_name" json:"regency_name"`
	MitraAOCoverageRegencyStatus    uint64     `gorm:"column:mitra_ao_cov_rgcy_status" json:"mitra_ao_cov_rgcy_status"`
	MitraAOCoverageRegencyCreatedBy string     `gorm:"column:mitra_ao_cov_rgcy_created_by" json:"mitra_ao_cov_rgcy_created_by"`
	MitraAOCoverageRegencyCreatedAt time.Time  `gorm:"column:mitra_ao_cov_rgcy_created_at" json:"mitra_ao_cov_rgcy_created_at"`
	MitraAOCoverageRegencyUpdatedBy string     `gorm:"column:mitra_ao_cov_rgcy_updated_by" json:"mitra_ao_cov_rgcy_updated_by"`
	MitraAOCoverageRegencyUpdatedAt *time.Time `gorm:"column:mitra_ao_cov_rgcy_updated_at" json:"mitra_ao_cov_rgcy_updated_at"`
	MitraAOCoverageRegencyDeletedBy string     `gorm:"column:mitra_ao_cov_rgcy_deleted_by" json:"mitra_ao_cov_rgcy_deleted_by"`
	MitraAOCoverageRegencyDeletedAt *time.Time `gorm:"column:mitra_ao_cov_rgcy_deleted_at" json:"mitra_ao_cov_rgcy_deleted_at"`
}

type ResponseMitraAOCoverageRegencys struct {
	Status  int                           `json:"status"`
	Message string                        `json:"message"`
	Data    *[]MitraAOCoverageRegencyData `json:"data"`
}

type ResponseMitraAOCoverageRegency struct {
	Status  int                         `json:"status"`
	Message string                      `json:"message"`
	Data    *MitraAOCoverageRegencyData `json:"data"`
}

type ResponseMitraAOCoverageRegencyCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraAOCoverageRegency) TableName() string {
	return "m_mitra_ao_coverage_regency"
}

func (MitraAOCoverageRegencyData) TableName() string {
	return "m_mitra_ao_coverage_regency"
}

func (p *MitraAOCoverageRegency) Prepare() {
	p.MitraAOCoverageRegencyID = nil
	p.MitraAOCoverageID = p.MitraAOCoverageID
	p.RegencyID = p.RegencyID
	p.MitraAOCoverageRegencyStatus = p.MitraAOCoverageRegencyStatus
	p.MitraAOCoverageRegencyCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraAOCoverageRegencyCreatedBy))
	p.MitraAOCoverageRegencyCreatedAt = time.Now()
	p.MitraAOCoverageRegencyUpdatedBy = html.EscapeString(strings.TrimSpace(p.MitraAOCoverageRegencyUpdatedBy))
	p.MitraAOCoverageRegencyUpdatedAt = p.MitraAOCoverageRegencyUpdatedAt
	p.MitraAOCoverageRegencyDeletedBy = html.EscapeString(strings.TrimSpace(p.MitraAOCoverageRegencyDeletedBy))
	p.MitraAOCoverageRegencyDeletedAt = p.MitraAOCoverageRegencyDeletedAt
}

func (p *MitraAOCoverageRegency) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.RegencyID == 0 {
			return errors.New("Required Regency")
		}
		return nil
	}
}

func (p *MitraAOCoverageRegency) SaveMitraAOCoverageRegency(db *gorm.DB) (*MitraAOCoverageRegency, error) {
	var err error
	err = db.Debug().Model(&MitraAOCoverageRegency{}).Create(&p).Error
	if err != nil {
		return &MitraAOCoverageRegency{}, err
	}
	return p, nil
}

func (p *MitraAOCoverageRegency) FindAllMitraAOCoverageRegency(db *gorm.DB) (*[]MitraAOCoverageRegencyData, error) {
	var err error
	mitraaocoverageregency := []MitraAOCoverageRegencyData{}
	err = db.Debug().Model(&MitraAOCoverageRegencyData{}).Limit(100).
		Select(`m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_id,
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
		Joins("JOIN m_mitra_ao_coverage on m_mitra_ao_coverage_regency.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Joins("JOIN m_regency on m_mitra_ao_coverage_regency.regency_id=m_regency.regency_id").
		Find(&mitraaocoverageregency).Error
	if err != nil {
		return &[]MitraAOCoverageRegencyData{}, err
	}
	return &mitraaocoverageregency, nil
}

func (p *MitraAOCoverageRegency) FindAllMitraAOCoverageRegencyStatus(db *gorm.DB, status uint64) (*[]MitraAOCoverageRegencyData, error) {
	var err error
	mitraaocoverageregency := []MitraAOCoverageRegencyData{}
	err = db.Debug().Model(&MitraAOCoverageRegencyData{}).Limit(100).
		Select(`m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_id,
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
		Joins("JOIN m_mitra_ao_coverage on m_mitra_ao_coverage_regency.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Joins("JOIN m_regency on m_mitra_ao_coverage_regency.regency_id=m_regency.regency_id").
		Where("mitra_ao_cov_rgcy_status = ?", status).
		Find(&mitraaocoverageregency).Error
	if err != nil {
		return &[]MitraAOCoverageRegencyData{}, err
	}
	return &mitraaocoverageregency, nil
}

func (p *MitraAOCoverageRegency) FindMitraAOCoverageRegencyDataByID(db *gorm.DB, pid uint64) (*MitraAOCoverageRegency, error) {
	var err error
	err = db.Debug().Model(&MitraAOCoverageRegency{}).Where("mitra_ao_cov_rgcy_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraAOCoverageRegency{}, err
	}
	return p, nil
}

func (p *MitraAOCoverageRegency) FindMitraAOCoverageRegencyByID(db *gorm.DB, pid uint64) (*MitraAOCoverageRegencyData, error) {
	var err error
	mitraaocoverageregency := MitraAOCoverageRegencyData{}
	err = db.Debug().Model(&MitraAOCoverageRegencyData{}).Limit(100).
		Select(`m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_id,
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
		Joins("JOIN m_mitra_ao_coverage on m_mitra_ao_coverage_regency.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Joins("JOIN m_regency on m_mitra_ao_coverage_regency.regency_id=m_regency.regency_id").
		Where("mitra_ao_cov_rgcy_id = ?", pid).
		Find(&mitraaocoverageregency).Error
	if err != nil {
		return &MitraAOCoverageRegencyData{}, err
	}
	return &mitraaocoverageregency, nil
}

func (p *MitraAOCoverageRegency) FindMitraAOCoverageRegencyStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraAOCoverageRegencyData, error) {
	var err error
	mitraaocoverageregency := MitraAOCoverageRegencyData{}
	err = db.Debug().Model(&MitraAOCoverageRegencyData{}).Limit(100).
		Select(`m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_id,
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
		Joins("JOIN m_mitra_ao_coverage on m_mitra_ao_coverage_regency.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Joins("JOIN m_regency on m_mitra_ao_coverage_regency.regency_id=m_regency.regency_id").
		Where("mitra_ao_cov_rgcy_id = ? AND mitra_ao_cov_rgcy_status = ?", pid, status).
		Find(&mitraaocoverageregency).Error
	if err != nil {
		return &MitraAOCoverageRegencyData{}, err
	}
	return &mitraaocoverageregency, nil
}

func (p *MitraAOCoverageRegency) UpdateMitraAOCoverageRegency(db *gorm.DB, pid uint64) (*MitraAOCoverageRegency, error) {
	var err error
	db = db.Debug().Model(&MitraAOCoverageRegency{}).Where("mitra_ao_cov_rgcy_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_ao_cov_id":              p.MitraAOCoverageID,
			"regency_id":                   p.RegencyID,
			"mitra_ao_cov_rgcy_status":     p.MitraAOCoverageRegencyStatus,
			"mitra_ao_cov_rgcy_updated_by": p.MitraAOCoverageRegencyUpdatedBy,
			"mitra_ao_cov_rgcy_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraAOCoverageRegency{}).Where("mitra_ao_cov_rgcy_id = ?", pid).Error
	if err != nil {
		return &MitraAOCoverageRegency{}, err
	}
	return p, nil
}

func (p *MitraAOCoverageRegency) UpdateMitraAOCoverageRegencyStatus(db *gorm.DB, pid uint64) (*MitraAOCoverageRegency, error) {
	var err error
	db = db.Debug().Model(&MitraAOCoverageRegency{}).Where("mitra_ao_cov_rgcy_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_ao_cov_rgcy_status":     p.MitraAOCoverageRegencyStatus,
			"mitra_ao_cov_rgcy_updated_by": p.MitraAOCoverageRegencyUpdatedBy,
			"mitra_ao_cov_rgcy_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraAOCoverageRegency{}).Where("mitra_ao_cov_rgcy_id = ?", pid).Error
	if err != nil {
		return &MitraAOCoverageRegency{}, err
	}
	return p, nil
}

func (p *MitraAOCoverageRegency) UpdateMitraAOCoverageRegencyStatusOnly(db *gorm.DB, pid uint64) (*MitraAOCoverageRegency, error) {
	var err error
	db = db.Debug().Model(&MitraAOCoverageRegency{}).Where("mitra_ao_cov_rgcy_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_ao_cov_rgcy_status": p.MitraAOCoverageRegencyStatus,
		},
	)
	err = db.Debug().Model(&MitraAOCoverageRegency{}).Where("mitra_ao_cov_rgcy_id = ?", pid).Error
	if err != nil {
		return &MitraAOCoverageRegency{}, err
	}
	return p, nil
}

func (p *MitraAOCoverageRegency) UpdateMitraAOCoverageRegencyStatusDelete(db *gorm.DB, pid uint64) (*MitraAOCoverageRegency, error) {
	var err error
	db = db.Debug().Model(&MitraAOCoverageRegency{}).Where("mitra_ao_cov_rgcy_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_ao_cov_rgcy_status":     3,
			"mitra_ao_cov_rgcy_deleted_by": p.MitraAOCoverageRegencyDeletedBy,
			"mitra_ao_cov_rgcy_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraAOCoverageRegency{}).Where("mitra_ao_cov_rgcy_id = ?", pid).Error
	if err != nil {
		return &MitraAOCoverageRegency{}, err
	}
	return p, nil
}

func (p *MitraAOCoverageRegency) DeleteMitraAOCoverageRegency(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraAOCoverageRegency{}).Where("mitra_ao_cov_rgcy_id = ?", pid).Take(&MitraAOCoverageRegency{}).Delete(&MitraAOCoverageRegency{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraAOCoverageRegency not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

type MitraAOPredefinedCoverageRegencyData struct {
	MitraAOCoverageRegencyID uint64 `gorm:"column:mitra_ao_cov_rgcy_id" json:"mitra_ao_cov_rgcy_id"`
	MitraAOCoverageID        uint64 `gorm:"column:mitra_ao_cov_id" json:"mitra_ao_cov_id"`
	RegencyID                uint64 `gorm:"column:regency_id" json:"regency_id"`
	RegencyName              string `gorm:"column:regency_name" json:"regency_name"`
}

type ResponsePredefinedMitraAOCoverageRegencys struct {
	Status  int                                     `json:"status"`
	Message string                                  `json:"message"`
	Data    *[]MitraAOPredefinedCoverageRegencyData `json:"data"`
}

func (MitraAOPredefinedCoverageRegencyData) TableName() string {
	return "m_mitra_ao_coverage_regency"
}

func (p *MitraAOCoverageRegency) FindMitraAOCoverageRegencyByMitraAOCoverageID(db *gorm.DB, pid uint64) ([]MitraAOCoverageRegencyData, error) {
	var err error
	mitraaocoverageprovince := []MitraAOCoverageRegencyData{}
	err = db.Debug().Model(&MitraAOCoverageRegencyData{}).Limit(100).
		Select(`m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_id,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_id,
				m_mitra_ao_coverage.mitra_ao_cov_id,
				m_mitra_ao_coverage.mitra_ao_cov_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_status,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_created_by,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_created_at`).
		Joins("JOIN m_mitra_ao_coverage on m_mitra_ao_coverage_regency.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Joins("JOIN m_regency on m_mitra_ao_coverage_regency.regency_id=m_regency.regency_id").
		Where("m_mitra_ao_coverage_regency.mitra_ao_cov_id = ?", pid).
		Find(&mitraaocoverageprovince).Error
	if err != nil {
		return []MitraAOCoverageRegencyData{}, err
	}
	return mitraaocoverageprovince, nil
}

func (p *MitraAOCoverageRegency) FindMitraAOCoverageRegencyStatusByMitraAOCoverageID(db *gorm.DB, mitra uint64, status uint64) ([]MitraAOCoverageRegencyData, error) {
	var err error
	mitraaocoverageprovince := []MitraAOCoverageRegencyData{}
	err = db.Debug().Model(&MitraAOCoverageRegencyData{}).Limit(100).
		Select(`m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_id,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_id,
				m_mitra_ao_coverage.mitra_ao_cov_id,
				m_mitra_ao_coverage.mitra_ao_cov_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_status,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_created_by,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_created_at`).
		Joins("JOIN m_mitra_ao_coverage on m_mitra_ao_coverage_regency.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Joins("JOIN m_regency on m_mitra_ao_coverage_regency.regency_id=m_regency.regency_id").
		Where("m_mitra_ao_coverage_regency.mitra_ao_cov_id = ? AND m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_status = ?", mitra, status).
		Find(&mitraaocoverageprovince).Error
	if err != nil {
		return []MitraAOCoverageRegencyData{}, err
	}
	return mitraaocoverageprovince, nil
}

func (p *MitraAOCoverageRegency) FindMitraAOPredefinedCoverageRegencyByMitraAOCoverageID(db *gorm.DB, mitra uint64, status uint64) ([]MitraAOPredefinedCoverageRegencyData, error) {
	var err error
	mitraaocoverageprovince := []MitraAOPredefinedCoverageRegencyData{}
	err = db.Debug().Model(&MitraAOPredefinedCoverageRegencyData{}).Limit(100).
		Select(`m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_id,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_id,
				m_mitra_ao_coverage.mitra_ao_cov_id,
				m_mitra_ao_coverage.mitra_ao_cov_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_status,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_created_by,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_created_at`).
		Joins("JOIN m_mitra_ao_coverage on m_mitra_ao_coverage_regency.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Joins("JOIN m_regency on m_mitra_ao_coverage_regency.regency_id=m_regency.regency_id").
		Where("m_mitra_ao_coverage_regency.mitra_ao_cov_id = ? AND m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_status = ?", mitra, status).
		Find(&mitraaocoverageprovince).Error
	if err != nil {
		return []MitraAOPredefinedCoverageRegencyData{}, err
	}
	return mitraaocoverageprovince, nil
}

func (p *MitraAOCoverageRegency) FindMitraAOPredefinedCoverageRegencyByMitraAOCoverageIDByProvince(db *gorm.DB, mitra uint64, province uint64, status uint64) ([]MitraAOPredefinedCoverageRegencyData, error) {
	var err error
	mitraaocoverageprovince := []MitraAOPredefinedCoverageRegencyData{}
	err = db.Debug().Model(&MitraAOPredefinedCoverageRegencyData{}).Limit(100).
		Select(`m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_id,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_id,
				m_mitra_ao_coverage.mitra_ao_cov_id,
				m_mitra_ao_coverage.mitra_ao_cov_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_status,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_created_by,
				m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_created_at`).
		Joins("JOIN m_mitra_ao_coverage on m_mitra_ao_coverage_regency.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Joins("JOIN m_regency on m_mitra_ao_coverage_regency.regency_id=m_regency.regency_id").
		Where("m_mitra_ao_coverage_regency.mitra_ao_cov_id = ? AND m_regency.province_id = ? AND m_mitra_ao_coverage_regency.mitra_ao_cov_rgcy_status = ?", mitra, province, status).
		Find(&mitraaocoverageprovince).Error
	if err != nil {
		return []MitraAOPredefinedCoverageRegencyData{}, err
	}
	return mitraaocoverageprovince, nil
}
