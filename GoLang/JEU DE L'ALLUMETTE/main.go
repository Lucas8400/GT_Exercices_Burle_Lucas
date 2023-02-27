package main

import (
	"fmt"
)

func main() {
	var allumettes int
	fmt.Println("Choisis un nombre d'allumettes supérieur à 4 :")
	fmt.Scanln(&allumettes)

	if allumettes < 4 {
		fmt.Println("Le nombre d'allumettes doit être supérieur à 4")
		fmt.Scanln(&allumettes)
	}
	joueur1 := "Joueur 1"
	joueur2 := "Joueur 2"
	joueur_actuel := joueur1

	for allumettes > 0 {
		allumettes_choisies := 0
		allumettes_restantes := 0
		fmt.Println("Au tour du ", joueur_actuel)
		fmt.Println("Choisissez entre 1 et 3 allumettes :")
		fmt.Scanln(&allumettes_choisies)
		if allumettes_choisies < 1 || allumettes_choisies > 3 {
			fmt.Println("Vous devez chosir entre 1 et 3 allumettes :")
			fmt.Scanln(&allumettes_choisies)
		} else {
			allumettes_restantes = allumettes - allumettes_choisies
		}
		if joueur_actuel == joueur1 && allumettes_restantes > 0 {
			joueur_actuel = joueur2
		} else if allumettes_restantes > 0 {
			joueur_actuel = joueur1
		}
		fmt.Println("Il reste", allumettes_restantes, "allumettes")
		allumettes = allumettes_restantes
		if allumettes_restantes == 0 {
			fmt.Println("Le ", joueur_actuel, "a perdu !")
		}
	}
}
