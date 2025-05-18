package main

import (
	"fmt"
	"strings"
)

// Location designates Top or Bottom region placement.
type Location int

const (
	Top Location = iota
	Bottom
)

// DivType categorizes the div's purpose in the menu.
type DivType int

const (
	TitleDiv DivType = iota
	InputDiv
	MenuDiv
	FenceDiv
	PagerDiv
)

// RecordItem pairs a display label with a value.
type RecordItem struct {
	Key   string
	Value string
}

// Div defines a block that occupies slots, with metadata for layout.
type Div struct {
	Name       string
	Type       DivType
	Priority   int
	Location   Location
	Record     []RecordItem
	Decorator  string
	Identifier []string
	Pager      bool
}

// Slot holds a reference to a Div and a specific RecordItem for rendering.
type Slot struct {
	Div   *Div
	Item  RecordItem
	Index int // Index within the Div's Record
}

// Menu orchestrates Divs in a fixed grid of width×height.
type Menu struct {
	Width, Height int
	Divs          []*Div
	PageState     map[string]int // Div.Name -> current page
}

// NewMenu constructs a menu with given dimensions.
func NewMenu(width, height int) *Menu {
	return &Menu{Width: width, Height: height, PageState: make(map[string]int)}
}

// AddDiv appends a div to the menu for later layout.
func (m *Menu) AddDiv(d *Div) *Menu {
	m.Divs = append(m.Divs, d)
	return m
}

// Layout assigns Div records to menu slots based on priority and location, with paging support.
func (m *Menu) Layout() []Slot {
	slots := make([]Slot, m.Height)
	used := make([]bool, m.Height)

	var topDivs, bottomDivs []*Div
	for _, div := range m.Divs {
		if div.Location == Top {
			topDivs = append(topDivs, div)
		} else {
			bottomDivs = append(bottomDivs, div)
		}
	}
	sortDivs := func(divs []*Div) []*Div {
		out := make([]*Div, len(divs))
		copy(out, divs)
		for i := 0; i < len(out)-1; i++ {
			for j := i + 1; j < len(out); j++ {
				if out[i].Priority > out[j].Priority {
					out[i], out[j] = out[j], out[i]
				}
			}
		}
		return out
	}

	// Helper for paged divs
	assignPaged := func(div *Div, slotIdx, maxSlots int, reverse bool) int {
		page := m.PageState[div.Name]
		total := len(div.Record)
		pageSize := maxSlots
		totalPages := (total + pageSize - 1) / pageSize
		if page < 0 {
			page = 0
		}
		if page >= totalPages {
			page = totalPages - 1
		}
		m.PageState[div.Name] = page
		start := page * pageSize
		end := start + pageSize
		if end > total {
			end = total
		}
		records := div.Record[start:end]
		for i, rec := range records {
			idx := slotIdx
			if reverse {
				idx = slotIdx - i
			} else {
				idx = slotIdx + i
			}
			if idx >= 0 && idx < m.Height {
				slots[idx] = Slot{Div: div, Item: rec, Index: start + i}
				used[idx] = true
			}
		}
		// Add pager slot if needed
		if totalPages > 1 {
			pagerText := ""
			switch {
			case page == 0:
				pagerText = "[>] More"
			case page == totalPages-1:
				pagerText = "[<] Prev"
			default:
				pagerText = "[< >] Prev/Next"
			}
			pagerSlot := Slot{
				Div:  &Div{Name: div.Name + "_pager", Type: PagerDiv},
				Item: RecordItem{Key: "", Value: pagerText},
			}
			pagerIdx := slotIdx + len(records)
			if reverse {
				pagerIdx = slotIdx - len(records)
			}
			if pagerIdx >= 0 && pagerIdx < m.Height {
				slots[pagerIdx] = pagerSlot
				used[pagerIdx] = true
			}
		}
		return len(records) + 1 // records + pager
	}

	// Assign top Divs from top down
	slotIdx := 0
	for _, div := range sortDivs(topDivs) {
		if div.Pager && len(div.Record) > (m.Height-slotIdx) {
			usedSlots := assignPaged(div, slotIdx, m.Height-slotIdx-1, false)
			slotIdx += usedSlots
		} else {
			for i, rec := range div.Record {
				if slotIdx < m.Height {
					slots[slotIdx] = Slot{Div: div, Item: rec, Index: i}
					used[slotIdx] = true
					slotIdx++
				}
			}
		}
	}

	// Assign bottom Divs from bottom up
	bottomIdx := m.Height - 1
	for i := len(bottomDivs) - 1; i >= 0; i-- {
		div := bottomDivs[i]
		if div.Pager && len(div.Record) > (bottomIdx+1) {
			usedSlots := assignPaged(div, bottomIdx, bottomIdx+1, true)
			bottomIdx -= usedSlots
		} else {
			for j := len(div.Record) - 1; j >= 0; j-- {
				rec := div.Record[j]
				for bottomIdx >= 0 && used[bottomIdx] {
					bottomIdx--
				}
				if bottomIdx >= 0 {
					slots[bottomIdx] = Slot{Div: div, Item: rec, Index: j}
					used[bottomIdx] = true
					bottomIdx--
				}
			}
		}
	}

	return slots
}

// Render prints the menu to the terminal using the current layout.
// It respects width, decorators, and fills unused slots with blanks.
func (m *Menu) Render() {
	m.ClearMenuArea()
	slots := m.Layout()
	for i := 0; i < m.Height; i++ {
		slot := slots[i]
		line := ""
		if slot.Div != nil {
			switch slot.Div.Type {
			case TitleDiv:
				// Center the title with decorators
				title := fmt.Sprintf(" %s ", slot.Item.Value)
				pad := (m.Width - len(title)) / 2
				if pad < 0 {
					pad = 0
				}
				line = strings.Repeat("=", pad) + title + strings.Repeat("=", m.Width-len(title)-pad)
			case FenceDiv:
				if slot.Div.Decorator != "" {
					line = slot.Div.Decorator
				} else {
					line = strings.Repeat("-", m.Width)
				}
			case MenuDiv, PagerDiv:
				line = fmt.Sprintf("[%s] %s", slot.Item.Key, slot.Item.Value)
			case InputDiv:
				line = slot.Item.Value
			default:
				line = slot.Item.Value
			}
		}
		// Pad or trim to width
		if len(line) > m.Width {
			line = line[:m.Width]
		} else if len(line) < m.Width {
			line = line + strings.Repeat(" ", m.Width-len(line))
		}
		fmt.Println(line)
	}
}

// ClearMenuArea clears only the menu area in the terminal (partial screen clearing).
func (m *Menu) ClearMenuArea() {
	for i := 0; i < m.Height; i++ {
		fmt.Print("\033[F") // Move cursor up one line
	}
	for i := 0; i < m.Height; i++ {
		fmt.Print("\033[2K") // Clear entire line
		fmt.Print("\033[E")  // Move cursor down one line
	}
	for i := 0; i < m.Height; i++ {
		fmt.Print("\033[F") // Move cursor up one line to return to menu start
	}
}
