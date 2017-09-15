---
title:  How to speed up your rails application ?
date: 2017-05-20 09:00
comments: true
external-url:
categories: rails
keywords: ruby, rails, speed up, increase performance, faster rails
excerpt: Some of the best ways to make your Rails app faster are dead-simple and quite often missed. The tragic part is that many of these problems come up over and over again and can be fixed quickly and incrementally. And the benefits of each of these improvements can be significant.
---

>The tragic part is that many of these problems come up over and over again and can be fixed quickly and incrementally. And the benefits of each of these improvements can be significant.

This post focuses on four quick fixes you can apply to your Rails app that require relatively little effort and often deliver a huge payoff.

## What makes your Rails application slow?

While there can be many reasons behind an application’s slowness, database queries usually play the biggest role in an application’s performance footprint. Loading too much data into memory, N+1 queries, lack of cached values, and the lack of proper databases indexes are the biggest culprits that can cause slow requests.

## 1. Group Your SQL Queries

You've probably heard of N+1 queries.

That's when an innocuous looking line of code (like the following) triggers many more queries than you expect.

```ruby
Client.limit(10).map(&:address)  
```

In this case, instead of doing a single query with a join for the addresses, you make a query for 10 clients, and then another query to get the address for each of the clients.

```ruby
SELECT * from "clients" LIMIT 10;  
SELECT * from "addresses" WHERE "id" = 7;  
SELECT * from "addresses" WHERE "id" = 8;  
SELECT * from "addresses" WHERE "id" = 10;  
SELECT * from "addresses" WHERE "id" = 12;  
SELECT * from "addresses" WHERE "id" = 13;  
SELECT * from "addresses" WHERE "id" = 15;  
SELECT * from "addresses" WHERE "id" = 16;  
SELECT * from "addresses" WHERE "id" = 17;  
SELECT * from "addresses" WHERE "id" = 21;  
```

You might have heard about this problem, and that the solution to the problem is the Rails eager loading API:

```ruby
Client.includes(:address).limit(10) 
```

This produces just two queries:

```ruby
SELECT * from "clients" LIMIT 10;  
SELECT * from "addresses" WHERE "client_id" IN (7, 8, 10, 12, 13, 15, 16, 17, 21);
```

## 2. Use Basic Fragment Caching

One of the easiest ways to get a decent performance boost out of a Rails app is to find ways to cache expensive HTML or JSON fragments.

Looking into `static/index.html.erb`, I see a few truly dynamic bits, like this:

```
<%= render partial: 'shared/flash',  
    locals: { flash: flash, class_name: "banner" } %>
```

But for the most part, it's a large template whose dynamic bits look like this:

```
<%= link_to "Sign Up For Free", signup_path, class: 'signup' %>
```

If you're like me, even talking about caching feels daunting. But Rails makes it really easy!

Just wrap the area of the template that you want to cache with a cache block.

```
<% cache(action_suffix: "primary") do %>  
<section class="hero">  
  <div class="container">
    ...
  </div>
</section>

<section class="data">  
  <div class="container">
    ...
  </div>
</section>  
...
<% end %>
```

When using fragment caching, remember three things:

- Pick a key that describes the fragment you are caching. 
You can use action_suffix, as in this example, if the key is unique only inside of this action. (You can also use an Active Record object as your cache key, which is quite convenient and simple in the right situations.)

- The easiest caching backend is memcached. This is because 
memcached automatically expires keys that haven't been used in a while (an "LRU" or "least recently used" expiration strategy), and cache expiration is the hardest part of a good caching strategy.

- Focus on the big spenders. It's tempting to spend a lot of 
time caching all of your HTML and trying to identify good cache keys for everything. In practice, you can get big wins by just caching expensive fragments that have easy cache keys (either because the template is relatively static, or because it's derived from an Active Record object, which has built-in caching support).

## 3. Eliminate Memory Bloat

Even if you're using a performance monitoring tool, it's very easy for the cost of memory bloat to go unnoticed.

That's because endpoints that create a huge amount of objects aren't necessarily the endpoints that experience the GC pauses. GC pauses are spread out throughout your entire app, so identifying the root cause can be tricky.

You can use gems like Sam Saffron's [memory_profiler](https://github.com/SamSaffron/memory_profiler) or Koichi Sasada's [allocation_tracer](https://github.com/ko1/allocation_tracer) to try to track down which actions are generating objects, and there are even Rack middlewares you can install that will collect the information automatically.

## 4. Move Third-Party Integration to Workers

If you're using Rails 4.2 or newer, [ActiveJob](http://guides.rubyonrails.org/active_job_basics.html) makes the process even simpler. Rails now bakes the notion of background jobs into the framework, complete with generators to get your started and seamless integration into the rest of Rails. I strongly recommend it.

## 5. Indexed Your Database

Missing database indexes on foreign keys and commonly searched columns or values that need to be sorted can make a huge difference. The missing index is an issue that is not even noticeable for tables with several thousand records. However, when you start hitting millions of records, the lookups in the table become painfully slow.

### The role of database indexes

When you create a database column, it’s vital to consider if you will need to find and retrieve records based on that column.

For example, let’s take a look at the internals of [Semaphore](https://semaphoreci.com/). We have a Project model, and every project has a name attribute. When someone visits a project on Semaphore, e.g. `https://semaphoreci.com/renderedtext/test-boosters`, the first thing we need to do in the projects controller is to find the project based on its name — test-boosters. 

```ruby
project = Project.find_by_name(params[:name])
```

Without an index, the database engine would need to check every record in the projects table, one by one, until a match is found.

However, if we introduce an index on the ‘projects’ table, as in the following example, the lookup will be much, much faster.

```ruby
class IndexProjectsOnName < ActiveRecord::Migration
  def change
    add_index :projects, :name
  end
end
```

A good way to think about indexes is to imagine them as the index section at the end of a book. If you want to find a word in a book, you can either read the whole book and find the word, or your can open the index section that contains a alphabetically sorted list of important words with a locator that points to the page that defines the word.

### What needs to be indexed?

A good rule of thumb is to create database indexes for everything that is referenced in the `WHERE`, `HAVING` and `ORDER BY` parts of your SQL queries.

**Indexes for unique lookups**
Any lookup based on a unique column value should have an index.

For example, the following queries:
```ruby
User.find_by_username("shiroyasha")
User.find_by_email("support@semaphoreci.com")
```

will benefit from an index of the username and email fields:

```ruby
add_index :users, :username
add_index :users, :email
```

**Indexes for foreign keys**

If you have belongs_to or has_many relationships, you will need to index the foreign keys to optimize for fast lookup.

For example, we have the branches that belong to projects:

```ruby
class Project < ActiveRecord::Base
  has_many :branches
end

class Branch < ActiveRecord::Base
  belongs_to :project
end
```

For fast lookup, we need to add the following index:

```ruby
add_index :branches, :project_id
```

For polymorphic associations, the owner of the project can either be a User or an `Organization`:

```ruby
class Organization < ActiveRecord::Base
  has_many :projects, :as => :owner
end

class User < ActiveRecord::Base
  has_many :projects, :as => :owner
end

class Project < ActiveRecord::Base
  belongs_to :owner, :polymorphic => true
end
```

We need to make sure that we create a double index:

```ruby
# Bad: This will not improve the lookup speed

add_index :projects, :owner_id
add_index :projects, :owner_type

# Good: This will create the proper index

add_index :projects, [:owner_id, :owner_type]
```

**Indexes for ordered values**
Any frequently used sorting can be improved by using a dedicated index.

For example:
```ruby
Build.order(:updated_at).take(10)
```

can be improved with a dedicated index:

```ruby
add_index :updated_at
```

**Should I always use indexes?**

While using indexes for important fields can immensely improve the performance of your application, sometimes the effect can be negligible, or it can even make your application slower.

For example, tables that have elements that are frequently deleted can negatively impact the performance of your database. Huge tables with many millions of records also require more storage for your indexes.

Always be concious about the changes you introduce in your database, and if in doubt, be sure to base your decisions on real world data and measurements.

Sources:

- [Faster Rails: Is Your Database Properly Indexed?](https://semaphoreci.com/blog/2017/05/09/faster-rails-is-your-database-properly-indexed.html)

- [4 Easy Ways to Speed Up Your Rails App](http://blog.skylight.io/4-easy-ways-to-speed-up-your-rails-app/)
