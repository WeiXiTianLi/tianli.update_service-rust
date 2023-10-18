package src

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	DatabasePath string `toml:"database_path"`
	Port         int    `toml:"port"`
}

// 检查配置文件是否存在
func checkConfigFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || os.IsExist(err)

}

// 创建默认配置
func createDefaultConfig(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	tomlObject := Config{
		DatabasePath: "data.db",
		Port:         12380,
	}
	encoder := toml.NewEncoder(file)
	err = encoder.Encode(tomlObject)
	if err != nil {
		return err
	}
	return nil
}

// 如果配置文件不存在，则创建默认配置
func createDefaultConfigIfNotExist(filePath string) error {
	if !checkConfigFileExist(filePath) {
		return createDefaultConfig(filePath)
	}
	return nil
}

func loadConfig(filePath string) (*Config, error) {
	// 如果配置文件不存在，则创建默认配置
	err := createDefaultConfigIfNotExist(filePath)
	if err != nil {
		return nil, err
	}
	var config Config
	_, err = toml.DecodeFile(filePath, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil

}
