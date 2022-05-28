package tests

import (
	"argtools/internal/generator/configDataExtractor"
	"argtools/internal/generator/configYaml"
	"argtools/pkg/argtoolsError"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractCommandDescriptionMapErrors(t *testing.T) {
	t.Parallel()

	testData := []struct {
		caseName            string
		commandDescriptions []*configYaml.CommandDescription
		expectedErrorCode   argtoolsError.Code
	}{
		{
			caseName: "single_empty_command_description",
			commandDescriptions: []*configYaml.CommandDescription{
				nil,
			},
			expectedErrorCode: argtoolsError.CodeUndefinedError,
		},
		{
			caseName: "empty_command_description_in_front",
			commandDescriptions: []*configYaml.CommandDescription{
				nil,
				{},
			},
			expectedErrorCode: argtoolsError.CodeUndefinedError,
		},
		{
			caseName: "empty_command_description_in_back",
			commandDescriptions: []*configYaml.CommandDescription{
				{},
				nil,
			},
			expectedErrorCode: argtoolsError.CodeUndefinedError,
		},

		{
			caseName: "duplicate_flag_descriptions_in_additional_commands_case_1",
			commandDescriptions: []*configYaml.CommandDescription{
				{
					Command: "command",
					AdditionalCommands: []configYaml.Command{
						"command",
					},
				},
			},
			expectedErrorCode: argtoolsError.CodeConfigContainsDuplicateCommands,
		},
		{
			caseName: "duplicate_flag_descriptions_in_additional_commands_case_10",
			commandDescriptions: []*configYaml.CommandDescription{
				{
					Command: "command",
					AdditionalCommands: []configYaml.Command{
						"command",
						"command1",
					},
				},
			},
			expectedErrorCode: argtoolsError.CodeConfigContainsDuplicateCommands,
		},
		{
			caseName: "duplicate_flag_descriptions_in_additional_commands_case_01",
			commandDescriptions: []*configYaml.CommandDescription{
				{
					Command: "command",
					AdditionalCommands: []configYaml.Command{
						"command1",
						"command",
					},
				},
			},
			expectedErrorCode: argtoolsError.CodeConfigContainsDuplicateCommands,
		},
		{
			caseName: "duplicate_flag_descriptions",
			commandDescriptions: []*configYaml.CommandDescription{
				{
					Command: "command",
				},
				{
					Command: "command",
				},
			},
			expectedErrorCode: argtoolsError.CodeConfigContainsDuplicateCommands,
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			flagDescriptionMap, err := configDataExtractor.ExtractCommandDescriptionMap(td.commandDescriptions)
			require.Nil(t, flagDescriptionMap)
			require.NotNil(t, err)
			require.Equal(t, td.expectedErrorCode, err.Code())
		})
	}
}

func TestExtractCommandDescriptionMap(t *testing.T) {
	t.Parallel()

	testData := []struct {
		caseName            string
		commandDescriptions []*configYaml.CommandDescription
		expectedMap         map[configYaml.Command]*configYaml.CommandDescription
	}{
		{
			caseName:            "no_flag_description",
			commandDescriptions: nil,
		},

		{
			caseName: "single_flag_description",
			commandDescriptions: []*configYaml.CommandDescription{
				{
					Command: "command",
				},
			},
			expectedMap: map[configYaml.Command]*configYaml.CommandDescription{
				"command": {
					Command: "command",
				},
			},
		},
		{
			caseName: "single_flag_description_with_additional_command",
			commandDescriptions: []*configYaml.CommandDescription{
				{
					Command: "command",
					AdditionalCommands: []configYaml.Command{
						"command1",
					},
				},
			},
			expectedMap: map[configYaml.Command]*configYaml.CommandDescription{
				"command": {
					Command: "command",
					AdditionalCommands: []configYaml.Command{
						"command1",
					},
				},
			},
		},

		{
			caseName: "two_flag_descriptions",
			commandDescriptions: []*configYaml.CommandDescription{
				{
					Command: "command1",
				},
				{
					Command: "command2",
				},
			},
			expectedMap: map[configYaml.Command]*configYaml.CommandDescription{
				"command1": {
					Command: "command1",
				},
				"command2": {
					Command: "command2",
				},
			},
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			flagDescriptionMap, err := configDataExtractor.ExtractCommandDescriptionMap(td.commandDescriptions)
			require.Nil(t, err)

			require.Equal(t, len(td.expectedMap), len(flagDescriptionMap))

			for command, expectedCommandDescription := range td.expectedMap {
				flagDescription, contain := flagDescriptionMap[command]
				require.True(t, contain)
				require.Equal(t, expectedCommandDescription.Command, flagDescription.Command)
			}
		})
	}
}
