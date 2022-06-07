# Margay Programming Language

## About The Project

University Project on developing a General Purpose Language.

Our programming language owes it's name to this cute [cat](https://en.wikipedia.org/wiki/Margay).

## Running the App

To run the REPL type `go run .` in the root folder.

To execute a file type `go run . <file path>`.

To see the Abstract Syntax Tree type `go run . <file path> -t`.

## Learning Margay

Margay has a C-like syntax, so it should be easy to learn if you know any C-like languages.
Let's take a look at some examples.

### Control Statements & Input/Output

```js
// Guess Number Game

write("\nWelcome to Guessing Game written in Margay!\n");
// writing message to console

win = rand(101); // random integer from 0 to 100
attempts = 5;
game = false;

write("\nGive me a number from 0 to 100 -> ");

// for loop works like while loop
for (attempts > 0) {

    // reading input from console and converting it to integer
    n = int(read());

    // if/else blocks
    if (n == win) {
        game = true;
        attempts = 0;
    } else {

        if (n < win) {
            write("\nGuess higher -> ");
        }

        if (n > win) {
            write("\nGuess lower -> ");
        }

        attempts = attempts - 1;
    }
}

if (game) {
    write("\nGuessed the number! You won!\n");
} else {
    write("\n\nNo attempts left! You lost!\n");
}
```

### Array Manipulation

```js
// arrays hold any data types,
// including functions or other arrays
arr = ["faf", 3.14, 203, true];

arr = pop(arr, 1);
// arr = ["faf", 203, true]

arr = push(arr, "team 11", 2);
// arr = ["faf", 203, "team 11" , true]

str = "";
i = 0;
// for loop works like while loop
for (i < len(arr) - 1) {
    str = str + arr[i] + " ";
    i = i + 1;
}

write(str, "\n"); // faf 203 team 11
```

### Closures

```js
// closure
greet = fn(greeting) {
    return fn(name) {
        write(greeting + " " + name + "!\n");
    }
}

sayHi = greet("Hi"); // returns a function

sayHi("Madalina"); // Hi Madalina!
sayHi("Valeria"); // Hi Valeria!
```

### Recursion

```js
fibonacci = fn(n) {
    if (n == 0) {
        return 1;
    }

    if (n == 1) {
        return 1;
    }

    // recursion
    return fibonacci(n-1) + fibonacci(n-2);
}

write(fibonacci(10)); // 89
```

For more examples check examples folder.
