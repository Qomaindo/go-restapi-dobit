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

type MitraCoverageProvince struct {
	MitraCoverageProvinceID        *uint64    `gorm:"column:mitra_cov_prov_id;primary_key;" json:"mitra_cov_prov_id"`
	MitraID                        uint64     `gorm:"column:mitra_id" json:"mitra_id"`
	ProvinceID                     uint64     `gorm:"column:province_id" json:"province_id"`
	MitraCoverageProvinceStatus    uint64     `gorm:"column:mitra_cov_prov_status;size:2" json:"mitra_cov_prov_status"`
	MitraCoverageProvinceCreatedBy string     `gorm:"column:mitra_cov_prov_created_by;size:125" json:"mitra_cov_prov_created_by"`
	MitraCoverageProvinceCreatedAt time.Time  `gorm:"column:mitra_cov_prov_created_at;default:CURRENT_TIMESTAMP" json:"mitra_cov_prov_created_at"`
	MitraCoverageProvinceUpdatedBy string     `gorm:"column:mitra_cov_prov_updated_by;size:125" json:"mitra_cov_prov_updated_by"`
	MitraCoverageProvinceUpdatedAt *time.Time `gorm:"column:mitra_cov_prov_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_cov_prov_created_at"`
	MitraCoverageProvinceDeletedBy string     `gorm:"column:mitra_cov_prov_deleted_by;size:125" json:"mitra_cov_prov_deleted_by"`
	MitraCoverageProvinceDeletedAt *time.Time `gorm:"column:mitra_cov_prov_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_cov_prov_deleted_at"`
}

type MitraCoverageProvinceData struct {
	MitraCoverageProvinceID        uint64     `gorm:"column:mitra_cov_prov_id" json:"mitra_cov_prov_id"`
	MitraID                        uint64     `gorm:"column:mitra_id" json:"mitra_id"`
	MitraName                      string     `gorm:"column:mitra_name" json:"mitra_name"`
	ProvinceID                     uint64     `gorm:"column:province_id" json:"province_id"`
	ProvinceName                   string     `gorm:"column:province_name" json:"province_name"`
	MitraCoverageProvinceStatus    uint64     `gorm:"column:mitra_cov_prov_status" json:"mitra_cov_prov_status"`
	MitraCoverageProvinceCreatedBy string     `gorm:"column:mitra_cov_prov_created_by" json:"mitra_cov_prov_created_by"`
	MitraCoverageProvinceCreatedAt time.Time  `gorm:"column:mitra_cov_prov_created_at" json:"mitra_cov_prov_created_at"`
	MitraCoverageProvinceUpdatedBy string     `gorm:"column:mitra_cov_prov_updated_by" json:"mitra_cov_prov_updated_by"`
	MitraCoverageProvinceUpdatedAt *time.Time `gorm:"column:mitra_cov_prov_updated_at" json:"mitra_cov_prov_updated_at"`
	MitraCoverageProvinceDeletedBy string     `gorm:"column:mitra_cov_prov_deleted_by" json:"mitra_cov_prov_deleted_by"`
	MitraCoverageProvinceDeletedAt *time.Time `gorm:"column:mitra_cov_prov_deleted_at" json:"mitra_cov_prov_deleted_at"`
}

type ResponseMitraCoverageProvinces struct {
	Status  int                          `json:"status"`
	Message string                       `json:"message"`
	Data    *[]MitraCoverageProvinceData `json:"data"`
}

type ResponseMitraCoverageProvince struct {
	Status  int                        `json:"status"`
	Message string                     `json:"message"`
	Data    *MitraCoverageProvinceData `json:"data"`
}

type ResponseMitraCoverageProvinceCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraCoverageProvince) TableName() string {
	return "m_mitra_coverage_province"
}

func (MitraCoverageProvinceData) TableName() string {
	return "m_mitra_coverage_province"
}

func (p *MitraCoverageProvince) Prepare() {
	p.MitraCoverageProvinceID = nil
	p.MitraID = p.MitraID
	p.ProvinceID = p.ProvinceID
	p.MitraCoverageProvinceStatus = p.MitraCoverageProvinceStatus
	p.MitraCoverageProvinceCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraCoverageProvinceCreatedBy))
	p.MitraCoverageProvinceCreatedAt = time.Now()
	p.MitraCoverageProvinceUpdatedBy = html.EscapeString(strings.TrimSpace(p.MitraCoverageProvinceUpdatedBy))
	p.MitraCoverageProvinceUpdatedAt = p.MitraCoverageProvinceUpdatedAt
	p.MitraCoverageProvinceDeletedBy = html.EscapeString(strings.TrimSpace(p.MitraCoverageProvinceDeletedBy))
	p.MitraCoverageProvinceDeletedAt = p.MitraCoverageProvinceDeletedAt
}

func (p *MitraCoverageProvince) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.MitraID == 0 {
			return errors.New("Required Country")
		}
		if p.ProvinceID == 0 {
			return errors.New("Required Province")
		}
		return nil
	}
}

func (p *MitraCoverageProvince) SaveMitraCoverageProvince(db *gorm.DB) (*MitraCoverageProvince, error) {
	var err error
	err = db.Debug().Model(&MitraCoverageProvince{}).Create(&p).Error
	if err != nil {
		return &MitraCoverageProvince{}, err
	}
	return p, nil
}

func (p *MitraCoverageProvince) FindAllMitraCoverageProvince(db *gorm.DB) (*[]MitraCoverageProvinceData, error) {
	var err error
	mitracoverageprovince := []MitraCoverageProvinceData{}
	err = db.Debug().Model(&MitraCoverageProvinceData{}).Limit(100).
		Select(`m_mitra_coverage_province.mitra_cov_prov_id,
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
		Joins("JOIN m_mitra on m_mitra_coverage_province.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_province on m_mitra_coverage_province.province_id=m_province.province_id").
		Find(&mitracoverageprovince).Error
	if err != nil {
		return &[]MitraCoverageProvinceData{}, err
	}
	return &mitracoverageprovince, nil
}

func (p *MitraCoverageProvince) FindAllMitraCoverageProvinceStatus(db *gorm.DB, status uint64) (*[]MitraCoverageProvinceData, error) {
	var err error
	mitracoverageprovince := []MitraCoverageProvinceData{}
	err = db.Debug().Model(&MitraCoverageProvinceData{}).Limit(100).
		Select(`m_mitra_coverage_province.mitra_cov_prov_id,
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
		Joins("JOIN m_mitra on m_mitra_coverage_province.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_province on m_mitra_coverage_province.province_id=m_province.province_id").
		Where("mitra_cov_prov_status = ?", status).
		Find(&mitracoverageprovince).Error
	if err != nil {
		return &[]MitraCoverageProvinceData{}, err
	}
	return &mitracoverageprovince, nil
}

func (p *MitraCoverageProvince) FindMitraCoverageProvinceDataByID(db *gorm.DB, pid uint64) (*MitraCoverageProvince, error) {
	var err error
	err = db.Debug().Model(&MitraCoverageProvince{}).Where("mitra_cov_prov_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraCoverageProvince{}, err
	}
	return p, nil
}

func (p *MitraCoverageProvince) FindMitraCoverageProvinceByID(db *gorm.DB, pid uint64) (*MitraCoverageProvinceData, error) {
	var err error
	mitracoverageprovince := MitraCoverageProvinceData{}
	err = db.Debug().Model(&MitraCoverageProvinceData{}).Limit(100).
		Select(`m_mitra_coverage_province.mitra_cov_prov_id,
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
		Joins("JOIN m_mitra on m_mitra_coverage_province.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_province on m_mitra_coverage_province.province_id=m_province.province_id").
		Where("mitra_cov_prov_id = ?", pid).
		Find(&mitracoverageprovince).Error
	if err != nil {
		return &MitraCoverageProvinceData{}, err
	}
	return &mitracoverageprovince, nil
}

func (p *MitraCoverageProvince) FindMitraCoverageProvinceStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraCoverageProvinceData, error) {
	var err error
	mitracoverageprovince := MitraCoverageProvinceData{}
	err = db.Debug().Model(&MitraCoverageProvinceData{}).Limit(100).
		Select(`m_mitra_coverage_province.mitra_cov_prov_id,
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
		Joins("JOIN m_mitra on m_mitra_coverage_province.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_province on m_mitra_coverage_province.province_id=m_province.province_id").
		Where("mitra_cov_prov_id = ? AND mitra_cov_prov_status = ?", pid, status).
		Find(&mitracoverageprovince).Error
	if err != nil {
		return &MitraCoverageProvinceData{}, err
	}
	return &mitracoverageprovince, nil
}

func (p *MitraCoverageProvince) UpdateMitraCoverageProvince(db *gorm.DB, pid uint64) (*MitraCoverageProvince, error) {
	var err error
	db = db.Debug().Model(&MitraCoverageProvince{}).Where("mitra_cov_prov_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_id":                  p.MitraID,
			"province_id":               p.ProvinceID,
			"mitra_cov_prov_status":     p.MitraCoverageProvinceStatus,
			"mitra_cov_prov_updated_by": p.MitraCoverageProvinceUpdatedBy,
			"mitra_cov_prov_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraCoverageProvince{}).Where("mitra_cov_prov_id = ?", pid).Error
	if err != nil {
		return &MitraCoverageProvince{}, err
	}
	return p, nil
}

func (p *MitraCoverageProvince) UpdateMitraCoverageProvinceStatus(db *gorm.DB, pid uint64) (*MitraCoverageProvince, error) {
	var err error
	db = db.Debug().Model(&MitraCoverageProvince{}).Where("mitra_cov_prov_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_cov_prov_status":     p.MitraCoverageProvinceStatus,
			"mitra_cov_prov_updated_by": p.MitraCoverageProvinceUpdatedBy,
			"mitra_cov_prov_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraCoverageProvince{}).Where("mitra_cov_prov_id = ?", pid).Error
	if err != nil {
		return &MitraCoverageProvince{}, err
	}
	return p, nil
}

func (p *MitraCoverageProvince) UpdateMitraCoverageProvinceStatusOnly(db *gorm.DB, pid uint64) (*MitraCoverageProvince, error) {
	var err error
	db = db.Debug().Model(&MitraCoverageProvince{}).Where("mitra_cov_prov_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_cov_prov_status": p.MitraCoverageProvinceStatus,
		},
	)
	err = db.Debug().Model(&MitraCoverageProvince{}).Where("mitra_cov_prov_id = ?", pid).Error
	if err != nil {
		return &MitraCoverageProvince{}, err
	}
	return p, nil
}

func (p *MitraCoverageProvince) UpdateMitraCoverageProvinceStatusDelete(db *gorm.DB, pid uint64) (*MitraCoverageProvince, error) {
	var err error
	db = db.Debug().Model(&MitraCoverageProvince{}).Where("mitra_cov_prov_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_cov_prov_status":     3,
			"mitra_cov_prov_deleted_by": p.MitraCoverageProvinceDeletedBy,
			"mitra_cov_prov_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraCoverageProvince{}).Where("mitra_cov_prov_id = ?", pid).Error
	if err != nil {
		return &MitraCoverageProvince{}, err
	}
	return p, nil
}

func (p *MitraCoverageProvince) DeleteMitraCoverageProvince(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraCoverageProvince{}).Where("mitra_cov_prov_id = ?", pid).Take(&MitraCoverageProvince{}).Delete(&MitraCoverageProvince{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraCoverageProvince not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

type MitraPredefinedCoverageProvinceData struct {
	MitraCoverageProvinceID uint64 `gorm:"column:mitra_cov_prov_id" json:"mitra_cov_prov_id"`
	MitraID                 uint64 `gorm:"column:mitra_id" json:"mitra_id"`
	ProvinceID              uint64 `gorm:"column:province_id" json:"province_id"`
	ProvinceName            string `gorm:"column:province_name" json:"province_name"`
}

type ResponseMitraPredefinedCoverageProvinces struct {
	Status  int                                    `json:"status"`
	Message string                                 `json:"message"`
	Data    *[]MitraPredefinedCoverageProvinceData `json:"data"`
}

func (MitraPredefinedCoverageProvinceData) TableName() string {
	return "m_mitra_coverage_province"
}

func (p *MitraCoverageProvince) FindMitraCoverageProvinceByMitraID(db *gorm.DB, mitra uint64) ([]MitraCoverageProvinceData, error) {
	var err error
	mitracoverageprovince := []MitraCoverageProvinceData{}
	err = db.Debug().Model(&MitraCoverageProvinceData{}).Limit(100).
		Select(`m_mitra_coverage_province.mitra_cov_prov_id,
				m_mitra_coverage_province.mitra_cov_prov_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_province.province_id,
				m_province.province_name,
				m_mitra_coverage_province.mitra_cov_prov_status,
				m_mitra_coverage_province.mitra_cov_prov_created_by,
				m_mitra_coverage_province.mitra_cov_prov_created_at`).
		Joins("JOIN m_mitra on m_mitra_coverage_province.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_province on m_mitra_coverage_province.province_id=m_province.province_id").
		Where("m_mitra_coverage_province.mitra_id = ?", mitra).
		Find(&mitracoverageprovince).Error
	if err != nil {
		return []MitraCoverageProvinceData{}, err
	}
	return mitracoverageprovince, nil
}

func (p *MitraCoverageProvince) FindMitraCoverageProvinceStatusByMitraID(db *gorm.DB, mitra uint64, status uint64) ([]MitraCoverageProvinceData, error) {
	var err error
	mitracoverageprovince := []MitraCoverageProvinceData{}
	err = db.Debug().Model(&MitraCoverageProvinceData{}).Limit(100).
		Select(`m_mitra_coverage_province.mitra_cov_prov_id,
				m_mitra_coverage_province.mitra_cov_prov_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_province.province_id,
				m_province.province_name,
				m_mitra_coverage_province.mitra_cov_prov_status,
				m_mitra_coverage_province.mitra_cov_prov_created_by,
				m_mitra_coverage_province.mitra_cov_prov_created_at`).
		Joins("JOIN m_mitra on m_mitra_coverage_province.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_province on m_mitra_coverage_province.province_id=m_province.province_id").
		Where("m_mitra_coverage_province.mitra_id = ? AND m_mitra_coverage_province.mitra_cov_prov_status = ?", mitra, status).
		Find(&mitracoverageprovince).Error
	if err != nil {
		return []MitraCoverageProvinceData{}, err
	}
	return mitracoverageprovince, nil
}

func (p *MitraCoverageProvince) FindMitraPredefinedCoverageProvinceByMitraID(db *gorm.DB, mitra uint64, status uint64) ([]MitraPredefinedCoverageProvinceData, error) {
	var err error
	mitracoverageprovince := []MitraPredefinedCoverageProvinceData{}
	err = db.Debug().Model(&MitraPredefinedCoverageProvinceData{}).Limit(100).
		Select(`m_mitra_coverage_province.mitra_cov_prov_id,
				m_mitra_coverage_province.mitra_cov_prov_id,
				m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_province.province_id,
				m_province.province_name,
				m_mitra_coverage_province.mitra_cov_prov_status,
				m_mitra_coverage_province.mitra_cov_prov_created_by,
				m_mitra_coverage_province.mitra_cov_prov_created_at`).
		Joins("JOIN m_mitra on m_mitra_coverage_province.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_province on m_mitra_coverage_province.province_id=m_province.province_id").
		Where("m_mitra_coverage_province.mitra_id = ? AND m_mitra_coverage_province.mitra_cov_prov_status = ?", mitra, status).
		Find(&mitracoverageprovince).Error
	if err != nil {
		return []MitraPredefinedCoverageProvinceData{}, err
	}
	return mitracoverageprovince, nil
}
