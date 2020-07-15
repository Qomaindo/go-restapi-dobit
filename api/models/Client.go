package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Client struct {
	ClientID            uint64     `gorm:"column:client_id;primary_key;" json:"client_id"`
	CustomerUserID      uint64     `gorm:"column:cust_user_id" json:"cust_user_id"`
	ClientAddressID     uint64     `gorm:"column:client_adrs_id" json:"client_adrs_id"`
	ClientCode          string     `gorm:"column:client_code;size:25" json:"client_code"`
	ClientKTP           string     `gorm:"column:client_ktp;size:255" json:"client_ktp"`
	ClientPassport      string     `gorm:"column:client_passport;size:125" json:"client_passport"`
	ClientNPWP          string     `gorm:"column:client_npwp;size:255" json:"client_npwp"`
	ClientName          string     `gorm:"column:client_name;size:255" json:"client_name"`
	ClientSex           string     `gorm:"column:client_sex;size:6" json:"client_sex"`
	ClientBirthPlace    string     `gorm:"column:client_birth_place;size:255" json:"client:"client_birth_place"`
	ClientBirthDate     string     `gorm:"column:client_birth_date;size:25" json:"client_birth_date"`
	ClientAge           string     `gorm:"column:client_age;size:3" json:"client_age"`
	ClientPhone         string     `gorm:"column:client_phone;size:20" json:"client_phone"`
	ClientPhoneWork     string     `gorm:"column:client_phone_work;size:20" json:"client_phone_work"`
	ClientEmail         string     `gorm:"column:client_email;size:255" json:"client_email"`
	ClientEmailWork     string     `gorm:"column:client_email_work;size:255" json:"client_email_work"`
	ClientMaritalStatus string     `gorm:"column:client_marital_status;size:125" json:"client_marital_status"`
	ClientImage         string     `gorm:"column:client_image;size:255" json:"client_image"`
	ClientDescription   string     `gorm:"column:client_desc" json:"client_desc"`
	ClientStatus        uint64     `gorm:"column:client_status;size:1" json:"client_status"`
	ClientCreatedAt     time.Time  `gorm:"column:client_created_at;default:CURRENT_TIMESTAMP" json:"client_created_at"`
	ClientCreatedBy     string     `gorm:"column:client_created_by;size:125" json:"client_created_by"`
	ClientUpdatedAt     *time.Time `gorm:"column:client_updated_at;default:CURRENT_TIMESTAMP" json:"client_updated_at"`
	ClientUpdatedBy     string     `gorm:"column:client_updated_by;size:125" json:"client_updated_by"`
	ClientDeletedAt     *time.Time `gorm:"column:client_deleted_at;default:CURRENT_TIMESTAMP" json:"client_deleted_at"`
	ClientDeletedBy     string     `gorm:"column:client_deleted_by;size:125" json:"client_deleted_by"`
}

type ClientData struct {
	ClientID                   uint64     `gorm:"column:client_id;primary_key;" json:"client_id"`
	CustomerUserID             uint64     `gorm:"column:cust_user_id" json:"cust_user_id"`
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
	ClientCode                 string     `gorm:"column:client_code" json:"client_code"`
	ClientKTP                  string     `gorm:"column:client_ktp" json:"client_ktp"`
	ClientPassport             string     `gorm:"column:client_passport" json:"client_passport"`
	ClientNPWP                 string     `gorm:"column:client_npwp" json:"client_npwp"`
	ClientName                 string     `gorm:"column:client_name" json:"client_name"`
	ClientSex                  string     `gorm:"column:client_sex" json:"client_sex"`
	ClientBirthPlace           string     `gorm:"column:client_birth_place" json:"client:"client_birth_place"`
	ClientBirthDate            string     `gorm:"column:client_birth_date" json:"client_birth_date"`
	ClientAge                  string     `gorm:"column:client_age" json:"client_age"`
	ClientPhone                string     `gorm:"column:client_phone" json:"client_phone"`
	ClientPhoneWork            string     `gorm:"column:client_phone_work" json:"client_phone_work"`
	ClientEmail                string     `gorm:"column:client_email" json:"client_email"`
	ClientEmailWork            string     `gorm:"column:client_email_work" json:"client_email_work"`
	ClientMaritalStatus        string     `gorm:"column:client_marital_status" json:"client_marital_status"`
	ClientImage                string     `gorm:"column:client_image" json:"client_image"`
	ClientDescription          string     `gorm:"column:client_desc" json:"client_desc"`
	ClientStatus               uint64     `gorm:"column:client_status" json:"client_status"`
	ClientCreatedAt            time.Time  `gorm:"column:client_created_at" json:"client_created_at"`
	ClientCreatedBy            string     `gorm:"column:client_created_by" json:"client_created_by"`
	ClientUpdatedAt            *time.Time `gorm:"column:client_updated_at" json:"client_updated_at"`
	ClientUpdatedBy            string     `gorm:"column:client_updated_by" json:"client_updated_by"`
	ClientDeletedAt            *time.Time `gorm:"column:client_deleted_at" json:"client_deleted_at"`
	ClientDeletedBy            string     `gorm:"column:client_deleted_by" json:"client_deleted_by"`
}

type ResponseClients struct {
	Status  uint64        `json:"status"`
	Message string        `json:"message"`
	Data    *[]ClientData `json:"data"`
}

type ResponseClient struct {
	Status  uint64      `json:"status"`
	Message string      `json:"message"`
	Data    *ClientData `json:"data"`
}

type ResponseClientCUD struct {
	Status  uint64 `json:"status"`
	Message string `json:"message"`
}

func (Client) TableName() string {
	return "m_client"
}

func (ClientData) TableName() string {
	return "m_client"
}

func (p *Client) Prepare() {
	p.CustomerUserID = p.CustomerUserID
	p.ClientAddressID = p.ClientAddressID
	p.ClientCode = html.EscapeString(strings.TrimSpace(p.ClientCode))
	p.ClientKTP = html.EscapeString(strings.TrimSpace(p.ClientKTP))
	p.ClientPassport = html.EscapeString(strings.TrimSpace(p.ClientPassport))
	p.ClientNPWP = html.EscapeString(strings.TrimSpace(p.ClientNPWP))
	p.ClientName = html.EscapeString(strings.TrimSpace(p.ClientName))
	p.ClientSex = html.EscapeString(strings.TrimSpace(p.ClientSex))
	p.ClientBirthPlace = html.EscapeString(strings.TrimSpace(p.ClientBirthPlace))
	p.ClientBirthDate = html.EscapeString(strings.TrimSpace(p.ClientBirthDate))
	p.ClientAge = html.EscapeString(strings.TrimSpace(p.ClientAge))
	p.ClientPhone = html.EscapeString(strings.TrimSpace(p.ClientPhone))
	p.ClientPhoneWork = html.EscapeString(strings.TrimSpace(p.ClientPhoneWork))
	p.ClientEmail = html.EscapeString(strings.TrimSpace(p.ClientEmail))
	p.ClientEmailWork = html.EscapeString(strings.TrimSpace(p.ClientEmailWork))
	p.ClientMaritalStatus = html.EscapeString(strings.TrimSpace(p.ClientMaritalStatus))
	p.ClientImage = html.EscapeString(strings.TrimSpace(p.ClientImage))
	p.ClientDescription = html.EscapeString(strings.TrimSpace(p.ClientDescription))
	p.ClientStatus = p.ClientStatus
	p.ClientCreatedAt = time.Now()
	p.ClientCreatedBy = p.ClientCreatedBy
	p.ClientUpdatedAt = p.ClientUpdatedAt
	p.ClientUpdatedBy = p.ClientUpdatedBy
	p.ClientDeletedAt = p.ClientDeletedAt
	p.ClientDeletedBy = p.ClientDeletedBy
}

func (p *Client) Validate(action string) error {
	switch strings.ToLower(action) {
	case "register":
		if p.ClientName == "" {
			return errors.New("Requred Client Name")
		}
		return nil

	default:
		if p.ClientCode == "" {
			return errors.New("Requred Client Code")
		}
		if p.ClientKTP == "" {
			return errors.New("Requred Client KTP Number")
		}
		if p.ClientPassport == "" {
			return errors.New("Requred Client Passport Number")
		}
		if p.ClientNPWP == "" {
			return errors.New("Requred Client NPWP Number")
		}
		if p.ClientName == "" {
			return errors.New("Requred Client Name")
		}
		if p.ClientSex == "" {
			return errors.New("Requred Client Sex")
		}
		if p.ClientBirthPlace == "" {
			return errors.New("Requred Client Birth Place")
		}
		if p.ClientBirthDate == "" {
			return errors.New("Requred Client Birth Date")
		}
		if p.ClientAge == "" {
			return errors.New("Requred Client Age")
		}
		if p.ClientPhoneWork == "" {
			return errors.New("Requred Client Phone Number Work")
		}
		if p.ClientEmailWork == "" {
			return errors.New("Requred Client Email Address Work")
		}
		if p.ClientMaritalStatus == "" {
			return errors.New("Requred Marital Status")
		}
		if p.ClientImage == "" {
			return errors.New("Requred Image Client")
		}
		return nil
	}
}

func (p *Client) SaveClient(db *gorm.DB) (*Client, error) {
	var err error
	err = db.Debug().Model(&Client{}).Create(&p).Error
	if err != nil {
		return &Client{}, err
	}
	return p, nil
}

func (p *Client) FindAllClient(db *gorm.DB) (*[]ClientData, error) {
	var err error
	client := []ClientData{}
	err = db.Debug().Model(&ClientData{}).Limit(100).
		Select(`m_client.client_id,
				m_client.cust_user_id,
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
				m_client.client_code,
				m_client.client_ktp,
				m_client.client_passport,
				m_client.client_npwp,
				m_client.client_name,
				m_client.client_sex,
				m_client.client_birth_place,
				m_client.client_birth_date,
				m_client.client_age,
				m_client.client_phone,
				m_client.client_phone_work,
				m_client.client_email,
				m_client.client_email_work,
				m_client.client_marital_status,
				m_client.client_image,
				m_client.client_desc,
				m_client.client_status,
				m_client.client_created_at,
				m_client.client_created_by,
				m_client.client_updated_at,
				m_client.client_updated_by,
				m_client.client_deleted_at,
				m_client.client_deleted_by`).
		Joins("JOIN m_client_address on m_client.client_adrs_id=m_client_address.client_adrs_id").
		Joins("JOIN m_address on m_client_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Find(&client).Error
	if err != nil {
		return &[]ClientData{}, err
	}
	return &client, nil
}

func (p *Client) FindAllClientStatus(db *gorm.DB, status uint64) (*[]ClientData, error) {
	var err error
	client := []ClientData{}
	err = db.Debug().Model(&ClientData{}).Limit(100).
		Select(`m_client.client_id,
				m_client.cust_user_id,
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
				m_client.client_code,
				m_client.client_ktp,
				m_client.client_passport,
				m_client.client_npwp,
				m_client.client_name,
				m_client.client_sex,
				m_client.client_birth_place,
				m_client.client_birth_date,
				m_client.client_age,
				m_client.client_phone,
				m_client.client_phone_work,
				m_client.client_email,
				m_client.client_email_work,
				m_client.client_marital_status,
				m_client.client_image,
				m_client.client_desc,
				m_client.client_status,
				m_client.client_created_at,
				m_client.client_created_by,
				m_client.client_updated_at,
				m_client.client_updated_by,
				m_client.client_deleted_at,
				m_client.client_deleted_by`).
		Joins("JOIN m_client_address on m_client.client_adrs_id=m_client_address.client_adrs_id").
		Joins("JOIN m_address on m_client_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("client_status = ?", status).
		Find(&client).Error
	if err != nil {
		return &[]ClientData{}, err
	}
	return &client, nil
}

func (p *Client) FindClientByID(db *gorm.DB, pid uint64) (*ClientData, error) {
	var err error
	client := ClientData{}
	err = db.Debug().Model(&ClientData{}).Limit(1).
		Select(`m_client.client_id,
				m_client.cust_user_id,
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
				m_client.client_code,
				m_client.client_ktp,
				m_client.client_passport,
				m_client.client_npwp,
				m_client.client_name,
				m_client.client_sex,
				m_client.client_birth_place,
				m_client.client_birth_date,
				m_client.client_age,
				m_client.client_phone,
				m_client.client_phone_work,
				m_client.client_email,
				m_client.client_email_work,
				m_client.client_marital_status,
				m_client.client_image,
				m_client.client_desc,
				m_client.client_status,
				m_client.client_created_at,
				m_client.client_created_by,
				m_client.client_updated_at,
				m_client.client_updated_by,
				m_client.client_deleted_at,
				m_client.client_deleted_by`).
		Joins("JOIN m_client_address on m_client.client_adrs_id=m_client_address.client_adrs_id").
		Joins("JOIN m_address on m_client_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("client_id = ?", pid).
		Find(&client).Error
	if err != nil {
		return &ClientData{}, err
	}
	return &client, nil
}

func (p *Client) FindClientStatusByID(db *gorm.DB, pid uint64, status uint64) (*ClientData, error) {
	var err error
	client := ClientData{}
	err = db.Debug().Model(&ClientData{}).Limit(1).
		Select(`m_client.client_id,
				m_client.cust_user_id,
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
				m_client.client_code,
				m_client.client_ktp,
				m_client.client_passport,
				m_client.client_npwp,
				m_client.client_name,
				m_client.client_sex,
				m_client.client_birth_place,
				m_client.client_birth_date,
				m_client.client_age,
				m_client.client_phone,
				m_client.client_phone_work,
				m_client.client_email,
				m_client.client_email_work,
				m_client.client_marital_status,
				m_client.client_image,
				m_client.client_desc,
				m_client.client_status,
				m_client.client_created_at,
				m_client.client_created_by,
				m_client.client_updated_at,
				m_client.client_updated_by,
				m_client.client_deleted_at,
				m_client.client_deleted_by`).
		Joins("JOIN m_client_address on m_client.client_adrs_id=m_client_address.client_adrs_id").
		Joins("JOIN m_address on m_client_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("client_id = ? AND client_status = ?", pid, status).
		Find(&client).Error
	if err != nil {
		return &ClientData{}, err
	}
	return &client, nil
}

func (p *Client) UpdateClient(db *gorm.DB, pid uint64) (*Client, error) {
	var err error
	db = db.Debug().Model(&Client{}).Where("client_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"client_ktp":            p.ClientKTP,
			"client_passport":       p.ClientPassport,
			"client_npwp":           p.ClientNPWP,
			"client_name":           p.ClientName,
			"client_sex":            p.ClientSex,
			"client_birth_date":     p.ClientBirthDate,
			"client_birth_place":    p.ClientBirthPlace,
			"client_age":            p.ClientAge,
			"client_phone":          p.ClientPhone,
			"client_phone_work":     p.ClientPhoneWork,
			"client_email":          p.ClientEmail,
			"client_email_work":     p.ClientEmailWork,
			"client_marital_status": p.ClientMaritalStatus,
			"client_image":          p.ClientImage,
			"client_status":         p.ClientStatus,
			"client_updated_by":     p.ClientUpdatedBy,
			"client_updated_at":     time.Now(),
		},
	)
	err = db.Debug().Model(&Client{}).Where("client_id = ?", pid).Error
	if err != nil {
		return &Client{}, err
	}
	return p, nil
}

func (p *Client) UpdateClientStatus(db *gorm.DB, pid uint64) (*Client, error) {
	var err error
	db = db.Debug().Model(&Client{}).Where("client_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"client_status":     p.ClientStatus,
			"client_updated_by": p.ClientUpdatedBy,
			"client_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&Client{}).Where("client_id = ?", pid).Error
	if err != nil {
		return &Client{}, err
	}
	return p, nil
}

func (p *Client) UpdateClientStatusDelete(db *gorm.DB, pid uint64) (*Client, error) {
	var err error
	db = db.Debug().Model(&Client{}).Where("client_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"client_status":     3,
			"client_deleted_by": p.ClientDeletedBy,
			"client_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&Client{}).Where("client_id = ?", pid).Error
	if err != nil {
		return &Client{}, err
	}
	return p, nil
}

func (p *Client) DeleteClient(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&Client{}).Where("client_id = ?", pid).Take(&Client{}).Delete(&Client{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Client not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// ADDITIONAL MEMORANDUM
// ===================================================================================================================================

type ClientMrdmData struct {
	ClientID                   uint64 `gorm:"column:client_id;" json:"client_id"`
	ClientKTP                  string `gorm:"column:client_ktp" json:"client_ktp"`
	ClientPassport             string `gorm:"column:client_passport" json:"client_passport"`
	ClientNPWP                 string `gorm:"column:client_npwp" json:"client_npwp"`
	ClientName                 string `gorm:"column:client_name" json:"client_name"`
	ClientSex                  string `gorm:"column:client_sex" json:"client_sex"`
	ClientBirthDate            string `gorm:"column:client_birth_date" json:"client_birth_date"`
	ClientBirthPlace           string `gorm:"column:client_birth_place" json:"client_birth_place"`
	ClientAge                  string `gorm:"column:client_age" json:"client_age"`
	ClientPhone                string `gorm:"column:client_phone" json:"client_phone"`
	ClientPhoneWork            string `gorm:"column:client_phone_work" json:"client_phone_work"`
	ClientEmail                string `gorm:"column:client_email" json:"client_email"`
	ClientEmailWork            string `gorm:"column:client_email_work" json:"client_email_work"`
	ClientMaritalStatus        string `gorm:"column:client_marital_status" json:"client_marital_status"`
	ClientImage                string `gorm:"column:client_image" json:"client_image"`
	ClientDescription          string `gorm:"column:client_desc" json:"client_desc"`
	ClientAddressStreet        string `gorm:"column:client_adrs_street" json:"client_adrs_street"`
	ClientAddressRT            string `gorm:"column:client_adrs_rt" json:"client_adrs_rt"`
	ClientAddressRW            string `gorm:"column:client_adrs_rw" json:"client_adrs_rw"`
	AddressID                  uint64 `gorm:"column:address_id;" json:"address_id"`
	CountryID                  uint64 `gorm:"column:country_id" json:"country_id"`
	CountryName                string `gorm:"column:country_name" json:"country_name"`
	ProvinceID                 uint64 `gorm:"column:province_id" json:"province_id"`
	ProvinceName               string `gorm:"column:province_name" json:"province_name"`
	RegencyID                  uint64 `gorm:"column:regency_id" json:"regency_id"`
	RegencyName                string `gorm:"column:regency_name" json:"regency_name"`
	DistrictID                 uint64 `gorm:"column:district_id" json:"district_id"`
	DistrictName               string `gorm:"column:district_name" json:"district_name"`
	VillageID                  uint64 `gorm:"column:village_id" json:"village_id"`
	VillageName                string `gorm:"column:village_name" json:"village_name"`
	ClientAddressZipCode       string `gorm:"column:client_adrs_zip_code" json:"client_adrs_zip_code"`
	ClientAddressOwnership     string `gorm:"column:client_adrs_ownership" json:"client_adrs_ownership"`
	ClientAddressLengthStay    uint64 `gorm:"column:client_adrs_length_stay" json:"client_adrs_length_stay"`
	ClientAddressAsResidential uint64 `gorm:"column:client_adrs_as_residential" json:"client_adrs_as_residential"`
}

func (ClientMrdmData) TableName() string {
	return "m_client"
}

func (p *Client) FindClientMrdmByClientID(db *gorm.DB, pid uint64) (ClientMrdmData, error) {
	var err error
	client := ClientMrdmData{}
	err = db.Debug().Model(&ClientMrdmData{}).Limit(100).
		Select(`m_client.client_id,
				m_client.cust_user_id,
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
				m_client.client_code,
				m_client.client_ktp,
				m_client.client_passport,
				m_client.client_npwp,
				m_client.client_name,
				m_client.client_sex,
				m_client.client_birth_place,
				m_client.client_birth_date,
				m_client.client_age,
				m_client.client_phone,
				m_client.client_phone_work,
				m_client.client_email,
				m_client.client_email_work,
				m_client.client_marital_status,
				m_client.client_image,
				m_client.client_desc,
				m_client.client_status,
				m_client.client_created_at,
				m_client.client_created_by,
				m_client.client_updated_at,
				m_client.client_updated_by,
				m_client.client_deleted_at,
				m_client.client_deleted_by`).
		Joins("JOIN m_client_address on m_client.client_adrs_id=m_client_address.client_adrs_id").
		Joins("JOIN m_address on m_client_address.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("m_client.client_id = ?", pid).
		Find(&client).Error
	if err != nil {
		return ClientMrdmData{}, err
	}
	return client, nil
}

// ADDITIONAL
// ===================================================================================================================================

type ResponseClientProfileImage struct {
	Status  uint64 `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func (p *Client) UpdateProfileClient(db *gorm.DB, uid uint64) (*Client, error) {
	var err error
	db = db.Debug().Model(&Client{}).Where("client_id = ?", uid).Take(&Client{}).UpdateColumns(
		map[string]interface{}{
			"client_name":           p.ClientName,
			"client_sex":            p.ClientSex,
			"client_birth_date":     p.ClientBirthDate,
			"client_birth_place":    p.ClientBirthPlace,
			"client_marital_status": p.ClientMaritalStatus,
			"client_age":            p.ClientAge,
			"client_ktp":            p.ClientKTP,
			"client_passport":       p.ClientPassport,
			"client_npwp":           p.ClientNPWP,
			"client_image":          p.ClientImage,
			"client_created_by":     p.ClientCreatedBy,
			"client_created_at":     time.Now(),
			"client_updated_by":     p.ClientUpdatedBy,
			"client_updated_at":     time.Now(),
			"client_deleted_by":     p.ClientDeletedBy,
			"client_deleted_at":     time.Now(),
		},
	)
	if db.Error != nil {
		return &Client{}, db.Error
	}
	err = db.Debug().Model(&Client{}).Where("client_id = ?", uid).Take(&p).Error
	if err != nil {
		return &Client{}, err
	}
	return p, nil
}

func (p *Client) UpdateClientImage(db *gorm.DB, pid uint64, image string) (*Client, error) {
	var err error
	db = db.Debug().Model(&Client{}).Where("client_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"client_image":      image,
			"client_updated_by": p.ClientUpdatedBy,
			"client_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&Client{}).Where("client_id = ?", pid).Error
	if err != nil {
		return &Client{}, err
	}
	return p, nil
}

func (p *Client) UpdateClientDescription(db *gorm.DB, pid uint64) (*Client, error) {
	var err error
	db = db.Debug().Model(&Client{}).Where("client_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"client_desc":       p.ClientDescription,
			"client_updated_by": p.ClientUpdatedBy,
			"client_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&Client{}).Where("client_id = ?", pid).Error
	if err != nil {
		return &Client{}, err
	}
	return p, nil
}
