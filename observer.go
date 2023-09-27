package main

import "fmt"

type Student struct {
	Name   string
	Grades []int
	Parent ParentObserver
}

type ParentObserver interface {
	Update(studentName string, newGrade int)
	Subscribe(studentName string)
	Unsubscribe()
}

type Parent struct {
	Name         string
	SubscribedTo string // The name of the child they are subscribed to
}

func NewParent(name string) *Parent {
	return &Parent{Name: name}
}

func (p *Parent) Update(studentName string, newGrade int) {
	fmt.Printf("Parent %s received a notification: %s got a new grade: %d\n", p.Name, studentName, newGrade)
}

func (p *Parent) Subscribe(studentName string) {
	p.SubscribedTo = studentName
}

func (p *Parent) Unsubscribe() {
	p.SubscribedTo = ""
}

type Diary struct {
	students  map[string]*Student
	observers map[string]ParentObserver
}

func NewDiary() *Diary {
	return &Diary{
		students:  make(map[string]*Student),
		observers: make(map[string]ParentObserver),
	}
}

func (d *Diary) AddStudent(s *Student) {
	d.students[s.Name] = s
}

func (d *Diary) AddParent(observer ParentObserver, studentName string) {
	d.observers[studentName] = observer
}

func (d *Diary) RemoveParent(studentName string) {
	delete(d.observers, studentName)
}

func (d *Diary) AddGrade(studentName string, newGrade int) {
	student, exists := d.students[studentName]
	if exists {
		student.Grades = append(student.Grades, newGrade)
		if observer, ok := d.observers[studentName]; ok {
			observer.Update(studentName, newGrade)
		}

	}
}

func main() {

	diary := NewDiary()

	alice := &Student{Name: "Alice", Grades: []int{95, 88}}
	bob := &Student{Name: "Bob", Grades: []int{78, 90}}
	diary.AddStudent(alice)
	diary.AddStudent(bob)

	parent1 := NewParent("Parent 1")
	parent2 := NewParent("Parent 2")

	parent1.Subscribe("Alice")
	parent2.Subscribe("Bob")

	diary.AddParent(parent1, "Alice")
	diary.AddParent(parent2, "Bob")

	diary.AddGrade("Alice", 92)

	diary.AddGrade("Bob", 85)

}
