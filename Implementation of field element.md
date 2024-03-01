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
	fmt.Printf("negate of field element 44 is : %v\n", f44.Negate())
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

Let's sovle the Substract problem here, for field element a, b, we want to find the the field element c such that c = a - b, notice that a -b is the same as a + (-b), and (-b) is the negate of b, which means c is a plus the negate of b, let's put this into code :
```go
func (f *FieldElement) Subtract(other *FieldElement) *FieldElement {
	//first find the negate of the other
	//add this and the negate of the other
	return f.Add(other.Negate())
}
```
Now let's add some code in main to run the Substract function:
```go
func main() {
    ....
fmt.Printf("field element 44 - 33 is : %v\n", f44.Substract(f33))
	fmt.Printf("field element 33 - 44 is : %v\n", f33.Subtract(f44))

	//it is easy to check (11+33)%57 == 44
	//check (46 + 44) % 57 == 33
	fmt.Printf("check 46 + 44 over modulur 57 is %d\n", (46+44)%57)
	//check by field element
	f46 := ecc.NewFieldElement(57, 46)
	fmt.Printf("field element 46 + 44 is %v\n", f46.Add(f44))
}
```
run the code and we can get the following result:
```go
field element 44 add to field element 33 is : FieldElement{order: 57, num: 20}
negate of field element 44 is : FieldElement{order: 57, num: 37}
field element 44 - 33 is : FieldElement{order: 57, num: 11}
field element 33 - 44 is : FieldElement{order: 57, num: 46}
check 46 + 44 over modulur 57 is 33
field element 46 + 44 is FieldElement{order: 57, num: 33}
```
we can do some simple arithmetic calculation, the result of (46+44) % 57 is indeed 33, which means the logic of our code is correct. Let's see how to add the operation of multiplie and power, there are arithmetic multiple and power over the modulur of order, the code is as following:
```go
func (f *FieldElement) checkOrder(other *FieldElement) {
	if other.order != f.order {
		panic("add need to do on field element with the same order")
	}
}

func (f *FieldElement) Multiplie(other *FieldElement) *FieldElement {
	f.checkOrder(other)
	//multiplie over modulur of order
	return NewFieldElement(f.order, (f.num*other.num)%f.order)
}

func (f *FieldElement) Power(power int64) *FieldElement {
	return NewFieldElement(f.order, uint64(math.Pow(float64(f.num), float64(power)))%f.order)
}
```
we run the newly add code for test, in main.go we add code like following:
```go
func main() {
...
    fmt.Printf("multiplie element 46 with itself is :%v\n", f46.Multiplie(f46))
    fmt.Printf("element 46 with power to 2 is %v\n", f46.Power(2))
}
```
The running result is:
```go
multiplie element 46 with itself is :FieldElement{order: 57, num: 7}
element 46 with power to 2 is FieldElement{order: 57, num: 7}
```
we can see that element 46 mutiplie itself is equal to compute its power of 2. question time now, for finite field with order 19, randomly select one element  k from the set,
compute {k . 0, k . 1, ... k . 18 } and what would you get?

let's use code to solve the problem, first we need to add method for the field element that it can multiplie scalar nmber:
```go
func (f *FieldElement) ScalarMul(val uint64) *FieldElement {
	return NewFieldElement(f.order, (f.num*val)%f.order)
}
```
Now goto main.go, and use the followint code to solve the problem:
```go
package main

import (
	ecc "elliptic_curve"
	"fmt"
	"math/rand"
)

func SolveField19MultiplieSet() {
	//randomly select a num from (1, 18)
	min := 1
	max := 18
	k := rand.Intn(max-min) + min
	fmt.Printf("randomly select k is : %d\n", k)
	element := ecc.NewFieldElement(19, uint64(k))
	for i := 0; i < 19; i++ {
		fmt.Printf("element %d multiplie with %d is %v\n", k, i, element.ScalarMul(uint64(i)))
	}

}

func main() {
	SolveField19MultiplieSet()
}

```
If you run the above code, you may get the following result:
```go
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

