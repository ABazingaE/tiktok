package consts

const (
	UserInfoTableName = "user_info"
	UserTableName     = "tiktok_user"
	FollowTableName   = "follow_relation"
	SecretKey         = "secret key"
	IdentityKey       = "user_id"
	ApiServiceName    = "api"
	UserServiceName   = "user"
	VideoServiceName  = "video"
	MySQLDefaultDSN   = "gorm:gorm@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	TCP               = "tcp"
	UserServiceAddr   = ":9000"
	VideoServiceAddr  = ":9001"
	ExportEndpoint    = ":4317"
	ETCDAddress       = "127.0.0.1:2379"
	AccessKeyId       = "LTAI5tDqxkmzXp1YkSb4tYXn"
	AccessKeySecret   = "fxPiIbyE8nqP4i5EzTUQdqiXqjcHGQ"
	TempPath          = "/home/bazinga/go/tiktok/"
	LikeServiceName   = "like"
	LikeServiceAddr   = ":9002"
)
