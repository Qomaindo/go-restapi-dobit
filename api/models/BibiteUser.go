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

type BibiteUser struct {
	BibiteUserID                      uint64     `gorm:"column:bbt_user_id;primary_key;" json:"bbt_user_id"`
	BibiteEmployeeID                  uint64     `gorm:"column:bbt_emp_id" json:"bbt_emp_id"`
	BibiteGroupID                     uint64     `gorm:"column:bbt_group_id" json:"bbt_group_id"`
	BibiteUserEmail                   string     `gorm:"column:bbt_user_email;size:255" json:"bbt_user_email"`
	BibiteUsername                    string     `gorm:"column:bbt_username;size:125" json:"bbt_username"`
	BibiteUserPassword                string     `gorm:"column:bbt_user_password;size:255" json:"bbt_user_password"`
	BibiteUserPhoneCode               string     `gorm:"column:bbt_user_phone_code;size:6" json:"bbt_user_phone_code"`
	BibiteUserPhone                   string     `gorm:"column:bbt_user_phone;size:20" json:"bbt_user_phone"`
	BibiteUserOTP                     string     `gorm:"column:bbt_user_otp;size:6" json:"bbt_user_otp"`
	BibiteUserOTPSecret               *time.Time `gorm:"column:bbt_user_otp_secret" json:"bbt_user_otp_secret"`
	BibiteUserSecret                  int64      `gorm:"column:bbt_user_secret" json:"bbt_user_secret"`
	BibiteUserPIN                     string     `gorm:"column:bbt_user_pin;size:6" json:"bbt_user_pin"`
	BibiteUserLastLogin               *time.Time `gorm:"column:bbt_user_last_login" json:"bbt_user_last_login"`
	BibiteUserIPAddress               string     `gorm:"column:bbt_user_ip_address;size:50" json:"bbt_user_ip_address"`
	BibiteUserStatus                  uint64     `gorm:"column:bbt_user_status" json:"bbt_user_status"`
	BibiteUserIsNew                   bool       `gorm:"column:bbt_user_is_new" json:"bbt_user_is_new"`
	BibiteUserIsRequestChangePassword bool       `gorm:"column:bbt_user_is_req_chpw" json:"bbt_user_is_req_chpw"`
	BibiteUserIsLocked                bool       `gorm:"column:bbt_user_is_locked" json:"bbt_user_is_locked"`
	BibiteUserFailAttempt             uint64     `gorm:"column:bbt_user_fail_attempt" json:"bbt_user_fail_attempt"`
	BibiteUserLockedReason            string     `gorm:"column:bbt_user_locked_reason;size:255" json:"bbt_user_locked_reason"`
	BibiteUserCreatedBy               string     `gorm:"column:bbt_user_created_by;size:125" json:"bbt_user_created_by"`
	BibiteUserCreatedAt               time.Time  `gorm:"column:bbt_user_created_at;default:CURRENT_TIMESTAMP" json:"bbt_user_created_at"`
	BibiteUserUpdatedBy               string     `gorm:"column:bbt_user_updated_by;size:125" json:"bbt_user_updated_by"`
	BibiteUserUpdatedAt               *time.Time `gorm:"column:bbt_user_updated_at;default:CURRENT_TIMESTAMP" json:"bbt_user_updated_at"`
	BibiteUserDeletedBy               string     `gorm:"column:bbt_user_deleted_by;size:125" json:"bbt_user_deleted_by"`
	BibiteUserDeletedAt               *time.Time `gorm:"column:bbt_user_deleted_at;default:CURRENT_TIMESTAMP" json:"bbt_user_deleted_at"`
}

type BibiteUserData struct {
	BibiteUserID                      uint64     `gorm:"column:bbt_user_id;primary_key;" json:"bbt_user_id"`
	BibiteEmployeeID                  uint64     `gorm:"column:bbt_emp_id" json:"bbt_emp_id"`
	BibiteEmployeeCode                string     `gorm:"column:bbt_emp_code" json:"bbt_emp_code"`
	BibiteEmployeeName                string     `gorm:"column:bbt_emp_name" json:"bbt_emp_name"`
	BibiteEmployeeGender              string     `gorm:"column:bbt_emp_gender" json:"bbt_emp_gender"`
	BibiteEmployeeBirthPlace          string     `gorm:"column:bbt_emp_birth_place" json:"bbt_emp_birth_place"`
	BibiteEmployeeBirthDate           string     `gorm:"column:bbt_emp_birth_date" json:"bbt_emp_birth_date"`
	BibiteEmployeeAddress             string     `gorm:"column:bbt_emp_address" json:"bbt_emp_address"`
	AddressID                         uint64     `gorm:"column:address_id" json:"address_id"`
	CountryID                         uint64     `gorm:"column:country_id" json:"country_id"`
	CountryName                       string     `gorm:"column:country_name" json:"country_name"`
	ProvinceID                        uint64     `gorm:"column:province_id" json:"province_id"`
	ProvinceName                      string     `gorm:"column:province_name" json:"province_name"`
	RegencyID                         uint64     `gorm:"column:regency_id" json:"regency_id"`
	RegencyName                       string     `gorm:"column:regency_name" json:"regency_name"`
	DistrictID                        uint64     `gorm:"column:district_id" json:"district_id"`
	DistrictName                      string     `gorm:"column:district_name" json:"district_name"`
	VillageID                         uint64     `gorm:"column:village_id" json:"village_id"`
	VillageName                       string     `gorm:"column:village_name" json:"village_name"`
	BibiteEmployeePhoto               string     `gorm:"column:bbt_emp_photo" json:"bbt_emp_photo"`
	BibiteGroupID                     uint64     `gorm:"column:bbt_group_id" json:"bbt_group_id"`
	BibiteGroupCode                   string     `gorm:"column:bbt_group_code" json:"bbt_group_code"`
	BibiteGroupName                   string     `gorm:"column:bbt_group_name" json:"bbt_group_name"`
	BibiteGroupRole                   string     `gorm:"column:bbt_group_role" json:"bbt_group_role"`
	BibiteGroupType                   uint64     `gorm:"column:bbt_group_type" json:"bbt_group_type"`
	BibiteUserEmail                   string     `gorm:"column:bbt_user_email" json:"bbt_user_email"`
	BibiteUsername                    string     `gorm:"column:bbt_username" json:"bbt_username"`
	BibiteUserPhoneCode               string     `gorm:"column:bbt_user_phone_code" json:"bbt_user_phone_code"`
	BibiteUserPhone                   string     `gorm:"column:bbt_user_phone" json:"bbt_user_phone"`
	BibiteUserLastLogin               time.Time  `gorm:"column:bbt_user_last_login" json:"bbt_user_last_login"`
	BibiteUserIPAddress               string     `gorm:"column:bbt_user_ip_address" json:"bbt_user_ip_address"`
	BibiteUserStatus                  uint64     `gorm:"column:bbt_user_status" json:"bbt_user_status"`
	BibiteUserIsNew                   bool       `gorm:"column:bbt_user_is_new" json:"bbt_user_is_new"`
	BibiteUserIsRequestChangePassword bool       `gorm:"column:bbt_user_is_req_chpw" json:"bbt_user_is_req_chpw"`
	BibiteUserIsLocked                bool       `gorm:"column:bbt_user_is_locked" json:"bbt_user_is_locked"`
	BibiteUserFailAttempt             uint64     `gorm:"column:bbt_user_fail_attempt" json:"bbt_user_fail_attempt"`
	BibiteUserLockedReason            string     `gorm:"column:bbt_user_locked_reason" json:"bbt_user_locked_reason"`
	BibiteUserCreatedBy               string     `gorm:"column:bbt_user_created_by" json:"bbt_user_created_by"`
	BibiteUserCreatedAt               time.Time  `gorm:"column:bbt_user_created_at" json:"bbt_user_created_at"`
	BibiteUserUpdatedBy               string     `gorm:"column:bbt_user_updated_by" json:"bbt_user_updated_by"`
	BibiteUserUpdatedAt               *time.Time `gorm:"column:bbt_user_updated_at" json:"bbt_user_updated_at"`
	BibiteUserDeletedBy               string     `gorm:"column:bbt_user_deleted_by" json:"bbt_user_deleted_by"`
	BibiteUserDeletedAt               *time.Time `gorm:"column:bbt_user_deleted_at" json:"bbt_user_deleted_at"`

	BibiteUserCount int `gorm:"-" json:"bbt_user_count"`
}

type ResponseBibiteUsers struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    *[]BibiteUserData `json:"data"`
}

type ResponseBibiteUser struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Data    *BibiteUserData `json:"data"`
}

type ResponseBibiteUserCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (BibiteUser) TableName() string {
	return "m_bibite_user"
}

func (BibiteUserData) TableName() string {
	return "m_bibite_user"
}

func (p *BibiteUser) Prepare() {
	p.BibiteEmployeeID = p.BibiteEmployeeID
	p.BibiteGroupID = p.BibiteGroupID
	p.BibiteUserEmail = html.EscapeString(strings.TrimSpace(p.BibiteUserEmail))
	p.BibiteUsername = html.EscapeString(strings.TrimSpace(p.BibiteUsername))
	p.BibiteUserPassword = p.BibiteUserPassword
	p.BibiteUserPhoneCode = html.EscapeString(strings.TrimSpace(p.BibiteUserPhoneCode))
	p.BibiteUserPhone = html.EscapeString(strings.TrimSpace(p.BibiteUserPhone))
	p.BibiteUserOTP = html.EscapeString(strings.TrimSpace(p.BibiteUserOTP))
	p.BibiteUserOTPSecret = nil
	p.BibiteUserSecret = p.BibiteUserSecret
	p.BibiteUserPIN = html.EscapeString(strings.TrimSpace(p.BibiteUserPIN))
	p.BibiteUserLastLogin = nil
	p.BibiteUserStatus = p.BibiteUserStatus
	p.BibiteUserIsNew = p.BibiteUserIsNew
	p.BibiteUserIsRequestChangePassword = p.BibiteUserIsRequestChangePassword
	p.BibiteUserIsLocked = p.BibiteUserIsLocked
	p.BibiteUserFailAttempt = p.BibiteUserFailAttempt
	p.BibiteUserPIN = html.EscapeString(strings.TrimSpace(p.BibiteUserPIN))
	p.BibiteUserCreatedBy = p.BibiteUserCreatedBy
	p.BibiteUserCreatedAt = time.Now()
	p.BibiteUserUpdatedBy = p.BibiteUserUpdatedBy
	p.BibiteUserUpdatedAt = p.BibiteUserUpdatedAt
	p.BibiteUserDeletedBy = p.BibiteUserDeletedBy
	p.BibiteUserDeletedAt = p.BibiteUserDeletedAt
}

func (p *BibiteUser) Validate(action string) error {
	switch strings.ToLower(action) {

	case "login":
		if p.BibiteUsername == "" {
			return errors.New("Username Wajib Diisi")
		}
		if p.BibiteUserPassword == "" {
			return errors.New("Password Wajib Diisi")
		}
		return nil

	default:
		if p.BibiteUserPassword == "" {
			return errors.New("Required BibiteUser Password")
		}
		if p.BibiteUserEmail == "" {
			return errors.New("Required BibiteUser Email")
		}
		return nil
	}
}

func (p *BibiteUser) SaveBibiteUser(db *gorm.DB) (*BibiteUser, error) {
	var err error
	err = db.Debug().Model(&BibiteUser{}).Create(&p).Error
	if err != nil {
		return &BibiteUser{}, err
	}
	return p, nil
}

func (p *BibiteUser) FindAllBibiteUser(db *gorm.DB) (*[]BibiteUserData, error) {
	var err error
	bibiteuser := []BibiteUserData{}
	err = db.Debug().Model(&BibiteUserData{}).Limit(100).
		Select(`m_bibite_user.bbt_user_id,
				m_bibite_user.bbt_emp_id,
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
				m_bibite_user.bbt_user_created_at at time zone 'Asia/Jakarta' as bbt_user_created_at,
				m_bibite_user.bbt_user_updated_by,
				m_bibite_user.bbt_user_updated_at at time zone 'Asia/Jakarta' as bbt_user_updated_at,
				m_bibite_user.bbt_user_deleted_by,
				m_bibite_user.bbt_user_deleted_at at time zone 'Asia/Jakarta' as bbt_user_deleted_at`).
		Joins("JOIN m_bibite_employee on m_bibite_user.bbt_emp_id=m_bibite_employee.bbt_emp_id").
		Joins("JOIN m_address on m_bibite_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Joins("JOIN m_bibite_group on m_bibite_user.bbt_group_id=m_bibite_group.bbt_group_id").
		Find(&bibiteuser).Error
	if err != nil {
		return &[]BibiteUserData{}, err
	}
	return &bibiteuser, nil
}

func (p *BibiteUser) FindAllBibiteUserStatus(db *gorm.DB, status uint64) (*[]BibiteUserData, error) {
	var err error
	bibiteuser := []BibiteUserData{}
	err = db.Debug().Model(&BibiteUserData{}).Limit(100).
		Select(`m_bibite_user.bbt_user_id,
				m_bibite_user.bbt_emp_id,
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
				m_bibite_user.bbt_user_created_at at time zone 'Asia/Jakarta' as bbt_user_created_at,
				m_bibite_user.bbt_user_updated_by,
				m_bibite_user.bbt_user_updated_at at time zone 'Asia/Jakarta' as bbt_user_updated_at,
				m_bibite_user.bbt_user_deleted_by,
				m_bibite_user.bbt_user_deleted_at at time zone 'Asia/Jakarta' as bbt_user_deleted_at`).
		Joins("JOIN m_bibite_employee on m_bibite_user.bbt_emp_id=m_bibite_employee.bbt_emp_id").
		Joins("JOIN m_address on m_bibite_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Joins("JOIN m_bibite_group on m_bibite_user.bbt_group_id=m_bibite_group.bbt_group_id").
		Where("bbt_user_status = ?", status).
		Find(&bibiteuser).Error
	if err != nil {
		return &[]BibiteUserData{}, err
	}
	return &bibiteuser, nil
}

func (p *BibiteUser) FindBibiteUserByID(db *gorm.DB, pid uint64) (*BibiteUserData, error) {
	var err error
	bibiteuser := BibiteUserData{}
	err = db.Debug().Model(&BibiteUserData{}).
		Select(`m_bibite_user.bbt_user_id,
				m_bibite_user.bbt_emp_id,
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
				m_bibite_user.bbt_user_created_at at time zone 'Asia/Jakarta' as bbt_user_created_at,
				m_bibite_user.bbt_user_updated_by,
				m_bibite_user.bbt_user_updated_at at time zone 'Asia/Jakarta' as bbt_user_updated_at,
				m_bibite_user.bbt_user_deleted_by,
				m_bibite_user.bbt_user_deleted_at at time zone 'Asia/Jakarta' as bbt_user_deleted_at`).
		Joins("JOIN m_bibite_employee on m_bibite_user.bbt_emp_id=m_bibite_employee.bbt_emp_id").
		Joins("JOIN m_address on m_bibite_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Joins("JOIN m_bibite_group on m_bibite_user.bbt_group_id=m_bibite_group.bbt_group_id").
		Where("bbt_user_id = ?", pid).
		Take(&bibiteuser).Error
	if err != nil {
		return &BibiteUserData{}, err
	}
	return &bibiteuser, nil
}

func (p *BibiteUser) FindBibiteUserStatusByID(db *gorm.DB, pid uint64, status uint64) (*BibiteUserData, error) {
	var err error
	bibiteuser := BibiteUserData{}
	err = db.Debug().Model(&BibiteUserData{}).
		Select(`m_bibite_user.bbt_user_id,
				m_bibite_user.bbt_emp_id,
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
				m_bibite_user.bbt_user_created_at at time zone 'Asia/Jakarta' as bbt_user_created_at,
				m_bibite_user.bbt_user_updated_by,
				m_bibite_user.bbt_user_updated_at at time zone 'Asia/Jakarta' as bbt_user_updated_at,
				m_bibite_user.bbt_user_deleted_by,
				m_bibite_user.bbt_user_deleted_at at time zone 'Asia/Jakarta' as bbt_user_deleted_at`).
		Joins("JOIN m_bibite_employee on m_bibite_user.bbt_emp_id=m_bibite_employee.bbt_emp_id").
		Joins("JOIN m_address on m_bibite_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Joins("JOIN m_bibite_group on m_bibite_user.bbt_group_id=m_bibite_group.bbt_group_id").
		Where("bbt_user_id = ? AND bbt_user_status = ?", pid, status).
		Take(&bibiteuser).Error
	if err != nil {
		return &BibiteUserData{}, err
	}
	return &bibiteuser, nil
}

func (p *BibiteUser) UpdateBibiteUser(db *gorm.DB, pid uint64) (*BibiteUser, error) {

	var err error
	db = db.Debug().Model(&BibiteUser{}).Where("bbt_user_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_group_id":        p.BibiteGroupID,
			"bbt_user_email":      p.BibiteUserEmail,
			"bbt_username":        p.BibiteUsername,
			"bbt_user_status":     p.BibiteUserStatus,
			"bbt_user_updated_by": p.BibiteUserUpdatedBy,
			"bbt_user_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&BibiteUser{}).Where("bbt_user_id = ?", pid).Error
	if err != nil {
		return &BibiteUser{}, err
	}
	return p, nil
}

func (p *BibiteUser) UpdateBibiteUserStatus(db *gorm.DB, pid uint64) (*BibiteUser, error) {

	var err error
	db = db.Debug().Model(&BibiteUser{}).Where("bbt_user_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_user_status":     p.BibiteUserStatus,
			"bbt_user_updated_by": p.BibiteUserUpdatedBy,
			"bbt_user_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&BibiteUser{}).Where("bbt_user_id = ?", pid).Error
	if err != nil {
		return &BibiteUser{}, err
	}
	return p, nil
}

func (p *BibiteUser) UpdateBibiteUserStatusOnly(db *gorm.DB, pid uint64) (*BibiteUser, error) {

	var err error
	db = db.Debug().Model(&BibiteUser{}).Where("bbt_user_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_user_status": p.BibiteUserStatus,
		},
	)
	err = db.Debug().Model(&BibiteUser{}).Where("bbt_user_id = ?", pid).Error
	if err != nil {
		return &BibiteUser{}, err
	}
	return p, nil
}

func (p *BibiteUser) UpdateBibiteUserStatusDelete(db *gorm.DB, pid uint64) (*BibiteUser, error) {
	var err error
	db = db.Debug().Model(&BibiteUser{}).Where("bbt_user_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_user_status":     3,
			"bbt_user_deleted_by": p.BibiteUserDeletedBy,
			"bbt_user_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&BibiteUser{}).Where("bbt_user_id = ?", pid).Error
	if err != nil {
		return &BibiteUser{}, err
	}
	return p, nil
}

func (p *BibiteUser) DeleteBibiteUser(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&BibiteUser{}).Where("bbt_user_id = ?", pid).Take(&BibiteUser{}).Delete(&BibiteUser{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("BibiteUser not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//BIBITE LOGIN USER
//====================================================================================================================================

func BibiteVerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (p *BibiteUser) UpdateBibiteUserLastLogin(db *gorm.DB, pid string) (*BibiteUser, error) {

	var err error
	db = db.Debug().Model(&BibiteUser{}).Where("bbt_username = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_user_last_login": time.Now(),
		},
	)
	err = db.Debug().Model(&BibiteUser{}).Where("bbt_username = ?", pid).Error
	if err != nil {
		return &BibiteUser{}, err
	}
	return p, nil
}

//ADDITIONAL FROM CONTROLLER
//====================================================================================================================================

func (p *BibiteUser) FindBibiteUserStatusByUsername(db *gorm.DB, username string) (*BibiteUserData, error) {
	var err error
	bibiteuser := BibiteUserData{}
	err = db.Debug().Model(&BibiteUserData{}).
		Select(`m_bibite_user.bbt_user_id,
				m_bibite_user.bbt_emp_id,
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
				m_bibite_user.bbt_user_created_at at time zone 'Asia/Jakarta' as bbt_user_created_at,
				m_bibite_user.bbt_user_updated_by,
				m_bibite_user.bbt_user_updated_at at time zone 'Asia/Jakarta' as bbt_user_updated_at,
				m_bibite_user.bbt_user_deleted_by,
				m_bibite_user.bbt_user_deleted_at at time zone 'Asia/Jakarta' as bbt_user_deleted_at`).
		Joins("JOIN m_bibite_employee on m_bibite_user.bbt_emp_id=m_bibite_employee.bbt_emp_id").
		Joins("JOIN m_address on m_bibite_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Joins("JOIN m_bibite_group on m_bibite_user.bbt_group_id=m_bibite_group.bbt_group_id").
		Where("bbt_username = ? AND bbt_user_status = ?", username, 1).
		Take(&bibiteuser).Error
	if err != nil {
		return &BibiteUserData{}, err
	}
	return &bibiteuser, nil
}

func (p *BibiteUser) CountBibiteUserStatusByUsername(db *gorm.DB, username string) (*BibiteUserData, error) {
	var err error
	bibiteuser := BibiteUserData{}
	err = db.Debug().Model(&BibiteUserData{}).
		Select(`m_bibite_user.bbt_username,
				m_bibite_user.bbt_user_status`).
		Where("bbt_username = ? AND bbt_user_status = ?", username, 1).
		Count(&bibiteuser.BibiteUserCount).
		Take(&bibiteuser).Error
	if err != nil {
		return &BibiteUserData{}, err
	}
	return &bibiteuser, nil
}

func (p *BibiteUser) UpdateBibiteUserProfile(db *gorm.DB, pid uint64) (*BibiteUser, error) {

	var err error
	db = db.Debug().Model(&BibiteUser{}).Where("bbt_user_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_user_email":      p.BibiteUserEmail,
			"bbt_username":        p.BibiteUsername,
			"bbt_user_phone_code": p.BibiteUserPhoneCode,
			"bbt_user_phone":      p.BibiteUserPhone,
			"bbt_user_status":     p.BibiteUserStatus,
			"bbt_user_updated_by": p.BibiteUserUpdatedBy,
			"bbt_user_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&BibiteUser{}).Where("bbt_user_id = ?", pid).Error
	if err != nil {
		return &BibiteUser{}, err
	}
	return p, nil
}

func (p *BibiteUser) BeforeUpdateBibitePassword() error {
	hashedPassword, err := HashBibiteUser(p.BibiteUserPassword)
	if err != nil {
		return err
	}
	p.BibiteUserPassword = string(hashedPassword)
	return nil
}

func (p *BibiteUser) UpdateBibiteUserPassword(db *gorm.DB, pid uint64) (*BibiteUser, error) {
	var err error
	err = p.BeforeUpdateBibitePassword()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&BibiteUser{}).Where("bbt_user_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_user_password":   p.BibiteUserPassword,
			"bbt_user_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&BibiteUser{}).Where("bbt_user_id = ?", pid).Error
	if err != nil {
		return &BibiteUser{}, err
	}
	return p, nil
}
