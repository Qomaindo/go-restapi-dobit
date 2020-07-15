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

type ProductCategoryDocumentGroupTemp struct {
	ProductCategoryDocumentGroupTempID           uint64    `gorm:"column:prod_ctgry_doc_group_temp_id;primary_key;" json:"prod_ctgry_doc_group_temp_id"`
	ProductCategoryDocumentGroupID               uint64    `gorm:"column:prod_ctgry_doc_group_id" json:"prod_ctgry_doc_group_id"`
	ProductDocumentTempID                        uint64    `gorm:"column:prod_doc_temp_id" json:"prod_doc_temp_id"`
	ProductSubCategoryTempID                     uint64    `gorm:"column:prod_sub_ctgry_temp_id" json:"prod_sub_ctgry_temp_id"`
	ProductCategoryDocumentGroupTempReason       string    `gorm:"column:prod_ctgry_doc_group_temp_reason" json:"prod_ctgry_doc_group_temp_reason"`
	ProductCategoryDocumentGroupTempStatus       uint64    `gorm:"column:prod_ctgry_doc_group_temp_status;size:2" json:"prod_ctgry_doc_group_temp_status"`
	ProductCategoryDocumentGroupTempActionBefore uint64    `gorm:"column:prod_ctgry_doc_group_temp_action_before;size:2" json:"prod_ctgry_doc_group_temp_action_before"`
	ProductCategoryDocumentGroupTempCreatedBy    string    `gorm:"column:prod_ctgry_doc_group_temp_created_by;size:125" json:"prod_ctgry_doc_group_temp_created_by"`
	ProductCategoryDocumentGroupTempCreatedAt    time.Time `gorm:"column:prod_ctgry_doc_group_temp_created_at;default:CURRENT_TIMESTAMP" json:"prod_ctgry_doc_group_temp_created_at"`
}

type ProductCategoryDocumentGroupTempPend struct {
	ProductCategoryDocumentGroupTempID           uint64    `gorm:"column:prod_ctgry_doc_group_temp_id;primary_key;" json:"prod_ctgry_doc_group_temp_id"`
	ProductCategoryDocumentGroupID               *uint64   `gorm:"column:prod_ctgry_doc_group_id" json:"prod_ctgry_doc_group_id"`
	ProductDocumentTempID                        uint64    `gorm:"column:prod_doc_temp_id" json:"prod_doc_temp_id"`
	ProductSubCategoryTempID                     uint64    `gorm:"column:prod_sub_ctgry_temp_id" json:"prod_sub_ctgry_temp_id"`
	ProductCategoryDocumentGroupTempReason       string    `gorm:"column:prod_ctgry_doc_group_temp_reason" json:"prod_ctgry_doc_group_temp_reason"`
	ProductCategoryDocumentGroupTempStatus       uint64    `gorm:"column:prod_ctgry_doc_group_temp_status;size:2" json:"prod_ctgry_doc_group_temp_status"`
	ProductCategoryDocumentGroupTempActionBefore uint64    `gorm:"column:prod_ctgry_doc_group_temp_action_before;size:2" json:"prod_ctgry_doc_group_temp_action_before"`
	ProductCategoryDocumentGroupTempCreatedBy    string    `gorm:"column:prod_ctgry_doc_group_temp_created_by;size:125" json:"prod_ctgry_doc_group_temp_created_by"`
	ProductCategoryDocumentGroupTempCreatedAt    time.Time `gorm:"column:prod_ctgry_doc_group_temp_created_at;default:CURRENT_TIMESTAMP" json:"prod_ctgry_doc_group_temp_created_at"`
}

type ProductCategoryDocumentGroupTempData struct {
	ProductCategoryDocumentGroupTempID           uint64    `gorm:"column:prod_ctgry_doc_group_temp_id" json:"prod_ctgry_doc_group_temp_id"`
	CountryTempID                                uint64    `gorm:"column:country_temp_id" json:"country_temp_id"`
	ProductDocumentTempID                        uint64    `gorm:"column:prod_doc_temp_id" json:"prod_doc_temp_id"`
	ProductSubCategoryTempID                     uint64    `gorm:"column:prod_sub_ctgry_temp_id" json:"prod_sub_ctgry_temp_id"`
	ProductCategoryDocumentGroupTempReason       string    `gorm:"column:prod_ctgry_doc_group_temp_reason" json:"prod_ctgry_doc_group_temp_reason"`
	ProductCategoryDocumentGroupTempStatus       uint64    `gorm:"column:prod_ctgry_doc_group_temp_status;size:2" json:"prod_ctgry_doc_group_temp_status"`
	ProductCategoryDocumentGroupTempActionBefore uint64    `gorm:"column:prod_ctgry_doc_group_temp_action_before;size:2" json:"prod_ctgry_doc_group_temp_action_before"`
	ProductCategoryDocumentGroupTempCreatedBy    string    `gorm:"column:prod_ctgry_doc_group_temp_created_by;size:125" json:"prod_ctgry_doc_group_temp_created_by"`
	ProductCategoryDocumentGroupTempCreatedAt    time.Time `gorm:"column:prod_ctgry_doc_group_temp_created_at" json:"prod_ctgry_doc_group_temp_created_at"`
	ProductCategoryDocumentGroupID               uint64    `gorm:"column:prod_ctgry_doc_group_id" json:"prod_ctgry_doc_group_id"`
	ProductDocumentID                            uint64    `gorm:"column:prod_doc_id" json:"prod_doc_id"`
	ProductDocumentName                          string    `gorm:"column:prod_doc_name" json:"prod_doc_name"`
	ProductSubCategoryID                         uint64    `gorm:"column:prod_sub_ctgry_id" json:"prod_sub_ctgry_id"`
	ProductSubCategoryName                       string    `gorm:"column:prod_sub_ctgry_name" json:"prod_sub_ctgry_name"`
	ProductCategoryDocumentGroupStatus           uint64    `gorm:"column:prod_ctgry_doc_group_status" json:"prod_ctgry_doc_group_status"`
	ProductCategoryDocumentGroupCreatedBy        string    `gorm:"column:prod_ctgry_doc_group_created_by" json:"prod_ctgry_doc_group_created_by"`
	ProductCategoryDocumentGroupCreatedAt        time.Time `gorm:"column:prod_ctgry_doc_group_created_at" json:"prod_ctgry_doc_group_created_at"`
	ProductCategoryDocumentGroupUpdatedBy        string    `gorm:"column:prod_ctgry_doc_group_updated_by" json:"prod_ctgry_doc_group_updated_by"`
	ProductCategoryDocumentGroupUpdatedAt        time.Time `gorm:"column:prod_ctgry_doc_group_updated_at" json:"prod_ctgry_doc_group_updated_at"`
	ProductCategoryDocumentGroupDeletedBy        string    `gorm:"column:prod_ctgry_doc_group_deleted_by" json:"prod_ctgry_doc_group_deleted_by"`
	ProductCategoryDocumentGroupDeletedAt        time.Time `gorm:"column:prod_ctgry_doc_group_deleted_at" json:"prod_ctgry_doc_group_deleted_at"`
}

type ResponseProductCategoryDocumentGroupTemps struct {
	Status  int                                     `json:"status"`
	Message string                                  `json:"message"`
	Data    *[]ProductCategoryDocumentGroupTempData `json:"data"`
}

type ResponseProductCategoryDocumentGroupTemp struct {
	Status  int                                   `json:"status"`
	Message string                                `json:"message"`
	Data    *ProductCategoryDocumentGroupTempData `json:"data"`
}

type ResponseProductCategoryDocumentGroupTempsPend struct {
	Status  int                                     `json:"status"`
	Message string                                  `json:"message"`
	Data    *[]ProductCategoryDocumentGroupTempPend `json:"data"`
}

type ResponseProductCategoryDocumentGroupTempPend struct {
	Status  int                                   `json:"status"`
	Message string                                `json:"message"`
	Data    *ProductCategoryDocumentGroupTempPend `json:"data"`
}

type ResponseProductCategoryDocumentGroupTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (ProductCategoryDocumentGroupTemp) TableName() string {
	return "m_product_category_document_group_temp"
}

func (ProductCategoryDocumentGroupTempPend) TableName() string {
	return "m_product_category_document_group_temp"
}

func (ProductCategoryDocumentGroupTempData) TableName() string {
	return "m_product_category_document_group_temp"
}

func (p *ProductCategoryDocumentGroupTemp) Prepare() {
	p.ProductCategoryDocumentGroupID = p.ProductCategoryDocumentGroupID
	p.ProductDocumentTempID = p.ProductDocumentTempID
	p.ProductSubCategoryTempID = p.ProductSubCategoryTempID
	p.ProductCategoryDocumentGroupTempReason = p.ProductCategoryDocumentGroupTempReason
	p.ProductCategoryDocumentGroupTempStatus = p.ProductCategoryDocumentGroupTempStatus
	p.ProductCategoryDocumentGroupTempCreatedBy = html.EscapeString(strings.TrimSpace(p.ProductCategoryDocumentGroupTempCreatedBy))
	p.ProductCategoryDocumentGroupTempCreatedAt = time.Now()
}

func (p *ProductCategoryDocumentGroupTempPend) Prepare() {
	p.ProductCategoryDocumentGroupID = nil
	p.ProductDocumentTempID = p.ProductDocumentTempID
	p.ProductSubCategoryTempID = p.ProductSubCategoryTempID
	p.ProductCategoryDocumentGroupTempReason = p.ProductCategoryDocumentGroupTempReason
	p.ProductCategoryDocumentGroupTempStatus = p.ProductCategoryDocumentGroupTempStatus
	p.ProductCategoryDocumentGroupTempCreatedBy = html.EscapeString(strings.TrimSpace(p.ProductCategoryDocumentGroupTempCreatedBy))
	p.ProductCategoryDocumentGroupTempCreatedAt = time.Now()
}

func (p *ProductCategoryDocumentGroupTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	case "insertupdate":
		if p.ProductDocumentTempID == 0 {
			return errors.New("Negara Wajib diisi")
		}
		if p.ProductSubCategoryTempID == 0 {
			return errors.New("Provinsi Wajib diisi")
		}
		if p.ProductCategoryDocumentGroupTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.ProductCategoryDocumentGroupTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	case "reason":
		if p.ProductCategoryDocumentGroupTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.ProductCategoryDocumentGroupTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	default:
		if p.ProductCategoryDocumentGroupTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil
	}
}

func (p *ProductCategoryDocumentGroupTemp) SaveProductCategoryDocumentGroupTemp(db *gorm.DB) (*ProductCategoryDocumentGroupTemp, error) {
	var err error
	err = db.Debug().Model(&ProductCategoryDocumentGroupTemp{}).Create(&p).Error
	if err != nil {
		return &ProductCategoryDocumentGroupTemp{}, err
	}
	return p, nil
}

func (p *ProductCategoryDocumentGroupTempPend) SaveProductCategoryDocumentGroupTempPend(db *gorm.DB) (*ProductCategoryDocumentGroupTempPend, error) {
	var err error
	err = db.Debug().Model(&ProductCategoryDocumentGroupTempPend{}).Create(&p).Error
	if err != nil {
		return &ProductCategoryDocumentGroupTempPend{}, err
	}
	return p, nil
}

func (p *ProductCategoryDocumentGroupTemp) FindAllProductCategoryDocumentGroupTemp(db *gorm.DB) (*[]ProductCategoryDocumentGroupTempPend, error) {
	var err error
	address := []ProductCategoryDocumentGroupTempPend{}
	err = db.Debug().Model(&ProductCategoryDocumentGroupTempPend{}).Limit(100).
		Select(`m_product_category_document_group_temp.prod_ctgry_doc_group_temp_id,
				m_product_category_document_group_temp.prod_ctgry_doc_group_id,
				m_product_category_document_group_temp.prod_doc_temp_id,
				m_product_category_document_group_temp.prod_sub_ctgry_temp_id,
				m_product_category_document_group_temp.prod_ctgry_doc_group_temp_reason,
				m_product_category_document_group_temp.prod_ctgry_doc_group_temp_status,
				m_product_category_document_group_temp.prod_ctgry_doc_group_temp_action_before,
				m_product_category_document_group_temp.prod_ctgry_doc_group_temp_created_by,
				m_product_category_document_group_temp.prod_ctgry_doc_group_temp_created_at`).
		Find(&address).Error
	if err != nil {
		return &[]ProductCategoryDocumentGroupTempPend{}, err
	}
	return &address, nil
}

func (p *ProductCategoryDocumentGroupTemp) FindAllProductCategoryDocumentGroupTempPendingActive(db *gorm.DB) (*[]ProductCategoryDocumentGroupTempPend, error) {
	var err error
	address := []ProductCategoryDocumentGroupTempPend{}
	err = db.Debug().Model(&ProductCategoryDocumentGroupTempPend{}).Limit(100).
		Select(`m_product_category_document_group_temp.prod_ctgry_doc_group_temp_id,
				m_product_category_document_group_temp.prod_ctgry_doc_group_id,
				m_product_category_document_group_temp.prod_doc_temp_id,
				m_product_category_document_group_temp.prod_sub_ctgry_temp_id,
				m_product_category_document_group_temp.prod_ctgry_doc_group_temp_reason,
				m_product_category_document_group_temp.prod_ctgry_doc_group_temp_status,
				m_product_category_document_group_temp.prod_ctgry_doc_group_temp_action_before,
				m_product_category_document_group_temp.prod_ctgry_doc_group_temp_created_by,
				m_product_category_document_group_temp.prod_ctgry_doc_group_temp_created_at`).
		Where("prod_ctgry_doc_group_temp_status = ?", 11).
		Find(&address).Error
	if err != nil {
		return &[]ProductCategoryDocumentGroupTempPend{}, err
	}
	return &address, nil
}

func (p *ProductCategoryDocumentGroupTemp) FindAllProductCategoryDocumentGroupTempStatus(db *gorm.DB, status uint64) (*[]ProductCategoryDocumentGroupTempData, error) {
	var err error
	address := []ProductCategoryDocumentGroupTempData{}
	err = db.Debug().Model(&ProductCategoryDocumentGroupTempData{}).Limit(100).
		Select(`m_product_category_document_group_temp.prod_ctgry_doc_group_temp_id,
				m_product_document_temp.prod_doc_id as prod_doc_temp_id,
				m_product_document_temp.prod_doc_name as prod_doc_temp_name,
				m_product_sub_category_temp.prod_sub_ctgry_id as prod_sub_ctgry_temp_id,
				m_product_sub_category_temp.prod_sub_ctgry_name as prod_sub_ctgry_temp_name,
				m_product_category_document_group_temp.prod_ctgry_doc_group_temp_reason,
				m_product_category_document_group_temp.prod_ctgry_doc_group_temp_status,
				m_product_category_document_group_temp.prod_ctgry_doc_group_temp_action_before,
				m_product_category_document_group_temp.prod_ctgry_doc_group_temp_created_by,
				m_product_category_document_group_temp.prod_ctgry_doc_group_temp_created_at,
				m_product_category_document_group.prod_ctgry_doc_group_id,
				m_product_document.prod_doc_id,
				m_product_document.prod_doc_name,
				m_product_sub_category.prod_sub_ctgry_id,
				m_product_sub_category.prod_sub_ctgry_name,
				m_product_category_document_group.prod_ctgry_doc_group_status,
				m_product_category_document_group.prod_ctgry_doc_group_created_by,
				m_product_category_document_group.prod_ctgry_doc_group_created_at,
				m_product_category_document_group.prod_ctgry_doc_group_updated_by,
				m_product_category_document_group.prod_ctgry_doc_group_updated_at,
				m_product_category_document_group.prod_ctgry_doc_group_deleted_by,
				m_product_category_document_group.prod_ctgry_doc_group_deleted_at`).
		Joins("JOIN m_product_document m_product_document_temp on m_product_category_document_group_temp.prod_doc_temp_id=m_product_document_temp.prod_doc_id").
		Joins("JOIN m_product_sub_category m_product_sub_category_temp on m_product_category_document_group_temp.prod_sub_ctgry_temp_id=m_product_sub_category_temp.prod_sub_ctgry_id").
		Joins("JOIN m_product_category_document_group on m_product_category_document_group_temp.prod_ctgry_doc_group_id=m_product_category_document_group.prod_ctgry_doc_group_id").
		Joins("JOIN m_product_document on m_product_category_document_group.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_sub_category on m_product_category_document_group.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Where("prod_ctgry_doc_group_temp_status = ?", status).
		Find(&address).Error
	if err != nil {
		return &[]ProductCategoryDocumentGroupTempData{}, err
	}
	return &address, nil
}

func (p *ProductCategoryDocumentGroupTemp) FindProductCategoryDocumentGroupTempDataByID(db *gorm.DB, pid uint64) (*ProductCategoryDocumentGroupTemp, error) {
	var err error
	err = db.Debug().Model(&ProductCategoryDocumentGroupTemp{}).Where("prod_ctgry_doc_group_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &ProductCategoryDocumentGroupTemp{}, err
	}
	return p, nil
}

func (p *ProductCategoryDocumentGroupTemp) FindProductCategoryDocumentGroupTempByIDPendingActive(db *gorm.DB, pid uint64) (*ProductCategoryDocumentGroupTempPend, error) {
	var err error
	address := ProductCategoryDocumentGroupTempPend{}
	err = db.Debug().Model(&ProductCategoryDocumentGroupTempPend{}).Limit(100).
		Select(`m_product_category_document_group_temp.prod_ctgry_doc_group_temp_id,
				m_product_category_document_group_temp.prod_ctgry_doc_group_id,
				m_product_category_document_group_temp.prod_doc_temp_id,
				m_product_category_document_group_temp.prod_sub_ctgry_temp_id,
				m_product_category_document_group_temp.prod_ctgry_doc_group_temp_reason,
				m_product_category_document_group_temp.prod_ctgry_doc_group_temp_status,
				m_product_category_document_group_temp.prod_ctgry_doc_group_temp_action_before,
				m_product_category_document_group_temp.prod_ctgry_doc_group_temp_created_by,
				m_product_category_document_group_temp.prod_ctgry_doc_group_temp_created_at`).
		Where("prod_ctgry_doc_group_temp_id = ? AND prod_ctgry_doc_group_temp_status = ?", pid, 11).
		Find(&address).Error
	if err != nil {
		return &ProductCategoryDocumentGroupTempPend{}, err
	}
	return &address, nil
}

func (p *ProductCategoryDocumentGroupTemp) FindProductCategoryDocumentGroupTempByID(db *gorm.DB, pid uint64) (*ProductCategoryDocumentGroupTempData, error) {
	var err error
	address := ProductCategoryDocumentGroupTempData{}
	err = db.Debug().Model(&ProductCategoryDocumentGroupTempData{}).Limit(100).
		Select(`m_product_category_document_group_temp.prod_ctgry_doc_group_temp_id,
				m_product_document_temp.prod_doc_id as prod_doc_temp_id,
				m_product_document_temp.prod_doc_name as prod_doc_temp_name,
				m_product_sub_category_temp.prod_sub_ctgry_id as prod_sub_ctgry_temp_id,
				m_product_sub_category_temp.prod_sub_ctgry_name as prod_sub_ctgry_temp_name,
				m_product_category_document_group_temp.prod_ctgry_doc_group_temp_reason,
				m_product_category_document_group_temp.prod_ctgry_doc_group_temp_status,
				m_product_category_document_group_temp.prod_ctgry_doc_group_temp_action_before,
				m_product_category_document_group_temp.prod_ctgry_doc_group_temp_created_by,
				m_product_category_document_group_temp.prod_ctgry_doc_group_temp_created_at,
				m_product_category_document_group.prod_ctgry_doc_group_id,
				m_product_document.prod_doc_id,
				m_product_document.prod_doc_name,
				m_product_sub_category.prod_sub_ctgry_id,
				m_product_sub_category.prod_sub_ctgry_name,
				m_product_category_document_group.prod_ctgry_doc_group_status,
				m_product_category_document_group.prod_ctgry_doc_group_created_by,
				m_product_category_document_group.prod_ctgry_doc_group_created_at,
				m_product_category_document_group.prod_ctgry_doc_group_updated_by,
				m_product_category_document_group.prod_ctgry_doc_group_updated_at,
				m_product_category_document_group.prod_ctgry_doc_group_deleted_by,
				m_product_category_document_group.prod_ctgry_doc_group_deleted_at`).
		Joins("JOIN m_product_document m_product_document_temp on m_product_category_document_group_temp.prod_doc_temp_id=m_product_document_temp.prod_doc_id").
		Joins("JOIN m_product_sub_category m_product_sub_category_temp on m_product_category_document_group_temp.prod_sub_ctgry_temp_id=m_product_sub_category_temp.prod_sub_ctgry_id").
		Joins("JOIN m_product_category_document_group on m_product_category_document_group_temp.prod_ctgry_doc_group_id=m_product_category_document_group.prod_ctgry_doc_group_id").
		Joins("JOIN m_product_document on m_product_category_document_group.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_sub_category on m_product_category_document_group.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Where("prod_ctgry_doc_group_temp_id = ?", pid).
		Find(&address).Error
	if err != nil {
		return &ProductCategoryDocumentGroupTempData{}, err
	}
	return &address, nil
}

func (p *ProductCategoryDocumentGroupTemp) FindProductCategoryDocumentGroupTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*ProductCategoryDocumentGroupTempData, error) {
	var err error
	address := ProductCategoryDocumentGroupTempData{}
	err = db.Debug().Model(&ProductCategoryDocumentGroupTempData{}).Limit(100).
		Select(`m_product_category_document_group_temp.prod_ctgry_doc_group_temp_id,
				m_product_document_temp.prod_doc_id as prod_doc_temp_id,
				m_product_document_temp.prod_doc_name as prod_doc_temp_name,
				m_product_sub_category_temp.prod_sub_ctgry_id as prod_sub_ctgry_temp_id,
				m_product_sub_category_temp.prod_sub_ctgry_name as prod_sub_ctgry_temp_name,
				m_product_category_document_group_temp.prod_ctgry_doc_group_temp_reason,
				m_product_category_document_group_temp.prod_ctgry_doc_group_temp_status,
				m_product_category_document_group_temp.prod_ctgry_doc_group_temp_action_before,
				m_product_category_document_group_temp.prod_ctgry_doc_group_temp_created_by,
				m_product_category_document_group_temp.prod_ctgry_doc_group_temp_created_at,
				m_product_category_document_group.prod_ctgry_doc_group_id,
				m_product_document.prod_doc_id,
				m_product_document.prod_doc_name,
				m_product_sub_category.prod_sub_ctgry_id,
				m_product_sub_category.prod_sub_ctgry_name,
				m_product_category_document_group.prod_ctgry_doc_group_status,
				m_product_category_document_group.prod_ctgry_doc_group_created_by,
				m_product_category_document_group.prod_ctgry_doc_group_created_at,
				m_product_category_document_group.prod_ctgry_doc_group_updated_by,
				m_product_category_document_group.prod_ctgry_doc_group_updated_at,
				m_product_category_document_group.prod_ctgry_doc_group_deleted_by,
				m_product_category_document_group.prod_ctgry_doc_group_deleted_at`).
		Joins("JOIN m_product_document m_product_document_temp on m_product_category_document_group_temp.prod_doc_temp_id=m_product_document_temp.prod_doc_id").
		Joins("JOIN m_product_sub_category m_product_sub_category_temp on m_product_category_document_group_temp.prod_sub_ctgry_temp_id=m_product_sub_category_temp.prod_sub_ctgry_id").
		Joins("JOIN m_product_category_document_group on m_product_category_document_group_temp.prod_ctgry_doc_group_id=m_product_category_document_group.prod_ctgry_doc_group_id").
		Joins("JOIN m_product_document on m_product_category_document_group.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_sub_category on m_product_category_document_group.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Where("prod_ctgry_doc_group_temp_id = ? AND prod_ctgry_doc_group_temp_status = ?", pid, status).
		Find(&address).Error
	if err != nil {
		return &ProductCategoryDocumentGroupTempData{}, err
	}
	return &address, nil
}

func (p *ProductCategoryDocumentGroupTemp) UpdateProductCategoryDocumentGroupTemp(db *gorm.DB, pid uint64) (*ProductCategoryDocumentGroupTemp, error) {
	var err error
	db = db.Debug().Model(&ProductCategoryDocumentGroupTemp{}).Where("prod_ctgry_doc_group_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_temp_id":                 p.ProductDocumentTempID,
			"prod_sub_ctgry_temp_id":           p.ProductSubCategoryTempID,
			"prod_ctgry_doc_group_temp_reason": p.ProductCategoryDocumentGroupTempReason,
		},
	)
	err = db.Debug().Model(&ProductCategoryDocumentGroupTemp{}).Where("prod_ctgry_doc_group_temp_id = ?", pid).Error
	if err != nil {
		return &ProductCategoryDocumentGroupTemp{}, err
	}
	return p, nil
}

func (p *ProductCategoryDocumentGroupTemp) UpdateProductCategoryDocumentGroupTempStatus(db *gorm.DB, pid uint64) (*ProductCategoryDocumentGroupTemp, error) {
	var err error
	db = db.Debug().Model(&ProductCategoryDocumentGroupTemp{}).Where("prod_ctgry_doc_group_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_ctgry_doc_group_temp_reason": p.ProductCategoryDocumentGroupTempReason,
			"prod_ctgry_doc_group_temp_status": p.ProductCategoryDocumentGroupTempStatus,
		},
	)
	err = db.Debug().Model(&ProductCategoryDocumentGroupTemp{}).Where("prod_ctgry_doc_group_temp_id = ?", pid).Error
	if err != nil {
		return &ProductCategoryDocumentGroupTemp{}, err
	}
	return p, nil
}

func (p *ProductCategoryDocumentGroupTemp) DeleteProductCategoryDocumentGroupTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&ProductCategoryDocumentGroupTemp{}).Where("prod_ctgry_doc_group_temp_id = ?", pid).Take(&ProductCategoryDocumentGroupTemp{}).Delete(&ProductCategoryDocumentGroupTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("ProductCategoryDocumentGroupTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================
