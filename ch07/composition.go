package main

import "fmt"

// 下位の構造体
type Employee struct {
	name string
	id string
}
func (e Employee) Description() string {
	return fmt.Sprintf("%s(%s)", e.name, e.id)
}

// 上位の構造体
type Manager struct {
	Employee //型のみを書く。型のフィールドが加わる。
	Reports []Employee
}
func (m Manager) addEmpoyees() []Employee {
	newEmployees := []Employee{
		{"太郎","100"},
		{"治郎","200"},
	}
	return newEmployees
}

func main() {
	m := Manager{
		Employee: Employee{"花子", "000"},
		Reports: []Employee{},
	}
	fmt.Println(m.id) // 000 
	fmt.Println(m.Description()) //000(花子)

	m.Reports = m.addEmpoyees()
	fmt.Println(m.Employee) // {花子 000}
	fmt.Println(m.Reports)	// [{太郎 100} {治郎 200}]	
}