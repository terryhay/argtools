package idTemplateDataCreator

// IDTemplateData - data for fill up templates
type IDTemplateData struct {
	id       string
	stringID string
	callName string
	comment  string
}

// NewIDTemplateData - IDTemplateData object constructor
func NewIDTemplateData(callName, id, stringID, comment string) *IDTemplateData {
	return &IDTemplateData{
		callName: callName,
		id:       id,
		stringID: stringID,
		comment:  comment,
	}
}

// GetID - id field getter
func (i *IDTemplateData) GetID() string {
	if i == nil {
		return ""
	}
	return i.id
}

// GetStringID - stringID field getter
func (i *IDTemplateData) GetStringID() string {
	if i == nil {
		return ""
	}
	return i.stringID
}

// GetCallName - field callName getter
func (i *IDTemplateData) GetCallName() string {
	if i == nil {
		return ""
	}
	return i.callName
}

// GetComment - field comment getter
func (i *IDTemplateData) GetComment() string {
	if i == nil {
		return ""
	}
	return i.comment
}
