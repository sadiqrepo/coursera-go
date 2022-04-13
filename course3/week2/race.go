package main

import "time"

/** Execute with race condition checker "go run -race ."

This is an example of race condition
2 goroutines tries to read&write sharedInt and there is no access control.
A race condition occurs when multiple threads try to access and modify the same data (memory address).
E.g., if one thread tries to increase an integer and another thread tries to read it,
this will cause a race condition. On the other hand, there won't be a race condition,
if the variable is read-only.
*/

var sharedInt int = 0
var unusedValue int = 0

func reader() {
	for {
		var val int = sharedInt
		if val%10 == 0 {
			unusedValue = unusedValue + 1
		}
	}
}

func writer() {
	for {
		sharedInt = sharedInt + 1
	}
}

func main() {
	go reader()
	go writer()
	time.Sleep(10 * time.Second)
}
