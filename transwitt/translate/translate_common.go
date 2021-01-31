package translate

const PapagoAPIURL = "https://openapi.naver.com/v1/papago/n2mt"

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
