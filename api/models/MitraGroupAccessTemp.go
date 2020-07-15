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

type MitraGroupAccessTemp struct {
	MitraGroupAccessTempID           uint64                              `gorm:"column:mitra_group_accs_temp_id;primary_key;" json:"mitra_group_accs_temp_id"`
	MitraGroupAccessID               uint64                              `gorm:"column:mitra_group_accs_id" json:"mitra_group_accs_id"`
	MitraTempID                      uint64                              `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	MitraGroupAccessTempCode         string                              `gorm:"column:mitra_group_accs_temp_code;size:25" json:"mitra_group_accs_temp_code"`
	MitraGroupAccessTempName         string                              `gorm:"column:mitra_group_accs_temp_name;size:255" json:"mitra_group_accs_temp_name"`
	MitraGroupAccessTempRole         string                              `gorm:"column:mitra_group_accs_temp_role;size:25" json:"mitra_group_accs_temp_role"`
	MitraGroupAccessTempType         int64                               `gorm:"column:mitra_group_accs_temp_type" json:"mitra_group_accs_temp_type"`
	MitraGroupAccessTempStatus       uint64                              `gorm:"column:mitra_group_accs_temp_status;size:2" json:"mitra_group_accs_temp_status"`
	MitraGroupAccessTempActionBefore uint64                              `gorm:"column:mitra_group_accs_temp_action_before;size:2" json:"mitra_group_accs_temp_action_before"`
	MitraGroupAccessTempReason       string                              `gorm:"column:mitra_group_accs_temp_reason" json:"mitra_group_accs_temp_reason"`
	MitraGroupAccessTempCreatedBy    string                              `gorm:"column:mitra_group_accs_temp_created_by;size:125" json:"mitra_group_accs_temp_created_by"`
	MitraGroupAccessTempCreatedAt    time.Time                           `gorm:"column:mitra_group_accs_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_group_accs_temp_created_at"`
	MitraGroupAccessTempPrivilege    []MitraGroupAccessTempPrivilegeTemp `gorm:"-" json:"mitra_group_accs_temp_privilege"`
}

type MitraGroupAccessTempPend struct {
	MitraGroupAccessTempID           uint64                              `gorm:"column:mitra_group_accs_temp_id;primary_key;" json:"mitra_group_accs_temp_id"`
	MitraGroupAccessID               *uint64                             `gorm:"column:mitra_group_accs_id" json:"mitra_group_accs_id"`
	MitraTempID                      uint64                              `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	MitraGroupAccessTempCode         string                              `gorm:"column:mitra_group_accs_temp_code;size:25" json:"mitra_group_accs_temp_code"`
	MitraGroupAccessTempName         string                              `gorm:"column:mitra_group_accs_temp_name;size:255" json:"mitra_group_accs_temp_name"`
	MitraGroupAccessTempRole         string                              `gorm:"column:mitra_group_accs_temp_role;size:25" json:"mitra_group_accs_temp_role"`
	MitraGroupAccessTempType         int64                               `gorm:"column:mitra_group_accs_temp_type" json:"mitra_group_accs_temp_type"`
	MitraGroupAccessTempStatus       uint64                              `gorm:"column:mitra_group_accs_temp_status;size:2" json:"mitra_group_accs_temp_status"`
	MitraGroupAccessTempActionBefore uint64                              `gorm:"column:mitra_group_accs_temp_action_before;size:2" json:"mitra_group_accs_temp_action_before"`
	MitraGroupAccessTempReason       string                              `gorm:"column:mitra_group_accs_temp_reason" json:"mitra_group_accs_temp_reason"`
	MitraGroupAccessTempCreatedBy    string                              `gorm:"column:mitra_group_accs_temp_created_by;size:125" json:"mitra_group_accs_temp_created_by"`
	MitraGroupAccessTempCreatedAt    time.Time                           `gorm:"column:mitra_group_accs_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_group_accs_temp_created_at"`
	MitraGroupAccessTempPrivilege    []MitraGroupAccessTempPrivilegeTemp `gorm:"-" json:"mitra_group_accs_temp_privilege"`
}

type MitraGroupAccessTempPendData struct {
	MitraGroupAccessTempID           uint64                              `gorm:"column:mitra_group_accs_temp_id;primary_key;" json:"mitra_group_accs_temp_id"`
	MitraGroupAccessID               uint64                              `gorm:"column:mitra_group_accs_id" json:"mitra_group_accs_id"`
	MitraTempID                      uint64                              `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	MitraGroupAccessTempCode         string                              `gorm:"column:mitra_group_accs_temp_code;size:25" json:"mitra_group_accs_temp_code"`
	MitraGroupAccessTempName         string                              `gorm:"column:mitra_group_accs_temp_name;size:255" json:"mitra_group_accs_temp_name"`
	MitraGroupAccessTempRole         string                              `gorm:"column:mitra_group_accs_temp_role;size:25" json:"mitra_group_accs_temp_role"`
	MitraGroupAccessTempType         int64                               `gorm:"column:mitra_group_accs_temp_type" json:"mitra_group_accs_temp_type"`
	MitraGroupAccessTempStatus       uint64                              `gorm:"column:mitra_group_accs_temp_status;size:2" json:"mitra_group_accs_temp_status"`
	MitraGroupAccessTempActionBefore uint64                              `gorm:"column:mitra_group_accs_temp_action_before;size:2" json:"mitra_group_accs_temp_action_before"`
	MitraGroupAccessTempReason       string                              `gorm:"column:mitra_group_accs_temp_reason" json:"mitra_group_accs_temp_reason"`
	MitraGroupAccessTempCreatedBy    string                              `gorm:"column:mitra_group_accs_temp_created_by;size:125" json:"mitra_group_accs_temp_created_by"`
	MitraGroupAccessTempCreatedAt    time.Time                           `gorm:"column:mitra_group_accs_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_group_accs_temp_created_at"`
	MitraGroupAccessTempPrivilege    []MitraGroupAccessTempPrivilegeData `gorm:"-" json:"mitra_group_accs_temp_privilege"`
}

type MitraGroupAccessTempDatas struct {
	MitraGroupAccessTempID           uint64    `gorm:"column:mitra_group_accs_temp_id" json:"mitra_group_accs_temp_id"`
	MitraTempID                      uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	MitraGroupAccessTempCode         string    `gorm:"column:mitra_group_accs_temp_code" json:"mitra_group_accs_temp_code"`
	MitraGroupAccessTempName         string    `gorm:"column:mitra_group_accs_temp_name" json:"mitra_group_accs_temp_name"`
	MitraGroupAccessTempRole         string    `gorm:"column:mitra_group_accs_temp_role" json:"mitra_group_accs_temp_role"`
	MitraGroupAccessTempType         uint64    `gorm:"column:mitra_group_accs_temp_type" json:"mitra_group_accs_temp_type"`
	MitraGroupAccessTempReason       string    `gorm:"column:mitra_group_accs_temp_reason" json:"mitra_group_accs_temp_reason"`
	MitraGroupAccessTempStatus       uint64    `gorm:"column:mitra_group_accs_temp_status;size:2" json:"mitra_group_accs_temp_status"`
	MitraGroupAccessTempActionBefore uint64    `gorm:"column:mitra_group_accs_temp_action_before;size:2" json:"mitra_group_accs_temp_action_before"`
	MitraGroupAccessTempCreatedBy    string    `gorm:"column:mitra_group_accs_temp_created_by;size:125" json:"mitra_group_accs_temp_created_by"`
	MitraGroupAccessTempCreatedAt    time.Time `gorm:"column:mitra_group_accs_temp_created_at" json:"mitra_group_accs_temp_created_at"`
	MitraGroupAccessID               uint64    `gorm:"column:mitra_group_accs_id" json:"mitra_group_accs_id"`
	MitraGroupAccessCode             string    `gorm:"column:mitra_group_accs_code" json:"mitra_group_accs_code"`
	MitraGroupAccessName             string    `gorm:"column:mitra_group_accs_name" json:"mitra_group_accs_name"`
	MitraGroupAccessRole             string    `gorm:"column:mitra_group_accs_role" json:"mitra_group_accs_role"`
	MitraGroupAccessType             string    `gorm:"column:mitra_group_accs_type" json:"mitra_group_accs_type"`
	MitraGroupAccessStatus           uint64    `gorm:"column:mitra_group_accs_status" json:"mitra_group_accs_status"`
	MitraGroupAccessCreatedBy        string    `gorm:"column:mitra_group_accs_created_by" json:"mitra_group_accs_created_by"`
	MitraGroupAccessCreatedAt        time.Time `gorm:"column:mitra_group_accs_created_at" json:"mitra_group_accs_created_at"`
	MitraGroupAccessUpdatedBy        string    `gorm:"column:mitra_group_accs_updated_by" json:"mitra_group_accs_updated_by"`
	MitraGroupAccessUpdatedAt        time.Time `gorm:"column:mitra_group_accs_updated_at" json:"mitra_group_accs_updated_at"`
	MitraGroupAccessDeletedBy        string    `gorm:"column:mitra_group_accs_deleted_by" json:"mitra_group_accs_deleted_by"`
	MitraGroupAccessDeletedAt        time.Time `gorm:"column:mitra_group_accs_deleted_at" json:"mitra_group_accs_deleted_at"`
}

type MitraGroupAccessTempData struct {
	MitraGroupAccessTempID           uint64                              `gorm:"column:mitra_group_accs_temp_id" json:"mitra_group_accs_temp_id"`
	MitraTempID                      uint64                              `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	MitraTempCode                    string                              `gorm:"column:mitra_temp_code;size:25" json:"mitra_temp_code"`
	MitraTempName                    string                              `gorm:"column:mitra_temp_name;size:255" json:"mitra_temp_name"`
	MitraTempWebsite                 string                              `gorm:"column:mitra_temp_website;size:255" json:"mitra_temp_website"`
	MitraTempEmail                   string                              `gorm:"column:mitra_temp_email;size:255" json:"mitra_temp_email"`
	MitraTempLogo                    string                              `gorm:"column:mitra_temp_logo;size:255" json:"mitra_temp_logo"`
	MitraGroupAccessID               uint64                              `gorm:"column:mitra_group_accs_id" json:"mitra_group_accs_id"`
	MitraGroupAccessCode             string                              `gorm:"column:mitra_group_accs_code;size:25" json:"mitra_group_accs_code"`
	MitraGroupAccessName             string                              `gorm:"column:mitra_group_accs_name;size:255" json:"mitra_group_accs_name"`
	MitraGroupAccessRole             string                              `gorm:"column:mitra_group_accs_role;size:25" json:"mitra_group_accs_role"`
	MitraGroupAccessType             int64                               `gorm:"column:mitra_group_accs_type" json:"mitra_group_accs_type"`
	MitraGroupAccessStatus           uint64                              `gorm:"column:mitra_group_accs_status;size:2" json:"mitra_group_accs_status"`
	MitraGroupAccessCreatedAt        time.Time                           `gorm:"column:mitra_group_accs_created_at;default:CURRENT_TIMESTAMP" json:"mitra_group_accs_created_at"`
	MitraGroupAccessPrivilege        []MitraGroupAccessPrivilegeData     `json:"mitra_group_accs_privilege"`
	MitraGroupAccessTempCode         string                              `gorm:"column:mitra_group_accs_temp_code;size:25" json:"mitra_group_accs_temp_code"`
	MitraGroupAccessTempName         string                              `gorm:"column:mitra_group_accs_temp_name;size:255" json:"mitra_group_accs_temp_name"`
	MitraGroupAccessTempRole         string                              `gorm:"column:mitra_group_accs_temp_role;size:25" json:"mitra_group_accs_temp_role"`
	MitraGroupAccessTempType         int64                               `gorm:"column:mitra_group_accs_temp_type" json:"mitra_group_accs_temp_type"`
	MitraGroupAccessTempStatus       uint64                              `gorm:"column:mitra_group_accs_temp_status;size:2" json:"mitra_group_accs_temp_status"`
	MitraGroupAccessTempReason       string                              `gorm:"column:mitra_group_accs_temp_reason" json:"mitra_group_accs_temp_reason"`
	MitraGroupAccessTempActionBefore uint64                              `gorm:"column:mitra_group_accs_temp_action_before;size:2" json:"mitra_group_accs_temp_action_before"`
	MitraGroupAccessTempCreatedBy    string                              `gorm:"column:mitra_group_accs_temp_created_by;size:125" json:"mitra_group_accs_temp_created_by"`
	MitraGroupAccessTempCreatedAt    time.Time                           `gorm:"column:mitra_group_accs_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_group_accs_temp_created_at"`
	MitraGroupAccessTempPrivilege    []MitraGroupAccessTempPrivilegeData `json:"mitra_group_accs_temp_privilege"`
}

type ResponseMitraGroupAccessTemps struct {
	Status  int                          `json:"status"`
	Message string                       `json:"message"`
	Data    *[]MitraGroupAccessTempDatas `json:"data"`
}

type ResponseMitraGroupAccessTemp struct {
	Status  int                       `json:"status"`
	Message string                    `json:"message"`
	Data    *MitraGroupAccessTempData `json:"data"`
}

type ResponseMitraGroupAccessTempsPend struct {
	Status  int                             `json:"status"`
	Message string                          `json:"message"`
	Data    *[]MitraGroupAccessTempPendData `json:"data"`
}

type ResponseMitraGroupAccessTempPend struct {
	Status  int                           `json:"status"`
	Message string                        `json:"message"`
	Data    *MitraGroupAccessTempPendData `json:"data"`
}

type ResponseMitraGroupAccessTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraGroupAccessTemp) TableName() string {
	return "m_mitra_group_access_temp"
}

func (MitraGroupAccessTempPend) TableName() string {
	return "m_mitra_group_access_temp"
}

func (MitraGroupAccessTempPendData) TableName() string {
	return "m_mitra_group_access_temp"
}

func (MitraGroupAccessTempData) TableName() string {
	return "m_mitra_group_access_temp"
}

func (MitraGroupAccessTempDatas) TableName() string {
	return "m_mitra_group_access_temp"
}

func (p *MitraGroupAccessTemp) Prepare() {
	p.MitraGroupAccessID = p.MitraGroupAccessID
	p.MitraTempID = p.MitraTempID
	p.MitraGroupAccessTempCode = html.EscapeString(strings.TrimSpace(p.MitraGroupAccessTempCode))
	p.MitraGroupAccessTempName = html.EscapeString(strings.TrimSpace(p.MitraGroupAccessTempName))
	p.MitraGroupAccessTempRole = html.EscapeString(strings.TrimSpace(p.MitraGroupAccessTempRole))
	p.MitraGroupAccessTempReason = html.EscapeString(strings.TrimSpace(p.MitraGroupAccessTempReason))
	p.MitraGroupAccessTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraGroupAccessTempCreatedBy))
	p.MitraGroupAccessTempType = p.MitraGroupAccessTempType
	p.MitraGroupAccessTempStatus = p.MitraGroupAccessTempStatus
	p.MitraGroupAccessTempCreatedAt = time.Now()
}

func (p *MitraGroupAccessTempPend) Prepare() {
	p.MitraGroupAccessID = nil
	p.MitraTempID = p.MitraTempID
	p.MitraGroupAccessTempCode = html.EscapeString(strings.TrimSpace(p.MitraGroupAccessTempCode))
	p.MitraGroupAccessTempName = html.EscapeString(strings.TrimSpace(p.MitraGroupAccessTempName))
	p.MitraGroupAccessTempRole = html.EscapeString(strings.TrimSpace(p.MitraGroupAccessTempRole))
	p.MitraGroupAccessTempReason = html.EscapeString(strings.TrimSpace(p.MitraGroupAccessTempReason))
	p.MitraGroupAccessTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraGroupAccessTempCreatedBy))
	p.MitraGroupAccessTempType = p.MitraGroupAccessTempType
	p.MitraGroupAccessTempStatus = p.MitraGroupAccessTempStatus
	p.MitraGroupAccessTempCreatedAt = time.Now()
}

func (p *MitraGroupAccessTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.MitraTempID == 0 {
			return errors.New("Required MitraGroupAccessTemp Code")
		}
		if p.MitraGroupAccessTempCode == "" {
			return errors.New("Required MitraGroupAccessTemp Code")
		}
		if p.MitraGroupAccessTempName == "" {
			return errors.New("Required MitraGroupAccessTemp Name")
		}
		if p.MitraGroupAccessTempRole == "" {
			return errors.New("Required MitraGroupAccessTemp Name")
		}
		if p.MitraGroupAccessTempType == 0 {
			return errors.New("Required MitraGroupAccessTemp Name")
		}
		return nil
	}
}

func (p *MitraGroupAccessTemp) SaveMitraGroupAccessTemp(db *gorm.DB) (*MitraGroupAccessTemp, error) {
	var err error
	err = db.Debug().Model(&MitraGroupAccessTemp{}).Create(&p).Error
	if err != nil {
		return &MitraGroupAccessTemp{}, err
	}
	return p, nil
}

func (p *MitraGroupAccessTempPend) SaveMitraGroupAccessTempPend(db *gorm.DB) (*MitraGroupAccessTempPend, error) {
	var err error
	err = db.Debug().Model(&MitraGroupAccessTempPend{}).Create(&p).Error
	if err != nil {
		return &MitraGroupAccessTempPend{}, err
	}
	return p, nil
}

func (p *MitraGroupAccessTemp) FindAllMitraGroupAccessTemp(db *gorm.DB, mitra uint64) (*[]MitraGroupAccessTempPendData, error) {
	var err error
	mitragroup := []MitraGroupAccessTempPendData{}
	err = db.Debug().Model(&MitraGroupAccessTempPendData{}).Limit(100).
		Select(`m_mitra_group_access_temp.mitra_group_accs_temp_id,
				m_mitra_group_access_temp.mitra_group_accs_id,
				m_mitra_group_access_temp.mitra_temp_id,
				m_mitra_group_access_temp.mitra_group_accs_temp_code,
				m_mitra_group_access_temp.mitra_group_accs_temp_name,
				m_mitra_group_access_temp.mitra_group_accs_temp_role,
				m_mitra_group_access_temp.mitra_group_accs_temp_type,
				m_mitra_group_access_temp.mitra_group_accs_temp_status,
				m_mitra_group_access_temp.mitra_group_accs_temp_action_before,
				m_mitra_group_access_temp.mitra_group_accs_temp_reason,
				m_mitra_group_access_temp.mitra_group_accs_temp_created_by,
				m_mitra_group_access_temp.mitra_group_accs_temp_created_at`).
		Where("mitra_temp_id = ?", mitra).
		Find(&mitragroup).Error
	if err != nil {
		return &[]MitraGroupAccessTempPendData{}, err
	}
	return &mitragroup, nil
}

func (p *MitraGroupAccessTemp) FindAllMitraGroupAccessTempStatusPendingActive(db *gorm.DB, mitra uint64) (*[]MitraGroupAccessTempPendData, error) {
	var err error
	mitragroup := []MitraGroupAccessTempPendData{}
	err = db.Debug().Model(&MitraGroupAccessTempPendData{}).Limit(100).
		Select(`m_mitra_group_access_temp.mitra_group_accs_temp_id,
				m_mitra_group_access_temp.mitra_group_accs_id,
				m_mitra_group_access_temp.mitra_temp_id,
				m_mitra_group_access_temp.mitra_group_accs_temp_code,
				m_mitra_group_access_temp.mitra_group_accs_temp_name,
				m_mitra_group_access_temp.mitra_group_accs_temp_role,
				m_mitra_group_access_temp.mitra_group_accs_temp_type,
				m_mitra_group_access_temp.mitra_group_accs_temp_status,
				m_mitra_group_access_temp.mitra_group_accs_temp_action_before,
				m_mitra_group_access_temp.mitra_group_accs_temp_reason,
				m_mitra_group_access_temp.mitra_group_accs_temp_created_by,
				m_mitra_group_access_temp.mitra_group_accs_temp_created_at`).
		Where("mitra_group_accs_temp_status = ? AND mitra_temp_id = ?", 11, mitra).
		Find(&mitragroup).Error
	if err != nil {
		return &[]MitraGroupAccessTempPendData{}, err
	}
	return &mitragroup, nil
}

func (p *MitraGroupAccessTemp) FindAllMitraGroupAccessTempStatus(db *gorm.DB, status uint64, mitra uint64) (*[]MitraGroupAccessTempDatas, error) {
	var err error
	mitragroup := []MitraGroupAccessTempDatas{}
	err = db.Debug().Model(&MitraGroupAccessTempDatas{}).Limit(100).
		Select(`m_mitra_group_access_temp.mitra_group_accs_temp_id,
				m_mitra_group_access_temp.mitra_temp_id,
				m_mitra_temp.mitra_code as mitra_temp_code,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_temp.mitra_website as mitra_temp_website,
				m_mitra_temp.mitra_email as mitra_temp_email,
				m_mitra_temp.mitra_logo as mitra_temp_logo,
				m_mitra_group_access_temp.mitra_group_accs_id,
				m_mitra_group_access.mitra_group_accs_code,
				m_mitra_group_access.mitra_group_accs_name,
				m_mitra_group_access.mitra_group_accs_role,
				m_mitra_group_access.mitra_group_accs_type,
				m_mitra_group_access.mitra_group_accs_status,
				m_mitra_group_access.mitra_group_accs_created_at,
				m_mitra_group_access_temp.mitra_group_accs_temp_code,
				m_mitra_group_access_temp.mitra_group_accs_temp_name,
				m_mitra_group_access_temp.mitra_group_accs_temp_role,
				m_mitra_group_access_temp.mitra_group_accs_temp_type,
				m_mitra_group_access_temp.mitra_group_accs_temp_status,
				m_mitra_group_access_temp.mitra_group_accs_temp_action_before,
				m_mitra_group_access_temp.mitra_group_accs_temp_reason,
				m_mitra_group_access_temp.mitra_group_accs_temp_created_by,
				m_mitra_group_access_temp.mitra_group_accs_temp_created_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_group_access_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_group_access on m_mitra_group_access_temp.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Where("mitra_group_accs_temp_status = ? AND mitra_temp_id = ?", status, mitra).
		Find(&mitragroup).Error
	if err != nil {
		return &[]MitraGroupAccessTempDatas{}, err
	}
	return &mitragroup, nil
}

func (p *MitraGroupAccessTemp) FindMitraGroupAccessTempDataByID(db *gorm.DB, pid uint64) (*MitraGroupAccessTemp, error) {
	var err error
	err = db.Debug().Model(&MitraGroupAccessTemp{}).Where("mitra_group_accs_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraGroupAccessTemp{}, err
	}
	return p, nil
}

func (p *MitraGroupAccessTemp) FindMitraGroupAccessTempByIDPendingActive(db *gorm.DB, pid uint64) (*MitraGroupAccessTempPendData, error) {
	var err error
	mitragroup := MitraGroupAccessTempPendData{}
	err = db.Debug().Model(&MitraGroupAccessTempPendData{}).Limit(100).
		Select(`m_mitra_group_access_temp.mitra_group_accs_temp_id,
				m_mitra_group_access_temp.mitra_group_accs_id,
				m_mitra_group_access_temp.mitra_temp_id,
				m_mitra_group_access_temp.mitra_group_accs_temp_code,
				m_mitra_group_access_temp.mitra_group_accs_temp_name,
				m_mitra_group_access_temp.mitra_group_accs_temp_role,
				m_mitra_group_access_temp.mitra_group_accs_temp_type,
				m_mitra_group_access_temp.mitra_group_accs_temp_status,
				m_mitra_group_access_temp.mitra_group_accs_temp_action_before,
				m_mitra_group_access_temp.mitra_group_accs_temp_reason,
				m_mitra_group_access_temp.mitra_group_accs_temp_created_by,
				m_mitra_group_access_temp.mitra_group_accs_temp_created_at`).
		Where("mitra_group_accs_temp_id = ?", pid).
		Find(&mitragroup).Error
	if err != nil {
		return &MitraGroupAccessTempPendData{}, err
	}
	return &mitragroup, nil
}

func (p *MitraGroupAccessTemp) FindMitraGroupAccessTempStatusByIDPendingActive(db *gorm.DB, pid uint64, mitra uint64) (*MitraGroupAccessTempPendData, error) {
	var err error
	mitragroup := MitraGroupAccessTempPendData{}
	err = db.Debug().Model(&MitraGroupAccessTempPendData{}).Limit(100).
		Select(`m_mitra_group_access_temp.mitra_group_accs_temp_id,
				m_mitra_group_access_temp.mitra_group_accs_id,
				m_mitra_group_access_temp.mitra_temp_id,
				m_mitra_group_access_temp.mitra_group_accs_temp_code,
				m_mitra_group_access_temp.mitra_group_accs_temp_name,
				m_mitra_group_access_temp.mitra_group_accs_temp_role,
				m_mitra_group_access_temp.mitra_group_accs_temp_type,
				m_mitra_group_access_temp.mitra_group_accs_temp_status,
				m_mitra_group_access_temp.mitra_group_accs_temp_action_before,
				m_mitra_group_access_temp.mitra_group_accs_temp_reason,
				m_mitra_group_access_temp.mitra_group_accs_temp_created_by,
				m_mitra_group_access_temp.mitra_group_accs_temp_created_at`).
		Where("mitra_group_accs_temp_id = ? AND mitra_group_accs_temp_status = ? AND mitra_temp_id = ?", pid, 11, mitra).
		Find(&mitragroup).Error
	if err != nil {
		return &MitraGroupAccessTempPendData{}, err
	}
	return &mitragroup, nil
}

func (p *MitraGroupAccessTemp) FindMitraGroupAccessTempByID(db *gorm.DB, pid uint64) (*MitraGroupAccessTempData, error) {
	var err error
	mitragroup := MitraGroupAccessTempData{}
	err = db.Debug().Model(&MitraGroupAccessTempData{}).
		Select(`m_mitra_group_access_temp.mitra_group_accs_temp_id,
				m_mitra_group_access_temp.mitra_temp_id,
				m_mitra_temp.mitra_code as mitra_temp_code,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_temp.mitra_website as mitra_temp_website,
				m_mitra_temp.mitra_email as mitra_temp_email,
				m_mitra_temp.mitra_logo as mitra_temp_logo,
				m_mitra_group_access_temp.mitra_group_accs_id,
				m_mitra_group_access.mitra_group_accs_code,
				m_mitra_group_access.mitra_group_accs_name,
				m_mitra_group_access.mitra_group_accs_role,
				m_mitra_group_access.mitra_group_accs_type,
				m_mitra_group_access.mitra_group_accs_status,
				m_mitra_group_access.mitra_group_accs_created_at,
				m_mitra_group_access_temp.mitra_group_accs_temp_code,
				m_mitra_group_access_temp.mitra_group_accs_temp_name,
				m_mitra_group_access_temp.mitra_group_accs_temp_role,
				m_mitra_group_access_temp.mitra_group_accs_temp_type,
				m_mitra_group_access_temp.mitra_group_accs_temp_status,
				m_mitra_group_access_temp.mitra_group_accs_temp_reason,
				m_mitra_group_access_temp.mitra_group_accs_temp_action_before,
				m_mitra_group_access_temp.mitra_group_accs_temp_created_by,
				m_mitra_group_access_temp.mitra_group_accs_temp_created_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_group_access_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_group_access on m_mitra_group_access_temp.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Where("mitra_group_accs_temp_id = ?", pid).
		Take(&mitragroup).Error
	if err != nil {
		return &MitraGroupAccessTempData{}, err
	}
	return &mitragroup, nil
}

func (p *MitraGroupAccessTemp) FindMitraGroupAccessTempByIDMitra(db *gorm.DB, pid uint64, mitra uint64) (*MitraGroupAccessTempData, error) {
	var err error
	mitragroup := MitraGroupAccessTempData{}
	err = db.Debug().Model(&MitraGroupAccessTempData{}).
		Select(`m_mitra_group_access_temp.mitra_group_accs_temp_id,
				m_mitra_group_access_temp.mitra_temp_id,
				m_mitra_temp.mitra_code as mitra_temp_code,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_temp.mitra_website as mitra_temp_website,
				m_mitra_temp.mitra_email as mitra_temp_email,
				m_mitra_temp.mitra_logo as mitra_temp_logo,
				m_mitra_group_access_temp.mitra_group_accs_id,
				m_mitra_group_access.mitra_group_accs_code,
				m_mitra_group_access.mitra_group_accs_name,
				m_mitra_group_access.mitra_group_accs_role,
				m_mitra_group_access.mitra_group_accs_type,
				m_mitra_group_access.mitra_group_accs_status,
				m_mitra_group_access.mitra_group_accs_created_at,
				m_mitra_group_access_temp.mitra_group_accs_temp_code,
				m_mitra_group_access_temp.mitra_group_accs_temp_name,
				m_mitra_group_access_temp.mitra_group_accs_temp_role,
				m_mitra_group_access_temp.mitra_group_accs_temp_type,
				m_mitra_group_access_temp.mitra_group_accs_temp_status,
				m_mitra_group_access_temp.mitra_group_accs_temp_action_before,
				m_mitra_group_access_temp.mitra_group_accs_temp_reason,
				m_mitra_group_access_temp.mitra_group_accs_temp_created_by,
				m_mitra_group_access_temp.mitra_group_accs_temp_created_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_group_access_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_group_access on m_mitra_group_access_temp.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Where("mitra_group_accs_temp_id = ? AND m_mitra_group_access_temp.mitra_temp_id = ?", pid, mitra).
		Take(&mitragroup).Error
	if err != nil {
		return &MitraGroupAccessTempData{}, err
	}
	return &mitragroup, nil
}

func (p *MitraGroupAccessTemp) FindMitraGroupAccessTempStatusByID(db *gorm.DB, pid uint64, status uint64, mitra uint64) (*MitraGroupAccessTempData, error) {
	var err error
	mitragroup := MitraGroupAccessTempData{}
	err = db.Debug().Model(&MitraGroupAccessTempData{}).
		Select(`m_mitra_group_access_temp.mitra_group_accs_temp_id,
				m_mitra_group_access_temp.mitra_temp_id,
				m_mitra_temp.mitra_code as mitra_temp_code,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_temp.mitra_website as mitra_temp_website,
				m_mitra_temp.mitra_email as mitra_temp_email,
				m_mitra_temp.mitra_logo as mitra_temp_logo,
				m_mitra_group_access_temp.mitra_group_accs_id,
				m_mitra_group_access.mitra_group_accs_code,
				m_mitra_group_access.mitra_group_accs_name,
				m_mitra_group_access.mitra_group_accs_role,
				m_mitra_group_access.mitra_group_accs_type,
				m_mitra_group_access.mitra_group_accs_status,
				m_mitra_group_access.mitra_group_accs_created_at,
				m_mitra_group_access_temp.mitra_group_accs_temp_code,
				m_mitra_group_access_temp.mitra_group_accs_temp_name,
				m_mitra_group_access_temp.mitra_group_accs_temp_role,
				m_mitra_group_access_temp.mitra_group_accs_temp_type,
				m_mitra_group_access_temp.mitra_group_accs_temp_status,
				m_mitra_group_access_temp.mitra_group_accs_temp_action_before,
				m_mitra_group_access_temp.mitra_group_accs_temp_reason,
				m_mitra_group_access_temp.mitra_group_accs_temp_created_by,
				m_mitra_group_access_temp.mitra_group_accs_temp_created_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_group_access_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_group_access on m_mitra_group_access_temp.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Where("mitra_group_accs_temp_id = ? AND mitra_group_accs_temp_status = ? AND mitra_temp_id = ?", pid, status, mitra).
		Take(&mitragroup).Error
	if err != nil {
		return &MitraGroupAccessTempData{}, err
	}
	return &mitragroup, nil
}

func (p *MitraGroupAccessTemp) UpdateMitraGroupAccessTemp(db *gorm.DB, pid uint64) (*MitraGroupAccessTemp, error) {

	var err error
	db = db.Debug().Model(&MitraGroupAccessTemp{}).Where("mitra_group_accs_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_temp_id":                    p.MitraTempID,
			"mitra_group_accs_temp_code":       p.MitraGroupAccessTempCode,
			"mitra_group_accs_temp_name":       p.MitraGroupAccessTempName,
			"mitra_group_accs_temp_role":       p.MitraGroupAccessTempRole,
			"mitra_group_accs_temp_type":       p.MitraGroupAccessTempType,
			"mitra_group_accs_temp_status":     p.MitraGroupAccessTempStatus,
			"mitra_group_accs_temp_reason":     p.MitraGroupAccessTempReason,
			"mitra_group_accs_temp_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraGroupAccessTemp{}).Where("mitra_group_accs_temp_id = ?", pid).Error
	if err != nil {
		return &MitraGroupAccessTemp{}, err
	}
	return p, nil
}

func (p *MitraGroupAccessTemp) UpdateMitraGroupAccessTempStatus(db *gorm.DB, pid uint64) (*MitraGroupAccessTemp, error) {

	var err error
	db = db.Debug().Model(&MitraGroupAccessTemp{}).Where("mitra_group_accs_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_group_accs_temp_status":     p.MitraGroupAccessTempStatus,
			"mitra_group_accs_temp_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraGroupAccessTemp{}).Where("mitra_group_accs_temp_id = ?", pid).Error
	if err != nil {
		return &MitraGroupAccessTemp{}, err
	}
	return p, nil
}

func (p *MitraGroupAccessTemp) DeleteMitraGroupAccessTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraGroupAccessTemp{}).Where("mitra_group_accs_temp_id = ?", pid).Take(&MitraGroupAccessTemp{}).Delete(&MitraGroupAccessTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraGroupAccessTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

type MitraGroupAccessTempPrivilegeTemp struct {
	MitraPrivilegeTempID uint64 `json:"mitra_priv_temp_id"`
}

type MitraGroupAccessTempPrivilege struct {
	MitraGroupAccessTempID uint64 `gorm:"column:mitra_group_accs_temp_id" json:"mitra_group_accs_temp_id"`
	MitraPrivilegeTempID   uint64 `gorm:"column:mitra_priv_temp_id" json:"mitra_priv_temp_id"`
}

type MitraGroupAccessTempPrivilegeData struct {
	MitraGroupAccessTempPrivilegeID uint64 `gorm:"column:mitra_group_accs_temp_priv_id" json:"mitra_group_accs_temp_priv_id"`
	MitraGroupAccessTempID          uint64 `gorm:"column:mitra_group_accs_temp_id" json:"mitra_group_accs_temp_id"`
	MitraPrivilegeTempID            uint64 `gorm:"column:mitra_priv_temp_id" json:"mitra_priv_temp_id"`
	MitraPrivilegeTempCode          string `gorm:"column:mitra_priv_temp_code" json:"mitra_priv_temp_code"`
	MitraPrivilegeTempName          string `gorm:"column:mitra_priv_temp_name" json:"mitra_priv_temp_name"`
	MitraPrivilegeCategoryTempID    uint64 `gorm:"column:mitra_priv_ctgry_temp_id" json:"mitra_priv_ctgry_temp_id"`
	MitraPrivilegeCategoryTempCode  string `gorm:"column:mitra_priv_ctgry_temp_code" json:"mitra_priv_ctgry_temp_code"`
	MitraPrivilegeCategoryTempName  string `gorm:"column:mitra_priv_ctgry_temp_name" json:"mitra_priv_ctgry_temp_name"`
}

func (MitraGroupAccessTempPrivilege) TableName() string {
	return "m_mitra_group_access_temp_privilege"
}

func (MitraGroupAccessTempPrivilegeData) TableName() string {
	return "m_mitra_group_access_temp_privilege"
}

func (p *MitraGroupAccessTempPrivilege) Prepare() {
	p.MitraGroupAccessTempID = p.MitraGroupAccessTempID
	p.MitraPrivilegeTempID = p.MitraPrivilegeTempID
}

func (p *MitraGroupAccessTempPrivilege) SaveMitraGroupAccessTempPrivilege(db *gorm.DB) (*MitraGroupAccessTempPrivilege, error) {
	var err error
	err = db.Debug().Model(&MitraGroupAccessTempPrivilege{}).Create(&p).Error
	if err != nil {
		return &MitraGroupAccessTempPrivilege{}, err
	}
	return p, nil
}

func (p *MitraGroupAccessTempPrivilege) FindAllMitraGroupAccessTempPrivilege(db *gorm.DB) (*[]MitraGroupAccessTempPrivilegeData, error) {
	var err error
	mitragroup := []MitraGroupAccessTempPrivilegeData{}
	err = db.Debug().Model(&MitraGroupAccessTempPrivilegeData{}).Limit(100).
		Select(`m_mitra_group_access_temp_privilege.mitra_group_accs_temp_priv_id,
				m_mitra_group_access_temp_privilege.mitra_group_accs_temp_id,
				m_mitra_privilege.mitra_priv_id as mitra_priv_temp_id,
				m_mitra_privilege.mitra_priv_code as mitra_priv_temp_code,
				m_mitra_privilege.mitra_priv_name as mitra_priv_temp_name,
				m_mitra_privilege_category.mitra_priv_ctgry_id as mitra_priv_ctgry_temp_id,
				m_mitra_privilege_category.mitra_priv_ctgry_code as mitra_priv_ctgry_temp_code,
				m_mitra_privilege_category.mitra_priv_ctgry_name as mitra_priv_ctgry_temp_name`).
		Joins("JOIN m_mitra_privilege on m_mitra_group_access_temp_privilege.mitra_priv_temp_id=m_mitra_privilege.mitra_priv_id").
		Joins("JOIN m_mitra_privilege_category on m_mitra_privilege.mitra_priv_ctgry_id=m_mitra_privilege_category.mitra_priv_ctgry_id").
		Find(&mitragroup).Error
	if err != nil {
		return &[]MitraGroupAccessTempPrivilegeData{}, err
	}
	return &mitragroup, nil
}

func (p *MitraGroupAccessTempPrivilege) FindMitraGroupAccessTempPrivilegeDataByID(db *gorm.DB, pid uint64) (*MitraGroupAccessTempPrivilege, error) {
	var err error
	err = db.Debug().Model(&MitraGroupAccessTempPrivilege{}).Where("mitra_group_accs_temp_priv_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraGroupAccessTempPrivilege{}, err
	}
	return p, nil
}

func (p *MitraGroupAccessTempPrivilege) FindMitraGroupAccessTempPrivilegeByID(db *gorm.DB, pid uint64) (*MitraGroupAccessTempPrivilegeData, error) {
	var err error
	mitragroup := MitraGroupAccessTempPrivilegeData{}
	err = db.Debug().Model(&MitraGroupAccessTempPrivilegeData{}).Limit(100).
		Select(`m_mitra_group_access_temp_privilege.mitra_group_accs_temp_priv_id,
				m_mitra_group_access_temp_privilege.mitra_group_accs_temp_id,
				m_mitra_privilege.mitra_priv_id as mitra_priv_temp_id,
				m_mitra_privilege.mitra_priv_code as mitra_priv_temp_code,
				m_mitra_privilege.mitra_priv_name as mitra_priv_temp_name,
				m_mitra_privilege_category.mitra_priv_ctgry_id as mitra_priv_ctgry_temp_id,
				m_mitra_privilege_category.mitra_priv_ctgry_code as mitra_priv_ctgry_temp_code,
				m_mitra_privilege_category.mitra_priv_ctgry_name as mitra_priv_ctgry_temp_name`).
		Joins("JOIN m_mitra_privilege on m_mitra_group_access_temp_privilege.mitra_priv_temp_id=m_mitra_privilege.mitra_priv_id").
		Joins("JOIN m_mitra_privilege_category on m_mitra_privilege.mitra_priv_ctgry_id=m_mitra_privilege_category.mitra_priv_ctgry_id").
		Where("mitra_group_accs_temp_priv_id = ?", pid).
		Find(&mitragroup).Error
	if err != nil {
		return &MitraGroupAccessTempPrivilegeData{}, err
	}
	return &mitragroup, nil
}

func (p *MitraGroupAccessTempPrivilege) FindMitraGroupAccessTempPrivilegeDataByMitraGroupAccessTempID(db *gorm.DB, pid uint64) (*MitraGroupAccessTempPrivilege, error) {
	var err error
	err = db.Debug().Model(&MitraGroupAccessTempPrivilege{}).Where("mitra_group_accs_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraGroupAccessTempPrivilege{}, err
	}
	return p, nil
}

func (p *MitraGroupAccessTempPrivilege) FindMitraGroupAccessTempPrivilegeByMitraGroupAccessTempID(db *gorm.DB, pid uint64) ([]MitraGroupAccessTempPrivilegeData, error) {
	var err error
	mitragroup := []MitraGroupAccessTempPrivilegeData{}
	err = db.Debug().Model(&MitraGroupAccessTempPrivilegeData{}).Limit(100).
		Select(`m_mitra_group_access_temp_privilege.mitra_group_accs_temp_priv_id,
				m_mitra_group_access_temp_privilege.mitra_group_accs_temp_id,
				m_mitra_privilege.mitra_priv_id as mitra_priv_temp_id,
				m_mitra_privilege.mitra_priv_code as mitra_priv_temp_code,
				m_mitra_privilege.mitra_priv_name as mitra_priv_temp_name,
				m_mitra_privilege_category.mitra_priv_ctgry_id as mitra_priv_ctgry_temp_id,
				m_mitra_privilege_category.mitra_priv_ctgry_code as mitra_priv_ctgry_temp_code,
				m_mitra_privilege_category.mitra_priv_ctgry_name as mitra_priv_ctgry_temp_name`).
		Joins("JOIN m_mitra_privilege on m_mitra_group_access_temp_privilege.mitra_priv_temp_id=m_mitra_privilege.mitra_priv_id").
		Joins("JOIN m_mitra_privilege_category on m_mitra_privilege.mitra_priv_ctgry_id=m_mitra_privilege_category.mitra_priv_ctgry_id").
		Where("mitra_group_accs_temp_id = ?", pid).
		Find(&mitragroup).Error
	if err != nil {
		return []MitraGroupAccessTempPrivilegeData{}, err
	}
	return mitragroup, nil
}

func (p *MitraGroupAccessTempPrivilege) UpdateMitraGroupAccessTempPrivilege(db *gorm.DB, pid uint64) (*MitraGroupAccessTempPrivilege, error) {
	var err error
	db = db.Debug().Model(&MitraGroupAccessTempPrivilege{}).Where("mitra_group_accs_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_group_accs_temp_id": p.MitraGroupAccessTempID,
			"mitra_priv_temp_id":       p.MitraPrivilegeTempID,
		},
	)
	err = db.Debug().Model(&MitraGroupAccessTempPrivilege{}).Where("mitra_group_accs_temp_id = ?", pid).Error
	if err != nil {
		return &MitraGroupAccessTempPrivilege{}, err
	}
	return p, nil
}

func (p *MitraGroupAccessTempPrivilege) DeleteMitraGroupAccessTempPrivilege(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraGroupAccessTempPrivilege{}).Where("mitra_group_accs_temp_priv_id = ?", pid).Take(&MitraGroupAccessTempPrivilege{}).Delete(&MitraGroupAccessTempPrivilege{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraGroupAccessTempPrivilege not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
