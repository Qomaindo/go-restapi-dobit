package models

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type ProductDocumentMitraAttribute struct {
	ProductDocumentMitraAttributeID        uint64     `gorm:"column:prod_doc_mitra_attr_id;primary_key;" json:"prod_doc_mitra_attr_id"`
	ProductDocumentMitraID                 uint64     `gorm:"column:prod_doc_mitra_id" json:"prod_doc_mitra_id"`
	ProductDocumentMitraAttributeCode      string     `gorm:"column:prod_doc_mitra_attr_code;size:15" json:"prod_doc_mitra_attr_code"`
	ProductDocumentMitraAttributeName      string     `gorm:"column:prod_doc_mitra_attr_name;size:255" json:"prod_doc_mitra_attr_name"`
	ProductDocumentMitraAttributeDesc      string     `gorm:"column:prod_doc_mitra_attr_desc" json:"prod_doc_mitra_attr_desc"`
	ProductDocumentMitraAttributeType      string     `gorm:"column:prod_doc_mitra_attr_type;size:255" json:"prod_doc_mitra_attr_type"`
	ProductDocumentMitraAttributeStatus    uint64     `gorm:"column:prod_doc_mitra_attr_status;size:2" json:"prod_doc_mitra_attr_status"`
	ProductDocumentMitraAttributeCreatedBy string     `gorm:"column:prod_doc_mitra_attr_created_by;size:125" json:"prod_doc_mitra_attr_created_by"`
	ProductDocumentMitraAttributeUpdatedBy string     `gorm:"column:prod_doc_mitra_attr_updated_by;size:125" json:"prod_doc_mitra_attr_updated_by"`
	ProductDocumentMitraAttributeDeletedBy string     `gorm:"column:prod_doc_mitra_attr_deleted_by;size:125" json:"prod_doc_mitra_attr_deleted_by"`
	ProductDocumentMitraAttributeCreatedAt time.Time  `gorm:"column:prod_doc_mitra_attr_created_at;default:CURRENT_TIMESTAMP" json:"prod_doc_mitra_attr_created_at"`
	ProductDocumentMitraAttributeUpdatedAt *time.Time `gorm:"column:prod_doc_mitra_attr_updated_at;default:CURRENT_TIMESTAMP" json:"prod_doc_mitra_attr_updated_at"`
	ProductDocumentMitraAttributeDeletedAt *time.Time `gorm:"column:prod_doc_mitra_attr_deleted_at;default:CURRENT_TIMESTAMP" json:"prod_doc_mitra_attr_deleted_at"`
}

type ProductDocumentMitraAttributeData struct {
	ProductDocumentMitraAttributeID        uint64     `gorm:"column:prod_doc_mitra_attr_id;primary_key;" json:"prod_doc_mitra_attr_id"`
	ProductDocumentMitraAttributeCode      string     `gorm:"column:prod_doc_mitra_attr_code;size:15" json:"prod_doc_mitra_attr_code"`
	ProductDocumentMitraAttributeName      string     `gorm:"column:prod_doc_mitra_attr_name;size:255" json:"prod_doc_mitra_attr_name"`
	ProductDocumentMitraAttributeDesc      string     `gorm:"column:prod_doc_mitra_attr_desc" json:"prod_doc_mitra_attr_desc"`
	ProductDocumentMitraAttributeType      string     `gorm:"column:prod_doc_mitra_attr_type;size:255" json:"prod_doc_mitra_attr_type"`
	ProductDocumentMitraAttributeStatus    uint64     `gorm:"column:prod_doc_mitra_attr_status;size:2" json:"prod_doc_mitra_attr_status"`
	ProductDocumentMitraID                 uint64     `gorm:"column:prod_doc_mitra_id" json:"prod_doc_mitra_id"`
	ProductDocumentMitraCode               string     `gorm:"column:prod_doc_mitra_code;size:25" json:"prod_doc_mitra_code"`
	ProductDocumentMitraName               string     `gorm:"column:prod_doc_mitra_name;size:255" json:"prod_doc_mitra_name"`
	ProductDocumentMitraStatus             uint64     `gorm:"column:prod_doc_mitra_status;size:25" json:"prod_doc_mitra_status"`
	ProductDocumentMitraAttributeCreatedBy string     `gorm:"column:prod_doc_mitra_attr_created_by;size:125" json:"prod_doc_mitra_attr_created_by"`
	ProductDocumentMitraAttributeUpdatedBy string     `gorm:"column:prod_doc_mitra_attr_updated_by;size:125" json:"prod_doc_mitra_attr_updated_by"`
	ProductDocumentMitraAttributeDeletedBy string     `gorm:"column:prod_doc_mitra_attr_deleted_by;size:125" json:"prod_doc_mitra_attr_deleted_by"`
	ProductDocumentMitraAttributeCreatedAt time.Time  `gorm:"column:prod_doc_mitra_attr_created_at;default:CURRENT_TIMESTAMP" json:"prod_doc_mitra_attr_created_at"`
	ProductDocumentMitraAttributeUpdatedAt *time.Time `gorm:"column:prod_doc_mitra_attr_updated_at;default:CURRENT_TIMESTAMP" json:"prod_doc_mitra_attr_updated_at"`
	ProductDocumentMitraAttributeDeletedAt *time.Time `gorm:"column:prod_doc_mitra_attr_deleted_at;default:CURRENT_TIMESTAMP" json:"prod_doc_mitra_attr_deleted_at"`
}

type ResponseProducDocumentMitraAttributes struct {
	Status  int                                  `json:"status"`
	Message string                               `json:"message"`
	Data    *[]ProductDocumentMitraAttributeData `json:"data"`
}

type ResponseProducDocumentMitraAttribute struct {
	Status  int                                `json:"status"`
	Message string                             `json:"message"`
	Data    *ProductDocumentMitraAttributeData `json:"data"`
}

type ResponseProducDocumentMitraAttributeIU struct {
	Status  int                            `json:"status"`
	Message string                         `json:"message"`
	Data    *ProductDocumentMitraAttribute `json:"data"`
}

type ResponseProducDocumentMitraAttributeDel struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (ProductDocumentMitraAttribute) TableName() string {
	return "m_product_document_mitra_attribute"
}

func (ProductDocumentMitraAttributeData) TableName() string {
	return "m_product_document_mitra_attribute"
}

func (p *ProductDocumentMitraAttribute) Prepare() {
	p.ProductDocumentMitraAttributeID = p.ProductDocumentMitraAttributeID
	p.ProductDocumentMitraID = p.ProductDocumentMitraID
	p.ProductDocumentMitraAttributeCode = p.ProductDocumentMitraAttributeCode
	p.ProductDocumentMitraAttributeName = p.ProductDocumentMitraAttributeName
	p.ProductDocumentMitraAttributeDesc = p.ProductDocumentMitraAttributeDesc
	p.ProductDocumentMitraAttributeType = p.ProductDocumentMitraAttributeType
	p.ProductDocumentMitraAttributeStatus = p.ProductDocumentMitraAttributeStatus
	p.ProductDocumentMitraAttributeCreatedBy = p.ProductDocumentMitraAttributeCreatedBy
	p.ProductDocumentMitraAttributeUpdatedBy = p.ProductDocumentMitraAttributeUpdatedBy
	p.ProductDocumentMitraAttributeDeletedBy = p.ProductDocumentMitraAttributeDeletedBy
	p.ProductDocumentMitraAttributeCreatedAt = time.Now()
	p.ProductDocumentMitraAttributeUpdatedAt = p.ProductDocumentMitraAttributeUpdatedAt
	p.ProductDocumentMitraAttributeDeletedAt = p.ProductDocumentMitraAttributeDeletedAt
}

func (p *ProductDocumentMitraAttribute) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.ProductDocumentMitraAttributeID == 0 {
			return errors.New("Required Product Document Mitra Attribute Code")
		}
		if p.ProductDocumentMitraID == 0 {
			return errors.New("Required Product Document Mitra Attribute Name")
		}
		return nil
	}
}

func (p *ProductDocumentMitraAttribute) SaveProductgoryDocumentMitraAttribute(db *gorm.DB) (*ProductDocumentMitraAttribute, error) {
	var err error
	err = db.Debug().Model(&ProductDocumentMitraAttribute{}).Create(&p).Error
	if err != nil {
		return &ProductDocumentMitraAttribute{}, err
	}
	return p, nil
}

func (p *ProductDocumentMitraAttribute) FindAllProductategoryDocumentMitraAttribute(db *gorm.DB) (*[]ProductDocumentMitraAttributeData, error) {
	var err error
	productdocumentmitraattribute := []ProductDocumentMitraAttributeData{}
	err = db.Debug().Model(&ProductDocumentMitraAttributeData{}).Limit(100).
		Select(`m_product_document_mitra_attribute.prod_doc_attr_id,
				m_product_document_mitra_attribute.prod_doc_attr_code,
				m_product_document_mitra_attribute.prod_doc_attr_name,
				m_product_document_mitra.prod_doc_mitra_id,
				m_product_document_mitra.prod_doc_mitra_code,
				m_product_document_mitra.prod_doc_mitra_name,
				m_product_document_mitra_attribute.prod_doc_attr_desc,
				m_product_document_mitra_attribute.prod_doc_attr_type,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_status,
				m_product_document_mitra_attribute.prod_doc_attr_created_by,
				m_product_document_mitra_attribute.prod_doc_attr_updated_by,
				m_product_document_mitra_attribute.prod_doc_attr_deleted_by,
				m_product_document_mitra_attribute.prod_doc_attr_created_at,
				m_product_document_mitra_attribute.prod_doc_attr_updated_at,
				m_product_document_mitra_attribute.prod_doc_attr_deleted_at
				`).
		// Joins("JOIN m_product_document on m_product_document_Mitra_attribute.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_document_mitra on m_product_document_mitra_attribute.prod_doc_mitra_id=m_product_document_mitra.prod_doc_mitra_id").
		Find(&productdocumentmitraattribute).Error
	if err != nil {
		return &[]ProductDocumentMitraAttributeData{}, err
	}
	return &productdocumentmitraattribute, nil
}

func (p *ProductDocumentMitraAttribute) FindAllProductategoryDocumentMitraAttributeStatus(db *gorm.DB, status uint64) (*[]ProductDocumentMitraAttributeData, error) {
	var err error
	productdocumentmitraattribute := []ProductDocumentMitraAttributeData{}
	err = db.Debug().Model(&ProductDocumentMitraAttributeData{}).Limit(100).
		Select(`m_product_document_mitra_attribute.prod_doc_attr_id,
				m_product_document_mitra_attribute.prod_doc_attr_code,
				m_product_document_mitra_attribute.prod_doc_attr_name,
				m_product_document_mitra.prod_doc_mitra_id,
				m_product_document_mitra.prod_doc_mitra_code,
				m_product_document_mitra.prod_doc_mitra_name,
				m_product_document_mitra_attribute.prod_doc_attr_desc,
				m_product_document_mitra_attribute.prod_doc_attr_type,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_status,
				m_product_document_mitra_attribute.prod_doc_attr_created_by,
				m_product_document_mitra_attribute.prod_doc_attr_updated_by,
				m_product_document_mitra_attribute.prod_doc_attr_deleted_by,
				m_product_document_mitra_attribute.prod_doc_attr_created_at,
				m_product_document_mitra_attribute.prod_doc_attr_updated_at,
				m_product_document_mitra_attribute.prod_doc_attr_deleted_at
				`).
		// Joins("JOIN m_product_document on m_product_document_mitra_attribute.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_document_mitra on m_product_document_mitra_attribute.prod_doc_mitra_id=m_product_document_mitra.prod_doc_mitra_id").
		Where("prod_doc_mitra_attr_status = ?", status).
		Find(&productdocumentmitraattribute).Error
	if err != nil {
		return &[]ProductDocumentMitraAttributeData{}, err
	}
	return &productdocumentmitraattribute, nil
}

func (p *ProductDocumentMitraAttribute) FindProductgoryDocumentMitraAttributeByID(db *gorm.DB, pid uint64) (*ProductDocumentMitraAttributeData, error) {
	var err error
	productdocumentmitraattribute := ProductDocumentMitraAttributeData{}
	err = db.Debug().Model(&ProductDocumentMitraAttributeData{}).Limit(100).
		Select(`m_product_document_mitra_attribute.prod_doc_attr_id,
				m_product_document_mitra_attribute.prod_doc_attr_code,
				m_product_document_mitra_attribute.prod_doc_attr_name,
				m_product_document_mitra.prod_doc_mitra_id,
				m_product_document_mitra.prod_doc_mitra_code,
				m_product_document_mitra.prod_doc_mitra_name,
				m_product_document_mitra_attribute.prod_doc_attr_desc,
				m_product_document_mitra_attribute.prod_doc_attr_type,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_status,
				m_product_document_mitra_attribute.prod_doc_attr_created_by,
				m_product_document_mitra_attribute.prod_doc_attr_updated_by,
				m_product_document_mitra_attribute.prod_doc_attr_deleted_by,
				m_product_document_mitra_attribute.prod_doc_attr_created_at,
				m_product_document_mitra_attribute.prod_doc_attr_updated_at,
				m_product_document_mitra_attribute.prod_doc_attr_deleted_at
				`).
		// Joins("JOIN m_product_document on m_product_document_mitra_attribute.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_document_mitra on m_product_document_mitra_attribute.prod_doc_mitra_id=m_product_document_mitra.prod_doc_mitra_id").
		Where("prod_doc_attr_id = ?", pid).
		Find(&productdocumentmitraattribute).Error
	if err != nil {
		return &ProductDocumentMitraAttributeData{}, err
	}
	return &productdocumentmitraattribute, nil
}

func (p *ProductDocumentMitraAttribute) FindProductgoryDocumentMitraAttributeStatusByID(db *gorm.DB, pid uint64, status uint64) (*ProductDocumentMitraAttributeData, error) {
	var err error
	productdocumentmitraattribute := ProductDocumentMitraAttributeData{}
	err = db.Debug().Model(&ProductDocumentMitraAttributeData{}).Limit(100).
		Select(`m_product_document_mitra_attribute.prod_doc_attr_id,
				m_product_document_mitra_attribute.prod_doc_attr_code,
				m_product_document_mitra_attribute.prod_doc_attr_name,
				m_product_document_mitra.prod_doc_mitra_id,
				m_product_document_mitra.prod_doc_mitra_code,
				m_product_document_mitra.prod_doc_mitra_name,
				m_product_document_mitra_attribute.prod_doc_attr_desc,
				m_product_document_mitra_attribute.prod_doc_attr_type,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_status,
				m_product_document_mitra_attribute.prod_doc_attr_created_by,
				m_product_document_mitra_attribute.prod_doc_attr_updated_by,
				m_product_document_mitra_attribute.prod_doc_attr_deleted_by,
				m_product_document_mitra_attribute.prod_doc_attr_created_at,
				m_product_document_mitra_attribute.prod_doc_attr_updated_at,
				m_product_document_mitra_attribute.prod_doc_attr_deleted_at
				`).
		// Joins("JOIN m_product_document on m_product_document_mitra_attribute.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_document_mitra on m_product_document_mitra_attribute.prod_doc_mitra_id=m_product_document_mitra.prod_doc_mitra_id").
		Where("prod_doc_attr_id = ? AND prod_doc_mitra_attr_status = ?", pid, status).
		Find(&productdocumentmitraattribute).Error
	if err != nil {
		return &ProductDocumentMitraAttributeData{}, err
	}
	return &productdocumentmitraattribute, nil
}

func (p *ProductDocumentMitraAttribute) FindProductgoryDocumentMitraAttributeBySubCategoryByID(db *gorm.DB, pid uint64, subcategory uint64) (*ProductDocumentMitraAttributeData, error) {
	var err error
	productdocumentmitraattribute := ProductDocumentMitraAttributeData{}
	err = db.Debug().Model(&ProductDocumentMitraAttributeData{}).Limit(100).
		Select(`m_product_document_mitra_attribute.prod_doc_attr_id,
				m_product_document_mitra_attribute.prod_doc_attr_code,
				m_product_document_mitra_attribute.prod_doc_attr_name,
				m_product_document_mitra.prod_doc_mitra_id,
				m_product_document_mitra.prod_doc_mitra_code,
				m_product_document_mfitra.prod_doc_mitra_name,
				m_product_document_mitra_attribute.prod_doc_attr_desc,
				m_product_document_mitra_attribute.prod_doc_attr_type,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_status,
				m_product_document_mitra_attribute.prod_doc_attr_created_by,
				m_product_document_mitra_attribute.prod_doc_attr_updated_by,
				m_product_document_mitra_attribute.prod_doc_attr_deleted_by,
				m_product_document_mitra_attribute.prod_doc_attr_created_at,
				m_product_document_mitra_attribute.prod_doc_attr_updated_at,
				m_product_document_mitra_attribute.prod_doc_attr_deleted_at
				`).
		// Joins("JOIN m_product_document on m_product_document_mitra_attribute.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_document_mitra on m_product_document_mitra_attribute.prod_doc_mitra_id=m_product_document_mitra.prod_doc_mitra_id").
		Where("prod_doc_attr_id = ? AND prod_doc_mitra_attr_status = ?", pid, subcategory).
		Find(&productdocumentmitraattribute).Error
	if err != nil {
		return &ProductDocumentMitraAttributeData{}, err
	}
	return &productdocumentmitraattribute, nil
}

func (p *ProductDocumentMitraAttribute) UpdateProductDocumentMitraAttribute(db *gorm.DB, pid uint64) (*ProductDocumentMitraAttribute, error) {
	var err error
	db = db.Debug().Model(&ProductDocumentMitraAttribute{}).Where("prod_doc_attr_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_mitra_id":          p.ProductDocumentMitraID,
			"prod_doc_mitra_attr_code":   p.ProductDocumentMitraAttributeCode,
			"prod_doc_mitra_attr_name":   p.ProductDocumentMitraAttributeName,
			"prod_doc_mitra_attr_desc":   p.ProductDocumentMitraAttributeDesc,
			"prod_doc_mitra_attr_type":   p.ProductDocumentMitraAttributeType,
			"prod_doc_attr_updated_by":   p.ProductDocumentMitraAttributeUpdatedBy,
			"prod_doc_mitra_attr_status": time.Now(),
		},
	)
	err = db.Debug().Model(&ProductDocumentMitraAttribute{}).Where("prod_doc_attr_id = ?", pid).Error
	if err != nil {
		return &ProductDocumentMitraAttribute{}, err
	}
	return p, nil
}

func (p *ProductDocumentMitraAttribute) UpdateProducttegoryDocumentMitraAttributeStatus(db *gorm.DB, pid uint64) (*ProductDocumentMitraAttribute, error) {
	var err error
	db = db.Debug().Model(&ProductDocumentMitraAttribute{}).Where("prod_doc_attr_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_mitra_attr_status":     p.ProductDocumentMitraAttributeStatus,
			"prod_doc_mitra_attr_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ProductDocumentMitraAttribute{}).Where("prod_doc_attr_id = ?", pid).Error
	if err != nil {
		return &ProductDocumentMitraAttribute{}, err
	}
	return p, nil
}

func (p *ProductDocumentMitraAttribute) UpdateProducttegoryDocumentMitraAttributeStatusDelete(db *gorm.DB, pid uint64) (*ProductDocumentMitraAttribute, error) {
	var err error
	db = db.Debug().Model(&ProductDocumentMitraAttribute{}).Where("prod_doc_attr_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_mitra_attr_status":     3,
			"prod_doc_mitra_attr_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ProductDocumentMitraAttribute{}).Where("prod_doc_attr_id = ?", pid).Error
	if err != nil {
		return &ProductDocumentMitraAttribute{}, err
	}
	return p, nil
}

func (p *ProductDocumentMitraAttribute) DeleteProducttegoryDocumentMitraAttribute(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&ProductDocumentMitraAttribute{}).Where("prod_doc_attr_id = ?", pid).Take(&ProductDocumentMitraAttribute{}).Delete(&ProductDocumentMitraAttribute{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Product Document Mitra Attribute not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
