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

type ProductDocumentTemp struct {
	ProductDocumentTempID           uint64    `gorm:"column:prod_doc_temp_id;primary_key;" json:"prod_doc_temp_id"`
	ProductDocumentID               uint64    `gorm:"column:prod_doc_id" json:"prod_doc_id"`
	ProductDocumentCategoryTempID   uint64    `gorm:"column:prod_doc_ctgry_temp_id" json:"prod_doc_ctgry_temp_id"`
	ProductDocumentTempCode         string    `gorm:"column:prod_doc_temp_code" json:"prod_doc_temp_code"`
	ProductDocumentTempName         string    `gorm:"column:prod_doc_temp_name" json:"prod_doc_temp_name"`
	ProductDocumentTempDescription  string    `gorm:"column:prod_doc_temp_desc" json:"prod_doc_temp_desc"`
	ProductDocumentTempReason       string    `gorm:"column:prod_doc_temp_reason" json:"prod_doc_temp_reason"`
	ProductDocumentTempStatus       uint64    `gorm:"column:prod_doc_temp_status;size:2" json:"prod_doc_temp_status"`
	ProductDocumentTempActionBefore uint64    `gorm:"column:prod_doc_temp_action_before;size:2" json:"prod_doc_temp_action_before"`
	ProductDocumentTempCreatedBy    string    `gorm:"column:prod_doc_temp_created_by;size:125" json:"prod_doc_temp_created_by"`
	ProductDocumentTempCreatedAt    time.Time `gorm:"column:prod_doc_temp_created_at;default:CURRENT_TIMESTAMP" json:"prod_doc_temp_created_at"`
}

type ProductDocumentTempPend struct {
	ProductDocumentTempID           uint64    `gorm:"column:prod_doc_temp_id;primary_key;" json:"prod_doc_temp_id"`
	ProductDocumentID               *uint64   `gorm:"column:prod_doc_id" json:"prod_doc_id"`
	ProductDocumentCategoryTempID   uint64    `gorm:"column:prod_doc_ctgry_temp_id" json:"prod_doc_ctgry_temp_id"`
	ProductDocumentTempCode         string    `gorm:"column:prod_doc_temp_code" json:"prod_doc_temp_code"`
	ProductDocumentTempName         string    `gorm:"column:prod_doc_temp_name" json:"prod_doc_temp_name"`
	ProductDocumentTempDescription  string    `gorm:"column:prod_doc_temp_desc" json:"prod_doc_temp_desc"`
	ProductDocumentTempReason       string    `gorm:"column:prod_doc_temp_reason" json:"prod_doc_temp_reason"`
	ProductDocumentTempStatus       uint64    `gorm:"column:prod_doc_temp_status;size:2" json:"prod_doc_temp_status"`
	ProductDocumentTempActionBefore uint64    `gorm:"column:prod_doc_temp_action_before;size:2" json:"prod_doc_temp_action_before"`
	ProductDocumentTempCreatedBy    string    `gorm:"column:prod_doc_temp_created_by;size:125" json:"prod_doc_temp_created_by"`
	ProductDocumentTempCreatedAt    time.Time `gorm:"column:prod_doc_temp_created_at;default:CURRENT_TIMESTAMP" json:"prod_doc_temp_created_at"`
}

type ProductDocumentTempData struct {
	ProductDocumentTempID           uint64    `gorm:"column:prod_doc_temp_id" json:"prod_doc_temp_id"`
	ProductDocumentCategoryTempID   uint64    `gorm:"column:prod_doc_ctgry_temp_id" json:"prod_doc_ctgry_temp_id"`
	ProductDocumentCategoryTempName string    `gorm:"column:prod_doc_ctgry_temp_name" json:"prod_doc_ctgry_temp_name"`
	ProductDocumentTempCode         string    `gorm:"column:prod_doc_temp_code" json:"prod_doc_temp_code"`
	ProductDocumentTempName         string    `gorm:"column:prod_doc_temp_name" json:"prod_doc_temp_name"`
	ProductDocumentTempDescription  string    `gorm:"column:prod_doc_temp_desc" json:"prod_doc_temp_desc"`
	ProductDocumentTempReason       string    `gorm:"column:prod_doc_temp_reason" json:"prod_doc_temp_reason"`
	ProductDocumentTempStatus       uint64    `gorm:"column:prod_doc_temp_status;size:2" json:"prod_doc_temp_status"`
	ProductDocumentTempActionBefore uint64    `gorm:"column:prod_doc_temp_action_before;size:2" json:"prod_doc_temp_action_before"`
	ProductDocumentTempCreatedBy    string    `gorm:"column:prod_doc_temp_created_by;size:125" json:"prod_doc_temp_created_by"`
	ProductDocumentTempCreatedAt    time.Time `gorm:"column:prod_doc_temp_created_at" json:"prod_doc_temp_created_at"`
	ProductDocumentID               uint64    `gorm:"column:prod_doc_id" json:"prod_doc_id"`
	ProductDocumentCategoryID       uint64    `gorm:"column:prod_doc_ctgry_id" json:"prod_doc_ctgry_id"`
	ProductDocumentCategoryName     string    `gorm:"column:prod_doc_ctgry_name" json:"prod_doc_ctgry_name"`
	ProductDocumentCode             string    `gorm:"column:prod_doc_code" json:"prod_doc_code"`
	ProductDocumentName             string    `gorm:"column:prod_doc_name" json:"prod_doc_name"`
	ProductDocumentDescription      string    `gorm:"column:prod_doc_description" json:"prod_doc_description"`
	ProductDocumentStatus           uint64    `gorm:"column:prod_doc_status" json:"prod_doc_status"`
	ProductDocumentCreatedBy        string    `gorm:"column:prod_doc_created_by" json:"prod_doc_created_by"`
	ProductDocumentCreatedAt        time.Time `gorm:"column:prod_doc_created_at" json:"prod_doc_created_at"`
	ProductDocumentUpdatedBy        string    `gorm:"column:prod_doc_updated_by" json:"prod_doc_updated_by"`
	ProductDocumentUpdatedAt        time.Time `gorm:"column:prod_doc_updated_at" json:"prod_doc_updated_at"`
	ProductDocumentDeletedBy        string    `gorm:"column:prod_doc_deleted_by" json:"prod_doc_deleted_by"`
	ProductDocumentDeletedAt        time.Time `gorm:"column:prod_doc_deleted_at" json:"prod_doc_deleted_at"`
}

type ResponseProductDocumentTemps struct {
	Status  int                        `json:"status"`
	Message string                     `json:"message"`
	Data    *[]ProductDocumentTempData `json:"data"`
}

type ResponseProductDocumentTemp struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    *ProductDocumentTempData `json:"data"`
}

type ResponseProductDocumentTempsPend struct {
	Status  int                        `json:"status"`
	Message string                     `json:"message"`
	Data    *[]ProductDocumentTempPend `json:"data"`
}

type ResponseProductDocumentTempPend struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    *ProductDocumentTempPend `json:"data"`
}

type ResponseProductDocumentTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (ProductDocumentTemp) TableName() string {
	return "m_product_document_temp"
}

func (ProductDocumentTempPend) TableName() string {
	return "m_product_document_temp"
}

func (ProductDocumentTempData) TableName() string {
	return "m_product_document_temp"
}

func (p *ProductDocumentTemp) Prepare() {
	p.ProductDocumentID = p.ProductDocumentID
	p.ProductDocumentCategoryTempID = p.ProductDocumentCategoryTempID
	p.ProductDocumentTempCode = p.ProductDocumentTempCode
	p.ProductDocumentTempName = p.ProductDocumentTempName
	p.ProductDocumentTempDescription = p.ProductDocumentTempDescription
	p.ProductDocumentTempReason = p.ProductDocumentTempReason
	p.ProductDocumentTempStatus = p.ProductDocumentTempStatus
	p.ProductDocumentTempCreatedBy = html.EscapeString(strings.TrimSpace(p.ProductDocumentTempCreatedBy))
	p.ProductDocumentTempCreatedAt = time.Now()
}

func (p *ProductDocumentTempPend) Prepare() {
	p.ProductDocumentID = nil
	p.ProductDocumentCategoryTempID = p.ProductDocumentCategoryTempID
	p.ProductDocumentTempCode = p.ProductDocumentTempCode
	p.ProductDocumentTempName = p.ProductDocumentTempName
	p.ProductDocumentTempDescription = p.ProductDocumentTempDescription
	p.ProductDocumentTempReason = p.ProductDocumentTempReason
	p.ProductDocumentTempStatus = p.ProductDocumentTempStatus
	p.ProductDocumentTempCreatedBy = html.EscapeString(strings.TrimSpace(p.ProductDocumentTempCreatedBy))
	p.ProductDocumentTempCreatedAt = time.Now()
}

func (p *ProductDocumentTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	case "insertupdate":
		if p.ProductDocumentCategoryTempID == 0 {
			return errors.New("Negara Wajib diisi")
		}
		if p.ProductDocumentTempCode == "" {
			return errors.New("Provinsi Wajib diisi")
		}
		if p.ProductDocumentTempName == "" {
			return errors.New("Kabupaten / Kota Wajib diisi")
		}
		return nil

	case "reason":
		if p.ProductDocumentTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.ProductDocumentTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	default:
		if p.ProductDocumentTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil
	}
}

func (p *ProductDocumentTemp) SaveProductDocumentTemp(db *gorm.DB) (*ProductDocumentTemp, error) {
	var err error
	err = db.Debug().Model(&ProductDocumentTemp{}).Create(&p).Error
	if err != nil {
		return &ProductDocumentTemp{}, err
	}
	return p, nil
}

func (p *ProductDocumentTempPend) SaveProductDocumentTempPend(db *gorm.DB) (*ProductDocumentTempPend, error) {
	var err error
	err = db.Debug().Model(&ProductDocumentTempPend{}).Create(&p).Error
	if err != nil {
		return &ProductDocumentTempPend{}, err
	}
	return p, nil
}

func (p *ProductDocumentTemp) FindAllProductDocumentTemp(db *gorm.DB) (*[]ProductDocumentTempPend, error) {
	var err error
	productdocument := []ProductDocumentTempPend{}
	err = db.Debug().Model(&ProductDocumentTempPend{}).Limit(100).
		Select(`m_product_document_temp.prod_doc_temp_id,
				m_product_document_temp.prod_doc_id,
				m_product_document_temp.prod_doc_ctgry_temp_id,
				m_product_document_temp.prod_doc_temp_code,
				m_product_document_temp.prod_doc_temp_name,
				m_product_document_temp.prod_doc_temp_desc,
				m_product_document_temp.prod_doc_temp_reason,
				m_product_document_temp.prod_doc_temp_status,
				m_product_document_temp.prod_doc_temp_action_before,
				m_product_document_temp.prod_doc_temp_created_by,
				m_product_document_temp.prod_doc_temp_created_at at time zone 'Asia/Jakarta' as prod_doc_temp_created_at`).
		Order("prod_doc_temp_created_at desc").
		Find(&productdocument).Error
	if err != nil {
		return &[]ProductDocumentTempPend{}, err
	}
	return &productdocument, nil
}

func (p *ProductDocumentTemp) FindAllProductDocumentTempStatusPendingActive(db *gorm.DB) (*[]ProductDocumentTempPend, error) {
	var err error
	productdocument := []ProductDocumentTempPend{}
	err = db.Debug().Model(&ProductDocumentTempPend{}).Limit(100).
		Select(`m_product_document_temp.prod_doc_temp_id,
				m_product_document_temp.prod_doc_id,
				m_product_document_temp.prod_doc_ctgry_temp_id,
				m_product_document_temp.prod_doc_temp_code,
				m_product_document_temp.prod_doc_temp_name,
				m_product_document_temp.prod_doc_temp_desc,
				m_product_document_temp.prod_doc_temp_reason,
				m_product_document_temp.prod_doc_temp_status,
				m_product_document_temp.prod_doc_temp_action_before,
				m_product_document_temp.prod_doc_temp_created_by,
				m_product_document_temp.prod_doc_temp_created_at at time zone 'Asia/Jakarta' as prod_doc_temp_created_at`).
		Where("prod_doc_temp_status = ?", 11).
		Order("prod_doc_temp_created_at desc").
		Find(&productdocument).Error
	if err != nil {
		return &[]ProductDocumentTempPend{}, err
	}
	return &productdocument, nil
}

func (p *ProductDocumentTemp) FindAllProductDocumentTempStatus(db *gorm.DB, status uint64) (*[]ProductDocumentTempData, error) {
	var err error
	productdocument := []ProductDocumentTempData{}
	err = db.Debug().Model(&ProductDocumentTempData{}).Limit(100).
		Select(`m_product_document_temp.prod_doc_temp_id,
				m_product_document_category_temp.prod_doc_ctgry_id as prod_doc_ctgry_temp_id,
				m_product_document_category_temp.prod_doc_ctgry_name as prod_doc_ctgry_temp_name,
				m_product_document_temp.prod_doc_temp_code,
				m_product_document_temp.prod_doc_temp_name,
				m_product_document_temp.prod_doc_temp_desc,
				m_product_document_temp.prod_doc_temp_reason,
				m_product_document_temp.prod_doc_temp_status,
				m_product_document_temp.prod_doc_temp_action_before,
				m_product_document_temp.prod_doc_temp_created_by,
				m_product_document_temp.prod_doc_temp_created_at at time zone 'Asia/Jakarta' as prod_doc_temp_created_at,
				m_product_document.prod_doc_id,
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
		Joins("JOIN m_product_document_category m_product_document_category_temp on m_product_document_temp.prod_doc_ctgry_temp_id=m_product_document_category_temp.prod_doc_ctgry_id").
		Joins("JOIN m_product_document on m_product_document_temp.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_document_category on m_product_document.prod_doc_ctgry_id=m_product_document_category.prod_doc_ctgry_id").
		Where("prod_doc_temp_status = ?", status).
		Order("prod_doc_temp_created_at desc").
		Find(&productdocument).Error
	if err != nil {
		return &[]ProductDocumentTempData{}, err
	}
	return &productdocument, nil
}

func (p *ProductDocumentTemp) FindProductDocumentTempDataByID(db *gorm.DB, pid uint64) (*ProductDocumentTemp, error) {
	var err error
	err = db.Debug().Model(&ProductDocumentTemp{}).Where("prod_doc_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &ProductDocumentTemp{}, err
	}
	return p, nil
}

func (p *ProductDocumentTemp) FindProductDocumentTempByIDPendingActive(db *gorm.DB, pid uint64) (*ProductDocumentTempPend, error) {
	var err error
	productdocument := ProductDocumentTempPend{}
	err = db.Debug().Model(&ProductDocumentTempPend{}).Limit(100).
		Select(`m_product_document_temp.prod_doc_temp_id,
				m_product_document_temp.prod_doc_id,
				m_product_document_temp.prod_doc_ctgry_temp_id,
				m_product_document_temp.prod_doc_temp_code,
				m_product_document_temp.prod_doc_temp_name,
				m_product_document_temp.prod_doc_temp_desc,
				m_product_document_temp.prod_doc_temp_reason,
				m_product_document_temp.prod_doc_temp_status,
				m_product_document_temp.prod_doc_temp_created_by,
				m_product_document_temp.prod_doc_temp_created_at at time zone 'Asia/Jakarta' as prod_doc_temp_created_at`).
		Where("prod_doc_temp_id = ?", pid).
		Order("prod_doc_temp_created_at desc").
		Find(&productdocument).Error
	if err != nil {
		return &ProductDocumentTempPend{}, err
	}
	return &productdocument, nil
}

func (p *ProductDocumentTemp) FindProductDocumentTempStatusByIDPendingActive(db *gorm.DB, pid uint64) (*ProductDocumentTempPend, error) {
	var err error
	productdocument := ProductDocumentTempPend{}
	err = db.Debug().Model(&ProductDocumentTempPend{}).Limit(100).
		Select(`m_product_document_temp.prod_doc_temp_id,
				m_product_document_temp.prod_doc_id,
				m_product_document_temp.prod_doc_ctgry_temp_id,
				m_product_document_temp.prod_doc_temp_code,
				m_product_document_temp.prod_doc_temp_name,
				m_product_document_temp.prod_doc_temp_desc,
				m_product_document_temp.prod_doc_temp_reason,
				m_product_document_temp.prod_doc_temp_status,
				m_product_document_temp.prod_doc_temp_action_before,
				m_product_document_temp.prod_doc_temp_created_by,
				m_product_document_temp.prod_doc_temp_created_at at time zone 'Asia/Jakarta' as prod_doc_temp_created_at`).
		Where("prod_doc_temp_id = ? AND prod_doc_temp_status = ?", pid, 11).
		Order("prod_doc_temp_created_at desc").
		Find(&productdocument).Error
	if err != nil {
		return &ProductDocumentTempPend{}, err
	}
	return &productdocument, nil
}

func (p *ProductDocumentTemp) FindProductDocumentTempByID(db *gorm.DB, pid uint64) (*ProductDocumentTempData, error) {
	var err error
	productdocument := ProductDocumentTempData{}
	err = db.Debug().Model(&ProductDocumentTempData{}).Limit(100).
		Select(`m_product_document_temp.prod_doc_temp_id,
				m_product_document_category_temp.prod_doc_ctgry_id as prod_doc_ctgry_temp_id,
				m_product_document_category_temp.prod_doc_ctgry_name as prod_doc_ctgry_temp_name,
				m_product_document_temp.prod_doc_temp_code,
				m_product_document_temp.prod_doc_temp_name,
				m_product_document_temp.prod_doc_temp_desc,
				m_product_document_temp.prod_doc_temp_reason,
				m_product_document_temp.prod_doc_temp_status,
				m_product_document_temp.prod_doc_temp_action_before,
				m_product_document_temp.prod_doc_temp_created_by,
				m_product_document_temp.prod_doc_temp_created_at at time zone 'Asia/Jakarta' as prod_doc_temp_created_at,
				m_product_document.prod_doc_id,
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
		Joins("JOIN m_product_document_category m_product_document_category_temp on m_product_document_temp.prod_doc_ctgry_temp_id=m_product_document_category_temp.prod_doc_ctgry_id").
		Joins("JOIN m_product_document on m_product_document_temp.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_document_category on m_product_document.prod_doc_ctgry_id=m_product_document_category.prod_doc_ctgry_id").
		Where("prod_doc_temp_id = ?", pid).
		Order("prod_doc_temp_created_at desc").
		Find(&productdocument).Error
	if err != nil {
		return &ProductDocumentTempData{}, err
	}
	return &productdocument, nil
}

func (p *ProductDocumentTemp) FindProductDocumentTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*ProductDocumentTempData, error) {
	var err error
	productdocument := ProductDocumentTempData{}
	err = db.Debug().Model(&ProductDocumentTempData{}).Limit(100).
		Select(`m_product_document_temp.prod_doc_temp_id,
				m_product_document_category_temp.prod_doc_ctgry_id as prod_doc_ctgry_temp_id,
				m_product_document_category_temp.prod_doc_ctgry_name as prod_doc_ctgry_temp_name,
				m_product_document_temp.prod_doc_temp_code,
				m_product_document_temp.prod_doc_temp_name,
				m_product_document_temp.prod_doc_temp_desc,
				m_product_document_temp.prod_doc_temp_reason,
				m_product_document_temp.prod_doc_temp_status,
				m_product_document_temp.prod_doc_temp_action_before,
				m_product_document_temp.prod_doc_temp_created_by,
				m_product_document_temp.prod_doc_temp_created_at at time zone 'Asia/Jakarta' as prod_doc_temp_created_at,
				m_product_document.prod_doc_id,
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
		Joins("JOIN m_product_document_category m_product_document_category_temp on m_product_document_temp.prod_doc_ctgry_temp_id=m_product_document_category_temp.prod_doc_ctgry_id").
		Joins("JOIN m_product_document on m_product_document_temp.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_document_category on m_product_document.prod_doc_ctgry_id=m_product_document_category.prod_doc_ctgry_id").
		Where("prod_doc_temp_id = ? AND prod_doc_temp_status = ?", pid, status).
		Order("prod_doc_temp_created_at desc").
		Find(&productdocument).Error
	if err != nil {
		return &ProductDocumentTempData{}, err
	}
	return &productdocument, nil
}

func (p *ProductDocumentTemp) UpdateProductDocumentTemp(db *gorm.DB, pid uint64) (*ProductDocumentTemp, error) {
	var err error
	db = db.Debug().Model(&ProductDocumentTemp{}).Where("prod_doc_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_ctgry_temp_id": p.ProductDocumentCategoryTempID,
			"prod_doc_temp_code":     p.ProductDocumentTempCode,
			"prod_doc_temp_name":     p.ProductDocumentTempName,
			"prod_doc_temp_desc":     p.ProductDocumentTempDescription,
			"prod_doc_temp_reason":   p.ProductDocumentTempReason,
		},
	)
	err = db.Debug().Model(&ProductDocumentTemp{}).Where("prod_doc_temp_id = ?", pid).Error
	if err != nil {
		return &ProductDocumentTemp{}, err
	}
	return p, nil
}

func (p *ProductDocumentTemp) UpdateProductDocumentTempStatus(db *gorm.DB, pid uint64) (*ProductDocumentTemp, error) {
	var err error
	db = db.Debug().Model(&ProductDocumentTemp{}).Where("prod_doc_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_temp_reason": p.ProductDocumentTempReason,
			"prod_doc_temp_status": p.ProductDocumentTempStatus,
		},
	)
	err = db.Debug().Model(&ProductDocumentTemp{}).Where("prod_doc_temp_id = ?", pid).Error
	if err != nil {
		return &ProductDocumentTemp{}, err
	}
	return p, nil
}

func (p *ProductDocumentTemp) DeleteProductDocumentTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&ProductDocumentTemp{}).Where("prod_doc_temp_id = ?", pid).Take(&ProductDocumentTemp{}).Delete(&ProductDocumentTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("ProductDocumentTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================
