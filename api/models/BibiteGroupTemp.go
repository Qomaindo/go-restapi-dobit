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

type BibiteGroupTemp struct {
	BibiteGroupTempID           uint64                         `gorm:"column:bbt_group_temp_id;primary_key;" json:"bbt_group_temp_id"`
	BibiteGroupID               uint64                         `gorm:"column:bbt_group_id" json:"bbt_group_id"`
	BibiteGroupTempCode         string                         `gorm:"column:bbt_group_temp_code" json:"bbt_group_temp_code"`
	BibiteGroupTempName         string                         `gorm:"column:bbt_group_temp_name" json:"bbt_group_temp_name"`
	BibiteGroupTempRole         string                         `gorm:"column:bbt_group_temp_role" json:"bbt_group_temp_role"`
	BibiteGroupTempType         uint64                         `gorm:"column:bbt_group_temp_type" json:"bbt_group_temp_type"`
	BibiteGroupTempReason       string                         `gorm:"column:bbt_group_temp_reason" json:"bbt_group_temp_reason"`
	BibiteGroupTempStatus       uint64                         `gorm:"column:bbt_group_temp_status;size:2" json:"bbt_group_temp_status"`
	BibiteGroupTempActionBefore uint64                         `gorm:"column:bbt_group_temp_action_before;size:2" json:"bbt_group_temp_action_before"`
	BibiteGroupTempCreatedBy    string                         `gorm:"column:bbt_group_temp_created_by;size:125" json:"bbt_group_temp_created_by"`
	BibiteGroupTempCreatedAt    time.Time                      `gorm:"column:bbt_group_temp_created_at;default:CURRENT_TIMESTAMP" json:"bbt_group_temp_created_at"`
	BibiteGroupTempPrivilege    []BibiteGroupTempPrivilegeTemp `gorm:"-" json:"bbt_group_temp_privilege"`
}

type BibiteGroupTempPend struct {
	BibiteGroupTempID           uint64                         `gorm:"column:bbt_group_temp_id;primary_key;" json:"bbt_group_temp_id"`
	BibiteGroupID               *uint64                        `gorm:"column:bbt_group_id" json:"bbt_group_id"`
	BibiteGroupTempCode         string                         `gorm:"column:bbt_group_temp_code" json:"bbt_group_temp_code"`
	BibiteGroupTempName         string                         `gorm:"column:bbt_group_temp_name" json:"bbt_group_temp_name"`
	BibiteGroupTempRole         string                         `gorm:"column:bbt_group_temp_role" json:"bbt_group_temp_role"`
	BibiteGroupTempType         uint64                         `gorm:"column:bbt_group_temp_type" json:"bbt_group_temp_type"`
	BibiteGroupTempReason       string                         `gorm:"column:bbt_group_temp_reason" json:"bbt_group_temp_reason"`
	BibiteGroupTempStatus       uint64                         `gorm:"column:bbt_group_temp_status;size:2" json:"bbt_group_temp_status"`
	BibiteGroupTempActionBefore uint64                         `gorm:"column:bbt_group_temp_action_before;size:2" json:"bbt_group_temp_action_before"`
	BibiteGroupTempCreatedBy    string                         `gorm:"column:bbt_group_temp_created_by;size:125" json:"bbt_group_temp_created_by"`
	BibiteGroupTempCreatedAt    time.Time                      `gorm:"column:bbt_group_temp_created_at;default:CURRENT_TIMESTAMP" json:"bbt_group_temp_created_at"`
	BibiteGroupTempPrivilege    []BibiteGroupTempPrivilegeTemp `gorm:"-" json:"bbt_group_temp_privilege"`
}

type BibiteGroupTempPendData struct {
	BibiteGroupTempID           uint64                         `gorm:"column:bbt_group_temp_id;primary_key;" json:"bbt_group_temp_id"`
	BibiteGroupID               *uint64                        `gorm:"column:bbt_group_id" json:"bbt_group_id"`
	BibiteGroupTempCode         string                         `gorm:"column:bbt_group_temp_code" json:"bbt_group_temp_code"`
	BibiteGroupTempName         string                         `gorm:"column:bbt_group_temp_name" json:"bbt_group_temp_name"`
	BibiteGroupTempRole         string                         `gorm:"column:bbt_group_temp_role" json:"bbt_group_temp_role"`
	BibiteGroupTempType         uint64                         `gorm:"column:bbt_group_temp_type" json:"bbt_group_temp_type"`
	BibiteGroupTempReason       string                         `gorm:"column:bbt_group_temp_reason" json:"bbt_group_temp_reason"`
	BibiteGroupTempStatus       uint64                         `gorm:"column:bbt_group_temp_status;size:2" json:"bbt_group_temp_status"`
	BibiteGroupTempActionBefore uint64                         `gorm:"column:bbt_group_temp_action_before;size:2" json:"bbt_group_temp_action_before"`
	BibiteGroupTempCreatedBy    string                         `gorm:"column:bbt_group_temp_created_by;size:125" json:"bbt_group_temp_created_by"`
	BibiteGroupTempCreatedAt    time.Time                      `gorm:"column:bbt_group_temp_created_at;default:CURRENT_TIMESTAMP" json:"bbt_group_temp_created_at"`
	BibiteGroupTempPrivilege    []BibiteGroupTempPrivilegeData `json:"bbt_group_temp_privilege"`
}

type BibiteGroupTempDatas struct {
	BibiteGroupTempID           uint64    `gorm:"column:bbt_group_temp_id" json:"bbt_group_temp_id"`
	BibiteGroupTempCode         string    `gorm:"column:bbt_group_temp_code" json:"bbt_group_temp_code"`
	BibiteGroupTempName         string    `gorm:"column:bbt_group_temp_name" json:"bbt_group_temp_name"`
	BibiteGroupTempRole         string    `gorm:"column:bbt_group_temp_role" json:"bbt_group_temp_role"`
	BibiteGroupTempType         uint64    `gorm:"column:bbt_group_temp_type" json:"bbt_group_temp_type"`
	BibiteGroupTempReason       string    `gorm:"column:bbt_group_temp_reason" json:"bbt_group_temp_reason"`
	BibiteGroupTempStatus       uint64    `gorm:"column:bbt_group_temp_status;size:2" json:"bbt_group_temp_status"`
	BibiteGroupTempActionBefore uint64    `gorm:"column:bbt_group_temp_action_before;size:2" json:"bbt_group_temp_action_before"`
	BibiteGroupTempCreatedBy    string    `gorm:"column:bbt_group_temp_created_by;size:125" json:"bbt_group_temp_created_by"`
	BibiteGroupTempCreatedAt    time.Time `gorm:"column:bbt_group_temp_created_at" json:"bbt_group_temp_created_at"`
	BibiteGroupID               uint64    `gorm:"column:bbt_group_id" json:"bbt_group_id"`
	BibiteGroupCode             string    `gorm:"column:bbt_group_code" json:"bbt_group_code"`
	BibiteGroupName             string    `gorm:"column:bbt_group_name" json:"bbt_group_name"`
	BibiteGroupRole             string    `gorm:"column:bbt_group_role" json:"bbt_group_role"`
	BibiteGroupType             string    `gorm:"column:bbt_group_type" json:"bbt_group_type"`
	BibiteGroupStatus           uint64    `gorm:"column:bbt_group_status" json:"bbt_group_status"`
	BibiteGroupCreatedBy        string    `gorm:"column:bbt_group_created_by" json:"bbt_group_created_by"`
	BibiteGroupCreatedAt        time.Time `gorm:"column:bbt_group_created_at" json:"bbt_group_created_at"`
	BibiteGroupUpdatedBy        string    `gorm:"column:bbt_group_updated_by" json:"bbt_group_updated_by"`
	BibiteGroupUpdatedAt        time.Time `gorm:"column:bbt_group_updated_at" json:"bbt_group_updated_at"`
	BibiteGroupDeletedBy        string    `gorm:"column:bbt_group_deleted_by" json:"bbt_group_deleted_by"`
	BibiteGroupDeletedAt        time.Time `gorm:"column:bbt_group_deleted_at" json:"bbt_group_deleted_at"`
}

type BibiteGroupTempData struct {
	BibiteGroupTempID           uint64                         `gorm:"column:bbt_group_temp_id" json:"bbt_group_temp_id"`
	BibiteGroupTempCode         string                         `gorm:"column:bbt_group_temp_code" json:"bbt_group_temp_code"`
	BibiteGroupTempName         string                         `gorm:"column:bbt_group_temp_name" json:"bbt_group_temp_name"`
	BibiteGroupTempRole         string                         `gorm:"column:bbt_group_temp_role" json:"bbt_group_temp_role"`
	BibiteGroupTempType         uint64                         `gorm:"column:bbt_group_temp_type" json:"bbt_group_temp_type"`
	BibiteGroupTempReason       string                         `gorm:"column:bbt_group_temp_reason" json:"bbt_group_temp_reason"`
	BibiteGroupTempStatus       uint64                         `gorm:"column:bbt_group_temp_status;size:2" json:"bbt_group_temp_status"`
	BibiteGroupTempActionBefore uint64                         `gorm:"column:bbt_group_temp_action_before;size:2" json:"bbt_group_temp_action_before"`
	BibiteGroupTempCreatedBy    string                         `gorm:"column:bbt_group_temp_created_by;size:125" json:"bbt_group_temp_created_by"`
	BibiteGroupTempCreatedAt    time.Time                      `gorm:"column:bbt_group_temp_created_at" json:"bbt_group_temp_created_at"`
	BibiteGroupTempPrivilege    []BibiteGroupTempPrivilegeData `json:"bbt_group_temp_privilege"`
	BibiteGroupID               uint64                         `gorm:"column:bbt_group_id" json:"bbt_group_id"`
	BibiteGroupCode             string                         `gorm:"column:bbt_group_code" json:"bbt_group_code"`
	BibiteGroupName             string                         `gorm:"column:bbt_group_name" json:"bbt_group_name"`
	BibiteGroupRole             string                         `gorm:"column:bbt_group_role" json:"bbt_group_role"`
	BibiteGroupType             string                         `gorm:"column:bbt_group_type" json:"bbt_group_type"`
	BibiteGroupStatus           uint64                         `gorm:"column:bbt_group_status" json:"bbt_group_status"`
	BibiteGroupCreatedBy        string                         `gorm:"column:bbt_group_created_by" json:"bbt_group_created_by"`
	BibiteGroupCreatedAt        time.Time                      `gorm:"column:bbt_group_created_at" json:"bbt_group_created_at"`
	BibiteGroupUpdatedBy        string                         `gorm:"column:bbt_group_updated_by" json:"bbt_group_updated_by"`
	BibiteGroupUpdatedAt        time.Time                      `gorm:"column:bbt_group_updated_at" json:"bbt_group_updated_at"`
	BibiteGroupDeletedBy        string                         `gorm:"column:bbt_group_deleted_by" json:"bbt_group_deleted_by"`
	BibiteGroupDeletedAt        time.Time                      `gorm:"column:bbt_group_deleted_at" json:"bbt_group_deleted_at"`
	BibiteGroupPrivilege        []BibiteGroupPrivilegeData     `json:"bbt_group_privilege"`
}

type ResponseBibiteGroupTemps struct {
	Status  int                     `json:"status"`
	Message string                  `json:"message"`
	Data    *[]BibiteGroupTempDatas `json:"data"`
}

type ResponseBibiteGroupTemp struct {
	Status  int                  `json:"status"`
	Message string               `json:"message"`
	Data    *BibiteGroupTempData `json:"data"`
}

type ResponseBibiteGroupTempsPend struct {
	Status  int                        `json:"status"`
	Message string                     `json:"message"`
	Data    *[]BibiteGroupTempPendData `json:"data"`
}

type ResponseBibiteGroupTempPend struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    *BibiteGroupTempPendData `json:"data"`
}

type ResponseBibiteGroupTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (BibiteGroupTemp) TableName() string {
	return "m_bibite_group_temp"
}

func (BibiteGroupTempPend) TableName() string {
	return "m_bibite_group_temp"
}

func (BibiteGroupTempPendData) TableName() string {
	return "m_bibite_group_temp"
}

func (BibiteGroupTempData) TableName() string {
	return "m_bibite_group_temp"
}

func (BibiteGroupTempDatas) TableName() string {
	return "m_bibite_group_temp"
}

func (p *BibiteGroupTemp) Prepare() {
	p.BibiteGroupTempCode = p.BibiteGroupTempCode
	p.BibiteGroupTempName = p.BibiteGroupTempName
	p.BibiteGroupTempRole = p.BibiteGroupTempRole
	p.BibiteGroupTempType = p.BibiteGroupTempType
	p.BibiteGroupTempReason = p.BibiteGroupTempReason
	p.BibiteGroupTempStatus = p.BibiteGroupTempStatus
	p.BibiteGroupTempCreatedBy = html.EscapeString(strings.TrimSpace(p.BibiteGroupTempCreatedBy))
	p.BibiteGroupTempCreatedAt = time.Now()
}

func (p *BibiteGroupTempPend) Prepare() {
	p.BibiteGroupID = nil
	p.BibiteGroupTempCode = p.BibiteGroupTempCode
	p.BibiteGroupTempName = p.BibiteGroupTempName
	p.BibiteGroupTempRole = p.BibiteGroupTempRole
	p.BibiteGroupTempType = p.BibiteGroupTempType
	p.BibiteGroupTempReason = p.BibiteGroupTempReason
	p.BibiteGroupTempStatus = p.BibiteGroupTempStatus
	p.BibiteGroupTempCreatedBy = html.EscapeString(strings.TrimSpace(p.BibiteGroupTempCreatedBy))
	p.BibiteGroupTempCreatedAt = time.Now()
}

func (p *BibiteGroupTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	case "insertupdate":
		if p.BibiteGroupTempCode == "" {
			return errors.New("Negara Wajib diisi")
		}
		if p.BibiteGroupTempName == "" {
			return errors.New("Provinsi Wajib diisi")
		}
		if p.BibiteGroupTempRole == "" {
			return errors.New("Kabupaten / Kota Wajib diisi")
		}
		if p.BibiteGroupTempType == 0 {
			return errors.New("Kecamatan Wajib diisi")
		}
		if p.BibiteGroupTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.BibiteGroupTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	case "reason":
		if p.BibiteGroupTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.BibiteGroupTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	default:
		if p.BibiteGroupTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil
	}
}

func (p *BibiteGroupTemp) SaveBibiteGroupTemp(db *gorm.DB) (*BibiteGroupTemp, error) {
	var err error
	err = db.Debug().Model(&BibiteGroupTemp{}).Create(&p).Error
	if err != nil {
		return &BibiteGroupTemp{}, err
	}
	return p, nil
}

func (p *BibiteGroupTempPend) SaveBibiteGroupTempPend(db *gorm.DB) (*BibiteGroupTempPend, error) {
	var err error
	err = db.Debug().Model(&BibiteGroupTempPend{}).Create(&p).Error
	if err != nil {
		return &BibiteGroupTempPend{}, err
	}
	return p, nil
}

func (p *BibiteGroupTemp) FindAllBibiteGroupTemp(db *gorm.DB) (*[]BibiteGroupTempPendData, error) {
	var err error
	bibitegroup := []BibiteGroupTempPendData{}
	err = db.Debug().Model(&BibiteGroupTempPendData{}).Limit(100).
		Select(`m_bibite_group_temp.bbt_group_temp_id,
				m_bibite_group_temp.bbt_group_id,
				m_bibite_group_temp.bbt_group_temp_code,
				m_bibite_group_temp.bbt_group_temp_name,
				m_bibite_group_temp.bbt_group_temp_role,
				m_bibite_group_temp.bbt_group_temp_type,
				m_bibite_group_temp.bbt_group_temp_reason,
				m_bibite_group_temp.bbt_group_temp_status,
				m_bibite_group_temp.bbt_group_temp_action_before,
				m_bibite_group_temp.bbt_group_temp_created_by,
				m_bibite_group_temp.bbt_group_temp_created_at at time zone 'Asia/Jakarta' as bbt_group_temp_created_at`).
		Order("bbt_group_temp_created_at desc").
		Find(&bibitegroup).Error
	if err != nil {
		return &[]BibiteGroupTempPendData{}, err
	}
	return &bibitegroup, nil
}

func (p *BibiteGroupTemp) FindAllBibiteGroupTempStatusPendingActive(db *gorm.DB) (*[]BibiteGroupTempPendData, error) {
	var err error
	bibitegroup := []BibiteGroupTempPendData{}
	err = db.Debug().Model(&BibiteGroupTempPendData{}).Limit(100).
		Select(`m_bibite_group_temp.bbt_group_temp_id,
				m_bibite_group_temp.bbt_group_id,
				m_bibite_group_temp.bbt_group_temp_code,
				m_bibite_group_temp.bbt_group_temp_name,
				m_bibite_group_temp.bbt_group_temp_role,
				m_bibite_group_temp.bbt_group_temp_type,
				m_bibite_group_temp.bbt_group_temp_reason,
				m_bibite_group_temp.bbt_group_temp_status,
				m_bibite_group_temp.bbt_group_temp_action_before,
				m_bibite_group_temp.bbt_group_temp_created_by,
				m_bibite_group_temp.bbt_group_temp_created_at at time zone 'Asia/Jakarta' as bbt_group_temp_created_at`).
		Where("bbt_group_temp_status = ?", 11).
		Order("bbt_group_temp_created_at desc").
		Find(&bibitegroup).Error
	if err != nil {
		return &[]BibiteGroupTempPendData{}, err
	}
	return &bibitegroup, nil
}

func (p *BibiteGroupTemp) FindAllBibiteGroupTempStatus(db *gorm.DB, status uint64) (*[]BibiteGroupTempDatas, error) {
	var err error
	bibitegroup := []BibiteGroupTempDatas{}
	err = db.Debug().Model(&BibiteGroupTempData{}).Limit(100).
		Select(`m_bibite_group_temp.bbt_group_temp_id,
				m_bibite_group_temp.bbt_group_temp_code,
				m_bibite_group_temp.bbt_group_temp_name,
				m_bibite_group_temp.bbt_group_temp_role,
				m_bibite_group_temp.bbt_group_temp_type,
				m_bibite_group_temp.bbt_group_temp_reason,
				m_bibite_group_temp.bbt_group_temp_status,
				m_bibite_group_temp.bbt_group_temp_action_before,
				m_bibite_group_temp.bbt_group_temp_created_by,
				m_bibite_group_temp.bbt_group_temp_created_at at time zone 'Asia/Jakarta' as bbt_group_temp_created_at,
				m_bibite_group.bbt_group_id,
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
		Joins("JOIN m_bibite_group on m_bibite_group_temp.bbt_group_id=m_bibite_group.bbt_group_id").
		Where("bbt_group_temp_status = ?", status).
		Order("bbt_group_temp_created_at desc").
		Find(&bibitegroup).Error
	if err != nil {
		return &[]BibiteGroupTempDatas{}, err
	}
	return &bibitegroup, nil
}

func (p *BibiteGroupTemp) FindBibiteGroupTempDataByID(db *gorm.DB, pid uint64) (*BibiteGroupTemp, error) {
	var err error
	err = db.Debug().Model(&BibiteGroupTemp{}).Where("bbt_group_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &BibiteGroupTemp{}, err
	}
	return p, nil
}

func (p *BibiteGroupTemp) FindBibiteGroupTempByIDPendingActive(db *gorm.DB, pid uint64) (*BibiteGroupTempPendData, error) {
	var err error
	bibitegroup := BibiteGroupTempPendData{}
	err = db.Debug().Model(&BibiteGroupTempPendData{}).Limit(100).
		Select(`m_bibite_group_temp.bbt_group_temp_id,
				m_bibite_group_temp.bbt_group_id,
				m_bibite_group_temp.bbt_group_temp_code,
				m_bibite_group_temp.bbt_group_temp_name,
				m_bibite_group_temp.bbt_group_temp_role,
				m_bibite_group_temp.bbt_group_temp_type,
				m_bibite_group_temp.bbt_group_temp_reason,
				m_bibite_group_temp.bbt_group_temp_status,
				m_bibite_group_temp.bbt_group_temp_action_before,
				m_bibite_group_temp.bbt_group_temp_created_by,
				m_bibite_group_temp.bbt_group_temp_created_at at time zone 'Asia/Jakarta' as bbt_group_temp_created_at`).
		Where("bbt_group_temp_id = ?", pid).
		Order("bbt_group_temp_created_at desc").
		Find(&bibitegroup).Error
	if err != nil {
		return &BibiteGroupTempPendData{}, err
	}
	return &bibitegroup, nil
}

func (p *BibiteGroupTemp) FindBibiteGroupTempStatusByIDPendingActive(db *gorm.DB, pid uint64) (*BibiteGroupTempPendData, error) {
	var err error
	bibitegroup := BibiteGroupTempPendData{}
	err = db.Debug().Model(&BibiteGroupTempPendData{}).Limit(100).
		Select(`m_bibite_group_temp.bbt_group_temp_id,
				m_bibite_group_temp.bbt_group_id,
				m_bibite_group_temp.bbt_group_temp_code,
				m_bibite_group_temp.bbt_group_temp_name,
				m_bibite_group_temp.bbt_group_temp_role,
				m_bibite_group_temp.bbt_group_temp_type,
				m_bibite_group_temp.bbt_group_temp_reason,
				m_bibite_group_temp.bbt_group_temp_status,
				m_bibite_group_temp.bbt_group_temp_action_before,
				m_bibite_group_temp.bbt_group_temp_created_by,
				m_bibite_group_temp.bbt_group_temp_created_at at time zone 'Asia/Jakarta' as bbt_group_temp_created_at`).
		Where("bbt_group_temp_id = ? AND bbt_group_temp_status = ?", pid, 11).
		Order("bbt_group_temp_created_at desc").
		Find(&bibitegroup).Error
	if err != nil {
		return &BibiteGroupTempPendData{}, err
	}
	return &bibitegroup, nil
}

func (p *BibiteGroupTemp) FindBibiteGroupTempByID(db *gorm.DB, pid uint64) (*BibiteGroupTempData, error) {
	var err error
	bibitegroup := BibiteGroupTempData{}
	err = db.Debug().Model(&BibiteGroupTempData{}).Limit(100).
		Select(`m_bibite_group_temp.bbt_group_temp_id,
				m_bibite_group_temp.bbt_group_temp_code,
				m_bibite_group_temp.bbt_group_temp_name,
				m_bibite_group_temp.bbt_group_temp_role,
				m_bibite_group_temp.bbt_group_temp_type,
				m_bibite_group_temp.bbt_group_temp_reason,
				m_bibite_group_temp.bbt_group_temp_status,
				m_bibite_group_temp.bbt_group_temp_action_before,
				m_bibite_group_temp.bbt_group_temp_created_by,
				m_bibite_group_temp.bbt_group_temp_created_at at time zone 'Asia/Jakarta' as bbt_group_temp_created_at,
				m_bibite_group.bbt_group_id,
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
		Joins("JOIN m_bibite_group on m_bibite_group_temp.bbt_group_id=m_bibite_group.bbt_group_id").
		Where("bbt_group_temp_id = ?", pid).
		Order("bbt_group_temp_created_at desc").
		Find(&bibitegroup).Error
	if err != nil {
		return &BibiteGroupTempData{}, err
	}
	return &bibitegroup, nil
}

func (p *BibiteGroupTemp) FindBibiteGroupTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*BibiteGroupTempData, error) {
	var err error
	bibitegroup := BibiteGroupTempData{}
	err = db.Debug().Model(&BibiteGroupTempData{}).Limit(100).
		Select(`m_bibite_group_temp.bbt_group_temp_id,
				m_bibite_group_temp.bbt_group_temp_code,
				m_bibite_group_temp.bbt_group_temp_name,
				m_bibite_group_temp.bbt_group_temp_role,
				m_bibite_group_temp.bbt_group_temp_type,
				m_bibite_group_temp.bbt_group_temp_reason,
				m_bibite_group_temp.bbt_group_temp_status,
				m_bibite_group_temp.bbt_group_temp_action_before,
				m_bibite_group_temp.bbt_group_temp_created_by,
				m_bibite_group_temp.bbt_group_temp_created_at at time zone 'Asia/Jakarta' as bbt_group_temp_created_at,
				m_bibite_group.bbt_group_id,
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
		Joins("JOIN m_bibite_group on m_bibite_group_temp.bbt_group_id=m_bibite_group.bbt_group_id").
		Where("bbt_group_temp_id = ? AND bbt_group_temp_status = ?", pid, status).
		Order("bbt_group_temp_created_at desc").
		Find(&bibitegroup).Error
	if err != nil {
		return &BibiteGroupTempData{}, err
	}
	return &bibitegroup, nil
}

func (p *BibiteGroupTemp) UpdateBibiteGroupTemp(db *gorm.DB, pid uint64) (*BibiteGroupTemp, error) {
	var err error
	db = db.Debug().Model(&BibiteGroupTemp{}).Where("bbt_group_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_group_temp_code":   p.BibiteGroupTempCode,
			"bbt_group_temp_name":   p.BibiteGroupTempName,
			"bbt_group_temp_role":   p.BibiteGroupTempRole,
			"bbt_group_temp_type":   p.BibiteGroupTempType,
			"bbt_group_temp_reason": p.BibiteGroupTempReason,
		},
	)
	err = db.Debug().Model(&BibiteGroupTemp{}).Where("bbt_group_temp_id = ?", pid).Error
	if err != nil {
		return &BibiteGroupTemp{}, err
	}
	return p, nil
}

func (p *BibiteGroupTemp) UpdateBibiteGroupTempStatus(db *gorm.DB, pid uint64) (*BibiteGroupTemp, error) {
	var err error
	db = db.Debug().Model(&BibiteGroupTemp{}).Where("bbt_group_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_group_temp_reason": p.BibiteGroupTempReason,
			"bbt_group_temp_status": p.BibiteGroupTempStatus,
		},
	)
	err = db.Debug().Model(&BibiteGroupTemp{}).Where("bbt_group_temp_id = ?", pid).Error
	if err != nil {
		return &BibiteGroupTemp{}, err
	}
	return p, nil
}

func (p *BibiteGroupTemp) DeleteBibiteGroupTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&BibiteGroupTemp{}).Where("bbt_group_temp_id = ?", pid).Take(&BibiteGroupTemp{}).Delete(&BibiteGroupTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("BibiteGroupTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

type BibiteGroupTempPrivilegeTemp struct {
	BibitePrivilegeTempID uint64 `json:"bbt_priv_temp_id"`
}

type BibiteGroupTempPrivilege struct {
	BibiteGroupTempID     uint64 `gorm:"column:bbt_group_temp_id" json:"bbt_group_temp_id"`
	BibitePrivilegeTempID uint64 `gorm:"column:bbt_priv_temp_id" json:"bbt_priv_temp_id"`
}

type BibiteGroupTempPrivilegeData struct {
	BibiteGroupTempPrivilegeID      uint64 `gorm:"column:bbt_group_temp_priv_id" json:"bbt_group_temp_priv_id"`
	BibiteGroupTempID               uint64 `gorm:"column:bbt_group_temp_id" json:"bbt_group_temp_id"`
	BibitePrivilegeTempID           uint64 `gorm:"column:bbt_priv_temp_id" json:"bbt_priv_temp_id"`
	BibitePrivilegeTempCode         string `gorm:"column:bbt_priv_temp_code" json:"bbt_priv_temp_code"`
	BibitePrivilegeTempName         string `gorm:"column:bbt_priv_temp_name" json:"bbt_priv_temp_name"`
	BibitePrivilegeCategoryTempID   uint64 `gorm:"column:bbt_priv_ctgry_temp_id" json:"bbt_priv_ctgry_temp_id"`
	BibitePrivilegeCategoryTempCode string `gorm:"column:bbt_priv_ctgry_temp_code" json:"bbt_priv_ctgry_temp_code"`
	BibitePrivilegeCategoryTempName string `gorm:"column:bbt_priv_ctgry_temp_name" json:"bbt_priv_ctgry_temp_name"`
}

func (BibiteGroupTempPrivilege) TableName() string {
	return "m_bibite_group_temp_privilege"
}

func (BibiteGroupTempPrivilegeData) TableName() string {
	return "m_bibite_group_temp_privilege"
}

func (p *BibiteGroupTempPrivilege) Prepare() {
	p.BibiteGroupTempID = p.BibiteGroupTempID
	p.BibitePrivilegeTempID = p.BibitePrivilegeTempID
}

func (p *BibiteGroupTempPrivilege) SaveBibiteGroupTempPrivilege(db *gorm.DB) (*BibiteGroupTempPrivilege, error) {
	var err error
	err = db.Debug().Model(&BibiteGroupTempPrivilege{}).Create(&p).Error
	if err != nil {
		return &BibiteGroupTempPrivilege{}, err
	}
	return p, nil
}

func (p *BibiteGroupTempPrivilege) FindAllBibiteGroupTempPrivilege(db *gorm.DB) (*[]BibiteGroupTempPrivilegeData, error) {
	var err error
	bibitegroup := []BibiteGroupTempPrivilegeData{}
	err = db.Debug().Model(&BibiteGroupTempPrivilegeData{}).Limit(100).
		Select(`m_bibite_group_temp_privilege.bbt_group_temp_priv_id,
				m_bibite_group_temp_privilege.bbt_group_temp_id,
				m_bibite_privilege.bbt_priv_id as bbt_priv_temp_id,
				m_bibite_privilege.bbt_priv_code as bbt_priv_temp_code,
				m_bibite_privilege.bbt_priv_name as bbt_priv_temp_name,
				m_bibite_privilege_category.bbt_priv_ctgry_id as bbt_priv_ctgry_temp_id,
				m_bibite_privilege_category.bbt_priv_ctgry_code as bbt_priv_ctgry_temp_code,
				m_bibite_privilege_category.bbt_priv_ctgry_name as bbt_priv_ctgry_temp_name`).
		Joins("JOIN m_bibite_privilege on m_bibite_group_temp_privilege.bbt_priv_temp_id=m_bibite_privilege.bbt_priv_id").
		Joins("JOIN m_bibite_privilege_category on m_bibite_privilege.bbt_priv_ctgry_id=m_bibite_privilege_category.bbt_priv_ctgry_id").
		Find(&bibitegroup).Error
	if err != nil {
		return &[]BibiteGroupTempPrivilegeData{}, err
	}
	return &bibitegroup, nil
}

func (p *BibiteGroupTempPrivilege) FindBibiteGroupTempPrivilegeDataByID(db *gorm.DB, pid uint64) (*BibiteGroupTempPrivilege, error) {
	var err error
	err = db.Debug().Model(&BibiteGroupTempPrivilege{}).Where("bbt_group_temp_priv_id = ?", pid).Take(&p).Error
	if err != nil {
		return &BibiteGroupTempPrivilege{}, err
	}
	return p, nil
}

func (p *BibiteGroupTempPrivilege) FindBibiteGroupTempPrivilegeByID(db *gorm.DB, pid uint64) (*BibiteGroupTempPrivilegeData, error) {
	var err error
	bibitegroup := BibiteGroupTempPrivilegeData{}
	err = db.Debug().Model(&BibiteGroupTempPrivilegeData{}).Limit(100).
		Select(`m_bibite_group_temp_privilege.bbt_group_temp_priv_id,
				m_bibite_group_temp_privilege.bbt_group_temp_id,
				m_bibite_privilege.bbt_priv_id as bbt_priv_temp_id,
				m_bibite_privilege.bbt_priv_code as bbt_priv_temp_code,
				m_bibite_privilege.bbt_priv_name as bbt_priv_temp_name,
				m_bibite_privilege_category.bbt_priv_ctgry_id as bbt_priv_ctgry_temp_id,
				m_bibite_privilege_category.bbt_priv_ctgry_code as bbt_priv_ctgry_temp_code,
				m_bibite_privilege_category.bbt_priv_ctgry_name as bbt_priv_ctgry_temp_name`).
		Joins("JOIN m_bibite_privilege on m_bibite_group_temp_privilege.bbt_priv_temp_id=m_bibite_privilege.bbt_priv_id").
		Joins("JOIN m_bibite_privilege_category on m_bibite_privilege.bbt_priv_ctgry_id=m_bibite_privilege_category.bbt_priv_ctgry_id").
		Where("bbt_group_temp_priv_id = ?", pid).
		Find(&bibitegroup).Error
	if err != nil {
		return &BibiteGroupTempPrivilegeData{}, err
	}
	return &bibitegroup, nil
}

func (p *BibiteGroupTempPrivilege) FindBibiteGroupTempPrivilegeDataByBibiteGroupTempID(db *gorm.DB, pid uint64) (*BibiteGroupTempPrivilege, error) {
	var err error
	err = db.Debug().Model(&BibiteGroupTempPrivilege{}).Where("bbt_group_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &BibiteGroupTempPrivilege{}, err
	}
	return p, nil
}

func (p *BibiteGroupTempPrivilege) FindBibiteGroupTempPrivilegeByBibiteGroupTempID(db *gorm.DB, pid uint64) ([]BibiteGroupTempPrivilegeData, error) {
	var err error
	bibitegroup := []BibiteGroupTempPrivilegeData{}
	err = db.Debug().Model(&BibiteGroupTempPrivilegeData{}).Limit(100).
		Select(`m_bibite_group_temp_privilege.bbt_group_temp_priv_id,
				m_bibite_group_temp_privilege.bbt_group_temp_id,
				m_bibite_privilege.bbt_priv_id as bbt_priv_temp_id,
				m_bibite_privilege.bbt_priv_code as bbt_priv_temp_code,
				m_bibite_privilege.bbt_priv_name as bbt_priv_temp_name,
				m_bibite_privilege_category.bbt_priv_ctgry_id as bbt_priv_ctgry_temp_id,
				m_bibite_privilege_category.bbt_priv_ctgry_code as bbt_priv_ctgry_temp_code,
				m_bibite_privilege_category.bbt_priv_ctgry_name as bbt_priv_ctgry_temp_name`).
		Joins("JOIN m_bibite_privilege on m_bibite_group_temp_privilege.bbt_priv_temp_id=m_bibite_privilege.bbt_priv_id").
		Joins("JOIN m_bibite_privilege_category on m_bibite_privilege.bbt_priv_ctgry_id=m_bibite_privilege_category.bbt_priv_ctgry_id").
		Where("bbt_group_temp_id = ?", pid).
		Find(&bibitegroup).Error
	if err != nil {
		return []BibiteGroupTempPrivilegeData{}, err
	}
	return bibitegroup, nil
}

func (p *BibiteGroupTempPrivilege) UpdateBibiteGroupTempPrivilege(db *gorm.DB, pid uint64) (*BibiteGroupTempPrivilege, error) {
	var err error
	db = db.Debug().Model(&BibiteGroupTempPrivilege{}).Where("bbt_group_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_group_temp_id": p.BibiteGroupTempID,
			"bbt_priv_temp_id":  p.BibitePrivilegeTempID,
		},
	)
	err = db.Debug().Model(&BibiteGroupTempPrivilege{}).Where("bbt_group_temp_id = ?", pid).Error
	if err != nil {
		return &BibiteGroupTempPrivilege{}, err
	}
	return p, nil
}

func (p *BibiteGroupTempPrivilege) DeleteBibiteGroupTempPrivilege(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&BibiteGroupTempPrivilege{}).Where("bbt_group_temp_priv_id = ?", pid).Take(&BibiteGroupTempPrivilege{}).Delete(&BibiteGroupTempPrivilege{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("BibiteGroupTempPrivilege not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
