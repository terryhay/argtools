package generate

import (
	"fmt"
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"github.com/terryhay/argtools/internal/generator/idTemplateDataCreator"
	"sort"
	"strings"
)

const (
	flagStringIDConstPrefixTemplate = `const (
`

	flagStringIDFirstConstTemplate = `	// %s - %s
	%s argParserConfig.Flag = "%s"
`
	flagStringIDConstTemplate = `	// %s - %s
	%s = "%s"
`
	flagStringIDConstPostfixTemplate = `)

`
)

type FlagStringIDListComponent string

func GenFlagStringIDConstants(flagsTemplateData map[configYaml.Flag]*idTemplateDataCreator.IDTemplateData) FlagStringIDListComponent {
	dataCount := len(flagsTemplateData)
	if dataCount == 0 {
		return ""
	}

	builder := strings.Builder{}
	builder.WriteString(flagStringIDConstPrefixTemplate)

	sortedFlagsTemplateData := make([]*idTemplateDataCreator.IDTemplateData, 0, dataCount)
	for _, data := range flagsTemplateData {
		sortedFlagsTemplateData = append(sortedFlagsTemplateData, data)
	}
	sort.Sort(byStringID(sortedFlagsTemplateData))

	templateData := sortedFlagsTemplateData[0]
	builder.WriteString(fmt.Sprintf(flagStringIDFirstConstTemplate,
		templateData.GetStringID(),
		templateData.GetComment(),
		templateData.GetStringID(),
		templateData.GetCallName()))

	for i := 1; i < len(sortedFlagsTemplateData); i++ {
		templateData = sortedFlagsTemplateData[i]
		builder.WriteString(fmt.Sprintf(flagStringIDConstTemplate, templateData.GetStringID(), templateData.GetComment(), templateData.GetStringID(), templateData.GetCallName()))
	}

	builder.WriteString(flagStringIDConstPostfixTemplate)
	return FlagStringIDListComponent(builder.String())
}
