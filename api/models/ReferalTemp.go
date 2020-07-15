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

type ReferalTemp struct {
	ReferalTempID               uint64    `gorm:"column:referal_temp_id;primary_key;" json:"referal_temp_id"`
	ReferalID                   uint64    `gorm:"column:referal_id" json:"referal_id"`
	ReferalTempName             string    `gorm:"column:referal_temp_name" json:"referal_temp_name"`
	ReferalTempDesc             string    `gorm:"column:referal_temp_desc" json:"referal_temp_desc"`
	ReferalTempProgram          string    `gorm:"column:referal_temp_program" json:"referal_temp_program"`
	ReferalTempImage            string    `gorm:"column:referal_temp_image" json:"referal_temp_image"`
	ReferalTempQuota            uint64    `gorm:"column:referal_temp_quota" json:"referal_temp_quota"`
	ReferalTempDateActiveStart  string    `gorm:"column:referal_temp_date_active_start" json:"referal_temp_date_active_start"`
	ReferalTempDateActiveFinish string    `gorm:"column:referal_temp_date_active_finish" json:"referal_temp_date_active_finish"`
	ReferalTempReason           string    `gorm:"column:referal_temp_reason" json:"referal_temp_reason"`
	ReferalTempStatus           uint64    `gorm:"column:referal_temp_status;size:2" json:"referal_temp_status"`
	ReferalTempActionBefore     uint64    `gorm:"column:referal_temp_action_before;size:2" json:"referal_temp_action_before"`
	ReferalTempCreatedBy        string    `gorm:"column:referal_temp_created_by;size:125" json:"referal_temp_created_by"`
	ReferalTempCreatedAt        time.Time `gorm:"column:referal_temp_created_at;default:CURRENT_TIMESTAMP" json:"referal_temp_created_at"`
}

type ReferalTempPend struct {
	ReferalTempID               uint64    `gorm:"column:referal_temp_id;primary_key;" json:"referal_temp_id"`
	ReferalID                   *uint64   `gorm:"column:referal_id" json:"referal_id"`
	ReferalTempName             string    `gorm:"column:referal_temp_name" json:"referal_temp_name"`
	ReferalTempDesc             string    `gorm:"column:referal_temp_desc" json:"referal_temp_desc"`
	ReferalTempProgram          string    `gorm:"column:referal_temp_program" json:"referal_temp_program"`
	ReferalTempImage            string    `gorm:"column:referal_temp_image" json:"referal_temp_image"`
	ReferalTempQuota            uint64    `gorm:"column:referal_temp_quota" json:"referal_temp_quota"`
	ReferalTempDateActiveStart  string    `gorm:"column:referal_temp_date_active_start" json:"referal_temp_date_active_start"`
	ReferalTempDateActiveFinish string    `gorm:"column:referal_temp_date_active_finish" json:"referal_temp_date_active_finish"`
	ReferalTempReason           string    `gorm:"column:referal_temp_reason" json:"referal_temp_reason"`
	ReferalTempStatus           uint64    `gorm:"column:referal_temp_status;size:2" json:"referal_temp_status"`
	ReferalTempActionBefore     uint64    `gorm:"column:referal_temp_action_before;size:2" json:"referal_temp_action_before"`
	ReferalTempCreatedBy        string    `gorm:"column:referal_temp_created_by;size:125" json:"referal_temp_created_by"`
	ReferalTempCreatedAt        time.Time `gorm:"column:referal_temp_created_at;default:CURRENT_TIMESTAMP" json:"referal_temp_created_at"`
}

type ReferalTempData struct {
	ReferalTempID               uint64    `gorm:"column:referal_temp_id" json:"referal_temp_id"`
	ReferalTempName             string    `gorm:"column:referal_temp_name" json:"referal_temp_name"`
	ReferalTempDesc             string    `gorm:"column:referal_temp_desc" json:"referal_temp_desc"`
	ReferalTempProgram          string    `gorm:"column:referal_temp_program" json:"referal_temp_program"`
	ReferalTempImage            string    `gorm:"column:referal_temp_image" json:"referal_temp_image"`
	ReferalTempQuota            uint64    `gorm:"column:referal_temp_quota" json:"referal_temp_quota"`
	ReferalTempDateActiveStart  string    `gorm:"column:referal_temp_date_active_start" json:"referal_temp_date_active_start"`
	ReferalTempDateActiveFinish string    `gorm:"column:referal_temp_date_active_finish" json:"referal_temp_date_active_finish"`
	ReferalTempReason           string    `gorm:"column:referal_temp_reason" json:"referal_temp_reason"`
	ReferalTempStatus           uint64    `gorm:"column:referal_temp_status;size:2" json:"referal_temp_status"`
	ReferalTempActionBefore     uint64    `gorm:"column:referal_temp_action_before;size:2" json:"referal_temp_action_before"`
	ReferalTempCreatedBy        string    `gorm:"column:referal_temp_created_by;size:125" json:"referal_temp_created_by"`
	ReferalTempCreatedAt        time.Time `gorm:"column:referal_temp_created_at" json:"referal_temp_created_at"`
	ReferalID                   uint64    `gorm:"column:referal_id" json:"referal_id"`
	ReferalName                 string    `gorm:"column:referal_name" json:"referal_name"`
	ReferalDesc                 string    `gorm:"column:referal_desc" json:"referal_desc"`
	ReferalProgram              string    `gorm:"column:referal_program" json:"referal_program"`
	ReferalImage                string    `gorm:"column:referal_image" json:"referal_image"`
	ReferalQuota                uint64    `gorm:"column:referal_quota" json:"referal_quota"`
	ReferalDateActiveStart      string    `gorm:"column:referal_date_active_start" json:"referal_date_active_start"`
	ReferalDateActiveFinish     string    `gorm:"column:referal_date_active_finish" json:"referal_date_active_finish"`
	ReferalStatus               uint64    `gorm:"column:referal_status" json:"referal_status"`
	ReferalCreatedBy            string    `gorm:"column:referal_created_by" json:"referal_created_by"`
	ReferalCreatedAt            time.Time `gorm:"column:referal_created_at" json:"referal_created_at"`
	ReferalUpdatedBy            string    `gorm:"column:referal_updated_by" json:"referal_updated_by"`
	ReferalUpdatedAt            time.Time `gorm:"column:referal_updated_at" json:"referal_updated_at"`
	ReferalDeletedBy            string    `gorm:"column:referal_deleted_by" json:"referal_deleted_by"`
	ReferalDeletedAt            time.Time `gorm:"column:referal_deleted_at" json:"referal_deleted_at"`
}

type ResponseReferalTemps struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    *[]ReferalTempData `json:"data"`
}

type ResponseReferalTemp struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    *ReferalTempData `json:"data"`
}

type ResponseReferalTempsPend struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    *[]ReferalTempPend `json:"data"`
}

type ResponseReferalTempPend struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    *ReferalTempPend `json:"data"`
}

type ResponseReferalTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (ReferalTemp) TableName() string {
	return "m_referal_temp"
}

func (ReferalTempPend) TableName() string {
	return "m_referal_temp"
}

func (ReferalTempData) TableName() string {
	return "m_referal_temp"
}

func (p *ReferalTemp) Prepare() {
	p.ReferalID = p.ReferalID
	p.ReferalTempName = html.EscapeString(strings.TrimSpace(p.ReferalTempName))
	p.ReferalTempDesc = p.ReferalTempDesc
	p.ReferalTempProgram = p.ReferalTempProgram
	p.ReferalTempImage = p.ReferalTempImage
	p.ReferalTempQuota = p.ReferalTempQuota
	p.ReferalTempDateActiveStart = p.ReferalTempDateActiveStart
	p.ReferalTempDateActiveFinish = p.ReferalTempDateActiveFinish
	p.ReferalTempReason = p.ReferalTempReason
	p.ReferalTempStatus = p.ReferalTempStatus
	p.ReferalTempCreatedBy = html.EscapeString(strings.TrimSpace(p.ReferalTempCreatedBy))
	p.ReferalTempCreatedAt = time.Now()
}

func (p *ReferalTempPend) Prepare() {
	p.ReferalID = nil
	p.ReferalTempName = html.EscapeString(strings.TrimSpace(p.ReferalTempName))
	p.ReferalTempDesc = p.ReferalTempDesc
	p.ReferalTempProgram = p.ReferalTempProgram
	p.ReferalTempImage = p.ReferalTempImage
	p.ReferalTempQuota = p.ReferalTempQuota
	p.ReferalTempDateActiveStart = p.ReferalTempDateActiveStart
	p.ReferalTempDateActiveFinish = p.ReferalTempDateActiveFinish
	p.ReferalTempReason = p.ReferalTempReason
	p.ReferalTempStatus = p.ReferalTempStatus
	p.ReferalTempCreatedBy = html.EscapeString(strings.TrimSpace(p.ReferalTempCreatedBy))
	p.ReferalTempCreatedAt = time.Now()
}

func (p *ReferalTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	case "insertupdate":
		if p.ReferalTempName == "" {
			return errors.New("Required Province")
		}
		if p.ReferalTempDesc == "" {
			return errors.New("Required Regency")
		}
		if p.ReferalTempProgram == "" {
			return errors.New("Required Regency")
		}
		if p.ReferalTempImage == "" {
			return errors.New("Required Regency")
		}
		if p.ReferalTempQuota == 0 {
			return errors.New("Required Regency")
		}
		if p.ReferalTempDateActiveStart == "" {
			return errors.New("Required Regency")
		}
		if p.ReferalTempDateActiveFinish == "" {
			return errors.New("Required Regency")
		}
		return nil

	case "reason":
		if p.ReferalTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.ReferalTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	default:
		if p.ReferalTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil
	}
}

func (p *ReferalTemp) SaveReferalTemp(db *gorm.DB) (*ReferalTemp, error) {
	var err error
	err = db.Debug().Model(&ReferalTemp{}).Create(&p).Error
	if err != nil {
		return &ReferalTemp{}, err
	}
	return p, nil
}

func (p *ReferalTempPend) SaveReferalTempPend(db *gorm.DB) (*ReferalTempPend, error) {
	var err error
	err = db.Debug().Model(&ReferalTempPend{}).Create(&p).Error
	if err != nil {
		return &ReferalTempPend{}, err
	}
	return p, nil
}

func (p *ReferalTemp) FindAllReferalTemp(db *gorm.DB) (*[]ReferalTempPend, error) {
	var err error
	referal := []ReferalTempPend{}
	err = db.Debug().Model(&ReferalTempPend{}).Limit(100).
		Select(`m_referal_temp.referal_temp_id,
				m_referal_temp.referal_id,
				m_referal_temp.referal_temp_name,
				m_referal_temp.referal_temp_desc,
				m_referal_temp.referal_temp_program,
				m_referal_temp.referal_temp_image,
				m_referal_temp.referal_temp_quota,
				m_referal_temp.referal_temp_date_active_start,
				m_referal_temp.referal_temp_date_active_finish,
				m_referal_temp.referal_temp_reason,
				m_referal_temp.referal_temp_status,
				m_referal_temp.referal_temp_action_before,
				m_referal_temp.referal_temp_created_by,
				m_referal_temp.referal_temp_created_at at time zone 'Asia/Jakarta' as referal_temp_created_at`).
		Order("address_temp_created_at desc").
		Find(&referal).Error
	if err != nil {
		return &[]ReferalTempPend{}, err
	}
	return &referal, nil
}

func (p *ReferalTemp) FindAllReferalTempPendingActive(db *gorm.DB) (*[]ReferalTempPend, error) {
	var err error
	referal := []ReferalTempPend{}
	err = db.Debug().Model(&ReferalTempPend{}).Limit(100).
		Select(`m_referal_temp.referal_temp_id,
				m_referal_temp.referal_id,
				m_referal_temp.referal_temp_name,
				m_referal_temp.referal_temp_desc,
				m_referal_temp.referal_temp_program,
				m_referal_temp.referal_temp_image,
				m_referal_temp.referal_temp_quota,
				m_referal_temp.referal_temp_date_active_start,
				m_referal_temp.referal_temp_date_active_finish,
				m_referal_temp.referal_temp_reason,
				m_referal_temp.referal_temp_status,
				m_referal_temp.referal_temp_action_before,
				m_referal_temp.referal_temp_created_by,
				m_referal_temp.referal_temp_created_at at time zone 'Asia/Jakarta' as referal_temp_created_at`).
		Where("referal_temp_status = ?", 11).
		Order("address_temp_created_at desc").
		Find(&referal).Error
	if err != nil {
		return &[]ReferalTempPend{}, err
	}
	return &referal, nil
}

func (p *ReferalTemp) FindAllReferalTempStatus(db *gorm.DB, status uint64) (*[]ReferalTempData, error) {
	var err error
	referal := []ReferalTempData{}
	err = db.Debug().Model(&ReferalTempData{}).Limit(100).
		Select(`m_referal_temp.referal_temp_id,
				m_referal_temp.referal_temp_name,
				m_referal_temp.referal_temp_desc,
				m_referal_temp.referal_temp_program,
				m_referal_temp.referal_temp_image,
				m_referal_temp.referal_temp_quota,
				m_referal_temp.referal_temp_date_active_start,
				m_referal_temp.referal_temp_date_active_finish,
				m_referal_temp.referal_temp_reason,
				m_referal_temp.referal_temp_status,
				m_referal_temp.referal_temp_action_before,
				m_referal_temp.referal_temp_created_by,
				m_referal_temp.referal_temp_created_at at time zone 'Asia/Jakarta' as referal_temp_created_at,
				m_referal.referal_id,
				m_referal.referal_code,
				m_referal.referal_name,
				m_referal.referal_desc,
				m_referal.referal_program,
				m_referal.referal_image,
				m_referal.referal_quota,
				m_referal.referal_date_active_start,
				m_referal.referal_date_active_finish,
				m_referal.referal_status,
				m_referal.referal_created_by,
				m_referal.referal_created_at at time zone 'Asia/Jakarta' as referal_created_at,
				m_referal.referal_updated_by,
				m_referal.referal_updated_at at time zone 'Asia/Jakarta' as referal_updated_at,
				m_referal.referal_deleted_by,
				m_referal.referal_deleted_at at time zone 'Asia/Jakarta' as referal_deleted_at`).
		Joins("JOIN m_referal on m_referal_temp.referal_id=m_referal.referal_id").
		Where("referal_temp_status = ?", status).
		Order("address_temp_created_at desc").
		Find(&referal).Error
	if err != nil {
		return &[]ReferalTempData{}, err
	}
	return &referal, nil
}

func (p *ReferalTemp) FindReferalTempDataByID(db *gorm.DB, pid uint64) (*ReferalTemp, error) {
	var err error
	err = db.Debug().Model(&ReferalTemp{}).Where("referal_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &ReferalTemp{}, err
	}
	return p, nil
}

func (p *ReferalTemp) FindReferalTempByIDPendingActive(db *gorm.DB, pid uint64) (*ReferalTempPend, error) {
	var err error
	referal := ReferalTempPend{}
	err = db.Debug().Model(&ReferalTempPend{}).Limit(100).
		Select(`m_referal_temp.referal_temp_id,
				m_referal_temp.referal_id,
				m_referal_temp.referal_temp_name,
				m_referal_temp.referal_temp_desc,
				m_referal_temp.referal_temp_program,
				m_referal_temp.referal_temp_image,
				m_referal_temp.referal_temp_quota,
				m_referal_temp.referal_temp_date_active_start,
				m_referal_temp.referal_temp_date_active_finish,
				m_referal_temp.referal_temp_reason,
				m_referal_temp.referal_temp_status,
				m_referal_temp.referal_temp_created_by,
				m_referal_temp.referal_temp_created_at at time zone 'Asia/Jakarta' as referal_temp_created_at`).
		Where("referal_temp_id = ?", pid).
		Order("address_temp_created_at desc").
		Find(&referal).Error
	if err != nil {
		return &ReferalTempPend{}, err
	}
	return &referal, nil
}

func (p *ReferalTemp) FindReferalTempStatusByIDPendingActive(db *gorm.DB, pid uint64) (*ReferalTempPend, error) {
	var err error
	referal := ReferalTempPend{}
	err = db.Debug().Model(&ReferalTempPend{}).Limit(100).
		Select(`m_referal_temp.referal_temp_id,
				m_referal_temp.referal_id,
				m_referal_temp.referal_temp_name,
				m_referal_temp.referal_temp_desc,
				m_referal_temp.referal_temp_program,
				m_referal_temp.referal_temp_image,
				m_referal_temp.referal_temp_quota,
				m_referal_temp.referal_temp_date_active_start,
				m_referal_temp.referal_temp_date_active_finish,
				m_referal_temp.referal_temp_reason,
				m_referal_temp.referal_temp_status,
				m_referal_temp.referal_temp_action_before,
				m_referal_temp.referal_temp_created_by,
				m_referal_temp.referal_temp_created_at at time zone 'Asia/Jakarta' as referal_temp_created_at`).
		Where("referal_temp_id = ? AND referal_temp_status = ?", pid, 11).
		Order("address_temp_created_at desc").
		Find(&referal).Error
	if err != nil {
		return &ReferalTempPend{}, err
	}
	return &referal, nil
}

func (p *ReferalTemp) FindReferalTempByID(db *gorm.DB, pid uint64) (*ReferalTempData, error) {
	var err error
	referal := ReferalTempData{}
	err = db.Debug().Model(&ReferalTempData{}).Limit(100).
		Select(`m_referal_temp.referal_temp_id,
				m_referal_temp.referal_temp_name,
				m_referal_temp.referal_temp_desc,
				m_referal_temp.referal_temp_program,
				m_referal_temp.referal_temp_image,
				m_referal_temp.referal_temp_quota,
				m_referal_temp.referal_temp_date_active_start,
				m_referal_temp.referal_temp_date_active_finish,
				m_referal_temp.referal_temp_reason,
				m_referal_temp.referal_temp_status,
				m_referal_temp.referal_temp_action_before,
				m_referal_temp.referal_temp_created_by,
				m_referal_temp.referal_temp_created_at at time zone 'Asia/Jakarta' as referal_temp_created_at,
				m_referal.referal_id,
				m_referal.referal_code,
				m_referal.referal_name,
				m_referal.referal_desc,
				m_referal.referal_program,
				m_referal.referal_image,
				m_referal.referal_quota,
				m_referal.referal_date_active_start,
				m_referal.referal_date_active_finish,
				m_referal.referal_status,
				m_referal.referal_created_by,
				m_referal.referal_created_at at time zone 'Asia/Jakarta' as referal_created_at,
				m_referal.referal_updated_by,
				m_referal.referal_updated_at at time zone 'Asia/Jakarta' as referal_updated_at,
				m_referal.referal_deleted_by,
				m_referal.referal_deleted_at at time zone 'Asia/Jakarta' as referal_deleted_at`).
		Joins("JOIN m_referal on m_referal_temp.referal_id=m_referal.referal_id").
		Where("referal_temp_id = ?", pid).
		Order("address_temp_created_at desc").
		Find(&referal).Error
	if err != nil {
		return &ReferalTempData{}, err
	}
	return &referal, nil
}

func (p *ReferalTemp) FindReferalTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*ReferalTempData, error) {
	var err error
	referal := ReferalTempData{}
	err = db.Debug().Model(&ReferalTempData{}).Limit(100).
		Select(`m_referal_temp.referal_temp_id,
				m_referal_temp.referal_temp_name,
				m_referal_temp.referal_temp_desc,
				m_referal_temp.referal_temp_program,
				m_referal_temp.referal_temp_image,
				m_referal_temp.referal_temp_quota,
				m_referal_temp.referal_temp_date_active_start,
				m_referal_temp.referal_temp_date_active_finish,
				m_referal_temp.referal_temp_reason,
				m_referal_temp.referal_temp_status,
				m_referal_temp.referal_temp_action_before,
				m_referal_temp.referal_temp_created_by,
				m_referal_temp.referal_temp_created_at at time zone 'Asia/Jakarta' as referal_temp_created_at,
				m_referal.referal_id,
				m_referal.referal_code,
				m_referal.referal_name,
				m_referal.referal_desc,
				m_referal.referal_program,
				m_referal.referal_image,
				m_referal.referal_quota,
				m_referal.referal_date_active_start,
				m_referal.referal_date_active_finish,
				m_referal.referal_status,
				m_referal.referal_created_by,
				m_referal.referal_created_at at time zone 'Asia/Jakarta' as referal_created_at,
				m_referal.referal_updated_by,
				m_referal.referal_updated_at at time zone 'Asia/Jakarta' as referal_updated_at,
				m_referal.referal_deleted_by,
				m_referal.referal_deleted_at at time zone 'Asia/Jakarta' as referal_deleted_at`).
		Joins("JOIN m_referal on m_referal_temp.referal_id=m_referal.referal_id").
		Where("referal_temp_id = ? AND referal_temp_status = ?", pid, status).
		Order("address_temp_created_at desc").
		Find(&referal).Error
	if err != nil {
		return &ReferalTempData{}, err
	}
	return &referal, nil
}

func (p *ReferalTemp) UpdateReferalTemp(db *gorm.DB, pid uint64) (*ReferalTemp, error) {
	var err error
	db = db.Debug().Model(&ReferalTemp{}).Where("referal_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"referal_temp_name":               p.ReferalTempName,
			"referal_temp_desc":               p.ReferalTempDesc,
			"referal_temp_program":            p.ReferalTempProgram,
			"referal_temp_image":              p.ReferalTempImage,
			"referal_temp_quota":              p.ReferalTempQuota,
			"referal_temp_date_active_start":  p.ReferalTempDateActiveStart,
			"referal_temp_date_active_finish": p.ReferalTempDateActiveFinish,
			"referal_temp_reason":             p.ReferalTempReason,
		},
	)
	err = db.Debug().Model(&ReferalTemp{}).Where("referal_temp_id = ?", pid).Error
	if err != nil {
		return &ReferalTemp{}, err
	}
	return p, nil
}

func (p *ReferalTemp) UpdateReferalTempStatus(db *gorm.DB, pid uint64) (*ReferalTemp, error) {
	var err error
	db = db.Debug().Model(&ReferalTemp{}).Where("referal_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"referal_temp_reason": p.ReferalTempReason,
			"referal_temp_status": p.ReferalTempStatus,
		},
	)
	err = db.Debug().Model(&ReferalTemp{}).Where("referal_temp_id = ?", pid).Error
	if err != nil {
		return &ReferalTemp{}, err
	}
	return p, nil
}

func (p *ReferalTemp) DeleteReferalTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&ReferalTemp{}).Where("referal_temp_id = ?", pid).Take(&ReferalTemp{}).Delete(&ReferalTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("ReferalTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================
