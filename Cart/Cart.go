package main

import (
	"Shopping_Cart/ItemId"
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var Cart = make(map[int]ItemId.Item)
var count = 1
var Calculate_Insurance float64

func addItem(id int, sellerid int, name string, price float32) {
	item, exists := ItemId.Id[id]
	item.SellerID = sellerid
	if !exists {
		fmt.Println("Item not found")
		return
	}
	var vasid int

	if id == 5580 {
		vasid = 7391
	} else if id == 4311 {
		vasid = 8080
	}

	vasitem, exists := ItemId.VasId[vasid]
	if !exists {
		fmt.Println("Item not found")
		return
	}
	price = price + float32(vasitem.Cargo)
	if count <= 10 {
		item.Name = name
		item.Price = price
		Cart[count] = item
	}
	count++
}
func ViewCart() {
	file, err := os.Open("CartRecord.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(fileContent))

}

func DeleteItem(id int) {
	delete(Cart, id)
}

func record() {
	file, err := os.Create("CartRecord.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Cart saved")
	}
	defer file.Close()

	for _, ItemID := range Cart {
		fmt.Fprintln(file, "ID:", ItemID.ID)
		fmt.Fprintln(file, "SellerId:", ItemID.SellerID)
		fmt.Fprintln(file, "Name:", ItemID.Name)
		fmt.Fprintln(file, "price:", ItemID.Price)
		fmt.Fprintln(file)
	}
}

func addInsurance(VasID int, CartId int) {
	vasitem, exists := ItemId.VasId[VasID]
	if !exists {
		fmt.Println("error")
		return
	}
	newItem := Cart[CartId]
	vasitem.Insurance = Cart[CartId].Price * 10 / 100
	newItem.Price = newItem.Price + vasitem.Insurance
	Cart[CartId] = newItem

}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "ViewCart" {
		ViewCart()
	} else {
		ItemId.ItemID()
		addItem(5580, 5141, "computer", 200)
		addItem(5580, 1986, "phone", 200)
		addItem(5580, 5141, "TV", 200)
		addItem(4311, 6753, "table", 100)
		addItem(4311, 6753, "chair", 300)
		addItem(4311, 6753, "carpet", 250)
		addInsurance(7391, 1)
		record()
		var total float32 = 0

		for i := 1; i <= count; i++ {
			total = total + Cart[i].Price
		}

		discount_price := ItemId.Discount(count, Cart)
		total = total - discount_price
		fmt.Println("discount total:", total)
	}
}

// indirim konusu tamamen hazır
// eğer aynı sellerId den 3 ya da daha fazla ürün var ise o 3 ürün kapsamında %5 indirim yapılıyor
