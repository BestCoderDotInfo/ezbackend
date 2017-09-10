---
title: What's the difference between ==, ===, eql? and equal? in Ruby
layout: post
date: '2017-09-10 08:00:00'
comments: true
categories: ruby
keywords: programming, ruby, triple equals, double equals, ==, ===, eql? method, equal? method
excerpt: The fact is, new beginners who fall in love with Ruby often mistaken for equality-related methods
---

# Double equals "==" and triple equals "==="

The fact is, new beginners who fall in love with Ruby often mistaken for equality-related methods: `==`, `===`, `eql?` and `equal?`.

As a programmer and often jump job, you must have encountered questions about the Javascript in the interview right? If not then I would like to reiterate 😅.

In Javascript, operator `===`` returns **true** only if both operands are of the same type and have the same value. If compare different types, the result will return **false**.

### So what about Ruby?

The `===` operator in Ruby is often referred to as the case equality operator with the `==` operator, also known as the general equality operator.

The `==` comparator is valid or not. This is the most common and most common way in most programming languages.

The `===` operator is really something very interesting. It is all around the place in Ruby but the most all that ever seen that made it. But, how does it go everywhere and no one ever sees it? It hidden inside a regular control panel, "case / when". Whenever you are using "case / when", in fact you are using the "===" operator and this makes the case statement on Ruby much more powerful so with languages like C or Java.

The `===` operator is simply interpreted as a case type comparison. Conditions of the case will be implemented with each corresponding class as:

- Range
- Regex
- Proc …

Example:

```ruby
(1...10) === 5 # => true
```

The `===` operator will check the array in a 5-character statement. If it have `5` in arry will return **true**, opposite return **false**.

```ruby
(1..5) === 3           # => true
(1..5) === 6           # => false
Integer === 42          # => true
Integer === 'fourtytwo' # => false
/ell/ === 'Hello'     # => true
/ell/ === 'Foobar'    # => false
"a" === "b" # false # different values, different objects
"a" === "a" # true # same values
```

Next:

```ruby
"test" == "test"  #=> true
"test" === "test" #=> true
```

#### What's the difference between "==" and "==="?

```ruby
String === "test"   #=> true
String == "test"    #=> false
```

So `===` is merely a comparison of values, not a comparison of objects whose case is used for comparison and `===` also known as case equality.

### Usage

#### Array.grep

The array has a method called grep use `===`.

```ruby
# grep(pattern) → array

(1..100).grep(38..44)
#=> [38, 39, 40, 41, 42, 43, 44]

names = %w(
  William
  Kate
  Adam
  Alexa
  James
  Natasha
)
names.grep(/am/)
# => %w(William Adam James)
```

The result is an array with elements satisfying `===` with the pattern of grep.

#### Ranges

The `===` operator will check to see if the object is one of the elements of that range.

```ruby
(2..4) == 3 # => false
(2..4) === 3 # => true
(2..4) === 6 # => false

(Date.new(2017, 8, 21)..Date.new(2017, 8, 27)) === Date.new(2017, 8, 27)
# => true

(Date.new(2017, 8, 21)..Date.new(2017, 8, 27)) === Date.new(2017, 8, 29)
# => false

("a".."z") === "a"
# => true

("a".."z") === "abc"
# => false
```

#### Class / Module

```ruby
mod === obj #→ true or false
```

The `===` operator returns true if it is an instance of mod or one of mod. Use limit mode for module, but can be used to distribute objects by layer. Go to the basic database as follows:

```ruby
obj.kind_of?(mod)
```

Example:

```ruby
"text".class.ancestors
# => [String, Comparable, Object, Kernel, BasicObject]

String === "text"
# => true

Object === "text"
# => true

Comparable === "text"
# => true

Numeric === "text"
# => false
```

#### Regexp

```ruby
rxp === str #→ true or false
```

Basic same as:

```ruby
rxp =~ str >= 0
```

Example:

```ruby
/^[a-z]*$/ === "HELLO"
#=> false

/^[A-Z]*$/ === "HELLO"
#=> true
```

#### Proc

```ruby
proc === obj # → result_of_proc
```

Example:

```ruby
is_today = -> (date) { Date.current === date }

is_today === Date.current
# => true

is_today === Date.tomorrow
# => false

is_today === Date.yesterday
# => false
```

#### Lambdas

Same as `Proc`

```ruby
is_even = -> (x) { x % 2 == 0 }
is_even == 4 # => false
is_even === 4 # => true
is_even === 5 # => false
```

#### Object

For all most objects, `===` similar `==`.

# eql? method and equal? method

Of the above methods is `equal?` It is best described in terms of equality by its name.

`equal?` will return **true** only when it (object called by `equal?`) and the parameter (object called by `equal?`) is the **same** object.

Example:

```ruby
some_word = "word"
some_other_word = some_word

some_word.equal? some_other_word # true
```

`eql?` will the return result be `true` when both objects are called by `eql?`And the object is called by `eql?` has a garia values.

```ruby
1 == 1.0 #=> true 
1.eql? 1.0 #=> false
"test".eql? "test" #=> true
```

# Finally

```ruby
a = 'a' # => 'a' 
other = a.dup # => 'a'

a == other # => true
a === other # => true
a.eql? other # => true
a.equal? other # => false
a.equal? a # => true
```

Resources:

- [What is the difference between ‘===’, ‘==’, ‘.equal?’ and ‘.eql?’ in Ruby?](http://www.rian.me/2013/10/15/what-is-the-difference-between-equals-equals-equals-and-equals-equals-in-ruby/)

- [The === (case equality) operator in Ruby](http://blog.arkency.com/the-equals-equals-equals-case-equality-operator-in-ruby/)

- [Ruby Basics - Equality operators in Ruby](https://mauricio.github.io/2011/05/30/ruby-basics-equality-operators-ruby.html)
- [The rarely used === in Ruby](https://coderwall.com/p/53xawg/the-rarely-used-in-ruby)
