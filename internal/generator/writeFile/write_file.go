package writeFile

import (
	"fmt"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"os"
	"unicode/utf8"
)

const (
	argParserDirName  = "argTools"
	argParserFileName = "arg_tools.go"
)

// Write - write file string by dir path
func Write(fileString, dirPath string) (err *argtoolsError.Error) {
	err = checkDirPath(dirPath)

	argParserDirPath := expandPath(dirPath, argParserDirName)
	if err = checkDirPath(argParserDirPath); err != nil {
		err = createArgParserDir(argParserDirPath)
		if err != nil {
			return err
		}
	}

	argParserFilePath := expandPath(argParserDirPath, argParserFileName)
	err = write(fileString, argParserFilePath)
	if err != nil {
		return err
	}

	return nil
}

func checkDirPath(dirPath string) *argtoolsError.Error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return argtoolsError.NewError(
			argtoolsError.CodeGeneratorInvalidGeneratePath,
			fmt.Errorf("check path exist error: %v\n", err),
		)
	}
	return nil
}

func expandPath(path, name string) string {
	generatePathRunes := []rune(path)
	backRune := generatePathRunes[len(generatePathRunes)-1]

	slashRune, _ := utf8.DecodeRuneInString("/")
	backSlashRune, _ := utf8.DecodeRuneInString("\\")

	if backRune != slashRune || backRune != backSlashRune {
		path += "/"
	}
	return path + name
}

func createArgParserDir(generatePath string) *argtoolsError.Error {
	if osErr := os.MkdirAll(generatePath, 0777); osErr != nil {
		return argtoolsError.NewError(
			argtoolsError.CodeGeneratorCreateDirError,
			fmt.Errorf("create dir error: %v\n", osErr))
	}
	return nil
}

func write(fileString, filePath string) *argtoolsError.Error {
	file, err := os.Create(filePath)
	if err != nil {
		return argtoolsError.NewError(argtoolsError.CodeGeneratorCreateFileError, fmt.Errorf("create file error: %v\n", err))
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			err = fmt.Errorf("can't close the file: %v\n", err)
		}
	}(file)

	_, err = file.WriteString(fileString)
	if err != nil {
		return argtoolsError.NewError(argtoolsError.CodeGeneratorWriteFileError, fmt.Errorf("file write error: %v\n", err))
	}

	return nil
}
