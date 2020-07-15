package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type ProductDocumentCustomer struct {
	ProductDocumentCustomerID        uint64     `gorm:"column:prod_doc_cust_id;primary_key;" json:"prod_doc_cust_id"`
	ProductDocumentCustomerCode      string     `gorm:"column:prod_doc_cust_code;size:25" json:"prod_doc_cust_code"`
	ProductDocumentCustomerName      string     `gorm:"column:prod_doc_cust_name;size:255" json:"prod_doc_cust_name"`
	ProductDocumentCustomerStatus    uint64     `gorm:"column:prod_doc_cust_status;size:25" json:"prod_doc_cust_status"`
	ProductDocumentCustomerCreatedBy string     `gorm:"column:prod_doc_cust_created_by;size:125" json:"prod_doc_cust_created_by"`
	ProductDocumentCustomerUpdatedBy string     `gorm:"column:prod_doc_cust_updated_by;size:125" json:"prod_doc_cust_updated_by"`
	ProductDocumentCustomerDeletedBy string     `gorm:"column:prod_doc_cust_deleted_by;size:125" json:"prod_doc_cust_deleted_by"`
	ProductDocumentCustomerCreatedAt time.Time  `gorm:"column:prod_doc_cust_created_at;default:CURRENT_TIMESTAMP" json:"prod_doc_cust_created_at"`
	ProductDocumentCustomerUpdatedAt *time.Time `gorm:"column:prod_doc_cust_updated_at;default:CURRENT_TIMESTAMP" json:"prod_doc_cust_updated_at"`
	ProductDocumentCustomerDeletedAt *time.Time `gorm:"column:prod_doc_cust_deleted_at;default:CURRENT_TIMESTAMP" json:"prod_doc_cust_deleted_at"`
}

type ResponseProductDocumentCustomers struct {
	Status  int                        `json:"status"`
	Message string                     `json:"message"`
	Data    *[]ProductDocumentCustomer `json:"data"`
}

type ResponseProductDocumentCustomer struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    *ProductDocumentCustomer `json:"data"`
}

type ResponseProductDocumentCustomerIU struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    *ProductDocumentCustomer `json:"data"`
}

type ResponseProductDocumentCustomerDel struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (ProductDocumentCustomer) TableName() string {
	return "m_product_document_customer"
}

func (p *ProductDocumentCustomer) Prepare() {
	p.ProductDocumentCustomerCode = html.EscapeString(strings.TrimSpace(p.ProductDocumentCustomerCode))
	p.ProductDocumentCustomerName = html.EscapeString(strings.TrimSpace(p.ProductDocumentCustomerName))
	p.ProductDocumentCustomerStatus = p.ProductDocumentCustomerStatus
	p.ProductDocumentCustomerCreatedAt = time.Now()
	p.ProductDocumentCustomerUpdatedAt = p.ProductDocumentCustomerUpdatedAt
	p.ProductDocumentCustomerDeletedAt = p.ProductDocumentCustomerDeletedAt
}

func (p *ProductDocumentCustomer) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.ProductDocumentCustomerCode == "" {
			return errors.New("Required ProductDocumentCustomer Code")
		}
		if p.ProductDocumentCustomerName == "" {
			return errors.New("Required ProductDocumentCustomer Name")
		}
		return nil
	}
}

func (p *ProductDocumentCustomer) SaveProductDocumentCustomer(db *gorm.DB) (*ProductDocumentCustomer, error) {
	var err error
	err = db.Debug().Model(&ProductDocumentCustomer{}).Create(&p).Error
	if err != nil {
		return &ProductDocumentCustomer{}, err
	}
	return p, nil
}

func (p *ProductDocumentCustomer) FindAllProductDocumentCustomer(db *gorm.DB) (*[]ProductDocumentCustomer, error) {
	var err error
	productdocumentcustomer := []ProductDocumentCustomer{}
	err = db.Debug().Model(&ProductDocumentCustomer{}).Limit(100).Find(&productdocumentcustomer).Error
	if err != nil {
		return &[]ProductDocumentCustomer{}, err
	}
	return &productdocumentcustomer, nil
}

func (p *ProductDocumentCustomer) FindAllProductDocumentCustomerStatus(db *gorm.DB, status uint64) (*[]ProductDocumentCustomer, error) {
	var err error
	productdocumentcustomer := []ProductDocumentCustomer{}
	err = db.Debug().Model(&ProductDocumentCustomer{}).Where("prod_doc_cust_status = ?", status).Limit(100).Find(&productdocumentcustomer).Error
	if err != nil {
		return &[]ProductDocumentCustomer{}, err
	}
	return &productdocumentcustomer, nil
}

func (p *ProductDocumentCustomer) FindProductDocumentCustomerByID(db *gorm.DB, pid uint64) (*ProductDocumentCustomer, error) {
	var err error
	err = db.Debug().Model(&ProductDocumentCustomer{}).Where("prod_doc_cust_id = ?", pid).Take(&p).Error
	if err != nil {
		return &ProductDocumentCustomer{}, err
	}
	return p, nil
}

func (p *ProductDocumentCustomer) FindProductDocumentCustomerStatusByID(db *gorm.DB, pid uint64, status uint64) (*ProductDocumentCustomer, error) {
	var err error
	err = db.Debug().Model(&ProductDocumentCustomer{}).Where("prod_doc_cust_id = ? AND prod_doc_cust_status = ?", pid, status).Take(&p).Error
	if err != nil {
		return &ProductDocumentCustomer{}, err
	}
	return p, nil
}

func (p *ProductDocumentCustomer) UpdateProductDocumentCustomer(db *gorm.DB, pid uint64) (*ProductDocumentCustomer, error) {
	var err error
	db = db.Debug().Model(&ProductDocumentCustomer{}).Where("prod_doc_cust_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_cust_code":       p.ProductDocumentCustomerCode,
			"prod_doc_cust_name":       p.ProductDocumentCustomerName,
			"prod_doc_cust_status":     p.ProductDocumentCustomerStatus,
			"prod_doc_cust_updated_by": p.ProductDocumentCustomerUpdatedBy,
			"prod_doc_cust_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ProductDocumentCustomer{}).Where("prod_doc_cust_id = ?", pid).Error
	if err != nil {
		return &ProductDocumentCustomer{}, err
	}
	return p, nil
}

func (p *ProductDocumentCustomer) UpdateProductDocumentCustomerStatus(db *gorm.DB, pid uint64) (*ProductDocumentCustomer, error) {
	var err error
	db = db.Debug().Model(&ProductDocumentCustomer{}).Where("prod_doc_cust_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_cust_status":     p.ProductDocumentCustomerStatus,
			"prod_doc_cust_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ProductDocumentCustomer{}).Where("prod_doc_cust_id = ?", pid).Error
	if err != nil {
		return &ProductDocumentCustomer{}, err
	}
	return p, nil
}

func (p *ProductDocumentCustomer) UpdateProductDocumentCustomerStatusDelete(db *gorm.DB, pid uint64) (*ProductDocumentCustomer, error) {
	var err error
	db = db.Debug().Model(&ProductDocumentCustomer{}).Where("prod_doc_cust_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_cust_status":     3,
			"prod_doc_cust_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ProductDocumentCustomer{}).Where("prod_doc_cust_id = ?", pid).Error
	if err != nil {
		return &ProductDocumentCustomer{}, err
	}
	return p, nil
}

func (p *ProductDocumentCustomer) DeleteProductDocumentCustomer(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&ProductDocumentCustomer{}).Where("prod_doc_cust_id = ?", pid).Take(&ProductDocumentCustomer{}).Delete(&ProductDocumentCustomer{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("ProductDocumentCustomer not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
