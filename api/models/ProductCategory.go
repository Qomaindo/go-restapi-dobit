package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ProductCategory struct {
	ProductCategoryID        uint64     `gorm:"column:prod_ctgry_id;primary_key;" json:"prod_ctgry_id"`
	ProductCategoryCode      string     `gorm:"column:prod_ctgry_code;size:25" json:"prod_ctgry_code"`
	ProductCategoryName      string     `gorm:"column:prod_ctgry_name;size:255" json:"prod_ctgry_name"`
	ProductCategoryStatus    uint64     `gorm:"column:prod_ctgry_status;size:2" json:"prod_ctgry_status"`
	ProductCategoryCreatedBy string     `gorm:"column:prod_ctgry_created_by;size:125" json:"prod_ctgry_created_by"`
	ProductCategoryUpdatedBy string     `gorm:"column:prod_ctgry_updated_by;size:125" json:"prod_ctgry_updated_by"`
	ProductCategoryDeletedBy string     `gorm:"column:prord_ctgry_deleted_by;size:125" json:"prord_ctgry_deleted_by"`
	ProductCategoryCreatedAt time.Time  `gorm:"column:prod_ctgry_created_at;default:CURRENT_TIMESTAMP" json:"prod_ctgry_created_at"`
	ProductCategoryUpdatedAt *time.Time `gorm:"column:prod_ctgry_updated_at;default:CURRENT_TIMESTAMP" json:"prod_ctgry_updated_at"`
	ProductCategoryDeletedAt *time.Time `gorm:"column:prod_ctgry_deleted_at;default:CURRENT_TIMESTAMP" json:"prod_ctgry_deleted_at"`
}

type ResponseProductCategorys struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    *[]ProductCategory `json:"data"`
}

type ResponseProductCategory struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    *ProductCategory `json:"data"`
}

func (ProductCategory) TableName() string {
	return "m_product_category"
}

func (p *ProductCategory) FindAllProductCategory(db *gorm.DB) (*[]ProductCategory, error) {
	var err error
	productcategory := []ProductCategory{}
	err = db.Debug().Model(&ProductCategory{}).Limit(100).Find(&productcategory).Error
	if err != nil {
		return &[]ProductCategory{}, err
	}
	return &productcategory, nil
}

func (p *ProductCategory) FindAllProductCategoryStatus(db *gorm.DB, status uint64) (*[]ProductCategory, error) {
	var err error
	productcategory := []ProductCategory{}
	err = db.Debug().Model(&ProductCategory{}).Where("prod_ctgry_status = ?", status).Limit(100).Find(&productcategory).Error
	if err != nil {
		return &[]ProductCategory{}, err
	}
	return &productcategory, nil
}

func (p *ProductCategory) FindProductCategoryByID(db *gorm.DB, pid uint64) (*ProductCategory, error) {
	var err error
	err = db.Debug().Model(&ProductCategory{}).Where("prod_ctgry_id = ?", pid).Take(&p).Error
	if err != nil {
		return &ProductCategory{}, err
	}
	return p, nil
}

func (p *ProductCategory) FindProductCategoryStatusByID(db *gorm.DB, pid uint64, status uint64) (*ProductCategory, error) {
	var err error
	err = db.Debug().Model(&ProductCategory{}).Where("prod_ctgry_id = ? AND prod_ctgry_status = ?", pid, status).Take(&p).Error
	if err != nil {
		return &ProductCategory{}, err
	}
	return p, nil
}
