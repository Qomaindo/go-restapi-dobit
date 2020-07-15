package models

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Message struct {
	MessageID        uint64     `gorm:"column:message_id;primary_key;" json:"message_id"`
	MessageFrom      string     `gorm:"column:message_from;" json:"message_from"`
	MessageSender    string     `gorm:"column:message_sender;" json:"message_sender"`
	MessageTo        string     `gorm:"column:message_to;" json:"message_to"`
	MessageLink      string     `gorm:"column:message_link;" json:"message_link"`
	MessageTitle     string     `gorm:"column:message_title" json:"message_title"`
	MessageDetail    string     `gorm:"column:message_detail" json:"message_detail"`
	MessageStatus    uint64     `gorm:"column:message_status;size:2" json:"message_status"`
	MessageCreatedBy string     `gorm:"column:message_created_by" json:"message_created_by"`
	MessageCreatedAt time.Time  `gorm:"column:message_created_at;default:CURRENT_TIMESTAMP" json:"message_created_at"`
	MessageDeletedBy string     `gorm:"column:message_deleted_by" json:"message_deleted_by"`
	MessageDeletedAt *time.Time `gorm:"column:message_deleted_at;default:CURRENT_TIMESTAMP" json:"message_deleted_at"`
}

type MessageData struct {
	MessageID        uint64     `gorm:"column:message_id;primary_key;" json:"message_id"`
	MessageFrom      string     `gorm:"column:message_from;" json:"message_from"`
	MessageSender    string     `gorm:"column:message_sender;" json:"message_sender"`
	MessageTo        string     `gorm:"column:message_to;" json:"message_to"`
	MessageLink      string     `gorm:"column:message_link;" json:"message_link"`
	MessageTitle     string     `gorm:"column:message_title" json:"message_title"`
	MessageDetail    string     `gorm:"column:message_detail" json:"message_detail"`
	MessageStatus    uint64     `gorm:"column:message_status;size:2" json:"message_status"`
	MessageCreatedBy string     `gorm:"column:message_created_by" json:"message_created_by"`
	MessageCreatedAt time.Time  `gorm:"column:message_created_at;default:CURRENT_TIMESTAMP" json:"message_created_at"`
	MessageDeletedBy string     `gorm:"column:message_deleted_by" json:"message_deleted_by"`
	MessageDeletedAt *time.Time `gorm:"column:message_deleted_at;default:CURRENT_TIMESTAMP" json:"message_deleted_at"`
}

type ResponseMessages struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    *[]MessageData `json:"data"`
}

type ResponseMessage struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    *MessageData `json:"data"`
}

type ResponseMessageIU struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    *Message `json:"data"`
}

type ResponseMessageDel struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (Message) TableName() string {
	return "t_message"
}

func (MessageData) TableName() string {
	return "t_message"
}

func (p *Message) Prepare() {

	// timeNow := fmt.Sprintf(time.Now().Format("2006-01-08 15:04:05"))
	// var timeWIB string
	// timeWIB = fmt.Sprintf(timeNow + "+00")

	p.MessageFrom = p.MessageFrom
	p.MessageSender = p.MessageSender
	p.MessageTo = p.MessageTo
	p.MessageLink = p.MessageLink
	p.MessageTitle = p.MessageTitle
	p.MessageDetail = p.MessageDetail
	p.MessageStatus = p.MessageStatus
	p.MessageCreatedBy = p.MessageCreatedBy
	p.MessageCreatedAt = time.Now()
	p.MessageDeletedBy = p.MessageDeletedBy
	p.MessageDeletedAt = p.MessageDeletedAt
}

func (p *Message) Validate(action string) error {
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

func (p *Message) SaveMessage(db *gorm.DB) (*Message, error) {
	var err error
	err = db.Debug().Model(&Message{}).Create(&p).Error
	if err != nil {
		return &Message{}, err
	}
	return p, nil
}

func (p *Message) FindAllMessage(db *gorm.DB, pid string) (*[]MessageData, error) {
	var err error
	loan_req_timeline := []MessageData{}
	err = db.Debug().Model(&MessageData{}).Limit(100).
		Select(`t_message.message_id,
				t_message.message_from,
				t_message.message_sender,
				t_message.message_to,
				t_message.message_link,
				t_message.message_title,
				t_message.message_detail,
				t_message.message_status,
				t_message.message_created_by,
				t_message.message_created_at at time zone 'Asia/Jakarta' as message_created_at,
				t_message.message_deleted_by,
				t_message.message_deleted_at at time zone 'Asia/Jakarta' as message_deleted_at`).
		Where("message_to = ?", pid).
		Order("t_message.message_id desc").
		Find(&loan_req_timeline).Error
	if err != nil {
		return &[]MessageData{}, err
	}
	return &loan_req_timeline, nil
}

func (p *Message) FindAllMessageStatus(db *gorm.DB, status uint64) (*[]MessageData, error) {
	var err error
	loan_req_timeline := []MessageData{}
	err = db.Debug().Model(&MessageData{}).Limit(100).
		Select(`t_message.message_id,
				t_message.message_from,
				t_message.message_sender,
				t_message.message_to,
				t_message.message_link,
				t_message.message_title,
				t_message.message_detail,
				t_message.message_status,
				t_message.message_created_by,
				t_message.message_created_at at time zone 'Asia/Jakarta' as message_created_at,
				t_message.message_deleted_by,
				t_message.message_deleted_at at time zone 'Asia/Jakarta' as message_deleted_at`).
		Where("message_status = ?", status).
		Order("t_message.message_id desc").
		Find(&loan_req_timeline).Error
	if err != nil {
		return &[]MessageData{}, err
	}
	return &loan_req_timeline, nil
}

func (p *Message) FindMessageDataByID(db *gorm.DB, pid uint64) (*Message, error) {
	var err error
	err = db.Debug().Model(&Message{}).Where("message_id = ?", pid).Take(&p).Error
	if err != nil {
		return &Message{}, err
	}
	return p, nil
}

func (p *Message) FindMessageByID(db *gorm.DB, pid uint64) (*MessageData, error) {
	var err error
	loan_req_timeline := MessageData{}
	err = db.Debug().Model(&MessageData{}).Limit(100).
		Select(`t_message.message_id,
				t_message.message_from,
				t_message.message_sender,
				t_message.message_to,
				t_message.message_link,
				t_message.message_title,
				t_message.message_detail,
				t_message.message_status,
				t_message.message_created_by,
				t_message.message_created_at at time zone 'Asia/Jakarta' as message_created_at,
				t_message.message_deleted_by,
				t_message.message_deleted_at at time zone 'Asia/Jakarta' as message_deleted_at`).
		Where("message_id = ?", pid).
		Order("t_message.message_id desc").
		Find(&loan_req_timeline).Error
	if err != nil {
		return &MessageData{}, err
	}
	return &loan_req_timeline, nil
}

func (p *Message) FindMessageStatusByID(db *gorm.DB, pid uint64, status uint64) (*MessageData, error) {
	var err error
	loan_req_timeline := MessageData{}
	err = db.Debug().Model(&MessageData{}).Limit(100).
		Select(`t_message.message_id,
				t_message.message_from,
				t_message.message_sender,
				t_message.message_to,
				t_message.message_link,
				t_message.message_title,
				t_message.message_detail,
				t_message.message_status,
				t_message.message_created_by,
				t_message.message_created_at at time zone 'Asia/Jakarta' as message_created_at,
				t_message.message_deleted_by,
				t_message.message_deleted_at at time zone 'Asia/Jakarta' as message_deleted_at`).
		Where("message_id = ? AND message_status = ?", pid, status).
		Order("t_message.message_id desc").
		Find(&loan_req_timeline).Error
	if err != nil {
		return &MessageData{}, err
	}
	return &loan_req_timeline, nil
}

func (p *Message) UpdateMessageStatus(db *gorm.DB, pid uint64) (*Message, error) {
	var err error
	db = db.Debug().Model(&Message{}).Where("message_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"message_status": p.MessageStatus,
		},
	)
	err = db.Debug().Model(&Message{}).Where("message_id = ?", pid).Error
	if err != nil {
		return &Message{}, err
	}
	return p, nil
}

func (p *Message) UpdateMessageStatusDelete(db *gorm.DB, pid uint64) (*Message, error) {
	var err error
	db = db.Debug().Model(&Message{}).Where("message_id = ?", pid).UpdateColumns(
		map[string]interface{}{
			"message_status":     2,
			"message_deleted_at": time.Now(),
		},
	)
	err = db.Debug().Model(&Message{}).Where("message_id = ?", pid).Error
	if err != nil {
		return &Message{}, err
	}
	return p, nil
}

func (p *Message) DeleteMessage(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&Message{}).Where("message_id = ?", pid).Take(&Message{}).Delete(&Message{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Message not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
