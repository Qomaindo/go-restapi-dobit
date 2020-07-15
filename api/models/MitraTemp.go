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

type MitraTemp struct {
	MitraTempID           uint64    `gorm:"column:mitra_temp_id;primary_key;" json:"mitra_temp_id"`
	MitraID               uint64    `gorm:"column:mitra_id" json:"mitra_id"`
	MitraTypeTempID       uint64    `gorm:"column:mitra_type_temp_id" json:"mitra_type_temp_id"`
	MitraTempCode         string    `gorm:"column:mitra_temp_code" json:"mitra_temp_code"`
	MitraTempName         string    `gorm:"column:mitra_temp_name" json:"mitra_temp_name"`
	MitraTempWebsite      string    `gorm:"column:mitra_temp_website" json:"mitra_temp_website"`
	MitraTempEmail        string    `gorm:"column:mitra_temp_email" json:"mitra_temp_email"`
	MitraTempLogo         string    `gorm:"column:mitra_temp_logo" json:"mitra_temp_logo"`
	MitraTempPKSFile      string    `gorm:"column:mitra_temp_pks_file" json:"mitra_temp_pks_file"`
	MitraTempReason       string    `gorm:"column:mitra_temp_reason" json:"mitra_temp_reason"`
	MitraTempStatus       uint64    `gorm:"column:mitra_temp_status;size:2" json:"mitra_temp_status"`
	MitraTempActionBefore uint64    `gorm:"column:mitra_temp_action_before;size:2" json:"mitra_temp_action_before"`
	MitraTempCreatedBy    string    `gorm:"column:mitra_temp_created_by;size:125" json:"mitra_temp_created_by"`
	MitraTempCreatedAt    time.Time `gorm:"column:mitra_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_temp_created_at"`
}

type MitraTempPend struct {
	MitraTempID           uint64    `gorm:"column:mitra_temp_id;primary_key;" json:"mitra_temp_id"`
	MitraID               *uint64   `gorm:"column:mitra_id" json:"mitra_id"`
	MitraTypeTempID       uint64    `gorm:"column:mitra_type_temp_id" json:"mitra_type_temp_id"`
	MitraTempCode         string    `gorm:"column:mitra_temp_code" json:"mitra_temp_code"`
	MitraTempName         string    `gorm:"column:mitra_temp_name" json:"mitra_temp_name"`
	MitraTempWebsite      string    `gorm:"column:mitra_temp_website" json:"mitra_temp_website"`
	MitraTempEmail        string    `gorm:"column:mitra_temp_email" json:"mitra_temp_email"`
	MitraTempLogo         string    `gorm:"column:mitra_temp_logo" json:"mitra_temp_logo"`
	MitraTempPKSFile      string    `gorm:"column:mitra_temp_pks_file" json:"mitra_temp_pks_file"`
	MitraTempReason       string    `gorm:"column:mitra_temp_reason" json:"mitra_temp_reason"`
	MitraTempStatus       uint64    `gorm:"column:mitra_temp_status;size:2" json:"mitra_temp_status"`
	MitraTempActionBefore uint64    `gorm:"column:mitra_temp_action_before;size:2" json:"mitra_temp_action_before"`
	MitraTempCreatedBy    string    `gorm:"column:mitra_temp_created_by;size:125" json:"mitra_temp_created_by"`
	MitraTempCreatedAt    time.Time `gorm:"column:mitra_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_temp_created_at"`

	MitraBranchTemp           *MitraBranchTempPendData            `json:"mitra_branch_temp"`
	MitraCoverageProvinceTemp []MitraCoverageProvinceTempPendData `json:"mitra_cov_prov_temp"`
	MitraCoverageRegencyTemp  []MitraCoverageRegencyTempPendData  `json:"mitra_cov_rgcy_temp"`
	MitraAdministratorTemp    []MitraAdministratorTempPendData    `json:"mitra_adm_temp"`
}

type MitraTempData struct {
	MitraTempID           uint64    `gorm:"column:mitra_temp_id" json:"mitra_temp_id"`
	MitraTypeTempID       uint64    `gorm:"column:mitra_type_temp_id" json:"mitra_type_temp_id"`
	MitraTypeTempCode     string    `gorm:"column:mitra_type_temp_code" json:"mitra_type_temp_code"`
	MitraTypeTempName     string    `gorm:"column:mitra_type_temp_name" json:"mitra_type_temp_name"`
	MitraTypeTempDesc     string    `gorm:"column:mitra_type_temp_desc" json:"mitra_type_temp_desc"`
	MitraTempCode         string    `gorm:"column:mitra_temp_code" json:"mitra_temp_code"`
	MitraTempName         string    `gorm:"column:mitra_temp_name" json:"mitra_temp_name"`
	MitraTempWebsite      string    `gorm:"column:mitra_temp_website" json:"mitra_temp_website"`
	MitraTempEmail        string    `gorm:"column:mitra_temp_email" json:"mitra_temp_email"`
	MitraTempLogo         string    `gorm:"column:mitra_temp_logo" json:"mitra_temp_logo"`
	MitraTempPKSFile      string    `gorm:"column:mitra_temp_pks_file" json:"mitra_temp_pks_file"`
	MitraTempReason       string    `gorm:"column:mitra_temp_reason" json:"mitra_temp_reason"`
	MitraTempStatus       uint64    `gorm:"column:mitra_temp_status;size:2" json:"mitra_temp_status"`
	MitraTempActionBefore uint64    `gorm:"column:mitra_temp_action_before;size:2" json:"mitra_temp_action_before"`
	MitraTempCreatedBy    string    `gorm:"column:mitra_temp_created_by;size:125" json:"mitra_temp_created_by"`
	MitraTempCreatedAt    time.Time `gorm:"column:mitra_temp_created_at" json:"mitra_temp_created_at"`

	MitraBranchTemp           *MitraBranchTempPendData        `json:"mitra_branch_temp"`
	MitraCoverageProvinceTemp []MitraCoverageProvinceTempData `json:"mitra_cov_prov_temp"`
	MitraCoverageRegencyTemp  []MitraCoverageRegencyTempData  `json:"mitra_cov_rgcy_temp"`

	MitraID          uint64    `gorm:"column:mitra_id" json:"mitra_id"`
	MitraTypeID      uint64    `gorm:"column:mitra_type_id" json:"mitra_type_id"`
	MitraTypeCode    string    `gorm:"column:mitra_type_code" json:"mitra_type_code"`
	MitraTypeName    string    `gorm:"column:mitra_type_name" json:"mitra_type_name"`
	MitraTypeDesc    string    `gorm:"column:mitra_type_desc" json:"mitra_type_desc"`
	MitraCode        string    `gorm:"column:mitra_code" json:"mitra_code"`
	MitraReferalCode string    `gorm:"column:mitra_referal_code" json:"mitra_referal_code"`
	MitraName        string    `gorm:"column:mitra_name" json:"mitra_name"`
	MitraWebsite     string    `gorm:"column:mitra_website" json:"mitra_website"`
	MitraEmail       string    `gorm:"column:mitra_email" json:"mitra_email"`
	MitraLogo        string    `gorm:"column:mitra_logo" json:"mitra_logo"`
	MitraPKSFile     string    `gorm:"column:mitra_pks_file" json:"mitra_pks_file"`
	MitraStatus      uint64    `gorm:"column:mitra_status" json:"mitra_status"`
	MitraCreatedBy   string    `gorm:"column:mitra_created_by" json:"mitra_created_by"`
	MitraCreatedAt   time.Time `gorm:"column:mitra_created_at" json:"mitra_created_at"`
	MitraUpdatedBy   string    `gorm:"column:mitra_updated_by" json:"mitra_updated_by"`
	MitraUpdatedAt   time.Time `gorm:"column:mitra_updated_at" json:"mitra_updated_at"`
	MitraDeletedBy   string    `gorm:"column:mitra_deleted_by" json:"mitra_deleted_by"`
	MitraDeletedAt   time.Time `gorm:"column:mitra_deleted_at" json:"mitra_deleted_at"`

	MitraBranch           *MitraBranchData            `json:"mitra_branch"`
	MitraCoverageProvince []MitraCoverageProvinceData `json:"mitra_cov_prov"`
	MitraCoverageRegency  []MitraCoverageRegencyData  `json:"mitra_cov_rgcy"`
}

type ResponseMitraTemps struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    *[]MitraTempData `json:"data"`
}

type ResponseMitraTemp struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    *MitraTempData `json:"data"`
}

type ResponseMitraTempsPend struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    *[]MitraTempPend `json:"data"`
}

type ResponseMitraTempPend struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    *MitraTempPend `json:"data"`
}

type ResponseMitraTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraTemp) TableName() string {
	return "m_mitra_temp"
}

func (MitraTempPend) TableName() string {
	return "m_mitra_temp"
}

func (MitraTempData) TableName() string {
	return "m_mitra_temp"
}

func (p *MitraTemp) Prepare() {
	p.MitraID = p.MitraID
	p.MitraTypeTempID = p.MitraTypeTempID
	p.MitraTempCode = p.MitraTempCode
	p.MitraTempName = p.MitraTempName
	p.MitraTempWebsite = p.MitraTempWebsite
	p.MitraTempEmail = p.MitraTempEmail
	p.MitraTempLogo = p.MitraTempLogo
	p.MitraTempPKSFile = p.MitraTempPKSFile
	p.MitraTempReason = p.MitraTempReason
	p.MitraTempStatus = p.MitraTempStatus
	p.MitraTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraTempCreatedBy))
	p.MitraTempCreatedAt = time.Now()
}

func (p *MitraTempPend) Prepare() {
	p.MitraID = nil
	p.MitraTypeTempID = p.MitraTypeTempID
	p.MitraTempCode = p.MitraTempCode
	p.MitraTempName = p.MitraTempName
	p.MitraTempWebsite = p.MitraTempWebsite
	p.MitraTempEmail = p.MitraTempEmail
	p.MitraTempLogo = p.MitraTempLogo
	p.MitraTempPKSFile = p.MitraTempPKSFile
	p.MitraTempReason = p.MitraTempReason
}

func (p *MitraTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	case "insertupdate":
		if p.MitraTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.MitraTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	case "reason":
		if p.MitraTempReason == "" {
			return errors.New("Alasan Wajib diisi")
		}
		if p.MitraTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil

	default:
		if p.MitraTempCreatedBy == "" {
			return errors.New("Username Wajib diisi")
		}
		return nil
	}
}

func (p *MitraTemp) SaveMitraTemp(db *gorm.DB) (*MitraTemp, error) {
	var err error
	err = db.Debug().Model(&MitraTemp{}).Create(&p).Error
	if err != nil {
		return &MitraTemp{}, err
	}
	return p, nil
}

func (p *MitraTempPend) SaveMitraTempPend(db *gorm.DB) (*MitraTempPend, error) {
	var err error
	err = db.Debug().Model(&MitraTempPend{}).Create(&p).Error
	if err != nil {
		return &MitraTempPend{}, err
	}
	return p, nil
}

func (p *MitraTemp) FindAllMitraTemp(db *gorm.DB) (*[]MitraTempPend, error) {
	var err error
	address := []MitraTempPend{}
	err = db.Debug().Model(&MitraTempPend{}).Limit(100).
		Select(`m_mitra_temp.mitra_temp_id,
				m_mitra_temp.mitra_id,
				m_mitra_temp.mitra_type_temp_id,
				m_mitra_temp.mitra_temp_code,
				m_mitra_temp.mitra_temp_name,
				m_mitra_temp.mitra_temp_website,
				m_mitra_temp.mitra_temp_email,
				m_mitra_temp.mitra_temp_logo,
				m_mitra_temp.mitra_temp_pks_file,
				m_mitra_temp.mitra_temp_reason,
				m_mitra_temp.mitra_temp_status,
				m_mitra_temp.mitra_temp_action_before,
				m_mitra_temp.mitra_temp_created_by,
				m_mitra_temp.mitra_temp_created_at at time zone 'Asia/Jakarta' as mitra_temp_created_at`).
		Order("mitra_temp_created_at desc").
		Find(&address).Error
	if err != nil {
		return &[]MitraTempPend{}, err
	}
	return &address, nil
}

func (p *MitraTemp) FindAllMitraTempPendingActive(db *gorm.DB) (*[]MitraTempPend, error) {
	var err error
	address := []MitraTempPend{}
	err = db.Debug().Model(&MitraTempPend{}).Limit(100).
		Select(`m_mitra_temp.mitra_temp_id,
				m_mitra_temp.mitra_id,
				m_mitra_temp.mitra_type_temp_id,
				m_mitra_temp.mitra_temp_code,
				m_mitra_temp.mitra_temp_name,
				m_mitra_temp.mitra_temp_website,
				m_mitra_temp.mitra_temp_email,
				m_mitra_temp.mitra_temp_logo,
				m_mitra_temp.mitra_temp_pks_file,
				m_mitra_temp.mitra_temp_reason,
				m_mitra_temp.mitra_temp_status,
				m_mitra_temp.mitra_temp_action_before,
				m_mitra_temp.mitra_temp_created_by,
				m_mitra_temp.mitra_temp_created_at at time zone 'Asia/Jakarta' as mitra_temp_created_at`).
		Where("mitra_temp_status = ?", 11).
		Order("mitra_temp_created_at desc").
		Find(&address).Error
	if err != nil {
		return &[]MitraTempPend{}, err
	}
	return &address, nil
}

func (p *MitraTemp) FindAllMitraTempStatus(db *gorm.DB, status uint64) (*[]MitraTempData, error) {
	var err error
	address := []MitraTempData{}
	err = db.Debug().Model(&MitraTempData{}).Limit(100).
		Select(`m_mitra_temp.mitra_temp_id,
				m_mitra_type_temp.mitra_type_id as mitra_type_temp_id,
				m_mitra_type_temp.mitra_type_code as mitra_type_temp_code,
				m_mitra_type_temp.mitra_type_name as mitra_type_temp_name,
				m_mitra_type_temp.mitra_type_desc as mitra_type_temp_desc,
				m_mitra_temp.mitra_temp_code,
				m_mitra_temp.mitra_temp_name,
				m_mitra_temp.mitra_temp_website,
				m_mitra_temp.mitra_temp_email,
				m_mitra_temp.mitra_temp_logo,
				m_mitra_temp.mitra_temp_pks_file,
				m_mitra_temp.mitra_temp_reason,
				m_mitra_temp.mitra_temp_status,
				m_mitra_temp.mitra_temp_action_before,
				m_mitra_temp.mitra_temp_created_by,
				m_mitra_temp.mitra_temp_created_at at time zone 'Asia/Jakarta' as mitra_temp_created_at,
				m_mitra.mitra_id,
				m_mitra_type.mitra_type_id,
				m_mitra_type.mitra_type_code,
				m_mitra_type.mitra_type_name,
				m_mitra_type.mitra_type_desc,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra.mitra_website,
				m_mitra.mitra_email,
				m_mitra.mitra_logo,
				m_mitra.mitra_pks_file,
				m_mitra.mitra_status,
				m_mitra.mitra_created_by,
				m_mitra.mitra_created_at at time zone 'Asia/Jakarta' as mitra_created_at,
				m_mitra.mitra_updated_by,
				m_mitra.mitra_updated_at at time zone 'Asia/Jakarta' as mitra_updated_at,
				m_mitra.mitra_deleted_by,
				m_mitra.mitra_deleted_at at time zone 'Asia/Jakarta' as mitra_deleted_at`).
		Joins("JOIN m_mitra_type m_mitra_type_temp on m_mitra_temp.mitra_type_temp_id=m_mitra_type_temp.mitra_type_id").
		Joins("JOIN m_mitra on m_mitra_temp.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_type on m_mitra.mitra_type_id=m_mitra_type.mitra_type_id").
		Where("mitra_temp_status = ?", status).
		Order("mitra_temp_created_at desc").
		Find(&address).Error
	if err != nil {
		return &[]MitraTempData{}, err
	}
	return &address, nil
}

func (p *MitraTemp) FindMitraTempDataByID(db *gorm.DB, pid uint64) (*MitraTemp, error) {
	var err error
	err = db.Debug().Model(&MitraTemp{}).Where("mitra_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraTemp{}, err
	}
	return p, nil
}

func (p *MitraTemp) FindMitraTempByIDPendingActive(db *gorm.DB, pid uint64) (*MitraTempPend, error) {
	var err error
	address := MitraTempPend{}
	err = db.Debug().Model(&MitraTempPend{}).Limit(100).
		Select(`m_mitra_temp.mitra_temp_id,
				m_mitra_temp.mitra_id,
				m_mitra_temp.mitra_type_temp_id,
				m_mitra_temp.mitra_temp_code,
				m_mitra_temp.mitra_temp_name,
				m_mitra_temp.mitra_temp_website,
				m_mitra_temp.mitra_temp_email,
				m_mitra_temp.mitra_temp_logo,
				m_mitra_temp.mitra_temp_pks_file,
				m_mitra_temp.mitra_temp_reason,
				m_mitra_temp.mitra_temp_status,
				m_mitra_temp.mitra_temp_action_before,
				m_mitra_temp.mitra_temp_created_by,
				m_mitra_temp.mitra_temp_created_at at time zone 'Asia/Jakarta' as mitra_temp_created_at`).
		Where("mitra_temp_id = ? AND mitra_temp_status = ?", pid, 11).
		Order("mitra_temp_created_at desc").
		Find(&address).Error
	if err != nil {
		return &MitraTempPend{}, err
	}
	return &address, nil
}

func (p *MitraTemp) FindMitraTempByID(db *gorm.DB, pid uint64) (*MitraTempData, error) {
	var err error
	address := MitraTempData{}
	err = db.Debug().Model(&MitraTempData{}).Limit(100).
		Select(`m_mitra_temp.mitra_temp_id,
				m_mitra_type_temp.mitra_type_id as mitra_type_temp_id,
				m_mitra_type_temp.mitra_type_code as mitra_type_temp_code,
				m_mitra_type_temp.mitra_type_name as mitra_type_temp_name,
				m_mitra_type_temp.mitra_type_desc as mitra_type_temp_desc,
				m_mitra_temp.mitra_temp_code,
				m_mitra_temp.mitra_temp_name,
				m_mitra_temp.mitra_temp_website,
				m_mitra_temp.mitra_temp_email,
				m_mitra_temp.mitra_temp_logo,
				m_mitra_temp.mitra_temp_pks_file,
				m_mitra_temp.mitra_temp_reason,
				m_mitra_temp.mitra_temp_status,
				m_mitra_temp.mitra_temp_action_before,
				m_mitra_temp.mitra_temp_created_by,
				m_mitra_temp.mitra_temp_created_at at time zone 'Asia/Jakarta' as mitra_temp_created_at,
				m_mitra.mitra_id,
				m_mitra_type.mitra_type_id,
				m_mitra_type.mitra_type_code,
				m_mitra_type.mitra_type_name,
				m_mitra_type.mitra_type_desc,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra.mitra_website,
				m_mitra.mitra_email,
				m_mitra.mitra_logo,
				m_mitra.mitra_pks_file,
				m_mitra.mitra_status,
				m_mitra.mitra_created_by,
				m_mitra.mitra_created_at at time zone 'Asia/Jakarta' as mitra_created_at,
				m_mitra.mitra_updated_by,
				m_mitra.mitra_updated_at at time zone 'Asia/Jakarta' as mitra_updated_at,
				m_mitra.mitra_deleted_by,
				m_mitra.mitra_deleted_at at time zone 'Asia/Jakarta' as mitra_deleted_at`).
		Joins("JOIN m_mitra_type m_mitra_type_temp on m_mitra_temp.mitra_type_temp_id=m_mitra_type_temp.mitra_type_id").
		Joins("JOIN m_mitra on m_mitra_temp.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_type on m_mitra.mitra_type_id=m_mitra_type.mitra_type_id").
		Where("mitra_temp_id = ?", pid).
		Order("mitra_temp_created_at desc").
		Find(&address).Error
	if err != nil {
		return &MitraTempData{}, err
	}
	return &address, nil
}

func (p *MitraTemp) FindMitraTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraTempData, error) {
	var err error
	address := MitraTempData{}
	err = db.Debug().Model(&MitraTempData{}).Limit(100).
		Select(`m_mitra_temp.mitra_temp_id,
				m_mitra_type_temp.mitra_type_id as mitra_type_temp_id,
				m_mitra_type_temp.mitra_type_code as mitra_type_temp_code,
				m_mitra_type_temp.mitra_type_name as mitra_type_temp_name,
				m_mitra_type_temp.mitra_type_desc as mitra_type_temp_desc,
				m_mitra_temp.mitra_temp_code,
				m_mitra_temp.mitra_temp_name,
				m_mitra_temp.mitra_temp_website,
				m_mitra_temp.mitra_temp_email,
				m_mitra_temp.mitra_temp_logo,
				m_mitra_temp.mitra_temp_pks_file,
				m_mitra_temp.mitra_temp_reason,
				m_mitra_temp.mitra_temp_status,
				m_mitra_temp.mitra_temp_action_before,
				m_mitra_temp.mitra_temp_created_by,
				m_mitra_temp.mitra_temp_created_at at time zone 'Asia/Jakarta' as mitra_temp_created_at,
				m_mitra.mitra_id,
				m_mitra_type.mitra_type_id,
				m_mitra_type.mitra_type_code,
				m_mitra_type.mitra_type_name,
				m_mitra_type.mitra_type_desc,
				m_mitra.mitra_code,
				m_mitra.mitra_name,
				m_mitra.mitra_website,
				m_mitra.mitra_email,
				m_mitra.mitra_logo,
				m_mitra.mitra_pks_file,
				m_mitra.mitra_status,
				m_mitra.mitra_created_by,
				m_mitra.mitra_created_at at time zone 'Asia/Jakarta' as mitra_created_at,
				m_mitra.mitra_updated_by,
				m_mitra.mitra_updated_at at time zone 'Asia/Jakarta' as mitra_updated_at,
				m_mitra.mitra_deleted_by,
				m_mitra.mitra_deleted_at at time zone 'Asia/Jakarta' as mitra_deleted_at`).
		Joins("JOIN m_mitra_type m_mitra_type_temp on m_mitra_temp.mitra_type_temp_id=m_mitra_type_temp.mitra_type_id").
		Joins("JOIN m_mitra on m_mitra_temp.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_type on m_mitra.mitra_type_id=m_mitra_type.mitra_type_id").
		Where("mitra_temp_id = ? AND mitra_temp_status = ?", pid, status).
		Order("mitra_temp_created_at desc").
		Find(&address).Error
	if err != nil {
		return &MitraTempData{}, err
	}
	return &address, nil
}

func (p *MitraTemp) UpdateMitraTemp(db *gorm.DB, pid uint64) (*MitraTemp, error) {
	var err error
	db = db.Debug().Model(&MitraTemp{}).Where("mitra_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_type_temp_id":  p.MitraTypeTempID,
			"mitra_temp_code":     p.MitraTempCode,
			"mitra_temp_name":     p.MitraTempName,
			"mitra_temp_website":  p.MitraTempWebsite,
			"mitra_temp_email":    p.MitraTempEmail,
			"mitra_temp_logo":     p.MitraTempLogo,
			"mitra_temp_pks_file": p.MitraTempPKSFile,
			"mitra_temp_reason":   p.MitraTempReason,
		},
	)
	err = db.Debug().Model(&MitraTemp{}).Where("mitra_temp_id = ?", pid).Error
	if err != nil {
		return &MitraTemp{}, err
	}
	return p, nil
}

func (p *MitraTemp) UpdateMitraTempStatus(db *gorm.DB, pid uint64) (*MitraTemp, error) {
	var err error
	db = db.Debug().Model(&MitraTemp{}).Where("mitra_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_temp_reason": p.MitraTempReason,
			"mitra_temp_status": p.MitraTempStatus,
		},
	)
	err = db.Debug().Model(&MitraTemp{}).Where("mitra_temp_id = ?", pid).Error
	if err != nil {
		return &MitraTemp{}, err
	}
	return p, nil
}

func (p *MitraTemp) DeleteMitraTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraTemp{}).Where("mitra_temp_id = ?", pid).Take(&MitraTemp{}).Delete(&MitraTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

type MitraBranchInputTemp struct {
	MitraBranchTempCode    string `gorm:"column:mitra_branch_temp_code" json:"mitra_branch_temp_code"`
	MitraBranchTempName    string `gorm:"column:mitra_branch_temp_name" json:"mitra_branch_temp_name"`
	MitraBranchTempType    string `gorm:"column:mitra_branch_temp_type" json:"mitra_branch_temp_type"`
	MitraBranchTempPhone   string `gorm:"column:mitra_branch_temp_phone" json:"mitra_branch_temp_phone"`
	MitraBranchTempFax     string `gorm:"column:mitra_branch_temp_fax" json:"mitra_branch_temp_fax"`
	MitraBranchTempAddress string `gorm:"column:mitra_branch_temp_address" json:"mitra_branch_temp_address"`
	CountryTempID          uint64 `gorm:"column:country_temp_id" json:"country_temp_id"`
	ProvinceTempID         uint64 `gorm:"column:province_temp_id" json:"province_temp_id"`
	RegencyTempID          uint64 `gorm:"column:regency_temp_id" json:"regency_temp_id"`
	DistrictTempID         uint64 `gorm:"column:district_temp_id" json:"district_temp_id"`
	VillageTempID          uint64 `gorm:"column:village_temp_id" json:"village_temp_id"`
}

type MitraCoverageProvinceInputTemp struct {
	ProvinceTempID uint64 `gorm:"column:province_temp_id" json:"province_temp_id"`
}

type MitraCoverageRegencyInputTemp struct {
	RegencyTempID uint64 `gorm:"column:regency_temp_id" json:"regency_temp_id"`
}

type MitraAdministratorInputTemp struct {
	MitraEmployeeTempName           string `gorm:"column:mitra_emp_temp_name" json:"mitra_emp_temp_name"`
	MitraEmployeeTempAddress        string `gorm:"column:mitra_emp_temp_address" json:"mitra_emp_temp_address"`
	MitraGroupAccessTempID          uint64 `gorm:"column:mitra_group_accs_temp_id" json:"mitra_group_accs_temp_id"`
	MitraAdministratorTempEmail     string `gorm:"column:mitra_adm_temp_email" json:"mitra_adm_temp_email"`
	MitraAdministratorTempUsername  string `gorm:"column:mitra_adm_temp_username" json:"mitra_adm_temp_username"`
	MitraAdministratorTempPassword  string `gorm:"column:mitra_adm_temp_password" json:"mitra_adm_temp_password"`
	MitraAdministratorTempPhoneCode string `gorm:"column:mitra_adm_temp_phone_code" json:"mitra_adm_temp_phone_code"`
	MitraAdministratorTempPhone     string `gorm:"column:mitra_adm_temp_phone" json:"mitra_adm_temp_phone"`
}
