package tools

import "github.com/spf13/viper"

func LoadConfig() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// Default values
	viper.SetDefault("BIND_ADDR", "127.0.0.1")
	viper.SetDefault("BIND_PORT", "3000")
	viper.SetDefault("ADDRESS_IPV4", "")
	viper.SetDefault("ADDRESS_IPV6", "")
	viper.SetDefault("BEHIND_PROXY", "false")
}

func ConfigValue(key string) string {
	return viper.GetString(key)
}
