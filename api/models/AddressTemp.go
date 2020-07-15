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

type AddressTemp struct {
	AddressTempID           uint64    `gorm:"column:address_temp_id;primary_key;" json:"address_temp_id"`
	AddressID               uint64    `gorm:"column:address_id" json:"address_id"`
	CountryTempID           uint64    `gorm:"column:country_temp_id" json:"country_temp_id"`
	ProvinceTempID          uint64    `gorm:"column:province_temp_id" json:"province_temp_id"`
	RegencyTempID           uint64    `gorm:"column:regency_temp_id" json:"regency_temp_id"`
	DistrictTempID          uint64    `gorm:"column:district_temp_id" json:"district_temp_id"`
	VillageTempID           uint64    `gorm:"column:village_temp_id" json:"village_temp_id"`
	AddressTempReason       string    `gorm:"column:address_temp_reason" json:"address_temp_reason"`
	AddressTempStatus       uint64    `gorm:"column:address_temp_status;size:2" json:"address_temp_status"`
	AddressTempActionBefore uint64    `gorm:"column:address_temp_action_before;size:2" json:"address_temp_action_before"`
	AddressTempCreatedBy    string    `gorm:"column:address_temp_created_by;size:125" json:"address_temp_created_by"`
	AddressTempCreatedAt    time.Time `gorm:"column:address_temp_created_at;default:CURRENT_TIMESTAMP" json:"address_temp_created_at"`
}

type AddressTempPend struct {
	AddressTempID           *uint64   `gorm:"column:address_temp_id;primary_key;" json:"address_temp_id"`
	AddressID               *uint64   `gorm:"column:address_id" json:"address_id"`
	CountryTempID           uint64    `gorm:"column:country_temp_id" json:"country_temp_id"`
	ProvinceTempID          uint64    `gorm:"column:province_temp_id" json:"province_temp_id"`
	RegencyTempID           uint64    `gorm:"column:regency_temp_id" json:"regency_temp_id"`
	DistrictTempID          uint64    `gorm:"column:district_temp_id" json:"district_temp_id"`
	VillageTempID           uint64    `gorm:"column:village_temp_id" json:"village_temp_id"`
	AddressTempReason       string    `gorm:"column:address_temp_reason" json:"address_temp_reason"`
	AddressTempStatus       uint64    `gorm:"column:address_temp_status;size:2" json:"address_temp_status"`
	AddressTempActionBefore uint64    `gorm:"column:address_temp_action_before;size:2" json:"address_temp_action_before"`
	AddressTempCreatedBy    string    `gorm:"column:address_temp_created_by;size:125" json:"address_temp_created_by"`
	AddressTempCreatedAt    time.Time `gorm:"column:address_temp_created_at;default:CURRENT_TIMESTAMP" json:"address_temp_created_at"`
}

type AddressTempData struct {
	AddressTempID           uint64    `gorm:"column:address_temp_id" json:"address_temp_id"`
	CountryTempID           uint64    `gorm:"column:country_temp_id" json:"country_temp_id"`
	CountryTempName         string    `gorm:"column:country_temp_name" json:"country_temp_name"`
	ProvinceTempID          uint64    `gorm:"column:province_temp_id" json:"province_temp_id"`
	ProvinceTempName        string    `gorm:"column:province_temp_name" json:"province_temp_name"`
	RegencyTempID           uint64    `gorm:"column:regency_temp_id" json:"regency_temp_id"`
	RegencyTempName         string    `gorm:"column:regency_temp_name" json:"regency_temp_name"`
	DistrictTempID          uint64    `gorm:"column:district_temp_id" json:"district_temp_id"`
	DistrictTempName        string    `gorm:"column:district_temp_name" json:"district_temp_name"`
	VillageTempID           uint64    `gorm:"column:village_temp_id" json:"village_temp_id"`
	VillageTempName         string    `gorm:"column:village_temp_name;size:255" json:"village_temp_name"`
	AddressTempReason       string    `gorm:"column:address_temp_reason" json:"address_temp_reason"`
	AddressTempStatus       uint64    `gorm:"column:address_temp_status;size:2" json:"address_temp_status"`
	AddressTempActionBefore uint64    `gorm:"column:address_temp_action_before;size:2" json:"address_temp_action_before"`
	AddressTempCreatedBy    string    `gorm:"column:address_temp_created_by;size:125" json:"address_temp_created_by"`
	AddressTempCreatedAt    time.Time `gorm:"column:address_temp_created_at" json:"address_temp_created_at"`
	AddressID               uint64    `gorm:"column:address_id" json:"address_id"`
	CountryID               uint64    `gorm:"column:country_id" json:"country_id"`
	CountryName             string    `gorm:"column:country_name" json:"country_name"`
	ProvinceID              uint64    `gorm:"column:province_id" json:"province_id"`
	ProvinceName            string    `gorm:"column:province_name" json:"province_name"`
	RegencyID               uint64    `gorm:"column:regency_id" json:"regency_id"`
	RegencyName             string    `gorm:"column:regency_name" json:"regency_name"`
	DistrictID              uint64    `gorm:"column:district_id" json:"district_id"`
	DistrictName            string    `gorm:"column:district_name" json:"district_name"`
	VillageID               uint64    `gorm:"column:village_id" json:"village_id"`
	VillageName             string    `gorm:"column:village_name" json:"village_name"`
	AddressStatus           uint64    `gorm:"column:address_status" json:"address_status"`
	AddressCreatedBy        string    `gorm:"column:address_created_by" json:"address_created_by"`
	AddressCreatedAt        time.Time `gorm:"column:address_created_at" json:"address_created_at"`
	AddressUpdatedBy        string    `gorm:"column:address_updated_by" json:"address_updated_by"`
	AddressUpdatedAt        time.Time `gorm:"column:address_updated_at" json:"address_updated_at"`
	AddressDeletedBy        string    `gorm:"column:address_deleted_by" json:"address_deleted_by"`
	AddressDeletedAt        time.Time `gorm:"column:address_deleted_at" json:"address_deleted_at"`
}

type ResponseAddressTemps struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    *[]AddressTempData `json:"data"`
}

type ResponseAddressTemp struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    *AddressTempData `json:"data"`
}

type ResponseAddressTempsPend struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    *[]AddressTempPend `json:"data"`
}

type ResponseAddressTempPend struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    *AddressTempPend `json:"data"`
}

type ResponseAddressTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (AddressTemp) TableName() string {
	return "m_address_temp"
}

func (AddressTempPend) TableName() string {
	return "m_address_temp"
}

func (AddressTempData) TableName() string {
	return "m_address_temp"
}

func (p *AddressTemp) Prepare() {
	p.AddressID = p.AddressID
	p.CountryTempID = p.CountryTempID
	p.ProvinceTempID = p.ProvinceTempID
	p.RegencyTempID = p.RegencyTempID
	p.DistrictTempID = p.DistrictTempID
	p.VillageTempID = p.VillageTempID
	p.AddressTempReason = p.AddressTempReason
	p.AddressTempStatus = p.AddressTempStatus
	p.AddressTempCreatedBy = html.EscapeString(strings.TrimSpace(p.AddressTempCreatedBy))
	p.AddressTempCreatedAt = time.Now()
}

func (p *AddressTempPend) Prepare() {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	p.AddressTempID = nil
	p.AddressID = nil
	p.CountryTempID = p.CountryTempID
	p.ProvinceTempID = p.ProvinceTempID
	p.RegencyTempID = p.RegencyTempID
	p.DistrictTempID = p.DistrictTempID
	p.VillageTempID = p.VillageTempID
	p.AddressTempReason = p.AddressTempReason
	p.AddressTempStatus = p.AddressTempStatus
	p.AddressTempCreatedBy = html.EscapeString(strings.TrimSpace(p.AddressTempCreatedBy))
	p.AddressTempCreatedAt = time.Now().In(loc)
}

func (p *AddressTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	case "insertupdate":
		if p.CountryTempID == 0 {
			return errors.New("Negara Wajib diisi")
		}
		if p.ProvinceTempID == 0 {
			return errors.New("Provinsi Wajib diisi")
		}
		if p.RegencyTempID == 0 {
			return errors.New("Kabupaten / Kota Wajib diisi")
		}
		if p.DistrictTempID == 0 {
			return errors.New("Kecamatan Wajib diisi")
		}
		if p.VillageTempID == 0 {
			return errors.New("Kelurahan Wajib diisi")
		}
		if p.AddressTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.AddressTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	case "reason":
		if p.AddressTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.AddressTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	default:
		if p.AddressTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil
	}
}

func (p *AddressTemp) SaveAddressTemp(db *gorm.DB) (*AddressTemp, error) {
	var err error
	err = db.Debug().Model(&AddressTemp{}).Create(&p).Error
	if err != nil {
		return &AddressTemp{}, err
	}
	return p, nil
}

func (p *AddressTempPend) SaveAddressTempPend(db *gorm.DB) (*AddressTempPend, error) {
	var err error
	err = db.Debug().Model(&AddressTempPend{}).Create(&p).Error
	if err != nil {
		return &AddressTempPend{}, err
	}
	return p, nil
}

func (p *AddressTemp) FindAllAddressTemp(db *gorm.DB) (*[]AddressTempPend, error) {
	var err error
	address := []AddressTempPend{}
	err = db.Debug().Model(&AddressTempPend{}).Limit(100).
		Select(`m_address_temp.address_temp_id,
				m_address_temp.address_id,
				m_address_temp.country_temp_id,
				m_address_temp.province_temp_id,
				m_address_temp.regency_temp_id,
				m_address_temp.district_temp_id,
				m_address_temp.village_temp_id,
				m_address_temp.address_temp_reason,
				m_address_temp.address_temp_status,
				m_address_temp.address_temp_action_before,
				m_address_temp.address_temp_created_by,
				m_address_temp.address_temp_created_at at time zone 'Asia/Jakarta' as address_temp_created_at`).
		Order("address_temp_created_at desc").
		Find(&address).Error
	if err != nil {
		return &[]AddressTempPend{}, err
	}
	return &address, nil
}

func (p *AddressTemp) FindAllAddressTempPendingActive(db *gorm.DB) (*[]AddressTempPend, error) {
	var err error
	address := []AddressTempPend{}
	err = db.Debug().Model(&AddressTempPend{}).Limit(100).
		Select(`m_address_temp.address_temp_id,
				m_address_temp.address_id,
				m_address_temp.country_temp_id,
				m_address_temp.province_temp_id,
				m_address_temp.regency_temp_id,
				m_address_temp.district_temp_id,
				m_address_temp.village_temp_id,
				m_address_temp.address_temp_reason,
				m_address_temp.address_temp_status,
				m_address_temp.address_temp_action_before,
				m_address_temp.address_temp_created_by,
				m_address_temp.address_temp_created_at at time zone 'Asia/Jakarta' as address_temp_created_at`).
		Where("address_temp_status = ?", 11).
		Order("address_temp_created_at desc").
		Find(&address).Error
	if err != nil {
		return &[]AddressTempPend{}, err
	}
	return &address, nil
}

func (p *AddressTemp) FindAllAddressTempStatus(db *gorm.DB, status uint64) (*[]AddressTempData, error) {
	var err error
	address := []AddressTempData{}
	err = db.Debug().Model(&AddressTempData{}).Limit(100).
		Select(`m_address_temp.address_temp_id,
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
				m_address_temp.address_temp_reason,
				m_address_temp.address_temp_status,
				m_address_temp.address_temp_action_before,
				m_address_temp.address_temp_created_by,
				m_address_temp.address_temp_created_at at time zone 'Asia/Jakarta' as address_temp_created_at,
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
				m_address.address_status,
				m_address.address_created_by,
				m_address.address_created_at at time zone 'Asia/Jakarta' as address_created_at,
				m_address.address_updated_by,
				m_address.address_updated_at at time zone 'Asia/Jakarta' as address_updated_at,
				m_address.address_deleted_by,
				m_address.address_deleted_at at time zone 'Asia/Jakarta' as address_deleted_at`).
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_address on m_address_temp.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("address_temp_status = ?", status).
		Order("address_temp_created_at desc").
		Find(&address).Error
	if err != nil {
		return &[]AddressTempData{}, err
	}
	return &address, nil
}

func (p *AddressTemp) FindAddressTempDataByID(db *gorm.DB, pid uint64) (*AddressTemp, error) {
	var err error
	err = db.Debug().Model(&AddressTemp{}).Where("address_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &AddressTemp{}, err
	}
	return p, nil
}

func (p *AddressTemp) FindAddressTempByIDPendingActive(db *gorm.DB, pid uint64) (*AddressTempPend, error) {
	var err error
	address := AddressTempPend{}
	err = db.Debug().Model(&AddressTempPend{}).Limit(100).
		Select(`m_address_temp.address_temp_id,
				m_address_temp.address_id,
				m_address_temp.country_temp_id,
				m_address_temp.province_temp_id,
				m_address_temp.regency_temp_id,
				m_address_temp.district_temp_id,
				m_address_temp.village_temp_id,
				m_address_temp.address_temp_reason,
				m_address_temp.address_temp_status,
				m_address_temp.address_temp_action_before,
				m_address_temp.address_temp_created_by,
				m_address_temp.address_temp_created_at at time zone 'Asia/Jakarta' as address_temp_created_at`).
		Where("address_temp_id = ? AND address_temp_status = ?", pid, 11).
		Order("address_temp_created_at desc").
		Find(&address).Error
	if err != nil {
		return &AddressTempPend{}, err
	}
	return &address, nil
}

func (p *AddressTemp) FindAddressTempByID(db *gorm.DB, pid uint64) (*AddressTempData, error) {
	var err error
	address := AddressTempData{}
	err = db.Debug().Model(&AddressTempData{}).Limit(100).
		Select(`m_address_temp.address_temp_id,
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
				m_address_temp.address_temp_reason,
				m_address_temp.address_temp_status,
				m_address_temp.address_temp_action_before,
				m_address_temp.address_temp_created_by,
				m_address_temp.address_temp_created_at at time zone 'Asia/Jakarta' as address_temp_created_at,
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
				m_address.address_status,
				m_address.address_created_by,
				m_address.address_created_at at time zone 'Asia/Jakarta' as address_created_at,
				m_address.address_updated_by,
				m_address.address_updated_at at time zone 'Asia/Jakarta' as address_updated_at,
				m_address.address_deleted_by,
				m_address.address_deleted_at at time zone 'Asia/Jakarta' as address_deleted_at`).
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_address on m_address_temp.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("address_temp_id = ?", pid).
		Order("address_temp_created_at desc").
		Find(&address).Error
	if err != nil {
		return &AddressTempData{}, err
	}
	return &address, nil
}

func (p *AddressTemp) FindAddressTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*AddressTempData, error) {
	var err error
	address := AddressTempData{}
	err = db.Debug().Model(&AddressTempData{}).Limit(100).
		Select(`m_address_temp.address_temp_id,
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
				m_address_temp.address_temp_reason,
				m_address_temp.address_temp_status,
				m_address_temp.address_temp_action_before,
				m_address_temp.address_temp_created_by,
				m_address_temp.address_temp_created_at at time zone 'Asia/Jakarta' as address_temp_created_at,
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
				m_address.address_status,
				m_address.address_created_by,
				m_address.address_created_at at time zone 'Asia/Jakarta' as address_created_at,
				m_address.address_updated_by,
				m_address.address_updated_at at time zone 'Asia/Jakarta' as address_updated_at,
				m_address.address_deleted_by,
				m_address.address_deleted_at at time zone 'Asia/Jakarta' as address_deleted_at`).
		Joins("JOIN m_country m_country_temp on m_address_temp.country_temp_id=m_country_temp.country_id").
		Joins("JOIN m_province m_province_temp on m_address_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_regency m_regency_temp on m_address_temp.regency_temp_id=m_regency_temp.regency_id").
		Joins("JOIN m_district m_district_temp on m_address_temp.district_temp_id=m_district_temp.district_id").
		Joins("JOIN m_village m_village_temp on m_address_temp.village_temp_id=m_village_temp.village_id").
		Joins("JOIN m_address on m_address_temp.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("address_temp_id = ? AND address_temp_status = ?", pid, status).
		Order("address_temp_created_at desc").
		Find(&address).Error
	if err != nil {
		return &AddressTempData{}, err
	}
	return &address, nil
}

func (p *AddressTemp) UpdateAddressTemp(db *gorm.DB, pid uint64) (*AddressTemp, error) {
	var err error
	db = db.Debug().Model(&AddressTemp{}).Where("address_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"country_temp_id":     p.CountryTempID,
			"province_temp_id":    p.ProvinceTempID,
			"regency_temp_id":     p.RegencyTempID,
			"district_temp_id":    p.DistrictTempID,
			"village_temp_id":     p.VillageTempID,
			"address_temp_reason": p.AddressTempReason,
		},
	)
	err = db.Debug().Model(&AddressTemp{}).Where("address_temp_id = ?", pid).Error
	if err != nil {
		return &AddressTemp{}, err
	}
	return p, nil
}

func (p *AddressTemp) UpdateAddressTempStatus(db *gorm.DB, pid uint64) (*AddressTemp, error) {
	var err error
	db = db.Debug().Model(&AddressTemp{}).Where("address_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"address_temp_reason": p.AddressTempReason,
			"address_temp_status": p.AddressTempStatus,
		},
	)
	err = db.Debug().Model(&AddressTemp{}).Where("address_temp_id = ?", pid).Error
	if err != nil {
		return &AddressTemp{}, err
	}
	return p, nil
}

func (p *AddressTemp) DeleteAddressTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&AddressTemp{}).Where("address_temp_id = ?", pid).Take(&AddressTemp{}).Delete(&AddressTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("AddressTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================
