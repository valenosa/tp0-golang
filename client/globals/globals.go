package globals

type Config struct {
	Ip      string `json:"ip"`
	Puerto  int    `json:"puerto"`
	Mensaje string `json:"mensaje"`
}

var ClientConfig *Config
