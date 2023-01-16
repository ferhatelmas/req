// Code generated by MockGen. DO NOT EDIT.
// Source: connection.go

// Package quic is a generated GoMock package.
package quic

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	protocol "github.com/imroc/req/v3/internal/quic-go/protocol"
	wire "github.com/imroc/req/v3/internal/quic-go/wire"
)

// MockStreamManager is a mock of StreamManager interface.
type MockStreamManager struct {
	ctrl     *gomock.Controller
	recorder *MockStreamManagerMockRecorder
}

// MockStreamManagerMockRecorder is the mock recorder for MockStreamManager.
type MockStreamManagerMockRecorder struct {
	mock *MockStreamManager
}

// NewMockStreamManager creates a new mock instance.
func NewMockStreamManager(ctrl *gomock.Controller) *MockStreamManager {
	mock := &MockStreamManager{ctrl: ctrl}
	mock.recorder = &MockStreamManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamManager) EXPECT() *MockStreamManagerMockRecorder {
	return m.recorder
}

// AcceptStream mocks base method.
func (m *MockStreamManager) AcceptStream(arg0 context.Context) (Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptStream", arg0)
	ret0, _ := ret[0].(Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptStream indicates an expected call of AcceptStream.
func (mr *MockStreamManagerMockRecorder) AcceptStream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptStream", reflect.TypeOf((*MockStreamManager)(nil).AcceptStream), arg0)
}

// AcceptUniStream mocks base method.
func (m *MockStreamManager) AcceptUniStream(arg0 context.Context) (ReceiveStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptUniStream", arg0)
	ret0, _ := ret[0].(ReceiveStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptUniStream indicates an expected call of AcceptUniStream.
func (mr *MockStreamManagerMockRecorder) AcceptUniStream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptUniStream", reflect.TypeOf((*MockStreamManager)(nil).AcceptUniStream), arg0)
}

// CloseWithError mocks base method.
func (m *MockStreamManager) CloseWithError(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CloseWithError", arg0)
}

// CloseWithError indicates an expected call of CloseWithError.
func (mr *MockStreamManagerMockRecorder) CloseWithError(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseWithError", reflect.TypeOf((*MockStreamManager)(nil).CloseWithError), arg0)
}

// DeleteStream mocks base method.
func (m *MockStreamManager) DeleteStream(arg0 protocol.StreamID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStream", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStream indicates an expected call of DeleteStream.
func (mr *MockStreamManagerMockRecorder) DeleteStream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStream", reflect.TypeOf((*MockStreamManager)(nil).DeleteStream), arg0)
}

// GetOrOpenReceiveStream mocks base method.
func (m *MockStreamManager) GetOrOpenReceiveStream(arg0 protocol.StreamID) (receiveStreamI, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrOpenReceiveStream", arg0)
	ret0, _ := ret[0].(receiveStreamI)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrOpenReceiveStream indicates an expected call of GetOrOpenReceiveStream.
func (mr *MockStreamManagerMockRecorder) GetOrOpenReceiveStream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrOpenReceiveStream", reflect.TypeOf((*MockStreamManager)(nil).GetOrOpenReceiveStream), arg0)
}

// GetOrOpenSendStream mocks base method.
func (m *MockStreamManager) GetOrOpenSendStream(arg0 protocol.StreamID) (sendStreamI, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrOpenSendStream", arg0)
	ret0, _ := ret[0].(sendStreamI)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrOpenSendStream indicates an expected call of GetOrOpenSendStream.
func (mr *MockStreamManagerMockRecorder) GetOrOpenSendStream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrOpenSendStream", reflect.TypeOf((*MockStreamManager)(nil).GetOrOpenSendStream), arg0)
}

// HandleMaxStreamsFrame mocks base method.
func (m *MockStreamManager) HandleMaxStreamsFrame(arg0 *wire.MaxStreamsFrame) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleMaxStreamsFrame", arg0)
}

// HandleMaxStreamsFrame indicates an expected call of HandleMaxStreamsFrame.
func (mr *MockStreamManagerMockRecorder) HandleMaxStreamsFrame(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleMaxStreamsFrame", reflect.TypeOf((*MockStreamManager)(nil).HandleMaxStreamsFrame), arg0)
}

// OpenStream mocks base method.
func (m *MockStreamManager) OpenStream() (Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenStream")
	ret0, _ := ret[0].(Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenStream indicates an expected call of OpenStream.
func (mr *MockStreamManagerMockRecorder) OpenStream() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenStream", reflect.TypeOf((*MockStreamManager)(nil).OpenStream))
}

// OpenStreamSync mocks base method.
func (m *MockStreamManager) OpenStreamSync(arg0 context.Context) (Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenStreamSync", arg0)
	ret0, _ := ret[0].(Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenStreamSync indicates an expected call of OpenStreamSync.
func (mr *MockStreamManagerMockRecorder) OpenStreamSync(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenStreamSync", reflect.TypeOf((*MockStreamManager)(nil).OpenStreamSync), arg0)
}

// OpenUniStream mocks base method.
func (m *MockStreamManager) OpenUniStream() (SendStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenUniStream")
	ret0, _ := ret[0].(SendStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenUniStream indicates an expected call of OpenUniStream.
func (mr *MockStreamManagerMockRecorder) OpenUniStream() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenUniStream", reflect.TypeOf((*MockStreamManager)(nil).OpenUniStream))
}

// OpenUniStreamSync mocks base method.
func (m *MockStreamManager) OpenUniStreamSync(arg0 context.Context) (SendStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenUniStreamSync", arg0)
	ret0, _ := ret[0].(SendStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenUniStreamSync indicates an expected call of OpenUniStreamSync.
func (mr *MockStreamManagerMockRecorder) OpenUniStreamSync(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenUniStreamSync", reflect.TypeOf((*MockStreamManager)(nil).OpenUniStreamSync), arg0)
}

// ResetFor0RTT mocks base method.
func (m *MockStreamManager) ResetFor0RTT() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ResetFor0RTT")
}

// ResetFor0RTT indicates an expected call of ResetFor0RTT.
func (mr *MockStreamManagerMockRecorder) ResetFor0RTT() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetFor0RTT", reflect.TypeOf((*MockStreamManager)(nil).ResetFor0RTT))
}

// UpdateLimits mocks base method.
func (m *MockStreamManager) UpdateLimits(arg0 *wire.TransportParameters) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateLimits", arg0)
}

// UpdateLimits indicates an expected call of UpdateLimits.
func (mr *MockStreamManagerMockRecorder) UpdateLimits(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLimits", reflect.TypeOf((*MockStreamManager)(nil).UpdateLimits), arg0)
}

// UseResetMaps mocks base method.
func (m *MockStreamManager) UseResetMaps() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UseResetMaps")
}

// UseResetMaps indicates an expected call of UseResetMaps.
func (mr *MockStreamManagerMockRecorder) UseResetMaps() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UseResetMaps", reflect.TypeOf((*MockStreamManager)(nil).UseResetMaps))
}
