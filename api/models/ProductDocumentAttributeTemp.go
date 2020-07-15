package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type ProductDocumentAttributeTemp struct {
	ProductDocumentAttributeTempID        uint64    `gorm:"column:prod_doc_attr_temp_id;primary_key;" json:"prod_doc_attr_temp_id"`
	ProductDocumentAttributeID            uint64    `gorm:"column:prod_doc_attr_id" json:"prod_doc_attr_id"`
	ProductDocumentAttributeTempCode      string    `gorm:"column:prod_doc_attr_temp_code;size:15" json:"prod_doc_attr_temp_code"`
	ProductDocumentAttributeTempName      string    `gorm:"column:prod_doc_attr_temp_name;size:255" json:"prod_doc_attr_temp_name"`
	ProductDocumentAttributeTempDesc      string    `gorm:"column:prod_doc_attr_temp_desc" json:"prod_doc_attr_temp_desc"`
	ProductDocumentAttributeTempType      string    `gorm:"column:prod_doc_attr_temp_type;size:255" json:"prod_doc_attr_temp_type"`
	ProductDocumentAttributeTempReason    string    `gorm:"column:prod_doc_attr_temp_reason" json:"prod_doc_attr_temp_reason"`
	ProductDocumentAttributeTempStatus    uint64    `gorm:"column:prod_doc_attr_temp_status;size:2" json:"prod_doc_attr_temp_status"`
	ProductDocumentAttributeTempCreatedBy string    `gorm:"column:prod_doc_attr_temp_created_by;size:125" json:"prod_doc_attr_temp_created_by"`
	ProductDocumentAttributeTempCreatedAt time.Time `gorm:"column:prod_doc_attr_temp_created_at;default:CURRENT_TIMESTAMP" json:"prod_doc_attr_temp_created_at"`
}

type ProductDocumentAttributeTempPend struct {
	ProductDocumentAttributeTempID        uint64    `gorm:"column:prod_doc_attr_temp_id;primary_key;" json:"prod_doc_attr_temp_id"`
	ProductDocumentAttributeID            *uint64   `gorm:"column:prod_doc_attr_id" json:"prod_doc_attr_id"`
	ProductDocumentAttributeTempCode      string    `gorm:"column:prod_doc_attr_temp_code;size:15" json:"prod_doc_attr_temp_code"`
	ProductDocumentAttributeTempName      string    `gorm:"column:prod_doc_attr_temp_name;size:255" json:"prod_doc_attr_temp_name"`
	ProductDocumentAttributeTempDesc      string    `gorm:"column:prod_doc_attr_temp_desc" json:"prod_doc_attr_temp_desc"`
	ProductDocumentAttributeTempType      string    `gorm:"column:prod_doc_attr_temp_type;size:255" json:"prod_doc_attr_temp_type"`
	ProductDocumentAttributeTempReason    string    `gorm:"column:prod_doc_attr_temp_reason" json:"prod_doc_attr_temp_reason"`
	ProductDocumentAttributeTempStatus    uint64    `gorm:"column:prod_doc_attr_temp_status;size:2" json:"prod_doc_attr_temp_status"`
	ProductDocumentAttributeTempCreatedBy string    `gorm:"column:prod_doc_attr_temp_created_by;size:125" json:"prod_doc_attr_temp_created_by"`
	ProductDocumentAttributeTempCreatedAt time.Time `gorm:"column:prod_doc_attr_temp_created_at;default:CURRENT_TIMESTAMP" json:"prod_doc_attr_temp_created_at"`
}

//kosong tanpa size yang untuk nampilin datanya
type ProductDocumentAttributeTempData struct {
	ProductDocumentAttributeTempID        uint64     `gorm:"column:prod_doc_attr_temp_id;primary_key;" json:"prod_doc_attr_temp_id"`
	ProductDocumentAttributeID            uint64     `gorm:"column:prod_doc_attr_id" json:"prod_doc_attr_id"`
	ProductDocumentAttributeTempCode      string     `gorm:"column:prod_doc_attr_temp_code" json:"prod_doc_attr_temp_code"`
	ProductDocumentAttributeTempName      string     `gorm:"column:prod_doc_attr_temp_name" json:"prod_doc_attr_temp_name"`
	ProductDocumentAttributeTempDesc      string     `gorm:"column:prod_doc_attr_temp_desc" json:"prod_doc_attr_temp_desc"`
	ProductDocumentAttributeTempType      string     `gorm:"column:prod_doc_attr_temp_type" json:"prod_doc_attr_temp_type"`
	ProductDocumentAttributeTempReason    string     `gorm:"column:prod_doc_attr_temp_reason" json:"prod_doc_attr_temp_reason"`
	ProductDocumentAttributeTempStatus    uint64     `gorm:"column:prod_doc_attr_temp_status" json:"prod_doc_attr_temp_status"`
	ProductDocumentAttributeTempCreatedBy string     `gorm:"column:prod_doc_attr_temp_created_by" json:"prod_doc_attr_temp_created_by"`
	ProductDocumentAttributeTempCreatedAt time.Time  `gorm:"column:prod_doc_attr_temp_created_at;default:CURRENT_TIMESTAMP" json:"prod_doc_attr_temp_created_at"`
	ProductDocumentAttributeCode          string     `gorm:"column:prod_doc_attr_code" json:"prod_doc_attr_code"`
	ProductDocumentAttributeName          string     `gorm:"column:prod_doc_attr_name" json:"prod_doc_attr_name"`
	ProductDocumentAttributeDesc          string     `gorm:"column:prod_doc_attr_desc" json:"prod_doc_attr_desc"`
	ProductDocumentAttributeType          string     `gorm:"column:prod_doc_attr_type" json:"prod_doc_attr_type"`
	ProductDocumentAttributeStatus        uint64     `gorm:"column:prod_doc_attr_status" json:"prod_doc_attr_status"`
	ProductDocumentAttributeCreatedBy     string     `gorm:"column:prod_doc_attr_created_by" json:"prod_doc_attr_created_by"`
	ProductDocumentAttributeCreatedAt     time.Time  `gorm:"column:prod_doc_attr_created_at" json:"prod_doc_attr_created_at"`
	ProductDocumentAttributeUpdatedBy     string     `gorm:"column:prod_doc_attr_updated_by" json:"prod_doc_attr_updated_by"`
	ProductDocumentAttributeUpdatedAt     *time.Time `gorm:"column:prod_doc_attr_updated_at" json:"prod_doc_attr_updated_at"`
	ProductDocumentAttributeDeletedBy     string     `gorm:"column:prod_doc_attr_deleted_by" json:"prod_doc_attr_deleted_by"`
	ProductDocumentAttributeDeletedAt     *time.Time `gorm:"column:prod_doc_attr_deleted_at" json:"prod_doc_attr_deleted_at"`
}

type ResponseProductDocumentAttributeTemps struct {
	Status  int                                 `json:"status"`
	Message string                              `json:"message"`
	Data    *[]ProductDocumentAttributeTempData `json:"data"`
}

type ResponseProductDocumentAttributeTemp struct {
	Status  int                               `json:"status"`
	Message string                            `json:"message"`
	Data    *ProductDocumentAttributeTempData `json:"data"`
}

type ResponseProductDocumentAttributeTempsPend struct {
	Status  int                                 `json:"status"`
	Message string                              `json:"message"`
	Data    *[]ProductDocumentAttributeTempPend `json:"data"`
}

type ResponseProductDocumentAttributeTempPend struct {
	Status  int                               `json:"status"`
	Message string                            `json:"message"`
	Data    *ProductDocumentAttributeTempPend `json:"data"`
}

// UI diganti jadi CUD
type ResponseProductDocumentAttributeTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (ProductDocumentAttributeTemp) TableName() string {
	return "m_product_document_attribute_temp"
}

func (ProductDocumentAttributeTempPend) TableName() string {
	return "m_product_document_attribute_temp"
}

func (ProductDocumentAttributeTempData) TableName() string {
	return "m_product_document_attribute_temp"
}

func (p *ProductDocumentAttributeTemp) Prepare() {
	p.ProductDocumentAttributeID = p.ProductDocumentAttributeID
	p.ProductDocumentAttributeTempCode = html.EscapeString(strings.TrimSpace(p.ProductDocumentAttributeTempCode))
	p.ProductDocumentAttributeTempName = html.EscapeString(strings.TrimSpace(p.ProductDocumentAttributeTempName))
	p.ProductDocumentAttributeTempDesc = html.EscapeString(strings.TrimSpace(p.ProductDocumentAttributeTempDesc))
	p.ProductDocumentAttributeTempType = html.EscapeString(strings.TrimSpace(p.ProductDocumentAttributeTempType))
	p.ProductDocumentAttributeTempReason = p.ProductDocumentAttributeTempReason
	p.ProductDocumentAttributeTempStatus = p.ProductDocumentAttributeTempStatus
	p.ProductDocumentAttributeTempCreatedBy = html.EscapeString(strings.TrimSpace(p.ProductDocumentAttributeTempCreatedBy))
	p.ProductDocumentAttributeTempCreatedAt = time.Now()
}

func (p *ProductDocumentAttributeTempPend) Prepare() {
	p.ProductDocumentAttributeID = nil
	p.ProductDocumentAttributeTempCode = p.ProductDocumentAttributeTempCode
	p.ProductDocumentAttributeTempName = p.ProductDocumentAttributeTempName
	p.ProductDocumentAttributeTempDesc = p.ProductDocumentAttributeTempDesc
	p.ProductDocumentAttributeTempType = p.ProductDocumentAttributeTempType
	p.ProductDocumentAttributeTempReason = p.ProductDocumentAttributeTempReason
	p.ProductDocumentAttributeTempStatus = p.ProductDocumentAttributeTempStatus
	p.ProductDocumentAttributeTempCreatedBy = html.EscapeString(strings.TrimSpace(p.ProductDocumentAttributeTempCreatedBy))
	p.ProductDocumentAttributeTempCreatedAt = time.Now()
}

//Nambah validasi dengan return b.indonesia
func (p *ProductDocumentAttributeTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	case "InsertUpdate":
		if p.ProductDocumentAttributeTempCode == "" {
			return errors.New("Kode Kategori Dokumen Wajib diisi")
		}
		if p.ProductDocumentAttributeTempName == "" {
			return errors.New("Nama Kategori Dokumen Wajib diisi")
		}
		if p.ProductDocumentAttributeTempDesc == "" {
			return errors.New("Deskripsi Kategori Dokumen Wajib diisi")
		}
		if p.ProductDocumentAttributeTempType == "" {
			return errors.New("Deskripsi Kategori Dokumen Wajib diisi")
		}
		return nil

	case "Reason":
		if p.ProductDocumentAttributeTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		return nil

	default:
		if p.ProductDocumentAttributeTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil
	}
}

func (p *ProductDocumentAttributeTemp) SaveProductDocumentAttributeTemp(db *gorm.DB) (*ProductDocumentAttributeTemp, error) {
	var err error
	err = db.Debug().Model(&ProductDocumentAttributeTemp{}).Create(&p).Error
	if err != nil {
		return &ProductDocumentAttributeTemp{}, err
	}
	return p, nil
}

func (p *ProductDocumentAttributeTempPend) SaveProductDocumentAttributeTempPend(db *gorm.DB) (*ProductDocumentAttributeTempPend, error) {
	var err error
	err = db.Debug().Model(&ProductDocumentAttributeTempPend{}).Create(&p).Error
	if err != nil {
		return &ProductDocumentAttributeTempPend{}, err
	}
	return p, nil
}

func (p *ProductDocumentAttributeTemp) FindAllProductDocumentAttributeTemp(db *gorm.DB) (*[]ProductDocumentAttributeTempPend, error) {
	var err error
	productdocumentcategorytemp := []ProductDocumentAttributeTempPend{}
	err = db.Debug().Model(&ProductDocumentAttributeTempPend{}).Limit(100).
		Select(`m_product_document_attribute_temp.prod_doc_attr_temp_id,
				m_product_document_attribute_temp.prod_doc_attr_id,
				m_product_document_attribute_temp.prod_doc_attr_temp_code,
				m_product_document_attribute_temp.prod_doc_attr_temp_name,
				m_product_document_attribute_temp.prod_doc_attr_temp_desc,
				m_product_document_attribute_temp.prod_doc_attr_temp_type,
				m_product_document_attribute_temp.prod_doc_attr_temp_reason,
				m_product_document_attribute_temp.prod_doc_attr_temp_status,
				m_product_document_attribute_temp.prod_doc_attr_temp_created_by,
				m_product_document_attribute_temp.prod_doc_attr_temp_created_at`).
		Find(&productdocumentcategorytemp).Error
	if err != nil {
		return &[]ProductDocumentAttributeTempPend{}, err
	}
	return &productdocumentcategorytemp, nil
}

func (p *ProductDocumentAttributeTemp) FindAllProductDocumentAttributeTempStatusPendingActive(db *gorm.DB, status uint64) (*[]ProductDocumentAttributeTempPend, error) {
	var err error
	productdocumentcategorytemp := []ProductDocumentAttributeTempPend{}
	err = db.Debug().Model(&ProductDocumentAttributeTempPend{}).Limit(100).
		Select(`m_product_document_attribute_temp.prod_doc_attr_temp_id,
				m_product_document_attribute_temp.prod_doc_attr_id,
				m_product_document_attribute_temp.prod_doc_attr_temp_code,
				m_product_document_attribute_temp.prod_doc_attr_temp_name,
				m_product_document_attribute_temp.prod_doc_attr_temp_desc,
				m_product_document_attribute_temp.prod_doc_attr_temp_type,
				m_product_document_attribute_temp.prod_doc_attr_temp_reason,
				m_product_document_attribute_temp.prod_doc_attr_temp_status,
				m_product_document_attribute_temp.prod_doc_attr_temp_created_by,
				m_product_document_attribute_temp.prod_doc_attr_temp_created_at`).
		Where("prod_doc_attr_temp_status = ?", status).
		Find(&productdocumentcategorytemp).Error
	if err != nil {
		return &[]ProductDocumentAttributeTempPend{}, err
	}
	return &productdocumentcategorytemp, nil
}

func (p *ProductDocumentAttributeTemp) FindAllProductDocumentAttributeTempStatus(db *gorm.DB, status uint64) (*[]ProductDocumentAttributeTempData, error) {
	var err error
	productdocumentcategorytemp := []ProductDocumentAttributeTempData{}
	err = db.Debug().Model(&ProductDocumentAttributeTempData{}).Limit(100).
		Select(`m_product_document_attribute_temp.prod_doc_attr_temp_id,
				m_product_document_attribute_temp.prod_doc_attr_id,
				m_product_document_attribute_temp.prod_doc_attr_temp_code,
				m_product_document_attribute_temp.prod_doc_attr_temp_name,
				m_product_document_attribute_temp.prod_doc_attr_temp_desc,
				m_product_document_attribute_temp.prod_doc_attr_temp_type,
				m_product_document_attribute_temp.prod_doc_attr_temp_reason,
				m_product_document_attribute_temp.prod_doc_attr_temp_status,
				m_product_document_attribute_temp.prod_doc_attr_temp_created_by,
				m_product_document_attribute_temp.prod_doc_attr_temp_created_at,
				m_product_document_attribute.prod_doc_attr_code,
				m_product_document_attribute.prod_doc_attr_name,
				m_product_document_attribute.prod_doc_attr_desc,
				m_product_document_attribute.prod_doc_attr_type,
				m_product_document_attribute.prod_doc_attr_status,
				m_product_document_attribute.prod_doc_attr_created_by,
				m_product_document_attribute.prod_doc_attr_created_at,
				m_product_document_attribute.prod_doc_attr_updated_by,
				m_product_document_attribute.prod_doc_attr_updated_at,
				m_product_document_attribute.prod_doc_attr_deleted_by,
				m_product_document_attribute.prod_doc_attr_deleted_at`).
		Joins("JOIN m_product_document_attribute on m_product_document_attribute_temp.prod_doc_attr_id=m_product_document_attribute.prod_doc_attr_id").
		Where("prod_doc_attr_temp_status = ?", status).
		Find(&productdocumentcategorytemp).Error
	if err != nil {
		return &[]ProductDocumentAttributeTempData{}, err
	}
	return &productdocumentcategorytemp, nil
}

func (p *ProductDocumentAttributeTemp) FindProductDocumentAttributeTempDataByID(db *gorm.DB, pid uint64) (*ProductDocumentAttributeTemp, error) {
	var err error
	err = db.Debug().Model(&ProductDocumentAttributeTemp{}).Where("prod_doc_attr_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &ProductDocumentAttributeTemp{}, err
	}
	return p, nil
}

func (p *ProductDocumentAttributeTemp) FindProductDocumentAttributeTempByIDPendingActive(db *gorm.DB, pid uint64) (*ProductDocumentAttributeTempPend, error) {
	var err error
	productdocumentcategorytemp := ProductDocumentAttributeTempPend{}
	err = db.Debug().Model(&ProductDocumentAttributeTempPend{}).Limit(100).
		Select(`m_product_document_attribute_temp.prod_doc_attr_temp_id,
				m_product_document_attribute_temp.prod_doc_attr_id,
				m_product_document_attribute_temp.prod_doc_attr_temp_code,
				m_product_document_attribute_temp.prod_doc_attr_temp_name,
				m_product_document_attribute_temp.prod_doc_attr_temp_desc,
				m_product_document_attribute_temp.prod_doc_attr_temp_type,
				m_product_document_attribute_temp.prod_doc_attr_temp_reason,
				m_product_document_attribute_temp.prod_doc_attr_temp_status,
				m_product_document_attribute_temp.prod_doc_attr_temp_created_by,
				m_product_document_attribute_temp.prod_doc_attr_temp_created_at`).
		Where("prod_doc_attr_temp_id = ?", pid).
		Find(&productdocumentcategorytemp).Error
	if err != nil {
		return &ProductDocumentAttributeTempPend{}, err
	}
	return &productdocumentcategorytemp, nil
}

func (p *ProductDocumentAttributeTemp) FindProductDocumentAttributeTempStatusByIDPendingActive(db *gorm.DB, pid uint64, status uint64) (*ProductDocumentAttributeTempPend, error) {
	var err error
	productdocumentcategorytemp := ProductDocumentAttributeTempPend{}
	err = db.Debug().Model(&ProductDocumentAttributeTempPend{}).Limit(100).
		Select(`m_product_document_attribute_temp.prod_doc_attr_temp_id,
				m_product_document_attribute_temp.address_id,
				m_product_document_attribute_temp.country_temp_id,
				m_product_document_attribute_temp.province_temp_id,
				m_product_document_attribute_temp.regency_temp_id,
				m_product_document_attribute_temp.district_temp_id,
				m_product_document_attribute_temp.village_temp_id,
				m_product_document_attribute_temp.address_temp_reason,
				m_product_document_attribute_temp.prod_doc_attr_temp_status,
				m_product_document_attribute_temp.address_temp_created_by,
				m_product_document_attribute_temp.address_temp_created_at`).
		Where("prod_doc_attr_temp_id = ? AND prod_doc_attr_temp_status = ?", pid, status).
		Find(&productdocumentcategorytemp).Error
	if err != nil {
		return &ProductDocumentAttributeTempPend{}, err
	}
	return &productdocumentcategorytemp, nil
}

func (p *ProductDocumentAttributeTemp) FindProductDocumentAttributeTempByID(db *gorm.DB, pid uint64) (*ProductDocumentAttributeTempData, error) {
	var err error
	productdocumentcategorytemp := ProductDocumentAttributeTempData{}
	err = db.Debug().Model(&ProductDocumentAttributeTempData{}).Limit(100).
		Select(`m_product_document_attribute_temp.prod_doc_attr_temp_id,
				m_product_document_attribute_temp.prod_doc_attr_temp_code,
				m_product_document_attribute_temp.prod_doc_attr_temp_name,
				m_product_document_attribute_temp.prod_doc_attr_temp_desc,
				m_product_document_attribute_temp.prod_doc_attr_temp_type,
				m_product_document_attribute_temp.prod_doc_attr_temp_reason,
				m_product_document_attribute_temp.prod_doc_attr_temp_status,
				m_product_document_attribute_temp.prod_doc_attr_temp_created_by,
				m_product_document_attribute_temp.prod_doc_attr_temp_created_at,
				m_product_document_attribute.prod_doc_attr_id,
				m_product_document_attribute.prod_doc_attr_code,
				m_product_document_attribute.prod_doc_attr_name,
				m_product_document_attribute.prod_doc_attr_desc,
				m_product_document_attribute.prod_doc_attr_type,
				m_product_document_attribute.prod_doc_attr_status,
				m_product_document_attribute.prod_doc_attr_created_by,
				m_product_document_attribute.prod_doc_attr_created_at,
				m_product_document_attribute.prod_doc_attr_updated_by,
				m_product_document_attribute.prod_doc_attr_updated_at,
				m_product_document_attribute.prod_doc_attr_deleted_by,
				m_product_document_attribute.prod_doc_attr_deleted_at`).
		Joins("JOIN m_product_document_attribute on m_product_document_attribute_temp.prod_doc_attr_id=m_product_document_attribute.prod_doc_attr_id").
		Where("prod_doc_attr_temp_id = ?", pid).
		Find(&productdocumentcategorytemp).Error
	if err != nil {
		return &ProductDocumentAttributeTempData{}, err
	}
	return &productdocumentcategorytemp, nil
}

func (p *ProductDocumentAttributeTemp) FindProductDocumentAttributeTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*ProductDocumentAttributeTempData, error) {
	var err error
	productdocumentcategorytemp := ProductDocumentAttributeTempData{}
	err = db.Debug().Model(&ProductDocumentAttributeTempData{}).Limit(100).
		Select(`m_product_document_attribute_temp.prod_doc_attr_temp_id,
				m_product_document_attribute_temp.prod_doc_attr_id,
				m_product_document_attribute_temp.prod_doc_attr_temp_code,
				m_product_document_attribute_temp.prod_doc_attr_temp_name,
				m_product_document_attribute_temp.prod_doc_attr_temp_desc,
				m_product_document_attribute_temp.prod_doc_attr_temp_type,
				m_product_document_attribute_temp.prod_doc_attr_temp_reason,
				m_product_document_attribute_temp.prod_doc_attr_temp_status,
				m_product_document_attribute_temp.prod_doc_attr_temp_created_by,
				m_product_document_attribute_temp.prod_doc_attr_temp_created_at,
				m_product_document_attribute.prod_doc_attr_code,
				m_product_document_attribute.prod_doc_attr_name,
				m_product_document_attribute.prod_doc_attr_desc,
				m_product_document_attribute.prod_doc_attr_type,
				m_product_document_attribute.prod_doc_attr_status,
				m_product_document_attribute.prod_doc_attr_created_by,
				m_product_document_attribute.prod_doc_attr_created_at,
				m_product_document_attribute.prod_doc_attr_updated_by,
				m_product_document_attribute.prod_doc_attr_updated_at,
				m_product_document_attribute.prod_doc_attr_deleted_by,
				m_product_document_attribute.prod_doc_attr_deleted_at`).
		Joins("JOIN m_product_document_attribute on m_product_document_attribute_temp.prod_doc_attr_id=m_product_document_attribute.prod_doc_attr_id").
		Where("prod_doc_attr_temp_id = ?", pid).Where("prod_doc_attr_temp_id = ? AND prod_doc_attr_temp_status = ?", pid, status).
		Find(&productdocumentcategorytemp).Error
	if err != nil {
		return &ProductDocumentAttributeTempData{}, err
	}
	return &productdocumentcategorytemp, nil
}

func (p *ProductDocumentAttributeTemp) UpdateProductDocumentAttributeTemp(db *gorm.DB, pid uint64) (*ProductDocumentAttributeTemp, error) {
	var err error
	db = db.Debug().Model(&ProductDocumentAttributeTemp{}).Where("prod_doc_attr_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_attr_temp_code":   p.ProductDocumentAttributeTempCode,
			"prod_doc_attr_temp_name":   p.ProductDocumentAttributeTempName,
			"prod_doc_attr_temp_desc":   p.ProductDocumentAttributeTempDesc,
			"prod_doc_attr_temp_type":   p.ProductDocumentAttributeTempType,
			"prod_doc_attr_temp_reason": p.ProductDocumentAttributeTempReason,
		},
	)
	err = db.Debug().Model(&ProductDocumentAttributeTemp{}).Where("prod_doc_attr_temp_id = ?", pid).Error
	if err != nil {
		return &ProductDocumentAttributeTemp{}, err
	}
	return p, nil
}

func (p *ProductDocumentAttributeTemp) UpdateProductDocumentAttributeTempStatus(db *gorm.DB, pid uint64) (*ProductDocumentAttributeTemp, error) {
	var err error
	db = db.Debug().Model(&ProductDocumentAttributeTemp{}).Where("prod_doc_attr_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_attr_temp_reason": p.ProductDocumentAttributeTempReason,
			"prod_doc_attr_temp_status": p.ProductDocumentAttributeTempStatus,
		},
	)
	err = db.Debug().Model(&ProductDocumentAttributeTemp{}).Where("prod_doc_attr_temp_id = ?", pid).Error
	if err != nil {
		return &ProductDocumentAttributeTemp{}, err
	}
	return p, nil
}

func (p *ProductDocumentAttributeTemp) DeleteProductDocumentAttributeTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&ProductDocumentAttributeTemp{}).Where("prod_doc_attr_temp_id = ?", pid).Take(&ProductDocumentAttributeTemp{}).Delete(&ProductDocumentAttributeTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Product Document Category Temp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
