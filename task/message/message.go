// message模块负责监听scrm中的员工的会话存档消息
package message

import (
	"sales_agent/common/log"
)

func Init() {
	log.GetLogger().Info("Init message")
}
