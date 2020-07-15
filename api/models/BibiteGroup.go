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

type BibiteGroup struct {
	BibiteGroupID        uint64     `gorm:"column:bbt_group_id;primary_key;" json:"bbt_group_id"`
	BibiteGroupCode      string     `gorm:"column:bbt_group_code" json:"bbt_group_code"`
	BibiteGroupName      string     `gorm:"column:bbt_group_name" json:"bbt_group_name"`
	BibiteGroupRole      string     `gorm:"column:bbt_group_role" json:"bbt_group_role"`
	BibiteGroupType      uint64     `gorm:"column:bbt_group_type" json:"bbt_group_type"`
	BibiteGroupStatus    uint64     `gorm:"column:bbt_group_status;size:2" json:"bbt_group_status"`
	BibiteGroupCreatedBy string     `gorm:"column:bbt_group_created_by;size:125" json:"bbt_group_created_by"`
	BibiteGroupCreatedAt time.Time  `gorm:"column:bbt_group_created_at;default:CURRENT_TIMESTAMP" json:"bbt_group_created_at"`
	BibiteGroupUpdatedBy string     `gorm:"column:bbt_group_updated_by;size:125" json:"bbt_group_updated_by"`
	BibiteGroupUpdatedAt *time.Time `gorm:"column:bbt_group_updated_at;default:CURRENT_TIMESTAMP" json:"bbt_group_created_at"`
	BibiteGroupDeletedBy string     `gorm:"column:bbt_group_deleted_by;size:125" json:"bbt_group_deleted_by"`
	BibiteGroupDeletedAt *time.Time `gorm:"column:bbt_group_deleted_at;default:CURRENT_TIMESTAMP" json:"bbt_group_deleted_at"`
}

type BibiteGroupData struct {
	BibiteGroupID        uint64                     `gorm:"column:bbt_group_id" json:"bbt_group_id"`
	BibiteGroupCode      string                     `gorm:"column:bbt_group_code" json:"bbt_group_code"`
	BibiteGroupName      string                     `gorm:"column:bbt_group_name" json:"bbt_group_name"`
	BibiteGroupRole      string                     `gorm:"column:bbt_group_role" json:"bbt_group_role"`
	BibiteGroupType      uint64                     `gorm:"column:bbt_group_type" json:"bbt_group_type"`
	BibiteGroupStatus    uint64                     `gorm:"column:bbt_group_status" json:"bbt_group_status"`
	BibiteGroupCreatedBy string                     `gorm:"column:bbt_group_created_by" json:"bbt_group_created_by"`
	BibiteGroupCreatedAt time.Time                  `gorm:"column:bbt_group_created_at" json:"bbt_group_created_at"`
	BibiteGroupUpdatedBy string                     `gorm:"column:bbt_group_updated_by" json:"bbt_group_updated_by"`
	BibiteGroupUpdatedAt *time.Time                 `gorm:"column:bbt_group_updated_at" json:"bbt_group_updated_at"`
	BibiteGroupDeletedBy string                     `gorm:"column:bbt_group_deleted_by" json:"bbt_group_deleted_by"`
	BibiteGroupDeletedAt *time.Time                 `gorm:"column:bbt_group_deleted_at" json:"bbt_group_deleted_at"`
	BibiteGroupPrivilege []BibiteGroupPrivilegeData `json:"bbt_group_privilege"`
}

type ResponseBibiteGroups struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    *[]BibiteGroupData `json:"data"`
}

type ResponseBibiteGroup struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    *BibiteGroupData `json:"data"`
}

type ResponseBibiteGroupCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (BibiteGroup) TableName() string {
	return "m_bibite_group"
}

func (BibiteGroupData) TableName() string {
	return "m_bibite_group"
}

func (p *BibiteGroup) Prepare() {
	p.BibiteGroupCode = p.BibiteGroupCode
	p.BibiteGroupName = p.BibiteGroupName
	p.BibiteGroupRole = p.BibiteGroupRole
	p.BibiteGroupType = p.BibiteGroupType
	p.BibiteGroupStatus = p.BibiteGroupStatus
	p.BibiteGroupCreatedBy = html.EscapeString(strings.TrimSpace(p.BibiteGroupCreatedBy))
	p.BibiteGroupCreatedAt = time.Now()
	p.BibiteGroupUpdatedBy = html.EscapeString(strings.TrimSpace(p.BibiteGroupUpdatedBy))
	p.BibiteGroupUpdatedAt = p.BibiteGroupUpdatedAt
	p.BibiteGroupDeletedBy = html.EscapeString(strings.TrimSpace(p.BibiteGroupDeletedBy))
	p.BibiteGroupDeletedAt = p.BibiteGroupDeletedAt
}

func (p *BibiteGroup) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.BibiteGroupCode == "" {
			return errors.New("Required Country")
		}
		if p.BibiteGroupName == "" {
			return errors.New("Required Province")
		}
		if p.BibiteGroupRole == "" {
			return errors.New("Required Regency")
		}
		if p.BibiteGroupType == 0 {
			return errors.New("Required District")
		}
		return nil
	}
}

func (p *BibiteGroup) SaveBibiteGroup(db *gorm.DB) (*BibiteGroup, error) {
	var err error
	err = db.Debug().Model(&BibiteGroup{}).Create(&p).Error
	if err != nil {
		return &BibiteGroup{}, err
	}
	return p, nil
}

func (p *BibiteGroup) FindAllBibiteGroup(db *gorm.DB) (*[]BibiteGroupData, error) {
	var err error
	bibitegroup := []BibiteGroupData{}
	err = db.Debug().Model(&BibiteGroupData{}).Limit(100).
		Select(`m_bibite_group.bbt_group_id,
				m_bibite_group.bbt_group_code,
				m_bibite_group.bbt_group_name,
				m_bibite_group.bbt_group_role,
				m_bibite_group.bbt_group_type,
				m_bibite_group.bbt_group_status,
				m_bibite_group.bbt_group_created_by,
				m_bibite_group.bbt_group_created_at at time zone 'Asia/Jakarta' as bbt_group_created_at,
				m_bibite_group.bbt_group_updated_by,
				m_bibite_group.bbt_group_updated_at at time zone 'Asia/Jakarta' as bbt_group_updated_at,
				m_bibite_group.bbt_group_deleted_by,
				m_bibite_group.bbt_group_deleted_at at time zone 'Asia/Jakarta' as bbt_group_deleted_at`).
		Find(&bibitegroup).Error
	if err != nil {
		return &[]BibiteGroupData{}, err
	}
	return &bibitegroup, nil
}

func (p *BibiteGroup) FindAllBibiteGroupStatus(db *gorm.DB, status uint64) (*[]BibiteGroupData, error) {
	var err error
	bibitegroup := []BibiteGroupData{}
	err = db.Debug().Model(&BibiteGroupData{}).Limit(100).
		Select(`m_bibite_group.bbt_group_id,
				m_bibite_group.bbt_group_code,
				m_bibite_group.bbt_group_name,
				m_bibite_group.bbt_group_role,
				m_bibite_group.bbt_group_type,
				m_bibite_group.bbt_group_status,
				m_bibite_group.bbt_group_created_by,
				m_bibite_group.bbt_group_created_at at time zone 'Asia/Jakarta' as bbt_group_created_at,
				m_bibite_group.bbt_group_updated_by,
				m_bibite_group.bbt_group_updated_at at time zone 'Asia/Jakarta' as bbt_group_updated_at,
				m_bibite_group.bbt_group_deleted_by,
				m_bibite_group.bbt_group_deleted_at at time zone 'Asia/Jakarta' as bbt_group_deleted_at`).
		Where("bbt_group_status = ?", status).
		Find(&bibitegroup).Error
	if err != nil {
		return &[]BibiteGroupData{}, err
	}
	return &bibitegroup, nil
}

func (p *BibiteGroup) FindBibiteGroupDataByID(db *gorm.DB, pid uint64) (*BibiteGroup, error) {
	var err error
	err = db.Debug().Model(&BibiteGroup{}).
		Select(`m_bibite_group.bbt_group_id,
				m_bibite_group.bbt_group_code,
				m_bibite_group.bbt_group_name,
				m_bibite_group.bbt_group_role,
				m_bibite_group.bbt_group_type,
				m_bibite_group.bbt_group_status,
				m_bibite_group.bbt_group_created_by,
				m_bibite_group.bbt_group_created_at at time zone 'Asia/Jakarta' as bbt_group_created_at,
				m_bibite_group.bbt_group_updated_by,
				m_bibite_group.bbt_group_updated_at at time zone 'Asia/Jakarta' as bbt_group_updated_at,
				m_bibite_group.bbt_group_deleted_by,
				m_bibite_group.bbt_group_deleted_at at time zone 'Asia/Jakarta' as bbt_group_deleted_at`).
		Where("bbt_group_id = ?", pid).Take(&p).Error
	if err != nil {
		return &BibiteGroup{}, err
	}
	return p, nil
}

func (p *BibiteGroup) FindBibiteGroupByID(db *gorm.DB, pid uint64) (*BibiteGroupData, error) {
	var err error
	bibitegroup := BibiteGroupData{}
	err = db.Debug().Model(&BibiteGroupData{}).Limit(100).
		Select(`m_bibite_group.bbt_group_id,
				m_bibite_group.bbt_group_code,
				m_bibite_group.bbt_group_name,
				m_bibite_group.bbt_group_role,
				m_bibite_group.bbt_group_type,
				m_bibite_group.bbt_group_status,
				m_bibite_group.bbt_group_created_by,
				m_bibite_group.bbt_group_created_at at time zone 'Asia/Jakarta' as bbt_group_created_at,
				m_bibite_group.bbt_group_updated_by,
				m_bibite_group.bbt_group_updated_at at time zone 'Asia/Jakarta' as bbt_group_updated_at,
				m_bibite_group.bbt_group_deleted_by,
				m_bibite_group.bbt_group_deleted_at at time zone 'Asia/Jakarta' as bbt_group_deleted_at`).
		Where("bbt_group_id = ?", pid).
		Find(&bibitegroup).Error
	if err != nil {
		return &BibiteGroupData{}, err
	}
	return &bibitegroup, nil
}

func (p *BibiteGroup) FindBibiteGroupStatusByID(db *gorm.DB, pid uint64, status uint64) (*BibiteGroupData, error) {
	var err error
	bibitegroup := BibiteGroupData{}
	err = db.Debug().Model(&BibiteGroupData{}).Limit(100).
		Select(`m_bibite_group.bbt_group_id,
				m_bibite_group.bbt_group_code,
				m_bibite_group.bbt_group_name,
				m_bibite_group.bbt_group_role,
				m_bibite_group.bbt_group_type,
				m_bibite_group.bbt_group_status,
				m_bibite_group.bbt_group_created_by,
				m_bibite_group.bbt_group_created_at at time zone 'Asia/Jakarta' as bbt_group_created_at,
				m_bibite_group.bbt_group_updated_by,
				m_bibite_group.bbt_group_updated_at at time zone 'Asia/Jakarta' as bbt_group_updated_at,
				m_bibite_group.bbt_group_deleted_by,
				m_bibite_group.bbt_group_deleted_at at time zone 'Asia/Jakarta' as bbt_group_deleted_at`).
		Where("bbt_group_id = ? AND bbt_group_status = ?", pid, status).
		Find(&bibitegroup).Error
	if err != nil {
		return &BibiteGroupData{}, err
	}
	return &bibitegroup, nil
}

func (p *BibiteGroup) UpdateBibiteGroup(db *gorm.DB, pid uint64) (*BibiteGroup, error) {
	var err error
	db = db.Debug().Model(&BibiteGroup{}).Where("bbt_group_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_group_code":       p.BibiteGroupCode,
			"bbt_group_name":       p.BibiteGroupName,
			"bbt_group_role":       p.BibiteGroupRole,
			"bbt_group_type":       p.BibiteGroupType,
			"bbt_group_status":     p.BibiteGroupStatus,
			"bbt_group_updated_by": p.BibiteGroupUpdatedBy,
			"bbt_group_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&BibiteGroup{}).Where("bbt_group_id = ?", pid).Error
	if err != nil {
		return &BibiteGroup{}, err
	}
	return p, nil
}

func (p *BibiteGroup) UpdateBibiteGroupStatus(db *gorm.DB, pid uint64) (*BibiteGroup, error) {
	var err error
	db = db.Debug().Model(&BibiteGroup{}).Where("bbt_group_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_group_status":     p.BibiteGroupStatus,
			"bbt_group_updated_by": p.BibiteGroupUpdatedBy,
			"bbt_group_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&BibiteGroup{}).Where("bbt_group_id = ?", pid).Error
	if err != nil {
		return &BibiteGroup{}, err
	}
	return p, nil
}

func (p *BibiteGroup) UpdateBibiteGroupStatusOnly(db *gorm.DB, pid uint64) (*BibiteGroup, error) {
	var err error
	db = db.Debug().Model(&BibiteGroup{}).Where("bbt_group_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_group_status": p.BibiteGroupStatus,
		},
	)
	err = db.Debug().Model(&BibiteGroup{}).Where("bbt_group_id = ?", pid).Error
	if err != nil {
		return &BibiteGroup{}, err
	}
	return p, nil
}

func (p *BibiteGroup) UpdateBibiteGroupStatusDelete(db *gorm.DB, pid uint64) (*BibiteGroup, error) {
	var err error
	db = db.Debug().Model(&BibiteGroup{}).Where("bbt_group_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_group_status":     3,
			"bbt_group_deleted_by": p.BibiteGroupDeletedBy,
			"bbt_group_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&BibiteGroup{}).Where("bbt_group_id = ?", pid).Error
	if err != nil {
		return &BibiteGroup{}, err
	}
	return p, nil
}

func (p *BibiteGroup) DeleteBibiteGroup(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&BibiteGroup{}).Where("bbt_group_id = ?", pid).Take(&BibiteGroup{}).Delete(&BibiteGroup{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("BibiteGroup not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

type BibiteGroupPrivilege struct {
	BibiteGroupID     uint64 `gorm:"column:bbt_group_id" json:"bbt_group_id"`
	BibitePrivilegeID uint64 `gorm:"column:bbt_priv_id" json:"bbt_priv_id"`
}

type BibiteGroupPrivilegeData struct {
	BibiteGroupPrivilegeID      uint64 `gorm:"column:bbt_group_priv_id" json:"bbt_group_priv_id"`
	BibiteGroupID               uint64 `gorm:"column:bbt_group_id" json:"bbt_group_id"`
	BibitePrivilegeID           uint64 `gorm:"column:bbt_priv_id" json:"bbt_priv_id"`
	BibitePrivilegeCode         string `gorm:"column:bbt_priv_code" json:"bbt_priv_code"`
	BibitePrivilegeName         string `gorm:"column:bbt_priv_name" json:"bbt_priv_name"`
	BibitePrivilegeCategoryID   uint64 `gorm:"column:bbt_priv_ctgry_id" json:"bbt_priv_ctgry_id"`
	BibitePrivilegeCategoryCode string `gorm:"column:bbt_priv_ctgry_code" json:"bbt_priv_ctgry_code"`
	BibitePrivilegeCategoryName string `gorm:"column:bbt_priv_ctgry_name" json:"bbt_priv_ctgry_name"`
}

type BibiteGroupPrivilegeDataLogin struct {
	BibitePrivilegeCode string `gorm:"column:bbt_priv_code" json:"bbt_priv_code"`
}

func (BibiteGroupPrivilege) TableName() string {
	return "m_bibite_group_privilege"
}

func (BibiteGroupPrivilegeData) TableName() string {
	return "m_bibite_group_privilege"
}

func (BibiteGroupPrivilegeDataLogin) TableName() string {
	return "m_bibite_group_privilege"
}

func (p *BibiteGroupPrivilege) Prepare() {
	p.BibiteGroupID = p.BibiteGroupID
	p.BibitePrivilegeID = p.BibitePrivilegeID
}

func (p *BibiteGroupPrivilege) SaveBibiteGroupPrivilege(db *gorm.DB) (*BibiteGroupPrivilege, error) {
	var err error
	err = db.Debug().Model(&BibiteGroupPrivilege{}).Create(&p).Error
	if err != nil {
		return &BibiteGroupPrivilege{}, err
	}
	return p, nil
}

func (p *BibiteGroupPrivilege) FindAllBibiteGroupPrivilege(db *gorm.DB) (*[]BibiteGroupPrivilegeData, error) {
	var err error
	bibitegroup := []BibiteGroupPrivilegeData{}
	err = db.Debug().Model(&BibiteGroupPrivilegeData{}).Limit(100).
		Select(`m_bibite_group_privilege.bbt_group_priv_id,
				m_bibite_group_privilege.bbt_group_id,
				m_bibite_privilege.bbt_priv_id,
				m_bibite_privilege.bbt_priv_code,
				m_bibite_privilege.bbt_priv_name,
				m_bibite_privilege_category.bbt_priv_ctgry_id,
				m_bibite_privilege_category.bbt_priv_ctgry_code,
				m_bibite_privilege_category.bbt_priv_ctgry_name`).
		Joins("JOIN m_bibite_privilege on m_bibite_group_privilege.bbt_priv_id=m_bibite_privilege.bbt_priv_id").
		Joins("JOIN m_bibite_privilege_category on m_bibite_privilege.bbt_priv_ctgry_id=m_bibite_privilege_category.bbt_priv_ctgry_id").
		Find(&bibitegroup).Error
	if err != nil {
		return &[]BibiteGroupPrivilegeData{}, err
	}
	return &bibitegroup, nil
}

func (p *BibiteGroupPrivilege) FindBibiteGroupPrivilegeDataByID(db *gorm.DB, pid uint64) (*BibiteGroupPrivilege, error) {
	var err error
	err = db.Debug().Model(&BibiteGroupPrivilege{}).Where("bbt_group_priv_id = ?", pid).Take(&p).Error
	if err != nil {
		return &BibiteGroupPrivilege{}, err
	}
	return p, nil
}

func (p *BibiteGroupPrivilege) FindBibiteGroupPrivilegeByID(db *gorm.DB, pid uint64) (*BibiteGroupPrivilegeData, error) {
	var err error
	bibitegroup := BibiteGroupPrivilegeData{}
	err = db.Debug().Model(&BibiteGroupPrivilegeData{}).Limit(100).
		Select(`m_bibite_group_privilege.bbt_group_priv_id,
				m_bibite_group_privilege.bbt_group_id,
				m_bibite_privilege.bbt_priv_id,
				m_bibite_privilege.bbt_priv_code,
				m_bibite_privilege.bbt_priv_name,
				m_bibite_privilege_category.bbt_priv_ctgry_id,
				m_bibite_privilege_category.bbt_priv_ctgry_code,
				m_bibite_privilege_category.bbt_priv_ctgry_name`).
		Joins("JOIN m_bibite_privilege on m_bibite_group_privilege.bbt_priv_id=m_bibite_privilege.bbt_priv_id").
		Joins("JOIN m_bibite_privilege_category on m_bibite_privilege.bbt_priv_ctgry_id=m_bibite_privilege_category.bbt_priv_ctgry_id").
		Where("bbt_group_priv_id = ?", pid).
		Find(&bibitegroup).Error
	if err != nil {
		return &BibiteGroupPrivilegeData{}, err
	}
	return &bibitegroup, nil
}

func (p *BibiteGroupPrivilege) FindBibiteGroupPrivilegeDataByBibiteGroupID(db *gorm.DB, pid uint64) (*BibiteGroupPrivilege, error) {
	var err error
	err = db.Debug().Model(&BibiteGroupPrivilege{}).Where("bbt_group_id = ?", pid).Take(&p).Error
	if err != nil {
		return &BibiteGroupPrivilege{}, err
	}
	return p, nil
}

func (p *BibiteGroupPrivilege) FindBibiteGroupPrivilegeByBibiteGroupID(db *gorm.DB, pid uint64) ([]BibiteGroupPrivilegeData, error) {
	var err error
	bibitegroup := []BibiteGroupPrivilegeData{}
	err = db.Debug().Model(&BibiteGroupPrivilegeData{}).Limit(100).
		Select(`m_bibite_group_privilege.bbt_group_priv_id,
				m_bibite_group_privilege.bbt_group_id,
				m_bibite_privilege.bbt_priv_id,
				m_bibite_privilege.bbt_priv_code,
				m_bibite_privilege.bbt_priv_name,
				m_bibite_privilege_category.bbt_priv_ctgry_id,
				m_bibite_privilege_category.bbt_priv_ctgry_code,
				m_bibite_privilege_category.bbt_priv_ctgry_name`).
		Joins("JOIN m_bibite_privilege on m_bibite_group_privilege.bbt_priv_id=m_bibite_privilege.bbt_priv_id").
		Joins("JOIN m_bibite_privilege_category on m_bibite_privilege.bbt_priv_ctgry_id=m_bibite_privilege_category.bbt_priv_ctgry_id").
		Where("bbt_group_id = ?", pid).
		Find(&bibitegroup).Error
	if err != nil {
		return []BibiteGroupPrivilegeData{}, err
	}
	return bibitegroup, nil
}

func (p *BibiteGroupPrivilege) FindBibiteGroupPrivilegeByBibiteGroupIDLogin(db *gorm.DB, pid uint64) ([]BibiteGroupPrivilegeDataLogin, error) {
	var err error
	bibitegroup := []BibiteGroupPrivilegeDataLogin{}
	err = db.Debug().Model(&BibiteGroupPrivilegeDataLogin{}).Limit(100).
		Select(`m_bibite_group_privilege.bbt_group_priv_id,
				m_bibite_group_privilege.bbt_group_id,
				m_bibite_privilege.bbt_priv_id,
				m_bibite_privilege.bbt_priv_code,
				m_bibite_privilege.bbt_priv_name,
				m_bibite_privilege_category.bbt_priv_ctgry_id,
				m_bibite_privilege_category.bbt_priv_ctgry_code,
				m_bibite_privilege_category.bbt_priv_ctgry_name`).
		Joins("JOIN m_bibite_privilege on m_bibite_group_privilege.bbt_priv_id=m_bibite_privilege.bbt_priv_id").
		Joins("JOIN m_bibite_privilege_category on m_bibite_privilege.bbt_priv_ctgry_id=m_bibite_privilege_category.bbt_priv_ctgry_id").
		Where("bbt_group_id = ?", pid).
		Find(&bibitegroup).Error
	if err != nil {
		return []BibiteGroupPrivilegeDataLogin{}, err
	}
	return bibitegroup, nil
}

func (p *BibiteGroupPrivilege) UpdateBibiteGroupPrivilege(db *gorm.DB, pid uint64) (*BibiteGroupPrivilege, error) {
	var err error
	db = db.Debug().Model(&BibiteGroupPrivilege{}).Where("bbt_group_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_group_id": p.BibiteGroupID,
			"bbt_priv_id":  p.BibitePrivilegeID,
		},
	)
	err = db.Debug().Model(&BibiteGroupPrivilege{}).Where("bbt_group_id = ?", pid).Error
	if err != nil {
		return &BibiteGroupPrivilege{}, err
	}
	return p, nil
}

func (p *BibiteGroupPrivilege) DeleteBibiteGroupPrivilege(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&BibiteGroupPrivilege{}).Where("bbt_group_id = ?", pid).Take(&BibiteGroupPrivilege{}).Delete(&BibiteGroupPrivilege{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("BibiteGroupPrivilege not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
