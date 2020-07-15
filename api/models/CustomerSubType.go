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

type CustomerSubType struct {
	CustomerSubTypeID        uint64     `gorm:"column:cust_sub_type_id;primary_key;" json:"cust_sub_type_id"`
	CustomerTypeID           uint64     `gorm:"column:cust_type_id" json:"cust_type_id"`
	CustomerSubTypeCode      string     `gorm:"column:cust_sub_type_code" json:"cust_sub_type_code"`
	CustomerSubTypeName      string     `gorm:"column:cust_sub_type_name" json:"cust_sub_type_name"`
	CustomerSubTypeStatus    uint64     `gorm:"column:cust_sub_type_status;size:2" json:"cust_sub_type_status"`
	CustomerSubTypeCreatedBy string     `gorm:"column:cust_sub_type_created_by;size:125" json:"cust_sub_type_created_by"`
	CustomerSubTypeCreatedAt time.Time  `gorm:"column:cust_sub_type_created_at;default:CURRENT_TIMESTAMP" json:"cust_sub_type_created_at"`
	CustomerSubTypeUpdatedBy string     `gorm:"column:cust_sub_type_updated_by;size:125" json:"cust_sub_type_updated_by"`
	CustomerSubTypeUpdatedAt *time.Time `gorm:"column:cust_sub_type_updated_at;default:CURRENT_TIMESTAMP" json:"cust_sub_type_created_at"`
	CustomerSubTypeDeletedBy string     `gorm:"column:cust_sub_type_deleted_by;size:125" json:"cust_sub_type_deleted_by"`
	CustomerSubTypeDeletedAt *time.Time `gorm:"column:cust_sub_type_deleted_at;default:CURRENT_TIMESTAMP" json:"cust_sub_type_deleted_at"`
}

type CustomerSubTypeData struct {
	CustomerSubTypeID        uint64     `gorm:"column:cust_sub_type_id" json:"cust_sub_type_id"`
	CustomerTypeID           uint64     `gorm:"column:cust_type_id" json:"cust_type_id"`
	CustomerTypeName         string     `gorm:"column:cust_type_name" json:"cust_type_name"`
	CustomerSubTypeCode      string     `gorm:"column:cust_sub_type_code" json:"cust_sub_type_code"`
	CustomerSubTypeName      string     `gorm:"column:cust_sub_type_name" json:"cust_sub_type_name"`
	CustomerSubTypeStatus    uint64     `gorm:"column:cust_sub_type_status" json:"cust_sub_type_status"`
	CustomerSubTypeCreatedBy string     `gorm:"column:cust_sub_type_created_by" json:"cust_sub_type_created_by"`
	CustomerSubTypeCreatedAt time.Time  `gorm:"column:cust_sub_type_created_at" json:"cust_sub_type_created_at"`
	CustomerSubTypeUpdatedBy string     `gorm:"column:cust_sub_type_updated_by" json:"cust_sub_type_updated_by"`
	CustomerSubTypeUpdatedAt *time.Time `gorm:"column:cust_sub_type_updated_at" json:"cust_sub_type_updated_at"`
	CustomerSubTypeDeletedBy string     `gorm:"column:cust_sub_type_deleted_by" json:"cust_sub_type_deleted_by"`
	CustomerSubTypeDeletedAt *time.Time `gorm:"column:cust_sub_type_deleted_at" json:"cust_sub_type_deleted_at"`
}

type ResponseCustomerSubTypes struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    *[]CustomerSubTypeData `json:"data"`
}

type ResponseCustomerSubType struct {
	Status  int                  `json:"status"`
	Message string               `json:"message"`
	Data    *CustomerSubTypeData `json:"data"`
}

type ResponseCustomerSubTypeCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (CustomerSubType) TableName() string {
	return "m_customer_sub_type"
}

func (CustomerSubTypeData) TableName() string {
	return "m_customer_sub_type"
}

func (p *CustomerSubType) Prepare() {
	p.CustomerTypeID = p.CustomerTypeID
	p.CustomerSubTypeCode = html.EscapeString(strings.TrimSpace(p.CustomerSubTypeCode))
	p.CustomerSubTypeName = html.EscapeString(strings.TrimSpace(p.CustomerSubTypeName))
	p.CustomerSubTypeStatus = p.CustomerSubTypeStatus
	p.CustomerSubTypeCreatedBy = html.EscapeString(strings.TrimSpace(p.CustomerSubTypeCreatedBy))
	p.CustomerSubTypeCreatedAt = time.Now()
	p.CustomerSubTypeUpdatedBy = html.EscapeString(strings.TrimSpace(p.CustomerSubTypeUpdatedBy))
	p.CustomerSubTypeUpdatedAt = p.CustomerSubTypeUpdatedAt
	p.CustomerSubTypeDeletedBy = html.EscapeString(strings.TrimSpace(p.CustomerSubTypeDeletedBy))
	p.CustomerSubTypeDeletedAt = p.CustomerSubTypeDeletedAt
}

func (p *CustomerSubType) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.CustomerTypeID == 0 {
			return errors.New("Required Country")
		}
		if p.CustomerSubTypeCode == "" {
			return errors.New("Required Province")
		}
		if p.CustomerSubTypeName == "" {
			return errors.New("Required Regency")
		}
		return nil
	}
}

func (p *CustomerSubType) SaveCustomerSubType(db *gorm.DB) (*CustomerSubType, error) {
	var err error
	err = db.Debug().Model(&CustomerSubType{}).Create(&p).Error
	if err != nil {
		return &CustomerSubType{}, err
	}
	return p, nil
}

func (p *CustomerSubType) FindAllCustomerSubType(db *gorm.DB) (*[]CustomerSubTypeData, error) {
	var err error
	customersubtype := []CustomerSubTypeData{}
	err = db.Debug().Model(&CustomerSubTypeData{}).Limit(100).
		Select(`m_customer_sub_type.cust_sub_type_id,
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
		Joins("JOIN m_customer_type on m_customer_sub_type.cust_type_id=m_customer_type.cust_type_id").
		Find(&customersubtype).Error
	if err != nil {
		return &[]CustomerSubTypeData{}, err
	}
	return &customersubtype, nil
}

func (p *CustomerSubType) FindAllCustomerSubTypeStatus(db *gorm.DB, status uint64) (*[]CustomerSubTypeData, error) {
	var err error
	customersubtype := []CustomerSubTypeData{}
	err = db.Debug().Model(&CustomerSubTypeData{}).Limit(100).
		Select(`m_customer_sub_type.cust_sub_type_id,
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
		Joins("JOIN m_customer_type on m_customer_sub_type.cust_type_id=m_customer_type.cust_type_id").
		Where("cust_sub_type_status = ?", status).
		Find(&customersubtype).Error
	if err != nil {
		return &[]CustomerSubTypeData{}, err
	}
	return &customersubtype, nil
}

func (p *CustomerSubType) FindCustomerSubTypeDataByID(db *gorm.DB, pid uint64) (*CustomerSubType, error) {
	var err error
	err = db.Debug().Model(&CustomerSubType{}).Where("cust_sub_type_id = ?", pid).Take(&p).Error
	if err != nil {
		return &CustomerSubType{}, err
	}
	return p, nil
}

func (p *CustomerSubType) FindCustomerSubTypeByID(db *gorm.DB, pid uint64) (*CustomerSubTypeData, error) {
	var err error
	customersubtype := CustomerSubTypeData{}
	err = db.Debug().Model(&CustomerSubTypeData{}).Limit(100).
		Select(`m_customer_sub_type.cust_sub_type_id,
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
		Joins("JOIN m_customer_type on m_customer_sub_type.cust_type_id=m_customer_type.cust_type_id").
		Where("cust_sub_type_id = ?", pid).
		Find(&customersubtype).Error
	if err != nil {
		return &CustomerSubTypeData{}, err
	}
	return &customersubtype, nil
}

func (p *CustomerSubType) FindCustomerSubTypeStatusByID(db *gorm.DB, pid uint64, status uint64) (*CustomerSubTypeData, error) {
	var err error
	customersubtype := CustomerSubTypeData{}
	err = db.Debug().Model(&CustomerSubTypeData{}).Limit(100).
		Select(`m_customer_sub_type.cust_sub_type_id,
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
		Joins("JOIN m_customer_type on m_customer_sub_type.cust_type_id=m_customer_type.cust_type_id").
		Where("cust_sub_type_id = ? AND cust_sub_type_status = ?", pid, status).
		Find(&customersubtype).Error
	if err != nil {
		return &CustomerSubTypeData{}, err
	}
	return &customersubtype, nil
}

func (p *CustomerSubType) UpdateCustomerSubType(db *gorm.DB, pid uint64) (*CustomerSubType, error) {
	var err error
	db = db.Debug().Model(&CustomerSubType{}).Where("cust_sub_type_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"cust_type_id":             p.CustomerTypeID,
			"cust_sub_type_code":       p.CustomerSubTypeCode,
			"cust_sub_type_name":       p.CustomerSubTypeName,
			"cust_sub_type_status":     p.CustomerSubTypeStatus,
			"cust_sub_type_updated_by": p.CustomerSubTypeUpdatedBy,
			"cust_sub_type_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&CustomerSubType{}).Where("cust_sub_type_id = ?", pid).Error
	if err != nil {
		return &CustomerSubType{}, err
	}
	return p, nil
}

func (p *CustomerSubType) UpdateCustomerSubTypeStatus(db *gorm.DB, pid uint64) (*CustomerSubType, error) {
	var err error
	db = db.Debug().Model(&CustomerSubType{}).Where("cust_sub_type_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"cust_sub_type_status":     p.CustomerSubTypeStatus,
			"cust_sub_type_updated_by": p.CustomerSubTypeUpdatedBy,
			"cust_sub_type_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&CustomerSubType{}).Where("cust_sub_type_id = ?", pid).Error
	if err != nil {
		return &CustomerSubType{}, err
	}
	return p, nil
}

func (p *CustomerSubType) UpdateCustomerSubTypeStatusOnly(db *gorm.DB, pid uint64) (*CustomerSubType, error) {
	var err error
	db = db.Debug().Model(&CustomerSubType{}).Where("cust_sub_type_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"cust_sub_type_status": p.CustomerSubTypeStatus,
		},
	)
	err = db.Debug().Model(&CustomerSubType{}).Where("cust_sub_type_id = ?", pid).Error
	if err != nil {
		return &CustomerSubType{}, err
	}
	return p, nil
}

func (p *CustomerSubType) UpdateCustomerSubTypeStatusDelete(db *gorm.DB, pid uint64) (*CustomerSubType, error) {
	var err error
	db = db.Debug().Model(&CustomerSubType{}).Where("cust_sub_type_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"cust_sub_type_status":     3,
			"cust_sub_type_deleted_by": p.CustomerSubTypeDeletedBy,
			"cust_sub_type_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&CustomerSubType{}).Where("cust_sub_type_id = ?", pid).Error
	if err != nil {
		return &CustomerSubType{}, err
	}
	return p, nil
}

func (p *CustomerSubType) DeleteCustomerSubType(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&CustomerSubType{}).Where("cust_sub_type_id = ?", pid).Take(&CustomerSubType{}).Delete(&CustomerSubType{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("CustomerSubType not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

func (p *CustomerSubType) FindCustomerSubTypeByCustomerTypeID(db *gorm.DB, pid uint64) (*[]CustomerSubTypeData, error) {
	var err error
	customersubtype := []CustomerSubTypeData{}
	err = db.Debug().Model(&CustomerSubTypeData{}).Limit(100).
		Select(`m_customer_sub_type.cust_sub_type_id,
				m_customer_type.cust_type_id,
				m_customer_type.cust_type_name,
				m_customer_sub_type.cust_sub_type_code,
				m_customer_sub_type.cust_sub_type_name,
				m_customer_sub_type.cust_sub_type_status,
				m_customer_sub_type.cust_sub_type_created_by,
				m_customer_sub_type.cust_sub_type_created_at,
				m_customer_sub_type.cust_sub_type_updated_by,
				m_customer_sub_type.cust_sub_type_updated_at,
				m_customer_sub_type.cust_sub_type_deleted_by,
				m_customer_sub_type.cust_sub_type_deleted_at`).
		Joins("JOIN m_customer_type on m_customer_sub_type.cust_type_id=m_customer_type.cust_type_id").
		Where("m_customer_type.cust_type_id = ?", pid).
		Find(&customersubtype).Error
	if err != nil {
		return &[]CustomerSubTypeData{}, err
	}
	return &customersubtype, nil
}
