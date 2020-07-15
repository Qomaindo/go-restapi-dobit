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

type MitraCoverageRegency struct {
	MitraCoverageRegencyID        *uint64    `gorm:"column:mitra_cov_rgcy_id;primary_key;" json:"mitra_cov_rgcy_id"`
	MitraID                       uint64     `gorm:"column:mitra_id" json:"mitra_id"`
	RegencyID                     uint64     `gorm:"column:regency_id" json:"regency_id"`
	MitraCoverageRegencyStatus    uint64     `gorm:"column:mitra_cov_rgcy_status;size:2" json:"mitra_cov_rgcy_status"`
	MitraCoverageRegencyCreatedBy string     `gorm:"column:mitra_cov_rgcy_created_by;size:125" json:"mitra_cov_rgcy_created_by"`
	MitraCoverageRegencyCreatedAt time.Time  `gorm:"column:mitra_cov_rgcy_created_at;default:CURRENT_TIMESTAMP" json:"mitra_cov_rgcy_created_at"`
	MitraCoverageRegencyUpdatedBy string     `gorm:"column:mitra_cov_rgcy_updated_by;size:125" json:"mitra_cov_rgcy_updated_by"`
	MitraCoverageRegencyUpdatedAt *time.Time `gorm:"column:mitra_cov_rgcy_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_cov_rgcy_created_at"`
	MitraCoverageRegencyDeletedBy string     `gorm:"column:mitra_cov_rgcy_deleted_by;size:125" json:"mitra_cov_rgcy_deleted_by"`
	MitraCoverageRegencyDeletedAt *time.Time `gorm:"column:mitra_cov_rgcy_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_cov_rgcy_deleted_at"`
}

type MitraCoverageRegencyData struct {
	MitraCoverageRegencyID        uint64     `gorm:"column:mitra_cov_rgcy_id" json:"mitra_cov_rgcy_id"`
	MitraID                       uint64     `gorm:"column:mitra_id" json:"mitra_id"`
	MitraName                     string     `gorm:"column:mitra_name" json:"mitra_name"`
	RegencyID                     uint64     `gorm:"column:regency_id" json:"regency_id"`
	RegencyName                   string     `gorm:"column:regency_name" json:"regency_name"`
	MitraCoverageRegencyStatus    uint64     `gorm:"column:mitra_cov_rgcy_status" json:"mitra_cov_rgcy_status"`
	MitraCoverageRegencyCreatedBy string     `gorm:"column:mitra_cov_rgcy_created_by" json:"mitra_cov_rgcy_created_by"`
	MitraCoverageRegencyCreatedAt time.Time  `gorm:"column:mitra_cov_rgcy_created_at" json:"mitra_cov_rgcy_created_at"`
	MitraCoverageRegencyUpdatedBy string     `gorm:"column:mitra_cov_rgcy_updated_by" json:"mitra_cov_rgcy_updated_by"`
	MitraCoverageRegencyUpdatedAt *time.Time `gorm:"column:mitra_cov_rgcy_updated_at" json:"mitra_cov_rgcy_updated_at"`
	MitraCoverageRegencyDeletedBy string     `gorm:"column:mitra_cov_rgcy_deleted_by" json:"mitra_cov_rgcy_deleted_by"`
	MitraCoverageRegencyDeletedAt *time.Time `gorm:"column:mitra_cov_rgcy_deleted_at" json:"mitra_cov_rgcy_deleted_at"`
}

type ResponseMitraCoverageRegencys struct {
	Status  int                         `json:"status"`
	Message string                      `json:"message"`
	Data    *[]MitraCoverageRegencyData `json:"data"`
}

type ResponseMitraCoverageRegency struct {
	Status  int                       `json:"status"`
	Message string                    `json:"message"`
	Data    *MitraCoverageRegencyData `json:"data"`
}

type ResponseMitraCoverageRegencyCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraCoverageRegency) TableName() string {
	return "m_mitra_coverage_regency"
}

func (MitraCoverageRegencyData) TableName() string {
	return "m_mitra_coverage_regency"
}

func (p *MitraCoverageRegency) Prepare() {
	p.MitraCoverageRegencyID = nil
	p.MitraID = p.MitraID
	p.RegencyID = p.RegencyID
	p.MitraCoverageRegencyStatus = p.MitraCoverageRegencyStatus
	p.MitraCoverageRegencyCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraCoverageRegencyCreatedBy))
	p.MitraCoverageRegencyCreatedAt = time.Now()
	p.MitraCoverageRegencyUpdatedBy = html.EscapeString(strings.TrimSpace(p.MitraCoverageRegencyUpdatedBy))
	p.MitraCoverageRegencyUpdatedAt = p.MitraCoverageRegencyUpdatedAt
	p.MitraCoverageRegencyDeletedBy = html.EscapeString(strings.TrimSpace(p.MitraCoverageRegencyDeletedBy))
	p.MitraCoverageRegencyDeletedAt = p.MitraCoverageRegencyDeletedAt
}

func (p *MitraCoverageRegency) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.MitraID == 0 {
			return errors.New("Required Country")
		}
		if p.RegencyID == 0 {
			return errors.New("Required Regency")
		}
		return nil
	}
}

func (p *MitraCoverageRegency) SaveMitraCoverageRegency(db *gorm.DB) (*MitraCoverageRegency, error) {
	var err error
	err = db.Debug().Model(&MitraCoverageRegency{}).Create(&p).Error
	if err != nil {
		return &MitraCoverageRegency{}, err
	}
	return p, nil
}

func (p *MitraCoverageRegency) FindAllMitraCoverageRegency(db *gorm.DB) (*[]MitraCoverageRegencyData, error) {
	var err error
	mitracoverageregency := []MitraCoverageRegencyData{}
	err = db.Debug().Model(&MitraCoverageRegencyData{}).Limit(100).
		Select(`m_mitra_coverage_regency.mitra_cov_rgcy_id,
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
		Joins("JOIN m_mitra on m_mitra_coverage_regency.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_regency on m_mitra_coverage_regency.regency_id=m_regency.regency_id").
		Find(&mitracoverageregency).Error
	if err != nil {
		return &[]MitraCoverageRegencyData{}, err
	}
	return &mitracoverageregency, nil
}

func (p *MitraCoverageRegency) FindAllMitraCoverageRegencyStatus(db *gorm.DB, status uint64) (*[]MitraCoverageRegencyData, error) {
	var err error
	mitracoverageregency := []MitraCoverageRegencyData{}
	err = db.Debug().Model(&MitraCoverageRegencyData{}).Limit(100).
		Select(`m_mitra_coverage_regency.mitra_cov_rgcy_id,
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
		Joins("JOIN m_mitra on m_mitra_coverage_regency.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_regency on m_mitra_coverage_regency.regency_id=m_regency.regency_id").
		Where("mitra_cov_rgcy_status = ?", status).
		Find(&mitracoverageregency).Error
	if err != nil {
		return &[]MitraCoverageRegencyData{}, err
	}
	return &mitracoverageregency, nil
}

func (p *MitraCoverageRegency) FindMitraCoverageRegencyDataByID(db *gorm.DB, pid uint64) (*MitraCoverageRegency, error) {
	var err error
	err = db.Debug().Model(&MitraCoverageRegency{}).Where("mitra_cov_rgcy_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraCoverageRegency{}, err
	}
	return p, nil
}

func (p *MitraCoverageRegency) FindMitraCoverageRegencyByID(db *gorm.DB, pid uint64) (*MitraCoverageRegencyData, error) {
	var err error
	mitracoverageregency := MitraCoverageRegencyData{}
	err = db.Debug().Model(&MitraCoverageRegencyData{}).Limit(100).
		Select(`m_mitra_coverage_regency.mitra_cov_rgcy_id,
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
		Joins("JOIN m_mitra on m_mitra_coverage_regency.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_regency on m_mitra_coverage_regency.regency_id=m_regency.regency_id").
		Where("mitra_cov_rgcy_id = ?", pid).
		Find(&mitracoverageregency).Error
	if err != nil {
		return &MitraCoverageRegencyData{}, err
	}
	return &mitracoverageregency, nil
}

func (p *MitraCoverageRegency) FindMitraCoverageRegencyStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraCoverageRegencyData, error) {
	var err error
	mitracoverageregency := MitraCoverageRegencyData{}
	err = db.Debug().Model(&MitraCoverageRegencyData{}).Limit(100).
		Select(`m_mitra_coverage_regency.mitra_cov_rgcy_id,
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
		Joins("JOIN m_mitra on m_mitra_coverage_regency.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_regency on m_mitra_coverage_regency.regency_id=m_regency.regency_id").
		Where("mitra_cov_rgcy_id = ? AND mitra_cov_rgcy_status = ?", pid, status).
		Find(&mitracoverageregency).Error
	if err != nil {
		return &MitraCoverageRegencyData{}, err
	}
	return &mitracoverageregency, nil
}

func (p *MitraCoverageRegency) UpdateMitraCoverageRegency(db *gorm.DB, pid uint64) (*MitraCoverageRegency, error) {
	var err error
	db = db.Debug().Model(&MitraCoverageRegency{}).Where("mitra_cov_rgcy_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_id":                  p.MitraID,
			"regency_id":                p.RegencyID,
			"mitra_cov_rgcy_status":     p.MitraCoverageRegencyStatus,
			"mitra_cov_rgcy_updated_by": p.MitraCoverageRegencyUpdatedBy,
			"mitra_cov_rgcy_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraCoverageRegency{}).Where("mitra_cov_rgcy_id = ?", pid).Error
	if err != nil {
		return &MitraCoverageRegency{}, err
	}
	return p, nil
}

func (p *MitraCoverageRegency) UpdateMitraCoverageRegencyStatus(db *gorm.DB, pid uint64) (*MitraCoverageRegency, error) {
	var err error
	db = db.Debug().Model(&MitraCoverageRegency{}).Where("mitra_cov_rgcy_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_cov_rgcy_status":     p.MitraCoverageRegencyStatus,
			"mitra_cov_rgcy_updated_by": p.MitraCoverageRegencyUpdatedBy,
			"mitra_cov_rgcy_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraCoverageRegency{}).Where("mitra_cov_rgcy_id = ?", pid).Error
	if err != nil {
		return &MitraCoverageRegency{}, err
	}
	return p, nil
}

func (p *MitraCoverageRegency) UpdateMitraCoverageRegencyStatusOnly(db *gorm.DB, pid uint64) (*MitraCoverageRegency, error) {
	var err error
	db = db.Debug().Model(&MitraCoverageRegency{}).Where("mitra_cov_rgcy_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_cov_rgcy_status": p.MitraCoverageRegencyStatus,
		},
	)
	err = db.Debug().Model(&MitraCoverageRegency{}).Where("mitra_cov_rgcy_id = ?", pid).Error
	if err != nil {
		return &MitraCoverageRegency{}, err
	}
	return p, nil
}

func (p *MitraCoverageRegency) UpdateMitraCoverageRegencyStatusDelete(db *gorm.DB, pid uint64) (*MitraCoverageRegency, error) {
	var err error
	db = db.Debug().Model(&MitraCoverageRegency{}).Where("mitra_cov_rgcy_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_cov_rgcy_status":     3,
			"mitra_cov_rgcy_deleted_by": p.MitraCoverageRegencyDeletedBy,
			"mitra_cov_rgcy_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraCoverageRegency{}).Where("mitra_cov_rgcy_id = ?", pid).Error
	if err != nil {
		return &MitraCoverageRegency{}, err
	}
	return p, nil
}

func (p *MitraCoverageRegency) DeleteMitraCoverageRegency(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraCoverageRegency{}).Where("mitra_cov_rgcy_id = ?", pid).Take(&MitraCoverageRegency{}).Delete(&MitraCoverageRegency{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraCoverageRegency not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

type MitraPredefinedCoverageRegencyData struct {
	MitraCoverageRegencyID uint64 `gorm:"column:mitra_cov_rgcy_id" json:"mitra_cov_rgcy_id"`
	MitraID                uint64 `gorm:"column:mitra_id" json:"mitra_id"`
	RegencyID              uint64 `gorm:"column:regency_id" json:"regency_id"`
	RegencyName            string `gorm:"column:regency_name" json:"regency_name"`
}

type ResponsePredefinedMitraCoverageRegencys struct {
	Status  int                                   `json:"status"`
	Message string                                `json:"message"`
	Data    *[]MitraPredefinedCoverageRegencyData `json:"data"`
}

func (MitraPredefinedCoverageRegencyData) TableName() string {
	return "m_mitra_coverage_regency"
}

func (p *MitraCoverageRegency) FindMitraCoverageRegencyByMitraID(db *gorm.DB, pid uint64) ([]MitraCoverageRegencyData, error) {
	var err error
	mitracoverageprovince := []MitraCoverageRegencyData{}
	err = db.Debug().Model(&MitraCoverageRegencyData{}).Limit(100).
		Select(`m_mitra_coverage_regency.mitra_cov_rgcy_id,
				m_mitra_coverage_regency.mitra_cov_rgcy_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_mitra_coverage_regency.mitra_cov_rgcy_status,
				m_mitra_coverage_regency.mitra_cov_rgcy_created_by,
				m_mitra_coverage_regency.mitra_cov_rgcy_created_at`).
		Joins("JOIN m_mitra on m_mitra_coverage_regency.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_regency on m_mitra_coverage_regency.regency_id=m_regency.regency_id").
		Where("m_mitra_coverage_regency.mitra_id = ?", pid).
		Find(&mitracoverageprovince).Error
	if err != nil {
		return []MitraCoverageRegencyData{}, err
	}
	return mitracoverageprovince, nil
}

func (p *MitraCoverageRegency) FindMitraCoverageRegencyStatusByMitraID(db *gorm.DB, mitra uint64, status uint64) ([]MitraCoverageRegencyData, error) {
	var err error
	mitracoverageprovince := []MitraCoverageRegencyData{}
	err = db.Debug().Model(&MitraCoverageRegencyData{}).Limit(100).
		Select(`m_mitra_coverage_regency.mitra_cov_rgcy_id,
				m_mitra_coverage_regency.mitra_cov_rgcy_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_mitra_coverage_regency.mitra_cov_rgcy_status,
				m_mitra_coverage_regency.mitra_cov_rgcy_created_by,
				m_mitra_coverage_regency.mitra_cov_rgcy_created_at`).
		Joins("JOIN m_mitra on m_mitra_coverage_regency.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_regency on m_mitra_coverage_regency.regency_id=m_regency.regency_id").
		Where("m_mitra_coverage_regency.mitra_id = ? AND m_mitra_coverage_regency.mitra_cov_rgcy_status = ?", mitra, status).
		Find(&mitracoverageprovince).Error
	if err != nil {
		return []MitraCoverageRegencyData{}, err
	}
	return mitracoverageprovince, nil
}

func (p *MitraCoverageRegency) FindMitraPredefinedCoverageRegencyByMitraID(db *gorm.DB, mitra uint64, status uint64) ([]MitraPredefinedCoverageRegencyData, error) {
	var err error
	mitracoverageprovince := []MitraPredefinedCoverageRegencyData{}
	err = db.Debug().Model(&MitraPredefinedCoverageRegencyData{}).Limit(100).
		Select(`m_mitra_coverage_regency.mitra_cov_rgcy_id,
				m_mitra_coverage_regency.mitra_cov_rgcy_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_mitra_coverage_regency.mitra_cov_rgcy_status,
				m_mitra_coverage_regency.mitra_cov_rgcy_created_by,
				m_mitra_coverage_regency.mitra_cov_rgcy_created_at`).
		Joins("JOIN m_mitra on m_mitra_coverage_regency.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_regency on m_mitra_coverage_regency.regency_id=m_regency.regency_id").
		Where("m_mitra_coverage_regency.mitra_id = ? AND m_mitra_coverage_regency.mitra_cov_rgcy_status = ?", mitra, status).
		Find(&mitracoverageprovince).Error
	if err != nil {
		return []MitraPredefinedCoverageRegencyData{}, err
	}
	return mitracoverageprovince, nil
}

func (p *MitraCoverageRegency) FindMitraPredefinedCoverageRegencyByMitraIDByProvince(db *gorm.DB, mitra uint64, province uint64, status uint64) ([]MitraPredefinedCoverageRegencyData, error) {
	var err error
	mitracoverageprovince := []MitraPredefinedCoverageRegencyData{}
	err = db.Debug().Model(&MitraPredefinedCoverageRegencyData{}).Limit(100).
		Select(`m_mitra_coverage_regency.mitra_cov_rgcy_id,
				m_mitra_coverage_regency.mitra_cov_rgcy_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_mitra_coverage_regency.mitra_cov_rgcy_status,
				m_mitra_coverage_regency.mitra_cov_rgcy_created_by,
				m_mitra_coverage_regency.mitra_cov_rgcy_created_at`).
		Joins("JOIN m_mitra on m_mitra_coverage_regency.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_regency on m_mitra_coverage_regency.regency_id=m_regency.regency_id").
		Where("m_mitra_coverage_regency.mitra_id = ? AND m_regency.province_id = ? AND m_mitra_coverage_regency.mitra_cov_rgcy_status = ?", mitra, province, status).
		Find(&mitracoverageprovince).Error
	if err != nil {
		return []MitraPredefinedCoverageRegencyData{}, err
	}
	return mitracoverageprovince, nil
}
