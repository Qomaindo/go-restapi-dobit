package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type LoanRequestCollateralImg struct {
	LoanRequestCollateralImgID    uint64 `gorm:"column:loan_req_coll_img_id;primary_key;" json:"loan_req_coll_img_id"`
	LoanRequestCollateralID       uint64 `gorm:"column:loan_req_coll_id;" json:"loan_req_coll_id"`
	LoanRequestCollateralImgValue string `gorm:"column:loan_req_coll_img_value;size:255" json:"loan_req_coll_img_value"`
	// LoanRequestCollateralImgStatus    uint64     `gorm:"column:loan_req_coll_img_status" json:"loan_req_coll_img_status"`
	LoanRequestCollateralImgCreatedBy string     `gorm:"column:loan_req_coll_img_created_by;size:125" json:"loan_req_coll_img_created_by"`
	LoanRequestCollateralImgCreatedAt time.Time  `gorm:"column:loan_req_coll_img_created_at;default:CURRENT_TIMESTAMP" json:"loan_req_coll_img_created_at"`
	LoanRequestCollateralImgUpdatedBy string     `gorm:"column:loan_req_coll_img_updated_by;size:125" json:"loan_req_coll_img_updated_by"`
	LoanRequestCollateralImgUpdatedAt *time.Time `gorm:"column:loan_req_coll_img_updated_at;default:CURRENT_TIMESTAMP" json:"loan_req_coll_img_updated_at"`
	LoanRequestCollateralImgDeletedBy string     `gorm:"column:loan_req_coll_img_deleted_by;size:125" json:"loan_req_coll_img_deleted_by"`
	LoanRequestCollateralImgDeletedAt *time.Time `gorm:"column:loan_req_coll_img_deleted_at;default:CURRENT_TIMESTAMP" json:"loan_req_coll_img_deleted_at"`
}

type LoanRequestCollateralImgData struct {
	LoanRequestCollateralImgID    uint64 `gorm:"column:loan_req_coll_img_id;primary_key;" json:"loan_req_coll_img_id"`
	LoanRequestCollateralID       uint64 `gorm:"column:loan_req_coll_id;" json:"loan_req_coll_id"`
	LoanSubmissionID              uint64 `gorm:"column:loan_sbmssn_id;" json:"loan_sbmssn_id"`
	LoanRequestID                 uint64 `gorm:"column:loan_req_id;" json:"loan_req_id"`
	LoanID                        uint64 `gorm:"column:loan_id" json:"loan_id"`
	LoanRequestCollateralName     string `gorm:"column:loan_req_coll_name;size:255" json:"loan_req_coll_name"`
	LoanRequestCollateralDesc     string `gorm:"column:loan_req_coll_desc;size:255" json:"loan_req_coll_desc"`
	LoanRequestCollateralImgValue string `gorm:"column:loan_req_coll_img_value;size:255" json:"loan_req_coll_img_value"`
	// LoanRequestCollateralImgStatus    uint64     `gorm:"column:loan_req_coll_img_status" json:"loan_req_coll_img_status"`
	LoanRequestCollateralImgCreatedBy string     `gorm:"column:loan_req_coll_img_created_by;size:125" json:"loan_req_coll_img_created_by"`
	LoanRequestCollateralImgCreatedAt time.Time  `gorm:"column:loan_req_coll_img_created_at;default:CURRENT_TIMESTAMP" json:"loan_req_coll_img_created_at"`
	LoanRequestCollateralImgUpdatedBy string     `gorm:"column:loan_req_coll_img_updated_by;size:125" json:"loan_req_coll_img_updated_by"`
	LoanRequestCollateralImgUpdatedAt *time.Time `gorm:"column:loan_req_coll_img_updated_at;default:CURRENT_TIMESTAMP" json:"loan_req_coll_img_updated_at"`
	LoanRequestCollateralImgDeletedBy string     `gorm:"column:loan_req_coll_img_deleted_by;size:125" json:"loan_req_coll_img_deleted_by"`
	LoanRequestCollateralImgDeletedAt *time.Time `gorm:"column:loan_req_coll_img_deleted_at;default:CURRENT_TIMESTAMP" json:"loan_req_coll_img_deleted_at"`
}

type ResponseLoanRequestCollateralImgs struct {
	Status  int                             `json:"status"`
	Message string                          `json:"message"`
	Data    *[]LoanRequestCollateralImgData `json:"data"`
}

type ResponseLoanRequestCollateralImg struct {
	Status  int                       `json:"status"`
	Message string                    `json:"message"`
	Data    *LoanRequestCollateralImg `json:"data"`
}

type ResponseLoanRequestCollateralImgCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (LoanRequestCollateralImg) TableName() string {
	return "t_loan_request_collateral_img"
}

func (LoanRequestCollateralImgData) TableName() string {
	return "t_loan_request_collateral_img"
}

func (p *LoanRequestCollateralImg) Prepare() {
	p.LoanRequestCollateralImgValue = html.EscapeString(strings.TrimSpace(p.LoanRequestCollateralImgValue))
	p.LoanRequestCollateralImgCreatedBy = p.LoanRequestCollateralImgCreatedBy
	p.LoanRequestCollateralImgCreatedAt = time.Now()
	p.LoanRequestCollateralImgUpdatedBy = p.LoanRequestCollateralImgUpdatedBy
	p.LoanRequestCollateralImgUpdatedAt = p.LoanRequestCollateralImgUpdatedAt
	p.LoanRequestCollateralImgDeletedBy = p.LoanRequestCollateralImgDeletedBy
	p.LoanRequestCollateralImgDeletedAt = p.LoanRequestCollateralImgDeletedAt
}

func (p *LoanRequestCollateralImg) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.LoanRequestCollateralImgValue == "" {
			return errors.New("Required Loan Collateral Name")
		}
		return nil
	}
}

func (p *LoanRequestCollateralImg) SaveLoanRequestCollateralImg(db *gorm.DB) (*LoanRequestCollateralImg, error) {
	var err error
	err = db.Debug().Model(&LoanRequestCollateralImg{}).Create(&p).Error
	if err != nil {
		return &LoanRequestCollateralImg{}, err
	}
	return p, nil
}

func (p *LoanRequestCollateralImg) FindAllLoanRequestCollateralImg(db *gorm.DB) (*[]LoanRequestCollateralImg, error) {
	var err error
	loanreqtcollateral := []LoanRequestCollateralImg{}
	err = db.Debug().Model(&LoanRequestCollateralImg{}).Limit(100).Find(&loanreqtcollateral).Error
	if err != nil {
		return &[]LoanRequestCollateralImg{}, err
	}
	return &loanreqtcollateral, nil
}

func (p *LoanRequestCollateralImg) FindLoanRequestCollateralImgByCollID(db *gorm.DB, pid uint64) (*LoanRequestCollateralImg, error) {
	var err error
	loanreqtcollateral := LoanRequestCollateralImg{}
	err = db.Debug().Model(&LoanRequestCollateralImg{}).
		Select(`t_loan_request_collateral_img.loan_req_coll_img_id,
				t_loan_request_collateral.loan_req_coll_id,
				t_loan_request_collateral.loan_req_id,
				t_loan_request_collateral.loan_sbmssn_id,
				t_loan_request_collateral.loan_id,
				t_loan_request_collateral.loan_req_coll_name,
				t_loan_request_collateral.loan_req_coll_desc,
				t_loan_request_collateral_img.loan_req_coll_img_value,
				t_loan_request_collateral_img.loan_req_coll_img_created_by,
				t_loan_request_collateral_img.loan_req_coll_img_created_at at time zone 'Asia/Jakarta' as loan_req_coll_img_created_at,
				t_loan_request_collateral_img.loan_req_coll_img_updated_by,
				t_loan_request_collateral_img.loan_req_coll_img_updated_at at time zone 'Asia/Jakarta' as loan_req_coll_img_updated_at,
				t_loan_request_collateral_img.loan_req_coll_img_deleted_by,
				t_loan_request_collateral_img.loan_req_coll_img_deleted_at at time zone 'Asia/Jakarta' as loan_req_coll_img_deleted_at
		`).
		Joins("JOIN t_loan_request on t_loan_request_collateral_img.loan_req_coll_id=t_loan_request.loan_req_coll_id").
		Where("t_loan_request_collateral_img.loan_req_coll_id = ?", pid).
		Find(&loanreqtcollateral).Error
	if err != nil {
		return &LoanRequestCollateralImg{}, err
	}
	return &loanreqtcollateral, nil
}

func (p *LoanRequestCollateralImg) FindAllLoanRequestCollateralImgByID(db *gorm.DB, pid uint64) (*[]LoanRequestCollateralImgData, error) {
	var err error
	loanreqtcollateral := []LoanRequestCollateralImgData{}
	err = db.Debug().Model(&LoanRequestCollateralImgData{}).
		Select(`t_loan_request_collateral_img.loan_req_coll_img_id,
				t_loan_request_collateral.loan_req_coll_id,
				t_loan_request_collateral.loan_req_id,
				t_loan_request_collateral.loan_sbmssn_id,
				t_loan_request_collateral.loan_id,
				t_loan_request_collateral.loan_req_coll_name,
				t_loan_request_collateral.loan_req_coll_desc,
				t_loan_request_collateral_img.loan_req_coll_img_value,
				t_loan_request_collateral_img.loan_req_coll_img_created_by,
				t_loan_request_collateral_img.loan_req_coll_img_created_at at time zone 'Asia/Jakarta' as loan_req_coll_img_created_at,
				t_loan_request_collateral_img.loan_req_coll_img_updated_by,
				t_loan_request_collateral_img.loan_req_coll_img_updated_at at time zone 'Asia/Jakarta' as loan_req_coll_img_updated_at,
				t_loan_request_collateral_img.loan_req_coll_img_deleted_by,
				t_loan_request_collateral_img.loan_req_coll_img_deleted_at at time zone 'Asia/Jakarta' as loan_req_coll_img_deleted_at
		`).
		Joins("JOIN t_loan_request_collateral on t_loan_request_collateral_img.loan_req_coll_id=t_loan_request_collateral.loan_req_coll_id").
		// Joins("JOIN t_loan_request on t_loan_request_collateral.loan_req_id=t_loan_request.loan_req_id").
		Where("t_loan_request_collateral.loan_req_id = ?", pid).
		Find(&loanreqtcollateral).Error
	if err != nil {
		return &[]LoanRequestCollateralImgData{}, err
	}
	return &loanreqtcollateral, nil
}

func (p *LoanRequestCollateralImg) FindLoanRequestCollateralImgStatusByID(db *gorm.DB, pid uint64, status uint64) (*LoanRequestCollateralImg, error) {
	var err error
	err = db.Debug().Model(&LoanRequestCollateralImg{}).Where("loan_req_coll_img_id = ? AND loan_req_coll_img_status = ?", pid, status).Take(&p).Error
	if err != nil {
		return &LoanRequestCollateralImg{}, err
	}
	return p, nil
}

func (p *LoanRequestCollateralImg) UpdateLoanRequestCollateralImg(db *gorm.DB, pid uint64) (*LoanRequestCollateralImg, error) {
	var err error
	db = db.Debug().Model(&LoanRequestCollateralImg{}).Where("loan_req_coll_img_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_req_coll_img_value":      p.LoanRequestCollateralImgValue,
			"loan_req_coll_img_updated_by": p.LoanRequestCollateralImgUpdatedBy,
			"loan_req_coll_img_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanRequestCollateralImg{}).Where("loan_req_coll_img_id = ?", pid).Error
	if err != nil {
		return &LoanRequestCollateralImg{}, err
	}
	return p, nil
}

func (p *LoanRequestCollateralImg) UpdateLoanRequestCollateralImgStatus(db *gorm.DB, pid uint64) (*LoanRequestCollateralImg, error) {
	var err error
	db = db.Debug().Model(&LoanRequestCollateralImg{}).Where("loan_req_coll_img_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_req_coll_img_updated_by": p.LoanRequestCollateralImgUpdatedBy,
			"loan_req_coll_img_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanRequestCollateralImg{}).Where("loan_req_coll_img_id = ?", pid).Error
	if err != nil {
		return &LoanRequestCollateralImg{}, err
	}
	return p, nil
}

func (p *LoanRequestCollateralImg) UpdateLoanRequestCollateralImgStatusDelete(db *gorm.DB, pid uint64) (*LoanRequestCollateralImg, error) {
	var err error
	db = db.Debug().Model(&LoanRequestCollateralImg{}).Where("loan_req_coll_img_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_req_coll_img_status":     3,
			"loan_req_coll_img_deleted_by": p.LoanRequestCollateralImgDeletedBy,
			"loan_req_coll_img_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanRequestCollateralImg{}).Where("loan_req_coll_img_id = ?", pid).Error
	if err != nil {
		return &LoanRequestCollateralImg{}, err
	}
	return p, nil
}

func (p *LoanRequestCollateralImg) DeleteLoanRequestCollateralImg(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&LoanRequestCollateralImg{}).Where("loan_req_coll_img_id = ?", pid).Take(&LoanRequestCollateralImg{}).Delete(&LoanRequestCollateralImg{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Loan Request Collateral Image not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// ADDITIONAL MEMORANDUM
// ===================================================================================================================================

type LoanRequestCollateralImgMrdmData struct {
	LoanRequestCollateralImgID        uint64 `gorm:"column:loan_req_coll_img_id;" json:"loan_req_coll_img_id"`
	LoanRequestCollateralID           uint64 `gorm:"column:loan_req_coll_id;" json:"loan_req_coll_id"`
	LoanRequestCollateralImgValue     string `gorm:"column:loan_req_coll_img_value" json:"loan_req_coll_img_value"`
	LoanRequestCollateralImgCreatedBy string `gorm:"column:loan_req_coll_img_created_by" json:"loan_req_coll_img_created_by"`
}

func (LoanRequestCollateralImgMrdmData) TableName() string {
	return "t_loan_request_collateral_img"
}

func (p *LoanRequestCollateralImg) FindAllLoanRequestCollateralImgMrdmByLoanRequestCollateralID(db *gorm.DB, pid uint64) ([]LoanRequestCollateralImgMrdmData, error) {
	var err error
	loanreqtcollateralimage := []LoanRequestCollateralImgMrdmData{}
	err = db.Debug().Model(&LoanRequestCollateralImgMrdmData{}).
		Select(`t_loan_request_collateral_img.loan_req_coll_img_id,
				t_loan_request_collateral_img.loan_req_coll_id,
				t_loan_request_collateral_img.loan_req_coll_img_value,
				t_loan_request_collateral_img.loan_req_coll_img_created_by`).
		Where("t_loan_request_collateral_img.loan_req_coll_id = ?", pid).
		Find(&loanreqtcollateralimage).Error
	if err != nil {
		return []LoanRequestCollateralImgMrdmData{}, err
	}
	return loanreqtcollateralimage, nil
}
