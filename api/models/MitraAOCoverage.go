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

type MitraAOCoverage struct {
	MitraAOCoverageID        *uint64    `gorm:"column:mitra_ao_cov_id;primary_key;" json:"mitra_ao_cov_id"`
	MitraID                  uint64     `gorm:"column:mitra_id" json:"mitra_id"`
	MitraAOCoverageCode      string     `gorm:"column:mitra_ao_cov_code" json:"mitra_ao_cov_code"`
	MitraAOCoverageName      string     `gorm:"column:mitra_ao_cov_name" json:"mitra_ao_cov_name"`
	MitraAOCoverageStatus    uint64     `gorm:"column:mitra_ao_cov_status;size:2" json:"mitra_ao_cov_status"`
	MitraAOCoverageCreatedBy string     `gorm:"column:mitra_ao_cov_created_by;size:125" json:"mitra_ao_cov_created_by"`
	MitraAOCoverageCreatedAt time.Time  `gorm:"column:mitra_ao_cov_created_at;default:CURRENT_TIMESTAMP" json:"mitra_ao_cov_created_at"`
	MitraAOCoverageUpdatedBy string     `gorm:"column:mitra_ao_cov_updated_by;size:125" json:"mitra_ao_cov_updated_by"`
	MitraAOCoverageUpdatedAt *time.Time `gorm:"column:mitra_ao_cov_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_ao_cov_created_at"`
	MitraAOCoverageDeletedBy string     `gorm:"column:mitra_ao_cov_deleted_by;size:125" json:"mitra_ao_cov_deleted_by"`
	MitraAOCoverageDeletedAt *time.Time `gorm:"column:mitra_ao_cov_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_ao_cov_deleted_at"`
}

type MitraAOCoverageData struct {
	MitraAOCoverageID        uint64     `gorm:"column:mitra_ao_cov_id" json:"mitra_ao_cov_id"`
	MitraID                  uint64     `gorm:"column:mitra_id" json:"mitra_id"`
	MitraName                string     `gorm:"column:mitra_name" json:"mitra_name"`
	MitraAOCoverageCode      string     `gorm:"column:mitra_ao_cov_code" json:"mitra_ao_cov_code"`
	MitraAOCoverageName      string     `gorm:"column:mitra_ao_cov_name" json:"mitra_ao_cov_name"`
	MitraAOCoverageStatus    uint64     `gorm:"column:mitra_ao_cov_status" json:"mitra_ao_cov_status"`
	MitraAOCoverageCreatedBy string     `gorm:"column:mitra_ao_cov_created_by" json:"mitra_ao_cov_created_by"`
	MitraAOCoverageCreatedAt time.Time  `gorm:"column:mitra_ao_cov_created_at" json:"mitra_ao_cov_created_at"`
	MitraAOCoverageUpdatedBy string     `gorm:"column:mitra_ao_cov_updated_by" json:"mitra_ao_cov_updated_by"`
	MitraAOCoverageUpdatedAt *time.Time `gorm:"column:mitra_ao_cov_updated_at" json:"mitra_ao_cov_updated_at"`
	MitraAOCoverageDeletedBy string     `gorm:"column:mitra_ao_cov_deleted_by" json:"mitra_ao_cov_deleted_by"`
	MitraAOCoverageDeletedAt *time.Time `gorm:"column:mitra_ao_cov_deleted_at" json:"mitra_ao_cov_deleted_at"`

	MitraAOCoverageRegency []MitraAOCoverageRegencyData `json:"mitra_ao_cov_regency"`
}

type ResponseMitraAOCoverages struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    *[]MitraAOCoverageData `json:"data"`
}

type ResponseMitraAOCoverage struct {
	Status  int                  `json:"status"`
	Message string               `json:"message"`
	Data    *MitraAOCoverageData `json:"data"`
}

type ResponseMitraAOCoverageCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraAOCoverage) TableName() string {
	return "m_mitra_ao_coverage"
}

func (MitraAOCoverageData) TableName() string {
	return "m_mitra_ao_coverage"
}

func (p *MitraAOCoverage) Prepare() {
	p.MitraAOCoverageID = nil
	p.MitraID = p.MitraID
	p.MitraAOCoverageCode = p.MitraAOCoverageCode
	p.MitraAOCoverageName = p.MitraAOCoverageName
	p.MitraAOCoverageStatus = p.MitraAOCoverageStatus
	p.MitraAOCoverageCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraAOCoverageCreatedBy))
	p.MitraAOCoverageCreatedAt = time.Now()
	p.MitraAOCoverageUpdatedBy = html.EscapeString(strings.TrimSpace(p.MitraAOCoverageUpdatedBy))
	p.MitraAOCoverageUpdatedAt = p.MitraAOCoverageUpdatedAt
	p.MitraAOCoverageDeletedBy = html.EscapeString(strings.TrimSpace(p.MitraAOCoverageDeletedBy))
	p.MitraAOCoverageDeletedAt = p.MitraAOCoverageDeletedAt
}

func (p *MitraAOCoverage) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.MitraAOCoverageCode == "" {
			return errors.New("Required Country")
		}
		if p.MitraAOCoverageName == "" {
			return errors.New("Required Province")
		}
		return nil
	}
}

func (p *MitraAOCoverage) SaveMitraAOCoverage(db *gorm.DB) (*MitraAOCoverage, error) {
	var err error
	err = db.Debug().Model(&MitraAOCoverage{}).Create(&p).Error
	if err != nil {
		return &MitraAOCoverage{}, err
	}
	return p, nil
}

func (p *MitraAOCoverage) FindAllMitraAOCoverage(db *gorm.DB) (*[]MitraAOCoverageData, error) {
	var err error
	mitraaocoverage := []MitraAOCoverageData{}
	err = db.Debug().Model(&MitraAOCoverageData{}).Limit(100).
		Select(`m_mitra_ao_coverage.mitra_ao_cov_id,
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
		Joins("JOIN m_mitra on m_mitra_ao_coverage.mitra_id=m_mitra.mitra_id").
		Find(&mitraaocoverage).Error
	if err != nil {
		return &[]MitraAOCoverageData{}, err
	}
	return &mitraaocoverage, nil
}

func (p *MitraAOCoverage) FindAllMitraAOCoverageStatus(db *gorm.DB, status uint64) (*[]MitraAOCoverageData, error) {
	var err error
	mitraaocoverage := []MitraAOCoverageData{}
	err = db.Debug().Model(&MitraAOCoverageData{}).Limit(100).
		Select(`m_mitra_ao_coverage.mitra_ao_cov_id,
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
		Joins("JOIN m_mitra on m_mitra_ao_coverage.mitra_id=m_mitra.mitra_id").
		Where("mitra_ao_cov_status = ?", status).
		Find(&mitraaocoverage).Error
	if err != nil {
		return &[]MitraAOCoverageData{}, err
	}
	return &mitraaocoverage, nil
}

func (p *MitraAOCoverage) FindMitraAOCoverageDataByID(db *gorm.DB, pid uint64) (*MitraAOCoverage, error) {
	var err error
	err = db.Debug().Model(&MitraAOCoverage{}).Where("mitra_ao_cov_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraAOCoverage{}, err
	}
	return p, nil
}

func (p *MitraAOCoverage) FindMitraAOCoverageByID(db *gorm.DB, pid uint64) (*MitraAOCoverageData, error) {
	var err error
	mitraaocoverage := MitraAOCoverageData{}
	err = db.Debug().Model(&MitraAOCoverageData{}).Limit(100).
		Select(`m_mitra_ao_coverage.mitra_ao_cov_id,
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
		Joins("JOIN m_mitra on m_mitra_ao_coverage.mitra_id=m_mitra.mitra_id").
		Where("mitra_ao_cov_id = ?", pid).
		Find(&mitraaocoverage).Error
	if err != nil {
		return &MitraAOCoverageData{}, err
	}
	return &mitraaocoverage, nil
}

func (p *MitraAOCoverage) FindMitraAOCoverageStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraAOCoverageData, error) {
	var err error
	mitraaocoverage := MitraAOCoverageData{}
	err = db.Debug().Model(&MitraAOCoverageData{}).Limit(100).
		Select(`m_mitra_ao_coverage.mitra_ao_cov_id,
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
		Joins("JOIN m_mitra on m_mitra_ao_coverage.mitra_id=m_mitra.mitra_id").
		Where("mitra_ao_cov_id = ? AND mitra_ao_cov_status = ?", pid, status).
		Find(&mitraaocoverage).Error
	if err != nil {
		return &MitraAOCoverageData{}, err
	}
	return &mitraaocoverage, nil
}

func (p *MitraAOCoverage) UpdateMitraAOCoverage(db *gorm.DB, pid uint64) (*MitraAOCoverage, error) {
	var err error
	db = db.Debug().Model(&MitraAOCoverage{}).Where("mitra_ao_cov_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_ao_cov_code":       p.MitraAOCoverageCode,
			"mitra_ao_cov_name":       p.MitraAOCoverageName,
			"mitra_ao_cov_status":     p.MitraAOCoverageStatus,
			"mitra_ao_cov_updated_by": p.MitraAOCoverageUpdatedBy,
			"mitra_ao_cov_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraAOCoverage{}).Where("mitra_ao_cov_id = ?", pid).Error
	if err != nil {
		return &MitraAOCoverage{}, err
	}
	return p, nil
}

func (p *MitraAOCoverage) UpdateMitraAOCoverageStatus(db *gorm.DB, pid uint64) (*MitraAOCoverage, error) {
	var err error
	db = db.Debug().Model(&MitraAOCoverage{}).Where("mitra_ao_cov_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_ao_cov_status":     p.MitraAOCoverageStatus,
			"mitra_ao_cov_updated_by": p.MitraAOCoverageUpdatedBy,
			"mitra_ao_cov_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraAOCoverage{}).Where("mitra_ao_cov_id = ?", pid).Error
	if err != nil {
		return &MitraAOCoverage{}, err
	}
	return p, nil
}

func (p *MitraAOCoverage) UpdateMitraAOCoverageStatusOnly(db *gorm.DB, pid uint64) (*MitraAOCoverage, error) {
	var err error
	db = db.Debug().Model(&MitraAOCoverage{}).Where("mitra_ao_cov_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_ao_cov_status": p.MitraAOCoverageStatus,
		},
	)
	err = db.Debug().Model(&MitraAOCoverage{}).Where("mitra_ao_cov_id = ?", pid).Error
	if err != nil {
		return &MitraAOCoverage{}, err
	}
	return p, nil
}

func (p *MitraAOCoverage) UpdateMitraAOCoverageStatusDelete(db *gorm.DB, pid uint64) (*MitraAOCoverage, error) {
	var err error
	db = db.Debug().Model(&MitraAOCoverage{}).Where("mitra_ao_cov_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_ao_cov_status":     3,
			"mitra_ao_cov_deleted_by": p.MitraAOCoverageDeletedBy,
			"mitra_ao_cov_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraAOCoverage{}).Where("mitra_ao_cov_id = ?", pid).Error
	if err != nil {
		return &MitraAOCoverage{}, err
	}
	return p, nil
}

func (p *MitraAOCoverage) DeleteMitraAOCoverage(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraAOCoverage{}).Where("mitra_ao_cov_id = ?", pid).Take(&MitraAOCoverage{}).Delete(&MitraAOCoverage{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraAOCoverage not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//BY MITRA ID
//====================================================================================================================================

func (p *MitraAOCoverage) FindAllMitraAOCoverageMitra(db *gorm.DB, mitra uint64) (*[]MitraAOCoverageData, error) {
	var err error
	mitraaocoverage := []MitraAOCoverageData{}
	err = db.Debug().Model(&MitraAOCoverageData{}).Limit(100).
		Select(`m_mitra_ao_coverage.mitra_ao_cov_id,
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
		Joins("JOIN m_mitra on m_mitra_ao_coverage.mitra_id=m_mitra.mitra_id").
		Where("m_mitra_ao_coverage.mitra_id = ?", mitra).
		Find(&mitraaocoverage).Error
	if err != nil {
		return &[]MitraAOCoverageData{}, err
	}
	return &mitraaocoverage, nil
}

func (p *MitraAOCoverage) FindAllMitraAOCoverageStatusMitra(db *gorm.DB, status uint64, mitra uint64) (*[]MitraAOCoverageData, error) {
	var err error
	mitraaocoverage := []MitraAOCoverageData{}
	err = db.Debug().Model(&MitraAOCoverageData{}).Limit(100).
		Select(`m_mitra_ao_coverage.mitra_ao_cov_id,
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
		Joins("JOIN m_mitra on m_mitra_ao_coverage.mitra_id=m_mitra.mitra_id").
		Where("mitra_ao_cov_status = ? AND m_mitra_ao_coverage.mitra_id = ?", status, mitra).
		Find(&mitraaocoverage).Error
	if err != nil {
		return &[]MitraAOCoverageData{}, err
	}
	return &mitraaocoverage, nil
}

func (p *MitraAOCoverage) FindMitraAOCoverageByIDMitra(db *gorm.DB, pid uint64, mitra uint64) (*MitraAOCoverageData, error) {
	var err error
	mitraaocoverage := MitraAOCoverageData{}
	err = db.Debug().Model(&MitraAOCoverageData{}).Limit(100).
		Select(`m_mitra_ao_coverage.mitra_ao_cov_id,
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
		Joins("JOIN m_mitra on m_mitra_ao_coverage.mitra_id=m_mitra.mitra_id").
		Where("mitra_ao_cov_id = ? AND m_mitra_ao_coverage.mitra_id = ?", pid, mitra).
		Find(&mitraaocoverage).Error
	if err != nil {
		return &MitraAOCoverageData{}, err
	}
	return &mitraaocoverage, nil
}

func (p *MitraAOCoverage) FindMitraAOCoverageStatusByIDMitra(db *gorm.DB, pid uint64, status uint64, mitra uint64) (*MitraAOCoverageData, error) {
	var err error
	mitraaocoverage := MitraAOCoverageData{}
	err = db.Debug().Model(&MitraAOCoverageData{}).Limit(100).
		Select(`m_mitra_ao_coverage.mitra_ao_cov_id,
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
		Joins("JOIN m_mitra on m_mitra_ao_coverage.mitra_id=m_mitra.mitra_id").
		Where("mitra_ao_cov_id = ? AND mitra_ao_cov_status = ? AND m_mitra_ao_coverage.mitra_id = ?", pid, status, mitra).
		Find(&mitraaocoverage).Error
	if err != nil {
		return &MitraAOCoverageData{}, err
	}
	return &mitraaocoverage, nil
}

//ADDITIONAL
//====================================================================================================================================

func (p *MitraAOCoverage) FindMitraAOCoverageByMitraID(db *gorm.DB, pid uint64) (*MitraAOCoverageData, error) {
	var err error
	mitraaocoverage := MitraAOCoverageData{}
	err = db.Debug().Model(&MitraAOCoverageData{}).Limit(100).
		Select(`m_mitra_ao_coverage.mitra_ao_cov_id,
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
		Joins("JOIN m_mitra on m_mitra_ao_coverage.mitra_id=m_mitra.mitra_id").
		Where("m_mitra_ao_coverage.mitra_id = ?", pid).
		Find(&mitraaocoverage).Error
	if err != nil {
		return &MitraAOCoverageData{}, err
	}
	return &mitraaocoverage, nil
}

func (p *MitraAOCoverage) FindMitraAOCoverageStatusByMitraID(db *gorm.DB, pid uint64, status uint64) (*MitraAOCoverageData, error) {
	var err error
	mitraaocoverage := MitraAOCoverageData{}
	err = db.Debug().Model(&MitraAOCoverageData{}).Limit(100).
		Select(`m_mitra_ao_coverage.mitra_ao_cov_id,
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
		Joins("JOIN m_mitra on m_mitra_ao_coverage.mitra_id=m_mitra.mitra_id").
		Where("m_mitra.mitra_id = ? AND m_mitra_ao_coverage.mitra_ao_cov_status = ?", pid, status).
		Find(&mitraaocoverage).Error
	if err != nil {
		return &MitraAOCoverageData{}, err
	}
	return &mitraaocoverage, nil
}
