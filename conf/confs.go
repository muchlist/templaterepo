package conf

type UserConfig struct {
	ServerPort    string
	SecretKey     string
	DebugPort     string
	DbDsn         string
	DbMaxOpenConn int
	DbMinOpenConn int
}

type ComplexConfig struct {
	ServerPort string
	SecretKey  string
}
