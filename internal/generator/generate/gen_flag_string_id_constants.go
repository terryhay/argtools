package generate

import (
	"argtools/internal/generator/configYaml"
	"argtools/internal/generator/idTemplateDataCreator"
	"fmt"
	"sort"
)

const (
	flagStringIDFirstConstTemplate = `	// %s - %s
	%s argParserConfig.Flag = "%s"
`
	flagStringIDConstTemplate = `	// %s - %s
	%s = "%s"
`
)

type FlagStringIDListComponent string

func GenFlagStringIDConstants(flagsTemplateData map[configYaml.Flag]*idTemplateDataCreator.IDTemplateData) FlagStringIDListComponent {
	dataCount := len(flagsTemplateData)
	if dataCount == 0 {
		return ""
	}

	sortedFlagsTemplateData := make([]*idTemplateDataCreator.IDTemplateData, 0, dataCount)
	for _, data := range flagsTemplateData {
		sortedFlagsTemplateData = append(sortedFlagsTemplateData, data)
	}
	sort.Sort(byStringID(sortedFlagsTemplateData))

	templateData := sortedFlagsTemplateData[0]
	res := fmt.Sprintf(flagStringIDFirstConstTemplate, templateData.GetStringID(), templateData.GetComment(), templateData.GetStringID(), templateData.GetCallName())

	for i := 1; i < len(sortedFlagsTemplateData); i++ {
		templateData = sortedFlagsTemplateData[i]
		res += fmt.Sprintf(flagStringIDConstTemplate, templateData.GetStringID(), templateData.GetComment(), templateData.GetStringID(), templateData.GetCallName())
	}

	return FlagStringIDListComponent(res)
}
