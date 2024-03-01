In bitcoin cryptography, we need to deal with Astronomical figures, the number is so huge that it is easily outpace the total count to atoms in our universe, and maybe 64bits value is 
not enough to represent the figure, therefore we may need the help of big package from golang, we will change our code of FieldNumber by using big.Int to represent its value field,
the code will change like the following:
```go
package elliptic_curve

import (
	"fmt"
	"math/big"
)

//using big package to deal with Astronomical figures

type FieldElement struct {
	order *big.Int //field order
	num   *big.Int //value of the given element in the field
}

func NewFieldElement(order *big.Int, num *big.Int) *FieldElement {
	/*
		constructor for FieldElement, its the __init__ if you are from python
	*/
	if order.Cmp(num) == -1 {
		err := fmt.Sprintf("Num not in the range from 0 to %v", order)
		panic(err)
	}

	return &FieldElement{
		order: order,
		num:   num,
	}
}

func (f *FieldElement) String() string {
	//format the object to printable string
	//its __repr__ if you are from python
	return fmt.Sprintf("FieldElement{order: %v, num: %v}", *f.order, *f.num)
}

func (f *FieldElement) EqualTo(other *FieldElement) bool {
	/*
		two field element is equal if their order and value are equal
	*/
	return f.order.Cmp(other.order) == 0 && f.num.Cmp(other.num) == 0
}

func (f *FieldElement) checkOrder(other *FieldElement) {
	if f.order.Cmp(other.order) != 0 {
		panic("add need to do on field element with the same order")
	}
}

func (f *FieldElement) Add(other *FieldElement) *FieldElement {

	f.checkOrder(other)
	//remember to do the modulur
	var op big.Int
	return NewFieldElement(f.order, op.Mod(op.Add(f.num, other.num), f.order))
}

func (f *FieldElement) Negate() *FieldElement {
	/*
		for a field element a, its negate is another element b in field such that
		(a + b) % order= 0(remember the modulur over order), because the value of element
		in the field are smaller than its order, we can easily get the negate of a by
		order - a,
	*/
	var op big.Int
	return NewFieldElement(f.order, op.Sub(f.order, f.num))
}

func (f *FieldElement) Subtract(other *FieldElement) *FieldElement {
	//first find the negate of the other
	//add this and the negate of the other
	return f.Add(other.Negate())
}

func (f *FieldElement) Multiplie(other *FieldElement) *FieldElement {
	f.checkOrder(other)
	//multiplie over modulur of order
	var op big.Int
	mul := op.Mul(f.num, other.num)
	return NewFieldElement(f.order, op.Mod(mul, f.order))
}

func (f *FieldElement) Power(power *big.Int) *FieldElement {
	var op big.Int
	powerRes := op.Exp(f.num, power, nil)
	modRes := op.Mod(powerRes, f.order)
	return NewFieldElement(f.order, modRes)
}

func (f *FieldElement) ScalarMul(val *big.Int) *FieldElement {
	var op big.Int
	res := op.Mul(f.num, val)
	res = op.Mod(res, f.order)
	return NewFieldElement(f.order, res)
}
```
Now we need to make sure the changes will not break our logic, let's run our tests again, in main.go, we have following code:
```go

```
