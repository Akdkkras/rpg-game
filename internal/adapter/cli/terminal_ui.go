package cli

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Akdkkras/rpg-game/internal/entity"
)

type TerminalUI struct {
	reader *bufio.Reader
}

func NewTerminalUI() *TerminalUI {
	return &TerminalUI{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (ui *TerminalUI) PrintGameStart() {
	fmt.Println("=== ДОБРО ПОЖАЛОВАТЬ В ТЕКСТОВУЮ RPG ИГРУ ===\n")
}

func (ui *TerminalUI) PrintGameEnd() {
	fmt.Println("🎉 ПОЗДРАВЛЯЕМ! Вы прошли игру! 🎉")
}

func (ui *TerminalUI) PrintLocationInfo(location *entity.Location) {
	fmt.Println("========================================")
	fmt.Printf("📍 ЛОКАЦИЯ: %s\n", location.Title)
	fmt.Println("========================================")
	fmt.Printf("%s\n\n", location.Description)
	fmt.Println("Доступные квесты:")
	for _, quest := range location.Quests {
		fmt.Printf("  - %s\n", quest.Title)
	}
	fmt.Println()
}

func (ui *TerminalUI) ChooseFromSuggested() string {
	fmt.Println("Доступные команды:")
	fmt.Println("  1. следующая локация")
	fmt.Println("  2. выполнить квест")
	fmt.Println("  3. завершить игру")
	fmt.Print("\n> ")

	command, _ := ui.reader.ReadString('\n')
	command = strings.TrimSpace(strings.ToLower(command))

	return command
}

func (ui *TerminalUI) PrintMoveOn() {
	fmt.Println("\n✨ Вы отправляетесь дальше... ✨\n")
}

func (ui *TerminalUI) ChooseQuestFromLocation(location *entity.Location) (int, error) {
	fmt.Println("\n⚔️ Выберите квест для прохождения: ⚔️\n")
	for idx, quest := range location.Quests {
		fmt.Printf("  %d. %s\n", idx+1, quest.Title)
	}
	fmt.Print("\n> ")

	questNumberString, _ := ui.reader.ReadString('\n')
	questNumberString = strings.TrimSpace(questNumberString)
	questNumber, err := strconv.Atoi(questNumberString)
	if err != nil {
		return 0, errors.New("необходимо ввести число")
	}

	questIdx := questNumber - 1
	if questIdx < 0 || questIdx > len(location.Quests) {
		return 0, errors.New("нет соответствующего квеста")
	}

	return questIdx, nil
}

func (ui *TerminalUI) PrintWrongChoice() {
	fmt.Println("❌ Некорректный ввод. Попробуйте снова.")
}

func (ui *TerminalUI) PrintWrongCommand() {
	fmt.Println("\n❌ Неизвестная команда. Попробуйте снова.\n")
}

func (ui *TerminalUI) PrintGameExit() {
	fmt.Println("\n👋 Спасибо за игру! До свидания!")
}

//func (ui *TerminalUI) Show(text string) {
//	fmt.Println(text)
//}

//func (ui *TerminalUI) ChooseAction(actions []string) (int, error) {
//	for i, a := range actions {
//		fmt.Printf("%d. %s\n", i+1, a)
//	}
//
//	fmt.Print("> ")
//
//	input, _ := ui.reader.ReadString('\n')
//	input = strings.TrimSpace(input)
//
//	n, err := strconv.Atoi(input)
//	if err != nil {
//		return -1, err
//	}
//
//	return n - 1, nil
//}
