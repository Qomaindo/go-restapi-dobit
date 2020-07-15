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

type Address struct {
	AddressID        *uint64    `gorm:"column:address_id;primary_key;" json:"address_id"`
	CountryID        uint64     `gorm:"column:country_id" json:"country_id"`
	ProvinceID       uint64     `gorm:"column:province_id" json:"province_id"`
	RegencyID        uint64     `gorm:"column:regency_id" json:"regency_id"`
	DistrictID       uint64     `gorm:"column:district_id" json:"district_id"`
	VillageID        uint64     `gorm:"column:village_id" json:"village_id"`
	AddressStatus    uint64     `gorm:"column:address_status;size:2" json:"address_status"`
	AddressCreatedBy string     `gorm:"column:address_created_by;size:125" json:"address_created_by"`
	AddressCreatedAt time.Time  `gorm:"column:address_created_at;default:CURRENT_TIMESTAMP" json:"address_created_at"`
	AddressUpdatedBy string     `gorm:"column:address_updated_by;size:125" json:"address_updated_by"`
	AddressUpdatedAt *time.Time `gorm:"column:address_updated_at;default:CURRENT_TIMESTAMP" json:"address_created_at"`
	AddressDeletedBy string     `gorm:"column:address_deleted_by;size:125" json:"address_deleted_by"`
	AddressDeletedAt *time.Time `gorm:"column:address_deleted_at;default:CURRENT_TIMESTAMP" json:"address_deleted_at"`
}

type AddressData struct {
	AddressID        uint64     `gorm:"column:address_id" json:"address_id"`
	CountryID        uint64     `gorm:"column:country_id" json:"country_id"`
	CountryName      string     `gorm:"column:country_name" json:"country_name"`
	ProvinceID       uint64     `gorm:"column:province_id" json:"province_id"`
	ProvinceName     string     `gorm:"column:province_name" json:"province_name"`
	RegencyID        uint64     `gorm:"column:regency_id" json:"regency_id"`
	RegencyName      string     `gorm:"column:regency_name" json:"regency_name"`
	DistrictID       uint64     `gorm:"column:district_id" json:"district_id"`
	DistrictName     string     `gorm:"column:district_name" json:"district_name"`
	VillageID        uint64     `gorm:"column:village_id" json:"village_id"`
	VillageName      string     `gorm:"column:village_name" json:"village_name"`
	AddressStatus    uint64     `gorm:"column:address_status" json:"address_status"`
	AddressCreatedBy string     `gorm:"column:address_created_by" json:"address_created_by"`
	AddressCreatedAt time.Time  `gorm:"column:address_created_at" json:"address_created_at"`
	AddressUpdatedBy string     `gorm:"column:address_updated_by" json:"address_updated_by"`
	AddressUpdatedAt *time.Time `gorm:"column:address_updated_at" json:"address_updated_at"`
	AddressDeletedBy string     `gorm:"column:address_deleted_by" json:"address_deleted_by"`
	AddressDeletedAt *time.Time `gorm:"column:address_deleted_at" json:"address_deleted_at"`
}

type ResponseAddresss struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    *[]AddressData `json:"data"`
}

type ResponseAddress struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    *AddressData `json:"data"`
}

type ResponseAddressCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (Address) TableName() string {
	return "m_address"
}

func (AddressData) TableName() string {
	return "m_address"
}

func (p *Address) Prepare() {
	p.AddressID = nil
	p.CountryID = p.CountryID
	p.ProvinceID = p.ProvinceID
	p.RegencyID = p.RegencyID
	p.DistrictID = p.DistrictID
	p.VillageID = p.VillageID
	p.AddressStatus = p.AddressStatus
	p.AddressCreatedBy = html.EscapeString(strings.TrimSpace(p.AddressCreatedBy))
	p.AddressCreatedAt = time.Now()
	p.AddressUpdatedBy = html.EscapeString(strings.TrimSpace(p.AddressUpdatedBy))
	p.AddressUpdatedAt = p.AddressUpdatedAt
	p.AddressDeletedBy = html.EscapeString(strings.TrimSpace(p.AddressDeletedBy))
	p.AddressDeletedAt = p.AddressDeletedAt
}

func (p *Address) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.CountryID == 0 {
			return errors.New("Required Country")
		}
		if p.ProvinceID == 0 {
			return errors.New("Required Province")
		}
		if p.RegencyID == 0 {
			return errors.New("Required Regency")
		}
		if p.DistrictID == 0 {
			return errors.New("Required District")
		}
		if p.VillageID == 0 {
			return errors.New("Required Village")
		}
		return nil
	}
}

func (p *Address) SaveAddress(db *gorm.DB) (*Address, error) {
	var err error
	err = db.Debug().Model(&Address{}).Create(&p).Error
	if err != nil {
		return &Address{}, err
	}
	return p, nil
}

func (p *Address) FindAllAddress(db *gorm.DB) (*[]AddressData, error) {
	var err error
	address := []AddressData{}
	err = db.Debug().Model(&AddressData{}).Limit(100).
		Select(`m_address.address_id,
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
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Find(&address).Error
	if err != nil {
		return &[]AddressData{}, err
	}
	return &address, nil
}

func (p *Address) FindAllAddressStatus(db *gorm.DB, status uint64) (*[]AddressData, error) {
	var err error
	address := []AddressData{}
	err = db.Debug().Model(&AddressData{}).Limit(100).
		Select(`m_address.address_id,
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
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("address_status = ?", status).
		Find(&address).Error
	if err != nil {
		return &[]AddressData{}, err
	}
	return &address, nil
}

func (p *Address) FindAddressDataByID(db *gorm.DB, pid uint64) (*Address, error) {
	var err error
	err = db.Debug().Model(&Address{}).Where("address_id = ?", pid).Take(&p).Error
	if err != nil {
		return &Address{}, err
	}
	return p, nil
}

func (p *Address) FindAddressByID(db *gorm.DB, pid uint64) (*AddressData, error) {
	var err error
	address := AddressData{}
	err = db.Debug().Model(&AddressData{}).Limit(100).
		Select(`m_address.address_id,
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
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("address_id = ?", pid).
		Find(&address).Error
	if err != nil {
		return &AddressData{}, err
	}
	return &address, nil
}

func (p *Address) FindAddressStatusByID(db *gorm.DB, pid uint64, status uint64) (*AddressData, error) {
	var err error
	address := AddressData{}
	err = db.Debug().Model(&AddressData{}).Limit(100).
		Select(`m_address.address_id,
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
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("address_id = ? AND address_status = ?", pid, status).
		Find(&address).Error
	if err != nil {
		return &AddressData{}, err
	}
	return &address, nil
}

func (p *Address) UpdateAddress(db *gorm.DB, pid uint64) (*Address, error) {
	var err error
	db = db.Debug().Model(&Address{}).Where("address_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"country_id":         p.CountryID,
			"province_id":        p.ProvinceID,
			"regency_id":         p.RegencyID,
			"district_id":        p.DistrictID,
			"village_id":         p.VillageID,
			"address_status":     p.AddressStatus,
			"address_updated_by": p.AddressUpdatedBy,
			"address_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&Address{}).Where("address_id = ?", pid).Error
	if err != nil {
		return &Address{}, err
	}
	return p, nil
}

func (p *Address) UpdateAddressStatus(db *gorm.DB, pid uint64) (*Address, error) {
	var err error
	db = db.Debug().Model(&Address{}).Where("address_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"address_status":     p.AddressStatus,
			"address_updated_by": p.AddressUpdatedBy,
			"address_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&Address{}).Where("address_id = ?", pid).Error
	if err != nil {
		return &Address{}, err
	}
	return p, nil
}

func (p *Address) UpdateAddressStatusOnly(db *gorm.DB, pid uint64) (*Address, error) {
	var err error
	db = db.Debug().Model(&Address{}).Where("address_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"address_status": p.AddressStatus,
		},
	)
	err = db.Debug().Model(&Address{}).Where("address_id = ?", pid).Error
	if err != nil {
		return &Address{}, err
	}
	return p, nil
}

func (p *Address) UpdateAddressStatusDelete(db *gorm.DB, pid uint64) (*Address, error) {
	var err error
	db = db.Debug().Model(&Address{}).Where("address_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"address_status":     3,
			"address_deleted_by": p.AddressDeletedBy,
			"address_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&Address{}).Where("address_id = ?", pid).Error
	if err != nil {
		return &Address{}, err
	}
	return p, nil
}

func (p *Address) DeleteAddress(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&Address{}).Where("address_id = ?", pid).Take(&Address{}).Delete(&Address{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Address not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

type AddressNil struct {
	AddressID        uint64     `gorm:"column:address_id;primary_key;" json:"address_id"`
	CountryID        *uint64    `gorm:"column:country_id" json:"country_id"`
	ProvinceID       *uint64    `gorm:"column:province_id" json:"province_id"`
	RegencyID        *uint64    `gorm:"column:regency_id" json:"regency_id"`
	DistrictID       *uint64    `gorm:"column:district_id" json:"district_id"`
	VillageID        *uint64    `gorm:"column:village_id" json:"village_id"`
	AddressStatus    uint64     `gorm:"column:address_status;size:2" json:"address_status"`
	AddressCreatedBy string     `gorm:"column:address_created_by;size:125" json:"address_created_by"`
	AddressCreatedAt time.Time  `gorm:"column:address_created_at;default:CURRENT_TIMESTAMP" json:"address_created_at"`
	AddressUpdatedBy string     `gorm:"column:address_updated_by;size:125" json:"address_updated_by"`
	AddressUpdatedAt *time.Time `gorm:"column:address_updated_at;default:CURRENT_TIMESTAMP" json:"address_created_at"`
	AddressDeletedBy string     `gorm:"column:address_deleted_by;size:125" json:"address_deleted_by"`
	AddressDeletedAt *time.Time `gorm:"column:address_deleted_at;default:CURRENT_TIMESTAMP" json:"address_deleted_at"`
}

func (AddressNil) TableName() string {
	return "m_address"
}

func (p *AddressNil) Prepare() {
	p.CountryID = nil
	p.ProvinceID = nil
	p.RegencyID = nil
	p.DistrictID = nil
	p.VillageID = nil
	p.AddressStatus = p.AddressStatus
	p.AddressCreatedBy = html.EscapeString(strings.TrimSpace(p.AddressCreatedBy))
	p.AddressCreatedAt = time.Now()
	p.AddressUpdatedBy = html.EscapeString(strings.TrimSpace(p.AddressUpdatedBy))
	p.AddressUpdatedAt = p.AddressUpdatedAt
	p.AddressDeletedBy = html.EscapeString(strings.TrimSpace(p.AddressDeletedBy))
	p.AddressDeletedAt = p.AddressDeletedAt
}

func (p *AddressNil) SaveAddresNil(db *gorm.DB) (*AddressNil, error) {
	var err error
	err = db.Debug().Model(&AddressNil{}).Create(&p).Error
	if err != nil {
		return &AddressNil{}, err
	}
	return p, nil
}
