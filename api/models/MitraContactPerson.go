package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type MitraContactPerson struct {
	MitraContactPersonID              *uint64    `gorm:"column:mitra_cp_id;primary_key;" json:"mitra_cp_id"`
	MitraBranchID                     uint64     `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	MitraContactPersonFirstName       string     `gorm:"column:mitra_cp_first_name;size:25" json:"mitra_cp_first_name"`
	MitraContactPersonLastName        string     `gorm:"column:mitra_cp_last_name;size:255" json:"mitra_cp_last_name"`
	MitraContactPersonJobTitle        string     `gorm:"column:mitra_cp_job_title;size:255" json:"mitra_cp_job_title"`
	MitraContactPersonMail            string     `gorm:"column:mitra_cp_mail;size:255" json:"mitra_cp_mail"`
	MitraContactPersonMailWork        string     `gorm:"column:mitra_cp_mail_work;size:255" json:"mitra_cp_mail_work"`
	MitraContactPersonPhone           string     `gorm:"column:mitra_cp_phone;size:25" json:"mitra_cp_phone"`
	MitraContactPersonPhoneWork       string     `gorm:"column:mitra_cp_phone_work;size:255" json:"mitra_cp_phone_work"`
	MitraContactPersonExtensionNumber string     `gorm:"column:mitra_cp_ext_no;size:255" json:"mitra_cp_ext_no"`
	MitraContactPersonStatus          uint64     `gorm:"column:mitra_cp_status;size:25" json:"mitra_cp_status"`
	MitraContactPersonCreatedBy       string     `gorm:"column:mitra_cp_created_by;size:125" json:"mitra_cp_created_by"`
	MitraContactPersonUpdatedBy       string     `gorm:"column:mitra_cp_updated_by;size:125" json:"mitra_cp_updated_by"`
	MitraContactPersonDeletedBy       string     `gorm:"column:mitra_cp_deleted_by;size:125" json:"mitra_cp_deleted_by"`
	MitraContactPersonCreatedAt       time.Time  `gorm:"column:mitra_cp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_cp_created_at"`
	MitraContactPersonUpdatedAt       *time.Time `gorm:"column:mitra_cp_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_cp_updated_at"`
	MitraContactPersonDeletedAt       *time.Time `gorm:"column:mitra_cp_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_cp_deleted_at"`
}

type MitraContactPersonData struct {
	MitraContactPersonID              uint64     `gorm:"column:mitra_cp_id" json:"mitra_cp_id"`
	MitraID                           uint64     `gorm:"column:mitra_id" json:"mitra_id"`
	MitraCode                         string     `gorm:"column:mitra_code;size:25" json:"mitra_code"`
	MitraName                         string     `gorm:"column:mitra_name;size:255" json:"mitra_name"`
	MitraWebsite                      string     `gorm:"column:mitra_website;size:255" json:"mitra_website"`
	MitraEmail                        string     `gorm:"column:mitra_email;size:255" json:"mitra_email"`
	MitraLogo                         string     `gorm:"column:mitra_logo;size:255" json:"mitra_logo"`
	MitraBranchID                     uint64     `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	MitraBranchName                   string     `gorm:"column:mitra_branch_name;size:255" json:"mitra_branch_name"`
	MitraBranchType                   string     `gorm:"column:mitra_branch_type;size:50" json:"mitra_branch_type"`
	MitraBranchPhone                  string     `gorm:"column:mitra_branch_phone;size:20" json:"mitra_branch_phone"`
	MitraBranchFax                    string     `gorm:"column:mitra_branch_fax;size:20" json:"mitra_branch_fax"`
	MitraBranchAddress                string     `gorm:"column:mitra_branch_address" json:"mitra_branch_address"`
	AddressID                         string     `gorm:"column:address_id" json:"address_id"`
	CountryID                         string     `gorm:"column:country_id" json:"country_id"`
	CountryName                       string     `gorm:"column:country_name;size:255" json:"country_name"`
	ProvinceID                        string     `gorm:"column:province_id" json:"province_id"`
	ProvinceName                      string     `gorm:"column:province_name;size:255" json:"province_name"`
	RegencyID                         string     `gorm:"column:regency_id" json:"regency_id"`
	RegencyName                       string     `gorm:"column:regency_name;size:255" json:"regency_name"`
	DistrictID                        string     `gorm:"column:district_id" json:"district_id"`
	DistrictName                      string     `gorm:"column:district_name;size:255" json:"district_name"`
	VillageID                         string     `gorm:"column:village_id" json:"village_id"`
	VillageName                       string     `gorm:"column:village_name;size:255" json:"village_name"`
	MitraContactPersonFirstName       string     `gorm:"column:mitra_cp_first_name;size:255" json:"mitra_cp_first_name"`
	MitraContactPersonLastName        string     `gorm:"column:mitra_cp_last_name;size:255" json:"mitra_cp_last_name"`
	MitraContactPersonJobTitle        string     `gorm:"column:mitra_cp_job_title;size:125" json:"mitra_cp_job_title"`
	MitraContactPersonMail            string     `gorm:"column:mitra_cp_mail;size:255" json:"mitra_cp_mail"`
	MitraContactPersonMailWork        string     `gorm:"column:mitra_cp_mail_work;size:255" json:"mitra_cp_mail_work"`
	MitraContactPersonPhone           string     `gorm:"column:mitra_cp_phone;size:25" json:"mitra_cp_phone"`
	MitraContactPersonPhoneWork       string     `gorm:"column:mitra_cp_phone_work;size:25" json:"mitra_cp_phone_work"`
	MitraContactPersonExtensionNumber string     `gorm:"column:mitra_cp_ext_no;size:255" json:"mitra_cp_ext_no"`
	MitraContactPersonStatus          uint64     `gorm:"column:mitra_cp_status;size:2" json:"mitra_cp_status"`
	MitraContactPersonCreatedBy       string     `gorm:"column:mitra_cp_created_by;size:125" json:"mitra_cp_created_by"`
	MitraContactPersonUpdatedBy       string     `gorm:"column:mitra_cp_updated_by;size:125" json:"mitra_cp_updated_by"`
	MitraContactPersonDeletedBy       string     `gorm:"column:mitra_cp_deleted_by;size:125" json:"mitra_cp_deleted_by"`
	MitraContactPersonCreatedAt       time.Time  `gorm:"column:mitra_cp_created_at" json:"mitra_cp_created_at"`
	MitraContactPersonUpdatedAt       *time.Time `gorm:"column:mitra_cp_updated_at" json:"mitra_cp_updated_at"`
	MitraContactPersonDeletedAt       *time.Time `gorm:"column:mitra_cp_deleted_at" json:"mitra_cp_deleted_at"`
}

type ResponseMitraContactPersons struct {
	Status  int                       `json:"status"`
	Message string                    `json:"message"`
	Data    *[]MitraContactPersonData `json:"data"`
}

type ResponseMitraContactPerson struct {
	Status  int                     `json:"status"`
	Message string                  `json:"message"`
	Data    *MitraContactPersonData `json:"data"`
}

type ResponseMitraContactPersonIU struct {
	Status  int                 `json:"status"`
	Message string              `json:"message"`
	Data    *MitraContactPerson `json:"data"`
}

type ResponseMitraContactPersonCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraContactPerson) TableName() string {
	return "m_mitra_contact_person"
}

func (MitraContactPersonData) TableName() string {
	return "m_mitra_contact_person"
}

func (p *MitraContactPerson) Prepare() {
	p.MitraContactPersonID = nil
	p.MitraBranchID = p.MitraBranchID
	p.MitraContactPersonFirstName = html.EscapeString(strings.TrimSpace(p.MitraContactPersonFirstName))
	p.MitraContactPersonLastName = html.EscapeString(strings.TrimSpace(p.MitraContactPersonLastName))
	p.MitraContactPersonJobTitle = html.EscapeString(strings.TrimSpace(p.MitraContactPersonJobTitle))
	p.MitraContactPersonMail = html.EscapeString(strings.TrimSpace(p.MitraContactPersonMail))
	p.MitraContactPersonMailWork = html.EscapeString(strings.TrimSpace(p.MitraContactPersonMailWork))
	p.MitraContactPersonPhone = html.EscapeString(strings.TrimSpace(p.MitraContactPersonPhone))
	p.MitraContactPersonPhoneWork = html.EscapeString(strings.TrimSpace(p.MitraContactPersonPhoneWork))
	p.MitraContactPersonExtensionNumber = html.EscapeString(strings.TrimSpace(p.MitraContactPersonExtensionNumber))
	p.MitraContactPersonStatus = p.MitraContactPersonStatus
	p.MitraContactPersonCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraContactPersonCreatedBy))
	p.MitraContactPersonUpdatedBy = html.EscapeString(strings.TrimSpace(p.MitraContactPersonUpdatedBy))
	p.MitraContactPersonDeletedBy = html.EscapeString(strings.TrimSpace(p.MitraContactPersonDeletedBy))
	p.MitraContactPersonCreatedAt = time.Now()
	p.MitraContactPersonUpdatedAt = p.MitraContactPersonUpdatedAt
	p.MitraContactPersonDeletedAt = p.MitraContactPersonDeletedAt
}

func (p *MitraContactPerson) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.MitraBranchID == 0 {
			return errors.New("Required MitraContactPerson Code")
		}
		if p.MitraContactPersonFirstName == "" {
			return errors.New("Required MitraContactPerson Name")
		}
		if p.MitraContactPersonLastName == "" {
			return errors.New("Required MitraContactPerson Name")
		}
		if p.MitraContactPersonJobTitle == "" {
			return errors.New("Required MitraContactPerson Name")
		}
		if p.MitraContactPersonMail == "" {
			return errors.New("Required MitraContactPerson Name")
		}
		if p.MitraContactPersonMailWork == "" {
			return errors.New("Required MitraContactPerson Name")
		}
		if p.MitraContactPersonPhone == "" {
			return errors.New("Required MitraContactPerson Name")
		}
		if p.MitraContactPersonPhoneWork == "" {
			return errors.New("Required MitraContactPerson Name")
		}
		return nil
	}
}

func (p *MitraContactPerson) SaveMitraContactPerson(db *gorm.DB) (*MitraContactPerson, error) {
	var err error
	err = db.Debug().Model(&MitraContactPerson{}).Create(&p).Error
	if err != nil {
		return &MitraContactPerson{}, err
	}
	return p, nil
}

func (p *MitraContactPerson) FindAllMitraContactPerson(db *gorm.DB) (*[]MitraContactPersonData, error) {
	var err error
	mitracontactperson := []MitraContactPersonData{}
	err = db.Debug().Model(&MitraContactPersonData{}).Limit(100).
		Select(`m_mitra_contact_person.mitra_cp_id,
				m_mitra.mitra_id,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra.mitra_website,
				m_mitra.mitra_email,
				m_mitra.mitra_logo,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_branch.mitra_branch_type,
				m_mitra_branch.mitra_branch_phone,
				m_mitra_branch.mitra_branch_fax,
				m_mitra_branch.mitra_branch_address,
				m_mitra_branch.address_id,
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
				m_mitra_contact_person.mitra_cp_first_name,
				m_mitra_contact_person.mitra_cp_last_name,
				m_mitra_contact_person.mitra_cp_job_title,
				m_mitra_contact_person.mitra_cp_mail,
				m_mitra_contact_person.mitra_cp_mail_work,
				m_mitra_contact_person.mitra_cp_phone,
				m_mitra_contact_person.mitra_cp_phone_work,
				m_mitra_contact_person.mitra_cp_ext_no,
				m_mitra_contact_person.mitra_cp_status,
				m_mitra_contact_person.mitra_cp_created_by,
				m_mitra_contact_person.mitra_cp_updated_by,
				m_mitra_contact_person.mitra_cp_deleted_by,
				m_mitra_contact_person.mitra_cp_created_at,
				m_mitra_contact_person.mitra_cp_updated_at,
				m_mitra_contact_person.mitra_cp_deleted_at
				`).
		Joins("JOIN m_mitra_branch on m_mitra_contact_person.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra on m_mitra_branch.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_address on m_mitra_branch.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Find(&mitracontactperson).Error
	if err != nil {
		return &[]MitraContactPersonData{}, err
	}
	return &mitracontactperson, nil
}

func (p *MitraContactPerson) FindAllMitraContactPersonStatus(db *gorm.DB, status uint64) (*[]MitraContactPersonData, error) {
	var err error
	mitracontactperson := []MitraContactPersonData{}
	err = db.Debug().Model(&MitraContactPersonData{}).Limit(100).
		Select(`m_mitra_contact_person.mitra_cp_id,
				m_mitra.mitra_id,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra.mitra_website,
				m_mitra.mitra_email,
				m_mitra.mitra_logo,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_branch.mitra_branch_type,
				m_mitra_branch.mitra_branch_phone,
				m_mitra_branch.mitra_branch_fax,
				m_mitra_branch.mitra_branch_address,
				m_mitra_branch.address_id,
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
				m_mitra_contact_person.mitra_cp_first_name,
				m_mitra_contact_person.mitra_cp_last_name,
				m_mitra_contact_person.mitra_cp_job_title,
				m_mitra_contact_person.mitra_cp_mail,
				m_mitra_contact_person.mitra_cp_mail_work,
				m_mitra_contact_person.mitra_cp_phone,
				m_mitra_contact_person.mitra_cp_phone_work,
				m_mitra_contact_person.mitra_cp_ext_no,
				m_mitra_contact_person.mitra_cp_status,
				m_mitra_contact_person.mitra_cp_created_by,
				m_mitra_contact_person.mitra_cp_updated_by,
				m_mitra_contact_person.mitra_cp_deleted_by,
				m_mitra_contact_person.mitra_cp_created_at,
				m_mitra_contact_person.mitra_cp_updated_at,
				m_mitra_contact_person.mitra_cp_deleted_at
				`).
		Joins("JOIN m_mitra_branch on m_mitra_contact_person.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra on m_mitra_branch.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_address on m_mitra_branch.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("mitra_cp_status = ?", status).
		Find(&mitracontactperson).Error
	if err != nil {
		return &[]MitraContactPersonData{}, err
	}
	return &mitracontactperson, nil
}

func (p *MitraContactPerson) FindAllMitraContactPersonByMitraBranch(db *gorm.DB, mitrabranch uint64) (*[]MitraContactPersonData, error) {
	var err error
	mitracontactperson := []MitraContactPersonData{}
	err = db.Debug().Model(&MitraContactPersonData{}).Limit(100).
		Select(`m_mitra_contact_person.mitra_cp_id,
				m_mitra.mitra_id,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra.mitra_website,
				m_mitra.mitra_email,
				m_mitra.mitra_logo,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_branch.mitra_branch_type,
				m_mitra_branch.mitra_branch_phone,
				m_mitra_branch.mitra_branch_fax,
				m_mitra_branch.mitra_branch_address,
				m_mitra_branch.address_id,
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
				m_mitra_contact_person.mitra_cp_first_name,
				m_mitra_contact_person.mitra_cp_last_name,
				m_mitra_contact_person.mitra_cp_job_title,
				m_mitra_contact_person.mitra_cp_mail,
				m_mitra_contact_person.mitra_cp_mail_work,
				m_mitra_contact_person.mitra_cp_phone,
				m_mitra_contact_person.mitra_cp_phone_work,
				m_mitra_contact_person.mitra_cp_ext_no,
				m_mitra_contact_person.mitra_cp_status,
				m_mitra_contact_person.mitra_cp_created_by,
				m_mitra_contact_person.mitra_cp_updated_by,
				m_mitra_contact_person.mitra_cp_deleted_by,
				m_mitra_contact_person.mitra_cp_created_at,
				m_mitra_contact_person.mitra_cp_updated_at,
				m_mitra_contact_person.mitra_cp_deleted_at
				`).
		Joins("JOIN m_mitra_branch on m_mitra_contact_person.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra on m_mitra_branch.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_address on m_mitra_branch.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("m_mitra_contact_person.mitra_branch_id = ?", mitrabranch).
		Find(&mitracontactperson).Error
	if err != nil {
		return &[]MitraContactPersonData{}, err
	}
	return &mitracontactperson, nil
}

func (p *MitraContactPerson) FindMitraContactPersonDataByID(db *gorm.DB, pid uint64) (*MitraContactPerson, error) {
	var err error
	err = db.Debug().Model(&MitraContactPerson{}).Where("mitra_cp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraContactPerson{}, err
	}
	return p, nil
}

func (p *MitraContactPerson) FindMitraContactPersonByID(db *gorm.DB, pid uint64) (*MitraContactPersonData, error) {
	var err error
	mitracontactperson := MitraContactPersonData{}
	err = db.Debug().Model(&MitraContactPersonData{}).
		Select(`m_mitra_contact_person.mitra_cp_id,
				m_mitra.mitra_id,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra.mitra_website,
				m_mitra.mitra_email,
				m_mitra.mitra_logo,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_branch.mitra_branch_type,
				m_mitra_branch.mitra_branch_phone,
				m_mitra_branch.mitra_branch_fax,
				m_mitra_branch.mitra_branch_address,
				m_mitra_branch.address_id,
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
				m_mitra_contact_person.mitra_cp_first_name,
				m_mitra_contact_person.mitra_cp_last_name,
				m_mitra_contact_person.mitra_cp_job_title,
				m_mitra_contact_person.mitra_cp_mail,
				m_mitra_contact_person.mitra_cp_mail_work,
				m_mitra_contact_person.mitra_cp_phone,
				m_mitra_contact_person.mitra_cp_phone_work,
				m_mitra_contact_person.mitra_cp_ext_no,
				m_mitra_contact_person.mitra_cp_status,
				m_mitra_contact_person.mitra_cp_created_by,
				m_mitra_contact_person.mitra_cp_updated_by,
				m_mitra_contact_person.mitra_cp_deleted_by,
				m_mitra_contact_person.mitra_cp_created_at,
				m_mitra_contact_person.mitra_cp_updated_at,
				m_mitra_contact_person.mitra_cp_deleted_at
			`).
		Joins("JOIN m_mitra_branch on m_mitra_contact_person.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra on m_mitra_branch.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_address on m_mitra_branch.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("mitra_cp_id = ?", pid).
		Take(&mitracontactperson).Error
	if err != nil {
		return &MitraContactPersonData{}, err
	}
	return &mitracontactperson, nil
}

func (p *MitraContactPerson) FindMitraContactPersonStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraContactPersonData, error) {
	var err error
	mitracontactperson := MitraContactPersonData{}
	err = db.Debug().Model(&MitraContactPersonData{}).
		Select(`m_mitra_contact_person.mitra_cp_id,
				m_mitra.mitra_id,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra.mitra_website,
				m_mitra.mitra_email,
				m_mitra.mitra_logo,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_branch.mitra_branch_type,
				m_mitra_branch.mitra_branch_phone,
				m_mitra_branch.mitra_branch_fax,
				m_mitra_branch.mitra_branch_address,
				m_mitra_branch.address_id,
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
				m_mitra_contact_person.mitra_cp_first_name,
				m_mitra_contact_person.mitra_cp_last_name,
				m_mitra_contact_person.mitra_cp_job_title,
				m_mitra_contact_person.mitra_cp_mail,
				m_mitra_contact_person.mitra_cp_mail_work,
				m_mitra_contact_person.mitra_cp_phone,
				m_mitra_contact_person.mitra_cp_phone_work,
				m_mitra_contact_person.mitra_cp_ext_no,
				m_mitra_contact_person.mitra_cp_status,
				m_mitra_contact_person.mitra_cp_created_by,
				m_mitra_contact_person.mitra_cp_updated_by,
				m_mitra_contact_person.mitra_cp_deleted_by,
				m_mitra_contact_person.mitra_cp_created_at,
				m_mitra_contact_person.mitra_cp_updated_at,
				m_mitra_contact_person.mitra_cp_deleted_at
			`).
		Joins("JOIN m_mitra_branch on m_mitra_contact_person.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra on m_mitra_branch.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_address on m_mitra_branch.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("mitra_cp_id = ? AND mitra_cp_status = ?", pid, status).
		Take(&mitracontactperson).Error
	if err != nil {
		return &MitraContactPersonData{}, err
	}
	return &mitracontactperson, nil
}

func (p *MitraContactPerson) FindMitraContactPersonByMitraBranchByID(db *gorm.DB, pid uint64, mitrabranch uint64) (*MitraContactPersonData, error) {
	var err error
	mitracontactperson := MitraContactPersonData{}
	err = db.Debug().Model(&MitraContactPersonData{}).
		Select(`m_mitra_contact_person.mitra_cp_id,
				m_mitra.mitra_id,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra.mitra_website,
				m_mitra.mitra_email,
				m_mitra.mitra_logo,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_branch.mitra_branch_type,
				m_mitra_branch.mitra_branch_phone,
				m_mitra_branch.mitra_branch_fax,
				m_mitra_branch.mitra_branch_address,
				m_mitra_branch.address_id,
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
				m_mitra_contact_person.mitra_cp_first_name,
				m_mitra_contact_person.mitra_cp_last_name,
				m_mitra_contact_person.mitra_cp_job_title,
				m_mitra_contact_person.mitra_cp_mail,
				m_mitra_contact_person.mitra_cp_mail_work,
				m_mitra_contact_person.mitra_cp_phone,
				m_mitra_contact_person.mitra_cp_phone_work,
				m_mitra_contact_person.mitra_cp_ext_no,
				m_mitra_contact_person.mitra_cp_status,
				m_mitra_contact_person.mitra_cp_created_by,
				m_mitra_contact_person.mitra_cp_updated_by,
				m_mitra_contact_person.mitra_cp_deleted_by,
				m_mitra_contact_person.mitra_cp_created_at,
				m_mitra_contact_person.mitra_cp_updated_at,
				m_mitra_contact_person.mitra_cp_deleted_at
			`).
		Joins("JOIN m_mitra_branch on m_mitra_contact_person.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra on m_mitra_branch.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_address on m_mitra_branch.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("mitra_cp_id = ? AND mitra_branch_id = ?", pid, mitrabranch).
		Take(&mitracontactperson).Error
	if err != nil {
		return &MitraContactPersonData{}, err
	}
	return &mitracontactperson, nil
}

func (p *MitraContactPerson) UpdateMitraContactPerson(db *gorm.DB, pid uint64) (*MitraContactPerson, error) {
	var err error
	db = db.Debug().Model(&MitraContactPerson{}).Where("mitra_cp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_branch_id":     p.MitraBranchID,
			"mitra_cp_first_name": p.MitraContactPersonFirstName,
			"mitra_cp_last_name":  p.MitraContactPersonLastName,
			"mitra_cp_job_title":  p.MitraContactPersonJobTitle,
			"mitra_cp_mail":       p.MitraContactPersonMail,
			"mitra_cp_mail_work":  p.MitraContactPersonMailWork,
			"mitra_cp_phone":      p.MitraContactPersonPhone,
			"mitra_cp_phone_work": p.MitraContactPersonPhoneWork,
			"mitra_cp_ext_no":     p.MitraContactPersonExtensionNumber,
			"mitra_cp_status":     p.MitraContactPersonStatus,
			"mitra_cp_updated_by": p.MitraContactPersonUpdatedBy,
			"mitra_cp_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraContactPerson{}).Where("mitra_cp_id = ?", pid).Error
	if err != nil {
		return &MitraContactPerson{}, err
	}
	return p, nil
}

func (p *MitraContactPerson) UpdateMitraContactPersonStatus(db *gorm.DB, pid uint64) (*MitraContactPerson, error) {
	var err error
	db = db.Debug().Model(&MitraContactPerson{}).Where("mitra_cp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_cp_status":     p.MitraContactPersonStatus,
			"mitra_cp_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraContactPerson{}).Where("mitra_cp_id = ?", pid).Error
	if err != nil {
		return &MitraContactPerson{}, err
	}
	return p, nil
}

func (p *MitraContactPerson) UpdateMitraContactPersonStatusOnly(db *gorm.DB, pid uint64) (*MitraContactPerson, error) {
	var err error
	db = db.Debug().Model(&MitraContactPerson{}).Where("mitra_cp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_cp_status": p.MitraContactPersonStatus,
		},
	)
	err = db.Debug().Model(&MitraContactPerson{}).Where("mitra_cp_id = ?", pid).Error
	if err != nil {
		return &MitraContactPerson{}, err
	}
	return p, nil
}

func (p *MitraContactPerson) UpdateMitraContactPersonStatusDelete(db *gorm.DB, pid uint64) (*MitraContactPerson, error) {
	var err error
	db = db.Debug().Model(&MitraContactPerson{}).Where("mitra_cp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_cp_status":     3,
			"mitra_cp_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraContactPerson{}).Where("mitra_cp_id = ?", pid).Error
	if err != nil {
		return &MitraContactPerson{}, err
	}
	return p, nil
}

func (p *MitraContactPerson) DeleteMitraContactPerson(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraContactPerson{}).Where("mitra_cp_id = ?", pid).Take(&MitraContactPerson{}).Delete(&MitraContactPerson{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraContactPerson not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

func (p *MitraContactPerson) FindMitraContactPersonByMitraBranchID(db *gorm.DB, mitra uint64) ([]MitraContactPersonData, error) {
	var err error
	mitracoverageprovince := []MitraContactPersonData{}
	err = db.Debug().Model(&MitraContactPersonData{}).Limit(100).
		Select(`m_mitra_contact_person.mitra_cp_id,
				m_mitra.mitra_id,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra.mitra_website,
				m_mitra.mitra_email,
				m_mitra.mitra_logo,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_branch.mitra_branch_type,
				m_mitra_branch.mitra_branch_phone,
				m_mitra_branch.mitra_branch_fax,
				m_mitra_branch.mitra_branch_address,
				m_mitra_branch.address_id,
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
				m_mitra_contact_person.mitra_cp_first_name,
				m_mitra_contact_person.mitra_cp_last_name,
				m_mitra_contact_person.mitra_cp_job_title,
				m_mitra_contact_person.mitra_cp_mail,
				m_mitra_contact_person.mitra_cp_mail_work,
				m_mitra_contact_person.mitra_cp_phone,
				m_mitra_contact_person.mitra_cp_phone_work,
				m_mitra_contact_person.mitra_cp_ext_no,
				m_mitra_contact_person.mitra_cp_status,
				m_mitra_contact_person.mitra_cp_created_by,
				m_mitra_contact_person.mitra_cp_updated_by,
				m_mitra_contact_person.mitra_cp_deleted_by,
				m_mitra_contact_person.mitra_cp_created_at,
				m_mitra_contact_person.mitra_cp_updated_at,
				m_mitra_contact_person.mitra_cp_deleted_at
			`).
		Joins("JOIN m_mitra_branch on m_mitra_contact_person.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra on m_mitra_branch.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_address on m_mitra_branch.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("m_mitra_contact_person.mitra_branch_id = ?", mitra).
		Find(&mitracoverageprovince).Error
	if err != nil {
		return []MitraContactPersonData{}, err
	}
	return mitracoverageprovince, nil
}

func (p *MitraContactPerson) FindMitraContactPersonStatusByMitraBranchID(db *gorm.DB, mitra uint64, status uint64) ([]MitraContactPersonData, error) {
	var err error
	mitracoverageprovince := []MitraContactPersonData{}
	err = db.Debug().Model(&MitraContactPersonData{}).Limit(100).
		Select(`m_mitra_contact_person.mitra_cp_id,
				m_mitra.mitra_id,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra.mitra_website,
				m_mitra.mitra_email,
				m_mitra.mitra_logo,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_branch.mitra_branch_type,
				m_mitra_branch.mitra_branch_phone,
				m_mitra_branch.mitra_branch_fax,
				m_mitra_branch.mitra_branch_address,
				m_mitra_branch.address_id,
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
				m_mitra_contact_person.mitra_cp_first_name,
				m_mitra_contact_person.mitra_cp_last_name,
				m_mitra_contact_person.mitra_cp_job_title,
				m_mitra_contact_person.mitra_cp_mail,
				m_mitra_contact_person.mitra_cp_mail_work,
				m_mitra_contact_person.mitra_cp_phone,
				m_mitra_contact_person.mitra_cp_phone_work,
				m_mitra_contact_person.mitra_cp_ext_no,
				m_mitra_contact_person.mitra_cp_status,
				m_mitra_contact_person.mitra_cp_created_by,
				m_mitra_contact_person.mitra_cp_updated_by,
				m_mitra_contact_person.mitra_cp_deleted_by,
				m_mitra_contact_person.mitra_cp_created_at,
				m_mitra_contact_person.mitra_cp_updated_at,
				m_mitra_contact_person.mitra_cp_deleted_at
			`).
		Joins("JOIN m_mitra_branch on m_mitra_contact_person.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_mitra on m_mitra_branch.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_address on m_mitra_branch.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Where("m_mitra_contact_person.mitra_branch_id = ? AND m_mitra_contact_person.mitra_cp_status = ?", mitra, status).
		Find(&mitracoverageprovince).Error
	if err != nil {
		return []MitraContactPersonData{}, err
	}
	return mitracoverageprovince, nil
}
