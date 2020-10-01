package main

import "fmt"

type task struct {
	name        string
	description string
	completed   bool
}

type taskList struct {
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

func (t *taskList) deleteTask(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index + 1:]...)
}

func (t *taskList) printList() {
	for _, task := range t.tasks {
		fmt.Println("Name", task.name)
		fmt.Println("Description", task.description)
	}
}

func (t *taskList) completedList() {
	for _, task := range t.tasks {
		if task.completed {
			fmt.Println("Name", task.name)
			fmt.Println("Description", task.description)
		}
	}
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

	t3 := &task{
		name:        "Complete my Node course",
		description: "Complete my Node course in Platzi on this week",
	}
	
	list := &taskList {
		tasks: [] *task {
			t1, t2,
		},
	}

	list.addToList(t3)

	// list.printList()

	list.tasks[0].markCompleted()

	// fmt.Println("Completed tasks")

	list.completedList()

	tasksMap := make(map[string]*taskList)

	tasksMap["Juan"] = list

	t4 := &task{
		name:        "Complete my Java course",
		description: "Complete my Java course in Platzi on this week",
	}

	t5 := &task{
		name:        "Complete my C# course",
		description: "Complete my C# course in Platzi on this week",
	}

	list2 := &taskList {
		tasks: [] *task {
			t4, t5,
		},
	}

	tasksMap["David"] = list2
	
	fmt.Println("Juan's tasks")
	tasksMap["Juan"].printList()

	fmt.Println("David's tasks")
	tasksMap["David"].printList()

	// for i := 0; i < len(list.tasks); i++ {
	// 	fmt.Println("Index", i, "Name", list.tasks[i].name)
	// }

	// for index, task := range list.tasks {
	// 	fmt.Println("Index", index, "Task", task.name)
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }
}
