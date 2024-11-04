package env

import (
	"fmt"
	"reflect"

	"github.com/shopspring/decimal"
	"github.com/spf13/viper"
)

// BindToml config 設置配置文件名 (不含副檔名)，path 路徑 ("." 為當前目錄)，target需要是指標
func BindToml(config, path string, target interface{}) error {
	viper.SetConfigName(config) // 設置配置文件名 (不含副檔名)
	viper.AddConfigPath(path)   // 添加配置文件所在的目錄
	viper.SetConfigType("toml") // 設置配置文件類型

	// 加入 decodeHook
	decimalDecodeHook := func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		if t != reflect.TypeOf(decimal.Decimal{}) {
			return data, nil
		}
		var err error
		d := decimal.Zero
		switch v := data.(type) {
		case string:
			d, err = decimal.NewFromString(v)
			if err != nil {
				return nil, fmt.Errorf("無法解析 decimal: %w", err)
			}
		case int64:
			d = decimal.NewFromInt(v)
		case float64:
			d = decimal.NewFromFloat(v)
		}
		return d, nil
	}

	// 讀取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// 將配置信息反序列化到 Config 結構體
	if err := viper.Unmarshal(target, viper.DecodeHook(decimalDecodeHook)); err != nil {
		return err
	}

	return nil
}
