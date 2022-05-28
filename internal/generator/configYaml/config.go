package configYaml

import "fmt"

// Version - yaml file format version
type Version string

// Flag - command flag name type
type Flag string

// Command - application command (which contains flags and arguments) type
type Command string

// Config - code struct of a config yaml file
type Config struct {
	Version                Version
	AppHelpDescription     *AppHelpDescription
	HelpCommandDescription *HelpCommandDescription

	// one or more of these field must be set
	NullCommandDescription *NullCommandDescription
	CommandDescriptions    []*CommandDescription

	// optional
	FlagDescriptions []*FlagDescription
}

// GetVersion - Version field getter
func (i *Config) GetVersion() Version {
	if i == nil {
		return ""
	}
	return i.Version
}

// GetAppHelpDescription - AppHelpDescription field getter
func (i *Config) GetAppHelpDescription() *AppHelpDescription {
	if i == nil {
		return nil
	}
	return i.AppHelpDescription
}

// GetHelpCommandDescription - HelpCommandDescription field getter
func (i *Config) GetHelpCommandDescription() *HelpCommandDescription {
	if i == nil {
		return nil
	}
	return i.HelpCommandDescription
}

// GetNullCommandDescription - NullCommandDescription field getter
func (i *Config) GetNullCommandDescription() *NullCommandDescription {
	if i == nil {
		return nil
	}
	return i.NullCommandDescription
}

// GetCommandDescriptions - CommandDescriptions field getter
func (i *Config) GetCommandDescriptions() []*CommandDescription {
	if i == nil {
		return nil
	}
	return i.CommandDescriptions
}

// GetFlagDescriptions - FlagDescriptions field getter
func (i *Config) GetFlagDescriptions() []*FlagDescription {
	if i == nil {
		return nil
	}
	return i.FlagDescriptions
}

type configSource struct {
	Version                string                  `yaml:"version"`
	AppHelpDescription     *AppHelpDescription     `yaml:"app_help_description"`
	HelpCommandDescription *HelpCommandDescription `yaml:"help_command_description"`

	// one or more of these field must be set
	NullCommandDescription *NullCommandDescription `yaml:"null_command_description"`
	CommandDescriptions    []*CommandDescription   `yaml:"command_descriptions"`

	// optional
	FlagDescriptions []*FlagDescription `yaml:"flag_descriptions"`
}

// UnmarshalYAML - custom unmarshal logic with checking required fields
func (i *Config) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	source := new(configSource)
	if err = unmarshal(&source); err != nil {
		return err
	}

	if len(source.Version) == 0 {
		return fmt.Errorf(`config unmarshal error: no required field "version"`)
	}
	i.Version = Version(source.Version)

	if source.AppHelpDescription == nil {
		return fmt.Errorf(`config unmarshal error: no required field "app_help_description"`)
	}
	i.AppHelpDescription = source.AppHelpDescription

	if source.HelpCommandDescription == nil {
		return fmt.Errorf(`config unmarshal error: no required field "help_command_description"`)
	}
	i.HelpCommandDescription = source.HelpCommandDescription

	if len(source.CommandDescriptions) == 0 && source.NullCommandDescription == nil {
		return fmt.Errorf(`config unmarshal error: one or more of felds "null_command_description" or "command_descriptions" must be set`)
	}

	// don't check optional fields
	i.CommandDescriptions = source.CommandDescriptions
	i.FlagDescriptions = source.FlagDescriptions
	i.NullCommandDescription = source.NullCommandDescription

	return nil
}
