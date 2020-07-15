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

type CompanyTemp struct {
	CompanyTempID           uint64    `gorm:"column:company_temp_id;primary_key;" json:"company_temp_id"`
	CompanyID               uint64    `gorm:"column:company_id" json:"company_id"`
	CompanyTempCode         string    `gorm:"column:company_temp_code" json:"company_temp_code"`
	CompanyTempName         string    `gorm:"column:company_temp_name" json:"company_temp_name"`
	CompanyTempWebsite      string    `gorm:"column:company_temp_website" json:"company_temp_website"`
	CompanyTempEmail        string    `gorm:"column:company_temp_email" json:"company_temp_email"`
	CompanyTempLogo         string    `gorm:"column:company_temp_logo" json:"company_temp_logo"`
	CompanyTempReason       string    `gorm:"column:company_temp_reason" json:"company_temp_reason"`
	CompanyTempStatus       uint64    `gorm:"column:company_temp_status;size:2" json:"company_temp_status"`
	CompanyTempActionBefore uint64    `gorm:"column:company_temp_action_before;size:2" json:"company_temp_action_before"`
	CompanyTempCreatedBy    string    `gorm:"column:company_temp_created_by;size:125" json:"company_temp_created_by"`
	CompanyTempCreatedAt    time.Time `gorm:"column:company_temp_created_at;default:CURRENT_TIMESTAMP" json:"company_temp_created_at"`
}

type CompanyTempPend struct {
	CompanyTempID           uint64    `gorm:"column:company_temp_id;primary_key;" json:"company_temp_id"`
	CompanyID               *uint64   `gorm:"column:company_id" json:"company_id"`
	CompanyTempCode         string    `gorm:"column:company_temp_code" json:"company_temp_code"`
	CompanyTempName         string    `gorm:"column:company_temp_name" json:"company_temp_name"`
	CompanyTempWebsite      string    `gorm:"column:company_temp_website" json:"company_temp_website"`
	CompanyTempEmail        string    `gorm:"column:company_temp_email" json:"company_temp_email"`
	CompanyTempLogo         string    `gorm:"column:company_temp_logo" json:"company_temp_logo"`
	CompanyTempReason       string    `gorm:"column:company_temp_reason" json:"company_temp_reason"`
	CompanyTempStatus       uint64    `gorm:"column:company_temp_status;size:2" json:"company_temp_status"`
	CompanyTempActionBefore uint64    `gorm:"column:company_temp_action_before;size:2" json:"company_temp_action_before"`
	CompanyTempCreatedBy    string    `gorm:"column:company_temp_created_by;size:125" json:"company_temp_created_by"`
	CompanyTempCreatedAt    time.Time `gorm:"column:company_temp_created_at;default:CURRENT_TIMESTAMP" json:"company_temp_created_at"`
}

type CompanyTempData struct {
	CompanyTempID           uint64    `gorm:"column:company_temp_id" json:"company_temp_id"`
	CompanyTempCode         string    `gorm:"column:company_temp_code" json:"company_temp_code"`
	CompanyTempName         string    `gorm:"column:company_temp_name" json:"company_temp_name"`
	CompanyTempWebsite      string    `gorm:"column:company_temp_website" json:"company_temp_website"`
	CompanyTempEmail        string    `gorm:"column:company_temp_email" json:"company_temp_email"`
	CompanyTempLogo         string    `gorm:"column:company_temp_logo" json:"company_temp_logo"`
	CompanyTempReason       string    `gorm:"column:company_temp_reason" json:"company_temp_reason"`
	CompanyTempStatus       uint64    `gorm:"column:company_temp_status;size:2" json:"company_temp_status"`
	CompanyTempActionBefore uint64    `gorm:"column:company_temp_action_before;size:2" json:"company_temp_action_before"`
	CompanyTempCreatedBy    string    `gorm:"column:company_temp_created_by;size:125" json:"company_temp_created_by"`
	CompanyTempCreatedAt    time.Time `gorm:"column:company_temp_created_at" json:"company_temp_created_at"`
	CompanyID               uint64    `gorm:"column:company_id" json:"company_id"`
	CompanyCode             string    `gorm:"column:company_code" json:"company_code"`
	CompanyReferalCode      string    `gorm:"column:company_referal_code" json:"company_referal_code"`
	CompanyName             string    `gorm:"column:company_name" json:"company_name"`
	CompanyWebsite          string    `gorm:"column:company_website" json:"company_website"`
	CompanyEmail            string    `gorm:"column:company_email" json:"company_email"`
	CompanyLogo             string    `gorm:"column:company_logo" json:"company_logo"`
	CompanyStatus           uint64    `gorm:"column:company_status" json:"company_status"`
	CompanyCreatedBy        string    `gorm:"column:company_created_by" json:"company_created_by"`
	CompanyCreatedAt        time.Time `gorm:"column:company_created_at" json:"company_created_at"`
	CompanyUpdatedBy        string    `gorm:"column:company_updated_by" json:"company_updated_by"`
	CompanyUpdatedAt        time.Time `gorm:"column:company_updated_at" json:"company_updated_at"`
	CompanyDeletedBy        string    `gorm:"column:company_deleted_by" json:"company_deleted_by"`
	CompanyDeletedAt        time.Time `gorm:"column:company_deleted_at" json:"company_deleted_at"`
}

type ResponseCompanyTemps struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    *[]CompanyTempData `json:"data"`
}

type ResponseCompanyTemp struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    *CompanyTempData `json:"data"`
}

type ResponseCompanyTempsPend struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    *[]CompanyTempPend `json:"data"`
}

type ResponseCompanyTempPend struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    *CompanyTempPend `json:"data"`
}

type ResponseCompanyTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (CompanyTemp) TableName() string {
	return "m_company_temp"
}

func (CompanyTempPend) TableName() string {
	return "m_company_temp"
}

func (CompanyTempData) TableName() string {
	return "m_company_temp"
}

func (p *CompanyTemp) Prepare() {
	p.CompanyTempCode = html.EscapeString(strings.TrimSpace(p.CompanyTempCode))
	p.CompanyTempName = html.EscapeString(strings.TrimSpace(p.CompanyTempName))
	p.CompanyTempWebsite = p.CompanyTempWebsite
	p.CompanyTempEmail = p.CompanyTempEmail
	p.CompanyTempLogo = p.CompanyTempLogo
	p.CompanyTempStatus = p.CompanyTempStatus
	p.CompanyTempCreatedBy = html.EscapeString(strings.TrimSpace(p.CompanyTempCreatedBy))
	p.CompanyTempCreatedAt = time.Now()
}

func (p *CompanyTempPend) Prepare() {
	p.CompanyID = nil
	p.CompanyTempCode = html.EscapeString(strings.TrimSpace(p.CompanyTempCode))
	p.CompanyTempName = html.EscapeString(strings.TrimSpace(p.CompanyTempName))
	p.CompanyTempWebsite = p.CompanyTempWebsite
	p.CompanyTempEmail = p.CompanyTempEmail
	p.CompanyTempLogo = p.CompanyTempLogo
	p.CompanyTempReason = p.CompanyTempReason
	p.CompanyTempStatus = p.CompanyTempStatus
	p.CompanyTempCreatedBy = html.EscapeString(strings.TrimSpace(p.CompanyTempCreatedBy))
	p.CompanyTempCreatedAt = time.Now()
}

func (p *CompanyTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	case "insertupdate":
		if p.CompanyTempCode == "" {
			return errors.New("Required Country")
		}
		if p.CompanyTempName == "" {
			return errors.New("Required Regency")
		}
		if p.CompanyTempWebsite == "" {
			return errors.New("Required District")
		}
		if p.CompanyTempEmail == "" {
			return errors.New("Required Village")
		}
		if p.CompanyTempLogo == "" {
			return errors.New("Required Village")
		}
		return nil

	case "reason":
		if p.CompanyTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.CompanyTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	default:
		if p.CompanyTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil
	}
}

func (p *CompanyTemp) SaveCompanyTemp(db *gorm.DB) (*CompanyTemp, error) {
	var err error
	err = db.Debug().Model(&CompanyTemp{}).Create(&p).Error
	if err != nil {
		return &CompanyTemp{}, err
	}
	return p, nil
}

func (p *CompanyTempPend) SaveCompanyTempPend(db *gorm.DB) (*CompanyTempPend, error) {
	var err error
	err = db.Debug().Model(&CompanyTempPend{}).Create(&p).Error
	if err != nil {
		return &CompanyTempPend{}, err
	}
	return p, nil
}

func (p *CompanyTemp) FindAllCompanyTemp(db *gorm.DB) (*[]CompanyTempPend, error) {
	var err error
	company := []CompanyTempPend{}
	err = db.Debug().Model(&CompanyTempPend{}).Limit(100).
		Select(`m_company_temp.company_temp_id,
				m_company_temp.company_id,
				m_company_temp.company_temp_code,
				m_company_temp.company_temp_name,
				m_company_temp.company_temp_website,
				m_company_temp.company_temp_email,
				m_company_temp.company_temp_logo,
				m_company_temp.company_temp_reason,
				m_company_temp.company_temp_status,
				m_company_temp.company_temp_action_before,
				m_company_temp.company_temp_created_by,
				m_company_temp.company_temp_created_at at time zone 'Asia/Jakarta' as company_temp_created_at`).
		Order("company_temp_created_at desc").
		Find(&company).Error
	if err != nil {
		return &[]CompanyTempPend{}, err
	}
	return &company, nil
}

func (p *CompanyTemp) FindAllCompanyTempStatusPendingActive(db *gorm.DB) (*[]CompanyTempPend, error) {
	var err error
	company := []CompanyTempPend{}
	err = db.Debug().Model(&CompanyTempPend{}).Limit(100).
		Select(`m_company_temp.company_temp_id,
				m_company_temp.company_id,
				m_company_temp.company_temp_code,
				m_company_temp.company_temp_name,
				m_company_temp.company_temp_website,
				m_company_temp.company_temp_email,
				m_company_temp.company_temp_logo,
				m_company_temp.company_temp_reason,
				m_company_temp.company_temp_status,
				m_company_temp.company_temp_action_before,
				m_company_temp.company_temp_created_by,
				m_company_temp.company_temp_created_at at time zone 'Asia/Jakarta' as company_temp_created_at`).
		Where("company_temp_status = ?", 11).
		Order("company_temp_created_at desc").
		Find(&company).Error
	if err != nil {
		return &[]CompanyTempPend{}, err
	}
	return &company, nil
}

func (p *CompanyTemp) FindAllCompanyTempStatus(db *gorm.DB, status uint64) (*[]CompanyTempData, error) {
	var err error
	company := []CompanyTempData{}
	err = db.Debug().Model(&CompanyTempData{}).Limit(100).
		Select(`m_company_temp.company_temp_id,
				m_company_temp.company_temp_code,
				m_company_temp.company_temp_name,
				m_company_temp.company_temp_website,
				m_company_temp.company_temp_email,
				m_company_temp.company_temp_logo,
				m_company_temp.company_temp_reason,
				m_company_temp.company_temp_status,
				m_company_temp.company_temp_action_before,
				m_company_temp.company_temp_created_by,
				m_company_temp.company_temp_created_at at time zone 'Asia/Jakarta' as company_temp_created_at,
				m_company.company_id,
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
		Joins("JOIN m_company on m_company_temp.company_id=m_company.company_id").
		Where("company_temp_status = ?", status).
		Order("company_temp_created_at desc").
		Find(&company).Error
	if err != nil {
		return &[]CompanyTempData{}, err
	}
	return &company, nil
}

func (p *CompanyTemp) FindCompanyTempDataByID(db *gorm.DB, pid uint64) (*CompanyTemp, error) {
	var err error
	err = db.Debug().Model(&CompanyTemp{}).Where("company_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &CompanyTemp{}, err
	}
	return p, nil
}

func (p *CompanyTemp) FindCompanyTempByIDPendingActive(db *gorm.DB, pid uint64) (*CompanyTempPend, error) {
	var err error
	company := CompanyTempPend{}
	err = db.Debug().Model(&CompanyTempPend{}).Limit(100).
		Select(`m_company_temp.company_temp_id,
				m_company_temp.company_id,
				m_company_temp.company_temp_code,
				m_company_temp.company_temp_name,
				m_company_temp.company_temp_website,
				m_company_temp.company_temp_email,
				m_company_temp.company_temp_logo,
				m_company_temp.company_temp_reason,
				m_company_temp.company_temp_status,
				m_company_temp.company_temp_created_by,
				m_company_temp.company_temp_created_at at time zone 'Asia/Jakarta' as company_temp_created_at`).
		Where("company_temp_id = ?", pid).
		Order("company_temp_created_at desc").
		Find(&company).Error
	if err != nil {
		return &CompanyTempPend{}, err
	}
	return &company, nil
}

func (p *CompanyTemp) FindCompanyTempStatusByIDPendingActive(db *gorm.DB, pid uint64) (*CompanyTempPend, error) {
	var err error
	company := CompanyTempPend{}
	err = db.Debug().Model(&CompanyTempPend{}).Limit(100).
		Select(`m_company_temp.company_temp_id,
				m_company_temp.company_id,
				m_company_temp.company_temp_code,
				m_company_temp.company_temp_name,
				m_company_temp.company_temp_website,
				m_company_temp.company_temp_email,
				m_company_temp.company_temp_logo,
				m_company_temp.company_temp_reason,
				m_company_temp.company_temp_status,
				m_company_temp.company_temp_action_before,
				m_company_temp.company_temp_created_by,
				m_company_temp.company_temp_created_at at time zone 'Asia/Jakarta' as company_temp_created_at`).
		Where("company_temp_id = ? and company_temp_status = ?", pid, 11).
		Order("company_temp_created_at desc").
		Find(&company).Error
	if err != nil {
		return &CompanyTempPend{}, err
	}
	return &company, nil
}

func (p *CompanyTemp) FindCompanyTempByID(db *gorm.DB, pid uint64) (*CompanyTempData, error) {
	var err error
	company := CompanyTempData{}
	err = db.Debug().Model(&CompanyTempData{}).Limit(100).
		Select(`m_company_temp.company_temp_id,
				m_company_temp.company_temp_code,
				m_company_temp.company_temp_name,
				m_company_temp.company_temp_website,
				m_company_temp.company_temp_email,
				m_company_temp.company_temp_logo,
				m_company_temp.company_temp_reason,
				m_company_temp.company_temp_status,
				m_company_temp.company_temp_action_before,
				m_company_temp.company_temp_created_by,
				m_company_temp.company_temp_created_at at time zone 'Asia/Jakarta' as company_temp_created_at,
				m_company.company_id,
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
		Joins("JOIN m_company on m_company_temp.company_id=m_company.company_id").
		Where("company_temp_id = ?", pid).
		Order("company_temp_created_at desc").
		Find(&company).Error
	if err != nil {
		return &CompanyTempData{}, err
	}
	return &company, nil
}

func (p *CompanyTemp) FindCompanyTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*CompanyTempData, error) {
	var err error
	company := CompanyTempData{}
	err = db.Debug().Model(&CompanyTempData{}).Limit(100).
		Select(`m_company_temp.company_temp_id,
				m_company_temp.company_temp_code,
				m_company_temp.company_temp_name,
				m_company_temp.company_temp_website,
				m_company_temp.company_temp_email,
				m_company_temp.company_temp_logo,
				m_company_temp.company_temp_reason,
				m_company_temp.company_temp_status,
				m_company_temp.company_temp_action_before,
				m_company_temp.company_temp_created_by,
				m_company_temp.company_temp_created_at at time zone 'Asia/Jakarta' as company_temp_created_at,
				m_company.company_id,
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
		Joins("JOIN m_company on m_company_temp.company_id=m_company.company_id").
		Where("company_temp_id = ? AND company_temp_status = ?", pid, status).
		Order("company_temp_created_at desc").
		Find(&company).Error
	if err != nil {
		return &CompanyTempData{}, err
	}
	return &company, nil
}

func (p *CompanyTemp) UpdateCompanyTemp(db *gorm.DB, pid uint64) (*CompanyTemp, error) {
	var err error
	db = db.Debug().Model(&CompanyTemp{}).Where("company_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"company_temp_code":    p.CompanyTempCode,
			"company_temp_name":    p.CompanyTempName,
			"company_temp_website": p.CompanyTempWebsite,
			"company_temp_email":   p.CompanyTempEmail,
			"company_temp_logo":    p.CompanyTempLogo,
			"company_temp_reason":  p.CompanyTempReason,
		},
	)
	err = db.Debug().Model(&CompanyTemp{}).Where("company_temp_id = ?", pid).Error
	if err != nil {
		return &CompanyTemp{}, err
	}
	return p, nil
}

func (p *CompanyTemp) UpdateCompanyTempStatus(db *gorm.DB, pid uint64) (*CompanyTemp, error) {
	var err error
	db = db.Debug().Model(&CompanyTemp{}).Where("company_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"company_temp_reason": p.CompanyTempReason,
			"company_temp_status": p.CompanyTempStatus,
		},
	)
	err = db.Debug().Model(&CompanyTemp{}).Where("company_temp_id = ?", pid).Error
	if err != nil {
		return &CompanyTemp{}, err
	}
	return p, nil
}

func (p *CompanyTemp) DeleteCompanyTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&CompanyTemp{}).Where("company_temp_id = ?", pid).Take(&CompanyTemp{}).Delete(&CompanyTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("CompanyTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================
