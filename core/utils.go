package core

func GenericTypeToStringList(items []GenericType) []string {
	var newItems []string
	for _, item := range items {
		newItems = append(newItems, item.Name)
	}
	return newItems
}
