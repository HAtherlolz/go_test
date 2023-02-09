package todo

type TodoList struct {
	Id int `json:"-"`
	Title string `json:"title"`
	Description string `json:"description"`
}

type UserList struct {
	Id int
	UserId string
	Description string
}

type TodoItem struct {
	Id int `json:"-"`
	Title string `json:"title"`
	Description string `json:"description"`
	Done bool `json:"done"`
}

type TodoItem struct {
	Id int
	ListId int
	ItemId int
}