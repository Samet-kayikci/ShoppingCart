
package ItemId

type VasItem struct {
	Insurance float32
	Cargo     int
	Id        int
}

type Item struct {
	ID       int
	SellerID int
	Name     string
	Price    float32
}

type DigitItem struct {
	Item
	VasItem
}

type DefaultItem struct {
	Item
	VasItem
}

var Id = make(map[int]Item)
var VasId = make(map[int]VasItem)

func ItemID() {
	digitItem := DigitItem{
		Item: Item{
			ID: 5580,
		},
		VasItem: VasItem{
			Id:    7391,
			Cargo: 3,
		},
	}

	defaultItem := DefaultItem{
		Item: Item{
			ID: 4311,
		},
		VasItem: VasItem{
			Id:    8080,
			Cargo: 4,
		},
	}
	Id[digitItem.ID] = digitItem.Item
	Id[defaultItem.ID] = defaultItem.Item
	VasId[defaultItem.Id] = defaultItem.VasItem
	VasId[digitItem.Id] = digitItem.VasItem
}
