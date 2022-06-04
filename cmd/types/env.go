package types

// User ...
type User struct {
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	Email    string `yaml:"Email"`
}

// Persistence ...
type Persistence struct {
	StorageClass       string `yaml:"StorageClass"`
	SizeRedis          string `yaml:"SizeRedis"`
	SizeDockerRegistry string `yaml:"SizeDockerRegistry"`
}

// Config ...
type Config struct {
	Namespace   string      `yaml:"Namespace"`
	User        User        `yaml:"User"`
	Persistence Persistence `yaml:"Persistence"`
}
