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

type ProductCost struct {
	ProductCostID        uint64     `gorm:"column:prod_cost_id;primary_key;" json:"prod_cost_id"`
	ProductCostCode      string     `gorm:"column:prod_cost_code" json:"prod_cost_code"`
	ProductCostName      string     `gorm:"column:prod_cost_name" json:"prod_cost_name"`
	ProductCostDesc      string     `gorm:"column:prod_cost_desc" json:"prod_cost_desc"`
	ProductCostStatus    uint64     `gorm:"column:prod_cost_status;size:2" json:"prod_cost_status"`
	ProductCostCreatedBy string     `gorm:"column:prod_cost_created_by;size:125" json:"prod_cost_created_by"`
	ProductCostCreatedAt time.Time  `gorm:"column:prod_cost_created_at;default:CURRENT_TIMESTAMP" json:"prod_cost_created_at"`
	ProductCostUpdatedBy string     `gorm:"column:prod_cost_updated_by;size:125" json:"prod_cost_updated_by"`
	ProductCostUpdatedAt *time.Time `gorm:"column:prod_cost_updated_at;default:CURRENT_TIMESTAMP" json:"prod_cost_created_at"`
	ProductCostDeletedBy string     `gorm:"column:prod_cost_deleted_by;size:125" json:"prod_cost_deleted_by"`
	ProductCostDeletedAt *time.Time `gorm:"column:prod_cost_deleted_at;default:CURRENT_TIMESTAMP" json:"prod_cost_deleted_at"`
}

type ProductCostData struct {
	ProductCostID        uint64     `gorm:"column:prod_cost_id" json:"prod_cost_id"`
	ProductCostCode      string     `gorm:"column:prod_cost_code" json:"prod_cost_code"`
	ProductCostName      string     `gorm:"column:prod_cost_name" json:"prod_cost_name"`
	ProductCostDesc      string     `gorm:"column:prod_cost_desc" json:"prod_cost_desc"`
	ProductCostStatus    uint64     `gorm:"column:prod_cost_status" json:"prod_cost_status"`
	ProductCostCreatedBy string     `gorm:"column:prod_cost_created_by" json:"prod_cost_created_by"`
	ProductCostCreatedAt time.Time  `gorm:"column:prod_cost_created_at" json:"prod_cost_created_at"`
	ProductCostUpdatedBy string     `gorm:"column:prod_cost_updated_by" json:"prod_cost_updated_by"`
	ProductCostUpdatedAt *time.Time `gorm:"column:prod_cost_updated_at" json:"prod_cost_updated_at"`
	ProductCostDeletedBy string     `gorm:"column:prod_cost_deleted_by" json:"prod_cost_deleted_by"`
	ProductCostDeletedAt *time.Time `gorm:"column:prod_cost_deleted_at" json:"prod_cost_deleted_at"`
}

type ResponseProductCosts struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    *[]ProductCostData `json:"data"`
}

type ResponseProductCost struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    *ProductCostData `json:"data"`
}

type ResponseProductCostCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (ProductCost) TableName() string {
	return "m_product_cost"
}

func (ProductCostData) TableName() string {
	return "m_product_cost"
}

func (p *ProductCost) Prepare() {
	p.ProductCostCode = html.EscapeString(strings.TrimSpace(p.ProductCostCode))
	p.ProductCostName = html.EscapeString(strings.TrimSpace(p.ProductCostName))
	p.ProductCostDesc = p.ProductCostDesc
	p.ProductCostStatus = p.ProductCostStatus
	p.ProductCostCreatedBy = html.EscapeString(strings.TrimSpace(p.ProductCostCreatedBy))
	p.ProductCostCreatedAt = time.Now()
	p.ProductCostUpdatedBy = html.EscapeString(strings.TrimSpace(p.ProductCostUpdatedBy))
	p.ProductCostUpdatedAt = p.ProductCostUpdatedAt
	p.ProductCostDeletedBy = html.EscapeString(strings.TrimSpace(p.ProductCostDeletedBy))
	p.ProductCostDeletedAt = p.ProductCostDeletedAt
}

func (p *ProductCost) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.ProductCostCode == "" {
			return errors.New("Required Country")
		}
		if p.ProductCostName == "" {
			return errors.New("Required Province")
		}
		if p.ProductCostDesc == "" {
			return errors.New("Required Regency")
		}
		return nil
	}
}

func (p *ProductCost) SaveProductCost(db *gorm.DB) (*ProductCost, error) {
	var err error
	err = db.Debug().Model(&ProductCost{}).Create(&p).Error
	if err != nil {
		return &ProductCost{}, err
	}
	return p, nil
}

func (p *ProductCost) FindAllProductCost(db *gorm.DB) (*[]ProductCostData, error) {
	var err error
	product_cost := []ProductCostData{}
	err = db.Debug().Model(&ProductCostData{}).Limit(100).
		Select(`m_product_cost.prod_cost_id,
				m_product_cost.prod_cost_code,
				m_product_cost.prod_cost_name,
				m_product_cost.prod_cost_desc,
				m_product_cost.prod_cost_status,
				m_product_cost.prod_cost_created_by,
				m_product_cost.prod_cost_created_at at time zone 'Asia/Jakarta' as prod_cost_created_at,
				m_product_cost.prod_cost_updated_by,
				m_product_cost.prod_cost_updated_at at time zone 'Asia/Jakarta' as prod_cost_updated_at,
				m_product_cost.prod_cost_deleted_by,
				m_product_cost.prod_cost_deleted_at at time zone 'Asia/Jakarta' as prod_cost_deleted_at`).
		Find(&product_cost).Error
	if err != nil {
		return &[]ProductCostData{}, err
	}
	return &product_cost, nil
}

func (p *ProductCost) FindAllProductCostStatus(db *gorm.DB, status uint64) (*[]ProductCostData, error) {
	var err error
	product_cost := []ProductCostData{}
	err = db.Debug().Model(&ProductCostData{}).Limit(100).
		Select(`m_product_cost.prod_cost_id,
				m_product_cost.prod_cost_code,
				m_product_cost.prod_cost_name,
				m_product_cost.prod_cost_desc,
				m_product_cost.prod_cost_status,
				m_product_cost.prod_cost_created_by,
				m_product_cost.prod_cost_created_at at time zone 'Asia/Jakarta' as prod_cost_created_at,
				m_product_cost.prod_cost_updated_by,
				m_product_cost.prod_cost_updated_at at time zone 'Asia/Jakarta' as prod_cost_updated_at,
				m_product_cost.prod_cost_deleted_by,
				m_product_cost.prod_cost_deleted_at at time zone 'Asia/Jakarta' as prod_cost_deleted_at`).
		Where("prod_cost_status = ?", status).
		Find(&product_cost).Error
	if err != nil {
		return &[]ProductCostData{}, err
	}
	return &product_cost, nil
}

func (p *ProductCost) FindProductCostDataByID(db *gorm.DB, pid uint64) (*ProductCost, error) {
	var err error
	err = db.Debug().Model(&ProductCost{}).
		Select(`m_product_cost.prod_cost_id,
				m_product_cost.prod_cost_code,
				m_product_cost.prod_cost_name,
				m_product_cost.prod_cost_desc,
				m_product_cost.prod_cost_status,
				m_product_cost.prod_cost_created_by,
				m_product_cost.prod_cost_created_at at time zone 'Asia/Jakarta' as prod_cost_created_at,
				m_product_cost.prod_cost_updated_by,
				m_product_cost.prod_cost_updated_at at time zone 'Asia/Jakarta' as prod_cost_updated_at,
				m_product_cost.prod_cost_deleted_by,
				m_product_cost.prod_cost_deleted_at at time zone 'Asia/Jakarta' as prod_cost_deleted_at`).
		Where("prod_cost_id = ?", pid).Take(&p).Error
	if err != nil {
		return &ProductCost{}, err
	}
	return p, nil
}

func (p *ProductCost) FindProductCostByID(db *gorm.DB, pid uint64) (*ProductCostData, error) {
	var err error
	product_cost := ProductCostData{}
	err = db.Debug().Model(&ProductCostData{}).Limit(100).
		Select(`m_product_cost.prod_cost_id,
				m_product_cost.prod_cost_code,
				m_product_cost.prod_cost_name,
				m_product_cost.prod_cost_desc,
				m_product_cost.prod_cost_status,
				m_product_cost.prod_cost_created_by,
				m_product_cost.prod_cost_created_at at time zone 'Asia/Jakarta' as prod_cost_created_at,
				m_product_cost.prod_cost_updated_by,
				m_product_cost.prod_cost_updated_at at time zone 'Asia/Jakarta' as prod_cost_updated_at,
				m_product_cost.prod_cost_deleted_by,
				m_product_cost.prod_cost_deleted_at at time zone 'Asia/Jakarta' as prod_cost_deleted_at`).
		Where("prod_cost_id = ?", pid).
		Find(&product_cost).Error
	if err != nil {
		return &ProductCostData{}, err
	}
	return &product_cost, nil
}

func (p *ProductCost) FindProductCostStatusByID(db *gorm.DB, pid uint64, status uint64) (*ProductCostData, error) {
	var err error
	product_cost := ProductCostData{}
	err = db.Debug().Model(&ProductCostData{}).Limit(100).
		Select(`m_product_cost.prod_cost_id,
				m_product_cost.prod_cost_code,
				m_product_cost.prod_cost_name,
				m_product_cost.prod_cost_desc,
				m_product_cost.prod_cost_status,
				m_product_cost.prod_cost_created_by,
				m_product_cost.prod_cost_created_at at time zone 'Asia/Jakarta' as prod_cost_created_at,
				m_product_cost.prod_cost_updated_by,
				m_product_cost.prod_cost_updated_at at time zone 'Asia/Jakarta' as prod_cost_updated_at,
				m_product_cost.prod_cost_deleted_by,
				m_product_cost.prod_cost_deleted_at at time zone 'Asia/Jakarta' as prod_cost_deleted_at`).
		Where("prod_cost_id = ? AND prod_cost_status = ?", pid, status).
		Find(&product_cost).Error
	if err != nil {
		return &ProductCostData{}, err
	}
	return &product_cost, nil
}

func (p *ProductCost) UpdateProductCost(db *gorm.DB, pid uint64) (*ProductCost, error) {
	var err error
	db = db.Debug().Model(&ProductCost{}).Where("prod_cost_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_cost_code":       p.ProductCostCode,
			"prod_cost_name":       p.ProductCostName,
			"prod_cost_desc":       p.ProductCostDesc,
			"prod_cost_status":     p.ProductCostStatus,
			"prod_cost_updated_by": p.ProductCostUpdatedBy,
			"prod_cost_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ProductCost{}).Where("prod_cost_id = ?", pid).Error
	if err != nil {
		return &ProductCost{}, err
	}
	return p, nil
}

func (p *ProductCost) UpdateProductCostStatus(db *gorm.DB, pid uint64) (*ProductCost, error) {
	var err error
	db = db.Debug().Model(&ProductCost{}).Where("prod_cost_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_cost_status":     p.ProductCostStatus,
			"prod_cost_updated_by": p.ProductCostUpdatedBy,
			"prod_cost_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ProductCost{}).Where("prod_cost_id = ?", pid).Error
	if err != nil {
		return &ProductCost{}, err
	}
	return p, nil
}

func (p *ProductCost) UpdateProductCostStatusOnly(db *gorm.DB, pid uint64) (*ProductCost, error) {
	var err error
	db = db.Debug().Model(&ProductCost{}).Where("prod_cost_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_cost_status": p.ProductCostStatus,
		},
	)
	err = db.Debug().Model(&ProductCost{}).Where("prod_cost_id = ?", pid).Error
	if err != nil {
		return &ProductCost{}, err
	}
	return p, nil
}

func (p *ProductCost) UpdateProductCostStatusDelete(db *gorm.DB, pid uint64) (*ProductCost, error) {
	var err error
	db = db.Debug().Model(&ProductCost{}).Where("prod_cost_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_cost_status":     3,
			"prod_cost_deleted_by": p.ProductCostDeletedBy,
			"prod_cost_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ProductCost{}).Where("prod_cost_id = ?", pid).Error
	if err != nil {
		return &ProductCost{}, err
	}
	return p, nil
}

func (p *ProductCost) DeleteProductCost(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&ProductCost{}).Where("prod_cost_id = ?", pid).Take(&ProductCost{}).Delete(&ProductCost{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("ProductCost not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================
