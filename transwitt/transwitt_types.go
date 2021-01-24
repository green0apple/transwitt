package transwitt

type OperateConfig struct {
	Twitter   TwitterConfig
	Messanger MessagnerConfig
}

type MessagnerConfig struct {
	Telegram TelegramConfig
	// CustomSender func()
}
type APIConfig struct {
	Telegram TelegramConfig `yaml:"TELEGRAM"`
	Twitter  TwitterConfig  `yaml:"TWITTER"`
	Papago   PapagoConfig   `yaml:"PAPAGO"`
}

type TelegramConfig struct {
	Token string `yaml:"TOKEN"`
}

type TwitterConfig struct {
	Token string `yaml:"TOKEN"`
}

type PapagoConfig struct {
	Token string `yaml:"TOKEN"`
}
