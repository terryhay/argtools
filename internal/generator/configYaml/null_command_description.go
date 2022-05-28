package configYaml

import "fmt"

// NullCommandDescription -
type NullCommandDescription struct {
	DescriptionHelpInfo string

	// optional
	RequiredFlags        []Flag
	OptionalFlags        []Flag
	ArgumentsDescription *ArgumentsDescription
}

// GetDescriptionHelpInfo - DescriptionHelpInfo field getter
func (i *NullCommandDescription) GetDescriptionHelpInfo() string {
	if i == nil {
		return ""
	}
	return i.DescriptionHelpInfo
}

// GetRequiredFlags - RequiredFlags field getter
func (i *NullCommandDescription) GetRequiredFlags() []Flag {
	if i == nil {
		return nil
	}
	return i.RequiredFlags
}

// GetOptionalFlags - OptionalFlags field getter
func (i *NullCommandDescription) GetOptionalFlags() []Flag {
	if i == nil {
		return nil
	}
	return i.OptionalFlags
}

// GetArgumentsDescription - ArgumentsDescription field getter
func (i *NullCommandDescription) GetArgumentsDescription() *ArgumentsDescription {
	if i == nil {
		return nil
	}
	return i.ArgumentsDescription
}

type nullCommandDescriptionSource struct {
	DescriptionHelpInfo string `yaml:"description_help_info"`

	// optional
	RequiredFlags        []Flag                `yaml:"required_flags"`
	OptionalFlags        []Flag                `yaml:"optional_flags"`
	ArgumentsDescription *ArgumentsDescription `yaml:"arguments_description"`
}

// UnmarshalYAML - custom unmarshal logic with checking required fields
func (i *NullCommandDescription) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	source := new(nullCommandDescriptionSource)
	if err = unmarshal(&source); err != nil {
		return err
	}

	if len(source.DescriptionHelpInfo) == 0 {
		return fmt.Errorf(`commandDescription unmarshal error: no required field "description_help_info"`)
	}
	i.DescriptionHelpInfo = source.DescriptionHelpInfo

	// don't check optional fields
	i.RequiredFlags = source.RequiredFlags
	i.OptionalFlags = source.OptionalFlags
	i.ArgumentsDescription = source.ArgumentsDescription

	return nil
}
