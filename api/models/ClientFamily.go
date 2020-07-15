package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type ClientFamily struct {
	ClientFamilyID           uint64     `gorm:"column:client_fmly_id;primary_key;" json:"client_fmly_id"`
	ClientID                 uint64     `gorm:"column:client_id" json:"client_id"`
	ClientAddressID          uint64     `gorm:"column:client_adrs_id" json:"client_adrs_id"`
	ClientFamilyKKNumber     string     `gorm:"column:client_fmly_kk_number;size:20" json:"client_fmly_kk_number"`
	ClientFamilyHeadName     string     `gorm:"column:client_fmly_head_name;size:255" json:"client_fmly_head_name"`
	ClientFamilyNumberMember uint64     `gorm:"column:client_fmly_num_member;" json:"client_fmly_num_member"`
	ClientFamilyStatus       uint64     `gorm:"column:client_fmly_status;size:2" json:"client_fmly_status"`
	ClientFamilyCreatedBy    string     `gorm:"column:client_fmly_created_by;size:125" json:"client_fmly_created_by"`
	ClientFamilyCreatedAt    time.Time  `gorm:"column:client_fmly_created_at;default:CURRENT_TIMESTAMP" json:"client_fmly_created_at"`
	ClientFamilyUpdatedBy    string     `gorm:"column:client_fmly_updated_by;size:125" json:"client_fmly_updated_by"`
	ClientFamilyUpdatedAt    *time.Time `gorm:"column:client_fmly_updated_at;default:CURRENT_TIMESTAMP" json:"client_fmly_updated_at"`
	ClientFamilyDeletedBy    string     `gorm:"column:client_fmly_deleted_by;size:125" json:"client_fmly_deleted_by"`
	ClientFamilyDeletedAt    *time.Time `gorm:"column:client_fmly_deleted_at;default:CURRENT_TIMESTAMP" json:"client_fmly_deleted_at"`
}

type ClientFamilyData struct {
	ClientFamilyID             uint64     `gorm:"column:client_fmly_id;primary_key;" json:"client_fmly_id"`
	ClientID                   uint64     `gorm:"column:client_id" json:"client_id"`
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
	ClientFamilyKKNumber       string     `gorm:"column:client_fmly_kk_number;size:20" json:"client_fmly_kk_number"`
	ClientFamilyHeadName       string     `gorm:"column:client_fmly_head_name;size:255" json:"client_fmly_head_name"`
	ClientFamilyNumberMember   uint64     `gorm:"column:client_fmly_num_member;" json:"client_fmly_num_member"`
	ClientFamilyStatus         uint64     `gorm:"column:client_fmly_status;size:2" json:"client_fmly_status"`
	ClientFamilyCreatedBy      string     `gorm:"column:client_fmly_created_by;size:125" json:"client_fmly_created_by"`
	ClientFamilyCreatedAt      time.Time  `gorm:"column:client_fmly_created_at;default:CURRENT_TIMESTAMP" json:"client_fmly_created_at"`
	ClientFamilyUpdatedBy      string     `gorm:"column:client_fmly_updated_by;size:125" json:"client_fmly_updated_by"`
	ClientFamilyUpdatedAt      *time.Time `gorm:"column:client_fmly_updated_at;default:CURRENT_TIMESTAMP" json:"client_fmly_updated_at"`
	ClientFamilyDeletedBy      string     `gorm:"column:client_fmly_deleted_by;size:125" json:"client_fmly_deleted_by"`
	ClientFamilyDeletedAt      *time.Time `gorm:"column:client_fmly_deleted_at;default:CURRENT_TIMESTAMP" json:"client_fmly_deleted_at"`
}

type ResponseClientFamilys struct {
	Status  uint64              `json:"status"`
	Message string              `json:"message"`
	Data    *[]ClientFamilyData `json:"data"`
}

type ResponseClientFamily struct {
	Status  uint64            `json:"status"`
	Message string            `json:"message"`
	Data    *ClientFamilyData `json:"data"`
}

type ResponseClientFamilyCUD struct {
	Status  uint64 `json:"status"`
	Message string `json:"message"`
}

func (ClientFamily) TableName() string {
	return "m_client_family"
}

func (p *ClientFamily) Prepare() {
	p.ClientID = p.ClientID
	p.ClientAddressID = p.ClientAddressID
	p.ClientFamilyKKNumber = html.EscapeString(strings.TrimSpace(p.ClientFamilyKKNumber))
	p.ClientFamilyHeadName = html.EscapeString(strings.TrimSpace(p.ClientFamilyHeadName))
	p.ClientFamilyNumberMember = p.ClientFamilyNumberMember
	p.ClientFamilyStatus = p.ClientFamilyStatus
	p.ClientFamilyCreatedAt = time.Now()
	p.ClientFamilyCreatedBy = p.ClientFamilyCreatedBy
	p.ClientFamilyUpdatedAt = p.ClientFamilyUpdatedAt
	p.ClientFamilyUpdatedBy = p.ClientFamilyUpdatedBy
	p.ClientFamilyDeletedAt = p.ClientFamilyDeletedAt
	p.ClientFamilyDeletedBy = p.ClientFamilyDeletedBy
}

func (p *ClientFamily) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.ClientAddressID == 0 {
			return errors.New("Required ClientFamily Code")
		}
		return nil
	}
}

func (p *ClientFamily) SaveClientFamily(db *gorm.DB) (*ClientFamily, error) {
	var err error
	err = db.Debug().Model(&ClientFamily{}).Create(&p).Error
	if err != nil {
		return &ClientFamily{}, err
	}
	return p, nil
}

func (p *ClientFamily) FindAllClientFamily(db *gorm.DB) (*[]ClientFamilyData, error) {
	var err error
	client := []ClientFamilyData{}
	err = db.Debug().Model(&ClientFamilyData{}).Limit(100).
		Select(`m_client_family.client_fmly_id,
				m_client_family.client_id,
				m_client_address.client_adrs_id,
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
				m_client_family.client_fmly_kk_number,
				m_client_family.client_fmly_head_name,
				m_client_family.client_fmly_num_member,
				m_client_family.client_fmly_status,
				m_client_family.client_fmly_created_by,
				m_client_family.client_fmly_created_at,
				m_client_family.client_fmly_updated_by,
				m_client_family.client_fmly_updated_at,
				m_client_family.client_fmly_deleted_by,
				m_client_family.client_fmly_deleted_at`).
		Joins("JOIN m_client_address on m_client.client_adrs_id=m_client_address.client_adrs_id").
		Joins("JOIN m_address on m_client_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Find(&client).Error
	if err != nil {
		return &[]ClientFamilyData{}, err
	}
	return &client, nil
}

func (p *ClientFamily) FindAllClientFamilyStatus(db *gorm.DB, status uint64) (*[]ClientFamilyData, error) {
	var err error
	client := []ClientFamilyData{}
	err = db.Debug().Model(&ClientFamilyData{}).Limit(100).
		Select(`m_client_family.client_fmly_id,
				m_client_family.client_id,
				m_client_address.client_adrs_id,
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
				m_client_family.client_fmly_kk_number,
				m_client_family.client_fmly_head_name,
				m_client_family.client_fmly_num_member,
				m_client_family.client_fmly_status,
				m_client_family.client_fmly_created_by,
				m_client_family.client_fmly_created_at,
				m_client_family.client_fmly_updated_by,
				m_client_family.client_fmly_updated_at,
				m_client_family.client_fmly_deleted_by,
				m_client_family.client_fmly_deleted_at`).
		Joins("JOIN m_client_address on m_client.client_adrs_id=m_client_address.client_adrs_id").
		Joins("JOIN m_address on m_client_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("client_fmly_status = ?", status).
		Find(&client).Error
	if err != nil {
		return &[]ClientFamilyData{}, err
	}
	return &client, nil
}

func (p *ClientFamily) FindClientFamilyByID(db *gorm.DB, pid uint64) (*ClientFamilyData, error) {
	var err error
	client := ClientFamilyData{}
	err = db.Debug().Model(&ClientFamilyData{}).Limit(1).
		Select(`m_client_family.client_fmly_id,
				m_client_family.client_id,
				m_client_address.client_adrs_id,
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
				m_client_family.client_fmly_kk_number,
				m_client_family.client_fmly_head_name,
				m_client_family.client_fmly_num_member,
				m_client_family.client_fmly_status,
				m_client_family.client_fmly_created_by,
				m_client_family.client_fmly_created_at,
				m_client_family.client_fmly_updated_by,
				m_client_family.client_fmly_updated_at,
				m_client_family.client_fmly_deleted_by,
				m_client_family.client_fmly_deleted_at`).
		Joins("JOIN m_client_address on m_client.client_adrs_id=m_client_address.client_adrs_id").
		Joins("JOIN m_address on m_client_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("client_fmly_id = ?", pid).
		Find(&client).Error
	if err != nil {
		return &ClientFamilyData{}, err
	}
	return &client, nil
}

func (p *ClientFamily) FindClientFamilyStatusByID(db *gorm.DB, pid uint64, status uint64) (*ClientFamilyData, error) {
	var err error
	client := ClientFamilyData{}
	err = db.Debug().Model(&ClientFamilyData{}).Limit(1).
		Select(`m_client_family.client_fmly_id,
				m_client_family.client_id,
				m_client_address.client_adrs_id,
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
				m_client_family.client_fmly_kk_number,
				m_client_family.client_fmly_head_name,
				m_client_family.client_fmly_num_member,
				m_client_family.client_fmly_status,
				m_client_family.client_fmly_created_by,
				m_client_family.client_fmly_created_at,
				m_client_family.client_fmly_updated_by,
				m_client_family.client_fmly_updated_at,
				m_client_family.client_fmly_deleted_by,
				m_client_family.client_fmly_deleted_at`).
		Joins("JOIN m_client_address on m_client.client_adrs_id=m_client_address.client_adrs_id").
		Joins("JOIN m_address on m_client_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("client_fmly_id = ? AND client_fmly_status = ?", pid, status).
		Find(&client).Error
	if err != nil {
		return &ClientFamilyData{}, err
	}
	return &client, nil
}

func (p *ClientFamily) UpdateClientFamily(db *gorm.DB, pid uint64) (*ClientFamily, error) {
	var err error
	db = db.Debug().Model(&ClientFamily{}).Where("client_fmly_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"client_fmly_kk_number":  p.ClientFamilyKKNumber,
			"client_fmly_head_name":  p.ClientFamilyHeadName,
			"client_fmly_num_member": p.ClientFamilyNumberMember,
			"client_fmly_status":     p.ClientFamilyStatus,
			"client_fmly_updated_by": p.ClientFamilyUpdatedBy,
			"client_fmly_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ClientFamily{}).Where("client_fmly_id = ?", pid).Error
	if err != nil {
		return &ClientFamily{}, err
	}
	return p, nil
}

func (p *ClientFamily) UpdateClientFamilyStatus(db *gorm.DB, pid uint64) (*ClientFamily, error) {
	var err error
	db = db.Debug().Model(&ClientFamily{}).Where("client_fmly_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"client_fmly_status":     p.ClientFamilyStatus,
			"client_fmly_updated_by": p.ClientFamilyUpdatedBy,
			"client_fmly_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ClientFamily{}).Where("client_fmly_id = ?", pid).Error
	if err != nil {
		return &ClientFamily{}, err
	}
	return p, nil
}

func (p *ClientFamily) UpdateClientFamilyStatusDelete(db *gorm.DB, pid uint64) (*ClientFamily, error) {
	var err error
	db = db.Debug().Model(&ClientFamily{}).Where("client_fmly_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"client_fmly_status":     3,
			"client_fmly_deleted_by": p.ClientFamilyDeletedBy,
			"client_fmly_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ClientFamily{}).Where("client_fmly_id = ?", pid).Error
	if err != nil {
		return &ClientFamily{}, err
	}
	return p, nil
}

func (p *ClientFamily) DeleteClientFamily(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&ClientFamily{}).Where("client_fmly_id = ?", pid).Take(&ClientFamily{}).Delete(&ClientFamily{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("ClientFamily not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// ADDITIONAL MEMORANDUM
// ===================================================================================================================================

type ClientFamilyMrdmData struct {
	ClientFamilyID             uint64                       `gorm:"column:client_fmly_id" json:"client_fmly_id"`
	ClientID                   uint64                       `gorm:"column:client_id" json:"client_id"`
	ClientAddressID            uint64                       `gorm:"column:client_adrs_id" json:"client_adrs_id"`
	ClientAddressStreet        string                       `gorm:"column:client_adrs_street" json:"client_adrs_street"`
	ClientAddressRT            string                       `gorm:"column:client_adrs_rt" json:"client_adrs_rt"`
	ClientAddressRW            string                       `gorm:"column:client_adrs_rw" json:"client_adrs_rw"`
	AddressID                  uint64                       `gorm:"column:address_id;" json:"address_id"`
	CountryID                  uint64                       `gorm:"column:country_id" json:"country_id"`
	CountryName                string                       `gorm:"column:country_name" json:"country_name"`
	ProvinceID                 uint64                       `gorm:"column:province_id" json:"province_id"`
	ProvinceName               string                       `gorm:"column:province_name" json:"province_name"`
	RegencyID                  uint64                       `gorm:"column:regency_id" json:"regency_id"`
	RegencyName                string                       `gorm:"column:regency_name" json:"regency_name"`
	DistrictID                 uint64                       `gorm:"column:district_id" json:"district_id"`
	DistrictName               string                       `gorm:"column:district_name" json:"district_name"`
	VillageID                  uint64                       `gorm:"column:village_id" json:"village_id"`
	VillageName                string                       `gorm:"column:village_name" json:"village_name"`
	ClientAddressOwnership     string                       `gorm:"column:client_adrs_ownership" json:"client_adrs_ownership"`
	ClientAddressLengthStay    uint64                       `gorm:"column:client_adrs_length_stay" json:"client_adrs_length_stay"`
	ClientAddressAsResidential uint64                       `gorm:"column:client_adrs_as_residential" json:"client_adrs_as_residential"`
	ClientFamilyKKNumber       string                       `gorm:"column:client_fmly_kk_number;size:20" json:"client_fmly_kk_number"`
	ClientFamilyHeadName       string                       `gorm:"column:client_fmly_head_name;size:255" json:"client_fmly_head_name"`
	ClientFamilyNumberMember   uint64                       `gorm:"column:client_fmly_num_member;" json:"client_fmly_num_member"`
	ClientFamilyMember         []ClientFamilyMemberMrdmData `gorm:"column:client_fmly_mmbr;" json:"client_fmly_mmbr"`
}

func (ClientFamilyMrdmData) TableName() string {
	return "m_client_family"
}

func (p *ClientFamily) FindClientFamilyMrdmByClientID(db *gorm.DB, pid uint64) (ClientFamilyMrdmData, error) {
	var err error
	clientfamily := ClientFamilyMrdmData{}
	err = db.Debug().Model(&ClientFamilyMrdmData{}).Limit(1).
		Select(`m_client_family.client_fmly_id,
				m_client_family.client_id,
				m_client_address.client_adrs_id,
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
				m_client_family.client_fmly_kk_number,
				m_client_family.client_fmly_head_name,
				m_client_family.client_fmly_num_member,
				m_client_family.client_fmly_status,
				m_client_family.client_fmly_created_by,
				m_client_family.client_fmly_created_at,
				m_client_family.client_fmly_updated_by,
				m_client_family.client_fmly_updated_at,
				m_client_family.client_fmly_deleted_by,
				m_client_family.client_fmly_deleted_at`).
		Joins("JOIN m_client_address on m_client_family.client_adrs_id=m_client_address.client_adrs_id").
		Joins("JOIN m_address on m_client_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("client_id = ?", pid).
		Find(&clientfamily).Error
	if err != nil {
		return ClientFamilyMrdmData{}, err
	}
	return clientfamily, nil
}
