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

# [Hello... go](https://github.com/hacksu/go-example/tree/70e55a6f72001daaaa82c729ab8f23d8e6f21599)

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


# [Variables](https://github.com/hacksu/go-example/tree/4bdea6691e466e107bbef0236b3e87ef1ba77bac)

Variables might be one of the stranger looking things in Go. First variables in
Go have type unlike languages like JavaScript or Python so we have to decided when
we create a variable what kind of stuff we want the variables to hold, but unlike C++
or Java go will by default guess the type of a variable which can make it feel
at times like an untyped language. The strangest part though is that the type goes at
the end so to define an int i we use `var i int = 0` or `var i = 0` and it will
assume a type of int if you want to leave out even the var we can do that by using
`:=` like `i := 1` so if we want to store who we're greeting as a variable we can
do it by defining a variable like `person := 0` (we'll CS people) then we can either
print it with `fmt.Println("Hello Person " + string(person))` note we have to convert
to a string with `string(...)` before we can add two strings. This holds with all
conversions. Go will never assume a conversion between two types. We can also just
pass it as another argument like `fmt.Println("Hello Person", person)`

# [Control](https://github.com/hacksu/go-example/tree/23b426cb7485b3dd752d64b7f994c87b3c411b74)

Right now we have the tools to make a very nice program that does exactly one thing
if we want to have it vary what it does we use things like for loops and if statements.

Go only has one type of loop, the for loop so they've done a lot to make it flexible.

The version that most closely matches what we're used to is something like `for i := 0; i < 100; i++`

It basically looks like C++, but we don't put parenthese around everything. We can
easily use this to great 10 people instead of just one

    for person := 0; person < 10; person++ {
       fmt.Println("Hello Person", person)
    }

People who have already done programing were probably asking from the moment I
started talking about how we're supposed to do while loops. It's easy enough,
we just leave the init and the step stages empty like `for is_more()` so the above
could also be written as

    person := 0
    for person < 10{
       fmt.Println("Hello Person", person)
       person++
    }

If we don't give any arguments like `for {}` we get an infinite loop.

If statements work the same way and interestingly we can also had pre statements. While I personally can't stand them switch statements also work exactly as expected though it is assumed
each block ends in break.

There is a one flow control statement in Go, however, that may be more confusing
for someone coming from C++ `defer` specifies that a function listed should be
executed at the end of the function. It's important for cleanup because it means
if something goes wrong and the function needs to get exited early critically
cleanup code will still get run. We're just going to use it for fun though so we
can do

    func main() {
   	  defer fmt.Println("two")
   	  fmt.Println("one")
    }

and it will get printed in the right order.

# [Functions](https://github.com/hacksu/go-example/commit/0f2ef9b88f510a7d69d50258152cdec0750a2e38)


You can probably guess how to define your own function but lets do it for real.

Lets define a function to return if a number is prime and if not the first number
it can be divided by. We will not be doing this efficiently

We need our function to take an int and for now return another int to hold the
first number it can be divided by (we'll just make it less than or equal to  0 if
we can't find anything) that would look like.

func is_prime(n int) int {
	return 0
}

before we get to actually implement the function I'd like to explain what exactly this means.

We start with the `func` keyword to tell the compiler we are defining a function.
we list the identifier next. If we want it to be available from other packages we'd
need to make this uppercase. Then we define the arguments with the type listed at
the end as normal. Finally, at the very end we define the type the function returns.

The simplest way to check if a number is a prime is to start at 2 and check if it
is divisible by any number between two and it's square root. We can do this with a
for loop like so `for i := 2; i <= n/2; ++i` We're going to use half n as a terrible
terrible approximation of sqrt(n). If you don't know % is the remainder operator.

If you remember doing long devision way back you may remember the remainder, it's
the bit that is left over after you perform integer devision so if the remainder
is ever 0 we know that the number was evenly divided so we end up with.

    func is_prime(n int) int {
  	 for i := 2; i <= n/2; i++ {
   	 	 if n % i == 0 {
   			 return i
   		 }
   	 }
  	 return -1;
    }

  One more thing, this is well and good, but most of the time we just want to know
  if the number is a prime not anything about it's factors. Luckily Go lets us
  return multiple things. We do it like this

     func is_prime(n int) (bool, int) {
       	for i := 2; i <= n/2; i++ {
       		if n % i == 0 {
       			return false, i
  	     	}
       	}
       	return true, -1;
     }

 While we can be done we can't then directly use it in an if statement. Lets just
 use the version that returns an int.



# Go... faster

What if we want to do something more complicated such as list all the primes between
two numbers. Let's make a function to do that.

   func primes(start, end int) []int {
	   primes := make([]int, 0)

   	for i := start; i < end; i++ {
	   	if is_prime(i) < 0 {
		   	primes = append(primes, i)
		   }
	   }
	   return primes
    }

This uses slices which are basically Go's version of vectors. They are an expandable
list of items


Lets run this to get a lot of big primes. Say `fmt.Println(primes(100000, 200000))`

It might take a while. Luckily one of Go's great strengths is in parallelization.
This is something you may have heard have and if so you probably think of it as
a huge pain. Luckily go supports multitasking on a very low level in the language
which makes tasks like this easy. We can start a Goroutine (Go's name a function
running on another thread) very simply. Just put `go` in front of a function call.

Let's make a simple example.

    package main

    import "fmt"
    import "time"

    func main() {
    	go say("hello")
    	go say("world")
    	var input string
    	fmt.Scanln(&input)
    	fmt.Println("done")
    }

    func say(message string) {
    	for {
    		fmt.Println(message)
    		time.Sleep(100*time.Millisecond)
    	}
    }

We need to read something at the end because if we don't the program will exit right
away when the main function ends

There's a problem though. We can't get values back from our Goroutines at the moment
we'll used something called a channel for that. It's a something in Go that's designed
to solve this problem. We can push something into our channel in our Goroutine
and pull it out in the main thread. It will even block (wait) if there is no space
in the channel. We create a channel with `make` again like `primes := make(chan []int)`.
We can make a channel of any type if we want.  

Lets do it. First lets create a new function named something like fast_primes
