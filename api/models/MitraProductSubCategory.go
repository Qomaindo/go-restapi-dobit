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

type MitraProductSubCategory struct {
	MitraProductSubCategoryID        *uint64    `gorm:"column:mitra_prod_sub_ctgry_id;primary_key;" json:"mitra_prod_sub_ctgry_id"`
	MitraProductID                   uint64     `gorm:"column:mitra_prod_id" json:"mitra_prod_id"`
	ProductSubCategoryID             uint64     `gorm:"column:prod_sub_ctgry_id" json:"prod_sub_ctgry_id"`
	MitraProductSubCategoryStatus    uint64     `gorm:"column:mitra_prod_sub_ctgry_status;size:2" json:"mitra_prod_sub_ctgry_status"`
	MitraProductSubCategoryCreatedBy string     `gorm:"column:mitra_prod_sub_ctgry_created_by;size:125" json:"mitra_prod_sub_ctgry_created_by"`
	MitraProductSubCategoryCreatedAt time.Time  `gorm:"column:mitra_prod_sub_ctgry_created_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_sub_ctgry_created_at"`
	MitraProductSubCategoryUpdatedBy string     `gorm:"column:mitra_prod_sub_ctgry_updated_by;size:125" json:"mitra_prod_sub_ctgry_updated_by"`
	MitraProductSubCategoryUpdatedAt *time.Time `gorm:"column:mitra_prod_sub_ctgry_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_sub_ctgry_created_at"`
	MitraProductSubCategoryDeletedBy string     `gorm:"column:mitra_prod_sub_ctgry_deleted_by;size:125" json:"mitra_prod_sub_ctgry_deleted_by"`
	MitraProductSubCategoryDeletedAt *time.Time `gorm:"column:mitra_prod_sub_ctgry_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_sub_ctgry_deleted_at"`
}

type MitraProductSubCategoryData struct {
	MitraProductSubCategoryID        uint64     `gorm:"column:mitra_prod_sub_ctgry_id" json:"mitra_prod_sub_ctgry_id"`
	MitraProductID                   uint64     `gorm:"column:mitra_prod_id" json:"mitra_prod_id"`
	MitraProductName                 string     `gorm:"column:mitra_prod_name" json:"mitra_prod_name"`
	ProductSubCategoryID             uint64     `gorm:"column:prod_sub_ctgry_id" json:"prod_sub_ctgry_id"`
	ProductSubCategoryName           string     `gorm:"column:prod_sub_ctgry_name" json:"prod_sub_ctgry_name"`
	MitraProductSubCategoryStatus    uint64     `gorm:"column:mitra_prod_sub_ctgry_status" json:"mitra_prod_sub_ctgry_status"`
	MitraProductSubCategoryCreatedBy string     `gorm:"column:mitra_prod_sub_ctgry_created_by" json:"mitra_prod_sub_ctgry_created_by"`
	MitraProductSubCategoryCreatedAt time.Time  `gorm:"column:mitra_prod_sub_ctgry_created_at" json:"mitra_prod_sub_ctgry_created_at"`
	MitraProductSubCategoryUpdatedBy string     `gorm:"column:mitra_prod_sub_ctgry_updated_by" json:"mitra_prod_sub_ctgry_updated_by"`
	MitraProductSubCategoryUpdatedAt *time.Time `gorm:"column:mitra_prod_sub_ctgry_updated_at" json:"mitra_prod_sub_ctgry_updated_at"`
	MitraProductSubCategoryDeletedBy string     `gorm:"column:mitra_prod_sub_ctgry_deleted_by" json:"mitra_prod_sub_ctgry_deleted_by"`
	MitraProductSubCategoryDeletedAt *time.Time `gorm:"column:mitra_prod_sub_ctgry_deleted_at" json:"mitra_prod_sub_ctgry_deleted_at"`
}

type ResponseMitraProductSubCategorys struct {
	Status  int                            `json:"status"`
	Message string                         `json:"message"`
	Data    *[]MitraProductSubCategoryData `json:"data"`
}

type ResponseMitraProductSubCategory struct {
	Status  int                          `json:"status"`
	Message string                       `json:"message"`
	Data    *MitraProductSubCategoryData `json:"data"`
}

type ResponseMitraProductSubCategoryCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraProductSubCategory) TableName() string {
	return "m_mitra_product_sub_category"
}

func (MitraProductSubCategoryData) TableName() string {
	return "m_mitra_product_sub_category"
}

func (p *MitraProductSubCategory) Prepare() {
	p.MitraProductSubCategoryID = nil
	p.MitraProductID = p.MitraProductID
	p.ProductSubCategoryID = p.ProductSubCategoryID
	p.MitraProductSubCategoryStatus = p.MitraProductSubCategoryStatus
	p.MitraProductSubCategoryCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraProductSubCategoryCreatedBy))
	p.MitraProductSubCategoryCreatedAt = time.Now()
	p.MitraProductSubCategoryUpdatedBy = html.EscapeString(strings.TrimSpace(p.MitraProductSubCategoryUpdatedBy))
	p.MitraProductSubCategoryUpdatedAt = p.MitraProductSubCategoryUpdatedAt
	p.MitraProductSubCategoryDeletedBy = html.EscapeString(strings.TrimSpace(p.MitraProductSubCategoryDeletedBy))
	p.MitraProductSubCategoryDeletedAt = p.MitraProductSubCategoryDeletedAt
}

func (p *MitraProductSubCategory) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.MitraProductID == 0 {
			return errors.New("Required Country")
		}
		if p.ProductSubCategoryID == 0 {
			return errors.New("Required ProductSubCategory")
		}
		return nil
	}
}

func (p *MitraProductSubCategory) SaveMitraProductSubCategory(db *gorm.DB) (*MitraProductSubCategory, error) {
	var err error
	err = db.Debug().Model(&MitraProductSubCategory{}).Create(&p).Error
	if err != nil {
		return &MitraProductSubCategory{}, err
	}
	return p, nil
}

func (p *MitraProductSubCategory) FindAllMitraProductSubCategory(db *gorm.DB) (*[]MitraProductSubCategoryData, error) {
	var err error
	mitraproductsubcategory := []MitraProductSubCategoryData{}
	err = db.Debug().Model(&MitraProductSubCategoryData{}).Limit(100).
		Select(`m_mitra_product_sub_category.mitra_prod_sub_ctgry_id,
				m_mitra_product.mitra_prod_id,
				m_mitra_product.mitra_prod_name,
				m_product_sub_category.prod_sub_ctgry_id,
				m_product_sub_category.prod_sub_ctgry_name,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_status,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_created_by,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_created_at,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_updated_by,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_updated_at,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_deleted_by,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_deleted_at`).
		Joins("JOIN m_mitra_product on m_mitra_product_sub_category.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_sub_category on m_mitra_product_sub_category.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Find(&mitraproductsubcategory).Error
	if err != nil {
		return &[]MitraProductSubCategoryData{}, err
	}
	return &mitraproductsubcategory, nil
}

func (p *MitraProductSubCategory) FindAllMitraProductSubCategoryStatus(db *gorm.DB, status uint64) (*[]MitraProductSubCategoryData, error) {
	var err error
	mitraproductsubcategory := []MitraProductSubCategoryData{}
	err = db.Debug().Model(&MitraProductSubCategoryData{}).Limit(100).
		Select(`m_mitra_product_sub_category.mitra_prod_sub_ctgry_id,
				m_mitra_product.mitra_prod_id,
				m_mitra_product.mitra_prod_name,
				m_product_sub_category.prod_sub_ctgry_id,
				m_product_sub_category.prod_sub_ctgry_name,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_status,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_created_by,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_created_at,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_updated_by,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_updated_at,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_deleted_by,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_deleted_at`).
		Joins("JOIN m_mitra_product on m_mitra_product_sub_category.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_sub_category on m_mitra_product_sub_category.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Where("mitra_prod_sub_ctgry_status = ?", status).
		Find(&mitraproductsubcategory).Error
	if err != nil {
		return &[]MitraProductSubCategoryData{}, err
	}
	return &mitraproductsubcategory, nil
}

func (p *MitraProductSubCategory) FindMitraProductSubCategoryDataByID(db *gorm.DB, pid uint64) (*MitraProductSubCategory, error) {
	var err error
	err = db.Debug().Model(&MitraProductSubCategory{}).Where("mitra_prod_sub_ctgry_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraProductSubCategory{}, err
	}
	return p, nil
}

func (p *MitraProductSubCategory) FindMitraProductSubCategoryByID(db *gorm.DB, pid uint64) (*MitraProductSubCategoryData, error) {
	var err error
	mitraproductsubcategory := MitraProductSubCategoryData{}
	err = db.Debug().Model(&MitraProductSubCategoryData{}).Limit(100).
		Select(`m_mitra_product_sub_category.mitra_prod_sub_ctgry_id,
				m_mitra_product.mitra_prod_id,
				m_mitra_product.mitra_prod_name,
				m_product_sub_category.prod_sub_ctgry_id,
				m_product_sub_category.prod_sub_ctgry_name,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_status,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_created_by,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_created_at,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_updated_by,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_updated_at,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_deleted_by,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_deleted_at`).
		Joins("JOIN m_mitra_product on m_mitra_product_sub_category.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_sub_category on m_mitra_product_sub_category.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Where("mitra_prod_sub_ctgry_id = ?", pid).
		Find(&mitraproductsubcategory).Error
	if err != nil {
		return &MitraProductSubCategoryData{}, err
	}
	return &mitraproductsubcategory, nil
}

func (p *MitraProductSubCategory) FindMitraProductSubCategoryStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraProductSubCategoryData, error) {
	var err error
	mitraproductsubcategory := MitraProductSubCategoryData{}
	err = db.Debug().Model(&MitraProductSubCategoryData{}).Limit(100).
		Select(`m_mitra_product_sub_category.mitra_prod_sub_ctgry_id,
				m_mitra_product.mitra_prod_id,
				m_mitra_product.mitra_prod_name,
				m_product_sub_category.prod_sub_ctgry_id,
				m_product_sub_category.prod_sub_ctgry_name,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_status,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_created_by,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_created_at,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_updated_by,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_updated_at,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_deleted_by,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_deleted_at`).
		Joins("JOIN m_mitra_product on m_mitra_product_sub_category.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_sub_category on m_mitra_product_sub_category.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Where("mitra_prod_sub_ctgry_id = ? AND mitra_prod_sub_ctgry_status = ?", pid, status).
		Find(&mitraproductsubcategory).Error
	if err != nil {
		return &MitraProductSubCategoryData{}, err
	}
	return &mitraproductsubcategory, nil
}

func (p *MitraProductSubCategory) UpdateMitraProductSubCategory(db *gorm.DB, pid uint64) (*MitraProductSubCategory, error) {
	var err error
	db = db.Debug().Model(&MitraProductSubCategory{}).Where("mitra_prod_sub_ctgry_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_prod_id":                   p.MitraProductID,
			"prod_sub_ctgry_id":               p.ProductSubCategoryID,
			"mitra_prod_sub_ctgry_status":     p.MitraProductSubCategoryStatus,
			"mitra_prod_sub_ctgry_updated_by": p.MitraProductSubCategoryUpdatedBy,
			"mitra_prod_sub_ctgry_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraProductSubCategory{}).Where("mitra_prod_sub_ctgry_id = ?", pid).Error
	if err != nil {
		return &MitraProductSubCategory{}, err
	}
	return p, nil
}

func (p *MitraProductSubCategory) UpdateMitraProductSubCategoryStatus(db *gorm.DB, pid uint64) (*MitraProductSubCategory, error) {
	var err error
	db = db.Debug().Model(&MitraProductSubCategory{}).Where("mitra_prod_sub_ctgry_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_prod_sub_ctgry_status":     p.MitraProductSubCategoryStatus,
			"mitra_prod_sub_ctgry_updated_by": p.MitraProductSubCategoryUpdatedBy,
			"mitra_prod_sub_ctgry_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraProductSubCategory{}).Where("mitra_prod_sub_ctgry_id = ?", pid).Error
	if err != nil {
		return &MitraProductSubCategory{}, err
	}
	return p, nil
}

func (p *MitraProductSubCategory) UpdateMitraProductSubCategoryStatusOnly(db *gorm.DB, pid uint64) (*MitraProductSubCategory, error) {
	var err error
	db = db.Debug().Model(&MitraProductSubCategory{}).Where("mitra_prod_sub_ctgry_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_prod_sub_ctgry_status": p.MitraProductSubCategoryStatus,
		},
	)
	err = db.Debug().Model(&MitraProductSubCategory{}).Where("mitra_prod_sub_ctgry_id = ?", pid).Error
	if err != nil {
		return &MitraProductSubCategory{}, err
	}
	return p, nil
}

func (p *MitraProductSubCategory) UpdateMitraProductSubCategoryStatusDelete(db *gorm.DB, pid uint64) (*MitraProductSubCategory, error) {
	var err error
	db = db.Debug().Model(&MitraProductSubCategory{}).Where("mitra_prod_sub_ctgry_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_prod_sub_ctgry_status":     3,
			"mitra_prod_sub_ctgry_deleted_by": p.MitraProductSubCategoryDeletedBy,
			"mitra_prod_sub_ctgry_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraProductSubCategory{}).Where("mitra_prod_sub_ctgry_id = ?", pid).Error
	if err != nil {
		return &MitraProductSubCategory{}, err
	}
	return p, nil
}

func (p *MitraProductSubCategory) DeleteMitraProductSubCategory(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraProductSubCategory{}).Where("mitra_prod_sub_ctgry_id = ?", pid).Take(&MitraProductSubCategory{}).Delete(&MitraProductSubCategory{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraProductSubCategory not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

type MitraProductPredefinedCoverageProductSubCategoryData struct {
	MitraProductSubCategoryID uint64 `gorm:"column:mitra_prod_sub_ctgry_id" json:"mitra_prod_sub_ctgry_id"`
	MitraProductID            uint64 `gorm:"column:mitra_prod_id" json:"mitra_prod_id"`
	ProductSubCategoryID      uint64 `gorm:"column:prod_sub_ctgry_id" json:"prod_sub_ctgry_id"`
	ProductSubCategoryName    string `gorm:"column:prod_sub_ctgry_name" json:"prod_sub_ctgry_name"`
}

type ResponsePredefinedMitraProductSubCategorys struct {
	Status  int                                                     `json:"status"`
	Message string                                                  `json:"message"`
	Data    *[]MitraProductPredefinedCoverageProductSubCategoryData `json:"data"`
}

func (MitraProductPredefinedCoverageProductSubCategoryData) TableName() string {
	return "m_mitra_product_sub_category"
}

func (p *MitraProductSubCategory) FindMitraProductSubCategoryByMitraProductID(db *gorm.DB, pid uint64) ([]MitraProductSubCategoryData, error) {
	var err error
	mitrabranchcoverageprovince := []MitraProductSubCategoryData{}
	err = db.Debug().Model(&MitraProductSubCategoryData{}).Limit(100).
		Select(`m_mitra_product_sub_category.mitra_prod_sub_ctgry_id,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_id,
				m_mitra_product.mitra_prod_id,
				m_mitra_product.mitra_prod_name,
				m_product_sub_category.prod_sub_ctgry_id,
				m_product_sub_category.prod_sub_ctgry_name,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_status,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_created_by,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_created_at`).
		Joins("JOIN m_mitra_product on m_mitra_product_sub_category.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_sub_category on m_mitra_product_sub_category.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Where("m_mitra_product_sub_category.mitra_prod_id = ?", pid).
		Find(&mitrabranchcoverageprovince).Error
	if err != nil {
		return []MitraProductSubCategoryData{}, err
	}
	return mitrabranchcoverageprovince, nil
}

func (p *MitraProductSubCategory) FindMitraProductSubCategoryStatusByMitraProductID(db *gorm.DB, mitra uint64, status uint64) ([]MitraProductSubCategoryData, error) {
	var err error
	mitrabranchcoverageprovince := []MitraProductSubCategoryData{}
	err = db.Debug().Model(&MitraProductSubCategoryData{}).Limit(100).
		Select(`m_mitra_product_sub_category.mitra_prod_sub_ctgry_id,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_id,
				m_mitra_product.mitra_prod_id,
				m_mitra_product.mitra_prod_name,
				m_product_sub_category.prod_sub_ctgry_id,
				m_product_sub_category.prod_sub_ctgry_name,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_status,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_created_by,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_created_at`).
		Joins("JOIN m_mitra_product on m_mitra_product_sub_category.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_sub_category on m_mitra_product_sub_category.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Where("m_mitra_product_sub_category.mitra_prod_id = ? AND m_mitra_product_sub_category.mitra_prod_sub_ctgry_status = ?", mitra, status).
		Find(&mitrabranchcoverageprovince).Error
	if err != nil {
		return []MitraProductSubCategoryData{}, err
	}
	return mitrabranchcoverageprovince, nil
}

func (p *MitraProductSubCategory) FindMitraProductPredefinedCoverageProductSubCategoryByMitraProductID(db *gorm.DB, mitra uint64, status uint64) ([]MitraProductPredefinedCoverageProductSubCategoryData, error) {
	var err error
	mitrabranchcoverageprovince := []MitraProductPredefinedCoverageProductSubCategoryData{}
	err = db.Debug().Model(&MitraProductPredefinedCoverageProductSubCategoryData{}).Limit(100).
		Select(`m_mitra_product_sub_category.mitra_prod_sub_ctgry_id,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_id,
				m_mitra_product.mitra_prod_id,
				m_mitra_product.mitra_prod_name,
				m_product_sub_category.prod_sub_ctgry_id,
				m_product_sub_category.prod_sub_ctgry_name,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_status,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_created_by,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_created_at`).
		Joins("JOIN m_mitra_product on m_mitra_product_sub_category.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_sub_category on m_mitra_product_sub_category.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Where("m_mitra_product_sub_category.mitra_prod_id = ? AND m_mitra_product_sub_category.mitra_prod_sub_ctgry_status = ?", mitra, status).
		Find(&mitrabranchcoverageprovince).Error
	if err != nil {
		return []MitraProductPredefinedCoverageProductSubCategoryData{}, err
	}
	return mitrabranchcoverageprovince, nil
}

func (p *MitraProductSubCategory) FindMitraProductPredefinedCoverageProductSubCategoryByMitraProductIDByProvince(db *gorm.DB, mitra uint64, province uint64, status uint64) ([]MitraProductPredefinedCoverageProductSubCategoryData, error) {
	var err error
	mitrabranchcoverageprovince := []MitraProductPredefinedCoverageProductSubCategoryData{}
	err = db.Debug().Model(&MitraProductPredefinedCoverageProductSubCategoryData{}).Limit(100).
		Select(`m_mitra_product_sub_category.mitra_prod_sub_ctgry_id,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_id,
				m_mitra_product.mitra_prod_id,
				m_mitra_product.mitra_prod_name,
				m_product_sub_category.prod_sub_ctgry_id,
				m_product_sub_category.prod_sub_ctgry_name,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_status,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_created_by,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_created_at`).
		Joins("JOIN m_mitra_product on m_mitra_product_sub_category.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_sub_category on m_mitra_product_sub_category.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Where("m_mitra_product_sub_category.mitra_prod_id = ? AND m_product_sub_category.province_id = ? AND m_mitra_product_sub_category.mitra_prod_sub_ctgry_status = ?", mitra, province, status).
		Find(&mitrabranchcoverageprovince).Error
	if err != nil {
		return []MitraProductPredefinedCoverageProductSubCategoryData{}, err
	}
	return mitrabranchcoverageprovince, nil
}
