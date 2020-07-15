package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

//BASE CRUD
//====================================================================================================================================

type BibiteEmployee struct {
	BibiteEmployeeID         uint64     `gorm:"column:bbt_emp_id;primary_key;" json:"bbt_emp_id"`
	BibiteEmployeeCode       string     `gorm:"column:bbt_emp_code" json:"bbt_emp_code"`
	BibiteEmployeeName       string     `gorm:"column:bbt_emp_name" json:"bbt_emp_name"`
	BibiteEmployeeGender     string     `gorm:"column:bbt_emp_gender" json:"bbt_emp_gender"`
	BibiteEmployeeBirthPlace string     `gorm:"column:bbt_emp_birth_place" json:"bbt_emp_birth_place"`
	BibiteEmployeeBirthDate  string     `gorm:"column:bbt_emp_birth_date" json:"bbt_emp_birth_date"`
	BibiteEmployeeAddress    string     `gorm:"column:bbt_emp_address" json:"bbt_emp_address"`
	AddressID                uint64     `gorm:"column:address_id" json:"address_id"`
	BibiteEmployeePhoto      string     `gorm:"column:bbt_emp_photo" json:"bbt_emp_photo"`
	BibiteEmployeeStatus     uint64     `gorm:"column:bbt_emp_status;size:2" json:"bbt_emp_status"`
	BibiteEmployeeCreatedBy  string     `gorm:"column:bbt_emp_created_by;size:125" json:"bbt_emp_created_by"`
	BibiteEmployeeCreatedAt  time.Time  `gorm:"column:bbt_emp_created_at;default:CURRENT_TIMESTAMP" json:"bbt_emp_created_at"`
	BibiteEmployeeUpdatedBy  string     `gorm:"column:bbt_emp_updated_by;size:125" json:"bbt_emp_updated_by"`
	BibiteEmployeeUpdatedAt  *time.Time `gorm:"column:bbt_emp_updated_at;default:CURRENT_TIMESTAMP" json:"bbt_emp_created_at"`
	BibiteEmployeeDeletedBy  string     `gorm:"column:bbt_emp_deleted_by;size:125" json:"bbt_emp_deleted_by"`
	BibiteEmployeeDeletedAt  *time.Time `gorm:"column:bbt_emp_deleted_at;default:CURRENT_TIMESTAMP" json:"bbt_emp_deleted_at"`
}

type BibiteEmployeeData struct {
	BibiteEmployeeID         uint64     `gorm:"column:bbt_emp_id" json:"bbt_emp_id"`
	BibiteEmployeeCode       string     `gorm:"column:bbt_emp_code" json:"bbt_emp_code"`
	BibiteEmployeeName       string     `gorm:"column:bbt_emp_name" json:"bbt_emp_name"`
	BibiteEmployeeGender     string     `gorm:"column:bbt_emp_gender" json:"bbt_emp_gender"`
	BibiteEmployeeBirthPlace string     `gorm:"column:bbt_emp_birth_place" json:"bbt_emp_birth_place"`
	BibiteEmployeeBirthDate  string     `gorm:"column:bbt_emp_birth_date" json:"bbt_emp_birth_date"`
	BibiteEmployeeAddress    string     `gorm:"column:bbt_emp_address" json:"bbt_emp_address"`
	AddressID                uint64     `gorm:"column:address_id" json:"address_id"`
	CountryID                uint64     `gorm:"column:country_id" json:"country_id"`
	CountryName              string     `gorm:"column:country_name" json:"country_name"`
	ProvinceID               uint64     `gorm:"column:province_id" json:"province_id"`
	ProvinceName             string     `gorm:"column:province_name" json:"province_name"`
	RegencyID                uint64     `gorm:"column:regency_id" json:"regency_id"`
	RegencyName              string     `gorm:"column:regency_name" json:"regency_name"`
	DistrictID               uint64     `gorm:"column:district_id" json:"district_id"`
	DistrictName             string     `gorm:"column:district_name" json:"district_name"`
	VillageID                uint64     `gorm:"column:village_id" json:"village_id"`
	VillageName              string     `gorm:"column:village_name" json:"village_name"`
	BibiteEmployeePhoto      string     `gorm:"column:bbt_emp_photo" json:"bbt_emp_photo"`
	BibiteEmployeeStatus     uint64     `gorm:"column:bbt_emp_status" json:"bbt_emp_status"`
	BibiteEmployeeCreatedBy  string     `gorm:"column:bbt_emp_created_by" json:"bbt_emp_created_by"`
	BibiteEmployeeCreatedAt  time.Time  `gorm:"column:bbt_emp_created_at" json:"bbt_emp_created_at"`
	BibiteEmployeeUpdatedBy  string     `gorm:"column:bbt_emp_updated_by" json:"bbt_emp_updated_by"`
	BibiteEmployeeUpdatedAt  *time.Time `gorm:"column:bbt_emp_updated_at" json:"bbt_emp_updated_at"`
	BibiteEmployeeDeletedBy  string     `gorm:"column:bbt_emp_deleted_by" json:"bbt_emp_deleted_by"`
	BibiteEmployeeDeletedAt  *time.Time `gorm:"column:bbt_emp_deleted_at" json:"bbt_emp_deleted_at"`
}

type ResponseBibiteEmployees struct {
	Status  int                   `json:"status"`
	Message string                `json:"message"`
	Data    *[]BibiteEmployeeData `json:"data"`
}

type ResponseBibiteEmployee struct {
	Status  int                 `json:"status"`
	Message string              `json:"message"`
	Data    *BibiteEmployeeData `json:"data"`
}

type ResponseBibiteEmployeeCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (BibiteEmployee) TableName() string {
	return "m_bibite_employee"
}

func (BibiteEmployeeData) TableName() string {
	return "m_bibite_employee"
}

func (p *BibiteEmployee) Prepare() {
	p.BibiteEmployeeCode = p.BibiteEmployeeCode
	p.BibiteEmployeeName = p.BibiteEmployeeName
	p.BibiteEmployeeGender = p.BibiteEmployeeGender
	p.BibiteEmployeeBirthPlace = p.BibiteEmployeeBirthPlace
	p.BibiteEmployeeBirthDate = p.BibiteEmployeeBirthDate
	p.BibiteEmployeeAddress = p.BibiteEmployeeAddress
	p.AddressID = p.AddressID
	p.BibiteEmployeePhoto = p.BibiteEmployeePhoto
	p.BibiteEmployeeStatus = p.BibiteEmployeeStatus
	p.BibiteEmployeeCreatedBy = html.EscapeString(strings.TrimSpace(p.BibiteEmployeeCreatedBy))
	p.BibiteEmployeeCreatedAt = time.Now()
	p.BibiteEmployeeUpdatedBy = html.EscapeString(strings.TrimSpace(p.BibiteEmployeeUpdatedBy))
	p.BibiteEmployeeUpdatedAt = p.BibiteEmployeeUpdatedAt
	p.BibiteEmployeeDeletedBy = html.EscapeString(strings.TrimSpace(p.BibiteEmployeeDeletedBy))
	p.BibiteEmployeeDeletedAt = p.BibiteEmployeeDeletedAt
}

func (p *BibiteEmployee) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.BibiteEmployeeCode == "" {
			return errors.New("Required Country")
		}
		if p.BibiteEmployeeName == "" {
			return errors.New("Required Province")
		}
		if p.BibiteEmployeeGender == "" {
			return errors.New("Required Regency")
		}
		if p.BibiteEmployeeBirthPlace == "" {
			return errors.New("Required District")
		}
		if p.BibiteEmployeeBirthDate == "" {
			return errors.New("Required Village")
		}
		if p.BibiteEmployeeAddress == "" {
			return errors.New("Required Village")
		}
		if p.AddressID == 0 {
			return errors.New("Required Village")
		}
		if p.BibiteEmployeePhoto == "" {
			return errors.New("Required Village")
		}
		return nil
	}
}

func (p *BibiteEmployee) SaveBibiteEmployee(db *gorm.DB) (*BibiteEmployee, error) {
	var err error
	err = db.Debug().Model(&BibiteEmployee{}).Create(&p).Error
	if err != nil {
		return &BibiteEmployee{}, err
	}
	return p, nil
}

func (p *BibiteEmployee) FindAllBibiteEmployee(db *gorm.DB) (*[]BibiteEmployeeData, error) {
	var err error
	address := []BibiteEmployeeData{}
	err = db.Debug().Model(&BibiteEmployeeData{}).Limit(100).
		Select(`m_bibite_employee.bbt_emp_id,
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
				m_bibite_employee.bbt_emp_status,
				m_bibite_employee.bbt_emp_created_by,
				m_bibite_employee.bbt_emp_created_at at time zone 'Asia/Jakarta' as address_deleted_at,
				m_bibite_employee.bbt_emp_updated_by,
				m_bibite_employee.bbt_emp_updated_at at time zone 'Asia/Jakarta' as address_deleted_at,
				m_bibite_employee.bbt_emp_deleted_by,
				m_bibite_employee.bbt_emp_deleted_at at time zone 'Asia/Jakarta' as address_deleted_at`).
		Joins("JOIN m_address on m_bibite_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Find(&address).Error
	if err != nil {
		return &[]BibiteEmployeeData{}, err
	}
	return &address, nil
}

func (p *BibiteEmployee) FindAllBibiteEmployeeStatus(db *gorm.DB, status uint64) (*[]BibiteEmployeeData, error) {
	var err error
	address := []BibiteEmployeeData{}
	err = db.Debug().Model(&BibiteEmployeeData{}).Limit(100).
		Select(`m_bibite_employee.bbt_emp_id,
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
				m_bibite_employee.bbt_emp_status,
				m_bibite_employee.bbt_emp_created_by,
				m_bibite_employee.bbt_emp_created_at at time zone 'Asia/Jakarta' as address_deleted_at,
				m_bibite_employee.bbt_emp_updated_by,
				m_bibite_employee.bbt_emp_updated_at at time zone 'Asia/Jakarta' as address_deleted_at,
				m_bibite_employee.bbt_emp_deleted_by,
				m_bibite_employee.bbt_emp_deleted_at at time zone 'Asia/Jakarta' as address_deleted_at`).
		Joins("JOIN m_address on m_bibite_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("bbt_emp_status = ?", status).
		Find(&address).Error
	if err != nil {
		return &[]BibiteEmployeeData{}, err
	}
	return &address, nil
}

func (p *BibiteEmployee) FindBibiteEmployeeDataByID(db *gorm.DB, pid uint64) (*BibiteEmployee, error) {
	var err error
	err = db.Debug().Model(&BibiteEmployee{}).Where("bbt_emp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &BibiteEmployee{}, err
	}
	return p, nil
}

func (p *BibiteEmployee) FindBibiteEmployeeByID(db *gorm.DB, pid uint64) (*BibiteEmployeeData, error) {
	var err error
	address := BibiteEmployeeData{}
	err = db.Debug().Model(&BibiteEmployeeData{}).Limit(100).
		Select(`m_bibite_employee.bbt_emp_id,
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
				m_bibite_employee.bbt_emp_status,
				m_bibite_employee.bbt_emp_created_by,
				m_bibite_employee.bbt_emp_created_at at time zone 'Asia/Jakarta' as address_deleted_at,
				m_bibite_employee.bbt_emp_updated_by,
				m_bibite_employee.bbt_emp_updated_at at time zone 'Asia/Jakarta' as address_deleted_at,
				m_bibite_employee.bbt_emp_deleted_by,
				m_bibite_employee.bbt_emp_deleted_at at time zone 'Asia/Jakarta' as address_deleted_at`).
		Joins("JOIN m_address on m_bibite_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("bbt_emp_id = ?", pid).
		Find(&address).Error
	if err != nil {
		return &BibiteEmployeeData{}, err
	}
	return &address, nil
}

func (p *BibiteEmployee) FindBibiteEmployeeStatusByID(db *gorm.DB, pid uint64, status uint64) (*BibiteEmployeeData, error) {
	var err error
	address := BibiteEmployeeData{}
	err = db.Debug().Model(&BibiteEmployeeData{}).Limit(100).
		Select(`m_bibite_employee.bbt_emp_id,
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
				m_bibite_employee.bbt_emp_status,
				m_bibite_employee.bbt_emp_created_by,
				m_bibite_employee.bbt_emp_created_at at time zone 'Asia/Jakarta' as address_deleted_at,
				m_bibite_employee.bbt_emp_updated_by,
				m_bibite_employee.bbt_emp_updated_at at time zone 'Asia/Jakarta' as address_deleted_at,
				m_bibite_employee.bbt_emp_deleted_by,
				m_bibite_employee.bbt_emp_deleted_at at time zone 'Asia/Jakarta' as address_deleted_at`).
		Joins("JOIN m_address on m_bibite_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("bbt_emp_id = ? AND bbt_emp_status = ?", pid, status).
		Find(&address).Error
	if err != nil {
		return &BibiteEmployeeData{}, err
	}
	return &address, nil
}

func (p *BibiteEmployee) UpdateBibiteEmployee(db *gorm.DB, pid uint64) (*BibiteEmployee, error) {
	var err error
	db = db.Debug().Model(&BibiteEmployee{}).Where("bbt_emp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_emp_code":        p.BibiteEmployeeCode,
			"bbt_emp_name":        p.BibiteEmployeeName,
			"bbt_emp_gender":      p.BibiteEmployeeGender,
			"bbt_emp_birth_place": p.BibiteEmployeeBirthPlace,
			"bbt_emp_birth_date":  p.BibiteEmployeeBirthDate,
			"bbt_emp_address":     p.BibiteEmployeeAddress,
			"bbt_emp_photo":       p.BibiteEmployeePhoto,
			"bbt_emp_status":      p.BibiteEmployeeStatus,
			"bbt_emp_updated_by":  p.BibiteEmployeeUpdatedBy,
			"bbt_emp_updated_at":  time.Now(),
		},
	)
	err = db.Debug().Model(&BibiteEmployee{}).Where("bbt_emp_id = ?", pid).Error
	if err != nil {
		return &BibiteEmployee{}, err
	}
	return p, nil
}

func (p *BibiteEmployee) UpdateBibiteEmployeeStatus(db *gorm.DB, pid uint64) (*BibiteEmployee, error) {
	var err error
	db = db.Debug().Model(&BibiteEmployee{}).Where("bbt_emp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_emp_status":     p.BibiteEmployeeStatus,
			"bbt_emp_updated_by": p.BibiteEmployeeUpdatedBy,
			"bbt_emp_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&BibiteEmployee{}).Where("bbt_emp_id = ?", pid).Error
	if err != nil {
		return &BibiteEmployee{}, err
	}
	return p, nil
}

func (p *BibiteEmployee) UpdateBibiteEmployeeStatusOnly(db *gorm.DB, pid uint64) (*BibiteEmployee, error) {
	var err error
	db = db.Debug().Model(&BibiteEmployee{}).Where("bbt_emp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_emp_status": p.BibiteEmployeeStatus,
		},
	)
	err = db.Debug().Model(&BibiteEmployee{}).Where("bbt_emp_id = ?", pid).Error
	if err != nil {
		return &BibiteEmployee{}, err
	}
	return p, nil
}

func (p *BibiteEmployee) UpdateBibiteEmployeeStatusDelete(db *gorm.DB, pid uint64) (*BibiteEmployee, error) {
	var err error
	db = db.Debug().Model(&BibiteEmployee{}).Where("bbt_emp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_emp_status":     3,
			"bbt_emp_deleted_by": p.BibiteEmployeeDeletedBy,
			"bbt_emp_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&BibiteEmployee{}).Where("bbt_emp_id = ?", pid).Error
	if err != nil {
		return &BibiteEmployee{}, err
	}
	return p, nil
}

func (p *BibiteEmployee) DeleteBibiteEmployee(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&BibiteEmployee{}).Where("bbt_emp_id = ?", pid).Take(&BibiteEmployee{}).Delete(&BibiteEmployee{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("BibiteEmployee not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//BIBITE LOGIN
//====================================================================================================================================

type BibiteLogin struct {
	BibiteType     string `json:"bbt_type"`
	BibiteUsername string `json:"bbt_username"`
	BibitePassword string `json:"bbt_password"`
}
type BibitePassword struct {
	BibiteAccountOldPassword string `json:"bbt_account_old_password"`
	BibiteAccountNewPassword string `json:"bbt_account_new_password"`
}

type ResponseBibiteLoginAuthenticationData struct {
	Token          string                          `json:"token"`
	Status         uint64                          `json:"status"`
	Name           string                          `json:"name"`
	Username       string                          `json:"username"`
	Photo          string                          `json:"photo"`
	IPAddress      string                          `json:"ip_address"`
	LastLogin      time.Time                       `json:"last_login"`
	GroupName      string                          `json:"group_name"`
	GroupRole      string                          `json:"group_role"`
	GroupPrivilege []BibiteGroupPrivilegeDataLogin `json:"group_privilege"`
}

type ResponseBibiteLoginAuthentication struct {
	Status  int                                   `json:"status"`
	Message string                                `json:"message"`
	Data    ResponseBibiteLoginAuthenticationData `json:"data"`
}

func (p *BibiteLogin) Validate(action string) error {
	switch strings.ToLower(action) {

	case "login":
		if p.BibiteUsername == "" {
			return errors.New("Username Wajib Diisi")
		}
		if p.BibitePassword == "" {
			return errors.New("Password Wajib Diisi")
		}
		return nil

	default:
		if p.BibiteUsername == "" {
			return errors.New("Username Wajib Diisi")
		}
		return nil
	}
}

//ADDITIONAL
//====================================================================================================================================

type BibiteProfileWeb struct {
	BibiteProfileEmail     string `gorm:"column:bbt_profile_email" json:"bbt_profile_email"`
	BibiteProfileUsername  string `gorm:"column:bbt_profile_username" json:"bbt_profile_username"`
	BibiteProfilePhoneCode string `gorm:"column:bbt_profile_phone_code" json:"bbt_profile_phone_code"`
	BibiteProfilePhone     string `gorm:"column:bbt_profile_phone" json:"bbt_profile_phone"`
}

func (p *BibiteProfileWeb) Prepare() {
	p.BibiteProfileEmail = p.BibiteProfileEmail
	p.BibiteProfileUsername = p.BibiteProfileUsername
	p.BibiteProfilePhoneCode = p.BibiteProfilePhoneCode
	p.BibiteProfilePhone = p.BibiteProfilePhone
}
