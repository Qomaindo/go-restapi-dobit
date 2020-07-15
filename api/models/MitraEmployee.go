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

type MitraEmployee struct {
	MitraEmployeeID         *uint64    `gorm:"column:mitra_emp_id;primary_key;" json:"mitra_emp_id"`
	MitraID                 uint64     `gorm:"column:mitra_id" json:"mitra_id"`
	MitraBranchID           uint64     `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	MitraEmployeeCode       string     `gorm:"column:mitra_emp_code" json:"mitra_emp_code"`
	MitraEmployeeName       string     `gorm:"column:mitra_emp_name" json:"mitra_emp_name"`
	MitraEmployeeGender     string     `gorm:"column:mitra_emp_gender" json:"mitra_emp_gender"`
	MitraEmployeeBirthPlace string     `gorm:"column:mitra_emp_birth_place" json:"mitra_emp_birth_place"`
	MitraEmployeeBirthDate  string     `gorm:"column:mitra_emp_birth_date" json:"mitra_emp_birth_date"`
	MitraEmployeeAddress    string     `gorm:"column:mitra_emp_address" json:"mitra_emp_address"`
	AddressID               uint64     `gorm:"column:address_id" json:"address_id"`
	MitraEmployeePhoto      string     `gorm:"column:mitra_emp_photo" json:"mitra_emp_photo"`
	MitraEmployeeStatus     uint64     `gorm:"column:mitra_emp_status;size:2" json:"mitra_emp_status"`
	MitraEmployeeCreatedBy  string     `gorm:"column:mitra_emp_created_by;size:125" json:"mitra_emp_created_by"`
	MitraEmployeeCreatedAt  time.Time  `gorm:"column:mitra_emp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_emp_created_at"`
	MitraEmployeeUpdatedBy  string     `gorm:"column:mitra_emp_updated_by;size:125" json:"mitra_emp_updated_by"`
	MitraEmployeeUpdatedAt  *time.Time `gorm:"column:mitra_emp_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_emp_created_at"`
	MitraEmployeeDeletedBy  string     `gorm:"column:mitra_emp_deleted_by;size:125" json:"mitra_emp_deleted_by"`
	MitraEmployeeDeletedAt  *time.Time `gorm:"column:mitra_emp_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_emp_deleted_at"`
}

type MitraEmployeeData struct {
	MitraEmployeeID         uint64     `gorm:"column:mitra_emp_id" json:"mitra_emp_id"`
	MitraID                 uint64     `gorm:"column:mitra_id" json:"mitra_id"`
	MitraName               string     `gorm:"column:mitra_name" json:"mitra_name"`
	MitraBranchID           uint64     `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	MitraBranchName         string     `gorm:"column:mitra_branch_name" json:"mitra_branch_name"`
	MitraEmployeeCode       string     `gorm:"column:mitra_emp_code" json:"mitra_emp_code"`
	MitraEmployeeName       string     `gorm:"column:mitra_emp_name" json:"mitra_emp_name"`
	MitraEmployeeGender     string     `gorm:"column:mitra_emp_gender" json:"mitra_emp_gender"`
	MitraEmployeeBirthPlace string     `gorm:"column:mitra_emp_birth_place" json:"mitra_emp_birth_place"`
	MitraEmployeeBirthDate  string     `gorm:"column:mitra_emp_birth_date" json:"mitra_emp_birth_date"`
	MitraEmployeeAddress    string     `gorm:"column:mitra_emp_address" json:"mitra_emp_address"`
	AddressID               uint64     `gorm:"column:address_id" json:"address_id"`
	CountryID               uint64     `gorm:"column:country_id" json:"country_id"`
	CountryName             string     `gorm:"column:country_name" json:"country_name"`
	ProvinceID              uint64     `gorm:"column:province_id" json:"province_id"`
	ProvinceName            string     `gorm:"column:province_name" json:"province_name"`
	RegencyID               uint64     `gorm:"column:regency_id" json:"regency_id"`
	RegencyName             string     `gorm:"column:regency_name" json:"regency_name"`
	DistrictID              uint64     `gorm:"column:district_id" json:"district_id"`
	DistrictName            string     `gorm:"column:district_name" json:"district_name"`
	VillageID               uint64     `gorm:"column:village_id" json:"village_id"`
	VillageName             string     `gorm:"column:village_name" json:"village_name"`
	MitraEmployeePhoto      string     `gorm:"column:mitra_emp_photo" json:"mitra_emp_photo"`
	MitraEmployeeStatus     uint64     `gorm:"column:mitra_emp_status" json:"mitra_emp_status"`
	MitraEmployeeCreatedBy  string     `gorm:"column:mitra_emp_created_by" json:"mitra_emp_created_by"`
	MitraEmployeeCreatedAt  time.Time  `gorm:"column:mitra_emp_created_at" json:"mitra_emp_created_at"`
	MitraEmployeeUpdatedBy  string     `gorm:"column:mitra_emp_updated_by" json:"mitra_emp_updated_by"`
	MitraEmployeeUpdatedAt  *time.Time `gorm:"column:mitra_emp_updated_at" json:"mitra_emp_updated_at"`
	MitraEmployeeDeletedBy  string     `gorm:"column:mitra_emp_deleted_by" json:"mitra_emp_deleted_by"`
	MitraEmployeeDeletedAt  *time.Time `gorm:"column:mitra_emp_deleted_at" json:"mitra_emp_deleted_at"`
}

type ResponseMitraEmployees struct {
	Status  int                  `json:"status"`
	Message string               `json:"message"`
	Data    *[]MitraEmployeeData `json:"data"`
}

type ResponseMitraEmployee struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    *MitraEmployeeData `json:"data"`
}

type ResponseMitraEmployeeCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraEmployee) TableName() string {
	return "m_mitra_employee"
}

func (MitraEmployeeData) TableName() string {
	return "m_mitra_employee"
}

func (p *MitraEmployee) Prepare() {
	p.MitraEmployeeID = nil
	p.MitraID = p.MitraID
	p.MitraBranchID = p.MitraBranchID
	p.MitraEmployeeCode = p.MitraEmployeeCode
	p.MitraEmployeeName = p.MitraEmployeeName
	p.MitraEmployeeGender = p.MitraEmployeeGender
	p.MitraEmployeeBirthPlace = p.MitraEmployeeBirthPlace
	p.MitraEmployeeBirthDate = p.MitraEmployeeBirthDate
	p.MitraEmployeeAddress = p.MitraEmployeeAddress
	p.AddressID = p.AddressID
	p.MitraEmployeePhoto = p.MitraEmployeePhoto
	p.MitraEmployeeStatus = p.MitraEmployeeStatus
	p.MitraEmployeeCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraEmployeeCreatedBy))
	p.MitraEmployeeCreatedAt = time.Now()
	p.MitraEmployeeUpdatedBy = html.EscapeString(strings.TrimSpace(p.MitraEmployeeUpdatedBy))
	p.MitraEmployeeUpdatedAt = p.MitraEmployeeUpdatedAt
	p.MitraEmployeeDeletedBy = html.EscapeString(strings.TrimSpace(p.MitraEmployeeDeletedBy))
	p.MitraEmployeeDeletedAt = p.MitraEmployeeDeletedAt
}

func (p *MitraEmployee) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.MitraEmployeeCode == "" {
			return errors.New("Required Country")
		}
		if p.MitraEmployeeName == "" {
			return errors.New("Required Province")
		}
		if p.MitraEmployeeGender == "" {
			return errors.New("Required Regency")
		}
		if p.MitraEmployeeBirthPlace == "" {
			return errors.New("Required District")
		}
		if p.MitraEmployeeBirthDate == "" {
			return errors.New("Required Village")
		}
		if p.MitraEmployeeAddress == "" {
			return errors.New("Required Village")
		}
		if p.AddressID == 0 {
			return errors.New("Required Village")
		}
		if p.MitraEmployeePhoto == "" {
			return errors.New("Required Village")
		}
		return nil
	}
}

func (p *MitraEmployee) SaveMitraEmployee(db *gorm.DB) (*MitraEmployee, error) {
	var err error
	err = db.Debug().Model(&MitraEmployee{}).Create(&p).Error
	if err != nil {
		return &MitraEmployee{}, err
	}
	return p, nil
}

func (p *MitraEmployee) FindAllMitraEmployee(db *gorm.DB) (*[]MitraEmployeeData, error) {
	var err error
	address := []MitraEmployeeData{}
	err = db.Debug().Model(&MitraEmployeeData{}).Limit(100).
		Select(`m_mitra_employee.mitra_emp_id,
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
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_address on m_mitra_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Find(&address).Error
	if err != nil {
		return &[]MitraEmployeeData{}, err
	}
	return &address, nil
}

func (p *MitraEmployee) FindAllMitraEmployeeStatus(db *gorm.DB, status uint64) (*[]MitraEmployeeData, error) {
	var err error
	address := []MitraEmployeeData{}
	err = db.Debug().Model(&MitraEmployeeData{}).Limit(100).
		Select(`m_mitra_employee.mitra_emp_id,
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
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_address on m_mitra_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("mitra_emp_status = ?", status).
		Find(&address).Error
	if err != nil {
		return &[]MitraEmployeeData{}, err
	}
	return &address, nil
}

func (p *MitraEmployee) FindMitraEmployeeDataByID(db *gorm.DB, pid uint64) (*MitraEmployee, error) {
	var err error
	err = db.Debug().Model(&MitraEmployee{}).Where("mitra_emp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraEmployee{}, err
	}
	return p, nil
}

func (p *MitraEmployee) FindMitraEmployeeByID(db *gorm.DB, pid uint64) (*MitraEmployeeData, error) {
	var err error
	address := MitraEmployeeData{}
	err = db.Debug().Model(&MitraEmployeeData{}).Limit(100).
		Select(`m_mitra_employee.mitra_emp_id,
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
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_address on m_mitra_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("mitra_emp_id = ?", pid).
		Find(&address).Error
	if err != nil {
		return &MitraEmployeeData{}, err
	}
	return &address, nil
}

func (p *MitraEmployee) FindMitraEmployeeStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraEmployeeData, error) {
	var err error
	address := MitraEmployeeData{}
	err = db.Debug().Model(&MitraEmployeeData{}).Limit(100).
		Select(`m_mitra_employee.mitra_emp_id,
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
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_address on m_mitra_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("mitra_emp_id = ? AND mitra_emp_status = ?", pid, status).
		Find(&address).Error
	if err != nil {
		return &MitraEmployeeData{}, err
	}
	return &address, nil
}

func (p *MitraEmployee) UpdateMitraEmployee(db *gorm.DB, pid uint64) (*MitraEmployee, error) {
	var err error
	db = db.Debug().Model(&MitraEmployee{}).Where("mitra_emp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_id":              p.MitraID,
			"mitra_branch_id":       p.MitraBranchID,
			"mitra_emp_code":        p.MitraEmployeeCode,
			"mitra_emp_name":        p.MitraEmployeeName,
			"mitra_emp_gender":      p.MitraEmployeeGender,
			"mitra_emp_birth_place": p.MitraEmployeeBirthPlace,
			"mitra_emp_birth_date":  p.MitraEmployeeBirthDate,
			"mitra_emp_address":     p.MitraEmployeeAddress,
			"address_id":            p.AddressID,
			"mitra_emp_photo":       p.MitraEmployeePhoto,
			"mitra_emp_status":      p.MitraEmployeeStatus,
			"mitra_emp_updated_by":  p.MitraEmployeeUpdatedBy,
			"mitra_emp_updated_at":  time.Now(),
		},
	)
	err = db.Debug().Model(&MitraEmployee{}).Where("mitra_emp_id = ?", pid).Error
	if err != nil {
		return &MitraEmployee{}, err
	}
	return p, nil
}

func (p *MitraEmployee) UpdateMitraEmployeeStatus(db *gorm.DB, pid uint64) (*MitraEmployee, error) {
	var err error
	db = db.Debug().Model(&MitraEmployee{}).Where("mitra_emp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_emp_status":     p.MitraEmployeeStatus,
			"mitra_emp_updated_by": p.MitraEmployeeUpdatedBy,
			"mitra_emp_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraEmployee{}).Where("mitra_emp_id = ?", pid).Error
	if err != nil {
		return &MitraEmployee{}, err
	}
	return p, nil
}

func (p *MitraEmployee) UpdateMitraEmployeeStatusOnly(db *gorm.DB, pid uint64) (*MitraEmployee, error) {
	var err error
	db = db.Debug().Model(&MitraEmployee{}).Where("mitra_emp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_emp_status": p.MitraEmployeeStatus,
		},
	)
	err = db.Debug().Model(&MitraEmployee{}).Where("mitra_emp_id = ?", pid).Error
	if err != nil {
		return &MitraEmployee{}, err
	}
	return p, nil
}

func (p *MitraEmployee) UpdateMitraEmployeeStatusDelete(db *gorm.DB, pid uint64) (*MitraEmployee, error) {
	var err error
	db = db.Debug().Model(&MitraEmployee{}).Where("mitra_emp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_emp_status":     3,
			"mitra_emp_deleted_by": p.MitraEmployeeDeletedBy,
			"mitra_emp_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraEmployee{}).Where("mitra_emp_id = ?", pid).Error
	if err != nil {
		return &MitraEmployee{}, err
	}
	return p, nil
}

func (p *MitraEmployee) DeleteMitraEmployee(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraEmployee{}).Where("mitra_emp_id = ?", pid).Take(&MitraEmployee{}).Delete(&MitraEmployee{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraEmployee not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//GET BY MITRA ID
//====================================================================================================================================

func (p *MitraEmployee) FindAllMitraEmployeeMitra(db *gorm.DB, mitra uint64) (*[]MitraEmployeeData, error) {
	var err error
	address := []MitraEmployeeData{}
	err = db.Debug().Model(&MitraEmployeeData{}).Limit(100).
		Select(`m_mitra_employee.mitra_emp_id,
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
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_address on m_mitra_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("m_mitra.mitra_id = ?", mitra).
		Find(&address).Error
	if err != nil {
		return &[]MitraEmployeeData{}, err
	}
	return &address, nil
}

func (p *MitraEmployee) FindAllMitraEmployeeStatusMitra(db *gorm.DB, status uint64, mitra uint64) (*[]MitraEmployeeData, error) {
	var err error
	address := []MitraEmployeeData{}
	err = db.Debug().Model(&MitraEmployeeData{}).Limit(100).
		Select(`m_mitra_employee.mitra_emp_id,
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
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_address on m_mitra_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("mitra_emp_status = ? AND m_mitra.mitra_id = ?", status, mitra).
		Find(&address).Error
	if err != nil {
		return &[]MitraEmployeeData{}, err
	}
	return &address, nil
}

func (p *MitraEmployee) FindMitraEmployeeDataByIDMitra(db *gorm.DB, pid uint64, mitra uint64) (*MitraEmployee, error) {
	var err error
	err = db.Debug().Model(&MitraEmployee{}).Where("mitra_emp_id = ? AND m_mitra.mitra_id = ?", pid, mitra).Take(&p).Error
	if err != nil {
		return &MitraEmployee{}, err
	}
	return p, nil
}

func (p *MitraEmployee) FindMitraEmployeeByIDMitra(db *gorm.DB, pid uint64, mitra uint64) (*MitraEmployeeData, error) {
	var err error
	address := MitraEmployeeData{}
	err = db.Debug().Model(&MitraEmployeeData{}).Limit(100).
		Select(`m_mitra_employee.mitra_emp_id,
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
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_address on m_mitra_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("mitra_emp_id = ? AND m_mitra.mitra_id = ?", pid, mitra).
		Find(&address).Error
	if err != nil {
		return &MitraEmployeeData{}, err
	}
	return &address, nil
}

func (p *MitraEmployee) FindMitraEmployeeStatusByIDMitra(db *gorm.DB, pid uint64, status uint64, mitra uint64) (*MitraEmployeeData, error) {
	var err error
	address := MitraEmployeeData{}
	err = db.Debug().Model(&MitraEmployeeData{}).Limit(100).
		Select(`m_mitra_employee.mitra_emp_id,
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
		Joins("JOIN m_mitra on m_mitra_employee.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on m_mitra_employee.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_address on m_mitra_employee.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("mitra_emp_id = ? AND mitra_emp_status = ? AND m_mitra.mitra_id = ?", pid, status, mitra).
		Find(&address).Error
	if err != nil {
		return &MitraEmployeeData{}, err
	}
	return &address, nil
}

//LOGIN
//====================================================================================================================================

type MitraLogin struct {
	MitraType     string `json:"mitra_type"`
	MitraCode     string `json:"mitra_code"`
	MitraUsername string `json:"mitra_username"`
	MitraPassword string `json:"mitra_password"`
}

type MitraPassword struct {
	MitraAccountOldPassword string `json:"mitra_account_old_password"`
	MitraAccountNewPassword string `json:"mitra_account_new_password"`
}

type ResponseMitraLoginAuthenticationData struct {
	ID                      uint64                               `json:"id"`
	Type                    string                               `json:"type"`
	Token                   string                               `json:"token"`
	Status                  uint64                               `json:"status"`
	Name                    string                               `json:"name"`
	Username                string                               `json:"username"`
	Photo                   string                               `json:"photo"`
	MitraID                 uint64                               `json:"mitra_id"`
	MitraName               string                               `json:"mitra_name"`
	IPAddress               string                               `json:"ip_address"`
	LastLogin               time.Time                            `json:"last_login"`
	IsNew                   bool                                 `json:"is_new"`
	IsRequestChangePassword bool                                 `json:"is_req_chpw"`
	IsLocked                bool                                 `json:"is_locked"`
	IsAO                    bool                                 `json:"is_ao"`
	IsCredit                bool                                 `json:"is_credit"`
	GroupName               string                               `json:"group_name"`
	GroupRole               string                               `json:"group_role"`
	GroupPrivilege          []MitraGroupAccessPrivilegeDataLogin `json:"group_privilege"`
}

type ResponseMitraLoginAuthentication struct {
	Status  int                                  `json:"status"`
	Message string                               `json:"message"`
	Data    ResponseMitraLoginAuthenticationData `json:"data"`
}

type ResponseMitraLoginGetIDData struct {
	MitraID       uint64 `json:"mitra_id"`
	MitraBranchID uint64 `json:"mitra_branch_id"`
}

type ResponseMitraLoginGetID struct {
	Status  int                         `json:"status"`
	Message string                      `json:"message"`
	Data    ResponseMitraLoginGetIDData `json:"data"`
}

func (p *MitraLogin) Validate(action string) error {
	switch strings.ToLower(action) {

	case "login":
		if p.MitraCode == "" {
			return errors.New("Required Mitra Code")
		}
		if p.MitraUsername == "" {
			return errors.New("Username Wajib Diisi")
		}
		if p.MitraPassword == "" {
			return errors.New("Password Wajib Diisi")
		}
		return nil

	default:
		if p.MitraUsername == "" {
			return errors.New("Username Wajib Diisi")
		}
		return nil
	}

}

//ADDITIONAL
//====================================================================================================================================

type MitraProfileWeb struct {
	MitraProfileEmail     string `gorm:"column:mitra_profile_email" json:"mitra_profile_email"`
	MitraProfileUsername  string `gorm:"column:mitra_profile_username" json:"mitra_profile_username"`
	MitraProfilePhoneCode string `gorm:"column:mitra_profile_phone_code" json:"mitra_profile_phone_code"`
	MitraProfilePhone     string `gorm:"column:mitra_profile_phone" json:"mitra_profile_phone"`
}

func (p *MitraProfileWeb) Prepare() {
	p.MitraProfileEmail = p.MitraProfileEmail
	p.MitraProfileUsername = p.MitraProfileUsername
	p.MitraProfilePhoneCode = p.MitraProfilePhoneCode
	p.MitraProfilePhone = p.MitraProfilePhone
}
