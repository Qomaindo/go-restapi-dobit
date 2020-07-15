package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type ProductDocumentMitraAttributeTemp struct {
	ProductDocumentMitraAttributeTempID        uint64    `gorm:"column:prod_doc_mitra_attr_temp_id;primary_key;" json:"prod_doc_mitra_attr_temp_id"`
	ProductDocumentMitraAttributeID            uint64    `gorm:"column:prod_doc_mitra_attr_id" json:"prod_doc_mitra_attr_id"`
	ProductDocumentMitraTempID                 uint64    `gorm:"column:prod_doc_mitra_temp_id" json:"prod_doc_mitra_temp_id"`
	ProductDocumentMitraAttributeTempCode      string    `gorm:"column:prod_doc_mitra_attr_temp_code;size:15" json:"prod_doc_mitra_attr_temp_code"`
	ProductDocumentMitraAttributeTempName      string    `gorm:"column:prod_doc_mitra_attr_temp_name;size:125" json:"prod_doc_mitra_attr_temp_name"`
	ProductDocumentMitraAttributeTempDesc      string    `gorm:"column:prod_doc_mitra_attr_temp_desc" json:"prod_doc_mitra_attr_temp_desc"`
	ProductDocumentMitraAttributeTempReason    string    `gorm:"column:prod_doc_mitra_attr_temp_reason" json:"prod_doc_mitra_attr_temp_reason"`
	ProductDocumentMitraAttributeTempStatus    uint64    `gorm:"column:prod_doc_mitra_attr_temp_status;size:2" json:"prod_doc_mitra_attr_temp_status"`
	ProductDocumentMitraAttributeTempCreatedBy string    `gorm:"column:prod_doc_mitra_attr_temp_created_by;size:125" json:"prod_doc_mitra_attr_temp_created_by"`
	ProductDocumentMitraAttributeTempCreatedAt time.Time `gorm:"column:prod_doc_mitra_attr_temp_created_at;default:CURRENT_TIMESTAMP" json:"prod_doc_mitra_attr_temp_created_at"`
}

type ProductDocumentMitraAttributeTempPend struct {
	ProductDocumentMitraAttributeTempID        uint64    `gorm:"column:prod_doc_mitra_attr_temp_id;primary_key;" json:"prod_doc_mitra_attr_temp_id"`
	ProductDocumentMitraAttributeID            *uint64   `gorm:"column:prod_doc_mitra_attr_id" json:"prod_doc_mitra_attr_id"`
	ProductDocumentMitraTempID                 uint64    `gorm:"column:prod_doc_mitra_temp_id" json:"prod_doc_mitra_temp_id"`
	ProductDocumentMitraAttributeTempCode      string    `gorm:"column:prod_doc_mitra_attr_temp_code;size:15" json:"prod_doc_mitra_attr_temp_code"`
	ProductDocumentMitraAttributeTempName      string    `gorm:"column:prod_doc_mitra_attr_temp_name;size:125" json:"prod_doc_mitra_attr_temp_name"`
	ProductDocumentMitraAttributeTempDesc      string    `gorm:"column:prod_doc_mitra_attr_temp_desc" json:"prod_doc_mitra_attr_temp_desc"`
	ProductDocumentMitraAttributeTempReason    string    `gorm:"column:prod_doc_mitra_attr_temp_reason" json:"prod_doc_mitra_attr_temp_reason"`
	ProductDocumentMitraAttributeTempStatus    uint64    `gorm:"column:prod_doc_mitra_attr_temp_status;size:2" json:"prod_doc_mitra_attr_temp_status"`
	ProductDocumentMitraAttributeTempCreatedBy string    `gorm:"column:prod_doc_mitra_attr_temp_created_by;size:125" json:"prod_doc_mitra_attr_temp_created_by"`
	ProductDocumentMitraAttributeTempCreatedAt time.Time `gorm:"column:prod_doc_mitra_attr_temp_created_at;default:CURRENT_TIMESTAMP" json:"prod_doc_mitra_attr_temp_created_at"`
}

//kosong tanpa size yang untuk nampilin datanya
type ProductDocumentMitraAttributeTempData struct {
	ProductDocumentMitraAttributeTempID        uint64     `gorm:"column:prod_doc_mitra_attr_temp_id;primary_key;" json:"prod_doc_mitra_attr_temp_id"`
	ProductDocumentMitraTempID                 uint64     `gorm:"column:prod_doc_mitra_temp_id" json:"prod_doc_mitra_temp_id"`
	ProductDocumentMitraAttributeID            uint64     `gorm:"column:prod_doc_mitra_attr_id" json:"prod_doc_mitra_attr_id"`
	ProductDocumentMitraAttributeTempCode      string     `gorm:"column:prod_doc_mitra_attr_temp_code" json:"prod_doc_mitra_attr_temp_code"`
	ProductDocumentMitraAttributeTempName      string     `gorm:"column:prod_doc_mitra_attr_temp_name" json:"prod_doc_mitra_attr_temp_name"`
	ProductDocumentMitraAttributeTempDesc      string     `gorm:"column:prod_doc_mitra_attr_temp_desc" json:"prod_doc_mitra_attr_temp_desc"`
	ProductDocumentMitraAttributeTempReason    string     `gorm:"column:prod_doc_mitra_attr_temp_reason" json:"prod_doc_mitra_attr_temp_reason"`
	ProductDocumentMitraAttributeTempStatus    uint64     `gorm:"column:prod_doc_mitra_attr_temp_status" json:"prod_doc_mitra_attr_temp_status"`
	ProductDocumentMitraAttributeTempCreatedBy string     `gorm:"column:prod_doc_mitra_attr_temp_created_by" json:"prod_doc_mitra_attr_temp_created_by"`
	ProductDocumentMitraAttributeTempCreatedAt time.Time  `gorm:"column:prod_doc_mitra_attr_temp_created_at;default:CURRENT_TIMESTAMP" json:"prod_doc_mitra_attr_temp_created_at"`
	ProductDocumentMitraAttributeCode          string     `gorm:"column:prod_doc_mitra_attr_code" json:"prod_doc_mitra_attr_code"`
	ProductDocumentMitraAttributeName          string     `gorm:"column:prod_doc_mitra_attr_name" json:"prod_doc_mitra_attr_name"`
	ProductDocumentMitraAttributeDesc          string     `gorm:"column:prod_doc_mitra_attr_desc" json:"prod_doc_mitra_attr_desc"`
	ProductDocumentMitraAttributeStatus        uint64     `gorm:"column:prod_doc_mitra_attr_status" json:"prod_doc_mitra_attr_status"`
	ProductDocumentMitraAttributeCreatedBy     string     `gorm:"column:prod_doc_mitra_attr_created_by" json:"prod_doc_mitra_attr_created_by"`
	ProductDocumentMitraAttributeCreatedAt     time.Time  `gorm:"column:prod_doc_mitra_attr_created_at" json:"prod_doc_mitra_attr_created_at"`
	ProductDocumentMitraAttributeUpdatedBy     string     `gorm:"column:prod_doc_mitra_attr_updated_by" json:"prod_doc_mitra_attr_updated_by"`
	ProductDocumentMitraAttributeUpdatedAt     *time.Time `gorm:"column:prod_doc_mitra_attr_updated_at" json:"prod_doc_mitra_attr_updated_at"`
	ProductDocumentMitraAttributeDeletedBy     string     `gorm:"column:prod_doc_mitra_attr_deleted_by" json:"prod_doc_mitra_attr_deleted_by"`
	ProductDocumentMitraAttributeDeletedAt     *time.Time `gorm:"column:prod_doc_mitra_attr_deleted_at" json:"prod_doc_mitra_attr_deleted_at"`
}

type ResponseProductDocumentMitraAttributeTemps struct {
	Status  int                                      `json:"status"`
	Message string                                   `json:"message"`
	Data    *[]ProductDocumentMitraAttributeTempData `json:"data"`
}

type ResponseProductDocumentMitraAttributeTemp struct {
	Status  int                                    `json:"status"`
	Message string                                 `json:"message"`
	Data    *ProductDocumentMitraAttributeTempData `json:"data"`
}

type ResponseProductDocumentMitraAttributeTempsPend struct {
	Status  int                                      `json:"status"`
	Message string                                   `json:"message"`
	Data    *[]ProductDocumentMitraAttributeTempPend `json:"data"`
}

type ResponseProductDocumentMitraAttributeTempPend struct {
	Status  int                                    `json:"status"`
	Message string                                 `json:"message"`
	Data    *ProductDocumentMitraAttributeTempPend `json:"data"`
}

// UI diganti jadi CUD
type ResponseProductDocumentMitraAttributeTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (ProductDocumentMitraAttributeTemp) TableName() string {
	return "m_product_document_mitra_attribute_temp"
}

func (ProductDocumentMitraAttributeTempPend) TableName() string {
	return "m_product_document_mitra_attribute_temp"
}

func (ProductDocumentMitraAttributeTempData) TableName() string {
	return "m_product_document_mitra_attribute_temp"
}

func (p *ProductDocumentMitraAttributeTemp) Prepare() {
	p.ProductDocumentMitraAttributeID = p.ProductDocumentMitraAttributeID
	p.ProductDocumentMitraTempID = p.ProductDocumentMitraTempID
	p.ProductDocumentMitraAttributeTempCode = html.EscapeString(strings.TrimSpace(p.ProductDocumentMitraAttributeTempCode))
	p.ProductDocumentMitraAttributeTempName = html.EscapeString(strings.TrimSpace(p.ProductDocumentMitraAttributeTempName))
	p.ProductDocumentMitraAttributeTempDesc = html.EscapeString(strings.TrimSpace(p.ProductDocumentMitraAttributeTempDesc))
	p.ProductDocumentMitraAttributeTempReason = p.ProductDocumentMitraAttributeTempReason
	p.ProductDocumentMitraAttributeTempStatus = p.ProductDocumentMitraAttributeTempStatus
	p.ProductDocumentMitraAttributeTempCreatedBy = html.EscapeString(strings.TrimSpace(p.ProductDocumentMitraAttributeTempCreatedBy))
	p.ProductDocumentMitraAttributeTempCreatedAt = time.Now()
}

func (p *ProductDocumentMitraAttributeTempPend) Prepare() {
	p.ProductDocumentMitraAttributeID = nil
	p.ProductDocumentMitraTempID = p.ProductDocumentMitraTempID
	p.ProductDocumentMitraAttributeTempCode = p.ProductDocumentMitraAttributeTempCode
	p.ProductDocumentMitraAttributeTempName = p.ProductDocumentMitraAttributeTempName
	p.ProductDocumentMitraAttributeTempDesc = p.ProductDocumentMitraAttributeTempDesc
	p.ProductDocumentMitraAttributeTempReason = p.ProductDocumentMitraAttributeTempReason
	p.ProductDocumentMitraAttributeTempStatus = p.ProductDocumentMitraAttributeTempStatus
	p.ProductDocumentMitraAttributeTempCreatedBy = html.EscapeString(strings.TrimSpace(p.ProductDocumentMitraAttributeTempCreatedBy))
	p.ProductDocumentMitraAttributeTempCreatedAt = time.Now()
}

//Nambah validasi dengan return b.indonesia
func (p *ProductDocumentMitraAttributeTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	case "InsertUpdate":
		if p.ProductDocumentMitraAttributeTempCode == "" {
			return errors.New("Kode Kategori Dokumen Wajib diisi")
		}
		if p.ProductDocumentMitraAttributeTempName == "" {
			return errors.New("Nama Kategori Dokumen Wajib diisi")
		}
		if p.ProductDocumentMitraAttributeTempDesc == "" {
			return errors.New("Deskripsi Kategori Dokumen Wajib diisi")
		}
		return nil

	case "Reason":
		if p.ProductDocumentMitraAttributeTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		return nil

	default:
		if p.ProductDocumentMitraAttributeTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil
	}
}

func (p *ProductDocumentMitraAttributeTemp) SaveProductDocumentMitraAttributeTemp(db *gorm.DB) (*ProductDocumentMitraAttributeTemp, error) {
	var err error
	err = db.Debug().Model(&ProductDocumentMitraAttributeTemp{}).Create(&p).Error
	if err != nil {
		return &ProductDocumentMitraAttributeTemp{}, err
	}
	return p, nil
}

func (p *ProductDocumentMitraAttributeTempPend) SaveProductDocumentMitraAttributeTempPend(db *gorm.DB) (*ProductDocumentMitraAttributeTempPend, error) {
	var err error
	err = db.Debug().Model(&ProductDocumentMitraAttributeTempPend{}).Create(&p).Error
	if err != nil {
		return &ProductDocumentMitraAttributeTempPend{}, err
	}
	return p, nil
}

func (p *ProductDocumentMitraAttributeTemp) FindAllProductDocumentMitraAttributeTemp(db *gorm.DB) (*[]ProductDocumentMitraAttributeTempPend, error) {
	var err error
	productdocumentcategorytemp := []ProductDocumentMitraAttributeTempPend{}
	err = db.Debug().Model(&ProductDocumentMitraAttributeTempPend{}).Limit(100).
		Select(`m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_id,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_temp_id,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_id,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_code,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_name,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_desc,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_reason,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_status,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_created_by,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_created_at`).
		Find(&productdocumentcategorytemp).Error
	if err != nil {
		return &[]ProductDocumentMitraAttributeTempPend{}, err
	}
	return &productdocumentcategorytemp, nil
}

func (p *ProductDocumentMitraAttributeTemp) FindAllProductDocumentMitraAttributeTempStatusPendingActive(db *gorm.DB, status uint64) (*[]ProductDocumentMitraAttributeTempPend, error) {
	var err error
	productdocumentcategorytemp := []ProductDocumentMitraAttributeTempPend{}
	err = db.Debug().Model(&ProductDocumentMitraAttributeTempPend{}).Limit(100).
		Select(`m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_id,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_temp_id,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_id,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_code,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_name,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_desc,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_reason,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_status,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_created_by,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_created_at`).
		Where("prod_doc_mitra_attr_temp_status = ?", status).
		Find(&productdocumentcategorytemp).Error
	if err != nil {
		return &[]ProductDocumentMitraAttributeTempPend{}, err
	}
	return &productdocumentcategorytemp, nil
}

func (p *ProductDocumentMitraAttributeTemp) FindAllProductDocumentMitraAttributeTempStatus(db *gorm.DB, status uint64) (*[]ProductDocumentMitraAttributeTempData, error) {
	var err error
	productdocumentcategorytemp := []ProductDocumentMitraAttributeTempData{}
	err = db.Debug().Model(&ProductDocumentMitraAttributeTempData{}).Limit(100).
		Select(`m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_id,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_temp_id,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_id,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_code,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_name,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_desc,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_reason,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_status,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_created_by,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_created_at,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_code,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_name,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_desc,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_status,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_created_by,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_created_at,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_updated_by,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_updated_at,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_deleted_by,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_deleted_at`).
		Joins("JOIN m_product_document_mitra_attribute on m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_id=m_product_document_mitra_attribute.prod_doc_mitra_attr_id").
		Where("prod_doc_mitra_attr_temp_status = ?", status).
		Find(&productdocumentcategorytemp).Error
	if err != nil {
		return &[]ProductDocumentMitraAttributeTempData{}, err
	}
	return &productdocumentcategorytemp, nil
}

func (p *ProductDocumentMitraAttributeTemp) FindProductDocumentMitraAttributeTempDataByID(db *gorm.DB, pid uint64) (*ProductDocumentMitraAttributeTemp, error) {
	var err error
	err = db.Debug().Model(&ProductDocumentMitraAttributeTemp{}).Where("prod_doc_mitra_attr_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &ProductDocumentMitraAttributeTemp{}, err
	}
	return p, nil
}

func (p *ProductDocumentMitraAttributeTemp) FindProductDocumentMitraAttributeTempByIDPendingActive(db *gorm.DB, pid uint64) (*ProductDocumentMitraAttributeTempPend, error) {
	var err error
	productdocumentcategorytemp := ProductDocumentMitraAttributeTempPend{}
	err = db.Debug().Model(&ProductDocumentMitraAttributeTempPend{}).Limit(100).
		Select(`m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_id,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_temp_id,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_id,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_code,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_name,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_desc,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_reason,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_status,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_created_by,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_created_at`).
		Where("prod_doc_mitra_attr_temp_id = ?", pid).
		Find(&productdocumentcategorytemp).Error
	if err != nil {
		return &ProductDocumentMitraAttributeTempPend{}, err
	}
	return &productdocumentcategorytemp, nil
}

func (p *ProductDocumentMitraAttributeTemp) FindProductDocumentMitraAttributeTempStatusByIDPendingActive(db *gorm.DB, pid uint64, status uint64) (*ProductDocumentMitraAttributeTempPend, error) {
	var err error
	productdocumentcategorytemp := ProductDocumentMitraAttributeTempPend{}
	err = db.Debug().Model(&ProductDocumentMitraAttributeTempPend{}).Limit(100).
		Select(`m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_id,
				m_product_document_mitra_attribute_temp.address_id,
				m_product_document_mitra_attribute_temp.country_temp_id,
				m_product_document_mitra_attribute_temp.province_temp_id,
				m_product_document_mitra_attribute_temp.regency_temp_id,
				m_product_document_mitra_attribute_temp.district_temp_id,
				m_product_document_mitra_attribute_temp.village_temp_id,
				m_product_document_mitra_attribute_temp.address_temp_reason,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_status,
				m_product_document_mitra_attribute_temp.address_temp_created_by,
				m_product_document_mitra_attribute_temp.address_temp_created_at`).
		Where("prod_doc_mitra_attr_temp_id = ? AND prod_doc_mitra_attr_temp_status = ?", pid, status).
		Find(&productdocumentcategorytemp).Error
	if err != nil {
		return &ProductDocumentMitraAttributeTempPend{}, err
	}
	return &productdocumentcategorytemp, nil
}

func (p *ProductDocumentMitraAttributeTemp) FindProductDocumentMitraAttributeTempByID(db *gorm.DB, pid uint64) (*ProductDocumentMitraAttributeTempData, error) {
	var err error
	productdocumentcategorytemp := ProductDocumentMitraAttributeTempData{}
	err = db.Debug().Model(&ProductDocumentMitraAttributeTempData{}).Limit(100).
		Select(`m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_id,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_temp_id,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_code,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_name,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_desc,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_reason,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_status,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_created_by,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_created_at,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_id,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_code,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_name,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_desc,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_status,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_created_by,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_created_at,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_updated_by,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_updated_at,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_deleted_by,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_deleted_at`).
		Joins("JOIN m_product_document_mitra_attribute on m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_id=m_product_document_mitra_attribute.prod_doc_mitra_attr_id").
		Where("prod_doc_mitra_attr_temp_id = ?", pid).
		Find(&productdocumentcategorytemp).Error
	if err != nil {
		return &ProductDocumentMitraAttributeTempData{}, err
	}
	return &productdocumentcategorytemp, nil
}

func (p *ProductDocumentMitraAttributeTemp) FindProductDocumentMitraAttributeTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*ProductDocumentMitraAttributeTempData, error) {
	var err error
	productdocumentcategorytemp := ProductDocumentMitraAttributeTempData{}
	err = db.Debug().Model(&ProductDocumentMitraAttributeTempData{}).Limit(100).
		Select(`m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_id,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_temp_id,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_id,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_code,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_name,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_desc,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_reason,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_status,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_created_by,
				m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_temp_created_at,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_code,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_name,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_desc,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_status,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_created_by,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_created_at,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_updated_by,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_updated_at,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_deleted_by,
				m_product_document_mitra_attribute.prod_doc_mitra_attr_deleted_at`).
		Joins("JOIN m_product_document_mitra_attribute on m_product_document_mitra_attribute_temp.prod_doc_mitra_attr_id=m_product_document_mitra_attribute.prod_doc_mitra_attr_id").
		Where("prod_doc_mitra_attr_temp_id = ?", pid).Where("prod_doc_mitra_attr_temp_id = ? AND prod_doc_mitra_attr_temp_status = ?", pid, status).
		Find(&productdocumentcategorytemp).Error
	if err != nil {
		return &ProductDocumentMitraAttributeTempData{}, err
	}
	return &productdocumentcategorytemp, nil
}

func (p *ProductDocumentMitraAttributeTemp) UpdateProductDocumentMitraAttributeTemp(db *gorm.DB, pid uint64) (*ProductDocumentMitraAttributeTemp, error) {
	var err error
	db = db.Debug().Model(&ProductDocumentMitraAttributeTemp{}).Where("prod_doc_mitra_attr_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_mitra_temp_id":          p.ProductDocumentMitraTempID,
			"prod_doc_mitra_attr_temp_code":   p.ProductDocumentMitraAttributeTempCode,
			"prod_doc_mitra_attr_temp_name":   p.ProductDocumentMitraAttributeTempName,
			"prod_doc_mitra_attr_temp_desc":   p.ProductDocumentMitraAttributeTempDesc,
			"prod_doc_mitra_attr_temp_reason": p.ProductDocumentMitraAttributeTempReason,
		},
	)
	err = db.Debug().Model(&ProductDocumentMitraAttributeTemp{}).Where("prod_doc_mitra_attr_temp_id = ?", pid).Error
	if err != nil {
		return &ProductDocumentMitraAttributeTemp{}, err
	}
	return p, nil
}

func (p *ProductDocumentMitraAttributeTemp) UpdateProductDocumentMitraAttributeTempStatus(db *gorm.DB, pid uint64) (*ProductDocumentMitraAttributeTemp, error) {
	var err error
	db = db.Debug().Model(&ProductDocumentMitraAttributeTemp{}).Where("prod_doc_mitra_attr_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_mitra_attr_temp_reason": p.ProductDocumentMitraAttributeTempReason,
			"prod_doc_mitra_attr_temp_status": p.ProductDocumentMitraAttributeTempStatus,
		},
	)
	err = db.Debug().Model(&ProductDocumentMitraAttributeTemp{}).Where("prod_doc_mitra_attr_temp_id = ?", pid).Error
	if err != nil {
		return &ProductDocumentMitraAttributeTemp{}, err
	}
	return p, nil
}

func (p *ProductDocumentMitraAttributeTemp) DeleteProductDocumentMitraAttributeTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&ProductDocumentMitraAttributeTemp{}).Where("prod_doc_mitra_attr_temp_id = ?", pid).Take(&ProductDocumentMitraAttributeTemp{}).Delete(&ProductDocumentMitraAttributeTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Product Document Category Temp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
