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

type ProductCostTemp struct {
	ProductCostTempID           uint64    `gorm:"column:prod_cost_temp_id;primary_key;" json:"prod_cost_temp_id"`
	ProductCostID               uint64    `gorm:"column:prod_cost_id" json:"prod_cost_id"`
	ProductCostTempCode         string    `gorm:"column:prod_cost_temp_code" json:"prod_cost_temp_code"`
	ProductCostTempName         string    `gorm:"column:prod_cost_temp_name" json:"prod_cost_temp_name"`
	ProductCostTempDesc         string    `gorm:"column:prod_cost_temp_desc" json:"prod_cost_temp_desc"`
	ProductCostTempReason       string    `gorm:"column:prod_cost_temp_reason" json:"prod_cost_temp_reason"`
	ProductCostTempStatus       uint64    `gorm:"column:prod_cost_temp_status;size:2" json:"prod_cost_temp_status"`
	ProductCostTempActionBefore uint64    `gorm:"column:prod_cost_temp_action_before;size:2" json:"prod_cost_temp_action_before"`
	ProductCostTempCreatedBy    string    `gorm:"column:prod_cost_temp_created_by;size:125" json:"prod_cost_temp_created_by"`
	ProductCostTempCreatedAt    time.Time `gorm:"column:prod_cost_temp_created_at;default:CURRENT_TIMESTAMP" json:"prod_cost_temp_created_at"`
}

type ProductCostTempPend struct {
	ProductCostTempID           uint64    `gorm:"column:prod_cost_temp_id;primary_key;" json:"prod_cost_temp_id"`
	ProductCostID               *uint64   `gorm:"column:prod_cost_id" json:"prod_cost_id"`
	ProductCostTempCode         string    `gorm:"column:prod_cost_temp_code" json:"prod_cost_temp_code"`
	ProductCostTempName         string    `gorm:"column:prod_cost_temp_name" json:"prod_cost_temp_name"`
	ProductCostTempDesc         string    `gorm:"column:prod_cost_temp_desc" json:"prod_cost_temp_desc"`
	ProductCostTempReason       string    `gorm:"column:prod_cost_temp_reason" json:"prod_cost_temp_reason"`
	ProductCostTempStatus       uint64    `gorm:"column:prod_cost_temp_status;size:2" json:"prod_cost_temp_status"`
	ProductCostTempActionBefore uint64    `gorm:"column:prod_cost_temp_action_before;size:2" json:"prod_cost_temp_action_before"`
	ProductCostTempCreatedBy    string    `gorm:"column:prod_cost_temp_created_by;size:125" json:"prod_cost_temp_created_by"`
	ProductCostTempCreatedAt    time.Time `gorm:"column:prod_cost_temp_created_at;default:CURRENT_TIMESTAMP" json:"prod_cost_temp_created_at"`
}

type ProductCostTempData struct {
	ProductCostTempID           uint64    `gorm:"column:prod_cost_temp_id" json:"prod_cost_temp_id"`
	ProductCostTempCode         string    `gorm:"column:prod_cost_temp_code" json:"prod_cost_temp_code"`
	ProductCostTempName         string    `gorm:"column:prod_cost_temp_name" json:"prod_cost_temp_name"`
	ProductCostTempDesc         string    `gorm:"column:prod_cost_temp_desc" json:"prod_cost_temp_desc"`
	ProductCostTempReason       string    `gorm:"column:prod_cost_temp_reason" json:"prod_cost_temp_reason"`
	ProductCostTempStatus       uint64    `gorm:"column:prod_cost_temp_status;size:2" json:"prod_cost_temp_status"`
	ProductCostTempActionBefore uint64    `gorm:"column:prod_cost_temp_action_before;size:2" json:"prod_cost_temp_action_before"`
	ProductCostTempCreatedBy    string    `gorm:"column:prod_cost_temp_created_by;size:125" json:"prod_cost_temp_created_by"`
	ProductCostTempCreatedAt    time.Time `gorm:"column:prod_cost_temp_created_at" json:"prod_cost_temp_created_at"`
	ProductCostID               uint64    `gorm:"column:prod_cost_id" json:"prod_cost_id"`
	ProductCostCode             string    `gorm:"column:prod_cost_code" json:"prod_cost_code"`
	ProductCostName             string    `gorm:"column:prod_cost_name" json:"prod_cost_name"`
	ProductCostDesc             string    `gorm:"column:prod_cost_desc" json:"prod_cost_desc"`
	ProductCostCreatedBy        string    `gorm:"column:prod_cost_created_by" json:"prod_cost_created_by"`
	ProductCostCreatedAt        time.Time `gorm:"column:prod_cost_created_at" json:"prod_cost_created_at"`
	ProductCostUpdatedBy        string    `gorm:"column:prod_cost_updated_by" json:"prod_cost_updated_by"`
	ProductCostUpdatedAt        time.Time `gorm:"column:prod_cost_updated_at" json:"prod_cost_updated_at"`
	ProductCostDeletedBy        string    `gorm:"column:prod_cost_deleted_by" json:"prod_cost_deleted_by"`
	ProductCostDeletedAt        time.Time `gorm:"column:prod_cost_deleted_at" json:"prod_cost_deleted_at"`
}

type ResponseProductCostTemps struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    *[]ProductCostTempData `json:"data"`
}

type ResponseProductCostTemp struct {
	Status  int                  `json:"status"`
	Message string               `json:"message"`
	Data    *ProductCostTempData `json:"data"`
}

type ResponseProductCostTempsPend struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    *[]ProductCostTempPend `json:"data"`
}

type ResponseProductCostTempPend struct {
	Status  int                  `json:"status"`
	Message string               `json:"message"`
	Data    *ProductCostTempPend `json:"data"`
}

type ResponseProductCostTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (ProductCostTemp) TableName() string {
	return "m_product_cost_temp"
}

func (ProductCostTempPend) TableName() string {
	return "m_product_cost_temp"
}

func (ProductCostTempData) TableName() string {
	return "m_product_cost_temp"
}

func (p *ProductCostTemp) Prepare() {
	p.ProductCostID = p.ProductCostID
	p.ProductCostTempCode = html.EscapeString(strings.TrimSpace(p.ProductCostTempCode))
	p.ProductCostTempName = html.EscapeString(strings.TrimSpace(p.ProductCostTempName))
	p.ProductCostTempDesc = p.ProductCostTempDesc
	p.ProductCostTempReason = p.ProductCostTempReason
	p.ProductCostTempStatus = p.ProductCostTempStatus
	p.ProductCostTempCreatedBy = html.EscapeString(strings.TrimSpace(p.ProductCostTempCreatedBy))
	p.ProductCostTempCreatedAt = time.Now()
}

func (p *ProductCostTempPend) Prepare() {
	p.ProductCostID = nil
	p.ProductCostTempCode = html.EscapeString(strings.TrimSpace(p.ProductCostTempCode))
	p.ProductCostTempName = html.EscapeString(strings.TrimSpace(p.ProductCostTempName))
	p.ProductCostTempDesc = p.ProductCostTempDesc
	p.ProductCostTempReason = p.ProductCostTempReason
	p.ProductCostTempStatus = p.ProductCostTempStatus
	p.ProductCostTempCreatedBy = html.EscapeString(strings.TrimSpace(p.ProductCostTempCreatedBy))
	p.ProductCostTempCreatedAt = time.Now()
}

func (p *ProductCostTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	case "insertupdate":
		if p.ProductCostTempCode == "" {
			return errors.New("Negara Wajib diisi")
		}
		if p.ProductCostTempName == "" {
			return errors.New("Provinsi Wajib diisi")
		}
		if p.ProductCostTempDesc == "" {
			return errors.New("Kabupaten / Kota Wajib diisi")
		}
		return nil

	case "reason":
		if p.ProductCostTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.ProductCostTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	default:
		if p.ProductCostTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil
	}
}

func (p *ProductCostTemp) SaveProductCostTemp(db *gorm.DB) (*ProductCostTemp, error) {
	var err error
	err = db.Debug().Model(&ProductCostTemp{}).Create(&p).Error
	if err != nil {
		return &ProductCostTemp{}, err
	}
	return p, nil
}

func (p *ProductCostTempPend) SaveProductCostTempPend(db *gorm.DB) (*ProductCostTempPend, error) {
	var err error
	err = db.Debug().Model(&ProductCostTempPend{}).Create(&p).Error
	if err != nil {
		return &ProductCostTempPend{}, err
	}
	return p, nil
}

func (p *ProductCostTemp) FindAllProductCostTemp(db *gorm.DB) (*[]ProductCostTempPend, error) {
	var err error
	product_cost := []ProductCostTempPend{}
	err = db.Debug().Model(&ProductCostTempPend{}).Limit(100).
		Select(`m_product_cost_temp.prod_cost_temp_id,
				m_product_cost_temp.prod_cost_id,
				m_product_cost_temp.prod_cost_temp_code,
				m_product_cost_temp.prod_cost_temp_name,
				m_product_cost_temp.prod_cost_temp_desc,
				m_product_cost_temp.prod_cost_temp_reason,
				m_product_cost_temp.prod_cost_temp_status,
				m_product_cost_temp.prod_cost_temp_action_before,
				m_product_cost_temp.prod_cost_temp_created_by,
				m_product_cost_temp.prod_cost_temp_created_at at time zone 'Asia/Jakarta' as prod_cost_temp_created_at`).
		Order("prod_cost_temp_created_at desc").
		Find(&product_cost).Error
	if err != nil {
		return &[]ProductCostTempPend{}, err
	}
	return &product_cost, nil
}

func (p *ProductCostTemp) FindAllProductCostTempPendingActive(db *gorm.DB) (*[]ProductCostTempPend, error) {
	var err error
	product_cost := []ProductCostTempPend{}
	err = db.Debug().Model(&ProductCostTempPend{}).Limit(100).
		Select(`m_product_cost_temp.prod_cost_temp_id,
				m_product_cost_temp.prod_cost_id,
				m_product_cost_temp.prod_cost_temp_code,
				m_product_cost_temp.prod_cost_temp_name,
				m_product_cost_temp.prod_cost_temp_desc,
				m_product_cost_temp.prod_cost_temp_reason,
				m_product_cost_temp.prod_cost_temp_status,
				m_product_cost_temp.prod_cost_temp_action_before,
				m_product_cost_temp.prod_cost_temp_created_by,
				m_product_cost_temp.prod_cost_temp_created_at at time zone 'Asia/Jakarta' as prod_cost_temp_created_at`).
		Where("prod_cost_temp_status = ?", 11).
		Order("prod_cost_temp_created_at desc").
		Find(&product_cost).Error
	if err != nil {
		return &[]ProductCostTempPend{}, err
	}
	return &product_cost, nil
}

func (p *ProductCostTemp) FindAllProductCostTempStatus(db *gorm.DB, status uint64) (*[]ProductCostTempData, error) {
	var err error
	product_cost := []ProductCostTempData{}
	err = db.Debug().Model(&ProductCostTempData{}).Limit(100).
		Select(`m_product_cost_temp.prod_cost_temp_id,
				m_product_cost_temp.prod_cost_temp_code,
				m_product_cost_temp.prod_cost_temp_name,
				m_product_cost_temp.prod_cost_temp_desc,
				m_product_cost_temp.prod_cost_temp_reason,
				m_product_cost_temp.prod_cost_temp_status,
				m_product_cost_temp.prod_cost_temp_action_before,
				m_product_cost_temp.prod_cost_temp_created_by,
				m_product_cost_temp.prod_cost_temp_created_at at time zone 'Asia/Jakarta' as prod_cost_temp_created_at,
				m_product_cost.prod_cost_id,
				m_product_cost.prod_cost_code,
				m_product_cost.prod_cost_name,
				m_product_cost.prod_cost_desc,
				m_product_cost.prod_cost_status,
				m_product_cost.prod_cost_created_by,
				m_product_cost.prod_cost_created_at at time zone 'Asia/Jakarta' as prod_cost_created_at,
				m_product_cost.prod_cost_updated_by,
				m_product_cost.prod_cost_updated_at at time zone 'Asia/Jakarta' as prod_cost_updated_at,
				m_product_cost.prod_cost_deleted_by,
				m_product_cost.prod_cost_deleted_at at time zone 'Asia/Jakarta' as prod_cost_deleted_at`).
		Joins("JOIN m_product_cost on m_product_cost_temp.prod_cost_id=m_product_cost.prod_cost_id").
		Where("prod_cost_temp_status = ?", status).
		Order("prod_cost_temp_created_at desc").
		Find(&product_cost).Error
	if err != nil {
		return &[]ProductCostTempData{}, err
	}
	return &product_cost, nil
}

func (p *ProductCostTemp) FindProductCostTempDataByID(db *gorm.DB, pid uint64) (*ProductCostTemp, error) {
	var err error
	err = db.Debug().Model(&ProductCostTemp{}).Where("prod_cost_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &ProductCostTemp{}, err
	}
	return p, nil
}

func (p *ProductCostTemp) FindProductCostTempByIDPendingActive(db *gorm.DB, pid uint64) (*ProductCostTempPend, error) {
	var err error
	product_cost := ProductCostTempPend{}
	err = db.Debug().Model(&ProductCostTempPend{}).Limit(100).
		Select(`m_product_cost_temp.prod_cost_temp_id,
				m_product_cost_temp.prod_cost_id,
				m_product_cost_temp.prod_cost_temp_code,
				m_product_cost_temp.prod_cost_temp_name,
				m_product_cost_temp.prod_cost_temp_desc,
				m_product_cost_temp.prod_cost_temp_reason,
				m_product_cost_temp.prod_cost_temp_status,
				m_product_cost_temp.prod_cost_temp_created_by,
				m_product_cost_temp.prod_cost_temp_created_at at time zone 'Asia/Jakarta' as prod_cost_temp_created_at`).
		Where("prod_cost_temp_id = ?", pid).
		Order("prod_cost_temp_created_at desc").
		Find(&product_cost).Error
	if err != nil {
		return &ProductCostTempPend{}, err
	}
	return &product_cost, nil
}

func (p *ProductCostTemp) FindProductCostTempStatusByIDPendingActive(db *gorm.DB, pid uint64) (*ProductCostTempPend, error) {
	var err error
	product_cost := ProductCostTempPend{}
	err = db.Debug().Model(&ProductCostTempPend{}).Limit(100).
		Select(`m_product_cost_temp.prod_cost_temp_id,
				m_product_cost_temp.prod_cost_id,
				m_product_cost_temp.prod_cost_temp_code,
				m_product_cost_temp.prod_cost_temp_name,
				m_product_cost_temp.prod_cost_temp_desc,
				m_product_cost_temp.prod_cost_temp_reason,
				m_product_cost_temp.prod_cost_temp_status,
				m_product_cost_temp.prod_cost_temp_action_before,
				m_product_cost_temp.prod_cost_temp_created_by,
				m_product_cost_temp.prod_cost_temp_created_at at time zone 'Asia/Jakarta' as prod_cost_temp_created_at`).
		Where("prod_cost_temp_id = ? AND prod_cost_temp_status = ?", pid, 11).
		Order("prod_cost_temp_created_at desc").
		Find(&product_cost).Error
	if err != nil {
		return &ProductCostTempPend{}, err
	}
	return &product_cost, nil
}

func (p *ProductCostTemp) FindProductCostTempByID(db *gorm.DB, pid uint64) (*ProductCostTempData, error) {
	var err error
	product_cost := ProductCostTempData{}
	err = db.Debug().Model(&ProductCostTempData{}).Limit(100).
		Select(`m_product_cost_temp.prod_cost_temp_id,
				m_product_cost_temp.prod_cost_temp_code,
				m_product_cost_temp.prod_cost_temp_name,
				m_product_cost_temp.prod_cost_temp_desc,
				m_product_cost_temp.prod_cost_temp_reason,
				m_product_cost_temp.prod_cost_temp_status,
				m_product_cost_temp.prod_cost_temp_action_before,
				m_product_cost_temp.prod_cost_temp_created_by,
				m_product_cost_temp.prod_cost_temp_created_at at time zone 'Asia/Jakarta' as prod_cost_temp_created_at,
				m_product_cost.prod_cost_id,
				m_product_cost.prod_cost_code,
				m_product_cost.prod_cost_name,
				m_product_cost.prod_cost_desc,
				m_product_cost.prod_cost_status,
				m_product_cost.prod_cost_created_by,
				m_product_cost.prod_cost_created_at at time zone 'Asia/Jakarta' as prod_cost_created_at,
				m_product_cost.prod_cost_updated_by,
				m_product_cost.prod_cost_updated_at at time zone 'Asia/Jakarta' as prod_cost_updated_at,
				m_product_cost.prod_cost_deleted_by,
				m_product_cost.prod_cost_deleted_at at time zone 'Asia/Jakarta' as prod_cost_deleted_at`).
		Joins("JOIN m_product_cost on m_product_cost_temp.prod_cost_id=m_product_cost.prod_cost_id").
		Where("prod_cost_temp_id = ?", pid).
		Order("prod_cost_temp_created_at desc").
		Find(&product_cost).Error
	if err != nil {
		return &ProductCostTempData{}, err
	}
	return &product_cost, nil
}

func (p *ProductCostTemp) FindProductCostTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*ProductCostTempData, error) {
	var err error
	product_cost := ProductCostTempData{}
	err = db.Debug().Model(&ProductCostTempData{}).Limit(100).
		Select(`m_product_cost_temp.prod_cost_temp_id,
				m_product_cost_temp.prod_cost_temp_code,
				m_product_cost_temp.prod_cost_temp_name,
				m_product_cost_temp.prod_cost_temp_desc,
				m_product_cost_temp.prod_cost_temp_reason,
				m_product_cost_temp.prod_cost_temp_status,
				m_product_cost_temp.prod_cost_temp_action_before,
				m_product_cost_temp.prod_cost_temp_created_by,
				m_product_cost_temp.prod_cost_temp_created_at at time zone 'Asia/Jakarta' as prod_cost_temp_created_at,
				m_product_cost.prod_cost_id,
				m_product_cost.prod_cost_code,
				m_product_cost.prod_cost_name,
				m_product_cost.prod_cost_desc,
				m_product_cost.prod_cost_status,
				m_product_cost.prod_cost_created_by,
				m_product_cost.prod_cost_created_at at time zone 'Asia/Jakarta' as prod_cost_created_at,
				m_product_cost.prod_cost_updated_by,
				m_product_cost.prod_cost_updated_at at time zone 'Asia/Jakarta' as prod_cost_updated_at,
				m_product_cost.prod_cost_deleted_by,
				m_product_cost.prod_cost_deleted_at at time zone 'Asia/Jakarta' as prod_cost_deleted_at`).
		Joins("JOIN m_product_cost on m_product_cost_temp.prod_cost_id=m_product_cost.prod_cost_id").
		Where("prod_cost_temp_id = ? AND prod_cost_temp_status = ?", pid, status).
		Order("prod_cost_temp_created_at desc").
		Find(&product_cost).Error
	if err != nil {
		return &ProductCostTempData{}, err
	}
	return &product_cost, nil
}

func (p *ProductCostTemp) UpdateProductCostTemp(db *gorm.DB, pid uint64) (*ProductCostTemp, error) {
	var err error
	db = db.Debug().Model(&ProductCostTemp{}).Where("prod_cost_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_cost_temp_code":   p.ProductCostTempCode,
			"prod_cost_temp_name":   p.ProductCostTempName,
			"prod_cost_temp_desc":   p.ProductCostTempDesc,
			"prod_cost_temp_reason": p.ProductCostTempReason,
		},
	)
	err = db.Debug().Model(&ProductCostTemp{}).Where("prod_cost_temp_id = ?", pid).Error
	if err != nil {
		return &ProductCostTemp{}, err
	}
	return p, nil
}

func (p *ProductCostTemp) UpdateProductCostTempStatus(db *gorm.DB, pid uint64) (*ProductCostTemp, error) {
	var err error
	db = db.Debug().Model(&ProductCostTemp{}).Where("prod_cost_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_cost_temp_reason": p.ProductCostTempReason,
			"prod_cost_temp_status": p.ProductCostTempStatus,
		},
	)
	err = db.Debug().Model(&ProductCostTemp{}).Where("prod_cost_temp_id = ?", pid).Error
	if err != nil {
		return &ProductCostTemp{}, err
	}
	return p, nil
}

func (p *ProductCostTemp) DeleteProductCostTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&ProductCostTemp{}).Where("prod_cost_temp_id = ?", pid).Take(&ProductCostTemp{}).Delete(&ProductCostTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("ProductCostTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================
