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

type MitraEmployeeTemp struct {
	MitraEmployeeTempID           uint64    `gorm:"column:mitra_emp_temp_id;primary_key;" json:"mitra_emp_temp_id"`
	MitraEmployeeID               uint64    `gorm:"column:mitra_emp_id" json:"mitra_emp_id"`
	MitraTempID                   uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	MitraBranchTempID             uint64    `gorm:"column:mitra_branch_temp_id" json:"mitra_branch_temp_id"`
	MitraEmployeeTempCode         string    `gorm:"column:mitra_emp_temp_code" json:"mitra_emp_temp_code"`
	MitraEmployeeTempName         string    `gorm:"column:mitra_emp_temp_name" json:"mitra_emp_temp_name"`
	MitraEmployeeTempGender       string    `gorm:"column:mitra_emp_temp_gender" json:"mitra_emp_temp_gender"`
	MitraEmployeeTempBirthPlace   string    `gorm:"column:mitra_emp_temp_birth_place" json:"mitra_emp_temp_birth_place"`
	MitraEmployeeTempBirthDate    string    `gorm:"column:mitra_emp_temp_birth_date" json:"mitra_emp_temp_birth_date"`
	MitraEmployeeTempAddress      string    `gorm:"column:mitra_emp_temp_address" json:"mitra_emp_temp_address"`
	AddressTempID                 uint64    `gorm:"column:address_temp_id" json:"address_temp_id"`
	MitraEmployeeTempPhoto        string    `gorm:"column:mitra_emp_temp_photo" json:"mitra_emp_temp_photo"`
	MitraEmployeeTempReason       string    `gorm:"column:mitra_emp_temp_reason" json:"mitra_emp_temp_reason"`
	MitraEmployeeTempStatus       uint64    `gorm:"column:mitra_emp_temp_status;size:2" json:"mitra_emp_temp_status"`
	MitraEmployeeTempActionBefore uint64    `gorm:"column:mitra_emp_temp_action_before;size:2" json:"mitra_emp_temp_action_before"`
	MitraEmployeeTempCreatedBy    string    `gorm:"column:mitra_emp_temp_created_by;size:125" json:"mitra_emp_temp_created_by"`
	MitraEmployeeTempCreatedAt    time.Time `gorm:"column:mitra_emp_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_emp_temp_created_at"`
}

type MitraEmployeeTempPend struct {
	MitraEmployeeTempID           *uint64   `gorm:"column:mitra_emp_temp_id;primary_key;" json:"mitra_emp_temp_id"`
	MitraEmployeeID               *uint64   `gorm:"column:mitra_emp_id" json:"mitra_emp_id"`
	MitraTempID                   uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	MitraBranchTempID             uint64    `gorm:"column:mitra_branch_temp_id" json:"mitra_branch_temp_id"`
	MitraEmployeeTempCode         string    `gorm:"column:mitra_emp_temp_code" json:"mitra_emp_temp_code"`
	MitraEmployeeTempName         string    `gorm:"column:mitra_emp_temp_name" json:"mitra_emp_temp_name"`
	MitraEmployeeTempGender       string    `gorm:"column:mitra_emp_temp_gender" json:"mitra_emp_temp_gender"`
	MitraEmployeeTempBirthPlace   string    `gorm:"column:mitra_emp_temp_birth_place" json:"mitra_emp_temp_birth_place"`
	MitraEmployeeTempBirthDate    string    `gorm:"column:mitra_emp_temp_birth_date" json:"mitra_emp_temp_birth_date"`
	MitraEmployeeTempAddress      string    `gorm:"column:mitra_emp_temp_address" json:"mitra_emp_temp_address"`
	AddressTempID                 uint64    `gorm:"column:address_temp_id" json:"address_temp_id"`
	MitraEmployeeTempPhoto        string    `gorm:"column:mitra_emp_temp_photo" json:"mitra_emp_temp_photo"`
	MitraEmployeeTempReason       string    `gorm:"column:mitra_emp_temp_reason" json:"mitra_emp_temp_reason"`
	MitraEmployeeTempStatus       uint64    `gorm:"column:mitra_emp_temp_status;size:2" json:"mitra_emp_temp_status"`
	MitraEmployeeTempActionBefore uint64    `gorm:"column:mitra_emp_temp_action_before;size:2" json:"mitra_emp_temp_action_before"`
	MitraEmployeeTempCreatedBy    string    `gorm:"column:mitra_emp_temp_created_by;size:125" json:"mitra_emp_temp_created_by"`
	MitraEmployeeTempCreatedAt    time.Time `gorm:"column:mitra_emp_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_emp_temp_created_at"`
}

type MitraEmployeeTempData struct {
	MitraEmployeeTempID           uint64    `gorm:"column:mitra_emp_temp_id" json:"mitra_emp_temp_id"`
	MitraTempID                   uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	MitraTempName                 string    `gorm:"column:mitra_temp_name" json:"mitra_temp_name"`
	MitraBranchTempID             uint64    `gorm:"column:mitra_branch_temp_id" json:"mitra_branch_temp_id"`
	MitraBranchTempName           string    `gorm:"column:mitra_branch_temp_name" json:"mitra_branch_temp_name"`
	MitraEmployeeTempCode         string    `gorm:"column:mitra_emp_temp_code" json:"mitra_emp_temp_code"`
	MitraEmployeeTempName         string    `gorm:"column:mitra_emp_temp_name" json:"mitra_emp_temp_name"`
	MitraEmployeeTempGender       string    `gorm:"column:mitra_emp_temp_gender" json:"mitra_emp_temp_gender"`
	MitraEmployeeTempBirthPlace   string    `gorm:"column:mitra_emp_temp_birth_place" json:"mitra_emp_temp_birth_place"`
	MitraEmployeeTempBirthDate    string    `gorm:"column:mitra_emp_temp_birth_date" json:"mitra_emp_temp_birth_date"`
	MitraEmployeeTempAddress      string    `gorm:"column:mitra_emp_temp_address" json:"mitra_emp_temp_address"`
	AddressTempID                 uint64    `gorm:"column:address_temp_id" json:"address_temp_id"`
	CountryTempID                 uint64    `gorm:"column:country_temp_id" json:"country_temp_id"`
	CountryTempName               string    `gorm:"column:country_temp_name" json:"country_temp_name"`
	ProvinceTempID                uint64    `gorm:"column:province_temp_id" json:"province_temp_id"`
	ProvinceTempName              string    `gorm:"column:province_temp_name" json:"province_temp_name"`
	RegencyTempID                 uint64    `gorm:"column:regency_temp_id" json:"regency_temp_id"`
	RegencyTempName               string    `gorm:"column:regency_temp_name" json:"regency_temp_name"`
	DistrictTempID                uint64    `gorm:"column:district_temp_id" json:"district_temp_id"`
	DistrictTempName              string    `gorm:"column:district_temp_name" json:"district_temp_name"`
	VillageTempID                 uint64    `gorm:"column:village_temp_id" json:"village_temp_id"`
	VillageTempName               string    `gorm:"column:village_temp_name" json:"village_temp_name"`
	MitraEmployeeTempPhoto        string    `gorm:"column:mitra_emp_temp_photo" json:"mitra_emp_temp_photo"`
	MitraEmployeeTempReason       string    `gorm:"column:mitra_emp_temp_reason" json:"mitra_emp_temp_reason"`
	MitraEmployeeTempStatus       uint64    `gorm:"column:mitra_emp_temp_status;size:2" json:"mitra_emp_temp_status"`
	MitraEmployeeTempActionBefore uint64    `gorm:"column:mitra_emp_temp_action_before;size:2" json:"mitra_emp_temp_action_before"`
	MitraEmployeeTempCreatedBy    string    `gorm:"column:mitra_emp_temp_created_by;size:125" json:"mitra_emp_temp_created_by"`
	MitraEmployeeTempCreatedAt    time.Time `gorm:"column:mitra_emp_temp_created_at" json:"mitra_emp_temp_created_at"`
	MitraEmployeeID               uint64    `gorm:"column:mitra_emp_id" json:"mitra_emp_id"`
	MitraID                       uint64    `gorm:"column:mitra_id" json:"mitra_id"`
	MitraName                     string    `gorm:"column:mitra_name" json:"mitra_name"`
	MitraBranchID                 uint64    `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	MitraBranchName               string    `gorm:"column:mitra_branch_name" json:"mitra_branch_name"`
	MitraEmployeeCode             string    `gorm:"column:mitra_emp_code" json:"mitra_emp_code"`
	MitraEmployeeName             string    `gorm:"column:mitra_emp_name" json:"mitra_emp_name"`
	MitraEmployeeGender           string    `gorm:"column:mitra_emp_gender" json:"mitra_emp_gender"`
	MitraEmployeeBirthPlace       string    `gorm:"column:mitra_emp_birth_place" json:"mitra_emp_birth_place"`
	MitraEmployeeBirthDate        string    `gorm:"column:mitra_emp_birth_date" json:"mitra_emp_birth_date"`
	MitraEmployeeAddress          string    `gorm:"column:mitra_emp_address" json:"mitra_emp_address"`
	AddressID                     uint64    `gorm:"column:address_id" json:"address_id"`
	CountryID                     uint64    `gorm:"column:country_id" json:"country_id"`
	CountryName                   string    `gorm:"column:country_name" json:"country_name"`
	ProvinceID                    uint64    `gorm:"column:province_id" json:"province_id"`
	ProvinceName                  string    `gorm:"column:province_name" json:"province_name"`
	RegencyID                     uint64    `gorm:"column:regency_id" json:"regency_id"`
	RegencyName                   string    `gorm:"column:regency_name" json:"regency_name"`
	DistrictID                    uint64    `gorm:"column:district_id" json:"district_id"`
	DistrictName                  string    `gorm:"column:district_name" json:"district_name"`
	VillageID                     uint64    `gorm:"column:village_id" json:"village_id"`
	VillageName                   string    `gorm:"column:village_name" json:"village_name"`
	MitraEmployeePhoto            string    `gorm:"column:mitra_emp_photo" json:"mitra_emp_photo"`
	MitraEmployeeStatus           uint64    `gorm:"column:mitra_emp_status" json:"mitra_emp_status"`
	MitraEmployeeCreatedBy        string    `gorm:"column:mitra_emp_created_by" json:"mitra_emp_created_by"`
	MitraEmployeeCreatedAt        time.Time `gorm:"column:mitra_emp_created_at" json:"mitra_emp_created_at"`
	MitraEmployeeUpdatedBy        string    `gorm:"column:mitra_emp_updated_by" json:"mitra_emp_updated_by"`
	MitraEmployeeUpdatedAt        time.Time `gorm:"column:mitra_emp_updated_at" json:"mitra_emp_updated_at"`
	MitraEmployeeDeletedBy        string    `gorm:"column:mitra_emp_deleted_by" json:"mitra_emp_deleted_by"`
	MitraEmployeeDeletedAt        time.Time `gorm:"column:mitra_emp_deleted_at" json:"mitra_emp_deleted_at"`
}

type ResponseMitraEmployeeTemps struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    *[]MitraEmployeeTempData `json:"data"`
}

type ResponseMitraEmployeeTemp struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    *MitraEmployeeTempData `json:"data"`
}

type ResponseMitraEmployeeTempsPend struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    *[]MitraEmployeeTempPend `json:"data"`
}

type ResponseMitraEmployeeTempPend struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    *MitraEmployeeTempPend `json:"data"`
}

type ResponseMitraEmployeeTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraEmployeeTemp) TableName() string {
	return "m_mitra_employee_temp"
}

func (MitraEmployeeTempPend) TableName() string {
	return "m_mitra_employee_temp"
}

func (MitraEmployeeTempData) TableName() string {
	return "m_mitra_employee_temp"
}

func (p *MitraEmployeeTemp) Prepare() {
	p.MitraEmployeeID = p.MitraEmployeeID
	p.MitraTempID = p.MitraTempID
	p.MitraBranchTempID = p.MitraBranchTempID
	p.MitraEmployeeTempCode = p.MitraEmployeeTempCode
	p.MitraEmployeeTempName = p.MitraEmployeeTempName
	p.MitraEmployeeTempGender = p.MitraEmployeeTempGender
	p.MitraEmployeeTempBirthPlace = p.MitraEmployeeTempBirthPlace
	p.MitraEmployeeTempBirthDate = p.MitraEmployeeTempBirthDate
	p.MitraEmployeeTempAddress = p.MitraEmployeeTempAddress
	p.AddressTempID = p.AddressTempID
	p.MitraEmployeeTempPhoto = p.MitraEmployeeTempPhoto
	p.MitraEmployeeTempReason = p.MitraEmployeeTempReason
	p.MitraEmployeeTempStatus = p.MitraEmployeeTempStatus
	p.MitraEmployeeTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraEmployeeTempCreatedBy))
	p.MitraEmployeeTempCreatedAt = time.Now()
}

func (p *MitraEmployeeTempPend) Prepare() {
	p.MitraEmployeeTempID = nil
	p.MitraEmployeeID = nil
	p.MitraTempID = p.MitraTempID
	p.MitraBranchTempID = p.MitraBranchTempID
	p.MitraEmployeeTempCode = p.MitraEmployeeTempCode
	p.MitraEmployeeTempName = p.MitraEmployeeTempName
	p.MitraEmployeeTempGender = p.MitraEmployeeTempGender
	p.MitraEmployeeTempBirthPlace = p.MitraEmployeeTempBirthPlace
	p.MitraEmployeeTempBirthDate = p.MitraEmployeeTempBirthDate
	p.MitraEmployeeTempAddress = p.MitraEmployeeTempAddress
	p.AddressTempID = p.AddressTempID
	p.MitraEmployeeTempPhoto = p.MitraEmployeeTempPhoto
	p.MitraEmployeeTempReason = p.MitraEmployeeTempReason
	p.MitraEmployeeTempStatus = p.MitraEmployeeTempStatus
	p.MitraEmployeeTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraEmployeeTempCreatedBy))
	p.MitraEmployeeTempCreatedAt = time.Now()
}

func (p *MitraEmployeeTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	case "insertupdate":
		if p.MitraEmployeeTempCode == "" {
			return errors.New("Negara Wajib diisi")
		}
		if p.MitraEmployeeTempName == "" {
			return errors.New("Provinsi Wajib diisi")
		}
		if p.MitraEmployeeTempGender == "" {
			return errors.New("Kabupaten / Kota Wajib diisi")
		}
		if p.MitraEmployeeTempBirthPlace == "" {
			return errors.New("Kecamatan Wajib diisi")
		}
		if p.MitraEmployeeTempBirthDate == "" {
			return errors.New("Kelurahan Wajib diisi")
		}
		if p.MitraEmployeeTempAddress == "" {
			return errors.New("Kelurahan Wajib diisi")
		}
		if p.AddressTempID == 0 {
			return errors.New("Kelurahan Wajib diisi")
		}
		if p.MitraEmployeeTempPhoto == "" {
			return errors.New("Kelurahan Wajib diisi")
		}
		if p.MitraEmployeeTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.MitraEmployeeTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	case "reason":
		if p.MitraEmployeeTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.MitraEmployeeTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	default:
		if p.MitraEmployeeTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil
	}
}

func (p *MitraEmployeeTemp) SaveMitraEmployeeTemp(db *gorm.DB) (*MitraEmployeeTemp, error) {
	var err error
	err = db.Debug().Model(&MitraEmployeeTemp{}).Create(&p).Error
	if err != nil {
		return &MitraEmployeeTemp{}, err
	}
	return p, nil
}

func (p *MitraEmployeeTempPend) SaveMitraEmployeeTempPend(db *gorm.DB) (*MitraEmployeeTempPend, error) {
	var err error
	err = db.Debug().Model(&MitraEmployeeTempPend{}).Create(&p).Error
	if err != nil {
		return &MitraEmployeeTempPend{}, err
	}
	return p, nil
}

func (p *MitraEmployeeTemp) FindAllMitraEmployeeTemp(db *gorm.DB) (*[]MitraEmployeeTempPend, error) {
	var err error
	mitraemployee := []MitraEmployeeTempPend{}
	err = db.Debug().Model(&MitraEmployeeTempPend{}).Limit(100).
		Select(`m_mitra_employee_temp.mitra_emp_temp_id,
				m_mitra_employee_temp.mitra_temp_id,
				m_mitra_employee_temp.mitra_branch_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_code,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_gender,
				m_mitra_employee_temp.mitra_emp_temp_birth_place,
				m_mitra_employee_temp.mitra_emp_temp_birth_date,
				m_mitra_employee_temp.mitra_emp_temp_address,
				m_mitra_employee_temp.address_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_photo,
				m_mitra_employee_temp.mitra_emp_temp_reason,
				m_mitra_employee_temp.mitra_emp_temp_status,
				m_mitra_employee_temp.mitra_emp_temp_action_before,
				m_mitra_employee_temp.mitra_emp_temp_created_by,
				m_mitra_employee_temp.mitra_emp_temp_created_at`).
		Find(&mitraemployee).Error
	if err != nil {
		return &[]MitraEmployeeTempPend{}, err
	}
	return &mitraemployee, nil
}

func (p *MitraEmployeeTemp) FindAllMitraEmployeeTempStatusPendingActive(db *gorm.DB) (*[]MitraEmployeeTempPend, error) {
	var err error
	mitraemployee := []MitraEmployeeTempPend{}
	err = db.Debug().Model(&MitraEmployeeTempPend{}).Limit(100).
		Select(`m_mitra_employee_temp.mitra_emp_temp_id,
				m_mitra_employee_temp.mitra_temp_id,
				m_mitra_employee_temp.mitra_branch_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_code,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_gender,
				m_mitra_employee_temp.mitra_emp_temp_birth_place,
				m_mitra_employee_temp.mitra_emp_temp_birth_date,
				m_mitra_employee_temp.mitra_emp_temp_address,
				m_mitra_employee_temp.address_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_photo,
				m_mitra_employee_temp.mitra_emp_temp_reason,
				m_mitra_employee_temp.mitra_emp_temp_status,
				m_mitra_employee_temp.mitra_emp_temp_action_before,
				m_mitra_employee_temp.mitra_emp_temp_created_by,
				m_mitra_employee_temp.mitra_emp_temp_created_at`).
		Where("mitra_emp_temp_status = ?", 11).
		Find(&mitraemployee).Error
	if err != nil {
		return &[]MitraEmployeeTempPend{}, err
	}
	return &mitraemployee, nil
}

func (p *MitraEmployeeTemp) FindAllMitraEmployeeTempStatus(db *gorm.DB, status uint64) (*[]MitraEmployeeTempData, error) {
	var err error
	mitraemployee := []MitraEmployeeTempData{}
	err = db.Debug().Model(&MitraEmployeeTempData{}).Limit(100).
		Select(`m_mitra_employee_temp.mitra_emp_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_code,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_gender,
				m_mitra_employee_temp.mitra_emp_temp_birth_place,
				m_mitra_employee_temp.mitra_emp_temp_birth_date,
				m_mitra_employee_temp.mitra_emp_temp_address,
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
				m_mitra_employee_temp.mitra_emp_temp_photo,
				m_mitra_employee_temp.mitra_emp_temp_status,
				m_mitra_employee_temp.mitra_emp_temp_action_before,
				m_mitra_employee_temp.mitra_emp_temp_created_by,
				m_mitra_employee.mitra_emp_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_employee.mitra_emp_code,
				m_mitra_employee.mitra_emp_name,
				m_mitra_employee.mitra_emp_gender,
				m_mitra_employee.mitra_emp_birth_place,
				m_mitra_employee.mitra_emp_birth_date,
				m_mitra_employee.mitra_emp_address,
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
				m_mitra_employee.mitra_emp_photo,
				m_mitra_employee.mitra_emp_status,
				m_mitra_employee.mitra_emp_created_by,
				m_mitra_employee.mitra_emp_created_at,
				m_mitra_employee.mitra_emp_updated_by,
				m_mitra_employee.mitra_emp_updated_at,
				m_mitra_employee.mitra_emp_deleted_by,
				m_mitra_employee.mitra_emp_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_employee_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_employee_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_address_temp on m_mitra_employee_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_mitra_employee on m_mitra_employee_temp.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_address on m_mitra_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("mitra_emp_temp_status = ?", status).
		Find(&mitraemployee).Error
	if err != nil {
		return &[]MitraEmployeeTempData{}, err
	}
	return &mitraemployee, nil
}

func (p *MitraEmployeeTemp) FindMitraEmployeeTempDataByID(db *gorm.DB, pid uint64) (*MitraEmployeeTemp, error) {
	var err error
	err = db.Debug().Model(&MitraEmployeeTemp{}).Where("mitra_emp_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraEmployeeTemp{}, err
	}
	return p, nil
}

func (p *MitraEmployeeTemp) FindMitraEmployeeTempByIDPendingActive(db *gorm.DB, pid uint64) (*MitraEmployeeTempPend, error) {
	var err error
	mitraemployee := MitraEmployeeTempPend{}
	err = db.Debug().Model(&MitraEmployeeTempPend{}).Limit(100).
		Select(`m_mitra_employee_temp.mitra_emp_temp_id,
				m_mitra_employee_temp.mitra_temp_id,
				m_mitra_employee_temp.mitra_branch_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_code,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_gender,
				m_mitra_employee_temp.mitra_emp_temp_birth_place,
				m_mitra_employee_temp.mitra_emp_temp_birth_date,
				m_mitra_employee_temp.mitra_emp_temp_address,
				m_mitra_employee_temp.address_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_photo,
				m_mitra_employee_temp.mitra_emp_temp_reason,
				m_mitra_employee_temp.mitra_emp_temp_status,
				m_mitra_employee_temp.mitra_emp_temp_action_before,
				m_mitra_employee_temp.mitra_emp_temp_created_by,
				m_mitra_employee_temp.mitra_emp_temp_created_at`).
		Where("mitra_emp_temp_id = ?", pid).
		Find(&mitraemployee).Error
	if err != nil {
		return &MitraEmployeeTempPend{}, err
	}
	return &mitraemployee, nil
}

func (p *MitraEmployeeTemp) FindMitraEmployeeTempStatusByIDPendingActive(db *gorm.DB, pid uint64) (*MitraEmployeeTempPend, error) {
	var err error
	mitraemployee := MitraEmployeeTempPend{}
	err = db.Debug().Model(&MitraEmployeeTempPend{}).Limit(100).
		Select(`m_mitra_employee_temp.mitra_emp_temp_id,
				m_mitra_employee_temp.mitra_temp_id,
				m_mitra_employee_temp.mitra_branch_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_code,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_gender,
				m_mitra_employee_temp.mitra_emp_temp_birth_place,
				m_mitra_employee_temp.mitra_emp_temp_birth_date,
				m_mitra_employee_temp.mitra_emp_temp_address,
				m_mitra_employee_temp.address_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_photo,
				m_mitra_employee_temp.mitra_emp_temp_reason,
				m_mitra_employee_temp.mitra_emp_temp_status,
				m_mitra_employee_temp.mitra_emp_temp_action_before,
				m_mitra_employee_temp.mitra_emp_temp_created_by,
				m_mitra_employee_temp.mitra_emp_temp_created_at`).
		Where("mitra_emp_temp_id = ? AND mitra_emp_temp_status = ?", pid, 11).
		Find(&mitraemployee).Error
	if err != nil {
		return &MitraEmployeeTempPend{}, err
	}
	return &mitraemployee, nil
}

func (p *MitraEmployeeTemp) FindMitraEmployeeTempByID(db *gorm.DB, pid uint64) (*MitraEmployeeTempData, error) {
	var err error
	mitraemployee := MitraEmployeeTempData{}
	err = db.Debug().Model(&MitraEmployeeTempData{}).Limit(100).
		Select(`m_mitra_employee_temp.mitra_emp_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_code,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_gender,
				m_mitra_employee_temp.mitra_emp_temp_birth_place,
				m_mitra_employee_temp.mitra_emp_temp_birth_date,
				m_mitra_employee_temp.mitra_emp_temp_address,
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
				m_mitra_employee_temp.mitra_emp_temp_photo,
				m_mitra_employee_temp.mitra_emp_temp_status,
				m_mitra_employee_temp.mitra_emp_temp_action_before,
				m_mitra_employee_temp.mitra_emp_temp_created_by,
				m_mitra_employee.mitra_emp_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_employee.mitra_emp_code,
				m_mitra_employee.mitra_emp_name,
				m_mitra_employee.mitra_emp_gender,
				m_mitra_employee.mitra_emp_birth_place,
				m_mitra_employee.mitra_emp_birth_date,
				m_mitra_employee.mitra_emp_address,
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
				m_mitra_employee.mitra_emp_photo,
				m_mitra_employee.mitra_emp_status,
				m_mitra_employee.mitra_emp_created_by,
				m_mitra_employee.mitra_emp_created_at,
				m_mitra_employee.mitra_emp_updated_by,
				m_mitra_employee.mitra_emp_updated_at,
				m_mitra_employee.mitra_emp_deleted_by,
				m_mitra_employee.mitra_emp_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_employee_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_employee_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_address_temp on m_mitra_employee_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_mitra_employee on m_mitra_employee_temp.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_address on m_mitra_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("mitra_emp_temp_id = ?", pid).
		Find(&mitraemployee).Error
	if err != nil {
		return &MitraEmployeeTempData{}, err
	}
	return &mitraemployee, nil
}

func (p *MitraEmployeeTemp) FindMitraEmployeeTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraEmployeeTempData, error) {
	var err error
	mitraemployee := MitraEmployeeTempData{}
	err = db.Debug().Model(&MitraEmployeeTempData{}).Limit(100).
		Select(`m_mitra_employee_temp.mitra_emp_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_code,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_gender,
				m_mitra_employee_temp.mitra_emp_temp_birth_place,
				m_mitra_employee_temp.mitra_emp_temp_birth_date,
				m_mitra_employee_temp.mitra_emp_temp_address,
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
				m_mitra_employee_temp.mitra_emp_temp_photo,
				m_mitra_employee_temp.mitra_emp_temp_status,
				m_mitra_employee_temp.mitra_emp_temp_action_before,
				m_mitra_employee_temp.mitra_emp_temp_created_by,
				m_mitra_employee.mitra_emp_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_employee.mitra_emp_code,
				m_mitra_employee.mitra_emp_name,
				m_mitra_employee.mitra_emp_gender,
				m_mitra_employee.mitra_emp_birth_place,
				m_mitra_employee.mitra_emp_birth_date,
				m_mitra_employee.mitra_emp_address,
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
				m_mitra_employee.mitra_emp_photo,
				m_mitra_employee.mitra_emp_status,
				m_mitra_employee.mitra_emp_created_by,
				m_mitra_employee.mitra_emp_created_at,
				m_mitra_employee.mitra_emp_updated_by,
				m_mitra_employee.mitra_emp_updated_at,
				m_mitra_employee.mitra_emp_deleted_by,
				m_mitra_employee.mitra_emp_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_employee_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_employee_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_address_temp on m_mitra_employee_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_mitra_employee on m_mitra_employee_temp.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_address on m_mitra_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("mitra_emp_temp_id = ? AND mitra_emp_temp_status = ?", pid, status).
		Find(&mitraemployee).Error
	if err != nil {
		return &MitraEmployeeTempData{}, err
	}
	return &mitraemployee, nil
}

func (p *MitraEmployeeTemp) UpdateMitraEmployeeTemp(db *gorm.DB, pid uint64) (*MitraEmployeeTemp, error) {
	var err error
	db = db.Debug().Model(&MitraEmployeeTemp{}).Where("mitra_emp_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_temp_id":              p.MitraTempID,
			"mitra_branch_temp_id":       p.MitraBranchTempID,
			"mitra_emp_temp_code":        p.MitraEmployeeTempCode,
			"mitra_emp_temp_name":        p.MitraEmployeeTempName,
			"mitra_emp_temp_gender":      p.MitraEmployeeTempGender,
			"mitra_emp_temp_birth_place": p.MitraEmployeeTempBirthPlace,
			"mitra_emp_temp_birth_date":  p.MitraEmployeeTempBirthDate,
			"mitra_emp_temp_address":     p.MitraEmployeeTempAddress,
			"address_temp_id":            p.AddressTempID,
			"mitra_emp_temp_photo":       p.MitraEmployeeTempPhoto,
			"mitra_emp_temp_reason":      p.MitraEmployeeTempReason,
		},
	)
	err = db.Debug().Model(&MitraEmployeeTemp{}).Where("mitra_emp_temp_id = ?", pid).Error
	if err != nil {
		return &MitraEmployeeTemp{}, err
	}
	return p, nil
}

func (p *MitraEmployeeTemp) UpdateMitraEmployeeTempStatus(db *gorm.DB, pid uint64) (*MitraEmployeeTemp, error) {
	var err error
	db = db.Debug().Model(&MitraEmployeeTemp{}).Where("mitra_emp_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_emp_temp_reason": p.MitraEmployeeTempReason,
			"mitra_emp_temp_status": p.MitraEmployeeTempStatus,
		},
	)
	err = db.Debug().Model(&MitraEmployeeTemp{}).Where("mitra_emp_temp_id = ?", pid).Error
	if err != nil {
		return &MitraEmployeeTemp{}, err
	}
	return p, nil
}

func (p *MitraEmployeeTemp) DeleteMitraEmployeeTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraEmployeeTemp{}).Where("mitra_emp_temp_id = ?", pid).Take(&MitraEmployeeTemp{}).Delete(&MitraEmployeeTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraEmployeeTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//GET BY MITRA ID
//====================================================================================================================================

func (p *MitraEmployeeTemp) FindAllMitraEmployeeTempMitra(db *gorm.DB, mitra uint64) (*[]MitraEmployeeTempPend, error) {
	var err error
	mitraemployee := []MitraEmployeeTempPend{}
	err = db.Debug().Model(&MitraEmployeeTempPend{}).Limit(100).
		Select(`m_mitra_employee_temp.mitra_emp_temp_id,
				m_mitra_employee_temp.mitra_temp_id,
				m_mitra_employee_temp.mitra_branch_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_code,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_gender,
				m_mitra_employee_temp.mitra_emp_temp_birth_place,
				m_mitra_employee_temp.mitra_emp_temp_birth_date,
				m_mitra_employee_temp.mitra_emp_temp_address,
				m_mitra_employee_temp.address_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_photo,
				m_mitra_employee_temp.mitra_emp_temp_reason,
				m_mitra_employee_temp.mitra_emp_temp_status,
				m_mitra_employee_temp.mitra_emp_temp_action_before,
				m_mitra_employee_temp.mitra_emp_temp_created_by,
				m_mitra_employee_temp.mitra_emp_temp_created_at`).
		Where("m_mitra_employee_temp.mitra_temp_id = ?", mitra).
		Find(&mitraemployee).Error
	if err != nil {
		return &[]MitraEmployeeTempPend{}, err
	}
	return &mitraemployee, nil
}

func (p *MitraEmployeeTemp) FindAllMitraEmployeeTempStatusPendingActiveMitra(db *gorm.DB, mitra uint64) (*[]MitraEmployeeTempPend, error) {
	var err error
	mitraemployee := []MitraEmployeeTempPend{}
	err = db.Debug().Model(&MitraEmployeeTempPend{}).Limit(100).
		Select(`m_mitra_employee_temp.mitra_emp_temp_id,
				m_mitra_employee_temp.mitra_temp_id,
				m_mitra_employee_temp.mitra_branch_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_code,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_gender,
				m_mitra_employee_temp.mitra_emp_temp_birth_place,
				m_mitra_employee_temp.mitra_emp_temp_birth_date,
				m_mitra_employee_temp.mitra_emp_temp_address,
				m_mitra_employee_temp.address_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_photo,
				m_mitra_employee_temp.mitra_emp_temp_reason,
				m_mitra_employee_temp.mitra_emp_temp_status,
				m_mitra_employee_temp.mitra_emp_temp_action_before,
				m_mitra_employee_temp.mitra_emp_temp_created_by,
				m_mitra_employee_temp.mitra_emp_temp_created_at`).
		Where("mitra_emp_temp_status = ? AND m_mitra_employee_temp.mitra_temp_id = ?", 11, mitra).
		Find(&mitraemployee).Error
	if err != nil {
		return &[]MitraEmployeeTempPend{}, err
	}
	return &mitraemployee, nil
}

func (p *MitraEmployeeTemp) FindAllMitraEmployeeTempStatusMitra(db *gorm.DB, status uint64, mitra uint64) (*[]MitraEmployeeTempData, error) {
	var err error
	mitraemployee := []MitraEmployeeTempData{}
	err = db.Debug().Model(&MitraEmployeeTempData{}).Limit(100).
		Select(`m_mitra_employee_temp.mitra_emp_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_code,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_gender,
				m_mitra_employee_temp.mitra_emp_temp_birth_place,
				m_mitra_employee_temp.mitra_emp_temp_birth_date,
				m_mitra_employee_temp.mitra_emp_temp_address,
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
				m_mitra_employee_temp.mitra_emp_temp_photo,
				m_mitra_employee_temp.mitra_emp_temp_status,
				m_mitra_employee_temp.mitra_emp_temp_action_before,
				m_mitra_employee_temp.mitra_emp_temp_created_by,
				m_mitra_employee.mitra_emp_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_employee.mitra_emp_code,
				m_mitra_employee.mitra_emp_name,
				m_mitra_employee.mitra_emp_gender,
				m_mitra_employee.mitra_emp_birth_place,
				m_mitra_employee.mitra_emp_birth_date,
				m_mitra_employee.mitra_emp_address,
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
				m_mitra_employee.mitra_emp_photo,
				m_mitra_employee.mitra_emp_status,
				m_mitra_employee.mitra_emp_created_by,
				m_mitra_employee.mitra_emp_created_at,
				m_mitra_employee.mitra_emp_updated_by,
				m_mitra_employee.mitra_emp_updated_at,
				m_mitra_employee.mitra_emp_deleted_by,
				m_mitra_employee.mitra_emp_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_employee_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_employee_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_address_temp on m_mitra_employee_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_mitra_employee on m_mitra_employee_temp.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_address on m_mitra_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("mitra_emp_temp_status = ? AND m_mitra_employee_temp.mitra_temp_id = ?", status, mitra).
		Find(&mitraemployee).Error
	if err != nil {
		return &[]MitraEmployeeTempData{}, err
	}
	return &mitraemployee, nil
}

func (p *MitraEmployeeTemp) FindMitraEmployeeTempDataByIDMitra(db *gorm.DB, pid uint64, mitra uint64) (*MitraEmployeeTemp, error) {
	var err error
	err = db.Debug().Model(&MitraEmployeeTemp{}).Where("mitra_emp_temp_id = ? AND m_mitra_employee_temp.mitra_temp_id = ?", pid, mitra).Take(&p).Error
	if err != nil {
		return &MitraEmployeeTemp{}, err
	}
	return p, nil
}

func (p *MitraEmployeeTemp) FindMitraEmployeeTempByIDPendingActiveMitra(db *gorm.DB, pid uint64, mitra uint64) (*MitraEmployeeTempPend, error) {
	var err error
	mitraemployee := MitraEmployeeTempPend{}
	err = db.Debug().Model(&MitraEmployeeTempPend{}).Limit(100).
		Select(`m_mitra_employee_temp.mitra_emp_temp_id,
				m_mitra_employee_temp.mitra_temp_id,
				m_mitra_employee_temp.mitra_branch_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_code,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_gender,
				m_mitra_employee_temp.mitra_emp_temp_birth_place,
				m_mitra_employee_temp.mitra_emp_temp_birth_date,
				m_mitra_employee_temp.mitra_emp_temp_address,
				m_mitra_employee_temp.address_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_photo,
				m_mitra_employee_temp.mitra_emp_temp_reason,
				m_mitra_employee_temp.mitra_emp_temp_status,
				m_mitra_employee_temp.mitra_emp_temp_action_before,
				m_mitra_employee_temp.mitra_emp_temp_created_by,
				m_mitra_employee_temp.mitra_emp_temp_created_at`).
		Where("mitra_emp_temp_id = ? AND m_mitra_employee_temp.mitra_temp_id = ?", pid, mitra).
		Find(&mitraemployee).Error
	if err != nil {
		return &MitraEmployeeTempPend{}, err
	}
	return &mitraemployee, nil
}

func (p *MitraEmployeeTemp) FindMitraEmployeeTempStatusByIDPendingActiveMitra(db *gorm.DB, pid uint64, mitra uint64) (*MitraEmployeeTempPend, error) {
	var err error
	mitraemployee := MitraEmployeeTempPend{}
	err = db.Debug().Model(&MitraEmployeeTempPend{}).Limit(100).
		Select(`m_mitra_employee_temp.mitra_emp_temp_id,
				m_mitra_employee_temp.mitra_temp_id,
				m_mitra_employee_temp.mitra_branch_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_code,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_gender,
				m_mitra_employee_temp.mitra_emp_temp_birth_place,
				m_mitra_employee_temp.mitra_emp_temp_birth_date,
				m_mitra_employee_temp.mitra_emp_temp_address,
				m_mitra_employee_temp.address_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_photo,
				m_mitra_employee_temp.mitra_emp_temp_reason,
				m_mitra_employee_temp.mitra_emp_temp_status,
				m_mitra_employee_temp.mitra_emp_temp_action_before,
				m_mitra_employee_temp.mitra_emp_temp_created_by,
				m_mitra_employee_temp.mitra_emp_temp_created_at`).
		Where("mitra_emp_temp_id = ? AND mitra_emp_temp_status = ? AND m_mitra_employee_temp.mitra_temp_id = ?", pid, 11, mitra).
		Find(&mitraemployee).Error
	if err != nil {
		return &MitraEmployeeTempPend{}, err
	}
	return &mitraemployee, nil
}

func (p *MitraEmployeeTemp) FindMitraEmployeeTempByIDMitra(db *gorm.DB, pid uint64, mitra uint64) (*MitraEmployeeTempData, error) {
	var err error
	mitraemployee := MitraEmployeeTempData{}
	err = db.Debug().Model(&MitraEmployeeTempData{}).Limit(100).
		Select(`m_mitra_employee_temp.mitra_emp_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_code,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_gender,
				m_mitra_employee_temp.mitra_emp_temp_birth_place,
				m_mitra_employee_temp.mitra_emp_temp_birth_date,
				m_mitra_employee_temp.mitra_emp_temp_address,
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
				m_mitra_employee_temp.mitra_emp_temp_photo,
				m_mitra_employee_temp.mitra_emp_temp_status,
				m_mitra_employee_temp.mitra_emp_temp_action_before,
				m_mitra_employee_temp.mitra_emp_temp_created_by,
				m_mitra_employee.mitra_emp_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_employee.mitra_emp_code,
				m_mitra_employee.mitra_emp_name,
				m_mitra_employee.mitra_emp_gender,
				m_mitra_employee.mitra_emp_birth_place,
				m_mitra_employee.mitra_emp_birth_date,
				m_mitra_employee.mitra_emp_address,
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
				m_mitra_employee.mitra_emp_photo,
				m_mitra_employee.mitra_emp_status,
				m_mitra_employee.mitra_emp_created_by,
				m_mitra_employee.mitra_emp_created_at,
				m_mitra_employee.mitra_emp_updated_by,
				m_mitra_employee.mitra_emp_updated_at,
				m_mitra_employee.mitra_emp_deleted_by,
				m_mitra_employee.mitra_emp_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_employee_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_employee_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_address_temp on m_mitra_employee_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_mitra_employee on m_mitra_employee_temp.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_address on m_mitra_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("mitra_emp_temp_id = ? AND m_mitra_employee_temp.mitra_temp_id = ?", pid, mitra).
		Find(&mitraemployee).Error
	if err != nil {
		return &MitraEmployeeTempData{}, err
	}
	return &mitraemployee, nil
}

func (p *MitraEmployeeTemp) FindMitraEmployeeTempStatusByIDMitra(db *gorm.DB, pid uint64, status uint64, mitra uint64) (*MitraEmployeeTempData, error) {
	var err error
	mitraemployee := MitraEmployeeTempData{}
	err = db.Debug().Model(&MitraEmployeeTempData{}).Limit(100).
		Select(`m_mitra_employee_temp.mitra_emp_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_code,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_gender,
				m_mitra_employee_temp.mitra_emp_temp_birth_place,
				m_mitra_employee_temp.mitra_emp_temp_birth_date,
				m_mitra_employee_temp.mitra_emp_temp_address,
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
				m_mitra_employee_temp.mitra_emp_temp_photo,
				m_mitra_employee_temp.mitra_emp_temp_status,
				m_mitra_employee_temp.mitra_emp_temp_action_before,
				m_mitra_employee_temp.mitra_emp_temp_created_by,
				m_mitra_employee.mitra_emp_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_employee.mitra_emp_code,
				m_mitra_employee.mitra_emp_name,
				m_mitra_employee.mitra_emp_gender,
				m_mitra_employee.mitra_emp_birth_place,
				m_mitra_employee.mitra_emp_birth_date,
				m_mitra_employee.mitra_emp_address,
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
				m_mitra_employee.mitra_emp_photo,
				m_mitra_employee.mitra_emp_status,
				m_mitra_employee.mitra_emp_created_by,
				m_mitra_employee.mitra_emp_created_at,
				m_mitra_employee.mitra_emp_updated_by,
				m_mitra_employee.mitra_emp_updated_at,
				m_mitra_employee.mitra_emp_deleted_by,
				m_mitra_employee.mitra_emp_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_employee_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_mitra_branch m_mitra_branch_temp on m_mitra_employee_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_address_temp on m_mitra_employee_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_mitra_employee on m_mitra_employee_temp.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_address on m_mitra_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("mitra_emp_temp_id = ? AND mitra_emp_temp_status = ? AND m_mitra_employee_temp.mitra_temp_id = ?", pid, status, mitra).
		Find(&mitraemployee).Error
	if err != nil {
		return &MitraEmployeeTempData{}, err
	}
	return &mitraemployee, nil
}

//ADDITIONAL
//====================================================================================================================================

func (p *MitraEmployeeTemp) FindMitraEmployeeTempByMitraTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraEmployeeTempPend, error) {
	var err error
	mitraemployee := []MitraEmployeeTempPend{}
	err = db.Debug().Model(&MitraEmployeeTempPend{}).Limit(100).
		Select(`m_mitra_employee_temp.mitra_emp_temp_id,
				m_mitra_employee_temp.mitra_temp_id,
				m_mitra_employee_temp.mitra_branch_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_code,
				m_mitra_employee_temp.mitra_emp_temp_name,
				m_mitra_employee_temp.mitra_emp_temp_gender,
				m_mitra_employee_temp.mitra_emp_temp_birth_place,
				m_mitra_employee_temp.mitra_emp_temp_birth_date,
				m_mitra_employee_temp.mitra_emp_temp_address,
				m_mitra_employee_temp.address_temp_id,
				m_mitra_employee_temp.mitra_emp_temp_photo,
				m_mitra_employee_temp.mitra_emp_temp_reason,
				m_mitra_employee_temp.mitra_emp_temp_status,
				m_mitra_employee_temp.mitra_emp_temp_created_by,
				m_mitra_employee_temp.mitra_emp_temp_created_at`).
		Where("mitra_temp_id = ? AND mitra_emp_temp_status = ?", pid, status).
		Find(&mitraemployee).Error
	if err != nil {
		return []MitraEmployeeTempPend{}, err
	}
	return mitraemployee, nil
}
