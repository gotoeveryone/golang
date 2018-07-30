package config

type (
	// Config 設定
	Config struct {
		Log   Log   `json:"log"`
		Cache Cache `json:"cache"`
		DB    DB    `json:"db"`
		Mail  Mail  `json:"mail"`
	}

	// Log ログ設定
	Log struct {
		Prefix string   `json:"prefix"`
		Path   string   `json:"path"`
		Level  LogLevel `json:"level"`
		Type   string   `json:"type"`
	}

	// Cache キャッシュサーバ接続設定
	Cache struct {
		Use  bool   `json:"use"`
		Host string `json:"host"`
		Port int    `json:"port"`
		Auth string `json:"auth"`
	}

	// DB データベース接続設定
	DB struct {
		Name     string `json:"name"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Timezone string `json:"timezone"`
	}

	// Mail メール接続設定
	Mail struct {
		SMTP      string   `json:"smtp"`
		Port      int      `json:"port"`
		UseTLS    bool     `json:"useTLS"`
		User      string   `json:"user"`
		Password  string   `json:"password"`
		From      string   `json:"from"`
		FromAlias string   `json:"fromAlias"`
		To        []string `json:"to"`
	}

	// LogLevel is log level text
	LogLevel string
)
