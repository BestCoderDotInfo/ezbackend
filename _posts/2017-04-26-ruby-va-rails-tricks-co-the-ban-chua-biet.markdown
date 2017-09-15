---
title: Ruby & Rails Tricks - Có Thể Bạn Chưa Biết ?
date: 2017-04-26 21:00
comments: true
external-url: 
categories: Tricks
keywords: ruby, php, rails
excerpt: Có thể các tricks này có người biết rồi hoặc sắp được biết nhưng mình cũng mạn phép được chia sẻ lại. Văn mình có thể không hay nhưng Ruby thì luôn đẹp 😆
---
>Có thể các tricks này có người biết rồi hoặc sắp được biết nhưng mình cũng mạn phép được chia sẻ lại. Văn mình có thể không hay nhưng Ruby thì luôn đẹp 😆.

## Ruby

### 1. Có nhiều cách để call một lamda

```ruby
my_lambda = -> { puts "welcome to summoner's rift"}
my_lambda.call
my_lambda[]
my_lambda.()
my_lambda.===
```

### 2. Tạo một hash từ một list value

```ruby
Hash['key1', 'value1', 'key2', 'value2']

# => {"key1"=>"value1", "key2"=>"value2"}
```
### 3 Tạo một pre-filled array

```ruby
Array.new(10) { rand 300 }
```
Kết quả ta sẽ có được một array với 10 số ngẫu nhiên.

### 4. Các object đặc biệt

Như các bạn đã biết trong ruby mọi thứ đề là object nên sẽ có `object_id` . Tuy nhiên một số đối tượng cá biệt có id cố định :

```ruby
false.object_id # 0
true.object_id  # 2
nil.object_id   # 4
 
1.object_id # 3
2.object_id # 5
```
Các id của `Fixnum` sẽ tính theo công thức : `(number * 2) + 1`.

Maximum Fixnum là `1073741823`, sau đó bạn sẽ có một Bignum object.

### 5. Convert bất kì giá trị nào sang kiểu boolean

```ruby
!!(1)   # true
!!(nil) # false
```

### 6.  Kiểm tra tên của các method đang dùng

```ruby
def foo
  bar
end
 
def bar
  puts caller
end
 
foo
```

Kết quả sẽ là:

```shell
-:3:in 'foo'
-:10:in '<main>'
```
Nếu bạn muốn biết tên của method hiện tại thì dùng `__method__` hoặc `  __callee__`

### Metaprogramming:

```ruby
['admin', 'marketer', 'sales'].each do |user_role|
    define_method "#{user_role}?" do
        role == user_role
    end
end
```

## Rails

### 7.  Xoá các space không cần thiết

```ruby
"    My    \r\n  \t   \n   books       ".squish # => "My books"
```

### 8. Dùng pluck thay vì map

`pluck` là method để lấy 1 column cho trước trong các record, mà không load toàn bộ các record đó. Vì thế mà tốc độ xử lý và RAM cũng hiệu quả hơn.

```ruby
def admin_user_ids
  User.where(admin: true).map(&:id)
end
```
thành 

```ruby
def admin_user_ids
  User.where(admin: true).pluck(:id)
end
```

### 9.  timezone trong Rails

Trong Rails, có 2 cách để setting timezone, cách 1 là setting trong application.rb, cách 2 là sử dụng timezone dựa theo biến số môi trường TZ. Nếu trong trường hợp setting giữa 2 cách này mâu thuẫn với nhau, sẽ nảy sinh ra những lỗi không thể dự đoán trước. Vì thế, tốt hơn là thống nhất chỉ sử dụng timezone trong application.rb.

Ví dụ, không dùng `Date.today` mà dùng `Date.current`, không dùng `Time.now` mà dùng `Time.current `( hoặc `Time.zone.now `)

### 10. Các method thay đổi string thành số nhiều, số ít, …

```ruby
"my_book".camelize # => "MyBook"

"MyBook".underscore # => "my_book"

"my_book".dasherize # => "my-book"

"book".pluralize            # => "books"
"person".pluralize          # => "people"
"fish".pluralize            # => "fish"
"book_and_person".pluralize # => "book_and_people"
"book and person".pluralize # => "book and people"
"BookAndPerson".pluralize   # => "BookAndPeople"

"books".singularize            # => "book"
"people".singularize           # => "person"
"books_and_people".singularize # => "books_and_person"
"books and people".singularize # => "books and person"
"BooksAndPeople".singularize   # => "BooksAndPerson"

"my_books".humanize # => "My books"

"my_books".titleize # => "My Books"

"my_book".classify  # => "MyBook"
"my_books".classify # => "MyBook"

"my_book".tableize # => "my_books"
"MyBook".tableize  # => "my_books"
```

### 11. Sử dụng Object.try(:method_name) thay vì kiểm tra nil

```ruby
if parent.children && parent.children.singleton?
  singleton = parent.children.first
  send_mail_to(singleton)
end
```

ta có thể viết

```ruby
# nếu children là nil thì try(:singleton?) cũng trả về nil 
# nếu children không nil thì children.singleton? được gọi như bình thường
if parent.children.try(:singleton?)
  singleton = parent.children.first
  send_mail_to(singleton)
end
```

### 12. Sử dụng presence

```ruby
if user.name.blank?
  name = "What's your name?"
else
  name = user.name
end
```
thành 

```ruby
name = user.name.presence || "What's your name?"
```

`"".presence` hoặc `[].presence` sẽ trả về `nil`.

```ruby
name = ""
puts name.presence || "What's your name?" # => What's your name?
```

Nguồn: google và sẽ update thêm.

>Happy Coding !