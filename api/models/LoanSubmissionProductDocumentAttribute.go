package models

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type LoanSubmissionProductDocumentAttribute struct {
	LoanSubmissionProductDocumentAttributeID          uint64     `gorm:"column:loan_sbmssn_prod_doc_attr_id;primary_key;" json:"loan_sbmssn_prod_doc_attr_id"`
	LoanSubmissionProductDocumentID                   uint64     `gorm:"column:loan_sbmssn_prod_doc_id; json:"loan_sbmssn_prod_doc_id"`
	LoanSubmissionProductDocumentAttributeCode        string     `gorm:"column:loan_sbmssn_prod_doc_attr_code;size:15" json:"loan_sbmssn_prod_doc_attr_code"`
	LoanSubmissionProductDocumentAttributeName        string     `gorm:"column:loan_sbmssn_prod_doc_attr_name;size:255" json:"loan_sbmssn_prod_doc_attr_name"`
	LoanSubmissionProductDocumentAttributeDescription string     `gorm:"column:loan_sbmssn_prod_doc_attr_desc" json:"loan_sbmssn_prod_doc_attr_desc"`
	LoanSubmissionProductDocumentAttributeType        string     `gorm:"column:loan_sbmssn_prod_doc_attr_type;size:255" json:"loan_sbmssn_prod_doc_attr_type"`
	LoanSubmissionProductDocumentAttributeStatus      uint64     `gorm:"column:loan_sbmssn_prod_doc_attr_status;size:2" json:"loan_sbmssn_prod_doc_attr_status"`
	LoanSubmissionProductDocumentAttributeCreatedAt   time.Time  `gorm:"column:loan_sbmssn_prod_doc_attr_created_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_prod_doc_attr_created_at"`
	LoanSubmissionProductDocumentAttributeCreatedBy   string     `gorm:"column:loan_sbmssn_prod_doc_attr_created_by;size:125" json:"loan_sbmssn_prod_doc_attr_created_by"`
	LoanSubmissionProductDocumentAttributeUpdatedAt   *time.Time `gorm:"column:loan_sbmssn_prod_doc_attr_updated_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_prod_doc_attr_updated_at"`
	LoanSubmissionProductDocumentAttributeUpdatedBy   string     `gorm:"column:loan_sbmssn_prod_doc_attr_updated_by;size:125" json:"loan_sbmssn_prod_doc_attr_updated_by"`
	LoanSubmissionProductDocumentAttributeDeletedAt   *time.Time `gorm:"column:loan_sbmssn_prod_doc_attr_deleted_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_prod_doc_attr_deleted_at"`
	LoanSubmissionProductDocumentAttributeDeletedBy   string     `gorm:"column:loan_sbmssn_prod_doc_attr_deleted_by;size:125" json:"loan_sbmssn_prod_doc_attr_deleted_by"`
}

type LoanSubmissionProductDocumentAttributeData struct {
	LoanSubmissionProductDocumentAttributeID          uint64     `gorm:"column:loan_sbmssn_prod_doc_attr_id" json:"loan_sbmssn_prod_doc_attr_id"`
	LoanSubmissionProductDocumentID                   uint64     `gorm:"column:loan_sbmssn_prod_doc_id; json:"loan_sbmssn_prod_doc_id"`
	LoanSubmissionProductDocumentAttributeCode        string     `gorm:"column:loan_sbmssn_prod_doc_attr_code" json:"loan_sbmssn_prod_doc_attr_code"`
	LoanSubmissionProductDocumentAttributeName        string     `gorm:"column:loan_sbmssn_prod_doc_attr_name" json:"loan_sbmssn_prod_doc_attr_name"`
	LoanSubmissionProductDocumentAttributeDescription string     `gorm:"column:loan_sbmssn_prod_doc_attr_desc" json:"loan_sbmssn_prod_doc_attr_desc"`
	LoanSubmissionProductDocumentAttributeType        string     `gorm:"column:loan_sbmssn_prod_doc_attr_type" json:"loan_sbmssn_prod_doc_attr_type"`
	LoanSubmissionProductDocumentAttributeStatus      uint64     `gorm:"column:loan_sbmssn_prod_doc_attr_status" json:"loan_sbmssn_prod_doc_attr_status"`
	LoanSubmissionProductDocumentAttributeCreatedAt   time.Time  `gorm:"column:loan_sbmssn_prod_doc_attr_created_at" json:"loan_sbmssn_prod_doc_attr_created_at"`
	LoanSubmissionProductDocumentAttributeCreatedBy   string     `gorm:"column:loan_sbmssn_prod_doc_attr_created_by" json:"loan_sbmssn_prod_doc_attr_created_by"`
	LoanSubmissionProductDocumentAttributeUpdatedAt   *time.Time `gorm:"column:loan_sbmssn_prod_doc_attr_updated_at" json:"loan_sbmssn_prod_doc_attr_updated_at"`
	LoanSubmissionProductDocumentAttributeUpdatedBy   string     `gorm:"column:loan_sbmssn_prod_doc_attr_updated_by" json:"loan_sbmssn_prod_doc_attr_updated_by"`
	LoanSubmissionProductDocumentAttributeDeletedAt   *time.Time `gorm:"column:loan_sbmssn_prod_doc_attr_deleted_at" json:"loan_sbmssn_prod_doc_attr_deleted_at"`
	LoanSubmissionProductDocumentAttributeDeletedBy   string     `gorm:"column:loan_sbmssn_prod_doc_attr_deleted_by" json:"loan_sbmssn_prod_doc_attr_deleted_by"`
}

type ResponseLoanSubmissionProductDocumentAttributes struct {
	Status  int                                           `json:"status"`
	Message string                                        `json:"message"`
	Data    *[]LoanSubmissionProductDocumentAttributeData `json:"data"`
}

type ResponseLoanSubmissionProductDocumentAttribute struct {
	Status  int                                         `json:"status"`
	Message string                                      `json:"message"`
	Data    *LoanSubmissionProductDocumentAttributeData `json:"data"`
}

type ResponseLoanSubmissionProductDocumentAttributeCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (LoanSubmissionProductDocumentAttribute) TableName() string {
	return "t_loan_submission_product_document_attribute"
}

func (LoanSubmissionProductDocumentAttributeData) TableName() string {
	return "t_loan_submission_product_document_attribute"
}

func (p *LoanSubmissionProductDocumentAttribute) Prepare() {
	p.LoanSubmissionProductDocumentID = p.LoanSubmissionProductDocumentID
	p.LoanSubmissionProductDocumentAttributeCode = p.LoanSubmissionProductDocumentAttributeCode
	p.LoanSubmissionProductDocumentAttributeName = p.LoanSubmissionProductDocumentAttributeName
	p.LoanSubmissionProductDocumentAttributeDescription = p.LoanSubmissionProductDocumentAttributeDescription
	p.LoanSubmissionProductDocumentAttributeType = p.LoanSubmissionProductDocumentAttributeType
	p.LoanSubmissionProductDocumentAttributeStatus = p.LoanSubmissionProductDocumentAttributeStatus
	p.LoanSubmissionProductDocumentAttributeCreatedAt = time.Now()
	p.LoanSubmissionProductDocumentAttributeCreatedBy = p.LoanSubmissionProductDocumentAttributeCreatedBy
	p.LoanSubmissionProductDocumentAttributeUpdatedAt = p.LoanSubmissionProductDocumentAttributeUpdatedAt
	p.LoanSubmissionProductDocumentAttributeUpdatedBy = p.LoanSubmissionProductDocumentAttributeUpdatedBy
	p.LoanSubmissionProductDocumentAttributeDeletedAt = p.LoanSubmissionProductDocumentAttributeDeletedAt
	p.LoanSubmissionProductDocumentAttributeDeletedBy = p.LoanSubmissionProductDocumentAttributeDeletedBy
}
func (p *LoanSubmissionProductDocumentAttribute) Validate(action string) error {
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

func (p *LoanSubmissionProductDocumentAttribute) SaveLoanSubmissionProductDocumentAttribute(db *gorm.DB) (*LoanSubmissionProductDocumentAttribute, error) {
	var err error
	err = db.Debug().Model(&LoanSubmissionProductDocumentAttribute{}).Create(&p).Error
	if err != nil {
		return &LoanSubmissionProductDocumentAttribute{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProductDocumentAttribute) FindAllLoanSubmissionProductDocumentAttribute(db *gorm.DB) (*[]LoanSubmissionProductDocumentAttributeData, error) {
	var err error
	prod_doc_mitra_loan_sbmssn_attr_ := []LoanSubmissionProductDocumentAttributeData{}
	err = db.Debug().Model(&LoanSubmissionProductDocumentAttributeData{}).Limit(100).
		Select(`t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_id,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_id,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_code,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_name,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_desc,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_type,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_status,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_created_at,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_created_by,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_updated_at,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_updated_by,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_deleted_at,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_deleted_by`).
		Find(&prod_doc_mitra_loan_sbmssn_attr_).Error
	if err != nil {
		return &[]LoanSubmissionProductDocumentAttributeData{}, err
	}
	return &prod_doc_mitra_loan_sbmssn_attr_, nil
}

func (p *LoanSubmissionProductDocumentAttribute) FindAllLoanSubmissionProductDocumentAttributeStatus(db *gorm.DB, status uint64) (*[]LoanSubmissionProductDocumentAttributeData, error) {
	var err error
	prod_doc_mitra_loan_sbmssn_attr_ := []LoanSubmissionProductDocumentAttributeData{}
	err = db.Debug().Model(&LoanSubmissionProductDocumentAttributeData{}).Limit(100).
		Select(`t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_id,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_id,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_code,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_name,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_desc,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_type,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_status,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_created_at,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_created_by,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_updated_at,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_updated_by,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_deleted_at,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_deleted_by`).
		Where("prod_doc_mitra_loan_sbmssn_attr_status = ?", status).
		Find(&prod_doc_mitra_loan_sbmssn_attr_).Error
	if err != nil {
		return &[]LoanSubmissionProductDocumentAttributeData{}, err
	}
	return &prod_doc_mitra_loan_sbmssn_attr_, nil
}

func (p *LoanSubmissionProductDocumentAttribute) FindLoanSubmissionProductDocumentAttributeDataByID(db *gorm.DB, pid uint64) (*LoanSubmissionProductDocumentAttribute, error) {
	var err error
	err = db.Debug().Model(&LoanSubmissionProductDocumentAttribute{}).Where("prod_doc_mitra_loan_sbmssn_attr_id = ?", pid).Take(&p).Error
	if err != nil {
		return &LoanSubmissionProductDocumentAttribute{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProductDocumentAttribute) FindLoanSubmissionProductDocumentAttributeByID(db *gorm.DB, pid uint64) (*LoanSubmissionProductDocumentAttributeData, error) {
	var err error
	prod_doc_mitra_loan_sbmssn_attr_ := LoanSubmissionProductDocumentAttributeData{}
	err = db.Debug().Model(&LoanSubmissionProductDocumentAttributeData{}).Limit(100).
		Select(`t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_id,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_id,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_code,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_name,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_desc,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_type,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_status,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_created_at,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_created_by,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_updated_at,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_updated_by,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_deleted_at,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_deleted_by`).
		Where("prod_doc_mitra_loan_sbmssn_attr_id = ?", pid).
		Find(&prod_doc_mitra_loan_sbmssn_attr_).Error
	if err != nil {
		return &LoanSubmissionProductDocumentAttributeData{}, err
	}
	return &prod_doc_mitra_loan_sbmssn_attr_, nil
}

func (p *LoanSubmissionProductDocumentAttribute) FindLoanSubmissionProductDocumentAttributeStatusByID(db *gorm.DB, pid uint64, status uint64) (*LoanSubmissionProductDocumentAttributeData, error) {
	var err error
	prod_doc_mitra_loan_sbmssn_attr_ := LoanSubmissionProductDocumentAttributeData{}
	err = db.Debug().Model(&LoanSubmissionProductDocumentAttributeData{}).Limit(100).
		Select(`t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_id,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_id,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_code,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_name,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_desc,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_type,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_status,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_created_at,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_created_by,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_updated_at,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_updated_by,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_deleted_at,
				t_loan_submission_product_document_attribute.prod_doc_mitra_loan_sbmssn_attr_deleted_by`).
		Where("prod_doc_mitra_loan_sbmssn_attr_id = ? AND prod_doc_mitra_loan_sbmssn_attr_status = ?", pid, status).
		Find(&prod_doc_mitra_loan_sbmssn_attr_).Error
	if err != nil {
		return &LoanSubmissionProductDocumentAttributeData{}, err
	}
	return &prod_doc_mitra_loan_sbmssn_attr_, nil
}

func (p *LoanSubmissionProductDocumentAttribute) UpdateLoanSubmissionProductDocumentAttribute(db *gorm.DB, pid uint64) (*LoanSubmissionProductDocumentAttribute, error) {

	var err error
	db = db.Debug().Model(&LoanSubmissionProductDocumentAttribute{}).Where("prod_doc_mitra_loan_sbmssn_attr_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_mitra_loan_sbmssn_id":              p.LoanSubmissionProductDocumentID,
			"prod_doc_mitra_loan_sbmssn_attr_code":       p.LoanSubmissionProductDocumentAttributeCode,
			"prod_doc_mitra_loan_sbmssn_attr_name":       p.LoanSubmissionProductDocumentAttributeName,
			"prod_doc_mitra_loan_sbmssn_attr_desc":       p.LoanSubmissionProductDocumentAttributeDescription,
			"prod_doc_mitra_loan_sbmssn_attr_type":       p.LoanSubmissionProductDocumentAttributeType,
			"prod_doc_mitra_loan_sbmssn_attr_status":     p.LoanSubmissionProductDocumentAttributeStatus,
			"prod_doc_mitra_loan_sbmssn_attr_updated_by": p.LoanSubmissionProductDocumentAttributeUpdatedBy,
			"prod_doc_mitra_loan_sbmssn_attr_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionProductDocumentAttribute{}).Where("prod_doc_mitra_loan_sbmssn_attr_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionProductDocumentAttribute{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProductDocumentAttribute) UpdateLoanSubmissionProductDocumentAttributeStatus(db *gorm.DB, pid uint64) (*LoanSubmissionProductDocumentAttribute, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionProductDocumentAttribute{}).Where("prod_doc_mitra_loan_sbmssn_attr_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_mitra_loan_sbmssn_attr_status":     p.LoanSubmissionProductDocumentAttributeStatus,
			"prod_doc_mitra_loan_sbmssn_attr_updated_by": p.LoanSubmissionProductDocumentAttributeUpdatedBy,
			"prod_doc_mitra_loan_sbmssn_attr_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionProductDocumentAttribute{}).Where("prod_doc_mitra_loan_sbmssn_attr_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionProductDocumentAttribute{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProductDocumentAttribute) UpdateLoanSubmissionProductDocumentAttributeStatusDelete(db *gorm.DB, pid uint64) (*LoanSubmissionProductDocumentAttribute, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionProductDocumentAttribute{}).Where("prod_doc_mitra_loan_sbmssn_attr_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"prod_doc_mitra_loan_sbmssn_attr_status":     3,
			"prod_doc_mitra_loan_sbmssn_attr_deleted_by": p.LoanSubmissionProductDocumentAttributeDeletedBy,
			"prod_doc_mitra_loan_sbmssn_attr_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionProductDocumentAttribute{}).Where("prod_doc_mitra_loan_sbmssn_attr_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionProductDocumentAttribute{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProductDocumentAttribute) DeleteLoanSubmissionProductDocumentAttribute(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&LoanSubmissionProductDocumentAttribute{}).Where("prod_doc_mitra_loan_sbmssn_attr_id = ?", pid).Take(&LoanSubmissionProductDocumentAttribute{}).Delete(&LoanSubmissionProductDocumentAttribute{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("LoanSubmissionProductDocumentAttribute not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
