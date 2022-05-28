package argParserConfig

type NullCommandDescription struct {
	ID                  CommandID
	DescriptionHelpInfo string
	ArgDescription      *ArgumentsDescription
	RequiredFlags       map[Flag]bool
	OptionalFlags       map[Flag]bool
}

// GetID - ID field getter
func (i *NullCommandDescription) GetID() CommandID {
	if i == nil {
		return CommandIDUndefined
	}
	return i.ID
}

// GetDescriptionHelpInfo - DescriptionHelpInfo field getter
func (i *NullCommandDescription) GetDescriptionHelpInfo() string {
	if i == nil {
		return ""
	}
	return i.DescriptionHelpInfo
}

// GetArgDescription - ArgDescription field getter
func (i *NullCommandDescription) GetArgDescription() *ArgumentsDescription {
	if i == nil {
		return nil
	}
	return i.ArgDescription
}

// GetRequiredFlags - RequiredFlags field getter
func (i *NullCommandDescription) GetRequiredFlags() map[Flag]bool {
	if i == nil {
		return nil
	}
	return i.RequiredFlags
}

// GetOptionalFlags - OptionalFlags field getter
func (i *NullCommandDescription) GetOptionalFlags() map[Flag]bool {
	if i == nil {
		return nil
	}
	return i.OptionalFlags
}
