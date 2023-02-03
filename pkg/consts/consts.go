package consts

const (
	UserInfoTableName = "user_info"
	UserTableName     = "tiktok_user"
	FollowTableName   = "follow_relation"
	SecretKey         = "secret key"
	IdentityKey       = "user_id"
	Total             = "total"
	ApiServiceName    = "api"
	UserServiceName   = "user"
	MySQLDefaultDSN   = "gorm:gorm@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	TCP               = "tcp"
	UserServiceAddr   = ":9000"
	ExportEndpoint    = ":4317"
	ETCDAddress       = "127.0.0.1:2379"
)
