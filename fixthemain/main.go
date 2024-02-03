package main

import "github.com/01-edu/z01"

type Door struct {
	state bool
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...")
	ptrDoor.state = true
}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...")
	ptrDoor.state = false
}

func IsDoorOpen(door Door) bool {
	PrintStr("is the Door opened ?")
	return door.state
}

func IsDoorClose(door Door) bool {
	PrintStr("is the Door closed ?")
	return !door.state
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(*door) {
		OpenDoor(door)
	}
	if IsDoorOpen(*door) {
		CloseDoor(door)
	}
	if door.state {
		CloseDoor(door)
	}
}
