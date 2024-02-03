package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	if order == "burger" {
		burger := food{}
		burger.preptime = 15
		return burger.preptime
	}
	if order == "chips" {
		chips := food{}
		chips.preptime = 10
		return chips.preptime
	}
	if order == "nuggets" {
		nuggets := food{}
		nuggets.preptime = 12
		return nuggets.preptime
	}
	return 404
}
