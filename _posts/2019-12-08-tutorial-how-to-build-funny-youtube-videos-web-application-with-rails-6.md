---
title: "[Tutorial] How to build Funny Youtube videos web application with Rails 6,
  Devise, Nokogiri and Bootstrap 4"
layout: post
author: derek
image: assets/funny-movies/1.png
date: 2019-12-08 12:00
comments: true
external-url: null
categories: Ruby on Rails
keywords: Ruby, Rails 6, Web Development, TechTechnology, Bootstrap 4, Devise, Nokogiri,
  crawling, tutorial
excerpt: In this tutorial, you will get a Funny Youtube Videos web application with
  Ruby on Rails up and running on your local server and will push your work into Github.
  How to make it alive?  If you are casually browsing and are maybe interested in
  doing a Ruby on Rails tutorial, stop what you are doing RIGHT NOW and give this
  podcast a listen! It will give you the fuel you need to continue down the web developer
  path; it definitely fueled me.
---

In this tutorial, you will get a Funny Youtube Videos web application with Ruby on Rails up and running on your local server and will push your work into Github. How to make it alive?  If you are casually browsing and are maybe interested in doing a Ruby on Rails tutorial, stop what you are doing RIGHT NOW and give this podcast a listen! It will give you the fuel you need to continue down the web developer path; it definitely fueled me.

Reading and executing tutorials is an amazing way to learn, so roll up your sleeves, take a slug of coffee, and let’s dive in! This tutorial assumes you have Ruby on Rails installed and all other pre-requisites, like github and a text editor and such.

I have put my github of tutorial at the end of the article.

#### 1. Design our web application

Here, I have been design our app with 2 main layouts:

- List youtubes videos
- Share youtube video url

1. Not signed in page
![](/assets/funny-movies/not_signed_in.png){:height="100%" width="100%"}

2. Singed in page
![](/assets/funny-movies/signed_in.png){:height="100%" width="100%"}

3. Share youtube video page
![](/assets/funny-movies/share_a_movie.png){:height="100%" width="100%"}


#### 2. Building our web application

##### 2.1 Setup Rails

As most Ruby on Rails fans might be aware, Rails 6 is coming soon, and bringing a number of eagerly awaited features and changes.  For starters, remember that Rails 6 requires Ruby 2.5+ . So, make sure you have a plan to upgrade your systems accordingly, in case you have not done so already.

We will use  Postgres for our project. So let write:

```ruby
rails new funny-youtube-videos --database=postgresql
```

A whole bunch of cool stuff happens in the terminal:

![](/assets/funny-movies/rails_setup.gif){:height="100%" width="100%"}

And wait amoment until it run out. Then:

```ruby
cd funny-youtube-videos
rails s
```

Now, navigate to `localhost/3000` to make sure everything is working:

Yay! You’re on Rails!

![](/assets/funny-movies/rails_welcome.png){:height="100%" width="100%"}

Next, we will use **Devise** gem for authentication. If you don't know Devise, please go to [https://github.com/plataformatec/devise](https://github.com/plataformatec/devise), and read. This gem is very helpful,it help you save a lot time with how to control authentication for web application.

Add the following line to your Gemfile:

```
gem 'devise'
```

Then run `bundle install`

Next, you need to run the generator:

```
$ rails generate devise:install
```

In our application, User is model contains user's information. Next, we will setup devise with **User**:

```
 rails generate devise User
```

Then run` rails db:migrate`

Now, check our folder `models`  and `db/schema.rb` we got `User` table with: email, password, etc... Restart application server again.

Next step, we need table to storage youtube video  shared by users. So we will create `Post` model with fields under:

- Id : integer
- Title: string
- Share Url: string
- Description: text
- User Id : who share youtube video

That's Ok! We run `rails g model Post title:string share_url:string description:text user:references` . We got migrate file look likes:

```ruby
class CreatePosts < ActiveRecord::Migration[6.0]
  def change
    create_table :posts do |t|
      t.string :share_url, null: false
      t.string :title
      t.text :description
      t.references :user, null: false, foreign_key: true

      t.timestamps
    end
  end
end
```

##### 2.2 Building controllers, views.

In our application, we need PostsController & UsersController. Open our `config/routes.rb` and adding bellow:

```ruby
Rails.application.routes.draw do
  root 'posts#index'
	
  resources :posts
  devise_for :users
	
  post 'access' => 'users#access', as: :create_session
  get 'share' => 'posts#new', as: :share_movie
end
```

In our `UsersController` we will control session with Devise. Help anonymous can sign up & sign in our application.  At  `access` , when anonymous enter correct their & password match with our users, it will create session and they signed in successfully. Opposite, it will create new user with email & password, then signed in.

```ruby
def access
    if valid_user_params?
      user = User.find_for_authentication(email: user_params[:email])
      valid_auth = user && user&.valid_password?(user_params[:password])
      user = User.create(user_params) unless valid_auth
    end
    if user
      sign_in(user)
      flash[:success] = "You have been access successfully!"
    else
      flash[:error] = "Invalid email or password. Please try again!"
    end
    redirect_to root_path
  end
```

At `PostsController` we have

a. `index` list youtube videos

```ruby
@posts = Post.order(id: :desc)
```

We will pagination for our list with gem `will_paginate`. Add our Gemfile `gem 'will_paginate', '~> 3.1.0'` then `bundle install`. Now:

```ruby
@posts = Post.order(id: :desc).paginate(page: params[:page], per_page: 10)
```

b. `create` post when cuser submit their youtube url at share page. I have been write script help we can get  youtube video information such as: title, description. I use `nokogiri` to do it. `nokogiri` is a Rubygem providing HTML, XML, SAX, and Reader parsers with XPath and CSS selector support. We can use it to crawling websites, ... etc.

I write my script in a service object `app/services/post_service.rb`

```ruby
require 'nokogiri'
require 'open-uri'
class PostService
  def initialize(current_user, params)
    @params = params
    @current_user = current_user
  end

  def create
    post = Post.new(@params)
    post.user = @current_user
    if valid_share_url?
      doc = get_doc(@params[:share_url]) if @params[:share_url].present?
      post.title = doc.css('title')&.first&.content&.strip
      post.description = doc.search('#watch-description-text')&.first&.content
    else
      post.errors.add(:share_url, 'Invalid Youtube url')
    end
    post
  end

  private

  def valid_share_url?
    @params[:share_url].match(/^(?:https?:\/\/)?(?:www\.)?youtu(?:\.be|be\.com)\/(?:watch\?v=)?([\w-]{10,})/).present?
  end

  def get_doc(path)
		user_agent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_0) AppleWebKit/535.2 (KHTML, like Gecko) Chrome/15.0.854.0 Safari/535.2"
		Nokogiri::HTML.parse(open(path, {'User-Agent' => user_agent}), nil, 'UTF-8')
	end
end
```

So, now our `create` of PostsController will look like:

```ruby
@post = PostService.new(current_user, post_params).create
respond_to do |format|
	if @post.errors.blank? && @post.save
		format.html { redirect_to root_url, notice: 'Your moive was successfully shared. Thank you!' }
		format.json { render :show, status: :created, location: @post }
	else
		flash[:error] = @post.errors.full_messages.first
		format.html { render :new }
		format.json { render json: @post.errors, status: :unprocessable_entity }
	end
end
```

That's all logic. Now we need implement our web site with front end. We will use [Bootstrap](https://getbootstrap.com/) for front end. It too easy to use. I have been use slim to write html.

In `views/layouts/application.html.slim`

```slim
doctype html
html
  head
    meta content=("text/html; charset=UTF-8") http-equiv="Content-Type" /
    title Funny Movies Site
    = csrf_meta_tags
    = csp_meta_tag
    = stylesheet_pack_tag 'application', media: 'all', 'data-turbolinks-track': 'reload'
    = javascript_pack_tag 'application', 'data-turbolinks-track': 'reload'
  body
    = render 'shared/nav'
    .container
      = yield
      = render 'shared/footer'
```

In `views/shared/_nav.html.slim`

```slim
.d-flex.flex-column.flex-md-row.align-items-center.p-3.px-md-4.mb-3.bg-white.border-bottom.box-shadow
  .col-sm-5
    = link_to root_path do
      h3.font-weight-normal Funny Movies Site
  .col-sm-7 style="display: flex; justify-content: flex-end;"
    - if current_user
      button.btn-light Welcome #{current_user.email}
      | &nbsp;	
      = link_to "Share a movie", share_movie_path, class: 'btn btn-success', role: "button"
      | &nbsp;
      = link_to "Logout", destroy_user_session_path, method: :delete, class: 'btn btn-danger', role: "button"
    - else
      = form_for :user, method: :post, url: create_session_path, html: { class: 'form-inline' } do |f|
        .form-group.mb-2
          = f.email_field :email, class: 'form-control', placeholder: 'email'
        .form-group.mx-sm-3.mb-2
          = f.password_field :password, class: 'form-control', placeholder: 'password'
        = f.submit 'Login / Register', class: 'btn btn-primary mb-2'
```

In home page, also index of PostsController.

`views/posts/index.html.slim`

```slim
.container
  - if @posts.empty?
    h5 No movies
  - @posts.each do |post|
    .card.mb-3
      .row.no-gutters
        .col-md-6
          iframe width="420" height="315" src=(post.share_url.gsub('watch?v=', 'embed/'))
        .col-md-6
          .card-body
            h5.card-title #{post.title&.gsub('- YouTube', '')}
            p.share Share by: #{post.user.email}
            p Description:
            p.card-text #{post.description&.truncate(250)}
  
  = will_paginate @posts
```

In share a youtube page.
`views/posts/new.html.slim`
```slim
.card style=("width: 40rem;margin: auto") 
  .card-header.text-center
    h5 Share a Youtube movie
  .card-body
    = render 'form'
```

And `views/posts/_form.html.slim`

```slim
= form_for @post do |f|
  .form-group.row
    .col-sm-2
      = f.label :share_url, class: 'col-form-label'
    .col-sm-10
      = f.text_field :share_url, class: 'form-control', required: true
  .form-group.row
    .col-sm-2
    .col-sm-10
      = f.submit 'Submit', class: 'btn btn-primary btn-lg btn-block'
```

##### 2.3 Make it alive!

I recommend guys using Heroku  for hostingRails application. [Heroku](https://heroku.com)  is a platform as a service (PaaS) that enables developers to build, run, and operate applications entirely in the cloud.  Register now!

This [artice](https://devcenter.heroku.com/articles/getting-started-with-rails5) will help you deploy our website to Heroku.

![](/assets/funny-movies/a_few_moments_later.jpg){:height="100%" width="100%"}

### Our website alive now: [https://funny-movies-site.herokuapp.com](https://funny-movies-site.herokuapp.com/)

#### We have awesome funny youtube videos with Rails 6.

![](/assets/funny-movies/1.png){:height="100%" width="100%"}

![](/assets/funny-movies/2.png){:height="100%" width="100%"}

![](/assets/funny-movies/3.png){:height="100%" width="100%"}

## Conclusion
In this tutorial, we can build Sharing Youtube Videos with Rails 6 and gems: Devise, nokogiri, will_paginate and Bootstrap 4. Learn how about crawling website.  And deploy our website to heroku. Make it alive!

[Github Funny youtube videos](https://github.com/dereknguyen269/funny-movies-site)

That's all. Thank for reading and enjoy coding!