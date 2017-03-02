---
layout: post
title: Functional Programming vs Object Oriented Programming
date: 2017-02-28 12:00
comments: true
external-url: 
categories: Programming
---
>When do you choose functional programming over object oriented?

### When you anticipate a different kind of software evolution:

- Object-oriented languages are good when you have a fixed set of `operations on things`, and as your code evolves, you primarily add new things. This can be accomplished by adding new classes which implement existing methods, and the existing classes are left alone.

- Functional languages are good when you have a fixed `set of things`, and as your code evolves, you primarily add `new operations` on existing things. This can be accomplished by adding new functions which compute with existing data types, and the existing functions are left alone.

### When evolution goes the wrong way, you have problems:

- Adding a new operation to an object-oriented program may require editing many class definitions to add a new method.

- Adding a new kind of thing to a functional program may require editing many function definitions to add a new case.


This problem has been well known for many years; in 1998, [Phil Wadler dubbed it the "expression problem"](http://www.daimi.au.dk/~madst/tool/papers/expression.txt). Although some researchers think that the expression problem can be addressed with such language features as mixins, a widely accepted solution has yet to hit the mainstream.

> ## What are the typical problem definitions where functional programming is a better choice?

Functional languages excel at manipulating symbolic data in tree form. A favorite example is compilers, where source and intermediate languages change seldom (mostly the same things), but compiler writers are always adding new translations and code improvements or optimizations (new operations on things). Compilation and translation more generally are "killer apps" for functional languages.

See more answers at [here](http://stackoverflow.com/questions/2078978/functional-programming-vs-object-oriented-programming/2079678#2079678)

This post via [stackoverflow.com](http://stackoverflow.com/questions/2078978/functional-programming-vs-object-oriented-programming/2079678#2079678)

>Happy Coding!
