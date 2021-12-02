package main

import (
	"fmt"

	"github.com/jinjio/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	err := dictionary.Update("tt", "Second")
	if err != nil {
		fmt.Println(err)
	}
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)
}

// import (
// 	"fmt"

// 	"github.com/jinjio/learngo/accounts"
// )

// func main() {
// 	account := accounts.NewAccount("jio")
// 	account.Deposit(10)
// 	err := account.Withdraw(20)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(account.Balance(), account.Owner())
// }
