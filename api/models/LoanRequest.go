package models

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type LoanRequest struct {
	LoanRequestID          uint64  `gorm:"column:loan_req_id;primary_key;" json:"loan_req_id"`
	MitraID                uint64  `gorm:"column:mitra_id" json:"mitra_id"`
	CustUserID             uint64  `gorm:"column:cust_user_id" json:"cust_user_id"`
	MitraUserID            *uint64 `gorm:"column:mitra_user_id" json:"mitra_user_id"`
	MitraBranchID          *uint64 `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	MitraProductID         *uint64 `gorm:"column:mitra_prod_id" json:"mitra_prod_id"`
	ProductSubCategoryID   uint64  `gorm:"column:prod_sub_ctgry_id" json:"prod_sub_ctgry_id"`
	LoanRequestNum         string  `gorm:"column:loan_req_num" json:"loan_req_num"`
	LoanRequestAmount      uint64  `gorm:"column:loan_req_amount" json:"loan_req_amount"`
	LoanRequestLong        uint64  `gorm:"column:loan_req_long" json:"loan_req_long"`
	LoanRequestPurpose     string  `gorm:"column:loan_req_purpose" json:"loan_req_purpose"`
	LoanRequestSalaryRange string  `gorm:"column:loan_req_salary_range;size:50" json:"loan_req_salary_range"`
	// LoanRequestPaymentPeriod    *string    `gorm:"column:loan_req_payment_period" json:"loan_req_payment_period"`
	LoanRequestReferral    *string    `gorm:"column:loan_req_referral" json:"loan_req_referral"`
	LoanRequestStatus      uint64     `gorm:"column:loan_req_status;size:2" json:"loan_req_status"`
	LoanRequestSchdlStatus uint64     `gorm:"column:loan_req_schdl_status;size:2" json:"loan_req_schdl_status"`
	LoanRequestDate        string     `gorm:"column:loan_req_date;size:10;" json:"loan_req_date"`
	LoanRequestSchdlMom    *string    `gorm:"column:loan_req_schdl_mom" json:"loan_req_schdl_mom"`
	LoanRequestJob         string     `gorm:"column:loan_req_job;size:24;" json:"loan_req_job"`
	LoanRequestCreatedAt   time.Time  `gorm:"column:loan_req_created_at;default:CURRENT_TIMESTAMP" json:"loan_req_created_at"`
	LoanRequestCreatedBy   string     `gorm:"column:loan_req_created_by;default:CURRENT_TIMESTAMP" json:"loan_req_created_by"`
	LoanRequestUpdatedBy   string     `gorm:"column:loan_req_updated_by;default:CURRENT_TIMESTAMP" json:"loan_req_updated_by"`
	LoanRequestUpdatedAt   *time.Time `gorm:"column:loan_req_updated_at;default:CURRENT_TIMESTAMP" json:"loan_req_updated_at"`
	LoanRequestDeletedBy   string     `gorm:"column:loan_req_deleted_by;default:CURRENT_TIMESTAMP" json:"loan_req_deleted_by"`
	LoanRequestDeletedAt   *time.Time `gorm:"column:loan_req_deleted_at;default:CURRENT_TIMESTAMP" json:"loan_req_deleted_at"`

	CustUserPhone string `gorm:"-" json:"cust_user_phone"`
	// MessageErrors []ErrorMessage `gorm:"-" json:"message_error"`
	MessageError string `gorm:"-" json:"message_error"`
}

type ErrorMessage struct {
	MessageError string `gorm:"-" json:"message_error"`
}

type LoanOffering struct {
	MitraID           uint64 `gorm:"column:mitra_id" json:"mitra_id"`
	MitraProductID    uint64 `gorm:"column:mitra_prod_id" json:"mitra_prod_id"`
	CustomerSubTypeID uint64 `gorm:"column:cust_sub_type_id;" json:"cust_sub_type_id"`

	// LoanRequestPaymentPeriod string `gorm:"column:loan_req_payment_period" json:"loan_req_payment_period"`

	CustUserPhone string `gorm:"-" json:"cust_user_phone"`
}
type LoanRequestData struct {
	LoanRequestID          uint64    `gorm:"column:loan_req_id" json:"loan_req_id"`
	ProductSubCategoryID   uint64    `gorm:"column:prod_sub_ctgry_id" json:"prod_sub_ctgry_id"`
	ProductSubCategoryName string    `gorm:"column:prod_sub_ctgry_name" json:"prod_sub_ctgry_name"`
	CustUserID             uint64    `gorm:"column:cust_user_id" json:"cust_user_id"`
	CustomerName           string    `gorm:"column:cust_name" json:"cust_name"`
	CustomerSex            string    `gorm:"column:cust_sex" json:"cust_sex"`
	CustomerBirthDate      string    `gorm:"column:cust_birth_date;size:25" json:"cust_birth_date"`
	CustomerMaritalStatus  string    `gorm:"column:cust_marital_status" json:"cust_marital_status"`
	CustomerAge            string    `gorm:"column:cust_age" json:"cust_age"`
	CustomerProvince       string    `gorm:"column:cust_province" json:"cust_province"`
	CustomerRegency        string    `gorm:"column:cust_regency" json:"cust_regency"`
	MitraID                uint64    `gorm:"column:mitra_id" json:"mitra_id"`
	LoanRequestAmount      uint64    `gorm:"column:loan_req_amount" json:"loan_req_amount"`
	LoanRequestLong        uint64    `gorm:"column:loan_req_long" json:"loan_req_long"`
	LoanRequestSalaryRange string    `gorm:"column:loan_req_salary_range" json:"loan_req_salary_range"`
	LoanRequestPurpose     string    `gorm:"column:loan_req_purpose" json:"loan_req_purpose"`
	LoanRequestJob         string    `gorm:"column:loan_req_job;size:24;" json:"loan_req_job"`
	LoanRequestStatus      uint64    `gorm:"column:loan_req_status" json:"loan_req_status"`
	LoanRequestDate        string    `gorm:"column:loan_req_date;size:10;" json:"loan_req_date"`
	LoanRequestCreatedAt   time.Time `gorm:"column:loan_req_created_at;default:CURRENT_TIMESTAMP" json:"loan_req_created_at"`
}

type LoanRequestScheduledData struct {
	LoanRequestID            uint64    `gorm:"column:loan_req_id" json:"loan_req_id"`
	LoanRequestScheduleDate  string    `gorm:"column:loan_req_schdl_date_1" json:"loan_req_schdl_date_1"`
	LoanRequestSchedulePlace string    `gorm:"column:loan_req_schdl_place" json:"loan_req_schdl_place"`
	LoanRequestScheduleTime  string    `gorm:"column:loan_req_schdl_time" json:"loan_req_schdl_time"`
	ProductSubCategoryID     uint64    `gorm:"column:prod_sub_ctgry_id" json:"prod_sub_ctgry_id"`
	ProductSubCategoryName   string    `gorm:"column:prod_sub_ctgry_name" json:"prod_sub_ctgry_name"`
	CustUserID               uint64    `gorm:"column:cust_user_id" json:"cust_user_id"`
	CustomerName             string    `gorm:"column:cust_name" json:"cust_name"`
	CustomerSex              string    `gorm:"column:cust_sex" json:"cust_sex"`
	CustomerMaritalStatus    string    `gorm:"column:cust_marital_status" json:"cust_marital_status"`
	CustomerAge              string    `gorm:"column:cust_age" json:"cust_age"`
	CustomerProvince         string    `gorm:"column:cust_province" json:"cust_province"`
	CustomerRegency          string    `gorm:"column:cust_regency" json:"cust_regency"`
	MitraID                  uint64    `gorm:"column:mitra_id" json:"mitra_id"`
	LoanRequestAmount        uint64    `gorm:"column:loan_req_amount" json:"loan_req_amount"`
	LoanRequestLong          uint64    `gorm:"column:loan_req_long" json:"loan_req_long"`
	LoanRequestStatus        uint64    `gorm:"column:loan_req_status" json:"loan_req_status"`
	LoanRequestDate          string    `gorm:"column:loan_req_date;size:10;" json:"loan_req_date"`
	LoanRequestCreatedAt     time.Time `gorm:"column:loan_req_created_at;default:CURRENT_TIMESTAMP" json:"loan_req_created_at"`
}

type LoanRequestDetailData struct {
	LoanRequestID           uint64    `gorm:"column:loan_req_id" json:"loan_req_id"`
	ProductSubCategoryID    uint64    `gorm:"column:prod_sub_ctgry_id" json:"prod_sub_ctgry_id"`
	ProductSubCategoryName  string    `gorm:"column:prod_sub_ctgry_name" json:"prod_sub_ctgry_name"`
	CustUserID              uint64    `gorm:"column:cust_user_id" json:"cust_user_id"`
	CustomerName            string    `gorm:"column:cust_name" json:"cust_name"`
	CustomerSex             string    `gorm:"column:cust_sex" json:"cust_sex"`
	CustomerBirthDate       string    `gorm:"column:cust_birth_date" json:"cust_birth_date"`
	CustomerUserPhone       string    `gorm:"column:cust_user_phone" json:"cust_user_phone"`
	CustomerMaritalStatus   string    `gorm:"column:cust_marital_status" json:"cust_marital_status"`
	CustomerAddressStreet   string    `gorm:"column:cust_adrs_street" json:"cust_adrs_street"`
	CustomerAddressRT       string    `gorm:"column:cust_adrs_rt" json:"cust_adrs_rt"`
	CustomerAddressRW       string    `gorm:"column:cust_adrs_rw" json:"cust_adrs_rw"`
	CustomerAddressZipCode  string    `gorm:"column:cust_adrs_zip_code" json:"cust_adrs_zip_code"`
	CountryID               uint64    `gorm:"column:country_id" json:"country_id"`
	CountryName             string    `gorm:"column:country_name" json:"country_name"`
	ProvinceID              uint64    `gorm:"column:province_id;" json:"province_id"`
	ProvinceName            string    `gorm:"column:province_name" json:"province_name"`
	RegencyID               uint64    `gorm:"column:regency_id" json:"regency_id"`
	RegencyName             string    `gorm:"column:regency_name" json:"regency_name"`
	MitraID                 uint64    `gorm:"column:mitra_id" json:"mitra_id"`
	MitraUserID             uint64    `gorm:"column:mitra_user_id" json:"mitra_user_id"`
	MitraUserPhone          string    `gorm:"column:mitra_user_phone;size:20" json:"mitra_user_phone"`
	LoanRequestNum          string    `gorm:"column:loan_req_num" json:"loan_req_num"`
	LoanRequestAmount       uint64    `gorm:"column:loan_req_amount" json:"loan_req_amount"`
	LoanRequestLong         uint64    `gorm:"column:loan_req_long" json:"loan_req_long"`
	LoanRequestPurpose      string    `gorm:"column:loan_req_purpose" json:"loan_req_purpose"`
	LoanRequestSalaryRange  string    `gorm:"column:loan_req_salary_range" json:"loan_req_salary_range"`
	LoanRequestStatus       uint64    `gorm:"column:loan_req_status" json:"loan_req_status"`
	LoanRequestDate         string    `gorm:"column:loan_req_date;size:10;" json:"loan_req_date"`
	LoanRequestJob          string    `gorm:"column:loan_req_job;size:24;" json:"loan_req_job"`
	LoanRequestCreatedAt    time.Time `gorm:"column:loan_req_created_at;default:CURRENT_TIMESTAMP" json:"loan_req_created_at"`
	LoanRequestCollID       uint64    `gorm:"column:loan_req_coll_id" json:"loan_req_coll_id"`
	LoanRequestCollName     string    `gorm:"column:loan_req_coll_name" json:"loan_req_coll_name"`
	LoanRequestCollDesc     string    `gorm:"column:loan_req_coll_desc" json:"loan_req_coll_name_desc"`
	LoanReqeustCollStatus   uint64    `gorm:"column:loan_req_coll_status" json:"loan_req_coll_status"`
	LoanRequestCollImgID    uint64    `gorm:"column:loan_req_coll_img_id" json:"loan_req_coll_img_id"`
	LoanRequestCollImgValue string    `gorm:"column:loan_req_coll_img_value" json:"loan_req_coll_img_value"`
}

type LoanOfferDetailData struct {
	LoanRequestID                  uint64  `gorm:"column:loan_req_id" json:"loan_req_id"`
	ProductSubCategoryID           uint64  `gorm:"column:prod_sub_ctgry_id" json:"prod_sub_ctgry_id"`
	ProductSubCategoryName         string  `gorm:"column:prod_sub_ctgry_name" json:"prod_sub_ctgry_name"`
	CustUserID                     uint64  `gorm:"column:cust_user_id" json:"cust_user_id"`
	CustomerName                   string  `gorm:"column:cust_name" json:"cust_name"`
	CustomerSex                    string  `gorm:"column:cust_sex" json:"cust_sex"`
	CustomerBirthDate              string  `gorm:"column:cust_birth_date" json:"cust_birth_date"`
	CustomerUserPhone              string  `gorm:"column:cust_user_phone" json:"cust_user_phone"`
	CountryID                      uint64  `gorm:"column:country_id" json:"country_id"`
	CountryName                    string  `gorm:"column:country_name" json:"country_name"`
	ProvinceID                     uint64  `gorm:"column:province_id;" json:"province_id"`
	ProvinceName                   string  `gorm:"column:province_name" json:"province_name"`
	RegencyID                      uint64  `gorm:"column:regency_id" json:"regency_id"`
	RegencyName                    string  `gorm:"column:regency_name" json:"regency_name"`
	MitraID                        uint64  `gorm:"column:mitra_id" json:"mitra_id"`
	MitraName                      string  `gorm:"column:mitra_name" json:"mitra_name"`
	MitraLogo                      string  `gorm:"column:mitra_logo" json:"mitra_logo"`
	MitraBranchID                  uint64  `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	MitraUserID                    uint64  `gorm:"column:mitra_user_id" json:"mitra_user_id"`
	MitraEmployeeName              string  `gorm:"column:mitra_emp_name" json:"mitra_emp_name"`
	MitraProductID                 uint64  `gorm:"column:mitra_prod_id" json:"mitra_prod_id"`
	MitraProductName               string  `gorm:"column:mitra_prod_name" json:"mitra_prod_name"`
	MitraProductInterestRate       float64 `gorm:"column:mitra_prod_interest_rate" json:"mitra_prod_interest_rate"`
	MitraProductInterestRatePeriod string  `gorm:"column:mitra_prod_interest_rate_period" json:"mitra_prod_interest_rate_period"`
	MitraProductInterestRateType   string  `gorm:"column:mitra_prod_interest_rate_type" json:"mitra_prod_interest_rate_type"`
	LoanRequestNum                 string  `gorm:"column:loan_req_num" json:"loan_req_num"`
	LoanRequestAmount              uint64  `gorm:"column:loan_req_amount" json:"loan_req_amount"`
	LoanRequestLong                uint64  `gorm:"column:loan_req_long" json:"loan_req_long"`
	LoanRequestPurpose             string  `gorm:"column:loan_req_purpose" json:"loan_req_purpose"`
	LoanRequestSalaryRange         string  `gorm:"column:loan_req_salary_range" json:"loan_req_salary_range"`
	// LoanRequestPaymentPeriod       string    `gorm:"column:loan_req_payment_period" json:"loan_req_payment_period"`
	LoanRequestStatus      uint64    `gorm:"column:loan_req_status" json:"loan_req_status"`
	LoanRequestSchdlStatus uint64    `gorm:"column:loan_req_schdl_status" json:"loan_req_schdl_status"`
	LoanRequestReferral    string    `gorm:"column:loan_req_referral" json:"loan_req_referral"`
	LoanRequestDate        string    `gorm:"column:loan_req_date;size:10;" json:"loan_req_date"`
	LoanRequestCreatedAt   time.Time `gorm:"column:loan_req_created_at;default:CURRENT_TIMESTAMP" json:"loan_req_created_at"`
}

type LoanRequestAcceptedData struct {
	LoanRequestID          uint64    `gorm:"column:loan_req_id" json:"loan_req_id"`
	ProductSubCategoryID   uint64    `gorm:"column:prod_sub_ctgry_id" json:"prod_sub_ctgry_id"`
	ProductSubCategoryName string    `gorm:"column:prod_sub_ctgry_name" json:"prod_sub_ctgry_name"`
	CustUserID             uint64    `gorm:"column:cust_user_id" json:"cust_user_id"`
	CustomerName           string    `gorm:"column:cust_name" json:"cust_name"`
	CountryID              uint64    `gorm:"column:country_id" json:"country_id"`
	CountryName            string    `gorm:"column:country_name" json:"country_name"`
	ProvinceID             uint64    `gorm:"column:province_id;" json:"province_id"`
	ProvinceName           string    `gorm:"column:province_name" json:"province_name"`
	RegencyID              uint64    `gorm:"column:regency_id" json:"regency_id"`
	RegencyName            string    `gorm:"column:regency_name" json:"regency_name"`
	MitraID                uint64    `gorm:"column:mitra_id" json:"mitra_id"`
	MitraUserID            uint64    `gorm:"column:mitra_user_id" json:"mitra_user_id"`
	LoanRequestAmount      uint64    `gorm:"column:loan_req_amount" json:"loan_req_amount"`
	LoanRequestLong        uint64    `gorm:"column:loan_req_long" json:"loan_req_long"`
	LoanRequestPurpose     string    `gorm:"column:loan_req_purpose" json:"loan_req_purpose"`
	LoanRequestSalaryRange string    `gorm:"column:loan_req_salary_range" json:"loan_req_salary_range"`
	LoanRequestStatus      uint64    `gorm:"column:loan_req_status" json:"loan_req_status"`
	LoanRequestDate        string    `gorm:"column:loan_req_date;size:10;" json:"loan_req_date"`
	LoanRequestCreatedAt   time.Time `gorm:"column:loan_req_created_at;default:CURRENT_TIMESTAMP" json:"loan_req_created_at"`
}

type LoanRequestAcceptedDetailData struct {
	LoanRequestID           uint64    `gorm:"column:loan_req_id" json:"loan_req_id"`
	ProductSubCategoryID    uint64    `gorm:"column:prod_sub_ctgry_id" json:"prod_sub_ctgry_id"`
	ProductSubCategoryName  string    `gorm:"column:prod_sub_ctgry_name" json:"prod_sub_ctgry_name"`
	CustUserID              uint64    `gorm:"column:cust_user_id" json:"cust_user_id"`
	CustomerName            string    `gorm:"column:cust_name" json:"cust_name"`
	CustomerUserPhone       string    `gorm:"column:cust_user_phone;size:20" json:"cust_user_phone" validate:"length(11|13)"`
	CustomerSex             string    `gorm:"column:cust_sex" json:"cust_sex"`
	CustomerBirthDate       string    `gorm:"column:cust_birth_date" json:"cust_birth_date"`
	CustomerMaritalStatus   string    `gorm:"column:cust_marital_status" json:"cust_marital_status"`
	CustomerAddressStreet   string    `gorm:"column:cust_adrs_street" json:"cust_adrs_street"`
	CustomerAddressRT       string    `gorm:"column:cust_adrs_rt" json:"cust_adrs_rt"`
	CustomerAddressRW       string    `gorm:"column:cust_adrs_rw" json:"cust_adrs_rw"`
	CustomerAddressZipCode  string    `gorm:"column:cust_adrs_zip_code" json:"cust_adrs_zip_code"`
	CountryID               uint64    `gorm:"column:country_id" json:"country_id"`
	CountryName             string    `gorm:"column:country_name" json:"country_name"`
	ProvinceID              uint64    `gorm:"column:province_id;" json:"province_id"`
	ProvinceName            string    `gorm:"column:province_name" json:"province_name"`
	RegencyID               uint64    `gorm:"column:regency_id" json:"regency_id"`
	RegencyName             string    `gorm:"column:regency_name" json:"regency_name"`
	MitraID                 uint64    `gorm:"column:mitra_id" json:"mitra_id"`
	MitraName               string    `gorm:"column:mitra_name" json:"mitra_name"`
	MitraUserID             uint64    `gorm:"column:mitra_user_id" json:"mitra_user_id"`
	MitraUserPhone          string    `gorm:"column:mitra_user_phone;size:20" json:"mitra_user_phone"`
	MitraEmployeeName       string    `gorm:"column:mitra_emp_name" json:"mitra_emp_name"`
	LoanRequestAmount       uint64    `gorm:"column:loan_req_amount" json:"loan_req_amount""`
	LoanRequestLong         uint64    `gorm:"column:loan_req_long" json:"loan_req_long"`
	LoanRequestPurpose      string    `gorm:"column:loan_req_purpose" json:"loan_req_purpose"`
	LoanRequestSalaryRange  string    `gorm:"column:loan_req_salary_range" json:"loan_req_salary_range"`
	LoanRequestStatus       uint64    `gorm:"column:loan_req_status" json:"loan_req_status"`
	LoanRequestDate         string    `gorm:"column:loan_req_date;size:10;" json:"loan_req_date"`
	LoanRequestJob          string    `gorm:"column:loan_req_job;size:24;" json:"loan_req_job"`
	LoanRequestCreatedAt    time.Time `gorm:"column:loan_req_created_at;default:CURRENT_TIMESTAMP" json:"loan_req_created_at"`
	LoanRequestCollID       uint64    `gorm:"column:loan_req_coll_id" json:"loan_req_coll_id"`
	LoanRequestCollName     string    `gorm:"column:loan_req_coll_name" json:"loan_req_coll_name"`
	LoanRequestCollDesc     string    `gorm:"column:loan_req_coll_desc" json:"loan_req_coll_name_desc"`
	LoanReqeustCollStatus   uint64    `gorm:"column:loan_req_coll_status" json:"loan_req_coll_status"`
	LoanRequestCollImgID    uint64    `gorm:"column:loan_req_coll_img_id" json:"loan_req_coll_img_id"`
	LoanRequestCollImgValue string    `gorm:"column:loan_req_coll_img_value" json:"loan_req_coll_img_value"`
}

type CustomerLoanStatusData struct {
	LoanRequestID          uint64     `gorm:"column:loan_req_id" json:"loan_req_id"`
	ProductSubCategoryID   uint64     `gorm:"column:prod_sub_ctgry_id" json:"prod_sub_ctgry_id"`
	ProductSubCategoryName string     `gorm:"column:prod_sub_ctgry_name" json:"prod_sub_ctgry_name"`
	MitraID                uint64     `gorm:"column:mitra_id" json:"mitra_id"`
	MitraName              string     `gorm:"column:mitra_name" json:"mitra_name"`
	MitraLogo              string     `gorm:"column:mitra_logo" json:"mitra_logo"`
	MitraUserID            uint64     `gorm:"column:mitra_user_id" json:"mitra_user_id"`
	MitraUserPhone         string     `gorm:"column:mitra_user_phone;size:20" json:"mitra_user_phone"`
	LoanRequestLong        uint64     `gorm:"column:loan_req_long" json:"loan_req_long"`
	LoanRequestStatus      uint64     `gorm:"column:loan_req_status" json:"loan_req_status"`
	LoanRequestSchdlStatus uint64     `gorm:"column:loan_req_schdl_status" json:"loan_req_schdl_status"`
	LoanRequestAmount      uint64     `gorm:"column:loan_req_amount" json:"loan_req_amount"`
	LoanRequestDate        string     `gorm:"column:loan_req_date;size:10;" json:"loan_req_date"`
	LoanRequestCreatedAt   time.Time  `gorm:"column:loan_req_created_at;default:CURRENT_TIMESTAMP" json:"loan_req_created_at"`
	LoanRequestUpdatedAt   *time.Time `gorm:"column:loan_req_updated_at;default:CURRENT_TIMESTAMP" json:"loan_req_updated_at"`
	LoanRequestDeletedAt   *time.Time `gorm:"column:loan_req_deleted_at;default:CURRENT_TIMESTAMP" json:"loan_req_deleted_at"`
}

type CustomerLoanStatusDetailData struct {
	LoanRequestID           uint64    `gorm:"column:loan_req_id" json:"loan_req_id"`
	ProductSubCategoryID    uint64    `gorm:"column:prod_sub_ctgry_id" json:"prod_sub_ctgry_id"`
	ProductSubCategoryName  string    `gorm:"column:prod_sub_ctgry_name" json:"prod_sub_ctgry_name"`
	CustUserID              uint64    `gorm:"column:cust_user_id" json:"cust_user_id"`
	CustomerName            string    `gorm:"column:cust_name" json:"cust_name"`
	CustomerSex             string    `gorm:"column:cust_sex" json:"cust_sex"`
	CustomerMaritalStatus   string    `gorm:"column:cust_marital_status" json:"cust_marital_status"`
	CustomerAddressStreet   string    `gorm:"column:cust_adrs_street" json:"cust_adrs_street"`
	MitraID                 uint64    `gorm:"column:mitra_id" json:"mitra_id"`
	MitraName               string    `gorm:"column:mitra_name" json:"mitra_name"`
	MitraLogo               string    `gorm:"column:mitra_logo" json:"mitra_logo"`
	CountryID               uint64    `gorm:"column:country_id" json:"country_id"`
	CountryName             string    `gorm:"column:country_name" json:"country_name"`
	ProvinceID              uint64    `gorm:"column:province_id;" json:"province_id"`
	ProvinceName            string    `gorm:"column:province_name" json:"province_name"`
	RegencyID               uint64    `gorm:"column:regency_id" json:"regency_id"`
	RegencyName             string    `gorm:"column:regency_name" json:"regency_name"`
	MitraUserID             uint64    `gorm:"column:mitra_user_id" json:"mitra_user_id"`
	MitraUserPhone          string    `gorm:"column:mitra_user_phone;size:20" json:"mitra_user_phone"`
	LoanRequestLong         uint64    `gorm:"column:loan_req_long" json:"loan_req_long"`
	LoanRequestAmount       uint64    `gorm:"column:loan_req_amount" json:"loan_req_amount"`
	LoanRequestSalaryRange  string    `gorm:"column:loan_req_salary_range" json:"loan_req_salary_range"`
	LoanRequestStatus       uint64    `gorm:"column:loan_req_status" json:"loan_req_status"`
	LoanRequestDate         string    `gorm:"column:loan_req_date;size:10;" json:"loan_req_date"`
	LoanRequestPurpose      string    `gorm:"column:loan_req_purpose" json:"loan_req_purpose"`
	LoanRequestJob          string    `gorm:"column:loan_req_job;size:24;" json:"loan_req_job"`
	LoanRequestSchdlStatus  uint64    `gorm:"column:loan_req_schdl_status" json:"loan_req_schdl_status"`
	LoanRequestCreatedAt    time.Time `gorm:"column:loan_req_created_at;default:CURRENT_TIMESTAMP" json:"loan_req_created_at"`
	LoanRequestCollID       uint64    `gorm:"column:loan_req_coll_id" json:"loan_req_coll_id"`
	LoanRequestCollName     string    `gorm:"column:loan_req_coll_name" json:"loan_req_coll_name"`
	LoanRequestCollDesc     string    `gorm:"column:loan_req_coll_desc" json:"loan_req_coll_name_desc"`
	LoanReqeustCollStatus   uint64    `gorm:"column:loan_req_coll_status" json:"loan_req_coll_status"`
	LoanRequestCollImgID    uint64    `gorm:"column:loan_req_coll_img_id" json:"loan_req_coll_img_id"`
	LoanRequestCollImgValue string    `gorm:"column:loan_req_coll_img_value" json:"loan_req_coll_img_value"`
}

type LoanRequestCountData struct {
	LoanRequestCount  int `gorm:"-" json:"loan_request_count"`
	LoanRequestAmount int `gorm:"column:loan_req_amount" json:"loan_request_amount"`
}

type ResponseGenerateCode struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    *GenerateCode `json:"data"`
}

type ResponseLoanRequests struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    *[]LoanRequestData `json:"data"`
}

type ResponseLoanSchedullingRequests struct {
	Status  int                         `json:"status"`
	Message string                      `json:"message"`
	Data    *[]LoanRequestScheduledData `json:"data"`
}

type ResponseLoanRequest struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    *LoanRequestData `json:"data"`
}

type ResponseLoanRequestIU struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    *LoanRequest `json:"data"`
}

type ResponseLoanRequestError struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    *LoanRequest `json:"data"`
}

type ResponseLoanRequestDel struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResponseLoanRequestDetails struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    *[]LoanRequestDetailData `json:"data"`
}

type ResponseLoanRequestDetail struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    *LoanRequestDetailData `json:"data"`
}

type ResponseLoanRequestAccepteds struct {
	Status  int                        `json:"status"`
	Message string                     `json:"message"`
	Data    *[]LoanRequestAcceptedData `json:"data"`
}

type ResponseLoanRequestAccepted struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    *LoanRequestAcceptedData `json:"data"`
}

type ResponseLoanRequestAcceptedDetails struct {
	Status  int                              `json:"status"`
	Message string                           `json:"message"`
	Data    *[]LoanRequestAcceptedDetailData `json:"data"`
}

type ResponseLoanRequestAcceptedDetail struct {
	Status  int                            `json:"status"`
	Message string                         `json:"message"`
	Data    *LoanRequestAcceptedDetailData `json:"data"`
}

type ResponseCustomerLoanStatusDetails struct {
	Status  int                       `json:"status"`
	Message string                    `json:"message"`
	Data    *[]CustomerLoanStatusData `json:"data"`
}

type ResponseCustomerLoanStatusCount struct {
	Status  int                   `json:"status"`
	Message string                `json:"message"`
	Data    *LoanRequestCountData `json:"data"`
}

type ResponseCustomerLoanStatusDetail struct {
	Status  int                           `json:"status"`
	Message string                        `json:"message"`
	Data    *CustomerLoanStatusDetailData `json:"data"`
}
type ResponseCustomerLoanOfferDetail struct {
	Status  int                  `json:"status"`
	Message string               `json:"message"`
	Data    *LoanOfferDetailData `json:"data"`
}

func (GenerateCode) TableName() string {
	return "t_loan_request"
}

func (LoanRequest) TableName() string {
	return "t_loan_request"
}

func (LoanRequestData) TableName() string {
	return "t_loan_request"
}

func (LoanRequestScheduledData) TableName() string {
	return "t_loan_request"
}

func (LoanRequestDetailData) TableName() string {
	return "t_loan_request"
}

func (LoanRequestAcceptedData) TableName() string {
	return "t_loan_request"
}

func (LoanRequestAcceptedDetailData) TableName() string {
	return "t_loan_request"
}

func (LoanOfferDetailData) TableName() string {
	return "t_loan_request"
}

func (CustomerLoanStatusData) TableName() string {
	return "t_loan_request"
}

func (CustomerLoanStatusDetailData) TableName() string {
	return "t_loan_request"
}

func (LoanRequestCountData) TableName() string {
	return "t_loan_request"
}

func (p *LoanRequest) Prepare() {
	p.MitraID = p.MitraID
	p.CustUserID = p.CustUserID
	p.MitraUserID = p.MitraUserID
	p.MitraBranchID = nil
	p.MitraProductID = nil
	p.ProductSubCategoryID = p.ProductSubCategoryID
	p.LoanRequestAmount = p.LoanRequestAmount
	p.LoanRequestLong = p.LoanRequestLong
	p.LoanRequestPurpose = p.LoanRequestPurpose
	p.LoanRequestSalaryRange = p.LoanRequestSalaryRange
	p.LoanRequestJob = p.LoanRequestJob
	p.LoanRequestReferral = nil
	p.LoanRequestSchdlStatus = p.LoanRequestSchdlStatus
	p.LoanRequestSchdlMom = nil
	p.LoanRequestStatus = p.LoanRequestStatus
	p.LoanRequestCreatedAt = time.Now()
	p.LoanRequestUpdatedAt = p.LoanRequestUpdatedAt
	p.LoanRequestDeletedAt = p.LoanRequestDeletedAt
}

func (p *LoanRequest) Validation() bool {

	// v := validator.New()

	// a := LoanRequest{}

	// err := v.Struct(a)

	// if err != nil {
	// 	for _, eror := range err.(validator.ValidationErrors) {
	// 		fmt.Println(eror)
	// 		var msg string
	// 		// return a.MessageError(e)
	// 		msg = fmt.Sprintf("Message Error : ", eror)
	// 		a.MessageError = msg
	// 		return false
	// 	}
	// }
	return true

}

func (p *LoanRequest) Validate(action string) error {
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

func (p *LoanRequest) SaveLoanRequest(db *gorm.DB) (*LoanRequest, error) {
	var err error
	err = db.Debug().Model(&LoanRequest{}).Create(&p).Error
	if err != nil {
		return &LoanRequest{}, err
	}
	return p, nil
}

func (p *LoanRequest) FindAllLoanRequestByMitra(db *gorm.DB, mitra uint64) (*[]LoanRequestData, error) {
	var err error
	loan_req := []LoanRequestData{}
	err = db.Debug().Model(&LoanRequestData{}).Limit(100).
		Select(`t_loan_request.loan_req_id,
				t_loan_request.cust_user_id,
				m_product_sub_category.prod_sub_ctgry_id, 	
				m_product_sub_category.prod_sub_ctgry_name, 	  	  
				m_customer.cust_name,	  	  
				m_customer.cust_sex,
				m_customer.cust_birth_date,
				m_customer.cust_marital_status,	  	  
				m_customer.cust_age,	     	  
				m_province.province_name as cust_province,	     	  
				m_regency.regency_name as cust_regency,	     
				t_loan_request.mitra_id,
				t_loan_request.loan_req_amount,
				t_loan_request.loan_req_long,
				t_loan_request.loan_req_purpose,
				t_loan_request.loan_req_salary_range,
				t_loan_request.loan_req_date,
				t_loan_request.loan_req_job,
				t_loan_request.loan_req_status,
				t_loan_request.loan_req_created_at at time zone 'Asia/Jakarta' as loan_req_created_at`).
		Joins("JOIN m_customer_user on t_loan_request.cust_user_id=m_customer_user.cust_user_id").
		Joins("JOIN m_customer on m_customer.cust_id=m_customer_user.cust_id").
		Joins("JOIN m_customer_address on m_customer.cust_id=m_customer_address.cust_id").
		Joins("JOIN m_address on m_customer_address.address_id=m_address.address_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_product_sub_category on t_loan_request.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Where("loan_req_status = 1 AND mitra_id = ?", mitra).
		Order("loan_req_created_at desc").
		Find(&loan_req).Error
	if err != nil {
		return &[]LoanRequestData{}, err
	}
	return &loan_req, nil
}

func (p *LoanRequest) FindAllLoanRequestKMGByMitra(db *gorm.DB, mitra uint64) (*[]LoanRequestData, error) {
	var err error
	loan_req := []LoanRequestData{}
	err = db.Debug().Model(&LoanRequestData{}).Limit(100).
		Select(`t_loan_request.loan_req_id,
				t_loan_request.cust_user_id,
				t_loan_request_collateral.loan_req_coll_id,
				t_loan_request_collateral.loan_req_coll_name,
				t_loan_request_collateral.loan_req_coll_desc,
				m_product_sub_category.prod_sub_ctgry_id, 	
				m_product_sub_category.prod_sub_ctgry_name, 	  	  
				m_customer.cust_name,	  	  
				m_customer.cust_sex,	  	  
				m_customer.cust_marital_status,	  	  
				m_customer.cust_age,	     	  
				m_province.province_name as cust_province,	     	  
				m_regency.regency_name as cust_regency,	     
				t_loan_request.mitra_id,
				t_loan_request.loan_req_amount,
				t_loan_request.loan_req_long,
				t_loan_request.loan_req_purpose,
				t_loan_request.loan_req_salary_range,
				t_loan_request.loan_req_date,
				t_loan_request.loan_req_job,
				t_loan_request.loan_req_status,
				t_loan_request.loan_req_created_at at time zone 'Asia/Jakarta' as loan_req_created_at`).
		Joins("JOIN m_customer_user on t_loan_request.cust_user_id=m_customer_user.cust_user_id").
		Joins("JOIN m_customer on m_customer.cust_id=m_customer_user.cust_id").
		Joins("JOIN m_customer_address on m_customer.cust_id=m_customer_address.cust_id").
		Joins("JOIN m_address on m_customer_address.address_id=m_address.address_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_product_sub_category on t_loan_request.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Joins("JOIN t_loan_request_collateral on t_loan_request.loan_req_id = t_loan_request_collateral.loan_req_id").
		Where("loan_req_status = 1 AND mitra_id = ?", mitra).
		Order("loan_req_created_at desc").
		Find(&loan_req).Error
	if err != nil {
		return &[]LoanRequestData{}, err
	}
	return &loan_req, nil
}

func (p *LoanRequest) FindAllLoanRequestBySchedulling(db *gorm.DB, phone string) (*[]LoanRequestScheduledData, error) {
	var err error
	loan_req := []LoanRequestScheduledData{}
	err = db.Debug().Model(&LoanRequestScheduledData{}).Limit(100).
		Select(`t_loan_request.loan_req_id,
				t_loan_request.cust_user_id,
				t_loan_request_scheduling.loan_req_schdl_date_1,
				t_loan_request_scheduling.loan_req_schdl_place,
				t_loan_request_scheduling.loan_req_schdl_time,
				m_product_sub_category.prod_sub_ctgry_id, 	
				m_product_sub_category.prod_sub_ctgry_name, 	  	  
				m_customer.cust_name,	  	  
				m_customer.cust_sex,	  	  
				m_customer.cust_marital_status,	  	  
				m_customer.cust_age,	     	  
				m_province.province_name as cust_province,	     	  
				m_regency.regency_name as cust_regency,	     
				t_loan_request.mitra_id,
				t_loan_request.loan_req_amount,
				t_loan_request.loan_req_long,
				t_loan_request.loan_req_date,
				t_loan_request.loan_req_job,
				t_loan_request.loan_req_status,
				t_loan_request.loan_req_created_at at time zone 'Asia/Jakarta' as loan_req_created_at`).
		Joins("JOIN t_loan_request_scheduling on t_loan_request.loan_req_id=t_loan_request_scheduling.loan_req_id").
		Joins("JOIN m_mitra_user on t_loan_request.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_customer_user on t_loan_request.cust_user_id=m_customer_user.cust_user_id").
		Joins("JOIN m_customer on m_customer.cust_id=m_customer_user.cust_id").
		Joins("JOIN m_customer_address on m_customer.cust_id=m_customer_address.cust_id").
		Joins("JOIN m_address on m_customer_address.address_id=m_address.address_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_product_sub_category on t_loan_request.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Where("t_loan_request.loan_req_status = 5 AND t_loan_request.loan_req_schdl_status = 3 AND t_loan_request_scheduling.loan_req_schdl_status = 1 AND m_mitra_user.mitra_user_phone = ?", phone).
		Order("loan_req_created_at desc").
		Find(&loan_req).Error
	if err != nil {
		return &[]LoanRequestScheduledData{}, err
	}
	return &loan_req, nil
}

func (p *LoanRequest) FindAllLoanRequestAcceptedByMitraUserID(db *gorm.DB, uid uint64) (*[]LoanRequestAcceptedData, error) {
	var err error
	loan_req := []LoanRequestAcceptedData{}
	err = db.Debug().Model(&LoanRequestAcceptedData{}).Limit(100).
		Select(`t_loan_request.loan_req_id,
				t_loan_request.cust_user_id,
				m_product_sub_category.prod_sub_ctgry_id, 	
				m_product_sub_category.prod_sub_ctgry_name, 	
				m_mitra_user.mitra_user_id,	   	  
				m_customer.cust_name,	  
				m_country.country_id,   
				m_country.country_name,	 
				m_province.province_id,   
				m_province.province_name,	 
				m_regency.regency_id,   
				m_regency.regency_name,	    
				m_mitra.mitra_id, 
				t_loan_request.loan_req_amount,
				t_loan_request.loan_req_long,
				t_loan_request.loan_req_purpose,
				t_loan_request.loan_req_salary_range,
				t_loan_request.loan_req_date,
				t_loan_request.loan_req_job,
				t_loan_request.loan_req_status,
				t_loan_request.loan_req_created_at at time zone 'Asia/Jakarta' as loan_req_created_at`).
		Joins("JOIN m_product_sub_category on t_loan_request.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Joins("JOIN m_customer_user on t_loan_request.cust_user_id=m_customer_user.cust_user_id").
		Joins("JOIN m_customer on m_customer.cust_id=m_customer_user.cust_id").
		Joins("JOIN m_customer_address on m_customer.cust_id=m_customer_address.cust_id").
		Joins("JOIN m_address on m_customer_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_mitra_user on t_loan_request.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra on t_loan_request.mitra_id=m_mitra.mitra_id").
		Where("loan_req_status != 1 AND m_mitra_user.mitra_user_id = ?", uid).
		Order("loan_req_created_at desc").
		Find(&loan_req).Error
	if err != nil {
		return &[]LoanRequestAcceptedData{}, err
	}
	return &loan_req, nil
}

func (p *LoanRequest) FindAllLoanOfferedByAO(db *gorm.DB, ao uint64) (*[]LoanRequestAcceptedData, error) {
	var err error
	loan_req := []LoanRequestAcceptedData{}
	err = db.Debug().Model(&LoanRequestAcceptedData{}).Limit(100).
		Select(`t_loan_request.loan_req_id,
				t_loan_request.cust_user_id,
				m_mitra_user.mitra_user_id,	   	  
				m_customer.cust_name,	    
				m_mitra.mitra_id, 
				t_loan_request.loan_req_amount,
				t_loan_request.loan_req_long,
				t_loan_request.loan_req_purpose,
				t_loan_request.loan_req_salary_range,
				t_loan_request.loan_req_date,
				t_loan_request.loan_req_job,
				t_loan_request.loan_req_status,
				t_loan_request.loan_req_schdl_status,
				t_loan_request.loan_req_created_at at time zone 'Asia/Jakarta' as loan_req_created_at`).
		Joins("LEFT JOIN m_customer_user on t_loan_request.cust_user_id=m_customer_user.cust_user_id").
		Joins("LEFT JOIN m_customer on m_customer.cust_id=m_customer_user.cust_id").
		Joins("LEFT JOIN m_mitra_user on t_loan_request.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("LEFT JOIN m_mitra on t_loan_request.mitra_id=m_mitra.mitra_id").
		Where("loan_req_status = 3 AND t_loan_request.mitra_user_id = ?", ao).
		Find(&loan_req).Error
	if err != nil {
		return &[]LoanRequestAcceptedData{}, err
	}
	return &loan_req, nil
}

func (p *LoanRequest) FindAllCustomerLoanStatusByCustomerID(db *gorm.DB, uid uint64) (*[]CustomerLoanStatusData, error) {
	var err error
	loan_req := []CustomerLoanStatusData{}
	err = db.Debug().Model(&CustomerLoanStatusData{}).Limit(100).
		Select(`t_loan_request.loan_req_id,
				t_loan_request.cust_user_id,
				m_customer_user.cust_id,
				m_product_sub_category.prod_sub_ctgry_id, 	
				m_product_sub_category.prod_sub_ctgry_name, 	  
				m_mitra.mitra_id, 	      
				m_mitra.mitra_name,	      
				m_mitra.mitra_logo, 
				m_mitra_user.mitra_user_id,
				m_mitra_user.mitra_user_phone,
				t_loan_request.loan_req_amount,
				t_loan_request.loan_req_long,
				t_loan_request.loan_req_status,
				t_loan_request.loan_req_schdl_status,
				t_loan_request.loan_req_date,
				t_loan_request.loan_req_job,
				t_loan_request.loan_req_created_at at time zone 'Asia/Jakarta' as loan_req_created_at,
				t_loan_request.loan_req_updated_at at time zone 'Asia/Jakarta' as loan_req_updated_at,
				t_loan_request.loan_req_deleted_at at time zone 'Asia/Jakarta' as loan_req_deleted_at
				`).
		Joins("LEFT JOIN m_customer_user on t_loan_request.cust_user_id=m_customer_user.cust_user_id").
		Joins("LEFT JOIN m_customer on m_customer_user.cust_id=m_customer.cust_id").
		Joins("LEFT JOIN m_mitra on t_loan_request.mitra_id=m_mitra.mitra_id").
		Joins("LEFT JOIN m_mitra_user on t_loan_request.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("LEFT JOIN m_product_sub_category on t_loan_request.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Where("m_customer_user.cust_id = ?", uid).
		Order("loan_req_created_at desc").
		Find(&loan_req).Error
	if err != nil {
		return &[]CustomerLoanStatusData{}, err
	}
	return &loan_req, nil
}

func (p *LoanRequest) FindAllLoanRequest(db *gorm.DB) (*[]LoanRequestData, error) {
	var err error
	loan_req := []LoanRequestData{}
	err = db.Debug().Model(&LoanRequestData{}).Limit(100).
		Select(`t_loan_request.loan_req_id,
				m_customer_user.cust_user_id,	     
				m_mitra_user.mitra_user_id,
				m_mitra.mitra_id,
				m_mitra_branch.mitra_branch_id,
				m_mitra_product.mitra_prod_id,
				t_loan_request.loan_req_amount,
				t_loan_request.loan_req_long,
				t_loan_request.loan_req_purpose,
				t_loan_request.loan_req_salary_range,
				t_loan_request.loan_req_referral,
				t_loan_request.loan_req_schdl_status,
				t_loan_request.loan_req_schdl_mom,
				t_loan_request.loan_req_date,
				t_loan_request.loan_req_job,
				t_loan_request.loan_req_created_at at time zone 'Asia/Jakarta' as loan_req_created_at,
				t_loan_request.loan_req_updated_at at time zone 'Asia/Jakarta' as loan_req_updated_at,
				t_loan_request.loan_req_deleted_at at time zone 'Asia/Jakarta' as loan_req_deleted_at`).
		Joins("JOIN m_customer_user on t_loan_request.cust_user_id=m_customer_user.cust_user_id").
		Joins("JOIN m_mitra_user on t_loan_request.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra on t_loan_request.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on t_loan_request.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra_product on t_loan_request.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Find(&loan_req).Error
	if err != nil {
		return &[]LoanRequestData{}, err
	}
	return &loan_req, nil
}

func (p *LoanRequest) FindAllLoanReqSchdlStatus(db *gorm.DB, status uint64) (*[]LoanRequestData, error) {
	var err error
	loan_req := []LoanRequestData{}
	err = db.Debug().Model(&LoanRequestData{}).Limit(100).
		Select(`t_loan_request.loan_req_id,
				t_loan_request.cust_user_id,	     
				t_loan_request.mitra_user_id,
				t_loan_request.mitra_id,
				t_loan_request.mitra_branch_id,
				t_loan_request.mitra_prod_id,
				t_loan_request.loan_req_amount,
				t_loan_request.loan_req_long,
				t_loan_request.loan_req_purpose,
				t_loan_request.loan_req_salary_range,
				t_loan_request.loan_req_referral,
				t_loan_request.loan_req_schdl_status,
				t_loan_request.loan_req_schdl_mom,
				t_loan_request.loan_req_date,
				t_loan_request.loan_req_job,
				t_loan_request.loan_req_created_at at time zone 'Asia/Jakarta' as loan_req_created_at,
				t_loan_request.loan_req_updated_at at time zone 'Asia/Jakarta' as loan_req_updated_at,
				t_loan_request.loan_req_deleted_at at time zone 'Asia/Jakarta' as loan_req_deleted_at`).
		Joins("JOIN m_customer_user on t_loan_request.cust_user_id=m_customer_user.cust_user_id").
		Joins("JOIN m_mitra_user on t_loan_request.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra on t_loan_request.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on t_loan_request.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra_product on t_loan_request.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Where("loan_req_schdl_status = ?", status).
		Find(&loan_req).Error
	if err != nil {
		return &[]LoanRequestData{}, err
	}
	return &loan_req, nil
}

func (p *LoanRequest) FindLoanRequestDataByID(db *gorm.DB, pid uint64) (*LoanRequest, error) {
	var err error
	err = db.Debug().Model(&LoanRequest{}).Where("loan_req_id = ?", pid).Take(&p).Error
	if err != nil {
		return &LoanRequest{}, err
	}
	return p, nil
}

func (p *LoanRequest) FindLoanRequestActByID(db *gorm.DB, pid uint64) (*LoanRequest, error) {
	var err error
	err = db.Debug().Model(&LoanRequest{}).Where("loan_req_id = ?", pid).Take(&p).Error
	if err != nil {
		return &LoanRequest{}, err
	}
	return p, nil
}

func (p *LoanRequest) LoanRequestDetailByID(db *gorm.DB, pid uint64) (*LoanRequestDetailData, error) {
	var err error
	loan_req := LoanRequestDetailData{}
	err = db.Debug().Model(&LoanRequestDetailData{}).Limit(100).
		Select(`t_loan_request.loan_req_id,
				t_loan_request.cust_user_id,
				t_loan_request.prod_sub_ctgry_id`).
		Where("loan_req_id = ?", pid).
		Find(&loan_req).Error
	if err != nil {
		return &LoanRequestDetailData{}, err
	}
	return &loan_req, nil
}

func (p *LoanRequest) LoanRequestDetailByMitraID(db *gorm.DB, pid uint64) (*LoanRequestDetailData, error) {
	var err error
	loan_req := LoanRequestDetailData{}
	err = db.Debug().Model(&LoanRequestDetailData{}).Limit(100).
		Select(`t_loan_request.loan_req_id,
				t_loan_request.cust_user_id,
				t_loan_request.prod_sub_ctgry_id`).
		Where("mitra_id = ?", pid).
		Find(&loan_req).Error
	if err != nil {
		return &LoanRequestDetailData{}, err
	}
	return &loan_req, nil
}

func (p *LoanRequest) FindLoanRequestDetailByID(db *gorm.DB, pid uint64) (*LoanRequestDetailData, error) {
	var err error
	loan_req := LoanRequestDetailData{}
	err = db.Debug().Model(&LoanRequestDetailData{}).Limit(100).
		Select(`t_loan_request.loan_req_id,
				t_loan_request.cust_user_id,
				m_mitra_user.mitra_user_id,
				m_mitra_user.mitra_user_phone,					  
				m_product_sub_category.prod_sub_ctgry_id, 	
				m_product_sub_category.prod_sub_ctgry_name, 	  
				m_customer.cust_name,	     
				m_customer.cust_sex,	     
				m_customer.cust_birth_date,   
				m_customer.cust_marital_status,   
				m_customer_user.cust_user_phone,
				m_customer_address.cust_adrs_street,
				m_customer_address.cust_adrs_rt,
				m_customer_address.cust_adrs_rw,
				m_customer_address.cust_adrs_zip_code,   
				m_country.country_id,   
				m_country.country_name,	 
				m_province.province_id,   
				m_province.province_name,	 
				m_regency.regency_id,   
				m_regency.regency_name,	     
				t_loan_request.mitra_id,
				t_loan_request.loan_req_amount,
				t_loan_request.loan_req_long,
				t_loan_request.loan_req_purpose,
				t_loan_request.loan_req_salary_range,
				t_loan_request.loan_req_date,
				t_loan_request.loan_req_job,
				t_loan_request.loan_req_status,
				t_loan_request.loan_req_schdl_status,
				t_loan_request.loan_req_created_at at time zone 'Asia/Jakarta' as loan_req_created_at`).
		Joins("LEFT JOIN m_product_sub_category on t_loan_request.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Joins("LEFT JOIN m_customer_user on t_loan_request.cust_user_id=m_customer_user.cust_user_id").
		Joins("LEFT JOIN m_customer on m_customer.cust_id=m_customer_user.cust_id").
		Joins("LEFT JOIN m_customer_address on m_customer.cust_id=m_customer_address.cust_id").
		Joins("LEFT JOIN m_mitra_user on t_loan_request.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("LEFT JOIN m_address on m_customer_address.address_id=m_address.address_id").
		Joins("LEFT JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("LEFT JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("LEFT JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Where("loan_req_id = ?", pid).
		Find(&loan_req).Error
	if err != nil {
		return &LoanRequestDetailData{}, err
	}
	return &loan_req, nil
}

func (p *LoanRequest) FindLoanRequestDetailByIDMitraByPhone(db *gorm.DB, mid uint64, phone string) (*[]LoanRequestDetailData, error) {
	var err error
	loan_req := []LoanRequestDetailData{}
	err = db.Debug().Model(&LoanRequestDetailData{}).Limit(100).
		Select(`t_loan_request.loan_req_id,
				t_loan_request.cust_user_id,
				m_product_sub_category.prod_sub_ctgry_id, 	
				m_product_sub_category.prod_sub_ctgry_name, 	  
				m_customer.cust_name,	     
				m_customer.cust_sex,	     
				m_customer.cust_birth_date,   
				m_customer.cust_marital_status,   
				m_customer_user.cust_user_phone,
				m_customer_address.cust_adrs_street,
				m_customer_address.cust_adrs_rt,
				m_customer_address.cust_adrs_rw,
				m_customer_address.cust_adrs_zip_code,   
				m_country.country_id,   
				m_country.country_name,	 
				m_province.province_id,   
				m_province.province_name,	 
				m_regency.regency_id,   
				m_regency.regency_name,	     
				t_loan_request.mitra_id,
				t_loan_request.loan_req_amount,
				t_loan_request.loan_req_long,
				t_loan_request.loan_req_purpose,
				t_loan_request.loan_req_salary_range,
				t_loan_request.loan_req_date,
				t_loan_request.loan_req_job,
				t_loan_request.loan_req_status,
				t_loan_request.loan_req_schdl_status,
				t_loan_request.loan_req_created_at at time zone 'Asia/Jakarta' as loan_req_created_at`).
		Joins("JOIN m_product_sub_category on t_loan_request.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Joins("JOIN m_customer_user on t_loan_request.cust_user_id=m_customer_user.cust_user_id").
		Joins("JOIN m_customer on m_customer.cust_id=m_customer_user.cust_id").
		Joins("JOIN m_customer_address on m_customer.cust_id=m_customer_address.cust_id").
		Joins("JOIN m_address on m_customer_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Where("mitra_id = ? AND cust_user_phone = ?", mid, phone).
		Find(&loan_req).Error
	if err != nil {
		return &[]LoanRequestDetailData{}, err
	}
	return &loan_req, nil
}

func (p *LoanRequest) FindLoanRequestAcceptedDetailByID(db *gorm.DB, pid uint64) (*LoanRequestAcceptedDetailData, error) {
	var err error
	loan_req := LoanRequestAcceptedDetailData{}
	err = db.Debug().Model(&LoanRequestAcceptedDetailData{}).Limit(100).
		Select(`t_loan_request.loan_req_id,
				t_loan_request.cust_user_id,
				m_product_sub_category.prod_sub_ctgry_id, 	
				m_product_sub_category.prod_sub_ctgry_name, 
				m_mitra_user.mitra_user_id,
				m_mitra_user.mitra_user_phone,
				m_mitra_employee.mitra_emp_name,	  	  
				m_customer.cust_name,	     
				m_customer.cust_sex,	     
				m_customer.cust_birth_date,
				m_customer.cust_marital_status,   
				m_customer_user.cust_user_phone,
				m_customer_address.cust_adrs_street,
				m_customer_address.cust_adrs_rt,
				m_customer_address.cust_adrs_rw,
				m_customer_address.cust_adrs_zip_code,   	     
				m_country.country_id,   
				m_country.country_name,	 
				m_province.province_id,   
				m_province.province_name,	 
				m_regency.regency_id,   
				m_regency.regency_name,	     
				m_mitra.mitra_id, 
				m_mitra.mitra_name,
				t_loan_request.loan_req_amount,
				t_loan_request.loan_req_long,
				t_loan_request.loan_req_purpose,
				t_loan_request.loan_req_salary_range,
				t_loan_request.loan_req_date,
				t_loan_request.loan_req_job,
				t_loan_request.loan_req_status,
				t_loan_request.loan_req_schdl_status,
				t_loan_request.loan_req_created_at at time zone 'Asia/Jakarta' as loan_req_created_at`).
		Joins("JOIN m_product_sub_category on t_loan_request.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Joins("JOIN m_customer_user on t_loan_request.cust_user_id=m_customer_user.cust_user_id").
		Joins("JOIN m_customer on m_customer.cust_id=m_customer_user.cust_id").
		Joins("JOIN m_customer_address on m_customer.cust_id=m_customer_address.cust_id").
		Joins("JOIN m_address on m_customer_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_mitra_user on t_loan_request.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra_employee on m_mitra_user.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra on t_loan_request.mitra_id=m_mitra.mitra_id").
		Where("loan_req_id = ?", pid).
		Find(&loan_req).Error
	if err != nil {
		return &LoanRequestAcceptedDetailData{}, err
	}
	return &loan_req, nil
}

func (p *LoanRequest) FindLoanRequestKMGAcceptedDetailByID(db *gorm.DB, pid uint64) (*LoanRequestAcceptedDetailData, error) {
	var err error
	loan_req := LoanRequestAcceptedDetailData{}
	err = db.Debug().Model(&LoanRequestAcceptedDetailData{}).Limit(100).
		Select(`t_loan_request.loan_req_id,
				t_loan_request.cust_user_id,
				m_product_sub_category.prod_sub_ctgry_id, 	
				m_product_sub_category.prod_sub_ctgry_name, 
				m_mitra_user.mitra_user_id,
				m_mitra_user.mitra_user_phone,
				m_mitra_employee.mitra_emp_name,	  	  
				m_customer.cust_name,	     
				m_customer.cust_sex,	     
				m_customer.cust_birth_date,
				m_customer.cust_marital_status,   
				m_customer_user.cust_user_phone,
				m_customer_address.cust_adrs_street,
				m_customer_address.cust_adrs_rt,
				m_customer_address.cust_adrs_rw,
				m_customer_address.cust_adrs_zip_code,   	     
				m_country.country_id,   
				m_country.country_name,	 
				m_province.province_id,   
				m_province.province_name,	 
				m_regency.regency_id,   
				m_regency.regency_name,	     
				m_mitra.mitra_id, 
				m_mitra.mitra_name,
				t_loan_request.loan_req_amount,
				t_loan_request.loan_req_long,
				t_loan_request.loan_req_purpose,
				t_loan_request.loan_req_salary_range,
				t_loan_request.loan_req_date,
				t_loan_request.loan_req_job,
				t_loan_request.loan_req_status,
				t_loan_request.loan_req_schdl_status,
				t_loan_request.loan_req_created_at at time zone 'Asia/Jakarta' as loan_req_created_at,
				t_loan_request_collateral.loan_req_coll_id,
				t_loan_request_collateral.loan_req_coll_name,
				t_loan_request_collateral.loan_req_coll_desc,
				t_loan_request_collateral_img.loan_req_coll_img_id,
				t_loan_request_collateral_img.loan_req_coll_img_value`).
		Joins("JOIN m_product_sub_category on t_loan_request.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Joins("JOIN m_customer_user on t_loan_request.cust_user_id=m_customer_user.cust_user_id").
		Joins("JOIN m_customer on m_customer.cust_id=m_customer_user.cust_id").
		Joins("JOIN m_customer_address on m_customer.cust_id=m_customer_address.cust_id").
		Joins("JOIN m_address on m_customer_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_mitra_user on t_loan_request.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra_employee on m_mitra_user.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra on t_loan_request.mitra_id=m_mitra.mitra_id").
		Joins("JOIN t_loan_request_collateral on t_loan_request.loan_req_id = t_loan_request_collateral.loan_req_id").
		Joins("JOIN t_loan_request_collateral_img on t_loan_request_collateral.loan_req_coll_id = t_loan_request_collateral_img.loan_req_coll_id").
		Where("t_loan_request.loan_req_id = ?", pid).
		Find(&loan_req).Error
	if err != nil {
		return &LoanRequestAcceptedDetailData{}, err
	}
	return &loan_req, nil
}

func (p *LoanRequest) FindLoanRequestKMGDetailByID(db *gorm.DB, pid uint64) (*LoanRequestDetailData, error) {
	var err error
	loan_req := LoanRequestDetailData{}
	err = db.Debug().Model(&LoanRequestDetailData{}).Limit(100).
		Select(`t_loan_request.loan_req_id,
				t_loan_request.cust_user_id,
				m_product_sub_category.prod_sub_ctgry_id, 	
				m_product_sub_category.prod_sub_ctgry_name, 	  
				m_mitra_user.mitra_user_id,
				m_mitra_user.mitra_user_phone,					  
				m_customer.cust_name,	     
				m_customer.cust_sex,	     
				m_customer.cust_birth_date,   
				m_customer.cust_marital_status,   
				m_customer_user.cust_user_phone,
				m_customer_address.cust_adrs_street,
				m_customer_address.cust_adrs_rt,
				m_customer_address.cust_adrs_rw,
				m_customer_address.cust_adrs_zip_code,   
				m_country.country_id,   
				m_country.country_name,	 
				m_province.province_id,   
				m_province.province_name,	 
				m_regency.regency_id,   
				m_regency.regency_name,	     
				t_loan_request.mitra_id,
				t_loan_request.loan_req_amount,
				t_loan_request.loan_req_long,
				t_loan_request.loan_req_purpose,
				t_loan_request.loan_req_salary_range,
				t_loan_request.loan_req_date,
				t_loan_request.loan_req_job,
				t_loan_request.loan_req_status,
				t_loan_request.loan_req_schdl_status,
				t_loan_request.loan_req_created_at at time zone 'Asia/Jakarta' as loan_req_created_at,
				t_loan_request_collateral.loan_req_coll_id,
				t_loan_request_collateral.loan_req_coll_name,
				t_loan_request_collateral.loan_req_coll_desc,
				t_loan_request_collateral_img.loan_req_coll_img_id,
				t_loan_request_collateral_img.loan_req_coll_img_value`).
		Joins("LEFT JOIN m_product_sub_category on t_loan_request.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Joins("LEFT JOIN m_customer_user on t_loan_request.cust_user_id=m_customer_user.cust_user_id").
		Joins("LEFT JOIN m_customer on m_customer.cust_id=m_customer_user.cust_id").
		Joins("LEFT JOIN m_customer_address on m_customer.cust_id=m_customer_address.cust_id").
		Joins("LEFT JOIN m_mitra_user on t_loan_request.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("LEFT JOIN m_address on m_customer_address.address_id=m_address.address_id").
		Joins("LEFT JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("LEFT JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("LEFT JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("LEFT JOIN t_loan_request_collateral on t_loan_request.loan_req_id = t_loan_request_collateral.loan_req_id").
		Joins("LEFT JOIN t_loan_request_collateral_img on t_loan_request_collateral.loan_req_coll_id = t_loan_request_collateral_img.loan_req_coll_id").
		Where("t_loan_request.loan_req_id = ?", pid).
		Find(&loan_req).Error
	if err != nil {
		return &LoanRequestDetailData{}, err
	}
	return &loan_req, nil
}

func (p *LoanRequest) FindAllLoanRequestKMGByID(db *gorm.DB, pid uint64) (*LoanRequestDetailData, error) {
	var err error
	loan_req := LoanRequestDetailData{}
	err = db.Debug().Model(&LoanRequestDetailData{}).Limit(100).
		Select(`t_loan_request.loan_req_id,
				t_loan_request.cust_user_id,
				m_product_sub_category.prod_sub_ctgry_id, 	
				m_product_sub_category.prod_sub_ctgry_name, 	  
				m_customer.cust_name,	     
				m_customer.cust_sex,	     
				m_customer.cust_birth_date,   
				m_customer.cust_marital_status,   
				m_customer_user.cust_user_phone,
				m_customer_address.cust_adrs_street,
				m_customer_address.cust_adrs_rt,
				m_customer_address.cust_adrs_rw,
				m_customer_address.cust_adrs_zip_code,   
				m_country.country_id,   
				m_country.country_name,	 
				m_province.province_id,   
				m_province.province_name,	 
				m_regency.regency_id,   
				m_regency.regency_name,	     
				t_loan_request.mitra_id,
				t_loan_request.loan_req_num,
				t_loan_request.loan_req_amount,
				t_loan_request.loan_req_long,
				t_loan_request.loan_req_purpose,
				t_loan_request.loan_req_salary_range,
				t_loan_request.loan_req_date,
				t_loan_request.loan_req_job,
				t_loan_request.loan_req_status,
				t_loan_request.loan_req_schdl_status,
				t_loan_request.loan_req_created_at at time zone 'Asia/Jakarta' as loan_req_created_at,
				t_loan_request_collateral.loan_req_coll_id,
				t_loan_request_collateral.loan_req_coll_name,
				t_loan_request_collateral.loan_req_coll_desc`).
		Joins("JOIN m_product_sub_category on t_loan_request.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Joins("JOIN m_customer_user on t_loan_request.cust_user_id=m_customer_user.cust_user_id").
		Joins("JOIN m_customer on m_customer.cust_id=m_customer_user.cust_id").
		Joins("JOIN m_customer_address on m_customer.cust_id=m_customer_address.cust_id").
		Joins("JOIN m_address on m_customer_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN t_loan_request_collateral on t_loan_request.loan_req_id = t_loan_request_collateral.loan_req_id").
		Where("t_loan_request.loan_req_id = ?", pid).
		Find(&loan_req).Error
	if err != nil {
		return &LoanRequestDetailData{}, err
	}
	return &loan_req, nil
}
func (p *LoanRequest) FindCustomerLoanStatusDetailByID(db *gorm.DB, pid uint64) (*CustomerLoanStatusDetailData, error) {
	var err error
	loan_req := CustomerLoanStatusDetailData{}
	err = db.Debug().Model(&CustomerLoanStatusDetailData{}).Limit(100).
		Select(`t_loan_request.loan_req_id,
				t_loan_request.cust_user_id,  
				m_product_sub_category.prod_sub_ctgry_id, 	
				m_product_sub_category.prod_sub_ctgry_name, 	  
				m_customer.cust_name,	     
				m_customer.cust_sex,
				m_customer.cust_marital_status,
				m_customer_address.cust_adrs_street,	      
				m_country.country_id,   
				m_country.country_name,	 
				m_province.province_id,   
				m_province.province_name,	 
				m_regency.regency_id,   
				m_regency.regency_name,	
				m_mitra.mitra_id, 
				m_mitra.mitra_name,
				m_mitra.mitra_logo,
				t_loan_request.mitra_user_id,	  	  
				t_loan_request.loan_req_amount,
				t_loan_request.loan_req_salary_range,
				t_loan_request.loan_req_date,
				t_loan_request.loan_req_long,
				t_loan_request.loan_req_purpose,
				t_loan_request.loan_req_job,
				t_loan_request.loan_req_status,
				t_loan_request.loan_req_schdl_status,
				t_loan_request.loan_req_created_at at time zone 'Asia/Jakarta' as loan_req_created_at`).
		Joins("JOIN m_customer_user on t_loan_request.cust_user_id=m_customer_user.cust_user_id").
		Joins("JOIN m_customer on m_customer.cust_id=m_customer_user.cust_id").
		Joins("JOIN m_mitra on t_loan_request.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_product_sub_category on t_loan_request.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Joins("JOIN m_customer_address on m_customer.cust_id=m_customer_address.cust_id").
		Joins("JOIN m_address on m_customer_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Where("loan_req_id = ?", pid).
		Find(&loan_req).Error
	if err != nil {
		return &CustomerLoanStatusDetailData{}, err
	}
	return &loan_req, nil
}

func (p *LoanRequest) FindCustomerLoanKMGStatusDetailByID(db *gorm.DB, pid uint64) (*CustomerLoanStatusDetailData, error) {
	var err error
	loan_req := CustomerLoanStatusDetailData{}
	err = db.Debug().Model(&CustomerLoanStatusDetailData{}).Limit(100).
		Select(`t_loan_request.loan_req_id,
				t_loan_request.cust_user_id,  
				m_product_sub_category.prod_sub_ctgry_id, 	
				m_product_sub_category.prod_sub_ctgry_name, 	  
				m_customer.cust_name,	     
				m_customer.cust_sex,
				m_customer.cust_marital_status,
				m_customer_address.cust_adrs_street,	      
				m_country.country_id,   
				m_country.country_name,	 
				m_province.province_id,   
				m_province.province_name,	 
				m_regency.regency_id,   
				m_regency.regency_name,	
				m_mitra.mitra_id, 
				m_mitra.mitra_name,
				m_mitra.mitra_logo,
				m_mitra_user.mitra_user_phone,					  
				t_loan_request.mitra_user_id,	  	  
				t_loan_request.loan_req_amount,
				t_loan_request.loan_req_salary_range,
				t_loan_request.loan_req_date,
				t_loan_request.loan_req_long,
				t_loan_request.loan_req_purpose,
				t_loan_request.loan_req_job,
				t_loan_request.loan_req_status,
				t_loan_request.loan_req_schdl_status,
				t_loan_request.loan_req_created_at at time zone 'Asia/Jakarta' as loan_req_created_at,
				t_loan_request_collateral.loan_req_coll_id,
				t_loan_request_collateral.loan_req_coll_name,
				t_loan_request_collateral.loan_req_coll_desc,
				t_loan_request_collateral_img.loan_req_coll_img_id,
				t_loan_request_collateral_img.loan_req_coll_img_value`).
		Joins("LEFT JOIN m_customer_user on t_loan_request.cust_user_id=m_customer_user.cust_user_id").
		Joins("LEFT JOIN m_customer on m_customer.cust_id=m_customer_user.cust_id").
		Joins("LEFT JOIN m_mitra on t_loan_request.mitra_id=m_mitra.mitra_id").
		Joins("LEFT JOIN m_mitra_user on t_loan_request.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("LEFT JOIN m_product_sub_category on t_loan_request.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Joins("LEFT JOIN m_customer_address on m_customer.cust_id=m_customer_address.cust_id").
		Joins("LEFT JOIN m_address on m_customer_address.address_id=m_address.address_id").
		Joins("LEFT JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("LEFT JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("LEFT JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("LEFT JOIN t_loan_request_collateral on t_loan_request.loan_req_id = t_loan_request_collateral.loan_req_id").
		Joins("LEFT JOIN t_loan_request_collateral_img on t_loan_request_collateral.loan_req_coll_id = t_loan_request_collateral_img.loan_req_coll_id").
		Where("t_loan_request.loan_req_id = ?", pid).
		Find(&loan_req).Error
	if err != nil {
		return &CustomerLoanStatusDetailData{}, err
	}
	return &loan_req, nil
}

func (p *LoanRequest) FindCustomerLoanOfferDetailByID(db *gorm.DB, pid uint64) (*LoanOfferDetailData, error) {
	var err error
	loan_req := LoanOfferDetailData{}
	err = db.Debug().Model(&LoanOfferDetailData{}).Limit(100).
		Select(`t_loan_request.loan_req_id,
				m_product_sub_category.prod_sub_ctgry_id, 	
				m_product_sub_category.prod_sub_ctgry_name,
				t_loan_request.cust_user_id,   	  
				m_customer.cust_name,	 	  
				m_customer.cust_sex,	 	  
				m_customer.cust_birth_place,	 	  
				m_customer_user.cust_user_phone,	      
				m_country.country_id,   
				m_country.country_name,	 
				m_province.province_id,   
				m_province.province_name,	 
				m_regency.regency_id,   
				m_regency.regency_name,	     
				m_mitra.mitra_id, 
				m_mitra.mitra_name,
				m_mitra.mitra_logo,
				m_mitra_user.mitra_user_id,
				m_mitra_employee.mitra_emp_name, 	  
				m_mitra_product.mitra_prod_id, 	  
				m_mitra_product.mitra_prod_name, 	  
				m_mitra_product.mitra_prod_interest_rate, 	  
				m_mitra_product.mitra_prod_interest_rate_period, 	  
				m_mitra_product.mitra_prod_interest_rate_type, 
				t_loan_request.mitra_branch_id,
				t_loan_request.loan_req_num,
				t_loan_request.loan_req_amount,
				t_loan_request.loan_req_long,
				t_loan_request.loan_req_purpose,
				t_loan_request.loan_req_salary_range,
				t_loan_request.loan_req_date,
				t_loan_request.loan_req_job,
				t_loan_request.loan_req_status,
				t_loan_request.loan_req_schdl_status,
				t_loan_request.loan_req_created_at at time zone 'Asia/Jakarta' as loan_req_created_at`).
		Joins("JOIN m_product_sub_category on t_loan_request.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Joins("JOIN m_customer_user on t_loan_request.cust_user_id=m_customer_user.cust_user_id").
		Joins("JOIN m_customer on m_customer.cust_id=m_customer_user.cust_id").
		Joins("JOIN m_customer_address on m_customer.cust_id=m_customer_address.cust_id").
		Joins("JOIN m_address on m_customer_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_mitra on t_loan_request.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_user on t_loan_request.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra_employee on m_mitra_user.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra_product on t_loan_request.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Where("loan_req_id = ?", pid).
		Find(&loan_req).Error
	if err != nil {
		return &LoanOfferDetailData{}, err
	}
	return &loan_req, nil
}

func (p *LoanRequest) FindLoanReqSchdlStatusByID(db *gorm.DB, pid uint64, status uint64) (*LoanRequestData, error) {
	var err error
	loan_req := LoanRequestData{}
	err = db.Debug().Model(&LoanRequestData{}).Limit(100).
		Select(`t_loan_request.loan_req_id,
				t_loan_request.cust_user_id,	     
				t_loan_request.mitra_user_id,
				t_loan_request.mitra_id,
				t_loan_request.mitra_branch_id,
				t_loan_request.mitra_prod_id,
				t_loan_request.loan_req_amount,
				t_loan_request.loan_req_long,
				t_loan_request.loan_req_purpose,
				t_loan_request.loan_req_salary_range,
				t_loan_request.loan_req_referral,
				t_loan_request.loan_req_date,
				t_loan_request.loan_req_schdl_status,
				t_loan_request.loan_req_schdl_mom,
				t_loan_request.loan_req_created_at at time zone 'Asia/Jakarta' as loan_req_created_at,
				t_loan_request.loan_req_updated_at at time zone 'Asia/Jakarta' as loan_req_updated_at,
				t_loan_request.loan_req_deleted_at at time zone 'Asia/Jakarta' as loan_req_deleted_at`).
		Where("loan_req_id = ? AND loan_req_schdl_status = ?", pid, status).
		Find(&loan_req).Error
	if err != nil {
		return &LoanRequestData{}, err
	}
	return &loan_req, nil
}

func (p *LoanRequest) UpdateLoanRequest(db *gorm.DB, pid uint64) (*LoanRequest, error) {

	var err error
	db = db.Debug().Model(&LoanRequest{}).Where("loan_req_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"cust_user_id":          p.CustUserID,
			"mitra_user_id":         p.MitraUserID,
			"mitra_id":              p.MitraID,
			"mitra_branch_id":       p.MitraBranchID,
			"mitra_prod_id":         p.MitraProductID,
			"loan_req_schdl_status": p.LoanRequestSchdlStatus,
			"loan_req_updated_at":   time.Now(),
		},
	)
	err = db.Debug().Model(&LoanRequest{}).Where("loan_req_id = ?", pid).Error
	if err != nil {
		return &LoanRequest{}, err
	}
	return p, nil
}

func (p *LoanRequest) UpdateLoanRequestedAccept(db *gorm.DB, pid uint64) (*LoanRequest, error) {

	var err error
	db = db.Debug().Model(&LoanRequest{}).Where("loan_req_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_user_id":         p.MitraUserID,
			"mitra_branch_id":       p.MitraBranchID,
			"loan_req_status":       p.LoanRequestStatus,
			"loan_req_schdl_status": p.LoanRequestSchdlStatus,
			"loan_req_updated_at":   time.Now(),
		},
	)
	err = db.Debug().Model(&LoanRequest{}).Where("loan_req_id = ?", pid).Error
	if err != nil {
		return &LoanRequest{}, err
	}
	return p, nil
}

func (p *LoanRequest) UpdateLoanRequestedRespond(db *gorm.DB, pid uint64) (*LoanRequest, error) {

	var err error
	db = db.Debug().Model(&LoanRequest{}).Where("loan_req_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_req_status":       p.LoanRequestStatus,
			"loan_req_schdl_status": p.LoanRequestSchdlStatus,
			"loan_req_updated_at":   time.Now(),
		},
	)
	err = db.Debug().Model(&LoanRequest{}).Where("loan_req_id = ?", pid).Error
	if err != nil {
		return &LoanRequest{}, err
	}
	return p, nil
}

func (p *LoanRequest) UpdateLoanRequestedStatus(db *gorm.DB, pid uint64) (*LoanRequest, error) {

	var err error
	db = db.Debug().Model(&LoanRequest{}).Where("loan_req_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_req_status":     p.LoanRequestStatus,
			"loan_req_updated_by": p.LoanRequestUpdatedBy,
			"loan_req_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanRequest{}).Where("loan_req_id = ?", pid).Error
	if err != nil {
		return &LoanRequest{}, err
	}
	return p, nil
}

func (p *LoanRequest) UpdateLoanRequestStatusOnly(db *gorm.DB, pid uint64) (*LoanRequest, error) {

	var err error
	db = db.Debug().Model(&LoanRequest{}).Where("loan_req_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_req_status": p.LoanRequestStatus,
		},
	)
	err = db.Debug().Model(&LoanRequest{}).Where("loan_req_id = ?", pid).Error
	if err != nil {
		return &LoanRequest{}, err
	}
	return p, nil
}

func (p *LoanRequest) CreateLoanOffering(db *gorm.DB, pid uint64) (*LoanRequest, error) {

	var err error
	db = db.Debug().Model(&LoanRequest{}).Where("loan_req_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_id":            p.MitraID,
			"mitra_prod_id":       p.MitraProductID,
			"loan_req_status":     p.LoanRequestStatus,
			"loan_req_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanRequest{}).Where("loan_req_id = ?", pid).Error
	if err != nil {
		return &LoanRequest{}, err
	}
	return p, nil
}

func (p *LoanRequest) UpdateLoanReqSchdlStatus(db *gorm.DB, pid uint64) (*LoanRequest, error) {
	var err error
	db = db.Debug().Model(&LoanRequest{}).Where("loan_req_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_req_schdl_status": p.LoanRequestSchdlStatus,
			"loan_req_updated_at":   time.Now(),
		},
	)
	err = db.Debug().Model(&LoanRequest{}).Where("loan_req_id = ?", pid).Error
	if err != nil {
		return &LoanRequest{}, err
	}
	return p, nil
}

func (p *LoanRequest) UpdateLoanReqSchdlStatusDelete(db *gorm.DB, pid uint64) (*LoanRequest, error) {
	var err error
	db = db.Debug().Model(&LoanRequest{}).Where("loan_req_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_req_schdl_status": 3,
			"loan_req_deleted_at":   time.Now(),
		},
	)
	err = db.Debug().Model(&LoanRequest{}).Where("loan_req_id = ?", pid).Error
	if err != nil {
		return &LoanRequest{}, err
	}
	return p, nil
}

func (p *LoanRequest) DeleteLoanRequest(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&LoanRequest{}).Where("loan_req_id = ?", pid).Take(&LoanRequest{}).Delete(&LoanRequest{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("LoanRequest not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// ADDITIONAL
// ===============================================================================

func (p *LoanRequest) FindAllLoanRequestActByAOPhone(db *gorm.DB, phone string) (*[]LoanRequestAcceptedData, error) {
	var err error
	loan_req := []LoanRequestAcceptedData{}
	err = db.Debug().Model(&LoanRequestAcceptedData{}).Limit(100).
		Select(`t_loan_request.loan_req_id,
				t_loan_request.cust_user_id,
				m_product_sub_category.prod_sub_ctgry_id, 	
				m_product_sub_category.prod_sub_ctgry_name, 	
				m_mitra_user.mitra_user_id,	   	  
				m_customer.cust_name,	  
				m_country.country_id,   
				m_country.country_name,	 
				m_province.province_id,   
				m_province.province_name,	 
				m_regency.regency_id,   
				m_regency.regency_name,	    
				m_mitra.mitra_id, 
				t_loan_request.loan_req_amount,
				t_loan_request.loan_req_long,
				t_loan_request.loan_req_purpose,
				t_loan_request.loan_req_salary_range,
				t_loan_request.loan_req_status,
				t_loan_request.loan_req_created_at at time zone 'Asia/Jakarta' as loan_req_created_at`).
		Joins("JOIN m_product_sub_category on t_loan_request.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Joins("JOIN m_customer_user on t_loan_request.cust_user_id=m_customer_user.cust_user_id").
		Joins("JOIN m_customer on m_customer.cust_id=m_customer_user.cust_id").
		Joins("JOIN m_customer_address on m_customer.cust_id=m_customer_address.cust_id").
		Joins("JOIN m_address on m_customer_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_mitra_user on t_loan_request.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra on t_loan_request.mitra_id=m_mitra.mitra_id").
		Where("loan_req_status = 13 AND m_mitra_user.mitra_user_phone = ?", phone).
		Order("loan_req_created_at desc").
		Find(&loan_req).Error
	if err != nil {
		return &[]LoanRequestAcceptedData{}, err
	}
	return &loan_req, nil
}

func (p *LoanRequest) CountAllLoanRequestProccessByMitraUserID(db *gorm.DB, uid uint64) (*LoanRequestCountData, error) {
	var err error
	loan_req := LoanRequestCountData{}
	err = db.Debug().Model(&LoanRequestCountData{}).Limit(100).
		Select(`t_loan_request.loan_req_id,
				t_loan_request.cust_user_id,
				m_product_sub_category.prod_sub_ctgry_id, 	
				m_product_sub_category.prod_sub_ctgry_name, 	
				m_mitra_user.mitra_user_id,	   	  
				m_customer.cust_name,	  
				m_country.country_id,   
				m_country.country_name,	 
				m_province.province_id,   
				m_province.province_name,	 
				m_regency.regency_id,   
				m_regency.regency_name,	    
				m_mitra.mitra_id, 
				t_loan_request.loan_req_amount,
				t_loan_request.loan_req_long,
				t_loan_request.loan_req_purpose,
				t_loan_request.loan_req_salary_range,
				t_loan_request.loan_req_status,
				t_loan_request.loan_req_created_at at time zone 'Asia/Jakarta' as loan_req_created_at`).
		Joins("JOIN m_product_sub_category on t_loan_request.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Joins("JOIN m_customer_user on t_loan_request.cust_user_id=m_customer_user.cust_user_id").
		Joins("JOIN m_customer on m_customer.cust_id=m_customer_user.cust_id").
		Joins("JOIN m_customer_address on m_customer.cust_id=m_customer_address.cust_id").
		Joins("JOIN m_address on m_customer_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_mitra_user on t_loan_request.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra on t_loan_request.mitra_id=m_mitra.mitra_id").
		Where("loan_req_status = 6 OR loan_req_status = 7 OR loan_req_status = 8 OR loan_req_status = 9 OR loan_req_status = 10 AND m_mitra_user.mitra_user_id = ?", uid).
		Count(&loan_req.LoanRequestCount).
		Take(&loan_req).Error
	if err != nil {
		return &LoanRequestCountData{}, err
	}
	return &loan_req, nil
}

func (p *LoanRequest) AmountAllLoanRequestProccessByMitraUserID(db *gorm.DB, uid uint64) (*LoanRequestCountData, error) {
	var err error
	loan_req := LoanRequestCountData{}
	err = db.Debug().Model(&LoanRequestCountData{}).Limit(100).
		Select(`sum(t_loan_request.loan_req_amount) as loan_req_amount`).
		Joins("JOIN m_product_sub_category on t_loan_request.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Joins("JOIN m_customer_user on t_loan_request.cust_user_id=m_customer_user.cust_user_id").
		Joins("JOIN m_customer on m_customer.cust_id=m_customer_user.cust_id").
		Joins("JOIN m_customer_address on m_customer.cust_id=m_customer_address.cust_id").
		Joins("JOIN m_address on m_customer_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_mitra_user on t_loan_request.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra on t_loan_request.mitra_id=m_mitra.mitra_id").
		Where("loan_req_status = 6 OR loan_req_status = 7 OR loan_req_status = 8 OR loan_req_status = 9 OR loan_req_status = 10 AND m_mitra_user.mitra_user_id = ?", uid).
		Find(&loan_req).
		Take(&loan_req).Error
	if err != nil {
		return &LoanRequestCountData{}, err
	}
	return &loan_req, nil
}

func (p *LoanRequest) CountAllLoanRequestActByMitraUserID(db *gorm.DB, uid uint64) (*LoanRequestCountData, error) {
	var err error
	loan_req := LoanRequestCountData{}
	err = db.Debug().Model(&LoanRequestCountData{}).Limit(100).
		Select(`t_loan_request.loan_req_amount,
				t_loan_request.loan_req_id,
				t_loan_request.cust_user_id,
				m_product_sub_category.prod_sub_ctgry_id, 	
				m_product_sub_category.prod_sub_ctgry_name, 	
				m_mitra_user.mitra_user_id,	   	  
				m_customer.cust_name,	  
				m_country.country_id,   
				m_country.country_name,	 
				m_province.province_id,   
				m_province.province_name,	 
				m_regency.regency_id,   
				m_regency.regency_name,	    
				m_mitra.mitra_id, 
				t_loan_request.loan_req_long,
				t_loan_request.loan_req_purpose,
				t_loan_request.loan_req_salary_range,
				t_loan_request.loan_req_status,
				t_loan_request.loan_req_created_at at time zone 'Asia/Jakarta' as loan_req_created_at`).
		Joins("JOIN m_product_sub_category on t_loan_request.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Joins("JOIN m_customer_user on t_loan_request.cust_user_id=m_customer_user.cust_user_id").
		Joins("JOIN m_customer on m_customer.cust_id=m_customer_user.cust_id").
		Joins("JOIN m_customer_address on m_customer.cust_id=m_customer_address.cust_id").
		Joins("JOIN m_address on m_customer_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_mitra_user on t_loan_request.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra on t_loan_request.mitra_id=m_mitra.mitra_id").
		Where("loan_req_status = 12 OR loan_req_status = 13 AND m_mitra_user.mitra_user_id = ?", uid).
		Count(&loan_req.LoanRequestCount).
		Take(&loan_req).Error
	if err != nil {
		return &LoanRequestCountData{}, err
	}
	return &loan_req, nil
}

func (p *LoanRequest) AmountAllLoanRequestActByMitraUserID(db *gorm.DB, uid uint64) (*LoanRequestCountData, error) {
	var err error
	loan_req := LoanRequestCountData{}
	err = db.Debug().Model(&LoanRequestCountData{}).Limit(100).
		Select(`sum(t_loan_request.loan_req_amount) as loan_req_amount`).
		Joins("JOIN m_product_sub_category on t_loan_request.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Joins("JOIN m_customer_user on t_loan_request.cust_user_id=m_customer_user.cust_user_id").
		Joins("JOIN m_customer on m_customer.cust_id=m_customer_user.cust_id").
		Joins("JOIN m_customer_address on m_customer.cust_id=m_customer_address.cust_id").
		Joins("JOIN m_address on m_customer_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_mitra_user on t_loan_request.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra on t_loan_request.mitra_id=m_mitra.mitra_id").
		Where("loan_req_status = 12 OR loan_req_status = 13 AND m_mitra_user.mitra_user_id = ?", uid).
		Find(&loan_req).
		Take(&loan_req).Error
	if err != nil {
		return &LoanRequestCountData{}, err
	}
	return &loan_req, nil
}

// Collateral
// =======================================

type LoanReqCollateral struct {
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

type ResponseLoanReqCollaterals struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    *[]LoanRequestCollateral `json:"data"`
}

func (LoanReqCollateral) TableName() string {
	return "t_loan_request_collateral"
}

func (p *LoanRequestCollateral) FindLoanReqCollateralByReqID(db *gorm.DB, pid uint64) (*LoanRequestCollateral, error) {
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

// Generate Code
// =======================================

type GenerateCode struct {
	LoanReqNumberBefore string    `gorm:"column:loan_req_num_before" json:"loan_req_num_before"`
	LoanReqCreatedAt    time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	LoanReqNumberAfter  string    `gorm:"column:loan_req_num_after" json:"loan_req_num_after"`
}

func (p *LoanRequest) GenerateCode(db *gorm.DB) (*GenerateCode, error) {
	var err error
	loan_req := GenerateCode{}
	err = db.Debug().Model(&GenerateCode{}).Limit(1).
		Select(`RIGHT(t_loan_request.loan_req_num, 9) as loan_req_num_before, loan_req_created_at as created_at`).
		Order("loan_req_id DESC").
		Find(&loan_req).
		Take(&loan_req).Error
	if err != nil {
		return &GenerateCode{}, err
	}
	return &loan_req, nil
}
