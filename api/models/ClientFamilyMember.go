package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type ClientFamilyMember struct {
	ClientFamilyMemberID           uint64     `gorm:"column:client_fmly_mmbr_id;primary_key;" json:"client_fmly_mmbr_id"`
	ClientFamilyID                 uint64     `gorm:"column:client_fmly_id" json:"client_fmly_id"`
	ClientFamilyMemberNIK          string     `gorm:"column:client_fmly_mmbr_nik;size:20" json:"client_fmly_mmbr_nik"`
	ClientFamilyMemberName         string     `gorm:"column:client_fmly_mmbr_name;size:255" json:"client_fmly_mmbr_name"`
	ClientFamilyMemberRelationship string     `gorm:"column:client_fmly_mmbr_relationship;size:125" json:"client_fmly_mmbr_relationship"`
	ClientFamilyMemberSex          string     `gorm:"column:client_fmly_mmbr_sex;size:1" json:"client_fmly_mmbr_sex"`
	ClientFamilyMemberAge          string     `gorm:"column:client_fmly_mmbr_age;size:3" json:"client_fmly_mmbr_age"`
	ClientFamilyMemberOccupation   string     `gorm:"column:client_fmly_mmbr_occupation;size:125" json:"client_fmly_mmbr_occupation"`
	ClientFamilyMemberCompany      string     `gorm:"column:client_fmly_mmbr_company;size:125" json:"client_fmly_mmbr_company"`
	ClientFamilyMemberEducation    string     `gorm:"column:client_fmly_mmbr_education;size:125" json:"client_fmly_mmbr_education"`
	ClientFamilyMemberSchool       string     `gorm:"column:client_fmly_mmbr_school;size:125" json:"client_fmly_mmbr_school"`
	ClientFamilyMemberMajor        string     `gorm:"column:client_fmly_mmbr_major;size:125" json:"client_fmly_mmbr_major"`
	ClientFamilyMemberRemark       string     `gorm:"column:client_fmly_mmbr_remark;size:125" json:"client_fmly_mmbr_remark"`
	ClientFamilyMemberStatus       uint64     `gorm:"column:client_fmly_mmbr_status;size:2" json:"client_fmly_mmbr_status"`
	ClientFamilyMemberCreatedBy    string     `gorm:"column:client_fmly_mmbr_created_by;size:125" json:"client_fmly_mmbr_created_by"`
	ClientFamilyMemberCreatedAt    time.Time  `gorm:"column:client_fmly_mmbr_created_at;default:CURRENT_TIMESTAMP" json:"client_fmly_mmbr_created_at"`
	ClientFamilyMemberUpdatedBy    string     `gorm:"column:client_fmly_mmbr_updated_by;size:125" json:"client_fmly_mmbr_updated_by"`
	ClientFamilyMemberUpdatedAt    *time.Time `gorm:"column:client_fmly_mmbr_updated_at;default:CURRENT_TIMESTAMP" json:"client_fmly_mmbr_updated_at"`
	ClientFamilyMemberDeletedBy    string     `gorm:"column:client_fmly_mmbr_deleted_by;size:125" json:"client_fmly_mmbr_deleted_by"`
	ClientFamilyMemberDeletedAt    *time.Time `gorm:"column:client_fmly_mmbr_deleted_at;default:CURRENT_TIMESTAMP" json:"client_fmly_mmbr_deleted_at"`
}

type ResponseClientFamilyMembers struct {
	Status  uint64                `json:"status"`
	Message string                `json:"message"`
	Data    *[]ClientFamilyMember `json:"data"`
}

type ResponseClientFamilyMember struct {
	Status  uint64              `json:"status"`
	Message string              `json:"message"`
	Data    *ClientFamilyMember `json:"data"`
}

type ResponseClientFamilyMemberCUD struct {
	Status  uint64 `json:"status"`
	Message string `json:"message"`
}

func (ClientFamilyMember) TableName() string {
	return "m_client_family"
}

func (p *ClientFamilyMember) Prepare() {
	p.ClientFamilyID = p.ClientFamilyID
	p.ClientFamilyMemberName = html.EscapeString(strings.TrimSpace(p.ClientFamilyMemberName))
	p.ClientFamilyMemberRelationship = html.EscapeString(strings.TrimSpace(p.ClientFamilyMemberRelationship))
	p.ClientFamilyMemberSex = html.EscapeString(strings.TrimSpace(p.ClientFamilyMemberSex))
	p.ClientFamilyMemberAge = html.EscapeString(strings.TrimSpace(p.ClientFamilyMemberAge))
	p.ClientFamilyMemberOccupation = html.EscapeString(strings.TrimSpace(p.ClientFamilyMemberOccupation))
	p.ClientFamilyMemberCompany = html.EscapeString(strings.TrimSpace(p.ClientFamilyMemberCompany))
	p.ClientFamilyMemberEducation = html.EscapeString(strings.TrimSpace(p.ClientFamilyMemberEducation))
	p.ClientFamilyMemberSchool = html.EscapeString(strings.TrimSpace(p.ClientFamilyMemberSchool))
	p.ClientFamilyMemberMajor = html.EscapeString(strings.TrimSpace(p.ClientFamilyMemberMajor))
	p.ClientFamilyMemberRemark = html.EscapeString(strings.TrimSpace(p.ClientFamilyMemberRemark))
	p.ClientFamilyMemberStatus = p.ClientFamilyMemberStatus
	p.ClientFamilyMemberCreatedAt = time.Now()
	p.ClientFamilyMemberCreatedBy = p.ClientFamilyMemberCreatedBy
	p.ClientFamilyMemberUpdatedAt = p.ClientFamilyMemberUpdatedAt
	p.ClientFamilyMemberUpdatedBy = p.ClientFamilyMemberUpdatedBy
	p.ClientFamilyMemberDeletedAt = p.ClientFamilyMemberDeletedAt
	p.ClientFamilyMemberDeletedBy = p.ClientFamilyMemberDeletedBy
}

func (p *ClientFamilyMember) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.ClientFamilyMemberName == "" {
			return errors.New("Required ClientFamilyMember Code")
		}
		if p.ClientFamilyMemberRelationship == "" {
			return errors.New("Required KTP Number")
		}
		if p.ClientFamilyMemberSex == "" {
			return errors.New("Required Passport Number")
		}
		if p.ClientFamilyMemberAge == "" {
			return errors.New("Required Passport Number")
		}
		if p.ClientFamilyMemberOccupation == "" {
			return errors.New("Required Passport Number")
		}
		if p.ClientFamilyMemberCompany == "" {
			return errors.New("Required Passport Number")
		}
		if p.ClientFamilyMemberEducation == "" {
			return errors.New("Required Passport Number")
		}
		if p.ClientFamilyMemberSchool == "" {
			return errors.New("Required Passport Number")
		}
		if p.ClientFamilyMemberMajor == "" {
			return errors.New("Required NPWP Number")
		}
		if p.ClientFamilyMemberRemark == "" {
			return errors.New("Required ClientFamilyMember Name")
		}
		return nil
	}
}

func (p *ClientFamilyMember) SaveClientFamilyMember(db *gorm.DB) (*ClientFamilyMember, error) {
	var err error
	err = db.Debug().Model(&ClientFamilyMember{}).Create(&p).Error
	if err != nil {
		return &ClientFamilyMember{}, err
	}
	return p, nil
}

func (p *ClientFamilyMember) FindAllClientFamilyMember(db *gorm.DB) (*[]ClientFamilyMember, error) {
	var err error
	clientfamily := []ClientFamilyMember{}
	err = db.Debug().Model(&ClientFamilyMember{}).Limit(100).Find(&clientfamily).Error
	if err != nil {
		return &[]ClientFamilyMember{}, err
	}
	return &clientfamily, nil
}

func (p *ClientFamilyMember) FindAllClientFamilyMemberStatus(db *gorm.DB, status uint64) (*[]ClientFamilyMember, error) {
	var err error
	clientfamily := []ClientFamilyMember{}
	err = db.Debug().Model(&ClientFamilyMember{}).Where("client_fmly_mmbr_status = ?", status).Limit(100).Find(&clientfamily).Error
	if err != nil {
		return &[]ClientFamilyMember{}, err
	}
	return &clientfamily, nil
}

func (p *ClientFamilyMember) FindClientFamilyMemberByID(db *gorm.DB, pid uint64) (*ClientFamilyMember, error) {
	var err error
	err = db.Debug().Model(&ClientFamilyMember{}).Where("client_fmly_mmbr_id = ?", pid).Take(&p).Error
	if err != nil {
		return &ClientFamilyMember{}, err
	}
	return p, nil
}

func (p *ClientFamilyMember) FindClientFamilyMemberStatusByID(db *gorm.DB, pid uint64, status uint64) (*ClientFamilyMember, error) {
	var err error
	err = db.Debug().Model(&ClientFamilyMember{}).Where("client_fmly_mmbr_id = ? AND client_fmly_mmbr_status = ?", pid, status).Take(&p).Error
	if err != nil {
		return &ClientFamilyMember{}, err
	}
	return p, nil
}

func (p *ClientFamilyMember) UpdateClientFamilyMember(db *gorm.DB, pid uint64) (*ClientFamilyMember, error) {
	var err error
	db = db.Debug().Model(&ClientFamilyMember{}).Where("client_fmly_mmbr_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"client_fmly_mmbr_nik":          p.ClientFamilyMemberNIK,
			"client_fmly_mmbr_name":         p.ClientFamilyMemberName,
			"client_fmly_mmbr_relationship": p.ClientFamilyMemberRelationship,
			"client_fmly_mmbr_sex":          p.ClientFamilyMemberSex,
			"client_fmly_mmbr_age":          p.ClientFamilyMemberAge,
			"client_fmly_mmbr_occupation":   p.ClientFamilyMemberOccupation,
			"client_fmly_mmbr_company":      p.ClientFamilyMemberCompany,
			"client_fmly_mmbr_education":    p.ClientFamilyMemberEducation,
			"client_fmly_mmbr_school":       p.ClientFamilyMemberSchool,
			"client_fmly_mmbr_major":        p.ClientFamilyMemberMajor,
			"client_fmly_mmbr_remark":       p.ClientFamilyMemberRemark,
			"client_fmly_mmbr_status":       p.ClientFamilyMemberStatus,
			"client_fmly_mmbr_updated_by":   p.ClientFamilyMemberUpdatedBy,
			"client_fmly_mmbr_updated_at":   time.Now(),
		},
	)
	err = db.Debug().Model(&ClientFamilyMember{}).Where("client_fmly_mmbr_id = ?", pid).Error
	if err != nil {
		return &ClientFamilyMember{}, err
	}
	return p, nil
}

func (p *ClientFamilyMember) UpdateClientFamilyMemberStatus(db *gorm.DB, pid uint64) (*ClientFamilyMember, error) {
	var err error
	db = db.Debug().Model(&ClientFamilyMember{}).Where("client_fmly_mmbr_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"client_fmly_mmbr_status":     p.ClientFamilyMemberStatus,
			"client_fmly_mmbr_updated_by": p.ClientFamilyMemberUpdatedBy,
			"client_fmly_mmbr_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ClientFamilyMember{}).Where("client_fmly_mmbr_id = ?", pid).Error
	if err != nil {
		return &ClientFamilyMember{}, err
	}
	return p, nil
}

func (p *ClientFamilyMember) UpdateClientFamilyMemberStatusDelete(db *gorm.DB, pid uint64) (*ClientFamilyMember, error) {
	var err error
	db = db.Debug().Model(&ClientFamilyMember{}).Where("client_fmly_mmbr_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"client_fmly_mmbr_status":     3,
			"client_fmly_mmbr_deleted_by": p.ClientFamilyMemberDeletedBy,
			"client_fmly_mmbr_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ClientFamilyMember{}).Where("client_fmly_mmbr_id = ?", pid).Error
	if err != nil {
		return &ClientFamilyMember{}, err
	}
	return p, nil
}

func (p *ClientFamilyMember) DeleteClientFamilyMember(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&ClientFamilyMember{}).Where("client_fmly_mmbr_id = ?", pid).Take(&ClientFamilyMember{}).Delete(&ClientFamilyMember{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("ClientFamilyMember not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// ADDITIONAL MEMORANDUM
// ===================================================================================================================================

type ClientFamilyMemberMrdmData struct {
	ClientFamilyMemberID           uint64 `gorm:"column:client_fmly_mmbr_id" json:"client_fmly_mmbr_id"`
	ClientFamilyID                 uint64 `gorm:"column:client_fmly_id" json:"client_fmly_id"`
	ClientFamilyMemberNIK          string `gorm:"column:client_fmly_mmbr_nik" json:"client_fmly_mmbr_nik"`
	ClientFamilyMemberName         string `gorm:"column:client_fmly_mmbr_name" json:"client_fmly_mmbr_name"`
	ClientFamilyMemberRelationship string `gorm:"column:client_fmly_mmbr_relationship" json:"client_fmly_mmbr_relationship"`
	ClientFamilyMemberSex          string `gorm:"column:client_fmly_mmbr_sex" json:"client_fmly_mmbr_sex"`
	ClientFamilyMemberAge          string `gorm:"column:client_fmly_mmbr_age" json:"client_fmly_mmbr_age"`
	ClientFamilyMemberOccupation   string `gorm:"column:client_fmly_mmbr_occupation" json:"client_fmly_mmbr_occupation"`
	ClientFamilyMemberCompany      string `gorm:"column:client_fmly_mmbr_company" json:"client_fmly_mmbr_company"`
	ClientFamilyMemberEducation    string `gorm:"column:client_fmly_mmbr_education" json:"client_fmly_mmbr_education"`
	ClientFamilyMemberSchool       string `gorm:"column:client_fmly_mmbr_school" json:"client_fmly_mmbr_school"`
	ClientFamilyMemberMajor        string `gorm:"column:client_fmly_mmbr_major" json:"client_fmly_mmbr_major"`
	ClientFamilyMemberRemark       string `gorm:"column:client_fmly_mmbr_remark" json:"client_fmly_mmbr_remark"`
}

func (ClientFamilyMemberMrdmData) TableName() string {
	return "m_client_family_member"
}

func (p *ClientFamilyMember) FindClientFamilyMemberMrdmByClientID(db *gorm.DB, pid uint64) ([]ClientFamilyMemberMrdmData, error) {
	var err error
	clientfamily := []ClientFamilyMemberMrdmData{}
	err = db.Debug().Model(&ClientFamilyMemberMrdmData{}).Where("client_fmly_id = ? AND client_fmly_mmbr_status = ?", pid, 1).Find(&clientfamily).Error
	if err != nil {
		return []ClientFamilyMemberMrdmData{}, err
	}
	return clientfamily, nil
}

func (p *ClientFamilyMember) GetClientFamilyMemberByClientFamilyID(db *gorm.DB, pid uint64) (*[]ClientFamilyMember, error) {
	var err error
	clientfamilymember := []ClientFamilyMember{}
	err = db.Debug().Model(&ClientFamilyMember{}).
		Where("client_fmly_id = ? AND client_fmly_mmbr_status = ?", pid, 1).
		Find(&clientfamilymember).Error
	if err != nil {
		return &[]ClientFamilyMember{}, err
	}
	return &clientfamilymember, nil
}
