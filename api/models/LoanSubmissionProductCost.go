package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type LoanSubmissionProductCost struct {
	LoanSubmissionProductCostID        *uint64    `gorm:"column:loan_sbmssn_prod_cost_id;primary_key;" json:"loan_sbmssn_prod_cost_id"`
	LoanSubmissionProductID            uint64     `gorm:"column:loan_sbmssn_prod_id" json:"loan_sbmssn_prod_id"`
	ProductCostID                      uint64     `gorm:"column:prod_cost_id" json:"prod_cost_id"`
	LoanSubmissionProductCostPeriod    string     `gorm:"column:loan_sbmssn_prod_cost_period;size:125" json:"loan_sbmssn_prod_cost_period"`
	LoanSubmissionProductCostValue     uint64     `gorm:"column:loan_sbmssn_prod_cost_value" json:"loan_sbmssn_prod_cost_value"`
	LoanSubmissionProductCostStatus    uint64     `gorm:"column:loan_sbmssn_prod_cost_status;size:2" json:"loan_sbmssn_prod_cost_status"`
	LoanSubmissionProductCostCreatedAt time.Time  `gorm:"column:loan_sbmssn_prod_cost_created_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_prod_cost_created_at"`
	LoanSubmissionProductCostCreatedBy string     `gorm:"column:loan_sbmssn_prod_cost_created_by;size:125" json:"loan_sbmssn_prod_cost_created_by"`
	LoanSubmissionProductCostUpdatedAt *time.Time `gorm:"column:loan_sbmssn_prod_cost_updated_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_prod_cost_updated_at"`
	LoanSubmissionProductCostUpdatedBy string     `gorm:"column:loan_sbmssn_prod_cost_updated_by;size:125" json:"loan_sbmssn_prod_cost_updated_by"`
	LoanSubmissionProductCostDeletedAt *time.Time `gorm:"column:loan_sbmssn_prod_cost_deleted_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_prod_cost_deleted_at"`
	LoanSubmissionProductCostDeletedBy string     `gorm:"column:loan_sbmssn_prod_cost_deleted_by;size:125" json:"loan_sbmssn_prod_cost_deleted_by"`
}

type LoanSubmissionProductCostData struct {
	LoanSubmissionProductCostID        uint64     `gorm:"column:loan_sbmssn_prod_cost_id" json:"loan_sbmssn_prod_cost_id"`
	LoanSubmissionProductID            uint64     `gorm:"column:loan_sbmssn_prod_id" json:"loan_sbmssn_prod_id"`
	LoanSubmissionProductName          string     `gorm:"column:loan_sbmssn_prod_name" json:"loan_sbmssn_prod_name"`
	ProductCostID                      uint64     `gorm:"column:prod_cost_id" json:"prod_cost_id"`
	ProductCostCode                    string     `gorm:"column:prod_cost_code" json:"prod_cost_code"`
	ProductCostName                    string     `gorm:"column:prod_cost_name" json:"prod_cost_name"`
	ProductCostSign                    string     `gorm:"column:prod_cost_sign" json:"prod_cost_sign"`
	LoanSubmissionProductCostPeriod    string     `gorm:"column:loan_sbmssn_prod_cost_period" json:"loan_sbmssn_prod_cost_period"`
	LoanSubmissionProductCostValue     uint64     `gorm:"column:loan_sbmssn_prod_cost_value" json:"loan_sbmssn_prod_cost_value"`
	LoanSubmissionProductCostStatus    uint64     `gorm:"column:loan_sbmssn_prod_cost_status" json:"loan_sbmssn_prod_cost_status"`
	LoanSubmissionProductCostCreatedAt time.Time  `gorm:"column:loan_sbmssn_prod_cost_created_at" json:"loan_sbmssn_prod_cost_created_at"`
	LoanSubmissionProductCostCreatedBy string     `gorm:"column:loan_sbmssn_prod_cost_created_by" json:"loan_sbmssn_prod_cost_created_by"`
	LoanSubmissionProductCostUpdatedAt *time.Time `gorm:"column:loan_sbmssn_prod_cost_updated_at" json:"loan_sbmssn_prod_cost_updated_at"`
	LoanSubmissionProductCostUpdatedBy string     `gorm:"column:loan_sbmssn_prod_cost_updated_by" json:"loan_sbmssn_prod_cost_updated_by"`
	LoanSubmissionProductCostDeletedAt *time.Time `gorm:"column:loan_sbmssn_prod_cost_deleted_at" json:"loan_sbmssn_prod_cost_deleted_at"`
	LoanSubmissionProductCostDeletedBy string     `gorm:"column:loan_sbmssn_prod_cost_deleted_by" json:"loan_sbmssn_prod_cost_deleted_by"`
}

type ResponseLoanSubmissionProductCosts struct {
	Status  int                              `json:"status"`
	Message string                           `json:"message"`
	Data    *[]LoanSubmissionProductCostData `json:"data"`
}

type ResponseLoanSubmissionProductCost struct {
	Status  int                            `json:"status"`
	Message string                         `json:"message"`
	Data    *LoanSubmissionProductCostData `json:"data"`
}

type ResponseLoanSubmissionProductCostCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (LoanSubmissionProductCost) TableName() string {
	return "t_loan_submission_product_cost"
}

func (LoanSubmissionProductCostData) TableName() string {
	return "t_loan_submission_product_cost"
}

func (p *LoanSubmissionProductCost) Prepare() {
	p.LoanSubmissionProductCostID = nil
	p.LoanSubmissionProductID = p.LoanSubmissionProductID
	p.ProductCostID = p.ProductCostID
	p.LoanSubmissionProductCostPeriod = html.EscapeString(strings.TrimSpace(p.LoanSubmissionProductCostPeriod))
	p.LoanSubmissionProductCostValue = p.LoanSubmissionProductCostValue
	p.LoanSubmissionProductCostStatus = p.LoanSubmissionProductCostStatus
	p.LoanSubmissionProductCostCreatedAt = time.Now()
	p.LoanSubmissionProductCostCreatedBy = p.LoanSubmissionProductCostCreatedBy
	p.LoanSubmissionProductCostUpdatedAt = p.LoanSubmissionProductCostUpdatedAt
	p.LoanSubmissionProductCostUpdatedBy = p.LoanSubmissionProductCostUpdatedBy
	p.LoanSubmissionProductCostDeletedAt = p.LoanSubmissionProductCostDeletedAt
	p.LoanSubmissionProductCostDeletedBy = p.LoanSubmissionProductCostDeletedBy
}

func (p *LoanSubmissionProductCost) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.LoanSubmissionProductID == 0 {
			return errors.New("Required LoanSubmissionProductCost Code")
		}
		if p.ProductCostID == 0 {
			return errors.New("Required LoanSubmissionProductCost ID")
		}
		if p.LoanSubmissionProductCostPeriod == "" {
			return errors.New("Required LoanSubmissionProductCost Period")
		}
		if p.LoanSubmissionProductCostValue == 0 {
			return errors.New("Required LoanSubmissionProductCost Value")
		}
		return nil
	}
}

func (p *LoanSubmissionProductCost) SaveLoanSubmissionProductCost(db *gorm.DB) (*LoanSubmissionProductCost, error) {
	var err error
	err = db.Debug().Model(&LoanSubmissionProductCost{}).Create(&p).Error
	if err != nil {
		return &LoanSubmissionProductCost{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProductCost) FindAllLoanSubmissionProductCost(db *gorm.DB) (*[]LoanSubmissionProductCostData, error) {
	var err error
	mitraproductcost := []LoanSubmissionProductCostData{}
	err = db.Debug().Model(&LoanSubmissionProductCostData{}).Limit(100).
		Select(`t_loan_submission_product_cost.loan_sbmssn_prod_cost_id,
				t_loan_submission_product.loan_sbmssn_prod_id,
				t_loan_submission_product.loan_sbmssn_prod_name,
				m_product_cost.prod_cost_id,
				m_product_cost.prod_cost_code,
				m_product_cost.prod_cost_name,
				m_product_cost.prod_cost_sign,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_period,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_value,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_status,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_created_at,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_created_by,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_updated_at,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_updated_by,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_deleted_at,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_deleted_by
				`).
		Joins("JOIN t_loan_submission_product on t_loan_submission_product_cost.loan_sbmssn_prod_id=t_loan_submission_product.loan_sbmssn_prod_id").
		Joins("JOIN m_product_cost on t_loan_submission_product_cost.prod_cost_id=m_product_cost.prod_cost_id").
		Find(&mitraproductcost).Error
	if err != nil {
		return &[]LoanSubmissionProductCostData{}, err
	}
	return &mitraproductcost, nil
}

func (p *LoanSubmissionProductCost) FindAllLoanSubmissionProductCostStatus(db *gorm.DB, status uint64) (*[]LoanSubmissionProductCostData, error) {
	var err error
	mitraproductcost := []LoanSubmissionProductCostData{}
	err = db.Debug().Model(&LoanSubmissionProductCostData{}).Limit(100).
		Select(`t_loan_submission_product_cost.loan_sbmssn_prod_cost_id,
				t_loan_submission_product.loan_sbmssn_prod_id,
				t_loan_submission_product.loan_sbmssn_prod_name,
				m_product_cost.prod_cost_id,
				m_product_cost.prod_cost_code,
				m_product_cost.prod_cost_name,
				m_product_cost.prod_cost_sign,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_period,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_value,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_status,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_created_at,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_created_by,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_updated_at,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_updated_by,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_deleted_at,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_deleted_by
				`).
		Joins("JOIN t_loan_submission_product on t_loan_submission_product_cost.loan_sbmssn_prod_id=t_loan_submission_product.loan_sbmssn_prod_id").
		Joins("JOIN m_product_cost on t_loan_submission_product_cost.prod_cost_id=m_product_cost.prod_cost_id").
		Where("loan_sbmssn_prod_cost_status = ?", status).
		Find(&mitraproductcost).Error
	if err != nil {
		return &[]LoanSubmissionProductCostData{}, err
	}
	return &mitraproductcost, nil
}

func (p *LoanSubmissionProductCost) FindAllLoanSubmissionProductCostByProduct(db *gorm.DB, product uint64) (*[]LoanSubmissionProductCostData, error) {
	var err error
	mitraproductcost := []LoanSubmissionProductCostData{}
	err = db.Debug().Model(&LoanSubmissionProductCostData{}).Limit(100).
		Select(`t_loan_submission_product_cost.loan_sbmssn_prod_cost_id,
				t_loan_submission_product.loan_sbmssn_prod_id,
				t_loan_submission_product.loan_sbmssn_prod_name,
				m_product_cost.prod_cost_id,
				m_product_cost.prod_cost_code,
				m_product_cost.prod_cost_name,
				m_product_cost.prod_cost_sign,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_period,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_value,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_status,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_created_at,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_created_by,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_updated_at,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_updated_by,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_deleted_at,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_deleted_by
				`).
		Joins("JOIN t_loan_submission_product on t_loan_submission_product_cost.loan_sbmssn_prod_id=t_loan_submission_product.loan_sbmssn_prod_id").
		Joins("JOIN m_product_cost on t_loan_submission_product_cost.prod_cost_id=m_product_cost.prod_cost_id").
		Where("loan_sbmssn_prod_id = ?", product).
		Find(&mitraproductcost).Error
	if err != nil {
		return &[]LoanSubmissionProductCostData{}, err
	}
	return &mitraproductcost, nil
}

func (p *LoanSubmissionProductCost) FindLoanSubmissionProductCostByID(db *gorm.DB, pid uint64) (*LoanSubmissionProductCostData, error) {
	var err error
	mitraproductcost := LoanSubmissionProductCostData{}
	err = db.Debug().Model(&LoanSubmissionProductCostData{}).Limit(100).
		Select(`t_loan_submission_product_cost.loan_sbmssn_prod_cost_id,
				t_loan_submission_product.loan_sbmssn_prod_id,
				t_loan_submission_product.loan_sbmssn_prod_name,
				m_product_cost.prod_cost_id,
				m_product_cost.prod_cost_code,
				m_product_cost.prod_cost_name,
				m_product_cost.prod_cost_sign,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_period,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_value,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_status,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_created_at,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_created_by,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_updated_at,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_updated_by,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_deleted_at,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_deleted_by
				`).
		Joins("JOIN t_loan_submission_product on t_loan_submission_product_cost.loan_sbmssn_prod_id=t_loan_submission_product.loan_sbmssn_prod_id").
		Joins("JOIN m_product_cost on t_loan_submission_product_cost.prod_cost_id=m_product_cost.prod_cost_id").
		Where("loan_sbmssn_prod_cost_id = ?", pid).
		Find(&mitraproductcost).Error
	if err != nil {
		return &LoanSubmissionProductCostData{}, err
	}
	return &mitraproductcost, nil
}

func (p *LoanSubmissionProductCost) FindLoanSubmissionProductCostStatusByID(db *gorm.DB, pid uint64, status uint64) (*LoanSubmissionProductCostData, error) {
	var err error
	mitraproductcost := LoanSubmissionProductCostData{}
	err = db.Debug().Model(&LoanSubmissionProductCostData{}).Limit(100).
		Select(`t_loan_submission_product_cost.loan_sbmssn_prod_cost_id,
				t_loan_submission_product.loan_sbmssn_prod_id,
				t_loan_submission_product.loan_sbmssn_prod_name,
				m_product_cost.prod_cost_id,
				m_product_cost.prod_cost_code,
				m_product_cost.prod_cost_name,
				m_product_cost.prod_cost_sign,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_period,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_value,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_status,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_created_at,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_created_by,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_updated_at,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_updated_by,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_deleted_at,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_deleted_by
			`).
		Joins("JOIN t_loan_submission_product on t_loan_submission_product_cost.loan_sbmssn_prod_id=t_loan_submission_product.loan_sbmssn_prod_id").
		Joins("JOIN m_product_cost on t_loan_submission_product_cost.prod_cost_id=m_product_cost.prod_cost_id").
		Where("loan_sbmssn_prod_cost_id = ? AND loan_sbmssn_prod_cost_status = ?", pid, status).
		Find(&mitraproductcost).Error
	if err != nil {
		return &LoanSubmissionProductCostData{}, err
	}
	return &mitraproductcost, nil
}

func (p *LoanSubmissionProductCost) FindLoanSubmissionProductCostByProductByID(db *gorm.DB, pid uint64, product uint64) (*LoanSubmissionProductCostData, error) {
	var err error
	mitraproductcost := LoanSubmissionProductCostData{}
	err = db.Debug().Model(&LoanSubmissionProductCostData{}).Limit(100).
		Select(`t_loan_submission_product_cost.loan_sbmssn_prod_cost_id,
				t_loan_submission_product.loan_sbmssn_prod_id,
				t_loan_submission_product.loan_sbmssn_prod_name,
				m_product_cost.prod_cost_id,
				m_product_cost.prod_cost_code,
				m_product_cost.prod_cost_name,
				m_product_cost.prod_cost_sign,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_period,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_value,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_status,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_created_at,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_created_by,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_updated_at,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_updated_by,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_deleted_at,
				t_loan_submission_product_cost.loan_sbmssn_prod_cost_deleted_by
			`).
		Joins("JOIN t_loan_submission_product on t_loan_submission_product_cost.loan_sbmssn_prod_id=t_loan_submission_product.loan_sbmssn_prod_id").
		Joins("JOIN m_product_cost on t_loan_submission_product_cost.prod_cost_id=m_product_cost.prod_cost_id").
		Where("loan_sbmssn_prod_cost_id = ? AND loan_sbmssn_prod_id = ?", pid, product).
		Find(&mitraproductcost).Error
	if err != nil {
		return &LoanSubmissionProductCostData{}, err
	}
	return &mitraproductcost, nil
}

func (p *LoanSubmissionProductCost) UpdateLoanSubmissionProductCost(db *gorm.DB, pid uint64) (*LoanSubmissionProductCost, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionProductCost{}).Where("loan_sbmssn_prod_cost_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_prod_id":              p.LoanSubmissionProductID,
			"prod_cost_id":                     p.ProductCostID,
			"loan_sbmssn_prod_cost_period":     p.LoanSubmissionProductCostPeriod,
			"loan_sbmssn_prod_cost_value":      p.LoanSubmissionProductCostValue,
			"loan_sbmssn_prod_cost_status":     p.LoanSubmissionProductCostStatus,
			"loan_sbmssn_prod_cost_updated_by": p.LoanSubmissionProductCostUpdatedBy,
			"loan_sbmssn_prod_cost_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionProductCost{}).Where("loan_sbmssn_prod_cost_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionProductCost{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProductCost) UpdateLoanSubmissionProductCostStatus(db *gorm.DB, pid uint64) (*LoanSubmissionProductCost, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionProductCost{}).Where("loan_sbmssn_prod_cost_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_prod_cost_status":     p.LoanSubmissionProductCostStatus,
			"loan_sbmssn_prod_cost_updated_by": p.LoanSubmissionProductCostUpdatedBy,
			"loan_sbmssn_prod_cost_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionProductCost{}).Where("loan_sbmssn_prod_cost_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionProductCost{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProductCost) UpdateLoanSubmissionProductCostStatusDelete(db *gorm.DB, pid uint64) (*LoanSubmissionProductCost, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionProductCost{}).Where("loan_sbmssn_prod_cost_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_prod_cost_status":     3,
			"loan_sbmssn_prod_cost_deleted_by": p.LoanSubmissionProductCostDeletedBy,
			"loan_sbmssn_prod_cost_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionProductCost{}).Where("loan_sbmssn_prod_cost_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionProductCost{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProductCost) DeleteLoanSubmissionProductCost(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&LoanSubmissionProductCost{}).Where("loan_sbmssn_prod_cost_id = ?", pid).Take(&LoanSubmissionProductCost{}).Delete(&LoanSubmissionProductCost{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("LoanSubmissionProductCost not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
