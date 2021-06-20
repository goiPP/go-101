# Basic Go
Go is not a OO language, it is static typing, pass by value
### :peach: Go CLI
- `go build` = compile go code file eg: `go build main.go` will produce main.exe, to execute Mac: ./main Win: click on main.exe
- `go run` = compile andexecute 1-2 files eg: go run main.go state.go
- `go fmt` = format all the code in each file in the current dir
- `go install` = compile and install a package
- `go get` = download the raw source code of package
- `go test` = run any tests associated with the current project

### :peach: Note

####  Package statement
1. Executable (package main) = generate a file that we can run, contains '*.go' files. (another name than main, when run go build it will not generate executable file. It requires func main() in the package, which is run auto when exec.
2. Reusable (package xxx) = code used as 'helpers where we put reusable logic

####  Import statement
- to import other package to be used in this package. ['fmt'](https://golang.org/pkg/fmt/) is a stand package of go.

####  Go is static typing language
- fundamental Go Types: bool, string, int, float64
> Static typing: types are known and checked by a compiler for correctness before running your program eg: C++, Java, Go. 
> So for Go `var card string = "hi"` is not neccessary. We can do `card:= "hi"`

> Dynamic typing: types are only known as your program is running eg: Ruby, Python, Javascript: var num = 123 (number) then num = "123" (string) is allowed

####  Decalring variable
- fundamental Go Types: bool, string, int, float64, Arrays (primative, cannot resize)
- Declaring ways:
    - :heavy_check_mark: Specify type (can do but not neccessary) `var card string = "Ace of Spades" `
    - :heavy_check_mark: First time initializing `value card := "Ace of Spades"`
    - :heavy_check_mark:  Reassigning `value card = "Five of Diamonds"`
    - :heavy_check_mark: Declare first then assign
    
             func main() {
                   var card int
                   card = 52
                   fmt.Println(card)
             }
    -  Declare value outside of func body
          :x:
           
            card := "Ace of Spade"
            func main() {...}
          :heavy_check_mark:

            var deckSize int
            func main() {
              deckSize = 50
              fmt.Println(deckSize)
            }   

####  Functions and Recievers
- Go is not a OO language
- Functions: function from the same package in different file can be used without need of importing
- Receivers: equivalent to "this" in java, go is avoiding this and self usage byb introduce this receiver.

        package main

        func main() {
            printState()
            cards := deck{"a card", newCard()}
            cards = append(cards, "another card") // return a new slice (not modify existing ones)
            cards.print() // usage of reciever
        }
        
        ---
        package main

        import "fmt"

        func printState() {
            fmt.Println("California")
        }
        
        ---
        package main

        import "fmt"
        
        type deck []string

        func (d deck) print() {           // (d deck) is a Reciever
            for i, card := range d {
                fmt.Println(i, card)
        }
#### Testing with Go
- Go does not have a test framework such as mocha, jasmine, selenium
- A test can be done with command `go test` that will run test on files (under main package) with suffix``"_test.go"`` .
- Errors that could be found when run `go test`
    - :bookmark_tabs: go: cannot find main module
    
                go: cannot find main module, but found .g\Workspace\Go\go-101    le there, run:
                    to create a module there, run:
                    cd .. && go mod init
                    
                    
        :bulb: The problem is due to the go default path. Run below command to will create `go.mod` file in cuurent directory

                `cd ~\Workspace\Go\go-101\cards`
                `go mod init cards`
    - :bookmark_tabs: Errorf call has arguments but no formatting directives         
        :bulb: wrong syntax, we can use 1 of these syntax

                t.Errorf("Expected deck of length 16, but got %v", len(d))

                or

                t.Error("Expected deck of length 16, but got ", len(d))

#### Pointer
        
        type person struct {
            firstName string
            lastName  string
        }
        aPerson = person{"first", "last"}
        aPersonPointer := &aPerson
        
- **`&variable`** (&aPerson) =  is an operator that give the memory address of the value that aPerson is pointing at = aPersonPointer (001f)
- **`*pointer`** (*aPersonPointer) = is an operator that give the value this memory address is pointing at = value at address 001f => person{"first", "last"}

        address point to 0001, value point to person{..}
        turn address into value   with *address
        turn value   into address with &value
    - `*type` = this is not an operator but just a type, pointer of a type eg: `(pointerToPerson *person)`

            package main

            import "fmt"

            func main() {
             name := "bill"

             namePointer := &name

             fmt.Println(&namePointer) =>0x0001
             printPointer(namePointer) =>0x0002
            }

            func printPointer(namePointer *string) {
             fmt.Println(&namePointer)
            }
    - &namePointer is the address of namePointer. while namePointer is the address of name

:low_brightness: Pointer shortcut
- With the Receiver define type as pointer, we can use function with that recevier with an object. It will auto retrieve the pointer of that object

        aPerson.update("new") === (&aPerson).update("new")  
        
        func (pointerToPerson *person) update(newName string) {
        (*pointerToPerson).name = newName 
        }
:low_brightness: Value Type and Reference Type
- Value Type = int, float, string, bool, structs
    - use pointers to change value in the function
- Reference Type = slice, map, channel, pointer, function
    - no need point, Go will still create a copy, but point to the same underlying data structure. see below example

:low_brightness: Pointer with Slice/ Reference type
- aSlice[index] = will actually modify the value at index of aSlice object
- how?
    - when we create a Slice we will get both Slice and Arrays (as a underlying of Slice in Go)

        aSlice:point_right:@0001 : "a", "b", "c" (:point_down:)

        aArray:point_right:@0002: []string{"a","b","c"}

    - when calling a function(aSlice) => Go will a copy of aSlice, but point to the same array of the original Slice

        aSlice:point_right:@0001 : "a", "b", "c" (:point_down:)

        aArray:point_right:@0002: []string{"a","b","c"}

        aSliceCopied:point_right:@0003: []string{"a","b","c"} (:point_up:)

#### Map Vs Struct
- Map
    - all keys/values must be same type
    - keys are indexed, we can iterate over them. map[key]
    - don't need to know all the keyes at compile time
    - Reference type
- Struct
    - Values can be differnet type
    - Keys don't support indexing
    - need to know all the fields at compile time
    - Value Type

#### Interface
- problem: We have englishBot and spanishBot that both have `func (englishBot/spanishBot type) printGreeting()` to print hello/hola. In this case we would have 2 identical logic in duplicate functions because having different type of receiver which is :(
### :sunflower: Useful Links
- https://play.golang.org/ (Go online playground)