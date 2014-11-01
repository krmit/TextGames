package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/krmit/TextGames/TextGamesLib"
	"os"
	"strconv"
	"strings"
)

func main() {

	max_number_of_game := 0
	flag.IntVar(&max_number_of_game, "n", 3, "Number of game to play in a row.")
	flag.Parse()

	fmt.Println("Lets play Rock-Paper-Scissors")
	wins, counter := 0.0, 0

	first_chooser := textgameslib.NewChooser()
	first_chooser.Add(textgameslib.Option{Id: 1, Description: "Rock!", Flags: []string{"r", "R"}})
	first_chooser.Add(textgameslib.Option{Id: 2, Description: "Paper!", Flags: []string{"p", "P"}})
	first_chooser.Add(textgameslib.Option{Id: 3, Description: "Scissors!", Flags: []string{"s", "S"}})

	var player_choose textgameslib.Option
	var err error

	for ; counter < max_number_of_game; counter++ {

		not_a_good_chooise := true
		for not_a_good_chooise {

			fmt.Print(first_chooser.Ask())
			fmt.Print("Enter text: ")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')

			text = strings.Trim(text, " \n")

			player_choose, err = first_chooser.Answer(text)

			not_a_good_chooise = false
			if err != nil {
				fmt.Println("Bad chooise!")
				not_a_good_chooise = true
			}
		}

		ai := first_chooser.RandOption()
		fmt.Println("AI choose: " + ai.String())
		fmt.Println("You have choosed: " + player_choose.String())

		if ai.Id == 3 && player_choose.Id == 1 || ai.Id < player_choose.Id {
			wins++
			fmt.Println("You Win!")
		} else if ai.Id == player_choose.Id {
			wins += 0.5
			fmt.Println("You get a Draw!")
		} else {
			fmt.Println("You Loose!")
		}
		fmt.Println("You have win " + strconv.FormatFloat(wins, 'f', 1, 64) + " of " + strconv.Itoa(counter+1) + ".")
	}

	fmt.Println()
	fmt.Println("Game over")
	fmt.Println()
	fmt.Println("You have win " + strconv.FormatFloat(wins, 'f', 1, 64) + " of " + strconv.Itoa(counter) + ".")
	fmt.Println("Your winrate: " + strconv.Itoa(int(wins*100)/counter) + "%.")
}
