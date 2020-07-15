package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type LoanSubmissionAccountAnalysis struct {
	LoanSubmissionAccountAnalysisID               uint64     `gorm:"column:loan_sbmssn_acnt_anlys_id;primary_key;" json:"loan_sbmssn_acnt_anlys_id"`
	LoanSubmissionID                              uint64     `gorm:"column:loan_sbmssn_id" json:"loan_sbmssn_id"`
	LoanSubmissionAccountAnalysisName             string     `gorm:"column:loan_sbmssn_acnt_anlys_name;size:125" json:"loan_sbmssn_acnt_anlys_name"`
	LoanSubmissionAccountAnalysisCode             string     `gorm:"column:loan_sbmssn_acnt_anlys_code;size:125" json:"loan_sbmssn_acnt_anlys_code"`
	LoanSubmissionAccountAnalysisReceiver         string     `gorm:"column:loan_sbmssn_acnt_anlys_receiver;size:125" json:"loan_sbmssn_acnt_anlys_receiver"`
	LoanSubmissionAccountAnalysisType             string     `gorm:"column:loan_sbmssn_acnt_anlys_type;size:125" json:"loan_sbmssn_acnt_anlys_type"`
	LoanSubmissionAccountAnalysisFacilities       string     `gorm:"column:loan_sbmssn_acnt_anlys_facilities;size:125" json:"loan_sbmssn_acnt_anlys_facilities"`
	LoanSubmissionAccountSalariesRecorded         uint64     `gorm:"column:loan_sbmssn_acnt_salaries_recorded" json:"loan_sbmssn_acnt_salaries_recorded"`
	LoanSubmissionAccountSalariesRecordedDocument string     `gorm:"column:loan_sbmssn_acnt_salaries_recorded_document;size:255" json:"loan_sbmssn_acnt_salaries_recorded_document"`
	LoanSubmissionAccountDescription              string     `gorm:"column:loan_sbmssn_acnt_desc;size:255" json:"loan_sbmssn_acnt_desc"`
	LoanSubmissionAccountStatus                   uint64     `gorm:"column:loan_sbmssn_acnt_status" json:"loan_sbmssn_acnt_status"`
	LoanSubmissionAccountCreatedBy                string     `gorm:"column:loan_sbmssn_acnt_created_by;size:125" json:"loan_sbmssn_acnt_created_by"`
	LoanSubmissionAccountCreatedAt                time.Time  `gorm:"column:loan_sbmssn_acnt_created_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_acnt_created_at"`
	LoanSubmissionAccountUpdatedBy                string     `gorm:"column:loan_sbmssn_acnt_updated_by;size:125" json:"loan_sbmssn_acnt_updated_by"`
	LoanSubmissionAccountUpdatedAt                *time.Time `gorm:"column:loan_sbmssn_acnt_updated_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_acnt_updated_at"`
	LoanSubmissionAccountDeletedBy                string     `gorm:"column:loan_sbmssn_acnt_deleted_by;size:125" json:"loan_sbmssn_acnt_deleted_by"`
	LoanSubmissionAccountDeletedAt                *time.Time `gorm:"column:loan_sbmssn_acnt_deleted_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_acnt_deleted_at"`
}

type LoanSubmissionAccountAnalysisData struct {
	LoanSubmissionAccountAnalysisID               uint64     `gorm:"column:loan_sbmssn_acnt_anlys_id;primary_key;" json:"loan_sbmssn_acnt_anlys_id"`
	LoanSubmissionID                              uint64     `gorm:"column:loan_sbmssn_id" json:"loan_sbmssn_id"`
	LoanSubmissionAccountAnalysisName             string     `gorm:"column:loan_sbmssn_acnt_anlys_name" json:"loan_sbmssn_acnt_anlys_name"`
	LoanSubmissionAccountAnalysisCode             string     `gorm:"column:loan_sbmssn_acnt_anlys_code" json:"loan_sbmssn_acnt_anlys_code"`
	LoanSubmissionAccountAnalysisReceiver         string     `gorm:"column:loan_sbmssn_acnt_anlys_receiver" json:"loan_sbmssn_acnt_anlys_receiver"`
	LoanSubmissionAccountAnalysisType             string     `gorm:"column:loan_sbmssn_acnt_anlys_type" json:"loan_sbmssn_acnt_anlys_type"`
	LoanSubmissionAccountAnalysisFacilities       string     `gorm:"column:loan_sbmssn_acnt_anlys_facilities" json:"loan_sbmssn_acnt_anlys_facilities"`
	LoanSubmissionAccountSalariesRecorded         uint64     `gorm:"column:loan_sbmssn_acnt_salaries_recorded" json:"loan_sbmssn_acnt_salaries_recorded"`
	LoanSubmissionAccountSalariesRecordedDocument string     `gorm:"column:loan_sbmssn_acnt_salaries_recorded_document" json:"loan_sbmssn_acnt_salaries_recorded_document"`
	LoanSubmissionAccountDescription              string     `gorm:"column:loan_sbmssn_acnt_desc" json:"loan_sbmssn_acnt_desc"`
	LoanSubmissionAccountStatus                   uint64     `gorm:"column:loan_sbmssn_acnt_status" json:"loan_sbmssn_acnt_status"`
	LoanSubmissionAccountCreatedBy                string     `gorm:"column:loan_sbmssn_acnt_created_by" json:"loan_sbmssn_acnt_created_by"`
	LoanSubmissionAccountCreatedAt                time.Time  `gorm:"column:loan_sbmssn_acnt_created_at" json:"loan_sbmssn_acnt_created_at"`
	LoanSubmissionAccountUpdatedBy                string     `gorm:"column:loan_sbmssn_acnt_updated_by" json:"loan_sbmssn_acnt_updated_by"`
	LoanSubmissionAccountUpdatedAt                *time.Time `gorm:"column:loan_sbmssn_acnt_updated_at" json:"loan_sbmssn_acnt_updated_at"`
	LoanSubmissionAccountDeletedBy                string     `gorm:"column:loan_sbmssn_acnt_deleted_by" json:"loan_sbmssn_acnt_deleted_by"`
	LoanSubmissionAccountDeletedAt                *time.Time `gorm:"column:loan_sbmssn_acnt_deleted_at" json:"loan_sbmssn_acnt_deleted_at"`
}

type ResponseLoanSubmissionAccountAnalysiss struct {
	Status  int                                  `json:"status"`
	Message string                               `json:"message"`
	Data    *[]LoanSubmissionAccountAnalysisData `json:"data"`
}

type ResponseLoanSubmissionAccountAnalysis struct {
	Status  int                                `json:"status"`
	Message string                             `json:"message"`
	Data    *LoanSubmissionAccountAnalysisData `json:"data"`
}

type ResponseLoanSubmissionAccountAnalysisCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (LoanSubmissionAccountAnalysis) TableName() string {
	return "t_loan_submission_account_analysis"
}

func (LoanSubmissionAccountAnalysisData) TableName() string {
	return "t_loan_submission_account_analysis"
}

func (p *LoanSubmissionAccountAnalysis) Prepare() {
	p.LoanSubmissionID = p.LoanSubmissionID
	p.LoanSubmissionAccountAnalysisName = html.EscapeString(strings.TrimSpace(p.LoanSubmissionAccountAnalysisName))
	p.LoanSubmissionAccountAnalysisCode = html.EscapeString(strings.TrimSpace(p.LoanSubmissionAccountAnalysisCode))
	p.LoanSubmissionAccountAnalysisReceiver = html.EscapeString(strings.TrimSpace(p.LoanSubmissionAccountAnalysisReceiver))
	p.LoanSubmissionAccountAnalysisType = html.EscapeString(strings.TrimSpace(p.LoanSubmissionAccountAnalysisType))
	p.LoanSubmissionAccountAnalysisFacilities = html.EscapeString(strings.TrimSpace(p.LoanSubmissionAccountAnalysisFacilities))
	p.LoanSubmissionAccountSalariesRecorded = p.LoanSubmissionAccountSalariesRecorded
	p.LoanSubmissionAccountSalariesRecordedDocument = html.EscapeString(strings.TrimSpace(p.LoanSubmissionAccountSalariesRecordedDocument))
	p.LoanSubmissionAccountDescription = html.EscapeString(strings.TrimSpace(p.LoanSubmissionAccountDescription))
	p.LoanSubmissionAccountStatus = p.LoanSubmissionAccountStatus
	p.LoanSubmissionAccountCreatedBy = p.LoanSubmissionAccountCreatedBy
	p.LoanSubmissionAccountCreatedAt = time.Now()
	p.LoanSubmissionAccountUpdatedBy = p.LoanSubmissionAccountUpdatedBy
	p.LoanSubmissionAccountUpdatedAt = p.LoanSubmissionAccountUpdatedAt
	p.LoanSubmissionAccountDeletedBy = p.LoanSubmissionAccountDeletedBy
	p.LoanSubmissionAccountDeletedAt = p.LoanSubmissionAccountDeletedAt
}

func (p *LoanSubmissionAccountAnalysis) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.LoanSubmissionID == 0 {
			return errors.New("Required Loan Submission ID")
		}
		if p.LoanSubmissionAccountAnalysisName == "" {
			return errors.New("Required Mitra accountanalysis Name")
		}
		if p.LoanSubmissionAccountAnalysisCode == "" {
			return errors.New("Required Mitra accountanalysis Type Customer")
		}
		if p.LoanSubmissionAccountAnalysisReceiver == "" {
			return errors.New("Required Receiver")
		}
		if p.LoanSubmissionAccountAnalysisType == "" {
			return errors.New("Required Type")
		}
		if p.LoanSubmissionAccountAnalysisFacilities == "" {
			return errors.New("Required Facilities")
		}
		if p.LoanSubmissionAccountSalariesRecorded == 0 {
			return errors.New("Required Salaries Recorded")
		}
		if p.LoanSubmissionAccountSalariesRecordedDocument == "" {
			return errors.New("Required Salaries Recorded Document")
		}
		if p.LoanSubmissionAccountDescription == "" {
			return errors.New("Required Description")
		}
		return nil
	}
}

func (p *LoanSubmissionAccountAnalysis) SaveLoanSubmissionAccountAnalysis(db *gorm.DB) (*LoanSubmissionAccountAnalysis, error) {
	var err error
	err = db.Debug().Model(&LoanSubmissionAccountAnalysis{}).Create(&p).Error
	if err != nil {
		return &LoanSubmissionAccountAnalysis{}, err
	}
	return p, nil
}

func (p *LoanSubmissionAccountAnalysis) FindAllLoanSubmissionAccountAnalysis(db *gorm.DB) (*[]LoanSubmissionAccountAnalysisData, error) {
	var err error
	accountanalysis := []LoanSubmissionAccountAnalysisData{}
	err = db.Debug().Model(&LoanSubmissionAccountAnalysisData{}).Limit(100).
		Select(`t_loan_submission_account_analysis.loan_sbmssn_acnt_anlys_id,
				t_loan_submission.loan_sbmssn_id,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_anlys_name,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_anlys_code,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_anlys_receiver,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_anlys_type,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_anlys_facilities,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_salaries_recorded,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_salaries_recorded_document,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_desc,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_status,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_created_by,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_created_at,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_updated_by,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_updated_at,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_deleted_by,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_deleted_at
				`).
		Joins("JOIN t_loan_submission on t_loan_submission_account_analysis.loan_sbmssn_id=t_loan_submission.loan_sbmssn_id").
		Find(&accountanalysis).Error
	if err != nil {
		return &[]LoanSubmissionAccountAnalysisData{}, err
	}
	return &accountanalysis, nil
}

func (p *LoanSubmissionAccountAnalysis) FindLoanSubmissionAccountAnalysisDataByID(db *gorm.DB, pid uint64) (*LoanSubmissionAccountAnalysis, error) {
	var err error
	err = db.Debug().Model(&LoanSubmissionAccountAnalysis{}).Where("loan_sbmssn_acnt_anlys_id = ?", pid).Take(&p).Error
	if err != nil {
		return &LoanSubmissionAccountAnalysis{}, err
	}
	return p, nil
}

func (p *LoanSubmissionAccountAnalysis) FindLoanSubmissionAccountAnalysisByID(db *gorm.DB, pid uint64) (*LoanSubmissionAccountAnalysisData, error) {
	var err error
	accountanalysis := LoanSubmissionAccountAnalysisData{}
	err = db.Debug().Model(&LoanSubmissionAccountAnalysisData{}).Limit(100).
		Select(`t_loan_submission_account_analysis.loan_sbmssn_acnt_anlys_id,
				t_loan_submission.loan_sbmssn_id,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_anlys_name,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_anlys_code,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_anlys_receiver,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_anlys_type,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_anlys_facilities,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_salaries_recorded,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_salaries_recorded_document,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_desc,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_status,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_created_by,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_created_at,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_updated_by,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_updated_at,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_deleted_by,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_deleted_at
				`).
		Joins("JOIN t_loan_submission on t_loan_submission_account_analysis.loan_sbmssn_id=t_loan_submission.loan_sbmssn_id").
		Where("loan_sbmssn_acnt_anlys_id = ?", pid).
		Find(&accountanalysis).Error
	if err != nil {
		return &LoanSubmissionAccountAnalysisData{}, err
	}
	return &accountanalysis, nil
}

func (p *LoanSubmissionAccountAnalysis) UpdateLoanSubmissionAccountAnalysis(db *gorm.DB, pid uint64) (*LoanSubmissionAccountAnalysis, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionAccountAnalysis{}).Where("loan_sbmssn_acnt_anlys_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_acnt_anlys_name":                 p.LoanSubmissionAccountAnalysisName,
			"loan_sbmssn_acnt_anlys_code":                 p.LoanSubmissionAccountAnalysisCode,
			"loan_sbmssn_acnt_anlys_receiver":             p.LoanSubmissionAccountAnalysisReceiver,
			"loan_sbmssn_acnt_anlys_type":                 p.LoanSubmissionAccountAnalysisType,
			"loan_sbmssn_acnt_anlys_facilities":           p.LoanSubmissionAccountAnalysisFacilities,
			"loan_sbmssn_acnt_salaries_recorded":          p.LoanSubmissionAccountSalariesRecorded,
			"loan_sbmssn_acnt_salaries_recorded_document": p.LoanSubmissionAccountSalariesRecordedDocument,
			"loan_sbmssn_acnt_desc":                       p.LoanSubmissionAccountDescription,
			"loan_sbmssn_acnt_updated_by":                 p.LoanSubmissionAccountUpdatedBy,
			"loan_sbmssn_acnt_updated_at":                 time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionAccountAnalysis{}).Where("loan_sbmssn_acnt_anlys_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionAccountAnalysis{}, err
	}
	return p, nil
}

func (p *LoanSubmissionAccountAnalysis) UpdateLoanSubmissionAccountAnalysisDelete(db *gorm.DB, pid uint64) (*LoanSubmissionAccountAnalysis, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionAccountAnalysis{}).Where("loan_sbmssn_acnt_anlys_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_acnt_status":     p.LoanSubmissionAccountStatus,
			"loan_sbmssn_acnt_deleted_by": p.LoanSubmissionAccountDeletedBy,
			"loan_sbmssn_acnt_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionAccountAnalysis{}).Where("loan_sbmssn_acnt_anlys_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionAccountAnalysis{}, err
	}
	return p, nil
}

func (p *LoanSubmissionAccountAnalysis) DeleteLoanSubmissionAccountAnalysis(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&LoanSubmissionAccountAnalysis{}).Where("loan_sbmssn_acnt_anlys_id = ?", pid).Take(&LoanSubmissionAccountAnalysis{}).Delete(&LoanSubmissionAccountAnalysis{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Mitra account analysis not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// ADDITIONAL MEMORANDUM
// ===================================================================================================================================

type LoanSubmissionAccountAnalysisMrdmData struct {
	LoanSubmissionAccountAnalysisID               uint64 `gorm:"column:loan_sbmssn_acnt_anlys_id" json:"loan_sbmssn_acnt_anlys_id"`
	LoanSubmissionID                              uint64 `gorm:"column:loan_sbmssn_id" json:"loan_sbmssn_id"`
	LoanSubmissionAccountAnalysisName             string `gorm:"column:loan_sbmssn_acnt_anlys_name" json:"loan_sbmssn_acnt_anlys_name"`
	LoanSubmissionAccountAnalysisCode             string `gorm:"column:loan_sbmssn_acnt_anlys_code" json:"loan_sbmssn_acnt_anlys_code"`
	LoanSubmissionAccountAnalysisReceiver         string `gorm:"column:loan_sbmssn_acnt_anlys_receiver" json:"loan_sbmssn_acnt_anlys_receiver"`
	LoanSubmissionAccountAnalysisType             string `gorm:"column:loan_sbmssn_acnt_anlys_type" json:"loan_sbmssn_acnt_anlys_type"`
	LoanSubmissionAccountAnalysisFacilities       string `gorm:"column:loan_sbmssn_acnt_anlys_facilities" json:"loan_sbmssn_acnt_anlys_facilities"`
	LoanSubmissionAccountSalariesRecorded         uint64 `gorm:"column:loan_sbmssn_acnt_salaries_recorded" json:"loan_sbmssn_acnt_salaries_recorded"`
	LoanSubmissionAccountSalariesRecordedDocument string `gorm:"column:loan_sbmssn_acnt_salaries_recorded_document" json:"loan_sbmssn_acnt_salaries_recorded_document"`
	LoanSubmissionAccountDescription              string `gorm:"column:loan_sbmssn_acnt_desc" json:"loan_sbmssn_acnt_desc"`
}

func (LoanSubmissionAccountAnalysisMrdmData) TableName() string {
	return "t_loan_submission_account_analysis"
}

func (p *LoanSubmissionAccountAnalysis) FindLoanSubmissionAccountAnalysisMrdmByLoanSubmissionID(db *gorm.DB, pid uint64) ([]LoanSubmissionAccountAnalysisMrdmData, error) {
	var err error
	accountanalysis := []LoanSubmissionAccountAnalysisMrdmData{}
	err = db.Debug().Model(&LoanSubmissionAccountAnalysisMrdmData{}).Limit(100).
		Select(`t_loan_submission_account_analysis.loan_sbmssn_acnt_anlys_id,
				t_loan_submission_account_analysis.loan_sbmssn_id,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_anlys_name,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_anlys_code,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_anlys_receiver,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_anlys_type,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_anlys_facilities,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_salaries_recorded,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_salaries_recorded_document,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_desc`).
		Where("t_loan_submission_account_analysis.loan_sbmssn_id = ? AND t_loan_submission_account_analysis.loan_sbmssn_acnt_status = ?", pid, 1).
		Find(&accountanalysis).Error
	if err != nil {
		return []LoanSubmissionAccountAnalysisMrdmData{}, err
	}
	return accountanalysis, nil
}

func (p *LoanSubmissionAccountAnalysis) GetLoanSubmissionAccountAnalysisByLoanSubmissionID(db *gorm.DB, pid uint64) (*[]LoanSubmissionAccountAnalysisData, error) {
	var err error
	accountanalysis := []LoanSubmissionAccountAnalysisData{}
	err = db.Debug().Model(&LoanSubmissionAccountAnalysisData{}).Limit(100).
		Select(`t_loan_submission_account_analysis.loan_sbmssn_acnt_anlys_id,
				t_loan_submission_account_analysis.loan_sbmssn_id,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_anlys_name,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_anlys_code,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_anlys_receiver,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_anlys_type,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_anlys_facilities,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_salaries_recorded,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_salaries_recorded_document,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_desc,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_created_by,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_created_at,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_updated_by,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_updated_at,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_deleted_by,
				t_loan_submission_account_analysis.loan_sbmssn_acnt_deleted_at`).
		Where("t_loan_submission_account_analysis.loan_sbmssn_id = ? AND t_loan_submission_account_analysis.loan_sbmssn_acnt_status = ?", pid, 1).
		Find(&accountanalysis).Error
	if err != nil {
		return &[]LoanSubmissionAccountAnalysisData{}, err
	}
	return &accountanalysis, nil
}
