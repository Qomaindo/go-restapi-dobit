package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type LoanSubmissionInvestigation struct {
	LoanSubmissionInvestigationID          uint64     `gorm:"column:loan_sbmssn_invsgt_id;primary_key;" json:"loan_sbmssn_invsgt_id"`
	LoanSubmissionID                       uint64     `gorm:"column:loan_sbmssn_id" json:"loan_sbmssn_id"`
	LoanSubmissionInvestigationLoanQuality string     `gorm:"column:loan_sbmssn_invsgt_loan_quality;size:255" json:"loan_sbmssn_invsgt_loan_quality"`
	LoanSubmissionInvestigationDate        string     `gorm:"column:loan_sbmssn_invsgt_date" json:"loan_sbmssn_invsgt_date"`
	LoanSubmissionInvestigationTime        string     `gorm:"column:loan_sbmssn_invsgt_time" json:"loan_sbmssn_invsgt_time"`
	LoanSubmissionInvestigationDescription string     `gorm:"column:loan_sbmssn_invsgt_desc;size:255" json:"loan_sbmssn_invsgt_desc"`
	LoanSubmissionInvestigationStatus      uint64     `gorm:"column:loan_sbmssn_invsgt_status" json:"loan_sbmssn_invsgt_status"`
	LoanSubmissionInvestigationCreatedBy   string     `gorm:"column:loan_sbmssn_invsgt_created_by;size:125" json:"loan_sbmssn_invsgt_created_by"`
	LoanSubmissionInvestigationCreatedAt   time.Time  `gorm:"column:loan_sbmssn_invsgt_created_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_invsgt_created_at"`
	LoanSubmissionInvestigationUpdatedBy   string     `gorm:"column:loan_sbmssn_invsgt_updated_by;size:125" json:"loan_sbmssn_invsgt_updated_by"`
	LoanSubmissionInvestigationUpdatedAt   *time.Time `gorm:"column:loan_sbmssn_invsgt_updated_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_invsgt_updated_at"`
	LoanSubmissionInvestigationDeletedBy   string     `gorm:"column:loan_sbmssn_invsgt_deleted_by;size:125" json:"loan_sbmssn_invsgt_deleted_by"`
	LoanSubmissionInvestigationDeletedAt   *time.Time `gorm:"column:loan_sbmssn_invsgt_deleted_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_invsgt_deleted_at"`
}

type ResponseLoanSubmissionInvestigations struct {
	Status  int                            `json:"status"`
	Message string                         `json:"message"`
	Data    *[]LoanSubmissionInvestigation `json:"data"`
}

type ResponseLoanSubmissionInvestigation struct {
	Status  int                          `json:"status"`
	Message string                       `json:"message"`
	Data    *LoanSubmissionInvestigation `json:"data"`
}

type ResponseLoanSubmissionInvestigationCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (LoanSubmissionInvestigation) TableName() string {
	return "t_loan_submission_investigation"
}

func (p *LoanSubmissionInvestigation) Prepare() {
	p.LoanSubmissionID = p.LoanSubmissionID
	p.LoanSubmissionInvestigationLoanQuality = html.EscapeString(strings.TrimSpace(p.LoanSubmissionInvestigationLoanQuality))
	p.LoanSubmissionInvestigationDate = p.LoanSubmissionInvestigationDate
	p.LoanSubmissionInvestigationTime = p.LoanSubmissionInvestigationTime
	p.LoanSubmissionInvestigationDescription = html.EscapeString(strings.TrimSpace(p.LoanSubmissionInvestigationDescription))
	p.LoanSubmissionInvestigationStatus = p.LoanSubmissionInvestigationStatus
	p.LoanSubmissionInvestigationCreatedBy = p.LoanSubmissionInvestigationCreatedBy
	p.LoanSubmissionInvestigationCreatedAt = time.Now()
	p.LoanSubmissionInvestigationUpdatedBy = p.LoanSubmissionInvestigationUpdatedBy
	p.LoanSubmissionInvestigationUpdatedAt = p.LoanSubmissionInvestigationUpdatedAt
	p.LoanSubmissionInvestigationDeletedBy = p.LoanSubmissionInvestigationDeletedBy
	p.LoanSubmissionInvestigationDeletedAt = p.LoanSubmissionInvestigationDeletedAt
}

func (p *LoanSubmissionInvestigation) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.LoanSubmissionID == 0 {
			return errors.New("Loan Submission ID Tidak Boleh Kosong")
		}
		if p.LoanSubmissionInvestigationLoanQuality == "" {
			return errors.New("Loan Quality Tidak Boleh Kosong")
		}
		if p.LoanSubmissionInvestigationDate == "" {
			return errors.New("Tanggal Tidak Boleh Kosong")
		}
		if p.LoanSubmissionInvestigationTime == "" {
			return errors.New("Waktu Tidak Boleh Kosong")
		}
		if p.LoanSubmissionInvestigationDescription == "" {
			return errors.New("Deskripsi Tidak Boleh Kosong")
		}
		return nil
	}
}

func (p *LoanSubmissionInvestigation) SaveLoanSubmissionInvestigation(db *gorm.DB) (*LoanSubmissionInvestigation, error) {
	var err error
	err = db.Debug().Model(&LoanSubmissionInvestigation{}).Create(&p).Error
	if err != nil {
		return &LoanSubmissionInvestigation{}, err
	}
	return p, nil
}

func (p *LoanSubmissionInvestigation) FindAllLoanSubmissionInvestigation(db *gorm.DB) (*[]LoanSubmissionInvestigation, error) {
	var err error
	loansubmissioninvestigation := []LoanSubmissionInvestigation{}
	err = db.Debug().Model(&LoanSubmissionInvestigation{}).Limit(100).Find(&loansubmissioninvestigation).Error
	if err != nil {
		return &[]LoanSubmissionInvestigation{}, err
	}
	return &loansubmissioninvestigation, nil
}

func (p *LoanSubmissionInvestigation) FindLoanSubmissionInvestigationByID(db *gorm.DB, pid uint64) (*LoanSubmissionInvestigation, error) {
	var err error
	err = db.Debug().Model(&LoanSubmissionInvestigation{}).Where("loan_sbmssn_invsgt_id = ?", pid).Take(&p).Error
	if err != nil {
		return &LoanSubmissionInvestigation{}, err
	}
	return p, nil
}

func (p *LoanSubmissionInvestigation) UpdateLoanSubmissionInvestigation(db *gorm.DB, pid uint64) (*LoanSubmissionInvestigation, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionInvestigation{}).Where("loan_sbmssn_invsgt_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_invsgt_loan_quality": p.LoanSubmissionInvestigationLoanQuality,
			"loan_sbmssn_invsgt_date":         p.LoanSubmissionInvestigationDate,
			"loan_sbmssn_invsgt_time":         p.LoanSubmissionInvestigationTime,
			"loan_sbmssn_invsgt_desc":         p.LoanSubmissionInvestigationDescription,
			"loan_sbmssn_invsgt_updated_by":   p.LoanSubmissionInvestigationUpdatedBy,
			"loan_sbmssn_invsgt_updated_at":   time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionInvestigation{}).Where("loan_sbmssn_invsgt_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionInvestigation{}, err
	}
	return p, nil
}

func (p *LoanSubmissionInvestigation) UpdateLoanSubmissionInvestigationDelete(db *gorm.DB, pid uint64) (*LoanSubmissionInvestigation, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionInvestigation{}).Where("loan_sbmssn_invsgt_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_invsgt_status":     p.LoanSubmissionInvestigationStatus,
			"loan_sbmssn_invsgt_deleted_by": p.LoanSubmissionInvestigationDeletedBy,
			"loan_sbmssn_invsgt_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionInvestigation{}).Where("loan_sbmssn_invsgt_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionInvestigation{}, err
	}
	return p, nil
}

func (p *LoanSubmissionInvestigation) DeleteLoanSubmissionInvestigation(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&LoanSubmissionInvestigation{}).Where("loan_sbmssn_invsgt_id = ?", pid).Take(&LoanSubmissionInvestigation{}).Delete(&LoanSubmissionInvestigation{})
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

type LoanSubmissionInvestigationMrdmData struct {
	LoanSubmissionInvestigationID          uint64                                         `gorm:"column:loan_sbmssn_invsgt_id" json:"loan_sbmssn_invsgt_id"`
	LoanSubmissionID                       uint64                                         `gorm:"column:loan_sbmssn_id" json:"loan_sbmssn_id"`
	LoanSubmissionInvestigationLoanQuality string                                         `gorm:"column:loan_sbmssn_invsgt_loan_quality" json:"loan_sbmssn_invsgt_loan_quality"`
	LoanSubmissionInvestigationDate        string                                         `gorm:"column:loan_sbmssn_invsgt_date" json:"loan_sbmssn_invsgt_date"`
	LoanSubmissionInvestigationTime        string                                         `gorm:"column:loan_sbmssn_invsgt_time" json:"loan_sbmssn_invsgt_time"`
	LoanSubmissionInvestigationDescription string                                         `gorm:"column:loan_sbmssn_invsgt_desc" json:"loan_sbmssn_invsgt_desc"`
	LoanSubmissionInvestigationReference   []LoanSubmissionInvestigationReferenceMrdmData `gorm:"column:loan_sbmssn_invsgt_ref" json:"loan_sbmssn_invsgt_ref"`
}

func (LoanSubmissionInvestigationMrdmData) TableName() string {
	return "t_loan_submission_investigation"
}

func (p *LoanSubmissionInvestigation) FindLoanSubmissionInvestigationByLoanSubmissionID(db *gorm.DB, pid uint64) (LoanSubmissionInvestigationMrdmData, error) {
	var err error
	loansubmissioninvestigation := LoanSubmissionInvestigationMrdmData{}
	err = db.Debug().Model(&LoanSubmissionInvestigationMrdmData{}).Where("loan_sbmssn_id = ?", pid).Find(&loansubmissioninvestigation).Error
	if err != nil {
		return LoanSubmissionInvestigationMrdmData{}, err
	}
	return loansubmissioninvestigation, nil
}

func (p *LoanSubmissionInvestigation) UpdateLoanSubmissionInvestigationByLoanSubmissionID(db *gorm.DB, pid uint64) (*LoanSubmissionInvestigation, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionInvestigation{}).Where("loan_sbmssn_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_invsgt_loan_quality": p.LoanSubmissionInvestigationLoanQuality,
			"loan_sbmssn_invsgt_date":         p.LoanSubmissionInvestigationDate,
			"loan_sbmssn_invsgt_time":         p.LoanSubmissionInvestigationTime,
			"loan_sbmssn_invsgt_desc":         p.LoanSubmissionInvestigationDescription,
			"loan_sbmssn_invsgt_updated_by":   p.LoanSubmissionInvestigationUpdatedBy,
			"loan_sbmssn_invsgt_updated_at":   time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionInvestigation{}).Where("loan_sbmssn_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionInvestigation{}, err
	}
	return p, nil
}
