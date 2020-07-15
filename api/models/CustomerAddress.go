package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type CustomerAddress struct {
	CustomerAddressID            uint64     `gorm:"column:cust_adrs_id;primary_key;" json:"cust_adrs_id"`
	CustomerID                   uint64     `gorm:"column:cust_id;" json:"cust_id"`
	CustomerAddressStreet        string     `gorm:"column:cust_adrs_street" json:"cust_adrs_street"`
	CustomerAddressRT            string     `gorm:"column:cust_adrs_rt;size:3" json:"cust_adrs_rt"`
	CustomerAddressRW            string     `gorm:"column:cust_adrs_rw;size:3" json:"cust_adrs_rw"`
	AddressID                    uint64     `gorm:"column:address_id;" json:"address_id"`
	CustomerAddressZipCode       string     `gorm:"column:cust_adrs_zip_code;size:10" json:"cust_adrs_zip_code"`
	CustomerAddressOwnership     string     `gorm:"column:cust_adrs_ownership;size:125" json:"cust_adrs_ownership"`
	CustomerAddressLengthStay    uint64     `gorm:"column:cust_adrs_length_stay;size:2" json:"cust_adrs_length_stay"`
	CustomerAddressAsResidential uint64     `gorm:"column:cust_adrs_as_residential;size:1" json:"cust_adrs_as_residential"`
	CustomerAddressStatus        uint64     `gorm:"column:cust_adrs_status;size:1" json:"cust_adrs_status"`
	CustomerAddressCreatedBy     string     `gorm:"column:cust_adrs_created_by;size:125" json:"cust_adrs_created_by"`
	CustomerAddressCreatedAt     time.Time  `gorm:"column:cust_adrs_created_at;default:CURRENT_TIMESTAMP" json:"cust_adrs_created_at"`
	CustomerAddressUpdatedBy     string     `gorm:"column:cust_adrs_updated_by;size:125" json:"cust_adrs_updated_by"`
	CustomerAddressUpdatedAt     *time.Time `gorm:"column:cust_adrs_updated_at;default:CURRENT_TIMESTAMP" json:"cust_adrs_updated_at"`
	CustomerAddressDeletedBy     string     `gorm:"column:cust_adrs_deleted_by;size:125" json:"cust_adrs_deleted_by"`
	CustomerAddressDeletedAt     *time.Time `gorm:"column:cust_adrs_deleted_at;default:CURRENT_TIMESTAMP" json:"cust_adrs_deleted_at"`
}

type CustomerAddressData struct {
	CustomerAddressID            uint64     `gorm:"column:cust_adrs_id" json:"cust_adrs_id"`
	CustomerID                   uint64     `gorm:"column:cust_id;" json:"cust_id"`
	CustomerCode                 string     `gorm:"column:cust_code;size:25" json:"cust_code"`
	CustomerKTP                  string     `gorm:"column:cust_ktp;size:255" json:"cust_ktp"`
	CustomerPassport             string     `gorm:"column:cust_passport;size:125" json:"cust_passport"`
	CustomerNPWP                 string     `gorm:"column:cust_npwp;size:255" json:"cust_npwp"`
	CustomerName                 string     `gorm:"column:cust_name;size:255" json:"cust_name"`
	CustomerSex                  string     `gorm:"column:cust_sex;size:6" json:"cust_sex"`
	CustomerBirthDate            string     `gorm:"column:cust_birth_date;size:25" json:"cust_birth_date"`
	CustomerBirthPlace           string     `gorm:"column:cust_birth_place;size:255" json:"cust_birth_place"`
	CustomerAge                  string     `gorm:"column:cust_age;size:3" json:"cust_age"`
	CustomerPhoneWork            string     `gorm:"column:cust_phone_work;size:20" json:"cust_phone_work"`
	CustomerEmailWork            string     `gorm:"column:cust_email_work;size:255" json:"cust_email_work"`
	CustomerMaritalStatus        string     `gorm:"column:cust_marital_status;size:125" json:"cust_marital_status"`
	CustomerImage                string     `gorm:"column:cust_image;size:255" json:"cust_image"`
	CustomerAddressStreet        string     `gorm:"column:cust_adrs_street" json:"cust_adrs_street"`
	CustomerAddressRT            string     `gorm:"column:cust_adrs_rt;size:3" json:"cust_adrs_rt"`
	CustomerAddressRW            string     `gorm:"column:cust_adrs_rw;size:3" json:"cust_adrs_rw"`
	AddressID                    uint64     `gorm:"column:address_id;" json:"address_id"`
	CountryID                    uint64     `gorm:"column:country_id" json:"country_id"`
	CountryName                  string     `gorm:"column:country_name;size:255" json:"country_name"`
	ProvinceID                   uint64     `gorm:"column:province_id" json:"province_id"`
	ProvinceName                 string     `gorm:"column:province_name;size:255" json:"province_name"`
	RegencyID                    uint64     `gorm:"column:regency_id" json:"regency_id"`
	RegencyName                  string     `gorm:"column:regency_name;size:255" json:"regency_name"`
	DistrictID                   uint64     `gorm:"column:district_id" json:"district_id"`
	DistrictName                 string     `gorm:"column:district_name;size:255" json:"district_name"`
	VillageID                    uint64     `gorm:"column:village_id" json:"village_id"`
	VillageName                  string     `gorm:"column:village_name;size:255" json:"village_name"`
	CustomerAddressZipCode       string     `gorm:"column:cust_adrs_zip_code;size:10" json:"cust_adrs_zip_code"`
	CustomerAddressOwnership     string     `gorm:"column:cust_adrs_ownership;size:125" json:"cust_adrs_ownership"`
	CustomerAddressLengthStay    uint64     `gorm:"column:cust_adrs_length_stay;size:2" json:"cust_adrs_length_stay"`
	CustomerAddressAsResidential uint64     `gorm:"column:cust_adrs_as_residential;size:1" json:"cust_adrs_as_residential"`
	CustomerAddressStatus        uint64     `gorm:"column:cust_adrs_status;size:1" json:"cust_adrs_status"`
	CustomerAddressCreatedBy     string     `gorm:"column:cust_adrs_created_by;size:125" json:"cust_adrs_created_by"`
	CustomerAddressCreatedAt     time.Time  `gorm:"column:cust_adrs_created_at;default:CURRENT_TIMESTAMP" json:"cust_adrs_created_at"`
	CustomerAddressUpdatedBy     string     `gorm:"column:cust_adrs_updated_by;size:125" json:"cust_adrs_updated_by"`
	CustomerAddressUpdatedAt     *time.Time `gorm:"column:cust_adrs_updated_at;default:CURRENT_TIMESTAMP" json:"cust_adrs_updated_at"`
	CustomerAddressDeletedBy     string     `gorm:"column:cust_adrs_deleted_by;size:125" json:"cust_adrs_deleted_by"`
	CustomerAddressDeletedAt     *time.Time `gorm:"column:cust_adrs_deleted_at;default:CURRENT_TIMESTAMP" json:"cust_adrs_deleted_at"`
}

type CustomerAddressCount struct {
	CustomerAddressCountValue uint64 `gorm:"column:cust_adrs_count_val" json:"cust_adrs_count_val"`
}

type ResponseCustomerAddresss struct {
	Status  uint64                 `json:"status"`
	Message string                 `json:"message"`
	Data    *[]CustomerAddressData `json:"data"`
}

type ResponseCustomerAddress struct {
	Status  uint64               `json:"status"`
	Message string               `json:"message"`
	Data    *CustomerAddressData `json:"data"`
}

type ResponseCustomerAddressIU struct {
	Status  uint64           `json:"status"`
	Message string           `json:"message"`
	Data    *CustomerAddress `json:"data"`
}

type ResponseCustomerAddressDel struct {
	Status  uint64 `json:"status"`
	Message string `json:"message"`
}

func (CustomerAddress) TableName() string {
	return "m_customer_address"
}

func (CustomerAddressData) TableName() string {
	return "m_customer_address"
}

func (CustomerAddressCount) TableName() string {
	return "m_customer_address"
}

func (u *CustomerAddress) Prepare() {
	u.CustomerID = u.CustomerID
	u.CustomerAddressStreet = html.EscapeString(strings.TrimSpace(u.CustomerAddressStreet))
	u.CustomerAddressRT = html.EscapeString(strings.TrimSpace(u.CustomerAddressRT))
	u.CustomerAddressRW = html.EscapeString(strings.TrimSpace(u.CustomerAddressRW))
	u.AddressID = u.AddressID
	u.CustomerAddressZipCode = html.EscapeString(strings.TrimSpace(u.CustomerAddressZipCode))
	u.CustomerAddressOwnership = html.EscapeString(strings.TrimSpace(u.CustomerAddressOwnership))
	u.CustomerAddressLengthStay = u.CustomerAddressLengthStay
	u.CustomerAddressAsResidential = u.CustomerAddressAsResidential
	u.CustomerAddressStatus = u.CustomerAddressStatus
	u.CustomerAddressCreatedBy = u.CustomerAddressCreatedBy
	u.CustomerAddressCreatedAt = time.Now()
	u.CustomerAddressUpdatedBy = u.CustomerAddressUpdatedBy
	u.CustomerAddressUpdatedAt = u.CustomerAddressUpdatedAt
	u.CustomerAddressDeletedBy = u.CustomerAddressDeletedBy
	u.CustomerAddressDeletedAt = u.CustomerAddressDeletedAt
}

func (u *CustomerAddress) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if u.CustomerID == 0 {
			return errors.New("Required Customer ID")
		}
		if u.CustomerAddressStreet == "" {
			return errors.New("Required User Code")
		}
		if u.AddressID == 0 {
			return errors.New("Required Email")
		}
		if u.CustomerAddressZipCode == "" {
			return errors.New("Username Wajib Diisi")
		}
		return nil
	}
}

func (u *CustomerAddress) SaveCustomerAddress(db *gorm.DB) (*CustomerAddress, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &CustomerAddress{}, err
	}
	return u, nil
}

func (p *CustomerAddress) FindAllCustomerAddress(db *gorm.DB) (*[]CustomerAddressData, error) {
	var err error
	productsubcategory := []CustomerAddressData{}
	err = db.Debug().Model(&CustomerAddressData{}).Limit(100).
		Select(`m_customer_address.cust_adrs_id,
				m_customer_address.cust_id,
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
				m_customer.cust_marital_status,
				m_customer.cust_image,
				m_customer_address.cust_adrs_street,
				m_customer_address.cust_adrs_rt,
				m_customer_address.cust_adrs_rw,
				m_customer_address.address_id,
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
				m_customer_address.cust_adrs_zip_code,
				m_customer_address.cust_adrs_ownership,
				m_customer_address.cust_adrs_length_stay,
				m_customer_address.cust_adrs_as_residential,
				m_customer_address.cust_adrs_status,
				m_customer_address.cust_adrs_created_by,
				m_customer_address.cust_adrs_created_at,
				m_customer_address.cust_adrs_updated_by,
				m_customer_address.cust_adrs_updated_at,
				m_customer_address.cust_adrs_deleted_by,
				m_customer_address.cust_adrs_deleted_at`).
		Joins("JOIN m_customer on m_customer_address.cust_id=m_customer.cust_id").
		Joins("JOIN m_address on m_customer_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Find(&productsubcategory).Error
	if err != nil {
		return &[]CustomerAddressData{}, err
	}
	return &productsubcategory, nil
}

func (p *CustomerAddress) FindAllCustomerAddressStatus(db *gorm.DB, status uint64) (*[]CustomerAddressData, error) {
	var err error
	productsubcategory := []CustomerAddressData{}
	err = db.Debug().Model(&CustomerAddressData{}).Limit(100).
		Select(`m_customer_address.cust_adrs_id,
				m_customer_address.cust_id,
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
				m_customer.cust_marital_status,
				m_customer.cust_image,
				m_customer_address.cust_adrs_street,
				m_customer_address.cust_adrs_rt,
				m_customer_address.cust_adrs_rw,
				m_customer_address.address_id,
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
				m_customer_address.cust_adrs_zip_code,
				m_customer_address.cust_adrs_ownership,
				m_customer_address.cust_adrs_length_stay,
				m_customer_address.cust_adrs_as_residential,
				m_customer_address.cust_adrs_status,
				m_customer_address.cust_adrs_created_by,
				m_customer_address.cust_adrs_created_at,
				m_customer_address.cust_adrs_updated_by,
				m_customer_address.cust_adrs_updated_at,
				m_customer_address.cust_adrs_deleted_by,
				m_customer_address.cust_adrs_deleted_at`).
		Joins("JOIN m_customer on m_customer_address.cust_id=m_customer.cust_id").
		Joins("JOIN m_address on m_customer_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("cust_adrs_status = ?", status).
		Find(&productsubcategory).Error
	if err != nil {
		return &[]CustomerAddressData{}, err
	}
	return &productsubcategory, nil
}

func (p *CustomerAddress) FindCustomerAddressByID(db *gorm.DB, pid uint64) (*CustomerAddressData, error) {
	var err error
	productsubcategory := CustomerAddressData{}
	err = db.Debug().Model(&CustomerAddressData{}).Limit(100).
		Select(`m_customer_address.cust_adrs_id,
				m_customer_address.cust_id,
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
				m_customer.cust_marital_status,
				m_customer.cust_image,
				m_customer_address.cust_adrs_street,
				m_customer_address.cust_adrs_rt,
				m_customer_address.cust_adrs_rw,
				m_customer_address.address_id,
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
				m_customer_address.cust_adrs_zip_code,
				m_customer_address.cust_adrs_ownership,
				m_customer_address.cust_adrs_length_stay,
				m_customer_address.cust_adrs_as_residential,
				m_customer_address.cust_adrs_status,
				m_customer_address.cust_adrs_created_by,
				m_customer_address.cust_adrs_created_at,
				m_customer_address.cust_adrs_updated_by,
				m_customer_address.cust_adrs_updated_at,
				m_customer_address.cust_adrs_deleted_by,
				m_customer_address.cust_adrs_deleted_at`).
		Joins("JOIN m_customer on m_customer_address.cust_id=m_customer.cust_id").
		Joins("JOIN m_address on m_customer_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("cust_adrs_id = ?", pid).
		Find(&productsubcategory).Error
	if err != nil {
		return &CustomerAddressData{}, err
	}
	return &productsubcategory, nil
}

func (p *CustomerAddress) FindCustomerAddressByPhone(db *gorm.DB, phone string) (*CustomerAddressData, error) {
	var err error
	productsubcategory := CustomerAddressData{}
	err = db.Debug().Model(&CustomerAddressData{}).Limit(100).
		Select(`m_customer_address.cust_adrs_id,
				m_customer_address.cust_id,
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
				m_customer.cust_marital_status,
				m_customer.cust_image,
				m_customer_user.cust_user_phone,
				m_customer_address.cust_adrs_street,
				m_customer_address.cust_adrs_rt,
				m_customer_address.cust_adrs_rw,
				m_customer_address.address_id,
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
				m_customer_address.cust_adrs_zip_code,
				m_customer_address.cust_adrs_ownership,
				m_customer_address.cust_adrs_length_stay,
				m_customer_address.cust_adrs_as_residential,
				m_customer_address.cust_adrs_status,
				m_customer_address.cust_adrs_created_by,
				m_customer_address.cust_adrs_created_at,
				m_customer_address.cust_adrs_updated_by,
				m_customer_address.cust_adrs_updated_at,
				m_customer_address.cust_adrs_deleted_by,
				m_customer_address.cust_adrs_deleted_at`).
		Joins("JOIN m_customer_user on m_customer_address.cust_id=m_customer_user.cust_id").
		Joins("JOIN m_customer on m_customer_address.cust_id=m_customer.cust_id").
		Joins("JOIN m_address on m_customer_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("m_customer_user.cust_user_phone = ?", phone).
		Find(&productsubcategory).Error
	if err != nil {
		return &CustomerAddressData{}, err
	}
	return &productsubcategory, nil
}

func (p *CustomerAddress) FindCustomerAddressStatusByID(db *gorm.DB, pid uint64, status uint64) (*CustomerAddressData, error) {
	var err error
	productsubcategory := CustomerAddressData{}
	err = db.Debug().Model(&CustomerAddressData{}).Limit(100).
		Select(`m_customer_address.cust_adrs_id,
				m_customer_address.cust_id,
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
				m_customer.cust_marital_status,
				m_customer.cust_image,
				m_customer_address.cust_adrs_street,
				m_customer_address.cust_adrs_rt,
				m_customer_address.cust_adrs_rw,
				m_customer_address.address_id,
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
				m_customer_address.cust_adrs_zip_code,
				m_customer_address.cust_adrs_ownership,
				m_customer_address.cust_adrs_length_stay,
				m_customer_address.cust_adrs_as_residential,
				m_customer_address.cust_adrs_status,
				m_customer_address.cust_adrs_created_by,
				m_customer_address.cust_adrs_created_at,
				m_customer_address.cust_adrs_updated_by,
				m_customer_address.cust_adrs_updated_at,
				m_customer_address.cust_adrs_deleted_by,
				m_customer_address.cust_adrs_deleted_at`).
		Joins("JOIN m_customer on m_customer_address.cust_id=m_customer.cust_id").
		Joins("JOIN m_address on m_customer_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("cust_adrs_id = ? AND cust_adrs_status = ?", pid, status).
		Find(&productsubcategory).Error
	if err != nil {
		return &CustomerAddressData{}, err
	}
	return &productsubcategory, nil
}

func (p *CustomerAddress) FindCustomerAddressByCustomerID(db *gorm.DB, pid uint64) (*CustomerAddressData, error) {
	var err error
	productsubcategory := CustomerAddressData{}
	err = db.Debug().Model(&CustomerAddressData{}).Limit(100).
		Select(`m_customer_address.cust_adrs_id,
				m_customer_address.cust_id,
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
				m_customer.cust_marital_status,
				m_customer.cust_image,
				m_customer_address.cust_adrs_street,
				m_customer_address.cust_adrs_rt,
				m_customer_address.cust_adrs_rw,
				m_customer_address.address_id,
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
				m_customer_address.cust_adrs_zip_code,
				m_customer_address.cust_adrs_ownership,
				m_customer_address.cust_adrs_length_stay,
				m_customer_address.cust_adrs_as_residential,
				m_customer_address.cust_adrs_status,
				m_customer_address.cust_adrs_created_by,
				m_customer_address.cust_adrs_created_at,
				m_customer_address.cust_adrs_updated_by,
				m_customer_address.cust_adrs_updated_at,
				m_customer_address.cust_adrs_deleted_by,
				m_customer_address.cust_adrs_deleted_at`).
		Joins("JOIN m_customer on m_customer_address.cust_id=m_customer.cust_id").
		Joins("JOIN m_address on m_customer_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("m_customer_address.cust_id = ?", pid).
		Find(&productsubcategory).Error
	if err != nil {
		return &CustomerAddressData{}, err
	}
	return &productsubcategory, nil
}

func (p *CustomerAddress) CountCustomerAddressByCustomerID(db *gorm.DB, pid uint64) (*CustomerAddressCount, error) {
	var err error
	productsubcategory := CustomerAddressCount{}
	err = db.Debug().Model(&CustomerAddressCount{}).
		Where("m_customer_address.cust_id = ?", pid).
		Count(&productsubcategory.CustomerAddressCountValue).
		Find(&productsubcategory).Error
	if err != nil {
		return &CustomerAddressCount{}, err
	}
	return &productsubcategory, nil
}

func (u *CustomerAddress) UpdateCustomerAddress(db *gorm.DB, uid uint64) (*CustomerAddress, error) {
	var err error
	db = db.Debug().Model(&CustomerAddress{}).Where("cust_adrs_id = ?", uid).Take(&CustomerAddress{}).UpdateColumns(
		map[string]interface{}{
			"cust_id":                  u.CustomerID,
			"cust_adrs_street":         u.CustomerAddressStreet,
			"cust_adrs_rt":             u.CustomerAddressRT,
			"cust_adrs_rw":             u.CustomerAddressRW,
			"address_id":               u.AddressID,
			"cust_adrs_zip_code":       u.CustomerAddressZipCode,
			"cust_adrs_ownership":      u.CustomerAddressOwnership,
			"cust_adrs_length_stay":    u.CustomerAddressLengthStay,
			"cust_adrs_as_residential": u.CustomerAddressAsResidential,
			"cust_adrs_status":         u.CustomerAddressStatus,
			"cust_adrs_updated_by":     u.CustomerAddressUpdatedBy,
			"cust_adrs_updated_at":     time.Now(),
		},
	)
	if db.Error != nil {
		return &CustomerAddress{}, db.Error
	}
	// This is the display the updated customeruser
	err = db.Debug().Model(&CustomerAddress{}).Where("cust_adrs_id = ?", uid).Take(&u).Error
	if err != nil {
		return &CustomerAddress{}, err
	}
	return u, nil
}

func (u *CustomerAddress) UpdateProfileCustomer(db *gorm.DB, uid uint64) (*CustomerAddress, error) {
	var err error
	db = db.Debug().Model(&CustomerAddress{}).Where("cust_id = ?", uid).Take(&CustomerAddress{}).UpdateColumns(
		map[string]interface{}{
			"cust_adrs_street":     u.CustomerAddressStreet,
			"cust_adrs_zip_code":   u.CustomerAddressZipCode,
			"cust_adrs_updated_by": u.CustomerAddressUpdatedBy,
			"cust_adrs_updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &CustomerAddress{}, db.Error
	}
	// This is the display the updated customeruser
	err = db.Debug().Model(&CustomerAddress{}).Where("cust_id = ?", uid).Take(&u).Error
	if err != nil {
		return &CustomerAddress{}, err
	}
	return u, nil
}

func (p *CustomerAddress) UpdateCustomerAddressStatus(db *gorm.DB, pid uint64) (*CustomerAddress, error) {

	var err error
	db = db.Debug().Model(&CustomerAddress{}).Where("cust_adrs_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"cust_adrs_status":     p.CustomerAddressStatus,
			"cust_adrs_updated_by": p.CustomerAddressUpdatedBy,
			"cust_adrs_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&CustomerAddress{}).Where("cust_adrs_id = ?", pid).Error
	if err != nil {
		return &CustomerAddress{}, err
	}
	return p, nil
}

func (p *CustomerAddress) UpdateCustomerAddressStatusDelete(db *gorm.DB, pid uint64) (*CustomerAddress, error) {
	var err error
	db = db.Debug().Model(&CustomerAddress{}).Where("cust_adrs_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"cust_adrs_status":     3,
			"cust_adrs_deleted_by": p.CustomerAddressDeletedBy,
			"cust_adrs_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&CustomerAddress{}).Where("cust_adrs_id = ?", pid).Error
	if err != nil {
		return &CustomerAddress{}, err
	}
	return p, nil
}

func (p *CustomerAddress) DeleteCustomerAddress(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&CustomerAddress{}).Where("cust_adrs_id = ?", pid).Take(&CustomerAddress{}).Delete(&CustomerAddress{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("CustomerAddress not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
