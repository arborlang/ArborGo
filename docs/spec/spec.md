# Spec for arbor

## Typing and Assignment

The first thing I wanted was Python like Typing; basically dynamically, but strongly typed. This means an expression like this: `'x' + 1` is invalid and will throw an error, but these two statements are valid: `x = 1; x = 'dd';`. However, to maintain safety, I will also like optional parameter type checking. If you define a functions like so: `(a:int, b:string)` then you would expect a to always be an int and string to always be a string.

The second thing I decided on was that everything must be assigned. In order to make the language simpler to implement, I did away with any special keywords to define a function. Unlike in Python, JavaScript, or C/C++ a function is inherently anonymous, unless assigned to a variable. The way to define a function would be: `() -> <function body>`. In order to keep that function around you would need to do something like: `foo = () -> <function body>`. Of course, this runs the risk of a programmer accidentally overwriting their function.

To make the language "safer", I decided that every variable had to be declared before you use it. This prevents a programmer,especially one with atrocious spelling like me, from accidentally declaring a variable because of a spelling error in one place. For right now, the only two keywords to define a variable is `let` and `const`. I decided on `const` because it is pretty self explanatory that the variable is constant. Plus C/C++ and Javascript use the `const` keyword, so I think it would be pretty easy for most developers to pick it up.

The choice of `let` has really nothing to do  with javascript, all though maybe it does a bit. I choose the keyword `let` to more closely align with "math speak" (i.e `let x be a value in universe`). I also chose let because it corresponds to a [lambda abstraction](https://en.wikipedia.org/wiki/Let_expression#Definition){:target="_blank"} and [lambda calculus](https://en.wikipedia.org/wiki/Lambda_calculus){:target='_blank'} is really the foundation of all pure functional languages.

I'm still torn about providing type declarations, such as `int`, `float`, `char`, and `array` because I don't see them as completely necessary. It may be nice to have if for no other reason than it makes the language easier to read. At the same time however, since the type can be inferred from what you are assigning a variable to, and function definitions provide optional type checking, I don't know if this is absolutely necessary. I am leaning towards no, but If I do add support for type declarations, it will be with the `let` keyword, not instead of.

## Functions Definitions

I also liked the [JavaScript arrow function](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Functions/Arrow_functions){:target="_blank"} syntax (this is part of the reason why arbor has `->` to define functions), and especially the behaviour where a single line means return and a function body means do this whole function body. But at the same time I like Python's no curly brace syntax. So what I did was do something that a lot of other languages do: I defined an end statement. So in Arbor, you would define a function two ways:

    let foo = (a, b, c) -> a + b + c;

or

    let foo = (a, b, c) -> {
        return a + b + c;
    }

Another thing I like about python is the way you can implement default parameters. In arbor it will be done similarly:

    (a = 1, b = 2) -> a + b;

One thing that a lot of people complain about in Python, and this also trips up people coming from languages such as Java or C/C++, is that you can pass in any type into a function. This could cause issue where you expect a string, but instead receive an int, causing your application to crash. All though this is something you're unit tests should catch, I also wanted to "fix" this with compile time type checking. Plus I think this makes the language that much more descriptive in my mind. If you want type checking and defaults, I think it should look like this:

    (a:int, b:int = 2) -> a + b;

Taking another principle from Python is packing and unpacking of variables in functions. Similarly to `function(...args)` in es6, python allows you to define lift over params as such: `def foo(a, *args)`. Arbor will do something similar: `foo = (a, b, *args) -> ...`. Then you can call this function like so: `foo(1, 2, 3, 4, 5, 6)`.
And like Python, I also want Arbor to support named leftover variables: `def foo(a, b, **kwargs)`, in arbor would be: `foo = (a, b **kwargs) -> ...`. Where `kwargs` is a dictionary.

And finally variable unpacking. Python has this really neat concept called unpacking if you have a function definition like

    def foo(a, b, c, d, e, f, g):
        pass

you can call that function like this:

    arr = [1, 2, 3, 4, 5, 6, 7]
    foo(*arr)

    # or

    vals = {
        "a": 1,
        "b": 2,
        "c": 3,
        "d": 4,
        "e", 5,
        "f": 6,
        "g": 7
    }
    foo(**arr)

With default parameters, any parameter not in the dict or array will be the default. I want Arbor to support this exactly the same way:

    foo = (a, b, c, d, e, f, g) -> null;
    arr  = [1, 2, 3, 4, 5, 6, 7];
    foo(*arr)

    // and

    dict = {
        a: 1,
        b: 2,
        c: 3,
        d: 4,
        e: 5,
        f: 6,
        g: 6,
    }
    foo(**dict)

## Control and flow

In true functional programming fasion, I decided to do away with loops. Instead all loops should be implemented using recursion constructs. Additionally, built in functions such as `forEach`, `map`, `filter`, `fold` or `reduce` will be implemented in order to make implementing loop behaviour easier.

Additionally, as well as having traditional control flow, `if`, `else`, `else if` statements, I will also have haskell like predicates. These could be similar enough to case statements in Elm. These would look like this:

    (a, b) -> {
        : a > b ->
            if (a != b)
                return "greater than"
        : a < b -> "less than";
        : true -> "equal to";
    }

This should be functionally equivalent to

    (a, b) -> {
        if (a > b) {
            return "greater than";
        }
        else if (a < b) {
            return "less than";
        }
        else {
            return "equal to";
        }
    }

And finally, ternary operators. I really like ternary operators. They are elegant and makes code easier to read for small stuff. However, I think JavaScript's and C/C++'s ternary operator leaves something to be desired. I really like Python's ternary operator and that is exactly how ternaries in Arbor should work: `value if <condition> else other value`.

## Data and Types

The only data types I want to include in Arbor are Integers, Floats, chars, Arrays, and Dictionaries. A string keyword will be available, but, like C/C++, it is really just an array of chars. Arbor will also provide `true` and `false` keywords that is really just 1 and 0 integers. Arbor should have a typedef operator that allows developers to define their own types. This would be similar to how C defines structs:

    Person = type {
        name: string,
        age: int,
        favorites: array
    }

This defines a type so that you can do things like

    person = instantiate(Person);
    person.name = "yoseph";
    person.age = 22;
    person.favorites = ["programming", "Arbor"]

or

    person = instantiate(Person, name="yoseph", age="22", favorites);

Functions are also first order citizens so that you can pass them as functions or in new types. Types can also refer to themselves, making the type composable and building complex structures like a tree.

## Error Handling

Arbor handles error handling differently than other languages. Instead of the try/catch of other languages, Arbor uses a concept called signals in order to enter recoverable states. This enables errors to be handled via Algebraic Effects, making more resilient software. For example, handling a null value using signals:

```arbor
fn SaveUserName (user: User, newUserName: String) -> User {
    if (newUserName == null) {
        newUserName = fatal new UserNameIsNullError();
    }
    user.userName = newUserName;
    return user;
}

fn SaveUser(user: User, newValues: User) -> User {
    try {
        SaveUserName(user, newValues.userName);
    } handle (UserNameIsNullError) {
        continue with "Oh snap, I forgot their name";
    }
}
```

This will raise a fatal signal, giving the call chain the opportunity to recover from the signal. If somewhere up the chain, the signal handle has a `continue` statement, then it will return from the point where the signal was raised and continue.

There are three levels of signals: `fatal`, `warn`, and `signal`. The difference between them is how they are handled in the un-handle case.

### `fatal` Signals

Fatal signals are closest to exceptions in other languages. When a fatal signal is emitted, Arbor will look up the call chain for the closest handle block that can handle the signal. If no handlers are found, then Arbor will crash the process and dump out its traceback, just like exceptions

### `warn` Signals

Warns are a level down from `fatal`s. Instead of crashing if no handler is found, Arbor will simply log the warning to the application logger but then continue its operation with a null value.

### `signal` Signals

This is the lowest level of signals. Nothing happens if there is no handler, and if there is no `continue` then it will continue with a null value.

### Resuming operation

After a signal is raised, then the programmer has the ability to resume the execution by calling `continue`. Continuing will tell arbor to go back to where the signal was raised and resume. You can pass values back by calling `continue with <value>`
