---
layout: post
title: Tìm Hiểu Về Locking Active Record Trong Ruby On Rails
date: 2017-02-19 13:00
comments: true
external-url:
categories: Ruby-on-Rails
---
>Như các bạn đã biết, **Data consistency** rất quan trọng trong nhiều ứng dụng, đặc biệt là cho các ứng dụng liên quan đến tài chính, ngân hàng, ... Một lỗi nhỏ có thể trở thành một thảm kịch nếu chúng ta không quan tâm đến nó một cách nghiêm túc. Lần này, tôi sẽ nói một chút về Locking và làm thế nào bạn có thể sử dụng nó hiệu quả.

## Tại Sao Locking Lại Cần Thiết Đến Thế ?

Hãy tưởng tượng bạn đang xây dựng một ứng dụng, trong đó mỗi người sẽ có một tài khoản với với một tiền ảo. Và người dùng có id là 5 đang truy cập vào trang web để mua một số món hàng, chúng ta nhận vào tài khoản như thế này :

```
account = Account.find_by_user_id(5)
```

Sau khi chọn được món hàng yêu thích của mình với giá $ 50, nhấp chuột kiểm tra và bắt đầu trả tiền cho món hàng đó. Trước khi thực hiện yêu cầu, đầu tiên chúng ta sẽ kiểm tra xem anh ta có đủ số tiền trong tài khoản của mình, và nếu anh ta thoả mãn điều kiện, chúng ta sau đó sẽ giảm số dư trong tài khoản của anh ấy một số tiền tương ứng với giá của mặt hàng đó.

```
if account.balance >= item.price
    account.balance = account.balance - item.price
    #some other long processes here
    account.save
end
```

Điều đó có vẻ dễ dàng phải không? Tuy nhiên, nếu những gì anh chàng này sẽ mở ra một tab của trang web, chọn một món hàng khác với giá $ 80 và bằng cách nào đó đồng thời nhấp chuột kiểm trên cả các tab. Mặc dù nó là rất hiếm, có thể có một cơ hội khi các yêu cầu trên tab đầu tiên và thứ hai đến máy chủ gần như cùng một thời điểm, và họ đều được xử lý bởi máy chủ đồng thời. Đây là cách mà request ở tab thứ nhất đã được thực hiện :

```
#account.balance = 100
account = Account.find_by_user_id(5) 


#item.price is 50
if account.balance >= item.price
  #it's good, allow user to buy this item
  account.balance = account.balance - item.price

  #account.balance is now 50

  account.save
end
```

Nhưng sau khi thực hiện `account.balance = account.balance - item.price` và trước khi lưu vào tài khoản, CPU thực hiện các yêu cầu thứ hai (với cùng code) :

```
account = Account.find_by_user_id(5) 
#account.balance is still 100

#item.price is 80
if account.balance >= item.price
  #it's good, allow user to buy this item
  account.balance = account.balance - item.price

  #account.balance is now 20

  account.save
end
```

Tôi chắc rằng bạn có thể thấy vấn đề bây giờ. Mặc dù sau khi mua món hàng đầu tiên, chúng ta sẽ nghĩ rằng người sử dụng chỉ có `$50` trong tài khoản của anh ta, và theo lý thuyết anh ta không thể mua một món hàng khác với giá cao hơn `$50`. Nhưng ở đây, anh ta có thể mua cả hai món hàng vì nó vượt qua các kiểm tra điều kiện.

Bằng cách sử dụng Locking, chúng ta có thể ngăn chặn tình trạng tương tự. Khi Locking được đặt ra, họ sẽ không cho phép hai tiến trình đồng thời cập nhật các đối tượng trong cùng một thời điểm.

Nói chung, có hai loại Locking : **Optimistic** và **Pessimistic**. Từ từ, tôi nghĩ rằng bạn cũng có thể phần nào đoán được ý nghĩa thật sự của chúng.

## Optimistic Locking

Trong loại này, nhiều người dùng có thể truy cập cùng một đối tượng để đọc giá trị của nó, nhưng nếu hai người dùng thực hiện cập nhật thì sẽ phát sinh mâu thuẫn, chỉ có một người sử dụng sẽ thành công và một trong những người khác sẽ không được thực hiện.

```
p1 = Person.find(1)
p2 = Person.find(1)

p1.first_name = "Michael"
p1.save

p2.first_name = "should fail"
p2.save # Raises a ActiveRecord::StaleObjectError
```

Để tạo **Optimistic locking**, bạn có thể tạo ra một field lock_version mà bạn muốn đặt khóa và Rails sẽ tự động kiểm tra trước khi cập nhật các đối tượng. Cơ chế của nó là khá đơn giản. Mỗi lần các đối tượng được cập nhật, giá trị lock_version sẽ được tăng lên. Do đó, nếu hai yêu cầu muốn thực hiện cùng một đối tượng, yêu cầu đầu tiên sẽ thành công vì lock_version của nó cũng giống như khi nó được đọc. Nhưng yêu cầu thứ hai sẽ thất bại vì lock_version đã được tăng lên trong cơ sở dữ liệu của các yêu cầu đầu tiên.

Trong loại locking này, bạn có trách nhiệm xử lý các ngoại lệ trả lại khi cập nhật hành động không thành. Bạn có thể đọc thêm ở đây:

[http://api.rubyonrails.org/classes/ActiveRecord/Locking/Optimistic.html](http://api.rubyonrails.org/classes/ActiveRecord/Locking/Optimistic.html)

## Pessimistic Locking

Với loại locking này, chỉ có người dùng đầu tiên truy cập đến các đối tượng sẽ có thể cập nhật nó. Tất cả những người dùng khác sẽ bị loại khỏi thậm chí đọc các đối tượng (hãy nhớ rằng trong **Optimistic locking**, chúng tôi chỉ khóa nó khi cập nhật dữ liệu và người dùng vẫn có thể đọc nó).

Rails sẽ thực hiện **Pessimistic Locking** bằng cách phát hành truy vấn đặc biệt trong cơ sở dữ liệu. Ví dụ, giả sử bạn muốn lấy đối tượng tài khoản và khóa nó cho đến khi bạn hoàn thành việc cập nhật:

```
account = Account.find_by_user_id(5)
account.lock!
#no other users can read this account, they have to wait until the lock is released
account.save! 
#lock is released, other users can read this account
```

Các bạn có thể tham khảo thêm :
[http://api.rubyonrails.org/classes/ActiveRecord/Locking/Pessimistic.html](http://api.rubyonrails.org/classes/ActiveRecord/Locking/Pessimistic.html)

## Kết Luận

**Locking** nên được sử dụng phụ thuộc vào yêu cầu. Nếu không có bất kỳ yêu cầu đặc biệt, **Optimistic locking** là đủ bởi vì nó là linh hoạt hơn và nhiều hơn nữa yêu cầu đồng thời có thể được phục vụ. Trong trường hợp của **Pessimistic Locking**, bạn cần phải chắc chắn rằng bạn mở khóa khi bạn hoàn thành việc cập nhật các đối tượng.

> ## Happy coding!