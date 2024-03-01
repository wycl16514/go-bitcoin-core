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

/*
p = 7, 11, 17, 19, 31, what would the following set be:
{1 ^(p-1), 2 ^ (p-1), ... (p-1)^(p-1)}
*/

func ComputeFieldOrderPower() {
	orders := []int{7, 11, 17, 19, 31}
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

	// f44 := ecc.NewFieldElement(big.NewInt(57), big.NewInt(44))
	// f33 := ecc.NewFieldElement(big.NewInt(57), big.NewInt(33))
	// // 44 + 33 equal to (44+33) % 57 is 20
	// res := f44.Add(f33)
	// fmt.Printf("field element 44 add to field element 33 is : %v\n", res)
	// //-44 is the negate of field element 44, which is 57 - 44 = 13
	// fmt.Printf("negate of field element 44 is : %v\n", f44.Negate())

	// fmt.Printf("field element 44 - 33 is : %v\n", f44.Subtract(f33))
	// fmt.Printf("field element 33 - 44 is : %v\n", f33.Subtract(f44))

	// //it is easy to check (11+33)%57 == 44
	// //check (46 + 44) % 57 == 33
	// fmt.Printf("check 46 + 44 over modulur 57 is %d\n", (46+44)%57)
	// //check by field element
	// f46 := ecc.NewFieldElement(big.NewInt(57), big.NewInt(46))
	// fmt.Printf("field element 46 + 44 is %v\n", f46.Add(f44))

	// SolveField19MultiplieSet()

	//ComputeFieldOrderPower()

	f2 := ecc.NewFieldElement(big.NewInt(int64(19)), big.NewInt(int64(2)))
	f7 := ecc.NewFieldElement(big.NewInt(int64(19)), big.NewInt(int64(7)))
	fmt.Printf("field element 2 / 7 with order 19 is %v\n", f2.Divide(f7))

	f46 := ecc.NewFieldElement(big.NewInt(57), big.NewInt(46))
	fmt.Printf("field element 46 * 46 with order 57: %v\n", f46.Multiply(f46))
	fmt.Printf("field element 46 ^ (58) is %v\n", f46.Power(big.NewInt(int64(58))))

}
