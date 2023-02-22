// Code generated by MockGen. DO NOT EDIT.
// Source: ./skus/credentials.go

// Package skus is a generated GoMock package.
package skus

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	go_uuid "github.com/satori/go.uuid"
	kafka_go "github.com/segmentio/kafka-go"
)

// MockOrderWorker is a mock of OrderWorker interface.
type MockOrderWorker struct {
	ctrl     *gomock.Controller
	recorder *MockOrderWorkerMockRecorder
}

// MockOrderWorkerMockRecorder is the mock recorder for MockOrderWorker.
type MockOrderWorkerMockRecorder struct {
	mock *MockOrderWorker
}

// NewMockOrderWorker creates a new mock instance.
func NewMockOrderWorker(ctrl *gomock.Controller) *MockOrderWorker {
	mock := &MockOrderWorker{ctrl: ctrl}
	mock.recorder = &MockOrderWorkerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderWorker) EXPECT() *MockOrderWorkerMockRecorder {
	return m.recorder
}

// SignOrderCreds mocks base method.
func (m *MockOrderWorker) SignOrderCreds(ctx context.Context, orderID go_uuid.UUID, issuer Issuer, blindedCreds []string) (*OrderCreds, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignOrderCreds", ctx, orderID, issuer, blindedCreds)
	ret0, _ := ret[0].(*OrderCreds)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignOrderCreds indicates an expected call of SignOrderCreds.
func (mr *MockOrderWorkerMockRecorder) SignOrderCreds(ctx, orderID, issuer, blindedCreds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignOrderCreds", reflect.TypeOf((*MockOrderWorker)(nil).SignOrderCreds), ctx, orderID, issuer, blindedCreds)
}

// MockSigningRequestWriter is a mock of SigningRequestWriter interface.
type MockSigningRequestWriter struct {
	ctrl     *gomock.Controller
	recorder *MockSigningRequestWriterMockRecorder
}

// MockSigningRequestWriterMockRecorder is the mock recorder for MockSigningRequestWriter.
type MockSigningRequestWriterMockRecorder struct {
	mock *MockSigningRequestWriter
}

// NewMockSigningRequestWriter creates a new mock instance.
func NewMockSigningRequestWriter(ctrl *gomock.Controller) *MockSigningRequestWriter {
	mock := &MockSigningRequestWriter{ctrl: ctrl}
	mock.recorder = &MockSigningRequestWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSigningRequestWriter) EXPECT() *MockSigningRequestWriterMockRecorder {
	return m.recorder
}

// WriteMessage mocks base method.
func (m *MockSigningRequestWriter) WriteMessage(ctx context.Context, message []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteMessage", ctx, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteMessage indicates an expected call of WriteMessage.
func (mr *MockSigningRequestWriterMockRecorder) WriteMessage(ctx, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteMessage", reflect.TypeOf((*MockSigningRequestWriter)(nil).WriteMessage), ctx, message)
}

// WriteMessages mocks base method.
func (m *MockSigningRequestWriter) WriteMessages(ctx context.Context, messages []SigningOrderRequestOutbox) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteMessages", ctx, messages)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteMessages indicates an expected call of WriteMessages.
func (mr *MockSigningRequestWriterMockRecorder) WriteMessages(ctx, messages interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteMessages", reflect.TypeOf((*MockSigningRequestWriter)(nil).WriteMessages), ctx, messages)
}

// MockDecoder is a mock of Decoder interface.
type MockDecoder struct {
	ctrl     *gomock.Controller
	recorder *MockDecoderMockRecorder
}

// MockDecoderMockRecorder is the mock recorder for MockDecoder.
type MockDecoderMockRecorder struct {
	mock *MockDecoder
}

// NewMockDecoder creates a new mock instance.
func NewMockDecoder(ctrl *gomock.Controller) *MockDecoder {
	mock := &MockDecoder{ctrl: ctrl}
	mock.recorder = &MockDecoderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDecoder) EXPECT() *MockDecoderMockRecorder {
	return m.recorder
}

// Decode mocks base method.
func (m *MockDecoder) Decode(message kafka_go.Message) (*SigningOrderResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", message)
	ret0, _ := ret[0].(*SigningOrderResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decode indicates an expected call of Decode.
func (mr *MockDecoderMockRecorder) Decode(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockDecoder)(nil).Decode), message)
}