// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	dto "daveslist/internal/core/domain/dto"
	model "daveslist/internal/core/domain/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCategoryRepository is a mock of CategoryRepository interface.
type MockCategoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryRepositoryMockRecorder
}

// MockCategoryRepositoryMockRecorder is the mock recorder for MockCategoryRepository.
type MockCategoryRepositoryMockRecorder struct {
	mock *MockCategoryRepository
}

// NewMockCategoryRepository creates a new mock instance.
func NewMockCategoryRepository(ctrl *gomock.Controller) *MockCategoryRepository {
	mock := &MockCategoryRepository{ctrl: ctrl}
	mock.recorder = &MockCategoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategoryRepository) EXPECT() *MockCategoryRepositoryMockRecorder {
	return m.recorder
}

// CountByQuery mocks base method.
func (m *MockCategoryRepository) CountByQuery(ctx context.Context, query *dto.CategoryQuery) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountByQuery", ctx, query)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountByQuery indicates an expected call of CountByQuery.
func (mr *MockCategoryRepositoryMockRecorder) CountByQuery(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountByQuery", reflect.TypeOf((*MockCategoryRepository)(nil).CountByQuery), ctx, query)
}

// FindByQuery mocks base method.
func (m *MockCategoryRepository) FindByQuery(ctx context.Context, query *dto.CategoryQuery) (dto.CategoryListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByQuery", ctx, query)
	ret0, _ := ret[0].(dto.CategoryListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByQuery indicates an expected call of FindByQuery.
func (mr *MockCategoryRepositoryMockRecorder) FindByQuery(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByQuery", reflect.TypeOf((*MockCategoryRepository)(nil).FindByQuery), ctx, query)
}

// Insert mocks base method.
func (m *MockCategoryRepository) Insert(ctx context.Context, data *model.Category) (*dto.CategoryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, data)
	ret0, _ := ret[0].(*dto.CategoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockCategoryRepositoryMockRecorder) Insert(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockCategoryRepository)(nil).Insert), ctx, data)
}

// UpdateOneByID mocks base method.
func (m *MockCategoryRepository) UpdateOneByID(ctx context.Context, id string, update *model.UpdateCategory) (*dto.CategoryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOneByID", ctx, id, update)
	ret0, _ := ret[0].(*dto.CategoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOneByID indicates an expected call of UpdateOneByID.
func (mr *MockCategoryRepositoryMockRecorder) UpdateOneByID(ctx, id, update interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOneByID", reflect.TypeOf((*MockCategoryRepository)(nil).UpdateOneByID), ctx, id, update)
}

// MockListingRepository is a mock of ListingRepository interface.
type MockListingRepository struct {
	ctrl     *gomock.Controller
	recorder *MockListingRepositoryMockRecorder
}

// MockListingRepositoryMockRecorder is the mock recorder for MockListingRepository.
type MockListingRepositoryMockRecorder struct {
	mock *MockListingRepository
}

// NewMockListingRepository creates a new mock instance.
func NewMockListingRepository(ctrl *gomock.Controller) *MockListingRepository {
	mock := &MockListingRepository{ctrl: ctrl}
	mock.recorder = &MockListingRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockListingRepository) EXPECT() *MockListingRepositoryMockRecorder {
	return m.recorder
}

// CountByQuery mocks base method.
func (m *MockListingRepository) CountByQuery(ctx context.Context, query *dto.ListingQuery) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountByQuery", ctx, query)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountByQuery indicates an expected call of CountByQuery.
func (mr *MockListingRepositoryMockRecorder) CountByQuery(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountByQuery", reflect.TypeOf((*MockListingRepository)(nil).CountByQuery), ctx, query)
}

// FindByQuery mocks base method.
func (m *MockListingRepository) FindByQuery(ctx context.Context, query *dto.ListingQuery) (dto.ListingListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByQuery", ctx, query)
	ret0, _ := ret[0].(dto.ListingListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByQuery indicates an expected call of FindByQuery.
func (mr *MockListingRepositoryMockRecorder) FindByQuery(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByQuery", reflect.TypeOf((*MockListingRepository)(nil).FindByQuery), ctx, query)
}

// FindOneByID mocks base method.
func (m *MockListingRepository) FindOneByID(ctx context.Context, id string) (*dto.ListingResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByID", ctx, id)
	ret0, _ := ret[0].(*dto.ListingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByID indicates an expected call of FindOneByID.
func (mr *MockListingRepositoryMockRecorder) FindOneByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByID", reflect.TypeOf((*MockListingRepository)(nil).FindOneByID), ctx, id)
}

// Insert mocks base method.
func (m *MockListingRepository) Insert(ctx context.Context, data *model.Listing) (*dto.ListingResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, data)
	ret0, _ := ret[0].(*dto.ListingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockListingRepositoryMockRecorder) Insert(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockListingRepository)(nil).Insert), ctx, data)
}

// UpdateByQuery mocks base method.
func (m *MockListingRepository) UpdateByQuery(ctx context.Context, query *dto.ListingQuery, update *model.UpdateListing) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByQuery", ctx, query, update)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateByQuery indicates an expected call of UpdateByQuery.
func (mr *MockListingRepositoryMockRecorder) UpdateByQuery(ctx, query, update interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByQuery", reflect.TypeOf((*MockListingRepository)(nil).UpdateByQuery), ctx, query, update)
}

// UpdateOneByID mocks base method.
func (m *MockListingRepository) UpdateOneByID(ctx context.Context, id string, update *model.UpdateListing) (*dto.ListingResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOneByID", ctx, id, update)
	ret0, _ := ret[0].(*dto.ListingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOneByID indicates an expected call of UpdateOneByID.
func (mr *MockListingRepositoryMockRecorder) UpdateOneByID(ctx, id, update interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOneByID", reflect.TypeOf((*MockListingRepository)(nil).UpdateOneByID), ctx, id, update)
}

// MockReplyListingRepository is a mock of ReplyListingRepository interface.
type MockReplyListingRepository struct {
	ctrl     *gomock.Controller
	recorder *MockReplyListingRepositoryMockRecorder
}

// MockReplyListingRepositoryMockRecorder is the mock recorder for MockReplyListingRepository.
type MockReplyListingRepositoryMockRecorder struct {
	mock *MockReplyListingRepository
}

// NewMockReplyListingRepository creates a new mock instance.
func NewMockReplyListingRepository(ctrl *gomock.Controller) *MockReplyListingRepository {
	mock := &MockReplyListingRepository{ctrl: ctrl}
	mock.recorder = &MockReplyListingRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReplyListingRepository) EXPECT() *MockReplyListingRepositoryMockRecorder {
	return m.recorder
}

// CountByQuery mocks base method.
func (m *MockReplyListingRepository) CountByQuery(ctx context.Context, query *dto.ReplyListingQuery) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountByQuery", ctx, query)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountByQuery indicates an expected call of CountByQuery.
func (mr *MockReplyListingRepositoryMockRecorder) CountByQuery(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountByQuery", reflect.TypeOf((*MockReplyListingRepository)(nil).CountByQuery), ctx, query)
}

// FindByQuery mocks base method.
func (m *MockReplyListingRepository) FindByQuery(ctx context.Context, query *dto.ReplyListingQuery) (dto.ReplyListingListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByQuery", ctx, query)
	ret0, _ := ret[0].(dto.ReplyListingListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByQuery indicates an expected call of FindByQuery.
func (mr *MockReplyListingRepositoryMockRecorder) FindByQuery(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByQuery", reflect.TypeOf((*MockReplyListingRepository)(nil).FindByQuery), ctx, query)
}

// Insert mocks base method.
func (m *MockReplyListingRepository) Insert(ctx context.Context, data *model.ReplyListing) (*dto.ReplyListingResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, data)
	ret0, _ := ret[0].(*dto.ReplyListingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockReplyListingRepositoryMockRecorder) Insert(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockReplyListingRepository)(nil).Insert), ctx, data)
}

// MockMessageRepository is a mock of MessageRepository interface.
type MockMessageRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMessageRepositoryMockRecorder
}

// MockMessageRepositoryMockRecorder is the mock recorder for MockMessageRepository.
type MockMessageRepositoryMockRecorder struct {
	mock *MockMessageRepository
}

// NewMockMessageRepository creates a new mock instance.
func NewMockMessageRepository(ctrl *gomock.Controller) *MockMessageRepository {
	mock := &MockMessageRepository{ctrl: ctrl}
	mock.recorder = &MockMessageRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageRepository) EXPECT() *MockMessageRepositoryMockRecorder {
	return m.recorder
}

// CountByQuery mocks base method.
func (m *MockMessageRepository) CountByQuery(ctx context.Context, query *dto.MessageQuery) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountByQuery", ctx, query)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountByQuery indicates an expected call of CountByQuery.
func (mr *MockMessageRepositoryMockRecorder) CountByQuery(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountByQuery", reflect.TypeOf((*MockMessageRepository)(nil).CountByQuery), ctx, query)
}

// FindByQuery mocks base method.
func (m *MockMessageRepository) FindByQuery(ctx context.Context, query *dto.MessageQuery) (dto.MessageListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByQuery", ctx, query)
	ret0, _ := ret[0].(dto.MessageListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByQuery indicates an expected call of FindByQuery.
func (mr *MockMessageRepositoryMockRecorder) FindByQuery(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByQuery", reflect.TypeOf((*MockMessageRepository)(nil).FindByQuery), ctx, query)
}

// Insert mocks base method.
func (m *MockMessageRepository) Insert(ctx context.Context, data *model.Message) (*dto.MessageResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, data)
	ret0, _ := ret[0].(*dto.MessageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockMessageRepositoryMockRecorder) Insert(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockMessageRepository)(nil).Insert), ctx, data)
}
