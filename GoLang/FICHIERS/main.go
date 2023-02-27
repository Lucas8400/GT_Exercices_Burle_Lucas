package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ChoixFichier() string {
	var choix string
	fmt.Println("Entrez le nom du fichier que vous souhaitez modifier :")
	fmt.Scanln(&choix)
	_, err := ioutil.ReadFile(choix + ".txt")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		fmt.Println("Entrez un autre nom de fichier :")
		fmt.Scanln(&choix)
	}
	return (choix + ".txt")
}

func Scan() int {
	var choix int
	fmt.Println("Que voulez-vous faire ?")
	fmt.Println("1 - Récupérer tout le texte contenu dans le fichier")
	fmt.Println("2 - Ajouter du texte dans le fichier")
	fmt.Println("3 - Supprimer tout le contenu du fichier")
	fmt.Println("4 - Remplacer le contenu du fichier par du texte donné")
	fmt.Print("Entrez le numéro correspondant à votre choix : ")
	fmt.Scanln(&choix)
	return choix
}

func main() {
	fichier := ChoixFichier()
	switch Scan() {
	case 1:
		contenu, _ := ioutil.ReadFile(fichier)
		fmt.Println(string(contenu))

	case 2:
		fmt.Print("Entrez le texte à ajouter : ")
		var texte string
		fmt.Scanln(&texte)
		fichier, _ := os.OpenFile(fichier, os.O_APPEND|os.O_WRONLY, 0644)
		defer fichier.Close()
		if _, err := fichier.WriteString(texte); err != nil {
			fmt.Println("Erreur lors de l'écriture dans le fichier :", err)
			return
		}
		fmt.Println("Le texte a été ajouté avec succès !")

	case 3:
		if err := os.Truncate(fichier, 0); err != nil {
			fmt.Println("Erreur lors de la suppression du contenu du fichier :", err)
			return
		}
		fmt.Println("Le contenu du fichier a été supprimé avec succès !")

	case 4:
		fmt.Print("Entrez le texte à écrire dans le fichier : ")
		var texte string
		fmt.Scanln(&texte)
		if err := ioutil.WriteFile(fichier, []byte(texte), 0644); err != nil {
			fmt.Println("Erreur lors de l'écriture dans le fichier :", err)
			return
		}
		fmt.Println("Le fichier a été mis à jour avec succès !")

	default:
		fmt.Println("Choix invalide.")
		Scan()
	}
}
