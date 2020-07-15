package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type LoanRequestCollateral struct {
	LoanRequestCollateralID        uint64     `gorm:"column:loan_req_coll_id;primary_key;" json:"loan_req_coll_id"`
	LoanRequestID                  uint64     `gorm:"column:loan_req_id;" json:"loan_req_id"`
	LoanSubmissionID               *uint64    `gorm:"column:loan_sbmssn_id;" json:"loan_sbmssn_id"`
	LoanID                         *uint64    `gorm:"column:loan_id" json:"loan_id"`
	LoanRequestCollateralName      string     `gorm:"column:loan_req_coll_name;size:255" json:"loan_req_coll_name"`
	LoanRequestCollateralDesc      string     `gorm:"column:loan_req_coll_desc;size:255" json:"loan_req_coll_desc"`
	LoanRequestCollateralStatus    uint64     `gorm:"column:loan_req_coll_status" json:"loan_req_coll_status"`
	LoanRequestCollateralCreatedBy string     `gorm:"column:loan_req_coll_created_by;size:125" json:"loan_req_coll_created_by"`
	LoanRequestCollateralCreatedAt time.Time  `gorm:"column:loan_req_coll_created_at;default:CURRENT_TIMESTAMP" json:"loan_req_coll_created_at"`
	LoanRequestCollateralUpdatedBy string     `gorm:"column:loan_req_coll_updated_by;size:125" json:"loan_req_coll_updated_by"`
	LoanRequestCollateralUpdatedAt *time.Time `gorm:"column:loan_req_coll_updated_at;default:CURRENT_TIMESTAMP" json:"loan_req_coll_updated_at"`
	LoanRequestCollateralDeletedBy string     `gorm:"column:loan_req_coll_deleted_by;size:125" json:"loan_req_coll_deleted_by"`
	LoanRequestCollateralDeletedAt *time.Time `gorm:"column:loan_req_coll_deleted_at;default:CURRENT_TIMESTAMP" json:"loan_req_coll_deleted_at"`
}

type ResponseLoanRequestCollaterals struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    *[]LoanRequestCollateral `json:"data"`
}

type ResponseLoanRequestCollateral struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    *LoanRequestCollateral `json:"data"`
}

type ResponseLoanRequestCollateralCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (LoanRequestCollateral) TableName() string {
	return "t_loan_request_collateral"
}

func (p *LoanRequestCollateral) Prepare() {
	p.LoanSubmissionID = nil
	p.LoanID = nil
	p.LoanRequestCollateralName = html.EscapeString(strings.TrimSpace(p.LoanRequestCollateralName))
	p.LoanRequestCollateralDesc = html.EscapeString(strings.TrimSpace(p.LoanRequestCollateralDesc))
	p.LoanRequestCollateralStatus = p.LoanRequestCollateralStatus
	p.LoanRequestCollateralCreatedBy = p.LoanRequestCollateralCreatedBy
	p.LoanRequestCollateralCreatedAt = time.Now()
	p.LoanRequestCollateralUpdatedBy = p.LoanRequestCollateralUpdatedBy
	p.LoanRequestCollateralUpdatedAt = p.LoanRequestCollateralUpdatedAt
	p.LoanRequestCollateralDeletedBy = p.LoanRequestCollateralDeletedBy
	p.LoanRequestCollateralDeletedAt = p.LoanRequestCollateralDeletedAt
}

func (p *LoanRequestCollateral) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.LoanRequestCollateralName == "" {
			return errors.New("Required Loan Collateral Name")
		}
		if p.LoanRequestCollateralDesc == "" {
			return errors.New("Required Loan Collateral Description")
		}
		return nil
	}
}

func (p *LoanRequestCollateral) SaveLoanRequestCollateral(db *gorm.DB) (*LoanRequestCollateral, error) {
	var err error
	err = db.Debug().Model(&LoanRequestCollateral{}).Create(&p).Error
	if err != nil {
		return &LoanRequestCollateral{}, err
	}
	return p, nil
}

func (p *LoanRequestCollateral) FindAllLoanRequestCollateral(db *gorm.DB) (*[]LoanRequestCollateral, error) {
	var err error
	loanreqtcollateral := []LoanRequestCollateral{}
	err = db.Debug().Model(&LoanRequestCollateral{}).Limit(100).Find(&loanreqtcollateral).Error
	if err != nil {
		return &[]LoanRequestCollateral{}, err
	}
	return &loanreqtcollateral, nil
}

func (p *LoanRequestCollateral) FindLoanRequestCollateralByReqID(db *gorm.DB, pid uint64) (*LoanRequestCollateral, error) {
	var err error
	loanreqtcollateral := LoanRequestCollateral{}
	err = db.Debug().Model(&LoanRequestCollateral{}).
		Select(`t_loan_request_collateral.loan_req_coll_id,
				t_loan_request_collateral.loan_req_id,
				t_loan_request_collateral.loan_sbmssn_id,
				t_loan_request_collateral.loan_id,
				t_loan_request_collateral.loan_req_coll_name,
				t_loan_request_collateral.loan_req_coll_desc,
				t_loan_request_collateral.loan_req_coll_status,
				t_loan_request_collateral.loan_req_coll_created_by,
				t_loan_request_collateral.loan_req_coll_created_at at time zone 'Asia/Jakarta' as loan_req_coll_created_at,
				t_loan_request_collateral.loan_req_coll_updated_by,
				t_loan_request_collateral.loan_req_coll_updated_at at time zone 'Asia/Jakarta' as loan_req_coll_updated_at,
				t_loan_request_collateral.loan_req_coll_deleted_by,
				t_loan_request_collateral.loan_req_coll_deleted_at at time zone 'Asia/Jakarta' as loan_req_coll_deleted_at
		`).
		Joins("JOIN t_loan_request on t_loan_request_collateral.loan_req_id=t_loan_request.loan_req_id").
		// Joins("JOIN t_loan_submission on t_loan_request_collateral.loan_req_id=t_loan_submission.loan_req_id").
		Joins("JOIN t_loan on t_loan_request_collateral.loan_id=t_loan.loan_id").
		Where("t_loan_request_collateral.loan_req_id = ?", pid).
		Find(&loanreqtcollateral).Error
	if err != nil {
		return &LoanRequestCollateral{}, err
	}
	return &loanreqtcollateral, nil
}

func (p *LoanRequestCollateral) FindAllLoanRequestCollateralByID(db *gorm.DB, pid uint64) (*[]LoanRequestCollateral, error) {
	var err error
	loanreqtcollateral := []LoanRequestCollateral{}
	err = db.Debug().Model(&LoanRequestCollateral{}).
		Select(`t_loan_request_collateral.loan_req_coll_id,
				t_loan_request_collateral.loan_req_id,
				t_loan_request_collateral.loan_sbmssn_id,
				t_loan_request_collateral.loan_id,
				t_loan_request_collateral.loan_req_coll_name,
				t_loan_request_collateral.loan_req_coll_desc,
				t_loan_request_collateral.loan_req_coll_status,
				t_loan_request_collateral.loan_req_coll_created_by,
				t_loan_request_collateral.loan_req_coll_created_at at time zone 'Asia/Jakarta' as loan_req_coll_created_at,
				t_loan_request_collateral.loan_req_coll_updated_by,
				t_loan_request_collateral.loan_req_coll_updated_at at time zone 'Asia/Jakarta' as loan_req_coll_updated_at,
				t_loan_request_collateral.loan_req_coll_deleted_by,
				t_loan_request_collateral.loan_req_coll_deleted_at at time zone 'Asia/Jakarta' as loan_req_coll_deleted_at
		`).
		Joins("JOIN t_loan_request on t_loan_request_collateral.loan_req_id=t_loan_request.loan_req_id").
		Joins("JOIN t_loan_submission on t_loan_request_collateral.loan_req_id=t_loan_submission.loan_req_id").
		Joins("JOIN t_loan on t_loan_request_collateral.loan_id=t_loan.loan_id").
		Where("t_loan_request_collateral.loan_req_coll_id = ?", pid).
		Find(&loanreqtcollateral).Error
	if err != nil {
		return &[]LoanRequestCollateral{}, err
	}
	return &loanreqtcollateral, nil
}

func (p *LoanRequestCollateral) FindLoanRequestCollateralStatusByID(db *gorm.DB, pid uint64, status uint64) (*LoanRequestCollateral, error) {
	var err error
	err = db.Debug().Model(&LoanRequestCollateral{}).Where("loan_req_coll_id = ? AND loan_req_coll_status = ?", pid, status).Take(&p).Error
	if err != nil {
		return &LoanRequestCollateral{}, err
	}
	return p, nil
}

func (p *LoanRequestCollateral) UpdateLoanRequestCollateral(db *gorm.DB, pid uint64) (*LoanRequestCollateral, error) {
	var err error
	db = db.Debug().Model(&LoanRequestCollateral{}).Where("loan_req_coll_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_req_coll_name":       p.LoanRequestCollateralName,
			"loan_req_coll_desc":       p.LoanRequestCollateralDesc,
			"loan_req_coll_updated_by": p.LoanRequestCollateralUpdatedBy,
			"loan_req_coll_status":     p.LoanRequestCollateralStatus,
			"loan_req_coll_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanRequestCollateral{}).Where("loan_req_coll_id = ?", pid).Error
	if err != nil {
		return &LoanRequestCollateral{}, err
	}
	return p, nil
}

func (p *LoanRequestCollateral) UpdateLoanRequestCollateralSbmssnID(db *gorm.DB, pid uint64) (*LoanRequestCollateral, error) {
	var err error
	db = db.Debug().Model(&LoanRequestCollateral{}).Where("loan_req_coll_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_id":           p.LoanSubmissionID,
			"loan_req_coll_status":     p.LoanRequestCollateralStatus,
			"loan_req_coll_updated_by": p.LoanRequestCollateralUpdatedBy,
			"loan_req_coll_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanRequestCollateral{}).Where("loan_req_coll_id = ?", pid).Error
	if err != nil {
		return &LoanRequestCollateral{}, err
	}
	return p, nil
}

func (p *LoanRequestCollateral) UpdateLoanRequestCollateralLoanID(db *gorm.DB, pid uint64) (*LoanRequestCollateral, error) {
	var err error
	db = db.Debug().Model(&LoanRequestCollateral{}).Where("loan_req_coll_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_id":                  p.LoanID,
			"loan_req_coll_status":     p.LoanRequestCollateralStatus,
			"loan_req_coll_updated_by": p.LoanRequestCollateralUpdatedBy,
			"loan_req_coll_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanRequestCollateral{}).Where("loan_req_coll_id = ?", pid).Error
	if err != nil {
		return &LoanRequestCollateral{}, err
	}
	return p, nil
}

func (p *LoanRequestCollateral) UpdateLoanRequestCollateralStatus(db *gorm.DB, pid uint64) (*LoanRequestCollateral, error) {
	var err error
	db = db.Debug().Model(&LoanRequestCollateral{}).Where("loan_req_coll_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_req_coll_status":     p.LoanRequestCollateralStatus,
			"loan_req_coll_updated_by": p.LoanRequestCollateralUpdatedBy,
			"loan_req_coll_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanRequestCollateral{}).Where("loan_req_coll_id = ?", pid).Error
	if err != nil {
		return &LoanRequestCollateral{}, err
	}
	return p, nil
}

func (p *LoanRequestCollateral) UpdateLoanRequestCollateralStatusDelete(db *gorm.DB, pid uint64) (*LoanRequestCollateral, error) {
	var err error
	db = db.Debug().Model(&LoanRequestCollateral{}).Where("loan_req_coll_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_req_coll_status":     3,
			"loan_req_coll_deleted_by": p.LoanRequestCollateralDeletedBy,
			"loan_req_coll_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanRequestCollateral{}).Where("loan_req_coll_id = ?", pid).Error
	if err != nil {
		return &LoanRequestCollateral{}, err
	}
	return p, nil
}

func (p *LoanRequestCollateral) DeleteLoanRequestCollateral(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&LoanRequestCollateral{}).Where("loan_req_coll_id = ?", pid).Take(&LoanRequestCollateral{}).Delete(&LoanRequestCollateral{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Loan Request Collateral not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// ADDITIONAL MEMORANDUM
// ===================================================================================================================================

type LoanRequestCollateralMrdmData struct {
	LoanRequestCollateralID     uint64                             `gorm:"column:loan_req_coll_id" json:"loan_req_coll_id"`
	LoanRequestCollateralName   string                             `gorm:"column:loan_req_coll_name" json:"loan_req_coll_name"`
	LoanRequestCollateralDesc   string                             `gorm:"column:loan_req_coll_desc" json:"loan_req_coll_desc"`
	LoanRequestCollateralStatus uint64                             `gorm:"column:loan_req_coll_status" json:"loan_req_coll_status"`
	LoanRequestCollateralImage  []LoanRequestCollateralImgMrdmData `gorm:"column:loan_req_coll_image" json:"loan_req_coll_image"`
}

func (LoanRequestCollateralMrdmData) TableName() string {
	return "t_loan_request_collateral"
}

func (p *LoanRequestCollateral) FindLoanRequestCollateralMrdmByID(db *gorm.DB, pid uint64) (LoanRequestCollateralMrdmData, error) {
	var err error
	loanreqtcollateralmemorandum := LoanRequestCollateralMrdmData{}
	err = db.Debug().Model(&LoanRequestCollateralMrdmData{}).Limit(100).
		Select(`t_loan_request_collateral.loan_req_coll_id,
				t_loan_request_collateral.loan_req_coll_name,
				t_loan_request_collateral.loan_req_coll_desc,
				t_loan_request_collateral.loan_req_coll_status`).
		Where("t_loan_request_collateral.loan_req_id = ?", pid).
		Find(&loanreqtcollateralmemorandum).Error
	if err != nil {
		return LoanRequestCollateralMrdmData{}, err
	}
	return loanreqtcollateralmemorandum, nil
}
