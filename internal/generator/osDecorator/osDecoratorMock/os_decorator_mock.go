package osDecoratorMock

import (
	"github.com/terryhay/argtools/internal/generator/osDecorator"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"os"
)

// MockOSDecoratorInit - init struct
type MockOSDecoratorInit struct {
	Args           []string
	CreateFunc     func(path string) (osDecorator.FileDecorator, error)
	ExitFunc       func(err *argtoolsError.Error)
	IsNotExistFunc func(err error) bool
	MkdirAllFunc   func(path string, perm os.FileMode) error
	StatFunc       func(path string) (os.FileInfo, error)
}

// NewMockOSDecorator - mocked os decorator instance constructor
func NewMockOSDecorator(init MockOSDecoratorInit) osDecorator.OSDecorator {
	return &osDecoratorMockImpl{
		mockArgs:           init.Args,
		mockCreateFunc:     init.CreateFunc,
		mockExit:           init.ExitFunc,
		mockIsNotExistFunc: init.IsNotExistFunc,
		mockMkdirAll:       init.MkdirAllFunc,
		mockStatFunc:       init.StatFunc,
	}
}

type osDecoratorMockImpl struct {
	mockArgs           []string
	mockCreateFunc     func(path string) (osDecorator.FileDecorator, error)
	mockExit           func(err *argtoolsError.Error)
	mockIsNotExistFunc func(err error) bool
	mockMkdirAll       func(path string, perm os.FileMode) error
	mockStatFunc       func(name string) (os.FileInfo, error)
}

// Args - returns command line arguments without application name
func (i *osDecoratorMockImpl) GetArgs() []string {
	return i.mockArgs
}

// Create - creates or truncates the named file
func (i *osDecoratorMockImpl) Create(path string) (osDecorator.FileDecorator, error) {
	return i.mockCreateFunc(path)
}

// Exit - causes the current program to exit with the given error
func (i *osDecoratorMockImpl) Exit(err *argtoolsError.Error) {
	i.mockExit(err)
}

// IsNotExist - checks if error is "not exist"
func (i *osDecoratorMockImpl) IsNotExist(err error) bool {
	return i.mockIsNotExistFunc(err)
}

// MkdirAll - creates a directory named path
func (i *osDecoratorMockImpl) MkdirAll(path string, perm os.FileMode) error {
	return i.mockMkdirAll(path, perm)
}

// Stat - returns a FileInfo describing the named file
func (i *osDecoratorMockImpl) Stat(name string) (os.FileInfo, error) {
	return i.mockStatFunc(name)
}
