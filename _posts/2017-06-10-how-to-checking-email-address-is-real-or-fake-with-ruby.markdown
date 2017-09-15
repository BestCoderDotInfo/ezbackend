---
title:  How to checking email address is real or fake with Ruby ?
date: 2017-06-10 12:00
comments: true
external-url:
categories: ruby
keywords: ruby, rails, checking email, email detected, ruby gem
excerpt: Email Detected is a simple tool for verifying an email address exists. It's free and quite easy to use 😄 .
---
>Email Detected is a simple tool for verifying an email address exists. It's free and quite easy to use 😄 .

## Introcude

Many times as developers we were putting validation statements for checking email addresses format. This gem will complete your existing setups with validator that actually connects with a given mail server and asks if the address in question exists for real.

## Installation

### Ruby on Rails
Add this line to your application's Gemfile:

```ruby
gem 'email_detected'
```

And then execute:

```
$ bundle
```


### Only Ruby

```
$ gem install email_detected
```

## Usage

To get info about realness of given email address, email_detected connects with a mail server that email's domain points to and pretends to send an email. Some smtp servers will not allow you to do this if you will not present yourself as a real user.

This only needs to be something the receiving SMTP server. We aren't actually sending any mail.

First thing you need to set up is placing something like this either in initializer or in `application.rb` file:

```ruby
EmailDetected.config do |config|
  config.verifier_email = "youremail@email.com"
end
```

Then just put this in your model e. g:

```ruby
validates_exist_email_of :email
```

Or - if you'd like to use it outside of your models: 

```ruby
EmailDetected.exist?(youremail)
```

This method will return with status `true || false` and message look like:

```javascript
{:status=>true, :message=>"The email address has already been registered."}
```

```javascript
{:status=>false, :message=>"The email address invalid."} 
```

or will throw an exception with nicely detailed info about what's wrong.

See more at [Email Detected](https://github.com/minhquan4080/email_detected)

>Happy Coding !