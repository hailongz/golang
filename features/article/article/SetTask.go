package article

type SetTask struct {
	Id	int64	`json:"id" name:"id" title:"动态ID"`
	State	interface{}	`json:"state,omitempty" name:"state" title:"状态 多个逗号分割"`
	Body	interface{}	`json:"body,omitempty" name:"body" title:"内容"`
	Options	interface{}	`json:"options,omitempty" name:"options" title:"其他数据 JSON 叠加数据"`
	Ctime	interface{}	`json:"ctime,omitempty" name:"ctime" title:"发布时间"`
}

func (T *SetTask) GetName() string {
	return "set.json"
}

func (T *SetTask) GetTitle() string {
	return "修改动态"
}

