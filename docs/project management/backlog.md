# CLI Menu System: Backlog

## 1. Core Types & Structs
- [x] 1.1. Define `Div`, `RecordItem`, `DivType`, and `Location` types to represent menu regions.
- [x] 1.2. Each `Div` has: name, type, priority, location (top/bottom), records (key/value pairs), decorator, slot count, identifier, pager flag.
- [x] 1.3. Define a `Menu` struct with configuration (width, height) and a slice of `Div`s.

## 2. Slot Assignment & Layout Algorithm
- [x] 2.1. Menu has a fixed number of slots (height), each slot is a line in the terminal.
- [x] 2.2. On each render, sort `Div`s by priority (ascending).
- [x] 2.3. Assign slots to `Div`s by location: top-down for `DivTop`, bottom-up for `DivBottom`.
- [x] 2.4. Each `Div` claims as many slots as its records require, plus one for a pager if needed.
- [x] 2.5. If not enough slots, lower-priority `Div`s are dropped or paged.
- [x] 2.6. Build a slice of slots, each holding a reference to its `Div`, the record, and the index.

## 3. Rendering Logic
- [x] 3.1. Initialize a blank canvas of `height` lines, each of `width` characters.
- [x] 3.2. For each slot, inject content from its `Div` and record, formatting with decorators, selection keys, and padding.
- [x] 3.3. Print the menu line-by-line to stdout.
- [x] 3.4. Support partial screen clearing: before each render, clear only the menu area (not the whole terminal).

## 4. Paging
- [x] 4.1. If a `Div`’s records exceed its assigned slots, show a subset and add a pager slot (with `<`, `>`, or `< >` as needed).
- [x] 4.2. Pager is always at the bottom of its region and is not selectable.

## 5. Menu Interaction
- [ ] 5.1. Each selectable item is assigned a number or key, displayed in the menu.
- [ ] 5.2. User input is mapped to the correct action or submenu based on the slot’s identifier.
- [ ] 5.3. After each action, the menu is re-laid out and re-rendered, updating only the menu area.

## 6. Builder API & Extensibility
- [x] 6.1. Provide a fluent builder API for constructing menus and adding `Div`s.
- [ ] 6.2. Allow configuration of decorators, padding, and ANSI colors for styling.
- [ ] 6.3. Support for dynamic menus: easily add, remove, or update `Div`s and records at runtime.

## 7. Robustness & Testing
- [ ] 7.1. Handle edge cases: overflow, empty slots, empty menus, and invalid input.
- [ ] 7.2. Provide unit tests for the layout algorithm and rendering logic.
- [ ] 7.3. Ensure error handling for all I/O and rendering operations.

## 8. Documentation & Comments
- [x] 8.1. Each type and method is documented with clear comments.
- [ ] 8.2. Example usage and extension points are provided in the code.

## 9. Advanced Features (Optional/Future)
- [ ] 9.1. Support for ANSI color profiles/themes.
- [ ] 9.2. Support for alternative navigation keys (arrows, tab, etc.).
- [ ] 9.3. Support for menu state/history (breadcrumbs, back/forward).
- [ ] 9.4. Support for localization/internationalization of menu text.
- [ ] 9.5. Efficient handling of very large menus (lazy loading, etc.).
- [ ] 9.6. Generate API documentation (GoDoc/Markdown).

---

## 5. Menu Interaction
- [ ] 5.1. Each selectable item is assigned a number or key, displayed in the menu.
- [ ] 5.2. User input is mapped to the correct action or submenu based on the slot’s identifier.
- [ ] 5.3. After each action, the menu is re-laid out and re-rendered, updating only the menu area.

## 6. Builder API & Extensibility
- [ ] 6.2. Allow configuration of decorators, padding, and ANSI colors for styling.
- [ ] 6.3. Support for dynamic menus: easily add, remove, or update `Div`s and records at runtime.

## 7. Robustness & Testing
- [ ] 7.1. Handle edge cases: overflow, empty slots, empty menus, and invalid input.
- [ ] 7.2. Provide unit tests for the layout algorithm and rendering logic.
- [ ] 7.3. Ensure error handling for all I/O and rendering operations.

## 8. Documentation & Comments
- [ ] 8.2. Example usage and extension points are provided in the code.

## 9. Advanced Features (Optional/Future)
- [ ] 9.1. Support for ANSI color profiles/themes.
- [ ] 9.2. Support for alternative navigation keys (arrows, tab, etc.).
- [ ] 9.3. Support for menu state/history (breadcrumbs, back/forward).
- [ ] 9.4. Support for localization/internationalization of menu text.
- [ ] 9.5. Efficient handling of very large menus (lazy loading, etc.).
- [ ] 9.6. Generate API documentation (GoDoc/Markdown).