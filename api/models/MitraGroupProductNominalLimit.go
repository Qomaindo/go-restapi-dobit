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

type MitraGroupProductNominalLimit struct {
	MitraGroupProductNominalLimitID        *uint64    `gorm:"column:mitra_group_prod_nom_lim_id;primary_key;" json:"mitra_group_prod_nom_lim_id"`
	MitraGroupProductLimitID               uint64     `gorm:"column:mitra_group_prod_lim_id" json:"mitra_group_prod_lim_id"`
	MitraProductID                         uint64     `gorm:"column:mitra_prod_id" json:"mitra_prod_id"`
	MitraGroupProductNominalLimitRole      string     `gorm:"column:mitra_group_prod_nom_lim_role" json:"mitra_group_prod_nom_lim_role"`
	MitraGroupProductNominalLimitMin       uint64     `gorm:"column:mitra_group_prod_nom_lim_min" json:"mitra_group_prod_nom_lim_min"`
	MitraGroupProductNominalLimitMax       uint64     `gorm:"column:mitra_group_prod_nom_lim_max" json:"mitra_group_prod_nom_lim_max"`
	MitraGroupProductNominalLimitStatus    uint64     `gorm:"column:mitra_group_prod_nom_lim_status;size:2" json:"mitra_group_prod_nom_lim_status"`
	MitraGroupProductNominalLimitCreatedBy string     `gorm:"column:mitra_group_prod_nom_lim_created_by;size:125" json:"mitra_group_prod_nom_lim_created_by"`
	MitraGroupProductNominalLimitCreatedAt time.Time  `gorm:"column:mitra_group_prod_nom_lim_created_at;default:CURRENT_TIMESTAMP" json:"mitra_group_prod_nom_lim_created_at"`
	MitraGroupProductNominalLimitUpdatedBy string     `gorm:"column:mitra_group_prod_nom_lim_updated_by;size:125" json:"mitra_group_prod_nom_lim_updated_by"`
	MitraGroupProductNominalLimitUpdatedAt *time.Time `gorm:"column:mitra_group_prod_nom_lim_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_group_prod_nom_lim_created_at"`
	MitraGroupProductNominalLimitDeletedBy string     `gorm:"column:mitra_group_prod_nom_lim_deleted_by;size:125" json:"mitra_group_prod_nom_lim_deleted_by"`
	MitraGroupProductNominalLimitDeletedAt *time.Time `gorm:"column:mitra_group_prod_nom_lim_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_group_prod_nom_lim_deleted_at"`
}

type MitraGroupProductNominalLimitData struct {
	MitraGroupProductNominalLimitID        uint64    `gorm:"column:mitra_group_prod_nom_lim_id" json:"mitra_group_prod_nom_lim_id"`
	MitraGroupProductLimitID               uint64    `gorm:"column:mitra_group_prod_lim_id" json:"mitra_group_prod_lim_id"`
	MitraGroupProductLimitName             string    `gorm:"column:mitra_group_prod_lim_name" json:"mitra_group_prod_lim_name"`
	MitraProductID                         uint64    `gorm:"column:mitra_prod_id" json:"mitra_prod_id"`
	MitraProductName                       string    `gorm:"column:mitra_prod_name" json:"mitra_prod_name"`
	MitraGroupProductNominalLimitRole      string    `gorm:"column:mitra_group_prod_nom_lim_role" json:"mitra_group_prod_nom_lim_role"`
	MitraGroupProductNominalLimitMin       uint64    `gorm:"column:mitra_group_prod_nom_lim_min" json:"mitra_group_prod_nom_lim_min"`
	MitraGroupProductNominalLimitMax       uint64    `gorm:"column:mitra_group_prod_nom_lim_max" json:"mitra_group_prod_nom_lim_max"`
	MitraGroupProductNominalLimitStatus    uint64    `gorm:"column:mitra_group_prod_nom_lim_status" json:"mitra_group_prod_nom_lim_status"`
	MitraGroupProductNominalLimitCreatedBy string    `gorm:"column:mitra_group_prod_nom_lim_created_by" json:"mitra_group_prod_nom_lim_created_by"`
	MitraGroupProductNominalLimitCreatedAt time.Time `gorm:"column:mitra_group_prod_nom_lim_created_at" json:"mitra_group_prod_nom_lim_created_at"`
}

type ResponseMitraGroupProductNominalLimits struct {
	Status  int                                  `json:"status"`
	Message string                               `json:"message"`
	Data    *[]MitraGroupProductNominalLimitData `json:"data"`
}

type ResponseMitraGroupProductNominalLimit struct {
	Status  int                                `json:"status"`
	Message string                             `json:"message"`
	Data    *MitraGroupProductNominalLimitData `json:"data"`
}

type ResponseMitraGroupProductNominalLimitCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraGroupProductNominalLimit) TableName() string {
	return "m_mitra_group_product_nominal_limit"
}

func (MitraGroupProductNominalLimitData) TableName() string {
	return "m_mitra_group_product_nominal_limit"
}

func (p *MitraGroupProductNominalLimit) Prepare() {
	p.MitraGroupProductNominalLimitID = nil
	p.MitraGroupProductLimitID = p.MitraGroupProductLimitID
	p.MitraProductID = p.MitraProductID
	p.MitraGroupProductNominalLimitRole = p.MitraGroupProductNominalLimitRole
	p.MitraGroupProductNominalLimitMin = p.MitraGroupProductNominalLimitMin
	p.MitraGroupProductNominalLimitMax = p.MitraGroupProductNominalLimitMax
	p.MitraGroupProductNominalLimitStatus = p.MitraGroupProductNominalLimitStatus
	p.MitraGroupProductNominalLimitCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraGroupProductNominalLimitCreatedBy))
	p.MitraGroupProductNominalLimitCreatedAt = time.Now()
	p.MitraGroupProductNominalLimitUpdatedBy = html.EscapeString(strings.TrimSpace(p.MitraGroupProductNominalLimitUpdatedBy))
	p.MitraGroupProductNominalLimitUpdatedAt = p.MitraGroupProductNominalLimitUpdatedAt
	p.MitraGroupProductNominalLimitDeletedBy = html.EscapeString(strings.TrimSpace(p.MitraGroupProductNominalLimitDeletedBy))
	p.MitraGroupProductNominalLimitDeletedAt = p.MitraGroupProductNominalLimitDeletedAt
}

func (p *MitraGroupProductNominalLimit) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.MitraGroupProductLimitID == 0 {
			return errors.New("Required Country")
		}
		if p.MitraProductID == 0 {
			return errors.New("Required Province")
		}
		return nil
	}
}

func (p *MitraGroupProductNominalLimit) SaveMitraGroupProductNominalLimit(db *gorm.DB) (*MitraGroupProductNominalLimit, error) {
	var err error
	err = db.Debug().Model(&MitraGroupProductNominalLimit{}).Create(&p).Error
	if err != nil {
		return &MitraGroupProductNominalLimit{}, err
	}
	return p, nil
}

func (p *MitraGroupProductNominalLimit) FindAllMitraGroupProductNominalLimit(db *gorm.DB) (*[]MitraGroupProductNominalLimitData, error) {
	var err error
	mitragroupproductnominallimit := []MitraGroupProductNominalLimitData{}
	err = db.Debug().Model(&MitraGroupProductNominalLimitData{}).Limit(100).
		Select(`m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_id,
				m_mitra_group_product_limit.mitra_group_prod_lim_id,
				m_mitra_group_product_limit.mitra_group_prod_lim_name,
				m_mitra_product.mitra_prod_id,
				m_mitra_product.mitra_prod_name,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_role,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_min,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_max,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_status,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_created_by,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_created_at,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_updated_by,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_updated_at,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_deleted_by,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_deleted_at`).
		Joins("JOIN m_mitra_group_product_limit on m_mitra_group_product_nominal_limit.mitra_group_prod_lim_id=m_mitra_group_product_limit.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_product on m_mitra_group_product_nominal_limit.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Find(&mitragroupproductnominallimit).Error
	if err != nil {
		return &[]MitraGroupProductNominalLimitData{}, err
	}
	return &mitragroupproductnominallimit, nil
}

func (p *MitraGroupProductNominalLimit) FindAllMitraGroupProductNominalLimitStatus(db *gorm.DB, status uint64) (*[]MitraGroupProductNominalLimitData, error) {
	var err error
	mitragroupproductnominallimit := []MitraGroupProductNominalLimitData{}
	err = db.Debug().Model(&MitraGroupProductNominalLimitData{}).Limit(100).
		Select(`m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_id,
				m_mitra_group_product_limit.mitra_group_prod_lim_id,
				m_mitra_group_product_limit.mitra_group_prod_lim_name,
				m_mitra_product.mitra_prod_id,
				m_mitra_product.mitra_prod_name,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_role,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_min,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_max,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_status,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_created_by,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_created_at,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_updated_by,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_updated_at,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_deleted_by,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_deleted_at`).
		Joins("JOIN m_mitra_group_product_limit on m_mitra_group_product_nominal_limit.mitra_group_prod_lim_id=m_mitra_group_product_limit.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_product on m_mitra_group_product_nominal_limit.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Where("mitra_group_prod_nom_lim_status = ?", status).
		Find(&mitragroupproductnominallimit).Error
	if err != nil {
		return &[]MitraGroupProductNominalLimitData{}, err
	}
	return &mitragroupproductnominallimit, nil
}

func (p *MitraGroupProductNominalLimit) FindMitraGroupProductNominalLimitDataByID(db *gorm.DB, pid uint64) (*MitraGroupProductNominalLimit, error) {
	var err error
	err = db.Debug().Model(&MitraGroupProductNominalLimit{}).Where("mitra_group_prod_nom_lim_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraGroupProductNominalLimit{}, err
	}
	return p, nil
}

func (p *MitraGroupProductNominalLimit) FindMitraGroupProductNominalLimitByID(db *gorm.DB, pid uint64) (*MitraGroupProductNominalLimitData, error) {
	var err error
	mitragroupproductnominallimit := MitraGroupProductNominalLimitData{}
	err = db.Debug().Model(&MitraGroupProductNominalLimitData{}).Limit(100).
		Select(`m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_id,
				m_mitra_group_product_limit.mitra_group_prod_lim_id,
				m_mitra_group_product_limit.mitra_group_prod_lim_name,
				m_mitra_product.mitra_prod_id,
				m_mitra_product.mitra_prod_name,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_role,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_min,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_max,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_status,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_created_by,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_created_at,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_updated_by,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_updated_at,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_deleted_by,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_deleted_at`).
		Joins("JOIN m_mitra_group_product_limit on m_mitra_group_product_nominal_limit.mitra_group_prod_lim_id=m_mitra_group_product_limit.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_product on m_mitra_group_product_nominal_limit.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Where("mitra_group_prod_nom_lim_id = ?", pid).
		Find(&mitragroupproductnominallimit).Error
	if err != nil {
		return &MitraGroupProductNominalLimitData{}, err
	}
	return &mitragroupproductnominallimit, nil
}

func (p *MitraGroupProductNominalLimit) FindMitraGroupProductNominalLimitStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraGroupProductNominalLimitData, error) {
	var err error
	mitragroupproductnominallimit := MitraGroupProductNominalLimitData{}
	err = db.Debug().Model(&MitraGroupProductNominalLimitData{}).Limit(100).
		Select(`m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_id,
				m_mitra_group_product_limit.mitra_group_prod_lim_id,
				m_mitra_group_product_limit.mitra_group_prod_lim_name,
				m_mitra_product.mitra_prod_id,
				m_mitra_product.mitra_prod_name,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_role,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_min,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_max,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_status,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_created_by,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_created_at,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_updated_by,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_updated_at,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_deleted_by,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_deleted_at`).
		Joins("JOIN m_mitra_group_product_limit on m_mitra_group_product_nominal_limit.mitra_group_prod_lim_id=m_mitra_group_product_limit.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_product on m_mitra_group_product_nominal_limit.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Where("mitra_group_prod_nom_lim_id = ? AND mitra_group_prod_nom_lim_status = ?", pid, status).
		Find(&mitragroupproductnominallimit).Error
	if err != nil {
		return &MitraGroupProductNominalLimitData{}, err
	}
	return &mitragroupproductnominallimit, nil
}

func (p *MitraGroupProductNominalLimit) UpdateMitraGroupProductNominalLimit(db *gorm.DB, pid uint64) (*MitraGroupProductNominalLimit, error) {
	var err error
	db = db.Debug().Model(&MitraGroupProductNominalLimit{}).Where("mitra_group_prod_nom_lim_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_group_prod_lim_id":             p.MitraGroupProductLimitID,
			"mitra_prod_id":                       p.MitraProductID,
			"mitra_group_prod_nom_lim_role":       p.MitraGroupProductNominalLimitRole,
			"mitra_group_prod_nom_lim_min":        p.MitraGroupProductNominalLimitMin,
			"mitra_group_prod_nom_lim_max":        p.MitraGroupProductNominalLimitMax,
			"mitra_group_prod_nom_lim_status":     p.MitraGroupProductNominalLimitStatus,
			"mitra_group_prod_nom_lim_updated_by": p.MitraGroupProductNominalLimitUpdatedBy,
			"mitra_group_prod_nom_lim_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraGroupProductNominalLimit{}).Where("mitra_group_prod_nom_lim_id = ?", pid).Error
	if err != nil {
		return &MitraGroupProductNominalLimit{}, err
	}
	return p, nil
}

func (p *MitraGroupProductNominalLimit) UpdateMitraGroupProductNominalLimitStatus(db *gorm.DB, pid uint64) (*MitraGroupProductNominalLimit, error) {
	var err error
	db = db.Debug().Model(&MitraGroupProductNominalLimit{}).Where("mitra_group_prod_nom_lim_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_group_prod_nom_lim_status":     p.MitraGroupProductNominalLimitStatus,
			"mitra_group_prod_nom_lim_updated_by": p.MitraGroupProductNominalLimitUpdatedBy,
			"mitra_group_prod_nom_lim_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraGroupProductNominalLimit{}).Where("mitra_group_prod_nom_lim_id = ?", pid).Error
	if err != nil {
		return &MitraGroupProductNominalLimit{}, err
	}
	return p, nil
}

func (p *MitraGroupProductNominalLimit) UpdateMitraGroupProductNominalLimitStatusOnly(db *gorm.DB, pid uint64) (*MitraGroupProductNominalLimit, error) {
	var err error
	db = db.Debug().Model(&MitraGroupProductNominalLimit{}).Where("mitra_group_prod_nom_lim_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_group_prod_nom_lim_status": p.MitraGroupProductNominalLimitStatus,
		},
	)
	err = db.Debug().Model(&MitraGroupProductNominalLimit{}).Where("mitra_group_prod_nom_lim_id = ?", pid).Error
	if err != nil {
		return &MitraGroupProductNominalLimit{}, err
	}
	return p, nil
}

func (p *MitraGroupProductNominalLimit) UpdateMitraGroupProductNominalLimitStatusDelete(db *gorm.DB, pid uint64) (*MitraGroupProductNominalLimit, error) {
	var err error
	db = db.Debug().Model(&MitraGroupProductNominalLimit{}).Where("mitra_group_prod_nom_lim_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_group_prod_nom_lim_status":     3,
			"mitra_group_prod_nom_lim_deleted_by": p.MitraGroupProductNominalLimitDeletedBy,
			"mitra_group_prod_nom_lim_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraGroupProductNominalLimit{}).Where("mitra_group_prod_nom_lim_id = ?", pid).Error
	if err != nil {
		return &MitraGroupProductNominalLimit{}, err
	}
	return p, nil
}

func (p *MitraGroupProductNominalLimit) DeleteMitraGroupProductNominalLimit(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraGroupProductNominalLimit{}).Where("mitra_group_prod_nom_lim_id = ?", pid).Take(&MitraGroupProductNominalLimit{}).Delete(&MitraGroupProductNominalLimit{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraGroupProductNominalLimit not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

func (p *MitraGroupProductNominalLimit) FindMitraGroupProductNominalLimitByMitraGroupProductLimitID(db *gorm.DB, pid uint64) ([]MitraGroupProductNominalLimitData, error) {
	var err error
	mitragroupproductnominallimit := []MitraGroupProductNominalLimitData{}
	err = db.Debug().Model(&MitraGroupProductNominalLimitData{}).Limit(100).
		Select(`m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_id,
				m_mitra_group_product_limit.mitra_group_prod_lim_id,
				m_mitra_group_product_limit.mitra_group_prod_lim_name,
				m_mitra_product.mitra_prod_id,
				m_mitra_product.mitra_prod_name,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_role,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_min,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_max,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_status,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_created_by,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_created_at,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_updated_by,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_updated_at,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_deleted_by,
				m_mitra_group_product_nominal_limit.mitra_group_prod_nom_lim_deleted_at`).
		Joins("JOIN m_mitra_group_product_limit on m_mitra_group_product_nominal_limit.mitra_group_prod_lim_id=m_mitra_group_product_limit.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_product on m_mitra_group_product_nominal_limit.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Where("m_mitra_group_product_nominal_limit.mitra_group_prod_lim_id = ?", pid).
		Find(&mitragroupproductnominallimit).Error
	if err != nil {
		return []MitraGroupProductNominalLimitData{}, err
	}
	return mitragroupproductnominallimit, nil
}
