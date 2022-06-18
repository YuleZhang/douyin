package constants

const (
	UserTableName            = "user"
	UserServiceName          = "user"
	VideoTableName           = "video"
	VideoServiceName         = "video"
	MySQLDefaultDSN          = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress              = "127.0.0.1:2379"
	ApiServiceName           = "demoapi"
	SecretKey                = "secret key"
	IdentityKey              = "id"
	CPURateLimit     float64 = 80.0
)
