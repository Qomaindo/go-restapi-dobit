package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Customer struct {
	CustomerID            uint64     `gorm:"column:cust_id;primary_key;" json:"cust_id"`
	CustomerCode          string     `gorm:"column:cust_code;size:25" json:"cust_code"`
	CustomerKTP           string     `gorm:"column:cust_ktp;size:255" json:"cust_ktp"`
	CustomerPassport      string     `gorm:"column:cust_passport;size:125" json:"cust_passport"`
	CustomerNPWP          string     `gorm:"column:cust_npwp;size:255" json:"cust_npwp"`
	CustomerName          string     `gorm:"column:cust_name;size:255" json:"cust_name"`
	CustomerSex           string     `gorm:"column:cust_sex;size:6" json:"cust_sex"`
	CustomerBirthDate     string     `gorm:"column:cust_birth_date;size:25" json:"cust_birth_date"`
	CustomerBirthPlace    string     `gorm:"column:cust_birth_place;size:255" json:"cust_birth_place"`
	CustomerAge           string     `gorm:"column:cust_age;size:3" json:"cust_age"`
	CustomerPhoneWork     string     `gorm:"column:cust_phone_work;size:25" json:"cust_phone_work"`
	CustomerEmailWork     string     `gorm:"column:cust_email_work;size:125;" json:"cust_email_work"`
	CustomerLastEducation string     `gorm:"column:cust_last_education;size:10;" json:"cust_last_education"`
	CustomerMaritalStatus string     `gorm:"column:cust_marital_status;size:25;" json:"cust_marital_status"`
	CustomerImage         string     `gorm:"column:cust_image;size:225" json:"cust_image"`
	CustomerStatus        uint64     `gorm:"column:cust_status;size:1" json:"cust_status"`
	CustomerCreatedBy     string     `gorm:"column:cust_created_by;size:125" json:"cust_created_by"`
	CustomerCreatedAt     time.Time  `gorm:"column:cust_created_at;default:CURRENT_TIMESTAMP" json:"cust_created_at"`
	CustomerUpdatedBy     string     `gorm:"column:cust_updated_by;size:125" json:"cust_updated_by"`
	CustomerUpdatedAt     *time.Time `gorm:"column:cust_updated_at;default:CURRENT_TIMESTAMP" json:"cust_updated_at"`
	CustomerDeletedBy     string     `gorm:"column:cust_deleted_by;size:125" json:"cust_deleted_by"`
	CustomerDeletedAt     *time.Time `gorm:"column:cust_deleted_at;default:CURRENT_TIMESTAMP" json:"cust_deleted_at"`
}

type ResponseCustomers struct {
	Status  uint64      `json:"status"`
	Message string      `json:"message"`
	Data    *[]Customer `json:"data"`
}

type ResponseCustomer struct {
	Status  uint64    `json:"status"`
	Message string    `json:"message"`
	Data    *Customer `json:"data"`
}
type ResponseProfileImage struct {
	Status  uint64 `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type ResponseCustomerIU struct {
	Status  uint64    `json:"status"`
	Message string    `json:"message"`
	Data    *Customer `json:"data"`
}

type ResponseCustomerDel struct {
	Status  uint64 `json:"status"`
	Message string `json:"message"`
}

func (Customer) TableName() string {
	return "m_customer"
}

func (p *Customer) Prepare() {
	p.CustomerCode = html.EscapeString(strings.TrimSpace(p.CustomerCode))
	p.CustomerKTP = html.EscapeString(strings.TrimSpace(p.CustomerKTP))
	p.CustomerPassport = html.EscapeString(strings.TrimSpace(p.CustomerPassport))
	p.CustomerNPWP = html.EscapeString(strings.TrimSpace(p.CustomerNPWP))
	p.CustomerName = html.EscapeString(strings.TrimSpace(p.CustomerName))
	p.CustomerSex = html.EscapeString(strings.TrimSpace(p.CustomerSex))
	p.CustomerBirthDate = p.CustomerBirthDate
	p.CustomerBirthPlace = html.EscapeString(strings.TrimSpace(p.CustomerBirthPlace))
	p.CustomerAge = html.EscapeString(strings.TrimSpace(p.CustomerAge))
	p.CustomerPhoneWork = html.EscapeString(strings.TrimSpace(p.CustomerPhoneWork))
	p.CustomerEmailWork = html.EscapeString(strings.TrimSpace(p.CustomerEmailWork))
	p.CustomerLastEducation = html.EscapeString(strings.TrimSpace(p.CustomerLastEducation))
	p.CustomerMaritalStatus = html.EscapeString(strings.TrimSpace(p.CustomerMaritalStatus))
	p.CustomerImage = p.CustomerImage
	p.CustomerStatus = p.CustomerStatus
	p.CustomerCreatedBy = p.CustomerCreatedBy
	p.CustomerCreatedAt = time.Now()
	p.CustomerUpdatedBy = p.CustomerUpdatedBy
	p.CustomerUpdatedAt = p.CustomerUpdatedAt
	p.CustomerDeletedBy = p.CustomerDeletedBy
	p.CustomerDeletedAt = p.CustomerDeletedAt
}

func (p *Customer) Validate(action string) error {
	switch strings.ToLower(action) {
	case "register":
		if p.CustomerName == "" {
			return errors.New("Required Customer Name")
		}
		return nil

	default:
		if p.CustomerCode == "" {
			return errors.New("Required Customer Code")
		}
		if p.CustomerKTP == "" {
			return errors.New("Required KTP Number")
		}
		if p.CustomerPassport == "" {
			return errors.New("Required Passport Number")
		}
		if p.CustomerNPWP == "" {
			return errors.New("Required NPWP Number")
		}
		if p.CustomerName == "" {
			return errors.New("Required Customer Name")
		}
		if p.CustomerSex == "" {
			return errors.New("Required Sex")
		}
		if p.CustomerBirthDate == "" {
			return errors.New("Required Birth Date")
		}
		if p.CustomerBirthPlace == "" {
			return errors.New("Required Birth Place")
		}
		if p.CustomerAge == "" {
			return errors.New("Required Age")
		}
		if p.CustomerPhoneWork == "" {
			return errors.New("Required Phone Number Work")
		}
		if p.CustomerEmailWork == "" {
			return errors.New("Required Email Address Work")
		}
		if p.CustomerMaritalStatus == "" {
			return errors.New("Required Marital Status")
		}
		if p.CustomerImage == "" {
			return errors.New("Required Image Customer")
		}
		return nil
	}
}

func (p *Customer) SaveCustomer(db *gorm.DB) (*Customer, error) {
	var err error
	err = db.Debug().Model(&Customer{}).Create(&p).Error
	if err != nil {
		return &Customer{}, err
	}
	return p, nil
}

func (p *Customer) FindAllCustomer(db *gorm.DB) (*[]Customer, error) {
	var err error
	customer := []Customer{}
	err = db.Debug().Model(&Customer{}).Limit(100).Find(&customer).Error
	if err != nil {
		return &[]Customer{}, err
	}
	return &customer, nil
}

func (p *Customer) FindAllCustomerStatus(db *gorm.DB, status uint64) (*[]Customer, error) {
	var err error
	customer := []Customer{}
	err = db.Debug().Model(&Customer{}).Where("cust_status = ?", status).Limit(100).Find(&customer).Error
	if err != nil {
		return &[]Customer{}, err
	}
	return &customer, nil
}

func (p *Customer) FindCustomerByID(db *gorm.DB, pid uint64) (*Customer, error) {
	var err error
	err = db.Debug().Model(&Customer{}).Where("cust_id = ?", pid).Take(&p).Error
	if err != nil {
		return &Customer{}, err
	}
	return p, nil
}

func (p *Customer) FindCustomerStatusByID(db *gorm.DB, pid uint64, status uint64) (*Customer, error) {
	var err error
	err = db.Debug().Model(&Customer{}).Where("cust_id = ? AND cust_status = ?", pid, status).Take(&p).Error
	if err != nil {
		return &Customer{}, err
	}
	return p, nil
}

func (p *Customer) UpdateCustomer(db *gorm.DB, pid uint64) (*Customer, error) {
	var err error
	db = db.Debug().Model(&Customer{}).Where("cust_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"cust_code":           p.CustomerCode,
			"cust_ktp":            p.CustomerKTP,
			"cust_passport":       p.CustomerPassport,
			"cust_npwp":           p.CustomerNPWP,
			"cust_name":           p.CustomerName,
			"cust_sex":            p.CustomerSex,
			"cust_birth_date":     p.CustomerBirthDate,
			"cust_birth_place":    p.CustomerBirthPlace,
			"cust_age":            p.CustomerAge,
			"cust_phone_work":     p.CustomerPhoneWork,
			"cust_email_work":     p.CustomerEmailWork,
			"cust_last_education": p.CustomerLastEducation,
			"cust_marital_status": p.CustomerMaritalStatus,
			"cust_image":          p.CustomerImage,
			"cust_status":         p.CustomerStatus,
			"cust_updated_by":     p.CustomerUpdatedBy,
			"cust_updated_at":     time.Now(),
		},
	)
	err = db.Debug().Model(&Customer{}).Where("cust_id = ?", pid).Error
	if err != nil {
		return &Customer{}, err
	}
	return p, nil
}

func (p *Customer) UpdateProfileCustomer(db *gorm.DB, uid uint64) (*Customer, error) {

	// To hash the password
	var err error
	db = db.Debug().Model(&Customer{}).Where("cust_id = ?", uid).Take(&Customer{}).UpdateColumns(
		map[string]interface{}{
			"cust_name":           p.CustomerName,
			"cust_sex":            p.CustomerSex,
			"cust_birth_date":     p.CustomerBirthDate,
			"cust_birth_place":    p.CustomerBirthPlace,
			"cust_last_education": p.CustomerLastEducation,
			"cust_marital_status": p.CustomerMaritalStatus,
			"cust_age":            p.CustomerAge,
			"cust_ktp":            p.CustomerKTP,
			"cust_passport":       p.CustomerPassport,
			"cust_npwp":           p.CustomerNPWP,
			"cust_image":          p.CustomerImage,
			"cust_updated_by":     p.CustomerUpdatedBy,
			"cust_updated_at":     time.Now(),
		},
	)
	if db.Error != nil {
		return &Customer{}, db.Error
	}
	// This is the display the updated customeruser
	err = db.Debug().Model(&Customer{}).Where("cust_id = ?", uid).Take(&p).Error
	if err != nil {
		return &Customer{}, err
	}
	return p, nil
}

func (p *Customer) UpdateCustomerImage(db *gorm.DB, pid uint64, image string) (*Customer, error) {
	var err error
	db = db.Debug().Model(&Customer{}).Where("cust_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"cust_image":      image,
			"cust_updated_by": p.CustomerUpdatedBy,
			"cust_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&Customer{}).Where("cust_id = ?", pid).Error
	if err != nil {
		return &Customer{}, err
	}
	return p, nil
}

func (p *Customer) UpdateCustomerStatus(db *gorm.DB, pid uint64) (*Customer, error) {
	var err error
	db = db.Debug().Model(&Customer{}).Where("cust_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"cust_status":     p.CustomerStatus,
			"cust_updated_by": p.CustomerUpdatedBy,
			"cust_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&Customer{}).Where("cust_id = ?", pid).Error
	if err != nil {
		return &Customer{}, err
	}
	return p, nil
}

func (p *Customer) UpdateCustomerStatusDelete(db *gorm.DB, pid uint64) (*Customer, error) {
	var err error
	db = db.Debug().Model(&Customer{}).Where("cust_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"cust_status":     3,
			"cust_deleted_by": p.CustomerDeletedBy,
			"cust_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&Customer{}).Where("cust_id = ?", pid).Error
	if err != nil {
		return &Customer{}, err
	}
	return p, nil
}

func (p *Customer) DeleteCustomer(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&Customer{}).Where("cust_id = ?", pid).Take(&Customer{}).Delete(&Customer{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Customer not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
