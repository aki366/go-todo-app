package models

// TodoItemは、APIのTodoアイテムの構造を定義
type TodoItem struct {
    ID          string `json:"id"`
    Description string `json:"description"`
    Completed   bool   `json:"completed"`
}
