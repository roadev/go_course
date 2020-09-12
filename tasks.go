package main

import "fmt"

type task struct {
	name        string
	description string
	completed   bool
}

type taskList struc {
	tasks [] *task
}

func (t *taskList) addToList(ta *task) {
	t.tasks = append(t.tasks, ta)
}

func (t *task) markCompleted() {
	t.completed = true
}

func (t *task) updateName(name string) {
	t.name = name
}

func (t *task) updateDescription(description string) {
	t.description = description
}

func main() {
	t1 := &task{
		name:        "Complete my Go course",
		description: "Complete my Go course in Platzi on this week",
	}

	t2 := &task{
		name:        "Complete my Python course",
		description: "Complete my Python course in Platzi on this week",
	}
	
	list := &taskList {
		tasks: [] *task {
			t1, t2
		}
	}
}
