---
layout: post
author: derek
image: assets/images/feature_default.jpg
featured: false
hidden: false
title: Tìm hiểu về sự khác nhau giữa eql? và equal? trong Ruby
date: '2017-09-04 16:00:00'
comments: true
categories: ruby
keywords: ruby, eql?, equal?, generic equality, object
excerpt: Một thực tế là những người mới chân ướt chân ráo bước vào lập trình với Ruby thường nhầm lẫn các phương thức liên quan đến bình đẳng `==`, `eql?` và `equal?`.
---
Một thực tế là những người mới chân ướt chân ráo bước vào lập trình với Ruby thường nhầm lẫn các phương thức liên quan đến bình đẳng `==`, `eql?` và `equal?`.

Nếu các bạn muốn hiểu thêm về `==` và `===` thì có thể xem [Bạn biết gì về toán tử "===" triple equals (case equality operator) trong Ruby?](https://bestcoder.info/ruby/2017/08/26/ban-biet-gi-ve-toan-tu-===-triple-equals-trong-ruby.html).

Trong số các phương thức trên thì `equal?` được mô tả dễ hiểu nhất về bình đẳng bởi cái tên của nó.

`equal?` sẽ trả về kết quả là `true` chỉ khi nó (đối tượng bị gọi bởi `equal?`) và tham số (đối tượng được gọi bởi `equal?`) là cùng một đối tượng (giống nhau hoàn toàn).

Ví dụ:

```ruby
some_word = "word"
some_other_word = some_word

some_word.equal? some_other_word # true
```

`eql?` sẽ trả về kết quả là `true` khi cả hai đối tượng bị gọi bởi `eql?` và đối tượng được gọi bởi `eql?` có gía trị giống nhau.

Ví dụ:

```ruby
1 == 1.0 #=> true
1.eql? 1.0 #=> false
"test".eql? "test" #=> true
```

Tổng kết lại ta có:

```ruby
a = 'a' # => 'a'
other = a.dup # => 'a'

a == other # => true
a === other # => true
a.eql? other # => true
a.equal? other # => false
a.equal? a # => true
```

Tham khảo:

- [What is the difference between '===', '==', '.equal?' and '.eql?' in Ruby?](http://www.rian.me/2013/10/15/what-is-the-difference-between-equals-equals-equals-and-equals-equals-in-ruby/)
