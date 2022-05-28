package argtoolsError

// Code is spec Code of error
type Code uint

const (
	// CodeUndefinedError - undefined internal error Code
	CodeUndefinedError Code = iota

	// CodeConfigContainsDuplicateCommands - some command is duplicating
	CodeConfigContainsDuplicateCommands

	// CodeConfigContainsDuplicateFlags - some flag is duplicating
	CodeConfigContainsDuplicateFlags

	// CodeConfigDoesNotContainCommandDescriptions - yaml config file doesn't contain command descriptions
	CodeConfigDoesNotContainCommandDescriptions

	// CodeConfigDoesNotContainFlagDescriptions - yaml config file doesn't contain flag descriptions
	CodeConfigDoesNotContainFlagDescriptions

	// CodeCantExtractFlagArgValue - error of extracting flag argument value
	CodeCantExtractFlagArgValue

	// CodeConfigFlagIsNotUsedInCommands - some flag is described, but not used in commands descriptions
	CodeConfigFlagIsNotUsedInCommands

	// CodeConfigUndefinedFlag - some flag is undefined in flag description list of yaml config file
	CodeConfigUndefinedFlag

	// CodeCantFindFlagNameInGroupSpec - unexpected flag name for determine using flag group
	CodeCantFindFlagNameInGroupSpec

	// CodeDuplicatedGroupFlag - some group flags are duplicated
	CodeDuplicatedGroupFlag

	// CodeFlagArgCountError - flag arg amount less or more than settings expected
	CodeFlagArgCountError

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

	// CodeGeneratorCreateDirError - create a dir error
	CodeGeneratorCreateDirError

	// CodeGeneratorCreateFileError - create a file error
	CodeGeneratorCreateFileError

	// CodeGeneratorCloseFileError - close a file error
	CodeGeneratorCloseFileError

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

	// CodeUnexpectedFlagName - unexpected flag was found
	CodeUnexpectedFlagName

	// CodeUnexpectedFlagArgTypeDescription - no logic for flag arg type description
	CodeUnexpectedFlagArgTypeDescription
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

// Code returns code of error, you must check if error == nil before
func (i *Error) Code() Code {
	return i.code
}

// Error decorates standard error interface
func (i *Error) Error() string {
	return i.err.Error()
}
