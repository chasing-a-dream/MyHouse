package initHouse

import (
	"MyHouse/family"
	. "MyHouse/furniture"
)

func InitHouse() {
	initFamily()
	initFurniture()

}

func initFamily() {
	print("Family\n")
	father := family.Member{}
	father.SetMember("Ivan", "male", 45, "father")
	father.GetMember()

	mother := family.Member{}
	mother.SetMember("Maria", "female", 38, "mother")
	mother.GetMember()

	son := family.Member{}
	son.SetMember("BOGdan", "male", 13, "son")
	son.GetMember()

	daughter := family.Member{}
	daughter.SetMember("Silfietta", "female", 16, "daughter")
	daughter.GetMember()

	granny := family.Member{}
	granny.SetMember("Avgustina", "female", 69, "grandmother")
	granny.GetMember()
}

func initFurniture() {
	print("Furniture\n")
	sofa := Furniture{}
	sofa.SetFurniture("sofa", "black", "wood", 70, 200)
	sofa.GetFurniture()

	bed := Furniture{}
	bed.SetFurniture("bed", "blue", "wood", 120, 210)
	bed.GetFurniture()

	table := Furniture{}
	table.SetFurniture("table", "white", "glass", 100, 100)
	table.GetFurniture()

	chair := Furniture{}
	chair.SetFurniture("chair", "orange", "steel", 50, 50)
	chair.GetFurniture()

	bookshelf := Furniture{}
	bookshelf.SetFurniture("bookshelf", "dark_green", "wood", 150, 40)
	bookshelf.GetFurniture()
}
