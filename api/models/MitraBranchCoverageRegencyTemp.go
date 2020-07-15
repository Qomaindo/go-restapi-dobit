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

type MitraBranchCoverageRegencyTemp struct {
	MitraBranchCoverageRegencyTempID           *uint64   `gorm:"column:mitra_branch_cov_rgcy_temp_id;primary_key;" json:"mitra_branch_cov_rgcy_temp_id"`
	MitraBranchCoverageRegencyID               uint64    `gorm:"column:mitra_branch_cov_rgcy_id" json:"mitra_branch_cov_rgcy_id"`
	MitraBranchTempID                          uint64    `gorm:"column:mitra_branch_temp_id" json:"mitra_branch_temp_id"`
	RegencyTempID                              uint64    `gorm:"column:regency_temp_id" json:"regency_temp_id"`
	MitraBranchCoverageRegencyTempReason       string    `gorm:"column:mitra_branch_cov_rgcy_temp_reason" json:"mitra_branch_cov_rgcy_temp_reason"`
	MitraBranchCoverageRegencyTempStatus       uint64    `gorm:"column:mitra_branch_cov_rgcy_temp_status;size:2" json:"mitra_branch_cov_rgcy_temp_status"`
	MitraBranchCoverageRegencyTempActionBefore uint64    `gorm:"column:mitra_branch_cov_rgcy_temp_action_before;size:2" json:"mitra_branch_cov_rgcy_temp_action_before"`
	MitraBranchCoverageRegencyTempCreatedBy    string    `gorm:"column:mitra_branch_cov_rgcy_temp_created_by;size:125" json:"mitra_branch_cov_rgcy_temp_created_by"`
	MitraBranchCoverageRegencyTempCreatedAt    time.Time `gorm:"column:mitra_branch_cov_rgcy_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_branch_cov_rgcy_temp_created_at"`
}

type MitraBranchCoverageRegencyTempPend struct {
	MitraBranchCoverageRegencyTempID           *uint64   `gorm:"column:mitra_branch_cov_rgcy_temp_id;primary_key;" json:"mitra_branch_cov_rgcy_temp_id"`
	MitraBranchCoverageRegencyID               *uint64   `gorm:"column:mitra_branch_cov_rgcy_id" json:"mitra_branch_cov_rgcy_id"`
	MitraBranchTempID                          uint64    `gorm:"column:mitra_branch_temp_id" json:"mitra_branch_temp_id"`
	RegencyTempID                              uint64    `gorm:"column:regency_temp_id" json:"regency_temp_id"`
	MitraBranchCoverageRegencyTempReason       string    `gorm:"column:mitra_branch_cov_rgcy_temp_reason" json:"mitra_branch_cov_rgcy_temp_reason"`
	MitraBranchCoverageRegencyTempStatus       uint64    `gorm:"column:mitra_branch_cov_rgcy_temp_status;size:2" json:"mitra_branch_cov_rgcy_temp_status"`
	MitraBranchCoverageRegencyTempActionBefore uint64    `gorm:"column:mitra_branch_cov_rgcy_temp_action_before;size:2" json:"mitra_branch_cov_rgcy_temp_action_before"`
	MitraBranchCoverageRegencyTempCreatedBy    string    `gorm:"column:mitra_branch_cov_rgcy_temp_created_by;size:125" json:"mitra_branch_cov_rgcy_temp_created_by"`
	MitraBranchCoverageRegencyTempCreatedAt    time.Time `gorm:"column:mitra_branch_cov_rgcy_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_branch_cov_rgcy_temp_created_at"`
}

type MitraBranchCoverageRegencyTempPendData struct {
	MitraBranchCoverageRegencyTempID           uint64    `gorm:"column:mitra_branch_cov_rgcy_temp_id;primary_key;" json:"mitra_branch_cov_rgcy_temp_id"`
	MitraBranchCoverageRegencyID               uint64    `gorm:"column:mitra_branch_cov_rgcy_id" json:"mitra_branch_cov_rgcy_id"`
	MitraBranchTempID                          uint64    `gorm:"column:mitra_branch_temp_id" json:"mitra_branch_temp_id"`
	RegencyTempID                              uint64    `gorm:"column:regency_temp_id" json:"regency_temp_id"`
	RegencyTempName                            string    `gorm:"column:regency_temp_name" json:"regency_temp_name"`
	MitraBranchCoverageRegencyTempReason       string    `gorm:"column:mitra_branch_cov_rgcy_temp_reason" json:"mitra_branch_cov_rgcy_temp_reason"`
	MitraBranchCoverageRegencyTempStatus       uint64    `gorm:"column:mitra_branch_cov_rgcy_temp_status;size:2" json:"mitra_branch_cov_rgcy_temp_status"`
	MitraBranchCoverageRegencyTempActionBefore uint64    `gorm:"column:mitra_branch_cov_rgcy_temp_action_before;size:2" json:"mitra_branch_cov_rgcy_temp_action_before"`
	MitraBranchCoverageRegencyTempCreatedBy    string    `gorm:"column:mitra_branch_cov_rgcy_temp_created_by;size:125" json:"mitra_branch_cov_rgcy_temp_created_by"`
	MitraBranchCoverageRegencyTempCreatedAt    time.Time `gorm:"column:mitra_branch_cov_rgcy_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_branch_cov_rgcy_temp_created_at"`
}

type MitraBranchCoverageRegencyTempData struct {
	MitraBranchCoverageRegencyTempID           uint64    `gorm:"column:mitra_branch_cov_rgcy_temp_id" json:"mitra_branch_cov_rgcy_temp_id"`
	MitraBranchTempID                          uint64    `gorm:"column:mitra_branch_temp_id" json:"mitra_branch_temp_id"`
	MitraBranchTempName                        string    `gorm:"column:mitra_branch_temp_name" json:"mitra_branch_temp_name"`
	RegencyTempID                              uint64    `gorm:"column:regency_temp_id" json:"regency_temp_id"`
	RegencyTempName                            string    `gorm:"column:regency_temp_name" json:"regency_temp_name"`
	MitraBranchID                              uint64    `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	MitraBranchName                            string    `gorm:"column:mitra_branch_name" json:"mitra_branch_name"`
	RegencyID                                  uint64    `gorm:"column:regency_id" json:"regency_id"`
	RegencyName                                string    `gorm:"column:regency_name" json:"regency_name"`
	MitraBranchCoverageRegencyTempReason       string    `gorm:"column:mitra_branch_cov_rgcy_temp_reason" json:"mitra_branch_cov_rgcy_temp_reason"`
	MitraBranchCoverageRegencyTempStatus       uint64    `gorm:"column:mitra_branch_cov_rgcy_temp_status;size:2" json:"mitra_branch_cov_rgcy_temp_status"`
	MitraBranchCoverageRegencyTempActionBefore uint64    `gorm:"column:mitra_branch_cov_rgcy_temp_action_before;size:2" json:"mitra_branch_cov_rgcy_temp_action_before"`
	MitraBranchCoverageRegencyTempCreatedBy    string    `gorm:"column:mitra_branch_cov_rgcy_temp_created_by;size:125" json:"mitra_branch_cov_rgcy_temp_created_by"`
	MitraBranchCoverageRegencyTempCreatedAt    time.Time `gorm:"column:mitra_branch_cov_rgcy_temp_created_at" json:"mitra_branch_cov_rgcy_temp_created_at"`
	MitraBranchCoverageRegencyID               uint64    `gorm:"column:mitra_branch_cov_rgcy_id" json:"mitra_branch_cov_rgcy_id"`
}

type ResponseMitraBranchCoverageRegencyTemps struct {
	Status  int                                   `json:"status"`
	Message string                                `json:"message"`
	Data    *[]MitraBranchCoverageRegencyTempData `json:"data"`
}

type ResponseMitraBranchCoverageRegencyTemp struct {
	Status  int                                 `json:"status"`
	Message string                              `json:"message"`
	Data    *MitraBranchCoverageRegencyTempData `json:"data"`
}

type ResponseMitraBranchCoverageRegencyTempsPend struct {
	Status  int                                   `json:"status"`
	Message string                                `json:"message"`
	Data    *[]MitraBranchCoverageRegencyTempPend `json:"data"`
}

type ResponseMitraBranchCoverageRegencyTempPend struct {
	Status  int                                 `json:"status"`
	Message string                              `json:"message"`
	Data    *MitraBranchCoverageRegencyTempPend `json:"data"`
}

type ResponseMitraBranchCoverageRegencyTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraBranchCoverageRegencyTemp) TableName() string {
	return "m_mitra_branch_coverage_regency_temp"
}

func (MitraBranchCoverageRegencyTempPend) TableName() string {
	return "m_mitra_branch_coverage_regency_temp"
}

func (MitraBranchCoverageRegencyTempPendData) TableName() string {
	return "m_mitra_branch_coverage_regency_temp"
}

func (MitraBranchCoverageRegencyTempData) TableName() string {
	return "m_mitra_branch_coverage_regency_temp"
}

func (p *MitraBranchCoverageRegencyTemp) Prepare() {
	p.MitraBranchCoverageRegencyTempID = nil
	p.MitraBranchCoverageRegencyID = p.MitraBranchCoverageRegencyID
	p.MitraBranchTempID = p.MitraBranchTempID
	p.RegencyTempID = p.RegencyTempID
	p.MitraBranchCoverageRegencyTempReason = p.MitraBranchCoverageRegencyTempReason
	p.MitraBranchCoverageRegencyTempStatus = p.MitraBranchCoverageRegencyTempStatus
	p.MitraBranchCoverageRegencyTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraBranchCoverageRegencyTempCreatedBy))
	p.MitraBranchCoverageRegencyTempCreatedAt = time.Now()
}

func (p *MitraBranchCoverageRegencyTempPend) Prepare() {
	p.MitraBranchCoverageRegencyTempID = nil
	p.MitraBranchCoverageRegencyID = nil
	p.MitraBranchTempID = p.MitraBranchTempID
	p.RegencyTempID = p.RegencyTempID
	p.MitraBranchCoverageRegencyTempReason = p.MitraBranchCoverageRegencyTempReason
	p.MitraBranchCoverageRegencyTempStatus = p.MitraBranchCoverageRegencyTempStatus
	p.MitraBranchCoverageRegencyTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraBranchCoverageRegencyTempCreatedBy))
	p.MitraBranchCoverageRegencyTempCreatedAt = time.Now()
}

func (p *MitraBranchCoverageRegencyTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	case "insertupdate":
		if p.MitraBranchTempID == 0 {
			return errors.New("Negara Wajib diisi")
		}
		if p.RegencyTempID == 0 {
			return errors.New("Provinsi Wajib diisi")
		}
		if p.MitraBranchCoverageRegencyTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.MitraBranchCoverageRegencyTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	case "reason":
		if p.MitraBranchCoverageRegencyTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.MitraBranchCoverageRegencyTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	default:
		if p.MitraBranchCoverageRegencyTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil
	}
}

func (p *MitraBranchCoverageRegencyTemp) SaveMitraBranchCoverageRegencyTemp(db *gorm.DB) (*MitraBranchCoverageRegencyTemp, error) {
	var err error
	err = db.Debug().Model(&MitraBranchCoverageRegencyTemp{}).Create(&p).Error
	if err != nil {
		return &MitraBranchCoverageRegencyTemp{}, err
	}
	return p, nil
}

func (p *MitraBranchCoverageRegencyTempPend) SaveMitraBranchCoverageRegencyTempPend(db *gorm.DB) (*MitraBranchCoverageRegencyTempPend, error) {
	var err error
	err = db.Debug().Model(&MitraBranchCoverageRegencyTempPend{}).Create(&p).Error
	if err != nil {
		return &MitraBranchCoverageRegencyTempPend{}, err
	}
	return p, nil
}

func (p *MitraBranchCoverageRegencyTemp) FindAllMitraBranchCoverageRegencyTemp(db *gorm.DB) (*[]MitraBranchCoverageRegencyTempPend, error) {
	var err error
	mitrabranchcoverageregency := []MitraBranchCoverageRegencyTempPend{}
	err = db.Debug().Model(&MitraBranchCoverageRegencyTempPend{}).Limit(100).
		Select(`m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_id,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_id,
				m_mitra_branch_coverage_regency_temp.mitra_branch_temp_id,
				m_mitra_branch_coverage_regency_temp.regency_temp_id,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_reason,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_status,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_action_before,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_created_by,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_created_at`).
		Find(&mitrabranchcoverageregency).Error
	if err != nil {
		return &[]MitraBranchCoverageRegencyTempPend{}, err
	}
	return &mitrabranchcoverageregency, nil
}

func (p *MitraBranchCoverageRegencyTemp) FindAllMitraBranchCoverageRegencyTempPendingActive(db *gorm.DB) (*[]MitraBranchCoverageRegencyTempPend, error) {
	var err error
	mitrabranchcoverageregency := []MitraBranchCoverageRegencyTempPend{}
	err = db.Debug().Model(&MitraBranchCoverageRegencyTempPend{}).Limit(100).
		Select(`m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_id,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_id,
				m_mitra_branch_coverage_regency_temp.mitra_branch_temp_id,
				m_mitra_branch_coverage_regency_temp.regency_temp_id,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_reason,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_status,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_action_before,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_created_by,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_created_at`).
		Where("mitra_branch_cov_rgcy_temp_status = ?", 11).
		Find(&mitrabranchcoverageregency).Error
	if err != nil {
		return &[]MitraBranchCoverageRegencyTempPend{}, err
	}
	return &mitrabranchcoverageregency, nil
}

func (p *MitraBranchCoverageRegencyTemp) FindAllMitraBranchCoverageRegencyTempStatus(db *gorm.DB, status uint64) (*[]MitraBranchCoverageRegencyTempData, error) {
	var err error
	mitrabranchcoverageregency := []MitraBranchCoverageRegencyTempData{}
	err = db.Debug().Model(&MitraBranchCoverageRegencyTempData{}).Limit(100).
		Select(`m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_id,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_regency_temp.regency_id as regency_temp_id,
				m_regency_temp.regency_name as regency_temp_name,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_reason,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_status,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_action_before,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_created_by,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_created_at,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_id,
				m_mitra_branch.mitra_branch_id,
				m_mitra.mitra_branch_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_status,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_created_by,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_created_at,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_updated_by,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_updated_at,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_deleted_by,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_deleted_at`).
		Joins("JOIN m_mitra m_mitra_branch_temp on m_mitra_branch_coverage_regency_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_regency m_regency_temp on m_mitra_branch_coverage_regency_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_mitra_branch_coverage_regency on m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_id=m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_id").
		Joins("JOIN m_mitra on m_mitra_branch_coverage_regency.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_regency on m_mitra_branch_coverage_regency.regency_id=m_regency.regency_id").
		Where("mitra_branch_cov_rgcy_temp_status = ?", status).
		Find(&mitrabranchcoverageregency).Error
	if err != nil {
		return &[]MitraBranchCoverageRegencyTempData{}, err
	}
	return &mitrabranchcoverageregency, nil
}

func (p *MitraBranchCoverageRegencyTemp) FindMitraBranchCoverageRegencyTempDataByID(db *gorm.DB, pid uint64) (*MitraBranchCoverageRegencyTemp, error) {
	var err error
	err = db.Debug().Model(&MitraBranchCoverageRegencyTemp{}).Where("mitra_branch_cov_rgcy_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraBranchCoverageRegencyTemp{}, err
	}
	return p, nil
}

func (p *MitraBranchCoverageRegencyTemp) FindMitraBranchCoverageRegencyTempByIDPendingActive(db *gorm.DB, pid uint64) (*MitraBranchCoverageRegencyTempPend, error) {
	var err error
	mitrabranchcoverageregency := MitraBranchCoverageRegencyTempPend{}
	err = db.Debug().Model(&MitraBranchCoverageRegencyTempPend{}).Limit(100).
		Select(`m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_id,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_id,
				m_mitra_branch_coverage_regency_temp.mitra_branch_temp_id,
				m_mitra_branch_coverage_regency_temp.regency_temp_id,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_reason,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_status,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_action_before,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_created_by,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_created_at`).
		Where("mitra_branch_cov_rgcy_temp_id = ? AND mitra_branch_cov_rgcy_temp_status = ?", pid, 11).
		Find(&mitrabranchcoverageregency).Error
	if err != nil {
		return &MitraBranchCoverageRegencyTempPend{}, err
	}
	return &mitrabranchcoverageregency, nil
}

func (p *MitraBranchCoverageRegencyTemp) FindMitraBranchCoverageRegencyTempByID(db *gorm.DB, pid uint64) (*MitraBranchCoverageRegencyTempData, error) {
	var err error
	mitrabranchcoverageregency := MitraBranchCoverageRegencyTempData{}
	err = db.Debug().Model(&MitraBranchCoverageRegencyTempData{}).Limit(100).
		Select(`m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_id,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_regency_temp.regency_id as regency_temp_id,
				m_regency_temp.regency_name as regency_temp_name,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_reason,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_status,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_action_before,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_created_by,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_created_at,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_id,
				m_mitra_branch.mitra_branch_id,
				m_mitra.mitra_branch_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_status,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_created_by,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_created_at,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_updated_by,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_updated_at,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_deleted_by,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_deleted_at`).
		Joins("JOIN m_mitra m_mitra_branch_temp on m_mitra_branch_coverage_regency_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_regency m_regency_temp on m_mitra_branch_coverage_regency_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_mitra_branch_coverage_regency on m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_id=m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_id").
		Joins("JOIN m_mitra on m_mitra_branch_coverage_regency.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_regency on m_mitra_branch_coverage_regency.regency_id=m_regency.regency_id").
		Where("mitra_branch_cov_rgcy_temp_id = ?", pid).
		Find(&mitrabranchcoverageregency).Error
	if err != nil {
		return &MitraBranchCoverageRegencyTempData{}, err
	}
	return &mitrabranchcoverageregency, nil
}

func (p *MitraBranchCoverageRegencyTemp) FindMitraBranchCoverageRegencyTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraBranchCoverageRegencyTempData, error) {
	var err error
	mitrabranchcoverageregency := MitraBranchCoverageRegencyTempData{}
	err = db.Debug().Model(&MitraBranchCoverageRegencyTempData{}).Limit(100).
		Select(`m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_id,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_regency_temp.regency_id as regency_temp_id,
				m_regency_temp.regency_name as regency_temp_name,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_reason,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_status,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_action_before,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_created_by,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_created_at,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_id,
				m_mitra_branch.mitra_branch_id,
				m_mitra.mitra_branch_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_status,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_created_by,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_created_at,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_updated_by,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_updated_at,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_deleted_by,
				m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_deleted_at`).
		Joins("JOIN m_mitra m_mitra_branch_temp on m_mitra_branch_coverage_regency_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_regency m_regency_temp on m_mitra_branch_coverage_regency_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_mitra_branch_coverage_regency on m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_id=m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_id").
		Joins("JOIN m_mitra on m_mitra_branch_coverage_regency.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_regency on m_mitra_branch_coverage_regency.regency_id=m_regency.regency_id").
		Where("mitra_branch_cov_rgcy_temp_id = ? AND mitra_branch_cov_rgcy_temp_status = ?", pid, status).
		Find(&mitrabranchcoverageregency).Error
	if err != nil {
		return &MitraBranchCoverageRegencyTempData{}, err
	}
	return &mitrabranchcoverageregency, nil
}

func (p *MitraBranchCoverageRegencyTemp) UpdateMitraBranchCoverageRegencyTemp(db *gorm.DB, pid uint64) (*MitraBranchCoverageRegencyTemp, error) {
	var err error
	db = db.Debug().Model(&MitraBranchCoverageRegencyTemp{}).Where("mitra_branch_cov_rgcy_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_branch_temp_id":              p.MitraBranchTempID,
			"regency_temp_id":                   p.RegencyTempID,
			"mitra_branch_cov_rgcy_temp_reason": p.MitraBranchCoverageRegencyTempReason,
		},
	)
	err = db.Debug().Model(&MitraBranchCoverageRegencyTemp{}).Where("mitra_branch_cov_rgcy_temp_id = ?", pid).Error
	if err != nil {
		return &MitraBranchCoverageRegencyTemp{}, err
	}
	return p, nil
}

func (p *MitraBranchCoverageRegencyTemp) UpdateMitraBranchCoverageRegencyTempStatus(db *gorm.DB, pid uint64) (*MitraBranchCoverageRegencyTemp, error) {
	var err error
	db = db.Debug().Model(&MitraBranchCoverageRegencyTemp{}).Where("mitra_branch_cov_rgcy_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_branch_cov_rgcy_temp_reason": p.MitraBranchCoverageRegencyTempReason,
			"mitra_branch_cov_rgcy_temp_status": p.MitraBranchCoverageRegencyTempStatus,
		},
	)
	err = db.Debug().Model(&MitraBranchCoverageRegencyTemp{}).Where("mitra_branch_cov_rgcy_temp_id = ?", pid).Error
	if err != nil {
		return &MitraBranchCoverageRegencyTemp{}, err
	}
	return p, nil
}

func (p *MitraBranchCoverageRegencyTemp) DeleteMitraBranchCoverageRegencyTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraBranchCoverageRegencyTemp{}).Where("mitra_branch_cov_rgcy_temp_id = ?", pid).Take(&MitraBranchCoverageRegencyTemp{}).Delete(&MitraBranchCoverageRegencyTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraBranchCoverageRegencyTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

func (p *MitraBranchCoverageRegencyTemp) FindMitraBranchCoverageRegencyTempPendByMitraBranchTempID(db *gorm.DB, pid uint64) ([]MitraBranchCoverageRegencyTempPendData, error) {
	var err error
	mitrabranchcoverageregency := []MitraBranchCoverageRegencyTempPendData{}
	err = db.Debug().Model(&MitraBranchCoverageRegencyTempPendData{}).Limit(100).
		Select(`m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_id,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_id,
				m_mitra_branch_coverage_regency_temp.mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_branch_coverage_regency_temp.regency_temp_id,
				m_regency_temp.regency_name as regency_temp_name,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_reason,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_status,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_action_before,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_created_by,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_created_at`).
		Joins("JOIN m_regency m_regency_temp on m_mitra_branch_coverage_regency_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_mitra_branch_temp on m_mitra_branch_coverage_regency_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_temp_id").
		Where("m_mitra_branch_coverage_regency_temp.mitra_branch_temp_id = ?", pid).
		Find(&mitrabranchcoverageregency).Error
	if err != nil {
		return []MitraBranchCoverageRegencyTempPendData{}, err
	}
	return mitrabranchcoverageregency, nil
}

func (p *MitraBranchCoverageRegencyTemp) FindMitraBranchCoverageRegencyTempPendStatusByMitraBranchTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraBranchCoverageRegencyTempPendData, error) {
	var err error
	mitrabranchcoverageregency := []MitraBranchCoverageRegencyTempPendData{}
	err = db.Debug().Model(&MitraBranchCoverageRegencyTempPendData{}).Limit(100).
		Select(`m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_id,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_id,
				m_mitra_branch_coverage_regency_temp.mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_branch_coverage_regency_temp.regency_temp_id,
				m_regency_temp.regency_name as regency_temp_name,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_reason,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_status,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_action_before,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_created_by,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_created_at`).
		Joins("JOIN m_regency m_regency_temp on m_mitra_branch_coverage_regency_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_mitra_branch_temp on m_mitra_branch_coverage_regency_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_temp_id").
		Where("m_mitra_branch_coverage_regency_temp.mitra_branch_temp_id = ? AND mitra_branch_cov_rgcy_temp_status = ?", pid, status).
		Find(&mitrabranchcoverageregency).Error
	if err != nil {
		return []MitraBranchCoverageRegencyTempPendData{}, err
	}
	return mitrabranchcoverageregency, nil
}

func (p *MitraBranchCoverageRegencyTemp) FindMitraBranchCoverageRegencyTempByMitraBranchTempID(db *gorm.DB, pid uint64) ([]MitraBranchCoverageRegencyTempData, error) {
	var err error
	mitrabranchcoverageregency := []MitraBranchCoverageRegencyTempData{}
	err = db.Debug().Model(&MitraBranchCoverageRegencyTempData{}).Limit(100).
		Select(`m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_id,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_id,
				m_mitra_branch_coverage_regency_temp.mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_branch_coverage_regency_temp.regency_temp_id,
				m_regency_temp.regency_name as regency_temp_name,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_reason,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_status,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_action_before,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_created_by,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_created_at`).
		Joins("JOIN m_regency m_regency_temp on m_mitra_branch_coverage_regency_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_mitra_branch_temp on m_mitra_branch_coverage_regency_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_temp_id").
		Where("m_mitra_branch_coverage_regency_temp.mitra_branch_temp_id = ?", pid).
		Find(&mitrabranchcoverageregency).Error
	if err != nil {
		return []MitraBranchCoverageRegencyTempData{}, err
	}
	return mitrabranchcoverageregency, nil
}

func (p *MitraBranchCoverageRegencyTemp) FindMitraBranchCoverageRegencyTempStatusByMitraBranchTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraBranchCoverageRegencyTempData, error) {
	var err error
	mitrabranchcoverageregency := []MitraBranchCoverageRegencyTempData{}
	err = db.Debug().Model(&MitraBranchCoverageRegencyTempData{}).Limit(100).
		Select(`m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_id,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_id,
				m_mitra_branch_coverage_regency_temp.mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_branch_coverage_regency_temp.regency_temp_id,
				m_regency_temp.regency_name as regency_temp_name,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_reason,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_status,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_action_before,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_created_by,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_created_at`).
		Joins("JOIN m_regency m_regency_temp on m_mitra_branch_coverage_regency_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_mitra_branch_temp on m_mitra_branch_coverage_regency_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_temp_id").
		Where("m_mitra_branch_coverage_regency_temp.mitra_branch_temp_id = ? AND mitra_branch_cov_rgcy_temp_status = ?", pid, status).
		Find(&mitrabranchcoverageregency).Error
	if err != nil {
		return []MitraBranchCoverageRegencyTempData{}, err
	}
	return mitrabranchcoverageregency, nil
}

func (p *MitraBranchCoverageRegencyTemp) FindMitraBranchCoverageRegencyTempDataStatusByMitraBranchTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraBranchCoverageRegencyTempData, error) {
	var err error
	mitrabranchcoverageregency := []MitraBranchCoverageRegencyTempData{}
	err = db.Debug().Model(&MitraBranchCoverageRegencyTempData{}).Limit(100).
		Select(`m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_id,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_id,
				m_mitra_branch_temp.mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_branch_coverage_regency_temp.regency_temp_id,
				m_regency_temp.regency_name as regency_temp_name,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_reason,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_status,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_action_before,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_created_by,
				m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_temp_created_at`).
		Joins("JOIN m_regency m_regency_temp on m_mitra_branch_coverage_regency_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_mitra_branch_temp on m_mitra_branch_coverage_regency_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_temp_id").
		Joins("JOIN m_mitra_branch_coverage_regency on m_mitra_branch_coverage_regency_temp.mitra_branch_cov_rgcy_id=m_mitra_branch_coverage_regency.mitra_branch_cov_rgcy_id").
		Joins("JOIN m_regency on m_mitra_branch_coverage_regency.regency_id=m_regency.regency_id").
		Joins("JOIN m_mitra_branch on m_mitra_branch_coverage_regency.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Where("m_mitra_branch_coverage_regency_temp.mitra_branch_temp_id = ? AND mitra_branch_cov_rgcy_temp_status = ?", pid, status).
		Find(&mitrabranchcoverageregency).Error
	if err != nil {
		return []MitraBranchCoverageRegencyTempData{}, err
	}
	return mitrabranchcoverageregency, nil
}
