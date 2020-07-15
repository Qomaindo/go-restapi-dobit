package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//BASE CRUD
//====================================================================================================================================

type MitraUser struct {
	MitraUserID                      uint64     `gorm:"column:mitra_user_id;primary_key;" json:"mitra_user_id"`
	MitraEmployeeID                  uint64     `gorm:"column:mitra_emp_id" json:"mitra_emp_id"`
	MitraGroupAccessID               uint64     `gorm:"column:mitra_group_accs_id" json:"mitra_group_accs_id"`
	MitraGroupProductLimitID         uint64     `gorm:"column:mitra_group_prod_lim_id" json:"mitra_group_prod_lim_id"`
	MitraAOCoverageID                uint64     `gorm:"column:mitra_ao_cov_id" json:"mitra_ao_cov_id"`
	MitraUserReferalCode             string     `gorm:"column:mitra_user_referal_code;size:25" json:"mitra_user_referal_code"`
	MitraUserEmail                   string     `gorm:"column:mitra_user_email;size:255" json:"mitra_user_email"`
	MitraUsername                    string     `gorm:"column:mitra_username;size:125" json:"mitra_username"`
	MitraUserPasswordWeb             string     `gorm:"column:mitra_user_password_web;size:255" json:"mitra_user_password_web"`
	MitraUserPasswordMobile          string     `gorm:"column:mitra_user_password_mobile;size:255" json:"mitra_user_password_mobile"`
	MitraUserPhoneCode               string     `gorm:"column:mitra_user_phone_code;size:6" json:"mitra_user_phone_code"`
	MitraUserPhone                   string     `gorm:"column:mitra_user_phone;size:20" json:"mitra_user_phone"`
	MitraUserOTP                     string     `gorm:"column:mitra_user_otp;size:6" json:"mitra_user_otp"`
	MitraUserOTPSecret               *time.Time `gorm:"column:mitra_user_otp_secret" json:"mitra_user_otp_secret"`
	MitraUserSecret                  uint64     `gorm:"column:mitra_user_secret" json:"mitra_user_secret"`
	MitraUserPIN                     string     `gorm:"column:mitra_user_pin" json:"mitra_user_pin"`
	MitraUserLastLogin               *time.Time `gorm:"column:mitra_user_last_login" json:"mitra_user_last_login"`
	MitraUserImei                    string     `gorm:"column:mitra_user_imei;size:50" json:"mitra_user_imei"`
	MitraUserIPAddress               string     `gorm:"column:mitra_user_ip_address;size:50" json:"mitra_user_ip_address"`
	MitraUserIsAO                    bool       `gorm:"column:mitra_user_is_ao" json:"mitra_user_is_ao"`
	MitraUserIsCredit                bool       `gorm:"column:mitra_user_is_credit" json:"mitra_user_is_credit"`
	MitraUserStatus                  uint64     `gorm:"column:mitra_user_status;size:1" json:"mitra_user_status"`
	MitraUserIsNew                   bool       `gorm:"column:mitra_user_is_new" json:"mitra_user_is_new"`
	MitraUserIsRequestChangePassword bool       `gorm:"column:mitra_user_is_req_chpw" json:"mitra_user_is_req_chpw"`
	MitraUserIsLocked                bool       `gorm:"column:mitra_user_is_locked" json:"mitra_user_is_locked"`
	MitraUserFailAttempt             uint64     `gorm:"column:mitra_user_fail_attempt;size:1" json:"mitra_user_fail_attempt"`
	MitraUserLockedReason            string     `gorm:"column:mitra_user_locked_reason;size:255" json:"mitra_user_locked_reason"`
	MitraUserCreatedBy               string     `gorm:"column:mitra_user_created_by;size:125" json:"mitra_user_created_by"`
	MitraUserUpdatedBy               string     `gorm:"column:mitra_user_updated_by;size:125" json:"mitra_user_updated_by"`
	MitraUserDeletedBy               string     `gorm:"column:mitra_user_deleted_by;size:125" json:"mitra_user_deleted_by"`
	MitraUserCreatedAt               time.Time  `gorm:"column:mitra_user_created_at;default:CURRENT_TIMESTAMP" json:"mitra_user_created_at"`
	MitraUserUpdatedAt               *time.Time `gorm:"column:mitra_user_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_user_updated_at"`
	MitraUserDeletedAt               *time.Time `gorm:"column:mitra_user_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_user_deleted_at"`

	MitraLoginStaticToken string    `gorm:"-" json:"mitra_login_static_token"`
	MitraUserPINOld       string    `gorm:"-" json:"mitra_user_pin_old"`
	MitraUserPINNew       string    `gorm:"-" json:"mitra_user_pin_new"`
	MitraUserOTPNew       string    `gorm:"-" json:"mitra_user_otp_new"`
	MitraCode             string    `gorm:"-" json:"mitra_code"`
	MitraUserOTPSecretNew time.Time `gorm:"-"`
	Msg_id                string    `gorm:"-"`
	Trx_id                string    `gorm:"-"`
	Status                uint64    `gorm:"-"`
	Secret                string    `gorm:"-"`
	Segment               uint64    `gorm:"-"`

	MitraUserPasswordNew     string `gorm:"-" json:"mitra_user_pw_new"`
	MitraUserPasswordKonfirm string `gorm:"-" json:"mitra_user_pw_konfirm"`
	MitraUserPasswordOld     string `gorm:"-" json:"mitra_user_pw_old"`
}

type MitraUserData struct {
	MitraUserID                      uint64     `gorm:"column:mitra_user_id;primary_key;" json:"mitra_user_id"`
	MitraID                          uint64     `gorm:"column:mitra_id" json:"mitra_id"`
	MitraName                        string     `gorm:"column:mitra_name" json:"mitra_name"`
	MitraLogo                        string     `gorm:"column:mitra_logo" json:"mitra_logo"`
	MitraBranchID                    uint64     `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	MitraBranchName                  string     `gorm:"column:mitra_branch_name" json:"mitra_branch_name"`
	MitraEmployeeID                  uint64     `gorm:"column:mitra_emp_id" json:"mitra_emp_id"`
	MitraEmployeeName                string     `gorm:"column:mitra_emp_name" json:"mitra_emp_name"`
	MitraEmployeeGender              string     `gorm:"column:mitra_emp_gender" json:"mitra_emp_gender"`
	MitraEmployeeBirthPlace          string     `gorm:"column:mitra_emp_birth_place" json:"mitra_emp_birth_place"`
	MitraEmployeeBirthDate           string     `gorm:"column:mitra_emp_birth_date" json:"mitra_emp_birth_date"`
	MitraEmployeeAddress             string     `gorm:"column:mitra_emp_address" json:"mitra_emp_address"`
	AddressID                        uint64     `gorm:"column:address_id" json:"address_id"`
	MitraEmployeePhoto               string     `gorm:"column:mitra_emp_photo" json:"mitra_emp_photo"`
	MitraGroupAccessID               uint64     `gorm:"column:mitra_group_accs_id" json:"mitra_group_accs_id"`
	MitraGroupAccessName             string     `gorm:"column:mitra_group_accs_name" json:"mitra_group_accs_name"`
	MitraGroupAccessRole             string     `gorm:"column:mitra_group_accs_role" json:"mitra_group_accs_role"`
	MitraGroupProductLimitID         uint64     `gorm:"column:mitra_group_prod_lim_id" json:"mitra_group_prod_lim_id"`
	MitraGroupProductLimitName       string     `gorm:"column:mitra_group_prod_lim_name" json:"mitra_group_prod_lim_name"`
	MitraAOCoverageID                uint64     `gorm:"column:mitra_ao_cov_id" json:"mitra_ao_cov_id"`
	MitraAOCoverageName              string     `gorm:"column:mitra_ao_cov_name" json:"mitra_ao_cov_name"`
	MitraUserReferalCode             string     `gorm:"column:mitra_user_referal_code" json:"mitra_user_referal_code"`
	MitraUserEmail                   string     `gorm:"column:mitra_user_email" json:"mitra_user_email"`
	MitraUsername                    string     `gorm:"column:mitra_username" json:"mitra_username"`
	MitraUserPIN                     string     `gorm:"column:mitra_user_pin;size:6" json:"mitra_user_pin"`
	MitraUserPhoneCode               string     `gorm:"column:mitra_user_phone_code" json:"mitra_user_phone_code"`
	MitraUserPhone                   string     `gorm:"column:mitra_user_phone" json:"mitra_user_phone"`
	MitraUserPasswordWeb             string     `gorm:"column:mitra_user_password_web;size:255" json:"mitra_user_password_web"`
	MitraUserPasswordMobile          string     `gorm:"column:mitra_user_password_mobile;size:255" json:"mitra_user_password_mobile"`
	MitraUserOTP                     string     `gorm:"column:mitra_user_otp" json:"mitra_user_otp"`
	MitraUserOTPSecret               *time.Time `gorm:"column:mitra_user_otp_secret;default:CURRENT_TIMESTAMP" json:"mitra_user_otp_secret"`
	MitraUserSecret                  uint64     `gorm:"column:mitra_user_secret" json:"mitra_user_secret"`
	MitraUserLastLogin               time.Time  `gorm:"column:mitra_user_last_login" json:"mitra_user_last_login"`
	MitraUserImei                    string     `gorm:"column:mitra_user_imei;size:50" json:"mitra_user_imei"`
	MitraUserIPAddress               string     `gorm:"column:mitra_user_ip_address;size:50" json:"mitra_user_ip_address"`
	MitraUserIsAO                    bool       `gorm:"column:mitra_user_is_ao" json:"mitra_user_is_ao"`
	MitraUserIsCredit                bool       `gorm:"column:mitra_user_is_credit" json:"mitra_user_is_credit"`
	MitraUserStatus                  uint64     `gorm:"column:mitra_user_status" json:"mitra_user_status"`
	MitraUserIsNew                   bool       `gorm:"column:mitra_user_is_new" json:"mitra_user_is_new"`
	MitraUserIsRequestChangePassword bool       `gorm:"column:mitra_user_is_req_chpw" json:"mitra_user_is_req_chpw"`
	MitraUserIsLocked                bool       `gorm:"column:mitra_user_is_locked" json:"mitra_user_is_locked"`
	MitraUserFailAttempt             uint64     `gorm:"column:mitra_user_fail_attempt" json:"mitra_user_fail_attempt"`
	MitraUserLockedReason            string     `gorm:"column:mitra_user_locked_reason" json:"mitra_user_locked_reason"`
	MitraUserCreatedBy               string     `gorm:"column:mitra_user_created_by;size:125" json:"mitra_user_created_by"`
	MitraUserCreatedAt               time.Time  `gorm:"column:mitra_user_created_at;default:CURRENT_TIMESTAMP" json:"mitra_user_created_at"`
	MitraUserUpdatedBy               string     `gorm:"column:mitra_user_updated_by;size:125" json:"mitra_user_updated_by"`
	MitraUserUpdatedAt               *time.Time `gorm:"column:mitra_user_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_user_updated_at"`
	MitraUserDeletedBy               string     `gorm:"column:mitra_user_deleted_by;size:125" json:"mitra_user_deleted_by"`
	MitraUserDeletedAt               *time.Time `gorm:"column:mitra_user_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_user_deleted_at"`

	MitraUserBranchAccess []MitraUserBranchAccessData `json:"mitra_user_branch_accs"`

	MitraUserCount int `gorm:"-"`
}

type ResponseMitraUsers struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    *[]MitraUserData `json:"data"`
}

type ResponseMitraUser struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    *MitraUserData `json:"data"`
}

type ResponseMitraUserCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraUser) TableName() string {
	return "m_mitra_user"
}

func (MitraUserData) TableName() string {
	return "m_mitra_user"
}

func (p *MitraUser) Prepare() {
	p.MitraEmployeeID = p.MitraEmployeeID
	p.MitraGroupAccessID = p.MitraGroupAccessID
	p.MitraGroupProductLimitID = p.MitraGroupProductLimitID
	p.MitraAOCoverageID = p.MitraAOCoverageID
	p.MitraUserReferalCode = html.EscapeString(strings.TrimSpace(p.MitraUserReferalCode))
	p.MitraUserEmail = html.EscapeString(strings.TrimSpace(p.MitraUserEmail))
	p.MitraUsername = html.EscapeString(strings.TrimSpace(p.MitraUsername))
	p.MitraUserPasswordWeb = html.EscapeString(strings.TrimSpace(p.MitraUserPasswordWeb))
	p.MitraUserPasswordMobile = html.EscapeString(strings.TrimSpace(p.MitraUserPasswordMobile))
	p.MitraUserPhoneCode = html.EscapeString(strings.TrimSpace(p.MitraUserPhoneCode))
	p.MitraUserPhone = html.EscapeString(strings.TrimSpace(p.MitraUserPhone))
	p.MitraUserOTP = html.EscapeString(strings.TrimSpace(p.MitraUserOTP))
	p.MitraUserOTPSecret = nil
	p.MitraUserSecret = p.MitraUserSecret
	p.MitraUserPIN = html.EscapeString(strings.TrimSpace(p.MitraUserPIN))
	p.MitraUserLastLogin = nil
	p.MitraUserImei = p.MitraUserImei
	p.MitraUserIPAddress = p.MitraUserIPAddress
	p.MitraUserIsAO = p.MitraUserIsAO
	p.MitraUserIsCredit = p.MitraUserIsCredit
	p.MitraUserStatus = p.MitraUserStatus
	p.MitraUserIsNew = p.MitraUserIsNew
	p.MitraUserIsRequestChangePassword = p.MitraUserIsRequestChangePassword
	p.MitraUserIsLocked = p.MitraUserIsLocked
	p.MitraUserFailAttempt = p.MitraUserFailAttempt
	p.MitraUserLockedReason = html.EscapeString(strings.TrimSpace(p.MitraUserLockedReason))
	p.MitraUserCreatedBy = p.MitraUserCreatedBy
	p.MitraUserUpdatedBy = p.MitraUserUpdatedBy
	p.MitraUserDeletedBy = p.MitraUserDeletedBy
	p.MitraUserCreatedAt = time.Now()
	p.MitraUserUpdatedAt = p.MitraUserUpdatedAt
	p.MitraUserDeletedAt = p.MitraUserDeletedAt
}

func (p *MitraUser) Validate(action string) error {
	switch strings.ToLower(action) {

	case "login":
		if p.MitraUsername == "" {
			return errors.New("Username Wajib Diisi")
		}
		if p.MitraUserPasswordWeb == "" {
			return errors.New("Password Wajib Diisi")
		}
		return nil

	case "loginmobile":
		if p.MitraUsername == "" {
			return errors.New("Username Wajib Diisi")
		}
		if p.MitraUserPasswordMobile == "" {
			return errors.New("Password Wajib Diisi")
		}
		return nil

	case "setpin":
		if p.MitraUserPIN == "" {
			return errors.New("Required Pin")
		}
		if p.MitraLoginStaticToken != "/@(Kh6f.7Qh_!_=y~R~!LfT_<S;)fj" {
			return errors.New("Required Static Token")
		}
		return nil

	case "username":
		if p.MitraUsername == "" {
			return errors.New("Username Wajib Diisi")
		}
		if p.MitraCode == "" {
			return errors.New("Required Mitra Code")
		}
		if p.MitraUserPasswordWeb == "" {
			return errors.New("Password Wajib Diisi")
		}
		if p.MitraLoginStaticToken != "/@(Kh6f.7Qh_!_=y~R~!LfT_<S;)fj" {
			return errors.New("Required Static Token")
		}
		return nil

	case "usernamemobile":
		if p.MitraUsername == "" {
			return errors.New("Username Wajib Diisi")
		}
		if p.MitraCode == "" {
			return errors.New("Required Mitra Code")
		}
		if p.MitraUserPasswordMobile == "" {
			return errors.New("Password Wajib Diisi")
		}
		if p.MitraLoginStaticToken != "/@(Kh6f.7Qh_!_=y~R~!LfT_<S;)fj" {
			return errors.New("Required Static Token")
		}
		return nil

	case "phone":
		if p.MitraUserPhone == "" {
			return errors.New("Required Phone")
		}
		if p.MitraLoginStaticToken != "/@(Kh6f.7Qh_!_=y~R~!LfT_<S;)fj" {
			return errors.New("Required Static Token")
		}
		return nil

	case "phone-profile":
		if p.MitraUserPhone == "" {
			return errors.New("Required Phone")
		}
		return nil

	case "verifyotp":
		if p.MitraUserOTPNew == "" {
			return errors.New("Required OTP")
		}
		if p.MitraUserPhone == "" {
			return errors.New("Required Phone")
		}
		return nil

	case "verifypin":
		if p.MitraUserPIN == "" {
			return errors.New("Required Pin")
		}
		// if p.MitraUserPhone == "" {
		// 	return errors.New("Required Phone")
		// }
		return nil

	case "verifypswdmobile":
		if p.MitraUserPasswordMobile == "" {
			return errors.New("Required Pin")
		}
		if p.MitraUserPhone == "" {
			return errors.New("Required Phone")
		}
		return nil

	case "updatepin":
		if p.MitraUserPINNew == "" {
			return errors.New("Required Pin")
		}
		if p.MitraUserPINOld == "" {
			return errors.New("Required Phone")
		}
		return nil

	case "updatepasswordmobile":
		if p.MitraUserPasswordNew == "" {
			return errors.New("Required Password")
		}
		if p.MitraUserPasswordOld == "" {
			return errors.New("Required Phone")
		}
		return nil

	case "updatepasswordweb":
		if p.MitraUserPasswordNew == "" {
			return errors.New("Required New Password")
		}
		if p.MitraUserPasswordOld == "" {
			return errors.New("Required Old Password")
		}
		if p.MitraUserPasswordKonfirm == "" {
			return errors.New("Required Konfirm Password")
		}
		return nil

	case "pin":
		if p.MitraUserPIN == "" {
			return errors.New("Required Pin")
		}
		if p.MitraUserPhone == "" {
			return errors.New("Required Phone")
		}
		return nil

	case "update":
		if p.MitraUserEmail == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(p.MitraUserEmail); err != nil {
			return errors.New("Invalid Email")
		}
		if p.MitraUsername == "" {
			return errors.New("Username Wajib Diisi")
		}
		if p.MitraUserPasswordWeb == "" {
			return errors.New("Password Wajib Diisi")
		}
		if p.MitraUserPhone == "" {
			return errors.New("Required Phone")
		}
		if p.MitraUserOTP == "" {
			return errors.New("Required OTP")
		}
		if p.MitraUserPIN == "" {
			return errors.New("Required Pin")
		}
		if p.MitraUserImei == "" {
			return errors.New("Required IMEI")
		}
		if p.MitraUserIPAddress == "" {
			return errors.New("Required IP Address")
		}
		if p.MitraUserStatus == 0 {
			return errors.New("Required User Status")
		}
		return nil

	default:
		if p.MitraUserPasswordWeb == "" {
			return errors.New("Required MitraUser Password")
		}
		if p.MitraUserEmail == "" {
			return errors.New("Required MitraUser Email")
		}
		return nil
	}
}

func (p *MitraUser) SaveMitraUser(db *gorm.DB) (*MitraUser, error) {
	var err error
	err = db.Debug().Model(&MitraUser{}).Create(&p).Error
	if err != nil {
		return &MitraUser{}, err
	}
	return p, nil
}

func (p *MitraUser) FindAllMitraUser(db *gorm.DB, mitra uint64) (*[]MitraUserData, error) {
	var err error
	mitrauser := []MitraUserData{}
	err = db.Debug().Model(&MitraUserData{}).Limit(100).
		Select(`m_mitra_user.mitra_user_id,
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
				m_mitra.mitra_logo,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_group_access.mitra_group_accs_id,
				m_mitra_group_access.mitra_group_accs_name,
				m_mitra_group_access.mitra_group_accs_role,
				m_mitra_group_product_limit.mitra_group_prod_lim_id,
				m_mitra_group_product_limit.mitra_group_prod_lim_name,
				m_mitra_ao_coverage.mitra_ao_cov_id as mitra_ao_cov_id,
				m_mitra_ao_coverage.mitra_ao_cov_name as mitra_ao_cov_name,
				m_mitra_user.mitra_user_referal_code,
				m_mitra_user.mitra_user_email,
				m_mitra_user.mitra_username,
				m_mitra_user.mitra_user_phone_code,
				m_mitra_user.mitra_user_phone,
				m_mitra_user.mitra_user_imei,
				m_mitra_user.mitra_user_ip_address,
				m_mitra_user.mitra_user_is_ao,
				m_mitra_user.mitra_user_is_credit,
				m_mitra_user.mitra_user_status,
				m_mitra_user.mitra_user_is_new,
				m_mitra_user.mitra_user_is_req_chpw,
				m_mitra_user.mitra_user_is_locked,
				m_mitra_user.mitra_user_fail_attempt,
				m_mitra_user.mitra_user_locked_reason,
				m_mitra_user.mitra_user_created_by,
				m_mitra_user.mitra_user_updated_by,
				m_mitra_user.mitra_user_deleted_by,
				m_mitra_user.mitra_user_created_at,
				m_mitra_user.mitra_user_updated_at,
				m_mitra_user.mitra_user_deleted_at`).
		Joins("JOIN m_mitra_employee on m_mitra_user.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra_group_access on m_mitra_user.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Joins("JOIN m_mitra_group_product_limit on m_mitra_user.mitra_group_prod_lim_id=m_mitra_group_product_limit.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_ao_coverage on m_mitra_user.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Where("m_mitra_employee.mitra_id = ?", mitra).
		Find(&mitrauser).Error
	if err != nil {
		return &[]MitraUserData{}, err
	}
	return &mitrauser, nil
}

func (p *MitraUser) FindAllMitraUserStatus(db *gorm.DB, status uint64, mitra uint64) (*[]MitraUserData, error) {
	var err error
	mitrauser := []MitraUserData{}
	err = db.Debug().Model(&MitraUserData{}).Limit(100).
		Select(`m_mitra_user.mitra_user_id,
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
				m_mitra.mitra_logo,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_group_access.mitra_group_accs_id,
				m_mitra_group_access.mitra_group_accs_name,
				m_mitra_group_access.mitra_group_accs_role,
				m_mitra_group_product_limit.mitra_group_prod_lim_id,
				m_mitra_group_product_limit.mitra_group_prod_lim_name,
				m_mitra_ao_coverage.mitra_ao_cov_id as mitra_ao_cov_id,
				m_mitra_ao_coverage.mitra_ao_cov_name as mitra_ao_cov_name,
				m_mitra_user.mitra_user_referal_code,
				m_mitra_user.mitra_user_email,
				m_mitra_user.mitra_username,
				m_mitra_user.mitra_user_phone_code,
				m_mitra_user.mitra_user_phone,
				m_mitra_user.mitra_user_imei,
				m_mitra_user.mitra_user_ip_address,
				m_mitra_user.mitra_user_is_ao,
				m_mitra_user.mitra_user_is_credit,
				m_mitra_user.mitra_user_status,
				m_mitra_user.mitra_user_is_new,
				m_mitra_user.mitra_user_is_req_chpw,
				m_mitra_user.mitra_user_is_locked,
				m_mitra_user.mitra_user_fail_attempt,
				m_mitra_user.mitra_user_locked_reason,
				m_mitra_user.mitra_user_created_by,
				m_mitra_user.mitra_user_updated_by,
				m_mitra_user.mitra_user_deleted_by,
				m_mitra_user.mitra_user_created_at,
				m_mitra_user.mitra_user_updated_at,
				m_mitra_user.mitra_user_deleted_at`).
		Joins("JOIN m_mitra_employee on m_mitra_user.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra_group_access on m_mitra_user.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Joins("JOIN m_mitra_group_product_limit on m_mitra_user.mitra_group_prod_lim_id=m_mitra_group_product_limit.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_ao_coverage on m_mitra_user.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Where("mitra_user_status = ? AND m_mitra_employee.mitra_id = ?", status, mitra).
		Find(&mitrauser).Error
	if err != nil {
		return &[]MitraUserData{}, err
	}
	return &mitrauser, nil
}

func (p *MitraUser) FindMitraUserByID(db *gorm.DB, pid uint64) (*MitraUserData, error) {
	var err error
	mitrauser := MitraUserData{}
	err = db.Debug().Model(&MitraUserData{}).
		Select(`m_mitra_user.mitra_user_id,
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
				m_mitra.mitra_logo,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_group_access.mitra_group_accs_id,
				m_mitra_group_access.mitra_group_accs_name,
				m_mitra_group_access.mitra_group_accs_role,
				m_mitra_group_product_limit.mitra_group_prod_lim_id,
				m_mitra_group_product_limit.mitra_group_prod_lim_name,
				m_mitra_ao_coverage.mitra_ao_cov_id as mitra_ao_cov_id,
				m_mitra_ao_coverage.mitra_ao_cov_name as mitra_ao_cov_name,
				m_mitra_user.mitra_user_referal_code,
				m_mitra_user.mitra_user_email,
				m_mitra_user.mitra_username,
				m_mitra_user.mitra_user_pin,
				m_mitra_user.mitra_user_password_mobile,
				m_mitra_user.mitra_user_phone_code,
				m_mitra_user.mitra_user_phone,
				m_mitra_user.mitra_user_imei,
				m_mitra_user.mitra_user_ip_address,
				m_mitra_user.mitra_user_is_ao,
				m_mitra_user.mitra_user_is_credit,
				m_mitra_user.mitra_user_status,
				m_mitra_user.mitra_user_is_new,
				m_mitra_user.mitra_user_is_req_chpw,
				m_mitra_user.mitra_user_is_locked,
				m_mitra_user.mitra_user_fail_attempt,
				m_mitra_user.mitra_user_locked_reason,
				m_mitra_user.mitra_user_created_by,
				m_mitra_user.mitra_user_updated_by,
				m_mitra_user.mitra_user_deleted_by,
				m_mitra_user.mitra_user_created_at,
				m_mitra_user.mitra_user_updated_at,
				m_mitra_user.mitra_user_deleted_at`).
		Joins("JOIN m_mitra_employee on m_mitra_user.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra_group_access on m_mitra_user.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Joins("JOIN m_mitra_group_product_limit on m_mitra_user.mitra_group_prod_lim_id=m_mitra_group_product_limit.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_ao_coverage on m_mitra_user.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Where("mitra_user_id = ?", pid).
		Take(&mitrauser).Error
	if err != nil {
		return &MitraUserData{}, err
	}
	return &mitrauser, nil
}

func (p *MitraUser) FindMitraUserStatusByID(db *gorm.DB, pid uint64, status uint64, mitra uint64) (*MitraUserData, error) {
	var err error
	mitrauser := MitraUserData{}
	err = db.Debug().Model(&MitraUserData{}).
		Select(`m_mitra_user.mitra_user_id,
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
				m_mitra.mitra_logo,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_group_access.mitra_group_accs_id,
				m_mitra_group_access.mitra_group_accs_name,
				m_mitra_group_access.mitra_group_accs_role,
				m_mitra_group_product_limit.mitra_group_prod_lim_id,
				m_mitra_group_product_limit.mitra_group_prod_lim_name,
				m_mitra_ao_coverage.mitra_ao_cov_id as mitra_ao_cov_id,
				m_mitra_ao_coverage.mitra_ao_cov_name as mitra_ao_cov_name,
				m_mitra_user.mitra_user_referal_code,
				m_mitra_user.mitra_user_email,
				m_mitra_user.mitra_username,
				m_mitra_user.mitra_user_phone_code,
				m_mitra_user.mitra_user_phone,
				m_mitra_user.mitra_user_imei,
				m_mitra_user.mitra_user_ip_address,
				m_mitra_user.mitra_user_is_ao,
				m_mitra_user.mitra_user_is_credit,
				m_mitra_user.mitra_user_status,
				m_mitra_user.mitra_user_is_new,
				m_mitra_user.mitra_user_is_req_chpw,
				m_mitra_user.mitra_user_is_locked,
				m_mitra_user.mitra_user_fail_attempt,
				m_mitra_user.mitra_user_locked_reason,
				m_mitra_user.mitra_user_created_by,
				m_mitra_user.mitra_user_updated_by,
				m_mitra_user.mitra_user_deleted_by,
				m_mitra_user.mitra_user_created_at,
				m_mitra_user.mitra_user_updated_at,
				m_mitra_user.mitra_user_deleted_at`).
		Joins("JOIN m_mitra_employee on m_mitra_user.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra_group_access on m_mitra_user.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Joins("JOIN m_mitra_group_product_limit on m_mitra_user.mitra_group_prod_lim_id=m_mitra_group_product_limit.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_ao_coverage on m_mitra_user.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Where("mitra_user_id = ? AND mitra_user_status = ? AND m_mitra_employee.mitra_id = ?", pid, status, mitra).
		Take(&mitrauser).Error
	if err != nil {
		return &MitraUserData{}, err
	}
	return &mitrauser, nil
}

func (p *MitraUser) UpdateMitraUser(db *gorm.DB, pid uint64) (*MitraUser, error) {

	var err error
	db = db.Debug().Model(&MitraUser{}).Where("mitra_user_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_group_accs_id":     p.MitraGroupAccessID,
			"mitra_group_prod_lim_id": p.MitraGroupProductLimitID,
			"mitra_user_email":        p.MitraUserEmail,
			"mitra_username":          p.MitraUsername,
			"mitra_user_is_ao":        p.MitraUserIsAO,
			"mitra_user_is_credit":    p.MitraUserIsCredit,
			"mitra_user_status":       p.MitraUserStatus,
			"mitra_user_updated_by":   p.MitraUserUpdatedBy,
			"mitra_user_updated_at":   time.Now(),
		},
	)
	err = db.Debug().Model(&MitraUser{}).Where("mitra_user_id = ?", pid).Error
	if err != nil {
		return &MitraUser{}, err
	}
	return p, nil
}

func (p *MitraUser) UpdateMitraUserProfile(db *gorm.DB, pid uint64) (*MitraUser, error) {

	var err error
	db = db.Debug().Model(&MitraUser{}).Where("mitra_user_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_user_email":      p.MitraUserEmail,
			"mitra_username":        p.MitraUsername,
			"mitra_user_phone_code": p.MitraUserPhoneCode,
			"mitra_user_phone":      p.MitraUserPhone,
			"mitra_user_status":     p.MitraUserStatus,
			"mitra_user_updated_by": p.MitraUserUpdatedBy,
			"mitra_user_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraUser{}).Where("mitra_user_id = ?", pid).Error
	if err != nil {
		return &MitraUser{}, err
	}
	return p, nil
}

func (p *MitraUser) BeforeUpdateMitraPassword() error {
	hashedPassword, err := HashMitraUser(p.MitraUserPasswordWeb)
	if err != nil {
		return err
	}
	p.MitraUserPasswordWeb = string(hashedPassword)
	return nil
}

func (p *MitraUser) UpdateMitraUserPasswordWeb(db *gorm.DB, pid uint64) (*MitraUser, error) {
	var err error
	err = p.BeforeUpdateMitraPassword()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&MitraUser{}).Where("mitra_user_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_user_password_web": p.MitraUserPasswordWeb,
			"mitra_user_updated_at":   time.Now(),
		},
	)
	err = db.Debug().Model(&MitraUser{}).Where("mitra_user_id = ?", pid).Error
	if err != nil {
		return &MitraUser{}, err
	}
	return p, nil
}

func (p *MitraUser) UpdateMitraUserStatus(db *gorm.DB, pid uint64) (*MitraUser, error) {

	var err error
	db = db.Debug().Model(&MitraUser{}).Where("mitra_user_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_user_status":     p.MitraUserStatus,
			"mitra_user_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraUser{}).Where("mitra_user_id = ?", pid).Error
	if err != nil {
		return &MitraUser{}, err
	}
	return p, nil
}

func (p *MitraUser) UpdateMitraUserStatusOnly(db *gorm.DB, pid uint64) (*MitraUser, error) {

	var err error
	db = db.Debug().Model(&MitraUser{}).Where("mitra_user_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_user_status": p.MitraUserStatus,
		},
	)
	err = db.Debug().Model(&MitraUser{}).Where("mitra_user_id = ?", pid).Error
	if err != nil {
		return &MitraUser{}, err
	}
	return p, nil
}

func (p *MitraUser) UpdateMitraUserStatusDelete(db *gorm.DB, pid uint64) (*MitraUser, error) {
	var err error
	db = db.Debug().Model(&MitraUser{}).Where("mitra_user_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_user_status":     3,
			"mitra_user_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraUser{}).Where("mitra_user_id = ?", pid).Error
	if err != nil {
		return &MitraUser{}, err
	}
	return p, nil
}

func (p *MitraUser) DeleteMitraUser(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraUser{}).Where("mitra_user_id = ?", pid).Take(&MitraUser{}).Delete(&MitraUser{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraUser not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// Mitra User Login
// ==================================================================================

type ResponseMitraRegister struct {
	Status  uint64 `json:"status"`
	Message string `json:"message"`
}

type ResponseMitraPhoneData struct {
	Status uint64 `json:"status"`
	Name   string `json:"name"`
	Photo  string `json:"photo"`
}

type ResponseMitraPhone struct {
	Status  uint64                 `json:"status"`
	Message string                 `json:"message"`
	Data    ResponseMitraPhoneData `json:"data"`
}

type ResponseMitraUsernameData struct {
	Status uint64 `json:"status"`
	Phone  string `json:"phone"`
	Token  string `json:"token"`
}

type ResponseMitraUsername struct {
	Status  uint64                    `json:"status"`
	Message string                    `json:"message"`
	Data    ResponseMitraUsernameData `json:"data"`
}

type ResponseMitraTokenData struct {
	Token  string `json:"token"`
	Status string `json:"status"`
	Name   string `json:"name"`
	Photo  string `json:"photo"`
}

type ResponseMitraSendData struct {
	Status       uint64 `json:"status"`
	Name         string `json:"name"`
	SMSStatus    string `json:"sms_status"`
	UpdateStatus uint64 `json:"update_status"`
}

type ResponseMitraUserSend struct {
	Status  uint64                `json:"status"`
	Message string                `json:"message"`
	Data    ResponseMitraSendData `json:"data"`
}

type ResponseMitraVerifyData struct {
	Token     string `json:"token"`
	Status    uint64 `json:"status"`
	Name      string `json:"name"`
	SMSStatus string `json:"sms_status"`
}

type ResponseMitraUserVerify struct {
	Status  uint64                  `json:"status"`
	Message string                  `json:"message"`
	Data    ResponseMitraVerifyData `json:"data"`
}

type ResponseMitraUserToken struct {
	Status  uint64                 `json:"status"`
	Message string                 `json:"message"`
	Data    ResponseMitraTokenData `json:"data"`
}

type ResponseMitraOTPData struct {
	Status    uint64 `json:"status"`
	Condition string `json:"condition"`
	Token     string `json:"token"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Photo     string `json:"photo"`
}

type ResponseMitraOTP struct {
	Status  uint64                 `json:"status"`
	Message string                 `json:"message"`
	Data    ResponseMitraPhoneData `json:"data"`
}

type ResponseMitraPinData struct {
	Status string `json:"status"`
	Name   string `json:"name"`
	Photo  string `json:"photo"`
}

type ResponseMitraPin struct {
	Status  uint64                 `json:"status"`
	Message string                 `json:"message"`
	Data    ResponseMitraPhoneData `json:"data"`
}

type ResponseMitraPswd struct {
	Status  uint64                 `json:"status"`
	Message string                 `json:"message"`
	Data    ResponseMitraPhoneData `json:"data"`
}

type ResponseMitraCodeOTP struct {
	TimeSend time.Time `json:"time_send"`
	Message  string    `json:"message"`
	Code     string    `json:"code_otp"`
}

type ResponseMitraDashboardData struct {
	MitraID uint64 `json:"mitra_id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Status  uint64 `json:"status"`
}

type ResponseMitraDashboard struct {
	Status  uint64                     `json:"status"`
	Message string                     `json:"message"`
	Data    ResponseMitraDashboardData `json:"data"`
}

type ResponseMitraProfileData struct {
	MitraID     uint64 `json:"mitra_id"`
	MitraName   string `json:"mitra_name"`
	MitraLogo   string `json:"mitra_logo"`
	ReferalCode string `json:"mitra_user_referal_code"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Sex         string `json:"gender"`
	BirthPlace  string `json:"birth_place"`
	BirthDate   string `json:"birth_date"`
	Photo       string `json:"photo"`
	Status      uint64 `json:"status"`
}

type ResponseMitraProfile struct {
	Status  uint64                   `json:"status"`
	Message string                   `json:"message"`
	Data    ResponseMitraProfileData `json:"data"`
}

type ResponseMitraUpdateProfileData struct {
	Mitra     uint64 `json:"mitra_id"`
	UserMitra uint64 `json:"mitra_user_id"`
}

type ResponseUpdateMitraProfile struct {
	Status  uint64                         `json:"status"`
	Message string                         `json:"message"`
	Data    ResponseMitraUpdateProfileData `json:"data"`
}

func (u *MitraUser) BeforeSaveMitra() error {
	hashedPassword, err := HashMitraUser(u.MitraUserPasswordWeb)
	if err != nil {
		return err
	}
	u.MitraUserPasswordWeb = string(hashedPassword)
	return nil
}

func (u *MitraUser) BeforeSaveMitraMobile() error {
	hashedPassword, err := HashMitraUser(u.MitraUserPasswordMobile)
	if err != nil {
		return err
	}
	u.MitraUserPasswordMobile = string(hashedPassword)
	return nil
}

func (u *MitraUser) BeforeUpdateMitraMobile() error {
	hashedPassword, err := HashMitraUser(u.MitraUserPasswordNew)
	if err != nil {
		return err
	}
	u.MitraUserPasswordNew = string(hashedPassword)
	return nil
}

func (u *MitraUser) BeforeUpdateMitraWeb() error {
	hashedPassword, err := HashMitraUser(u.MitraUserPasswordNew)
	if err != nil {
		return err
	}
	u.MitraUserPasswordNew = string(hashedPassword)
	return nil
}

func (u *MitraUser) BeforeSaveMitraPin() error {
	hashedPassword, err := HashMitraUser(u.MitraUserPINNew)
	if err != nil {
		return err
	}
	u.MitraUserPINNew = string(hashedPassword)
	return nil
}

func MitraHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func MitraVerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func MitraVerifyPin(hashedPin, pin string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPin), []byte(pin))
}

func (u *MitraUser) MitraBeforeSave() error {
	hashedPassword, err := MitraHash(u.MitraUserPasswordWeb)
	if err != nil {
		return err
	}
	u.MitraUserPasswordWeb = string(hashedPassword)
	return nil
}

func (u *MitraUser) MitraBeforeSaveMobile() error {

	hashedPassword, err := MitraHash(u.MitraUserPasswordMobile)
	if err != nil {
		return err
	}
	u.MitraUserPasswordMobile = string(hashedPassword)
	return nil
}

func (p *MitraUser) FindAllMitraUserReferalCode(db *gorm.DB, referal string) (*MitraUser, error) {
	var err error
	err = db.Debug().Model(&MitraUser{}).Where("mitra_user_referal_code = ? AND mitra_status = ?", referal, 1).Take(&p).Error
	if err != nil {
		return &MitraUser{}, err
	}
	return p, nil
}

func (p *MitraUser) FindMitraUserStatusByIDUpdate(db *gorm.DB, pid uint64, status uint64) (*MitraUserData, error) {
	var err error
	mitrauser := MitraUserData{}
	err = db.Debug().Model(&MitraUserData{}).
		Select(`m_mitra_user.mitra_user_id,
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
				m_mitra_user.mitra_user_referal_code,
				m_mitra_user.mitra_user_email,
				m_mitra_user.mitra_username,
				m_mitra_user.mitra_user_phone_code,
				m_mitra_user.mitra_user_phone,
				m_mitra_user.mitra_user_imei,
				m_mitra_user.mitra_user_ip_address,
				m_mitra_user.mitra_user_is_ao,
				m_mitra_user.mitra_user_is_credit,
				m_mitra_user.mitra_user_status,
				m_mitra_user.mitra_user_is_new,
				m_mitra_user.mitra_user_is_req_chpw,
				m_mitra_user.mitra_user_is_locked,
				m_mitra_user.mitra_user_fail_attempt,
				m_mitra_user.mitra_user_locked_reason,
				m_mitra_user.mitra_user_created_by,
				m_mitra_user.mitra_user_updated_by,
				m_mitra_user.mitra_user_deleted_by,
				m_mitra_user.mitra_user_created_at,
				m_mitra_user.mitra_user_updated_at,
				m_mitra_user.mitra_user_deleted_at`).
		Joins("JOIN m_mitra_employee on m_mitra_user.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra_group_access on m_mitra_user.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Joins("JOIN m_mitra_group_product_limit on m_mitra_user.mitra_group_prod_lim_id=m_mitra_group_product_limit.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_ao_coverage on m_mitra_user.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Where("mitra_user_id = ? AND mitra_user_status = ?", pid, status).
		Take(&mitrauser).Error
	if err != nil {
		return &MitraUserData{}, err
	}
	return &mitrauser, nil
}

func (p *MitraUser) LoginMitraUserByPhone(db *gorm.DB, phone string) (*MitraUserData, error) {
	var err error
	mitrauser := MitraUserData{}
	err = db.Debug().Model(&MitraUserData{}).
		Select(`m_mitra_user.mitra_user_id,
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
				m_mitra_user.mitra_user_referal_code,
				m_mitra_user.mitra_user_email,
				m_mitra_user.mitra_username,
				m_mitra_user.mitra_user_pin,
				m_mitra_user.mitra_user_phone_code,
				m_mitra_user.mitra_user_phone,
				m_mitra_user.mitra_user_imei,
				m_mitra_user.mitra_user_ip_address,
				m_mitra_user.mitra_user_is_ao,
				m_mitra_user.mitra_user_is_credit,
				m_mitra_user.mitra_user_status,
				m_mitra_user.mitra_user_is_new,
				m_mitra_user.mitra_user_is_req_chpw,
				m_mitra_user.mitra_user_is_locked,
				m_mitra_user.mitra_user_fail_attempt,
				m_mitra_user.mitra_user_locked_reason,
				m_mitra_user.mitra_user_created_by,
				m_mitra_user.mitra_user_updated_by,
				m_mitra_user.mitra_user_deleted_by,
				m_mitra_user.mitra_user_created_at,
				m_mitra_user.mitra_user_updated_at,
				m_mitra_user.mitra_user_deleted_at`).
		Joins("JOIN m_mitra_employee on m_mitra_user.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra_group_access on m_mitra_user.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Joins("JOIN m_mitra_group_product_limit on m_mitra_user.mitra_group_prod_lim_id=m_mitra_group_product_limit.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_ao_coverage on m_mitra_user.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Where("mitra_user_phone = ?", phone).
		Take(&mitrauser).Error
	if err != nil {
		return &MitraUserData{}, err
	}
	return &mitrauser, nil
}

func (p *MitraUser) CheckMitraUserByPhone(db *gorm.DB, phone string) (*MitraUserData, error) {
	var err error
	mitrauser := MitraUserData{}
	err = db.Debug().Model(&MitraUserData{}).
		Select(`m_mitra_user.mitra_user_id,
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
				m_mitra_user.mitra_user_referal_code,
				m_mitra_user.mitra_user_email,
				m_mitra_user.mitra_username,
				m_mitra_user.mitra_user_phone_code,
				m_mitra_user.mitra_user_phone,
				m_mitra_user.mitra_user_otp,
				m_mitra_user.mitra_user_otp_secret,
				m_mitra_user.mitra_user_imei,
				m_mitra_user.mitra_user_ip_address,
				m_mitra_user.mitra_user_is_ao,
				m_mitra_user.mitra_user_is_credit,
				m_mitra_user.mitra_user_status,
				m_mitra_user.mitra_user_is_new,
				m_mitra_user.mitra_user_is_req_chpw,
				m_mitra_user.mitra_user_is_locked,
				m_mitra_user.mitra_user_fail_attempt,
				m_mitra_user.mitra_user_locked_reason,
				m_mitra_user.mitra_user_created_by,
				m_mitra_user.mitra_user_updated_by,
				m_mitra_user.mitra_user_deleted_by,
				m_mitra_user.mitra_user_created_at,
				m_mitra_user.mitra_user_updated_at,
				m_mitra_user.mitra_user_deleted_at`).
		Joins("JOIN m_mitra_employee on m_mitra_user.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra_group_access on m_mitra_user.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Joins("JOIN m_mitra_group_product_limit on m_mitra_user.mitra_group_prod_lim_id=m_mitra_group_product_limit.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_ao_coverage on m_mitra_user.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Where("mitra_user_phone = ?", phone).
		Take(&mitrauser).Error
	if err != nil {
		return &MitraUserData{}, err
	}
	return &mitrauser, nil
}

func (p *MitraUser) CheckMitraUserByUsername(db *gorm.DB, username string, mitra string) (*MitraUserData, error) {
	var err error
	mitrauser := MitraUserData{}
	err = db.Debug().Model(&MitraUserData{}).
		Select(`m_mitra_user.mitra_user_id,
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
				m_mitra_user.mitra_user_referal_code,
				m_mitra_user.mitra_user_email,
				m_mitra_user.mitra_username,
				m_mitra_user.mitra_user_phone_code,
				m_mitra_user.mitra_user_phone,
				m_mitra_user.mitra_user_imei,
				m_mitra_user.mitra_user_ip_address,
				m_mitra_user.mitra_user_is_ao,
				m_mitra_user.mitra_user_is_credit,
				m_mitra_user.mitra_user_status,
				m_mitra_user.mitra_user_is_new,
				m_mitra_user.mitra_user_is_req_chpw,
				m_mitra_user.mitra_user_is_locked,
				m_mitra_user.mitra_user_fail_attempt,
				m_mitra_user.mitra_user_locked_reason,
				m_mitra_user.mitra_user_created_by,
				m_mitra_user.mitra_user_updated_by,
				m_mitra_user.mitra_user_deleted_by,
				m_mitra_user.mitra_user_created_at,
				m_mitra_user.mitra_user_updated_at,
				m_mitra_user.mitra_user_deleted_at`).
		Joins("JOIN m_mitra_employee on m_mitra_user.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra_group_access on m_mitra_user.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Joins("JOIN m_mitra_group_product_limit on m_mitra_user.mitra_group_prod_lim_id=m_mitra_group_product_limit.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_ao_coverage on m_mitra_user.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Where("mitra_username = ? AND mitra_code = ?", username, mitra).
		Take(&mitrauser).Error
	if err != nil {
		return &MitraUserData{}, err
	}
	return &mitrauser, nil
}

func (p *MitraUser) LoginMitraUserByPin(db *gorm.DB, phone, pin string) (*MitraUserData, error) {
	var err error
	mitrauser := MitraUserData{}
	err = db.Debug().Model(&MitraUserData{}).
		Select(`m_mitra_user.mitra_user_id,
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
				m_mitra_user.mitra_user_referal_code,
				m_mitra_user.mitra_user_email,
				m_mitra_user.mitra_username,
				m_mitra_user.mitra_user_phone_code,
				m_mitra_user.mitra_user_phone,
				m_mitra_user.mitra_user_imei,
				m_mitra_user.mitra_user_ip_address,
				m_mitra_user.mitra_user_is_ao,
				m_mitra_user.mitra_user_is_credit,
				m_mitra_user.mitra_user_status,
				m_mitra_user.mitra_user_is_new,
				m_mitra_user.mitra_user_is_req_chpw,
				m_mitra_user.mitra_user_is_locked,
				m_mitra_user.mitra_user_fail_attempt,
				m_mitra_user.mitra_user_locked_reason,
				m_mitra_user.mitra_user_created_by,
				m_mitra_user.mitra_user_updated_by,
				m_mitra_user.mitra_user_deleted_by,
				m_mitra_user.mitra_user_created_at,
				m_mitra_user.mitra_user_updated_at,
				m_mitra_user.mitra_user_deleted_at`).
		Joins("JOIN m_mitra_employee on m_mitra_user.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra_group_access on m_mitra_user.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Joins("JOIN m_mitra_group_product_limit on m_mitra_user.mitra_group_prod_lim_id=m_mitra_group_product_limit.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_ao_coverage on m_mitra_user.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Where("mitra_user_phone = ? AND mitra_user_pin = ? ", phone, pin).
		Take(&mitrauser).Error
	if err != nil {
		return &MitraUserData{}, err
	}
	return &mitrauser, nil
}

func (p *MitraUser) LoginMitraUserByPswdMobile(db *gorm.DB, phone, pswd string) (*MitraUserData, error) {
	var err error
	mitrauser := MitraUserData{}
	err = db.Debug().Model(&MitraUserData{}).
		Select(`m_mitra_user.mitra_user_id,
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
				m_mitra_user.mitra_user_referal_code,
				m_mitra_user.mitra_user_email,
				m_mitra_user.mitra_username,
				m_mitra_user.mitra_user_phone_code,
				m_mitra_user.mitra_user_phone,
				m_mitra_user.mitra_user_imei,
				m_mitra_user.mitra_user_ip_address,
				m_mitra_user.mitra_user_is_ao,
				m_mitra_user.mitra_user_is_credit,
				m_mitra_user.mitra_user_status,
				m_mitra_user.mitra_user_is_new,
				m_mitra_user.mitra_user_is_req_chpw,
				m_mitra_user.mitra_user_is_locked,
				m_mitra_user.mitra_user_fail_attempt,
				m_mitra_user.mitra_user_locked_reason,
				m_mitra_user.mitra_user_created_by,
				m_mitra_user.mitra_user_updated_by,
				m_mitra_user.mitra_user_deleted_by,
				m_mitra_user.mitra_user_created_at,
				m_mitra_user.mitra_user_updated_at,
				m_mitra_user.mitra_user_deleted_at`).
		Joins("JOIN m_mitra_employee on m_mitra_user.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra_group_access on m_mitra_user.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Joins("JOIN m_mitra_group_product_limit on m_mitra_user.mitra_group_prod_lim_id=m_mitra_group_product_limit.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_ao_coverage on m_mitra_user.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Where("mitra_user_phone = ? AND mitra_user_password_mobile = ? ", phone, pswd).
		Take(&mitrauser).Error
	if err != nil {
		return &MitraUserData{}, err
	}
	return &mitrauser, nil
}

func (p *MitraUser) FindMitraUserStatusByUsername(db *gorm.DB, username string) (*MitraUserData, error) {
	var err error
	mitrauser := MitraUserData{}
	err = db.Debug().Model(&MitraUserData{}).
		Select(`m_mitra_user.mitra_user_id,
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
				m_mitra_user.mitra_user_referal_code,
				m_mitra_user.mitra_user_email,
				m_mitra_user.mitra_username,
				m_mitra_user.mitra_user_phone_code,
				m_mitra_user.mitra_user_phone,
				m_mitra_user.mitra_user_imei,
				m_mitra_user.mitra_user_ip_address,
				m_mitra_user.mitra_user_is_ao,
				m_mitra_user.mitra_user_is_credit,
				m_mitra_user.mitra_user_status,
				m_mitra_user.mitra_user_is_new,
				m_mitra_user.mitra_user_is_req_chpw,
				m_mitra_user.mitra_user_is_locked,
				m_mitra_user.mitra_user_fail_attempt,
				m_mitra_user.mitra_user_locked_reason,
				m_mitra_user.mitra_user_created_by,
				m_mitra_user.mitra_user_updated_by,
				m_mitra_user.mitra_user_deleted_by,
				m_mitra_user.mitra_user_created_at,
				m_mitra_user.mitra_user_updated_at,
				m_mitra_user.mitra_user_deleted_at`).
		Joins("JOIN m_mitra_employee on m_mitra_user.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra_group_access on m_mitra_user.mitra_group_accs_id=m_mitra_group_access.mitra_group_accs_id").
		Joins("JOIN m_mitra_group_product_limit on m_mitra_user.mitra_group_prod_lim_id=m_mitra_group_product_limit.mitra_group_prod_lim_id").
		Joins("JOIN m_mitra_ao_coverage on m_mitra_user.mitra_ao_cov_id=m_mitra_ao_coverage.mitra_ao_cov_id").
		Where("mitra_username = ? AND mitra_user_status = ?", username, 1).
		Take(&mitrauser).Error
	if err != nil {
		return &MitraUserData{}, err
	}
	return &mitrauser, nil
}

func (p *MitraUser) CountMitraUserStatusByUsername(db *gorm.DB, username string) (*MitraUserData, error) {
	var err error
	mitrauser := MitraUserData{}
	err = db.Debug().Model(&MitraUserData{}).
		Select(`m_mitra_user.mitra_username,
				m_mitra_user.mitra_user_status`).
		Where("mitra_username = ? AND mitra_user_status = ?", username, 1).
		Count(&mitrauser.MitraUserCount).
		Take(&mitrauser).Error
	if err != nil {
		return &MitraUserData{}, err
	}
	return &mitrauser, nil
}

func (u *MitraUser) UpdateMitraUserOTP(db *gorm.DB, uid uint64) (*MitraUser, error) {

	// To hash the password
	err := u.BeforeSaveMitra()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&MitraUser{}).Where("mitra_user_id = ?", uid).Take(&MitraUser{}).UpdateColumns(
		map[string]interface{}{
			"mitra_user_otp":        u.MitraUserOTP,
			"mitra_user_secret":     u.MitraUserSecret,
			"mitra_user_updated_by": u.MitraUserUpdatedBy,
			"mitra_user_otp_secret": time.Now(),
			"mitra_updated_at":      time.Now(),
		},
	)
	if db.Error != nil {
		return &MitraUser{}, db.Error
	}
	// This is the display the updated mitrauser
	err = db.Debug().Model(&MitraUser{}).Where("mitra_user_id = ?", uid).Take(&u).Error
	if err != nil {
		return &MitraUser{}, err
	}
	return u, nil
}

func (u *MitraUser) SetPinMitraUser(db *gorm.DB, uid uint64) (*MitraUser, error) {

	// To hash the password
	err := u.BeforeSaveMitra()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&MitraUser{}).Where("mitra_user_phone = ?", uid).Take(&MitraUser{}).UpdateColumns(
		map[string]interface{}{
			"mitra_user_pin":   u.MitraUserPIN,
			"mitra_updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &MitraUser{}, db.Error
	}
	// This is the display the updated mitrauser
	err = db.Debug().Model(&MitraUser{}).Where("mitra_user_phone = ?", uid).Take(&u).Error
	if err != nil {
		return &MitraUser{}, err
	}
	return u, nil
}

func (u *MitraUser) UpdatePinMitraUser(db *gorm.DB, uid uint64) (*MitraUser, error) {

	// To hash the password
	err := u.BeforeSaveMitraPin()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&MitraUser{}).Where("mitra_user_id = ?", uid).Take(&MitraUser{}).UpdateColumns(
		map[string]interface{}{
			"mitra_user_pin":        u.MitraUserPINNew,
			"mitra_user_updated_by": u.MitraUserUpdatedBy,
			"mitra_updated_at":      time.Now(),
		},
	)
	if db.Error != nil {
		return &MitraUser{}, db.Error
	}
	// This is the display the updated mitrauser
	err = db.Debug().Model(&MitraUser{}).Where("mitra_user_id = ?", uid).Take(&u).Error
	if err != nil {
		return &MitraUser{}, err
	}
	return u, nil
}

func (u *MitraUser) UpdatePswdMobileMitraUser(db *gorm.DB, uid string) (*MitraUser, error) {

	// To hash the password
	err := u.BeforeUpdateMitraMobile()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&MitraUser{}).Where("mitra_user_phone = ?", uid).Take(&MitraUser{}).UpdateColumns(
		map[string]interface{}{
			"mitra_user_password_mobile": u.MitraUserPasswordNew,
			"mitra_user_updated_by":      u.MitraUserUpdatedBy,
			"mitra_updated_at":           time.Now(),
		},
	)
	if db.Error != nil {
		return &MitraUser{}, db.Error
	}
	// This is the display the updated mitrauser
	err = db.Debug().Model(&MitraUser{}).Where("mitra_user_phone = ?", uid).Take(&u).Error
	if err != nil {
		return &MitraUser{}, err
	}
	return u, nil
}

func (u *MitraUser) UpdatePswdWebMitraUser(db *gorm.DB, uid string) (*MitraUser, error) {

	// To hash the password
	err := u.BeforeUpdateMitraMobile()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&MitraUser{}).Where("mitra_user_phone = ?", uid).Take(&MitraUser{}).UpdateColumns(
		map[string]interface{}{
			"mitra_user_password_web": u.MitraUserPasswordKonfirm,
			"mitra_user_updated_by":   u.MitraUserUpdatedBy,
			"mitra_updated_at":        time.Now(),
		},
	)
	if db.Error != nil {
		return &MitraUser{}, db.Error
	}
	// This is the display the updated mitrauser
	err = db.Debug().Model(&MitraUser{}).Where("mitra_user_phone = ?", uid).Take(&u).Error
	if err != nil {
		return &MitraUser{}, err
	}
	return u, nil
}

func (u *MitraUser) UpdateProfileMitra(db *gorm.DB, uid uint64) (*MitraUser, error) {

	// To hash the password
	err := u.BeforeSaveMitra()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&MitraUser{}).Where("mitra_user_id = ?", uid).Take(&MitraUser{}).UpdateColumns(
		map[string]interface{}{
			"mitra_user_phone":      u.MitraUserPhone,
			"mitra_user_email":      u.MitraUserEmail,
			"mitra_user_updated_by": u.MitraUserUpdatedBy,
			"mitra_updated_at":      time.Now(),
		},
	)
	if db.Error != nil {
		return &MitraUser{}, db.Error
	}
	// This is the display the updated mitrauser
	err = db.Debug().Model(&MitraUser{}).Where("mitra_user_id = ?", uid).Error
	if err != nil {
		return &MitraUser{}, err
	}
	return u, nil
}

func (p *MitraUser) UpdateMitraUserSecret(db *gorm.DB, pid uint64) (*MitraUser, error) {

	var err error
	db = db.Debug().Model(&MitraUser{}).Where("mitra_user_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_user_secret":     p.MitraUserSecret,
			"mitra_user_updated_by": p.MitraUserUpdatedBy,
			"mitra_user_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraUser{}).Where("mitra_user_id = ?", pid).Error
	if err != nil {
		return &MitraUser{}, err
	}
	return p, nil
}

func (p *MitraUser) UpdateMitraUserOTPBlank(db *gorm.DB, pid uint64) (*MitraUser, error) {

	var err error
	db = db.Debug().Model(&MitraUser{}).Where("mitra_user_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_user_otp":        nil,
			"mitra_user_otp_secret": nil,
			"mitra_user_updated_by": p.MitraUserUpdatedBy,
			"mitra_user_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraUser{}).Where("mitra_user_id = ?", pid).Error
	if err != nil {
		return &MitraUser{}, err
	}
	return p, nil
}

func (p *MitraUser) UpdateMitraUserLastLogin(db *gorm.DB, username string) (*MitraUser, error) {

	var err error
	db = db.Debug().Model(&MitraUser{}).Where("mitra_username = ?", username).UpdateColumns(
		map[string]interface{}{
			"mitra_user_last_login": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraUser{}).Where("mitra_username = ?", username).Error
	if err != nil {
		return &MitraUser{}, err
	}
	return p, nil
}
