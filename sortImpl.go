package sortImpl

import (
	"fmt"
	"sort" // important for our case
)

type Employee struct { 
	Name string 
	Surname string
	IdentificationNumber int 
	PhoneNumber string 
	Department string
}

type SortType []Employee

//Implement sort.Sort(i Interface) methods.
func (arr SortType) Len() int {
	return len(arr)
}

// Sort function Employee.IdentificationNumber.
func (arr SortType) Less(x, y int) bool {
	return arr[x].IdentificationNumber < arr[y].IdentificationNumber
}

func (arr SortType) Swap(x, y int) {
	arr[x], arr[y] = arr[y], arr[x] //Golang feature don't need to use temp variable.
}

// Format custom type
func (e Employee) String() string {
	return fmt.Sprintf("%s %s %d %s %s \n", e.Name, e.Surname, e.IdentificationNumber, e.Department, e.PhoneNumber)				
}

// Only if first letter is upper case then it can be use export
func CustomSort(arr *[]Employee) {
	sort.Sort(SortType(*arr))	
	for _, e := range *arr {
		fmt.Println(e)
	}
}