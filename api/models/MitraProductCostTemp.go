package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type MitraProductCostTemp struct {
	MitraProductCostTempID           *uint64   `gorm:"column:mitra_prod_cost_temp_id;primary_key;" json:"mitra_prod_cost_temp_id"`
	MitraProductCostID               uint64    `gorm:"column:mitra_prod_cost_id" json:"mitra_prod_cost_id"`
	MitraProductTempID               uint64    `gorm:"column:mitra_prod_temp_id" json:"mitra_prod_temp_id"`
	ProductCostTempID                uint64    `gorm:"column:prod_cost_temp_id" json:"prod_cost_temp_id"`
	MitraProductCostTempSign         string    `gorm:"column:mitra_prod_cost_temp_sign;size:25" json:"mitra_prod_cost_temp_sign"`
	MitraProductCostTempPeriod       string    `gorm:"column:mitra_prod_cost_temp_period;size:255" json:"mitra_prod_cost_temp_period"`
	MitraProductCostTempValue        uint64    `gorm:"column:mitra_prod_cost_temp_value;size:255" json:"mitra_prod_cost_temp_value"`
	MitraProductCostTempReason       string    `gorm:"column:mitra_prod_cost_temp_reason" json:"mitra_prod_cost_temp_reason"`
	MitraProductCostTempActionBefore uint64    `gorm:"column:mitra_prod_cost_temp_action_before" json:"mitra_prod_cost_temp_action_before"`
	MitraProductCostTempStatus       uint64    `gorm:"column:mitra_prod_cost_temp_status;size:25" json:"mitra_prod_cost_temp_status"`
	MitraProductCostTempCreatedBy    string    `gorm:"column:mitra_prod_cost_temp_created_by;size:125" json:"mitra_prod_cost_temp_created_by"`
	MitraProductCostTempCreatedAt    time.Time `gorm:"column:mitra_prod_cost_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_cost_temp_created_at"`
}

type MitraProductCostTempPend struct {
	MitraProductCostTempID           *uint64   `gorm:"column:mitra_prod_cost_temp_id;primary_key;" json:"mitra_prod_cost_temp_id"`
	MitraProductCostID               *uint64   `gorm:"column:mitra_prod_cost_id" json:"mitra_prod_cost_id"`
	MitraProductTempID               uint64    `gorm:"column:mitra_prod_temp_id" json:"mitra_prod_temp_id"`
	ProductCostTempID                uint64    `gorm:"column:prod_cost_temp_id" json:"prod_cost_temp_id"`
	MitraProductCostTempSign         string    `gorm:"column:mitra_prod_cost_temp_sign;size:25" json:"mitra_prod_cost_temp_sign"`
	MitraProductCostTempPeriod       string    `gorm:"column:mitra_prod_cost_temp_period;size:255" json:"mitra_prod_cost_temp_period"`
	MitraProductCostTempValue        uint64    `gorm:"column:mitra_prod_cost_temp_value;size:255" json:"mitra_prod_cost_temp_value"`
	MitraProductCostTempReason       string    `gorm:"column:mitra_prod_cost_temp_reason" json:"mitra_prod_cost_temp_reason"`
	MitraProductCostTempActionBefore uint64    `gorm:"column:mitra_prod_cost_temp_action_before" json:"mitra_prod_cost_temp_action_before"`
	MitraProductCostTempStatus       uint64    `gorm:"column:mitra_prod_cost_temp_status;size:25" json:"mitra_prod_cost_temp_status"`
	MitraProductCostTempCreatedBy    string    `gorm:"column:mitra_prod_cost_temp_created_by;size:125" json:"mitra_prod_cost_temp_created_by"`
	MitraProductCostTempCreatedAt    time.Time `gorm:"column:mitra_prod_cost_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_cost_temp_created_at"`
}

type MitraProductCostTempPendData struct {
	MitraProductCostTempID           uint64    `gorm:"column:mitra_prod_cost_temp_id;primary_key;" json:"mitra_prod_cost_temp_id"`
	MitraProductCostID               uint64    `gorm:"column:mitra_prod_cost_id" json:"mitra_prod_cost_id"`
	MitraProductTempID               uint64    `gorm:"column:mitra_prod_temp_id" json:"mitra_prod_temp_id"`
	MitraProductTempName             string    `gorm:"column:mitra_prod_temp_name" json:"mitra_prod_temp_name"`
	ProductCostTempID                uint64    `gorm:"column:prod_cost_temp_id" json:"prod_cost_temp_id"`
	ProductCostTempName              string    `gorm:"column:prod_cost_temp_name" json:"prod_cost_temp_name"`
	MitraProductCostTempSign         string    `gorm:"column:mitra_prod_cost_temp_sign;size:25" json:"mitra_prod_cost_temp_sign"`
	MitraProductCostTempPeriod       string    `gorm:"column:mitra_prod_cost_temp_period;size:255" json:"mitra_prod_cost_temp_period"`
	MitraProductCostTempValue        uint64    `gorm:"column:mitra_prod_cost_temp_value;size:255" json:"mitra_prod_cost_temp_value"`
	MitraProductCostTempReason       string    `gorm:"column:mitra_prod_cost_temp_reason" json:"mitra_prod_cost_temp_reason"`
	MitraProductCostTempActionBefore uint64    `gorm:"column:mitra_prod_cost_temp_action_before" json:"mitra_prod_cost_temp_action_before"`
	MitraProductCostTempStatus       uint64    `gorm:"column:mitra_prod_cost_temp_status;size:25" json:"mitra_prod_cost_temp_status"`
	MitraProductCostTempCreatedBy    string    `gorm:"column:mitra_prod_cost_temp_created_by;size:125" json:"mitra_prod_cost_temp_created_by"`
	MitraProductCostTempCreatedAt    time.Time `gorm:"column:mitra_prod_cost_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_cost_temp_created_at"`
}

type MitraProductCostTempData struct {
	MitraProductCostTempID           uint64     `gorm:"column:mitra_prod_cost_temp_id;primary_key;" json:"mitra_prod_cost_temp_id"`
	MitraProductTempID               uint64     `gorm:"column:mitra_prod_temp_id" json:"mitra_prod_temp_id"`
	MitraProductTempName             string     `gorm:"column:mitra_prod_temp_name" json:"mitra_prod_temp_name"`
	ProductCostTempID                uint64     `gorm:"column:prod_cost_temp_id" json:"prod_cost_temp_id"`
	ProductCostTempName              string     `gorm:"column:prod_cost_temp_name" json:"prod_cost_temp_name"`
	MitraProductCostTempSign         string     `gorm:"column:mitra_prod_cost_temp_sign;size:25" json:"mitra_prod_cost_temp_sign"`
	MitraProductCostTempPeriod       string     `gorm:"column:mitra_prod_cost_temp_period;size:255" json:"mitra_prod_cost_temp_period"`
	MitraProductCostTempValue        uint64     `gorm:"column:mitra_prod_cost_temp_value;size:255" json:"mitra_prod_cost_temp_value"`
	MitraProductCostTempReason       string     `gorm:"column:mitra_prod_cost_temp_reason" json:"mitra_prod_cost_temp_reason"`
	MitraProductCostTempActionBefore uint64     `gorm:"column:mitra_prod_cost_temp_action_before" json:"mitra_prod_cost_temp_action_before"`
	MitraProductCostTempStatus       uint64     `gorm:"column:mitra_prod_cost_temp_status;size:25" json:"mitra_prod_cost_temp_status"`
	MitraProductCostTempCreatedBy    string     `gorm:"column:mitra_prod_cost_temp_created_by;size:125" json:"mitra_prod_cost_temp_created_by"`
	MitraProductCostTempCreatedAt    time.Time  `gorm:"column:mitra_prod_cost_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_cost_temp_created_at"`
	MitraProductCostID               uint64     `gorm:"column:mitra_prod_cost_id;primary_key;" json:"mitra_prod_cost_id"`
	MitraProductID                   uint64     `gorm:"column:mitra_prod_id" json:"mitra_prod_id"`
	MitraProductName                 string     `gorm:"column:mitra_prod_name" json:"mitra_prod_name"`
	ProductCostID                    uint64     `gorm:"column:prod_cost_id" json:"prod_cost_id"`
	ProductCostName                  string     `gorm:"column:prod_cost_name" json:"prod_cost_name"`
	MitraProductCostSign             string     `gorm:"column:mitra_prod_cost_sign;size:25" json:"mitra_prod_cost_sign"`
	MitraProductCostPeriod           string     `gorm:"column:mitra_prod_cost_period;size:255" json:"mitra_prod_cost_period"`
	MitraProductCostValue            uint64     `gorm:"column:mitra_prod_cost_value;size:255" json:"mitra_prod_cost_value"`
	MitraProductCostStatus           uint64     `gorm:"column:mitra_prod_cost_status;size:25" json:"mitra_prod_cost_status"`
	MitraProductCostCreatedBy        string     `gorm:"column:mitra_prod_cost_created_by;size:125" json:"mitra_prod_cost_created_by"`
	MitraProductCostCreatedAt        time.Time  `gorm:"column:mitra_prod_cost_created_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_cost_created_at"`
	MitraProductCostUpdatedBy        string     `gorm:"column:mitra_prod_cost_updated_by;size:125" json:"mitra_prod_cost_updated_by"`
	MitraProductCostUpdatedAt        *time.Time `gorm:"column:mitra_prod_cost_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_cost_updated_at"`
	MitraProductCostDeletedBy        string     `gorm:"column:mitra_prod_cost_deleted_by;size:125" json:"mitra_prod_cost_deleted_by"`
	MitraProductCostDeletedAt        *time.Time `gorm:"column:mitra_prod_cost_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_cost_deleted_at"`
}

type ResponseMitraProductCostTemps struct {
	Status  int                         `json:"status"`
	Message string                      `json:"message"`
	Data    *[]MitraProductCostTempData `json:"data"`
}

type ResponseMitraProductCostTempsPend struct {
	Status  int                         `json:"status"`
	Message string                      `json:"message"`
	Data    *[]MitraProductCostTempPend `json:"data"`
}

type ResponseMitraProductCostTemp struct {
	Status  int                       `json:"status"`
	Message string                    `json:"message"`
	Data    *MitraProductCostTempData `json:"data"`
}

type ResponseMitraProductCostTempPend struct {
	Status  int                       `json:"status"`
	Message string                    `json:"message"`
	Data    *MitraProductCostTempPend `json:"data"`
}

type ResponseMitraProductCostTempIU struct {
	Status  int                   `json:"status"`
	Message string                `json:"message"`
	Data    *MitraProductCostTemp `json:"data"`
}

type ResponseMitraProductCostTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraProductCostTemp) TableName() string {
	return "m_mitra_product_cost_temp"
}

func (MitraProductCostTempPend) TableName() string {
	return "m_mitra_product_cost_temp"
}

func (MitraProductCostTempPendData) TableName() string {
	return "m_mitra_product_cost_temp"
}

func (MitraProductCostTempData) TableName() string {
	return "m_mitra_product_cost_temp"
}

func (p *MitraProductCostTemp) Prepare() {
	p.MitraProductCostTempID = nil
	p.MitraProductCostID = p.MitraProductCostID
	p.MitraProductTempID = p.MitraProductTempID
	p.ProductCostTempID = p.ProductCostTempID
	p.MitraProductCostTempSign = html.EscapeString(strings.TrimSpace(p.MitraProductCostTempSign))
	p.MitraProductCostTempPeriod = html.EscapeString(strings.TrimSpace(p.MitraProductCostTempPeriod))
	p.MitraProductCostTempValue = p.MitraProductCostTempValue
	p.MitraProductCostTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraProductCostTempCreatedBy))
	p.MitraProductCostTempReason = html.EscapeString(strings.TrimSpace(p.MitraProductCostTempReason))
	p.MitraProductCostTempStatus = p.MitraProductCostTempStatus
	p.MitraProductCostTempCreatedAt = time.Now()
}

func (p *MitraProductCostTempPend) Prepare() {
	p.MitraProductCostTempID = nil
	p.MitraProductCostID = nil
	p.MitraProductTempID = p.MitraProductTempID
	p.ProductCostTempID = p.ProductCostTempID
	p.MitraProductCostTempSign = html.EscapeString(strings.TrimSpace(p.MitraProductCostTempSign))
	p.MitraProductCostTempPeriod = html.EscapeString(strings.TrimSpace(p.MitraProductCostTempPeriod))
	p.MitraProductCostTempValue = p.MitraProductCostTempValue
	p.MitraProductCostTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraProductCostTempCreatedBy))
	p.MitraProductCostTempReason = html.EscapeString(strings.TrimSpace(p.MitraProductCostTempReason))
	p.MitraProductCostTempStatus = p.MitraProductCostTempStatus
	p.MitraProductCostTempCreatedAt = time.Now()
}

func (p *MitraProductCostTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.MitraProductTempID == 0 {
			return errors.New("Required MitraProductCost Code")
		}
		if p.ProductCostTempID == 0 {
			return errors.New("Required MitraProductCost Code")
		}
		if p.MitraProductCostTempSign == "" {
			return errors.New("Required MitraProductCost Name")
		}
		if p.MitraProductCostTempPeriod == "" {
			return errors.New("Required MitraProductCost Name")
		}
		if p.MitraProductCostTempValue == 0 {
			return errors.New("Required MitraProductCost Name")
		}
		return nil
	}
}

func (p *MitraProductCostTemp) SaveMitraProductCostTemp(db *gorm.DB) (*MitraProductCostTemp, error) {
	var err error
	err = db.Debug().Model(&MitraProductCostTemp{}).Create(&p).Error
	if err != nil {
		return &MitraProductCostTemp{}, err
	}
	return p, nil
}

func (p *MitraProductCostTempPend) SaveMitraProductCostTempPend(db *gorm.DB) (*MitraProductCostTempPend, error) {
	var err error
	err = db.Debug().Model(&MitraProductCostTempPend{}).Create(&p).Error
	if err != nil {
		return &MitraProductCostTempPend{}, err
	}
	return p, nil
}

func (p *MitraProductCostTemp) FindAllMitraProductCostTemp(db *gorm.DB) (*[]MitraProductCostTempPend, error) {
	var err error
	mitraproductcosttemp := []MitraProductCostTempPend{}
	err = db.Debug().Model(&MitraProductCostTempPend{}).Limit(100).
		Select(`m_mitra_product_cost_temp.mitra_prod_cost_temp_id,
				m_mitra_product_cost_temp.mitra_prod_cost_id,
				m_mitra_product_cost_temp.mitra_prod_temp_id,
				m_mitra_product_cost_temp.prod_cost_temp_id,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_sign,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_period,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_value,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_status,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_reason,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_created_by,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_created_at`).
		Find(&mitraproductcosttemp).Error
	if err != nil {
		return &[]MitraProductCostTempPend{}, err
	}
	return &mitraproductcosttemp, nil
}

func (p *MitraProductCostTemp) FindAllMitraProductCostTempStatusPendingActive(db *gorm.DB) (*[]MitraProductCostTempPend, error) {
	var err error
	mitraproductcosttemp := []MitraProductCostTempPend{}
	err = db.Debug().Model(&MitraProductCostTempPend{}).Limit(100).
		Select(`m_mitra_product_cost_temp.mitra_prod_cost_temp_id,
				m_mitra_product_cost_temp.mitra_prod_cost_id,
				m_mitra_product_cost_temp.mitra_prod_temp_id,
				m_mitra_product_cost_temp.prod_cost_temp_id,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_sign,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_period,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_value,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_status,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_reason,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_created_by,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_created_at`).
		Where("mitra_prod_cost_temp_status = ?", 11).
		Find(&mitraproductcosttemp).Error
	if err != nil {
		return &[]MitraProductCostTempPend{}, err
	}
	return &mitraproductcosttemp, nil
}

func (p *MitraProductCostTemp) FindAllMitraProductCostTempStatus(db *gorm.DB, status uint64) (*[]MitraProductCostTempData, error) {
	var err error
	mitraproductcosttemp := []MitraProductCostTempData{}
	err = db.Debug().Model(&MitraProductCostTempData{}).Limit(100).
		Select(`m_mitra_product_cost_temp.mitra_prod_cost_temp_id,
				m_mitra_product_temp.mitra_prod_temp_id,
				m_mitra_product_temp.mitra_prod_temp_name,
				m_product_cost_temp.prod_cost_id as prod_cost_temp_id,
				m_product_cost_temp.prod_cost_name as prod_cost_temp_name,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_sign,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_period,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_value,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_status,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_reason,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_created_by,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_created_at,
				m_mitra_product_cost.mitra_prod_cost_id,
				m_mitra_product.mitra_prod_id,
				m_mitra_product.mitra_prod_name,
				m_product_cost.prod_cost_id,
				m_product_cost.prod_cost_name,
				m_mitra_product_cost.mitra_prod_cost_sign,
				m_mitra_product_cost.mitra_prod_cost_period,
				m_mitra_product_cost.mitra_prod_cost_value,
				m_mitra_product_cost.mitra_prod_cost_status,
				m_mitra_product_cost.mitra_prod_cost_created_by,
				m_mitra_product_cost.mitra_prod_cost_updated_by,
				m_mitra_product_cost.mitra_prod_cost_deleted_by,
				m_mitra_product_cost.mitra_prod_cost_created_at,
				m_mitra_product_cost.mitra_prod_cost_updated_at,
				m_mitra_product_cost.mitra_prod_cost_deleted_at`).
		Joins("JOIN m_mitra_product_temp on m_mitra_product_cost_temp.mitra_prod_temp_id=m_mitra_product_temp.mitra_prod_temp_id").
		Joins("JOIN m_product_cost m_product_cost_temp on m_mitra_product_cost_temp.prod_cost_temp_id=m_product_cost_temp.prod_cost_id").
		Joins("JOIN m_mitra_product_cost on m_mitra_product_cost_temp.mitra_prod_id=m_mitra_product_cost.mitra_prod_id").
		Joins("JOIN m_mitra_product on m_mitra_product_cost.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_cost on m_mitra_product_cost.prod_cost_id=m_product_cost.prod_cost_id").
		Where("mitra_prod_cost_temp_status = ?", status).
		Find(&mitraproductcosttemp).Error
	if err != nil {
		return &[]MitraProductCostTempData{}, err
	}
	return &mitraproductcosttemp, nil
}

func (p *MitraProductCostTemp) FindAllMitraProductCostTempByMitraProduct(db *gorm.DB, mitrabranchtemp uint64) (*[]MitraProductCostTempPend, error) {
	var err error
	mitraproductcosttemp := []MitraProductCostTempPend{}
	err = db.Debug().Model(&MitraProductCostTempPend{}).Limit(100).
		Select(`m_mitra_product_cost_temp.mitra_prod_cost_temp_id,
				m_mitra_product_cost_temp.mitra_prod_cost_id,
				m_mitra_product_cost_temp.mitra_prod_temp_id,
				m_mitra_product_cost_temp.prod_cost_temp_id,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_sign,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_period,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_value,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_status,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_reason,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_created_by,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_created_at`).
		Where("mitra_prod_temp_id = ?", mitrabranchtemp).
		Find(&mitraproductcosttemp).Error
	if err != nil {
		return &[]MitraProductCostTempPend{}, err
	}
	return &mitraproductcosttemp, nil
}

func (p *MitraProductCostTemp) FindMitraProductCostTempDataByID(db *gorm.DB, pid uint64) (*MitraProductCostTemp, error) {
	var err error
	err = db.Debug().Model(&MitraProductCostTemp{}).Where("mitra_prod_cost_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraProductCostTemp{}, err
	}
	return p, nil
}

func (p *MitraProductCostTemp) FindMitraProductCostTempByIDPendingActive(db *gorm.DB, pid uint64) (*MitraProductCostTempPend, error) {
	var err error
	mitraproductcosttemp := MitraProductCostTempPend{}
	err = db.Debug().Model(&MitraProductCostTempPend{}).Limit(100).
		Select(`m_mitra_product_cost_temp.mitra_prod_cost_temp_id,
				m_mitra_product_cost_temp.mitra_prod_cost_id,
				m_mitra_product_cost_temp.mitra_prod_temp_id,
				m_mitra_product_cost_temp.prod_cost_temp_id,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_sign,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_period,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_value,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_status,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_reason,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_created_by,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_created_at`).
		Where("mitra_prod_cost_temp_id = ?", pid).
		Find(&mitraproductcosttemp).Error
	if err != nil {
		return &MitraProductCostTempPend{}, err
	}
	return &mitraproductcosttemp, nil
}

func (p *MitraProductCostTemp) FindMitraProductCostTempStatusByIDPendingActive(db *gorm.DB, pid uint64, status uint64) (*MitraProductCostTempPend, error) {
	var err error
	mitraproductcosttemp := MitraProductCostTempPend{}
	err = db.Debug().Model(&MitraProductCostTempPend{}).Limit(100).
		Select(`m_mitra_product_cost_temp.mitra_prod_cost_temp_id,
				m_mitra_product_cost_temp.mitra_prod_cost_id,
				m_mitra_product_cost_temp.mitra_prod_temp_id,
				m_mitra_product_cost_temp.prod_cost_temp_id,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_sign,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_period,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_value,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_status,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_reason,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_created_by,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_created_at`).
		Where("mitra_prod_cost_temp_id = ? AND mitra_prod_cost_temp_status = ?", pid, status).
		Find(&mitraproductcosttemp).Error
	if err != nil {
		return &MitraProductCostTempPend{}, err
	}
	return &mitraproductcosttemp, nil
}

func (p *MitraProductCostTemp) FindMitraProductCostTempByID(db *gorm.DB, pid uint64) (*MitraProductCostTempData, error) {
	var err error
	mitraproductcosttemp := MitraProductCostTempData{}
	err = db.Debug().Model(&MitraProductCostTempData{}).Limit(100).
		Select(`m_mitra_product_cost_temp.mitra_prod_cost_temp_id,
				m_mitra_product_temp.mitra_prod_temp_id,
				m_mitra_product_temp.mitra_prod_temp_name,
				m_product_cost_temp.prod_cost_id as prod_cost_temp_id,
				m_product_cost_temp.prod_cost_name as prod_cost_temp_name,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_sign,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_period,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_value,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_status,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_reason,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_created_by,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_created_at,
				m_mitra_product_cost.mitra_prod_cost_id,
				m_mitra_product.mitra_prod_id,
				m_mitra_product.mitra_prod_name,
				m_product_cost.prod_cost_id,
				m_product_cost.prod_cost_name,
				m_mitra_product_cost.mitra_prod_cost_sign,
				m_mitra_product_cost.mitra_prod_cost_period,
				m_mitra_product_cost.mitra_prod_cost_value,
				m_mitra_product_cost.mitra_prod_cost_status,
				m_mitra_product_cost.mitra_prod_cost_created_by,
				m_mitra_product_cost.mitra_prod_cost_updated_by,
				m_mitra_product_cost.mitra_prod_cost_deleted_by,
				m_mitra_product_cost.mitra_prod_cost_created_at,
				m_mitra_product_cost.mitra_prod_cost_updated_at,
				m_mitra_product_cost.mitra_prod_cost_deleted_at`).
		Joins("JOIN m_mitra_product_temp on m_mitra_product_cost_temp.mitra_prod_temp_id=m_mitra_product_temp.mitra_prod_temp_id").
		Joins("JOIN m_product_cost m_product_cost_temp on m_mitra_product_cost_temp.prod_cost_temp_id=m_product_cost_temp.prod_cost_id").
		Joins("JOIN m_mitra_product_cost on m_mitra_product_cost_temp.mitra_prod_id=m_mitra_product_cost.mitra_prod_id").
		Joins("JOIN m_mitra_product on m_mitra_product_cost.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_cost on m_mitra_product_cost.prod_cost_id=m_product_cost.prod_cost_id").
		Where("mitra_prod_cost_temp_id = ?", pid).
		Find(&mitraproductcosttemp).Error
	if err != nil {
		return &MitraProductCostTempData{}, err
	}
	return &mitraproductcosttemp, nil
}

func (p *MitraProductCostTemp) FindMitraProductCostTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraProductCostTempData, error) {
	var err error
	mitraproductcosttemp := MitraProductCostTempData{}
	err = db.Debug().Model(&MitraProductCostTempData{}).Limit(100).
		Select(`m_mitra_product_cost_temp.mitra_prod_cost_temp_id,
				m_mitra_product_temp.mitra_prod_temp_id,
				m_mitra_product_temp.mitra_prod_temp_name,
				m_product_cost_temp.prod_cost_id as prod_cost_temp_id,
				m_product_cost_temp.prod_cost_name as prod_cost_temp_name,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_sign,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_period,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_value,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_status,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_reason,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_created_by,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_created_at,
				m_mitra_product_cost.mitra_prod_cost_id,
				m_mitra_product.mitra_prod_id,
				m_mitra_product.mitra_prod_name,
				m_product_cost.prod_cost_id,
				m_product_cost.prod_cost_name,
				m_mitra_product_cost.mitra_prod_cost_sign,
				m_mitra_product_cost.mitra_prod_cost_period,
				m_mitra_product_cost.mitra_prod_cost_value,
				m_mitra_product_cost.mitra_prod_cost_status,
				m_mitra_product_cost.mitra_prod_cost_created_by,
				m_mitra_product_cost.mitra_prod_cost_updated_by,
				m_mitra_product_cost.mitra_prod_cost_deleted_by,
				m_mitra_product_cost.mitra_prod_cost_created_at,
				m_mitra_product_cost.mitra_prod_cost_updated_at,
				m_mitra_product_cost.mitra_prod_cost_deleted_at`).
		Joins("JOIN m_mitra_product_temp on m_mitra_product_cost_temp.mitra_prod_temp_id=m_mitra_product_temp.mitra_prod_temp_id").
		Joins("JOIN m_product_cost m_product_cost_temp on m_mitra_product_cost_temp.prod_cost_temp_id=m_product_cost_temp.prod_cost_id").
		Joins("JOIN m_mitra_product_cost on m_mitra_product_cost_temp.mitra_prod_id=m_mitra_product_cost.mitra_prod_id").
		Joins("JOIN m_mitra_product on m_mitra_product_cost.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_cost on m_mitra_product_cost.prod_cost_id=m_product_cost.prod_cost_id").
		Where("mitra_prod_cost_temp_id = ? AND mitra_prod_cost_temp_status = ?", pid, status).
		Find(&mitraproductcosttemp).Error
	if err != nil {
		return &MitraProductCostTempData{}, err
	}
	return &mitraproductcosttemp, nil
}

func (p *MitraProductCostTemp) FindMitraProductCostTempByMitraProductByID(db *gorm.DB, pid uint64, mitrabranchtemp uint64) (*MitraProductCostTempPend, error) {
	var err error
	mitraproductcosttemp := MitraProductCostTempPend{}
	err = db.Debug().Model(&MitraProductCostTempPend{}).Limit(100).
		Select(`m_mitra_product_cost_temp.mitra_prod_cost_temp_id,
				m_mitra_product_cost_temp.mitra_prod_cost_id,
				m_mitra_product_cost_temp.mitra_prod_temp_id,
				m_mitra_product_cost_temp.prod_cost_temp_id,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_sign,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_period,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_value,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_status,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_reason,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_created_by,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_created_at`).
		Where("mitra_prod_cost_temp_id = ? AND mitra_prod_temp_id = ?", pid, mitrabranchtemp).
		Find(&mitraproductcosttemp).Error
	if err != nil {
		return &MitraProductCostTempPend{}, err
	}
	return &mitraproductcosttemp, nil
}
func (p *MitraProductCostTemp) UpdateMitraProductCostBranchTempID(db *gorm.DB, pid uint64) (*MitraProductCostTemp, error) {
	var err error
	db = db.Debug().Model(&MitraProductCostTemp{}).Where("mitra_prod_cost_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_prod_temp_id":              p.MitraProductTempID,
			"mitra_prod_cost_temp_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraProductCostTemp{}).Where("mitra_prod_cost_temp_id = ?", pid).Error
	if err != nil {
		return &MitraProductCostTemp{}, err
	}
	return p, nil
}

func (p *MitraProductCostTemp) UpdateMitraProductCostTemp(db *gorm.DB, pid uint64) (*MitraProductCostTemp, error) {
	var err error
	db = db.Debug().Model(&MitraProductCostTemp{}).Where("mitra_prod_cost_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_prod_cost_id":          p.MitraProductCostID,
			"mitra_prod_temp_id":          p.MitraProductTempID,
			"prod_cost_temp_id":           p.ProductCostTempID,
			"mitra_prod_cost_temp_sign":   p.MitraProductCostTempSign,
			"mitra_prod_cost_temp_period": p.MitraProductCostTempPeriod,
			"mitra_prod_cost_temp_value":  p.MitraProductCostTempValue,
			"mitra_prod_cost_temp_reason": p.MitraProductCostTempReason,
			"mitra_prod_cost_temp_status": p.MitraProductCostTempStatus,
		},
	)
	err = db.Debug().Model(&MitraProductCostTemp{}).Where("mitra_prod_cost_temp_id = ?", pid).Error
	if err != nil {
		return &MitraProductCostTemp{}, err
	}
	return p, nil
}

func (p *MitraProductCostTemp) UpdateMitraProductCostTempStatus(db *gorm.DB, pid uint64) (*MitraProductCostTemp, error) {
	var err error
	db = db.Debug().Model(&MitraProductCostTemp{}).Where("mitra_prod_cost_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_prod_cost_temp_status": p.MitraProductCostTempStatus,
		},
	)
	err = db.Debug().Model(&MitraProductCostTemp{}).Where("mitra_prod_cost_temp_id = ?", pid).Error
	if err != nil {
		return &MitraProductCostTemp{}, err
	}
	return p, nil
}

func (p *MitraProductCostTemp) DeleteMitraProductCostTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraProductCostTemp{}).Where("mitra_prod_cost_temp_id = ?", pid).Take(&MitraProductCostTemp{}).Delete(&MitraProductCostTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraProductCostTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

func (p *MitraProductCostTemp) FindMitraProductCostTempPendByMitraProductTempID(db *gorm.DB, pid uint64) ([]MitraProductCostTempPendData, error) {
	var err error
	mitrabranchcoverageregency := []MitraProductCostTempPendData{}
	err = db.Debug().Model(&MitraProductCostTempPendData{}).Limit(100).
		Select(`m_mitra_product_cost_temp.mitra_prod_cost_temp_id,
				m_mitra_product_cost_temp.mitra_prod_cost_id,
				m_mitra_product_temp.mitra_prod_temp_id,
				m_mitra_product_temp.mitra_prod_temp_name,
				m_product_cost_temp.prod_cost_id as prod_cost_temp_id,
				m_product_cost_temp.prod_cost_name as prod_cost_temp_name,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_sign,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_period,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_value,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_status,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_reason,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_created_by,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_created_at`).
		Joins("JOIN m_mitra_product_temp on m_mitra_product_cost_temp.mitra_prod_temp_id=m_mitra_product_temp.mitra_prod_temp_id").
		Joins("JOIN m_product_cost m_product_cost_temp on m_mitra_product_cost_temp.prod_cost_temp_id=m_product_cost_temp.prod_cost_id").
		Where("m_mitra_product_cost_temp.mitra_prod_temp_id = ?", pid).
		Find(&mitrabranchcoverageregency).Error
	if err != nil {
		return []MitraProductCostTempPendData{}, err
	}
	return mitrabranchcoverageregency, nil
}

func (p *MitraProductCostTemp) FindMitraProductCostTempPendStatusByMitraProductTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraProductCostTempPendData, error) {
	var err error
	mitrabranchcoverageregency := []MitraProductCostTempPendData{}
	err = db.Debug().Model(&MitraProductCostTempPendData{}).Limit(100).
		Select(`m_mitra_product_cost_temp.mitra_prod_cost_temp_id,
				m_mitra_product_cost_temp.mitra_prod_cost_id,
				m_mitra_product_temp.mitra_prod_temp_id,
				m_mitra_product_temp.mitra_prod_temp_name,
				m_product_cost_temp.prod_cost_id as prod_cost_temp_id,
				m_product_cost_temp.prod_cost_name as prod_cost_temp_name,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_sign,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_period,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_value,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_status,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_reason,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_created_by,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_created_at`).
		Joins("JOIN m_mitra_product_temp on m_mitra_product_cost_temp.mitra_prod_temp_id=m_mitra_product_temp.mitra_prod_temp_id").
		Joins("JOIN m_product_cost m_product_cost_temp on m_mitra_product_cost_temp.prod_cost_temp_id=m_product_cost_temp.prod_cost_id").
		Where("m_mitra_product_cost_temp.mitra_prod_temp_id = ? AND mitra_prod_cost_temp_status = ?", pid, status).
		Find(&mitrabranchcoverageregency).Error
	if err != nil {
		return []MitraProductCostTempPendData{}, err
	}
	return mitrabranchcoverageregency, nil
}

func (p *MitraProductCostTemp) FindMitraProductCostTempByMitraProductTempID(db *gorm.DB, pid uint64) ([]MitraProductCostTempData, error) {
	var err error
	mitrabranchcoverageregency := []MitraProductCostTempData{}
	err = db.Debug().Model(&MitraProductCostTempData{}).Limit(100).
		Select(`m_mitra_product_cost_temp.mitra_prod_cost_temp_id,
				m_mitra_product_cost_temp.mitra_prod_cost_id,
				m_mitra_product_temp.mitra_prod_temp_id,
				m_mitra_product_temp.mitra_prod_temp_name,
				m_product_cost_temp.prod_cost_id as prod_cost_temp_id,
				m_product_cost_temp.prod_cost_name as prod_cost_temp_name,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_sign,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_period,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_value,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_status,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_reason,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_created_by,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_created_at`).
		Joins("JOIN m_mitra_product_temp on m_mitra_product_cost_temp.mitra_prod_temp_id=m_mitra_product_temp.mitra_prod_temp_id").
		Joins("JOIN m_product_cost m_product_cost_temp on m_mitra_product_cost_temp.prod_cost_temp_id=m_product_cost_temp.prod_cost_id").
		Where("m_mitra_product_cost_temp.mitra_prod_temp_id = ?", pid).
		Find(&mitrabranchcoverageregency).Error
	if err != nil {
		return []MitraProductCostTempData{}, err
	}
	return mitrabranchcoverageregency, nil
}

func (p *MitraProductCostTemp) FindMitraProductCostTempStatusByMitraProductTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraProductCostTempPendData, error) {
	var err error
	mitrabranchcoverageregency := []MitraProductCostTempPendData{}
	err = db.Debug().Model(&MitraProductCostTempPendData{}).Limit(100).
		Select(`m_mitra_product_cost_temp.mitra_prod_cost_temp_id,
				m_mitra_product_cost_temp.mitra_prod_cost_id,
				m_mitra_product_temp.mitra_prod_temp_id,
				m_mitra_product_temp.mitra_prod_temp_name,
				m_product_cost_temp.prod_cost_id as prod_cost_temp_id,
				m_product_cost_temp.prod_cost_name as prod_cost_temp_name,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_sign,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_period,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_value,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_status,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_reason,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_created_by,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_created_at`).
		Joins("JOIN m_mitra_product_temp on m_mitra_product_cost_temp.mitra_prod_temp_id=m_mitra_product_temp.mitra_prod_temp_id").
		Joins("JOIN m_product_cost m_product_cost_temp on m_mitra_product_cost_temp.prod_cost_temp_id=m_product_cost_temp.prod_cost_id").
		Where("m_mitra_product_cost_temp.mitra_prod_temp_id = ? AND mitra_prod_cost_temp_status = ?", pid, status).
		Find(&mitrabranchcoverageregency).Error
	if err != nil {
		return []MitraProductCostTempPendData{}, err
	}
	return mitrabranchcoverageregency, nil
}

func (p *MitraProductCostTemp) FindMitraProductCostTempDataStatusByMitraProductTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraProductCostTempData, error) {
	var err error
	mitrabranchcoverageregency := []MitraProductCostTempData{}
	err = db.Debug().Model(&MitraProductCostTempData{}).Limit(100).
		Select(`m_mitra_product_cost_temp.mitra_prod_cost_temp_id,
				m_mitra_product_cost_temp.mitra_prod_cost_id,
				m_mitra_product_temp.mitra_prod_temp_id,
				m_mitra_product_temp.mitra_prod_temp_name,
				m_product_cost_temp.prod_cost_id as prod_cost_temp_id,
				m_product_cost_temp.prod_cost_name as prod_cost_temp_name,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_sign,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_period,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_value,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_status,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_reason,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_created_by,
				m_mitra_product_cost_temp.mitra_prod_cost_temp_created_at`).
		Joins("JOIN m_mitra_product_temp on m_mitra_product_cost_temp.mitra_prod_temp_id=m_mitra_product_temp.mitra_prod_temp_id").
		Joins("JOIN m_product_cost m_product_cost_temp on m_mitra_product_cost_temp.prod_cost_temp_id=m_product_cost_temp.prod_cost_id").
		Where("m_mitra_product_cost_temp.mitra_prod_temp_id = ? AND mitra_prod_cost_temp_status = ?", pid, status).
		Find(&mitrabranchcoverageregency).Error
	if err != nil {
		return []MitraProductCostTempData{}, err
	}
	return mitrabranchcoverageregency, nil
}
