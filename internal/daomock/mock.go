package daomock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	business "github.com/liluhao/ginstarter/pkg/business"
	dao "github.com/liluhao/ginstarter/pkg/dao"
)

//MockMemberDAO是MemberDAO接口的一个mock
type MockMemberDAO struct {
	ctrl     *gomock.Controller
	recorder *MockMemberDAOMockRecorder
}

//MockMemberDAOMockRecorder是MemberDAO接口的一个mock recorder
type MockMemberDAOMockRecorder struct {
	mock *MockMemberDAO
}

//创造一个新的mock实例
func NewMockMemberDAO(ctrl *gomock.Controller) *MockMemberDAO {
	mock := &MockMemberDAO{ctrl: ctrl}
	mock.recorder = &MockMemberDAOMockRecorder{mock}
	return mock
}

//返回一个对象，该对象允许调用方指示预期的使用。
func (m *MockMemberDAO) EXPECT() *MockMemberDAOMockRecorder {
	return m.recorder
}

//mock基本方法
func (m *MockMemberDAO) Create(arg0 dao.Member) (string, *business.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*business.Error)
	return ret0, ret1
}

//指示一个Create的预期调用
func (mr *MockMemberDAOMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMemberDAO)(nil).Create), arg0)
}

//mock基本方法
func (m *MockMemberDAO) Delete(arg0 string) *business.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(*business.Error)
	return ret0
}

//指示一个Delete的预期调用
func (mr *MockMemberDAOMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMemberDAO)(nil).Delete), arg0)
}

//mock基本方法
func (m *MockMemberDAO) Get(arg0 string) (dao.Member, *business.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(dao.Member)
	ret1, _ := ret[1].(*business.Error)
	return ret0, ret1
}

//指示一个Get的预期调用
func (mr *MockMemberDAOMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMemberDAO)(nil).Get), arg0)
}

//mock基本方法
func (m *MockMemberDAO) Update(arg0 dao.Member) *business.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*business.Error)
	return ret0
}

//指示一个Update的预期调用
func (mr *MockMemberDAOMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMemberDAO)(nil).Update), arg0)
}
