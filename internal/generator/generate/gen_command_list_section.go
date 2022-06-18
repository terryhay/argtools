package generate

import (
	"fmt"
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"github.com/terryhay/argtools/internal/generator/idTemplateDataCreator"
	"sort"
	"strings"
)

const (
	commandStringIDFirstConstPattern = `
	// %s - %s
	%s argParserConfig.Command = "%s"`
	commandStringIDConstTemplate = `
	// %s - %s
	%s = "%s"`
)

// CommandListSection - string with command constant definitions list
type CommandListSection string

// byNameID - type for sorting IDTemplateData pointers by name id
type byNameID []*idTemplateDataCreator.IDTemplateData

// Len - implementation of Len sort interface method
func (i byNameID) Len() int {
	return len(i)
}

// Less - implementation of Less sort interface method
func (i byNameID) Less(left, right int) bool {
	return i[left].GetNameID() < i[right].GetNameID()
}

// Swap - implementation of Swap sort interface method
func (i byNameID) Swap(left, right int) {
	i[left], i[right] = i[right], i[left]
}

// GenCommandListSection creates a paste section with commands
func GenCommandListSection(
	commandsTemplateData map[configYaml.Command]*idTemplateDataCreator.IDTemplateData,
) CommandListSection {

	sortedCommandsTemplateData := make([]*idTemplateDataCreator.IDTemplateData, 0, len(commandsTemplateData))
	for _, data := range commandsTemplateData {
		sortedCommandsTemplateData = append(sortedCommandsTemplateData, data)
	}
	sort.Sort(byNameID(sortedCommandsTemplateData))

	templateData := sortedCommandsTemplateData[0]

	builder := strings.Builder{}
	builder.WriteString("const (")

	builder.WriteString(fmt.Sprintf(commandStringIDFirstConstPattern,
		templateData.GetNameID(), templateData.GetComment(), templateData.GetNameID(), templateData.GetCallName()))

	for i := 1; i < len(sortedCommandsTemplateData); i++ {
		templateData = sortedCommandsTemplateData[i]
		builder.WriteString(fmt.Sprintf(commandStringIDConstTemplate,
			templateData.GetNameID(), templateData.GetComment(), templateData.GetNameID(), templateData.GetCallName()))
	}

	builder.WriteString("\n)")

	return CommandListSection(builder.String())
}
