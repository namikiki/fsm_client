package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"fsm_client/pkg/types"

	"github.com/BurntSushi/toml"
	"github.com/google/uuid"
)

const (
	Application   = "fsm"
	DefaultConfig = "[Device]\nClientID = \"%s\"\n" +
		"Platform = \"%s\"\n" +
		"[Server]\nBaseUrl = \"%s\"\n" +
		"WebSocketUrl = \"%s\"\n"
)

func ReadConfigFile() (*types.Config, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(dir, Application, "config")
	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
		log.Println("配置文件不存在，正在生成默认配置")
		if err := WriteConfigToFile(); err != nil {
			return nil, err
		}
	}

	var cf types.Config
	if _, err := toml.DecodeFile(configPath, &cf); err != nil {
		panic(err)
	}

	return &cf, nil
}

func DeleteConfig() error {
	dir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	configPath := filepath.Join(dir, Application, "config")
	return os.Remove(configPath)

}

func WriteDefaultConfig() {

}

func WriteConfigToFile() error {
	dir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	file, err := os.OpenFile(filepath.Join(dir, Application, "config"), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf(DefaultConfig,
		GenerateClientID(),
		GetPlatformType(),
		"http://127.0.0.1:8080",
		"ws://127.0.0.1:8080",
	))

	return err
}

func GenerateClientID() string {
	return uuid.NewString()
}

func GetPlatformType() string {
	return runtime.GOOS
}

func NewIgnoreConfig() (*types.Ignore, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}
	ignorePath := filepath.Join(dir, Application, "ignore")

	var ignore types.Ignore
	_, err = toml.DecodeFile(ignorePath, &ignore)
	return &ignore, err
}
