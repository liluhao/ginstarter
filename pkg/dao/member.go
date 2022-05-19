package dao

import (
	"github.com/liluhao/ginstarter/pkg/business"
	"time"
)

//注意这里面是没有密码，存的是PasswordDigest更安全
type Member struct {
	ID             string    `json:"id"`
	Email          string    `json:"email" binding:"required,email"`
	PasswordDigest string    `json:"-"` //“-”是指Marshal序列化时候会忽略掉该字段
	Name           string    `json:"name" binding:"required"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	DeletedAt      time.Time `pg:",soft_delete" json:"-"`
}

//面向对象编程增删改查
type MemberDAO interface {
	Create(member Member) (string, *business.Error)
	Get(memberID string) (Member, *business.Error)
	Update(member Member) *business.Error
	Delete(memberID string) *business.Error
}
