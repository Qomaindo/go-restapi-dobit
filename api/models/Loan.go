package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Loan struct {
	LoanID           uint64  `gorm:"column:loan_id;primary_key;" json:"loan_id"`
	LoanCode         string  `gorm:"column:loan_code;size:50" json:"loan_code"`
	LoanNum          string  `gorm:"column:loan_num; size:25;" json:"loan_num"`
	LoanAmount       uint64  `gorm:"column:loan_amount" json:"loan_amount"`
	LoanLong         uint64  `gorm:"column:loan_long" json:"loan_long"`
	LoanInterestRate float64 `gorm:"column:loan_interest_rate" json:"loan_interest_rate"`
	// LoanDesc         string    `gorm:"column:loan_desc" json:"loan_desc"`
	LoanDate         string    `gorm:"column:loan_date; size:10;" json:"loan_date"`
	LoanStatus       uint64    `gorm:"column:loan_status" json:"loan_status"`
	LoanCreatedBy    string    `gorm:"column:loan_created_by;size:125" json:"loan_created_by"`
	LoanCreatedAt    time.Time `gorm:"column:loan_created_at;default:CURRENT_TIMESTAMP" json:"loan_created_at"`
	LoanUpdatedBy    string    `gorm:"column:loan_updated_by;size:125" json:"loan_updated_by"`
	LoanUpdatedAt    time.Time `gorm:"column:loan_updated_at;default:CURRENT_TIMESTAMP" json:"loan_updated_at"`
	LoanDeletedBy    string    `gorm:"column:loan_deleted_by;size:125" json:"loan_deleted_by"`
	LoanDeletedAt    time.Time `gorm:"column:loan_deleted_at;default:CURRENT_TIMESTAMP" json:"loan_deleted_at"`
	LoanSubmissionID uint64    `gorm:"column:loan_sbmssn_id" json:"loan_sbmssn_id"`
}

type LoanSubmissionStatus struct {
	LoanSubmissionID     uint64 `gorm:"column:loan_sbmssn_id" json:"loan_sbmssn_id"`
	LoanSubmissionStatus uint64 `gorm:"column:loan_sbmssn_status" json:"loan_sbmssn_status"`
}

type LoanData struct {
	LoanID           uint64    `gorm:"column:loan_id;primary_key;" json:"loan_id"`
	LoanSubmissionID uint64    `gorm:"column:loan_sbmssn_id" json:"loan_sbmssn_id"`
	LoanCode         string    `gorm:"column:loan_code;size:50" json:"loan_code"`
	LoanAmount       uint64    `gorm:"column:loan_amount" json:"loan_amount"`
	LoanNum          string    `gorm:"column:loan_num; size:25;" json:"loan_num"`
	LoanLong         uint64    `gorm:"column:loan_long" json:"loan_long"`
	LoanInterestRate float64   `gorm:"column:loan_interest_rate" json:"loan_interest_rate"`
	LoanDate         string    `gorm:"column:loan_date; size:10;" json:"loan_date"`
	LoanStatus       uint64    `gorm:"column:loan_status" json:"loan_status"`
	LoanCreatedBy    string    `gorm:"column:loan_created_by" json:"loan_created_by"`
	LoanCreatedAt    time.Time `gorm:"column:loan_created_at" json:"loan_created_at"`
	LoanUpdatedBy    string    `gorm:"column:loan_updated_by" json:"loan_updated_by"`
	LoanUpdatedAt    time.Time `gorm:"column:loan_updated_at" json:"loan_updated_at"`
	LoanDeletedBy    string    `gorm:"column:loan_deleted_by" json:"loan_deleted_by"`
	LoanDeletedAt    time.Time `gorm:"column:loan_deleted_at" json:"loan_deleted_at"`
}

type LoanMeet struct {
	LoanID          uint64 `gorm:"column:loan_id;primary_key;" json:"loan_id"`
	LoanDateMeet    string `gorm:"column:loan_meet_date" json:"loan_meet_date"`
	LoanPlaceMeet   string `gorm:"column:loan_meet_place;size:125" json:"loan_meet_place"`
	LoanTimeMeet    string `gorm:"column:loan_meet_time" json:"loan_meet_time"`
	LoanMeetAddress string `gorm:"column:loan_meet_address" json:"loan_meet_address"`
	LoanMeetPurpose string `gorm:"column:loan_meet_purpose" json:"loan_meet_purpose"`
}

type LoanMeetData struct {
	LoanID          uint64 `gorm:"column:loan_id" json:"loan_id"`
	LoanDateMeet    string `gorm:"column:loan_meet_date" json:"loan_meet_date"`
	LoanPlaceMeet   string `gorm:"column:loan_meet_place;size:125" json:"loan_meet_place"`
	LoanTimeMeet    string `gorm:"column:loan_meet_time" json:"loan_meet_time"`
	LoanMeetAddress string `gorm:"column:loan_meet_address" json:"loan_meet_address"`
	LoanMeetPurpose string `gorm:"column:loan_meet_purpose" json:"loan_meet_purpose"`
}

type LoanSignature struct {
	LoanID               uint64 `gorm:"column:loan_id;primary_key;" json:"loan_id"`
	LoanSubmissionID     uint64 `gorm:"column:loan_sbmssn_id;" json:"loan_sbmssn_id"`
	LoanSignatureDate    string `gorm:"column:loan_signature_date" json:"loan_signature_date"`
	LoanSignaturePlace   string `gorm:"column:loan_signature_place;size:125" json:"loan_signature_place"`
	LoanSignatureTime    string `gorm:"column:loan_signature_time" json:"loan_signature_time"`
	LoanSignatureAddress string `gorm:"column:loan_signature_address" json:"loan_signature_address"`
	LoanSignaturePurpose string `gorm:"column:loan_signature_purpose" json:"loan_signature_purpose"`
}

type LoanSignatureData struct {
	LoanID               uint64 `gorm:"column:loan_id" json:"loan_id"`
	LoanSignatureDate    string `gorm:"column:loan_signature_date" json:"loan_signature_date"`
	LoanSignaturePlace   string `gorm:"column:loan_signature_place;size:125" json:"loan_signature_place"`
	LoanSignatureTime    string `gorm:"column:loan_signature_time" json:"loan_signature_time"`
	LoanSignatureAddress string `gorm:"column:loan_signature_address" json:"loan_signature_address"`
	LoanSignaturePurpose string `gorm:"column:loan_signature_purpose" json:"loan_signature_purpose"`
}

type LoanMeetSignatureData struct {
	LoanID               uint64 `gorm:"column:loan_id" json:"loan_id"`
	LoanDateMeet         string `gorm:"column:loan_meet_date" json:"loan_meet_date"`
	LoanPlaceMeet        string `gorm:"column:loan_meet_place;size:125" json:"loan_meet_place"`
	LoanTimeMeet         string `gorm:"column:loan_meet_time" json:"loan_meet_time"`
	LoanMeetAddress      string `gorm:"column:loan_meet_address" json:"loan_meet_address"`
	LoanMeetPurpose      string `gorm:"column:loan_meet_purpose" json:"loan_meet_purpose"`
	LoanSignatureDate    string `gorm:"column:loan_signature_date" json:"loan_signature_date"`
	LoanSignaturePlace   string `gorm:"column:loan_signature_place;size:125" json:"loan_signature_place"`
	LoanSignatureTime    string `gorm:"column:loan_signature_time" json:"loan_signature_time"`
	LoanSignatureAddress string `gorm:"column:loan_signature_address" json:"loan_signature_address"`
	LoanSignaturePurpose string `gorm:"column:loan_signature_purpose" json:"loan_signature_purpose"`
}

type ResponseLoanMeetData struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    *LoanMeetData `json:"data"`
}

type ResponseLoanSignature struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResponseLoanSignatureData struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    *LoanSignatureData `json:"data"`
}
type ResponseLoanSignatureDatas struct {
	Status  int                  `json:"status"`
	Message string               `json:"message"`
	Data    *[]LoanSignatureData `json:"data"`
}

type ResponseLoanMeetSignatureData struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    *LoanMeetSignatureData `json:"data"`
}

type ResponseLoanMeet struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResponseLoans struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    *[]LoanData `json:"data"`
}

type ResponseLoan struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    *LoanData `json:"data"`
}

type ResponseLoanIU struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    *Loan  `json:"data"`
}

type ResponseLoanDel struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResponseGenerateCodeLoan struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    *GenerateCodeLoan `json:"data"`
}

func (Loan) TableName() string {
	return "t_loan"
}

func (LoanSubmissionStatus) TableName() string {
	return "t_loan_submission"
}

func (LoanData) TableName() string {
	return "t_loan"
}

func (LoanMeet) TableName() string {
	return "t_loan"
}

func (LoanMeetData) TableName() string {
	return "t_loan"
}

func (LoanSignature) TableName() string {
	return "t_loan"
}

func (LoanSignatureData) TableName() string {
	return "t_loan"
}

func (LoanMeetSignatureData) TableName() string {
	return "t_loan"
}

func (GenerateCodeLoan) TableName() string {
	return "t_loan"
}

func (p *Loan) Prepare() {
	p.LoanSubmissionID = p.LoanSubmissionID
	p.LoanCode = html.EscapeString(strings.TrimSpace(p.LoanCode))
	p.LoanAmount = p.LoanAmount
	p.LoanLong = p.LoanLong
	p.LoanInterestRate = p.LoanInterestRate
	p.LoanStatus = p.LoanStatus
	p.LoanCreatedBy = p.LoanCreatedBy
	p.LoanCreatedAt = p.LoanCreatedAt
	p.LoanUpdatedBy = p.LoanUpdatedBy
	p.LoanUpdatedAt = p.LoanUpdatedAt
	p.LoanDeletedBy = p.LoanDeletedBy
	p.LoanDeletedAt = time.Now()
}

func (p *LoanMeet) Prepare() {
	p.LoanDateMeet = p.LoanDateMeet
	p.LoanPlaceMeet = html.EscapeString(strings.TrimSpace(p.LoanPlaceMeet))
	p.LoanTimeMeet = p.LoanTimeMeet
}

func (p *LoanSignature) Prepare() {
	p.LoanSignatureDate = p.LoanSignatureDate
	p.LoanSignaturePlace = html.EscapeString(strings.TrimSpace(p.LoanSignaturePlace))
	p.LoanSignatureTime = p.LoanSignatureTime
}

func (p *LoanMeet) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.LoanDateMeet == "" {
			return errors.New("Required Loan Meet date")
		}
		if p.LoanPlaceMeet == "" {
			return errors.New("Required Loan Code")
		}
		if p.LoanTimeMeet == "" {
			return errors.New("Required Loan Amount")
		}
		return nil
	}
}

func (p *LoanSignature) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.LoanSignatureDate == "" {
			return errors.New("Required Loan Signature Date")
		}
		if p.LoanSignaturePlace == "" {
			return errors.New("Required Loan Code")
		}
		if p.LoanSignatureTime == "" {
			return errors.New("Required Loan Amount")
		}
		return nil
	}
}

func (p *Loan) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.LoanSubmissionID == 0 {
			return errors.New("Required Loan Submission ID")
		}
		if p.LoanCode == "" {
			return errors.New("Required Loan Code")
		}
		if p.LoanAmount == 0 {
			return errors.New("Required Loan Amount")
		}
		if p.LoanLong == 0 {
			return errors.New("Required Loan Long")
		}
		if p.LoanInterestRate == 0 {
			return errors.New("Required Loan Interest Rate")
		}
		return nil
	}
}

func (p *Loan) SaveLoan(db *gorm.DB) (*Loan, error) {
	var err error
	err = db.Debug().Model(&Loan{}).Create(&p).Error
	if err != nil {
		return &Loan{}, err
	}
	return p, nil
}

func (p *Loan) FindAllLoan(db *gorm.DB) (*[]LoanData, error) {
	var err error
	loan := []LoanData{}
	err = db.Debug().Model(&Loan{}).Limit(100).
		Select(`t_loan.loan_id,
				t_loan.loan_code,
				t_loan.loan_amount,
				t_loan.loan_long,
				t_loan.loan_interest_rate,
				t_loan.loan_status,
				t_loan.loan_meet_date,
				t_loan.loan_meet_place,
				t_loan.loan_meet_time,
				t_loan.loan_signature_date,
				t_loan.loan_signature_time,
				t_loan.loan_signature_place,
				t_loan.loan_date,
				t_loan.loan_created_by,
				t_loan.loan_created_at at time zone 'Asia/Jakarta' as loan_created_at,
				t_loan.loan_updated_by,
				t_loan.loan_updated_at,
				t_loan.loan_deleted_by,
				t_loan.loan_deleted_at`).
		Find(&loan).Error
	if err != nil {
		return &[]LoanData{}, err
	}
	return &loan, nil
}

func (p *Loan) FindDataLoanByStatus(db *gorm.DB, pid string) (*[]LoanData, error) {
	var err error
	loan := []LoanData{}
	err = db.Debug().Model(&LoanData{}).Limit(100).
		Select(`t_loan.loan_id,
				t_loan.loan_sbmssn_id,
				t_loan.loan_code,
				t_loan.loan_amount,
				t_loan.loan_long,
				t_loan.loan_interest_rate,
				t_loan.loan_status,
				t_loan.loan_meet_date,
				t_loan.loan_meet_place,
				t_loan.loan_meet_time,
				t_loan.loan_signature_date,
				t_loan.loan_signature_place,
				t_loan.loan_signature_time,
				t_loan.loan_date,
				t_loan.loan_created_by,
				t_loan.loan_created_at at time zone 'Asia/Jakarta' as loan_created_at,
				t_loan.loan_updated_by,
				t_loan.loan_updated_at,
				t_loan.loan_deleted_by,
				t_loan.loan_deleted_at`).
		Joins("JOIN t_loan_submission on t_loan.loan_sbmssn_id=t_loan_submission.loan_sbmssn_id").
		Joins("JOIN m_customer_user on t_loan_submission.cust_user_id=m_customer_user.cust_user_id").
		Where("cust_user_phone = ? AND loan_status = 13", pid).
		// Where("loan_status = 13").
		Find(&loan).Error
	if err != nil {
		return &[]LoanData{}, err
	}
	return &loan, nil
}

func (p *Loan) FindAllDataLoanByStatus(db *gorm.DB) (*[]LoanData, error) {
	var err error
	loan := []LoanData{}
	err = db.Debug().Model(&LoanData{}).Limit(100).
		Select(`t_loan.loan_id,
				t_loan.loan_sbmssn_id,
				t_loan.loan_code,
				t_loan.loan_amount,
				t_loan.loan_long,
				t_loan.loan_interest_rate,
				t_loan.loan_status,
				t_loan.loan_meet_date,
				t_loan.loan_meet_place,
				t_loan.loan_meet_time,
				t_loan.loan_signature_date,
				t_loan.loan_signature_place,
				t_loan.loan_signature_time,
				t_loan.loan_date,
				t_loan.loan_created_by,
				t_loan.loan_created_at at time zone 'Asia/Jakarta' as loan_created_at,
				t_loan.loan_updated_by,
				t_loan.loan_updated_at,
				t_loan.loan_deleted_by,
				t_loan.loan_deleted_at`).
		Joins("JOIN t_loan_submission on t_loan.loan_sbmssn_id=t_loan_submission.loan_sbmssn_id").
		Joins("JOIN m_customer_user on t_loan_submission.cust_user_id=m_customer_user.cust_user_id").
		// Where("cust_user_phone = ? AND loan_status = 13", pid).
		Where("loan_status = 13").
		Find(&loan).Error
	if err != nil {
		return &[]LoanData{}, err
	}
	return &loan, nil
}

func (p *Loan) FindAllLoanStatus(db *gorm.DB) (*[]LoanData, error) {
	var err error
	loan := []LoanData{}
	err = db.Debug().Model(&LoanData{}).Limit(100).
		Select(`t_loan.loan_id,
				t_loan.loan_sbmssn_id,
				t_loan.loan_code,
				t_loan.loan_amount,
				t_loan.loan_long,
				t_loan.loan_interest_rate,
				t_loan.loan_status,
				t_loan.loan_meet_date,
				t_loan.loan_meet_place,
				t_loan.loan_meet_time,
				t_loan.loan_signature_date,
				t_loan.loan_signature_place,
				t_loan.loan_signature_time,
				t_loan.loan_date,
				t_loan.loan_created_by,
				t_loan.loan_created_at at time zone 'Asia/Jakarta' as loan_created_at,
				t_loan.loan_updated_by,
				t_loan.loan_updated_at,
				t_loan.loan_deleted_by,
				t_loan.loan_deleted_at`).
		Where("loan_status = 9 OR loan_status = 10 OR loan_status = 11 OR loan_status = 12").
		Find(&loan).Error
	if err != nil {
		return &[]LoanData{}, err
	}
	return &loan, nil
}

func (p *Loan) FindLoanStatusByID(db *gorm.DB, pid uint64, status uint64) (*Loan, error) {
	var err error
	loan := Loan{}
	err = db.Debug().Model(&Loan{}).Limit(100).
		Select(`t_loan.loan_id,
				t_loan.loan_sbmssn_id,
				t_loan.loan_code,
				t_loan.loan_amount,
				t_loan.loan_long,
				t_loan.loan_interest_rate,
				t_loan.loan_status,
				t_loan.loan_meet_date,
				t_loan.loan_meet_place,
				t_loan.loan_meet_time,
				t_loan.loan_signature_date,
				t_loan.loan_signature_time,
				t_loan.loan_signature_place,
				t_loan.loan_date,
				t_loan.loan_created_by,
				t_loan.loan_created_at at time zone 'Asia/Jakarta' as loan_created_at,
				t_loan.loan_updated_by,
				t_loan.loan_updated_at,
				t_loan.loan_deleted_by,
				t_loan.loan_deleted_at`).
		Where("loan_id = ? AND loan_status = ?", pid, status).
		Find(&loan).Error
	if err != nil {
		return &Loan{}, err
	}
	return &loan, nil
}

func (p *Loan) FindLoanStatusBySbmssnID(db *gorm.DB, pid uint64) (*Loan, error) {
	var err error
	loan := Loan{}
	err = db.Debug().Model(&Loan{}).Limit(100).
		Select(`t_loan.loan_id,
				t_loan.loan_sbmssn_id,
				t_loan.loan_code,
				t_loan.loan_amount,
				t_loan.loan_long,
				t_loan.loan_interest_rate,
				t_loan.loan_status,
				t_loan.loan_meet_date,
				t_loan.loan_meet_place,
				t_loan.loan_meet_time,
				t_loan.loan_signature_date,
				t_loan.loan_signature_time,
				t_loan.loan_signature_place,
				t_loan.loan_date,
				t_loan.loan_created_by,
				t_loan.loan_created_at at time zone 'Asia/Jakarta' as loan_created_at,
				t_loan.loan_updated_by,
				t_loan.loan_updated_at,
				t_loan.loan_deleted_by,
				t_loan.loan_deleted_at`).
		Where("loan_sbmssn_id = ?", pid).
		Find(&loan).Error
	if err != nil {
		return &Loan{}, err
	}
	return &loan, nil
}
func (p *Loan) FindLoanDataByID(db *gorm.DB, pid uint64) (*Loan, error) {
	var err error
	err = db.Debug().Model(&Loan{}).Where("loan_id = ?", pid).Take(&p).Error
	if err != nil {
		return &Loan{}, err
	}
	return p, nil
}

func (p *Loan) FindLoanByID(db *gorm.DB, pid uint64) (*LoanData, error) {
	var err error
	loan := LoanData{}
	err = db.Debug().Model(&LoanData{}).Limit(100).
		Select(`t_loan.loan_id,
				t_loan.loan_sbmssn_id,
				t_loan.loan_code,
				t_loan.loan_amount,
				t_loan.loan_long,
				t_loan.loan_interest_rate,
				t_loan.loan_status,
				t_loan.loan_meet_date,
				t_loan.loan_meet_place,
				t_loan.loan_meet_time,
				t_loan.loan_signature_date,
				t_loan.loan_signature_time,
				t_loan.loan_signature_place,
				t_loan.loan_date,
				t_loan.loan_created_by,
				t_loan.loan_created_at at time zone 'Asia/Jakarta' as loan_created_at,
				t_loan.loan_updated_by,
				t_loan.loan_updated_at,
				t_loan.loan_deleted_by,
				t_loan.loan_deleted_at`).
		Where("loan_id = ?", pid).
		Find(&loan).Error
	if err != nil {
		return &LoanData{}, err
	}
	return &loan, nil
}

func (p *Loan) UpdateLoan(db *gorm.DB, pid uint64) (*Loan, error) {
	var err error
	db = db.Debug().Model(&Loan{}).Where("loan_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_id":     p.LoanSubmissionID,
			"loan_code":          p.LoanCode,
			"loan_amount":        p.LoanAmount,
			"loan_long":          p.LoanLong,
			"loan_interest_rate": p.LoanInterestRate,
			"loan_status":        p.LoanStatus,
			"loan_updated_by":    p.LoanUpdatedBy,
			"loan_updated_at":    time.Now,
		},
	)
	err = db.Debug().Model(&Loan{}).Where("loan_id = ?", pid).Error
	if err != nil {
		return &Loan{}, err
	}
	return p, nil
}

func (p *Loan) UpdateLoanStatus(db *gorm.DB, pid uint64) (*Loan, error) {
	var err error
	db = db.Debug().Model(&Loan{}).Where("loan_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_status":     p.LoanStatus,
			"loan_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&Loan{}).Where("loan_id = ?", pid).Error
	if err != nil {
		return &Loan{}, err
	}
	return p, nil
}

func (p *Loan) UpdateLoanConfirmStatus(db *gorm.DB, pid uint64) (*Loan, error) {
	var err error
	db = db.Debug().Model(&Loan{}).Where("loan_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			// "loan_desc":       p.LoanDesc,
			"loan_status":     p.LoanStatus,
			"loan_updated_by": p.LoanUpdatedBy,
			"loan_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&Loan{}).Where("loan_id = ?", pid).Error
	if err != nil {
		return &Loan{}, err
	}
	return p, nil
}

func (p *LoanSubmissionStatus) UpdateLoanSubmissionStatus(db *gorm.DB, pid uint64) (*LoanSubmissionStatus, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionStatus{}).Where("loan_sbmssn_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_status":     p.LoanSubmissionStatus,
			"loan_sbmssn_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionStatus{}).Where("loan_sbmssn_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionStatus{}, err
	}
	return p, nil
}

func (p *Loan) DeleteHard(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&Loan{}).Where("loan_id = ?", pid).Take(&Loan{}).Delete(&Loan{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Loan not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// Additional ================================================================

// =====================================================================================================
// Insert Date Meet

func (p *LoanMeet) UpdateLoanMeet(db *gorm.DB, pid uint64) (*LoanMeet, error) {
	var err error
	db = db.Debug().Model(&LoanMeet{}).Where("loan_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_meet_date":    p.LoanDateMeet,
			"loan_meet_place":   p.LoanPlaceMeet,
			"loan_meet_time":    p.LoanTimeMeet,
			"loan_meet_address": p.LoanMeetAddress,
			"loan_meet_purpose": p.LoanMeetPurpose,
		},
	)
	err = db.Debug().Model(&LoanMeet{}).Where("loan_id = ?", pid).Error
	if err != nil {
		return &LoanMeet{}, err
	}
	return p, nil
}

func (p *LoanMeet) FindLoanMeetByID(db *gorm.DB, pid string) (*LoanMeetData, error) {
	var err error
	loanmeet := LoanMeetData{}
	err = db.Debug().Model(&LoanMeetData{}).Limit(100).
		Select(`t_loan.loan_id,
				t_loan.loan_meet_date,
				t_loan.loan_meet_place,
				t_loan.loan_meet_time,
				t_loan.loan_meet_address,
				t_loan.loan_meet_purpose`).
		Joins("JOIN t_loan_submission on t_loan.loan_sbmssn_id=t_loan_submission.loan_sbmssn_id").
		Joins("JOIN m_customer_user on t_loan_submission.cust_user_id=m_customer_user.cust_user_id").
		Where("cust_user_phone = ?", pid).
		Find(&loanmeet).Error
	if err != nil {
		return &LoanMeetData{}, err
	}
	return &loanmeet, nil
}

// =====================================================================================================
// Insert Date Signature

func (p *LoanSignature) UpdateLoanSignature(db *gorm.DB, pid uint64) (*LoanSignature, error) {
	var err error
	db = db.Debug().Model(&LoanSignature{}).Where("loan_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_signature_date":    p.LoanSignatureDate,
			"loan_signature_place":   p.LoanSignaturePlace,
			"loan_signature_time":    p.LoanSignatureTime,
			"loan_signature_address": p.LoanSignatureAddress,
			"loan_signature_purpose": p.LoanSignaturePurpose,
		},
	)
	err = db.Debug().Model(&LoanSignature{}).Where("loan_id = ?", pid).Error
	if err != nil {
		return &LoanSignature{}, err
	}
	return p, nil
}

func (p *LoanSignature) FindLoanSignatureByPhone(db *gorm.DB, pid string) (*[]LoanSignatureData, error) {
	var err error
	loansignature := []LoanSignatureData{}
	err = db.Debug().Model(&LoanSignatureData{}).Limit(100).
		Select(`t_loan.loan_id,
				t_loan.loan_signature_date,
				t_loan.loan_signature_place,
				t_loan.loan_signature_time,
				t_loan.loan_signature_address,
				t_loan.loan_signature_purpose`).
		Joins("JOIN t_loan_submission on t_loan.loan_sbmssn_id=t_loan_submission.loan_sbmssn_id").
		Joins("JOIN m_customer_user on t_loan_submission.cust_user_id=m_customer_user.cust_user_id").
		Where("cust_user_phone = ?", pid).
		Find(&loansignature).Error
	if err != nil {
		return &[]LoanSignatureData{}, err
	}
	return &loansignature, nil
}

func (p *LoanSignature) FindLoanSignatureByID(db *gorm.DB, pid string) (*LoanSignatureData, error) {
	var err error
	loansignature := LoanSignatureData{}
	err = db.Debug().Model(&LoanSignatureData{}).Limit(100).
		Select(`t_loan.loan_id,
				t_loan.loan_signature_date,
				t_loan.loan_signature_place,
				t_loan.loan_signature_time,
				t_loan.loan_signature_address,
				t_loan.loan_signature_purpose`).
		Joins("JOIN t_loan_submission on t_loan.loan_sbmssn_id=t_loan_submission.loan_sbmssn_id").
		Joins("JOIN m_customer_user on t_loan_submission.cust_user_id=m_customer_user.cust_user_id").
		Where("loan_id = ?", pid).
		Find(&loansignature).Error
	if err != nil {
		return &LoanSignatureData{}, err
	}
	return &loansignature, nil
}

// Get All Jadwal Pertemuan By User ID
// ========================================================================================================

func (p *Loan) FindLoanSchaduleByIDUser(db *gorm.DB, pid uint64, uid uint64) (*LoanMeetSignatureData, error) {
	var err error
	loan := LoanMeetSignatureData{}
	err = db.Debug().Model(&LoanMeetSignatureData{}).Limit(100).
		Select(`t_loan.loan_id,
				t_loan.loan_sbmssn_id,
				t_loan_submission.mitra_user_id,
				t_loan.loan_status,
				t_loan.loan_meet_date,
				t_loan.loan_meet_place,
				t_loan.loan_meet_time,
				t_loan.loan_meet_address,
				t_loan.loan_meet_purpose,
				t_loan.loan_signature_date,
				t_loan.loan_signature_time,
				t_loan.loan_signature_place,
				t_loan.loan_signature_address,
				t_loan.loan_signature_purpose,
				t_loan.loan_created_by,
				t_loan.loan_created_at at time zone 'Asia/Jakarta' as loan_created_at,
				t_loan.loan_updated_by,
				t_loan.loan_updated_at,
				t_loan.loan_deleted_by,
				t_loan.loan_deleted_at`).
		Joins("JOIN t_loan_submission on t_loan.loan_sbmssn_id=t_loan_submission.loan_sbmssn_id").
		Where("loan_id = ? AND mitra_user_id = ?", pid, uid).
		Find(&loan).Error
	if err != nil {
		return &LoanMeetSignatureData{}, err
	}
	return &loan, nil
}

// GENERATE CODE
// ===================================================================

type GenerateCodeLoan struct {
	LoanLastNum   string    `gorm:"column:loan_last_num" json:"loan_last_num"`
	LoanCreatedAt time.Time `gorm:"column:loan_created_at" json:"loan_created_at"`
	LoanNum       string    `gorm:"column:loan_num" json:"loan_num"`
}

func (p *Loan) GenerateCode(db *gorm.DB) (*GenerateCodeLoan, error) {
	var err error
	loan := GenerateCodeLoan{}
	err = db.Debug().Model(&GenerateCodeLoan{}).Limit(1).
		Select(`RIGHT(t_loan.loan_num, 5) as loan_last_num, loan_created_at`).
		Order("loan_created_at desc").
		Find(&loan).
		Take(&loan).Error
	if err != nil {
		return &GenerateCodeLoan{}, err
	}
	return &loan, nil
}
