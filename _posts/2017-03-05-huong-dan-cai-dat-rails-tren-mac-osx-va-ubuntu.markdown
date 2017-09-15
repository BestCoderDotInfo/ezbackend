---
title: Hướng Dẫn Cài Đặt Rails Trên MAC OSX Và Ubuntu
date: 2017-03-05 06:00
comments: true
external-url:
categories: Ruby-on-Rails
keywords: rails, ruby on rails, cai dat ruby on rails, mac osx, cai dat rails tren mac, cai dat rails tren osx, cai dat rails tren ubuntu
excerpt: Trước tiên mình xin nói thẳng, MAC OSX là một môi trường hoàn hảo để cài lập trình Ruby on Rails. Nếu có điều kiện thì các bạn hãy trang bị ngay cho bản thân một con Macbook để làm việc hiệu quả hơn. Còn không thì các bạn phải cài đặt hệ điều hành Ubuntu để làm hệ điều hành chính, hoặc máy ảo chạy Ubuntu cho các bạn thích xài Window nhưng vẫn muốn lập trình Ruby. Nhưng các thư viện của Rails hỗ trợ tốt nhất vẫn dành cho MAC OS. Cho nên đôi khi chúng không sử dụng được trên Ubuntu.
---
Khi lập trình Ruby on Rails các bạn phải tập làm quen với Terminal ( Cmd ) để gõ các lệnh.

# Cài Đặt Ruby on Rails trên MAC OSX

1 . Cài Đặt Homebrew :

Đầu tiên các bạn cần phải cài đặt [Homebrew](http://brew.sh/) là một công cụ tiện ích cần thiết đối với bất kì lập trình viên nào khi lập trình trên MAC OSX. Homebrew giúp cài thêm các phần mềm , thư viện có trong Unix, Linux mà MAC OSX không có. Và Homebrew cũng được viết bằng ngôn ngữ Ruby 🙂

Bật Terminal lên và gõ : 

```bash
ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```

Mình sẽ có riêng một bài viết về cách sử dụng Homebrew khi lập trình.

2. Cài Đặt Ruby :

Phiên bản Ruby được đề nghị cho thời gian hiện tại là : 2.4.0

```bash
brew install rbenv ruby-build

# Add rbenv to bash so that it loads every time you open a terminal
echo 'if which rbenv &gt; /dev/null; then eval "$(rbenv init -)"; fi' &gt;&gt; ~/.bash_profile
source ~/.bash_profile

# Install Ruby
rbenv install 2.4.0
rbenv global 2.4.0
ruby -v
```

3. Cài  Đặt Rails :

Phiên bản Rails điệc đề nghĩ cho thời gian hiện tại là : 4.2.4

Mở terminal lên và gõ : 

```bash
gem install rails -v 4.2.4
```

Rails đã được cài đặt nhưng để sử dụng các bạn cần gõ tiếp :

```
rbenv rehash
```

Để kiểm tra phiên bản rails hiện tại bạn gõ:

```bash
rails -v
# Rails 4.2.4
```

4. Tạo Ứng Dụng Rails Đầu Tiên:

Các bạn mở terminal lên và gõ : 

```bash
rails new myapp
cd myapp
rails s
```

Mở trình duyệt lên và gõ vào đường dẫn : http://localhost:3000/ để truy cập vào ứng dụng của bạn.

# Cài Đặt Ruby on Rails trên Ubuntu 16.04 Xenial Xerus 

1. Cài Đặt Ruby : 

```bash
sudo apt-get update
sudo apt-get install git-core curl zlib1g-dev build-essential libssl-dev libreadline-dev libyaml-dev libsqlite3-dev sqlite3 libxml2-dev libxslt1-dev libcurl4-openssl-dev python-software-properties libffi-dev
```

```bash
cd
git clone git://github.com/sstephenson/rbenv.git .rbenv
echo 'export PATH="$HOME/.rbenv/bin:$PATH"' &gt;&gt; ~/.bashrc
echo 'eval "$(rbenv init -)"' &gt;&gt; ~/.bashrc
exec $SHELL

git clone git://github.com/sstephenson/ruby-build.git ~/.rbenv/plugins/ruby-build
echo 'export PATH="$HOME/.rbenv/plugins/ruby-build/bin:$PATH"' &gt;&gt; ~/.bashrc
exec $SHELL

git clone https://github.com/sstephenson/rbenv-gem-rehash.git ~/.rbenv/plugins/rbenv-gem-rehash

rbenv install 2.2.3
rbenv global 2.2.3
ruby -v
```


```bash
echo "gem: --no-ri --no-rdoc" &gt; ~/.gemrc
gem install bundler
```

2. Cài Đặt Rails

– Cài đặt NodeJs

```bash
curl -sL https://deb.nodesource.com/setup_4.x | sudo -E bash -
sudo apt-get install -y nodejs
```

– Cài đặt Rails:

```
gem install rails -v 4.2.4
```

```
rbenv rehash
```

```
rails -v
```

4. Tạo Ứng Dụng Rails Đầu Tiên:

Các bạn mở terminal lên và gõ : 

```bash
rails new myapp
cd myapp
rails s
```

Mở trình duyệt lên và gõ vào đường dẫn : http://localhost:3000/ để truy cập vào ứng dụng của bạn.

Trong bài tiếp theo mình sẽ hưỡng dẫn các bạn sử dụng Database cho ứng dụng Ruby on Rails.

>Happy Coding!