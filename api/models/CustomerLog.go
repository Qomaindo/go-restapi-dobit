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

type CustomerLog struct {
	CustomerLogID          uint64    `gorm:"column:cust_log_id;primary_key;" json:"cust_log_id"`
	CustomerID             uint64    `gorm:"column:cust_id" json:"cust_id"`
	CustomerLogAction      string    `gorm:"column:cust_log_action;size:255" json:"cust_log_action"`
	CustomerLogDescription string    `gorm:"column:cust_log_description" json:"cust_log_description"`
	CustomerLogCreatedAt   time.Time `gorm:"column:cust_log_created_at;default:CURRENT_TIMESTAMP" json:"cust_log_created_at"`
}

type CustomerLogData struct {
	CustomerLogID          uint64    `gorm:"column:cust_log_id;primary_key;" json:"cust_log_id"`
	CustomerID             uint64    `gorm:"column:cust_id" json:"cust_id"`
	MitraEmployeeName      string    `gorm:"column:cust_name" json:"cust_name"`
	CustomerLogAction      string    `gorm:"column:cust_log_action" json:"cust_log_action"`
	CustomerLogDescription string    `gorm:"column:cust_log_description" json:"cust_log_description"`
	CustomerLogCreatedAt   time.Time `gorm:"column:cust_log_created_at" json:"cust_log_created_at"`
}

type ResponseCustomerLogs struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    *[]CustomerLogData `json:"data"`
}

type ResponseCustomerLog struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    *CustomerLogData `json:"data"`
}

type ResponseCustomerLogCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (CustomerLog) TableName() string {
	return "t_customer_log"
}
func (CustomerLogData) TableName() string {
	return "t_customer_log"
}

func (p *CustomerLog) Prepare() {
	p.CustomerID = p.CustomerID
	p.CustomerLogAction = html.EscapeString(strings.TrimSpace(p.CustomerLogAction))
	p.CustomerLogDescription = p.CustomerLogDescription
	p.CustomerLogCreatedAt = time.Now()
}

func (p *CustomerLog) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.CustomerID == 0 {
			return errors.New("Required Customer ID")
		}
		if p.CustomerLogAction == "" {
			return errors.New("Required Customer Log Action")
		}
		if p.CustomerLogDescription == "" {
			return errors.New("Required Customer Log Description")
		}
		return nil
	}
}

func (p *CustomerLog) SaveCustomerLog(db *gorm.DB) (*CustomerLog, error) {
	var err error
	err = db.Debug().Model(&CustomerLog{}).Create(&p).Error
	if err != nil {
		return &CustomerLog{}, err
	}
	return p, nil
}

func (p *CustomerLog) FindAllCustomerLogs(db *gorm.DB) (*[]CustomerLogData, error) {
	var err error
	mitrauserlog := []CustomerLogData{}
	err = db.Debug().Model(&CustomerLogData{}).Limit(100).
		Select(`t_customer_log.cust_log_id,
				t_customer_log.cust_id,
				t_customer_log.cust_log_action,
				t_customer_log.cust_log_description,
				t_customer_log.cust_log_created_at at time zone 'Asia/Jakarta' as cust_log_created_at`).
		Joins("JOIN m_customer on t_customer_log.cust_id=m_customer.cust_id").
		Find(&mitrauserlog).Error
	if err != nil {
		return &[]CustomerLogData{}, err
	}
	return &mitrauserlog, nil
}

func (p *CustomerLog) FindCustomerLogByID(db *gorm.DB, pid uint64) (*CustomerLogData, error) {
	var err error
	mitrauserlog := CustomerLogData{}
	err = db.Debug().Model(&CustomerLogData{}).Limit(100).
		Select(`t_customer_log.cust_log_id,
				t_customer_log.cust_id,
				t_customer_log.cust_log_action,
				t_customer_log.cust_log_description,
				t_customer_log.cust_log_created_at at time zone 'Asia/Jakarta' as cust_log_created_at`).
		Joins("JOIN m_customer on t_customer_log.cust_id=m_customer.cust_id").
		Where("cust_log_id = ?", pid).
		Find(&mitrauserlog).Error
	if err != nil {
		return &CustomerLogData{}, err
	}
	return &mitrauserlog, nil
}

func (p *CustomerLog) UpdateCustomerLog(db *gorm.DB, pid uint64) (*CustomerLog, error) {

	var err error
	db = db.Debug().Model(&CustomerLog{}).Where("cust_log_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"cust_id":              p.CustomerID,
			"cust_log_action":      p.CustomerLogAction,
			"cust_log_description": p.CustomerLogDescription,
		},
	)
	err = db.Debug().Model(&CustomerLog{}).Where("cust_log_id = ?", pid).Error
	if err != nil {
		return &CustomerLog{}, err
	}
	return p, nil
}

func (p *CustomerLog) DeleteCustomerLog(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&CustomerLog{}).Where("cust_log_id = ?", pid).Take(&CustomerLog{}).Delete(&CustomerLog{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("CustomerLog not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

func (p *CustomerLog) FindCustomerLogUserByCustomerID(db *gorm.DB, pid uint64) (*CustomerLogData, error) {
	var err error
	mitrauserlog := CustomerLogData{}
	err = db.Debug().Model(&CustomerLogData{}).Limit(100).
		Select(`t_customer_log.cust_log_id,
				t_customer_log.cust_id,
				t_customer_log.cust_log_action,
				t_customer_log.cust_log_description,
				t_customer_log.cust_log_created_at at time zone 'Asia/Jakarta' as cust_log_created_at`).
		Joins("JOIN m_customer on t_customer_log.cust_id=m_customer.cust_id").
		Joins("JOIN m_mitra_employee on m_customer.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Where("t_customer_log.cust_id = ?", pid).
		Find(&mitrauserlog).Error
	if err != nil {
		return &CustomerLogData{}, err
	}
	return &mitrauserlog, nil
}
