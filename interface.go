package main

import (
	"fmt"
)

/* INTERFACE
 * Keys to understand interfaces:
 * It is a way to define a set of methods that are shared for multiple objects but with a different outcome.
 * See example below
 */
type Bot interface {
	Login()
	Reply()
}

func doBotThings(b Bot) {
	b.Login()
	b.Reply()
}

type myDiscordBot struct{}

func (db myDiscordBot) Login() {
	fmt.Println("Online on Discord!")
}

func (db myDiscordBot) Reply() {
	fmt.Println("Replying to a discord message! Hi there!")
}

type myChatBot struct{}

func (cb myChatBot) Login() {
	fmt.Println("Online on Chat!")
}

func (db myChatBot) Reply() {
	fmt.Println("Replying to a personal chat! Hi!")
}

// This is the caller function
func run() {
	bot1 := new(myDiscordBot)
	bot2 := new(myChatBot)

	doBotThings(bot1)
	doBotThings(bot2)
}
