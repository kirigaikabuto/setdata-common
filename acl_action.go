package setdata_common

type AclAction string

var (
	ActionRead        AclAction = "read"
	ActionWrite       AclAction = "write"
	ActionFullControl AclAction = "full_control"
)

var (
	aclActionToString = map[AclAction]string{
		ActionRead:        "read",
		ActionWrite:       "write",
		ActionFullControl: "full_control",
	}
	stringToAclAction = map[string]AclAction{
		"read":         ActionRead,
		"write":        ActionWrite,
		"full_control": ActionFullControl,
	}
)

func GetAclAction(action string) (AclAction, error) {
	val, ok := stringToAclAction[action]
	if !ok {
		return "", ErrCurrentAclActionNotExist
	}
	return val, nil
}

func GetAclActiomString(action AclAction) (string, error) {
	val, ok := aclActionToString[action]
	if !ok {
		return "", ErrCurrentAclActionNotExist
	}
	return val, nil
}
