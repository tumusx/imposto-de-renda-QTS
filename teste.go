package main

import "fmt"

func main() {
	var salarioDev float64
	fmt.Scanf("%g", &salarioDev)
	if salarioDev <= 1903.98 {
		fmt.Println("não deve pagar receita federal")
	} else if salarioDev > 1903.98 && salarioDev < 2826.65 {
		fmt.Println("voce deve pagar a receita federal a deducao de 7.5%, portanto:", salarioDev*0.075, "Reais")
	} else if salarioDev > 2826.65 && salarioDev < 3751.05 {
		fmt.Println("voce deve pagar a receita federal a deducao de 15%, portanto:", salarioDev*0.15, "Reais")
	} else if salarioDev > 3751.06 && salarioDev < 4664.68 {
		fmt.Println("voce deve pagar a receita federal a deducao de 22,5%, portanto:", salarioDev*0.225, "Reais")
	} else {
		fmt.Println("voce deve pagar a receita federal a deducao de 27,5%, portanto:", salarioDev*0.275, "Reais")
	}
}
