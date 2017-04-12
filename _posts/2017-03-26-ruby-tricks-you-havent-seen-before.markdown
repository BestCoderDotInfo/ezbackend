---
layout: post
title: 11 Ruby Tricks You Haven't Seen Before
date: 2017-03-26 00:10
comments: true
external-url:
categories: Ruby
keywords: ruby, ruby tricks, rails
excerpt: When you copy an object that contains other objects, like an Array, only a reference to these objects is copied.
image:
  feature: ../../assets/ruby-tips-and-tricks-1-638.jpg
---

## 1. Deep copy

When you `copy` an object that contains other objects, like an `Array`, only a reference to these objects is copied.
You can see that in action here:

```ruby
food = %w( bread milk orange )

food.map(&:object_id)       # [35401044, 35401020, 35400996]

food.clone.map(&:object_id) # [35401044, 35401020, 35400996]
```

Using the **Marshal class**, which is normally used for [serialization](https://en.wikipedia.org/wiki/Serialization), you can create a ‘deep copy’ of an object.

```ruby
def deep_copy(obj)
  Marshal.load(Marshal.dump(obj))
end
```

The results:

```ruby
deep_copy(food).map(&:object_id) # [42975648, 42975624, 42975612]
```

## 2. Different ways to call a lambda

```ruby
my_lambda = -> { puts 'Hello' }
my_lambda.call
my_lambda[]
my_lambda.()
my_lambda.===
```

If possible, you should stick with the first one (**call**), because it’s the one most people know.

## 3. Creating a pre-filled array

The Array class can take an argument **+ a block**, which lets you create an array with `n` elements. By default these elements are `nil`, but if you have a block, the values will come from it.

Example:

```ruby
Array.new(10) { rand 300 }
```

This will generate an array with 10 [random numbers](http://www.blackbytes.info/2015/03/ruby-random/) which are between 0 and 299.

## 4. True, false and nil are objects

```ruby
true.class  # TrueClass


false.class # FalseClass


nil.class   # NilClass
```

There is only one copy of these objects, and you can’t create more even if you wanted.
This is the [singleton pattern](http://c2.com/cgi/wiki?SingletonPattern) in action.

## 5. Lambdas are strict about arguments, but Procs don’t care

```ruby
my_lambda = ->(a, b)  { a + b }
my_proc   = Proc.new  { |a, b| a + b }

my_lambda.call(2)
# ArgumentError: wrong number of arguments (1 for 2)

my_proc.call(2)
# TypeError: nil can't be coerced into Fixnum
```

## 6. Execute code directly without irb or files

The **ruby** command has a number of interesting options you can use.
For example, with the **-e** flag you can pass in a snippet of code to be executed.

```bash
ruby -e '5.times { puts "Fun with Ruby" }'
```

You can find more by using the **-h** flag.

## Your own mini-irb in one command

Ever wanted to know how **irb** works? Well, this is a super-simple version of it.
Remember what ‘REPL’ stands for: Read-Eval-Print Loop.

```bash
ruby -n -e 'p eval($_)'
```

You won’t get a prompt, but go ahead and type some Ruby code.

```bash
"A" * 5

"AAAAA"
```

This works because the **-n** flag does this:

```bash
-n    assume 'while gets(); ... end' loop around your script
```

And **$_ is** a global variable. Which contains the following:

`The last input line of string by gets or readline.`

## 8. Unfreeze an object (danger!) 

There isn’t any Ruby method to unfreeze an object, but using the **Fiddle** class you can reach into Ruby internals to make it happen.

```ruby
require 'fiddle'

str = 'water'.freeze

str.frozen? # true

memory_address = str.object_id * 2

Fiddle::Pointer.new(memory_address)[1] &= ~8

str.frozen? # false
```

Don’t try this at home!

## 9. Objects with special identity

Ruby objects have an identifier or ‘id’ number you can access using the **object_id** method. Some objects have a fixed id: Fixnums, true, false & nil.

```ruby
false.object_id # 0

true.object_id  # 2

nil.object_id   # 4

1.object_id # 3

2.object_id # 5
```

Fixnum ids use this formula: (number * 2) + 1.
Bonus: The maximum Fixnum is `1073741823`, after that you get a Bignum object.

## 10. Avoid big output in irb or pry

If you are working in `irb` and want to avoid filling your screen with the contents of some really big array or string you can just append `;` at the end of your code.

```ruby
require 'rest-client'

RestClient.get('blackbytes.info');
```

Try again without the `;` to see the difference

## 11. Using the caller method to get the current call stack

Here is a code example:

```ruby
def foo
  bar
end

def bar
  puts caller
end

foo
```

Output:

```bash
-:3:in 'foo'

-:10:in ''
```

If you need the current method name you can use `__method__` or `__callee__`.

Bonus! Convert any value into a boolean.

```ruby
!!(1)   # true

!!(nil) # false
```

**That’s all!**