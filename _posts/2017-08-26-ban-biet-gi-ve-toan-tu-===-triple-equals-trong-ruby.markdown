---
title: Bạn biết gì về toán tử "===" triple equals (case equality operator) trong Ruby?
date: '2017-08-26 09:00:00'
comments: true
categories: ruby
keywords: ruby, ===, toan tu, operator, operator in ruby, the === operator
excerpt: Toán tử triple equals (===) trong Ruby thực sự là một cái gì đó rất thú vị. Nó có ở khắp mọi nơi trong Ruby nhưng hầu hết mọi người chưa bao giờ thấy nó thực ra ở đó. Nhưng, làm thế nào đến nó ở khắp mọi nơi và không ai đã từng nhìn thấy nó?
---

Là một lập trình viên và thường xuyên nhảy việc thì chắc hẳn các bạn cũng đã từng gặp câu hỏi về `===` trong Javascript khi đi phỏng vấn rồi phải không? Nếu chưa thì mình xin nhắc lại 😅.

Trong Javascript, toán tử `===` chỉ trả về **true** nếu như cả hai toán hạng đều cùng một loại và có cùng giá trị. Nếu so sánh khác loại, kết quả sẽ trả về **false**. Cái này sẽ có giải thích và ví dụ cụ thể ở [đây](https://codeaholicguy.com/2016/06/14/nen-dung-hay-de-so-sanh-trong-javascript/). Xin chân thành cảm ơn bác @codeaholicguy , chả hiểu sao cứ search lại ra bài bác đầu tiên 😂 . 

## Vậy còn trong Ruby thì sao?

Toán tử `===` trong Ruby thường được gọi là case equality operator khác với toán tử `==` hay còn gọi là generic equality. 

Thằng `==` so sánh có cùng giá trị hay không. Đây là cách so sánh phổ biến và cơ bản nhất trong hầu hết các ngôn ngữ lập trình. 

Toán tử `===` thực sự là một cái gì đó rất thú vị. Nó có ở khắp mọi nơi trong Ruby nhưng hầu hết mọi người chưa bao giờ thấy nó thực ra ở đó. Nhưng, làm thế nào đến nó ở khắp mọi nơi và không ai đã từng nhìn thấy nó? Nó ẩn bên trong một cấu trúc điều khiển thông thường, "case / when". Bất cứ khi nào bạn đang sử dụng "case / when", trên thực tế bạn đang sử dụng toán tử "===" và điều này làm cho câu lệnh case trên Ruby mạnh hơn nhiều so với các ngôn ngữ như C hay Java

Toán tử `===` được hiểu đơn giản là so sánh theo kiểu trường hợp.  Các điều kiện của case sẽ đc implement với mỗi class tương ứng như:

- Range
- Regex
- Proc
...

Ví dụ đơn giản:

```ruby
(1...10) === 5 # => true
```

Toán tử `===` sẽ kiểm tra array trong mệnh đề trên có thằng 5 hay không. Nếu có sẽ trả về true nếu không trả về false.

```ruby
(1..5) === 3           # => true
(1..5) === 6           # => false
Integer === 42          # => true
Integer === 'fourtytwo' # => false
/ell/ === 'Hello'     # => true
/ell/ === 'Foobar'    # => false
"a" === "b" # false # different values, different objects
"a" === "a" # true # same values
```

Tiếp theo ta có:

```ruby
"test" == "test"  #=> true
"test" === "test" #=> true
```

Thế `==` và `===` có gì khác nhau?

```ruby
String === "test"   #=> true
String == "test"    #=> false
```

Như vậy `===` cũng đơn thuần là so sánh giá trị chứ không phải là so sánh object có điều nó dùng case để so sánh và `===` còn được gọi là Case equality.

## Sử dụng

### Array.grep

Mảng có một phương pháp được gọi là `grep` sử dụng `===`.

```ruby
# grep(pattern) → array

(1..100).grep(38..44)
#=> [38, 39, 40, 41, 42, 43, 44]

names = %w(
  William
  Kate
  Adam
  Alexa
  James
  Natasha
)
names.grep(/am/)
# => %w(William Adam James)
```

Kết quả ta có được là một mảng với các phần tử thỏa `===` với pattern của grep.

### Ranges

`===` kiểm tra để xem đối tượng đó là một trong các phần tử của range đó hay không.

```ruby
(2..4) == 3 # => false
(2..4) === 3 # => true
(2..4) === 6 # => false

(Date.new(2017, 8, 21)..Date.new(2017, 8, 27)) === Date.new(2017, 8, 27)
# => true

(Date.new(2017, 8, 21)..Date.new(2017, 8, 27)) === Date.new(2017, 8, 29)
# => false

("a".."z") === "a"
# => true

("a".."z") === "abc"
# => false
```

### Class / Module

```ruby
mod === obj #→ true or false
```

`===` trả về true nếu obj là một instance của mod hoặc một trong những hậu duệ của mod. Việc sử dụng hạn chế cho module, nhưng có thể được sử dụng để phân loại các đối tượng theo class. Về cơ bản thực hiện như:

```ruby
obj.kind_of?(mod)
```

Ví dụ:

```ruby
"text".class.ancestors
# => [String, Comparable, Object, Kernel, BasicObject]

String === "text"
# => true

Object === "text"
# => true

Comparable === "text"
# => true

Numeric === "text"
# => false
```

### Regexp

```ruby
rxp === str #→ true or false
```

Về cơ bản thực hiện như:

```ruby
rxp =~ str >= 0
```

Ví dụ:

```ruby
/^[a-z]*$/ === "HELLO"
#=> false

/^[A-Z]*$/ === "HELLO"
#=> true
```

### Proc

```ruby
proc === obj # → result_of_proc
```

Gọi block với một đối tượng như tham số giống như `#call`.

Ví dụ:

```ruby
is_today = -> (date) { Date.current === date }

is_today === Date.current
# => true

is_today === Date.tomorrow
# => false

is_today === Date.yesterday
# => false
```


### Lambdas

Tương tự như `Proc`:

```ruby
is_even = -> (x) { x % 2 == 0 }
is_even == 4 # => false
is_even === 4 # => true
is_even === 5 # => false
```

### Object

Đối với hầu hết các Object, `===` tương tự như `==`.

Mọi ý kiến đóng góp và thảo luận xin bình luận bên dưới. Xin cảm ơn 🙇🙇🙇

---

Bài viết được tham khảo từ:

- [Chuyện xưa cũ: Nên dùng === hay == để so sánh trong Javascript?](https://codeaholicguy.com/2016/06/14/nen-dung-hay-de-so-sanh-trong-javascript/)

- [The === (case equality) operator in Ruby](http://blog.arkency.com/the-equals-equals-equals-case-equality-operator-in-ruby/)

- [Ruby Basics - Equality operators in Ruby ](https://mauricio.github.io/2011/05/30/ruby-basics-equality-operators-ruby.html)

- [The rarely used === in Ruby](https://coderwall.com/p/53xawg/the-rarely-used-in-ruby)
