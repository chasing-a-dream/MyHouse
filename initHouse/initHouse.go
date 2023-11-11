package initHouse

import (
	"MyHouse/family"
	. "MyHouse/furniture"
	"MyHouse/homeAppliances"
)

func InitHouse() {
	initFamily()
	initFurniture()
	initAppliances()
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

func initAppliances() {
	print("Appliances\n")
	fridge := homeAppliances.Appliances{}
	fridge.SetAppliance("fridge", "white", 220, 70, 70)
	fridge.GetAppliance()

	washingMachine := homeAppliances.Appliances{}
	washingMachine.SetAppliance("washing machine", "white", 220, 80, 60)
	washingMachine.GetAppliance()
}
