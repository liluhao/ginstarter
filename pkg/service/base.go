package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liluhao/ginstarter/pkg/business"
	"github.com/liluhao/ginstarter/pkg/dao"
)

type BaseService struct {
	memberDAO dao.MemberDAO
}

func NewService(memberDAO dao.MemberDAO) *BaseService {
	return &BaseService{memberDAO: memberDAO}
}

//请求方法错误
func (s *BaseService) HandleMethodNotAllowed(c *gin.Context) {
	s.responseWithError(c, business.NewError(business.MethodNowAllowed, http.StatusMethodNotAllowed, "http method not allowed", nil))
}

//请求路径错误
func (s *BaseService) HandlePathNotFound(c *gin.Context) {
	s.responseWithError(c, business.NewError(business.PathNotFound, http.StatusNotFound, "http path not found", nil))
}

func (s *BaseService) responseWithError(c *gin.Context, businessError *business.Error) {
	c.Abort()
	c.Error(businessError)
}

func (s *BaseService) responseWithSuccess(c *gin.Context, businessSuccess *business.Success) {
	c.Set("success", businessSuccess)
}
