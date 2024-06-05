# Go Tutorial

Go is compiled , Go tool can be run direclty on the system , as far as Go in installed.
Unlike Java , we dont get any Byte code in order to run it on some other machine.
There are no swtich cases here , and it misses a lot more features try-catch no semi-colons etc.
We dont have classes in Go , we have structs here
We dont have Method Overloading here.

Golang.org - official site

Go mod init - gives go.mod file 
Install go extension in Vscode 

Go help <command> - to know about the command

go mod init hello - to create the go.mod file

Go run main.go - is to run the file main.go

Go env GOPATH - to get the path 

Lexer - The job of the lexer is to check the code grammatically (indentation, semi colons , variables)

Types in Go Lang - Almost everything in Go is kind of a type . Example : String , Bool , Integer , Floating , Complex (number with imaginary values) , Arrays, Slices , Maps , Structs , Pointer , Functions , Channels.

fp - is the shortcut to print fmt.println()

walrus operator - via using this no need to add var or datatype keyword 
        Exmaple  : numberOfUser := 3000
        But this is only allowed inside Functions

https://pkg.go.dev/ - is the place where we can search for any package in Go

go build - is for builind the file

GOOS="windows" go build - will build the file for windows OS , from which we get mytime.exe , which can run layer 

In Go memory allocation and deallocation happens on its own
new() - allocate memory but without initialization , gives zeroed storage
make() - allocate memory and INIT , gives non-zeroed storage
Garbage Collection happens automatically
