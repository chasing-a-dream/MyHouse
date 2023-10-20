package furniture

type Furniture struct {
	name     string
	color    string
	material string
	width    int
	length   int
}

func (furniture *Furniture) GetFurniture() {
	println(furniture.color, furniture.name, "made of", furniture.material, furniture.width, "X", furniture.length)
}
func (furniture *Furniture) SetFurniture(name string, color string, material string, width int, length int) {
	furniture.name = name
	furniture.color = color
	furniture.material = material
	furniture.length = length
	furniture.width = width
}
