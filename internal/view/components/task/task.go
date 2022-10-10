package task

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Task struct {
	Name        binding.String
	Description binding.String
	Complete    binding.Bool
	Container   *fyne.Container
}

func CreateTask(title, description string) *Task {
	newTask := new(Task)
	fmt.Println(newTask)
	newTask.Complete = binding.NewBool()
	newTask.Complete.Set(false)
	newTask.Name = binding.NewString()
	newTask.Name.Set(title)
	newTask.Description = binding.NewString()
	newTask.Description.Set(description)
	element := container.New(
		layout.NewHBoxLayout(),
		widget.NewLabelWithData(newTask.Name),
	)
	newTask.Container = element
	return newTask

}
