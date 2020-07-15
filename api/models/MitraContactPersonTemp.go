package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type MitraContactPersonTemp struct {
	MitraContactPersonTempID              *uint64   `gorm:"column:mitra_cp_temp_id;primary_key;" json:"mitra_cp_temp_id"`
	MitraContactPersonID                  uint64    `gorm:"column:mitra_cp_id" json:"mitra_cp_id"`
	MitraBranchTempID                     uint64    `gorm:"column:mitra_branch_temp_id" json:"mitra_branch_temp_id"`
	MitraContactPersonTempFirstName       string    `gorm:"column:mitra_cp_temp_first_name;size:25" json:"mitra_cp_temp_first_name"`
	MitraContactPersonTempLastName        string    `gorm:"column:mitra_cp_temp_last_name;size:255" json:"mitra_cp_temp_last_name"`
	MitraContactPersonTempJobTitle        string    `gorm:"column:mitra_cp_temp_job_title;size:255" json:"mitra_cp_temp_job_title"`
	MitraContactPersonTempMail            string    `gorm:"column:mitra_cp_temp_mail;size:255" json:"mitra_cp_temp_mail"`
	MitraContactPersonTempMailWork        string    `gorm:"column:mitra_cp_temp_mail_work;size:255" json:"mitra_cp_temp_mail_work"`
	MitraContactPersonTempPhone           string    `gorm:"column:mitra_cp_temp_phone;size:255" json:"mitra_cp_temp_phone"`
	MitraContactPersonTempPhoneWork       string    `gorm:"column:mitra_cp_temp_phone_work;size:255" json:"mitra_cp_temp_phone_work"`
	MitraContactPersonTempExtensionNumber string    `gorm:"column:mitra_cp_temp_ext_no;size:255" json:"mitra_cp_temp_ext_no"`
	MitraContactPersonTempReason          string    `gorm:"column:mitra_cp_temp_reason" json:"mitra_cp_temp_reason"`
	MitraContactPersonTempActionBefore    uint64    `gorm:"column:mitra_cp_temp_action_before" json:"mitra_cp_temp_action_before"`
	MitraContactPersonTempStatus          uint64    `gorm:"column:mitra_cp_temp_status;size:25" json:"mitra_cp_temp_status"`
	MitraContactPersonTempCreatedBy       string    `gorm:"column:mitra_cp_temp_created_by;size:125" json:"mitra_cp_temp_created_by"`
	MitraContactPersonTempCreatedAt       time.Time `gorm:"column:mitra_cp_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_cp_temp_created_at"`
}

type MitraContactPersonTempPend struct {
	MitraContactPersonTempID              *uint64   `gorm:"column:mitra_cp_temp_id;primary_key;" json:"mitra_cp_temp_id"`
	MitraContactPersonID                  *uint64   `gorm:"column:mitra_cp_id" json:"mitra_cp_id"`
	MitraBranchTempID                     uint64    `gorm:"column:mitra_branch_temp_id" json:"mitra_branch_temp_id"`
	MitraContactPersonTempFirstName       string    `gorm:"column:mitra_cp_temp_first_name;size:25" json:"mitra_cp_temp_first_name"`
	MitraContactPersonTempLastName        string    `gorm:"column:mitra_cp_temp_last_name;size:255" json:"mitra_cp_temp_last_name"`
	MitraContactPersonTempJobTitle        string    `gorm:"column:mitra_cp_temp_job_title;size:255" json:"mitra_cp_temp_job_title"`
	MitraContactPersonTempMail            string    `gorm:"column:mitra_cp_temp_mail;size:255" json:"mitra_cp_temp_mail"`
	MitraContactPersonTempMailWork        string    `gorm:"column:mitra_cp_temp_mail_work;size:255" json:"mitra_cp_temp_mail_work"`
	MitraContactPersonTempPhone           string    `gorm:"column:mitra_cp_temp_phone;size:255" json:"mitra_cp_temp_phone"`
	MitraContactPersonTempPhoneWork       string    `gorm:"column:mitra_cp_temp_phone_work;size:255" json:"mitra_cp_temp_phone_work"`
	MitraContactPersonTempExtensionNumber string    `gorm:"column:mitra_cp_temp_ext_no;size:255" json:"mitra_cp_temp_ext_no"`
	MitraContactPersonTempReason          string    `gorm:"column:mitra_cp_temp_reason" json:"mitra_cp_temp_reason"`
	MitraContactPersonTempActionBefore    uint64    `gorm:"column:mitra_cp_temp_action_before" json:"mitra_cp_temp_action_before"`
	MitraContactPersonTempStatus          uint64    `gorm:"column:mitra_cp_temp_status;size:25" json:"mitra_cp_temp_status"`
	MitraContactPersonTempCreatedBy       string    `gorm:"column:mitra_cp_temp_created_by;size:125" json:"mitra_cp_temp_created_by"`
	MitraContactPersonTempCreatedAt       time.Time `gorm:"column:mitra_cp_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_cp_temp_created_at"`
}

type MitraContactPersonTempPendData struct {
	MitraContactPersonTempID              uint64    `gorm:"column:mitra_cp_temp_id;primary_key;" json:"mitra_cp_temp_id"`
	MitraContactPersonID                  uint64    `gorm:"column:mitra_cp_id" json:"mitra_cp_id"`
	MitraBranchTempID                     uint64    `gorm:"column:mitra_branch_temp_id" json:"mitra_branch_temp_id"`
	MitraBranchTempName                   string    `gorm:"column:mitra_branch_temp_name" json:"mitra_branch_temp_name"`
	MitraContactPersonTempFirstName       string    `gorm:"column:mitra_cp_temp_first_name;size:25" json:"mitra_cp_temp_first_name"`
	MitraContactPersonTempLastName        string    `gorm:"column:mitra_cp_temp_last_name;size:255" json:"mitra_cp_temp_last_name"`
	MitraContactPersonTempJobTitle        string    `gorm:"column:mitra_cp_temp_job_title;size:255" json:"mitra_cp_temp_job_title"`
	MitraContactPersonTempMail            string    `gorm:"column:mitra_cp_temp_mail;size:255" json:"mitra_cp_temp_mail"`
	MitraContactPersonTempMailWork        string    `gorm:"column:mitra_cp_temp_mail_work;size:255" json:"mitra_cp_temp_mail_work"`
	MitraContactPersonTempPhone           string    `gorm:"column:mitra_cp_temp_phone;size:255" json:"mitra_cp_temp_phone"`
	MitraContactPersonTempPhoneWork       string    `gorm:"column:mitra_cp_temp_phone_work;size:255" json:"mitra_cp_temp_phone_work"`
	MitraContactPersonTempExtensionNumber string    `gorm:"column:mitra_cp_temp_ext_no;size:255" json:"mitra_cp_temp_ext_no"`
	MitraContactPersonTempReason          string    `gorm:"column:mitra_cp_temp_reason" json:"mitra_cp_temp_reason"`
	MitraContactPersonTempActionBefore    uint64    `gorm:"column:mitra_cp_temp_action_before" json:"mitra_cp_temp_action_before"`
	MitraContactPersonTempStatus          uint64    `gorm:"column:mitra_cp_temp_status;size:25" json:"mitra_cp_temp_status"`
	MitraContactPersonTempCreatedBy       string    `gorm:"column:mitra_cp_temp_created_by;size:125" json:"mitra_cp_temp_created_by"`
	MitraContactPersonTempCreatedAt       time.Time `gorm:"column:mitra_cp_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_cp_temp_created_at"`
}

type MitraContactPersonTempData struct {
	MitraContactPersonTempID              uint64    `gorm:"column:mitra_cp_temp_id" json:"mitra_cp_temp_id"`
	MitraContactPersonID                  uint64    `gorm:"column:mitra_cp_id" json:"mitra_cp_id"`
	MitraID                               uint64    `gorm:"column:mitra_id" json:"mitra_id"`
	MitraCode                             string    `gorm:"column:mitra_code;size:25" json:"mitra_code"`
	MitraName                             string    `gorm:"column:mitra_name;size:255" json:"mitra_name"`
	MitraWebsite                          string    `gorm:"column:mitra_website;size:255" json:"mitra_website"`
	MitraEmail                            string    `gorm:"column:mitra_email;size:255" json:"mitra_email"`
	MitraLogo                             string    `gorm:"column:mitra_logo;size:255" json:"mitra_logo"`
	MitraBranchID                         uint64    `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	MitraBranchName                       string    `gorm:"column:mitra_branch_name;size:255" json:"mitra_branch_name"`
	MitraBranchType                       string    `gorm:"column:mitra_branch_type;size:50" json:"mitra_branch_type"`
	MitraBranchPhone                      string    `gorm:"column:mitra_branch_phone;size:20" json:"mitra_branch_phone"`
	MitraBranchFax                        string    `gorm:"column:mitra_branch_fax;size:20" json:"mitra_branch_fax"`
	MitraBranchAddress                    string    `gorm:"column:mitra_branch_address" json:"mitra_branch_address"`
	AddressID                             string    `gorm:"column:address_id" json:"address_id"`
	CountryID                             string    `gorm:"column:country_id" json:"country_id"`
	CountryName                           string    `gorm:"column:country_name;size:255" json:"country_name"`
	ProvinceID                            string    `gorm:"column:province_id" json:"province_id"`
	ProvinceName                          string    `gorm:"column:province_name;size:255" json:"province_name"`
	RegencyID                             string    `gorm:"column:regency_id" json:"regency_id"`
	RegencyName                           string    `gorm:"column:regency_name;size:255" json:"regency_name"`
	DistrictID                            string    `gorm:"column:district_id" json:"district_id"`
	DistrictName                          string    `gorm:"column:district_name;size:255" json:"district_name"`
	VillageID                             string    `gorm:"column:village_id" json:"village_id"`
	VillageName                           string    `gorm:"column:village_name;size:255" json:"village_name"`
	MitraContactPersonFirstName           string    `gorm:"column:mitra_cp_first_name;size:255" json:"mitra_cp_first_name"`
	MitraContactPersonLastName            string    `gorm:"column:mitra_cp_last_name;size:255" json:"mitra_cp_last_name"`
	MitraContactPersonJobTitle            string    `gorm:"column:mitra_cp_job_title;size:255" json:"mitra_cp_job_title"`
	MitraContactPersonMail                string    `gorm:"column:mitra_cp_mail;size:255" json:"mitra_cp_mail"`
	MitraContactPersonMailWork            string    `gorm:"column:mitra_cp_mail_work;size:255" json:"mitra_cp_mail_work"`
	MitraContactPersonPhone               string    `gorm:"column:mitra_cp_phone;size:25" json:"mitra_cp_phone"`
	MitraContactPersonPhoneWork           string    `gorm:"column:mitra_cp_phone_work;size:25" json:"mitra_cp_phone_work"`
	MitraContactPersonExtensionNumber     string    `gorm:"column:mitra_cp_ext_no;size:255" json:"mitra_cp_ext_no"`
	MitraContactPersonStatus              uint64    `gorm:"column:mitra_cp_status;size:2" json:"mitra_cp_status"`
	MitraContactPersonCreatedBy           time.Time `gorm:"column:mitra_cp_created_by" json:"mitra_cp_created_by"`
	MitraContactPersonCreatedAt           time.Time `gorm:"column:mitra_cp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_cp_created_at"`
	MitraBranchTempID                     uint64    `gorm:"column:mitra_branch_temp_id" json:"mitra_branch_temp_id"`
	MitraContactPersonTempFirstName       string    `gorm:"column:mitra_cp_temp_first_name;size:255" json:"mitra_cp_temp_first_name"`
	MitraContactPersonTempLastName        string    `gorm:"column:mitra_cp_temp_last_name;size:255" json:"mitra_cp_temp_last_name"`
	MitraContactPersonTempJobTitle        string    `gorm:"column:mitra_cp_temp_job_title;size:255" json:"mitra_cp_temp_job_title"`
	MitraContactPersonTempMail            string    `gorm:"column:mitra_cp_temp_mail;size:255" json:"mitra_cp_temp_mail"`
	MitraContactPersonTempMailWork        string    `gorm:"column:mitra_cp_temp_mail_work;size:255" json:"mitra_cp_temp_mail_work"`
	MitraContactPersonTempPhone           string    `gorm:"column:mitra_cp_temp_phone;size:25" json:"mitra_cp_temp_phone"`
	MitraContactPersonTempPhoneWork       string    `gorm:"column:mitra_cp_temp_phone_work;size:25" json:"mitra_cp_temp_phone_work"`
	MitraContactPersonTempExtensionNumber string    `gorm:"column:mitra_cp_temp_ext_no;size:255" json:"mitra_cp_temp_ext_no"`
	MitraContactPersonTempReason          string    `gorm:"column:mitra_cp_temp_reason" json:"mitra_cp_temp_reason"`
	MitraContactPersonTempActionBefore    uint64    `gorm:"column:mitra_cp_temp_action_before" json:"mitra_cp_temp_action_before"`
	MitraContactPersonTempStatus          uint64    `gorm:"column:mitra_cp_temp_status;size:2" json:"mitra_cp_temp_status"`
	MitraContactPersonTempCreatedBy       string    `gorm:"column:mitra_cp_temp_created_by;size:125" json:"mitra_cp_temp_created_by"`
	MitraContactPersonTempCreatedAt       time.Time `gorm:"column:mitra_cp_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_cp_temp_created_at"`
}

type ResponseMitraContactPersonTemps struct {
	Status  int                           `json:"status"`
	Message string                        `json:"message"`
	Data    *[]MitraContactPersonTempData `json:"data"`
}

type ResponseMitraContactPersonTempsPend struct {
	Status  int                           `json:"status"`
	Message string                        `json:"message"`
	Data    *[]MitraContactPersonTempPend `json:"data"`
}

type ResponseMitraContactPersonTemp struct {
	Status  int                         `json:"status"`
	Message string                      `json:"message"`
	Data    *MitraContactPersonTempData `json:"data"`
}

type ResponseMitraContactPersonTempPend struct {
	Status  int                         `json:"status"`
	Message string                      `json:"message"`
	Data    *MitraContactPersonTempPend `json:"data"`
}

type ResponseMitraContactPersonTempIU struct {
	Status  int                     `json:"status"`
	Message string                  `json:"message"`
	Data    *MitraContactPersonTemp `json:"data"`
}

type ResponseMitraContactPersonTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraContactPersonTemp) TableName() string {
	return "m_mitra_contact_person_temp"
}

func (MitraContactPersonTempPend) TableName() string {
	return "m_mitra_contact_person_temp"
}

func (MitraContactPersonTempPendData) TableName() string {
	return "m_mitra_contact_person_temp"
}

func (MitraContactPersonTempData) TableName() string {
	return "m_mitra_contact_person_temp"
}

func (p *MitraContactPersonTemp) Prepare() {
	p.MitraContactPersonTempID = nil
	p.MitraContactPersonID = p.MitraContactPersonID
	p.MitraBranchTempID = p.MitraBranchTempID
	p.MitraContactPersonTempFirstName = html.EscapeString(strings.TrimSpace(p.MitraContactPersonTempFirstName))
	p.MitraContactPersonTempLastName = html.EscapeString(strings.TrimSpace(p.MitraContactPersonTempLastName))
	p.MitraContactPersonTempJobTitle = html.EscapeString(strings.TrimSpace(p.MitraContactPersonTempJobTitle))
	p.MitraContactPersonTempMail = html.EscapeString(strings.TrimSpace(p.MitraContactPersonTempMail))
	p.MitraContactPersonTempMailWork = html.EscapeString(strings.TrimSpace(p.MitraContactPersonTempMailWork))
	p.MitraContactPersonTempPhone = html.EscapeString(strings.TrimSpace(p.MitraContactPersonTempPhone))
	p.MitraContactPersonTempPhoneWork = html.EscapeString(strings.TrimSpace(p.MitraContactPersonTempPhoneWork))
	p.MitraContactPersonTempExtensionNumber = html.EscapeString(strings.TrimSpace(p.MitraContactPersonTempExtensionNumber))
	p.MitraContactPersonTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraContactPersonTempCreatedBy))
	p.MitraContactPersonTempReason = html.EscapeString(strings.TrimSpace(p.MitraContactPersonTempReason))
	p.MitraContactPersonTempStatus = p.MitraContactPersonTempStatus
	p.MitraContactPersonTempCreatedAt = time.Now()
}

func (p *MitraContactPersonTempPend) Prepare() {
	p.MitraContactPersonTempID = nil
	p.MitraContactPersonID = nil
	p.MitraBranchTempID = p.MitraBranchTempID
	p.MitraContactPersonTempFirstName = html.EscapeString(strings.TrimSpace(p.MitraContactPersonTempFirstName))
	p.MitraContactPersonTempLastName = html.EscapeString(strings.TrimSpace(p.MitraContactPersonTempLastName))
	p.MitraContactPersonTempJobTitle = html.EscapeString(strings.TrimSpace(p.MitraContactPersonTempJobTitle))
	p.MitraContactPersonTempMail = html.EscapeString(strings.TrimSpace(p.MitraContactPersonTempMail))
	p.MitraContactPersonTempMailWork = html.EscapeString(strings.TrimSpace(p.MitraContactPersonTempMailWork))
	p.MitraContactPersonTempPhone = html.EscapeString(strings.TrimSpace(p.MitraContactPersonTempPhone))
	p.MitraContactPersonTempPhoneWork = html.EscapeString(strings.TrimSpace(p.MitraContactPersonTempPhoneWork))
	p.MitraContactPersonTempExtensionNumber = html.EscapeString(strings.TrimSpace(p.MitraContactPersonTempExtensionNumber))
	p.MitraContactPersonTempStatus = p.MitraContactPersonTempStatus
	p.MitraContactPersonTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraContactPersonTempCreatedBy))
	p.MitraContactPersonTempReason = html.EscapeString(strings.TrimSpace(p.MitraContactPersonTempReason))
	p.MitraContactPersonTempCreatedAt = time.Now()
}

func (p *MitraContactPersonTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.MitraBranchTempID == 0 {
			return errors.New("Required MitraContactPerson Code")
		}
		if p.MitraContactPersonTempFirstName == "" {
			return errors.New("Required MitraContactPerson Name")
		}
		if p.MitraContactPersonTempLastName == "" {
			return errors.New("Required MitraContactPerson Name")
		}
		if p.MitraContactPersonTempJobTitle == "" {
			return errors.New("Required MitraContactPerson Name")
		}
		if p.MitraContactPersonTempMail == "" {
			return errors.New("Required MitraContactPerson Name")
		}
		if p.MitraContactPersonTempMailWork == "" {
			return errors.New("Required MitraContactPerson Name")
		}
		if p.MitraContactPersonTempPhone == "" {
			return errors.New("Required MitraContactPerson Name")
		}
		if p.MitraContactPersonTempPhoneWork == "" {
			return errors.New("Required MitraContactPerson Name")
		}
		return nil
	}
}

func (p *MitraContactPersonTemp) SaveMitraContactPersonTemp(db *gorm.DB) (*MitraContactPersonTemp, error) {
	var err error
	err = db.Debug().Model(&MitraContactPersonTemp{}).Create(&p).Error
	if err != nil {
		return &MitraContactPersonTemp{}, err
	}
	return p, nil
}

func (p *MitraContactPersonTempPend) SaveMitraContactPersonTempPend(db *gorm.DB) (*MitraContactPersonTempPend, error) {
	var err error
	err = db.Debug().Model(&MitraContactPersonTempPend{}).Create(&p).Error
	if err != nil {
		return &MitraContactPersonTempPend{}, err
	}
	return p, nil
}

func (p *MitraContactPersonTemp) FindAllMitraContactPersonTemp(db *gorm.DB) (*[]MitraContactPersonTempPend, error) {
	var err error
	mitracontactpersontemp := []MitraContactPersonTempPend{}
	err = db.Debug().Model(&MitraContactPersonTempPend{}).Limit(100).
		Select(`m_mitra_contact_person_temp.mitra_cp_temp_id,
				m_mitra_contact_person_temp.mitra_cp_id,
				m_mitra_contact_person_temp.mitra_branch_temp_id,
				m_mitra_contact_person_temp.mitra_cp_temp_first_name,
				m_mitra_contact_person_temp.mitra_cp_temp_last_name,
				m_mitra_contact_person_temp.mitra_cp_temp_job_title,
				m_mitra_contact_person_temp.mitra_cp_temp_mail,
				m_mitra_contact_person_temp.mitra_cp_temp_mail_work,
				m_mitra_contact_person_temp.mitra_cp_temp_phone,
				m_mitra_contact_person_temp.mitra_cp_temp_phone_work,
				m_mitra_contact_person_temp.mitra_cp_temp_ext_no,
				m_mitra_contact_person_temp.mitra_cp_temp_status,
				m_mitra_contact_person_temp.mitra_cp_temp_action_before,
				m_mitra_contact_person_temp.mitra_cp_temp_reason,
				m_mitra_contact_person_temp.mitra_cp_temp_created_by,
				m_mitra_contact_person_temp.mitra_cp_temp_created_at`).
		Find(&mitracontactpersontemp).Error
	if err != nil {
		return &[]MitraContactPersonTempPend{}, err
	}
	return &mitracontactpersontemp, nil
}

func (p *MitraContactPersonTemp) FindAllMitraContactPersonTempStatusPendingActive(db *gorm.DB) (*[]MitraContactPersonTempPend, error) {
	var err error
	mitracontactpersontemp := []MitraContactPersonTempPend{}
	err = db.Debug().Model(&MitraContactPersonTempPend{}).Limit(100).
		Select(`m_mitra_contact_person_temp.mitra_cp_temp_id,
				m_mitra_contact_person_temp.mitra_cp_id,
				m_mitra_contact_person_temp.mitra_branch_temp_id,
				m_mitra_contact_person_temp.mitra_cp_temp_first_name,
				m_mitra_contact_person_temp.mitra_cp_temp_last_name,
				m_mitra_contact_person_temp.mitra_cp_temp_job_title,
				m_mitra_contact_person_temp.mitra_cp_temp_mail,
				m_mitra_contact_person_temp.mitra_cp_temp_mail_work,
				m_mitra_contact_person_temp.mitra_cp_temp_phone,
				m_mitra_contact_person_temp.mitra_cp_temp_phone_work,
				m_mitra_contact_person_temp.mitra_cp_temp_ext_no,
				m_mitra_contact_person_temp.mitra_cp_temp_status,
				m_mitra_contact_person_temp.mitra_cp_temp_action_before,
				m_mitra_contact_person_temp.mitra_cp_temp_reason,
				m_mitra_contact_person_temp.mitra_cp_temp_created_by,
				m_mitra_contact_person_temp.mitra_cp_temp_created_at`).
		Where("mitra_cp_temp_status = ?", 11).
		Find(&mitracontactpersontemp).Error
	if err != nil {
		return &[]MitraContactPersonTempPend{}, err
	}
	return &mitracontactpersontemp, nil
}

func (p *MitraContactPersonTemp) FindAllMitraContactPersonTempStatus(db *gorm.DB, status uint64) (*[]MitraContactPersonTempData, error) {
	var err error
	mitracontactpersontemp := []MitraContactPersonTempData{}
	err = db.Debug().Model(&MitraContactPersonTempData{}).Limit(100).
		Select(`m_mitra_contact_person_temp.mitra_cp_temp_id,
				m_mitra_contact_person.mitra_cp_id,
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
				m_mitra_contact_person.mitra_cp_created_at,
				m_mitra_contact_person_temp.mitra_branch_temp_id,
				m_mitra_contact_person_temp.mitra_cp_temp_first_name,
				m_mitra_contact_person_temp.mitra_cp_temp_last_name,
				m_mitra_contact_person_temp.mitra_cp_temp_job_title,
				m_mitra_contact_person_temp.mitra_cp_temp_mail,
				m_mitra_contact_person_temp.mitra_cp_temp_mail_work,
				m_mitra_contact_person_temp.mitra_cp_temp_phone,
				m_mitra_contact_person_temp.mitra_cp_temp_phone_work,
				m_mitra_contact_person_temp.mitra_cp_temp_ext_no,
				m_mitra_contact_person_temp.mitra_cp_temp_status,
				m_mitra_contact_person_temp.mitra_cp_temp_action_before,
				m_mitra_contact_person_temp.mitra_cp_temp_reason,
				m_mitra_contact_person_temp.mitra_cp_temp_created_by,
				m_mitra_contact_person_temp.mitra_cp_temp_created_at`).
		Joins("JOIN m_mitra_contact_person on m_mitra_contact_person_temp.mitra_cp_id=m_mitra_contact_person.mitra_cp_id").
		Joins("JOIN m_mitra_branch on m_mitra_contact_person_temp.mitra_branch_temp_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_address on m_mitra_branch.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Joins("JOIN m_mitra on m_mitra_branch.mitra_id=m_mitra.mitra_id").
		Where("mitra_cp_temp_status = ?", status).
		Find(&mitracontactpersontemp).Error
	if err != nil {
		return &[]MitraContactPersonTempData{}, err
	}
	return &mitracontactpersontemp, nil
}

func (p *MitraContactPersonTemp) FindAllMitraContactPersonTempByMitraBranch(db *gorm.DB, mitrabranchtemp uint64) (*[]MitraContactPersonTempPend, error) {
	var err error
	mitracontactpersontemp := []MitraContactPersonTempPend{}
	err = db.Debug().Model(&MitraContactPersonTempPend{}).Limit(100).
		Select(`m_mitra_contact_person_temp.mitra_cp_temp_id,
				m_mitra_contact_person_temp.mitra_cp_id,
				m_mitra_contact_person_temp.mitra_branch_temp_id,
				m_mitra_contact_person_temp.mitra_cp_temp_first_name,
				m_mitra_contact_person_temp.mitra_cp_temp_last_name,
				m_mitra_contact_person_temp.mitra_cp_temp_job_title,
				m_mitra_contact_person_temp.mitra_cp_temp_mail,
				m_mitra_contact_person_temp.mitra_cp_temp_mail_work,
				m_mitra_contact_person_temp.mitra_cp_temp_phone,
				m_mitra_contact_person_temp.mitra_cp_temp_phone_work,
				m_mitra_contact_person_temp.mitra_cp_temp_ext_no,
				m_mitra_contact_person_temp.mitra_cp_temp_status,
				m_mitra_contact_person_temp.mitra_cp_temp_action_before,
				m_mitra_contact_person_temp.mitra_cp_temp_reason,
				m_mitra_contact_person_temp.mitra_cp_temp_created_by,
				m_mitra_contact_person_temp.mitra_cp_temp_created_at`).
		Where("mitra_branch_temp_id = ?", mitrabranchtemp).
		Find(&mitracontactpersontemp).Error
	if err != nil {
		return &[]MitraContactPersonTempPend{}, err
	}
	return &mitracontactpersontemp, nil
}

func (p *MitraContactPersonTemp) FindMitraContactPersonTempDataByID(db *gorm.DB, pid uint64) (*MitraContactPersonTemp, error) {
	var err error
	err = db.Debug().Model(&MitraContactPersonTemp{}).Where("mitra_cp_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraContactPersonTemp{}, err
	}
	return p, nil
}

func (p *MitraContactPersonTemp) FindMitraContactPersonTempByIDPendingActive(db *gorm.DB, pid uint64) (*MitraContactPersonTempPend, error) {
	var err error
	mitracontactpersontemp := MitraContactPersonTempPend{}
	err = db.Debug().Model(&MitraContactPersonTempPend{}).Limit(100).
		Select(`m_mitra_contact_person_temp.mitra_cp_temp_id,
				m_mitra_contact_person_temp.mitra_cp_id,
				m_mitra_contact_person_temp.mitra_branch_temp_id,
				m_mitra_contact_person_temp.mitra_cp_temp_first_name,
				m_mitra_contact_person_temp.mitra_cp_temp_last_name,
				m_mitra_contact_person_temp.mitra_cp_temp_job_title,
				m_mitra_contact_person_temp.mitra_cp_temp_mail,
				m_mitra_contact_person_temp.mitra_cp_temp_mail_work,
				m_mitra_contact_person_temp.mitra_cp_temp_phone,
				m_mitra_contact_person_temp.mitra_cp_temp_phone_work,
				m_mitra_contact_person_temp.mitra_cp_temp_ext_no,
				m_mitra_contact_person_temp.mitra_cp_temp_status,
				m_mitra_contact_person_temp.mitra_cp_temp_action_before,
				m_mitra_contact_person_temp.mitra_cp_temp_reason,
				m_mitra_contact_person_temp.mitra_cp_temp_created_by,
				m_mitra_contact_person_temp.mitra_cp_temp_created_at`).
		Where("mitra_cp_temp_id = ?", pid).
		Find(&mitracontactpersontemp).Error
	if err != nil {
		return &MitraContactPersonTempPend{}, err
	}
	return &mitracontactpersontemp, nil
}

func (p *MitraContactPersonTemp) FindMitraContactPersonTempStatusByIDPendingActive(db *gorm.DB, pid uint64, status uint64) (*MitraContactPersonTempPend, error) {
	var err error
	mitracontactpersontemp := MitraContactPersonTempPend{}
	err = db.Debug().Model(&MitraContactPersonTempPend{}).Limit(100).
		Select(`m_mitra_contact_person_temp.mitra_cp_temp_id,
				m_mitra_contact_person_temp.mitra_cp_id,
				m_mitra_contact_person_temp.mitra_branch_temp_id,
				m_mitra_contact_person_temp.mitra_cp_temp_first_name,
				m_mitra_contact_person_temp.mitra_cp_temp_last_name,
				m_mitra_contact_person_temp.mitra_cp_temp_job_title,
				m_mitra_contact_person_temp.mitra_cp_temp_mail,
				m_mitra_contact_person_temp.mitra_cp_temp_mail_work,
				m_mitra_contact_person_temp.mitra_cp_temp_phone,
				m_mitra_contact_person_temp.mitra_cp_temp_phone_work,
				m_mitra_contact_person_temp.mitra_cp_temp_ext_no,
				m_mitra_contact_person_temp.mitra_cp_temp_status,
				m_mitra_contact_person_temp.mitra_cp_temp_action_before,
				m_mitra_contact_person_temp.mitra_cp_temp_reason,
				m_mitra_contact_person_temp.mitra_cp_temp_created_by,
				m_mitra_contact_person_temp.mitra_cp_temp_created_at`).
		Where("mitra_cp_temp_id = ? AND mitra_cp_temp_status = ?", pid, status).
		Find(&mitracontactpersontemp).Error
	if err != nil {
		return &MitraContactPersonTempPend{}, err
	}
	return &mitracontactpersontemp, nil
}

func (p *MitraContactPersonTemp) FindMitraContactPersonTempByID(db *gorm.DB, pid uint64) (*MitraContactPersonTempData, error) {
	var err error
	mitracontactpersontemp := MitraContactPersonTempData{}
	err = db.Debug().Model(&MitraContactPersonTempData{}).Limit(100).
		Select(`m_mitra_contact_person_temp.mitra_cp_temp_id,
				m_mitra_contact_person.mitra_cp_id,
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
				m_mitra_contact_person.mitra_cp_created_at,
				m_mitra_contact_person_temp.mitra_branch_temp_id,
				m_mitra_contact_person_temp.mitra_cp_temp_first_name,
				m_mitra_contact_person_temp.mitra_cp_temp_last_name,
				m_mitra_contact_person_temp.mitra_cp_temp_job_title,
				m_mitra_contact_person_temp.mitra_cp_temp_mail,
				m_mitra_contact_person_temp.mitra_cp_temp_mail_work,
				m_mitra_contact_person_temp.mitra_cp_temp_phone,
				m_mitra_contact_person_temp.mitra_cp_temp_phone_work,
				m_mitra_contact_person_temp.mitra_cp_temp_ext_no,
				m_mitra_contact_person_temp.mitra_cp_temp_status,
				m_mitra_contact_person_temp.mitra_cp_temp_action_before,
				m_mitra_contact_person_temp.mitra_cp_temp_reason,
				m_mitra_contact_person_temp.mitra_cp_temp_created_by,
				m_mitra_contact_person_temp.mitra_cp_temp_created_at`).
		Joins("JOIN m_mitra_contact_person on m_mitra_contact_person_temp.mitra_cp_id=m_mitra_contact_person.mitra_cp_id").
		Joins("JOIN m_mitra_branch on m_mitra_contact_person_temp.mitra_branch_temp_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_address on m_mitra_branch.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Joins("JOIN m_mitra on m_mitra_branch.mitra_id=m_mitra.mitra_id").
		Where("mitra_cp_temp_id = ?", pid).
		Find(&mitracontactpersontemp).Error
	if err != nil {
		return &MitraContactPersonTempData{}, err
	}
	return &mitracontactpersontemp, nil
}

func (p *MitraContactPersonTemp) FindMitraContactPersonTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraContactPersonTempData, error) {
	var err error
	mitracontactpersontemp := MitraContactPersonTempData{}
	err = db.Debug().Model(&MitraContactPersonTempData{}).Limit(100).
		Select(`m_mitra_contact_person_temp.mitra_cp_temp_id,
				m_mitra_contact_person.mitra_cp_id,
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
				m_mitra_contact_person.mitra_cp_created_at,
				m_mitra_contact_person_temp.mitra_branch_temp_id,
				m_mitra_contact_person_temp.mitra_cp_temp_first_name,
				m_mitra_contact_person_temp.mitra_cp_temp_last_name,
				m_mitra_contact_person_temp.mitra_cp_temp_job_title,
				m_mitra_contact_person_temp.mitra_cp_temp_mail,
				m_mitra_contact_person_temp.mitra_cp_temp_mail_work,
				m_mitra_contact_person_temp.mitra_cp_temp_phone,
				m_mitra_contact_person_temp.mitra_cp_temp_phone_work,
				m_mitra_contact_person_temp.mitra_cp_temp_ext_no,
				m_mitra_contact_person_temp.mitra_cp_temp_status,
				m_mitra_contact_person_temp.mitra_cp_temp_action_before,
				m_mitra_contact_person_temp.mitra_cp_temp_reason,
				m_mitra_contact_person_temp.mitra_cp_temp_created_by,
				m_mitra_contact_person_temp.mitra_cp_temp_created_at`).
		Joins("JOIN m_mitra_contact_person on m_mitra_contact_person_temp.mitra_cp_id=m_mitra_contact_person.mitra_cp_id").
		Joins("JOIN m_mitra_branch on m_mitra_contact_person_temp.mitra_branch_temp_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_address on m_mitra_branch.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Joins("JOIN m_mitra on m_mitra_branch.mitra_id=m_mitra.mitra_id").
		Where("mitra_cp_temp_id = ? AND mitra_cp_temp_status = ?", pid, status).
		Find(&mitracontactpersontemp).Error
	if err != nil {
		return &MitraContactPersonTempData{}, err
	}
	return &mitracontactpersontemp, nil
}

func (p *MitraContactPersonTemp) FindMitraContactPersonTempByMitraBranchByID(db *gorm.DB, pid uint64, mitrabranchtemp uint64) (*MitraContactPersonTempPend, error) {
	var err error
	mitracontactpersontemp := MitraContactPersonTempPend{}
	err = db.Debug().Model(&MitraContactPersonTempPend{}).Limit(100).
		Select(`m_mitra_contact_person_temp.mitra_cp_temp_id,
				m_mitra_contact_person_temp.mitra_cp_id,
				m_mitra_contact_person_temp.mitra_branch_temp_id,
				m_mitra_contact_person_temp.mitra_cp_temp_first_name,
				m_mitra_contact_person_temp.mitra_cp_temp_last_name,
				m_mitra_contact_person_temp.mitra_cp_temp_job_title,
				m_mitra_contact_person_temp.mitra_cp_temp_mail,
				m_mitra_contact_person_temp.mitra_cp_temp_mail_work,
				m_mitra_contact_person_temp.mitra_cp_temp_phone,
				m_mitra_contact_person_temp.mitra_cp_temp_phone_work,
				m_mitra_contact_person_temp.mitra_cp_temp_ext_no,
				m_mitra_contact_person_temp.mitra_cp_temp_status,
				m_mitra_contact_person_temp.mitra_cp_temp_action_before,
				m_mitra_contact_person_temp.mitra_cp_temp_reason,
				m_mitra_contact_person_temp.mitra_cp_temp_created_by,
				m_mitra_contact_person_temp.mitra_cp_temp_created_at`).
		Where("mitra_cp_temp_id = ? AND mitra_branch_temp_id = ?", pid, mitrabranchtemp).
		Find(&mitracontactpersontemp).Error
	if err != nil {
		return &MitraContactPersonTempPend{}, err
	}
	return &mitracontactpersontemp, nil
}
func (p *MitraContactPersonTemp) UpdateMitraContactPersonBranchTempID(db *gorm.DB, pid uint64) (*MitraContactPersonTemp, error) {
	var err error
	db = db.Debug().Model(&MitraContactPersonTemp{}).Where("mitra_cp_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_branch_temp_id":     p.MitraBranchTempID,
			"mitra_cp_temp_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraContactPersonTemp{}).Where("mitra_cp_temp_id = ?", pid).Error
	if err != nil {
		return &MitraContactPersonTemp{}, err
	}
	return p, nil
}

func (p *MitraContactPersonTemp) UpdateMitraContactPersonTemp(db *gorm.DB, pid uint64) (*MitraContactPersonTemp, error) {
	var err error
	db = db.Debug().Model(&MitraContactPersonTemp{}).Where("mitra_cp_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_cp_id":              p.MitraContactPersonID,
			"mitra_branch_temp_id":     p.MitraBranchTempID,
			"mitra_cp_temp_first_name": p.MitraContactPersonTempFirstName,
			"mitra_cp_temp_last_name":  p.MitraContactPersonTempLastName,
			"mitra_cp_temp_job_title":  p.MitraContactPersonTempJobTitle,
			"mitra_cp_temp_mail":       p.MitraContactPersonTempMail,
			"mitra_cp_temp_mail_work":  p.MitraContactPersonTempMailWork,
			"mitra_cp_temp_phone":      p.MitraContactPersonTempPhone,
			"mitra_cp_temp_phone_work": p.MitraContactPersonTempPhoneWork,
			"mitra_cp_temp_ext_no":     p.MitraContactPersonTempExtensionNumber,
			"mitra_cp_temp_reason":     p.MitraContactPersonTempReason,
			"mitra_cp_temp_status":     p.MitraContactPersonTempStatus,
		},
	)
	err = db.Debug().Model(&MitraContactPersonTemp{}).Where("mitra_cp_temp_id = ?", pid).Error
	if err != nil {
		return &MitraContactPersonTemp{}, err
	}
	return p, nil
}

func (p *MitraContactPersonTemp) UpdateMitraContactPersonTempStatus(db *gorm.DB, pid uint64) (*MitraContactPersonTemp, error) {
	var err error
	db = db.Debug().Model(&MitraContactPersonTemp{}).Where("mitra_cp_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_cp_temp_status": p.MitraContactPersonTempStatus,
		},
	)
	err = db.Debug().Model(&MitraContactPersonTemp{}).Where("mitra_cp_temp_id = ?", pid).Error
	if err != nil {
		return &MitraContactPersonTemp{}, err
	}
	return p, nil
}

func (p *MitraContactPersonTemp) DeleteMitraContactPersonTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraContactPersonTemp{}).Where("mitra_cp_temp_id = ?", pid).Take(&MitraContactPersonTemp{}).Delete(&MitraContactPersonTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraContactPersonTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

func (p *MitraContactPersonTemp) FindMitraContactPersonTempPendByMitraBranchTempID(db *gorm.DB, pid uint64) ([]MitraContactPersonTempPendData, error) {
	var err error
	mitrabranchcoverageregency := []MitraContactPersonTempPendData{}
	err = db.Debug().Model(&MitraContactPersonTempPendData{}).Limit(100).
		Select(`m_mitra_contact_person_temp.mitra_cp_temp_id,
				m_mitra_contact_person_temp.mitra_cp_id,
				m_mitra_branch_temp.mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_contact_person_temp.mitra_cp_temp_first_name,
				m_mitra_contact_person_temp.mitra_cp_temp_last_name,
				m_mitra_contact_person_temp.mitra_cp_temp_job_title,
				m_mitra_contact_person_temp.mitra_cp_temp_mail,
				m_mitra_contact_person_temp.mitra_cp_temp_mail_work,
				m_mitra_contact_person_temp.mitra_cp_temp_phone,
				m_mitra_contact_person_temp.mitra_cp_temp_phone_work,
				m_mitra_contact_person_temp.mitra_cp_temp_ext_no,
				m_mitra_contact_person_temp.mitra_cp_temp_status,
				m_mitra_contact_person_temp.mitra_cp_temp_action_before,
				m_mitra_contact_person_temp.mitra_cp_temp_reason,
				m_mitra_contact_person_temp.mitra_cp_temp_created_by,
				m_mitra_contact_person_temp.mitra_cp_temp_created_at`).
		Joins("JOIN m_mitra_branch_temp on m_mitra_contact_person_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_temp_id").
		Where("m_mitra_contact_person_temp.mitra_branch_temp_id = ?", pid).
		Find(&mitrabranchcoverageregency).Error
	if err != nil {
		return []MitraContactPersonTempPendData{}, err
	}
	return mitrabranchcoverageregency, nil
}

func (p *MitraContactPersonTemp) FindMitraContactPersonTempPendStatusByMitraBranchTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraContactPersonTempPendData, error) {
	var err error
	mitrabranchcoverageregency := []MitraContactPersonTempPendData{}
	err = db.Debug().Model(&MitraContactPersonTempPendData{}).Limit(100).
		Select(`m_mitra_contact_person_temp.mitra_cp_temp_id,
				m_mitra_contact_person_temp.mitra_cp_id,
				m_mitra_branch_temp.mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_contact_person_temp.mitra_cp_temp_first_name,
				m_mitra_contact_person_temp.mitra_cp_temp_last_name,
				m_mitra_contact_person_temp.mitra_cp_temp_job_title,
				m_mitra_contact_person_temp.mitra_cp_temp_mail,
				m_mitra_contact_person_temp.mitra_cp_temp_mail_work,
				m_mitra_contact_person_temp.mitra_cp_temp_phone,
				m_mitra_contact_person_temp.mitra_cp_temp_phone_work,
				m_mitra_contact_person_temp.mitra_cp_temp_ext_no,
				m_mitra_contact_person_temp.mitra_cp_temp_status,
				m_mitra_contact_person_temp.mitra_cp_temp_action_before,
				m_mitra_contact_person_temp.mitra_cp_temp_reason,
				m_mitra_contact_person_temp.mitra_cp_temp_created_by,
				m_mitra_contact_person_temp.mitra_cp_temp_created_at`).
		Joins("JOIN m_mitra_branch_temp on m_mitra_contact_person_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_temp_id").
		Where("m_mitra_contact_person_temp.mitra_branch_temp_id = ? AND mitra_cp_temp_status = ?", pid, status).
		Find(&mitrabranchcoverageregency).Error
	if err != nil {
		return []MitraContactPersonTempPendData{}, err
	}
	return mitrabranchcoverageregency, nil
}

func (p *MitraContactPersonTemp) FindMitraContactPersonTempByMitraBranchTempID(db *gorm.DB, pid uint64) ([]MitraContactPersonTempData, error) {
	var err error
	mitrabranchcoverageregency := []MitraContactPersonTempData{}
	err = db.Debug().Model(&MitraContactPersonTempData{}).Limit(100).
		Select(`m_mitra_contact_person_temp.mitra_cp_temp_id,
				m_mitra_contact_person_temp.mitra_cp_id,
				m_mitra_branch_temp.mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_contact_person_temp.mitra_cp_temp_first_name,
				m_mitra_contact_person_temp.mitra_cp_temp_last_name,
				m_mitra_contact_person_temp.mitra_cp_temp_job_title,
				m_mitra_contact_person_temp.mitra_cp_temp_mail,
				m_mitra_contact_person_temp.mitra_cp_temp_mail_work,
				m_mitra_contact_person_temp.mitra_cp_temp_phone,
				m_mitra_contact_person_temp.mitra_cp_temp_phone_work,
				m_mitra_contact_person_temp.mitra_cp_temp_ext_no,
				m_mitra_contact_person_temp.mitra_cp_temp_status,
				m_mitra_contact_person_temp.mitra_cp_temp_action_before,
				m_mitra_contact_person_temp.mitra_cp_temp_reason,
				m_mitra_contact_person_temp.mitra_cp_temp_created_by,
				m_mitra_contact_person_temp.mitra_cp_temp_created_at`).
		Joins("JOIN m_mitra_branch_temp on m_mitra_contact_person_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_temp_id").
		Where("m_mitra_contact_person_temp.mitra_branch_temp_id = ?", pid).
		Find(&mitrabranchcoverageregency).Error
	if err != nil {
		return []MitraContactPersonTempData{}, err
	}
	return mitrabranchcoverageregency, nil
}

func (p *MitraContactPersonTemp) FindMitraContactPersonTempStatusByMitraBranchTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraContactPersonTempData, error) {
	var err error
	mitrabranchcoverageregency := []MitraContactPersonTempData{}
	err = db.Debug().Model(&MitraContactPersonTempData{}).Limit(100).
		Select(`m_mitra_contact_person_temp.mitra_cp_temp_id,
				m_mitra_contact_person_temp.mitra_cp_id,
				m_mitra_branch_temp.mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_contact_person_temp.mitra_cp_temp_first_name,
				m_mitra_contact_person_temp.mitra_cp_temp_last_name,
				m_mitra_contact_person_temp.mitra_cp_temp_job_title,
				m_mitra_contact_person_temp.mitra_cp_temp_mail,
				m_mitra_contact_person_temp.mitra_cp_temp_mail_work,
				m_mitra_contact_person_temp.mitra_cp_temp_phone,
				m_mitra_contact_person_temp.mitra_cp_temp_phone_work,
				m_mitra_contact_person_temp.mitra_cp_temp_ext_no,
				m_mitra_contact_person_temp.mitra_cp_temp_status,
				m_mitra_contact_person_temp.mitra_cp_temp_action_before,
				m_mitra_contact_person_temp.mitra_cp_temp_reason,
				m_mitra_contact_person_temp.mitra_cp_temp_created_by,
				m_mitra_contact_person_temp.mitra_cp_temp_created_at`).
		Joins("JOIN m_mitra_branch_temp on m_mitra_contact_person_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_temp_id").
		Where("m_mitra_contact_person_temp.mitra_branch_temp_id = ? AND mitra_cp_temp_status = ?", pid, status).
		Find(&mitrabranchcoverageregency).Error
	if err != nil {
		return []MitraContactPersonTempData{}, err
	}
	return mitrabranchcoverageregency, nil
}

func (p *MitraContactPersonTemp) FindMitraContactPersonTempDataStatusByMitraBranchTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraContactPersonTempData, error) {
	var err error
	mitrabranchcoverageregency := []MitraContactPersonTempData{}
	err = db.Debug().Model(&MitraContactPersonTempData{}).Limit(100).
		Select(`m_mitra_contact_person_temp.mitra_cp_temp_id,
				m_mitra_contact_person_temp.mitra_cp_id,
				m_mitra_branch_temp.mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_contact_person_temp.mitra_cp_temp_first_name,
				m_mitra_contact_person_temp.mitra_cp_temp_last_name,
				m_mitra_contact_person_temp.mitra_cp_temp_job_title,
				m_mitra_contact_person_temp.mitra_cp_temp_mail,
				m_mitra_contact_person_temp.mitra_cp_temp_mail_work,
				m_mitra_contact_person_temp.mitra_cp_temp_phone,
				m_mitra_contact_person_temp.mitra_cp_temp_phone_work,
				m_mitra_contact_person_temp.mitra_cp_temp_ext_no,
				m_mitra_contact_person_temp.mitra_cp_temp_status,
				m_mitra_contact_person_temp.mitra_cp_temp_action_before,
				m_mitra_contact_person_temp.mitra_cp_temp_reason,
				m_mitra_contact_person_temp.mitra_cp_temp_created_by,
				m_mitra_contact_person_temp.mitra_cp_temp_created_at`).
		Joins("JOIN m_mitra_branch_temp on m_mitra_contact_person_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_temp_id").
		Where("m_mitra_contact_person_temp.mitra_branch_temp_id = ? AND mitra_cp_temp_status = ?", pid, status).
		Find(&mitrabranchcoverageregency).Error
	if err != nil {
		return []MitraContactPersonTempData{}, err
	}
	return mitrabranchcoverageregency, nil
}
