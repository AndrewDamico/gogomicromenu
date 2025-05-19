package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AndrewDamico/gogomicromenu"
)

func main() {
	menu := gogomicromenu.NewMenu(40, 8).
		AddDiv(&gogomicromenu.Div{
			Name:      "Header",
			Type:      gogomicromenu.TitleDiv,
			Priority:  1,
			Location:  gogomicromenu.Top,
			Record:    []gogomicromenu.RecordItem{{Key: "", Value: "Demo Menu"}},
			Decorator: "=",
		}).
		AddDiv(&gogomicromenu.Div{
			Name:     "Main",
			Type:     gogomicromenu.MenuDiv,
			Priority: 2,
			Location: gogomicromenu.Top,
			Record: []gogomicromenu.RecordItem{
				{Key: "1", Value: "Option 1"},
				{Key: "2", Value: "Option 2"},
				{Key: "3", Value: "Option 3"},
				{Key: "4", Value: "Option 4"},
				{Key: "5", Value: "Option 5"},
				{Key: "6", Value: "Option 6"},
				{Key: "7", Value: "Option 7"},
				{Key: "8", Value: "Option 8"},
				{Key: "9", Value: "Option 9"},
			},
			Pager: true,
		})

	reader := bufio.NewReader(os.Stdin)
	menu.PageState["Main"] = 0

	for {
		menu.Render()
		fmt.Print("Press < or > to page, q to quit: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "q" {
			break
		}
		if input == ">" {
			menu.PageState["Main"]++
		}
		if input == "<" && menu.PageState["Main"] > 0 {
			menu.PageState["Main"]--
		}
	}
}
