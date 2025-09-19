Set implementation in Go

```go
package main

import (
	"fmt"

	"github.com/akthrmsx/myset"
)

func main() {
	books := myset.New("Harry Potter and the Philosopher's Stone")

	books.Add("Harry Potter and the Chamber of Secrets")
	books.Add("Harry Potter and the Prisoner of Azkaban")

	if !books.Has("Harry Potter and the Goblet of Fire") {
		fmt.Printf("I have %d books, but I don't have 'Harry Potter and the Goblet of Fire'.\n", books.Len())
	}

	books.Remove("Harry Potter and the Chamber of Secrets")

	for book := range books.Iter() {
		fmt.Println(book)
	}
}
```
