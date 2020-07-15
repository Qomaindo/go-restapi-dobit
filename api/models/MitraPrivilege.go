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

type MitraPrivilege struct {
	MitraPrivilegeID         uint64     `gorm:"column:mitra_priv_id;primary_key;" json:"mitra_priv_id"`
	MitraPrivilegeCategoryID uint64     `gorm:"column:mitra_priv_ctgry_id" json:"mitra_priv_ctgry_id"`
	MitraPrivilegeCode       string     `gorm:"column:mitra_priv_code;size:25" json:"mitra_priv_code"`
	MitraPrivilegeName       string     `gorm:"column:mitra_priv_name;size:255" json:"mitra_priv_name"`
	MitraPrivilegeStatus     uint64     `gorm:"column:mitra_priv_status;size:25" json:"mitra_priv_status"`
	MitraPrivilegeCreatedBy  string     `gorm:"column:mitra_priv_created_by;size:125" json:"mitra_priv_created_by"`
	MitraPrivilegeCreatedAt  time.Time  `gorm:"column:mitra_priv_created_at;default:CURRENT_TIMESTAMP" json:"mitra_priv_created_at"`
	MitraPrivilegeUpdatedBy  string     `gorm:"column:mitra_priv_updated_by;size:125" json:"mitra_priv_updated_by"`
	MitraPrivilegeUpdatedAt  *time.Time `gorm:"column:mitra_priv_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_priv_updated_at"`
	MitraPrivilegeDeletedBy  string     `gorm:"column:mitra_priv_deleted_by;size:125" json:"mitra_priv_deleted_by"`
	MitraPrivilegeDeletedAt  *time.Time `gorm:"column:mitra_priv_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_priv_deleted_at"`
}

type MitraPrivilegeData struct {
	MitraPrivilegeID           uint64     `gorm:"column:mitra_priv_id" json:"mitra_priv_id"`
	MitraPrivilegeCategoryID   uint64     `gorm:"column:mitra_priv_ctgry_id" json:"mitra_priv_ctgry_id"`
	MitraPrivilegeCategoryCode string     `gorm:"column:mitra_priv_ctgry_code;size:10" json:"mitra_priv_ctgry_code"`
	MitraPrivilegeCategoryName string     `gorm:"column:mitra_priv_ctgry_name;size:255" json:"mitra_priv_ctgry_name"`
	MitraPrivilegeCode         string     `gorm:"column:mitra_priv_code;size:25" json:"mitra_priv_code"`
	MitraPrivilegeName         string     `gorm:"column:mitra_priv_name;size:255" json:"mitra_priv_name"`
	MitraPrivilegeStatus       uint64     `gorm:"column:mitra_priv_status;size:2" json:"mitra_priv_status"`
	MitraPrivilegeCreatedBy    string     `gorm:"column:mitra_priv_created_by;size:125" json:"mitra_priv_created_by"`
	MitraPrivilegeCreatedAt    time.Time  `gorm:"column:mitra_priv_created_at;default:CURRENT_TIMESTAMP" json:"mitra_priv_created_at"`
	MitraPrivilegeUpdatedBy    string     `gorm:"column:mitra_priv_updated_by;size:125" json:"mitra_priv_updated_by"`
	MitraPrivilegeUpdatedAt    *time.Time `gorm:"column:mitra_priv_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_priv_updated_at"`
	MitraPrivilegeDeletedBy    string     `gorm:"column:mitra_priv_deleted_by;size:125" json:"mitra_priv_deleted_by"`
	MitraPrivilegeDeletedAt    *time.Time `gorm:"column:mitra_priv_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_priv_deleted_at"`
}

type ResponseMitraPrivileges struct {
	Status  int                   `json:"status"`
	Message string                `json:"message"`
	Data    *[]MitraPrivilegeData `json:"data"`
}

type ResponseMitraPrivilege struct {
	Status  int                 `json:"status"`
	Message string              `json:"message"`
	Data    *MitraPrivilegeData `json:"data"`
}

type ResponseMitraPrivilegeCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraPrivilege) TableName() string {
	return "m_mitra_privilege"
}

func (MitraPrivilegeData) TableName() string {
	return "m_mitra_privilege"
}

func (p *MitraPrivilege) Prepare() {
	p.MitraPrivilegeCategoryID = p.MitraPrivilegeCategoryID
	p.MitraPrivilegeCode = html.EscapeString(strings.TrimSpace(p.MitraPrivilegeCode))
	p.MitraPrivilegeName = html.EscapeString(strings.TrimSpace(p.MitraPrivilegeName))
	p.MitraPrivilegeStatus = p.MitraPrivilegeStatus
	p.MitraPrivilegeCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraPrivilegeCreatedBy))
	p.MitraPrivilegeUpdatedBy = html.EscapeString(strings.TrimSpace(p.MitraPrivilegeUpdatedBy))
	p.MitraPrivilegeDeletedBy = html.EscapeString(strings.TrimSpace(p.MitraPrivilegeDeletedBy))
	p.MitraPrivilegeCreatedAt = time.Now()
	p.MitraPrivilegeUpdatedAt = p.MitraPrivilegeUpdatedAt
	p.MitraPrivilegeDeletedAt = p.MitraPrivilegeDeletedAt
}

func (p *MitraPrivilege) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.MitraPrivilegeCategoryID == 0 {
			return errors.New("Required MitraPrivilegeCategory ID")
		}
		if p.MitraPrivilegeCode == "" {
			return errors.New("Required MitraPrivilege Code")
		}
		if p.MitraPrivilegeName == "" {
			return errors.New("Required MitraPrivilege Name")
		}
		return nil
	}
}

func (p *MitraPrivilege) SaveMitraPrivilege(db *gorm.DB) (*MitraPrivilege, error) {
	p.MitraPrivilegeStatus = 1
	var err error
	err = db.Debug().Model(&MitraPrivilege{}).Create(&p).Error
	if err != nil {
		return &MitraPrivilege{}, err
	}
	return p, nil
}

func (p *MitraPrivilege) SaveMitraPrivilegePending(db *gorm.DB) (*MitraPrivilege, error) {

	p.MitraPrivilegeStatus = 0

	var err error
	err = db.Debug().Model(&MitraPrivilege{}).Create(&p).Error
	if err != nil {
		return &MitraPrivilege{}, err
	}
	return p, nil
}

func (p *MitraPrivilege) FindAllMitraPrivileges(db *gorm.DB) (*[]MitraPrivilegeData, error) {
	var err error
	mitracontactperson := []MitraPrivilegeData{}
	err = db.Debug().Model(&MitraPrivilegeData{}).Limit(100).
		Select(`m_mitra_privilege.mitra_priv_id,
				m_mitra_privilege_category.mitra_priv_ctgry_id,
				m_mitra_privilege_category.mitra_priv_ctgry_code,
				m_mitra_privilege_category.mitra_priv_ctgry_name,
				m_mitra_privilege.mitra_priv_code,
				m_mitra_privilege.mitra_priv_name,
				m_mitra_privilege.mitra_priv_status,
				m_mitra_privilege.mitra_priv_created_by,
				m_mitra_privilege.mitra_priv_updated_by,
				m_mitra_privilege.mitra_priv_deleted_by,
				m_mitra_privilege.mitra_priv_created_at at time zone 'Asia/Jakarta' as mitra_priv_created_at,
				m_mitra_privilege.mitra_priv_updated_at at time zone 'Asia/Jakarta' as mitra_priv_updated_at,
				m_mitra_privilege.mitra_priv_deleted_at at time zone 'Asia/Jakarta' as mitra_priv_deleted_at
				`).
		Joins("JOIN m_mitra_privilege_category on m_mitra_privilege.mitra_priv_ctgry_id=m_mitra_privilege_category.mitra_priv_ctgry_id").
		Find(&mitracontactperson).Error
	if err != nil {
		return &[]MitraPrivilegeData{}, err
	}
	return &mitracontactperson, nil
}

func (p *MitraPrivilege) FindAllMitraPrivilegesActive(db *gorm.DB) (*[]MitraPrivilegeData, error) {
	var err error
	mitracontactperson := []MitraPrivilegeData{}
	err = db.Debug().Model(&MitraPrivilegeData{}).Limit(100).
		Select(`m_mitra_privilege.mitra_priv_id,
				m_mitra_privilege_category.mitra_priv_ctgry_id,
				m_mitra_privilege_category.mitra_priv_ctgry_code,
				m_mitra_privilege_category.mitra_priv_ctgry_name,
				m_mitra_privilege.mitra_priv_code,
				m_mitra_privilege.mitra_priv_name,
				m_mitra_privilege.mitra_priv_status,
				m_mitra_privilege.mitra_priv_created_by,
				m_mitra_privilege.mitra_priv_updated_by,
				m_mitra_privilege.mitra_priv_deleted_by,
				m_mitra_privilege.mitra_priv_created_at at time zone 'Asia/Jakarta' as mitra_priv_created_at,
				m_mitra_privilege.mitra_priv_updated_at at time zone 'Asia/Jakarta' as mitra_priv_updated_at,
				m_mitra_privilege.mitra_priv_deleted_at at time zone 'Asia/Jakarta' as mitra_priv_deleted_at
				`).
		Joins("JOIN m_mitra_privilege_category on m_mitra_privilege.mitra_priv_ctgry_id=m_mitra_privilege_category.mitra_priv_ctgry_id").
		Where("mitra_priv_status = ?", 1).
		Find(&mitracontactperson).Error
	if err != nil {
		return &[]MitraPrivilegeData{}, err
	}
	return &mitracontactperson, nil
}

func (p *MitraPrivilege) FindAllMitraPrivilegesDeleted(db *gorm.DB) (*[]MitraPrivilegeData, error) {
	var err error
	mitracontactperson := []MitraPrivilegeData{}
	err = db.Debug().Model(&MitraPrivilegeData{}).Limit(100).
		Select(`m_mitra_privilege.mitra_priv_id,
				m_mitra_privilege_category.mitra_priv_ctgry_id,
				m_mitra_privilege_category.mitra_priv_ctgry_code,
				m_mitra_privilege_category.mitra_priv_ctgry_name,
				m_mitra_privilege.mitra_priv_code,
				m_mitra_privilege.mitra_priv_name,
				m_mitra_privilege.mitra_priv_status,
				m_mitra_privilege.mitra_priv_created_by,
				m_mitra_privilege.mitra_priv_updated_by,
				m_mitra_privilege.mitra_priv_deleted_by,
				m_mitra_privilege.mitra_priv_created_at at time zone 'Asia/Jakarta' as mitra_priv_created_at,
				m_mitra_privilege.mitra_priv_updated_at at time zone 'Asia/Jakarta' as mitra_priv_updated_at,
				m_mitra_privilege.mitra_priv_deleted_at at time zone 'Asia/Jakarta' as mitra_priv_deleted_at
				`).
		Joins("JOIN m_mitra_privilege_category on m_mitra_privilege.mitra_priv_ctgry_id=m_mitra_privilege_category.mitra_priv_ctgry_id").
		Where("mitra_priv_status = ?", 3).
		Find(&mitracontactperson).Error
	if err != nil {
		return &[]MitraPrivilegeData{}, err
	}
	return &mitracontactperson, nil
}

func (p *MitraPrivilege) FindMitraPrivilegeByID(db *gorm.DB, pid uint64) (*MitraPrivilegeData, error) {
	var err error
	mitracontactperson := MitraPrivilegeData{}
	err = db.Debug().Model(&MitraPrivilegeData{}).Limit(100).
		Select(`m_mitra_privilege.mitra_priv_id,
				m_mitra_privilege_category.mitra_priv_ctgry_id,
				m_mitra_privilege_category.mitra_priv_ctgry_code,
				m_mitra_privilege_category.mitra_priv_ctgry_name,
				m_mitra_privilege.mitra_priv_code,
				m_mitra_privilege.mitra_priv_name,
				m_mitra_privilege.mitra_priv_status,
				m_mitra_privilege.mitra_priv_created_by,
				m_mitra_privilege.mitra_priv_updated_by,
				m_mitra_privilege.mitra_priv_deleted_by,
				m_mitra_privilege.mitra_priv_created_at at time zone 'Asia/Jakarta' as mitra_priv_created_at,
				m_mitra_privilege.mitra_priv_updated_at at time zone 'Asia/Jakarta' as mitra_priv_updated_at,
				m_mitra_privilege.mitra_priv_deleted_at at time zone 'Asia/Jakarta' as mitra_priv_deleted_at
				`).
		Joins("JOIN m_mitra_privilege_category on m_mitra_privilege.mitra_priv_ctgry_id=m_mitra_privilege_category.mitra_priv_ctgry_id").
		Where("mitra_priv_id = ?", pid).
		Find(&mitracontactperson).Error
	if err != nil {
		return &MitraPrivilegeData{}, err
	}
	return &mitracontactperson, nil
}

func (p *MitraPrivilege) FindMitraPrivilegeActiveByID(db *gorm.DB, pid uint64) (*MitraPrivilegeData, error) {
	var err error
	mitracontactperson := MitraPrivilegeData{}
	err = db.Debug().Model(&MitraPrivilegeData{}).Limit(100).
		Select(`m_mitra_privilege.mitra_priv_id,
				m_mitra_privilege_category.mitra_priv_ctgry_id,
				m_mitra_privilege_category.mitra_priv_ctgry_code,
				m_mitra_privilege_category.mitra_priv_ctgry_name,
				m_mitra_privilege.mitra_priv_code,
				m_mitra_privilege.mitra_priv_name,
				m_mitra_privilege.mitra_priv_status,
				m_mitra_privilege.mitra_priv_created_by,
				m_mitra_privilege.mitra_priv_updated_by,
				m_mitra_privilege.mitra_priv_deleted_by,
				m_mitra_privilege.mitra_priv_created_at at time zone 'Asia/Jakarta' as mitra_priv_created_at,
				m_mitra_privilege.mitra_priv_updated_at at time zone 'Asia/Jakarta' as mitra_priv_updated_at,
				m_mitra_privilege.mitra_priv_deleted_at at time zone 'Asia/Jakarta' as mitra_priv_deleted_at
				`).
		Joins("JOIN m_mitra_privilege_category on m_mitra_privilege.mitra_priv_ctgry_id=m_mitra_privilege_category.mitra_priv_ctgry_id").
		Where("mitra_priv_id = ? AND mitra_priv_status = ?", pid, 1).
		Find(&mitracontactperson).Error
	if err != nil {
		return &MitraPrivilegeData{}, err
	}
	return &mitracontactperson, nil
}

func (p *MitraPrivilege) FindMitraPrivilegeDeletedByID(db *gorm.DB, pid uint64) (*MitraPrivilegeData, error) {
	var err error
	mitracontactperson := MitraPrivilegeData{}
	err = db.Debug().Model(&MitraPrivilegeData{}).Limit(100).
		Select(`m_mitra_privilege.mitra_priv_id,
				m_mitra_privilege_category.mitra_priv_ctgry_id,
				m_mitra_privilege_category.mitra_priv_ctgry_code,
				m_mitra_privilege_category.mitra_priv_ctgry_name,
				m_mitra_privilege.mitra_priv_code,
				m_mitra_privilege.mitra_priv_name,
				m_mitra_privilege.mitra_priv_status,
				m_mitra_privilege.mitra_priv_created_by,
				m_mitra_privilege.mitra_priv_updated_by,
				m_mitra_privilege.mitra_priv_deleted_by,
				m_mitra_privilege.mitra_priv_created_at at time zone 'Asia/Jakarta' as mitra_priv_created_at,
				m_mitra_privilege.mitra_priv_updated_at at time zone 'Asia/Jakarta' as mitra_priv_updated_at,
				m_mitra_privilege.mitra_priv_deleted_at at time zone 'Asia/Jakarta' as mitra_priv_deleted_at
				`).
		Joins("JOIN m_mitra_privilege_category on m_mitra_privilege.mitra_priv_ctgry_id=m_mitra_privilege_category.mitra_priv_ctgry_id").
		Where("mitra_priv_id = ? AND mitra_priv_status = ?", pid, 3).
		Find(&mitracontactperson).Error
	if err != nil {
		return &MitraPrivilegeData{}, err
	}
	return &mitracontactperson, nil
}

func (p *MitraPrivilege) UpdateMitraPrivilege(db *gorm.DB, pid uint64) (*MitraPrivilege, error) {

	var err error
	db = db.Debug().Model(&MitraPrivilege{}).Where("mitra_priv_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_priv_ctgry_id":   p.MitraPrivilegeCategoryID,
			"mitra_priv_code":       p.MitraPrivilegeCode,
			"mitra_priv_name":       p.MitraPrivilegeName,
			"mitra_priv_updated_by": p.MitraPrivilegeUpdatedBy,
			"mitra_priv_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraPrivilege{}).Where("mitra_priv_id = ?", pid).Error
	if err != nil {
		return &MitraPrivilege{}, err
	}
	return p, nil
}

func (p *MitraPrivilege) UpdateMitraPrivilegeStatus(db *gorm.DB, pid uint64) (*MitraPrivilege, error) {

	var err error
	db = db.Debug().Model(&MitraPrivilege{}).Where("mitra_priv_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_priv_status": p.MitraPrivilegeStatus,
		},
	)
	err = db.Debug().Model(&MitraPrivilege{}).Where("mitra_priv_id = ?", pid).Error
	if err != nil {
		return &MitraPrivilege{}, err
	}
	return p, nil
}

func (p *MitraPrivilege) UpdateMitraPrivilegeDeleted(db *gorm.DB, pid uint64) (*MitraPrivilege, error) {

	var err error
	db = db.Debug().Model(&MitraPrivilege{}).Where("mitra_priv_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_priv_status":     3,
			"mitra_priv_deleted_by": p.MitraPrivilegeDeletedBy,
			"mitra_priv_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraPrivilege{}).Where("mitra_priv_id = ?", pid).Error
	if err != nil {
		return &MitraPrivilege{}, err
	}
	return p, nil
}

func (p *MitraPrivilege) DeleteMitraPrivilege(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&MitraPrivilege{}).Where("mitra_priv_id = ?", pid).Take(&MitraPrivilege{}).Delete(&MitraPrivilege{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraPrivilege not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================
