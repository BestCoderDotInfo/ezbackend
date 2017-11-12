---
title: Sharing models between Rails apps
date: '2017-11-11 18:00'
comments: true
categories: Rails
keywords: programming, ruby, rails, sharing models
---

Note 1: If you have an option to use micro services and/or event-sourcing, go for it! Rails solutions based on shared models and a single shared database, can bring you in a longer perspective more harm than good.

Note 2: This approach is great when your data is tightly coupled and you can’t easily switch from a single app to a distributed model.

## Sharing models between Rails apps – basics

The concept of shared models is really well described by Fabio Akita in following articles:

- [Sharing models between Rails apps – Part 1](http://www.akitaonrails.com/2016/10/03/sharing-models-between-rails-apps-part-1)
- [Sharing models between Rails apps – Part 2](http://www.akitaonrails.com/2016/10/03/sharing-models-between-rails-apps-part-2)

However, his approach towards migrations comes from his really specific use-case (two DBs in a single app – one shared and one private).

## Managing migrations – standard Rails engine approach

Standard Rails engine approach assumes that your migrations will be copied from the engine into the application when you run following command:

```
bundle exec rake engine_name_engine:install:migrations
```

This is great when:

- You have a single “master” application that you want to decompose with engines
- You have multiple applications with separate databases and you want to use business logic from the engine from each of them
- If you want to to have a single “master” application that is supposed to run all the migrations from the engine

However with some benefits, you get a huge (in my case) drawback – when you copy migrations, their timestamp is being changed. It means that if you share same database across multiple applications that also share the same engine, you will end up with a single migration being executed (assuming you install the migrations) from each of the separate applications.

## Single database and no master application

This won’t do if your case is similar to mine, that is:

- Single database
- Multiple applications that need to share same models and scopes
- Migrations should be executed in the first application that is being deployed after the model engine change (not from the “master” app)
- There should not be any patching / adding  code into any of the apps that will use shared models gem

![](/assets/2017-11-12/shared_models.png){:height="100%" width="100%"}

## Keeping your migrations inside your model engine

Solution for such a case (in which all the models are being kept inside the gem) is pretty simple: you just need to append migrations into your apps migrations path without:

- Copying them from the model engine gem
- Changing the timestamp
- Executing the same migrations multiple times

To achieve such a behavior, we will take advantage of how Rails config paths, migrations and initializers work:

- Config paths aren’t bound to the Rails.root directory (which means that they can use files from gems and other locations)
- Config paths are appendable (which means we can add our gem migrations into the app migration list without changing timestamps and copying files)
- Engine initializer allow us to bind this process from the model gem, keeping the apps untouched (they will think that those migrations are theirs)
- Rails migrations execution details are stored in schema_migrations table, so unless executed exactly the same moment (so transactions overlap) a single gem migration will not be executed twice.

All of this comes down few lines of Ruby inside Rails engine engine class (engine_path/lib/engine_name/engine.rb):

```ruby
initializer :append_migrations do |app|
  # This prevents migrations from being loaded twice from the inside of the
  # gem itself (dummy test app)
  if app.root.to_s !~ /#{root}/
    config.paths['db/migrate'].expanded.each do |migration_path|
      app.config.paths['db/migrate'] << migration_path
    end
  end
end
```

## TL;DR – Final solution

`engine_path/lib/engine_name/engine.rb:`

```ruby
module ModelEngine
  class Engine < ::Rails::Engine
    initializer :append_migrations do |app|
      # This prevents migrations from being loaded twice from the inside of the
      # gem itself (dummy test app)
      if app.root.to_s !~ /#{root}/
        config.paths['db/migrate'].expanded.each do |migration_path|
          app.config.paths['db/migrate'] << migration_path
        end
      end
    end
  end
end
```

## Summary

Most of the time sharing models is bad, but there are some cases app data is really tightly coupled together and exposing API with building microservices around it would mean a huge overhead. For such cases model gem with internal migrations might be a great solution.

Warning: If you decide to go that road, please make sure, that:

- Your models are stable
- Your models are slim and without any business logic
- Your models don’t have any callbacks or external dependencies
- If your models have external dependencies, make them model gem dependencies
- Your models are loosely coupled (if you follow Akitas approach with concerns it won’t be hard)
- Your applications are well tested
- Your model gem is well tested
- You don’t use model validations – instead you can use Reform, Dry-Validations or any other solution that allows you to move validations logic out of models
- All the model related things and migrations are inside the model gem
- Migrations from external gems like Devise or FriendlyId are also inside the gem.

>Happy Coding!

Resources:

- [Sharing models between Rails apps – Keeping Rails engine migrations in the engine](https://mensfeld.pl/2017/01/sharing-models-between-rails-apps-keeping-rails-engine-migrations-in-the-engine/)