---
layout: post
author: derek
image: assets/images/feature_default.jpg
featured: false
hidden: false
title: Setup Docker for Rails Application using Postgres (Mysql) with nginx
date: '2017-07-16 09:30:00'
comments: true
external-url: null
categories: programming
keywords: programming, ruby, rails, api, docker, docker-rails, nginx, postgres, mysql
excerpt: Docker is a tool that packages, provisions and runs containers independent
  of the OS...
---

Docker is a tool that packages, provisions and runs containers independent of the OS. Container technology is available through the operating system: A container packages the application service or function with all of the libraries, configuration files, dependencies and other necessary parts to operate. Each container shares the services of one underlying operating system.

Docker was created to work on the Linux platform, but has extended to offer greater support for non-Linux operating systems, including Microsoft Windows and Apple OS X. Versions of Docker for Amazon Web Services (AWS) and Microsoft Azure are available.

In the article, I will write some code to config Docker for Rails with Mysql and Postgres both.

# 1. Config Database

## 1.1 Using Postgres

At `Gemfile`, use Postgres as the database for Active Record:

```
...
gem 'pg'
...
```


Next, at `config/database.yml` :

```
default: &default
  adapter: postgresql
  encoding: utf8
  username: root
  pool: <%= ENV.fetch("RAILS_MAX_THREADS") { 5 } %>
  min_messages: log


development:
  <<: *default
  database: rails_app_<%= Rails.env %>

test:
  <<: *default
  database: rails_app_<%= Rails.env %>

production:
  <<: *default
  min_messages: notice
  url: <%= ENV['DATABASE_URL'] %>
```

##1.2 Using Mysql

At `Gemfile`, Use mysql as the database for Active Record:

```
...
gem 'mysql2'
...
```

Next, at `config/database.yml` :

```
default: &default
  adapter: mysql2
  encoding: utf8
  username: root
  pool: <%= ENV.fetch("RAILS_MAX_THREADS") { 5 } %>
  min_messages: log


development:
  <<: *default
  database: rails_app_<%= Rails.env %>

test:
  <<: *default
  database: rails_app_<%= Rails.env %>

production:
  <<: *default
  min_messages: notice
  url: <%= ENV['DATABASE_URL'] %>

```

# 2. Setup Docker

## 2.1 Dockerfile

```
FROM ruby:2.3-alpine

MAINTAINER Quan Nguyen <quannguyen@bestcoder.info>

RUN apk add --no-cache \
  alpine-sdk \
  tzdata \
  nodejs \
  mariadb-dev \
  postgresql-dev \
  && rm -rf /var/cache/apk/*

RUN gem install bundler

ENV APP_ROOT /opt/app

WORKDIR $APP_ROOT

COPY Gemfile* $APP_ROOT/
RUN bundle install -j4

ARG RAILS_ENV
ENV RAILS_ENV ${RAILS_ENV:-production}
COPY . $APP_ROOT

# Assets precompile
RUN if [ $RAILS_ENV = 'production' ]; then bundle exec rake assets:precompile --trace; fi
# Expose assets for web container
VOLUME $APP_ROOT/public
```


## 2.2 docker-compose.yml

### a. Using Postgres

```
version: '3'
services:
  app:
    build:
      context: .
      args:
        RAILS_ENV: development
    volumes:
      - public_data:/opt/app/public
      - ./log:/opt/app/log
      - assets:/usr/app/${APP_NAME}/public/assets
    environment:
      SECRET_KEY_BASE: 'your_rails_secret_key_base'
      DATABASE_URL: 'postgresql://postgres:postgres@db/rails_app_development'
    links:
      - db
    command: [bundle, exec, rails, server, -b, 0.0.0.0]
    # If your rails app using Puma as web server
    # command: bundle exec rails s Puma -b 0.0.0.0
  db:
    image: postgres
    environment:
      DATABASE: rails_app_development
      USER: postgres
      PASSWORD: postgres
    ports:
      - 5432:5432
    volumes:
      # Mount the DB dumps folder into the container, to be able to create & access database dumps:
      - ./db/dumps:/db/dumps
      # Mount out tmp folder, we might want to have access to something there during development:
      - ./tmp:/tmp
      # Mount our 'restoredb' script:
      - ./bin/restoredb:/bin/restoredb:ro
      # Mount our 'dumpdb' script:
      - ./bin/dumpdb:/bin/dumpdb:ro
      # We'll mount the 'postgres-data' volume into the location Postgres stores it's data:
      - postgres-data:/var/lib/postgresql/data
  web:
    build: containers/nginx
    volumes:
      - ./log:/opt/app/log
      - ./tmp:/opt/app/tmp
      - public_data:/opt/app/public
    ports:
      - "80:80"
    links:
      - app
volumes:
  postgres-data:
    driver: local
  assets:
    external: false
  db_data:
  cache_data:
  public_data:

```

### b. Using mysql

```
version: '3'
services:
  app:
    build:
      context: .
      args:
        RAILS_ENV: development
    volumes:
      - public_data:/opt/app/public
      - ./log:/opt/app/log
      - assets:/usr/app/${APP_NAME}/public/assets
    environment:
      SECRET_KEY_BASE: 'your_rails_app_secret_key_base'
      DATABASE_URL: 'mysql2://db_user:db_pass@db/rails_app_development'
    links:
      - db
    command: bundle exec rails s Puma -b 0.0.0.0
  db:
    image: mysql:5.6
    environment:
      MYSQL_ROOT_PASSWORD: db_pass
      MYSQL_DATABASE: rails_app_development
      MYSQL_USER: db_user
      MYSQL_PASSWORD: db_pass
    volumes:
      - /var/lib/mysql
  web:
    build: containers/nginx
    volumes:
      - ./log:/opt/app/log
      - ./tmp:/opt/app/tmp
      - public_data:/opt/app/public
    ports:
      - "80:80"
    links:
      - app
volumes:
  assets:
    external: false
  db_data:
  cache_data:
  public_data:

```

## 2.3 Config nginx

First, create folder **containers** in your root application:

```
mkdir nginx
```

Then, create `Dockerfile` and `nginx.conf` inside this folder:

**Dockerfile**

```
FROM nginx:stable-alpine

MAINTAINER Quan Nguyen <quannguyen@bestcoder.info>

COPY nginx.conf /etc/nginx/conf.d/default.conf
```

**nginx.conf**

```
upstream app {
  server app:3000;
}

server {
  listen 80 default_server;
  listen 443 default_server;
  server_name _;
  keepalive_timeout 5;
  root /opt/app/public;
  access_log /opt/app/log/nginx.access.log;
  error_log /opt/app/log/nginx.error.log info;

  location / {
    proxy_pass http://app;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header Host $http_host;
  }
}
```

Finally, we will build containers with docker.

# 3. Build and run Rails application with Docker

Download latest **Docker** at [docker.com](https://www.docker.com).

Openter your terminal, in root rails application run:

```
docker-compose build
```

and wait untill successfully.

After build successfully, start your app with Docker:

```
docker-compose up
```

This app will run at: [localhost:80](http://localhost:80)

Setup Database first time for your application:

```
docker-compose run app rake db:create && rake db:migrate && rake db:seed
```

Stop your application:

```
docker-compose stop
```

**Note:**

If you want remove all containers have been build, run:

```
docker-compose rm -v
```


---

Look good! In the next article, I will introcude deploy Rails Application using Docker with AWS.

>Thanks your reading. Happy Coding!