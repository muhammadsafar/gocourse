package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main72() {

	//tmpl := template.New("example")

	//tmpl, err := template.New("example").Parse("Welcome, {{.name}}! How are you doing?")

	// if err != nil {
	// 	panic(err)
	// }
	//template.Must - > automatically panic so no needed error handling
	tmpl := template.Must(template.New("example").Parse("Welcome, {{.name}}! How are you doing?\n"))

	//define data for the welcome message template

	data := map[string]interface{}{
		"name": "Abdullah",
		"age":  12,
	}

	err := tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name >> ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	//define namd templates for different type of

	templates := map[string]string{
		"welcome":      "Welcome, {{.name}}! we're glad you joined",
		"notification": "{{.nm}}, you have a new notification: {{.notif}}",
		"error":        "Oops! An error occured : {{.errorMessage}}",
	}

	// Parse and store template
	parsedTemplates := make(map[string]*template.Template)
	for name, tmpl := range templates {
		parsedTemplates[name] = template.Must(template.New(name).Parse(tmpl))
	}

	for {
		//show menu
		fmt.Println("\n\nMnu:")
		fmt.Println("1. Join")
		fmt.Println("2. Get Notification")
		fmt.Println("3. Get Error")
		fmt.Println("4. Exit")
		fmt.Println("Choose an option:")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]interface {
		}

		var tmpl *template.Template

		switch choice {
		case "1":
			tmpl = parsedTemplates["welcome"]
			data = map[string]interface{}{"name": name}

		case "2":
			fmt.Print("Enter your notification message: ")
			notif, _ := reader.ReadString('\n')
			notif = strings.TrimSpace(notif)
			tmpl = parsedTemplates["notification"]
			data = map[string]interface{}{"nm": name, "notif": notif}

		case "3":
			fmt.Print("Enter your error message: ")
			errMsg, _ := reader.ReadString('\n')
			errMsg = strings.TrimSpace(errMsg)
			tmpl = parsedTemplates["error"]
			data = map[string]interface{}{"errorMessage": errMsg}

		case "4":
			fmt.Println("Exitting...")
		default:
			fmt.Println("Invalid choice. please select a valid option.")
			continue

		}

		//render and print tehe tamplate to the console

		err := tmpl.Execute(os.Stdout, data)
		if err != nil {
			fmt.Println("Error Executing template", err)
		}
	}

}
