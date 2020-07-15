package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type ClientWorkExperience struct {
	ClientWorkExperienceID        uint64     `gorm:"column:client_wexp_id;primary_key;" json:"client_wexp_id"`
	ClientID                      uint64     `gorm:"column:client_id" json:"client_id"`
	ClientWorkExperienceYear      string     `gorm:"column:client_wexp_year;size:10" json:"client_wexp_year"`
	ClientWorkExperienceCompany   string     `gorm:"column:client_wexp_company;size:255" json:"client_wexp_company"`
	ClientWorkExperienceSalary    string     `gorm:"column:client_wexp_salary;size:15" json:"client_wexp_salary"`
	ClientWorkExperienceJobTitle  string     `gorm:"column:client_wexp_job_title;size:125" json:"client_wexp_job_title"`
	ClientWorkExperienceLongWork  string     `gorm:"column:client_wexp_long_work;size:25" json:"client_wexp_long_work"`
	ClientWorkExperienceStatus    uint64     `gorm:"column:client_wexp_status;size:2" json:"client_wexp_status"`
	ClientWorkExperienceCreatedBy string     `gorm:"column:client_wexp_created_by;size:125" json:"client_wexp_created_by"`
	ClientWorkExperienceCreatedAt time.Time  `gorm:"column:client_wexp_created_at;default:CURRENT_TIMESTAMP" json:"prod_ctgry_created_at"`
	ClientWorkExperienceUpdatedBy string     `gorm:"column:client_wexp_updated_by;size:125" json:"client_wexp_updated_by"`
	ClientWorkExperienceUpdatedAt *time.Time `gorm:"column:client_wexp_updated_at;default:CURRENT_TIMESTAMP" json:"prod_ctgry_updated_at"`
	ClientWorkExperienceDeletedBy string     `gorm:"column:client_wexp_deleted_by;size:125" json:"client_wexp_deleted_by"`
	ClientWorkExperienceDeletedAt *time.Time `gorm:"column:client_wexp_deleted_at;default:CURRENT_TIMESTAMP" json:"prod_ctgry_deleted_at"`
}

type ResponseClientWorkExperiences struct {
	Status  int                     `json:"status"`
	Message string                  `json:"message"`
	Data    *[]ClientWorkExperience `json:"data"`
}

type ResponseClientWorkExperience struct {
	Status  int                   `json:"status"`
	Message string                `json:"message"`
	Data    *ClientWorkExperience `json:"data"`
}

type ResponseClientWorkExperienceCUD struct {
	Status  uint64 `json:"status"`
	Message string `json:"message"`
}

func (ClientWorkExperience) TableName() string {
	return "m_client_workexp"
}

func (p *ClientWorkExperience) Prepare() {
	p.ClientID = p.ClientID
	p.ClientWorkExperienceYear = html.EscapeString(strings.TrimSpace(p.ClientWorkExperienceYear))
	p.ClientWorkExperienceCompany = html.EscapeString(strings.TrimSpace(p.ClientWorkExperienceCompany))
	p.ClientWorkExperienceSalary = html.EscapeString(strings.TrimSpace(p.ClientWorkExperienceSalary))
	p.ClientWorkExperienceJobTitle = html.EscapeString(strings.TrimSpace(p.ClientWorkExperienceJobTitle))
	p.ClientWorkExperienceLongWork = html.EscapeString(strings.TrimSpace(p.ClientWorkExperienceLongWork))
	p.ClientWorkExperienceStatus = p.ClientWorkExperienceStatus
	p.ClientWorkExperienceCreatedAt = time.Now()
	p.ClientWorkExperienceCreatedBy = p.ClientWorkExperienceCreatedBy
	p.ClientWorkExperienceUpdatedAt = p.ClientWorkExperienceUpdatedAt
	p.ClientWorkExperienceUpdatedBy = p.ClientWorkExperienceUpdatedBy
	p.ClientWorkExperienceDeletedAt = p.ClientWorkExperienceDeletedAt
	p.ClientWorkExperienceDeletedBy = p.ClientWorkExperienceDeletedBy
}

func (p *ClientWorkExperience) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.ClientWorkExperienceYear == "" {
			return errors.New("Required ClientWorkExperience Code")
		}
		if p.ClientWorkExperienceCompany == "" {
			return errors.New("Required KTP Number")
		}
		if p.ClientWorkExperienceSalary == "" {
			return errors.New("Required Passport Number")
		}
		if p.ClientWorkExperienceJobTitle == "" {
			return errors.New("Required Passport Number")
		}
		if p.ClientWorkExperienceLongWork == "" {
			return errors.New("Required Passport Number")
		}
		return nil
	}
}

func (p *ClientWorkExperience) SaveClientWorkExperience(db *gorm.DB) (*ClientWorkExperience, error) {
	var err error
	err = db.Debug().Model(&ClientWorkExperience{}).Create(&p).Error
	if err != nil {
		return &ClientWorkExperience{}, err
	}
	return p, nil
}

func (p *ClientWorkExperience) FindAllClientWorkExperience(db *gorm.DB) (*[]ClientWorkExperience, error) {
	var err error
	clientworkexp := []ClientWorkExperience{}
	err = db.Debug().Model(&ClientWorkExperience{}).Limit(100).Find(&clientworkexp).Error
	if err != nil {
		return &[]ClientWorkExperience{}, err
	}
	return &clientworkexp, nil
}

func (p *ClientWorkExperience) FindAllClientWorkExperienceStatus(db *gorm.DB, status uint64) (*[]ClientWorkExperience, error) {
	var err error
	clientworkexp := []ClientWorkExperience{}
	err = db.Debug().Model(&ClientWorkExperience{}).Where("client_wexp_status = ?", status).Limit(100).Find(&clientworkexp).Error
	if err != nil {
		return &[]ClientWorkExperience{}, err
	}
	return &clientworkexp, nil
}

func (p *ClientWorkExperience) FindClientWorkExperienceByID(db *gorm.DB, pid uint64) (*ClientWorkExperience, error) {
	var err error
	err = db.Debug().Model(&ClientWorkExperience{}).Where("client_wexp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &ClientWorkExperience{}, err
	}
	return p, nil
}

func (p *ClientWorkExperience) FindClientWorkExperienceStatusByID(db *gorm.DB, pid uint64, status uint64) (*ClientWorkExperience, error) {
	var err error
	err = db.Debug().Model(&ClientWorkExperience{}).Where("client_wexp_id = ? AND client_wexp_status = ?", pid, status).Take(&p).Error
	if err != nil {
		return &ClientWorkExperience{}, err
	}
	return p, nil
}

func (p *ClientWorkExperience) UpdateClientWorkExperience(db *gorm.DB, pid uint64) (*ClientWorkExperience, error) {
	var err error
	db = db.Debug().Model(&ClientWorkExperience{}).Where("client_wexp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"client_wexp_year":       p.ClientWorkExperienceYear,
			"client_wexp_company":    p.ClientWorkExperienceCompany,
			"client_wexp_salary":     p.ClientWorkExperienceSalary,
			"client_wexp_job_title":  p.ClientWorkExperienceJobTitle,
			"client_wexp_long_work":  p.ClientWorkExperienceLongWork,
			"client_wexp_updated_by": p.ClientWorkExperienceUpdatedBy,
			"client_wexp_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ClientWorkExperience{}).Where("client_wexp_id = ?", pid).Error
	if err != nil {
		return &ClientWorkExperience{}, err
	}
	return p, nil
}

func (p *ClientWorkExperience) UpdateClientWorkExperienceStatus(db *gorm.DB, pid uint64) (*ClientWorkExperience, error) {
	var err error
	db = db.Debug().Model(&ClientWorkExperience{}).Where("client_wexp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"client_wexp_status":     p.ClientWorkExperienceStatus,
			"client_wexp_updated_by": p.ClientWorkExperienceUpdatedBy,
			"client_wexp_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ClientWorkExperience{}).Where("client_wexp_id = ?", pid).Error
	if err != nil {
		return &ClientWorkExperience{}, err
	}
	return p, nil
}

func (p *ClientWorkExperience) UpdateClientWorkExperienceStatusDelete(db *gorm.DB, pid uint64) (*ClientWorkExperience, error) {
	var err error
	db = db.Debug().Model(&ClientWorkExperience{}).Where("client_wexp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"client_wexp_status":     3,
			"client_wexp_deleted_by": p.ClientWorkExperienceDeletedBy,
			"client_wexp_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ClientWorkExperience{}).Where("client_wexp_id = ?", pid).Error
	if err != nil {
		return &ClientWorkExperience{}, err
	}
	return p, nil
}

func (p *ClientWorkExperience) DeleteClientWorkExperience(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&ClientWorkExperience{}).Where("client_wexp_id = ?", pid).Take(&ClientWorkExperience{}).Delete(&ClientWorkExperience{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("ClientWorkExperience not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// ADDITIONAL MEMORANDUM
// ===================================================================================================================================

type ClientWorkExperienceMrdmData struct {
	ClientWorkExperienceID       uint64 `gorm:"column:client_wexp_id" json:"client_wexp_id"`
	ClientID                     uint64 `gorm:"column:client_id" json:"client_id"`
	ClientWorkExperienceYear     string `gorm:"column:client_wexp_year" json:"client_wexp_year"`
	ClientWorkExperienceCompany  string `gorm:"column:client_wexp_company" json:"client_wexp_company"`
	ClientWorkExperienceSalary   string `gorm:"column:client_wexp_salary" json:"client_wexp_salary"`
	ClientWorkExperienceJobTitle string `gorm:"column:client_wexp_job_title" json:"client_wexp_job_title"`
	ClientWorkExperienceLongWork string `gorm:"column:client_wexp_long_work" json:"client_wexp_long_work"`
}

func (ClientWorkExperienceMrdmData) TableName() string {
	return "m_client_workexp"
}

func (p *ClientWorkExperience) FindClientWorkExperienceMrdmByClientId(db *gorm.DB, pid uint64) ([]ClientWorkExperienceMrdmData, error) {
	var err error
	clientworkexp := []ClientWorkExperienceMrdmData{}
	err = db.Debug().Model(&ClientWorkExperienceMrdmData{}).Where("client_id = ? AND client_wexp_status = ?", pid, 1).Limit(100).Find(&clientworkexp).Error
	if err != nil {
		return []ClientWorkExperienceMrdmData{}, err
	}
	return clientworkexp, nil
}

func (p *ClientWorkExperience) GetClientWorkExperienceByClientID(db *gorm.DB, pid uint64) (*[]ClientWorkExperience, error) {
	var err error
	clientworkexperience := []ClientWorkExperience{}
	err = db.Debug().Model(&ClientWorkExperience{}).
		Where("client_id = ? AND client_wexp_status = ?", pid, 1).
		Find(&clientworkexperience).Error
	if err != nil {
		return &[]ClientWorkExperience{}, err
	}
	return &clientworkexperience, nil
}
