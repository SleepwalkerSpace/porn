package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Load 导入配置 github.com/spf13/viper
func Load(cfgPath, cfgName, cfgtype string, config interface{}) error {
	imp := viper.New()
	imp.AddConfigPath(cfgPath)
	imp.SetConfigName(cfgName)
	imp.SetConfigType(cfgtype)
	imp.ReadInConfig()
	fmt.Println(imp.AllKeys())
	if err := imp.Unmarshal(config); err != nil {
		return err
	}
	return nil
}
