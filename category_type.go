package setdata_common

type CategoryType string

var (
	Panel  CategoryType = "panel"
	Filler CategoryType = "filler"
)

var (
	categoryTypeToString = map[CategoryType]string{
		Panel:  "panel",
		Filler: "filler",
	}
	stringToCategoryType = map[string]CategoryType{
		"panel":  Panel,
		"filler": Filler,
	}
)

func (c CategoryType) ToString() string {
	return categoryTypeToString[c]
}

func ToCategoryType(s string) CategoryType {
	return stringToCategoryType[s]
}

func IsCategoryExist(s string) bool {
	categoryTypes := []string{"panel", "filler"}
	for _, v := range categoryTypes {
		if v == s {
			return true
		}
	}
	return false
}
