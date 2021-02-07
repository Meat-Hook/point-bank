// Code generated by MockGen. DO NOT EDIT.
// Source: ../internal/api/rpc/pb/session.pb.go

// Package client_test is a generated GoMock package.
package client_test

import (
	context "context"
	reflect "reflect"

	pb "github.com/Meat-Hook/back-template/internal/modules/session/internal/api/rpc/pb"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockSessionClient is a mock of SessionClient interface
type MockSessionClient struct {
	ctrl     *gomock.Controller
	recorder *MockSessionClientMockRecorder
}

// MockSessionClientMockRecorder is the mock recorder for MockSessionClient
type MockSessionClientMockRecorder struct {
	mock *MockSessionClient
}

// NewMockSessionClient creates a new mock instance
func NewMockSessionClient(ctrl *gomock.Controller) *MockSessionClient {
	mock := &MockSessionClient{ctrl: ctrl}
	mock.recorder = &MockSessionClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSessionClient) EXPECT() *MockSessionClientMockRecorder {
	return m.recorder
}

// Session mocks base method
func (m *MockSessionClient) Session(ctx context.Context, in *pb.RequestSession, opts ...grpc.CallOption) (*pb.SessionInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Session", varargs...)
	ret0, _ := ret[0].(*pb.SessionInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Session indicates an expected call of Session
func (mr *MockSessionClientMockRecorder) Session(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Session", reflect.TypeOf((*MockSessionClient)(nil).Session), varargs...)
}

// MockSessionServer is a mock of SessionServer interface
type MockSessionServer struct {
	ctrl     *gomock.Controller
	recorder *MockSessionServerMockRecorder
}

// MockSessionServerMockRecorder is the mock recorder for MockSessionServer
type MockSessionServerMockRecorder struct {
	mock *MockSessionServer
}

// NewMockSessionServer creates a new mock instance
func NewMockSessionServer(ctrl *gomock.Controller) *MockSessionServer {
	mock := &MockSessionServer{ctrl: ctrl}
	mock.recorder = &MockSessionServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSessionServer) EXPECT() *MockSessionServerMockRecorder {
	return m.recorder
}

// Session mocks base method
func (m *MockSessionServer) Session(arg0 context.Context, arg1 *pb.RequestSession) (*pb.SessionInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Session", arg0, arg1)
	ret0, _ := ret[0].(*pb.SessionInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Session indicates an expected call of Session
func (mr *MockSessionServerMockRecorder) Session(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Session", reflect.TypeOf((*MockSessionServer)(nil).Session), arg0, arg1)
}