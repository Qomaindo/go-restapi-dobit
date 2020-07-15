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

type ProductDocumentCategoryTemp struct {
	ProductDocumentCategoryTempID           uint64    `gorm:"column:prod_doc_ctgry_temp_id;primary_key;" json:"prod_doc_ctgry_temp_id"`
	ProductDocumentCategoryID               uint64    `gorm:"column:prod_doc_ctgry_id" json:"prod_doc_ctgry_id"`
	ProductDocumentCategoryTempCode         string    `gorm:"column:prod_doc_ctgry_temp_code" json:"prod_doc_ctgry_temp_code"`
	ProductDocumentCategoryTempName         string    `gorm:"column:prod_doc_ctgry_temp_name" json:"prod_doc_ctgry_temp_name"`
	ProductDocumentCategoryTempDesc         string    `gorm:"column:prod_doc_ctgry_temp_desc" json:"prod_doc_ctgry_temp_desc"`
	ProductDocumentCategoryTempReason       string    `gorm:"column:prod_doc_ctgry_temp_reason" json:"prod_doc_ctgry_temp_reason"`
	ProductDocumentCategoryTempStatus       uint64    `gorm:"column:prod_doc_ctgry_temp_status;size:2" json:"prod_doc_ctgry_temp_status"`
	ProductDocumentCategoryTempActionBefore uint64    `gorm:"column:prod_doc_ctgry_temp_action_before;size:2" json:"prod_doc_ctgry_temp_action_before"`
	ProductDocumentCategoryTempCreatedBy    string    `gorm:"column:prod_doc_ctgry_temp_created_by;size:125" json:"prod_doc_ctgry_temp_created_by"`
	ProductDocumentCategoryTempCreatedAt    time.Time `gorm:"column:prod_doc_ctgry_temp_created_at;default:CURRENT_TIMESTAMP" json:"prod_doc_ctgry_temp_created_at"`
}

type ProductDocumentCategoryTempPend struct {
	ProductDocumentCategoryTempID           uint64    `gorm:"column:prod_doc_ctgry_temp_id;primary_key;" json:"prod_doc_ctgry_temp_id"`
	ProductDocumentCategoryID               *uint64   `gorm:"column:prod_doc_ctgry_id" json:"prod_doc_ctgry_id"`
	ProductDocumentCategoryTempCode         string    `gorm:"column:prod_doc_ctgry_temp_code" json:"prod_doc_ctgry_temp_code"`
	ProductDocumentCategoryTempName         string    `gorm:"column:prod_doc_ctgry_temp_name" json:"prod_doc_ctgry_temp_name"`
	ProductDocumentCategoryTempDesc         string    `gorm:"column:prod_doc_ctgry_temp_desc" json:"prod_doc_ctgry_temp_desc"`
	ProductDocumentCategoryTempReason       string    `gorm:"column:prod_doc_ctgry_temp_reason" json:"prod_doc_ctgry_temp_reason"`
	ProductDocumentCategoryTempStatus       uint64    `gorm:"column:prod_doc_ctgry_temp_status;size:2" json:"prod_doc_ctgry_temp_status"`
	ProductDocumentCategoryTempActionBefore uint64    `gorm:"column:prod_doc_ctgry_temp_action_before;size:2" json:"prod_doc_ctgry_temp_action_before"`
	ProductDocumentCategoryTempCreatedBy    string    `gorm:"column:prod_doc_ctgry_temp_created_by;size:125" json:"prod_doc_ctgry_temp_created_by"`
	ProductDocumentCategoryTempCreatedAt    time.Time `gorm:"column:prod_doc_ctgry_temp_created_at;default:CURRENT_TIMESTAMP" json:"prod_doc_ctgry_temp_created_at"`
}

type ProductDocumentCategoryTempData struct {
	ProductDocumentCategoryTempID           uint64    `gorm:"column:prod_doc_ctgry_temp_id" json:"prod_doc_ctgry_temp_id"`
	ProductDocumentCategoryTempCode         string    `gorm:"column:prod_doc_ctgry_temp_code" json:"prod_doc_ctgry_temp_code"`
	ProductDocumentCategoryTempName         string    `gorm:"column:prod_doc_ctgry_temp_name" json:"prod_doc_ctgry_temp_name"`
	ProductDocumentCategoryTempDesc         string    `gorm:"column:prod_doc_ctgry_temp_desc" json:"prod_doc_ctgry_temp_desc"`
	ProductDocumentCategoryTempReason       string    `gorm:"column:prod_doc_ctgry_temp_reason" json:"prod_doc_ctgry_temp_reason"`
	ProductDocumentCategoryTempStatus       uint64    `gorm:"column:prod_doc_ctgry_temp_status;size:2" json:"prod_doc_ctgry_temp_status"`
	ProductDocumentCategoryTempActionBefore uint64    `gorm:"column:prod_doc_ctgry_temp_action_before;size:2" json:"prod_doc_ctgry_temp_action_before"`
	ProductDocumentCategoryTempCreatedBy    string    `gorm:"column:prod_doc_ctgry_temp_created_by;size:125" json:"prod_doc_ctgry_temp_created_by"`
	ProductDocumentCategoryTempCreatedAt    time.Time `gorm:"column:prod_doc_ctgry_temp_created_at" json:"prod_doc_ctgry_temp_created_at"`
	ProductDocumentCategoryID               uint64    `gorm:"column:prod_doc_ctgry_id" json:"prod_doc_ctgry_id"`
	ProductDocumentCategoryCode             string    `gorm:"column:prod_doc_ctgry_code" json:"prod_doc_ctgry_code"`
	ProductDocumentCategoryName             string    `gorm:"column:prod_doc_ctgry_name" json:"prod_doc_ctgry_name"`
	ProductDocumentCategoryDesc             string    `gorm:"column:prod_doc_ctgry_desc" json:"prod_doc_ctgry_desc"`
	ProductDocumentCategoryStatus           uint64    `gorm:"column:prod_doc_ctgry_status" json:"prod_doc_ctgry_status"`
	ProductDocumentCategoryCreatedBy        string    `gorm:"column:prod_doc_ctgry_created_by" json:"prod_doc_ctgry_created_by"`
	ProductDocumentCategoryCreatedAt        time.Time `gorm:"column:prod_doc_ctgry_created_at" json:"prod_doc_ctgry_created_at"`
	ProductDocumentCategoryUpdatedBy        string    `gorm:"column:prod_doc_ctgry_updated_by" json:"prod_doc_ctgry_updated_by"`
	ProductDocumentCategoryUpdatedAt        time.Time `gorm:"column:prod_doc_ctgry_updated_at" json:"prod_doc_ctgry_updated_at"`
	ProductDocumentCategoryDeletedBy        string    `gorm:"column:prod_doc_ctgry_deleted_by" json:"prod_doc_ctgry_deleted_by"`
	ProductDocumentCategoryDeletedAt        time.Time `gorm:"column:prod_doc_ctgry_deleted_at" json:"prod_doc_ctgry_deleted_at"`
}

type ResponseProductDocumentCategoryTemps struct {
	Status  int                                `json:"status"`
	Message string                             `json:"message"`
	Data    *[]ProductDocumentCategoryTempData `json:"data"`
}

type ResponseProductDocumentCategoryTemp struct {
	Status  int                              `json:"status"`
	Message string                           `json:"message"`
	Data    *ProductDocumentCategoryTempData `json:"data"`
}

type ResponseProductDocumentCategoryTempsPend struct {
	Status  int                                `json:"status"`
	Message string                             `json:"message"`
	Data    *[]ProductDocumentCategoryTempPend `json:"data"`
}

type ResponseProductDocumentCategoryTempPend struct {
	Status  int                              `json:"status"`
	Message string                           `json:"message"`
	Data    *ProductDocumentCategoryTempPend `json:"data"`
}

type ResponseProductDocumentCategoryTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (ProductDocumentCategoryTemp) TableName() string {
	return "m_product_document_category_temp"
}

func (ProductDocumentCategoryTempPend) TableName() string {
	return "m_product_document_category_temp"
}

func (ProductDocumentCategoryTempData) TableName() string {
	return "m_product_document_category_temp"
}

func (p *ProductDocumentCategoryTemp) Prepare() {
	p.ProductDocumentCategoryID = p.ProductDocumentCategoryID
	p.ProductDocumentCategoryTempCode = html.EscapeString(strings.TrimSpace(p.ProductDocumentCategoryTempCode))
	p.ProductDocumentCategoryTempName = html.EscapeString(strings.TrimSpace(p.ProductDocumentCategoryTempName))
	p.ProductDocumentCategoryTempDesc = p.ProductDocumentCategoryTempDesc
	p.ProductDocumentCategoryTempReason = p.ProductDocumentCategoryTempReason
	p.ProductDocumentCategoryTempStatus = p.ProductDocumentCategoryTempStatus
	p.ProductDocumentCategoryTempCreatedBy = html.EscapeString(strings.TrimSpace(p.ProductDocumentCategoryTempCreatedBy))
	p.ProductDocumentCategoryTempCreatedAt = time.Now()
}

func (p *ProductDocumentCategoryTempPend) Prepare() {
	p.ProductDocumentCategoryID = nil
	p.ProductDocumentCategoryTempCode = html.EscapeString(strings.TrimSpace(p.ProductDocumentCategoryTempCode))
	p.ProductDocumentCategoryTempName = html.EscapeString(strings.TrimSpace(p.ProductDocumentCategoryTempName))
	p.ProductDocumentCategoryTempDesc = p.ProductDocumentCategoryTempDesc
	p.ProductDocumentCategoryTempReason = p.ProductDocumentCategoryTempReason
	p.ProductDocumentCategoryTempStatus = p.ProductDocumentCategoryTempStatus
	p.ProductDocumentCategoryTempCreatedBy = html.EscapeString(strings.TrimSpace(p.ProductDocumentCategoryTempCreatedBy))
	p.ProductDocumentCategoryTempCreatedAt = time.Now()
}

func (p *ProductDocumentCategoryTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	case "insertupdate":
		if p.ProductDocumentCategoryTempCode == "" {
			return errors.New("Negara Wajib diisi")
		}
		if p.ProductDocumentCategoryTempName == "" {
			return errors.New("Provinsi Wajib diisi")
		}
		if p.ProductDocumentCategoryTempDesc == "" {
			return errors.New("Kabupaten / Kota Wajib diisi")
		}
		return nil

	case "reason":
		if p.ProductDocumentCategoryTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.ProductDocumentCategoryTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	default:
		if p.ProductDocumentCategoryTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil
	}
}

func (p *ProductDocumentCategoryTemp) SaveProductDocumentCategoryTemp(db *gorm.DB) (*ProductDocumentCategoryTemp, error) {
	var err error
	err = db.Debug().Model(&ProductDocumentCategoryTemp{}).Create(&p).Error
	if err != nil {
		return &ProductDocumentCategoryTemp{}, err
	}
	return p, nil
}

func (p *ProductDocumentCategoryTempPend) SaveProductDocumentCategoryTempPend(db *gorm.DB) (*ProductDocumentCategoryTempPend, error) {
	var err error
	err = db.Debug().Model(&ProductDocumentCategoryTempPend{}).Create(&p).Error
	if err != nil {
		return &ProductDocumentCategoryTempPend{}, err
	}
	return p, nil
}

func (p *ProductDocumentCategoryTemp) FindAllProductDocumentCategoryTemp(db *gorm.DB) (*[]ProductDocumentCategoryTempPend, error) {
	var err error
	productdocumentcategory := []ProductDocumentCategoryTempPend{}
	err = db.Debug().Model(&ProductDocumentCategoryTempPend{}).Limit(100).
		Select(`m_product_document_category_temp.prod_doc_ctgry_temp_id,
				m_product_document_category_temp.prod_doc_ctgry_id,
				m_product_document_category_temp.prod_doc_ctgry_temp_code,
				m_product_document_category_temp.prod_doc_ctgry_temp_name,
				m_product_document_category_temp.prod_doc_ctgry_temp_desc,
				m_product_document_category_temp.prod_doc_ctgry_temp_reason,
				m_product_document_category_temp.prod_doc_ctgry_temp_status,
				m_product_document_category_temp.prod_doc_ctgry_temp_action_before,
				m_product_document_category_temp.prod_doc_ctgry_temp_created_by,
				m_product_document_category_temp.prod_doc_ctgry_temp_created_at at time zone 'Asia/Jakarta' as prod_doc_ctgry_temp_created_at`).
		Order("prod_doc_ctgry_temp_created_at desc").
		Find(&productdocumentcategory).Error
	if err != nil {
		return &[]ProductDocumentCategoryTempPend{}, err
	}
	return &productdocumentcategory, nil
}

func (p *ProductDocumentCategoryTemp) FindAllProductDocumentCategoryTempPendingActive(db *gorm.DB) (*[]ProductDocumentCategoryTempPend, error) {
	var err error
	productdocumentcategory := []ProductDocumentCategoryTempPend{}
	err = db.Debug().Model(&ProductDocumentCategoryTempPend{}).Limit(100).
		Select(`m_product_document_category_temp.prod_doc_ctgry_temp_id,
				m_product_document_category_temp.prod_doc_ctgry_id,
				m_product_document_category_temp.prod_doc_ctgry_temp_code,
				m_product_document_category_temp.prod_doc_ctgry_temp_name,
				m_product_document_category_temp.prod_doc_ctgry_temp_desc,
				m_product_document_category_temp.prod_doc_ctgry_temp_reason,
				m_product_document_category_temp.prod_doc_ctgry_temp_status,
				m_product_document_category_temp.prod_doc_ctgry_temp_action_before,
				m_product_document_category_temp.prod_doc_ctgry_temp_created_by,
				m_product_document_category_temp.prod_doc_ctgry_temp_created_at at time zone 'Asia/Jakarta' as prod_doc_ctgry_temp_created_at`).
		Where("prod_doc_ctgry_temp_status = ?", 11).
		Order("prod_doc_ctgry_temp_created_at desc").
		Find(&productdocumentcategory).Error
	if err != nil {
		return &[]ProductDocumentCategoryTempPend{}, err
	}
	return &productdocumentcategory, nil
}

func (p *ProductDocumentCategoryTemp) FindAllProductDocumentCategoryTempStatus(db *gorm.DB, status uint64) (*[]ProductDocumentCategoryTempData, error) {
	var err error
	productdocumentcategory := []ProductDocumentCategoryTempData{}
	err = db.Debug().Model(&ProductDocumentCategoryTempData{}).Limit(100).
		Select(`m_product_document_category_temp.prod_doc_ctgry_temp_id,
				m_product_document_category_temp.prod_doc_ctgry_temp_code,
				m_product_document_category_temp.prod_doc_ctgry_temp_name,
				m_product_document_category_temp.prod_doc_ctgry_temp_desc,
				m_product_document_category_temp.prod_doc_ctgry_temp_reason,
				m_product_document_category_temp.prod_doc_ctgry_temp_status,
				m_product_document_category_temp.prod_doc_ctgry_temp_action_before,
				m_product_document_category_temp.prod_doc_ctgry_temp_created_by,
				m_product_document_category_temp.prod_doc_ctgry_temp_created_at at time zone 'Asia/Jakarta' as prod_doc_ctgry_temp_created_at,
				m_product_document_category.prod_doc_ctgry_id,
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
		Joins("JOIN m_product_document_category on m_product_document_category_temp.prod_doc_ctgry_id=m_product_document_category.prod_doc_ctgry_id").
		Where("prod_doc_ctgry_temp_status = ?", status).
		Order("prod_doc_ctgry_temp_created_at desc").
		Find(&productdocumentcategory).Error
	if err != nil {
		return &[]ProductDocumentCategoryTempData{}, err
	}
	return &productdocumentcategory, nil
}

func (p *ProductDocumentCategoryTemp) FindProductDocumentCategoryTempDataByID(db *gorm.DB, pid uint64) (*ProductDocumentCategoryTemp, error) {
	var err error
	err = db.Debug().Model(&ProductDocumentCategoryTemp{}).Where("prod_doc_ctgry_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &ProductDocumentCategoryTemp{}, err
	}
	return p, nil
}

func (p *ProductDocumentCategoryTemp) FindProductDocumentCategoryTempByIDPendingActive(db *gorm.DB, pid uint64) (*ProductDocumentCategoryTempPend, error) {
	var err error
	productdocumentcategory := ProductDocumentCategoryTempPend{}
	err = db.Debug().Model(&ProductDocumentCategoryTempPend{}).Limit(100).
		Select(`m_product_document_category_temp.prod_doc_ctgry_temp_id,
				m_product_document_category_temp.prod_doc_ctgry_id,
				m_product_document_category_temp.prod_doc_ctgry_temp_code,
				m_product_document_category_temp.prod_doc_ctgry_temp_name,
				m_product_document_category_temp.prod_doc_ctgry_temp_desc,
				m_product_document_category_temp.prod_doc_ctgry_temp_reason,
				m_product_document_category_temp.prod_doc_ctgry_temp_status,
				m_product_document_category_temp.prod_doc_ctgry_temp_created_by,
				m_product_document_category_temp.prod_doc_ctgry_temp_created_at at time zone 'Asia/Jakarta' as prod_doc_ctgry_temp_created_at`).
		Where("prod_doc_ctgry_temp_id = ?", pid).
		Order("prod_doc_ctgry_temp_created_at desc").
		Find(&productdocumentcategory).Error
	if err != nil {
		return &ProductDocumentCategoryTempPend{}, err
	}
	return &productdocumentcategory, nil
}

func (p *ProductDocumentCategoryTemp) FindProductDocumentCategoryTempByID(db *gorm.DB, pid uint64) (*ProductDocumentCategoryTempData, error) {
	var err error
	productdocumentcategory := ProductDocumentCategoryTempData{}
	err = db.Debug().Model(&ProductDocumentCategoryTempData{}).Limit(100).
		Select(`m_product_document_category_temp.prod_doc_ctgry_temp_id,
				m_product_document_category_temp.prod_doc_ctgry_temp_code,
				m_product_document_category_temp.prod_doc_ctgry_temp_name,
				m_product_document_category_temp.prod_doc_ctgry_temp_desc,
				m_product_document_category_temp.prod_doc_ctgry_temp_reason,
				m_product_document_category_temp.prod_doc_ctgry_temp_status,
				m_product_document_category_temp.prod_doc_ctgry_temp_action_before,
				m_product_document_category_temp.prod_doc_ctgry_temp_created_by,
				m_product_document_category_temp.prod_doc_ctgry_temp_created_at at time zone 'Asia/Jakarta' as prod_doc_ctgry_temp_created_at,
				m_product_document_category.prod_doc_ctgry_id,
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
		Joins("JOIN m_product_document_category on m_product_document_category_temp.prod_doc_ctgry_id=m_product_document_category.prod_doc_ctgry_id").
		Where("prod_doc_ctgry_temp_id = ?", pid).
		Order("prod_doc_ctgry_temp_created_at desc").
		Find(&productdocumentcategory).Error
	if err != nil {
		return &ProductDocumentCategoryTempData{}, err
	}
	return &productdocumentcategory, nil
}

func (p *ProductDocumentCategoryTemp) FindProductDocumentCategoryTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*ProductDocumentCategoryTempData, error) {
	var err error
	productdocumentcategory := ProductDocumentCategoryTempData{}
	err = db.Debug().Model(&ProductDocumentCategoryTempData{}).Limit(100).
		Select(`m_product_document_category_temp.prod_doc_ctgry_temp_id,
				m_product_document_category_temp.prod_doc_ctgry_temp_code,
				m_product_document_category_temp.prod_doc_ctgry_temp_name,
				m_product_document_category_temp.prod_doc_ctgry_temp_desc,
				m_product_document_category_temp.prod_doc_ctgry_temp_reason,
				m_product_document_category_temp.prod_doc_ctgry_temp_status,
				m_product_document_category_temp.prod_doc_ctgry_temp_action_before,
				m_product_document_category_temp.prod_doc_ctgry_temp_created_by,
				m_product_document_category_temp.prod_doc_ctgry_temp_created_at at time zone 'Asia/Jakarta' as prod_doc_ctgry_temp_created_at,
				m_product_document_category.prod_doc_ctgry_id,
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
		Joins("JOIN m_product_document_category on m_product_document_category_temp.prod_doc_ctgry_id=m_product_document_category.prod_doc_ctgry_id").
		Where("prod_doc_ctgry_temp_id = ? AND prod_doc_ctgry_temp_status = ?", pid, status).
		Order("prod_doc_ctgry_temp_created_at desc").
		Find(&productdocumentcategory).Error
	if err != nil {
		return &ProductDocumentCategoryTempData{}, err
	}
	return &productdocumentcategory, nil
}

func (p *ProductDocumentCategoryTemp) UpdateProductDocumentCategoryTemp(db *gorm.DB, pid uint64) (*ProductDocumentCategoryTemp, error) {
	var err error
	db = db.Debug().Model(&ProductDocumentCategoryTemp{}).Where("prod_doc_ctgry_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_ctgry_temp_code":   p.ProductDocumentCategoryTempCode,
			"prod_doc_ctgry_temp_name":   p.ProductDocumentCategoryTempName,
			"prod_doc_ctgry_temp_desc":   p.ProductDocumentCategoryTempDesc,
			"prod_doc_ctgry_temp_reason": p.ProductDocumentCategoryTempReason,
		},
	)
	err = db.Debug().Model(&ProductDocumentCategoryTemp{}).Where("prod_doc_ctgry_temp_id = ?", pid).Error
	if err != nil {
		return &ProductDocumentCategoryTemp{}, err
	}
	return p, nil
}

func (p *ProductDocumentCategoryTemp) UpdateProductDocumentCategoryTempStatus(db *gorm.DB, pid uint64) (*ProductDocumentCategoryTemp, error) {
	var err error
	db = db.Debug().Model(&ProductDocumentCategoryTemp{}).Where("prod_doc_ctgry_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_ctgry_temp_reason": p.ProductDocumentCategoryTempReason,
			"prod_doc_ctgry_temp_status": p.ProductDocumentCategoryTempStatus,
		},
	)
	err = db.Debug().Model(&ProductDocumentCategoryTemp{}).Where("prod_doc_ctgry_temp_id = ?", pid).Error
	if err != nil {
		return &ProductDocumentCategoryTemp{}, err
	}
	return p, nil
}

func (p *ProductDocumentCategoryTemp) DeleteProductDocumentCategoryTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&ProductDocumentCategoryTemp{}).Where("prod_doc_ctgry_temp_id = ?", pid).Take(&ProductDocumentCategoryTemp{}).Delete(&ProductDocumentCategoryTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("ProductDocumentCategoryTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================
