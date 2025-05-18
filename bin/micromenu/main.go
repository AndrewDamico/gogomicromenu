package main

import "github.com/yourusername/micromenu"

func main() {
	menu := micromenu.NewMenu(40, 8).
		AddDiv(&micromenu.Div{
			Name:      "Header",
			Type:      micromenu.TitleDiv,
			Priority:  1,
			Location:  micromenu.Top,
			Record:    []micromenu.RecordItem{{Key: "", Value: "Demo Menu"}},
			Decorator: "=",
		}).
		AddDiv(&micromenu.Div{
			Name:     "Main",
			Type:     micromenu.MenuDiv,
			Priority: 2,
			Location: micromenu.Top,
			Record: []micromenu.RecordItem{
				{Key: "1", Value: "Option 1"},
				{Key: "2", Value: "Option 2"},
			},
		})

	menu.Render()
}
