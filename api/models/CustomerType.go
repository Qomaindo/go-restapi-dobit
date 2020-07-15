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

type CustomerType struct {
	CustomerTypeID        uint64     `gorm:"column:cust_type_id;primary_key;" json:"cust_type_id"`
	CustomerTypeCode      string     `gorm:"column:cust_type_code" json:"cust_type_code"`
	CustomerTypeName      string     `gorm:"column:cust_type_name" json:"cust_type_name"`
	CustomerTypeStatus    uint64     `gorm:"column:cust_type_status;size:2" json:"cust_type_status"`
	CustomerTypeCreatedBy string     `gorm:"column:cust_type_created_by;size:125" json:"cust_type_created_by"`
	CustomerTypeCreatedAt time.Time  `gorm:"column:cust_type_created_at;default:CURRENT_TIMESTAMP" json:"cust_type_created_at"`
	CustomerTypeUpdatedBy string     `gorm:"column:cust_type_updated_by;size:125" json:"cust_type_updated_by"`
	CustomerTypeUpdatedAt *time.Time `gorm:"column:cust_type_updated_at;default:CURRENT_TIMESTAMP" json:"cust_type_created_at"`
	CustomerTypeDeletedBy string     `gorm:"column:cust_type_deleted_by;size:125" json:"cust_type_deleted_by"`
	CustomerTypeDeletedAt *time.Time `gorm:"column:cust_type_deleted_at;default:CURRENT_TIMESTAMP" json:"cust_type_deleted_at"`
}

type CustomerTypeData struct {
	CustomerTypeID        uint64     `gorm:"column:cust_type_id" json:"cust_type_id"`
	CustomerTypeCode      string     `gorm:"column:cust_type_code" json:"cust_type_code"`
	CustomerTypeName      string     `gorm:"column:cust_type_name" json:"cust_type_name"`
	CustomerTypeStatus    uint64     `gorm:"column:cust_type_status" json:"cust_type_status"`
	CustomerTypeCreatedBy string     `gorm:"column:cust_type_created_by" json:"cust_type_created_by"`
	CustomerTypeCreatedAt time.Time  `gorm:"column:cust_type_created_at" json:"cust_type_created_at"`
	CustomerTypeUpdatedBy string     `gorm:"column:cust_type_updated_by" json:"cust_type_updated_by"`
	CustomerTypeUpdatedAt *time.Time `gorm:"column:cust_type_updated_at" json:"cust_type_updated_at"`
	CustomerTypeDeletedBy string     `gorm:"column:cust_type_deleted_by" json:"cust_type_deleted_by"`
	CustomerTypeDeletedAt *time.Time `gorm:"column:cust_type_deleted_at" json:"cust_type_deleted_at"`
}

type ResponseCustomerTypes struct {
	Status  int                 `json:"status"`
	Message string              `json:"message"`
	Data    *[]CustomerTypeData `json:"data"`
}

type ResponseCustomerType struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    *CustomerTypeData `json:"data"`
}

type ResponseCustomerTypeCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (CustomerType) TableName() string {
	return "m_customer_type"
}

func (CustomerTypeData) TableName() string {
	return "m_customer_type"
}

func (p *CustomerType) Prepare() {
	p.CustomerTypeCode = html.EscapeString(strings.TrimSpace(p.CustomerTypeCode))
	p.CustomerTypeName = html.EscapeString(strings.TrimSpace(p.CustomerTypeName))
	p.CustomerTypeStatus = p.CustomerTypeStatus
	p.CustomerTypeCreatedBy = html.EscapeString(strings.TrimSpace(p.CustomerTypeCreatedBy))
	p.CustomerTypeCreatedAt = time.Now()
	p.CustomerTypeUpdatedBy = html.EscapeString(strings.TrimSpace(p.CustomerTypeUpdatedBy))
	p.CustomerTypeUpdatedAt = p.CustomerTypeUpdatedAt
	p.CustomerTypeDeletedBy = html.EscapeString(strings.TrimSpace(p.CustomerTypeDeletedBy))
	p.CustomerTypeDeletedAt = p.CustomerTypeDeletedAt
}

func (p *CustomerType) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.CustomerTypeCode == "" {
			return errors.New("Required Country")
		}
		if p.CustomerTypeName == "" {
			return errors.New("Required Province")
		}
		return nil
	}
}

func (p *CustomerType) SaveCustomerType(db *gorm.DB) (*CustomerType, error) {
	var err error
	err = db.Debug().Model(&CustomerType{}).Create(&p).Error
	if err != nil {
		return &CustomerType{}, err
	}
	return p, nil
}

func (p *CustomerType) FindAllCustomerType(db *gorm.DB) (*[]CustomerTypeData, error) {
	var err error
	address := []CustomerTypeData{}
	err = db.Debug().Model(&CustomerTypeData{}).Limit(100).Find(&address).Error
	if err != nil {
		return &[]CustomerTypeData{}, err
	}
	return &address, nil
}

func (p *CustomerType) FindAllCustomerTypeStatus(db *gorm.DB, status uint64) (*[]CustomerTypeData, error) {
	var err error
	address := []CustomerTypeData{}
	err = db.Debug().Model(&CustomerTypeData{}).Limit(100).Where("cust_type_status = ?", status).
		Find(&address).Error
	if err != nil {
		return &[]CustomerTypeData{}, err
	}
	return &address, nil
}

func (p *CustomerType) FindCustomerTypeDataByID(db *gorm.DB, pid uint64) (*CustomerType, error) {
	var err error
	err = db.Debug().Model(&CustomerType{}).Where("cust_type_id = ?", pid).Take(&p).Error
	if err != nil {
		return &CustomerType{}, err
	}
	return p, nil
}

func (p *CustomerType) FindCustomerTypeData(db *gorm.DB, pid uint64) (*CustomerTypeData, error) {
	var err error
	customertypedata := CustomerTypeData{}
	err = db.Debug().Model(&CustomerTypeData{}).Limit(100).
		Select(`*`).
		Where("cust_type_id = ?", pid).
		Find(&customertypedata).Error
	if err != nil {
		return &CustomerTypeData{}, err
	}
	return &customertypedata, nil
}

func (p *CustomerType) FindCustomerTypeByID(db *gorm.DB, pid uint64) (*CustomerTypeData, error) {
	var err error
	address := CustomerTypeData{}
	err = db.Debug().Model(&CustomerTypeData{}).Limit(100).Where("cust_type_id = ?", pid).
		Find(&address).Error
	if err != nil {
		return &CustomerTypeData{}, err
	}
	return &address, nil
}

func (p *CustomerType) FindCustomerTypeStatusByID(db *gorm.DB, pid uint64, status uint64) (*CustomerTypeData, error) {
	var err error
	address := CustomerTypeData{}
	err = db.Debug().Model(&CustomerTypeData{}).Limit(100).Where("cust_type_id = ? AND cust_type_status = ?", pid, status).
		Find(&address).Error
	if err != nil {
		return &CustomerTypeData{}, err
	}
	return &address, nil
}

func (p *CustomerType) UpdateCustomerType(db *gorm.DB, pid uint64) (*CustomerType, error) {
	var err error
	db = db.Debug().Model(&CustomerType{}).Where("cust_type_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"cust_type_code":       p.CustomerTypeCode,
			"cust_type_name":       p.CustomerTypeName,
			"cust_type_status":     p.CustomerTypeStatus,
			"cust_type_updated_by": p.CustomerTypeUpdatedBy,
			"cust_type_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&CustomerType{}).Where("cust_type_id = ?", pid).Error
	if err != nil {
		return &CustomerType{}, err
	}
	return p, nil
}

func (p *CustomerType) UpdateCustomerTypeStatus(db *gorm.DB, pid uint64) (*CustomerType, error) {
	var err error
	db = db.Debug().Model(&CustomerType{}).Where("cust_type_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"cust_type_status":     p.CustomerTypeStatus,
			"cust_type_updated_by": p.CustomerTypeUpdatedBy,
			"cust_type_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&CustomerType{}).Where("cust_type_id = ?", pid).Error
	if err != nil {
		return &CustomerType{}, err
	}
	return p, nil
}

func (p *CustomerType) UpdateCustomerTypeStatusOnly(db *gorm.DB, pid uint64) (*CustomerType, error) {
	var err error
	db = db.Debug().Model(&CustomerType{}).Where("cust_type_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"cust_type_status": p.CustomerTypeStatus,
		},
	)
	err = db.Debug().Model(&CustomerType{}).Where("cust_type_id = ?", pid).Error
	if err != nil {
		return &CustomerType{}, err
	}
	return p, nil
}

func (p *CustomerType) UpdateCustomerTypeStatusDelete(db *gorm.DB, pid uint64) (*CustomerType, error) {
	var err error
	db = db.Debug().Model(&CustomerType{}).Where("cust_type_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"cust_type_status":     3,
			"cust_type_deleted_by": p.CustomerTypeDeletedBy,
			"cust_type_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&CustomerType{}).Where("cust_type_id = ?", pid).Error
	if err != nil {
		return &CustomerType{}, err
	}
	return p, nil
}

func (p *CustomerType) DeleteCustomerType(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&CustomerType{}).Where("cust_type_id = ?", pid).Take(&CustomerType{}).Delete(&CustomerType{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("CustomerType not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================
