package models

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type ProductDocumentCustomerAttribute struct {
	ProductDocumentCustomerAttributeID        uint64     `gorm:"column:prod_doc_cust_attr_id;primary_key;" json:"prod_doc_cust_attr_id"`
	ProductDocumentCustomerID                 uint64     `gorm:"column:prod_doc_cust_id" json:"prod_doc_cust_id"`
	ProductDocumentCustomerAttributeCode      string     `gorm:"column:prod_doc_cust_attr_code;size:15" json:"prod_doc_cust_attr_code"`
	ProductDocumentCustomerAttributeName      string     `gorm:"column:prod_doc_cust_attr_name;size:255" json:"prod_doc_cust_attr_name"`
	ProductDocumentCustomerAttributeDesc      string     `gorm:"column:prod_doc_cust_attr_desc" json:"prod_doc_cust_attr_desc"`
	ProductDocumentCustomerAttributeType      string     `gorm:"column:prod_doc_cust_attr_type;size:255" json:"prod_doc_cust_attr_type"`
	ProductDocumentCustomerAttributeStatus    uint64     `gorm:"column:prod_doc_cust_attr_status;size:2" json:"prod_doc_cust_attr_status"`
	ProductDocumentCustomerAttributeCreatedBy string     `gorm:"column:prod_doc_cust_attr_created_by;size:125" json:"prod_doc_cust_attr_created_by"`
	ProductDocumentCustomerAttributeUpdatedBy string     `gorm:"column:prod_doc_cust_attr_updated_by;size:125" json:"prod_doc_cust_attr_updated_by"`
	ProductDocumentCustomerAttributeDeletedBy string     `gorm:"column:prod_doc_cust_attr_deleted_by;size:125" json:"prod_doc_cust_attr_deleted_by"`
	ProductDocumentCustomerAttributeCreatedAt time.Time  `gorm:"column:prod_doc_cust_attr_created_at;default:CURRENT_TIMESTAMP" json:"prod_doc_cust_attr_created_at"`
	ProductDocumentCustomerAttributeUpdatedAt *time.Time `gorm:"column:prod_doc_cust_attr_updated_at;default:CURRENT_TIMESTAMP" json:"prod_doc_cust_attr_updated_at"`
	ProductDocumentCustomerAttributeDeletedAt *time.Time `gorm:"column:prod_doc_cust_attr_deleted_at;default:CURRENT_TIMESTAMP" json:"prod_doc_cust_attr_deleted_at"`
}

type ProductDocumentCustomerAttributeData struct {
	ProductDocumentCustomerAttributeID        uint64     `gorm:"column:prod_doc_cust_attr_id;primary_key;" json:"prod_doc_cust_attr_id"`
	ProductDocumentCustomerAttributeCode      string     `gorm:"column:prod_doc_cust_attr_code;size:15" json:"prod_doc_cust_attr_code"`
	ProductDocumentCustomerAttributeName      string     `gorm:"column:prod_doc_cust_attr_name;size:255" json:"prod_doc_cust_attr_name"`
	ProductDocumentCustomerAttributeDesc      string     `gorm:"column:prod_doc_cust_attr_desc" json:"prod_doc_cust_attr_desc"`
	ProductDocumentCustomerAttributeType      string     `gorm:"column:prod_doc_cust_attr_type;size:255" json:"prod_doc_cust_attr_type"`
	ProductDocumentCustomerAttributeStatus    uint64     `gorm:"column:prod_doc_cust_attr_status;size:2" json:"prod_doc_cust_attr_status"`
	ProductDocumentCustomerID                 uint64     `gorm:"column:prod_doc_cust_id" json:"prod_doc_cust_id"`
	ProductDocumentCustomerCode               string     `gorm:"column:prod_doc_cust_code;size:25" json:"prod_doc_cust_code"`
	ProductDocumentCustomerName               string     `gorm:"column:prod_doc_cust_name;size:255" json:"prod_doc_cust_name"`
	ProductDocumentCustomerStatus             uint64     `gorm:"column:prod_doc_cust_status;size:25" json:"prod_doc_cust_status"`
	ProductDocumentCustomerAttributeCreatedBy string     `gorm:"column:prod_doc_cust_attr_created_by;size:125" json:"prod_doc_cust_attr_created_by"`
	ProductDocumentCustomerAttributeUpdatedBy string     `gorm:"column:prod_doc_cust_attr_updated_by;size:125" json:"prod_doc_cust_attr_updated_by"`
	ProductDocumentCustomerAttributeDeletedBy string     `gorm:"column:prod_doc_cust_attr_deleted_by;size:125" json:"prod_doc_cust_attr_deleted_by"`
	ProductDocumentCustomerAttributeCreatedAt time.Time  `gorm:"column:prod_doc_cust_attr_created_at;default:CURRENT_TIMESTAMP" json:"prod_doc_cust_attr_created_at"`
	ProductDocumentCustomerAttributeUpdatedAt *time.Time `gorm:"column:prod_doc_cust_attr_updated_at;default:CURRENT_TIMESTAMP" json:"prod_doc_cust_attr_updated_at"`
	ProductDocumentCustomerAttributeDeletedAt *time.Time `gorm:"column:prod_doc_cust_attr_deleted_at;default:CURRENT_TIMESTAMP" json:"prod_doc_cust_attr_deleted_at"`
}

type ResponseProducDocumentCustomerAttributes struct {
	Status  int                                     `json:"status"`
	Message string                                  `json:"message"`
	Data    *[]ProductDocumentCustomerAttributeData `json:"data"`
}

type ResponseProducDocumentCustomerAttribute struct {
	Status  int                                   `json:"status"`
	Message string                                `json:"message"`
	Data    *ProductDocumentCustomerAttributeData `json:"data"`
}

type ResponseProducDocumentCustomerAttributeIU struct {
	Status  int                               `json:"status"`
	Message string                            `json:"message"`
	Data    *ProductDocumentCustomerAttribute `json:"data"`
}

type ResponseProducDocumentCustomerAttributeDel struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (ProductDocumentCustomerAttribute) TableName() string {
	return "m_product_document_customer_attribute"
}

func (ProductDocumentCustomerAttributeData) TableName() string {
	return "m_product_document_customer_attribute"
}

func (p *ProductDocumentCustomerAttribute) Prepare() {
	p.ProductDocumentCustomerAttributeID = p.ProductDocumentCustomerAttributeID
	p.ProductDocumentCustomerID = p.ProductDocumentCustomerID
	p.ProductDocumentCustomerAttributeCode = p.ProductDocumentCustomerAttributeCode
	p.ProductDocumentCustomerAttributeName = p.ProductDocumentCustomerAttributeName
	p.ProductDocumentCustomerAttributeDesc = p.ProductDocumentCustomerAttributeDesc
	p.ProductDocumentCustomerAttributeType = p.ProductDocumentCustomerAttributeType
	p.ProductDocumentCustomerAttributeStatus = p.ProductDocumentCustomerAttributeStatus
	p.ProductDocumentCustomerAttributeCreatedBy = p.ProductDocumentCustomerAttributeCreatedBy
	p.ProductDocumentCustomerAttributeUpdatedBy = p.ProductDocumentCustomerAttributeUpdatedBy
	p.ProductDocumentCustomerAttributeDeletedBy = p.ProductDocumentCustomerAttributeDeletedBy
	p.ProductDocumentCustomerAttributeCreatedAt = time.Now()
	p.ProductDocumentCustomerAttributeUpdatedAt = p.ProductDocumentCustomerAttributeUpdatedAt
	p.ProductDocumentCustomerAttributeDeletedAt = p.ProductDocumentCustomerAttributeDeletedAt
}

func (p *ProductDocumentCustomerAttribute) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.ProductDocumentCustomerAttributeID == 0 {
			return errors.New("Required Product Document Customer Attribute Code")
		}
		if p.ProductDocumentCustomerID == 0 {
			return errors.New("Required Product Document Customer Attribute Name")
		}
		return nil
	}
}

func (p *ProductDocumentCustomerAttribute) SaveProductgoryDocumentCustomerAttribute(db *gorm.DB) (*ProductDocumentCustomerAttribute, error) {
	var err error
	err = db.Debug().Model(&ProductDocumentCustomerAttribute{}).Create(&p).Error
	if err != nil {
		return &ProductDocumentCustomerAttribute{}, err
	}
	return p, nil
}

func (p *ProductDocumentCustomerAttribute) FindAllProductategoryDocumentCustomerAttribute(db *gorm.DB) (*[]ProductDocumentCustomerAttributeData, error) {
	var err error
	productdocumentcustomerattribute := []ProductDocumentCustomerAttributeData{}
	err = db.Debug().Model(&ProductDocumentCustomerAttributeData{}).Limit(100).
		Select(`m_product_document_customer_attribute.prod_doc_attr_id,
				m_product_document_customer_attribute.prod_doc_attr_code,
				m_product_document_customer_attribute.prod_doc_attr_name,
				m_product_document_customer.prod_doc_cust_id,
				m_product_document_customer.prod_doc_cust_code,
				m_product_document_customer.prod_doc_cust_name,
				m_product_document_customer_attribute.prod_doc_attr_desc,
				m_product_document_customer_attribute.prod_doc_attr_type,
				m_product_document_customer_attribute.prod_doc_cust_attr_status,
				m_product_document_customer_attribute.prod_doc_attr_created_by,
				m_product_document_customer_attribute.prod_doc_attr_updated_by,
				m_product_document_customer_attribute.prod_doc_attr_deleted_by,
				m_product_document_customer_attribute.prod_doc_attr_created_at,
				m_product_document_customer_attribute.prod_doc_attr_updated_at,
				m_product_document_customer_attribute.prod_doc_attr_deleted_at
				`).
		// Joins("JOIN m_product_document on m_product_document_customer_attribute.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_document_customer on m_product_document_customer_attribute.prod_doc_cust_id=m_product_document_customer.prod_doc_cust_id").
		Find(&productdocumentcustomerattribute).Error
	if err != nil {
		return &[]ProductDocumentCustomerAttributeData{}, err
	}
	return &productdocumentcustomerattribute, nil
}

func (p *ProductDocumentCustomerAttribute) FindAllProductategoryDocumentCustomerAttributeStatus(db *gorm.DB, status uint64) (*[]ProductDocumentCustomerAttributeData, error) {
	var err error
	productdocumentcustomerattribute := []ProductDocumentCustomerAttributeData{}
	err = db.Debug().Model(&ProductDocumentCustomerAttributeData{}).Limit(100).
		Select(`m_product_document_customer_attribute.prod_doc_attr_id,
				m_product_document_customer_attribute.prod_doc_attr_code,
				m_product_document_customer_attribute.prod_doc_attr_name,
				m_product_document_customer.prod_doc_cust_id,
				m_product_document_customer.prod_doc_cust_code,
				m_product_document_customer.prod_doc_cust_name,
				m_product_document_customer_attribute.prod_doc_attr_desc,
				m_product_document_customer_attribute.prod_doc_attr_type,
				m_product_document_customer_attribute.prod_doc_cust_attr_status,
				m_product_document_customer_attribute.prod_doc_attr_created_by,
				m_product_document_customer_attribute.prod_doc_attr_updated_by,
				m_product_document_customer_attribute.prod_doc_attr_deleted_by,
				m_product_document_customer_attribute.prod_doc_attr_created_at,
				m_product_document_customer_attribute.prod_doc_attr_updated_at,
				m_product_document_customer_attribute.prod_doc_attr_deleted_at
				`).
		// Joins("JOIN m_product_document on m_product_document_customer_attribute.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_document_customer on m_product_document_customer_attribute.prod_doc_cust_id=m_product_document_customer.prod_doc_cust_id").
		Where("prod_doc_cust_attr_status = ?", status).
		Find(&productdocumentcustomerattribute).Error
	if err != nil {
		return &[]ProductDocumentCustomerAttributeData{}, err
	}
	return &productdocumentcustomerattribute, nil
}

func (p *ProductDocumentCustomerAttribute) FindProductgoryDocumentCustomerAttributeByID(db *gorm.DB, pid uint64) (*ProductDocumentCustomerAttributeData, error) {
	var err error
	productdocumentcustomerattribute := ProductDocumentCustomerAttributeData{}
	err = db.Debug().Model(&ProductDocumentCustomerAttributeData{}).Limit(100).
		Select(`m_product_document_customer_attribute.prod_doc_attr_id,
				m_product_document_customer_attribute.prod_doc_attr_code,
				m_product_document_customer_attribute.prod_doc_attr_name,
				m_product_document_customer.prod_doc_cust_id,
				m_product_document_customer.prod_doc_cust_code,
				m_product_document_customer.prod_doc_cust_name,
				m_product_document_customer_attribute.prod_doc_attr_desc,
				m_product_document_customer_attribute.prod_doc_attr_type,
				m_product_document_customer_attribute.prod_doc_cust_attr_status,
				m_product_document_customer_attribute.prod_doc_attr_created_by,
				m_product_document_customer_attribute.prod_doc_attr_updated_by,
				m_product_document_customer_attribute.prod_doc_attr_deleted_by,
				m_product_document_customer_attribute.prod_doc_attr_created_at,
				m_product_document_customer_attribute.prod_doc_attr_updated_at,
				m_product_document_customer_attribute.prod_doc_attr_deleted_at
				`).
		// Joins("JOIN m_product_document on m_product_document_customer_attribute.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_document_customer on m_product_document_customer_attribute.prod_doc_cust_id=m_product_document_customer.prod_doc_cust_id").
		Where("prod_doc_attr_id = ?", pid).
		Find(&productdocumentcustomerattribute).Error
	if err != nil {
		return &ProductDocumentCustomerAttributeData{}, err
	}
	return &productdocumentcustomerattribute, nil
}

func (p *ProductDocumentCustomerAttribute) FindProductgoryDocumentCustomerAttributeStatusByID(db *gorm.DB, pid uint64, status uint64) (*ProductDocumentCustomerAttributeData, error) {
	var err error
	productdocumentcustomerattribute := ProductDocumentCustomerAttributeData{}
	err = db.Debug().Model(&ProductDocumentCustomerAttributeData{}).Limit(100).
		Select(`m_product_document_customer_attribute.prod_doc_attr_id,
				m_product_document_customer_attribute.prod_doc_attr_code,
				m_product_document_customer_attribute.prod_doc_attr_name,
				m_product_document_customer.prod_doc_cust_id,
				m_product_document_customer.prod_doc_cust_code,
				m_product_document_customer.prod_doc_cust_name,
				m_product_document_customer_attribute.prod_doc_attr_desc,
				m_product_document_customer_attribute.prod_doc_attr_type,
				m_product_document_customer_attribute.prod_doc_cust_attr_status,
				m_product_document_customer_attribute.prod_doc_attr_created_by,
				m_product_document_customer_attribute.prod_doc_attr_updated_by,
				m_product_document_customer_attribute.prod_doc_attr_deleted_by,
				m_product_document_customer_attribute.prod_doc_attr_created_at,
				m_product_document_customer_attribute.prod_doc_attr_updated_at,
				m_product_document_customer_attribute.prod_doc_attr_deleted_at
				`).
		// Joins("JOIN m_product_document on m_product_document_customer_attribute.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_document_customer on m_product_document_customer_attribute.prod_doc_cust_id=m_product_document_customer.prod_doc_cust_id").
		Where("prod_doc_attr_id = ? AND prod_doc_cust_attr_status = ?", pid, status).
		Find(&productdocumentcustomerattribute).Error
	if err != nil {
		return &ProductDocumentCustomerAttributeData{}, err
	}
	return &productdocumentcustomerattribute, nil
}

func (p *ProductDocumentCustomerAttribute) FindProductgoryDocumentCustomerAttributeBySubCategoryByID(db *gorm.DB, pid uint64, subcategory uint64) (*ProductDocumentCustomerAttributeData, error) {
	var err error
	productdocumentcustomerattribute := ProductDocumentCustomerAttributeData{}
	err = db.Debug().Model(&ProductDocumentCustomerAttributeData{}).Limit(100).
		Select(`m_product_document_customer_attribute.prod_doc_attr_id,
				m_product_document_customer_attribute.prod_doc_attr_code,
				m_product_document_customer_attribute.prod_doc_attr_name,
				m_product_document_customer.prod_doc_cust_id,
				m_product_document_customer.prod_doc_cust_code,
				m_product_document_customer.prod_doc_cust_name,
				m_product_document_customer_attribute.prod_doc_attr_desc,
				m_product_document_customer_attribute.prod_doc_attr_type,
				m_product_document_customer_attribute.prod_doc_cust_attr_status,
				m_product_document_customer_attribute.prod_doc_attr_created_by,
				m_product_document_customer_attribute.prod_doc_attr_updated_by,
				m_product_document_customer_attribute.prod_doc_attr_deleted_by,
				m_product_document_customer_attribute.prod_doc_attr_created_at,
				m_product_document_customer_attribute.prod_doc_attr_updated_at,
				m_product_document_customer_attribute.prod_doc_attr_deleted_at
				`).
		// Joins("JOIN m_product_document on m_product_document_customer_attribute.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_document_customer on m_product_document_customer_attribute.prod_doc_cust_id=m_product_document_customer.prod_doc_cust_id").
		Where("prod_doc_attr_id = ? AND prod_doc_cust_attr_status = ?", pid, subcategory).
		Find(&productdocumentcustomerattribute).Error
	if err != nil {
		return &ProductDocumentCustomerAttributeData{}, err
	}
	return &productdocumentcustomerattribute, nil
}

func (p *ProductDocumentCustomerAttribute) UpdateProductDocumentCustomerAttribute(db *gorm.DB, pid uint64) (*ProductDocumentCustomerAttribute, error) {
	var err error
	db = db.Debug().Model(&ProductDocumentCustomerAttribute{}).Where("prod_doc_attr_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_cust_id":          p.ProductDocumentCustomerID,
			"prod_doc_cust_attr_code":   p.ProductDocumentCustomerAttributeCode,
			"prod_doc_cust_attr_name":   p.ProductDocumentCustomerAttributeName,
			"prod_doc_cust_attr_desc":   p.ProductDocumentCustomerAttributeDesc,
			"prod_doc_cust_attr_type":   p.ProductDocumentCustomerAttributeType,
			"prod_doc_attr_updated_by":  p.ProductDocumentCustomerAttributeUpdatedBy,
			"prod_doc_cust_attr_status": time.Now(),
		},
	)
	err = db.Debug().Model(&ProductDocumentCustomerAttribute{}).Where("prod_doc_attr_id = ?", pid).Error
	if err != nil {
		return &ProductDocumentCustomerAttribute{}, err
	}
	return p, nil
}

func (p *ProductDocumentCustomerAttribute) UpdateProducttegoryDocumentCustomerAttributeStatus(db *gorm.DB, pid uint64) (*ProductDocumentCustomerAttribute, error) {
	var err error
	db = db.Debug().Model(&ProductDocumentCustomerAttribute{}).Where("prod_doc_attr_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_cust_attr_status":     p.ProductDocumentCustomerAttributeStatus,
			"prod_doc_cust_attr_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ProductDocumentCustomerAttribute{}).Where("prod_doc_attr_id = ?", pid).Error
	if err != nil {
		return &ProductDocumentCustomerAttribute{}, err
	}
	return p, nil
}

func (p *ProductDocumentCustomerAttribute) UpdateProducttegoryDocumentCustomerAttributeStatusDelete(db *gorm.DB, pid uint64) (*ProductDocumentCustomerAttribute, error) {
	var err error
	db = db.Debug().Model(&ProductDocumentCustomerAttribute{}).Where("prod_doc_attr_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_cust_attr_status":     3,
			"prod_doc_cust_attr_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ProductDocumentCustomerAttribute{}).Where("prod_doc_attr_id = ?", pid).Error
	if err != nil {
		return &ProductDocumentCustomerAttribute{}, err
	}
	return p, nil
}

func (p *ProductDocumentCustomerAttribute) DeleteProducttegoryDocumentCustomerAttribute(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&ProductDocumentCustomerAttribute{}).Where("prod_doc_attr_id = ?", pid).Take(&ProductDocumentCustomerAttribute{}).Delete(&ProductDocumentCustomerAttribute{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("ProductDocumentCustomerAttribute not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
