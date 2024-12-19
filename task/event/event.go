// event模块负责监听账户下员工的任务变更
package event

import (
	"sales_agent/common/log"
)

func Init() {
	log.GetLogger().Info("Init event")
}
