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

type MitraBranchCoverageRegency struct {
	MitraBranchCoverageRegencyID        *uint64    `gorm:"column:mitra_branch_cov_rgcy_id;primary_key;" json:"mitra_branch_cov_rgcy_id"`
	MitraBranchID                       uint64     `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	RegencyID                           uint64     `gorm:"column:regency_id" json:"regency_id"`
	MitraBranchCoverageRegencyStatus    uint64     `gorm:"column:mitra_branch_cov_rgcy_status;size:2" json:"mitra_branch_cov_rgcy_status"`
	MitraBranchCoverageRegencyCreatedBy string     `gorm:"column:mitra_branch_cov_rgcy_created_by;size:125" json:"mitra_branch_cov_rgcy_created_by"`
	MitraBranchCoverageRegencyCreatedAt time.Time  `gorm:"column:mitra_branch_cov_rgcy_created_at;default:CURRENT_TIMESTAMP" json:"mitra_branch_cov_rgcy_created_at"`
	MitraBranchCoverageRegencyUpdatedBy string     `gorm:"column:mitra_branch_cov_rgcy_updated_by;size:125" json:"mitra_branch_cov_rgcy_updated_by"`
	MitraBranchCoverageRegencyUpdatedAt *time.Time `gorm:"column:mitra_branch_cov_rgcy_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_branch_cov_rgcy_created_at"`
	MitraBranchCoverageRegencyDeletedBy string     `gorm:"column:mitra_branch_cov_rgcy_deleted_by;size:125" json:"mitra_branch_cov_rgcy_deleted_by"`
	MitraBranchCoverageRegencyDeletedAt *time.Time `gorm:"column:mitra_branch_cov_rgcy_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_branch_cov_rgcy_deleted_at"`
}

type MitraBranchCoverageRegencyData struct {
	MitraBranchCoverageRegencyID        uint64     `gorm:"column:mitra_branch_cov_rgcy_id" json:"mitra_branch_cov_rgcy_id"`
	MitraBranchID                       uint64     `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	MitraBranchName                     string     `gorm:"column:mitra_branch_name" json:"mitra_branch_name"`
	RegencyID                           uint64     `gorm:"column:regency_id" json:"regency_id"`
	RegencyName                         string     `gorm:"column:regency_name" json:"regency_name"`
	MitraBranchCoverageRegencyStatus    uint64     `gorm:"column:mitra_branch_cov_rgcy_status" json:"mitra_branch_cov_rgcy_status"`
	MitraBranchCoverageRegencyCreatedBy string     `gorm:"column:mitra_branch_cov_rgcy_created_by" json:"mitra_branch_cov_rgcy_created_by"`
	MitraBranchCoverageRegencyCreatedAt time.Time  `gorm:"column:mitra_branch_cov_rgcy_created_at" json:"mitra_branch_cov_rgcy_created_at"`
	MitraBranchCoverageRegencyUpdatedBy string     `gorm:"column:mitra_branch_cov_rgcy_updated_by" json:"mitra_branch_cov_rgcy_updated_by"`
	MitraBranchCoverageRegencyUpdatedAt *time.Time `gorm:"column:mitra_branch_cov_rgcy_updated_at" json:"mitra_branch_cov_rgcy_updated_at"`
	MitraBranchCoverageRegencyDeletedBy string     `gorm:"column:mitra_branch_cov_rgcy_deleted_by" json:"mitra_branch_cov_rgcy_deleted_by"`
	MitraBranchCoverageRegencyDeletedAt *time.Time `gorm:"column:mitra_branch_cov_rgcy_deleted_at" json:"mitra_branch_cov_rgcy_deleted_at"`
}

type ResponseMitraBranchCoverageRegencys struct {
	Status  int                               `json:"status"`
	Message string                            `json:"message"`
	Data    *[]MitraBranchCoverageRegencyData `json:"data"`
}

type ResponseMitraBranchCoverageRegency struct {
	Status  int                             `json:"status"`
	Message string                          `json:"message"`
	Data    *MitraBranchCoverageRegencyData `json:"data"`
}

type ResponseMitraBranchCoverageRegencyCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraBranchCoverageRegency) TableName() string {
	return "m_mitra_branch_coverage_regency"
}

func (MitraBranchCoverageRegencyData) TableName() string {
	return "m_mitra_branch_coverage_regency"
}

func (p *MitraBranchCoverageRegency) Prepare() {
	p.MitraBranchCoverageRegencyID = nil
	p.MitraBranchID = p.MitraBranchID
	p.RegencyID = p.RegencyID
	p.MitraBranchCoverageRegencyStatus = p.MitraBranchCoverageRegencyStatus
	p.MitraBranchCoverageRegencyCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraBranchCoverageRegencyCreatedBy))
	p.MitraBranchCoverageRegencyCreatedAt = time.Now()
	p.MitraBranchCoverageRegencyUpdatedBy = html.EscapeString(strings.TrimSpace(p.MitraBranchCoverageRegencyUpdatedBy))
	p.MitraBranchCoverageRegencyUpdatedAt = p.MitraBranchCoverageRegencyUpdatedAt
	p.MitraBranchCoverageRegencyDeletedBy = html.EscapeString(strings.TrimSpace(p.MitraBranchCoverageRegencyDeletedBy))
	p.MitraBranchCoverageRegencyDeletedAt = p.MitraBranchCoverageRegencyDeletedAt
}

func (p *MitraBranchCoverageRegency) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.MitraBranchID == 0 {
			return errors.New("Required Country")
		}
		if p.RegencyID == 0 {
			return errors.New("Required Regency")
		}
		return nil
	}
}

func (p *MitraBranchCoverageRegency) SaveMitraBranchCoverageRegency(db *gorm.DB) (*MitraBranchCoverageRegency, error) {
	var err error
	err = db.Debug().Model(&MitraBranchCoverageRegency{}).Create(&p).Error
	if err != nil {
		return &MitraBranchCoverageRegency{}, err
	}
	return p, nil
}

func (p *MitraBranchCoverageRegency) FindAllMitraBranchCoverageRegency(db *gorm.DB) (*[]MitraBranchCoverageRegencyData, error) {
	var err error
	mitrabranchcoverageregency := []MitraBranchCoverageRegencyData{}
	err = db.Debug().Model(&MitraBranchCoverageRegencyData{}).Limit(100).
		Select(`m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_id,
				m_mitra_branch.mitra_branch_id,
				m_mitra.mitra_branch_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_status,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_created_by,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_created_at,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_updated_by,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_updated_at,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_deleted_by,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_deleted_at`).
		Joins("JOIN m_mitra on m_mitra_branch_coverage_regency.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_regency on m_mitra_branch_coverage_regency.regency_id=m_regency.regency_id").
		Find(&mitrabranchcoverageregency).Error
	if err != nil {
		return &[]MitraBranchCoverageRegencyData{}, err
	}
	return &mitrabranchcoverageregency, nil
}

func (p *MitraBranchCoverageRegency) FindAllMitraBranchCoverageRegencyStatus(db *gorm.DB, status uint64) (*[]MitraBranchCoverageRegencyData, error) {
	var err error
	mitrabranchcoverageregency := []MitraBranchCoverageRegencyData{}
	err = db.Debug().Model(&MitraBranchCoverageRegencyData{}).Limit(100).
		Select(`m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_id,
				m_mitra_branch.mitra_branch_id,
				m_mitra.mitra_branch_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_status,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_created_by,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_created_at,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_updated_by,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_updated_at,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_deleted_by,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_deleted_at`).
		Joins("JOIN m_mitra on m_mitra_branch_coverage_regency.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_regency on m_mitra_branch_coverage_regency.regency_id=m_regency.regency_id").
		Where("mitra_branch_cov_rgcy_status = ?", status).
		Find(&mitrabranchcoverageregency).Error
	if err != nil {
		return &[]MitraBranchCoverageRegencyData{}, err
	}
	return &mitrabranchcoverageregency, nil
}

func (p *MitraBranchCoverageRegency) FindMitraBranchCoverageRegencyDataByID(db *gorm.DB, pid uint64) (*MitraBranchCoverageRegency, error) {
	var err error
	err = db.Debug().Model(&MitraBranchCoverageRegency{}).Where("mitra_branch_cov_rgcy_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraBranchCoverageRegency{}, err
	}
	return p, nil
}

func (p *MitraBranchCoverageRegency) FindMitraBranchCoverageRegencyByID(db *gorm.DB, pid uint64) (*MitraBranchCoverageRegencyData, error) {
	var err error
	mitrabranchcoverageregency := MitraBranchCoverageRegencyData{}
	err = db.Debug().Model(&MitraBranchCoverageRegencyData{}).Limit(100).
		Select(`m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_id,
				m_mitra_branch.mitra_branch_id,
				m_mitra.mitra_branch_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_status,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_created_by,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_created_at,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_updated_by,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_updated_at,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_deleted_by,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_deleted_at`).
		Joins("JOIN m_mitra on m_mitra_branch_coverage_regency.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_regency on m_mitra_branch_coverage_regency.regency_id=m_regency.regency_id").
		Where("mitra_branch_cov_rgcy_id = ?", pid).
		Find(&mitrabranchcoverageregency).Error
	if err != nil {
		return &MitraBranchCoverageRegencyData{}, err
	}
	return &mitrabranchcoverageregency, nil
}

func (p *MitraBranchCoverageRegency) FindMitraBranchCoverageRegencyStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraBranchCoverageRegencyData, error) {
	var err error
	mitrabranchcoverageregency := MitraBranchCoverageRegencyData{}
	err = db.Debug().Model(&MitraBranchCoverageRegencyData{}).Limit(100).
		Select(`m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_id,
				m_mitra_branch.mitra_branch_id,
				m_mitra.mitra_branch_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_status,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_created_by,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_created_at,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_updated_by,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_updated_at,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_deleted_by,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_deleted_at`).
		Joins("JOIN m_mitra on m_mitra_branch_coverage_regency.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_regency on m_mitra_branch_coverage_regency.regency_id=m_regency.regency_id").
		Where("mitra_branch_cov_rgcy_id = ? AND mitra_branch_cov_rgcy_status = ?", pid, status).
		Find(&mitrabranchcoverageregency).Error
	if err != nil {
		return &MitraBranchCoverageRegencyData{}, err
	}
	return &mitrabranchcoverageregency, nil
}

func (p *MitraBranchCoverageRegency) UpdateMitraBranchCoverageRegency(db *gorm.DB, pid uint64) (*MitraBranchCoverageRegency, error) {
	var err error
	db = db.Debug().Model(&MitraBranchCoverageRegency{}).Where("mitra_branch_cov_rgcy_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_branch_id":                  p.MitraBranchID,
			"regency_id":                       p.RegencyID,
			"mitra_branch_cov_rgcy_status":     p.MitraBranchCoverageRegencyStatus,
			"mitra_branch_cov_rgcy_updated_by": p.MitraBranchCoverageRegencyUpdatedBy,
			"mitra_branch_cov_rgcy_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraBranchCoverageRegency{}).Where("mitra_branch_cov_rgcy_id = ?", pid).Error
	if err != nil {
		return &MitraBranchCoverageRegency{}, err
	}
	return p, nil
}

func (p *MitraBranchCoverageRegency) UpdateMitraBranchCoverageRegencyStatus(db *gorm.DB, pid uint64) (*MitraBranchCoverageRegency, error) {
	var err error
	db = db.Debug().Model(&MitraBranchCoverageRegency{}).Where("mitra_branch_cov_rgcy_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_branch_cov_rgcy_status":     p.MitraBranchCoverageRegencyStatus,
			"mitra_branch_cov_rgcy_updated_by": p.MitraBranchCoverageRegencyUpdatedBy,
			"mitra_branch_cov_rgcy_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraBranchCoverageRegency{}).Where("mitra_branch_cov_rgcy_id = ?", pid).Error
	if err != nil {
		return &MitraBranchCoverageRegency{}, err
	}
	return p, nil
}

func (p *MitraBranchCoverageRegency) UpdateMitraBranchCoverageRegencyStatusOnly(db *gorm.DB, pid uint64) (*MitraBranchCoverageRegency, error) {
	var err error
	db = db.Debug().Model(&MitraBranchCoverageRegency{}).Where("mitra_branch_cov_rgcy_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_branch_cov_rgcy_status": p.MitraBranchCoverageRegencyStatus,
		},
	)
	err = db.Debug().Model(&MitraBranchCoverageRegency{}).Where("mitra_branch_cov_rgcy_id = ?", pid).Error
	if err != nil {
		return &MitraBranchCoverageRegency{}, err
	}
	return p, nil
}

func (p *MitraBranchCoverageRegency) UpdateMitraBranchCoverageRegencyStatusDelete(db *gorm.DB, pid uint64) (*MitraBranchCoverageRegency, error) {
	var err error
	db = db.Debug().Model(&MitraBranchCoverageRegency{}).Where("mitra_branch_cov_rgcy_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_branch_cov_rgcy_status":     3,
			"mitra_branch_cov_rgcy_deleted_by": p.MitraBranchCoverageRegencyDeletedBy,
			"mitra_branch_cov_rgcy_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraBranchCoverageRegency{}).Where("mitra_branch_cov_rgcy_id = ?", pid).Error
	if err != nil {
		return &MitraBranchCoverageRegency{}, err
	}
	return p, nil
}

func (p *MitraBranchCoverageRegency) DeleteMitraBranchCoverageRegency(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraBranchCoverageRegency{}).Where("mitra_branch_cov_rgcy_id = ?", pid).Take(&MitraBranchCoverageRegency{}).Delete(&MitraBranchCoverageRegency{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraBranchCoverageRegency not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

type MitraBranchPredefinedCoverageRegencyData struct {
	MitraBranchCoverageRegencyID uint64 `gorm:"column:mitra_branch_cov_rgcy_id" json:"mitra_branch_cov_rgcy_id"`
	MitraBranchID                uint64 `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	RegencyID                    uint64 `gorm:"column:regency_id" json:"regency_id"`
	RegencyName                  string `gorm:"column:regency_name" json:"regency_name"`
}

type ResponsePredefinedMitraBranchCoverageRegencys struct {
	Status  int                                         `json:"status"`
	Message string                                      `json:"message"`
	Data    *[]MitraBranchPredefinedCoverageRegencyData `json:"data"`
}

func (MitraBranchPredefinedCoverageRegencyData) TableName() string {
	return "m_mitra_branch_coverage_regency"
}

func (p *MitraBranchCoverageRegency) FindMitraBranchCoverageRegencyByMitraBranchID(db *gorm.DB, pid uint64) ([]MitraBranchCoverageRegencyData, error) {
	var err error
	mitrabranchcoverageprovince := []MitraBranchCoverageRegencyData{}
	err = db.Debug().Model(&MitraBranchCoverageRegencyData{}).Limit(100).
		Select(`m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_id,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_id,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_status,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_created_by,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_created_at`).
		Joins("JOIN m_mitra_branch on m_mitra_branch_coverage_regency.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_regency on m_mitra_branch_coverage_regency.regency_id=m_regency.regency_id").
		Where("m_mitra_branch_coverage_regency.mitra_branch_id = ?", pid).
		Find(&mitrabranchcoverageprovince).Error
	if err != nil {
		return []MitraBranchCoverageRegencyData{}, err
	}
	return mitrabranchcoverageprovince, nil
}

func (p *MitraBranchCoverageRegency) FindMitraBranchCoverageRegencyStatusByMitraBranchID(db *gorm.DB, mitra uint64, status uint64) ([]MitraBranchCoverageRegencyData, error) {
	var err error
	mitrabranchcoverageprovince := []MitraBranchCoverageRegencyData{}
	err = db.Debug().Model(&MitraBranchCoverageRegencyData{}).Limit(100).
		Select(`m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_id,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_id,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_status,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_created_by,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_created_at`).
		Joins("JOIN m_mitra_branch on m_mitra_branch_coverage_regency.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_regency on m_mitra_branch_coverage_regency.regency_id=m_regency.regency_id").
		Where("m_mitra_branch_coverage_regency.mitra_branch_id = ? AND m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_status = ?", mitra, status).
		Find(&mitrabranchcoverageprovince).Error
	if err != nil {
		return []MitraBranchCoverageRegencyData{}, err
	}
	return mitrabranchcoverageprovince, nil
}

func (p *MitraBranchCoverageRegency) FindMitraBranchPredefinedCoverageRegencyByMitraBranchID(db *gorm.DB, mitra uint64, status uint64) ([]MitraBranchPredefinedCoverageRegencyData, error) {
	var err error
	mitrabranchcoverageprovince := []MitraBranchPredefinedCoverageRegencyData{}
	err = db.Debug().Model(&MitraBranchPredefinedCoverageRegencyData{}).Limit(100).
		Select(`m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_id,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_id,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_status,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_created_by,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_created_at`).
		Joins("JOIN m_mitra_branch on m_mitra_branch_coverage_regency.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_regency on m_mitra_branch_coverage_regency.regency_id=m_regency.regency_id").
		Where("m_mitra_branch_coverage_regency.mitra_branch_id = ? AND m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_status = ?", mitra, status).
		Find(&mitrabranchcoverageprovince).Error
	if err != nil {
		return []MitraBranchPredefinedCoverageRegencyData{}, err
	}
	return mitrabranchcoverageprovince, nil
}

func (p *MitraBranchCoverageRegency) FindMitraBranchPredefinedCoverageRegencyByMitraBranchIDByProvince(db *gorm.DB, mitra uint64, province uint64, status uint64) ([]MitraBranchPredefinedCoverageRegencyData, error) {
	var err error
	mitrabranchcoverageprovince := []MitraBranchPredefinedCoverageRegencyData{}
	err = db.Debug().Model(&MitraBranchPredefinedCoverageRegencyData{}).Limit(100).
		Select(`m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_id,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_id,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_status,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_created_by,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_created_at`).
		Joins("JOIN m_mitra_branch on m_mitra_branch_coverage_regency.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_regency on m_mitra_branch_coverage_regency.regency_id=m_regency.regency_id").
		Where("m_mitra_branch_coverage_regency.mitra_branch_id = ? AND m_regency.province_id = ? AND m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_status = ?", mitra, province, status).
		Find(&mitrabranchcoverageprovince).Error
	if err != nil {
		return []MitraBranchPredefinedCoverageRegencyData{}, err
	}
	return mitrabranchcoverageprovince, nil
}
