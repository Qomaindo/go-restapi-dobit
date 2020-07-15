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

type MitraAdministrator struct {
	MitraAdministratorID                      *uint64    `gorm:"column:mitra_adm_id;primary_key;" json:"mitra_adm_id"`
	MitraEmployeeID                           uint64     `gorm:"column:mitra_emp_id" json:"mitra_emp_id"`
	MitraGroupAccessID                        uint64     `gorm:"column:mitra_group_accs_id" json:"mitra_group_accs_id"`
	MitraAdministratorEmail                   string     `gorm:"column:mitra_adm_email" json:"mitra_adm_email"`
	MitraAdministratorUsername                string     `gorm:"column:mitra_adm_username" json:"mitra_adm_username"`
	MitraAdministratorPassword                string     `gorm:"column:mitra_adm_password" json:"mitra_adm_password"`
	MitraAdministratorPhoneCode               string     `gorm:"column:mitra_adm_phone_code" json:"mitra_adm_phone_code"`
	MitraAdministratorPhone                   string     `gorm:"column:mitra_adm_phone" json:"mitra_adm_phone"`
	MitraAdministratorOTP                     string     `gorm:"column:mitra_adm_otp" json:"mitra_adm_otp"`
	MitraAdministratorOTPSecret               *time.Time `gorm:"column:mitra_adm_otp_secret" json:"mitra_adm_otp_secret"`
	MitraAdministratorSecret                  uint64     `gorm:"column:mitra_adm_secret" json:"mitra_adm_secret"`
	MitraAdministratorPIN                     string     `gorm:"column:mitra_adm_pin" json:"mitra_adm_pin"`
	MitraAdministratorLastLogin               *time.Time `gorm:"column:mitra_adm_last_login" json:"mitra_adm_last_login"`
	MitraAdministratorIPAddress               string     `gorm:"column:mitra_adm_ip_address" json:"mitra_adm_ip_address"`
	MitraAdministratorStatus                  uint64     `gorm:"column:mitra_adm_status" json:"mitra_adm_status"`
	MitraAdministratorIsNew                   bool       `gorm:"column:mitra_adm_is_new" json:"mitra_adm_is_new"`
	MitraAdministratorIsRequestChangePassword bool       `gorm:"column:mitra_adm_is_req_chpw" json:"mitra_adm_is_req_chpw"`
	MitraAdministratorIsLocked                bool       `gorm:"column:mitra_adm_is_locked" json:"mitra_adm_is_locked"`
	MitraAdministratorFailAttempt             uint64     `gorm:"column:mitra_adm_fail_attempt" json:"mitra_adm_fail_attempt"`
	MitraAdministratorLockedReason            string     `gorm:"column:mitra_adm_locked_reason" json:"mitra_adm_locked_reason"`
	MitraAdministratorCreatedBy               string     `gorm:"column:mitra_adm_created_by;size:125" json:"mitra_adm_created_by"`
	MitraAdministratorCreatedAt               time.Time  `gorm:"column:mitra_adm_created_at;default:CURRENT_TIMESTAMP" json:"mitra_adm_created_at"`
	MitraAdministratorUpdatedBy               string     `gorm:"column:mitra_adm_updated_by;size:125" json:"mitra_adm_updated_by"`
	MitraAdministratorUpdatedAt               *time.Time `gorm:"column:mitra_adm_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_adm_created_at"`
	MitraAdministratorDeletedBy               string     `gorm:"column:mitra_adm_deleted_by;size:125" json:"mitra_adm_deleted_by"`
	MitraAdministratorDeletedAt               *time.Time `gorm:"column:mitra_adm_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_adm_deleted_at"`
}

type MitraAdministratorData struct {
	MitraAdministratorID                      uint64     `gorm:"column:mitra_adm_id;primary_key;" json:"mitra_adm_id"`
	MitraEmployeeID                           uint64     `gorm:"column:mitra_emp_id" json:"mitra_emp_id"`
	MitraEmployeeCode                         string     `gorm:"column:mitra_emp_code" json:"mitra_emp_code"`
	MitraEmployeeName                         string     `gorm:"column:mitra_emp_name" json:"mitra_emp_name"`
	MitraEmployeeGender                       string     `gorm:"column:mitra_emp_gender" json:"mitra_emp_gender"`
	MitraEmployeeBirthPlace                   string     `gorm:"column:mitra_emp_birth_place" json:"mitra_emp_birth_place"`
	MitraEmployeeBirthDate                    string     `gorm:"column:mitra_emp_birth_date" json:"mitra_emp_birth_date"`
	MitraEmployeeAddress                      string     `gorm:"column:mitra_emp_address" json:"mitra_emp_address"`
	AddressID                                 uint64     `gorm:"column:address_id" json:"address_id"`
	CountryID                                 uint64     `gorm:"column:country_id" json:"country_id"`
	CountryName                               string     `gorm:"column:country_name" json:"country_name"`
	ProvinceID                                uint64     `gorm:"column:province_id" json:"province_id"`
	ProvinceName                              string     `gorm:"column:province_name" json:"province_name"`
	RegencyID                                 uint64     `gorm:"column:regency_id" json:"regency_id"`
	RegencyName                               string     `gorm:"column:regency_name" json:"regency_name"`
	DistrictID                                uint64     `gorm:"column:district_id" json:"district_id"`
	DistrictName                              string     `gorm:"column:district_name" json:"district_name"`
	VillageID                                 uint64     `gorm:"column:village_id" json:"village_id"`
	VillageName                               string     `gorm:"column:village_name" json:"village_name"`
	MitraID                                   uint64     `gorm:"column:mitra_id" json:"mitra_id"`
	MitraName                                 string     `gorm:"column:mitra_name" json:"mitra_name"`
	MitraBranchID                             uint64     `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	MitraBranchName                           string     `gorm:"column:mitra_branch_name" json:"mitra_branch_name"`
	MitraEmployeePhoto                        string     `gorm:"column:mitra_emp_photo" json:"mitra_emp_photo"`
	MitraGroupAccessID                        uint64     `gorm:"column:mitra_group_accs_id" json:"mitra_group_accs_id"`
	MitraGroupAccessCode                      string     `gorm:"column:mitra_group_accs_code" json:"mitra_group_accs_code"`
	MitraGroupAccessName                      string     `gorm:"column:mitra_group_accs_name" json:"mitra_group_accs_name"`
	MitraGroupAccessRole                      string     `gorm:"column:mitra_group_accs_role" json:"mitra_group_accs_role"`
	MitraGroupAccessType                      uint64     `gorm:"column:mitra_group_accs_type" json:"mitra_group_accs_type"`
	MitraAdministratorEmail                   string     `gorm:"column:mitra_adm_email" json:"mitra_adm_email"`
	MitraAdministratorUsername                string     `gorm:"column:mitra_adm_username" json:"mitra_adm_username"`
	MitraAdministratorPhoneCode               string     `gorm:"column:mitra_adm_phone_code" json:"mitra_adm_phone_code"`
	MitraAdministratorPhone                   string     `gorm:"column:mitra_adm_phone" json:"mitra_adm_phone"`
	MitraAdministratorLastLogin               *time.Time `gorm:"column:mitra_adm_last_login" json:"mitra_adm_last_login"`
	MitraAdministratorIPAddress               string     `gorm:"column:mitra_adm_ip_address" json:"mitra_adm_ip_address"`
	MitraAdministratorStatus                  uint64     `gorm:"column:mitra_adm_status" json:"mitra_adm_status"`
	MitraAdministratorIsNew                   bool       `gorm:"column:mitra_adm_is_new" json:"mitra_adm_is_new"`
	MitraAdministratorIsRequestChangePassword bool       `gorm:"column:mitra_adm_is_req_chpw" json:"mitra_adm_is_req_chpw"`
	MitraAdministratorIsLocked                bool       `gorm:"column:mitra_adm_is_locked" json:"mitra_adm_is_locked"`
	MitraAdministratorFailAttempt             uint64     `gorm:"column:mitra_adm_fail_attempt" json:"mitra_adm_fail_attempt"`
	MitraAdministratorLockedReason            string     `gorm:"column:mitra_adm_locked_reason" json:"mitra_adm_locked_reason"`
	MitraAdministratorCreatedBy               string     `gorm:"column:mitra_adm_created_by;size:125" json:"mitra_adm_created_by"`
	MitraAdministratorCreatedAt               time.Time  `gorm:"column:mitra_adm_created_at;default:CURRENT_TIMESTAMP" json:"mitra_adm_created_at"`
	MitraAdministratorUpdatedBy               string     `gorm:"column:mitra_adm_updated_by;size:125" json:"mitra_adm_updated_by"`
	MitraAdministratorUpdatedAt               *time.Time `gorm:"column:mitra_adm_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_adm_created_at"`
	MitraAdministratorDeletedBy               string     `gorm:"column:mitra_adm_deleted_by;size:125" json:"mitra_adm_deleted_by"`
	MitraAdministratorDeletedAt               *time.Time `gorm:"column:mitra_adm_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_adm_deleted_at"`

	MitraAdministratorCount int `gorm:"-" json:"mitra_adm_count"`
}

type ResponseMitraAdministrators struct {
	Status  int                       `json:"status"`
	Message string                    `json:"message"`
	Data    *[]MitraAdministratorData `json:"data"`
}

type ResponseMitraAdministrator struct {
	Status  int                     `json:"status"`
	Message string                  `json:"message"`
	Data    *MitraAdministratorData `json:"data"`
}

type ResponseMitraAdministratorCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraAdministrator) TableName() string {
	return "m_mitra_administrator"
}

func (MitraAdministratorData) TableName() string {
	return "m_mitra_administrator"
}

func (p *MitraAdministrator) Prepare() {
	p.MitraAdministratorID = nil
	p.MitraEmployeeID = p.MitraEmployeeID
	p.MitraGroupAccessID = p.MitraGroupAccessID
	p.MitraAdministratorEmail = html.EscapeString(strings.TrimSpace(p.MitraAdministratorEmail))
	p.MitraAdministratorUsername = html.EscapeString(strings.TrimSpace(p.MitraAdministratorUsername))
	p.MitraAdministratorPassword = p.MitraAdministratorPassword
	p.MitraAdministratorPhoneCode = html.EscapeString(strings.TrimSpace(p.MitraAdministratorPhoneCode))
	p.MitraAdministratorPhone = html.EscapeString(strings.TrimSpace(p.MitraAdministratorPhone))
	p.MitraAdministratorOTP = html.EscapeString(strings.TrimSpace(p.MitraAdministratorOTP))
	p.MitraAdministratorOTPSecret = nil
	p.MitraAdministratorSecret = p.MitraAdministratorSecret
	p.MitraAdministratorPIN = html.EscapeString(strings.TrimSpace(p.MitraAdministratorPIN))
	p.MitraAdministratorLastLogin = nil
	p.MitraAdministratorStatus = p.MitraAdministratorStatus
	p.MitraAdministratorIsNew = p.MitraAdministratorIsNew
	p.MitraAdministratorIsRequestChangePassword = p.MitraAdministratorIsRequestChangePassword
	p.MitraAdministratorIsLocked = p.MitraAdministratorIsLocked
	p.MitraAdministratorFailAttempt = p.MitraAdministratorFailAttempt
	p.MitraAdministratorPIN = html.EscapeString(strings.TrimSpace(p.MitraAdministratorPIN))
	p.MitraAdministratorStatus = p.MitraAdministratorStatus
	p.MitraAdministratorCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraAdministratorCreatedBy))
	p.MitraAdministratorCreatedAt = time.Now()
	p.MitraAdministratorUpdatedBy = html.EscapeString(strings.TrimSpace(p.MitraAdministratorUpdatedBy))
	p.MitraAdministratorUpdatedAt = p.MitraAdministratorUpdatedAt
	p.MitraAdministratorDeletedBy = html.EscapeString(strings.TrimSpace(p.MitraAdministratorDeletedBy))
	p.MitraAdministratorDeletedAt = p.MitraAdministratorDeletedAt
}

func (p *MitraAdministrator) Validate(action string) error {
	switch strings.ToLower(action) {

	case "login":
		if p.MitraAdministratorUsername == "" {
			return errors.New("Username Wajib Diisi")
		}
		if p.MitraAdministratorPassword == "" {
			return errors.New("Password Wajib Diisi")
		}
		return nil

	default:
		if p.MitraAdministratorPassword == "" {
			return errors.New("Required MitraAdministrator Password")
		}
		if p.MitraAdministratorEmail == "" {
			return errors.New("Required MitraAdministrator Email")
		}
		return nil
	}
}

func (p *MitraAdministrator) SaveMitraAdministrator(db *gorm.DB) (*MitraAdministrator, error) {
	var err error
	err = db.Debug().Model(&MitraAdministrator{}).Create(&p).Error
	if err != nil {
		return &MitraAdministrator{}, err
	}
	return p, nil
}

func (p *MitraAdministrator) FindAllMitraAdministrator(db *gorm.DB) (*[]MitraAdministratorData, error) {
	var err error
	mitraadministrator := []MitraAdministratorData{}
	err = db.Debug().Model(&MitraAdministratorData{}).Limit(100).
		Select(`m_mitra_administrator.mitra_adm_id,
				m_mitra_employee.mitra_emp_id,
				m_mitra_employee.mitra_emp_code,
				m_mitra_employee.mitra_emp_name,
				m_mitra_employee.mitra_emp_gender,
				m_mitra_employee.mitra_emp_birth_place,
				m_mitra_employee.mitra_emp_birth_date,
				m_mitra_employee.mitra_emp_address,
				m_mitra_employee.address_id,
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
				m_mitra_employee.mitra_emp_status,
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
				m_mitra_administrator.mitra_adm_created_at,
				m_mitra_administrator.mitra_adm_updated_by,
				m_mitra_administrator.mitra_adm_updated_at,
				m_mitra_administrator.mitra_adm_deleted_by,
				m_mitra_administrator.mitra_adm_deleted_at`).
		Joins("JOIN m_mitra_employee on m_mitra_administrator.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra_group_access on m_mitra_administrator.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Joins("JOIN m_address on m_mitra_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Find(&mitraadministrator).Error
	if err != nil {
		return &[]MitraAdministratorData{}, err
	}
	return &mitraadministrator, nil
}

func (p *MitraAdministrator) FindAllMitraAdministratorStatus(db *gorm.DB, status uint64) (*[]MitraAdministratorData, error) {
	var err error
	mitraadministrator := []MitraAdministratorData{}
	err = db.Debug().Model(&MitraAdministratorData{}).Limit(100).
		Select(`m_mitra_administrator.mitra_adm_id,
				m_mitra_employee.mitra_emp_id,
				m_mitra_employee.mitra_emp_code,
				m_mitra_employee.mitra_emp_name,
				m_mitra_employee.mitra_emp_gender,
				m_mitra_employee.mitra_emp_birth_place,
				m_mitra_employee.mitra_emp_birth_date,
				m_mitra_employee.mitra_emp_address,
				m_mitra_employee.address_id,
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
				m_mitra_employee.mitra_emp_status,
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
				m_mitra_administrator.mitra_adm_created_at,
				m_mitra_administrator.mitra_adm_updated_by,
				m_mitra_administrator.mitra_adm_updated_at,
				m_mitra_administrator.mitra_adm_deleted_by,
				m_mitra_administrator.mitra_adm_deleted_at`).
		Joins("JOIN m_mitra_employee on m_mitra_administrator.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra_group_access on m_mitra_administrator.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Joins("JOIN m_address on m_mitra_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Where("mitra_adm_status = ?", status).
		Find(&mitraadministrator).Error
	if err != nil {
		return &[]MitraAdministratorData{}, err
	}
	return &mitraadministrator, nil
}

func (p *MitraAdministrator) FindMitraAdministratorDataByID(db *gorm.DB, pid uint64) (*MitraAdministrator, error) {
	var err error
	err = db.Debug().Model(&MitraAdministrator{}).Where("mitra_adm_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraAdministrator{}, err
	}
	return p, nil
}

func (p *MitraAdministrator) FindMitraAdministratorByID(db *gorm.DB, pid uint64) (*MitraAdministratorData, error) {
	var err error
	mitraadministrator := MitraAdministratorData{}
	err = db.Debug().Model(&MitraAdministratorData{}).Limit(100).
		Select(`m_mitra_administrator.mitra_adm_id,
				m_mitra_employee.mitra_emp_id,
				m_mitra_employee.mitra_emp_code,
				m_mitra_employee.mitra_emp_name,
				m_mitra_employee.mitra_emp_gender,
				m_mitra_employee.mitra_emp_birth_place,
				m_mitra_employee.mitra_emp_birth_date,
				m_mitra_employee.mitra_emp_address,
				m_mitra_employee.address_id,
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
				m_mitra_employee.mitra_emp_status,
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
				m_mitra_administrator.mitra_adm_created_at,
				m_mitra_administrator.mitra_adm_updated_by,
				m_mitra_administrator.mitra_adm_updated_at,
				m_mitra_administrator.mitra_adm_deleted_by,
				m_mitra_administrator.mitra_adm_deleted_at`).
		Joins("JOIN m_mitra_employee on m_mitra_administrator.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra_group_access on m_mitra_administrator.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Joins("JOIN m_address on m_mitra_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Where("mitra_adm_id = ?", pid).
		Find(&mitraadministrator).Error
	if err != nil {
		return &MitraAdministratorData{}, err
	}
	return &mitraadministrator, nil
}

func (p *MitraAdministrator) FindMitraAdministratorStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraAdministratorData, error) {
	var err error
	mitraadministrator := MitraAdministratorData{}
	err = db.Debug().Model(&MitraAdministratorData{}).Limit(100).
		Select(`m_mitra_administrator.mitra_adm_id,
				m_mitra_employee.mitra_emp_id,
				m_mitra_employee.mitra_emp_code,
				m_mitra_employee.mitra_emp_name,
				m_mitra_employee.mitra_emp_gender,
				m_mitra_employee.mitra_emp_birth_place,
				m_mitra_employee.mitra_emp_birth_date,
				m_mitra_employee.mitra_emp_address,
				m_mitra_employee.address_id,
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
				m_mitra_employee.mitra_emp_status,
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
				m_mitra_administrator.mitra_adm_created_at,
				m_mitra_administrator.mitra_adm_updated_by,
				m_mitra_administrator.mitra_adm_updated_at,
				m_mitra_administrator.mitra_adm_deleted_by,
				m_mitra_administrator.mitra_adm_deleted_at`).
		Joins("JOIN m_mitra_employee on m_mitra_administrator.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra_group_access on m_mitra_administrator.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Joins("JOIN m_address on m_mitra_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Where("mitra_adm_id = ? AND mitra_adm_status = ?", pid, status).
		Find(&mitraadministrator).Error
	if err != nil {
		return &MitraAdministratorData{}, err
	}
	return &mitraadministrator, nil
}

func (p *MitraAdministrator) UpdateMitraAdministrator(db *gorm.DB, pid uint64) (*MitraAdministrator, error) {
	var err error
	db = db.Debug().Model(&MitraAdministrator{}).Where("mitra_adm_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_group_accs_id":  p.MitraGroupAccessID,
			"mitra_adm_email":      p.MitraAdministratorEmail,
			"mitra_adm_username":   p.MitraAdministratorUsername,
			"mitra_adm_status":     p.MitraAdministratorStatus,
			"mitra_adm_updated_by": p.MitraAdministratorUpdatedBy,
			"mitra_adm_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraAdministrator{}).Where("mitra_adm_id = ?", pid).Error
	if err != nil {
		return &MitraAdministrator{}, err
	}
	return p, nil
}

func (p *MitraAdministrator) UpdateMitraAdministratorProfile(db *gorm.DB, pid uint64) (*MitraAdministrator, error) {
	var err error
	db = db.Debug().Model(&MitraAdministrator{}).Where("mitra_adm_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_adm_email":      p.MitraAdministratorEmail,
			"mitra_adm_username":   p.MitraAdministratorUsername,
			"mitra_adm_phone_code": p.MitraAdministratorPhoneCode,
			"mitra_adm_phone":      p.MitraAdministratorPhone,
			"mitra_adm_updated_by": p.MitraAdministratorUpdatedBy,
			"mitra_adm_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraAdministrator{}).Where("mitra_adm_id = ?", pid).Error
	if err != nil {
		return &MitraAdministrator{}, err
	}
	return p, nil
}

func (p *MitraAdministrator) BeforeUpdateMitraPassword() error {
	hashedPassword, err := HashMitraAdministrator(p.MitraAdministratorPassword)
	if err != nil {
		return err
	}
	p.MitraAdministratorPassword = string(hashedPassword)
	return nil
}

func (p *MitraAdministrator) UpdateMitraAdministratorPassword(db *gorm.DB, pid uint64) (*MitraAdministrator, error) {
	var err error
	err = p.BeforeUpdateMitraPassword()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&MitraAdministrator{}).Where("bbt_adm_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_adm_password":   p.MitraAdministratorPassword,
			"bbt_adm_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraAdministrator{}).Where("bbt_adm_id = ?", pid).Error
	if err != nil {
		return &MitraAdministrator{}, err
	}
	return p, nil
}

func (p *MitraAdministrator) UpdateMitraAdministratorStatus(db *gorm.DB, pid uint64) (*MitraAdministrator, error) {
	var err error
	db = db.Debug().Model(&MitraAdministrator{}).Where("mitra_adm_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_adm_status":     p.MitraAdministratorStatus,
			"mitra_adm_updated_by": p.MitraAdministratorUpdatedBy,
			"mitra_adm_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraAdministrator{}).Where("mitra_adm_id = ?", pid).Error
	if err != nil {
		return &MitraAdministrator{}, err
	}
	return p, nil
}

func (p *MitraAdministrator) UpdateMitraAdministratorStatusOnly(db *gorm.DB, pid uint64) (*MitraAdministrator, error) {
	var err error
	db = db.Debug().Model(&MitraAdministrator{}).Where("mitra_adm_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_adm_status": p.MitraAdministratorStatus,
		},
	)
	err = db.Debug().Model(&MitraAdministrator{}).Where("mitra_adm_id = ?", pid).Error
	if err != nil {
		return &MitraAdministrator{}, err
	}
	return p, nil
}

func (p *MitraAdministrator) UpdateMitraAdministratorStatusDelete(db *gorm.DB, pid uint64) (*MitraAdministrator, error) {
	var err error
	db = db.Debug().Model(&MitraAdministrator{}).Where("mitra_adm_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_adm_status":     3,
			"mitra_adm_deleted_by": p.MitraAdministratorDeletedBy,
			"mitra_adm_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraAdministrator{}).Where("mitra_adm_id = ?", pid).Error
	if err != nil {
		return &MitraAdministrator{}, err
	}
	return p, nil
}

func (p *MitraAdministrator) DeleteMitraAdministrator(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraAdministrator{}).Where("mitra_adm_id = ?", pid).Take(&MitraAdministrator{}).Delete(&MitraAdministrator{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraAdministrator not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//GET BY MITRA ID
//====================================================================================================================================

func (p *MitraAdministrator) FindAllMitraAdministratorMitra(db *gorm.DB, mitra uint64) (*[]MitraAdministratorData, error) {
	var err error
	mitraadministrator := []MitraAdministratorData{}
	err = db.Debug().Model(&MitraAdministratorData{}).Limit(100).
		Select(`m_mitra_administrator.mitra_adm_id,
				m_mitra_employee.mitra_emp_id,
				m_mitra_employee.mitra_emp_code,
				m_mitra_employee.mitra_emp_name,
				m_mitra_employee.mitra_emp_gender,
				m_mitra_employee.mitra_emp_birth_place,
				m_mitra_employee.mitra_emp_birth_date,
				m_mitra_employee.mitra_emp_address,
				m_mitra_employee.address_id,
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
				m_mitra_employee.mitra_emp_status,
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
				m_mitra_administrator.mitra_adm_created_at,
				m_mitra_administrator.mitra_adm_updated_by,
				m_mitra_administrator.mitra_adm_updated_at,
				m_mitra_administrator.mitra_adm_deleted_by,
				m_mitra_administrator.mitra_adm_deleted_at`).
		Joins("JOIN m_mitra_employee on m_mitra_administrator.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra_group_access on m_mitra_administrator.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Joins("JOIN m_address on m_mitra_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Where("m_mitra_employee.mitra_id = ?", mitra).
		Find(&mitraadministrator).Error
	if err != nil {
		return &[]MitraAdministratorData{}, err
	}
	return &mitraadministrator, nil
}

func (p *MitraAdministrator) FindAllMitraAdministratorStatusMitra(db *gorm.DB, status uint64, mitra uint64) (*[]MitraAdministratorData, error) {
	var err error
	mitraadministrator := []MitraAdministratorData{}
	err = db.Debug().Model(&MitraAdministratorData{}).Limit(100).
		Select(`m_mitra_administrator.mitra_adm_id,
				m_mitra_employee.mitra_emp_id,
				m_mitra_employee.mitra_emp_code,
				m_mitra_employee.mitra_emp_name,
				m_mitra_employee.mitra_emp_gender,
				m_mitra_employee.mitra_emp_birth_place,
				m_mitra_employee.mitra_emp_birth_date,
				m_mitra_employee.mitra_emp_address,
				m_mitra_employee.address_id,
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
				m_mitra_employee.mitra_emp_status,
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
				m_mitra_administrator.mitra_adm_created_at,
				m_mitra_administrator.mitra_adm_updated_by,
				m_mitra_administrator.mitra_adm_updated_at,
				m_mitra_administrator.mitra_adm_deleted_by,
				m_mitra_administrator.mitra_adm_deleted_at`).
		Joins("JOIN m_mitra_employee on m_mitra_administrator.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra_group_access on m_mitra_administrator.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Joins("JOIN m_address on m_mitra_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Where("mitra_adm_status = ? AND m_mitra_employee.mitra_id = ?", status, mitra).
		Find(&mitraadministrator).Error
	if err != nil {
		return &[]MitraAdministratorData{}, err
	}
	return &mitraadministrator, nil
}

func (p *MitraAdministrator) FindMitraAdministratorDataByIDMitra(db *gorm.DB, pid uint64, mitra uint64) (*MitraAdministrator, error) {
	var err error
	err = db.Debug().Model(&MitraAdministrator{}).Where("mitra_adm_id = ? AND m_mitra_employee.mitra_id = ?", pid, mitra).Take(&p).Error
	if err != nil {
		return &MitraAdministrator{}, err
	}
	return p, nil
}

func (p *MitraAdministrator) FindMitraAdministratorByIDMitra(db *gorm.DB, pid uint64, mitra uint64) (*MitraAdministratorData, error) {
	var err error
	mitraadministrator := MitraAdministratorData{}
	err = db.Debug().Model(&MitraAdministratorData{}).Limit(100).
		Select(`m_mitra_administrator.mitra_adm_id,
				m_mitra_employee.mitra_emp_id,
				m_mitra_employee.mitra_emp_code,
				m_mitra_employee.mitra_emp_name,
				m_mitra_employee.mitra_emp_gender,
				m_mitra_employee.mitra_emp_birth_place,
				m_mitra_employee.mitra_emp_birth_date,
				m_mitra_employee.mitra_emp_address,
				m_mitra_employee.address_id,
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
				m_mitra_employee.mitra_emp_status,
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
				m_mitra_administrator.mitra_adm_created_at,
				m_mitra_administrator.mitra_adm_updated_by,
				m_mitra_administrator.mitra_adm_updated_at,
				m_mitra_administrator.mitra_adm_deleted_by,
				m_mitra_administrator.mitra_adm_deleted_at`).
		Joins("JOIN m_mitra_employee on m_mitra_administrator.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra_group_access on m_mitra_administrator.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Joins("JOIN m_address on m_mitra_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Where("mitra_adm_id = ? AND m_mitra_employee.mitra_id = ?", pid, mitra).
		Find(&mitraadministrator).Error
	if err != nil {
		return &MitraAdministratorData{}, err
	}
	return &mitraadministrator, nil
}

func (p *MitraAdministrator) FindMitraAdministratorStatusByIDMitra(db *gorm.DB, pid uint64, status uint64, mitra uint64) (*MitraAdministratorData, error) {
	var err error
	mitraadministrator := MitraAdministratorData{}
	err = db.Debug().Model(&MitraAdministratorData{}).Limit(100).
		Select(`m_mitra_administrator.mitra_adm_id,
				m_mitra_employee.mitra_emp_id,
				m_mitra_employee.mitra_emp_code,
				m_mitra_employee.mitra_emp_name,
				m_mitra_employee.mitra_emp_gender,
				m_mitra_employee.mitra_emp_birth_place,
				m_mitra_employee.mitra_emp_birth_date,
				m_mitra_employee.mitra_emp_address,
				m_mitra_employee.address_id,
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
				m_mitra_employee.mitra_emp_status,
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
				m_mitra_administrator.mitra_adm_created_at,
				m_mitra_administrator.mitra_adm_updated_by,
				m_mitra_administrator.mitra_adm_updated_at,
				m_mitra_administrator.mitra_adm_deleted_by,
				m_mitra_administrator.mitra_adm_deleted_at`).
		Joins("JOIN m_mitra_employee on m_mitra_administrator.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra_group_access on m_mitra_administrator.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Joins("JOIN m_address on m_mitra_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Where("mitra_adm_id = ? AND mitra_adm_status = ? AND m_mitra_employee.mitra_id = ?", pid, status, mitra).
		Find(&mitraadministrator).Error
	if err != nil {
		return &MitraAdministratorData{}, err
	}
	return &mitraadministrator, nil
}

//MITRA LOGIN USER
//====================================================================================================================================

func MitraAdministratorVerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (p *MitraAdministrator) UpdateMitraAdministratorLastLogin(db *gorm.DB, username string) (*MitraAdministrator, error) {

	var err error
	db = db.Debug().Model(&MitraAdministrator{}).Where("mitra_adm_username = ?", username).UpdateColumns(
		map[string]interface{}{
			"mitra_adm_last_login": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraAdministrator{}).Where("mitra_adm_username = ?", username).Error
	if err != nil {
		return &MitraAdministrator{}, err
	}
	return p, nil
}

//ADDITIONAL
//====================================================================================================================================

func (p *MitraAdministrator) FindMitraAdministratorStatusByUsername(db *gorm.DB, username string) (*MitraAdministratorData, error) {
	var err error
	mitraadministrator := MitraAdministratorData{}
	err = db.Debug().Model(&MitraAdministratorData{}).
		Select(`m_mitra_administrator.mitra_adm_id,
				m_mitra_administrator.mitra_emp_id,
				m_mitra_employee.mitra_id,
				m_mitra_employee.mitra_branch_id,
				m_mitra_employee.mitra_emp_code,
				m_mitra_employee.mitra_emp_name,
				m_mitra_employee.mitra_emp_gender,
				m_mitra_employee.mitra_emp_birth_place,
				m_mitra_employee.mitra_emp_birth_date,
				m_mitra_employee.mitra_emp_address,
				m_mitra_employee.address_id,
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
				m_mitra_administrator.mitra_adm_ip_address,
				m_mitra_administrator.mitra_adm_last_login,
				m_mitra_administrator.mitra_adm_status,
				m_mitra_administrator.mitra_adm_is_new,
				m_mitra_administrator.mitra_adm_is_req_chpw,
				m_mitra_administrator.mitra_adm_is_locked,
				m_mitra_administrator.mitra_adm_fail_attempt,
				m_mitra_administrator.mitra_adm_locked_reason,
				m_mitra_administrator.mitra_adm_created_at,
				m_mitra_administrator.mitra_adm_updated_at,
				m_mitra_administrator.mitra_adm_deleted_at`).
		Joins("JOIN m_mitra_employee on m_mitra_administrator.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra_group_access on m_mitra_administrator.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Where("mitra_adm_username = ? AND mitra_adm_status = ?", username, 1).
		Take(&mitraadministrator).Error
	if err != nil {
		return &MitraAdministratorData{}, err
	}
	return &mitraadministrator, nil
}

func (p *MitraAdministrator) CountMitraAdministratorStatusByUsername(db *gorm.DB, username string) (*MitraAdministratorData, error) {
	var err error
	mitraadministrator := MitraAdministratorData{}
	err = db.Debug().Model(&MitraAdministratorData{}).
		Select(`m_mitra_administrator.mitra_adm_username,
				m_mitra_administrator.mitra_adm_status`).
		Where("mitra_adm_username = ? AND mitra_adm_status = ?", username, 1).
		Count(&mitraadministrator.MitraAdministratorCount).
		Take(&mitraadministrator).Error
	if err != nil {
		return &MitraAdministratorData{}, err
	}
	return &mitraadministrator, nil
}
