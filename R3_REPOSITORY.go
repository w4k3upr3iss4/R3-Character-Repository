package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		showWelcome()
	}
}

func showWelcome() {
	clearScreen()

	fmt.Println(`

░▒▓███████▓▒░  ░▒▓███████▓▒░  
░▒▓█▓▒░░▒▓█▓▒░        ░▒▓█▓▒░ 
░▒▓█▓▒░░▒▓█▓▒░        ░▒▓█▓▒░ 
░▒▓███████▓▒░  ░▒▓███████▓▒░  
░▒▓█▓▒░░▒▓█▓▒░        ░▒▓█▓▒░ 
░▒▓█▓▒░░▒▓█▓▒░        ░▒▓█▓▒░ 
░▒▓█▓▒░░▒▓█▓▒░ ░▒▓███████▓▒░  
                              
Добро пожаловать в репозиторий R3N3G4T3.
Хранилище персонажей, воспоминаний и хронологий. 
Ты всегда можешь обратиться сюда, чтобы почитать про все, что тебе как игроку может быть интересно
Заходя, вытирай ноги и складывай оружие`)
	fmt.Println()
	fmt.Println("Выбери опцию:")
	fmt.Println()
	fmt.Println("1. Просмотр персонажа")
	fmt.Println("2. Просмотр хронологии событий сюжета")
	fmt.Println("3. Выход")

	switch readChoice() {
	case "1":
		showCharacterMenu()
	case "2":
		fmt.Println("\n[Заглушка] Переход к хронологии событий...")
		waitEnter()
	case "3":
		fmt.Println("\nЗавершение работы. До встречи.")
		os.Exit(0)
	default:
		fmt.Println("\nНеверный ввод. Пиши то, что сказано выше. Никакой отсебятины.")
		waitEnter()
	}
}

func readChoice() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(">> ")
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func waitEnter() {
	fmt.Print("\nНажмите Enter, чтобы продолжить...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func clearScreen() {
	// работает в Windows (примерно), не идеально, но чистит
	fmt.Print("\033[H\033[2J")
}

func showCharacterMenu() {
	clearScreen()
	fmt.Println(`
	
 ░▒▓██████▓▒░░▒▓█▓▒░░▒▓█▓▒░░▒▓██████▓▒░ ░▒▓██████▓▒░ ░▒▓███████▓▒░▒▓████████▓▒░ 
░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░        
░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░        
░▒▓█▓▒░      ░▒▓████████▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░░▒▓██████▓▒░░▒▓██████▓▒░   
░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░      ░▒▓█▓▒░▒▓█▓▒░        
░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░      ░▒▓█▓▒░▒▓█▓▒░        
 ░▒▓██████▓▒░░▒▓█▓▒░░▒▓█▓▒░░▒▓██████▓▒░ ░▒▓██████▓▒░░▒▓███████▓▒░░▒▓████████▓▒░ 
                                                                                
                                                                                `)
	fmt.Println()
	fmt.Println("Выбери игру, откуда хочешь взять персонажа:")
	fmt.Println("1. The Brink")
	fmt.Println("2. Cyberpunk RED")
	fmt.Println("3. Tennessee Chronicles (Вестерн)")
	fmt.Println("4. Назад")

	switch readChoice() {
	case "1":
		clearScreen()
		fmt.Println(`
▄▄▄█████▓ ██░ ██ ▓█████     ▄▄▄▄    ██▀███   ██▓ ███▄    █  ██ ▄█▀
▓  ██▒ ▓▒▓██░ ██▒▓█   ▀    ▓█████▄ ▓██ ▒ ██▒▓██▒ ██ ▀█   █  ██▄█▒ 
▒ ▓██░ ▒░▒██▀▀██░▒███      ▒██▒ ▄██▓██ ░▄█ ▒▒██▒▓██  ▀█ ██▒▓███▄░ 
░ ▓██▓ ░ ░▓█ ░██ ▒▓█  ▄    ▒██░█▀  ▒██▀▀█▄  ░██░▓██▒  ▐▌██▒▓██ █▄ 
  ▒██▒ ░ ░▓█▒░██▓░▒████▒   ░▓█  ▀█▓░██▓ ▒██▒░██░▒██░   ▓██░▒██▒ █▄
  ▒ ░░    ▒ ░░▒░▒░░ ▒░ ░   ░▒▓███▀▒░ ▒▓ ░▒▓░░▓  ░ ▒░   ▒ ▒ ▒ ▒▒ ▓▒
    ░     ▒ ░▒░ ░ ░ ░  ░   ▒░▒   ░   ░▒ ░ ▒░ ▒ ░░ ░░   ░ ▒░░ ░▒ ▒░
  ░       ░  ░░ ░   ░       ░    ░   ░░   ░  ▒ ░   ░   ░ ░ ░ ░░ ░ 
          ░  ░  ░   ░  ░    ░         ░      ░           ░ ░  ░   
                                 ░                                `)
		fmt.Println()
		fmt.Println("1. Скол")
		fmt.Println("2. Назад")
		switch readChoice() {
		case "1":
			clearScreen()
			fmt.Println(`Скол. Основатель THE FACTOR UNIT. По неподтвержденной информации беженец из ОДК.
Основная деятельность:
- Информационные диверсии: сливы данных, сбор компромата, шантаж мировых элит
- Логистические и стратегические диверсии: уничтожение путей транспортировок, 
убийства членов командования
Вооружен Mk18 и Glock19
Уставная форма отсуствует.
Отличительный знак - соответсвующий шеврон TFU на левом плече
Не выбирает сторону. Работает против всех крупных группировок.
Преследуемые цели неизвестны.
Контакта избегать. ЛЮБОЙ ЦЕНОЙ`)
			waitEnter()
		case "2":
			// ничего — просто вернёмся
		}
	case "2":
		clearScreen()
		fmt.Println(`
		
._______ ._______      .______  ._______.______  
:_.  ___\: ____  |     : __   \ : .____/:_ _   \ 
|  : |/\ |    :  |     |  \____|| : _/\ |   |   |
|    /  \|   |___|     |   :  \ |   /  \| . |   |
|. _____/|___|         |   |___\|_.: __/|. ____/ 
 :/                    |___|       :/    :/      
 :                                       :        
                                                 
		`)
		waitEnter()
	case "3":
		clearScreen()
		fmt.Println(` _____                                      
(_   _)                                     
  | | ___  ____  ____  ___  __  __ ___  ___ 
  ( )( o_)( __ )( __ )( o_)(_' (_'( o_)( o_)
  /_\ \(  /_\/_\/_\/_\ \(  /__)/__)\(   \(  
                                            
 ___  _                   _     _           
(  _)( )                 (_)   ( )          
| |  | |_  __  ___  ____  _  __| | ___  __  
( )_ ( _ )( _)( o )( __ )( )/ /( )( o_)(_'  
/___\/_\||/_\  \_/ /_\/_\/_\\_\/_\ \(  /__) `)
		fmt.Println()
		fmt.Println("Выбери персонажа:")
		fmt.Println("1. Рик Граймс")
		fmt.Println("2. Гуд Мосли")
		fmt.Println("3. Кид Харпер")
		fmt.Println("0. Назад")
		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			showCharacterFromFile("data/TennesseeChronicles/playableChars/rickGrimes.txt")
		case "2":
			showCharacterFromFile("data/TennesseeChronicles/playableChars/goodMos.txt")
		case "3":
			showCharacterFromFile("data/TennesseeChronicles/playableChars/eliHarper.txt")
		case "0":
			return
		default:
			fmt.Println("Неверный выбор.")
		}
		waitEnter()
	case "4":
		// ничего — просто вернёмся
	default:
		fmt.Println("Неверный ввод.")
		waitEnter()
	}
}
func showCharacterFromFile(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	fmt.Println("\n==============================")
	fmt.Println(string(content))
	fmt.Println("==============================")

}
