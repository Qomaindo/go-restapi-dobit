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

type BibiteAdministratorTemp struct {
	BibiteAdministratorTempID           uint64    `gorm:"column:bbt_adm_temp_id;primary_key;" json:"bbt_adm_temp_id"`
	BibiteAdministratorID               uint64    `gorm:"column:bbt_adm_id" json:"bbt_adm_id"`
	BibiteEmployeeTempID                uint64    `gorm:"column:bbt_emp_temp_id" json:"bbt_emp_temp_id"`
	BibiteGroupTempID                   uint64    `gorm:"column:bbt_group_temp_id" json:"bbt_group_temp_id"`
	BibiteAdministratorTempEmail        string    `gorm:"column:bbt_adm_temp_email;size:255" json:"bbt_adm_temp_email"`
	BibiteAdministratorTempUsername     string    `gorm:"column:bbt_adm_temp_username;size:125" json:"bbt_adm_temp_username"`
	BibiteAdministratorTempPassword     string    `gorm:"column:bbt_adm_temp_password;size:125" json:"bbt_adm_temp_password"`
	BibiteAdministratorTempPhoneCode    string    `gorm:"column:bbt_adm_temp_phone_code;size:6" json:"bbt_adm_temp_phone_code"`
	BibiteAdministratorTempPhone        string    `gorm:"column:bbt_adm_temp_phone;size:20" json:"bbt_adm_temp_phone"`
	BibiteAdministratorTempPIN          string    `gorm:"column:bbt_adm_temp_pin;size:255" json:"bbt_adm_temp_pin"`
	BibiteAdministratorTempReason       string    `gorm:"column:bbt_adm_temp_reason" json:"bbt_adm_temp_reason"`
	BibiteAdministratorTempStatus       uint64    `gorm:"column:bbt_adm_temp_status;size:2" json:"bbt_adm_temp_status"`
	BibiteAdministratorTempActionBefore uint64    `gorm:"column:bbt_adm_temp_action_before;size:2" json:"bbt_adm_temp_action_before"`
	BibiteAdministratorTempCreatedBy    string    `gorm:"column:bbt_adm_temp_created_by;size:125" json:"bbt_adm_temp_created_by"`
	BibiteAdministratorTempCreatedAt    time.Time `gorm:"column:bbt_adm_temp_created_at;default:CURRENT_TIMESTAMP" json:"bbt_adm_temp_created_at"`
}

type BibiteAdministratorTempPend struct {
	BibiteAdministratorTempID           uint64    `gorm:"column:bbt_adm_temp_id;primary_key;" json:"bbt_adm_temp_id"`
	BibiteAdministratorID               *uint64   `gorm:"column:bbt_adm_id" json:"bbt_adm_id"`
	BibiteEmployeeTempID                uint64    `gorm:"column:bbt_emp_temp_id" json:"bbt_emp_temp_id"`
	BibiteGroupTempID                   uint64    `gorm:"column:bbt_group_temp_id" json:"bbt_group_temp_id"`
	BibiteAdministratorTempEmail        string    `gorm:"column:bbt_adm_temp_email;size:255" json:"bbt_adm_temp_email"`
	BibiteAdministratorTempUsername     string    `gorm:"column:bbt_adm_temp_username;size:125" json:"bbt_adm_temp_username"`
	BibiteAdministratorTempPassword     string    `gorm:"column:bbt_adm_temp_password;size:125" json:"bbt_adm_temp_password"`
	BibiteAdministratorTempPhoneCode    string    `gorm:"column:bbt_adm_temp_phone_code;size:6" json:"bbt_adm_temp_phone_code"`
	BibiteAdministratorTempPhone        string    `gorm:"column:bbt_adm_temp_phone;size:20" json:"bbt_adm_temp_phone"`
	BibiteAdministratorTempPIN          string    `gorm:"column:bbt_adm_temp_pin;size:255" json:"bbt_adm_temp_pin"`
	BibiteAdministratorTempReason       string    `gorm:"column:bbt_adm_temp_reason" json:"bbt_adm_temp_reason"`
	BibiteAdministratorTempStatus       uint64    `gorm:"column:bbt_adm_temp_status;size:2" json:"bbt_adm_temp_status"`
	BibiteAdministratorTempActionBefore uint64    `gorm:"column:bbt_adm_temp_action_before;size:2" json:"bbt_adm_temp_action_before"`
	BibiteAdministratorTempCreatedBy    string    `gorm:"column:bbt_adm_temp_created_by;size:125" json:"bbt_adm_temp_created_by"`
	BibiteAdministratorTempCreatedAt    time.Time `gorm:"column:bbt_adm_temp_created_at;default:CURRENT_TIMESTAMP" json:"bbt_adm_temp_created_at"`
}

type BibiteAdministratorTempPendData struct {
	BibiteAdministratorTempID           uint64    `gorm:"column:bbt_adm_temp_id;primary_key;" json:"bbt_adm_temp_id"`
	BibiteAdministratorID               *uint64   `gorm:"column:bbt_adm_id" json:"bbt_adm_id"`
	BibiteEmployeeTempID                uint64    `gorm:"column:bbt_emp_temp_id" json:"bbt_emp_temp_id"`
	BibiteEmployeeTempName              string    `gorm:"column:bbt_emp_temp_name" json:"bbt_emp_temp_name"`
	BibiteEmployeeTempGender            string    `gorm:"column:bbt_emp_temp_gender;size:255" json:"bbt_emp_temp_gender"`
	BibiteEmployeeTempBirthPlace        string    `gorm:"column:bbt_emp_temp_birth_place;size:255" json:"bbt_emp_temp_birth_place"`
	BibiteEmployeeTempBirthDate         string    `gorm:"column:bbt_emp_temp_birth_date;size:255" json:"bbt_emp_temp_birth_date"`
	BibiteEmployeeTempAddress           string    `gorm:"column:bbt_emp_temp_address;size:255" json:"bbt_emp_temp_address"`
	AddressTempID                       uint64    `gorm:"column:address_temp_id;size:255" json:"address_id"`
	CountryTempID                       uint64    `gorm:"column:country_temp_id" json:"country_temp_id"`
	CountryTempName                     string    `gorm:"column:country_temp_name" json:"country_temp_name"`
	ProvinceTempID                      uint64    `gorm:"column:province_temp_id" json:"province_temp_id"`
	ProvinceTempName                    string    `gorm:"column:province_temp_name" json:"province_temp_name"`
	RegencyTempID                       uint64    `gorm:"column:regency_temp_id" json:"regency_temp_id"`
	RegencyTempName                     string    `gorm:"column:regency_temp_name" json:"regency_temp_name"`
	DistrictTempID                      uint64    `gorm:"column:district_temp_id" json:"district_temp_id"`
	DistrictTempName                    string    `gorm:"column:district_temp_name" json:"district_temp_name"`
	VillageTempID                       uint64    `gorm:"column:village_temp_id" json:"village_temp_id"`
	VillageTempName                     string    `gorm:"column:village_temp_name" json:"village_temp_name"`
	BibiteEmployeeTempPhoto             string    `gorm:"column:bbt_emp_temp_photo;size:255" json:"bbt_emp_temp_photo"`
	BibiteGroupTempID                   uint64    `gorm:"column:bbt_group_temp_id" json:"bbt_group_temp_id"`
	BibiteGroupTemName                  string    `gorm:"column:bbt_group_temp_name" json:"bbt_group_temp_name"`
	BibiteAdministratorTempEmail        string    `gorm:"column:bbt_adm_temp_email;size:255" json:"bbt_adm_temp_email"`
	BibiteAdministratorTempUsername     string    `gorm:"column:bbt_adm_temp_username;size:125" json:"bbt_adm_temp_username"`
	BibiteAdministratorTempPassword     string    `gorm:"column:bbt_adm_temp_password;size:125" json:"bbt_adm_temp_password"`
	BibiteAdministratorTempPhoneCode    string    `gorm:"column:bbt_adm_temp_phone_code;size:6" json:"bbt_adm_temp_phone_code"`
	BibiteAdministratorTempPhone        string    `gorm:"column:bbt_adm_temp_phone;size:20" json:"bbt_adm_temp_phone"`
	BibiteAdministratorTempPIN          string    `gorm:"column:bbt_adm_temp_pin;size:255" json:"bbt_adm_temp_pin"`
	BibiteAdministratorTempReason       string    `gorm:"column:bbt_adm_temp_reason" json:"bbt_adm_temp_reason"`
	BibiteAdministratorTempStatus       uint64    `gorm:"column:bbt_adm_temp_status;size:2" json:"bbt_adm_temp_status"`
	BibiteAdministratorTempActionBefore uint64    `gorm:"column:bbt_adm_temp_action_before;size:2" json:"bbt_adm_temp_action_before"`
	BibiteAdministratorTempCreatedBy    string    `gorm:"column:bbt_adm_temp_created_by;size:125" json:"bbt_adm_temp_created_by"`
	BibiteAdministratorTempCreatedAt    time.Time `gorm:"column:bbt_adm_temp_created_at;default:CURRENT_TIMESTAMP" json:"bbt_adm_temp_created_at"`
}

type BibiteAdministratorTempData struct {
	BibiteAdministratorTempID                  uint64    `gorm:"column:bbt_adm_temp_id" json:"bbt_adm_temp_id"`
	BibiteEmployeeTempID                       uint64    `gorm:"column:bbt_emp_temp_id" json:"bbt_emp_temp_id"`
	BibiteEmployeeTempName                     string    `gorm:"column:bbt_emp_temp_name;size:255" json:"bbt_emp_temp_name"`
	BibiteEmployeeTempGender                   string    `gorm:"column:bbt_emp_temp_gender;size:255" json:"bbt_emp_temp_gender"`
	BibiteEmployeeTempBirthPlace               string    `gorm:"column:bbt_emp_temp_birth_place;size:255" json:"bbt_emp_temp_birth_place"`
	BibiteEmployeeTempBirthDate                string    `gorm:"column:bbt_emp_temp_birth_date;size:255" json:"bbt_emp_temp_birth_date"`
	BibiteEmployeeTempAddress                  string    `gorm:"column:bbt_emp_temp_address;size:255" json:"bbt_emp_temp_address"`
	AddressTempID                              uint64    `gorm:"column:address_temp_id;size:255" json:"address_id"`
	CountryTempID                              uint64    `gorm:"column:country_temp_id" json:"country_temp_id"`
	CountryTempName                            string    `gorm:"column:country_temp_name" json:"country_temp_name"`
	ProvinceTempID                             uint64    `gorm:"column:province_temp_id" json:"province_temp_id"`
	ProvinceTempName                           string    `gorm:"column:province_temp_name" json:"province_temp_name"`
	RegencyTempID                              uint64    `gorm:"column:regency_temp_id" json:"regency_temp_id"`
	RegencyTempName                            string    `gorm:"column:regency_temp_name" json:"regency_temp_name"`
	DistrictTempID                             uint64    `gorm:"column:district_temp_id" json:"district_temp_id"`
	DistrictTempName                           string    `gorm:"column:district_temp_name" json:"district_temp_name"`
	VillageTempID                              uint64    `gorm:"column:village_temp_id" json:"village_temp_id"`
	VillageTempName                            string    `gorm:"column:village_temp_name" json:"village_temp_name"`
	BibiteEmployeeTempPhoto                    string    `gorm:"column:bbt_emp_temp_photo;size:255" json:"bbt_emp_temp_photo"`
	BibiteGroupTempID                          uint64    `gorm:"column:bbt_group_temp_id" json:"bbt_group_temp_id"`
	BibiteGroupTempCode                        string    `gorm:"column:bbt_group_temp_code;size:25" json:"bbt_group_temp_code"`
	BibiteGroupTempName                        string    `gorm:"column:bbt_group_temp_name;size:255" json:"bbt_group_temp_name"`
	BibiteGroupTempRole                        string    `gorm:"column:bbt_group_temp_role;size:125" json:"bbt_group_temp_role"`
	BibiteGroupTempType                        uint64    `gorm:"column:bbt_group_temp_type;size:2" json:"bbt_group_temp_type"`
	BibiteAdministratorTempEmail               string    `gorm:"column:bbt_adm_temp_email;size:255" json:"bbt_adm_temp_email"`
	BibiteAdministratorTempUsername            string    `gorm:"column:bbt_adm_temp_username;size:125" json:"bbt_adm_temp_username"`
	BibiteAdministratorTempPhoneCode           string    `gorm:"column:bbt_adm_temp_phone_code;size:6" json:"bbt_adm_temp_phone_code"`
	BibiteAdministratorTempPhone               string    `gorm:"column:bbt_adm_temp_phone;size:20" json:"bbt_adm_temp_phone"`
	BibiteAdministratorTempReason              string    `gorm:"column:bbt_adm_temp_reason" json:"bbt_adm_temp_reason"`
	BibiteAdministratorTempStatus              int64     `gorm:"column:bbt_adm_temp_status;size:2" json:"bbt_adm_temp_status"`
	BibiteAdministratorTempActionBefore        uint64    `gorm:"column:bbt_adm_temp_action_before;size:2" json:"bbt_adm_temp_action_before"`
	BibiteAdministratorTempCreatedBy           string    `gorm:"column:bbt_adm_temp_created_by;size:125" json:"bbt_adm_temp_created_by"`
	BibiteAdministratorTempCreatedAt           time.Time `gorm:"column:bbt_adm_temp_created_at;default:CURRENT_TIMESTAMP" json:"bbt_adm_temp_created_at"`
	BibiteAdministratorID                      uint64    `gorm:"column:bbt_adm_id" json:"bbt_adm_id"`
	BibiteEmployeeID                           uint64    `gorm:"column:bbt_emp_id" json:"bbt_emp_id"`
	BibiteEmployeeCode                         string    `gorm:"column:bbt_emp_code;size:25" json:"bbt_emp_code"`
	BibiteEmployeeName                         string    `gorm:"column:bbt_emp_name;size:255" json:"bbt_emp_name"`
	BibiteEmployeeGender                       string    `gorm:"column:bbt_emp_gender;size:255" json:"bbt_emp_gender"`
	BibiteEmployeeBirthPlace                   string    `gorm:"column:bbt_emp_birth_place;size:255" json:"bbt_emp_birth_place"`
	BibiteEmployeeBirthDate                    string    `gorm:"column:bbt_emp_birth_date;size:255" json:"bbt_emp_birth_date"`
	BibiteEmployeeAddress                      string    `gorm:"column:bbt_emp_address;size:255" json:"bbt_emp_address"`
	AddressID                                  uint64    `gorm:"column:address_id;size:255" json:"address_id"`
	CountryID                                  uint64    `gorm:"column:country_id" json:"country_id"`
	CountryName                                string    `gorm:"column:country_name" json:"country_name"`
	ProvinceID                                 uint64    `gorm:"column:province_id" json:"province_id"`
	ProvinceName                               string    `gorm:"column:province_name" json:"province_name"`
	RegencyID                                  uint64    `gorm:"column:regency_id" json:"regency_id"`
	RegencyName                                string    `gorm:"column:regency_name" json:"regency_name"`
	DistrictID                                 uint64    `gorm:"column:district_id" json:"district_id"`
	DistrictName                               string    `gorm:"column:district_name" json:"district_name"`
	VillageID                                  uint64    `gorm:"column:village_id" json:"village_id"`
	VillageName                                string    `gorm:"column:village_name" json:"village_name"`
	BibiteEmployeePhoto                        string    `gorm:"column:bbt_emp_photo;size:255" json:"bbt_emp_photo"`
	BibiteGroupID                              uint64    `gorm:"column:bbt_group_id" json:"bbt_group_id"`
	BibiteGroupCode                            string    `gorm:"column:bbt_group_code;size:25" json:"bbt_group_code"`
	BibiteGroupName                            string    `gorm:"column:bbt_group_name;size:255" json:"bbt_group_name"`
	BibiteGroupRole                            string    `gorm:"column:bbt_group_role;size:125" json:"bbt_group_role"`
	BibiteGroupType                            uint64    `gorm:"column:bbt_group_type;size:2" json:"bbt_group_type"`
	BibiteAdministratorEmail                   string    `gorm:"column:bbt_adm_email" json:"bbt_adm_email"`
	BibiteAdministratorUsername                string    `gorm:"column:bbt_username" json:"bbt_username"`
	BibiteAdministratorPhoneCode               string    `gorm:"column:bbt_adm_phone_code;size:6" json:"bbt_adm_phone_code"`
	BibiteAdministratorPhone                   string    `gorm:"column:bbt_adm_phone;size:20" json:"bbt_adm_phone"`
	BibiteAdministratorStatus                  int64     `gorm:"column:bbt_adm_status;size:1" json:"bbt_adm_status"`
	BibiteAdministratorIsNew                   bool      `gorm:"column:bbt_adm_is_new" json:"bbt_adm_is_new"`
	BibiteAdministratorIsRequestChangePassword bool      `gorm:"column:bbt_adm_is_req_chpw" json:"bbt_adm_is_req_chpw"`
	BibiteAdministratorIsLocked                bool      `gorm:"column:bbt_adm_is_locked" json:"bbt_adm_is_locked"`
	BibiteAdministratorFailAttempt             uint64    `gorm:"column:bbt_adm_fail_attempt;size:1" json:"bbt_adm_fail_attempt"`
	BibiteAdministratorLockedReason            string    `gorm:"column:bbt_adm_locked_reason;size:255" json:"bbt_adm_locked_reason"`
	BibiteAdministratorCreatedBy               string    `gorm:"column:bbt_adm_created_by" json:"bbt_adm_created_by"`
	BibiteAdministratorCreatedAt               time.Time `gorm:"column:bbt_adm_created_at" json:"bbt_adm_created_at"`
	BibiteAdministratorUpdatedBy               string    `gorm:"column:bbt_adm_updated_by" json:"bbt_adm_updated_by"`
	BibiteAdministratorUpdatedAt               time.Time `gorm:"column:bbt_adm_updated_at" json:"bbt_adm_updated_at"`
	BibiteAdministratorDeletedBy               string    `gorm:"column:bbt_adm_deleted_by" json:"bbt_adm_deleted_by"`
	BibiteAdministratorDeletedAt               time.Time `gorm:"column:bbt_adm_deleted_at" json:"bbt_adm_deleted_at"`
}

type ResponseBibiteAdministratorTemps struct {
	Status  int                            `json:"status"`
	Message string                         `json:"message"`
	Data    *[]BibiteAdministratorTempData `json:"data"`
}

type ResponseBibiteAdministratorTemp struct {
	Status  int                          `json:"status"`
	Message string                       `json:"message"`
	Data    *BibiteAdministratorTempData `json:"data"`
}

type ResponseBibiteAdministratorTempsPend struct {
	Status  int                                `json:"status"`
	Message string                             `json:"message"`
	Data    *[]BibiteAdministratorTempPendData `json:"data"`
}

type ResponseBibiteAdministratorTempPend struct {
	Status  int                              `json:"status"`
	Message string                           `json:"message"`
	Data    *BibiteAdministratorTempPendData `json:"data"`
}

type ResponseBibiteAdministratorTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (BibiteAdministratorTemp) TableName() string {
	return "m_bibite_administrator_temp"
}

func (BibiteAdministratorTempPend) TableName() string {
	return "m_bibite_administrator_temp"
}

func (BibiteAdministratorTempPendData) TableName() string {
	return "m_bibite_administrator_temp"
}

func (BibiteAdministratorTempData) TableName() string {
	return "m_bibite_administrator_temp"
}

func (p *BibiteAdministratorTemp) Prepare() {
	p.BibiteAdministratorID = p.BibiteAdministratorID
	p.BibiteEmployeeTempID = p.BibiteEmployeeTempID
	p.BibiteGroupTempID = p.BibiteGroupTempID
	p.BibiteAdministratorTempEmail = html.EscapeString(strings.TrimSpace(p.BibiteAdministratorTempEmail))
	p.BibiteAdministratorTempUsername = html.EscapeString(strings.TrimSpace(p.BibiteAdministratorTempUsername))
	p.BibiteAdministratorTempPassword = p.BibiteAdministratorTempPassword
	p.BibiteAdministratorTempPhoneCode = html.EscapeString(strings.TrimSpace(p.BibiteAdministratorTempPhoneCode))
	p.BibiteAdministratorTempPhone = html.EscapeString(strings.TrimSpace(p.BibiteAdministratorTempPhone))
	p.BibiteAdministratorTempPIN = html.EscapeString(strings.TrimSpace(p.BibiteAdministratorTempPIN))
	p.BibiteAdministratorTempPIN = html.EscapeString(strings.TrimSpace(p.BibiteAdministratorTempPIN))
	p.BibiteAdministratorTempReason = p.BibiteAdministratorTempReason
	p.BibiteAdministratorTempCreatedBy = p.BibiteAdministratorTempCreatedBy
	p.BibiteAdministratorTempCreatedAt = time.Now()
}

func (p *BibiteAdministratorTempPend) Prepare() {
	p.BibiteAdministratorID = nil
	p.BibiteEmployeeTempID = p.BibiteEmployeeTempID
	p.BibiteGroupTempID = p.BibiteGroupTempID
	p.BibiteAdministratorTempEmail = html.EscapeString(strings.TrimSpace(p.BibiteAdministratorTempEmail))
	p.BibiteAdministratorTempUsername = html.EscapeString(strings.TrimSpace(p.BibiteAdministratorTempUsername))
	p.BibiteAdministratorTempPassword = p.BibiteAdministratorTempPassword
	p.BibiteAdministratorTempPhoneCode = html.EscapeString(strings.TrimSpace(p.BibiteAdministratorTempPhoneCode))
	p.BibiteAdministratorTempPhone = html.EscapeString(strings.TrimSpace(p.BibiteAdministratorTempPhone))
	p.BibiteAdministratorTempPIN = html.EscapeString(strings.TrimSpace(p.BibiteAdministratorTempPIN))
	p.BibiteAdministratorTempPIN = html.EscapeString(strings.TrimSpace(p.BibiteAdministratorTempPIN))
	p.BibiteAdministratorTempReason = p.BibiteAdministratorTempReason
	p.BibiteAdministratorTempCreatedBy = p.BibiteAdministratorTempCreatedBy
	p.BibiteAdministratorTempCreatedAt = time.Now()
}

func (p *BibiteAdministratorTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.BibiteAdministratorTempPassword == "" {
			return errors.New("Required BibiteAdministratorTemp Password")
		}
		if p.BibiteAdministratorTempEmail == "" {
			return errors.New("Required BibiteAdministratorTemp Email")
		}
		return nil
	}
}

func (p *BibiteAdministratorTemp) SaveBibiteAdministratorTemp(db *gorm.DB) (*BibiteAdministratorTemp, error) {
	var err error
	err = p.BeforeSaveBibiteAdministratorTemp()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Debug().Model(&BibiteAdministratorTemp{}).Create(&p).Error
	if err != nil {
		return &BibiteAdministratorTemp{}, err
	}
	return p, nil
}

func (p *BibiteAdministratorTempPend) SaveBibiteAdministratorTempPend(db *gorm.DB) (*BibiteAdministratorTempPend, error) {
	var err error
	err = p.BeforeSaveBibiteAdministratorTemp()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Debug().Model(&BibiteAdministratorTempPend{}).Create(&p).Error
	if err != nil {
		return &BibiteAdministratorTempPend{}, err
	}
	return p, nil
}

func (p *BibiteAdministratorTemp) FindAllBibiteAdministratorTemp(db *gorm.DB) (*[]BibiteAdministratorTempPendData, error) {
	var err error
	bibiteadministrator := []BibiteAdministratorTempPendData{}
	err = db.Debug().Model(&BibiteAdministratorTempPendData{}).Limit(100).
		Select(`m_bibite_administrator_temp.bbt_adm_temp_id,
				m_bibite_administrator_temp.bbt_emp_temp_id,
				m_bibite_employee_temp.bbt_emp_temp_name,
				m_bibite_employee_temp.bbt_emp_temp_gender,
				m_bibite_employee_temp.bbt_emp_temp_birth_place,
				m_bibite_employee_temp.bbt_emp_temp_birth_date,
				m_bibite_employee_temp.bbt_emp_temp_address,
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
				m_bibite_employee_temp.bbt_emp_temp_photo,
				m_bibite_administrator_temp.bbt_group_temp_id,
				m_bibite_group_temp.bbt_group_name as bbt_group_temp_name,
				m_bibite_administrator_temp.bbt_adm_temp_email,
				m_bibite_administrator_temp.bbt_adm_temp_username,
				m_bibite_administrator_temp.bbt_adm_temp_phone_code,
				m_bibite_administrator_temp.bbt_adm_temp_phone,
				m_bibite_administrator_temp.bbt_adm_temp_reason,
				m_bibite_administrator_temp.bbt_adm_temp_status,
				m_bibite_administrator_temp.bbt_adm_temp_action_before,
				m_bibite_administrator_temp.bbt_adm_temp_created_by,
				m_bibite_administrator_temp.bbt_adm_temp_created_at at time zone 'Asia/Jakarta' as bbt_adm_temp_created_at`).
		Joins("JOIN m_bibite_employee_temp on m_bibite_administrator_temp.bbt_emp_temp_id=m_bibite_employee_temp.bbt_emp_temp_id").
		Joins("JOIN m_address_temp on m_bibite_employee_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_bibite_group m_bibite_group_temp on m_bibite_administrator_temp.bbt_group_temp_id=m_bibite_group_temp.bbt_group_id").
		Order("bbt_adm_temp_created_at desc").
		Find(&bibiteadministrator).Error
	if err != nil {
		return &[]BibiteAdministratorTempPendData{}, err
	}
	return &bibiteadministrator, nil
}

func (p *BibiteAdministratorTemp) FindAllBibiteAdministratorTempPendingActive(db *gorm.DB) (*[]BibiteAdministratorTempPendData, error) {
	var err error
	bibiteadministrator := []BibiteAdministratorTempPendData{}
	err = db.Debug().Model(&BibiteAdministratorTempPendData{}).Limit(100).
		Select(`m_bibite_administrator_temp.bbt_adm_temp_id,
				m_bibite_administrator_temp.bbt_emp_temp_id,
				m_bibite_employee_temp.bbt_emp_temp_name,
				m_bibite_employee_temp.bbt_emp_temp_gender,
				m_bibite_employee_temp.bbt_emp_temp_birth_place,
				m_bibite_employee_temp.bbt_emp_temp_birth_date,
				m_bibite_employee_temp.bbt_emp_temp_address,
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
				m_bibite_employee_temp.bbt_emp_temp_photo,
				m_bibite_administrator_temp.bbt_group_temp_id,
				m_bibite_group_temp.bbt_group_name as bbt_group_temp_name,
				m_bibite_administrator_temp.bbt_adm_temp_email,
				m_bibite_administrator_temp.bbt_adm_temp_username,
				m_bibite_administrator_temp.bbt_adm_temp_phone_code,
				m_bibite_administrator_temp.bbt_adm_temp_phone,
				m_bibite_administrator_temp.bbt_adm_temp_reason,
				m_bibite_administrator_temp.bbt_adm_temp_status,
				m_bibite_administrator_temp.bbt_adm_temp_action_before,
				m_bibite_administrator_temp.bbt_adm_temp_created_by,
				m_bibite_administrator_temp.bbt_adm_temp_created_at at time zone 'Asia/Jakarta' as bbt_adm_temp_created_at`).
		Joins("JOIN m_bibite_employee_temp on m_bibite_administrator_temp.bbt_emp_temp_id=m_bibite_employee_temp.bbt_emp_temp_id").
		Joins("JOIN m_address_temp on m_bibite_employee_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_bibite_group m_bibite_group_temp on m_bibite_administrator_temp.bbt_group_temp_id=m_bibite_group_temp.bbt_group_id").
		Where("bbt_adm_temp_status = ?", 11).
		Order("bbt_adm_temp_created_at desc").
		Find(&bibiteadministrator).Error
	if err != nil {
		return &[]BibiteAdministratorTempPendData{}, err
	}
	return &bibiteadministrator, nil
}

func (p *BibiteAdministratorTemp) FindAllBibiteAdministratorTempStatus(db *gorm.DB, status uint64) (*[]BibiteAdministratorTempData, error) {
	var err error
	bibiteadministrator := []BibiteAdministratorTempData{}
	err = db.Debug().Model(&BibiteAdministratorTempData{}).Limit(100).
		Select(`m_bibite_administrator_temp.bbt_adm_temp_id,
				m_bibite_administrator_temp.bbt_emp_temp_id,
				m_bibite_employee_temp.bbt_emp_temp_code,
				m_bibite_employee_temp.bbt_emp_temp_name,
				m_bibite_employee_temp.bbt_emp_temp_gender,
				m_bibite_employee_temp.bbt_emp_temp_birth_place,
				m_bibite_employee_temp.bbt_emp_temp_birth_date,
				m_bibite_employee_temp.bbt_emp_temp_address,
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
				m_bibite_employee_temp.bbt_emp_temp_photo,
				m_bibite_administrator_temp.bbt_group_temp_id,
				m_bibite_group_temp.bbt_group_code as bbt_group_temp_code,
				m_bibite_group_temp.bbt_group_name as bbt_group_temp_name,
				m_bibite_group_temp.bbt_group_role as bbt_group_temp_role,
				m_bibite_group_temp.bbt_group_type as bbt_group_temp_type,
				m_bibite_administrator_temp.bbt_adm_temp_email,
				m_bibite_administrator_temp.bbt_adm_temp_username,
				m_bibite_administrator_temp.bbt_adm_temp_phone_code,
				m_bibite_administrator_temp.bbt_adm_temp_phone,
				m_bibite_administrator_temp.bbt_adm_temp_reason,
				m_bibite_administrator_temp.bbt_adm_temp_status,
				m_bibite_administrator_temp.bbt_adm_temp_action_before,
				m_bibite_administrator_temp.bbt_adm_temp_created_by,
				m_bibite_administrator_temp.bbt_adm_temp_created_at at time zone 'Asia/Jakarta' as bbt_adm_temp_created_at,
				m_bibite_administrator.bbt_adm_id,
				m_bibite_administrator.bbt_emp_id,
				m_bibite_employee.bbt_emp_code,
				m_bibite_employee.bbt_emp_name,
				m_bibite_employee.bbt_emp_gender,
				m_bibite_employee.bbt_emp_birth_place,
				m_bibite_employee.bbt_emp_birth_date,
				m_bibite_employee.bbt_emp_address,
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
				m_bibite_employee.bbt_emp_photo,
				m_bibite_group.bbt_group_id,
				m_bibite_group.bbt_group_code,
				m_bibite_group.bbt_group_name,
				m_bibite_group.bbt_group_role,
				m_bibite_group.bbt_group_type,
				m_bibite_administrator.bbt_adm_email,
				m_bibite_administrator.bbt_adm_username,
				m_bibite_administrator.bbt_adm_phone_code,
				m_bibite_administrator.bbt_adm_phone,
				m_bibite_administrator.bbt_adm_status,
				m_bibite_administrator.bbt_adm_is_new,
				m_bibite_administrator.bbt_adm_is_req_chpw,
				m_bibite_administrator.bbt_adm_is_locked,
				m_bibite_administrator.bbt_adm_fail_attempt,
				m_bibite_administrator.bbt_adm_locked_reason,
				m_bibite_administrator.bbt_adm_created_by,
				m_bibite_administrator.bbt_adm_created_at at time zone 'Asia/Jakarta' as bbt_adm_created_at`).
		Joins("JOIN m_bibite_employee_temp on m_bibite_administrator_temp.bbt_emp_temp_id=m_bibite_employee_temp.bbt_emp_temp_id").
		Joins("JOIN m_address_temp on m_bibite_employee_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_bibite_group m_bibite_group_temp on m_bibite_administrator_temp.bbt_group_temp_id=m_bibite_group_temp.bbt_group_id").
		Joins("JOIN m_bibite_administrator on m_bibite_administrator_temp.bbt_adm_id=m_bibite_administrator.bbt_adm_id").
		Joins("JOIN m_bibite_employee on m_bibite_administrator.bbt_emp_id=m_bibite_employee.bbt_emp_id").
		Joins("JOIN m_address on m_bibite_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Joins("JOIN m_bibite_group on m_bibite_administrator.bbt_group_id=m_bibite_group.bbt_group_id").
		Where("bbt_adm_temp_status = ?", status).
		Order("bbt_adm_temp_created_at desc").
		Find(&bibiteadministrator).Error
	if err != nil {
		return &[]BibiteAdministratorTempData{}, err
	}
	return &bibiteadministrator, nil
}

func (p *BibiteAdministratorTemp) FindBibiteAdministratorTempDataByID(db *gorm.DB, pid uint64) (*BibiteAdministratorTemp, error) {
	var err error
	err = db.Debug().Model(&BibiteAdministratorTemp{}).Where("bbt_adm_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &BibiteAdministratorTemp{}, err
	}
	return p, nil
}

func (p *BibiteAdministratorTemp) FindBibiteAdministratorTempByIDPendingActive(db *gorm.DB, pid uint64) (*BibiteAdministratorTempPendData, error) {
	var err error
	bibiteadministrator := BibiteAdministratorTempPendData{}
	err = db.Debug().Model(&BibiteAdministratorTempPendData{}).Limit(100).
		Select(`m_bibite_administrator_temp.bbt_adm_temp_id,
				m_bibite_administrator_temp.bbt_emp_temp_id,
				m_bibite_employee_temp.bbt_emp_temp_name,
				m_bibite_employee_temp.bbt_emp_temp_gender,
				m_bibite_employee_temp.bbt_emp_temp_birth_place,
				m_bibite_employee_temp.bbt_emp_temp_birth_date,
				m_bibite_employee_temp.bbt_emp_temp_address,
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
				m_bibite_employee_temp.bbt_emp_temp_photo,
				m_bibite_administrator_temp.bbt_group_temp_id,
				m_bibite_group_temp.bbt_group_name as bbt_group_temp_name,
				m_bibite_administrator_temp.bbt_adm_temp_email,
				m_bibite_administrator_temp.bbt_adm_temp_username,
				m_bibite_administrator_temp.bbt_adm_temp_phone_code,
				m_bibite_administrator_temp.bbt_adm_temp_phone,
				m_bibite_administrator_temp.bbt_adm_temp_reason,
				m_bibite_administrator_temp.bbt_adm_temp_status,
				m_bibite_administrator_temp.bbt_adm_temp_action_before,
				m_bibite_administrator_temp.bbt_adm_temp_created_by,
				m_bibite_administrator_temp.bbt_adm_temp_created_at at time zone 'Asia/Jakarta' as bbt_adm_temp_created_at`).
		Joins("JOIN m_bibite_employee_temp on m_bibite_administrator_temp.bbt_emp_temp_id=m_bibite_employee_temp.bbt_emp_temp_id").
		Joins("JOIN m_address_temp on m_bibite_employee_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_bibite_group m_bibite_group_temp on m_bibite_administrator_temp.bbt_group_temp_id=m_bibite_group_temp.bbt_group_id").
		Where("bbt_adm_temp_id = ? AND bbt_adm_temp_status = ?", pid, 11).
		Order("bbt_adm_temp_created_at desc").
		Find(&bibiteadministrator).Error
	if err != nil {
		return &BibiteAdministratorTempPendData{}, err
	}
	return &bibiteadministrator, nil
}

func (p *BibiteAdministratorTemp) FindBibiteAdministratorTempByID(db *gorm.DB, pid uint64) (*BibiteAdministratorTempData, error) {
	var err error
	bibiteadministrator := BibiteAdministratorTempData{}
	err = db.Debug().Model(&BibiteAdministratorTempData{}).Limit(100).
		Select(`m_bibite_administrator_temp.bbt_adm_temp_id,
				m_bibite_administrator_temp.bbt_emp_temp_id,
				m_bibite_employee_temp.bbt_emp_temp_code,
				m_bibite_employee_temp.bbt_emp_temp_name,
				m_bibite_employee_temp.bbt_emp_temp_gender,
				m_bibite_employee_temp.bbt_emp_temp_birth_place,
				m_bibite_employee_temp.bbt_emp_temp_birth_date,
				m_bibite_employee_temp.bbt_emp_temp_address,
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
				m_bibite_employee_temp.bbt_emp_temp_photo,
				m_bibite_administrator_temp.bbt_group_temp_id,
				m_bibite_group_temp.bbt_group_code as bbt_group_temp_code,
				m_bibite_group_temp.bbt_group_name as bbt_group_temp_name,
				m_bibite_group_temp.bbt_group_role as bbt_group_temp_role,
				m_bibite_group_temp.bbt_group_type as bbt_group_temp_type,
				m_bibite_administrator_temp.bbt_adm_temp_email,
				m_bibite_administrator_temp.bbt_adm_temp_username,
				m_bibite_administrator_temp.bbt_adm_temp_phone_code,
				m_bibite_administrator_temp.bbt_adm_temp_phone,
				m_bibite_administrator_temp.bbt_adm_temp_reason,
				m_bibite_administrator_temp.bbt_adm_temp_status,
				m_bibite_administrator_temp.bbt_adm_temp_action_before,
				m_bibite_administrator_temp.bbt_adm_temp_created_by,
				m_bibite_administrator_temp.bbt_adm_temp_created_at at time zone 'Asia/Jakarta' as bbt_adm_temp_created_at,
				m_bibite_administrator.bbt_adm_id,
				m_bibite_administrator.bbt_emp_id,
				m_bibite_employee.bbt_emp_code,
				m_bibite_employee.bbt_emp_name,
				m_bibite_employee.bbt_emp_gender,
				m_bibite_employee.bbt_emp_birth_place,
				m_bibite_employee.bbt_emp_birth_date,
				m_bibite_employee.bbt_emp_address,
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
				m_bibite_employee.bbt_emp_photo,
				m_bibite_group.bbt_group_id,
				m_bibite_group.bbt_group_code,
				m_bibite_group.bbt_group_name,
				m_bibite_group.bbt_group_role,
				m_bibite_group.bbt_group_type,
				m_bibite_administrator.bbt_adm_email,
				m_bibite_administrator.bbt_adm_username,
				m_bibite_administrator.bbt_adm_phone_code,
				m_bibite_administrator.bbt_adm_phone,
				m_bibite_administrator.bbt_adm_status,
				m_bibite_administrator.bbt_adm_is_new,
				m_bibite_administrator.bbt_adm_is_req_chpw,
				m_bibite_administrator.bbt_adm_is_locked,
				m_bibite_administrator.bbt_adm_fail_attempt,
				m_bibite_administrator.bbt_adm_locked_reason,
				m_bibite_administrator.bbt_adm_created_by,
				m_bibite_administrator.bbt_adm_created_at at time zone 'Asia/Jakarta' as bbt_adm_created_at`).
		Joins("JOIN m_bibite_employee_temp on m_bibite_administrator_temp.bbt_emp_temp_id=m_bibite_employee_temp.bbt_emp_temp_id").
		Joins("JOIN m_address_temp on m_bibite_employee_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_bibite_group m_bibite_group_temp on m_bibite_administrator_temp.bbt_group_temp_id=m_bibite_group_temp.bbt_group_id").
		Joins("JOIN m_bibite_administrator on m_bibite_administrator_temp.bbt_adm_id=m_bibite_administrator.bbt_adm_id").
		Joins("JOIN m_bibite_employee on m_bibite_administrator.bbt_emp_id=m_bibite_employee.bbt_emp_id").
		Joins("JOIN m_address on m_bibite_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Joins("JOIN m_bibite_group on m_bibite_administrator.bbt_group_id=m_bibite_group.bbt_group_id").
		Where("bbt_adm_temp_id = ?", pid).
		Order("bbt_adm_temp_created_at desc").
		Find(&bibiteadministrator).Error
	if err != nil {
		return &BibiteAdministratorTempData{}, err
	}
	return &bibiteadministrator, nil
}

func (p *BibiteAdministratorTemp) FindBibiteAdministratorTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*BibiteAdministratorTempData, error) {
	var err error
	bibiteadministrator := BibiteAdministratorTempData{}
	err = db.Debug().Model(&BibiteAdministratorTempData{}).Limit(100).
		Select(`m_bibite_administrator_temp.bbt_adm_temp_id,
				m_bibite_administrator_temp.bbt_emp_temp_id,
				m_bibite_employee_temp.bbt_emp_temp_code,
				m_bibite_employee_temp.bbt_emp_temp_name,
				m_bibite_employee_temp.bbt_emp_temp_gender,
				m_bibite_employee_temp.bbt_emp_temp_birth_place,
				m_bibite_employee_temp.bbt_emp_temp_birth_date,
				m_bibite_employee_temp.bbt_emp_temp_address,
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
				m_bibite_employee_temp.bbt_emp_temp_photo,
				m_bibite_administrator_temp.bbt_group_temp_id,
				m_bibite_group_temp.bbt_group_code as bbt_group_temp_code,
				m_bibite_group_temp.bbt_group_name as bbt_group_temp_name,
				m_bibite_group_temp.bbt_group_role as bbt_group_temp_role,
				m_bibite_group_temp.bbt_group_type as bbt_group_temp_type,
				m_bibite_administrator_temp.bbt_adm_temp_email,
				m_bibite_administrator_temp.bbt_adm_temp_username,
				m_bibite_administrator_temp.bbt_adm_temp_phone_code,
				m_bibite_administrator_temp.bbt_adm_temp_phone,
				m_bibite_administrator_temp.bbt_adm_temp_reason,
				m_bibite_administrator_temp.bbt_adm_temp_status,
				m_bibite_administrator_temp.bbt_adm_temp_action_before,
				m_bibite_administrator_temp.bbt_adm_temp_created_by,
				m_bibite_administrator_temp.bbt_adm_temp_created_at at time zone 'Asia/Jakarta' as bbt_adm_temp_created_at,
				m_bibite_administrator.bbt_adm_id,
				m_bibite_administrator.bbt_emp_id,
				m_bibite_employee.bbt_emp_code,
				m_bibite_employee.bbt_emp_name,
				m_bibite_employee.bbt_emp_gender,
				m_bibite_employee.bbt_emp_birth_place,
				m_bibite_employee.bbt_emp_birth_date,
				m_bibite_employee.bbt_emp_address,
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
				m_bibite_employee.bbt_emp_photo,
				m_bibite_group.bbt_group_id,
				m_bibite_group.bbt_group_code,
				m_bibite_group.bbt_group_name,
				m_bibite_group.bbt_group_role,
				m_bibite_group.bbt_group_type,
				m_bibite_administrator.bbt_adm_email,
				m_bibite_administrator.bbt_adm_username,
				m_bibite_administrator.bbt_adm_phone_code,
				m_bibite_administrator.bbt_adm_phone,
				m_bibite_administrator.bbt_adm_status,
				m_bibite_administrator.bbt_adm_is_new,
				m_bibite_administrator.bbt_adm_is_req_chpw,
				m_bibite_administrator.bbt_adm_is_locked,
				m_bibite_administrator.bbt_adm_fail_attempt,
				m_bibite_administrator.bbt_adm_locked_reason,
				m_bibite_administrator.bbt_adm_created_by,
				m_bibite_administrator.bbt_adm_created_at at time zone 'Asia/Jakarta' as bbt_adm_created_at`).
		Joins("JOIN m_bibite_employee_temp on m_bibite_administrator_temp.bbt_emp_temp_id=m_bibite_employee_temp.bbt_emp_temp_id").
		Joins("JOIN m_address_temp on m_bibite_employee_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_bibite_group m_bibite_group_temp on m_bibite_administrator_temp.bbt_group_temp_id=m_bibite_group_temp.bbt_group_id").
		Joins("JOIN m_bibite_administrator on m_bibite_administrator_temp.bbt_adm_id=m_bibite_administrator.bbt_adm_id").
		Joins("JOIN m_bibite_employee on m_bibite_administrator.bbt_emp_id=m_bibite_employee.bbt_emp_id").
		Joins("JOIN m_address on m_bibite_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Joins("JOIN m_bibite_group on m_bibite_administrator.bbt_group_id=m_bibite_group.bbt_group_id").
		Where("bbt_adm_temp_id = ? AND bbt_adm_temp_status = ?", pid, status).
		Order("bbt_adm_temp_created_at desc").
		Find(&bibiteadministrator).Error
	if err != nil {
		return &BibiteAdministratorTempData{}, err
	}
	return &bibiteadministrator, nil
}

func (p *BibiteAdministratorTemp) UpdateBibiteAdministratorTemp(db *gorm.DB, pid uint64) (*BibiteAdministratorTemp, error) {
	var err error
	db = db.Debug().Model(&BibiteAdministratorTemp{}).Where("bbt_adm_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_emp_temp_id":       p.BibiteEmployeeTempID,
			"bbt_group_temp_id":     p.BibiteGroupTempID,
			"bbt_adm_temp_email":    p.BibiteAdministratorTempEmail,
			"bbt_adm_temp_username": p.BibiteAdministratorTempUsername,
			"bbt_adm_temp_status":   p.BibiteAdministratorTempStatus,
			"bbt_adm_temp_reason":   p.BibiteAdministratorTempReason,
		},
	)
	err = db.Debug().Model(&BibiteAdministratorTemp{}).Where("bbt_adm_temp_id = ?", pid).Error
	if err != nil {
		return &BibiteAdministratorTemp{}, err
	}
	return p, nil
}

func (p *BibiteAdministratorTemp) UpdateBibiteAdministratorTempStatus(db *gorm.DB, pid uint64) (*BibiteAdministratorTemp, error) {
	var err error
	db = db.Debug().Model(&BibiteAdministratorTemp{}).Where("bbt_adm_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_adm_temp_reason": p.BibiteAdministratorTempReason,
			"bbt_adm_temp_status": p.BibiteAdministratorTempStatus,
		},
	)
	err = db.Debug().Model(&BibiteAdministratorTemp{}).Where("bbt_adm_temp_id = ?", pid).Error
	if err != nil {
		return &BibiteAdministratorTemp{}, err
	}
	return p, nil
}

func (p *BibiteAdministratorTemp) DeleteBibiteAdministratorTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&BibiteAdministratorTemp{}).Where("bbt_adm_temp_id = ?", pid).Take(&BibiteAdministratorTemp{}).Delete(&BibiteAdministratorTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("BibiteAdministratorTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADMINISTRATOR LOGIN
//====================================================================================================================================

func HashBibiteAdministrator(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (p *BibiteAdministratorTemp) BeforeSaveBibiteAdministratorTemp() error {
	hashedPassword, err := HashBibiteAdministrator(p.BibiteAdministratorTempPassword)
	if err != nil {
		return err
	}
	p.BibiteAdministratorTempPassword = string(hashedPassword)
	return nil
}

func (p *BibiteAdministratorTempPend) BeforeSaveBibiteAdministratorTemp() error {
	hashedPassword, err := HashBibiteAdministrator(p.BibiteAdministratorTempPassword)
	if err != nil {
		return err
	}
	p.BibiteAdministratorTempPassword = string(hashedPassword)
	return nil
}

//ADDITIONAL FROM CONTROLLER
//====================================================================================================================================
