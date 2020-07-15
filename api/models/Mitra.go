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

type Mitra struct {
	MitraID          uint64     `gorm:"column:mitra_id;primary_key;" json:"mitra_id"`
	MitraTypeID      uint64     `gorm:"column:mitra_type_id" json:"mitra_type_id"`
	MitraCode        string     `gorm:"column:mitra_code" json:"mitra_code"`
	MitraReferalCode string     `gorm:"column:mitra_referal_code" json:"mitra_referal_code"`
	MitraName        string     `gorm:"column:mitra_name" json:"mitra_name"`
	MitraWebsite     string     `gorm:"column:mitra_website" json:"mitra_website"`
	MitraEmail       string     `gorm:"column:mitra_email" json:"mitra_email"`
	MitraLogo        string     `gorm:"column:mitra_logo" json:"mitra_logo"`
	MitraInitial     string     `gorm:"column:mitra_initial" json:"mitra_initial"`
	MitraPKSFile     string     `gorm:"column:mitra_pks_file" json:"mitra_pks_file"`
	MitraStatus      uint64     `gorm:"column:mitra_status;size:2" json:"mitra_status"`
	MitraCreatedBy   string     `gorm:"column:mitra_created_by;size:125" json:"mitra_created_by"`
	MitraCreatedAt   time.Time  `gorm:"column:mitra_created_at;default:CURRENT_TIMESTAMP" json:"mitra_created_at"`
	MitraUpdatedBy   string     `gorm:"column:mitra_updated_by;size:125" json:"mitra_updated_by"`
	MitraUpdatedAt   *time.Time `gorm:"column:mitra_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_created_at"`
	MitraDeletedBy   string     `gorm:"column:mitra_deleted_by;size:125" json:"mitra_deleted_by"`
	MitraDeletedAt   *time.Time `gorm:"column:mitra_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_deleted_at"`

	LoanRequestAmount uint64 `gorm:"-" json:"loan_req_amount"`
	LoanRequestLong   uint64 `gorm:"-" json:"loan_req_long"`
}

type MitraData struct {
	MitraID          uint64     `gorm:"column:mitra_id" json:"mitra_id"`
	MitraTypeID      uint64     `gorm:"column:mitra_type_id" json:"mitra_type_id"`
	MitraTypeCode    string     `gorm:"column:mitra_type_code" json:"mitra_type_code"`
	MitraTypeName    string     `gorm:"column:mitra_type_name" json:"mitra_type_name"`
	MitraTypeDesc    string     `gorm:"column:mitra_type_desc" json:"mitra_type_desc"`
	MitraCode        string     `gorm:"column:mitra_code" json:"mitra_code"`
	MitraReferalCode string     `gorm:"column:mitra_referal_code" json:"mitra_referal_code"`
	MitraName        string     `gorm:"column:mitra_name" json:"mitra_name"`
	MitraWebsite     string     `gorm:"column:mitra_website" json:"mitra_website"`
	MitraEmail       string     `gorm:"column:mitra_email" json:"mitra_email"`
	MitraLogo        string     `gorm:"column:mitra_logo" json:"mitra_logo"`
	MitraInitial     string     `gorm:"column:mitra_initial" json:"mitra_initial"`
	MitraPKSFile     string     `gorm:"column:mitra_pks_file" json:"mitra_pks_file"`
	MitraStatus      uint64     `gorm:"column:mitra_status" json:"mitra_status"`
	MitraCreatedBy   string     `gorm:"column:mitra_created_by" json:"mitra_created_by"`
	MitraCreatedAt   time.Time  `gorm:"column:mitra_created_at" json:"mitra_created_at"`
	MitraUpdatedBy   string     `gorm:"column:mitra_updated_by" json:"mitra_updated_by"`
	MitraUpdatedAt   *time.Time `gorm:"column:mitra_updated_at" json:"mitra_updated_at"`
	MitraDeletedBy   string     `gorm:"column:mitra_deleted_by" json:"mitra_deleted_by"`
	MitraDeletedAt   *time.Time `gorm:"column:mitra_deleted_at" json:"mitra_deleted_at"`

	MitraBranch           *MitraBranchData            `json:"mitra_branch"`
	MitraCoverageProvince []MitraCoverageProvinceData `json:"mitra_cov_prov"`
	MitraCoverageRegency  []MitraCoverageRegencyData  `json:"mitra_cov_rgcy"`
}

type ResponseMitras struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    *[]MitraData `json:"data"`
}

type ResponseMitra struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *MitraData `json:"data"`
}

type ResponseMitraCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (Mitra) TableName() string {
	return "m_mitra"
}

func (MitraData) TableName() string {
	return "m_mitra"
}

func (p *Mitra) Prepare() {
	p.MitraTypeID = p.MitraTypeID
	p.MitraCode = p.MitraCode
	p.MitraReferalCode = p.MitraReferalCode
	p.MitraName = p.MitraName
	p.MitraWebsite = p.MitraWebsite
	p.MitraEmail = p.MitraEmail
	p.MitraLogo = p.MitraLogo
	p.MitraPKSFile = p.MitraPKSFile
	p.MitraStatus = p.MitraStatus
	p.MitraCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraCreatedBy))
	p.MitraCreatedAt = time.Now()
	p.MitraUpdatedBy = html.EscapeString(strings.TrimSpace(p.MitraUpdatedBy))
	p.MitraUpdatedAt = p.MitraUpdatedAt
	p.MitraDeletedBy = html.EscapeString(strings.TrimSpace(p.MitraDeletedBy))
	p.MitraDeletedAt = p.MitraDeletedAt
}

func (p *Mitra) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.MitraTypeID == 0 {
			return errors.New("Required Country")
		}
		if p.MitraCode == "" {
			return errors.New("Required Province")
		}
		if p.MitraReferalCode == "" {
			return errors.New("Required Regency")
		}
		if p.MitraName == "" {
			return errors.New("Required District")
		}
		if p.MitraEmail == "" {
			return errors.New("Required Village")
		}
		return nil
	}
}

func (p *Mitra) SaveMitra(db *gorm.DB) (*Mitra, error) {
	var err error
	err = db.Debug().Model(&Mitra{}).Create(&p).Error
	if err != nil {
		return &Mitra{}, err
	}
	return p, nil
}

func (p *Mitra) FindAllMitra(db *gorm.DB) (*[]MitraData, error) {
	var err error
	address := []MitraData{}
	err = db.Debug().Model(&MitraData{}).Limit(100).
		Select(`m_mitra.mitra_id,
				m_mitra_type.mitra_type_id,
				m_mitra_type.mitra_type_code,
				m_mitra_type.mitra_type_name,
				m_mitra_type.mitra_type_desc,
				m_mitra.mitra_code,
				m_mitra.mitra_referal_code,
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
		Joins("JOIN m_mitra_type on m_mitra.mitra_type_id=m_mitra_type.mitra_type_id").
		Find(&address).Error
	if err != nil {
		return &[]MitraData{}, err
	}
	return &address, nil
}

func (p *Mitra) FindAllMitraStatus(db *gorm.DB, status uint64) (*[]MitraData, error) {
	var err error
	address := []MitraData{}
	err = db.Debug().Model(&MitraData{}).Limit(100).
		Select(`m_mitra.mitra_id,
				m_mitra_type.mitra_type_id,
				m_mitra_type.mitra_type_code,
				m_mitra_type.mitra_type_name,
				m_mitra_type.mitra_type_desc,
				m_mitra.mitra_code,
				m_mitra.mitra_referal_code,
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
		Joins("JOIN m_mitra_type on m_mitra.mitra_type_id=m_mitra_type.mitra_type_id").
		Where("mitra_status = ?", status).
		Find(&address).Error
	if err != nil {
		return &[]MitraData{}, err
	}
	return &address, nil
}

func (p *Mitra) FindMitraByID(db *gorm.DB, pid uint64) (*Mitra, error) {
	var err error
	err = db.Debug().Model(&Mitra{}).Where("mitra_id = ?", pid).Take(&p).Error
	if err != nil {
		return &Mitra{}, err
	}
	return p, nil
}

func (p *Mitra) FindMitraDataByID(db *gorm.DB, pid uint64) (*MitraData, error) {
	var err error
	address := MitraData{}
	err = db.Debug().Model(&MitraData{}).Limit(100).
		Select(`m_mitra.mitra_id,
				m_mitra_type.mitra_type_id,
				m_mitra_type.mitra_type_code,
				m_mitra_type.mitra_type_name,
				m_mitra_type.mitra_type_desc,
				m_mitra.mitra_code,
				m_mitra.mitra_initial,
				m_mitra.mitra_referal_code,
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
		Joins("JOIN m_mitra_type on m_mitra.mitra_type_id=m_mitra_type.mitra_type_id").
		Where("mitra_id = ?", pid).
		Find(&address).Error
	if err != nil {
		return &MitraData{}, err
	}
	return &address, nil
}

func (p *Mitra) FindMitraStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraData, error) {
	var err error
	address := MitraData{}
	err = db.Debug().Model(&MitraData{}).Limit(100).
		Select(`m_mitra.mitra_id,
				m_mitra_type.mitra_type_id,
				m_mitra_type.mitra_type_code,
				m_mitra_type.mitra_type_name,
				m_mitra_type.mitra_type_desc,
				m_mitra.mitra_code,
				m_mitra.mitra_referal_code,
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
		Joins("JOIN m_mitra_type on m_mitra.mitra_type_id=m_mitra_type.mitra_type_id").
		Where("mitra_id = ? AND mitra_status = ?", pid, status).
		Find(&address).Error
	if err != nil {
		return &MitraData{}, err
	}
	return &address, nil
}

func (p *Mitra) UpdateMitra(db *gorm.DB, pid uint64) (*Mitra, error) {
	var err error
	db = db.Debug().Model(&Mitra{}).Where("mitra_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_type_id":    p.MitraTypeID,
			"mitra_code":       p.MitraCode,
			"mitra_name":       p.MitraName,
			"mitra_website":    p.MitraWebsite,
			"mitra_email":      p.MitraEmail,
			"mitra_logo":       p.MitraLogo,
			"mitra_pks_file":   p.MitraPKSFile,
			"mitra_status":     p.MitraStatus,
			"mitra_updated_by": p.MitraUpdatedBy,
			"mitra_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&Mitra{}).Where("mitra_id = ?", pid).Error
	if err != nil {
		return &Mitra{}, err
	}
	return p, nil
}

func (p *Mitra) UpdateMitraStatus(db *gorm.DB, pid uint64) (*Mitra, error) {
	var err error
	db = db.Debug().Model(&Mitra{}).Where("mitra_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_status":     p.MitraStatus,
			"mitra_updated_by": p.MitraUpdatedBy,
			"mitra_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&Mitra{}).Where("mitra_id = ?", pid).Error
	if err != nil {
		return &Mitra{}, err
	}
	return p, nil
}

func (p *Mitra) UpdateMitraStatusOnly(db *gorm.DB, pid uint64) (*Mitra, error) {
	var err error
	db = db.Debug().Model(&Mitra{}).Where("mitra_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_status": p.MitraStatus,
		},
	)
	err = db.Debug().Model(&Mitra{}).Where("mitra_id = ?", pid).Error
	if err != nil {
		return &Mitra{}, err
	}
	return p, nil
}

func (p *Mitra) UpdateMitraStatusDelete(db *gorm.DB, pid uint64) (*Mitra, error) {
	var err error
	db = db.Debug().Model(&Mitra{}).Where("mitra_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_status":     3,
			"mitra_deleted_by": p.MitraDeletedBy,
			"mitra_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&Mitra{}).Where("mitra_id = ?", pid).Error
	if err != nil {
		return &Mitra{}, err
	}
	return p, nil
}

func (p *Mitra) DeleteMitra(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&Mitra{}).Where("mitra_id = ?", pid).Take(&Mitra{}).Delete(&Mitra{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Mitra not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

type MitraLoanList struct {
	MitraID                        uint64 `gorm:"column:mitra_id;primary_key;" json:"mitra_id"`
	MitraType                      string `gorm:"column:mitra_type;size:25" json:"mitra_type"`
	MitraName                      string `gorm:"column:mitra_name;size:255" json:"mitra_name"`
	MitraWebsite                   string `gorm:"column:mitra_website;size:255" json:"mitra_website"`
	MitraEmail                     string `gorm:"column:mitra_email;size:255" json:"mitra_email"`
	MitraLogo                      string `gorm:"column:mitra_logo;size:255" json:"mitra_logo"`
	MitraBranchName                string `gorm:"column:mitra_branch_name;size:255" json:"mitra_branch_name"`
	MitraBranchPhone               string `gorm:"column:mitra_branch_phone;size:20" json:"mitra_branch_phone"`
	MitraBranchFax                 string `gorm:"column:mitra_branch_fax;size:20" json:"mitra_branch_fax"`
	MitraBranchAddress             string `gorm:"column:mitra_branch_address" json:"mitra_branch_address"`
	CountryName                    string `gorm:"column:country_name;size:255" json:"country_name"`
	ProvinceName                   string `gorm:"column:province_name;size:255" json:"province_name"`
	RegencyName                    string `gorm:"column:regency_name;size:255" json:"regency_name"`
	DistrictName                   string `gorm:"column:district_name;size:255" json:"district_name"`
	VillageName                    string `gorm:"column:village_name;size:255" json:"village_name"`
	MitraProductLimitMax           uint64 `gorm:"column:mitra_prod_limit_max" json:"mitra_prod_limit_max"`
	MitraProductInterestRate       uint64 `gorm:"column:mitra_prod_interest_rate" json:"mitra_prod_interest_rate"`
	MitraProductInterestRatePeriod string `gorm:"column:mitra_prod_interest_rate_period;size:125" json:"mitra_prod_interest_rate_period"`
	ProductSubCategoryName         string `gorm:"column:prod_sub_ctgry_name;size:255" json:"prod_sub_ctgry_name"`
}

type ResponseMitraLoanLists struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    *[]MitraLoanList `json:"data"`
}

func (MitraLoanList) TableName() string {
	return "m_mitra"
}

func (p *Mitra) FindMitraLoanList(db *gorm.DB, loanreqamount uint64, loanreqlong uint64, ctgryid uint64, regencyid uint64) (*[]MitraLoanList, error) {
	var err error
	mitraloanlist := []MitraLoanList{}
	err = db.Debug().Model(&MitraLoanList{}).Limit(100).
		Select(`m_mitra.mitra_id,
				m_mitra.mitra_name,
				m_mitra.mitra_website,
				m_mitra.mitra_email,
				m_mitra.mitra_logo,
				m_mitra_branch.mitra_branch_name,
				m_mitra_branch.mitra_branch_phone,
				m_mitra_branch.mitra_branch_fax,
				m_mitra_branch.mitra_branch_address,
				m_country.country_name,
				m_province.province_name,
				m_regency.regency_name,
				m_district.district_name,
				m_village.village_name,
				m_mitra_product.mitra_prod_limit_max,
				m_mitra_product.mitra_prod_interest_rate,
				m_mitra_product.mitra_prod_interest_rate_period,
				m_product_sub_category.prod_sub_ctgry_name
				`).
		Joins("JOIN m_mitra_branch on m_mitra.mitra_id=m_mitra_branch.mitra_id").
		Joins("JOIN m_address on m_mitra_branch.address_id=m_address.address_id").
		Joins("JOIN m_country on m_address.country_id=m_country.country_id").
		Joins("JOIN m_province on m_address.province_id=m_province.province_id").
		Joins("JOIN m_regency on m_address.regency_id=m_regency.regency_id").
		Joins("JOIN m_district on m_address.district_id=m_district.district_id").
		Joins("JOIN m_village on m_address.village_id=m_village.village_id").
		Joins("JOIN m_mitra_product on m_mitra.mitra_id=m_mitra_product.mitra_id").
		Joins("JOIN m_mitra_product_sub_category on m_mitra_product.mitra_prod_id=m_mitra_product_sub_category.mitra_prod_id").
		Joins("JOIN m_product_sub_category on m_mitra_product_sub_category.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Where("mitra_branch_type = ?", `PST`).
		Where("m_mitra_product_sub_category.prod_sub_ctgry_id = ?", ctgryid).
		Where("m_address.regency_id = ?", regencyid).
		Where("? BETWEEN mitra_prod_limit_min AND mitra_prod_limit_max", loanreqamount).
		Where("? BETWEEN mitra_prod_tenor_min AND mitra_prod_tenor_max", loanreqlong).
		Group(`m_mitra.mitra_id,
			   m_mitra.mitra_name,
			   m_mitra.mitra_website,
			   m_mitra.mitra_email,
			   m_mitra.mitra_logo,
			   m_mitra_branch.mitra_branch_name,
			   m_mitra_branch.mitra_branch_phone,
			   m_mitra_branch.mitra_branch_fax,
			   m_mitra_branch.mitra_branch_address,
			   m_country.country_name,
			   m_province.province_name,
			   m_regency.regency_name,
			   m_district.district_name,
			   m_village.village_name,
			   m_mitra_product.mitra_prod_limit_max,
			   m_mitra_product.mitra_prod_interest_rate,
			   m_mitra_product.mitra_prod_interest_rate_period,
			   m_product_sub_category.prod_sub_ctgry_name
			   `).
		Find(&mitraloanlist).Error
	if err != nil {
		return &[]MitraLoanList{}, err
	}
	return &mitraloanlist, nil
}

func (p *Mitra) FindAllMitraReferalCode(db *gorm.DB, referal string) (*Mitra, error) {
	var err error
	err = db.Debug().Model(&Mitra{}).Where("mitra_referal_code = ? AND mitra_status = ?", referal, 1).Take(&p).Error
	if err != nil {
		return &Mitra{}, err
	}
	return p, nil
}

func (p *Mitra) FindMitraStatusByCode(db *gorm.DB, code string) (*Mitra, error) {
	var err error
	err = db.Debug().Model(&Mitra{}).Where("mitra_code = ? AND mitra_status = ?", code, 1).Take(&p).Error
	if err != nil {
		return &Mitra{}, err
	}
	return p, nil
}
