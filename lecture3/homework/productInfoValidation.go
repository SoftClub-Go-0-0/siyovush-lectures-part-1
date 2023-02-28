package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	fmt.Print("Enter product creation mode (single/multiple): ")
	var creationMode string
	fmt.Scan(&creationMode)

	var (
		title          string
		description    string
		price          int
		discount       int
		categoryID     int
		manufacturerID int
		image          string
	)

	stdin := bufio.NewScanner(os.Stdin)

	switch creationMode {
	case "single":
		fmt.Println("Okay, let's create one product")
	GetInfoSingle:
		fmt.Println("Enter product's info")

		fmt.Println("\nTitle:")
		stdin.Scan()
		title = stdin.Text()

		fmt.Println("\nDescription:")
		stdin.Scan()
		description = stdin.Text()

		fmt.Println("\nPrice:")
		fmt.Scan(&price)

		fmt.Println("\nDiscount:")
		fmt.Scan(&discount)

		fmt.Println("\nCategory ID:")
		fmt.Scan(&categoryID)

		fmt.Println("\nManufacturer ID:")
		fmt.Scan(&manufacturerID)

		fmt.Println("\nImage:")
		stdin.Scan()
		image = stdin.Text()

		if !validate(
			title,
			description,
			price,
			discount,
			categoryID,
			manufacturerID,
			image,
		) {
			fmt.Println("===========================================")
			fmt.Println("= Error: bad data. Re-enter all the info. =")
			fmt.Println("===========================================")
			goto GetInfoSingle
		}

		fmt.Println("\nAll good. Created product:")
		fmt.Printf("\tID:\t\t%d\n", rand.Int()+1)
		fmt.Printf("\tTitle:\t\t%s\n", title)
		fmt.Printf("\tDescription:\t%s\n", description)
		fmt.Printf("\tPrice:\t\t%d\n", price)
		fmt.Printf("\tDiscount:\t%d\n", discount)
		fmt.Printf("\tCategoryID:\t%d\n", categoryID)
		fmt.Printf("\tManufacturerID:\t%d\n", manufacturerID)
		fmt.Printf("\tImage:\t\t%s\n", image)

	case "multiple":
		fmt.Println("Okay, let's create multiple products")
		fmt.Println("How many products do you want to create?")
		var productsNumber int
		fmt.Scan(&productsNumber)
		for i := 0; i < productsNumber; i++ {
		GetInfoMultiple:
			fmt.Println("\nEnter", i+1, "product's info")

			fmt.Println("\nTitle:")
			stdin.Scan()
			title = stdin.Text()

			fmt.Println("\nDescription:")
			stdin.Scan()
			description = stdin.Text()

			fmt.Println("\nPrice:")
			fmt.Scan(&price)

			fmt.Println("\nDiscount:")
			fmt.Scan(&discount)

			fmt.Println("\nCategory ID:")
			fmt.Scan(&categoryID)

			fmt.Println("\nManufacturer ID:")
			fmt.Scan(&manufacturerID)

			fmt.Println("\nImage:")
			stdin.Scan()
			image = stdin.Text()

			if !validate(
				title,
				description,
				price,
				discount,
				categoryID,
				manufacturerID,
				image,
			) {
				fmt.Println("===========================================")
				fmt.Println("= Error: bad data. Re-enter all the info. =")
				fmt.Println("===========================================")
				goto GetInfoMultiple
			}

			fmt.Println("\nAll good. Created product:")
			fmt.Printf("\tID:\t\t%d\n", rand.Int()+1)
			fmt.Printf("\tTitle:\t\t%s\n", title)
			fmt.Printf("\tDescription:\t%s\n", description)
			fmt.Printf("\tPrice:\t\t%d\n", price)
			fmt.Printf("\tDiscount:\t%d\n", discount)
			fmt.Printf("\tCategoryID:\t%d\n", categoryID)
			fmt.Printf("\tManufacturerID:\t%d\n", manufacturerID)
			fmt.Printf("\tImage:\t\t%s\n", image)
		}
		fmt.Println("\nCreated", productsNumber, "products")
	default:
		fmt.Println("Wrong creation mode!")
	}
}

func validate(title, description string, price, discount, categoryID, manufacturerID int, image string) bool {
	if !(len(title) != 0 && len(title) <= 100) {
		return false
	}

	if !(len(description) != 0 && len(description) <= 1_000) {
		return false
	}

	if !(price >= 20 && price <= 50_000) {
		return false
	}

	if !(discount >= 0 && discount <= 95) {
		return false
	}

	if !(categoryID >= 1 && categoryID <= 10) {
		return false
	}

	if !(manufacturerID >= 1 && manufacturerID <= 50) {
		return false
	}

	if !(len(image) <= 300) {
		return false
	}

	return true
}
