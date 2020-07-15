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

type BibiteUserTemp struct {
	BibiteUserTempID           uint64    `gorm:"column:bbt_user_temp_id;primary_key;" json:"bbt_user_temp_id"`
	BibiteUserID               uint64    `gorm:"column:bbt_user_id" json:"bbt_user_id"`
	BibiteEmployeeTempID       uint64    `gorm:"column:bbt_emp_temp_id" json:"bbt_emp_temp_id"`
	BibiteGroupTempID          uint64    `gorm:"column:bbt_group_temp_id" json:"bbt_group_temp_id"`
	BibiteUserTempEmail        string    `gorm:"column:bbt_user_temp_email;size:255" json:"bbt_user_temp_email"`
	BibiteUserTempUsername     string    `gorm:"column:bbt_user_temp_username;size:125" json:"bbt_user_temp_username"`
	BibiteUserTempPassword     string    `gorm:"column:bbt_user_temp_password;size:125" json:"bbt_user_temp_password"`
	BibiteUserTempPhoneCode    string    `gorm:"column:bbt_user_temp_phone_code;size:6" json:"bbt_user_temp_phone_code"`
	BibiteUserTempPhone        string    `gorm:"column:bbt_user_temp_phone;size:20" json:"bbt_user_temp_phone"`
	BibiteUserTempPIN          string    `gorm:"column:bbt_user_temp_pin;size:255" json:"bbt_user_temp_pin"`
	BibiteUserTempReason       string    `gorm:"column:bbt_user_temp_reason" json:"bbt_user_temp_reason"`
	BibiteUserTempStatus       uint64    `gorm:"column:bbt_user_temp_status;size:2" json:"bbt_user_temp_status"`
	BibiteUserTempActionBefore uint64    `gorm:"column:bbt_user_temp_action_before" json:"bbt_user_temp_action_before"`
	BibiteUserTempCreatedBy    string    `gorm:"column:bbt_user_temp_created_by;size:125" json:"bbt_user_temp_created_by"`
	BibiteUserTempCreatedAt    time.Time `gorm:"column:bbt_user_temp_created_at;default:CURRENT_TIMESTAMP" json:"bbt_user_temp_created_at"`
}

type BibiteUserTempPend struct {
	BibiteUserTempID           uint64    `gorm:"column:bbt_user_temp_id;primary_key;" json:"bbt_user_temp_id"`
	BibiteUserID               *uint64   `gorm:"column:bbt_user_id" json:"bbt_user_id"`
	BibiteEmployeeTempID       uint64    `gorm:"column:bbt_emp_temp_id" json:"bbt_emp_temp_id"`
	BibiteGroupTempID          uint64    `gorm:"column:bbt_group_temp_id" json:"bbt_group_temp_id"`
	BibiteUserTempEmail        string    `gorm:"column:bbt_user_temp_email;size:255" json:"bbt_user_temp_email"`
	BibiteUserTempUsername     string    `gorm:"column:bbt_user_temp_username;size:125" json:"bbt_user_temp_username"`
	BibiteUserTempPassword     string    `gorm:"column:bbt_user_temp_password;size:125" json:"bbt_user_temp_password"`
	BibiteUserTempPhoneCode    string    `gorm:"column:bbt_user_temp_phone_code;size:6" json:"bbt_user_temp_phone_code"`
	BibiteUserTempPhone        string    `gorm:"column:bbt_user_temp_phone;size:20" json:"bbt_user_temp_phone"`
	BibiteUserTempPIN          string    `gorm:"column:bbt_user_temp_pin;size:255" json:"bbt_user_temp_pin"`
	BibiteUserTempReason       string    `gorm:"column:bbt_user_temp_reason" json:"bbt_user_temp_reason"`
	BibiteUserTempStatus       uint64    `gorm:"column:bbt_user_temp_status;size:2" json:"bbt_user_temp_status"`
	BibiteUserTempActionBefore uint64    `gorm:"column:bbt_user_temp_action_before" json:"bbt_user_temp_action_before"`
	BibiteUserTempCreatedBy    string    `gorm:"column:bbt_user_temp_created_by;size:125" json:"bbt_user_temp_created_by"`
	BibiteUserTempCreatedAt    time.Time `gorm:"column:bbt_user_temp_created_at;default:CURRENT_TIMESTAMP" json:"bbt_user_temp_created_at"`
}

type BibiteUserTempPendData struct {
	BibiteUserTempID             uint64    `gorm:"column:bbt_user_temp_id;primary_key;" json:"bbt_user_temp_id"`
	BibiteUserID                 *uint64   `gorm:"column:bbt_user_id" json:"bbt_user_id"`
	BibiteEmployeeTempID         uint64    `gorm:"column:bbt_emp_temp_id" json:"bbt_emp_temp_id"`
	BibiteEmployeeTempName       string    `gorm:"column:bbt_emp_temp_name" json:"bbt_emp_temp_name"`
	BibiteEmployeeTempGender     string    `gorm:"column:bbt_emp_temp_gender;size:255" json:"bbt_emp_temp_gender"`
	BibiteEmployeeTempBirthPlace string    `gorm:"column:bbt_emp_temp_birth_place;size:255" json:"bbt_emp_temp_birth_place"`
	BibiteEmployeeTempBirthDate  string    `gorm:"column:bbt_emp_temp_birth_date;size:255" json:"bbt_emp_temp_birth_date"`
	BibiteEmployeeTempAddress    string    `gorm:"column:bbt_emp_temp_address;size:255" json:"bbt_emp_temp_address"`
	AddressTempID                uint64    `gorm:"column:address_temp_id;size:255" json:"address_id"`
	CountryTempID                uint64    `gorm:"column:country_temp_id" json:"country_temp_id"`
	CountryTempName              string    `gorm:"column:country_temp_name" json:"country_temp_name"`
	ProvinceTempID               uint64    `gorm:"column:province_temp_id" json:"province_temp_id"`
	ProvinceTempName             string    `gorm:"column:province_temp_name" json:"province_temp_name"`
	RegencyTempID                uint64    `gorm:"column:regency_temp_id" json:"regency_temp_id"`
	RegencyTempName              string    `gorm:"column:regency_temp_name" json:"regency_temp_name"`
	DistrictTempID               uint64    `gorm:"column:district_temp_id" json:"district_temp_id"`
	DistrictTempName             string    `gorm:"column:district_temp_name" json:"district_temp_name"`
	VillageTempID                uint64    `gorm:"column:village_temp_id" json:"village_temp_id"`
	VillageTempName              string    `gorm:"column:village_temp_name" json:"village_temp_name"`
	BibiteEmployeeTempPhoto      string    `gorm:"column:bbt_emp_temp_photo;size:255" json:"bbt_emp_temp_photo"`
	BibiteGroupTempID            uint64    `gorm:"column:bbt_group_temp_id" json:"bbt_group_temp_id"`
	BibiteGroupTempName          string    `gorm:"column:bbt_group_temp_name" json:"bbt_group_temp_name"`
	BibiteUserTempEmail          string    `gorm:"column:bbt_user_temp_email;size:255" json:"bbt_user_temp_email"`
	BibiteUserTempUsername       string    `gorm:"column:bbt_user_temp_username;size:125" json:"bbt_user_temp_username"`
	BibiteUserTempPassword       string    `gorm:"column:bbt_user_temp_password;size:125" json:"bbt_user_temp_password"`
	BibiteUserTempPhoneCode      string    `gorm:"column:bbt_user_temp_phone_code;size:6" json:"bbt_user_temp_phone_code"`
	BibiteUserTempPhone          string    `gorm:"column:bbt_user_temp_phone;size:20" json:"bbt_user_temp_phone"`
	BibiteUserTempPIN            string    `gorm:"column:bbt_user_temp_pin;size:255" json:"bbt_user_temp_pin"`
	BibiteUserTempReason         string    `gorm:"column:bbt_user_temp_reason" json:"bbt_user_temp_reason"`
	BibiteUserTempStatus         uint64    `gorm:"column:bbt_user_temp_status;size:2" json:"bbt_user_temp_status"`
	BibiteUserTempActionBefore   uint64    `gorm:"column:bbt_user_temp_action_before" json:"bbt_user_temp_action_before"`
	BibiteUserTempCreatedBy      string    `gorm:"column:bbt_user_temp_created_by;size:125" json:"bbt_user_temp_created_by"`
	BibiteUserTempCreatedAt      time.Time `gorm:"column:bbt_user_temp_created_at;default:CURRENT_TIMESTAMP" json:"bbt_user_temp_created_at"`
}

type BibiteUserTempData struct {
	BibiteUserTempID                  uint64    `gorm:"column:bbt_user_temp_id" json:"bbt_user_temp_id"`
	BibiteEmployeeTempID              uint64    `gorm:"column:bbt_emp_temp_id" json:"bbt_emp_temp_id"`
	BibiteEmployeeTempCode            string    `gorm:"column:bbt_emp_temp_code;size:25" json:"bbt_emp_temp_code"`
	BibiteEmployeeTempName            string    `gorm:"column:bbt_emp_temp_name;size:255" json:"bbt_emp_temp_name"`
	BibiteEmployeeTempGender          string    `gorm:"column:bbt_emp_temp_gender;size:255" json:"bbt_emp_temp_gender"`
	BibiteEmployeeTempBirthPlace      string    `gorm:"column:bbt_emp_temp_birth_place;size:255" json:"bbt_emp_temp_birth_place"`
	BibiteEmployeeTempBirthDate       string    `gorm:"column:bbt_emp_temp_birth_date;size:255" json:"bbt_emp_temp_birth_date"`
	BibiteEmployeeTempAddress         string    `gorm:"column:bbt_emp_temp_address;size:255" json:"bbt_emp_temp_address"`
	AddressTempID                     uint64    `gorm:"column:address_temp_id;size:255" json:"address_id"`
	CountryTempID                     uint64    `gorm:"column:country_temp_id" json:"country_temp_id"`
	CountryTempName                   string    `gorm:"column:country_temp_name" json:"country_temp_name"`
	ProvinceTempID                    uint64    `gorm:"column:province_temp_id" json:"province_temp_id"`
	ProvinceTempName                  string    `gorm:"column:province_temp_name" json:"province_temp_name"`
	RegencyTempID                     uint64    `gorm:"column:regency_temp_id" json:"regency_temp_id"`
	RegencyTempName                   string    `gorm:"column:regency_temp_name" json:"regency_temp_name"`
	DistrictTempID                    uint64    `gorm:"column:district_temp_id" json:"district_temp_id"`
	DistrictTempName                  string    `gorm:"column:district_temp_name" json:"district_temp_name"`
	VillageTempID                     uint64    `gorm:"column:village_temp_id" json:"village_temp_id"`
	VillageTempName                   string    `gorm:"column:village_temp_name" json:"village_temp_name"`
	BibiteEmployeeTempPhoto           string    `gorm:"column:bbt_emp_temp_photo;size:255" json:"bbt_emp_temp_photo"`
	BibiteGroupTempID                 uint64    `gorm:"column:bbt_group_temp_id" json:"bbt_group_temp_id"`
	BibiteGroupTempCode               string    `gorm:"column:bbt_group_temp_code;size:25" json:"bbt_group_temp_code"`
	BibiteGroupTempName               string    `gorm:"column:bbt_group_temp_name;size:255" json:"bbt_group_temp_name"`
	BibiteGroupTempRole               string    `gorm:"column:bbt_group_temp_role;size:125" json:"bbt_group_temp_role"`
	BibiteGroupTempType               uint64    `gorm:"column:bbt_group_temp_type;size:2" json:"bbt_group_temp_type"`
	BibiteUserTempEmail               string    `gorm:"column:bbt_user_temp_email;size:255" json:"bbt_user_temp_email"`
	BibiteUserTempUsername            string    `gorm:"column:bbt_user_temp_username;size:125" json:"bbt_user_temp_username"`
	BibiteUserTempPhoneCode           string    `gorm:"column:bbt_user_temp_phone_code;size:6" json:"bbt_user_temp_phone_code"`
	BibiteUserTempPhone               string    `gorm:"column:bbt_user_temp_phone;size:20" json:"bbt_user_temp_phone"`
	BibiteUserTempReason              string    `gorm:"column:bbt_user_temp_reason" json:"bbt_user_temp_reason"`
	BibiteUserTempStatus              uint64    `gorm:"column:bbt_user_temp_status;size:2" json:"bbt_user_temp_status"`
	BibiteUserTempActionBefore        uint64    `gorm:"column:bbt_user_temp_action_before" json:"bbt_user_temp_action_before"`
	BibiteUserTempCreatedBy           string    `gorm:"column:bbt_user_temp_created_by;size:125" json:"bbt_user_temp_created_by"`
	BibiteUserTempCreatedAt           time.Time `gorm:"column:bbt_user_temp_created_at;default:CURRENT_TIMESTAMP" json:"bbt_user_temp_created_at"`
	BibiteUserID                      uint64    `gorm:"column:bbt_user_id" json:"bbt_user_id"`
	BibiteEmployeeID                  uint64    `gorm:"column:bbt_emp_id" json:"bbt_emp_id"`
	BibiteEmployeeCode                string    `gorm:"column:bbt_emp_code;size:25" json:"bbt_emp_code"`
	BibiteEmployeeName                string    `gorm:"column:bbt_emp_name;size:255" json:"bbt_emp_name"`
	BibiteEmployeeGender              string    `gorm:"column:bbt_emp_gender;size:255" json:"bbt_emp_gender"`
	BibiteEmployeeBirthPlace          string    `gorm:"column:bbt_emp_birth_place;size:255" json:"bbt_emp_birth_place"`
	BibiteEmployeeBirthDate           string    `gorm:"column:bbt_emp_birth_date;size:255" json:"bbt_emp_birth_date"`
	BibiteEmployeeAddress             string    `gorm:"column:bbt_emp_address;size:255" json:"bbt_emp_address"`
	AddressID                         uint64    `gorm:"column:address_id;size:255" json:"address_id"`
	CountryID                         uint64    `gorm:"column:country_id" json:"country_id"`
	CountryName                       string    `gorm:"column:country_name" json:"country_name"`
	ProvinceID                        uint64    `gorm:"column:province_id" json:"province_id"`
	ProvinceName                      string    `gorm:"column:province_name" json:"province_name"`
	RegencyID                         uint64    `gorm:"column:regency_id" json:"regency_id"`
	RegencyName                       string    `gorm:"column:regency_name" json:"regency_name"`
	DistrictID                        uint64    `gorm:"column:district_id" json:"district_id"`
	DistrictName                      string    `gorm:"column:district_name" json:"district_name"`
	VillageID                         uint64    `gorm:"column:village_id" json:"village_id"`
	VillageName                       string    `gorm:"column:village_name" json:"village_name"`
	BibiteEmployeePhoto               string    `gorm:"column:bbt_emp_photo;size:255" json:"bbt_emp_photo"`
	BibiteGroupID                     uint64    `gorm:"column:bbt_group_id" json:"bbt_group_id"`
	BibiteGroupCode                   string    `gorm:"column:bbt_group_code;size:25" json:"bbt_group_code"`
	BibiteGroupName                   string    `gorm:"column:bbt_group_name;size:255" json:"bbt_group_name"`
	BibiteGroupRole                   string    `gorm:"column:bbt_group_role;size:125" json:"bbt_group_role"`
	BibiteGroupType                   uint64    `gorm:"column:bbt_group_type;size:2" json:"bbt_group_type"`
	BibiteUserEmail                   string    `gorm:"column:bbt_user_email" json:"bbt_user_email"`
	BibiteUsername                    string    `gorm:"column:bbt_username" json:"bbt_username"`
	BibiteUserPhoneCode               string    `gorm:"column:bbt_user_phone_code;size:6" json:"bbt_user_phone_code"`
	BibiteUserPhone                   string    `gorm:"column:bbt_user_phone;size:20" json:"bbt_user_phone"`
	BibiteUserStatus                  uint64    `gorm:"column:bbt_user_status;size:1" json:"bbt_user_status"`
	BibiteUserIsNew                   bool      `gorm:"column:bbt_user_is_new" json:"bbt_user_is_new"`
	BibiteUserIsRequestChangePassword bool      `gorm:"column:bbt_user_is_req_chpw" json:"bbt_user_is_req_chpw"`
	BibiteUserIsLocked                bool      `gorm:"column:bbt_user_is_locked" json:"bbt_user_is_locked"`
	BibiteUserFailAttempt             uint64    `gorm:"column:bbt_user_fail_attempt;size:1" json:"bbt_user_fail_attempt"`
	BibiteUserLockedReason            string    `gorm:"column:bbt_user_locked_reason;size:255" json:"bbt_user_locked_reason"`
	BibiteUserCreatedBy               string    `gorm:"column:bbt_user_created_by;size:125" json:"bbt_user_created_by"`
	BibiteUserCreatedAt               time.Time `gorm:"column:bbt_user_created_at;default:CURRENT_TIMESTAMP" json:"bbt_user_created_at"`
}

type ResponseBibiteUserTemps struct {
	Status  int                   `json:"status"`
	Message string                `json:"message"`
	Data    *[]BibiteUserTempData `json:"data"`
}

type ResponseBibiteUserTemp struct {
	Status  int                 `json:"status"`
	Message string              `json:"message"`
	Data    *BibiteUserTempData `json:"data"`
}

type ResponseBibiteUserTempsPend struct {
	Status  int                       `json:"status"`
	Message string                    `json:"message"`
	Data    *[]BibiteUserTempPendData `json:"data"`
}

type ResponseBibiteUserTempPend struct {
	Status  int                     `json:"status"`
	Message string                  `json:"message"`
	Data    *BibiteUserTempPendData `json:"data"`
}

type ResponseBibiteUserTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (BibiteUserTemp) TableName() string {
	return "m_bibite_user_temp"
}

func (BibiteUserTempPend) TableName() string {
	return "m_bibite_user_temp"
}

func (BibiteUserTempPendData) TableName() string {
	return "m_bibite_user_temp"
}

func (BibiteUserTempData) TableName() string {
	return "m_bibite_user_temp"
}

func (p *BibiteUserTemp) Prepare() {
	p.BibiteUserID = p.BibiteUserID
	p.BibiteEmployeeTempID = p.BibiteEmployeeTempID
	p.BibiteGroupTempID = p.BibiteGroupTempID
	p.BibiteUserTempEmail = html.EscapeString(strings.TrimSpace(p.BibiteUserTempEmail))
	p.BibiteUserTempUsername = html.EscapeString(strings.TrimSpace(p.BibiteUserTempUsername))
	p.BibiteUserTempPassword = p.BibiteUserTempPassword
	p.BibiteUserTempPhoneCode = html.EscapeString(strings.TrimSpace(p.BibiteUserTempPhoneCode))
	p.BibiteUserTempPhone = html.EscapeString(strings.TrimSpace(p.BibiteUserTempPhone))
	p.BibiteUserTempPIN = html.EscapeString(strings.TrimSpace(p.BibiteUserTempPIN))
	p.BibiteUserTempPIN = html.EscapeString(strings.TrimSpace(p.BibiteUserTempPIN))
	p.BibiteUserTempReason = p.BibiteUserTempReason
	p.BibiteUserTempCreatedBy = p.BibiteUserTempCreatedBy
	p.BibiteUserTempCreatedAt = time.Now()
}

func (p *BibiteUserTempPend) Prepare() {
	p.BibiteUserID = nil
	p.BibiteEmployeeTempID = p.BibiteEmployeeTempID
	p.BibiteGroupTempID = p.BibiteGroupTempID
	p.BibiteUserTempEmail = html.EscapeString(strings.TrimSpace(p.BibiteUserTempEmail))
	p.BibiteUserTempUsername = html.EscapeString(strings.TrimSpace(p.BibiteUserTempUsername))
	p.BibiteUserTempPassword = p.BibiteUserTempPassword
	p.BibiteUserTempPhoneCode = html.EscapeString(strings.TrimSpace(p.BibiteUserTempPhoneCode))
	p.BibiteUserTempPhone = html.EscapeString(strings.TrimSpace(p.BibiteUserTempPhone))
	p.BibiteUserTempPIN = html.EscapeString(strings.TrimSpace(p.BibiteUserTempPIN))
	p.BibiteUserTempPIN = html.EscapeString(strings.TrimSpace(p.BibiteUserTempPIN))
	p.BibiteUserTempReason = p.BibiteUserTempReason
	p.BibiteUserTempCreatedBy = p.BibiteUserTempCreatedBy
	p.BibiteUserTempCreatedAt = time.Now()
}

func (p *BibiteUserTempPend) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.BibiteUserTempPassword == "" {
			return errors.New("Required BibiteUserTemp Password")
		}
		if p.BibiteUserTempEmail == "" {
			return errors.New("Required BibiteUserTemp Email")
		}
		return nil
	}
}

func (p *BibiteUserTemp) SaveBibiteUserTemp(db *gorm.DB) (*BibiteUserTemp, error) {
	var err error
	err = p.BeforeSaveBibiteUser()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Debug().Model(&BibiteUserTemp{}).Create(&p).Error
	if err != nil {
		return &BibiteUserTemp{}, err
	}
	return p, nil
}

func (p *BibiteUserTempPend) SaveBibiteUserTempPend(db *gorm.DB) (*BibiteUserTempPend, error) {
	var err error
	err = p.BeforeSaveBibiteUser()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Debug().Model(&BibiteUserTempPend{}).Create(&p).Error
	if err != nil {
		return &BibiteUserTempPend{}, err
	}
	return p, nil
}

func (p *BibiteUserTemp) FindAllBibiteUserTemp(db *gorm.DB) (*[]BibiteUserTempPendData, error) {
	var err error
	bibiteuser := []BibiteUserTempPendData{}
	err = db.Debug().Model(&BibiteUserTempPendData{}).Limit(100).
		Select(`m_bibite_user_temp.bbt_user_temp_id,
				m_bibite_user_temp.bbt_emp_temp_id,
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
				m_bibite_user_temp.bbt_group_temp_id,
				m_bibite_group_temp.bbt_group_name as bbt_group_temp_name,
				m_bibite_user_temp.bbt_user_temp_email,
				m_bibite_user_temp.bbt_user_temp_username,
				m_bibite_user_temp.bbt_user_temp_phone_code,
				m_bibite_user_temp.bbt_user_temp_phone,
				m_bibite_user_temp.bbt_user_temp_reason,
				m_bibite_user_temp.bbt_user_temp_status,
				m_bibite_user_temp.bbt_user_temp_action_before,
				m_bibite_user_temp.bbt_user_temp_created_by,
				m_bibite_user_temp.bbt_user_temp_created_at at time zone 'Asia/Jakarta' as bbt_user_temp_created_at`).
		Joins("JOIN m_bibite_employee_temp on m_bibite_user_temp.bbt_emp_temp_id=m_bibite_employee_temp.bbt_emp_temp_id").
		Joins("JOIN m_address_temp on m_bibite_employee_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_bibite_group m_bibite_group_temp on m_bibite_user_temp.bbt_group_temp_id=m_bibite_group_temp.bbt_group_id").
		Order("bbt_user_temp_created_at desc").
		Find(&bibiteuser).Error
	if err != nil {
		return &[]BibiteUserTempPendData{}, err
	}
	return &bibiteuser, nil
}

func (p *BibiteUserTemp) FindAllBibiteUserTempStatusPendingActive(db *gorm.DB) (*[]BibiteUserTempPendData, error) {
	var err error
	bibiteuser := []BibiteUserTempPendData{}
	err = db.Debug().Model(&BibiteUserTempPendData{}).Limit(100).
		Select(`m_bibite_user_temp.bbt_user_temp_id,
				m_bibite_user_temp.bbt_emp_temp_id,
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
				m_bibite_user_temp.bbt_group_temp_id,
				m_bibite_group_temp.bbt_group_name as bbt_group_temp_name,
				m_bibite_user_temp.bbt_user_temp_email,
				m_bibite_user_temp.bbt_user_temp_username,
				m_bibite_user_temp.bbt_user_temp_phone_code,
				m_bibite_user_temp.bbt_user_temp_phone,
				m_bibite_user_temp.bbt_user_temp_reason,
				m_bibite_user_temp.bbt_user_temp_status,
				m_bibite_user_temp.bbt_user_temp_action_before,
				m_bibite_user_temp.bbt_user_temp_created_by,
				m_bibite_user_temp.bbt_user_temp_created_at at time zone 'Asia/Jakarta' as bbt_user_temp_created_at`).
		Joins("JOIN m_bibite_employee_temp on m_bibite_user_temp.bbt_emp_temp_id=m_bibite_employee_temp.bbt_emp_temp_id").
		Joins("JOIN m_address_temp on m_bibite_employee_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_bibite_group m_bibite_group_temp on m_bibite_user_temp.bbt_group_temp_id=m_bibite_group_temp.bbt_group_id").
		Where("bbt_user_temp_status = ?", 11).
		Order("bbt_user_temp_created_at desc").
		Find(&bibiteuser).Error
	if err != nil {
		return &[]BibiteUserTempPendData{}, err
	}
	return &bibiteuser, nil
}

func (p *BibiteUserTemp) FindAllBibiteUserTempStatus(db *gorm.DB, status uint64) (*[]BibiteUserTempData, error) {
	var err error
	bibiteuser := []BibiteUserTempData{}
	err = db.Debug().Model(&BibiteUserTempData{}).Limit(100).
		Select(`m_bibite_user_temp.bbt_user_temp_id,
				m_bibite_user_temp.bbt_emp_temp_id,
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
				m_bibite_user_temp.bbt_group_temp_id,
				m_bibite_group_temp.bbt_group_code as bbt_group_temp_code,
				m_bibite_group_temp.bbt_group_name as bbt_group_temp_name,
				m_bibite_group_temp.bbt_group_role as bbt_group_temp_role,
				m_bibite_group_temp.bbt_group_type as bbt_group_temp_type,
				m_bibite_user_temp.bbt_user_temp_email,
				m_bibite_user_temp.bbt_user_temp_username,
				m_bibite_user_temp.bbt_user_temp_phone_code,
				m_bibite_user_temp.bbt_user_temp_phone,
				m_bibite_user_temp.bbt_user_temp_reason,
				m_bibite_user_temp.bbt_user_temp_status,
				m_bibite_user_temp.bbt_user_temp_action_before,
				m_bibite_user_temp.bbt_user_temp_created_by,
				m_bibite_user_temp.bbt_user_temp_created_at at time zone 'Asia/Jakarta' as bbt_user_temp_created_at,
				m_bibite_user.bbt_user_id,
				m_bibite_user.bbt_emp_id,
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
				m_bibite_user.bbt_user_email,
				m_bibite_user.bbt_username,
				m_bibite_user.bbt_user_phone_code,
				m_bibite_user.bbt_user_phone,
				m_bibite_user.bbt_user_status,
				m_bibite_user.bbt_user_is_new,
				m_bibite_user.bbt_user_is_req_chpw,
				m_bibite_user.bbt_user_is_locked,
				m_bibite_user.bbt_user_fail_attempt,
				m_bibite_user.bbt_user_locked_reason,
				m_bibite_user.bbt_user_created_by,
				m_bibite_user.bbt_user_created_at at time zone 'Asia/Jakarta' as bbt_user_created_at`).
		Joins("JOIN m_bibite_employee_temp on m_bibite_user_temp.bbt_emp_temp_id=m_bibite_employee_temp.bbt_emp_temp_id").
		Joins("JOIN m_address_temp on m_bibite_employee_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_bibite_group m_bibite_group_temp on m_bibite_user_temp.bbt_group_temp_id=m_bibite_group_temp.bbt_group_id").
		Joins("JOIN m_bibite_user on m_bibite_user_temp.bbt_user_id=m_bibite_user.bbt_user_id").
		Joins("JOIN m_bibite_employee on m_bibite_user.bbt_emp_id=m_bibite_employee.bbt_emp_id").
		Joins("JOIN m_address on m_bibite_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Joins("JOIN m_bibite_group on m_bibite_user.bbt_group_id=m_bibite_group.bbt_group_id").
		Where("bbt_user_temp_status = ?", status).
		Order("bbt_user_temp_created_at desc").
		Find(&bibiteuser).Error
	if err != nil {
		return &[]BibiteUserTempData{}, err
	}
	return &bibiteuser, nil
}

func (p *BibiteUserTemp) FindBibiteUserTempDataByID(db *gorm.DB, pid uint64) (*BibiteUserTemp, error) {
	var err error
	err = db.Debug().Model(&BibiteUserTemp{}).Where("bbt_user_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &BibiteUserTemp{}, err
	}
	return p, nil
}

func (p *BibiteUserTemp) FindBibiteUserTempByIDPendingActive(db *gorm.DB, pid uint64) (*BibiteUserTempPendData, error) {
	var err error
	bibiteuser := BibiteUserTempPendData{}
	err = db.Debug().Model(&BibiteUserTempPendData{}).Limit(100).
		Select(`m_bibite_user_temp.bbt_user_temp_id,
				m_bibite_user_temp.bbt_emp_temp_id,
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
				m_bibite_user_temp.bbt_group_temp_id,
				m_bibite_group_temp.bbt_group_name as bbt_group_temp_name,
				m_bibite_user_temp.bbt_user_temp_email,
				m_bibite_user_temp.bbt_user_temp_username,
				m_bibite_user_temp.bbt_user_temp_phone_code,
				m_bibite_user_temp.bbt_user_temp_phone,
				m_bibite_user_temp.bbt_user_temp_reason,
				m_bibite_user_temp.bbt_user_temp_status,
				m_bibite_user_temp.bbt_user_temp_action_before,
				m_bibite_user_temp.bbt_user_temp_created_by,
				m_bibite_user_temp.bbt_user_temp_created_at at time zone 'Asia/Jakarta' as bbt_user_temp_created_at`).
		Joins("JOIN m_bibite_employee_temp on m_bibite_user_temp.bbt_emp_temp_id=m_bibite_employee_temp.bbt_emp_temp_id").
		Joins("JOIN m_address_temp on m_bibite_employee_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_bibite_group m_bibite_group_temp on m_bibite_user_temp.bbt_group_temp_id=m_bibite_group_temp.bbt_group_id").
		Where("bbt_user_temp_id = ?", pid).
		Order("bbt_user_temp_created_at desc").
		Find(&bibiteuser).Error
	if err != nil {
		return &BibiteUserTempPendData{}, err
	}
	return &bibiteuser, nil
}

func (p *BibiteUserTempPendAct) FindBibiteUserTempStatusByIDPendAll(db *gorm.DB, pid uint64) (*BibiteUserTempPendAct, error) {
	var err error
	bibiteuser := BibiteUserTempPendAct{}
	err = db.Debug().Model(&BibiteUserTempPendAct{}).Limit(100).
		Select(`m_bibite_user_temp.bbt_user_temp_id,
				m_bibite_user_temp.bbt_emp_temp_id,
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
				m_bibite_user_temp.bbt_group_temp_id,
				m_bibite_group_temp.bbt_group_name as bbt_group_temp_name,
				m_bibite_user_temp.bbt_user_temp_email,
				m_bibite_user_temp.bbt_user_temp_username,
				m_bibite_user_temp.bbt_user_temp_phone_code,
				m_bibite_user_temp.bbt_user_temp_phone,
				m_bibite_user_temp.bbt_user_temp_reason,
				m_bibite_user_temp.bbt_user_temp_status,
				m_bibite_user_temp.bbt_user_temp_action_before,
				m_bibite_user_temp.bbt_user_temp_created_by,
				m_bibite_user_temp.bbt_user_temp_created_at at time zone 'Asia/Jakarta' as bbt_user_temp_created_at`).
		Joins("JOIN m_bibite_employee_temp on m_bibite_user_temp.bbt_emp_temp_id=m_bibite_employee_temp.bbt_emp_temp_id").
		Joins("JOIN m_address_temp on m_bibite_employee_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_bibite_group m_bibite_group_temp on m_bibite_user_temp.bbt_group_temp_id=m_bibite_group_temp.bbt_group_id").
		Where("bbt_user_temp_id = ? AND bbt_user_temp_status IN (11,12,13,14)", pid).
		Order("bbt_user_temp_created_at desc").
		Find(&bibiteuser).Error
	if err != nil {
		return &BibiteUserTempPendAct{}, err
	}
	return &bibiteuser, nil
}

func (p *BibiteUserTemp) FindBibiteUserTempByID(db *gorm.DB, pid uint64) (*BibiteUserTempData, error) {
	var err error
	bibiteuser := BibiteUserTempData{}
	err = db.Debug().Model(&BibiteUserTempData{}).
		Select(`m_bibite_user_temp.bbt_user_temp_id,
				m_bibite_user_temp.bbt_emp_temp_id,
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
				m_bibite_user_temp.bbt_group_temp_id,
				m_bibite_group_temp.bbt_group_code as bbt_group_temp_code,
				m_bibite_group_temp.bbt_group_name as bbt_group_temp_name,
				m_bibite_group_temp.bbt_group_role as bbt_group_temp_role,
				m_bibite_group_temp.bbt_group_type as bbt_group_temp_type,
				m_bibite_user_temp.bbt_user_temp_email,
				m_bibite_user_temp.bbt_user_temp_username,
				m_bibite_user_temp.bbt_user_temp_phone_code,
				m_bibite_user_temp.bbt_user_temp_phone,
				m_bibite_user_temp.bbt_user_temp_reason,
				m_bibite_user_temp.bbt_user_temp_status,
				m_bibite_user_temp.bbt_user_temp_action_before,
				m_bibite_user_temp.bbt_user_temp_created_by,
				m_bibite_user_temp.bbt_user_temp_created_at at time zone 'Asia/Jakarta' as bbt_user_temp_created_at,
				m_bibite_user.bbt_user_id,
				m_bibite_user.bbt_emp_id,
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
				m_bibite_user.bbt_user_email,
				m_bibite_user.bbt_username,
				m_bibite_user.bbt_user_phone_code,
				m_bibite_user.bbt_user_phone,
				m_bibite_user.bbt_user_status,
				m_bibite_user.bbt_user_is_new,
				m_bibite_user.bbt_user_is_req_chpw,
				m_bibite_user.bbt_user_is_locked,
				m_bibite_user.bbt_user_fail_attempt,
				m_bibite_user.bbt_user_locked_reason,
				m_bibite_user.bbt_user_created_by,
				m_bibite_user.bbt_user_created_at at time zone 'Asia/Jakarta' as bbt_user_created_at`).
		Joins("JOIN m_bibite_employee_temp on m_bibite_user_temp.bbt_emp_temp_id=m_bibite_employee_temp.bbt_emp_temp_id").
		Joins("JOIN m_address_temp on m_bibite_employee_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_bibite_group m_bibite_group_temp on m_bibite_user_temp.bbt_group_temp_id=m_bibite_group_temp.bbt_group_id").
		Joins("JOIN m_bibite_user on m_bibite_user_temp.bbt_user_id=m_bibite_user.bbt_user_id").
		Joins("JOIN m_bibite_employee on m_bibite_user.bbt_emp_id=m_bibite_employee.bbt_emp_id").
		Joins("JOIN m_address on m_bibite_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Joins("JOIN m_bibite_group on m_bibite_user.bbt_group_id=m_bibite_group.bbt_group_id").
		Where("bbt_user_temp_id = ?", pid).
		Order("bbt_user_temp_created_at desc").
		Take(&bibiteuser).Error
	if err != nil {
		return &BibiteUserTempData{}, err
	}
	return &bibiteuser, nil
}

func (p *BibiteUserTemp) FindBibiteUserTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*BibiteUserTempData, error) {
	var err error
	bibiteuser := BibiteUserTempData{}
	err = db.Debug().Model(&BibiteUserTempData{}).
		Select(`m_bibite_user_temp.bbt_user_temp_id,
				m_bibite_user_temp.bbt_emp_temp_id,
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
				m_bibite_user_temp.bbt_group_temp_id,
				m_bibite_group_temp.bbt_group_code as bbt_group_temp_code,
				m_bibite_group_temp.bbt_group_name as bbt_group_temp_name,
				m_bibite_group_temp.bbt_group_role as bbt_group_temp_role,
				m_bibite_group_temp.bbt_group_type as bbt_group_temp_type,
				m_bibite_user_temp.bbt_user_temp_email,
				m_bibite_user_temp.bbt_user_temp_username,
				m_bibite_user_temp.bbt_user_temp_phone_code,
				m_bibite_user_temp.bbt_user_temp_phone,
				m_bibite_user_temp.bbt_user_temp_reason,
				m_bibite_user_temp.bbt_user_temp_status,
				m_bibite_user_temp.bbt_user_temp_action_before,
				m_bibite_user_temp.bbt_user_temp_created_by,
				m_bibite_user_temp.bbt_user_temp_created_at at time zone 'Asia/Jakarta' as bbt_user_temp_created_at,
				m_bibite_user.bbt_user_id,
				m_bibite_user.bbt_emp_id,
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
				m_bibite_user.bbt_user_email,
				m_bibite_user.bbt_username,
				m_bibite_user.bbt_user_phone_code,
				m_bibite_user.bbt_user_phone,
				m_bibite_user.bbt_user_status,
				m_bibite_user.bbt_user_is_new,
				m_bibite_user.bbt_user_is_req_chpw,
				m_bibite_user.bbt_user_is_locked,
				m_bibite_user.bbt_user_fail_attempt,
				m_bibite_user.bbt_user_locked_reason,
				m_bibite_user.bbt_user_created_by,
				m_bibite_user.bbt_user_created_at at time zone 'Asia/Jakarta' as bbt_user_created_at`).
		Joins("JOIN m_bibite_employee_temp on m_bibite_user_temp.bbt_emp_temp_id=m_bibite_employee_temp.bbt_emp_temp_id").
		Joins("JOIN m_address_temp on m_bibite_employee_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_bibite_group m_bibite_group_temp on m_bibite_user_temp.bbt_group_temp_id=m_bibite_group_temp.bbt_group_id").
		Joins("JOIN m_bibite_user on m_bibite_user_temp.bbt_user_id=m_bibite_user.bbt_user_id").
		Joins("JOIN m_bibite_employee on m_bibite_user.bbt_emp_id=m_bibite_employee.bbt_emp_id").
		Joins("JOIN m_address on m_bibite_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Joins("JOIN m_bibite_group on m_bibite_user.bbt_group_id=m_bibite_group.bbt_group_id").
		Where("bbt_user_temp_id = ? AND bbt_user_temp_status = ?", pid, status).
		Order("bbt_user_temp_created_at desc").
		Take(&bibiteuser).Error
	if err != nil {
		return &BibiteUserTempData{}, err
	}
	return &bibiteuser, nil
}

func (p *BibiteUserTemp) UpdateBibiteUserTemp(db *gorm.DB, pid uint64) (*BibiteUserTemp, error) {

	var err error
	db = db.Debug().Model(&BibiteUserTemp{}).Where("bbt_user_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_emp_temp_id":          p.BibiteEmployeeTempID,
			"bbt_group_temp_id":        p.BibiteGroupTempID,
			"bbt_user_temp_email":      p.BibiteUserTempEmail,
			"bbt_user_temp_username":   p.BibiteUserTempUsername,
			"bbt_user_temp_reason":     p.BibiteUserTempReason,
			"bbt_user_temp_status":     p.BibiteUserTempStatus,
			"bbt_user_temp_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&BibiteUserTemp{}).Where("bbt_user_temp_id = ?", pid).Error
	if err != nil {
		return &BibiteUserTemp{}, err
	}
	return p, nil
}

func (p *BibiteUserTemp) UpdateBibiteUserTempStatus(db *gorm.DB, pid uint64) (*BibiteUserTemp, error) {

	var err error
	db = db.Debug().Model(&BibiteUserTemp{}).Where("bbt_user_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_user_temp_status":     p.BibiteUserTempStatus,
			"bbt_user_temp_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&BibiteUserTemp{}).Where("bbt_user_temp_id = ?", pid).Error
	if err != nil {
		return &BibiteUserTemp{}, err
	}
	return p, nil
}

func (p *BibiteUserTemp) DeleteBibiteUserTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&BibiteUserTemp{}).Where("bbt_user_temp_id = ?", pid).Take(&BibiteUserTemp{}).Delete(&BibiteUserTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("BibiteUserTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//BIBITE LOGIN USER
//====================================================================================================================================

func HashBibiteUser(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (p *BibiteUserTemp) BeforeSaveBibiteUser() error {
	hashedPassword, err := HashBibiteUser(p.BibiteUserTempPassword)
	if err != nil {
		return err
	}
	p.BibiteUserTempPassword = string(hashedPassword)
	return nil
}

func (p *BibiteUserTempPend) BeforeSaveBibiteUser() error {
	hashedPassword, err := Hash(p.BibiteUserTempPassword)
	if err != nil {
		return err
	}
	p.BibiteUserTempPassword = string(hashedPassword)
	return nil
}

//ADDITIONAL FROM CONTROLLER
//====================================================================================================================================

type BibiteUserTempAct struct {
	BibiteUserTempID           uint64    `gorm:"column:bbt_user_temp_id;primary_key;" json:"bbt_user_temp_id"`
	BibiteUserID               uint64    `gorm:"column:bbt_user_id" json:"bbt_user_id"`
	BibiteEmployeeTempID       uint64    `gorm:"column:bbt_emp_temp_id" json:"bbt_emp_temp_id"`
	BibiteGroupTempID          uint64    `gorm:"column:bbt_group_temp_id" json:"bbt_group_temp_id"`
	BibiteEmployeeTempName     string    `gorm:"column:bbt_emp_temp_name" json:"bbt_emp_temp_name"`
	BibiteUserTempEmail        string    `gorm:"column:bbt_user_temp_email;size:255" json:"bbt_user_temp_email"`
	BibiteUserTempUsername     string    `gorm:"column:bbt_user_temp_username;size:125" json:"bbt_user_temp_username"`
	BibiteUserTempPassword     string    `gorm:"column:bbt_user_temp_password;size:125" json:"bbt_user_temp_password"`
	BibiteUserTempPhoneCode    string    `gorm:"column:bbt_user_temp_phone_code;size:6" json:"bbt_user_temp_phone_code"`
	BibiteUserTempPhone        string    `gorm:"column:bbt_user_temp_phone;size:20" json:"bbt_user_temp_phone"`
	BibiteUserTempPIN          string    `gorm:"column:bbt_user_temp_pin;size:255" json:"bbt_user_temp_pin"`
	BibiteUserTempReason       string    `gorm:"column:bbt_user_temp_reason" json:"bbt_user_temp_reason"`
	BibiteUserTempStatus       uint64    `gorm:"column:bbt_user_temp_status;size:2" json:"bbt_user_temp_status"`
	BibiteUserTempActionBefore uint64    `gorm:"column:bbt_user_temp_action_before" json:"bbt_user_temp_action_before"`
	BibiteUserTempCreatedBy    string    `gorm:"column:bbt_user_temp_created_by;size:125" json:"bbt_user_temp_created_by"`
	BibiteUserTempCreatedAt    time.Time `gorm:"column:bbt_user_temp_created_at;default:CURRENT_TIMESTAMP" json:"bbt_user_temp_created_at"`
}

type BibiteUserTempPendAct struct {
	BibiteUserTempID           uint64    `gorm:"column:bbt_user_temp_id;primary_key;" json:"bbt_user_temp_id"`
	BibiteUserID               *uint64   `gorm:"column:bbt_user_id" json:"bbt_user_id"`
	BibiteEmployeeTempID       uint64    `gorm:"column:bbt_emp_temp_id" json:"bbt_emp_temp_id"`
	BibiteGroupTempID          uint64    `gorm:"column:bbt_group_temp_id" json:"bbt_group_temp_id"`
	BibiteEmployeeTempName     string    `gorm:"column:bbt_emp_temp_name" json:"bbt_emp_temp_name"`
	BibiteUserTempEmail        string    `gorm:"column:bbt_user_temp_email;size:255" json:"bbt_user_temp_email"`
	BibiteUserTempUsername     string    `gorm:"column:bbt_user_temp_username;size:125" json:"bbt_user_temp_username"`
	BibiteUserTempPassword     string    `gorm:"column:bbt_user_temp_password;size:125" json:"bbt_user_temp_password"`
	BibiteUserTempPhoneCode    string    `gorm:"column:bbt_user_temp_phone_code;size:6" json:"bbt_user_temp_phone_code"`
	BibiteUserTempPhone        string    `gorm:"column:bbt_user_temp_phone;size:20" json:"bbt_user_temp_phone"`
	BibiteUserTempPIN          string    `gorm:"column:bbt_user_temp_pin;size:255" json:"bbt_user_temp_pin"`
	BibiteUserTempReason       string    `gorm:"column:bbt_user_temp_reason" json:"bbt_user_temp_reason"`
	BibiteUserTempStatus       uint64    `gorm:"column:bbt_user_temp_status;size:2" json:"bbt_user_temp_status"`
	BibiteUserTempActionBefore uint64    `gorm:"column:bbt_user_temp_action_before" json:"bbt_user_temp_action_before"`
	BibiteUserTempCreatedBy    string    `gorm:"column:bbt_user_temp_created_by;size:125" json:"bbt_user_temp_created_by"`
	BibiteUserTempCreatedAt    time.Time `gorm:"column:bbt_user_temp_created_at;default:CURRENT_TIMESTAMP" json:"bbt_user_temp_created_at"`
}

type ResponseBibiteUserTempPendAct struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    *BibiteUserTempPendAct `json:"data"`
}

func (BibiteUserTempPendAct) TableName() string {
	return "m_bibite_user_temp"
}

func (p *BibiteUserTemp) FindBibiteUserTempStatusByIDPendingActive(db *gorm.DB, pid uint64) (*BibiteUserTempPendData, error) {
	var err error
	bibiteuser := BibiteUserTempPendData{}
	err = db.Debug().Model(&BibiteUserTempPendData{}).Limit(100).
		Select(`m_bibite_user_temp.bbt_user_temp_id,
				m_bibite_user_temp.bbt_emp_temp_id,
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
				m_bibite_user_temp.bbt_group_temp_id,
				m_bibite_group_temp.bbt_group_name as bbt_group_temp_name,
				m_bibite_user_temp.bbt_user_temp_email,
				m_bibite_user_temp.bbt_user_temp_username,
				m_bibite_user_temp.bbt_user_temp_phone_code,
				m_bibite_user_temp.bbt_user_temp_phone,
				m_bibite_user_temp.bbt_user_temp_reason,
				m_bibite_user_temp.bbt_user_temp_status,
				m_bibite_user_temp.bbt_user_temp_action_before,
				m_bibite_user_temp.bbt_user_temp_created_by,
				m_bibite_user_temp.bbt_user_temp_created_at at time zone 'Asia/Jakarta' as bbt_user_temp_created_at`).
		Joins("JOIN m_bibite_employee_temp on m_bibite_user_temp.bbt_emp_temp_id=m_bibite_employee_temp.bbt_emp_temp_id").
		Joins("JOIN m_address_temp on m_bibite_employee_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_bibite_group m_bibite_group_temp on m_bibite_user_temp.bbt_group_temp_id=m_bibite_group_temp.bbt_group_id").
		Where("bbt_user_temp_id = ? AND bbt_user_temp_status = ?", pid, 11).
		Order("bbt_user_temp_created_at desc").
		Find(&bibiteuser).Error
	if err != nil {
		return &BibiteUserTempPendData{}, err
	}
	return &bibiteuser, nil
}

func (p *BibiteUserTempPendAct) FindBibiteUserTempStatusByIDPendActive(db *gorm.DB, pid uint64, status uint64) (*BibiteUserTempPendAct, error) {
	var err error
	bibiteuser := BibiteUserTempPendAct{}
	err = db.Debug().Model(&BibiteUserTempPendAct{}).Limit(100).
		Select(`m_bibite_user_temp.bbt_user_temp_id,
				m_bibite_user_temp.bbt_emp_temp_id,
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
				m_bibite_user_temp.bbt_group_temp_id,
				m_bibite_group_temp.bbt_group_name as bbt_group_temp_name,
				m_bibite_user_temp.bbt_user_temp_email,
				m_bibite_user_temp.bbt_user_temp_username,
				m_bibite_user_temp.bbt_user_temp_phone_code,
				m_bibite_user_temp.bbt_user_temp_phone,
				m_bibite_user_temp.bbt_user_temp_reason,
				m_bibite_user_temp.bbt_user_temp_status,
				m_bibite_user_temp.bbt_user_temp_action_before,
				m_bibite_user_temp.bbt_user_temp_created_by,
				m_bibite_user_temp.bbt_user_temp_created_at at time zone 'Asia/Jakarta' as bbt_user_temp_created_at`).
		Joins("JOIN m_bibite_employee_temp on m_bibite_user_temp.bbt_emp_temp_id=m_bibite_employee_temp.bbt_emp_temp_id").
		Joins("JOIN m_address_temp on m_bibite_employee_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_bibite_group m_bibite_group_temp on m_bibite_user_temp.bbt_group_temp_id=m_bibite_group_temp.bbt_group_id").
		Where("bbt_user_temp_id = ? AND bbt_user_temp_status = ?", pid, status).
		Order("bbt_user_temp_created_at desc").
		Find(&bibiteuser).Error
	if err != nil {
		return &BibiteUserTempPendAct{}, err
	}
	return &bibiteuser, nil
}
