package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type ProductDocumentAttribute struct {
	ProductDocumentAttributeID        uint64     `gorm:"column:prod_doc_attr_id;primary_key;" json:"prod_doc_attr_id"`
	ProductDocumentAttributeCode      string     `gorm:"column:prod_doc_attr_code;size:25" json:"prod_doc_attr_code"`
	ProductDocumentAttributeName      string     `gorm:"column:prod_doc_attr_name;size:255" json:"prod_doc_attr_name"`
	ProductDocumentAttributeStatus    uint64     `gorm:"column:prod_doc_attr_status;size:25" json:"prod_doc_attr_status"`
	ProductDocumentAttributeCreatedBy string     `gorm:"column:prod_doc_attr_created_by;size:125" json:"prod_doc_attr_created_by"`
	ProductDocumentAttributeUpdatedBy string     `gorm:"column:prod_doc_attr_updated_by;size:125" json:"prod_doc_attr_updated_by"`
	ProductDocumentAttributeDeletedBy string     `gorm:"column:prod_doc_attr_deleted_by;size:125" json:"prod_doc_attr_deleted_by"`
	ProductDocumentAttributeCreatedAt time.Time  `gorm:"column:prod_doc_attr_created_at;default:CURRENT_TIMESTAMP" json:"prod_doc_attr_created_at"`
	ProductDocumentAttributeUpdatedAt *time.Time `gorm:"column:prod_doc_attr_updated_at;default:CURRENT_TIMESTAMP" json:"prod_doc_attr_updated_at"`
	ProductDocumentAttributeDeletedAt *time.Time `gorm:"column:prod_doc_attr_deleted_at;default:CURRENT_TIMESTAMP" json:"prod_doc_attr_deleted_at"`
}

type ResponseProductDocumentAttributes struct {
	Status  int                         `json:"status"`
	Message string                      `json:"message"`
	Data    *[]ProductDocumentAttribute `json:"data"`
}

type ResponseProductDocumentAttribute struct {
	Status  int                       `json:"status"`
	Message string                    `json:"message"`
	Data    *ProductDocumentAttribute `json:"data"`
}

type ResponseProductDocumentAttributeIU struct {
	Status  int                       `json:"status"`
	Message string                    `json:"message"`
	Data    *ProductDocumentAttribute `json:"data"`
}

type ResponseProductDocumentAttributeDel struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (ProductDocumentAttribute) TableName() string {
	return "m_product_document_attribute"
}

func (p *ProductDocumentAttribute) Prepare() {
	p.ProductDocumentAttributeCode = html.EscapeString(strings.TrimSpace(p.ProductDocumentAttributeCode))
	p.ProductDocumentAttributeName = html.EscapeString(strings.TrimSpace(p.ProductDocumentAttributeName))
	p.ProductDocumentAttributeStatus = p.ProductDocumentAttributeStatus
	p.ProductDocumentAttributeCreatedBy = html.EscapeString(strings.TrimSpace(p.ProductDocumentAttributeCreatedBy))
	p.ProductDocumentAttributeUpdatedBy = html.EscapeString(strings.TrimSpace(p.ProductDocumentAttributeUpdatedBy))
	p.ProductDocumentAttributeDeletedBy = html.EscapeString(strings.TrimSpace(p.ProductDocumentAttributeDeletedBy))
	p.ProductDocumentAttributeCreatedAt = time.Now()
	p.ProductDocumentAttributeUpdatedAt = p.ProductDocumentAttributeUpdatedAt
	p.ProductDocumentAttributeDeletedAt = p.ProductDocumentAttributeDeletedAt
}

func (p *ProductDocumentAttribute) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.ProductDocumentAttributeCode == "" {
			return errors.New("Required ProductDocumentAttribute Code")
		}
		if p.ProductDocumentAttributeName == "" {
			return errors.New("Required ProductDocumentAttribute Name")
		}
		return nil
	}
}

func (p *ProductDocumentAttribute) SaveProductDocumentAttribute(db *gorm.DB) (*ProductDocumentAttribute, error) {
	var err error
	err = db.Debug().Model(&ProductDocumentAttribute{}).Create(&p).Error
	if err != nil {
		return &ProductDocumentAttribute{}, err
	}
	return p, nil
}

func (p *ProductDocumentAttribute) FindAllProductDocumentAttribute(db *gorm.DB) (*[]ProductDocumentAttribute, error) {
	var err error
	productdocumentattribute := []ProductDocumentAttribute{}
	err = db.Debug().Model(&ProductDocumentAttribute{}).Limit(100).Find(&productdocumentattribute).Error
	if err != nil {
		return &[]ProductDocumentAttribute{}, err
	}
	return &productdocumentattribute, nil
}

func (p *ProductDocumentAttribute) FindAllProductDocumentAttributeStatus(db *gorm.DB, status uint64) (*[]ProductDocumentAttribute, error) {
	var err error
	productdocumentattribute := []ProductDocumentAttribute{}
	err = db.Debug().Model(&ProductDocumentAttribute{}).Where("prod_doc_attr_status = ?", status).Limit(100).Find(&productdocumentattribute).Error
	if err != nil {
		return &[]ProductDocumentAttribute{}, err
	}
	return &productdocumentattribute, nil
}

func (p *ProductDocumentAttribute) FindProductDocumentAttributeByID(db *gorm.DB, pid uint64) (*ProductDocumentAttribute, error) {
	var err error
	err = db.Debug().Model(&ProductDocumentAttribute{}).Where("prod_doc_attr_id = ?", pid).Take(&p).Error
	if err != nil {
		return &ProductDocumentAttribute{}, err
	}
	return p, nil
}

func (p *ProductDocumentAttribute) FindProductDocumentAttributeStatusByID(db *gorm.DB, pid uint64, status uint64) (*ProductDocumentAttribute, error) {
	var err error
	err = db.Debug().Model(&ProductDocumentAttribute{}).Where("prod_doc_attr_id = ? AND prod_doc_attr_status = ?", pid, status).Take(&p).Error
	if err != nil {
		return &ProductDocumentAttribute{}, err
	}
	return p, nil
}

func (p *ProductDocumentAttribute) UpdateProductDocumentAttribute(db *gorm.DB, pid uint64) (*ProductDocumentAttribute, error) {
	var err error
	db = db.Debug().Model(&ProductDocumentAttribute{}).Where("prod_doc_attr_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_attr_code":       p.ProductDocumentAttributeCode,
			"prod_doc_attr_name":       p.ProductDocumentAttributeName,
			"prod_doc_attr_status":     p.ProductDocumentAttributeStatus,
			"prod_doc_attr_updated_by": p.ProductDocumentAttributeUpdatedBy,
			"prod_doc_attr_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ProductDocumentAttribute{}).Where("prod_doc_attr_id = ?", pid).Error
	if err != nil {
		return &ProductDocumentAttribute{}, err
	}
	return p, nil
}

func (p *ProductDocumentAttribute) UpdateProductDocumentAttributeStatus(db *gorm.DB, pid uint64) (*ProductDocumentAttribute, error) {
	var err error
	db = db.Debug().Model(&ProductDocumentAttribute{}).Where("prod_doc_attr_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_attr_status":     p.ProductDocumentAttributeStatus,
			"prod_doc_attr_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ProductDocumentAttribute{}).Where("prod_doc_attr_id = ?", pid).Error
	if err != nil {
		return &ProductDocumentAttribute{}, err
	}
	return p, nil
}

func (p *ProductDocumentAttribute) UpdateProductDocumentAttributeStatusDelete(db *gorm.DB, pid uint64) (*ProductDocumentAttribute, error) {
	var err error
	db = db.Debug().Model(&ProductDocumentAttribute{}).Where("prod_doc_attr_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_attr_status":     3,
			"prod_doc_attr_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ProductDocumentAttribute{}).Where("prod_doc_attr_id = ?", pid).Error
	if err != nil {
		return &ProductDocumentAttribute{}, err
	}
	return p, nil
}

func (p *ProductDocumentAttribute) DeleteProductDocumentAttribute(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&ProductDocumentAttribute{}).Where("prod_doc_attr_id = ?", pid).Take(&ProductDocumentAttribute{}).Delete(&ProductDocumentAttribute{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Product Document Attribute not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
