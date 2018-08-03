---
layout: post
author: derek
image: assets/images/feature_default.jpg
featured: false
hidden: false
title: Ruby Web Application With Rack
date: '2017-07-28 22:00:00'
comments: true
external-url: null
categories: programming
keywords: programming, ruby, rack, web application
excerpt: As you know, popular framework of Ruby is Rails. Rails is a framework for building websites. As such, Rails establishes conventions for easier collaboration and maintenance.
---

>As you know, popular framework of Ruby is Rails. Rails is a framework for building websites.

As such, Rails establishes conventions for easier collaboration and maintenance. Rails like an Full-stack framework and it really strong. So, it have a lot of things no need for an small web application.

There are a few reasons you may want to ditch the framework and go bare-bones:

- **Speed** - Rails can't match just Rack for speed. A "Hello World!" in Rails takes about 1.13 seconds to render. A simple Rack application takes < 10 milliseconds. Rails has some tricks to bring that time down on subsequent requests, but from a cold start using shotgun, those are the numbers.

- **Size** - Rails also uses a lot more memory than a bare-bones Rack application.

- **Simplicity** - Rails creates a lot of files to wade through. Sometimes a single `config.ru` is enough.

- **Future-Proofing** - Rails has a very aggressive release schedule and isn't afraid to deprecate things and leaving vulnerabilities in old versions unpatched. Rack, on the other hand, tends to remain stable for long periods of time.

## What is Rack?

[Rack](https://github.com/rack/rack) is the standard that all Ruby web servers use to interact with Ruby web applications. Rails is a Rack application. Sinatra is a Rack application. So are Volt, Lotus, Grape and just about everything else.

Rack is both a standard for Ruby web apps and also a gem. It's not necessary to use the gem but it does tend to make life easier.

### Hello World With Rack

```ruby
run -> env { [200, {'Content-Type' => 'text/plain'}, ['Hello World!']] }
```

There it is. That's all you need to have a basic Rack app. Stick it in a `config.ru` file, run rackup and visit `http://localhost:9292` and you should see your first bare-bones Ruby web app.

### We'll take it step by step:

- `config.ru` - The `.ru` stands for Rackup file. It's just a Ruby file, but it's designed to be read by a Ruby web server. If you try running ruby config.ru you'll get an error.

- `run` - This is a method defined by the Ruby web server. It expects a Rack application as an argument.

- `->` - This is called a stabby lambda. It's shorthand for a anonymous function. Everything inside `{ }` will be run when the `#call` method is called on it.

- `env` - This is an argument for the lambda. Every Rack application must have a `#call` method that takes one argument. That argument is an environment hash which contains all the information about the HTTP request that was made.

- `[]` - The return value of the `#call` method of a Rack app must be an array with three values.

- `200` - The first element of the returned array is the status code to be returned.

- `{'Content-Type' => 'text/plain'}` - This is a hash of the headers that will be sent with the HTTP response.

- `['Hello World']` - This is the body of the response. This is where HTML would normally go. Notice it's a string wrapped in an array. Whatever object you give here must respond to #each, which a Ruby array does of course.

### Making It Better

```ruby
class MyApp
  attr_reader :request

  def initialize(request)
    @request = request
  end

  def status
    if homepage?
      200
    else
      404
    end
  end

  def headers
    {'Content-Type' => 'text/html', 'Content-Length' => body.size.to_s}
  end

  def body
    content = if homepage?
      "Your IP: #{request.ip}"
    else
      "Page Not Found"
    end

    layout(content)
  end

  private

  def homepage?
    request.path_info == '/'
  end

  def layout(content)
%{<!DOCTYPE html>
<html lang="en">
  <head>

    <meta charset="utf-8">
    <title>Your IP</title>
  </head>
  <body>
    #{content}
  </body>
</html>}
  end
end

class MyApp::Rack
  def call(env)
    request = Rack::Request.new(env)
    my_app = MyApp.new(request)

    [my_app.status, my_app.headers, [my_app.body]]
  end
end

run MyApp::Rack.new
```


>Happy Coding!




