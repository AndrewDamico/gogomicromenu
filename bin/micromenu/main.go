package main

import "github.com/AndrewDamico/gogomicromenu"

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
			},
		})

	menu.Render()
}
