package argParserConfig

// NewArgParserConfig - ArgParserConfig constructor
func NewArgParserConfig(
	appDescription ApplicationDescription,
	flagDescriptions map[Flag]*FlagDescription,
	commandDescriptions []*CommandDescription,
	nullCommandDescription *NullCommandDescription,
) ArgParserConfig {

	return ArgParserConfig{
		AppDescription:         appDescription,
		FlagDescriptions:       flagDescriptions,
		CommandDescriptions:    commandDescriptions,
		NullCommandDescription: nullCommandDescription,
	}
}
