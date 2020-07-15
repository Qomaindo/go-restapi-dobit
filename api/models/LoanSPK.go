package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type LoanSPK struct {
	LoanSPKID        uint64     `gorm:"column:loan_spk_id;primary_key;" json:"loan_spk_id"`
	LoanID           uint64     `gorm:"column:loan_id" json:"loan_id"`
	LoanSPKValue     string     `gorm:"column:loan_spk_value;size:255" json:"loan_spk_value"`
	LoanSPKStatus    uint64     `gorm:"column:loan_spk_status" json:"loan_spk_status"`
	LoanSPKCreatedBy string     `gorm:"column:loan_spk_created_by;size:125" json:"loan_spk_created_by"`
	LoanSPKCreatedAt time.Time  `gorm:"column:loan_spk_created_at;default:CURRENT_TIMESTAMP" json:"loan_spk_created_at"`
	LoanSPKUpdatedBy string     `gorm:"column:loan_spk_updated_by;size:125" json:"loan_spk_updated_by"`
	LoanSPKUpdatedAt *time.Time `gorm:"column:loan_spk_updated_at;default:CURRENT_TIMESTAMP" json:"loan_spk_updated_at"`
	LoanSPKDeletedBy string     `gorm:"column:loan_spk_deleted_by;size:125" json:"loan_spk_deleted_by"`
	LoanSPKDeletedAt *time.Time `gorm:"column:loan_spk_deleted_at;default:CURRENT_TIMESTAMP" json:"loan_spk_deleted_at"`
}

type ResponseLoanSPKs struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *[]LoanSPK `json:"data"`
}

type ResponseLoanSPK struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    *LoanSPK `json:"data"`
}

type ResponseLoanSPKCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (LoanSPK) TableName() string {
	return "t_loan_spk"
}

func (p *LoanSPK) Prepare() {
	p.LoanID = p.LoanID
	p.LoanSPKValue = html.EscapeString(strings.TrimSpace(p.LoanSPKValue))
	p.LoanSPKStatus = p.LoanSPKStatus
	p.LoanSPKCreatedBy = p.LoanSPKCreatedBy
	p.LoanSPKCreatedAt = time.Now()
	p.LoanSPKUpdatedBy = p.LoanSPKUpdatedBy
	p.LoanSPKUpdatedAt = p.LoanSPKUpdatedAt
	p.LoanSPKDeletedBy = p.LoanSPKDeletedBy
	p.LoanSPKDeletedAt = p.LoanSPKDeletedAt
}

func (p *LoanSPK) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.LoanID == 0 {
			return errors.New("Loan ID Tidak Boleh Kosong")
		}
		if p.LoanSPKValue == "" {
			return errors.New("Loan SPK Value Tidak Boleh Kosong")
		}
		return nil
	}
}

func (p *LoanSPK) SaveLoanSPK(db *gorm.DB) (*LoanSPK, error) {
	var err error
	err = db.Debug().Model(&LoanSPK{}).Create(&p).Error
	if err != nil {
		return &LoanSPK{}, err
	}
	return p, nil
}

func (p *LoanSPK) FindAllLoanSPK(db *gorm.DB) (*[]LoanSPK, error) {
	var err error
	loanspk := []LoanSPK{}
	err = db.Debug().Model(&LoanSPK{}).Limit(100).Find(&loanspk).Error
	if err != nil {
		return &[]LoanSPK{}, err
	}
	return &loanspk, nil
}

func (p *LoanSPK) FindAllLoanSPKStatus(db *gorm.DB, status uint64) (*[]LoanSPK, error) {
	var err error
	loanspk := []LoanSPK{}
	err = db.Debug().Model(&LoanSPK{}).Where("referal_status = ?", status).Limit(100).Find(&loanspk).Error
	if err != nil {
		return &[]LoanSPK{}, err
	}
	return &loanspk, nil
}

func (p *LoanSPK) FindLoanSPKByID(db *gorm.DB, pid uint64) (*LoanSPK, error) {
	var err error
	loanspk := LoanSPK{}
	err = db.Debug().Model(&LoanSPK{}).
		Select(`t_loan_spk.loan_id,
			t_loan_spk.loan_spk_value,
			t_loan_spk.loan_spk_status,
			t_loan_spk.loan_spk_created_by,
			t_loan_spk.loan_spk_created_at at time zone 'Asia/Jakarta' as loan_spk_created_at,
			t_loan_spk.loan_spk_updated_by,
			t_loan_spk.loan_spk_updated_at at time zone 'Asia/Jakarta' as loan_spk_updated_at,
			t_loan_spk.loan_spk_deleted_by,
			t_loan_spk.loan_spk_deleted_at at time zone 'Asia/Jakarta' as loan_spk_deleted_at
	`).
		Joins("JOIN t_loan on t_loan_spk.loan_id=t_loan.loan_id").
		Joins("JOIN t_loan_submission on t_loan.loan_sbmssn_id=t_loan_submission.loan_sbmssn_id").
		// Joins("JOIN t_loan on t_loan_spk.loan_id=t_loan.loan_id").
		Where("t_loan_submission.loan_req_id = ?", pid).
		Find(&loanspk).Error
	if err != nil {
		return &LoanSPK{}, err
	}
	return &loanspk, nil
}

func (p *LoanSPK) FindAllLoanSPKByID(db *gorm.DB, pid uint64) (*[]LoanSPK, error) {
	var err error
	loanspk := []LoanSPK{}
	err = db.Debug().Model(&LoanSPK{}).
		Select(`t_loan_spk.loan_id,
			t_loan_spk.loan_spk_value,
			t_loan_spk.loan_spk_status,
			t_loan_spk.loan_spk_created_by,
			t_loan_spk.loan_spk_created_at at time zone 'Asia/Jakarta' as loan_spk_created_at,
			t_loan_spk.loan_spk_updated_by,
			t_loan_spk.loan_spk_updated_at at time zone 'Asia/Jakarta' as loan_spk_updated_at,
			t_loan_spk.loan_spk_deleted_by,
			t_loan_spk.loan_spk_deleted_at at time zone 'Asia/Jakarta' as loan_spk_deleted_at
	`).
		Joins("JOIN t_loan on t_loan_spk.loan_id=t_loan.loan_id").
		Joins("JOIN t_loan_submission on t_loan.loan_sbmssn_id=t_loan_submission.loan_sbmssn_id").
		// Joins("JOIN t_loan on t_loan_spk.loan_id=t_loan.loan_id").
		Where("t_loan_submission.loan_req_id = ?", pid).
		Find(&loanspk).Error
	if err != nil {
		return &[]LoanSPK{}, err
	}
	return &loanspk, nil
}

func (p *LoanSPK) FindLoanSPKStatusByID(db *gorm.DB, pid uint64, status uint64) (*LoanSPK, error) {
	var err error
	err = db.Debug().Model(&LoanSPK{}).Where("loan_spk_id = ? AND loan_spk_status = ?", pid, status).Take(&p).Error
	if err != nil {
		return &LoanSPK{}, err
	}
	return p, nil
}

func (p *LoanSPK) FindLoanSPKStatusByCode(db *gorm.DB, code string) (*LoanSPK, error) {
	var err error
	err = db.Debug().Model(&LoanSPK{}).Where("referal_code = ? AND loan_spk_status = ?", code, 1).Take(&p).Error
	if err != nil {
		return &LoanSPK{}, err
	}
	return p, nil
}

func (p *LoanSPK) UpdateLoanSPK(db *gorm.DB, pid uint64) (*LoanSPK, error) {
	var err error
	db = db.Debug().Model(&LoanSPK{}).Where("loan_spk_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_spk_value":      p.LoanSPKValue,
			"loan_spk_updated_by": p.LoanSPKUpdatedBy,
			"loan_spk_status":     p.LoanSPKStatus,
			"loan_spk_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSPK{}).Where("loan_spk_id = ?", pid).Error
	if err != nil {
		return &LoanSPK{}, err
	}
	return p, nil
}

func (p *LoanSPK) UpdateLoanSPKStatus(db *gorm.DB, pid uint64) (*LoanSPK, error) {
	var err error
	db = db.Debug().Model(&LoanSPK{}).Where("loan_spk_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_spk_status":     p.LoanSPKStatus,
			"loan_spk_updated_by": p.LoanSPKUpdatedBy,
			"loan_spk_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSPK{}).Where("loan_spk_id = ?", pid).Error
	if err != nil {
		return &LoanSPK{}, err
	}
	return p, nil
}

func (p *LoanSPK) UpdateLoanSPKStatusDelete(db *gorm.DB, pid uint64) (*LoanSPK, error) {
	var err error
	db = db.Debug().Model(&LoanSPK{}).Where("loan_spk_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_spk_status":     3,
			"loan_spk_deleted_by": p.LoanSPKDeletedBy,
			"loan_spk_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSPK{}).Where("loan_spk_id = ?", pid).Error
	if err != nil {
		return &LoanSPK{}, err
	}
	return p, nil
}

func (p *LoanSPK) DeleteLoanSPK(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&LoanSPK{}).Where("loan_spk_id = ?", pid).Take(&LoanSPK{}).Delete(&LoanSPK{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("LoanSPK not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
