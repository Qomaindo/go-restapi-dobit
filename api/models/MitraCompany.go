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

type MitraCompany struct {
	MitraCompanyID          uint64     `gorm:"column:mitra_company_id;primary_key;" json:"mitra_company_id"`
	MitraID                 uint64     `gorm:"column:mitra_id" json:"mitra_id"`
	CompanyID               uint64     `gorm:"column:company_id" json:"company_id"`
	MitraCompanyReferalCode string     `gorm:"column:mitra_company_referal_code" json:"mitra_company_referal_code"`
	MitraCompanyStatus      uint64     `gorm:"column:mitra_company_status;size:2" json:"mitra_company_status"`
	MitraCompanyCreatedBy   string     `gorm:"column:mitra_company_created_by;size:125" json:"mitra_company_created_by"`
	MitraCompanyCreatedAt   time.Time  `gorm:"column:mitra_company_created_at;default:CURRENT_TIMESTAMP" json:"mitra_company_created_at"`
	MitraCompanyUpdatedBy   string     `gorm:"column:mitra_company_updated_by;size:125" json:"mitra_company_updated_by"`
	MitraCompanyUpdatedAt   *time.Time `gorm:"column:mitra_company_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_company_created_at"`
	MitraCompanyDeletedBy   string     `gorm:"column:mitra_company_deleted_by;size:125" json:"mitra_company_deleted_by"`
	MitraCompanyDeletedAt   *time.Time `gorm:"column:mitra_company_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_company_deleted_at"`
}

type MitraCompanyData struct {
	MitraCompanyID          uint64     `gorm:"column:mitra_company_id" json:"mitra_company_id"`
	MitraID                 uint64     `gorm:"column:mitra_id" json:"mitra_id"`
	MitraName               string     `gorm:"column:mitra_name" json:"mitra_name"`
	CompanyID               uint64     `gorm:"column:company_id" json:"company_id"`
	CompanyName             string     `gorm:"column:company_name" json:"company_name"`
	MitraCompanyReferalCode string     `gorm:"column:mitra_company_referal_code" json:"mitra_company_referal_code"`
	MitraCompanyStatus      uint64     `gorm:"column:mitra_company_status" json:"mitra_company_status"`
	MitraCompanyCreatedBy   string     `gorm:"column:mitra_company_created_by" json:"mitra_company_created_by"`
	MitraCompanyCreatedAt   time.Time  `gorm:"column:mitra_company_created_at" json:"mitra_company_created_at"`
	MitraCompanyUpdatedBy   string     `gorm:"column:mitra_company_updated_by" json:"mitra_company_updated_by"`
	MitraCompanyUpdatedAt   *time.Time `gorm:"column:mitra_company_updated_at" json:"mitra_company_updated_at"`
	MitraCompanyDeletedBy   string     `gorm:"column:mitra_company_deleted_by" json:"mitra_company_deleted_by"`
	MitraCompanyDeletedAt   *time.Time `gorm:"column:mitra_company_deleted_at" json:"mitra_company_deleted_at"`
}

type ResponseMitraCompanys struct {
	Status  int                 `json:"status"`
	Message string              `json:"message"`
	Data    *[]MitraCompanyData `json:"data"`
}

type ResponseMitraCompany struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    *MitraCompanyData `json:"data"`
}

type ResponseMitraCompanyCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraCompany) TableName() string {
	return "m_mitra_company"
}

func (MitraCompanyData) TableName() string {
	return "m_mitra_company"
}

func (p *MitraCompany) Prepare() {
	p.MitraID = p.MitraID
	p.CompanyID = p.CompanyID
	p.MitraCompanyReferalCode = p.MitraCompanyReferalCode
	p.MitraCompanyStatus = p.MitraCompanyStatus
	p.MitraCompanyCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraCompanyCreatedBy))
	p.MitraCompanyCreatedAt = time.Now()
	p.MitraCompanyUpdatedBy = html.EscapeString(strings.TrimSpace(p.MitraCompanyUpdatedBy))
	p.MitraCompanyUpdatedAt = p.MitraCompanyUpdatedAt
	p.MitraCompanyDeletedBy = html.EscapeString(strings.TrimSpace(p.MitraCompanyDeletedBy))
	p.MitraCompanyDeletedAt = p.MitraCompanyDeletedAt
}

func (p *MitraCompany) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.MitraID == 0 {
			return errors.New("Required Country")
		}
		if p.CompanyID == 0 {
			return errors.New("Required Province")
		}
		if p.MitraCompanyReferalCode == "0" {
			return errors.New("Required Regency")
		}
		return nil
	}
}

func (p *MitraCompany) SaveMitraCompany(db *gorm.DB) (*MitraCompany, error) {
	var err error
	err = db.Debug().Model(&MitraCompany{}).Create(&p).Error
	if err != nil {
		return &MitraCompany{}, err
	}
	return p, nil
}

func (p *MitraCompany) FindAllMitraCompany(db *gorm.DB) (*[]MitraCompanyData, error) {
	var err error
	mitra_company := []MitraCompanyData{}
	err = db.Debug().Model(&MitraCompanyData{}).Limit(100).
		Select(`m_mitra_company.mitra_company_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_company.company_id,
				m_company.company_name,
				m_mitra_company.mitra_company_referal_code,
				m_mitra_company.mitra_company_status,
				m_mitra_company.mitra_company_created_by,
				m_mitra_company.mitra_company_created_at at time zone 'Asia/Jakarta' as mitra_company_created_at,
				m_mitra_company.mitra_company_updated_by,
				m_mitra_company.mitra_company_updated_at at time zone 'Asia/Jakarta' as mitra_company_updated_at,
				m_mitra_company.mitra_company_deleted_by,
				m_mitra_company.mitra_company_deleted_at at time zone 'Asia/Jakarta' as mitra_company_deleted_at`).
		Joins("JOIN m_mitra on m_mitra_company.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_company on m_mitra_company.company_id=m_company.company_id").
		Find(&mitra_company).Error
	if err != nil {
		return &[]MitraCompanyData{}, err
	}
	return &mitra_company, nil
}

func (p *MitraCompany) FindAllMitraCompanyStatus(db *gorm.DB, status uint64) (*[]MitraCompanyData, error) {
	var err error
	mitra_company := []MitraCompanyData{}
	err = db.Debug().Model(&MitraCompanyData{}).Limit(100).
		Select(`m_mitra_company.mitra_company_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_company.company_id,
				m_company.company_name,
				m_mitra_company.mitra_company_referal_code,
				m_mitra_company.mitra_company_status,
				m_mitra_company.mitra_company_created_by,
				m_mitra_company.mitra_company_created_at at time zone 'Asia/Jakarta' as mitra_company_created_at,
				m_mitra_company.mitra_company_updated_by,
				m_mitra_company.mitra_company_updated_at at time zone 'Asia/Jakarta' as mitra_company_updated_at,
				m_mitra_company.mitra_company_deleted_by,
				m_mitra_company.mitra_company_deleted_at at time zone 'Asia/Jakarta' as mitra_company_deleted_at`).
		Joins("JOIN m_mitra on m_mitra_company.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_company on m_mitra_company.company_id=m_company.company_id").
		Where("mitra_company_status = ?", status).
		Find(&mitra_company).Error
	if err != nil {
		return &[]MitraCompanyData{}, err
	}
	return &mitra_company, nil
}

func (p *MitraCompany) FindMitraCompanyDataByID(db *gorm.DB, pid uint64) (*MitraCompany, error) {
	var err error
	err = db.Debug().Model(&MitraCompany{}).Where("mitra_company_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraCompany{}, err
	}
	return p, nil
}

func (p *MitraCompany) FindMitraCompanyByID(db *gorm.DB, pid uint64) (*MitraCompanyData, error) {
	var err error
	mitra_company := MitraCompanyData{}
	err = db.Debug().Model(&MitraCompanyData{}).Limit(100).
		Select(`m_mitra_company.mitra_company_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_company.company_id,
				m_company.company_name,
				m_mitra_company.mitra_company_referal_code,
				m_mitra_company.mitra_company_status,
				m_mitra_company.mitra_company_created_by,
				m_mitra_company.mitra_company_created_at at time zone 'Asia/Jakarta' as mitra_company_created_at,
				m_mitra_company.mitra_company_updated_by,
				m_mitra_company.mitra_company_updated_at at time zone 'Asia/Jakarta' as mitra_company_updated_at,
				m_mitra_company.mitra_company_deleted_by,
				m_mitra_company.mitra_company_deleted_at at time zone 'Asia/Jakarta' as mitra_company_deleted_at`).
		Joins("JOIN m_mitra on m_mitra_company.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_company on m_mitra_company.company_id=m_company.company_id").
		Where("mitra_company_id = ?", pid).
		Find(&mitra_company).Error
	if err != nil {
		return &MitraCompanyData{}, err
	}
	return &mitra_company, nil
}

func (p *MitraCompany) FindMitraCompanyStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraCompanyData, error) {
	var err error
	mitra_company := MitraCompanyData{}
	err = db.Debug().Model(&MitraCompanyData{}).Limit(100).
		Select(`m_mitra_company.mitra_company_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_company.company_id,
				m_company.company_name,
				m_mitra_company.mitra_company_referal_code,
				m_mitra_company.mitra_company_status,
				m_mitra_company.mitra_company_created_by,
				m_mitra_company.mitra_company_created_at at time zone 'Asia/Jakarta' as mitra_company_created_at,
				m_mitra_company.mitra_company_updated_by,
				m_mitra_company.mitra_company_updated_at at time zone 'Asia/Jakarta' as mitra_company_updated_at,
				m_mitra_company.mitra_company_deleted_by,
				m_mitra_company.mitra_company_deleted_at at time zone 'Asia/Jakarta' as mitra_company_deleted_at`).
		Joins("JOIN m_mitra on m_mitra_company.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_company on m_mitra_company.company_id=m_company.company_id").
		Where("mitra_company_id = ? AND mitra_company_status = ?", pid, status).
		Find(&mitra_company).Error
	if err != nil {
		return &MitraCompanyData{}, err
	}
	return &mitra_company, nil
}

func (p *MitraCompany) UpdateMitraCompany(db *gorm.DB, pid uint64) (*MitraCompany, error) {
	var err error
	db = db.Debug().Model(&MitraCompany{}).Where("mitra_company_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_id":   p.MitraID,
			"company_id": p.CompanyID,
			// "mitra_company_referal_code": p.MitraCompanyReferalCode,
			"mitra_company_status":     p.MitraCompanyStatus,
			"mitra_company_updated_by": p.MitraCompanyUpdatedBy,
			"mitra_company_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraCompany{}).Where("mitra_company_id = ?", pid).Error
	if err != nil {
		return &MitraCompany{}, err
	}
	return p, nil
}

func (p *MitraCompany) UpdateMitraCompanyStatus(db *gorm.DB, pid uint64) (*MitraCompany, error) {
	var err error
	db = db.Debug().Model(&MitraCompany{}).Where("mitra_company_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_company_status":     p.MitraCompanyStatus,
			"mitra_company_updated_by": p.MitraCompanyUpdatedBy,
			"mitra_company_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraCompany{}).Where("mitra_company_id = ?", pid).Error
	if err != nil {
		return &MitraCompany{}, err
	}
	return p, nil
}

func (p *MitraCompany) UpdateMitraCompanyStatusOnly(db *gorm.DB, pid uint64) (*MitraCompany, error) {
	var err error
	db = db.Debug().Model(&MitraCompany{}).Where("mitra_company_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_company_status": p.MitraCompanyStatus,
		},
	)
	err = db.Debug().Model(&MitraCompany{}).Where("mitra_company_id = ?", pid).Error
	if err != nil {
		return &MitraCompany{}, err
	}
	return p, nil
}

func (p *MitraCompany) UpdateMitraCompanyStatusDelete(db *gorm.DB, pid uint64) (*MitraCompany, error) {
	var err error
	db = db.Debug().Model(&MitraCompany{}).Where("mitra_company_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_company_status":     3,
			"mitra_company_deleted_by": p.MitraCompanyDeletedBy,
			"mitra_company_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraCompany{}).Where("mitra_company_id = ?", pid).Error
	if err != nil {
		return &MitraCompany{}, err
	}
	return p, nil
}

func (p *MitraCompany) DeleteMitraCompany(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraCompany{}).Where("mitra_company_id = ?", pid).Take(&MitraCompany{}).Delete(&MitraCompany{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraCompany not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================
