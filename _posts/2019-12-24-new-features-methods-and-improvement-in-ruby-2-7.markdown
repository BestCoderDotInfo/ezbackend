---
layout: post
author: derek
image: assets/ruby-2.7.jpg
featured: false
hidden: false
title: " New features, methods and improvements in Ruby 2.7"
date: 2019-12-24 12:00
comments: true
external-url:
categories: Ruby
keywords: Ruby, Rails, Web Development
excerpt: Ruby never stops improving!
Version 2.7 is around the corner with new features & methods. It’s scheduled for release on December 25, 2019.
---
**Ruby never stops improving!**

Version 2.7 is around the corner with new features & methods. It’s scheduled for release on December 25, 2019.

### Enumerable#tally

This is a new Ruby method that counts all the elements in an array & returns a hash with their counts.

You can do this yourself, but this `tally` method saves you work.

Example:

```ruby
%w(a a a b b c).tally
```

Result:

```
{"a"=>3, "b"=>2, "c"=>1}
```

### Numbered Parameters For Blocks [Experimental]

An interesting new feature, which I’ve been wanting for a while, is a default name for block parameters.

Here’s a regular block, with one parameter:

```ruby
[1,2,3].each { |n| puts n }
```

This `|n|` is the parameter, which you have to define to use it.

But what if we had a default name?

Well, that’s one of the goodies that Ruby 2.7 brings with it, although as an experimental feature, it can save us a lot of typing.

```ruby
[1,2,3].each { puts _1 }
```

Where `_1` refers to the first parameter, you can use `_2` if you have a 2nd one, and yes,` _3` for a 3rd parameter, etc.

This feature was first proposed to look like `@1`, but that looks too much like an instance variable, so after some discussion, it was settled on using `_1` instead.

### Array#intersection

New method, but no new functionality. This is more like an `alias`!
Ruby 2.6 introduced Array `union` & `difference` methods, to match the more succinct equivalents of Array#|, and Array#-.

But the missing operation was `intersection`, which has another hard to remember short version.
The `Array#& method`.

Now:

What do these methods do & how do they work?

```ruby
[1, 2, 3].intersection([2, 3, 4])
# [2, 3]
[1, 2, 3] & [2, 3, 4]
# [2, 3]
```

The name kind of gives it away, `intersection` finds the intersection between two arrays. In other words, it finds which elements are in common.

### Enumerable#filter_map

This `filter_map` method is an attempt to combine the `select` & `map` methods into one.

Why?

Because it’s a common operation to filter a list first, then map the remaining elements.

```ruby
(1..8).select(&:even?).map { |n| n ** 2 }
# OR
(1..8).map { |n| n ** 2 if n.even? }.compact
```

I’ve used both, but I tend to lean into the first because the intention is more clear.

In Ruby 2.7

```ruby
(1..8).filter_map { |n| n ** 2 if n.even? }
# [4, 16, 36, 64]
```
I’m not the biggest fan of having if statements inside blocks, but sometimes they’re needed to get the job done.
**Besides that, there is something you should know.**
The `filter_map` operation doesn’t behave like `map` + `compact`, because it removes false objects, and `compact` doesn’t.

### Enumerator#produce

Here’s another new method you may find interesting, but it may require some creativity to get the most out of it.

```ruby
Enumerator.produce(1, &:next).take(5)
# [1, 2, 3, 4, 5]
```

In this example, it doesn’t matter if you do `take(10)` or `take(10_000)`, you’ll get an infinite number of values from it.

Btw, `1` is the initial value.

And `&:next` is the method called on that value to produce the next element in the sequence.
