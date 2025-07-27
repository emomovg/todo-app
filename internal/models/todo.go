package models

const (
	TodoListTable  = "todo_lists"
	UserListTable  = "users_lists"
	TodoItemTable  = "todo_items"
	ListsItemTable = "lists_items"
)

type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description""`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        string `json:"done"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}
