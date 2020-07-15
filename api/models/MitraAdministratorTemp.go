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

type MitraAdministratorTemp struct {
	MitraAdministratorTempID           uint64    `gorm:"column:mitra_adm_temp_id;primary_key;" json:"mitra_adm_temp_id"`
	MitraAdministratorID               uint64    `gorm:"column:mitra_adm_id" json:"mitra_adm_id"`
	MitraEmployeeTempID                uint64    `gorm:"column:mitra_emp_temp_id" json:"mitra_emp_temp_id"`
	MitraGroupAccessTempID             uint64    `gorm:"column:mitra_group_accs_temp_id" json:"mitra_group_accs_temp_id"`
	MitraAdministratorTempEmail        string    `gorm:"column:mitra_adm_temp_email;size:255" json:"mitra_adm_temp_email"`
	MitraAdministratorTempUsername     string    `gorm:"column:mitra_adm_temp_username;size:125" json:"mitra_adm_temp_username"`
	MitraAdministratorTempPassword     string    `gorm:"column:mitra_adm_temp_password;size:125" json:"mitra_adm_temp_password"`
	MitraAdministratorTempPhoneCode    string    `gorm:"column:mitra_adm_temp_phone_code;size:6" json:"mitra_adm_temp_phone_code"`
	MitraAdministratorTempPhone        string    `gorm:"column:mitra_adm_temp_phone;size:20" json:"mitra_adm_temp_phone"`
	MitraAdministratorTempPIN          string    `gorm:"column:mitra_adm_temp_pin;size:255" json:"mitra_adm_temp_pin"`
	MitraAdministratorTempReason       string    `gorm:"column:mitra_adm_temp_reason" json:"mitra_adm_temp_reason"`
	MitraAdministratorTempStatus       uint64    `gorm:"column:mitra_adm_temp_status;size:2" json:"mitra_adm_temp_status"`
	MitraAdministratorTempActionBefore uint64    `gorm:"column:mitra_adm_temp_action_before;size:2" json:"mitra_adm_temp_action_before"`
	MitraAdministratorTempCreatedBy    string    `gorm:"column:mitra_adm_temp_created_by;size:125" json:"mitra_adm_temp_created_by"`
	MitraAdministratorTempCreatedAt    time.Time `gorm:"column:mitra_adm_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_adm_temp_created_at"`
}

type MitraAdministratorTempPend struct {
	MitraAdministratorTempID           *uint64   `gorm:"column:mitra_adm_temp_id;primary_key;" json:"mitra_adm_temp_id"`
	MitraAdministratorID               *uint64   `gorm:"column:mitra_adm_id" json:"mitra_adm_id"`
	MitraEmployeeTempID                uint64    `gorm:"column:mitra_emp_temp_id" json:"mitra_emp_temp_id"`
	MitraGroupAccessTempID             uint64    `gorm:"column:mitra_group_accs_temp_id" json:"mitra_group_accs_temp_id"`
	MitraAdministratorTempEmail        string    `gorm:"column:mitra_adm_temp_email;size:255" json:"mitra_adm_temp_email"`
	MitraAdministratorTempUsername     string    `gorm:"column:mitra_adm_temp_username;size:125" json:"mitra_adm_temp_username"`
	MitraAdministratorTempPassword     string    `gorm:"column:mitra_adm_temp_password;size:125" json:"mitra_adm_temp_password"`
	MitraAdministratorTempPhoneCode    string    `gorm:"column:mitra_adm_temp_phone_code;size:6" json:"mitra_adm_temp_phone_code"`
	MitraAdministratorTempPhone        string    `gorm:"column:mitra_adm_temp_phone;size:20" json:"mitra_adm_temp_phone"`
	MitraAdministratorTempPIN          string    `gorm:"column:mitra_adm_temp_pin;size:255" json:"mitra_adm_temp_pin"`
	MitraAdministratorTempReason       string    `gorm:"column:mitra_adm_temp_reason" json:"mitra_adm_temp_reason"`
	MitraAdministratorTempStatus       uint64    `gorm:"column:mitra_adm_temp_status;size:2" json:"mitra_adm_temp_status"`
	MitraAdministratorTempActionBefore uint64    `gorm:"column:mitra_adm_temp_action_before;size:2" json:"mitra_adm_temp_action_before"`
	MitraAdministratorTempCreatedBy    string    `gorm:"column:mitra_adm_temp_created_by;size:125" json:"mitra_adm_temp_created_by"`
	MitraAdministratorTempCreatedAt    time.Time `gorm:"column:mitra_adm_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_adm_temp_created_at"`
}

type MitraAdministratorTempPendData struct {
	MitraAdministratorTempID           uint64    `gorm:"column:mitra_adm_temp_id;primary_key;" json:"mitra_adm_temp_id"`
	MitraAdministratorID               *uint64   `gorm:"column:mitra_adm_id" json:"mitra_adm_id"`
	MitraEmployeeTempID                uint64    `gorm:"column:mitra_emp_temp_id;primary_key;" json:"mitra_emp_temp_id"`
	MitraEmployeeTempName              string    `gorm:"column:mitra_emp_temp_name" json:"mitra_emp_temp_name"`
	MitraTempID                        uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	MitraTempName                      string    `gorm:"column:mitra_temp_name" json:"mitra_temp_name"`
	MitraBranchTempID                  uint64    `gorm:"column:mitra_branch_temp_id" json:"mitra_branch_temp_id"`
	MitraBranchTempName                string    `gorm:"column:mitra_branch_temp_name" json:"mitra_branch_temp_name"`
	MitraGroupAccessTempID             uint64    `gorm:"column:mitra_group_accs_temp_id" json:"mitra_group_accs_temp_id"`
	MitraGroupAccessTempName           string    `gorm:"column:mitra_group_accs_temp_name" json:"mitra_group_accs_temp_name"`
	MitraAdministratorTempEmail        string    `gorm:"column:mitra_adm_temp_email;size:255" json:"mitra_adm_temp_email"`
	MitraAdministratorTempUsername     string    `gorm:"column:mitra_adm_temp_username;size:125" json:"mitra_adm_temp_username"`
	MitraAdministratorTempPassword     string    `gorm:"column:mitra_adm_temp_password;size:125" json:"mitra_adm_temp_password"`
	MitraAdministratorTempPhoneCode    string    `gorm:"column:mitra_adm_temp_phone_code;size:6" json:"mitra_adm_temp_phone_code"`
	MitraAdministratorTempPhone        string    `gorm:"column:mitra_adm_temp_phone;size:20" json:"mitra_adm_temp_phone"`
	MitraAdministratorTempPIN          string    `gorm:"column:mitra_adm_temp_pin;size:255" json:"mitra_adm_temp_pin"`
	MitraAdministratorTempReason       string    `gorm:"column:mitra_adm_temp_reason" json:"mitra_adm_temp_reason"`
	MitraAdministratorTempStatus       uint64    `gorm:"column:mitra_adm_temp_status;size:2" json:"mitra_adm_temp_status"`
	MitraAdministratorTempActionBefore uint64    `gorm:"column:mitra_adm_temp_action_before;size:2" json:"mitra_adm_temp_action_before"`
	MitraAdministratorTempCreatedBy    string    `gorm:"column:mitra_adm_temp_created_by;size:125" json:"mitra_adm_temp_created_by"`
	MitraAdministratorTempCreatedAt    time.Time `gorm:"column:mitra_adm_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_adm_temp_created_at"`
}

type MitraAdministratorTempData struct {
	MitraAdministratorTempID                  uint64    `gorm:"column:mitra_adm_temp_id" json:"mitra_adm_temp_id"`
	MitraEmployeeTempID                       uint64    `gorm:"column:mitra_emp_temp_id" json:"mitra_emp_temp_id"`
	MitraEmployeeTempCode                     string    `gorm:"column:mitra_emp_temp_code;size:25" json:"mitra_emp_temp_code"`
	MitraEmployeeTempName                     string    `gorm:"column:mitra_emp_temp_name;size:255" json:"mitra_emp_temp_name"`
	MitraEmployeeTempGender                   string    `gorm:"column:mitra_emp_temp_gender;size:255" json:"mitra_emp_temp_gender"`
	MitraEmployeeTempBirthPlace               string    `gorm:"column:mitra_emp_temp_birth_place;size:255" json:"mitra_emp_temp_birth_place"`
	MitraEmployeeTempBirthDate                string    `gorm:"column:mitra_emp_temp_birth_date;size:255" json:"mitra_emp_temp_birth_date"`
	MitraEmployeeTempAddress                  string    `gorm:"column:mitra_emp_temp_address;size:255" json:"mitra_emp_temp_address"`
	AddressTempID                             uint64    `gorm:"column:address_temp_id;size:255" json:"address_id"`
	CountryTempID                             uint64    `gorm:"column:country_temp_id" json:"country_temp_id"`
	CountryTempName                           string    `gorm:"column:country_temp_name" json:"country_temp_name"`
	ProvinceTempID                            uint64    `gorm:"column:province_temp_id" json:"province_temp_id"`
	ProvinceTempName                          string    `gorm:"column:province_temp_name" json:"province_temp_name"`
	RegencyTempID                             uint64    `gorm:"column:regency_temp_id" json:"regency_temp_id"`
	RegencyTempName                           string    `gorm:"column:regency_temp_name" json:"regency_temp_name"`
	DistrictTempID                            uint64    `gorm:"column:district_temp_id" json:"district_temp_id"`
	DistrictTempName                          string    `gorm:"column:district_temp_name" json:"district_temp_name"`
	VillageTempID                             uint64    `gorm:"column:village_temp_id" json:"village_temp_id"`
	VillageTempName                           string    `gorm:"column:village_temp_name" json:"village_temp_name"`
	MitraEmployeeTempPhoto                    string    `gorm:"column:mitra_emp_temp_photo;size:255" json:"mitra_emp_temp_photo"`
	MitraGroupAccessTempID                    uint64    `gorm:"column:mitra_group_accs_temp_id" json:"mitra_group_accs_temp_id"`
	MitraGroupAccessTempCode                  string    `gorm:"column:mitra_group_accs_temp_code;size:25" json:"mitra_group_accs_temp_code"`
	MitraGroupAccessTempName                  string    `gorm:"column:mitra_group_accs_temp_name;size:255" json:"mitra_group_accs_temp_name"`
	MitraGroupAccessTempRole                  string    `gorm:"column:mitra_group_accs_temp_role;size:125" json:"mitra_group_accs_temp_role"`
	MitraGroupAccessTempType                  uint64    `gorm:"column:mitra_group_accs_temp_type;size:2" json:"mitra_group_accs_temp_type"`
	MitraTempID                               uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	MitraTempName                             string    `gorm:"column:mitra_temp_name" json:"mitra_temp_name"`
	MitraBranchTempID                         uint64    `gorm:"column:mitra_branch_temp_id" json:"mitra_branch_temp_id"`
	MitraBranchTempName                       string    `gorm:"column:mitra_branch_temp_name" json:"mitra_branch_temp_name"`
	MitraAdministratorTempEmail               string    `gorm:"column:mitra_adm_temp_email;size:255" json:"mitra_adm_temp_email"`
	MitraAdministratorTempUsername            string    `gorm:"column:mitra_adm_temp_username;size:125" json:"mitra_adm_temp_username"`
	MitraAdministratorTempPhoneCode           string    `gorm:"column:mitra_adm_temp_phone_code;size:6" json:"mitra_adm_temp_phone_code"`
	MitraAdministratorTempPhone               string    `gorm:"column:mitra_adm_temp_phone;size:20" json:"mitra_adm_temp_phone"`
	MitraAdministratorTempReason              string    `gorm:"column:mitra_adm_temp_reason" json:"mitra_adm_temp_reason"`
	MitraAdministratorTempStatus              int64     `gorm:"column:mitra_adm_temp_status;size:2" json:"mitra_adm_temp_status"`
	MitraAdministratorTempActionBefore        uint64    `gorm:"column:mitra_adm_temp_action_before;size:2" json:"mitra_adm_temp_action_before"`
	MitraAdministratorTempCreatedBy           string    `gorm:"column:mitra_adm_temp_created_by;size:125" json:"mitra_adm_temp_created_by"`
	MitraAdministratorTempCreatedAt           time.Time `gorm:"column:mitra_adm_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_adm_temp_created_at"`
	MitraAdministratorID                      uint64    `gorm:"column:mitra_adm_id" json:"mitra_adm_id"`
	MitraEmployeeID                           uint64    `gorm:"column:mitra_emp_id" json:"mitra_emp_id"`
	MitraEmployeeCode                         string    `gorm:"column:mitra_emp_code;size:25" json:"mitra_emp_code"`
	MitraEmployeeName                         string    `gorm:"column:mitra_emp_name;size:255" json:"mitra_emp_name"`
	MitraEmployeeGender                       string    `gorm:"column:mitra_emp_gender;size:255" json:"mitra_emp_gender"`
	MitraEmployeeBirthPlace                   string    `gorm:"column:mitra_emp_birth_place;size:255" json:"mitra_emp_birth_place"`
	MitraEmployeeBirthDate                    string    `gorm:"column:mitra_emp_birth_date;size:255" json:"mitra_emp_birth_date"`
	MitraEmployeeAddress                      string    `gorm:"column:mitra_emp_address;size:255" json:"mitra_emp_address"`
	AddressID                                 uint64    `gorm:"column:address_id;size:255" json:"address_id"`
	CountryID                                 uint64    `gorm:"column:country_id" json:"country_id"`
	CountryName                               string    `gorm:"column:country_name" json:"country_name"`
	ProvinceID                                uint64    `gorm:"column:province_id" json:"province_id"`
	ProvinceName                              string    `gorm:"column:province_name" json:"province_name"`
	RegencyID                                 uint64    `gorm:"column:regency_id" json:"regency_id"`
	RegencyName                               string    `gorm:"column:regency_name" json:"regency_name"`
	DistrictID                                uint64    `gorm:"column:district_id" json:"district_id"`
	DistrictName                              string    `gorm:"column:district_name" json:"district_name"`
	VillageID                                 uint64    `gorm:"column:village_id" json:"village_id"`
	VillageName                               string    `gorm:"column:village_name" json:"village_name"`
	MitraEmployeePhoto                        string    `gorm:"column:mitra_emp_photo;size:255" json:"mitra_emp_photo"`
	MitraGroupAccessID                        uint64    `gorm:"column:mitra_group_accs_id" json:"mitra_group_accs_id"`
	MitraGroupAccessCode                      string    `gorm:"column:mitra_group_accs_code;size:25" json:"mitra_group_accs_code"`
	MitraGroupAccessName                      string    `gorm:"column:mitra_group_accs_name;size:255" json:"mitra_group_accs_name"`
	MitraGroupAccessRole                      string    `gorm:"column:mitra_group_accs_role;size:125" json:"mitra_group_accs_role"`
	MitraGroupAccessType                      uint64    `gorm:"column:mitra_group_accs_type;size:2" json:"mitra_group_accs_type"`
	MitraID                                   uint64    `gorm:"column:mitra_id" json:"mitra_id"`
	MitraName                                 string    `gorm:"column:mitra_name" json:"mitra_name"`
	MitraBranchID                             uint64    `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	MitraBranchName                           string    `gorm:"column:mitra_branch_name" json:"mitra_branch_name"`
	MitraAdministratorEmail                   string    `gorm:"column:mitra_adm_email" json:"mitra_adm_email"`
	MitraAdministratorUsername                string    `gorm:"column:mitra_username" json:"mitra_username"`
	MitraAdministratorPhoneCode               string    `gorm:"column:mitra_adm_phone_code;size:6" json:"mitra_adm_phone_code"`
	MitraAdministratorPhone                   string    `gorm:"column:mitra_adm_phone;size:20" json:"mitra_adm_phone"`
	MitraAdministratorStatus                  int64     `gorm:"column:mitra_adm_status;size:1" json:"mitra_adm_status"`
	MitraAdministratorIsNew                   bool      `gorm:"column:mitra_adm_is_new" json:"mitra_adm_is_new"`
	MitraAdministratorIsRequestChangePassword bool      `gorm:"column:mitra_adm_is_req_chpw" json:"mitra_adm_is_req_chpw"`
	MitraAdministratorIsLocked                bool      `gorm:"column:mitra_adm_is_locked" json:"mitra_adm_is_locked"`
	MitraAdministratorFailAttempt             uint64    `gorm:"column:mitra_adm_fail_attempt;size:1" json:"mitra_adm_fail_attempt"`
	MitraAdministratorLockedReason            string    `gorm:"column:mitra_adm_locked_reason;size:255" json:"mitra_adm_locked_reason"`
	MitraAdministratorCreatedBy               string    `gorm:"column:mitra_adm_created_by" json:"mitra_adm_created_by"`
	MitraAdministratorCreatedAt               time.Time `gorm:"column:mitra_adm_created_at" json:"mitra_adm_created_at"`
	MitraAdministratorUpdatedBy               string    `gorm:"column:mitra_adm_updated_by" json:"mitra_adm_updated_by"`
	MitraAdministratorUpdatedAt               time.Time `gorm:"column:mitra_adm_updated_at" json:"mitra_adm_updated_at"`
	MitraAdministratorDeletedBy               string    `gorm:"column:mitra_adm_deleted_by" json:"mitra_adm_deleted_by"`
	MitraAdministratorDeletedAt               time.Time `gorm:"column:mitra_adm_deleted_at" json:"mitra_adm_deleted_at"`
}

type ResponseMitraAdministratorTemps struct {
	Status  int                           `json:"status"`
	Message string                        `json:"message"`
	Data    *[]MitraAdministratorTempData `json:"data"`
}

type ResponseMitraAdministratorTemp struct {
	Status  int                         `json:"status"`
	Message string                      `json:"message"`
	Data    *MitraAdministratorTempData `json:"data"`
}

type ResponseMitraAdministratorTempsPend struct {
	Status  int                               `json:"status"`
	Message string                            `json:"message"`
	Data    *[]MitraAdministratorTempPendData `json:"data"`
}

type ResponseMitraAdministratorTempPend struct {
	Status  int                             `json:"status"`
	Message string                          `json:"message"`
	Data    *MitraAdministratorTempPendData `json:"data"`
}

type ResponseMitraAdministratorTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraAdministratorTemp) TableName() string {
	return "m_mitra_administrator_temp"
}

func (MitraAdministratorTempPend) TableName() string {
	return "m_mitra_administrator_temp"
}

func (MitraAdministratorTempPendData) TableName() string {
	return "m_mitra_administrator_temp"
}

func (MitraAdministratorTempData) TableName() string {
	return "m_mitra_administrator_temp"
}

func (p *MitraAdministratorTemp) Prepare() {
	p.MitraAdministratorID = p.MitraAdministratorID
	p.MitraEmployeeTempID = p.MitraEmployeeTempID
	p.MitraGroupAccessTempID = p.MitraGroupAccessTempID
	p.MitraAdministratorTempEmail = html.EscapeString(strings.TrimSpace(p.MitraAdministratorTempEmail))
	p.MitraAdministratorTempUsername = html.EscapeString(strings.TrimSpace(p.MitraAdministratorTempUsername))
	p.MitraAdministratorTempPassword = html.EscapeString(strings.TrimSpace(p.MitraAdministratorTempPassword))
	p.MitraAdministratorTempPhoneCode = html.EscapeString(strings.TrimSpace(p.MitraAdministratorTempPhoneCode))
	p.MitraAdministratorTempPhone = html.EscapeString(strings.TrimSpace(p.MitraAdministratorTempPhone))
	p.MitraAdministratorTempPIN = html.EscapeString(strings.TrimSpace(p.MitraAdministratorTempPIN))
	p.MitraAdministratorTempPIN = html.EscapeString(strings.TrimSpace(p.MitraAdministratorTempPIN))
	p.MitraAdministratorTempReason = p.MitraAdministratorTempReason
	p.MitraAdministratorTempCreatedBy = p.MitraAdministratorTempCreatedBy
	p.MitraAdministratorTempCreatedAt = time.Now()
}

func (p *MitraAdministratorTempPend) Prepare() {
	p.MitraAdministratorTempID = nil
	p.MitraAdministratorID = nil
	p.MitraEmployeeTempID = p.MitraEmployeeTempID
	p.MitraGroupAccessTempID = p.MitraGroupAccessTempID
	p.MitraAdministratorTempEmail = html.EscapeString(strings.TrimSpace(p.MitraAdministratorTempEmail))
	p.MitraAdministratorTempUsername = html.EscapeString(strings.TrimSpace(p.MitraAdministratorTempUsername))
	p.MitraAdministratorTempPassword = html.EscapeString(strings.TrimSpace(p.MitraAdministratorTempPassword))
	p.MitraAdministratorTempPhoneCode = html.EscapeString(strings.TrimSpace(p.MitraAdministratorTempPhoneCode))
	p.MitraAdministratorTempPhone = html.EscapeString(strings.TrimSpace(p.MitraAdministratorTempPhone))
	p.MitraAdministratorTempPIN = html.EscapeString(strings.TrimSpace(p.MitraAdministratorTempPIN))
	p.MitraAdministratorTempPIN = html.EscapeString(strings.TrimSpace(p.MitraAdministratorTempPIN))
	p.MitraAdministratorTempReason = p.MitraAdministratorTempReason
	p.MitraAdministratorTempCreatedBy = p.MitraAdministratorTempCreatedBy
	p.MitraAdministratorTempCreatedAt = time.Now()
}

func (p *MitraAdministratorTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.MitraAdministratorTempPassword == "" {
			return errors.New("Required MitraAdministratorTemp Password")
		}
		if p.MitraAdministratorTempEmail == "" {
			return errors.New("Required MitraAdministratorTemp Email")
		}
		return nil
	}
}

func (p *MitraAdministratorTemp) SaveMitraAdministratorTemp(db *gorm.DB) (*MitraAdministratorTemp, error) {
	var err error
	err = p.BeforeSaveMitraAdministratorTemp()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Debug().Model(&MitraAdministratorTemp{}).Create(&p).Error
	if err != nil {
		return &MitraAdministratorTemp{}, err
	}
	return p, nil
}

func (p *MitraAdministratorTempPend) SaveMitraAdministratorTempPend(db *gorm.DB) (*MitraAdministratorTempPend, error) {
	var err error
	err = p.BeforeSaveMitraAdministratorTemp()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Debug().Model(&MitraAdministratorTempPend{}).Create(&p).Error
	if err != nil {
		return &MitraAdministratorTempPend{}, err
	}
	return p, nil
}

func (p *MitraAdministratorTemp) FindAllMitraAdministratorTemp(db *gorm.DB) (*[]MitraAdministratorTempPendData, error) {
	var err error
	mitraadministrator := []MitraAdministratorTempPendData{}
	err = db.Debug().Model(&MitraAdministratorTempPendData{}).Limit(100).
		Select(`m_mitra_administrator_temp.mitra_adm_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_name,
				m_mitra_group_access_temp.mitra_group_accs_id as mitra_group_accs_temp_id,
				m_mitra_group_access_temp.mitra_group_accs_name as mitra_group_accs_temp_name,
				m_mitra_administrator_temp.mitra_adm_temp_email,
				m_mitra_administrator_temp.mitra_adm_temp_username,
				m_mitra_administrator_temp.mitra_adm_temp_phone_code,
				m_mitra_administrator_temp.mitra_adm_temp_phone,
				m_mitra_administrator_temp.mitra_adm_temp_reason,
				m_mitra_administrator_temp.mitra_adm_temp_status,
				m_mitra_administrator_temp.mitra_adm_temp_action_before,
				m_mitra_administrator_temp.mitra_adm_temp_created_by,
				m_mitra_administrator_temp.mitra_adm_temp_created_at`).
		Joins("JOIN m_mitra_employee_temp on m_mitra_administrator_temp.mitra_emp_temp_id=m_mitra_employee_temp.mitra_emp_temp_id").
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_employee_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_employee_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_mitra_group_access m_mitra_group_access_temp on m_mitra_administrator_temp.mitra_group_accs_temp_id=m_mitra_group_access_temp.mitra_group_accs_id").
		Find(&mitraadministrator).Error
	if err != nil {
		return &[]MitraAdministratorTempPendData{}, err
	}
	return &mitraadministrator, nil
}

func (p *MitraAdministratorTemp) FindAllMitraAdministratorTempPendingActive(db *gorm.DB) (*[]MitraAdministratorTempPendData, error) {
	var err error
	mitraadministrator := []MitraAdministratorTempPendData{}
	err = db.Debug().Model(&MitraAdministratorTempPendData{}).Limit(100).
		Select(`m_mitra_administrator_temp.mitra_adm_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_name,
				m_mitra_group_access_temp.mitra_group_accs_id as mitra_group_accs_temp_id,
				m_mitra_group_access_temp.mitra_group_accs_name as mitra_group_accs_temp_name,
				m_mitra_administrator_temp.mitra_adm_temp_email,
				m_mitra_administrator_temp.mitra_adm_temp_username,
				m_mitra_administrator_temp.mitra_adm_temp_phone_code,
				m_mitra_administrator_temp.mitra_adm_temp_phone,
				m_mitra_administrator_temp.mitra_adm_temp_reason,
				m_mitra_administrator_temp.mitra_adm_temp_status,
				m_mitra_administrator_temp.mitra_adm_temp_action_before,
				m_mitra_administrator_temp.mitra_adm_temp_created_by,
				m_mitra_administrator_temp.mitra_adm_temp_created_at`).
		Joins("JOIN m_mitra_employee_temp on m_mitra_administrator_temp.mitra_emp_temp_id=m_mitra_employee_temp.mitra_emp_temp_id").
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_employee_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_employee_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_mitra_group_access m_mitra_group_access_temp on m_mitra_administrator_temp.mitra_group_accs_temp_id=m_mitra_group_access_temp.mitra_group_accs_id").
		Where("mitra_adm_temp_status = ?", 11).
		Find(&mitraadministrator).Error
	if err != nil {
		return &[]MitraAdministratorTempPendData{}, err
	}
	return &mitraadministrator, nil
}

func (p *MitraAdministratorTemp) FindAllMitraAdministratorTempStatus(db *gorm.DB, status uint64) (*[]MitraAdministratorTempData, error) {
	var err error
	mitraadministrator := []MitraAdministratorTempData{}
	err = db.Debug().Model(&MitraAdministratorTempData{}).Limit(100).
		Select(`m_mitra_administrator_temp.mitra_adm_temp_id,
				m_mitra_administrator_temp.mitra_emp_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_code,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_gender,
				m_mitra_employee_temp.mitra_emp_temp_birth_place,
				m_mitra_employee_temp.mitra_emp_temp_birth_date,
				m_mitra_employee_temp.mitra_emp_temp_address,
				m_address_temp.address_temp_id,
				m_country_temp.country_id as country_temp_id,
				m_country_temp.country_name as country_temp_name,
				m_province_temp.province_id as province_temp_id,
				m_province_temp.province_name as province_temp_name,
				m_regency_temp.regency_id as regency_temp_id,
				m_regency_temp.regency_name as regency_temp_name,
				m_district_temp.district_id as district_temp_id,
				m_district_temp.district_name as district_temp_name,
				m_village_temp.village_id as village_temp_id,
				m_village_temp.village_name as village_temp_name,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_photo,
				m_mitra_administrator_temp.mitra_group_accs_temp_id,
				m_mitra_group_access_temp.mitra_group_accs_code as mitra_group_accs_temp_code,
				m_mitra_group_access_temp.mitra_group_accs_name as mitra_group_accs_temp_name,
				m_mitra_group_access_temp.mitra_group_accs_role as mitra_group_accs_temp_role,
				m_mitra_group_access_temp.mitra_group_accs_type as mitra_group_accs_temp_type,
				m_mitra_administrator_temp.mitra_adm_temp_email,
				m_mitra_administrator_temp.mitra_adm_temp_username,
				m_mitra_administrator_temp.mitra_adm_temp_phone_code,
				m_mitra_administrator_temp.mitra_adm_temp_phone,
				m_mitra_administrator_temp.mitra_adm_temp_reason,
				m_mitra_administrator_temp.mitra_adm_temp_status,
				m_mitra_administrator_temp.mitra_adm_temp_action_before,
				m_mitra_administrator_temp.mitra_adm_temp_created_by,
				m_mitra_administrator_temp.mitra_adm_temp_created_at,
				m_mitra_administrator.mitra_adm_id,
				m_mitra_administrator.mitra_emp_id,
				m_mitra_employee.mitra_emp_code,
				m_mitra_employee.mitra_emp_name,
				m_mitra_employee.mitra_emp_gender,
				m_mitra_employee.mitra_emp_birth_place,
				m_mitra_employee.mitra_emp_birth_date,
				m_mitra_employee.mitra_emp_address,
				m_address.address_id,
				m_country.country_id,
				m_country.country_name,
				m_province.province_id,
				m_province.province_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_district.district_id,
				m_district.district_name,
				m_village.village_id,
				m_village.village_name,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_employee.mitra_emp_photo,
				m_mitra_group_access.mitra_group_accs_id,
				m_mitra_group_access.mitra_group_accs_code,
				m_mitra_group_access.mitra_group_accs_name,
				m_mitra_group_access.mitra_group_accs_role,
				m_mitra_group_access.mitra_group_accs_type,
				m_mitra_administrator.mitra_adm_email,
				m_mitra_administrator.mitra_adm_username,
				m_mitra_administrator.mitra_adm_phone_code,
				m_mitra_administrator.mitra_adm_phone,
				m_mitra_administrator.mitra_adm_status,
				m_mitra_administrator.mitra_adm_is_new,
				m_mitra_administrator.mitra_adm_is_req_chpw,
				m_mitra_administrator.mitra_adm_is_locked,
				m_mitra_administrator.mitra_adm_fail_attempt,
				m_mitra_administrator.mitra_adm_locked_reason,
				m_mitra_administrator.mitra_adm_created_by,
				m_mitra_administrator.mitra_adm_created_at`).
		Joins("JOIN m_mitra_employee_temp on m_mitra_administrator_temp.mitra_emp_temp_id=m_mitra_employee_temp.mitra_emp_temp_id").
		Joins("JOIN m_address_temp on m_mitra_employee_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_employee_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_employee_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_mitra_group_access m_mitra_group_access_temp on m_mitra_administrator_temp.mitra_group_accs_temp_id=m_mitra_group_access_temp.mitra_group_accs_id").
		Joins("JOIN m_mitra_administrator on m_mitra_administrator_temp.mitra_adm_id=m_mitra_administrator.mitra_adm_id").
		Joins("JOIN m_mitra_employee on m_mitra_administrator.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_address on m_mitra_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra_group_access on m_mitra_administrator.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Where("mitra_adm_temp_status = ?", status).
		Find(&mitraadministrator).Error
	if err != nil {
		return &[]MitraAdministratorTempData{}, err
	}
	return &mitraadministrator, nil
}

func (p *MitraAdministratorTemp) FindMitraAdministratorTempDataByID(db *gorm.DB, pid uint64) (*MitraAdministratorTemp, error) {
	var err error
	err = db.Debug().Model(&MitraAdministratorTemp{}).Where("mitra_adm_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraAdministratorTemp{}, err
	}
	return p, nil
}

func (p *MitraAdministratorTemp) FindMitraAdministratorTempByIDPendingActive(db *gorm.DB, pid uint64) (*MitraAdministratorTempPendData, error) {
	var err error
	mitraadministrator := MitraAdministratorTempPendData{}
	err = db.Debug().Model(&MitraAdministratorTempPendData{}).Limit(100).
		Select(`m_mitra_administrator_temp.mitra_adm_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_name,
				m_mitra_group_access_temp.mitra_group_accs_id as mitra_group_accs_temp_id,
				m_mitra_group_access_temp.mitra_group_accs_name as mitra_group_accs_temp_name,
				m_mitra_administrator_temp.mitra_adm_temp_email,
				m_mitra_administrator_temp.mitra_adm_temp_username,
				m_mitra_administrator_temp.mitra_adm_temp_password,
				m_mitra_administrator_temp.mitra_adm_temp_phone_code,
				m_mitra_administrator_temp.mitra_adm_temp_phone,
				m_mitra_administrator_temp.mitra_adm_temp_reason,
				m_mitra_administrator_temp.mitra_adm_temp_status,
				m_mitra_administrator_temp.mitra_adm_temp_action_before,
				m_mitra_administrator_temp.mitra_adm_temp_created_by,
				m_mitra_administrator_temp.mitra_adm_temp_created_at`).
		Joins("JOIN m_mitra_employee_temp on m_mitra_administrator_temp.mitra_emp_temp_id=m_mitra_employee_temp.mitra_emp_temp_id").
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_employee_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_employee_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_mitra_group_access m_mitra_group_access_temp on m_mitra_administrator_temp.mitra_group_accs_temp_id=m_mitra_group_access_temp.mitra_group_accs_id").
		Where("mitra_adm_temp_id = ? AND mitra_adm_temp_status = ?", pid, 11).
		Find(&mitraadministrator).Error
	if err != nil {
		return &MitraAdministratorTempPendData{}, err
	}
	return &mitraadministrator, nil
}

func (p *MitraAdministratorTemp) FindMitraAdministratorTempByID(db *gorm.DB, pid uint64) (*MitraAdministratorTempData, error) {
	var err error
	mitraadministrator := MitraAdministratorTempData{}
	err = db.Debug().Model(&MitraAdministratorTempData{}).Limit(100).
		Select(`m_mitra_administrator_temp.mitra_adm_temp_id,
				m_mitra_administrator_temp.mitra_emp_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_code,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_gender,
				m_mitra_employee_temp.mitra_emp_temp_birth_place,
				m_mitra_employee_temp.mitra_emp_temp_birth_date,
				m_mitra_employee_temp.mitra_emp_temp_address,
				m_address_temp.address_temp_id,
				m_country_temp.country_id as country_temp_id,
				m_country_temp.country_name as country_temp_name,
				m_province_temp.province_id as province_temp_id,
				m_province_temp.province_name as province_temp_name,
				m_regency_temp.regency_id as regency_temp_id,
				m_regency_temp.regency_name as regency_temp_name,
				m_district_temp.district_id as district_temp_id,
				m_district_temp.district_name as district_temp_name,
				m_village_temp.village_id as village_temp_id,
				m_village_temp.village_name as village_temp_name,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_photo,
				m_mitra_administrator_temp.mitra_group_accs_temp_id,
				m_mitra_group_access_temp.mitra_group_accs_code as mitra_group_accs_temp_code,
				m_mitra_group_access_temp.mitra_group_accs_name as mitra_group_accs_temp_name,
				m_mitra_group_access_temp.mitra_group_accs_role as mitra_group_accs_temp_role,
				m_mitra_group_access_temp.mitra_group_accs_type as mitra_group_accs_temp_type,
				m_mitra_administrator_temp.mitra_adm_temp_email,
				m_mitra_administrator_temp.mitra_adm_temp_username,
				m_mitra_administrator_temp.mitra_adm_temp_phone_code,
				m_mitra_administrator_temp.mitra_adm_temp_phone,
				m_mitra_administrator_temp.mitra_adm_temp_reason,
				m_mitra_administrator_temp.mitra_adm_temp_status,
				m_mitra_administrator_temp.mitra_adm_temp_action_before,
				m_mitra_administrator_temp.mitra_adm_temp_created_by,
				m_mitra_administrator_temp.mitra_adm_temp_created_at,
				m_mitra_administrator.mitra_adm_id,
				m_mitra_administrator.mitra_emp_id,
				m_mitra_employee.mitra_emp_code,
				m_mitra_employee.mitra_emp_name,
				m_mitra_employee.mitra_emp_gender,
				m_mitra_employee.mitra_emp_birth_place,
				m_mitra_employee.mitra_emp_birth_date,
				m_mitra_employee.mitra_emp_address,
				m_address.address_id,
				m_country.country_id,
				m_country.country_name,
				m_province.province_id,
				m_province.province_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_district.district_id,
				m_district.district_name,
				m_village.village_id,
				m_village.village_name,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_employee.mitra_emp_photo,
				m_mitra_group_access.mitra_group_accs_id,
				m_mitra_group_access.mitra_group_accs_code,
				m_mitra_group_access.mitra_group_accs_name,
				m_mitra_group_access.mitra_group_accs_role,
				m_mitra_group_access.mitra_group_accs_type,
				m_mitra_administrator.mitra_adm_email,
				m_mitra_administrator.mitra_adm_username,
				m_mitra_administrator.mitra_adm_phone_code,
				m_mitra_administrator.mitra_adm_phone,
				m_mitra_administrator.mitra_adm_status,
				m_mitra_administrator.mitra_adm_is_new,
				m_mitra_administrator.mitra_adm_is_req_chpw,
				m_mitra_administrator.mitra_adm_is_locked,
				m_mitra_administrator.mitra_adm_fail_attempt,
				m_mitra_administrator.mitra_adm_locked_reason,
				m_mitra_administrator.mitra_adm_created_by,
				m_mitra_administrator.mitra_adm_created_at`).
		Joins("JOIN m_mitra_employee_temp on m_mitra_administrator_temp.mitra_emp_temp_id=m_mitra_employee_temp.mitra_emp_temp_id").
		Joins("JOIN m_address_temp on m_mitra_employee_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_employee_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_employee_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_mitra_group_access m_mitra_group_access_temp on m_mitra_administrator_temp.mitra_group_accs_temp_id=m_mitra_group_access_temp.mitra_group_accs_id").
		Joins("JOIN m_mitra_administrator on m_mitra_administrator_temp.mitra_adm_id=m_mitra_administrator.mitra_adm_id").
		Joins("JOIN m_mitra_employee on m_mitra_administrator.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_address on m_mitra_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra_group_access on m_mitra_administrator.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Where("mitra_adm_temp_id = ?", pid).
		Find(&mitraadministrator).Error
	if err != nil {
		return &MitraAdministratorTempData{}, err
	}
	return &mitraadministrator, nil
}

func (p *MitraAdministratorTemp) FindMitraAdministratorTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraAdministratorTempData, error) {
	var err error
	mitraadministrator := MitraAdministratorTempData{}
	err = db.Debug().Model(&MitraAdministratorTempData{}).Limit(100).
		Select(`m_mitra_administrator_temp.mitra_adm_temp_id,
				m_mitra_administrator_temp.mitra_emp_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_code,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_gender,
				m_mitra_employee_temp.mitra_emp_temp_birth_place,
				m_mitra_employee_temp.mitra_emp_temp_birth_date,
				m_mitra_employee_temp.mitra_emp_temp_address,
				m_address_temp.address_temp_id,
				m_country_temp.country_id as country_temp_id,
				m_country_temp.country_name as country_temp_name,
				m_province_temp.province_id as province_temp_id,
				m_province_temp.province_name as province_temp_name,
				m_regency_temp.regency_id as regency_temp_id,
				m_regency_temp.regency_name as regency_temp_name,
				m_district_temp.district_id as district_temp_id,
				m_district_temp.district_name as district_temp_name,
				m_village_temp.village_id as village_temp_id,
				m_village_temp.village_name as village_temp_name,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_photo,
				m_mitra_administrator_temp.mitra_group_accs_temp_id,
				m_mitra_group_access_temp.mitra_group_accs_code as mitra_group_accs_temp_code,
				m_mitra_group_access_temp.mitra_group_accs_name as mitra_group_accs_temp_name,
				m_mitra_group_access_temp.mitra_group_accs_role as mitra_group_accs_temp_role,
				m_mitra_group_access_temp.mitra_group_accs_type as mitra_group_accs_temp_type,
				m_mitra_administrator_temp.mitra_adm_temp_email,
				m_mitra_administrator_temp.mitra_adm_temp_username,
				m_mitra_administrator_temp.mitra_adm_temp_phone_code,
				m_mitra_administrator_temp.mitra_adm_temp_phone,
				m_mitra_administrator_temp.mitra_adm_temp_reason,
				m_mitra_administrator_temp.mitra_adm_temp_status,
				m_mitra_administrator_temp.mitra_adm_temp_action_before,
				m_mitra_administrator_temp.mitra_adm_temp_created_by,
				m_mitra_administrator_temp.mitra_adm_temp_created_at,
				m_mitra_administrator.mitra_adm_id,
				m_mitra_administrator.mitra_emp_id,
				m_mitra_employee.mitra_emp_code,
				m_mitra_employee.mitra_emp_name,
				m_mitra_employee.mitra_emp_gender,
				m_mitra_employee.mitra_emp_birth_place,
				m_mitra_employee.mitra_emp_birth_date,
				m_mitra_employee.mitra_emp_address,
				m_address.address_id,
				m_country.country_id,
				m_country.country_name,
				m_province.province_id,
				m_province.province_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_district.district_id,
				m_district.district_name,
				m_village.village_id,
				m_village.village_name,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_employee.mitra_emp_photo,
				m_mitra_group_access.mitra_group_accs_id,
				m_mitra_group_access.mitra_group_accs_code,
				m_mitra_group_access.mitra_group_accs_name,
				m_mitra_group_access.mitra_group_accs_role,
				m_mitra_group_access.mitra_group_accs_type,
				m_mitra_administrator.mitra_adm_email,
				m_mitra_administrator.mitra_adm_username,
				m_mitra_administrator.mitra_adm_phone_code,
				m_mitra_administrator.mitra_adm_phone,
				m_mitra_administrator.mitra_adm_status,
				m_mitra_administrator.mitra_adm_is_new,
				m_mitra_administrator.mitra_adm_is_req_chpw,
				m_mitra_administrator.mitra_adm_is_locked,
				m_mitra_administrator.mitra_adm_fail_attempt,
				m_mitra_administrator.mitra_adm_locked_reason,
				m_mitra_administrator.mitra_adm_created_by,
				m_mitra_administrator.mitra_adm_created_at`).
		Joins("JOIN m_mitra_employee_temp on m_mitra_administrator_temp.mitra_emp_temp_id=m_mitra_employee_temp.mitra_emp_temp_id").
		Joins("JOIN m_address_temp on m_mitra_employee_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_employee_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_employee_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_mitra_group_access m_mitra_group_access_temp on m_mitra_administrator_temp.mitra_group_accs_temp_id=m_mitra_group_access_temp.mitra_group_accs_id").
		Joins("JOIN m_mitra_administrator on m_mitra_administrator_temp.mitra_adm_id=m_mitra_administrator.mitra_adm_id").
		Joins("JOIN m_mitra_employee on m_mitra_administrator.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_address on m_mitra_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra_group_access on m_mitra_administrator.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Where("mitra_adm_temp_id = ? AND mitra_adm_temp_status = ?", pid, status).
		Find(&mitraadministrator).Error
	if err != nil {
		return &MitraAdministratorTempData{}, err
	}
	return &mitraadministrator, nil
}

func (p *MitraAdministratorTemp) UpdateMitraAdministratorTemp(db *gorm.DB, pid uint64) (*MitraAdministratorTemp, error) {
	var err error
	db = db.Debug().Model(&MitraAdministratorTemp{}).Where("mitra_adm_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_emp_temp_id":        p.MitraEmployeeTempID,
			"mitra_group_accs_temp_id": p.MitraGroupAccessTempID,
			"mitra_adm_temp_email":     p.MitraAdministratorTempEmail,
			"mitra_adm_temp_username":  p.MitraAdministratorTempUsername,
			"mitra_adm_temp_status":    p.MitraAdministratorTempStatus,
			"mitra_adm_temp_reason":    p.MitraAdministratorTempReason,
		},
	)
	err = db.Debug().Model(&MitraAdministratorTemp{}).Where("mitra_adm_temp_id = ?", pid).Error
	if err != nil {
		return &MitraAdministratorTemp{}, err
	}
	return p, nil
}

func (p *MitraAdministratorTemp) UpdateMitraAdministratorTempStatus(db *gorm.DB, pid uint64) (*MitraAdministratorTemp, error) {
	var err error
	db = db.Debug().Model(&MitraAdministratorTemp{}).Where("mitra_adm_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_adm_temp_reason": p.MitraAdministratorTempReason,
			"mitra_adm_temp_status": p.MitraAdministratorTempStatus,
		},
	)
	err = db.Debug().Model(&MitraAdministratorTemp{}).Where("mitra_adm_temp_id = ?", pid).Error
	if err != nil {
		return &MitraAdministratorTemp{}, err
	}
	return p, nil
}

func (p *MitraAdministratorTemp) DeleteMitraAdministratorTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraAdministratorTemp{}).Where("mitra_adm_temp_id = ?", pid).Take(&MitraAdministratorTemp{}).Delete(&MitraAdministratorTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraAdministratorTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//GET BY MITRA ID
//===================================================================================================================================

func (p *MitraAdministratorTemp) FindAllMitraAdministratorTempMitra(db *gorm.DB, mitra uint64) (*[]MitraAdministratorTempPendData, error) {
	var err error
	mitraadministrator := []MitraAdministratorTempPendData{}
	err = db.Debug().Model(&MitraAdministratorTempPendData{}).Limit(100).
		Select(`m_mitra_administrator_temp.mitra_adm_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_name,
				m_mitra_group_access_temp.mitra_group_accs_id as mitra_group_accs_temp_id,
				m_mitra_group_access_temp.mitra_group_accs_name as mitra_group_accs_temp_name,
				m_mitra_administrator_temp.mitra_adm_temp_email,
				m_mitra_administrator_temp.mitra_adm_temp_username,
				m_mitra_administrator_temp.mitra_adm_temp_phone_code,
				m_mitra_administrator_temp.mitra_adm_temp_phone,
				m_mitra_administrator_temp.mitra_adm_temp_reason,
				m_mitra_administrator_temp.mitra_adm_temp_status,
				m_mitra_administrator_temp.mitra_adm_temp_action_before,
				m_mitra_administrator_temp.mitra_adm_temp_created_by,
				m_mitra_administrator_temp.mitra_adm_temp_created_at`).
		Joins("JOIN m_mitra_employee_temp on m_mitra_administrator_temp.mitra_emp_temp_id=m_mitra_employee_temp.mitra_emp_temp_id").
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_employee_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_employee_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_mitra_group_access m_mitra_group_access_temp on m_mitra_administrator_temp.mitra_group_accs_temp_id=m_mitra_group_access_temp.mitra_group_accs_id").
		Where("m_mitra_employee_temp.mitra_emp_temp_id = ?", mitra).
		Find(&mitraadministrator).Error
	if err != nil {
		return &[]MitraAdministratorTempPendData{}, err
	}
	return &mitraadministrator, nil
}

func (p *MitraAdministratorTemp) FindAllMitraAdministratorTempPendingActiveMitra(db *gorm.DB, mitra uint64) (*[]MitraAdministratorTempPendData, error) {
	var err error
	mitraadministrator := []MitraAdministratorTempPendData{}
	err = db.Debug().Model(&MitraAdministratorTempPendData{}).Limit(100).
		Select(`m_mitra_administrator_temp.mitra_adm_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_name,
				m_mitra_group_access_temp.mitra_group_accs_id as mitra_group_accs_temp_id,
				m_mitra_group_access_temp.mitra_group_accs_name as mitra_group_accs_temp_name,
				m_mitra_administrator_temp.mitra_adm_temp_email,
				m_mitra_administrator_temp.mitra_adm_temp_username,
				m_mitra_administrator_temp.mitra_adm_temp_phone_code,
				m_mitra_administrator_temp.mitra_adm_temp_phone,
				m_mitra_administrator_temp.mitra_adm_temp_reason,
				m_mitra_administrator_temp.mitra_adm_temp_status,
				m_mitra_administrator_temp.mitra_adm_temp_action_before,
				m_mitra_administrator_temp.mitra_adm_temp_created_by,
				m_mitra_administrator_temp.mitra_adm_temp_created_at`).
		Joins("JOIN m_mitra_employee_temp on m_mitra_administrator_temp.mitra_emp_temp_id=m_mitra_employee_temp.mitra_emp_temp_id").
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_employee_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_employee_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_mitra_group_access m_mitra_group_access_temp on m_mitra_administrator_temp.mitra_group_accs_temp_id=m_mitra_group_access_temp.mitra_group_accs_id").
		Where("mitra_adm_temp_status = ? AND m_mitra_employee_temp.mitra_emp_temp_id = ?", 11, mitra).
		Find(&mitraadministrator).Error
	if err != nil {
		return &[]MitraAdministratorTempPendData{}, err
	}
	return &mitraadministrator, nil
}

func (p *MitraAdministratorTemp) FindAllMitraAdministratorTempStatusMitra(db *gorm.DB, status uint64, mitra uint64) (*[]MitraAdministratorTempData, error) {
	var err error
	mitraadministrator := []MitraAdministratorTempData{}
	err = db.Debug().Model(&MitraAdministratorTempData{}).Limit(100).
		Select(`m_mitra_administrator_temp.mitra_adm_temp_id,
				m_mitra_administrator_temp.mitra_emp_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_code,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_gender,
				m_mitra_employee_temp.mitra_emp_temp_birth_place,
				m_mitra_employee_temp.mitra_emp_temp_birth_date,
				m_mitra_employee_temp.mitra_emp_temp_address,
				m_address_temp.address_temp_id,
				m_country_temp.country_id as country_temp_id,
				m_country_temp.country_name as country_temp_name,
				m_province_temp.province_id as province_temp_id,
				m_province_temp.province_name as province_temp_name,
				m_regency_temp.regency_id as regency_temp_id,
				m_regency_temp.regency_name as regency_temp_name,
				m_district_temp.district_id as district_temp_id,
				m_district_temp.district_name as district_temp_name,
				m_village_temp.village_id as village_temp_id,
				m_village_temp.village_name as village_temp_name,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_photo,
				m_mitra_administrator_temp.mitra_group_accs_temp_id,
				m_mitra_group_access_temp.mitra_group_accs_code as mitra_group_accs_temp_code,
				m_mitra_group_access_temp.mitra_group_accs_name as mitra_group_accs_temp_name,
				m_mitra_group_access_temp.mitra_group_accs_role as mitra_group_accs_temp_role,
				m_mitra_group_access_temp.mitra_group_accs_type as mitra_group_accs_temp_type,
				m_mitra_administrator_temp.mitra_adm_temp_email,
				m_mitra_administrator_temp.mitra_adm_temp_username,
				m_mitra_administrator_temp.mitra_adm_temp_phone_code,
				m_mitra_administrator_temp.mitra_adm_temp_phone,
				m_mitra_administrator_temp.mitra_adm_temp_reason,
				m_mitra_administrator_temp.mitra_adm_temp_status,
				m_mitra_administrator_temp.mitra_adm_temp_action_before,
				m_mitra_administrator_temp.mitra_adm_temp_created_by,
				m_mitra_administrator_temp.mitra_adm_temp_created_at,
				m_mitra_administrator.mitra_adm_id,
				m_mitra_administrator.mitra_emp_id,
				m_mitra_employee.mitra_emp_code,
				m_mitra_employee.mitra_emp_name,
				m_mitra_employee.mitra_emp_gender,
				m_mitra_employee.mitra_emp_birth_place,
				m_mitra_employee.mitra_emp_birth_date,
				m_mitra_employee.mitra_emp_address,
				m_address.address_id,
				m_country.country_id,
				m_country.country_name,
				m_province.province_id,
				m_province.province_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_district.district_id,
				m_district.district_name,
				m_village.village_id,
				m_village.village_name,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_employee.mitra_emp_photo,
				m_mitra_group_access.mitra_group_accs_id,
				m_mitra_group_access.mitra_group_accs_code,
				m_mitra_group_access.mitra_group_accs_name,
				m_mitra_group_access.mitra_group_accs_role,
				m_mitra_group_access.mitra_group_accs_type,
				m_mitra_administrator.mitra_adm_email,
				m_mitra_administrator.mitra_adm_username,
				m_mitra_administrator.mitra_adm_phone_code,
				m_mitra_administrator.mitra_adm_phone,
				m_mitra_administrator.mitra_adm_status,
				m_mitra_administrator.mitra_adm_is_new,
				m_mitra_administrator.mitra_adm_is_req_chpw,
				m_mitra_administrator.mitra_adm_is_locked,
				m_mitra_administrator.mitra_adm_fail_attempt,
				m_mitra_administrator.mitra_adm_locked_reason,
				m_mitra_administrator.mitra_adm_created_by,
				m_mitra_administrator.mitra_adm_created_at`).
		Joins("JOIN m_mitra_employee_temp on m_mitra_administrator_temp.mitra_emp_temp_id=m_mitra_employee_temp.mitra_emp_temp_id").
		Joins("JOIN m_address_temp on m_mitra_employee_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_employee_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_employee_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_mitra_group_access m_mitra_group_access_temp on m_mitra_administrator_temp.mitra_group_accs_temp_id=m_mitra_group_access_temp.mitra_group_accs_id").
		Joins("JOIN m_mitra_administrator on m_mitra_administrator_temp.mitra_adm_id=m_mitra_administrator.mitra_adm_id").
		Joins("JOIN m_mitra_employee on m_mitra_administrator.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_address on m_mitra_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra_group_access on m_mitra_administrator.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Where("mitra_adm_temp_status = ? AND m_mitra_employee_temp.mitra_temp_id = ?", status, mitra).
		Find(&mitraadministrator).Error
	if err != nil {
		return &[]MitraAdministratorTempData{}, err
	}
	return &mitraadministrator, nil
}

func (p *MitraAdministratorTemp) FindMitraAdministratorTempDataByIDMitra(db *gorm.DB, pid uint64, mitra uint64) (*MitraAdministratorTemp, error) {
	var err error
	err = db.Debug().Model(&MitraAdministratorTemp{}).Where("mitra_adm_temp_id = ? AND m_mitra_employee_temp.mitra_emp_temp_id = ?", pid, mitra).Take(&p).Error
	if err != nil {
		return &MitraAdministratorTemp{}, err
	}
	return p, nil
}

func (p *MitraAdministratorTemp) FindMitraAdministratorTempByIDPendingActiveMitra(db *gorm.DB, pid uint64, mitra uint64) (*MitraAdministratorTempPendData, error) {
	var err error
	mitraadministrator := MitraAdministratorTempPendData{}
	err = db.Debug().Model(&MitraAdministratorTempPendData{}).Limit(100).
		Select(`m_mitra_administrator_temp.mitra_adm_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_name,
				m_mitra_group_access_temp.mitra_group_accs_id as mitra_group_accs_temp_id,
				m_mitra_group_access_temp.mitra_group_accs_name as mitra_group_accs_temp_name,
				m_mitra_administrator_temp.mitra_adm_temp_email,
				m_mitra_administrator_temp.mitra_adm_temp_username,
				m_mitra_administrator_temp.mitra_adm_temp_password,
				m_mitra_administrator_temp.mitra_adm_temp_phone_code,
				m_mitra_administrator_temp.mitra_adm_temp_phone,
				m_mitra_administrator_temp.mitra_adm_temp_reason,
				m_mitra_administrator_temp.mitra_adm_temp_status,
				m_mitra_administrator_temp.mitra_adm_temp_action_before,
				m_mitra_administrator_temp.mitra_adm_temp_created_by,
				m_mitra_administrator_temp.mitra_adm_temp_created_at`).
		Joins("JOIN m_mitra_employee_temp on m_mitra_administrator_temp.mitra_emp_temp_id=m_mitra_employee_temp.mitra_emp_temp_id").
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_employee_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_employee_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_mitra_group_access m_mitra_group_access_temp on m_mitra_administrator_temp.mitra_group_accs_temp_id=m_mitra_group_access_temp.mitra_group_accs_id").
		Where("mitra_adm_temp_id = ? AND mitra_adm_temp_status = ? AND m_mitra_employee_temp.mitra_emp_temp_id = ?", pid, 11, mitra).
		Find(&mitraadministrator).Error
	if err != nil {
		return &MitraAdministratorTempPendData{}, err
	}
	return &mitraadministrator, nil
}

func (p *MitraAdministratorTemp) FindMitraAdministratorTempByIDMitra(db *gorm.DB, pid uint64, mitra uint64) (*MitraAdministratorTempData, error) {
	var err error
	mitraadministrator := MitraAdministratorTempData{}
	err = db.Debug().Model(&MitraAdministratorTempData{}).Limit(100).
		Select(`m_mitra_administrator_temp.mitra_adm_temp_id,
				m_mitra_administrator_temp.mitra_emp_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_code,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_gender,
				m_mitra_employee_temp.mitra_emp_temp_birth_place,
				m_mitra_employee_temp.mitra_emp_temp_birth_date,
				m_mitra_employee_temp.mitra_emp_temp_address,
				m_address_temp.address_temp_id,
				m_country_temp.country_id as country_temp_id,
				m_country_temp.country_name as country_temp_name,
				m_province_temp.province_id as province_temp_id,
				m_province_temp.province_name as province_temp_name,
				m_regency_temp.regency_id as regency_temp_id,
				m_regency_temp.regency_name as regency_temp_name,
				m_district_temp.district_id as district_temp_id,
				m_district_temp.district_name as district_temp_name,
				m_village_temp.village_id as village_temp_id,
				m_village_temp.village_name as village_temp_name,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_photo,
				m_mitra_administrator_temp.mitra_group_accs_temp_id,
				m_mitra_group_access_temp.mitra_group_accs_code as mitra_group_accs_temp_code,
				m_mitra_group_access_temp.mitra_group_accs_name as mitra_group_accs_temp_name,
				m_mitra_group_access_temp.mitra_group_accs_role as mitra_group_accs_temp_role,
				m_mitra_group_access_temp.mitra_group_accs_type as mitra_group_accs_temp_type,
				m_mitra_administrator_temp.mitra_adm_temp_email,
				m_mitra_administrator_temp.mitra_adm_temp_username,
				m_mitra_administrator_temp.mitra_adm_temp_phone_code,
				m_mitra_administrator_temp.mitra_adm_temp_phone,
				m_mitra_administrator_temp.mitra_adm_temp_reason,
				m_mitra_administrator_temp.mitra_adm_temp_status,
				m_mitra_administrator_temp.mitra_adm_temp_action_before,
				m_mitra_administrator_temp.mitra_adm_temp_created_by,
				m_mitra_administrator_temp.mitra_adm_temp_created_at,
				m_mitra_administrator.mitra_adm_id,
				m_mitra_administrator.mitra_emp_id,
				m_mitra_employee.mitra_emp_code,
				m_mitra_employee.mitra_emp_name,
				m_mitra_employee.mitra_emp_gender,
				m_mitra_employee.mitra_emp_birth_place,
				m_mitra_employee.mitra_emp_birth_date,
				m_mitra_employee.mitra_emp_address,
				m_address.address_id,
				m_country.country_id,
				m_country.country_name,
				m_province.province_id,
				m_province.province_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_district.district_id,
				m_district.district_name,
				m_village.village_id,
				m_village.village_name,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_employee.mitra_emp_photo,
				m_mitra_group_access.mitra_group_accs_id,
				m_mitra_group_access.mitra_group_accs_code,
				m_mitra_group_access.mitra_group_accs_name,
				m_mitra_group_access.mitra_group_accs_role,
				m_mitra_group_access.mitra_group_accs_type,
				m_mitra_administrator.mitra_adm_email,
				m_mitra_administrator.mitra_adm_username,
				m_mitra_administrator.mitra_adm_phone_code,
				m_mitra_administrator.mitra_adm_phone,
				m_mitra_administrator.mitra_adm_status,
				m_mitra_administrator.mitra_adm_is_new,
				m_mitra_administrator.mitra_adm_is_req_chpw,
				m_mitra_administrator.mitra_adm_is_locked,
				m_mitra_administrator.mitra_adm_fail_attempt,
				m_mitra_administrator.mitra_adm_locked_reason,
				m_mitra_administrator.mitra_adm_created_by,
				m_mitra_administrator.mitra_adm_created_at`).
		Joins("JOIN m_mitra_employee_temp on m_mitra_administrator_temp.mitra_emp_temp_id=m_mitra_employee_temp.mitra_emp_temp_id").
		Joins("JOIN m_address_temp on m_mitra_employee_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_employee_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_employee_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_mitra_group_access m_mitra_group_access_temp on m_mitra_administrator_temp.mitra_group_accs_temp_id=m_mitra_group_access_temp.mitra_group_accs_id").
		Joins("JOIN m_mitra_administrator on m_mitra_administrator_temp.mitra_adm_id=m_mitra_administrator.mitra_adm_id").
		Joins("JOIN m_mitra_employee on m_mitra_administrator.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_address on m_mitra_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra_group_access on m_mitra_administrator.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Where("mitra_adm_temp_id = ? AND m_mitra_employee_temp.mitra_emp_temp_id = ?", pid, mitra).
		Find(&mitraadministrator).Error
	if err != nil {
		return &MitraAdministratorTempData{}, err
	}
	return &mitraadministrator, nil
}

func (p *MitraAdministratorTemp) FindMitraAdministratorTempStatusByIDMitra(db *gorm.DB, pid uint64, status uint64, mitra uint64) (*MitraAdministratorTempData, error) {
	var err error
	mitraadministrator := MitraAdministratorTempData{}
	err = db.Debug().Model(&MitraAdministratorTempData{}).Limit(100).
		Select(`m_mitra_administrator_temp.mitra_adm_temp_id,
				m_mitra_administrator_temp.mitra_emp_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_code,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_gender,
				m_mitra_employee_temp.mitra_emp_temp_birth_place,
				m_mitra_employee_temp.mitra_emp_temp_birth_date,
				m_mitra_employee_temp.mitra_emp_temp_address,
				m_address_temp.address_temp_id,
				m_country_temp.country_id as country_temp_id,
				m_country_temp.country_name as country_temp_name,
				m_province_temp.province_id as province_temp_id,
				m_province_temp.province_name as province_temp_name,
				m_regency_temp.regency_id as regency_temp_id,
				m_regency_temp.regency_name as regency_temp_name,
				m_district_temp.district_id as district_temp_id,
				m_district_temp.district_name as district_temp_name,
				m_village_temp.village_id as village_temp_id,
				m_village_temp.village_name as village_temp_name,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_photo,
				m_mitra_administrator_temp.mitra_group_accs_temp_id,
				m_mitra_group_access_temp.mitra_group_accs_code as mitra_group_accs_temp_code,
				m_mitra_group_access_temp.mitra_group_accs_name as mitra_group_accs_temp_name,
				m_mitra_group_access_temp.mitra_group_accs_role as mitra_group_accs_temp_role,
				m_mitra_group_access_temp.mitra_group_accs_type as mitra_group_accs_temp_type,
				m_mitra_administrator_temp.mitra_adm_temp_email,
				m_mitra_administrator_temp.mitra_adm_temp_username,
				m_mitra_administrator_temp.mitra_adm_temp_phone_code,
				m_mitra_administrator_temp.mitra_adm_temp_phone,
				m_mitra_administrator_temp.mitra_adm_temp_reason,
				m_mitra_administrator_temp.mitra_adm_temp_status,
				m_mitra_administrator_temp.mitra_adm_temp_action_before,
				m_mitra_administrator_temp.mitra_adm_temp_created_by,
				m_mitra_administrator_temp.mitra_adm_temp_created_at,
				m_mitra_administrator.mitra_adm_id,
				m_mitra_administrator.mitra_emp_id,
				m_mitra_employee.mitra_emp_code,
				m_mitra_employee.mitra_emp_name,
				m_mitra_employee.mitra_emp_gender,
				m_mitra_employee.mitra_emp_birth_place,
				m_mitra_employee.mitra_emp_birth_date,
				m_mitra_employee.mitra_emp_address,
				m_address.address_id,
				m_country.country_id,
				m_country.country_name,
				m_province.province_id,
				m_province.province_name,
				m_regency.regency_id,
				m_regency.regency_name,
				m_district.district_id,
				m_district.district_name,
				m_village.village_id,
				m_village.village_name,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_employee.mitra_emp_photo,
				m_mitra_group_access.mitra_group_accs_id,
				m_mitra_group_access.mitra_group_accs_code,
				m_mitra_group_access.mitra_group_accs_name,
				m_mitra_group_access.mitra_group_accs_role,
				m_mitra_group_access.mitra_group_accs_type,
				m_mitra_administrator.mitra_adm_email,
				m_mitra_administrator.mitra_adm_username,
				m_mitra_administrator.mitra_adm_phone_code,
				m_mitra_administrator.mitra_adm_phone,
				m_mitra_administrator.mitra_adm_status,
				m_mitra_administrator.mitra_adm_is_new,
				m_mitra_administrator.mitra_adm_is_req_chpw,
				m_mitra_administrator.mitra_adm_is_locked,
				m_mitra_administrator.mitra_adm_fail_attempt,
				m_mitra_administrator.mitra_adm_locked_reason,
				m_mitra_administrator.mitra_adm_created_by,
				m_mitra_administrator.mitra_adm_created_at`).
		Joins("JOIN m_mitra_employee_temp on m_mitra_administrator_temp.mitra_emp_temp_id=m_mitra_employee_temp.mitra_emp_temp_id").
		Joins("JOIN m_address_temp on m_mitra_employee_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_employee_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_employee_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_mitra_group_access m_mitra_group_access_temp on m_mitra_administrator_temp.mitra_group_accs_temp_id=m_mitra_group_access_temp.mitra_group_accs_id").
		Joins("JOIN m_mitra_administrator on m_mitra_administrator_temp.mitra_adm_id=m_mitra_administrator.mitra_adm_id").
		Joins("JOIN m_mitra_employee on m_mitra_administrator.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_address on m_mitra_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra_group_access on m_mitra_administrator.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Where("mitra_adm_temp_id = ? AND mitra_adm_temp_status = ? AND m_mitra_employee_temp.mitra_temp_id = ?", pid, status, mitra).
		Find(&mitraadministrator).Error
	if err != nil {
		return &MitraAdministratorTempData{}, err
	}
	return &mitraadministrator, nil
}

//ADMINISTRATOR LOGIN
//====================================================================================================================================

func HashMitraAdministrator(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (p *MitraAdministratorTemp) BeforeSaveMitraAdministratorTemp() error {
	hashedPassword, err := HashMitraAdministrator(p.MitraAdministratorTempPassword)
	if err != nil {
		return err
	}
	p.MitraAdministratorTempPassword = string(hashedPassword)
	return nil
}

func (p *MitraAdministratorTempPend) BeforeSaveMitraAdministratorTemp() error {
	hashedPassword, err := HashMitraAdministrator(p.MitraAdministratorTempPassword)
	if err != nil {
		return err
	}
	p.MitraAdministratorTempPassword = string(hashedPassword)
	return nil
}

//ADDITIONAL
//====================================================================================================================================

func (p *MitraAdministratorTemp) FindMitraAdministratorTempByEmployeeTempID(db *gorm.DB, pid uint64) (*MitraAdministratorTempPend, error) {
	var err error
	mitraadministrator := MitraAdministratorTempPend{}
	err = db.Debug().Model(&MitraAdministratorTempPend{}).Limit(100).
		Select(`m_mitra_administrator_temp.mitra_adm_temp_id,
				m_mitra_administrator_temp.mitra_emp_temp_id,
				m_mitra_administrator_temp.mitra_group_accs_temp_id,
				m_mitra_administrator_temp.mitra_adm_temp_email,
				m_mitra_administrator_temp.mitra_adm_temp_username,
				m_mitra_administrator_temp.mitra_adm_temp_password,
				m_mitra_administrator_temp.mitra_adm_temp_phone_code,
				m_mitra_administrator_temp.mitra_adm_temp_phone,
				m_mitra_administrator_temp.mitra_adm_temp_reason,
				m_mitra_administrator_temp.mitra_adm_temp_status,
				m_mitra_administrator_temp.mitra_adm_temp_action_before,
				m_mitra_administrator_temp.mitra_adm_temp_created_by,
				m_mitra_administrator_temp.mitra_adm_temp_created_at`).
		Where("mitra_emp_temp_id = ?", pid).
		Find(&mitraadministrator).Error
	if err != nil {
		return &MitraAdministratorTempPend{}, err
	}
	return &mitraadministrator, nil
}

func (p *MitraAdministratorTemp) FindMitraAdministratorTempStatusByEmployeeTempID(db *gorm.DB, pid uint64, status uint64) (*MitraAdministratorTempPend, error) {
	var err error
	mitraadministrator := MitraAdministratorTempPend{}
	err = db.Debug().Model(&MitraAdministratorTempPend{}).Limit(100).
		Select(`m_mitra_administrator_temp.mitra_adm_temp_id,
				m_mitra_administrator_temp.mitra_emp_temp_id,
				m_mitra_administrator_temp.mitra_group_accs_temp_id,
				m_mitra_administrator_temp.mitra_adm_temp_email,
				m_mitra_administrator_temp.mitra_adm_temp_username,
				m_mitra_administrator_temp.mitra_adm_temp_password,
				m_mitra_administrator_temp.mitra_adm_temp_phone_code,
				m_mitra_administrator_temp.mitra_adm_temp_phone,
				m_mitra_administrator_temp.mitra_adm_temp_reason,
				m_mitra_administrator_temp.mitra_adm_temp_status,
				m_mitra_administrator_temp.mitra_adm_temp_action_before,
				m_mitra_administrator_temp.mitra_adm_temp_created_by,
				m_mitra_administrator_temp.mitra_adm_temp_created_at`).
		Where("mitra_emp_temp_id = ? AND mitra_adm_temp_status = ?", pid, status).
		Find(&mitraadministrator).Error
	if err != nil {
		return &MitraAdministratorTempPend{}, err
	}
	return &mitraadministrator, nil
}

func (p *MitraAdministratorTemp) FindMitraAdministratorTempByMitraTempID(db *gorm.DB, pid uint64) ([]MitraAdministratorTempPendData, error) {
	var err error
	mitraemployee := []MitraAdministratorTempPendData{}
	err = db.Debug().Model(&MitraAdministratorTempPendData{}).Limit(100).
		Select(`m_mitra_administrator_temp.mitra_adm_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_id,
				m_mitra_employee_temp.mitra_temp_id,
				m_mitra_employee_temp.mitra_branch_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_code,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_gender,
				m_mitra_employee_temp.mitra_emp_temp_birth_place,
				m_mitra_employee_temp.mitra_emp_temp_birth_date,
				m_mitra_employee_temp.mitra_emp_temp_address,
				m_mitra_employee_temp.address_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_photo,
				m_mitra_employee_temp.mitra_emp_temp_reason,
				m_mitra_employee_temp.mitra_emp_temp_status,
				m_mitra_employee_temp.mitra_emp_temp_created_by,
				m_mitra_employee_temp.mitra_emp_temp_created_at,
				m_mitra_administrator_temp.mitra_group_accs_temp_id,
				m_mitra_administrator_temp.mitra_adm_temp_email,
				m_mitra_administrator_temp.mitra_adm_temp_username,
				m_mitra_administrator_temp.mitra_adm_temp_phone_code,
				m_mitra_administrator_temp.mitra_adm_temp_phone,
				m_mitra_administrator_temp.mitra_adm_temp_reason,
				m_mitra_administrator_temp.mitra_adm_temp_status,
				m_mitra_administrator_temp.mitra_adm_temp_action_before,
				m_mitra_administrator_temp.mitra_adm_temp_created_by,
				m_mitra_administrator_temp.mitra_adm_temp_created_at`).
		Joins("JOIN m_mitra_employee_temp on m_mitra_administrator_temp.mitra_emp_temp_id=m_mitra_employee_temp.mitra_emp_temp_id").
		Where("mitra_temp_id = ?", pid).
		Find(&mitraemployee).Error
	if err != nil {
		return []MitraAdministratorTempPendData{}, err
	}
	return mitraemployee, nil
}

func (p *MitraAdministratorTemp) FindMitraAdministratorTempStatusByMitraTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraAdministratorTempPendData, error) {
	var err error
	mitraemployee := []MitraAdministratorTempPendData{}
	err = db.Debug().Model(&MitraAdministratorTempPendData{}).Limit(100).
		Select(`m_mitra_administrator_temp.mitra_adm_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_id,
				m_mitra_employee_temp.mitra_temp_id,
				m_mitra_employee_temp.mitra_branch_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_code,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_gender,
				m_mitra_employee_temp.mitra_emp_temp_birth_place,
				m_mitra_employee_temp.mitra_emp_temp_birth_date,
				m_mitra_employee_temp.mitra_emp_temp_address,
				m_mitra_employee_temp.address_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_photo,
				m_mitra_employee_temp.mitra_emp_temp_reason,
				m_mitra_employee_temp.mitra_emp_temp_status,
				m_mitra_employee_temp.mitra_emp_temp_created_by,
				m_mitra_employee_temp.mitra_emp_temp_created_at,
				m_mitra_administrator_temp.mitra_group_accs_temp_id,
				m_mitra_administrator_temp.mitra_adm_temp_email,
				m_mitra_administrator_temp.mitra_adm_temp_username,
				m_mitra_administrator_temp.mitra_adm_temp_phone_code,
				m_mitra_administrator_temp.mitra_adm_temp_phone,
				m_mitra_administrator_temp.mitra_adm_temp_reason,
				m_mitra_administrator_temp.mitra_adm_temp_status,
				m_mitra_administrator_temp.mitra_adm_temp_action_before,
				m_mitra_administrator_temp.mitra_adm_temp_created_by,
				m_mitra_administrator_temp.mitra_adm_temp_created_at`).
		Joins("JOIN m_mitra_employee_temp on m_mitra_administrator_temp.mitra_emp_temp_id=m_mitra_employee_temp.mitra_emp_temp_id").
		Where("mitra_temp_id = ? AND m_mitra_administrator_temp.mitra_adm_temp_status = ?", pid, status).
		Find(&mitraemployee).Error
	if err != nil {
		return []MitraAdministratorTempPendData{}, err
	}
	return mitraemployee, nil
}
