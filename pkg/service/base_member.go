package service

import (
	"net/http"

	"github.com/liluhao/ginstarter/pkg/business"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/liluhao/ginstarter/pkg/dao"
)

//创建用户
func (s *BaseService) CreateMember(c *gin.Context) {
	var request struct {
		dao.Member
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		s.responseWithError(c, business.NewError(business.InvalidBodyParse, http.StatusBadRequest, "invalid parse member's fields", err))
		return
	}
	request.ID = uuid.NewV4().String()
	if digest, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost); err != nil { //对密码进行加密
		s.responseWithError(c, business.NewError(business.Unknown, http.StatusBadRequest, "internal server error", err))
		return
	} else {
		request.PasswordDigest = string(digest) //上述返回的digest是[]byte，所以需要再进行转换
	}

	member := dao.Member{ID: request.ID, Name: request.Name, Email: request.Email, PasswordDigest: request.PasswordDigest}
	memberID, err := s.memberDAO.Create(member)
	if err != nil {
		s.responseWithError(c, business.NewError(business.Unknown, http.StatusBadRequest, "internal server error", err))
		return
	}
	s.responseWithSuccess(c, business.NewSuccess(http.StatusCreated, gin.H{"memberId": memberID})) //把ID返回
}

//获取用户消息
func (s *BaseService) GetMember(c *gin.Context) {
	memberID := c.Param("id")
	if memberID == "" {
		s.responseWithError(c, business.NewError(business.InvalidBodyParse, http.StatusBadRequest, "invalid memberID", nil)) //由于Param函数不会返回error，所以传入nil
		return
	}

	member, err := s.memberDAO.Get(memberID)
	if err != nil {
		s.responseWithError(c, err)
		return
	}
	s.responseWithSuccess(c, business.NewSuccess(http.StatusOK, member))
}

//更新用户信息
func (s *BaseService) UpdateMember(c *gin.Context) {
	var requestMemberID struct {
		ID string `uri:"id" binding:"uuid4"`
	}
	if err := c.ShouldBindUri(&requestMemberID); err != nil {
		s.responseWithError(c, business.NewError(business.InvalidBodyParse, http.StatusBadRequest, "invalid parse member's fields", err))
		return
	}

	memberID := requestMemberID.ID
	var requestMember struct {
		dao.Member
	}
	if err := c.ShouldBindJSON(&requestMember); err != nil {
		s.responseWithError(c, business.NewError(business.InvalidBodyParse, http.StatusBadRequest, "invalid parse member's fields", err))
		return
	}

	member := dao.Member{ID: memberID, Email: requestMember.Email, Name: requestMember.Name}
	err := s.memberDAO.Update(member)
	if err != nil {
		s.responseWithError(c, err)
		return
	}
	s.responseWithSuccess(c, business.NewSuccess(http.StatusNoContent, nil))
}

//删除用户信息
func (s *BaseService) DeleteMember(c *gin.Context) {
	var requestMemberID struct {
		ID string `uri:"id" binding:"uuid4"`
	}
	if err := c.ShouldBindUri(&requestMemberID); err != nil {
		s.responseWithError(c, business.NewError(business.InvalidBodyParse, http.StatusBadRequest, "invalid parse member's fields", err))
		return
	}

	memberID := requestMemberID.ID
	err := s.memberDAO.Delete(memberID)
	if err != nil {
		s.responseWithError(c, err)
		return
	}
	s.responseWithSuccess(c, business.NewSuccess(http.StatusNoContent, nil))
}
