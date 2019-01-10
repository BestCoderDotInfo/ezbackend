---
layout: post
author: derek
image: assets/Ruby Interview Questions and Answers.jpg
featured: false
hidden: false
title: 15 Questions to ask during a Ruby Interview
date: 2018-08-06 14:00
comments: true
external-url:
categories: Interview
keywords: ruby, interview, questions
excerpt: Candidates will typically give you a range of responses based on their experience and personality. So it's up to you to decide the correctness of their answer.
---

First, Thank **Ryan Sobol** very much!

## What to expect

Candidates will typically give you a range of responses based on their experience and personality. So it's up to you to decide the correctness of their answer.

There are many solutions to most of these questions -- some aren't listed here. Candidates get my respect for knowing esoteric solutions, but I'm looking for developers with similar practices to my own.

Make no mistake, this list is not complete, by any means. But it does provide a useful tool to put a candidate's reaction to the same series of questions in perspective. Especially during an interview, when your mouth goes dry and and your hands get all sweaty.

## Begin!

Senior programmers won't have a problem with these, while junior programmers will usually give only half-answers.

### What is a class?

A text-book answer: classes are a blue-print for constructing computer models for real or virtual objects... boring.

In reality: classes hold **data**, have **methods** that interact with that data, and are used to **instantiate objects**.

Like this.

```ruby
class WhatAreClasses
  def initialize
    @data = "I'm instance data of this object. Hello."
  end

  def method
    puts @data.gsub("instance", "altered")
  end
end

object = WhatAreClasses.new
object.method
 #=> I'm altered data of this object. Hello.
```

### What is an object?

An instance of a class.

To some, it's also the root class in ruby (Object).

Classes themselves descend from the Object root class. (Kudos to Ezra)

### What is a module? Can you tell me the difference between classes and modules?
Modules serve as a mechanism for **namespaces**.

```ruby
module ANamespace
  class AClass
    def initialize
      puts "Another object, coming right up!"
    end
  end
end

ANamespace::AClass.new
 #=> Another object, coming right up!
```

Also, modules provide as a mechanism for multiple inheritance via **mix-ins** and **cannot be instantiated** like classes can.

```ruby
module AMixIn
  def who_am_i?
    puts "An existentialist, that's who."
  end
end

# String is already the parent class
class DeepString < String
  # extend adds instance methods from AMixIn as class methods
  extend AMixIn
end

DeepString.who_am_i?
 #=> An existentialist, that's who.

AMixIn.new
 #=> NoMethodError: undefined method ‘new’ for AMixIn:Module
```

### Can you tell me the three levels of method access control for classes and modules? What do they imply about the method?

**All methods**, no matter the access control, can be **accessed within the class**. But what about outside callers?

*Public* methods enforce **no access control** -- they can be called in any scope.

*Protected* methods are only accessible to **other objects of the same class**.

*Private* methods are only accessible within the **context of the current object**.

```ruby
class AccessLevel
  def something_interesting
    another = AccessLevel.new
    another.public_method
    another.protected_method
    another.private_method
  end

  def public_method
    puts "Public method. Nice to meet you."
  end

  protected

  def protected_method
    puts "Protected method. Sweet!"
  end

  private

  def private_method
    puts "Incoming exception!"
  end
end

AccessLevel.new.something_interesting
 #=> Public method.  Nice to meet you.
 #=> Protected method.  Sweet!
 #=> NoMethodError: private method ‘private_method’ called for
 #=>  #<AccessLevel:0x898c8>
```

### There are three ways to invoke a method in ruby. Can you give me at least two?
Here, I'm looking for the **dot operator** (or period operator), the **Object#send** method, or **method(:foo).call**

```ruby
object = Object.new
puts object.object_id
 #=> 282660

puts object.send(:object_id)
 #=> 282660

puts object.method(:object_id).call # (Kudos to Ezra)
 #=> 282660
```

## Separating the professional from the hobbyist

Senior programmers should be able to give competent answers for all questions. Junior programmers should answer some correct, but usually won't know them all.

### Explain this ruby idiom: a ||= b

A common idiom that strong ruby developers use all the time.

```ruby
# a = b when a == false
# otherwise a remains unchanged
a || a = b # (Kudos to Markus Prinz)


a = 1
b = 2
a ||= b #=> a = 1

a = nil
b = 2
a ||= b #=> a = 2

a = false
b = 2
a ||= b #=> a = 2
```

### What does self mean?

*self* always refers to the current object. But this question is more difficult than it seems because Classes are also objects in ruby. (Kudos to Stephen)

```ruby
class WhatIsSelf
  def test
    puts "At the instance level, self is #{self}"
  end

  def self.test
    puts "At the class level, self is #{self}"
  end
end

WhatIsSelf.test
 #=> At the class level, self is WhatIsSelf

WhatIsSelf.new.test
 #=> At the instance level, self is #<WhatIsSelf:0x28190>
```

This short snippet indicates two things:

  - at the *class level*, self is the **class**, in this case WhatIsSelf.
  - at the *instance level*, self is the **instance in context**, in this case the instance of WhatIsSelf at memory location 0x28190.

### What is a Proc?

Everyone usually confuses procs with blocks, but the strongest rubyist can grok the true meaning of the question.

Essentially, Procs are **anonymous methods** (or nameless functions) containing code. They can be placed inside a variable and passed around like any other object or scalar value. They are created by **Proc.new**, **lambda**, and **blocks** (invoked by the yield keyword).

**Note**: Procs and lambdas do have subtle, but important, differences in ruby v1.8.6. However, I wouldn't expect a candidate talk about these nitty-gritty details during an interview. (Kudos to Noah Thorp)

```ruby
# wants a proc, a lambda, AND a block
def three_ways(proc, lambda, &block)
  proc.call
  lambda.call
  yield # like block.call
  puts "#{proc.inspect} #{lambda.inspect} #{block.inspect}"
end

anonymous = Proc.new { puts "I'm a Proc for sure." }
nameless  = lambda { puts "But what about me?" }

three_ways(anonymous, nameless) do
  puts "I'm a block, but could it be???"
end
 #=> I'm a Proc for sure.
 #=> But what about me?
 #=> I'm a block, but could it be???
 #=> #<Proc:0x00089d64> #<Proc:0x00089c74> #<Proc:0x00089b34>
```

### What is unit testing (in classical terms)? What is the primary technique when writing a test?

The strongest candidates should be quite comfortable with test or behavior driven development.

Unit testing, simply put, is testing methods -- the smallest unit in object-oriented programming. Strong candidates will argue that it allows a developer to flesh out their API before it's consumed by other systems in the application.

The primary way to achieve this is to **assert** that the **actual** result of the method matches an expected result.

```ruby
require "test/unit"

class Brokened
  def uh_oh
    "I needs fixing"
  end
end

class BrokenedTest < Test::Unit::TestCase
  def test_uh_oh
    actual = Brokened.new
    assert_equal("I'm all better!", actual.uh_oh)
  end
end
 #=> Started
 #=> F
 #=> Finished in 0.663831 seconds.
 #=>
 #=>   1) Failure:
 #=> test_uh_oh:11
 #=> <"I'm all better!"> expected but was
 #=> <"I needs fixing">.
 #=>
 #=> 1 tests, 1 assertions, 1 failures, 0 errors
```

Resources:
  - https://gist.github.com/ryansobol/5252653
