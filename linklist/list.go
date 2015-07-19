package main

const Next = byte(0)
const Prev = byte(1)

/*
//CALLBACK SIGNATURES:
set func(*, byte, *)
get func(*, byte) (*)
*/

// creates a new list
func Mklist(set func(*, byte, *), list *) {
	set(list, Prev, list)
	set(list, Next, list)
}

// inserts item to list
func Insert(set func(*, byte, *), get func(*, byte)(*), list *, elm *) {
	set(elm, Prev, list)
	set(elm, Next, get(list, Next))
	set(list, Next, elm)
	set(get(elm, Next), Prev, elm)
}

// remove item from list.
func Remove(set func(*, byte, *), get func(*, byte)(*), elm *) {
	set(get(elm, Prev), Next, get(elm, Next))
	set(get(elm, Next), Prev, get(elm, Prev))
}

// is list empty?
func Empty(get func(*, byte)(*), list *) bool {
	return get(list, Next) == list;
}

// do count of the list items
func Len(get func(*, byte)(*), list *) (count int) {
	e := get(list, Next)
	for e != list {
		e = get(e, Next)
		count++;
	}
	return count
}



