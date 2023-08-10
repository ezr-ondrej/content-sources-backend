// Code generated by mockery v2.20.0. DO NOT EDIT.

package pulp_client

import (
	zest "github.com/content-services/zest/release/v2023"
	mock "github.com/stretchr/testify/mock"
)

// MockPulpClient is an autogenerated mock type for the PulpClient type
type MockPulpClient struct {
	mock.Mock
}

// CreateRpmDistribution provides a mock function with given fields: publicationHref, name, basePath
func (_m *MockPulpClient) CreateRpmDistribution(publicationHref string, name string, basePath string) (*string, error) {
	ret := _m.Called(publicationHref, name, basePath)

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (*string, error)); ok {
		return rf(publicationHref, name, basePath)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) *string); ok {
		r0 = rf(publicationHref, name, basePath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(publicationHref, name, basePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRpmPublication provides a mock function with given fields: versionHref
func (_m *MockPulpClient) CreateRpmPublication(versionHref string) (*string, error) {
	ret := _m.Called(versionHref)

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*string, error)); ok {
		return rf(versionHref)
	}
	if rf, ok := ret.Get(0).(func(string) *string); ok {
		r0 = rf(versionHref)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(versionHref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRpmRemote provides a mock function with given fields: name, url
func (_m *MockPulpClient) CreateRpmRemote(name string, url string) (*zest.RpmRpmRemoteResponse, error) {
	ret := _m.Called(name, url)

	var r0 *zest.RpmRpmRemoteResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*zest.RpmRpmRemoteResponse, error)); ok {
		return rf(name, url)
	}
	if rf, ok := ret.Get(0).(func(string, string) *zest.RpmRpmRemoteResponse); ok {
		r0 = rf(name, url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*zest.RpmRpmRemoteResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRpmRepository provides a mock function with given fields: uuid, rpmRemotePulpRef
func (_m *MockPulpClient) CreateRpmRepository(uuid string, rpmRemotePulpRef *string) (*zest.RpmRpmRepositoryResponse, error) {
	ret := _m.Called(uuid, rpmRemotePulpRef)

	var r0 *zest.RpmRpmRepositoryResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *string) (*zest.RpmRpmRepositoryResponse, error)); ok {
		return rf(uuid, rpmRemotePulpRef)
	}
	if rf, ok := ret.Get(0).(func(string, *string) *zest.RpmRpmRepositoryResponse); ok {
		r0 = rf(uuid, rpmRemotePulpRef)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*zest.RpmRpmRepositoryResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *string) error); ok {
		r1 = rf(uuid, rpmRemotePulpRef)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRpmDistribution provides a mock function with given fields: rpmDistributionHref
func (_m *MockPulpClient) DeleteRpmDistribution(rpmDistributionHref string) (string, error) {
	ret := _m.Called(rpmDistributionHref)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(rpmDistributionHref)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(rpmDistributionHref)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(rpmDistributionHref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRpmRemote provides a mock function with given fields: pulpHref
func (_m *MockPulpClient) DeleteRpmRemote(pulpHref string) (string, error) {
	ret := _m.Called(pulpHref)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(pulpHref)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(pulpHref)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(pulpHref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRpmRepository provides a mock function with given fields: rpmRepositoryHref
func (_m *MockPulpClient) DeleteRpmRepository(rpmRepositoryHref string) (string, error) {
	ret := _m.Called(rpmRepositoryHref)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(rpmRepositoryHref)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(rpmRepositoryHref)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(rpmRepositoryHref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRpmRepositoryVersion provides a mock function with given fields: href
func (_m *MockPulpClient) DeleteRpmRepositoryVersion(href string) (string, error) {
	ret := _m.Called(href)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(href)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(href)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(href)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindDistributionByPath provides a mock function with given fields: path
func (_m *MockPulpClient) FindDistributionByPath(path string) (*zest.RpmRpmDistributionResponse, error) {
	ret := _m.Called(path)

	var r0 *zest.RpmRpmDistributionResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*zest.RpmRpmDistributionResponse, error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) *zest.RpmRpmDistributionResponse); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*zest.RpmRpmDistributionResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindRpmPublicationByVersion provides a mock function with given fields: versionHref
func (_m *MockPulpClient) FindRpmPublicationByVersion(versionHref string) (*zest.RpmRpmPublicationResponse, error) {
	ret := _m.Called(versionHref)

	var r0 *zest.RpmRpmPublicationResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*zest.RpmRpmPublicationResponse, error)); ok {
		return rf(versionHref)
	}
	if rf, ok := ret.Get(0).(func(string) *zest.RpmRpmPublicationResponse); ok {
		r0 = rf(versionHref)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*zest.RpmRpmPublicationResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(versionHref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRpmRemoteByName provides a mock function with given fields: name
func (_m *MockPulpClient) GetRpmRemoteByName(name string) (*zest.RpmRpmRemoteResponse, error) {
	ret := _m.Called(name)

	var r0 *zest.RpmRpmRemoteResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*zest.RpmRpmRemoteResponse, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *zest.RpmRpmRemoteResponse); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*zest.RpmRpmRemoteResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRpmRemoteList provides a mock function with given fields:
func (_m *MockPulpClient) GetRpmRemoteList() ([]zest.RpmRpmRemoteResponse, error) {
	ret := _m.Called()

	var r0 []zest.RpmRpmRemoteResponse
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]zest.RpmRpmRemoteResponse, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []zest.RpmRpmRemoteResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]zest.RpmRpmRemoteResponse)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRpmRepositoryByName provides a mock function with given fields: name
func (_m *MockPulpClient) GetRpmRepositoryByName(name string) (*zest.RpmRpmRepositoryResponse, error) {
	ret := _m.Called(name)

	var r0 *zest.RpmRpmRepositoryResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*zest.RpmRpmRepositoryResponse, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *zest.RpmRpmRepositoryResponse); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*zest.RpmRpmRepositoryResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRpmRepositoryByRemote provides a mock function with given fields: pulpHref
func (_m *MockPulpClient) GetRpmRepositoryByRemote(pulpHref string) (*zest.RpmRpmRepositoryResponse, error) {
	ret := _m.Called(pulpHref)

	var r0 *zest.RpmRpmRepositoryResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*zest.RpmRpmRepositoryResponse, error)); ok {
		return rf(pulpHref)
	}
	if rf, ok := ret.Get(0).(func(string) *zest.RpmRpmRepositoryResponse); ok {
		r0 = rf(pulpHref)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*zest.RpmRpmRepositoryResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(pulpHref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRpmRepositoryVersion provides a mock function with given fields: href
func (_m *MockPulpClient) GetRpmRepositoryVersion(href string) (*zest.RepositoryVersionResponse, error) {
	ret := _m.Called(href)

	var r0 *zest.RepositoryVersionResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*zest.RepositoryVersionResponse, error)); ok {
		return rf(href)
	}
	if rf, ok := ret.Get(0).(func(string) *zest.RepositoryVersionResponse); ok {
		r0 = rf(href)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*zest.RepositoryVersionResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(href)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTask provides a mock function with given fields: taskHref
func (_m *MockPulpClient) GetTask(taskHref string) (zest.TaskResponse, error) {
	ret := _m.Called(taskHref)

	var r0 zest.TaskResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (zest.TaskResponse, error)); ok {
		return rf(taskHref)
	}
	if rf, ok := ret.Get(0).(func(string) zest.TaskResponse); ok {
		r0 = rf(taskHref)
	} else {
		r0 = ret.Get(0).(zest.TaskResponse)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(taskHref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LookupDomain provides a mock function with given fields: name
func (_m *MockPulpClient) LookupDomain(name string) (string, error) {
	ret := _m.Called(name)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LookupOrCreateDomain provides a mock function with given fields: name
func (_m *MockPulpClient) LookupOrCreateDomain(name string) (string, error) {
	ret := _m.Called(name)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PollTask provides a mock function with given fields: taskHref
func (_m *MockPulpClient) PollTask(taskHref string) (*zest.TaskResponse, error) {
	ret := _m.Called(taskHref)

	var r0 *zest.TaskResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*zest.TaskResponse, error)); ok {
		return rf(taskHref)
	}
	if rf, ok := ret.Get(0).(func(string) *zest.TaskResponse); ok {
		r0 = rf(taskHref)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*zest.TaskResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(taskHref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Status provides a mock function with given fields:
func (_m *MockPulpClient) Status() (*zest.StatusResponse, error) {
	ret := _m.Called()

	var r0 *zest.StatusResponse
	var r1 error
	if rf, ok := ret.Get(0).(func() (*zest.StatusResponse, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *zest.StatusResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*zest.StatusResponse)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SyncRpmRepository provides a mock function with given fields: rpmRpmRepositoryHref, remoteHref
func (_m *MockPulpClient) SyncRpmRepository(rpmRpmRepositoryHref string, remoteHref *string) (string, error) {
	ret := _m.Called(rpmRpmRepositoryHref, remoteHref)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *string) (string, error)); ok {
		return rf(rpmRpmRepositoryHref, remoteHref)
	}
	if rf, ok := ret.Get(0).(func(string, *string) string); ok {
		r0 = rf(rpmRpmRepositoryHref, remoteHref)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, *string) error); ok {
		r1 = rf(rpmRpmRepositoryHref, remoteHref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDomainIfNeeded provides a mock function with given fields: name
func (_m *MockPulpClient) UpdateDomainIfNeeded(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateRpmRemoteUrl provides a mock function with given fields: pulpHref, url
func (_m *MockPulpClient) UpdateRpmRemoteUrl(pulpHref string, url string) (string, error) {
	ret := _m.Called(pulpHref, url)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(pulpHref, url)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(pulpHref, url)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(pulpHref, url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockPulpClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockPulpClient creates a new instance of MockPulpClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockPulpClient(t mockConstructorTestingTNewMockPulpClient) *MockPulpClient {
	mock := &MockPulpClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
