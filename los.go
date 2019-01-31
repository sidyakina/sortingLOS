package main

import "fmt"

type LOS struct {
	number int
	next   *LOS
}

func (los LOS) String() string {
	result := fmt.Sprintf("%v", los.number)
	ptr := los
	for {
		if ptr.next != nil {
			ptr = *(ptr.next)
			result = result + fmt.Sprintf(", %v", ptr.number)
		} else {
			return result
		}

	}
}

func newLOS(number int) *LOS {
	return &LOS{number: number}
}

func NewList(maxNumber int) *LOS {
	los := newLOS(1)
	ptr := los
	for i := 2; i < maxNumber + 1; i++ {
		ptr.next = newLOS(i)
		ptr = ptr.next
	}
	return los
}


func (los *LOS) Sort() {
	los.swap()
}


func (los *LOS) swap() {
	ptr := los             // текущий указатель на изменяемый элемент
	oldNext := ptr.next // старый следующий элемент изменяемого элемента
	ptr0, ptr1 := getLasts(oldNext)
	if ptr1 == nil {
		// мы в конце списка, все отсортировано
		return
	}
	// ставим следующим после текущего последний и делаем последним предыдущий последний
	ptr.next = ptr1
	ptr0.next = nil
	ptr1.next = oldNext
	oldNext.swap()
}

func getLasts(los *LOS) (*LOS, *LOS) {
	if los.next == nil {
		// элемент и так последний
		return los, nil
	}
	ptr0 := los
	ptr1 := los.next
	for ptr1.next != nil {
		ptr0 = ptr1
		ptr1 = ptr0.next
	}
	//fmt.Printf("/n %v %v", ptr0.number, ptr1.number)
	return ptr0, ptr1
}
