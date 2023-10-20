package main

import (
	"MyHouse/family"
	. "MyHouse/furniture"
)

func main() {
	chair := Chair{}
	chair.SetChair("Steel")
	chair.GetChair()

	father := family.Member{}
	father.SetMember("Ivan", "male", 48, "father")
	father.GetMember()
}

func getError(number int) (int, error) {
	if number%2 == 0 {
	}

}
