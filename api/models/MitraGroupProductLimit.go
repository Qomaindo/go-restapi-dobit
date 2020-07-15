package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type MitraGroupProductLimit struct {
	MitraGroupProductLimitID        uint64     `gorm:"column:mitra_group_prod_lim_id;primary_key;" json:"mitra_group_prod_lim_id"`
	MitraID                         uint64     `gorm:"column:mitra_id" json:"mitra_id"`
	MitraGroupProductLimitCode      string     `gorm:"column:mitra_group_prod_lim_code;size:25" json:"mitra_group_prod_lim_code"`
	MitraGroupProductLimitName      string     `gorm:"column:mitra_group_prod_lim_name;size:255" json:"mitra_group_prod_lim_name"`
	MitraGroupProductLimitStatus    uint64     `gorm:"column:mitra_group_prod_lim_status;size:2" json:"mitra_group_prod_lim_status"`
	MitraGroupProductLimitCreatedBy string     `gorm:"column:mitra_group_prod_lim_created_by;size:125" json:"mitra_group_prod_lim_created_by"`
	MitraGroupProductLimitCreatedAt time.Time  `gorm:"column:mitra_group_prod_lim_created_at;default:CURRENT_TIMESTAMP" json:"mitra_group_prod_lim_created_at"`
	MitraGroupProductLimitUpdatedBy string     `gorm:"column:mitra_group_prod_lim_updated_by;size:125" json:"mitra_group_prod_lim_updated_by"`
	MitraGroupProductLimitUpdatedAt *time.Time `gorm:"column:mitra_group_prod_lim_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_group_prod_lim_updated_at"`
	MitraGroupProductLimitDeletedBy string     `gorm:"column:mitra_group_prod_lim_deleted_by;size:125" json:"mitra_group_prod_lim_deleted_by"`
	MitraGroupProductLimitDeletedAt *time.Time `gorm:"column:mitra_group_prod_lim_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_group_prod_lim_deleted_at"`
}

type MitraGroupProductLimitData struct {
	MitraGroupProductLimitID        uint64                              `gorm:"column:mitra_group_prod_lim_id" json:"mitra_group_prod_lim_id"`
	MitraID                         uint64                              `gorm:"column:mitra_id" json:"mitra_id"`
	MitraCode                       string                              `gorm:"column:mitra_code" json:"mitra_code"`
	MitraName                       string                              `gorm:"column:mitra_name" json:"mitra_name"`
	MitraWebsite                    string                              `gorm:"column:mitra_website" json:"mitra_website"`
	MitraEmail                      string                              `gorm:"column:mitra_email" json:"mitra_email"`
	MitraLogo                       string                              `gorm:"column:mitra_logo" json:"mitra_logo"`
	MitraGroupProductLimitCode      string                              `gorm:"column:mitra_group_prod_lim_code;size:25" json:"mitra_group_prod_lim_code"`
	MitraGroupProductLimitName      string                              `gorm:"column:mitra_group_prod_lim_name;size:255" json:"mitra_group_prod_lim_name"`
	MitraGroupProductLimitStatus    uint64                              `gorm:"column:mitra_group_prod_lim_status;size:2" json:"mitra_group_prod_lim_status"`
	MitraGroupProductLimitCreatedBy string                              `gorm:"column:mitra_group_prod_lim_created_by;size:125" json:"mitra_group_prod_lim_created_by"`
	MitraGroupProductLimitCreatedAt time.Time                           `gorm:"column:mitra_group_prod_lim_created_at;default:CURRENT_TIMESTAMP" json:"mitra_group_prod_lim_created_at"`
	MitraGroupProductLimitUpdatedBy string                              `gorm:"column:mitra_group_prod_lim_updated_by;size:125" json:"mitra_group_prod_lim_updated_by"`
	MitraGroupProductLimitUpdatedAt *time.Time                          `gorm:"column:mitra_group_prod_lim_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_group_prod_lim_updated_at"`
	MitraGroupProductLimitDeletedBy string                              `gorm:"column:mitra_group_prod_lim_deleted_by;size:125" json:"mitra_group_prod_lim_deleted_by"`
	MitraGroupProductLimitDeletedAt *time.Time                          `gorm:"column:mitra_group_prod_lim_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_group_prod_lim_deleted_at"`
	MitraGroupProductNominalLimit   []MitraGroupProductNominalLimitData `json:"mitra_group_prod_lim_nominal"`
}

type ResponseMitraGroupProductLimits struct {
	Status  int                           `json:"status"`
	Message string                        `json:"message"`
	Data    *[]MitraGroupProductLimitData `json:"data"`
}

type ResponseMitraGroupProductLimit struct {
	Status  int                         `json:"status"`
	Message string                      `json:"message"`
	Data    *MitraGroupProductLimitData `json:"data"`
}

type ResponseMitraGroupProductLimitCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraGroupProductLimit) TableName() string {
	return "m_mitra_group_product_limit"
}

func (MitraGroupProductLimitData) TableName() string {
	return "m_mitra_group_product_limit"
}

func (p *MitraGroupProductLimit) Prepare() {
	p.MitraID = p.MitraID
	p.MitraGroupProductLimitCode = p.MitraGroupProductLimitCode
	p.MitraGroupProductLimitName = p.MitraGroupProductLimitName
	p.MitraGroupProductLimitStatus = p.MitraGroupProductLimitStatus
	p.MitraGroupProductLimitCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraGroupProductLimitCreatedBy))
	p.MitraGroupProductLimitCreatedAt = time.Now()
	p.MitraGroupProductLimitUpdatedBy = html.EscapeString(strings.TrimSpace(p.MitraGroupProductLimitUpdatedBy))
	p.MitraGroupProductLimitUpdatedAt = p.MitraGroupProductLimitUpdatedAt
	p.MitraGroupProductLimitDeletedBy = html.EscapeString(strings.TrimSpace(p.MitraGroupProductLimitDeletedBy))
	p.MitraGroupProductLimitDeletedAt = p.MitraGroupProductLimitDeletedAt
}

func (p *MitraGroupProductLimit) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.MitraGroupProductLimitCode == "" {
			return errors.New("Required MitraGroupProductLimit Code")
		}
		if p.MitraGroupProductLimitName == "" {
			return errors.New("Required MitraGroupProductLimit Name")
		}
		return nil
	}
}

func (p *MitraGroupProductLimit) SaveMitraGroupProductLimit(db *gorm.DB) (*MitraGroupProductLimit, error) {
	var err error
	err = db.Debug().Model(&MitraGroupProductLimit{}).Create(&p).Error
	if err != nil {
		return &MitraGroupProductLimit{}, err
	}
	return p, nil
}

func (p *MitraGroupProductLimit) FindAllMitraGroupProductLimit(db *gorm.DB, mitra uint64) (*[]MitraGroupProductLimitData, error) {
	var err error
	mitragroupproductlimit := []MitraGroupProductLimitData{}
	err = db.Debug().Model(&MitraGroupProductLimitData{}).Limit(100).
		Select(`m_mitra_group_product_limit.mitra_group_prod_lim_id,
				m_mitra_group_product_limit.mitra_id,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra.mitra_website,
				m_mitra.mitra_email,
				m_mitra.mitra_logo,
				m_mitra_group_product_limit.mitra_group_prod_lim_code,
				m_mitra_group_product_limit.mitra_group_prod_lim_name,
				m_mitra_group_product_limit.mitra_group_prod_lim_status,
				m_mitra_group_product_limit.mitra_group_prod_lim_created_by,
				m_mitra_group_product_limit.mitra_group_prod_lim_updated_by,
				m_mitra_group_product_limit.mitra_group_prod_lim_deleted_by,
				m_mitra_group_product_limit.mitra_group_prod_lim_created_at,
				m_mitra_group_product_limit.mitra_group_prod_lim_updated_at,
				m_mitra_group_product_limit.mitra_group_prod_lim_deleted_at`).
		Joins("JOIN m_mitra on m_mitra_group_product_limit.mitra_id=m_mitra.mitra_id").
		Where("m_mitra_group_product_limit.mitra_id = ?", mitra).
		Find(&mitragroupproductlimit).Error
	if err != nil {
		return &[]MitraGroupProductLimitData{}, err
	}
	return &mitragroupproductlimit, nil
}

func (p *MitraGroupProductLimit) FindAllMitraGroupProductLimitStatus(db *gorm.DB, status uint64, mitra uint64) (*[]MitraGroupProductLimitData, error) {
	var err error
	mitragroupproductlimit := []MitraGroupProductLimitData{}
	err = db.Debug().Model(&MitraGroupProductLimitData{}).Limit(100).
		Select(`m_mitra_group_product_limit.mitra_group_prod_lim_id,
				m_mitra_group_product_limit.mitra_id,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra.mitra_website,
				m_mitra.mitra_email,
				m_mitra.mitra_logo,
				m_mitra_group_product_limit.mitra_group_prod_lim_code,
				m_mitra_group_product_limit.mitra_group_prod_lim_name,
				m_mitra_group_product_limit.mitra_group_prod_lim_status,
				m_mitra_group_product_limit.mitra_group_prod_lim_created_by,
				m_mitra_group_product_limit.mitra_group_prod_lim_updated_by,
				m_mitra_group_product_limit.mitra_group_prod_lim_deleted_by,
				m_mitra_group_product_limit.mitra_group_prod_lim_created_at,
				m_mitra_group_product_limit.mitra_group_prod_lim_updated_at,
				m_mitra_group_product_limit.mitra_group_prod_lim_deleted_at`).
		Joins("JOIN m_mitra on m_mitra_group_product_limit.mitra_id=m_mitra.mitra_id").
		Where("mitra_group_prod_lim_status = ? AND m_mitra_group_product_limit.mitra_id = ?", status, mitra).
		Find(&mitragroupproductlimit).Error
	if err != nil {
		return &[]MitraGroupProductLimitData{}, err
	}
	return &mitragroupproductlimit, nil
}

func (p *MitraGroupProductLimit) FindMitraGroupProductLimitDataByID(db *gorm.DB, pid uint64) (*MitraGroupProductLimit, error) {
	var err error
	err = db.Debug().Model(&MitraGroupProductLimit{}).Where("mitra_group_prod_lim_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraGroupProductLimit{}, err
	}
	return p, nil
}

func (p *MitraGroupProductLimit) FindMitraGroupProductLimitByID(db *gorm.DB, pid uint64) (*MitraGroupProductLimitData, error) {
	var err error
	mitragroupproductlimit := MitraGroupProductLimitData{}
	err = db.Debug().Model(&MitraGroupProductLimitData{}).
		Select(`m_mitra_group_product_limit.mitra_group_prod_lim_id,
				m_mitra_group_product_limit.mitra_id,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra.mitra_website,
				m_mitra.mitra_email,
				m_mitra.mitra_logo,
				m_mitra_group_product_limit.mitra_group_prod_lim_code,
				m_mitra_group_product_limit.mitra_group_prod_lim_name,
				m_mitra_group_product_limit.mitra_group_prod_lim_status,
				m_mitra_group_product_limit.mitra_group_prod_lim_created_by,
				m_mitra_group_product_limit.mitra_group_prod_lim_updated_by,
				m_mitra_group_product_limit.mitra_group_prod_lim_deleted_by,
				m_mitra_group_product_limit.mitra_group_prod_lim_created_at,
				m_mitra_group_product_limit.mitra_group_prod_lim_updated_at,
				m_mitra_group_product_limit.mitra_group_prod_lim_deleted_at`).
		Joins("JOIN m_mitra on m_mitra_group_product_limit.mitra_id=m_mitra.mitra_id").
		Where("mitra_group_prod_lim_id = ?", pid).
		Take(&mitragroupproductlimit).Error
	if err != nil {
		return &MitraGroupProductLimitData{}, err
	}
	return &mitragroupproductlimit, nil
}

func (p *MitraGroupProductLimit) FindMitraGroupProductLimitByIDMitra(db *gorm.DB, pid uint64, mitra uint64) (*MitraGroupProductLimitData, error) {
	var err error
	mitragroupproductlimit := MitraGroupProductLimitData{}
	err = db.Debug().Model(&MitraGroupProductLimitData{}).
		Select(`m_mitra_group_product_limit.mitra_group_prod_lim_id,
				m_mitra_group_product_limit.mitra_id,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra.mitra_website,
				m_mitra.mitra_email,
				m_mitra.mitra_logo,
				m_mitra_group_product_limit.mitra_group_prod_lim_code,
				m_mitra_group_product_limit.mitra_group_prod_lim_name,
				m_mitra_group_product_limit.mitra_group_prod_lim_status,
				m_mitra_group_product_limit.mitra_group_prod_lim_created_by,
				m_mitra_group_product_limit.mitra_group_prod_lim_updated_by,
				m_mitra_group_product_limit.mitra_group_prod_lim_deleted_by,
				m_mitra_group_product_limit.mitra_group_prod_lim_created_at,
				m_mitra_group_product_limit.mitra_group_prod_lim_updated_at,
				m_mitra_group_product_limit.mitra_group_prod_lim_deleted_at`).
		Joins("JOIN m_mitra on m_mitra_group_product_limit.mitra_id=m_mitra.mitra_id").
		Where("mitra_group_prod_lim_id = ? AND m_mitra_group_product_limit.mitra_id = ?", pid, mitra).
		Take(&mitragroupproductlimit).Error
	if err != nil {
		return &MitraGroupProductLimitData{}, err
	}
	return &mitragroupproductlimit, nil
}

func (p *MitraGroupProductLimit) FindMitraGroupProductLimitStatusByID(db *gorm.DB, pid uint64, status uint64, mitra uint64) (*MitraGroupProductLimitData, error) {
	var err error
	mitragroupproductlimit := MitraGroupProductLimitData{}
	err = db.Debug().Model(&MitraGroupProductLimitData{}).
		Select(`m_mitra_group_product_limit.mitra_group_prod_lim_id,
				m_mitra_group_product_limit.mitra_id,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra.mitra_website,
				m_mitra.mitra_email,
				m_mitra.mitra_logo,
				m_mitra_group_product_limit.mitra_group_prod_lim_code,
				m_mitra_group_product_limit.mitra_group_prod_lim_name,
				m_mitra_group_product_limit.mitra_group_prod_lim_status,
				m_mitra_group_product_limit.mitra_group_prod_lim_created_by,
				m_mitra_group_product_limit.mitra_group_prod_lim_updated_by,
				m_mitra_group_product_limit.mitra_group_prod_lim_deleted_by,
				m_mitra_group_product_limit.mitra_group_prod_lim_created_at,
				m_mitra_group_product_limit.mitra_group_prod_lim_updated_at,
				m_mitra_group_product_limit.mitra_group_prod_lim_deleted_at`).
		Joins("JOIN m_mitra on m_mitra_group_product_limit.mitra_id=m_mitra.mitra_id").
		Where("mitra_group_prod_lim_id = ? AND mitra_group_prod_lim_status = ? AND m_mitra_group_product_limit.mitra_id = ?", pid, status, mitra).
		Take(&mitragroupproductlimit).Error
	if err != nil {
		return &MitraGroupProductLimitData{}, err
	}
	return &mitragroupproductlimit, nil
}

func (p *MitraGroupProductLimit) UpdateMitraGroupProductLimit(db *gorm.DB, pid uint64) (*MitraGroupProductLimit, error) {

	var err error
	db = db.Debug().Model(&MitraGroupProductLimit{}).Where("mitra_group_prod_lim_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_id":                        p.MitraID,
			"mitra_group_prod_lim_code":       p.MitraGroupProductLimitCode,
			"mitra_group_prod_lim_name":       p.MitraGroupProductLimitName,
			"mitra_group_prod_lim_status":     p.MitraGroupProductLimitStatus,
			"mitra_group_prod_lim_updated_by": p.MitraGroupProductLimitUpdatedBy,
			"mitra_group_prod_lim_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraGroupProductLimit{}).Where("mitra_group_prod_lim_id = ?", pid).Error
	if err != nil {
		return &MitraGroupProductLimit{}, err
	}
	return p, nil
}

func (p *MitraGroupProductLimit) UpdateMitraGroupProductLimitStatus(db *gorm.DB, pid uint64) (*MitraGroupProductLimit, error) {

	var err error
	db = db.Debug().Model(&MitraGroupProductLimit{}).Where("mitra_group_prod_lim_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_group_prod_lim_status":     p.MitraGroupProductLimitStatus,
			"mitra_group_prod_lim_updated_by": p.MitraGroupProductLimitUpdatedBy,
			"mitra_group_prod_lim_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraGroupProductLimit{}).Where("mitra_group_prod_lim_id = ?", pid).Error
	if err != nil {
		return &MitraGroupProductLimit{}, err
	}
	return p, nil
}

func (p *MitraGroupProductLimit) UpdateMitraGroupProductLimitStatusOnly(db *gorm.DB, pid uint64) (*MitraGroupProductLimit, error) {
	var err error
	db = db.Debug().Model(&MitraGroupProductLimit{}).Where("mitra_group_prod_lim_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_group_prod_lim_status": p.MitraGroupProductLimitStatus,
		},
	)
	err = db.Debug().Model(&MitraGroupProductLimit{}).Where("mitra_group_prod_lim_id = ?", pid).Error
	if err != nil {
		return &MitraGroupProductLimit{}, err
	}
	return p, nil
}

func (p *MitraGroupProductLimit) UpdateMitraGroupProductLimitStatusDelete(db *gorm.DB, pid uint64) (*MitraGroupProductLimit, error) {
	var err error
	db = db.Debug().Model(&MitraGroupProductLimit{}).Where("mitra_group_prod_lim_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_group_prod_lim_status":     3,
			"mitra_group_prod_lim_deleted_by": p.MitraGroupProductLimitDeletedBy,
			"mitra_group_prod_lim_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraGroupProductLimit{}).Where("mitra_group_prod_lim_id = ?", pid).Error
	if err != nil {
		return &MitraGroupProductLimit{}, err
	}
	return p, nil
}

func (p *MitraGroupProductLimit) DeleteMitraGroupProductLimit(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraGroupProductLimit{}).Where("mitra_group_prod_lim_id = ?", pid).Take(&MitraGroupProductLimit{}).Delete(&MitraGroupProductLimit{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraGroupProductLimit not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================
