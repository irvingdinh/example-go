// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	repository "github.com/irvingdinh/example-go/internal/internal/service/repository"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// TaskRepository provides a mock function with given fields:
func (_m *Repository) TaskRepository() repository.TaskRepository {
	ret := _m.Called()

	var r0 repository.TaskRepository
	if rf, ok := ret.Get(0).(func() repository.TaskRepository); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(repository.TaskRepository)
		}
	}

	return r0
}
