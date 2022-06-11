package argParserConfig

// NewArgParserConfig - ArgParserConfig constructor
func NewArgParserConfig(
	appDescription ApplicationDescription,
	flagDescriptions map[Flag]*FlagDescription,
	commandDescriptions []*CommandDescription,
	namelessCommandDescription *NamelessCommandDescription,
) ArgParserConfig {

	return ArgParserConfig{
		AppDescription:             appDescription,
		FlagDescriptions:           flagDescriptions,
		CommandDescriptions:        commandDescriptions,
		NamelessCommandDescription: namelessCommandDescription,
	}
}
