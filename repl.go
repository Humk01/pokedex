package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	commandpkg "pokedexcli/commands"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startRepl() {
	cfg := &commandpkg.Config{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		words := cleanInput(input)

		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		cmd, exists := commandpkg.Commands()[commandName]

		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		cfg.Args = words[1:]

		if err := cmd.Callback(cfg); err != nil {
			fmt.Println(err)
		}
	}
}
