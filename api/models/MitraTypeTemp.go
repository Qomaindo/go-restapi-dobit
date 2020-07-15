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

type MitraTypeTemp struct {
	MitraTypeTempID           uint64    `gorm:"column:mitra_type_temp_id;primary_key;" json:"mitra_type_temp_id"`
	MitraTypeID               uint64    `gorm:"column:mitra_type_id" json:"mitra_type_id"`
	MitraTypeTempCode         string    `gorm:"column:mitra_type_temp_code" json:"mitra_type_temp_code"`
	MitraTypeTempName         string    `gorm:"column:mitra_type_temp_name" json:"mitra_type_temp_name"`
	MitraTypeTempDesc         string    `gorm:"column:mitra_type_temp_desc" json:"mitra_type_temp_desc"`
	MitraTypeTempReason       string    `gorm:"column:mitra_type_temp_reason" json:"mitra_type_temp_reason"`
	MitraTypeTempStatus       uint64    `gorm:"column:mitra_type_temp_status;size:2" json:"mitra_type_temp_status"`
	MitraTypeTempActionBefore uint64    `gorm:"column:mitra_type_temp_action_before;size:2" json:"mitra_type_temp_action_before"`
	MitraTypeTempCreatedBy    string    `gorm:"column:mitra_type_temp_created_by;size:125" json:"mitra_type_temp_created_by"`
	MitraTypeTempCreatedAt    time.Time `gorm:"column:mitra_type_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_type_temp_created_at"`
}

type MitraTypeTempPend struct {
	MitraTypeTempID           uint64    `gorm:"column:mitra_type_temp_id;primary_key;" json:"mitra_type_temp_id"`
	MitraTypeID               *uint64   `gorm:"column:mitra_type_id" json:"mitra_type_id"`
	MitraTypeTempCode         string    `gorm:"column:mitra_type_temp_code" json:"mitra_type_temp_code"`
	MitraTypeTempName         string    `gorm:"column:mitra_type_temp_name" json:"mitra_type_temp_name"`
	MitraTypeTempDesc         string    `gorm:"column:mitra_type_temp_desc" json:"mitra_type_temp_desc"`
	MitraTypeTempReason       string    `gorm:"column:mitra_type_temp_reason" json:"mitra_type_temp_reason"`
	MitraTypeTempStatus       uint64    `gorm:"column:mitra_type_temp_status;size:2" json:"mitra_type_temp_status"`
	MitraTypeTempActionBefore uint64    `gorm:"column:mitra_type_temp_action_before;size:2" json:"mitra_type_temp_action_before"`
	MitraTypeTempCreatedBy    string    `gorm:"column:mitra_type_temp_created_by;size:125" json:"mitra_type_temp_created_by"`
	MitraTypeTempCreatedAt    time.Time `gorm:"column:mitra_type_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_type_temp_created_at"`
}

type MitraTypeTempData struct {
	MitraTypeTempID           uint64    `gorm:"column:mitra_type_temp_id" json:"mitra_type_temp_id"`
	MitraTypeTempCode         string    `gorm:"column:mitra_type_temp_code" json:"mitra_type_temp_code"`
	MitraTypeTempName         string    `gorm:"column:mitra_type_temp_name" json:"mitra_type_temp_name"`
	MitraTypeTempDesc         string    `gorm:"column:mitra_type_temp_desc" json:"mitra_type_temp_desc"`
	MitraTypeTempReason       string    `gorm:"column:mitra_type_temp_reason" json:"mitra_type_temp_reason"`
	MitraTypeTempStatus       uint64    `gorm:"column:mitra_type_temp_status;size:2" json:"mitra_type_temp_status"`
	MitraTypeTempActionBefore uint64    `gorm:"column:mitra_type_temp_action_before;size:2" json:"mitra_type_temp_action_before"`
	MitraTypeTempCreatedBy    string    `gorm:"column:mitra_type_temp_created_by;size:125" json:"mitra_type_temp_created_by"`
	MitraTypeTempCreatedAt    time.Time `gorm:"column:mitra_type_temp_created_at" json:"mitra_type_temp_created_at"`
	MitraTypeID               uint64    `gorm:"column:mitra_type_id" json:"mitra_type_id"`
	MitraTypeCode             string    `gorm:"column:mitra_type_code" json:"mitra_type_code"`
	MitraTypeName             string    `gorm:"column:mitra_type_name" json:"mitra_type_name"`
	MitraTypeDesc             string    `gorm:"column:mitra_type_desc" json:"mitra_type_desc"`
	MitraTypeStatus           uint64    `gorm:"column:mitra_type_status" json:"mitra_type_status"`
	MitraTypeCreatedBy        string    `gorm:"column:mitra_type_created_by" json:"mitra_type_created_by"`
	MitraTypeCreatedAt        time.Time `gorm:"column:mitra_type_created_at" json:"mitra_type_created_at"`
	MitraTypeUpdatedBy        string    `gorm:"column:mitra_type_updated_by" json:"mitra_type_updated_by"`
	MitraTypeUpdatedAt        time.Time `gorm:"column:mitra_type_updated_at" json:"mitra_type_updated_at"`
	MitraTypeDeletedBy        string    `gorm:"column:mitra_type_deleted_by" json:"mitra_type_deleted_by"`
	MitraTypeDeletedAt        time.Time `gorm:"column:mitra_type_deleted_at" json:"mitra_type_deleted_at"`
}

type ResponseMitraTypeTemps struct {
	Status  int                  `json:"status"`
	Message string               `json:"message"`
	Data    *[]MitraTypeTempData `json:"data"`
}

type ResponseMitraTypeTemp struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    *MitraTypeTempData `json:"data"`
}

type ResponseMitraTypeTempsPend struct {
	Status  int                  `json:"status"`
	Message string               `json:"message"`
	Data    *[]MitraTypeTempPend `json:"data"`
}

type ResponseMitraTypeTempPend struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    *MitraTypeTempPend `json:"data"`
}

type ResponseMitraTypeTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraTypeTemp) TableName() string {
	return "m_mitra_type_temp"
}

func (MitraTypeTempPend) TableName() string {
	return "m_mitra_type_temp"
}

func (MitraTypeTempData) TableName() string {
	return "m_mitra_type_temp"
}

func (p *MitraTypeTemp) Prepare() {
	p.MitraTypeID = p.MitraTypeID
	p.MitraTypeTempCode = html.EscapeString(strings.TrimSpace(p.MitraTypeTempCode))
	p.MitraTypeTempName = html.EscapeString(strings.TrimSpace(p.MitraTypeTempName))
	p.MitraTypeTempDesc = p.MitraTypeTempDesc
	p.MitraTypeTempReason = p.MitraTypeTempReason
	p.MitraTypeTempStatus = p.MitraTypeTempStatus
	p.MitraTypeTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraTypeTempCreatedBy))
	p.MitraTypeTempCreatedAt = time.Now()
}

func (p *MitraTypeTempPend) Prepare() {
	p.MitraTypeID = nil
	p.MitraTypeTempCode = html.EscapeString(strings.TrimSpace(p.MitraTypeTempCode))
	p.MitraTypeTempName = html.EscapeString(strings.TrimSpace(p.MitraTypeTempName))
	p.MitraTypeTempDesc = p.MitraTypeTempDesc
	p.MitraTypeTempReason = p.MitraTypeTempReason
	p.MitraTypeTempStatus = p.MitraTypeTempStatus
	p.MitraTypeTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraTypeTempCreatedBy))
	p.MitraTypeTempCreatedAt = time.Now()
}

func (p *MitraTypeTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	case "insertupdate":
		if p.MitraTypeTempCode == "" {
			return errors.New("Negara Wajib diisi")
		}
		if p.MitraTypeTempName == "" {
			return errors.New("Provinsi Wajib diisi")
		}
		if p.MitraTypeTempDesc == "" {
			return errors.New("Kabupaten / Kota Wajib diisi")
		}
		if p.MitraTypeTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.MitraTypeTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	case "reason":
		if p.MitraTypeTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.MitraTypeTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	default:
		if p.MitraTypeTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil
	}
}

func (p *MitraTypeTemp) SaveMitraTypeTemp(db *gorm.DB) (*MitraTypeTemp, error) {
	var err error
	err = db.Debug().Model(&MitraTypeTemp{}).Create(&p).Error
	if err != nil {
		return &MitraTypeTemp{}, err
	}
	return p, nil
}

func (p *MitraTypeTempPend) SaveMitraTypeTempPend(db *gorm.DB) (*MitraTypeTempPend, error) {
	var err error
	err = db.Debug().Model(&MitraTypeTempPend{}).Create(&p).Error
	if err != nil {
		return &MitraTypeTempPend{}, err
	}
	return p, nil
}

func (p *MitraTypeTemp) FindAllMitraTypeTemp(db *gorm.DB) (*[]MitraTypeTempPend, error) {
	var err error
	mitratype := []MitraTypeTempPend{}
	err = db.Debug().Model(&MitraTypeTempPend{}).Limit(100).
		Select(`m_mitra_type_temp.mitra_type_temp_id,
				m_mitra_type_temp.mitra_type_id,
				m_mitra_type_temp.mitra_type_temp_code,
				m_mitra_type_temp.mitra_type_temp_name,
				m_mitra_type_temp.mitra_type_temp_desc,
				m_mitra_type_temp.mitra_type_temp_reason,
				m_mitra_type_temp.mitra_type_temp_status,
				m_mitra_type_temp.mitra_type_temp_action_before,
				m_mitra_type_temp.mitra_type_temp_created_by,
				m_mitra_type_temp.mitra_type_temp_created_at at time zone 'Asia/Jakarta' as mitra_type_temp_created_at`).
		Order("mitra_type_temp_created_at desc").
		Find(&mitratype).Error
	if err != nil {
		return &[]MitraTypeTempPend{}, err
	}
	return &mitratype, nil
}

func (p *MitraTypeTemp) FindAllMitraTypeTempPendingActive(db *gorm.DB) (*[]MitraTypeTempPend, error) {
	var err error
	mitratype := []MitraTypeTempPend{}
	err = db.Debug().Model(&MitraTypeTempPend{}).Limit(100).
		Select(`m_mitra_type_temp.mitra_type_temp_id,
				m_mitra_type_temp.mitra_type_id,
				m_mitra_type_temp.mitra_type_temp_code,
				m_mitra_type_temp.mitra_type_temp_name,
				m_mitra_type_temp.mitra_type_temp_desc,
				m_mitra_type_temp.mitra_type_temp_reason,
				m_mitra_type_temp.mitra_type_temp_status,
				m_mitra_type_temp.mitra_type_temp_action_before,
				m_mitra_type_temp.mitra_type_temp_created_by,
				m_mitra_type_temp.mitra_type_temp_created_at at time zone 'Asia/Jakarta' as mitra_type_temp_created_at`).
		Where("mitra_type_temp_status = ?", 11).
		Order("mitra_type_temp_created_at desc").
		Find(&mitratype).Error
	if err != nil {
		return &[]MitraTypeTempPend{}, err
	}
	return &mitratype, nil
}

func (p *MitraTypeTemp) FindAllMitraTypeTempStatus(db *gorm.DB, status uint64) (*[]MitraTypeTempData, error) {
	var err error
	mitratype := []MitraTypeTempData{}
	err = db.Debug().Model(&MitraTypeTempData{}).Limit(100).
		Select(`m_mitra_type_temp.mitra_type_temp_id,
				m_mitra_type_temp.mitra_type_temp_code,
				m_mitra_type_temp.mitra_type_temp_name,
				m_mitra_type_temp.mitra_type_temp_desc,
				m_mitra_type_temp.mitra_type_temp_reason,
				m_mitra_type_temp.mitra_type_temp_status,
				m_mitra_type_temp.mitra_type_temp_action_before,
				m_mitra_type_temp.mitra_type_temp_created_by,
				m_mitra_type_temp.mitra_type_temp_created_at at time zone 'Asia/Jakarta' as mitra_type_temp_created_at,
				m_mitra_type.mitra_type_id,
				m_mitra_type.mitra_type_code,
				m_mitra_type.mitra_type_name,
				m_mitra_type.mitra_type_desc,
				m_mitra_type.mitra_type_status,
				m_mitra_type.mitra_type_created_by,
				m_mitra_type.mitra_type_created_at at time zone 'Asia/Jakarta' as mitra_type_created_at,
				m_mitra_type.mitra_type_updated_by,
				m_mitra_type.mitra_type_updated_at at time zone 'Asia/Jakarta' as mitra_type_updated_at,
				m_mitra_type.mitra_type_deleted_by,
				m_mitra_type.mitra_type_deleted_at at time zone 'Asia/Jakarta' as mitra_type_deleted_at`).
		Joins("JOIN m_mitra_type on m_mitra_type_temp.mitra_type_id=m_mitra_type.mitra_type_id").
		Where("mitra_type_temp_status = ?", status).
		Order("mitra_type_temp_created_at desc").
		Find(&mitratype).Error
	if err != nil {
		return &[]MitraTypeTempData{}, err
	}
	return &mitratype, nil
}

func (p *MitraTypeTemp) FindMitraTypeTempDataByID(db *gorm.DB, pid uint64) (*MitraTypeTemp, error) {
	var err error
	err = db.Debug().Model(&MitraTypeTemp{}).Where("mitra_type_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraTypeTemp{}, err
	}
	return p, nil
}

func (p *MitraTypeTemp) FindMitraTypeTempByIDPendingActive(db *gorm.DB, pid uint64) (*MitraTypeTempPend, error) {
	var err error
	mitratype := MitraTypeTempPend{}
	err = db.Debug().Model(&MitraTypeTempPend{}).Limit(100).
		Select(`m_mitra_type_temp.mitra_type_temp_id,
				m_mitra_type_temp.mitra_type_id,
				m_mitra_type_temp.mitra_type_temp_code,
				m_mitra_type_temp.mitra_type_temp_name,
				m_mitra_type_temp.mitra_type_temp_desc,
				m_mitra_type_temp.mitra_type_temp_reason,
				m_mitra_type_temp.mitra_type_temp_status,
				m_mitra_type_temp.mitra_type_temp_action_before,
				m_mitra_type_temp.mitra_type_temp_created_by,
				m_mitra_type_temp.mitra_type_temp_created_at at time zone 'Asia/Jakarta' as mitra_type_temp_created_at`).
		Where("mitra_type_temp_id = ? AND mitra_type_temp_status = ?", pid, 11).
		Order("mitra_type_temp_created_at desc").
		Find(&mitratype).Error
	if err != nil {
		return &MitraTypeTempPend{}, err
	}
	return &mitratype, nil
}

func (p *MitraTypeTemp) FindMitraTypeTempByID(db *gorm.DB, pid uint64) (*MitraTypeTempData, error) {
	var err error
	mitratype := MitraTypeTempData{}
	err = db.Debug().Model(&MitraTypeTempData{}).Limit(100).
		Select(`m_mitra_type_temp.mitra_type_temp_id,
				m_mitra_type_temp.mitra_type_temp_code,
				m_mitra_type_temp.mitra_type_temp_name,
				m_mitra_type_temp.mitra_type_temp_desc,
				m_mitra_type_temp.mitra_type_temp_reason,
				m_mitra_type_temp.mitra_type_temp_status,
				m_mitra_type_temp.mitra_type_temp_action_before,
				m_mitra_type_temp.mitra_type_temp_created_by,
				m_mitra_type_temp.mitra_type_temp_created_at at time zone 'Asia/Jakarta' as mitra_type_temp_created_at,
				m_mitra_type.mitra_type_id,
				m_mitra_type.mitra_type_code,
				m_mitra_type.mitra_type_name,
				m_mitra_type.mitra_type_desc,
				m_mitra_type.mitra_type_status,
				m_mitra_type.mitra_type_created_by,
				m_mitra_type.mitra_type_created_at at time zone 'Asia/Jakarta' as mitra_type_created_at,
				m_mitra_type.mitra_type_updated_by,
				m_mitra_type.mitra_type_updated_at at time zone 'Asia/Jakarta' as mitra_type_updated_at,
				m_mitra_type.mitra_type_deleted_by,
				m_mitra_type.mitra_type_deleted_at at time zone 'Asia/Jakarta' as mitra_type_deleted_at`).
		Joins("JOIN m_mitra_type on m_mitra_type_temp.mitra_type_id=m_mitra_type.mitra_type_id").
		Where("mitra_type_temp_id = ?", pid).
		Order("mitra_type_temp_created_at desc").
		Find(&mitratype).Error
	if err != nil {
		return &MitraTypeTempData{}, err
	}
	return &mitratype, nil
}

func (p *MitraTypeTemp) FindMitraTypeTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraTypeTempData, error) {
	var err error
	mitratype := MitraTypeTempData{}
	err = db.Debug().Model(&MitraTypeTempData{}).Limit(100).
		Select(`m_mitra_type_temp.mitra_type_temp_id,
				m_mitra_type_temp.mitra_type_temp_code,
				m_mitra_type_temp.mitra_type_temp_name,
				m_mitra_type_temp.mitra_type_temp_desc,
				m_mitra_type_temp.mitra_type_temp_reason,
				m_mitra_type_temp.mitra_type_temp_status,
				m_mitra_type_temp.mitra_type_temp_action_before,
				m_mitra_type_temp.mitra_type_temp_created_by,
				m_mitra_type_temp.mitra_type_temp_created_at at time zone 'Asia/Jakarta' as mitra_type_temp_created_at,
				m_mitra_type.mitra_type_id,
				m_mitra_type.mitra_type_code,
				m_mitra_type.mitra_type_name,
				m_mitra_type.mitra_type_desc,
				m_mitra_type.mitra_type_status,
				m_mitra_type.mitra_type_created_by,
				m_mitra_type.mitra_type_created_at at time zone 'Asia/Jakarta' as mitra_type_created_at,
				m_mitra_type.mitra_type_updated_by,
				m_mitra_type.mitra_type_updated_at at time zone 'Asia/Jakarta' as mitra_type_updated_at,
				m_mitra_type.mitra_type_deleted_by,
				m_mitra_type.mitra_type_deleted_at at time zone 'Asia/Jakarta' as mitra_type_deleted_at`).
		Joins("JOIN m_mitra_type on m_mitra_type_temp.mitra_type_id=m_mitra_type.mitra_type_id").
		Order("mitra_type_temp_created_at desc").
		Find(&mitratype).Error
	if err != nil {
		return &MitraTypeTempData{}, err
	}
	return &mitratype, nil
}

func (p *MitraTypeTemp) UpdateMitraTypeTemp(db *gorm.DB, pid uint64) (*MitraTypeTemp, error) {
	var err error
	db = db.Debug().Model(&MitraTypeTemp{}).Where("mitra_type_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_type_temp_code":   p.MitraTypeTempCode,
			"mitra_type_temp_name":   p.MitraTypeTempName,
			"mitra_type_temp_desc":   p.MitraTypeTempDesc,
			"mitra_type_temp_reason": p.MitraTypeTempReason,
		},
	)
	err = db.Debug().Model(&MitraTypeTemp{}).Where("mitra_type_temp_id = ?", pid).Error
	if err != nil {
		return &MitraTypeTemp{}, err
	}
	return p, nil
}

func (p *MitraTypeTemp) UpdateMitraTypeTempStatus(db *gorm.DB, pid uint64) (*MitraTypeTemp, error) {
	var err error
	db = db.Debug().Model(&MitraTypeTemp{}).Where("mitra_type_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_type_temp_reason": p.MitraTypeTempReason,
			"mitra_type_temp_status": p.MitraTypeTempStatus,
		},
	)
	err = db.Debug().Model(&MitraTypeTemp{}).Where("mitra_type_temp_id = ?", pid).Error
	if err != nil {
		return &MitraTypeTemp{}, err
	}
	return p, nil
}

func (p *MitraTypeTemp) DeleteMitraTypeTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraTypeTemp{}).Where("mitra_type_temp_id = ?", pid).Take(&MitraTypeTemp{}).Delete(&MitraTypeTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraTypeTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================
