package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const port = ":5500"

var columnNumberMap = map[string]int{
	"A": 0,
	"B": 1,
	"C": 2,
}

func main() {
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router := mux.NewRouter()
	router.HandleFunc("/next-move", nextMove).Methods("POST")

	fmt.Println("Serving @ http://localhost" + port)

	log.Fatal(http.ListenAndServe(port, handlers.CORS(header, methods, origins)(router)))
}

func nextMove(w http.ResponseWriter, r *http.Request) {
	var nextMoveRequest NextMoveRequest
	err := json.NewDecoder(r.Body).Decode(&nextMoveRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var cell string = findOptimalMove(nextMoveRequest)
	move := TicTacToeMove{Cell: cell, Player: nextMoveRequest.Player}
	body, err := json.Marshal(move)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("content-type", "applicaton/json")
		w.Header().Set("Access-Control-Allow-Methods", "HEAD, POST, OPTIONS")
		w.Write(body)
	}
}

func convertMovesToBoard(moves []TicTacToeMove) [3][3]string {
	board := [3][3]string{}
	for _, move := range moves {
		colIndex := columnNumberMap[move.Cell[0:1]]
		row := move.Cell[1:2]
		rowNumber, err := strconv.ParseInt(row, 10, 64)
		if err != nil {
			fmt.Println("an error occured when converting moves to board")
		}
		rowIndex := rowNumber - 1
		board[colIndex][rowIndex] = move.Player
	}
	return board
}

func findOptimalMove(nextMoveRequest NextMoveRequest) string {
	alphaArray := [3]string{"A", "B", "C"}
	var board = convertMovesToBoard(nextMoveRequest.Moves)
	colIndex, rowIndex := searchBoardForBestMove(board, nextMoveRequest.Player)
	columnAlpha := alphaArray[colIndex]

	fmt.Printf("%v\n", board)
	return columnAlpha + strconv.Itoa(rowIndex+1)
}

func isWinningCombo(a string, b string, c string) bool {
	if (a != "") && a == b && a == c {
		return true
	}
	return false
}

func getWinner(board [3][3]string) string {

	//check for row win
	for c := 0; c < 3; c++ {
		isRowWin := isWinningCombo(board[c][0], board[c][1], board[c][2])
		if isRowWin {
			return board[c][0]
		}
	}

	//check for column win
	for c := 0; c < 3; c++ {
		for r := 0; r < 3; r++ {
			isColumnWin := isWinningCombo(board[0][r], board[1][r], board[2][r])
			if isColumnWin {
				return board[0][r]
			}
		}
	}

	//check for descending diagnal win
	isDescendingWin := isWinningCombo(board[0][0], board[1][1], board[2][2])
	if isDescendingWin {
		return board[0][0]
	}

	//check for ascending diagnal win
	isAscendingWin := isWinningCombo(board[0][2], board[1][1], board[2][0])
	if isAscendingWin {
		return board[0][2]
	}

	return ""
}

func isGameOver(board [3][3]string) bool {
	winner := getWinner(board)
	if winner != "" {
		return true
	}
	for c := 0; c < 3; c++ {
		for r := 0; r < 3; r++ {
			cellValue := board[c][r]
			if cellValue == "" {
				return false
			}
		}
	}
	return true
}

func minimax(board [3][3]string, depth int, isMax bool, computer string, opponent string) int {
	if isGameOver(board) {
		return getScore(board, depth, computer)
	}
	if isMax {
		bestScore := -1000
		for c := 0; c < 3; c++ {
			for r := 0; r < 3; r++ {
				if board[c][r] == "" {
					board[c][r] = computer
					hypotheticalScore := minimax(board, depth+1, !isMax, computer, opponent)
					if hypotheticalScore > bestScore {
						bestScore = hypotheticalScore
					}
					board[c][r] = "" //undo move
				}
			}
		}
		return bestScore
	} else {
		bestScore := 1000
		for c := 0; c < 3; c++ {
			for r := 0; r < 3; r++ {
				if board[c][r] == "" {
					board[c][r] = opponent
					hypotheticalScore := minimax(board, depth+1, !isMax, computer, opponent)
					if hypotheticalScore < bestScore {
						bestScore = hypotheticalScore
					}
					board[c][r] = "" //undo move
				}
			}
		}
		return bestScore
	}
}

func searchBoardForBestMove(board [3][3]string, computer string) (int, int) {
	bestScore := -1000
	bestColumn := 0
	bestRow := 0
	opponent := "X"

	if computer == "X" {
		opponent = "O"
	}

	for c := 0; c < 3; c++ {
		for r := 0; r < 3; r++ {
			if board[c][r] == "" {
				board[c][r] = computer
				moveScore := minimax(board, 0, false, computer, opponent)
				if moveScore > bestScore {
					bestColumn = c
					bestRow = r
					bestScore = moveScore
				}
				board[c][r] = ""
			}
		}
	}

	return bestColumn, bestRow
}

func getScore(board [3][3]string, depth int, computer string) int {
	winner := getWinner(board)
	if winner == computer {
		return 10 - depth
	}
	if winner != "" && winner != computer {
		return depth - 10
	}
	return 0
}

type TicTacToeMove struct {
	Player string `json:"player"`
	Cell   string `json:"cell"`
}

type NextMoveRequest struct {
	Player string          `json:"player"`
	Moves  []TicTacToeMove `json:"moves"`
}
