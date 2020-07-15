package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type MitraGroupAccess struct {
	MitraGroupAccessID        uint64     `gorm:"column:mitra_group_accs_id;primary_key;" json:"mitra_group_accs_id"`
	MitraID                   uint64     `gorm:"column:mitra_id" json:"mitra_id"`
	MitraGroupAccessCode      string     `gorm:"column:mitra_group_accs_code;size:25" json:"mitra_group_accs_code"`
	MitraGroupAccessName      string     `gorm:"column:mitra_group_accs_name;size:255" json:"mitra_group_accs_name"`
	MitraGroupAccessRole      string     `gorm:"column:mitra_group_accs_role;size:25" json:"mitra_group_accs_role"`
	MitraGroupAccessType      int64      `gorm:"column:mitra_group_accs_type;size:2" json:"mitra_group_accs_type"`
	MitraGroupAccessStatus    uint64     `gorm:"column:mitra_group_accs_status;size:2" json:"mitra_group_accs_status"`
	MitraGroupAccessCreatedBy string     `gorm:"column:mitra_group_accs_created_by;size:125" json:"mitra_group_accs_created_by"`
	MitraGroupAccessUpdatedBy string     `gorm:"column:mitra_group_accs_updated_by;size:125" json:"mitra_group_accs_updated_by"`
	MitraGroupAccessDeletedBy string     `gorm:"column:mitra_group_accs_deleted_by;size:125" json:"mitra_group_accs_deleted_by"`
	MitraGroupAccessCreatedAt time.Time  `gorm:"column:mitra_group_accs_created_at;default:CURRENT_TIMESTAMP" json:"mitra_group_accs_created_at"`
	MitraGroupAccessUpdatedAt *time.Time `gorm:"column:mitra_group_accs_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_group_accs_updated_at"`
	MitraGroupAccessDeletedAt *time.Time `gorm:"column:mitra_group_accs_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_group_accs_deleted_at"`
}

type MitraGroupAccessData struct {
	MitraGroupAccessID        uint64                          `gorm:"column:mitra_group_accs_id" json:"mitra_group_accs_id"`
	MitraID                   uint64                          `gorm:"column:mitra_id" json:"mitra_id"`
	MitraCode                 string                          `gorm:"column:mitra_code" json:"mitra_code"`
	MitraName                 string                          `gorm:"column:mitra_name" json:"mitra_name"`
	MitraWebsite              string                          `gorm:"column:mitra_website" json:"mitra_website"`
	MitraEmail                string                          `gorm:"column:mitra_email" json:"mitra_email"`
	MitraLogo                 string                          `gorm:"column:mitra_logo" json:"mitra_logo"`
	MitraGroupAccessCode      string                          `gorm:"column:mitra_group_accs_code;size:25" json:"mitra_group_accs_code"`
	MitraGroupAccessName      string                          `gorm:"column:mitra_group_accs_name;size:255" json:"mitra_group_accs_name"`
	MitraGroupAccessRole      string                          `gorm:"column:mitra_group_accs_role;size:25" json:"mitra_group_accs_role"`
	MitraGroupAccessType      int64                           `gorm:"column:mitra_group_accs_type;size:2" json:"mitra_group_accs_type"`
	MitraGroupAccessStatus    uint64                          `gorm:"column:mitra_group_accs_status;size:2" json:"mitra_group_accs_status"`
	MitraGroupAccessCreatedBy string                          `gorm:"column:mitra_group_accs_created_by;size:125" json:"mitra_group_accs_created_by"`
	MitraGroupAccessCreatedAt time.Time                       `gorm:"column:mitra_group_accs_created_at;default:CURRENT_TIMESTAMP" json:"mitra_group_accs_created_at"`
	MitraGroupAccessUpdatedBy string                          `gorm:"column:mitra_group_accs_updated_by;size:125" json:"mitra_group_accs_updated_by"`
	MitraGroupAccessUpdatedAt *time.Time                      `gorm:"column:mitra_group_accs_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_group_accs_updated_at"`
	MitraGroupAccessDeletedBy string                          `gorm:"column:mitra_group_accs_deleted_by;size:125" json:"mitra_group_accs_deleted_by"`
	MitraGroupAccessDeletedAt *time.Time                      `gorm:"column:mitra_group_accs_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_group_accs_deleted_at"`
	MitraGroupAccessPrivilege []MitraGroupAccessPrivilegeData `json:"mitra_group_accs_privilege"`
}

type ResponseMitraGroupAccesss struct {
	Status  int                     `json:"status"`
	Message string                  `json:"message"`
	Data    *[]MitraGroupAccessData `json:"data"`
}

type ResponseMitraGroupAccess struct {
	Status  int                   `json:"status"`
	Message string                `json:"message"`
	Data    *MitraGroupAccessData `json:"data"`
}

type ResponseMitraGroupAccessCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraGroupAccess) TableName() string {
	return "m_mitra_group_access"
}

func (MitraGroupAccessData) TableName() string {
	return "m_mitra_group_access"
}

func (p *MitraGroupAccess) Prepare() {
	p.MitraID = p.MitraID
	p.MitraGroupAccessCode = p.MitraGroupAccessCode
	p.MitraGroupAccessName = p.MitraGroupAccessName
	p.MitraGroupAccessRole = p.MitraGroupAccessRole
	p.MitraGroupAccessType = p.MitraGroupAccessType
	p.MitraGroupAccessStatus = p.MitraGroupAccessStatus
	p.MitraGroupAccessCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraGroupAccessCreatedBy))
	p.MitraGroupAccessCreatedAt = time.Now()
	p.MitraGroupAccessUpdatedBy = html.EscapeString(strings.TrimSpace(p.MitraGroupAccessUpdatedBy))
	p.MitraGroupAccessUpdatedAt = p.MitraGroupAccessUpdatedAt
	p.MitraGroupAccessDeletedBy = html.EscapeString(strings.TrimSpace(p.MitraGroupAccessDeletedBy))
	p.MitraGroupAccessDeletedAt = p.MitraGroupAccessDeletedAt
}

func (p *MitraGroupAccess) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.MitraGroupAccessCode == "" {
			return errors.New("Required MitraGroupAccess Code")
		}
		if p.MitraGroupAccessName == "" {
			return errors.New("Required MitraGroupAccess Name")
		}
		if p.MitraGroupAccessRole == "" {
			return errors.New("Required MitraGroupAccess Name")
		}
		if p.MitraGroupAccessType == 0 {
			return errors.New("Required MitraGroupAccess Name")
		}
		return nil
	}
}

func (p *MitraGroupAccess) SaveMitraGroupAccess(db *gorm.DB) (*MitraGroupAccess, error) {
	var err error
	err = db.Debug().Model(&MitraGroupAccess{}).Create(&p).Error
	if err != nil {
		return &MitraGroupAccess{}, err
	}
	return p, nil
}

func (p *MitraGroupAccess) FindAllMitraGroupAccess(db *gorm.DB, mitra uint64) (*[]MitraGroupAccessData, error) {
	var err error
	mitragroup := []MitraGroupAccessData{}
	err = db.Debug().Model(&MitraGroupAccessData{}).Limit(100).
		Select(`m_mitra_group_access.mitra_group_accs_id,
				m_mitra_group_access.mitra_id,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra.mitra_website,
				m_mitra.mitra_email,
				m_mitra.mitra_logo,
				m_mitra_group_access.mitra_group_accs_code,
				m_mitra_group_access.mitra_group_accs_name,
				m_mitra_group_access.mitra_group_accs_role,
				m_mitra_group_access.mitra_group_accs_type,
				m_mitra_group_access.mitra_group_accs_status,
				m_mitra_group_access.mitra_group_accs_created_by,
				m_mitra_group_access.mitra_group_accs_updated_by,
				m_mitra_group_access.mitra_group_accs_deleted_by,
				m_mitra_group_access.mitra_group_accs_created_at,
				m_mitra_group_access.mitra_group_accs_updated_at,
				m_mitra_group_access.mitra_group_accs_deleted_at`).
		Joins("JOIN m_mitra on m_mitra_group_access.mitra_id=m_mitra.mitra_id").
		Where("m_mitra_group_access.mitra_id = ?", mitra).
		Find(&mitragroup).Error
	if err != nil {
		return &[]MitraGroupAccessData{}, err
	}
	return &mitragroup, nil
}

func (p *MitraGroupAccess) FindAllMitraGroupAccessStatus(db *gorm.DB, status uint64, mitra uint64) (*[]MitraGroupAccessData, error) {
	var err error
	mitragroup := []MitraGroupAccessData{}
	err = db.Debug().Model(&MitraGroupAccessData{}).Limit(100).
		Select(`m_mitra_group_access.mitra_group_accs_id,
				m_mitra_group_access.mitra_id,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra.mitra_website,
				m_mitra.mitra_email,
				m_mitra.mitra_logo,
				m_mitra_group_access.mitra_group_accs_code,
				m_mitra_group_access.mitra_group_accs_name,
				m_mitra_group_access.mitra_group_accs_role,
				m_mitra_group_access.mitra_group_accs_type,
				m_mitra_group_access.mitra_group_accs_status,
				m_mitra_group_access.mitra_group_accs_created_by,
				m_mitra_group_access.mitra_group_accs_updated_by,
				m_mitra_group_access.mitra_group_accs_deleted_by,
				m_mitra_group_access.mitra_group_accs_created_at,
				m_mitra_group_access.mitra_group_accs_updated_at,
				m_mitra_group_access.mitra_group_accs_deleted_at`).
		Joins("JOIN m_mitra on m_mitra_group_access.mitra_id=m_mitra.mitra_id").
		Where("mitra_group_accs_status = ? AND m_mitra_group_access.mitra_id = ?", status, mitra).
		Find(&mitragroup).Error
	if err != nil {
		return &[]MitraGroupAccessData{}, err
	}
	return &mitragroup, nil
}

func (p *MitraGroupAccess) FindMitraGroupAccessDataByID(db *gorm.DB, pid uint64) (*MitraGroupAccess, error) {
	var err error
	err = db.Debug().Model(&MitraGroupAccess{}).Where("mitra_group_accs_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraGroupAccess{}, err
	}
	return p, nil
}

func (p *MitraGroupAccess) FindMitraGroupAccessByID(db *gorm.DB, pid uint64) (*MitraGroupAccessData, error) {
	var err error
	mitragroup := MitraGroupAccessData{}
	err = db.Debug().Model(&MitraGroupAccessData{}).
		Select(`m_mitra_group_access.mitra_group_accs_id,
				m_mitra_group_access.mitra_id,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra.mitra_website,
				m_mitra.mitra_email,
				m_mitra.mitra_logo,
				m_mitra_group_access.mitra_group_accs_code,
				m_mitra_group_access.mitra_group_accs_name,
				m_mitra_group_access.mitra_group_accs_role,
				m_mitra_group_access.mitra_group_accs_type,
				m_mitra_group_access.mitra_group_accs_status,
				m_mitra_group_access.mitra_group_accs_created_by,
				m_mitra_group_access.mitra_group_accs_updated_by,
				m_mitra_group_access.mitra_group_accs_deleted_by,
				m_mitra_group_access.mitra_group_accs_created_at,
				m_mitra_group_access.mitra_group_accs_updated_at,
				m_mitra_group_access.mitra_group_accs_deleted_at`).
		Joins("JOIN m_mitra on m_mitra_group_access.mitra_id=m_mitra.mitra_id").
		Where("mitra_group_accs_id = ?", pid).
		Take(&mitragroup).Error
	if err != nil {
		return &MitraGroupAccessData{}, err
	}
	return &mitragroup, nil
}

func (p *MitraGroupAccess) FindMitraGroupAccessByIDMitra(db *gorm.DB, pid uint64, mitra uint64) (*MitraGroupAccessData, error) {
	var err error
	mitragroup := MitraGroupAccessData{}
	err = db.Debug().Model(&MitraGroupAccessData{}).
		Select(`m_mitra_group_access.mitra_group_accs_id,
				m_mitra_group_access.mitra_id,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra.mitra_website,
				m_mitra.mitra_email,
				m_mitra.mitra_logo,
				m_mitra_group_access.mitra_group_accs_code,
				m_mitra_group_access.mitra_group_accs_name,
				m_mitra_group_access.mitra_group_accs_role,
				m_mitra_group_access.mitra_group_accs_type,
				m_mitra_group_access.mitra_group_accs_status,
				m_mitra_group_access.mitra_group_accs_created_by,
				m_mitra_group_access.mitra_group_accs_updated_by,
				m_mitra_group_access.mitra_group_accs_deleted_by,
				m_mitra_group_access.mitra_group_accs_created_at,
				m_mitra_group_access.mitra_group_accs_updated_at,
				m_mitra_group_access.mitra_group_accs_deleted_at`).
		Joins("JOIN m_mitra on m_mitra_group_access.mitra_id=m_mitra.mitra_id").
		Where("mitra_group_accs_id = ? AND m_mitra_group_access.mitra_id = ?", pid, mitra).
		Take(&mitragroup).Error
	if err != nil {
		return &MitraGroupAccessData{}, err
	}
	return &mitragroup, nil
}

func (p *MitraGroupAccess) FindMitraGroupAccessStatusByID(db *gorm.DB, pid uint64, status uint64, mitra uint64) (*MitraGroupAccessData, error) {
	var err error
	mitragroup := MitraGroupAccessData{}
	err = db.Debug().Model(&MitraGroupAccessData{}).
		Select(`m_mitra_group_access.mitra_group_accs_id,
				m_mitra_group_access.mitra_id,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra.mitra_website,
				m_mitra.mitra_email,
				m_mitra.mitra_logo,
				m_mitra_group_access.mitra_group_accs_code,
				m_mitra_group_access.mitra_group_accs_name,
				m_mitra_group_access.mitra_group_accs_role,
				m_mitra_group_access.mitra_group_accs_type,
				m_mitra_group_access.mitra_group_accs_status,
				m_mitra_group_access.mitra_group_accs_created_by,
				m_mitra_group_access.mitra_group_accs_updated_by,
				m_mitra_group_access.mitra_group_accs_deleted_by,
				m_mitra_group_access.mitra_group_accs_created_at,
				m_mitra_group_access.mitra_group_accs_updated_at,
				m_mitra_group_access.mitra_group_accs_deleted_at`).
		Joins("JOIN m_mitra on m_mitra_group_access.mitra_id=m_mitra.mitra_id").
		Where("mitra_group_accs_id = ? AND mitra_group_accs_status = ? AND m_mitra_group_access.mitra_id = ?", pid, status, mitra).
		Take(&mitragroup).Error
	if err != nil {
		return &MitraGroupAccessData{}, err
	}
	return &mitragroup, nil
}

func (p *MitraGroupAccess) UpdateMitraGroupAccess(db *gorm.DB, pid uint64) (*MitraGroupAccess, error) {

	var err error
	db = db.Debug().Model(&MitraGroupAccess{}).Where("mitra_group_accs_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_id":                    p.MitraID,
			"mitra_group_accs_code":       p.MitraGroupAccessCode,
			"mitra_group_accs_name":       p.MitraGroupAccessName,
			"mitra_group_accs_role":       p.MitraGroupAccessRole,
			"mitra_group_accs_type":       p.MitraGroupAccessType,
			"mitra_group_accs_status":     p.MitraGroupAccessStatus,
			"mitra_group_accs_updated_by": p.MitraGroupAccessUpdatedBy,
			"mitra_group_accs_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraGroupAccess{}).Where("mitra_group_accs_id = ?", pid).Error
	if err != nil {
		return &MitraGroupAccess{}, err
	}
	return p, nil
}

func (p *MitraGroupAccess) UpdateMitraGroupAccessStatus(db *gorm.DB, pid uint64) (*MitraGroupAccess, error) {

	var err error
	db = db.Debug().Model(&MitraGroupAccess{}).Where("mitra_group_accs_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_group_accs_status":     p.MitraGroupAccessStatus,
			"mitra_group_accs_updated_by": p.MitraGroupAccessUpdatedBy,
			"mitra_group_accs_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraGroupAccess{}).Where("mitra_group_accs_id = ?", pid).Error
	if err != nil {
		return &MitraGroupAccess{}, err
	}
	return p, nil
}

func (p *MitraGroupAccess) UpdateMitraGroupAccessStatusOnly(db *gorm.DB, pid uint64) (*MitraGroupAccess, error) {
	var err error
	db = db.Debug().Model(&MitraGroupAccess{}).Where("mitra_group_accs_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_group_accs_status": p.MitraGroupAccessStatus,
		},
	)
	err = db.Debug().Model(&MitraGroupAccess{}).Where("mitra_group_accs_id = ?", pid).Error
	if err != nil {
		return &MitraGroupAccess{}, err
	}
	return p, nil
}

func (p *MitraGroupAccess) UpdateMitraGroupAccessStatusDelete(db *gorm.DB, pid uint64) (*MitraGroupAccess, error) {
	var err error
	db = db.Debug().Model(&MitraGroupAccess{}).Where("mitra_group_accs_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_group_accs_status":     3,
			"mitra_group_accs_deleted_by": p.MitraGroupAccessDeletedBy,
			"mitra_group_accs_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraGroupAccess{}).Where("mitra_group_accs_id = ?", pid).Error
	if err != nil {
		return &MitraGroupAccess{}, err
	}
	return p, nil
}

func (p *MitraGroupAccess) DeleteMitraGroupAccess(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraGroupAccess{}).Where("mitra_group_accs_id = ?", pid).Take(&MitraGroupAccess{}).Delete(&MitraGroupAccess{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraGroupAccess not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

type MitraGroupAccessPrivilege struct {
	MitraGroupAccessID uint64 `gorm:"column:mitra_group_accs_id" json:"mitra_group_accs_id"`
	MitraPrivilegeID   uint64 `gorm:"column:mitra_priv_id" json:"mitra_priv_id"`
}

type MitraGroupAccessPrivilegeData struct {
	MitraGroupAccessPrivilegeID uint64 `gorm:"column:mitra_group_accs_priv_id" json:"mitra_group_accs_priv_id"`
	MitraGroupAccessID          uint64 `gorm:"column:mitra_group_accs_id" json:"mitra_group_accs_id"`
	MitraPrivilegeID            uint64 `gorm:"column:mitra_priv_id" json:"mitra_priv_id"`
	MitraPrivilegeCode          string `gorm:"column:mitra_priv_code" json:"mitra_priv_code"`
	MitraPrivilegeName          string `gorm:"column:mitra_priv_name" json:"mitra_priv_name"`
	MitraPrivilegeCategoryID    uint64 `gorm:"column:mitra_priv_ctgry_id" json:"mitra_priv_ctgry_id"`
	MitraPrivilegeCategoryCode  string `gorm:"column:mitra_priv_ctgry_code" json:"mitra_priv_ctgry_code"`
	MitraPrivilegeCategoryName  string `gorm:"column:mitra_priv_ctgry_name" json:"mitra_priv_ctgry_name"`
}

type MitraGroupAccessPrivilegeDataLogin struct {
	MitraPrivilegeCode string `gorm:"column:mitra_priv_code" json:"mitra_priv_code"`
}

func (MitraGroupAccessPrivilege) TableName() string {
	return "m_mitra_group_access_privilege"
}

func (MitraGroupAccessPrivilegeData) TableName() string {
	return "m_mitra_group_access_privilege"
}

func (MitraGroupAccessPrivilegeDataLogin) TableName() string {
	return "m_mitra_group_access_privilege"
}

func (p *MitraGroupAccessPrivilege) Prepare() {
	p.MitraGroupAccessID = p.MitraGroupAccessID
	p.MitraPrivilegeID = p.MitraPrivilegeID
}

func (p *MitraGroupAccessPrivilege) SaveMitraGroupAccessPrivilege(db *gorm.DB) (*MitraGroupAccessPrivilege, error) {
	var err error
	err = db.Debug().Model(&MitraGroupAccessPrivilege{}).Create(&p).Error
	if err != nil {
		return &MitraGroupAccessPrivilege{}, err
	}
	return p, nil
}

func (p *MitraGroupAccessPrivilege) FindAllMitraGroupAccessPrivilege(db *gorm.DB) (*[]MitraGroupAccessPrivilegeData, error) {
	var err error
	mitragroup := []MitraGroupAccessPrivilegeData{}
	err = db.Debug().Model(&MitraGroupAccessPrivilegeData{}).Limit(100).
		Select(`m_mitra_group_access_privilege.mitra_group_accs_priv_id,
				m_mitra_group_access_privilege.mitra_group_accs_id,
				m_mitra_privilege.mitra_priv_id,
				m_mitra_privilege.mitra_priv_code,
				m_mitra_privilege.mitra_priv_name,
				m_mitra_privilege_category.mitra_priv_ctgry_id,
				m_mitra_privilege_category.mitra_priv_ctgry_code,
				m_mitra_privilege_category.mitra_priv_ctgry_name`).
		Joins("JOIN m_mitra_privilege on m_mitra_group_access_privilege.mitra_priv_id=m_mitra_privilege.mitra_priv_id").
		Joins("JOIN m_mitra_privilege_category on m_mitra_privilege.mitra_priv_ctgry_id=m_mitra_privilege_category.mitra_priv_ctgry_id").
		Find(&mitragroup).Error
	if err != nil {
		return &[]MitraGroupAccessPrivilegeData{}, err
	}
	return &mitragroup, nil
}

func (p *MitraGroupAccessPrivilege) FindMitraGroupAccessPrivilegeDataByID(db *gorm.DB, pid uint64) (*MitraGroupAccessPrivilege, error) {
	var err error
	err = db.Debug().Model(&MitraGroupAccessPrivilege{}).Where("mitra_group_accs_priv_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraGroupAccessPrivilege{}, err
	}
	return p, nil
}

func (p *MitraGroupAccessPrivilege) FindMitraGroupAccessPrivilegeByID(db *gorm.DB, pid uint64) (*MitraGroupAccessPrivilegeData, error) {
	var err error
	mitragroup := MitraGroupAccessPrivilegeData{}
	err = db.Debug().Model(&MitraGroupAccessPrivilegeData{}).Limit(100).
		Select(`m_mitra_group_access_privilege.mitra_group_accs_priv_id,
				m_mitra_group_access_privilege.mitra_group_accs_id,
				m_mitra_privilege.mitra_priv_id,
				m_mitra_privilege.mitra_priv_code,
				m_mitra_privilege.mitra_priv_name,
				m_mitra_privilege_category.mitra_priv_ctgry_id,
				m_mitra_privilege_category.mitra_priv_ctgry_code,
				m_mitra_privilege_category.mitra_priv_ctgry_name`).
		Joins("JOIN m_mitra_privilege on m_mitra_group_access_privilege.mitra_priv_id=m_mitra_privilege.mitra_priv_id").
		Joins("JOIN m_mitra_privilege_category on m_mitra_privilege.mitra_priv_ctgry_id=m_mitra_privilege_category.mitra_priv_ctgry_id").
		Where("mitra_group_accs_priv_id = ?", pid).
		Find(&mitragroup).Error
	if err != nil {
		return &MitraGroupAccessPrivilegeData{}, err
	}
	return &mitragroup, nil
}

func (p *MitraGroupAccessPrivilege) FindMitraGroupAccessPrivilegeDataByMitraGroupAccessID(db *gorm.DB, pid uint64) (*MitraGroupAccessPrivilege, error) {
	var err error
	err = db.Debug().Model(&MitraGroupAccessPrivilege{}).Where("mitra_group_accs_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraGroupAccessPrivilege{}, err
	}
	return p, nil
}

func (p *MitraGroupAccessPrivilege) FindMitraGroupAccessPrivilegeByMitraGroupAccessID(db *gorm.DB, pid uint64) ([]MitraGroupAccessPrivilegeData, error) {
	var err error
	mitragroup := []MitraGroupAccessPrivilegeData{}
	err = db.Debug().Model(&MitraGroupAccessPrivilegeData{}).Limit(100).
		Select(`m_mitra_group_access_privilege.mitra_group_accs_priv_id,
				m_mitra_group_access_privilege.mitra_group_accs_id,
				m_mitra_privilege.mitra_priv_id,
				m_mitra_privilege.mitra_priv_code,
				m_mitra_privilege.mitra_priv_name,
				m_mitra_privilege_category.mitra_priv_ctgry_id,
				m_mitra_privilege_category.mitra_priv_ctgry_code,
				m_mitra_privilege_category.mitra_priv_ctgry_name`).
		Joins("JOIN m_mitra_privilege on m_mitra_group_access_privilege.mitra_priv_id=m_mitra_privilege.mitra_priv_id").
		Joins("JOIN m_mitra_privilege_category on m_mitra_privilege.mitra_priv_ctgry_id=m_mitra_privilege_category.mitra_priv_ctgry_id").
		Where("mitra_group_accs_id = ?", pid).
		Find(&mitragroup).Error
	if err != nil {
		return []MitraGroupAccessPrivilegeData{}, err
	}
	return mitragroup, nil
}

func (p *MitraGroupAccessPrivilege) FindMitraGroupAccessPrivilegeByMitraGroupAccessIDLogin(db *gorm.DB, pid uint64) ([]MitraGroupAccessPrivilegeDataLogin, error) {
	var err error
	mitragroup := []MitraGroupAccessPrivilegeDataLogin{}
	err = db.Debug().Model(&MitraGroupAccessPrivilegeDataLogin{}).Limit(100).
		Select(`m_mitra_group_access_privilege.mitra_group_accs_priv_id,
				m_mitra_group_access_privilege.mitra_group_accs_id,
				m_mitra_privilege.mitra_priv_id,
				m_mitra_privilege.mitra_priv_code,
				m_mitra_privilege.mitra_priv_name,
				m_mitra_privilege_category.mitra_priv_ctgry_id,
				m_mitra_privilege_category.mitra_priv_ctgry_code,
				m_mitra_privilege_category.mitra_priv_ctgry_name`).
		Joins("JOIN m_mitra_privilege on m_mitra_group_access_privilege.mitra_priv_id=m_mitra_privilege.mitra_priv_id").
		Joins("JOIN m_mitra_privilege_category on m_mitra_privilege.mitra_priv_ctgry_id=m_mitra_privilege_category.mitra_priv_ctgry_id").
		Where("mitra_group_accs_id = ?", pid).
		Find(&mitragroup).Error
	if err != nil {
		return []MitraGroupAccessPrivilegeDataLogin{}, err
	}
	return mitragroup, nil
}

func (p *MitraGroupAccessPrivilege) UpdateMitraGroupAccessPrivilege(db *gorm.DB, pid uint64) (*MitraGroupAccessPrivilege, error) {
	var err error
	db = db.Debug().Model(&MitraGroupAccessPrivilege{}).Where("mitra_group_accs_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_group_accs_id": p.MitraGroupAccessID,
			"mitra_priv_id":       p.MitraPrivilegeID,
		},
	)
	err = db.Debug().Model(&MitraGroupAccessPrivilege{}).Where("mitra_group_accs_id = ?", pid).Error
	if err != nil {
		return &MitraGroupAccessPrivilege{}, err
	}
	return p, nil
}

func (p *MitraGroupAccessPrivilege) DeleteMitraGroupAccessPrivilege(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraGroupAccessPrivilege{}).Where("mitra_group_accs_id = ?", pid).Take(&MitraGroupAccessPrivilege{}).Delete(&MitraGroupAccessPrivilege{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraGroupAccessPrivilege not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
