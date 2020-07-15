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

type BibiteEmployeeTemp struct {
	BibiteEmployeeTempID           uint64    `gorm:"column:bbt_emp_temp_id;primary_key;" json:"bbt_emp_temp_id"`
	BibiteEmployeeID               uint64    `gorm:"column:bbt_emp_id" json:"bbt_emp_id"`
	BibiteEmployeeTempCode         string    `gorm:"column:bbt_emp_temp_code" json:"bbt_emp_temp_code"`
	BibiteEmployeeTempName         string    `gorm:"column:bbt_emp_temp_name" json:"bbt_emp_temp_name"`
	BibiteEmployeeTempGender       string    `gorm:"column:bbt_emp_temp_gender" json:"bbt_emp_temp_gender"`
	BibiteEmployeeTempBirthPlace   string    `gorm:"column:bbt_emp_temp_birth_place" json:"bbt_emp_temp_birth_place"`
	BibiteEmployeeTempBirthDate    string    `gorm:"column:bbt_emp_temp_birth_date" json:"bbt_emp_temp_birth_date"`
	BibiteEmployeeTempAddress      string    `gorm:"column:bbt_emp_temp_address" json:"bbt_emp_temp_address"`
	AddressTempID                  uint64    `gorm:"column:address_temp_id" json:"address_temp_id"`
	BibiteEmployeeTempPhoto        string    `gorm:"column:bbt_emp_temp_photo" json:"bbt_emp_temp_photo"`
	BibiteEmployeeTempReason       string    `gorm:"column:bbt_emp_temp_reason" json:"bbt_emp_temp_reason"`
	BibiteEmployeeTempStatus       uint64    `gorm:"column:bbt_emp_temp_status;size:2" json:"bbt_emp_temp_status"`
	BibiteEmployeeTempActionBefore uint64    `gorm:"column:bbt_emp_temp_action_before;size:2" json:"bbt_emp_temp_action_before"`
	BibiteEmployeeTempCreatedBy    string    `gorm:"column:bbt_emp_temp_created_by;size:125" json:"bbt_emp_temp_created_by"`
	BibiteEmployeeTempCreatedAt    time.Time `gorm:"column:bbt_emp_temp_created_at;default:CURRENT_TIMESTAMP" json:"bbt_emp_temp_created_at"`
}

type BibiteEmployeeTempPend struct {
	BibiteEmployeeTempID           uint64    `gorm:"column:bbt_emp_temp_id;primary_key;" json:"bbt_emp_temp_id"`
	BibiteEmployeeID               *uint64   `gorm:"column:bbt_emp_id" json:"bbt_emp_id"`
	BibiteEmployeeTempCode         string    `gorm:"column:bbt_emp_temp_code" json:"bbt_emp_temp_code"`
	BibiteEmployeeTempName         string    `gorm:"column:bbt_emp_temp_name" json:"bbt_emp_temp_name"`
	BibiteEmployeeTempGender       string    `gorm:"column:bbt_emp_temp_gender" json:"bbt_emp_temp_gender"`
	BibiteEmployeeTempBirthPlace   string    `gorm:"column:bbt_emp_temp_birth_place" json:"bbt_emp_temp_birth_place"`
	BibiteEmployeeTempBirthDate    string    `gorm:"column:bbt_emp_temp_birth_date" json:"bbt_emp_temp_birth_date"`
	BibiteEmployeeTempAddress      string    `gorm:"column:bbt_emp_temp_address" json:"bbt_emp_temp_address"`
	AddressTempID                  uint64    `gorm:"column:address_temp_id" json:"address_temp_id"`
	BibiteEmployeeTempPhoto        string    `gorm:"column:bbt_emp_temp_photo" json:"bbt_emp_temp_photo"`
	BibiteEmployeeTempReason       string    `gorm:"column:bbt_emp_temp_reason" json:"bbt_emp_temp_reason"`
	BibiteEmployeeTempStatus       uint64    `gorm:"column:bbt_emp_temp_status;size:2" json:"bbt_emp_temp_status"`
	BibiteEmployeeTempActionBefore uint64    `gorm:"column:bbt_emp_temp_action_before;size:2" json:"bbt_emp_temp_action_before"`
	BibiteEmployeeTempCreatedBy    string    `gorm:"column:bbt_emp_temp_created_by;size:125" json:"bbt_emp_temp_created_by"`
	BibiteEmployeeTempCreatedAt    time.Time `gorm:"column:bbt_emp_temp_created_at;default:CURRENT_TIMESTAMP" json:"bbt_emp_temp_created_at"`
}

type BibiteEmployeeTempData struct {
	BibiteEmployeeTempID           uint64    `gorm:"column:bbt_emp_temp_id" json:"bbt_emp_temp_id"`
	BibiteEmployeeTempCode         string    `gorm:"column:bbt_emp_temp_code" json:"bbt_emp_temp_code"`
	BibiteEmployeeTempName         string    `gorm:"column:bbt_emp_temp_name" json:"bbt_emp_temp_name"`
	BibiteEmployeeTempGender       string    `gorm:"column:bbt_emp_temp_gender" json:"bbt_emp_temp_gender"`
	BibiteEmployeeTempBirthPlace   string    `gorm:"column:bbt_emp_temp_birth_place" json:"bbt_emp_temp_birth_place"`
	BibiteEmployeeTempBirthDate    string    `gorm:"column:bbt_emp_temp_birth_date" json:"bbt_emp_temp_birth_date"`
	BibiteEmployeeTempAddress      string    `gorm:"column:bbt_emp_temp_address" json:"bbt_emp_temp_address"`
	AddressTempID                  uint64    `gorm:"column:address_temp_id" json:"address_temp_id"`
	CountryTempID                  uint64    `gorm:"column:country_temp_id" json:"country_temp_id"`
	CountryTempName                string    `gorm:"column:country_temp_name" json:"country_temp_name"`
	ProvinceTempID                 uint64    `gorm:"column:province_temp_id" json:"province_temp_id"`
	ProvinceTempName               string    `gorm:"column:province_temp_name" json:"province_temp_name"`
	RegencyTempID                  uint64    `gorm:"column:regency_temp_id" json:"regency_temp_id"`
	RegencyTempName                string    `gorm:"column:regency_temp_name" json:"regency_temp_name"`
	DistrictTempID                 uint64    `gorm:"column:district_temp_id" json:"district_temp_id"`
	DistrictTempName               string    `gorm:"column:district_temp_name" json:"district_temp_name"`
	VillageTempID                  uint64    `gorm:"column:village_temp_id" json:"village_temp_id"`
	VillageTempName                string    `gorm:"column:village_temp_name" json:"village_temp_name"`
	BibiteEmployeeTempPhoto        string    `gorm:"column:bbt_emp_temp_photo" json:"bbt_emp_temp_photo"`
	BibiteEmployeeTempReason       string    `gorm:"column:bbt_emp_temp_reason" json:"bbt_emp_temp_reason"`
	BibiteEmployeeTempStatus       uint64    `gorm:"column:bbt_emp_temp_status;size:2" json:"bbt_emp_temp_status"`
	BibiteEmployeeTempActionBefore uint64    `gorm:"column:bbt_emp_temp_action_before;size:2" json:"bbt_emp_temp_action_before"`
	BibiteEmployeeTempCreatedBy    string    `gorm:"column:bbt_emp_temp_created_by;size:125" json:"bbt_emp_temp_created_by"`
	BibiteEmployeeTempCreatedAt    time.Time `gorm:"column:bbt_emp_temp_created_at" json:"bbt_emp_temp_created_at"`
	BibiteEmployeeID               uint64    `gorm:"column:bbt_emp_id" json:"bbt_emp_id"`
	BibiteEmployeeCode             string    `gorm:"column:bbt_emp_code" json:"bbt_emp_code"`
	BibiteEmployeeName             string    `gorm:"column:bbt_emp_name" json:"bbt_emp_name"`
	BibiteEmployeeGender           string    `gorm:"column:bbt_emp_gender" json:"bbt_emp_gender"`
	BibiteEmployeeBirthPlace       string    `gorm:"column:bbt_emp_birth_place" json:"bbt_emp_birth_place"`
	BibiteEmployeeBirthDate        string    `gorm:"column:bbt_emp_birth_date" json:"bbt_emp_birth_date"`
	BibiteEmployeeAddress          string    `gorm:"column:bbt_emp_address" json:"bbt_emp_address"`
	AddressID                      uint64    `gorm:"column:address_id" json:"address_id"`
	CountryID                      uint64    `gorm:"column:country_id" json:"country_id"`
	CountryName                    string    `gorm:"column:country_name" json:"country_name"`
	ProvinceID                     uint64    `gorm:"column:province_id" json:"province_id"`
	ProvinceName                   string    `gorm:"column:province_name" json:"province_name"`
	RegencyID                      uint64    `gorm:"column:regency_id" json:"regency_id"`
	RegencyName                    string    `gorm:"column:regency_name" json:"regency_name"`
	DistrictID                     uint64    `gorm:"column:district_id" json:"district_id"`
	DistrictName                   string    `gorm:"column:district_name" json:"district_name"`
	VillageID                      uint64    `gorm:"column:village_id" json:"village_id"`
	VillageName                    string    `gorm:"column:village_name" json:"village_name"`
	BibiteEmployeePhoto            string    `gorm:"column:bbt_emp_photo" json:"bbt_emp_photo"`
	BibiteEmployeeStatus           uint64    `gorm:"column:bbt_emp_status" json:"bbt_emp_status"`
	BibiteEmployeeCreatedBy        string    `gorm:"column:bbt_emp_created_by" json:"bbt_emp_created_by"`
	BibiteEmployeeCreatedAt        time.Time `gorm:"column:bbt_emp_created_at" json:"bbt_emp_created_at"`
	BibiteEmployeeUpdatedBy        string    `gorm:"column:bbt_emp_updated_by" json:"bbt_emp_updated_by"`
	BibiteEmployeeUpdatedAt        time.Time `gorm:"column:bbt_emp_updated_at" json:"bbt_emp_updated_at"`
	BibiteEmployeeDeletedBy        string    `gorm:"column:bbt_emp_deleted_by" json:"bbt_emp_deleted_by"`
	BibiteEmployeeDeletedAt        time.Time `gorm:"column:bbt_emp_deleted_at" json:"bbt_emp_deleted_at"`
}

type ResponseBibiteEmployeeTemps struct {
	Status  int                       `json:"status"`
	Message string                    `json:"message"`
	Data    *[]BibiteEmployeeTempData `json:"data"`
}

type ResponseBibiteEmployeeTemp struct {
	Status  int                     `json:"status"`
	Message string                  `json:"message"`
	Data    *BibiteEmployeeTempData `json:"data"`
}

type ResponseBibiteEmployeeTempsPend struct {
	Status  int                       `json:"status"`
	Message string                    `json:"message"`
	Data    *[]BibiteEmployeeTempPend `json:"data"`
}

type ResponseBibiteEmployeeTempPend struct {
	Status  int                     `json:"status"`
	Message string                  `json:"message"`
	Data    *BibiteEmployeeTempPend `json:"data"`
}

type ResponseBibiteEmployeeTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (BibiteEmployeeTemp) TableName() string {
	return "m_bibite_employee_temp"
}

func (BibiteEmployeeTempPend) TableName() string {
	return "m_bibite_employee_temp"
}

func (BibiteEmployeeTempData) TableName() string {
	return "m_bibite_employee_temp"
}

func (p *BibiteEmployeeTemp) Prepare() {
	p.BibiteEmployeeID = p.BibiteEmployeeID
	p.BibiteEmployeeTempCode = p.BibiteEmployeeTempCode
	p.BibiteEmployeeTempName = p.BibiteEmployeeTempName
	p.BibiteEmployeeTempGender = p.BibiteEmployeeTempGender
	p.BibiteEmployeeTempBirthPlace = p.BibiteEmployeeTempBirthPlace
	p.BibiteEmployeeTempBirthDate = p.BibiteEmployeeTempBirthDate
	p.BibiteEmployeeTempAddress = p.BibiteEmployeeTempAddress
	p.AddressTempID = p.AddressTempID
	p.BibiteEmployeeTempPhoto = p.BibiteEmployeeTempPhoto
	p.BibiteEmployeeTempReason = p.BibiteEmployeeTempReason
	p.BibiteEmployeeTempStatus = p.BibiteEmployeeTempStatus
	p.BibiteEmployeeTempCreatedBy = html.EscapeString(strings.TrimSpace(p.BibiteEmployeeTempCreatedBy))
	p.BibiteEmployeeTempCreatedAt = time.Now()
}

func (p *BibiteEmployeeTempPend) Prepare() {
	p.BibiteEmployeeID = nil
	p.BibiteEmployeeTempCode = p.BibiteEmployeeTempCode
	p.BibiteEmployeeTempName = p.BibiteEmployeeTempName
	p.BibiteEmployeeTempGender = p.BibiteEmployeeTempGender
	p.BibiteEmployeeTempBirthPlace = p.BibiteEmployeeTempBirthPlace
	p.BibiteEmployeeTempBirthDate = p.BibiteEmployeeTempBirthDate
	p.BibiteEmployeeTempAddress = p.BibiteEmployeeTempAddress
	p.AddressTempID = p.AddressTempID
	p.BibiteEmployeeTempPhoto = p.BibiteEmployeeTempPhoto
	p.BibiteEmployeeTempReason = p.BibiteEmployeeTempReason
	p.BibiteEmployeeTempStatus = p.BibiteEmployeeTempStatus
	p.BibiteEmployeeTempCreatedBy = html.EscapeString(strings.TrimSpace(p.BibiteEmployeeTempCreatedBy))
	p.BibiteEmployeeTempCreatedAt = time.Now()
}

func (p *BibiteEmployeeTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	case "insertupdate":
		if p.BibiteEmployeeTempCode == "" {
			return errors.New("Negara Wajib diisi")
		}
		if p.BibiteEmployeeTempName == "" {
			return errors.New("Provinsi Wajib diisi")
		}
		if p.BibiteEmployeeTempGender == "" {
			return errors.New("Kabupaten / Kota Wajib diisi")
		}
		if p.BibiteEmployeeTempBirthPlace == "" {
			return errors.New("Kecamatan Wajib diisi")
		}
		if p.BibiteEmployeeTempBirthDate == "" {
			return errors.New("Kelurahan Wajib diisi")
		}
		if p.BibiteEmployeeTempAddress == "" {
			return errors.New("Kelurahan Wajib diisi")
		}
		if p.AddressTempID == 0 {
			return errors.New("Kelurahan Wajib diisi")
		}
		if p.BibiteEmployeeTempPhoto == "" {
			return errors.New("Kelurahan Wajib diisi")
		}
		if p.BibiteEmployeeTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.BibiteEmployeeTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	case "reason":
		if p.BibiteEmployeeTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.BibiteEmployeeTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	default:
		if p.BibiteEmployeeTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil
	}
}

func (p *BibiteEmployeeTemp) SaveBibiteEmployeeTemp(db *gorm.DB) (*BibiteEmployeeTemp, error) {
	var err error
	err = db.Debug().Model(&BibiteEmployeeTemp{}).Create(&p).Error
	if err != nil {
		return &BibiteEmployeeTemp{}, err
	}
	return p, nil
}

func (p *BibiteEmployeeTempPend) SaveBibiteEmployeeTempPend(db *gorm.DB) (*BibiteEmployeeTempPend, error) {
	var err error
	err = db.Debug().Model(&BibiteEmployeeTempPend{}).Create(&p).Error
	if err != nil {
		return &BibiteEmployeeTempPend{}, err
	}
	return p, nil
}

func (p *BibiteEmployeeTemp) FindAllBibiteEmployeeTemp(db *gorm.DB) (*[]BibiteEmployeeTempPend, error) {
	var err error
	bibiteemployee := []BibiteEmployeeTempPend{}
	err = db.Debug().Model(&BibiteEmployeeTempPend{}).Limit(100).
		Select(`m_bibite_employee_temp.bbt_emp_temp_id,
				m_bibite_employee_temp.bbt_emp_temp_code,
				m_bibite_employee_temp.bbt_emp_temp_name,
				m_bibite_employee_temp.bbt_emp_temp_gender,
				m_bibite_employee_temp.bbt_emp_temp_birth_place,
				m_bibite_employee_temp.bbt_emp_temp_birth_date,
				m_bibite_employee_temp.bbt_emp_temp_address,
				m_bibite_employee_temp.address_temp_id,
				m_bibite_employee_temp.bbt_emp_temp_photo,
				m_bibite_employee_temp.bbt_emp_temp_reason,
				m_bibite_employee_temp.bbt_emp_temp_status,
				m_bibite_employee_temp.bbt_emp_temp_action_before,
				m_bibite_employee_temp.bbt_emp_temp_created_by,
				m_bibite_employee_temp.bbt_emp_temp_created_at at time zone 'Asia/Jakarta' as address_deleted_at`).
		Find(&bibiteemployee).Error
	if err != nil {
		return &[]BibiteEmployeeTempPend{}, err
	}
	return &bibiteemployee, nil
}

func (p *BibiteEmployeeTemp) FindAllBibiteEmployeeTempStatusPendingActive(db *gorm.DB) (*[]BibiteEmployeeTempPend, error) {
	var err error
	bibiteemployee := []BibiteEmployeeTempPend{}
	err = db.Debug().Model(&BibiteEmployeeTempPend{}).Limit(100).
		Select(`m_bibite_employee_temp.bbt_emp_temp_id,
				m_bibite_employee_temp.bbt_emp_temp_code,
				m_bibite_employee_temp.bbt_emp_temp_name,
				m_bibite_employee_temp.bbt_emp_temp_gender,
				m_bibite_employee_temp.bbt_emp_temp_birth_place,
				m_bibite_employee_temp.bbt_emp_temp_birth_date,
				m_bibite_employee_temp.bbt_emp_temp_address,
				m_bibite_employee_temp.address_temp_id,
				m_bibite_employee_temp.bbt_emp_temp_photo,
				m_bibite_employee_temp.bbt_emp_temp_reason,
				m_bibite_employee_temp.bbt_emp_temp_status,
				m_bibite_employee_temp.bbt_emp_temp_action_before,
				m_bibite_employee_temp.bbt_emp_temp_created_by,
				m_bibite_employee_temp.bbt_emp_temp_created_at at time zone 'Asia/Jakarta' as address_deleted_at`).
		Where("bbt_emp_temp_status = ?", 11).
		Find(&bibiteemployee).Error
	if err != nil {
		return &[]BibiteEmployeeTempPend{}, err
	}
	return &bibiteemployee, nil
}

func (p *BibiteEmployeeTemp) FindAllBibiteEmployeeTempStatus(db *gorm.DB, status uint64) (*[]BibiteEmployeeTempData, error) {
	var err error
	bibiteemployee := []BibiteEmployeeTempData{}
	err = db.Debug().Model(&BibiteEmployeeTempData{}).Limit(100).
		Select(`m_bibite_employee_temp.bbt_emp_temp_id,
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
				m_bibite_employee_temp.bbt_emp_temp_status,
				m_bibite_employee_temp.bbt_emp_temp_action_before,
				m_bibite_employee_temp.bbt_emp_temp_created_by,
				m_bibite_employee_temp.bbt_emp_temp_created_at at time zone 'Asia/Jakarta' as address_deleted_at,
				m_bibite_employee.bbt_emp_id,
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
		Joins("JOIN m_address_temp on m_bibite_employee_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_bibite_employee on m_bibite_employee_temp.bbt_emp_id=m_bibite_employee.bbt_emp_id").
		Joins("JOIN m_address on m_bibite_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("bbt_emp_temp_status = ?", status).
		Find(&bibiteemployee).Error
	if err != nil {
		return &[]BibiteEmployeeTempData{}, err
	}
	return &bibiteemployee, nil
}

func (p *BibiteEmployeeTemp) FindBibiteEmployeeTempDataByID(db *gorm.DB, pid uint64) (*BibiteEmployeeTemp, error) {
	var err error
	err = db.Debug().Model(&BibiteEmployeeTemp{}).Where("bbt_emp_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &BibiteEmployeeTemp{}, err
	}
	return p, nil
}

func (p *BibiteEmployeeTemp) FindBibiteEmployeeTempByIDPendingActive(db *gorm.DB, pid uint64) (*BibiteEmployeeTempPend, error) {
	var err error
	bibiteemployee := BibiteEmployeeTempPend{}
	err = db.Debug().Model(&BibiteEmployeeTempPend{}).Limit(100).
		Select(`m_bibite_employee_temp.bbt_emp_temp_id,
				m_bibite_employee_temp.bbt_emp_temp_code,
				m_bibite_employee_temp.bbt_emp_temp_name,
				m_bibite_employee_temp.bbt_emp_temp_gender,
				m_bibite_employee_temp.bbt_emp_temp_birth_place,
				m_bibite_employee_temp.bbt_emp_temp_birth_date,
				m_bibite_employee_temp.bbt_emp_temp_address,
				m_bibite_employee_temp.address_temp_id,
				m_bibite_employee_temp.bbt_emp_temp_photo,
				m_bibite_employee_temp.bbt_emp_temp_reason,
				m_bibite_employee_temp.bbt_emp_temp_status,
				m_bibite_employee_temp.bbt_emp_temp_action_before,
				m_bibite_employee_temp.bbt_emp_temp_created_by,
				m_bibite_employee_temp.bbt_emp_temp_created_at at time zone 'Asia/Jakarta' as address_deleted_at`).
		Where("bbt_emp_temp_id = ?", pid).
		Find(&bibiteemployee).Error
	if err != nil {
		return &BibiteEmployeeTempPend{}, err
	}
	return &bibiteemployee, nil
}

func (p *BibiteEmployeeTemp) FindBibiteEmployeeTempStatusByIDPendingActive(db *gorm.DB, pid uint64) (*BibiteEmployeeTempPend, error) {
	var err error
	bibiteemployee := BibiteEmployeeTempPend{}
	err = db.Debug().Model(&BibiteEmployeeTempPend{}).Limit(100).
		Select(`m_bibite_employee_temp.bbt_emp_temp_id,
				m_bibite_employee_temp.bbt_emp_temp_code,
				m_bibite_employee_temp.bbt_emp_temp_name,
				m_bibite_employee_temp.bbt_emp_temp_gender,
				m_bibite_employee_temp.bbt_emp_temp_birth_place,
				m_bibite_employee_temp.bbt_emp_temp_birth_date,
				m_bibite_employee_temp.bbt_emp_temp_address,
				m_bibite_employee_temp.address_temp_id,
				m_bibite_employee_temp.bbt_emp_temp_photo,
				m_bibite_employee_temp.bbt_emp_temp_reason,
				m_bibite_employee_temp.bbt_emp_temp_status,
				m_bibite_employee_temp.bbt_emp_temp_action_before,
				m_bibite_employee_temp.bbt_emp_temp_created_by,
				m_bibite_employee_temp.bbt_emp_temp_created_at at time zone 'Asia/Jakarta' as address_deleted_at`).
		Where("bbt_emp_temp_id = ? AND bbt_emp_temp_status = ?", pid, 11).
		Find(&bibiteemployee).Error
	if err != nil {
		return &BibiteEmployeeTempPend{}, err
	}
	return &bibiteemployee, nil
}

func (p *BibiteEmployeeTemp) FindBibiteEmployeeTempByID(db *gorm.DB, pid uint64) (*BibiteEmployeeTempData, error) {
	var err error
	bibiteemployee := BibiteEmployeeTempData{}
	err = db.Debug().Model(&BibiteEmployeeTempData{}).Limit(100).
		Select(`m_bibite_employee_temp.bbt_emp_temp_id,
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
				m_bibite_employee_temp.bbt_emp_temp_status,
				m_bibite_employee_temp.bbt_emp_temp_action_before,
				m_bibite_employee_temp.bbt_emp_temp_created_by,
				m_bibite_employee_temp.bbt_emp_temp_created_at at time zone 'Asia/Jakarta' as address_deleted_at,
				m_bibite_employee.bbt_emp_id,
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
		Joins("JOIN m_address_temp on m_bibite_employee_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_bibite_employee on m_bibite_employee_temp.bbt_emp_id=m_bibite_employee.bbt_emp_id").
		Joins("JOIN m_address on m_bibite_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("bbt_emp_temp_id = ?", pid).
		Find(&bibiteemployee).Error
	if err != nil {
		return &BibiteEmployeeTempData{}, err
	}
	return &bibiteemployee, nil
}

func (p *BibiteEmployeeTemp) FindBibiteEmployeeTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*BibiteEmployeeTempData, error) {
	var err error
	bibiteemployee := BibiteEmployeeTempData{}
	err = db.Debug().Model(&BibiteEmployeeTempData{}).Limit(100).
		Select(`m_bibite_employee_temp.bbt_emp_temp_id,
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
				m_bibite_employee_temp.bbt_emp_temp_status,
				m_bibite_employee_temp.bbt_emp_temp_action_before,
				m_bibite_employee_temp.bbt_emp_temp_created_by,
				m_bibite_employee_temp.bbt_emp_temp_created_at at time zone 'Asia/Jakarta' as address_deleted_at,
				m_bibite_employee.bbt_emp_id,
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
		Joins("JOIN m_address_temp on m_bibite_employee_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_bibite_employee on m_bibite_employee_temp.bbt_emp_id=m_bibite_employee.bbt_emp_id").
		Joins("JOIN m_address on m_bibite_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("bbt_emp_temp_id = ? AND bbt_emp_temp_status = ?", pid, status).
		Find(&bibiteemployee).Error
	if err != nil {
		return &BibiteEmployeeTempData{}, err
	}
	return &bibiteemployee, nil
}

func (p *BibiteEmployeeTemp) UpdateBibiteEmployeeTemp(db *gorm.DB, pid uint64) (*BibiteEmployeeTemp, error) {
	var err error
	db = db.Debug().Model(&BibiteEmployeeTemp{}).Where("bbt_emp_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_emp_temp_code":        p.BibiteEmployeeTempCode,
			"bbt_emp_temp_name":        p.BibiteEmployeeTempName,
			"bbt_emp_temp_gender":      p.BibiteEmployeeTempGender,
			"bbt_emp_temp_birth_place": p.BibiteEmployeeTempBirthPlace,
			"bbt_emp_temp_birth_date":  p.BibiteEmployeeTempBirthDate,
			"bbt_emp_temp_address":     p.BibiteEmployeeTempAddress,
			"address_temp_id":          p.AddressTempID,
			"bbt_emp_temp_photo":       p.BibiteEmployeeTempPhoto,
			"bbt_emp_temp_reason":      p.BibiteEmployeeTempReason,
		},
	)
	err = db.Debug().Model(&BibiteEmployeeTemp{}).Where("bbt_emp_temp_id = ?", pid).Error
	if err != nil {
		return &BibiteEmployeeTemp{}, err
	}
	return p, nil
}

func (p *BibiteEmployeeTemp) UpdateBibiteEmployeeTempStatus(db *gorm.DB, pid uint64) (*BibiteEmployeeTemp, error) {
	var err error
	db = db.Debug().Model(&BibiteEmployeeTemp{}).Where("bbt_emp_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_emp_temp_reason": p.BibiteEmployeeTempReason,
			"bbt_emp_temp_status": p.BibiteEmployeeTempStatus,
		},
	)
	err = db.Debug().Model(&BibiteEmployeeTemp{}).Where("bbt_emp_temp_id = ?", pid).Error
	if err != nil {
		return &BibiteEmployeeTemp{}, err
	}
	return p, nil
}

func (p *BibiteEmployeeTemp) DeleteBibiteEmployeeTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&BibiteEmployeeTemp{}).Where("bbt_emp_temp_id = ?", pid).Take(&BibiteEmployeeTemp{}).Delete(&BibiteEmployeeTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("BibiteEmployeeTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================
