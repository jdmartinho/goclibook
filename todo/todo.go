package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type ItemList []item

func (l *ItemList) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, t)
}

func (l *ItemList) Complete(i int) error {
	ls := *l
	if i < 0 || i > len(ls) {
		return fmt.Errorf("item %d doesn't exist", i)
	}

	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

func (l *ItemList) Delete(i int) error {
	ls := *l
	if i < 0 || i > len(ls) {
		return fmt.Errorf("item %d doesn't exist", i)
	}

	*l = append(ls[:i-1], ls[i:]...)
	return nil
}

func (l *ItemList) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, js, 0644)
}

func (l *ItemList) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l)
}
