package dao

import (
	"github.com/liluhao/ginstarter/pkg/business"
	"time"
)

type Member struct {
	ID             string    `json:"id"`
	Email          string    `json:"email" binding:"required,email"`
	PasswordDigest string    `json:"-"`
	Name           string    `json:"name" binding:"required"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	DeletedAt      time.Time `pg:",soft_delete" json:"-"`
}

type MemberDAO interface {
	Create(member Member) (string, *business.Error)
	Get(memberID string) (Member, *business.Error)
	Update(member Member) *business.Error
	Delete(memberID string) *business.Error
}
