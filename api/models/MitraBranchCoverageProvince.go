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

type MitraBranchCoverageProvince struct {
	MitraBranchCoverageProvinceID        *uint64    `gorm:"column:mitra_branch_cov_prov_id;primary_key;" json:"mitra_branch_cov_prov_id"`
	MitraBranchID                        uint64     `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	ProvinceID                           uint64     `gorm:"column:province_id" json:"province_id"`
	MitraBranchCoverageProvinceStatus    uint64     `gorm:"column:mitra_branch_cov_prov_status;size:2" json:"mitra_branch_cov_prov_status"`
	MitraBranchCoverageProvinceCreatedBy string     `gorm:"column:mitra_branch_cov_prov_created_by;size:125" json:"mitra_branch_cov_prov_created_by"`
	MitraBranchCoverageProvinceCreatedAt time.Time  `gorm:"column:mitra_branch_cov_prov_created_at;default:CURRENT_TIMESTAMP" json:"mitra_branch_cov_prov_created_at"`
	MitraBranchCoverageProvinceUpdatedBy string     `gorm:"column:mitra_branch_cov_prov_updated_by;size:125" json:"mitra_branch_cov_prov_updated_by"`
	MitraBranchCoverageProvinceUpdatedAt *time.Time `gorm:"column:mitra_branch_cov_prov_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_branch_cov_prov_created_at"`
	MitraBranchCoverageProvinceDeletedBy string     `gorm:"column:mitra_branch_cov_prov_deleted_by;size:125" json:"mitra_branch_cov_prov_deleted_by"`
	MitraBranchCoverageProvinceDeletedAt *time.Time `gorm:"column:mitra_branch_cov_prov_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_branch_cov_prov_deleted_at"`
}

type MitraBranchCoverageProvinceData struct {
	MitraBranchCoverageProvinceID        uint64     `gorm:"column:mitra_branch_cov_prov_id" json:"mitra_branch_cov_prov_id"`
	MitraBranchID                        uint64     `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	MitraBranchName                      string     `gorm:"column:mitra_branch_name" json:"mitra_branch_name"`
	ProvinceID                           uint64     `gorm:"column:province_id" json:"province_id"`
	ProvinceName                         string     `gorm:"column:province_name" json:"province_name"`
	MitraBranchCoverageProvinceStatus    uint64     `gorm:"column:mitra_branch_cov_prov_status" json:"mitra_branch_cov_prov_status"`
	MitraBranchCoverageProvinceCreatedBy string     `gorm:"column:mitra_branch_cov_prov_created_by" json:"mitra_branch_cov_prov_created_by"`
	MitraBranchCoverageProvinceCreatedAt time.Time  `gorm:"column:mitra_branch_cov_prov_created_at" json:"mitra_branch_cov_prov_created_at"`
	MitraBranchCoverageProvinceUpdatedBy string     `gorm:"column:mitra_branch_cov_prov_updated_by" json:"mitra_branch_cov_prov_updated_by"`
	MitraBranchCoverageProvinceUpdatedAt *time.Time `gorm:"column:mitra_branch_cov_prov_updated_at" json:"mitra_branch_cov_prov_updated_at"`
	MitraBranchCoverageProvinceDeletedBy string     `gorm:"column:mitra_branch_cov_prov_deleted_by" json:"mitra_branch_cov_prov_deleted_by"`
	MitraBranchCoverageProvinceDeletedAt *time.Time `gorm:"column:mitra_branch_cov_prov_deleted_at" json:"mitra_branch_cov_prov_deleted_at"`
}

type ResponseMitraBranchCoverageProvinces struct {
	Status  int                                `json:"status"`
	Message string                             `json:"message"`
	Data    *[]MitraBranchCoverageProvinceData `json:"data"`
}

type ResponseMitraBranchCoverageProvince struct {
	Status  int                              `json:"status"`
	Message string                           `json:"message"`
	Data    *MitraBranchCoverageProvinceData `json:"data"`
}

type ResponseMitraBranchCoverageProvinceCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraBranchCoverageProvince) TableName() string {
	return "m_mitra_branch_coverage_province"
}

func (MitraBranchCoverageProvinceData) TableName() string {
	return "m_mitra_branch_coverage_province"
}

func (p *MitraBranchCoverageProvince) Prepare() {
	p.MitraBranchCoverageProvinceID = nil
	p.MitraBranchID = p.MitraBranchID
	p.ProvinceID = p.ProvinceID
	p.MitraBranchCoverageProvinceStatus = p.MitraBranchCoverageProvinceStatus
	p.MitraBranchCoverageProvinceCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraBranchCoverageProvinceCreatedBy))
	p.MitraBranchCoverageProvinceCreatedAt = time.Now()
	p.MitraBranchCoverageProvinceUpdatedBy = html.EscapeString(strings.TrimSpace(p.MitraBranchCoverageProvinceUpdatedBy))
	p.MitraBranchCoverageProvinceUpdatedAt = p.MitraBranchCoverageProvinceUpdatedAt
	p.MitraBranchCoverageProvinceDeletedBy = html.EscapeString(strings.TrimSpace(p.MitraBranchCoverageProvinceDeletedBy))
	p.MitraBranchCoverageProvinceDeletedAt = p.MitraBranchCoverageProvinceDeletedAt
}

func (p *MitraBranchCoverageProvince) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.MitraBranchID == 0 {
			return errors.New("Required Country")
		}
		if p.ProvinceID == 0 {
			return errors.New("Required Province")
		}
		return nil
	}
}

func (p *MitraBranchCoverageProvince) SaveMitraBranchCoverageProvince(db *gorm.DB) (*MitraBranchCoverageProvince, error) {
	var err error
	err = db.Debug().Model(&MitraBranchCoverageProvince{}).Create(&p).Error
	if err != nil {
		return &MitraBranchCoverageProvince{}, err
	}
	return p, nil
}

func (p *MitraBranchCoverageProvince) FindAllMitraBranchCoverageProvince(db *gorm.DB) (*[]MitraBranchCoverageProvinceData, error) {
	var err error
	mitrabranchcoverageprovince := []MitraBranchCoverageProvinceData{}
	err = db.Debug().Model(&MitraBranchCoverageProvinceData{}).Limit(100).
		Select(`m_mitra_branch_coverage_province.mitra_branch_cov_prov_id,
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
		Joins("JOIN m_mitra on m_mitra_branch_coverage_province.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_province on m_mitra_branch_coverage_province.province_id=m_province.province_id").
		Find(&mitrabranchcoverageprovince).Error
	if err != nil {
		return &[]MitraBranchCoverageProvinceData{}, err
	}
	return &mitrabranchcoverageprovince, nil
}

func (p *MitraBranchCoverageProvince) FindAllMitraBranchCoverageProvinceStatus(db *gorm.DB, status uint64) (*[]MitraBranchCoverageProvinceData, error) {
	var err error
	mitrabranchcoverageprovince := []MitraBranchCoverageProvinceData{}
	err = db.Debug().Model(&MitraBranchCoverageProvinceData{}).Limit(100).
		Select(`m_mitra_branch_coverage_province.mitra_branch_cov_prov_id,
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
		Joins("JOIN m_mitra on m_mitra_branch_coverage_province.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_province on m_mitra_branch_coverage_province.province_id=m_province.province_id").
		Where("mitra_branch_cov_prov_status = ?", status).
		Find(&mitrabranchcoverageprovince).Error
	if err != nil {
		return &[]MitraBranchCoverageProvinceData{}, err
	}
	return &mitrabranchcoverageprovince, nil
}

func (p *MitraBranchCoverageProvince) FindMitraBranchCoverageProvinceDataByID(db *gorm.DB, pid uint64) (*MitraBranchCoverageProvince, error) {
	var err error
	err = db.Debug().Model(&MitraBranchCoverageProvince{}).Where("mitra_branch_cov_prov_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraBranchCoverageProvince{}, err
	}
	return p, nil
}

func (p *MitraBranchCoverageProvince) FindMitraBranchCoverageProvinceByID(db *gorm.DB, pid uint64) (*MitraBranchCoverageProvinceData, error) {
	var err error
	mitrabranchcoverageprovince := MitraBranchCoverageProvinceData{}
	err = db.Debug().Model(&MitraBranchCoverageProvinceData{}).Limit(100).
		Select(`m_mitra_branch_coverage_province.mitra_branch_cov_prov_id,
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
		Joins("JOIN m_mitra on m_mitra_branch_coverage_province.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_province on m_mitra_branch_coverage_province.province_id=m_province.province_id").
		Where("mitra_branch_cov_prov_id = ?", pid).
		Find(&mitrabranchcoverageprovince).Error
	if err != nil {
		return &MitraBranchCoverageProvinceData{}, err
	}
	return &mitrabranchcoverageprovince, nil
}

func (p *MitraBranchCoverageProvince) FindMitraBranchCoverageProvinceStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraBranchCoverageProvinceData, error) {
	var err error
	mitrabranchcoverageprovince := MitraBranchCoverageProvinceData{}
	err = db.Debug().Model(&MitraBranchCoverageProvinceData{}).Limit(100).
		Select(`m_mitra_branch_coverage_province.mitra_branch_cov_prov_id,
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
		Joins("JOIN m_mitra on m_mitra_branch_coverage_province.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_province on m_mitra_branch_coverage_province.province_id=m_province.province_id").
		Where("mitra_branch_cov_prov_id = ? AND mitra_branch_cov_prov_status = ?", pid, status).
		Find(&mitrabranchcoverageprovince).Error
	if err != nil {
		return &MitraBranchCoverageProvinceData{}, err
	}
	return &mitrabranchcoverageprovince, nil
}

func (p *MitraBranchCoverageProvince) UpdateMitraBranchCoverageProvince(db *gorm.DB, pid uint64) (*MitraBranchCoverageProvince, error) {
	var err error
	db = db.Debug().Model(&MitraBranchCoverageProvince{}).Where("mitra_branch_cov_prov_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_branch_id":                  p.MitraBranchID,
			"province_id":                      p.ProvinceID,
			"mitra_branch_cov_prov_status":     p.MitraBranchCoverageProvinceStatus,
			"mitra_branch_cov_prov_updated_by": p.MitraBranchCoverageProvinceUpdatedBy,
			"mitra_branch_cov_prov_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraBranchCoverageProvince{}).Where("mitra_branch_cov_prov_id = ?", pid).Error
	if err != nil {
		return &MitraBranchCoverageProvince{}, err
	}
	return p, nil
}

func (p *MitraBranchCoverageProvince) UpdateMitraBranchCoverageProvinceStatus(db *gorm.DB, pid uint64) (*MitraBranchCoverageProvince, error) {
	var err error
	db = db.Debug().Model(&MitraBranchCoverageProvince{}).Where("mitra_branch_cov_prov_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_branch_cov_prov_status":     p.MitraBranchCoverageProvinceStatus,
			"mitra_branch_cov_prov_updated_by": p.MitraBranchCoverageProvinceUpdatedBy,
			"mitra_branch_cov_prov_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraBranchCoverageProvince{}).Where("mitra_branch_cov_prov_id = ?", pid).Error
	if err != nil {
		return &MitraBranchCoverageProvince{}, err
	}
	return p, nil
}

func (p *MitraBranchCoverageProvince) UpdateMitraBranchCoverageProvinceStatusOnly(db *gorm.DB, pid uint64) (*MitraBranchCoverageProvince, error) {
	var err error
	db = db.Debug().Model(&MitraBranchCoverageProvince{}).Where("mitra_branch_cov_prov_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_branch_cov_prov_status": p.MitraBranchCoverageProvinceStatus,
		},
	)
	err = db.Debug().Model(&MitraBranchCoverageProvince{}).Where("mitra_branch_cov_prov_id = ?", pid).Error
	if err != nil {
		return &MitraBranchCoverageProvince{}, err
	}
	return p, nil
}

func (p *MitraBranchCoverageProvince) UpdateMitraBranchCoverageProvinceStatusDelete(db *gorm.DB, pid uint64) (*MitraBranchCoverageProvince, error) {
	var err error
	db = db.Debug().Model(&MitraBranchCoverageProvince{}).Where("mitra_branch_cov_prov_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_branch_cov_prov_status":     3,
			"mitra_branch_cov_prov_deleted_by": p.MitraBranchCoverageProvinceDeletedBy,
			"mitra_branch_cov_prov_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraBranchCoverageProvince{}).Where("mitra_branch_cov_prov_id = ?", pid).Error
	if err != nil {
		return &MitraBranchCoverageProvince{}, err
	}
	return p, nil
}

func (p *MitraBranchCoverageProvince) DeleteMitraBranchCoverageProvince(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraBranchCoverageProvince{}).Where("mitra_branch_cov_prov_id = ?", pid).Take(&MitraBranchCoverageProvince{}).Delete(&MitraBranchCoverageProvince{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraBranchCoverageProvince not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

type MitraBranchPredefinedCoverageProvinceData struct {
	MitraBranchCoverageProvinceID uint64 `gorm:"column:mitra_branch_cov_prov_id" json:"mitra_branch_cov_prov_id"`
	MitraBranchID                 uint64 `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	ProvinceID                    uint64 `gorm:"column:province_id" json:"province_id"`
	ProvinceName                  string `gorm:"column:province_name" json:"province_name"`
}

type ResponseMitraBranchPredefinedCoverageProvinces struct {
	Status  int                                          `json:"status"`
	Message string                                       `json:"message"`
	Data    *[]MitraBranchPredefinedCoverageProvinceData `json:"data"`
}

func (MitraBranchPredefinedCoverageProvinceData) TableName() string {
	return "m_mitra_branch_coverage_province"
}

func (p *MitraBranchCoverageProvince) FindMitraBranchCoverageProvinceByMitraBranchID(db *gorm.DB, mitra uint64) ([]MitraBranchCoverageProvinceData, error) {
	var err error
	mitrabranchcoverageprovince := []MitraBranchCoverageProvinceData{}
	err = db.Debug().Model(&MitraBranchCoverageProvinceData{}).Limit(100).
		Select(`m_mitra_branch_coverage_province.mitra_branch_cov_prov_id,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_id,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_province.province_id,
				m_province.province_name,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_status,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_created_by,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_created_at`).
		Joins("JOIN m_mitra_branch on m_mitra_branch_coverage_province.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_province on m_mitra_branch_coverage_province.province_id=m_province.province_id").
		Where("m_mitra_branch_coverage_province.mitra_branch_id = ?", mitra).
		Find(&mitrabranchcoverageprovince).Error
	if err != nil {
		return []MitraBranchCoverageProvinceData{}, err
	}
	return mitrabranchcoverageprovince, nil
}

func (p *MitraBranchCoverageProvince) FindMitraBranchCoverageProvinceStatusByMitraBranchID(db *gorm.DB, mitra uint64, status uint64) ([]MitraBranchCoverageProvinceData, error) {
	var err error
	mitrabranchcoverageprovince := []MitraBranchCoverageProvinceData{}
	err = db.Debug().Model(&MitraBranchCoverageProvinceData{}).Limit(100).
		Select(`m_mitra_branch_coverage_province.mitra_branch_cov_prov_id,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_id,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_province.province_id,
				m_province.province_name,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_status,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_created_by,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_created_at`).
		Joins("JOIN m_mitra_branch on m_mitra_branch_coverage_province.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_province on m_mitra_branch_coverage_province.province_id=m_province.province_id").
		Where("m_mitra_branch_coverage_province.mitra_branch_id = ? AND m_mitra_branch_coverage_province.mitra_branch_cov_prov_status = ?", mitra, status).
		Find(&mitrabranchcoverageprovince).Error
	if err != nil {
		return []MitraBranchCoverageProvinceData{}, err
	}
	return mitrabranchcoverageprovince, nil
}

func (p *MitraBranchCoverageProvince) FindMitraBranchPredefinedCoverageProvinceByMitraBranchID(db *gorm.DB, mitra uint64, status uint64) ([]MitraBranchPredefinedCoverageProvinceData, error) {
	var err error
	mitrabranchcoverageprovince := []MitraBranchPredefinedCoverageProvinceData{}
	err = db.Debug().Model(&MitraBranchPredefinedCoverageProvinceData{}).Limit(100).
		Select(`m_mitra_branch_coverage_province.mitra_branch_cov_prov_id,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_id,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_province.province_id,
				m_province.province_name,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_status,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_created_by,
				m_mitra_branch_coverage_province.mitra_branch_cov_prov_created_at`).
		Joins("JOIN m_mitra_branch on m_mitra_branch_coverage_province.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Joins("JOIN m_province on m_mitra_branch_coverage_province.province_id=m_province.province_id").
		Where("m_mitra_branch_coverage_province.mitra_branch_id = ? AND m_mitra_branch_coverage_province.mitra_branch_cov_prov_status = ?", mitra, status).
		Find(&mitrabranchcoverageprovince).Error
	if err != nil {
		return []MitraBranchPredefinedCoverageProvinceData{}, err
	}
	return mitrabranchcoverageprovince, nil
}
