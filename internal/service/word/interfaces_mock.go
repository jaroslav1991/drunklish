// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package word is a generated GoMock package.
package word

import (
	model "drunklish/internal/model"
	dto "drunklish/internal/service/word/dto"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepository) Create(word dto.CreateWordRequest) (*model.Word, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", word)
	ret0, _ := ret[0].(*model.Word)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(word interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), word)
}

// DeleteWord mocks base method.
func (m *MockRepository) DeleteWord(word dto.RequestForDeletingWord) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWord", word)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWord indicates an expected call of DeleteWord.
func (mr *MockRepositoryMockRecorder) DeleteWord(word interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWord", reflect.TypeOf((*MockRepository)(nil).DeleteWord), word)
}

// GetWords mocks base method.
func (m *MockRepository) GetWords(word dto.RequestForGettingWord) (*dto.ResponseWords, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWords", word)
	ret0, _ := ret[0].(*dto.ResponseWords)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWords indicates an expected call of GetWords.
func (mr *MockRepositoryMockRecorder) GetWords(word interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWords", reflect.TypeOf((*MockRepository)(nil).GetWords), word)
}

// GetWordsByCreated mocks base method.
func (m *MockRepository) GetWordsByCreated(userId int64, createdAt time.Time) (*model.Word, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWordsByCreated", userId, createdAt)
	ret0, _ := ret[0].(*model.Word)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWordsByCreated indicates an expected call of GetWordsByCreated.
func (mr *MockRepositoryMockRecorder) GetWordsByCreated(userId, createdAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWordsByCreated", reflect.TypeOf((*MockRepository)(nil).GetWordsByCreated), userId, createdAt)
}
