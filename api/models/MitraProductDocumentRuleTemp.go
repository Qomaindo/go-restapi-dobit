package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type MitraProductDocumentRuleTemp struct {
	MitraProductDocumentRuleTempID           *uint64   `gorm:"column:mitra_prod_doc_rule_temp_id;primary_key;" json:"mitra_prod_doc_rule_temp_id"`
	MitraProductDocumentRuleID               uint64    `gorm:"column:mitra_prod_doc_rule_id" json:"mitra_prod_doc_rule_id"`
	MitraProductTempID                       uint64    `gorm:"column:mitra_prod_temp_id" json:"mitra_prod_temp_id"`
	ProductDocumentTempID                    uint64    `gorm:"column:prod_doc_temp_id" json:"prod_doc_temp_id"`
	CustomerSubType1Temp                     bool      `gorm:"column:cust_sub_type1_temp" json:"cust_sub_type1_temp"`
	CustomerSubType2Temp                     bool      `gorm:"column:cust_sub_type2_temp" json:"cust_sub_type2_temp"`
	CustomerSubType3Temp                     bool      `gorm:"column:cust_sub_type3_temp" json:"cust_sub_type3_temp"`
	CustomerSubType4Temp                     bool      `gorm:"column:cust_sub_type4_temp" json:"cust_sub_type4_temp"`
	CustomerSubType5Temp                     bool      `gorm:"column:cust_sub_type5_temp" json:"cust_sub_type5_temp"`
	CustomerSubType6Temp                     bool      `gorm:"column:cust_sub_type6_temp" json:"cust_sub_type6_temp"`
	CustomerSubType7Temp                     bool      `gorm:"column:cust_sub_type7_temp" json:"cust_sub_type7_temp"`
	CustomerSubType8Temp                     bool      `gorm:"column:cust_sub_type8_temp" json:"cust_sub_type8_temp"`
	MitraProductDocumentRuleTempIsMandatory  bool      `gorm:"column:mitra_prod_doc_rule_temp_is_mandatory" json:"mitra_prod_doc_rule_temp_is_mandatory"`
	MitraProductDocumentRuleTempIsManyImage  bool      `gorm:"column:mitra_prod_doc_rule_temp_is_many_img" json:"mitra_prod_doc_rule_temp_is_many_img"`
	MitraProductDocumentRuleTempReason       string    `gorm:"column:mitra_prod_doc_rule_temp_reason" json:"mitra_prod_doc_rule_temp_reason"`
	MitraProductDocumentRuleTempActionBefore uint64    `gorm:"column:mitra_prod_doc_rule_temp_action_before" json:"mitra_prod_doc_rule_temp_action_before"`
	MitraProductDocumentRuleTempStatus       uint64    `gorm:"column:mitra_prod_doc_rule_temp_status;size:25" json:"mitra_prod_doc_rule_temp_status"`
	MitraProductDocumentRuleTempCreatedBy    string    `gorm:"column:mitra_prod_doc_rule_temp_created_by;size:125" json:"mitra_prod_doc_rule_temp_created_by"`
	MitraProductDocumentRuleTempCreatedAt    time.Time `gorm:"column:mitra_prod_doc_rule_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_doc_rule_temp_created_at"`
}

type MitraProductDocumentRuleTempPend struct {
	MitraProductDocumentRuleTempID           *uint64   `gorm:"column:mitra_prod_doc_rule_temp_id;primary_key;" json:"mitra_prod_doc_rule_temp_id"`
	MitraProductDocumentRuleID               *uint64   `gorm:"column:mitra_prod_doc_rule_id" json:"mitra_prod_doc_rule_id"`
	MitraProductTempID                       uint64    `gorm:"column:mitra_prod_temp_id" json:"mitra_prod_temp_id"`
	ProductDocumentTempID                    uint64    `gorm:"column:prod_doc_temp_id" json:"prod_doc_temp_id"`
	CustomerSubType1Temp                     bool      `gorm:"column:cust_sub_type1_temp" json:"cust_sub_type1_temp"`
	CustomerSubType2Temp                     bool      `gorm:"column:cust_sub_type2_temp" json:"cust_sub_type2_temp"`
	CustomerSubType3Temp                     bool      `gorm:"column:cust_sub_type3_temp" json:"cust_sub_type3_temp"`
	CustomerSubType4Temp                     bool      `gorm:"column:cust_sub_type4_temp" json:"cust_sub_type4_temp"`
	CustomerSubType5Temp                     bool      `gorm:"column:cust_sub_type5_temp" json:"cust_sub_type5_temp"`
	CustomerSubType6Temp                     bool      `gorm:"column:cust_sub_type6_temp" json:"cust_sub_type6_temp"`
	CustomerSubType7Temp                     bool      `gorm:"column:cust_sub_type7_temp" json:"cust_sub_type7_temp"`
	CustomerSubType8Temp                     bool      `gorm:"column:cust_sub_type8_temp" json:"cust_sub_type8_temp"`
	MitraProductDocumentRuleTempIsMandatory  bool      `gorm:"column:mitra_prod_doc_rule_temp_is_mandatory" json:"mitra_prod_doc_rule_temp_is_mandatory"`
	MitraProductDocumentRuleTempIsManyImage  bool      `gorm:"column:mitra_prod_doc_rule_temp_is_many_img" json:"mitra_prod_doc_rule_temp_is_many_img"`
	MitraProductDocumentRuleTempReason       string    `gorm:"column:mitra_prod_doc_rule_temp_reason" json:"mitra_prod_doc_rule_temp_reason"`
	MitraProductDocumentRuleTempActionBefore uint64    `gorm:"column:mitra_prod_doc_rule_temp_action_before" json:"mitra_prod_doc_rule_temp_action_before"`
	MitraProductDocumentRuleTempStatus       uint64    `gorm:"column:mitra_prod_doc_rule_temp_status;size:25" json:"mitra_prod_doc_rule_temp_status"`
	MitraProductDocumentRuleTempCreatedBy    string    `gorm:"column:mitra_prod_doc_rule_temp_created_by;size:125" json:"mitra_prod_doc_rule_temp_created_by"`
	MitraProductDocumentRuleTempCreatedAt    time.Time `gorm:"column:mitra_prod_doc_rule_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_doc_rule_temp_created_at"`
}

type MitraProductDocumentRuleTempPendData struct {
	MitraProductDocumentRuleTempID           uint64    `gorm:"column:mitra_prod_doc_rule_temp_id;primary_key;" json:"mitra_prod_doc_rule_temp_id"`
	MitraProductDocumentRuleID               uint64    `gorm:"column:mitra_prod_doc_rule_id" json:"mitra_prod_doc_rule_id"`
	MitraProductTempID                       uint64    `gorm:"column:mitra_prod_temp_id" json:"mitra_prod_temp_id"`
	MitraProductTempName                     string    `gorm:"column:mitra_prod_temp_name" json:"mitra_prod_temp_name"`
	ProductDocumentTempID                    uint64    `gorm:"column:prod_doc_temp_id" json:"prod_doc_temp_id"`
	ProductDocumentTempCode                  string    `gorm:"column:prod_doc_temp_code" json:"prod_doc_temp_code"`
	ProductDocumentTempName                  string    `gorm:"column:prod_doc_temp_name" json:"prod_doc_temp_name"`
	ProductDocumentCategoryTempName          string    `gorm:"column:prod_doc_ctgry_temp_name" json:"prod_doc_ctgry_temp_name"`
	ProductDocumentCategoryTempDesc          string    `gorm:"column:prod_doc_ctgry_temp_desc" json:"prod_doc_ctgry_temp_desc"`
	CustomerSubType1Temp                     bool      `gorm:"column:cust_sub_type1_temp" json:"cust_sub_type1_temp"`
	CustomerSubType2Temp                     bool      `gorm:"column:cust_sub_type2_temp" json:"cust_sub_type2_temp"`
	CustomerSubType3Temp                     bool      `gorm:"column:cust_sub_type3_temp" json:"cust_sub_type3_temp"`
	CustomerSubType4Temp                     bool      `gorm:"column:cust_sub_type4_temp" json:"cust_sub_type4_temp"`
	CustomerSubType5Temp                     bool      `gorm:"column:cust_sub_type5_temp" json:"cust_sub_type5_temp"`
	CustomerSubType6Temp                     bool      `gorm:"column:cust_sub_type6_temp" json:"cust_sub_type6_temp"`
	CustomerSubType7Temp                     bool      `gorm:"column:cust_sub_type7_temp" json:"cust_sub_type7_temp"`
	CustomerSubType8Temp                     bool      `gorm:"column:cust_sub_type8_temp" json:"cust_sub_type8_temp"`
	MitraProductDocumentRuleTempIsMandatory  bool      `gorm:"column:mitra_prod_doc_rule_temp_is_mandatory" json:"mitra_prod_doc_rule_temp_is_mandatory"`
	MitraProductDocumentRuleTempIsManyImage  bool      `gorm:"column:mitra_prod_doc_rule_temp_is_many_img" json:"mitra_prod_doc_rule_temp_is_many_img"`
	MitraProductDocumentRuleTempReason       string    `gorm:"column:mitra_prod_doc_rule_temp_reason" json:"mitra_prod_doc_rule_temp_reason"`
	MitraProductDocumentRuleTempActionBefore uint64    `gorm:"column:mitra_prod_doc_rule_temp_action_before" json:"mitra_prod_doc_rule_temp_action_before"`
	MitraProductDocumentRuleTempStatus       uint64    `gorm:"column:mitra_prod_doc_rule_temp_status;size:25" json:"mitra_prod_doc_rule_temp_status"`
	MitraProductDocumentRuleTempCreatedBy    string    `gorm:"column:mitra_prod_doc_rule_temp_created_by;size:125" json:"mitra_prod_doc_rule_temp_created_by"`
	MitraProductDocumentRuleTempCreatedAt    time.Time `gorm:"column:mitra_prod_doc_rule_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_doc_rule_temp_created_at"`
}

type MitraProductDocumentRuleTempData struct {
	MitraProductDocumentRuleTempID           uint64     `gorm:"column:mitra_prod_doc_rule_temp_id" json:"mitra_prod_doc_rule_temp_id"`
	MitraProductTempID                       uint64     `gorm:"column:mitra_prod_temp_id" json:"mitra_prod_temp_id"`
	MitraProductTempName                     string     `gorm:"column:mitra_prod_temp_name" json:"mitra_prod_temp_name"`
	ProductDocumentTempID                    uint64     `gorm:"column:prod_doc_temp_id" json:"prod_doc_temp_id"`
	ProductDocumentTempCode                  string     `gorm:"column:prod_doc_temp_code" json:"prod_doc_temp_code"`
	ProductDocumentTempName                  string     `gorm:"column:prod_doc_temp_name" json:"prod_doc_temp_name"`
	ProductDocumentCategoryTempName          string     `gorm:"column:prod_doc_ctgry_temp_name" json:"prod_doc_ctgry_temp_name"`
	ProductDocumentCategoryTempDesc          string     `gorm:"column:prod_doc_ctgry_temp_desc" json:"prod_doc_ctgry_temp_desc"`
	CustomerSubType1Temp                     bool       `gorm:"column:cust_sub_type1_temp" json:"cust_sub_type1_temp"`
	CustomerSubType2Temp                     bool       `gorm:"column:cust_sub_type2_temp" json:"cust_sub_type2_temp"`
	CustomerSubType3Temp                     bool       `gorm:"column:cust_sub_type3_temp" json:"cust_sub_type3_temp"`
	CustomerSubType4Temp                     bool       `gorm:"column:cust_sub_type4_temp" json:"cust_sub_type4_temp"`
	CustomerSubType5Temp                     bool       `gorm:"column:cust_sub_type5_temp" json:"cust_sub_type5_temp"`
	CustomerSubType6Temp                     bool       `gorm:"column:cust_sub_type6_temp" json:"cust_sub_type6_temp"`
	CustomerSubType7Temp                     bool       `gorm:"column:cust_sub_type7_temp" json:"cust_sub_type7_temp"`
	CustomerSubType8Temp                     bool       `gorm:"column:cust_sub_type8_temp" json:"cust_sub_type8_temp"`
	MitraProductDocumentRuleTempIsMandatory  bool       `gorm:"column:mitra_prod_doc_rule_temp_is_mandatory" json:"mitra_prod_doc_rule_temp_is_mandatory"`
	MitraProductDocumentRuleTempIsManyImage  bool       `gorm:"column:mitra_prod_doc_rule_temp_is_many_img" json:"mitra_prod_doc_rule_temp_is_many_img"`
	MitraProductDocumentRuleTempReason       string     `gorm:"column:mitra_prod_doc_rule_temp_reason" json:"mitra_prod_doc_rule_temp_reason"`
	MitraProductDocumentRuleTempActionBefore uint64     `gorm:"column:mitra_prod_doc_rule_temp_action_before" json:"mitra_prod_doc_rule_temp_action_before"`
	MitraProductDocumentRuleTempStatus       uint64     `gorm:"column:mitra_prod_doc_rule_temp_status;size:25" json:"mitra_prod_doc_rule_temp_status"`
	MitraProductDocumentRuleTempCreatedBy    string     `gorm:"column:mitra_prod_doc_rule_temp_created_by;size:125" json:"mitra_prod_doc_rule_temp_created_by"`
	MitraProductDocumentRuleTempCreatedAt    time.Time  `gorm:"column:mitra_prod_doc_rule_temp_created_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_doc_rule_temp_created_at"`
	MitraProductDocumentRuleID               uint64     `gorm:"column:mitra_prod_doc_rule_id" json:"mitra_prod_doc_rule_id"`
	MitraProductID                           uint64     `gorm:"column:mitra_prod_id" json:"mitra_prod_id"`
	MitraProductName                         string     `gorm:"column:mitra_prod_name" json:"mitra_prod_name"`
	ProductDocumentID                        uint64     `gorm:"column:prod_doc_id" json:"prod_doc_id"`
	ProductDocumentCode                      string     `gorm:"column:prod_doc_code" json:"prod_doc_code"`
	ProductDocumentName                      string     `gorm:"column:prod_doc_name" json:"prod_doc_name"`
	ProductDocumentCategoryName              string     `gorm:"column:prod_doc_ctgry_name" json:"prod_doc_ctgry_name"`
	ProductDocumentCategoryDesc              string     `gorm:"column:prod_doc_ctgry_desc" json:"prod_doc_ctgry_desc"`
	CustomerSubType1                         bool       `gorm:"column:cust_sub_type1" json:"cust_sub_type1"`
	CustomerSubType2                         bool       `gorm:"column:cust_sub_type2" json:"cust_sub_type2"`
	CustomerSubType3                         bool       `gorm:"column:cust_sub_type3" json:"cust_sub_type3"`
	CustomerSubType4                         bool       `gorm:"column:cust_sub_type4" json:"cust_sub_type4"`
	CustomerSubType5                         bool       `gorm:"column:cust_sub_type5" json:"cust_sub_type5"`
	CustomerSubType6                         bool       `gorm:"column:cust_sub_type6" json:"cust_sub_type6"`
	CustomerSubType7                         bool       `gorm:"column:cust_sub_type7" json:"cust_sub_type7"`
	CustomerSubType8                         bool       `gorm:"column:cust_sub_type8" json:"cust_sub_type8"`
	MitraProductDocumentRuleIsMandatory      bool       `gorm:"column:mitra_prod_doc_rule_is_mandatory" json:"mitra_prod_doc_rule_is_mandatory"`
	MitraProductDocumentRuleIsManyImage      bool       `gorm:"column:mitra_prod_doc_rule_is_many_img" json:"mitra_prod_doc_rule_is_many_img"`
	MitraProductDocumentRuleStatus           uint64     `gorm:"column:mitra_prod_doc_rule_status" json:"mitra_prod_doc_rule_status"`
	MitraProductDocumentRuleCreatedBy        string     `gorm:"column:mitra_prod_doc_rule_created_by;size:125" json:"mitra_prod_doc_rule_created_by"`
	MitraProductDocumentRuleCreatedAt        time.Time  `gorm:"column:mitra_prod_doc_rule_created_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_doc_rule_created_at"`
	MitraProductDocumentRuleUpdatedBy        string     `gorm:"column:mitra_prod_doc_rule_updated_by;size:125" json:"mitra_prod_doc_rule_updated_by"`
	MitraProductDocumentRuleUpdatedAt        *time.Time `gorm:"column:mitra_prod_doc_rule_updated_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_doc_rule_updated_at"`
	MitraProductDocumentRuleDeletedBy        string     `gorm:"column:mitra_prod_doc_rule_deleted_by;size:125" json:"mitra_prod_doc_rule_deleted_by"`
	MitraProductDocumentRuleDeletedAt        *time.Time `gorm:"column:mitra_prod_doc_rule_deleted_at;default:CURRENT_TIMESTAMP" json:"mitra_prod_doc_rule_deleted_at"`
}

type ResponseMitraProductDocumentRuleTemps struct {
	Status  int                                 `json:"status"`
	Message string                              `json:"message"`
	Data    *[]MitraProductDocumentRuleTempData `json:"data"`
}

type ResponseMitraProductDocumentRuleTempsPend struct {
	Status  int                                 `json:"status"`
	Message string                              `json:"message"`
	Data    *[]MitraProductDocumentRuleTempPend `json:"data"`
}

type ResponseMitraProductDocumentRuleTemp struct {
	Status  int                               `json:"status"`
	Message string                            `json:"message"`
	Data    *MitraProductDocumentRuleTempData `json:"data"`
}

type ResponseMitraProductDocumentRuleTempPend struct {
	Status  int                               `json:"status"`
	Message string                            `json:"message"`
	Data    *MitraProductDocumentRuleTempPend `json:"data"`
}

type ResponseMitraProductDocumentRuleTempIU struct {
	Status  int                           `json:"status"`
	Message string                        `json:"message"`
	Data    *MitraProductDocumentRuleTemp `json:"data"`
}

type ResponseMitraProductDocumentRuleTempCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (MitraProductDocumentRuleTemp) TableName() string {
	return "m_mitra_product_document_rule_temp"
}

func (MitraProductDocumentRuleTempPend) TableName() string {
	return "m_mitra_product_document_rule_temp"
}

func (MitraProductDocumentRuleTempPendData) TableName() string {
	return "m_mitra_product_document_rule_temp"
}

func (MitraProductDocumentRuleTempData) TableName() string {
	return "m_mitra_product_document_rule_temp"
}

func (p *MitraProductDocumentRuleTemp) Prepare() {
	p.MitraProductDocumentRuleTempID = nil
	p.MitraProductDocumentRuleID = p.MitraProductDocumentRuleID
	p.MitraProductTempID = p.MitraProductTempID
	p.ProductDocumentTempID = p.ProductDocumentTempID
	p.CustomerSubType1Temp = p.CustomerSubType1Temp
	p.CustomerSubType2Temp = p.CustomerSubType2Temp
	p.CustomerSubType3Temp = p.CustomerSubType3Temp
	p.CustomerSubType4Temp = p.CustomerSubType4Temp
	p.CustomerSubType5Temp = p.CustomerSubType5Temp
	p.CustomerSubType6Temp = p.CustomerSubType6Temp
	p.CustomerSubType7Temp = p.CustomerSubType7Temp
	p.CustomerSubType8Temp = p.CustomerSubType8Temp
	p.MitraProductDocumentRuleTempIsMandatory = p.MitraProductDocumentRuleTempIsMandatory
	p.MitraProductDocumentRuleTempIsManyImage = p.MitraProductDocumentRuleTempIsManyImage
	p.MitraProductDocumentRuleTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraProductDocumentRuleTempCreatedBy))
	p.MitraProductDocumentRuleTempReason = html.EscapeString(strings.TrimSpace(p.MitraProductDocumentRuleTempReason))
	p.MitraProductDocumentRuleTempStatus = p.MitraProductDocumentRuleTempStatus
	p.MitraProductDocumentRuleTempCreatedAt = time.Now()
}

func (p *MitraProductDocumentRuleTempPend) Prepare() {
	p.MitraProductDocumentRuleTempID = nil
	p.MitraProductDocumentRuleID = nil
	p.MitraProductTempID = p.MitraProductTempID
	p.ProductDocumentTempID = p.ProductDocumentTempID
	p.CustomerSubType1Temp = p.CustomerSubType1Temp
	p.CustomerSubType2Temp = p.CustomerSubType2Temp
	p.CustomerSubType3Temp = p.CustomerSubType3Temp
	p.CustomerSubType4Temp = p.CustomerSubType4Temp
	p.CustomerSubType5Temp = p.CustomerSubType5Temp
	p.CustomerSubType6Temp = p.CustomerSubType6Temp
	p.CustomerSubType7Temp = p.CustomerSubType7Temp
	p.CustomerSubType8Temp = p.CustomerSubType8Temp
	p.MitraProductDocumentRuleTempIsMandatory = p.MitraProductDocumentRuleTempIsMandatory
	p.MitraProductDocumentRuleTempIsManyImage = p.MitraProductDocumentRuleTempIsManyImage
	p.MitraProductDocumentRuleTempCreatedBy = html.EscapeString(strings.TrimSpace(p.MitraProductDocumentRuleTempCreatedBy))
	p.MitraProductDocumentRuleTempReason = html.EscapeString(strings.TrimSpace(p.MitraProductDocumentRuleTempReason))
	p.MitraProductDocumentRuleTempStatus = p.MitraProductDocumentRuleTempStatus
	p.MitraProductDocumentRuleTempCreatedAt = time.Now()
}

func (p *MitraProductDocumentRuleTemp) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		if p.MitraProductTempID == 0 {
			return errors.New("Required MitraProductDocumentRule Code")
		}
		if p.ProductDocumentTempID == 0 {
			return errors.New("Required MitraProductDocumentRule Code")
		}
		return nil
	}
}

func (p *MitraProductDocumentRuleTemp) SaveMitraProductDocumentRuleTemp(db *gorm.DB) (*MitraProductDocumentRuleTemp, error) {
	var err error
	err = db.Debug().Model(&MitraProductDocumentRuleTemp{}).Create(&p).Error
	if err != nil {
		return &MitraProductDocumentRuleTemp{}, err
	}
	return p, nil
}

func (p *MitraProductDocumentRuleTempPend) SaveMitraProductDocumentRuleTempPend(db *gorm.DB) (*MitraProductDocumentRuleTempPend, error) {
	var err error
	err = db.Debug().Model(&MitraProductDocumentRuleTempPend{}).Create(&p).Error
	if err != nil {
		return &MitraProductDocumentRuleTempPend{}, err
	}
	return p, nil
}

func (p *MitraProductDocumentRuleTemp) FindAllMitraProductDocumentRuleTemp(db *gorm.DB) (*[]MitraProductDocumentRuleTempPend, error) {
	var err error
	mitraproductdocumentruletemp := []MitraProductDocumentRuleTempPend{}
	err = db.Debug().Model(&MitraProductDocumentRuleTempPend{}).Limit(100).
		Select(`m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_id,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_id,
				m_mitra_product_document_rule_temp.mitra_prod_temp_id,
				m_mitra_product_document_rule_temp.prod_doc_temp_id,
				m_mitra_product_document_rule_temp.cust_sub_type1_temp,
				m_mitra_product_document_rule_temp.cust_sub_type2_temp,
				m_mitra_product_document_rule_temp.cust_sub_type3_temp,
				m_mitra_product_document_rule_temp.cust_sub_type4_temp,
				m_mitra_product_document_rule_temp.cust_sub_type5_temp,
				m_mitra_product_document_rule_temp.cust_sub_type6_temp,
				m_mitra_product_document_rule_temp.cust_sub_type7_temp,
				m_mitra_product_document_rule_temp.cust_sub_type8_temp,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_is_mandatory,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_is_many_img,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_status,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_reason,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_created_by,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_created_at`).
		Order("m_mitra_product_document_rule_temp.prod_doc_temp_id asc").
		Find(&mitraproductdocumentruletemp).Error
	if err != nil {
		return &[]MitraProductDocumentRuleTempPend{}, err
	}
	return &mitraproductdocumentruletemp, nil
}

func (p *MitraProductDocumentRuleTemp) FindAllMitraProductDocumentRuleTempStatusPendingActive(db *gorm.DB) (*[]MitraProductDocumentRuleTempPend, error) {
	var err error
	mitraproductdocumentruletemp := []MitraProductDocumentRuleTempPend{}
	err = db.Debug().Model(&MitraProductDocumentRuleTempPend{}).Limit(100).
		Select(`m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_id,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_id,
				m_mitra_product_document_rule_temp.mitra_prod_temp_id,
				m_mitra_product_document_rule_temp.prod_doc_temp_id,
				m_mitra_product_document_rule_temp.cust_sub_type1_temp,
				m_mitra_product_document_rule_temp.cust_sub_type2_temp,
				m_mitra_product_document_rule_temp.cust_sub_type3_temp,
				m_mitra_product_document_rule_temp.cust_sub_type4_temp,
				m_mitra_product_document_rule_temp.cust_sub_type5_temp,
				m_mitra_product_document_rule_temp.cust_sub_type6_temp,
				m_mitra_product_document_rule_temp.cust_sub_type7_temp,
				m_mitra_product_document_rule_temp.cust_sub_type8_temp,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_is_mandatory,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_is_many_img,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_status,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_reason,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_created_by,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_created_at`).
		Order("m_mitra_product_document_rule_temp.prod_doc_temp_id asc").
		Where("mitra_prod_doc_rule_temp_status = ?", 11).
		Find(&mitraproductdocumentruletemp).Error
	if err != nil {
		return &[]MitraProductDocumentRuleTempPend{}, err
	}
	return &mitraproductdocumentruletemp, nil
}

func (p *MitraProductDocumentRuleTemp) FindAllMitraProductDocumentRuleTempStatus(db *gorm.DB, status uint64) (*[]MitraProductDocumentRuleTempData, error) {
	var err error
	mitraproductdocumentruletemp := []MitraProductDocumentRuleTempData{}
	err = db.Debug().Model(&MitraProductDocumentRuleTempData{}).Limit(100).
		Select(`m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_id,
				m_mitra_product_temp.mitra_prod_temp_id,
				m_mitra_product_temp.mitra_prod_temp_name,
				m_product_document_temp.prod_doc_id as prod_doc_temp_id,
				m_product_document_temp.prod_doc_code as prod_doc_temp_code,
				m_product_document_temp.prod_doc_name as prod_doc_temp_name,
				m_product_document_category_temp.prod_doc_ctgry_name as prod_doc_ctgry_temp_name,
				m_product_document_category_temp.prod_doc_ctgry_desc as prod_doc_ctgry_temp_desc,
				m_mitra_product_document_rule_temp.cust_sub_type1_temp,
				m_mitra_product_document_rule_temp.cust_sub_type2_temp,
				m_mitra_product_document_rule_temp.cust_sub_type3_temp,
				m_mitra_product_document_rule_temp.cust_sub_type4_temp,
				m_mitra_product_document_rule_temp.cust_sub_type5_temp,
				m_mitra_product_document_rule_temp.cust_sub_type6_temp,
				m_mitra_product_document_rule_temp.cust_sub_type7_temp,
				m_mitra_product_document_rule_temp.cust_sub_type8_temp,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_is_mandatory,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_is_many_img,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_status,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_reason,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_created_by,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_created_at,
				m_mitra_product_document_rule.mitra_prod_doc_rule_id,
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
		Joins("JOIN m_mitra_product_temp on m_mitra_product_document_rule_temp.mitra_prod_temp_id=m_mitra_product_temp.mitra_prod_temp_id").
		Joins("JOIN m_product_document m_product_document_temp on m_mitra_product_document_rule_temp.prod_doc_temp_id=m_product_document_temp.prod_doc_id").
		Joins("JOIN m_product_document_category m_product_document_category_temp on m_product_document_temp.prod_doc_ctgry_id=m_product_document_category_temp.prod_doc_ctgry_id").
		Joins("JOIN m_mitra_product_document_rule on m_mitra_product_document_rule_temp.mitra_prod_doc_rule_id=m_mitra_product_document_rule.mitra_prod_doc_rule_id").
		Joins("JOIN m_mitra_product on m_mitra_product_document_rule.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_document on m_mitra_product_document_rule.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_document_category on m_product_document.prod_doc_ctgry_id=m_product_document_category.prod_doc_ctgry_id").
		Where("mitra_prod_doc_rule_temp_status = ?", status).
		Order("m_mitra_product_document_rule_temp.prod_doc_temp_id asc").
		Find(&mitraproductdocumentruletemp).Error
	if err != nil {
		return &[]MitraProductDocumentRuleTempData{}, err
	}
	return &mitraproductdocumentruletemp, nil
}

func (p *MitraProductDocumentRuleTemp) FindAllMitraProductDocumentRuleTempByMitraProduct(db *gorm.DB, mitrabranchtemp uint64) (*[]MitraProductDocumentRuleTempPend, error) {
	var err error
	mitraproductdocumentruletemp := []MitraProductDocumentRuleTempPend{}
	err = db.Debug().Model(&MitraProductDocumentRuleTempPend{}).Limit(100).
		Select(`m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_id,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_id,
				m_mitra_product_document_rule_temp.mitra_prod_temp_id,
				m_mitra_product_document_rule_temp.prod_doc_temp_id,
				m_mitra_product_document_rule_temp.cust_sub_type1_temp,
				m_mitra_product_document_rule_temp.cust_sub_type2_temp,
				m_mitra_product_document_rule_temp.cust_sub_type3_temp,
				m_mitra_product_document_rule_temp.cust_sub_type4_temp,
				m_mitra_product_document_rule_temp.cust_sub_type5_temp,
				m_mitra_product_document_rule_temp.cust_sub_type6_temp,
				m_mitra_product_document_rule_temp.cust_sub_type7_temp,
				m_mitra_product_document_rule_temp.cust_sub_type8_temp,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_is_mandatory,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_is_many_img,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_status,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_reason,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_created_by,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_created_at`).
		Where("mitra_prod_temp_id = ?", mitrabranchtemp).
		Order("m_mitra_product_document_rule_temp.prod_doc_temp_id asc").
		Find(&mitraproductdocumentruletemp).Error
	if err != nil {
		return &[]MitraProductDocumentRuleTempPend{}, err
	}
	return &mitraproductdocumentruletemp, nil
}

func (p *MitraProductDocumentRuleTemp) FindMitraProductDocumentRuleTempDataByID(db *gorm.DB, pid uint64) (*MitraProductDocumentRuleTemp, error) {
	var err error
	err = db.Debug().Model(&MitraProductDocumentRuleTemp{}).Where("mitra_prod_doc_rule_temp_id = ?", pid).Take(&p).Error
	if err != nil {
		return &MitraProductDocumentRuleTemp{}, err
	}
	return p, nil
}

func (p *MitraProductDocumentRuleTemp) FindMitraProductDocumentRuleTempByIDPendingActive(db *gorm.DB, pid uint64) (*MitraProductDocumentRuleTempPend, error) {
	var err error
	mitraproductdocumentruletemp := MitraProductDocumentRuleTempPend{}
	err = db.Debug().Model(&MitraProductDocumentRuleTempPend{}).Limit(100).
		Select(`m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_id,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_id,
				m_mitra_product_document_rule_temp.mitra_prod_temp_id,
				m_mitra_product_document_rule_temp.prod_doc_temp_id,
				m_mitra_product_document_rule_temp.cust_sub_type1_temp,
				m_mitra_product_document_rule_temp.cust_sub_type2_temp,
				m_mitra_product_document_rule_temp.cust_sub_type3_temp,
				m_mitra_product_document_rule_temp.cust_sub_type4_temp,
				m_mitra_product_document_rule_temp.cust_sub_type5_temp,
				m_mitra_product_document_rule_temp.cust_sub_type6_temp,
				m_mitra_product_document_rule_temp.cust_sub_type7_temp,
				m_mitra_product_document_rule_temp.cust_sub_type8_temp,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_is_mandatory,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_is_many_img,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_status,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_reason,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_created_by,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_created_at`).
		Where("mitra_prod_doc_rule_temp_id = ?", pid).
		Order("m_mitra_product_document_rule_temp.prod_doc_temp_id asc").
		Find(&mitraproductdocumentruletemp).Error
	if err != nil {
		return &MitraProductDocumentRuleTempPend{}, err
	}
	return &mitraproductdocumentruletemp, nil
}

func (p *MitraProductDocumentRuleTemp) FindMitraProductDocumentRuleTempStatusByIDPendingActive(db *gorm.DB, pid uint64, status uint64) (*MitraProductDocumentRuleTempPend, error) {
	var err error
	mitraproductdocumentruletemp := MitraProductDocumentRuleTempPend{}
	err = db.Debug().Model(&MitraProductDocumentRuleTempPend{}).Limit(100).
		Select(`m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_id,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_id,
				m_mitra_product_document_rule_temp.mitra_prod_temp_id,
				m_mitra_product_document_rule_temp.prod_doc_temp_id,
				m_mitra_product_document_rule_temp.cust_sub_type1_temp,
				m_mitra_product_document_rule_temp.cust_sub_type2_temp,
				m_mitra_product_document_rule_temp.cust_sub_type3_temp,
				m_mitra_product_document_rule_temp.cust_sub_type4_temp,
				m_mitra_product_document_rule_temp.cust_sub_type5_temp,
				m_mitra_product_document_rule_temp.cust_sub_type6_temp,
				m_mitra_product_document_rule_temp.cust_sub_type7_temp,
				m_mitra_product_document_rule_temp.cust_sub_type8_temp,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_is_mandatory,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_is_many_img,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_status,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_reason,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_created_by,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_created_at`).
		Where("mitra_prod_doc_rule_temp_id = ? AND mitra_prod_doc_rule_temp_status = ?", pid, status).
		Order("m_mitra_product_document_rule_temp.prod_doc_temp_id asc").
		Find(&mitraproductdocumentruletemp).Error
	if err != nil {
		return &MitraProductDocumentRuleTempPend{}, err
	}
	return &mitraproductdocumentruletemp, nil
}

func (p *MitraProductDocumentRuleTemp) FindMitraProductDocumentRuleTempByID(db *gorm.DB, pid uint64) (*MitraProductDocumentRuleTempData, error) {
	var err error
	mitraproductdocumentruletemp := MitraProductDocumentRuleTempData{}
	err = db.Debug().Model(&MitraProductDocumentRuleTempData{}).Limit(100).
		Select(`m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_id,
				m_mitra_product_temp.mitra_prod_temp_id,
				m_mitra_product_temp.mitra_prod_temp_name,
				m_product_document_temp.prod_doc_id as prod_doc_temp_id,
				m_product_document_temp.prod_doc_code as prod_doc_temp_code,
				m_product_document_temp.prod_doc_name as prod_doc_temp_name,
				m_product_document_category_temp.prod_doc_ctgry_name as prod_doc_ctgry_temp_name,
				m_product_document_category_temp.prod_doc_ctgry_desc as prod_doc_ctgry_temp_desc,
				m_mitra_product_document_rule_temp.cust_sub_type1_temp,
				m_mitra_product_document_rule_temp.cust_sub_type2_temp,
				m_mitra_product_document_rule_temp.cust_sub_type3_temp,
				m_mitra_product_document_rule_temp.cust_sub_type4_temp,
				m_mitra_product_document_rule_temp.cust_sub_type5_temp,
				m_mitra_product_document_rule_temp.cust_sub_type6_temp,
				m_mitra_product_document_rule_temp.cust_sub_type7_temp,
				m_mitra_product_document_rule_temp.cust_sub_type8_temp,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_is_mandatory,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_is_many_img,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_status,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_reason,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_created_by,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_created_at,
				m_mitra_product_document_rule.mitra_prod_doc_rule_id,
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
		Joins("JOIN m_mitra_product_temp on m_mitra_product_document_rule_temp.mitra_prod_temp_id=m_mitra_product_temp.mitra_prod_temp_id").
		Joins("JOIN m_product_document m_product_document_temp on m_mitra_product_document_rule_temp.prod_doc_temp_id=m_product_document_temp.prod_doc_id").
		Joins("JOIN m_product_document_category m_product_document_category_temp on m_product_document_temp.prod_doc_ctgry_id=m_product_document_category_temp.prod_doc_ctgry_id").
		Joins("JOIN m_mitra_product_document_rule on m_mitra_product_document_rule_temp.mitra_prod_doc_rule_id=m_mitra_product_document_rule.mitra_prod_doc_rule_id").
		Joins("JOIN m_mitra_product on m_mitra_product_document_rule.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_document on m_mitra_product_document_rule.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_document_category on m_product_document.prod_doc_ctgry_id=m_product_document_category.prod_doc_ctgry_id").
		Where("mitra_prod_doc_rule_temp_id = ?", pid).
		Order("m_mitra_product_document_rule_temp.prod_doc_temp_id asc").
		Find(&mitraproductdocumentruletemp).Error
	if err != nil {
		return &MitraProductDocumentRuleTempData{}, err
	}
	return &mitraproductdocumentruletemp, nil
}

func (p *MitraProductDocumentRuleTemp) FindMitraProductDocumentRuleTempStatusByID(db *gorm.DB, pid uint64, status uint64) (*MitraProductDocumentRuleTempData, error) {
	var err error
	mitraproductdocumentruletemp := MitraProductDocumentRuleTempData{}
	err = db.Debug().Model(&MitraProductDocumentRuleTempData{}).Limit(100).
		Select(`m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_id,
				m_mitra_product_temp.mitra_prod_temp_id,
				m_mitra_product_temp.mitra_prod_temp_name,
				m_product_document_temp.prod_doc_id as prod_doc_temp_id,
				m_product_document_temp.prod_doc_code as prod_doc_temp_code,
				m_product_document_temp.prod_doc_name as prod_doc_temp_name,
				m_product_document_category_temp.prod_doc_ctgry_name as prod_doc_ctgry_temp_name,
				m_product_document_category_temp.prod_doc_ctgry_desc as prod_doc_ctgry_temp_desc,
				m_mitra_product_document_rule_temp.cust_sub_type1_temp,
				m_mitra_product_document_rule_temp.cust_sub_type2_temp,
				m_mitra_product_document_rule_temp.cust_sub_type3_temp,
				m_mitra_product_document_rule_temp.cust_sub_type4_temp,
				m_mitra_product_document_rule_temp.cust_sub_type5_temp,
				m_mitra_product_document_rule_temp.cust_sub_type6_temp,
				m_mitra_product_document_rule_temp.cust_sub_type7_temp,
				m_mitra_product_document_rule_temp.cust_sub_type8_temp,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_is_mandatory,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_is_many_img,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_status,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_reason,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_created_by,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_created_at,
				m_mitra_product_document_rule.mitra_prod_doc_rule_id,
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
		Joins("JOIN m_mitra_product_temp on m_mitra_product_document_rule_temp.mitra_prod_temp_id=m_mitra_product_temp.mitra_prod_temp_id").
		Joins("JOIN m_product_document m_product_document_temp on m_mitra_product_document_rule_temp.prod_doc_temp_id=m_product_document_temp.prod_doc_id").
		Joins("JOIN m_product_document_category m_product_document_category_temp on m_product_document_temp.prod_doc_ctgry_id=m_product_document_category_temp.prod_doc_ctgry_id").
		Joins("JOIN m_mitra_product_document_rule on m_mitra_product_document_rule_temp.mitra_prod_doc_rule_id=m_mitra_product_document_rule.mitra_prod_doc_rule_id").
		Joins("JOIN m_mitra_product on m_mitra_product_document_rule.mitra_prod_id=m_mitra_product.mitra_prod_id").
		Joins("JOIN m_product_document on m_mitra_product_document_rule.prod_doc_id=m_product_document.prod_doc_id").
		Joins("JOIN m_product_document_category on m_product_document.prod_doc_ctgry_id=m_product_document_category.prod_doc_ctgry_id").
		Where("mitra_prod_doc_rule_temp_id = ? AND mitra_prod_doc_rule_temp_status = ?", pid, status).
		Order("m_mitra_product_document_rule_temp.prod_doc_temp_id asc").
		Find(&mitraproductdocumentruletemp).Error
	if err != nil {
		return &MitraProductDocumentRuleTempData{}, err
	}
	return &mitraproductdocumentruletemp, nil
}

func (p *MitraProductDocumentRuleTemp) FindMitraProductDocumentRuleTempByMitraProductByID(db *gorm.DB, pid uint64, mitrabranchtemp uint64) (*MitraProductDocumentRuleTempPend, error) {
	var err error
	mitraproductdocumentruletemp := MitraProductDocumentRuleTempPend{}
	err = db.Debug().Model(&MitraProductDocumentRuleTempPend{}).Limit(100).
		Select(`m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_id,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_id,
				m_mitra_product_document_rule_temp.mitra_prod_temp_id,
				m_mitra_product_document_rule_temp.prod_doc_temp_id,
				m_mitra_product_document_rule_temp.cust_sub_type1_temp,
				m_mitra_product_document_rule_temp.cust_sub_type2_temp,
				m_mitra_product_document_rule_temp.cust_sub_type3_temp,
				m_mitra_product_document_rule_temp.cust_sub_type4_temp,
				m_mitra_product_document_rule_temp.cust_sub_type5_temp,
				m_mitra_product_document_rule_temp.cust_sub_type6_temp,
				m_mitra_product_document_rule_temp.cust_sub_type7_temp,
				m_mitra_product_document_rule_temp.cust_sub_type8_temp,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_is_mandatory,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_is_many_img,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_status,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_reason,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_created_by,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_created_at`).
		Where("mitra_prod_doc_rule_temp_id = ? AND mitra_prod_temp_id = ?", pid, mitrabranchtemp).
		Order("m_mitra_product_document_rule_temp.prod_doc_temp_id asc").
		Find(&mitraproductdocumentruletemp).Error
	if err != nil {
		return &MitraProductDocumentRuleTempPend{}, err
	}
	return &mitraproductdocumentruletemp, nil
}
func (p *MitraProductDocumentRuleTemp) UpdateMitraProductDocumentRuleBranchTempID(db *gorm.DB, pid uint64) (*MitraProductDocumentRuleTemp, error) {
	var err error
	db = db.Debug().Model(&MitraProductDocumentRuleTemp{}).Where("mitra_prod_doc_rule_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_prod_temp_id":                  p.MitraProductTempID,
			"mitra_prod_doc_rule_temp_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&MitraProductDocumentRuleTemp{}).Where("mitra_prod_doc_rule_temp_id = ?", pid).Error
	if err != nil {
		return &MitraProductDocumentRuleTemp{}, err
	}
	return p, nil
}

func (p *MitraProductDocumentRuleTemp) UpdateMitraProductDocumentRuleTemp(db *gorm.DB, pid uint64) (*MitraProductDocumentRuleTemp, error) {
	var err error
	db = db.Debug().Model(&MitraProductDocumentRuleTemp{}).Where("mitra_prod_doc_rule_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_prod_doc_rule_id":                p.MitraProductDocumentRuleID,
			"mitra_prod_temp_id":                    p.MitraProductTempID,
			"prod_doc_temp_id":                      p.ProductDocumentTempID,
			"cust_sub_type1_temp":                   p.CustomerSubType1Temp,
			"cust_sub_type2_temp":                   p.CustomerSubType2Temp,
			"cust_sub_type3_temp":                   p.CustomerSubType3Temp,
			"cust_sub_type4_temp":                   p.CustomerSubType4Temp,
			"cust_sub_type5_temp":                   p.CustomerSubType5Temp,
			"cust_sub_type6_temp":                   p.CustomerSubType6Temp,
			"cust_sub_type7_temp":                   p.CustomerSubType7Temp,
			"cust_sub_type8_temp":                   p.CustomerSubType8Temp,
			"mitra_prod_doc_rule_temp_is_mandatory": p.MitraProductDocumentRuleTempIsMandatory,
			"mitra_prod_doc_rule_temp_is_many_img":  p.MitraProductDocumentRuleTempIsManyImage,
			"mitra_prod_doc_rule_temp_reason":       p.MitraProductDocumentRuleTempReason,
			"mitra_prod_doc_rule_temp_status":       p.MitraProductDocumentRuleTempStatus,
		},
	)
	err = db.Debug().Model(&MitraProductDocumentRuleTemp{}).Where("mitra_prod_doc_rule_temp_id = ?", pid).Error
	if err != nil {
		return &MitraProductDocumentRuleTemp{}, err
	}
	return p, nil
}

func (p *MitraProductDocumentRuleTemp) UpdateMitraProductDocumentRuleTempStatus(db *gorm.DB, pid uint64) (*MitraProductDocumentRuleTemp, error) {
	var err error
	db = db.Debug().Model(&MitraProductDocumentRuleTemp{}).Where("mitra_prod_doc_rule_temp_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"mitra_prod_doc_rule_temp_status": p.MitraProductDocumentRuleTempStatus,
		},
	)
	err = db.Debug().Model(&MitraProductDocumentRuleTemp{}).Where("mitra_prod_doc_rule_temp_id = ?", pid).Error
	if err != nil {
		return &MitraProductDocumentRuleTemp{}, err
	}
	return p, nil
}

func (p *MitraProductDocumentRuleTemp) DeleteMitraProductDocumentRuleTemp(db *gorm.DB, pid uint64) (int64, error) {
	db = db.Debug().Model(&MitraProductDocumentRuleTemp{}).Where("mitra_prod_doc_rule_temp_id = ?", pid).Take(&MitraProductDocumentRuleTemp{}).Delete(&MitraProductDocumentRuleTemp{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("MitraProductDocumentRuleTemp not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

//ADDITIONAL
//====================================================================================================================================

func (p *MitraProductDocumentRuleTemp) FindMitraProductDocumentRuleTempPendByMitraProductTempID(db *gorm.DB, pid uint64) ([]MitraProductDocumentRuleTempPendData, error) {
	var err error
	mitrabranchcoverageregency := []MitraProductDocumentRuleTempPendData{}
	err = db.Debug().Model(&MitraProductDocumentRuleTempPendData{}).Limit(100).
		Select(`m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_id,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_id,
				m_mitra_product_temp.mitra_prod_temp_id,
				m_mitra_product_temp.mitra_prod_temp_name,
				m_product_document_temp.prod_doc_id as prod_doc_temp_id,
				m_product_document_temp.prod_doc_code as prod_doc_temp_code,
				m_product_document_temp.prod_doc_name as prod_doc_temp_name,
				m_product_document_category_temp.prod_doc_ctgry_name as prod_doc_ctgry_temp_name,
				m_product_document_category_temp.prod_doc_ctgry_desc as prod_doc_ctgry_temp_desc,
				m_mitra_product_document_rule_temp.cust_sub_type1_temp,
				m_mitra_product_document_rule_temp.cust_sub_type2_temp,
				m_mitra_product_document_rule_temp.cust_sub_type3_temp,
				m_mitra_product_document_rule_temp.cust_sub_type4_temp,
				m_mitra_product_document_rule_temp.cust_sub_type5_temp,
				m_mitra_product_document_rule_temp.cust_sub_type6_temp,
				m_mitra_product_document_rule_temp.cust_sub_type7_temp,
				m_mitra_product_document_rule_temp.cust_sub_type8_temp,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_is_mandatory,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_is_many_img,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_status,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_reason,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_created_by,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_created_at`).
		Joins("JOIN m_mitra_product_temp on m_mitra_product_document_rule_temp.mitra_prod_temp_id=m_mitra_product_temp.mitra_prod_temp_id").
		Joins("JOIN m_product_document m_product_document_temp on m_mitra_product_document_rule_temp.prod_doc_temp_id=m_product_document_temp.prod_doc_id").
		Joins("JOIN m_product_document_category m_product_document_category_temp on m_product_document_temp.prod_doc_ctgry_id=m_product_document_category_temp.prod_doc_ctgry_id").
		Where("m_mitra_product_document_rule_temp.mitra_prod_temp_id = ?", pid).
		Order("m_mitra_product_document_rule_temp.prod_doc_temp_id asc").
		Find(&mitrabranchcoverageregency).Error
	if err != nil {
		return []MitraProductDocumentRuleTempPendData{}, err
	}
	return mitrabranchcoverageregency, nil
}

func (p *MitraProductDocumentRuleTemp) FindMitraProductDocumentRuleTempPendStatusByMitraProductTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraProductDocumentRuleTempPendData, error) {
	var err error
	mitrabranchcoverageregency := []MitraProductDocumentRuleTempPendData{}
	err = db.Debug().Model(&MitraProductDocumentRuleTempPendData{}).Limit(100).
		Select(`m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_id,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_id,
				m_mitra_product_temp.mitra_prod_temp_id,
				m_mitra_product_temp.mitra_prod_temp_name,
				m_product_document_temp.prod_doc_id as prod_doc_temp_id,
				m_product_document_temp.prod_doc_code as prod_doc_temp_code,
				m_product_document_temp.prod_doc_name as prod_doc_temp_name,
				m_product_document_category_temp.prod_doc_ctgry_name as prod_doc_ctgry_temp_name,
				m_product_document_category_temp.prod_doc_ctgry_desc as prod_doc_ctgry_temp_desc,
				m_mitra_product_document_rule_temp.cust_sub_type1_temp,
				m_mitra_product_document_rule_temp.cust_sub_type2_temp,
				m_mitra_product_document_rule_temp.cust_sub_type3_temp,
				m_mitra_product_document_rule_temp.cust_sub_type4_temp,
				m_mitra_product_document_rule_temp.cust_sub_type5_temp,
				m_mitra_product_document_rule_temp.cust_sub_type6_temp,
				m_mitra_product_document_rule_temp.cust_sub_type7_temp,
				m_mitra_product_document_rule_temp.cust_sub_type8_temp,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_is_mandatory,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_is_many_img,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_status,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_reason,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_created_by,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_created_at`).
		Joins("JOIN m_mitra_product_temp on m_mitra_product_document_rule_temp.mitra_prod_temp_id=m_mitra_product_temp.mitra_prod_temp_id").
		Joins("JOIN m_product_document m_product_document_temp on m_mitra_product_document_rule_temp.prod_doc_temp_id=m_product_document_temp.prod_doc_id").
		Joins("JOIN m_product_document_category m_product_document_category_temp on m_product_document_temp.prod_doc_ctgry_id=m_product_document_category_temp.prod_doc_ctgry_id").
		Where("m_mitra_product_document_rule_temp.mitra_prod_temp_id = ? AND mitra_prod_doc_rule_temp_status = ?", pid, status).
		Order("m_mitra_product_document_rule_temp.prod_doc_temp_id asc").
		Find(&mitrabranchcoverageregency).Error
	if err != nil {
		return []MitraProductDocumentRuleTempPendData{}, err
	}
	return mitrabranchcoverageregency, nil
}

func (p *MitraProductDocumentRuleTemp) FindMitraProductDocumentRuleTempByMitraProductTempID(db *gorm.DB, pid uint64) ([]MitraProductDocumentRuleTempData, error) {
	var err error
	mitrabranchcoverageregency := []MitraProductDocumentRuleTempData{}
	err = db.Debug().Model(&MitraProductDocumentRuleTempData{}).Limit(100).
		Select(`m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_id,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_id,
				m_mitra_product_temp.mitra_prod_temp_id,
				m_mitra_product_temp.mitra_prod_temp_name,
				m_product_document_temp.prod_doc_id as prod_doc_temp_id,
				m_product_document_temp.prod_doc_code as prod_doc_temp_code,
				m_product_document_temp.prod_doc_name as prod_doc_temp_name,
				m_product_document_category_temp.prod_doc_ctgry_name as prod_doc_ctgry_temp_name,
				m_product_document_category_temp.prod_doc_ctgry_desc as prod_doc_ctgry_temp_desc,
				m_mitra_product_document_rule_temp.cust_sub_type1_temp,
				m_mitra_product_document_rule_temp.cust_sub_type2_temp,
				m_mitra_product_document_rule_temp.cust_sub_type3_temp,
				m_mitra_product_document_rule_temp.cust_sub_type4_temp,
				m_mitra_product_document_rule_temp.cust_sub_type5_temp,
				m_mitra_product_document_rule_temp.cust_sub_type6_temp,
				m_mitra_product_document_rule_temp.cust_sub_type7_temp,
				m_mitra_product_document_rule_temp.cust_sub_type8_temp,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_is_mandatory,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_is_many_img,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_status,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_reason,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_created_by,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_created_at`).
		Joins("JOIN m_mitra_product_temp on m_mitra_product_document_rule_temp.mitra_prod_temp_id=m_mitra_product_temp.mitra_prod_temp_id").
		Joins("JOIN m_product_document m_product_document_temp on m_mitra_product_document_rule_temp.prod_doc_temp_id=m_product_document_temp.prod_doc_id").
		Joins("JOIN m_product_document_category m_product_document_category_temp on m_product_document_temp.prod_doc_ctgry_id=m_product_document_category_temp.prod_doc_ctgry_id").
		Where("m_mitra_product_document_rule_temp.mitra_prod_temp_id = ?", pid).
		Order("m_mitra_product_document_rule_temp.prod_doc_temp_id asc").
		Find(&mitrabranchcoverageregency).Error
	if err != nil {
		return []MitraProductDocumentRuleTempData{}, err
	}
	return mitrabranchcoverageregency, nil
}

func (p *MitraProductDocumentRuleTemp) FindMitraProductDocumentRuleTempStatusByMitraProductTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraProductDocumentRuleTempPendData, error) {
	var err error
	mitrabranchcoverageregency := []MitraProductDocumentRuleTempPendData{}
	err = db.Debug().Model(&MitraProductDocumentRuleTempPendData{}).Limit(100).
		Select(`m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_id,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_id,
				m_mitra_product_temp.mitra_prod_temp_id,
				m_mitra_product_temp.mitra_prod_temp_name,
				m_product_document_temp.prod_doc_id as prod_doc_temp_id,
				m_product_document_temp.prod_doc_code as prod_doc_temp_code,
				m_product_document_temp.prod_doc_name as prod_doc_temp_name,
				m_product_document_category_temp.prod_doc_ctgry_name as prod_doc_ctgry_temp_name,
				m_product_document_category_temp.prod_doc_ctgry_desc as prod_doc_ctgry_temp_desc,
				m_mitra_product_document_rule_temp.cust_sub_type1_temp,
				m_mitra_product_document_rule_temp.cust_sub_type2_temp,
				m_mitra_product_document_rule_temp.cust_sub_type3_temp,
				m_mitra_product_document_rule_temp.cust_sub_type4_temp,
				m_mitra_product_document_rule_temp.cust_sub_type5_temp,
				m_mitra_product_document_rule_temp.cust_sub_type6_temp,
				m_mitra_product_document_rule_temp.cust_sub_type7_temp,
				m_mitra_product_document_rule_temp.cust_sub_type8_temp,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_is_mandatory,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_is_many_img,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_status,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_reason,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_created_by,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_created_at`).
		Joins("JOIN m_mitra_product_temp on m_mitra_product_document_rule_temp.mitra_prod_temp_id=m_mitra_product_temp.mitra_prod_temp_id").
		Joins("JOIN m_product_document m_product_document_temp on m_mitra_product_document_rule_temp.prod_doc_temp_id=m_product_document_temp.prod_doc_id").
		Joins("JOIN m_product_document_category m_product_document_category_temp on m_product_document_temp.prod_doc_ctgry_id=m_product_document_category_temp.prod_doc_ctgry_id").
		Where("m_mitra_product_document_rule_temp.mitra_prod_temp_id = ? AND mitra_prod_doc_rule_temp_status = ?", pid, status).
		Order("m_mitra_product_document_rule_temp.prod_doc_temp_id asc").
		Find(&mitrabranchcoverageregency).Error
	if err != nil {
		return []MitraProductDocumentRuleTempPendData{}, err
	}
	return mitrabranchcoverageregency, nil
}

func (p *MitraProductDocumentRuleTemp) FindMitraProductDocumentRuleTempDataStatusByMitraProductTempID(db *gorm.DB, pid uint64, status uint64) ([]MitraProductDocumentRuleTempData, error) {
	var err error
	mitrabranchcoverageregency := []MitraProductDocumentRuleTempData{}
	err = db.Debug().Model(&MitraProductDocumentRuleTempData{}).Limit(100).
		Select(`m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_id,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_id,
				m_mitra_product_temp.mitra_prod_temp_id,
				m_mitra_product_temp.mitra_prod_temp_name,
				m_product_document_temp.prod_doc_id as prod_doc_temp_id,
				m_product_document_temp.prod_doc_code as prod_doc_temp_code,
				m_product_document_temp.prod_doc_name as prod_doc_temp_name,
				m_product_document_category_temp.prod_doc_ctgry_name as prod_doc_ctgry_temp_name,
				m_product_document_category_temp.prod_doc_ctgry_desc as prod_doc_ctgry_temp_desc,
				m_mitra_product_document_rule_temp.cust_sub_type1_temp,
				m_mitra_product_document_rule_temp.cust_sub_type2_temp,
				m_mitra_product_document_rule_temp.cust_sub_type3_temp,
				m_mitra_product_document_rule_temp.cust_sub_type4_temp,
				m_mitra_product_document_rule_temp.cust_sub_type5_temp,
				m_mitra_product_document_rule_temp.cust_sub_type6_temp,
				m_mitra_product_document_rule_temp.cust_sub_type7_temp,
				m_mitra_product_document_rule_temp.cust_sub_type8_temp,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_is_mandatory,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_is_many_img,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_status,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_reason,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_created_by,
				m_mitra_product_document_rule_temp.mitra_prod_doc_rule_temp_created_at`).
		Joins("JOIN m_mitra_product_temp on m_mitra_product_document_rule_temp.mitra_prod_temp_id=m_mitra_product_temp.mitra_prod_temp_id").
		Joins("JOIN m_product_document m_product_document_temp on m_mitra_product_document_rule_temp.prod_doc_temp_id=m_product_document_temp.prod_doc_id").
		Joins("JOIN m_product_document_category m_product_document_category_temp on m_product_document_temp.prod_doc_ctgry_id=m_product_document_category_temp.prod_doc_ctgry_id").
		Where("m_mitra_product_document_rule_temp.mitra_prod_temp_id = ? AND mitra_prod_doc_rule_temp_status = ?", pid, status).
		Order("m_mitra_product_document_rule_temp.prod_doc_temp_id asc").
		Find(&mitrabranchcoverageregency).Error
	if err != nil {
		return []MitraProductDocumentRuleTempData{}, err
	}
	return mitrabranchcoverageregency, nil
}
