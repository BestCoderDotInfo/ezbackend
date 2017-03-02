---
layout: post
title: Setting Up Rails Application With Action Mailer In Development And Sendgrid In Production Using Heroku
date: 2017-02-27 12:00
comments: true
external-url: 
categories: Ruby-on-Rails
---
>Here is a demo of how I created an email with user-mailer in rails

First off, from the [Rails Guide](http://guides.rubyonrails.org/action_mailer_basics.html) generate a UserMailer.

Then in then `user_mailer.rb` create a function that sends the email :

```
class UserMailer < ActionMailer::Base
   default from: "me@gmail.com" 

   def welcome_email(user)
        @user = user
        @url = 'http://www.google.com'
        mail(to: @user.email, subject: 'Test welcome email')
   end
end
```

Pretty simple right?  Now we need to create the view.  This is where you can design the email and whatever text you want to put in it. Create a file in the `app/views/user_mailer` folder.  The name of the file should match the name of the function you defined in the user_mailer.rb file.  Mine is `welcome_email.html.erb`

```
<!DOCTYPE html>
<html>
<head>
    <meta content='text/html; charset=UTF-8' http-equiv='Content-Type' />
</head>
<body>
    <h1>Hello, </h1>
    <p>
        Welcome to our site !
    </p>
        To login to the site, just follow this link : <%= @url %>.
</body>
</html>
```

Now that we have the function to send the email, and the text of the email itself,  we just need to add a line in the controller to have it send the email when we want it to.

You can call the mailer function wherever you want to send the email, in this case we can just send it when a new user signs up

```
class UserController < ApplicationController
    def create
        @user = User.new(params[:user])

        response_to do |format|
            if @user.save
                UserMailer.welcome_email(@user).deliver
            end
        end
    end
end
```

Now we need to set up the config files for the email address we will use to send emails. For development, in the `config/environment/development.rb` file:

```
config.action_mailer.raise_delivery_errors = true
config.action_mailer.default_url_options = {:host => 'localhost:3000'}

config.action_mailer.delivery_method = :smtp
config.action_mailer.smtp_settings = {
    address: "smtp.gmail.com",
    port: "587",
    domain: "gmail.com",
    authentication: "plain",
    enable_starttls_auto: true,
    user_name: 'your-email-username',
    password: 'your-email-password'
}
```

Set the `config.action_mailer.default_url_options` to `localhost:3000` in development, or whatever port you are using to test.  In production, you would set it to the url of your production app.  Then set up the email smtp settings according to your email provider.  I’ve shown the settings for gmail here.

So great, now we have it set up so a new user will receive an email, at least in development.  Now getting this working in production can be slightly harder.  But I’m going to walk through one super easy solution of how to set this up using [Heroku](https://www.heroku.com/) and [SendGrid](http://sendgrid.com/), but there are probably a million different ways to do this depending on your hosting, your email, etc.

So the first thing to do before moving to production is to protect the sensitive email sign in information that we just put in the `config/environment/development.rb`.  The easiest way to do this is to use environment variables, and a simple way to do that is to use the [figaro gem](https://github.com/laserlemon/figaro).  I wrote a quick post on how to setup environment variables with the [figaro gem here](https://howilearnedrails.wordpress.com/2013/08/05/environment-variables-with-the-figaro-gem/), so take a minute and set that up in your app first.  This will keep your email login info from getting pushed to github, in case you are using a public repository.  It should then look something like this:


```
user_name: ENV['GMAIL_USERNAME_DEV']
password: ENV['GMAIL_PASSWORD_DEV']
```

So, assuming you have your app pushed to heroku for production, the next thing we need to do is setup an email add-on for your heroku app.  Of course you could fully create your own email solution, but if you know how to do that then you wouldn’t be reading this blog, would you? So, for a simpler solution, first you can choose one of the [email related add-ons](https://addons.heroku.com/#email-sms) from the [heroku add-ons page](https://addons.heroku.com/).  I happen to use [SendGrid](https://addons.heroku.com/sendgrid) as I think its one of the simplest to set up and use, so that’s what I’ll go through here.

Add SendGrid to your heroku app by running

`heroku addons:add sendgrid:starter`

Now SendGrid will have automatically generated a username and password that you can see by running

`heroku config:get SENDGRID_USERNAME`
`heroku config:get SENDGRID_PASSWORD`

or simply running heroku config will show you all the heroku config variables.

Next, go to `config/environment/production.rb` and add our email settings there:


```
config.action_mailer.smtp_settings = {
    address: "smtp.sendgrid.ner",
    port: "587",
    authentication: "plain",
    enable_starttls_auto: true,
    user_name: ENV['SENDGRID_USERNAME'],
    password: ENV['SENDGRID_PASSWORD']
}
```

So you can see that SendGrid has set up environment variables on heroku for us, giving us the **SENDGRID_USERNAME** and **SENDGRID_PASSWORD**.  This is the same idea that we just did in development using the figaro gem.

Ok, so now push to heroku, migrate the database, and see if the mailer is working in production.  Not too difficult right?  Stay tuned for a follow up post where I’ll describe how to send email asynchronously with [sidekiq](http://sidekiq.org/) and [redis](http://redis.io/).

>Happy Coding!