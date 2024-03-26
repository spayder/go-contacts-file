package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func saveToFile(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()
	encoder := json.NewEncoder(file)

	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}

	return nil
}

func loadFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(contacts)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var contacts []Contact
	err := loadFromFile(&contacts)
	if err != nil {
		fmt.Println("Error loading contacts: ", err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("==== Contact Manager====\n",
			"1. Add a contact\n",
			"2. Show all contacts\n",
			"3. Exit\n",
			"Select an option: ")

		var option int
		_, err = fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error reading user option: ", err)
			return
		}

		switch option {
		case 1:
			var c Contact
			fmt.Print("Name: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Phone: ")
			c.Phone, _ = reader.ReadString('\n')

			contacts = append(contacts, c)
			if err := saveToFile(contacts); err != nil {
				fmt.Println("Error saving contact: ", err)
			}

		case 2:
			fmt.Println("=====================================")
			for index, contact := range contacts {
				fmt.Printf("%d. Name: %s Email: %s Phone: %s", index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("=====================================")

		case 3:
			return
		default:
			fmt.Println("Invalid option ...")
		}
	}
}
