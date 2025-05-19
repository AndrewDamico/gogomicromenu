# Pager Logic in gogomicromenu

## Overview

Paging allows menus to display more items than fit on the screen, using navigation controls like `<` and `>`.

## Key Components

- **PagerState:** Tracks the current page for each paged Div.
- **PagerInfo:** Holds metadata for paging (total pages, current page, page size).
- **PagerDivType:** Special DivType for pager slots, rendered as navigation hints.

## How It Works

- When a Div has more records than available slots, only a subset is shown.
- A pager slot is added, displaying `[<] Prev`, `[>] More`, or `[< >] Prev/Next`.
- User input (`<` or `>`) updates the PagerState, and the menu is re-rendered.

## Example

If a menu has 12 options and only 5 can be shown at once, the user can page through them:
```
[1] Option 1
[2] Option 2
[3] Option 3
[4] Option 4
[>] More
```
Pressing `>` shows the next page.

## Why It's Important

- Keeps the CLI clean and usable for large menus.
- Makes navigation intuitive and efficient.
- Lays the groundwork for more advanced navigation (arrows, search, etc).

---