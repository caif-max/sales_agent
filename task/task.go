// task模块负责从外界获取任务，并转化为AI任务，提交给agent处理
// todo 可能还要抽象出来一层，专门处理业务逻辑
package task

import (
	"sales_agent/common/log"
	"sales_agent/task/event"
	"sales_agent/task/message"
)

func Init() {
	log.GetLogger().Info("Init task")
	// todo 初始化定时器，每天凌晨从robot server中拉取开通了服务的账户和员工（或者从计费中）

	// todo 遍历所有账户，从scrm中获取员工对应的任务，并初始化task列表

	// todo 初始化event，根据task列表监听scrm中的员工任务变更
	event.Init()
	// todo 初始化message，根据task列表监听scrm中的员工的会话存档消息
	message.Init()

	// todo 初始化第二个定时器，每天8-9点，根据销售机会任务来生成AI任务，提交给agent处理

	// todo 初始化第三个定时器，每天凌晨0点，清空任务列表
}
