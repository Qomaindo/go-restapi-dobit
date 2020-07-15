package models

import (
	"errors"
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type MitraProductDocumentRule struct {
	MitraProductDocumentRuleID          *uint64    `gorm:"column:mitra_prod_doc_rule_id;primary_key;" json:"mitra_prod_doc_rule_id"`
	MitraProductID                      uint64     `gorm:"column:mitra_prod_id" json:"mitra_prod_id"`
	ProductDocumentID                   uint64     `gorm:"column:prod_doc_id" json:"prod_doc_id"`
	CustomerSubType1                    bool       `gorm:"column:cust_sub_type1" json:"cust_sub_type1"`
	CustomerSubType2                    bool       `gorm:"column:cust_sub_type2" json:"cust_sub_type2"`
	CustomerSubType3                    bool       `gorm:"column:cust_sub_type3" json:"cust_sub_type3"`
	CustomerSubType4                    bool       `gorm:"column:cust_sub_type4" json:"cust_sub_type4"`
	CustomerSubType5                    bool       `gorm:"column:cust_sub_type5" json:"cust_sub_type5"`
	CustomerSubType6                    bool       `gorm:"column:cust_sub_type6" json:"cust_sub_type6"`
	CustomerSubType7                    bool       `gorm:"column:cust_sub_type7" json:"cust_sub_type7"`
	CustomerSubType8                    bool       `gorm:"column:cust_sub_type8" json:"cust_sub_type8"`
	MitraProductDocumentRuleIsMandatory bool       `gorm:"column:mitra_prod_doc_rule_is_mandatory" json:"mitra_prod_doc_rule_is_mandatory"`
	MitraProductDocumentRuleIsManyImage bool       `gorm:"column:mitra_prod_doc_rule_is_many_img" json:"mitra_prod_doc_rule_is_many_img"`
	MitraProductDocumentRuleStatus      uint64     `gorm:"column:mitra_prod_doc_rule_status" json:"mitra_prod_doc_rule_status"`
	MitraProductDocumentRuleCreatedBy   string     `gorm:"column:mitra_prod_doc_rule_created_by;size:125" json:"mitra_prod_doc_rule_created_by"`
	MitraProductDocumentRuleCreatedAt   time.Time  `gorm:"column:mitra_prod_doc_rule_created_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_doc_rule_created_at"`
	MitraProductDocumentRuleUpdatedBy   string     `gorm:"column:mitra_prod_doc_rule_updated_by;size:125" json:"mitra_prod_doc_rule_updated_by"`
	MitraProductDocumentRuleUpdatedAt   *time.Time `gorm:"column:mitra_prod_doc_rule_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_doc_rule_updated_at"`
	MitraProductDocumentRuleDeletedBy   string     `gorm:"column:mitra_prod_doc_rule_deleted_by;size:125" json:"mitra_prod_doc_rule_deleted_by"`
	MitraProductDocumentRuleDeletedAt   *time.Time `gorm:"column:mitra_prod_doc_rule_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_doc_rule_deleted_at"`
}

type MitraProductDocumentRuleData struct {
	MitraProductDocumentRuleID          uint64     `gorm:"column:mitra_prod_doc_rule_id" json:"mitra_prod_doc_rule_id"`
	MitraProductID                      uint64     `gorm:"column:mitra_prod_id" json:"mitra_prod_id"`
	MitraProductName                    string     `gorm:"column:mitra_prod_name" json:"mitra_prod_name"`
	ProductDocumentID                   uint64     `gorm:"column:prod_doc_id" json:"prod_doc_id"`
	ProductDocumentCode                 string     `gorm:"column:prod_doc_code" json:"prod_doc_code"`
	ProductDocumentName                 string     `gorm:"column:prod_doc_name" json:"prod_doc_name"`
	ProductDocumentCategoryName         string     `gorm:"column:prod_doc_ctgry_name" json:"prod_doc_ctgry_name"`
	ProductDocumentCategoryDesc         string     `gorm:"column:prod_doc_ctgry_desc" json:"prod_doc_ctgry_desc"`
	CustomerSubType1                    bool       `gorm:"column:cust_sub_type1" json:"cust_sub_type1"`
	CustomerSubType2                    bool       `gorm:"column:cust_sub_type2" json:"cust_sub_type2"`
	CustomerSubType3                    bool       `gorm:"column:cust_sub_type3" json:"cust_sub_type3"`
	CustomerSubType4                    bool       `gorm:"column:cust_sub_type4" json:"cust_sub_type4"`
	CustomerSubType5                    bool       `gorm:"column:cust_sub_type5" json:"cust_sub_type5"`
	CustomerSubType6                    bool       `gorm:"column:cust_sub_type6" json:"cust_sub_type6"`
	CustomerSubType7                    bool       `gorm:"column:cust_sub_type7" json:"cust_sub_type7"`
	CustomerSubType8                    bool       `gorm:"column:cust_sub_type8" json:"cust_sub_type8"`
	MitraProductDocumentRuleIsMandatory bool       `gorm:"column:mitra_prod_doc_rule_is_mandatory" json:"mitra_prod_doc_rule_is_mandatory"`
	MitraProductDocumentRuleIsManyImage bool       `gorm:"column:mitra_prod_doc_rule_is_many_img" json:"mitra_prod_doc_rule_is_many_img"`
	MitraProductDocumentRuleStatus      uint64     `gorm:"column:mitra_prod_doc_rule_status" json:"mitra_prod_doc_rule_status"`
	MitraProductDocumentRuleCreatedBy   string     `gorm:"column:mitra_prod_doc_rule_created_by;size:125" json:"mitra_prod_doc_rule_created_by"`
	MitraProductDocumentRuleCreatedAt   time.Time  `gorm:"column:mitra_prod_doc_rule_created_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_doc_rule_created_at"`
	MitraProductDocumentRuleUpdatedBy   string     `gorm:"column:mitra_prod_doc_rule_updated_by;size:125" json:"mitra_prod_doc_rule_updated_by"`
	MitraProductDocumentRuleUpdatedAt   *time.Time `gorm:"column:mitra_prod_doc_rule_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_doc_rule_updated_at"`
	MitraProductDocumentRuleDeletedBy   string     `gorm:"column:mitra_prod_doc_rule_deleted_by;size:125" json:"mitra_prod_doc_rule_deleted_by"`
	MitraProductDocumentRuleDeletedAt   *time.Time `gorm:"column:mitra_prod_doc_rule_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_doc_rule_deleted_at"`
}

type ResponseMitraProductDocumentRules struct {
	Status  int                             `json:"status"`
	Message string                          `json:"message"`
	Data    *[]MitraProductDocumentRuleData `json:"data"`
}

type ResponseMitraProductDocumentRule struct {
	Status  int                           `json:"status"`
	Message string                        `json:"message"`
	Data    *MitraProductDocumentRuleData `json:"data"`
}

type ResponseMitraProductDocumentRuleIU struct {
	Status  int                       `json:"status"`
	Message string                    `json:"message"`
	Data    *MitraProductDocumentRule `json:"data"`
}

type ResponseMitraProductDocumentRuleCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraProductDocumentRule) TableName() string {
	return "m_mitra_product_document_rule"
}

func (MitraProductDocumentRuleData) TableName() string {
	return "m_mitra_product_document_rule"
}

func (p *MitraProductDocumentRule) Prepare() {
	p.MitraProductDocumentRuleID = nil
	p.MitraProductID = p.MitraProductID
	p.ProductDocumentID = p.ProductDocumentID
	p.CustomerSubType1 = p.CustomerSubType1
	p.CustomerSubType2 = p.CustomerSubType2
	p.CustomerSubType3 = p.CustomerSubType3
	p.CustomerSubType4 = p.CustomerSubType4
	p.CustomerSubType5 = p.CustomerSubType5
	p.CustomerSubType6 = p.CustomerSubType6
	p.CustomerSubType7 = p.CustomerSubType7
	p.CustomerSubType8 = p.CustomerSubType8
	p.MitraProductDocumentRuleIsMandatory = p.MitraProductDocumentRuleIsMandatory
	p.MitraProductDocumentRuleIsManyImage = p.MitraProductDocumentRuleIsManyImage
	p.MitraProductDocumentRuleStatus = p.MitraProductDocumentRuleStatus
	p.MitraProductDocumentRuleCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraProductDocumentRuleCreatedBy))
	p.MitraProductDocumentRuleCreatedAt = time.Now()
	p.MitraProductDocumentRuleUpdatedBy = html.EscapeString(strings.TrimSpace(p.MitraProductDocumentRuleUpdatedBy))
	p.MitraProductDocumentRuleUpdatedAt = p.MitraProductDocumentRuleUpdatedAt
	p.MitraProductDocumentRuleDeletedBy = html.EscapeString(strings.TrimSpace(p.MitraProductDocumentRuleDeletedBy))
	p.MitraProductDocumentRuleDeletedAt = p.MitraProductDocumentRuleDeletedAt
}

func (p *MitraProductDocumentRule) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.MitraProductID == 0 {
			return errors.New("Required MitraProductDocumentRule Code")
		}
		if p.ProductDocumentID == 0 {
			return errors.New("Required MitraProductDocumentRule Code")
		}
		return nil
	}
}

func (p *MitraProductDocumentRule) SaveMitraProductDocumentRule(db *gorm.DB) (*MitraProductDocumentRule, error) {
	var err error
	err = db.Debug().Model(&MitraProductDocumentRule{}).Create(&p).Error
	if err != nil {
		return &MitraProductDocumentRule{}, err
	}
	return p, nil
}

func (p *MitraProductDocumentRule) FindAllMitraProductDocumentRule(db *gorm.DB) (*[]MitraProductDocumentRuleData, error) {
	var err error
	mitraproductdocumentrule := []MitraProductDocumentRuleData{}
	err = db.Debug().Model(&MitraProductDocumentRuleData{}).Limit(100).
		Select(`m_mitra_product_document_rule.mitra_prod_doc_rule_id,
				m_mitra_product.mitra_prod_id,
				m_mitra_product.mitra_prod_name,
				m_product_document.prod_doc_id,
				m_product_document.prod_doc_code,
				m_product_document.prod_doc_name,
				m_product_document_category.prod_doc_ctgry_name,
				m_product_document_category.prod_doc_ctgry_desc,
				m_mitra_product_document_rule.cust_sub_type1,
				m_mitra_product_document_rule.cust_sub_type2,
				m_mitra_product_document_rule.cust_sub_type3,
				m_mitra_product_document_rule.cust_sub_type4,
				m_mitra_product_document_rule.cust_sub_type5,
				m_mitra_product_document_rule.cust_sub_type6,
				m_mitra_product_document_rule.cust_sub_type7,
				m_mitra_product_document_rule.cust_sub_type8,
				m_mitra_product_document_rule.mitra_prod_doc_rule_is_mandatory,
				m_mitra_product_document_rule.mitra_prod_doc_rule_is_many_img,
				m_mitra_product_document_rule.mitra_prod_doc_rule_status,
				m_mitra_product_document_rule.mitra_prod_doc_rule_created_by,
				m_mitra_product_document_rule.mitra_prod_doc_rule_updated_by,
				m_mitra_product_document_rule.mitra_prod_doc_rule_deleted_by,
				m_mitra_product_document_rule.mitra_prod_doc_rule_created_at,
				m_mitra_product_document_rule.mitra_prod_doc_rule_updated_at,
				m_mitra_product_document_rule.mitra_prod_doc_rule_deleted_at`).
		Joins("JOIN m_mitra_product on m_mitra_product_document_rule.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_document on m_mitra_product_document_rule.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_document_category on m_product_document.prod_doc_ctgry_id=m_product_document_category.prod_doc_ctgry_id").
		Order("m_mitra_product_document_rule.prod_doc_id asc").
		Find(&mitraproductdocumentrule).Error
	if err != nil {
		return &[]MitraProductDocumentRuleData{}, err
	}
	return &mitraproductdocumentrule, nil
}

func (p *MitraProductDocumentRule) FindAllMitraProductDocumentRuleStatus(db *gorm.DB, status uint64) (*[]MitraProductDocumentRuleData, error) {
	var err error
	mitraproductdocumentrule := []MitraProductDocumentRuleData{}
	err = db.Debug().Model(&MitraProductDocumentRuleData{}).Limit(100).
		Select(`m_mitra_product_document_rule.mitra_prod_doc_rule_id,
				m_mitra_product.mitra_prod_id,
				m_mitra_product.mitra_prod_name,
				m_product_document.prod_doc_id,
				m_product_document.prod_doc_code,
				m_product_document.prod_doc_name,
				m_product_document_category.prod_doc_ctgry_name,
				m_product_document_category.prod_doc_ctgry_desc,
				m_mitra_product_document_rule.cust_sub_type1,
				m_mitra_product_document_rule.cust_sub_type2,
				m_mitra_product_document_rule.cust_sub_type3,
				m_mitra_product_document_rule.cust_sub_type4,
				m_mitra_product_document_rule.cust_sub_type5,
				m_mitra_product_document_rule.cust_sub_type6,
				m_mitra_product_document_rule.cust_sub_type7,
				m_mitra_product_document_rule.cust_sub_type8,
				m_mitra_product_document_rule.mitra_prod_doc_rule_is_mandatory,
				m_mitra_product_document_rule.mitra_prod_doc_rule_is_many_img,
				m_mitra_product_document_rule.mitra_prod_doc_rule_status,
				m_mitra_product_document_rule.mitra_prod_doc_rule_created_by,
				m_mitra_product_document_rule.mitra_prod_doc_rule_updated_by,
				m_mitra_product_document_rule.mitra_prod_doc_rule_deleted_by,
				m_mitra_product_document_rule.mitra_prod_doc_rule_created_at,
				m_mitra_product_document_rule.mitra_prod_doc_rule_updated_at,
				m_mitra_product_document_rule.mitra_prod_doc_rule_deleted_at`).
		Joins("JOIN m_mitra_product on m_mitra_product_document_rule.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_document on m_mitra_product_document_rule.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_document_category on m_product_document.prod_doc_ctgry_id=m_product_document_category.prod_doc_ctgry_id").
		Where("mitra_prod_doc_rule_status = ?", status).
		Order("m_mitra_product_document_rule.prod_doc_id asc").
		Find(&mitraproductdocumentrule).Error
	if err != nil {
		return &[]MitraProductDocumentRuleData{}, err
	}
	return &mitraproductdocumentrule, nil
}

func (p *MitraProductDocumentRule) FindAllMitraProductDocumentRuleByMitraProduct(db *gorm.DB, mitrabranch uint64) (*[]MitraProductDocumentRuleData, error) {
	var err error
	mitraproductdocumentrule := []MitraProductDocumentRuleData{}
	err = db.Debug().Model(&MitraProductDocumentRuleData{}).Limit(100).
		Select(`m_mitra_product_document_rule.mitra_prod_doc_rule_id,
				m_mitra_product.mitra_prod_id,
				m_mitra_product.mitra_prod_name,
				m_product_document.prod_doc_id,
				m_product_document.prod_doc_code,
				m_product_document.prod_doc_name,
				m_product_document_category.prod_doc_ctgry_name,
				m_product_document_category.prod_doc_ctgry_desc,
				m_mitra_product_document_rule.cust_sub_type1,
				m_mitra_product_document_rule.cust_sub_type2,
				m_mitra_product_document_rule.cust_sub_type3,
				m_mitra_product_document_rule.cust_sub_type4,
				m_mitra_product_document_rule.cust_sub_type5,
				m_mitra_product_document_rule.cust_sub_type6,
				m_mitra_product_document_rule.cust_sub_type7,
				m_mitra_product_document_rule.cust_sub_type8,
				m_mitra_product_document_rule.mitra_prod_doc_rule_is_mandatory,
				m_mitra_product_document_rule.mitra_prod_doc_rule_is_many_img,
				m_mitra_product_document_rule.mitra_prod_doc_rule_status,
				m_mitra_product_document_rule.mitra_prod_doc_rule_created_by,
				m_mitra_product_document_rule.mitra_prod_doc_rule_updated_by,
				m_mitra_product_document_rule.mitra_prod_doc_rule_deleted_by,
				m_mitra_product_document_rule.mitra_prod_doc_rule_created_at,
				m_mitra_product_document_rule.mitra_prod_doc_rule_updated_at,
				m_mitra_product_document_rule.mitra_prod_doc_rule_deleted_at`).
		Joins("JOIN m_mitra_product on m_mitra_product_document_rule.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_document on m_mitra_product_document_rule.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_document_category on m_product_document.prod_doc_ctgry_id=m_product_document_category.prod_doc_ctgry_id").
		Where("m_mitra_product_document_rule.mitra_prod_id = ?", mitrabranch).
		Order("m_mitra_product_document_rule.prod_doc_id asc").
		Find(&mitraproductdocumentrule).Error
	if err != nil {
		return &[]MitraProductDocumentRuleData{}, err
	}
	return &mitraproductdocumentrule, nil
}

func (p *MitraProductDocumentRule) FindMitraProductDocumentRuleDataByID(db *gorm.DB, pid uint64) (*MitraProductDocumentRule, error) {
	var err error
	err = db.Debug().Model(&MitraProductDocumentRule{}).Where("mitra_prod_doc_rule_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraProductDocumentRule{}, err
	}
	return p, nil
}

func (p *MitraProductDocumentRule) FindMitraProductDocumentRuleByID(db *gorm.DB, pid uint64) (*MitraProductDocumentRuleData, error) {
	var err error
	mitraproductdocumentrule := MitraProductDocumentRuleData{}
	err = db.Debug().Model(&MitraProductDocumentRuleData{}).
		Select(`m_mitra_product_document_rule.mitra_prod_doc_rule_id,
				m_mitra_product.mitra_prod_id,
				m_mitra_product.mitra_prod_name,
				m_product_document.prod_doc_id,
				m_product_document.prod_doc_code,
				m_product_document.prod_doc_name,
				m_product_document_category.prod_doc_ctgry_name,
				m_product_document_category.prod_doc_ctgry_desc,
				m_mitra_product_document_rule.cust_sub_type1,
				m_mitra_product_document_rule.cust_sub_type2,
				m_mitra_product_document_rule.cust_sub_type3,
				m_mitra_product_document_rule.cust_sub_type4,
				m_mitra_product_document_rule.cust_sub_type5,
				m_mitra_product_document_rule.cust_sub_type6,
				m_mitra_product_document_rule.cust_sub_type7,
				m_mitra_product_document_rule.cust_sub_type8,
				m_mitra_product_document_rule.mitra_prod_doc_rule_is_mandatory,
				m_mitra_product_document_rule.mitra_prod_doc_rule_is_many_img,
				m_mitra_product_document_rule.mitra_prod_doc_rule_status,
				m_mitra_product_document_rule.mitra_prod_doc_rule_created_by,
				m_mitra_product_document_rule.mitra_prod_doc_rule_updated_by,
				m_mitra_product_document_rule.mitra_prod_doc_rule_deleted_by,
				m_mitra_product_document_rule.mitra_prod_doc_rule_created_at,
				m_mitra_product_document_rule.mitra_prod_doc_rule_updated_at,
				m_mitra_product_document_rule.mitra_prod_doc_rule_deleted_at`).
		Joins("JOIN m_mitra_product on m_mitra_product_document_rule.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_document on m_mitra_product_document_rule.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_document_category on m_product_document.prod_doc_ctgry_id=m_product_document_category.prod_doc_ctgry_id").
		Where("mitra_prod_doc_rule_id = ?", pid).
		Order("m_mitra_product_document_rule.prod_doc_id asc").
		Take(&mitraproductdocumentrule).Error
	if err != nil {
		return &MitraProductDocumentRuleData{}, err
	}
	return &mitraproductdocumentrule, nil
}

func (p *MitraProductDocumentRule) FindMitraProductDocumentRuleStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraProductDocumentRuleData, error) {
	var err error
	mitraproductdocumentrule := MitraProductDocumentRuleData{}
	err = db.Debug().Model(&MitraProductDocumentRuleData{}).
		Select(`m_mitra_product_document_rule.mitra_prod_doc_rule_id,
				m_mitra_product.mitra_prod_id,
				m_mitra_product.mitra_prod_name,
				m_product_document.prod_doc_id,
				m_product_document.prod_doc_code,
				m_product_document.prod_doc_name,
				m_product_document_category.prod_doc_ctgry_name,
				m_product_document_category.prod_doc_ctgry_desc,
				m_mitra_product_document_rule.cust_sub_type1,
				m_mitra_product_document_rule.cust_sub_type2,
				m_mitra_product_document_rule.cust_sub_type3,
				m_mitra_product_document_rule.cust_sub_type4,
				m_mitra_product_document_rule.cust_sub_type5,
				m_mitra_product_document_rule.cust_sub_type6,
				m_mitra_product_document_rule.cust_sub_type7,
				m_mitra_product_document_rule.cust_sub_type8,
				m_mitra_product_document_rule.mitra_prod_doc_rule_is_mandatory,
				m_mitra_product_document_rule.mitra_prod_doc_rule_is_many_img,
				m_mitra_product_document_rule.mitra_prod_doc_rule_status,
				m_mitra_product_document_rule.mitra_prod_doc_rule_created_by,
				m_mitra_product_document_rule.mitra_prod_doc_rule_updated_by,
				m_mitra_product_document_rule.mitra_prod_doc_rule_deleted_by,
				m_mitra_product_document_rule.mitra_prod_doc_rule_created_at,
				m_mitra_product_document_rule.mitra_prod_doc_rule_updated_at,
				m_mitra_product_document_rule.mitra_prod_doc_rule_deleted_at`).
		Joins("JOIN m_mitra_product on m_mitra_product_document_rule.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_document on m_mitra_product_document_rule.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_document_category on m_product_document.prod_doc_ctgry_id=m_product_document_category.prod_doc_ctgry_id").
		Where("mitra_prod_doc_rule_id = ? AND mitra_prod_doc_rule_status = ?", pid, status).
		Order("m_mitra_product_document_rule.prod_doc_id asc").
		Take(&mitraproductdocumentrule).Error
	if err != nil {
		return &MitraProductDocumentRuleData{}, err
	}
	return &mitraproductdocumentrule, nil
}

func (p *MitraProductDocumentRule) FindMitraProductDocumentRuleByMitraProductByID(db *gorm.DB, pid uint64, mitrabranch uint64) (*MitraProductDocumentRuleData, error) {
	var err error
	mitraproductdocumentrule := MitraProductDocumentRuleData{}
	err = db.Debug().Model(&MitraProductDocumentRuleData{}).
		Select(`m_mitra_product_document_rule.mitra_prod_doc_rule_id,
				m_mitra_product.mitra_prod_id,
				m_mitra_product.mitra_prod_name,
				m_product_document.prod_doc_id,
				m_product_document.prod_doc_code,
				m_product_document.prod_doc_name,
				m_product_document_category.prod_doc_ctgry_name,
				m_product_document_category.prod_doc_ctgry_desc,
				m_mitra_product_document_rule.cust_sub_type1,
				m_mitra_product_document_rule.cust_sub_type2,
				m_mitra_product_document_rule.cust_sub_type3,
				m_mitra_product_document_rule.cust_sub_type4,
				m_mitra_product_document_rule.cust_sub_type5,
				m_mitra_product_document_rule.cust_sub_type6,
				m_mitra_product_document_rule.cust_sub_type7,
				m_mitra_product_document_rule.cust_sub_type8,
				m_mitra_product_document_rule.mitra_prod_doc_rule_is_mandatory,
				m_mitra_product_document_rule.mitra_prod_doc_rule_is_many_img,
				m_mitra_product_document_rule.mitra_prod_doc_rule_status,
				m_mitra_product_document_rule.mitra_prod_doc_rule_created_by,
				m_mitra_product_document_rule.mitra_prod_doc_rule_updated_by,
				m_mitra_product_document_rule.mitra_prod_doc_rule_deleted_by,
				m_mitra_product_document_rule.mitra_prod_doc_rule_created_at,
				m_mitra_product_document_rule.mitra_prod_doc_rule_updated_at,
				m_mitra_product_document_rule.mitra_prod_doc_rule_deleted_at`).
		Joins("JOIN m_mitra_product on m_mitra_product_document_rule.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_document on m_mitra_product_document_rule.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_document_category on m_product_document.prod_doc_ctgry_id=m_product_document_category.prod_doc_ctgry_id").
		Where("mitra_prod_doc_rule_id = ? AND mitra_prod_id = ?", pid, mitrabranch).
		Order("m_mitra_product_document_rule.prod_doc_id asc").
		Take(&mitraproductdocumentrule).Error
	if err != nil {
		return &MitraProductDocumentRuleData{}, err
	}
	return &mitraproductdocumentrule, nil
}

func (p *MitraProductDocumentRule) UpdateMitraProductDocumentRule(db *gorm.DB, pid uint64) (*MitraProductDocumentRule, error) {
	var err error
	db = db.Debug().Model(&MitraProductDocumentRule{}).Where("mitra_prod_doc_rule_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_prod_id":                    p.MitraProductID,
			"cust_sub_type1":                   p.CustomerSubType1,
			"cust_sub_type2":                   p.CustomerSubType2,
			"cust_sub_type3":                   p.CustomerSubType3,
			"cust_sub_type4":                   p.CustomerSubType4,
			"cust_sub_type5":                   p.CustomerSubType5,
			"cust_sub_type6":                   p.CustomerSubType6,
			"cust_sub_type7":                   p.CustomerSubType7,
			"cust_sub_type8":                   p.CustomerSubType8,
			"mitra_prod_doc_rule_is_mandatory": p.MitraProductDocumentRuleIsMandatory,
			"mitra_prod_doc_rule_is_many_img":  p.MitraProductDocumentRuleIsManyImage,
			"mitra_prod_doc_rule_status":       p.MitraProductDocumentRuleStatus,
			"mitra_prod_doc_rule_updated_by":   p.MitraProductDocumentRuleUpdatedBy,
			"mitra_prod_doc_rule_updated_at":   time.Now(),
		},
	)
	err = db.Debug().Model(&MitraProductDocumentRule{}).Where("mitra_prod_doc_rule_id = ?", pid).Error
	if err != nil {
		return &MitraProductDocumentRule{}, err
	}
	return p, nil
}

func (p *MitraProductDocumentRule) UpdateMitraProductDocumentRuleStatus(db *gorm.DB, pid uint64) (*MitraProductDocumentRule, error) {
	var err error
	db = db.Debug().Model(&MitraProductDocumentRule{}).Where("mitra_prod_doc_rule_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_prod_doc_rule_status":     p.MitraProductDocumentRuleStatus,
			"mitra_prod_doc_rule_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraProductDocumentRule{}).Where("mitra_prod_doc_rule_id = ?", pid).Error
	if err != nil {
		return &MitraProductDocumentRule{}, err
	}
	return p, nil
}

func (p *MitraProductDocumentRule) UpdateMitraProductDocumentRuleStatusOnly(db *gorm.DB, pid uint64) (*MitraProductDocumentRule, error) {
	var err error
	db = db.Debug().Model(&MitraProductDocumentRule{}).Where("mitra_prod_doc_rule_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_prod_doc_rule_status": p.MitraProductDocumentRuleStatus,
		},
	)
	err = db.Debug().Model(&MitraProductDocumentRule{}).Where("mitra_prod_doc_rule_id = ?", pid).Error
	if err != nil {
		return &MitraProductDocumentRule{}, err
	}
	return p, nil
}

func (p *MitraProductDocumentRule) UpdateMitraProductDocumentRuleStatusDelete(db *gorm.DB, pid uint64) (*MitraProductDocumentRule, error) {
	var err error
	db = db.Debug().Model(&MitraProductDocumentRule{}).Where("mitra_prod_doc_rule_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_prod_doc_rule_status":     3,
			"mitra_prod_doc_rule_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraProductDocumentRule{}).Where("mitra_prod_doc_rule_id = ?", pid).Error
	if err != nil {
		return &MitraProductDocumentRule{}, err
	}
	return p, nil
}

func (p *MitraProductDocumentRule) DeleteMitraProductDocumentRule(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraProductDocumentRule{}).Where("mitra_prod_doc_rule_id = ?", pid).Take(&MitraProductDocumentRule{}).Delete(&MitraProductDocumentRule{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraProductDocumentRule not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

func (p *MitraProductDocumentRule) FindMitraProductDocumentRuleByMitraProductID(db *gorm.DB, mitra uint64) ([]MitraProductDocumentRuleData, error) {
	var err error
	mitracoverageprovince := []MitraProductDocumentRuleData{}
	err = db.Debug().Model(&MitraProductDocumentRuleData{}).Limit(100).
		Select(`m_mitra_product_document_rule.mitra_prod_doc_rule_id,
				m_mitra_product.mitra_prod_id,
				m_mitra_product.mitra_prod_name,
				m_product_document.prod_doc_id,
				m_product_document.prod_doc_code,
				m_product_document.prod_doc_name,
				m_product_document_category.prod_doc_ctgry_name,
				m_product_document_category.prod_doc_ctgry_desc,
				m_mitra_product_document_rule.cust_sub_type1,
				m_mitra_product_document_rule.cust_sub_type2,
				m_mitra_product_document_rule.cust_sub_type3,
				m_mitra_product_document_rule.cust_sub_type4,
				m_mitra_product_document_rule.cust_sub_type5,
				m_mitra_product_document_rule.cust_sub_type6,
				m_mitra_product_document_rule.cust_sub_type7,
				m_mitra_product_document_rule.cust_sub_type8,
				m_mitra_product_document_rule.mitra_prod_doc_rule_is_mandatory,
				m_mitra_product_document_rule.mitra_prod_doc_rule_is_many_img,
				m_mitra_product_document_rule.mitra_prod_doc_rule_status,
				m_mitra_product_document_rule.mitra_prod_doc_rule_created_by,
				m_mitra_product_document_rule.mitra_prod_doc_rule_updated_by,
				m_mitra_product_document_rule.mitra_prod_doc_rule_deleted_by,
				m_mitra_product_document_rule.mitra_prod_doc_rule_created_at,
				m_mitra_product_document_rule.mitra_prod_doc_rule_updated_at,
				m_mitra_product_document_rule.mitra_prod_doc_rule_deleted_at`).
		Joins("JOIN m_mitra_product on m_mitra_product_document_rule.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_document on m_mitra_product_document_rule.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_document_category on m_product_document.prod_doc_ctgry_id=m_product_document_category.prod_doc_ctgry_id").
		Where("m_mitra_product_document_rule.mitra_prod_id = ?", mitra).
		Order("m_mitra_product_document_rule.prod_doc_id asc").
		Find(&mitracoverageprovince).Error
	if err != nil {
		return []MitraProductDocumentRuleData{}, err
	}
	return mitracoverageprovince, nil
}

func (p *MitraProductDocumentRule) FindMitraProductDocumentRuleStatusByMitraProductID(db *gorm.DB, mitra uint64, status uint64) ([]MitraProductDocumentRuleData, error) {
	var err error
	mitracoverageprovince := []MitraProductDocumentRuleData{}
	err = db.Debug().Model(&MitraProductDocumentRuleData{}).Limit(100).
		Select(`m_mitra_product_document_rule.mitra_prod_doc_rule_id,
				m_mitra_product.mitra_prod_id,
				m_mitra_product.mitra_prod_name,
				m_product_document.prod_doc_id,
				m_product_document.prod_doc_code,
				m_product_document.prod_doc_name,
				m_product_document_category.prod_doc_ctgry_name,
				m_product_document_category.prod_doc_ctgry_desc,
				m_mitra_product_document_rule.cust_sub_type1,
				m_mitra_product_document_rule.cust_sub_type2,
				m_mitra_product_document_rule.cust_sub_type3,
				m_mitra_product_document_rule.cust_sub_type4,
				m_mitra_product_document_rule.cust_sub_type5,
				m_mitra_product_document_rule.cust_sub_type6,
				m_mitra_product_document_rule.cust_sub_type7,
				m_mitra_product_document_rule.cust_sub_type8,
				m_mitra_product_document_rule.mitra_prod_doc_rule_is_mandatory,
				m_mitra_product_document_rule.mitra_prod_doc_rule_is_many_img,
				m_mitra_product_document_rule.mitra_prod_doc_rule_status,
				m_mitra_product_document_rule.mitra_prod_doc_rule_created_by,
				m_mitra_product_document_rule.mitra_prod_doc_rule_updated_by,
				m_mitra_product_document_rule.mitra_prod_doc_rule_deleted_by,
				m_mitra_product_document_rule.mitra_prod_doc_rule_created_at,
				m_mitra_product_document_rule.mitra_prod_doc_rule_updated_at,
				m_mitra_product_document_rule.mitra_prod_doc_rule_deleted_at`).
		Joins("JOIN m_mitra_product on m_mitra_product_document_rule.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_document on m_mitra_product_document_rule.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_document_category on m_product_document.prod_doc_ctgry_id=m_product_document_category.prod_doc_ctgry_id").
		Where("m_mitra_product_document_rule.mitra_prod_id = ? AND m_mitra_product_document_rule.mitra_prod_doc_rule_status = ?", mitra, status).
		Order("m_mitra_product_document_rule.prod_doc_id asc").
		Find(&mitracoverageprovince).Error
	if err != nil {
		return []MitraProductDocumentRuleData{}, err
	}
	return mitracoverageprovince, nil
}

func (p *MitraProductDocumentRule) FindMitraProductDocumentRuleByProduct(db *gorm.DB, product uint64, custsubtype uint64) ([]MitraProductDocumentRuleData, error) {
	var err error
	SQL := fmt.Sprintf("m_mitra_product.mitra_prod_id = ? AND m_mitra_product_document_rule.cust_sub_type%d = true", custsubtype)
	mitraproductcost := []MitraProductDocumentRuleData{}
	err = db.Debug().Model(&[]MitraProductDocumentRuleData{}).Limit(100).
		Select(`m_mitra_product_document_rule.mitra_prod_doc_rule_id,
				m_mitra_product.mitra_prod_id,
				m_mitra_product.mitra_prod_name,
				m_product_document.prod_doc_id,
				m_product_document.prod_doc_code,
				m_product_document.prod_doc_name,
				m_product_document_category.prod_doc_ctgry_name,
				m_product_document_category.prod_doc_ctgry_desc,
				m_mitra_product_document_rule.cust_sub_type1,
				m_mitra_product_document_rule.cust_sub_type2,
				m_mitra_product_document_rule.cust_sub_type3,
				m_mitra_product_document_rule.cust_sub_type4,
				m_mitra_product_document_rule.cust_sub_type5,
				m_mitra_product_document_rule.cust_sub_type6,
				m_mitra_product_document_rule.cust_sub_type7,
				m_mitra_product_document_rule.cust_sub_type8,
				m_mitra_product_document_rule.mitra_prod_doc_rule_is_mandatory,
				m_mitra_product_document_rule.mitra_prod_doc_rule_is_many_img,
				m_mitra_product_document_rule.mitra_prod_doc_rule_status,
				m_mitra_product_document_rule.mitra_prod_doc_rule_created_by,
				m_mitra_product_document_rule.mitra_prod_doc_rule_updated_by,
				m_mitra_product_document_rule.mitra_prod_doc_rule_deleted_by,
				m_mitra_product_document_rule.mitra_prod_doc_rule_created_at,
				m_mitra_product_document_rule.mitra_prod_doc_rule_updated_at,
				m_mitra_product_document_rule.mitra_prod_doc_rule_deleted_at`).
		Joins("JOIN m_mitra_product on m_mitra_product_document_rule.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_document on m_mitra_product_document_rule.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_document_category on m_product_document.prod_doc_ctgry_id=m_product_document_category.prod_doc_ctgry_id").
		Where(SQL, product).
		Order("m_mitra_product_document_rule.prod_doc_id asc").
		Find(&mitraproductcost).Error
	if err != nil {
		return []MitraProductDocumentRuleData{}, err
	}
	return mitraproductcost, nil
}
