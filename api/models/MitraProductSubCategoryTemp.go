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

type MitraProductSubCategoryTemp struct {
	MitraProductSubCategoryTempID           *uint64   `gorm:"column:mitra_prod_sub_ctgry_temp_id;primary_key;" json:"mitra_prod_sub_ctgry_temp_id"`
	MitraProductSubCategoryID               uint64    `gorm:"column:mitra_prod_sub_ctgry_id" json:"mitra_prod_sub_ctgry_id"`
	MitraProductTempID                      uint64    `gorm:"column:mitra_prod_temp_id" json:"mitra_prod_temp_id"`
	ProductSubCategoryTempID                uint64    `gorm:"column:prod_sub_ctgry_temp_id" json:"prod_sub_ctgry_temp_id"`
	MitraProductSubCategoryTempReason       string    `gorm:"column:mitra_prod_sub_ctgry_temp_reason" json:"mitra_prod_sub_ctgry_temp_reason"`
	MitraProductSubCategoryTempStatus       uint64    `gorm:"column:mitra_prod_sub_ctgry_temp_status;size:2" json:"mitra_prod_sub_ctgry_temp_status"`
	MitraProductSubCategoryTempActionBefore uint64    `gorm:"column:mitra_prod_sub_ctgry_temp_action_before;size:2" json:"mitra_prod_sub_ctgry_temp_action_before"`
	MitraProductSubCategoryTempCreatedBy    string    `gorm:"column:mitra_prod_sub_ctgry_temp_created_by;size:125" json:"mitra_prod_sub_ctgry_temp_created_by"`
	MitraProductSubCategoryTempCreatedAt    time.Time `gorm:"column:mitra_prod_sub_ctgry_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_sub_ctgry_temp_created_at"`
}

type MitraProductSubCategoryTempPend struct {
	MitraProductSubCategoryTempID           *uint64   `gorm:"column:mitra_prod_sub_ctgry_temp_id;primary_key;" json:"mitra_prod_sub_ctgry_temp_id"`
	MitraProductSubCategoryID               *uint64   `gorm:"column:mitra_prod_sub_ctgry_id" json:"mitra_prod_sub_ctgry_id"`
	MitraProductTempID                      uint64    `gorm:"column:mitra_prod_temp_id" json:"mitra_prod_temp_id"`
	ProductSubCategoryTempID                uint64    `gorm:"column:prod_sub_ctgry_temp_id" json:"prod_sub_ctgry_temp_id"`
	MitraProductSubCategoryTempReason       string    `gorm:"column:mitra_prod_sub_ctgry_temp_reason" json:"mitra_prod_sub_ctgry_temp_reason"`
	MitraProductSubCategoryTempStatus       uint64    `gorm:"column:mitra_prod_sub_ctgry_temp_status;size:2" json:"mitra_prod_sub_ctgry_temp_status"`
	MitraProductSubCategoryTempActionBefore uint64    `gorm:"column:mitra_prod_sub_ctgry_temp_action_before;size:2" json:"mitra_prod_sub_ctgry_temp_action_before"`
	MitraProductSubCategoryTempCreatedBy    string    `gorm:"column:mitra_prod_sub_ctgry_temp_created_by;size:125" json:"mitra_prod_sub_ctgry_temp_created_by"`
	MitraProductSubCategoryTempCreatedAt    time.Time `gorm:"column:mitra_prod_sub_ctgry_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_sub_ctgry_temp_created_at"`
}

type MitraProductSubCategoryTempPendData struct {
	MitraProductSubCategoryTempID           uint64    `gorm:"column:mitra_prod_sub_ctgry_temp_id;primary_key;" json:"mitra_prod_sub_ctgry_temp_id"`
	MitraProductSubCategoryID               uint64    `gorm:"column:mitra_prod_sub_ctgry_id" json:"mitra_prod_sub_ctgry_id"`
	MitraProductTempID                      uint64    `gorm:"column:mitra_prod_temp_id" json:"mitra_prod_temp_id"`
	ProductSubCategoryTempID                uint64    `gorm:"column:prod_sub_ctgry_temp_id" json:"prod_sub_ctgry_temp_id"`
	ProductSubCategoryTempName              string    `gorm:"column:prod_sub_ctgry_temp_name" json:"prod_sub_ctgry_temp_name"`
	MitraProductSubCategoryTempReason       string    `gorm:"column:mitra_prod_sub_ctgry_temp_reason" json:"mitra_prod_sub_ctgry_temp_reason"`
	MitraProductSubCategoryTempStatus       uint64    `gorm:"column:mitra_prod_sub_ctgry_temp_status;size:2" json:"mitra_prod_sub_ctgry_temp_status"`
	MitraProductSubCategoryTempActionBefore uint64    `gorm:"column:mitra_prod_sub_ctgry_temp_action_before;size:2" json:"mitra_prod_sub_ctgry_temp_action_before"`
	MitraProductSubCategoryTempCreatedBy    string    `gorm:"column:mitra_prod_sub_ctgry_temp_created_by;size:125" json:"mitra_prod_sub_ctgry_temp_created_by"`
	MitraProductSubCategoryTempCreatedAt    time.Time `gorm:"column:mitra_prod_sub_ctgry_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_sub_ctgry_temp_created_at"`
}

type MitraProductSubCategoryTempData struct {
	MitraProductSubCategoryTempID           uint64    `gorm:"column:mitra_prod_sub_ctgry_temp_id" json:"mitra_prod_sub_ctgry_temp_id"`
	MitraProductTempID                      uint64    `gorm:"column:mitra_prod_temp_id" json:"mitra_prod_temp_id"`
	MitraProductTempName                    string    `gorm:"column:mitra_prod_temp_name" json:"mitra_prod_temp_name"`
	ProductSubCategoryTempID                uint64    `gorm:"column:prod_sub_ctgry_temp_id" json:"prod_sub_ctgry_temp_id"`
	ProductSubCategoryTempName              string    `gorm:"column:prod_sub_ctgry_temp_name" json:"prod_sub_ctgry_temp_name"`
	MitraProductSubCategoryTempReason       string    `gorm:"column:mitra_prod_sub_ctgry_temp_reason" json:"mitra_prod_sub_ctgry_temp_reason"`
	MitraProductSubCategoryTempStatus       uint64    `gorm:"column:mitra_prod_sub_ctgry_temp_status;size:2" json:"mitra_prod_sub_ctgry_temp_status"`
	MitraProductSubCategoryTempActionBefore uint64    `gorm:"column:mitra_prod_sub_ctgry_temp_action_before;size:2" json:"mitra_prod_sub_ctgry_temp_action_before"`
	MitraProductSubCategoryTempCreatedBy    string    `gorm:"column:mitra_prod_sub_ctgry_temp_created_by;size:125" json:"mitra_prod_sub_ctgry_temp_created_by"`
	MitraProductSubCategoryTempCreatedAt    time.Time `gorm:"column:mitra_prod_sub_ctgry_temp_created_at" json:"mitra_prod_sub_ctgry_temp_created_at"`
	MitraProductSubCategoryID               uint64    `gorm:"column:mitra_prod_sub_ctgry_id" json:"mitra_prod_sub_ctgry_id"`
}

type ResponseMitraProductSubCategoryTemps struct {
	Status  int                                `json:"status"`
	Message string                             `json:"message"`
	Data    *[]MitraProductSubCategoryTempData `json:"data"`
}

type ResponseMitraProductSubCategoryTemp struct {
	Status  int                              `json:"status"`
	Message string                           `json:"message"`
	Data    *MitraProductSubCategoryTempData `json:"data"`
}

type ResponseMitraProductSubCategoryTempsPend struct {
	Status  int                                `json:"status"`
	Message string                             `json:"message"`
	Data    *[]MitraProductSubCategoryTempPend `json:"data"`
}

type ResponseMitraProductSubCategoryTempPend struct {
	Status  int                              `json:"status"`
	Message string                           `json:"message"`
	Data    *MitraProductSubCategoryTempPend `json:"data"`
}

type ResponseMitraProductSubCategoryTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraProductSubCategoryTemp) TableName() string {
	return "m_mitra_product_sub_category_temp"
}

func (MitraProductSubCategoryTempPend) TableName() string {
	return "m_mitra_product_sub_category_temp"
}

func (MitraProductSubCategoryTempPendData) TableName() string {
	return "m_mitra_product_sub_category_temp"
}

func (MitraProductSubCategoryTempData) TableName() string {
	return "m_mitra_product_sub_category_temp"
}

func (p *MitraProductSubCategoryTemp) Prepare() {
	p.MitraProductSubCategoryTempID = nil
	p.MitraProductSubCategoryID = p.MitraProductSubCategoryID
	p.MitraProductTempID = p.MitraProductTempID
	p.ProductSubCategoryTempID = p.ProductSubCategoryTempID
	p.MitraProductSubCategoryTempReason = p.MitraProductSubCategoryTempReason
	p.MitraProductSubCategoryTempStatus = p.MitraProductSubCategoryTempStatus
	p.MitraProductSubCategoryTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraProductSubCategoryTempCreatedBy))
	p.MitraProductSubCategoryTempCreatedAt = time.Now()
}

func (p *MitraProductSubCategoryTempPend) Prepare() {
	p.MitraProductSubCategoryTempID = nil
	p.MitraProductSubCategoryID = nil
	p.MitraProductTempID = p.MitraProductTempID
	p.ProductSubCategoryTempID = p.ProductSubCategoryTempID
	p.MitraProductSubCategoryTempReason = p.MitraProductSubCategoryTempReason
	p.MitraProductSubCategoryTempStatus = p.MitraProductSubCategoryTempStatus
	p.MitraProductSubCategoryTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraProductSubCategoryTempCreatedBy))
	p.MitraProductSubCategoryTempCreatedAt = time.Now()
}

func (p *MitraProductSubCategoryTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	case "insertupdate":
		if p.MitraProductTempID == 0 {
			return errors.New("Negara Wajib diisi")
		}
		if p.ProductSubCategoryTempID == 0 {
			return errors.New("Provinsi Wajib diisi")
		}
		if p.MitraProductSubCategoryTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.MitraProductSubCategoryTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	case "reason":
		if p.MitraProductSubCategoryTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.MitraProductSubCategoryTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	default:
		if p.MitraProductSubCategoryTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil
	}
}

func (p *MitraProductSubCategoryTemp) SaveMitraProductSubCategoryTemp(db *gorm.DB) (*MitraProductSubCategoryTemp, error) {
	var err error
	err = db.Debug().Model(&MitraProductSubCategoryTemp{}).Create(&p).Error
	if err != nil {
		return &MitraProductSubCategoryTemp{}, err
	}
	return p, nil
}

func (p *MitraProductSubCategoryTempPend) SaveMitraProductSubCategoryTempPend(db *gorm.DB) (*MitraProductSubCategoryTempPend, error) {
	var err error
	err = db.Debug().Model(&MitraProductSubCategoryTempPend{}).Create(&p).Error
	if err != nil {
		return &MitraProductSubCategoryTempPend{}, err
	}
	return p, nil
}

func (p *MitraProductSubCategoryTemp) FindAllMitraProductSubCategoryTemp(db *gorm.DB) (*[]MitraProductSubCategoryTempPend, error) {
	var err error
	mitraproductsubcategory := []MitraProductSubCategoryTempPend{}
	err = db.Debug().Model(&MitraProductSubCategoryTempPend{}).Limit(100).
		Select(`m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_id,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_id,
				m_mitra_product_sub_category_temp.mitra_prod_temp_id,
				m_mitra_product_sub_category_temp.prod_sub_ctgry_temp_id,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_reason,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_status,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_action_before,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_created_by,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_created_at`).
		Find(&mitraproductsubcategory).Error
	if err != nil {
		return &[]MitraProductSubCategoryTempPend{}, err
	}
	return &mitraproductsubcategory, nil
}

func (p *MitraProductSubCategoryTemp) FindAllMitraProductSubCategoryTempPendingActive(db *gorm.DB) (*[]MitraProductSubCategoryTempPend, error) {
	var err error
	mitraproductsubcategory := []MitraProductSubCategoryTempPend{}
	err = db.Debug().Model(&MitraProductSubCategoryTempPend{}).Limit(100).
		Select(`m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_id,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_id,
				m_mitra_product_sub_category_temp.mitra_prod_temp_id,
				m_mitra_product_sub_category_temp.prod_sub_ctgry_temp_id,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_reason,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_status,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_action_before,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_created_by,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_created_at`).
		Where("mitra_prod_sub_ctgry_temp_status = ?", 11).
		Find(&mitraproductsubcategory).Error
	if err != nil {
		return &[]MitraProductSubCategoryTempPend{}, err
	}
	return &mitraproductsubcategory, nil
}

func (p *MitraProductSubCategoryTemp) FindAllMitraProductSubCategoryTempStatus(db *gorm.DB, status uint64) (*[]MitraProductSubCategoryTempData, error) {
	var err error
	mitraproductsubcategory := []MitraProductSubCategoryTempData{}
	err = db.Debug().Model(&MitraProductSubCategoryTempData{}).Limit(100).
		Select(`m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_id,
				m_mitra_product_temp.mitra_prod_id as mitra_prod_temp_id,
				m_mitra_product_temp.mitra_prod_name as mitra_prod_temp_name,
				m_product_sub_category_temp.prod_sub_ctgry_id as prod_sub_ctgry_temp_id,
				m_product_sub_category_temp.prod_sub_ctgry_name as prod_sub_ctgry_temp_name,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_reason,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_status,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_action_before,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_created_by,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_created_at,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_id,
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
		Joins("JOIN m_mitra_product m_mitra_product_temp on m_mitra_product_sub_category_temp.mitra_prod_temp_id=m_mitra_product_temp.mitra_prod_id").
		Joins("JOIN m_product_sub_category m_product_sub_category_temp on m_mitra_product_sub_category_temp.prod_sub_ctgry_temp_id=m_product_sub_category_temp.prod_sub_ctgry_id").
		Joins("JOIN m_mitra_product_sub_category on m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_id=m_mitra_product_sub_category.mitra_prod_sub_ctgry_id").
		Joins("JOIN m_mitra_product on m_mitra_product_sub_category.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_sub_category on m_mitra_product_sub_category.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Where("mitra_prod_sub_ctgry_temp_status = ?", status).
		Find(&mitraproductsubcategory).Error
	if err != nil {
		return &[]MitraProductSubCategoryTempData{}, err
	}
	return &mitraproductsubcategory, nil
}

func (p *MitraProductSubCategoryTemp) FindMitraProductSubCategoryTempDataByID(db *gorm.DB, pid uint64) (*MitraProductSubCategoryTemp, error) {
	var err error
	err = db.Debug().Model(&MitraProductSubCategoryTemp{}).Where("mitra_prod_sub_ctgry_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraProductSubCategoryTemp{}, err
	}
	return p, nil
}

func (p *MitraProductSubCategoryTemp) FindMitraProductSubCategoryTempByIDPendingActive(db *gorm.DB, pid uint64) (*MitraProductSubCategoryTempPend, error) {
	var err error
	mitraproductsubcategory := MitraProductSubCategoryTempPend{}
	err = db.Debug().Model(&MitraProductSubCategoryTempPend{}).Limit(100).
		Select(`m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_id,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_id,
				m_mitra_product_sub_category_temp.mitra_prod_temp_id,
				m_mitra_product_sub_category_temp.prod_sub_ctgry_temp_id,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_reason,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_status,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_action_before,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_created_by,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_created_at`).
		Where("mitra_prod_sub_ctgry_temp_id = ? AND mitra_prod_sub_ctgry_temp_status = ?", pid, 11).
		Find(&mitraproductsubcategory).Error
	if err != nil {
		return &MitraProductSubCategoryTempPend{}, err
	}
	return &mitraproductsubcategory, nil
}

func (p *MitraProductSubCategoryTemp) FindMitraProductSubCategoryTempByID(db *gorm.DB, pid uint64) (*MitraProductSubCategoryTempData, error) {
	var err error
	mitraproductsubcategory := MitraProductSubCategoryTempData{}
	err = db.Debug().Model(&MitraProductSubCategoryTempData{}).Limit(100).
		Select(`m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_id,
				m_mitra_product_temp.mitra_prod_id as mitra_prod_temp_id,
				m_mitra_product_temp.mitra_prod_name as mitra_prod_temp_name,
				m_product_sub_category_temp.prod_sub_ctgry_id as prod_sub_ctgry_temp_id,
				m_product_sub_category_temp.prod_sub_ctgry_name as prod_sub_ctgry_temp_name,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_reason,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_status,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_action_before,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_created_by,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_created_at,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_id,
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
		Joins("JOIN m_mitra_product m_mitra_product_temp on m_mitra_product_sub_category_temp.mitra_prod_temp_id=m_mitra_product_temp.mitra_prod_id").
		Joins("JOIN m_product_sub_category m_product_sub_category_temp on m_mitra_product_sub_category_temp.prod_sub_ctgry_temp_id=m_product_sub_category_temp.prod_sub_ctgry_id").
		Joins("JOIN m_mitra_product_sub_category on m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_id=m_mitra_product_sub_category.mitra_prod_sub_ctgry_id").
		Joins("JOIN m_mitra_product on m_mitra_product_sub_category.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_sub_category on m_mitra_product_sub_category.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Where("mitra_prod_sub_ctgry_temp_id = ?", pid).
		Find(&mitraproductsubcategory).Error
	if err != nil {
		return &MitraProductSubCategoryTempData{}, err
	}
	return &mitraproductsubcategory, nil
}

func (p *MitraProductSubCategoryTemp) FindMitraProductSubCategoryTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraProductSubCategoryTempData, error) {
	var err error
	mitraproductsubcategory := MitraProductSubCategoryTempData{}
	err = db.Debug().Model(&MitraProductSubCategoryTempData{}).Limit(100).
		Select(`m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_id,
				m_mitra_product_temp.mitra_prod_id as mitra_prod_temp_id,
				m_mitra_product_temp.mitra_prod_name as mitra_prod_temp_name,
				m_product_sub_category_temp.prod_sub_ctgry_id as prod_sub_ctgry_temp_id,
				m_product_sub_category_temp.prod_sub_ctgry_name as prod_sub_ctgry_temp_name,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_reason,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_status,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_action_before,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_created_by,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_created_at,
				m_mitra_product_sub_category.mitra_prod_sub_ctgry_id,
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
		Joins("JOIN m_mitra_product m_mitra_product_temp on m_mitra_product_sub_category_temp.mitra_prod_temp_id=m_mitra_product_temp.mitra_prod_id").
		Joins("JOIN m_product_sub_category m_product_sub_category_temp on m_mitra_product_sub_category_temp.prod_sub_ctgry_temp_id=m_product_sub_category_temp.prod_sub_ctgry_id").
		Joins("JOIN m_mitra_product_sub_category on m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_id=m_mitra_product_sub_category.mitra_prod_sub_ctgry_id").
		Joins("JOIN m_mitra_product on m_mitra_product_sub_category.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_sub_category on m_mitra_product_sub_category.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Where("mitra_prod_sub_ctgry_temp_id = ? AND mitra_prod_sub_ctgry_temp_status = ?", pid, status).
		Find(&mitraproductsubcategory).Error
	if err != nil {
		return &MitraProductSubCategoryTempData{}, err
	}
	return &mitraproductsubcategory, nil
}

func (p *MitraProductSubCategoryTemp) UpdateMitraProductSubCategoryTemp(db *gorm.DB, pid uint64) (*MitraProductSubCategoryTemp, error) {
	var err error
	db = db.Debug().Model(&MitraProductSubCategoryTemp{}).Where("mitra_prod_sub_ctgry_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_prod_temp_id":               p.MitraProductTempID,
			"prod_sub_ctgry_temp_id":           p.ProductSubCategoryTempID,
			"mitra_prod_sub_ctgry_temp_reason": p.MitraProductSubCategoryTempReason,
		},
	)
	err = db.Debug().Model(&MitraProductSubCategoryTemp{}).Where("mitra_prod_sub_ctgry_temp_id = ?", pid).Error
	if err != nil {
		return &MitraProductSubCategoryTemp{}, err
	}
	return p, nil
}

func (p *MitraProductSubCategoryTemp) UpdateMitraProductSubCategoryTempStatus(db *gorm.DB, pid uint64) (*MitraProductSubCategoryTemp, error) {
	var err error
	db = db.Debug().Model(&MitraProductSubCategoryTemp{}).Where("mitra_prod_sub_ctgry_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_prod_sub_ctgry_temp_reason": p.MitraProductSubCategoryTempReason,
			"mitra_prod_sub_ctgry_temp_status": p.MitraProductSubCategoryTempStatus,
		},
	)
	err = db.Debug().Model(&MitraProductSubCategoryTemp{}).Where("mitra_prod_sub_ctgry_temp_id = ?", pid).Error
	if err != nil {
		return &MitraProductSubCategoryTemp{}, err
	}
	return p, nil
}

func (p *MitraProductSubCategoryTemp) DeleteMitraProductSubCategoryTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraProductSubCategoryTemp{}).Where("mitra_prod_sub_ctgry_temp_id = ?", pid).Take(&MitraProductSubCategoryTemp{}).Delete(&MitraProductSubCategoryTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraProductSubCategoryTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

func (p *MitraProductSubCategoryTemp) FindMitraProductSubCategoryTempPendByMitraProductTempID(db *gorm.DB, pid uint64) ([]MitraProductSubCategoryTempPendData, error) {
	var err error
	mitraproductsubcategory := []MitraProductSubCategoryTempPendData{}
	err = db.Debug().Model(&MitraProductSubCategoryTempPendData{}).Limit(100).
		Select(`m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_id,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_id,
				m_mitra_product_temp.mitra_prod_temp_id,
				m_mitra_product_temp.mitra_prod_temp_name,
				m_product_sub_category_temp.prod_sub_ctgry_id as prod_sub_ctgry_temp_id,
				m_product_sub_category_temp.prod_sub_ctgry_name as prod_sub_ctgry_temp_name,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_reason,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_status,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_action_before,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_created_by,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_created_at`).
		Joins("JOIN m_mitra_product_temp on m_mitra_product_sub_category_temp.mitra_prod_temp_id=m_mitra_product_temp.mitra_prod_temp_id").
		Joins("JOIN m_product_sub_category m_product_sub_category_temp on m_mitra_product_sub_category_temp.prod_sub_ctgry_temp_id=m_product_sub_category_temp.prod_sub_ctgry_id").
		Where("m_mitra_product_sub_category_temp.mitra_prod_temp_id = ?", pid).
		Find(&mitraproductsubcategory).Error
	if err != nil {
		return []MitraProductSubCategoryTempPendData{}, err
	}
	return mitraproductsubcategory, nil
}

func (p *MitraProductSubCategoryTemp) FindMitraProductSubCategoryTempPendStatusByMitraProductTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraProductSubCategoryTempPendData, error) {
	var err error
	mitraproductsubcategory := []MitraProductSubCategoryTempPendData{}
	err = db.Debug().Model(&MitraProductSubCategoryTempPendData{}).Limit(100).
		Select(`m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_id,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_id,
				m_mitra_product_temp.mitra_prod_temp_id,
				m_mitra_product_temp.mitra_prod_temp_name,
				m_product_sub_category_temp.prod_sub_ctgry_id as prod_sub_ctgry_temp_id,
				m_product_sub_category_temp.prod_sub_ctgry_name as prod_sub_ctgry_temp_name,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_reason,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_status,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_action_before,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_created_by,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_created_at`).
		Joins("JOIN m_mitra_product_temp on m_mitra_product_sub_category_temp.mitra_prod_temp_id=m_mitra_product_temp.mitra_prod_temp_id").
		Joins("JOIN m_product_sub_category m_product_sub_category_temp on m_mitra_product_sub_category_temp.prod_sub_ctgry_temp_id=m_product_sub_category_temp.prod_sub_ctgry_id").
		Where("m_mitra_product_sub_category_temp.mitra_prod_temp_id = ? AND mitra_prod_sub_ctgry_temp_status = ?", pid, status).
		Find(&mitraproductsubcategory).Error
	if err != nil {
		return []MitraProductSubCategoryTempPendData{}, err
	}
	return mitraproductsubcategory, nil
}

func (p *MitraProductSubCategoryTemp) FindMitraProductSubCategoryTempByMitraProductTempID(db *gorm.DB, pid uint64) ([]MitraProductSubCategoryTempData, error) {
	var err error
	mitraproductsubcategory := []MitraProductSubCategoryTempData{}
	err = db.Debug().Model(&MitraProductSubCategoryTempData{}).Limit(100).
		Select(`m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_id,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_id,
				m_mitra_product_temp.mitra_prod_temp_id,
				m_mitra_product_temp.mitra_prod_temp_name,
				m_product_sub_category_temp.prod_sub_ctgry_id as prod_sub_ctgry_temp_id,
				m_product_sub_category_temp.prod_sub_ctgry_name as prod_sub_ctgry_temp_name,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_reason,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_status,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_action_before,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_created_by,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_created_at`).
		Joins("JOIN m_mitra_product_temp on m_mitra_product_sub_category_temp.mitra_prod_temp_id=m_mitra_product_temp.mitra_prod_temp_id").
		Joins("JOIN m_product_sub_category m_product_sub_category_temp on m_mitra_product_sub_category_temp.prod_sub_ctgry_temp_id=m_product_sub_category_temp.prod_sub_ctgry_id").
		Where("m_mitra_product_sub_category_temp.mitra_prod_temp_id = ?", pid).
		Find(&mitraproductsubcategory).Error
	if err != nil {
		return []MitraProductSubCategoryTempData{}, err
	}
	return mitraproductsubcategory, nil
}

func (p *MitraProductSubCategoryTemp) FindMitraProductSubCategoryTempStatusByMitraProductTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraProductSubCategoryTempPendData, error) {
	var err error
	mitraproductsubcategory := []MitraProductSubCategoryTempPendData{}
	err = db.Debug().Model(&MitraProductSubCategoryTempPendData{}).Limit(100).
		Select(`m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_id,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_id,
				m_mitra_product_temp.mitra_prod_temp_id,
				m_mitra_product_temp.mitra_prod_temp_name,
				m_product_sub_category_temp.prod_sub_ctgry_id as prod_sub_ctgry_temp_id,
				m_product_sub_category_temp.prod_sub_ctgry_name as prod_sub_ctgry_temp_name,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_reason,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_status,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_action_before,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_created_by,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_created_at`).
		Joins("JOIN m_mitra_product_temp on m_mitra_product_sub_category_temp.mitra_prod_temp_id=m_mitra_product_temp.mitra_prod_temp_id").
		Joins("JOIN m_product_sub_category m_product_sub_category_temp on m_mitra_product_sub_category_temp.prod_sub_ctgry_temp_id=m_product_sub_category_temp.prod_sub_ctgry_id").
		Find(&mitraproductsubcategory).Error
	if err != nil {
		return []MitraProductSubCategoryTempPendData{}, err
	}
	return mitraproductsubcategory, nil
}

func (p *MitraProductSubCategoryTemp) FindMitraProductSubCategoryTempDataStatusByMitraProductTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraProductSubCategoryTempData, error) {
	var err error
	mitraproductsubcategory := []MitraProductSubCategoryTempData{}
	err = db.Debug().Model(&MitraProductSubCategoryTempData{}).Limit(100).
		Select(`m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_id,
				m_mitra_product_temp.mitra_prod_temp_id,
				m_mitra_product_temp.mitra_prod_temp_name,
				m_product_sub_category_temp.prod_sub_ctgry_id as prod_sub_ctgry_temp_id,
				m_product_sub_category_temp.prod_sub_ctgry_name as prod_sub_ctgry_temp_name,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_id,
				m_mitra_product.mitra_prod_id,
				m_mitra_product.mitra_prod_name,
				m_product_sub_category.prod_sub_ctgry_id,
				m_product_sub_category.prod_sub_ctgry_name,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_reason,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_status,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_action_before,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_created_by,
				m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_temp_created_at`).
		Joins("JOIN m_mitra_product_temp on m_mitra_product_sub_category_temp.mitra_prod_temp_id=m_mitra_product_temp.mitra_prod_temp_id").
		Joins("JOIN m_product_sub_category m_product_sub_category_temp on m_mitra_product_sub_category_temp.prod_sub_ctgry_temp_id=m_product_sub_category_temp.prod_sub_ctgry_id").
		Joins("JOIN m_mitra_product_sub_category on m_mitra_product_sub_category_temp.mitra_prod_sub_ctgry_id=m_mitra_product_sub_category.mitra_prod_sub_ctgry_id").
		Joins("JOIN m_mitra_product on m_mitra_product_sub_category.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_sub_category on m_mitra_product_sub_category.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Where("m_mitra_product_sub_category_temp.mitra_prod_temp_id = ? AND mitra_prod_sub_ctgry_temp_status = ?", pid, status).
		Find(&mitraproductsubcategory).Error
	if err != nil {
		return []MitraProductSubCategoryTempData{}, err
	}
	return mitraproductsubcategory, nil
}
