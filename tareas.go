package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) addToList(tl *task) {
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) deleteTaskFromList(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *taskList) printList() {
	fmt.Println(":::: Lista de Tareas ::::")
	for _, task := range t.tasks {
		fmt.Println("Name: ", task.name)
		fmt.Println("Description: ", task.name)
	}
}

func (t *taskList) printListOfCompleteTasks() {
	fmt.Println(":::: Lista de Tareas Completadas ::::")
	for _, task := range t.tasks {
		if task.finished {
			fmt.Println("Name: ", task.name)
			fmt.Println("Description: ", task.name)
		}
	}
}

type task struct {
	name        string
	description string
	finished    bool
}

func (t *task) markAsDone() {
	t.finished = true
}

func (t *task) updateDescription(description string) {
	t.description = description
}

func (t *task) updateName(name string) {
	t.name = name
}

func main() {
	t1 := &task{
		name:        "Completar mi curso de Go",
		description: "Completar mi curso de Go de platzi en esta semana",
		finished:    false,
	}

	t2 := &task{
		name:        "Completar mi curso de Node.js",
		description: "Completar mi curso de Node.js de platzi en esta semana",
		finished:    false,
	}

	t3 := &task{
		name:        "Completar mi curso de React.js",
		description: "Completar mi curso de React.js de platzi en esta semana",
		finished:    false,
	}

	list := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}

	list.tasks[0].markAsDone()

	list.addToList(t3)
	list.printList()
	list.printListOfCompleteTasks()
	// fmt.Println(len(list.tasks))

	/* list.deleteTaskFromList(1)
	fmt.Println(len(list.tasks)) */

	/* for i := 0; i < len(list.tasks); i++ {
		fmt.Println("Index: ", i, "Nombre: ", list.tasks[i].name)
	} */

	/* for index, task := range list.tasks {
		fmt.Println("Index: ", index, "Nombre: ", task.name)
	} */

	tasksMap := make(map[string]*taskList)

	tasksMap["Maria"] = list

	t4 := &task{
		name:        "Completar mi curso de C#",
		description: "Completar mi curso de C# de platzi en esta semana",
		finished:    false,
	}

	t5 := &task{
		name:        "Completar mi curso de Java",
		description: "Completar mi curso de Java de platzi en esta semana",
		finished:    false,
	}

	list2 := &taskList{
		tasks: []*task{
			t4, t5,
		},
	}

	tasksMap["Ricardo"] = list2

	fmt.Println("Tareas de Maria")
	tasksMap["Maria"].printList()

	fmt.Println("Tareas de Ricardo")
	tasksMap["Ricardo"].printList()

}
