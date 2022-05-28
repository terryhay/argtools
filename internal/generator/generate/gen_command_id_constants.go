package generate

import (
	"argtools/internal/generator/configYaml"
	"argtools/internal/generator/idTemplateDataCreator"
	"fmt"
	"sort"
)

const (
	commandIDFirstConstTemplate = `	// %s - %s
	%s argParserConfig.CommandID = iota + 1
`

	commandIDConstTemplate = `	// %s - %s
	%s
`
)

type CommandIDListComponent string

type byID []*idTemplateDataCreator.IDTemplateData

func (i byID) Len() int { return len(i) }

func (i byID) Less(left, right int) bool { return i[left].GetStringID() < i[right].GetStringID() }

func (i byID) Swap(left, right int) { i[left], i[right] = i[right], i[left] }

func GenCommandIDConstants(
	commandsTemplateData map[configYaml.Command]*idTemplateDataCreator.IDTemplateData,
	nullCommandIDTemplateData *idTemplateDataCreator.IDTemplateData,
) CommandIDListComponent {

	dataCount := 0
	if nullCommandIDTemplateData != nil {
		dataCount++
	}
	dataCount += len(commandsTemplateData)

	if dataCount == 0 {
		return ""
	}

	checkDuplicates := make(map[string]bool, dataCount)
	sortedCommandsTemplateData := make([]*idTemplateDataCreator.IDTemplateData, 0, dataCount)
	var contains bool

	if nullCommandIDTemplateData != nil {
		checkDuplicates[nullCommandIDTemplateData.GetID()] = true
		sortedCommandsTemplateData = append(sortedCommandsTemplateData, nullCommandIDTemplateData)
	}

	for _, data := range commandsTemplateData {
		if _, contains = checkDuplicates[data.GetID()]; contains {
			continue
		}
		checkDuplicates[data.GetID()] = true

		sortedCommandsTemplateData = append(sortedCommandsTemplateData, data)
	}
	sort.Sort(byID(sortedCommandsTemplateData))

	templateData := sortedCommandsTemplateData[0]
	res := fmt.Sprintf(commandIDFirstConstTemplate, templateData.GetID(), templateData.GetComment(), templateData.GetID())

	for i := 1; i < len(sortedCommandsTemplateData); i++ {
		templateData = sortedCommandsTemplateData[i]
		res += fmt.Sprintf(commandIDConstTemplate, templateData.GetID(), templateData.GetComment(), templateData.GetID())
	}

	return CommandIDListComponent(res)
}
