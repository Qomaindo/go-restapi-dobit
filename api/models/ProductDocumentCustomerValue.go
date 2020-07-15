package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type ProductDocumentCustomerValue struct {
	ProductDocumentCustomerValueID        uint64     `gorm:"column:prod_doc_cust_val_id;primary_key;" json:"prod_doc_cust_val_id"`
	ProductDocumentCustomerID             uint64     `gorm:"column:prod_doc_cust_id" json:"prod_doc_cust_id"`
	ProductDocumentCustomerValue          string     `gorm:"column:prod_doc_cust_val_value;size:255" json:"prod_doc_cust_val_value"`
	ProductDocumentCustomerValueCreatedBy string     `gorm:"column:prod_doc_cust_val_created_by;size:125" json:"prod_doc_cust_val_created_by"`
	ProductDocumentCustomerValueCreatedAt time.Time  `gorm:"column:prod_doc_cust_id;default:CURRENT_TIMESTAMP" json:"prod_doc_cust_id"`
	ProductDocumentCustomerValueUpdatedBy string     `gorm:"column:prod_doc_cust_val_updated_by;size:125" json:"prod_doc_cust_val_updated_by"`
	ProductDocumentCustomerValueUpdatedAt *time.Time `gorm:"column:prod_doc_cust_id;default:CURRENT_TIMESTAMP" json:"prod_doc_cust_id"`
	ProductDocumentCustomerValueDeletedBy string     `gorm:"column:prod_doc_cust_val_deleted_by;size:125" json:"prod_doc_cust_val_deleted_by"`
	ProductDocumentCustomerValueDeletedAt *time.Time `gorm:"column:prod_doc_cust_id;default:CURRENT_TIMESTAMP" json:"prod_doc_cust_id"`
}

type ProductDocumentCustomerValueData struct {
	ProductDocumentCustomerValueID        uint64     `gorm:"column:prod_doc_cust_val_id;primary_key;" json:"prod_doc_cust_val_id"`
	ProductDocumentCustomerID             uint64     `gorm:"column:prod_doc_cust_id;" json:"prod_doc_cust_id"`
	ProductDocumentCustomerValue          string     `gorm:"column:prod_doc_cust_id;size:255" json:"prod_doc_cust_id"`
	ProductDocumentCustomerValueCreatedBy string     `gorm:"column:prod_doc_cust_val_created_by;size:125" json:"prod_doc_cust_val_created_by"`
	ProductDocumentCustomerValueCreatedAt time.Time  `gorm:"column:prod_doc_cust_id;default:CURRENT_TIMESTAMP" json:"prod_doc_cust_id"`
	ProductDocumentCustomerValueUpdatedBy string     `gorm:"column:prod_doc_cust_val_updated_by;size:125" json:"prod_doc_cust_val_updated_by"`
	ProductDocumentCustomerValueUpdatedAt *time.Time `gorm:"column:prod_doc_cust_id;default:CURRENT_TIMESTAMP" json:"prod_doc_cust_id"`
	ProductDocumentCustomerValueDeletedBy string     `gorm:"column:prod_doc_cust_val_deleted_by;size:125" json:"prod_doc_cust_val_deleted_by"`
	ProductDocumentCustomerValueDeletedAt *time.Time `gorm:"column:prod_doc_cust_id;default:CURRENT_TIMESTAMP" json:"prod_doc_cust_id"`
}

type ResponseProductDocumentCustomerValues struct {
	Status  uint64                              `json:"status"`
	Message string                              `json:"message"`
	Data    *[]ProductDocumentCustomerValueData `json:"data"`
}

type ResponseProductDocumentCustomerValue struct {
	Status  uint64                        `json:"status"`
	Message string                        `json:"message"`
	Data    *ProductDocumentCustomerValue `json:"data"`
}
type ResponseProductDocumentCustomerValueData struct {
	Status  uint64                            `json:"status"`
	Message string                            `json:"message"`
	Data    *ProductDocumentCustomerValueData `json:"data"`
}
type ResponseProductDocumentCustomerValueIU struct {
	Status  uint64                        `json:"status"`
	Message string                        `json:"message"`
	Data    *ProductDocumentCustomerValue `json:"data"`
}

type ResponseProductDocumentCustomerValueDel struct {
	Status  uint64 `json:"status"`
	Message string `json:"message"`
}

func (ProductDocumentCustomerValue) TableName() string {
	return "m_product_document_customer_value"
}

func (ProductDocumentCustomerValueData) TableName() string {
	return "m_product_document_customer_value"
}

func (u *ProductDocumentCustomerValue) Prepare() {
	u.ProductDocumentCustomerValueID = u.ProductDocumentCustomerValueID
	u.ProductDocumentCustomerValue = html.EscapeString(strings.TrimSpace(u.ProductDocumentCustomerValue))
	u.ProductDocumentCustomerValueCreatedBy = u.ProductDocumentCustomerValueCreatedBy
	u.ProductDocumentCustomerValueCreatedAt = time.Now()
	u.ProductDocumentCustomerValueUpdatedBy = u.ProductDocumentCustomerValueUpdatedBy
	u.ProductDocumentCustomerValueUpdatedAt = u.ProductDocumentCustomerValueUpdatedAt
	u.ProductDocumentCustomerValueDeletedBy = u.ProductDocumentCustomerValueDeletedBy
	u.ProductDocumentCustomerValueDeletedAt = u.ProductDocumentCustomerValueDeletedAt
}

func (u *ProductDocumentCustomerValue) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		// if u.clientomerAddressID == 0 {
		// 	return errors.New("Required clientomer ID")
		// }
		// if u.clientomerAddressImageName == "" {
		// 	return errors.New("Required User Code")
		// }
		// if u.AddressID == 0 {
		// 	return errors.New("Required Email")
		// }
		// if u.clientomerAddressImageZipCode == "" {
		// 	return errors.New("Required Username")
		// }
		return nil
	}
}

func (u *ProductDocumentCustomerValue) SaveProductDocumentCustomerValue(db *gorm.DB) (*ProductDocumentCustomerValue, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &ProductDocumentCustomerValue{}, err
	}
	return u, nil
}

func (p *ProductDocumentCustomerValue) FindAllProductDocumentCustomerValue(db *gorm.DB) (*[]ProductDocumentCustomerValueData, error) {
	var err error
	productdocumentcustomerimage := []ProductDocumentCustomerValueData{}
	err = db.Debug().Model(&ProductDocumentCustomerValueData{}).Limit(100).
		Select(`m_product_document_customer_value.prod_doc_cust_val_id,
				m_product_document_customer_value.prod_doc_cust_id,
				m_product_document_customer_value.prod_doc_cust_val_value,
				m_product_document_customer_value.prod_doc_cust_val_created_by,
				m_product_document_customer_value.prod_doc_cust_val_updated_by,
				m_product_document_customer_value.prod_doc_cust_val_deleted_by,
				m_product_document_customer_value.prod_doc_cust_val_created_at,
				m_product_document_customer_value.prod_doc_cust_val_updated_at,
				m_product_document_customer_value.prod_doc_cust_val_deleted_at`).
		Joins("JOIN m_client_address on m_product_document_customer_value.prod_doc_cust_id=m_client_address.prod_doc_cust_id").
		Find(&productdocumentcustomerimage).Error
	if err != nil {
		return &[]ProductDocumentCustomerValueData{}, err
	}
	return &productdocumentcustomerimage, nil
}

func (p *ProductDocumentCustomerValue) FindAllProductDocumentCustomerValueStatus(db *gorm.DB, status uint64) (*[]ProductDocumentCustomerValueData, error) {
	var err error
	productdocumentcustomerimage := []ProductDocumentCustomerValueData{}
	err = db.Debug().Model(&ProductDocumentCustomerValueData{}).Limit(100).
		Select(`m_product_document_customer_value.prod_doc_cust_val_id,
				m_product_document_customer_value.prod_doc_cust_id,
				m_product_document_customer_value.prod_doc_cust_val_value,
				m_product_document_customer_value.prod_doc_cust_val_created_by,
				m_product_document_customer_value.prod_doc_cust_val_updated_by,
				m_product_document_customer_value.prod_doc_cust_val_deleted_by,
				m_product_document_customer_value.prod_doc_cust_val_created_at,
				m_product_document_customer_value.prod_doc_cust_val_updated_at,
				m_product_document_customer_value.prod_doc_cust_val_deleted_at`).
		Joins("JOIN m_client_address on m_product_document_customer_value.prod_doc_cust_id=m_client_address.prod_doc_cust_id").
		Where("client_adrs_img_status = ?", status).
		Find(&productdocumentcustomerimage).Error
	if err != nil {
		return &[]ProductDocumentCustomerValueData{}, err
	}
	return &productdocumentcustomerimage, nil
}

func (p *ProductDocumentCustomerValue) FindProductDocumentCustomerValueByIDAddress(db *gorm.DB, pid uint64) (*ProductDocumentCustomerValueData, error) {
	var err error
	productdocumentcustomerimage := ProductDocumentCustomerValueData{}
	err = db.Debug().Model(&ProductDocumentCustomerValueData{}).Limit(100).
		Select(`m_product_document_customer_value.prod_doc_cust_val_id,
				m_product_document_customer_value.prod_doc_cust_id,
				m_product_document_customer_value.prod_doc_cust_val_value,
				m_product_document_customer_value.prod_doc_cust_val_created_by,
				m_product_document_customer_value.prod_doc_cust_val_updated_by,
				m_product_document_customer_value.prod_doc_cust_val_deleted_by,
				m_product_document_customer_value.prod_doc_cust_val_created_at,
				m_product_document_customer_value.prod_doc_cust_val_updated_at,
				m_product_document_customer_value.prod_doc_cust_val_deleted_at`).
		// Joins("JOIN m_clientomer_address on m_clientomer_address_image.prod_doc_cust_id=m_clientomer_address.prod_doc_cust_id").
		Where("prod_doc_cust_val_id = ?", pid).
		Find(&productdocumentcustomerimage).Error
	if err != nil {
		return &ProductDocumentCustomerValueData{}, err
	}
	return &productdocumentcustomerimage, nil
}

func (p *ProductDocumentCustomerValue) FindProductDocumentCustomerValueStatusByID(db *gorm.DB, pid uint64, status uint64) (*ProductDocumentCustomerValueData, error) {
	var err error
	productdocumentcustomerimage := ProductDocumentCustomerValueData{}
	err = db.Debug().Model(&ProductDocumentCustomerValueData{}).Limit(100).
		Select(`m_product_document_customer_value.prod_doc_cust_val_id,
				m_client_address.prod_doc_cust_id,
				m_product_document_customer_value.client_adrs_img_name,
				m_product_document_customer_value.client_adrs_img_desc,
				m_product_document_customer_value.prod_doc_cust_id,
				m_product_document_customer_value.prod_doc_cust_id,
				m_product_document_customer_value.prod_doc_cust_id`).
		Joins("JOIN m_client_address on m_product_document_customer_value.prod_doc_cust_id=m_client_address.prod_doc_cust_id").
		Where("prod_doc_cust_val_id = ? AND client_adrs_img_status = ?", pid, status).
		Find(&productdocumentcustomerimage).Error
	if err != nil {
		return &ProductDocumentCustomerValueData{}, err
	}
	return &productdocumentcustomerimage, nil
}

func (p *ProductDocumentCustomerValue) FindProductDocumentCustomerValueByProductDocumentCustomerValueID(db *gorm.DB, pid uint64) (*ProductDocumentCustomerValueData, error) {
	var err error
	productdocumentcustomerimage := ProductDocumentCustomerValueData{}
	err = db.Debug().Model(&ProductDocumentCustomerValueData{}).Limit(100).
		Select(`m_product_document_customer_value.prod_doc_cust_val_id,
		m_product_document_customer_value.prod_doc_cust_id,
				m_product_document_customer_value.prod_doc_cust_val_value,
				m_product_document_customer_value.prod_doc_cust_val_created_by,
				m_product_document_customer_value.prod_doc_cust_val_updated_by,
				m_product_document_customer_value.prod_doc_cust_val_deleted_by,
				m_product_document_customer_value.prod_doc_cust_val_created_at,
				m_product_document_customer_value.prod_doc_cust_val_updated_at,
				m_product_document_customer_value.prod_doc_cust_val_deleted_at`).
		Joins("JOIN m_client_address on m_product_document_customer_value.prod_doc_cust_id=m_client_address.prod_doc_cust_id").
		Where("m_product_document_customer_value.prod_doc_cust_val_id = ?", pid).
		Find(&productdocumentcustomerimage).Error
	if err != nil {
		return &ProductDocumentCustomerValueData{}, err
	}
	return &productdocumentcustomerimage, nil
}

func (p *ProductDocumentCustomerValue) DeleteProductDocumentCustomerValue(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&ProductDocumentCustomerValue{}).Where("prod_doc_cust_val_id = ?", pid).Take(&ProductDocumentCustomerValue{}).Delete(&ProductDocumentCustomerValue{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("ProductDocumentCustomerValue not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
