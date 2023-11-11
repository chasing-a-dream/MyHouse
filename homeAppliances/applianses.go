package homeAppliances

type Appliances struct {
	name    string
	color   string
	voltage int
	width   int
	length  int
}

func (appliance *Appliances) GetAppliance() {
	println(appliance.color, appliance.name, "with voltage", appliance.voltage, appliance.width, "X", appliance.length)
}

func (appliance *Appliances) SetAppliance(name string, color string, voltage int, width int, length int) {
	appliance.name = name
	appliance.color = color
	appliance.voltage = voltage
	appliance.length = length
	appliance.width = width
}
