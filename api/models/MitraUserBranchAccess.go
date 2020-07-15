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

type MitraUserBranchAccess struct {
	MitraUserBranchAccessID        *uint64    `gorm:"column:mitra_user_branch_accs_id;primary_key;" json:"mitra_user_branch_accs_id"`
	MitraUserID                    *uint64    `gorm:"column:mitra_user_id" json:"mitra_user_id"`
	MitraBranchID                  uint64     `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	MitraUserBranchAccessStatus    uint64     `gorm:"column:mitra_user_branch_accs_status;size:2" json:"mitra_user_branch_accs_status"`
	MitraUserBranchAccessCreatedBy string     `gorm:"column:mitra_user_branch_accs_created_by;size:125" json:"mitra_user_branch_accs_created_by"`
	MitraUserBranchAccessCreatedAt time.Time  `gorm:"column:mitra_user_branch_accs_created_at;default:CURRENT_TIMESTAMP" json:"mitra_user_branch_accs_created_at"`
	MitraUserBranchAccessUpdatedBy string     `gorm:"column:mitra_user_branch_accs_updated_by;size:125" json:"mitra_user_branch_accs_updated_by"`
	MitraUserBranchAccessUpdatedAt *time.Time `gorm:"column:mitra_user_branch_accs_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_user_branch_accs_created_at"`
	MitraUserBranchAccessDeletedBy string     `gorm:"column:mitra_user_branch_accs_deleted_by;size:125" json:"mitra_user_branch_accs_deleted_by"`
	MitraUserBranchAccessDeletedAt *time.Time `gorm:"column:mitra_user_branch_accs_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_user_branch_accs_deleted_at"`
}

type MitraUserBranchAccessData struct {
	MitraUserBranchAccessID     uint64 `gorm:"column:mitra_user_branch_accs_id" json:"mitra_user_branch_accs_id"`
	MitraUserID                 uint64 `gorm:"column:mitra_user_id" json:"mitra_user_id"`
	MitraUserName               string `gorm:"column:mitra_username" json:"mitra_username"`
	MitraBranchID               uint64 `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	MitraBranchName             string `gorm:"column:mitra_branch_name" json:"mitra_branch_name"`
	MitraUserBranchAccessStatus uint64 `gorm:"column:mitra_user_branch_accs_status" json:"mitra_user_branch_accs_status"`
	// MitraUserBranchAccessCreatedBy string     `gorm:"column:mitra_user_branch_accs_created_by" json:"mitra_user_branch_accs_created_by"`
	// MitraUserBranchAccessCreatedAt time.Time  `gorm:"column:mitra_user_branch_accs_created_at" json:"mitra_user_branch_accs_created_at"`
	// MitraUserBranchAccessUpdatedBy string     `gorm:"column:mitra_user_branch_accs_updated_by" json:"mitra_user_branch_accs_updated_by"`
	// MitraUserBranchAccessUpdatedAt *time.Time `gorm:"column:mitra_user_branch_accs_updated_at" json:"mitra_user_branch_accs_updated_at"`
	// MitraUserBranchAccessDeletedBy string     `gorm:"column:mitra_user_branch_accs_deleted_by" json:"mitra_user_branch_accs_deleted_by"`
	// MitraUserBranchAccessDeletedAt *time.Time `gorm:"column:mitra_user_branch_accs_deleted_at" json:"mitra_user_branch_accs_deleted_at"`
}

type ResponseMitraUserBranchAccesss struct {
	Status  int                          `json:"status"`
	Message string                       `json:"message"`
	Data    *[]MitraUserBranchAccessData `json:"data"`
}

type ResponseMitraUserBranchAccess struct {
	Status  int                        `json:"status"`
	Message string                     `json:"message"`
	Data    *MitraUserBranchAccessData `json:"data"`
}

type ResponseMitraUserBranchAccessCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraUserBranchAccess) TableName() string {
	return "m_mitra_user_branch_access"
}

func (MitraUserBranchAccessData) TableName() string {
	return "m_mitra_user_branch_access"
}

func (p *MitraUserBranchAccess) Prepare() {
	p.MitraUserBranchAccessID = nil
	p.MitraUserID = p.MitraUserID
	p.MitraBranchID = p.MitraBranchID
	p.MitraUserBranchAccessStatus = p.MitraUserBranchAccessStatus
	p.MitraUserBranchAccessCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraUserBranchAccessCreatedBy))
	p.MitraUserBranchAccessCreatedAt = time.Now()
	p.MitraUserBranchAccessUpdatedBy = html.EscapeString(strings.TrimSpace(p.MitraUserBranchAccessUpdatedBy))
	p.MitraUserBranchAccessUpdatedAt = p.MitraUserBranchAccessUpdatedAt
	p.MitraUserBranchAccessDeletedBy = html.EscapeString(strings.TrimSpace(p.MitraUserBranchAccessDeletedBy))
	p.MitraUserBranchAccessDeletedAt = p.MitraUserBranchAccessDeletedAt
}

func (p *MitraUserBranchAccess) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.MitraBranchID == 0 {
			return errors.New("Required MitraBranch")
		}
		return nil
	}
}

func (p *MitraUserBranchAccess) SaveMitraUserBranchAccess(db *gorm.DB) (*MitraUserBranchAccess, error) {
	var err error
	err = db.Debug().Model(&MitraUserBranchAccess{}).Create(&p).Error
	if err != nil {
		return &MitraUserBranchAccess{}, err
	}
	return p, nil
}

func (p *MitraUserBranchAccess) FindAllMitraUserBranchAccess(db *gorm.DB) (*[]MitraUserBranchAccessData, error) {
	var err error
	mitrauserbranchaccess := []MitraUserBranchAccessData{}
	err = db.Debug().Model(&MitraUserBranchAccessData{}).Limit(100).
		Select(`m_mitra_user_branch_access.mitra_user_branch_accs_id,
				m_mitra_user.mitra_user_id,
				m_mitra_user.mitra_username,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_user_branch_access.mitra_user_branch_accs_status,
				m_mitra_user_branch_access.mitra_user_branch_accs_created_by,
				m_mitra_user_branch_access.mitra_user_branch_accs_created_at,
				m_mitra_user_branch_access.mitra_user_branch_accs_updated_by,
				m_mitra_user_branch_access.mitra_user_branch_accs_updated_at,
				m_mitra_user_branch_access.mitra_user_branch_accs_deleted_by,
				m_mitra_user_branch_access.mitra_user_branch_accs_deleted_at`).
		Joins("JOIN m_mitra_user on m_mitra_user_branch_access.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra_branch on m_mitra_user_branch_access.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Find(&mitrauserbranchaccess).Error
	if err != nil {
		return &[]MitraUserBranchAccessData{}, err
	}
	return &mitrauserbranchaccess, nil
}

func (p *MitraUserBranchAccess) FindAllMitraUserBranchAccessStatus(db *gorm.DB, status uint64) (*[]MitraUserBranchAccessData, error) {
	var err error
	mitrauserbranchaccess := []MitraUserBranchAccessData{}
	err = db.Debug().Model(&MitraUserBranchAccessData{}).Limit(100).
		Select(`m_mitra_user_branch_access.mitra_user_branch_accs_id,
				m_mitra_user.mitra_user_id,
				m_mitra_user.mitra_username,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_user_branch_access.mitra_user_branch_accs_status,
				m_mitra_user_branch_access.mitra_user_branch_accs_created_by,
				m_mitra_user_branch_access.mitra_user_branch_accs_created_at,
				m_mitra_user_branch_access.mitra_user_branch_accs_updated_by,
				m_mitra_user_branch_access.mitra_user_branch_accs_updated_at,
				m_mitra_user_branch_access.mitra_user_branch_accs_deleted_by,
				m_mitra_user_branch_access.mitra_user_branch_accs_deleted_at`).
		Joins("JOIN m_mitra_user on m_mitra_user_branch_access.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra_branch on m_mitra_user_branch_access.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Where("mitra_user_branch_accs_status = ?", status).
		Find(&mitrauserbranchaccess).Error
	if err != nil {
		return &[]MitraUserBranchAccessData{}, err
	}
	return &mitrauserbranchaccess, nil
}

func (p *MitraUserBranchAccess) FindMitraUserBranchAccessDataByID(db *gorm.DB, pid uint64) (*MitraUserBranchAccess, error) {
	var err error
	err = db.Debug().Model(&MitraUserBranchAccess{}).Where("mitra_user_branch_accs_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraUserBranchAccess{}, err
	}
	return p, nil
}

func (p *MitraUserBranchAccess) FindMitraUserBranchAccessByID(db *gorm.DB, pid uint64) (*MitraUserBranchAccessData, error) {
	var err error
	mitrauserbranchaccess := MitraUserBranchAccessData{}
	err = db.Debug().Model(&MitraUserBranchAccessData{}).Limit(100).
		Select(`m_mitra_user_branch_access.mitra_user_branch_accs_id,
				m_mitra_user.mitra_user_id,
				m_mitra_user.mitra_username,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_user_branch_access.mitra_user_branch_accs_status,
				m_mitra_user_branch_access.mitra_user_branch_accs_created_by,
				m_mitra_user_branch_access.mitra_user_branch_accs_created_at,
				m_mitra_user_branch_access.mitra_user_branch_accs_updated_by,
				m_mitra_user_branch_access.mitra_user_branch_accs_updated_at,
				m_mitra_user_branch_access.mitra_user_branch_accs_deleted_by,
				m_mitra_user_branch_access.mitra_user_branch_accs_deleted_at`).
		Joins("JOIN m_mitra_user on m_mitra_user_branch_access.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra_branch on m_mitra_user_branch_access.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Where("mitra_user_branch_accs_id = ?", pid).
		Find(&mitrauserbranchaccess).Error
	if err != nil {
		return &MitraUserBranchAccessData{}, err
	}
	return &mitrauserbranchaccess, nil
}

func (p *MitraUserBranchAccess) FindMitraUserBranchAccessStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraUserBranchAccessData, error) {
	var err error
	mitrauserbranchaccess := MitraUserBranchAccessData{}
	err = db.Debug().Model(&MitraUserBranchAccessData{}).Limit(100).
		Select(`m_mitra_user_branch_access.mitra_user_branch_accs_id,
				m_mitra_user.mitra_user_id,
				m_mitra_user.mitra_username,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_user_branch_access.mitra_user_branch_accs_status,
				m_mitra_user_branch_access.mitra_user_branch_accs_created_by,
				m_mitra_user_branch_access.mitra_user_branch_accs_created_at,
				m_mitra_user_branch_access.mitra_user_branch_accs_updated_by,
				m_mitra_user_branch_access.mitra_user_branch_accs_updated_at,
				m_mitra_user_branch_access.mitra_user_branch_accs_deleted_by,
				m_mitra_user_branch_access.mitra_user_branch_accs_deleted_at`).
		Joins("JOIN m_mitra_user on m_mitra_user_branch_access.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra_branch on m_mitra_user_branch_access.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Where("mitra_user_branch_accs_id = ? AND mitra_user_branch_accs_status = ?", pid, status).
		Find(&mitrauserbranchaccess).Error
	if err != nil {
		return &MitraUserBranchAccessData{}, err
	}
	return &mitrauserbranchaccess, nil
}

func (p *MitraUserBranchAccess) UpdateMitraUserBranchAccess(db *gorm.DB, pid uint64) (*MitraUserBranchAccess, error) {
	var err error
	db = db.Debug().Model(&MitraUserBranchAccess{}).Where("mitra_user_branch_accs_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_user_id":                     p.MitraUserID,
			"mitra_branch_id":                   p.MitraBranchID,
			"mitra_user_branch_accs_status":     p.MitraUserBranchAccessStatus,
			"mitra_user_branch_accs_updated_by": p.MitraUserBranchAccessUpdatedBy,
			"mitra_user_branch_accs_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraUserBranchAccess{}).Where("mitra_user_branch_accs_id = ?", pid).Error
	if err != nil {
		return &MitraUserBranchAccess{}, err
	}
	return p, nil
}

func (p *MitraUserBranchAccess) UpdateMitraUserBranchAccessStatus(db *gorm.DB, pid uint64) (*MitraUserBranchAccess, error) {
	var err error
	db = db.Debug().Model(&MitraUserBranchAccess{}).Where("mitra_user_branch_accs_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_user_branch_accs_status":     p.MitraUserBranchAccessStatus,
			"mitra_user_branch_accs_updated_by": p.MitraUserBranchAccessUpdatedBy,
			"mitra_user_branch_accs_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraUserBranchAccess{}).Where("mitra_user_branch_accs_id = ?", pid).Error
	if err != nil {
		return &MitraUserBranchAccess{}, err
	}
	return p, nil
}

func (p *MitraUserBranchAccess) UpdateMitraUserBranchAccessStatusOnly(db *gorm.DB, pid uint64) (*MitraUserBranchAccess, error) {
	var err error
	db = db.Debug().Model(&MitraUserBranchAccess{}).Where("mitra_user_branch_accs_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_user_branch_accs_status": p.MitraUserBranchAccessStatus,
		},
	)
	err = db.Debug().Model(&MitraUserBranchAccess{}).Where("mitra_user_branch_accs_id = ?", pid).Error
	if err != nil {
		return &MitraUserBranchAccess{}, err
	}
	return p, nil
}

func (p *MitraUserBranchAccess) UpdateMitraUserBranchAccessStatusDelete(db *gorm.DB, pid uint64) (*MitraUserBranchAccess, error) {
	var err error
	db = db.Debug().Model(&MitraUserBranchAccess{}).Where("mitra_user_branch_accs_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_user_branch_accs_status":     3,
			"mitra_user_branch_accs_deleted_by": p.MitraUserBranchAccessDeletedBy,
			"mitra_user_branch_accs_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraUserBranchAccess{}).Where("mitra_user_branch_accs_id = ?", pid).Error
	if err != nil {
		return &MitraUserBranchAccess{}, err
	}
	return p, nil
}

func (p *MitraUserBranchAccess) DeleteMitraUserBranchAccess(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraUserBranchAccess{}).Where("mitra_user_branch_accs_id = ?", pid).Take(&MitraUserBranchAccess{}).Delete(&MitraUserBranchAccess{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraUserBranchAccess not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

type MitarUserBranchAccessData struct {
	MitraUserBranchAccessID uint64 `gorm:"column:mitra_user_branch_accs_id" json:"mitra_user_branch_accs_id"`
	MitraUserID             uint64 `gorm:"column:mitra_user_id" json:"mitra_user_id"`
	MitraBranchID           uint64 `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	MitraBranchName         string `gorm:"column:mitra_branch_name" json:"mitra_branch_name"`
}

type ResponsePredefinedMitraUserBranchAccesss struct {
	Status  int                          `json:"status"`
	Message string                       `json:"message"`
	Data    *[]MitarUserBranchAccessData `json:"data"`
}

func (MitarUserBranchAccessData) TableName() string {
	return "m_mitra_user_branch_access"
}

func (p *MitraUserBranchAccess) FindMitraUserBranchAccessByMitraUserID(db *gorm.DB, pid uint64) ([]MitraUserBranchAccessData, error) {
	var err error
	mitraaocoverageprovince := []MitraUserBranchAccessData{}
	err = db.Debug().Model(&MitraUserBranchAccessData{}).Limit(100).
		Select(`m_mitra_user_branch_access.mitra_user_branch_accs_id,
				m_mitra_user_branch_access.mitra_user_branch_accs_id,
				m_mitra_user.mitra_user_id,
				m_mitra_user.mitra_username,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_user_branch_access.mitra_user_branch_accs_status,
				m_mitra_user_branch_access.mitra_user_branch_accs_created_by,
				m_mitra_user_branch_access.mitra_user_branch_accs_created_at`).
		Joins("JOIN m_mitra_user on m_mitra_user_branch_access.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra_branch on m_mitra_user_branch_access.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Where("m_mitra_user_branch_access.mitra_user_id = ?", pid).
		Find(&mitraaocoverageprovince).Error
	if err != nil {
		return []MitraUserBranchAccessData{}, err
	}
	return mitraaocoverageprovince, nil
}

func (p *MitraUserBranchAccess) FindMitraUserBranchAccessStatusByMitraUserID(db *gorm.DB, mitra uint64, status uint64) ([]MitraUserBranchAccessData, error) {
	var err error
	mitraaocoverageprovince := []MitraUserBranchAccessData{}
	err = db.Debug().Model(&MitraUserBranchAccessData{}).Limit(100).
		Select(`m_mitra_user_branch_access.mitra_user_branch_accs_id,
				m_mitra_user_branch_access.mitra_user_branch_accs_id,
				m_mitra_user.mitra_user_id,
				m_mitra_user.mitra_username,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_user_branch_access.mitra_user_branch_accs_status,
				m_mitra_user_branch_access.mitra_user_branch_accs_created_by,
				m_mitra_user_branch_access.mitra_user_branch_accs_created_at`).
		Joins("JOIN m_mitra_user on m_mitra_user_branch_access.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra_branch on m_mitra_user_branch_access.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Where("m_mitra_user_branch_access.mitra_user_id = ? AND m_mitra_user_branch_access.mitra_user_branch_accs_status = ?", mitra, status).
		Find(&mitraaocoverageprovince).Error
	if err != nil {
		return []MitraUserBranchAccessData{}, err
	}
	return mitraaocoverageprovince, nil
}

func (p *MitraUserBranchAccess) FindMitraAOPredefinedCoverageMitraBranchByMitraUserID(db *gorm.DB, mitra uint64, status uint64) ([]MitarUserBranchAccessData, error) {
	var err error
	mitraaocoverageprovince := []MitarUserBranchAccessData{}
	err = db.Debug().Model(&MitarUserBranchAccessData{}).Limit(100).
		Select(`m_mitra_user_branch_access.mitra_user_branch_accs_id,
				m_mitra_user_branch_access.mitra_user_branch_accs_id,
				m_mitra_user.mitra_user_id,
				m_mitra_user.mitra_username,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_user_branch_access.mitra_user_branch_accs_status,
				m_mitra_user_branch_access.mitra_user_branch_accs_created_by,
				m_mitra_user_branch_access.mitra_user_branch_accs_created_at`).
		Joins("JOIN m_mitra_user on m_mitra_user_branch_access.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra_branch on m_mitra_user_branch_access.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Where("m_mitra_user_branch_access.mitra_user_id = ? AND m_mitra_user_branch_access.mitra_user_branch_accs_status = ?", mitra, status).
		Find(&mitraaocoverageprovince).Error
	if err != nil {
		return []MitarUserBranchAccessData{}, err
	}
	return mitraaocoverageprovince, nil
}

func (p *MitraUserBranchAccess) FindMitraAOPredefinedCoverageMitraBranchByMitraUserIDByProvince(db *gorm.DB, mitra uint64, province uint64, status uint64) ([]MitarUserBranchAccessData, error) {
	var err error
	mitraaocoverageprovince := []MitarUserBranchAccessData{}
	err = db.Debug().Model(&MitarUserBranchAccessData{}).Limit(100).
		Select(`m_mitra_user_branch_access.mitra_user_branch_accs_id,
				m_mitra_user_branch_access.mitra_user_branch_accs_id,
				m_mitra_user.mitra_user_id,
				m_mitra_user.mitra_username,
				m_mitra_branch.mitra_branch_id,
				m_mitra_branch.mitra_branch_name,
				m_mitra_user_branch_access.mitra_user_branch_accs_status,
				m_mitra_user_branch_access.mitra_user_branch_accs_created_by,
				m_mitra_user_branch_access.mitra_user_branch_accs_created_at`).
		Joins("JOIN m_mitra_user on m_mitra_user_branch_access.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra_branch on m_mitra_user_branch_access.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Where("m_mitra_user_branch_access.mitra_user_id = ? AND m_mitra_branch.province_id = ? AND m_mitra_user_branch_access.mitra_user_branch_accs_status = ?", mitra, province, status).
		Find(&mitraaocoverageprovince).Error
	if err != nil {
		return []MitarUserBranchAccessData{}, err
	}
	return mitraaocoverageprovince, nil
}
