---
layout: post
title: Upload Hình Ảnh Lên S3 Với Carrierwave Cho Ứng Dụng Ruby On Rails
date: 2017-03-03 12:00
comments: true
external-url:
categories: Ruby-on-Rails
tags: amazone s3, s3, upload with rails, carrierwave
keywords: amazone s3, s3, upload with rails, carrierwave
excerpt: Đối với ứng dụng web sử sụng Rails hay bất cứ ngôn ngữ hay framework nào khác thì việc lưu trữ các hình ảnh đổi lúc sẽ phát sinh nhiều bất tiện. Với dịch vụ S3 mà Amazon cung cấp chúng ta có thể lưu trữ và các hình ảnh đó sẽ tự động resize với kích thước mà chúng ta không cần làm gì hết.
image:
  feature: ../../assets/aws-s3-carrierwave-on-rails.jpg
  credit: 
  creditlink:
---
>Đối với ứng dụng web sử sụng Rails hay bất cứ ngôn ngữ hay framework nào khác thì việc lưu trữ các hình ảnh đổi lúc sẽ phát sinh nhiều bất tiện. Với dịch vụ S3 mà Amazon cung cấp chúng ta có thể lưu trữ và các hình ảnh đó sẽ tự động resize với kích thước mà chúng ta không cần làm gì hết. Để sử dụng S3 cho Rails app chúng ta cần thực hiện theo các bước sau:

>Đầu tiên chúng ta cần thêm vào `Gemfile` :

```
gem 'carrierwave'

gem 'fog' # Support storing file in AWS
```

[Carrierwave](https://goo.gl/XWJHJI) là một gem sẽ giúp chúng ta đơn giản hóa việc upload hình ảnh, file một cách dễ dàng và nhanh chóng.

>Tạo một file `uploader`

`rails generate uploader CoverImage` ở đây tui dùng để upload cho các hình ảnh của cover chẳng hạn. Và chúng ta tạo được một file mới:

`app/uploader/cover_image_uploader.rb`

>Cấu hình trong file uploader mà ta vừa tạo:

Nó nên có cấu trúc như thế này :

```
class CoverImageUploader < CarrierWave::Uploader::Base

  storage :fog

  
  ...

  def store_dir
    "#{Rails.env}/#{model.class.to_s.underscore}/#{mounted_as}/#{model.id}"
  end

  def extension_white_list
    %w(jpg jpeg gif png)
  end

  ...
end
```


>Chỉnh sử lại model tương ứng 

```
class Post < ActiveRecord::Base
  mount_uploader :cover_image, CoverImageUploader

  ...
end
```


>Cấu hình upload lên S3

Ở đây chúng ta sử dụng gem [fog](http://goo.gl/dKmvS1) để hỗ trợ lưu trữ các file ở AWS, Google, Local và Rackspace.

Tạo file:

`config/initializers/carrierwave.rb`

với cấu trúc như sau :

```
CarrierWave.configure do |config|
  config.fog_credentials = {
    provider: 'AWS',
    aws_access_key_id: ENV['AWS_KEY'],
    aws_secret_access_key: ENV['AWS_SECRET_KEY'],
    region: 'ap-southeast-1'
  }
  config.fog_directory = ENV['S3_BUCKET_NAME']
end
```

**region** : nên là một trong [ ‘eu-west-1’, ‘us-east-1’, ‘ap-southeast-1’, ‘us-west-1’, ‘ap-northeast-1’ ]

**fog_directory** : là tên Bucket Name ở S3

>Happy Coding!
