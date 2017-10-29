---
title: What is the difference between extend and include in Ruby
date: '2017-10-29 09:00'
comments: true
categories: Ruby
keywords: programming, developer should know, ruby, rails, include, extend
excerpt: As you know, we always using include and extend when programming with Ruby. Sometimes we don't really why do we use it and when do we use it?
---

As you know, we always using include and extend when programming with Ruby. Sometimes we don't really why do we use it and when do we use it?

`include` mixes in specified module methods as instance methods in the target class.

`extend` mixes in specified module methods as class methods in the target class.

Given the following class definitions:

```ruby
module ReusableModule
  def module_method
    puts "Module Method: Hi there! I'm a module method"
  end
end

class ClassThatIncludes
  include ReusableModule
end
class ClassThatExtends
  extend ReusableModule
end
```

Here’s how `ClassThatIncludes` behaves:

```ruby
# A class method does not exist
>> ClassThatIncludes.module_method
NoMethodError: undefined method `module_method' for ClassThatIncludes:Class

# A valid instance method exists
>> ClassThatIncludes.new.module_method
Module Method: Hi there! I'm a module method
=> nil
```

Here’s how `ClassThatExtends` behaves:

```ruby
# A valid class method exists
>> ClassThatExtends.module_method
Module Method: Hi there! I'm a module method
=> nil

# An instance method does not exist
ClassThatExtends.new.module_method
NoMethodError: undefined method `module_method' for #<ClassThatExtends:0x007ffa1e0317e8>
```

We should mention that `object.extend ExampleModule` makes `ExampleModule` methods available as singleton methods in the object.


>Happy Coding!

Resource: 
- [21 Essential Ruby Interview Questions](https://www.toptal.com/ruby/interview-questions)
