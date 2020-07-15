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

type ProductCategoryDocumentGroup struct {
	ProductCategoryDocumentGroupID        uint64     `gorm:"column:prod_ctgry_doc_group_id;primary_key;" json:"prod_ctgry_doc_group_id"`
	ProductDocumentID                     uint64     `gorm:"column:prod_doc_id" json:"prod_doc_id"`
	ProductSubCategoryID                  uint64     `gorm:"column:prod_sub_ctgry_id" json:"prod_sub_ctgry_id"`
	ProductCategoryDocumentGroupStatus    uint64     `gorm:"column:prod_ctgry_doc_group_status;size:2" json:"prod_ctgry_doc_group_status"`
	ProductCategoryDocumentGroupCreatedBy string     `gorm:"column:prod_ctgry_doc_group_created_by;size:125" json:"prod_ctgry_doc_group_created_by"`
	ProductCategoryDocumentGroupCreatedAt time.Time  `gorm:"column:prod_ctgry_doc_group_created_at;default:CURRENT_TIMESTAMP" json:"prod_ctgry_doc_group_created_at"`
	ProductCategoryDocumentGroupUpdatedBy string     `gorm:"column:prod_ctgry_doc_group_updated_by;size:125" json:"prod_ctgry_doc_group_updated_by"`
	ProductCategoryDocumentGroupUpdatedAt *time.Time `gorm:"column:prod_ctgry_doc_group_updated_at;default:CURRENT_TIMESTAMP" json:"prod_ctgry_doc_group_created_at"`
	ProductCategoryDocumentGroupDeletedBy string     `gorm:"column:prod_ctgry_doc_group_deleted_by;size:125" json:"prod_ctgry_doc_group_deleted_by"`
	ProductCategoryDocumentGroupDeletedAt *time.Time `gorm:"column:prod_ctgry_doc_group_deleted_at;default:CURRENT_TIMESTAMP" json:"prod_ctgry_doc_group_deleted_at"`
}

type ProductCategoryDocumentGroupData struct {
	ProductCategoryDocumentGroupID        uint64     `gorm:"column:prod_ctgry_doc_group_id" json:"prod_ctgry_doc_group_id"`
	ProductDocumentID                     uint64     `gorm:"column:prod_doc_id" json:"prod_doc_id"`
	ProductDocumentName                   uint64     `gorm:"column:prod_doc_name" json:"prod_doc_name"`
	ProductSubCategoryID                  uint64     `gorm:"column:prod_sub_ctgry_id" json:"prod_sub_ctgry_id"`
	ProductSubCategoryName                uint64     `gorm:"column:prod_sub_ctgry_name" json:"prod_sub_ctgry_name"`
	ProductCategoryDocumentGroupStatus    uint64     `gorm:"column:prod_ctgry_doc_group_status" json:"prod_ctgry_doc_group_status"`
	ProductCategoryDocumentGroupCreatedBy string     `gorm:"column:prod_ctgry_doc_group_created_by" json:"prod_ctgry_doc_group_created_by"`
	ProductCategoryDocumentGroupCreatedAt time.Time  `gorm:"column:prod_ctgry_doc_group_created_at" json:"prod_ctgry_doc_group_created_at"`
	ProductCategoryDocumentGroupUpdatedBy string     `gorm:"column:prod_ctgry_doc_group_updated_by" json:"prod_ctgry_doc_group_updated_by"`
	ProductCategoryDocumentGroupUpdatedAt *time.Time `gorm:"column:prod_ctgry_doc_group_updated_at" json:"prod_ctgry_doc_group_updated_at"`
	ProductCategoryDocumentGroupDeletedBy string     `gorm:"column:prod_ctgry_doc_group_deleted_by" json:"prod_ctgry_doc_group_deleted_by"`
	ProductCategoryDocumentGroupDeletedAt *time.Time `gorm:"column:prod_ctgry_doc_group_deleted_at" json:"prod_ctgry_doc_group_deleted_at"`
}

type ResponseProductCategoryDocumentGroups struct {
	Status  int                                 `json:"status"`
	Message string                              `json:"message"`
	Data    *[]ProductCategoryDocumentGroupData `json:"data"`
}

type ResponseProductCategoryDocumentGroup struct {
	Status  int                               `json:"status"`
	Message string                            `json:"message"`
	Data    *ProductCategoryDocumentGroupData `json:"data"`
}

type ResponseProductCategoryDocumentGroupCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (ProductCategoryDocumentGroup) TableName() string {
	return "m_product_category_document_group"
}

func (ProductCategoryDocumentGroupData) TableName() string {
	return "m_product_category_document_group"
}

func (p *ProductCategoryDocumentGroup) Prepare() {
	p.ProductDocumentID = p.ProductDocumentID
	p.ProductSubCategoryID = p.ProductSubCategoryID
	p.ProductCategoryDocumentGroupStatus = p.ProductCategoryDocumentGroupStatus
	p.ProductCategoryDocumentGroupCreatedBy = html.EscapeString(strings.TrimSpace(p.ProductCategoryDocumentGroupCreatedBy))
	p.ProductCategoryDocumentGroupCreatedAt = time.Now()
	p.ProductCategoryDocumentGroupUpdatedBy = html.EscapeString(strings.TrimSpace(p.ProductCategoryDocumentGroupUpdatedBy))
	p.ProductCategoryDocumentGroupUpdatedAt = p.ProductCategoryDocumentGroupUpdatedAt
	p.ProductCategoryDocumentGroupDeletedBy = html.EscapeString(strings.TrimSpace(p.ProductCategoryDocumentGroupDeletedBy))
	p.ProductCategoryDocumentGroupDeletedAt = p.ProductCategoryDocumentGroupDeletedAt
}

func (p *ProductCategoryDocumentGroup) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.ProductDocumentID == 0 {
			return errors.New("Required Country")
		}
		if p.ProductSubCategoryID == 0 {
			return errors.New("Required Province")
		}
		return nil
	}
}

func (p *ProductCategoryDocumentGroup) SaveProductCategoryDocumentGroup(db *gorm.DB) (*ProductCategoryDocumentGroup, error) {
	var err error
	err = db.Debug().Model(&ProductCategoryDocumentGroup{}).Create(&p).Error
	if err != nil {
		return &ProductCategoryDocumentGroup{}, err
	}
	return p, nil
}

func (p *ProductCategoryDocumentGroup) FindAllProductCategoryDocumentGroup(db *gorm.DB) (*[]ProductCategoryDocumentGroupData, error) {
	var err error
	address := []ProductCategoryDocumentGroupData{}
	err = db.Debug().Model(&ProductCategoryDocumentGroupData{}).Limit(100).
		Select(`m_product_category_document_group.prod_ctgry_doc_group_id,
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
		Joins("JOIN m_product_document on m_product_category_document_group.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_sub_category on m_product_category_document_group.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Find(&address).Error
	if err != nil {
		return &[]ProductCategoryDocumentGroupData{}, err
	}
	return &address, nil
}

func (p *ProductCategoryDocumentGroup) FindAllProductCategoryDocumentGroupStatus(db *gorm.DB, status uint64) (*[]ProductCategoryDocumentGroupData, error) {
	var err error
	address := []ProductCategoryDocumentGroupData{}
	err = db.Debug().Model(&ProductCategoryDocumentGroupData{}).Limit(100).
		Select(`m_product_category_document_group.prod_ctgry_doc_group_id,
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
		Joins("JOIN m_product_document on m_product_category_document_group.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_sub_category on m_product_category_document_group.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Where("prod_ctgry_doc_group_status = ?", status).
		Find(&address).Error
	if err != nil {
		return &[]ProductCategoryDocumentGroupData{}, err
	}
	return &address, nil
}

func (p *ProductCategoryDocumentGroup) FindProductCategoryDocumentGroupDataByID(db *gorm.DB, pid uint64) (*ProductCategoryDocumentGroup, error) {
	var err error
	err = db.Debug().Model(&ProductCategoryDocumentGroup{}).Where("prod_ctgry_doc_group_id = ?", pid).Take(&p).Error
	if err != nil {
		return &ProductCategoryDocumentGroup{}, err
	}
	return p, nil
}

func (p *ProductCategoryDocumentGroup) FindProductCategoryDocumentGroupByID(db *gorm.DB, pid uint64) (*ProductCategoryDocumentGroupData, error) {
	var err error
	address := ProductCategoryDocumentGroupData{}
	err = db.Debug().Model(&ProductCategoryDocumentGroupData{}).Limit(100).
		Select(`m_product_category_document_group.prod_ctgry_doc_group_id,
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
		Joins("JOIN m_product_document on m_product_category_document_group.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_sub_category on m_product_category_document_group.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Where("prod_ctgry_doc_group_id = ?", pid).
		Find(&address).Error
	if err != nil {
		return &ProductCategoryDocumentGroupData{}, err
	}
	return &address, nil
}

func (p *ProductCategoryDocumentGroup) FindProductCategoryDocumentGroupStatusByID(db *gorm.DB, pid uint64, status uint64) (*ProductCategoryDocumentGroupData, error) {
	var err error
	address := ProductCategoryDocumentGroupData{}
	err = db.Debug().Model(&ProductCategoryDocumentGroupData{}).Limit(100).
		Select(`m_product_category_document_group.prod_ctgry_doc_group_id,
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
		Joins("JOIN m_product_document on m_product_category_document_group.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_sub_category on m_product_category_document_group.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Where("prod_ctgry_doc_group_id = ? AND prod_ctgry_doc_group_status = ?", pid, status).
		Find(&address).Error
	if err != nil {
		return &ProductCategoryDocumentGroupData{}, err
	}
	return &address, nil
}

func (p *ProductCategoryDocumentGroup) UpdateProductCategoryDocumentGroup(db *gorm.DB, pid uint64) (*ProductCategoryDocumentGroup, error) {
	var err error
	db = db.Debug().Model(&ProductCategoryDocumentGroup{}).Where("prod_ctgry_doc_group_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_id":                     p.ProductDocumentID,
			"prod_sub_ctgry_id":               p.ProductSubCategoryID,
			"prod_ctgry_doc_group_status":     p.ProductCategoryDocumentGroupStatus,
			"prod_ctgry_doc_group_updated_by": p.ProductCategoryDocumentGroupUpdatedBy,
			"prod_ctgry_doc_group_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ProductCategoryDocumentGroup{}).Where("prod_ctgry_doc_group_id = ?", pid).Error
	if err != nil {
		return &ProductCategoryDocumentGroup{}, err
	}
	return p, nil
}

func (p *ProductCategoryDocumentGroup) UpdateProductCategoryDocumentGroupStatus(db *gorm.DB, pid uint64) (*ProductCategoryDocumentGroup, error) {
	var err error
	db = db.Debug().Model(&ProductCategoryDocumentGroup{}).Where("prod_ctgry_doc_group_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_ctgry_doc_group_status":     p.ProductCategoryDocumentGroupStatus,
			"prod_ctgry_doc_group_updated_by": p.ProductCategoryDocumentGroupUpdatedBy,
			"prod_ctgry_doc_group_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ProductCategoryDocumentGroup{}).Where("prod_ctgry_doc_group_id = ?", pid).Error
	if err != nil {
		return &ProductCategoryDocumentGroup{}, err
	}
	return p, nil
}

func (p *ProductCategoryDocumentGroup) UpdateProductCategoryDocumentGroupStatusOnly(db *gorm.DB, pid uint64) (*ProductCategoryDocumentGroup, error) {
	var err error
	db = db.Debug().Model(&ProductCategoryDocumentGroup{}).Where("prod_ctgry_doc_group_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_ctgry_doc_group_status": p.ProductCategoryDocumentGroupStatus,
		},
	)
	err = db.Debug().Model(&ProductCategoryDocumentGroup{}).Where("prod_ctgry_doc_group_id = ?", pid).Error
	if err != nil {
		return &ProductCategoryDocumentGroup{}, err
	}
	return p, nil
}

func (p *ProductCategoryDocumentGroup) UpdateProductCategoryDocumentGroupStatusDelete(db *gorm.DB, pid uint64) (*ProductCategoryDocumentGroup, error) {
	var err error
	db = db.Debug().Model(&ProductCategoryDocumentGroup{}).Where("prod_ctgry_doc_group_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_ctgry_doc_group_status":     3,
			"prod_ctgry_doc_group_deleted_by": p.ProductCategoryDocumentGroupDeletedBy,
			"prod_ctgry_doc_group_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ProductCategoryDocumentGroup{}).Where("prod_ctgry_doc_group_id = ?", pid).Error
	if err != nil {
		return &ProductCategoryDocumentGroup{}, err
	}
	return p, nil
}

func (p *ProductCategoryDocumentGroup) DeleteProductCategoryDocumentGroup(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&ProductCategoryDocumentGroup{}).Where("prod_ctgry_doc_group_id = ?", pid).Take(&ProductCategoryDocumentGroup{}).Delete(&ProductCategoryDocumentGroup{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("ProductCategoryDocumentGroup not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================
