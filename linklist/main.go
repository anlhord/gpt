package main




type foo struct {
	prev *foo
	next *foo
	foo int
}

//////////////////////////////////////////////////////////////
// begin of type specific linklist code

func set(l *foo, i byte, to *foo) {
	if i == Next {
		l.next = to
		return
	}
	l.prev = to
}

func get(l *foo, i byte) (link *foo) {
	if i == Next {
		return l.next
	}
	return l.prev
}
// end of type specific linklist code
//////////////////////////////////////////////////////////////

func main() {
	var x, y, z, w foo
	x = y
	z = w

	Mklist(set, &x)
	
	print(Empty(get, &x))
	Insert(set, get, &x, &y)

	print(Empty(get, &x))
	Insert(set, get, &x, &z)
	

	print(Empty(get, &x))
	Insert(set, get, &x, &w)


	print(Empty(get, &x))
	Remove(set, get, &y)
	
	print("hello")
}

