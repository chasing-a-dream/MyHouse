package furniture

type Chair struct {
	material string
}

func (chair *Chair) GetChair() {
	println("Chair made of", chair.material)
}
func (chair *Chair) SetChair(material string) {
	chair.material = material
}
