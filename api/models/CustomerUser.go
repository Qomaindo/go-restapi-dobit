package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type CustomerUser struct {
	CustomerUserID                      uint64     `gorm:"column:cust_user_id;primary_key;" json:"cust_user_id"`
	CustomerID                          uint64     `gorm:"column:cust_id;" json:"cust_id"`
	CustomerSubTypeID                   uint64     `gorm:"column:cust_sub_type_id;" json:"cust_sub_type_id"`
	CustomerUserReferalCode             string     `gorm:"column:cust_user_referal_code;size:25" json:"cust_user_referal_code"`
	CustomerUserReferalCodeFrom         string     `gorm:"column:cust_user_referal_code_from;size:25" json:"cust_user_referal_code_from"`
	CustomerUserEmail                   string     `gorm:"column:cust_user_email;size:255" json:"cust_user_email" validate:"required,email"`
	CustomerUsername                    string     `gorm:"column:cust_username;size:125" json:"cust_username"`
	CustomerUserPassword                string     `gorm:"column:cust_user_password;size:255" json:"cust_user_password"`
	CustomerUserPhoneCode               string     `gorm:"column:cust_user_phone_code;size:6" json:"cust_user_phone_code"`
	CustomerUserPhone                   string     `gorm:"column:cust_user_phone;size:20" json:"cust_user_phone" validate:"length(11|13)"`
	CustomerUserOTP                     string     `gorm:"column:cust_user_otp;size:6" json:"cust_user_otp"`
	CustomerUserOTPSecret               *time.Time `gorm:"column:cust_user_otp_secret;default:CURRENT_TIMESTAMP" json:"cust_user_otp_secret"`
	CustomerUserSecret                  uint64     `gorm:"column:cust_user_secret;size:2" json:"cust_user_secret"`
	CustomerUserPin                     string     `gorm:"column:cust_user_pin;size:6" json:"cust_user_pin"`
	CustomerUserLastLogin               *time.Time `gorm:"column:cust_user_last_login;default:CURRENT_TIMESTAMP" json:"cust_user_last_login"`
	CustomerUserImei                    string     `gorm:"column:cust_user_imei;size:125" json:"cust_user_imei"`
	CustomerUserIPAddress               string     `gorm:"column:cust_user_ip_address;size:125" json:"cust_user_ip_address"`
	CustomerUserEmailVerifiedAt         *time.Time `gorm:"column:cust_user_email_verified_at;default:CURRENT_TIMESTAMP" json:"cust_user_email_verified_at"`
	CustomerUserStatus                  uint64     `gorm:"column:cust_user_status;size:1" json:"cust_user_status"`
	CustomerUserIsNew                   bool       `gorm:"column:cust_user_is_new" json:"cust_user_is_new"`
	CustomerUserIsRequestChangePassword bool       `gorm:"column:cust_user_is_req_chpw" json:"cust_user_is_req_chpw"`
	CustomerUserIsLocked                bool       `gorm:"column:cust_user_is_locked" json:"cust_user_is_locked"`
	CustomerUserFailAttempt             uint64     `gorm:"column:cust_user_fail_attempt" json:"cust_user_fail_attempt"`
	CustomerUserLockedReason            string     `gorm:"column:cust_user_locked_reason;size:255" json:"cust_user_locked_reason"`
	CustomerUserCreatedBy               string     `gorm:"column:cust_user_created_by;size:125" json:"cust_user_created_by"`
	CustomerUserCreatedAt               time.Time  `gorm:"column:cust_user_created_at;default:CURRENT_TIMESTAMP" json:"cust_user_created_at"`
	CustomerUserUpdatedBy               string     `gorm:"column:cust_user_updated_by;size:125" json:"cust_user_updated_by"`
	CustomerUserUpdatedAt               *time.Time `gorm:"column:cust_user_updated_at;default:CURRENT_TIMESTAMP" json:"cust_user_updated_at"`
	CustomerUserDeletedBy               string     `gorm:"column:cust_user_deleted_by;size:125" json:"cust_user_deleted_by"`
	CustomerUserDeletedAt               *time.Time `gorm:"column:cust_user_deleted_at;default:CURRENT_TIMESTAMP" json:"cust_user_deleted_at"`

	CustomerLoginStaticToken string    `gorm:"-" json:"cust_login_static_token"`
	CustomerUserPinOld       string    `gorm:"-" json:"cust_user_pin_old"`
	CustomerUserPinNew       string    `gorm:"-" json:"cust_user_pin_new"`
	CustomerUserOTPNew       string    `gorm:"-" json:"cust_user_otp_new"`
	CustomerUserOTPSecretNew time.Time `gorm:"-"`
	Msg_id                   string    `gorm:"-"`
	Trx_id                   string    `gorm:"-"`
	Status                   uint64    `gorm:"-"`
	Secret                   string    `gorm:"-"`
	Segment                  uint64    `gorm:"-"`

	CustomerUserPswdOld string `gorm:"-" json:"cust_user_pswd_old"`
	CustomerUserPswdNew string `gorm:"-" json:"cust_user_pswd_new"`
}

type CustomerUserData struct {
	CustomerUserID                      uint64     `gorm:"column:cust_user_id" json:"cust_user_id"`
	CustomerID                          uint64     `gorm:"column:cust_id;" json:"cust_id"`
	CustomerCode                        string     `gorm:"column:cust_code;size:25" json:"cust_code"`
	CustomerKTP                         string     `gorm:"column:cust_ktp;size:255" json:"cust_ktp"`
	CustomerPassport                    string     `gorm:"column:cust_passport;size:125" json:"cust_passport"`
	CustomerNPWP                        string     `gorm:"column:cust_npwp;size:255" json:"cust_npwp"`
	CustomerName                        string     `gorm:"column:cust_name;size:255" json:"cust_name"`
	CustomerSex                         string     `gorm:"column:cust_sex;size:6" json:"cust_sex"`
	CustomerBirthDate                   string     `gorm:"column:cust_birth_date;size:25" json:"cust_birth_date"`
	CustomerBirthPlace                  string     `gorm:"column:cust_birth_place;size:255" json:"cust_birth_place"`
	CustomerAge                         string     `gorm:"column:cust_age;size:3" json:"cust_age"`
	CustomerPhoneWork                   string     `gorm:"column:cust_phone_work;size:20" json:"cust_phone_work"`
	CustomerEmailWork                   string     `gorm:"column:cust_email_work;size:255" json:"cust_email_work"`
	CustomerLastEducation               string     `gorm:"column:cust_last_education;size:10" json:"cust_last_education"`
	CustomerMaritalStatus               string     `gorm:"column:cust_marital_status;size:125" json:"cust_marital_status"`
	CustomerImage                       string     `gorm:"column:cust_image;size:255" json:"cust_image"`
	CustomerUserReferalCode             string     `gorm:"column:cust_user_referal_code" json:"cust_user_referal_code"`
	CustomerUserReferalCodeFrom         string     `gorm:"column:cust_user_referal_code_from;size:25" json:"cust_user_referal_code_from"`
	CustomerUserEmail                   string     `gorm:"column:cust_user_email;size:255" json:"cust_user_email"`
	CustomerUsername                    string     `gorm:"column:cust_username;size:125" json:"cust_username"`
	CustomerUserPassword                string     `gorm:"column:cust_user_password;size:255" json:"cust_user_password"`
	CustomerUserPhoneCode               string     `gorm:"column:cust_user_phone_code;size:6" json:"cust_user_phone_code"`
	CustomerUserPhone                   string     `gorm:"column:cust_user_phone;size:20" json:"cust_user_phone"`
	CustomerUserOTP                     string     `gorm:"column:cust_user_otp;size:6" json:"cust_user_otp"`
	CustomerUserOTPSecret               *time.Time `gorm:"column:cust_user_otp_secret;default:CURRENT_TIMESTAMP" json:"cust_user_otp_secret"`
	CustomerUserSecret                  uint64     `gorm:"column:cust_user_secret;size:2" json:"cust_user_secret"`
	CustomerUserPin                     string     `gorm:"column:cust_user_pin;size:6" json:"cust_user_pin"`
	CustomerUserLastLogin               *time.Time `gorm:"column:cust_user_last_login;default:CURRENT_TIMESTAMP" json:"cust_user_last_login"`
	CustomerUserImei                    string     `gorm:"column:cust_user_imei;size:125" json:"cust_user_imei"`
	CustomerUserIPAddress               string     `gorm:"column:cust_user_ip_address;size:125" json:"cust_user_ip_address"`
	CustomerUserEmailVerifiedAt         *time.Time `gorm:"column:cust_user_email_verified_at;default:CURRENT_TIMESTAMP" json:"cust_user_email_verified_at"`
	CustomerUserStatus                  uint64     `gorm:"column:cust_user_status;size:1" json:"cust_user_status"`
	CustomerUserIsNew                   bool       `gorm:"column:cust_user_is_new" json:"cust_user_is_new"`
	CustomerUserIsRequestChangePassword bool       `gorm:"column:cust_user_is_req_chpw" json:"cust_user_is_req_chpw"`
	CustomerUserIsLocked                bool       `gorm:"column:cust_user_is_locked" json:"cust_user_is_locked"`
	CustomerUserFailAttempt             uint64     `gorm:"column:cust_user_fail_attempt" json:"cust_user_fail_attempt"`
	CustomerUserLockedReason            string     `gorm:"column:cust_user_locked_reason;size:255" json:"cust_user_locked_reason;size:255"`
	CustomerUserCreatedBy               string     `gorm:"column:cust_user_created_by;size:125" json:"cust_user_created_by"`
	CustomerUserCreatedAt               time.Time  `gorm:"column:cust_user_created_at;default:CURRENT_TIMESTAMP" json:"cust_user_created_at"`
	CustomerUserUpdatedBy               string     `gorm:"column:cust_user_updated_by;size:125" json:"cust_user_updated_by"`
	CustomerUserUpdatedAt               *time.Time `gorm:"column:cust_user_updated_at;default:CURRENT_TIMESTAMP" json:"cust_user_updated_at"`
	CustomerUserDeletedBy               string     `gorm:"column:cust_user_deleted_by;size:125" json:"cust_user_deleted_by"`
	CustomerUserDeletedAt               *time.Time `gorm:"column:cust_user_deleted_at;default:CURRENT_TIMESTAMP" json:"cust_user_deleted_at"`
	CustomerSubTypeID                   uint64     `gorm:"column:cust_sub_type_id;" json:"cust_sub_type_id"`
	CustomerSubTypeName                 string     `gorm:"column:cust_sub_type_name" json:"cust_sub_type_name"`
	CustomerTypeID                      uint64     `gorm:"column:cust_type_id;" json:"cust_type_id"`
	CustomerTypeName                    string     `gorm:"column:cust_type_name" json:"cust_type_name"`
	CustomerLoginStaticToken            string     `gorm:"-" json:"cust_login_static_token"`
}

type DetailMobileUser struct {
	CustomerUserPhone  string    `gorm:"column:cust_user_phone;size:20" json:"cust_user_phone"`
	CustomerUserID     string    `gorm:"column:cust_user_id" json:"cust_user_id"`
	CustomerKTP        string    `gorm:"column:cust_ktp;size:255" json:"cust_ktp"`
	CustomerName       string    `gorm:"column:cust_name;size:255" json:"cust_name"`
	CustomerSex        string    `gorm:"column:cust_sex;size:6" json:"cust_sex"`
	CustomerBirthDate  string    `gorm:"column:cust_birth_date;size:25" json:"cust_birth_date"`
	CustomerBirthPlace string    `gorm:"column:cust_birth_place;size:255" json:"cust_birth_place"`
	CustomerAge        string    `gorm:"column:cust_age;size:3" json:"cust_age"`
	CustomerUserEmail  string    `gorm:"column:cust_user_email;size:255" json:"cust_user_email"`
	CustomerUserStatus uint64    `gorm:"column:cust_user_status;size:1" json:"cust_user_status"`
	CreatedUserBy      string    `gorm:"column:cust_user_created_by;size:125" json:"cust_user_created_by"`
	CreatedUserAt      time.Time `gorm:"column:cust_user_created_at;default:CURRENT_TIMESTAMP" json:"cust_user_created_at"`

	CountPhoneCustomer int `gorm:"-" json:"count_phone_customer"`
	CountEmailCustomer int `gorm:"-" json:"cust_user_email"`
}

type MobileUserData struct {
	CustomerUserID     string `gorm:"column:cust_user_id" json:"cust_user_id"`
	CustomerUserPhone  string `gorm:"column:cust_user_phone;size:20" json:"cust_user_phone"`
	CustomerUserEmail  string `gorm:"column:cust_user_email;size:255" json:"cust_user_email"`
	CustomerName       string `gorm:"column:cust_name;size:125" json:"cust_name"`
	CustomerUserStatus uint64 `gorm:"column:cust_user_status;size:1" json:"cust_user_status"`
}

type ResponseDetailMobileUser struct {
	Status  uint64            `json:"status"`
	Message string            `json:"message"`
	Data    *DetailMobileUser `json:"data"`
}

type ResponseMobileUser struct {
	Status  uint64            `json:"status"`
	Message string            `json:"message"`
	Data    *[]MobileUserData `json:"data"`
}

type ResponseCustomerUsers struct {
	Status  uint64              `json:"status"`
	Message string              `json:"message"`
	Data    *[]CustomerUserData `json:"data"`
}

type ResponseCustomerUser struct {
	Status  uint64            `json:"status"`
	Message string            `json:"message"`
	Data    *CustomerUserData `json:"data"`
}

type ResponseCustomerUserIU struct {
	Status  uint64        `json:"status"`
	Message string        `json:"message"`
	Data    *CustomerUser `json:"data"`
}

type ResponseCustomerUserDel struct {
	Status  uint64 `json:"status"`
	Message string `json:"message"`
}

type ResponseCustomerRegister struct {
	Status  uint64 `json:"status"`
	Message string `json:"message"`
}

type ResponseCustomerPhoneData struct {
	Status uint64 `json:"status"`
	Name   string `json:"name"`
	Photo  string `json:"photo"`
	Token  string `json:"token"`
}

type ResponseCustomerPhone struct {
	Status  int                       `json:"status"`
	Message string                    `json:"message"`
	Data    ResponseCustomerPhoneData `json:"data"`
}

type ResponseCustomerTokenData struct {
	Token  string `json:"token"`
	Status string `json:"status"`
	Name   string `json:"name"`
	Photo  string `json:"photo"`
}

type ResponseCustomerSendData struct {
	Status       uint64 `json:"status"`
	Name         string `json:"name"`
	SMSStatus    string `json:"sms_status"`
	UpdateStatus uint64 `json:"update_status"`
}

type ResponseCustomerUserSend struct {
	Status  uint64                   `json:"status"`
	Message string                   `json:"message"`
	Data    ResponseCustomerSendData `json:"data"`
}

type ResponseCustomerVerifyData struct {
	Token     string `json:"token"`
	Status    uint64 `json:"status"`
	Name      string `json:"name"`
	SMSStatus string `json:"sms_status"`
}

type ResponseCustomerUserVerify struct {
	Status  uint64                     `json:"status"`
	Message string                     `json:"message"`
	Data    ResponseCustomerVerifyData `json:"data"`
}

type ResponseCustomerUserToken struct {
	Status  uint64                    `json:"status"`
	Message string                    `json:"message"`
	Data    ResponseCustomerTokenData `json:"data"`
}

type ResponseCustomerOTPData struct {
	Status    int    `json:"status"`
	Condition string `json:"condition"`
	Token     string `json:"token"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Photo     string `json:"photo"`
}

type ResponseCustomerOTP struct {
	Status  uint64                    `json:"status"`
	Message string                    `json:"message"`
	Data    ResponseCustomerPhoneData `json:"data"`
}

type ResponseCustomerPinData struct {
	Status string `json:"status"`
	Name   string `json:"name"`
	Photo  string `json:"photo"`
}

type ResponseCustomerPin struct {
	Status  uint64                    `json:"status"`
	Message string                    `json:"message"`
	Data    ResponseCustomerPhoneData `json:"data"`
}

type ResponseCustomerPswd struct {
	Status  uint64                    `json:"status"`
	Message string                    `json:"message"`
	Data    ResponseCustomerPhoneData `json:"data"`
}

type ResponseCustomerCodeOTP struct {
	TimeSend time.Time `json:"time_send"`
	Message  string    `json:"message"`
	Code     string    `json:"code_otp"`
}

type ResponseCustomerDashboardData struct {
	Name         string    `json:"name"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	RegisterDate time.Time `json:"register_date"`
	Status       uint64    `json:"status"`
}

type ResponseCustomerDashboard struct {
	Status  uint64                        `json:"status"`
	Message string                        `json:"message"`
	Data    ResponseCustomerDashboardData `json:"data"`
}

type ResponseCustomerProfileData struct {
	ReferalCode   string    `json:"referal_code"`
	Name          string    `json:"name"`
	Phone         string    `json:"phone"`
	Email         string    `json:"email"`
	Sex           string    `json:"sex"`
	BirthPlace    string    `json:"birth_place"`
	BirthDate     string    `json:"birth_date"`
	NIK           string    `json:"nik"`
	NPWP          string    `json:"npwp"`
	Passport      string    `json:"passport"`
	Age           string    `json:"age"`
	LastEdcuation string    `json:"last_education"`
	MaritalStatus string    `json:"marital_status"`
	Photo         string    `json:"photo"`
	RegisterDate  time.Time `json:"register_date"`
	Address       string    `json:"address"`
	Status        uint64    `json:"status"`
}

type ResponseCustomerProfile struct {
	Status  uint64                      `json:"status"`
	Message string                      `json:"message"`
	Data    ResponseCustomerProfileData `json:"data"`
}

type ResponseCustomerProfileAddressData struct {
	ReferalCode     string    `json:"referal_code"`
	ReferalCodeFrom string    `json:"referal_code_from"`
	Name            string    `json:"name"`
	Phone           string    `json:"phone"`
	Email           string    `json:"email"`
	Sex             string    `json:"sex"`
	BirthPlace      string    `json:"birth_place"`
	BirthDate       string    `json:"birth_date"`
	NIK             string    `json:"nik"`
	NPWP            string    `json:"npwp"`
	Passport        string    `json:"passport"`
	Age             string    `json:"age"`
	LastEdcuation   string    `json:"last_education"`
	MaritalStatus   string    `json:"marital_status"`
	Photo           string    `json:"photo"`
	Address         string    `json:"address"`
	CountryID       uint64    `json:"country_id"`
	Country         string    `json:"country_name"`
	ProvinceID      uint64    `json:"province_id"`
	Province        string    `json:"province_name"`
	RegencyID       uint64    `json:"regency_id"`
	Regency         string    `json:"regency_name"`
	DistrictID      uint64    `json:"district_id"`
	District        string    `json:"district_name"`
	VillageID       uint64    `json:"village_id"`
	Village         string    `json:"village_name"`
	ZipCode         string    `json:"zip_code"`
	RegisterDate    time.Time `json:"register_date"`
	Status          uint64    `json:"status"`
}

type ResponseCustomerProfileAddress struct {
	Status  uint64                             `json:"status"`
	Message string                             `json:"message"`
	Data    ResponseCustomerProfileAddressData `json:"data"`
}

type ResponseCustomerUpdateProfileData struct {
	Customer     uint64 `json:"customer_id"`
	UserCustomer uint64 `json:"customer_user_id"`
}

type ResponseUpdateCustomerProfile struct {
	Status  uint64                            `json:"status"`
	Message string                            `json:"message"`
	Data    ResponseCustomerUpdateProfileData `json:"data"`
}

func (DetailMobileUser) TableName() string {
	return "m_customer_user"
}

func (MobileUserData) TableName() string {
	return "m_customer_user"
}

func (CustomerUser) TableName() string {
	return "m_customer_user"
}

func (CustomerUserData) TableName() string {
	return "m_customer_user"
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *CustomerUser) Prepare() {
	u.CustomerID = u.CustomerID
	u.CustomerUserReferalCode = html.EscapeString(strings.TrimSpace(u.CustomerUserReferalCode))
	u.CustomerUserEmail = html.EscapeString(strings.TrimSpace(u.CustomerUserEmail))
	u.CustomerUsername = html.EscapeString(strings.TrimSpace(u.CustomerUsername))
	u.CustomerUserPassword = html.EscapeString(strings.TrimSpace(u.CustomerUserPassword))
	u.CustomerUserPhoneCode = "62"
	u.CustomerUserPhone = html.EscapeString(strings.TrimSpace(u.CustomerUserPhone))
	u.CustomerUserOTP = html.EscapeString(strings.TrimSpace(u.CustomerUserOTP))
	u.CustomerUserOTPSecret = u.CustomerUserOTPSecret
	u.CustomerUserPin = html.EscapeString(strings.TrimSpace(u.CustomerUserPin))
	u.CustomerUserLastLogin = u.CustomerUserLastLogin
	u.CustomerUserImei = html.EscapeString(strings.TrimSpace(u.CustomerUserImei))
	u.CustomerUserIPAddress = html.EscapeString(strings.TrimSpace(u.CustomerUserIPAddress))
	u.CustomerUserEmailVerifiedAt = u.CustomerUserEmailVerifiedAt
	u.CustomerUserStatus = u.CustomerUserStatus
	u.CustomerUserIsNew = u.CustomerUserIsNew
	u.CustomerUserIsRequestChangePassword = u.CustomerUserIsRequestChangePassword
	u.CustomerUserIsLocked = u.CustomerUserIsLocked
	u.CustomerUserFailAttempt = u.CustomerUserFailAttempt
	u.CustomerUserLockedReason = u.CustomerUserLockedReason
	u.CustomerUserCreatedBy = u.CustomerUserCreatedBy
	u.CustomerUserCreatedAt = time.Now()
	u.CustomerUserUpdatedBy = u.CustomerUserUpdatedBy
	u.CustomerUserUpdatedAt = u.CustomerUserUpdatedAt
	u.CustomerUserDeletedBy = u.CustomerUserDeletedBy
	u.CustomerUserDeletedAt = u.CustomerUserDeletedAt
}

func (u CustomerUser) Validation() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.CustomerUserPhone, validation.Length(11, 13)),
	)
}

func VerifyPinMobile(hashedPin, pin string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPin), []byte(pin))
}

func (u *CustomerUser) Validate(action string) error {
	switch strings.ToLower(action) {
	case "register":
		if u.CustomerUserEmail == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.CustomerUserEmail); err != nil {
			return errors.New("Invalid Email")
		}
		if u.CustomerUserPhone == "" {
			return errors.New("Required Phone")
		}
		if u.CustomerUserImei == "" {
			return errors.New("Required IMEI")
		}
		if u.CustomerLoginStaticToken != "SaCCQVw]aY8a=:s3!'j*cAIL5uoLyp" {
			return errors.New("Required Static Token")
		}
		return nil

	case "setpin":
		if u.CustomerUserPin == "" {
			return errors.New("Required PIN")
		}
		if u.CustomerLoginStaticToken != "SaCCQVw]aY8a=:s3!'j*cAIL5uoLyp" {
			return errors.New("Required Static Token")
		}
		return nil

	case "phone":
		if u.CustomerUserPhone == "" {
			return errors.New("Required Phone")
		}
		if u.CustomerLoginStaticToken != "SaCCQVw]aY8a=:s3!'j*cAIL5uoLyp" {
			return errors.New("Required Static Token")
		}
		return nil

	case "phone-profile":
		if u.CustomerUserPhone == "" {
			return errors.New("Required Phone")
		}
		return nil

	case "sendotp":
		if u.CustomerLoginStaticToken != "SaCCQVw]aY8a=:s3!'j*cAIL5uoLyp" {
			return errors.New("Required Static Token")
		}
		if u.CustomerUserPhone == "" {
			return errors.New("Required Phone")
		}
		return nil

	case "verifyotp":
		if u.CustomerLoginStaticToken != "SaCCQVw]aY8a=:s3!'j*cAIL5uoLyp" {
			return errors.New("Required Static Token")
		}
		if u.CustomerUserOTPNew == "" {
			return errors.New("Required OTP")
		}
		if u.CustomerUserPhone == "" {
			return errors.New("Required Phone")
		}
		return nil

	case "verifypin":
		if u.CustomerUserPin == "" {
			return errors.New("Required PIN")
		}
		// if u.CustomerUserPhone == "" {
		// 	return errors.New("Required Phone")
		// }
		return nil

	case "updatepin":
		if u.CustomerUserPinNew == "" {
			return errors.New("Required PIN")
		}
		if u.CustomerUserPinOld == "" {
			return errors.New("Required Phone")
		}
		return nil

	case "pin":
		if u.CustomerUserPin == "" {
			return errors.New("Required PIN")
		}
		if u.CustomerUserPhone == "" {
			return errors.New("Required Phone")
		}
		return nil

	case "verifypswd":
		if u.CustomerUserPassword == "" {
			return errors.New("Required Password")
		}
		if u.CustomerUserPhone == "" {
			return errors.New("Required Phone")
		}
		return nil

	case "updatepswd":
		if u.CustomerUserPswdNew == "" {
			return errors.New("Required Password")
		}
		if u.CustomerUserPswdOld == "" {
			return errors.New("Required Phone")
		}
		return nil

	case "update":
		if u.CustomerUserEmail == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.CustomerUserEmail); err != nil {
			return errors.New("Invalid Email")
		}
		if u.CustomerUsername == "" {
			return errors.New("Username Wajib Diisi")
		}
		if u.CustomerUserPassword == "" {
			return errors.New("Password Wajib Diisi")
		}
		if u.CustomerUserPhone == "" {
			return errors.New("Required Phone")
		}
		if u.CustomerUserOTP == "" {
			return errors.New("Required OTP")
		}
		if u.CustomerUserPin == "" {
			return errors.New("Required PIN")
		}
		if u.CustomerUserImei == "" {
			return errors.New("Required IMEI")
		}
		if u.CustomerUserIPAddress == "" {
			return errors.New("Required IP Address")
		}
		if u.CustomerUserStatus == 0 {
			return errors.New("Required User Status")
		}
		return nil

	default:
		if u.CustomerID == 0 {
			return errors.New("Required Customer ID")
		}
		return nil
	}
}

func (u *CustomerUser) SaveCustomerUser(db *gorm.DB) (*CustomerUser, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &CustomerUser{}, err
	}
	return u, nil
}

func (p *CustomerUser) FindMobileUser(db *gorm.DB) (*[]MobileUserData, error) {
	var err error
	productsubcategory := []MobileUserData{}
	err = db.Debug().Model(&MobileUserData{}).Limit(100).
		Select(`m_customer_user.cust_user_phone,
				m_customer_user.cust_user_id,
				m_customer_user.cust_id,
				m_customer_user.cust_user_email,
				m_customer.cust_name,
				m_customer_user.cust_user_status`).
		Joins("JOIN m_customer on m_customer_user.cust_id=m_customer.cust_id").
		Find(&productsubcategory).Error
	if err != nil {
		return &[]MobileUserData{}, err
	}
	return &productsubcategory, nil
}

func (p *CustomerUser) FindMobileUserByID(db *gorm.DB, pid uint64) (*DetailMobileUser, error) {
	var err error
	productsubcategory := DetailMobileUser{}
	err = db.Debug().Model(&DetailMobileUser{}).Limit(100).
		Select(`m_customer_user.cust_user_phone,
				m_customer_user.cust_user_id,
				m_customer.cust_ktp,
				m_customer.cust_name,
				m_customer.cust_sex,
				m_customer.cust_birth_date,
				m_customer.cust_birth_place,
				m_customer.cust_age,
				m_customer_user.cust_user_email,
				m_customer_user.cust_user_status,
				m_customer_user.cust_user_created_by,
				m_customer_user.cust_user_created_at`).
		Joins("JOIN m_customer on m_customer_user.cust_id=m_customer.cust_id").
		Where("cust_user_id = ?", pid).
		Find(&productsubcategory).Error
	if err != nil {
		return &DetailMobileUser{}, err
	}
	return &productsubcategory, nil
}

func (p *CustomerUser) FindAllCustomerUser(db *gorm.DB) (*[]CustomerUserData, error) {
	var err error
	productsubcategory := []CustomerUserData{}
	err = db.Debug().Model(&CustomerUserData{}).Limit(100).
		Select(`m_customer_user.cust_user_id,
				m_customer_user.cust_id,
				m_customer.cust_code,
				m_customer.cust_ktp,
				m_customer.cust_passport,
				m_customer.cust_npwp,
				m_customer.cust_name,
				m_customer.cust_sex,
				m_customer.cust_birth_date,
				m_customer.cust_birth_place,
				m_customer.cust_age,
				m_customer.cust_phone_work,
				m_customer.cust_email_work,
				m_customer.cust_last_education,
				m_customer.cust_marital_status,
				m_customer.cust_image,
				m_customer_user.cust_user_referal_code,
				m_customer_user.cust_user_referal_code_from,
				m_customer_user.cust_user_email,
				m_customer_user.cust_username,
				m_customer_user.cust_user_password,
				m_customer_user.cust_user_phone_code,
				m_customer_user.cust_user_phone,
				m_customer_user.cust_user_otp,
				m_customer_user.cust_user_otp_secret,
				m_customer_user.cust_user_secret,
				m_customer_user.cust_user_pin,
				m_customer_user.cust_user_last_login,
				m_customer_user.cust_user_imei,
				m_customer_user.cust_user_ip_address,
				m_customer_user.cust_user_email_verified_at,
				m_customer_user.cust_user_status,
				m_customer_user.cust_user_is_new,
				m_customer_user.cust_user_is_req_chpw,
				m_customer_user.cust_user_is_locked,
				m_customer_user.cust_user_fail_attempt,
				m_customer_user.cust_user_pin,
				m_customer_user.cust_user_locked_reason,
				m_customer_user.cust_user_created_by,
				m_customer_user.cust_user_created_at,
				m_customer_user.cust_user_updated_by,
				m_customer_user.cust_user_updated_at,
				m_customer_user.cust_user_deleted_by,
				m_customer_user.cust_user_deleted_at`).
		Joins("JOIN m_customer on m_customer_user.cust_id=m_customer.cust_id").
		Find(&productsubcategory).Error
	if err != nil {
		return &[]CustomerUserData{}, err
	}
	return &productsubcategory, nil
}

func (p *CustomerUser) FindAllCustomerUserStatus(db *gorm.DB, status uint64) (*[]CustomerUserData, error) {
	var err error
	productsubcategory := []CustomerUserData{}
	err = db.Debug().Model(&CustomerUserData{}).Limit(100).
		Select(`m_customer_user.cust_user_id,
				m_customer_user.cust_id,
				m_customer.cust_code,
				m_customer.cust_ktp,
				m_customer.cust_passport,
				m_customer.cust_npwp,
				m_customer.cust_name,
				m_customer.cust_sex,
				m_customer.cust_birth_date,
				m_customer.cust_birth_place,
				m_customer.cust_age,
				m_customer.cust_phone_work,
				m_customer.cust_email_work,
				m_customer.cust_last_education,
				m_customer.cust_marital_status,
				m_customer.cust_image,
				m_customer_user.cust_user_referal_code,
				m_customer_user.cust_user_referal_code_from,
				m_customer_user.cust_user_email,
				m_customer_user.cust_username,
				m_customer_user.cust_user_password,
				m_customer_user.cust_user_phone_code,
				m_customer_user.cust_user_phone,
				m_customer_user.cust_user_otp,
				m_customer_user.cust_user_otp_secret,
				m_customer_user.cust_user_secret,
				m_customer_user.cust_user_pin,
				m_customer_user.cust_user_last_login,
				m_customer_user.cust_user_imei,
				m_customer_user.cust_user_ip_address,
				m_customer_user.cust_user_email_verified_at,
				m_customer_user.cust_user_status,
				m_customer_user.cust_user_is_new,
				m_customer_user.cust_user_is_req_chpw,
				m_customer_user.cust_user_is_locked,
				m_customer_user.cust_user_fail_attempt,
				m_customer_user.cust_user_pin,
				m_customer_user.cust_user_locked_reason,
				m_customer_user.cust_user_created_by,
				m_customer_user.cust_user_created_at,
				m_customer_user.cust_user_updated_by,
				m_customer_user.cust_user_updated_at,
				m_customer_user.cust_user_deleted_by,
				m_customer_user.cust_user_deleted_at`).
		Joins("JOIN m_customer on m_customer_user.cust_id=m_customer.cust_id").
		Where("cust_user_status = ?", status).
		Joins("JOIN m_product_category on m_customer_user.prod_ctgry_id=m_product_category.prod_ctgry_id").
		Find(&productsubcategory).Error
	if err != nil {
		return &[]CustomerUserData{}, err
	}
	return &productsubcategory, nil
}

func (p *CustomerUser) FindCustomerUserByID(db *gorm.DB, pid uint64) (*CustomerUserData, error) {
	var err error
	productsubcategory := CustomerUserData{}
	err = db.Debug().Model(&CustomerUserData{}).Limit(100).
		Select(`m_customer_user.cust_user_id,
				m_customer_user.cust_id,
				m_customer_user.cust_sub_type_id,
				m_customer.cust_code,
				m_customer.cust_ktp,
				m_customer.cust_passport,
				m_customer.cust_npwp,
				m_customer.cust_name,
				m_customer.cust_sex,
				m_customer.cust_birth_date,
				m_customer.cust_birth_place,
				m_customer.cust_age,
				m_customer.cust_phone_work,
				m_customer.cust_email_work,
				m_customer.cust_last_education,
				m_customer.cust_marital_status,
				m_customer.cust_image,
				m_customer_user.cust_sub_type_id,
				m_customer_user.cust_user_referal_code,
				m_customer_user.cust_user_referal_code_from,
				m_customer_user.cust_user_email,
				m_customer_user.cust_username,
				m_customer_user.cust_user_password,
				m_customer_user.cust_user_phone_code,
				m_customer_user.cust_user_phone,
				m_customer_user.cust_user_otp,
				m_customer_user.cust_user_otp_secret,
				m_customer_user.cust_user_secret,
				m_customer_user.cust_user_pin,
				m_customer_user.cust_user_last_login,
				m_customer_user.cust_user_imei,
				m_customer_user.cust_user_ip_address,
				m_customer_user.cust_user_email_verified_at,
				m_customer_user.cust_user_status,
				m_customer_user.cust_user_is_new,
				m_customer_user.cust_user_is_req_chpw,
				m_customer_user.cust_user_is_locked,
				m_customer_user.cust_user_fail_attempt,
				m_customer_user.cust_user_pin,
				m_customer_user.cust_user_locked_reason,
				m_customer_user.cust_user_created_by,
				m_customer_user.cust_user_created_at,
				m_customer_user.cust_user_updated_by,
				m_customer_user.cust_user_updated_at,
				m_customer_user.cust_user_deleted_by,
				m_customer_user.cust_user_deleted_at`).
		Joins("JOIN m_customer on m_customer_user.cust_id=m_customer.cust_id").
		Where("cust_user_id = ?", pid).
		Find(&productsubcategory).Error
	if err != nil {
		return &CustomerUserData{}, err
	}
	return &productsubcategory, nil
}

func (p *CustomerUser) FindCustomerUserByCustomerID(db *gorm.DB, pid uint64) (*CustomerUserData, error) {
	var err error
	productsubcategory := CustomerUserData{}
	err = db.Debug().Model(&CustomerUserData{}).Limit(100).
		Select(`m_customer_user.cust_user_id,
				m_customer_user.cust_id,
				m_customer_user.cust_sub_type_id,
				m_customer.cust_code,
				m_customer.cust_ktp,
				m_customer.cust_passport,
				m_customer.cust_npwp,
				m_customer.cust_name,
				m_customer.cust_sex,
				m_customer.cust_birth_date,
				m_customer.cust_birth_place,
				m_customer.cust_age,
				m_customer.cust_phone_work,
				m_customer.cust_email_work,
				m_customer.cust_last_education,
				m_customer.cust_marital_status,
				m_customer.cust_image,
				m_customer_user.cust_sub_type_id,
				m_customer_user.cust_user_referal_code,
				m_customer_user.cust_user_referal_code_from,
				m_customer_user.cust_user_email,
				m_customer_user.cust_username,
				m_customer_user.cust_user_password,
				m_customer_user.cust_user_phone_code,
				m_customer_user.cust_user_phone,
				m_customer_user.cust_user_otp,
				m_customer_user.cust_user_otp_secret,
				m_customer_user.cust_user_secret,
				m_customer_user.cust_user_pin,
				m_customer_user.cust_user_last_login,
				m_customer_user.cust_user_imei,
				m_customer_user.cust_user_ip_address,
				m_customer_user.cust_user_email_verified_at,
				m_customer_user.cust_user_status,
				m_customer_user.cust_user_is_new,
				m_customer_user.cust_user_is_req_chpw,
				m_customer_user.cust_user_is_locked,
				m_customer_user.cust_user_fail_attempt,
				m_customer_user.cust_user_pin,
				m_customer_user.cust_user_locked_reason,
				m_customer_user.cust_user_created_by,
				m_customer_user.cust_user_created_at,
				m_customer_user.cust_user_updated_by,
				m_customer_user.cust_user_updated_at,
				m_customer_user.cust_user_deleted_by,
				m_customer_user.cust_user_deleted_at`).
		Joins("JOIN m_customer on m_customer_user.cust_id=m_customer.cust_id").
		Where("m_customer_user.cust_id = ?", pid).
		Find(&productsubcategory).Error
	if err != nil {
		return &CustomerUserData{}, err
	}
	return &productsubcategory, nil
}

func (p *CustomerUser) FindCustomerSubTypeByUserID(db *gorm.DB, pid uint64) (*CustomerUserData, error) {
	var err error
	productsubcategory := CustomerUserData{}
	err = db.Debug().Model(&CustomerUserData{}).Limit(100).
		Select(`m_customer_user.cust_user_id,
				m_customer_user.cust_id,
				m_customer_user.cust_sub_type_id,
				m_customer.cust_code,
				m_customer.cust_ktp,
				m_customer.cust_passport,
				m_customer.cust_npwp,
				m_customer.cust_name,
				m_customer.cust_sex,
				m_customer.cust_birth_date,
				m_customer.cust_birth_place,
				m_customer.cust_age,
				m_customer.cust_phone_work,
				m_customer.cust_email_work,
				m_customer.cust_last_education,
				m_customer.cust_marital_status,
				m_customer.cust_image,
				m_customer_user.cust_user_referal_code,
				m_customer_user.cust_user_referal_code_from,
				m_customer_user.cust_user_email,
				m_customer_user.cust_username,
				m_customer_user.cust_user_password,
				m_customer_user.cust_user_phone_code,
				m_customer_user.cust_user_phone,
				m_customer_user.cust_user_otp,
				m_customer_user.cust_user_otp_secret,
				m_customer_user.cust_user_secret,
				m_customer_user.cust_user_pin,
				m_customer_user.cust_user_last_login,
				m_customer_user.cust_user_imei,
				m_customer_user.cust_user_ip_address,
				m_customer_user.cust_user_email_verified_at,
				m_customer_user.cust_user_status,
				m_customer_user.cust_user_is_new,
				m_customer_user.cust_user_is_req_chpw,
				m_customer_user.cust_user_is_locked,
				m_customer_user.cust_user_fail_attempt,
				m_customer_user.cust_user_pin,
				m_customer_user.cust_user_locked_reason,
				m_customer_user.cust_user_created_by,
				m_customer_user.cust_user_created_at,
				m_customer_user.cust_user_updated_by,
				m_customer_user.cust_user_updated_at,
				m_customer_user.cust_user_deleted_by,
				m_customer_user.cust_user_deleted_at,
				m_customer_sub_type.cust_sub_type_id,
				m_customer_sub_type.cust_sub_type_name,
				m_customer_type.cust_type_id,
				m_customer_type.cust_type_name`).
		Joins("JOIN m_customer on m_customer_user.cust_id=m_customer.cust_id").
		Joins("JOIN m_customer_sub_type on m_customer_user.cust_sub_type_id=m_customer_sub_type.cust_sub_type_id").
		Joins("JOIN m_customer_type on m_customer_sub_type.cust_type_id=m_customer_type.cust_type_id").
		Where("cust_user_id = ?", pid).
		Find(&productsubcategory).Error
	if err != nil {
		return &CustomerUserData{}, err
	}
	return &productsubcategory, nil
}

func (p *CustomerUser) FindAllCustomerReferalCode(db *gorm.DB, referal string) (*CustomerUser, error) {
	var err error
	err = db.Debug().Model(&CustomerUser{}).Where("cust_user_referal_code = ? AND mitra_status = ?", referal, 1).Take(&p).Error
	if err != nil {
		return &CustomerUser{}, err
	}
	return p, nil
}

func (p *CustomerUser) FindCustomerUserStatusByID(db *gorm.DB, pid uint64, status uint64) (*CustomerUserData, error) {
	var err error
	productsubcategory := CustomerUserData{}
	err = db.Debug().Model(&CustomerUserData{}).Limit(100).
		Select(`m_customer_user.cust_user_id,
				m_customer_user.cust_id,
				m_customer.cust_code,
				m_customer.cust_ktp,
				m_customer.cust_passport,
				m_customer.cust_npwp,
				m_customer.cust_name,
				m_customer.cust_sex,
				m_customer.cust_birth_date,
				m_customer.cust_birth_place,
				m_customer.cust_age,
				m_customer.cust_phone_work,
				m_customer.cust_email_work,
				m_customer.cust_last_education,
				m_customer.cust_marital_status,
				m_customer.cust_image,
				m_customer_user.cust_user_referal_code,
				m_customer_user.cust_user_referal_code_from,
				m_customer_user.cust_user_email,
				m_customer_user.cust_username,
				m_customer_user.cust_user_password,
				m_customer_user.cust_user_phone_code,
				m_customer_user.cust_user_phone,
				m_customer_user.cust_user_otp,
				m_customer_user.cust_user_otp_secret,
				m_customer_user.cust_user_secret,
				m_customer_user.cust_user_pin,
				m_customer_user.cust_user_last_login,
				m_customer_user.cust_user_imei,
				m_customer_user.cust_user_ip_address,
				m_customer_user.cust_user_email_verified_at,
				m_customer_user.cust_user_status,
				m_customer_user.cust_user_is_new,
				m_customer_user.cust_user_is_req_chpw,
				m_customer_user.cust_user_is_locked,
				m_customer_user.cust_user_fail_attempt,
				m_customer_user.cust_user_pin,
				m_customer_user.cust_user_locked_reason,
				m_customer_user.cust_user_created_by,
				m_customer_user.cust_user_created_at,
				m_customer_user.cust_user_updated_by,
				m_customer_user.cust_user_updated_at,
				m_customer_user.cust_user_deleted_by,
				m_customer_user.cust_user_deleted_at`).
		Joins("JOIN m_customer on m_customer_user.cust_id=m_customer.cust_id").
		Where("cust_user_id = ? AND cust_user_status = ?", pid, status).
		Find(&productsubcategory).Error
	if err != nil {
		return &CustomerUserData{}, err
	}
	return &productsubcategory, nil
}

func (p *CustomerUser) LoginCustomerUserByPhone(db *gorm.DB, phone string) (*CustomerUserData, error) {
	var err error
	productsubcategory := CustomerUserData{}
	err = db.Debug().Model(&CustomerUserData{}).Limit(100).
		Select(`m_customer_user.cust_user_id,
				m_customer.cust_id,
				m_customer.cust_code,
				m_customer.cust_ktp,
				m_customer.cust_passport,
				m_customer.cust_npwp,
				m_customer.cust_name,
				m_customer.cust_sex,
				m_customer.cust_birth_date,
				m_customer.cust_birth_place,
				m_customer.cust_age,
				m_customer.cust_phone_work,
				m_customer.cust_email_work,
				m_customer.cust_last_education,
				m_customer.cust_marital_status,
				m_customer.cust_image,
				m_customer_user.cust_user_referal_code,
				m_customer_user.cust_user_referal_code_from,
				m_customer_user.cust_user_email,
				m_customer_user.cust_username,
				m_customer_user.cust_user_password,
				m_customer_user.cust_user_phone_code,
				m_customer_user.cust_user_phone,
				m_customer_user.cust_user_otp,
				m_customer_user.cust_user_otp_secret,
				m_customer_user.cust_user_secret,
				m_customer_user.cust_user_pin,
				m_customer_user.cust_user_last_login,
				m_customer_user.cust_user_imei,
				m_customer_user.cust_user_ip_address,
				m_customer_user.cust_user_email_verified_at,
				m_customer_user.cust_user_status,
				m_customer_user.cust_user_is_new,
				m_customer_user.cust_user_is_req_chpw,
				m_customer_user.cust_user_is_locked,
				m_customer_user.cust_user_fail_attempt,
				m_customer_user.cust_user_pin,
				m_customer_user.cust_user_locked_reason,
				m_customer_user.cust_user_created_by,
				m_customer_user.cust_user_created_at,
				m_customer_user.cust_user_updated_by,
				m_customer_user.cust_user_updated_at,
				m_customer_user.cust_user_deleted_by,
				m_customer_user.cust_user_deleted_at`).
		Joins("JOIN m_customer on m_customer_user.cust_id=m_customer.cust_id").
		Where("cust_user_phone = ?", phone).
		Find(&productsubcategory).Error
	if err != nil {
		return &CustomerUserData{}, err
	}
	return &productsubcategory, nil
}

func (p *CustomerUser) CheckCustomerUserByPhone(db *gorm.DB, phone string) (*CustomerUser, error) {
	var err error
	productsubcategory := CustomerUser{}
	err = db.Debug().Model(&CustomerUser{}).Limit(100).
		Select(`m_customer_user.cust_user_id,
				m_customer_user.cust_id,
				m_customer_user.cust_user_referal_code,
				m_customer_user.cust_user_referal_code_from,
				m_customer_user.cust_user_email,
				m_customer_user.cust_username,
				m_customer_user.cust_user_password,
				m_customer_user.cust_user_phone_code,
				m_customer_user.cust_user_phone,
				m_customer_user.cust_user_otp,
				m_customer_user.cust_user_otp_secret,
				m_customer_user.cust_user_secret,
				m_customer_user.cust_user_pin,
				m_customer_user.cust_user_last_login,
				m_customer_user.cust_user_imei,
				m_customer_user.cust_user_ip_address,
				m_customer_user.cust_user_email_verified_at,
				m_customer_user.cust_user_status,
				m_customer_user.cust_user_is_new,
				m_customer_user.cust_user_is_req_chpw,
				m_customer_user.cust_user_is_locked,
				m_customer_user.cust_user_fail_attempt,
				m_customer_user.cust_user_pin,
				m_customer_user.cust_user_locked_reason,
				m_customer_user.cust_user_created_by,
				m_customer_user.cust_user_created_at,
				m_customer_user.cust_user_updated_by,
				m_customer_user.cust_user_updated_at,
				m_customer_user.cust_user_deleted_by,
				m_customer_user.cust_user_deleted_at`).
		Where("cust_user_phone = ?", phone).
		Find(&productsubcategory).Error
	if err != nil {
		return &CustomerUser{}, err
	}
	return &productsubcategory, nil
}

func (p *CustomerUser) LoginCustomerUserByPin(db *gorm.DB, phone, pin string) (*CustomerUserData, error) {
	var err error
	productsubcategory := CustomerUserData{}
	err = db.Debug().Model(&CustomerUserData{}).Limit(100).
		Select(`m_customer_user.cust_user_id,
				m_customer_user.cust_id,
				m_customer.cust_code,
				m_customer.cust_ktp,
				m_customer.cust_passport,
				m_customer.cust_npwp,
				m_customer.cust_name,
				m_customer.cust_sex,
				m_customer.cust_birth_date,
				m_customer.cust_birth_place,
				m_customer.cust_age,
				m_customer.cust_phone_work,
				m_customer.cust_email_work,
				m_customer.cust_last_education,
				m_customer.cust_marital_status,
				m_customer.cust_image,
				m_customer_user.cust_user_referal_code,
				m_customer_user.cust_user_referal_code_from,
				m_customer_user.cust_user_email,
				m_customer_user.cust_username,
				m_customer_user.cust_user_password,
				m_customer_user.cust_user_phone_code,
				m_customer_user.cust_user_phone,
				m_customer_user.cust_user_otp,
				m_customer_user.cust_user_otp_secret,
				m_customer_user.cust_user_secret,
				m_customer_user.cust_user_pin,
				m_customer_user.cust_user_last_login,
				m_customer_user.cust_user_imei,
				m_customer_user.cust_user_ip_address,
				m_customer_user.cust_user_email_verified_at,
				m_customer_user.cust_user_status,
				m_customer_user.cust_user_is_new,
				m_customer_user.cust_user_is_req_chpw,
				m_customer_user.cust_user_is_locked,
				m_customer_user.cust_user_fail_attempt,
				m_customer_user.cust_user_pin,
				m_customer_user.cust_user_locked_reason,
				m_customer_user.cust_user_created_by,
				m_customer_user.cust_user_created_at,
				m_customer_user.cust_user_updated_by,
				m_customer_user.cust_user_updated_at,
				m_customer_user.cust_user_deleted_by,
				m_customer_user.cust_user_deleted_at`).
		Joins("JOIN m_customer on m_customer_user.cust_id=m_customer.cust_id").
		Where("cust_user_phone = ? AND cust_user_pin = ?", phone, pin).
		Find(&productsubcategory).Error
	if err != nil {
		return &CustomerUserData{}, err
	}
	return &productsubcategory, nil
}

func (p *CustomerUser) LoginCustomerUserByPswd(db *gorm.DB, phone, pswd string) (*CustomerUserData, error) {
	var err error
	productsubcategory := CustomerUserData{}
	err = db.Debug().Model(&CustomerUserData{}).Limit(100).
		Select(`m_customer_user.cust_user_id,
				m_customer_user.cust_id,
				m_customer.cust_code,
				m_customer.cust_ktp,
				m_customer.cust_passport,
				m_customer.cust_npwp,
				m_customer.cust_name,
				m_customer.cust_sex,
				m_customer.cust_birth_date,
				m_customer.cust_birth_place,
				m_customer.cust_age,
				m_customer.cust_phone_work,
				m_customer.cust_email_work,
				m_customer.cust_last_education,
				m_customer.cust_marital_status,
				m_customer.cust_image,
				m_customer_user.cust_user_referal_code,
				m_customer_user.cust_user_referal_code_from,
				m_customer_user.cust_user_email,
				m_customer_user.cust_username,
				m_customer_user.cust_user_password,
				m_customer_user.cust_user_phone_code,
				m_customer_user.cust_user_phone,
				m_customer_user.cust_user_otp,
				m_customer_user.cust_user_otp_secret,
				m_customer_user.cust_user_secret,
				m_customer_user.cust_user_pin,
				m_customer_user.cust_user_last_login,
				m_customer_user.cust_user_imei,
				m_customer_user.cust_user_ip_address,
				m_customer_user.cust_user_email_verified_at,
				m_customer_user.cust_user_status,
				m_customer_user.cust_user_is_new,
				m_customer_user.cust_user_is_req_chpw,
				m_customer_user.cust_user_is_locked,
				m_customer_user.cust_user_fail_attempt,
				m_customer_user.cust_user_pin,
				m_customer_user.cust_user_locked_reason,
				m_customer_user.cust_user_created_by,
				m_customer_user.cust_user_created_at,
				m_customer_user.cust_user_updated_by,
				m_customer_user.cust_user_updated_at,
				m_customer_user.cust_user_deleted_by,
				m_customer_user.cust_user_deleted_at`).
		Joins("JOIN m_customer on m_customer_user.cust_id=m_customer.cust_id").
		Where("cust_user_phone = ? AND cust_user_password = ?", phone, pswd).
		Find(&productsubcategory).Error
	if err != nil {
		return &CustomerUserData{}, err
	}
	return &productsubcategory, nil
}

func (p *CustomerUser) CheckPhoneCustomerUser(db *gorm.DB, phone string) (*DetailMobileUser, error) {
	var err error
	productsubcategory := DetailMobileUser{}
	err = db.Debug().Model(&DetailMobileUser{}).Limit(100).
		Select(`m_customer_user.cust_user_phone,
				m_customer_user.cust_user_id,
				m_customer_user.cust_user_created_at`).
		Where("cust_user_phone = ?", phone).
		Count(&productsubcategory.CountPhoneCustomer).
		Find(&productsubcategory).Error
	if err != nil {
		return &DetailMobileUser{}, err
	}
	return &productsubcategory, nil
}

func (p *CustomerUser) CheckEmailCustomerUser(db *gorm.DB, email string) (*DetailMobileUser, error) {
	var err error
	productsubcategory := DetailMobileUser{}
	err = db.Debug().Model(&DetailMobileUser{}).Limit(100).
		Select(`m_customer_user.cust_user_email,
				m_customer_user.cust_user_id,
				m_customer_user.cust_user_created_at`).
		Where("cust_user_email = ?", email).
		Count(&productsubcategory.CountEmailCustomer).
		Find(&productsubcategory).Error
	if err != nil {
		return &DetailMobileUser{}, err
	}
	return &productsubcategory, nil
}

func (u *CustomerUser) UpdateCustomerUser(db *gorm.DB, uid uint64) (*CustomerUser, error) {
	// To hash the password
	err := u.BeforeSaveCustomer()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&CustomerUser{}).Where("cust_user_id = ?", uid).Take(&CustomerUser{}).UpdateColumns(
		map[string]interface{}{
			"cust_user_email":      u.CustomerUserEmail,
			"cust_username":        u.CustomerUsername,
			"cust_user_phone":      u.CustomerUserPhone,
			"cust_user_otp":        u.CustomerUserOTP,
			"cust_user_secret":     u.CustomerUserSecret,
			"cust_user_otp_secret": time.Now(),
			"cust_user_imei":       u.CustomerUserImei,
			"cust_user_ip_address": u.CustomerUserIPAddress,
			"cust_user_status":     u.CustomerUserStatus,
			"cust_user_updated_by": u.CustomerUserUpdatedBy,
			"cust_user_updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &CustomerUser{}, db.Error
	}
	// This is the display the updated customeruser
	err = db.Debug().Model(&CustomerUser{}).Where("cust_user_id = ?", uid).Take(&u).Error
	if err != nil {
		return &CustomerUser{}, err
	}
	return u, nil
}

func (u *CustomerUser) UpdateCustomerUserJob(db *gorm.DB, uid uint64) (*CustomerUser, error) {
	// To hash the password
	err := u.BeforeSaveCustomer()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&CustomerUser{}).Where("cust_user_id = ?", uid).Take(&CustomerUser{}).UpdateColumns(
		map[string]interface{}{
			"cust_sub_type_id":     u.CustomerSubTypeID,
			"cust_user_updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &CustomerUser{}, db.Error
	}
	// This is the display the updated customeruser
	err = db.Debug().Model(&CustomerUser{}).Where("cust_user_id = ?", uid).Take(&u).Error
	if err != nil {
		return &CustomerUser{}, err
	}
	return u, nil
}

func (u *CustomerUser) SuspendCustomerUser(db *gorm.DB, uid uint64) (*CustomerUser, error) {
	var err error
	db = db.Debug().Model(&CustomerUser{}).Where("cust_user_id = ?", uid).Take(&CustomerUser{}).UpdateColumns(
		map[string]interface{}{
			"cust_user_is_locked":     u.CustomerUserIsLocked,
			"cust_user_locked_reason": u.CustomerUserLockedReason,
			"cust_user_status":        u.CustomerUserStatus,
			"cust_user_updated_by":    u.CustomerUserUpdatedBy,
			"cust_user_updated_at":    time.Now(),
		},
	)
	if db.Error != nil {
		return &CustomerUser{}, db.Error
	}
	// This is the display the updated customeruser
	err = db.Debug().Model(&CustomerUser{}).Where("cust_user_id = ?", uid).Take(&u).Error
	if err != nil {
		return &CustomerUser{}, err
	}
	return u, nil
}

func (u *CustomerUser) UpdateCustomerUserOTP(db *gorm.DB, uid uint64) (*CustomerUser, error) {

	// To hash the password
	err := u.BeforeSaveCustomer()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&CustomerUser{}).Where("cust_user_id = ?", uid).Take(&CustomerUser{}).UpdateColumns(
		map[string]interface{}{
			"cust_user_otp":        u.CustomerUserOTP,
			"cust_user_secret":     u.CustomerUserSecret,
			"cust_user_otp_secret": time.Now(),
			"cust_user_updated_by": u.CustomerUserUpdatedBy,
			"cust_user_updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &CustomerUser{}, db.Error
	}
	// This is the display the updated customeruser
	err = db.Debug().Model(&CustomerUser{}).Where("cust_user_id = ?", uid).Take(&u).Error
	if err != nil {
		return &CustomerUser{}, err
	}
	return u, nil
}

func (u *CustomerUser) SetPinCustomerUser(db *gorm.DB, uid uint64) (*CustomerUser, error) {

	// To hash the password
	err := u.BeforeSaveCustomerPin()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&CustomerUser{}).Where("cust_id = ?", uid).Take(&CustomerUser{}).UpdateColumns(
		map[string]interface{}{
			"cust_user_pin":        u.CustomerUserPin,
			"cust_user_updated_by": u.CustomerUserUpdatedBy,
			"cust_user_updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &CustomerUser{}, db.Error
	}
	// This is the display the updated customeruser
	err = db.Debug().Model(&CustomerUser{}).Where("cust_id = ?", uid).Take(&u).Error
	if err != nil {
		return &CustomerUser{}, err
	}
	return u, nil
}

func (u *CustomerUser) UpdatePinCustomerUser(db *gorm.DB, uid uint64) (*CustomerUser, error) {

	// To hash the password
	err := u.BeforeUpdateCustomerPin()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&CustomerUser{}).Where("cust_id = ?", uid).Take(&CustomerUser{}).UpdateColumns(
		map[string]interface{}{
			"cust_user_pin":        u.CustomerUserPinNew,
			"cust_user_updated_by": u.CustomerUserUpdatedBy,
			"cust_user_updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &CustomerUser{}, db.Error
	}
	// This is the display the updated customeruser
	err = db.Debug().Model(&CustomerUser{}).Where("cust_id = ?", uid).Take(&u).Error
	if err != nil {
		return &CustomerUser{}, err
	}
	return u, nil
}

func (u *CustomerUser) UpdatePswdCustomerUser(db *gorm.DB, uid uint64) (*CustomerUser, error) {

	// To hash the password
	err := u.BeforeSaveCustomer()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&CustomerUser{}).Where("cust_id = ?", uid).Take(&CustomerUser{}).UpdateColumns(
		map[string]interface{}{
			"cust_user_password":   u.CustomerUserPswdNew,
			"cust_user_updated_by": u.CustomerUserUpdatedBy,
			"cust_user_updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &CustomerUser{}, db.Error
	}
	// This is the display the updated customeruser
	err = db.Debug().Model(&CustomerUser{}).Where("cust_id = ?", uid).Take(&u).Error
	if err != nil {
		return &CustomerUser{}, err
	}
	return u, nil
}

func (u *CustomerUser) UpdateProfileCustomer(db *gorm.DB, uid uint64) (*CustomerUser, error) {

	// To hash the password
	err := u.BeforeSaveCustomer()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&CustomerUser{}).Where("cust_id = ?", uid).Take(&CustomerUser{}).UpdateColumns(
		map[string]interface{}{
			"cust_user_phone":      u.CustomerUserPhone,
			"cust_user_email":      u.CustomerUserEmail,
			"cust_user_updated_by": u.CustomerUserUpdatedBy,
			"cust_user_updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &CustomerUser{}, db.Error
	}
	// This is the display the updated customeruser
	err = db.Debug().Model(&CustomerUser{}).Where("cust_id = ?", uid).Error
	if err != nil {
		return &CustomerUser{}, err
	}
	return u, nil
}

func (p *CustomerUser) UpdateCustomerUserSecret(db *gorm.DB, pid uint64) (*CustomerUser, error) {

	var err error
	db = db.Debug().Model(&CustomerUser{}).Where("cust_user_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"cust_user_secret":     p.CustomerUserSecret,
			"cust_user_updated_by": p.CustomerUserUpdatedBy,
			"cust_user_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&CustomerUser{}).Where("cust_user_id = ?", pid).Error
	if err != nil {
		return &CustomerUser{}, err
	}
	return p, nil
}

func (p *CustomerUser) UpdateCustomerUserOTPBlank(db *gorm.DB, pid uint64) (*CustomerUser, error) {

	var err error
	db = db.Debug().Model(&CustomerUser{}).Where("cust_user_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"cust_user_otp":        nil,
			"cust_user_otp_secret": nil,
			"cust_user_updated_by": p.CustomerUserUpdatedBy,
			"cust_user_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&CustomerUser{}).Where("cust_user_id = ?", pid).Error
	if err != nil {
		return &CustomerUser{}, err
	}
	return p, nil
}

func (u *CustomerUser) DeleteCustomerUser(db *gorm.DB, uid uint64) (int64, error) {

	db = db.Debug().Model(&CustomerUser{}).Where("cust_user_id = ?", uid).Take(&CustomerUser{}).Delete(&CustomerUser{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//Customer User LOGIN
//====================================================================================================================================

func (u *CustomerUser) BeforeSaveCustomer() error {
	hashedPassword, err := Hash(u.CustomerUserPassword)
	if err != nil {
		return err
	}
	u.CustomerUserPassword = string(hashedPassword)
	return nil
}

func (u *CustomerUser) BeforeSaveCustomerPin() error {
	hashedPassword, err := Hash(u.CustomerUserPin)
	if err != nil {
		return err
	}
	u.CustomerUserPin = string(hashedPassword)
	return nil
}

func (u *CustomerUser) BeforeUpdateCustomerPin() error {
	hashedPassword, err := Hash(u.CustomerUserPinNew)
	if err != nil {
		return err
	}
	u.CustomerUserPinNew = string(hashedPassword)
	return nil
}
