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

type Company struct {
	CompanyID          uint64     `gorm:"column:company_id;primary_key;" json:"company_id"`
	CompanyCode        string     `gorm:"column:company_code" json:"company_code"`
	CompanyReferalCode string     `gorm:"column:company_referal_code" json:"company_referal_code"`
	CompanyName        string     `gorm:"column:company_name" json:"company_name"`
	CompanyWebsite     string     `gorm:"column:company_website" json:"company_website"`
	CompanyEmail       string     `gorm:"column:company_email" json:"company_email"`
	CompanyLogo        string     `gorm:"column:company_logo" json:"company_logo"`
	CompanyStatus      uint64     `gorm:"column:company_status;size:2" json:"company_status"`
	CompanyCreatedBy   string     `gorm:"column:company_created_by;size:125" json:"company_created_by"`
	CompanyCreatedAt   time.Time  `gorm:"column:company_created_at;default:CURRENT_TIMESTAMP" json:"company_created_at"`
	CompanyUpdatedBy   string     `gorm:"column:company_updated_by;size:125" json:"company_updated_by"`
	CompanyUpdatedAt   *time.Time `gorm:"column:company_updated_at;default:CURRENT_TIMESTAMP" json:"company_created_at"`
	CompanyDeletedBy   string     `gorm:"column:company_deleted_by;size:125" json:"company_deleted_by"`
	CompanyDeletedAt   *time.Time `gorm:"column:company_deleted_at;default:CURRENT_TIMESTAMP" json:"company_deleted_at"`
}

type CompanyData struct {
	CompanyID          uint64     `gorm:"column:company_id" json:"company_id"`
	CompanyCode        string     `gorm:"column:company_code" json:"company_code"`
	CompanyReferalCode string     `gorm:"column:company_referal_code" json:"company_referal_code"`
	CompanyName        string     `gorm:"column:company_name" json:"company_name"`
	CompanyWebsite     string     `gorm:"column:company_website" json:"company_website"`
	CompanyEmail       string     `gorm:"column:company_email" json:"company_email"`
	CompanyLogo        string     `gorm:"column:company_logo" json:"company_logo"`
	CompanyStatus      uint64     `gorm:"column:company_status;size:2" json:"company_status"`
	CompanyCreatedBy   string     `gorm:"column:company_created_by" json:"company_created_by"`
	CompanyCreatedAt   time.Time  `gorm:"column:company_created_at" json:"company_created_at"`
	CompanyUpdatedBy   string     `gorm:"column:company_updated_by" json:"company_updated_by"`
	CompanyUpdatedAt   *time.Time `gorm:"column:company_updated_at" json:"company_updated_at"`
	CompanyDeletedBy   string     `gorm:"column:company_deleted_by" json:"company_deleted_by"`
	CompanyDeletedAt   *time.Time `gorm:"column:company_deleted_at" json:"company_deleted_at"`
}

type ResponseCompanys struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    *[]CompanyData `json:"data"`
}

type ResponseCompany struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    *CompanyData `json:"data"`
}

type ResponseCompanyCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (Company) TableName() string {
	return "m_company"
}

func (CompanyData) TableName() string {
	return "m_company"
}

func (p *Company) Prepare() {
	p.CompanyCode = html.EscapeString(strings.TrimSpace(p.CompanyCode))
	p.CompanyReferalCode = p.CompanyReferalCode
	p.CompanyName = html.EscapeString(strings.TrimSpace(p.CompanyName))
	p.CompanyWebsite = p.CompanyWebsite
	p.CompanyEmail = p.CompanyEmail
	p.CompanyLogo = p.CompanyLogo
	p.CompanyStatus = p.CompanyStatus
	p.CompanyCreatedBy = html.EscapeString(strings.TrimSpace(p.CompanyCreatedBy))
	p.CompanyCreatedAt = time.Now()
	p.CompanyUpdatedBy = html.EscapeString(strings.TrimSpace(p.CompanyUpdatedBy))
	p.CompanyUpdatedAt = p.CompanyUpdatedAt
	p.CompanyDeletedBy = html.EscapeString(strings.TrimSpace(p.CompanyDeletedBy))
	p.CompanyDeletedAt = p.CompanyDeletedAt
}

func (p *Company) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.CompanyCode == "" {
			return errors.New("Required Country")
		}
		if p.CompanyReferalCode == "" {
			return errors.New("Required Province")
		}
		if p.CompanyName == "" {
			return errors.New("Required Regency")
		}
		if p.CompanyWebsite == "" {
			return errors.New("Required District")
		}
		if p.CompanyEmail == "" {
			return errors.New("Required Village")
		}
		if p.CompanyLogo == "" {
			return errors.New("Required Village")
		}
		return nil
	}
}

func (p *Company) SaveCompany(db *gorm.DB) (*Company, error) {
	var err error
	err = db.Debug().Model(&Company{}).Create(&p).Error
	if err != nil {
		return &Company{}, err
	}
	return p, nil
}

func (p *Company) FindAllCompany(db *gorm.DB) (*[]CompanyData, error) {
	var err error
	company := []CompanyData{}
	err = db.Debug().Model(&CompanyData{}).Limit(100).
		Select(`m_company.company_id,
				m_company.company_code,
				m_company.company_referal_code,
				m_company.company_name,
				m_company.company_website,
				m_company.company_email,
				m_company.company_logo,
				m_company.company_status,
				m_company.company_created_by,
				m_company.company_created_at at time zone 'Asia/Jakarta' as company_created_at,
				m_company.company_updated_by,
				m_company.company_updated_at at time zone 'Asia/Jakarta' as company_updated_at,
				m_company.company_deleted_by,
				m_company.company_deleted_at at time zone 'Asia/Jakarta' as company_deleted_at`).
		Find(&company).Error
	if err != nil {
		return &[]CompanyData{}, err
	}
	return &company, nil
}

func (p *Company) FindAllCompanyStatus(db *gorm.DB, status uint64) (*[]CompanyData, error) {
	var err error
	company := []CompanyData{}
	err = db.Debug().Model(&CompanyData{}).Limit(100).
		Select(`m_company.company_id,
				m_company.company_code,
				m_company.company_referal_code,
				m_company.company_name,
				m_company.company_website,
				m_company.company_email,
				m_company.company_logo,
				m_company.company_status,
				m_company.company_created_by,
				m_company.company_created_at at time zone 'Asia/Jakarta' as company_created_at,
				m_company.company_updated_by,
				m_company.company_updated_at at time zone 'Asia/Jakarta' as company_updated_at,
				m_company.company_deleted_by,
				m_company.company_deleted_at at time zone 'Asia/Jakarta' as company_deleted_at`).
		Where("company_status = ?", status).
		Find(&company).Error
	if err != nil {
		return &[]CompanyData{}, err
	}
	return &company, nil
}

func (p *Company) FindCompanyDataByID(db *gorm.DB, pid uint64) (*Company, error) {
	var err error
	err = db.Debug().Model(&Company{}).
		Select(`m_company.company_id,
				m_company.company_code,
				m_company.company_referal_code,
				m_company.company_name,
				m_company.company_website,
				m_company.company_email,
				m_company.company_logo,
				m_company.company_status,
				m_company.company_created_by,
				m_company.company_created_at at time zone 'Asia/Jakarta' as company_created_at,
				m_company.company_updated_by,
				m_company.company_updated_at at time zone 'Asia/Jakarta' as company_updated_at,
				m_company.company_deleted_by,
				m_company.company_deleted_at at time zone 'Asia/Jakarta' as company_deleted_at`).
		Where("company_id = ?", pid).
		Take(&p).Error
	if err != nil {
		return &Company{}, err
	}
	return p, nil
}

func (p *Company) FindCompanyByID(db *gorm.DB, pid uint64) (*CompanyData, error) {
	var err error
	company := CompanyData{}
	err = db.Debug().Model(&CompanyData{}).Limit(100).
		Select(`m_company.company_id,
				m_company.company_code,
				m_company.company_referal_code,
				m_company.company_name,
				m_company.company_website,
				m_company.company_email,
				m_company.company_logo,
				m_company.company_status,
				m_company.company_created_by,
				m_company.company_created_at at time zone 'Asia/Jakarta' as company_created_at,
				m_company.company_updated_by,
				m_company.company_updated_at at time zone 'Asia/Jakarta' as company_updated_at,
				m_company.company_deleted_by,
				m_company.company_deleted_at at time zone 'Asia/Jakarta' as company_deleted_at`).
		Where("company_id = ?", pid).
		Find(&company).Error
	if err != nil {
		return &CompanyData{}, err
	}
	return &company, nil
}

func (p *Company) FindCompanyStatusByID(db *gorm.DB, pid uint64, status uint64) (*CompanyData, error) {
	var err error
	company := CompanyData{}
	err = db.Debug().Model(&CompanyData{}).Limit(100).
		Select(`m_company.company_id,
				m_company.company_code,
				m_company.company_referal_code,
				m_company.company_name,
				m_company.company_website,
				m_company.company_email,
				m_company.company_logo,
				m_company.company_status,
				m_company.company_created_by,
				m_company.company_created_at at time zone 'Asia/Jakarta' as company_created_at,
				m_company.company_updated_by,
				m_company.company_updated_at at time zone 'Asia/Jakarta' as company_updated_at,
				m_company.company_deleted_by,
				m_company.company_deleted_at at time zone 'Asia/Jakarta' as company_deleted_at`).
		Where("company_id = ? AND company_status = ?", pid, status).
		Find(&company).Error
	if err != nil {
		return &CompanyData{}, err
	}
	return &company, nil
}

func (p *Company) UpdateCompany(db *gorm.DB, pid uint64) (*Company, error) {
	var err error
	db = db.Debug().Model(&Company{}).Where("company_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"company_code": p.CompanyCode,
			// "company_referal_code": p.CompanyReferalCode,
			"company_name":       p.CompanyName,
			"company_website":    p.CompanyWebsite,
			"company_email":      p.CompanyEmail,
			"company_logo":       p.CompanyLogo,
			"company_status":     p.CompanyStatus,
			"company_updated_by": p.CompanyUpdatedBy,
			"company_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&Company{}).Where("company_id = ?", pid).Error
	if err != nil {
		return &Company{}, err
	}
	return p, nil
}

func (p *Company) UpdateCompanyStatus(db *gorm.DB, pid uint64) (*Company, error) {
	var err error
	db = db.Debug().Model(&Company{}).Where("company_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"company_status":     p.CompanyStatus,
			"company_updated_by": p.CompanyUpdatedBy,
			"company_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&Company{}).Where("company_id = ?", pid).Error
	if err != nil {
		return &Company{}, err
	}
	return p, nil
}

func (p *Company) UpdateCompanyStatusOnly(db *gorm.DB, pid uint64) (*Company, error) {
	var err error
	db = db.Debug().Model(&Company{}).Where("company_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"company_status": p.CompanyStatus,
		},
	)
	err = db.Debug().Model(&Company{}).Where("company_id = ?", pid).Error
	if err != nil {
		return &Company{}, err
	}
	return p, nil
}

func (p *Company) UpdateCompanyStatusDelete(db *gorm.DB, pid uint64) (*Company, error) {
	var err error
	db = db.Debug().Model(&Company{}).Where("company_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"company_status":     3,
			"company_deleted_by": p.CompanyDeletedBy,
			"company_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&Company{}).Where("company_id = ?", pid).Error
	if err != nil {
		return &Company{}, err
	}
	return p, nil
}

func (p *Company) DeleteCompany(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&Company{}).Where("company_id = ?", pid).Take(&Company{}).Delete(&Company{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Company not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

func (p *Company) FindAllCompanyReferalCode(db *gorm.DB, referal string) (*Company, error) {
	var err error
	err = db.Debug().Model(&Company{}).Where("company_referal_code = ? AND mitra_status = ?", referal, 1).Take(&p).Error
	if err != nil {
		return &Company{}, err
	}
	return p, nil
}
