package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type LoanSubmissionConclution struct {
	LoanSubmissionConclutionID                   uint64     `gorm:"column:loan_sbmssn_conclu_id;primary_key;" json:"loan_sbmssn_conclu_id"`
	LoanSubmissionID                             uint64     `gorm:"column:loan_sbmssn_id" json:"loan_sbmssn_id"`
	LoanSubmissionConclutionAnalysisPro          string     `gorm:"column:loan_sbmssn_conclu_anlys_pro" json:"loan_sbmssn_conclu_anlys_pro"`
	LoanSubmissionConclutionAnalysisContra       string     `gorm:"column:loan_sbmssn_conclu_anlys_contra" json:"loan_sbmssn_conclu_anlys_contra"`
	LoanSubmissionConclutionAnalysisMitigation   string     `gorm:"column:loan_sbmssn_conclu_anlys_mitigation" json:"loan_sbmssn_conclu_anlys_mitigation"`
	LoanSubmissionConclutionAsesmentPefindo      string     `gorm:"column:loan_sbmssn_conclu_assmnt_pefindo" json:"loan_sbmssn_conclu_assmnt_pefindo"`
	LoanSubmissionConclutionAsesmentCreditRating string     `gorm:"column:loan_sbmssn_conclu_assmnt_credit_rating" json:"loan_sbmssn_conclu_assmnt_credit_rating"`
	LoanSubmissionConclutionDescription          string     `gorm:"column:loan_sbmssn_conclu_desc" json:"loan_sbmssn_conclu_desc"`
	LoanSubmissionConclutionStatus               uint64     `gorm:"column:loan_sbmssn_conclu_status" json:"loan_sbmssn_conclu_status"`
	LoanSubmissionConclutionCreatedBy            string     `gorm:"column:loan_sbmssn_conclu_created_by;size:125" json:"loan_sbmssn_conclu_created_by"`
	LoanSubmissionConclutionCreatedAt            time.Time  `gorm:"column:loan_sbmssn_conclu_created_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_conclu_created_at"`
	LoanSubmissionConclutionUpdatedBy            string     `gorm:"column:loan_sbmssn_conclu_updated_by;size:125" json:"loan_sbmssn_conclu_updated_by"`
	LoanSubmissionConclutionUpdatedAt            *time.Time `gorm:"column:loan_sbmssn_conclu_updated_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_conclu_updated_at"`
	LoanSubmissionConclutionDeletedBy            string     `gorm:"column:loan_sbmssn_conclu_deleted_by;size:125" json:"loan_sbmssn_conclu_deleted_by"`
	LoanSubmissionConclutionDeletedAt            *time.Time `gorm:"column:loan_sbmssn_conclu_deleted_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_conclu_deleted_at"`
}

type ResponseLoanSubmissionConclutions struct {
	Status  int                         `json:"status"`
	Message string                      `json:"message"`
	Data    *[]LoanSubmissionConclution `json:"data"`
}

type ResponseLoanSubmissionConclution struct {
	Status  int                       `json:"status"`
	Message string                    `json:"message"`
	Data    *LoanSubmissionConclution `json:"data"`
}

type ResponseLoanSubmissionConclutionCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (LoanSubmissionConclution) TableName() string {
	return "t_loan_submission_conclution"
}

func (p *LoanSubmissionConclution) Prepare() {
	p.LoanSubmissionID = p.LoanSubmissionID
	p.LoanSubmissionConclutionAnalysisPro = p.LoanSubmissionConclutionAnalysisPro
	p.LoanSubmissionConclutionAnalysisContra = p.LoanSubmissionConclutionAnalysisContra
	p.LoanSubmissionConclutionAnalysisMitigation = p.LoanSubmissionConclutionAnalysisMitigation
	p.LoanSubmissionConclutionAsesmentPefindo = p.LoanSubmissionConclutionAsesmentPefindo
	p.LoanSubmissionConclutionAsesmentCreditRating = p.LoanSubmissionConclutionAsesmentCreditRating
	p.LoanSubmissionConclutionDescription = html.EscapeString(strings.TrimSpace(p.LoanSubmissionConclutionDescription))
	p.LoanSubmissionConclutionStatus = p.LoanSubmissionConclutionStatus
	p.LoanSubmissionConclutionCreatedBy = p.LoanSubmissionConclutionCreatedBy
	p.LoanSubmissionConclutionCreatedAt = time.Now()
	p.LoanSubmissionConclutionUpdatedBy = p.LoanSubmissionConclutionUpdatedBy
	p.LoanSubmissionConclutionUpdatedAt = p.LoanSubmissionConclutionUpdatedAt
	p.LoanSubmissionConclutionDeletedBy = p.LoanSubmissionConclutionDeletedBy
	p.LoanSubmissionConclutionDeletedAt = p.LoanSubmissionConclutionDeletedAt
}

func (p *LoanSubmissionConclution) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.LoanSubmissionID == 0 {
			return errors.New("Loan Submission ID Tidak Boleh Kosong")
		}
		if p.LoanSubmissionConclutionAnalysisPro == "" {
			return errors.New("Pro Tidak Boleh Kosong")
		}
		if p.LoanSubmissionConclutionAnalysisContra == "" {
			return errors.New("Contra Tidak Boleh Kosong")
		}
		if p.LoanSubmissionConclutionAnalysisMitigation == "" {
			return errors.New("Mitigation Tidak Boleh Kosong")
		}
		if p.LoanSubmissionConclutionAsesmentPefindo == "" {
			return errors.New("Pefindo Tidak Boleh Kosong")
		}
		if p.LoanSubmissionConclutionAsesmentCreditRating == "" {
			return errors.New("Credit Rating Tidak Boleh Kosong")
		}
		if p.LoanSubmissionConclutionDescription == "" {
			return errors.New("Deskripsi Tidak Boleh Kosong")
		}
		return nil
	}
}

func (p *LoanSubmissionConclution) SaveLoanSubmissionConclution(db *gorm.DB) (*LoanSubmissionConclution, error) {
	var err error
	err = db.Debug().Model(&LoanSubmissionConclution{}).Create(&p).Error
	if err != nil {
		return &LoanSubmissionConclution{}, err
	}
	return p, nil
}

func (p *LoanSubmissionConclution) FindAllLoanSubmissionConclution(db *gorm.DB) (*[]LoanSubmissionConclution, error) {
	var err error
	loansbmssnconclu := []LoanSubmissionConclution{}
	err = db.Debug().Model(&LoanSubmissionConclution{}).Limit(100).Find(&loansbmssnconclu).Error
	if err != nil {
		return &[]LoanSubmissionConclution{}, err
	}
	return &loansbmssnconclu, nil
}

func (p *LoanSubmissionConclution) FindLoanSubmissionConclutionByID(db *gorm.DB, pid uint64) (*LoanSubmissionConclution, error) {
	var err error
	err = db.Debug().Model(&LoanSubmissionConclution{}).Where("loan_sbmssn_conclu_id = ?", pid).Take(&p).Error
	if err != nil {
		return &LoanSubmissionConclution{}, err
	}
	return p, nil
}

func (p *LoanSubmissionConclution) UpdateLoanSubmissionConclution(db *gorm.DB, pid uint64) (*LoanSubmissionConclution, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionConclution{}).Where("loan_sbmssn_conclu_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_conclu_anlys_pro":            p.LoanSubmissionConclutionAnalysisPro,
			"loan_sbmssn_conclu_anlys_contra":         p.LoanSubmissionConclutionAnalysisContra,
			"loan_sbmssn_conclu_anlys_mitigation":     p.LoanSubmissionConclutionAnalysisMitigation,
			"loan_sbmssn_conclu_assmnt_pefindo":       p.LoanSubmissionConclutionAsesmentPefindo,
			"loan_sbmssn_conclu_assmnt_credit_rating": p.LoanSubmissionConclutionAsesmentCreditRating,
			"loan_sbmssn_conclu_desc":                 p.LoanSubmissionConclutionDescription,
			"loan_sbmssn_conclu_updated_by":           p.LoanSubmissionConclutionUpdatedBy,
			"loan_sbmssn_conclu_updated_at":           time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionConclution{}).Where("loan_sbmssn_conclu_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionConclution{}, err
	}
	return p, nil
}

func (p *LoanSubmissionConclution) UpdateLoanSubmissionConclutionDelete(db *gorm.DB, pid uint64) (*LoanSubmissionConclution, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionConclution{}).Where("loan_sbmssn_conclu_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_conclu_status":     p.LoanSubmissionConclutionStatus,
			"loan_sbmssn_conclu_deleted_by": p.LoanSubmissionConclutionDeletedBy,
			"loan_sbmssn_conclu_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionConclution{}).Where("loan_sbmssn_conclu_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionConclution{}, err
	}
	return p, nil
}

func (p *LoanSubmissionConclution) DeleteLoanSubmissionConclution(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&LoanSubmissionConclution{}).Where("loan_sbmssn_conclu_id = ?", pid).Take(&LoanSubmissionConclution{}).Delete(&LoanSubmissionConclution{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Loan Submission Investigation not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// ADDITIONAL MEMORANDUM
// ===================================================================================================================================

type LoanSubmissionConlutionMrdmData struct {
	LoanSubmissionConlutionID                uint64 `gorm:"column:loan_sbmssn_conclu_id" json:"loan_sbmssn_conclu_id"`
	LoanSubmissionID                         uint64 `gorm:"column:loan_sbmssn_id" json:"loan_sbmssn_id"`
	LoanSubmissionConlutionAnlysPro          string `gorm:"column:loan_sbmssn_conclu_anlys_pro" json:"loan_sbmssn_conclu_anlys_pro"`
	LoanSubmissionConlutionAnlysContra       string `gorm:"column:loan_sbmssn_conclu_anlys_contra" json:"loan_sbmssn_conclu_anlys_contra"`
	LoanSubmissionConlutionAnlysMtgtn        string `gorm:"column:loan_sbmssn_conclu_anlys_mitigation" json:"loan_sbmssn_conclu_anlys_mitigation"`
	LoanSubmissionConlutionAssmtPefindo      string `gorm:"column:loan_sbmssn_conclu_assmnt_pefindo" json:"loan_sbmssn_conclu_assmnt_pefindo"`
	LoanSubmissionConlutionAssmtCreditRating string `gorm:"column:loan_sbmssn_conclu_assmnt_credit_rating" json:"loan_sbmssn_conclu_assmnt_credit_rating"`
	LoanSubmissionConlutionDesc              string `gorm:"column:loan_sbmssn_conclu_desc" json:"loan_sbmssn_conclu_desc"`
	LoanSubmissionConlutionStatus            uint64 `gorm:"column:loan_sbmssn_conclu_status" json:"loan_sbmssn_conclu_status"`
}

func (LoanSubmissionConlutionMrdmData) TableName() string {
	return "t_loan_submission_conclution"
}

func (p *LoanSubmissionConclution) FindLoanSubmissionConlutionByLoanSubmissionID(db *gorm.DB, pid uint64) (LoanSubmissionConlutionMrdmData, error) {
	var err error
	loansbmssnconclu := LoanSubmissionConlutionMrdmData{}
	err = db.Debug().Model(&LoanSubmissionConlutionMrdmData{}).Where("loan_sbmssn_id = ?", pid).Find(&loansbmssnconclu).Error
	if err != nil {
		return LoanSubmissionConlutionMrdmData{}, err
	}
	return loansbmssnconclu, nil
}

func (p *LoanSubmissionConclution) UpdateLoanSubmissionConclutionByLoanSubmissionID(db *gorm.DB, pid uint64) (*LoanSubmissionConclution, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionConclution{}).Where("loan_sbmssn_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_conclu_anlys_pro":            p.LoanSubmissionConclutionAnalysisPro,
			"loan_sbmssn_conclu_anlys_contra":         p.LoanSubmissionConclutionAnalysisContra,
			"loan_sbmssn_conclu_anlys_mitigation":     p.LoanSubmissionConclutionAnalysisMitigation,
			"loan_sbmssn_conclu_assmnt_pefindo":       p.LoanSubmissionConclutionAsesmentPefindo,
			"loan_sbmssn_conclu_assmnt_credit_rating": p.LoanSubmissionConclutionAsesmentCreditRating,
			"loan_sbmssn_conclu_desc":                 p.LoanSubmissionConclutionDescription,
			"loan_sbmssn_conclu_updated_by":           p.LoanSubmissionConclutionUpdatedBy,
			"loan_sbmssn_conclu_updated_at":           time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionConclution{}).Where("loan_sbmssn_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionConclution{}, err
	}
	return p, nil
}
