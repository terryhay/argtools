package argParserConfig

// NamelessCommandDescription - contains a specification of a command without call name
type NamelessCommandDescription struct {
	ID                  CommandID
	DescriptionHelpInfo string
	ArgDescription      *ArgumentsDescription
	RequiredFlags       map[Flag]bool
	OptionalFlags       map[Flag]bool
}

// GetID - ID field getter
func (i *NamelessCommandDescription) GetID() CommandID {
	if i == nil {
		return CommandIDUndefined
	}
	return i.ID
}

// GetDescriptionHelpInfo - DescriptionHelpInfo field getter
func (i *NamelessCommandDescription) GetDescriptionHelpInfo() string {
	if i == nil {
		return ""
	}
	return i.DescriptionHelpInfo
}

// GetArgDescription - ArgDescription field getter
func (i *NamelessCommandDescription) GetArgDescription() *ArgumentsDescription {
	if i == nil {
		return nil
	}
	return i.ArgDescription
}

// GetRequiredFlags - RequiredFlags field getter
func (i *NamelessCommandDescription) GetRequiredFlags() map[Flag]bool {
	if i == nil {
		return nil
	}
	return i.RequiredFlags
}

// GetOptionalFlags - OptionalFlags field getter
func (i *NamelessCommandDescription) GetOptionalFlags() map[Flag]bool {
	if i == nil {
		return nil
	}
	return i.OptionalFlags
}
