package lib

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Todo struct {
	createdAt time.Time
	title string
	content string
	completed bool
}

func (t *Todo) Title() string {
	return t.title
}

func (t *Todo) Content() string {
	return t.content
}

func (t *Todo) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Todo) Completed() bool {
	return t.completed
}

func (t *Todo) setTitle(title string) {
	t.title = title
}

func (t *Todo) setContent(content string) {
	t.content = content
}

func NewTodo(title string, content string) *Todo {
	var todo Todo
	todo.title = title
	todo.content = content
	todo.completed = false
	todo.createdAt = time.Now()
	fmt.Println(todo.title, todo.content)
	return &todo
}

func (t Todo) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Title string
		Content string
		CreatedAt time.Time
		Completed bool
	}{
		Title: t.title,
		Content: t.content,
		CreatedAt: t.createdAt,
		Completed: t.completed,
	})
}

func (t *Todo) UnmarshalJSON(data []byte) (error) {
	aux := &struct {
		Title string
		Content string
		CreatedAt time.Time
		Completed bool
	}{}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	t.title = aux.Title
	t.content = aux.Content
	t.completed = aux.Completed
	t.createdAt = aux.CreatedAt

	return nil
}

func LoadTodosJSON(file *os.File) ([]*Todo, error) {
	reader := bufio.NewReader(file)

	var todos []*Todo
	err := json.NewDecoder(reader).Decode(&todos)

	return todos, err
}

func StoreTodosJSON(file *os.File, todos *[]*Todo) (error) {
	writer := bufio.NewWriter(file)

	file.Truncate(0)
	file.Seek(0, 0)
	err := json.NewEncoder(writer).Encode(todos)
	writer.Flush()

	return err
}
