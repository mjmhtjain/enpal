package domain

type Product string

const (
	solarPanels = "SolarPanels"
	heatpumps   = "Heatpumps"
)

func GetValidProductsMap() map[Product]bool {
	var validProductMap = map[Product]bool{
		solarPanels: false,
		heatpumps:   false,
	}

	return validProductMap
}

func (l *Product) ToString() string {
	return string(*l)
}
