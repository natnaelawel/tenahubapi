package entity

import (
	"time"
)

// User represents user data
type User struct {
	ID          uint   `json:"id"`
	FirstName   string `json:"firstname" gorm:"varchar(255);not null"`
	LastName    string `json:"lastname" gorm:"varchar(255);not null"`
	Email       string `json:"email" gorm:"unique;varchar(255);not null"`
	Password    string `json:"password" gorm:"varchar(255);not null"`
	ProfilePic  string `json:"profilepic" gorm:"varchar(255);"`
	PhoneNumber string `json:"phonenum" gorm:"unique;varchar(255);not null"`
	Role        string `json:"role" gorm:"varchar(255);not null"`
}

type Agent struct {
	ID          uint `json:"id" `
	FirstName   string `json:"firstname" gorm:"type:varchar(255);not null"`
	LastName    string `json:"lastname" gorm:"type:varchar(255);not null"`
	UserName    string `json:"username" gorm:"type:varchar(255);not null"`
	Email       string `json:"email" gorm:"type:varchar(255);not null"`
	Password    string `json:"password" gorm:"type:varchar(255);not null"`
	ProfilePic  string `json:"profilepic" gorm:"type:varchar(255);"`
	PhoneNumber string `json:"phonenum" gorm:"type:varchar(255);not null"`
}



type Admin struct {
	ID          uint `json:"id"`
	FirstName   string `json:"firstname" gorm:"type:varchar(255);not null"`
	LastName    string `json:"lastname" gorm:"type:varchar(255);not null"`
	UserName    string `json:"username" gorm:"type:varchar(255);not null"`
	Email       string `json:"email" gorm:"type:varchar(255);not null"`
	Password    string `json:"password" gorm:"type:varchar(255);not null"`
	ProfilePic  string `json:"profilepic" gorm:"type:varchar(255);"`
	PhoneNumber string `json:"phonenum" gorm:"type:varchar(255);not null"`
}


// Rating represents users rating
type Rating struct {
	ID             uint
	Value          uint
	UserID         uint
	HealthCenterID uint
	PlacedAt       time.Time `sql:"DEFAULT:current_timestamp"`
}

// Comment represents users comment
type Comment struct {
	ID             uint      `json:"id"`
	Comment        string    `json:"comment" gorm:"varchar(255); not null"`
	UserID         uint      `json:"userid"`
	HealthCenterID uint      `json:"healthcenterid"`
	Rating         uint      `json:"rating"`
	PlacedAt       time.Time `json:"placedat" sql:"DEFAULT:current_timestamp"`
}

// UserComment joins feedback givers first name with feedback
type UserComment struct {
	FirstName string
	Comment
}

// HealthCenter represents health centers data
type HealthCenter struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name" gorm:"unique;varchar(255); not null"`
	Email       string    `json:"email" gorm:"unique;varchar(255); not null"`
	PhoneNumber string    `json:"phonenumber" gorm:"unique;varchar(255); not null"`
	Password    string    `json:"password" gorm:"varchar(255);not null"`
	City        string    `json:"city" gorm:"varchar(255); not null"`
	ProfilePic  string    `json:"profilepic" gorm:"varchar(255);"`
	AgentID     uint      `json:"agentid"`
	Services    []Service `gorm:"foreignkey:HealthCenterID"`
}

// admating represents healthcenters with rating
type Hcrating struct{
	HealthCenter
	Rating float64 `json:"rating"`
}

// Service represents health centers services
type Service struct {
	ID             uint   `json:"id"`
	Name           string `json:"name" gorm:"varchar(255); not null"`
	Description    string `json:"description" gorm:"varchar(255); not null"`
	HealthCenterID uint   `json:"healthcenterid"`
	// HealthCenter   HealthCenter
	Status string `json:"status" gorm:"varchar(255); not null;default:'pending'"`
}

//Session represents login user session
type Session struct {
	ID         uint
	UUID       string `gorm:"type:varchar(255);not null" json:"uuid"`
	Expires    int64  `gorm:"type:varchar(255);not null" json:"expires"`
	SigningKey []byte `gorm:"type:varchar(255);not null" json:"signing_key"`
}