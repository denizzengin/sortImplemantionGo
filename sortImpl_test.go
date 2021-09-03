package sortImpl

import (
	"testing"	
)

func TestCustomSort(t *testing.T){

	employees := []Employee{ 
		{Name:"deniz", Surname: "zengin", IdentificationNumber: 28, Department: "Software", PhoneNumber:"111-11-11"},
		{Name:"steven", Surname: "zengin", IdentificationNumber: 27, Department: "System", PhoneNumber:"111-11-12"},
		{Name:"kamil", Surname: "zengin", IdentificationNumber: 30, Department: "Finance", PhoneNumber:"111-11-13"},
	 } 
	 CustomSort(&employees)
	 if len(employees) == 0 {
		t.Errorf("not passed")
	 }

	 for _, e := range employees {		
		t.Logf("%s %s %d %s %s \n", e.Name, e.Surname, e.IdentificationNumber, e.Department, e.PhoneNumber)
	}

}