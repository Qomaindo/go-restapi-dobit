package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type ProductDocumentMitra struct {
	ProductDocumentMitraID   uint64 `gorm:"column:prod_doc_mitra_id;primary_key;" json:"prod_doc_mitra_id"`
	ProductDocumentMitraCode string `gorm:"column:prod_doc_mitra_code;size:25" json:"prod_doc_mitra_code"`
	ProductDocumentMitraName string `gorm:"column:prod_doc_mitra_name;size:255" json:"prod_doc_mitra_name"`
	// ProductDocumentMitraDesc      string     `gorm:"column:prod_doc_mitra_desc" json:"prod_doc_mitra_desc"`
	ProductDocumentMitraStatus    uint64     `gorm:"column:prod_doc_mitra_status;size:25" json:"prod_doc_mitra_status"`
	ProductDocumentMitraCreatedBy string     `gorm:"column:prod_doc_mitra_created_by;size:125" json:"prod_doc_mitra_created_by"`
	ProductDocumentMitraUpdatedBy string     `gorm:"column:prod_doc_mitra_updated_by;size:125" json:"prod_doc_mitra_updated_by"`
	ProductDocumentMitraDeletedBy string     `gorm:"column:prod_doc_mitra_deleted_by;size:125" json:"prod_doc_mitra_deleted_by"`
	ProductDocumentMitraCreatedAt time.Time  `gorm:"column:prod_doc_mitra_created_at;default:CURRENT_TIMESTAMP" json:"prod_doc_mitra_created_at"`
	ProductDocumentMitraUpdatedAt *time.Time `gorm:"column:prod_doc_mitra_updated_at;default:CURRENT_TIMESTAMP" json:"prod_doc_mitra_updated_at"`
	ProductDocumentMitraDeletedAt *time.Time `gorm:"column:prod_doc_mitra_deleted_at;default:CURRENT_TIMESTAMP" json:"prod_doc_mitra_deleted_at"`
}

type ResponseProductDocumentMitras struct {
	Status  int                     `json:"status"`
	Message string                  `json:"message"`
	Data    *[]ProductDocumentMitra `json:"data"`
}

type ResponseProductDocumentMitra struct {
	Status  int                   `json:"status"`
	Message string                `json:"message"`
	Data    *ProductDocumentMitra `json:"data"`
}

type ResponseProductDocumentMitraIU struct {
	Status  int                   `json:"status"`
	Message string                `json:"message"`
	Data    *ProductDocumentMitra `json:"data"`
}

type ResponseProductDocumentMitraDel struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (ProductDocumentMitra) TableName() string {
	return "m_product_document_mitra"
}

func (p *ProductDocumentMitra) Prepare() {
	p.ProductDocumentMitraCode = html.EscapeString(strings.TrimSpace(p.ProductDocumentMitraCode))
	p.ProductDocumentMitraName = html.EscapeString(strings.TrimSpace(p.ProductDocumentMitraName))
	p.ProductDocumentMitraStatus = p.ProductDocumentMitraStatus
	p.ProductDocumentMitraCreatedBy = html.EscapeString(strings.TrimSpace(p.ProductDocumentMitraCreatedBy))
	p.ProductDocumentMitraUpdatedBy = html.EscapeString(strings.TrimSpace(p.ProductDocumentMitraUpdatedBy))
	p.ProductDocumentMitraDeletedBy = html.EscapeString(strings.TrimSpace(p.ProductDocumentMitraDeletedBy))
	p.ProductDocumentMitraCreatedAt = time.Now()
	p.ProductDocumentMitraUpdatedAt = p.ProductDocumentMitraUpdatedAt
	p.ProductDocumentMitraDeletedAt = p.ProductDocumentMitraDeletedAt
}

func (p *ProductDocumentMitra) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.ProductDocumentMitraCode == "" {
			return errors.New("Required ProductDocumentMitra Code")
		}
		if p.ProductDocumentMitraName == "" {
			return errors.New("Required ProductDocumentMitra Name")
		}
		return nil
	}
}

func (p *ProductDocumentMitra) SaveProductDocumentMitra(db *gorm.DB) (*ProductDocumentMitra, error) {
	var err error
	err = db.Debug().Model(&ProductDocumentMitra{}).Create(&p).Error
	if err != nil {
		return &ProductDocumentMitra{}, err
	}
	return p, nil
}

func (p *ProductDocumentMitra) FindAllProductDocumentMitra(db *gorm.DB) (*[]ProductDocumentMitra, error) {
	var err error
	productdocumentmitra := []ProductDocumentMitra{}
	err = db.Debug().Model(&ProductDocumentMitra{}).Limit(100).Find(&productdocumentmitra).Error
	if err != nil {
		return &[]ProductDocumentMitra{}, err
	}
	return &productdocumentmitra, nil
}

func (p *ProductDocumentMitra) FindAllProductDocumentMitraStatus(db *gorm.DB, status uint64) (*[]ProductDocumentMitra, error) {
	var err error
	productdocumentmitra := []ProductDocumentMitra{}
	err = db.Debug().Model(&ProductDocumentMitra{}).Where("prod_doc_mitra_status = ?", status).Limit(100).Find(&productdocumentmitra).Error
	if err != nil {
		return &[]ProductDocumentMitra{}, err
	}
	return &productdocumentmitra, nil
}

func (p *ProductDocumentMitra) FindProductDocumentMitraByID(db *gorm.DB, pid uint64) (*ProductDocumentMitra, error) {
	var err error
	err = db.Debug().Model(&ProductDocumentMitra{}).Where("prod_doc_mitra_id = ?", pid).Take(&p).Error
	if err != nil {
		return &ProductDocumentMitra{}, err
	}
	return p, nil
}

func (p *ProductDocumentMitra) FindProductDocumentMitraStatusByID(db *gorm.DB, pid uint64, status uint64) (*ProductDocumentMitra, error) {
	var err error
	err = db.Debug().Model(&ProductDocumentMitra{}).Where("prod_doc_mitra_id = ? AND prod_doc_mitra_status = ?", pid, status).Take(&p).Error
	if err != nil {
		return &ProductDocumentMitra{}, err
	}
	return p, nil
}

func (p *ProductDocumentMitra) UpdateProductDocumentMitra(db *gorm.DB, pid uint64) (*ProductDocumentMitra, error) {
	var err error
	db = db.Debug().Model(&ProductDocumentMitra{}).Where("prod_doc_mitra_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_mitra_code":       p.ProductDocumentMitraCode,
			"prod_doc_mitra_name":       p.ProductDocumentMitraName,
			"prod_doc_mitra_status":     p.ProductDocumentMitraStatus,
			"prod_doc_mitra_updated_by": p.ProductDocumentMitraUpdatedBy,
			"prod_doc_mitra_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ProductDocumentMitra{}).Where("prod_doc_mitra_id = ?", pid).Error
	if err != nil {
		return &ProductDocumentMitra{}, err
	}
	return p, nil
}

func (p *ProductDocumentMitra) UpdateProductDocumentMitraStatus(db *gorm.DB, pid uint64) (*ProductDocumentMitra, error) {
	var err error
	db = db.Debug().Model(&ProductDocumentMitra{}).Where("prod_doc_mitra_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_mitra_status":     p.ProductDocumentMitraStatus,
			"prod_doc_mitra_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ProductDocumentMitra{}).Where("prod_doc_mitra_id = ?", pid).Error
	if err != nil {
		return &ProductDocumentMitra{}, err
	}
	return p, nil
}

func (p *ProductDocumentMitra) UpdateProductDocumentMitraStatusDelete(db *gorm.DB, pid uint64) (*ProductDocumentMitra, error) {
	var err error
	db = db.Debug().Model(&ProductDocumentMitra{}).Where("prod_doc_mitra_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_mitra_status":     3,
			"prod_doc_mitra_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ProductDocumentMitra{}).Where("prod_doc_mitra_id = ?", pid).Error
	if err != nil {
		return &ProductDocumentMitra{}, err
	}
	return p, nil
}

func (p *ProductDocumentMitra) DeleteProductDocumentMitra(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&ProductDocumentMitra{}).Where("prod_doc_mitra_id = ?", pid).Take(&ProductDocumentMitra{}).Delete(&ProductDocumentMitra{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("ProductDocumentMitra not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
