package config

import (
	"log"
	"os"
	"testing"

	"fsm_client/pkg/types"

	"github.com/BurntSushi/toml"
)

func TestGetUserConfigDir(t *testing.T) {
	dir, err := os.UserConfigDir()
	if err != nil {
		log.Println(err)
	}

	log.Println(dir)
}

func TestWriteConfigToFile(t *testing.T) {
	err := WriteConfigToFile()
	if err != nil {
		log.Println(err)
	}
}

func TestReadConfigFile(t *testing.T) {
	_, _ = ReadConfigFile()

}

func TestDeleteConfig(t *testing.T) {
	err := DeleteConfig()
	if err != nil {
		log.Println(err)
	}
}

func TestTUY(t *testing.T) {
	// Config 是一个配置对象

	// 从文件中读取配置并解码为 Config 对象
	var cfg types.Ignore
	if _, err := toml.DecodeFile("config.toml", &cfg); err != nil {
		log.Fatal(err)
	}

	log.Println(cfg)

	//// 将修改后的配置对象编码为 TOML 格式并写入文件
	//f, err := os.Create("config.toml")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer f.Close()
	//
	//if err := toml.NewEncoder(f).Encode(cfg); err != nil {
	//	log.Fatal(err)
	//}

}

func TestName1(t *testing.T) {
	stat, err := os.Stat("/adwdawdawd")
	log.Println(stat, err)
}
