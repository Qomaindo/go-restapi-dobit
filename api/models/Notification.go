package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Notification struct {
	NotificationID        uint64     `gorm:"column:notification_id;primary_key;" json:"notification_id"`
	CustomerUserID        uint64     `gorm:"column:cust_user_id;" json:"cust_user_id"`
	MitraUserID           uint64     `gorm:"column:mitra_user_id;" json:"mitra_user_id"`
	NotificationTitle     string     `gorm:"column:notification_title;" json:"notification_title"`
	NotificationMessage   string     `gorm:"column:notification_message;" json:"notification_message"`
	NotificationSender    string     `gorm:"column:notification_sender;" json:"notification_sender"`
	NotificationLink      string     `gorm:"column:notification_link;" json:"notification_link"`
	NotificationIncoming  uint64     `gorm:"column:notification_incoming;" json:"notification_incoming"`
	NotificationReceiver  string     `gorm:"column:notification_receiver;" json:"notification_receiver"`
	NotificationStatus    uint64     `gorm:"column:notification_status;" json:"notification_status"`
	NotificationCreatedAt time.Time  `gorm:"column:notification_created_at;default:CURRENT_TIMESTAMP" json:"notification_created_at"`
	NotificationDeletedAt *time.Time `gorm:"column:notification_deleted_at;default:CURRENT_TIMESTAMP" json:"notification_deleted_at"`
}

func (p *Notification) Prepare() {
	p.CustomerUserID = p.CustomerUserID
	p.MitraUserID = p.MitraUserID
	p.NotificationTitle = p.NotificationTitle
	p.NotificationMessage = p.NotificationMessage
	p.NotificationSender = p.NotificationSender
	p.NotificationLink = p.NotificationLink
}

type ResponseGetNotification struct {
	Status  uint64          `json:"status"`
	Message string          `json:"message"`
	Data    *[]Notification `json:"data"`
}

type ResponseGetNotificationDetail struct {
	Status  uint64        `json:"status"`
	Message string        `json:"message"`
	Data    *Notification `json:"data"`
}

type ResponseGetNotificationIU struct {
	Status  uint64 `json:"status"`
	Message string `json:"message"`
}

func (Notification) TableName() string {
	return "t_notification"
}

func (u *Notification) SaveNotification(db *gorm.DB) (*Notification, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Notification{}, err
	}
	return u, nil
}

func (p *Notification) FindNotificationCustomer(db *gorm.DB, pid uint64, penerima string) (*[]Notification, error) {
	var err error
	notifMasuk := []Notification{}
	err = db.Debug().Model(&Notification{}).Limit(100).
		Select(`t_notification.notification_id,
				t_notification.cust_user_id,
				t_notification.mitra_user_id,
				t_notification.notification_title,
				t_notification.notification_sender,
				t_notification.notification_link,
				t_notification.notification_message,
				t_notification.notification_incoming,
				t_notification.notification_receiver,
				t_notification.notification_status,
				t_notification.notification_deleted_at at time zone 'Asia/Jakarta' as notification_deleted_at,
				t_notification.notification_created_at at time zone 'Asia/Jakarta' as notification_created_at`).
		Where("notification_receiver = all OR cust_user_id = ? AND notification_receiver = ?", pid, penerima).
		Order("notification_id desc").
		Find(&notifMasuk).Error
	if err != nil {
		return &[]Notification{}, err
	}
	return &notifMasuk, nil
}

func (p *Notification) FindNotificationCustomerBy(db *gorm.DB, pid uint64, penerima string) (*[]Notification, error) {
	var err error
	notifMasuk := []Notification{}
	err = db.Debug().Model(&Notification{}).Limit(100).
		Select(`t_notification.notification_id,
				t_notification.cust_user_id,
				t_notification.mitra_user_id,
				t_notification.notification_title,
				t_notification.notification_sender,
				t_notification.notification_link,
				t_notification.notification_message,
				t_notification.notification_incoming,
				t_notification.notification_receiver,
				t_notification.notification_status,
				t_notification.notification_deleted_at at time zone 'Asia/Jakarta' as notification_deleted_at,
				t_notification.notification_created_at at time zone 'Asia/Jakarta' as notification_created_at`).
		Where("cust_user_id = ? AND notification_receiver = ?", pid, penerima).
		Order("notification_id desc").
		Find(&notifMasuk).Error
	if err != nil {
		return &[]Notification{}, err
	}
	return &notifMasuk, nil
}

func (p *Notification) FindNotificationMitraUser(db *gorm.DB, pid uint64, penerima string) (*[]Notification, error) {
	var err error
	notifMasuk := []Notification{}
	err = db.Debug().Model(&Notification{}).Limit(100).
		Select(`t_notification.notification_id,
				t_notification.cust_user_id,
				t_notification.mitra_user_id,
				t_notification.notification_title,
				t_notification.notification_sender,
				t_notification.notification_link,
				t_notification.notification_message,
				t_notification.notification_incoming,
				t_notification.notification_receiver,
				t_notification.notification_status,
				t_notification.notification_deleted_at at time zone 'Asia/Jakarta' as notification_deleted_at,
				t_notification.notification_created_at at time zone 'Asia/Jakarta' as notification_created_at`).
		Where("mitra_user_id = ? AND notification_receiver = ?", pid, penerima).
		Order("t_notification.notification_id desc").
		Find(&notifMasuk).Error
	if err != nil {
		return &[]Notification{}, err
	}
	return &notifMasuk, nil
}

// Get Notification In Mobile
// ==============================================================================================

func (p *Notification) FindNotificationCustomerByCustomerID(db *gorm.DB, pid uint64, penerima string) (*[]Notification, error) {
	var err error
	notifMasuk := []Notification{}
	err = db.Debug().Model(&Notification{}).Limit(100).
		Select(`t_notification.notification_id,
				t_notification.mitra_user_id,
				m_customer_user.cust_user_id,
				t_notification.notification_title,
				t_notification.notification_sender,
				t_notification.notification_link,
				t_notification.notification_message,
				t_notification.notification_incoming,
				t_notification.notification_receiver,
				t_notification.notification_status,
				t_notification.notification_deleted_at at time zone 'Asia/Jakarta' as notification_deleted_at,
				t_notification.notification_created_at at time zone 'Asia/Jakarta' as notification_created_at`).
		Joins("LEFT JOIN m_customer_user on t_notification.cust_user_id=m_customer_user.cust_user_id").
		Joins("LEFT JOIN m_customer on m_customer_user.cust_id=m_customer.cust_id").
		Where("notification_receiver = ? OR m_customer.cust_id = ? AND notification_receiver = ?", "allcustomer", pid, penerima).
		Order("notification_id desc").
		Find(&notifMasuk).Error
	if err != nil {
		return &[]Notification{}, err
	}
	return &notifMasuk, nil
}

func (p *Notification) FindNotificationMitraUserByMitraUserID(db *gorm.DB, pid uint64, penerima string) (*[]Notification, error) {
	var err error
	notifMasuk := []Notification{}
	err = db.Debug().Model(&Notification{}).Limit(100).
		Select(`t_notification.notification_id,
				m_mitra_user.mitra_user_id,
				t_notification.cust_user_id,
				t_notification.notification_title,
				t_notification.notification_sender,
				t_notification.notification_link,
				t_notification.notification_message,
				t_notification.notification_incoming,
				t_notification.notification_receiver,
				t_notification.notification_status,
				t_notification.notification_deleted_at at time zone 'Asia/Jakarta' as notification_deleted_at,
				t_notification.notification_created_at at time zone 'Asia/Jakarta' as notification_created_at`).
		Where("notification_receiver = ? OR m_mitra_user.mitra_user_id = ? AND notification_receiver = ?", "allmitrauser", pid, penerima).
		Joins("LEFT JOIN m_mitra_user on t_notification.mitra_user_id=m_mitra_user.mitra_user_id").
		Order("t_notification.notification_id desc").
		Find(&notifMasuk).Error
	if err != nil {
		return &[]Notification{}, err
	}
	return &notifMasuk, nil
}

func (p *Notification) UpdateNotificationStatus(db *gorm.DB, pid uint64) (*Notification, error) {
	var err error
	db = db.Debug().Model(&Notification{}).Where("notification_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"notification_status": p.NotificationStatus,
		},
	)
	err = db.Debug().Model(&Notification{}).Where("notification_id = ?", pid).Error
	if err != nil {
		return &Notification{}, err
	}
	return p, nil
}

func (p *Notification) UpdateNotificationIncoming(db *gorm.DB, pid uint64) (*Notification, error) {
	var err error
	db = db.Debug().Model(&Notification{}).Where("notification_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"notification_incoming": 1,
		},
	)
	err = db.Debug().Model(&Notification{}).Where("notification_id = ?", pid).Error
	if err != nil {
		return &Notification{}, err
	}
	return p, nil
}
