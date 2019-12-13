// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/paaguti/quic-go (interfaces: ReceiveStreamI)

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	protocol "github.com/paaguti/quic-go/internal/protocol"
	wire "github.com/paaguti/quic-go/internal/wire"
)

// MockReceiveStreamI is a mock of ReceiveStreamI interface
type MockReceiveStreamI struct {
	ctrl     *gomock.Controller
	recorder *MockReceiveStreamIMockRecorder
}

// MockReceiveStreamIMockRecorder is the mock recorder for MockReceiveStreamI
type MockReceiveStreamIMockRecorder struct {
	mock *MockReceiveStreamI
}

// NewMockReceiveStreamI creates a new mock instance
func NewMockReceiveStreamI(ctrl *gomock.Controller) *MockReceiveStreamI {
	mock := &MockReceiveStreamI{ctrl: ctrl}
	mock.recorder = &MockReceiveStreamIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReceiveStreamI) EXPECT() *MockReceiveStreamIMockRecorder {
	return m.recorder
}

// CancelRead mocks base method
func (m *MockReceiveStreamI) CancelRead(arg0 protocol.ApplicationErrorCode) error {
	ret := m.ctrl.Call(m, "CancelRead", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelRead indicates an expected call of CancelRead
func (mr *MockReceiveStreamIMockRecorder) CancelRead(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelRead", reflect.TypeOf((*MockReceiveStreamI)(nil).CancelRead), arg0)
}

// Read mocks base method
func (m *MockReceiveStreamI) Read(arg0 []byte) (int, error) {
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockReceiveStreamIMockRecorder) Read(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockReceiveStreamI)(nil).Read), arg0)
}

// SetReadDeadline mocks base method
func (m *MockReceiveStreamI) SetReadDeadline(arg0 time.Time) error {
	ret := m.ctrl.Call(m, "SetReadDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReadDeadline indicates an expected call of SetReadDeadline
func (mr *MockReceiveStreamIMockRecorder) SetReadDeadline(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadDeadline", reflect.TypeOf((*MockReceiveStreamI)(nil).SetReadDeadline), arg0)
}

// StreamID mocks base method
func (m *MockReceiveStreamI) StreamID() protocol.StreamID {
	ret := m.ctrl.Call(m, "StreamID")
	ret0, _ := ret[0].(protocol.StreamID)
	return ret0
}

// StreamID indicates an expected call of StreamID
func (mr *MockReceiveStreamIMockRecorder) StreamID() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamID", reflect.TypeOf((*MockReceiveStreamI)(nil).StreamID))
}

// closeForShutdown mocks base method
func (m *MockReceiveStreamI) closeForShutdown(arg0 error) {
	m.ctrl.Call(m, "closeForShutdown", arg0)
}

// closeForShutdown indicates an expected call of closeForShutdown
func (mr *MockReceiveStreamIMockRecorder) closeForShutdown(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "closeForShutdown", reflect.TypeOf((*MockReceiveStreamI)(nil).closeForShutdown), arg0)
}

// getWindowUpdate mocks base method
func (m *MockReceiveStreamI) getWindowUpdate() protocol.ByteCount {
	ret := m.ctrl.Call(m, "getWindowUpdate")
	ret0, _ := ret[0].(protocol.ByteCount)
	return ret0
}

// getWindowUpdate indicates an expected call of getWindowUpdate
func (mr *MockReceiveStreamIMockRecorder) getWindowUpdate() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getWindowUpdate", reflect.TypeOf((*MockReceiveStreamI)(nil).getWindowUpdate))
}

// handleRstStreamFrame mocks base method
func (m *MockReceiveStreamI) handleRstStreamFrame(arg0 *wire.RstStreamFrame) error {
	ret := m.ctrl.Call(m, "handleRstStreamFrame", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// handleRstStreamFrame indicates an expected call of handleRstStreamFrame
func (mr *MockReceiveStreamIMockRecorder) handleRstStreamFrame(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleRstStreamFrame", reflect.TypeOf((*MockReceiveStreamI)(nil).handleRstStreamFrame), arg0)
}

// handleStreamFrame mocks base method
func (m *MockReceiveStreamI) handleStreamFrame(arg0 *wire.StreamFrame) error {
	ret := m.ctrl.Call(m, "handleStreamFrame", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// handleStreamFrame indicates an expected call of handleStreamFrame
func (mr *MockReceiveStreamIMockRecorder) handleStreamFrame(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleStreamFrame", reflect.TypeOf((*MockReceiveStreamI)(nil).handleStreamFrame), arg0)
}
