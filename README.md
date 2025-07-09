# Features of CLI:
- navigation to the different parts of the system
	- ability to save for quick access
- filter by date, type, or other criteria

# Features of interpreter for the cli:
- interpret more natural/flexibal commands like:
	- `show me all tasks from last week`
	- `goto project <- X in corp` - possible variants are `project`, `proj` `folder`, `dict`, etc
	- `goto project <- last opened` - `<-` command that indicates that we want to insert result from subcommand to the "base of the current command"
	- `goto project <- 1st in 'corp_list'` - we can create lists
		- if we just specify `goto <- work`, then it will go to the parent of the projects, or it will go to the last been in the list
*/

## think:
`<-` is to much of the symbols, `<` is faster, but it is for comparison, how to handle that?

> Why not use AI for that? because AI is not free -> not always available. we will add this later as an option.