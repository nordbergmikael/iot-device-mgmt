// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package database

import (
	"io"
	"sync"
	"time"
)

// Ensure, that DatastoreMock does implement Datastore.
// If this is not the case, regenerate this file with moq.
var _ Datastore = &DatastoreMock{}

// DatastoreMock is a mock implementation of Datastore.
//
// 	func TestSomethingThatUsesDatastore(t *testing.T) {
//
// 		// make and configure a mocked Datastore
// 		mockedDatastore := &DatastoreMock{
// 			CreateDeviceFunc: func(devEUI string, deviceId string, name string, description string, environment string, sensorType string, tenant string, latitude float64, longitude float64, types []string, active bool) (Device, error) {
// 				panic("mock out the CreateDevice method")
// 			},
// 			GetAllFunc: func() ([]Device, error) {
// 				panic("mock out the GetAll method")
// 			},
// 			GetDeviceFromDevEUIFunc: func(eui string) (Device, error) {
// 				panic("mock out the GetDeviceFromDevEUI method")
// 			},
// 			GetDeviceFromIDFunc: func(deviceID string) (Device, error) {
// 				panic("mock out the GetDeviceFromID method")
// 			},
// 			ListEnvironmentsFunc: func() ([]Environment, error) {
// 				panic("mock out the ListEnvironments method")
// 			},
// 			SeedFunc: func(r io.Reader) error {
// 				panic("mock out the Seed method")
// 			},
// 			UpdateDeviceFunc: func(deviceID string, fields map[string]interface{}) (Device, error) {
// 				panic("mock out the UpdateDevice method")
// 			},
// 			UpdateLastObservedOnDeviceFunc: func(deviceID string, timestamp time.Time) error {
// 				panic("mock out the UpdateLastObservedOnDevice method")
// 			},
// 		}
//
// 		// use mockedDatastore in code that requires Datastore
// 		// and then make assertions.
//
// 	}
type DatastoreMock struct {
	// CreateDeviceFunc mocks the CreateDevice method.
	CreateDeviceFunc func(devEUI string, deviceId string, name string, description string, environment string, sensorType string, tenant string, latitude float64, longitude float64, types []string, active bool) (Device, error)

	// GetAllFunc mocks the GetAll method.
	GetAllFunc func() ([]Device, error)

	// GetDeviceFromDevEUIFunc mocks the GetDeviceFromDevEUI method.
	GetDeviceFromDevEUIFunc func(eui string) (Device, error)

	// GetDeviceFromIDFunc mocks the GetDeviceFromID method.
	GetDeviceFromIDFunc func(deviceID string) (Device, error)

	// ListEnvironmentsFunc mocks the ListEnvironments method.
	ListEnvironmentsFunc func() ([]Environment, error)

	// SeedFunc mocks the Seed method.
	SeedFunc func(r io.Reader) error

	// UpdateDeviceFunc mocks the UpdateDevice method.
	UpdateDeviceFunc func(deviceID string, fields map[string]interface{}) (Device, error)

	// UpdateLastObservedOnDeviceFunc mocks the UpdateLastObservedOnDevice method.
	UpdateLastObservedOnDeviceFunc func(deviceID string, timestamp time.Time) error

	// calls tracks calls to the methods.
	calls struct {
		// CreateDevice holds details about calls to the CreateDevice method.
		CreateDevice []struct {
			// DevEUI is the devEUI argument value.
			DevEUI string
			// DeviceId is the deviceId argument value.
			DeviceId string
			// Name is the name argument value.
			Name string
			// Description is the description argument value.
			Description string
			// Environment is the environment argument value.
			Environment string
			// SensorType is the sensorType argument value.
			SensorType string
			// Tenant is the tenant argument value.
			Tenant string
			// Latitude is the latitude argument value.
			Latitude float64
			// Longitude is the longitude argument value.
			Longitude float64
			// Types is the types argument value.
			Types []string
			// Active is the active argument value.
			Active bool
		}
		// GetAll holds details about calls to the GetAll method.
		GetAll []struct {
		}
		// GetDeviceFromDevEUI holds details about calls to the GetDeviceFromDevEUI method.
		GetDeviceFromDevEUI []struct {
			// Eui is the eui argument value.
			Eui string
		}
		// GetDeviceFromID holds details about calls to the GetDeviceFromID method.
		GetDeviceFromID []struct {
			// DeviceID is the deviceID argument value.
			DeviceID string
		}
		// ListEnvironments holds details about calls to the ListEnvironments method.
		ListEnvironments []struct {
		}
		// Seed holds details about calls to the Seed method.
		Seed []struct {
			// R is the r argument value.
			R io.Reader
		}
		// UpdateDevice holds details about calls to the UpdateDevice method.
		UpdateDevice []struct {
			// DeviceID is the deviceID argument value.
			DeviceID string
			// Fields is the fields argument value.
			Fields map[string]interface{}
		}
		// UpdateLastObservedOnDevice holds details about calls to the UpdateLastObservedOnDevice method.
		UpdateLastObservedOnDevice []struct {
			// DeviceID is the deviceID argument value.
			DeviceID string
			// Timestamp is the timestamp argument value.
			Timestamp time.Time
		}
	}
	lockCreateDevice               sync.RWMutex
	lockGetAll                     sync.RWMutex
	lockGetDeviceFromDevEUI        sync.RWMutex
	lockGetDeviceFromID            sync.RWMutex
	lockListEnvironments           sync.RWMutex
	lockSeed                       sync.RWMutex
	lockUpdateDevice               sync.RWMutex
	lockUpdateLastObservedOnDevice sync.RWMutex
}

// CreateDevice calls CreateDeviceFunc.
func (mock *DatastoreMock) CreateDevice(devEUI string, deviceId string, name string, description string, environment string, sensorType string, tenant string, latitude float64, longitude float64, types []string, active bool) (Device, error) {
	if mock.CreateDeviceFunc == nil {
		panic("DatastoreMock.CreateDeviceFunc: method is nil but Datastore.CreateDevice was just called")
	}
	callInfo := struct {
		DevEUI      string
		DeviceId    string
		Name        string
		Description string
		Environment string
		SensorType  string
		Tenant      string
		Latitude    float64
		Longitude   float64
		Types       []string
		Active      bool
	}{
		DevEUI:      devEUI,
		DeviceId:    deviceId,
		Name:        name,
		Description: description,
		Environment: environment,
		SensorType:  sensorType,
		Tenant:      tenant,
		Latitude:    latitude,
		Longitude:   longitude,
		Types:       types,
		Active:      active,
	}
	mock.lockCreateDevice.Lock()
	mock.calls.CreateDevice = append(mock.calls.CreateDevice, callInfo)
	mock.lockCreateDevice.Unlock()
	return mock.CreateDeviceFunc(devEUI, deviceId, name, description, environment, sensorType, tenant, latitude, longitude, types, active)
}

// CreateDeviceCalls gets all the calls that were made to CreateDevice.
// Check the length with:
//     len(mockedDatastore.CreateDeviceCalls())
func (mock *DatastoreMock) CreateDeviceCalls() []struct {
	DevEUI      string
	DeviceId    string
	Name        string
	Description string
	Environment string
	SensorType  string
	Tenant      string
	Latitude    float64
	Longitude   float64
	Types       []string
	Active      bool
} {
	var calls []struct {
		DevEUI      string
		DeviceId    string
		Name        string
		Description string
		Environment string
		SensorType  string
		Tenant      string
		Latitude    float64
		Longitude   float64
		Types       []string
		Active      bool
	}
	mock.lockCreateDevice.RLock()
	calls = mock.calls.CreateDevice
	mock.lockCreateDevice.RUnlock()
	return calls
}

// GetAll calls GetAllFunc.
func (mock *DatastoreMock) GetAll() ([]Device, error) {
	if mock.GetAllFunc == nil {
		panic("DatastoreMock.GetAllFunc: method is nil but Datastore.GetAll was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetAll.Lock()
	mock.calls.GetAll = append(mock.calls.GetAll, callInfo)
	mock.lockGetAll.Unlock()
	return mock.GetAllFunc()
}

// GetAllCalls gets all the calls that were made to GetAll.
// Check the length with:
//     len(mockedDatastore.GetAllCalls())
func (mock *DatastoreMock) GetAllCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetAll.RLock()
	calls = mock.calls.GetAll
	mock.lockGetAll.RUnlock()
	return calls
}

// GetDeviceFromDevEUI calls GetDeviceFromDevEUIFunc.
func (mock *DatastoreMock) GetDeviceFromDevEUI(eui string) (Device, error) {
	if mock.GetDeviceFromDevEUIFunc == nil {
		panic("DatastoreMock.GetDeviceFromDevEUIFunc: method is nil but Datastore.GetDeviceFromDevEUI was just called")
	}
	callInfo := struct {
		Eui string
	}{
		Eui: eui,
	}
	mock.lockGetDeviceFromDevEUI.Lock()
	mock.calls.GetDeviceFromDevEUI = append(mock.calls.GetDeviceFromDevEUI, callInfo)
	mock.lockGetDeviceFromDevEUI.Unlock()
	return mock.GetDeviceFromDevEUIFunc(eui)
}

// GetDeviceFromDevEUICalls gets all the calls that were made to GetDeviceFromDevEUI.
// Check the length with:
//     len(mockedDatastore.GetDeviceFromDevEUICalls())
func (mock *DatastoreMock) GetDeviceFromDevEUICalls() []struct {
	Eui string
} {
	var calls []struct {
		Eui string
	}
	mock.lockGetDeviceFromDevEUI.RLock()
	calls = mock.calls.GetDeviceFromDevEUI
	mock.lockGetDeviceFromDevEUI.RUnlock()
	return calls
}

// GetDeviceFromID calls GetDeviceFromIDFunc.
func (mock *DatastoreMock) GetDeviceFromID(deviceID string) (Device, error) {
	if mock.GetDeviceFromIDFunc == nil {
		panic("DatastoreMock.GetDeviceFromIDFunc: method is nil but Datastore.GetDeviceFromID was just called")
	}
	callInfo := struct {
		DeviceID string
	}{
		DeviceID: deviceID,
	}
	mock.lockGetDeviceFromID.Lock()
	mock.calls.GetDeviceFromID = append(mock.calls.GetDeviceFromID, callInfo)
	mock.lockGetDeviceFromID.Unlock()
	return mock.GetDeviceFromIDFunc(deviceID)
}

// GetDeviceFromIDCalls gets all the calls that were made to GetDeviceFromID.
// Check the length with:
//     len(mockedDatastore.GetDeviceFromIDCalls())
func (mock *DatastoreMock) GetDeviceFromIDCalls() []struct {
	DeviceID string
} {
	var calls []struct {
		DeviceID string
	}
	mock.lockGetDeviceFromID.RLock()
	calls = mock.calls.GetDeviceFromID
	mock.lockGetDeviceFromID.RUnlock()
	return calls
}

// ListEnvironments calls ListEnvironmentsFunc.
func (mock *DatastoreMock) ListEnvironments() ([]Environment, error) {
	if mock.ListEnvironmentsFunc == nil {
		panic("DatastoreMock.ListEnvironmentsFunc: method is nil but Datastore.ListEnvironments was just called")
	}
	callInfo := struct {
	}{}
	mock.lockListEnvironments.Lock()
	mock.calls.ListEnvironments = append(mock.calls.ListEnvironments, callInfo)
	mock.lockListEnvironments.Unlock()
	return mock.ListEnvironmentsFunc()
}

// ListEnvironmentsCalls gets all the calls that were made to ListEnvironments.
// Check the length with:
//     len(mockedDatastore.ListEnvironmentsCalls())
func (mock *DatastoreMock) ListEnvironmentsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockListEnvironments.RLock()
	calls = mock.calls.ListEnvironments
	mock.lockListEnvironments.RUnlock()
	return calls
}

// Seed calls SeedFunc.
func (mock *DatastoreMock) Seed(r io.Reader) error {
	if mock.SeedFunc == nil {
		panic("DatastoreMock.SeedFunc: method is nil but Datastore.Seed was just called")
	}
	callInfo := struct {
		R io.Reader
	}{
		R: r,
	}
	mock.lockSeed.Lock()
	mock.calls.Seed = append(mock.calls.Seed, callInfo)
	mock.lockSeed.Unlock()
	return mock.SeedFunc(r)
}

// SeedCalls gets all the calls that were made to Seed.
// Check the length with:
//     len(mockedDatastore.SeedCalls())
func (mock *DatastoreMock) SeedCalls() []struct {
	R io.Reader
} {
	var calls []struct {
		R io.Reader
	}
	mock.lockSeed.RLock()
	calls = mock.calls.Seed
	mock.lockSeed.RUnlock()
	return calls
}

// UpdateDevice calls UpdateDeviceFunc.
func (mock *DatastoreMock) UpdateDevice(deviceID string, fields map[string]interface{}) (Device, error) {
	if mock.UpdateDeviceFunc == nil {
		panic("DatastoreMock.UpdateDeviceFunc: method is nil but Datastore.UpdateDevice was just called")
	}
	callInfo := struct {
		DeviceID string
		Fields   map[string]interface{}
	}{
		DeviceID: deviceID,
		Fields:   fields,
	}
	mock.lockUpdateDevice.Lock()
	mock.calls.UpdateDevice = append(mock.calls.UpdateDevice, callInfo)
	mock.lockUpdateDevice.Unlock()
	return mock.UpdateDeviceFunc(deviceID, fields)
}

// UpdateDeviceCalls gets all the calls that were made to UpdateDevice.
// Check the length with:
//     len(mockedDatastore.UpdateDeviceCalls())
func (mock *DatastoreMock) UpdateDeviceCalls() []struct {
	DeviceID string
	Fields   map[string]interface{}
} {
	var calls []struct {
		DeviceID string
		Fields   map[string]interface{}
	}
	mock.lockUpdateDevice.RLock()
	calls = mock.calls.UpdateDevice
	mock.lockUpdateDevice.RUnlock()
	return calls
}

// UpdateLastObservedOnDevice calls UpdateLastObservedOnDeviceFunc.
func (mock *DatastoreMock) UpdateLastObservedOnDevice(deviceID string, timestamp time.Time) error {
	if mock.UpdateLastObservedOnDeviceFunc == nil {
		panic("DatastoreMock.UpdateLastObservedOnDeviceFunc: method is nil but Datastore.UpdateLastObservedOnDevice was just called")
	}
	callInfo := struct {
		DeviceID  string
		Timestamp time.Time
	}{
		DeviceID:  deviceID,
		Timestamp: timestamp,
	}
	mock.lockUpdateLastObservedOnDevice.Lock()
	mock.calls.UpdateLastObservedOnDevice = append(mock.calls.UpdateLastObservedOnDevice, callInfo)
	mock.lockUpdateLastObservedOnDevice.Unlock()
	return mock.UpdateLastObservedOnDeviceFunc(deviceID, timestamp)
}

// UpdateLastObservedOnDeviceCalls gets all the calls that were made to UpdateLastObservedOnDevice.
// Check the length with:
//     len(mockedDatastore.UpdateLastObservedOnDeviceCalls())
func (mock *DatastoreMock) UpdateLastObservedOnDeviceCalls() []struct {
	DeviceID  string
	Timestamp time.Time
} {
	var calls []struct {
		DeviceID  string
		Timestamp time.Time
	}
	mock.lockUpdateLastObservedOnDevice.RLock()
	calls = mock.calls.UpdateLastObservedOnDevice
	mock.lockUpdateLastObservedOnDevice.RUnlock()
	return calls
}
