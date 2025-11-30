package config

import (
	"backend/utils/logs"

	"github.com/spf13/viper"
)

var (
	JwtSignKey string //jwt签名密钥
	JwtExpTime int64  //jwt过期时间 单位：分钟
	Port       string //服务端口号
	Username   string //用户名
	Password   string //密码
	UploadsDir string //uploads目录路径
)

type ReturnDate struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// 构造函数
/*总是初始化完整状态，这种写法适合返回给前端，有默认值。
什么都不传的话。返回的json如下：
{
  "status": 200,
  "message": "",
  "data": {}
}
*/
func NewReturnDate() ReturnDate {
	returnDate := ReturnDate{}
	returnDate.Status = 200
	Data := make(map[string]interface{})
	returnDate.Data = Data
	return returnDate
}

func init() {
	//获取程序启动端口的配置
	viper.SetDefault("Port", "8080")
	//配置jwt加密的secret，生成JWT令牌，可以使用oliver作为密钥加密，生产环境改成随机字符串
	viper.SetDefault("JWT_SIGN_KEY", "ZtIace633OWYtpH0Y7Z6IFb3rGstbzzw") //
	//配置jwt过期时间配置
	viper.SetDefault("JWT_EXP_TIME", 120)
	//配置用户名和密码，默认值：oliver，oliver6
	viper.SetDefault("USERNAME", "admin")
	viper.SetDefault("PASSWORD", "123456")
	//配置uploads目录路径(注意！根据生产环境配置)
	viper.SetDefault("UPLOADS_DIR", "C:\\Users\\Xyh\\Desktop\\Ad-go\\backend\\uploads")
	logs.Debug(nil, "开始加载配置")

	//获取用户名和密码
	Username = viper.GetString("USERNAME")
	Password = viper.GetString("PASSWORD")
	//获取jwt签名密钥
	JwtSignKey = viper.GetString("JWT_SIGN_KEY")
	//获取jwt过期时间
	JwtExpTime = viper.GetInt64("JWT_EXP_TIME")
	//获取端口号
	Port = viper.GetString("PORT")
	//获取uploads目录路径
	UploadsDir = viper.GetString("UPLOADS_DIR")

	//调试可以打印配置信息
	logs.Debug(map[string]interface{}{"用户名": Username, "密码": Password}, "获取到程序用户名密码配置")

}
