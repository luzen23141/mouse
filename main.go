package main

import (
	"fmt"
	"mouse/cmd"
	"mouse/pkg/helper"
	"mouse/pkg/lib/env"
	"mouse/pkg/lib/log"
)

func main() {
	err := env.BindToml("config", ".", &helper.EnvCfg)
	if err != nil {
		log.LogError(log.FileError, err, "環境變數綁定失敗")
		fmt.Println("環境變數綁定失敗", err)
		return
	}

	cmd.Execute()
}
