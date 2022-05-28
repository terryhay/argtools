package helpPrinter

import (
	"fmt"
)

// CreateNameChapter - creates name help chapter
func CreateNameChapter(appName, nameHelpInfo string) string {
	return fmt.Sprintf("\u001B[1mNAME\u001B[0m\n\t\u001B[1m%s\u001B[0m – %s\n\n", appName, nameHelpInfo)
}
