Let's begin our code jouney of finite field, go to an empty directory, create a new folder named bitcoin, in the bitcoind dir, create a new 
folder named elliptic-curve, let's run the command to init a new package there:
```go
go init mod elliptic_curve
```

then we create a new file finite-element, we will have our code for  element of the finite field here, let's put these code into the file:
```go
package elliptic_curve

import (
	"fmt"
)

type FieldElement struct {
	order uint64 //field order
	num   uint64 //value of the given element in the field
}

func NewFieldElement(order uint64, num uint64) FieldElement {
	/*
		constructor for FieldElement, its the __init__ if you are from python
	*/

	if num >= order || num < 0 {
		err := fmt.Sprintf("Num not in the range from 0 to %d", order)
		panic(err)
	}

	return FieldElement{
		order: order,
		num:   num,
	}
}

func (f FieldElement) String() string {
	//format the object to printable string
	//its __repr__ if you are from python
	return fmt.Sprintf("FieldElement{order: %d, num: %d}", f.order, f.num)
}

func (f FieldElement) EqualTo(other FieldElement) bool {
	/*
		two field element is equal if their order and value are equal
	*/
	return f.order == other.order && f.num == other.num
}

```

now we have the bare born for finite field element, let's add more method on it.we have two kind of operation on the field element, one is "+", it is arithmetic add with modulur, and "." it is arithmetic multiplie with modulur, let's see how we can do the addition:
```go
func (f *FieldElement) Add(other *FieldElement) *FieldElement {
	if other.order != f.order {
		panic("add need to do on field element with the same order")
	}

	//remember to do the modulur
	return NewFieldElement(f.order, (f.num+other.num)%f.order)
}

func (f *FieldElement) Negate() *FieldElement {
	/*
		for a field element a, its negate is another element b in field such that
		(a + b) % order= 0(remember the modulur over order), because the value of element
		in the field are smaller than its order, we can easily get the negate of a by
		order - a,
	*/

	return NewFieldElement(f.order, f.order-f.num)
}

func (f *FieldElement) Substract(other *FieldElement) *FieldElement {
	//How to do ?
	return nil
}

```

let's call the above code to check for result, in main.go have the code like following:
```go
package main

import (
	ecc "elliptic_curve"
	"fmt"
)

func main() {
	/*
		construct field with order 57 and do add and substract
	*/
	f44 := ecc.NewFieldElement(57, 44)
	f33 := ecc.NewFieldElement(57, 33)
	// 44 + 33 equal to (44+33) % 57 is 20
	res := f44.Add(f33)
	fmt.Printf("field element 44 add to field element 33 is : %v\n", res)
	//-44 is the negate of field element 44, which is 57 - 44 = 13
	fmt.Printf("negate of field element 44 is : %v\n", res.Negate())
}
```
then run the command below:
```go
go run main.go
```
if everything go smoothly, you will see the following result:
```go
field element 44 add to field element 33 is : FieldElement{order: 57, num: 20}
negate of field element 44 is : FieldElement{order: 57, num: 37}
```
