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

type FlagStringIDListSection string

func GenFlagStringIDConstants(flagsTemplateData map[configYaml.Flag]*idTemplateDataCreator.IDTemplateData) FlagStringIDListSection {
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
	sort.Sort(byNameID(sortedFlagsTemplateData))

	templateData := sortedFlagsTemplateData[0]
	builder.WriteString(fmt.Sprintf(flagStringIDFirstConstTemplate,
		templateData.GetNameID(),
		templateData.GetComment(),
		templateData.GetNameID(),
		templateData.GetCallName()))

	for i := 1; i < len(sortedFlagsTemplateData); i++ {
		templateData = sortedFlagsTemplateData[i]
		builder.WriteString(fmt.Sprintf(flagStringIDConstTemplate, templateData.GetNameID(), templateData.GetComment(), templateData.GetNameID(), templateData.GetCallName()))
	}

	builder.WriteString(flagStringIDConstPostfixTemplate)
	return FlagStringIDListSection(builder.String())
}
