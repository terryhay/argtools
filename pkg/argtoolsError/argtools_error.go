package argtoolsError

// Code is spec Code of error
type Code uint

const (
	// CodeNone - null value, no error
	CodeNone Code = iota

	// CodeUndefinedError - undefined internal error Code
	CodeUndefinedError

	// CodeConfigContainsDuplicateCommands - some command is duplicating
	CodeConfigContainsDuplicateCommands

	// CodeConfigContainsDuplicateFlags - some flag is duplicating
	CodeConfigContainsDuplicateFlags

	// CodeConfigFlagIsNotUsedInCommands - some flag is described, but not used in commands descriptions
	CodeConfigFlagIsNotUsedInCommands

	// CodeConfigUndefinedFlag - some flag is undefined in flag description list of yaml config file
	CodeConfigUndefinedFlag

	// CodeCantFindFlagNameInGroupSpec - unexpected flag name for determine using flag group
	CodeCantFindFlagNameInGroupSpec

	// CodeFlagHasNilSpec - flag has nil pointer to spec
	CodeFlagHasNilSpec

	// CodeFlagMustHaveDashPrefix - some not group flag doesn't have a dash "-" prefix
	CodeFlagMustHaveDashPrefix

	// CodeFlagMustBeInGroup - some flag is not in commonFlagSpecMap
	CodeFlagMustBeInGroup

	// CodeFlagMustHaveSpec - some flag from groupSpecSlice is not found in groupSpecSlice
	CodeFlagMustHaveSpec

	// CodeGeneratorInvalidGeneratePath - path is not exist
	CodeGeneratorInvalidGeneratePath

	// CodeGeneratorFileCloseError - file close error
	CodeGeneratorFileCloseError

	// CodeGeneratorCreateDirError - create a dir error
	CodeGeneratorCreateDirError

	// CodeGeneratorCreateFileError - create a file error
	CodeGeneratorCreateFileError

	// CodeGeneratorWriteFileError - write file error
	CodeGeneratorWriteFileError

	// CodeGetConfigReadFileError - can't read yaml config file
	CodeGetConfigReadFileError

	// CodeGetConfigUnmarshalError - some unmarshal yaml config file error
	CodeGetConfigUnmarshalError

	// CodeParserIsNotInitialized - looks like Init method was not called or was called with nil CmdArgSpec pointer
	CodeParserIsNotInitialized

	// CodeRequiredFlagIsNotSet - some required flag of group is not set
	CodeRequiredFlagIsNotSet
)

// Error is detail of parser work error
type Error struct {
	code Code
	err  error
}

// NewError creates Error object and returns pointer
func NewError(code Code, err error) *Error {
	return &Error{
		code: code,
		err:  err,
	}
}

// ReInit - updates error fields
func (i *Error) ReInit(code Code, err error) {
	i.code = code
	i.err = err
}

// Code returns code of error, you must check if error == nil before
func (i *Error) Code() Code {
	return i.code
}

// Error decorates standard error interface
func (i *Error) Error() string {
	return i.err.Error()
}
