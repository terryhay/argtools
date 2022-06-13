package idTemplateDataCreator

import (
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"regexp"
	"unicode"
)

const (
	PrefixCommandID       = "CommandID"
	PrefixCommandStringID = "Command"
	PrefixFlagStringID    = "Flag"

	NamelessCommandIDPostfix = "NamelessCommand"
)

const (
	helpCommandIDStr   = "PrintHelpInfo"
	helpCommandComment = "print help info"
)

// IDTemplateDataCreator - creates slices of id template data
type IDTemplateDataCreator struct {
	dashRemover *regexp.Regexp
}

// NewIDTemplateCreator - IDTemplateDataCreator object constructor
func NewIDTemplateCreator() IDTemplateDataCreator {
	return IDTemplateDataCreator{dashRemover: regexp.MustCompile("-+")}
}

// RemoveDashes - removes all dashes from a string
func (i IDTemplateDataCreator) RemoveDashes(str string) string {
	return i.dashRemover.ReplaceAllString(str, "")
}

// CreateID - creates ID string by call name
func (i IDTemplateDataCreator) CreateID(prefix string, callName string) string {
	callName = i.RemoveDashes(callName)

	callNameRunes := []rune(callName)
	callNameRuneCount := len(callNameRunes)

	if callNameRuneCount == 0 {
		return ""
	}

	res := prefix + string(unicode.ToUpper(callNameRunes[0]))
	if callNameRuneCount > 1 {
		res += string(callNameRunes[1:])
	}

	return res
}

// CreateIDTemplateData - creates IDTemplateData slices for commands and flags
func (i IDTemplateDataCreator) CreateIDTemplateData(
	commandDescriptions []*configYaml.CommandDescription,
	helpCommandDescription *configYaml.HelpCommandDescription,
	nullCommandDescription *configYaml.NamelessCommandDescription,
	flagDescriptionMap map[configYaml.Flag]*configYaml.FlagDescription,
) (
	commandsIDTemplateData map[configYaml.Command]*IDTemplateData,
	nullCommandIDTemplateData *IDTemplateData,
	flagsIDTemplateData map[configYaml.Flag]*IDTemplateData) {

	var (
		j, k               int
		callName           string
		commandId          string
		commandDescription *configYaml.CommandDescription
	)

	commandsIDTemplateData = make(map[configYaml.Command]*IDTemplateData, len(commandDescriptions))
	flagsIDTemplateData = make(map[configYaml.Flag]*IDTemplateData, len(flagDescriptionMap))

	// standard commands
	for j = range commandDescriptions {
		commandDescription = commandDescriptions[j]

		callName = string(commandDescription.GetCommand())
		commandId = i.CreateID(PrefixCommandID, callName)

		commandsIDTemplateData[commandDescription.GetCommand()] = NewIDTemplateData(
			commandId,
			i.CreateID(PrefixCommandStringID, callName),
			callName,
			commandDescription.GetDescriptionHelpInfo())

		for k = range commandDescription.GetAdditionalCommands() {
			callName = string(commandDescription.GetAdditionalCommands()[k])
			commandsIDTemplateData[commandDescription.GetAdditionalCommands()[k]] = NewIDTemplateData(
				commandId,
				i.CreateID(PrefixCommandStringID, callName),
				callName,
				commandDescription.GetDescriptionHelpInfo())
		}
	}

	// help command
	commandId = i.CreateID(PrefixCommandID, helpCommandIDStr)
	callName = string(helpCommandDescription.GetCommand())

	commandsIDTemplateData[helpCommandDescription.GetCommand()] = NewIDTemplateData(
		commandId,
		i.CreateID(PrefixCommandStringID, callName),
		callName,
		helpCommandComment)

	for k = range helpCommandDescription.GetAdditionalCommands() {
		callName = string(helpCommandDescription.GetAdditionalCommands()[k])
		commandsIDTemplateData[helpCommandDescription.GetAdditionalCommands()[k]] = NewIDTemplateData(
			commandId,
			i.CreateID(PrefixCommandStringID, callName),
			callName,
			helpCommandComment)
	}

	// null command
	if nullCommandDescription != nil {
		nullCommandIDTemplateData = NewIDTemplateData(
			i.CreateID(PrefixCommandID, NamelessCommandIDPostfix),
			"",
			"",
			nullCommandDescription.GetDescriptionHelpInfo())
	}

	// flags
	for _, flagDescription := range flagDescriptionMap {
		callName = string(flagDescription.GetFlag())
		flagsIDTemplateData[flagDescription.GetFlag()] = NewIDTemplateData(
			"",
			i.CreateID(PrefixFlagStringID, callName),
			callName,
			flagDescription.GetDescriptionHelpInfo())
	}

	return commandsIDTemplateData, nullCommandIDTemplateData, flagsIDTemplateData
}
