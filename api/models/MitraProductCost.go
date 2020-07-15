package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type MitraProductCost struct {
	MitraProductCostID        *uint64    `gorm:"column:mitra_prod_cost_id;primary_key;" json:"mitra_prod_cost_id"`
	MitraProductID            uint64     `gorm:"column:mitra_prod_id" json:"mitra_prod_id"`
	ProductCostID             uint64     `gorm:"column:prod_cost_id" json:"prod_cost_id"`
	MitraProductCostSign      string     `gorm:"column:mitra_prod_cost_sign;size:25" json:"mitra_prod_cost_sign"`
	MitraProductCostPeriod    string     `gorm:"column:mitra_prod_cost_period;size:255" json:"mitra_prod_cost_period"`
	MitraProductCostValue     uint64     `gorm:"column:mitra_prod_cost_value;size:255" json:"mitra_prod_cost_value"`
	MitraProductCostStatus    uint64     `gorm:"column:mitra_prod_cost_status;size:25" json:"mitra_prod_cost_status"`
	MitraProductCostCreatedBy string     `gorm:"column:mitra_prod_cost_created_by;size:125" json:"mitra_prod_cost_created_by"`
	MitraProductCostCreatedAt time.Time  `gorm:"column:mitra_prod_cost_created_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_cost_created_at"`
	MitraProductCostUpdatedBy string     `gorm:"column:mitra_prod_cost_updated_by;size:125" json:"mitra_prod_cost_updated_by"`
	MitraProductCostUpdatedAt *time.Time `gorm:"column:mitra_prod_cost_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_cost_updated_at"`
	MitraProductCostDeletedBy string     `gorm:"column:mitra_prod_cost_deleted_by;size:125" json:"mitra_prod_cost_deleted_by"`
	MitraProductCostDeletedAt *time.Time `gorm:"column:mitra_prod_cost_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_cost_deleted_at"`
}

type MitraProductCostData struct {
	MitraProductCostID        uint64     `gorm:"column:mitra_prod_cost_id;primary_key;" json:"mitra_prod_cost_id"`
	MitraProductID            uint64     `gorm:"column:mitra_prod_id" json:"mitra_prod_id"`
	MitraProductName          string     `gorm:"column:mitra_prod_name" json:"mitra_prod_name"`
	ProductCostID             uint64     `gorm:"column:prod_cost_id" json:"prod_cost_id"`
	ProductCostName           string     `gorm:"column:prod_cost_name" json:"prod_cost_name"`
	MitraProductCostSign      string     `gorm:"column:mitra_prod_cost_sign;size:25" json:"mitra_prod_cost_sign"`
	MitraProductCostPeriod    string     `gorm:"column:mitra_prod_cost_period;size:255" json:"mitra_prod_cost_period"`
	MitraProductCostValue     uint64     `gorm:"column:mitra_prod_cost_value;size:255" json:"mitra_prod_cost_value"`
	MitraProductCostStatus    uint64     `gorm:"column:mitra_prod_cost_status;size:25" json:"mitra_prod_cost_status"`
	MitraProductCostCreatedBy string     `gorm:"column:mitra_prod_cost_created_by;size:125" json:"mitra_prod_cost_created_by"`
	MitraProductCostCreatedAt time.Time  `gorm:"column:mitra_prod_cost_created_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_cost_created_at"`
	MitraProductCostUpdatedBy string     `gorm:"column:mitra_prod_cost_updated_by;size:125" json:"mitra_prod_cost_updated_by"`
	MitraProductCostUpdatedAt *time.Time `gorm:"column:mitra_prod_cost_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_cost_updated_at"`
	MitraProductCostDeletedBy string     `gorm:"column:mitra_prod_cost_deleted_by;size:125" json:"mitra_prod_cost_deleted_by"`
	MitraProductCostDeletedAt *time.Time `gorm:"column:mitra_prod_cost_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_cost_deleted_at"`
}

type ResponseMitraProductCosts struct {
	Status  int                     `json:"status"`
	Message string                  `json:"message"`
	Data    *[]MitraProductCostData `json:"data"`
}

type ResponseMitraProductCost struct {
	Status  int                   `json:"status"`
	Message string                `json:"message"`
	Data    *MitraProductCostData `json:"data"`
}

type ResponseMitraProductCostIU struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    *MitraProductCost `json:"data"`
}

type ResponseMitraProductCostCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraProductCost) TableName() string {
	return "m_mitra_product_cost"
}

func (MitraProductCostData) TableName() string {
	return "m_mitra_product_cost"
}

func (p *MitraProductCost) Prepare() {
	p.MitraProductCostID = nil
	p.MitraProductID = p.MitraProductID
	p.ProductCostID = p.ProductCostID
	p.MitraProductCostSign = html.EscapeString(strings.TrimSpace(p.MitraProductCostSign))
	p.MitraProductCostPeriod = html.EscapeString(strings.TrimSpace(p.MitraProductCostPeriod))
	p.MitraProductCostValue = p.MitraProductCostValue
	p.MitraProductCostStatus = p.MitraProductCostStatus
	p.MitraProductCostCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraProductCostCreatedBy))
	p.MitraProductCostUpdatedBy = html.EscapeString(strings.TrimSpace(p.MitraProductCostUpdatedBy))
	p.MitraProductCostDeletedBy = html.EscapeString(strings.TrimSpace(p.MitraProductCostDeletedBy))
	p.MitraProductCostCreatedAt = time.Now()
	p.MitraProductCostUpdatedAt = p.MitraProductCostUpdatedAt
	p.MitraProductCostDeletedAt = p.MitraProductCostDeletedAt
}

func (p *MitraProductCost) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.MitraProductID == 0 {
			return errors.New("Required MitraProductCost Code")
		}
		if p.ProductCostID == 0 {
			return errors.New("Required MitraProductCost Code")
		}
		if p.MitraProductCostSign == "" {
			return errors.New("Required MitraProductCost Name")
		}
		if p.MitraProductCostPeriod == "" {
			return errors.New("Required MitraProductCost Name")
		}
		if p.MitraProductCostValue == 0 {
			return errors.New("Required MitraProductCost Name")
		}
		return nil
	}
}

func (p *MitraProductCost) SaveMitraProductCost(db *gorm.DB) (*MitraProductCost, error) {
	var err error
	err = db.Debug().Model(&MitraProductCost{}).Create(&p).Error
	if err != nil {
		return &MitraProductCost{}, err
	}
	return p, nil
}

func (p *MitraProductCost) FindAllMitraProductCost(db *gorm.DB) (*[]MitraProductCostData, error) {
	var err error
	mitraproductcost := []MitraProductCostData{}
	err = db.Debug().Model(&MitraProductCostData{}).Limit(100).
		Select(`m_mitra_product_cost.mitra_prod_cost_id,
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
				m_mitra_product_cost.mitra_prod_cost_deleted_at
				`).
		Joins("JOIN m_mitra_product on m_mitra_product_cost.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_cost on m_mitra_product_cost.prod_cost_id=m_product_cost.prod_cost_id").
		Find(&mitraproductcost).Error
	if err != nil {
		return &[]MitraProductCostData{}, err
	}
	return &mitraproductcost, nil
}

func (p *MitraProductCost) FindAllMitraProductCostStatus(db *gorm.DB, status uint64) (*[]MitraProductCostData, error) {
	var err error
	mitraproductcost := []MitraProductCostData{}
	err = db.Debug().Model(&MitraProductCostData{}).Limit(100).
		Select(`m_mitra_product_cost.mitra_prod_cost_id,
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
				m_mitra_product_cost.mitra_prod_cost_deleted_at
				`).
		Joins("JOIN m_mitra_product on m_mitra_product_cost.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_cost on m_mitra_product_cost.prod_cost_id=m_product_cost.prod_cost_id").
		Where("mitra_prod_cost_status = ?", status).
		Find(&mitraproductcost).Error
	if err != nil {
		return &[]MitraProductCostData{}, err
	}
	return &mitraproductcost, nil
}

func (p *MitraProductCost) FindAllMitraProductCostByMitraProduct(db *gorm.DB, mitrabranch uint64) (*[]MitraProductCostData, error) {
	var err error
	mitraproductcost := []MitraProductCostData{}
	err = db.Debug().Model(&MitraProductCostData{}).Limit(100).
		Select(`m_mitra_product_cost.mitra_prod_cost_id,
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
				m_mitra_product_cost.mitra_prod_cost_deleted_at
				`).
		Joins("JOIN m_mitra_product on m_mitra_product_cost.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_cost on m_mitra_product_cost.prod_cost_id=m_product_cost.prod_cost_id").
		Where("m_mitra_product_cost.mitra_prod_id = ?", mitrabranch).
		Find(&mitraproductcost).Error
	if err != nil {
		return &[]MitraProductCostData{}, err
	}
	return &mitraproductcost, nil
}

func (p *MitraProductCost) FindMitraProductCostDataByID(db *gorm.DB, pid uint64) (*MitraProductCost, error) {
	var err error
	err = db.Debug().Model(&MitraProductCost{}).Where("mitra_prod_cost_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraProductCost{}, err
	}
	return p, nil
}

func (p *MitraProductCost) FindMitraProductCostByID(db *gorm.DB, pid uint64) (*MitraProductCostData, error) {
	var err error
	mitraproductcost := MitraProductCostData{}
	err = db.Debug().Model(&MitraProductCostData{}).
		Select(`m_mitra_product_cost.mitra_prod_cost_id,
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
				m_mitra_product_cost.mitra_prod_cost_deleted_at
				`).
		Joins("JOIN m_mitra_product on m_mitra_product_cost.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_cost on m_mitra_product_cost.prod_cost_id=m_product_cost.prod_cost_id").
		Where("mitra_prod_cost_id = ?", pid).
		Take(&mitraproductcost).Error
	if err != nil {
		return &MitraProductCostData{}, err
	}
	return &mitraproductcost, nil
}

func (p *MitraProductCost) FindMitraProductCostStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraProductCostData, error) {
	var err error
	mitraproductcost := MitraProductCostData{}
	err = db.Debug().Model(&MitraProductCostData{}).
		Select(`m_mitra_product_cost.mitra_prod_cost_id,
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
				m_mitra_product_cost.mitra_prod_cost_deleted_at
				`).
		Joins("JOIN m_mitra_product on m_mitra_product_cost.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_cost on m_mitra_product_cost.prod_cost_id=m_product_cost.prod_cost_id").
		Where("mitra_prod_cost_id = ? AND mitra_prod_cost_status = ?", pid, status).
		Take(&mitraproductcost).Error
	if err != nil {
		return &MitraProductCostData{}, err
	}
	return &mitraproductcost, nil
}

func (p *MitraProductCost) FindMitraProductCostByMitraProductByID(db *gorm.DB, pid uint64, mitrabranch uint64) (*MitraProductCostData, error) {
	var err error
	mitraproductcost := MitraProductCostData{}
	err = db.Debug().Model(&MitraProductCostData{}).
		Select(`m_mitra_product_cost.mitra_prod_cost_id,
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
				m_mitra_product_cost.mitra_prod_cost_deleted_at
				`).
		Joins("JOIN m_mitra_product on m_mitra_product_cost.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_cost on m_mitra_product_cost.prod_cost_id=m_product_cost.prod_cost_id").
		Where("mitra_prod_cost_id = ? AND mitra_prod_id = ?", pid, mitrabranch).
		Take(&mitraproductcost).Error
	if err != nil {
		return &MitraProductCostData{}, err
	}
	return &mitraproductcost, nil
}

func (p *MitraProductCost) UpdateMitraProductCost(db *gorm.DB, pid uint64) (*MitraProductCost, error) {
	var err error
	db = db.Debug().Model(&MitraProductCost{}).Where("mitra_prod_cost_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_prod_id":              p.MitraProductID,
			"prod_cost_id":               p.ProductCostID,
			"mitra_prod_cost_sign":       p.MitraProductCostSign,
			"mitra_prod_cost_period":     p.MitraProductCostPeriod,
			"mitra_prod_cost_value":      p.MitraProductCostValue,
			"mitra_prod_cost_status":     p.MitraProductCostStatus,
			"mitra_prod_cost_updated_by": p.MitraProductCostUpdatedBy,
			"mitra_prod_cost_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraProductCost{}).Where("mitra_prod_cost_id = ?", pid).Error
	if err != nil {
		return &MitraProductCost{}, err
	}
	return p, nil
}

func (p *MitraProductCost) UpdateMitraProductCostStatus(db *gorm.DB, pid uint64) (*MitraProductCost, error) {
	var err error
	db = db.Debug().Model(&MitraProductCost{}).Where("mitra_prod_cost_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_prod_cost_status":     p.MitraProductCostStatus,
			"mitra_prod_cost_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraProductCost{}).Where("mitra_prod_cost_id = ?", pid).Error
	if err != nil {
		return &MitraProductCost{}, err
	}
	return p, nil
}

func (p *MitraProductCost) UpdateMitraProductCostStatusOnly(db *gorm.DB, pid uint64) (*MitraProductCost, error) {
	var err error
	db = db.Debug().Model(&MitraProductCost{}).Where("mitra_prod_cost_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_prod_cost_status": p.MitraProductCostStatus,
		},
	)
	err = db.Debug().Model(&MitraProductCost{}).Where("mitra_prod_cost_id = ?", pid).Error
	if err != nil {
		return &MitraProductCost{}, err
	}
	return p, nil
}

func (p *MitraProductCost) UpdateMitraProductCostStatusDelete(db *gorm.DB, pid uint64) (*MitraProductCost, error) {
	var err error
	db = db.Debug().Model(&MitraProductCost{}).Where("mitra_prod_cost_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_prod_cost_status":     3,
			"mitra_prod_cost_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraProductCost{}).Where("mitra_prod_cost_id = ?", pid).Error
	if err != nil {
		return &MitraProductCost{}, err
	}
	return p, nil
}

func (p *MitraProductCost) DeleteMitraProductCost(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraProductCost{}).Where("mitra_prod_cost_id = ?", pid).Take(&MitraProductCost{}).Delete(&MitraProductCost{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraProductCost not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

func (p *MitraProductCost) FindMitraProductCostByMitraProductID(db *gorm.DB, mitra uint64) ([]MitraProductCostData, error) {
	var err error
	mitracoverageprovince := []MitraProductCostData{}
	err = db.Debug().Model(&MitraProductCostData{}).Limit(100).
		Select(`m_mitra_product_cost.mitra_prod_cost_id,
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
				m_mitra_product_cost.mitra_prod_cost_deleted_at
				`).
		Joins("JOIN m_mitra_product on m_mitra_product_cost.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_cost on m_mitra_product_cost.prod_cost_id=m_product_cost.prod_cost_id").
		Where("m_mitra_product_cost.mitra_prod_id = ?", mitra).
		Find(&mitracoverageprovince).Error
	if err != nil {
		return []MitraProductCostData{}, err
	}
	return mitracoverageprovince, nil
}

func (p *MitraProductCost) FindMitraProductCostStatusByMitraProductID(db *gorm.DB, mitra uint64, status uint64) ([]MitraProductCostData, error) {
	var err error
	mitracoverageprovince := []MitraProductCostData{}
	err = db.Debug().Model(&MitraProductCostData{}).Limit(100).
		Select(`m_mitra_product_cost.mitra_prod_cost_id,
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
				m_mitra_product_cost.mitra_prod_cost_deleted_at
				`).
		Joins("JOIN m_mitra_product on m_mitra_product_cost.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_cost on m_mitra_product_cost.prod_cost_id=m_product_cost.prod_cost_id").
		Where("m_mitra_product_cost.mitra_prod_id = ? AND m_mitra_product_cost.mitra_prod_cost_status = ?", mitra, status).
		Find(&mitracoverageprovince).Error
	if err != nil {
		return []MitraProductCostData{}, err
	}
	return mitracoverageprovince, nil
}

// ADDITIONAL
// ===============================================================================

func (p *MitraProductCost) FindMitraProductCostByProduct(db *gorm.DB, product uint64) ([]MitraProductCost, error) {
	var err error
	mitraproductcost := []MitraProductCost{}
	err = db.Debug().Model(&[]MitraProductCost{}).Limit(100).
		Select(`m_mitra_product_cost.mitra_prod_cost_id,
				m_mitra_product.mitra_prod_id,
				m_mitra_product.mitra_prod_name,
				m_product_cost.prod_cost_id,
				m_product_cost.prod_cost_code,
				m_product_cost.prod_cost_name,
				m_mitra_product_cost.mitra_prod_cost_period,
				m_mitra_product_cost.mitra_prod_cost_value,
				m_mitra_product_cost.mitra_prod_cost_status,
				m_mitra_product_cost.mitra_prod_cost_created_by,
				m_mitra_product_cost.mitra_prod_cost_updated_by,
				m_mitra_product_cost.mitra_prod_cost_deleted_by,
				m_mitra_product_cost.mitra_prod_cost_created_at,
				m_mitra_product_cost.mitra_prod_cost_updated_at,
				m_mitra_product_cost.mitra_prod_cost_deleted_at
				`).
		Joins("JOIN m_mitra_product on m_mitra_product_cost.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_cost on m_mitra_product_cost.prod_cost_id=m_product_cost.prod_cost_id").
		Where("m_mitra_product_cost.mitra_prod_id = ?", product).
		Find(&mitraproductcost).Error
	if err != nil {
		return []MitraProductCost{}, err
	}
	return mitraproductcost, nil
}
