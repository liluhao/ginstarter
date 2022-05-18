package daomock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	business "github.com/liluhao/ginstarter/pkg/business"
	dao "github.com/liluhao/ginstarter/pkg/dao"
)

// MockMemberDAO is a mock of MemberDAO interface.
type MockMemberDAO struct {
	ctrl     *gomock.Controller
	recorder *MockMemberDAOMockRecorder
}

// MockMemberDAOMockRecorder is the mock recorder for MockMemberDAO.
type MockMemberDAOMockRecorder struct {
	mock *MockMemberDAO
}

// NewMockMemberDAO creates a new mock instance.
func NewMockMemberDAO(ctrl *gomock.Controller) *MockMemberDAO {
	mock := &MockMemberDAO{ctrl: ctrl}
	mock.recorder = &MockMemberDAOMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMemberDAO) EXPECT() *MockMemberDAOMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockMemberDAO) Create(arg0 dao.Member) (string, *business.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*business.Error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMemberDAOMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMemberDAO)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockMemberDAO) Delete(arg0 string) *business.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(*business.Error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMemberDAOMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMemberDAO)(nil).Delete), arg0)
}

// Get mocks base method.
func (m *MockMemberDAO) Get(arg0 string) (dao.Member, *business.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(dao.Member)
	ret1, _ := ret[1].(*business.Error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockMemberDAOMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMemberDAO)(nil).Get), arg0)
}

// Update mocks base method.
func (m *MockMemberDAO) Update(arg0 dao.Member) *business.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*business.Error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockMemberDAOMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMemberDAO)(nil).Update), arg0)
}
