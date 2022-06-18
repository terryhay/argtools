package generate

import (
	"fmt"
	"github.com/terryhay/argtools/internal/generator/idTemplateDataCreator"
	"strings"
)

const (
	commandIDConstantsPrefixPart = `
const (`
	commandIDConstantsFirstLinePart = `
	// %s - %s
	%s argParserConfig.CommandID = iota + 1`
	commandIDConstantsLinePart = `
	// %s - %s
	%s`
	commandIDConstantsPostfixPart = `
)`
)

// CommandIDListSection - string with command id constant definitions list
type CommandIDListSection string

// GenCommandIDListSection creates a paste section with command ids
func GenCommandIDListSection(
	commandsTemplateData map[string]*idTemplateDataCreator.IDTemplateData,
	nullCommandIDTemplateData *idTemplateDataCreator.IDTemplateData,
) CommandIDListSection {

	sortedCommandsTemplateData := sortCommandsTemplateData(commandsTemplateData, nullCommandIDTemplateData)

	builder := strings.Builder{}
	builder.WriteString(commandIDConstantsPrefixPart)

	templateData := sortedCommandsTemplateData[0]
	builder.WriteString(fmt.Sprintf(commandIDConstantsFirstLinePart,
		templateData.GetID(), templateData.GetComment(), templateData.GetID()))

	for i := 1; i < len(sortedCommandsTemplateData); i++ {
		templateData = sortedCommandsTemplateData[i]
		builder.WriteString(fmt.Sprintf(commandIDConstantsLinePart,
			templateData.GetID(), templateData.GetComment(), templateData.GetID()))
	}
	builder.WriteString(commandIDConstantsPostfixPart)

	return CommandIDListSection(builder.String())
}
