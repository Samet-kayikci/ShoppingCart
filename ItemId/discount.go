package ItemId

func Discount(count int, Cart map[int]Item) float32 {
	var SellerCount []int
	var Scount int = 0
	for i := 1; i < count; i++ {
		if i == 1 {
			SellerCount = append(SellerCount, Cart[i].SellerID)
			Scount++
		} else {
			for j := 1; j < i; j++ {
				if Cart[j].SellerID == Cart[i].SellerID {
					break
				} else if j == i-1 {
					SellerCount = append(SellerCount, Cart[i].SellerID)
					Scount++
				}
			}
		}
	}
	var discount_limited int = 0
	var record_cart []int
	a := 0
	var total_price float32
	var discount_price float32 = 0
	for i := 0; i < Scount; i++ {
		for j := 1; j < count; j++ {
			if SellerCount[i] == Cart[j].SellerID {
				discount_limited++
				record_cart = append(record_cart, j)
				a++
			}
		}

		if discount_limited >= 3 {
			for i := 0; i < a; i++ {
				index := record_cart[i]
				total_price = total_price + Cart[index].Price
			}
			discount_price = 0
			discount_price = discount_price + (total_price * 5 / 100)
			a = 0

			discount_limited = 0
			record_cart = append(record_cart[:0], record_cart[3:]...)
		}

	}
	return discount_price

}

// program çalışıyor discount fonksiyonunda tüm id ler tek arraya atandı mesela 4 id var bunlar 3 3 4 1 ise bu id ler 3 4 1 şeklinde atanmış oldu
// sırada bu id lerin map karşılıkları ile karşılaştırılması var eğer sepette 3 ve daha fazla aynı sellerid içeriği varsa o 3 ürünün toplam fiyatında % 5 indirim olacak

// en son record_card daki indexleri hallettim 1, 2, 3	4, 5, 6 diye devam ediyor sırada kağıt üzerinde bunların işlemini devam ettircem
// program çalıştığındaki çıktı doğrumu ona bakacam
