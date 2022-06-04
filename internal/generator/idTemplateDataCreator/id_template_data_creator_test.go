package idTemplateDataCreator

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"testing"
)

func TestIDTemplateDataCreator(t *testing.T) {
	t.Parallel()

	command := configYaml.Command(gofakeit.Color())
	additionalCommand := configYaml.Command(gofakeit.Color())
	commandDescriptionHelpInfo := gofakeit.Name()

	helpCommand := configYaml.Command(gofakeit.Color())
	additionalHelpCommand := configYaml.Command(gofakeit.Color())

	flag := configYaml.Flag(gofakeit.Color())

	creator := NewIDTemplateCreator()
	commandsIDTemplateData, nullCommandIDTemplateData, flagsIDTemplateData := creator.CreateIDTemplateData(
		[]*configYaml.CommandDescription{
			{
				Command: command,
				AdditionalCommands: []configYaml.Command{
					additionalCommand,
				},
				DescriptionHelpInfo: commandDescriptionHelpInfo,
			},
			{
				// fake empty command
			},
		},
		&configYaml.HelpCommandDescription{
			Command: helpCommand,
			AdditionalCommands: []configYaml.Command{
				additionalHelpCommand,
			},
		},
		&configYaml.NullCommandDescription{},
		map[configYaml.Flag]*configYaml.FlagDescription{
			flag: {
				Flag: flag,
			},
		})

	expectedCommandID := creator.CreateID(PrefixCommandID, string(command))
	expectedHelpCommandID := creator.CreateID(PrefixCommandID, string(helpCommand))
	expectedCommandsIDTemplateData := map[configYaml.Command]*IDTemplateData{
		command: {
			id:       expectedCommandID,
			stringID: creator.CreateID(PrefixCommandStringID, string(command)),
			callName: string(command),
			comment:  commandDescriptionHelpInfo,
		},
		additionalCommand: {
			id:       expectedCommandID,
			stringID: creator.CreateID(PrefixCommandStringID, string(additionalCommand)),
			callName: string(additionalCommand),
			comment:  commandDescriptionHelpInfo,
		},
		helpCommand: {
			id:       expectedHelpCommandID,
			stringID: creator.CreateID(PrefixCommandStringID, string(helpCommand)),
			callName: string(helpCommand),
			comment:  helpCommandComment,
		},
		additionalHelpCommand: {
			id:       expectedHelpCommandID,
			stringID: creator.CreateID(PrefixCommandStringID, string(additionalHelpCommand)),
			callName: string(additionalHelpCommand),
			comment:  helpCommandComment,
		},
		"": {},
	}

	require.Equal(t, len(expectedCommandsIDTemplateData), len(commandsIDTemplateData))
	for expectedCommand, expectedIDTemplateData := range expectedCommandsIDTemplateData {
		idTemplateData, ok := commandsIDTemplateData[expectedCommand]
		require.True(t, ok)

		require.Equal(t, expectedIDTemplateData, idTemplateData)
	}

	require.Equal(t, &IDTemplateData{id: "CommandIDNullCommand"}, nullCommandIDTemplateData)

	flagIDTemplateData, ok := flagsIDTemplateData[flag]
	require.True(t, ok)
	require.Equal(t, &IDTemplateData{
		id:       "",
		stringID: creator.CreateID(PrefixFlagStringID, string(flag)),
		callName: string(flag),
	}, flagIDTemplateData)
}
