// Code generated by MockGen. DO NOT EDIT.
// Source: internal/storage/database/database.go

// Package database is a generated GoMock package.
package database

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/leguminosa/kounat/internal/entity"
)

// MockCharacterDB is a mock of CharacterDB interface.
type MockCharacterDB struct {
	ctrl     *gomock.Controller
	recorder *MockCharacterDBMockRecorder
}

// MockCharacterDBMockRecorder is the mock recorder for MockCharacterDB.
type MockCharacterDBMockRecorder struct {
	mock *MockCharacterDB
}

// NewMockCharacterDB creates a new mock instance.
func NewMockCharacterDB(ctrl *gomock.Controller) *MockCharacterDB {
	mock := &MockCharacterDB{ctrl: ctrl}
	mock.recorder = &MockCharacterDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCharacterDB) EXPECT() *MockCharacterDBMockRecorder {
	return m.recorder
}

// GetByID mocks base method.
func (m *MockCharacterDB) GetByID(ctx context.Context, id int) (*entity.Character, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*entity.Character)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockCharacterDBMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockCharacterDB)(nil).GetByID), ctx, id)
}