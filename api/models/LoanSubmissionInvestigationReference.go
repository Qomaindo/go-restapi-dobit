package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type LoanSubmissionInvestigationReference struct {
	LoanSubmissionInvestigationReferenceID          uint64     `gorm:"column:loan_sbmssn_invsgt_ref_id;primary_key;" json:"loan_sbmssn_invsgt_ref_id"`
	LoanSubmissionInvestigationID                   uint64     `gorm:"column:loan_sbmssn_invsgt_id" json:"loan_sbmssn_invsgt_id"`
	LoanSubmissionInvestigationReferenceType        string     `gorm:"column:loan_sbmssn_invsgt_ref_type;size:255" json:"loan_sbmssn_invsgt_ref_type"`
	LoanSubmissionInvestigationReferenceName        string     `gorm:"column:loan_sbmssn_invsgt_ref_name" json:"loan_sbmssn_invsgt_ref_name"`
	LoanSubmissionInvestigationReferenceRelation    string     `gorm:"column:loan_sbmssn_invsgt_ref_relation" json:"loan_sbmssn_invsgt_ref_relation"`
	LoanSubmissionInvestigationReferencePhone       string     `gorm:"column:loan_sbmssn_invsgt_ref_phone" json:"loan_sbmssn_invsgt_ref_phone"`
	LoanSubmissionInvestigationReferenceEmail       string     `gorm:"column:loan_sbmssn_invsgt_ref_email" json:"loan_sbmssn_invsgt_ref_email"`
	LoanSubmissionInvestigationReferenceDate        string     `gorm:"column:loan_sbmssn_invsgt_ref_date" json:"loan_sbmssn_invsgt_ref_date"`
	LoanSubmissionInvestigationReferenceTime        string     `gorm:"column:loan_sbmssn_invsgt_ref_time" json:"loan_sbmssn_invsgt_ref_time"`
	LoanSubmissionInvestigationReferenceDescription string     `gorm:"column:loan_sbmssn_invsgt_ref_desc;size:255" json:"loan_sbmssn_invsgt_ref_desc"`
	LoanSubmissionInvestigationReferenceStatus      uint64     `gorm:"column:loan_sbmssn_invsgt_ref_status" json:"loan_sbmssn_invsgt_ref_status"`
	LoanSubmissionInvestigationReferenceCreatedBy   string     `gorm:"column:loan_sbmssn_invsgt_ref_created_by;size:125" json:"loan_sbmssn_invsgt_ref_created_by"`
	LoanSubmissionInvestigationReferenceCreatedAt   time.Time  `gorm:"column:loan_sbmssn_invsgt_ref_created_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_invsgt_ref_created_at"`
	LoanSubmissionInvestigationReferenceUpdatedBy   string     `gorm:"column:loan_sbmssn_invsgt_ref_updated_by;size:125" json:"loan_sbmssn_invsgt_ref_updated_by"`
	LoanSubmissionInvestigationReferenceUpdatedAt   *time.Time `gorm:"column:loan_sbmssn_invsgt_ref_updated_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_invsgt_ref_updated_at"`
	LoanSubmissionInvestigationReferenceDeletedBy   string     `gorm:"column:loan_sbmssn_invsgt_ref_deleted_by;size:125" json:"loan_sbmssn_invsgt_ref_deleted_by"`
	LoanSubmissionInvestigationReferenceDeletedAt   *time.Time `gorm:"column:loan_sbmssn_invsgt_info_deleted_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_invsgt_info_deleted_at"`
}

type ResponseLoanSubmissionInvestigationReferences struct {
	Status  int                                     `json:"status"`
	Message string                                  `json:"message"`
	Data    *[]LoanSubmissionInvestigationReference `json:"data"`
}

type ResponseLoanSubmissionInvestigationReference struct {
	Status  int                                   `json:"status"`
	Message string                                `json:"message"`
	Data    *LoanSubmissionInvestigationReference `json:"data"`
}

type ResponseLoanSubmissionInvestigationReferenceCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (LoanSubmissionInvestigationReference) TableName() string {
	return "t_loan_submission_investigation_reference"
}

func (p *LoanSubmissionInvestigationReference) Prepare() {
	p.LoanSubmissionInvestigationID = p.LoanSubmissionInvestigationID
	p.LoanSubmissionInvestigationReferenceType = html.EscapeString(strings.TrimSpace(p.LoanSubmissionInvestigationReferenceType))
	p.LoanSubmissionInvestigationReferenceName = html.EscapeString(strings.TrimSpace(p.LoanSubmissionInvestigationReferenceName))
	p.LoanSubmissionInvestigationReferenceRelation = html.EscapeString(strings.TrimSpace(p.LoanSubmissionInvestigationReferenceRelation))
	p.LoanSubmissionInvestigationReferencePhone = html.EscapeString(strings.TrimSpace(p.LoanSubmissionInvestigationReferencePhone))
	p.LoanSubmissionInvestigationReferenceEmail = html.EscapeString(strings.TrimSpace(p.LoanSubmissionInvestigationReferenceEmail))
	p.LoanSubmissionInvestigationReferenceDate = p.LoanSubmissionInvestigationReferenceDate
	p.LoanSubmissionInvestigationReferenceTime = p.LoanSubmissionInvestigationReferenceTime
	p.LoanSubmissionInvestigationReferenceDescription = html.EscapeString(strings.TrimSpace(p.LoanSubmissionInvestigationReferenceDescription))
	p.LoanSubmissionInvestigationReferenceStatus = p.LoanSubmissionInvestigationReferenceStatus
	p.LoanSubmissionInvestigationReferenceCreatedBy = p.LoanSubmissionInvestigationReferenceCreatedBy
	p.LoanSubmissionInvestigationReferenceCreatedAt = time.Now()
	p.LoanSubmissionInvestigationReferenceUpdatedBy = p.LoanSubmissionInvestigationReferenceUpdatedBy
	p.LoanSubmissionInvestigationReferenceUpdatedAt = p.LoanSubmissionInvestigationReferenceUpdatedAt
	p.LoanSubmissionInvestigationReferenceDeletedBy = p.LoanSubmissionInvestigationReferenceDeletedBy
	p.LoanSubmissionInvestigationReferenceDeletedAt = p.LoanSubmissionInvestigationReferenceDeletedAt
}

func (p *LoanSubmissionInvestigationReference) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.LoanSubmissionInvestigationID == 0 {
			return errors.New("Loan Investigation ID Tidak Boleh Kosong")
		}
		if p.LoanSubmissionInvestigationReferenceType == "" {
			return errors.New("Loan Quality Tidak Boleh Kosong")
		}
		if p.LoanSubmissionInvestigationReferenceName == "" {
			return errors.New("Tanggal Tidak Boleh Kosong")
		}
		if p.LoanSubmissionInvestigationReferenceRelation == "" {
			return errors.New("Waktu Tidak Boleh Kosong")
		}
		if p.LoanSubmissionInvestigationReferencePhone == "" {
			return errors.New("Deskripsi Tidak Boleh Kosong")
		}
		if p.LoanSubmissionInvestigationReferenceEmail == "" {
			return errors.New("Deskripsi Tidak Boleh Kosong")
		}
		if p.LoanSubmissionInvestigationReferenceDate == "" {
			return errors.New("Deskripsi Tidak Boleh Kosong")
		}
		if p.LoanSubmissionInvestigationReferenceTime == "" {
			return errors.New("Deskripsi Tidak Boleh Kosong")
		}
		if p.LoanSubmissionInvestigationReferenceDescription == "" {
			return errors.New("Deskripsi Tidak Boleh Kosong")
		}
		return nil
	}
}

func (p *LoanSubmissionInvestigationReference) SaveLoanSubmissionInvestigationReference(db *gorm.DB) (*LoanSubmissionInvestigationReference, error) {
	var err error
	err = db.Debug().Model(&LoanSubmissionInvestigationReference{}).Create(&p).Error
	if err != nil {
		return &LoanSubmissionInvestigationReference{}, err
	}
	return p, nil
}

func (p *LoanSubmissionInvestigationReference) FindAllLoanSubmissionInvestigation(db *gorm.DB) (*[]LoanSubmissionInvestigationReference, error) {
	var err error
	loansubmissioninvestigation := []LoanSubmissionInvestigationReference{}
	err = db.Debug().Model(&LoanSubmissionInvestigationReference{}).Limit(100).Find(&loansubmissioninvestigation).Error
	if err != nil {
		return &[]LoanSubmissionInvestigationReference{}, err
	}
	return &loansubmissioninvestigation, nil
}

func (p *LoanSubmissionInvestigationReference) FindLoanSubmissionInvestigationReferenceByID(db *gorm.DB, pid uint64) (*LoanSubmissionInvestigationReference, error) {
	var err error
	err = db.Debug().Model(&LoanSubmissionInvestigationReference{}).Where("loan_sbmssn_invsgt_ref_id = ?", pid).Take(&p).Error
	if err != nil {
		return &LoanSubmissionInvestigationReference{}, err
	}
	return p, nil
}

func (p *LoanSubmissionInvestigationReference) UpdateLoanSubmissionInvestigationReference(db *gorm.DB, pid uint64) (*LoanSubmissionInvestigationReference, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionInvestigationReference{}).Where("loan_sbmssn_invsgt_ref_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_invsgt_ref_type":       p.LoanSubmissionInvestigationReferenceType,
			"loan_sbmssn_invsgt_ref_name":       p.LoanSubmissionInvestigationReferenceName,
			"loan_sbmssn_invsgt_ref_relation":   p.LoanSubmissionInvestigationReferenceRelation,
			"loan_sbmssn_invsgt_ref_phone":      p.LoanSubmissionInvestigationReferencePhone,
			"loan_sbmssn_invsgt_ref_email":      p.LoanSubmissionInvestigationReferenceEmail,
			"loan_sbmssn_invsgt_ref_date":       p.LoanSubmissionInvestigationReferenceDate,
			"loan_sbmssn_invsgt_ref_time":       p.LoanSubmissionInvestigationReferenceTime,
			"loan_sbmssn_invsgt_ref_desc":       p.LoanSubmissionInvestigationReferenceDescription,
			"loan_sbmssn_invsgt_ref_updated_by": p.LoanSubmissionInvestigationReferenceUpdatedBy,
			"loan_sbmssn_invsgt_ref_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionInvestigationReference{}).Where("loan_sbmssn_invsgt_ref_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionInvestigationReference{}, err
	}
	return p, nil
}

func (p *LoanSubmissionInvestigationReference) UpdateLoanSubmissionInvestigationReferenceDelete(db *gorm.DB, pid uint64) (*LoanSubmissionInvestigationReference, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionInvestigationReference{}).Where("loan_sbmssn_invsgt_ref_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_invsgt_ref_status":     p.LoanSubmissionInvestigationReferenceStatus,
			"loan_sbmssn_invsgt_ref_deleted_by": p.LoanSubmissionInvestigationReferenceDeletedBy,
			"loan_sbmssn_invsgt_ref_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionInvestigationReference{}).Where("loan_sbmssn_invsgt_ref_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionInvestigationReference{}, err
	}
	return p, nil
}

func (p *LoanSubmissionInvestigationReference) DeleteLoanSubmissionInvestigationReference(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&LoanSubmissionInvestigationReference{}).Where("loan_sbmssn_invsgt_ref_id = ?", pid).Take(&LoanSubmissionInvestigationReference{}).Delete(&LoanSubmissionInvestigationReference{})
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

type LoanSubmissionInvestigationReferenceMrdmData struct {
	LoanSubmissionInvestigationReferenceID          uint64 `gorm:"column:loan_sbmssn_invsgt_ref_id" json:"loan_sbmssn_invsgt_ref_id"`
	LoanSubmissionInvestigationID                   uint64 `gorm:"column:loan_sbmssn_invsgt_id" json:"loan_sbmssn_invsgt_id"`
	LoanSubmissionInvestigationReferenceType        string `gorm:"column:loan_sbmssn_invsgt_ref_type" json:"loan_sbmssn_invsgt_ref_type"`
	LoanSubmissionInvestigationReferenceName        string `gorm:"column:loan_sbmssn_invsgt_ref_name" json:"loan_sbmssn_invsgt_ref_name"`
	LoanSubmissionInvestigationReferenceRelation    string `gorm:"column:loan_sbmssn_invsgt_ref_relation" json:"loan_sbmssn_invsgt_ref_relation"`
	LoanSubmissionInvestigationReferencePhone       string `gorm:"column:loan_sbmssn_invsgt_ref_phone" json:"loan_sbmssn_invsgt_ref_phone"`
	LoanSubmissionInvestigationReferenceEmail       string `gorm:"column:loan_sbmssn_invsgt_ref_email" json:"loan_sbmssn_invsgt_ref_email"`
	LoanSubmissionInvestigationReferenceDate        string `gorm:"column:loan_sbmssn_invsgt_ref_date" json:"loan_sbmssn_invsgt_ref_date"`
	LoanSubmissionInvestigationReferenceTime        string `gorm:"column:loan_sbmssn_invsgt_ref_time" json:"loan_sbmssn_invsgt_ref_time"`
	LoanSubmissionInvestigationReferenceDescription string `gorm:"column:loan_sbmssn_invsgt_ref_desc" json:"loan_sbmssn_invsgt_ref_desc"`
}

func (LoanSubmissionInvestigationReferenceMrdmData) TableName() string {
	return "t_loan_submission_investigation_reference"
}

func (p *LoanSubmissionInvestigationReference) FindLoanSubmissionInvestigationReferenceByLoanSubmissionInvestigationID(db *gorm.DB, pid uint64) ([]LoanSubmissionInvestigationReferenceMrdmData, error) {
	var err error
	loansubmissioninvestigation := []LoanSubmissionInvestigationReferenceMrdmData{}
	err = db.Debug().Model(&LoanSubmissionInvestigationReferenceMrdmData{}).Where("loan_sbmssn_invsgt_id = ? AND loan_sbmssn_invsgt_ref_status = ?", pid, 1).Find(&loansubmissioninvestigation).Error
	if err != nil {
		return []LoanSubmissionInvestigationReferenceMrdmData{}, err
	}
	return loansubmissioninvestigation, nil
}

func (p *LoanSubmissionInvestigationReference) GetLoanSubmissionInvestigationReferenceByLoanInvestigationID(db *gorm.DB, pid uint64) (*[]LoanSubmissionInvestigationReference, error) {
	var err error
	loansubmissioninvestigation := []LoanSubmissionInvestigationReference{}
	err = db.Debug().Model(&LoanSubmissionInvestigationReference{}).Where("loan_sbmssn_invsgt_id = ? AND loan_sbmssn_invsgt_ref_status = ?", pid, 1).Limit(100).Find(&loansubmissioninvestigation).Error
	if err != nil {
		return &[]LoanSubmissionInvestigationReference{}, err
	}
	return &loansubmissioninvestigation, nil
}
