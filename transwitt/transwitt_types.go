package transwitt

import (
	"time"
)

type OperateConfig struct {
	Twitter   TwitterConfig
	Messenger MessegnerConfig
}

type MessegnerConfig struct {
	Telegram TelegramConfig
	// CustomSender func()
}

type TelegramConfig struct {
	Token string
	Admin int64
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
	TweetTime time.Time
}

type PapagoConfig struct {
	Token string
}
