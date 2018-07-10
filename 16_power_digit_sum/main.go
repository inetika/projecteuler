package main

import (
	"fmt"
	"math/big"
)

func main() {
	//calculating 2^1000
	power := new(big.Int).Exp(big.NewInt(2), big.NewInt(1000), nil)

	//Initiating the sum of digits variable
	sum := new(big.Int)

	//Going through each digit while remainder is > 10
	for power.Cmp(big.NewInt(10)) > 0 {
		//power%10
		remainder := new(big.Int).Mod(power, big.NewInt(10))
		//adding remainder to the sum variable
		sum = new(big.Int).Add(sum, remainder)
		//separating remainder from the power to proceed with the next digit
		power = power.Div(power, big.NewInt(10))
	}
	//We are left with either 10 or number less than 10

	//If we have 10 then just add 1 to the total sum (+0 and +1)
	if power.Cmp(big.NewInt(10)) == 0 {
		sum = new(big.Int).Add(sum, big.NewInt(1))
		//If we have any digit less than 10 then we add this digit to the total sum
	} else {
		sum = new(big.Int).Add(sum, power)
	}

	//Print result
	fmt.Println(sum)
}
