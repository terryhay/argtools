package generate

import (
	"fmt"
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"github.com/terryhay/argtools/internal/generator/idTemplateDataCreator"
	"sort"
	"strings"
)

const (
	commandIDFirstConstPattern = `
	// %s - %s
	%s argParserConfig.CommandID = iota + 1`
	commandIDConstPattern = `
	// %s - %s
	%s`
)

// CommandIDListSection - string with command id constant definitions list
type CommandIDListSection string

// byID - type for sorting IDTemplateData pointers by id
type byID []*idTemplateDataCreator.IDTemplateData

// Len - implementation of Len sort interface method
func (i byID) Len() int {
	return len(i)
}

// Less - implementation of Less sort interface method
func (i byID) Less(left, right int) bool {
	return i[left].GetNameID() < i[right].GetNameID()
}

// Swap - implementation of Swap sort interface method
func (i byID) Swap(left, right int) {
	i[left], i[right] = i[right], i[left]
}

// GenCommandIDListSection creates a paste section with command ids
func GenCommandIDListSection(
	commandsTemplateData map[configYaml.Command]*idTemplateDataCreator.IDTemplateData,
	nullCommandIDTemplateData *idTemplateDataCreator.IDTemplateData,
) CommandIDListSection {

	sortedCommandsTemplateData := sortCommandsTemplateData(commandsTemplateData, nullCommandIDTemplateData)

	builder := strings.Builder{}
	builder.WriteString("const(")

	templateData := sortedCommandsTemplateData[0]
	builder.WriteString(fmt.Sprintf(commandIDFirstConstPattern,
		templateData.GetID(), templateData.GetComment(), templateData.GetID()))

	for i := 1; i < len(sortedCommandsTemplateData); i++ {
		templateData = sortedCommandsTemplateData[i]
		builder.WriteString(fmt.Sprintf(commandIDConstPattern,
			templateData.GetID(), templateData.GetComment(), templateData.GetID()))
	}
	builder.WriteString("\n)")

	return CommandIDListSection(builder.String())
}

func sortCommandsTemplateData(
	commandsTemplateData map[configYaml.Command]*idTemplateDataCreator.IDTemplateData,
	nullCommandIDTemplateData *idTemplateDataCreator.IDTemplateData,
) []*idTemplateDataCreator.IDTemplateData {

	dataCount := 0
	if nullCommandIDTemplateData != nil {
		dataCount++
	}
	dataCount += len(commandsTemplateData)

	checkDuplicates := make(map[string]bool, dataCount)
	sortedCommandsTemplateData := make([]*idTemplateDataCreator.IDTemplateData, 0, dataCount)
	var contains bool

	if nullCommandIDTemplateData != nil {
		checkDuplicates[nullCommandIDTemplateData.GetID()] = true
		sortedCommandsTemplateData = append(sortedCommandsTemplateData, nullCommandIDTemplateData)
	}

	for _, idTemplateData := range commandsTemplateData {
		if _, contains = checkDuplicates[idTemplateData.GetID()]; contains {
			continue
		}
		checkDuplicates[idTemplateData.GetID()] = true

		sortedCommandsTemplateData = append(sortedCommandsTemplateData, idTemplateData)
	}
	sort.Sort(byID(sortedCommandsTemplateData))

	return sortedCommandsTemplateData
}
