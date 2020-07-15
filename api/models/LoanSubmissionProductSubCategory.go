package models

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type LoanSubmissionProductSubCategory struct {
	LoanSubmissionProductSubCategoryID        *uint64    `gorm:"column:loan_sbmssn_prod_sub_ctgry_id;primary_key;" json:"loan_sbmssn_prod_sub_ctgry_id"`
	LoanSubmissionProductID                   uint64     `gorm:"column:loan_sbmssn_prod_id" json:"loan_sbmssn_prod_id"`
	ProductSubCategoryID                      uint64     `gorm:"column:prod_sub_ctgry_id" json:"prod_sub_ctgry_id"`
	LoanSubmissionProductSubCategoryStatus    uint64     `gorm:"column:loan_sbmssn_prod_sub_ctgry_status;size:2" json:"loan_sbmssn_prod_sub_ctgry_status"`
	LoanSubmissionProductSubCategoryCreatedAt time.Time  `gorm:"column:loan_sbmssn_prod_sub_ctgry_created_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_prod_sub_ctgry_created_at"`
	LoanSubmissionProductSubCategoryCreatedBy string     `gorm:"column:loan_sbmssn_prod_sub_ctgry_created_by;size:125" json:"loan_sbmssn_prod_sub_ctgry_created_by"`
	LoanSubmissionProductSubCategoryUpdatedAt *time.Time `gorm:"column:loan_sbmssn_prod_sub_ctgry_updated_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_prod_sub_ctgry_updated_at"`
	LoanSubmissionProductSubCategoryUpdatedBy string     `gorm:"column:loan_sbmssn_prod_sub_ctgry_updated_by;size:125" json:"loan_sbmssn_prod_sub_ctgry_updated_by"`
	LoanSubmissionProductSubCategoryDeletedAt *time.Time `gorm:"column:loan_sbmssn_prod_sub_ctgry_deleted_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_prod_sub_ctgry_deleted_at"`
	LoanSubmissionProductSubCategoryDeletedBy string     `gorm:"column:loan_sbmssn_prod_sub_ctgry_deleted_by;size:125" json:"loan_sbmssn_prod_sub_ctgry_deleted_by"`
}

type LoanSubmissionProductSubCategoryData struct {
	LoanSubmissionProductSubCategoryID        uint64     `gorm:"column:loan_sbmssn_prod_sub_ctgry_id" json:"loan_sbmssn_prod_sub_ctgry_id"`
	LoanSubmissionProductID                   uint64     `gorm:"column:loan_sbmssn_prod_id" json:"loan_sbmssn_prod_id"`
	LoanSubmissionProductName                 string     `gorm:"column:loan_sbmssn_prod_name" json:"loan_sbmssn_prod_name"`
	ProductSubCategoryID                      uint64     `gorm:"column:prod_sub_ctgry_id" json:"prod_sub_ctgry_id"`
	ProductSubCategoryCode                    string     `gorm:"column:prod_sub_ctgry_code" json:"prod_sub_ctgry_code"`
	ProductSubCategoryName                    string     `gorm:"column:prod_sub_ctgry_name" json:"prod_sub_ctgry_name"`
	LoanSubmissionProductSubCategoryStatus    uint64     `gorm:"column:loan_sbmssn_prod_sub_ctgry_status" json:"loan_sbmssn_prod_sub_ctgry_status"`
	LoanSubmissionProductSubCategoryCreatedAt time.Time  `gorm:"column:loan_sbmssn_prod_sub_ctgry_created_at" json:"loan_sbmssn_prod_sub_ctgry_created_at"`
	LoanSubmissionProductSubCategoryCreatedBy string     `gorm:"column:loan_sbmssn_prod_sub_ctgry_created_by" json:"loan_sbmssn_prod_sub_ctgry_created_by"`
	LoanSubmissionProductSubCategoryUpdatedAt *time.Time `gorm:"column:loan_sbmssn_prod_sub_ctgry_updated_at" json:"loan_sbmssn_prod_sub_ctgry_updated_at"`
	LoanSubmissionProductSubCategoryUpdatedBy string     `gorm:"column:loan_sbmssn_prod_sub_ctgry_updated_by" json:"loan_sbmssn_prod_sub_ctgry_updated_by"`
	LoanSubmissionProductSubCategoryDeletedAt *time.Time `gorm:"column:loan_sbmssn_prod_sub_ctgry_deleted_at" json:"loan_sbmssn_prod_sub_ctgry_deleted_at"`
	LoanSubmissionProductSubCategoryDeletedBy string     `gorm:"column:loan_sbmssn_prod_sub_ctgry_deleted_by" json:"loan_sbmssn_prod_sub_ctgry_deleted_by"`
}

type ResponseLoanSubmissionProductSubCategorys struct {
	Status  int                                     `json:"status"`
	Message string                                  `json:"message"`
	Data    *[]LoanSubmissionProductSubCategoryData `json:"data"`
}

type ResponseLoanSubmissionProductSubCategory struct {
	Status  int                                   `json:"status"`
	Message string                                `json:"message"`
	Data    *LoanSubmissionProductSubCategoryData `json:"data"`
}

type ResponseLoanSubmissionProductSubCategoryIU struct {
	Status  int                               `json:"status"`
	Message string                            `json:"message"`
	Data    *LoanSubmissionProductSubCategory `json:"data"`
}

type ResponseLoanSubmissionProductSubCategoryDel struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (LoanSubmissionProductSubCategory) TableName() string {
	return "t_loan_submission_product_sub_category"
}

func (LoanSubmissionProductSubCategoryData) TableName() string {
	return "t_loan_submission_product_sub_category"
}

func (p *LoanSubmissionProductSubCategory) Prepare() {
	p.LoanSubmissionProductSubCategoryID = nil
	p.LoanSubmissionProductID = p.LoanSubmissionProductID
	p.ProductSubCategoryID = p.ProductSubCategoryID
	p.LoanSubmissionProductSubCategoryStatus = p.LoanSubmissionProductSubCategoryStatus
	p.LoanSubmissionProductSubCategoryCreatedAt = time.Now()
	p.LoanSubmissionProductSubCategoryCreatedBy = p.LoanSubmissionProductSubCategoryCreatedBy
	p.LoanSubmissionProductSubCategoryUpdatedAt = p.LoanSubmissionProductSubCategoryUpdatedAt
	p.LoanSubmissionProductSubCategoryUpdatedBy = p.LoanSubmissionProductSubCategoryUpdatedBy
	p.LoanSubmissionProductSubCategoryDeletedAt = p.LoanSubmissionProductSubCategoryDeletedAt
	p.LoanSubmissionProductSubCategoryDeletedBy = p.LoanSubmissionProductSubCategoryDeletedBy
}

func (p *LoanSubmissionProductSubCategory) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.LoanSubmissionProductID == 0 {
			return errors.New("Required LoanSubmissionProductSubCategory Code")
		}
		if p.ProductSubCategoryID == 0 {
			return errors.New("Required LoanSubmissionProductSubCategory Name")
		}
		return nil
	}
}

func (p *LoanSubmissionProductSubCategory) SaveLoanSubmissionProductSubCategory(db *gorm.DB) (*LoanSubmissionProductSubCategory, error) {
	var err error
	err = db.Debug().Model(&LoanSubmissionProductSubCategory{}).Create(&p).Error
	if err != nil {
		return &LoanSubmissionProductSubCategory{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProductSubCategory) FindAllLoanSubmissionProductSubCategory(db *gorm.DB) (*[]LoanSubmissionProductSubCategoryData, error) {
	var err error
	mitraproductsubcategory := []LoanSubmissionProductSubCategoryData{}
	err = db.Debug().Model(&LoanSubmissionProductSubCategoryData{}).Limit(100).
		Select(`t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_id,
				t_loan_produsbmssn_ct.loan_sbmssn_prod_id,
				t_loan_produsbmssn_ct.loan_sbmssn_prod_name,
				m_product_sub_category.prod_sub_ctgry_id,
				m_product_sub_category.prod_sub_ctgry_code,
				m_product_sub_category.prod_sub_ctgry_name,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_status,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_created_at,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_created_by,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_updated_at,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_updated_by,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_deleted_at,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_deleted_by`).
		Joins("JOIN t_loan_produsbmssn_ct on t_loan_submission_product_sub_category.loan_sbmssn_prod_id=t_loan_produsbmssn_ct.loan_sbmssn_prod_id").
		Joins("JOIN m_product_sub_category on t_loan_submission_product_sub_category.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Find(&mitraproductsubcategory).Error
	if err != nil {
		return &[]LoanSubmissionProductSubCategoryData{}, err
	}
	return &mitraproductsubcategory, nil
}

func (p *LoanSubmissionProductSubCategory) FindAllLoanSubmissionProductSubCategoryStatus(db *gorm.DB, status uint64) (*[]LoanSubmissionProductSubCategoryData, error) {
	var err error
	mitraproductsubcategory := []LoanSubmissionProductSubCategoryData{}
	err = db.Debug().Model(&LoanSubmissionProductSubCategoryData{}).Limit(100).
		Select(`t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_id,
				t_loan_produsbmssn_ct.loan_sbmssn_prod_id,
				t_loan_produsbmssn_ct.loan_sbmssn_prod_name,
				m_product_sub_category.prod_sub_ctgry_id,
				m_product_sub_category.prod_sub_ctgry_code,
				m_product_sub_category.prod_sub_ctgry_name,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_status,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_created_at,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_created_by,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_updated_at,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_updated_by,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_deleted_at,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_deleted_by`).
		Joins("JOIN t_loan_produsbmssn_ct on t_loan_submission_product_sub_category.loan_sbmssn_prod_id=t_loan_produsbmssn_ct.loan_sbmssn_prod_id").
		Joins("JOIN m_product_sub_category on t_loan_submission_product_sub_category.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Where("loan_sbmssn_prod_sub_ctgry_status = ?", status).
		Find(&mitraproductsubcategory).Error
	if err != nil {
		return &[]LoanSubmissionProductSubCategoryData{}, err
	}
	return &mitraproductsubcategory, nil
}

func (p *LoanSubmissionProductSubCategory) FindAllLoanSubmissionProductSubCategoryByProduct(db *gorm.DB, product uint64) (*[]LoanSubmissionProductSubCategoryData, error) {
	var err error
	mitraproductsubcategory := []LoanSubmissionProductSubCategoryData{}
	err = db.Debug().Model(&LoanSubmissionProductSubCategoryData{}).Limit(100).
		Select(`t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_id,
				t_loan_produsbmssn_ct.loan_sbmssn_prod_id,
				t_loan_produsbmssn_ct.loan_sbmssn_prod_name,
				m_product_sub_category.prod_sub_ctgry_id,
				m_product_sub_category.prod_sub_ctgry_code,
				m_product_sub_category.prod_sub_ctgry_name,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_status,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_created_at,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_created_by,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_updated_at,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_updated_by,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_deleted_at,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_deleted_by`).
		Joins("JOIN t_loan_produsbmssn_ct on t_loan_submission_product_sub_category.loan_sbmssn_prod_id=t_loan_produsbmssn_ct.loan_sbmssn_prod_id").
		Joins("JOIN m_product_sub_category on t_loan_submission_product_sub_category.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Where("loan_sbmssn_prod_id = ?", product).
		Find(&mitraproductsubcategory).Error
	if err != nil {
		return &[]LoanSubmissionProductSubCategoryData{}, err
	}
	return &mitraproductsubcategory, nil
}

func (p *LoanSubmissionProductSubCategory) FindLoanSubmissionProductSubCategoryByID(db *gorm.DB, pid uint64) (*LoanSubmissionProductSubCategoryData, error) {
	var err error
	mitraproductcost := LoanSubmissionProductSubCategoryData{}
	err = db.Debug().Model(&LoanSubmissionProductSubCategoryData{}).Limit(100).
		Select(`t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_id,
				t_loan_produsbmssn_ct.loan_sbmssn_prod_id,
				t_loan_produsbmssn_ct.loan_sbmssn_prod_name,
				m_product_sub_category.prod_sub_ctgry_id,
				m_product_sub_category.prod_sub_ctgry_code,
				m_product_sub_category.prod_sub_ctgry_name,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_status,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_created_at,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_created_by,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_updated_at,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_updated_by,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_deleted_at,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_deleted_by`).
		Joins("JOIN t_loan_produsbmssn_ct on t_loan_submission_product_sub_category.loan_sbmssn_prod_id=t_loan_produsbmssn_ct.loan_sbmssn_prod_id").
		Joins("JOIN m_product_sub_category on t_loan_submission_product_sub_category.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Where("loan_sbmssn_prod_sub_ctgry_id = ?", pid).
		Find(&mitraproductcost).Error
	if err != nil {
		return &LoanSubmissionProductSubCategoryData{}, err
	}
	return &mitraproductcost, nil
}

func (p *LoanSubmissionProductSubCategory) FindLoanSubmissionProductSubCategoryStatusByID(db *gorm.DB, pid uint64, status uint64) (*LoanSubmissionProductSubCategoryData, error) {
	var err error
	mitraproductcost := LoanSubmissionProductSubCategoryData{}
	err = db.Debug().Model(&LoanSubmissionProductSubCategoryData{}).Limit(100).
		Select(`t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_id,
				t_loan_produsbmssn_ct.loan_sbmssn_prod_id,
				t_loan_produsbmssn_ct.loan_sbmssn_prod_name,
				m_product_sub_category.prod_sub_ctgry_id,
				m_product_sub_category.prod_sub_ctgry_code,
				m_product_sub_category.prod_sub_ctgry_name,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_status,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_created_at,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_created_by,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_updated_at,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_updated_by,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_deleted_at,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_deleted_by`).
		Joins("JOIN t_loan_produsbmssn_ct on t_loan_submission_product_sub_category.loan_sbmssn_prod_id=t_loan_produsbmssn_ct.loan_sbmssn_prod_id").
		Joins("JOIN m_product_sub_category on t_loan_submission_product_sub_category.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Where("loan_sbmssn_prod_sub_ctgry_id = ? AND loan_sbmssn_prod_sub_ctgry_status = ?", pid, status).
		Find(&mitraproductcost).Error
	if err != nil {
		return &LoanSubmissionProductSubCategoryData{}, err
	}
	return &mitraproductcost, nil
}

func (p *LoanSubmissionProductSubCategory) FindLoanSubmissionProductSubCategoryByProductByID(db *gorm.DB, pid uint64, product uint64) (*LoanSubmissionProductSubCategoryData, error) {
	var err error
	mitraproductcost := LoanSubmissionProductSubCategoryData{}
	err = db.Debug().Model(&LoanSubmissionProductSubCategoryData{}).Limit(100).
		Select(`t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_id,
				t_loan_produsbmssn_ct.loan_sbmssn_prod_id,
				t_loan_produsbmssn_ct.loan_sbmssn_prod_name,
				m_product_sub_category.prod_sub_ctgry_id,
				m_product_sub_category.prod_sub_ctgry_code,
				m_product_sub_category.prod_sub_ctgry_name,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_status,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_created_at,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_created_by,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_updated_at,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_updated_by,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_deleted_at,
				t_loan_submission_product_sub_category.loan_sbmssn_prod_sub_ctgry_deleted_by`).
		Joins("JOIN t_loan_produsbmssn_ct on t_loan_submission_product_sub_category.loan_sbmssn_prod_id=t_loan_produsbmssn_ct.loan_sbmssn_prod_id").
		Joins("JOIN m_product_sub_category on t_loan_submission_product_sub_category.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Where("loan_sbmssn_prod_sub_ctgry_id = ? AND loan_sbmssn_prod_id = ?", pid, product).
		Find(&mitraproductcost).Error
	if err != nil {
		return &LoanSubmissionProductSubCategoryData{}, err
	}
	return &mitraproductcost, nil
}

func (p *LoanSubmissionProductSubCategory) UpdateLoanSubmissionProductSubCategory(db *gorm.DB, pid uint64) (*LoanSubmissionProductSubCategory, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionProductSubCategory{}).Where("loan_sbmssn_prod_sub_ctgry_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_prod_id":                   p.LoanSubmissionProductID,
			"prod_sub_ctgry_id":                     p.ProductSubCategoryID,
			"loan_sbmssn_prod_sub_ctgry_updated_by": p.LoanSubmissionProductSubCategoryUpdatedBy,
			"loan_sbmssn_prod_sub_ctgry_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionProductSubCategory{}).Where("loan_sbmssn_prod_sub_ctgry_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionProductSubCategory{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProductSubCategory) UpdateLoanSubmissionProductSubCategoryStatus(db *gorm.DB, pid uint64) (*LoanSubmissionProductSubCategory, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionProductSubCategory{}).Where("loan_sbmssn_prod_sub_ctgry_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_prod_sub_ctgry_status":     p.LoanSubmissionProductSubCategoryStatus,
			"loan_sbmssn_prod_sub_ctgry_updated_by": p.LoanSubmissionProductSubCategoryUpdatedBy,
			"loan_sbmssn_prod_sub_ctgry_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionProductSubCategory{}).Where("loan_sbmssn_prod_sub_ctgry_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionProductSubCategory{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProductSubCategory) UpdateLoanSubmissionProductSubCategoryStatusDelete(db *gorm.DB, pid uint64) (*LoanSubmissionProductSubCategory, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionProductSubCategory{}).Where("loan_sbmssn_prod_sub_ctgry_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_prod_sub_ctgry_status":     3,
			"loan_sbmssn_prod_sub_ctgry_deleted_by": p.LoanSubmissionProductSubCategoryDeletedBy,
			"loan_sbmssn_prod_sub_ctgry_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionProductSubCategory{}).Where("loan_sbmssn_prod_sub_ctgry_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionProductSubCategory{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProductSubCategory) DeleteLoanSubmissionProductSubCategory(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&LoanSubmissionProductSubCategory{}).Where("loan_sbmssn_prod_sub_ctgry_id = ?", pid).Take(&LoanSubmissionProductSubCategory{}).Delete(&LoanSubmissionProductSubCategory{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("LoanSubmissionProductSubCategory not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
