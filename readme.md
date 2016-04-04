Here at HacKSU we've talked about a lot of different programming languages
Python, JavaScript, more JavaScript, but today I figured it would be neat to touch
on a language that we've not talked about before. Go. No don't go anywhere that's
the name of the language. Yes I know it's hard to Google, that's why everyone calls
it Golang, so if your interested search for that and you should find plenty of info.

There a couple of things that make Go special. First it's C like which should make
it familiar to everyone here as it's in the same family of languages as Java,
JavaScript, and C++, but unlike all those languages its a new language built with
lessens learned from those languages. It's also fully compiled so is meant to be
fast and responsive while making typing kind of more optional. It also has great
support for paralyzation which is the best way to make your program work faster.

My goal here like in all things is not to teach you everything about Go. I don't know enough,
but just to teach you enough to know it's out there and a bit of how to use it
and what you might use it for.

# Installing

First thing first we need to install Go. Don't worry if you don't have it installed
already you can follow along online at [play.golang.org](https://play.golang.org/).

If you have the time follow these instructions to install it locally
[golang.org/doc/install](https://golang.org/doc/install). Make sure you set the
environmental variable GOPATH to a directory that will store all your go stuff. Put some
thing like `export GOPATH=$HOME/work` into your `.profile` if your on a Mac or a
Linux box and search for environmental variables and add a new one with the name
GOPATH if your on Windows.

# Hello... go

Go is a bit strange in that everything about using it revolves around the a workspace
directory defined by the above mentioned environmental variable. To create our code
we need to cd into this directory and add a directory to put your source code in.
`mkdir src` cd into this directory and make a new directory to hold our project.
Technically we should create a hierarchy of directories to make sure that our package is
unique like `github.com/IApark/hello_world`. We won't be doing that today. Why?
Because it's too much work. Instead just choose a name that makes sense to you
and that you're remember (spaces make it harder), add the directories again using
mkdir, like `mkdir HacKSU_project` Open that directories in your text editor of
choice. You can open the current directory in Atom with `Atom .`

Make file named `main.go`, the name doesn't matter just the .go at the end which
will tell the compiler it is a Go source file. The first thing we need to do in
this file is tell go what package we are in. In this case we are in the `main`
package (no relation to the name of the file). You can think of this kind of like
a namespace in C++ if you know what that is. Don't worry if you don't. To do
this add `package main`

Next we will need to import the `fmt` package so we can do simple things like printing
text. To do this we just type `import "fmt"` Note the quotes around `fmt` this is
because fmt is directory and could have spaces or other strange characters in it

Lets add a main function

    func main() {
    }

This is the function that will get called when our program starts running.

In this we just need to add something to print hello world. `fmt.Println("Hello World")`
here we have the `fmt` because that's the package Println is in. Println is uppercase
because all public identifiers in golang are.

Compile it by typing typing `go install` either by itself from inside the project
directory or type `go install ` and the name you gave the directory. Then cd into
the bin directory of your workspace and execute your script with `./HacKSU_project`
(assuming you named the directory HacKSU_project).
