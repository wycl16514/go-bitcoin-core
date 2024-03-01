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
	return fmt.Sprintf("FieldElement{order: %s, num: %s}",
		f.order.String(), f.num.String())
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

func (f *FieldElement) Multiply(other *FieldElement) *FieldElement {
	f.checkOrder(other)
	//multiplie over modulur of order
	var op big.Int
	mul := op.Mul(f.num, other.num)
	return NewFieldElement(f.order, op.Mod(mul, f.order))
}

func (f *FieldElement) Power(power *big.Int) *FieldElement {
	/*
		k ^ (p-1) % p = 1, we can compute t = power % (p-1)
		and then k ^ power % p == k ^ t %p
	*/
	var op big.Int
	t := op.Mod(power, op.Sub(f.order, big.NewInt(int64(1))))
	powerRes := op.Exp(f.num, t, nil)
	modRes := op.Mod(powerRes, f.order)
	return NewFieldElement(f.order, modRes)
}

func (f *FieldElement) ScalarMul(val *big.Int) *FieldElement {
	var op big.Int
	res := op.Mul(f.num, val)
	res = op.Mod(res, f.order)
	return NewFieldElement(f.order, res)
}

func (f *FieldElement) Divide(other *FieldElement) *FieldElement {
	f.checkOrder(other)
	var op big.Int
	otherReverse := other.Power(op.Sub(f.order, big.NewInt(int64(2))))
	return f.Multiply(otherReverse)
}
