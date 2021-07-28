package setdata_common

type LoginType string

var (
	Email       LoginType = "email"
	PhoneNumber LoginType = "phone_number"
)

var (
	loginTypeToString = map[LoginType]string{
		Email:       "email",
		PhoneNumber: "phone_number",
	}
	stringToLoginType = map[string]LoginType{
		"email":        Email,
		"phone_number": PhoneNumber,
	}
)

func (l LoginType) ToString() string {
	return loginTypeToString[l]
}

func ToLoginType(s string) LoginType {
	return stringToLoginType[s]
}
