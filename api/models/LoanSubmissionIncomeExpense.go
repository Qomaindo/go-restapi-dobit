package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type LoanSubmissionIncomeExpense struct {
	LoanSubmissionIncomeExpenseID                 uint64     `gorm:"column:loan_sbmssn_incm_expn_id;primary_key;" json:"loan_sbmssn_incm_expn_id"`
	LoanSubmissionID                              uint64     `gorm:"column:loan_sbmssn_id" json:"loan_sbmssn_id"`
	LoanSubmissionIncomeExpenseSalary             uint64     `gorm:"column:loan_sbmssn_incm_expn_salary" json:"loan_sbmssn_incm_expn_salary"`
	LoanSubmissionIncomeExpenseSalaryImage        string     `gorm:"column:loan_sbmssn_incm_expn_salary_image;size:255" json:"loan_sbmssn_incm_expn_salary_image"`
	LoanSubmissionIncomeExpenseSalaryPartner      uint64     `gorm:"column:loan_sbmssn_incm_expn_salary_partner" json:"loan_sbmssn_incm_expn_salary_partner"`
	LoanSubmissionIncomeExpenseSalaryPartnerImage string     `gorm:"column:loan_sbmssn_incm_expn_salary_partner_image;size:255" json:"loan_sbmssn_incm_expn_salary_partner_image"`
	LoanSubmissionIncomeExpenseOther              uint64     `gorm:"column:loan_sbmssn_incm_expn_other" json:"loan_sbmssn_incm_expn_other"`
	LoanSubmissionIncomeExpenseOtherImage         string     `gorm:"column:loan_sbmssn_incm_expn_other_image;size:255" json:"loan_sbmssn_incm_expn_other_image"`
	LoanSubmissionIncomeExpenseCostHousehold      uint64     `gorm:"column:loan_sbmssn_incm_expn_cost_household" json:"loan_sbmssn_incm_expn_cost_household"`
	LoanSubmissionIncomeExpenseCostEducation      uint64     `gorm:"column:loan_sbmssn_incm_expn_cost_education" json:"loan_sbmssn_incm_expn_cost_education"`
	LoanSubmissionIncomeExpenseCostUtility        uint64     `gorm:"column:loan_sbmssn_incm_expn_cost_utility" json:"loan_sbmssn_incm_expn_cost_utility"`
	LoanSubmissionIncomeExpenseCostRent           uint64     `gorm:"column:loan_sbmssn_incm_expn_cost_rent" json:"loan_sbmssn_incm_expn_cost_rent"`
	LoanSubmissionIncomeExpenseCostInstallment    uint64     `gorm:"column:loan_sbmssn_incm_expn_cost_installment" json:"loan_sbmssn_incm_expn_cost_installment"`
	LoanSubmissionIncomeExpenseStatus             uint64     `gorm:"column:loan_sbmssn_incm_expn_status" json:"loan_sbmssn_incm_expn_status"`
	LoanSubmissionIncomeExpenseCreatedBy          string     `gorm:"column:loan_sbmssn_incm_expn_created_by;size:125" json:"loan_sbmssn_incm_expn_created_by"`
	LoanSubmissionIncomeExpenseCreatedAt          time.Time  `gorm:"column:loan_sbmssn_incm_expn_created_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_incm_expn_created_at"`
	LoanSubmissionIncomeExpenseUpdatedBy          string     `gorm:"column:loan_sbmssn_incm_expn_updated_by;size:125" json:"loan_sbmssn_incm_expn_updated_by"`
	LoanSubmissionIncomeExpenseUpdatedAt          *time.Time `gorm:"column:loan_sbmssn_incm_expn_updated_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_incm_expn_updated_at"`
	LoanSubmissionIncomeExpenseDeletedBy          string     `gorm:"column:loan_sbmssn_incm_expn_deleted_by;size:125" json:"loan_sbmssn_incm_expn_deleted_by"`
	LoanSubmissionIncomeExpenseDeletedAt          *time.Time `gorm:"column:loan_sbmssn_incm_expn_deleted_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_incm_expn_deleted_at"`
}

type LoanSubmissionIncomeExpenseData struct {
	LoanSubmissionIncomeExpenseID                 uint64     `gorm:"column:loan_sbmssn_incm_expn_id" json:"loan_sbmssn_incm_expn_id"`
	LoanSubmissionID                              uint64     `gorm:"column:loan_sbmssn_id" json:"loan_sbmssn_id"`
	LoanSubmissionIncomeExpenseSalary             uint64     `gorm:"column:loan_sbmssn_incm_expn_salary" json:"loan_sbmssn_incm_expn_salary"`
	LoanSubmissionIncomeExpenseSalaryImage        string     `gorm:"column:loan_sbmssn_incm_expn_salary_image" json:"loan_sbmssn_incm_expn_salary_image"`
	LoanSubmissionIncomeExpenseSalaryPartner      uint64     `gorm:"column:loan_sbmssn_incm_expn_salary_partner" json:"loan_sbmssn_incm_expn_salary_partner"`
	LoanSubmissionIncomeExpenseSalaryPartnerImage string     `gorm:"column:loan_sbmssn_incm_expn_salary_partner_image" json:"loan_sbmssn_incm_expn_salary_partner_image"`
	LoanSubmissionIncomeExpenseOther              uint64     `gorm:"column:loan_sbmssn_incm_expn_other" json:"loan_sbmssn_incm_expn_other"`
	LoanSubmissionIncomeExpenseOtherImage         string     `gorm:"column:loan_sbmssn_incm_expn_other_image" json:"loan_sbmssn_incm_expn_other_image"`
	LoanSubmissionIncomeExpenseCostHousehold      uint64     `gorm:"column:loan_sbmssn_incm_expn_cost_household" json:"loan_sbmssn_incm_expn_cost_household"`
	LoanSubmissionIncomeExpenseCostEducation      uint64     `gorm:"column:loan_sbmssn_incm_expn_cost_education" json:"loan_sbmssn_incm_expn_cost_education"`
	LoanSubmissionIncomeExpenseCostUtility        uint64     `gorm:"column:loan_sbmssn_incm_expn_cost_utility" json:"loan_sbmssn_incm_expn_cost_utility"`
	LoanSubmissionIncomeExpenseCostRent           uint64     `gorm:"column:loan_sbmssn_incm_expn_cost_rent" json:"loan_sbmssn_incm_expn_cost_rent"`
	LoanSubmissionIncomeExpenseCostInstallment    uint64     `gorm:"column:loan_sbmssn_incm_expn_cost_installment" json:"loan_sbmssn_incm_expn_cost_installment"`
	LoanSubmissionIncomeExpenseStatus             uint64     `gorm:"column:loan_sbmssn_incm_expn_status" json:"loan_sbmssn_incm_expn_status"`
	LoanSubmissionIncomeExpenseCreatedBy          string     `gorm:"column:loan_sbmssn_incm_expn_created_by" json:"loan_sbmssn_incm_expn_created_by"`
	LoanSubmissionIncomeExpenseCreatedAt          time.Time  `gorm:"column:loan_sbmssn_incm_expn_created_at" json:"loan_sbmssn_incm_expn_created_at"`
	LoanSubmissionIncomeExpenseUpdatedBy          string     `gorm:"column:loan_sbmssn_incm_expn_updated_by" json:"loan_sbmssn_incm_expn_updated_by"`
	LoanSubmissionIncomeExpenseUpdatedAt          *time.Time `gorm:"column:loan_sbmssn_incm_expn_updated_at" json:"loan_sbmssn_incm_expn_updated_at"`
	LoanSubmissionIncomeExpenseDeletedBy          string     `gorm:"column:loan_sbmssn_incm_expn_deleted_by" json:"loan_sbmssn_incm_expn_deleted_by"`
	LoanSubmissionIncomeExpenseDeletedAt          *time.Time `gorm:"column:loan_sbmssn_incm_expn_deleted_at" json:"loan_sbmssn_incm_expn_deleted_at"`
}

type ResponseLoanSubmissionIncomeExpenses struct {
	Status  int                                `json:"status"`
	Message string                             `json:"message"`
	Data    *[]LoanSubmissionIncomeExpenseData `json:"data"`
}

type ResponseLoanSubmissionIncomeExpense struct {
	Status  int                              `json:"status"`
	Message string                           `json:"message"`
	Data    *LoanSubmissionIncomeExpenseData `json:"data"`
}

type ResponseLoanSubmissionIncomeExpenseCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (LoanSubmissionIncomeExpense) TableName() string {
	return "t_loan_submission_income_expense"
}

func (LoanSubmissionIncomeExpenseData) TableName() string {
	return "t_loan_submission_income_expense"
}

func (p *LoanSubmissionIncomeExpense) Prepare() {
	p.LoanSubmissionID = p.LoanSubmissionID
	p.LoanSubmissionIncomeExpenseSalary = p.LoanSubmissionIncomeExpenseSalary
	p.LoanSubmissionIncomeExpenseSalaryImage = html.EscapeString(strings.TrimSpace(p.LoanSubmissionIncomeExpenseSalaryImage))
	p.LoanSubmissionIncomeExpenseSalaryPartner = p.LoanSubmissionIncomeExpenseSalaryPartner
	p.LoanSubmissionIncomeExpenseSalaryPartnerImage = html.EscapeString(strings.TrimSpace(p.LoanSubmissionIncomeExpenseSalaryPartnerImage))
	p.LoanSubmissionIncomeExpenseOther = p.LoanSubmissionIncomeExpenseOther
	p.LoanSubmissionIncomeExpenseOtherImage = html.EscapeString(strings.TrimSpace(p.LoanSubmissionIncomeExpenseOtherImage))
	p.LoanSubmissionIncomeExpenseCostHousehold = p.LoanSubmissionIncomeExpenseCostHousehold
	p.LoanSubmissionIncomeExpenseCostEducation = p.LoanSubmissionIncomeExpenseCostEducation
	p.LoanSubmissionIncomeExpenseCostUtility = p.LoanSubmissionIncomeExpenseCostUtility
	p.LoanSubmissionIncomeExpenseCostRent = p.LoanSubmissionIncomeExpenseCostRent
	p.LoanSubmissionIncomeExpenseCostInstallment = p.LoanSubmissionIncomeExpenseCostInstallment
	p.LoanSubmissionIncomeExpenseStatus = p.LoanSubmissionIncomeExpenseStatus
	p.LoanSubmissionIncomeExpenseCreatedBy = p.LoanSubmissionIncomeExpenseCreatedBy
	p.LoanSubmissionIncomeExpenseCreatedAt = time.Now()
	p.LoanSubmissionIncomeExpenseUpdatedBy = p.LoanSubmissionIncomeExpenseUpdatedBy
	p.LoanSubmissionIncomeExpenseUpdatedAt = p.LoanSubmissionIncomeExpenseUpdatedAt
	p.LoanSubmissionIncomeExpenseDeletedBy = p.LoanSubmissionIncomeExpenseDeletedBy
	p.LoanSubmissionIncomeExpenseDeletedAt = p.LoanSubmissionIncomeExpenseDeletedAt
}

func (p *LoanSubmissionIncomeExpense) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.LoanSubmissionID == 0 {
			return errors.New("Required Loan Submission ID")
		}
		if p.LoanSubmissionIncomeExpenseSalary == 0 {
			return errors.New("Required Salary")
		}
		if p.LoanSubmissionIncomeExpenseSalaryPartnerImage == "" {
			return errors.New("Required salary partner Image")
		}
		if p.LoanSubmissionIncomeExpenseSalaryPartner == 0 {
			return errors.New("Required Partner salary")
		}
		if p.LoanSubmissionIncomeExpenseSalaryPartnerImage == "" {
			return errors.New("Required Partner Image")
		}
		if p.LoanSubmissionIncomeExpenseOther == 0 {
			return errors.New("Required Income Other")
		}
		if p.LoanSubmissionIncomeExpenseOtherImage == "" {
			return errors.New("Required other image")
		}
		if p.LoanSubmissionIncomeExpenseCostHousehold == 0 {
			return errors.New("Required Cost Household")
		}
		if p.LoanSubmissionIncomeExpenseCostEducation == 0 {
			return errors.New("Required cost education")
		}
		if p.LoanSubmissionIncomeExpenseCostUtility == 0 {
			return errors.New("Required Cost Utility")
		}
		if p.LoanSubmissionIncomeExpenseCostRent == 0 {
			return errors.New("Required Cost Rent")
		}
		if p.LoanSubmissionIncomeExpenseCostInstallment == 0 {
			return errors.New("Required Cost Installment")
		}
		return nil
	}
}

func (p *LoanSubmissionIncomeExpense) SaveLoanSubmissionIncomeExpense(db *gorm.DB) (*LoanSubmissionIncomeExpense, error) {
	var err error
	err = db.Debug().Model(&LoanSubmissionIncomeExpense{}).Create(&p).Error
	if err != nil {
		return &LoanSubmissionIncomeExpense{}, err
	}
	return p, nil
}

func (p *LoanSubmissionIncomeExpense) FindAllLoanSubmissionIncomeExpense(db *gorm.DB) (*[]LoanSubmissionIncomeExpenseData, error) {
	var err error
	loansbmssnincmexpn := []LoanSubmissionIncomeExpenseData{}
	err = db.Debug().Model(&LoanSubmissionIncomeExpenseData{}).Limit(100).
		Select(`t_loan_submission_income_expense.loan_sbmssn_incm_expn_id,
				t_loan_submission_income_expense.loan_sbmssn_id,
				t_loan_submission_income_expense.loan_sbmssn_incm_expn_salary,
				t_loan_submission_income_expense.loan_sbmssn_incm_expn_salary_image,
				t_loan_submission_income_expense.loan_sbmssn_incm_expn_salary_partner,
				t_loan_submission_income_expense.loan_sbmssn_incm_expn_salary_partner_image,
				t_loan_submission_income_expense.loan_sbmssn_incm_expn_other,
				t_loan_submission_income_expense.loan_sbmssn_incm_expn_other_image,
				t_loan_submission_income_expense.loan_sbmssn_incm_expn_cost_household,
				t_loan_submission_income_expense.loan_sbmssn_incm_expn_cost_education,
				t_loan_submission_income_expense.loan_sbmssn_incm_expn_cost_utility,
				t_loan_submission_income_expense.loan_sbmssn_incm_expn_cost_rent,
				t_loan_submission_income_expense.loan_sbmssn_incm_expn_cost_installment,
				t_loan_submission_income_expense.loan_sbmssn_incm_expn_status,
				t_loan_submission_income_expense.loan_sbmssn_incm_expn_created_by,
				t_loan_submission_income_expense.loan_sbmssn_incm_expn_created_at,
				t_loan_submission_income_expense.loan_sbmssn_incm_expn_updated_by,
				t_loan_submission_income_expense.loan_sbmssn_incm_expn_updated_at,
				t_loan_submission_income_expense.loan_sbmssn_incm_expn_deleted_by,
				t_loan_submission_income_expense.loan_sbmssn_incm_expn_deleted_at
				`).
		Find(&loansbmssnincmexpn).Error
	if err != nil {
		return &[]LoanSubmissionIncomeExpenseData{}, err
	}
	return &loansbmssnincmexpn, nil
}

func (p *LoanSubmissionIncomeExpense) FindLoanSubmissionIncomeExpenseDataByID(db *gorm.DB, pid uint64) (*LoanSubmissionIncomeExpense, error) {
	var err error
	err = db.Debug().Model(&LoanSubmissionIncomeExpense{}).Where("loan_sbmssn_incm_expn_id = ?", pid).Take(&p).Error
	if err != nil {
		return &LoanSubmissionIncomeExpense{}, err
	}
	return p, nil
}

func (p *LoanSubmissionIncomeExpense) FindLoanSubmissionIncomeExpenseByID(db *gorm.DB, pid uint64) (*LoanSubmissionIncomeExpenseData, error) {
	var err error
	loansbmssnincmexpn := LoanSubmissionIncomeExpenseData{}
	err = db.Debug().Model(&LoanSubmissionIncomeExpenseData{}).Limit(100).
		Select(`t_loan_submission_income_expense.loan_sbmssn_incm_expn_id,
		t_loan_submission_income_expense.loan_sbmssn_id,
		t_loan_submission_income_expense.loan_sbmssn_incm_expn_salary,
		t_loan_submission_income_expense.loan_sbmssn_incm_expn_salary_image,
		t_loan_submission_income_expense.loan_sbmssn_incm_expn_salary_partner,
		t_loan_submission_income_expense.loan_sbmssn_incm_expn_salary_partner_image,
		t_loan_submission_income_expense.loan_sbmssn_incm_expn_other,
		t_loan_submission_income_expense.loan_sbmssn_incm_expn_other_image,
		t_loan_submission_income_expense.loan_sbmssn_incm_expn_cost_household,
		t_loan_submission_income_expense.loan_sbmssn_incm_expn_cost_education,
		t_loan_submission_income_expense.loan_sbmssn_incm_expn_cost_utility,
		t_loan_submission_income_expense.loan_sbmssn_incm_expn_cost_rent,
		t_loan_submission_income_expense.loan_sbmssn_incm_expn_cost_installment,
		t_loan_submission_income_expense.loan_sbmssn_incm_expn_status,
		t_loan_submission_income_expense.loan_sbmssn_incm_expn_created_by,
		t_loan_submission_income_expense.loan_sbmssn_incm_expn_created_at,
		t_loan_submission_income_expense.loan_sbmssn_incm_expn_updated_by,
		t_loan_submission_income_expense.loan_sbmssn_incm_expn_updated_at,
		t_loan_submission_income_expense.loan_sbmssn_incm_expn_deleted_by,
		t_loan_submission_income_expense.loan_sbmssn_incm_expn_deleted_at
				`).
		Where("loan_sbmssn_incm_expn_id = ?", pid).
		Find(&loansbmssnincmexpn).Error
	if err != nil {
		return &LoanSubmissionIncomeExpenseData{}, err
	}
	return &loansbmssnincmexpn, nil
}

func (p *LoanSubmissionIncomeExpense) UpdateLoanSubmissionIncomeExpense(db *gorm.DB, pid uint64) (*LoanSubmissionIncomeExpense, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionIncomeExpense{}).Where("loan_sbmssn_incm_expn_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_incm_expn_salary":               p.LoanSubmissionIncomeExpenseSalary,
			"loan_sbmssn_incm_expn_salary_image":         p.LoanSubmissionIncomeExpenseSalaryImage,
			"loan_sbmssn_incm_expn_salary_partner":       p.LoanSubmissionIncomeExpenseSalaryPartner,
			"loan_sbmssn_incm_expn_salary_partner_image": p.LoanSubmissionIncomeExpenseSalaryPartnerImage,
			"loan_sbmssn_incm_expn_other":                p.LoanSubmissionIncomeExpenseOther,
			"loan_sbmssn_incm_expn_other_image":          p.LoanSubmissionIncomeExpenseSalaryPartnerImage,
			"loan_sbmssn_incm_expn_cost_household":       p.LoanSubmissionIncomeExpenseCostHousehold,
			"loan_sbmssn_incm_expn_cost_education":       p.LoanSubmissionIncomeExpenseCostEducation,
			"loan_sbmssn_incm_expn_cost_utility":         p.LoanSubmissionIncomeExpenseCostUtility,
			"loan_sbmssn_incm_expn_cost_rent":            p.LoanSubmissionIncomeExpenseCostRent,
			"loan_sbmssn_incm_expn_cost_installment":     p.LoanSubmissionIncomeExpenseCostInstallment,
			"loan_sbmssn_incm_expn_updated_by":           p.LoanSubmissionIncomeExpenseUpdatedBy,
			"loan_sbmssn_incm_expn_updated_at":           time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionIncomeExpense{}).Where("loan_sbmssn_incm_expn_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionIncomeExpense{}, err
	}
	return p, nil
}

func (p *LoanSubmissionIncomeExpense) UpdateLoanSubmissionIncomeExpenseDelete(db *gorm.DB, pid uint64) (*LoanSubmissionIncomeExpense, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionIncomeExpense{}).Where("loan_sbmssn_incm_expn_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_incm_expn_status":     p.LoanSubmissionIncomeExpenseStatus,
			"loan_sbmssn_incm_expn_deleted_by": p.LoanSubmissionIncomeExpenseDeletedBy,
			"loan_sbmssn_incm_expn_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionIncomeExpense{}).Where("loan_sbmssn_incm_expn_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionIncomeExpense{}, err
	}
	return p, nil
}

func (p *LoanSubmissionIncomeExpense) DeleteLoanSubmissionIncomeExpense(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&LoanSubmissionIncomeExpense{}).Where("loan_sbmssn_incm_expn_id = ?", pid).Take(&LoanSubmissionIncomeExpense{}).Delete(&LoanSubmissionIncomeExpense{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Loan Submission Income Expense not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// ADDITIONAL MEMORANDUM
// ===================================================================================================================================

type LoanSubmissionIncomeMrdmData struct {
	LoanSubmissionIncomeExpenseID                 uint64 `gorm:"column:loan_sbmssn_incm_expn_id" json:"loan_sbmssn_incm_expn_id"`
	LoanSubmissionID                              uint64 `gorm:"column:loan_sbmssn_id" json:"loan_sbmssn_id"`
	LoanSubmissionIncomeExpenseSalary             uint64 `gorm:"column:loan_sbmssn_incm_expn_salary" json:"loan_sbmssn_incm_expn_salary"`
	LoanSubmissionIncomeExpenseSalaryImage        string `gorm:"column:loan_sbmssn_incm_expn_salary_image" json:"loan_sbmssn_incm_expn_salary_image"`
	LoanSubmissionIncomeExpenseSalaryPartner      uint64 `gorm:"column:loan_sbmssn_incm_expn_salary_partner" json:"loan_sbmssn_incm_expn_salary_partner"`
	LoanSubmissionIncomeExpenseSalaryPartnerImage string `gorm:"column:loan_sbmssn_incm_expn_salary_partner_image" json:"loan_sbmssn_incm_expn_salary_partner_image"`
	LoanSubmissionIncomeExpenseOther              uint64 `gorm:"column:loan_sbmssn_incm_expn_other" json:"loan_sbmssn_incm_expn_other"`
	LoanSubmissionIncomeExpenseOtherImage         string `gorm:"column:loan_sbmssn_incm_expn_other_image" json:"loan_sbmssn_incm_expn_other_image"`
}

type LoanSubmissionExpenseMrdmData struct {
	LoanSubmissionIncomeExpenseID              uint64 `gorm:"column:loan_sbmssn_incm_expn_id" json:"loan_sbmssn_incm_expn_id"`
	LoanSubmissionID                           uint64 `gorm:"column:loan_sbmssn_id" json:"loan_sbmssn_id"`
	LoanSubmissionIncomeExpenseCostHousehold   uint64 `gorm:"column:loan_sbmssn_incm_expn_cost_household" json:"loan_sbmssn_incm_expn_cost_household"`
	LoanSubmissionIncomeExpenseCostEducation   uint64 `gorm:"column:loan_sbmssn_incm_expn_cost_education" json:"loan_sbmssn_incm_expn_cost_education"`
	LoanSubmissionIncomeExpenseCostUtility     uint64 `gorm:"column:loan_sbmssn_incm_expn_cost_utility" json:"loan_sbmssn_incm_expn_cost_utility"`
	LoanSubmissionIncomeExpenseCostRent        uint64 `gorm:"column:loan_sbmssn_incm_expn_cost_rent" json:"loan_sbmssn_incm_expn_cost_rent"`
	LoanSubmissionIncomeExpenseCostInstallment uint64 `gorm:"column:loan_sbmssn_incm_expn_cost_installment" json:"loan_sbmssn_incm_expn_cost_installment"`
}

func (LoanSubmissionIncomeMrdmData) TableName() string {
	return "t_loan_submission_income_expense"
}

func (LoanSubmissionExpenseMrdmData) TableName() string {
	return "t_loan_submission_income_expense"
}

func (p *LoanSubmissionIncomeExpense) FindLoanSubmissionIncomeMrdmByLoanSubmissionID(db *gorm.DB, pid uint64) (LoanSubmissionIncomeMrdmData, error) {
	var err error
	loansbmssnincmexpn := LoanSubmissionIncomeMrdmData{}
	err = db.Debug().Model(&LoanSubmissionIncomeMrdmData{}).Limit(100).
		Select(`t_loan_submission_income_expense.loan_sbmssn_incm_expn_id,
				t_loan_submission_income_expense.loan_sbmssn_id,
				t_loan_submission_income_expense.loan_sbmssn_incm_expn_salary,
				t_loan_submission_income_expense.loan_sbmssn_incm_expn_salary_image,
				t_loan_submission_income_expense.loan_sbmssn_incm_expn_salary_partner,
				t_loan_submission_income_expense.loan_sbmssn_incm_expn_salary_partner_image,
				t_loan_submission_income_expense.loan_sbmssn_incm_expn_other,
				t_loan_submission_income_expense.loan_sbmssn_incm_expn_other_image`).
		Where("loan_sbmssn_id = ?", pid).
		Find(&loansbmssnincmexpn).Error
	if err != nil {
		return LoanSubmissionIncomeMrdmData{}, err
	}
	return loansbmssnincmexpn, nil
}

func (p *LoanSubmissionIncomeExpense) FindLoanSubmissionExpenseMrdmByLoanSubmissionID(db *gorm.DB, pid uint64) (LoanSubmissionExpenseMrdmData, error) {
	var err error
	loansbmssnincmexpn := LoanSubmissionExpenseMrdmData{}
	err = db.Debug().Model(&LoanSubmissionExpenseMrdmData{}).Limit(100).
		Select(`t_loan_submission_income_expense.loan_sbmssn_incm_expn_id,
				t_loan_submission_income_expense.loan_sbmssn_id,
				t_loan_submission_income_expense.loan_sbmssn_incm_expn_cost_household,
				t_loan_submission_income_expense.loan_sbmssn_incm_expn_cost_education,
				t_loan_submission_income_expense.loan_sbmssn_incm_expn_cost_utility,
				t_loan_submission_income_expense.loan_sbmssn_incm_expn_cost_rent,
				t_loan_submission_income_expense.loan_sbmssn_incm_expn_cost_installment`).
		Where("loan_sbmssn_id = ?", pid).
		Find(&loansbmssnincmexpn).Error
	if err != nil {
		return LoanSubmissionExpenseMrdmData{}, err
	}
	return loansbmssnincmexpn, nil
}

func (p *LoanSubmissionIncomeExpense) UpdateLoanSubmissionIncomeByLoanSubmissionID(db *gorm.DB, pid uint64) (*LoanSubmissionIncomeExpense, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionIncomeExpense{}).Where("loan_sbmssn_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_incm_expn_salary":               p.LoanSubmissionIncomeExpenseSalary,
			"loan_sbmssn_incm_expn_salary_image":         p.LoanSubmissionIncomeExpenseSalaryImage,
			"loan_sbmssn_incm_expn_salary_partner":       p.LoanSubmissionIncomeExpenseSalaryPartner,
			"loan_sbmssn_incm_expn_salary_partner_image": p.LoanSubmissionIncomeExpenseSalaryPartnerImage,
			"loan_sbmssn_incm_expn_other":                p.LoanSubmissionIncomeExpenseOther,
			"loan_sbmssn_incm_expn_other_image":          p.LoanSubmissionIncomeExpenseSalaryPartnerImage,
			"loan_sbmssn_incm_expn_updated_by":           p.LoanSubmissionIncomeExpenseUpdatedBy,
			"loan_sbmssn_incm_expn_updated_at":           time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionIncomeExpense{}).Where("loan_sbmssn_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionIncomeExpense{}, err
	}
	return p, nil
}

func (p *LoanSubmissionIncomeExpense) UpdateLoanSubmissionExpenseByLoanSubmissionID(db *gorm.DB, pid uint64) (*LoanSubmissionIncomeExpense, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionIncomeExpense{}).Where("loan_sbmssn_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_incm_expn_cost_household":   p.LoanSubmissionIncomeExpenseCostHousehold,
			"loan_sbmssn_incm_expn_cost_education":   p.LoanSubmissionIncomeExpenseCostEducation,
			"loan_sbmssn_incm_expn_cost_utility":     p.LoanSubmissionIncomeExpenseCostUtility,
			"loan_sbmssn_incm_expn_cost_rent":        p.LoanSubmissionIncomeExpenseCostRent,
			"loan_sbmssn_incm_expn_cost_installment": p.LoanSubmissionIncomeExpenseCostInstallment,
			"loan_sbmssn_incm_expn_updated_by":       p.LoanSubmissionIncomeExpenseUpdatedBy,
			"loan_sbmssn_incm_expn_updated_at":       time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionIncomeExpense{}).Where("loan_sbmssn_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionIncomeExpense{}, err
	}
	return p, nil
}
