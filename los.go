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
		if ptr.next == nil {
			return result
		}
		ptr = *(ptr.next)
		result = result + fmt.Sprintf(", %v", ptr.number)

	}
}

func newElement(number int) *LOS {
	return &LOS{number: number}
}

func NewLOS(maxNumber int) *LOS {
	if maxNumber < 1 {
		return nil
	}
	los := newElement(1)
	ptr := los
	for i := 2; i < maxNumber+1; i++ {
		ptr.next = newElement(i)
		ptr = ptr.next
	}
	return los
}

func (los *LOS) getLasts() (*LOS, *LOS) {
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

func (los *LOS) getLength() int {
	l := 0
	ptr := los
	for ptr != nil {
		l++
		ptr = ptr.next
	}
	return l
}

func (los *LOS) cutSecondPart(length int) *LOS {
	ptr := los
	if length == 1 {
		return nil
	}
	if length%2 == 1 {
		length += 1
	}
	separator := int(float32(length) / 2)
	for i := 1; i < separator; i++ {
		ptr = ptr.next
	}
	newPtr := ptr.next
	ptr.next = nil
	return newPtr
}

//----------------------------------------------------

func (los *LOS) Sort1() {
	los.swap1()
}

func (los *LOS) swap1() {
	ptr := los          // текущий указатель на изменяемый элемент
	oldNext := ptr.next // старый следующий элемент изменяемого элемента
	if oldNext == nil {
		// нет следующего элемента, может быть если список состоит из одного элемента
		return
	}
	ptr0, ptr1 := oldNext.getLasts()
	if ptr1 == nil {
		// мы в конце списка, все отсортировано
		return
	}
	// ставим следующим после текущего последний и делаем последним предыдущий последний
	ptr.next = ptr1
	ptr0.next = nil
	ptr1.next = oldNext
	oldNext.swap1()
}

//----------------------------------------------------

func (los *LOS) Sort2() {
	length := los.getLength()
	los.swap2(length)
}

func (los *LOS) swap2(length int) {
	changeList := los
	insertList := changeList.cutSecondPart(length)
	if insertList == nil {
		return
	}
	for {
		old := changeList.next
		newLast, last := insertList.getLasts()
		if last == nil {
			changeList.next = newLast
			newLast.next = old
			return
		}
		changeList.next = last
		newLast.next = nil
		last.next = old
		changeList = old
	}
}

//----------------------------------------------------

func (los *LOS) Sort3() {
	length := los.getLength()
	los.swap3(length)
}

func (los *LOS) swap3(length int) {
	changeList := los
	insertList := changeList.cutSecondPart(length)
	if insertList == nil {
		return
	}
	tInsertList := transformLOS(*insertList)
	for {
		old := changeList.next
		if tInsertList.next == nil {
			changeList.next = tInsertList.element
			tInsertList.element.next = old
			return
		}
		changeList.next = tInsertList.element
		tInsertList.element.next = old
		tInsertList = tInsertList.next
		changeList = old
	}
}

type transformedLOS struct {
	element *LOS
	next *transformedLOS
}

func transformLOS(los LOS) *transformedLOS{
	ptrT := &transformedLOS{&los, nil}
	ptrL := &los
	for ptrL.next != nil {
		ptrL = ptrL.next
		ptrT = &transformedLOS{ptrL, ptrT}
	}
	return ptrT
}

//----------------------------------------------------

func (los *LOS) Sort4() {
	length := los.getLength()
	if length < 101 {
		los.swap2(length)
	} else {
		los.swap3(length)
	}
}