package transwitt

import (
	"time"
)

type OperateConfig struct {
	Twitter    TwitterConfig
	Messenger  MessegnerConfig
	Translator TranslatorConfig
}

type MessegnerConfig struct {
	Telegram TelegramConfig
	Discord  DiscordConfig
	// CustomSender func()
}

type TelegramConfig struct {
	Token string
	Admin int64
}

type DiscordConfig struct {
	Token string
	Admin string
}

type TwitterConfig struct {
	ConsumerKey     string
	ConsumerSecret  string
	Users           TwitterUsers
	TPS             float32
	CountPerRequest int
}

type TwitterUsers []TwitterUser

type TwitterUser struct {
	ScreenID  string
	Nickname  string
	Language  TranslateLanguage
	TweetTime time.Time
}

type TranslateLanguage struct {
	Source string
	Target string
}

type TranslatorConfig struct {
	Papago PapagoConfig
}

type PapagoConfig struct {
	ClientID     string
	ClientSecret string
}
