package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type ClientWorkSupervisor struct {
	ClientWorkSupervisorID          uint64     `gorm:"column:client_wspv_id;primary_key;" json:"client_wspv_id"`
	ClientID                        uint64     `gorm:"column:client_id" json:"client_id"`
	ClientWorkSupervisorName        string     `gorm:"column:client_wspv_name;size:125" json:"client_wspv_name"`
	ClientWorkSupervisorJobTitle    string     `gorm:"column:client_wspv_job_title;size:125" json:"client_wspv_job_title"`
	ClientWorkSupervisorHrdName     string     `gorm:"column:client_wspv_hrd_name;size:125" json:"client_wspv_hrd_name"`
	ClientWorkSupervisorHrdJobTitle string     `gorm:"column:client_wspv_hrd_job_title;size:125" json:"client_wspv_hrd_job_title"`
	ClientWorkSupervisorStatus      uint64     `gorm:"column:client_wspv_status;size:2" json:"client_wspv_status"`
	ClientWorkSupervisorCreatedBy   string     `gorm:"column:client_wspv_created_by;size:125" json:"client_wspv_created_by"`
	ClientWorkSupervisorCreatedAt   time.Time  `gorm:"column:client_wspv_created_at;default:CURRENT_TIMESTAMP" json:"client_wspv_created_at"`
	ClientWorkSupervisorUpdatedBy   string     `gorm:"column:client_wspv_updated_by;size:125" json:"client_wspv_updated_by"`
	ClientWorkSupervisorUpdatedAt   *time.Time `gorm:"column:client_wspv_updated_at;default:CURRENT_TIMESTAMP" json:"client_wspv_updated_at"`
	ClientWorkSupervisorDeletedBy   string     `gorm:"column:client_wspv_deleted_by;size:125" json:"client_wspv_deleted_by"`
	ClientWorkSupervisorDeletedAt   *time.Time `gorm:"column:client_wspv_deleted_at;default:CURRENT_TIMESTAMP" json:"client_wspv_deleted_at"`
}

type ResponseClientWorkSupervisors struct {
	Status  int                     `json:"status"`
	Message string                  `json:"message"`
	Data    *[]ClientWorkSupervisor `json:"data"`
}

type ResponseClientWorkSupervisor struct {
	Status  int                   `json:"status"`
	Message string                `json:"message"`
	Data    *ClientWorkSupervisor `json:"data"`
}

type ResponseClientWorkSupervisorCUD struct {
	Status  uint64 `json:"status"`
	Message string `json:"message"`
}

func (ClientWorkSupervisor) TableName() string {
	return "m_client_workspv"
}

func (p *ClientWorkSupervisor) Prepare() {
	p.ClientID = p.ClientID
	p.ClientWorkSupervisorName = html.EscapeString(strings.TrimSpace(p.ClientWorkSupervisorName))
	p.ClientWorkSupervisorJobTitle = html.EscapeString(strings.TrimSpace(p.ClientWorkSupervisorJobTitle))
	p.ClientWorkSupervisorHrdName = html.EscapeString(strings.TrimSpace(p.ClientWorkSupervisorHrdName))
	p.ClientWorkSupervisorHrdJobTitle = html.EscapeString(strings.TrimSpace(p.ClientWorkSupervisorHrdJobTitle))
	p.ClientWorkSupervisorStatus = p.ClientWorkSupervisorStatus
	p.ClientWorkSupervisorCreatedAt = time.Now()
	p.ClientWorkSupervisorCreatedBy = p.ClientWorkSupervisorCreatedBy
	p.ClientWorkSupervisorUpdatedAt = p.ClientWorkSupervisorUpdatedAt
	p.ClientWorkSupervisorUpdatedBy = p.ClientWorkSupervisorUpdatedBy
	p.ClientWorkSupervisorDeletedAt = p.ClientWorkSupervisorDeletedAt
	p.ClientWorkSupervisorDeletedAt = p.ClientWorkSupervisorDeletedAt
}

func (p *ClientWorkSupervisor) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.ClientWorkSupervisorName == "" {
			return errors.New("Required ClientWorkSupervisor Code")
		}
		if p.ClientWorkSupervisorJobTitle == "" {
			return errors.New("Required KTP Number")
		}
		if p.ClientWorkSupervisorHrdName == "" {
			return errors.New("Required Passport Number")
		}
		if p.ClientWorkSupervisorHrdJobTitle == "" {
			return errors.New("Required Passport Number")
		}
		return nil
	}
}

func (p *ClientWorkSupervisor) SaveClientWorkSupervisor(db *gorm.DB) (*ClientWorkSupervisor, error) {
	var err error
	err = db.Debug().Model(&ClientWorkSupervisor{}).Create(&p).Error
	if err != nil {
		return &ClientWorkSupervisor{}, err
	}
	return p, nil
}

func (p *ClientWorkSupervisor) FindAllClientWorkSupervisor(db *gorm.DB) (*[]ClientWorkSupervisor, error) {
	var err error
	clientworkspv := []ClientWorkSupervisor{}
	err = db.Debug().Model(&ClientWorkSupervisor{}).Limit(100).Find(&clientworkspv).Error
	if err != nil {
		return &[]ClientWorkSupervisor{}, err
	}
	return &clientworkspv, nil
}

func (p *ClientWorkSupervisor) FindAllClientWorkSupervisorStatus(db *gorm.DB, status uint64) (*[]ClientWorkSupervisor, error) {
	var err error
	clientworkspv := []ClientWorkSupervisor{}
	err = db.Debug().Model(&ClientWorkSupervisor{}).Where("client_wspv_status = ?", status).Limit(100).Find(&clientworkspv).Error
	if err != nil {
		return &[]ClientWorkSupervisor{}, err
	}
	return &clientworkspv, nil
}

func (p *ClientWorkSupervisor) FindClientWorkSupervisorByID(db *gorm.DB, pid uint64) (*ClientWorkSupervisor, error) {
	var err error
	err = db.Debug().Model(&ClientWorkSupervisor{}).Where("client_wspv_id = ?", pid).Take(&p).Error
	if err != nil {
		return &ClientWorkSupervisor{}, err
	}
	return p, nil
}

func (p *ClientWorkSupervisor) FindClientWorkSupervisorStatusByID(db *gorm.DB, pid uint64, status uint64) (*ClientWorkSupervisor, error) {
	var err error
	err = db.Debug().Model(&ClientWorkSupervisor{}).Where("client_wspv_id = ? AND client_wspv_status = ?", pid, status).Take(&p).Error
	if err != nil {
		return &ClientWorkSupervisor{}, err
	}
	return p, nil
}

func (p *ClientWorkSupervisor) UpdateClientWorkSupervisor(db *gorm.DB, pid uint64) (*ClientWorkSupervisor, error) {
	var err error
	db = db.Debug().Model(&ClientWorkSupervisor{}).Where("client_wspv_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"client_wspv_name":          p.ClientWorkSupervisorName,
			"client_wspv_job_title":     p.ClientWorkSupervisorJobTitle,
			"client_wspv_hrd_name":      p.ClientWorkSupervisorHrdName,
			"client_wspv_hrd_job_title": p.ClientWorkSupervisorHrdJobTitle,
			"client_wspv_status":        p.ClientWorkSupervisorStatus,
			"client_wspv_updated_by":    p.ClientWorkSupervisorUpdatedBy,
			"client_wspv_updated_at":    time.Now(),
		},
	)
	err = db.Debug().Model(&ClientWorkSupervisor{}).Where("client_wspv_id = ?", pid).Error
	if err != nil {
		return &ClientWorkSupervisor{}, err
	}
	return p, nil
}

func (p *ClientWorkSupervisor) UpdateClientWorkSupervisorStatus(db *gorm.DB, pid uint64) (*ClientWorkSupervisor, error) {
	var err error
	db = db.Debug().Model(&ClientWorkSupervisor{}).Where("client_wspv_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"client_wspv_status":     p.ClientWorkSupervisorStatus,
			"client_wspv_updated_by": p.ClientWorkSupervisorUpdatedBy,
			"client_wspv_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ClientWorkSupervisor{}).Where("client_wspv_id = ?", pid).Error
	if err != nil {
		return &ClientWorkSupervisor{}, err
	}
	return p, nil
}

func (p *ClientWorkSupervisor) UpdateClientWorkSupervisorStatusDelete(db *gorm.DB, pid uint64) (*ClientWorkSupervisor, error) {
	var err error
	db = db.Debug().Model(&ClientWorkSupervisor{}).Where("client_wspv_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"client_wspv_status":     3,
			"client_wspv_deleted_by": p.ClientWorkSupervisorDeletedBy,
			"client_wspv_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&ClientWorkSupervisor{}).Where("client_wspv_id = ?", pid).Error
	if err != nil {
		return &ClientWorkSupervisor{}, err
	}
	return p, nil
}

func (p *ClientWorkSupervisor) DeleteClientWorkSupervisor(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&ClientWorkSupervisor{}).Where("client_wspv_id = ?", pid).Take(&ClientWorkSupervisor{}).Delete(&ClientWorkSupervisor{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("ClientWorkSupervisor not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// ADDITIONAL MEMORANDUM
// ===================================================================================================================================

type ClientWorkSupervisorMrdmData struct {
	ClientWorkSupervisorID          uint64 `gorm:"column:client_wspv_id" json:"client_wspv_id"`
	ClientID                        uint64 `gorm:"column:client_id" json:"client_id"`
	ClientWorkSupervisorName        string `gorm:"column:client_wspv_name" json:"client_wspv_name"`
	ClientWorkSupervisorJobTitle    string `gorm:"column:client_wspv_job_title" json:"client_wspv_job_title"`
	ClientWorkSupervisorHrdName     string `gorm:"column:client_wspv_hrd_name" json:"client_wspv_hrd_name"`
	ClientWorkSupervisorHrdJobTitle string `gorm:"column:client_wspv_hrd_job_title" json:"client_wspv_hrd_job_title"`
}

func (ClientWorkSupervisorMrdmData) TableName() string {
	return "m_client_workspv"
}

func (p *ClientWorkSupervisor) FindClientWorkSupervisorMrdmByClientID(db *gorm.DB, pid uint64) ([]ClientWorkSupervisorMrdmData, error) {
	var err error
	clientworkspv := []ClientWorkSupervisorMrdmData{}
	err = db.Debug().Model(&ClientWorkSupervisorMrdmData{}).Where("client_id = ? AND client_wspv_status = ?", pid, 1).Limit(100).Find(&clientworkspv).Error
	if err != nil {
		return []ClientWorkSupervisorMrdmData{}, err
	}
	return clientworkspv, nil
}

func (p *ClientWorkSupervisor) GetClientWorkSupervisorByClientID(db *gorm.DB, pid uint64) (*[]ClientWorkSupervisor, error) {
	var err error
	clientworksupervisor := []ClientWorkSupervisor{}
	err = db.Debug().Model(&ClientWorkSupervisor{}).
		Where("client_id = ? AND client_wexp_status = ?", pid, 1).
		Find(&clientworksupervisor).Error
	if err != nil {
		return &[]ClientWorkSupervisor{}, err
	}
	return &clientworksupervisor, nil
}
