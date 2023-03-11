package config

import (
	"log"
	"os"
	"testing"
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
