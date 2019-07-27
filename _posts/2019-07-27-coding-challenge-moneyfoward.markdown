---
layout: post
author: derek
image: assets/images/feature_default.jpg
featured: false
hidden: false
title: "Coding Challenge MoneyFoward"
date: 2019-07-27 12:00
comments: true
external-url:
categories: Coding Challenge
keywords: Ruby, Rails, Coding Challenge
excerpt: This is Coding Challenge MoneyFoward. You can use any programming language. I am a Rubyst, so I will write some script with Ruby.
---

## Question

[https://sample-accounts-api.herokuapp.com/](https://sample-accounts-api.herokuapp.com/)

*There are three APIs below. (These are actual API, so you can access them.)
Please create a program with some of these APIs. It should take a user ID as input and return the name, account list and balances of that user. Please note that it is required to use object-oriented programming.*

```
- Web API
  - https://sample-accounts-api.herokuapp.com/users/1
    - レスポンス例 / Sample response
      {
        attributes: {
          id: 1,
          name: "Alice",
          account_ids: [
            1,
            3,
            5
          ]
        }
      }
  - https://sample-accounts-api.herokuapp.com/users/1/accounts
    - レスポンス例 / Sample response
      [
        {
          attributes: {
            id: 1,
            user_id: 1,
            name: "A銀行",
            balance: 20000
          }
        },
        {
          attributes: {
            id: 3,
            user_id: 1,
            name: "C信用金庫",
            balance: 120000
          }
        },
        {
          attributes: {
            id: 5,
            user_id: 1,
            name: "E銀行",
            balance: 5000
          }
        }
      ]  　

  - https://sample-accounts-api.herokuapp.com/accounts/2
    - レスポンス例 / Sample response
      {
        attributes: {
          id: 2,
          user_id: 2,
          name: "Bカード",
          balance: 200
        }
      }
```

## My Answer

First, I write `api_request.rb` to handle requests to  API endpoint.

```ruby
require 'net/http'
require 'json'

module ApiRequest

  def get(url)
    begin
      uri = URI.parse(url)
      http = Net::HTTP.new(uri.host, uri.port)
      http.use_ssl = true
      http.read_timeout = 100
      request = Net::HTTP::Get.new(uri.request_uri)
      response = http.request(request)
      content = JSON.parse(response.body)
    rescue => e
      STDERR.puts "Errors: #{e} \n"
      nil
    end
  end

  module_function :get
  
end
```

Next, I write main application with functions match with question.

```ruby
require './api_request'

class App
  attr_accessor :user_id
  BASE_URL = "https://sample-accounts-api.herokuapp.com".freeze

  include ApiRequest

  def initialize(user_id)
    @user_id = user_id
  end

  def get_user_info
    return unless @user_id
    url = [BASE_URL, 'users', @user_id].join('/').to_s
    STDERR.puts "Starting request to #{url} \n"
    @user_info = get(url)
    STDERR.puts "Response: \n #{@user_info} \n"
  end

  def get_accounts
    url = [BASE_URL, 'users', @user_id, 'accounts'].join('/').to_s
    STDERR.puts "Starting request to #{url} \n"
    @user_accounts = get(url)
    STDERR.puts "Response: \n #{@user_accounts} \n"
  end

  def get_account_info
    @user_info ||= get_user_info
    if @user_info
      account_ids = @user_info['attributes']['account_ids']
      if account_ids
        account_ids.each{|aid| App.get_account_info aid }
      else
        STDERR.puts "This user doesn't have accounts \n"
      end
    else
      STDERR.puts "This user doesn't have accounts \n"
    end
  end

  class << self
    def get_account_info(account_id)
      return unless account_id
      url = [BASE_URL, 'accounts', account_id].join('/').to_s
      STDERR.puts "Starting request to #{url} \n"
      account_info = ApiRequest.get(url)
      STDERR.puts "Response: \n #{account_info} \n"
    end
  end

end
```

Done, it's all! It looks so easy. Happy Conding!

