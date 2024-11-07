package helper

type EnvCfgStruct struct {
	Gin struct {
		Debug bool `mapstructure:"debug"`
		Port  int  `mapstructure:"port"`
	} `mapstructure:"gin"`
	Redis map[string]struct {
		Addr     string `mapstructure:"addr"`
		DB       int    `mapstructure:"db"`
		Password string `mapstructure:"password"`
	}
	Telegram struct {
		Token  string `mapstructure:"token"`
		ChatID string `mapstructure:"chat_id"`
	}
}
