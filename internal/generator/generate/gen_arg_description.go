package generate

import (
	"fmt"
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"strings"
)

const (
	argumentsDescriptionPrefix = `%[1]sArgDescription: &argParserConfig.ArgumentsDescription{
%[1]s	AmountType:              %[2]s,
%[1]s	SynopsisHelpDescription: "%[3]s",
`
	argumentsDescriptionDefaultValuesPrefix = `%s	DefaultValues: []string{
`
	argumentsDescriptionAllowedValuesPrefix = `%s	AllowedValues: map[string]bool{
`
	argumentsDescriptionVariantValue = `%s		"%s",
`
	argumentsDescriptionMapVariantValue = `%s		"%s": true,
`
	argumentsDescriptionVariantValuesPostfix = `%s	},
`
	argumentsDescriptionPostfix = `%s},
`
)

type ArgDescriptionElement string

func GenArgDescriptionElement(
	argumentsDescription *configYaml.ArgumentsDescription,
	indent string,
) ArgDescriptionElement {
	builder := new(strings.Builder)

	builder.WriteString(fmt.Sprintf(argumentsDescriptionPrefix,
		indent,
		getArgAmountTypeElement(argumentsDescription.GetAmountType()),
		argumentsDescription.GetSynopsisHelpDescription()))

	if defaultValues := argumentsDescription.GetDefaultValues(); len(defaultValues) > 0 {
		builder.WriteString(argumentsDescriptionDefaultValuesPrefix)
		for _, value := range defaultValues {
			builder.WriteString(fmt.Sprintf(argumentsDescriptionVariantValue, indent, value))
		}
		builder.WriteString(fmt.Sprintf(argumentsDescriptionVariantValuesPostfix, indent))
	}

	if allowedValues := argumentsDescription.GetAllowedValues(); len(allowedValues) > 0 {
		builder.WriteString(argumentsDescriptionAllowedValuesPrefix)
		for _, value := range allowedValues {
			builder.WriteString(fmt.Sprintf(argumentsDescriptionMapVariantValue, indent, value))
		}
		builder.WriteString(fmt.Sprintf(argumentsDescriptionVariantValuesPostfix, indent))
	}

	builder.WriteString(fmt.Sprintf(argumentsDescriptionPostfix, indent))

	return ArgDescriptionElement(builder.String())
}

func getArgAmountTypeElement(argAmountType argParserConfig.ArgAmountType) string {
	argAmountTypeElement := "argParserConfig.ArgAmountTypeNoArgs"
	switch argAmountType {
	case argParserConfig.ArgAmountTypeSingle:
		argAmountTypeElement = "argParserConfig.ArgAmountTypeSingle"
	case argParserConfig.ArgAmountTypeList:
		argAmountTypeElement = "argParserConfig.ArgAmountTypeList"
	}
	return argAmountTypeElement
}
