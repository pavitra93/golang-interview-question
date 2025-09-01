package main

import (
	"fmt"

	"github.com/pavitra93/trello/enums/privacy"
	"github.com/pavitra93/trello/repository"
	"github.com/pavitra93/trello/service"
)

func main() {
	fmt.Println("Welcome to Trello App")
	fmt.Println("===== User Operations =====")
	UserRepository := repository.NewUserRepository()
	userService := service.UserService{UserRepository: UserRepository}

	U1, _ := userService.AddUser("Pavitra Mehta", "pavitramehta4@gmail.com")
	U2, _ := userService.AddUser("Monika Sharma", "monika@gmail.com")
	U3, _ := userService.AddUser("Sonika Sharma", "sonika@gmail.com")

	userService.SingleUser(U1)
	userService.SingleUser(U2)
	userService.SingleUser(U3)

	fmt.Println("===== Boards Operations =====")

	BoardRepository := repository.NewBoardRepository()
	boardService := service.BoardService{BoardRepository: BoardRepository}

	B1, _ := boardService.CreateBoard("Board 1", "Sprint 1 Board 1", privacy.PUBLIC)
	B2, _ := boardService.CreateBoard("Board 2", "Sprint 1 Board 2", privacy.PUBLIC)

	boardService.SingleBoard(B1)
	boardService.SingleBoard(B2)
	//boards, err := boardService.ListBoards()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for _, board := range boards {
	//	fmt.Println(board)
	//}

	fmt.Println("===== Lists Operations =====")

	ListRepository := repository.NewListRepository()
	listService := service.ListService{ListRepository: ListRepository}

	L1, _ := listService.CreateList("List 1")
	L2, _ := listService.CreateList("List 2")

	listService.SingleList(L1)
	listService.SingleList(L2)
	//boards, err := boardService.ListBoards()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for _, board := range boards {
	//	fmt.Println(board)
	//}

	fmt.Println("===== Cards Operations =====")

	CardRepository := repository.NewCardRepository()
	cardService := service.CardService{
		CardRepository:  CardRepository,
		BoardRepository: BoardRepository,
		ListRepository:  ListRepository,
	}

	C1, _ := cardService.CreateCard("Card 1", "Card 2 Desc")
	C2, _ := cardService.CreateCard("Card 2", "Card 2 Desc")

	cardService.SingleCard(C1)
	cardService.SingleCard(C2)
	//boards, err := boardService.ListBoards()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for _, board := range boards {
	//	fmt.Println(board)
	//}

	fmt.Println("===== Assignment Operations =====")

	boardService.AddUsersByBoardID(B1, []int{U1, U2, U3})
	boardService.AddListsByBoardID(B1, []int{L1, L2})

	listService.AddCardsByListID(L1, []int{C1, C2})

	cardService.AssignUserByCardID(C1, U1)
	cardService.AssignUserByCardID(C2, U2)

	fmt.Println("===== Assignment Prints Operations =====")

	boardService.SingleBoard(B1)
	listService.SingleList(L1)
	cardService.SingleCard(C1)
	cardService.SingleCard(C2)
	userService.SingleUser(U1)
	userService.SingleUser(U2)

	fmt.Println("===== Unassigment Operations =====")

	cardService.UnAssignUserByCardID(C1, U3)
	cardService.SingleCard(C1)

	fmt.Println("===== Card Movement Operations =====")
	cardService.MoveCardAcrossListInSameBoard(B1, L1, L2, C2)
	listService.SingleList(L1)
	listService.SingleList(L2)
}
