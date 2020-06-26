package main

import (
	"flag"
	"fmt"
	"os"
	"udemygo/crud/dictionnary"
)

func main() {
	action := flag.String("action", "list", "Action to perform on the dictionnary") //On initialise notre CLI

	d, err := dictionnary.New("./badger")
	handleError(err)
	defer d.Close()

	//d.Add("python", "Language très simple")
	//entry, _ := d.Get("java")

	//d.Remove("java")

	flag.Parse() //On parse ce qu'on reçoit en ligne de commande
	switch *action {
	case "list":
		actionList(d)
	case "add":
		actionAdd(d, flag.Args())
	case "define":
		actionDefine(d, flag.Args())
	case "remove":
		actionRemove(d, flag.Args())
	default:
		fmt.Println("Action inconnu: :", *action)
	}
}

func actionList(d *dictionnary.Dictionnary) {
	words, entries, err := d.List()
	handleError(err)
	fmt.Println("Contenu de notre base de donnée")
	for _, word := range words {
		fmt.Println(entries[word])
	}
}

func actionAdd(d *dictionnary.Dictionnary, args []string) {
	word := args[0]
	definition := args[1]
	err := d.Add(word, definition)
	handleError(err)
	fmt.Printf("'%v' ajouté à notre base de donné\n", word)
}

func actionDefine(d *dictionnary.Dictionnary, args []string) {
	word := args[0]
	entry, err := d.Get(word)
	handleError(err)
	fmt.Println(entry)
}

func actionRemove(d *dictionnary.Dictionnary, args []string) {
	word := args[0]
	err := d.Remove(word)
	handleError(err)
	fmt.Printf("'%v' a été supprimé de notre base de donnée\n", word)
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("Erreur : %v\n", err)
		os.Exit(1)
	}
}
