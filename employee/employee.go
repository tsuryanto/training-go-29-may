package employee

import "fmt"

type Employee struct {
	EmpNo int
	Name  string
}

func (e Employee) GetName() string {
	e.EmpNo = e.setEmpNo()
	return e.Name
}

func (e Employee) GetEmpNo() int { // public
	return e.EmpNo
}

func (e Employee) setEmpNo() int { // private
	return e.EmpNo
}

func (e Employee) CountAge() int {
	// tanggal lahir - hari ini
	return e.EmpNo
}

func (e Employee) AksesMasuk() {
	fmt.Println("akses masuk")
}

type Pakaian interface {
	Pakai()
}

func (e Employee) PakaiPakaian(pakaian Pakaian) {
	pakaian.Pakai()
}
