// 本地服务配置文件读取，优先级低于环境变量

package config

import "bee-boilerplate/public/logger"

var s_logger = logger.New(logger.LocalConfig)
