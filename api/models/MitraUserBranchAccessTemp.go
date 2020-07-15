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

type MitraUserBranchAccessTemp struct {
	MitraUserBranchAccessTempID           *uint64   `gorm:"column:mitra_user_branch_accs_temp_id;primary_key;" json:"mitra_user_branch_accs_temp_id"`
	MitraUserBranchAccessID               uint64    `gorm:"column:mitra_user_branch_accs_id" json:"mitra_user_branch_accs_id"`
	MitraUserTempID                       uint64    `gorm:"column:mitra_user_temp_id" json:"mitra_user_temp_id"`
	MitraBranchTempID                     uint64    `gorm:"column:mitra_branch_temp_id" json:"mitra_branch_temp_id"`
	MitraUserBranchAccessTempReason       string    `gorm:"column:mitra_user_branch_accs_temp_reason" json:"mitra_user_branch_accs_temp_reason"`
	MitraUserBranchAccessTempStatus       uint64    `gorm:"column:mitra_user_branch_accs_temp_status;size:2" json:"mitra_user_branch_accs_temp_status"`
	MitraUserBranchAccessTempActionBefore uint64    `gorm:"column:mitra_user_branch_accs_temp_action_before;size:2" json:"mitra_user_branch_accs_temp_action_before"`
	MitraUserBranchAccessTempCreatedBy    string    `gorm:"column:mitra_user_branch_accs_temp_created_by;size:125" json:"mitra_user_branch_accs_temp_created_by"`
	MitraUserBranchAccessTempCreatedAt    time.Time `gorm:"column:mitra_user_branch_accs_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_user_branch_accs_temp_created_at"`
}

type MitraUserBranchAccessTempPend struct {
	MitraUserBranchAccessTempID           *uint64   `gorm:"column:mitra_user_branch_accs_temp_id;primary_key;" json:"mitra_user_branch_accs_temp_id"`
	MitraUserBranchAccessID               *uint64   `gorm:"column:mitra_user_branch_accs_id" json:"mitra_user_branch_accs_id"`
	MitraUserTempID                       uint64    `gorm:"column:mitra_user_temp_id" json:"mitra_user_temp_id"`
	MitraBranchTempID                     uint64    `gorm:"column:mitra_branch_temp_id" json:"mitra_branch_temp_id"`
	MitraUserBranchAccessTempReason       string    `gorm:"column:mitra_user_branch_accs_temp_reason" json:"mitra_user_branch_accs_temp_reason"`
	MitraUserBranchAccessTempStatus       uint64    `gorm:"column:mitra_user_branch_accs_temp_status;size:2" json:"mitra_user_branch_accs_temp_status"`
	MitraUserBranchAccessTempActionBefore uint64    `gorm:"column:mitra_user_branch_accs_temp_action_before;size:2" json:"mitra_user_branch_accs_temp_action_before"`
	MitraUserBranchAccessTempCreatedBy    string    `gorm:"column:mitra_user_branch_accs_temp_created_by;size:125" json:"mitra_user_branch_accs_temp_created_by"`
	MitraUserBranchAccessTempCreatedAt    time.Time `gorm:"column:mitra_user_branch_accs_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_user_branch_accs_temp_created_at"`
}

type MitraUserBranchAccessTempPendData struct {
	MitraUserBranchAccessTempID           uint64    `gorm:"column:mitra_user_branch_accs_temp_id;primary_key;" json:"mitra_user_branch_accs_temp_id"`
	MitraUserBranchAccessID               uint64    `gorm:"column:mitra_user_branch_accs_id" json:"mitra_user_branch_accs_id"`
	MitraUserTempID                       uint64    `gorm:"column:mitra_user_temp_id" json:"mitra_user_temp_id"`
	MitraBranchTempID                     uint64    `gorm:"column:mitra_branch_temp_id" json:"mitra_branch_temp_id"`
	MitraBranchTempName                   string    `gorm:"column:mitra_branch_temp_name" json:"mitra_branch_temp_name"`
	MitraUserBranchAccessTempReason       string    `gorm:"column:mitra_user_branch_accs_temp_reason" json:"mitra_user_branch_accs_temp_reason"`
	MitraUserBranchAccessTempStatus       uint64    `gorm:"column:mitra_user_branch_accs_temp_status;size:2" json:"mitra_user_branch_accs_temp_status"`
	MitraUserBranchAccessTempActionBefore uint64    `gorm:"column:mitra_user_branch_accs_temp_action_before;size:2" json:"mitra_user_branch_accs_temp_action_before"`
	MitraUserBranchAccessTempCreatedBy    string    `gorm:"column:mitra_user_branch_accs_temp_created_by;size:125" json:"mitra_user_branch_accs_temp_created_by"`
	MitraUserBranchAccessTempCreatedAt    time.Time `gorm:"column:mitra_user_branch_accs_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_user_branch_accs_temp_created_at"`
}

type MitraUserBranchAccessTempData struct {
	MitraUserBranchAccessTempID           uint64    `gorm:"column:mitra_user_branch_accs_temp_id" json:"mitra_user_branch_accs_temp_id"`
	MitraUserBranchAccessID               uint64    `gorm:"column:mitra_user_branch_accs_id" json:"mitra_user_branch_accs_id"`
	MitraUserTempID                       uint64    `gorm:"column:mitra_user_temp_id" json:"mitra_user_temp_id"`
	MitraUserTempName                     string    `gorm:"column:mitra_user_temp_username" json:"mitra_user_temp_username"`
	MitraBranchTempID                     uint64    `gorm:"column:mitra_branch_temp_id" json:"mitra_branch_temp_id"`
	MitraBranchTempName                   string    `gorm:"column:mitra_branch_temp_name" json:"mitra_branch_temp_name"`
	MitraUserID                           uint64    `gorm:"column:mitra_user_id" json:"mitra_user_id"`
	MitraUserName                         string    `gorm:"column:mitra_username" json:"mitra_username"`
	MitraBranchID                         uint64    `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	MitraBranchName                       string    `gorm:"column:mitra_branch_name" json:"mitra_branch_name"`
	MitraUserBranchAccessTempReason       string    `gorm:"column:mitra_user_branch_accs_temp_reason" json:"mitra_user_branch_accs_temp_reason"`
	MitraUserBranchAccessTempStatus       uint64    `gorm:"column:mitra_user_branch_accs_temp_status;size:2" json:"mitra_user_branch_accs_temp_status"`
	MitraUserBranchAccessTempActionBefore uint64    `gorm:"column:mitra_user_branch_accs_temp_action_before;size:2" json:"mitra_user_branch_accs_temp_action_before"`
	MitraUserBranchAccessTempCreatedBy    string    `gorm:"column:mitra_user_branch_accs_temp_created_by;size:125" json:"mitra_user_branch_accs_temp_created_by"`
	MitraUserBranchAccessTempCreatedAt    time.Time `gorm:"column:mitra_user_branch_accs_temp_created_at" json:"mitra_user_branch_accs_temp_created_at"`
}

type ResponseMitraUserBranchAccessTemps struct {
	Status  int                              `json:"status"`
	Message string                           `json:"message"`
	Data    *[]MitraUserBranchAccessTempData `json:"data"`
}

type ResponseMitraUserBranchAccessTemp struct {
	Status  int                            `json:"status"`
	Message string                         `json:"message"`
	Data    *MitraUserBranchAccessTempData `json:"data"`
}

type ResponseMitraUserBranchAccessTempsPend struct {
	Status  int                              `json:"status"`
	Message string                           `json:"message"`
	Data    *[]MitraUserBranchAccessTempPend `json:"data"`
}

type ResponseMitraUserBranchAccessTempPend struct {
	Status  int                            `json:"status"`
	Message string                         `json:"message"`
	Data    *MitraUserBranchAccessTempPend `json:"data"`
}

type ResponseMitraUserBranchAccessTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraUserBranchAccessTemp) TableName() string {
	return "m_mitra_user_branch_access_temp"
}

func (MitraUserBranchAccessTempPend) TableName() string {
	return "m_mitra_user_branch_access_temp"
}

func (MitraUserBranchAccessTempPendData) TableName() string {
	return "m_mitra_user_branch_access_temp"
}

func (MitraUserBranchAccessTempData) TableName() string {
	return "m_mitra_user_branch_access_temp"
}

func (p *MitraUserBranchAccessTemp) Prepare() {
	p.MitraUserBranchAccessTempID = nil
	p.MitraUserBranchAccessID = p.MitraUserBranchAccessID
	p.MitraUserTempID = p.MitraUserTempID
	p.MitraBranchTempID = p.MitraBranchTempID
	p.MitraUserBranchAccessTempReason = p.MitraUserBranchAccessTempReason
	p.MitraUserBranchAccessTempStatus = p.MitraUserBranchAccessTempStatus
	p.MitraUserBranchAccessTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraUserBranchAccessTempCreatedBy))
	p.MitraUserBranchAccessTempCreatedAt = time.Now()
}

func (p *MitraUserBranchAccessTempPend) Prepare() {
	p.MitraUserBranchAccessTempID = nil
	p.MitraUserBranchAccessID = nil
	p.MitraUserTempID = p.MitraUserTempID
	p.MitraBranchTempID = p.MitraBranchTempID
	p.MitraUserBranchAccessTempReason = p.MitraUserBranchAccessTempReason
	p.MitraUserBranchAccessTempStatus = p.MitraUserBranchAccessTempStatus
	p.MitraUserBranchAccessTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraUserBranchAccessTempCreatedBy))
	p.MitraUserBranchAccessTempCreatedAt = time.Now()
}

func (p *MitraUserBranchAccessTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	case "insertupdate":
		if p.MitraUserTempID == 0 {
			return errors.New("Negara Wajib diisi")
		}
		if p.MitraBranchTempID == 0 {
			return errors.New("Provinsi Wajib diisi")
		}
		if p.MitraUserBranchAccessTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.MitraUserBranchAccessTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	case "reason":
		if p.MitraUserBranchAccessTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.MitraUserBranchAccessTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	default:
		if p.MitraUserBranchAccessTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil
	}
}

func (p *MitraUserBranchAccessTemp) SaveMitraUserBranchAccessTemp(db *gorm.DB) (*MitraUserBranchAccessTemp, error) {
	var err error
	err = db.Debug().Model(&MitraUserBranchAccessTemp{}).Create(&p).Error
	if err != nil {
		return &MitraUserBranchAccessTemp{}, err
	}
	return p, nil
}

func (p *MitraUserBranchAccessTempPend) SaveMitraUserBranchAccessTempPend(db *gorm.DB) (*MitraUserBranchAccessTempPend, error) {
	var err error
	err = db.Debug().Model(&MitraUserBranchAccessTempPend{}).Create(&p).Error
	if err != nil {
		return &MitraUserBranchAccessTempPend{}, err
	}
	return p, nil
}

func (p *MitraUserBranchAccessTemp) FindAllMitraUserBranchAccessTemp(db *gorm.DB) (*[]MitraUserBranchAccessTempPend, error) {
	var err error
	mitrauserbranchaccess := []MitraUserBranchAccessTempPend{}
	err = db.Debug().Model(&MitraUserBranchAccessTempPend{}).Limit(100).
		Select(`m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_id,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_id,
				m_mitra_user_branch_access_temp.mitra_user_temp_id,
				m_mitra_user_branch_access_temp.mitra_branch_temp_id,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_reason,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_status,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_action_before,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_created_by,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_created_at`).
		Find(&mitrauserbranchaccess).Error
	if err != nil {
		return &[]MitraUserBranchAccessTempPend{}, err
	}
	return &mitrauserbranchaccess, nil
}

func (p *MitraUserBranchAccessTemp) FindAllMitraUserBranchAccessTempPendingActive(db *gorm.DB) (*[]MitraUserBranchAccessTempPend, error) {
	var err error
	mitrauserbranchaccess := []MitraUserBranchAccessTempPend{}
	err = db.Debug().Model(&MitraUserBranchAccessTempPend{}).Limit(100).
		Select(`m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_id,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_id,
				m_mitra_user_branch_access_temp.mitra_user_temp_id,
				m_mitra_user_branch_access_temp.mitra_branch_temp_id,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_reason,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_status,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_action_before,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_created_by,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_created_at`).
		Where("mitra_user_branch_accs_temp_status = ?", 11).
		Find(&mitrauserbranchaccess).Error
	if err != nil {
		return &[]MitraUserBranchAccessTempPend{}, err
	}
	return &mitrauserbranchaccess, nil
}

func (p *MitraUserBranchAccessTemp) FindAllMitraUserBranchAccessTempStatus(db *gorm.DB, status uint64) (*[]MitraUserBranchAccessTempData, error) {
	var err error
	mitrauserbranchaccess := []MitraUserBranchAccessTempData{}
	err = db.Debug().Model(&MitraUserBranchAccessTempData{}).Limit(100).
		Select(`m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_id,
				m_mitra_user_temp.mitra_user_id as mitra_user_temp_id,
				m_mitra_user_temp.mitra_username as mitra_user_temp_username,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_reason,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_status,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_action_before,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_created_by,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_created_at,
				m_mitra_user_branch_access.mitra_user_branch_accs_id,
				m_mitra_user.mitra_user_id,
				m_mitra_user.mitra_username,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_user_branch_access.mitra_user_branch_accs_status,
				m_mitra_user_branch_access.mitra_user_branch_accs_created_by,
				m_mitra_user_branch_access.mitra_user_branch_accs_created_at,
				m_mitra_user_branch_access.mitra_user_branch_accs_updated_by,
				m_mitra_user_branch_access.mitra_user_branch_accs_updated_at,
				m_mitra_user_branch_access.mitra_user_branch_accs_deleted_by,
				m_mitra_user_branch_access.mitra_user_branch_accs_deleted_at`).
		Joins("JOIN m_mitra_user m_mitra_user_temp on m_mitra_user_branch_access_temp.mitra_user_temp_id=m_mitra_user_temp.mitra_user_temp_id").
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_user_branch_access_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_mitra_user_branch_access on m_mitra_user_branch_access_temp.mitra_user_branch_accs_id=m_mitra_user_branch_access.mitra_user_branch_accs_id").
		Joins("JOIN m_mitra_user on m_mitra_user_branch_access.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra_branch on m_mitra_user_branch_access.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Where("mitra_user_branch_accs_temp_status = ?", status).
		Find(&mitrauserbranchaccess).Error
	if err != nil {
		return &[]MitraUserBranchAccessTempData{}, err
	}
	return &mitrauserbranchaccess, nil
}

func (p *MitraUserBranchAccessTemp) FindMitraUserBranchAccessTempDataByID(db *gorm.DB, pid uint64) (*MitraUserBranchAccessTemp, error) {
	var err error
	err = db.Debug().Model(&MitraUserBranchAccessTemp{}).Where("mitra_user_branch_accs_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraUserBranchAccessTemp{}, err
	}
	return p, nil
}

func (p *MitraUserBranchAccessTemp) FindMitraUserBranchAccessTempByIDPendingActive(db *gorm.DB, pid uint64) (*MitraUserBranchAccessTempPend, error) {
	var err error
	mitrauserbranchaccess := MitraUserBranchAccessTempPend{}
	err = db.Debug().Model(&MitraUserBranchAccessTempPend{}).Limit(100).
		Select(`m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_id,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_id,
				m_mitra_user_branch_access_temp.mitra_user_temp_id,
				m_mitra_user_branch_access_temp.mitra_branch_temp_id,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_reason,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_status,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_action_before,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_created_by,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_created_at`).
		Where("mitra_user_branch_accs_temp_id = ? AND mitra_user_branch_accs_temp_status = ?", pid, 11).
		Find(&mitrauserbranchaccess).Error
	if err != nil {
		return &MitraUserBranchAccessTempPend{}, err
	}
	return &mitrauserbranchaccess, nil
}

func (p *MitraUserBranchAccessTemp) FindMitraUserBranchAccessTempByID(db *gorm.DB, pid uint64) (*MitraUserBranchAccessTempData, error) {
	var err error
	mitrauserbranchaccess := MitraUserBranchAccessTempData{}
	err = db.Debug().Model(&MitraUserBranchAccessTempData{}).Limit(100).
		Select(`m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_id,
				m_mitra_user_temp.mitra_user_id as mitra_user_temp_id,
				m_mitra_user_temp.mitra_username as mitra_user_temp_username,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_reason,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_status,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_action_before,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_created_by,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_created_at,
				m_mitra_user_branch_access.mitra_user_branch_accs_id,
				m_mitra_user.mitra_user_id,
				m_mitra_user.mitra_username,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_user_branch_access.mitra_user_branch_accs_status,
				m_mitra_user_branch_access.mitra_user_branch_accs_created_by,
				m_mitra_user_branch_access.mitra_user_branch_accs_created_at,
				m_mitra_user_branch_access.mitra_user_branch_accs_updated_by,
				m_mitra_user_branch_access.mitra_user_branch_accs_updated_at,
				m_mitra_user_branch_access.mitra_user_branch_accs_deleted_by,
				m_mitra_user_branch_access.mitra_user_branch_accs_deleted_at`).
		Joins("JOIN m_mitra_user m_mitra_user_temp on m_mitra_user_branch_access_temp.mitra_user_temp_id=m_mitra_user_temp.mitra_user_temp_id").
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_user_branch_access_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_mitra_user_branch_access on m_mitra_user_branch_access_temp.mitra_user_branch_accs_id=m_mitra_user_branch_access.mitra_user_branch_accs_id").
		Joins("JOIN m_mitra_user on m_mitra_user_branch_access.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra_branch on m_mitra_user_branch_access.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Where("mitra_user_branch_accs_temp_id = ?", pid).
		Find(&mitrauserbranchaccess).Error
	if err != nil {
		return &MitraUserBranchAccessTempData{}, err
	}
	return &mitrauserbranchaccess, nil
}

func (p *MitraUserBranchAccessTemp) FindMitraUserBranchAccessTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraUserBranchAccessTempData, error) {
	var err error
	mitrauserbranchaccess := MitraUserBranchAccessTempData{}
	err = db.Debug().Model(&MitraUserBranchAccessTempData{}).Limit(100).
		Select(`m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_id,
				m_mitra_user_temp.mitra_user_id as mitra_user_temp_id,
				m_mitra_user_temp.mitra_username as mitra_user_temp_username,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_reason,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_status,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_action_before,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_created_by,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_created_at,
				m_mitra_user_branch_access.mitra_user_branch_accs_id,
				m_mitra_user.mitra_user_id,
				m_mitra_user.mitra_username,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_user_branch_access.mitra_user_branch_accs_status,
				m_mitra_user_branch_access.mitra_user_branch_accs_created_by,
				m_mitra_user_branch_access.mitra_user_branch_accs_created_at,
				m_mitra_user_branch_access.mitra_user_branch_accs_updated_by,
				m_mitra_user_branch_access.mitra_user_branch_accs_updated_at,
				m_mitra_user_branch_access.mitra_user_branch_accs_deleted_by,
				m_mitra_user_branch_access.mitra_user_branch_accs_deleted_at`).
		Joins("JOIN m_mitra_user m_mitra_user_temp on m_mitra_user_branch_access_temp.mitra_user_temp_id=m_mitra_user_temp.mitra_user_temp_id").
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_user_branch_access_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_mitra_user_branch_access on m_mitra_user_branch_access_temp.mitra_user_branch_accs_id=m_mitra_user_branch_access.mitra_user_branch_accs_id").
		Joins("JOIN m_mitra_user on m_mitra_user_branch_access.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra_branch on m_mitra_user_branch_access.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Where("mitra_user_branch_accs_temp_id = ? AND mitra_user_branch_accs_temp_status = ?", pid, status).
		Find(&mitrauserbranchaccess).Error
	if err != nil {
		return &MitraUserBranchAccessTempData{}, err
	}
	return &mitrauserbranchaccess, nil
}

func (p *MitraUserBranchAccessTemp) UpdateMitraUserBranchAccessTemp(db *gorm.DB, pid uint64) (*MitraUserBranchAccessTemp, error) {
	var err error
	db = db.Debug().Model(&MitraUserBranchAccessTemp{}).Where("mitra_user_branch_accs_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_user_id":                      p.MitraUserTempID,
			"mitra_branch_temp_id":               p.MitraBranchTempID,
			"mitra_user_branch_accs_temp_reason": p.MitraUserBranchAccessTempReason,
		},
	)
	err = db.Debug().Model(&MitraUserBranchAccessTemp{}).Where("mitra_user_branch_accs_temp_id = ?", pid).Error
	if err != nil {
		return &MitraUserBranchAccessTemp{}, err
	}
	return p, nil
}

func (p *MitraUserBranchAccessTemp) UpdateMitraUserBranchAccessTempStatus(db *gorm.DB, pid uint64) (*MitraUserBranchAccessTemp, error) {
	var err error
	db = db.Debug().Model(&MitraUserBranchAccessTemp{}).Where("mitra_user_branch_accs_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_user_branch_accs_temp_reason": p.MitraUserBranchAccessTempReason,
			"mitra_user_branch_accs_temp_status": p.MitraUserBranchAccessTempStatus,
		},
	)
	err = db.Debug().Model(&MitraUserBranchAccessTemp{}).Where("mitra_user_branch_accs_temp_id = ?", pid).Error
	if err != nil {
		return &MitraUserBranchAccessTemp{}, err
	}
	return p, nil
}

func (p *MitraUserBranchAccessTemp) DeleteMitraUserBranchAccessTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraUserBranchAccessTemp{}).Where("mitra_user_branch_accs_temp_id = ?", pid).Take(&MitraUserBranchAccessTemp{}).Delete(&MitraUserBranchAccessTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraUserBranchAccessTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

func (p *MitraUserBranchAccessTemp) FindMitraUserBranchAccessTempPendByMitraUserTempID(db *gorm.DB, pid uint64) ([]MitraUserBranchAccessTempPendData, error) {
	var err error
	mitrauserbranchaccess := []MitraUserBranchAccessTempPendData{}
	err = db.Debug().Model(&MitraUserBranchAccessTempPendData{}).Limit(100).
		Select(`m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_id,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_id,
				m_mitra_user_branch_access_temp.mitra_user_temp_id,
				m_mitra_user_temp.mitra_user_temp_username,
				m_mitra_user_branch_access_temp.mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_reason,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_status,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_action_before,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_created_by,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_created_at`).
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_user_branch_access_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_mitra_user_temp on m_mitra_user_branch_access_temp.mitra_user_temp_id=m_mitra_user_temp.mitra_user_temp_id").
		Where("m_mitra_user_branch_access_temp.mitra_user_temp_id = ?", pid).
		Find(&mitrauserbranchaccess).Error
	if err != nil {
		return []MitraUserBranchAccessTempPendData{}, err
	}
	return mitrauserbranchaccess, nil
}

func (p *MitraUserBranchAccessTemp) FindMitraUserBranchAccessTempPendStatusByMitraUserTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraUserBranchAccessTempPendData, error) {
	var err error
	mitrauserbranchaccess := []MitraUserBranchAccessTempPendData{}
	err = db.Debug().Model(&MitraUserBranchAccessTempPendData{}).Limit(100).
		Select(`m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_id,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_id,
				m_mitra_user_branch_access_temp.mitra_user_temp_id,
				m_mitra_user_temp.mitra_user_temp_username,
				m_mitra_user_branch_access_temp.mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_reason,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_status,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_action_before,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_created_by,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_created_at`).
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_user_branch_access_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_mitra_user_temp on m_mitra_user_branch_access_temp.mitra_user_temp_id=m_mitra_user_temp.mitra_user_temp_id").
		Where("m_mitra_user_branch_access_temp.mitra_user_temp_id = ? AND mitra_user_branch_accs_temp_status = ?", pid, status).
		Find(&mitrauserbranchaccess).Error
	if err != nil {
		return []MitraUserBranchAccessTempPendData{}, err
	}
	return mitrauserbranchaccess, nil
}

func (p *MitraUserBranchAccessTemp) FindMitraUserBranchAccessTempByMitraUserTempID(db *gorm.DB, pid uint64) ([]MitraUserBranchAccessTempData, error) {
	var err error
	mitrauserbranchaccess := []MitraUserBranchAccessTempData{}
	err = db.Debug().Model(&MitraUserBranchAccessTempData{}).Limit(100).
		Select(`m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_id,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_id,
				m_mitra_user_branch_access_temp.mitra_user_temp_id,
				m_mitra_user_temp.mitra_user_temp_username,
				m_mitra_user_branch_access_temp.mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_reason,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_status,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_action_before,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_created_by,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_created_at`).
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_user_branch_access_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_mitra_user_temp on m_mitra_user_branch_access_temp.mitra_user_temp_id=m_mitra_user_temp.mitra_user_temp_id").
		Where("m_mitra_user_branch_access_temp.mitra_user_temp_id = ?", pid).
		Find(&mitrauserbranchaccess).Error
	if err != nil {
		return []MitraUserBranchAccessTempData{}, err
	}
	return mitrauserbranchaccess, nil
}

func (p *MitraUserBranchAccessTemp) FindMitraUserBranchAccessTempStatusByMitraUserTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraUserBranchAccessTempData, error) {
	var err error
	mitrauserbranchaccess := []MitraUserBranchAccessTempData{}
	err = db.Debug().Model(&MitraUserBranchAccessTempData{}).Limit(100).
		Select(`m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_id,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_id,
				m_mitra_user_branch_access_temp.mitra_user_temp_id,
				m_mitra_user_temp.mitra_user_temp_username,
				m_mitra_user_branch_access_temp.mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_reason,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_status,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_action_before,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_created_by,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_created_at`).
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_user_branch_access_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_mitra_user_temp on m_mitra_user_branch_access_temp.mitra_user_temp_id=m_mitra_user_temp.mitra_user_temp_id").
		Where("m_mitra_user_branch_access_temp.mitra_user_temp_id = ? AND mitra_user_branch_accs_temp_status = ?", pid, status).
		Find(&mitrauserbranchaccess).Error
	if err != nil {
		return []MitraUserBranchAccessTempData{}, err
	}
	return mitrauserbranchaccess, nil
}

func (p *MitraUserBranchAccessTemp) FindMitraUserBranchAccessTempDataStatusByMitraUserTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraUserBranchAccessTempData, error) {
	var err error
	mitrauserbranchaccess := []MitraUserBranchAccessTempData{}
	err = db.Debug().Model(&MitraUserBranchAccessTempData{}).Limit(100).
		Select(`m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_id,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_id,
				m_mitra_user_temp.mitra_user_id,
				m_mitra_user_temp.mitra_user_temp_username,
				m_mitra_user_branch_access_temp.mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_mitra_user.mitra_user_id,
				m_mitra_user.mitra_username,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_reason,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_status,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_action_before,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_created_by,
				m_mitra_user_branch_access_temp.mitra_user_branch_accs_temp_created_at`).
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_user_branch_access_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_mitra_user_temp on m_mitra_user_branch_access_temp.mitra_user_temp_id=m_mitra_user_temp.mitra_user_temp_id").
		Joins("JOIN m_mitra_user_branch_access on m_mitra_user_branch_access_temp.mitra_user_branch_accs_id=m_mitra_user_branch_access.mitra_user_branch_accs_id").
		Joins("JOIN m_mitra_branch on m_mitra_user_branch_access.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra_user on m_mitra_user_branch_access.mitra_user_id=m_mitra_user.mitra_user_id").
		Where("m_mitra_user_branch_access_temp.mitra_user_temp_id = ? AND mitra_user_branch_accs_temp_status = ?", pid, status).
		Find(&mitrauserbranchaccess).Error
	if err != nil {
		return []MitraUserBranchAccessTempData{}, err
	}
	return mitrauserbranchaccess, nil
}
