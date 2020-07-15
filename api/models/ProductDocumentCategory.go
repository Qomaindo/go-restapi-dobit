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

type ProductDocumentCategory struct {
	ProductDocumentCategoryID        uint64     `gorm:"column:prod_doc_ctgry_id;primary_key;" json:"prod_doc_ctgry_id"`
	ProductDocumentCategoryCode      string     `gorm:"column:prod_doc_ctgry_code" json:"prod_doc_ctgry_code"`
	ProductDocumentCategoryName      string     `gorm:"column:prod_doc_ctgry_name" json:"prod_doc_ctgry_name"`
	ProductDocumentCategoryDesc      string     `gorm:"column:prod_doc_ctgry_desc" json:"prod_doc_ctgry_desc"`
	ProductDocumentCategoryStatus    uint64     `gorm:"column:prod_doc_ctgry_status;size:2" json:"prod_doc_ctgry_status"`
	ProductDocumentCategoryCreatedBy string     `gorm:"column:prod_doc_ctgry_created_by;size:125" json:"prod_doc_ctgry_created_by"`
	ProductDocumentCategoryCreatedAt time.Time  `gorm:"column:prod_doc_ctgry_created_at;default:CURRENT_TIMESTAMP" json:"prod_doc_ctgry_created_at"`
	ProductDocumentCategoryUpdatedBy string     `gorm:"column:prod_doc_ctgry_updated_by;size:125" json:"prod_doc_ctgry_updated_by"`
	ProductDocumentCategoryUpdatedAt *time.Time `gorm:"column:prod_doc_ctgry_updated_at;default:CURRENT_TIMESTAMP" json:"prod_doc_ctgry_created_at"`
	ProductDocumentCategoryDeletedBy string     `gorm:"column:prod_doc_ctgry_deleted_by;size:125" json:"prod_doc_ctgry_deleted_by"`
	ProductDocumentCategoryDeletedAt *time.Time `gorm:"column:prod_doc_ctgry_deleted_at;default:CURRENT_TIMESTAMP" json:"prod_doc_ctgry_deleted_at"`
}

type ProductDocumentCategoryData struct {
	ProductDocumentCategoryID        uint64     `gorm:"column:prod_doc_ctgry_id" json:"prod_doc_ctgry_id"`
	ProductDocumentCategoryCode      string     `gorm:"column:prod_doc_ctgry_code" json:"prod_doc_ctgry_code"`
	ProductDocumentCategoryName      string     `gorm:"column:prod_doc_ctgry_name" json:"prod_doc_ctgry_name"`
	ProductDocumentCategoryDesc      string     `gorm:"column:prod_doc_ctgry_desc" json:"prod_doc_ctgry_desc"`
	ProductDocumentCategoryStatus    uint64     `gorm:"column:prod_doc_ctgry_status" json:"prod_doc_ctgry_status"`
	ProductDocumentCategoryCreatedBy string     `gorm:"column:prod_doc_ctgry_created_by" json:"prod_doc_ctgry_created_by"`
	ProductDocumentCategoryCreatedAt time.Time  `gorm:"column:prod_doc_ctgry_created_at" json:"prod_doc_ctgry_created_at"`
	ProductDocumentCategoryUpdatedBy string     `gorm:"column:prod_doc_ctgry_updated_by" json:"prod_doc_ctgry_updated_by"`
	ProductDocumentCategoryUpdatedAt *time.Time `gorm:"column:prod_doc_ctgry_updated_at" json:"prod_doc_ctgry_updated_at"`
	ProductDocumentCategoryDeletedBy string     `gorm:"column:prod_doc_ctgry_deleted_by" json:"prod_doc_ctgry_deleted_by"`
	ProductDocumentCategoryDeletedAt *time.Time `gorm:"column:prod_doc_ctgry_deleted_at" json:"prod_doc_ctgry_deleted_at"`
}

type ResponseProductDocumentCategorys struct {
	Status  int                            `json:"status"`
	Message string                         `json:"message"`
	Data    *[]ProductDocumentCategoryData `json:"data"`
}

type ResponseProductDocumentCategory struct {
	Status  int                          `json:"status"`
	Message string                       `json:"message"`
	Data    *ProductDocumentCategoryData `json:"data"`
}

type ResponseProductDocumentCategoryCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (ProductDocumentCategory) TableName() string {
	return "m_product_document_category"
}

func (ProductDocumentCategoryData) TableName() string {
	return "m_product_document_category"
}

func (p *ProductDocumentCategory) Prepare() {
	p.ProductDocumentCategoryCode = html.EscapeString(strings.TrimSpace(p.ProductDocumentCategoryCode))
	p.ProductDocumentCategoryName = html.EscapeString(strings.TrimSpace(p.ProductDocumentCategoryName))
	p.ProductDocumentCategoryDesc = p.ProductDocumentCategoryDesc
	p.ProductDocumentCategoryStatus = p.ProductDocumentCategoryStatus
	p.ProductDocumentCategoryCreatedBy = html.EscapeString(strings.TrimSpace(p.ProductDocumentCategoryCreatedBy))
	p.ProductDocumentCategoryCreatedAt = time.Now()
	p.ProductDocumentCategoryUpdatedBy = html.EscapeString(strings.TrimSpace(p.ProductDocumentCategoryUpdatedBy))
	p.ProductDocumentCategoryUpdatedAt = p.ProductDocumentCategoryUpdatedAt
	p.ProductDocumentCategoryDeletedBy = html.EscapeString(strings.TrimSpace(p.ProductDocumentCategoryDeletedBy))
	p.ProductDocumentCategoryDeletedAt = p.ProductDocumentCategoryDeletedAt
}

func (p *ProductDocumentCategory) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.ProductDocumentCategoryCode == "" {
			return errors.New("Required Country")
		}
		if p.ProductDocumentCategoryName == "" {
			return errors.New("Required Province")
		}
		if p.ProductDocumentCategoryDesc == "" {
			return errors.New("Required Regency")
		}
		return nil
	}
}

func (p *ProductDocumentCategory) SaveProductDocumentCategory(db *gorm.DB) (*ProductDocumentCategory, error) {
	var err error
	err = db.Debug().Model(&ProductDocumentCategory{}).Create(&p).Error
	if err != nil {
		return &ProductDocumentCategory{}, err
	}
	return p, nil
}

func (p *ProductDocumentCategory) FindAllProductDocumentCategory(db *gorm.DB) (*[]ProductDocumentCategoryData, error) {
	var err error
	address := []ProductDocumentCategoryData{}
	err = db.Debug().Model(&ProductDocumentCategoryData{}).Limit(100).
		Select(`m_product_document_category.prod_doc_ctgry_id,
				m_product_document_category.prod_doc_ctgry_code,
				m_product_document_category.prod_doc_ctgry_name,
				m_product_document_category.prod_doc_ctgry_desc,
				m_product_document_category.prod_doc_ctgry_status,
				m_product_document_category.prod_doc_ctgry_created_by,
				m_product_document_category.prod_doc_ctgry_created_at at time zone 'Asia/Jakarta' as prod_doc_ctgry_created_at,
				m_product_document_category.prod_doc_ctgry_updated_by,
				m_product_document_category.prod_doc_ctgry_updated_at at time zone 'Asia/Jakarta' as prod_doc_ctgry_updated_at,
				m_product_document_category.prod_doc_ctgry_deleted_by,
				m_product_document_category.prod_doc_ctgry_deleted_at at time zone 'Asia/Jakarta' as prod_doc_ctgry_deleted_at`).
		Find(&address).Error
	if err != nil {
		return &[]ProductDocumentCategoryData{}, err
	}
	return &address, nil
}

func (p *ProductDocumentCategory) FindAllProductDocumentCategoryStatus(db *gorm.DB, status uint64) (*[]ProductDocumentCategoryData, error) {
	var err error
	address := []ProductDocumentCategoryData{}
	err = db.Debug().Model(&ProductDocumentCategoryData{}).Limit(100).
		Select(`m_product_document_category.prod_doc_ctgry_id,
				m_product_document_category.prod_doc_ctgry_code,
				m_product_document_category.prod_doc_ctgry_name,
				m_product_document_category.prod_doc_ctgry_desc,
				m_product_document_category.prod_doc_ctgry_status,
				m_product_document_category.prod_doc_ctgry_created_by,
				m_product_document_category.prod_doc_ctgry_created_at at time zone 'Asia/Jakarta' as prod_doc_ctgry_created_at,
				m_product_document_category.prod_doc_ctgry_updated_by,
				m_product_document_category.prod_doc_ctgry_updated_at at time zone 'Asia/Jakarta' as prod_doc_ctgry_updated_at,
				m_product_document_category.prod_doc_ctgry_deleted_by,
				m_product_document_category.prod_doc_ctgry_deleted_at at time zone 'Asia/Jakarta' as prod_doc_ctgry_deleted_at`).
		Where("prod_doc_ctgry_status = ?", status).
		Find(&address).Error
	if err != nil {
		return &[]ProductDocumentCategoryData{}, err
	}
	return &address, nil
}

func (p *ProductDocumentCategory) FindProductDocumentCategoryDataByID(db *gorm.DB, pid uint64) (*ProductDocumentCategory, error) {
	var err error
	err = db.Debug().Model(&ProductDocumentCategory{}).
		Select(`m_product_document_category.prod_doc_ctgry_id,
				m_product_document_category.prod_doc_ctgry_code,
				m_product_document_category.prod_doc_ctgry_name,
				m_product_document_category.prod_doc_ctgry_desc,
				m_product_document_category.prod_doc_ctgry_status,
				m_product_document_category.prod_doc_ctgry_created_by,
				m_product_document_category.prod_doc_ctgry_created_at at time zone 'Asia/Jakarta' as prod_doc_ctgry_created_at,
				m_product_document_category.prod_doc_ctgry_updated_by,
				m_product_document_category.prod_doc_ctgry_updated_at at time zone 'Asia/Jakarta' as prod_doc_ctgry_updated_at,
				m_product_document_category.prod_doc_ctgry_deleted_by,
				m_product_document_category.prod_doc_ctgry_deleted_at at time zone 'Asia/Jakarta' as prod_doc_ctgry_deleted_at`).
		Where("prod_doc_ctgry_id = ?", pid).
		Take(&p).Error
	if err != nil {
		return &ProductDocumentCategory{}, err
	}
	return p, nil
}

func (p *ProductDocumentCategory) FindProductDocumentCategoryData(db *gorm.DB, pid uint64) (*ProductDocumentCategoryData, error) {
	var err error
	documentdata := ProductDocumentCategoryData{}
	err = db.Debug().Model(&ProductDocumentCategoryData{}).Limit(100).
		Select(`m_product_document_category.prod_doc_ctgry_id,
				m_product_document_category.prod_doc_ctgry_code,
				m_product_document_category.prod_doc_ctgry_name,
				m_product_document_category.prod_doc_ctgry_desc,
				m_product_document_category.prod_doc_ctgry_status,
				m_product_document_category.prod_doc_ctgry_created_by,
				m_product_document_category.prod_doc_ctgry_created_at at time zone 'Asia/Jakarta' as prod_doc_ctgry_created_at,
				m_product_document_category.prod_doc_ctgry_updated_by,
				m_product_document_category.prod_doc_ctgry_updated_at at time zone 'Asia/Jakarta' as prod_doc_ctgry_updated_at,
				m_product_document_category.prod_doc_ctgry_deleted_by,
				m_product_document_category.prod_doc_ctgry_deleted_at at time zone 'Asia/Jakarta' as prod_doc_ctgry_deleted_at`).
		Where("prod_doc_ctgry_id = ?", pid).
		Find(&documentdata).Error
	if err != nil {
		return &ProductDocumentCategoryData{}, err
	}
	return &documentdata, nil
}

func (p *ProductDocumentCategory) FindProductDocumentCategoryByID(db *gorm.DB, pid uint64) (*ProductDocumentCategoryData, error) {
	var err error
	address := ProductDocumentCategoryData{}
	err = db.Debug().Model(&ProductDocumentCategoryData{}).Limit(100).
		Select(`m_product_document_category.prod_doc_ctgry_id,
				m_product_document_category.prod_doc_ctgry_code,
				m_product_document_category.prod_doc_ctgry_name,
				m_product_document_category.prod_doc_ctgry_desc,
				m_product_document_category.prod_doc_ctgry_status,
				m_product_document_category.prod_doc_ctgry_created_by,
				m_product_document_category.prod_doc_ctgry_created_at at time zone 'Asia/Jakarta' as prod_doc_ctgry_created_at,
				m_product_document_category.prod_doc_ctgry_updated_by,
				m_product_document_category.prod_doc_ctgry_updated_at at time zone 'Asia/Jakarta' as prod_doc_ctgry_updated_at,
				m_product_document_category.prod_doc_ctgry_deleted_by,
				m_product_document_category.prod_doc_ctgry_deleted_at at time zone 'Asia/Jakarta' as prod_doc_ctgry_deleted_at`).
		Where("prod_doc_ctgry_id = ?", pid).
		Find(&address).Error
	if err != nil {
		return &ProductDocumentCategoryData{}, err
	}
	return &address, nil
}

func (p *ProductDocumentCategory) FindProductDocumentCategoryStatusByID(db *gorm.DB, pid uint64, status uint64) (*ProductDocumentCategoryData, error) {
	var err error
	address := ProductDocumentCategoryData{}
	err = db.Debug().Model(&ProductDocumentCategoryData{}).Limit(100).
		Select(`m_product_document_category.prod_doc_ctgry_id,
				m_product_document_category.prod_doc_ctgry_code,
				m_product_document_category.prod_doc_ctgry_name,
				m_product_document_category.prod_doc_ctgry_desc,
				m_product_document_category.prod_doc_ctgry_status,
				m_product_document_category.prod_doc_ctgry_created_by,
				m_product_document_category.prod_doc_ctgry_created_at at time zone 'Asia/Jakarta' as prod_doc_ctgry_created_at,
				m_product_document_category.prod_doc_ctgry_updated_by,
				m_product_document_category.prod_doc_ctgry_updated_at at time zone 'Asia/Jakarta' as prod_doc_ctgry_updated_at,
				m_product_document_category.prod_doc_ctgry_deleted_by,
				m_product_document_category.prod_doc_ctgry_deleted_at at time zone 'Asia/Jakarta' as prod_doc_ctgry_deleted_at`).
		Where("prod_doc_ctgry_id = ? AND prod_doc_ctgry_status = ?", pid, status).
		Find(&address).Error
	if err != nil {
		return &ProductDocumentCategoryData{}, err
	}
	return &address, nil
}

func (p *ProductDocumentCategory) UpdateProductDocumentCategory(db *gorm.DB, pid uint64) (*ProductDocumentCategory, error) {
	var err error
	db = db.Debug().Model(&ProductDocumentCategory{}).Where("prod_doc_ctgry_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_ctgry_code":       p.ProductDocumentCategoryCode,
			"prod_doc_ctgry_name":       p.ProductDocumentCategoryName,
			"prod_doc_ctgry_desc":       p.ProductDocumentCategoryDesc,
			"prod_doc_ctgry_status":     p.ProductDocumentCategoryStatus,
			"prod_doc_ctgry_updated_by": p.ProductDocumentCategoryUpdatedBy,
			"prod_doc_ctgry_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ProductDocumentCategory{}).Where("prod_doc_ctgry_id = ?", pid).Error
	if err != nil {
		return &ProductDocumentCategory{}, err
	}
	return p, nil
}

func (p *ProductDocumentCategory) UpdateProductDocumentCategoryStatus(db *gorm.DB, pid uint64) (*ProductDocumentCategory, error) {
	var err error
	db = db.Debug().Model(&ProductDocumentCategory{}).Where("prod_doc_ctgry_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_ctgry_status":     p.ProductDocumentCategoryStatus,
			"prod_doc_ctgry_updated_by": p.ProductDocumentCategoryUpdatedBy,
			"prod_doc_ctgry_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ProductDocumentCategory{}).Where("prod_doc_ctgry_id = ?", pid).Error
	if err != nil {
		return &ProductDocumentCategory{}, err
	}
	return p, nil
}

func (p *ProductDocumentCategory) UpdateProductDocumentCategoryStatusOnly(db *gorm.DB, pid uint64) (*ProductDocumentCategory, error) {
	var err error
	db = db.Debug().Model(&ProductDocumentCategory{}).Where("prod_doc_ctgry_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_ctgry_status": p.ProductDocumentCategoryStatus,
		},
	)
	err = db.Debug().Model(&ProductDocumentCategory{}).Where("prod_doc_ctgry_id = ?", pid).Error
	if err != nil {
		return &ProductDocumentCategory{}, err
	}
	return p, nil
}

func (p *ProductDocumentCategory) UpdateProductDocumentCategoryStatusDelete(db *gorm.DB, pid uint64) (*ProductDocumentCategory, error) {
	var err error
	db = db.Debug().Model(&ProductDocumentCategory{}).Where("prod_doc_ctgry_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_ctgry_status":     3,
			"prod_doc_ctgry_deleted_by": p.ProductDocumentCategoryDeletedBy,
			"prod_doc_ctgry_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ProductDocumentCategory{}).Where("prod_doc_ctgry_id = ?", pid).Error
	if err != nil {
		return &ProductDocumentCategory{}, err
	}
	return p, nil
}

func (p *ProductDocumentCategory) DeleteProductDocumentCategory(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&ProductDocumentCategory{}).Where("prod_doc_ctgry_id = ?", pid).Take(&ProductDocumentCategory{}).Delete(&ProductDocumentCategory{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("ProductDocumentCategory not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================
