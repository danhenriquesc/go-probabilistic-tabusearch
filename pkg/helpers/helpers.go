package helpers

func Contains(needed string, tabuList *[]string) bool{
	for _, item := range *tabuList {
		if item == needed {
			return true
		}
	}
	return false
}