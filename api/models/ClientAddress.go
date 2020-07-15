package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type ClientAddress struct {
	ClientAddressID            *uint64    `gorm:"column:client_adrs_id;primary_key;" json:"client_adrs_id"`
	ClientAddressStreet        string     `gorm:"column:client_adrs_street" json:"client_adrs_street"`
	ClientAddressRT            string     `gorm:"column:client_adrs_rt;size:3" json:"client_adrs_rt"`
	ClientAddressRW            string     `gorm:"column:client_adrs_rw;size:3" json:"client_adrs_rw"`
	AddressID                  uint64     `gorm:"column:address_id;" json:"address_id"`
	ClientAddressOwnership     string     `gorm:"column:client_adrs_ownership;size:125" json:"client_adrs_ownership"`
	ClientAddressLengthStay    uint64     `gorm:"column:client_adrs_length_stay;size:2" json:"client_adrs_length_stay"`
	ClientAddressAsResidential uint64     `gorm:"column:client_adrs_as_residential;size:1" json:"client_adrs_as_residential"`
	ClientAddressStatus        uint64     `gorm:"column:client_adrs_status;size:1" json:"client_adrs_status"`
	ClientAddressCreatedBy     string     `gorm:"column:client_adrs_created_by;size:125" json:"client_adrs_created_by"`
	ClientAddressCreatedAt     time.Time  `gorm:"column:client_adrs_created_at;default:CURRENT_TIMESTAMP" json:"client_adrs_created_at"`
	ClientAddressUpdatedBy     string     `gorm:"column:client_adrs_updated_by;size:125" json:"client_adrs_updated_by"`
	ClientAddressUpdatedAt     *time.Time `gorm:"column:client_adrs_updated_at;default:CURRENT_TIMESTAMP" json:"client_adrs_updated_at"`
	ClientAddressDeletedBy     string     `gorm:"column:client_adrs_deleted_by;size:125" json:"client_adrs_deleted_by"`
	ClientAddressDeletedAt     *time.Time `gorm:"column:client_adrs_deleted_at;default:CURRENT_TIMESTAMP" json:"client_adrs_deleted_at"`
}

type ClientAddressData struct {
	ClientAddressID            uint64     `gorm:"column:client_adrs_id" json:"client_adrs_id"`
	ClientAddressStreet        string     `gorm:"column:client_adrs_street" json:"client_adrs_street"`
	ClientAddressRT            string     `gorm:"column:client_adrs_rt" json:"client_adrs_rt"`
	ClientAddressRW            string     `gorm:"column:client_adrs_rw" json:"client_adrs_rw"`
	AddressID                  uint64     `gorm:"column:address_id;" json:"address_id"`
	CountryID                  uint64     `gorm:"column:country_id" json:"country_id"`
	CountryName                string     `gorm:"column:country_name" json:"country_name"`
	ProvinceID                 uint64     `gorm:"column:province_id" json:"province_id"`
	ProvinceName               string     `gorm:"column:province_name" json:"province_name"`
	RegencyID                  uint64     `gorm:"column:regency_id" json:"regency_id"`
	RegencyName                string     `gorm:"column:regency_name" json:"regency_name"`
	DistrictID                 uint64     `gorm:"column:district_id" json:"district_id"`
	DistrictName               string     `gorm:"column:district_name" json:"district_name"`
	VillageID                  uint64     `gorm:"column:village_id" json:"village_id"`
	VillageName                string     `gorm:"column:village_name" json:"village_name"`
	ClientAddressOwnership     string     `gorm:"column:client_adrs_ownership" json:"client_adrs_ownership"`
	ClientAddressLengthStay    uint64     `gorm:"column:client_adrs_length_stay" json:"client_adrs_length_stay"`
	ClientAddressAsResidential uint64     `gorm:"column:client_adrs_as_residential" json:"client_adrs_as_residential"`
	ClientAddressStatus        uint64     `gorm:"column:client_adrs_status" json:"client_adrs_status"`
	ClientAddressCreatedBy     string     `gorm:"column:client_adrs_created_by" json:"client_adrs_created_by"`
	ClientAddressCreatedAt     time.Time  `gorm:"column:client_adrs_created_at" json:"client_adrs_created_at"`
	ClientAddressUpdatedBy     string     `gorm:"column:client_adrs_updated_by" json:"client_adrs_updated_by"`
	ClientAddressUpdatedAt     *time.Time `gorm:"column:client_adrs_updated_at" json:"client_adrs_updated_at"`
	ClientAddressDeletedBy     string     `gorm:"column:client_adrs_deleted_by" json:"client_adrs_deleted_by"`
	ClientAddressDeletedAt     *time.Time `gorm:"column:client_adrs_deleted_at" json:"client_adrs_deleted_at"`
}

type ResponseClientAddresss struct {
	Status  uint64               `json:"status"`
	Message string               `json:"message"`
	Data    *[]ClientAddressData `json:"data"`
}

type ResponseClientAddress struct {
	Status  uint64             `json:"status"`
	Message string             `json:"message"`
	Data    *ClientAddressData `json:"data"`
}

type ResponseClientAddressCUD struct {
	Status  uint64 `json:"status"`
	Message string `json:"message"`
}

func (ClientAddress) TableName() string {
	return "m_client_address"
}

func (ClientAddressData) TableName() string {
	return "m_client_address"
}

func (u *ClientAddress) Prepare() {
	u.ClientAddressID = nil
	u.ClientAddressStreet = html.EscapeString(strings.TrimSpace(u.ClientAddressStreet))
	u.ClientAddressRT = html.EscapeString(strings.TrimSpace(u.ClientAddressRT))
	u.ClientAddressRW = html.EscapeString(strings.TrimSpace(u.ClientAddressRW))
	u.AddressID = u.AddressID
	u.ClientAddressOwnership = html.EscapeString(strings.TrimSpace(u.ClientAddressOwnership))
	u.ClientAddressLengthStay = u.ClientAddressLengthStay
	u.ClientAddressAsResidential = u.ClientAddressAsResidential
	u.ClientAddressStatus = u.ClientAddressStatus
	u.ClientAddressCreatedBy = u.ClientAddressCreatedBy
	u.ClientAddressCreatedAt = time.Now()
	u.ClientAddressUpdatedBy = u.ClientAddressUpdatedBy
	u.ClientAddressUpdatedAt = u.ClientAddressUpdatedAt
	u.ClientAddressDeletedBy = u.ClientAddressDeletedBy
	u.ClientAddressDeletedAt = u.ClientAddressDeletedAt
}

func (u *ClientAddress) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if u.ClientAddressStreet == "" {
			return errors.New("Required User Code")
		}
		if u.AddressID == 0 {
			return errors.New("Required Email")
		}
		return nil
	}
}

func (u *ClientAddress) SaveClientAddress(db *gorm.DB) (*ClientAddress, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &ClientAddress{}, err
	}
	return u, nil
}

func (p *ClientAddress) FindAllClientAddress(db *gorm.DB) (*[]ClientAddressData, error) {
	var err error
	clientaddress := []ClientAddressData{}
	err = db.Debug().Model(&ClientAddressData{}).Limit(100).
		Select(`m_client_address.client_adrs_id,
				m_client_address.client_adrs_street,
				m_client_address.client_adrs_rt,
				m_client_address.client_adrs_rw,
				m_client_address.address_id,
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
				m_client_address.client_adrs_ownership,
				m_client_address.client_adrs_length_stay,
				m_client_address.client_adrs_as_residential,
				m_client_address.client_adrs_status,
				m_client_address.client_adrs_created_by,
				m_client_address.client_adrs_created_at,
				m_client_address.client_adrs_updated_by,
				m_client_address.client_adrs_updated_at,
				m_client_address.client_adrs_deleted_by,
				m_client_address.client_adrs_deleted_at`).
		Joins("JOIN m_address on m_client_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Find(&clientaddress).Error
	if err != nil {
		return &[]ClientAddressData{}, err
	}
	return &clientaddress, nil
}

func (p *ClientAddress) FindAllClientAddressStatus(db *gorm.DB, status uint64) (*[]ClientAddressData, error) {
	var err error
	clientaddress := []ClientAddressData{}
	err = db.Debug().Model(&ClientAddressData{}).Limit(100).
		Select(`m_client_address.client_adrs_id,
				m_client_address.client_adrs_street,
				m_client_address.client_adrs_rt,
				m_client_address.client_adrs_rw,
				m_client_address.address_id,
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
				m_client_address.client_adrs_ownership,
				m_client_address.client_adrs_length_stay,
				m_client_address.client_adrs_as_residential,
				m_client_address.client_adrs_status,
				m_client_address.client_adrs_created_by,
				m_client_address.client_adrs_created_at,
				m_client_address.client_adrs_updated_by,
				m_client_address.client_adrs_updated_at,
				m_client_address.client_adrs_deleted_by,
				m_client_address.client_adrs_deleted_at`).
		Joins("JOIN m_address on m_client_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("client_adrs_status = ?", status).
		Find(&clientaddress).Error
	if err != nil {
		return &[]ClientAddressData{}, err
	}
	return &clientaddress, nil
}

func (p *ClientAddress) FindClientAddressByID(db *gorm.DB, pid uint64) (*ClientAddressData, error) {
	var err error
	clientaddress := ClientAddressData{}
	err = db.Debug().Model(&ClientAddressData{}).Limit(100).
		Select(`m_client_address.client_adrs_id,
				m_client_address.client_adrs_street,
				m_client_address.client_adrs_rt,
				m_client_address.client_adrs_rw,
				m_client_address.address_id,
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
				m_client_address.client_adrs_ownership,
				m_client_address.client_adrs_length_stay,
				m_client_address.client_adrs_as_residential,
				m_client_address.client_adrs_status,
				m_client_address.client_adrs_created_by,
				m_client_address.client_adrs_created_at,
				m_client_address.client_adrs_updated_by,
				m_client_address.client_adrs_updated_at,
				m_client_address.client_adrs_deleted_by,
				m_client_address.client_adrs_deleted_at`).
		Joins("JOIN m_address on m_client_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("client_adrs_id = ?", pid).
		Find(&clientaddress).Error
	if err != nil {
		return &ClientAddressData{}, err
	}
	return &clientaddress, nil
}

func (p *ClientAddress) FindClientAddressStatusByID(db *gorm.DB, pid uint64, status uint64) (*ClientAddressData, error) {
	var err error
	clientaddress := ClientAddressData{}
	err = db.Debug().Model(&ClientAddressData{}).Limit(100).
		Select(`m_client_address.client_adrs_id,
				m_client_address.client_adrs_street,
				m_client_address.client_adrs_rt,
				m_client_address.client_adrs_rw,
				m_client_address.address_id,
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
				m_client_address.client_adrs_ownership,
				m_client_address.client_adrs_length_stay,
				m_client_address.client_adrs_as_residential,
				m_client_address.client_adrs_status,
				m_client_address.client_adrs_created_by,
				m_client_address.client_adrs_created_at,
				m_client_address.client_adrs_updated_by,
				m_client_address.client_adrs_updated_at,
				m_client_address.client_adrs_deleted_by,
				m_client_address.client_adrs_deleted_at`).
		Joins("JOIN m_address on m_client_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("client_adrs_id = ? AND client_adrs_status = ?", pid, status).
		Find(&clientaddress).Error
	if err != nil {
		return &ClientAddressData{}, err
	}
	return &clientaddress, nil
}

func (u *ClientAddress) UpdateClientAddress(db *gorm.DB, uid uint64) (*ClientAddress, error) {
	var err error
	db = db.Debug().Model(&ClientAddress{}).Where("client_adrs_id = ?", uid).Take(&ClientAddress{}).UpdateColumns(
		map[string]interface{}{
			"client_adrs_street":      u.ClientAddressStreet,
			"client_adrs_rt":          u.ClientAddressRT,
			"client_adrs_rw":          u.ClientAddressRW,
			"address_id":              u.AddressID,
			"client_adrs_ownership":   u.ClientAddressOwnership,
			"client_adrs_length_stay": u.ClientAddressLengthStay,
			"client_adrs_status":      u.ClientAddressStatus,
			"client_adrs_updated_by":  u.ClientAddressUpdatedBy,
			"client_adrs_updated_at":  time.Now(),
		},
	)
	if db.Error != nil {
		return &ClientAddress{}, db.Error
	}
	err = db.Debug().Model(&ClientAddress{}).Where("client_adrs_id = ?", uid).Take(&u).Error
	if err != nil {
		return &ClientAddress{}, err
	}
	return u, nil
}

func (p *ClientAddress) UpdateClientAddressStatus(db *gorm.DB, pid uint64) (*ClientAddress, error) {

	var err error
	db = db.Debug().Model(&ClientAddress{}).Where("client_adrs_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"client_adrs_status":     p.ClientAddressStatus,
			"client_adrs_updated_by": p.ClientAddressUpdatedBy,
			"client_adrs_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ClientAddress{}).Where("client_adrs_id = ?", pid).Error
	if err != nil {
		return &ClientAddress{}, err
	}
	return p, nil
}

func (p *ClientAddress) UpdateClientAddressStatusDelete(db *gorm.DB, pid uint64) (*ClientAddress, error) {
	var err error
	db = db.Debug().Model(&ClientAddress{}).Where("client_adrs_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"client_adrs_status":     3,
			"client_adrs_deleted_by": p.ClientAddressDeletedBy,
			"client_adrs_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ClientAddress{}).Where("client_adrs_id = ?", pid).Error
	if err != nil {
		return &ClientAddress{}, err
	}
	return p, nil
}

func (p *ClientAddress) DeleteClientAddress(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&ClientAddress{}).Where("client_adrs_id = ?", pid).Take(&ClientAddress{}).Delete(&ClientAddress{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("ClientAddress not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// ADDITIONAL
// ===================================================================================================================================
