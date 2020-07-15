package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type LoanSubmissionProduct struct {
	LoanSubmissionProductID                  uint64     `gorm:"column:loan_sbmssn_prod_id;primary_key;" json:"loan_sbmssn_prod_id"`
	LoanSubmissionID                         uint64     `gorm:"column:loan_sbmssn_id" json:"loan_sbmssn_id"`
	MitraProductID                           uint64     `gorm:"column:mitra_prod_id" json:"mitra_prod_id"`
	LoanSubmissionProductName                string     `gorm:"column:loan_sbmssn_prod_name;size:125" json:"loan_sbmssn_prod_name"`
	LoanSubmissionProductLimitMin            uint64     `gorm:"column:loan_sbmssn_prod_limit_min" json:"loan_sbmssn_prod_limit_min"`
	LoanSubmissionProductLimitMax            uint64     `gorm:"column:loan_sbmssn_prod_limit_max" json:"loan_sbmssn_prod_limit_max"`
	LoanSubmissionProductTenorMin            uint64     `gorm:"column:loan_sbmssn_prod_tenor_min" json:"loan_sbmssn_prod_tenor_min"`
	LoanSubmissionProductTenorMax            uint64     `gorm:"column:loan_sbmssn_prod_tenor_max" json:"loan_sbmssn_prod_tenor_max"`
	LoanSubmissionProductInterestRate        float64    `gorm:"column:loan_sbmssn_prod_interest_rate" json:"loan_sbmssn_prod_interest_rate"`
	LoanSubmissionProductInterestRatePeriod  string     `gorm:"column:loan_sbmssn_prod_interest_rate_period;size:125" json:"loan_sbmssn_prod_interest_rate_period"`
	LoanSubmissionProductInterestRateType    string     `gorm:"column:loan_sbmssn_prod_interest_rate_formula;size:125" json:"loan_sbmssn_prod_interest_rate_formula"`
	LoanSubmissionProductInterestRateDiscMax float64    `gorm:"column:loan_sbmssn_prod_interest_rate_disc_max" json:"loan_sbmssn_prod_interest_rate_disc_max"`
	LoanSubmissionProductNumReviewer         uint64     `gorm:"column:loan_sbmssn_prod_num_reviewer;size:2" json:"loan_sbmssn_prod_num_reviewer"`
	LoanSubmissionProductNumApprover         uint64     `gorm:"column:loan_sbmssn_prod_num_approver;size:2" json:"loan_sbmssn_prod_num_approver"`
	LoanSubmissionProductStatus              uint64     `gorm:"column:loan_sbmssn_prod_status;size:2" json:"loan_sbmssn_prod_status"`
	LoanSubmissionProductCreatedAt           time.Time  `gorm:"column:loan_sbmssn_prod_created_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_prod_created_at"`
	LoanSubmissionProductCreatedBy           string     `gorm:"column:loan_sbmssn_prod_created_by;size:125" json:"loan_sbmssn_prod_created_by"`
	LoanSubmissionProductUpdatedAt           *time.Time `gorm:"column:loan_sbmssn_prod_updated_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_prod_updated_at"`
	LoanSubmissionProductUpdatedBy           string     `gorm:"column:loan_sbmssn_prod_updated_by;size:125" json:"loan_sbmssn_prod_updated_by"`
	LoanSubmissionProductDeletedAt           *time.Time `gorm:"column:loan_sbmssn_prod_deleted_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_prod_deleted_at"`
	LoanSubmissionProductDeletedBy           string     `gorm:"column:loan_sbmssn_prod_deleted_by;size:125" json:"loan_sbmssn_prod_deleted_by"`
}

type LoanSubmissionProductData struct {
	LoanSubmissionProductID                  uint64     `gorm:"column:loan_sbmssn_prod_id" json:"loan_sbmssn_prod_id"`
	LoanSubmissionID                         uint64     `gorm:"column:loan_sbmssn_id" json:"loan_sbmssn_id"`
	MitraProductID                           uint64     `gorm:"column:mitra_prod_id" json:"mitra_prod_id"`
	LoanSubmissionProductName                string     `gorm:"column:loan_sbmssn_prod_name" json:"loan_sbmssn_prod_name"`
	LoanSubmissionProductLimitMin            uint64     `gorm:"column:loan_sbmssn_prod_limit_min" json:"loan_sbmssn_prod_limit_min"`
	LoanSubmissionProductLimitMax            uint64     `gorm:"column:loan_sbmssn_prod_limit_max" json:"loan_sbmssn_prod_limit_max"`
	LoanSubmissionProductTenorMin            uint64     `gorm:"column:loan_sbmssn_prod_tenor_min" json:"loan_sbmssn_prod_tenor_min"`
	LoanSubmissionProductTenorMax            uint64     `gorm:"column:loan_sbmssn_prod_tenor_max" json:"loan_sbmssn_prod_tenor_max"`
	LoanSubmissionProductInterestRate        float64    `gorm:"column:loan_sbmssn_prod_interest_rate" json:"loan_sbmssn_prod_interest_rate"`
	LoanSubmissionProductInterestRatePeriod  string     `gorm:"column:loan_sbmssn_prod_interest_rate_period" json:"loan_sbmssn_prod_interest_rate_period"`
	LoanSubmissionProductInterestRateType    string     `gorm:"column:loan_sbmssn_prod_interest_rate_formula" json:"loan_sbmssn_prod_interest_rate_formula"`
	LoanSubmissionProductInterestRateDiscMax float64    `gorm:"column:loan_sbmssn_prod_interest_rate_disc_max" json:"loan_sbmssn_prod_interest_rate_disc_max"`
	LoanSubmissionProductNumReviewer         uint64     `gorm:"column:loan_sbmssn_prod_num_reviewer" json:"loan_sbmssn_prod_num_reviewer"`
	LoanSubmissionProductNumApprover         uint64     `gorm:"column:loan_sbmssn_prod_num_approver" json:"loan_sbmssn_prod_num_approver"`
	LoanSubmissionProductStatus              uint64     `gorm:"column:loan_sbmssn_prod_status" json:"loan_sbmssn_prod_status"`
	LoanSubmissionProductCreatedBy           string     `gorm:"column:loan_sbmssn_prod_created_by" json:"loan_sbmssn_prod_created_by"`
	LoanSubmissionProductUpdatedAt           *time.Time `gorm:"column:loan_sbmssn_prod_updated_at" json:"loan_sbmssn_prod_updated_at"`
	LoanSubmissionProductUpdatedBy           string     `gorm:"column:loan_sbmssn_prod_updated_by" json:"loan_sbmssn_prod_updated_by"`
	LoanSubmissionProductDeletedAt           *time.Time `gorm:"column:loan_sbmssn_prod_deleted_at" json:"loan_sbmssn_prod_deleted_at"`
	LoanSubmissionProductDeletedBy           string     `gorm:"column:loan_sbmssn_prod_deleted_by" json:"loan_sbmssn_prod_deleted_by"`
}

type ResponseLoanSubmissionProducts struct {
	Status  int                          `json:"status"`
	Message string                       `json:"message"`
	Data    *[]LoanSubmissionProductData `json:"data"`
}

type ResponseLoanSubmissionProduct struct {
	Status  int                        `json:"status"`
	Message string                     `json:"message"`
	Data    *LoanSubmissionProductData `json:"data"`
}

type ResponseLoanSubmissionProductCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (LoanSubmissionProduct) TableName() string {
	return "t_loan_submission_product"
}

func (LoanSubmissionProductData) TableName() string {
	return "t_loan_submission_product"
}

func (p *LoanSubmissionProduct) Prepare() {
	p.LoanSubmissionID = p.LoanSubmissionID
	p.MitraProductID = p.MitraProductID
	p.LoanSubmissionProductName = html.EscapeString(strings.TrimSpace(p.LoanSubmissionProductName))
	p.LoanSubmissionProductLimitMin = p.LoanSubmissionProductLimitMin
	p.LoanSubmissionProductLimitMax = p.LoanSubmissionProductLimitMax
	p.LoanSubmissionProductTenorMin = p.LoanSubmissionProductTenorMin
	p.LoanSubmissionProductTenorMax = p.LoanSubmissionProductTenorMax
	p.LoanSubmissionProductInterestRate = p.LoanSubmissionProductInterestRate
	p.LoanSubmissionProductInterestRateDiscMax = p.LoanSubmissionProductInterestRateDiscMax
	p.LoanSubmissionProductNumReviewer = p.LoanSubmissionProductNumReviewer
	p.LoanSubmissionProductNumApprover = p.LoanSubmissionProductNumApprover
	p.LoanSubmissionProductStatus = p.LoanSubmissionProductStatus
	p.LoanSubmissionProductCreatedAt = time.Now()
	p.LoanSubmissionProductCreatedBy = p.LoanSubmissionProductCreatedBy
	p.LoanSubmissionProductUpdatedAt = p.LoanSubmissionProductUpdatedAt
	p.LoanSubmissionProductUpdatedBy = p.LoanSubmissionProductUpdatedBy
	p.LoanSubmissionProductDeletedAt = p.LoanSubmissionProductDeletedAt
	p.LoanSubmissionProductDeletedBy = p.LoanSubmissionProductDeletedBy
}

func (p *LoanSubmissionProduct) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.LoanSubmissionID == 0 {
			return errors.New("Required Mitra Product Name")
		}
		if p.MitraProductID == 0 {
			return errors.New("Required Mitra Product Name")
		}
		if p.LoanSubmissionProductName == "" {
			return errors.New("Required Mitra Product Name")
		}
		if p.LoanSubmissionProductName == "" {
			return errors.New("Required Mitra Product Name")
		}
		if p.LoanSubmissionProductLimitMin == 0 {
			return errors.New("Required Mitra Product Name")
		}
		if p.LoanSubmissionProductLimitMax == 0 {
			return errors.New("Required Mitra Product Name")
		}
		if p.LoanSubmissionProductTenorMin == 0 {
			return errors.New("Required Mitra Product Name")
		}
		if p.LoanSubmissionProductTenorMax == 0 {
			return errors.New("Required Mitra Product Name")
		}
		if p.LoanSubmissionProductInterestRate == 0 {
			return errors.New("Required Mitra Product Name")
		}
		if p.LoanSubmissionProductInterestRatePeriod == "" {
			return errors.New("Required Mitra Product Name")
		}
		if p.LoanSubmissionProductInterestRateType == "" {
			return errors.New("Required Mitra Product Name")
		}
		if p.LoanSubmissionProductInterestRateDiscMax == 0 {
			return errors.New("Required Mitra Product Name")
		}
		if p.LoanSubmissionProductNumReviewer == 0 {
			return errors.New("Required Mitra Product Name")
		}
		if p.LoanSubmissionProductNumApprover == 0 {
			return errors.New("Required Mitra Product Name")
		}
		return nil
	}
}

func (p *LoanSubmissionProduct) SaveLoanSubmissionProduct(db *gorm.DB) (*LoanSubmissionProduct, error) {
	var err error
	err = db.Debug().Model(&LoanSubmissionProduct{}).Create(&p).Error
	if err != nil {
		return &LoanSubmissionProduct{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProduct) FindAllLoanSubmissionProduct(db *gorm.DB) (*[]LoanSubmissionProductData, error) {
	var err error
	product := []LoanSubmissionProductData{}
	err = db.Debug().Model(&LoanSubmissionProductData{}).Limit(100).
		Select(`t_loan_submission_product.loan_sbmssn_prod_id,
				t_loan_submission_product.loan_sbmssn_id,
				t_loan_submission_product.mitra_prod_id,
				t_loan_submission_product.loan_sbmssn_prod_name,
				t_loan_submission_product.loan_sbmssn_prod_limit_min,
				t_loan_submission_product.loan_sbmssn_prod_limit_max,
				t_loan_submission_product.loan_sbmssn_prod_tenor_min,
				t_loan_submission_product.loan_sbmssn_prod_tenor_max,
				t_loan_submission_product.loan_sbmssn_prod_interest_rate,
				t_loan_submission_product.loan_sbmssn_prod_interest_rate_period,
				t_loan_submission_product.loan_sbmssn_prod_interest_rate_formula,
				t_loan_submission_product.loan_sbmssn_prod_interest_rate_disc_max,
				t_loan_submission_product.loan_sbmssn_prod_num_reviewer,
				t_loan_submission_product.loan_sbmssn_prod_num_approver,
				t_loan_submission_product.loan_sbmssn_prod_status,
				t_loan_submission_product.loan_sbmssn_prod_created_at,
				t_loan_submission_product.loan_sbmssn_prod_created_by,
				t_loan_submission_product.loan_sbmssn_prod_updated_at,
				t_loan_submission_product.loan_sbmssn_prod_updated_by,
				t_loan_submission_product.loan_sbmssn_prod_deleted_at,
				t_loan_submission_product.loan_sbmssn_prod_deleted_by
				`).
		Find(&product).Error
	if err != nil {
		return &[]LoanSubmissionProductData{}, err
	}
	return &product, nil
}

func (p *LoanSubmissionProduct) FindAllLoanSubmissionProductStatus(db *gorm.DB, status uint64) (*[]LoanSubmissionProductData, error) {
	var err error
	product := []LoanSubmissionProductData{}
	err = db.Debug().Model(&LoanSubmissionProductData{}).Limit(100).
		Select(`t_loan_submission_product.loan_sbmssn_prod_id,
				t_loan_submission_product.loan_sbmssn_id,
				t_loan_submission_product.mitra_prod_id,
				t_loan_submission_product.loan_sbmssn_prod_name,
				t_loan_submission_product.loan_sbmssn_prod_limit_min,
				t_loan_submission_product.loan_sbmssn_prod_limit_max,
				t_loan_submission_product.loan_sbmssn_prod_tenor_min,
				t_loan_submission_product.loan_sbmssn_prod_tenor_max,
				t_loan_submission_product.loan_sbmssn_prod_interest_rate,
				t_loan_submission_product.loan_sbmssn_prod_interest_rate_period,
				t_loan_submission_product.loan_sbmssn_prod_interest_rate_formula,
				t_loan_submission_product.loan_sbmssn_prod_interest_rate_disc_max,
				t_loan_submission_product.loan_sbmssn_prod_num_reviewer,
				t_loan_submission_product.loan_sbmssn_prod_num_approver,
				t_loan_submission_product.loan_sbmssn_prod_status,
				t_loan_submission_product.loan_sbmssn_prod_created_at,
				t_loan_submission_product.loan_sbmssn_prod_created_by,
				t_loan_submission_product.loan_sbmssn_prod_updated_at,
				t_loan_submission_product.loan_sbmssn_prod_updated_by,
				t_loan_submission_product.loan_sbmssn_prod_deleted_at,
				t_loan_submission_product.loan_sbmssn_prod_deleted_by
				`).
		Where("loan_sbmssn_prod_status = ?", status).
		Find(&product).Error
	if err != nil {
		return &[]LoanSubmissionProductData{}, err
	}
	return &product, nil
}

func (p *LoanSubmissionProduct) FindLoanSubmissionProductDataByID(db *gorm.DB, pid uint64) (*LoanSubmissionProduct, error) {
	var err error
	err = db.Debug().Model(&LoanSubmissionProduct{}).Where("loan_sbmssn_prod_id = ?", pid).Take(&p).Error
	if err != nil {
		return &LoanSubmissionProduct{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProduct) FindLoanSubmissionProductByID(db *gorm.DB, pid uint64) (*LoanSubmissionProductData, error) {
	var err error
	product := LoanSubmissionProductData{}
	err = db.Debug().Model(&LoanSubmissionProductData{}).Limit(100).
		Select(`t_loan_submission_product.loan_sbmssn_prod_id,
				t_loan_submission_product.loan_sbmssn_id,
				t_loan_submission_product.mitra_prod_id,
				t_loan_submission_product.loan_sbmssn_prod_name,
				t_loan_submission_product.loan_sbmssn_prod_limit_min,
				t_loan_submission_product.loan_sbmssn_prod_limit_max,
				t_loan_submission_product.loan_sbmssn_prod_tenor_min,
				t_loan_submission_product.loan_sbmssn_prod_tenor_max,
				t_loan_submission_product.loan_sbmssn_prod_interest_rate,
				t_loan_submission_product.loan_sbmssn_prod_interest_rate_period,
				t_loan_submission_product.loan_sbmssn_prod_interest_rate_formula,
				t_loan_submission_product.loan_sbmssn_prod_interest_rate_disc_max,
				t_loan_submission_product.loan_sbmssn_prod_num_reviewer,
				t_loan_submission_product.loan_sbmssn_prod_num_approver,
				t_loan_submission_product.loan_sbmssn_prod_status,
				t_loan_submission_product.loan_sbmssn_prod_created_at,
				t_loan_submission_product.loan_sbmssn_prod_created_by,
				t_loan_submission_product.loan_sbmssn_prod_updated_at,
				t_loan_submission_product.loan_sbmssn_prod_updated_by,
				t_loan_submission_product.loan_sbmssn_prod_deleted_at,
				t_loan_submission_product.loan_sbmssn_prod_deleted_by
				`).
		Where("loan_sbmssn_prod_id = ?", pid).
		Find(&product).Error
	if err != nil {
		return &LoanSubmissionProductData{}, err
	}
	return &product, nil
}

func (p *LoanSubmissionProduct) FindLoanSubmissionProductStatusByID(db *gorm.DB, pid uint64, status uint64) (*LoanSubmissionProductData, error) {
	var err error
	product := LoanSubmissionProductData{}
	err = db.Debug().Model(&LoanSubmissionProductData{}).Limit(100).
		Select(`t_loan_submission_product.loan_sbmssn_prod_id,
				t_loan_submission_product.loan_sbmssn_id,
				t_loan_submission_product.mitra_prod_id,
				t_loan_submission_product.loan_sbmssn_prod_name,
				t_loan_submission_product.loan_sbmssn_prod_limit_min,
				t_loan_submission_product.loan_sbmssn_prod_limit_max,
				t_loan_submission_product.loan_sbmssn_prod_tenor_min,
				t_loan_submission_product.loan_sbmssn_prod_tenor_max,
				t_loan_submission_product.loan_sbmssn_prod_interest_rate,
				t_loan_submission_product.loan_sbmssn_prod_interest_rate_period,
				t_loan_submission_product.loan_sbmssn_prod_interest_rate_formula,
				t_loan_submission_product.loan_sbmssn_prod_interest_rate_disc_max,
				t_loan_submission_product.loan_sbmssn_prod_num_reviewer,
				t_loan_submission_product.loan_sbmssn_prod_num_approver,
				t_loan_submission_product.loan_sbmssn_prod_status,
				t_loan_submission_product.loan_sbmssn_prod_created_at,
				t_loan_submission_product.loan_sbmssn_prod_created_by,
				t_loan_submission_product.loan_sbmssn_prod_updated_at,
				t_loan_submission_product.loan_sbmssn_prod_updated_by,
				t_loan_submission_product.loan_sbmssn_prod_deleted_at,
				t_loan_submission_product.loan_sbmssn_prod_deleted_by
			`).
		Where("loan_sbmssn_prod_id = ? AND loan_sbmssn_prod_status = ?", pid, status).
		Find(&product).Error
	if err != nil {
		return &LoanSubmissionProductData{}, err
	}
	return &product, nil
}

func (p *LoanSubmissionProduct) UpdateLoanSubmissionProduct(db *gorm.DB, pid uint64) (*LoanSubmissionProduct, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionProduct{}).Where("loan_sbmssn_prod_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_prod_name":                   p.LoanSubmissionProductName,
			"loan_sbmssn_prod_limit_min":              p.LoanSubmissionProductLimitMin,
			"loan_sbmssn_prod_limit_max":              p.LoanSubmissionProductLimitMax,
			"loan_sbmssn_prod_tenor_min":              p.LoanSubmissionProductTenorMin,
			"loan_sbmssn_prod_tenor_max":              p.LoanSubmissionProductTenorMax,
			"loan_sbmssn_prod_interest_rate":          p.LoanSubmissionProductInterestRate,
			"loan_sbmssn_prod_interest_rate_period":   p.LoanSubmissionProductInterestRatePeriod,
			"loan_sbmssn_prod_interest_rate_formula":  p.LoanSubmissionProductInterestRateType,
			"loan_sbmssn_prod_interest_rate_disc_max": p.LoanSubmissionProductInterestRateDiscMax,
			"loan_sbmssn_prod_num_reviewer":           p.LoanSubmissionProductNumReviewer,
			"loan_sbmssn_prod_num_approver":           p.LoanSubmissionProductNumApprover,
			"loan_sbmssn_prod_status":                 p.LoanSubmissionProductStatus,
			"loan_sbmssn_prod_updated_by":             p.LoanSubmissionProductUpdatedBy,
			"loan_sbmssn_prod_updated_at":             time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionProduct{}).Where("loan_sbmssn_prod_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionProduct{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProduct) UpdateLoanSubmissionProductStatus(db *gorm.DB, pid uint64) (*LoanSubmissionProduct, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionProduct{}).Where("loan_sbmssn_prod_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_prod_status":     p.LoanSubmissionProductStatus,
			"loan_sbmssn_prod_updated_by": p.LoanSubmissionProductUpdatedBy,
			"loan_sbmssn_prod_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionProduct{}).Where("loan_sbmssn_prod_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionProduct{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProduct) UpdateLoanSubmissionProductStatusDelete(db *gorm.DB, pid uint64) (*LoanSubmissionProduct, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionProduct{}).Where("loan_sbmssn_prod_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_prod_status":     3,
			"loan_sbmssn_prod_deleted_by": p.LoanSubmissionProductDeletedBy,
			"loan_sbmssn_prod_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionProduct{}).Where("loan_sbmssn_prod_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionProduct{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProduct) DeleteLoanSubmissionProduct(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&LoanSubmissionProduct{}).Where("loan_sbmssn_prod_id = ?", pid).Take(&LoanSubmissionProduct{}).Delete(&LoanSubmissionProduct{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Mitra Product not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
