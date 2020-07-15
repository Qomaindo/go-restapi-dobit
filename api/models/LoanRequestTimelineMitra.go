package models

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type LoanRequestTimelineMitra struct {
	LoanRequestTimelineMitraID        uint64     `gorm:"column:loan_req_timeline_mitra_id;primary_key;" json:"loan_req_timeline_mitra_id"`
	LoanRequestID                     uint64     `gorm:"column:loan_req_id;" json:"loan_req_id"`
	LoanRequestTimelineMitraTitle     string     `gorm:"column:loan_req_timeline_mitra_title" json:"loan_req_timeline_mitra_title"`
	LoanRequestTimelineMitraDesc      string     `gorm:"column:loan_req_timeline_mitra_desc" json:"loan_req_timeline_mitra_desc"`
	LoanRequestTimelineMitraStatus    uint64     `gorm:"column:loan_req_timeline_mitra_status;size:2" json:"loan_req_timeline_mitra_status"`
	LoanRequestTimelineMitraCreatedAt time.Time  `gorm:"column:loan_req_timeline_mitra_created_at;default:CURRENT_TIMESTAMP" json:"loan_req_timeline_mitra_created_at"`
	LoanRequestTimelineMitraUpdatedAt *time.Time `gorm:"column:loan_req_timeline_mitra_updated_at;default:CURRENT_TIMESTAMP" json:"loan_req_timeline_mitra_updated_at"`
	LoanRequestTimelineMitraDeletedAt *time.Time `gorm:"column:loan_req_timeline_mitra_deleted_at;default:CURRENT_TIMESTAMP" json:"loan_req_timeline_mitra_deleted_at"`
}

type LoanRequestTimelineMitraData struct {
	LoanRequestTimelineMitraID        uint64     `gorm:"column:loan_req_timeline_mitra_id;primary_key;" json:"loan_req_timeline_mitra_id"`
	LoanRequestID                     uint64     `gorm:"column:loan_req_id;" json:"loan_req_id"`
	LoanRequestTimelineMitraTitle     string     `gorm:"column:loan_req_timeline_mitra_title" json:"loan_req_timeline_mitra_title"`
	LoanRequestTimelineMitraDesc      string     `gorm:"column:loan_req_timeline_mitra_desc" json:"loan_req_timeline_mitra_desc"`
	LoanRequestTimelineMitraStatus    uint64     `gorm:"column:loan_req_timeline_mitra_status;size:2" json:"loan_req_timeline_mitra_status"`
	LoanRequestTimelineMitraCreatedAt time.Time  `gorm:"column:loan_req_timeline_mitra_created_at;" json:"loan_req_timeline_mitra_created_at"`
	LoanRequestTimelineMitraUpdatedAt *time.Time `gorm:"column:loan_req_timeline_mitra_updated_at;" json:"loan_req_timeline_mitra_updated_at"`
	LoanRequestTimelineMitraDeletedAt *time.Time `gorm:"column:loan_req_timeline_mitra_deleted_at;default:CURRENT_TIMESTAMP" json:"loan_req_timeline_mitra_deleted_at"`
}

type ResponseLoanRequestTimelineMitras struct {
	Status  int                             `json:"status"`
	Message string                          `json:"message"`
	Data    *[]LoanRequestTimelineMitraData `json:"data"`
	// Data2 *[]LoanRequestTimelineMitraData.LoanRequestTimelineMitraCreatedAt `gorm:"-" json:"loan_req_timeline_mitra_created_ats"`
}

type ResponseLoanRequestTimelineMitra struct {
	Status  int                           `json:"status"`
	Message string                        `json:"message"`
	Data    *LoanRequestTimelineMitraData `json:"data"`
}

type ResponseLoanRequestTimelineMitraIU struct {
	Status  int                       `json:"status"`
	Message string                    `json:"message"`
	Data    *LoanRequestTimelineMitra `json:"data"`
}

type ResponseLoanRequestTimelineMitraDel struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (LoanRequestTimelineMitra) TableName() string {
	return "t_loan_request_timeline_mitra"
}

func (LoanRequestTimelineMitraData) TableName() string {
	return "t_loan_request_timeline_mitra"
}

func (p *LoanRequestTimelineMitra) Prepare() {

	p.LoanRequestID = p.LoanRequestID
	p.LoanRequestTimelineMitraTitle = p.LoanRequestTimelineMitraTitle
	p.LoanRequestTimelineMitraDesc = p.LoanRequestTimelineMitraDesc
	p.LoanRequestTimelineMitraStatus = p.LoanRequestTimelineMitraStatus
	p.LoanRequestTimelineMitraCreatedAt = time.Now()
	p.LoanRequestTimelineMitraUpdatedAt = p.LoanRequestTimelineMitraUpdatedAt
	p.LoanRequestTimelineMitraDeletedAt = p.LoanRequestTimelineMitraDeletedAt
}
func (p *LoanRequestTimelineMitra) Validate(action string) error {
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

func (p *LoanRequestTimelineMitra) SaveLoanRequestTimelineMitra(db *gorm.DB) (*LoanRequestTimelineMitra, error) {
	var err error
	err = db.Debug().Model(&LoanRequestTimelineMitra{}).Create(&p).Error
	if err != nil {
		return &LoanRequestTimelineMitra{}, err
	}
	return p, nil
}

func (p *LoanRequestTimelineMitra) FindAllLoanRequestTimelineMitra(db *gorm.DB) (*[]LoanRequestTimelineMitraData, error) {
	var err error
	loan_req_timeline := []LoanRequestTimelineMitraData{}
	err = db.Debug().Model(&LoanRequestTimelineMitraData{}).Limit(100).
		Select(`t_loan_request_timeline_mitra.loan_req_timeline_mitra_id,
				t_loan_request.loan_req_id,     
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_title,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_desc,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_status,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_created_at at time zone 'Asia/Jakarta' as loan_req_timeline_mitra_created_at,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_created_by,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_updated_by,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_updated_at at time zone 'Asia/Jakarta' as loan_req_timeline_mitra_updated_at,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_deleted_by,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_deleted_at at time zone 'Asia/Jakarta' as loan_req_timeline_mitra_deleted_at`).
		Joins("JOIN t_loan_request on t_loan_request_timeline_mitra.loan_req_id=t_loan_request.loan_req_id").
		Find(&loan_req_timeline).Error
	if err != nil {
		return &[]LoanRequestTimelineMitraData{}, err
	}
	return &loan_req_timeline, nil
}

func (p *LoanRequestTimelineMitra) FindAllLoanRequestTimelineMitraStatus(db *gorm.DB, status uint64) (*[]LoanRequestTimelineMitraData, error) {
	var err error
	loan_req_timeline := []LoanRequestTimelineMitraData{}
	err = db.Debug().Model(&LoanRequestTimelineMitraData{}).Limit(100).
		Select(`t_loan_request_timeline_mitra.loan_req_timeline_mitra_id,
				t_loan_request.loan_req_id,     
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_title,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_desc,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_status,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_created_at at time zone 'Asia/Jakarta' as loan_req_timeline_mitra_created_at,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_created_by,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_updated_by,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_updated_at at time zone 'Asia/Jakarta' as loan_req_timeline_mitra_updated_at,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_deleted_by,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_deleted_at at time zone 'Asia/Jakarta' as loan_req_timeline_mitra_deleted_at`).
		Joins("JOIN t_loan_request on t_loan_request_timeline_mitra.loan_req_id=t_loan_request.loan_req_id").
		Where("loan_req_timeline_mitra_status = ?", status).
		Find(&loan_req_timeline).Error
	if err != nil {
		return &[]LoanRequestTimelineMitraData{}, err
	}
	return &loan_req_timeline, nil
}

func (p *LoanRequestTimelineMitra) FindLoanRequestTimelineMitraDataByID(db *gorm.DB, pid uint64) (*LoanRequestTimelineMitra, error) {
	var err error
	err = db.Debug().Model(&LoanRequestTimelineMitra{}).Where("loan_req_timeline_mitra_id = ?", pid).Take(&p).Error
	if err != nil {
		return &LoanRequestTimelineMitra{}, err
	}
	return p, nil
}

func (p *LoanRequestTimelineMitra) FindLoanRequestTimelineMitraByID(db *gorm.DB, pid uint64) (*LoanRequestTimelineMitraData, error) {
	var err error
	loan_req_timeline := LoanRequestTimelineMitraData{}
	err = db.Debug().Model(&LoanRequestTimelineMitraData{}).Limit(100).
		Select(`t_loan_request_timeline_mitra.loan_req_timeline_mitra_id,
				t_loan_request.loan_req_id,     
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_title,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_desc,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_status,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_created_at at time zone 'Asia/Jakarta' as loan_req_timeline_mitra_created_at,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_created_by,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_updated_by,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_updated_at at time zone 'Asia/Jakarta' as loan_req_timeline_mitra_updated_at,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_deleted_by,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_deleted_at at time zone 'Asia/Jakarta' as loan_req_timeline_mitra_deleted_at`).
		Joins("JOIN t_loan_request on t_loan_request_timeline_mitra.loan_req_id=t_loan_request.loan_req_id").
		Where("loan_req_timeline_mitra_id = ?", pid).
		Find(&loan_req_timeline).Error
	if err != nil {
		return &LoanRequestTimelineMitraData{}, err
	}
	return &loan_req_timeline, nil
}

func (p *LoanRequestTimelineMitra) FindLoanRequestTimelineMitraByReqID(db *gorm.DB, pid string) (*[]LoanRequestTimelineMitraData, error) {
	var err error
	loan_req_timeline := []LoanRequestTimelineMitraData{}
	err = db.Debug().Model(&LoanRequestTimelineMitraData{}).Limit(100).
		Select(`t_loan_request_timeline_mitra.loan_req_timeline_mitra_id,
				t_loan_request.loan_req_id,
				m_mitra_user.mitra_user_id,
				m_mitra_user.mitra_user_phone,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_title,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_desc,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_status,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_created_at at time zone 'Asia/Jakarta' as loan_req_timeline_mitra_created_at,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_created_by,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_updated_by,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_updated_at at time zone 'Asia/Jakarta' as loan_req_timeline_mitra_updated_at,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_deleted_by,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_deleted_at at time zone 'Asia/Jakarta' as loan_req_timeline_mitra_deleted_at`).
		Joins("JOIN t_loan_request on t_loan_request_timeline_mitra.loan_req_id=t_loan_request.loan_req_id").
		Joins("JOIN m_mitra_user on t_loan_request.mitra_user_id=m_mitra_user.mitra_user_id").
		Where("t_loan_request.loan_req_id = ?", pid).
		Order("t_loan_request_timeline_mitra.loan_req_timeline_mitra_id asc").
		Find(&loan_req_timeline).Error
	if err != nil {
		return &[]LoanRequestTimelineMitraData{}, err
	}
	return &loan_req_timeline, nil
}

func (p *LoanRequestTimelineMitra) FindLoanRequestTimelineMitraByPhone(db *gorm.DB, pid string) (*LoanRequestTimelineMitraData, error) {
	var err error
	loan_req_timeline := LoanRequestTimelineMitraData{}
	err = db.Debug().Model(&LoanRequestTimelineMitraData{}).Limit(100).
		Select(`t_loan_request_timeline_mitra.loan_req_timeline_mitra_id,
				t_loan_request.loan_req_id,
				m_mitra_user.mitra_user_id,
				m_mitra_user.mitra_user_phone,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_title,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_desc,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_status,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_created_at at time zone 'Asia/Jakarta' as loan_req_timeline_mitra_created_at,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_created_by,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_updated_by,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_updated_at at time zone 'Asia/Jakarta' as loan_req_timeline_mitra_updated_at,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_deleted_by,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_deleted_at at time zone 'Asia/Jakarta' as loan_req_timeline_mitra_deleted_at`).
		Joins("JOIN t_loan_request on t_loan_request_timeline_mitra.loan_req_id=t_loan_request.loan_req_id").
		Joins("JOIN m_mitra_user on t_loan_request.mitra_user_id=m_mitra_user.mitra_user_id").
		Where("m_mitra_user.mitra_user_phone = ?", pid).
		Find(&loan_req_timeline).Error
	if err != nil {
		return &LoanRequestTimelineMitraData{}, err
	}
	return &loan_req_timeline, nil
}

func (p *LoanRequestTimelineMitra) FindLoanRequestTimelineMitraActiveByLoanRequestID(db *gorm.DB, pid uint64) (*LoanRequestTimelineMitraData, error) {
	var err error
	loan_req_timeline := LoanRequestTimelineMitraData{}
	err = db.Debug().Model(&LoanRequestTimelineMitraData{}).Limit(100).
		Select(`t_loan_request_timeline_mitra.loan_req_timeline_mitra_id,
				t_loan_request.loan_req_id,     
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_title,
				t_loan_request_timeline_mitra.loan_req_timeline_place,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_status,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_created_at at time zone 'Asia/Jakarta' as loan_req_timeline_mitra_created_at,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_created_by,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_updated_by,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_updated_at at time zone 'Asia/Jakarta' as loan_req_timeline_mitra_updated_at,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_deleted_by,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_deleted_at at time zone 'Asia/Jakarta' as loan_req_timeline_mitra_deleted_at`).
		Joins("JOIN t_loan_request on t_loan_request_timeline_mitra.loan_req_id=t_loan_request.loan_req_id").
		Where("t_loan_request_timeline_mitra.loan_req_id = ? AND t_loan_request_timeline_mitra.loan_req_timeline_mitra_status = ?", pid, 1).
		Find(&loan_req_timeline).Error
	if err != nil {
		return &LoanRequestTimelineMitraData{}, err
	}
	return &loan_req_timeline, nil
}

func (p *LoanRequestTimelineMitra) FindLoanRequestTimelineMitraStatusByID(db *gorm.DB, pid uint64, status uint64) (*LoanRequestTimelineMitraData, error) {
	var err error
	loan_req_timeline := LoanRequestTimelineMitraData{}
	err = db.Debug().Model(&LoanRequestTimelineMitraData{}).Limit(100).
		Select(`t_loan_request_timeline_mitra.loan_req_timeline_mitra_id,
				t_loan_request.loan_req_id,     
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_title,
				t_loan_request_timeline_mitra.loan_req_timeline_place,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_status,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_created_at at time zone 'Asia/Jakarta' as loan_req_timeline_mitra_created_at,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_created_by,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_updated_by,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_updated_at at time zone 'Asia/Jakarta' as loan_req_timeline_mitra_updated_at,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_deleted_by,
				t_loan_request_timeline_mitra.loan_req_timeline_mitra_deleted_at at time zone 'Asia/Jakarta' as loan_req_timeline_mitra_deleted_at`).
		Joins("JOIN t_loan_request on t_loan_request_timeline_mitra.loan_req_id=t_loan_request.loan_req_id").
		Where("loan_req_timeline_mitra_id = ? AND loan_req_timeline_mitra_status = ?", pid, status).
		Find(&loan_req_timeline).Error
	if err != nil {
		return &LoanRequestTimelineMitraData{}, err
	}
	return &loan_req_timeline, nil
}

func (p *LoanRequestTimelineMitra) UpdateLoanRequestTimelineMitra(db *gorm.DB, pid uint64) (*LoanRequestTimelineMitra, error) {

	var err error
	db = db.Debug().Model(&LoanRequestTimelineMitra{}).Where("loan_req_timeline_mitra_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_req_id":                        p.LoanRequestID,
			"loan_req_timeline_mitra_title":      p.LoanRequestTimelineMitraTitle,
			"loan_req_timeline_mitra_desc":       p.LoanRequestTimelineMitraDesc,
			"loan_req_timeline_mitra_status":     p.LoanRequestTimelineMitraStatus,
			"loan_req_timeline_mitra_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanRequestTimelineMitra{}).Where("loan_req_timeline_mitra_id = ?", pid).Error
	if err != nil {
		return &LoanRequestTimelineMitra{}, err
	}
	return p, nil
}

func (p *LoanRequestTimelineMitra) UpdateLoanRequestTimelineMitraStatus(db *gorm.DB, pid uint64) (*LoanRequestTimelineMitra, error) {
	var err error
	db = db.Debug().Model(&LoanRequestTimelineMitra{}).Where("loan_req_timeline_mitra_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_req_timeline_mitra_status":     p.LoanRequestTimelineMitraStatus,
			"loan_req_timeline_mitra_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanRequestTimelineMitra{}).Where("loan_req_timeline_mitra_id = ?", pid).Error
	if err != nil {
		return &LoanRequestTimelineMitra{}, err
	}
	return p, nil
}

func (p *LoanRequestTimelineMitra) UpdateLoanRequestTimelineMitraStatusUpdate(db *gorm.DB, pid uint64) (*LoanRequestTimelineMitra, error) {
	var err error
	db = db.Debug().Model(&LoanRequestTimelineMitra{}).Where("loan_req_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_req_timeline_mitra_status":     p.LoanRequestTimelineMitraStatus,
			"loan_req_timeline_mitra_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanRequestTimelineMitra{}).Where("loan_req_id = ?", pid).Error
	if err != nil {
		return &LoanRequestTimelineMitra{}, err
	}
	return p, nil
}

func (p *LoanRequestTimelineMitra) UpdateLoanRequestTimelineMitraStatusDelete(db *gorm.DB, pid uint64) (*LoanRequestTimelineMitra, error) {
	var err error
	db = db.Debug().Model(&LoanRequestTimelineMitra{}).Where("loan_req_timeline_mitra_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_req_timeline_mitra_status":     3,
			"loan_req_timeline_mitra_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanRequestTimelineMitra{}).Where("loan_req_timeline_mitra_id = ?", pid).Error
	if err != nil {
		return &LoanRequestTimelineMitra{}, err
	}
	return p, nil
}

func (p *LoanRequestTimelineMitra) DeleteLoanRequestTimelineMitra(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&LoanRequestTimelineMitra{}).Where("loan_req_timeline_mitra_id = ?", pid).Take(&LoanRequestTimelineMitra{}).Delete(&LoanRequestTimelineMitra{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("LoanRequestTimelineMitra not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
