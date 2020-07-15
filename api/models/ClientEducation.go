package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type ClientEducation struct {
	ClientEducationID             uint64     `gorm:"column:client_edu_id;primary_key;" json:"client_edu_id"`
	ClientID                      uint64     `gorm:"column:client_id" json:"client_id"`
	ClientEducationGraduationYear string     `gorm:"column:client_edu_grad_year;size:4" json:"client_edu_grad_year"`
	ClientEducationInstitution    string     `gorm:"column:client_edu_institution;size:125" json:"client_edu_institution"`
	ClientEducationCity           string     `gorm:"column:client_edu_city;size:125" json:"client_edu_city"`
	ClientEducationMajor          string     `gorm:"column:client_edu_major;size:125" json:"client_edu_major"`
	ClientEducationRemark         string     `gorm:"column:client_edu_remark;size:255" json:"client_edu_remark"`
	ClientEducationStatus         uint64     `gorm:"column:client_edu_status;size:2" json:"client_edu_status"`
	ClientEducationCreatedBy      string     `gorm:"column:client_edu_created_by;size:125" json:"client_edu_created_by"`
	ClientEducationCreatedAt      time.Time  `gorm:"column:client_edu_created_at;default:CURRENT_TIMESTAMP" json:"prod_ctgry_created_at"`
	ClientEducationUpdatedBy      string     `gorm:"column:client_edu_updated_by;size:125" json:"client_edu_updated_by"`
	ClientEducationUpdatedAt      *time.Time `gorm:"column:client_edu_updated_at;default:CURRENT_TIMESTAMP" json:"prod_ctgry_updated_at"`
	ClientEducationDeletedBy      string     `gorm:"column:client_edu_deleted_by;size:125" json:"client_edu_deleted_by"`
	ClientEducationDeletedAt      *time.Time `gorm:"column:client_edu_deleted_at;default:CURRENT_TIMESTAMP" json:"prod_ctgry_deleted_at"`
}

type ResponseClientEducations struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    *[]ClientEducation `json:"data"`
}

type ResponseClientEducation struct {
	Status  uint64           `json:"status"`
	Message string           `json:"message"`
	Data    *ClientEducation `json:"data"`
}

type ResponseClientEducationCUD struct {
	Status  uint64 `json:"status"`
	Message string `json:"message"`
}

func (ClientEducation) TableName() string {
	return "m_client_education"
}

func (p *ClientEducation) Prepare() {
	p.ClientID = p.ClientID
	p.ClientEducationGraduationYear = html.EscapeString(strings.TrimSpace(p.ClientEducationGraduationYear))
	p.ClientEducationInstitution = html.EscapeString(strings.TrimSpace(p.ClientEducationInstitution))
	p.ClientEducationCity = html.EscapeString(strings.TrimSpace(p.ClientEducationCity))
	p.ClientEducationMajor = html.EscapeString(strings.TrimSpace(p.ClientEducationMajor))
	p.ClientEducationRemark = html.EscapeString(strings.TrimSpace(p.ClientEducationRemark))
	p.ClientEducationStatus = p.ClientEducationStatus
	p.ClientEducationCreatedAt = time.Now()
	p.ClientEducationCreatedBy = p.ClientEducationCreatedBy
	p.ClientEducationUpdatedAt = p.ClientEducationUpdatedAt
	p.ClientEducationUpdatedBy = p.ClientEducationUpdatedBy
	p.ClientEducationDeletedAt = p.ClientEducationDeletedAt
	p.ClientEducationDeletedBy = p.ClientEducationDeletedBy
}

func (p *ClientEducation) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.ClientEducationGraduationYear == "" {
			return errors.New("Required ClientEducation Code")
		}
		if p.ClientEducationInstitution == "" {
			return errors.New("Required KTP Number")
		}
		if p.ClientEducationCity == "" {
			return errors.New("Required Passport Number")
		}
		if p.ClientEducationMajor == "" {
			return errors.New("Required NPWP Number")
		}
		if p.ClientEducationRemark == "" {
			return errors.New("Required ClientEducation Name")
		}
		return nil
	}
}

func (p *ClientEducation) SaveClientEducation(db *gorm.DB) (*ClientEducation, error) {
	var err error
	err = db.Debug().Model(&ClientEducation{}).Create(&p).Error
	if err != nil {
		return &ClientEducation{}, err
	}
	return p, nil
}

func (p *ClientEducation) FindAllClientEducation(db *gorm.DB) (*[]ClientEducation, error) {
	var err error
	clienteducation := []ClientEducation{}
	err = db.Debug().Model(&ClientEducation{}).Limit(100).Find(&clienteducation).Error
	if err != nil {
		return &[]ClientEducation{}, err
	}
	return &clienteducation, nil
}

func (p *ClientEducation) FindAllClientEducationStatus(db *gorm.DB, status uint64) (*[]ClientEducation, error) {
	var err error
	clienteducation := []ClientEducation{}
	err = db.Debug().Model(&ClientEducation{}).Where("client_edu_status = ?", status).Limit(100).Find(&clienteducation).Error
	if err != nil {
		return &[]ClientEducation{}, err
	}
	return &clienteducation, nil
}

func (p *ClientEducation) FindClientEducationByID(db *gorm.DB, pid uint64) (*ClientEducation, error) {
	var err error
	err = db.Debug().Model(&ClientEducation{}).Where("client_edu_id = ?", pid).Take(&p).Error
	if err != nil {
		return &ClientEducation{}, err
	}
	return p, nil
}

func (p *ClientEducation) FindClientEducationStatusByID(db *gorm.DB, pid uint64, status uint64) (*ClientEducation, error) {
	var err error
	err = db.Debug().Model(&ClientEducation{}).Where("client_edu_id = ? AND client_edu_status = ?", pid, status).Take(&p).Error
	if err != nil {
		return &ClientEducation{}, err
	}
	return p, nil
}

func (p *ClientEducation) UpdateClientEducation(db *gorm.DB, pid uint64) (*ClientEducation, error) {
	var err error
	db = db.Debug().Model(&ClientEducation{}).Where("client_edu_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"client_edu_grad_year":   p.ClientEducationGraduationYear,
			"client_edu_institution": p.ClientEducationInstitution,
			"client_edu_city":        p.ClientEducationCity,
			"client_edu_major":       p.ClientEducationMajor,
			"client_edu_remark":      p.ClientEducationRemark,
			"client_edu_status":      p.ClientEducationStatus,
			"client_edu_updated_by":  p.ClientEducationUpdatedBy,
			"client_edu_updated_at":  time.Now(),
		},
	)
	err = db.Debug().Model(&ClientEducation{}).Where("client_edu_id = ?", pid).Error
	if err != nil {
		return &ClientEducation{}, err
	}
	return p, nil
}

func (p *ClientEducation) UpdateClientEducationStatus(db *gorm.DB, pid uint64) (*ClientEducation, error) {
	var err error
	db = db.Debug().Model(&ClientEducation{}).Where("client_edu_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"client_edu_status":     p.ClientEducationStatus,
			"client_edu_updated_by": p.ClientEducationUpdatedBy,
			"client_edu_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ClientEducation{}).Where("client_edu_id = ?", pid).Error
	if err != nil {
		return &ClientEducation{}, err
	}
	return p, nil
}

func (p *ClientEducation) UpdateClientEducationStatusDelete(db *gorm.DB, pid uint64) (*ClientEducation, error) {
	var err error
	db = db.Debug().Model(&ClientEducation{}).Where("client_edu_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"client_edu_status":     3,
			"client_edu_deleted_by": p.ClientEducationDeletedBy,
			"client_edu_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ClientEducation{}).Where("client_edu_id = ?", pid).Error
	if err != nil {
		return &ClientEducation{}, err
	}
	return p, nil
}

func (p *ClientEducation) DeleteClientEducation(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&ClientEducation{}).Where("client_edu_id = ?", pid).Take(&ClientEducation{}).Delete(&ClientEducation{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("ClientEducation not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// ADDITIONAL MEMORANDUM
// ===================================================================================================================================

type ClientEducationMrdmData struct {
	ClientEducationID             uint64 `gorm:"column:client_edu_id" json:"client_edu_id"`
	ClientID                      uint64 `gorm:"column:client_id" json:"client_id"`
	ClientEducationGraduationYear string `gorm:"column:client_edu_grad_year" json:"client_edu_grad_year"`
	ClientEducationInstitution    string `gorm:"column:client_edu_institution" json:"client_edu_institution"`
	ClientEducationCity           string `gorm:"column:client_edu_city" json:"client_edu_city"`
	ClientEducationMajor          string `gorm:"column:client_edu_major" json:"client_edu_major"`
	ClientEducationRemark         string `gorm:"column:client_edu_remark" json:"client_edu_remark"`
}

func (ClientEducationMrdmData) TableName() string {
	return "m_client_education"
}

func (p *ClientEducation) FindClientEducationMrdmByClientID(db *gorm.DB, pid uint64) ([]ClientEducationMrdmData, error) {
	var err error
	clienteducation := []ClientEducationMrdmData{}
	err = db.Debug().Model(&ClientEducationMrdmData{}).Where("client_id = ? AND client_edu_status = ?", pid, 1).Limit(100).Find(&clienteducation).Error
	if err != nil {
		return []ClientEducationMrdmData{}, err
	}
	return clienteducation, nil
}

func (p *ClientEducation) GetClientEducationByClientID(db *gorm.DB, pid uint64) (*[]ClientEducation, error) {
	var err error
	clienteducation := []ClientEducation{}
	err = db.Debug().Model(&ClientEducation{}).
		Where("client_id = ? AND client_edu_status = ?", pid, 1).
		Find(&clienteducation).Error
	if err != nil {
		return &[]ClientEducation{}, err
	}
	return &clienteducation, nil
}
