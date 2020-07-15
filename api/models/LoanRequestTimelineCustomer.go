package models

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type LoanRequestTimelineCust struct {
	LoanRequestTimelineCustID        uint64     `gorm:"column:loan_req_timeline_cust_id;primary_key;" json:"loan_req_timeline_cust_id"`
	LoanRequestID                    uint64     `gorm:"column:loan_req_id;" json:"loan_req_id"`
	LoanRequestTimelineCustTitle     string     `gorm:"column:loan_req_timeline_cust_title" json:"loan_req_timeline_cust_title"`
	LoanRequestTimelineCustDesc      string     `gorm:"column:loan_req_timeline_cust_desc" json:"loan_req_timeline_cust_desc"`
	LoanRequestTimelineCustStatus    uint64     `gorm:"column:loan_req_timeline_cust_status;size:2" json:"loan_req_timeline_cust_status"`
	LoanRequestTimelineCustCreatedBy string     `gorm:"column:loan_req_timeline_cust_created_by" json:"loan_req_timeline_cust_created_by"`
	LoanRequestTimelineCustCreatedAt time.Time  `gorm:"column:loan_req_timeline_cust_created_at;default:CURRENT_TIMESTAMP" json:"loan_req_timeline_cust_created_at"`
	LoanRequestTimelineCustUpdatedBy string     `gorm:"column:loan_req_timeline_cust_updated_by" json:"loan_req_timeline_cust_updated_by"`
	LoanRequestTimelineCustUpdatedAt *time.Time `gorm:"column:loan_req_timeline_cust_updated_at;default:CURRENT_TIMESTAMP" json:"loan_req_timeline_cust_updated_at"`
	LoanRequestTimelineCustDeletedBy string     `gorm:"column:loan_req_timeline_cust_deleted_by" json:"loan_req_timeline_cust_deleted_by"`
	LoanRequestTimelineCustDeletedAt *time.Time `gorm:"column:loan_req_timeline_cust_deleted_at;default:CURRENT_TIMESTAMP" json:"loan_req_timeline_cust_deleted_at"`
}

type LoanRequestTimelineCustData struct {
	LoanRequestTimelineCustID        uint64     `gorm:"column:loan_req_timeline_cust_id;primary_key;" json:"loan_req_timeline_cust_id"`
	LoanRequestID                    uint64     `gorm:"column:loan_req_id;" json:"loan_req_id"`
	LoanRequestTimelineCustTitle     string     `gorm:"column:loan_req_timeline_cust_title" json:"loan_req_timeline_cust_title"`
	LoanRequestTimelineCustDesc      string     `gorm:"column:loan_req_timeline_cust_desc" json:"loan_req_timeline_cust_desc"`
	LoanRequestTimelineCustStatus    uint64     `gorm:"column:loan_req_timeline_cust_status;size:2" json:"loan_req_timeline_cust_status"`
	LoanRequestTimelineCustCreatedBy string     `gorm:"column:loan_req_timeline_cust_created_by" json:"loan_req_timeline_cust_created_by"`
	LoanRequestTimelineCustCreatedAt time.Time  `gorm:"column:loan_req_timeline_cust_created_at;default:CURRENT_TIMESTAMP" json:"loan_req_timeline_cust_created_at"`
	LoanRequestTimelineCustUpdatedBy string     `gorm:"column:loan_req_timeline_cust_updated_by" json:"loan_req_timeline_cust_updated_by"`
	LoanRequestTimelineCustUpdatedAt *time.Time `gorm:"column:loan_req_timeline_cust_updated_at;default:CURRENT_TIMESTAMP" json:"loan_req_timeline_cust_updated_at"`
	LoanRequestTimelineCustDeletedBy string     `gorm:"column:loan_req_timeline_cust_deleted_by" json:"loan_req_timeline_cust_deleted_by"`
	LoanRequestTimelineCustDeletedAt *time.Time `gorm:"co
	func (p *LoanRequestTimelineCust) UpdateLoanRequestTimelineCustDate(db *gorm.DB, pid uint64) (*LoanRequestTimelineCust, error) {
		var err error
		db = db.Debug().Model(&LoanRequestTimelineCust{}).Where("loan_req_id = ?", pid).UpdateColumns(
			map[string]interface{}{
				"loan_req_timeline_cust_title":      p.LoanRequestTimelineCustDate1,
				"loan_req_timeline_cust_desc":       p.LoanRequestTimelineCustDate2,
				"loan_req_timeline_cust_updated_at": time.Now(),
			},
		)
		err = db.Debug().Model(&LoanRequestTimelineCust{}).Where("loan_req_id = ?", pid).Error
		if err != nil {
			return &LoanRequestTimelineCust{}, err
		}
		return p, nil
	}lumn:loan_req_timeline_cust_deleted_at;default:CURRENT_TIMESTAMP" json:"loan_req_timeline_cust_deleted_at"`
}

type ResponseLoanRequestTimelineCusts struct {
	Status  int                            `json:"status"`
	Message string                         `json:"message"`
	Data    *[]LoanRequestTimelineCustData `json:"data"`
}

type ResponseLoanRequestTimelineCust struct {
	Status  int                          `json:"status"`
	Message string                       `json:"message"`
	Data    *LoanRequestTimelineCustData `json:"data"`
}

type ResponseLoanRequestTimelineCustIU struct {
	Status  int                      `json:"status"`
	Message string                   `json:"message"`
	Data    *LoanRequestTimelineCust `json:"data"`
}

type ResponseLoanRequestTimelineCustDel struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (LoanRequestTimelineCust) TableName() string {
	return "t_loan_request_timeline_customer"
}

func (LoanRequestTimelineCustData) TableName() string {
	return "t_loan_request_timeline_customer"
}

func (p *LoanRequestTimelineCust) Prepare() {

	p.LoanRequestID = p.LoanRequestID
	p.LoanRequestTimelineCustTitle = p.LoanRequestTimelineCustTitle
	p.LoanRequestTimelineCustDesc = p.LoanRequestTimelineCustDesc
	p.LoanRequestTimelineCustStatus = p.LoanRequestTimelineCustStatus
	p.LoanRequestTimelineCustCreatedBy = p.LoanRequestTimelineCustCreatedBy
	p.LoanRequestTimelineCustCreatedAt = time.Now()
	p.LoanRequestTimelineCustUpdatedBy = p.LoanRequestTimelineCustUpdatedBy
	p.LoanRequestTimelineCustUpdatedAt = p.LoanRequestTimelineCustUpdatedAt
	p.LoanRequestTimelineCustDeletedBy = p.LoanRequestTimelineCustDeletedBy
	p.LoanRequestTimelineCustDeletedAt = p.LoanRequestTimelineCustDeletedAt
}
func (p *LoanRequestTimelineCust) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		// if p.CountryID == 0 {
		// 	return errors.New("Required Region Country")
		// }
		// if p.ProvinceID == 0 {
		// 	return errors.New("Required Region Province")
		// }
		// if p.RegencyID == 0 {
		// 	return errors.New("Required Region Regency")
		// }
		// if p.DistrictID == 0 {
		// 	return errors.New("Required Region District")
		// }
		// if p.VillageID == 0 {
		// 	return errors.New("Required Region Village")
		// }
		return nil
	}
}

func (p *LoanRequestTimelineCust) SaveLoanRequestTimelineCust(db *gorm.DB) (*LoanRequestTimelineCust, error) {
	var err error
	err = db.Debug().Model(&LoanRequestTimelineCust{}).Create(&p).Error
	if err != nil {
		return &LoanRequestTimelineCust{}, err
	}
	return p, nil
}

func (p *LoanRequestTimelineCust) FindAllLoanRequestTimelineCust(db *gorm.DB) (*[]LoanRequestTimelineCustData, error) {
	var err error
	loan_req_timeline := []LoanRequestTimelineCustData{}
	err = db.Debug().Model(&LoanRequestTimelineCustData{}).Limit(100).
		Select(`t_loan_request_timeline_customer.loan_req_timeline_cust_id,
				t_loan_request.loan_req_id,     
				t_loan_request_timeline_customer.loan_req_timeline_cust_title,
				t_loan_request_timeline_customer.loan_req_timeline_cust_desc,
				t_loan_request_timeline_customer.loan_req_timeline_cust_status,
				t_loan_request_timeline_customer.loan_req_timeline_cust_created_by,
				t_loan_request_timeline_customer.loan_req_timeline_cust_created_at at time zone 'Asia/Jakarta' as loan_req_timeline_cust_created_at,
				t_loan_request_timeline_customer.loan_req_timeline_cust_updated_by,
				t_loan_request_timeline_customer.loan_req_timeline_cust_updated_at at time zone 'Asia/Jakarta' as loan_req_timeline_cust_updated_at,
				t_loan_request_timeline_customer.loan_req_timeline_cust_deleted_by,
				t_loan_request_timeline_customer.loan_req_timeline_cust_deleted_at at time zone 'Asia/Jakarta' as loan_req_timeline_cust_deleted_at`).
		Joins("JOIN t_loan_request on t_loan_request_timeline_customer.loan_req_id=t_loan_request.loan_req_id").
		Find(&loan_req_timeline).Error
	if err != nil {
		return &[]LoanRequestTimelineCustData{}, err
	}
	return &loan_req_timeline, nil
}

func (p *LoanRequestTimelineCust) FindAllLoanRequestTimelineCustStatus(db *gorm.DB, status uint64) (*[]LoanRequestTimelineCustData, error) {
	var err error
	loan_req_timeline := []LoanRequestTimelineCustData{}
	err = db.Debug().Model(&LoanRequestTimelineCustData{}).Limit(100).
		Select(`t_loan_request_timeline_customer.loan_req_timeline_cust_id,
				t_loan_request.loan_req_id,     
				t_loan_request_timeline_customer.loan_req_timeline_cust_title,
				t_loan_request_timeline_customer.loan_req_timeline_cust_desc,
				t_loan_request_timeline_customer.loan_req_timeline_cust_status,
				t_loan_request_timeline_customer.loan_req_timeline_cust_created_by,
				t_loan_request_timeline_customer.loan_req_timeline_cust_created_at at time zone 'Asia/Jakarta' as loan_req_timeline_cust_created_at,
				t_loan_request_timeline_customer.loan_req_timeline_cust_updated_by,
				t_loan_request_timeline_customer.loan_req_timeline_cust_updated_at at time zone 'Asia/Jakarta' as loan_req_timeline_cust_updated_at,
				t_loan_request_timeline_customer.loan_req_timeline_cust_deleted_by,
				t_loan_request_timeline_customer.loan_req_timeline_cust_deleted_at at time zone 'Asia/Jakarta' as loan_req_timeline_cust_deleted_at`).
		Joins("JOIN t_loan_request on t_loan_request_timeline_customer.loan_req_id=t_loan_request.loan_req_id").
		Where("loan_req_timeline_cust_status = ?", status).
		Find(&loan_req_timeline).Error
	if err != nil {
		return &[]LoanRequestTimelineCustData{}, err
	}
	return &loan_req_timeline, nil
}

func (p *LoanRequestTimelineCust) FindLoanRequestTimelineCustDataByID(db *gorm.DB, pid uint64) (*LoanRequestTimelineCust, error) {
	var err error
	err = db.Debug().Model(&LoanRequestTimelineCust{}).Where("loan_req_timeline_cust_id = ?", pid).Take(&p).Error
	if err != nil {
		return &LoanRequestTimelineCust{}, err
	}
	return p, nil
}

func (p *LoanRequestTimelineCust) FindLoanRequestTimelineCustByReqID(db *gorm.DB, pid string) (*[]LoanRequestTimelineCustData, error) {
	var err error
	loan_req_timeline := []LoanRequestTimelineCustData{}
	err = db.Debug().Model(&LoanRequestTimelineCustData{}).Limit(100).
		Select(`t_loan_request_timeline_customer.loan_req_timeline_cust_id,
				t_loan_request.loan_req_id,
				m_customer_user.cust_user_id,
				m_customer_user.cust_user_phone,
				t_loan_request_timeline_customer.loan_req_timeline_cust_title,
				t_loan_request_timeline_customer.loan_req_timeline_cust_desc,
				t_loan_request_timeline_customer.loan_req_timeline_cust_status,
				t_loan_request_timeline_customer.loan_req_timeline_cust_created_at at time zone 'Asia/Jakarta' as loan_req_timeline_cust_created_at,
				t_loan_request_timeline_customer.loan_req_timeline_cust_created_by,
				t_loan_request_timeline_customer.loan_req_timeline_cust_updated_at at time zone 'Asia/Jakarta' as loan_req_timeline_cust_updated_at,
				t_loan_request_timeline_customer.loan_req_timeline_cust_updated_by,
				t_loan_request_timeline_customer.loan_req_timeline_cust_deleted_at at time zone 'Asia/Jakarta' as loan_req_timeline_cust_deleted_at,
				t_loan_request_timeline_customer.loan_req_timeline_cust_deleted_by`).
		Joins("JOIN t_loan_request on t_loan_request_timeline_customer.loan_req_id=t_loan_request.loan_req_id").
		Joins("JOIN m_customer_user on t_loan_request.cust_user_id =m_customer_user.cust_user_id").
		Where("t_loan_request.loan_req_id = ?", pid).
		Order("t_loan_request_timeline_customer.loan_req_timeline_cust_id asc").
		Find(&loan_req_timeline).Error
	if err != nil {
		return &[]LoanRequestTimelineCustData{}, err
	}
	return &loan_req_timeline, nil
}

func (p *LoanRequestTimelineCust) FindLoanRequestTimelineCustByPhone(db *gorm.DB, pid string) (*[]LoanRequestTimelineCustData, error) {
	var err error
	loan_req_timeline := []LoanRequestTimelineCustData{}
	err = db.Debug().Model(&LoanRequestTimelineCustData{}).Limit(100).
		Select(`t_loan_request_timeline_customer.loan_req_timeline_cust_id,
				t_loan_request.loan_req_id,
				m_customer_user.cust_user_id,
				m_customer_user.cust_user_phone,
				t_loan_request_timeline_customer.loan_req_timeline_cust_title,
				t_loan_request_timeline_customer.loan_req_timeline_cust_desc,
				t_loan_request_timeline_customer.loan_req_timeline_cust_status,
				t_loan_request_timeline_customer.loan_req_timeline_cust_created_at at time zone 'Asia/Jakarta' as loan_req_timeline_cust_created_at,
				t_loan_request_timeline_customer.loan_req_timeline_cust_created_by,
				t_loan_request_timeline_customer.loan_req_timeline_cust_updated_at at time zone 'Asia/Jakarta' as loan_req_timeline_cust_updated_at,
				t_loan_request_timeline_customer.loan_req_timeline_cust_updated_by,
				t_loan_request_timeline_customer.loan_req_timeline_cust_deleted_at at time zone 'Asia/Jakarta' as loan_req_timeline_cust_deleted_at,
				t_loan_request_timeline_customer.loan_req_timeline_cust_deleted_by`).
		Joins("JOIN t_loan_request on t_loan_request_timeline_customer.loan_req_id=t_loan_request.loan_req_id").
		Joins("JOIN m_customer_user on t_loan_request.cust_user_id =m_customer_user.cust_user_id").
		Where("m_customer_user.cust_user_phone = ?", pid).
		Find(&loan_req_timeline).Error
	if err != nil {
		return &[]LoanRequestTimelineCustData{}, err
	}
	return &loan_req_timeline, nil
}

func (p *LoanRequestTimelineCust) FindLoanRequestTimelineCustByID(db *gorm.DB, pid uint64) (*LoanRequestTimelineCustData, error) {
	var err error
	loan_req_timeline := LoanRequestTimelineCustData{}
	err = db.Debug().Model(&LoanRequestTimelineCustData{}).Limit(100).
		Select(`t_loan_request_timeline_customer.loan_req_timeline_cust_id,
				t_loan_request.loan_req_id,     
				t_loan_request_timeline_customer.loan_req_timeline_cust_title,
				t_loan_request_timeline_customer.loan_req_timeline_cust_desc,
				t_loan_request_timeline_customer.loan_req_timeline_cust_status,
				t_loan_request_timeline_customer.loan_req_timeline_cust_created_at at time zone 'Asia/Jakarta' as loan_req_timeline_cust_created_at,
				t_loan_request_timeline_customer.loan_req_timeline_cust_created_by,
				t_loan_request_timeline_customer.loan_req_timeline_cust_updated_at at time zone 'Asia/Jakarta' as loan_req_timeline_cust_updated_at,
				t_loan_request_timeline_customer.loan_req_timeline_cust_updated_by,
				t_loan_request_timeline_customer.loan_req_timeline_cust_deleted_at at time zone 'Asia/Jakarta' as loan_req_timeline_cust_deleted_at,
				t_loan_request_timeline_customer.loan_req_timeline_cust_deleted_by`).
		Joins("JOIN t_loan_request on t_loan_request_timeline_customer.loan_req_id=t_loan_request.loan_req_id").
		Where("loan_req_timeline_cust_id = ?", pid).
		Find(&loan_req_timeline).Error
	if err != nil {
		return &LoanRequestTimelineCustData{}, err
	}
	return &loan_req_timeline, nil
}

func (p *LoanRequestTimelineCust) FindLoanRequestTimelineCustActiveByLoanRequestID(db *gorm.DB, pid uint64) (*LoanRequestTimelineCustData, error) {
	var err error
	loan_req_timeline := LoanRequestTimelineCustData{}
	err = db.Debug().Model(&LoanRequestTimelineCustData{}).Limit(100).
		Select(`t_loan_request_timeline_customer.loan_req_timeline_cust_id,
				t_loan_request.loan_req_id,     
				t_loan_request_timeline_customer.loan_req_timeline_cust_title,
				t_loan_request_timeline_customer.loan_req_timeline_place,
				t_loan_request_timeline_customer.loan_req_timeline_cust_status,
				t_loan_request_timeline_customer.loan_req_timeline_cust_created_at at time zone 'Asia/Jakarta' as loan_req_timeline_cust_created_at,
				t_loan_request_timeline_customer.loan_req_timeline_cust_created_by,
				t_loan_request_timeline_customer.loan_req_timeline_cust_updated_at at time zone 'Asia/Jakarta' as loan_req_timeline_cust_updated_at,
				t_loan_request_timeline_customer.loan_req_timeline_cust_updated_by,
				t_loan_request_timeline_customer.loan_req_timeline_cust_deleted_at at time zone 'Asia/Jakarta' as loan_req_timeline_cust_deleted_at,
				t_loan_request_timeline_customer.loan_req_timeline_cust_deleted_by`).
		Joins("JOIN t_loan_request on t_loan_request_timeline_customer.loan_req_id=t_loan_request.loan_req_id").
		Where("t_loan_request_timeline_customer.loan_req_id = ? AND t_loan_request_timeline_customer.loan_req_timeline_cust_status = ?", pid, 1).
		Find(&loan_req_timeline).Error
	if err != nil {
		return &LoanRequestTimelineCustData{}, err
	}
	return &loan_req_timeline, nil
}

func (p *LoanRequestTimelineCust) FindLoanRequestTimelineCustStatusByID(db *gorm.DB, pid uint64, status uint64) (*LoanRequestTimelineCustData, error) {
	var err error
	loan_req_timeline := LoanRequestTimelineCustData{}
	err = db.Debug().Model(&LoanRequestTimelineCustData{}).Limit(100).
		Select(`t_loan_request_timeline_customer.loan_req_timeline_cust_id,
				t_loan_request.loan_req_id,     
				t_loan_request_timeline_customer.loan_req_timeline_cust_title,
				t_loan_request_timeline_customer.loan_req_timeline_place,
				t_loan_request_timeline_customer.loan_req_timeline_cust_status,
				t_loan_request_timeline_customer.loan_req_timeline_cust_created_by,
				t_loan_request_timeline_customer.loan_req_timeline_cust_created_at at time zone 'Asia/Jakarta' as loan_req_timeline_cust_created_at,
				t_loan_request_timeline_customer.loan_req_timeline_cust_updated_by,
				t_loan_request_timeline_customer.loan_req_timeline_cust_updated_at at time zone 'Asia/Jakarta' as loan_req_timeline_cust_updated_at,
				t_loan_request_timeline_customer.loan_req_timeline_cust_deleted_by,
				t_loan_request_timeline_customer.loan_req_timeline_cust_deleted_at at time zone 'Asia/Jakarta' as loan_req_timeline_cust_deleted_at`).
		Joins("JOIN t_loan_request on t_loan_request_timeline_customer.loan_req_id=t_loan_request.loan_req_id").
		Where("loan_req_timeline_cust_id = ? AND loan_req_timeline_cust_status = ?", pid, status).
		Find(&loan_req_timeline).Error
	if err != nil {
		return &LoanRequestTimelineCustData{}, err
	}
	return &loan_req_timeline, nil
}

func (p *LoanRequestTimelineCust) UpdateLoanRequestTimelineCust(db *gorm.DB, pid uint64) (*LoanRequestTimelineCust, error) {

	var err error
	db = db.Debug().Model(&LoanRequestTimelineCust{}).Where("loan_req_timeline_cust_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_req_id":                       p.LoanRequestID,
			"loan_req_timeline_cust_title":      p.LoanRequestTimelineCustTitle,
			"loan_req_timeline_cust_desc":       p.LoanRequestTimelineCustDesc,
			"loan_req_timeline_cust_status":     p.LoanRequestTimelineCustStatus,
			"loan_req_timeline_cust_updated_by": p.LoanRequestTimelineCustUpdatedBy,
			"loan_req_timeline_cust_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanRequestTimelineCust{}).Where("loan_req_timeline_cust_id = ?", pid).Error
	if err != nil {
		return &LoanRequestTimelineCust{}, err
	}
	return p, nil
}

func (p *LoanRequestTimelineCust) UpdateLoanRequestTimelineCustStatus(db *gorm.DB, pid uint64) (*LoanRequestTimelineCust, error) {
	var err error
	db = db.Debug().Model(&LoanRequestTimelineCust{}).Where("loan_req_timeline_cust_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_req_timeline_cust_status":     p.LoanRequestTimelineCustStatus,
			"loan_req_timeline_cust_updated_by": p.LoanRequestTimelineCustUpdatedBy,
			"loan_req_timeline_cust_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanRequestTimelineCust{}).Where("loan_req_timeline_cust_id = ?", pid).Error
	if err != nil {
		return &LoanRequestTimelineCust{}, err
	}
	return p, nil
}

func (p *LoanRequestTimelineCust) UpdateLoanRequestTimelineCustStatusUpdate(db *gorm.DB, pid uint64) (*LoanRequestTimelineCust, error) {
	var err error
	db = db.Debug().Model(&LoanRequestTimelineCust{}).Where("loan_req_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_req_timeline_cust_status":     p.LoanRequestTimelineCustStatus,
			"loan_req_timeline_cust_updated_by": p.LoanRequestTimelineCustUpdatedBy,
			"loan_req_timeline_cust_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanRequestTimelineCust{}).Where("loan_req_id = ?", pid).Error
	if err != nil {
		return &LoanRequestTimelineCust{}, err
	}
	return p, nil
}

func (p *LoanRequestTimelineCust) UpdateLoanRequestTimelineCustStatusDelete(db *gorm.DB, pid uint64) (*LoanRequestTimelineCust, error) {
	var err error
	db = db.Debug().Model(&LoanRequestTimelineCust{}).Where("loan_req_timeline_cust_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_req_timeline_cust_status":     3,
			"loan_req_timeline_cust_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanRequestTimelineCust{}).Where("loan_req_timeline_cust_id = ?", pid).Error
	if err != nil {
		return &LoanRequestTimelineCust{}, err
	}
	return p, nil
}

func (p *LoanRequestTimelineCust) DeleteLoanRequestTimelineCust(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&LoanRequestTimelineCust{}).Where("loan_req_timeline_cust_id = ?", pid).Take(&LoanRequestTimelineCust{}).Delete(&LoanRequestTimelineCust{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("LoanRequestTimelineCust not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
