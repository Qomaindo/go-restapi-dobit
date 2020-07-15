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

type MitraBranchTemp struct {
	MitraBranchTempID           uint64    `gorm:"column:mitra_branch_temp_id;primary_key;" json:"mitra_branch_temp_id"`
	MitraBranchID               uint64    `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	MitraTempID                 uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	MitraBranchParentTempID     uint64    `gorm:"column:mitra_branch_parent_temp_id" json:"mitra_branch_parent_temp_id"`
	MitraBranchTempCode         string    `gorm:"column:mitra_branch_temp_code" json:"mitra_branch_temp_code"`
	MitraBranchTempName         string    `gorm:"column:mitra_branch_temp_name" json:"mitra_branch_temp_name"`
	MitraBranchTempType         string    `gorm:"column:mitra_branch_temp_type" json:"mitra_branch_temp_type"`
	MitraBranchTempPhone        string    `gorm:"column:mitra_branch_temp_phone" json:"mitra_branch_temp_phone"`
	MitraBranchTempFax          string    `gorm:"column:mitra_branch_temp_fax" json:"mitra_branch_temp_fax"`
	MitraBranchTempAddress      string    `gorm:"column:mitra_branch_temp_address" json:"mitra_branch_temp_address"`
	AddressTempID               uint64    `gorm:"column:address_temp_id" json:"address_temp_id"`
	MitraBranchTempReason       string    `gorm:"column:mitra_branch_temp_reason" json:"mitra_branch_temp_reason"`
	MitraBranchTempStatus       uint64    `gorm:"column:mitra_branch_temp_status;size:2" json:"mitra_branch_temp_status"`
	MitraBranchTempActionBefore uint64    `gorm:"column:mitra_branch_temp_action_before;size:2" json:"mitra_branch_temp_action_before"`
	MitraBranchTempCreatedBy    string    `gorm:"column:mitra_branch_temp_created_by;size:125" json:"mitra_branch_temp_created_by"`
	MitraBranchTempCreatedAt    time.Time `gorm:"column:mitra_branch_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_branch_temp_created_at"`

	MitraBranchCoverageProvinceTemp []MitraBranchCoverageProvinceTempPendData `json:"mitra_branch_cov_prov_temp"`
	MitraBranchCoverageRegencyTemp  []MitraBranchCoverageRegencyTempPendData  `json:"mitra_branch_cov_rgcy_temp"`
	MitraContactPersonTemp          []MitraContactPersonTempPendData          `json:"mitra_cp_temp"`
}

type MitraBranchTempPend struct {
	MitraBranchTempID           uint64    `gorm:"column:mitra_branch_temp_id;primary_key;" json:"mitra_branch_temp_id"`
	MitraBranchID               *uint64   `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	MitraTempID                 uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	MitraBranchParentTempID     uint64    `gorm:"column:mitra_branch_parent_temp_id" json:"mitra_branch_parent_temp_id"`
	MitraBranchTempCode         string    `gorm:"column:mitra_branch_temp_code" json:"mitra_branch_temp_code"`
	MitraBranchTempName         string    `gorm:"column:mitra_branch_temp_name" json:"mitra_branch_temp_name"`
	MitraBranchTempType         string    `gorm:"column:mitra_branch_temp_type" json:"mitra_branch_temp_type"`
	MitraBranchTempPhone        string    `gorm:"column:mitra_branch_temp_phone" json:"mitra_branch_temp_phone"`
	MitraBranchTempFax          string    `gorm:"column:mitra_branch_temp_fax" json:"mitra_branch_temp_fax"`
	MitraBranchTempAddress      string    `gorm:"column:mitra_branch_temp_address" json:"mitra_branch_temp_address"`
	AddressTempID               uint64    `gorm:"column:address_temp_id" json:"address_temp_id"`
	MitraBranchTempReason       string    `gorm:"column:mitra_branch_temp_reason" json:"mitra_branch_temp_reason"`
	MitraBranchTempStatus       uint64    `gorm:"column:mitra_branch_temp_status;size:2" json:"mitra_branch_temp_status"`
	MitraBranchTempActionBefore uint64    `gorm:"column:mitra_branch_temp_action_before;size:2" json:"mitra_branch_temp_action_before"`
	MitraBranchTempCreatedBy    string    `gorm:"column:mitra_branch_temp_created_by;size:125" json:"mitra_branch_temp_created_by"`
	MitraBranchTempCreatedAt    time.Time `gorm:"column:mitra_branch_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_branch_temp_created_at"`

	MitraBranchCoverageProvinceTemp []MitraBranchCoverageProvinceTempPendData `json:"mitra_branch_cov_prov_temp"`
	MitraBranchCoverageRegencyTemp  []MitraBranchCoverageRegencyTempPendData  `json:"mitra_branch_cov_rgcy_temp"`
	MitraContactPersonTemp          []MitraContactPersonTempPendData          `json:"mitra_cp_temp"`
}

type MitraBranchTempPendData struct {
	MitraBranchTempID           uint64    `gorm:"column:mitra_branch_temp_id;primary_key;" json:"mitra_branch_temp_id"`
	MitraBranchID               uint64    `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	MitraTempID                 uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	MitraBranchParentTempID     uint64    `gorm:"column:mitra_branch_parent_temp_id" json:"mitra_branch_parent_temp_id"`
	MitraBranchTempCode         string    `gorm:"column:mitra_branch_temp_code" json:"mitra_branch_temp_code"`
	MitraBranchTempName         string    `gorm:"column:mitra_branch_temp_name" json:"mitra_branch_temp_name"`
	MitraBranchTempType         string    `gorm:"column:mitra_branch_temp_type" json:"mitra_branch_temp_type"`
	MitraBranchTempPhone        string    `gorm:"column:mitra_branch_temp_phone" json:"mitra_branch_temp_phone"`
	MitraBranchTempFax          string    `gorm:"column:mitra_branch_temp_fax" json:"mitra_branch_temp_fax"`
	MitraBranchTempAddress      string    `gorm:"column:mitra_branch_temp_address" json:"mitra_branch_temp_address"`
	AddressTempID               uint64    `gorm:"column:address_temp_id" json:"address_temp_id"`
	CountryTempID               uint64    `gorm:"column:country_temp_id" json:"country_temp_id"`
	CountryTempName             string    `gorm:"column:country_temp_name" json:"country_temp_name"`
	ProvinceTempID              uint64    `gorm:"column:province_temp_id" json:"province_temp_id"`
	ProvinceTempName            string    `gorm:"column:province_temp_name" json:"province_temp_name"`
	RegencyTempID               uint64    `gorm:"column:regency_temp_id" json:"regency_temp_id"`
	RegencyTempName             string    `gorm:"column:regency_temp_name" json:"regency_temp_name"`
	DistrictTempID              uint64    `gorm:"column:district_temp_id" json:"district_temp_id"`
	DistrictTempName            string    `gorm:"column:district_temp_name" json:"district_temp_name"`
	VillageTempID               uint64    `gorm:"column:village_temp_id" json:"village_temp_id"`
	VillageTempName             string    `gorm:"column:village_temp_name;size:255" json:"village_temp_name"`
	MitraBranchTempReason       string    `gorm:"column:mitra_branch_temp_reason" json:"mitra_branch_temp_reason"`
	MitraBranchTempStatus       uint64    `gorm:"column:mitra_branch_temp_status;size:2" json:"mitra_branch_temp_status"`
	MitraBranchTempActionBefore uint64    `gorm:"column:mitra_branch_temp_action_before;size:2" json:"mitra_branch_temp_action_before"`
	MitraBranchTempCreatedBy    string    `gorm:"column:mitra_branch_temp_created_by;size:125" json:"mitra_branch_temp_created_by"`
	MitraBranchTempCreatedAt    time.Time `gorm:"column:mitra_branch_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_branch_temp_created_at"`
}

type MitraBranchTempData struct {
	MitraBranchTempID           uint64    `gorm:"column:mitra_branch_temp_id" json:"mitra_branch_temp_id"`
	MitraTempID                 uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	MitraBranchParentTempID     uint64    `gorm:"column:mitra_branch_parent_temp_id" json:"mitra_branch_parent_temp_id"`
	MitraBranchTempCode         string    `gorm:"column:mitra_branch_temp_code" json:"mitra_branch_temp_code"`
	MitraBranchTempName         string    `gorm:"column:mitra_branch_temp_name" json:"mitra_branch_temp_name"`
	MitraBranchTempType         string    `gorm:"column:mitra_branch_temp_type" json:"mitra_branch_temp_type"`
	MitraBranchTempPhone        string    `gorm:"column:mitra_branch_temp_phone" json:"mitra_branch_temp_phone"`
	MitraBranchTempFax          string    `gorm:"column:mitra_branch_temp_fax" json:"mitra_branch_temp_fax"`
	MitraBranchTempAddress      string    `gorm:"column:mitra_branch_temp_address" json:"mitra_branch_temp_address"`
	AddressTempID               uint64    `gorm:"column:address_temp_id" json:"address_temp_id"`
	CountryTempID               uint64    `gorm:"column:country_temp_id" json:"country_temp_id"`
	CountryTempName             string    `gorm:"column:country_temp_name" json:"country_temp_name"`
	ProvinceTempID              uint64    `gorm:"column:province_temp_id" json:"province_temp_id"`
	ProvinceTempName            string    `gorm:"column:province_temp_name" json:"province_temp_name"`
	RegencyTempID               uint64    `gorm:"column:regency_temp_id" json:"regency_temp_id"`
	RegencyTempName             string    `gorm:"column:regency_temp_name" json:"regency_temp_name"`
	DistrictTempID              uint64    `gorm:"column:district_temp_id" json:"district_temp_id"`
	DistrictTempName            string    `gorm:"column:district_temp_name" json:"district_temp_name"`
	VillageTempID               uint64    `gorm:"column:village_temp_id" json:"village_temp_id"`
	VillageTempName             string    `gorm:"column:village_temp_name;size:255" json:"village_temp_name"`
	MitraBranchTempReason       string    `gorm:"column:mitra_branch_temp_reason" json:"mitra_branch_temp_reason"`
	MitraBranchTempStatus       uint64    `gorm:"column:mitra_branch_temp_status;size:2" json:"mitra_branch_temp_status"`
	MitraBranchTempActionBefore uint64    `gorm:"column:mitra_branch_temp_action_before;size:2" json:"mitra_branch_temp_action_before"`
	MitraBranchTempCreatedBy    string    `gorm:"column:mitra_branch_temp_created_by;size:125" json:"mitra_branch_temp_created_by"`
	MitraBranchTempCreatedAt    time.Time `gorm:"column:mitra_branch_temp_created_at" json:"mitra_branch_temp_created_at"`

	MitraBranchCoverageProvinceTemp []MitraBranchCoverageProvinceTempData `json:"mitra_branch_cov_prov_temp"`
	MitraBranchCoverageRegencyTemp  []MitraBranchCoverageRegencyTempData  `json:"mitra_branch_cov_rgcy_temp"`
	MitraContactPersonTemp          []MitraContactPersonTempPendData      `json:"mitra_cp_temp"`

	MitraBranchID        uint64    `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	MitraBranchParentID  uint64    `gorm:"column:mitra_branch_parent_id" json:"mitra_branch_parent_id"`
	MitraBranchCode      string    `gorm:"column:mitra_branch_code" json:"mitra_branch_code"`
	MitraBranchName      string    `gorm:"column:mitra_branch_name" json:"mitra_branch_name"`
	MitraBranchType      string    `gorm:"column:mitra_branch_type" json:"mitra_branch_type"`
	MitraBranchPhone     string    `gorm:"column:mitra_branch_phone" json:"mitra_branch_phone"`
	MitraBranchFax       string    `gorm:"column:mitra_branch_fax" json:"mitra_branch_fax"`
	MitraBranchAddress   string    `gorm:"column:mitra_branch_address" json:"mitra_branch_address"`
	AddressID            uint64    `gorm:"column:address_id" json:"address_id"`
	CountryID            uint64    `gorm:"column:country_id" json:"country_id"`
	CountryName          string    `gorm:"column:country_name" json:"country_name"`
	ProvinceID           uint64    `gorm:"column:province_id" json:"province_id"`
	ProvinceName         string    `gorm:"column:province_name" json:"province_name"`
	RegencyID            uint64    `gorm:"column:regency_id" json:"regency_id"`
	RegencyName          string    `gorm:"column:regency_name" json:"regency_name"`
	DistrictID           uint64    `gorm:"column:district_id" json:"district_id"`
	DistrictName         string    `gorm:"column:district_name" json:"district_name"`
	VillageID            uint64    `gorm:"column:village_id" json:"village_id"`
	VillageName          string    `gorm:"column:village_name" json:"village_name"`
	MitraBranchStatus    uint64    `gorm:"column:mitra_branch_status" json:"mitra_branch_status"`
	MitraBranchCreatedBy string    `gorm:"column:mitra_branch_created_by" json:"mitra_branch_created_by"`
	MitraBranchCreatedAt time.Time `gorm:"column:mitra_branch_created_at" json:"mitra_branch_created_at"`
	MitraBranchUpdatedBy string    `gorm:"column:mitra_branch_updated_by" json:"mitra_branch_updated_by"`
	MitraBranchUpdatedAt time.Time `gorm:"column:mitra_branch_updated_at" json:"mitra_branch_updated_at"`
	MitraBranchDeletedBy string    `gorm:"column:mitra_branch_deleted_by" json:"mitra_branch_deleted_by"`
	MitraBranchDeletedAt time.Time `gorm:"column:mitra_branch_deleted_at" json:"mitra_branch_deleted_at"`

	MitraBranchCoverageProvince []MitraBranchCoverageProvinceData `json:"mitra_branch_cov_prov"`
	MitraBranchCoverageRegency  []MitraBranchCoverageRegencyData  `json:"mitra_branch_cov_rgcy"`
	MitraContactPerson          []MitraContactPersonData          `json:"mitra_cp"`
}

type ResponseMitraBranchTemps struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    *[]MitraBranchTempData `json:"data"`
}

type ResponseMitraBranchTemp struct {
	Status  int                  `json:"status"`
	Message string               `json:"message"`
	Data    *MitraBranchTempData `json:"data"`
}

type ResponseMitraBranchTempsPend struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    *[]MitraBranchTempPend `json:"data"`
}

type ResponseMitraBranchTempPend struct {
	Status  int                  `json:"status"`
	Message string               `json:"message"`
	Data    *MitraBranchTempPend `json:"data"`
}

type ResponseMitraBranchTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraBranchTemp) TableName() string {
	return "m_mitra_branch_temp"
}

func (MitraBranchTempPend) TableName() string {
	return "m_mitra_branch_temp"
}

func (MitraBranchTempPendData) TableName() string {
	return "m_mitra_branch_temp"
}

func (MitraBranchTempData) TableName() string {
	return "m_mitra_branch_temp"
}

func (p *MitraBranchTemp) Prepare() {
	p.MitraBranchID = p.MitraBranchID
	p.MitraBranchParentTempID = p.MitraBranchParentTempID
	p.MitraBranchTempCode = p.MitraBranchTempCode
	p.MitraBranchTempName = p.MitraBranchTempName
	p.MitraBranchTempType = p.MitraBranchTempType
	p.MitraBranchTempPhone = p.MitraBranchTempPhone
	p.MitraBranchTempFax = p.MitraBranchTempFax
	p.MitraBranchTempAddress = p.MitraBranchTempAddress
	p.AddressTempID = p.AddressTempID
	p.MitraBranchTempReason = p.MitraBranchTempReason
	p.MitraBranchTempStatus = p.MitraBranchTempStatus
	p.MitraBranchTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraBranchTempCreatedBy))
	p.MitraBranchTempCreatedAt = time.Now()
}

func (p *MitraBranchTempPend) Prepare() {
	p.MitraBranchID = nil
	p.MitraBranchParentTempID = p.MitraBranchParentTempID
	p.MitraBranchTempCode = p.MitraBranchTempCode
	p.MitraBranchTempName = p.MitraBranchTempName
	p.MitraBranchTempType = p.MitraBranchTempType
	p.MitraBranchTempPhone = p.MitraBranchTempPhone
	p.MitraBranchTempFax = p.MitraBranchTempFax
	p.MitraBranchTempAddress = p.MitraBranchTempAddress
	p.AddressTempID = p.AddressTempID
	p.MitraBranchTempReason = p.MitraBranchTempReason
	p.MitraBranchTempStatus = p.MitraBranchTempStatus
	p.MitraBranchTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraBranchTempCreatedBy))
	p.MitraBranchTempCreatedAt = time.Now()
}

func (p *MitraBranchTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	case "insertupdate":
		if p.MitraBranchTempName == "" {
			return errors.New("Negara Wajib diisi")
		}
		if p.MitraBranchTempType == "" {
			return errors.New("Provinsi Wajib diisi")
		}
		if p.MitraBranchTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.MitraBranchTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	case "reason":
		if p.MitraBranchTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.MitraBranchTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	default:
		if p.MitraBranchTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil
	}
}

func (p *MitraBranchTemp) SaveMitraBranchTemp(db *gorm.DB) (*MitraBranchTemp, error) {
	var err error
	err = db.Debug().Model(&MitraBranchTemp{}).Create(&p).Error
	if err != nil {
		return &MitraBranchTemp{}, err
	}
	return p, nil
}

func (p *MitraBranchTempPend) SaveMitraBranchTempPend(db *gorm.DB) (*MitraBranchTempPend, error) {
	var err error
	err = db.Debug().Model(&MitraBranchTempPend{}).Create(&p).Error
	if err != nil {
		return &MitraBranchTempPend{}, err
	}
	return p, nil
}

func (p *MitraBranchTemp) FindAllMitraBranchTemp(db *gorm.DB) (*[]MitraBranchTempPend, error) {
	var err error
	mitrabranch := []MitraBranchTempPend{}
	err = db.Debug().Model(&MitraBranchTempPend{}).Limit(100).
		Select(`m_mitra_branch_temp.mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_id,
				m_mitra_branch_temp.mitra_temp_id,
				m_mitra_branch_temp.mitra_branch_parent_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_code,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_branch_temp.mitra_branch_temp_type,
				m_mitra_branch_temp.mitra_branch_temp_phone,
				m_mitra_branch_temp.mitra_branch_temp_fax,
				m_mitra_branch_temp.mitra_branch_temp_address,
				m_mitra_branch_temp.address_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_reason,
				m_mitra_branch_temp.mitra_branch_temp_status,
				m_mitra_branch_temp.mitra_branch_temp_action_before,
				m_mitra_branch_temp.mitra_branch_temp_created_by,
				m_mitra_branch_temp.mitra_branch_temp_created_at`).
		Find(&mitrabranch).Error
	if err != nil {
		return &[]MitraBranchTempPend{}, err
	}
	return &mitrabranch, nil
}

func (p *MitraBranchTemp) FindAllMitraBranchTempPendingActive(db *gorm.DB) (*[]MitraBranchTempPend, error) {
	var err error
	mitrabranch := []MitraBranchTempPend{}
	err = db.Debug().Model(&MitraBranchTempPend{}).Limit(100).
		Select(`m_mitra_branch_temp.mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_id,
				m_mitra_branch_temp.mitra_temp_id,
				m_mitra_branch_temp.mitra_branch_parent_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_code,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_branch_temp.mitra_branch_temp_type,
				m_mitra_branch_temp.mitra_branch_temp_phone,
				m_mitra_branch_temp.mitra_branch_temp_fax,
				m_mitra_branch_temp.mitra_branch_temp_address,
				m_mitra_branch_temp.address_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_reason,
				m_mitra_branch_temp.mitra_branch_temp_status,
				m_mitra_branch_temp.mitra_branch_temp_action_before,
				m_mitra_branch_temp.mitra_branch_temp_created_by,
				m_mitra_branch_temp.mitra_branch_temp_created_at`).
		Where("mitra_branch_temp_status = ?", 11).
		Find(&mitrabranch).Error
	if err != nil {
		return &[]MitraBranchTempPend{}, err
	}
	return &mitrabranch, nil
}

func (p *MitraBranchTemp) FindAllMitraBranchTempStatus(db *gorm.DB, status uint64) (*[]MitraBranchTempData, error) {
	var err error
	mitrabranch := []MitraBranchTempData{}
	err = db.Debug().Model(&MitraBranchTempData{}).Limit(100).
		Select(`m_mitra_branch_temp.mitra_branch_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_parent_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_code,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_branch_temp.mitra_branch_temp_type,
				m_mitra_branch_temp.mitra_branch_temp_phone,
				m_mitra_branch_temp.mitra_branch_temp_fax,
				m_mitra_branch_temp.mitra_branch_temp_address,
				m_mitra_branch_temp.address_temp_id,
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
				m_mitra_branch_temp.mitra_branch_temp_reason,
				m_mitra_branch_temp.mitra_branch_temp_status,
				m_mitra_branch_temp.mitra_branch_temp_action_before,
				m_mitra_branch_temp.mitra_branch_temp_created_by,
				m_mitra_branch_temp.mitra_branch_temp_created_at,
				m_mitra_branch.mitra_branch_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_mitra_branch.mitra_branch_parent_id,
				m_mitra_branch.mitra_branch_code,
				m_mitra_branch.mitra_branch_name,
				m_mitra_branch.mitra_branch_type,
				m_mitra_branch.mitra_branch_phone,
				m_mitra_branch.mitra_branch_fax,
				m_mitra_branch.mitra_branch_address,
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
				m_mitra_branch.mitra_branch_status,
				m_mitra_branch.mitra_branch_created_by,
				m_mitra_branch.mitra_branch_created_at,
				m_mitra_branch.mitra_branch_updated_by,
				m_mitra_branch.mitra_branch_updated_at,
				m_mitra_branch.mitra_branch_deleted_by,
				m_mitra_branch.mitra_branch_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_branch_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_address_temp on m_mitra_branch_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_mitra_branch on m_mitra_branch_temp.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra on m_mitra_branch.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_address on m_mitra_branch.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("mitra_branch_temp_status = ?", status).
		Find(&mitrabranch).Error
	if err != nil {
		return &[]MitraBranchTempData{}, err
	}
	return &mitrabranch, nil
}

func (p *MitraBranchTemp) FindMitraBranchTempDataByID(db *gorm.DB, pid uint64) (*MitraBranchTemp, error) {
	var err error
	err = db.Debug().Model(&MitraBranchTemp{}).Where("mitra_branch_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraBranchTemp{}, err
	}
	return p, nil
}

func (p *MitraBranchTemp) FindMitraBranchTempByIDPendingActive(db *gorm.DB, pid uint64) (*MitraBranchTempPend, error) {
	var err error
	mitrabranch := MitraBranchTempPend{}
	err = db.Debug().Model(&MitraBranchTempPend{}).Limit(100).
		Select(`m_mitra_branch_temp.mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_id,
				m_mitra_branch_temp.mitra_temp_id,
				m_mitra_branch_temp.mitra_branch_parent_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_code,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_branch_temp.mitra_branch_temp_type,
				m_mitra_branch_temp.mitra_branch_temp_phone,
				m_mitra_branch_temp.mitra_branch_temp_fax,
				m_mitra_branch_temp.mitra_branch_temp_address,
				m_mitra_branch_temp.address_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_reason,
				m_mitra_branch_temp.mitra_branch_temp_status,
				m_mitra_branch_temp.mitra_branch_temp_action_before,
				m_mitra_branch_temp.mitra_branch_temp_created_by,
				m_mitra_branch_temp.mitra_branch_temp_created_at`).
		Where("mitra_branch_temp_id = ? AND mitra_branch_temp_status = ?", pid, 11).
		Find(&mitrabranch).Error
	if err != nil {
		return &MitraBranchTempPend{}, err
	}
	return &mitrabranch, nil
}

func (p *MitraBranchTemp) FindMitraBranchTempByID(db *gorm.DB, pid uint64) (*MitraBranchTempData, error) {
	var err error
	mitrabranch := MitraBranchTempData{}
	err = db.Debug().Model(&MitraBranchTempData{}).Limit(100).
		Select(`m_mitra_branch_temp.mitra_branch_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_parent_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_code,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_branch_temp.mitra_branch_temp_type,
				m_mitra_branch_temp.mitra_branch_temp_phone,
				m_mitra_branch_temp.mitra_branch_temp_fax,
				m_mitra_branch_temp.mitra_branch_temp_address,
				m_mitra_branch_temp.address_temp_id,
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
				m_mitra_branch_temp.mitra_branch_temp_reason,
				m_mitra_branch_temp.mitra_branch_temp_status,
				m_mitra_branch_temp.mitra_branch_temp_action_before,
				m_mitra_branch_temp.mitra_branch_temp_created_by,
				m_mitra_branch_temp.mitra_branch_temp_created_at,
				m_mitra_branch.mitra_branch_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_mitra_branch.mitra_branch_parent_id,
				m_mitra_branch.mitra_branch_code,
				m_mitra_branch.mitra_branch_name,
				m_mitra_branch.mitra_branch_type,
				m_mitra_branch.mitra_branch_phone,
				m_mitra_branch.mitra_branch_fax,
				m_mitra_branch.mitra_branch_address,
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
				m_mitra_branch.mitra_branch_status,
				m_mitra_branch.mitra_branch_created_by,
				m_mitra_branch.mitra_branch_created_at,
				m_mitra_branch.mitra_branch_updated_by,
				m_mitra_branch.mitra_branch_updated_at,
				m_mitra_branch.mitra_branch_deleted_by,
				m_mitra_branch.mitra_branch_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_branch_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_address_temp on m_mitra_branch_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_mitra_branch on m_mitra_branch_temp.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra on m_mitra_branch.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_address on m_mitra_branch.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("mitra_branch_temp_id = ?", pid).
		Find(&mitrabranch).Error
	if err != nil {
		return &MitraBranchTempData{}, err
	}
	return &mitrabranch, nil
}

func (p *MitraBranchTemp) FindMitraBranchTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraBranchTempData, error) {
	var err error
	mitrabranch := MitraBranchTempData{}
	err = db.Debug().Model(&MitraBranchTempData{}).Limit(100).
		Select(`m_mitra_branch_temp.mitra_branch_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_parent_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_code,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_branch_temp.mitra_branch_temp_type,
				m_mitra_branch_temp.mitra_branch_temp_phone,
				m_mitra_branch_temp.mitra_branch_temp_fax,
				m_mitra_branch_temp.mitra_branch_temp_address,
				m_mitra_branch_temp.address_temp_id,
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
				m_mitra_branch_temp.mitra_branch_temp_reason,
				m_mitra_branch_temp.mitra_branch_temp_status,
				m_mitra_branch_temp.mitra_branch_temp_action_before,
				m_mitra_branch_temp.mitra_branch_temp_created_by,
				m_mitra_branch_temp.mitra_branch_temp_created_at,
				m_mitra_branch.mitra_branch_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_mitra_branch.mitra_branch_parent_id,
				m_mitra_branch.mitra_branch_code,
				m_mitra_branch.mitra_branch_name,
				m_mitra_branch.mitra_branch_type,
				m_mitra_branch.mitra_branch_phone,
				m_mitra_branch.mitra_branch_fax,
				m_mitra_branch.mitra_branch_address,
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
				m_mitra_branch.mitra_branch_status,
				m_mitra_branch.mitra_branch_created_by,
				m_mitra_branch.mitra_branch_created_at,
				m_mitra_branch.mitra_branch_updated_by,
				m_mitra_branch.mitra_branch_updated_at,
				m_mitra_branch.mitra_branch_deleted_by,
				m_mitra_branch.mitra_branch_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_branch_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_address_temp on m_mitra_branch_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_mitra_branch on m_mitra_branch_temp.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra on m_mitra_branch.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_address on m_mitra_branch.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("mitra_branch_temp_id = ? AND mitra_branch_temp_status = ?", pid, status).
		Find(&mitrabranch).Error
	if err != nil {
		return &MitraBranchTempData{}, err
	}
	return &mitrabranch, nil
}

func (p *MitraBranchTemp) UpdateMitraBranchTemp(db *gorm.DB, pid uint64) (*MitraBranchTemp, error) {
	var err error
	db = db.Debug().Model(&MitraBranchTemp{}).Where("mitra_branch_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_branch_parent_temp_id": p.MitraBranchParentTempID,
			"mitra_branch_temp_code":      p.MitraBranchTempCode,
			"mitra_branch_temp_name":      p.MitraBranchTempName,
			"mitra_branch_temp_type":      p.MitraBranchTempType,
			"mitra_branch_temp_phone":     p.MitraBranchTempPhone,
			"mitra_branch_temp_fax":       p.MitraBranchTempFax,
			"mitra_branch_temp_address":   p.MitraBranchTempAddress,
			"address_temp_id":             p.AddressTempID,
			"mitra_branch_temp_reason":    p.MitraBranchTempReason,
		},
	)
	err = db.Debug().Model(&MitraBranchTemp{}).Where("mitra_branch_temp_id = ?", pid).Error
	if err != nil {
		return &MitraBranchTemp{}, err
	}
	return p, nil
}

func (p *MitraBranchTemp) UpdateMitraBranchTempStatus(db *gorm.DB, pid uint64) (*MitraBranchTemp, error) {
	var err error
	db = db.Debug().Model(&MitraBranchTemp{}).Where("mitra_branch_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_branch_temp_reason": p.MitraBranchTempReason,
			"mitra_branch_temp_status": p.MitraBranchTempStatus,
		},
	)
	err = db.Debug().Model(&MitraBranchTemp{}).Where("mitra_branch_temp_id = ?", pid).Error
	if err != nil {
		return &MitraBranchTemp{}, err
	}
	return p, nil
}

func (p *MitraBranchTemp) DeleteMitraBranchTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraBranchTemp{}).Where("mitra_branch_temp_id = ?", pid).Take(&MitraBranchTemp{}).Delete(&MitraBranchTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraBranchTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//BY MITRA ID
//===================================================================================================================================

func (p *MitraBranchTemp) FindAllMitraBranchTempMitra(db *gorm.DB, mitra uint64) (*[]MitraBranchTempPend, error) {
	var err error
	mitrabranch := []MitraBranchTempPend{}
	err = db.Debug().Model(&MitraBranchTempPend{}).Limit(100).
		Select(`m_mitra_branch_temp.mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_id,
				m_mitra_branch_temp.mitra_temp_id,
				m_mitra_branch_temp.mitra_branch_parent_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_code,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_branch_temp.mitra_branch_temp_type,
				m_mitra_branch_temp.mitra_branch_temp_phone,
				m_mitra_branch_temp.mitra_branch_temp_fax,
				m_mitra_branch_temp.mitra_branch_temp_address,
				m_mitra_branch_temp.address_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_reason,
				m_mitra_branch_temp.mitra_branch_temp_status,
				m_mitra_branch_temp.mitra_branch_temp_action_before,
				m_mitra_branch_temp.mitra_branch_temp_created_by,
				m_mitra_branch_temp.mitra_branch_temp_created_at`).
		Where("m_mitra_branch_temp.mitra_temp_id = ?", mitra).
		Find(&mitrabranch).Error
	if err != nil {
		return &[]MitraBranchTempPend{}, err
	}
	return &mitrabranch, nil
}

func (p *MitraBranchTemp) FindAllMitraBranchTempPendingActiveMitra(db *gorm.DB, mitra uint64) (*[]MitraBranchTempPend, error) {
	var err error
	mitrabranch := []MitraBranchTempPend{}
	err = db.Debug().Model(&MitraBranchTempPend{}).Limit(100).
		Select(`m_mitra_branch_temp.mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_id,
				m_mitra_branch_temp.mitra_temp_id,
				m_mitra_branch_temp.mitra_branch_parent_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_code,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_branch_temp.mitra_branch_temp_type,
				m_mitra_branch_temp.mitra_branch_temp_phone,
				m_mitra_branch_temp.mitra_branch_temp_fax,
				m_mitra_branch_temp.mitra_branch_temp_address,
				m_mitra_branch_temp.address_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_reason,
				m_mitra_branch_temp.mitra_branch_temp_status,
				m_mitra_branch_temp.mitra_branch_temp_action_before,
				m_mitra_branch_temp.mitra_branch_temp_created_by,
				m_mitra_branch_temp.mitra_branch_temp_created_at`).
		Where("mitra_branch_temp_status = ? AND m_mitra_branch_temp.mitra_temp_id = ?", 11, mitra).
		Find(&mitrabranch).Error
	if err != nil {
		return &[]MitraBranchTempPend{}, err
	}
	return &mitrabranch, nil
}

func (p *MitraBranchTemp) FindAllMitraBranchTempStatusMitra(db *gorm.DB, status uint64, mitra uint64) (*[]MitraBranchTempData, error) {
	var err error
	mitrabranch := []MitraBranchTempData{}
	err = db.Debug().Model(&MitraBranchTempData{}).Limit(100).
		Select(`m_mitra_branch_temp.mitra_branch_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_parent_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_code,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_branch_temp.mitra_branch_temp_type,
				m_mitra_branch_temp.mitra_branch_temp_phone,
				m_mitra_branch_temp.mitra_branch_temp_fax,
				m_mitra_branch_temp.mitra_branch_temp_address,
				m_mitra_branch_temp.address_temp_id,
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
				m_mitra_branch_temp.mitra_branch_temp_reason,
				m_mitra_branch_temp.mitra_branch_temp_status,
				m_mitra_branch_temp.mitra_branch_temp_action_before,
				m_mitra_branch_temp.mitra_branch_temp_created_by,
				m_mitra_branch_temp.mitra_branch_temp_created_at,
				m_mitra_branch.mitra_branch_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_mitra_branch.mitra_branch_parent_id,
				m_mitra_branch.mitra_branch_code,
				m_mitra_branch.mitra_branch_name,
				m_mitra_branch.mitra_branch_type,
				m_mitra_branch.mitra_branch_phone,
				m_mitra_branch.mitra_branch_fax,
				m_mitra_branch.mitra_branch_address,
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
				m_mitra_branch.mitra_branch_status,
				m_mitra_branch.mitra_branch_created_by,
				m_mitra_branch.mitra_branch_created_at,
				m_mitra_branch.mitra_branch_updated_by,
				m_mitra_branch.mitra_branch_updated_at,
				m_mitra_branch.mitra_branch_deleted_by,
				m_mitra_branch.mitra_branch_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_branch_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_address_temp on m_mitra_branch_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_mitra_branch on m_mitra_branch_temp.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra on m_mitra_branch.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_address on m_mitra_branch.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("mitra_branch_temp_status = ? AND m_mitra_branch_temp.mitra_temp_id = ?", status, mitra).
		Find(&mitrabranch).Error
	if err != nil {
		return &[]MitraBranchTempData{}, err
	}
	return &mitrabranch, nil
}

func (p *MitraBranchTemp) FindMitraBranchTempByIDPendingActiveMitra(db *gorm.DB, pid uint64, mitra uint64) (*MitraBranchTempPend, error) {
	var err error
	mitrabranch := MitraBranchTempPend{}
	err = db.Debug().Model(&MitraBranchTempPend{}).Limit(100).
		Select(`m_mitra_branch_temp.mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_id,
				m_mitra_branch_temp.mitra_temp_id,
				m_mitra_branch_temp.mitra_branch_parent_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_code,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_branch_temp.mitra_branch_temp_type,
				m_mitra_branch_temp.mitra_branch_temp_phone,
				m_mitra_branch_temp.mitra_branch_temp_fax,
				m_mitra_branch_temp.mitra_branch_temp_address,
				m_mitra_branch_temp.address_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_reason,
				m_mitra_branch_temp.mitra_branch_temp_status,
				m_mitra_branch_temp.mitra_branch_temp_action_before,
				m_mitra_branch_temp.mitra_branch_temp_created_by,
				m_mitra_branch_temp.mitra_branch_temp_created_at`).
		Where("mitra_branch_temp_id = ? AND mitra_branch_temp_status = ? AND m_mitra_branch_temp.mitra_temp_id = ?", pid, 11, mitra).
		Find(&mitrabranch).Error
	if err != nil {
		return &MitraBranchTempPend{}, err
	}
	return &mitrabranch, nil
}

func (p *MitraBranchTemp) FindMitraBranchTempByIDMitra(db *gorm.DB, pid uint64, mitra uint64) (*MitraBranchTempData, error) {
	var err error
	mitrabranch := MitraBranchTempData{}
	err = db.Debug().Model(&MitraBranchTempData{}).Limit(100).
		Select(`m_mitra_branch_temp.mitra_branch_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_parent_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_code,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_branch_temp.mitra_branch_temp_type,
				m_mitra_branch_temp.mitra_branch_temp_phone,
				m_mitra_branch_temp.mitra_branch_temp_fax,
				m_mitra_branch_temp.mitra_branch_temp_address,
				m_mitra_branch_temp.address_temp_id,
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
				m_mitra_branch_temp.mitra_branch_temp_reason,
				m_mitra_branch_temp.mitra_branch_temp_status,
				m_mitra_branch_temp.mitra_branch_temp_action_before,
				m_mitra_branch_temp.mitra_branch_temp_created_by,
				m_mitra_branch_temp.mitra_branch_temp_created_at,
				m_mitra_branch.mitra_branch_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_mitra_branch.mitra_branch_parent_id,
				m_mitra_branch.mitra_branch_code,
				m_mitra_branch.mitra_branch_name,
				m_mitra_branch.mitra_branch_type,
				m_mitra_branch.mitra_branch_phone,
				m_mitra_branch.mitra_branch_fax,
				m_mitra_branch.mitra_branch_address,
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
				m_mitra_branch.mitra_branch_status,
				m_mitra_branch.mitra_branch_created_by,
				m_mitra_branch.mitra_branch_created_at,
				m_mitra_branch.mitra_branch_updated_by,
				m_mitra_branch.mitra_branch_updated_at,
				m_mitra_branch.mitra_branch_deleted_by,
				m_mitra_branch.mitra_branch_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_branch_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_address_temp on m_mitra_branch_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_mitra_branch on m_mitra_branch_temp.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra on m_mitra_branch.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_address on m_mitra_branch.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("mitra_branch_temp_id = ? AND m_mitra_branch_temp.mitra_temp_id = ?", pid, mitra).
		Find(&mitrabranch).Error
	if err != nil {
		return &MitraBranchTempData{}, err
	}
	return &mitrabranch, nil
}

func (p *MitraBranchTemp) FindMitraBranchTempStatusByIDMitra(db *gorm.DB, pid uint64, status uint64, mitra uint64) (*MitraBranchTempData, error) {
	var err error
	mitrabranch := MitraBranchTempData{}
	err = db.Debug().Model(&MitraBranchTempData{}).Limit(100).
		Select(`m_mitra_branch_temp.mitra_branch_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_parent_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_code,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_branch_temp.mitra_branch_temp_type,
				m_mitra_branch_temp.mitra_branch_temp_phone,
				m_mitra_branch_temp.mitra_branch_temp_fax,
				m_mitra_branch_temp.mitra_branch_temp_address,
				m_mitra_branch_temp.address_temp_id,
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
				m_mitra_branch_temp.mitra_branch_temp_reason,
				m_mitra_branch_temp.mitra_branch_temp_status,
				m_mitra_branch_temp.mitra_branch_temp_action_before,
				m_mitra_branch_temp.mitra_branch_temp_created_by,
				m_mitra_branch_temp.mitra_branch_temp_created_at,
				m_mitra_branch.mitra_branch_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_mitra_branch.mitra_branch_parent_id,
				m_mitra_branch.mitra_branch_code,
				m_mitra_branch.mitra_branch_name,
				m_mitra_branch.mitra_branch_type,
				m_mitra_branch.mitra_branch_phone,
				m_mitra_branch.mitra_branch_fax,
				m_mitra_branch.mitra_branch_address,
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
				m_mitra_branch.mitra_branch_status,
				m_mitra_branch.mitra_branch_created_by,
				m_mitra_branch.mitra_branch_created_at,
				m_mitra_branch.mitra_branch_updated_by,
				m_mitra_branch.mitra_branch_updated_at,
				m_mitra_branch.mitra_branch_deleted_by,
				m_mitra_branch.mitra_branch_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_branch_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_address_temp on m_mitra_branch_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_mitra_branch on m_mitra_branch_temp.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra on m_mitra_branch.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_address on m_mitra_branch.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("mitra_branch_temp_id = ? AND mitra_branch_temp_status = ? AND m_mitra_branch_temp.mitra_temp_id = ?", pid, status, mitra).
		Find(&mitrabranch).Error
	if err != nil {
		return &MitraBranchTempData{}, err
	}
	return &mitrabranch, nil
}

//ADDITIONAL
//====================================================================================================================================

func (p *MitraBranchTemp) FindMitraBranchTempPendByMitraTempID(db *gorm.DB, pid uint64) (*MitraBranchTempPendData, error) {
	var err error
	mitrabranch := MitraBranchTempPendData{}
	err = db.Debug().Model(&MitraBranchTempPendData{}).Limit(100).
		Select(`m_mitra_branch_temp.mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_id,
				m_mitra_branch_temp.mitra_temp_id,
				m_mitra_branch_temp.mitra_branch_parent_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_code,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_branch_temp.mitra_branch_temp_type,
				m_mitra_branch_temp.mitra_branch_temp_phone,
				m_mitra_branch_temp.mitra_branch_temp_fax,
				m_mitra_branch_temp.mitra_branch_temp_address,
				m_mitra_branch_temp.address_temp_id,
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
				m_mitra_branch_temp.mitra_branch_temp_reason,
				m_mitra_branch_temp.mitra_branch_temp_status,
				m_mitra_branch_temp.mitra_branch_temp_action_before,
				m_mitra_branch_temp.mitra_branch_temp_created_by,
				m_mitra_branch_temp.mitra_branch_temp_created_at`).
		Joins("JOIN m_address_temp on m_mitra_branch_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Where("mitra_temp_id = ?", pid).
		Find(&mitrabranch).Error
	if err != nil {
		return &MitraBranchTempPendData{}, err
	}
	return &mitrabranch, nil
}

func (p *MitraBranchTemp) FindMitraBranchTempPendStatusByMitraTempID(db *gorm.DB, pid uint64, status uint64) (*MitraBranchTempPendData, error) {
	var err error
	mitrabranch := MitraBranchTempPendData{}
	err = db.Debug().Model(&MitraBranchTempPendData{}).Limit(100).
		Select(`m_mitra_branch_temp.mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_id,
				m_mitra_branch_temp.mitra_temp_id,
				m_mitra_branch_temp.mitra_branch_parent_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_code,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_branch_temp.mitra_branch_temp_type,
				m_mitra_branch_temp.mitra_branch_temp_phone,
				m_mitra_branch_temp.mitra_branch_temp_fax,
				m_mitra_branch_temp.mitra_branch_temp_address,
				m_mitra_branch_temp.address_temp_id,
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
				m_mitra_branch_temp.mitra_branch_temp_reason,
				m_mitra_branch_temp.mitra_branch_temp_status,
				m_mitra_branch_temp.mitra_branch_temp_action_before,
				m_mitra_branch_temp.mitra_branch_temp_created_by,
				m_mitra_branch_temp.mitra_branch_temp_created_at`).
		Joins("JOIN m_address_temp on m_mitra_branch_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Where("mitra_temp_id = ? AND mitra_branch_temp_status = ?", pid, status).
		Find(&mitrabranch).Error
	if err != nil {
		return &MitraBranchTempPendData{}, err
	}
	return &mitrabranch, nil
}

func (p *MitraBranchTemp) FindMitraBranchTempByMitraTempID(db *gorm.DB, pid uint64) (*MitraBranchTempData, error) {
	var err error
	mitrabranch := MitraBranchTempData{}
	err = db.Debug().Model(&MitraBranchTempData{}).Limit(100).
		Select(`m_mitra_branch_temp.mitra_branch_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_parent_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_code,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_branch_temp.mitra_branch_temp_type,
				m_mitra_branch_temp.mitra_branch_temp_phone,
				m_mitra_branch_temp.mitra_branch_temp_fax,
				m_mitra_branch_temp.mitra_branch_temp_address,
				m_mitra_branch_temp.address_temp_id,
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
				m_mitra_branch_temp.mitra_branch_temp_reason,
				m_mitra_branch_temp.mitra_branch_temp_status,
				m_mitra_branch_temp.mitra_branch_temp_action_before,
				m_mitra_branch_temp.mitra_branch_temp_created_by,
				m_mitra_branch_temp.mitra_branch_temp_created_at,
				m_mitra_branch.mitra_branch_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_mitra_branch.mitra_branch_parent_id,
				m_mitra_branch.mitra_branch_code,
				m_mitra_branch.mitra_branch_name,
				m_mitra_branch.mitra_branch_type,
				m_mitra_branch.mitra_branch_phone,
				m_mitra_branch.mitra_branch_fax,
				m_mitra_branch.mitra_branch_address,
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
				m_mitra_branch.mitra_branch_status,
				m_mitra_branch.mitra_branch_created_by,
				m_mitra_branch.mitra_branch_created_at,
				m_mitra_branch.mitra_branch_updated_by,
				m_mitra_branch.mitra_branch_updated_at,
				m_mitra_branch.mitra_branch_deleted_by,
				m_mitra_branch.mitra_branch_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_branch_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_address_temp on m_mitra_branch_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_mitra_branch on m_mitra_branch_temp.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra on m_mitra_branch.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_address on m_mitra_branch.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("m_mitra_branch_temp.mitra_temp_id = ?", pid).
		Find(&mitrabranch).Error
	if err != nil {
		return &MitraBranchTempData{}, err
	}
	return &mitrabranch, nil
}

func (p *MitraBranchTemp) FindMitraBranchTempStatusByMitraTempID(db *gorm.DB, pid uint64, status uint64) (*MitraBranchTempData, error) {
	var err error
	mitrabranch := MitraBranchTempData{}
	err = db.Debug().Model(&MitraBranchTempData{}).Limit(100).
		Select(`m_mitra_branch_temp.mitra_branch_temp_id,
				m_mitra_temp.mitra_temp_id,
				m_mitra_temp.mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_parent_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_code,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_branch_temp.mitra_branch_temp_type,
				m_mitra_branch_temp.mitra_branch_temp_phone,
				m_mitra_branch_temp.mitra_branch_temp_fax,
				m_mitra_branch_temp.mitra_branch_temp_address,
				m_mitra_branch_temp.address_temp_id,
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
				m_mitra_branch_temp.mitra_branch_temp_reason,
				m_mitra_branch_temp.mitra_branch_temp_status,
				m_mitra_branch_temp.mitra_branch_temp_action_before,
				m_mitra_branch_temp.mitra_branch_temp_created_by,
				m_mitra_branch_temp.mitra_branch_temp_created_at,
				m_mitra_branch.mitra_branch_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_mitra_branch.mitra_branch_parent_id,
				m_mitra_branch.mitra_branch_code,
				m_mitra_branch.mitra_branch_name,
				m_mitra_branch.mitra_branch_type,
				m_mitra_branch.mitra_branch_phone,
				m_mitra_branch.mitra_branch_fax,
				m_mitra_branch.mitra_branch_address,
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
				m_mitra_branch.mitra_branch_status,
				m_mitra_branch.mitra_branch_created_by,
				m_mitra_branch.mitra_branch_created_at,
				m_mitra_branch.mitra_branch_updated_by,
				m_mitra_branch.mitra_branch_updated_at,
				m_mitra_branch.mitra_branch_deleted_by,
				m_mitra_branch.mitra_branch_deleted_at`).
		Joins("JOIN m_mitra_temp on m_mitra_branch_temp.mitra_temp_id=m_mitra_temp.mitra_temp_id").
		Joins("JOIN m_address_temp on m_mitra_branch_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_mitra_branch on m_mitra_branch_temp.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra on m_mitra_branch.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_address on m_mitra_branch.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("m_mitra_branch_temp.mitra_temp_id = ? AND m_mitra_branch_temp.mitra_branch_temp_status = ?", pid, status).
		Find(&mitrabranch).Error
	if err != nil {
		return &MitraBranchTempData{}, err
	}
	return &mitrabranch, nil
}

func (p *MitraBranchTemp) FindMitraBranchTempPendactStatusByMitraTempID(db *gorm.DB, pid uint64, status uint64) (*MitraBranchTempData, error) {
	var err error
	mitrabranch := MitraBranchTempData{}
	err = db.Debug().Model(&MitraBranchTempData{}).Limit(100).
		Select(`m_mitra_branch_temp.mitra_branch_temp_id,
				m_mitra_temp.mitra_temp_id,
				m_mitra_temp.mitra_temp_name,
				m_mitra_branch_temp.mitra_branch_parent_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_code,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_branch_temp.mitra_branch_temp_type,
				m_mitra_branch_temp.mitra_branch_temp_phone,
				m_mitra_branch_temp.mitra_branch_temp_fax,
				m_mitra_branch_temp.mitra_branch_temp_address,
				m_mitra_branch_temp.address_temp_id,
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
				m_mitra_branch_temp.mitra_branch_temp_reason,
				m_mitra_branch_temp.mitra_branch_temp_status,
				m_mitra_branch_temp.mitra_branch_temp_action_before,
				m_mitra_branch_temp.mitra_branch_temp_created_by,
				m_mitra_branch_temp.mitra_branch_temp_created_at`).
		Joins("JOIN m_mitra_temp on m_mitra_branch_temp.mitra_temp_id=m_mitra_temp.mitra_temp_id").
		Joins("JOIN m_address_temp on m_mitra_branch_temp.address_temp_id=m_address_temp.address_temp_id").
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Where("m_mitra_branch_temp.mitra_temp_id = ? AND m_mitra_branch_temp.mitra_branch_temp_status = ?", pid, status).
		Find(&mitrabranch).Error
	if err != nil {
		return &MitraBranchTempData{}, err
	}
	return &mitrabranch, nil
}
