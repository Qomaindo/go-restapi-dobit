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

type BibiteAdministrator struct {
	BibiteAdministratorID                      uint64     `gorm:"column:bbt_adm_id;primary_key;" json:"bbt_adm_id"`
	BibiteEmployeeID                           uint64     `gorm:"column:bbt_emp_id" json:"bbt_emp_id"`
	BibiteGroupID                              uint64     `gorm:"column:bbt_group_id" json:"bbt_group_id"`
	BibiteAdministratorEmail                   string     `gorm:"column:bbt_adm_email" json:"bbt_adm_email"`
	BibiteAdministratorUsername                string     `gorm:"column:bbt_adm_username" json:"bbt_adm_username"`
	BibiteAdministratorPassword                string     `gorm:"column:bbt_adm_password" json:"bbt_adm_password"`
	BibiteAdministratorPhoneCode               string     `gorm:"column:bbt_adm_phone_code" json:"bbt_adm_phone_code"`
	BibiteAdministratorPhone                   string     `gorm:"column:bbt_adm_phone" json:"bbt_adm_phone"`
	BibiteAdministratorOTP                     string     `gorm:"column:bbt_adm_otp" json:"bbt_adm_otp"`
	BibiteAdministratorOTPSecret               *time.Time `gorm:"column:bbt_adm_otp_secret" json:"bbt_adm_otp_secret"`
	BibiteAdministratorSecret                  uint64     `gorm:"column:bbt_adm_secret" json:"bbt_adm_secret"`
	BibiteAdministratorPIN                     string     `gorm:"column:bbt_adm_pin" json:"bbt_adm_pin"`
	BibiteAdministratorLastLogin               *time.Time `gorm:"column:bbt_adm_last_login" json:"bbt_adm_last_login"`
	BibiteAdministratorIPAddress               string     `gorm:"column:bbt_adm_ip_address" json:"bbt_adm_ip_address"`
	BibiteAdministratorStatus                  uint64     `gorm:"column:bbt_adm_status" json:"bbt_adm_status"`
	BibiteAdministratorIsNew                   bool       `gorm:"column:bbt_adm_is_new" json:"bbt_adm_is_new"`
	BibiteAdministratorIsRequestChangePassword bool       `gorm:"column:bbt_adm_is_req_chpw" json:"bbt_adm_is_req_chpw"`
	BibiteAdministratorIsLocked                bool       `gorm:"column:bbt_adm_is_locked" json:"bbt_adm_is_locked"`
	BibiteAdministratorFailAttempt             uint64     `gorm:"column:bbt_adm_fail_attempt" json:"bbt_adm_fail_attempt"`
	BibiteAdministratorLockedReason            string     `gorm:"column:bbt_adm_locked_reason" json:"bbt_adm_locked_reason"`
	BibiteAdministratorCreatedBy               string     `gorm:"column:bbt_adm_created_by;size:125" json:"bbt_adm_created_by"`
	BibiteAdministratorCreatedAt               time.Time  `gorm:"column:bbt_adm_created_at;default:CURRENT_TIMESTAMP" json:"bbt_adm_created_at"`
	BibiteAdministratorUpdatedBy               string     `gorm:"column:bbt_adm_updated_by;size:125" json:"bbt_adm_updated_by"`
	BibiteAdministratorUpdatedAt               *time.Time `gorm:"column:bbt_adm_updated_at;default:CURRENT_TIMESTAMP" json:"bbt_adm_updated_at"`
	BibiteAdministratorDeletedBy               string     `gorm:"column:bbt_adm_deleted_by;size:125" json:"bbt_adm_deleted_by"`
	BibiteAdministratorDeletedAt               *time.Time `gorm:"column:bbt_adm_deleted_at;default:CURRENT_TIMESTAMP" json:"bbt_adm_deleted_at"`
}

type BibiteAdministratorData struct {
	BibiteAdministratorID                      uint64     `gorm:"column:bbt_adm_id;primary_key;" json:"bbt_adm_id"`
	BibiteEmployeeID                           uint64     `gorm:"column:bbt_emp_id" json:"bbt_emp_id"`
	BibiteEmployeeCode                         string     `gorm:"column:bbt_emp_code" json:"bbt_emp_code"`
	BibiteEmployeeName                         string     `gorm:"column:bbt_emp_name" json:"bbt_emp_name"`
	BibiteEmployeeGender                       string     `gorm:"column:bbt_emp_gender" json:"bbt_emp_gender"`
	BibiteEmployeeBirthPlace                   string     `gorm:"column:bbt_emp_birth_place" json:"bbt_emp_birth_place"`
	BibiteEmployeeBirthDate                    string     `gorm:"column:bbt_emp_birth_date" json:"bbt_emp_birth_date"`
	BibiteEmployeeAddress                      string     `gorm:"column:bbt_emp_address" json:"bbt_emp_address"`
	AddressID                                  uint64     `gorm:"column:address_id" json:"address_id"`
	CountryID                                  uint64     `gorm:"column:country_id" json:"country_id"`
	CountryName                                string     `gorm:"column:country_name" json:"country_name"`
	ProvinceID                                 uint64     `gorm:"column:province_id" json:"province_id"`
	ProvinceName                               string     `gorm:"column:province_name" json:"province_name"`
	RegencyID                                  uint64     `gorm:"column:regency_id" json:"regency_id"`
	RegencyName                                string     `gorm:"column:regency_name" json:"regency_name"`
	DistrictID                                 uint64     `gorm:"column:district_id" json:"district_id"`
	DistrictName                               string     `gorm:"column:district_name" json:"district_name"`
	VillageID                                  uint64     `gorm:"column:village_id" json:"village_id"`
	VillageName                                string     `gorm:"column:village_name" json:"village_name"`
	BibiteEmployeePhoto                        string     `gorm:"column:bbt_emp_photo" json:"bbt_emp_photo"`
	BibiteGroupID                              uint64     `gorm:"column:bbt_group_id" json:"bbt_group_id"`
	BibiteGroupCode                            string     `gorm:"column:bbt_group_code" json:"bbt_group_code"`
	BibiteGroupName                            string     `gorm:"column:bbt_group_name" json:"bbt_group_name"`
	BibiteGroupRole                            string     `gorm:"column:bbt_group_role" json:"bbt_group_role"`
	BibiteGroupType                            uint64     `gorm:"column:bbt_group_type" json:"bbt_group_type"`
	BibiteAdministratorEmail                   string     `gorm:"column:bbt_adm_email" json:"bbt_adm_email"`
	BibiteAdministratorUsername                string     `gorm:"column:bbt_adm_username" json:"bbt_adm_username"`
	BibiteAdministratorPhoneCode               string     `gorm:"column:bbt_adm_phone_code" json:"bbt_adm_phone_code"`
	BibiteAdministratorPhone                   string     `gorm:"column:bbt_adm_phone" json:"bbt_adm_phone"`
	BibiteAdministratorLastLogin               time.Time  `gorm:"column:bbt_adm_last_login" json:"bbt_adm_last_login"`
	BibiteAdministratorIPAddress               string     `gorm:"column:bbt_adm_ip_address" json:"bbt_adm_ip_address"`
	BibiteAdministratorStatus                  uint64     `gorm:"column:bbt_adm_status" json:"bbt_adm_status"`
	BibiteAdministratorIsNew                   bool       `gorm:"column:bbt_adm_is_new" json:"bbt_adm_is_new"`
	BibiteAdministratorIsRequestChangePassword bool       `gorm:"column:bbt_adm_is_req_chpw" json:"bbt_adm_is_req_chpw"`
	BibiteAdministratorIsLocked                bool       `gorm:"column:bbt_adm_is_locked" json:"bbt_adm_is_locked"`
	BibiteAdministratorFailAttempt             uint64     `gorm:"column:bbt_adm_fail_attempt" json:"bbt_adm_fail_attempt"`
	BibiteAdministratorLockedReason            string     `gorm:"column:bbt_adm_locked_reason" json:"bbt_adm_locked_reason"`
	BibiteAdministratorCreatedBy               string     `gorm:"column:bbt_adm_created_by" json:"bbt_adm_created_by"`
	BibiteAdministratorCreatedAt               time.Time  `gorm:"column:bbt_adm_created_at" json:"bbt_adm_created_at"`
	BibiteAdministratorUpdatedBy               string     `gorm:"column:bbt_adm_updated_by" json:"bbt_adm_updated_by"`
	BibiteAdministratorUpdatedAt               *time.Time `gorm:"column:bbt_adm_updated_at" json:"bbt_adm_updated_at"`
	BibiteAdministratorDeletedBy               string     `gorm:"column:bbt_adm_deleted_by" json:"bbt_adm_deleted_by"`
	BibiteAdministratorDeletedAt               *time.Time `gorm:"column:bbt_adm_deleted_at" json:"bbt_adm_deleted_at"`

	BibiteAdministratorCount int `gorm:"-" json:"bbt_user_count"`
}

type ResponseBibiteAdministrators struct {
	Status  int                        `json:"status"`
	Message string                     `json:"message"`
	Data    *[]BibiteAdministratorData `json:"data"`
}

type ResponseBibiteAdministrator struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    *BibiteAdministratorData `json:"data"`
}

type ResponseBibiteAdministratorCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (BibiteAdministrator) TableName() string {
	return "m_bibite_administrator"
}

func (BibiteAdministratorData) TableName() string {
	return "m_bibite_administrator"
}

func (p *BibiteAdministrator) Prepare() {
	p.BibiteEmployeeID = p.BibiteEmployeeID
	p.BibiteGroupID = p.BibiteGroupID
	p.BibiteAdministratorEmail = html.EscapeString(strings.TrimSpace(p.BibiteAdministratorEmail))
	p.BibiteAdministratorUsername = html.EscapeString(strings.TrimSpace(p.BibiteAdministratorUsername))
	p.BibiteAdministratorPassword = p.BibiteAdministratorPassword
	p.BibiteAdministratorPhoneCode = html.EscapeString(strings.TrimSpace(p.BibiteAdministratorPhoneCode))
	p.BibiteAdministratorPhone = html.EscapeString(strings.TrimSpace(p.BibiteAdministratorPhone))
	p.BibiteAdministratorOTP = html.EscapeString(strings.TrimSpace(p.BibiteAdministratorOTP))
	p.BibiteAdministratorOTPSecret = nil
	p.BibiteAdministratorSecret = p.BibiteAdministratorSecret
	p.BibiteAdministratorPIN = html.EscapeString(strings.TrimSpace(p.BibiteAdministratorPIN))
	p.BibiteAdministratorLastLogin = nil
	p.BibiteAdministratorStatus = p.BibiteAdministratorStatus
	p.BibiteAdministratorIsNew = p.BibiteAdministratorIsNew
	p.BibiteAdministratorIsRequestChangePassword = p.BibiteAdministratorIsRequestChangePassword
	p.BibiteAdministratorIsLocked = p.BibiteAdministratorIsLocked
	p.BibiteAdministratorFailAttempt = p.BibiteAdministratorFailAttempt
	p.BibiteAdministratorPIN = html.EscapeString(strings.TrimSpace(p.BibiteAdministratorPIN))
	p.BibiteAdministratorStatus = p.BibiteAdministratorStatus
	p.BibiteAdministratorCreatedBy = html.EscapeString(strings.TrimSpace(p.BibiteAdministratorCreatedBy))
	p.BibiteAdministratorCreatedAt = time.Now()
	p.BibiteAdministratorUpdatedBy = html.EscapeString(strings.TrimSpace(p.BibiteAdministratorUpdatedBy))
	p.BibiteAdministratorUpdatedAt = p.BibiteAdministratorUpdatedAt
	p.BibiteAdministratorDeletedBy = html.EscapeString(strings.TrimSpace(p.BibiteAdministratorDeletedBy))
	p.BibiteAdministratorDeletedAt = p.BibiteAdministratorDeletedAt
}

func (p *BibiteAdministrator) Validate(action string) error {
	switch strings.ToLower(action) {

	case "login":
		if p.BibiteAdministratorUsername == "" {
			return errors.New("Username Wajib Diisi")
		}
		if p.BibiteAdministratorPassword == "" {
			return errors.New("Password Wajib Diisi")
		}
		return nil

	default:
		if p.BibiteAdministratorPassword == "" {
			return errors.New("Required BibiteAdministrator Password")
		}
		if p.BibiteAdministratorEmail == "" {
			return errors.New("Required BibiteAdministrator Email")
		}
		return nil
	}
}

func (p *BibiteAdministrator) SaveBibiteAdministrator(db *gorm.DB) (*BibiteAdministrator, error) {
	var err error
	err = db.Debug().Model(&BibiteAdministrator{}).Create(&p).Error
	if err != nil {
		return &BibiteAdministrator{}, err
	}
	return p, nil
}

func (p *BibiteAdministrator) FindAllBibiteAdministrator(db *gorm.DB) (*[]BibiteAdministratorData, error) {
	var err error
	bibiteadministrator := []BibiteAdministratorData{}
	err = db.Debug().Model(&BibiteAdministratorData{}).Limit(100).
		Select(`m_bibite_administrator.bbt_adm_id,
				m_bibite_employee.bbt_emp_id,
				m_bibite_employee.bbt_emp_code,
				m_bibite_employee.bbt_emp_name,
				m_bibite_employee.bbt_emp_gender,
				m_bibite_employee.bbt_emp_birth_place,
				m_bibite_employee.bbt_emp_birth_date,
				m_bibite_employee.bbt_emp_address,
				m_bibite_employee.address_id,
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
				m_bibite_employee.bbt_emp_status,
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
				m_bibite_administrator.bbt_adm_created_at at time zone 'Asia/Jakarta' as bbt_adm_created_at,
				m_bibite_administrator.bbt_adm_updated_by,
				m_bibite_administrator.bbt_adm_updated_at at time zone 'Asia/Jakarta' as bbt_adm_updated_at,
				m_bibite_administrator.bbt_adm_deleted_by,
				m_bibite_administrator.bbt_adm_deleted_at at time zone 'Asia/Jakarta' as bbt_adm_deleted_at`).
		Joins("JOIN m_bibite_employee on m_bibite_administrator.bbt_emp_id=m_bibite_employee.bbt_emp_id").
		Joins("JOIN m_bibite_group on m_bibite_administrator.bbt_group_id=m_bibite_group.bbt_group_id").
		Joins("JOIN m_address on m_bibite_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Find(&bibiteadministrator).Error
	if err != nil {
		return &[]BibiteAdministratorData{}, err
	}
	return &bibiteadministrator, nil
}

func (p *BibiteAdministrator) FindAllBibiteAdministratorStatus(db *gorm.DB, status uint64) (*[]BibiteAdministratorData, error) {
	var err error
	bibiteadministrator := []BibiteAdministratorData{}
	err = db.Debug().Model(&BibiteAdministratorData{}).Limit(100).
		Select(`m_bibite_administrator.bbt_adm_id,
				m_bibite_employee.bbt_emp_id,
				m_bibite_employee.bbt_emp_code,
				m_bibite_employee.bbt_emp_name,
				m_bibite_employee.bbt_emp_gender,
				m_bibite_employee.bbt_emp_birth_place,
				m_bibite_employee.bbt_emp_birth_date,
				m_bibite_employee.bbt_emp_address,
				m_bibite_employee.address_id,
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
				m_bibite_employee.bbt_emp_status,
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
				m_bibite_administrator.bbt_adm_created_at at time zone 'Asia/Jakarta' as bbt_adm_created_at,
				m_bibite_administrator.bbt_adm_updated_by,
				m_bibite_administrator.bbt_adm_updated_at at time zone 'Asia/Jakarta' as bbt_adm_updated_at,
				m_bibite_administrator.bbt_adm_deleted_by,
				m_bibite_administrator.bbt_adm_deleted_at at time zone 'Asia/Jakarta' as bbt_adm_deleted_at`).
		Joins("JOIN m_bibite_employee on m_bibite_administrator.bbt_emp_id=m_bibite_employee.bbt_emp_id").
		Joins("JOIN m_bibite_group on m_bibite_administrator.bbt_group_id=m_bibite_group.bbt_group_id").
		Joins("JOIN m_address on m_bibite_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("bbt_adm_status = ?", status).
		Find(&bibiteadministrator).Error
	if err != nil {
		return &[]BibiteAdministratorData{}, err
	}
	return &bibiteadministrator, nil
}

func (p *BibiteAdministrator) FindBibiteAdministratorDataByID(db *gorm.DB, pid uint64) (*BibiteAdministrator, error) {
	var err error
	err = db.Debug().Model(&BibiteAdministrator{}).Where("bbt_adm_id = ?", pid).Take(&p).Error
	if err != nil {
		return &BibiteAdministrator{}, err
	}
	return p, nil
}

func (p *BibiteAdministrator) FindBibiteAdministratorByID(db *gorm.DB, pid uint64) (*BibiteAdministratorData, error) {
	var err error
	bibiteadministrator := BibiteAdministratorData{}
	err = db.Debug().Model(&BibiteAdministratorData{}).Limit(100).
		Select(`m_bibite_administrator.bbt_adm_id,
				m_bibite_employee.bbt_emp_id,
				m_bibite_employee.bbt_emp_code,
				m_bibite_employee.bbt_emp_name,
				m_bibite_employee.bbt_emp_gender,
				m_bibite_employee.bbt_emp_birth_place,
				m_bibite_employee.bbt_emp_birth_date,
				m_bibite_employee.bbt_emp_address,
				m_bibite_employee.address_id,
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
				m_bibite_employee.bbt_emp_status,
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
				m_bibite_administrator.bbt_adm_created_at at time zone 'Asia/Jakarta' as bbt_adm_created_at,
				m_bibite_administrator.bbt_adm_updated_by,
				m_bibite_administrator.bbt_adm_updated_at at time zone 'Asia/Jakarta' as bbt_adm_updated_at,
				m_bibite_administrator.bbt_adm_deleted_by,
				m_bibite_administrator.bbt_adm_deleted_at at time zone 'Asia/Jakarta' as bbt_adm_deleted_at`).
		Joins("JOIN m_bibite_employee on m_bibite_administrator.bbt_emp_id=m_bibite_employee.bbt_emp_id").
		Joins("JOIN m_bibite_group on m_bibite_administrator.bbt_group_id=m_bibite_group.bbt_group_id").
		Joins("JOIN m_address on m_bibite_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("bbt_adm_id = ?", pid).
		Find(&bibiteadministrator).Error
	if err != nil {
		return &BibiteAdministratorData{}, err
	}
	return &bibiteadministrator, nil
}

func (p *BibiteAdministrator) FindBibiteAdministratorStatusByID(db *gorm.DB, pid uint64, status uint64) (*BibiteAdministratorData, error) {
	var err error
	bibiteadministrator := BibiteAdministratorData{}
	err = db.Debug().Model(&BibiteAdministratorData{}).Limit(100).
		Select(`m_bibite_administrator.bbt_adm_id,
				m_bibite_employee.bbt_emp_id,
				m_bibite_employee.bbt_emp_code,
				m_bibite_employee.bbt_emp_name,
				m_bibite_employee.bbt_emp_gender,
				m_bibite_employee.bbt_emp_birth_place,
				m_bibite_employee.bbt_emp_birth_date,
				m_bibite_employee.bbt_emp_address,
				m_bibite_employee.address_id,
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
				m_bibite_employee.bbt_emp_status,
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
				m_bibite_administrator.bbt_adm_created_at at time zone 'Asia/Jakarta' as bbt_adm_created_at,
				m_bibite_administrator.bbt_adm_updated_by,
				m_bibite_administrator.bbt_adm_updated_at at time zone 'Asia/Jakarta' as bbt_adm_updated_at,
				m_bibite_administrator.bbt_adm_deleted_by,
				m_bibite_administrator.bbt_adm_deleted_at at time zone 'Asia/Jakarta' as bbt_adm_deleted_at`).
		Joins("JOIN m_bibite_employee on m_bibite_administrator.bbt_emp_id=m_bibite_employee.bbt_emp_id").
		Joins("JOIN m_bibite_group on m_bibite_administrator.bbt_group_id=m_bibite_group.bbt_group_id").
		Joins("JOIN m_address on m_bibite_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("bbt_adm_id = ? AND bbt_adm_status = ?", pid, status).
		Find(&bibiteadministrator).Error
	if err != nil {
		return &BibiteAdministratorData{}, err
	}
	return &bibiteadministrator, nil
}

func (p *BibiteAdministrator) UpdateBibiteAdministrator(db *gorm.DB, pid uint64) (*BibiteAdministrator, error) {
	var err error
	db = db.Debug().Model(&BibiteAdministrator{}).Where("bbt_adm_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_group_id":       p.BibiteGroupID,
			"bbt_adm_email":      p.BibiteAdministratorEmail,
			"bbt_adm_username":   p.BibiteAdministratorUsername,
			"bbt_adm_phone_code": p.BibiteAdministratorPhoneCode,
			"bbt_adm_phone":      p.BibiteAdministratorPhone,
			"bbt_adm_status":     p.BibiteAdministratorStatus,
			"bbt_adm_updated_by": p.BibiteAdministratorUpdatedBy,
			"bbt_adm_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&BibiteAdministrator{}).Where("bbt_adm_id = ?", pid).Error
	if err != nil {
		return &BibiteAdministrator{}, err
	}
	return p, nil
}

func (p *BibiteAdministrator) UpdateBibiteAdministratorStatus(db *gorm.DB, pid uint64) (*BibiteAdministrator, error) {
	var err error
	db = db.Debug().Model(&BibiteAdministrator{}).Where("bbt_adm_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_adm_status":     p.BibiteAdministratorStatus,
			"bbt_adm_updated_by": p.BibiteAdministratorUpdatedBy,
			"bbt_adm_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&BibiteAdministrator{}).Where("bbt_adm_id = ?", pid).Error
	if err != nil {
		return &BibiteAdministrator{}, err
	}
	return p, nil
}

func (p *BibiteAdministrator) UpdateBibiteAdministratorStatusOnly(db *gorm.DB, pid uint64) (*BibiteAdministrator, error) {
	var err error
	db = db.Debug().Model(&BibiteAdministrator{}).Where("bbt_adm_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_adm_status": p.BibiteAdministratorStatus,
		},
	)
	err = db.Debug().Model(&BibiteAdministrator{}).Where("bbt_adm_id = ?", pid).Error
	if err != nil {
		return &BibiteAdministrator{}, err
	}
	return p, nil
}

func (p *BibiteAdministrator) UpdateBibiteAdministratorStatusDelete(db *gorm.DB, pid uint64) (*BibiteAdministrator, error) {
	var err error
	db = db.Debug().Model(&BibiteAdministrator{}).Where("bbt_adm_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_adm_status":     3,
			"bbt_adm_deleted_by": p.BibiteAdministratorDeletedBy,
			"bbt_adm_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&BibiteAdministrator{}).Where("bbt_adm_id = ?", pid).Error
	if err != nil {
		return &BibiteAdministrator{}, err
	}
	return p, nil
}

func (p *BibiteAdministrator) DeleteBibiteAdministrator(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&BibiteAdministrator{}).Where("bbt_adm_id = ?", pid).Take(&BibiteAdministrator{}).Delete(&BibiteAdministrator{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("BibiteAdministrator not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//BIBITE LOGIN USER
//====================================================================================================================================

func BibiteAdministratorVerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (p *BibiteAdministrator) UpdateBibiteAdministratorLastLogin(db *gorm.DB, pid string) (*BibiteAdministrator, error) {

	var err error
	db = db.Debug().Model(&BibiteAdministrator{}).Where("bbt_adm_username = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_adm_last_login": time.Now(),
		},
	)
	err = db.Debug().Model(&BibiteAdministrator{}).Where("bbt_adm_username = ?", pid).Error
	if err != nil {
		return &BibiteAdministrator{}, err
	}
	return p, nil
}

//ADDITIONAL FROM CONTROLLER
//====================================================================================================================================

func (p *BibiteAdministrator) FindBibiteAdministratorStatusByUsername(db *gorm.DB, username string) (*BibiteAdministratorData, error) {
	var err error
	bibiteuser := BibiteAdministratorData{}
	err = db.Debug().Model(&BibiteAdministratorData{}).
		Select(`m_bibite_administrator.bbt_adm_id,
				m_bibite_administrator.bbt_emp_id,
				m_bibite_employee.bbt_emp_code,
				m_bibite_employee.bbt_emp_name,
				m_bibite_employee.bbt_emp_gender,
				m_bibite_employee.bbt_emp_birth_place,
				m_bibite_employee.bbt_emp_birth_date,
				m_bibite_employee.bbt_emp_address,
				m_bibite_employee.address_id,
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
				m_bibite_administrator.bbt_adm_ip_address,
				m_bibite_administrator.bbt_adm_last_login,
				m_bibite_administrator.bbt_adm_status,
				m_bibite_administrator.bbt_adm_is_new,
				m_bibite_administrator.bbt_adm_is_req_chpw,
				m_bibite_administrator.bbt_adm_is_locked,
				m_bibite_administrator.bbt_adm_fail_attempt,
				m_bibite_administrator.bbt_adm_locked_reason,
				m_bibite_administrator.bbt_adm_created_at at time zone 'Asia/Jakarta' as bbt_adm_created_at,
				m_bibite_administrator.bbt_adm_created_by,
				m_bibite_administrator.bbt_adm_updated_at at time zone 'Asia/Jakarta' as bbt_adm_updated_at,
				m_bibite_administrator.bbt_adm_updated_by,
				m_bibite_administrator.bbt_adm_deleted_at at time zone 'Asia/Jakarta' as bbt_adm_deleted_at,
				m_bibite_administrator.bbt_adm_deleted_by`).
		Joins("JOIN m_bibite_employee on m_bibite_administrator.bbt_emp_id=m_bibite_employee.bbt_emp_id").
		Joins("JOIN m_bibite_group on m_bibite_administrator.bbt_group_id=m_bibite_group.bbt_group_id").
		Where("bbt_adm_username = ? AND bbt_adm_status = ?", username, 1).
		Take(&bibiteuser).Error
	if err != nil {
		return &BibiteAdministratorData{}, err
	}
	return &bibiteuser, nil
}

func (p *BibiteAdministrator) CountBibiteAdministratorStatusByUsername(db *gorm.DB, username string) (*BibiteAdministratorData, error) {
	var err error
	bibiteuser := BibiteAdministratorData{}
	err = db.Debug().Model(&BibiteAdministratorData{}).
		Select(`m_bibite_administrator.bbt_adm_username,
				m_bibite_administrator.bbt_adm_status`).
		Where("bbt_adm_username = ? AND bbt_adm_status = ?", username, 1).
		Count(&bibiteuser.BibiteAdministratorCount).
		Take(&bibiteuser).Error
	if err != nil {
		return &BibiteAdministratorData{}, err
	}
	return &bibiteuser, nil
}

func (p *BibiteAdministrator) UpdateBibiteAdministratorProfile(db *gorm.DB, pid uint64) (*BibiteAdministrator, error) {
	var err error
	db = db.Debug().Model(&BibiteAdministrator{}).Where("bbt_adm_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_adm_email":      p.BibiteAdministratorEmail,
			"bbt_adm_username":   p.BibiteAdministratorUsername,
			"bbt_adm_phone_code": p.BibiteAdministratorPhoneCode,
			"bbt_adm_phone":      p.BibiteAdministratorPhone,
			"bbt_adm_updated_by": p.BibiteAdministratorUpdatedBy,
			"bbt_adm_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&BibiteAdministrator{}).Where("bbt_adm_id = ?", pid).Error
	if err != nil {
		return &BibiteAdministrator{}, err
	}
	return p, nil
}

func (p *BibiteAdministrator) BeforeUpdateBibitePassword() error {
	hashedPassword, err := HashBibiteAdministrator(p.BibiteAdministratorPassword)
	if err != nil {
		return err
	}
	p.BibiteAdministratorPassword = string(hashedPassword)
	return nil
}

func (p *BibiteAdministrator) UpdateBibiteAdministratorPassword(db *gorm.DB, pid uint64) (*BibiteAdministrator, error) {
	var err error
	err = p.BeforeUpdateBibitePassword()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&BibiteAdministrator{}).Where("bbt_adm_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_adm_password":   p.BibiteAdministratorPassword,
			"bbt_adm_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&BibiteAdministrator{}).Where("bbt_adm_id = ?", pid).Error
	if err != nil {
		return &BibiteAdministrator{}, err
	}
	return p, nil
}
