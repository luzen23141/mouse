package main

import (
	"fmt"
	"github.com/luzen23141/mouse/cmd"
	"github.com/luzen23141/mouse/pkg/helper"
	"github.com/luzen23141/mouse/pkg/lib/env"
	"github.com/luzen23141/mouse/pkg/lib/log"
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
