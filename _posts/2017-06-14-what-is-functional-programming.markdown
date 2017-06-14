---
layout: post
title:  What is Functional Programming ?
date: 2017-06-14 09:00
comments: true
external-url:
categories: programming
keywords: programming, functional programming
excerpt: Functional programming is a paradigm which concentrates on computing results rather than on performing actions.
---
>Learning Functional Programming takes a while. So be patient.

## What is Functional Programming ?

Functional programming is a paradigm which concentrates on **computing results** rather than on **performing actions**.  That is, when you call a function, the only significant effect that the function has is usually to compute a value and return it. Of course, behind the scenes the function is using CPU time, allocating and writing memory, but from the programmer's point of view, the primary effect is the return value.  Objects in a functional programming language are often immutable (a.k.a. const or final); instead of changing an object, you allocate a new object which looks like the old one except for the change.  Compare this with an **imperative** programming language like Java, where progress is made by changing objects' fields, inserting them into sets, etc.

In a **pure** functional language, like Haskell, most functions are guaranteed by the type system not to perform any other actions.  In an **impure** functional language, like ML, a function may have other side effects, such as querying a database or server, generating random numbers, reading or writing the disk, etc.

Read [Functional programming on Wiki](https://en.wikipedia.org/wiki/Functional_programming)

Functional programming is a style of programming that emphasizes the evaluation of expressions rather than the execution of commands. **Erlang** programming language is described as a functional programming language. Erlang avoids the use of global variables that can be used in common by multiple functions since changing such a variable in part of a program may have unexpected effects in another part.

In an earlier definition from the ITU-TS, functional programming is "a method for structuring programs mainly as sequences of possibly nested function procedure calls." A function procedure is a relatively simple program that is called by other programs and derives and returns a value to the program that called it.

Read more at: [So You Want to be a Functional Programmer on Medium](https://medium.com/@cscalfani/so-you-want-to-be-a-functional-programmer-part-1-1f15e387e536)

>Happy Coding!


