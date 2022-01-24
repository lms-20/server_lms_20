package model

import (
	"lms-api/internal/abstraction"
	"lms-api/pkg/constant"
	"lms-api/pkg/util/date"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type roleType string

const (
	student roleType = "student"
	admin   roleType = "admin"
)

type UserEntity struct {
	Name         string   `json:"fullname" validate:"required"`
	Email        string   `json:"emailAddress" validate:"required,email" gorm:"index:idx_user_email,unique"`
	Occupation   string   `json:"occupation" validate:"required"`
	PasswordHash string   `json:"-"`
	Password     string   `json:"password" validate:"required" gorm:"-"`
	Avatar       string   `json:"avatar"`
	Age          int      `json:"age"`
	PhoneNumber  string   `json:"phoneNumber"`
	Role         roleType `sql:"roleType" json:"role"`
}

type UserEntityModel struct {
	// abstraction
	abstraction.Entity

	// entity
	UserEntity

	// relationals

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (UserEntityModel) TableName() string {
	return "users"
}

func (m *UserEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	m.hashPassword()
	m.Password = ""
	m.Role = "student"
	return
}

func (m *UserEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.Name
	return
}

func (m *UserEntityModel) hashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	m.PasswordHash = string(bytes)
}

func (m *UserEntityModel) GenerateToken() (string, error) {
	var (
		jwtKey = os.Getenv("JWT_KEY")
	)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     m.ID,
		"email":  m.Email,
		"name":   m.Name,
		"avatar": m.Avatar,
		"role":   m.Role,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtKey))
	return tokenString, err
}
