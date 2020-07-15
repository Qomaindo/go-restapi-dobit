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

type MitraProductTemp struct {
	MitraProductTempID                  uint64    `gorm:"column:mitra_prod_temp_id;primary_key;" json:"mitra_prod_temp_id"`
	MitraProductID                      uint64    `gorm:"column:mitra_prod_id" json:"mitra_prod_id"`
	MitraTempID                         uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	CustomerTypeTempID                  uint64    `gorm:"column:cust_type_temp_id" json:"cust_type_temp_id"`
	MitraProductTempName                string    `gorm:"column:mitra_prod_temp_name;size:125" json:"mitra_prod_temp_name"`
	MitraProductTempLimitMin            uint64    `gorm:"column:mitra_prod_temp_limit_min" json:"mitra_prod_temp_limit_min"`
	MitraProductTempLimitMax            uint64    `gorm:"column:mitra_prod_temp_limit_max" json:"mitra_prod_temp_limit_max"`
	MitraProductTempTenorMin            uint64    `gorm:"column:mitra_prod_temp_tenor_min" json:"mitra_prod_temp_tenor_min"`
	MitraProductTempTenorMax            uint64    `gorm:"column:mitra_prod_temp_tenor_max" json:"mitra_prod_temp_tenor_max"`
	MitraProductTempInterestRate        float64   `gorm:"column:mitra_prod_temp_interest_rate" json:"mitra_prod_temp_interest_rate"`
	MitraProductTempInterestRatePeriod  string    `gorm:"column:mitra_prod_temp_interest_rate_period;size:125" json:"mitra_prod_temp_interest_rate_period"`
	MitraProductTempInterestRateType    string    `gorm:"column:mitra_prod_temp_interest_rate_type;size:125" json:"mitra_prod_temp_interest_rate_type"`
	MitraProductTempInterestRateFormula string    `gorm:"column:mitra_prod_temp_interest_rate_formula;size:125" json:"mitra_prod_temp_interest_rate_formula"`
	MitraProductTempInterestRateDiscMax float64   `gorm:"column:mitra_prod_temp_interest_rate_disc_max" json:"mitra_prod_temp_interest_rate_disc_max"`
	MitraProductTempNumReviewer         uint64    `gorm:"column:mitra_prod_temp_num_reviewer" json:"mitra_prod_temp_num_reviewer"`
	MitraProductTempNumApprover         uint64    `gorm:"column:mitra_prod_temp_num_approver" json:"mitra_prod_temp_num_approver"`
	MitraProductTempDateActiveStart     string    `gorm:"column:mitra_prod_temp_date_active_start" json:"mitra_prod_temp_date_active_start"`
	MitraProductTempDateActiveFinish    string    `gorm:"column:mitra_prod_temp_date_active_finish" json:"mitra_prod_temp_date_active_finish"`
	MitraProductTempReason              string    `gorm:"column:mitra_prod_temp_reason" json:"mitra_prod_temp_reason"`
	MitraProductTempStatus              uint64    `gorm:"column:mitra_prod_temp_status;size:2" json:"mitra_prod_temp_status"`
	MitraProductTempActionBefore        uint64    `gorm:"column:mitra_prod_temp_action_before;size:2" json:"mitra_prod_temp_action_before"`
	MitraProductTempCreatedBy           string    `gorm:"column:mitra_prod_temp_created_by;size:125" json:"mitra_prod_temp_created_by"`
	MitraProductTempCreatedAt           time.Time `gorm:"column:mitra_prod_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_temp_created_at"`

	MitraProductSubCategoryTemp  []MitraProductSubCategoryTempPendData  `json:"mitra_prod_sub_ctgry_temp"`
	MitraProductCostTemp         []MitraProductCostTempPendData         `json:"mitra_prod_cost_temp"`
	MitraProductDocumentRuleTemp []MitraProductDocumentRuleTempPendData `json:"mitra_prod_doc_rule_temp"`
}

type MitraProductTempPend struct {
	MitraProductTempID                  uint64    `gorm:"column:mitra_prod_temp_id;primary_key;" json:"mitra_prod_temp_id"`
	MitraProductID                      *uint64   `gorm:"column:mitra_prod_id" json:"mitra_prod_id"`
	MitraTempID                         uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	CustomerTypeTempID                  uint64    `gorm:"column:cust_type_temp_id" json:"cust_type_temp_id"`
	MitraProductTempName                string    `gorm:"column:mitra_prod_temp_name;size:125" json:"mitra_prod_temp_name"`
	MitraProductTempLimitMin            uint64    `gorm:"column:mitra_prod_temp_limit_min" json:"mitra_prod_temp_limit_min"`
	MitraProductTempLimitMax            uint64    `gorm:"column:mitra_prod_temp_limit_max" json:"mitra_prod_temp_limit_max"`
	MitraProductTempTenorMin            uint64    `gorm:"column:mitra_prod_temp_tenor_min" json:"mitra_prod_temp_tenor_min"`
	MitraProductTempTenorMax            uint64    `gorm:"column:mitra_prod_temp_tenor_max" json:"mitra_prod_temp_tenor_max"`
	MitraProductTempInterestRate        float64   `gorm:"column:mitra_prod_temp_interest_rate" json:"mitra_prod_temp_interest_rate"`
	MitraProductTempInterestRatePeriod  string    `gorm:"column:mitra_prod_temp_interest_rate_period;size:125" json:"mitra_prod_temp_interest_rate_period"`
	MitraProductTempInterestRateType    string    `gorm:"column:mitra_prod_temp_interest_rate_type;size:125" json:"mitra_prod_temp_interest_rate_type"`
	MitraProductTempInterestRateFormula string    `gorm:"column:mitra_prod_temp_interest_rate_formula;size:125" json:"mitra_prod_temp_interest_rate_formula"`
	MitraProductTempInterestRateDiscMax float64   `gorm:"column:mitra_prod_temp_interest_rate_disc_max" json:"mitra_prod_temp_interest_rate_disc_max"`
	MitraProductTempNumReviewer         uint64    `gorm:"column:mitra_prod_temp_num_reviewer" json:"mitra_prod_temp_num_reviewer"`
	MitraProductTempNumApprover         uint64    `gorm:"column:mitra_prod_temp_num_approver" json:"mitra_prod_temp_num_approver"`
	MitraProductTempDateActiveStart     string    `gorm:"column:mitra_prod_temp_date_active_start" json:"mitra_prod_temp_date_active_start"`
	MitraProductTempDateActiveFinish    string    `gorm:"column:mitra_prod_temp_date_active_finish" json:"mitra_prod_temp_date_active_finish"`
	MitraProductTempReason              string    `gorm:"column:mitra_prod_temp_reason" json:"mitra_prod_temp_reason"`
	MitraProductTempStatus              uint64    `gorm:"column:mitra_prod_temp_status;size:2" json:"mitra_prod_temp_status"`
	MitraProductTempActionBefore        uint64    `gorm:"column:mitra_prod_temp_action_before;size:2" json:"mitra_prod_temp_action_before"`
	MitraProductTempCreatedBy           string    `gorm:"column:mitra_prod_temp_created_by;size:125" json:"mitra_prod_temp_created_by"`
	MitraProductTempCreatedAt           time.Time `gorm:"column:mitra_prod_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_temp_created_at"`

	MitraProductSubCategoryTemp  []MitraProductSubCategoryTempPend  `json:"mitra_prod_sub_ctgry_temp"`
	MitraProductCostTemp         []MitraProductCostTempPend         `json:"mitra_prod_cost_temp"`
	MitraProductDocumentRuleTemp []MitraProductDocumentRuleTempPend `json:"mitra_prod_doc_rule_temp"`
}

type MitraProductTempPendData struct {
	MitraProductTempID                  uint64    `gorm:"column:mitra_prod_temp_id;primary_key;" json:"mitra_prod_temp_id"`
	MitraProductID                      *uint64   `gorm:"column:mitra_prod_id" json:"mitra_prod_id"`
	MitraTempID                         uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	CustomerTypeTempID                  uint64    `gorm:"column:cust_type_temp_id" json:"cust_type_temp_id"`
	MitraProductTempName                string    `gorm:"column:mitra_prod_temp_name;size:125" json:"mitra_prod_temp_name"`
	MitraProductTempLimitMin            uint64    `gorm:"column:mitra_prod_temp_limit_min" json:"mitra_prod_temp_limit_min"`
	MitraProductTempLimitMax            uint64    `gorm:"column:mitra_prod_temp_limit_max" json:"mitra_prod_temp_limit_max"`
	MitraProductTempTenorMin            uint64    `gorm:"column:mitra_prod_temp_tenor_min" json:"mitra_prod_temp_tenor_min"`
	MitraProductTempTenorMax            uint64    `gorm:"column:mitra_prod_temp_tenor_max" json:"mitra_prod_temp_tenor_max"`
	MitraProductTempInterestRate        float64   `gorm:"column:mitra_prod_temp_interest_rate" json:"mitra_prod_temp_interest_rate"`
	MitraProductTempInterestRatePeriod  string    `gorm:"column:mitra_prod_temp_interest_rate_period;size:125" json:"mitra_prod_temp_interest_rate_period"`
	MitraProductTempInterestRateType    string    `gorm:"column:mitra_prod_temp_interest_rate_type;size:125" json:"mitra_prod_temp_interest_rate_type"`
	MitraProductTempInterestRateFormula string    `gorm:"column:mitra_prod_temp_interest_rate_formula;size:125" json:"mitra_prod_temp_interest_rate_formula"`
	MitraProductTempInterestRateDiscMax float64   `gorm:"column:mitra_prod_temp_interest_rate_disc_max" json:"mitra_prod_temp_interest_rate_disc_max"`
	MitraProductTempNumReviewer         uint64    `gorm:"column:mitra_prod_temp_num_reviewer" json:"mitra_prod_temp_num_reviewer"`
	MitraProductTempNumApprover         uint64    `gorm:"column:mitra_prod_temp_num_approver" json:"mitra_prod_temp_num_approver"`
	MitraProductTempDateActiveStart     string    `gorm:"column:mitra_prod_temp_date_active_start" json:"mitra_prod_temp_date_active_start"`
	MitraProductTempDateActiveFinish    string    `gorm:"column:mitra_prod_temp_date_active_finish" json:"mitra_prod_temp_date_active_finish"`
	MitraProductTempReason              string    `gorm:"column:mitra_prod_temp_reason" json:"mitra_prod_temp_reason"`
	MitraProductTempStatus              uint64    `gorm:"column:mitra_prod_temp_status;size:2" json:"mitra_prod_temp_status"`
	MitraProductTempActionBefore        uint64    `gorm:"column:mitra_prod_temp_action_before;size:2" json:"mitra_prod_temp_action_before"`
	MitraProductTempCreatedBy           string    `gorm:"column:mitra_prod_temp_created_by;size:125" json:"mitra_prod_temp_created_by"`
	MitraProductTempCreatedAt           time.Time `gorm:"column:mitra_prod_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_temp_created_at"`

	MitraProductSubCategoryTemp  []MitraProductSubCategoryTempPendData  `json:"mitra_prod_sub_ctgry_temp"`
	MitraProductCostTemp         []MitraProductCostTempPendData         `json:"mitra_prod_cost_temp"`
	MitraProductDocumentRuleTemp []MitraProductDocumentRuleTempPendData `json:"mitra_prod_doc_rule_temp"`
}

type MitraProductTempData struct {
	MitraProductTempID                  uint64    `gorm:"column:mitra_prod_temp_id;primary_key;" json:"mitra_prod_temp_id"`
	MitraTempID                         uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	CustomerTypeTempID                  uint64    `gorm:"column:cust_type_temp_id" json:"cust_type_temp_id"`
	CustomerTypeTempName                string    `gorm:"column:cust_type_temp_name" json:"cust_type_temp_name"`
	MitraTempCode                       string    `gorm:"column:mitra_temp_code;size:25" json:"mitra_temp_code"`
	MitraTempName                       string    `gorm:"column:mitra_temp_name;size:255" json:"mitra_temp_name"`
	MitraTempWebsite                    string    `gorm:"column:mitra_temp_website;size:255" json:"mitra_temp_website"`
	MitraTempEmail                      string    `gorm:"column:mitra_temp_email;size:255" json:"mitra_temp_email"`
	MitraTempLogo                       string    `gorm:"column:mitra_temp_logo;size:255" json:"mitra_temp_logo"`
	MitraProductTempName                string    `gorm:"column:mitra_prod_temp_name;size:125" json:"mitra_prod_temp_name"`
	MitraProductTempLimitMin            uint64    `gorm:"column:mitra_prod_temp_limit_min" json:"mitra_prod_temp_limit_min"`
	MitraProductTempLimitMax            uint64    `gorm:"column:mitra_prod_temp_limit_max" json:"mitra_prod_temp_limit_max"`
	MitraProductTempTenorMin            uint64    `gorm:"column:mitra_prod_temp_tenor_min" json:"mitra_prod_temp_tenor_min"`
	MitraProductTempTenorMax            uint64    `gorm:"column:mitra_prod_temp_tenor_max" json:"mitra_prod_temp_tenor_max"`
	MitraProductTempInterestRate        float64   `gorm:"column:mitra_prod_temp_interest_rate" json:"mitra_prod_temp_interest_rate"`
	MitraProductTempInterestRatePeriod  string    `gorm:"column:mitra_prod_temp_interest_rate_period;size:125" json:"mitra_prod_temp_interest_rate_period"`
	MitraProductTempInterestRateType    string    `gorm:"column:mitra_prod_temp_interest_rate_type;size:125" json:"mitra_prod_temp_interest_rate_type"`
	MitraProductTempInterestRateFormula string    `gorm:"column:mitra_prod_temp_interest_rate_formula;size:125" json:"mitra_prod_temp_interest_rate_formula"`
	MitraProductTempInterestRateDiscMax float64   `gorm:"column:mitra_prod_temp_interest_rate_disc_max" json:"mitra_prod_temp_interest_rate_disc_max"`
	MitraProductTempNumReviewer         uint64    `gorm:"column:mitra_prod_temp_num_reviewer" json:"mitra_prod_temp_num_reviewer"`
	MitraProductTempNumApprover         uint64    `gorm:"column:mitra_prod_temp_num_approver" json:"mitra_prod_temp_num_approver"`
	MitraProductTempDateActiveStart     string    `gorm:"column:mitra_prod_temp_date_active_start" json:"mitra_prod_temp_date_active_start"`
	MitraProductTempDateActiveFinish    string    `gorm:"column:mitra_prod_temp_date_active_finish" json:"mitra_prod_temp_date_active_finish"`
	MitraProductTempReason              string    `gorm:"column:mitra_prod_temp_reason" json:"mitra_prod_temp_reason"`
	MitraProductTempStatus              uint64    `gorm:"column:mitra_prod_temp_status;size:2" json:"mitra_prod_temp_status"`
	MitraProductTempActionBefore        uint64    `gorm:"column:mitra_prod_temp_action_before;size:2" json:"mitra_prod_temp_action_before"`
	MitraProductTempCreatedBy           string    `gorm:"column:mitra_prod_temp_created_by;size:125" json:"mitra_prod_temp_created_by"`
	MitraProductTempCreatedAt           time.Time `gorm:"column:mitra_prod_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_temp_created_at"`

	MitraProductSubCategoryTemp  []MitraProductSubCategoryTempPendData  `json:"mitra_prod_sub_ctgry_temp"`
	MitraProductCostTemp         []MitraProductCostTempPendData         `json:"mitra_prod_cost_temp"`
	MitraProductDocumentRuleTemp []MitraProductDocumentRuleTempPendData `json:"mitra_prod_doc_rule_temp"`

	MitraProductID                  uint64     `gorm:"column:mitra_prod_id" json:"mitra_prod_id"`
	MitraID                         uint64     `gorm:"column:mitra_id" json:"mitra_id"`
	MitraCode                       string     `gorm:"column:mitra_code;size:25" json:"mitra_code"`
	MitraName                       string     `gorm:"column:mitra_name;size:255" json:"mitra_name"`
	MitraWebsite                    string     `gorm:"column:mitra_website;size:255" json:"mitra_website"`
	MitraEmail                      string     `gorm:"column:mitra_email;size:255" json:"mitra_email"`
	MitraLogo                       string     `gorm:"column:mitra_logo;size:255" json:"mitra_logo"`
	CustomerTypeID                  uint64     `gorm:"column:cust_type_id" json:"cust_type_id"`
	CustomerTypeName                string     `gorm:"column:cust_type_name" json:"cust_type_name"`
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
	MitraProductDateActiveStart     string     `gorm:"column:mitra_prod_date_active_start" json:"mitra_prod_date_active_start"`
	MitraProductDateActiveFinish    string     `gorm:"column:mitra_prod_date_active_finish" json:"mitra_prod_date_active_finish"`
	MitraProductStatus              uint64     `gorm:"column:mitra_prod_status;size:2" json:"mitra_prod_status"`
	MitraProductCreatedBy           string     `gorm:"column:mitra_prod_created_by;size:125" json:"mitra_prod_created_by"`
	MitraProductCreatedAt           time.Time  `gorm:"column:mitra_prod_created_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_created_at"`
	MitraProductUpdatedBy           string     `gorm:"column:mitra_prod_created_by;size:125" json:"mitra_prod_created_by"`
	MitraProductUpdatedAt           *time.Time `gorm:"column:mitra_prod_created_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_created_at"`
	MitraProductDeletedBy           string     `gorm:"column:mitra_prod_created_by;size:125" json:"mitra_prod_created_by"`
	MitraProductDeletedAt           *time.Time `gorm:"column:mitra_prod_created_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_created_at"`

	MitraProductSubCategory  []MitraProductSubCategoryData  `json:"mitra_prod_sub_ctgry"`
	MitraProductCost         []MitraProductCostData         `json:"mitra_prod_cost"`
	MitraProductDocumentRule []MitraProductDocumentRuleData `json:"mitra_prod_doc_rule"`
}

type ResponseMitraProductTemps struct {
	Status  int                     `json:"status"`
	Message string                  `json:"message"`
	Data    *[]MitraProductTempData `json:"data"`
}

type ResponseMitraProductTemp struct {
	Status  int                   `json:"status"`
	Message string                `json:"message"`
	Data    *MitraProductTempData `json:"data"`
}

type ResponseMitraProductTempsPend struct {
	Status  int                 `json:"status"`
	Message string              `json:"message"`
	Data    *[]MitraProductTemp `json:"data"`
}

type ResponseMitraProductTempPend struct {
	Status  int                   `json:"status"`
	Message string                `json:"message"`
	Data    *MitraProductTempPend `json:"data"`
}

type ResponseMitraProductTempPendData struct {
	Status  int                       `json:"status"`
	Message string                    `json:"message"`
	Data    *MitraProductTempPendData `json:"data"`
}

type ResponseMitraProductTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraProductTemp) TableName() string {
	return "m_mitra_product_temp"
}

func (MitraProductTempPend) TableName() string {
	return "m_mitra_product_temp"
}

func (MitraProductTempPendData) TableName() string {
	return "m_mitra_product_temp"
}

func (MitraProductTempData) TableName() string {
	return "m_mitra_product_temp"
}

func (p *MitraProductTemp) Prepare() {
	p.MitraProductID = p.MitraProductID
	p.MitraTempID = p.MitraTempID
	p.CustomerTypeTempID = p.CustomerTypeTempID
	p.MitraProductTempName = html.EscapeString(strings.TrimSpace(p.MitraProductTempName))
	p.MitraProductTempLimitMin = p.MitraProductTempLimitMin
	p.MitraProductTempLimitMax = p.MitraProductTempLimitMax
	p.MitraProductTempTenorMin = p.MitraProductTempTenorMin
	p.MitraProductTempTenorMax = p.MitraProductTempTenorMax
	p.MitraProductTempInterestRate = p.MitraProductTempInterestRate
	p.MitraProductTempInterestRatePeriod = p.MitraProductTempInterestRatePeriod
	p.MitraProductTempInterestRateType = p.MitraProductTempInterestRateType
	p.MitraProductTempInterestRateFormula = p.MitraProductTempInterestRateFormula
	p.MitraProductTempInterestRateDiscMax = p.MitraProductTempInterestRateDiscMax
	p.MitraProductTempNumReviewer = p.MitraProductTempNumReviewer
	p.MitraProductTempNumApprover = p.MitraProductTempNumApprover
	p.MitraProductTempDateActiveStart = p.MitraProductTempDateActiveStart
	p.MitraProductTempDateActiveFinish = p.MitraProductTempDateActiveFinish
	p.MitraProductTempStatus = p.MitraProductTempStatus
	p.MitraProductTempReason = p.MitraProductTempReason
	p.MitraProductTempCreatedBy = p.MitraProductTempCreatedBy
	p.MitraProductTempCreatedAt = time.Now()
}

func (p *MitraProductTempPend) Prepare() {
	p.MitraProductID = nil
	p.MitraTempID = p.MitraTempID
	p.CustomerTypeTempID = p.CustomerTypeTempID
	p.MitraProductTempName = html.EscapeString(strings.TrimSpace(p.MitraProductTempName))
	p.MitraProductTempLimitMin = p.MitraProductTempLimitMin
	p.MitraProductTempLimitMax = p.MitraProductTempLimitMax
	p.MitraProductTempTenorMin = p.MitraProductTempTenorMin
	p.MitraProductTempTenorMax = p.MitraProductTempTenorMax
	p.MitraProductTempInterestRate = p.MitraProductTempInterestRate
	p.MitraProductTempInterestRatePeriod = p.MitraProductTempInterestRatePeriod
	p.MitraProductTempInterestRateType = p.MitraProductTempInterestRateType
	p.MitraProductTempInterestRateFormula = p.MitraProductTempInterestRateFormula
	p.MitraProductTempInterestRateDiscMax = p.MitraProductTempInterestRateDiscMax
	p.MitraProductTempNumReviewer = p.MitraProductTempNumReviewer
	p.MitraProductTempNumApprover = p.MitraProductTempNumApprover
	p.MitraProductTempDateActiveStart = p.MitraProductTempDateActiveStart
	p.MitraProductTempDateActiveFinish = p.MitraProductTempDateActiveFinish
	p.MitraProductTempStatus = p.MitraProductTempStatus
	p.MitraProductTempReason = p.MitraProductTempReason
	p.MitraProductTempCreatedBy = p.MitraProductTempCreatedBy
	p.MitraProductTempCreatedAt = time.Now()
}

func (p *MitraProductTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.MitraTempID == 0 {
			return errors.New("Required Quadrant ID")
		}
		if p.MitraProductTempName == "" {
			return errors.New("Required Mitra Product Name")
		}
		if p.MitraProductTempLimitMin == 0 {
			return errors.New("Required Mitra Product Name")
		}
		if p.MitraProductTempLimitMax == 0 {
			return errors.New("Required Mitra Product Name")
		}
		if p.MitraProductTempTenorMin == 0 {
			return errors.New("Required Mitra Product Name")
		}
		if p.MitraProductTempTenorMax == 0 {
			return errors.New("Required Mitra Product Name")
		}
		if p.MitraProductTempInterestRate == 0 {
			return errors.New("Required Mitra Product Name")
		}
		if p.MitraProductTempInterestRatePeriod == "" {
			return errors.New("Required Mitra Product Name")
		}
		if p.MitraProductTempInterestRateType == "" {
			return errors.New("Required Mitra Product Name")
		}
		if p.MitraProductTempInterestRateFormula == "" {
			return errors.New("Required Mitra Product Name")
		}
		if p.MitraProductTempInterestRateDiscMax == 0 {
			return errors.New("Required Mitra Product Name")
		}
		return nil
	}
}

func (p *MitraProductTemp) SaveMitraProductTemp(db *gorm.DB) (*MitraProductTemp, error) {
	var err error
	err = db.Debug().Model(&MitraProductTemp{}).Create(&p).Error
	if err != nil {
		return &MitraProductTemp{}, err
	}
	return p, nil
}

func (p *MitraProductTempPend) SaveMitraProductTempPend(db *gorm.DB) (*MitraProductTempPend, error) {
	var err error
	err = db.Debug().Model(&MitraProductTempPend{}).Create(&p).Error
	if err != nil {
		return &MitraProductTempPend{}, err
	}
	return p, nil
}

func (p *MitraProductTemp) FindAllMitraProductTemp(db *gorm.DB) (*[]MitraProductTemp, error) {
	var err error
	product := []MitraProductTemp{}
	err = db.Debug().Model(&MitraProductTemp{}).Limit(100).
		Select(`m_mitra_product_temp.mitra_prod_temp_id,
				m_mitra_product_temp.mitra_temp_id,
				m_mitra_product_temp.cust_type_temp_id,
				m_mitra_product_temp.mitra_prod_temp_name,
				m_mitra_product_temp.mitra_prod_temp_limit_min,
				m_mitra_product_temp.mitra_prod_temp_limit_max,
				m_mitra_product_temp.mitra_prod_temp_tenor_min,
				m_mitra_product_temp.mitra_prod_temp_tenor_max,
				m_mitra_product_temp.mitra_prod_temp_interest_rate,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_period,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_type,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_formula,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_disc_max,
				m_mitra_product_temp.mitra_prod_temp_num_reviewer,
				m_mitra_product_temp.mitra_prod_temp_num_approver,
				m_mitra_product_temp.mitra_prod_temp_date_active_start,
				m_mitra_product_temp.mitra_prod_temp_date_active_finish,
				m_mitra_product_temp.mitra_prod_temp_status,
				m_mitra_product_temp.mitra_prod_temp_action_before,
				m_mitra_product_temp.mitra_prod_temp_reason,
				m_mitra_product_temp.mitra_prod_temp_created_by,
				m_mitra_product_temp.mitra_prod_temp_created_at`).
		Find(&product).Error
	if err != nil {
		return &[]MitraProductTemp{}, err
	}
	return &product, nil
}

func (p *MitraProductTemp) FindAllMitraProductTempStatusPendingActive(db *gorm.DB, status uint64, mitra uint64) (*[]MitraProductTemp, error) {
	var err error
	product := []MitraProductTemp{}
	err = db.Debug().Model(&MitraProductTemp{}).Limit(100).
		Select(`m_mitra_product_temp.mitra_prod_temp_id,
				m_mitra_product_temp.mitra_temp_id,
				m_mitra_product_temp.cust_type_temp_id,
				m_mitra_product_temp.mitra_prod_temp_name,
				m_mitra_product_temp.mitra_prod_temp_limit_min,
				m_mitra_product_temp.mitra_prod_temp_limit_max,
				m_mitra_product_temp.mitra_prod_temp_tenor_min,
				m_mitra_product_temp.mitra_prod_temp_tenor_max,
				m_mitra_product_temp.mitra_prod_temp_interest_rate,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_period,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_type,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_formula,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_disc_max,
				m_mitra_product_temp.mitra_prod_temp_num_reviewer,
				m_mitra_product_temp.mitra_prod_temp_num_approver,
				m_mitra_product_temp.mitra_prod_temp_date_active_start,
				m_mitra_product_temp.mitra_prod_temp_date_active_finish,
				m_mitra_product_temp.mitra_prod_temp_status,
				m_mitra_product_temp.mitra_prod_temp_action_before,
				m_mitra_product_temp.mitra_prod_temp_reason,
				m_mitra_product_temp.mitra_prod_temp_created_by,
				m_mitra_product_temp.mitra_prod_temp_created_at`).
		Where("m_mitra_product_temp.mitra_prod_temp_status = ? AND m_mitra_product_temp.mitra_temp_id = ?", status, mitra).
		Find(&product).Error
	if err != nil {
		return &[]MitraProductTemp{}, err
	}
	return &product, nil
}

func (p *MitraProductTemp) FindAllMitraProductTempStatus(db *gorm.DB, status uint64, mitra uint64) (*[]MitraProductTempData, error) {
	var err error
	mitraproduct := []MitraProductTempData{}
	err = db.Debug().Model(&MitraProductTempData{}).Limit(100).
		Select(`m_mitra_product_temp.mitra_prod_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_code as mitra_temp_code,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_temp.mitra_website as mitra_temp_website,
				m_mitra_temp.mitra_email as mitra_temp_email,
				m_mitra_temp.mitra_logo as mitra_temp_logo,
				m_customer_type_temp.cust_type_id as cust_type_temp_id,
				m_customer_type_temp.cust_type_name as cust_type_temp_name,
				m_mitra_product_temp.mitra_prod_temp_name,
				m_mitra_product_temp.mitra_prod_temp_limit_min,
				m_mitra_product_temp.mitra_prod_temp_limit_max,
				m_mitra_product_temp.mitra_prod_temp_tenor_min,
				m_mitra_product_temp.mitra_prod_temp_tenor_max,
				m_mitra_product_temp.mitra_prod_temp_interest_rate,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_period,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_type,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_formula,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_disc_max,
				m_mitra_product_temp.mitra_prod_temp_num_reviewer,
				m_mitra_product_temp.mitra_prod_temp_num_approver,
				m_mitra_product_temp.mitra_prod_temp_date_active_start,
				m_mitra_product_temp.mitra_prod_temp_date_active_finish,
				m_mitra_product_temp.mitra_prod_temp_status,
				m_mitra_product_temp.mitra_prod_temp_action_before,
				m_mitra_product_temp.mitra_prod_temp_reason,
				m_mitra_product_temp.mitra_prod_temp_created_by,
				m_mitra_product_temp.mitra_prod_temp_created_at,
				m_mitra_product.mitra_prod_id,
				m_mitra.mitra_id as mitra_id,
				m_mitra.mitra_code as mitra_code,
				m_mitra.mitra_name as mitra_name,
				m_mitra.mitra_website as mitra_website,
				m_mitra.mitra_email as mitra_email,
				m_mitra.mitra_logo as mitra_logo,
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
				m_mitra_product.mitra_prod_created_at,
				m_mitra_product.mitra_prod_updated_by,
				m_mitra_product.mitra_prod_updated_at,
				m_mitra_product.mitra_prod_deleted_by,
				m_mitra_product.mitra_prod_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_product_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_customer_type m_customer_type_temp on m_mitra_product_temp.cust_type_temp_id=m_customer_type_temp.cust_type_id").
		Joins("JOIN m_mitra_product on m_mitra_product_temp.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_customer_type m_customer_type on m_mitra_product.cust_type_id=m_customer_type.cust_type_id").
		Joins("JOIN m_mitra on m_mitra_product.mitra_id=m_mitra.mitra_id").
		Where("m_mitra_product_temp.mitra_prod_temp_status = ? AND m_mitra_product_temp.mitra_temp_id = ?", status, mitra).
		Find(&mitraproduct).Error
	if err != nil {
		return &[]MitraProductTempData{}, err
	}
	return &mitraproduct, nil
}

func (p *MitraProductTemp) FindMitraProductTempDataByID(db *gorm.DB, pid uint64) (*MitraProductTemp, error) {
	var err error
	err = db.Debug().Model(&MitraProductTemp{}).Where("mitra_prod_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraProductTemp{}, err
	}
	return p, nil
}

func (p *MitraProductTemp) FindMitraProductTempByIDPendingActive(db *gorm.DB, pid uint64) (*MitraProductTempPend, error) {
	var err error
	product := MitraProductTempPend{}
	err = db.Debug().Model(&MitraProductTempPend{}).Limit(100).
		Select(`m_mitra_product_temp.mitra_prod_temp_id,
				m_mitra_product_temp.mitra_temp_id,
				m_mitra_product_temp.cust_type_temp_id,
				m_mitra_product_temp.mitra_prod_temp_name,
				m_mitra_product_temp.mitra_prod_temp_limit_min,
				m_mitra_product_temp.mitra_prod_temp_limit_max,
				m_mitra_product_temp.mitra_prod_temp_tenor_min,
				m_mitra_product_temp.mitra_prod_temp_tenor_max,
				m_mitra_product_temp.mitra_prod_temp_interest_rate,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_period,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_type,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_formula,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_disc_max,
				m_mitra_product_temp.mitra_prod_temp_num_reviewer,
				m_mitra_product_temp.mitra_prod_temp_num_approver,
				m_mitra_product_temp.mitra_prod_temp_date_active_start,
				m_mitra_product_temp.mitra_prod_temp_date_active_finish,
				m_mitra_product_temp.mitra_prod_temp_status,
				m_mitra_product_temp.mitra_prod_temp_action_before,
				m_mitra_product_temp.mitra_prod_temp_reason,
				m_mitra_product_temp.mitra_prod_temp_created_by,
				m_mitra_product_temp.mitra_prod_temp_created_at`).
		Where("m_mitra_product_temp.mitra_prod_temp_id = ?", pid).
		Find(&product).Error
	if err != nil {
		return &MitraProductTempPend{}, err
	}
	return &product, nil
}

func (p *MitraProductTemp) FindMitraProductTempStatusByIDPendingActive(db *gorm.DB, pid uint64, status uint64, mitra uint64) (*MitraProductTempPendData, error) {
	var err error
	product := MitraProductTempPendData{}
	err = db.Debug().Model(&MitraProductTempPendData{}).Limit(100).
		Select(`m_mitra_product_temp.mitra_prod_temp_id,
				m_mitra_product_temp.mitra_temp_id,
				m_mitra_product_temp.cust_type_temp_id,
				m_mitra_product_temp.mitra_prod_temp_name,
				m_mitra_product_temp.mitra_prod_temp_limit_min,
				m_mitra_product_temp.mitra_prod_temp_limit_max,
				m_mitra_product_temp.mitra_prod_temp_tenor_min,
				m_mitra_product_temp.mitra_prod_temp_tenor_max,
				m_mitra_product_temp.mitra_prod_temp_interest_rate,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_period,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_type,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_formula,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_disc_max,
				m_mitra_product_temp.mitra_prod_temp_num_reviewer,
				m_mitra_product_temp.mitra_prod_temp_num_approver,
				m_mitra_product_temp.mitra_prod_temp_date_active_start,
				m_mitra_product_temp.mitra_prod_temp_date_active_finish,
				m_mitra_product_temp.mitra_prod_temp_status,
				m_mitra_product_temp.mitra_prod_temp_action_before,
				m_mitra_product_temp.mitra_prod_temp_reason,
				m_mitra_product_temp.mitra_prod_temp_created_by,
				m_mitra_product_temp.mitra_prod_temp_created_at`).
		Where("m_mitra_product_temp.mitra_prod_temp_id = ? AND m_mitra_product_temp.mitra_prod_temp_status = ? AND m_mitra_product_temp.mitra_temp_id = ?", pid, status, mitra).
		Find(&product).Error
	if err != nil {
		return &MitraProductTempPendData{}, err
	}
	return &product, nil
}

func (p *MitraProductTemp) FindMitraProductTempByID(db *gorm.DB, pid uint64) (*MitraProductTempData, error) {
	var err error
	product := MitraProductTempData{}
	err = db.Debug().Model(&MitraProductTempData{}).Limit(100).
		Select(`m_mitra_product_temp.mitra_prod_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_code as mitra_temp_code,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_temp.mitra_website as mitra_temp_website,
				m_mitra_temp.mitra_email as mitra_temp_email,
				m_mitra_temp.mitra_logo as mitra_temp_logo,
				m_customer_type_temp.cust_type_id as cust_type_temp_id,
				m_customer_type_temp.cust_type_name as cust_type_temp_name,
				m_mitra_product_temp.mitra_prod_temp_name,
				m_mitra_product_temp.mitra_prod_temp_limit_min,
				m_mitra_product_temp.mitra_prod_temp_limit_max,
				m_mitra_product_temp.mitra_prod_temp_tenor_min,
				m_mitra_product_temp.mitra_prod_temp_tenor_max,
				m_mitra_product_temp.mitra_prod_temp_interest_rate,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_period,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_type,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_formula,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_disc_max,
				m_mitra_product_temp.mitra_prod_temp_num_reviewer,
				m_mitra_product_temp.mitra_prod_temp_num_approver,
				m_mitra_product_temp.mitra_prod_temp_date_active_start,
				m_mitra_product_temp.mitra_prod_temp_date_active_finish,
				m_mitra_product_temp.mitra_prod_temp_status,
				m_mitra_product_temp.mitra_prod_temp_action_before,
				m_mitra_product_temp.mitra_prod_temp_reason,
				m_mitra_product_temp.mitra_prod_temp_created_by,
				m_mitra_product_temp.mitra_prod_temp_created_at,
				m_mitra_product.mitra_prod_id,
				m_mitra.mitra_id as mitra_id,
				m_mitra.mitra_code as mitra_code,
				m_mitra.mitra_name as mitra_name,
				m_mitra.mitra_website as mitra_website,
				m_mitra.mitra_email as mitra_email,
				m_mitra.mitra_logo as mitra_logo,
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
				m_mitra_product.mitra_prod_created_at,
				m_mitra_product.mitra_prod_updated_by,
				m_mitra_product.mitra_prod_updated_at,
				m_mitra_product.mitra_prod_deleted_by,
				m_mitra_product.mitra_prod_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_product_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_customer_type m_customer_type_temp on m_mitra_product_temp.cust_type_temp_id=m_customer_type_temp.cust_type_id").
		Joins("JOIN m_mitra_product on m_mitra_product_temp.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_customer_type m_customer_type on m_mitra_product.cust_type_id=m_customer_type.cust_type_id").
		Joins("JOIN m_mitra on m_mitra_product.mitra_id=m_mitra.mitra_id").
		Where("m_mitra_product_temp.mitra_prod_temp_id = ?", pid).
		Find(&product).Error
	if err != nil {
		return &MitraProductTempData{}, err
	}
	return &product, nil
}

func (p *MitraProductTemp) FindMitraProductTempStatusByID(db *gorm.DB, pid uint64, status uint64, mitra uint64) (*MitraProductTempData, error) {
	var err error
	product := MitraProductTempData{}
	err = db.Debug().Model(&MitraProductTempData{}).Limit(100).
		Select(`m_mitra_product_temp.mitra_prod_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_code as mitra_temp_code,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_temp.mitra_website as mitra_temp_website,
				m_mitra_temp.mitra_email as mitra_temp_email,
				m_mitra_temp.mitra_logo as mitra_temp_logo,
				m_customer_type_temp.cust_type_id as cust_type_temp_id,
				m_customer_type_temp.cust_type_name as cust_type_temp_name,
				m_mitra_product_temp.mitra_prod_temp_name,
				m_mitra_product_temp.mitra_prod_temp_limit_min,
				m_mitra_product_temp.mitra_prod_temp_limit_max,
				m_mitra_product_temp.mitra_prod_temp_tenor_min,
				m_mitra_product_temp.mitra_prod_temp_tenor_max,
				m_mitra_product_temp.mitra_prod_temp_interest_rate,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_period,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_type,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_formula,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_disc_max,
				m_mitra_product_temp.mitra_prod_temp_num_reviewer,
				m_mitra_product_temp.mitra_prod_temp_num_approver,
				m_mitra_product_temp.mitra_prod_temp_date_active_start,
				m_mitra_product_temp.mitra_prod_temp_date_active_finish,
				m_mitra_product_temp.mitra_prod_temp_status,
				m_mitra_product_temp.mitra_prod_temp_action_before,
				m_mitra_product_temp.mitra_prod_temp_reason,
				m_mitra_product_temp.mitra_prod_temp_created_by,
				m_mitra_product_temp.mitra_prod_temp_created_at,
				m_mitra_product.mitra_prod_id,
				m_mitra.mitra_id as mitra_id,
				m_mitra.mitra_code as mitra_code,
				m_mitra.mitra_name as mitra_name,
				m_mitra.mitra_website as mitra_website,
				m_mitra.mitra_email as mitra_email,
				m_mitra.mitra_logo as mitra_logo,
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
				m_mitra_product.mitra_prod_created_at,
				m_mitra_product.mitra_prod_updated_by,
				m_mitra_product.mitra_prod_updated_at,
				m_mitra_product.mitra_prod_deleted_by,
				m_mitra_product.mitra_prod_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_product_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_customer_type m_customer_type_temp on m_mitra_product_temp.cust_type_temp_id=m_customer_type_temp.cust_type_id").
		Joins("JOIN m_mitra_product on m_mitra_product_temp.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_customer_type m_customer_type on m_mitra_product.cust_type_id=m_customer_type.cust_type_id").
		Joins("JOIN m_mitra on m_mitra_product.mitra_id=m_mitra.mitra_id").
		Where("m_mitra_product_temp.mitra_prod_temp_id = ? AND m_mitra_product_temp.mitra_prod_temp_status = ? AND m_mitra_product_temp.mitra_temp_id = ?", pid, status, mitra).
		Find(&product).Error
	if err != nil {
		return &MitraProductTempData{}, err
	}
	return &product, nil
}

func (p *MitraProductTemp) UpdateMitraProductTemp(db *gorm.DB, pid uint64) (*MitraProductTemp, error) {
	var err error
	db = db.Debug().Model(&MitraProductTemp{}).Where("mitra_prod_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_temp_id":                         p.MitraTempID,
			"cust_type_temp_id":                     p.CustomerTypeTempID,
			"mitra_prod_temp_name":                  p.MitraProductTempName,
			"mitra_prod_temp_limit_min":             p.MitraProductTempLimitMin,
			"mitra_prod_temp_limit_max":             p.MitraProductTempLimitMax,
			"mitra_prod_temp_tenor_min":             p.MitraProductTempTenorMin,
			"mitra_prod_temp_tenor_max":             p.MitraProductTempTenorMax,
			"mitra_prod_temp_interest_rate":         p.MitraProductTempInterestRate,
			"mitra_prod_temp_interest_rate_period":  p.MitraProductTempInterestRatePeriod,
			"mitra_prod_temp_interest_rate_type":    p.MitraProductTempInterestRateType,
			"mitra_prod_temp_interest_rate_formula": p.MitraProductTempInterestRateFormula,
			"mitra_prod_temp_num_reviewer":          p.MitraProductTempNumReviewer,
			"mitra_prod_temp_num_approver":          p.MitraProductTempNumApprover,
			"mitra_prod_temp_active_start":          p.MitraProductTempDateActiveStart,
			"mitra_prod_temp_active_finish":         p.MitraProductTempDateActiveFinish,
			"mitra_prod_temp_status":                p.MitraProductTempStatus,
			"mitra_prod_temp_reason":                p.MitraProductTempReason,
		},
	)
	err = db.Debug().Model(&MitraProductTemp{}).Where("mitra_prod_temp_id = ?", pid).Error
	if err != nil {
		return &MitraProductTemp{}, err
	}
	return p, nil
}

func (p *MitraProductTemp) UpdateMitraProductTempStatus(db *gorm.DB, pid uint64) (*MitraProductTemp, error) {
	var err error
	db = db.Debug().Model(&MitraProductTemp{}).Where("mitra_prod_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_prod_temp_reason": p.MitraProductTempReason,
			"mitra_prod_temp_status": p.MitraProductTempStatus,
		},
	)
	err = db.Debug().Model(&MitraProductTemp{}).Where("mitra_prod_temp_id = ?", pid).Error
	if err != nil {
		return &MitraProductTemp{}, err
	}
	return p, nil
}

func (p *MitraProductTemp) DeleteMitraProductTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraProductTemp{}).Where("mitra_prod_temp_id = ?", pid).Take(&MitraProductTemp{}).Delete(&MitraProductTemp{})
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

//ADDITIONAL
//====================================================================================================================================

func (p *MitraProductTemp) FindMitraProductTempByMitraID(db *gorm.DB, pid uint64, mitra uint64) (*MitraProductTempData, error) {
	var err error
	product := MitraProductTempData{}
	err = db.Debug().Model(&MitraProductTempData{}).Limit(100).
		Select(`m_mitra_product_temp.mitra_prod_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_code as mitra_temp_code,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_temp.mitra_website as mitra_temp_website,
				m_mitra_temp.mitra_email as mitra_temp_email,
				m_mitra_temp.mitra_logo as mitra_temp_logo,
				m_customer_type_temp.cust_type_id as cust_type_temp_id,
				m_customer_type_temp.cust_type_name as cust_type_temp_name,
				m_mitra_product_temp.mitra_prod_temp_name,
				m_mitra_product_temp.mitra_prod_temp_limit_min,
				m_mitra_product_temp.mitra_prod_temp_limit_max,
				m_mitra_product_temp.mitra_prod_temp_tenor_min,
				m_mitra_product_temp.mitra_prod_temp_tenor_max,
				m_mitra_product_temp.mitra_prod_temp_interest_rate,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_period,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_type,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_formula,
				m_mitra_product_temp.mitra_prod_temp_interest_rate_disc_max,
				m_mitra_product_temp.mitra_prod_temp_num_reviewer,
				m_mitra_product_temp.mitra_prod_temp_num_approver,
				m_mitra_product_temp.mitra_prod_temp_date_active_start,
				m_mitra_product_temp.mitra_prod_temp_date_active_finish,
				m_mitra_product_temp.mitra_prod_temp_status,
				m_mitra_product_temp.mitra_prod_temp_action_before,
				m_mitra_product_temp.mitra_prod_temp_reason,
				m_mitra_product_temp.mitra_prod_temp_created_by,
				m_mitra_product_temp.mitra_prod_temp_created_at,
				m_mitra_product.mitra_prod_id,
				m_mitra.mitra_id as mitra_id,
				m_mitra.mitra_code as mitra_code,
				m_mitra.mitra_name as mitra_name,
				m_mitra.mitra_website as mitra_website,
				m_mitra.mitra_email as mitra_email,
				m_mitra.mitra_logo as mitra_logo,
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
				m_mitra_product.mitra_prod_created_at,
				m_mitra_product.mitra_prod_updated_by,
				m_mitra_product.mitra_prod_updated_at,
				m_mitra_product.mitra_prod_deleted_by,
				m_mitra_product.mitra_prod_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_product_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_customer_type m_customer_type_temp on m_mitra_product_temp.cust_type_temp_id=m_customer_type_temp.cust_type_id").
		Joins("JOIN m_mitra_product on m_mitra_product_temp.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_customer_type m_customer_type on m_mitra_product.cust_type_id=m_customer_type.cust_type_id").
		Joins("JOIN m_mitra on m_mitra_product.mitra_id=m_mitra.mitra_id").
		Where("m_mitra_product_temp.mitra_prod_temp_id = ? AND m_mitra_product_temp.mitra_temp_id = ?", pid, mitra).
		Find(&product).Error
	if err != nil {
		return &MitraProductTempData{}, err
	}
	return &product, nil
}
