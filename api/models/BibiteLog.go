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

type BibiteLog struct {
	BibiteLogID          uint64    `gorm:"column:bbt_log_id;primary_key;" json:"bbt_log_id"`
	BibiteUserID         uint64    `gorm:"column:bbt_user_id" json:"bbt_user_id"`
	BibiteLogType        string    `gorm:"column:bbt_log_type;size:3" json:"bbt_log_type"`
	BibiteLogAction      string    `gorm:"column:bbt_log_action;size:255" json:"bbt_log_action"`
	BibiteLogDescription string    `gorm:"column:bbt_log_description" json:"bbt_log_description"`
	BibiteLogCreatedAt   time.Time `gorm:"column:bbt_log_created_at;default:CURRENT_TIMESTAMP" json:"bbt_log_created_at"`
}

type BibiteLogData struct {
	BibiteLogID          uint64    `gorm:"column:bbt_log_id;primary_key;" json:"bbt_log_id"`
	BibiteUserID         uint64    `gorm:"column:bbt_user_id" json:"bbt_user_id"`
	BibiteEmployeeName   string    `gorm:"column:bbt_emp_name" json:"bbt_emp_name"`
	BibiteLogType        string    `gorm:"column:bbt_log_type" json:"bbt_log_type"`
	BibiteLogAction      string    `gorm:"column:bbt_log_action" json:"bbt_log_action"`
	BibiteLogDescription string    `gorm:"column:bbt_log_description" json:"bbt_log_description"`
	BibiteLogCreatedAt   time.Time `gorm:"column:bbt_log_created_at" json:"bbt_log_created_at"`
}

type ResponseBibiteLogs struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    *[]BibiteLogData `json:"data"`
}

type ResponseBibiteLog struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    *BibiteLogData `json:"data"`
}

type ResponseBibiteLogCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (BibiteLog) TableName() string {
	return "t_bibite_log"
}
func (BibiteLogData) TableName() string {
	return "t_bibite_log"
}

func (p *BibiteLog) Prepare() {
	p.BibiteUserID = p.BibiteUserID
	p.BibiteLogType = html.EscapeString(strings.TrimSpace(p.BibiteLogType))
	p.BibiteLogAction = html.EscapeString(strings.TrimSpace(p.BibiteLogAction))
	p.BibiteLogDescription = p.BibiteLogDescription
	p.BibiteLogCreatedAt = time.Now()
}

func (p *BibiteLog) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.BibiteUserID == 0 {
			return errors.New("Required Bibite User ID")
		}
		if p.BibiteLogType == "" {
			return errors.New("Required Bibite User Log Type")
		}
		if p.BibiteLogAction == "" {
			return errors.New("Required Bibite User Log Action")
		}
		if p.BibiteLogDescription == "" {
			return errors.New("Required Bibite User Log Description")
		}
		return nil
	}
}

func (p *BibiteLog) SaveBibiteLog(db *gorm.DB) (*BibiteLog, error) {
	var err error
	err = db.Debug().Model(&BibiteLog{}).Create(&p).Error
	if err != nil {
		return &BibiteLog{}, err
	}
	return p, nil
}

func (p *BibiteLog) FindAllBibiteLogs(db *gorm.DB) (*[]BibiteLogData, error) {
	var err error
	bibiteuserlog := []BibiteLogData{}
	err = db.Debug().Model(&BibiteLogData{}).Limit(100).
		Select(`t_bibite_log.bbt_log_id,
				t_bibite_log.bbt_user_id,
				m_bibite_employee.bbt_emp_name,
				t_bibite_log.bbt_log_type,
				t_bibite_log.bbt_log_action,
				t_bibite_log.bbt_log_description,
				t_bibite_log.bbt_log_created_at at time zone 'Asia/Jakarta' as bbt_log_created_at`).
		Joins("JOIN m_bibite_user on t_bibite_log.bbt_user_id=m_bibite_user.bbt_user_id").
		Joins("JOIN m_bibite_employee on m_bibite_user.bbt_emp_id=m_bibite_employee.bbt_emp_id").
		Find(&bibiteuserlog).Error
	if err != nil {
		return &[]BibiteLogData{}, err
	}
	return &bibiteuserlog, nil
}

func (p *BibiteLog) FindBibiteLogByID(db *gorm.DB, pid uint64) (*BibiteLogData, error) {
	var err error
	bibiteuserlog := BibiteLogData{}
	err = db.Debug().Model(&BibiteLogData{}).Limit(100).
		Select(`t_bibite_log.bbt_log_id,
				t_bibite_log.bbt_user_id,
				m_bibite_employee.bbt_emp_name,
				t_bibite_log.bbt_log_type,
				t_bibite_log.bbt_log_action,
				t_bibite_log.bbt_log_description,
				t_bibite_log.bbt_log_created_at at time zone 'Asia/Jakarta' as bbt_log_created_at`).
		Joins("JOIN m_bibite_user on t_bibite_log.bbt_user_id=m_bibite_user.bbt_user_id").
		Joins("JOIN m_bibite_employee on m_bibite_user.bbt_emp_id=m_bibite_employee.bbt_emp_id").
		Where("bbt_log_id = ?", pid).
		Find(&bibiteuserlog).Error
	if err != nil {
		return &BibiteLogData{}, err
	}
	return &bibiteuserlog, nil
}

func (p *BibiteLog) UpdateBibiteLog(db *gorm.DB, pid uint64) (*BibiteLog, error) {

	var err error
	db = db.Debug().Model(&BibiteLog{}).Where("bbt_log_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"bbt_user_id":         p.BibiteUserID,
			"bbt_log_type":        p.BibiteLogType,
			"bbt_log_action":      p.BibiteLogAction,
			"bbt_log_description": p.BibiteLogDescription,
		},
	)
	err = db.Debug().Model(&BibiteLog{}).Where("bbt_log_id = ?", pid).Error
	if err != nil {
		return &BibiteLog{}, err
	}
	return p, nil
}

func (p *BibiteLog) DeleteBibiteLog(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&BibiteLog{}).Where("bbt_log_id = ?", pid).Take(&BibiteLog{}).Delete(&BibiteLog{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("BibiteLog not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

func (p *BibiteLog) FindBibiteLogUserByBibiteUserID(db *gorm.DB, pid uint64) (*BibiteLogData, error) {
	var err error
	bibiteuserlog := BibiteLogData{}
	err = db.Debug().Model(&BibiteLogData{}).Limit(100).
		Select(`t_bibite_log.bbt_log_id,
				t_bibite_log.bbt_user_id,
				m_bibite_employee.bbt_emp_name,
				t_bibite_log.bbt_log_type,
				t_bibite_log.bbt_log_action,
				t_bibite_log.bbt_log_description,
				t_bibite_log.bbt_log_created_at at time zone 'Asia/Jakarta' as bbt_log_created_at`).
		Joins("JOIN m_bibite_user on t_bibite_log.bbt_user_id=m_bibite_user.bbt_user_id").
		Joins("JOIN m_bibite_employee on m_bibite_user.bbt_emp_id=m_bibite_employee.bbt_emp_id").
		Where("t_bibite_log.bbt_user_id = ?", pid).
		Find(&bibiteuserlog).Error
	if err != nil {
		return &BibiteLogData{}, err
	}
	return &bibiteuserlog, nil
}

func (p *BibiteLog) FindBibiteLogAdminByBibiteUserID(db *gorm.DB, pid uint64) (*BibiteLogData, error) {
	var err error
	bibiteuserlog := BibiteLogData{}
	err = db.Debug().Model(&BibiteLogData{}).Limit(100).
		Select(`t_bibite_log.bbt_log_id,
				t_bibite_log.bbt_user_id,
				m_bibite_employee.bbt_emp_name,
				t_bibite_log.bbt_log_type,
				t_bibite_log.bbt_log_action,
				t_bibite_log.bbt_log_description,
				t_bibite_log.bbt_log_created_at at time zone 'Asia/Jakarta' as bbt_log_created_at`).
		Joins("JOIN m_bibite_administrator on t_bibite_log.bbt_user_id=m_bibite_administrator.bbt_adm_id").
		Joins("JOIN m_bibite_employee on m_bibite_administrator.bbt_emp_id=m_bibite_employee.bbt_emp_id").
		Where("t_bibite_log.bbt_user_id = ?", pid).
		Find(&bibiteuserlog).Error
	if err != nil {
		return &BibiteLogData{}, err
	}
	return &bibiteuserlog, nil
}
