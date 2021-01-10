package main

import "fmt"

// tictactoe driver
func main() {
    // initalization
    game := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
    placeholder := [9]string{"0️⃣", "1️⃣", "2️⃣", "3️⃣", "4️⃣", "5️⃣", "6️⃣", "7️⃣", "8️⃣"}
    icons := [2]string{"❌", "⭕"}
    loops := 0
    player := 2
    startscreen()

    // game loop
    for {
        display(game, placeholder, icons)

        // swaps between player 1 and 2
        switch player {
        case 2:
            player = 1
        default:
            player = 2
        }

        fmt.Printf("%s  pick a position: ", icons[player-1])
        pos := input()

        // quits program
        if pos == 9 {
            break
        }

        game = play(player, game, pos)
        winner := winner(game)

        if winner != 0 {
            fmt.Printf("%s  is the winner!\n", icons[winner-1])
        } else if loops == 9 {
            fmt.Println("!NO WINNER TIED GAME!")
        }

        // reset game if a tie or a winner
        if loops == 9 || winner != 0 {
            display(game, placeholder, icons)
            fmt.Println("+ Game Reset +")
            fmt.Println("")
            loops = 0
            game = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
        }
        loops++
    }
}

// game instructions
func startscreen() {
    fmt.Println("+++++ Instructions +++++")
    fmt.Println("========================")
    fmt.Println("- The numbers on screen")
    fmt.Println("  correspond to the")
    fmt.Println("  position your piece")
    fmt.Println("  will be placed.")
    fmt.Println("")
    fmt.Println("- Enter 9 to Quit.")
    fmt.Println("")
}

// displays the game
func display(game [9]int, placeholder [9]string, icons [2]string) {
    for i, v := range game {
        if v == 0 {
            fmt.Print(placeholder[i])
        } else {
            fmt.Print(icons[v-1])
        }
        if i%3 == 2 {
            fmt.Print("\n")
        } else {
            fmt.Print("  ")
        }
    }
    fmt.Println("")
}

// collect user input
func input() int {
    for {
        var pos int
        fmt.Scan(&pos)
        // valid inputs between 0-9
        if (pos < 10) && (pos >= 0) {
            return pos
        } else {
            fmt.Print("Pick a position 0-8 or 9 to quit: ")
        }
    }
}

// executes players move
func play(player int, game [9]int, pos int) [9]int {
    if game[pos] != 0 {
        fmt.Print("Position occupied. Try again: ")
        pos = input()
        game = play(player, game, pos)
    } else {
        game[pos] = player
    }
    return game
}

// determines if board contains a winner
func winner(game [9]int) int {
    x := 0
    // checks rows
    for i := 0; i <= 6; i += 3 {
        if (game[i] == game[i+1]) && (game[i+1] == game[i+2]) {
            x += game[i]
        }
    }
    // checks columns
    for i := 0; i <= 2; i++ {
        if (game[i] == game[i+3]) && (game[i+3] == game[i+6]) {
            x += game[i]
        }
    }
    // checks diagonals
    for i := 0; i <= 2; i += 2 {
        if (game[i] == game[4]) && (game[4] == game[8-i]) {
            // takes care of special case were player scores 2 diagonals
            if x == 0 {
                x += game[i]
            }
        }
    }
    return x
}
