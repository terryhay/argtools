package generate

import (
	"fmt"
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"github.com/terryhay/argtools/internal/generator/idTemplateDataCreator"
	"sort"
)

const (
	commandStringIDFirstConstTemplate = `	// %s - %s
	%s argParserConfig.Command = "%s"
`

	commandStringIDConstTemplate = `	// %s - %s
	%s = "%s"
`
)

type CommandStringIDListComponent string

type byStringID []*idTemplateDataCreator.IDTemplateData

func (i byStringID) Len() int {
	return len(i)
}

func (i byStringID) Less(left, right int) bool {
	return i[left].GetStringID() < i[right].GetStringID()
}

func (i byStringID) Swap(left, right int) {
	i[left], i[right] = i[right], i[left]
}

func GenCommandStringIDConstants(
	commandsTemplateData map[configYaml.Command]*idTemplateDataCreator.IDTemplateData,
) CommandStringIDListComponent {

	sortedCommandsTemplateData := make([]*idTemplateDataCreator.IDTemplateData, 0, len(commandsTemplateData))
	for _, data := range commandsTemplateData {
		sortedCommandsTemplateData = append(sortedCommandsTemplateData, data)
	}
	sort.Sort(byStringID(sortedCommandsTemplateData))

	templateData := sortedCommandsTemplateData[0]
	res := fmt.Sprintf(commandStringIDFirstConstTemplate, templateData.GetStringID(), templateData.GetComment(), templateData.GetStringID(), templateData.GetCallName())

	for i := 1; i < len(sortedCommandsTemplateData); i++ {
		templateData = sortedCommandsTemplateData[i]
		res += fmt.Sprintf(commandStringIDConstTemplate, templateData.GetStringID(), templateData.GetComment(), templateData.GetStringID(), templateData.GetCallName())
	}

	return CommandStringIDListComponent(res)
}
