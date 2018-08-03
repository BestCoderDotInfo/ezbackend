---
layout: post
author: derek
image: assets/images/feature_default.jpg
featured: false
hidden: false
title: Things You Should And Shouldn't Do In Development With Rails Part 1
date: '2017-10-19 11:00'
comments: true
categories: Rails
keywords: programming, developer should know, ruby, rails, should do, should not, development
excerpt: Rails is the awesome web application framework designed to work with the Ruby programming language. There are a few reasons why Rails is so great for beginners. The first reason is simply that Ruby is really great for beginners, so that’s a big win. Learning to program in Ruby is much easier than in other languages because the language is super flexible and very forgiving, which translates to more time spent absorbing programming fundamentals and less time banging your head against your desk.
---

## Introduce

Rails is the awesome web application framework designed to work with the Ruby programming language. There are a few reasons why Rails is so great for beginners. The first reason is simply that Ruby is really great for beginners, so that’s a big win. Learning to program in Ruby is much easier than in other languages because the language is super flexible and very forgiving, which translates to more time spent absorbing programming fundamentals and less time banging your head against your desk.

Rails is designed with the best practices in mind so it basically guides you into writing awesome code even if you don’t want to (or wouldn’t know how).

This means that if you want to write robust web applications that will scale as you need them to and be easy to maintain as you go forward, Rails is an awesome way to go.


One of the reasons why Rails is so popular with tech startups is that it’s really great for rapid prototyping. You could think of a Rails app you want to build and have it up online and ready to go within a few hours. There are very few frameworks that could make that possible.


What are you thinking of building?!! Really if it’s a web application you can build it with Rails. Just check out some of these examples of websites that are built with Rails: Hulu, Airbnb, Basecamp.

## Things You Should And Shouldn't Do In Development With Rails

### 1. Don’t use default_scope. Ever.

When you would like a scope to be applied across all queries on a model, you can use `default_scope`. See more in the [ActiveRecord Query Guide](http://guides.rubyonrails.org/active_record_querying.html#applying-a-default-scope) and [Rails docs](http://api.rubyonrails.org/classes/ActiveRecord/Scoping/Default/ClassMethods.html#method-i-default_scope).

Instead of use `default_scope`:

```ruby
# app/models/post.rb
class Post < ActiveRecord::Base
  default_scope { where(hidden: false) }
end
```

use explicit scopes:

```ruby
# app/models/post.rb
class Post < ActiveRecord::Base
  scope, :published -> { where(hidden: false) }
end
```

then you can use: `Post.published`.

**But why?**

Two reasons. Both to do with avoiding later confusion and bug hunting.

Adding a default scope affects your model initialization. In the example, `Post.new` is defaulted to `hidden = false` whether you are expecting it or not.

Trying not to use your defined default scope is a pain. To remove the `default_scope` when you don’t need it you have to use the `unscoped` scope (!) which removes all applied conditions including associations.

`Post.first.comments.unscoped` would return every comment in your database, not just those for the first Post.

The explicit use of named scopes is a clearer solution. Using `default_scope` will lead to many hours of bug hunting. Don’t do it to yourself.

### 2.Use Rails’ naming conventions for dates & times

Rails includes the default managed timestamps `updated_at` and `created_at` for ActiveRecord models.

However, on many applications, diving into a schema.rb or migration often reveals something_date as a field name on a model.

Instead of, including the words `date` or `time` in your database columns:

```ruby
class NaughtyMigration < ActiveRecord::Migration[5.1]
  add_column :users, :logged_in_date, :datetime
  add_column :users, :logged_out_time, :date
end
```

use the suffix at for times and on for dates:

```ruby
class AwesomeMigration < ActiveRecord::Migration[5.1]
  add_column :users, :logged_in_at, :datetime
  add_column :users, :logged_out_on, :date
end
```

**But why?**

Including the word time or date in the variable name is redundant and adds to the visual noise of the code. You don’t say `first_name_string`, do you?

Given Rails’ conventions, something like a due_on field lets you know to expect a date. You give instant feedback to anyone reading your code about the expected data stored in the database.

## 3.Rescue specific errors. Avoid rescuing StandardError. Don’t rescue Exception.

There are many built-in error classes in Ruby and Rails. Most of these errors are subclasses of Ruby’s StandardError. You can find more information in the relevant [Ruby docs](http://ruby-doc.org/core-2.4.2/StandardError.html).

Instead of use rescuing `Exception`.

```ruby
def your_method
  # do something
rescue Exception => e
  # saved ALL THE THINGS
end
```

or a non-specific `rescue` that implicitly rescues `StandardError`.

```ruby
def your_method
  # do something
rescue => e
  # saved StandardError and all subclasses
end
```

use `rescue` on a specific named error.


```ruby
def your_method
  # do something
rescue SpecificProblemError => e
  # saved only what you meant to
rescue AnotherProblemError => e
  # saved a different kind of error
end
```

**But why?**

Ruby’s Exception is the parent class to all errors. “Great” you might say, “I want to catch all errors”. But you don’t.

`Exception` includes the class of errors that can occur outside your application. Things like memory errors, or `SignalException::Interrupt` (sent when you manually quit your application by hitting Control-C). These are errors you don’t want to catch in your application as they are generally serious and related to external factors. Rescuing Exception can cause very unexpected behaviour.

`StandardError` is the parent of most Ruby and Rails errors. If you catch `StandardError` you’re not introducing the problems of rescuing Exception, but it is not a great idea. Rescuing all application-level errors might cover up unrelated bugs you don’t know about.

The safest approach is to rescue the error (or errors) you are expecting and deal with the consequences of that error inside the rescue block.

In the event of an unexpected error in your application you want to know that a new error has occurred and deal with the consequences of that new error inside its own rescue block.

Being specific with `rescue` means your code doesn’t accidentally swallow new errors. You avoid subtle hidden errors that lead to unexpected behaviour for your users and bug hunting for you.

>Happy Coding!

Resources:

- [Don’t use default_scope. Ever.](https://andycroll.com/ruby/dont-use-default-scope/)

- [Use Rails’ naming conventions for dates & times](https://andycroll.com/ruby/use-rails-naming-conventions-for-dates-and-times/)

- [Rescue specific errors. Avoid rescuing StandardError. Don’t rescue Exception.](https://andycroll.com/ruby/rescue-specific-errors-avoid-standarderror-do-not-rescue-exception/)

