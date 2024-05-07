package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {

	var accountBalance, err = fileops.ReadFloatFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------------------")
	}

	fmt.Println("Bienvenido al Bank Suruka")
	fmt.Println(randomdata.City())

	//for i == 0; i < 20; i++ {}
	for {

		Introduction()

		var numero int
		fmt.Print("Ingesa el numero de opción: ")
		fmt.Scan(&numero)

		//fmt.Println("Elegiste la opción: ", numero)

		switch numero {

		case 1:

			fmt.Println("El saldo de tu cuenta es:", accountBalance, err)

		case 2:

			var deposito float64
			fmt.Print("Ingrese el monto de depositar: ")
			fmt.Scan(&deposito)

			if deposito <= 0 {
				println("Valor ingreso inválido, del monto debe ser mayor a 0")
				continue
			}

			accountBalance += deposito

			fmt.Print("Tu nuevo saldo es: ", accountBalance)

			fileops.WriteFloatTofile(accountBalance, accountBalanceFile)

		case 3:

			var retiro float64
			fmt.Print("Ingrese el monto a retirar: ")
			fmt.Scan(&retiro)

			if retiro > accountBalance {
				println("Valor ingreso inválido, del monto debe ser menor al saldo total")
				continue

			}

			accountBalance -= retiro

			fmt.Print("Tu nuevo saldo es: ", accountBalance)

			fileops.WriteFloatTofile(accountBalance, accountBalanceFile)

		default:

			println("Gracias por visitar Bank Suruka")
			return

		}

		/* if numero == 1 {

			fmt.Println("El saldo de tu cuenta es:", accountBalance)

		} else if numero == 2 {

			var deposito float64
			fmt.Print("Ingrese el monto de depositar: ")
			fmt.Scan(&deposito)

			if deposito <= 0 {
				println("Valor ingreso inválido, del monto debe ser mayor a 0")
				continue
			}

			accountBalance += deposito

			fmt.Print("Tu nuevo saldo es: ", accountBalance)

		} else if numero == 3 {

			var retiro float64
			fmt.Print("Ingrese el monto a retirar: ")
			fmt.Scan(&retiro)

			if retiro > accountBalance {
				println("Valor ingreso inválido, del monto debe ser menor al saldo total")
				continue

			}

			accountBalance -= retiro

			fmt.Print("Tu nuevo saldo es: ", accountBalance)
		} else {

			println("Gracias por visitar Bank")
			break
		} */

	}

}
