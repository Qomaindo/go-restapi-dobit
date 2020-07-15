package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

//BASE CRUD
//====================================================================================================================================

type MitraProduct struct {
	MitraProductID                  uint64     `gorm:"column:mitra_prod_id;primary_key;" json:"mitra_prod_id"`
	MitraID                         uint64     `gorm:"column:mitra_id" json:"mitra_id"`
	CustomerTypeID                  uint64     `gorm:"column:cust_type_id" json:"cust_type_id"`
	MitraProductName                string     `gorm:"column:mitra_prod_name;size:125" json:"mitra_prod_name"`
	MitraProductLimitMin            uint64     `gorm:"column:mitra_prod_limit_min" json:"mitra_prod_limit_min"`
	MitraProductLimitMax            uint64     `gorm:"column:mitra_prod_limit_max" json:"mitra_prod_limit_max"`
	MitraProductTenorMin            uint64     `gorm:"column:mitra_prod_tenor_min" json:"mitra_prod_tenor_min"`
	MitraProductTenorMax            uint64     `gorm:"column:mitra_prod_tenor_max" json:"mitra_prod_tenor_max"`
	MitraProductInterestRate        float64    `gorm:"column:mitra_prod_interest_rate" json:"mitra_prod_interest_rate"`
	MitraProductInterestRatePeriod  string     `gorm:"column:mitra_prod_interest_rate_period;size:125" json:"mitra_prod_interest_rate_period"`
	MitraProductInterestRateType    string     `gorm:"column:mitra_prod_interest_rate_type;size:125" json:"mitra_prod_interest_rate_type"`
	MitraProductInterestRateFormula string     `gorm:"column:mitra_prod_interest_rate_formula;size:125" json:"mitra_prod_interest_rate_formula"`
	MitraProductInterestRateDiscMax float64    `gorm:"column:mitra_prod_interest_rate_disc_max" json:"mitra_prod_interest_rate_disc_max"`
	MitraProductNumReviewer         uint64     `gorm:"column:mitra_prod_num_reviewer" json:"mitra_prod_num_reviewer"`
	MitraProductNumApprover         uint64     `gorm:"column:mitra_prod_num_approver" json:"mitra_prod_num_approver"`
	MitraProductDateActiveStart     string     `gorm:"column:mitra_prod_date_active_start;size:25" json:"mitra_prod_date_active_start"`
	MitraProductDateActiveFinish    string     `gorm:"column:mitra_prod_date_active_finish;size:25" json:"mitra_prod_date_active_finish"`
	MitraProductStatus              uint64     `gorm:"column:mitra_prod_status" json:"mitra_prod_status"`
	MitraProductCreatedBy           string     `gorm:"column:mitra_prod_created_by;size:125" json:"mitra_prod_created_by"`
	MitraProductCreatedAt           time.Time  `gorm:"column:mitra_prod_created_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_created_at"`
	MitraProductUpdatedBy           string     `gorm:"column:mitra_prod_updated_by;size:125" json:"mitra_prod_updated_by"`
	MitraProductUpdatedAt           *time.Time `gorm:"column:mitra_prod_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_updated_at"`
	MitraProductDeletedBy           string     `gorm:"column:mitra_prod_deleted_by;size:125" json:"mitra_prod_deleted_by"`
	MitraProductDeletedAt           *time.Time `gorm:"column:mitra_prod_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_deleted_at"`
}

type MitraProductData struct {
	MitraProductID                  uint64     `gorm:"column:mitra_prod_id;primary_key;" json:"mitra_prod_id"`
	MitraID                         uint64     `gorm:"column:mitra_id" json:"mitra_id"`
	MitraCode                       string     `gorm:"column:mitra_code" json:"mitra_code"`
	MitraName                       string     `gorm:"column:mitra_name" json:"mitra_name"`
	MitraWebsite                    string     `gorm:"column:mitra_website" json:"mitra_website"`
	MitraEmail                      string     `gorm:"column:mitra_email" json:"mitra_email"`
	MitraLogo                       string     `gorm:"column:mitra_logo" json:"mitra_logo"`
	CustomerTypeID                  uint64     `gorm:"column:cust_type_id" json:"cust_type_id"`
	CustomerTypeName                string     `gorm:"column:cust_type_name" json:"cust_type_name"`
	MitraProductName                string     `gorm:"column:mitra_prod_name" json:"mitra_prod_name"`
	MitraProductLimitMin            uint64     `gorm:"column:mitra_prod_limit_min" json:"mitra_prod_limit_min"`
	MitraProductLimitMax            uint64     `gorm:"column:mitra_prod_limit_max" json:"mitra_prod_limit_max"`
	MitraProductTenorMin            uint64     `gorm:"column:mitra_prod_tenor_min" json:"mitra_prod_tenor_min"`
	MitraProductTenorMax            uint64     `gorm:"column:mitra_prod_tenor_max" json:"mitra_prod_tenor_max"`
	MitraProductInterestRate        float64    `gorm:"column:mitra_prod_interest_rate" json:"mitra_prod_interest_rate"`
	MitraProductInterestRatePeriod  string     `gorm:"column:mitra_prod_interest_rate_period" json:"mitra_prod_interest_rate_period"`
	MitraProductInterestRateType    string     `gorm:"column:mitra_prod_interest_rate_type" json:"mitra_prod_interest_rate_type"`
	MitraProductInterestRateFormula string     `gorm:"column:mitra_prod_interest_rate_formula" json:"mitra_prod_interest_rate_formula"`
	MitraProductInterestRateDiscMax float64    `gorm:"column:mitra_prod_interest_rate_disc_max" json:"mitra_prod_interest_rate_disc_max"`
	MitraProductNumReviewer         uint64     `gorm:"column:mitra_prod_num_reviewer" json:"mitra_prod_num_reviewer"`
	MitraProductNumApprover         uint64     `gorm:"column:mitra_prod_num_approver" json:"mitra_prod_num_approver"`
	MitraProductDateActiveStart     string     `gorm:"column:mitra_prod_date_active_start" json:"mitra_prod_date_active_start"`
	MitraProductDateActiveFinish    string     `gorm:"column:mitra_prod_date_active_finish" json:"mitra_prod_date_active_finish"`
	MitraProductStatus              uint64     `gorm:"column:mitra_prod_status" json:"mitra_prod_status"`
	MitraProductCreatedBy           string     `gorm:"column:mitra_prod_created_by" json:"mitra_prod_created_by"`
	MitraProductCreatedAt           time.Time  `gorm:"column:mitra_prod_created_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_created_at"`
	MitraProductUpdatedBy           string     `gorm:"column:mitra_prod_updated_by" json:"mitra_prod_updated_by"`
	MitraProductUpdatedAt           *time.Time `gorm:"column:mitra_prod_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_updated_at"`
	MitraProductDeletedBy           string     `gorm:"column:mitra_prod_deleted_by" json:"mitra_prod_deleted_by"`
	MitraProductDeletedAt           *time.Time `gorm:"column:mitra_prod_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_deleted_at"`

	MitraProductSubCategory  []MitraProductSubCategoryData  `json:"mitra_prod_sub_ctgry"`
	MitraProductCost         []MitraProductCostData         `json:"mitra_prod_cost"`
	MitraProductDocumentRule []MitraProductDocumentRuleData `json:"mitra_prod_doc_rule"`
}

type ResponseMitraProducts struct {
	Status  int                 `json:"status"`
	Message string              `json:"message"`
	Data    *[]MitraProductData `json:"data"`
}

type ResponseMitraProduct struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    *MitraProductData `json:"data"`
}

type ResponseMitraProductCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraProduct) TableName() string {
	return "m_mitra_product"
}

func (MitraProductData) TableName() string {
	return "m_mitra_product"
}

func (p *MitraProduct) Prepare() {
	p.MitraID = p.MitraID
	p.CustomerTypeID = p.CustomerTypeID
	p.MitraProductName = html.EscapeString(strings.TrimSpace(p.MitraProductName))
	p.MitraProductLimitMin = p.MitraProductLimitMin
	p.MitraProductLimitMax = p.MitraProductLimitMax
	p.MitraProductTenorMin = p.MitraProductTenorMin
	p.MitraProductTenorMax = p.MitraProductTenorMax
	p.MitraProductInterestRate = p.MitraProductInterestRate
	p.MitraProductInterestRateType = p.MitraProductInterestRateType
	p.MitraProductInterestRateFormula = p.MitraProductInterestRateFormula
	p.MitraProductInterestRateDiscMax = p.MitraProductInterestRateDiscMax
	p.MitraProductNumReviewer = p.MitraProductNumReviewer
	p.MitraProductNumApprover = p.MitraProductNumApprover
	p.MitraProductDateActiveStart = p.MitraProductDateActiveStart
	p.MitraProductDateActiveFinish = p.MitraProductDateActiveFinish
	p.MitraProductStatus = p.MitraProductStatus
	p.MitraProductCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraProductCreatedBy))
	p.MitraProductCreatedAt = time.Now()
	p.MitraProductUpdatedBy = html.EscapeString(strings.TrimSpace(p.MitraProductUpdatedBy))
	p.MitraProductUpdatedAt = p.MitraProductUpdatedAt
	p.MitraProductDeletedBy = html.EscapeString(strings.TrimSpace(p.MitraProductDeletedBy))
	p.MitraProductDeletedAt = p.MitraProductDeletedAt
}

func (p *MitraProduct) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.MitraID == 0 {
			return errors.New("Required Quadrant ID")
		}
		if p.MitraProductName == "" {
			return errors.New("Required Mitra Product Name")
		}
		if p.MitraProductLimitMin == 0 {
			return errors.New("Required Mitra Product Name")
		}
		if p.MitraProductLimitMax == 0 {
			return errors.New("Required Mitra Product Name")
		}
		if p.MitraProductTenorMin == 0 {
			return errors.New("Required Mitra Product Name")
		}
		if p.MitraProductTenorMax == 0 {
			return errors.New("Required Mitra Product Name")
		}
		if p.MitraProductInterestRate == 0 {
			return errors.New("Required Mitra Product Name")
		}
		if p.MitraProductInterestRatePeriod == "" {
			return errors.New("Required Mitra Product Name")
		}
		if p.MitraProductInterestRateType == "" {
			return errors.New("Required Mitra Product Name")
		}
		if p.MitraProductInterestRateDiscMax == 0 {
			return errors.New("Required Mitra Product Name")
		}
		if p.MitraProductDateActiveStart == "" {
			return errors.New("Required Mitra Product Name")
		}
		if p.MitraProductDateActiveFinish == "" {
			return errors.New("Required Mitra Product Name")
		}
		return nil
	}
}

func (p *MitraProduct) SaveMitraProduct(db *gorm.DB) (*MitraProduct, error) {
	var err error
	err = db.Debug().Model(&MitraProduct{}).Create(&p).Error
	if err != nil {
		return &MitraProduct{}, err
	}
	return p, nil
}

func (p *MitraProduct) FindAllMitraProduct(db *gorm.DB) (*[]MitraProductData, error) {
	var err error
	product := []MitraProductData{}
	err = db.Debug().Model(&MitraProductData{}).Limit(100).
		Select(`m_mitra_product.mitra_prod_id,
				m_mitra.mitra_id,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra.mitra_website,
				m_mitra.mitra_email,
				m_mitra.mitra_logo,
				m_customer_type.cust_type_id,
				m_customer_type.cust_type_name,
				m_mitra_product.mitra_prod_name,
				m_mitra_product.mitra_prod_limit_min,
				m_mitra_product.mitra_prod_limit_max,
				m_mitra_product.mitra_prod_tenor_min,
				m_mitra_product.mitra_prod_tenor_max,
				m_mitra_product.mitra_prod_interest_rate,
				m_mitra_product.mitra_prod_interest_rate_period,
				m_mitra_product.mitra_prod_interest_rate_type,
				m_mitra_product.mitra_prod_interest_rate_formula,
				m_mitra_product.mitra_prod_interest_rate_disc_max,
				m_mitra_product.mitra_prod_num_reviewer,
				m_mitra_product.mitra_prod_num_approver,
				m_mitra_product.mitra_prod_date_active_start,
				m_mitra_product.mitra_prod_date_active_finish,
				m_mitra_product.mitra_prod_status,
				m_mitra_product.mitra_prod_created_by,
				m_mitra_product.mitra_prod_updated_by,
				m_mitra_product.mitra_prod_deleted_by,
				m_mitra_product.mitra_prod_created_at,
				m_mitra_product.mitra_prod_updated_at,
				m_mitra_product.mitra_prod_deleted_at`).
		Joins("JOIN m_mitra on m_mitra_product.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_customer_type m_customer_type on m_mitra_product.cust_type_id=m_customer_type.cust_type_id").
		Find(&product).Error
	if err != nil {
		return &[]MitraProductData{}, err
	}
	return &product, nil
}

func (p *MitraProduct) FindAllMitraProductStatus(db *gorm.DB, status uint64, mitra uint64) (*[]MitraProductData, error) {
	var err error
	product := []MitraProductData{}
	err = db.Debug().Model(&MitraProductData{}).Limit(100).
		Select(`m_mitra_product.mitra_prod_id,
				m_mitra.mitra_id,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra.mitra_website,
				m_mitra.mitra_email,
				m_mitra.mitra_logo,
				m_customer_type.cust_type_id,
				m_customer_type.cust_type_name,
				m_mitra_product.mitra_prod_name,
				m_mitra_product.mitra_prod_limit_min,
				m_mitra_product.mitra_prod_limit_max,
				m_mitra_product.mitra_prod_tenor_min,
				m_mitra_product.mitra_prod_tenor_max,
				m_mitra_product.mitra_prod_interest_rate,
				m_mitra_product.mitra_prod_interest_rate_period,
				m_mitra_product.mitra_prod_interest_rate_type,
				m_mitra_product.mitra_prod_interest_rate_formula,
				m_mitra_product.mitra_prod_interest_rate_disc_max,
				m_mitra_product.mitra_prod_num_reviewer,
				m_mitra_product.mitra_prod_num_approver,
				m_mitra_product.mitra_prod_date_active_start,
				m_mitra_product.mitra_prod_date_active_finish,
				m_mitra_product.mitra_prod_status,
				m_mitra_product.mitra_prod_created_by,
				m_mitra_product.mitra_prod_updated_by,
				m_mitra_product.mitra_prod_deleted_by,
				m_mitra_product.mitra_prod_created_at,
				m_mitra_product.mitra_prod_updated_at,
				m_mitra_product.mitra_prod_deleted_at`).
		Joins("JOIN m_mitra on m_mitra_product.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_customer_type m_customer_type on m_mitra_product.cust_type_id=m_customer_type.cust_type_id").
		Where("m_mitra_product.mitra_prod_status = ? AND m_mitra_product.mitra_id = ?", status, mitra).
		Find(&product).Error
	if err != nil {
		return &[]MitraProductData{}, err
	}
	return &product, nil
}

func (p *MitraProduct) FindMitraProductByID(db *gorm.DB, pid uint64) (*MitraProduct, error) {
	var err error
	err = db.Debug().Model(&MitraProduct{}).Where("mitra_prod_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraProduct{}, err
	}
	return p, nil
}

func (p *MitraProduct) FindMitraProductDataByID(db *gorm.DB, pid uint64) (*MitraProductData, error) {
	var err error
	product := MitraProductData{}
	err = db.Debug().Model(&MitraProductData{}).Limit(100).
		Select(`m_mitra_product.mitra_prod_id,
				m_mitra.mitra_id,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra.mitra_website,
				m_mitra.mitra_email,
				m_mitra.mitra_logo,
				m_customer_type.cust_type_id,
				m_customer_type.cust_type_name,
				m_mitra_product.mitra_prod_name,
				m_mitra_product.mitra_prod_limit_min,
				m_mitra_product.mitra_prod_limit_max,
				m_mitra_product.mitra_prod_tenor_min,
				m_mitra_product.mitra_prod_tenor_max,
				m_mitra_product.mitra_prod_interest_rate,
				m_mitra_product.mitra_prod_interest_rate_period,
				m_mitra_product.mitra_prod_interest_rate_type,
				m_mitra_product.mitra_prod_interest_rate_formula,
				m_mitra_product.mitra_prod_interest_rate_disc_max,
				m_mitra_product.mitra_prod_num_reviewer,
				m_mitra_product.mitra_prod_num_approver,
				m_mitra_product.mitra_prod_date_active_start,
				m_mitra_product.mitra_prod_date_active_finish,
				m_mitra_product.mitra_prod_status,
				m_mitra_product.mitra_prod_created_by,
				m_mitra_product.mitra_prod_updated_by,
				m_mitra_product.mitra_prod_deleted_by,
				m_mitra_product.mitra_prod_created_at,
				m_mitra_product.mitra_prod_updated_at,
				m_mitra_product.mitra_prod_deleted_at`).
		Joins("JOIN m_mitra on m_mitra_product.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_customer_type m_customer_type on m_mitra_product.cust_type_id=m_customer_type.cust_type_id").
		Joins("JOIN m_mitra on m_mitra_product.mitra_id=m_mitra.mitra_id").
		Where("m_mitra_product.mitra_prod_id = ?", pid).
		Find(&product).Error
	if err != nil {
		return &MitraProductData{}, err
	}
	return &product, nil
}

func (p *MitraProduct) FindMitraProductStatusByID(db *gorm.DB, pid uint64, status uint64, mitra uint64) (*MitraProductData, error) {
	var err error
	product := MitraProductData{}
	err = db.Debug().Model(&MitraProductData{}).Limit(100).
		Select(`m_mitra_product.mitra_prod_id,
				m_mitra.mitra_id,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra.mitra_website,
				m_mitra.mitra_email,
				m_mitra.mitra_logo,
				m_customer_type.cust_type_id,
				m_customer_type.cust_type_name,
				m_mitra_product.mitra_prod_name,
				m_mitra_product.mitra_prod_limit_min,
				m_mitra_product.mitra_prod_limit_max,
				m_mitra_product.mitra_prod_tenor_min,
				m_mitra_product.mitra_prod_tenor_max,
				m_mitra_product.mitra_prod_interest_rate,
				m_mitra_product.mitra_prod_interest_rate_period,
				m_mitra_product.mitra_prod_interest_rate_type,
				m_mitra_product.mitra_prod_interest_rate_formula,
				m_mitra_product.mitra_prod_interest_rate_disc_max,
				m_mitra_product.mitra_prod_num_reviewer,
				m_mitra_product.mitra_prod_num_approver,
				m_mitra_product.mitra_prod_date_active_start,
				m_mitra_product.mitra_prod_date_active_finish,
				m_mitra_product.mitra_prod_status,
				m_mitra_product.mitra_prod_created_by,
				m_mitra_product.mitra_prod_updated_by,
				m_mitra_product.mitra_prod_deleted_by,
				m_mitra_product.mitra_prod_created_at,
				m_mitra_product.mitra_prod_updated_at,
				m_mitra_product.mitra_prod_deleted_at`).
		Joins("JOIN m_mitra on m_mitra_product.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_customer_type m_customer_type on m_mitra_product.cust_type_id=m_customer_type.cust_type_id").
		Where("m_mitra_product.mitra_prod_id = ? AND m_mitra_product.mitra_prod_status = ? AND m_mitra_product.mitra_id = ?", pid, status, mitra).
		Find(&product).Error
	if err != nil {
		return &MitraProductData{}, err
	}
	return &product, nil
}

func (p *MitraProduct) UpdateMitraProduct(db *gorm.DB, pid uint64) (*MitraProduct, error) {
	var err error
	db = db.Debug().Model(&MitraProduct{}).Where("mitra_prod_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_id":                          p.MitraID,
			"cust_type_id":                      p.CustomerTypeID,
			"mitra_prod_name":                   p.MitraProductName,
			"mitra_prod_limit_min":              p.MitraProductLimitMin,
			"mitra_prod_limit_max":              p.MitraProductLimitMax,
			"mitra_prod_tenor_min":              p.MitraProductTenorMin,
			"mitra_prod_tenor_max":              p.MitraProductTenorMax,
			"mitra_prod_interest_rate":          p.MitraProductInterestRate,
			"mitra_prod_interest_rate_period":   p.MitraProductInterestRatePeriod,
			"mitra_prod_interest_rate_type":     p.MitraProductInterestRateType,
			"mitra_prod_interest_rate_disc_max": p.MitraProductInterestRateDiscMax,
			"mitra_prod_num_reviewer":           p.MitraProductNumReviewer,
			"mitra_prod_num_approver":           p.MitraProductNumApprover,
			"mitra_prod_date_active_start":      p.MitraProductDateActiveStart,
			"mitra_prod_date_active_finish":     p.MitraProductDateActiveFinish,
			"mitra_prod_status":                 p.MitraProductStatus,
			"mitra_prod_updated_by":             p.MitraProductUpdatedBy,
			"mitra_prod_updated_at":             time.Now(),
		},
	)
	err = db.Debug().Model(&MitraProduct{}).Where("mitra_prod_id = ?", pid).Error
	if err != nil {
		return &MitraProduct{}, err
	}
	return p, nil
}

func (p *MitraProduct) UpdateMitraProductStatus(db *gorm.DB, pid uint64) (*MitraProduct, error) {
	var err error
	db = db.Debug().Model(&MitraProduct{}).Where("mitra_prod_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_prod_status":     p.MitraProductStatus,
			"mitra_prod_updated_by": p.MitraProductUpdatedBy,
			"mitra_prod_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraProduct{}).Where("mitra_prod_id = ?", pid).Error
	if err != nil {
		return &MitraProduct{}, err
	}
	return p, nil
}

func (p *MitraProduct) UpdateMitraProductStatusOnly(db *gorm.DB, pid uint64) (*MitraProduct, error) {
	var err error
	db = db.Debug().Model(&MitraProduct{}).Where("mitra_prod_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_prod_status": p.MitraProductStatus,
		},
	)
	err = db.Debug().Model(&MitraProduct{}).Where("mitra_prod_id = ?", pid).Error
	if err != nil {
		return &MitraProduct{}, err
	}
	return p, nil
}

func (p *MitraProduct) UpdateMitraProductStatusDelete(db *gorm.DB, pid uint64) (*MitraProduct, error) {
	var err error
	db = db.Debug().Model(&MitraProduct{}).Where("mitra_prod_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_prod_status":     3,
			"mitra_prod_deleted_by": p.MitraProductDeletedBy,
			"mitra_prod_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraProduct{}).Where("mitra_prod_id = ?", pid).Error
	if err != nil {
		return &MitraProduct{}, err
	}
	return p, nil
}

func (p *MitraProduct) DeleteMitraProduct(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraProduct{}).Where("mitra_prod_id = ?", pid).Take(&MitraProduct{}).Delete(&MitraProduct{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Mitra Product not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//PRODUCT WORKFLOW
//====================================================================================================================================

type MitraProductWorkflow struct {
	MitraProductWorkflow []MitraProductWorkflowPend `gorm:"column:mitra_prod_workflow" json:"mitra_prod_workflow"`
}

type MitraProductWorkflowPend struct {
	MitraProductID                uint64 `gorm:"column:mitra_prod_id" json:"mitra_prod_id"`
	MitraProductNumReviewer       uint64 `gorm:"column:mitra_prod_num_reviewer" json:"mitra_prod_num_reviewer"`
	MitraProductNumApprover       uint64 `gorm:"column:mitra_prod_num_approver" json:"mitra_prod_num_approver"`
	MitraProductWorkflowUpdatedBy string `gorm:"column:mitra_prod_workflow_updated_by" json:"mitra_prod_workflow_updated_by"`
}

type MitraProductWorkflowData struct {
	MitraProductID          uint64     `gorm:"column:mitra_prod_id;primary_key;" json:"mitra_prod_id"`
	MitraID                 uint64     `gorm:"column:mitra_id" json:"mitra_id"`
	MitraCode               string     `gorm:"column:mitra_code" json:"mitra_code"`
	MitraName               string     `gorm:"column:mitra_name" json:"mitra_name"`
	MitraProductName        string     `gorm:"column:mitra_prod_name" json:"mitra_prod_name"`
	MitraProductNumReviewer uint64     `gorm:"column:mitra_prod_num_reviewer" json:"mitra_prod_num_reviewer"`
	MitraProductNumApprover uint64     `gorm:"column:mitra_prod_num_approver" json:"mitra_prod_num_approver"`
	MitraProductStatus      uint64     `gorm:"column:mitra_prod_status" json:"mitra_prod_status"`
	MitraProductCreatedAt   time.Time  `gorm:"column:mitra_prod_created_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_created_at"`
	MitraProductUpdatedAt   *time.Time `gorm:"column:mitra_prod_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_updated_at"`
	MitraProductDeletedAt   *time.Time `gorm:"column:mitra_prod_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_deleted_at"`
}

type ResponseMitraProductsWorkflow struct {
	Status  int                         `json:"status"`
	Message string                      `json:"message"`
	Data    *[]MitraProductWorkflowData `json:"data"`
}

type ResponseMitraProductWorkflow struct {
	Status  int                       `json:"status"`
	Message string                    `json:"message"`
	Data    *MitraProductWorkflowData `json:"data"`
}

func (MitraProductWorkflowData) TableName() string {
	return "m_mitra_product"
}

func (p *MitraProduct) FindAllMitraProductWorkflow(db *gorm.DB, status uint64, mitra uint64) (*[]MitraProductWorkflowData, error) {
	var err error
	product := []MitraProductWorkflowData{}
	err = db.Debug().Model(&MitraProductWorkflowData{}).Limit(100).
		Select(`m_mitra_product.mitra_prod_id,
				m_mitra.mitra_id,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra_product.mitra_prod_name,
				m_mitra_product.mitra_prod_num_reviewer,
				m_mitra_product.mitra_prod_num_approver,
				m_mitra_product.mitra_prod_status,
				m_mitra_product.mitra_prod_created_at,
				m_mitra_product.mitra_prod_updated_at,
				m_mitra_product.mitra_prod_deleted_at`).
		Joins("JOIN m_mitra on m_mitra_product.mitra_id=m_mitra.mitra_id").
		Where("m_mitra_product.mitra_prod_status = ? AND m_mitra_product.mitra_id = ?", status, mitra).
		Find(&product).Error
	if err != nil {
		return &[]MitraProductWorkflowData{}, err
	}
	return &product, nil
}

func (p *MitraProduct) FindMitraProductWorkflowByID(db *gorm.DB, pid uint64, status uint64, mitra uint64) (*MitraProductWorkflowData, error) {
	var err error
	product := MitraProductWorkflowData{}
	err = db.Debug().Model(&MitraProductWorkflowData{}).Limit(100).
		Select(`m_mitra_product.mitra_prod_id,
				m_mitra.mitra_id,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra_product.mitra_prod_name,
				m_mitra_product.mitra_prod_num_reviewer,
				m_mitra_product.mitra_prod_num_approver,
				m_mitra_product.mitra_prod_status,
				m_mitra_product.mitra_prod_created_at,
				m_mitra_product.mitra_prod_updated_at,
				m_mitra_product.mitra_prod_deleted_at
			`).
		Joins("JOIN m_mitra on m_mitra_product.mitra_id=m_mitra.mitra_id").
		Where("m_mitra_product.mitra_prod_id = ? AND m_mitra_product.mitra_prod_status = ? AND m_mitra_product.mitra_id = ?", pid, status, mitra).
		Find(&product).Error
	if err != nil {
		return &MitraProductWorkflowData{}, err
	}
	return &product, nil
}

func (p *MitraProduct) UpdateMitraProductWorkflow(db *gorm.DB, pid uint64) (*MitraProduct, error) {
	var err error
	db = db.Debug().Model(&MitraProduct{}).Where("mitra_prod_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_prod_num_reviewer": p.MitraProductNumReviewer,
			"mitra_prod_num_approver": p.MitraProductNumApprover,
			"mitra_prod_updated_by":   p.MitraProductUpdatedBy,
			"mitra_prod_updated_at":   time.Now(),
		},
	)
	err = db.Debug().Model(&MitraProduct{}).Where("mitra_prod_id = ?", pid).Error
	if err != nil {
		return &MitraProduct{}, err
	}
	return p, nil
}

//ADDITIONAL
//====================================================================================================================================

type MitraProductOfferingListData struct {
	MitraProductID     uint64 `gorm:"column:mitra_prod_id;primary_key;" json:"mitra_prod_id"`
	MitraProductName   string `gorm:"column:mitra_prod_name;size:255" json:"mitra_prod_name"`
	MitraProductStatus uint64 `gorm:"column:mitra_prod_status" json:"mitra_prod_status"`
}

type ResponseLoanOfferingProduct struct {
	Status  int                             `json:"status"`
	Message string                          `json:"message"`
	Data    *[]MitraProductOfferingListData `json:"data"`
}

type ResponseLoanOfferingProductByMitra struct {
	Status  int                 `json:"status"`
	Message string              `json:"message"`
	Data    *[]MitraProductData `json:"data"`
}

type ResponseLoanOfferingProductDetail struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    *MitraProductData `json:"data"`
}

func (MitraProductOfferingListData) TableName() string {
	return "m_mitra_product"
}

func (p *MitraProduct) FindAllMitraProductStatusByMitraByCtgry(db *gorm.DB, mitra uint64, subctgry uint64) (*[]MitraProductOfferingListData, error) {
	var err error
	product := []MitraProductOfferingListData{}
	err = db.Debug().Model(&MitraProductOfferingListData{}).Limit(100).
		Select(`m_mitra_product.mitra_prod_id,
				m_mitra.mitra_id,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra.mitra_website,
				m_mitra.mitra_email,
				m_mitra.mitra_logo,
				m_customer_type.cust_type_id,
				m_customer_type.cust_type_name,
				m_mitra_product.mitra_prod_name,
				m_mitra_product.mitra_prod_limit_min,
				m_mitra_product.mitra_prod_limit_max,
				m_mitra_product.mitra_prod_tenor_min,
				m_mitra_product.mitra_prod_tenor_max,
				m_mitra_product.mitra_prod_interest_rate,
				m_mitra_product.mitra_prod_interest_rate_period,
				m_mitra_product.mitra_prod_interest_rate_type,
				m_mitra_product.mitra_prod_interest_rate_disc_max,
				m_mitra_product.mitra_prod_num_reviewer,
				m_mitra_product.mitra_prod_num_approver,
				m_mitra_product.mitra_prod_status,
				m_mitra_product.mitra_prod_created_at,
				m_mitra_product.mitra_prod_updated_at,
				m_mitra_product.mitra_prod_deleted_at
				`).
		Joins("JOIN m_mitra_product_sub_category on m_mitra_product.mitra_prod_id=m_mitra_product_sub_category.mitra_prod_id").
		Joins("JOIN m_mitra on m_mitra_product.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_customer_type m_customer_type on m_mitra_product.cust_type_id=m_customer_type.cust_type_id").
		Where("m_mitra_product.mitra_id = ? AND prod_sub_ctgry_id = ? AND m_mitra_product.mitra_prod_status = ?", mitra, subctgry, 1).
		Find(&product).Error
	if err != nil {
		return &[]MitraProductOfferingListData{}, err
	}
	return &product, nil
}

func (p *MitraProduct) FindAllMitraProductStatusByMitra(db *gorm.DB, mitra uint64) (*[]MitraProductData, error) {
	var err error
	product := []MitraProductData{}
	err = db.Debug().Model(&MitraProductData{}).Limit(100).
		Select(`m_mitra_product.mitra_prod_id,
				m_mitra.mitra_id,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra.mitra_website,
				m_mitra.mitra_email,
				m_mitra.mitra_logo,
				m_customer_type.cust_type_id,
				m_customer_type.cust_type_name,
				m_mitra_product.mitra_prod_name,
				m_mitra_product.mitra_prod_limit_min,
				m_mitra_product.mitra_prod_limit_max,
				m_mitra_product.mitra_prod_tenor_min,
				m_mitra_product.mitra_prod_tenor_max,
				m_mitra_product.mitra_prod_interest_rate,
				m_mitra_product.mitra_prod_interest_rate_period,
				m_mitra_product.mitra_prod_interest_rate_type,
				m_mitra_product.mitra_prod_interest_rate_disc_max,
				m_mitra_product.mitra_prod_num_reviewer,
				m_mitra_product.mitra_prod_num_approver,
				m_mitra_product.mitra_prod_status,
				m_mitra_product.mitra_prod_created_at,
				m_mitra_product.mitra_prod_updated_at,
				m_mitra_product.mitra_prod_deleted_at
				`).
		Joins("JOIN m_mitra on m_mitra_product.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_customer_type m_customer_type on m_mitra_product.cust_type_id=m_customer_type.cust_type_id").
		Where("m_mitra_product.mitra_id = ? AND m_mitra_product.mitra_prod_status = ?", mitra, 1).
		Find(&product).Error
	if err != nil {
		return &[]MitraProductData{}, err
	}
	return &product, nil
}

func (p *MitraProduct) FindMitraProductByMitraUserID(db *gorm.DB, uid uint64) (*[]MitraProductData, error) {
	var err error
	product := []MitraProductData{}
	err = db.Debug().Model(&MitraProductData{}).Limit(100).
		Select(`m_mitra_user.mitra_user_id,
				m_mitra_product.mitra_prod_id,
				m_mitra.mitra_id,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra.mitra_website,
				m_mitra.mitra_email,
				m_mitra.mitra_logo,
				m_customer_type.cust_type_id,
				m_customer_type.cust_type_name,
				m_mitra_product.mitra_prod_name,
				m_mitra_product.mitra_prod_limit_min,
				m_mitra_product.mitra_prod_limit_max,
				m_mitra_product.mitra_prod_tenor_min,
				m_mitra_product.mitra_prod_tenor_max,
				m_mitra_product.mitra_prod_interest_rate,
				m_mitra_product.mitra_prod_interest_rate_period,
				m_mitra_product.mitra_prod_interest_rate_type,
				m_mitra_product.mitra_prod_interest_rate_disc_max,
				m_mitra_product.mitra_prod_num_reviewer,
				m_mitra_product.mitra_prod_num_approver,
				m_mitra_product.mitra_prod_status,
				m_mitra_product.mitra_prod_created_by,
				m_mitra_product.mitra_prod_updated_by,
				m_mitra_product.mitra_prod_deleted_by,
				m_mitra_product.mitra_prod_created_at,
				m_mitra_product.mitra_prod_updated_at,
				m_mitra_product.mitra_prod_deleted_at
				`).
		Joins("JOIN m_mitra on m_mitra_product.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_employee on m_mitra_product.mitra_id=m_mitra_employee.mitra_id").
		Joins("JOIN m_mitra_user on m_mitra_employee.mitra_emp_id=m_mitra_user.mitra_emp_id").
		Joins("JOIN m_customer_type m_customer_type on m_mitra_product.cust_type_id=m_customer_type.cust_type_id").
		Where("mitra_user_id = ?", uid).
		Find(&product).Error
	if err != nil {
		return &[]MitraProductData{}, err
	}
	return &product, nil
}

func (p *MitraProduct) FindMitraProductByLimit(db *gorm.DB, pid uint64, limitmin uint64, limitmax uint64) (*MitraProductData, error) {
	var err error
	product := MitraProductData{}
	err = db.Debug().Model(&MitraProductData{}).Limit(100).
		Select(`m_mitra_product.mitra_prod_id,
				m_mitra.mitra_id,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra.mitra_website,
				m_mitra.mitra_email,
				m_mitra.mitra_logo,
				m_customer_type.cust_type_id,
				m_customer_type.cust_type_name,
				m_mitra_product.mitra_prod_name,
				m_mitra_product.mitra_prod_limit_min,
				m_mitra_product.mitra_prod_limit_max,
				m_mitra_product.mitra_prod_tenor_min,
				m_mitra_product.mitra_prod_tenor_max,
				m_mitra_product.mitra_prod_interest_rate,
				m_mitra_product.mitra_prod_interest_rate_period,
				m_mitra_product.mitra_prod_interest_rate_type,
				m_mitra_product.mitra_prod_interest_rate_disc_max,
				m_mitra_product.mitra_prod_num_reviewer,
				m_mitra_product.mitra_prod_num_approver,
				m_mitra_product.mitra_prod_status,
				m_mitra_product.mitra_prod_created_by,
				m_mitra_product.mitra_prod_updated_by,
				m_mitra_product.mitra_prod_deleted_by,
				m_mitra_product.mitra_prod_created_at,
				m_mitra_product.mitra_prod_updated_at,
				m_mitra_product.mitra_prod_deleted_at
				`).
		Joins("JOIN m_mitra on m_mitra_product.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_customer_type m_customer_type on m_mitra_product.cust_type_id=m_customer_type.cust_type_id").
		Where("mitra_prod_id = ?", pid).
		Find(&product).Error
	if err != nil {
		return &MitraProductData{}, err
	}
	return &product, nil
}

func (p *MitraProduct) FindMitraProductStatusByOnlyID(db *gorm.DB, pid uint64) (*MitraProductData, error) {
	var err error
	product := MitraProductData{}
	err = db.Debug().Model(&MitraProductData{}).Limit(100).
		Select(`m_mitra_product.mitra_prod_id,
			m_mitra.mitra_id,
			m_mitra.mitra_code,
			m_mitra.mitra_name,
			m_mitra.mitra_website,
			m_mitra.mitra_email,
			m_mitra.mitra_logo,
			m_customer_type.cust_type_id,
			m_customer_type.cust_type_name,
			m_mitra_product.mitra_prod_name,
			m_mitra_product.mitra_prod_limit_min,
			m_mitra_product.mitra_prod_limit_max,
			m_mitra_product.mitra_prod_tenor_min,
			m_mitra_product.mitra_prod_tenor_max,
			m_mitra_product.mitra_prod_interest_rate,
			m_mitra_product.mitra_prod_interest_rate_period,
			m_mitra_product.mitra_prod_interest_rate_type,
			m_mitra_product.mitra_prod_interest_rate_disc_max,
			m_mitra_product.mitra_prod_num_reviewer,
			m_mitra_product.mitra_prod_num_approver,
			m_mitra_product.mitra_prod_status,
			m_mitra_product.mitra_prod_created_at,
			m_mitra_product.mitra_prod_updated_at,
			m_mitra_product.mitra_prod_deleted_at
			`).
		Joins("JOIN m_mitra on m_mitra_product.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_customer_type m_customer_type on m_mitra_product.cust_type_id=m_customer_type.cust_type_id").
		Where("mitra_prod_id = ? AND mitra_prod_status = ?", pid, 1).
		Find(&product).Error
	if err != nil {
		return &MitraProductData{}, err
	}
	return &product, nil
}

func (p *MitraProduct) FindMitraProductByMitraID(db *gorm.DB, mitra uint64, pid uint64) (*MitraProductData, error) {
	var err error
	product := MitraProductData{}
	err = db.Debug().Model(&MitraProductData{}).Limit(100).
		Select(`m_mitra_product.mitra_prod_id,
				m_mitra.mitra_id,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra.mitra_website,
				m_mitra.mitra_email,
				m_mitra.mitra_logo,
				m_customer_type.cust_type_id,
				m_customer_type.cust_type_name,
				m_mitra_product.mitra_prod_name,
				m_mitra_product.mitra_prod_limit_min,
				m_mitra_product.mitra_prod_limit_max,
				m_mitra_product.mitra_prod_tenor_min,
				m_mitra_product.mitra_prod_tenor_max,
				m_mitra_product.mitra_prod_interest_rate,
				m_mitra_product.mitra_prod_interest_rate_period,
				m_mitra_product.mitra_prod_interest_rate_type,
				m_mitra_product.mitra_prod_interest_rate_formula,
				m_mitra_product.mitra_prod_interest_rate_disc_max,
				m_mitra_product.mitra_prod_num_reviewer,
				m_mitra_product.mitra_prod_num_approver,
				m_mitra_product.mitra_prod_date_active_start,
				m_mitra_product.mitra_prod_date_active_finish,
				m_mitra_product.mitra_prod_status,
				m_mitra_product.mitra_prod_created_by,
				m_mitra_product.mitra_prod_updated_by,
				m_mitra_product.mitra_prod_deleted_by,
				m_mitra_product.mitra_prod_created_at,
				m_mitra_product.mitra_prod_updated_at,
				m_mitra_product.mitra_prod_deleted_at`).
		Joins("JOIN m_mitra on m_mitra_product.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_customer_type m_customer_type on m_mitra_product.cust_type_id=m_customer_type.cust_type_id").
		Where("m_mitra_product.mitra_prod_id = ? AND m_mitra_product.mitra_id = ?", pid, mitra).
		Find(&product).Error
	if err != nil {
		return &MitraProductData{}, err
	}
	return &product, nil
}
