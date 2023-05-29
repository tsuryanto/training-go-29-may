package employee

type Manager struct {
	Title string
}

type Programmer struct {
	Employee
	Manager
	CodeEditor string
}

func NewProgrammer(name string, codeEditor string) Programmer {
	return Programmer{
		Employee: Employee{
			EmpNo: 1,
			Name:  name,
		},
		Manager: Manager{
			Title: "Manager Development",
		},
		CodeEditor: codeEditor,
	}
}
