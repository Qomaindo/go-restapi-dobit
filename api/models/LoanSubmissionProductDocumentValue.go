package models

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type LoanSubmissionProductDocumentValue struct {
	LoanSubmissionProductDocumentValueID        uint64     `gorm:"column:loan_sbmssn_prod_doc_val_id;primary_key" json:"loan_sbmssn_prod_doc_val_id"`
	LoanSubmissionProductDocumentAttributeID    uint64     `gorm:"column:loan_sbmssn_prod_doc_attr_id" json:"loan_sbmssn_prod_doc_attr_id"`
	LoanSubmissionProductDocumentValueValue     string     `gorm:"column:loan_sbmssn_prod_doc_val_value;size:255" json:"loan_sbmssn_prod_doc_val_value"`
	LoanSubmissionProductDocumentValueStatus    uint64     `gorm:"column:loan_sbmssn_prod_doc_val_status" json:"loan_sbmssn_prod_doc_val_status"`
	LoanSubmissionProductDocumentValueCreatedAt time.Time  `gorm:"column:loan_sbmssn_prod_doc_val_created_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_prod_doc_val_created_at"`
	LoanSubmissionProductDocumentValueCreatedBy string     `gorm:"column:loan_sbmssn_prod_doc_val_created_by;size:125" json:"loan_sbmssn_prod_doc_val_created_by"`
	LoanSubmissionProductDocumentValueUpdatedAt *time.Time `gorm:"column:loan_sbmssn_prod_doc_val_updated_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_prod_doc_val_created_at"`
	LoanSubmissionProductDocumentValueUpdatedBy string     `gorm:"column:loan_sbmssn_prod_doc_val_updated_by;size:125" json:"loan_sbmssn_prod_doc_val_updated_by"`
	LoanSubmissionProductDocumentValueDeletedAt *time.Time `gorm:"column:loan_sbmssn_prod_doc_val_deleted_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_prod_doc_val_deleted_at"`
	LoanSubmissionProductDocumentValueDeletedBy string     `gorm:"column:loan_sbmssn_prod_doc_val_deleted_by;size:125" json:"loan_sbmssn_prod_doc_val_deleted_by"`
}

type LoanSubmissionProductDocumentValueData struct {
	LoanSubmissionProductDocumentValueID        uint64     `gorm:"column:loan_sbmssn_prod_doc_val_id" json:"loan_sbmssn_prod_doc_val_id"`
	LoanSubmissionProductDocumentAttributeID    uint64     `gorm:"column:loan_sbmssn_prod_doc_attr_id" json:"loan_sbmssn_prod_doc_attr_id"`
	LoanSubmissionProductDocumentValueValue     string     `gorm:"column:loan_sbmssn_prod_doc_val_value" json:"loan_sbmssn_prod_doc_val_value"`
	LoanSubmissionProductDocumentValueStatus    uint64     `gorm:"column:loan_sbmssn_prod_doc_val_status" json:"loan_sbmssn_prod_doc_val_status"`
	LoanSubmissionProductDocumentValueCreatedAt time.Time  `gorm:"column:loan_sbmssn_prod_doc_val_created_at" json:"loan_sbmssn_prod_doc_val_created_at"`
	LoanSubmissionProductDocumentValueCreatedBy string     `gorm:"column:loan_sbmssn_prod_doc_val_created_by" json:"loan_sbmssn_prod_doc_val_created_by"`
	LoanSubmissionProductDocumentValueUpdatedAt *time.Time `gorm:"column:loan_sbmssn_prod_doc_val_updated_at" json:"loan_sbmssn_prod_doc_val_created_at"`
	LoanSubmissionProductDocumentValueUpdatedBy string     `gorm:"column:loan_sbmssn_prod_doc_val_updated_by" json:"loan_sbmssn_prod_doc_val_updated_by"`
	LoanSubmissionProductDocumentValueDeletedAt *time.Time `gorm:"column:loan_sbmssn_prod_doc_val_deleted_at" json:"loan_sbmssn_prod_doc_val_deleted_at"`
	LoanSubmissionProductDocumentValueDeletedBy string     `gorm:"column:loan_sbmssn_prod_doc_val_deleted_by" json:"loan_sbmssn_prod_doc_val_deleted_by"`
}

type ResponseLoanSubmissionProductDocumentValues struct {
	Status  int                                       `json:"status"`
	Message string                                    `json:"message"`
	Data    *[]LoanSubmissionProductDocumentValueData `json:"data"`
}

type ResponseLoanSubmissionProductDocumentValue struct {
	Status  int                                     `json:"status"`
	Message string                                  `json:"message"`
	Data    *LoanSubmissionProductDocumentValueData `json:"data"`
}

type ResponseLoanSubmissionProductDocumentValueCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (LoanSubmissionProductDocumentValue) TableName() string {
	return "t_loan_submission_product_document_value"
}

func (LoanSubmissionProductDocumentValueData) TableName() string {
	return "t_loan_submission_product_document_value"
}

func (p *LoanSubmissionProductDocumentValue) Prepare() {
	p.LoanSubmissionProductDocumentAttributeID = p.LoanSubmissionProductDocumentAttributeID
	p.LoanSubmissionProductDocumentValueValue = p.LoanSubmissionProductDocumentValueValue
	p.LoanSubmissionProductDocumentValueStatus = p.LoanSubmissionProductDocumentValueStatus
	p.LoanSubmissionProductDocumentValueCreatedAt = time.Now()
	p.LoanSubmissionProductDocumentValueCreatedBy = p.LoanSubmissionProductDocumentValueCreatedBy
	p.LoanSubmissionProductDocumentValueUpdatedAt = p.LoanSubmissionProductDocumentValueUpdatedAt
	p.LoanSubmissionProductDocumentValueUpdatedBy = p.LoanSubmissionProductDocumentValueUpdatedBy
	p.LoanSubmissionProductDocumentValueDeletedAt = p.LoanSubmissionProductDocumentValueDeletedAt
	p.LoanSubmissionProductDocumentValueDeletedBy = p.LoanSubmissionProductDocumentValueDeletedBy
}
func (p *LoanSubmissionProductDocumentValue) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		// if p.CountryID == 0 {
		// 	return errors.New("Required Region Country")
		// }
		// if p.ProvinceID == 0 {
		// 	return errors.New("Required Region Province")
		// }
		// if p.RegencyID == 0 {
		// 	return errors.New("Required Region Regency")
		// }
		// if p.DistrictID == 0 {
		// 	return errors.New("Required Region District")
		// }
		// if p.VillageID == 0 {
		// 	return errors.New("Required Region Village")
		// }
		return nil
	}
}

func (p *LoanSubmissionProductDocumentValue) SaveLoanSubmissionProductDocumentValue(db *gorm.DB) (*LoanSubmissionProductDocumentValue, error) {
	var err error
	err = db.Debug().Model(&LoanSubmissionProductDocumentValue{}).Create(&p).Error
	if err != nil {
		return &LoanSubmissionProductDocumentValue{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProductDocumentValue) FindAllLoanSubmissionProductDocumentValue(db *gorm.DB) (*[]LoanSubmissionProductDocumentValueData, error) {
	var err error
	loan_sbmssn_prod_doc_val := []LoanSubmissionProductDocumentValueData{}
	err = db.Debug().Model(&LoanSubmissionProductDocumentValueData{}).Limit(100).
		Select(`t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_id,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_attr_id,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_value,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_status,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_created_at,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_created_by,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_updated_at,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_updated_by,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_deleted_at,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_deleted_by`).
		Find(&loan_sbmssn_prod_doc_val).Error
	if err != nil {
		return &[]LoanSubmissionProductDocumentValueData{}, err
	}
	return &loan_sbmssn_prod_doc_val, nil
}

func (p *LoanSubmissionProductDocumentValue) FindAllLoanSubmissionProductDocumentValueStatus(db *gorm.DB, status uint64) (*[]LoanSubmissionProductDocumentValueData, error) {
	var err error
	loan_sbmssn_prod_doc_val := []LoanSubmissionProductDocumentValueData{}
	err = db.Debug().Model(&LoanSubmissionProductDocumentValueData{}).Limit(100).
		Select(`t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_id,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_attr_id,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_value,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_status,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_created_at,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_created_by,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_updated_at,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_updated_by,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_deleted_at,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_deleted_by`).
		Where("loan_sbmssn_prod_doc_val_status = ?", status).
		Find(&loan_sbmssn_prod_doc_val).Error
	if err != nil {
		return &[]LoanSubmissionProductDocumentValueData{}, err
	}
	return &loan_sbmssn_prod_doc_val, nil
}

func (p *LoanSubmissionProductDocumentValue) FindLoanSubmissionProductDocumentValueDataByID(db *gorm.DB, pid uint64) (*LoanSubmissionProductDocumentValue, error) {
	var err error
	err = db.Debug().Model(&LoanSubmissionProductDocumentValue{}).Where("loan_sbmssn_prod_doc_val_id = ?", pid).Take(&p).Error
	if err != nil {
		return &LoanSubmissionProductDocumentValue{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProductDocumentValue) FindLoanSubmissionProductDocumentValueByID(db *gorm.DB, pid uint64) (*LoanSubmissionProductDocumentValueData, error) {
	var err error
	loan_sbmssn_prod_doc_val := LoanSubmissionProductDocumentValueData{}
	err = db.Debug().Model(&LoanSubmissionProductDocumentValueData{}).Limit(100).
		Select(`t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_id,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_attr_id,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_value,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_status,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_created_at,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_created_by,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_updated_at,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_updated_by,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_deleted_at,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_deleted_by`).
		Where("loan_sbmssn_prod_doc_val_id = ?", pid).
		Find(&loan_sbmssn_prod_doc_val).Error
	if err != nil {
		return &LoanSubmissionProductDocumentValueData{}, err
	}
	return &loan_sbmssn_prod_doc_val, nil
}

func (p *LoanSubmissionProductDocumentValue) FindLoanSubmissionProductDocumentValueStatusByID(db *gorm.DB, pid uint64, status uint64) (*LoanSubmissionProductDocumentValueData, error) {
	var err error
	loan_sbmssn_prod_doc_val := LoanSubmissionProductDocumentValueData{}
	err = db.Debug().Model(&LoanSubmissionProductDocumentValueData{}).Limit(100).
		Select(`t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_id,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_attr_id,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_value,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_status,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_created_at,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_created_by,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_updated_at,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_updated_by,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_deleted_at,
				t_loan_submission_product_document_value.loan_sbmssn_prod_doc_val_deleted_by`).
		Where("loan_sbmssn_prod_doc_val_id = ? AND loan_sbmssn_prod_doc_val_status = ?", pid, status).
		Find(&loan_sbmssn_prod_doc_val).Error
	if err != nil {
		return &LoanSubmissionProductDocumentValueData{}, err
	}
	return &loan_sbmssn_prod_doc_val, nil
}

func (p *LoanSubmissionProductDocumentValue) UpdateLoanSubmissionProductDocumentValue(db *gorm.DB, pid uint64) (*LoanSubmissionProductDocumentValue, error) {

	var err error
	db = db.Debug().Model(&LoanSubmissionProductDocumentValue{}).Where("loan_sbmssn_prod_doc_val_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_prod_doc_attr_id":        p.LoanSubmissionProductDocumentAttributeID,
			"loan_sbmssn_prod_doc_val_value":      p.LoanSubmissionProductDocumentValueValue,
			"loan_sbmssn_prod_doc_val_updated_by": p.LoanSubmissionProductDocumentValueUpdatedBy,
			"loan_sbmssn_prod_doc_val_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionProductDocumentValue{}).Where("loan_sbmssn_prod_doc_val_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionProductDocumentValue{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProductDocumentValue) UpdateLoanSubmissionProductDocumentValueStatus(db *gorm.DB, pid uint64) (*LoanSubmissionProductDocumentValue, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionProductDocumentValue{}).Where("loan_sbmssn_prod_doc_val_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_prod_doc_val_status":     p.LoanSubmissionProductDocumentValueStatus,
			"loan_sbmssn_prod_doc_val_updated_by": p.LoanSubmissionProductDocumentValueUpdatedBy,
			"loan_sbmssn_prod_doc_val_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionProductDocumentValue{}).Where("loan_sbmssn_prod_doc_val_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionProductDocumentValue{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProductDocumentValue) UpdateLoanSubmissionProductDocumentValueStatusDelete(db *gorm.DB, pid uint64) (*LoanSubmissionProductDocumentValue, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionProductDocumentValue{}).Where("loan_sbmssn_prod_doc_val_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_prod_doc_val_status":     3,
			"loan_sbmssn_prod_doc_val_deleted_by": p.LoanSubmissionProductDocumentValueDeletedBy,
			"loan_sbmssn_prod_doc_val_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionProductDocumentValue{}).Where("loan_sbmssn_prod_doc_val_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionProductDocumentValue{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProductDocumentValue) DeleteLoanSubmissionProductDocumentValue(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&LoanSubmissionProductDocumentValue{}).Where("loan_sbmssn_prod_doc_val_id = ?", pid).Take(&LoanSubmissionProductDocumentValue{}).Delete(&LoanSubmissionProductDocumentValue{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("LoanSubmissionProductDocumentValue not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
