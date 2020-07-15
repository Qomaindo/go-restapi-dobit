package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type ProductDocumentCustomerImage struct {
	ProductDocumentCustomerImageID        uint64     `gorm:"column:prod_doc_cust_img_id;primary_key;" json:"prod_doc_cust_img_id"`
	ProductDocumentCustomerID             uint64     `gorm:"column:prod_doc_cust_id" json:"prod_doc_cust_id"`
	ProductDocumentCustomerImageValue     string     `gorm:"column:prod_doc_cust_img_value;size:255" json:"prod_doc_cust_img_value"`
	ProductDocumentCustomerImageCreatedBy string     `gorm:"column:prod_doc_cust_img_created_by;size:125" json:"prod_doc_cust_img_created_by"`
	ProductDocumentCustomerImageUpdatedBy string     `gorm:"column:prod_doc_cust_img_updated_by;size:125" json:"prod_doc_cust_img_updated_by"`
	ProductDocumentCustomerImageDeletedBy string     `gorm:"column:prod_doc_cust_img_deleted_by;size:125" json:"prod_doc_cust_img_deleted_by"`
	ProductDocumentCustomerImageCreatedAt time.Time  `gorm:"column:prod_doc_cust_id;default:CURRENT_TIMESTAMP" json:"prod_doc_cust_id"`
	ProductDocumentCustomerImageUpdatedAt *time.Time `gorm:"column:prod_doc_cust_id;default:CURRENT_TIMESTAMP" json:"prod_doc_cust_id"`
	ProductDocumentCustomerImageDeletedAt *time.Time `gorm:"column:prod_doc_cust_id;default:CURRENT_TIMESTAMP" json:"prod_doc_cust_id"`
}

type ProductDocumentCustomerImageData struct {
	ProductDocumentCustomerImageID        uint64     `gorm:"column:prod_doc_cust_img_id;primary_key;" json:"prod_doc_cust_img_id"`
	ProductDocumentCustomerID             uint64     `gorm:"column:prod_doc_cust_id;" json:"prod_doc_cust_id"`
	ProductDocumentCustomerImageValue     string     `gorm:"column:prod_doc_cust_id;size:255" json:"prod_doc_cust_id"`
	ProductDocumentCustomerImageCreatedBy string     `gorm:"column:prod_doc_cust_img_created_by;size:125" json:"prod_doc_cust_img_created_by"`
	ProductDocumentCustomerImageUpdatedBy string     `gorm:"column:prod_doc_cust_img_updated_by;size:125" json:"prod_doc_cust_img_updated_by"`
	ProductDocumentCustomerImageDeletedBy string     `gorm:"column:prod_doc_cust_img_deleted_by;size:125" json:"prod_doc_cust_img_deleted_by"`
	ProductDocumentCustomerImageCreatedAt time.Time  `gorm:"column:prod_doc_cust_id;default:CURRENT_TIMESTAMP" json:"prod_doc_cust_id"`
	ProductDocumentCustomerImageUpdatedAt *time.Time `gorm:"column:prod_doc_cust_id;default:CURRENT_TIMESTAMP" json:"prod_doc_cust_id"`
	ProductDocumentCustomerImageDeletedAt *time.Time `gorm:"column:prod_doc_cust_id;default:CURRENT_TIMESTAMP" json:"prod_doc_cust_id"`
}

type ResponseProductDocumentCustomerImages struct {
	Status  uint64                              `json:"status"`
	Message string                              `json:"message"`
	Data    *[]ProductDocumentCustomerImageData `json:"data"`
}

type ResponseProductDocumentCustomerImage struct {
	Status  uint64                        `json:"status"`
	Message string                        `json:"message"`
	Data    *ProductDocumentCustomerImage `json:"data"`
}
type ResponseProductDocumentCustomerImageData struct {
	Status  uint64                            `json:"status"`
	Message string                            `json:"message"`
	Data    *ProductDocumentCustomerImageData `json:"data"`
}
type ResponseProductDocumentCustomerImageIU struct {
	Status  uint64                        `json:"status"`
	Message string                        `json:"message"`
	Data    *ProductDocumentCustomerImage `json:"data"`
}

type ResponseProductDocumentCustomerImageDel struct {
	Status  uint64 `json:"status"`
	Message string `json:"message"`
}

func (ProductDocumentCustomerImage) TableName() string {
	return "m_product_document_customer_image"
}

func (ProductDocumentCustomerImageData) TableName() string {
	return "m_product_document_customer_image"
}

func (u *ProductDocumentCustomerImage) Prepare() {
	u.ProductDocumentCustomerImageID = u.ProductDocumentCustomerImageID
	u.ProductDocumentCustomerImageValue = html.EscapeString(strings.TrimSpace(u.ProductDocumentCustomerImageValue))
	u.ProductDocumentCustomerImageCreatedBy = u.ProductDocumentCustomerImageCreatedBy
	u.ProductDocumentCustomerImageUpdatedBy = u.ProductDocumentCustomerImageUpdatedBy
	u.ProductDocumentCustomerImageDeletedBy = u.ProductDocumentCustomerImageDeletedBy
	u.ProductDocumentCustomerImageCreatedAt = time.Now()
	u.ProductDocumentCustomerImageUpdatedAt = u.ProductDocumentCustomerImageUpdatedAt
	u.ProductDocumentCustomerImageDeletedAt = u.ProductDocumentCustomerImageDeletedAt
}

func (u *ProductDocumentCustomerImage) Validate(action string) error {
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

func (u *ProductDocumentCustomerImage) SaveProductDocumentCustomerImage(db *gorm.DB) (*ProductDocumentCustomerImage, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &ProductDocumentCustomerImage{}, err
	}
	return u, nil
}

func (p *ProductDocumentCustomerImage) FindAllProductDocumentCustomerImage(db *gorm.DB) (*[]ProductDocumentCustomerImageData, error) {
	var err error
	productdocumentcustomerimage := []ProductDocumentCustomerImageData{}
	err = db.Debug().Model(&ProductDocumentCustomerImageData{}).Limit(100).
		Select(`m_product_document_customer_image.prod_doc_cust_img_id,
				m_product_document_customer_image.prod_doc_cust_id,
				m_product_document_customer_image.prod_doc_cust_img_value,
				m_product_document_customer_image.prod_doc_cust_img_created_by,
				m_product_document_customer_image.prod_doc_cust_img_updated_by,
				m_product_document_customer_image.prod_doc_cust_img_deleted_by,
				m_product_document_customer_image.prod_doc_cust_img_created_at,
				m_product_document_customer_image.prod_doc_cust_img_updated_at,
				m_product_document_customer_image.prod_doc_cust_img_deleted_at`).
		Joins("JOIN m_client_address on m_product_document_customer_image.prod_doc_cust_id=m_client_address.prod_doc_cust_id").
		Find(&productdocumentcustomerimage).Error
	if err != nil {
		return &[]ProductDocumentCustomerImageData{}, err
	}
	return &productdocumentcustomerimage, nil
}

func (p *ProductDocumentCustomerImage) FindAllProductDocumentCustomerImageStatus(db *gorm.DB, status uint64) (*[]ProductDocumentCustomerImageData, error) {
	var err error
	productdocumentcustomerimage := []ProductDocumentCustomerImageData{}
	err = db.Debug().Model(&ProductDocumentCustomerImageData{}).Limit(100).
		Select(`m_product_document_customer_image.prod_doc_cust_img_id,
				m_product_document_customer_image.prod_doc_cust_id,
				m_product_document_customer_image.prod_doc_cust_img_value,
				m_product_document_customer_image.prod_doc_cust_img_created_by,
				m_product_document_customer_image.prod_doc_cust_img_updated_by,
				m_product_document_customer_image.prod_doc_cust_img_deleted_by,
				m_product_document_customer_image.prod_doc_cust_img_created_at,
				m_product_document_customer_image.prod_doc_cust_img_updated_at,
				m_product_document_customer_image.prod_doc_cust_img_deleted_at`).
		Joins("JOIN m_client_address on m_product_document_customer_image.prod_doc_cust_id=m_client_address.prod_doc_cust_id").
		Where("client_adrs_img_status = ?", status).
		Find(&productdocumentcustomerimage).Error
	if err != nil {
		return &[]ProductDocumentCustomerImageData{}, err
	}
	return &productdocumentcustomerimage, nil
}

func (p *ProductDocumentCustomerImage) FindProductDocumentCustomerImageByIDAddress(db *gorm.DB, pid uint64) (*ProductDocumentCustomerImageData, error) {
	var err error
	productdocumentcustomerimage := ProductDocumentCustomerImageData{}
	err = db.Debug().Model(&ProductDocumentCustomerImageData{}).Limit(100).
		Select(`m_product_document_customer_image.prod_doc_cust_img_id,
				m_product_document_customer_image.prod_doc_cust_id,
				m_product_document_customer_image.prod_doc_cust_img_value,
				m_product_document_customer_image.prod_doc_cust_img_created_by,
				m_product_document_customer_image.prod_doc_cust_img_updated_by,
				m_product_document_customer_image.prod_doc_cust_img_deleted_by,
				m_product_document_customer_image.prod_doc_cust_img_created_at,
				m_product_document_customer_image.prod_doc_cust_img_updated_at,
				m_product_document_customer_image.prod_doc_cust_img_deleted_at`).
		// Joins("JOIN m_clientomer_address on m_clientomer_address_image.prod_doc_cust_id=m_clientomer_address.prod_doc_cust_id").
		Where("prod_doc_cust_img_id = ?", pid).
		Find(&productdocumentcustomerimage).Error
	if err != nil {
		return &ProductDocumentCustomerImageData{}, err
	}
	return &productdocumentcustomerimage, nil
}

// func (p *ProductDocumentCustomerImage) FindProductDocumentCustomerImageStatusByID(db *gorm.DB, pid uint64, status uint64) (*ProductDocumentCustomerImageData, error) {
// 	var err error
// 	productdocumentcustomerimage := ProductDocumentCustomerImageData{}
// 	err = db.Debug().Model(&ProductDocumentCustomerImageData{}).Limit(100).
// 		Select(`m_product_document_customer_image.prod_doc_cust_img_id,
// 				m_client_address.prod_doc_cust_id,
// 				m_product_document_customer_image.client_adrs_img_name,
// 				m_product_document_customer_image.client_adrs_img_desc,
// 				m_product_document_customer_image.prod_doc_cust_id,
// 				m_product_document_customer_image.prod_doc_cust_id,
// 				m_product_document_customer_image.prod_doc_cust_id`).
// 		Joins("JOIN m_client_address on m_product_document_customer_image.prod_doc_cust_id=m_client_address.prod_doc_cust_id").
// 		Where("prod_doc_cust_img_id = ? AND client_adrs_img_status = ?", pid, status).
// 		Find(&productdocumentcustomerimage).Error
// 	if err != nil {
// 		return &ProductDocumentCustomerImageData{}, err
// 	}
// 	return &productdocumentcustomerimage, nil
// }

func (p *ProductDocumentCustomerImage) FindProductDocumentCustomerImageByProductDocumentCustomerImageID(db *gorm.DB, pid uint64) (*ProductDocumentCustomerImageData, error) {
	var err error
	productdocumentcustomerimage := ProductDocumentCustomerImageData{}
	err = db.Debug().Model(&ProductDocumentCustomerImageData{}).Limit(100).
		Select(`m_product_document_customer_image.prod_doc_cust_img_id,
		m_product_document_customer_image.prod_doc_cust_id,
				m_product_document_customer_image.prod_doc_cust_img_value,
				m_product_document_customer_image.prod_doc_cust_img_created_by,
				m_product_document_customer_image.prod_doc_cust_img_updated_by,
				m_product_document_customer_image.prod_doc_cust_img_deleted_by,
				m_product_document_customer_image.prod_doc_cust_img_created_at,
				m_product_document_customer_image.prod_doc_cust_img_updated_at,
				m_product_document_customer_image.prod_doc_cust_img_deleted_at`).
		Joins("JOIN m_client_address on m_product_document_customer_image.prod_doc_cust_id=m_client_address.prod_doc_cust_id").
		Where("m_product_document_customer_image.prod_doc_cust_img_id = ?", pid).
		Find(&productdocumentcustomerimage).Error
	if err != nil {
		return &ProductDocumentCustomerImageData{}, err
	}
	return &productdocumentcustomerimage, nil
}

func (p *ProductDocumentCustomerImage) DeleteProductDocumentCustomerImage(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&ProductDocumentCustomerImage{}).Where("prod_doc_cust_img_id = ?", pid).Take(&ProductDocumentCustomerImage{}).Delete(&ProductDocumentCustomerImage{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("ProductDocumentCustomerImage not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
