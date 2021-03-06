package translate

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func (p Papago) AvailableSourceLanguages() []string {
	return []string{"ko", "en", "ja", "zh-CN", "zh-TW", "vi", "id", "th", "de", "ru", "es", "it", "fr"}
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

func (p Papago) GetTranslate(sSource, sTarget, sText string) (string, error) {
	if indexOf(sSource, p.AvailableSourceLanguages()) == -1 {
		return "", errors.New("Source language " + sSource + " is not available in Papago")
	}
	if indexOf(sTarget, p.AvailableTargetLanguages(sSource)) == -1 {
		return "", errors.New("Target language " + sTarget + " cannot translate from " + sSource + " in Papago")
	}
	preq := PapagoRequest{
		Target: sTarget,
		Source: sSource,
		Text:   sText,
	}
	bufRequest, err := json.Marshal(preq)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest(http.MethodPost, PapagoAPIURL, bytes.NewBuffer(bufRequest))
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("X-Naver-Client-Id", p.ClientID)
	req.Header.Add("X-Naver-Client-Secret", p.ClientSecret)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer req.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	pres := PapagoResponse{}
	if err = json.Unmarshal(resBody, &pres); err != nil {
		return "", err
	}
	if pres.ErrorMessage != "" {
		return "", errors.New(string(resBody))
	}
	return pres.Message.Result.TranslatedText, nil
}
