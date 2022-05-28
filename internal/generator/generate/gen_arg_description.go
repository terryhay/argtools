package generate

import (
	"fmt"
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"strings"
)

const (
	argumentsDescriptionPrefix = `				ArgDescription: &argParserConfig.ArgumentsDescription{
					AmountType:              %s,
					SynopsisHelpDescription: "%s",
`
	argumentsDescriptionDefaultValuesPrefix = `					DefaultValues: []string{
`
	argumentsDescriptionAllowedValuesPrefix = `					AllowedValues: map[string]bool{
`
	argumentsDescriptionVariantValue = `						"%s",
`
	argumentsDescriptionMapVariantValue = `						"%s": true,
`
	argumentsDescriptionVariantValuesPostfix = `					},
`
	argumentsDescriptionPostfix = `				},
`
)

type ArgDescriptionElement string

func GenArgDescriptionElement(argumentsDescription *configYaml.ArgumentsDescription) ArgDescriptionElement {
	builder := new(strings.Builder)

	builder.WriteString(fmt.Sprintf(argumentsDescriptionPrefix,
		getArgAmountTypeElement(argumentsDescription.GetAmountType()),
		argumentsDescription.GetSynopsisHelpDescription()))

	if defaultValues := argumentsDescription.GetDefaultValues(); len(defaultValues) > 0 {
		builder.WriteString(argumentsDescriptionDefaultValuesPrefix)
		for _, value := range defaultValues {
			builder.WriteString(fmt.Sprintf(argumentsDescriptionVariantValue, value))
		}
		builder.WriteString(argumentsDescriptionVariantValuesPostfix)
	}

	if allowedValues := argumentsDescription.GetAllowedValues(); len(allowedValues) > 0 {
		builder.WriteString(argumentsDescriptionAllowedValuesPrefix)
		for _, value := range allowedValues {
			builder.WriteString(fmt.Sprintf(argumentsDescriptionMapVariantValue, value))
		}
		builder.WriteString(argumentsDescriptionVariantValuesPostfix)
	}

	builder.WriteString(argumentsDescriptionPostfix)

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
