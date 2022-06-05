package osDecoratorMock

import (
	"github.com/terryhay/argtools/internal/generator/osDecorator"
)

// NewMockFileDecorator - mocked file decorator instance constructor
func NewMockFileDecorator(
	mockClose func() error,
	mockWriteString func(s string) error,
) osDecorator.FileDecorator {

	return &fileDecoratorMockImpl{
		mockClose:       mockClose,
		mockWriteString: mockWriteString,
	}
}

type fileDecoratorMockImpl struct {
	mockClose       func() error
	mockWriteString func(s string) error
}

// Close - closes the file
func (i *fileDecoratorMockImpl) Close() error {
	return i.mockClose()
}

// WriteString - writes a string into the file
func (i *fileDecoratorMockImpl) WriteString(s string) error {
	return i.mockWriteString(s)
}
