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

type BibitePrivilege struct {
	BibitePrivilegeID         uint64     `gorm:"column:bbt_priv_id;primary_key;" json:"bbt_priv_id"`
	BibitePrivilegeCategoryID uint64     `gorm:"column:bbt_priv_ctgry_id" json:"bbt_priv_ctgry_id"`
	BibitePrivilegeCode       string     `gorm:"column:bbt_priv_code;size:10" json:"bbt_priv_code"`
	BibitePrivilegeName       string     `gorm:"column:bbt_priv_name;size:255" json:"bbt_priv_name"`
	BibitePrivilegeStatus     uint64     `gorm:"column:bbt_priv_status;size:1" json:"bbt_priv_status"`
	BibitePrivilegeCreatedBy  string     `gorm:"column:bbt_priv_created_by;size:125" json:"bbt_priv_created_by"`
	BibitePrivilegeCreatedAt  time.Time  `gorm:"column:bbt_priv_created_at;default:CURRENT_TIMESTAMP" json:"bbt_priv_created_at"`
	BibitePrivilegeUpdatedBy  string     `gorm:"column:bbt_priv_updated_by;size:125" json:"bbt_priv_updated_by"`
	BibitePrivilegeUpdatedAt  *time.Time `gorm:"column:bbt_priv_updated_at;default:CURRENT_TIMESTAMP" json:"bbt_priv_updated_at"`
	BibitePrivilegeDeletedBy  string     `gorm:"column:bbt_priv_deleted_by;size:125" json:"bbt_priv_deleted_by"`
	BibitePrivilegeDeletedAt  *time.Time `gorm:"column:bbt_priv_deleted_at;default:CURRENT_TIMESTAMP" json:"bbt_priv_deleted_at"`
}

type BibitePrivilegeData struct {
	BibitePrivilegeID           uint64     `gorm:"column:bbt_priv_id;primary_key;" json:"bbt_priv_id"`
	BibitePrivilegeCategoryID   uint64     `gorm:"column:bbt_priv_ctgry_id" json:"bbt_priv_ctgry_id"`
	BibitePrivilegeCategoryCode string     `gorm:"column:bbt_priv_ctgry_code" json:"bbt_priv_ctgry_code"`
	BibitePrivilegeCategoryName string     `gorm:"column:bbt_priv_ctgry_name" json:"bbt_priv_ctgry_name"`
	BibitePrivilegeCode         string     `gorm:"column:bbt_priv_code;size:10" json:"bbt_priv_code"`
	BibitePrivilegeName         string     `gorm:"column:bbt_priv_name;size:255" json:"bbt_priv_name"`
	BibitePrivilegeStatus       uint64     `gorm:"column:bbt_priv_status;size:1" json:"bbt_priv_status"`
	BibitePrivilegeCreatedBy    string     `gorm:"column:bbt_priv_created_by;size:125" json:"bbt_priv_created_by"`
	BibitePrivilegeCreatedAt    time.Time  `gorm:"column:bbt_priv_created_at;default:CURRENT_TIMESTAMP" json:"bbt_priv_created_at"`
	BibitePrivilegeUpdatedBy    string     `gorm:"column:bbt_priv_updated_by;size:125" json:"bbt_priv_updated_by"`
	BibitePrivilegeUpdatedAt    *time.Time `gorm:"column:bbt_priv_updated_at;default:CURRENT_TIMESTAMP" json:"bbt_priv_updated_at"`
	BibitePrivilegeDeletedBy    string     `gorm:"column:bbt_priv_deleted_by;size:125" json:"bbt_priv_deleted_by"`
	BibitePrivilegeDeletedAt    *time.Time `gorm:"column:bbt_priv_deleted_at;default:CURRENT_TIMESTAMP" json:"bbt_priv_deleted_at"`
}

type ResponseBibitePrivileges struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    *[]BibitePrivilegeData `json:"data"`
}

type ResponseBibitePrivilege struct {
	Status  int                  `json:"status"`
	Message string               `json:"message"`
	Data    *BibitePrivilegeData `json:"data"`
}

type ResponseBibitePrivilegeCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (BibitePrivilege) TableName() string {
	return "m_bibite_privilege"
}

func (BibitePrivilegeData) TableName() string {
	return "m_bibite_privilege"
}

func (p *BibitePrivilege) Prepare() {
	p.BibitePrivilegeCategoryID = p.BibitePrivilegeCategoryID
	p.BibitePrivilegeCode = html.EscapeString(strings.TrimSpace(p.BibitePrivilegeCode))
	p.BibitePrivilegeName = html.EscapeString(strings.TrimSpace(p.BibitePrivilegeName))
	p.BibitePrivilegeStatus = p.BibitePrivilegeStatus
	p.BibitePrivilegeCreatedBy = p.BibitePrivilegeCreatedBy
	p.BibitePrivilegeCreatedAt = time.Now()
	p.BibitePrivilegeUpdatedBy = p.BibitePrivilegeUpdatedBy
	p.BibitePrivilegeUpdatedAt = p.BibitePrivilegeUpdatedAt
	p.BibitePrivilegeDeletedBy = p.BibitePrivilegeDeletedBy
	p.BibitePrivilegeDeletedAt = p.BibitePrivilegeDeletedAt
}

func (p *BibitePrivilege) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.BibitePrivilegeCategoryID == 0 {
			return errors.New("Required BibitePrivilegeCategory ID")
		}
		if p.BibitePrivilegeCode == "" {
			return errors.New("Required BibitePrivilege Code")
		}
		if p.BibitePrivilegeName == "" {
			return errors.New("Required BibitePrivilege Name")
		}
		return nil
	}
}

func (p *BibitePrivilege) SaveBibitePrivilege(db *gorm.DB) (*BibitePrivilege, error) {
	p.BibitePrivilegeStatus = 1
	var err error
	err = db.Debug().Model(&BibitePrivilege{}).Create(&p).Error
	if err != nil {
		return &BibitePrivilege{}, err
	}
	return p, nil
}

func (p *BibitePrivilege) SaveBibitePrivilegePending(db *gorm.DB) (*BibitePrivilege, error) {

	p.BibitePrivilegeStatus = 0

	var err error
	err = db.Debug().Model(&BibitePrivilege{}).Create(&p).Error
	if err != nil {
		return &BibitePrivilege{}, err
	}
	return p, nil
}

func (p *BibitePrivilege) FindAllBibitePrivileges(db *gorm.DB) (*[]BibitePrivilegeData, error) {
	var err error
	mitracontactperson := []BibitePrivilegeData{}
	err = db.Debug().Model(&BibitePrivilegeData{}).Limit(100).
		Select(`m_bibite_privilege.bbt_priv_id,
				m_bibite_privilege_category.bbt_priv_ctgry_id,
				m_bibite_privilege_category.bbt_priv_ctgry_code,
				m_bibite_privilege_category.bbt_priv_ctgry_name,
				m_bibite_privilege.bbt_priv_code,
				m_bibite_privilege.bbt_priv_name,
				m_bibite_privilege.bbt_priv_status,
				m_bibite_privilege.bbt_priv_created_by,
				m_bibite_privilege.bbt_priv_created_at at time zone 'Asia/Jakarta' as bbt_priv_created_at,
				m_bibite_privilege.bbt_priv_updated_by,
				m_bibite_privilege.bbt_priv_updated_at at time zone 'Asia/Jakarta' as bbt_priv_updated_at,
				m_bibite_privilege.bbt_priv_deleted_by,
				m_bibite_privilege.bbt_priv_deleted_at at time zone 'Asia/Jakarta' as bbt_priv_deleted_at
				`).
		Joins("JOIN m_bibite_privilege_category on m_bibite_privilege.bbt_priv_ctgry_id=m_bibite_privilege_category.bbt_priv_ctgry_id").
		Find(&mitracontactperson).Error
	if err != nil {
		return &[]BibitePrivilegeData{}, err
	}
	return &mitracontactperson, nil
}

func (p *BibitePrivilege) FindAllBibitePrivilegesActive(db *gorm.DB) (*[]BibitePrivilegeData, error) {
	var err error
	mitracontactperson := []BibitePrivilegeData{}
	err = db.Debug().Model(&BibitePrivilegeData{}).Limit(100).
		Select(`m_bibite_privilege.bbt_priv_id,
				m_bibite_privilege_category.bbt_priv_ctgry_id,
				m_bibite_privilege_category.bbt_priv_ctgry_code,
				m_bibite_privilege_category.bbt_priv_ctgry_name,
				m_bibite_privilege.bbt_priv_code,
				m_bibite_privilege.bbt_priv_name,
				m_bibite_privilege.bbt_priv_status,
				m_bibite_privilege.bbt_priv_created_by,
				m_bibite_privilege.bbt_priv_created_at at time zone 'Asia/Jakarta' as bbt_priv_created_at,
				m_bibite_privilege.bbt_priv_updated_by,
				m_bibite_privilege.bbt_priv_updated_at at time zone 'Asia/Jakarta' as bbt_priv_updated_at,
				m_bibite_privilege.bbt_priv_deleted_by,
				m_bibite_privilege.bbt_priv_deleted_at at time zone 'Asia/Jakarta' as bbt_priv_deleted_at
				`).
		Joins("JOIN m_bibite_privilege_category on m_bibite_privilege.bbt_priv_ctgry_id=m_bibite_privilege_category.bbt_priv_ctgry_id").
		Where("bbt_priv_status = ?", 1).
		Find(&mitracontactperson).Error
	if err != nil {
		return &[]BibitePrivilegeData{}, err
	}
	return &mitracontactperson, nil
}

func (p *BibitePrivilege) FindAllBibitePrivilegesDeleted(db *gorm.DB) (*[]BibitePrivilegeData, error) {
	var err error
	mitracontactperson := []BibitePrivilegeData{}
	err = db.Debug().Model(&BibitePrivilegeData{}).Limit(100).
		Select(`m_bibite_privilege.bbt_priv_id,
				m_bibite_privilege_category.bbt_priv_ctgry_id,
				m_bibite_privilege_category.bbt_priv_ctgry_code,
				m_bibite_privilege_category.bbt_priv_ctgry_name,
				m_bibite_privilege.bbt_priv_code,
				m_bibite_privilege.bbt_priv_name,
				m_bibite_privilege.bbt_priv_status,
				m_bibite_privilege.bbt_priv_created_by,
				m_bibite_privilege.bbt_priv_created_at at time zone 'Asia/Jakarta' as bbt_priv_created_at,
				m_bibite_privilege.bbt_priv_updated_by,
				m_bibite_privilege.bbt_priv_updated_at at time zone 'Asia/Jakarta' as bbt_priv_updated_at,
				m_bibite_privilege.bbt_priv_deleted_by,
				m_bibite_privilege.bbt_priv_deleted_at at time zone 'Asia/Jakarta' as bbt_priv_deleted_at
				`).
		Joins("JOIN m_bibite_privilege_category on m_bibite_privilege.bbt_priv_ctgry_id=m_bibite_privilege_category.bbt_priv_ctgry_id").
		Where("bbt_priv_status = ?", 3).
		Find(&mitracontactperson).Error
	if err != nil {
		return &[]BibitePrivilegeData{}, err
	}
	return &mitracontactperson, nil
}

func (p *BibitePrivilege) FindBibitePrivilegeByID(db *gorm.DB, pid uint64) (*BibitePrivilegeData, error) {
	var err error
	mitracontactperson := BibitePrivilegeData{}
	err = db.Debug().Model(&BibitePrivilegeData{}).Limit(100).
		Select(`m_bibite_privilege.bbt_priv_id,
				m_bibite_privilege_category.bbt_priv_ctgry_id,
				m_bibite_privilege_category.bbt_priv_ctgry_code,
				m_bibite_privilege_category.bbt_priv_ctgry_name,
				m_bibite_privilege.bbt_priv_code,
				m_bibite_privilege.bbt_priv_name,
				m_bibite_privilege.bbt_priv_status,
				m_bibite_privilege.bbt_priv_created_by,
				m_bibite_privilege.bbt_priv_created_at at time zone 'Asia/Jakarta' as bbt_priv_created_at,
				m_bibite_privilege.bbt_priv_updated_by,
				m_bibite_privilege.bbt_priv_updated_at at time zone 'Asia/Jakarta' as bbt_priv_updated_at,
				m_bibite_privilege.bbt_priv_deleted_by,
				m_bibite_privilege.bbt_priv_deleted_at at time zone 'Asia/Jakarta' as bbt_priv_deleted_at
				`).
		Joins("JOIN m_bibite_privilege_category on m_bibite_privilege.bbt_priv_ctgry_id=m_bibite_privilege_category.bbt_priv_ctgry_id").
		Where("bbt_priv_id = ?", pid).
		Find(&mitracontactperson).Error
	if err != nil {
		return &BibitePrivilegeData{}, err
	}
	return &mitracontactperson, nil
}

func (p *BibitePrivilege) FindBibitePrivilegeActiveByID(db *gorm.DB, pid uint64) (*BibitePrivilegeData, error) {
	var err error
	mitracontactperson := BibitePrivilegeData{}
	err = db.Debug().Model(&BibitePrivilegeData{}).Limit(100).
		Select(`m_bibite_privilege.bbt_priv_id,
				m_bibite_privilege_category.bbt_priv_ctgry_id,
				m_bibite_privilege_category.bbt_priv_ctgry_code,
				m_bibite_privilege_category.bbt_priv_ctgry_name,
				m_bibite_privilege.bbt_priv_code,
				m_bibite_privilege.bbt_priv_name,
				m_bibite_privilege.bbt_priv_status,
				m_bibite_privilege.bbt_priv_created_by,
				m_bibite_privilege.bbt_priv_created_at at time zone 'Asia/Jakarta' as bbt_priv_created_at,
				m_bibite_privilege.bbt_priv_updated_by,
				m_bibite_privilege.bbt_priv_updated_at at time zone 'Asia/Jakarta' as bbt_priv_updated_at,
				m_bibite_privilege.bbt_priv_deleted_by,
				m_bibite_privilege.bbt_priv_deleted_at at time zone 'Asia/Jakarta' as bbt_priv_deleted_at
				`).
		Joins("JOIN m_bibite_privilege_category on m_bibite_privilege.bbt_priv_ctgry_id=m_bibite_privilege_category.bbt_priv_ctgry_id").
		Where("bbt_priv_id = ? AND bbt_priv_status = ?", pid, 1).
		Find(&mitracontactperson).Error
	if err != nil {
		return &BibitePrivilegeData{}, err
	}
	return &mitracontactperson, nil
}

func (p *BibitePrivilege) FindBibitePrivilegeDeletedByID(db *gorm.DB, pid uint64) (*BibitePrivilegeData, error) {
	var err error
	mitracontactperson := BibitePrivilegeData{}
	err = db.Debug().Model(&BibitePrivilegeData{}).Limit(100).
		Select(`m_bibite_privilege.bbt_priv_id,
				m_bibite_privilege_category.bbt_priv_ctgry_id,
				m_bibite_privilege_category.bbt_priv_ctgry_code,
				m_bibite_privilege_category.bbt_priv_ctgry_name,
				m_bibite_privilege.bbt_priv_code,
				m_bibite_privilege.bbt_priv_name,
				m_bibite_privilege.bbt_priv_status,
				m_bibite_privilege.bbt_priv_created_by,
				m_bibite_privilege.bbt_priv_created_at at time zone 'Asia/Jakarta' as bbt_priv_created_at,
				m_bibite_privilege.bbt_priv_updated_by,
				m_bibite_privilege.bbt_priv_updated_at at time zone 'Asia/Jakarta' as bbt_priv_updated_at,
				m_bibite_privilege.bbt_priv_deleted_by,
				m_bibite_privilege.bbt_priv_deleted_at at time zone 'Asia/Jakarta' as bbt_priv_deleted_at
				`).
		Joins("JOIN m_bibite_privilege_category on m_bibite_privilege.bbt_priv_ctgry_id=m_bibite_privilege_category.bbt_priv_ctgry_id").
		Where("bbt_priv_id = ? AND bbt_priv_status = ?", pid, 3).
		Find(&mitracontactperson).Error
	if err != nil {
		return &BibitePrivilegeData{}, err
	}
	return &mitracontactperson, nil
}

func (p *BibitePrivilege) UpdateBibitePrivilege(db *gorm.DB, pid uint64) (*BibitePrivilege, error) {

	var err error
	db = db.Debug().Model(&BibitePrivilege{}).Where("bbt_priv_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_priv_ctgry_id":   p.BibitePrivilegeCategoryID,
			"bbt_priv_code":       p.BibitePrivilegeCode,
			"bbt_priv_name":       p.BibitePrivilegeName,
			"bbt_priv_updated_by": p.BibitePrivilegeUpdatedBy,
			"bbt_priv_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&BibitePrivilege{}).Where("bbt_priv_id = ?", pid).Error
	if err != nil {
		return &BibitePrivilege{}, err
	}
	return p, nil
}

func (p *BibitePrivilege) UpdateBibitePrivilegeStatus(db *gorm.DB, pid uint64) (*BibitePrivilege, error) {

	var err error
	db = db.Debug().Model(&BibitePrivilege{}).Where("bbt_priv_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_priv_status": p.BibitePrivilegeStatus,
		},
	)
	err = db.Debug().Model(&BibitePrivilege{}).Where("bbt_priv_id = ?", pid).Error
	if err != nil {
		return &BibitePrivilege{}, err
	}
	return p, nil
}

func (p *BibitePrivilege) UpdateBibitePrivilegeDeleted(db *gorm.DB, pid uint64) (*BibitePrivilege, error) {

	var err error
	db = db.Debug().Model(&BibitePrivilege{}).Where("bbt_priv_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_priv_status":     3,
			"bbt_priv_deleted_by": p.BibitePrivilegeDeletedBy,
			"bbt_priv_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&BibitePrivilege{}).Where("bbt_priv_id = ?", pid).Error
	if err != nil {
		return &BibitePrivilege{}, err
	}
	return p, nil
}

func (p *BibitePrivilege) DeleteBibitePrivilege(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&BibitePrivilege{}).Where("bbt_priv_id = ?", pid).Take(&BibitePrivilege{}).Delete(&BibitePrivilege{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("BibitePrivilege not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================
