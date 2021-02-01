package translate

type Papago struct {
	ClientID     string
	ClientSecret string
}

type PapagoResponse struct {
	Message      PapagoResponseMessage `json:"message"`
	ErrorMessage string                `json:"errorMessage"`
	ErrorCode    string                `json:"errorCode"`
}

type PapagoResponseMessage struct {
	Type    string               `json:"@type"`
	Service string               `json:"@service"`
	Version string               `json:"@version"`
	Result  PapagoResponseResult `json:"result"`
}

type PapagoResponseResult struct {
	TranslatedText string `json:"translatedText"`
}

type PapagoRequest struct {
	Target string `json:"target"`
	Source string `json:"source"`
	Text   string `json:"text"`
}
