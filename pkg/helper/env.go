package helper

type EnvCfgStruct struct {
	Gin struct {
		Debug bool `mapstructure:"debug"`
		Port  int  `mapstructure:"port"`
	} `mapstructure:"gin"`
	Redis map[string]struct {
		Addr     string `mapstructure:"addr"`
		Db       int    `mapstructure:"db"`
		Password string `mapstructure:"password"`
	}
	Telegram struct {
		Token  string `mapstructure:"token"`
		ChatId string `mapstructure:"chat_id"`
	}
}
