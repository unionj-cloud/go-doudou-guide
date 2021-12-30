package vo

//go:generate go-doudou name --file $GOFILE -o

type PageFilter struct {
	// 真实姓名，前缀匹配
	Name string `json:"name,omitempty"`
	// 所属部门ID
	Dept int `json:"dept,omitempty"`
}

type Order struct {
	Col  string `json:"col,omitempty"`
	Sort string `json:"sort,omitempty"`
}

type Page struct {
	// 排序规则
	Orders []Order `json:"orders,omitempty"`
	// 页码
	PageNo int `json:"pageNo,omitempty"`
	// 每页行数
	Size int `json:"size,omitempty"`
}

// 分页筛选条件
type PageQuery struct {
	Filter PageFilter `json:"filter,omitempty"`
	Page   Page       `json:"page,omitempty"`
}

type PageRet struct {
	Items    interface{} `json:"items,omitempty"`
	PageNo   int         `json:"pageNo,omitempty"`
	PageSize int         `json:"pageSize,omitempty"`
	Total    int         `json:"total,omitempty"`
	HasNext  bool        `json:"hasNext,omitempty"`
}

type UserVo struct {
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
	Dept  string `json:"dept,omitempty"`
}
