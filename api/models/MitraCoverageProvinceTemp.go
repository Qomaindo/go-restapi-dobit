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

type MitraCoverageProvinceTemp struct {
	MitraCoverageProvinceTempID           uint64    `gorm:"column:mitra_cov_prov_temp_id;primary_key;" json:"mitra_cov_prov_temp_id"`
	MitraCoverageProvinceID               uint64    `gorm:"column:mitra_cov_prov_id" json:"mitra_cov_prov_id"`
	MitraTempID                           uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	ProvinceTempID                        uint64    `gorm:"column:province_temp_id" json:"province_temp_id"`
	MitraCoverageProvinceTempReason       string    `gorm:"column:mitra_cov_prov_temp_reason" json:"mitra_cov_prov_temp_reason"`
	MitraCoverageProvinceTempStatus       uint64    `gorm:"column:mitra_cov_prov_temp_status;size:2" json:"mitra_cov_prov_temp_status"`
	MitraCoverageProvinceTempActionBefore uint64    `gorm:"column:mitra_cov_prov_temp_action_before;size:2" json:"mitra_cov_prov_temp_action_before"`
	MitraCoverageProvinceTempCreatedBy    string    `gorm:"column:mitra_cov_prov_temp_created_by;size:125" json:"mitra_cov_prov_temp_created_by"`
	MitraCoverageProvinceTempCreatedAt    time.Time `gorm:"column:mitra_cov_prov_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_cov_prov_temp_created_at"`
}

type MitraCoverageProvinceTempPend struct {
	MitraCoverageProvinceTempID           *uint64   `gorm:"column:mitra_cov_prov_temp_id;primary_key;" json:"mitra_cov_prov_temp_id"`
	MitraCoverageProvinceID               *uint64   `gorm:"column:mitra_cov_prov_id" json:"mitra_cov_prov_id"`
	MitraTempID                           uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	ProvinceTempID                        uint64    `gorm:"column:province_temp_id" json:"province_temp_id"`
	MitraCoverageProvinceTempReason       string    `gorm:"column:mitra_cov_prov_temp_reason" json:"mitra_cov_prov_temp_reason"`
	MitraCoverageProvinceTempStatus       uint64    `gorm:"column:mitra_cov_prov_temp_status;size:2" json:"mitra_cov_prov_temp_status"`
	MitraCoverageProvinceTempActionBefore uint64    `gorm:"column:mitra_cov_prov_temp_action_before;size:2" json:"mitra_cov_prov_temp_action_before"`
	MitraCoverageProvinceTempCreatedBy    string    `gorm:"column:mitra_cov_prov_temp_created_by;size:125" json:"mitra_cov_prov_temp_created_by"`
	MitraCoverageProvinceTempCreatedAt    time.Time `gorm:"column:mitra_cov_prov_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_cov_prov_temp_created_at"`
}

type MitraCoverageProvinceTempPendData struct {
	MitraCoverageProvinceTempID           uint64    `gorm:"column:mitra_cov_prov_temp_id;primary_key;" json:"mitra_cov_prov_temp_id"`
	MitraCoverageProvinceID               uint64    `gorm:"column:mitra_cov_prov_id" json:"mitra_cov_prov_id"`
	MitraTempID                           uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	ProvinceTempID                        uint64    `gorm:"column:province_temp_id" json:"province_temp_id"`
	ProvinceTempName                      string    `gorm:"column:province_temp_name" json:"province_temp_name"`
	MitraCoverageProvinceTempReason       string    `gorm:"column:mitra_cov_prov_temp_reason" json:"mitra_cov_prov_temp_reason"`
	MitraCoverageProvinceTempStatus       uint64    `gorm:"column:mitra_cov_prov_temp_status;size:2" json:"mitra_cov_prov_temp_status"`
	MitraCoverageProvinceTempActionBefore uint64    `gorm:"column:mitra_cov_prov_temp_action_before;size:2" json:"mitra_cov_prov_temp_action_before"`
	MitraCoverageProvinceTempCreatedBy    string    `gorm:"column:mitra_cov_prov_temp_created_by;size:125" json:"mitra_cov_prov_temp_created_by"`
	MitraCoverageProvinceTempCreatedAt    time.Time `gorm:"column:mitra_cov_prov_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_cov_prov_temp_created_at"`
}

type MitraCoverageProvinceTempData struct {
	MitraCoverageProvinceTempID           uint64    `gorm:"column:mitra_cov_prov_temp_id" json:"mitra_cov_prov_temp_id"`
	MitraTempID                           uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	MitraTempName                         string    `gorm:"column:mitra_temp_name" json:"mitra_temp_name"`
	ProvinceTempID                        uint64    `gorm:"column:province_temp_id" json:"province_temp_id"`
	ProvinceTempName                      string    `gorm:"column:province_temp_name" json:"province_temp_name"`
	MitraID                               uint64    `gorm:"column:mitra_id" json:"mitra_id"`
	MitraName                             string    `gorm:"column:mitra_name" json:"mitra_name"`
	ProvinceID                            uint64    `gorm:"column:province_id" json:"province_id"`
	ProvinceName                          string    `gorm:"column:province_name" json:"province_name"`
	MitraCoverageProvinceTempReason       string    `gorm:"column:mitra_cov_prov_temp_reason" json:"mitra_cov_prov_temp_reason"`
	MitraCoverageProvinceTempStatus       uint64    `gorm:"column:mitra_cov_prov_temp_status;size:2" json:"mitra_cov_prov_temp_status"`
	MitraCoverageProvinceTempActionBefore uint64    `gorm:"column:mitra_cov_prov_temp_action_before;size:2" json:"mitra_cov_prov_temp_action_before"`
	MitraCoverageProvinceTempCreatedBy    string    `gorm:"column:mitra_cov_prov_temp_created_by;size:125" json:"mitra_cov_prov_temp_created_by"`
	MitraCoverageProvinceTempCreatedAt    time.Time `gorm:"column:mitra_cov_prov_temp_created_at" json:"mitra_cov_prov_temp_created_at"`
	MitraCoverageProvinceID               uint64    `gorm:"column:mitra_cov_prov_id" json:"mitra_cov_prov_id"`
}

type ResponseMitraCoverageProvinceTemps struct {
	Status  int                              `json:"status"`
	Message string                           `json:"message"`
	Data    *[]MitraCoverageProvinceTempData `json:"data"`
}

type ResponseMitraCoverageProvinceTemp struct {
	Status  int                            `json:"status"`
	Message string                         `json:"message"`
	Data    *MitraCoverageProvinceTempData `json:"data"`
}

type ResponseMitraCoverageProvinceTempsPend struct {
	Status  int                              `json:"status"`
	Message string                           `json:"message"`
	Data    *[]MitraCoverageProvinceTempPend `json:"data"`
}

type ResponseMitraCoverageProvinceTempPend struct {
	Status  int                            `json:"status"`
	Message string                         `json:"message"`
	Data    *MitraCoverageProvinceTempPend `json:"data"`
}

type ResponseMitraCoverageProvinceTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraCoverageProvinceTemp) TableName() string {
	return "m_mitra_coverage_province_temp"
}

func (MitraCoverageProvinceTempPend) TableName() string {
	return "m_mitra_coverage_province_temp"
}

func (MitraCoverageProvinceTempPendData) TableName() string {
	return "m_mitra_coverage_province_temp"
}

func (MitraCoverageProvinceTempData) TableName() string {
	return "m_mitra_coverage_province_temp"
}

func (p *MitraCoverageProvinceTemp) Prepare() {
	p.MitraCoverageProvinceID = p.MitraCoverageProvinceID
	p.MitraTempID = p.MitraTempID
	p.ProvinceTempID = p.ProvinceTempID
	p.MitraCoverageProvinceTempReason = p.MitraCoverageProvinceTempReason
	p.MitraCoverageProvinceTempStatus = p.MitraCoverageProvinceTempStatus
	p.MitraCoverageProvinceTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraCoverageProvinceTempCreatedBy))
	p.MitraCoverageProvinceTempCreatedAt = time.Now()
}

func (p *MitraCoverageProvinceTempPend) Prepare() {
	p.MitraCoverageProvinceTempID = nil
	p.MitraCoverageProvinceID = nil
	p.MitraTempID = p.MitraTempID
	p.ProvinceTempID = p.ProvinceTempID
	p.MitraCoverageProvinceTempReason = p.MitraCoverageProvinceTempReason
	p.MitraCoverageProvinceTempStatus = p.MitraCoverageProvinceTempStatus
	p.MitraCoverageProvinceTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraCoverageProvinceTempCreatedBy))
	p.MitraCoverageProvinceTempCreatedAt = time.Now()
}

func (p *MitraCoverageProvinceTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	case "insertupdate":
		if p.MitraTempID == 0 {
			return errors.New("Negara Wajib diisi")
		}
		if p.ProvinceTempID == 0 {
			return errors.New("Provinsi Wajib diisi")
		}
		if p.MitraCoverageProvinceTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.MitraCoverageProvinceTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	case "reason":
		if p.MitraCoverageProvinceTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.MitraCoverageProvinceTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	default:
		if p.MitraCoverageProvinceTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil
	}
}

func (p *MitraCoverageProvinceTemp) SaveMitraCoverageProvinceTemp(db *gorm.DB) (*MitraCoverageProvinceTemp, error) {
	var err error
	err = db.Debug().Model(&MitraCoverageProvinceTemp{}).Create(&p).Error
	if err != nil {
		return &MitraCoverageProvinceTemp{}, err
	}
	return p, nil
}

func (p *MitraCoverageProvinceTempPend) SaveMitraCoverageProvinceTempPend(db *gorm.DB) (*MitraCoverageProvinceTempPend, error) {
	var err error
	err = db.Debug().Model(&MitraCoverageProvinceTempPend{}).Create(&p).Error
	if err != nil {
		return &MitraCoverageProvinceTempPend{}, err
	}
	return p, nil
}

func (p *MitraCoverageProvinceTemp) FindAllMitraCoverageProvinceTemp(db *gorm.DB) (*[]MitraCoverageProvinceTempPend, error) {
	var err error
	mitracoverageprovince := []MitraCoverageProvinceTempPend{}
	err = db.Debug().Model(&MitraCoverageProvinceTempPend{}).Limit(100).
		Select(`m_mitra_coverage_province_temp.mitra_cov_prov_temp_id,
				m_mitra_coverage_province_temp.mitra_cov_prov_id,
				m_mitra_coverage_province_temp.mitra_temp_id,
				m_mitra_coverage_province_temp.province_temp_id,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_reason,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_status,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_action_before,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_created_by,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_created_at`).
		Find(&mitracoverageprovince).Error
	if err != nil {
		return &[]MitraCoverageProvinceTempPend{}, err
	}
	return &mitracoverageprovince, nil
}

func (p *MitraCoverageProvinceTemp) FindAllMitraCoverageProvinceTempPendingActive(db *gorm.DB) (*[]MitraCoverageProvinceTempPend, error) {
	var err error
	mitracoverageprovince := []MitraCoverageProvinceTempPend{}
	err = db.Debug().Model(&MitraCoverageProvinceTempPend{}).Limit(100).
		Select(`m_mitra_coverage_province_temp.mitra_cov_prov_temp_id,
				m_mitra_coverage_province_temp.mitra_cov_prov_id,
				m_mitra_coverage_province_temp.mitra_temp_id,
				m_mitra_coverage_province_temp.province_temp_id,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_reason,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_status,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_action_before,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_created_by,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_created_at`).
		Where("mitra_cov_prov_temp_status = ?", 11).
		Find(&mitracoverageprovince).Error
	if err != nil {
		return &[]MitraCoverageProvinceTempPend{}, err
	}
	return &mitracoverageprovince, nil
}

func (p *MitraCoverageProvinceTemp) FindAllMitraCoverageProvinceTempStatus(db *gorm.DB, status uint64) (*[]MitraCoverageProvinceTempData, error) {
	var err error
	mitracoverageprovince := []MitraCoverageProvinceTempData{}
	err = db.Debug().Model(&MitraCoverageProvinceTempData{}).Limit(100).
		Select(`m_mitra_coverage_province_temp.mitra_cov_prov_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_province_temp.province_id as province_temp_id,
				m_province_temp.province_name as province_temp_name,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_reason,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_status,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_action_before,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_created_by,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_created_at,
				m_mitra_coverage_province.mitra_cov_prov_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_province.province_id,
				m_province.province_name,
				m_mitra_coverage_province.mitra_cov_prov_status,
				m_mitra_coverage_province.mitra_cov_prov_created_by,
				m_mitra_coverage_province.mitra_cov_prov_created_at,
				m_mitra_coverage_province.mitra_cov_prov_updated_by,
				m_mitra_coverage_province.mitra_cov_prov_updated_at,
				m_mitra_coverage_province.mitra_cov_prov_deleted_by,
				m_mitra_coverage_province.mitra_cov_prov_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_coverage_province_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_province m_province_temp on m_mitra_coverage_province_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_mitra_coverage_province on m_mitra_coverage_province_temp.mitra_cov_prov_id=m_mitra_coverage_province.mitra_cov_prov_id").
		Joins("JOIN m_mitra on m_mitra_coverage_province.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_province on m_mitra_coverage_province.province_id=m_province.province_id").
		Where("mitra_cov_prov_temp_status = ?", status).
		Find(&mitracoverageprovince).Error
	if err != nil {
		return &[]MitraCoverageProvinceTempData{}, err
	}
	return &mitracoverageprovince, nil
}

func (p *MitraCoverageProvinceTemp) FindMitraCoverageProvinceTempDataByID(db *gorm.DB, pid uint64) (*MitraCoverageProvinceTemp, error) {
	var err error
	err = db.Debug().Model(&MitraCoverageProvinceTemp{}).Where("mitra_cov_prov_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraCoverageProvinceTemp{}, err
	}
	return p, nil
}

func (p *MitraCoverageProvinceTemp) FindMitraCoverageProvinceTempByIDPendingActive(db *gorm.DB, pid uint64) (*MitraCoverageProvinceTempPend, error) {
	var err error
	mitracoverageprovince := MitraCoverageProvinceTempPend{}
	err = db.Debug().Model(&MitraCoverageProvinceTempPend{}).Limit(100).
		Select(`m_mitra_coverage_province_temp.mitra_cov_prov_temp_id,
				m_mitra_coverage_province_temp.mitra_cov_prov_id,
				m_mitra_coverage_province_temp.mitra_temp_id,
				m_mitra_coverage_province_temp.province_temp_id,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_reason,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_status,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_action_before,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_created_by,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_created_at`).
		Where("mitra_cov_prov_temp_id = ? AND mitra_cov_prov_temp_status = ?", pid, 11).
		Find(&mitracoverageprovince).Error
	if err != nil {
		return &MitraCoverageProvinceTempPend{}, err
	}
	return &mitracoverageprovince, nil
}

func (p *MitraCoverageProvinceTemp) FindMitraCoverageProvinceTempByID(db *gorm.DB, pid uint64) (*MitraCoverageProvinceTempData, error) {
	var err error
	mitracoverageprovince := MitraCoverageProvinceTempData{}
	err = db.Debug().Model(&MitraCoverageProvinceTempData{}).Limit(100).
		Select(`m_mitra_coverage_province_temp.mitra_cov_prov_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_province_temp.province_id as province_temp_id,
				m_province_temp.province_name as province_temp_name,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_reason,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_status,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_action_before,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_created_by,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_created_at,
				m_mitra_coverage_province.mitra_cov_prov_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_province.province_id,
				m_province.province_name,
				m_mitra_coverage_province.mitra_cov_prov_status,
				m_mitra_coverage_province.mitra_cov_prov_created_by,
				m_mitra_coverage_province.mitra_cov_prov_created_at,
				m_mitra_coverage_province.mitra_cov_prov_updated_by,
				m_mitra_coverage_province.mitra_cov_prov_updated_at,
				m_mitra_coverage_province.mitra_cov_prov_deleted_by,
				m_mitra_coverage_province.mitra_cov_prov_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_coverage_province_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_province m_province_temp on m_mitra_coverage_province_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_mitra_coverage_province on m_mitra_coverage_province_temp.mitra_cov_prov_id=m_mitra_coverage_province.mitra_cov_prov_id").
		Joins("JOIN m_mitra on m_mitra_coverage_province.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_province on m_mitra_coverage_province.province_id=m_province.province_id").
		Where("mitra_cov_prov_temp_id = ?", pid).
		Find(&mitracoverageprovince).Error
	if err != nil {
		return &MitraCoverageProvinceTempData{}, err
	}
	return &mitracoverageprovince, nil
}

func (p *MitraCoverageProvinceTemp) FindMitraCoverageProvinceTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraCoverageProvinceTempData, error) {
	var err error
	mitracoverageprovince := MitraCoverageProvinceTempData{}
	err = db.Debug().Model(&MitraCoverageProvinceTempData{}).Limit(100).
		Select(`m_mitra_coverage_province_temp.mitra_cov_prov_temp_id,
				m_mitra_temp.mitra_id as mitra_temp_id,
				m_mitra_temp.mitra_name as mitra_temp_name,
				m_province_temp.province_id as province_temp_id,
				m_province_temp.province_name as province_temp_name,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_reason,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_status,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_action_before,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_created_by,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_created_at,
				m_mitra_coverage_province.mitra_cov_prov_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_province.province_id,
				m_province.province_name,
				m_mitra_coverage_province.mitra_cov_prov_status,
				m_mitra_coverage_province.mitra_cov_prov_created_by,
				m_mitra_coverage_province.mitra_cov_prov_created_at,
				m_mitra_coverage_province.mitra_cov_prov_updated_by,
				m_mitra_coverage_province.mitra_cov_prov_updated_at,
				m_mitra_coverage_province.mitra_cov_prov_deleted_by,
				m_mitra_coverage_province.mitra_cov_prov_deleted_at`).
		Joins("JOIN m_mitra m_mitra_temp on m_mitra_coverage_province_temp.mitra_temp_id=m_mitra_temp.mitra_id").
		Joins("JOIN m_province m_province_temp on m_mitra_coverage_province_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_mitra_coverage_province on m_mitra_coverage_province_temp.mitra_cov_prov_id=m_mitra_coverage_province.mitra_cov_prov_id").
		Joins("JOIN m_mitra on m_mitra_coverage_province.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_province on m_mitra_coverage_province.province_id=m_province.province_id").
		Where("mitra_cov_prov_temp_id = ? AND mitra_cov_prov_temp_status = ?", pid, status).
		Find(&mitracoverageprovince).Error
	if err != nil {
		return &MitraCoverageProvinceTempData{}, err
	}
	return &mitracoverageprovince, nil
}

func (p *MitraCoverageProvinceTemp) UpdateMitraCoverageProvinceTemp(db *gorm.DB, pid uint64) (*MitraCoverageProvinceTemp, error) {
	var err error
	db = db.Debug().Model(&MitraCoverageProvinceTemp{}).Where("mitra_cov_prov_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_temp_id":              p.MitraTempID,
			"province_temp_id":           p.ProvinceTempID,
			"mitra_cov_prov_temp_reason": p.MitraCoverageProvinceTempReason,
		},
	)
	err = db.Debug().Model(&MitraCoverageProvinceTemp{}).Where("mitra_cov_prov_temp_id = ?", pid).Error
	if err != nil {
		return &MitraCoverageProvinceTemp{}, err
	}
	return p, nil
}

func (p *MitraCoverageProvinceTemp) UpdateMitraCoverageProvinceTempStatus(db *gorm.DB, pid uint64) (*MitraCoverageProvinceTemp, error) {
	var err error
	db = db.Debug().Model(&MitraCoverageProvinceTemp{}).Where("mitra_cov_prov_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_cov_prov_temp_reason": p.MitraCoverageProvinceTempReason,
			"mitra_cov_prov_temp_status": p.MitraCoverageProvinceTempStatus,
		},
	)
	err = db.Debug().Model(&MitraCoverageProvinceTemp{}).Where("mitra_cov_prov_temp_id = ?", pid).Error
	if err != nil {
		return &MitraCoverageProvinceTemp{}, err
	}
	return p, nil
}

func (p *MitraCoverageProvinceTemp) DeleteMitraCoverageProvinceTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraCoverageProvinceTemp{}).Where("mitra_cov_prov_temp_id = ?", pid).Take(&MitraCoverageProvinceTemp{}).Delete(&MitraCoverageProvinceTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraCoverageProvinceTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

func (p *MitraCoverageProvinceTemp) FindMitraCoverageProvinceTempPendByMitraTempID(db *gorm.DB, pid uint64) ([]MitraCoverageProvinceTempPendData, error) {
	var err error
	mitracoverageprovince := []MitraCoverageProvinceTempPendData{}
	err = db.Debug().Model(&MitraCoverageProvinceTempPendData{}).Limit(100).
		Select(`m_mitra_coverage_province_temp.mitra_cov_prov_temp_id,
				m_mitra_coverage_province_temp.mitra_cov_prov_id,
				m_mitra_coverage_province_temp.mitra_temp_id,
				m_mitra_temp.mitra_temp_name,
				m_mitra_coverage_province_temp.province_temp_id,
				m_province_temp.province_name as province_temp_name,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_reason,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_status,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_action_before,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_created_by,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_created_at`).
		Joins("JOIN m_province m_province_temp on m_mitra_coverage_province_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_mitra_temp on m_mitra_coverage_province_temp.mitra_temp_id=m_mitra_temp.mitra_temp_id").
		Where("m_mitra_coverage_province_temp.mitra_temp_id = ?", pid).
		Find(&mitracoverageprovince).Error
	if err != nil {
		return []MitraCoverageProvinceTempPendData{}, err
	}
	return mitracoverageprovince, nil
}

func (p *MitraCoverageProvinceTemp) FindMitraCoverageProvinceTempPendStatusByMitraTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraCoverageProvinceTempPendData, error) {
	var err error
	mitracoverageprovince := []MitraCoverageProvinceTempPendData{}
	err = db.Debug().Model(&MitraCoverageProvinceTempPendData{}).Limit(100).
		Select(`m_mitra_coverage_province_temp.mitra_cov_prov_temp_id,
				m_mitra_coverage_province_temp.mitra_cov_prov_id,
				m_mitra_coverage_province_temp.mitra_temp_id,
				m_mitra_temp.mitra_temp_name,
				m_mitra_coverage_province_temp.province_temp_id,
				m_province_temp.province_name as province_temp_name,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_reason,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_status,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_action_before,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_created_by,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_created_at`).
		Joins("JOIN m_province m_province_temp on m_mitra_coverage_province_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_mitra_temp on m_mitra_coverage_province_temp.mitra_temp_id=m_mitra_temp.mitra_temp_id").
		Where("m_mitra_coverage_province_temp.mitra_temp_id = ? AND mitra_cov_prov_temp_status = ?", pid, status).
		Find(&mitracoverageprovince).Error
	if err != nil {
		return []MitraCoverageProvinceTempPendData{}, err
	}
	return mitracoverageprovince, nil
}

func (p *MitraCoverageProvinceTemp) FindMitraCoverageProvinceTempByMitraTempID(db *gorm.DB, pid uint64) ([]MitraCoverageProvinceTempData, error) {
	var err error
	mitracoverageprovince := []MitraCoverageProvinceTempData{}
	err = db.Debug().Model(&MitraCoverageProvinceTempData{}).Limit(100).
		Select(`m_mitra_coverage_province_temp.mitra_cov_prov_temp_id,
				m_mitra_coverage_province_temp.mitra_cov_prov_id,
				m_mitra_coverage_province_temp.mitra_temp_id,
				m_mitra_temp.mitra_temp_name,
				m_mitra_coverage_province_temp.province_temp_id,
				m_province_temp.province_name as province_temp_name,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_reason,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_status,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_action_before,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_created_by,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_created_at`).
		Joins("JOIN m_province m_province_temp on m_mitra_coverage_province_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_mitra_temp on m_mitra_coverage_province_temp.mitra_temp_id=m_mitra_temp.mitra_temp_id").
		Where("m_mitra_coverage_province_temp.mitra_temp_id = ?", pid).
		Find(&mitracoverageprovince).Error
	if err != nil {
		return []MitraCoverageProvinceTempData{}, err
	}
	return mitracoverageprovince, nil
}

func (p *MitraCoverageProvinceTemp) FindMitraCoverageProvinceTempStatusByMitraTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraCoverageProvinceTempData, error) {
	var err error
	mitracoverageprovince := []MitraCoverageProvinceTempData{}
	err = db.Debug().Model(&MitraCoverageProvinceTempData{}).Limit(100).
		Select(`m_mitra_coverage_province_temp.mitra_cov_prov_temp_id,
				m_mitra_coverage_province_temp.mitra_cov_prov_id,
				m_mitra_coverage_province_temp.mitra_temp_id,
				m_mitra_temp.mitra_temp_name,
				m_mitra_coverage_province_temp.province_temp_id,
				m_province_temp.province_name as province_temp_name,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_reason,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_status,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_action_before,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_created_by,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_created_at`).
		Joins("JOIN m_province m_province_temp on m_mitra_coverage_province_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_mitra_temp on m_mitra_coverage_province_temp.mitra_temp_id=m_mitra_temp.mitra_temp_id").
		Where("m_mitra_coverage_province_temp.mitra_temp_id = ? AND mitra_cov_prov_temp_status = ?", pid, status).
		Find(&mitracoverageprovince).Error
	if err != nil {
		return []MitraCoverageProvinceTempData{}, err
	}
	return mitracoverageprovince, nil
}

func (p *MitraCoverageProvinceTemp) FindMitraCoverageProvinceTempDataStatusByMitraTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraCoverageProvinceTempData, error) {
	var err error
	mitracoverageprovince := []MitraCoverageProvinceTempData{}
	err = db.Debug().Model(&MitraCoverageProvinceTempData{}).Limit(100).
		Select(`m_mitra_coverage_province_temp.mitra_cov_prov_temp_id,
				m_mitra_coverage_province_temp.mitra_cov_prov_id,
				m_mitra_temp.mitra_temp_id,
				m_mitra_temp.mitra_temp_name,
				m_mitra_coverage_province_temp.province_temp_id,
				m_province_temp.province_name as province_temp_name,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_province.province_id,
				m_province.province_name,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_reason,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_status,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_action_before,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_created_by,
				m_mitra_coverage_province_temp.mitra_cov_prov_temp_created_at`).
		Joins("JOIN m_province m_province_temp on m_mitra_coverage_province_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_mitra_temp on m_mitra_coverage_province_temp.mitra_temp_id=m_mitra_temp.mitra_temp_id").
		Joins("JOIN m_mitra_coverage_province on m_mitra_coverage_province_temp.mitra_cov_prov_id=m_mitra_coverage_province.mitra_cov_prov_id").
		Joins("JOIN m_province on m_mitra_coverage_province.province_id=m_province.province_id").
		Joins("JOIN m_mitra on m_mitra_coverage_province.mitra_id=m_mitra.mitra_id").
		Where("m_mitra_coverage_province_temp.mitra_temp_id = ? AND mitra_cov_prov_temp_status = ?", pid, status).
		Find(&mitracoverageprovince).Error
	if err != nil {
		return []MitraCoverageProvinceTempData{}, err
	}
	return mitracoverageprovince, nil
}
