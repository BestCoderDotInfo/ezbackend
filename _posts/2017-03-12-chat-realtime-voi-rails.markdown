---
layout: post
author: derek
image: assets/images/feature_default.jpg
featured: false
hidden: false
title: Chat Realtime Với Rails
date: 2017-03-12 00:00
comments: true
external-url:
categories: Ruby-on-Rails
keywords: ruby, rails, ruby on rails, chat realtime, chat realtime voi rails
excerpt: Cũng như các bạn đã biết, đã số các website chat bây giờ hoặc có tính tương tác trực tiếp với người dùng cao thì thường dùng nodejs để làm real time. Hôm nay , mình cũng xin giới thiệu tới các bạn cách xây dựng 1 ứng dụng chat real time chỉ hoàn toàn sử dụng ruby on rails.
---
>Cũng như các bạn đã biết, đã số các website chat bây giờ hoặc có tính tương tác trực tiếp với người dùng cao thì thường dùng nodejs để làm real time.

Hôm nay , mình cũng xin giới thiệu tới các bạn cách xây dựng 1 ứng dụng chat real time chỉ hoàn toàn sử dụng ruby on rails.

Sau khi khởi tạo app rails thì chúng ta thêm vào Gemfile:

```ruby
gem 'private_pub'
gem 'thin'
```

chạy lệnh bundle install .
chạy tiếp lệnh :

`rails g private_pub:install`

`rackup private_pub.ru -s thin -E production`

Chú ý phải thêm đoạn sau vào application.js  của app :

```
//= require private_pub
```

Thêm đoạn sau vào layout của app :

```
<%= subscribe_to "/messages/new" %>
```

`/messages/new` là url ví dụ của mình .

ở phần action của  controller  , ví dụ sau khi bạn create message, action create  sẽ trả về response là json hoặcchuyển trang. Nhưng ở đây chúng ta sẽ tạo 1 file `create.js.erb` tương ứng với action đó. Và  file đó  sẽ chứa đoan sau :

```
<% publish_to "/messages/new" do %>
  $("#chat").append("<%= j render(@messages) %>");
<% end %>
```

đây là đoạn code mà sau khi bạn create sẽ tự động render ra và realtime có nghĩa là ai đang lướt app của bạn đều sẽ nhìn thấy đoạn chat mới đó.
