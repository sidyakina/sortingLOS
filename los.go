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

func (los *LOS) getLasts() (penultimate, last *LOS) {
	penultimate = los
	last = penultimate.next
	if last == nil {
		// переданный элемент и так последний
		return
	}
	for last.next != nil {
		penultimate = last
		last = penultimate.next
	}
	//fmt.Printf("/n %v %v", penultimate.number, last.number)
	return
}

func (los *LOS) getLength() int {
	length := 0
	ptr := los
	for ptr != nil {
		length++
		ptr = ptr.next
	}
	return length
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
	penultimate, last := oldNext.getLasts()
	if last == nil {
		// мы в конце списка, все отсортировано
		return
	}
	// ставим последний элемент между текущим и старым следующим
	// после этого предпоследний становится последним элементом
	ptr.next = last
	penultimate.next = nil
	last.next = oldNext
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
		oldNext := changeList.next
		penultimate, last := insertList.getLasts()
		if last == nil {
			changeList.next = penultimate
			penultimate.next = oldNext
			return
		}
		changeList.next = last
		penultimate.next = nil
		last.next = oldNext
		changeList = oldNext
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
		oldNext := changeList.next
		if tInsertList.next == nil {
			changeList.next = tInsertList.element
			tInsertList.element.next = oldNext
			return
		}
		changeList.next = tInsertList.element
		tInsertList.element.next = oldNext
		tInsertList = tInsertList.next
		changeList = oldNext
	}
}

type transformedLOS struct {
	element *LOS
	next    *transformedLOS
}

func transformLOS(los LOS) *transformedLOS {
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
