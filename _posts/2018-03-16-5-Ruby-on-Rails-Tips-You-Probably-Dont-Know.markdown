---
title: 5 Ruby on Rails Tips You Probably Don’t Know
date: '2018-03-16 17:00'
comments: true
categories: Rails
keywords: programming, ruby, rails, tips, tricks
---

## Hash#dig
I didn’t meet this in anyone’s code for 7 years of RoR development and discovered recently. And it’s obvious why :-) The first Ruby version I worked with was 1.8, but this method was introduced in 2.3.

How many times did you do something like this:

```ruby
... if params[:user] && params[:user][:address] && params[:user][:address][:somewhere_deep]
```

Think of dig as kind of safe navigation operator `&`. but for Hash objects. So now you could rewrite such things:

```ruby
... if params.dig(:user, :address, :somewhere_deep)
```


## Object#presence_in

This one I found in the very good article about Query Objects in Ruby on Rails. It allows you to replace conditionals (often ternary) with single method call when you don’t actually need a boolean result of inclusion check, but a checked object itself. For example, your code looks like:

```ruby
sort_options = [:by_date, :by_title, :by_author]
...
sort = sort_options.include?(params[:sort])
  ? params[:sort]
  : :by_date
# Another option
sort = (sort_options.include?(params[:sort]) && params[:sort]) || :by_date
```

Doesn’t this look better?

```ruby
params[:sort].presence_in(sort_options) || :by_date
```

## Module#alias_attribute

Well, I found this very useful when I worked on a project with a legacy database. It had a table with weird column names like `SERNUM_0`, `ITMDES1_0` and other. We mapped this table to an `ActiveRecord` model and instead of writing queries and scopes on it like `WeirdTable.where(SERNUM_0: ‘123’)` , we came up to using `alias_attribute` since it doesn’t just generate getter and setter (as well as a predicate method), but works in queries as well:


```ruby
alias_attribute :name, :ITMDES1_0
...
scope :by_name, -> (name) { where(name: name) }
```

## Object#presence

This one is more popular than others. There’s pretty good explanation on ApiDock. So, `object.presence` is equivalent to:

```ruby
object.present? ? object : nil
```

## Module#delegate

Still rarely used by most of the developers for some reason. The main purpose of this method is loose coupling and following the Law of Demeter. One of the good articles on this theme coming to mind is Avdi Grimm’s “Demeter: It’s not just a good idea. It’s the law.” Also, check out short Rails Best Practices article on utilizing delegate in context of applying the Law of Demeter. The snippet below describes this as well:

```ruby
class Profile < ApplicationRecord
  belongs_to :user
  delegate :email, to: :user
end
...
profile.email # equivalent to profile.user.email
```

>Happy Coding!

Resources:
[5 Ruby on Rails Tips You Probably Don’t Know](https://hackernoon.com/5-ruby-on-rails-tips-you-probably-dont-know-8b80b4a0890f)