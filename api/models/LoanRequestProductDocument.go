package models

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type LoanRequestProdDoc struct {
	LoanRequestProdDocID        uint64     `gorm:"column:loan_prod_doc_req_id;primary_key;" json:"loan_prod_doc_req_id"`
	LoanRequestID               uint64     `gorm:"column:loan_req_id; json:"loan_req_id"`
	LoanRequestProdDocCode      string     `gorm:"column:loan_prod_doc_req_code;size:15" json:"loan_prod_doc_req_code"`
	LoanRequestProdDocStatus    uint64     `gorm:"column:loan_prod_doc_req_status;size:2" json:"loan_prod_doc_req_status"`
	LoanRequestProdDocName      string     `gorm:"column:loan_prod_doc_req_name;size:255" json:"loan_prod_doc_req_name"`
	LoanRequestProdDocDesc      string     `gorm:"column:loan_prod_doc_req_desc" json:"loan_prod_doc_req_desc"`
	LoanRequestProdDocCreatedAt time.Time  `gorm:"column:loan_prod_doc_req_created_at;default:CURRENT_TIMESTAMP" json:"loan_prod_doc_req_created_at"`
	LoanRequestProdDocUpdatedAt *time.Time `gorm:"column:loan_prod_doc_req_updated_at;default:CURRENT_TIMESTAMP" json:"loan_prod_doc_req_created_at"`
	LoanRequestProdDocDeletedAt *time.Time `gorm:"column:loan_prod_doc_req_deleted_at;default:CURRENT_TIMESTAMP" json:"loan_prod_doc_req_deleted_at"`
}

type LoanRequestProdDocData struct {
	LoanRequestProdDocID        uint64     `gorm:"column:loan_prod_doc_req_id;primary_key;" json:"loan_prod_doc_req_id"`
	LoanRequestID               uint64     `gorm:"column:loan_req_id; json:"loan_req_id"`
	LoanRequestProdDocCode      string     `gorm:"column:loan_prod_doc_req_code;size:15" json:"loan_prod_doc_req_code"`
	LoanRequestProdDocStatus    uint64     `gorm:"column:loan_prod_doc_req_status;size:2" json:"loan_prod_doc_req_status"`
	LoanRequestProdDocName      string     `gorm:"column:loan_prod_doc_req_name;size:255" json:"loan_prod_doc_req_name"`
	LoanRequestProdDocDesc      string     `gorm:"column:loan_prod_doc_req_desc" json:"loan_prod_doc_req_desc"`
	LoanRequestProdDocCreatedAt time.Time  `gorm:"column:loan_prod_doc_req_created_at;default:CURRENT_TIMESTAMP" json:"loan_prod_doc_req_created_at"`
	LoanRequestProdDocUpdatedAt *time.Time `gorm:"column:loan_prod_doc_req_updated_at;default:CURRENT_TIMESTAMP" json:"loan_prod_doc_req_created_at"`
	LoanRequestProdDocDeletedAt *time.Time `gorm:"column:loan_prod_doc_req_deleted_at;default:CURRENT_TIMESTAMP" json:"loan_prod_doc_req_deleted_at"`
}

type ResponseLoanRequestProdDocs struct {
	Status  int                       `json:"status"`
	Message string                    `json:"message"`
	Data    *[]LoanRequestProdDocData `json:"data"`
}

type ResponseLoanRequestProdDoc struct {
	Status  int                     `json:"status"`
	Message string                  `json:"message"`
	Data    *LoanRequestProdDocData `json:"data"`
}

type ResponseLoanRequestProdDocIU struct {
	Status  int                 `json:"status"`
	Message string              `json:"message"`
	Data    *LoanRequestProdDoc `json:"data"`
}

type ResponseLoanRequestProdDocDel struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (LoanRequestProdDoc) TableName() string {
	return "l_loan_request_product_document"
}

func (LoanRequestProdDocData) TableName() string {
	return "l_loan_request_product_document"
}

func (p *LoanRequestProdDoc) Prepare() {
	p.LoanRequestID = p.LoanRequestID
	p.LoanRequestProdDocCode = p.LoanRequestProdDocCode
	p.LoanRequestProdDocStatus = p.LoanRequestProdDocStatus
	p.LoanRequestProdDocName = p.LoanRequestProdDocName
	p.LoanRequestProdDocDesc = p.LoanRequestProdDocDesc
	p.LoanRequestProdDocCreatedAt = time.Now()
	p.LoanRequestProdDocUpdatedAt = p.LoanRequestProdDocUpdatedAt
	p.LoanRequestProdDocDeletedAt = p.LoanRequestProdDocDeletedAt
}
func (p *LoanRequestProdDoc) Validate(action string) error {
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

func (p *LoanRequestProdDoc) SaveLoanRequestProdDoc(db *gorm.DB) (*LoanRequestProdDoc, error) {
	var err error
	err = db.Debug().Model(&LoanRequestProdDoc{}).Create(&p).Error
	if err != nil {
		return &LoanRequestProdDoc{}, err
	}
	return p, nil
}

func (p *LoanRequestProdDoc) FindAllLoanRequestProdDoc(db *gorm.DB) (*[]LoanRequestProdDocData, error) {
	var err error
	loan_prod_doc_req := []LoanRequestProdDocData{}
	err = db.Debug().Model(&LoanRequestProdDocData{}).Limit(100).
		Select(`l_loan_request_product_document.loan_prod_doc_req_id,
				t_loan_request.loan_req_id,
				l_loan_request_product_document.loan_prod_doc_req_code,
				l_loan_request_product_document.loan_prod_doc_req_status,
				l_loan_request_product_document.loan_prod_doc_req_name,
				l_loan_request_product_document.loan_prod_doc_req_desc,
				l_loan_request_product_document.loan_prod_doc_req_created_at at time zone 'Asia/Jakarta' as loan_prod_doc_req_created_at,
				l_loan_request_product_document.loan_prod_doc_req_updated_at at time zone 'Asia/Jakarta' as loan_prod_doc_req_updated_at,
				l_loan_request_product_document.loan_prod_doc_req_deleted_at at time zone 'Asia/Jakarta' as loan_prod_doc_req_deleted_at`).
		Joins("JOIN t_loan_request on l_loan_request_product_document.loan_req_id=t_loan_request.loan_req_id").
		Find(&loan_prod_doc_req).Error
	if err != nil {
		return &[]LoanRequestProdDocData{}, err
	}
	return &loan_prod_doc_req, nil
}

func (p *LoanRequestProdDoc) FindAllLoanRequestProdDocStatus(db *gorm.DB, status uint64) (*[]LoanRequestProdDocData, error) {
	var err error
	loan_prod_doc_req := []LoanRequestProdDocData{}
	err = db.Debug().Model(&LoanRequestProdDocData{}).Limit(100).
		Select(`l_loan_request_product_document.loan_prod_doc_req_id,
				t_loan_request.loan_req_id,
				l_loan_request_product_document.loan_prod_doc_req_code,
				l_loan_request_product_document.loan_prod_doc_req_status,
				l_loan_request_product_document.loan_prod_doc_req_name,
				l_loan_request_product_document.loan_prod_doc_req_desc,
				l_loan_request_product_document.loan_prod_doc_req_created_at at time zone 'Asia/Jakarta' as loan_prod_doc_req_created_at,
				l_loan_request_product_document.loan_prod_doc_req_updated_at at time zone 'Asia/Jakarta' as loan_prod_doc_req_updated_at,
				l_loan_request_product_document.loan_prod_doc_req_deleted_at at time zone 'Asia/Jakarta' as loan_prod_doc_req_deleted_at`).
		Joins("JOIN t_loan_request on l_loan_request_product_document.loan_req_id=t_loan_request.loan_req_id").
		Where("loan_prod_doc_req_status = ?", status).
		Find(&loan_prod_doc_req).Error
	if err != nil {
		return &[]LoanRequestProdDocData{}, err
	}
	return &loan_prod_doc_req, nil
}

func (p *LoanRequestProdDoc) FindLoanRequestProdDocDataByID(db *gorm.DB, pid uint64) (*LoanRequestProdDoc, error) {
	var err error
	err = db.Debug().Model(&LoanRequestProdDoc{}).Where("loan_prod_doc_req_id = ?", pid).Take(&p).Error
	if err != nil {
		return &LoanRequestProdDoc{}, err
	}
	return p, nil
}

func (p *LoanRequestProdDoc) FindLoanRequestProdDocByID(db *gorm.DB, pid uint64) (*LoanRequestProdDocData, error) {
	var err error
	loan_prod_doc_req := LoanRequestProdDocData{}
	err = db.Debug().Model(&LoanRequestProdDocData{}).Limit(100).
		Select(`l_loan_request_product_document.loan_prod_doc_req_id,
				t_loan_request.loan_req_id,
				l_loan_request_product_document.loan_prod_doc_req_code,
				l_loan_request_product_document.loan_prod_doc_req_status,
				l_loan_request_product_document.loan_prod_doc_req_name,
				l_loan_request_product_document.loan_prod_doc_req_desc,
				l_loan_request_product_document.loan_prod_doc_req_created_at at time zone 'Asia/Jakarta' as loan_prod_doc_req_created_at,
				l_loan_request_product_document.loan_prod_doc_req_updated_at at time zone 'Asia/Jakarta' as loan_prod_doc_req_updated_at,
				l_loan_request_product_document.loan_prod_doc_req_deleted_at at time zone 'Asia/Jakarta' as loan_prod_doc_req_deleted_at`).
		Joins("JOIN t_loan_request on l_loan_request_product_document.loan_req_id=t_loan_request.loan_req_id").
		Where("loan_prod_doc_req_id = ?", pid).
		Find(&loan_prod_doc_req).Error
	if err != nil {
		return &LoanRequestProdDocData{}, err
	}
	return &loan_prod_doc_req, nil
}

func (p *LoanRequestProdDoc) FindLoanRequestProdDocStatusByID(db *gorm.DB, pid uint64, status uint64) (*LoanRequestProdDocData, error) {
	var err error
	loan_prod_doc_req := LoanRequestProdDocData{}
	err = db.Debug().Model(&LoanRequestProdDocData{}).Limit(100).
		Select(`l_loan_request_product_document.loan_prod_doc_req_id,
				t_loan_request.loan_req_id,
				l_loan_request_product_document.loan_prod_doc_req_code,
				l_loan_request_product_document.loan_prod_doc_req_status,
				l_loan_request_product_document.loan_prod_doc_req_name,
				l_loan_request_product_document.loan_prod_doc_req_desc,
				l_loan_request_product_document.loan_prod_doc_req_created_at at time zone 'Asia/Jakarta' as loan_prod_doc_req_created_at,
				l_loan_request_product_document.loan_prod_doc_req_updated_at at time zone 'Asia/Jakarta' as loan_prod_doc_req_updated_at,
				l_loan_request_product_document.loan_prod_doc_req_deleted_at at time zone 'Asia/Jakarta' as loan_prod_doc_req_deleted_at`).
		Joins("JOIN t_loan_request on l_loan_request_product_document.loan_req_id=t_loan_request.loan_req_id").
		Where("loan_prod_doc_req_id = ? AND loan_prod_doc_req_status = ?", pid, status).
		Find(&loan_prod_doc_req).Error
	if err != nil {
		return &LoanRequestProdDocData{}, err
	}
	return &loan_prod_doc_req, nil
}

func (p *LoanRequestProdDoc) UpdateLoanRequestProdDoc(db *gorm.DB, pid uint64) (*LoanRequestProdDoc, error) {

	var err error
	db = db.Debug().Model(&LoanRequestProdDoc{}).Where("loan_prod_doc_req_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_req_id":                  p.LoanRequestID,
			"loan_prod_doc_req_code":       p.LoanRequestProdDocCode,
			"loan_prod_doc_req_name":       p.LoanRequestProdDocName,
			"loan_prod_doc_req_desc":       p.LoanRequestProdDocDesc,
			"loan_prod_doc_req_status":     p.LoanRequestProdDocStatus,
			"loan_prod_doc_req_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanRequestProdDoc{}).Where("loan_prod_doc_req_id = ?", pid).Error
	if err != nil {
		return &LoanRequestProdDoc{}, err
	}
	return p, nil
}

func (p *LoanRequestProdDoc) UpdateLoanRequestProdDocStatus(db *gorm.DB, pid uint64) (*LoanRequestProdDoc, error) {
	var err error
	db = db.Debug().Model(&LoanRequestProdDoc{}).Where("loan_prod_doc_req_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_prod_doc_req_status":     p.LoanRequestProdDocStatus,
			"loan_prod_doc_req_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanRequestProdDoc{}).Where("loan_prod_doc_req_id = ?", pid).Error
	if err != nil {
		return &LoanRequestProdDoc{}, err
	}
	return p, nil
}

func (p *LoanRequestProdDoc) UpdateLoanRequestProdDocStatusDelete(db *gorm.DB, pid uint64) (*LoanRequestProdDoc, error) {
	var err error
	db = db.Debug().Model(&LoanRequestProdDoc{}).Where("loan_prod_doc_req_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_prod_doc_req_status":     3,
			"loan_prod_doc_req_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanRequestProdDoc{}).Where("loan_prod_doc_req_id = ?", pid).Error
	if err != nil {
		return &LoanRequestProdDoc{}, err
	}
	return p, nil
}

func (p *LoanRequestProdDoc) DeleteLoanRequestProdDoc(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&LoanRequestProdDoc{}).Where("loan_prod_doc_req_id = ?", pid).Take(&LoanRequestProdDoc{}).Delete(&LoanRequestProdDoc{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("LoanRequestProdDoc not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
