package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type LoanEvidenceTransfer struct {
	LoanEvidenceTransferID        uint64     `gorm:"column:loan_evidence_trf_id;primary_key;" json:"loan_evidence_trf_id"`
	LoanID                        uint64     `gorm:"column:loan_id" json:"loan_id"`
	LoanEvidenceTransferValue     string     `gorm:"column:loan_evidence_trf_value;size:255" json:"loan_evidence_trf_value"`
	LoanEvidenceTransferStatus    uint64     `gorm:"column:loan_evidence_trf_status" json:"loan_evidence_trf_status"`
	LoanEvidenceTransferCreatedBy string     `gorm:"column:loan_evidence_trf_created_by;size:125" json:"loan_evidence_trf_created_by"`
	LoanEvidenceTransferCreatedAt time.Time  `gorm:"column:loan_evidence_trf_created_at;default:CURRENT_TIMESTAMP" json:"loan_evidence_trf_created_at"`
	LoanEvidenceTransferUpdatedBy string     `gorm:"column:loan_evidence_trf_updated_by;size:125" json:"loan_evidence_trf_updated_by"`
	LoanEvidenceTransferUpdatedAt *time.Time `gorm:"column:loan_evidence_trf_updated_at;default:CURRENT_TIMESTAMP" json:"loan_evidence_trf_updated_at"`
	LoanEvidenceTransferDeletedBy string     `gorm:"column:loan_evidence_trf_deleted_by;size:125" json:"loan_evidence_trf_deleted_by"`
	LoanEvidenceTransferDeletedAt *time.Time `gorm:"column:loan_evidence_trf_deleted_at;default:CURRENT_TIMESTAMP" json:"loan_evidence_trf_deleted_at"`
}

type ResponseLoanEvidenceTransfers struct {
	Status  int                     `json:"status"`
	Message string                  `json:"message"`
	Data    *[]LoanEvidenceTransfer `json:"data"`
}

type ResponseLoanEvidenceTransfer struct {
	Status  int                   `json:"status"`
	Message string                `json:"message"`
	Data    *LoanEvidenceTransfer `json:"data"`
}

type ResponseLoanEvidenceTransferCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (LoanEvidenceTransfer) TableName() string {
	return "t_loan_evidence_transfer"
}

func (p *LoanEvidenceTransfer) Prepare() {
	p.LoanID = p.LoanID
	p.LoanEvidenceTransferValue = html.EscapeString(strings.TrimSpace(p.LoanEvidenceTransferValue))
	p.LoanEvidenceTransferStatus = p.LoanEvidenceTransferStatus
	p.LoanEvidenceTransferCreatedAt = time.Now()
	p.LoanEvidenceTransferCreatedBy = p.LoanEvidenceTransferCreatedBy
	p.LoanEvidenceTransferUpdatedBy = p.LoanEvidenceTransferUpdatedBy
	p.LoanEvidenceTransferUpdatedAt = p.LoanEvidenceTransferUpdatedAt
	p.LoanEvidenceTransferDeletedBy = p.LoanEvidenceTransferDeletedBy
	p.LoanEvidenceTransferDeletedAt = p.LoanEvidenceTransferDeletedAt
}

func (p *LoanEvidenceTransfer) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.LoanID == 0 {
			return errors.New("Loan ID Tidak Boleh Kosong")
		}
		if p.LoanEvidenceTransferValue == "" {
			return errors.New("Loan Evidence Transfer Value Tidak Boleh Kosong")
		}
		return nil
	}
}

func (p *LoanEvidenceTransfer) SaveLoanEvidenceTransfer(db *gorm.DB) (*LoanEvidenceTransfer, error) {
	var err error
	err = db.Debug().Model(&LoanEvidenceTransfer{}).Create(&p).Error
	if err != nil {
		return &LoanEvidenceTransfer{}, err
	}
	return p, nil
}

func (p *LoanEvidenceTransfer) FindAllLoanEvidenceTransfer(db *gorm.DB) (*[]LoanEvidenceTransfer, error) {
	var err error
	loanevidencetransfer := []LoanEvidenceTransfer{}
	err = db.Debug().Model(&LoanEvidenceTransfer{}).Limit(100).Find(&loanevidencetransfer).Error
	if err != nil {
		return &[]LoanEvidenceTransfer{}, err
	}
	return &loanevidencetransfer, nil
}

func (p *LoanEvidenceTransfer) FindAllLoanEvidenceTransferStatus(db *gorm.DB, status uint64) (*[]LoanEvidenceTransfer, error) {
	var err error
	loanevidencetransfer := []LoanEvidenceTransfer{}
	err = db.Debug().Model(&LoanEvidenceTransfer{}).Where("referal_status = ?", status).Limit(100).Find(&loanevidencetransfer).Error
	if err != nil {
		return &[]LoanEvidenceTransfer{}, err
	}
	return &loanevidencetransfer, nil
}

func (p *LoanEvidenceTransfer) FindLoanEvidenceTransferByID(db *gorm.DB, pid uint64) (*LoanEvidenceTransfer, error) {
	var err error
	err = db.Debug().Model(&LoanEvidenceTransfer{}).Where("loan_evidence_trf_id = ?", pid).Take(&p).Error
	if err != nil {
		return &LoanEvidenceTransfer{}, err
	}
	return p, nil
}

func (p *LoanEvidenceTransfer) FindAllLoanEvidenceTransferByID(db *gorm.DB, pid uint64) (*[]LoanEvidenceTransfer, error) {
	var err error
	loanevidencetransfer := []LoanEvidenceTransfer{}
	err = db.Debug().Model(&LoanEvidenceTransfer{}).
		Select(`t_loan_evidence_transfer.loan_id,
				t_loan_evidence_transfer.loan_evidence_trf_value,
				t_loan_evidence_transfer.loan_evidence_trf_status,
				t_loan_evidence_transfer.loan_evidence_trf_created_by,
				t_loan_evidence_transfer.loan_evidence_trf_created_at at time zone 'Asia/Jakarta' as loan_evidence_trf_created_at,
				t_loan_evidence_transfer.loan_evidence_trf_updated_by,
				t_loan_evidence_transfer.loan_evidence_trf_updated_at,
				t_loan_evidence_transfer.loan_evidence_trf_deleted_by,
				t_loan_evidence_transfer.loan_evidence_trf_deleted_at`).
		Joins("JOIN t_loan on t_loan_evidence_transfer.loan_id=t_loan.loan_id").
		Joins("JOIN t_loan_submission on t_loan.loan_sbmssn_id=t_loan_submission.loan_sbmssn_id").
		Where("t_loan_submission.loan_req_id = ?", pid).
		Find(&loanevidencetransfer).Error
	if err != nil {
		return &[]LoanEvidenceTransfer{}, err
	}
	return &loanevidencetransfer, nil
}

func (p *LoanEvidenceTransfer) FindLoanEvidenceTransferStatusByID(db *gorm.DB, pid uint64, status uint64) (*LoanEvidenceTransfer, error) {
	var err error
	err = db.Debug().Model(&LoanEvidenceTransfer{}).Where("loan_evidence_trf_id = ? AND referal_status = ?", pid, status).Take(&p).Error
	if err != nil {
		return &LoanEvidenceTransfer{}, err
	}
	return p, nil
}

func (p *LoanEvidenceTransfer) FindLoanEvidenceTransferStatusByCode(db *gorm.DB, code string) (*LoanEvidenceTransfer, error) {
	var err error
	err = db.Debug().Model(&LoanEvidenceTransfer{}).Where("referal_code = ? AND referal_status = ?", code, 1).Take(&p).Error
	if err != nil {
		return &LoanEvidenceTransfer{}, err
	}
	return p, nil
}

func (p *LoanEvidenceTransfer) UpdateLoanEvidenceTransfer(db *gorm.DB, pid uint64) (*LoanEvidenceTransfer, error) {
	var err error
	db = db.Debug().Model(&LoanEvidenceTransfer{}).Where("loan_evidence_trf_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_spk_value":               p.LoanEvidenceTransferValue,
			"loan_evidence_trf_updated_by": p.LoanEvidenceTransferUpdatedBy,
			"loan_evidence_trf_status":     p.LoanEvidenceTransferStatus,
			"loan_evidence_trf_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanEvidenceTransfer{}).Where("loan_evidence_trf_id = ?", pid).Error
	if err != nil {
		return &LoanEvidenceTransfer{}, err
	}
	return p, nil
}

func (p *LoanEvidenceTransfer) UpdateLoanEvidenceTransferStatus(db *gorm.DB, pid uint64) (*LoanEvidenceTransfer, error) {
	var err error
	db = db.Debug().Model(&LoanEvidenceTransfer{}).Where("loan_evidence_trf_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_evidence_trf_status":     p.LoanEvidenceTransferStatus,
			"loan_evidence_trf_updated_by": p.LoanEvidenceTransferUpdatedBy,
			"loan_evidence_trf_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanEvidenceTransfer{}).Where("loan_evidence_trf_id = ?", pid).Error
	if err != nil {
		return &LoanEvidenceTransfer{}, err
	}
	return p, nil
}

func (p *LoanEvidenceTransfer) UpdateLoanEvidenceTransferStatusDelete(db *gorm.DB, pid uint64) (*LoanEvidenceTransfer, error) {
	var err error
	db = db.Debug().Model(&LoanEvidenceTransfer{}).Where("loan_evidence_trf_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_evidence_trf_status":     3,
			"loan_evidence_trf_deleted_by": p.LoanEvidenceTransferDeletedBy,
			"loan_evidence_trf_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanEvidenceTransfer{}).Where("loan_evidence_trf_id = ?", pid).Error
	if err != nil {
		return &LoanEvidenceTransfer{}, err
	}
	return p, nil
}

func (p *LoanEvidenceTransfer) DeleteLoanEvidenceTransfer(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&LoanEvidenceTransfer{}).Where("loan_evidence_trf_id = ?", pid).Take(&LoanEvidenceTransfer{}).Delete(&LoanEvidenceTransfer{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("LoanEvidenceTransfer not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
