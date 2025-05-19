# micromenu Demo Instructions

## Overview

This demo will show the current capabilities of the `micromenu` package, focusing on:
- Menu struct creation and builder API
- Slot assignment and layout
- Rendering logic (including partial screen clearing)
- Paging for Divs with many records

---

## 1. Basic Menu Creation & Rendering

**Goal:** Show how to define a menu with header and main options, and render it.

**Steps:**
1. Open `main.go` in `micromenu/bin/micromenu/`.
2. Replace the `fmt.Println("Menu struct created...")` line with:
   ```go
   menu.Render()
   ```
3. Run the program:
   ```sh
   go run main.go
   ```
4. Observe the output:  
   - The header should be centered and decorated.
   - The main options should be listed with their keys.
   - The menu fills the specified width and height.

---

## 2. Demonstrate Paging

**Goal:** Show how the menu handles more options than available slots.

**Steps:**
1. In `main.go`, modify the "Main" Div to have more than 6 options and set `Pager: true`:
   ```go
   AddDiv(&micromenu.Div{
       Name:     "Main",
       Type:     micromenu.MenuDiv,
       Priority: 2,
       Location: micromenu.Top,
       Pager:    true,
       Record: []micromenu.RecordItem{
           {Key: "1", Value: "Option 1"},
           {Key: "2", Value: "Option 2"},
           {Key: "3", Value: "Option 3"},
           {Key: "4", Value: "Option 4"},
           {Key: "5", Value: "Option 5"},
           {Key: "6", Value: "Option 6"},
           {Key: "7", Value: "Option 7"},
           {Key: "8", Value: "Option 8"},
       },
   })
   ```
2. In your `main()` function, set the page state and render:
   ```go
   menu.PageState["Main"] = 0 // First page
   menu.Render()
   menu.PageState["Main"] = 1 // Next page
   menu.Render()
   ```
3. Run the program and observe:
   - Only a subset of options are shown per page.
   - A pager line (`[>] More`, `[<] Prev`, or `[< >] Prev/Next`) appears as needed.

---

## 3. Partial Screen Clearing

**Goal:** Show that only the menu area is cleared/redrawn between renders.

**Steps:**
1. Add a `fmt.Println("Before menu")` before `menu.Render()` and `fmt.Println("After menu")` after.
2. Run the program and observe:
   - The lines before and after the menu remain, only the menu area is cleared/redrawn.

---

## 4. Highlighted Features

- **Builder API:** Easily chain `.AddDiv()` calls to construct complex menus.
- **Slot Assignment:** Divs are assigned to top/bottom regions by priority, with paging support.
- **Partial Screen Clearing:** Only the menu area is cleared, preserving other terminal output.
- **Paging:** Large menus are automatically paged, with navigation hints.

---

## 5. (Optional) Experiment

- Try adding a bottom Div (e.g., a fence or input prompt) and see how it appears.
- Change the menu width/height and observe the layout adapt.

---

## Summary

- The demo shows menu creation, rendering, paging, and partial screen clearing.
- The code is modular and ready for further extension (input handling, dynamic updates, etc.).

---