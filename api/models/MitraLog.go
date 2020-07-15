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

type MitraLog struct {
	MitraLogID          uint64    `gorm:"column:mitra_log_id;primary_key;" json:"mitra_log_id"`
	MitraUserID         uint64    `gorm:"column:mitra_user_id" json:"mitra_user_id"`
	MitraLogType        string    `gorm:"column:mitra_log_type;size:3" json:"mitra_log_type"`
	MitraLogAction      string    `gorm:"column:mitra_log_action;size:255" json:"mitra_log_action"`
	MitraLogDescription string    `gorm:"column:mitra_log_description" json:"mitra_log_description"`
	MitraLogCreatedAt   time.Time `gorm:"column:mitra_log_created_at;default:CURRENT_TIMESTAMP" json:"mitra_log_created_at"`
}

type MitraLogData struct {
	MitraLogID          uint64    `gorm:"column:mitra_log_id;primary_key;" json:"mitra_log_id"`
	MitraUserID         uint64    `gorm:"column:mitra_user_id" json:"mitra_user_id"`
	MitraEmployeeName   string    `gorm:"column:mitra_emp_name" json:"mitra_emp_name"`
	MitraLogType        string    `gorm:"column:mitra_log_type" json:"mitra_log_type"`
	MitraLogAction      string    `gorm:"column:mitra_log_action" json:"mitra_log_action"`
	MitraLogDescription string    `gorm:"column:mitra_log_description" json:"mitra_log_description"`
	MitraLogCreatedAt   time.Time `gorm:"column:mitra_log_created_at" json:"mitra_log_created_at"`
}

type ResponseMitraLogs struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Data    *[]MitraLogData `json:"data"`
}

type ResponseMitraLog struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    *MitraLogData `json:"data"`
}

type ResponseMitraLogCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraLog) TableName() string {
	return "t_mitra_log"
}
func (MitraLogData) TableName() string {
	return "t_mitra_log"
}

func (p *MitraLog) Prepare() {
	p.MitraUserID = p.MitraUserID
	p.MitraLogType = html.EscapeString(strings.TrimSpace(p.MitraLogType))
	p.MitraLogAction = html.EscapeString(strings.TrimSpace(p.MitraLogAction))
	p.MitraLogDescription = p.MitraLogDescription
	p.MitraLogCreatedAt = time.Now()
}

func (p *MitraLog) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.MitraUserID == 0 {
			return errors.New("Required Mitra User ID")
		}
		if p.MitraLogType == "" {
			return errors.New("Required Mitra User Log Type")
		}
		if p.MitraLogAction == "" {
			return errors.New("Required Mitra User Log Action")
		}
		if p.MitraLogDescription == "" {
			return errors.New("Required Mitra User Log Description")
		}
		return nil
	}
}

func (p *MitraLog) SaveMitraLog(db *gorm.DB) (*MitraLog, error) {
	var err error
	err = db.Debug().Model(&MitraLog{}).Create(&p).Error
	if err != nil {
		return &MitraLog{}, err
	}
	return p, nil
}

func (p *MitraLog) FindAllMitraLogs(db *gorm.DB) (*[]MitraLogData, error) {
	var err error
	mitrauserlog := []MitraLogData{}
	err = db.Debug().Model(&MitraLogData{}).Limit(100).
		Select(`t_mitra_log.mitra_log_id,
				t_mitra_log.mitra_user_id,
				m_mitra_employee.mitra_emp_name,
				t_mitra_log.mitra_log_type,
				t_mitra_log.mitra_log_action,
				t_mitra_log.mitra_log_description,
				t_mitra_log.mitra_log_created_at at time zone 'Asia/Jakarta' as mitra_log_created_at`).
		Joins("JOIN m_mitra_user on t_mitra_log.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra_employee on m_mitra_user.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Find(&mitrauserlog).Error
	if err != nil {
		return &[]MitraLogData{}, err
	}
	return &mitrauserlog, nil
}

func (p *MitraLog) FindMitraLogByID(db *gorm.DB, pid uint64) (*MitraLogData, error) {
	var err error
	mitrauserlog := MitraLogData{}
	err = db.Debug().Model(&MitraLogData{}).Limit(100).
		Select(`t_mitra_log.mitra_log_id,
				t_mitra_log.mitra_user_id,
				m_mitra_employee.mitra_emp_name,
				t_mitra_log.mitra_log_type,
				t_mitra_log.mitra_log_action,
				t_mitra_log.mitra_log_description,
				t_mitra_log.mitra_log_created_at at time zone 'Asia/Jakarta' as mitra_log_created_at`).
		Joins("JOIN m_mitra_user on t_mitra_log.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra_employee on m_mitra_user.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Where("mitra_log_id = ?", pid).
		Find(&mitrauserlog).Error
	if err != nil {
		return &MitraLogData{}, err
	}
	return &mitrauserlog, nil
}

func (p *MitraLog) UpdateMitraLog(db *gorm.DB, pid uint64) (*MitraLog, error) {

	var err error
	db = db.Debug().Model(&MitraLog{}).Where("mitra_log_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_user_id":         p.MitraUserID,
			"mitra_log_type":        p.MitraLogType,
			"mitra_log_action":      p.MitraLogAction,
			"mitra_log_description": p.MitraLogDescription,
		},
	)
	err = db.Debug().Model(&MitraLog{}).Where("mitra_log_id = ?", pid).Error
	if err != nil {
		return &MitraLog{}, err
	}
	return p, nil
}

func (p *MitraLog) DeleteMitraLog(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&MitraLog{}).Where("mitra_log_id = ?", pid).Take(&MitraLog{}).Delete(&MitraLog{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraLog not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

func (p *MitraLog) FindMitraLogUserByMitraUserID(db *gorm.DB, pid uint64) (*MitraLogData, error) {
	var err error
	mitrauserlog := MitraLogData{}
	err = db.Debug().Model(&MitraLogData{}).Limit(100).
		Select(`t_mitra_log.mitra_log_id,
				t_mitra_log.mitra_user_id,
				m_mitra_employee.mitra_emp_name,
				t_mitra_log.mitra_log_type,
				t_mitra_log.mitra_log_action,
				t_mitra_log.mitra_log_description,
				t_mitra_log.mitra_log_created_at at time zone 'Asia/Jakarta' as mitra_log_created_at`).
		Joins("JOIN m_mitra_user on t_mitra_log.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra_employee on m_mitra_user.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Where("t_mitra_log.mitra_user_id = ?", pid).
		Find(&mitrauserlog).Error
	if err != nil {
		return &MitraLogData{}, err
	}
	return &mitrauserlog, nil
}

func (p *MitraLog) FindMitraLogAdminByMitraUserID(db *gorm.DB, pid uint64) (*MitraLogData, error) {
	var err error
	mitrauserlog := MitraLogData{}
	err = db.Debug().Model(&MitraLogData{}).Limit(100).
		Select(`t_mitra_log.mitra_log_id,
				t_mitra_log.mitra_user_id,
				m_mitra_employee.mitra_emp_name,
				t_mitra_log.mitra_log_type,
				t_mitra_log.mitra_log_action,
				t_mitra_log.mitra_log_description,
				t_mitra_log.mitra_log_created_at at time zone 'Asia/Jakarta' as mitra_log_created_at`).
		Joins("JOIN m_mitra_administrator on t_mitra_log.mitra_user_id=m_mitra_administrator.mitra_adm_id").
		Joins("JOIN m_mitra_employee on m_mitra_administrator.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Where("t_mitra_log.mitra_user_id = ?", pid).
		Find(&mitrauserlog).Error
	if err != nil {
		return &MitraLogData{}, err
	}
	return &mitrauserlog, nil
}
