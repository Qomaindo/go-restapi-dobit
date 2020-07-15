package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//BASE CRUD
//====================================================================================================================================

type MitraUserTemp struct {
	MitraUserTempID              uint64    `gorm:"column:mitra_user_temp_id;primary_key;" json:"mitra_user_temp_id"`
	MitraUserID                  uint64    `gorm:"column:mitra_user_id" json:"mitra_user_id"`
	MitraEmployeeTempID          uint64    `gorm:"column:mitra_emp_temp_id" json:"mitra_emp_temp_id"`
	MitraGroupAccessTempID       uint64    `gorm:"column:mitra_group_accs_temp_id" json:"mitra_group_accs_temp_id"`
	MitraGroupProductLimitTempID uint64    `gorm:"column:mitra_group_prod_lim_temp_id" json:"mitra_group_prod_lim_temp_id"`
	MitraAOCoverageTempID        uint64    `gorm:"column:mitra_ao_cov_temp_id" json:"mitra_ao_cov_temp_id"`
	MitraUserTempEmail           string    `gorm:"column:mitra_user_temp_email;size:255" json:"mitra_user_temp_email"`
	MitraUserTempUsername        string    `gorm:"column:mitra_user_temp_username;size:125" json:"mitra_user_temp_username"`
	MitraUserTempPasswordWeb     string    `gorm:"column:mitra_user_temp_password_web" json:"mitra_user_temp_password_web"`
	MitraUserTempPasswordMobile  string    `gorm:"column:mitra_user_temp_password_mobile" json:"mitra_user_temp_password_mobile"`
	MitraUserTempPhoneCode       string    `gorm:"column:mitra_user_temp_phone_code;size:6" json:"mitra_user_temp_phone_code"`
	MitraUserTempPhone           string    `gorm:"column:mitra_user_temp_phone;size:20" json:"mitra_user_temp_phone"`
	MitraUserTempPIN             string    `gorm:"column:mitra_user_temp_pin;size:255" json:"mitra_user_temp_pin"`
	MitraUserTempIsAO            bool      `gorm:"column:mitra_user_temp_is_ao" json:"mitra_user_temp_is_ao"`
	MitraUserTempIsCredit        bool      `gorm:"column:mitra_user_temp_is_credit" json:"mitra_user_temp_is_credit"`
	MitraUserTempReason          string    `gorm:"column:mitra_user_temp_reason" json:"mitra_user_temp_reason"`
	MitraUserTempActionBefore    uint64    `gorm:"column:mitra_user_temp_action_before" json:"mitra_user_temp_action_before"`
	MitraUserTempStatus          uint64    `gorm:"column:mitra_user_temp_status;size:2" json:"mitra_user_temp_status"`
	MitraUserTempCreatedBy       string    `gorm:"column:mitra_user_temp_created_by;size:125" json:"mitra_user_temp_created_by"`
	MitraUserTempCreatedAt       time.Time `gorm:"column:mitra_user_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_user_temp_created_at"`

	MitraUserBranchAccessTemp []MitraUserBranchAccessTempPendData `json:"mitra_user_branch_accs_temp"`
}

type MitraUserTempPend struct {
	MitraUserTempID              uint64    `gorm:"column:mitra_user_temp_id;primary_key;" json:"mitra_user_temp_id"`
	MitraUserID                  *uint64   `gorm:"column:mitra_user_id" json:"mitra_user_id"`
	MitraEmployeeTempID          uint64    `gorm:"column:mitra_emp_temp_id" json:"mitra_emp_temp_id"`
	MitraGroupAccessTempID       uint64    `gorm:"column:mitra_group_accs_temp_id" json:"mitra_group_accs_temp_id"`
	MitraGroupProductLimitTempID uint64    `gorm:"column:mitra_group_prod_lim_temp_id" json:"mitra_group_prod_lim_temp_id"`
	MitraAOCoverageTempID        uint64    `gorm:"column:mitra_ao_cov_temp_id" json:"mitra_ao_cov_temp_id"`
	MitraUserTempEmail           string    `gorm:"column:mitra_user_temp_email;size:255" json:"mitra_user_temp_email"`
	MitraUserTempUsername        string    `gorm:"column:mitra_user_temp_username;size:125" json:"mitra_user_temp_username"`
	MitraUserTempPasswordWeb     string    `gorm:"column:mitra_user_temp_password_web" json:"mitra_user_temp_password_web"`
	MitraUserTempPasswordMobile  string    `gorm:"column:mitra_user_temp_password_mobile" json:"mitra_user_temp_password_mobile"`
	MitraUserTempPhoneCode       string    `gorm:"column:mitra_user_temp_phone_code;size:6" json:"mitra_user_temp_phone_code"`
	MitraUserTempPhone           string    `gorm:"column:mitra_user_temp_phone;size:20" json:"mitra_user_temp_phone"`
	MitraUserTempPIN             string    `gorm:"column:mitra_user_temp_pin;size:255" json:"mitra_user_temp_pin"`
	MitraUserTempIsAO            bool      `gorm:"column:mitra_user_temp_is_ao" json:"mitra_user_temp_is_ao"`
	MitraUserTempIsCredit        bool      `gorm:"column:mitra_user_temp_is_credit" json:"mitra_user_temp_is_credit"`
	MitraUserTempReason          string    `gorm:"column:mitra_user_temp_reason" json:"mitra_user_temp_reason"`
	MitraUserTempActionBefore    uint64    `gorm:"column:mitra_user_temp_action_before" json:"mitra_user_temp_action_before"`
	MitraUserTempStatus          uint64    `gorm:"column:mitra_user_temp_status;size:2" json:"mitra_user_temp_status"`
	MitraUserTempCreatedBy       string    `gorm:"column:mitra_user_temp_created_by;size:125" json:"mitra_user_temp_created_by"`
	MitraUserTempCreatedAt       time.Time `gorm:"column:mitra_user_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_user_temp_created_at"`

	MitraUserBranchAccessTemp []MitraUserBranchAccessTempPendData `json:"mitra_user_branch_accs_temp"`
}

type MitraUserTempPendData struct {
	MitraUserTempID                uint64    `gorm:"column:mitra_user_temp_id;primary_key;" json:"mitra_user_temp_id"`
	MitraUserID                    *uint64   `gorm:"column:mitra_user_id" json:"mitra_user_id"`
	MitraEmployeeTempID            uint64    `gorm:"column:mitra_emp_temp_id" json:"mitra_emp_temp_id"`
	MitraEmployeeTempName          string    `gorm:"column:mitra_emp_temp_name" json:"mitra_emp_temp_name"`
	MitraEmployeeTempGender        string    `gorm:"column:mitra_emp_temp_gender" json:"mitra_emp_temp_gender"`
	MitraEmployeeTempBirthPlace    string    `gorm:"column:mitra_emp_temp_birth_place" json:"mitra_emp_temp_birth_place"`
	MitraEmployeeTempBirthDate     string    `gorm:"column:mitra_emp_temp_birth_date" json:"mitra_emp_temp_birth_date"`
	MitraEmployeeTempAddress       string    `gorm:"column:mitra_emp_temp_address" json:"mitra_emp_temp_address"`
	AddressTempID                  uint64    `gorm:"column:address_temp_id;size:255" json:"address_id"`
	MitraEmployeeTempPhoto         string    `gorm:"column:mitra_emp_temp_photo;size:255" json:"mitra_emp_temp_photo"`
	MitraTempID                    uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	MitraTempName                  string    `gorm:"column:mitra_temp_name" json:"mitra_temp_name"`
	MitraBranchTempID              uint64    `gorm:"column:mitra_branch_temp_id" json:"mitra_branch_temp_id"`
	MitraBranchTempName            string    `gorm:"column:mitra_branch_temp_name" json:"mitra_branch_temp_name"`
	MitraGroupAccessTempID         uint64    `gorm:"column:mitra_group_accs_temp_id" json:"mitra_group_accs_temp_id"`
	MitraGroupAccessTempName       string    `gorm:"column:mitra_group_accs_temp_name" json:"mitra_group_accs_temp_name"`
	MitraGroupProductLimitTempID   uint64    `gorm:"column:mitra_group_prod_lim_temp_id" json:"mitra_group_prod_lim_temp_id"`
	MitraAOCoverageTempID          uint64    `gorm:"column:mitra_ao_cov_temp_id" json:"mitra_ao_cov_temp_id"`
	MitraAOCoverageTempName        string    `gorm:"column:mitra_ao_cov_temp_name" json:"mitra_ao_cov_temp_name"`
	MitraGroupProductLimitTempName string    `gorm:"column:mitra_group_prod_lim_temp_name" json:"mitra_group_prod_lim_temp_name"`
	MitraUserTempEmail             string    `gorm:"column:mitra_user_temp_email;size:255" json:"mitra_user_temp_email"`
	MitraUserTempUsername          string    `gorm:"column:mitra_user_temp_username;size:125" json:"mitra_user_temp_username"`
	MitraUserTempPasswordWeb       string    `gorm:"column:mitra_user_temp_password_web" json:"mitra_user_temp_password_web"`
	MitraUserTempPasswordMobile    string    `gorm:"column:mitra_user_temp_password_mobile" json:"mitra_user_temp_password_mobile"`
	MitraUserTempPhoneCode         string    `gorm:"column:mitra_user_temp_phone_code;size:6" json:"mitra_user_temp_phone_code"`
	MitraUserTempPhone             string    `gorm:"column:mitra_user_temp_phone;size:20" json:"mitra_user_temp_phone"`
	MitraUserTempPIN               string    `gorm:"column:mitra_user_temp_pin;size:255" json:"mitra_user_temp_pin"`
	MitraUserTempIsAO              bool      `gorm:"column:mitra_user_temp_is_ao" json:"mitra_user_temp_is_ao"`
	MitraUserTempIsCredit          bool      `gorm:"column:mitra_user_temp_is_credit" json:"mitra_user_temp_is_credit"`
	MitraUserTempReason            string    `gorm:"column:mitra_user_temp_reason" json:"mitra_user_temp_reason"`
	MitraUserTempActionBefore      uint64    `gorm:"column:mitra_user_temp_action_before" json:"mitra_user_temp_action_before"`
	MitraUserTempStatus            uint64    `gorm:"column:mitra_user_temp_status;size:2" json:"mitra_user_temp_status"`
	MitraUserTempCreatedBy         string    `gorm:"column:mitra_user_temp_created_by;size:125" json:"mitra_user_temp_created_by"`
	MitraUserTempCreatedAt         time.Time `gorm:"column:mitra_user_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_user_temp_created_at"`

	MitraUserBranchAccessTemp []MitraUserBranchAccessTempPendData `json:"mitra_user_branch_accs_temp"`
}

type MitraUserTempData struct {
	MitraUserTempID                uint64    `gorm:"column:mitra_user_temp_id" json:"mitra_user_temp_id"`
	MitraEmployeeTempID            uint64    `gorm:"column:mitra_emp_temp_id" json:"mitra_emp_temp_id"`
	MitraEmployeeTempName          string    `gorm:"column:mitra_emp_temp_name;size:255" json:"mitra_emp_temp_name"`
	MitraEmployeeTempGender        string    `gorm:"column:mitra_emp_temp_gender;size:255" json:"mitra_emp_temp_gender"`
	MitraEmployeeTempBirthPlace    string    `gorm:"column:mitra_emp_temp_birth_place;size:255" json:"mitra_emp_temp_birth_place"`
	MitraEmployeeTempBirthDate     string    `gorm:"column:mitra_emp_temp_birth_date;size:255" json:"mitra_emp_temp_birth_date"`
	MitraEmployeeTempAddress       string    `gorm:"column:mitra_emp_temp_address;size:255" json:"mitra_emp_temp_address"`
	AddressTempID                  uint64    `gorm:"column:address_temp_id;size:255" json:"address_id"`
	MitraEmployeeTempPhoto         string    `gorm:"column:mitra_emp_temp_photo;size:255" json:"mitra_emp_temp_photo"`
	MitraTempID                    uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	MitraTempName                  string    `gorm:"column:mitra_temp_name" json:"mitra_temp_name"`
	MitraBranchTempID              uint64    `gorm:"column:mitra_branch_temp_id" json:"mitra_branch_temp_id"`
	MitraBranchTempName            string    `gorm:"column:mitra_branch_temp_name" json:"mitra_branch_temp_name"`
	MitraGroupAccessTempID         uint64    `gorm:"column:mitra_group_accs_temp_id" json:"mitra_group_accs_temp_id"`
	MitraGroupAccessTempName       string    `gorm:"column:mitra_group_accs_temp_name;size:255" json:"mitra_group_accs_temp_name"`
	MitraGroupProductLimitTempID   uint64    `gorm:"column:mitra_group_prod_lim_temp_id" json:"mitra_group_prod_lim_temp_id"`
	MitraGroupProductLimitTempName string    `gorm:"column:mitra_group_prod_lim_temp_name" json:"mitra_group_prod_lim_temp_name"`
	MitraAOCoverageTempID          uint64    `gorm:"column:mitra_ao_cov_temp_id" json:"mitra_ao_cov_temp_id"`
	MitraAOCoverageTempName        string    `gorm:"column:mitra_ao_cov_temp_name" json:"mitra_ao_cov_temp_name"`
	MitraUserTempEmail             string    `gorm:"column:mitra_user_temp_email;size:255" json:"mitra_user_temp_email"`
	MitraUserTempUsername          string    `gorm:"column:mitra_user_temp_username;size:125" json:"mitra_user_temp_username"`
	MitraUserTempPhoneCode         string    `gorm:"column:mitra_user_temp_phone_code;size:6" json:"mitra_user_temp_phone_code"`
	MitraUserTempPhone             string    `gorm:"column:mitra_user_temp_phone;size:20" json:"mitra_user_temp_phone"`
	MitraUserTempIsAO              bool      `gorm:"column:mitra_user_temp_is_ao" json:"mitra_user_temp_is_ao"`
	MitraUserTempIsCredit          bool      `gorm:"column:mitra_user_temp_is_credit" json:"mitra_user_temp_is_credit"`
	MitraUserTempReason            string    `gorm:"column:mitra_user_temp_reason" json:"mitra_user_temp_reason"`
	MitraUserTempActionBefore      uint64    `gorm:"column:mitra_user_temp_action_before" json:"mitra_user_temp_action_before"`
	MitraUserTempStatus            uint64    `gorm:"column:mitra_user_temp_status;size:2" json:"mitra_user_temp_status"`
	MitraUserTempCreatedBy         string    `gorm:"column:mitra_user_temp_created_by;size:125" json:"mitra_user_temp_created_by"`
	MitraUserTempCreatedAt         time.Time `gorm:"column:mitra_user_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_user_temp_created_at"`

	MitraUserBranchAccessTemp []MitraUserBranchAccessTempData `json:"mitra_user_branch_accs_temp"`

	MitraUserID                      uint64    `gorm:"column:mitra_user_id" json:"mitra_user_id"`
	MitraEmployeeID                  uint64    `gorm:"column:mitra_emp_id" json:"mitra_emp_id"`
	MitraEmployeeName                string    `gorm:"column:mitra_emp_name;size:255" json:"mitra_emp_name"`
	MitraEmployeeGender              string    `gorm:"column:mitra_emp_gender;size:255" json:"mitra_emp_gender"`
	MitraEmployeeBirthPlace          string    `gorm:"column:mitra_emp_birth_place;size:255" json:"mitra_emp_birth_place"`
	MitraEmployeeBirthDate           string    `gorm:"column:mitra_emp_birth_date;size:255" json:"mitra_emp_birth_date"`
	MitraEmployeeAddress             string    `gorm:"column:mitra_emp_address;size:255" json:"mitra_emp_address"`
	AddressID                        uint64    `gorm:"column:address_id;size:255" json:"address_id"`
	MitraEmployeePhoto               string    `gorm:"column:mitra_emp_photo;size:255" json:"mitra_emp_photo"`
	MitraID                          uint64    `gorm:"column:mitra_id" json:"mitra_id"`
	MitraName                        string    `gorm:"column:mitra_name" json:"mitra_name"`
	MitraBranchID                    uint64    `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	MitraBranchName                  string    `gorm:"column:mitra_branch_name" json:"mitra_branch_name"`
	MitraGroupAccessID               uint64    `gorm:"column:mitra_group_accs_id" json:"mitra_group_accs_id"`
	MitraGroupAccessName             string    `gorm:"column:mitra_group_accs_name;size:255" json:"mitra_group_accs_name"`
	MitraGroupProductLimitID         uint64    `gorm:"column:mitra_group_prod_lim_id" json:"mitra_group_prod_lim_id"`
	MitraGroupProductLimitName       string    `gorm:"column:mitra_group_prod_lim_name" json:"mitra_group_prod_lim_name"`
	MitraAOCoverageID                uint64    `gorm:"column:mitra_ao_cov_id" json:"mitra_ao_cov_id"`
	MitraAOCoverageName              string    `gorm:"column:mitra_ao_cov_name" json:"mitra_ao_cov_name"`
	MitraUserEmail                   string    `gorm:"column:mitra_user_email" json:"mitra_user_email"`
	MitraUsername                    string    `gorm:"column:mitra_username" json:"mitra_username"`
	MitraUserPhoneCode               string    `gorm:"column:mitra_user_phone_code;size:6" json:"mitra_user_phone_code"`
	MitraUserPhone                   string    `gorm:"column:mitra_user_phone;size:20" json:"mitra_user_phone"`
	MitraUserIsAO                    bool      `gorm:"column:mitra_user_is_ao" json:"mitra_user_is_ao"`
	MitraUserIsCredit                bool      `gorm:"column:mitra_user_is_credit" json:"mitra_user_is_credit"`
	MitraUserStatus                  uint64    `gorm:"column:mitra_user_status;size:1" json:"mitra_user_status"`
	MitraUserIsNew                   bool      `gorm:"column:mitra_user_is_new" json:"mitra_user_is_new"`
	MitraUserIsRequestChangePassword bool      `gorm:"column:mitra_user_is_req_chpw" json:"mitra_user_is_req_chpw"`
	MitraUserIsLocked                bool      `gorm:"column:mitra_user_is_locked" json:"mitra_user_is_locked"`
	MitraUserFailAttempt             uint64    `gorm:"column:mitra_user_fail_attempt;size:1" json:"mitra_user_fail_attempt"`
	MitraUserLockedReason            string    `gorm:"column:mitra_user_locked_reason" json:"mitra_user_locked_reason"`
	MitraUserCreatedBy               string    `gorm:"column:mitra_user_created_by;size:125" json:"mitra_user_created_by"`
	MitraUserCreatedAt               time.Time `gorm:"column:mitra_user_created_at;default:CURRENT_TIMESTAMP" json:"mitra_user_created_at"`

	MitraUserBranchAccess []MitraUserBranchAccessData `json:"mitra_user_branch_accs"`
}

type ResponseMitraUserTemps struct {
	Status  int                  `json:"status"`
	Message string               `json:"message"`
	Data    *[]MitraUserTempData `json:"data"`
}

type ResponseMitraUserTemp struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    *MitraUserTempData `json:"data"`
}

type ResponseMitraUserTempsPend struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    *[]MitraUserTempPendData `json:"data"`
}

type ResponseMitraUserTempPend struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    *MitraUserTempPendData `json:"data"`
}

type ResponseMitraUserTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraUserTemp) TableName() string {
	return "m_mitra_user_temp"
}

func (MitraUserTempPend) TableName() string {
	return "m_mitra_user_temp"
}

func (MitraUserTempPendData) TableName() string {
	return "m_mitra_user_temp"
}

func (MitraUserTempData) TableName() string {
	return "m_mitra_user_temp"
}

func (p *MitraUserTemp) Prepare() {
	p.MitraUserID = p.MitraUserID
	p.MitraEmployeeTempID = p.MitraEmployeeTempID
	p.MitraGroupAccessTempID = p.MitraGroupAccessTempID
	p.MitraGroupProductLimitTempID = p.MitraGroupProductLimitTempID
	p.MitraAOCoverageTempID = p.MitraAOCoverageTempID
	p.MitraUserTempEmail = html.EscapeString(strings.TrimSpace(p.MitraUserTempEmail))
	p.MitraUserTempUsername = html.EscapeString(strings.TrimSpace(p.MitraUserTempUsername))
	p.MitraUserTempPasswordWeb = html.EscapeString(strings.TrimSpace(p.MitraUserTempPasswordWeb))
	p.MitraUserTempPasswordMobile = html.EscapeString(strings.TrimSpace(p.MitraUserTempPasswordMobile))
	p.MitraUserTempPhoneCode = html.EscapeString(strings.TrimSpace(p.MitraUserTempPhoneCode))
	p.MitraUserTempPhone = html.EscapeString(strings.TrimSpace(p.MitraUserTempPhone))
	p.MitraUserTempPIN = html.EscapeString(strings.TrimSpace(p.MitraUserTempPIN))
	p.MitraUserTempIsAO = p.MitraUserTempIsAO
	p.MitraUserTempIsCredit = p.MitraUserTempIsCredit
	p.MitraUserTempReason = html.EscapeString(strings.TrimSpace(p.MitraUserTempReason))
	p.MitraUserTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraUserTempCreatedBy))
	p.MitraUserTempCreatedAt = time.Now()
}

func (p *MitraUserTempPend) Prepare() {
	p.MitraUserID = nil
	p.MitraEmployeeTempID = p.MitraEmployeeTempID
	p.MitraGroupAccessTempID = p.MitraGroupAccessTempID
	p.MitraGroupProductLimitTempID = p.MitraGroupProductLimitTempID
	p.MitraAOCoverageTempID = p.MitraAOCoverageTempID
	p.MitraUserTempEmail = html.EscapeString(strings.TrimSpace(p.MitraUserTempEmail))
	p.MitraUserTempUsername = html.EscapeString(strings.TrimSpace(p.MitraUserTempUsername))
	p.MitraUserTempPasswordWeb = html.EscapeString(strings.TrimSpace(p.MitraUserTempPasswordWeb))
	p.MitraUserTempPasswordMobile = html.EscapeString(strings.TrimSpace(p.MitraUserTempPasswordMobile))
	p.MitraUserTempPhoneCode = html.EscapeString(strings.TrimSpace(p.MitraUserTempPhoneCode))
	p.MitraUserTempPhone = html.EscapeString(strings.TrimSpace(p.MitraUserTempPhone))
	p.MitraUserTempPIN = html.EscapeString(strings.TrimSpace(p.MitraUserTempPIN))
	p.MitraUserTempIsAO = p.MitraUserTempIsAO
	p.MitraUserTempIsCredit = p.MitraUserTempIsCredit
	p.MitraUserTempReason = html.EscapeString(strings.TrimSpace(p.MitraUserTempReason))
	p.MitraUserTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraUserTempCreatedBy))
	p.MitraUserTempCreatedAt = time.Now()
}

func (p *MitraUserTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.MitraUserTempPasswordWeb == "" {
			return errors.New("Required MitraUserTemp Password")
		}
		if p.MitraUserTempEmail == "" {
			return errors.New("Required MitraUserTemp Email")
		}
		return nil
	}
}

func (p *MitraUserTemp) SaveMitraUserTemp(db *gorm.DB) (*MitraUserTemp, error) {
	var err error
	err = p.BeforeSaveMitraUser()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Debug().Model(&MitraUserTemp{}).Create(&p).Error
	if err != nil {
		return &MitraUserTemp{}, err
	}
	return p, nil
}

func (p *MitraUserTempPend) SaveMitraUserTempPend(db *gorm.DB) (*MitraUserTempPend, error) {
	var err error
	err = p.BeforeSaveMitraUser()
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().Model(&MitraUserTempPend{}).Create(&p).Error
	if err != nil {
		return &MitraUserTempPend{}, err
	}
	return p, nil
}

func (p *MitraUserTemp) FindAllMitraUserTemp(db *gorm.DB) (*[]MitraUserTempPendData, error) {
	var err error
	mitrauser := []MitraUserTempPendData{}
	err = db.Debug().Model(&MitraUserTempPendData{}).Limit(100).
		Select(`m_mitra_user_temp.mitra_user_temp_id,
				m_mitra_user_temp.mitra_user_id,
				m_mitra_employee_temp.mitra_emp_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_gender,
				m_mitra_employee_temp.mitra_emp_temp_birth_place,
				m_mitra_employee_temp.mitra_emp_temp_birth_date,
				m_mitra_employee_temp.mitra_emp_temp_address,
				m_mitra_employee_temp.address_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_photo,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_mitra_group_access_temp.mitra_group_accs_id as mitra_group_accs_temp_id,
				m_mitra_group_access_temp.mitra_group_accs_name as mitra_group_accs_temp_name,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_id as mitra_group_prod_lim_temp_id,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_name as mitra_group_prod_lim_temp_name,
				m_mitra_ao_coverage_temp.mitra_ao_cov_id as mitra_ao_cov_temp_id,
				m_mitra_ao_coverage_temp.mitra_ao_cov_name as mitra_ao_cov_temp_name,
				m_mitra_user_temp.mitra_user_temp_email,
				m_mitra_user_temp.mitra_user_temp_username,
				m_mitra_user_temp.mitra_user_temp_phone_code,
				m_mitra_user_temp.mitra_user_temp_phone,
				m_mitra_user_temp.mitra_user_temp_is_ao,
				m_mitra_user_temp.mitra_user_temp_is_credit,
				m_mitra_user_temp.mitra_user_temp_status,
				m_mitra_user_temp.mitra_user_temp_reason,
				m_mitra_user_temp.mitra_user_temp_action_before,
				m_mitra_user_temp.mitra_user_temp_created_by,
				m_mitra_user_temp.mitra_user_temp_created_at`).
		Joins("JOIN m_mitra_employee_temp on m_mitra_user_temp.mitra_emp_temp_id=m_mitra_employee_temp.mitra_emp_temp_id").
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_employee_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_employee_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_mitra_group_access m_mitra_group_access_temp on m_mitra_user_temp.mitra_group_accs_temp_id=m_mitra_group_access_temp.mitra_group_accs_id").
		Joins("JOIN m_mitra_group_product_limit m_mitra_group_product_limit_temp on m_mitra_user_temp.mitra_group_prod_lim_temp_id=m_mitra_group_product_limit_temp.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_ao_coverage m_mitra_ao_coverage_temp on m_mitra_user_temp.mitra_ao_cov_temp_id=m_mitra_ao_coverage_temp.mitra_ao_cov_id").
		Find(&mitrauser).Error
	if err != nil {
		return &[]MitraUserTempPendData{}, err
	}
	return &mitrauser, nil
}

func (p *MitraUserTemp) FindAllMitraUserTempStatusPendingActive(db *gorm.DB, mitra uint64, status uint64) (*[]MitraUserTempPendData, error) {
	var err error
	mitrauser := []MitraUserTempPendData{}
	err = db.Debug().Model(&MitraUserTempPendData{}).Limit(100).
		Select(`m_mitra_user_temp.mitra_user_temp_id,
				m_mitra_user_temp.mitra_user_id,
				m_mitra_employee_temp.mitra_emp_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_gender,
				m_mitra_employee_temp.mitra_emp_temp_birth_place,
				m_mitra_employee_temp.mitra_emp_temp_birth_date,
				m_mitra_employee_temp.mitra_emp_temp_address,
				m_mitra_employee_temp.address_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_photo,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_mitra_group_access_temp.mitra_group_accs_id as mitra_group_accs_temp_id,
				m_mitra_group_access_temp.mitra_group_accs_name as mitra_group_accs_temp_name,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_id as mitra_group_prod_lim_temp_id,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_name as mitra_group_prod_lim_temp_name,
				m_mitra_ao_coverage_temp.mitra_ao_cov_id as mitra_ao_cov_temp_id,
				m_mitra_ao_coverage_temp.mitra_ao_cov_name as mitra_ao_cov_temp_name,
				m_mitra_user_temp.mitra_user_temp_email,
				m_mitra_user_temp.mitra_user_temp_username,
				m_mitra_user_temp.mitra_user_temp_phone_code,
				m_mitra_user_temp.mitra_user_temp_phone,
				m_mitra_user_temp.mitra_user_temp_is_ao,
				m_mitra_user_temp.mitra_user_temp_is_credit,
				m_mitra_user_temp.mitra_user_temp_status,
				m_mitra_user_temp.mitra_user_temp_reason,
				m_mitra_user_temp.mitra_user_temp_action_before,
				m_mitra_user_temp.mitra_user_temp_created_by,
				m_mitra_user_temp.mitra_user_temp_created_at`).
		Joins("JOIN m_mitra_employee_temp on m_mitra_user_temp.mitra_emp_temp_id=m_mitra_employee_temp.mitra_emp_temp_id").
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_employee_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_employee_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_mitra_group_access m_mitra_group_access_temp on m_mitra_user_temp.mitra_group_accs_temp_id=m_mitra_group_access_temp.mitra_group_accs_id").
		Joins("JOIN m_mitra_group_product_limit m_mitra_group_product_limit_temp on m_mitra_user_temp.mitra_group_prod_lim_temp_id=m_mitra_group_product_limit_temp.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_ao_coverage m_mitra_ao_coverage_temp on m_mitra_user_temp.mitra_ao_cov_temp_id=m_mitra_ao_coverage_temp.mitra_ao_cov_id").
		Where("m_mitra_user_temp.mitra_user_temp_status = ? AND m_mitra_employee_temp.mitra_temp_id = ?", status, mitra).
		Find(&mitrauser).Error
	if err != nil {
		return &[]MitraUserTempPendData{}, err
	}
	return &mitrauser, nil
}

func (p *MitraUserTemp) FindAllMitraUserTempStatus(db *gorm.DB, status uint64, mitra uint64) (*[]MitraUserTempData, error) {
	var err error
	mitrauser := []MitraUserTempData{}
	err = db.Debug().Model(&MitraUserTempData{}).Limit(100).
		Select(`m_mitra_user_temp.mitra_user_temp_id,
				m_mitra_user_temp.mitra_user_id,
				m_mitra_employee_temp.mitra_emp_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_gender,
				m_mitra_employee_temp.mitra_emp_temp_birth_place,
				m_mitra_employee_temp.mitra_emp_temp_birth_date,
				m_mitra_employee_temp.mitra_emp_temp_address,
				m_mitra_employee_temp.address_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_photo,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_mitra_group_access_temp.mitra_group_accs_id as mitra_group_accs_temp_id,
				m_mitra_group_access_temp.mitra_group_accs_name as mitra_group_accs_temp_name,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_id as mitra_group_prod_lim_temp_id,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_name as mitra_group_prod_lim_temp_name,
				m_mitra_ao_coverage_temp.mitra_ao_cov_id as mitra_ao_cov_temp_id,
				m_mitra_ao_coverage_temp.mitra_ao_cov_name as mitra_ao_cov_temp_name,
				m_mitra_user_temp.mitra_user_temp_email,
				m_mitra_user_temp.mitra_user_temp_username,
				m_mitra_user_temp.mitra_user_temp_phone_code,
				m_mitra_user_temp.mitra_user_temp_phone,
				m_mitra_user_temp.mitra_user_temp_is_ao,
				m_mitra_user_temp.mitra_user_temp_is_credit,
				m_mitra_user_temp.mitra_user_temp_status,
				m_mitra_user_temp.mitra_user_temp_reason,
				m_mitra_user_temp.mitra_user_temp_action_before,
				m_mitra_user_temp.mitra_user_temp_created_by,
				m_mitra_user_temp.mitra_user_temp_created_at,
				m_mitra_user_temp.mitra_user_temp_id,
				m_mitra_user.mitra_user_id,
				m_mitra_employee.mitra_emp_id,
				m_mitra_employee.mitra_emp_name,
				m_mitra_employee.mitra_emp_gender,
				m_mitra_employee.mitra_emp_birth_place,
				m_mitra_employee.mitra_emp_birth_date,
				m_mitra_employee.mitra_emp_address,
				m_mitra_employee.address_id,
				m_mitra_employee.mitra_emp_photo,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_group_access.mitra_group_accs_id,
				m_mitra_group_access.mitra_group_accs_name,
				m_mitra_group_product_limit.mitra_group_prod_lim_id,
				m_mitra_group_product_limit.mitra_group_prod_lim_name,
				m_mitra_ao_coverage.mitra_ao_cov_id as mitra_ao_cov_id,
				m_mitra_ao_coverage.mitra_ao_cov_name as mitra_ao_cov_name,
				m_mitra_user.mitra_user_email,
				m_mitra_user.mitra_username,
				m_mitra_user.mitra_user_phone_code,
				m_mitra_user.mitra_user_phone,
				m_mitra_user.mitra_user_is_ao,
				m_mitra_user.mitra_user_is_credit,
				m_mitra_user.mitra_user_status,
				m_mitra_user.mitra_user_is_new,
				m_mitra_user.mitra_user_is_req_chpw,
				m_mitra_user.mitra_user_is_locked,
				m_mitra_user.mitra_user_fail_attempt,
				m_mitra_user.mitra_user_locked_reason,
				m_mitra_user.mitra_user_created_by,
				m_mitra_user.mitra_user_created_at`).
		Joins("JOIN m_mitra_employee_temp on m_mitra_user_temp.mitra_emp_temp_id=m_mitra_employee_temp.mitra_emp_temp_id").
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_employee_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_employee_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_mitra_group_access m_mitra_group_access_temp on m_mitra_user_temp.mitra_group_accs_temp_id=m_mitra_group_access_temp.mitra_group_accs_id").
		Joins("JOIN m_mitra_group_product_limit m_mitra_group_product_limit_temp on m_mitra_user_temp.mitra_group_prod_lim_temp_id=m_mitra_group_product_limit_temp.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_ao_coverage m_mitra_ao_coverage_temp on m_mitra_user_temp.mitra_ao_cov_temp_id=m_mitra_ao_coverage_temp.mitra_ao_cov_id").
		Joins("JOIN m_mitra_user on m_mitra_user_temp.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra_employee on m_mitra_user.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra_group_access on m_mitra_user.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Joins("JOIN m_mitra_group_product_limit on m_mitra_user.mitra_group_prod_lim_id=m_mitra_group_product_limit.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_ao_coverage on m_mitra_user.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Where("mitra_user_temp_status = ? AND m_mitra_employee_temp.mitra_temp_id = ?", status, mitra).
		Find(&mitrauser).Error
	if err != nil {
		return &[]MitraUserTempData{}, err
	}
	return &mitrauser, nil
}

func (p *MitraUserTemp) FindMitraUserTempDataByID(db *gorm.DB, pid uint64) (*MitraUserTemp, error) {
	var err error
	err = db.Debug().Model(&MitraUserTemp{}).Where("mitra_user_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraUserTemp{}, err
	}
	return p, nil
}

func (p *MitraUserTemp) FindMitraUserTempByIDPendingActive(db *gorm.DB, pid uint64) (*MitraUserTempPendData, error) {
	var err error
	mitrauser := MitraUserTempPendData{}
	err = db.Debug().Model(&MitraUserTempPendData{}).Limit(100).
		Select(`m_mitra_user_temp.mitra_user_temp_id,
				m_mitra_user_temp.mitra_user_id,
				m_mitra_employee_temp.mitra_emp_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_gender,
				m_mitra_employee_temp.mitra_emp_temp_birth_place,
				m_mitra_employee_temp.mitra_emp_temp_birth_date,
				m_mitra_employee_temp.mitra_emp_temp_address,
				m_mitra_employee_temp.address_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_photo,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_mitra_group_access_temp.mitra_group_accs_id as mitra_group_accs_temp_id,
				m_mitra_group_access_temp.mitra_group_accs_name as mitra_group_accs_temp_name,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_id as mitra_group_prod_lim_temp_id,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_name as mitra_group_prod_lim_temp_name,
				m_mitra_ao_coverage_temp.mitra_ao_cov_id as mitra_ao_cov_temp_id,
				m_mitra_ao_coverage_temp.mitra_ao_cov_name as mitra_ao_cov_temp_name,
				m_mitra_user_temp.mitra_user_temp_email,
				m_mitra_user_temp.mitra_user_temp_username,
				m_mitra_user_temp.mitra_user_temp_pin,
				m_mitra_user_temp.mitra_user_temp_phone_code,
				m_mitra_user_temp.mitra_user_temp_phone,
				m_mitra_user_temp.mitra_user_temp_is_ao,
				m_mitra_user_temp.mitra_user_temp_is_credit,
				m_mitra_user_temp.mitra_user_temp_status,
				m_mitra_user_temp.mitra_user_temp_reason,
				m_mitra_user_temp.mitra_user_temp_action_before,
				m_mitra_user_temp.mitra_user_temp_created_by,
				m_mitra_user_temp.mitra_user_temp_created_at`).
		Joins("JOIN m_mitra_employee_temp on m_mitra_user_temp.mitra_emp_temp_id=m_mitra_employee_temp.mitra_emp_temp_id").
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_employee_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_employee_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_mitra_group_access m_mitra_group_access_temp on m_mitra_user_temp.mitra_group_accs_temp_id=m_mitra_group_access_temp.mitra_group_accs_id").
		Joins("JOIN m_mitra_group_product_limit m_mitra_group_product_limit_temp on m_mitra_user_temp.mitra_group_prod_lim_temp_id=m_mitra_group_product_limit_temp.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_ao_coverage m_mitra_ao_coverage_temp on m_mitra_user_temp.mitra_ao_cov_temp_id=m_mitra_ao_coverage_temp.mitra_ao_cov_id").
		Where("mitra_user_temp_id = ?", pid).
		Find(&mitrauser).Error
	if err != nil {
		return &MitraUserTempPendData{}, err
	}
	return &mitrauser, nil
}

func (p *MitraUserTemp) FindMitraUserTempStatusByIDPendingActive(db *gorm.DB, pid uint64, status uint64, mitra uint64) (*MitraUserTempPendData, error) {
	var err error
	mitrauser := MitraUserTempPendData{}
	err = db.Debug().Model(&MitraUserTempPendData{}).Limit(100).
		Select(`m_mitra_user_temp.mitra_user_temp_id,
				m_mitra_user_temp.mitra_user_id,
				m_mitra_employee_temp.mitra_emp_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_gender,
				m_mitra_employee_temp.mitra_emp_temp_birth_place,
				m_mitra_employee_temp.mitra_emp_temp_birth_date,
				m_mitra_employee_temp.mitra_emp_temp_address,
				m_mitra_employee_temp.address_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_photo,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_mitra_group_access_temp.mitra_group_accs_id as mitra_group_accs_temp_id,
				m_mitra_group_access_temp.mitra_group_accs_name as mitra_group_accs_temp_name,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_id as mitra_group_prod_lim_temp_id,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_name as mitra_group_prod_lim_temp_name,
				m_mitra_ao_coverage_temp.mitra_ao_cov_id as mitra_ao_cov_temp_id,
				m_mitra_ao_coverage_temp.mitra_ao_cov_name as mitra_ao_cov_temp_name,
				m_mitra_user_temp.mitra_user_temp_email,
				m_mitra_user_temp.mitra_user_temp_username,
				m_mitra_user_temp.mitra_user_temp_phone_code,
				m_mitra_user_temp.mitra_user_temp_phone,
				m_mitra_user_temp.mitra_user_temp_is_ao,
				m_mitra_user_temp.mitra_user_temp_is_credit,
				m_mitra_user_temp.mitra_user_temp_status,
				m_mitra_user_temp.mitra_user_temp_reason,
				m_mitra_user_temp.mitra_user_temp_action_before,
				m_mitra_user_temp.mitra_user_temp_created_by,
				m_mitra_user_temp.mitra_user_temp_created_at`).
		Joins("JOIN m_mitra_employee_temp on m_mitra_user_temp.mitra_emp_temp_id=m_mitra_employee_temp.mitra_emp_temp_id").
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_employee_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_employee_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_mitra_group_access m_mitra_group_access_temp on m_mitra_user_temp.mitra_group_accs_temp_id=m_mitra_group_access_temp.mitra_group_accs_id").
		Joins("JOIN m_mitra_group_product_limit m_mitra_group_product_limit_temp on m_mitra_user_temp.mitra_group_prod_lim_temp_id=m_mitra_group_product_limit_temp.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_ao_coverage m_mitra_ao_coverage_temp on m_mitra_user_temp.mitra_ao_cov_temp_id=m_mitra_ao_coverage_temp.mitra_ao_cov_id").
		Where("m_mitra_user_temp.mitra_user_temp_id = ? AND m_mitra_user_temp.mitra_user_temp_status = ? AND m_mitra_employee_temp.mitra_temp_id = ?", pid, status, mitra).
		Find(&mitrauser).Error
	if err != nil {
		return &MitraUserTempPendData{}, err
	}
	return &mitrauser, nil
}

func (p *MitraUserTemp) FindMitraUserTempByID(db *gorm.DB, pid uint64) (*MitraUserTempData, error) {
	var err error
	mitrauser := MitraUserTempData{}
	err = db.Debug().Model(&MitraUserTempData{}).
		Select(`m_mitra_user_temp.mitra_user_temp_id,
				m_mitra_user_temp.mitra_user_id,
				m_mitra_employee_temp.mitra_emp_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_gender,
				m_mitra_employee_temp.mitra_emp_temp_birth_place,
				m_mitra_employee_temp.mitra_emp_temp_birth_date,
				m_mitra_employee_temp.mitra_emp_temp_address,
				m_mitra_employee_temp.address_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_photo,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_mitra_group_access_temp.mitra_group_accs_id as mitra_group_accs_temp_id,
				m_mitra_group_access_temp.mitra_group_accs_name as mitra_group_accs_temp_name,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_id as mitra_group_prod_lim_temp_id,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_name as mitra_group_prod_lim_temp_name,
				m_mitra_ao_coverage_temp.mitra_ao_cov_id as mitra_ao_cov_temp_id,
				m_mitra_ao_coverage_temp.mitra_ao_cov_name as mitra_ao_cov_temp_name,
				m_mitra_user_temp.mitra_user_temp_email,
				m_mitra_user_temp.mitra_user_temp_username,
				m_mitra_user_temp.mitra_user_temp_phone_code,
				m_mitra_user_temp.mitra_user_temp_phone,
				m_mitra_user_temp.mitra_user_temp_is_ao,
				m_mitra_user_temp.mitra_user_temp_is_credit,
				m_mitra_user_temp.mitra_user_temp_status,
				m_mitra_user_temp.mitra_user_temp_reason,
				m_mitra_user_temp.mitra_user_temp_action_before,
				m_mitra_user_temp.mitra_user_temp_created_by,
				m_mitra_user_temp.mitra_user_temp_created_at,
				m_mitra_user_temp.mitra_user_temp_id,
				m_mitra_user.mitra_user_id,
				m_mitra_employee.mitra_emp_id,
				m_mitra_employee.mitra_emp_name,
				m_mitra_employee.mitra_emp_gender,
				m_mitra_employee.mitra_emp_birth_place,
				m_mitra_employee.mitra_emp_birth_date,
				m_mitra_employee.mitra_emp_address,
				m_mitra_employee.address_id,
				m_mitra_employee.mitra_emp_photo,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_group_access.mitra_group_accs_id,
				m_mitra_group_access.mitra_group_accs_name,
				m_mitra_group_product_limit.mitra_group_prod_lim_id,
				m_mitra_group_product_limit.mitra_group_prod_lim_name,
				m_mitra_ao_coverage.mitra_ao_cov_id as mitra_ao_cov_id,
				m_mitra_ao_coverage.mitra_ao_cov_name as mitra_ao_cov_name,
				m_mitra_user.mitra_user_email,
				m_mitra_user.mitra_username,
				m_mitra_user.mitra_user_phone_code,
				m_mitra_user.mitra_user_phone,
				m_mitra_user.mitra_user_is_ao,
				m_mitra_user.mitra_user_is_credit,
				m_mitra_user.mitra_user_status,
				m_mitra_user.mitra_user_is_new,
				m_mitra_user.mitra_user_is_req_chpw,
				m_mitra_user.mitra_user_is_locked,
				m_mitra_user.mitra_user_fail_attempt,
				m_mitra_user.mitra_user_locked_reason,
				m_mitra_user.mitra_user_created_by,
				m_mitra_user.mitra_user_created_at`).
		Joins("JOIN m_mitra_employee_temp on m_mitra_user_temp.mitra_emp_temp_id=m_mitra_employee_temp.mitra_emp_temp_id").
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_employee_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_employee_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_mitra_group_access m_mitra_group_access_temp on m_mitra_user_temp.mitra_group_accs_temp_id=m_mitra_group_access_temp.mitra_group_accs_id").
		Joins("JOIN m_mitra_group_product_limit m_mitra_group_product_limit_temp on m_mitra_user_temp.mitra_group_prod_lim_temp_id=m_mitra_group_product_limit_temp.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_ao_coverage m_mitra_ao_coverage_temp on m_mitra_user_temp.mitra_ao_cov_temp_id=m_mitra_ao_coverage_temp.mitra_ao_cov_id").
		Joins("JOIN m_mitra_user on m_mitra_user_temp.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra_employee on m_mitra_user.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra_group_access on m_mitra_user.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Joins("JOIN m_mitra_group_product_limit on m_mitra_user.mitra_group_prod_lim_id=m_mitra_group_product_limit.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_ao_coverage on m_mitra_user.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Where("mitra_user_temp_id = ?", pid).
		Take(&mitrauser).Error
	if err != nil {
		return &MitraUserTempData{}, err
	}
	return &mitrauser, nil
}

func (p *MitraUserTemp) FindMitraUserTempStatusByID(db *gorm.DB, pid uint64, status uint64, mitra uint64) (*MitraUserTempData, error) {
	var err error
	mitrauser := MitraUserTempData{}
	err = db.Debug().Model(&MitraUserTempData{}).
		Select(`m_mitra_user_temp.mitra_user_temp_id,
				m_mitra_user_temp.mitra_user_id,
				m_mitra_employee_temp.mitra_emp_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_gender,
				m_mitra_employee_temp.mitra_emp_temp_birth_place,
				m_mitra_employee_temp.mitra_emp_temp_birth_date,
				m_mitra_employee_temp.mitra_emp_temp_address,
				m_mitra_employee_temp.address_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_photo,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_mitra_group_access_temp.mitra_group_accs_id as mitra_group_accs_temp_id,
				m_mitra_group_access_temp.mitra_group_accs_name as mitra_group_accs_temp_name,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_id as mitra_group_prod_lim_temp_id,
				m_mitra_group_product_limit_temp.mitra_group_prod_lim_name as mitra_group_prod_lim_temp_name,
				m_mitra_ao_coverage_temp.mitra_ao_cov_id as mitra_ao_cov_temp_id,
				m_mitra_ao_coverage_temp.mitra_ao_cov_name as mitra_ao_cov_temp_name,
				m_mitra_user_temp.mitra_user_temp_email,
				m_mitra_user_temp.mitra_user_temp_username,
				m_mitra_user_temp.mitra_user_temp_phone_code,
				m_mitra_user_temp.mitra_user_temp_phone,
				m_mitra_user_temp.mitra_user_temp_is_ao,
				m_mitra_user_temp.mitra_user_temp_is_credit,
				m_mitra_user_temp.mitra_user_temp_status,
				m_mitra_user_temp.mitra_user_temp_reason,
				m_mitra_user_temp.mitra_user_temp_action_before,
				m_mitra_user_temp.mitra_user_temp_created_by,
				m_mitra_user_temp.mitra_user_temp_created_at,
				m_mitra_user_temp.mitra_user_temp_id,
				m_mitra_user.mitra_user_id,
				m_mitra_employee.mitra_emp_id,
				m_mitra_employee.mitra_emp_name,
				m_mitra_employee.mitra_emp_gender,
				m_mitra_employee.mitra_emp_birth_place,
				m_mitra_employee.mitra_emp_birth_date,
				m_mitra_employee.mitra_emp_address,
				m_mitra_employee.address_id,
				m_mitra_employee.mitra_emp_photo,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_group_access.mitra_group_accs_id,
				m_mitra_group_access.mitra_group_accs_name,
				m_mitra_group_product_limit.mitra_group_prod_lim_id,
				m_mitra_group_product_limit.mitra_group_prod_lim_name,
				m_mitra_ao_coverage.mitra_ao_cov_id as mitra_ao_cov_id,
				m_mitra_ao_coverage.mitra_ao_cov_name as mitra_ao_cov_name,
				m_mitra_user.mitra_user_email,
				m_mitra_user.mitra_username,
				m_mitra_user.mitra_user_phone_code,
				m_mitra_user.mitra_user_phone,
				m_mitra_user.mitra_user_is_ao,
				m_mitra_user.mitra_user_is_credit,
				m_mitra_user.mitra_user_status,
				m_mitra_user.mitra_user_is_new,
				m_mitra_user.mitra_user_is_req_chpw,
				m_mitra_user.mitra_user_is_locked,
				m_mitra_user.mitra_user_fail_attempt,
				m_mitra_user.mitra_user_locked_reason,
				m_mitra_user.mitra_user_created_by,
				m_mitra_user.mitra_user_created_at`).
		Joins("JOIN m_mitra_employee_temp on m_mitra_user_temp.mitra_emp_temp_id=m_mitra_employee_temp.mitra_emp_temp_id").
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_employee_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_employee_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_mitra_group_access m_mitra_group_access_temp on m_mitra_user_temp.mitra_group_accs_temp_id=m_mitra_group_access_temp.mitra_group_accs_id").
		Joins("JOIN m_mitra_group_product_limit m_mitra_group_product_limit_temp on m_mitra_user_temp.mitra_group_prod_lim_temp_id=m_mitra_group_product_limit_temp.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_ao_coverage m_mitra_ao_coverage_temp on m_mitra_user_temp.mitra_ao_cov_temp_id=m_mitra_ao_coverage_temp.mitra_ao_cov_id").
		Joins("JOIN m_mitra_user on m_mitra_user_temp.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra_employee on m_mitra_user.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra_group_access on m_mitra_user.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Joins("JOIN m_mitra_group_product_limit on m_mitra_user.mitra_group_prod_lim_id=m_mitra_group_product_limit.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_ao_coverage on m_mitra_user.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Where("mitra_user_temp_id = ? AND mitra_user_temp_status = ? AND m_mitra_employee.mitra_id = ?", pid, status, mitra).
		Take(&mitrauser).Error
	if err != nil {
		return &MitraUserTempData{}, err
	}
	return &mitrauser, nil
}

func (p *MitraUserTemp) UpdateMitraUserTemp(db *gorm.DB, pid uint64) (*MitraUserTemp, error) {

	var err error
	db = db.Debug().Model(&MitraUserTemp{}).Where("mitra_user_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_emp_temp_id":          p.MitraEmployeeTempID,
			"mitra_group_accs_temp_id":   p.MitraGroupAccessTempID,
			"mitra_user_temp_email":      p.MitraUserTempEmail,
			"mitra_user_temp_username":   p.MitraUserTempUsername,
			"mitra_user_temp_status":     p.MitraUserTempStatus,
			"mitra_user_temp_reason":     p.MitraUserTempReason,
			"mitra_user_temp_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraUserTemp{}).Where("mitra_user_temp_id = ?", pid).Error
	if err != nil {
		return &MitraUserTemp{}, err
	}
	return p, nil
}

func (p *MitraUserTemp) UpdateMitraUserTempStatus(db *gorm.DB, pid uint64) (*MitraUserTemp, error) {

	var err error
	db = db.Debug().Model(&MitraUserTemp{}).Where("mitra_user_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_user_temp_status":     p.MitraUserTempStatus,
			"mitra_user_temp_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraUserTemp{}).Where("mitra_user_temp_id = ?", pid).Error
	if err != nil {
		return &MitraUserTemp{}, err
	}
	return p, nil
}

func (p *MitraUserTemp) DeleteMitraUserTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraUserTemp{}).Where("mitra_user_temp_id = ?", pid).Take(&MitraUserTemp{}).Delete(&MitraUserTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraUserTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

func HashMitraUser(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPasswordMitraUser(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (p *MitraUserTemp) BeforeSaveMitraUser() error {
	hashedPasswordWeb, err := Hash(p.MitraUserTempPasswordWeb)
	hashedPasswordMobile, err := Hash(p.MitraUserTempPasswordMobile)
	if err != nil {
		return err
	}
	p.MitraUserTempPasswordWeb = string(hashedPasswordWeb)
	p.MitraUserTempPasswordMobile = string(hashedPasswordMobile)
	return nil
}

func (p *MitraUserTempPend) BeforeSaveMitraUser() error {
	hashedPasswordWeb, err := Hash(p.MitraUserTempPasswordWeb)
	hashedPasswordMobile, err := Hash(p.MitraUserTempPasswordMobile)
	hashedPIN, err := Hash(p.MitraUserTempPIN)
	if err != nil {
		return err
	}
	p.MitraUserTempPasswordWeb = string(hashedPasswordWeb)
	p.MitraUserTempPasswordMobile = string(hashedPasswordMobile)
	p.MitraUserTempPIN = string(hashedPIN)
	return nil
}
