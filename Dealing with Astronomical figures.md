In bitcoin cryptography, we need to deal with Astronomical figures, the number is so huge that it is easily outpace the total count to atoms in our universe, and maybe 64bits value is 
not enough to represent the figure, and doing operations like add , multipy, power will have result causing overflow if we use 64bit ineger,therefore we may need the help of big package from golang, we will change our code of FieldNumber by using big.Int to represent its value field,
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
package main

import (
	ecc "elliptic_curve"
	"fmt"
	"math/big"
	"math/rand"
)

func SolveField19MultiplieSet() {
	//randomly select a num from (1, 18)
	min := 1
	max := 18
	k := rand.Intn(max-min) + min
	fmt.Printf("randomly select k is : %d\n", k)
	element := ecc.NewFieldElement(big.NewInt(19), big.NewInt(int64(k)))
	for i := 0; i < 19; i++ {
		fmt.Printf("element %d multiplie with %d is %v\n", k, i,
			element.ScalarMul(big.NewInt(int64(i))))
	}

}

func main() {
	f44 := ecc.NewFieldElement(big.NewInt(57), big.NewInt(44))
	f33 := ecc.NewFieldElement(big.NewInt(57), big.NewInt(33))
	// 44 + 33 equal to (44+33) % 57 is 20
	res := f44.Add(f33)
	fmt.Printf("field element 44 add to field element 33 is : %v\n", res)
	//-44 is the negate of field element 44, which is 57 - 44 = 13
	fmt.Printf("negate of field element 44 is : %v\n", f44.Negate())

	fmt.Printf("field element 44 - 33 is : %v\n", f44.Subtract(f33))
	fmt.Printf("field element 33 - 44 is : %v\n", f33.Subtract(f44))

	//it is easy to check (11+33)%57 == 44
	//check (46 + 44) % 57 == 33
	fmt.Printf("check 46 + 44 over modulur 57 is %d\n", (46+44)%57)
	//check by field element
	f46 := ecc.NewFieldElement(big.NewInt(57), big.NewInt(46))
	fmt.Printf("field element 46 + 44 is %v\n", f46.Add(f44))

	SolveField19MultiplieSet()
}
```
running the above code will get the following result:
```go
field element 44 add to field element 33 is : FieldElement{order: 57, num: 20}
negate of field element 44 is : FieldElement{order: 57, num: 13}
field element 44 - 33 is : FieldElement{order: 57, num: 11}
field element 33 - 44 is : FieldElement{order: 57, num: 46}
check 46 + 44 over modulur 57 is 33
field element 46 + 44 is FieldElement{order: 57, num: 33}
randomly select k is : 2
element 2 multiplie with 0 is FieldElement{order: 19, num: 0}
element 2 multiplie with 1 is FieldElement{order: 19, num: 2}
element 2 multiplie with 2 is FieldElement{order: 19, num: 4}
element 2 multiplie with 3 is FieldElement{order: 19, num: 6}
element 2 multiplie with 4 is FieldElement{order: 19, num: 8}
element 2 multiplie with 5 is FieldElement{order: 19, num: 10}
element 2 multiplie with 6 is FieldElement{order: 19, num: 12}
element 2 multiplie with 7 is FieldElement{order: 19, num: 14}
element 2 multiplie with 8 is FieldElement{order: 19, num: 16}
element 2 multiplie with 9 is FieldElement{order: 19, num: 18}
element 2 multiplie with 10 is FieldElement{order: 19, num: 1}
element 2 multiplie with 11 is FieldElement{order: 19, num: 3}
element 2 multiplie with 12 is FieldElement{order: 19, num: 5}
element 2 multiplie with 13 is FieldElement{order: 19, num: 7}
element 2 multiplie with 14 is FieldElement{order: 19, num: 9}
element 2 multiplie with 15 is FieldElement{order: 19, num: 11}
element 2 multiplie with 16 is FieldElement{order: 19, num: 13}
element 2 multiplie with 17 is FieldElement{order: 19, num: 15}
element 2 multiplie with 18 is FieldElement{order: 19, num: 17}
```
by checking the result, we can make sure the changes in FieldElement is not break our logic before. Now Let's give a thought to following problem:
p = 7, 11, 17, 19, 31, what would the following set be:
{1 ^(p-1), 2 ^ (p-1), ... (p-1)^(p-1)}
Let's have code to solve it in main.go: 
```go
func ComputeFieldOrderPower() {
	orders := []int{7, 11, 17, 31}
	for _, p := range orders {
		fmt.Printf("value of p is: %d\n", p)
		for i := 1; i < p; i++ {
			elm := ecc.NewFieldElement(big.NewInt(int64(p)), big.NewInt(int64(i)))
			fmt.Printf("for element: %v, its power of p - 1 is: %v\n", elm,
				elm.Power(big.NewInt(int64(p-1))))
		}
		fmt.Println("-------------------------------")
	}
}

func main() {
    ComputeFieldOrderPower()
}
```
The result is following:
```go
value of p is: 7
for element: FieldElement{order: 7, num: 1}, its power of p - 1 is: FieldElement{order: 7, num: 1}
for element: FieldElement{order: 7, num: 2}, its power of p - 1 is: FieldElement{order: 7, num: 1}
for element: FieldElement{order: 7, num: 3}, its power of p - 1 is: FieldElement{order: 7, num: 1}
for element: FieldElement{order: 7, num: 4}, its power of p - 1 is: FieldElement{order: 7, num: 1}
for element: FieldElement{order: 7, num: 5}, its power of p - 1 is: FieldElement{order: 7, num: 1}
for element: FieldElement{order: 7, num: 6}, its power of p - 1 is: FieldElement{order: 7, num: 1}
-------------------------------
value of p is: 11
for element: FieldElement{order: 11, num: 1}, its power of p - 1 is: FieldElement{order: 11, num: 1}
for element: FieldElement{order: 11, num: 2}, its power of p - 1 is: FieldElement{order: 11, num: 1}
for element: FieldElement{order: 11, num: 3}, its power of p - 1 is: FieldElement{order: 11, num: 1}
for element: FieldElement{order: 11, num: 4}, its power of p - 1 is: FieldElement{order: 11, num: 1}
for element: FieldElement{order: 11, num: 5}, its power of p - 1 is: FieldElement{order: 11, num: 1}
for element: FieldElement{order: 11, num: 6}, its power of p - 1 is: FieldElement{order: 11, num: 1}
for element: FieldElement{order: 11, num: 7}, its power of p - 1 is: FieldElement{order: 11, num: 1}
for element: FieldElement{order: 11, num: 8}, its power of p - 1 is: FieldElement{order: 11, num: 1}
for element: FieldElement{order: 11, num: 9}, its power of p - 1 is: FieldElement{order: 11, num: 1}
for element: FieldElement{order: 11, num: 10}, its power of p - 1 is: FieldElement{order: 11, num: 1}
-------------------------------
value of p is: 17
for element: FieldElement{order: 17, num: 1}, its power of p - 1 is: FieldElement{order: 17, num: 1}
for element: FieldElement{order: 17, num: 2}, its power of p - 1 is: FieldElement{order: 17, num: 1}
for element: FieldElement{order: 17, num: 3}, its power of p - 1 is: FieldElement{order: 17, num: 1}
for element: FieldElement{order: 17, num: 4}, its power of p - 1 is: FieldElement{order: 17, num: 1}
for element: FieldElement{order: 17, num: 5}, its power of p - 1 is: FieldElement{order: 17, num: 1}
for element: FieldElement{order: 17, num: 6}, its power of p - 1 is: FieldElement{order: 17, num: 1}
for element: FieldElement{order: 17, num: 7}, its power of p - 1 is: FieldElement{order: 17, num: 1}
for element: FieldElement{order: 17, num: 8}, its power of p - 1 is: FieldElement{order: 17, num: 1}
for element: FieldElement{order: 17, num: 9}, its power of p - 1 is: FieldElement{order: 17, num: 1}
for element: FieldElement{order: 17, num: 10}, its power of p - 1 is: FieldElement{order: 17, num: 1}
for element: FieldElement{order: 17, num: 11}, its power of p - 1 is: FieldElement{order: 17, num: 1}
for element: FieldElement{order: 17, num: 12}, its power of p - 1 is: FieldElement{order: 17, num: 1}
for element: FieldElement{order: 17, num: 13}, its power of p - 1 is: FieldElement{order: 17, num: 1}
for element: FieldElement{order: 17, num: 14}, its power of p - 1 is: FieldElement{order: 17, num: 1}
for element: FieldElement{order: 17, num: 15}, its power of p - 1 is: FieldElement{order: 17, num: 1}
for element: FieldElement{order: 17, num: 16}, its power of p - 1 is: FieldElement{order: 17, num: 1}
-------------------------------
value of p is: 31
for element: FieldElement{order: 31, num: 1}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 2}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 3}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 4}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 5}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 6}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 7}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 8}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 9}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 10}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 11}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 12}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 13}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 14}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 15}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 16}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 17}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 18}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 19}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 20}, its power of p - 1 is: FieldElement{order: 31, num: 1}
my@MACdeMacBook-Air bitcoin % go run main.go
value of p is: 7
for element: FieldElement{order: 7, num: 1}, its power of p - 1 is: FieldElement{order: 7, num: 1}
for element: FieldElement{order: 7, num: 2}, its power of p - 1 is: FieldElement{order: 7, num: 1}
for element: FieldElement{order: 7, num: 3}, its power of p - 1 is: FieldElement{order: 7, num: 1}
for element: FieldElement{order: 7, num: 4}, its power of p - 1 is: FieldElement{order: 7, num: 1}
for element: FieldElement{order: 7, num: 5}, its power of p - 1 is: FieldElement{order: 7, num: 1}
for element: FieldElement{order: 7, num: 6}, its power of p - 1 is: FieldElement{order: 7, num: 1}
-------------------------------
value of p is: 11
for element: FieldElement{order: 11, num: 1}, its power of p - 1 is: FieldElement{order: 11, num: 1}
for element: FieldElement{order: 11, num: 2}, its power of p - 1 is: FieldElement{order: 11, num: 1}
for element: FieldElement{order: 11, num: 3}, its power of p - 1 is: FieldElement{order: 11, num: 1}
for element: FieldElement{order: 11, num: 4}, its power of p - 1 is: FieldElement{order: 11, num: 1}
for element: FieldElement{order: 11, num: 5}, its power of p - 1 is: FieldElement{order: 11, num: 1}
for element: FieldElement{order: 11, num: 6}, its power of p - 1 is: FieldElement{order: 11, num: 1}
for element: FieldElement{order: 11, num: 7}, its power of p - 1 is: FieldElement{order: 11, num: 1}
for element: FieldElement{order: 11, num: 8}, its power of p - 1 is: FieldElement{order: 11, num: 1}
for element: FieldElement{order: 11, num: 9}, its power of p - 1 is: FieldElement{order: 11, num: 1}
for element: FieldElement{order: 11, num: 10}, its power of p - 1 is: FieldElement{order: 11, num: 1}
-------------------------------
value of p is: 17
for element: FieldElement{order: 17, num: 1}, its power of p - 1 is: FieldElement{order: 17, num: 1}
for element: FieldElement{order: 17, num: 2}, its power of p - 1 is: FieldElement{order: 17, num: 1}
for element: FieldElement{order: 17, num: 3}, its power of p - 1 is: FieldElement{order: 17, num: 1}
for element: FieldElement{order: 17, num: 4}, its power of p - 1 is: FieldElement{order: 17, num: 1}
for element: FieldElement{order: 17, num: 5}, its power of p - 1 is: FieldElement{order: 17, num: 1}
for element: FieldElement{order: 17, num: 6}, its power of p - 1 is: FieldElement{order: 17, num: 1}
for element: FieldElement{order: 17, num: 7}, its power of p - 1 is: FieldElement{order: 17, num: 1}
for element: FieldElement{order: 17, num: 8}, its power of p - 1 is: FieldElement{order: 17, num: 1}
for element: FieldElement{order: 17, num: 9}, its power of p - 1 is: FieldElement{order: 17, num: 1}
for element: FieldElement{order: 17, num: 10}, its power of p - 1 is: FieldElement{order: 17, num: 1}
for element: FieldElement{order: 17, num: 11}, its power of p - 1 is: FieldElement{order: 17, num: 1}
for element: FieldElement{order: 17, num: 12}, its power of p - 1 is: FieldElement{order: 17, num: 1}
for element: FieldElement{order: 17, num: 13}, its power of p - 1 is: FieldElement{order: 17, num: 1}
for element: FieldElement{order: 17, num: 14}, its power of p - 1 is: FieldElement{order: 17, num: 1}
for element: FieldElement{order: 17, num: 15}, its power of p - 1 is: FieldElement{order: 17, num: 1}
for element: FieldElement{order: 17, num: 16}, its power of p - 1 is: FieldElement{order: 17, num: 1}
-------------------------------
value of p is: 19
for element: FieldElement{order: 19, num: 1}, its power of p - 1 is: FieldElement{order: 19, num: 1}
for element: FieldElement{order: 19, num: 2}, its power of p - 1 is: FieldElement{order: 19, num: 1}
for element: FieldElement{order: 19, num: 3}, its power of p - 1 is: FieldElement{order: 19, num: 1}
for element: FieldElement{order: 19, num: 4}, its power of p - 1 is: FieldElement{order: 19, num: 1}
for element: FieldElement{order: 19, num: 5}, its power of p - 1 is: FieldElement{order: 19, num: 1}
for element: FieldElement{order: 19, num: 6}, its power of p - 1 is: FieldElement{order: 19, num: 1}
for element: FieldElement{order: 19, num: 7}, its power of p - 1 is: FieldElement{order: 19, num: 1}
for element: FieldElement{order: 19, num: 8}, its power of p - 1 is: FieldElement{order: 19, num: 1}
for element: FieldElement{order: 19, num: 9}, its power of p - 1 is: FieldElement{order: 19, num: 1}
for element: FieldElement{order: 19, num: 10}, its power of p - 1 is: FieldElement{order: 19, num: 1}
for element: FieldElement{order: 19, num: 11}, its power of p - 1 is: FieldElement{order: 19, num: 1}
for element: FieldElement{order: 19, num: 12}, its power of p - 1 is: FieldElement{order: 19, num: 1}
for element: FieldElement{order: 19, num: 13}, its power of p - 1 is: FieldElement{order: 19, num: 1}
for element: FieldElement{order: 19, num: 14}, its power of p - 1 is: FieldElement{order: 19, num: 1}
for element: FieldElement{order: 19, num: 15}, its power of p - 1 is: FieldElement{order: 19, num: 1}
for element: FieldElement{order: 19, num: 16}, its power of p - 1 is: FieldElement{order: 19, num: 1}
for element: FieldElement{order: 19, num: 17}, its power of p - 1 is: FieldElement{order: 19, num: 1}
for element: FieldElement{order: 19, num: 18}, its power of p - 1 is: FieldElement{order: 19, num: 1}
-------------------------------
value of p is: 31
for element: FieldElement{order: 31, num: 1}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 2}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 3}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 4}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 5}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 6}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 7}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 8}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 9}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 10}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 11}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 12}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 13}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 14}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 15}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 16}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 17}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 18}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 19}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 20}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 21}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 22}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 23}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 24}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 25}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 26}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 27}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 28}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 29}, its power of p - 1 is: FieldElement{order: 31, num: 1}
for element: FieldElement{order: 31, num: 30}, its power of p - 1 is: FieldElement{order: 31, num: 1}
-------------------------------
```
you can see the set is all 1 regardless the field order, that means for any finit field with order p, for any element k in the field, we would have:
k ^(p-1) % p == 1
This is an important conclution, we will use it to drive our cryptograhpy algorithm in later videos
