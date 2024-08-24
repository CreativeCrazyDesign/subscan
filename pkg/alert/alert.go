package alert

import "fmt"

func Error(text string) {
	fmt.Println(
		fmt.Sprint(RED, "[-] ", COLOR_RESET, text),
	)
}

func Warn(text string) {
	fmt.Println(
		fmt.Sprint(CYAN, "[!] ", COLOR_RESET, text),
	)
}

func Success(text string) {
	fmt.Println(
		fmt.Sprint(GREEN, "[+] ", COLOR_RESET, text),
	)
}

func Info(text string) {
	fmt.Println(
		fmt.Sprint(BLUE, "[?] ", COLOR_RESET, text),
	)
}

func Impo(text string) {
	fmt.Println(
		fmt.Sprint(YELLOW, "[*] ", COLOR_RESET, text),
	)
}
