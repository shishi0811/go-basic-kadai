package main

import "fmt"

type ToDoItem struct {
	Id        int
	Title     string
	Completed bool
}

func (item *ToDoItem) Complete() {
	fmt.Printf("ID：%dのToDoを完了に更新します\n", item.Id)
	item.Completed = true
}

func printTodos(todos []ToDoItem) {
	todoState := "完了"
	for _, v := range todos {
		if v.Completed == false {
			todoState = "未完了"
		}
		fmt.Printf("[%s](ID:%d)%s\n", todoState, v.Id, v.Title)
	}
}

func main() {
	// []ToDo型のtodosスライスを作成
	todos := make([]ToDoItem, 0, 3)

	// 3つのToDoアイテムを以下の初期値で作り、todosスライスに追加
	todo1 := ToDoItem{
		Id:        1,
		Title:     "学習計画",
		Completed: false,
	}
	todo2 := ToDoItem{
		Id:        2,
		Title:     "環境構築",
		Completed: false,
	}
	todo3 := ToDoItem{
		Id:        3,
		Title:     "基礎文法",
		Completed: false,
	}
	todos = append(todos, todo1, todo2, todo3)

	// printTodos()関数にtodosスライスを渡し、初期状態のToDoリストを表示
	printTodos(todos)

	// IDが1・2のToDoアイテムをComplete()メソッドで完了状態に更新
	todo1.Complete()
	todo2.Complete()

	// printTodos()関数にtodosスライスを渡し、最終的なToDoリストを表示
	printTodos(todos)

	// ※簡潔に書いてみた
	// todos := []ToDoItem{
	// 	{
	// 		Id:        1,
	// 		Title:     "学習計画",
	// 		Completed: false,
	// 	},
	// 	{
	// 		Id:        2,
	// 		Title:     "環境構築",
	// 		Completed: false,
	// 	},
	// 	{
	// 		Id:        3,
	// 		Title:     "基礎文法",
	// 		Completed: false,
	// 	},
	// }

	// printTodos(todos)

	// todos[0].Complete()
	// todos[1].Complete()

	// printTodos(todos)
}
