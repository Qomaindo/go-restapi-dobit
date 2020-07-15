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

type MitraCompanyTemp struct {
	MitraCompanyTempID           uint64    `gorm:"column:mitra_company_temp_id;primary_key;" json:"mitra_company_temp_id"`
	MitraCompanyID               uint64    `gorm:"column:mitra_company_id" json:"mitra_company_id"`
	MitraTempID                  uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	CompanyTempID                uint64    `gorm:"column:company_temp_id" json:"company_temp_id"`
	MitraCompanyTempReason       string    `gorm:"column:mitra_company_temp_reason" json:"mitra_company_temp_reason"`
	MitraCompanyTempStatus       uint64    `gorm:"column:mitra_company_temp_status;size:2" json:"mitra_company_temp_status"`
	MitraCompanyTempActionBefore uint64    `gorm:"column:mitra_company_temp_action_before;size:2" json:"mitra_company_temp_action_before"`
	MitraCompanyTempCreatedBy    string    `gorm:"column:mitra_company_temp_created_by;size:125" json:"mitra_company_temp_created_by"`
	MitraCompanyTempCreatedAt    time.Time `gorm:"column:mitra_company_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_company_temp_created_at"`
}

type MitraCompanyTempPend struct {
	MitraCompanyTempID           uint64    `gorm:"column:mitra_company_temp_id;primary_key;" json:"mitra_company_temp_id"`
	MitraCompanyID               *uint64   `gorm:"column:mitra_company_id" json:"mitra_company_id"`
	MitraTempID                  uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	CompanyTempID                uint64    `gorm:"column:company_temp_id" json:"company_temp_id"`
	MitraCompanyTempReason       string    `gorm:"column:mitra_company_temp_reason" json:"mitra_company_temp_reason"`
	MitraCompanyTempStatus       uint64    `gorm:"column:mitra_company_temp_status;size:2" json:"mitra_company_temp_status"`
	MitraCompanyTempActionBefore uint64    `gorm:"column:mitra_company_temp_action_before;size:2" json:"mitra_company_temp_action_before"`
	MitraCompanyTempCreatedBy    string    `gorm:"column:mitra_company_temp_created_by;size:125" json:"mitra_company_temp_created_by"`
	MitraCompanyTempCreatedAt    time.Time `gorm:"column:mitra_company_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_company_temp_created_at"`
}

type MitraCompanyTempPendData struct {
	MitraCompanyTempID           uint64    `gorm:"column:mitra_company_temp_id;primary_key;" json:"mitra_company_temp_id"`
	MitraCompanyID               *uint64   `gorm:"column:mitra_company_id" json:"mitra_company_id"`
	MitraTempID                  uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	MitraTempName                string    `gorm:"column:mitra_temp_name" json:"mitra_temp_name"`
	CompanyTempID                uint64    `gorm:"column:company_temp_id" json:"company_temp_id"`
	CompanyTempName              string    `gorm:"column:company_temp_name" json:"company_temp_name"`
	MitraCompanyTempReason       string    `gorm:"column:mitra_company_temp_reason" json:"mitra_company_temp_reason"`
	MitraCompanyTempStatus       uint64    `gorm:"column:mitra_company_temp_status;size:2" json:"mitra_company_temp_status"`
	MitraCompanyTempActionBefore uint64    `gorm:"column:mitra_company_temp_action_before;size:2" json:"mitra_company_temp_action_before"`
	MitraCompanyTempCreatedBy    string    `gorm:"column:mitra_company_temp_created_by;size:125" json:"mitra_company_temp_created_by"`
	MitraCompanyTempCreatedAt    time.Time `gorm:"column:mitra_company_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_company_temp_created_at"`
}

type MitraCompanyTempData struct {
	MitraCompanyTempID           uint64    `gorm:"column:mitra_company_temp_id" json:"mitra_company_temp_id"`
	MitraTempID                  uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	MitraTempName                string    `gorm:"column:mitra_temp_name" json:"mitra_temp_name"`
	CompanyTempID                uint64    `gorm:"column:company_temp_id" json:"company_temp_id"`
	CompanyTempName              string    `gorm:"column:company_temp_name" json:"company_temp_name"`
	MitraCompanyTempReason       string    `gorm:"column:mitra_company_temp_reason" json:"mitra_company_temp_reason"`
	MitraCompanyTempStatus       uint64    `gorm:"column:mitra_company_temp_status;size:2" json:"mitra_company_temp_status"`
	MitraCompanyTempActionBefore uint64    `gorm:"column:mitra_company_temp_action_before;size:2" json:"mitra_company_temp_action_before"`
	MitraCompanyTempCreatedBy    string    `gorm:"column:mitra_company_temp_created_by;size:125" json:"mitra_company_temp_created_by"`
	MitraCompanyTempCreatedAt    time.Time `gorm:"column:mitra_company_temp_created_at" json:"mitra_company_temp_created_at"`
	MitraCompanyID               uint64    `gorm:"column:mitra_company_id" json:"mitra_company_id"`
	MitraID                      uint64    `gorm:"column:mitra_id" json:"mitra_id"`
	MitraName                    string    `gorm:"column:mitra_name" json:"mitra_name"`
	CompanyID                    uint64    `gorm:"column:company_id" json:"company_id"`
	CompanyName                  string    `gorm:"column:company_name" json:"company_name"`
	MitraCompanyStatus           uint64    `gorm:"column:mitra_company_status" json:"mitra_company_status"`
	MitraCompanyCreatedBy        string    `gorm:"column:mitra_company_created_by" json:"mitra_company_created_by"`
	MitraCompanyCreatedAt        time.Time `gorm:"column:mitra_company_created_at" json:"mitra_company_created_at"`
	MitraCompanyUpdatedBy        string    `gorm:"column:mitra_company_updated_by" json:"mitra_company_updated_by"`
	MitraCompanyUpdatedAt        time.Time `gorm:"column:mitra_company_updated_at" json:"mitra_company_updated_at"`
	MitraCompanyDeletedBy        string    `gorm:"column:mitra_company_deleted_by" json:"mitra_company_deleted_by"`
	MitraCompanyDeletedAt        time.Time `gorm:"column:mitra_company_deleted_at" json:"mitra_company_deleted_at"`
}

type ResponseMitraCompanyTemps struct {
	Status  int                     `json:"status"`
	Message string                  `json:"message"`
	Data    *[]MitraCompanyTempData `json:"data"`
}

type ResponseMitraCompanyTemp struct {
	Status  int                   `json:"status"`
	Message string                `json:"message"`
	Data    *MitraCompanyTempData `json:"data"`
}

type ResponseMitraCompanyTempsPend struct {
	Status  int                         `json:"status"`
	Message string                      `json:"message"`
	Data    *[]MitraCompanyTempPendData `json:"data"`
}

type ResponseMitraCompanyTempPend struct {
	Status  int                       `json:"status"`
	Message string                    `json:"message"`
	Data    *MitraCompanyTempPendData `json:"data"`
}

type ResponseMitraCompanyTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraCompanyTemp) TableName() string {
	return "m_mitra_company_temp"
}

func (MitraCompanyTempPend) TableName() string {
	return "m_mitra_company_temp"
}

func (MitraCompanyTempPendData) TableName() string {
	return "m_mitra_company_temp"
}

func (MitraCompanyTempData) TableName() string {
	return "m_mitra_company_temp"
}

func (p *MitraCompanyTemp) Prepare() {
	p.MitraCompanyID = p.MitraCompanyID
	p.MitraTempID = p.MitraTempID
	p.CompanyTempID = p.CompanyTempID
	p.MitraCompanyTempReason = p.MitraCompanyTempReason
	p.MitraCompanyTempStatus = p.MitraCompanyTempStatus
	p.MitraCompanyTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraCompanyTempCreatedBy))
	p.MitraCompanyTempCreatedAt = time.Now()
}

func (p *MitraCompanyTempPend) Prepare() {
	p.MitraCompanyID = nil
	p.MitraTempID = p.MitraTempID
	p.CompanyTempID = p.CompanyTempID
	p.MitraCompanyTempReason = p.MitraCompanyTempReason
	p.MitraCompanyTempStatus = p.MitraCompanyTempStatus
	p.MitraCompanyTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraCompanyTempCreatedBy))
	p.MitraCompanyTempCreatedAt = time.Now()
}

func (p *MitraCompanyTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	case "insertupdate":
		if p.MitraTempID == 0 {
			return errors.New("Negara Wajib diisi")
		}
		if p.CompanyTempID == 0 {
			return errors.New("Provinsi Wajib diisi")
		}
		if p.MitraCompanyTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.MitraCompanyTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	case "reason":
		if p.MitraCompanyTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.MitraCompanyTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	default:
		if p.MitraCompanyTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil
	}
}

func (p *MitraCompanyTemp) SaveMitraCompanyTemp(db *gorm.DB) (*MitraCompanyTemp, error) {
	var err error
	err = db.Debug().Model(&MitraCompanyTemp{}).Create(&p).Error
	if err != nil {
		return &MitraCompanyTemp{}, err
	}
	return p, nil
}

func (p *MitraCompanyTempPend) SaveMitraCompanyTempPend(db *gorm.DB) (*MitraCompanyTempPend, error) {
	var err error
	err = db.Debug().Model(&MitraCompanyTempPend{}).Create(&p).Error
	if err != nil {
		return &MitraCompanyTempPend{}, err
	}
	return p, nil
}

func (p *MitraCompanyTemp) FindAllMitraCompanyTemp(db *gorm.DB) (*[]MitraCompanyTempPendData, error) {
	var err error
	mitra_company := []MitraCompanyTempPendData{}
	err = db.Debug().Model(&MitraCompanyTempPendData{}).Limit(100).
		Select(`m_mitra_company_temp.mitra_company_temp_id,
				m_mitra_company_temp.mitra_company_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_company_temp.company_id as company_temp_id,
				m_company_temp.company_name as company_temp_name,
				m_mitra_company_temp.mitra_company_temp_reason,
				m_mitra_company_temp.mitra_company_temp_status,
				m_mitra_company_temp.mitra_company_temp_action_before,
				m_mitra_company_temp.mitra_company_temp_created_by,
				m_mitra_company_temp.mitra_company_temp_created_at at time zone 'Asia/Jakarta' as mitra_company_temp_created_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_company_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_company m_company_temp on m_mitra_company_temp.company_temp_id=m_company_temp.company_id").
		Order("mitra_company_temp_created_at desc").
		Find(&mitra_company).Error
	if err != nil {
		return &[]MitraCompanyTempPendData{}, err
	}
	return &mitra_company, nil
}

func (p *MitraCompanyTemp) FindAllMitraCompanyTempStatusPendingActive(db *gorm.DB) (*[]MitraCompanyTempPendData, error) {
	var err error
	mitra_company := []MitraCompanyTempPendData{}
	err = db.Debug().Model(&MitraCompanyTempPendData{}).Limit(100).
		Select(`m_mitra_company_temp.mitra_company_temp_id,
				m_mitra_company_temp.mitra_company_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_company_temp.company_id as company_temp_id,
				m_company_temp.company_name as company_temp_name,
				m_mitra_company_temp.mitra_company_temp_reason,
				m_mitra_company_temp.mitra_company_temp_status,
				m_mitra_company_temp.mitra_company_temp_action_before,
				m_mitra_company_temp.mitra_company_temp_created_by,
				m_mitra_company_temp.mitra_company_temp_created_at at time zone 'Asia/Jakarta' as mitra_company_temp_created_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_company_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_company m_company_temp on m_mitra_company_temp.company_temp_id=m_company_temp.company_id").
		Where("mitra_company_temp_status = ?", 11).
		Order("mitra_company_temp_created_at desc").
		Find(&mitra_company).Error
	if err != nil {
		return &[]MitraCompanyTempPendData{}, err
	}
	return &mitra_company, nil
}

func (p *MitraCompanyTemp) FindAllMitraCompanyTempStatus(db *gorm.DB, status uint64) (*[]MitraCompanyTempData, error) {
	var err error
	mitra_company := []MitraCompanyTempData{}
	err = db.Debug().Model(&MitraCompanyTempData{}).Limit(100).
		Select(`m_mitra_company_temp.mitra_company_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_company_temp.company_id as company_temp_id,
				m_company_temp.company_name as company_temp_name,
				m_mitra_company_temp.mitra_company_temp_reason,
				m_mitra_company_temp.mitra_company_temp_status,
				m_mitra_company_temp.mitra_company_temp_action_before,
				m_mitra_company_temp.mitra_company_temp_created_by,
				m_mitra_company_temp.mitra_company_temp_created_at at time zone 'Asia/Jakarta' as mitra_company_temp_created_at,
				m_mitra_company.mitra_company_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_company.company_id,
				m_company.company_name,
				m_mitra_company.mitra_company_status,
				m_mitra_company.mitra_company_created_by,
				m_mitra_company.mitra_company_created_at at time zone 'Asia/Jakarta' as mitra_company_created_at,
				m_mitra_company.mitra_company_updated_by,
				m_mitra_company.mitra_company_updated_at at time zone 'Asia/Jakarta' as mitra_company_updated_at,
				m_mitra_company.mitra_company_deleted_by,
				m_mitra_company.mitra_company_deleted_at at time zone 'Asia/Jakarta' as mitra_company_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_company_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_company m_company_temp on m_mitra_company_temp.company_temp_id=m_company_temp.company_id").
		Joins("JOIN m_mitra_company on m_mitra_company_temp.mitra_company_id=m_mitra_company.mitra_company_id").
		Joins("JOIN m_mitra on m_mitra_company.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_company on m_mitra_company.company_id=m_company.company_id").
		Where("mitra_company_temp_status = ?", status).
		Order("mitra_company_temp_created_at desc").
		Find(&mitra_company).Error
	if err != nil {
		return &[]MitraCompanyTempData{}, err
	}
	return &mitra_company, nil
}

func (p *MitraCompanyTemp) FindMitraCompanyTempDataByID(db *gorm.DB, pid uint64) (*MitraCompanyTemp, error) {
	var err error
	err = db.Debug().Model(&MitraCompanyTemp{}).Where("mitra_company_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraCompanyTemp{}, err
	}
	return p, nil
}

func (p *MitraCompanyTemp) FindMitraCompanyTempByIDPendingActive(db *gorm.DB, pid uint64) (*MitraCompanyTempPendData, error) {
	var err error
	mitra_company := MitraCompanyTempPendData{}
	err = db.Debug().Model(&MitraCompanyTempPendData{}).Limit(100).
		Select(`m_mitra_company_temp.mitra_company_temp_id,
				m_mitra_company_temp.mitra_company_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_company_temp.company_id as company_temp_id,
				m_company_temp.company_name as company_temp_name,
				m_mitra_company_temp.mitra_company_temp_reason,
				m_mitra_company_temp.mitra_company_temp_status,
				m_mitra_company_temp.mitra_company_temp_action_before,
				m_mitra_company_temp.mitra_company_temp_created_by,
				m_mitra_company_temp.mitra_company_temp_created_at at time zone 'Asia/Jakarta' as mitra_company_temp_created_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_company_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_company m_company_temp on m_mitra_company_temp.company_temp_id=m_company_temp.company_id").
		Where("mitra_company_temp_id = ?", pid).
		Order("mitra_company_temp_created_at desc").
		Find(&mitra_company).Error
	if err != nil {
		return &MitraCompanyTempPendData{}, err
	}
	return &mitra_company, nil
}

func (p *MitraCompanyTemp) FindMitraCompanyTempStatusByIDPendingActive(db *gorm.DB, pid uint64) (*MitraCompanyTempPendData, error) {
	var err error
	mitra_company := MitraCompanyTempPendData{}
	err = db.Debug().Model(&MitraCompanyTempPendData{}).Limit(100).
		Select(`m_mitra_company_temp.mitra_company_temp_id,
				m_mitra_company_temp.mitra_company_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_company_temp.company_id as company_temp_id,
				m_company_temp.company_name as company_temp_name,
				m_mitra_company_temp.mitra_company_temp_reason,
				m_mitra_company_temp.mitra_company_temp_status,
				m_mitra_company_temp.mitra_company_temp_action_before,
				m_mitra_company_temp.mitra_company_temp_created_by,
				m_mitra_company_temp.mitra_company_temp_created_at at time zone 'Asia/Jakarta' as mitra_company_temp_created_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_company_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_company m_company_temp on m_mitra_company_temp.company_temp_id=m_company_temp.company_id").
		Where("mitra_company_temp_id = ? AND mitra_company_temp_status = ?", pid, 11).
		Order("mitra_company_temp_created_at desc").
		Find(&mitra_company).Error
	if err != nil {
		return &MitraCompanyTempPendData{}, err
	}
	return &mitra_company, nil
}

func (p *MitraCompanyTemp) FindMitraCompanyTempByID(db *gorm.DB, pid uint64) (*MitraCompanyTempData, error) {
	var err error
	mitra_company := MitraCompanyTempData{}
	err = db.Debug().Model(&MitraCompanyTempData{}).Limit(100).
		Select(`m_mitra_company_temp.mitra_company_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_company_temp.company_id as company_temp_id,
				m_company_temp.company_name as company_temp_name,
				m_mitra_company_temp.mitra_company_temp_reason,
				m_mitra_company_temp.mitra_company_temp_status,
				m_mitra_company_temp.mitra_company_temp_action_before,
				m_mitra_company_temp.mitra_company_temp_created_by,
				m_mitra_company_temp.mitra_company_temp_created_at at time zone 'Asia/Jakarta' as mitra_company_temp_created_at,
				m_mitra_company.mitra_company_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_company.company_id,
				m_company.company_name,
				m_mitra_company.mitra_company_status,
				m_mitra_company.mitra_company_created_by,
				m_mitra_company.mitra_company_created_at at time zone 'Asia/Jakarta' as mitra_company_created_at,
				m_mitra_company.mitra_company_updated_by,
				m_mitra_company.mitra_company_updated_at at time zone 'Asia/Jakarta' as mitra_company_updated_at,
				m_mitra_company.mitra_company_deleted_by,
				m_mitra_company.mitra_company_deleted_at at time zone 'Asia/Jakarta' as mitra_company_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_company_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_company m_company_temp on m_mitra_company_temp.company_temp_id=m_company_temp.company_id").
		Joins("JOIN m_mitra_company on m_mitra_company_temp.mitra_company_id=m_mitra_company.mitra_company_id").
		Joins("JOIN m_mitra on m_mitra_company.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_company on m_mitra_company.company_id=m_company.company_id").
		Where("mitra_company_temp_id = ?", pid).
		Order("mitra_company_temp_created_at desc").
		Find(&mitra_company).Error
	if err != nil {
		return &MitraCompanyTempData{}, err
	}
	return &mitra_company, nil
}

func (p *MitraCompanyTemp) FindMitraCompanyTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraCompanyTempData, error) {
	var err error
	mitra_company := MitraCompanyTempData{}
	err = db.Debug().Model(&MitraCompanyTempData{}).Limit(100).
		Select(`m_mitra_company_temp.mitra_company_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_company_temp.company_id as company_temp_id,
				m_company_temp.company_name as company_temp_name,
				m_mitra_company_temp.mitra_company_temp_reason,
				m_mitra_company_temp.mitra_company_temp_status,
				m_mitra_company_temp.mitra_company_temp_action_before,
				m_mitra_company_temp.mitra_company_temp_created_by,
				m_mitra_company_temp.mitra_company_temp_created_at at time zone 'Asia/Jakarta' as mitra_company_temp_created_at,
				m_mitra_company.mitra_company_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_company.company_id,
				m_company.company_name,
				m_mitra_company.mitra_company_status,
				m_mitra_company.mitra_company_created_by,
				m_mitra_company.mitra_company_created_at at time zone 'Asia/Jakarta' as mitra_company_created_at,
				m_mitra_company.mitra_company_updated_by,
				m_mitra_company.mitra_company_updated_at at time zone 'Asia/Jakarta' as mitra_company_updated_at,
				m_mitra_company.mitra_company_deleted_by,
				m_mitra_company.mitra_company_deleted_at at time zone 'Asia/Jakarta' as mitra_company_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_company_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_company m_company_temp on m_mitra_company_temp.company_temp_id=m_company_temp.company_id").
		Joins("JOIN m_mitra_company on m_mitra_company_temp.mitra_company_id=m_mitra_company.mitra_company_id").
		Joins("JOIN m_mitra on m_mitra_company.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_company on m_mitra_company.company_id=m_company.company_id").
		Where("mitra_company_temp_id = ? AND mitra_company_temp_status = ?", pid, status).
		Order("mitra_company_temp_created_at desc").
		Find(&mitra_company).Error
	if err != nil {
		return &MitraCompanyTempData{}, err
	}
	return &mitra_company, nil
}

func (p *MitraCompanyTemp) UpdateMitraCompanyTemp(db *gorm.DB, pid uint64) (*MitraCompanyTemp, error) {
	var err error
	db = db.Debug().Model(&MitraCompanyTemp{}).Where("mitra_company_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_temp_id":             p.MitraTempID,
			"company_temp_id":           p.CompanyTempID,
			"mitra_company_temp_reason": p.MitraCompanyTempReason,
		},
	)
	err = db.Debug().Model(&MitraCompanyTemp{}).Where("mitra_company_temp_id = ?", pid).Error
	if err != nil {
		return &MitraCompanyTemp{}, err
	}
	return p, nil
}

func (p *MitraCompanyTemp) UpdateMitraCompanyTempStatus(db *gorm.DB, pid uint64) (*MitraCompanyTemp, error) {
	var err error
	db = db.Debug().Model(&MitraCompanyTemp{}).Where("mitra_company_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_company_temp_reason": p.MitraCompanyTempReason,
			"mitra_company_temp_status": p.MitraCompanyTempStatus,
		},
	)
	err = db.Debug().Model(&MitraCompanyTemp{}).Where("mitra_company_temp_id = ?", pid).Error
	if err != nil {
		return &MitraCompanyTemp{}, err
	}
	return p, nil
}

func (p *MitraCompanyTemp) DeleteMitraCompanyTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraCompanyTemp{}).Where("mitra_company_temp_id = ?", pid).Take(&MitraCompanyTemp{}).Delete(&MitraCompanyTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraCompanyTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

func (p *MitraCompany) FindAllMitraCompanyReferalCode(db *gorm.DB, referal string) (*MitraCompany, error) {
	var err error
	err = db.Debug().Model(&MitraCompany{}).Where("mitra_company_referal_code = ? AND mitra_status = ?", referal, 1).Take(&p).Error
	if err != nil {
		return &MitraCompany{}, err
	}
	return p, nil
}
