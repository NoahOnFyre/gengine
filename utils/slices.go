package utils

func RemoveElement(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}
