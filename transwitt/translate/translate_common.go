package translate

const APIDefaultURL = "https://openapi.naver.com/v1/papago"
const NMTAPIPath = "/n2mt"

func (p Papago) AvailableSourceLanguages() []string {
	return []string{"ko", "en", "jp", "zh-CN", "zh-TW", "vi", "id", "th", "de", "ru", "es", "it", "fr"}
}

func (p Papago) AvailableTargetLanguages(sSource string) []string {
	switch sSource {
	case "ko":
		return []string{"en", "ja", "zh-CN", "zh-TW", "vi", "id", "th", "de", "ru", "es", "it", "fr"}
	case "en":
		return []string{"ja", "fr", "zh-CN", "zh-TW", "ko"}
	case "ja":
		return []string{"zh-CN", "zh-TW", "ko", "en"}
	case "zh-CN":
		return []string{"zh-TW", "ko", "en", "jp"}
	case "zh-TW":
		return []string{"ko", "en", "ja", "zh-CN"}
	case "vi":
		return []string{"ko"}
	case "id":
		return []string{"ko"}
	case "th":
		return []string{"de"}
	case "de":
		return []string{"ko"}
	case "ru":
		return []string{"ko"}
	case "es":
		return []string{"ko"}
	case "it":
		return []string{"ko"}
	case "fr":
		return []string{"ko", "en"}
	}
	return nil
}
