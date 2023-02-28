package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Привет! Меня зовут Softik, и я хорош в матиматике и не только :)")
		fmt.Println("Вот команды, которые я могу выполнять:")
		fmt.Println("\t1. Распечатать таблицу умножения для заданного вами числа")
		fmt.Println("\t2. Распечатать и посчитать все числа в заданном отрезке, которые кратны заданному вами числу")
		fmt.Println("\t3. Взломать Пентагон")
		fmt.Println("\t4. Мирно завершить работу")
		fmt.Print("Что бы вы хотели? Введите номер команды: ")
		var command int
		fmt.Scan(&command)
		fmt.Println()

		if command == 4 {
			break
		}

		switch command {
		case 1:
			fmt.Println("Распечатать таблицу умножения? Не проблема!")
			fmt.Print("Введите число для которого нужно посчитать: ")
			var multiplier int
			fmt.Scan(&multiplier)
			for i := 1; i <= 10; i++ {
				fmt.Printf("%d x %d = %d\n", multiplier, i, multiplier*i)
			}
			fmt.Println("Вуаля, готово!")
		case 2:
			fmt.Println("Все, базару нет, дайте мне отрезок (два числа через пробел): ")
			var start, finish int
			fmt.Scan(&start, &finish)
			if start > finish {
				start, finish = finish, start
			}
			fmt.Print("Понял, а теперь дайте мне число, кратные которому я должен вывести и посчитать: ")
			var divisor, counter int
			fmt.Scan(&divisor)
			fmt.Println("Окей, смотрите:")
			for i := start; i <= finish; i++ {
				if i%divisor == 0 {
					fmt.Print(i, " ")
					counter++
				}
			}
			fmt.Println()
			fmt.Printf("В отрезке от %d до %d включительно всего %d число(ел) кратных %d. Еще что-нибудь?\n", start, finish, counter, divisor)

		case 3:
			fmt.Println("Принято, запускаю скрипты. Взламываю... Это займет не более 10 секунд )")
			for i := 1; i <= 10; i++ {
				if i == 10 {
					fmt.Println("Я передумал, давайте не будем портить кому-то день на работе ))")
					fmt.Println("Кому-то же придется еще долго восстанавливать все :)")
					time.Sleep(time.Second * 3)
					break
				}
				fmt.Println(i, "...")
				time.Sleep(time.Second)
			}
		default:
			fmt.Println("wrong command")
		}
		fmt.Println()
	}

	fmt.Println("Лааадно, мирно завершаю работу...")
	fmt.Println("Пока! Если что, я тут :)")
}
