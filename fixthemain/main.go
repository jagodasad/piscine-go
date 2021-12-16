package main

import "github.com/01-edu/z01"

type Door struct {
	state bool
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func OpenDoor(ptrDoor *Door) bool {
	PrintStr("Door Opening...\n")
	ptrDoor.state = true
	return true
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...\n")
	ptrDoor.state = false
	return true
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door opened ?\n")
	return ptrDoor.state == true
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?\n")
	return ptrDoor.state == false
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == true {
		CloseDoor(door)
	}
}
