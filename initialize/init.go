package initialize

import (
	"github.com/6.2/global"
	"github.com/winksai/utils/mysql"
)

func Mysql() {
	db := mysql.InitMysql("root", "root", "127.0.0.1", "day03", 3306)
	global.DB = db
}
