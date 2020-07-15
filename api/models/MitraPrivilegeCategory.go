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

type MitraPrivilegeCategory struct {
	MitraPrivilegeCategoryID        uint64     `gorm:"column:mitra_priv_ctgry_id;primary_key;" json:"mitra_priv_ctgry_id"`
	MitraPrivilegeCategoryCode      string     `gorm:"column:mitra_priv_ctgry_code;size:25" json:"mitra_priv_ctgry_code"`
	MitraPrivilegeCategoryName      string     `gorm:"column:mitra_priv_ctgry_name;size:255" json:"mitra_priv_ctgry_name"`
	MitraPrivilegeCategoryStatus    uint64     `gorm:"column:mitra_priv_ctgry_status;size:25" json:"mitra_priv_ctgry_status"`
	MitraPrivilegeCategoryCreatedBy string     `gorm:"column:mitra_priv_ctgry_created_by;size:125" json:"mitra_priv_ctgry_created_by"`
	MitraPrivilegeCategoryCreatedAt time.Time  `gorm:"column:mitra_priv_ctgry_created_at;default:CURRENT_TIMESTAMP" json:"mitra_priv_ctgry_created_at"`
	MitraPrivilegeCategoryUpdatedBy string     `gorm:"column:mitra_priv_ctgry_updated_by;size:125" json:"mitra_priv_ctgry_updated_by"`
	MitraPrivilegeCategoryUpdatedAt *time.Time `gorm:"column:mitra_priv_ctgry_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_priv_ctgry_updated_at"`
	MitraPrivilegeCategoryDeletedBy string     `gorm:"column:mitra_priv_ctgry_deleted_by;size:125" json:"mitra_priv_ctgry_deleted_by"`
	MitraPrivilegeCategoryDeletedAt *time.Time `gorm:"column:mitra_priv_ctgry_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_priv_ctgry_deleted_at"`
}

type ResponseMitraPrivilegeCategorys struct {
	Status  int                       `json:"status"`
	Message string                    `json:"message"`
	Data    *[]MitraPrivilegeCategory `json:"data"`
}

type ResponseMitraPrivilegeCategory struct {
	Status  int                     `json:"status"`
	Message string                  `json:"message"`
	Data    *MitraPrivilegeCategory `json:"data"`
}

type ResponseMitraPrivilegeCategoryCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraPrivilegeCategory) TableName() string {
	return "m_mitra_privilege_category"
}

func (p *MitraPrivilegeCategory) Prepare() {
	p.MitraPrivilegeCategoryCode = html.EscapeString(strings.TrimSpace(p.MitraPrivilegeCategoryCode))
	p.MitraPrivilegeCategoryName = html.EscapeString(strings.TrimSpace(p.MitraPrivilegeCategoryName))
	p.MitraPrivilegeCategoryStatus = p.MitraPrivilegeCategoryStatus
	p.MitraPrivilegeCategoryCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraPrivilegeCategoryCreatedBy))
	p.MitraPrivilegeCategoryUpdatedBy = html.EscapeString(strings.TrimSpace(p.MitraPrivilegeCategoryUpdatedBy))
	p.MitraPrivilegeCategoryDeletedBy = html.EscapeString(strings.TrimSpace(p.MitraPrivilegeCategoryDeletedBy))
	p.MitraPrivilegeCategoryCreatedAt = time.Now()
	p.MitraPrivilegeCategoryUpdatedAt = p.MitraPrivilegeCategoryUpdatedAt
	p.MitraPrivilegeCategoryDeletedAt = p.MitraPrivilegeCategoryDeletedAt
}

func (p *MitraPrivilegeCategory) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.MitraPrivilegeCategoryCode == "" {
			return errors.New("Required MitraPrivilegeCategory Code")
		}
		if p.MitraPrivilegeCategoryName == "" {
			return errors.New("Required MitraPrivilegeCategory Name")
		}
		return nil
	}
}

func (p *MitraPrivilegeCategory) SaveMitraPrivilegeCategory(db *gorm.DB) (*MitraPrivilegeCategory, error) {
	p.MitraPrivilegeCategoryStatus = 1
	var err error
	err = db.Debug().Model(&MitraPrivilegeCategory{}).Create(&p).Error
	if err != nil {
		return &MitraPrivilegeCategory{}, err
	}
	return p, nil
}

func (p *MitraPrivilegeCategory) FindAllMitraPrivilegeCategory(db *gorm.DB) (*[]MitraPrivilegeCategory, error) {
	var err error

	mItras := []MitraPrivilegeCategory{}
	err = db.Debug().Model(&MitraPrivilegeCategory{}).Limit(100).
		Select(`m_mitra_privilege_category.mitra_priv_ctgry_id,
				m_mitra_privilege_category.mitra_priv_ctgry_code,
				m_mitra_privilege_category.mitra_priv_ctgry_name,
				m_mitra_privilege_category.mitra_priv_ctgry_status,
				m_mitra_privilege_category.mitra_priv_ctgry_created_by,
				m_mitra_privilege_category.mitra_priv_ctgry_created_at at time zone 'Asia/Jakarta' as mitra_priv_ctgry_created_at,
				m_mitra_privilege_category.mitra_priv_ctgry_updated_by,
				m_mitra_privilege_category.mitra_priv_ctgry_updated_at at time zone 'Asia/Jakarta' as mitra_priv_ctgry_updated_at,
				m_mitra_privilege_category.mitra_priv_ctgry_deleted_by,
				m_mitra_privilege_category.mitra_priv_ctgry_deleted_at at time zone 'Asia/Jakarta' as mitra_priv_ctgry_deleted_at
			`).
		Find(&mItras).Error
	if err != nil {
		return &[]MitraPrivilegeCategory{}, err
	}
	return &mItras, nil
}

func (p *MitraPrivilegeCategory) FindAllMitraPrivilegeCategoryActive(db *gorm.DB) (*[]MitraPrivilegeCategory, error) {
	var err error

	mItras := []MitraPrivilegeCategory{}
	err = db.Debug().Model(&MitraPrivilegeCategory{}).
		Select(`m_mitra_privilege_category.mitra_priv_ctgry_id,
				m_mitra_privilege_category.mitra_priv_ctgry_code,
				m_mitra_privilege_category.mitra_priv_ctgry_name,
				m_mitra_privilege_category.mitra_priv_ctgry_status,
				m_mitra_privilege_category.mitra_priv_ctgry_created_by,
				m_mitra_privilege_category.mitra_priv_ctgry_created_at at time zone 'Asia/Jakarta' as mitra_priv_ctgry_created_at,
				m_mitra_privilege_category.mitra_priv_ctgry_updated_by,
				m_mitra_privilege_category.mitra_priv_ctgry_updated_at at time zone 'Asia/Jakarta' as mitra_priv_ctgry_updated_at,
				m_mitra_privilege_category.mitra_priv_ctgry_deleted_by,
				m_mitra_privilege_category.mitra_priv_ctgry_deleted_at at time zone 'Asia/Jakarta' as mitra_priv_ctgry_deleted_at
			`).
		Where("mitra_priv_ctgry_status = ?", 1).Limit(100).Find(&mItras).Error
	if err != nil {
		return &[]MitraPrivilegeCategory{}, err
	}
	return &mItras, nil
}

func (p *MitraPrivilegeCategory) FindAllMitraPrivilegeCategoryDeleted(db *gorm.DB) (*[]MitraPrivilegeCategory, error) {
	var err error

	mItras := []MitraPrivilegeCategory{}
	err = db.Debug().Model(&MitraPrivilegeCategory{}).
		Select(`m_mitra_privilege_category.mitra_priv_ctgry_id,
				m_mitra_privilege_category.mitra_priv_ctgry_code,
				m_mitra_privilege_category.mitra_priv_ctgry_name,
				m_mitra_privilege_category.mitra_priv_ctgry_status,
				m_mitra_privilege_category.mitra_priv_ctgry_created_by,
				m_mitra_privilege_category.mitra_priv_ctgry_created_at at time zone 'Asia/Jakarta' as mitra_priv_ctgry_created_at,
				m_mitra_privilege_category.mitra_priv_ctgry_updated_by,
				m_mitra_privilege_category.mitra_priv_ctgry_updated_at at time zone 'Asia/Jakarta' as mitra_priv_ctgry_updated_at,
				m_mitra_privilege_category.mitra_priv_ctgry_deleted_by,
				m_mitra_privilege_category.mitra_priv_ctgry_deleted_at at time zone 'Asia/Jakarta' as mitra_priv_ctgry_deleted_at
			`).
		Where("mitra_priv_ctgry_status = ?", 3).Limit(100).Find(&mItras).Error
	if err != nil {
		return &[]MitraPrivilegeCategory{}, err
	}
	return &mItras, nil
}

func (p *MitraPrivilegeCategory) FindMitraPrivilegeCategoryByID(db *gorm.DB, pid uint64) (*MitraPrivilegeCategory, error) {
	var err error
	err = db.Debug().Model(&MitraPrivilegeCategory{}).
		Select(`m_mitra_privilege_category.mitra_priv_ctgry_id,
				m_mitra_privilege_category.mitra_priv_ctgry_code,
				m_mitra_privilege_category.mitra_priv_ctgry_name,
				m_mitra_privilege_category.mitra_priv_ctgry_status,
				m_mitra_privilege_category.mitra_priv_ctgry_created_by,
				m_mitra_privilege_category.mitra_priv_ctgry_created_at at time zone 'Asia/Jakarta' as mitra_priv_ctgry_created_at,
				m_mitra_privilege_category.mitra_priv_ctgry_updated_by,
				m_mitra_privilege_category.mitra_priv_ctgry_updated_at at time zone 'Asia/Jakarta' as mitra_priv_ctgry_updated_at,
				m_mitra_privilege_category.mitra_priv_ctgry_deleted_by,
				m_mitra_privilege_category.mitra_priv_ctgry_deleted_at at time zone 'Asia/Jakarta' as mitra_priv_ctgry_deleted_at
			`).
		Where("mitra_priv_ctgry_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraPrivilegeCategory{}, err
	}
	return p, nil
}

func (p *MitraPrivilegeCategory) FindMitraPrivilegeCategoryActiveByID(db *gorm.DB, pid uint64) (*MitraPrivilegeCategory, error) {
	var err error
	err = db.Debug().Model(&MitraPrivilegeCategory{}).
		Select(`m_mitra_privilege_category.mitra_priv_ctgry_id,
				m_mitra_privilege_category.mitra_priv_ctgry_code,
				m_mitra_privilege_category.mitra_priv_ctgry_name,
				m_mitra_privilege_category.mitra_priv_ctgry_status,
				m_mitra_privilege_category.mitra_priv_ctgry_created_by,
				m_mitra_privilege_category.mitra_priv_ctgry_created_at at time zone 'Asia/Jakarta' as mitra_priv_ctgry_created_at,
				m_mitra_privilege_category.mitra_priv_ctgry_updated_by,
				m_mitra_privilege_category.mitra_priv_ctgry_updated_at at time zone 'Asia/Jakarta' as mitra_priv_ctgry_updated_at,
				m_mitra_privilege_category.mitra_priv_ctgry_deleted_by,
				m_mitra_privilege_category.mitra_priv_ctgry_deleted_at at time zone 'Asia/Jakarta' as mitra_priv_ctgry_deleted_at
			`).
		Where("mitra_priv_ctgry_id = ? AND mitra_priv_ctgry_status = ?", pid, 1).Take(&p).Error
	if err != nil {
		return &MitraPrivilegeCategory{}, err
	}
	return p, nil
}

func (p *MitraPrivilegeCategory) FindMitraPrivilegeCategoryDeletedByID(db *gorm.DB, pid uint64) (*MitraPrivilegeCategory, error) {
	var err error
	err = db.Debug().Model(&MitraPrivilegeCategory{}).
		Select(`m_mitra_privilege_category.mitra_priv_ctgry_id,
				m_mitra_privilege_category.mitra_priv_ctgry_code,
				m_mitra_privilege_category.mitra_priv_ctgry_name,
				m_mitra_privilege_category.mitra_priv_ctgry_status,
				m_mitra_privilege_category.mitra_priv_ctgry_created_by,
				m_mitra_privilege_category.mitra_priv_ctgry_created_at at time zone 'Asia/Jakarta' as mitra_priv_ctgry_created_at,
				m_mitra_privilege_category.mitra_priv_ctgry_updated_by,
				m_mitra_privilege_category.mitra_priv_ctgry_updated_at at time zone 'Asia/Jakarta' as mitra_priv_ctgry_updated_at,
				m_mitra_privilege_category.mitra_priv_ctgry_deleted_by,
				m_mitra_privilege_category.mitra_priv_ctgry_deleted_at at time zone 'Asia/Jakarta' as mitra_priv_ctgry_deleted_at
			`).
		Where("mitra_priv_ctgry_id = ? AND mitra_priv_ctgry_status = ?", pid, 3).Take(&p).Error
	if err != nil {
		return &MitraPrivilegeCategory{}, err
	}
	return p, nil
}

func (p *MitraPrivilegeCategory) UpdateMitraPrivilegeCategory(db *gorm.DB, pid uint64) (*MitraPrivilegeCategory, error) {

	var err error
	db = db.Debug().Model(&MitraPrivilegeCategory{}).Where("mitra_priv_ctgry_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_priv_ctgry_code":       p.MitraPrivilegeCategoryCode,
			"mitra_priv_ctgry_name":       p.MitraPrivilegeCategoryName,
			"mitra_priv_ctgry_updated_by": p.MitraPrivilegeCategoryUpdatedBy,
			"mitra_priv_ctgry_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraPrivilegeCategory{}).Where("mitra_priv_ctgry_id = ?", pid).Error
	if err != nil {
		return &MitraPrivilegeCategory{}, err
	}
	return p, nil
}

func (p *MitraPrivilegeCategory) UpdateMitraPrivilegeCategoryStatus(db *gorm.DB, pid uint64) (*MitraPrivilegeCategory, error) {

	var err error
	db = db.Debug().Model(&MitraPrivilegeCategory{}).Where("mitra_priv_ctgry_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_priv_ctgry_status": p.MitraPrivilegeCategoryStatus,
		},
	)
	err = db.Debug().Model(&MitraPrivilegeCategory{}).Where("mitra_priv_ctgry_id = ?", pid).Error
	if err != nil {
		return &MitraPrivilegeCategory{}, err
	}
	return p, nil
}

func (p *MitraPrivilegeCategory) UpdateMitraPrivilegeCategoryDeleted(db *gorm.DB, pid uint64) (*MitraPrivilegeCategory, error) {

	var err error
	db = db.Debug().Model(&MitraPrivilegeCategory{}).Where("mitra_priv_ctgry_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_priv_ctgry_status":     3,
			"mitra_priv_ctgry_deleted_by": p.MitraPrivilegeCategoryDeletedBy,
			"mitra_priv_ctgry_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraPrivilegeCategory{}).Where("mitra_priv_ctgry_id = ?", pid).Error
	if err != nil {
		return &MitraPrivilegeCategory{}, err
	}
	return p, nil
}

func (p *MitraPrivilegeCategory) DeleteMitraPrivilegeCategory(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&MitraPrivilegeCategory{}).Where("mitra_priv_ctgry_id = ?", pid).Take(&MitraPrivilegeCategory{}).Delete(&MitraPrivilegeCategory{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraPrivilegeCategory not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================
