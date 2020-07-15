package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ProductSubCategory struct {
	ProductSubCategoryID        uint64     `gorm:"column:prod_sub_ctgry_id;primary_key;" json:"prod_sub_ctgry_id"`
	ProductCategoryID           uint64     `gorm:"column:prod_ctgry_id" json:"prod_ctgry_id"`
	ProductSubCategoryCode      string     `gorm:"column:prod_sub_ctgry_code;size:25" json:"prod_sub_ctgry_code"`
	ProductSubCategoryName      string     `gorm:"column:prod_sub_ctgry_name;size:255" json:"prod_sub_ctgry_name"`
	ProductSubCategoryInitial   string     `gorm:"column:prod_sub_ctgry_initial;size:1" json:"prod_sub_ctgry_initial"`
	ProductSubCategoryStatus    uint64     `gorm:"column:prod_sub_ctgry_status;size:2" json:"prod_sub_ctgry_status"`
	ProductSubCategoryCreatedBy string     `gorm:"column:prod_sub_ctgry_created_by;size:125" json:"prod_sub_ctgry_created_by"`
	ProductSubCategoryUpdatedBy string     `gorm:"column:prod_sub_ctgry_updated_by;size:125" json:"prod_sub_ctgry_updated_by"`
	ProductSubCategoryDeletedBy string     `gorm:"column:prod_sub_ctgry_deleted_by;size:125" json:"prod_sub_ctgry_deleted_by"`
	ProductSubCategoryCreatedAt time.Time  `gorm:"column:prod_sub_ctgry_created_at;default:CURRENT_TIMESTAMP" json:"prod_sub_ctgry_created_at"`
	ProductSubCategoryUpdatedAt *time.Time `gorm:"column:prod_sub_ctgry_updated_at;default:CURRENT_TIMESTAMP" json:"prod_sub_ctgry_updated_at"`
	ProductSubCategoryDeletedAt *time.Time `gorm:"column:prod_sub_ctgry_deleted_at;default:CURRENT_TIMESTAMP" json:"prod_sub_ctgry_deleted_at"`
}

type ProductSubCategoryData struct {
	ProductSubCategoryID        uint64     `gorm:"column:prod_sub_ctgry_id" json:"prod_sub_ctgry_id"`
	ProductCategoryID           uint64     `gorm:"column:prod_ctgry_id" json:"prod_ctgry_id"`
	ProductCategoryCode         string     `gorm:"column:prod_ctgry_code;size:25" json:"prod_ctgry_code"`
	ProductCategoryName         string     `gorm:"column:prod_ctgry_name;size:255" json:"prod_ctgry_name"`
	ProductSubCategoryCode      string     `gorm:"column:prod_sub_ctgry_code;size:15" json:"prod_sub_ctgry_code"`
	ProductSubCategoryName      string     `gorm:"column:prod_sub_ctgry_name;size:255" json:"prod_sub_ctgry_name"`
	ProductSubCategoryInitial   string     `gorm:"column:prod_sub_ctgry_initial;size:1" json:"prod_sub_ctgry_initial"`
	ProductSubCategoryStatus    uint64     `gorm:"column:prod_sub_ctgry_status;size:2" json:"prod_sub_ctgry_status"`
	ProductSubCategoryCreatedBy string     `gorm:"column:prod_sub_ctgry_created_by;size:125" json:"prod_sub_ctgry_created_by"`
	ProductSubCategoryUpdatedBy string     `gorm:"column:prod_sub_ctgry_updated_by;size:125" json:"prod_sub_ctgry_updated_by"`
	ProductSubCategoryDeletedBy string     `gorm:"column:prod_sub_ctgry_deleted_by;size:125" json:"prod_sub_ctgry_deleted_by"`
	ProductSubCategoryCreatedAt time.Time  `gorm:"column:prod_sub_ctgry_created_at;default:CURRENT_TIMESTAMP" json:"prod_sub_ctgry_created_at"`
	ProductSubCategoryUpdatedAt *time.Time `gorm:"column:prod_sub_ctgry_updated_at;default:CURRENT_TIMESTAMP" json:"prod_sub_ctgry_updated_at"`
	ProductSubCategoryDeletedAt *time.Time `gorm:"column:prod_sub_ctgry_deleted_at;default:CURRENT_TIMESTAMP" json:"prod_sub_ctgry_deleted_at"`
}

type ResponseProductSubCategorys struct {
	Status  int                       `json:"status"`
	Message string                    `json:"message"`
	Data    *[]ProductSubCategoryData `json:"data"`
}

type ResponseProductSubCategory struct {
	Status  int                     `json:"status"`
	Message string                  `json:"message"`
	Data    *ProductSubCategoryData `json:"data"`
}

func (ProductSubCategory) TableName() string {
	return "m_product_sub_category"
}

func (ProductSubCategoryData) TableName() string {
	return "m_product_sub_category"
}

func (p *ProductSubCategory) FindAllProductSubCategory(db *gorm.DB) (*[]ProductSubCategoryData, error) {
	var err error
	productsubcategory := []ProductSubCategoryData{}
	err = db.Debug().Model(&ProductSubCategoryData{}).Limit(100).
		Select(`m_product_sub_category.prod_sub_ctgry_id,
				m_product_category.prod_ctgry_id,
				m_product_category.prod_ctgry_code,
				m_product_category.prod_ctgry_name,
				m_product_category.prod_ctgry_status,
				m_product_sub_category.prod_sub_ctgry_code,
				m_product_sub_category.prod_sub_ctgry_name,
				m_product_sub_category.prod_sub_ctgry_initial,
				m_product_sub_category.prod_sub_ctgry_status,
				m_product_sub_category.prod_sub_ctgry_created_by,
				m_product_sub_category.prod_sub_ctgry_updated_by,
				m_product_sub_category.prod_sub_ctgry_deleted_by,
				m_product_sub_category.prod_sub_ctgry_created_at,
				m_product_sub_category.prod_sub_ctgry_updated_at,
				m_product_sub_category.prod_sub_ctgry_deleted_at`).
		Joins("JOIN m_product_category on m_product_sub_category.prod_ctgry_id=m_product_category.prod_ctgry_id").
		Find(&productsubcategory).Error
	if err != nil {
		return &[]ProductSubCategoryData{}, err
	}
	return &productsubcategory, nil
}

func (p *ProductSubCategory) FindAllProductSubCategoryStatus(db *gorm.DB, status uint64) (*[]ProductSubCategoryData, error) {
	var err error
	productsubcategory := []ProductSubCategoryData{}
	err = db.Debug().Model(&ProductSubCategoryData{}).Limit(100).
		Select(`m_product_sub_category.prod_sub_ctgry_id,
				m_product_category.prod_ctgry_id,
				m_product_category.prod_ctgry_code,
				m_product_category.prod_ctgry_name,
				m_product_category.prod_ctgry_status,
				m_product_sub_category.prod_sub_ctgry_code,
				m_product_sub_category.prod_sub_ctgry_name,
				m_product_sub_category.prod_sub_ctgry_initial,
				m_product_sub_category.prod_sub_ctgry_status,
				m_product_sub_category.prod_sub_ctgry_created_by,
				m_product_sub_category.prod_sub_ctgry_updated_by,
				m_product_sub_category.prod_sub_ctgry_deleted_by,
				m_product_sub_category.prod_sub_ctgry_created_at,
				m_product_sub_category.prod_sub_ctgry_updated_at,
				m_product_sub_category.prod_sub_ctgry_deleted_at`).
		Where("prod_sub_ctgry_status = ?", status).
		Joins("JOIN m_product_category on m_product_sub_category.prod_ctgry_id=m_product_category.prod_ctgry_id").
		Find(&productsubcategory).Error
	if err != nil {
		return &[]ProductSubCategoryData{}, err
	}
	return &productsubcategory, nil
}

func (p *ProductSubCategory) FindProductSubCategoryByID(db *gorm.DB, pid uint64) (*ProductSubCategoryData, error) {
	var err error
	productsubcategory := ProductSubCategoryData{}
	err = db.Debug().Model(&ProductSubCategoryData{}).Limit(100).
		Select(`m_product_sub_category.prod_sub_ctgry_id,
				m_product_category.prod_ctgry_id,
				m_product_category.prod_ctgry_code,
				m_product_category.prod_ctgry_name,
				m_product_category.prod_ctgry_status,
				m_product_sub_category.prod_sub_ctgry_code,
				m_product_sub_category.prod_sub_ctgry_name,
				m_product_sub_category.prod_sub_ctgry_initial,
				m_product_sub_category.prod_sub_ctgry_status,
				m_product_sub_category.prod_sub_ctgry_created_by,
				m_product_sub_category.prod_sub_ctgry_updated_by,
				m_product_sub_category.prod_sub_ctgry_deleted_by,
				m_product_sub_category.prod_sub_ctgry_created_at,
				m_product_sub_category.prod_sub_ctgry_updated_at,
				m_product_sub_category.prod_sub_ctgry_deleted_at`).
		Where("prod_sub_ctgry_id = ?", pid).
		Joins("JOIN m_product_category on m_product_sub_category.prod_ctgry_id=m_product_category.prod_ctgry_id").
		Find(&productsubcategory).Error
	if err != nil {
		return &ProductSubCategoryData{}, err
	}
	return &productsubcategory, nil
}

func (p *ProductSubCategory) FindProductSubCategoryStatusByID(db *gorm.DB, pid uint64, status uint64) (*ProductSubCategoryData, error) {
	var err error
	productsubcategory := ProductSubCategoryData{}
	err = db.Debug().Model(&ProductSubCategoryData{}).Limit(100).
		Select(`m_product_sub_category.prod_sub_ctgry_id,
				m_product_category.prod_ctgry_id,
				m_product_category.prod_ctgry_code,
				m_product_category.prod_ctgry_name,
				m_product_category.prod_ctgry_status,
				m_product_sub_category.prod_sub_ctgry_code,
				m_product_sub_category.prod_sub_ctgry_name,
				m_product_sub_category.prod_sub_ctgry_initial,
				m_product_sub_category.prod_sub_ctgry_status,
				m_product_sub_category.prod_sub_ctgry_created_by,
				m_product_sub_category.prod_sub_ctgry_updated_by,
				m_product_sub_category.prod_sub_ctgry_deleted_by,
				m_product_sub_category.prod_sub_ctgry_created_at,
				m_product_sub_category.prod_sub_ctgry_updated_at,
				m_product_sub_category.prod_sub_ctgry_deleted_at`).
		Where("prod_sub_ctgry_id = ? AND prod_sub_ctgry_status = ?", pid, status).
		Joins("JOIN m_product_category on m_product_sub_category.prod_ctgry_id=m_product_category.prod_ctgry_id").
		Find(&productsubcategory).Error
	if err != nil {
		return &ProductSubCategoryData{}, err
	}
	return &productsubcategory, nil
}

//ADDITIONAL
// ============================================================================================================================

func (p *ProductSubCategory) FindAllProductSubCategoryOnly(db *gorm.DB) (*[]ProductSubCategoryData, error) {
	var err error
	productsubcategory := []ProductSubCategoryData{}
	err = db.Debug().Model(&ProductSubCategoryData{}).Limit(100).
		Select(`m_product_sub_category.prod_sub_ctgry_id,
				m_product_category.prod_ctgry_id,
				m_product_category.prod_ctgry_code,
				m_product_category.prod_ctgry_name,
				m_product_category.prod_ctgry_status,
				m_product_sub_category.prod_sub_ctgry_code,
				m_product_sub_category.prod_sub_ctgry_name,
				m_product_sub_category.prod_sub_ctgry_initial,
				m_product_sub_category.prod_sub_ctgry_status,
				m_product_sub_category.prod_sub_ctgry_created_by,
				m_product_sub_category.prod_sub_ctgry_updated_by,
				m_product_sub_category.prod_sub_ctgry_deleted_by,
				m_product_sub_category.prod_sub_ctgry_created_at,
				m_product_sub_category.prod_sub_ctgry_updated_at,
				m_product_sub_category.prod_sub_ctgry_deleted_at`).
		Where("m_product_sub_category.prod_sub_ctgry_initial = ? OR m_product_sub_category.prod_sub_ctgry_initial = ?", "A", "B").
		Joins("JOIN m_product_category on m_product_sub_category.prod_ctgry_id=m_product_category.prod_ctgry_id").
		Find(&productsubcategory).Error
	if err != nil {
		return &[]ProductSubCategoryData{}, err
	}
	return &productsubcategory, nil
}

func (p *ProductSubCategory) FindAllProductSubCategoryByCategory(db *gorm.DB, category uint64) (*[]ProductSubCategoryData, error) {
	var err error
	productsubcategory := []ProductSubCategoryData{}
	err = db.Debug().Model(&ProductSubCategoryData{}).Limit(100).
		Select(`m_product_sub_category.prod_sub_ctgry_id,
				m_product_category.prod_ctgry_id,
				m_product_category.prod_ctgry_code,
				m_product_category.prod_ctgry_name,
				m_product_category.prod_ctgry_status,
				m_product_sub_category.prod_sub_ctgry_code,
				m_product_sub_category.prod_sub_ctgry_name,
				m_product_sub_category.prod_sub_ctgry_initial,
				m_product_sub_category.prod_sub_ctgry_status,
				m_product_sub_category.prod_sub_ctgry_created_by,
				m_product_sub_category.prod_sub_ctgry_updated_by,
				m_product_sub_category.prod_sub_ctgry_deleted_by,
				m_product_sub_category.prod_sub_ctgry_created_at,
				m_product_sub_category.prod_sub_ctgry_updated_at,
				m_product_sub_category.prod_sub_ctgry_deleted_at`).
		Where("m_product_sub_category.prod_sub_ctgry_initial = ? OR m_product_sub_category.prod_sub_ctgry_initial = ?", "A", "B").
		Where("m_product_sub_category.prod_ctgry_id = ?", category).
		Joins("JOIN m_product_category on m_product_sub_category.prod_ctgry_id=m_product_category.prod_ctgry_id").
		Find(&productsubcategory).Error
	if err != nil {
		return &[]ProductSubCategoryData{}, err
	}
	return &productsubcategory, nil
}
