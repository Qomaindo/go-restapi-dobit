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

type ProductDocument struct {
	ProductDocumentID          uint64     `gorm:"column:prod_doc_id;primary_key;" json:"prod_doc_id"`
	ProductDocumentCategoryID  uint64     `gorm:"column:prod_doc_ctgry_id" json:"prod_doc_ctgry_id"`
	ProductDocumentCode        string     `gorm:"column:prod_doc_code" json:"prod_doc_code"`
	ProductDocumentName        string     `gorm:"column:prod_doc_name" json:"prod_doc_name"`
	ProductDocumentDescription string     `gorm:"column:prod_doc_desc" json:"prod_doc_desc"`
	ProductDocumentStatus      uint64     `gorm:"column:prod_doc_status;size:2" json:"prod_doc_status"`
	ProductDocumentCreatedBy   string     `gorm:"column:prod_doc_created_by;size:125" json:"prod_doc_created_by"`
	ProductDocumentCreatedAt   time.Time  `gorm:"column:prod_doc_created_at;default:CURRENT_TIMESTAMP" json:"prod_doc_created_at"`
	ProductDocumentUpdatedBy   string     `gorm:"column:prod_doc_updated_by;size:125" json:"prod_doc_updated_by"`
	ProductDocumentUpdatedAt   *time.Time `gorm:"column:prod_doc_updated_at;default:CURRENT_TIMESTAMP" json:"prod_doc_created_at"`
	ProductDocumentDeletedBy   string     `gorm:"column:prod_doc_deleted_by;size:125" json:"prod_doc_deleted_by"`
	ProductDocumentDeletedAt   *time.Time `gorm:"column:prod_doc_deleted_at;default:CURRENT_TIMESTAMP" json:"prod_doc_deleted_at"`
}

type ProductDocumentData struct {
	ProductDocumentID           uint64     `gorm:"column:prod_doc_id" json:"prod_doc_id"`
	ProductDocumentCategoryID   uint64     `gorm:"column:prod_doc_ctgry_id" json:"prod_doc_ctgry_id"`
	ProductDocumentCategoryName string     `gorm:"column:prod_doc_ctgry_name" json:"prod_doc_ctgry_name"`
	ProductDocumentCode         string     `gorm:"column:prod_doc_code" json:"prod_doc_code"`
	ProductDocumentName         string     `gorm:"column:prod_doc_name" json:"prod_doc_name"`
	ProductDocumentDescription  string     `gorm:"column:prod_doc_desc" json:"prod_doc_desc"`
	ProductDocumentStatus       uint64     `gorm:"column:prod_doc_status" json:"prod_doc_status"`
	ProductDocumentCreatedBy    string     `gorm:"column:prod_doc_created_by" json:"prod_doc_created_by"`
	ProductDocumentCreatedAt    time.Time  `gorm:"column:prod_doc_created_at" json:"prod_doc_created_at"`
	ProductDocumentUpdatedBy    string     `gorm:"column:prod_doc_updated_by" json:"prod_doc_updated_by"`
	ProductDocumentUpdatedAt    *time.Time `gorm:"column:prod_doc_updated_at" json:"prod_doc_updated_at"`
	ProductDocumentDeletedBy    string     `gorm:"column:prod_doc_deleted_by" json:"prod_doc_deleted_by"`
	ProductDocumentDeletedAt    *time.Time `gorm:"column:prod_doc_deleted_at" json:"prod_doc_deleted_at"`
}

type ResponseProductDocuments struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    *[]ProductDocumentData `json:"data"`
}

type ResponseProductDocument struct {
	Status  int                  `json:"status"`
	Message string               `json:"message"`
	Data    *ProductDocumentData `json:"data"`
}

type ResponseProductDocumentCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (ProductDocument) TableName() string {
	return "m_product_document"
}

func (ProductDocumentData) TableName() string {
	return "m_product_document"
}

func (p *ProductDocument) Prepare() {
	p.ProductDocumentCategoryID = p.ProductDocumentCategoryID
	p.ProductDocumentCode = html.EscapeString(strings.TrimSpace(p.ProductDocumentCode))
	p.ProductDocumentName = html.EscapeString(strings.TrimSpace(p.ProductDocumentName))
	p.ProductDocumentDescription = p.ProductDocumentDescription
	p.ProductDocumentStatus = p.ProductDocumentStatus
	p.ProductDocumentCreatedBy = html.EscapeString(strings.TrimSpace(p.ProductDocumentCreatedBy))
	p.ProductDocumentCreatedAt = time.Now()
	p.ProductDocumentUpdatedBy = html.EscapeString(strings.TrimSpace(p.ProductDocumentUpdatedBy))
	p.ProductDocumentUpdatedAt = p.ProductDocumentUpdatedAt
	p.ProductDocumentDeletedBy = html.EscapeString(strings.TrimSpace(p.ProductDocumentDeletedBy))
	p.ProductDocumentDeletedAt = p.ProductDocumentDeletedAt
}

func (p *ProductDocument) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.ProductDocumentCategoryID == 0 {
			return errors.New("Required Country")
		}
		if p.ProductDocumentCode == "" {
			return errors.New("Required Province")
		}
		if p.ProductDocumentName == "" {
			return errors.New("Required Regency")
		}
		return nil
	}
}

func (p *ProductDocument) SaveProductDocument(db *gorm.DB) (*ProductDocument, error) {
	var err error
	err = db.Debug().Model(&ProductDocument{}).Create(&p).Error
	if err != nil {
		return &ProductDocument{}, err
	}
	return p, nil
}

func (p *ProductDocument) FindAllProductDocument(db *gorm.DB) (*[]ProductDocumentData, error) {
	var err error
	productdocument := []ProductDocumentData{}
	err = db.Debug().Model(&ProductDocumentData{}).Limit(100).
		Select(`m_product_document.prod_doc_id,
				m_product_document_category.prod_doc_ctgry_id,
				m_product_document_category.prod_doc_ctgry_name,
				m_product_document.prod_doc_code,
				m_product_document.prod_doc_name,
				m_product_document.prod_doc_desc,
				m_product_document.prod_doc_status,
				m_product_document.prod_doc_created_by,
				m_product_document.prod_doc_created_at at time zone 'Asia/Jakarta' as prod_doc_created_at,
				m_product_document.prod_doc_updated_by,
				m_product_document.prod_doc_updated_at at time zone 'Asia/Jakarta' as prod_doc_updated_at,
				m_product_document.prod_doc_deleted_by,
				m_product_document.prod_doc_deleted_at at time zone 'Asia/Jakarta' as prod_doc_deleted_at`).
		Joins("JOIN m_product_document_category on m_product_document.prod_doc_ctgry_id=m_product_document_category.prod_doc_ctgry_id").
		Find(&productdocument).Error
	if err != nil {
		return &[]ProductDocumentData{}, err
	}
	return &productdocument, nil
}

func (p *ProductDocument) FindAllProductDocumentStatus(db *gorm.DB, status uint64) (*[]ProductDocumentData, error) {
	var err error
	productdocument := []ProductDocumentData{}
	err = db.Debug().Model(&ProductDocumentData{}).Limit(100).
		Select(`m_product_document.prod_doc_id,
				m_product_document_category.prod_doc_ctgry_id,
				m_product_document_category.prod_doc_ctgry_name,
				m_product_document.prod_doc_code,
				m_product_document.prod_doc_name,
				m_product_document.prod_doc_desc,
				m_product_document.prod_doc_status,
				m_product_document.prod_doc_created_by,
				m_product_document.prod_doc_created_at at time zone 'Asia/Jakarta' as prod_doc_created_at,
				m_product_document.prod_doc_updated_by,
				m_product_document.prod_doc_updated_at at time zone 'Asia/Jakarta' as prod_doc_updated_at,
				m_product_document.prod_doc_deleted_by,
				m_product_document.prod_doc_deleted_at at time zone 'Asia/Jakarta' as prod_doc_deleted_at`).
		Joins("JOIN m_product_document_category on m_product_document.prod_doc_ctgry_id=m_product_document_category.prod_doc_ctgry_id").
		Where("prod_doc_status = ?", status).
		Find(&productdocument).Error
	if err != nil {
		return &[]ProductDocumentData{}, err
	}
	return &productdocument, nil
}

func (p *ProductDocument) FindProductDocumentDataByID(db *gorm.DB, pid uint64) (*ProductDocument, error) {
	var err error
	err = db.Debug().Model(&ProductDocument{}).Where("prod_doc_id = ?", pid).Take(&p).Error
	if err != nil {
		return &ProductDocument{}, err
	}
	return p, nil
}

func (p *ProductDocument) FindProductDocumentByID(db *gorm.DB, pid uint64) (*ProductDocumentData, error) {
	var err error
	productdocument := ProductDocumentData{}
	err = db.Debug().Model(&ProductDocumentData{}).Limit(100).
		Select(`m_product_document.prod_doc_id,
				m_product_document_category.prod_doc_ctgry_id,
				m_product_document_category.prod_doc_ctgry_name,
				m_product_document.prod_doc_code,
				m_product_document.prod_doc_name,
				m_product_document.prod_doc_desc,
				m_product_document.prod_doc_status,
				m_product_document.prod_doc_created_by,
				m_product_document.prod_doc_created_at at time zone 'Asia/Jakarta' as prod_doc_created_at,
				m_product_document.prod_doc_updated_by,
				m_product_document.prod_doc_updated_at at time zone 'Asia/Jakarta' as prod_doc_updated_at,
				m_product_document.prod_doc_deleted_by,
				m_product_document.prod_doc_deleted_at at time zone 'Asia/Jakarta' as prod_doc_deleted_at`).
		Joins("JOIN m_product_document_category on m_product_document.prod_doc_ctgry_id=m_product_document_category.prod_doc_ctgry_id").
		Where("prod_doc_id = ?", pid).
		Find(&productdocument).Error
	if err != nil {
		return &ProductDocumentData{}, err
	}
	return &productdocument, nil
}

func (p *ProductDocument) FindProductDocumentStatusByID(db *gorm.DB, pid uint64, status uint64) (*ProductDocumentData, error) {
	var err error
	productdocument := ProductDocumentData{}
	err = db.Debug().Model(&ProductDocumentData{}).Limit(100).
		Select(`m_product_document.prod_doc_id,
				m_product_document_category.prod_doc_ctgry_id,
				m_product_document_category.prod_doc_ctgry_name,
				m_product_document.prod_doc_code,
				m_product_document.prod_doc_name,
				m_product_document.prod_doc_desc,
				m_product_document.prod_doc_status,
				m_product_document.prod_doc_created_by,
				m_product_document.prod_doc_created_at at time zone 'Asia/Jakarta' as prod_doc_created_at,
				m_product_document.prod_doc_updated_by,
				m_product_document.prod_doc_updated_at at time zone 'Asia/Jakarta' as prod_doc_updated_at,
				m_product_document.prod_doc_deleted_by,
				m_product_document.prod_doc_deleted_at at time zone 'Asia/Jakarta' as prod_doc_deleted_at`).
		Joins("JOIN m_product_document_category on m_product_document.prod_doc_ctgry_id=m_product_document_category.prod_doc_ctgry_id").
		Where("prod_doc_id = ? AND prod_doc_status = ?", pid, status).
		Find(&productdocument).Error
	if err != nil {
		return &ProductDocumentData{}, err
	}
	return &productdocument, nil
}

func (p *ProductDocument) UpdateProductDocument(db *gorm.DB, pid uint64) (*ProductDocument, error) {
	var err error
	db = db.Debug().Model(&ProductDocument{}).Where("prod_doc_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_ctgry_id":   p.ProductDocumentCategoryID,
			"prod_doc_code":       p.ProductDocumentCode,
			"prod_doc_name":       p.ProductDocumentName,
			"prod_doc_desc":       p.ProductDocumentDescription,
			"prod_doc_status":     p.ProductDocumentStatus,
			"prod_doc_updated_by": p.ProductDocumentUpdatedBy,
			"prod_doc_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ProductDocument{}).Where("prod_doc_id = ?", pid).Error
	if err != nil {
		return &ProductDocument{}, err
	}
	return p, nil
}

func (p *ProductDocument) UpdateProductDocumentStatus(db *gorm.DB, pid uint64) (*ProductDocument, error) {
	var err error
	db = db.Debug().Model(&ProductDocument{}).Where("prod_doc_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_status":     p.ProductDocumentStatus,
			"prod_doc_updated_by": p.ProductDocumentUpdatedBy,
			"prod_doc_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ProductDocument{}).Where("prod_doc_id = ?", pid).Error
	if err != nil {
		return &ProductDocument{}, err
	}
	return p, nil
}

func (p *ProductDocument) UpdateProductDocumentStatusOnly(db *gorm.DB, pid uint64) (*ProductDocument, error) {
	var err error
	db = db.Debug().Model(&ProductDocument{}).Where("prod_doc_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_status": p.ProductDocumentStatus,
		},
	)
	err = db.Debug().Model(&ProductDocument{}).Where("prod_doc_id = ?", pid).Error
	if err != nil {
		return &ProductDocument{}, err
	}
	return p, nil
}

func (p *ProductDocument) UpdateProductDocumentStatusDelete(db *gorm.DB, pid uint64) (*ProductDocument, error) {
	var err error
	db = db.Debug().Model(&ProductDocument{}).Where("prod_doc_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_status":     3,
			"prod_doc_deleted_by": p.ProductDocumentDeletedBy,
			"prod_doc_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ProductDocument{}).Where("prod_doc_id = ?", pid).Error
	if err != nil {
		return &ProductDocument{}, err
	}
	return p, nil
}

func (p *ProductDocument) DeleteProductDocument(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&ProductDocument{}).Where("prod_doc_id = ?", pid).Take(&ProductDocument{}).Delete(&ProductDocument{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("ProductDocument not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================
