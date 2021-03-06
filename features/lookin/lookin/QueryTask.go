package lookin

type QueryTask struct {
	Tid	int64	`json:"tid" name:"tid" title:"目标"`
	Iid	interface{}	`json:"iid,omitempty" name:"iid" title:"项ID 默认 0"`
	Uid	interface{}	`json:"uid,omitempty" name:"uid" title:"用户ID"`
	Fuid	interface{}	`json:"fuid,omitempty" name:"fuid" title:"用户ID"`
	Flevel	interface{}	`json:"flevel,omitempty" name:"flevel" title:"好友级别，多个逗号分割"`
	GroupBy	interface{}	`json:"groupBy,omitempty" name:"groupby" title:"分组"`
	P	interface{}	`json:"p,omitempty" name:"p" title:"分页位置, 从1开始, 0 不处理分页"`
	N	interface{}	`json:"n,omitempty" name:"n" title:"分页大小，默认 20"`
}

func (T *QueryTask) GetName() string {
	return "query.json"
}

func (T *QueryTask) GetTitle() string {
	return "查询"
}

