package argParserImpl

import "github.com/terryhay/argtools/pkg/argParserConfig"

func nameless2commandDescription(namelessCommandDescription *argParserConfig.NamelessCommandDescription) *argParserConfig.CommandDescription {
	if namelessCommandDescription == nil {
		return nil
	}
	return &argParserConfig.CommandDescription{
		ID:                  namelessCommandDescription.GetID(),
		DescriptionHelpInfo: namelessCommandDescription.GetDescriptionHelpInfo(),
		ArgDescription:      namelessCommandDescription.GetArgDescription(),
		RequiredFlags:       namelessCommandDescription.GetRequiredFlags(),
		OptionalFlags:       namelessCommandDescription.GetOptionalFlags(),
	}
}
