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

type CustomerTypeTemp struct {
	CustomerTypeTempID           uint64    `gorm:"column:cust_type_temp_id;primary_key;" json:"cust_type_temp_id"`
	CustomerTypeID               uint64    `gorm:"column:cust_type_id" json:"cust_type_id"`
	CustomerTypeTempCode         string    `gorm:"column:cust_type_temp_code" json:"cust_type_temp_code"`
	CustomerTypeTempName         string    `gorm:"column:cust_type_temp_name" json:"cust_type_temp_name"`
	CustomerTypeTempReason       string    `gorm:"column:cust_type_temp_reason" json:"cust_type_temp_reason"`
	CustomerTypeTempStatus       uint64    `gorm:"column:cust_type_temp_status;size:2" json:"cust_type_temp_status"`
	CustomerTypeTempActionBefore uint64    `gorm:"column:cust_type_temp_action_before;size:2" json:"cust_type_temp_action_before"`
	CustomerTypeTempCreatedBy    string    `gorm:"column:cust_type_temp_created_by;size:125" json:"cust_type_temp_created_by"`
	CustomerTypeTempCreatedAt    time.Time `gorm:"column:cust_type_temp_created_at;default:CURRENT_TIMESTAMP" json:"cust_type_temp_created_at"`
}

type CustomerTypeTempPend struct {
	CustomerTypeTempID           uint64    `gorm:"column:cust_type_temp_id;primary_key;" json:"cust_type_temp_id"`
	CustomerTypeID               *uint64   `gorm:"column:cust_type_id" json:"cust_type_id"`
	CustomerTypeTempCode         string    `gorm:"column:cust_type_temp_code" json:"cust_type_temp_code"`
	CustomerTypeTempName         string    `gorm:"column:cust_type_temp_name" json:"cust_type_temp_name"`
	CustomerTypeTempReason       string    `gorm:"column:cust_type_temp_reason" json:"cust_type_temp_reason"`
	CustomerTypeTempStatus       uint64    `gorm:"column:cust_type_temp_status;size:2" json:"cust_type_temp_status"`
	CustomerTypeTempActionBefore uint64    `gorm:"column:cust_type_temp_action_before;size:2" json:"cust_type_temp_action_before"`
	CustomerTypeTempCreatedBy    string    `gorm:"column:cust_type_temp_created_by;size:125" json:"cust_type_temp_created_by"`
	CustomerTypeTempCreatedAt    time.Time `gorm:"column:cust_type_temp_created_at;default:CURRENT_TIMESTAMP" json:"cust_type_temp_created_at"`
}

type CustomerTypeTempData struct {
	CustomerTypeTempID           uint64    `gorm:"column:cust_type_temp_id" json:"cust_type_temp_id"`
	CustomerTypeTempCode         string    `gorm:"column:cust_type_temp_code" json:"cust_type_temp_code"`
	CustomerTypeTempName         string    `gorm:"column:cust_type_temp_name" json:"cust_type_temp_name"`
	CustomerTypeTempReason       string    `gorm:"column:cust_type_temp_reason" json:"cust_type_temp_reason"`
	CustomerTypeTempStatus       uint64    `gorm:"column:cust_type_temp_status;size:2" json:"cust_type_temp_status"`
	CustomerTypeTempActionBefore uint64    `gorm:"column:cust_type_temp_action_before;size:2" json:"cust_type_temp_action_before"`
	CustomerTypeTempCreatedBy    string    `gorm:"column:cust_type_temp_created_by;size:125" json:"cust_type_temp_created_by"`
	CustomerTypeTempCreatedAt    time.Time `gorm:"column:cust_type_temp_created_at" json:"cust_type_temp_created_at"`
	CustomerTypeID               uint64    `gorm:"column:cust_type_id" json:"cust_type_id"`
	CustomerTypeCode             string    `gorm:"column:cust_type_code" json:"cust_type_code"`
	CustomerTypeName             string    `gorm:"column:cust_type_name" json:"cust_type_name"`
	CustomerTypeStatus           uint64    `gorm:"column:cust_type_status" json:"cust_type_status"`
	CustomerTypeCreatedBy        string    `gorm:"column:cust_type_created_by" json:"cust_type_created_by"`
	CustomerTypeCreatedAt        time.Time `gorm:"column:cust_type_created_at" json:"cust_type_created_at"`
	CustomerTypeUpdatedBy        string    `gorm:"column:cust_type_updated_by" json:"cust_type_updated_by"`
	CustomerTypeUpdatedAt        time.Time `gorm:"column:cust_type_updated_at" json:"cust_type_updated_at"`
	CustomerTypeDeletedBy        string    `gorm:"column:cust_type_deleted_by" json:"cust_type_deleted_by"`
	CustomerTypeDeletedAt        time.Time `gorm:"column:cust_type_deleted_at" json:"cust_type_deleted_at"`
}

type ResponseCustomerTypeTemps struct {
	Status  int                     `json:"status"`
	Message string                  `json:"message"`
	Data    *[]CustomerTypeTempData `json:"data"`
}

type ResponseCustomerTypeTemp struct {
	Status  int                   `json:"status"`
	Message string                `json:"message"`
	Data    *CustomerTypeTempData `json:"data"`
}

type ResponseCustomerTypeTempsPend struct {
	Status  int                     `json:"status"`
	Message string                  `json:"message"`
	Data    *[]CustomerTypeTempPend `json:"data"`
}

type ResponseCustomerTypeTempPend struct {
	Status  int                   `json:"status"`
	Message string                `json:"message"`
	Data    *CustomerTypeTempPend `json:"data"`
}

type ResponseCustomerTypeTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (CustomerTypeTemp) TableName() string {
	return "m_customer_type_temp"
}

func (CustomerTypeTempPend) TableName() string {
	return "m_customer_type_temp"
}

func (CustomerTypeTempData) TableName() string {
	return "m_customer_type_temp"
}

func (p *CustomerTypeTemp) Prepare() {
	p.CustomerTypeID = p.CustomerTypeID
	p.CustomerTypeTempCode = html.EscapeString(strings.TrimSpace(p.CustomerTypeTempCode))
	p.CustomerTypeTempName = html.EscapeString(strings.TrimSpace(p.CustomerTypeTempName))
	p.CustomerTypeTempReason = p.CustomerTypeTempReason
	p.CustomerTypeTempStatus = p.CustomerTypeTempStatus
	p.CustomerTypeTempCreatedBy = html.EscapeString(strings.TrimSpace(p.CustomerTypeTempCreatedBy))
	p.CustomerTypeTempCreatedAt = time.Now()
}

func (p *CustomerTypeTempPend) Prepare() {
	p.CustomerTypeID = nil
	p.CustomerTypeTempCode = html.EscapeString(strings.TrimSpace(p.CustomerTypeTempCode))
	p.CustomerTypeTempName = html.EscapeString(strings.TrimSpace(p.CustomerTypeTempName))
	p.CustomerTypeTempReason = p.CustomerTypeTempReason
	p.CustomerTypeTempStatus = p.CustomerTypeTempStatus
	p.CustomerTypeTempCreatedBy = html.EscapeString(strings.TrimSpace(p.CustomerTypeTempCreatedBy))
	p.CustomerTypeTempCreatedAt = time.Now()
}

func (p *CustomerTypeTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	case "insertupdate":
		if p.CustomerTypeTempCode == "" {
			return errors.New("Negara Wajib diisi")
		}
		if p.CustomerTypeTempName == "" {
			return errors.New("Provinsi Wajib diisi")
		}
		return nil

	case "reason":
		if p.CustomerTypeTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.CustomerTypeTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	default:
		if p.CustomerTypeTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil
	}
}

func (p *CustomerTypeTemp) SaveCustomerTypeTemp(db *gorm.DB) (*CustomerTypeTemp, error) {
	var err error
	err = db.Debug().Model(&CustomerTypeTemp{}).Create(&p).Error
	if err != nil {
		return &CustomerTypeTemp{}, err
	}
	return p, nil
}

func (p *CustomerTypeTempPend) SaveCustomerTypeTempPend(db *gorm.DB) (*CustomerTypeTempPend, error) {
	var err error
	err = db.Debug().Model(&CustomerTypeTempPend{}).Create(&p).Error
	if err != nil {
		return &CustomerTypeTempPend{}, err
	}
	return p, nil
}

func (p *CustomerTypeTemp) FindAllCustomerTypeTemp(db *gorm.DB) (*[]CustomerTypeTempPend, error) {
	var err error
	customertype := []CustomerTypeTempPend{}
	err = db.Debug().Model(&CustomerTypeTempPend{}).Limit(100).
		Select(`m_customer_type_temp.cust_type_temp_id,
				m_customer_type_temp.cust_type_id,
				m_customer_type_temp.cust_type_temp_code,
				m_customer_type_temp.cust_type_temp_name,
				m_customer_type_temp.cust_type_temp_reason,
				m_customer_type_temp.cust_type_temp_status,
				m_customer_type_temp.cust_type_temp_action_before,
				m_customer_type_temp.cust_type_temp_created_by,
				m_customer_type_temp.cust_type_temp_created_at at time zone 'Asia/Jakarta' as address_deleted_at`).
		Find(&customertype).Error
	if err != nil {
		return &[]CustomerTypeTempPend{}, err
	}
	return &customertype, nil
}

func (p *CustomerTypeTemp) FindAllCustomerTypeTempPendingActive(db *gorm.DB) (*[]CustomerTypeTempPend, error) {
	var err error
	customertype := []CustomerTypeTempPend{}
	err = db.Debug().Model(&CustomerTypeTempPend{}).Limit(100).
		Select(`m_customer_type_temp.cust_type_temp_id,
				m_customer_type_temp.cust_type_id,
				m_customer_type_temp.cust_type_temp_code,
				m_customer_type_temp.cust_type_temp_name,
				m_customer_type_temp.cust_type_temp_reason,
				m_customer_type_temp.cust_type_temp_status,
				m_customer_type_temp.cust_type_temp_action_before,
				m_customer_type_temp.cust_type_temp_created_by,
				m_customer_type_temp.cust_type_temp_created_at at time zone 'Asia/Jakarta' as address_deleted_at`).
		Where("cust_type_temp_status = ?", 11).
		Find(&customertype).Error
	if err != nil {
		return &[]CustomerTypeTempPend{}, err
	}
	return &customertype, nil
}

func (p *CustomerTypeTemp) FindAllCustomerTypeTempStatus(db *gorm.DB, status uint64) (*[]CustomerTypeTempData, error) {
	var err error
	customertype := []CustomerTypeTempData{}
	err = db.Debug().Model(&CustomerTypeTempData{}).Limit(100).
		Select(`m_customer_type_temp.cust_type_temp_id,
				m_customer_type_temp.cust_type_temp_code,
				m_customer_type_temp.cust_type_temp_name,
				m_customer_type_temp.cust_type_temp_reason,
				m_customer_type_temp.cust_type_temp_status,
				m_customer_type_temp.cust_type_temp_action_before,
				m_customer_type_temp.cust_type_temp_created_by,
				m_customer_type_temp.cust_type_temp_created_at at time zone 'Asia/Jakarta' as address_deleted_at,
				m_customer_type.cust_type_id,
				m_customer_type.cust_type_code,
				m_customer_type.cust_type_name,
				m_customer_type.cust_type_status,
				m_customer_type.cust_type_created_by,
				m_customer_type.cust_type_created_at at time zone 'Asia/Jakarta' as address_deleted_at,
				m_customer_type.cust_type_updated_by,
				m_customer_type.cust_type_updated_at at time zone 'Asia/Jakarta' as address_deleted_at,
				m_customer_type.cust_type_deleted_by,
				m_customer_type.cust_type_deleted_at at time zone 'Asia/Jakarta' as address_deleted_at`).
		Joins("JOIN m_customer_type on m_customer_type_temp.cust_type_id=m_customer_type.cust_type_id").
		Where("cust_type_temp_status = ?", status).
		Find(&customertype).Error
	if err != nil {
		return &[]CustomerTypeTempData{}, err
	}
	return &customertype, nil
}

func (p *CustomerTypeTemp) FindCustomerTypeTempDataByID(db *gorm.DB, pid uint64) (*CustomerTypeTemp, error) {
	var err error
	err = db.Debug().Model(&CustomerTypeTemp{}).Where("cust_type_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &CustomerTypeTemp{}, err
	}
	return p, nil
}

func (p *CustomerTypeTemp) FindCustomerTypeTempByIDPendingActive(db *gorm.DB, pid uint64) (*CustomerTypeTempPend, error) {
	var err error
	customertype := CustomerTypeTempPend{}
	err = db.Debug().Model(&CustomerTypeTempPend{}).Limit(100).
		Select(`m_customer_type_temp.cust_type_temp_id,
				m_customer_type_temp.cust_type_id,
				m_customer_type_temp.cust_type_temp_code,
				m_customer_type_temp.cust_type_temp_name,
				m_customer_type_temp.cust_type_temp_reason,
				m_customer_type_temp.cust_type_temp_status,
				m_customer_type_temp.cust_type_temp_created_by,
				m_customer_type_temp.cust_type_temp_created_at at time zone 'Asia/Jakarta' as address_deleted_at`).
		Where("cust_type_temp_id = ?", pid).
		Find(&customertype).Error
	if err != nil {
		return &CustomerTypeTempPend{}, err
	}
	return &customertype, nil
}

func (p *CustomerTypeTemp) FindCustomerTypeTempByID(db *gorm.DB, pid uint64) (*CustomerTypeTempData, error) {
	var err error
	customertype := CustomerTypeTempData{}
	err = db.Debug().Model(&CustomerTypeTempData{}).Limit(100).
		Select(`m_customer_type_temp.cust_type_temp_id,
				m_customer_type_temp.cust_type_temp_code,
				m_customer_type_temp.cust_type_temp_name,
				m_customer_type_temp.cust_type_temp_reason,
				m_customer_type_temp.cust_type_temp_status,
				m_customer_type_temp.cust_type_temp_action_before,
				m_customer_type_temp.cust_type_temp_created_by,
				m_customer_type_temp.cust_type_temp_created_at at time zone 'Asia/Jakarta' as address_deleted_at,
				m_customer_type.cust_type_id,
				m_customer_type.cust_type_code,
				m_customer_type.cust_type_name,
				m_customer_type.cust_type_status,
				m_customer_type.cust_type_created_by,
				m_customer_type.cust_type_created_at at time zone 'Asia/Jakarta' as address_deleted_at,
				m_customer_type.cust_type_updated_by,
				m_customer_type.cust_type_updated_at at time zone 'Asia/Jakarta' as address_deleted_at,
				m_customer_type.cust_type_deleted_by,
				m_customer_type.cust_type_deleted_at at time zone 'Asia/Jakarta' as address_deleted_at`).
		Joins("JOIN m_customer_type on m_customer_type_temp.cust_type_id=m_customer_type.cust_type_id").
		Where("cust_type_temp_id = ?", pid).
		Find(&customertype).Error
	if err != nil {
		return &CustomerTypeTempData{}, err
	}
	return &customertype, nil
}

func (p *CustomerTypeTemp) FindCustomerTypeTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*CustomerTypeTempData, error) {
	var err error
	customertype := CustomerTypeTempData{}
	err = db.Debug().Model(&CustomerTypeTempData{}).Limit(100).
		Select(`m_customer_type_temp.cust_type_temp_id,
				m_customer_type_temp.cust_type_temp_code,
				m_customer_type_temp.cust_type_temp_name,
				m_customer_type_temp.cust_type_temp_reason,
				m_customer_type_temp.cust_type_temp_status,
				m_customer_type_temp.cust_type_temp_action_before,
				m_customer_type_temp.cust_type_temp_created_by,
				m_customer_type_temp.cust_type_temp_created_at at time zone 'Asia/Jakarta' as address_deleted_at,
				m_customer_type.cust_type_id,
				m_customer_type.cust_type_code,
				m_customer_type.cust_type_name,
				m_customer_type.cust_type_status,
				m_customer_type.cust_type_created_by,
				m_customer_type.cust_type_created_at at time zone 'Asia/Jakarta' as address_deleted_at,
				m_customer_type.cust_type_updated_by,
				m_customer_type.cust_type_updated_at at time zone 'Asia/Jakarta' as address_deleted_at,
				m_customer_type.cust_type_deleted_by,
				m_customer_type.cust_type_deleted_at at time zone 'Asia/Jakarta' as address_deleted_at`).
		Joins("JOIN m_customer_type on m_customer_type_temp.cust_type_id=m_customer_type.cust_type_id").
		Where("cust_type_temp_id = ? AND cust_type_temp_status = ?", pid, status).
		Find(&customertype).Error
	if err != nil {
		return &CustomerTypeTempData{}, err
	}
	return &customertype, nil
}

func (p *CustomerTypeTemp) UpdateCustomerTypeTemp(db *gorm.DB, pid uint64) (*CustomerTypeTemp, error) {
	var err error
	db = db.Debug().Model(&CustomerTypeTemp{}).Where("cust_type_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"cust_type_temp_code":   p.CustomerTypeTempCode,
			"cust_type_temp_name":   p.CustomerTypeTempName,
			"cust_type_temp_reason": p.CustomerTypeTempReason,
		},
	)
	err = db.Debug().Model(&CustomerTypeTemp{}).Where("cust_type_temp_id = ?", pid).Error
	if err != nil {
		return &CustomerTypeTemp{}, err
	}
	return p, nil
}

func (p *CustomerTypeTemp) UpdateCustomerTypeTempStatus(db *gorm.DB, pid uint64) (*CustomerTypeTemp, error) {
	var err error
	db = db.Debug().Model(&CustomerTypeTemp{}).Where("cust_type_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"cust_type_temp_reason": p.CustomerTypeTempReason,
			"cust_type_temp_status": p.CustomerTypeTempStatus,
		},
	)
	err = db.Debug().Model(&CustomerTypeTemp{}).Where("cust_type_temp_id = ?", pid).Error
	if err != nil {
		return &CustomerTypeTemp{}, err
	}
	return p, nil
}

func (p *CustomerTypeTemp) DeleteCustomerTypeTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&CustomerTypeTemp{}).Where("cust_type_temp_id = ?", pid).Take(&CustomerTypeTemp{}).Delete(&CustomerTypeTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("CustomerTypeTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================
