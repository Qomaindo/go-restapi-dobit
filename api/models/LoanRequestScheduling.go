package models

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type LoanRequestScheduling struct {
	LoanRequestSchedulingID        uint64     `gorm:"column:loan_req_schdl_id;primary_key;" json:"loan_req_schdl_id"`
	LoanRequestID                  uint64     `gorm:"column:loan_req_id;" json:"loan_req_id"`
	LoanRequestSchedulingDate1     string     `gorm:"column:loan_req_schdl_date_1" json:"loan_req_schdl_date_1"`
	LoanRequestSchedulingDate2     string     `gorm:"column:loan_req_schdl_date_2" json:"loan_req_schdl_date_2"`
	LoanRequestSchedulingTime      string     `gorm:"column:loan_req_schdl_time;size:25" json:"loan_req_schdl_time"`
	LoanRequestSchedulingPlace     string     `gorm:"column:loan_req_schdl_place;size:255" json:"loan_req_schdl_place"`
	LoanRequestSchedulingResult    string     `gorm:"column:loan_req_schdl_result;size:2" json:"loan_req_schdl_result"`
	LoanRequestSchedulingStatus    uint64     `gorm:"column:loan_req_schdl_status;size:2" json:"loan_req_schdl_status"`
	LoanRequestSchedulingCreatedAt time.Time  `gorm:"column:loan_req_schdl_created_at;default:CURRENT_TIMESTAMP" json:"loan_req_schdl_created_at"`
	LoanRequestSchedulingUpdatedAt *time.Time `gorm:"column:loan_req_schdl_updated_at;default:CURRENT_TIMESTAMP" json:"loan_req_schdl_created_at"`
	LoanRequestSchedulingDeletedAt *time.Time `gorm:"column:loan_req_schdl_deleted_at;default:CURRENT_TIMESTAMP" json:"loan_req_schdl_deleted_at"`
}

type LoanRequestSchedulingData struct {
	LoanRequestSchedulingID        uint64     `gorm:"column:loan_req_schdl_id;primary_key;" json:"loan_req_schdl_id"`
	LoanRequestID                  uint64     `gorm:"column:loan_req_id;" json:"loan_req_id"`
	ProductSubCategoryID           uint64     `gorm:"column:prod_sub_ctgry_id" json:"prod_sub_ctgry_id"`
	ProductSubCategoryName         string     `gorm:"column:prod_sub_ctgry_name" json:"prod_sub_ctgry_name"`
	CustUserID                     uint64     `gorm:"column:cust_user_id" json:"cust_user_id"`
	CustomerName                   string     `gorm:"column:cust_name" json:"cust_name"`
	CustomerUserPhone              string     `gorm:"column:cust_user_phone;size:20" json:"cust_user_phone" validate:"length(11|13)"`
	LoanRequestAmount              uint64     `gorm:"column:loan_req_amount" json:"loan_req_amount"`
	LoanRequestLong                uint64     `gorm:"column:loan_req_long" json:"loan_req_long"`
	LoanRequestSchedulingDate1     string     `gorm:"column:loan_req_schdl_date_1" json:"loan_req_schdl_date_1"`
	LoanRequestSchedulingDate2     string     `gorm:"column:loan_req_schdl_date_2" json:"loan_req_schdl_date_2"`
	LoanRequestSchedulingTime      string     `gorm:"column:loan_req_schdl_time;size:25" json:"loan_req_schdl_time"`
	LoanRequestSchedulingPlace     string     `gorm:"column:loan_req_schdl_place;size:255" json:"loan_req_schdl_place"`
	LoanRequestSchedulingResult    string     `gorm:"column:loan_req_schdl_result;size:2" json:"loan_req_schdl_result"`
	LoanRequestSchedulingStatus    uint64     `gorm:"column:loan_req_schdl_status;size:2" json:"loan_req_schdl_status"`
	LoanRequestSchedulingCreatedAt time.Time  `gorm:"column:loan_req_schdl_created_at;default:CURRENT_TIMESTAMP" json:"loan_req_schdl_created_at"`
	LoanRequestSchedulingUpdatedAt *time.Time `gorm:"column:loan_req_schdl_updated_at;default:CURRENT_TIMESTAMP" json:"loan_req_schdl_created_at"`
	LoanRequestSchedulingDeletedAt *time.Time `gorm:"column:loan_req_schdl_deleted_at;default:CURRENT_TIMESTAMP" json:"loan_req_schdl_deleted_at"`
}

type ResponseLoanRequestSchedulings struct {
	Status  int                          `json:"status"`
	Message string                       `json:"message"`
	Data    *[]LoanRequestSchedulingData `json:"data"`
}

type ResponseLoanRequestScheduling struct {
	Status  int                        `json:"status"`
	Message string                     `json:"message"`
	Data    *LoanRequestSchedulingData `json:"data"`
}

type ResponseLoanRequestSchedulingIU struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    *LoanRequestScheduling `json:"data"`
}

type ResponseLoanRequestSchedulingDel struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (LoanRequestScheduling) TableName() string {
	return "t_loan_request_scheduling"
}

func (LoanRequestSchedulingData) TableName() string {
	return "t_loan_request_scheduling"
}

func (p *LoanRequestScheduling) Prepare() {
	p.LoanRequestID = p.LoanRequestID
	p.LoanRequestSchedulingDate1 = p.LoanRequestSchedulingDate1
	p.LoanRequestSchedulingDate2 = p.LoanRequestSchedulingDate2
	p.LoanRequestSchedulingTime = p.LoanRequestSchedulingTime
	p.LoanRequestSchedulingPlace = p.LoanRequestSchedulingPlace
	p.LoanRequestSchedulingResult = p.LoanRequestSchedulingResult
	p.LoanRequestSchedulingStatus = p.LoanRequestSchedulingStatus
	p.LoanRequestSchedulingCreatedAt = time.Now()
	p.LoanRequestSchedulingUpdatedAt = p.LoanRequestSchedulingUpdatedAt
	p.LoanRequestSchedulingDeletedAt = p.LoanRequestSchedulingDeletedAt
}
func (p *LoanRequestScheduling) Validate(action string) error {
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

func (p *LoanRequestScheduling) SaveLoanRequestScheduling(db *gorm.DB) (*LoanRequestScheduling, error) {
	var err error
	err = db.Debug().Model(&LoanRequestScheduling{}).Create(&p).Error
	if err != nil {
		return &LoanRequestScheduling{}, err
	}
	return p, nil
}

func (p *LoanRequestScheduling) FindAllLoanRequestScheduling(db *gorm.DB) (*[]LoanRequestSchedulingData, error) {
	var err error
	loan_req_schdl := []LoanRequestSchedulingData{}
	err = db.Debug().Model(&LoanRequestSchedulingData{}).Limit(100).
		Select(`t_loan_request_scheduling.loan_req_schdl_id,
				t_loan_request.loan_req_id,     
				t_loan_request_scheduling.loan_req_schdl_date_1,
				t_loan_request_scheduling.loan_req_schdl_date_2,
				t_loan_request_scheduling.loan_req_schdl_time,
				t_loan_request_scheduling.loan_req_schdl_place,
				t_loan_request_scheduling.loan_req_schdl_status,
				t_loan_request_scheduling.loan_req_schdl_result,
				t_loan_request_scheduling.loan_req_schdl_created_at at time zone 'Asia/Jakarta' as loan_req_schdl_created_at,
				t_loan_request_scheduling.loan_req_schdl_updated_at at time zone 'Asia/Jakarta' as loan_req_schdl_updated_at,
				t_loan_request_scheduling.loan_req_schdl_deleted_at at time zone 'Asia/Jakarta' as loan_req_schdl_deleted_at`).
		Joins("JOIN t_loan_request on t_loan_request_scheduling.loan_req_id=t_loan_request.loan_req_id").
		Find(&loan_req_schdl).Error
	if err != nil {
		return &[]LoanRequestSchedulingData{}, err
	}
	return &loan_req_schdl, nil
}

func (p *LoanRequestScheduling) FindAllLoanRequestSchedulingByCutstomerID(db *gorm.DB, uid uint64) (*[]LoanRequestSchedulingData, error) {
	var err error
	loan_req_schdl := []LoanRequestSchedulingData{}
	err = db.Debug().Model(&LoanRequestSchedulingData{}).Limit(100).
		Select(`t_loan_request_scheduling.loan_req_schdl_id,
				t_loan_request.loan_req_id,     
				t_loan_request.loan_req_amount,
				t_loan_request.loan_req_long,
				m_product_sub_category.prod_sub_ctgry_id, 	
				m_product_sub_category.prod_sub_ctgry_name, 
				m_customer_user.cust_user_phone,
				t_loan_request.cust_user_id,
				m_customer.cust_name,	     
				t_loan_request_scheduling.loan_req_schdl_date_1,
				t_loan_request_scheduling.loan_req_schdl_date_2,
				t_loan_request_scheduling.loan_req_schdl_time,
				t_loan_request_scheduling.loan_req_schdl_place,
				t_loan_request_scheduling.loan_req_schdl_status,
				t_loan_request_scheduling.loan_req_schdl_result,
				t_loan_request_scheduling.loan_req_schdl_created_at at time zone 'Asia/Jakarta' as loan_req_schdl_created_at,
				t_loan_request_scheduling.loan_req_schdl_updated_at at time zone 'Asia/Jakarta' as loan_req_schdl_updated_at,
				t_loan_request_scheduling.loan_req_schdl_deleted_at at time zone 'Asia/Jakarta' as loan_req_schdl_deleted_at`).
		Joins("JOIN t_loan_request on t_loan_request_scheduling.loan_req_id=t_loan_request.loan_req_id").
		Joins("JOIN m_product_sub_category on t_loan_request.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Joins("JOIN m_customer_user on t_loan_request.cust_user_id=m_customer_user.cust_user_id").
		Joins("JOIN m_customer on m_customer.cust_id=m_customer_user.cust_id").
		Where("m_customer.cust_id = ?", uid).
		Order("t_loan_request_scheduling.loan_req_schdl_id desc").
		Find(&loan_req_schdl).Error
	if err != nil {
		return &[]LoanRequestSchedulingData{}, err
	}
	return &loan_req_schdl, nil
}

func (p *LoanRequestScheduling) FindAllLoanRequestSchedulingByMitraUserID(db *gorm.DB, uid uint64) (*[]LoanRequestSchedulingData, error) {
	var err error
	loan_req_schdl := []LoanRequestSchedulingData{}
	err = db.Debug().Model(&LoanRequestSchedulingData{}).Limit(100).
		Select(`t_loan_request_scheduling.loan_req_schdl_id,
				t_loan_request.loan_req_id,     
				t_loan_request.loan_req_amount,
				t_loan_request.loan_req_long,
				m_product_sub_category.prod_sub_ctgry_id, 	
				m_product_sub_category.prod_sub_ctgry_name, 
				m_customer_user.cust_user_phone,
				t_loan_request.cust_user_id,
				m_customer.cust_name,	     
				t_loan_request_scheduling.loan_req_schdl_date_1,
				t_loan_request_scheduling.loan_req_schdl_date_2,
				t_loan_request_scheduling.loan_req_schdl_time,
				t_loan_request_scheduling.loan_req_schdl_place,
				t_loan_request_scheduling.loan_req_schdl_status,
				t_loan_request_scheduling.loan_req_schdl_result,
				t_loan_request_scheduling.loan_req_schdl_created_at at time zone 'Asia/Jakarta' as loan_req_schdl_created_at,
				t_loan_request_scheduling.loan_req_schdl_updated_at at time zone 'Asia/Jakarta' as loan_req_schdl_updated_at,
				t_loan_request_scheduling.loan_req_schdl_deleted_at at time zone 'Asia/Jakarta' as loan_req_schdl_deleted_at`).
		Joins("JOIN t_loan_request on t_loan_request_scheduling.loan_req_id=t_loan_request.loan_req_id").
		Joins("JOIN m_product_sub_category on t_loan_request.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Joins("JOIN m_mitra_user on t_loan_request.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_customer_user on t_loan_request.cust_user_id=m_customer_user.cust_user_id").
		Joins("JOIN m_customer on m_customer.cust_id=m_customer_user.cust_id").
		Where("m_mitra_user.mitra_user_id = ?", uid).
		Order("t_loan_request_scheduling.loan_req_schdl_id desc").
		Find(&loan_req_schdl).Error
	if err != nil {
		return &[]LoanRequestSchedulingData{}, err
	}
	return &loan_req_schdl, nil
}

func (p *LoanRequestScheduling) FindAllLoanRequestSchedulingStatus(db *gorm.DB, status uint64) (*[]LoanRequestSchedulingData, error) {
	var err error
	loan_req_schdl := []LoanRequestSchedulingData{}
	err = db.Debug().Model(&LoanRequestSchedulingData{}).Limit(100).
		Select(`t_loan_request_scheduling.loan_req_schdl_id,
				t_loan_request.loan_req_id,     
				t_loan_request_scheduling.loan_req_schdl_date_1,
				t_loan_request_scheduling.loan_req_schdl_date_2,
				t_loan_request_scheduling.loan_req_schdl_time,
				t_loan_request_scheduling.loan_req_schdl_place,
				t_loan_request_scheduling.loan_req_schdl_status,
				t_loan_request_scheduling.loan_req_schdl_result,
				t_loan_request_scheduling.loan_req_schdl_created_at at time zone 'Asia/Jakarta' as loan_req_schdl_created_at,
				t_loan_request_scheduling.loan_req_schdl_updated_at at time zone 'Asia/Jakarta' as loan_req_schdl_updated_at,
				t_loan_request_scheduling.loan_req_schdl_deleted_at at time zone 'Asia/Jakarta' as loan_req_schdl_deleted_at`).
		Joins("JOIN t_loan_request on t_loan_request_scheduling.loan_req_id=t_loan_request.loan_req_id").
		Where("loan_req_schdl_status = ?", status).
		Find(&loan_req_schdl).Error
	if err != nil {
		return &[]LoanRequestSchedulingData{}, err
	}
	return &loan_req_schdl, nil
}

func (p *LoanRequestScheduling) FindLoanRequestSchedulingDataByID(db *gorm.DB, pid uint64) (*LoanRequestScheduling, error) {
	var err error
	err = db.Debug().Model(&LoanRequestScheduling{}).Where("loan_req_schdl_id = ?", pid).Take(&p).Error
	if err != nil {
		return &LoanRequestScheduling{}, err
	}
	return p, nil
}

func (p *LoanRequestScheduling) FindLoanRequestSchedulingByID(db *gorm.DB, pid uint64) (*LoanRequestSchedulingData, error) {
	var err error
	loan_req_schdl := LoanRequestSchedulingData{}
	err = db.Debug().Model(&LoanRequestSchedulingData{}).Limit(100).
		Select(`t_loan_request_scheduling.loan_req_schdl_id,
				t_loan_request.loan_req_id,     
				t_loan_request_scheduling.loan_req_schdl_date_1,
				t_loan_request_scheduling.loan_req_schdl_date_2,
				t_loan_request_scheduling.loan_req_schdl_time,
				t_loan_request_scheduling.loan_req_schdl_place,
				t_loan_request_scheduling.loan_req_schdl_status,
				t_loan_request_scheduling.loan_req_schdl_result,
				t_loan_request_scheduling.loan_req_schdl_deleted_at at time zone 'Asia/Jakarta' as loan_req_schdl_deleted_at`).
		Joins("JOIN t_loan_request on t_loan_request_scheduling.loan_req_id=t_loan_request.loan_req_id").
		Where("loan_req_schdl_id = ?", pid).
		Find(&loan_req_schdl).Error
	if err != nil {
		return &LoanRequestSchedulingData{}, err
	}
	return &loan_req_schdl, nil
}

func (p *LoanRequestScheduling) FindLoanRequestSchedulingActiveByLoanRequestID(db *gorm.DB, pid uint64) (*LoanRequestSchedulingData, error) {
	var err error
	loan_req_schdl := LoanRequestSchedulingData{}
	err = db.Debug().Model(&LoanRequestSchedulingData{}).Limit(100).
		Select(`t_loan_request_scheduling.loan_req_schdl_id,
				t_loan_request.loan_req_id,     
				t_loan_request_scheduling.loan_req_schdl_date_1,
				t_loan_request_scheduling.loan_req_schdl_date_2,
				t_loan_request_scheduling.loan_req_schdl_time,
				t_loan_request_scheduling.loan_req_schdl_place,
				t_loan_request_scheduling.loan_req_schdl_status,
				t_loan_request_scheduling.loan_req_schdl_result,
				t_loan_request_scheduling.loan_req_schdl_deleted_at at time zone 'Asia/Jakarta' as loan_req_schdl_deleted_at`).
		Joins("JOIN t_loan_request on t_loan_request_scheduling.loan_req_id=t_loan_request.loan_req_id").
		Where("t_loan_request_scheduling.loan_req_id = ? AND t_loan_request_scheduling.loan_req_schdl_status = ?", pid, 1).
		Find(&loan_req_schdl).Error
	if err != nil {
		return &LoanRequestSchedulingData{}, err
	}
	return &loan_req_schdl, nil
}

func (p *LoanRequestScheduling) FindLoanRequestSchedulingActiveByLoanRequestByCustomerID(db *gorm.DB, uid uint64) (*LoanRequestSchedulingData, error) {
	var err error
	loan_req_schdl := LoanRequestSchedulingData{}
	err = db.Debug().Model(&LoanRequestSchedulingData{}).Limit(100).
		Select(`t_loan_request_scheduling.loan_req_schdl_id,
				t_loan_request.loan_req_id,     
				t_loan_request_scheduling.loan_req_schdl_date_1,
				t_loan_request_scheduling.loan_req_schdl_date_2,
				t_loan_request_scheduling.loan_req_schdl_time,
				t_loan_request_scheduling.loan_req_schdl_place,
				t_loan_request_scheduling.loan_req_schdl_status,
				t_loan_request_scheduling.loan_req_schdl_result,
				t_loan_request_scheduling.loan_req_schdl_deleted_at at time zone 'Asia/Jakarta' as loan_req_schdl_deleted_at`).
		Joins("JOIN t_loan_request on t_loan_request_scheduling.loan_req_id=t_loan_request.loan_req_id").
		Joins("JOIN m_customer_user on t_loan_request.cust_user_id=m_customer_user.cust_user_id").
		Where("m_customer_user.cust_id = ? AND t_loan_request_scheduling.loan_req_schdl_status = ?", uid, 1).
		Find(&loan_req_schdl).Error
	if err != nil {
		return &LoanRequestSchedulingData{}, err
	}
	return &loan_req_schdl, nil
}

func (p *LoanRequestScheduling) FindLoanRequestSchedulingStatusByID(db *gorm.DB, pid uint64, status uint64) (*LoanRequestSchedulingData, error) {
	var err error
	loan_req_schdl := LoanRequestSchedulingData{}
	err = db.Debug().Model(&LoanRequestSchedulingData{}).Limit(100).
		Select(`t_loan_request_scheduling.loan_req_schdl_id,
				t_loan_request.loan_req_id,     
				t_loan_request_scheduling.loan_req_schdl_date_1,
				t_loan_request_scheduling.loan_req_schdl_date_2,
				t_loan_request_scheduling.loan_req_schdl_time,
				t_loan_request_scheduling.loan_req_schdl_place,
				t_loan_request_scheduling.loan_req_schdl_status,
				t_loan_request_scheduling.loan_req_schdl_result,
				t_loan_request_scheduling.loan_req_schdl_created_at at time zone 'Asia/Jakarta' as loan_req_schdl_created_at,
				t_loan_request_scheduling.loan_req_schdl_updated_at at time zone 'Asia/Jakarta' as loan_req_schdl_updated_at,
				t_loan_request_scheduling.loan_req_schdl_deleted_at at time zone 'Asia/Jakarta' as loan_req_schdl_deleted_at`).
		Joins("JOIN t_loan_request on t_loan_request_scheduling.loan_req_id=t_loan_request.loan_req_id").
		Where("loan_req_schdl_id = ? AND loan_req_schdl_status = ?", pid, status).
		Find(&loan_req_schdl).Error
	if err != nil {
		return &LoanRequestSchedulingData{}, err
	}
	return &loan_req_schdl, nil
}

func (p *LoanRequestScheduling) UpdateLoanRequestScheduling(db *gorm.DB, pid uint64) (*LoanRequestScheduling, error) {

	var err error
	db = db.Debug().Model(&LoanRequestScheduling{}).Where("loan_req_schdl_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_req_id":               p.LoanRequestID,
			"loan_req_schdl_date_1":     p.LoanRequestSchedulingDate1,
			"loan_req_schdl_date_2":     p.LoanRequestSchedulingDate2,
			"loan_req_schdl_time":       p.LoanRequestSchedulingTime,
			"loan_req_schdl_place":      p.LoanRequestSchedulingPlace,
			"loan_req_schdl_result":     p.LoanRequestSchedulingResult,
			"loan_req_schdl_status":     p.LoanRequestSchedulingStatus,
			"loan_req_schdl_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanRequestScheduling{}).Where("loan_req_schdl_id = ?", pid).Error
	if err != nil {
		return &LoanRequestScheduling{}, err
	}
	return p, nil
}

func (p *LoanRequestScheduling) UpdateLoanRequestSchedulingStatus(db *gorm.DB, pid uint64) (*LoanRequestScheduling, error) {
	var err error
	db = db.Debug().Model(&LoanRequestScheduling{}).Where("loan_req_schdl_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_req_schdl_status":     p.LoanRequestSchedulingStatus,
			"loan_req_schdl_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanRequestScheduling{}).Where("loan_req_schdl_id = ?", pid).Error
	if err != nil {
		return &LoanRequestScheduling{}, err
	}
	return p, nil
}

func (p *LoanRequestScheduling) UpdateLoanRequestSchedulingStatusUpdate(db *gorm.DB, pid uint64) (*LoanRequestScheduling, error) {
	var err error
	db = db.Debug().Model(&LoanRequestScheduling{}).Where("loan_req_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_req_schdl_status":     p.LoanRequestSchedulingStatus,
			"loan_req_schdl_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanRequestScheduling{}).Where("loan_req_id = ?", pid).Error
	if err != nil {
		return &LoanRequestScheduling{}, err
	}
	return p, nil
}

func (p *LoanRequestScheduling) UpdateLoanRequestSchedulingDate(db *gorm.DB, pid uint64) (*LoanRequestScheduling, error) {
	var err error
	db = db.Debug().Model(&LoanRequestScheduling{}).Where("loan_req_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_req_schdl_date_1":     p.LoanRequestSchedulingDate1,
			"loan_req_schdl_date_2":     p.LoanRequestSchedulingDate2,
			"loan_req_schdl_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanRequestScheduling{}).Where("loan_req_id = ?", pid).Error
	if err != nil {
		return &LoanRequestScheduling{}, err
	}
	return p, nil
}

func (p *LoanRequestScheduling) UpdateLoanRequestSchedulingStatusDelete(db *gorm.DB, pid uint64) (*LoanRequestScheduling, error) {
	var err error
	db = db.Debug().Model(&LoanRequestScheduling{}).Where("loan_req_schdl_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_req_schdl_status":     3,
			"loan_req_schdl_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanRequestScheduling{}).Where("loan_req_schdl_id = ?", pid).Error
	if err != nil {
		return &LoanRequestScheduling{}, err
	}
	return p, nil
}

func (p *LoanRequestScheduling) DeleteLoanRequestScheduling(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&LoanRequestScheduling{}).Where("loan_req_schdl_id = ?", pid).Take(&LoanRequestScheduling{}).Delete(&LoanRequestScheduling{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("LoanRequestScheduling not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// ADDITIONAL MEMORANDUM
// ===================================================================================================================================

type LoanRequestSchedulingMrdmData struct {
	LoanRequestSchedulingID     uint64 `gorm:"column:loan_req_schdl_id" json:"loan_req_schdl_id"`
	MitraUserID                 uint64 `gorm:"column:mitra_user_id;" json:"mitra_user_id"`
	MitraEmployeeID             uint64 `gorm:"column:mitra_emp_id;" json:"mitra_emp_id"`
	MitraUserEmail              string `gorm:"column:mitra_user_email;" json:"mitra_user_email"`
	MitraUserPhone              string `gorm:"column:mitra_user_phone;" json:"mitra_user_phone"`
	MitraEmployeeCode           string `gorm:"column:mitra_emp_code;" json:"mitra_emp_code"`
	MitraEmployeeName           string `gorm:"column:mitra_emp_name;" json:"mitra_emp_name"`
	MitraEmployeeGender         string `gorm:"column:mitra_emp_gender;" json:"mitra_emp_gender"`
	MitraEmployeeBirthPlace     string `gorm:"column:mitra_emp_birth_place;" json:"mitra_emp_birth_place"`
	MitraEmployeeBirthDate      string `gorm:"column:mitra_emp_birth_date;" json:"mitra_emp_birth_date"`
	MitraEmployeeAddress        string `gorm:"column:mitra_emp_address;" json:"mitra_emp_address"`
	MitraEmployeePhoto          string `gorm:"column:mitra_emp_photo;" json:"mitra_emp_photo"`
	LoanRequestID               uint64 `gorm:"column:loan_req_id;" json:"loan_req_id"`
	LoanRequestSchedulingDate1  string `gorm:"column:loan_req_schdl_date_1" json:"loan_req_schdl_date_1"`
	LoanRequestSchedulingDate2  string `gorm:"column:loan_req_schdl_date_2" json:"loan_req_schdl_date_2"`
	LoanRequestSchedulingTime   string `gorm:"column:loan_req_schdl_time" json:"loan_req_schdl_time"`
	LoanRequestSchedulingPlace  string `gorm:"column:loan_req_schdl_place" json:"loan_req_schdl_place"`
	LoanRequestSchedulingResult string `gorm:"column:loan_req_schdl_result" json:"loan_req_schdl_result"`
	LoanRequestSchedulingStatus uint64 `gorm:"column:loan_req_schdl_status" json:"loan_req_schdl_status"`
}

func (LoanRequestSchedulingMrdmData) TableName() string {
	return "t_loan_request_scheduling"
}

func (p *LoanRequestScheduling) FindLoanRequestSchedulingMrdmByID(db *gorm.DB, pid uint64) (LoanRequestSchedulingMrdmData, error) {
	var err error
	loan_req_schdl := LoanRequestSchedulingMrdmData{}
	err = db.Debug().Model(&LoanRequestSchedulingMrdmData{}).Limit(100).
		Select(`t_loan_request_scheduling.loan_req_schdl_id,
				m_mitra_user.mitra_user_id,
				m_mitra_employee.mitra_emp_id,
				m_mitra_user.mitra_user_email,    
				m_mitra_user.mitra_user_phone,    
				m_mitra_employee.mitra_emp_code,   
				m_mitra_employee.mitra_emp_name,   
				m_mitra_employee.mitra_emp_gender,   
				m_mitra_employee.mitra_emp_birth_place,   
				m_mitra_employee.mitra_emp_birth_date,   
				m_mitra_employee.mitra_emp_address,   
				m_mitra_employee.mitra_emp_photo,     
				t_loan_request.loan_req_id,
				t_loan_request_scheduling.loan_req_schdl_date_1,
				t_loan_request_scheduling.loan_req_schdl_date_2,
				t_loan_request_scheduling.loan_req_schdl_time,
				t_loan_request_scheduling.loan_req_schdl_place,
				t_loan_request_scheduling.loan_req_schdl_result,
				t_loan_request_scheduling.loan_req_schdl_status`).
		Joins("JOIN t_loan_request on t_loan_request_scheduling.loan_req_id=t_loan_request.loan_req_id").
		Joins("JOIN m_mitra_user on t_loan_request.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra_employee on m_mitra_user.mitra_emp_id=m_mitra_employee.mitra_emp_id").
		Where("t_loan_request_scheduling.loan_req_id = ? AND t_loan_request_scheduling.loan_req_schdl_status = ?", pid, 1).
		Find(&loan_req_schdl).Error
	if err != nil {
		return LoanRequestSchedulingMrdmData{}, err
	}
	return loan_req_schdl, nil
}
