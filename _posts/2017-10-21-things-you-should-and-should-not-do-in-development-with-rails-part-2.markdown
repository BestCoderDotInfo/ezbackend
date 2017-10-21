---
title: Things You Should And Shouldn't Do In Development With Rails Part 2
date: '2017-10-21 18:00'
comments: true
categories: Rails
keywords: programming, developer should know, ruby, rails, should do, should not, development
excerpt: Use Time.zone.now instead of Time.now .Replace Complex Creation with Factory Method
---

## 1.Use Time.zone.now instead of Time.now

### Why?

Using default Ruby `Time`, `Date` and `DateTime` classes will not show times in the time zone specified by `config.time_zone` in `application.rb`.

```ruby
Time.zone = "Alaska"
Time.now
Date.today
```
These show the local time, not the time in Alaska (unless you're already in Alaska).

### Refactor

You should instead use ActiveSupport methods of `Time.zone` to pickup the Rails time zone.

```ruby
Time.zone.now
Time.zone.today
```

This is well-described in [The Exhaustive Guide to Rails Time Zones](http://danilenko.org/2012/7/6/rails_timezones/). This should be easy to write a `rails_best_practice` rule to implement.

Time zone bugs are particularly tricky when the production server is set to a different time zone (often UTC) than the development machine. Using 'Time.zone` avoids this breakdown of dev/prod parity.

## 2.Replace Complex Creation with Factory Method

### Bad

```ruby
class InvoicesController < ApplicationController
  def create
    @invoice = Invoice.new(params[:invoice])
    @invoice.address = current_user.address
    @invoice.phone = current_user.phone
    @invoice.vip = (@invoice.amount > 1000)

    if Time.now.day > 15
      @invoice.delivery_time = Time.now + 2.month
    else
      @invoice.delivery_time = Time.now + 1.month
    end

    @invoice.save
  end
end
```

The logic to create an invoice is too complex, it makes InvoicesController a bit difficult to read. And the controller should not know too much things about how to create a model, we should move the the logic of creating an invoice into the Invoice model.

### Good

```ruby
class Invoice < ActiveRecord::Base
  def self.new_by_user(params, user)
    invoice = self.new(params)
    invoice.address = user.address
    invoice.phone = user.phone
    invoice.vip = (invoice.amount > 1000)

    if Time.now.day > 15
      invoice.delivery_time = Time.now + 2.month
    else
      invoice.delivery_time = Time.now + 1.month
    end
  end
end

class InvoicesController < ApplicationController
  def create
    @invoice = Invoice.new_by_user(params[:invoice], current_user)
    @invoice.save
  end
end
```

Now we define a new_by_user method in Invoice model, it takes charge of the invoice creation. So the InvoicesController can just call the new_by_user method to build the invoice object.

Keep in mind the principle "Skinny Controller, Fat Model"

## 3.Use model association

### Bad

```ruby
class PostsController < ApplicationController
  def create
    @post = Post.new(params[:post])
    @post.user_id = current_user.id
    @post.save
  end
end
```

In this example, `user_id is` assigned to `@post` explicitly. It's not too big problem, but we can save this line by using model association.

### Good

```ruby
class PostsController < ApplicationController
  def create
    @post = current_user.posts.build(params[:post])
    @post.save
  end
end

class User < ActiveRecord::Base
  has_many :posts
end
```

We define the association that user has many posts, then we can just use `current_user.posts.build` or `current_user.posts.create` to generate a post, and the current_user's id is assigned to `the user_id` of the post automatically by activerecord.

## 4.Nested Model Forms

### Bad

```ruby
class Product < ActiveRecord::Base
  has_one :detail
end

class Detail < ActiveRecord::Base
  belongs_to :product
end

<% form_for :product do |f| %>
  <%= f.text_field :title %>
  <% fields_for :detail do |detail| %>
    <%= detail.text_field :manufacturer %>
  <% end %>
<% end %>

class ProductsController < ApplicationController
  def create
    @product = Product.new(params[:product])
    @detail = Detail.new(params[:detail])

    Product.transaction do
      @product.save!
      @detail.product = @product
      @detail.save
    end
  end
end
```

Product and Detail models are one-to-one association, we want to create a product with a new manufacturer. In this example, we create a product and a detail object, and associate them in a transaction. But why we do the complex creation in the controller, Product model should take charge of creating a detail.

### Good

```ruby
class Product < ActiveRecord::Base
  has_one :detail
  accepts_nested_attributes_for :detail
end

<% form_for :product do |f| %>
  <%= f.text_field :title %>
  <% f.fields_for :detail do |detail| %>
    <%= detail.text_field :manufacturer %>
  <% end %>
<% end %>

class ProductsController < ApplicationController
  def create
    @product = Product.new(params[:product])
    @product.save
  end
end
```

We add the accepts_nested_attributes_for in Product model, so Product model can handle the creation of Detail model. That's so cool!

And we can also use accepts_nested_attributes_for to simplify the one-to-many association creation.


```ruby
class Project < ActiveRecord::Base
  has_many :tasks
  accepts_nested_attributes_for :tasks
end

class Task < ActiveRecord::Base
  belongs_to :project
end

<% form_for @project do |f| %>
  <%= f.text_field :name %>
  <% f.fields_for :tasks do |tasks_form| %>
    <%= tasks_form.text_field :name %>
  <% end %>
<% end %>
```

>Happy Coding!

Resources:

- [Rails Best Practices](https://rails-bestpractices.com)
