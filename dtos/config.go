package dtos

type Config struct {
	DB    DBConfig    `mapstructure:"db"`
	Admin AdminConfig `mapstructure:"adm"`
	JWT   JWTConfig   `mapstructure:"jwt"`
}

type DBConfig struct {
	Host string `mapstructure:"host"`
	User string `mapstructure:"user"`
	Pass string `mapstructure:"pwd"`
	Name string `mapstructure:"name"`
	Port int    `mapstructure:"port"`
}
type JWTConfig struct {
	Secret string `mapstructure:"secret"`
}
type AdminConfig struct {
	Password string `mapstructure:"password"`
	Email    string `mapstructure:"email"`
}
