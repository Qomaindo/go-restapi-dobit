package models

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type LoanSubmissionProductDocument struct {
	LoanSubmissionProductDocumentID                  *uint64    `gorm:"column:loan_sbmssn_prod_doc_id;primary_key;" json:"loan_sbmssn_prod_doc_id"`
	LoanSubmissionProductID                          uint64     `gorm:"column:loan_sbmssn_prod_id" json:"loan_sbmssn_prod_id"`
	LoanSubmissionProductDocumentCategoryName        string     `gorm:"column:loan_sbmssn_prod_doc_ctgry_name;size:255" json:"loan_sbmssn_prod_doc_ctgry_name"`
	LoanSubmissionProductDocumentCategoryDesc        string     `gorm:"column:loan_sbmssn_prod_doc_ctgry_desc" json:"loan_sbmssn_prod_doc_ctgry_desc"`
	LoanSubmissionProductDocumentCustomerTypeName    string     `gorm:"column:loan_sbmssn_prod_doc_cust_type_name;size:125" json:"loan_sbmssn_prod_doc_cust_type_name"`
	LoanSubmissionProductDocumentCustomerSubTypeName string     `gorm:"column:loan_sbmssn_prod_doc_sub_cust_type_name;size:125" json:"loan_sbmssn_prod_doc_sub_cust_type_name"`
	LoanSubmissionProductDocumentIsMandatory         bool       `gorm:"column:loan_sbmssn_prod_doc_is_mandatory" json:"loan_sbmssn_prod_doc_is_mandatory"`
	LoanSubmissionProductDocumentIsManyImage         bool       `gorm:"column:loan_sbmssn_prod_doc_is_many_image" json:"loan_sbmssn_prod_doc_is_many_image"`
	LoanSubmissionProductDocumentCode                string     `gorm:"column:loan_sbmssn_prod_doc_code;size:15" json:"loan_sbmssn_prod_doc_code"`
	LoanSubmissionProductDocumentName                string     `gorm:"column:loan_sbmssn_prod_doc_name;size:255" json:"loan_sbmssn_prod_doc_name"`
	LoanSubmissionProductDocumentStatus              uint64     `gorm:"column:loan_sbmssn_prod_doc_status;size:2" json:"loan_sbmssn_prod_doc_status"`
	LoanSubmissionProductDocumentCreatedAt           time.Time  `gorm:"column:loan_sbmssn_prod_doc_created_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_prod_doc_created_at"`
	LoanSubmissionProductDocumentCreatedBy           string     `gorm:"column:loan_sbmssn_prod_doc_created_by;size:125" json:"loan_sbmssn_prod_doc_created_by"`
	LoanSubmissionProductDocumentUpdatedAt           *time.Time `gorm:"column:loan_sbmssn_prod_doc_updated_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_prod_doc_updated_at"`
	LoanSubmissionProductDocumentUpdatedBy           string     `gorm:"column:loan_sbmssn_prod_doc_updated_by;size:125" json:"loan_sbmssn_prod_doc_updated_by"`
	LoanSubmissionProductDocumentDeletedAt           *time.Time `gorm:"column:loan_sbmssn_prod_doc_deleted_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_prod_doc_deleted_at"`
	LoanSubmissionProductDocumentDeletedBy           string     `gorm:"column:loan_sbmssn_prod_doc_deleted_by;size:125" json:"loan_sbmssn_prod_doc_deleted_by"`
}

type LoanSubmissionProductDocumentData struct {
	LoanSubmissionProductDocumentID                  uint64     `gorm:"column:loan_sbmssn_prod_doc_id" json:"loan_sbmssn_prod_doc_id"`
	LoanSubmissionProductID                          uint64     `gorm:"column:loan_sbmssn_prod_id" json:"loan_sbmssn_prod_id"`
	LoanSubmissionProductDocumentCategoryName        string     `gorm:"column:loan_sbmssn_prod_doc_ctgry_name" json:"loan_sbmssn_prod_doc_ctgry_name"`
	LoanSubmissionProductDocumentCategoryDesc        string     `gorm:"column:loan_sbmssn_prod_doc_ctgry_desc" json:"loan_sbmssn_prod_doc_ctgry_desc"`
	LoanSubmissionProductDocumentCustomerTypeName    string     `gorm:"column:loan_sbmssn_prod_doc_cust_type_name" json:"loan_sbmssn_prod_doc_cust_type_name"`
	LoanSubmissionProductDocumentCustomerSubTypeName string     `gorm:"column:loan_sbmssn_prod_doc_sub_cust_type_name" json:"loan_sbmssn_prod_doc_sub_cust_type_name"`
	LoanSubmissionProductDocumentIsMandatory         bool       `gorm:"column:loan_sbmssn_prod_doc_is_mandatory" json:"loan_sbmssn_prod_doc_is_mandatory"`
	LoanSubmissionProductDocumentIsManyImage         bool       `gorm:"column:loan_sbmssn_prod_doc_is_many_image" json:"loan_sbmssn_prod_doc_is_many_image"`
	LoanSubmissionProductDocumentName                string     `gorm:"column:loan_sbmssn_prod_doc_name" json:"loan_sbmssn_prod_doc_name"`
	LoanSubmissionProductDocumentStatus              uint64     `gorm:"column:loan_sbmssn_prod_doc_status" json:"loan_sbmssn_prod_doc_status"`
	LoanSubmissionProductDocumentCreatedAt           time.Time  `gorm:"column:loan_sbmssn_prod_doc_created_at" json:"loan_sbmssn_prod_doc_created_at"`
	LoanSubmissionProductDocumentCreatedBy           string     `gorm:"column:loan_sbmssn_prod_doc_created_by" json:"loan_sbmssn_prod_doc_created_by"`
	LoanSubmissionProductDocumentUpdatedAt           *time.Time `gorm:"column:loan_sbmssn_prod_doc_updated_at" json:"loan_sbmssn_prod_doc_updated_at"`
	LoanSubmissionProductDocumentUpdatedBy           string     `gorm:"column:loan_sbmssn_prod_doc_updated_by" json:"loan_sbmssn_prod_doc_updated_by"`
	LoanSubmissionProductDocumentDeletedAt           *time.Time `gorm:"column:loan_sbmssn_prod_doc_deleted_at" json:"loan_sbmssn_prod_doc_deleted_at"`
	LoanSubmissionProductDocumentDeletedBy           string     `gorm:"column:loan_sbmssn_prod_doc_deleted_by" json:"loan_sbmssn_prod_doc_deleted_by"`
}

type ResponseLoanSubmissionProductDocuments struct {
	Status  int                                  `json:"status"`
	Message string                               `json:"message"`
	Data    *[]LoanSubmissionProductDocumentData `json:"data"`
}

type ResponseLoanSubmissionProductDocument struct {
	Status  int                                `json:"status"`
	Message string                             `json:"message"`
	Data    *LoanSubmissionProductDocumentData `json:"data"`
}

type ResponseLoanSubmissionProductDocumentCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (LoanSubmissionProductDocument) TableName() string {
	return "t_loan_submission_product_document"
}

func (LoanSubmissionProductDocumentData) TableName() string {
	return "t_loan_submission_product_document"
}

func (p *LoanSubmissionProductDocument) Prepare() {
	p.LoanSubmissionProductDocumentID = nil
	p.LoanSubmissionProductID = p.LoanSubmissionProductID
	p.LoanSubmissionProductDocumentCategoryName = p.LoanSubmissionProductDocumentCategoryName
	p.LoanSubmissionProductDocumentCategoryDesc = p.LoanSubmissionProductDocumentCategoryDesc
	p.LoanSubmissionProductDocumentCustomerTypeName = p.LoanSubmissionProductDocumentCustomerTypeName
	p.LoanSubmissionProductDocumentCustomerSubTypeName = p.LoanSubmissionProductDocumentCustomerSubTypeName
	p.LoanSubmissionProductDocumentIsMandatory = p.LoanSubmissionProductDocumentIsMandatory
	p.LoanSubmissionProductDocumentIsManyImage = p.LoanSubmissionProductDocumentIsManyImage
	p.LoanSubmissionProductDocumentCode = p.LoanSubmissionProductDocumentCode
	p.LoanSubmissionProductDocumentStatus = p.LoanSubmissionProductDocumentStatus
	p.LoanSubmissionProductDocumentName = p.LoanSubmissionProductDocumentName
	p.LoanSubmissionProductDocumentCreatedAt = time.Now()
	p.LoanSubmissionProductDocumentCreatedBy = p.LoanSubmissionProductDocumentCreatedBy
	p.LoanSubmissionProductDocumentUpdatedAt = p.LoanSubmissionProductDocumentUpdatedAt
	p.LoanSubmissionProductDocumentUpdatedBy = p.LoanSubmissionProductDocumentUpdatedBy
	p.LoanSubmissionProductDocumentDeletedAt = p.LoanSubmissionProductDocumentDeletedAt
	p.LoanSubmissionProductDocumentDeletedBy = p.LoanSubmissionProductDocumentDeletedBy
}
func (p *LoanSubmissionProductDocument) Validate(action string) error {
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

func (p *LoanSubmissionProductDocument) SaveLoanSubmissionProductDocument(db *gorm.DB) (*LoanSubmissionProductDocument, error) {
	var err error
	err = db.Debug().Model(&LoanSubmissionProductDocument{}).Create(&p).Error
	if err != nil {
		return &LoanSubmissionProductDocument{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProductDocument) FindAllLoanSubmissionProductDocument(db *gorm.DB) (*[]LoanSubmissionProductDocumentData, error) {
	var err error
	loan_sbmssn_prod_doc := []LoanSubmissionProductDocumentData{}
	err = db.Debug().Model(&LoanSubmissionProductDocumentData{}).Limit(100).
		Select(`t_loan_submission_product_document.loan_sbmssn_prod_doc_id,
				t_loan_submission_product_document.loan_sbmssn_prod_id,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_ctgry_name,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_ctgry_desc,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_cust_type_name,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_sub_cust_type_name,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_is_mandatory,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_is_many_image,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_code,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_name,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_status,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_created_at,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_created_by,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_updated_at,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_updated_by,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_deleted_at,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_deleted_by`).
		Find(&loan_sbmssn_prod_doc).Error
	if err != nil {
		return &[]LoanSubmissionProductDocumentData{}, err
	}
	return &loan_sbmssn_prod_doc, nil
}

func (p *LoanSubmissionProductDocument) FindAllLoanSubmissionProductDocumentStatus(db *gorm.DB, status uint64) (*[]LoanSubmissionProductDocumentData, error) {
	var err error
	loan_sbmssn_prod_doc := []LoanSubmissionProductDocumentData{}
	err = db.Debug().Model(&LoanSubmissionProductDocumentData{}).Limit(100).
		Select(`t_loan_submission_product_document.loan_sbmssn_prod_doc_id,
				t_loan_submission_product_document.loan_sbmssn_prod_id,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_ctgry_name,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_ctgry_desc,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_cust_type_name,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_sub_cust_type_name,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_is_mandatory,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_is_many_image,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_code,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_name,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_status,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_created_at,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_created_by,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_updated_at,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_updated_by,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_deleted_at,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_deleted_by`).
		Where("loan_sbmssn_prod_doc_status = ?", status).
		Find(&loan_sbmssn_prod_doc).Error
	if err != nil {
		return &[]LoanSubmissionProductDocumentData{}, err
	}
	return &loan_sbmssn_prod_doc, nil
}

func (p *LoanSubmissionProductDocument) FindAllLoanSubmissionProductDocumentByLoanSubmissionID(db *gorm.DB, pid uint64) (*[]LoanSubmissionProductDocumentData, error) {
	var err error
	loan_sbmssn_prod_doc := []LoanSubmissionProductDocumentData{}
	err = db.Debug().Model(&LoanSubmissionProductDocumentData{}).Limit(100).
		Select(`t_loan_submission_product_document.loan_sbmssn_prod_doc_id,
				t_loan_submission_product_document.loan_sbmssn_prod_id,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_ctgry_name,																																																																																																																												
				t_loan_submission_product_document.loan_sbmssn_prod_doc_ctgry_desc,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_cust_type_name,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_sub_cust_type_name,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_is_mandatory,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_is_many_image,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_code,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_name,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_status,																												
				t_loan_submission_product_document.loan_sbmssn_prod_doc_created_at,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_created_by,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_updated_at,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_updated_by,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_deleted_at,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_deleted_by`).
		Joins("JOIN t_loan_submission_product on t_loan_submission_product_document.loan_sbmssn_prod_id=t_loan_submission_product.loan_sbmssn_prod_id").
		Where("t_loan_submission_product.loan_sbmssn_id = ? ", pid).
		Find(&loan_sbmssn_prod_doc).Error
	if err != nil {
		return &[]LoanSubmissionProductDocumentData{}, err
	}
	return &loan_sbmssn_prod_doc, nil
}

func (p *LoanSubmissionProductDocument) FindLoanSubmissionProductDocumentDataByID(db *gorm.DB, pid uint64) (*LoanSubmissionProductDocument, error) {
	var err error
	err = db.Debug().Model(&LoanSubmissionProductDocument{}).Where("loan_sbmssn_prod_doc_id = ?", pid).Take(&p).Error
	if err != nil {
		return &LoanSubmissionProductDocument{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProductDocument) FindLoanSubmissionProductDocumentByID(db *gorm.DB, pid uint64) (*LoanSubmissionProductDocumentData, error) {
	var err error
	loan_sbmssn_prod_doc := LoanSubmissionProductDocumentData{}
	err = db.Debug().Model(&LoanSubmissionProductDocumentData{}).Limit(100).
		Select(`t_loan_submission_product_document.loan_sbmssn_prod_doc_id,
				t_loan_submission_product_document.loan_sbmssn_prod_id,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_ctgry_name,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_ctgry_desc,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_cust_type_name,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_sub_cust_type_name,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_is_mandatory,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_is_many_image,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_code,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_name,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_status,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_created_at,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_created_by,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_updated_at,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_updated_by,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_deleted_at,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_deleted_by`).
		Where("loan_sbmssn_prod_doc_id = ?", pid).
		Find(&loan_sbmssn_prod_doc).Error
	if err != nil {
		return &LoanSubmissionProductDocumentData{}, err
	}
	return &loan_sbmssn_prod_doc, nil
}

func (p *LoanSubmissionProductDocument) FindLoanSubmissionProductDocumentStatusByID(db *gorm.DB, pid uint64, status uint64) (*LoanSubmissionProductDocumentData, error) {
	var err error
	loan_sbmssn_prod_doc := LoanSubmissionProductDocumentData{}
	err = db.Debug().Model(&LoanSubmissionProductDocumentData{}).Limit(100).
		Select(`t_loan_submission_product_document.loan_sbmssn_prod_doc_id,
				t_loan_submission_product_document.loan_sbmssn_prod_id,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_ctgry_name,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_ctgry_desc,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_cust_type_name,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_sub_cust_type_name,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_is_mandatory,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_is_many_image,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_code,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_name,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_status,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_created_at,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_created_by,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_updated_at,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_updated_by,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_deleted_at,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_deleted_by`).
		Where("loan_sbmssn_prod_doc_id = ? AND loan_sbmssn_prod_doc_status = ?", pid, status).
		Find(&loan_sbmssn_prod_doc).Error
	if err != nil {
		return &LoanSubmissionProductDocumentData{}, err
	}
	return &loan_sbmssn_prod_doc, nil
}

func (p *LoanSubmissionProductDocument) UpdateLoanSubmissionProductDocument(db *gorm.DB, pid uint64) (*LoanSubmissionProductDocument, error) {

	var err error
	db = db.Debug().Model(&LoanSubmissionProductDocument{}).Where("loan_sbmssn_prod_doc_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_prod_id":                     p.LoanSubmissionProductID,
			"loan_sbmssn_prod_doc_ctgry_name":         p.LoanSubmissionProductDocumentCategoryName,
			"loan_sbmssn_prod_doc_ctgry_desc":         p.LoanSubmissionProductDocumentCategoryDesc,
			"loan_sbmssn_prod_doc_cust_type_name":     p.LoanSubmissionProductDocumentCustomerTypeName,
			"loan_sbmssn_prod_doc_sub_cust_type_name": p.LoanSubmissionProductDocumentCustomerSubTypeName,
			"loan_sbmssn_prod_doc_is_mandatory":       p.LoanSubmissionProductDocumentIsMandatory,
			"loan_sbmssn_prod_doc_is_many_image":      p.LoanSubmissionProductDocumentIsManyImage,
			"loan_sbmssn_prod_doc_code":               p.LoanSubmissionProductDocumentCode,
			"loan_sbmssn_prod_doc_name":               p.LoanSubmissionProductDocumentName,
			"loan_sbmssn_prod_doc_status":             p.LoanSubmissionProductDocumentStatus,
			"loan_sbmssn_prod_doc_updated_by":         p.LoanSubmissionProductDocumentUpdatedBy,
			"loan_sbmssn_prod_doc_updated_at":         time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionProductDocument{}).Where("loan_sbmssn_prod_doc_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionProductDocument{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProductDocument) UpdateLoanSubmissionProductDocumentStatus(db *gorm.DB, pid uint64) (*LoanSubmissionProductDocument, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionProductDocument{}).Where("loan_sbmssn_prod_doc_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_prod_doc_status":     p.LoanSubmissionProductDocumentStatus,
			"loan_sbmssn_prod_doc_updated_by": p.LoanSubmissionProductDocumentUpdatedBy,
			"loan_sbmssn_prod_doc_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionProductDocument{}).Where("loan_sbmssn_prod_doc_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionProductDocument{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProductDocument) UpdateLoanSubmissionProductDocumentStatusDelete(db *gorm.DB, pid uint64) (*LoanSubmissionProductDocument, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionProductDocument{}).Where("loan_sbmssn_prod_doc_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_prod_doc_status":     3,
			"loan_sbmssn_prod_doc_deleted_by": p.LoanSubmissionProductDocumentDeletedBy,
			"loan_sbmssn_prod_doc_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionProductDocument{}).Where("loan_sbmssn_prod_doc_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionProductDocument{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProductDocument) DeleteLoanSubmissionProductDocument(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&LoanSubmissionProductDocument{}).Where("loan_sbmssn_prod_doc_id = ?", pid).Take(&LoanSubmissionProductDocument{}).Delete(&LoanSubmissionProductDocument{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("LoanSubmissionProductDocument not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// ADDITIONAL MEMORANDUM
// ===================================================================================================================================

type LoanSubmissionProductDocumentMrdmData struct {
	LoanSubmissionProductDocumentID                  uint64 `gorm:"column:loan_sbmssn_prod_doc_id;" json:"loan_sbmssn_prod_doc_id"`
	LoanSubmissionProductDocumentCategoryName        string `gorm:"column:loan_sbmssn_prod_doc_ctgry_name" json:"loan_sbmssn_prod_doc_ctgry_name`
	LoanSubmissionProductDocumentCategoryDesc        string `gorm:"column:loan_sbmssn_prod_doc_ctgry_desc;" json:"loan_sbmssn_prod_doc_ctgry_desc"`
	LoanSubmissionProductDocumentCustomerTypeName    string `gorm:"column:loan_sbmssn_prod_doc_cust_type_name" json:"loan_sbmssn_prod_doc_cust_type_name"`
	LoanSubmissionProductDocumentCustomerSubTypeName string `gorm:"column:loan_sbmssn_prod_doc_sub_cust_type_name" json:"loan_sbmssn_prod_doc_sub_cust_type_name"`
	LoanSubmissionProductDocumentIsMandatory         bool   `gorm:"column:loan_sbmssn_prod_doc_is_mandatory" json:"loan_sbmssn_prod_doc_is_mandatory"`
	LoanSubmissionProductDocumentIsManyImage         bool   `gorm:"column:loan_sbmssn_prod_doc_is_many_image" json:"loan_sbmssn_prod_doc_is_many_image"`
	LoanSubmissionProductDocumentCode                string `gorm:"column:loan_sbmssn_prod_doc_code" json:"loan_sbmssn_prod_doc_code"`
	LoanSubmissionProductDocumentName                string `gorm:"column:loan_sbmssn_prod_doc_name" json:"loan_sbmssn_prod_doc_name"`
	LoanSubmissionProductDocumentStatus              uint64 `gorm:"column:loan_sbmssn_prod_doc_status" json:"loan_sbmssn_prod_doc_status"`
}

func (LoanSubmissionProductDocumentMrdmData) TableName() string {
	return "t_loan_submission_product_document"
}

func (p *LoanSubmissionProductDocument) FindLoanSubmissionProductDocumentMrdmByLoanSubmissionProductID(db *gorm.DB, pid uint64) ([]LoanSubmissionProductDocumentMrdmData, error) {
	var err error
	loan_sbmssn_prod_doc := []LoanSubmissionProductDocumentMrdmData{}
	err = db.Debug().Model(&LoanSubmissionProductDocumentMrdmData{}).Limit(100).
		Select(`t_loan_submission_product_document.loan_sbmssn_prod_doc_id,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_ctgry_name,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_ctgry_desc,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_cust_type_name,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_sub_cust_type_name,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_is_mandatory,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_is_many_image,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_code,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_name,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_status`).
		Where("t_loan_submission_product_document.loan_sbmssn_prod_id = ? AND loan_sbmssn_prod_doc_status = ?", pid, 1).
		Find(&loan_sbmssn_prod_doc).Error
	if err != nil {
		return []LoanSubmissionProductDocumentMrdmData{}, err
	}
	return loan_sbmssn_prod_doc, nil
}
