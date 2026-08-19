package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		user_input := scanner.Text()
		cleaned_user_input := cleanInput(user_input)
		if len(cleaned_user_input) == 0 {
			continue
		}
		cmd, ok := commandRegister[cleaned_user_input[0]]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := cmd.callback(&configuration, cleaned_user_input[1:])
			if err != nil {
				fmt.Printf("error when calling command: %s\n", err)
			}
		}
	}
}