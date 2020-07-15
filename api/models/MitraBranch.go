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

type MitraBranch struct {
	MitraBranchID        uint64     `gorm:"column:mitra_branch_id;primary_key;" json:"mitra_branch_id"`
	MitraID              uint64     `gorm:"column:mitra_id" json:"mitra_id"`
	MitraBranchParentID  uint64     `gorm:"column:mitra_branch_parent_id" json:"mitra_branch_parent_id"`
	MitraBranchCode      string     `gorm:"column:mitra_branch_code" json:"mitra_branch_code"`
	MitraBranchName      string     `gorm:"column:mitra_branch_name" json:"mitra_branch_name"`
	MitraBranchType      string     `gorm:"column:mitra_branch_type" json:"mitra_branch_type"`
	MitraBranchPhone     string     `gorm:"column:mitra_branch_phone" json:"mitra_branch_phone"`
	MitraBranchFax       string     `gorm:"column:mitra_branch_fax" json:"mitra_branch_fax"`
	MitraBranchAddress   string     `gorm:"column:mitra_branch_address" json:"mitra_branch_address"`
	AddressID            uint64     `gorm:"column:address_id" json:"address_id"`
	MitraBranchStatus    uint64     `gorm:"column:mitra_branch_status;size:2" json:"mitra_branch_status"`
	MitraBranchCreatedBy string     `gorm:"column:mitra_branch_created_by;size:125" json:"mitra_branch_created_by"`
	MitraBranchCreatedAt time.Time  `gorm:"column:mitra_branch_created_at;default:CURRENT_TIMESTAMP" json:"mitra_branch_created_at"`
	MitraBranchUpdatedBy string     `gorm:"column:mitra_branch_updated_by;size:125" json:"mitra_branch_updated_by"`
	MitraBranchUpdatedAt *time.Time `gorm:"column:mitra_branch_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_branch_created_at"`
	MitraBranchDeletedBy string     `gorm:"column:mitra_branch_deleted_by;size:125" json:"mitra_branch_deleted_by"`
	MitraBranchDeletedAt *time.Time `gorm:"column:mitra_branch_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_branch_deleted_at"`
}

type MitraBranchData struct {
	MitraBranchID        uint64     `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	MitraID              uint64     `gorm:"column:mitra_id" json:"mitra_id"`
	MitraName            string     `gorm:"column:mitra_name" json:"mitra_name"`
	MitraBranchParentID  uint64     `gorm:"column:mitra_branch_parent_id" json:"mitra_branch_parent_id"`
	MitraBranchCode      string     `gorm:"column:mitra_branch_code" json:"mitra_branch_code"`
	MitraBranchName      string     `gorm:"column:mitra_branch_name" json:"mitra_branch_name"`
	MitraBranchType      string     `gorm:"column:mitra_branch_type" json:"mitra_branch_type"`
	MitraBranchPhone     string     `gorm:"column:mitra_branch_phone" json:"mitra_branch_phone"`
	MitraBranchFax       string     `gorm:"column:mitra_branch_fax" json:"mitra_branch_fax"`
	MitraBranchAddress   string     `gorm:"column:mitra_branch_address" json:"mitra_branch_address"`
	AddressID            uint64     `gorm:"column:address_id" json:"address_id"`
	CountryID            uint64     `gorm:"column:country_id" json:"country_id"`
	CountryName          string     `gorm:"column:country_name" json:"country_name"`
	ProvinceID           uint64     `gorm:"column:province_id" json:"province_id"`
	ProvinceName         string     `gorm:"column:province_name" json:"province_name"`
	RegencyID            uint64     `gorm:"column:regency_id" json:"regency_id"`
	RegencyName          string     `gorm:"column:regency_name" json:"regency_name"`
	DistrictID           uint64     `gorm:"column:district_id" json:"district_id"`
	DistrictName         string     `gorm:"column:district_name" json:"district_name"`
	VillageID            uint64     `gorm:"column:village_id" json:"village_id"`
	VillageName          string     `gorm:"column:village_name" json:"village_name"`
	MitraBranchStatus    uint64     `gorm:"column:mitra_branch_status" json:"mitra_branch_status"`
	MitraBranchCreatedBy string     `gorm:"column:mitra_branch_created_by" json:"mitra_branch_created_by"`
	MitraBranchCreatedAt time.Time  `gorm:"column:mitra_branch_created_at" json:"mitra_branch_created_at"`
	MitraBranchUpdatedBy string     `gorm:"column:mitra_branch_updated_by" json:"mitra_branch_updated_by"`
	MitraBranchUpdatedAt *time.Time `gorm:"column:mitra_branch_updated_at" json:"mitra_branch_updated_at"`
	MitraBranchDeletedBy string     `gorm:"column:mitra_branch_deleted_by" json:"mitra_branch_deleted_by"`
	MitraBranchDeletedAt *time.Time `gorm:"column:mitra_branch_deleted_at" json:"mitra_branch_deleted_at"`

	MitraBranchCoverageProvince []MitraBranchCoverageProvinceData `json:"mitra_branch_cov_prov"`
	MitraBranchCoverageRegency  []MitraBranchCoverageRegencyData  `json:"mitra_branch_cov_rgcy"`
	MitraContactPerson          []MitraContactPersonData          `json:"mitra_cp"`
}

type ResponseMitraBranchs struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    *[]MitraBranchData `json:"data"`
}

type ResponseMitraBranch struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    *MitraBranchData `json:"data"`
}

type ResponseMitraBranchCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraBranch) TableName() string {
	return "m_mitra_branch"
}

func (MitraBranchData) TableName() string {
	return "m_mitra_branch"
}

func (p *MitraBranch) Prepare() {
	p.MitraID = p.MitraID
	p.MitraBranchParentID = p.MitraBranchParentID
	p.MitraBranchCode = p.MitraBranchCode
	p.MitraBranchName = p.MitraBranchName
	p.MitraBranchType = p.MitraBranchType
	p.MitraBranchPhone = p.MitraBranchPhone
	p.MitraBranchFax = p.MitraBranchFax
	p.MitraBranchAddress = p.MitraBranchAddress
	p.AddressID = p.AddressID
	p.MitraBranchStatus = p.MitraBranchStatus
	p.MitraBranchCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraBranchCreatedBy))
	p.MitraBranchCreatedAt = time.Now()
	p.MitraBranchUpdatedBy = html.EscapeString(strings.TrimSpace(p.MitraBranchUpdatedBy))
	p.MitraBranchUpdatedAt = p.MitraBranchUpdatedAt
	p.MitraBranchDeletedBy = html.EscapeString(strings.TrimSpace(p.MitraBranchDeletedBy))
	p.MitraBranchDeletedAt = p.MitraBranchDeletedAt
}

func (p *MitraBranch) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.MitraBranchCode == "" {
			return errors.New("Required Country")
		}
		if p.MitraBranchName == "" {
			return errors.New("Required Province")
		}
		if p.MitraBranchType == "" {
			return errors.New("Required Regency")
		}
		if p.MitraBranchPhone == "" {
			return errors.New("Required District")
		}
		if p.MitraBranchFax == "" {
			return errors.New("Required Village")
		}
		return nil
	}
}

func (p *MitraBranch) SaveMitraBranch(db *gorm.DB) (*MitraBranch, error) {
	var err error
	err = db.Debug().Model(&MitraBranch{}).Create(&p).Error
	if err != nil {
		return &MitraBranch{}, err
	}
	return p, nil
}

func (p *MitraBranch) FindAllMitraBranch(db *gorm.DB) (*[]MitraBranchData, error) {
	var err error
	mitrabranch := []MitraBranchData{}
	err = db.Debug().Model(&MitraBranchData{}).Limit(100).
		Select(`m_mitra_branch.mitra_branch_id,
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
		Joins("JOIN m_mitra on m_mitra_branch.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_address on m_mitra_branch.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Find(&mitrabranch).Error
	if err != nil {
		return &[]MitraBranchData{}, err
	}
	return &mitrabranch, nil
}

func (p *MitraBranch) FindAllMitraBranchStatus(db *gorm.DB, status uint64) (*[]MitraBranchData, error) {
	var err error
	mitrabranch := []MitraBranchData{}
	err = db.Debug().Model(&MitraBranchData{}).Limit(100).
		Select(`m_mitra_branch.mitra_branch_id,
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
		Joins("JOIN m_mitra on m_mitra_branch.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_address on m_mitra_branch.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("mitra_branch_status = ?", status).
		Find(&mitrabranch).Error
	if err != nil {
		return &[]MitraBranchData{}, err
	}
	return &mitrabranch, nil
}

func (p *MitraBranch) FindMitraBranchDataByID(db *gorm.DB, pid uint64) (*MitraBranch, error) {
	var err error
	err = db.Debug().Model(&MitraBranch{}).Where("mitra_branch_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraBranch{}, err
	}
	return p, nil
}

func (p *MitraBranch) FindMitraBranchByID(db *gorm.DB, pid uint64) (*MitraBranchData, error) {
	var err error
	mitrabranch := MitraBranchData{}
	err = db.Debug().Model(&MitraBranchData{}).Limit(100).
		Select(`m_mitra_branch.mitra_branch_id,
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
		Joins("JOIN m_mitra on m_mitra_branch.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_address on m_mitra_branch.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("mitra_branch_id = ?", pid).
		Find(&mitrabranch).Error
	if err != nil {
		return &MitraBranchData{}, err
	}
	return &mitrabranch, nil
}

func (p *MitraBranch) FindMitraBranchStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraBranchData, error) {
	var err error
	mitrabranch := MitraBranchData{}
	err = db.Debug().Model(&MitraBranchData{}).Limit(100).
		Select(`m_mitra_branch.mitra_branch_id,
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
		Joins("JOIN m_mitra on m_mitra_branch.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_address on m_mitra_branch.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("mitra_branch_id = ? AND mitra_branch_status = ?", pid, status).
		Find(&mitrabranch).Error
	if err != nil {
		return &MitraBranchData{}, err
	}
	return &mitrabranch, nil
}

func (p *MitraBranch) UpdateMitraBranch(db *gorm.DB, pid uint64) (*MitraBranch, error) {
	var err error
	db = db.Debug().Model(&MitraBranch{}).Where("mitra_branch_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_branch_parent_id":  p.MitraBranchParentID,
			"mitra_branch_code":       p.MitraBranchCode,
			"mitra_branch_name":       p.MitraBranchName,
			"mitra_branch_type":       p.MitraBranchType,
			"mitra_branch_phone":      p.MitraBranchPhone,
			"mitra_branch_fax":        p.MitraBranchFax,
			"mitra_branch_address":    p.MitraBranchAddress,
			"mitra_branch_status":     p.MitraBranchStatus,
			"mitra_branch_updated_by": p.MitraBranchUpdatedBy,
			"mitra_branch_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraBranch{}).Where("mitra_branch_id = ?", pid).Error
	if err != nil {
		return &MitraBranch{}, err
	}
	return p, nil
}

func (p *MitraBranch) UpdateMitraBranchStatus(db *gorm.DB, pid uint64) (*MitraBranch, error) {
	var err error
	db = db.Debug().Model(&MitraBranch{}).Where("mitra_branch_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_branch_status":     p.MitraBranchStatus,
			"mitra_branch_updated_by": p.MitraBranchUpdatedBy,
			"mitra_branch_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraBranch{}).Where("mitra_branch_id = ?", pid).Error
	if err != nil {
		return &MitraBranch{}, err
	}
	return p, nil
}

func (p *MitraBranch) UpdateMitraBranchStatusOnly(db *gorm.DB, pid uint64) (*MitraBranch, error) {
	var err error
	db = db.Debug().Model(&MitraBranch{}).Where("mitra_branch_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_branch_status": p.MitraBranchStatus,
		},
	)
	err = db.Debug().Model(&MitraBranch{}).Where("mitra_branch_id = ?", pid).Error
	if err != nil {
		return &MitraBranch{}, err
	}
	return p, nil
}

func (p *MitraBranch) UpdateMitraBranchStatusDelete(db *gorm.DB, pid uint64) (*MitraBranch, error) {
	var err error
	db = db.Debug().Model(&MitraBranch{}).Where("mitra_branch_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_branch_status":     3,
			"mitra_branch_deleted_by": p.MitraBranchDeletedBy,
			"mitra_branch_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraBranch{}).Where("mitra_branch_id = ?", pid).Error
	if err != nil {
		return &MitraBranch{}, err
	}
	return p, nil
}

func (p *MitraBranch) DeleteMitraBranch(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraBranch{}).Where("mitra_branch_id = ?", pid).Take(&MitraBranch{}).Delete(&MitraBranch{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraBranch not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//BY MITRA ID
//====================================================================================================================================

func (p *MitraBranch) FindAllMitraBranchByMitraID(db *gorm.DB, mitra uint64) (*[]MitraBranchData, error) {
	var err error
	mitrabranch := []MitraBranchData{}
	err = db.Debug().Model(&MitraBranchData{}).Limit(100).
		Select(`m_mitra_branch.mitra_branch_id,
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
		Joins("JOIN m_mitra on m_mitra_branch.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_address on m_mitra_branch.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("m_mitra_branch.mitra_id = ?", mitra).
		Find(&mitrabranch).Error
	if err != nil {
		return &[]MitraBranchData{}, err
	}
	return &mitrabranch, nil
}

func (p *MitraBranch) FindMitraBranchByMitraID(db *gorm.DB, mitra uint64, pid uint64) (*MitraBranchData, error) {
	var err error
	mitrabranch := MitraBranchData{}
	err = db.Debug().Model(&MitraBranchData{}).Limit(100).
		Select(`m_mitra_branch.mitra_branch_id,
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
		Joins("JOIN m_mitra on m_mitra_branch.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_address on m_mitra_branch.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("mitra_branch_id = ? AND m_mitra_branch.mitra_id = ?", pid, mitra).
		Find(&mitrabranch).Error
	if err != nil {
		return &MitraBranchData{}, err
	}
	return &mitrabranch, nil
}

func (p *MitraBranch) FindAllMitraBranchStatusByMitraID(db *gorm.DB, mitra uint64, status uint64) (*[]MitraBranchData, error) {
	var err error
	mitrabranch := []MitraBranchData{}
	err = db.Debug().Model(&MitraBranchData{}).Limit(100).
		Select(`m_mitra_branch.mitra_branch_id,
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
		Joins("JOIN m_mitra on m_mitra_branch.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_address on m_mitra_branch.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("mitra_branch_status = ? AND m_mitra_branch.mitra_id = ?", status, mitra).
		Find(&mitrabranch).Error
	if err != nil {
		return &[]MitraBranchData{}, err
	}
	return &mitrabranch, nil
}

func (p *MitraBranch) FindMitraBranchStatusByMitraID(db *gorm.DB, mitra uint64, pid uint64, status uint64) (*MitraBranchData, error) {
	var err error
	mitrabranch := MitraBranchData{}
	err = db.Debug().Model(&MitraBranchData{}).Limit(100).
		Select(`m_mitra_branch.mitra_branch_id,
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
		Joins("JOIN m_mitra on m_mitra_branch.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_address on m_mitra_branch.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("mitra_branch_id = ? AND mitra_branch_status = ? AND m_mitra_branch.mitra_id = ?", pid, status, mitra).
		Find(&mitrabranch).Error
	if err != nil {
		return &MitraBranchData{}, err
	}
	return &mitrabranch, nil
}

//ADDITIONAL
//====================================================================================================================================

func (p *MitraBranch) FindAllMitraBranchByMitra(db *gorm.DB, mitra uint64) (*MitraBranchData, error) {
	var err error
	mitrabranch := MitraBranchData{}
	err = db.Debug().Model(&MitraBranchData{}).Limit(100).
		Select(`m_mitra_branch.mitra_branch_id,
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
		Joins("JOIN m_mitra on m_mitra_branch.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_address on m_mitra_branch.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("m_mitra_branch.mitra_id = ?", mitra).
		Find(&mitrabranch).Error
	if err != nil {
		return &MitraBranchData{}, err
	}
	return &mitrabranch, nil
}

func (p *MitraBranch) FindAllMitraBranchStatusByMitra(db *gorm.DB, mitra uint64, status uint64) (*MitraBranchData, error) {
	var err error
	mitrabranch := MitraBranchData{}
	err = db.Debug().Model(&MitraBranchData{}).Limit(100).
		Select(`m_mitra_branch.mitra_branch_id,
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
		Joins("JOIN m_mitra on m_mitra_branch.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_address on m_mitra_branch.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("mitra_branch_status = ? AND m_mitra_branch.mitra_id = ?", status, mitra).
		Find(&mitrabranch).Error
	if err != nil {
		return &MitraBranchData{}, err
	}
	return &mitrabranch, nil
}

type ResponseMitraBranchHierarchy struct {
	Status  int                 `json:"status"`
	Message string              `json:"message"`
	Data    *MitraBranchPSTData `json:"data"`
}

type MitraBranchPSTData struct {
	MitraBranchID       uint64               `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	MitraBranchName     string               `gorm:"column:mitra_branch_name" json:"mitra_branch_name"`
	MitraBranchParentID uint64               `gorm:"column:mitra_branch_parent_id" json:"mitra_branch_parent_id"`
	MitraBranchCBG      []MitraBranchCBGData `json:"mitra_branch_cabang"`
}

type MitraBranchCBGData struct {
	MitraBranchID       uint64               `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	MitraBranchName     string               `gorm:"column:mitra_branch_name" json:"mitra_branch_name"`
	MitraBranchParentID uint64               `gorm:"column:mitra_branch_parent_id" json:"mitra_branch_parent_id"`
	MitraBranchKAS      []MitraBranchKASData `json:"mitra_branch_kas"`
}

type MitraBranchKASData struct {
	MitraBranchID       uint64 `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	MitraBranchName     string `gorm:"column:mitra_branch_name" json:"mitra_branch_name"`
	MitraBranchParentID uint64 `gorm:"column:mitra_branch_parent_id" json:"mitra_branch_parent_id"`
}

func (MitraBranchPSTData) TableName() string {
	return "m_mitra_branch"
}

func (MitraBranchCBGData) TableName() string {
	return "m_mitra_branch"
}

func (MitraBranchKASData) TableName() string {
	return "m_mitra_branch"
}

func (p *MitraBranch) FindAllMitraBranchPSTByMitraID(db *gorm.DB, mitra uint64) (*MitraBranchPSTData, error) {
	var err error
	mitrabranch := MitraBranchPSTData{}
	err = db.Debug().Model(&MitraBranchPSTData{}).Limit(100).
		Select(`m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_parent_id,
				m_mitra_branch.mitra_branch_name`).
		Where("m_mitra_branch.mitra_id = ? AND m_mitra_branch.mitra_branch_status = ? AND m_mitra_branch.mitra_branch_type = ?", mitra, 1, "PST").
		Find(&mitrabranch).Error
	if err != nil {
		return &MitraBranchPSTData{}, err
	}
	return &mitrabranch, nil
}

func (p *MitraBranch) FindAllMitraBranchCBGByMitraID(db *gorm.DB, mitrabranchparent uint64) ([]MitraBranchCBGData, error) {
	var err error
	mitrabranch := []MitraBranchCBGData{}
	err = db.Debug().Model(MitraBranchCBGData{}).Limit(100).
		Select(`m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_parent_id,
				m_mitra_branch.mitra_branch_name`).
		Where("m_mitra_branch.mitra_branch_parent_id = ? AND m_mitra_branch.mitra_branch_status = ? AND m_mitra_branch.mitra_branch_type = ?", mitrabranchparent, 1, "CBG").
		Find(&mitrabranch).Error
	if err != nil {
		return []MitraBranchCBGData{}, err
	}
	return mitrabranch, nil
}

func (p *MitraBranch) FindAllMitraBranchKASByMitraID(db *gorm.DB, mitrabranchparent uint64) ([]MitraBranchKASData, error) {
	var err error
	mitrabranch := []MitraBranchKASData{}
	err = db.Debug().Model(&MitraBranchKASData{}).Limit(100).
		Select(`m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_parent_id,
				m_mitra_branch.mitra_branch_name`).
		Where("m_mitra_branch.mitra_branch_parent_id = ? AND m_mitra_branch.mitra_branch_status = ? AND m_mitra_branch.mitra_branch_type = ?", mitrabranchparent, 1, "KAS").
		Find(&mitrabranch).Error
	if err != nil {
		return []MitraBranchKASData{}, err
	}
	return mitrabranch, nil
}

type ResponseMitraBranchHierarchyCBG struct {
	Status  int                 `json:"status"`
	Message string              `json:"message"`
	Data    *MitraBranchCBGData `json:"data"`
}

func (p *MitraBranch) FindAllMitraBranchCBGByID(db *gorm.DB, mitrabranchparent uint64) (*MitraBranchCBGData, error) {
	var err error
	mitrabranch := MitraBranchCBGData{}
	err = db.Debug().Model(MitraBranchCBGData{}).Limit(100).
		Select(`m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_parent_id,
				m_mitra_branch.mitra_branch_name`).
		Where("m_mitra_branch.mitra_branch_id = ? AND m_mitra_branch.mitra_branch_status = ? AND m_mitra_branch.mitra_branch_type = ?", mitrabranchparent, 1, "CBG").
		Find(&mitrabranch).Error
	if err != nil {
		return &MitraBranchCBGData{}, err
	}
	return &mitrabranch, nil
}
