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

type CustomerSubTypeTemp struct {
	CustomerSubTypeTempID           uint64    `gorm:"column:cust_sub_type_temp_id;primary_key;" json:"cust_sub_type_temp_id"`
	CustomerSubTypeID               uint64    `gorm:"column:cust_sub_type_id" json:"cust_sub_type_id"`
	CustomerTypeTempID              uint64    `gorm:"column:cust_type_temp_id" json:"cust_type_temp_id"`
	CustomerSubTypeTempCode         string    `gorm:"column:cust_sub_type_temp_code" json:"cust_sub_type_temp_code"`
	CustomerSubTypeTempName         string    `gorm:"column:cust_sub_type_temp_name" json:"cust_sub_type_temp_name"`
	CustomerSubTypeTempReason       string    `gorm:"column:cust_sub_type_temp_reason" json:"cust_sub_type_temp_reason"`
	CustomerSubTypeTempStatus       uint64    `gorm:"column:cust_sub_type_temp_status;size:2" json:"cust_sub_type_temp_status"`
	CustomerSubTypeTempActionBefore uint64    `gorm:"column:cust_sub_type_temp_action_before;size:2" json:"cust_sub_type_temp_action_before"`
	CustomerSubTypeTempCreatedBy    string    `gorm:"column:cust_sub_type_temp_created_by;size:125" json:"cust_sub_type_temp_created_by"`
	CustomerSubTypeTempCreatedAt    time.Time `gorm:"column:cust_sub_type_temp_created_at;default:CURRENT_TIMESTAMP" json:"cust_sub_type_temp_created_at"`
}

type CustomerSubTypeTempPend struct {
	CustomerSubTypeTempID           uint64    `gorm:"column:cust_sub_type_temp_id;primary_key;" json:"cust_sub_type_temp_id"`
	CustomerSubTypeID               *uint64   `gorm:"column:cust_sub_type_id" json:"cust_sub_type_id"`
	CustomerTypeTempID              uint64    `gorm:"column:cust_type_temp_id" json:"cust_type_temp_id"`
	CustomerSubTypeTempCode         string    `gorm:"column:cust_sub_type_temp_code" json:"cust_sub_type_temp_code"`
	CustomerSubTypeTempName         string    `gorm:"column:cust_sub_type_temp_name" json:"cust_sub_type_temp_name"`
	CustomerSubTypeTempReason       string    `gorm:"column:cust_sub_type_temp_reason" json:"cust_sub_type_temp_reason"`
	CustomerSubTypeTempStatus       uint64    `gorm:"column:cust_sub_type_temp_status;size:2" json:"cust_sub_type_temp_status"`
	CustomerSubTypeTempActionBefore uint64    `gorm:"column:cust_sub_type_temp_action_before;size:2" json:"cust_sub_type_temp_action_before"`
	CustomerSubTypeTempCreatedBy    string    `gorm:"column:cust_sub_type_temp_created_by;size:125" json:"cust_sub_type_temp_created_by"`
	CustomerSubTypeTempCreatedAt    time.Time `gorm:"column:cust_sub_type_temp_created_at;default:CURRENT_TIMESTAMP" json:"cust_sub_type_temp_created_at"`
}

type CustomerSubTypeTempData struct {
	CustomerSubTypeTempID           uint64    `gorm:"column:cust_sub_type_temp_id" json:"cust_sub_type_temp_id"`
	CustomerTypeTempID              uint64    `gorm:"column:cust_type_temp_id" json:"cust_type_temp_id"`
	CustomerTypeTempName            string    `gorm:"column:cust_type_temp_name" json:"cust_type_temp_name"`
	CustomerSubTypeTempCode         string    `gorm:"column:cust_sub_type_temp_code" json:"cust_sub_type_temp_code"`
	CustomerSubTypeTempName         string    `gorm:"column:cust_sub_type_temp_name" json:"cust_sub_type_temp_name"`
	CustomerSubTypeTempReason       string    `gorm:"column:cust_sub_type_temp_reason" json:"cust_sub_type_temp_reason"`
	CustomerSubTypeTempStatus       uint64    `gorm:"column:cust_sub_type_temp_status;size:2" json:"cust_sub_type_temp_status"`
	CustomerSubTypeTempActionBefore uint64    `gorm:"column:cust_sub_type_temp_action_before;size:2" json:"cust_sub_type_temp_action_before"`
	CustomerSubTypeTempCreatedBy    string    `gorm:"column:cust_sub_type_temp_created_by;size:125" json:"cust_sub_type_temp_created_by"`
	CustomerSubTypeTempCreatedAt    time.Time `gorm:"column:cust_sub_type_temp_created_at" json:"cust_sub_type_temp_created_at"`
	CustomerSubTypeID               uint64    `gorm:"column:cust_sub_type_id" json:"cust_sub_type_id"`
	CustomerTypeID                  uint64    `gorm:"column:cust_type_id" json:"cust_type_id"`
	CustomerTypeName                string    `gorm:"column:cust_type_name" json:"cust_type_name"`
	CustomerSubTypeCode             string    `gorm:"column:cust_sub_type_code" json:"cust_sub_type_code"`
	CustomerSubTypeName             string    `gorm:"column:cust_sub_type_name" json:"cust_sub_type_name"`
	CustomerSubTypeStatus           uint64    `gorm:"column:cust_sub_type_status" json:"cust_sub_type_status"`
	CustomerSubTypeCreatedBy        string    `gorm:"column:cust_sub_type_created_by" json:"cust_sub_type_created_by"`
	CustomerSubTypeCreatedAt        time.Time `gorm:"column:cust_sub_type_created_at" json:"cust_sub_type_created_at"`
	CustomerSubTypeUpdatedBy        string    `gorm:"column:cust_sub_type_updated_by" json:"cust_sub_type_updated_by"`
	CustomerSubTypeUpdatedAt        time.Time `gorm:"column:cust_sub_type_updated_at" json:"cust_sub_type_updated_at"`
	CustomerSubTypeDeletedBy        string    `gorm:"column:cust_sub_type_deleted_by" json:"cust_sub_type_deleted_by"`
	CustomerSubTypeDeletedAt        time.Time `gorm:"column:cust_sub_type_deleted_at" json:"cust_sub_type_deleted_at"`
}

type ResponseCustomerSubTypeTemps struct {
	Status  int                        `json:"status"`
	Message string                     `json:"message"`
	Data    *[]CustomerSubTypeTempData `json:"data"`
}

type ResponseCustomerSubTypeTemp struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    *CustomerSubTypeTempData `json:"data"`
}

type ResponseCustomerSubTypeTempsPend struct {
	Status  int                        `json:"status"`
	Message string                     `json:"message"`
	Data    *[]CustomerSubTypeTempPend `json:"data"`
}

type ResponseCustomerSubTypeTempPend struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    *CustomerSubTypeTempPend `json:"data"`
}

type ResponseCustomerSubTypeTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (CustomerSubTypeTemp) TableName() string {
	return "m_customer_sub_type_temp"
}

func (CustomerSubTypeTempPend) TableName() string {
	return "m_customer_sub_type_temp"
}

func (CustomerSubTypeTempData) TableName() string {
	return "m_customer_sub_type_temp"
}

func (p *CustomerSubTypeTemp) Prepare() {
	p.CustomerSubTypeID = p.CustomerSubTypeID
	p.CustomerTypeTempID = p.CustomerTypeTempID
	p.CustomerSubTypeTempCode = p.CustomerSubTypeTempCode
	p.CustomerSubTypeTempName = p.CustomerSubTypeTempName
	p.CustomerSubTypeTempReason = p.CustomerSubTypeTempReason
	p.CustomerSubTypeTempStatus = p.CustomerSubTypeTempStatus
	p.CustomerSubTypeTempCreatedBy = html.EscapeString(strings.TrimSpace(p.CustomerSubTypeTempCreatedBy))
	p.CustomerSubTypeTempCreatedAt = time.Now()
}

func (p *CustomerSubTypeTempPend) Prepare() {
	p.CustomerSubTypeID = nil
	p.CustomerTypeTempID = p.CustomerTypeTempID
	p.CustomerSubTypeTempCode = p.CustomerSubTypeTempCode
	p.CustomerSubTypeTempName = p.CustomerSubTypeTempName
	p.CustomerSubTypeTempReason = p.CustomerSubTypeTempReason
	p.CustomerSubTypeTempStatus = p.CustomerSubTypeTempStatus
	p.CustomerSubTypeTempCreatedBy = html.EscapeString(strings.TrimSpace(p.CustomerSubTypeTempCreatedBy))
	p.CustomerSubTypeTempCreatedAt = time.Now()
}

func (p *CustomerSubTypeTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	case "insertupdate":
		if p.CustomerTypeTempID == 0 {
			return errors.New("Negara Wajib diisi")
		}
		if p.CustomerSubTypeTempCode == "" {
			return errors.New("Provinsi Wajib diisi")
		}
		if p.CustomerSubTypeTempName == "" {
			return errors.New("Kabupaten / Kota Wajib diisi")
		}
		return nil

	case "reason":
		if p.CustomerSubTypeTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.CustomerSubTypeTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	default:
		if p.CustomerSubTypeTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil
	}
}

func (p *CustomerSubTypeTemp) SaveCustomerSubTypeTemp(db *gorm.DB) (*CustomerSubTypeTemp, error) {
	var err error
	err = db.Debug().Model(&CustomerSubTypeTemp{}).Create(&p).Error
	if err != nil {
		return &CustomerSubTypeTemp{}, err
	}
	return p, nil
}

func (p *CustomerSubTypeTempPend) SaveCustomerSubTypeTempPend(db *gorm.DB) (*CustomerSubTypeTempPend, error) {
	var err error
	err = db.Debug().Model(&CustomerSubTypeTempPend{}).Create(&p).Error
	if err != nil {
		return &CustomerSubTypeTempPend{}, err
	}
	return p, nil
}

func (p *CustomerSubTypeTemp) FindAllCustomerSubTypeTemp(db *gorm.DB) (*[]CustomerSubTypeTempPend, error) {
	var err error
	customersubtype := []CustomerSubTypeTempPend{}
	err = db.Debug().Model(&CustomerSubTypeTempPend{}).Limit(100).
		Select(`m_customer_sub_type_temp.cust_sub_type_temp_id,
				m_customer_sub_type_temp.cust_sub_type_id,
				m_customer_sub_type_temp.cust_type_temp_id,
				m_customer_sub_type_temp.cust_sub_type_temp_code,
				m_customer_sub_type_temp.cust_sub_type_temp_name,
				m_customer_sub_type_temp.cust_sub_type_temp_reason,
				m_customer_sub_type_temp.cust_sub_type_temp_status,
				m_customer_sub_type_temp.cust_sub_type_temp_action_before,
				m_customer_sub_type_temp.cust_sub_type_temp_created_by,
				m_customer_sub_type_temp.cust_sub_type_temp_created_at at time zone 'Asia/Jakarta' as cust_sub_type_temp_created_at`).
		Order("cust_sub_type_temp_created_at desc").
		Find(&customersubtype).Error
	if err != nil {
		return &[]CustomerSubTypeTempPend{}, err
	}
	return &customersubtype, nil
}

func (p *CustomerSubTypeTemp) FindAllCustomerSubTypeTempStatusPendingActive(db *gorm.DB) (*[]CustomerSubTypeTempPend, error) {
	var err error
	customersubtype := []CustomerSubTypeTempPend{}
	err = db.Debug().Model(&CustomerSubTypeTempPend{}).Limit(100).
		Select(`m_customer_sub_type_temp.cust_sub_type_temp_id,
				m_customer_sub_type_temp.cust_sub_type_id,
				m_customer_sub_type_temp.cust_type_temp_id,
				m_customer_sub_type_temp.cust_sub_type_temp_code,
				m_customer_sub_type_temp.cust_sub_type_temp_name,
				m_customer_sub_type_temp.cust_sub_type_temp_reason,
				m_customer_sub_type_temp.cust_sub_type_temp_status,
				m_customer_sub_type_temp.cust_sub_type_temp_action_before,
				m_customer_sub_type_temp.cust_sub_type_temp_created_by,
				m_customer_sub_type_temp.cust_sub_type_temp_created_at at time zone 'Asia/Jakarta' as cust_sub_type_temp_created_at`).
		Where("cust_sub_type_temp_status = ?", 11).
		Order("cust_sub_type_temp_created_at desc").
		Find(&customersubtype).Error
	if err != nil {
		return &[]CustomerSubTypeTempPend{}, err
	}
	return &customersubtype, nil
}

func (p *CustomerSubTypeTemp) FindAllCustomerSubTypeTempStatus(db *gorm.DB, status uint64) (*[]CustomerSubTypeTempData, error) {
	var err error
	customersubtype := []CustomerSubTypeTempData{}
	err = db.Debug().Model(&CustomerSubTypeTempData{}).Limit(100).
		Select(`m_customer_sub_type_temp.cust_sub_type_temp_id,
				m_customer_type_temp.cust_type_id as cust_type_temp_id,
				m_customer_type_temp.cust_type_name as cust_type_temp_name,
				m_customer_sub_type_temp.cust_sub_type_temp_code,
				m_customer_sub_type_temp.cust_sub_type_temp_name,
				m_customer_sub_type_temp.cust_sub_type_temp_reason,
				m_customer_sub_type_temp.cust_sub_type_temp_status,
				m_customer_sub_type_temp.cust_sub_type_temp_action_before,
				m_customer_sub_type_temp.cust_sub_type_temp_created_by,
				m_customer_sub_type_temp.cust_sub_type_temp_created_at at time zone 'Asia/Jakarta' as cust_sub_type_temp_created_at,
				m_customer_sub_type.cust_sub_type_id,
				m_customer_type.cust_type_id,
				m_customer_type.cust_type_name,
				m_customer_sub_type.cust_sub_type_code,
				m_customer_sub_type.cust_sub_type_name,
				m_customer_sub_type.cust_sub_type_status,
				m_customer_sub_type.cust_sub_type_created_by,
				m_customer_sub_type.cust_sub_type_created_at at time zone 'Asia/Jakarta' as cust_sub_type_created_at,
				m_customer_sub_type.cust_sub_type_updated_by,
				m_customer_sub_type.cust_sub_type_updated_at at time zone 'Asia/Jakarta' as cust_sub_type_updated_at,
				m_customer_sub_type.cust_sub_type_deleted_by,
				m_customer_sub_type.cust_sub_type_deleted_at at time zone 'Asia/Jakarta' as cust_sub_type_deleted_at`).
		Joins("JOIN m_customer_type m_customer_type_temp on m_customer_sub_type_temp.cust_type_temp_id=m_customer_type_temp.cust_type_id").
		Joins("JOIN m_customer_sub_type on m_customer_sub_type_temp.cust_sub_type_id=m_customer_sub_type.cust_sub_type_id").
		Joins("JOIN m_customer_type on m_customer_sub_type.cust_type_id=m_customer_type.cust_type_id").
		Where("cust_sub_type_temp_status = ?", status).
		Order("cust_sub_type_temp_created_at desc").
		Find(&customersubtype).Error
	if err != nil {
		return &[]CustomerSubTypeTempData{}, err
	}
	return &customersubtype, nil
}

func (p *CustomerSubTypeTemp) FindCustomerSubTypeTempDataByID(db *gorm.DB, pid uint64) (*CustomerSubTypeTemp, error) {
	var err error
	err = db.Debug().Model(&CustomerSubTypeTemp{}).Where("cust_sub_type_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &CustomerSubTypeTemp{}, err
	}
	return p, nil
}

func (p *CustomerSubTypeTemp) FindCustomerSubTypeTempByIDPendingActive(db *gorm.DB, pid uint64) (*CustomerSubTypeTempPend, error) {
	var err error
	customersubtype := CustomerSubTypeTempPend{}
	err = db.Debug().Model(&CustomerSubTypeTempPend{}).Limit(100).
		Select(`m_customer_sub_type_temp.cust_sub_type_temp_id,
				m_customer_sub_type_temp.cust_sub_type_id,
				m_customer_sub_type_temp.cust_type_temp_id,
				m_customer_sub_type_temp.cust_sub_type_temp_code,
				m_customer_sub_type_temp.cust_sub_type_temp_name,
				m_customer_sub_type_temp.cust_sub_type_temp_reason,
				m_customer_sub_type_temp.cust_sub_type_temp_status,
				m_customer_sub_type_temp.cust_sub_type_temp_created_by,
				m_customer_sub_type_temp.cust_sub_type_temp_created_at at time zone 'Asia/Jakarta' as cust_sub_type_temp_created_at`).
		Where("cust_sub_type_temp_id = ?", pid).
		Order("cust_sub_type_temp_created_at desc").
		Find(&customersubtype).Error
	if err != nil {
		return &CustomerSubTypeTempPend{}, err
	}
	return &customersubtype, nil
}

func (p *CustomerSubTypeTemp) FindCustomerSubTypeTempStatusByIDPendingActive(db *gorm.DB, pid uint64) (*CustomerSubTypeTempPend, error) {
	var err error
	customersubtype := CustomerSubTypeTempPend{}
	err = db.Debug().Model(&CustomerSubTypeTempPend{}).Limit(100).
		Select(`m_customer_sub_type_temp.cust_sub_type_temp_id,
				m_customer_sub_type_temp.cust_sub_type_id,
				m_customer_sub_type_temp.cust_type_temp_id,
				m_customer_sub_type_temp.cust_sub_type_temp_code,
				m_customer_sub_type_temp.cust_sub_type_temp_name,
				m_customer_sub_type_temp.cust_sub_type_temp_reason,
				m_customer_sub_type_temp.cust_sub_type_temp_status,
				m_customer_sub_type_temp.cust_sub_type_temp_action_before,
				m_customer_sub_type_temp.cust_sub_type_temp_created_by,
				m_customer_sub_type_temp.cust_sub_type_temp_created_at at time zone 'Asia/Jakarta' as cust_sub_type_temp_created_at`).
		Where("cust_sub_type_temp_id = ? AND cust_sub_type_temp_status = ?", pid, 11).
		Order("cust_sub_type_temp_created_at desc").
		Find(&customersubtype).Error
	if err != nil {
		return &CustomerSubTypeTempPend{}, err
	}
	return &customersubtype, nil
}

func (p *CustomerSubTypeTemp) FindCustomerSubTypeTempByID(db *gorm.DB, pid uint64) (*CustomerSubTypeTempData, error) {
	var err error
	customersubtype := CustomerSubTypeTempData{}
	err = db.Debug().Model(&CustomerSubTypeTempData{}).Limit(100).
		Select(`m_customer_sub_type_temp.cust_sub_type_temp_id,
				m_customer_type_temp.cust_type_id as cust_type_temp_id,
				m_customer_type_temp.cust_type_name as cust_type_temp_name,
				m_customer_sub_type_temp.cust_sub_type_temp_code,
				m_customer_sub_type_temp.cust_sub_type_temp_name,
				m_customer_sub_type_temp.cust_sub_type_temp_reason,
				m_customer_sub_type_temp.cust_sub_type_temp_status,
				m_customer_sub_type_temp.cust_sub_type_temp_action_before,
				m_customer_sub_type_temp.cust_sub_type_temp_created_by,
				m_customer_sub_type_temp.cust_sub_type_temp_created_at at time zone 'Asia/Jakarta' as cust_sub_type_temp_created_at,
				m_customer_sub_type.cust_sub_type_id,
				m_customer_type.cust_type_id,
				m_customer_type.cust_type_name,
				m_customer_sub_type.cust_sub_type_code,
				m_customer_sub_type.cust_sub_type_name,
				m_customer_sub_type.cust_sub_type_status,
				m_customer_sub_type.cust_sub_type_created_by,
				m_customer_sub_type.cust_sub_type_created_at at time zone 'Asia/Jakarta' as cust_sub_type_created_at,
				m_customer_sub_type.cust_sub_type_updated_by,
				m_customer_sub_type.cust_sub_type_updated_at at time zone 'Asia/Jakarta' as cust_sub_type_updated_at,
				m_customer_sub_type.cust_sub_type_deleted_by,
				m_customer_sub_type.cust_sub_type_deleted_at at time zone 'Asia/Jakarta' as cust_sub_type_deleted_at`).
		Joins("JOIN m_customer_type m_customer_type_temp on m_customer_sub_type_temp.cust_type_temp_id=m_customer_type_temp.cust_type_id").
		Joins("JOIN m_customer_sub_type on m_customer_sub_type_temp.cust_sub_type_id=m_customer_sub_type.cust_sub_type_id").
		Joins("JOIN m_customer_type on m_customer_sub_type.cust_type_id=m_customer_type.cust_type_id").
		Order("cust_sub_type_temp_created_at desc").
		Where("cust_sub_type_temp_id = ?", pid).
		Find(&customersubtype).Error
	if err != nil {
		return &CustomerSubTypeTempData{}, err
	}
	return &customersubtype, nil
}

func (p *CustomerSubTypeTemp) FindCustomerSubTypeTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*CustomerSubTypeTempData, error) {
	var err error
	customersubtype := CustomerSubTypeTempData{}
	err = db.Debug().Model(&CustomerSubTypeTempData{}).Limit(100).
		Select(`m_customer_sub_type_temp.cust_sub_type_temp_id,
				m_customer_type_temp.cust_type_id as cust_type_temp_id,
				m_customer_type_temp.cust_type_name as cust_type_temp_name,
				m_customer_sub_type_temp.cust_sub_type_temp_code,
				m_customer_sub_type_temp.cust_sub_type_temp_name,
				m_customer_sub_type_temp.cust_sub_type_temp_reason,
				m_customer_sub_type_temp.cust_sub_type_temp_status,
				m_customer_sub_type_temp.cust_sub_type_temp_action_before,
				m_customer_sub_type_temp.cust_sub_type_temp_created_by,
				m_customer_sub_type_temp.cust_sub_type_temp_created_at,
				m_customer_sub_type.cust_sub_type_id,
				m_customer_type.cust_type_id,
				m_customer_type.cust_type_name,
				m_customer_sub_type.cust_sub_type_code,
				m_customer_sub_type.cust_sub_type_name,
				m_customer_sub_type.cust_sub_type_status,
				m_customer_sub_type.cust_sub_type_created_by,
				m_customer_sub_type.cust_sub_type_created_at at time zone 'Asia/Jakarta' as cust_sub_type_created_at,
				m_customer_sub_type.cust_sub_type_updated_by,
				m_customer_sub_type.cust_sub_type_updated_at at time zone 'Asia/Jakarta' as cust_sub_type_updated_at,
				m_customer_sub_type.cust_sub_type_deleted_by,
				m_customer_sub_type.cust_sub_type_deleted_at at time zone 'Asia/Jakarta' as cust_sub_type_deleted_at`).
		Joins("JOIN m_customer_type m_customer_type_temp on m_customer_sub_type_temp.cust_type_temp_id=m_customer_type_temp.cust_type_id").
		Joins("JOIN m_customer_sub_type on m_customer_sub_type_temp.cust_sub_type_id=m_customer_sub_type.cust_sub_type_id").
		Joins("JOIN m_customer_type on m_customer_sub_type.cust_type_id=m_customer_type.cust_type_id").
		Where("cust_sub_type_temp_id = ? AND cust_sub_type_temp_status = ?", pid, status).
		Order("cust_sub_type_temp_created_at desc").
		Find(&customersubtype).Error
	if err != nil {
		return &CustomerSubTypeTempData{}, err
	}
	return &customersubtype, nil
}

func (p *CustomerSubTypeTemp) UpdateCustomerSubTypeTemp(db *gorm.DB, pid uint64) (*CustomerSubTypeTemp, error) {
	var err error
	db = db.Debug().Model(&CustomerSubTypeTemp{}).Where("cust_sub_type_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"cust_type_temp_id":         p.CustomerTypeTempID,
			"cust_sub_type_temp_code":   p.CustomerSubTypeTempCode,
			"cust_sub_type_temp_name":   p.CustomerSubTypeTempName,
			"cust_sub_type_temp_reason": p.CustomerSubTypeTempReason,
		},
	)
	err = db.Debug().Model(&CustomerSubTypeTemp{}).Where("cust_sub_type_temp_id = ?", pid).Error
	if err != nil {
		return &CustomerSubTypeTemp{}, err
	}
	return p, nil
}

func (p *CustomerSubTypeTemp) UpdateCustomerSubTypeTempStatus(db *gorm.DB, pid uint64) (*CustomerSubTypeTemp, error) {
	var err error
	db = db.Debug().Model(&CustomerSubTypeTemp{}).Where("cust_sub_type_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"cust_sub_type_temp_reason": p.CustomerSubTypeTempReason,
			"cust_sub_type_temp_status": p.CustomerSubTypeTempStatus,
		},
	)
	err = db.Debug().Model(&CustomerSubTypeTemp{}).Where("cust_sub_type_temp_id = ?", pid).Error
	if err != nil {
		return &CustomerSubTypeTemp{}, err
	}
	return p, nil
}

func (p *CustomerSubTypeTemp) DeleteCustomerSubTypeTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&CustomerSubTypeTemp{}).Where("cust_sub_type_temp_id = ?", pid).Take(&CustomerSubTypeTemp{}).Delete(&CustomerSubTypeTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("CustomerSubTypeTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================
