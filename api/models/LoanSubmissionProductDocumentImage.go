package models

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type LoanSubmissionProductDocumentImage struct {
	LoanSubmissionProductDocumentImageID        uint64     `gorm:"column:loan_sbmssn_prod_doc_img_id;primary_key;" json:"loan_sbmssn_prod_doc_img_id"`
	LoanSubmissionProductDocumentID             uint64     `gorm:"column:loan_sbmssn_prod_doc_id;" json:"loan_sbmssn_prod_doc_id"`
	LoanSubmissionProductDocumentImageValue     string     `gorm:"column:loan_sbmssn_prod_doc_img_value;size:255" json:"loan_sbmssn_prod_doc_img_value"`
	LoanSubmissionProductDocumentImageStatus    uint64     `gorm:"column:loan_sbmssn_prod_doc_img_status;" json:"loan_sbmssn_prod_doc_img_status"`
	LoanSubmissionProductDocumentImageCreatedAt time.Time  `gorm:"column:loan_sbmssn_prod_doc_img_created_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_prod_doc_img_created_at"`
	LoanSubmissionProductDocumentImageCreatedBy string     `gorm:"column:loan_sbmssn_prod_doc_img_created_by;size:255" json:"loan_sbmssn_prod_doc_img_created_by"`
	LoanSubmissionProductDocumentImageUpdatedAt *time.Time `gorm:"column:loan_sbmssn_prod_doc_img_updated_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_prod_doc_img_created_at"`
	LoanSubmissionProductDocumentImageUpdatedBy string     `gorm:"column:loan_sbmssn_prod_doc_img_updated_by;size:255" json:"loan_sbmssn_prod_doc_img_updated_by"`
	LoanSubmissionProductDocumentImageDeletedAt *time.Time `gorm:"column:loan_sbmssn_prod_doc_img_deleted_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_prod_doc_img_deleted_at"`
	LoanSubmissionProductDocumentImageDeletedBy string     `gorm:"column:loan_sbmssn_prod_doc_img_deleted_by;size:255" json:"loan_sbmssn_prod_doc_img_deleted_by"`
}

type LoanSubmissionProductDocumentImageData struct {
	LoanSubmissionProductDocumentImageID        uint64     `gorm:"column:loan_sbmssn_prod_doc_img_id;" json:"loan_sbmssn_prod_doc_img_id"`
	LoanSubmissionProductDocumentID             uint64     `gorm:"column:loan_sbmssn_prod_doc_id;" json:"loan_sbmssn_prod_doc_id"`
	LoanSubmissionProductDocumentImageValue     string     `gorm:"column:loan_sbmssn_prod_doc_img_value" json:"loan_sbmssn_prod_doc_img_value"`
	LoanSubmissionProductDocumentImageStatus    uint64     `gorm:"column:loan_sbmssn_prod_doc_img_status;" json:"loan_sbmssn_prod_doc_img_status"`
	LoanSubmissionProductDocumentImageCreatedAt time.Time  `gorm:"column:loan_sbmssn_prod_doc_img_created_at" json:"loan_sbmssn_prod_doc_img_created_at"`
	LoanSubmissionProductDocumentImageCreatedBy string     `gorm:"column:loan_sbmssn_prod_doc_img_created_by" json:"loan_sbmssn_prod_doc_img_created_by"`
	LoanSubmissionProductDocumentImageUpdatedAt *time.Time `gorm:"column:loan_sbmssn_prod_doc_img_updated_at" json:"loan_sbmssn_prod_doc_img_created_at"`
	LoanSubmissionProductDocumentImageUpdatedBy string     `gorm:"column:loan_sbmssn_prod_doc_img_updated_by" json:"loan_sbmssn_prod_doc_img_updated_by"`
	LoanSubmissionProductDocumentImageDeletedAt *time.Time `gorm:"column:loan_sbmssn_prod_doc_img_deleted_at" json:"loan_sbmssn_prod_doc_img_deleted_at"`
	LoanSubmissionProductDocumentImageDeletedBy string     `gorm:"column:loan_sbmssn_prod_doc_img_deleted_by" json:"loan_sbmssn_prod_doc_img_deleted_by"`

	LoanDocumentCount int `gorm:"-"`
}

type ResponseLoanSubmissionProductDocumentImages struct {
	Status  int                                       `json:"status"`
	Message string                                    `json:"message"`
	Data    *[]LoanSubmissionProductDocumentImageData `json:"data"`
}

type ResponseLoanSubmissionProductDocumentImage struct {
	Status  int                                     `json:"status"`
	Message string                                  `json:"message"`
	Data    *LoanSubmissionProductDocumentImageData `json:"data"`
}

type ResponseLoanSubmissionProductDocumentImageCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (LoanSubmissionProductDocumentImage) TableName() string {
	return "t_loan_submission_product_document_image"
}

func (LoanSubmissionProductDocumentImageData) TableName() string {
	return "t_loan_submission_product_document_image"
}

func (p *LoanSubmissionProductDocumentImage) Prepare() {
	p.LoanSubmissionProductDocumentID = p.LoanSubmissionProductDocumentID
	p.LoanSubmissionProductDocumentImageValue = p.LoanSubmissionProductDocumentImageValue
	p.LoanSubmissionProductDocumentImageCreatedAt = time.Now()
	p.LoanSubmissionProductDocumentImageCreatedBy = p.LoanSubmissionProductDocumentImageCreatedBy
	p.LoanSubmissionProductDocumentImageUpdatedAt = p.LoanSubmissionProductDocumentImageUpdatedAt
	p.LoanSubmissionProductDocumentImageUpdatedBy = p.LoanSubmissionProductDocumentImageUpdatedBy
	p.LoanSubmissionProductDocumentImageDeletedAt = p.LoanSubmissionProductDocumentImageDeletedAt
	p.LoanSubmissionProductDocumentImageDeletedBy = p.LoanSubmissionProductDocumentImageDeletedBy
}
func (p *LoanSubmissionProductDocumentImage) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.LoanSubmissionProductDocumentID == 0 {
			return errors.New("Required Product Document ID")
		}
		if p.LoanSubmissionProductDocumentImageValue == "" {
			return errors.New("Required Document Image")
		}
		return nil
	}
}

func (p *LoanSubmissionProductDocumentImage) SaveLoanSubmissionProductDocumentImage(db *gorm.DB) (*LoanSubmissionProductDocumentImage, error) {
	var err error
	err = db.Debug().Model(&LoanSubmissionProductDocumentImage{}).Create(&p).Error
	if err != nil {
		return &LoanSubmissionProductDocumentImage{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProductDocumentImage) FindAllLoanSubmissionProductDocumentImage(db *gorm.DB) (*[]LoanSubmissionProductDocumentImageData, error) {
	var err error
	loan_sbmssn_prod_doc_img := []LoanSubmissionProductDocumentImageData{}
	err = db.Debug().Model(&LoanSubmissionProductDocumentImageData{}).Limit(100).
		Select(`t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_id,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_id,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_value,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_status,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_created_at,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_created_by,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_updated_at,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_updated_by,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_deleted_at,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_deleted_by`).
		Find(&loan_sbmssn_prod_doc_img).Error
	if err != nil {
		return &[]LoanSubmissionProductDocumentImageData{}, err
	}
	return &loan_sbmssn_prod_doc_img, nil
}

func (p *LoanSubmissionProductDocumentImage) FindAllLoanSubmissionProductDocumentImageStatus(db *gorm.DB, status uint64) (*[]LoanSubmissionProductDocumentImageData, error) {
	var err error
	loan_sbmssn_prod_doc_img := []LoanSubmissionProductDocumentImageData{}
	err = db.Debug().Model(&LoanSubmissionProductDocumentImageData{}).Limit(100).
		Select(`t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_id,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_id,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_value,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_status,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_created_at,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_created_by,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_updated_at,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_updated_by,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_deleted_at,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_deleted_by`).
		Where("loan_sbmssn_prod_doc_img_status = ?", status).
		Find(&loan_sbmssn_prod_doc_img).Error
	if err != nil {
		return &[]LoanSubmissionProductDocumentImageData{}, err
	}
	return &loan_sbmssn_prod_doc_img, nil
}

func (p *LoanSubmissionProductDocumentImage) FindLoanSubmissionProductDocumentImageDataByID(db *gorm.DB, pid uint64) (*LoanSubmissionProductDocumentImage, error) {
	var err error
	err = db.Debug().Model(&LoanSubmissionProductDocumentImage{}).Where("loan_sbmssn_prod_doc_img_id = ?", pid).Take(&p).Error
	if err != nil {
		return &LoanSubmissionProductDocumentImage{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProductDocumentImage) FindLoanSubmissionProductDocumentImageByID(db *gorm.DB, pid uint64) (*LoanSubmissionProductDocumentImageData, error) {
	var err error
	loan_sbmssn_prod_doc_img := LoanSubmissionProductDocumentImageData{}
	err = db.Debug().Model(&LoanSubmissionProductDocumentImageData{}).Limit(100).
		Select(`t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_id,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_id,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_value,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_status,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_created_at,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_created_by,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_updated_at,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_updated_by,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_deleted_at,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_deleted_by`).
		Where("loan_sbmssn_prod_doc_img_id = ?", pid).
		Find(&loan_sbmssn_prod_doc_img).Error
	if err != nil {
		return &LoanSubmissionProductDocumentImageData{}, err
	}
	return &loan_sbmssn_prod_doc_img, nil
}

func (p *LoanSubmissionProductDocumentImage) FindLoanSubmissionProductDocumentImageStatusByID(db *gorm.DB, pid uint64, status uint64) (*LoanSubmissionProductDocumentImageData, error) {
	var err error
	loan_sbmssn_prod_doc_img := LoanSubmissionProductDocumentImageData{}
	err = db.Debug().Model(&LoanSubmissionProductDocumentImageData{}).Limit(100).
		Select(`t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_id,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_id,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_value,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_status,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_created_at,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_created_by,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_updated_at,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_updated_by,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_deleted_at,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_deleted_by`).
		Joins("JOIN t_loan_submission_product_document on t_loan_submission_product_document_image.loan_sbmssn_prod_doc_id=t_loan_submission_product_document.loan_sbmssn_prod_doc_id").
		Where("t_loan_submission_product_document.loan_sbmssn_prod_doc_id = ? AND loan_sbmssn_prod_doc_img_status = ?", pid, status).
		Find(&loan_sbmssn_prod_doc_img).Error
	if err != nil {
		return &LoanSubmissionProductDocumentImageData{}, err
	}
	return &loan_sbmssn_prod_doc_img, nil
}

func (p *LoanSubmissionProductDocumentImage) UpdateLoanSubmissionProductDocumentImage(db *gorm.DB, pid uint64) (*LoanSubmissionProductDocumentImage, error) {

	var err error
	db = db.Debug().Model(&LoanSubmissionProductDocumentImage{}).Where("loan_sbmssn_prod_doc_img_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_prod_doc_id":             p.LoanSubmissionProductDocumentID,
			"loan_sbmssn_prod_doc_img_value":      p.LoanSubmissionProductDocumentImageValue,
			"loan_sbmssn_prod_doc_img_updated_by": p.LoanSubmissionProductDocumentImageUpdatedBy,
			"loan_sbmssn_prod_doc_img_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionProductDocumentImage{}).Where("loan_sbmssn_prod_doc_img_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionProductDocumentImage{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProductDocumentImage) UpdateLoanSubmissionProductDocumentImageStatus(db *gorm.DB, pid uint64) (*LoanSubmissionProductDocumentImage, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionProductDocumentImage{}).Where("loan_sbmssn_prod_doc_img_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_prod_doc_img_status":     p.LoanSubmissionProductDocumentImageStatus,
			"loan_sbmssn_prod_doc_img_updated_by": p.LoanSubmissionProductDocumentImageUpdatedBy,
			"loan_sbmssn_prod_doc_img_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionProductDocumentImage{}).Where("loan_sbmssn_prod_doc_img_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionProductDocumentImage{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProductDocumentImage) UpdateLoanSubmissionProductDocumentImageStatusDelete(db *gorm.DB, pid uint64) (*LoanSubmissionProductDocumentImage, error) {
	var err error
	db = db.Debug().Model(&LoanSubmissionProductDocumentImage{}).Where("loan_sbmssn_prod_doc_img_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_prod_doc_img_status":     3,
			"loan_sbmssn_prod_doc_img_deleted_by": p.LoanSubmissionProductDocumentImageDeletedBy,
			"loan_sbmssn_prod_doc_img_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmissionProductDocumentImage{}).Where("loan_sbmssn_prod_doc_img_id = ?", pid).Error
	if err != nil {
		return &LoanSubmissionProductDocumentImage{}, err
	}
	return p, nil
}

func (p *LoanSubmissionProductDocumentImage) DeleteLoanSubmissionProductDocumentImage(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&LoanSubmissionProductDocumentImage{}).Where("loan_sbmssn_prod_doc_img_id = ?", pid).Take(&LoanSubmissionProductDocumentImage{}).Delete(&LoanSubmissionProductDocumentImage{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("LoanSubmissionProductDocumentImage not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// =====================================================================================================
// ADDITIONAL

type LoanSubmissionProductDocumentImageJoin struct {
	LoanSubmissionProductDocumentImageID        uint64     `gorm:"column:loan_sbmssn_prod_doc_img_id;primary_key;" json:"loan_sbmssn_prod_doc_img_id"`
	LoanSubmissionProductDocumentID             uint64     `gorm:"column:loan_sbmssn_prod_doc_id; json:"loan_sbmssn_prod_doc_id"`
	LoanSubmissionProductID                     uint64     `gorm:"column:loan_sbmssn_prod_id; json:"loan_sbmssn_prod_id"`
	LoanSubmissionID                            uint64     `gorm:"column:loan_req_id; json:"loan_req_id"`
	LoanReqID                                   uint64     `gorm:"column:loan_id; json:"loan_id"`
	LoanID                                      uint64     `gorm:"column:loan_id; json:"loan_id"`
	LoanSubmissionProductDocumentImageValue     string     `gorm:"column:loan_sbmssn_prod_doc_img_value;size:255" json:"loan_sbmssn_prod_doc_img_value"`
	LoanSubmissionProductDocumentImageStatus    uint64     `gorm:"column:loan_sbmssn_prod_doc_img_status;" json:"loan_sbmssn_prod_doc_img_status"`
	LoanSubmissionProductDocumentImageCreatedAt time.Time  `gorm:"column:loan_sbmssn_prod_doc_img_created_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_prod_doc_img_created_at"`
	LoanSubmissionProductDocumentImageCreatedBy string     `gorm:"column:loan_sbmssn_prod_doc_img_created_by;size:255" json:"loan_sbmssn_prod_doc_img_created_by"`
	LoanSubmissionProductDocumentImageUpdatedAt *time.Time `gorm:"column:loan_sbmssn_prod_doc_img_updated_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_prod_doc_img_created_at"`
	LoanSubmissionProductDocumentImageUpdatedBy string     `gorm:"column:loan_sbmssn_prod_doc_img_updated_by;size:255" json:"loan_sbmssn_prod_doc_img_updated_by"`
	LoanSubmissionProductDocumentImageDeletedAt *time.Time `gorm:"column:loan_sbmssn_prod_doc_img_deleted_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_prod_doc_img_deleted_at"`
	LoanSubmissionProductDocumentImageDeletedBy string     `gorm:"column:loan_sbmssn_prod_doc_img_deleted_by;size:255" json:"loan_sbmssn_prod_doc_img_deleted_by"`
}

type ResponseLoanSubmissionProductDocumentImageJoin struct {
	Status  int                                     `json:"status"`
	Message string                                  `json:"message"`
	Data    *LoanSubmissionProductDocumentImageJoin `json:"data"`
}

func (p *LoanSubmissionProductDocumentImage) FindLoanSubmissionProductDocumentImageByProdID(db *gorm.DB, pid uint64) (*LoanSubmissionProductDocumentImageData, error) {
	var err error
	loan_sbmssn_prod_doc_img := LoanSubmissionProductDocumentImageData{}
	err = db.Debug().Model(&LoanSubmissionProductDocumentImageData{}).Limit(100).
		Select(`t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_id,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_id,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_value,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_status,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_created_at,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_created_by,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_updated_at,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_updated_by,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_deleted_at,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_deleted_by`).
		Where("loan_sbmssn_prod_doc_id = ?", pid).
		Find(&loan_sbmssn_prod_doc_img).Error
	if err != nil {
		return &LoanSubmissionProductDocumentImageData{}, err
	}
	return &loan_sbmssn_prod_doc_img, nil
}

func (p *LoanSubmissionProductDocumentImage) FindLoanSubmissionProductDocumentImageByProdIDCount(db *gorm.DB, pid uint64) (*LoanSubmissionProductDocumentImageData, error) {
	var err error
	loan_sbmssn_prod_doc_img := LoanSubmissionProductDocumentImageData{}
	err = db.Debug().Model(&LoanSubmissionProductDocumentImageData{}).Limit(100).
		Select(`count(distinct(loan_sbmssn_prod_doc_id, loan_sbmssn_prod_doc_img_status))`).
		Where("loan_sbmssn_prod_doc_id = ? AND loan_sbmssn_prod_doc_img_status = ?", pid, 1).
		Find(&loan_sbmssn_prod_doc_img).
		Count(&loan_sbmssn_prod_doc_img.LoanDocumentCount).
		Take(&loan_sbmssn_prod_doc_img).Error
	if err != nil {
		return &LoanSubmissionProductDocumentImageData{}, err
	}
	return &loan_sbmssn_prod_doc_img, nil
}

func (p *LoanSubmissionProductDocumentImage) FindLoanSubmissionByProdDocumentImgID(db *gorm.DB, pid uint64) (*LoanSubmissionProductDocumentImageJoin, error) {
	var err error
	loan_sbmssn_prod_doc_img := LoanSubmissionProductDocumentImageJoin{}
	err = db.Debug().Model(&LoanSubmissionProductDocumentImageJoin{}).Limit(100).
		Select(`t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_id,
				t_loan_submission_product_document.loan_sbmssn_prod_doc_id,
				t_loan_submission_product.loan_sbmssn_prod_id,
				t_loan_submission.loan_sbmssn_id,
				t_loan_submission.loan_req_id,
				t_loan.loan_id,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_id,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_value,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_status,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_created_at,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_updated_at,
				t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_deleted_at`).
		Joins("t_loan_submission_product_document on t_loan_submission_product_document_image.loan_sbmssn_prod_doc_id = t_loan_submission_product_document.loan_sbmssn_prod_doc_id").
		Joins("t_loan_submission_product on t_loan_submission_product_document.loan_sbmssn_prod_id = t_loan_submission_product.loan_sbmssn_prod_id").
		Joins("t_loan_submission on t_loan_submission_product.loan_sbmssn_id = t_loan_submission.loan_sbmssn_id").
		Joins("t_loan on t_loan_submission.loan_sbmssn_id = t_loan.loan_sbmssn_id").
		Where("t_loan_submission_product_document_image.loan_sbmssn_prod_doc_img_id = ?", pid).
		Find(&loan_sbmssn_prod_doc_img).Error
	if err != nil {
		return &LoanSubmissionProductDocumentImageJoin{}, err
	}
	return &loan_sbmssn_prod_doc_img, nil
}
