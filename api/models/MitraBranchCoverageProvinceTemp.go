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

type MitraBranchCoverageProvinceTemp struct {
	MitraBranchCoverageProvinceTempID           *uint64   `gorm:"column:mitra_branch_cov_prov_temp_id;primary_key;" json:"mitra_branch_cov_prov_temp_id"`
	MitraBranchCoverageProvinceID               uint64    `gorm:"column:mitra_branch_cov_prov_id" json:"mitra_branch_cov_prov_id"`
	MitraBranchTempID                           uint64    `gorm:"column:mitra_branch_temp_id" json:"mitra_branch_temp_id"`
	ProvinceTempID                              uint64    `gorm:"column:province_temp_id" json:"province_temp_id"`
	MitraBranchCoverageProvinceTempReason       string    `gorm:"column:mitra_branch_cov_prov_temp_reason" json:"mitra_branch_cov_prov_temp_reason"`
	MitraBranchCoverageProvinceTempStatus       uint64    `gorm:"column:mitra_branch_cov_prov_temp_status;size:2" json:"mitra_branch_cov_prov_temp_status"`
	MitraBranchCoverageProvinceTempActionBefore uint64    `gorm:"column:mitra_branch_cov_prov_temp_action_before;size:2" json:"mitra_branch_cov_prov_temp_action_before"`
	MitraBranchCoverageProvinceTempCreatedBy    string    `gorm:"column:mitra_branch_cov_prov_temp_created_by;size:125" json:"mitra_branch_cov_prov_temp_created_by"`
	MitraBranchCoverageProvinceTempCreatedAt    time.Time `gorm:"column:mitra_branch_cov_prov_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_branch_cov_prov_temp_created_at"`
}

type MitraBranchCoverageProvinceTempPend struct {
	MitraBranchCoverageProvinceTempID           *uint64   `gorm:"column:mitra_branch_cov_prov_temp_id;primary_key;" json:"mitra_branch_cov_prov_temp_id"`
	MitraBranchCoverageProvinceID               *uint64   `gorm:"column:mitra_branch_cov_prov_id" json:"mitra_branch_cov_prov_id"`
	MitraBranchTempID                           uint64    `gorm:"column:mitra_branch_temp_id" json:"mitra_branch_temp_id"`
	ProvinceTempID                              uint64    `gorm:"column:province_temp_id" json:"province_temp_id"`
	MitraBranchCoverageProvinceTempReason       string    `gorm:"column:mitra_branch_cov_prov_temp_reason" json:"mitra_branch_cov_prov_temp_reason"`
	MitraBranchCoverageProvinceTempStatus       uint64    `gorm:"column:mitra_branch_cov_prov_temp_status;size:2" json:"mitra_branch_cov_prov_temp_status"`
	MitraBranchCoverageProvinceTempActionBefore uint64    `gorm:"column:mitra_branch_cov_prov_temp_action_before;size:2" json:"mitra_branch_cov_prov_temp_action_before"`
	MitraBranchCoverageProvinceTempCreatedBy    string    `gorm:"column:mitra_branch_cov_prov_temp_created_by;size:125" json:"mitra_branch_cov_prov_temp_created_by"`
	MitraBranchCoverageProvinceTempCreatedAt    time.Time `gorm:"column:mitra_branch_cov_prov_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_branch_cov_prov_temp_created_at"`
}

type MitraBranchCoverageProvinceTempPendData struct {
	MitraBranchCoverageProvinceTempID           uint64    `gorm:"column:mitra_branch_cov_prov_temp_id;primary_key;" json:"mitra_branch_cov_prov_temp_id"`
	MitraBranchCoverageProvinceID               uint64    `gorm:"column:mitra_branch_cov_prov_id" json:"mitra_branch_cov_prov_id"`
	MitraBranchTempID                           uint64    `gorm:"column:mitra_branch_temp_id" json:"mitra_branch_temp_id"`
	ProvinceTempID                              uint64    `gorm:"column:province_temp_id" json:"province_temp_id"`
	ProvinceTempName                            string    `gorm:"column:province_temp_name" json:"province_temp_name"`
	MitraBranchCoverageProvinceTempReason       string    `gorm:"column:mitra_branch_cov_prov_temp_reason" json:"mitra_branch_cov_prov_temp_reason"`
	MitraBranchCoverageProvinceTempStatus       uint64    `gorm:"column:mitra_branch_cov_prov_temp_status;size:2" json:"mitra_branch_cov_prov_temp_status"`
	MitraBranchCoverageProvinceTempActionBefore uint64    `gorm:"column:mitra_branch_cov_prov_temp_action_before;size:2" json:"mitra_branch_cov_prov_temp_action_before"`
	MitraBranchCoverageProvinceTempCreatedBy    string    `gorm:"column:mitra_branch_cov_prov_temp_created_by;size:125" json:"mitra_branch_cov_prov_temp_created_by"`
	MitraBranchCoverageProvinceTempCreatedAt    time.Time `gorm:"column:mitra_branch_cov_prov_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_branch_cov_prov_temp_created_at"`
}

type MitraBranchCoverageProvinceTempData struct {
	MitraBranchCoverageProvinceTempID           uint64    `gorm:"column:mitra_branch_cov_prov_temp_id" json:"mitra_branch_cov_prov_temp_id"`
	MitraBranchTempID                           uint64    `gorm:"column:mitra_branch_temp_id" json:"mitra_branch_temp_id"`
	MitraBranchTempName                         string    `gorm:"column:mitra_branch_temp_name" json:"mitra_branch_temp_name"`
	ProvinceTempID                              uint64    `gorm:"column:province_temp_id" json:"province_temp_id"`
	ProvinceTempName                            string    `gorm:"column:province_temp_name" json:"province_temp_name"`
	MitraBranchID                               uint64    `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	MitraBranchName                             string    `gorm:"column:mitra_branch_name" json:"mitra_branch_name"`
	ProvinceID                                  uint64    `gorm:"column:province_id" json:"province_id"`
	ProvinceName                                string    `gorm:"column:province_name" json:"province_name"`
	MitraBranchCoverageProvinceTempReason       string    `gorm:"column:mitra_branch_cov_prov_temp_reason" json:"mitra_branch_cov_prov_temp_reason"`
	MitraBranchCoverageProvinceTempStatus       uint64    `gorm:"column:mitra_branch_cov_prov_temp_status;size:2" json:"mitra_branch_cov_prov_temp_status"`
	MitraBranchCoverageProvinceTempActionBefore uint64    `gorm:"column:mitra_branch_cov_prov_temp_action_before;size:2" json:"mitra_branch_cov_prov_temp_action_before"`
	MitraBranchCoverageProvinceTempCreatedBy    string    `gorm:"column:mitra_branch_cov_prov_temp_created_by;size:125" json:"mitra_branch_cov_prov_temp_created_by"`
	MitraBranchCoverageProvinceTempCreatedAt    time.Time `gorm:"column:mitra_branch_cov_prov_temp_created_at" json:"mitra_branch_cov_prov_temp_created_at"`
	MitraBranchCoverageProvinceID               uint64    `gorm:"column:mitra_branch_cov_prov_id" json:"mitra_branch_cov_prov_id"`
}

type ResponseMitraBranchCoverageProvinceTemps struct {
	Status  int                                    `json:"status"`
	Message string                                 `json:"message"`
	Data    *[]MitraBranchCoverageProvinceTempData `json:"data"`
}

type ResponseMitraBranchCoverageProvinceTemp struct {
	Status  int                                  `json:"status"`
	Message string                               `json:"message"`
	Data    *MitraBranchCoverageProvinceTempData `json:"data"`
}

type ResponseMitraBranchCoverageProvinceTempsPend struct {
	Status  int                                    `json:"status"`
	Message string                                 `json:"message"`
	Data    *[]MitraBranchCoverageProvinceTempPend `json:"data"`
}

type ResponseMitraBranchCoverageProvinceTempPend struct {
	Status  int                                  `json:"status"`
	Message string                               `json:"message"`
	Data    *MitraBranchCoverageProvinceTempPend `json:"data"`
}

type ResponseMitraBranchCoverageProvinceTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraBranchCoverageProvinceTemp) TableName() string {
	return "m_mitra_branch_coverage_province_temp"
}

func (MitraBranchCoverageProvinceTempPend) TableName() string {
	return "m_mitra_branch_coverage_province_temp"
}

func (MitraBranchCoverageProvinceTempPendData) TableName() string {
	return "m_mitra_branch_coverage_province_temp"
}

func (MitraBranchCoverageProvinceTempData) TableName() string {
	return "m_mitra_branch_coverage_province_temp"
}

func (p *MitraBranchCoverageProvinceTemp) Prepare() {
	p.MitraBranchCoverageProvinceTempID = nil
	p.MitraBranchCoverageProvinceID = p.MitraBranchCoverageProvinceID
	p.MitraBranchTempID = p.MitraBranchTempID
	p.ProvinceTempID = p.ProvinceTempID
	p.MitraBranchCoverageProvinceTempReason = p.MitraBranchCoverageProvinceTempReason
	p.MitraBranchCoverageProvinceTempStatus = p.MitraBranchCoverageProvinceTempStatus
	p.MitraBranchCoverageProvinceTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraBranchCoverageProvinceTempCreatedBy))
	p.MitraBranchCoverageProvinceTempCreatedAt = time.Now()
}

func (p *MitraBranchCoverageProvinceTempPend) Prepare() {
	p.MitraBranchCoverageProvinceTempID = nil
	p.MitraBranchCoverageProvinceID = nil
	p.MitraBranchTempID = p.MitraBranchTempID
	p.ProvinceTempID = p.ProvinceTempID
	p.MitraBranchCoverageProvinceTempReason = p.MitraBranchCoverageProvinceTempReason
	p.MitraBranchCoverageProvinceTempStatus = p.MitraBranchCoverageProvinceTempStatus
	p.MitraBranchCoverageProvinceTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraBranchCoverageProvinceTempCreatedBy))
	p.MitraBranchCoverageProvinceTempCreatedAt = time.Now()
}

func (p *MitraBranchCoverageProvinceTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	case "insertupdate":
		if p.MitraBranchTempID == 0 {
			return errors.New("Negara Wajib diisi")
		}
		if p.ProvinceTempID == 0 {
			return errors.New("Provinsi Wajib diisi")
		}
		if p.MitraBranchCoverageProvinceTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.MitraBranchCoverageProvinceTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	case "reason":
		if p.MitraBranchCoverageProvinceTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.MitraBranchCoverageProvinceTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	default:
		if p.MitraBranchCoverageProvinceTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil
	}
}

func (p *MitraBranchCoverageProvinceTemp) SaveMitraBranchCoverageProvinceTemp(db *gorm.DB) (*MitraBranchCoverageProvinceTemp, error) {
	var err error
	err = db.Debug().Model(&MitraBranchCoverageProvinceTemp{}).Create(&p).Error
	if err != nil {
		return &MitraBranchCoverageProvinceTemp{}, err
	}
	return p, nil
}

func (p *MitraBranchCoverageProvinceTempPend) SaveMitraBranchCoverageProvinceTempPend(db *gorm.DB) (*MitraBranchCoverageProvinceTempPend, error) {
	var err error
	err = db.Debug().Model(&MitraBranchCoverageProvinceTempPend{}).Create(&p).Error
	if err != nil {
		return &MitraBranchCoverageProvinceTempPend{}, err
	}
	return p, nil
}

func (p *MitraBranchCoverageProvinceTemp) FindAllMitraBranchCoverageProvinceTemp(db *gorm.DB) (*[]MitraBranchCoverageProvinceTempPend, error) {
	var err error
	mitrabranchcoverageprovince := []MitraBranchCoverageProvinceTempPend{}
	err = db.Debug().Model(&MitraBranchCoverageProvinceTempPend{}).Limit(100).
		Select(`m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_id,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_id,
				m_mitra_branch_coverage_province_temp.mitra_branch_temp_id,
				m_mitra_branch_coverage_province_temp.province_temp_id,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_reason,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_status,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_action_before,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_created_by,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_created_at`).
		Find(&mitrabranchcoverageprovince).Error
	if err != nil {
		return &[]MitraBranchCoverageProvinceTempPend{}, err
	}
	return &mitrabranchcoverageprovince, nil
}

func (p *MitraBranchCoverageProvinceTemp) FindAllMitraBranchCoverageProvinceTempPendingActive(db *gorm.DB) (*[]MitraBranchCoverageProvinceTempPend, error) {
	var err error
	mitrabranchcoverageprovince := []MitraBranchCoverageProvinceTempPend{}
	err = db.Debug().Model(&MitraBranchCoverageProvinceTempPend{}).Limit(100).
		Select(`m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_id,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_id,
				m_mitra_branch_coverage_province_temp.mitra_branch_temp_id,
				m_mitra_branch_coverage_province_temp.province_temp_id,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_reason,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_status,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_action_before,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_created_by,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_created_at`).
		Where("mitra_branch_cov_prov_temp_status = ?", 11).
		Find(&mitrabranchcoverageprovince).Error
	if err != nil {
		return &[]MitraBranchCoverageProvinceTempPend{}, err
	}
	return &mitrabranchcoverageprovince, nil
}

func (p *MitraBranchCoverageProvinceTemp) FindAllMitraBranchCoverageProvinceTempStatus(db *gorm.DB, status uint64) (*[]MitraBranchCoverageProvinceTempData, error) {
	var err error
	mitrabranchcoverageprovince := []MitraBranchCoverageProvinceTempData{}
	err = db.Debug().Model(&MitraBranchCoverageProvinceTempData{}).Limit(100).
		Select(`m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_id,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_province_temp.province_id as province_temp_id,
				m_province_temp.province_name as province_temp_name,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_reason,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_status,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_action_before,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_created_by,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_created_at,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_id,
				m_mitra_branch.mitra_branch_id,
				m_mitra.mitra_branch_name,
				m_province.province_id,
				m_province.province_name,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_status,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_created_by,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_created_at,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_updated_by,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_updated_at,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_deleted_by,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_deleted_at`).
		Joins("JOIN m_mitra m_mitra_branch_temp on m_mitra_branch_coverage_province_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_province m_province_temp on m_mitra_branch_coverage_province_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_mitra_branch_coverage_province on m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_id=m_mitra_branch_coverage_province.mitra_branch_cov_prov_id").
		Joins("JOIN m_mitra on m_mitra_branch_coverage_province.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_province on m_mitra_branch_coverage_province.province_id=m_province.province_id").
		Where("mitra_branch_cov_prov_temp_status = ?", status).
		Find(&mitrabranchcoverageprovince).Error
	if err != nil {
		return &[]MitraBranchCoverageProvinceTempData{}, err
	}
	return &mitrabranchcoverageprovince, nil
}

func (p *MitraBranchCoverageProvinceTemp) FindMitraBranchCoverageProvinceTempDataByID(db *gorm.DB, pid uint64) (*MitraBranchCoverageProvinceTemp, error) {
	var err error
	err = db.Debug().Model(&MitraBranchCoverageProvinceTemp{}).Where("mitra_branch_cov_prov_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraBranchCoverageProvinceTemp{}, err
	}
	return p, nil
}

func (p *MitraBranchCoverageProvinceTemp) FindMitraBranchCoverageProvinceTempByIDPendingActive(db *gorm.DB, pid uint64) (*MitraBranchCoverageProvinceTempPend, error) {
	var err error
	mitrabranchcoverageprovince := MitraBranchCoverageProvinceTempPend{}
	err = db.Debug().Model(&MitraBranchCoverageProvinceTempPend{}).Limit(100).
		Select(`m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_id,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_id,
				m_mitra_branch_coverage_province_temp.mitra_branch_temp_id,
				m_mitra_branch_coverage_province_temp.province_temp_id,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_reason,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_status,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_action_before,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_created_by,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_created_at`).
		Where("mitra_branch_cov_prov_temp_id = ? AND mitra_branch_cov_prov_temp_status = ?", pid, 11).
		Find(&mitrabranchcoverageprovince).Error
	if err != nil {
		return &MitraBranchCoverageProvinceTempPend{}, err
	}
	return &mitrabranchcoverageprovince, nil
}

func (p *MitraBranchCoverageProvinceTemp) FindMitraBranchCoverageProvinceTempByID(db *gorm.DB, pid uint64) (*MitraBranchCoverageProvinceTempData, error) {
	var err error
	mitrabranchcoverageprovince := MitraBranchCoverageProvinceTempData{}
	err = db.Debug().Model(&MitraBranchCoverageProvinceTempData{}).Limit(100).
		Select(`m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_id,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_province_temp.province_id as province_temp_id,
				m_province_temp.province_name as province_temp_name,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_reason,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_status,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_action_before,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_created_by,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_created_at,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_id,
				m_mitra_branch.mitra_branch_id,
				m_mitra.mitra_branch_name,
				m_province.province_id,
				m_province.province_name,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_status,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_created_by,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_created_at,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_updated_by,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_updated_at,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_deleted_by,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_deleted_at`).
		Joins("JOIN m_mitra m_mitra_branch_temp on m_mitra_branch_coverage_province_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_province m_province_temp on m_mitra_branch_coverage_province_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_mitra_branch_coverage_province on m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_id=m_mitra_branch_coverage_province.mitra_branch_cov_prov_id").
		Joins("JOIN m_mitra on m_mitra_branch_coverage_province.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_province on m_mitra_branch_coverage_province.province_id=m_province.province_id").
		Where("mitra_branch_cov_prov_temp_id = ?", pid).
		Find(&mitrabranchcoverageprovince).Error
	if err != nil {
		return &MitraBranchCoverageProvinceTempData{}, err
	}
	return &mitrabranchcoverageprovince, nil
}

func (p *MitraBranchCoverageProvinceTemp) FindMitraBranchCoverageProvinceTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraBranchCoverageProvinceTempData, error) {
	var err error
	mitrabranchcoverageprovince := MitraBranchCoverageProvinceTempData{}
	err = db.Debug().Model(&MitraBranchCoverageProvinceTempData{}).Limit(100).
		Select(`m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_id,
				m_mitra_branch_temp.mitra_branch_id as mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_name as mitra_branch_temp_name,
				m_province_temp.province_id as province_temp_id,
				m_province_temp.province_name as province_temp_name,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_reason,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_status,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_action_before,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_created_by,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_created_at,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_id,
				m_mitra_branch.mitra_branch_id,
				m_mitra.mitra_branch_name,
				m_province.province_id,
				m_province.province_name,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_status,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_created_by,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_created_at,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_updated_by,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_updated_at,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_deleted_by,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_deleted_at`).
		Joins("JOIN m_mitra m_mitra_branch_temp on m_mitra_branch_coverage_province_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_id").
		Joins("JOIN m_province m_province_temp on m_mitra_branch_coverage_province_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_mitra_branch_coverage_province on m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_id=m_mitra_branch_coverage_province.mitra_branch_cov_prov_id").
		Joins("JOIN m_mitra on m_mitra_branch_coverage_province.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_province on m_mitra_branch_coverage_province.province_id=m_province.province_id").
		Where("mitra_branch_cov_prov_temp_id = ? AND mitra_branch_cov_prov_temp_status = ?", pid, status).
		Find(&mitrabranchcoverageprovince).Error
	if err != nil {
		return &MitraBranchCoverageProvinceTempData{}, err
	}
	return &mitrabranchcoverageprovince, nil
}

func (p *MitraBranchCoverageProvinceTemp) UpdateMitraBranchCoverageProvinceTemp(db *gorm.DB, pid uint64) (*MitraBranchCoverageProvinceTemp, error) {
	var err error
	db = db.Debug().Model(&MitraBranchCoverageProvinceTemp{}).Where("mitra_branch_cov_prov_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_branch_temp_id":              p.MitraBranchTempID,
			"province_temp_id":                  p.ProvinceTempID,
			"mitra_branch_cov_prov_temp_reason": p.MitraBranchCoverageProvinceTempReason,
		},
	)
	err = db.Debug().Model(&MitraBranchCoverageProvinceTemp{}).Where("mitra_branch_cov_prov_temp_id = ?", pid).Error
	if err != nil {
		return &MitraBranchCoverageProvinceTemp{}, err
	}
	return p, nil
}

func (p *MitraBranchCoverageProvinceTemp) UpdateMitraBranchCoverageProvinceTempStatus(db *gorm.DB, pid uint64) (*MitraBranchCoverageProvinceTemp, error) {
	var err error
	db = db.Debug().Model(&MitraBranchCoverageProvinceTemp{}).Where("mitra_branch_cov_prov_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_branch_cov_prov_temp_reason": p.MitraBranchCoverageProvinceTempReason,
			"mitra_branch_cov_prov_temp_status": p.MitraBranchCoverageProvinceTempStatus,
		},
	)
	err = db.Debug().Model(&MitraBranchCoverageProvinceTemp{}).Where("mitra_branch_cov_prov_temp_id = ?", pid).Error
	if err != nil {
		return &MitraBranchCoverageProvinceTemp{}, err
	}
	return p, nil
}

func (p *MitraBranchCoverageProvinceTemp) DeleteMitraBranchCoverageProvinceTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraBranchCoverageProvinceTemp{}).Where("mitra_branch_cov_prov_temp_id = ?", pid).Take(&MitraBranchCoverageProvinceTemp{}).Delete(&MitraBranchCoverageProvinceTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraBranchCoverageProvinceTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

func (p *MitraBranchCoverageProvinceTemp) FindMitraBranchCoverageProvinceTempPendByMitraBranchTempID(db *gorm.DB, pid uint64) ([]MitraBranchCoverageProvinceTempPendData, error) {
	var err error
	mitrabranchcoverageprovince := []MitraBranchCoverageProvinceTempPendData{}
	err = db.Debug().Model(&MitraBranchCoverageProvinceTempPendData{}).Limit(100).
		Select(`m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_id,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_id,
				m_mitra_branch_coverage_province_temp.mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_branch_coverage_province_temp.province_temp_id,
				m_province_temp.province_name as province_temp_name,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_reason,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_status,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_action_before,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_created_by,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_created_at`).
		Joins("JOIN m_province m_province_temp on m_mitra_branch_coverage_province_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_mitra_branch_temp on m_mitra_branch_coverage_province_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_temp_id").
		Where("m_mitra_branch_coverage_province_temp.mitra_branch_temp_id = ?", pid).
		Find(&mitrabranchcoverageprovince).Error
	if err != nil {
		return []MitraBranchCoverageProvinceTempPendData{}, err
	}
	return mitrabranchcoverageprovince, nil
}

func (p *MitraBranchCoverageProvinceTemp) FindMitraBranchCoverageProvinceTempPendStatusByMitraBranchTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraBranchCoverageProvinceTempPendData, error) {
	var err error
	mitrabranchcoverageprovince := []MitraBranchCoverageProvinceTempPendData{}
	err = db.Debug().Model(&MitraBranchCoverageProvinceTempPendData{}).Limit(100).
		Select(`m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_id,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_id,
				m_mitra_branch_coverage_province_temp.mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_branch_coverage_province_temp.province_temp_id,
				m_province_temp.province_name as province_temp_name,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_reason,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_status,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_action_before,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_created_by,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_created_at`).
		Joins("JOIN m_province m_province_temp on m_mitra_branch_coverage_province_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_mitra_branch_temp on m_mitra_branch_coverage_province_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_temp_id").
		Where("m_mitra_branch_coverage_province_temp.mitra_branch_temp_id = ? AND mitra_branch_cov_prov_temp_status = ?", pid, status).
		Find(&mitrabranchcoverageprovince).Error
	if err != nil {
		return []MitraBranchCoverageProvinceTempPendData{}, err
	}
	return mitrabranchcoverageprovince, nil
}

func (p *MitraBranchCoverageProvinceTemp) FindMitraBranchCoverageProvinceTempByMitraBranchTempID(db *gorm.DB, pid uint64) ([]MitraBranchCoverageProvinceTempData, error) {
	var err error
	mitrabranchcoverageprovince := []MitraBranchCoverageProvinceTempData{}
	err = db.Debug().Model(&MitraBranchCoverageProvinceTempData{}).Limit(100).
		Select(`m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_id,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_id,
				m_mitra_branch_coverage_province_temp.mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_branch_coverage_province_temp.province_temp_id,
				m_province_temp.province_name as province_temp_name,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_reason,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_status,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_action_before,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_created_by,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_created_at`).
		Joins("JOIN m_province m_province_temp on m_mitra_branch_coverage_province_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_mitra_branch_temp on m_mitra_branch_coverage_province_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_temp_id").
		Where("m_mitra_branch_coverage_province_temp.mitra_branch_temp_id = ?", pid).
		Find(&mitrabranchcoverageprovince).Error
	if err != nil {
		return []MitraBranchCoverageProvinceTempData{}, err
	}
	return mitrabranchcoverageprovince, nil
}

func (p *MitraBranchCoverageProvinceTemp) FindMitraBranchCoverageProvinceTempStatusByMitraBranchTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraBranchCoverageProvinceTempData, error) {
	var err error
	mitrabranchcoverageprovince := []MitraBranchCoverageProvinceTempData{}
	err = db.Debug().Model(&MitraBranchCoverageProvinceTempData{}).Limit(100).
		Select(`m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_id,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_id,
				m_mitra_branch_coverage_province_temp.mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_branch_coverage_province_temp.province_temp_id,
				m_province_temp.province_name as province_temp_name,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_reason,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_status,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_action_before,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_created_by,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_created_at`).
		Joins("JOIN m_province m_province_temp on m_mitra_branch_coverage_province_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_mitra_branch_temp on m_mitra_branch_coverage_province_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_temp_id").
		Where("m_mitra_branch_coverage_province_temp.mitra_branch_temp_id = ? AND mitra_branch_cov_prov_temp_status = ?", pid, status).
		Find(&mitrabranchcoverageprovince).Error
	if err != nil {
		return []MitraBranchCoverageProvinceTempData{}, err
	}
	return mitrabranchcoverageprovince, nil
}

func (p *MitraBranchCoverageProvinceTemp) FindMitraBranchCoverageProvinceTempDataStatusByMitraBranchTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraBranchCoverageProvinceTempData, error) {
	var err error
	mitrabranchcoverageprovince := []MitraBranchCoverageProvinceTempData{}
	err = db.Debug().Model(&MitraBranchCoverageProvinceTempData{}).Limit(100).
		Select(`m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_id,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_id,
				m_mitra_branch_temp.mitra_branch_temp_id,
				m_mitra_branch_temp.mitra_branch_temp_name,
				m_mitra_branch_coverage_province_temp.province_temp_id,
				m_province_temp.province_name as province_temp_name,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_province.province_id,
				m_province.province_name,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_reason,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_status,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_action_before,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_created_by,
				m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_temp_created_at`).
		Joins("JOIN m_province m_province_temp on m_mitra_branch_coverage_province_temp.province_temp_id=m_province_temp.province_id").
		Joins("JOIN m_mitra_branch_temp on m_mitra_branch_coverage_province_temp.mitra_branch_temp_id=m_mitra_branch_temp.mitra_branch_temp_id").
		Joins("JOIN m_mitra_branch_coverage_province on m_mitra_branch_coverage_province_temp.mitra_branch_cov_prov_id=m_mitra_branch_coverage_province.mitra_branch_cov_prov_id").
		Joins("JOIN m_province on m_mitra_branch_coverage_province.province_id=m_province.province_id").
		Joins("JOIN m_mitra_branch on m_mitra_branch_coverage_province.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Where("m_mitra_branch_coverage_province_temp.mitra_branch_temp_id = ? AND mitra_branch_cov_prov_temp_status = ?", pid, status).
		Find(&mitrabranchcoverageprovince).Error
	if err != nil {
		return []MitraBranchCoverageProvinceTempData{}, err
	}
	return mitrabranchcoverageprovince, nil
}
