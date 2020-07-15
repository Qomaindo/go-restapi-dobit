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

type MitraType struct {
	MitraTypeID        uint64     `gorm:"column:mitra_type_id;primary_key;" json:"mitra_type_id"`
	MitraTypeCode      string     `gorm:"column:mitra_type_code" json:"mitra_type_code"`
	MitraTypeName      string     `gorm:"column:mitra_type_name" json:"mitra_type_name"`
	MitraTypeDesc      string     `gorm:"column:mitra_type_desc" json:"mitra_type_desc"`
	MitraTypeStatus    uint64     `gorm:"column:mitra_type_status;size:2" json:"mitra_type_status"`
	MitraTypeCreatedBy string     `gorm:"column:mitra_type_created_by;size:125" json:"mitra_type_created_by"`
	MitraTypeCreatedAt time.Time  `gorm:"column:mitra_type_created_at;default:CURRENT_TIMESTAMP" json:"mitra_type_created_at"`
	MitraTypeUpdatedBy string     `gorm:"column:mitra_type_updated_by;size:125" json:"mitra_type_updated_by"`
	MitraTypeUpdatedAt *time.Time `gorm:"column:mitra_type_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_type_created_at"`
	MitraTypeDeletedBy string     `gorm:"column:mitra_type_deleted_by;size:125" json:"mitra_type_deleted_by"`
	MitraTypeDeletedAt *time.Time `gorm:"column:mitra_type_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_type_deleted_at"`
}

type MitraTypeData struct {
	MitraTypeID        uint64     `gorm:"column:mitra_type_id" json:"mitra_type_id"`
	MitraTypeCode      string     `gorm:"column:mitra_type_code" json:"mitra_type_code"`
	MitraTypeName      string     `gorm:"column:mitra_type_name" json:"mitra_type_name"`
	MitraTypeDesc      string     `gorm:"column:mitra_type_desc" json:"mitra_type_desc"`
	MitraTypeStatus    uint64     `gorm:"column:mitra_type_status" json:"mitra_type_status"`
	MitraTypeCreatedBy string     `gorm:"column:mitra_type_created_by" json:"mitra_type_created_by"`
	MitraTypeCreatedAt time.Time  `gorm:"column:mitra_type_created_at" json:"mitra_type_created_at"`
	MitraTypeUpdatedBy string     `gorm:"column:mitra_type_updated_by" json:"mitra_type_updated_by"`
	MitraTypeUpdatedAt *time.Time `gorm:"column:mitra_type_updated_at" json:"mitra_type_updated_at"`
	MitraTypeDeletedBy string     `gorm:"column:mitra_type_deleted_by" json:"mitra_type_deleted_by"`
	MitraTypeDeletedAt *time.Time `gorm:"column:mitra_type_deleted_at" json:"mitra_type_deleted_at"`
}

type ResponseMitraTypes struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    *[]MitraTypeData `json:"data"`
}

type ResponseMitraType struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    *MitraTypeData `json:"data"`
}

type ResponseMitraTypeCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraType) TableName() string {
	return "m_mitra_type"
}

func (MitraTypeData) TableName() string {
	return "m_mitra_type"
}

func (p *MitraType) Prepare() {
	p.MitraTypeCode = html.EscapeString(strings.TrimSpace(p.MitraTypeCode))
	p.MitraTypeName = html.EscapeString(strings.TrimSpace(p.MitraTypeName))
	p.MitraTypeDesc = p.MitraTypeDesc
	p.MitraTypeStatus = p.MitraTypeStatus
	p.MitraTypeCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraTypeCreatedBy))
	p.MitraTypeCreatedAt = time.Now()
	p.MitraTypeUpdatedBy = html.EscapeString(strings.TrimSpace(p.MitraTypeUpdatedBy))
	p.MitraTypeUpdatedAt = p.MitraTypeUpdatedAt
	p.MitraTypeDeletedBy = html.EscapeString(strings.TrimSpace(p.MitraTypeDeletedBy))
	p.MitraTypeDeletedAt = p.MitraTypeDeletedAt
}

func (p *MitraType) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.MitraTypeCode == "" {
			return errors.New("Required Country")
		}
		if p.MitraTypeName == "" {
			return errors.New("Required Province")
		}
		if p.MitraTypeDesc == "" {
			return errors.New("Required Regency")
		}
		return nil
	}
}

func (p *MitraType) SaveMitraType(db *gorm.DB) (*MitraType, error) {
	var err error
	err = db.Debug().Model(&MitraType{}).Create(&p).Error
	if err != nil {
		return &MitraType{}, err
	}
	return p, nil
}

func (p *MitraType) FindAllMitraType(db *gorm.DB) (*[]MitraTypeData, error) {
	var err error
	mitratype := []MitraTypeData{}
	err = db.Debug().Model(&MitraTypeData{}).Limit(100).
		Select(`m_mitra_type.mitra_type_id,
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
		Find(&mitratype).Error
	if err != nil {
		return &[]MitraTypeData{}, err
	}
	return &mitratype, nil
}

func (p *MitraType) FindAllMitraTypeStatus(db *gorm.DB, status uint64) (*[]MitraTypeData, error) {
	var err error
	mitratype := []MitraTypeData{}
	err = db.Debug().Model(&MitraTypeData{}).Limit(100).
		Select(`m_mitra_type.mitra_type_id,
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
		Where("mitra_type_status = ?", status).
		Find(&mitratype).Error
	if err != nil {
		return &[]MitraTypeData{}, err
	}
	return &mitratype, nil
}

func (p *MitraType) FindMitraTypeDataByID(db *gorm.DB, pid uint64) (*MitraType, error) {
	var err error
	err = db.Debug().Model(&MitraType{}).
		Select(`m_mitra_type.mitra_type_id,
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
		Where("mitra_type_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraType{}, err
	}
	return p, nil
}

func (p *MitraType) FindMitraTypeByID(db *gorm.DB, pid uint64) (*MitraTypeData, error) {
	var err error
	mitratype := MitraTypeData{}
	err = db.Debug().Model(&MitraTypeData{}).Limit(100).
		Select(`m_mitra_type.mitra_type_id,
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
		Where("mitra_type_id = ?", pid).
		Find(&mitratype).Error
	if err != nil {
		return &MitraTypeData{}, err
	}
	return &mitratype, nil
}

func (p *MitraType) FindMitraTypeStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraTypeData, error) {
	var err error
	mitratype := MitraTypeData{}
	err = db.Debug().Model(&MitraTypeData{}).Limit(100).
		Select(`m_mitra_type.mitra_type_id,
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
		Where("mitra_type_id = ? AND mitra_type_status = ?", pid, status).
		Find(&mitratype).Error
	if err != nil {
		return &MitraTypeData{}, err
	}
	return &mitratype, nil
}

func (p *MitraType) UpdateMitraType(db *gorm.DB, pid uint64) (*MitraType, error) {
	var err error
	db = db.Debug().Model(&MitraType{}).Where("mitra_type_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_type_code":       p.MitraTypeCode,
			"mitra_type_name":       p.MitraTypeName,
			"mitra_type_desc":       p.MitraTypeDesc,
			"mitra_type_status":     p.MitraTypeStatus,
			"mitra_type_updated_by": p.MitraTypeUpdatedBy,
			"mitra_type_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraType{}).Where("mitra_type_id = ?", pid).Error
	if err != nil {
		return &MitraType{}, err
	}
	return p, nil
}

func (p *MitraType) UpdateMitraTypeStatus(db *gorm.DB, pid uint64) (*MitraType, error) {
	var err error
	db = db.Debug().Model(&MitraType{}).Where("mitra_type_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_type_status":     p.MitraTypeStatus,
			"mitra_type_updated_by": p.MitraTypeUpdatedBy,
			"mitra_type_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraType{}).Where("mitra_type_id = ?", pid).Error
	if err != nil {
		return &MitraType{}, err
	}
	return p, nil
}

func (p *MitraType) UpdateMitraTypeStatusOnly(db *gorm.DB, pid uint64) (*MitraType, error) {
	var err error
	db = db.Debug().Model(&MitraType{}).Where("mitra_type_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_type_status": p.MitraTypeStatus,
		},
	)
	err = db.Debug().Model(&MitraType{}).Where("mitra_type_id = ?", pid).Error
	if err != nil {
		return &MitraType{}, err
	}
	return p, nil
}

func (p *MitraType) UpdateMitraTypeStatusDelete(db *gorm.DB, pid uint64) (*MitraType, error) {
	var err error
	db = db.Debug().Model(&MitraType{}).Where("mitra_type_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_type_status":     3,
			"mitra_type_deleted_by": p.MitraTypeDeletedBy,
			"mitra_type_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraType{}).Where("mitra_type_id = ?", pid).Error
	if err != nil {
		return &MitraType{}, err
	}
	return p, nil
}

func (p *MitraType) DeleteMitraType(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraType{}).Where("mitra_type_id = ?", pid).Take(&MitraType{}).Delete(&MitraType{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraType not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================
