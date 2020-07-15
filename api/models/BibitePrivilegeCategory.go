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

type BibitePrivilegeCategory struct {
	BibitePrivilegeCategoryID        uint64     `gorm:"column:bbt_priv_ctgry_id;primary_key;" json:"bbt_priv_ctgry_id"`
	BibitePrivilegeCategoryCode      string     `gorm:"column:bbt_priv_ctgry_code;size:10" json:"bbt_priv_ctgry_code"`
	BibitePrivilegeCategoryName      string     `gorm:"column:bbt_priv_ctgry_name;size:255" json:"bbt_priv_ctgry_name"`
	BibitePrivilegeCategoryStatus    uint64     `gorm:"column:bbt_priv_ctgry_status;size:1" json:"bbt_priv_ctgry_status"`
	BibitePrivilegeCategoryCreatedBy string     `gorm:"column:bbt_priv_ctgry_created_by;size:125" json:"bbt_priv_ctgry_created_by"`
	BibitePrivilegeCategoryCreatedAt time.Time  `gorm:"column:bbt_priv_ctgry_created_at;default:CURRENT_TIMESTAMP" json:"bbt_priv_ctgry_created_at"`
	BibitePrivilegeCategoryUpdatedBy string     `gorm:"column:bbt_priv_ctgry_updated_by;size:125" json:"bbt_priv_ctgry_updated_by"`
	BibitePrivilegeCategoryUpdatedAt *time.Time `gorm:"column:bbt_priv_ctgry_updated_at;default:CURRENT_TIMESTAMP" json:"bbt_priv_ctgry_updated_at"`
	BibitePrivilegeCategoryDeletedBy string     `gorm:"column:bbt_priv_ctgry_deleted_by;size:125" json:"bbt_priv_ctgry_deleted_by"`
	BibitePrivilegeCategoryDeletedAt *time.Time `gorm:"column:bbt_priv_ctgry_deleted_at;default:CURRENT_TIMESTAMP" json:"bbt_priv_ctgry_deleted_at"`
}

type ResponseBibitePrivilegeCategorys struct {
	Status  int                        `json:"status"`
	Message string                     `json:"message"`
	Data    *[]BibitePrivilegeCategory `json:"data"`
}

type ResponseBibitePrivilegeCategory struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    *BibitePrivilegeCategory `json:"data"`
}

type ResponseBibitePrivilegeCategoryCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (BibitePrivilegeCategory) TableName() string {
	return "m_bibite_privilege_category"
}

func (p *BibitePrivilegeCategory) Prepare() {
	p.BibitePrivilegeCategoryCode = html.EscapeString(strings.TrimSpace(p.BibitePrivilegeCategoryCode))
	p.BibitePrivilegeCategoryName = html.EscapeString(strings.TrimSpace(p.BibitePrivilegeCategoryName))
	p.BibitePrivilegeCategoryStatus = p.BibitePrivilegeCategoryStatus
	p.BibitePrivilegeCategoryCreatedBy = p.BibitePrivilegeCategoryCreatedBy
	p.BibitePrivilegeCategoryCreatedAt = time.Now()
	p.BibitePrivilegeCategoryUpdatedBy = p.BibitePrivilegeCategoryUpdatedBy
	p.BibitePrivilegeCategoryUpdatedAt = p.BibitePrivilegeCategoryUpdatedAt
	p.BibitePrivilegeCategoryDeletedBy = p.BibitePrivilegeCategoryDeletedBy
	p.BibitePrivilegeCategoryDeletedAt = p.BibitePrivilegeCategoryDeletedAt
}

func (p *BibitePrivilegeCategory) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.BibitePrivilegeCategoryCode == "" {
			return errors.New("Required BibitePrivilegeCategory Code")
		}
		if p.BibitePrivilegeCategoryName == "" {
			return errors.New("Required BibitePrivilegeCategory Name")
		}
		return nil
	}
}

func (p *BibitePrivilegeCategory) SaveBibitePrivilegeCategory(db *gorm.DB) (*BibitePrivilegeCategory, error) {
	p.BibitePrivilegeCategoryStatus = 1
	var err error
	err = db.Debug().Model(&BibitePrivilegeCategory{}).Create(&p).Error
	if err != nil {
		return &BibitePrivilegeCategory{}, err
	}
	return p, nil
}

func (p *BibitePrivilegeCategory) SaveBibitePrivilegeCategoryPending(db *gorm.DB) (*BibitePrivilegeCategory, error) {

	p.BibitePrivilegeCategoryStatus = 0

	var err error
	err = db.Debug().Model(&BibitePrivilegeCategory{}).Create(&p).Error
	if err != nil {
		return &BibitePrivilegeCategory{}, err
	}
	return p, nil
}

func (p *BibitePrivilegeCategory) FindAllBibitePrivilegeCategory(db *gorm.DB) (*[]BibitePrivilegeCategory, error) {
	var err error

	mItras := []BibitePrivilegeCategory{}
	err = db.Debug().Model(&BibitePrivilegeCategory{}).Limit(100).
		Select(`m_bibite_privilege_category.bbt_priv_ctgry_id,
				m_bibite_privilege_category.bbt_priv_ctgry_code,
				m_bibite_privilege_category.bbt_priv_ctgry_name,
				m_bibite_privilege_category.bbt_priv_ctgry_status,
				m_bibite_privilege_category.bbt_priv_ctgry_created_by,
				m_bibite_privilege_category.bbt_priv_ctgry_created_at at time zone 'Asia/Jakarta' as bbt_priv_ctgry_created_at,
				m_bibite_privilege_category.bbt_priv_ctgry_updated_by,
				m_bibite_privilege_category.bbt_priv_ctgry_updated_at at time zone 'Asia/Jakarta' as bbt_priv_ctgry_updated_at,
				m_bibite_privilege_category.bbt_priv_ctgry_deleted_by,
				m_bibite_privilege_category.bbt_priv_ctgry_deleted_at at time zone 'Asia/Jakarta' as bbt_priv_ctgry_deleted_at
			`).
		Find(&mItras).Error
	if err != nil {
		return &[]BibitePrivilegeCategory{}, err
	}
	return &mItras, nil
}

func (p *BibitePrivilegeCategory) FindAllBibiteGroupPrivilegeCategoryByID(db *gorm.DB, id uint64) (*[]BibitePrivilegeCategory, error) {
	var err error
	bibitegorupprivilege := []BibitePrivilegeCategory{}
	err = db.Debug().Model(&BibitePrivilegeCategory{}).Limit(100).
		Select(`m_bibite_privilege_category.bbt_priv_ctgry_id,
				m_bibite_privilege_category.bbt_priv_ctgry_code,
				m_bibite_privilege_category.bbt_priv_ctgry_name`).
		Find(&bibitegorupprivilege).Error
	if err != nil {
		return &[]BibitePrivilegeCategory{}, err
	}
	return &bibitegorupprivilege, nil
}

func (p *BibitePrivilegeCategory) FindAllBibitePrivilegeCategoryActive(db *gorm.DB) (*[]BibitePrivilegeCategory, error) {
	var err error

	mItras := []BibitePrivilegeCategory{}
	err = db.Debug().Model(&BibitePrivilegeCategory{}).
		Select(`m_bibite_privilege_category.bbt_priv_ctgry_id,
				m_bibite_privilege_category.bbt_priv_ctgry_code,
				m_bibite_privilege_category.bbt_priv_ctgry_name,
				m_bibite_privilege_category.bbt_priv_ctgry_status,
				m_bibite_privilege_category.bbt_priv_ctgry_created_by,
				m_bibite_privilege_category.bbt_priv_ctgry_created_at at time zone 'Asia/Jakarta' as bbt_priv_ctgry_created_at,
				m_bibite_privilege_category.bbt_priv_ctgry_updated_by,
				m_bibite_privilege_category.bbt_priv_ctgry_updated_at at time zone 'Asia/Jakarta' as bbt_priv_ctgry_updated_at,
				m_bibite_privilege_category.bbt_priv_ctgry_deleted_by,
				m_bibite_privilege_category.bbt_priv_ctgry_deleted_at at time zone 'Asia/Jakarta' as bbt_priv_ctgry_deleted_at
			`).
		Where("bbt_priv_ctgry_status = ?", 1).Limit(100).Find(&mItras).Error
	if err != nil {
		return &[]BibitePrivilegeCategory{}, err
	}
	return &mItras, nil
}

func (p *BibitePrivilegeCategory) FindAllBibitePrivilegeCategoryPending(db *gorm.DB) (*[]BibitePrivilegeCategory, error) {
	var err error

	mItras := []BibitePrivilegeCategory{}
	err = db.Debug().Model(&BibitePrivilegeCategory{}).
		Select(`m_bibite_privilege_category.bbt_priv_ctgry_id,
				m_bibite_privilege_category.bbt_priv_ctgry_code,
				m_bibite_privilege_category.bbt_priv_ctgry_name,
				m_bibite_privilege_category.bbt_priv_ctgry_status,
				m_bibite_privilege_category.bbt_priv_ctgry_created_by,
				m_bibite_privilege_category.bbt_priv_ctgry_created_at at time zone 'Asia/Jakarta' as bbt_priv_ctgry_created_at,
				m_bibite_privilege_category.bbt_priv_ctgry_updated_by,
				m_bibite_privilege_category.bbt_priv_ctgry_updated_at at time zone 'Asia/Jakarta' as bbt_priv_ctgry_updated_at,
				m_bibite_privilege_category.bbt_priv_ctgry_deleted_by,
				m_bibite_privilege_category.bbt_priv_ctgry_deleted_at at time zone 'Asia/Jakarta' as bbt_priv_ctgry_deleted_at
			`).
		Where("bbt_priv_ctgry_status = ?", 0).Limit(100).Find(&mItras).Error
	if err != nil {
		return &[]BibitePrivilegeCategory{}, err
	}
	return &mItras, nil
}

func (p *BibitePrivilegeCategory) FindAllBibitePrivilegeCategoryDeleted(db *gorm.DB) (*[]BibitePrivilegeCategory, error) {
	var err error

	mItras := []BibitePrivilegeCategory{}
	err = db.Debug().Model(&BibitePrivilegeCategory{}).
		Select(`m_bibite_privilege_category.bbt_priv_ctgry_id,
				m_bibite_privilege_category.bbt_priv_ctgry_code,
				m_bibite_privilege_category.bbt_priv_ctgry_name,
				m_bibite_privilege_category.bbt_priv_ctgry_status,
				m_bibite_privilege_category.bbt_priv_ctgry_created_by,
				m_bibite_privilege_category.bbt_priv_ctgry_created_at at time zone 'Asia/Jakarta' as bbt_priv_ctgry_created_at,
				m_bibite_privilege_category.bbt_priv_ctgry_updated_by,
				m_bibite_privilege_category.bbt_priv_ctgry_updated_at at time zone 'Asia/Jakarta' as bbt_priv_ctgry_updated_at,
				m_bibite_privilege_category.bbt_priv_ctgry_deleted_by,
				m_bibite_privilege_category.bbt_priv_ctgry_deleted_at at time zone 'Asia/Jakarta' as bbt_priv_ctgry_deleted_at
			`).
		Where("bbt_priv_ctgry_status = ?", 3).Limit(100).Find(&mItras).Error
	if err != nil {
		return &[]BibitePrivilegeCategory{}, err
	}
	return &mItras, nil
}

func (p *BibitePrivilegeCategory) FindBibitePrivilegeCategoryByID(db *gorm.DB, pid uint64) (*BibitePrivilegeCategory, error) {
	var err error
	err = db.Debug().Model(&BibitePrivilegeCategory{}).
		Select(`m_bibite_privilege_category.bbt_priv_ctgry_id,
				m_bibite_privilege_category.bbt_priv_ctgry_code,
				m_bibite_privilege_category.bbt_priv_ctgry_name,
				m_bibite_privilege_category.bbt_priv_ctgry_status,
				m_bibite_privilege_category.bbt_priv_ctgry_created_by,
				m_bibite_privilege_category.bbt_priv_ctgry_created_at at time zone 'Asia/Jakarta' as bbt_priv_ctgry_created_at,
				m_bibite_privilege_category.bbt_priv_ctgry_updated_by,
				m_bibite_privilege_category.bbt_priv_ctgry_updated_at at time zone 'Asia/Jakarta' as bbt_priv_ctgry_updated_at,
				m_bibite_privilege_category.bbt_priv_ctgry_deleted_by,
				m_bibite_privilege_category.bbt_priv_ctgry_deleted_at at time zone 'Asia/Jakarta' as bbt_priv_ctgry_deleted_at
			`).
		Where("bbt_priv_ctgry_id = ?", pid).Take(&p).Error
	if err != nil {
		return &BibitePrivilegeCategory{}, err
	}
	return p, nil
}

func (p *BibitePrivilegeCategory) FindBibitePrivilegeCategoryActiveByID(db *gorm.DB, pid uint64) (*BibitePrivilegeCategory, error) {
	var err error
	err = db.Debug().Model(&BibitePrivilegeCategory{}).
		Select(`m_bibite_privilege_category.bbt_priv_ctgry_id,
				m_bibite_privilege_category.bbt_priv_ctgry_code,
				m_bibite_privilege_category.bbt_priv_ctgry_name,
				m_bibite_privilege_category.bbt_priv_ctgry_status,
				m_bibite_privilege_category.bbt_priv_ctgry_created_by,
				m_bibite_privilege_category.bbt_priv_ctgry_created_at at time zone 'Asia/Jakarta' as bbt_priv_ctgry_created_at,
				m_bibite_privilege_category.bbt_priv_ctgry_updated_by,
				m_bibite_privilege_category.bbt_priv_ctgry_updated_at at time zone 'Asia/Jakarta' as bbt_priv_ctgry_updated_at,
				m_bibite_privilege_category.bbt_priv_ctgry_deleted_by,
				m_bibite_privilege_category.bbt_priv_ctgry_deleted_at at time zone 'Asia/Jakarta' as bbt_priv_ctgry_deleted_at
			`).
		Where("bbt_priv_ctgry_id = ? AND bbt_priv_ctgry_status = ?", pid, 1).Take(&p).Error
	if err != nil {
		return &BibitePrivilegeCategory{}, err
	}
	return p, nil
}

func (p *BibitePrivilegeCategory) FindBibitePrivilegeCategoryPendingByID(db *gorm.DB, pid uint64) (*BibitePrivilegeCategory, error) {
	var err error
	err = db.Debug().Model(&BibitePrivilegeCategory{}).
		Select(`m_bibite_privilege_category.bbt_priv_ctgry_id,
				m_bibite_privilege_category.bbt_priv_ctgry_code,
				m_bibite_privilege_category.bbt_priv_ctgry_name,
				m_bibite_privilege_category.bbt_priv_ctgry_status,
				m_bibite_privilege_category.bbt_priv_ctgry_created_by,
				m_bibite_privilege_category.bbt_priv_ctgry_created_at at time zone 'Asia/Jakarta' as bbt_priv_ctgry_created_at,
				m_bibite_privilege_category.bbt_priv_ctgry_updated_by,
				m_bibite_privilege_category.bbt_priv_ctgry_updated_at at time zone 'Asia/Jakarta' as bbt_priv_ctgry_updated_at,
				m_bibite_privilege_category.bbt_priv_ctgry_deleted_by,
				m_bibite_privilege_category.bbt_priv_ctgry_deleted_at at time zone 'Asia/Jakarta' as bbt_priv_ctgry_deleted_at
			`).
		Where("bbt_priv_ctgry_id = ? AND bbt_priv_ctgry_status = ?", pid, 0).Take(&p).Error
	if err != nil {
		return &BibitePrivilegeCategory{}, err
	}
	return p, nil
}

func (p *BibitePrivilegeCategory) FindBibitePrivilegeCategoryDeletedByID(db *gorm.DB, pid uint64) (*BibitePrivilegeCategory, error) {
	var err error
	err = db.Debug().Model(&BibitePrivilegeCategory{}).
		Select(`m_bibite_privilege_category.bbt_priv_ctgry_id,
				m_bibite_privilege_category.bbt_priv_ctgry_code,
				m_bibite_privilege_category.bbt_priv_ctgry_name,
				m_bibite_privilege_category.bbt_priv_ctgry_status,
				m_bibite_privilege_category.bbt_priv_ctgry_created_by,
				m_bibite_privilege_category.bbt_priv_ctgry_created_at at time zone 'Asia/Jakarta' as bbt_priv_ctgry_created_at,
				m_bibite_privilege_category.bbt_priv_ctgry_updated_by,
				m_bibite_privilege_category.bbt_priv_ctgry_updated_at at time zone 'Asia/Jakarta' as bbt_priv_ctgry_updated_at,
				m_bibite_privilege_category.bbt_priv_ctgry_deleted_by,
				m_bibite_privilege_category.bbt_priv_ctgry_deleted_at at time zone 'Asia/Jakarta' as bbt_priv_ctgry_deleted_at
			`).
		Where("bbt_priv_ctgry_id = ? AND bbt_priv_ctgry_status = ?", pid, 3).Take(&p).Error
	if err != nil {
		return &BibitePrivilegeCategory{}, err
	}
	return p, nil
}

func (p *BibitePrivilegeCategory) UpdateBibitePrivilegeCategory(db *gorm.DB, pid uint64) (*BibitePrivilegeCategory, error) {

	var err error
	db = db.Debug().Model(&BibitePrivilegeCategory{}).Where("bbt_priv_ctgry_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_priv_ctgry_code":       p.BibitePrivilegeCategoryCode,
			"bbt_priv_ctgry_name":       p.BibitePrivilegeCategoryName,
			"bbt_priv_ctgry_updated_by": p.BibitePrivilegeCategoryUpdatedBy,
			"bbt_priv_ctgry_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&BibitePrivilegeCategory{}).Where("bbt_priv_ctgry_id = ?", pid).Error
	if err != nil {
		return &BibitePrivilegeCategory{}, err
	}
	return p, nil
}

func (p *BibitePrivilegeCategory) UpdateBibitePrivilegeCategoryStatus(db *gorm.DB, pid uint64) (*BibitePrivilegeCategory, error) {

	var err error
	db = db.Debug().Model(&BibitePrivilegeCategory{}).Where("bbt_priv_ctgry_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_priv_ctgry_status": p.BibitePrivilegeCategoryStatus,
		},
	)
	err = db.Debug().Model(&BibitePrivilegeCategory{}).Where("bbt_priv_ctgry_id = ?", pid).Error
	if err != nil {
		return &BibitePrivilegeCategory{}, err
	}
	return p, nil
}

func (p *BibitePrivilegeCategory) UpdateBibitePrivilegeCategoryDeleted(db *gorm.DB, pid uint64) (*BibitePrivilegeCategory, error) {

	var err error
	db = db.Debug().Model(&BibitePrivilegeCategory{}).Where("bbt_priv_ctgry_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_priv_ctgry_status":     3,
			"bbt_priv_ctgry_deleted_by": p.BibitePrivilegeCategoryDeletedBy,
			"bbt_priv_ctgry_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&BibitePrivilegeCategory{}).Where("bbt_priv_ctgry_id = ?", pid).Error
	if err != nil {
		return &BibitePrivilegeCategory{}, err
	}
	return p, nil
}

func (p *BibitePrivilegeCategory) DeleteBibitePrivilegeCategory(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&BibitePrivilegeCategory{}).Where("bbt_priv_ctgry_id = ?", pid).Take(&BibitePrivilegeCategory{}).Delete(&BibitePrivilegeCategory{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("BibitePrivilegeCategory not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================
