package constants

const (
	UserTableName            = "user"
	UserServiceName          = "user"
	VideoTableName           = "video"
	VideoServiceName         = "video"
	MySQLDefaultDSN          = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress              = "127.0.0.1:2379"
	WSLAdress                = "http://172.28.131.59:8080/static/"
	OssAdress                = "https://douyin-1259511784.cos.ap-chongqing.myqcloud.com"
	OssSecretID              = "your secret id"
	OssSecretKey             = "your secret key"
	ApiServiceName           = "demoapi"
	SecretKey                = "secret key"
	IdentityKey              = "id"
	CPURateLimit     float64 = 80.0
)
