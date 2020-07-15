package models

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type LoanSubmission struct {
	LoanSubmissionID           uint64     `gorm:"column:loan_sbmssn_id;primary_key;" json:"loan_sbmssn_id"`
	LoanRequestID              uint64     `gorm:"column:loan_req_id;primary_key;" json:"loan_req_id"`
	MitraID                    uint64     `gorm:"column:mitra_id" json:"mitra_id"`
	MitraBranchID              uint64     `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	MitraUserID                uint64     `gorm:"column:mitra_user_id" json:"mitra_user_id"`
	CustUserID                 uint64     `gorm:"column:cust_user_id" json:"cust_user_id"`
	ClientID                   uint64     `gorm:"column:client_id" json:"client_id"`
	LoanSubmissionCode         string     `gorm:"column:loan_sbmssn_code" json:"loan_sbmssn_code"`
	LoanSubmissionNum          string     `gorm:"column:loan_sbmssn_num" json:"loan_sbmssn_num"`
	LoanSubmissionAmount       uint64     `gorm:"column:loan_sbmssn_amount" json:"loan_sbmssn_amount"`
	LoanSubmissionLong         uint64     `gorm:"column:loan_sbmssn_long" json:"loan_sbmssn_long"`
	LoanSubmissionInterestRate float64    `gorm:"column:loan_sbmssn_interest_rate" json:"loan_sbmssn_interest_rate"`
	LoanSubmissionInterestDisc float64    `gorm:"column:loan_sbmssn_interest_disc" json:"loan_sbmssn_interest_disc"`
	LoanSubmissionDate         string     `gorm:"column:loan_sbmssn_date" json:"loan_sbmssn_date"`
	LoanSubmissionStatus       uint64     `gorm:"column:loan_sbmssn_status" json:"loan_sbmssn_status"`
	LoanSubmissionCreatedAt    time.Time  `gorm:"column:loan_sbmssn_created_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_created_at"`
	LoanSubmissionUpdatedAt    *time.Time `gorm:"column:loan_sbmssn_updated_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_created_at"`
	LoanSubmissionDeletedAt    *time.Time `gorm:"column:loan_sbmssn_deleted_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_deleted_at"`
}

type LoanSubmissionData struct {
	LoanSubmissionID           uint64     `gorm:"column:loan_sbmssn_id;primary_key;" json:"loan_sbmssn_id"`
	LoanRequestID              uint64     `gorm:"column:loan_req_id;primary_key;" json:"loan_req_id"`
	MitraID                    uint64     `gorm:"column:mitra_id" json:"mitra_id"`
	MitraBranchID              uint64     `gorm:"column:mitra_branch_id" json:"mitra_branch_id"`
	MitraUserID                uint64     `gorm:"column:mitra_user_id" json:"mitra_user_id"`
	CustUserID                 uint64     `gorm:"column:cust_user_id" json:"cust_user_id"`
	CustName                   string     `gorm:"column:cust_name" json:"cust_name"`
	ClientID                   uint64     `gorm:"column:client_id" json:"client_id"`
	LoanSubmissionCode         string     `gorm:"column:loan_sbmssn_code" json:"loan_sbmssn_code"`
	LoanSubmissionNum          string     `gorm:"column:loan_sbmssn_num" json:"loan_sbmssn_num"`
	LoanSubmissionAmount       uint64     `gorm:"column:loan_sbmssn_amount" json:"loan_sbmssn_amount"`
	LoanSubmissionLong         uint64     `gorm:"column:loan_sbmssn_long" json:"loan_sbmssn_long"`
	LoanSubmissionInterestRate float64    `gorm:"column:loan_sbmssn_interest_rate" json:"loan_sbmssn_interest_rate"`
	LoanSubmissionInterestDisc float64    `gorm:"column:loan_sbmssn_interest_disc" json:"loan_sbmssn_interest_disc"`
	LoanSubmissionDate         string     `gorm:"column:loan_sbmssn_date" json:"loan_sbmssn_date"`
	LoanSubmissionStatus       uint64     `gorm:"column:loan_sbmssn_status" json:"loan_sbmssn_status"`
	LoanSubmissionCreatedAt    time.Time  `gorm:"column:loan_sbmssn_created_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_created_at"`
	LoanSubmissionUpdatedAt    *time.Time `gorm:"column:loan_sbmssn_updated_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_created_at"`
	LoanSubmissionDeletedAt    *time.Time `gorm:"column:loan_sbmssn_deleted_at;default:CURRENT_TIMESTAMP" json:"loan_sbmssn_deleted_at"`
}

type ResponseLoanSubmissions struct {
	Status  int                   `json:"status"`
	Message string                `json:"message"`
	Data    *[]LoanSubmissionData `json:"data"`
}

type ResponseLoanSubmission struct {
	Status  int                 `json:"status"`
	Message string              `json:"message"`
	Data    *LoanSubmissionData `json:"data"`
}

type ResponseLoanSubmissionCUD struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (LoanSubmission) TableName() string {
	return "t_loan_submission"
}

func (LoanSubmissionData) TableName() string {
	return "t_loan_submission"
}

func (p *LoanSubmission) Prepare() {
	p.LoanRequestID = p.LoanRequestID
	p.MitraID = p.MitraID
	p.MitraBranchID = p.MitraBranchID
	p.MitraUserID = p.MitraUserID
	p.CustUserID = p.CustUserID
	p.ClientID = p.ClientID
	p.LoanSubmissionCode = p.LoanSubmissionCode
	p.LoanSubmissionAmount = p.LoanSubmissionAmount
	p.LoanSubmissionLong = p.LoanSubmissionLong
	p.LoanSubmissionInterestRate = p.LoanSubmissionInterestRate
	p.LoanSubmissionStatus = p.LoanSubmissionStatus
	p.LoanSubmissionCreatedAt = time.Now()
	p.LoanSubmissionUpdatedAt = p.LoanSubmissionUpdatedAt
	p.LoanSubmissionDeletedAt = p.LoanSubmissionDeletedAt
}
func (p *LoanSubmission) Validate(action string) error {
	switch strings.ToLower(action) {

	default:
		// if p.CountryID == 0 {
		// 	return errors.New("Required Region Country")
		// }
		// if p.ProvinceID == 0 {
		// 	return errors.New("Required Region Province")
		// }
		// if p.RegencyID == 0 {
		// 	return errors.New("Required Region Regency")
		// }
		// if p.DistrictID == 0 {
		// 	return errors.New("Required Region District")
		// }
		// if p.VillageID == 0 {
		// 	return errors.New("Required Region Village")
		// }
		return nil
	}
}

func (p *LoanSubmission) SaveLoanSubmission(db *gorm.DB) (*LoanSubmission, error) {
	var err error
	err = db.Debug().Model(&LoanSubmission{}).Create(&p).Error
	if err != nil {
		return &LoanSubmission{}, err
	}
	return p, nil
}

func (p *LoanSubmission) FindAllLoanSubmission(db *gorm.DB) (*[]LoanSubmissionData, error) {
	var err error
	loan_sbmssn := []LoanSubmissionData{}
	err = db.Debug().Model(&LoanSubmissionData{}).Limit(100).
		Select(`t_loan_submission.loan_sbmssn_id,
				t_loan_request.loan_req_id,
				m_customer_user.cust_user_id,	  
				m_customer.cust_name,	     
				m_mitra_user.mitra_user_id,
				m_mitra.mitra_id,
				m_mitra_branch.mitra_branch_id,
				t_loan_submission.loan_sbmssn_num,
				t_loan_submission.loan_sbmssn_amount,
				t_loan_submission.loan_sbmssn_long,
				t_loan_submission.loan_amount_reference,
				t_loan_submission.loan_sbmssn_long_reference,
				t_loan_submission.loan_sbmssn_interest_disc,
				t_loan_submission.loan_sbmssn_interest_rate,
				t_loan_submission.loan_sbmssn_status,
				t_loan_submission.loan_sbmssn_long_reference,
				t_loan_submission.loan_sbmssn_date,
				t_loan_submission.loan_sbmssn_created_at,
				t_loan_submission.loan_sbmssn_updated_at,
				t_loan_submission.loan_sbmssn_deleted_at`).
		Joins("JOIN t_loan_request on t_loan_submission.loan_req_id=t_loan_request.loan_req_id").
		Joins("JOIN m_customer_user on t_loan_submission.cust_user_id=m_customer_user.cust_user_id").
		Joins("JOIN m_customer on m_customer_user.cust_id=m_customer.cust_id").
		Joins("JOIN m_mitra_user on t_loan_submission.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra on t_loan_submission.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on t_loan_submission.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Find(&loan_sbmssn).Error
	if err != nil {
		return &[]LoanSubmissionData{}, err
	}
	return &loan_sbmssn, nil
}

func (p *LoanSubmission) FindAllLoanSubmissionSchdlStatus(db *gorm.DB, status uint64) (*[]LoanSubmissionData, error) {
	var err error
	loan_sbmssn := []LoanSubmissionData{}
	err = db.Debug().Model(&LoanSubmissionData{}).Limit(100).
		Select(`t_loan_submission.loan_sbmssn_id,
			t_loan_request.loan_req_id,
			m_customer_user.cust_user_id,	  
			m_customer.cust_name,	     
			m_mitra_user.mitra_user_id,
			m_mitra.mitra_id,
			m_mitra_branch.mitra_branch_id,
			t_loan_submission.loan_sbmssn_num,
			t_loan_submission.loan_sbmssn_amount,
			t_loan_submission.loan_sbmssn_long,
			t_loan_submission.loan_amount_reference,
			t_loan_submission.loan_sbmssn_long_reference,
			t_loan_submission.loan_sbmssn_interest_disc,
			t_loan_submission.loan_sbmssn_interest_rate,
			t_loan_submission.loan_sbmssn_status,
			t_loan_submission.loan_sbmssn_long_reference,
			t_loan_submission.loan_sbmssn_date,
			t_loan_submission.loan_sbmssn_created_at,
			t_loan_submission.loan_sbmssn_updated_at,
			t_loan_submission.loan_sbmssn_deleted_at`).
		Joins("JOIN t_loan_request on t_loan_submission.loan_req_id=t_loan_request.loan_req_id").
		Joins("JOIN m_customer_user on t_loan_submission.cust_user_id=m_customer_user.cust_user_id").
		Joins("JOIN m_customer on m_customer_user.cust_id=m_customer.cust_id").
		Joins("JOIN m_mitra_user on t_loan_submission.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra on t_loan_submission.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on t_loan_submission.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Where("loan_sbmssn_status = ?", status).
		Find(&loan_sbmssn).Error
	if err != nil {
		return &[]LoanSubmissionData{}, err
	}
	return &loan_sbmssn, nil
}

func (p *LoanSubmission) FindLoanSubmissionDataByID(db *gorm.DB, pid uint64) (*LoanSubmission, error) {
	var err error
	err = db.Debug().Model(&LoanSubmission{}).Where("loan_sbmssn_id = ?", pid).Take(&p).Error
	if err != nil {
		return &LoanSubmission{}, err
	}
	return p, nil
}

func (p *LoanSubmission) FindLoanSubmissionByID(db *gorm.DB, pid uint64) (*LoanSubmissionData, error) {
	var err error
	loan_sbmssn := LoanSubmissionData{}
	err = db.Debug().Model(&LoanSubmissionData{}).Limit(100).
		Select(`t_loan_submission.loan_sbmssn_id,
				t_loan_request.loan_req_id,
				m_customer_user.cust_user_id,	  
				m_customer.cust_name,	     
				m_mitra_user.mitra_user_id,
				m_mitra.mitra_id,
				m_mitra_branch.mitra_branch_id,
				t_loan_submission.loan_sbmssn_num,
				t_loan_submission.loan_sbmssn_amount,
				t_loan_submission.loan_sbmssn_long,
				t_loan_submission.loan_amount_reference,
				t_loan_submission.loan_sbmssn_long_reference,
				t_loan_submission.loan_sbmssn_interest_disc,
				t_loan_submission.loan_sbmssn_interest_rate,
				t_loan_submission.loan_sbmssn_status,
				t_loan_submission.loan_sbmssn_long_reference,
				t_loan_submission.loan_sbmssn_date,
				t_loan_submission.loan_sbmssn_created_at,
				t_loan_submission.loan_sbmssn_updated_at,
				t_loan_submission.loan_sbmssn_deleted_at`).
		Joins("JOIN t_loan_request on t_loan_submission.loan_req_id=t_loan_request.loan_req_id").
		Joins("JOIN m_customer_user on t_loan_submission.cust_user_id=m_customer_user.cust_user_id").
		Joins("JOIN m_customer on m_customer_user.cust_id=m_customer.cust_id").
		Joins("JOIN m_mitra_user on t_loan_submission.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra on t_loan_submission.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on t_loan_submission.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Where("t_loan_submission.loan_sbmssn_id = ?", pid).
		Find(&loan_sbmssn).Error
	if err != nil {
		return &LoanSubmissionData{}, err
	}
	return &loan_sbmssn, nil
}

func (p *LoanSubmission) FindLoanSubmissionBySbmssnID(db *gorm.DB, pid uint64) (*LoanSubmissionData, error) {
	var err error
	loan_sbmssn := LoanSubmissionData{}
	err = db.Debug().Model(&LoanSubmissionData{}).Limit(100).
		Select(`t_loan_submission.loan_sbmssn_id,
				t_loan_request.loan_req_id,
				m_customer_user.cust_user_id,	  
				m_customer.cust_name,	     
				m_mitra_user.mitra_user_id,
				m_mitra.mitra_id,
				m_mitra_branch.mitra_branch_id,
				t_loan_submission.loan_sbmssn_num,
				t_loan_submission.loan_sbmssn_amount,
				t_loan_submission.loan_sbmssn_long,
				t_loan_submission.loan_amount_reference,
				t_loan_submission.loan_sbmssn_long_reference,
				t_loan_submission.loan_sbmssn_interest_disc,
				t_loan_submission.loan_sbmssn_interest_rate,
				t_loan_submission.loan_sbmssn_status,
				t_loan_submission.loan_sbmssn_long_reference,
				t_loan_submission.loan_sbmssn_date,
				t_loan_submission.loan_sbmssn_created_at,
				t_loan_submission.loan_sbmssn_updated_at,
				t_loan_submission.loan_sbmssn_deleted_at`).
		Joins("JOIN t_loan_request on t_loan_submission.loan_req_id=t_loan_request.loan_req_id").
		Joins("JOIN m_customer_user on t_loan_submission.cust_user_id=m_customer_user.cust_user_id").
		Joins("JOIN m_customer on m_customer_user.cust_id=m_customer.cust_id").
		Joins("JOIN m_mitra_user on t_loan_submission.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra on t_loan_submission.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on t_loan_submission.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Where("t_loan_request.loan_req_id = ?", pid).
		Find(&loan_sbmssn).Error
	if err != nil {
		return &LoanSubmissionData{}, err
	}
	return &loan_sbmssn, nil
}

func (p *LoanSubmission) FindLoanSubmissionSchdlStatusByID(db *gorm.DB, pid uint64, status uint64) (*LoanSubmissionData, error) {
	var err error
	loan_sbmssn := LoanSubmissionData{}
	err = db.Debug().Model(&LoanSubmissionData{}).Limit(100).
		Select(`t_loan_submission.loan_sbmssn_id,
				t_loan_request.loan_req_id,
				m_customer_user.cust_user_id,	  
				m_customer.cust_name,	     
				m_mitra_user.mitra_user_id,
				m_mitra.mitra_id,
				m_mitra_branch.mitra_branch_id,
				t_loan_submission.loan_sbmssn_num,
				t_loan_submission.loan_sbmssn_amount,
				t_loan_submission.loan_sbmssn_long,
				t_loan_submission.loan_amount_reference,
				t_loan_submission.loan_sbmssn_long_reference,
				t_loan_submission.loan_sbmssn_interest_disc,
				t_loan_submission.loan_sbmssn_interest_rate,
				t_loan_submission.loan_sbmssn_status,
				t_loan_submission.loan_sbmssn_long_reference,
				t_loan_submission.loan_sbmssn_date,
				t_loan_submission.loan_sbmssn_created_at,
				t_loan_submission.loan_sbmssn_updated_at,
				t_loan_submission.loan_sbmssn_deleted_at`).
		Joins("JOIN t_loan_request on t_loan_submission.loan_req_id=t_loan_request.loan_req_id").
		Joins("JOIN m_customer_user on t_loan_submission.cust_user_id=m_customer_user.cust_user_id").
		Joins("JOIN m_customer on m_customer_user.cust_id=m_customer.cust_id").
		Joins("JOIN m_mitra_user on t_loan_submission.mitra_user_id=m_mitra_user.mitra_user_id").
		Joins("JOIN m_mitra on t_loan_submission.mitra_id=m_mitra.mitra_id").
		Joins("JOIN m_mitra_branch on t_loan_submission.mitra_branch_id=m_mitra_branch.mitra_branch_id").
		Where("loan_sbmssn_id = ? AND loan_sbmssn_status = ?", pid, status).
		Find(&loan_sbmssn).Error
	if err != nil {
		return &LoanSubmissionData{}, err
	}
	return &loan_sbmssn, nil
}

func (p *LoanSubmission) UpdateLoanSubmission(db *gorm.DB, pid uint64) (*LoanSubmission, error) {

	var err error
	db = db.Debug().Model(&LoanSubmission{}).Where("loan_sbmssn_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_req_id":               p.LoanRequestID,
			"mitra_id":                  p.MitraID,
			"mitra_branch_id":           p.MitraBranchID,
			"mitra_user_id":             p.MitraUserID,
			"cust_user_id":              p.CustUserID,
			"client_id":                 p.ClientID,
			"loan_sbmssn_code":          p.LoanSubmissionCode,
			"loan_sbmssn_amount":        p.LoanSubmissionAmount,
			"loan_sbmssn_long":          p.LoanSubmissionLong,
			"loan_sbmssn_interest_rate": p.LoanSubmissionInterestRate,
			"loan_sbmssn_interest_disc": p.LoanSubmissionInterestDisc,
			"loan_sbmssn_status":        p.LoanSubmissionStatus,
			"loan_sbmssn_updated_at":    time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmission{}).Where("loan_sbmssn_id = ?", pid).Error
	if err != nil {
		return &LoanSubmission{}, err
	}
	return p, nil
}

func (p *LoanSubmission) UpdateLoanSubmissionStatus(db *gorm.DB, pid uint64) (*LoanSubmission, error) {
	var err error
	db = db.Debug().Model(&LoanSubmission{}).Where("loan_sbmssn_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_status":     p.LoanSubmissionStatus,
			"loan_sbmssn_updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmission{}).Where("loan_sbmssn_id = ?", pid).Error
	if err != nil {
		return &LoanSubmission{}, err
	}
	return p, nil
}

func (p *LoanSubmission) UpdateLoanSubmissionSchdlStatusDelete(db *gorm.DB, pid uint64) (*LoanSubmission, error) {
	var err error
	db = db.Debug().Model(&LoanSubmission{}).Where("loan_sbmssn_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"loan_sbmssn_status":     3,
			"loan_sbmssn_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&LoanSubmission{}).Where("loan_sbmssn_id = ?", pid).Error
	if err != nil {
		return &LoanSubmission{}, err
	}
	return p, nil
}

func (p *LoanSubmission) DeleteLoanSubmission(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&LoanSubmission{}).Where("loan_sbmssn_id = ?", pid).Take(&LoanSubmission{}).Delete(&LoanSubmission{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("LoanSubmission not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// LOAN MEMORANDUM
// ===================================================================================================================================

type ResponseLoanMemorandum struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    LoanMemorandumData `json:"data"`
}

type LoanMemorandumData struct {
	SubmissionData            LoanMemorandumSubmissionData            `json:"data_pengaju"`
	DocumentData              []LoanSubmissionProductDocumentMrdmData `json:"data_dokumen"`
	ClientData                LoanMemorandumClientData                `json:"data_pribadi"`
	FinancialData             LoanMemorandumFinancialData             `json:"data_finansial"`
	InvestigationData         LoanSubmissionInvestigationMrdmData     `json:"data_investigasi"`
	ConclutionRecommendations LoanSubmissionConlutionMrdmData         `json:"kesimpulan_rekomendasi"`
}

type LoanMemorandumSubmissionData struct {
	SubmissionInformation LoanSubmissionMrdmData        `json:"informasi_pengajuan"`
	SchedulingData        LoanRequestSchedulingMrdmData `json:"jadwal_pertemuan"`
	CollateralData        LoanRequestCollateralMrdmData `json:"data_jaminan"`
}

type LoanMemorandumClientData struct {
	PersonalData         ClientMrdmData                 `json:"data_diri"`
	FamilyData           ClientFamilyMrdmData           `json:"data_keluarga"`
	EducationData        []ClientEducationMrdmData      `json:"data_pendidikan"`
	WorkExperienceData   []ClientWorkExperienceMrdmData `json:"data_pekerjaan"`
	WorkSupervisor       []ClientWorkSupervisorMrdmData `json:"data_atasan"`
	CharacterDescription string                         `json:"deskripsi_karakter"`
}

type LoanMemorandumFinancialData struct {
	AccountAnalysis []LoanSubmissionAccountAnalysisMrdmData `json:"analisa_akun"`
	MonthlyIncome   LoanSubmissionIncomeMrdmData            `json:"pendapatan_per_bulan"`
	MonthlyExpense  LoanSubmissionExpenseMrdmData           `json:"pengeluaran_per_bulan"`
}

// ADDITIONAL MEMORANDUM
// ===================================================================================================================================

type LoanSubmissionMrdmData struct {
	LoanSubmissionID                         uint64  `gorm:"column:loan_sbmssn_id" json:"loan_sbmssn_id"`
	LoanRequestID                            uint64  `gorm:"column:loan_req_id" json:"loan_req_id"`
	LoanSubmissionCode                       string  `gorm:"column:loan_sbmssn_code" json:"loan_sbmssn_code"`
	LoanSubmissionDate                       string  `gorm:"column:loan_sbmssn_date" json:"loan_sbmssn_date"`
	ProductSubCategoryID                     uint64  `gorm:"column:prod_sub_ctgry_id" json:"prod_sub_ctgry_id"`
	ProductSubCategoryCode                   string  `gorm:"column:prod_sub_ctgry_code" json:"prod_sub_ctgry_code"`
	ProductSubCategoryName                   string  `gorm:"column:prod_sub_ctgry_name" json:"prod_sub_ctgry_name"`
	ProductSubCategoryInitial                string  `gorm:"column:prod_sub_ctgry_initial" json:"prod_sub_ctgry_initial"`
	ClientID                                 uint64  `gorm:"column:client_id" json:"client_id"`
	ClientCode                               string  `gorm:"column:client_code" json:"client_code"`
	ClientName                               string  `gorm:"column:client_name" json:"client_name"`
	ClientSex                                string  `gorm:"column:client_sex" json:"client_sex"`
	ClientAge                                string  `gorm:"column:client_age" json:"client_age"`
	ClientJob                                string  `gorm:"column:client_job" json:"client_job"`
	LoanSubmissionStatus                     uint64  `gorm:"column:loan_sbmssn_status" json:"loan_sbmssn_status"`
	LoanSubmissionStatusExplain              string  `gorm:"column:loan_sbmssn_status_explain" json:"loan_sbmssn_status_explain"`
	LoanRequestAmount                        uint64  `gorm:"column:loan_req_amount" json:"loan_req_amount"`
	LoanRequestLong                          uint64  `gorm:"column:loan_req_long" json:"loan_req_long"`
	LoanRequestPurpose                       string  `gorm:"column:loan_req_purpose" json:"loan_req_purpose"`
	LoanRequestSalaryRange                   string  `gorm:"column:loan_req_salary_range" json:"loan_req_salary_range"`
	LoanSubmissionProductID                  uint64  `gorm:"column:loan_sbmssn_prod_id;primary_key;" json:"loan_sbmssn_prod_id"`
	LoanSubmissionProductName                string  `gorm:"column:loan_sbmssn_prod_name" json:"loan_sbmssn_prod_name"`
	LoanSubmissionProductLimitMin            uint64  `gorm:"column:loan_sbmssn_prod_limit_min" json:"loan_sbmssn_prod_limit_min"`
	LoanSubmissionProductLimitMax            uint64  `gorm:"column:loan_sbmssn_prod_limit_max" json:"loan_sbmssn_prod_limit_max"`
	LoanSubmissionProductTenorMin            uint64  `gorm:"column:loan_sbmssn_prod_tenor_min" json:"loan_sbmssn_prod_tenor_min"`
	LoanSubmissionProductTenorMax            uint64  `gorm:"column:loan_sbmssn_prod_tenor_max" json:"loan_sbmssn_prod_tenor_max"`
	LoanSubmissionProductInterestRate        float64 `gorm:"column:loan_sbmssn_prod_interest_rate" json:"loan_sbmssn_prod_interest_rate"`
	LoanSubmissionProductInterestRatePeriod  string  `gorm:"column:loan_sbmssn_prod_interest_rate_period" json:"loan_sbmssn_prod_interest_rate_period"`
	LoanSubmissionProductInterestRateFormula string  `gorm:"column:loan_sbmssn_prod_interest_rate_formula" json:"loan_sbmssn_prod_interest_rate_formula"`
	LoanSubmissionProductInterestRateDiscMax float64 `gorm:"column:loan_sbmssn_prod_interest_rate_disc_max" json:"loan_sbmssn_prod_interest_rate_disc_max"`
	LoanSubmissionProductNumReviewer         uint64  `gorm:"column:loan_sbmssn_prod_num_reviewer" json:"loan_sbmssn_prod_num_reviewer"`
	LoanSubmissionProductNumApprover         uint64  `gorm:"column:loan_sbmssn_prod_num_approver" json:"loan_sbmssn_prod_num_approver"`
}

func (LoanSubmissionMrdmData) TableName() string {
	return "t_loan_submission"
}

func (p *LoanSubmission) FindLoanSubmissionMrdmByID(db *gorm.DB, pid uint64) (LoanSubmissionMrdmData, error) {
	var err error
	loan_sbmssn := LoanSubmissionMrdmData{}
	err = db.Debug().Model(&LoanSubmissionMrdmData{}).Limit(100).
		Select(`t_loan_submission.loan_sbmssn_id,
				t_loan_request.loan_req_id,
				t_loan_submission.loan_sbmssn_code,
				t_loan_submission.loan_sbmssn_date,
				m_product_sub_category.prod_sub_ctgry_id,
				m_product_sub_category.prod_sub_ctgry_code,
				m_product_sub_category.prod_sub_ctgry_name,
				m_product_sub_category.prod_sub_ctgry_initial,
				m_client.client_id,
				m_client.client_code,
				m_client.client_name,
				m_client.client_sex,
				m_client.client_age,
				t_loan_request.loan_req_job as client_job,
				t_loan_submission.loan_sbmssn_status,
				t_loan_request.loan_req_amount,
				t_loan_request.loan_req_long,
				t_loan_request.loan_req_purpose,
				t_loan_request.loan_req_salary_range,
				t_loan_submission_product.loan_sbmssn_prod_id,
				t_loan_submission_product.loan_sbmssn_prod_name,
				t_loan_submission_product.loan_sbmssn_prod_limit_min,
				t_loan_submission_product.loan_sbmssn_prod_limit_max,
				t_loan_submission_product.loan_sbmssn_prod_tenor_min,
				t_loan_submission_product.loan_sbmssn_prod_tenor_max,
				t_loan_submission_product.loan_sbmssn_prod_interest_rate,
				t_loan_submission_product.loan_sbmssn_prod_interest_rate_period,
				t_loan_submission_product.loan_sbmssn_prod_interest_rate_formula,
				t_loan_submission_product.loan_sbmssn_prod_interest_rate_disc_max,
				t_loan_submission_product.loan_sbmssn_prod_num_reviewer,
				t_loan_submission_product.loan_sbmssn_prod_num_approver`).
		Joins("JOIN t_loan_request on t_loan_submission.loan_req_id=t_loan_request.loan_req_id").
		Joins("JOIN t_loan_submission_product on t_loan_submission.loan_sbmssn_id=t_loan_submission_product.loan_sbmssn_id").
		Joins("JOIN t_loan_submission_product_sub_category on t_loan_submission_product.loan_sbmssn_prod_id=t_loan_submission_product_sub_category.loan_sbmssn_prod_id").
		Joins("JOIN m_product_sub_category on t_loan_submission_product_sub_category.prod_sub_ctgry_id=m_product_sub_category.prod_sub_ctgry_id").
		Joins("JOIN m_client on t_loan_submission.client_id=m_client.client_id").
		Where("t_loan_submission.loan_sbmssn_id = ?", pid).
		Find(&loan_sbmssn).Error
	if err != nil {
		return LoanSubmissionMrdmData{}, err
	}
	return loan_sbmssn, nil
}
