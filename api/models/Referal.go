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

type Referal struct {
	ReferalID               uint64     `gorm:"column:referal_id;primary_key;" json:"referal_id"`
	ReferalCode             string     `gorm:"column:referal_code" json:"referal_code"`
	ReferalName             string     `gorm:"column:referal_name" json:"referal_name"`
	ReferalDesc             string     `gorm:"column:referal_desc" json:"referal_desc"`
	ReferalProgram          string     `gorm:"column:referal_program" json:"referal_program"`
	ReferalImage            string     `gorm:"column:referal_image" json:"referal_image"`
	ReferalQuota            uint64     `gorm:"column:referal_quota" json:"referal_quota"`
	ReferalDateActiveStart  string     `gorm:"column:referal_date_active_start" json:"referal_date_active_start"`
	ReferalDateActiveFinish string     `gorm:"column:referal_date_active_finish" json:"referal_date_active_finish"`
	ReferalStatus           uint64     `gorm:"column:referal_status;size:2" json:"referal_status"`
	ReferalCreatedBy        string     `gorm:"column:referal_created_by;size:125" json:"referal_created_by"`
	ReferalCreatedAt        time.Time  `gorm:"column:referal_created_at;default:CURRENT_TIMESTAMP" json:"referal_created_at"`
	ReferalUpdatedBy        string     `gorm:"column:referal_updated_by;size:125" json:"referal_updated_by"`
	ReferalUpdatedAt        *time.Time `gorm:"column:referal_updated_at;default:CURRENT_TIMESTAMP" json:"referal_created_at"`
	ReferalDeletedBy        string     `gorm:"column:referal_deleted_by;size:125" json:"referal_deleted_by"`
	ReferalDeletedAt        *time.Time `gorm:"column:referal_deleted_at;default:CURRENT_TIMESTAMP" json:"referal_deleted_at"`
}

type ReferalData struct {
	ReferalID               uint64     `gorm:"column:referal_id" json:"referal_id"`
	ReferalCode             string     `gorm:"column:referal_code" json:"referal_code"`
	ReferalName             string     `gorm:"column:referal_name" json:"referal_name"`
	ReferalDesc             string     `gorm:"column:referal_desc" json:"referal_desc"`
	ReferalProgram          string     `gorm:"column:referal_program" json:"referal_program"`
	ReferalImage            string     `gorm:"column:referal_image" json:"referal_image"`
	ReferalQuota            uint64     `gorm:"column:referal_quota" json:"referal_quota"`
	ReferalDateActiveStart  string     `gorm:"column:referal_date_active_start" json:"referal_date_active_start"`
	ReferalDateActiveFinish string     `gorm:"column:referal_date_active_finish" json:"referal_date_active_finish"`
	ReferalStatus           uint64     `gorm:"column:referal_status" json:"referal_status"`
	ReferalCreatedBy        string     `gorm:"column:referal_created_by" json:"referal_created_by"`
	ReferalCreatedAt        time.Time  `gorm:"column:referal_created_at" json:"referal_created_at"`
	ReferalUpdatedBy        string     `gorm:"column:referal_updated_by" json:"referal_updated_by"`
	ReferalUpdatedAt        *time.Time `gorm:"column:referal_updated_at" json:"referal_updated_at"`
	ReferalDeletedBy        string     `gorm:"column:referal_deleted_by" json:"referal_deleted_by"`
	ReferalDeletedAt        *time.Time `gorm:"column:referal_deleted_at" json:"referal_deleted_at"`
}

type ResponseReferals struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    *[]ReferalData `json:"data"`
}

type ResponseReferal struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    *ReferalData `json:"data"`
}

type ResponseReferalCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (Referal) TableName() string {
	return "m_referal"
}

func (ReferalData) TableName() string {
	return "m_referal"
}

func (p *Referal) Prepare() {
	p.ReferalCode = html.EscapeString(strings.TrimSpace(p.ReferalCode))
	p.ReferalName = html.EscapeString(strings.TrimSpace(p.ReferalName))
	p.ReferalDesc = p.ReferalDesc
	p.ReferalProgram = p.ReferalProgram
	p.ReferalImage = p.ReferalImage
	p.ReferalQuota = p.ReferalQuota
	p.ReferalDateActiveStart = p.ReferalDateActiveStart
	p.ReferalDateActiveFinish = p.ReferalDateActiveFinish
	p.ReferalStatus = p.ReferalStatus
	p.ReferalCreatedBy = html.EscapeString(strings.TrimSpace(p.ReferalCreatedBy))
	p.ReferalCreatedAt = time.Now()
	p.ReferalUpdatedBy = html.EscapeString(strings.TrimSpace(p.ReferalUpdatedBy))
	p.ReferalUpdatedAt = p.ReferalUpdatedAt
	p.ReferalDeletedBy = html.EscapeString(strings.TrimSpace(p.ReferalDeletedBy))
	p.ReferalDeletedAt = p.ReferalDeletedAt
}

func (p *Referal) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.ReferalCode == "" {
			return errors.New("Required Country")
		}
		if p.ReferalName == "" {
			return errors.New("Required Province")
		}
		if p.ReferalDesc == "" {
			return errors.New("Required Regency")
		}
		if p.ReferalProgram == "" {
			return errors.New("Required Regency")
		}
		if p.ReferalImage == "" {
			return errors.New("Required Regency")
		}
		if p.ReferalQuota == 0 {
			return errors.New("Required Regency")
		}
		if p.ReferalDateActiveStart == "" {
			return errors.New("Required Regency")
		}
		if p.ReferalDateActiveFinish == "" {
			return errors.New("Required Regency")
		}
		return nil
	}
}

func (p *Referal) SaveReferal(db *gorm.DB) (*Referal, error) {
	var err error
	err = db.Debug().Model(&Referal{}).Create(&p).Error
	if err != nil {
		return &Referal{}, err
	}
	return p, nil
}

func (p *Referal) FindAllReferal(db *gorm.DB) (*[]ReferalData, error) {
	var err error
	referal := []ReferalData{}
	err = db.Debug().Model(&ReferalData{}).Limit(100).
		Select(`m_referal.referal_id,
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
		Find(&referal).Error
	if err != nil {
		return &[]ReferalData{}, err
	}
	return &referal, nil
}

func (p *Referal) FindAllReferalStatus(db *gorm.DB, status uint64) (*[]ReferalData, error) {
	var err error
	referal := []ReferalData{}
	err = db.Debug().Model(&ReferalData{}).Limit(100).
		Select(`m_referal.referal_id,
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
		Where("referal_status = ?", status).
		Find(&referal).Error
	if err != nil {
		return &[]ReferalData{}, err
	}
	return &referal, nil
}

func (p *Referal) FindReferalDataByID(db *gorm.DB, pid uint64) (*Referal, error) {
	var err error
	err = db.Debug().Model(&Referal{}).
		Select(`m_referal.referal_id,
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
		Where("referal_id = ?", pid).
		Take(&p).Error
	if err != nil {
		return &Referal{}, err
	}
	return p, nil
}

func (p *Referal) FindReferalByID(db *gorm.DB, pid uint64) (*ReferalData, error) {
	var err error
	referal := ReferalData{}
	err = db.Debug().Model(&ReferalData{}).Limit(100).
		Select(`m_referal.referal_id,
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
		Where("referal_id = ?", pid).
		Find(&referal).Error
	if err != nil {
		return &ReferalData{}, err
	}
	return &referal, nil
}

func (p *Referal) FindReferalStatusByID(db *gorm.DB, pid uint64, status uint64) (*ReferalData, error) {
	var err error
	referal := ReferalData{}
	err = db.Debug().Model(&ReferalData{}).Limit(100).
		Select(`m_referal.referal_id,
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
		Where("referal_id = ? AND referal_status = ?", pid, status).
		Find(&referal).Error
	if err != nil {
		return &ReferalData{}, err
	}
	return &referal, nil
}

func (p *Referal) UpdateReferal(db *gorm.DB, pid uint64) (*Referal, error) {
	var err error
	db = db.Debug().Model(&Referal{}).Where("referal_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			// "referal_code":               p.ReferalCode,
			"referal_name":               p.ReferalName,
			"referal_desc":               p.ReferalDesc,
			"referal_program":            p.ReferalProgram,
			"referal_image":              p.ReferalImage,
			"referal_quota":              p.ReferalQuota,
			"referal_date_active_start":  p.ReferalDateActiveStart,
			"referal_date_active_finish": p.ReferalDateActiveFinish,
			"referal_status":             p.ReferalStatus,
			"referal_updated_by":         p.ReferalUpdatedBy,
			"referal_updated_at":         time.Now(),
		},
	)
	err = db.Debug().Model(&Referal{}).Where("referal_id = ?", pid).Error
	if err != nil {
		return &Referal{}, err
	}
	return p, nil
}

func (p *Referal) UpdateReferalStatus(db *gorm.DB, pid uint64) (*Referal, error) {
	var err error
	db = db.Debug().Model(&Referal{}).Where("referal_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"referal_status":     p.ReferalStatus,
			"referal_updated_by": p.ReferalUpdatedBy,
			"referal_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&Referal{}).Where("referal_id = ?", pid).Error
	if err != nil {
		return &Referal{}, err
	}
	return p, nil
}

func (p *Referal) UpdateReferalStatusOnly(db *gorm.DB, pid uint64) (*Referal, error) {
	var err error
	db = db.Debug().Model(&Referal{}).Where("referal_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"referal_status": p.ReferalStatus,
		},
	)
	err = db.Debug().Model(&Referal{}).Where("referal_id = ?", pid).Error
	if err != nil {
		return &Referal{}, err
	}
	return p, nil
}

func (p *Referal) UpdateReferalStatusDelete(db *gorm.DB, pid uint64) (*Referal, error) {
	var err error
	db = db.Debug().Model(&Referal{}).Where("referal_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"referal_status":     3,
			"referal_deleted_by": p.ReferalDeletedBy,
			"referal_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&Referal{}).Where("referal_id = ?", pid).Error
	if err != nil {
		return &Referal{}, err
	}
	return p, nil
}

func (p *Referal) DeleteReferal(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&Referal{}).Where("referal_id = ?", pid).Take(&Referal{}).Delete(&Referal{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Referal not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

func (p *Referal) FindAllReferalCode(db *gorm.DB, referal string) (*Referal, error) {
	var err error
	err = db.Debug().Model(&Referal{}).Where("referal_code = ? AND referal_status = ?", referal, 1).Take(&p).Error
	if err != nil {
		return &Referal{}, err
	}
	return p, nil
}
