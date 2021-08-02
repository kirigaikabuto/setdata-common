package setdata_common

type Field struct {
	Name        string    `json:"name"`
	Type        FieldType `json:"type"`
	Placeholder string    `json:"placeholder"`
}

type FieldType string

var (
	Text              FieldType = "text"
	Radio             FieldType = "radio"
	CheckBox          FieldType = "checkbox"
	fieldTypeToString           = map[FieldType]string{
		Text:     "text",
		Radio:    "radio",
		CheckBox: "checkbox",
	}
	stringToFieldType = map[string]FieldType{
		"text":     Text,
		"radio":    Radio,
		"checkbox": CheckBox,
	}
)

func GetFieldType(typeVal string) (FieldType, error) {
	val, ok := stringToFieldType[typeVal]
	if !ok {
		return "", ErrCurrentAclActionNotExist
	}
	return val, nil
}

func GetFieldTypeString(typeVal FieldType) (string, error) {
	val, ok := fieldTypeToString[typeVal]
	if !ok {
		return "", ErrCurrentAclActionNotExist
	}
	return val, nil
}
