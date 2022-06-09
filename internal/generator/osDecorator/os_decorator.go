package osDecorator

import (
	"fmt"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"os"
)

// OSDecorator - os methods decorator interface
type OSDecorator interface {
	// GetArgs Args - returns command line arguments without application name
	GetArgs() []string

	// Create - creates or truncates the named file
	Create(path string) (FileDecorator, error)

	// Exit - causes the current program to exit with the given error
	Exit(err *argtoolsError.Error)

	// IsNotExist - checks if error is "not exist"
	IsNotExist(err error) bool

	// MkdirAll - creates a directory named path
	MkdirAll(path string, perm os.FileMode) error

	// Stat - returns a FileInfo describing the named file
	Stat(name string) (os.FileInfo, error)
}

// NewOSDecorator - os decorator instance constructor
func NewOSDecorator() OSDecorator {
	return &osDecoratorImpl{}
}

type osDecoratorImpl struct{}

// GetArgs Args - returns command line arguments without application name
func (osDecoratorImpl) GetArgs() []string {
	return os.Args[1:]
}

// Create - creates or truncates the named file
func (osDecoratorImpl) Create(path string) (FileDecorator, error) {
	file, err := os.Create(path)
	return NewFileDecorator(file), err
}

// Exit - causes the current program to exit with the given error
func (osDecoratorImpl) Exit(err *argtoolsError.Error) {
	exitCode := 0
	if err != nil {
		fmt.Println("argParser generator: " + err.Error())
		exitCode = int(err.Code())
	}
	os.Exit(exitCode)
}

// IsNotExist - checks if error is "not exist"
func (osDecoratorImpl) IsNotExist(err error) bool {
	return os.IsNotExist(err)
}

// MkdirAll - creates a directory named path
func (osDecoratorImpl) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

// Stat - returns a FileInfo describing the named file
func (osDecoratorImpl) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}
