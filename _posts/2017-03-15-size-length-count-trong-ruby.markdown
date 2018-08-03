---
layout: post
author: derek
image: assets/images/feature_default.jpg
featured: false
hidden: false
layout: post
author: derek
image: assets/images/feature_default.jpg
featured: false
hidden: false
title: Size, length và count trong Ruby
date: 2017-03-15 12:00
comments: true
external-url:
categories: Ruby
keywords: ruby, array ruby, size, length, count
excerpt: Tại sao lại có count, size, length ? count, size, length đều dùng để tính số lượng. Vậy tại sao có đến tận 3 method?.
---
>Tại sao lại có count, size, length ?

**count**, **size**, **length** đều dùng để tính số lượng. Vậy tại sao có đến tận 3 method?.

Chúng ta hãy cùng nhau tìm hiểu qua bài viết này. Trước hết ta có bối cảnh như sau:

- Bảng users lưu trữ các thông tin về các users
- Bảng comments lưu các thông tin về các comments của mỗi User.
- Một user có nhiều comments.

```ruby
class User
  has_many :comments
end

class Comment
  belongs_to :user, counter_cache: true
end
```

## 1. Trong trường hợp không dùng counter cache:

Counter cache là gì? Các bạn có thể tham khảo bài viết [Counter Cache](https://viblo.asia/nguyen.thanh.luan/posts/zb7vD8ALvjKd)

Chú ý đừng nhầm lẫn với từ `cache` (lưu vào bộ nhớ tạm) trong bài viết này sử dụng.

### a. count

Count có vẻ như được biết tới nhiều hơn. Hãy cùng xem count hoạt động như thế nào

```bash
User.count

   (0.3ms)  SELECT COUNT(*) FROM "users"

 => 52

 users = User.all

 users.count

   (0.3ms)  SELECT COUNT(*) FROM "users"

 => 52
```

=> Không có gì bất ngờ, count `đã vào database` để đếm số lượng bản ghi của scope hiện tại.

### b. length

Hãy xem length thực hiện như thế nào...

```ruby
users = User.all

 users.length

 User Load (1.3ms)  SELECT "users".* FROM "users"

=> 52

users.length

=> 52

User.length

undefined method 'length' for #
```

=> Như vậy rails đã vào database để load tất cả các users về và đo kích thước của mảng kết quả. Thông thường cách này sẽ lâu hơn việt sử dụng `User.all.count`

Và cũng bởi ý nghĩa như vậy nên không có phương thức `length` cho class `User`.

Sử dụng **length** có vẻ là lựa chọn tồi hơn khi muốn đếm số phần tử mà lại đi load database về và đếm nhưng hãy để ý.

Khi gọi `users.length` lần thứ 2 chúng ta có kết quả 52 vì length sẽ đếm số phần tử hiện tại của mảng.
=> vậy là nó cũng có ưu điểm nhất định.

### c. size

Vậy còn size?

```ruby
User.all.size

   (0.3ms)  SELECT COUNT(*) FROM "users"

=> 52

 users = User.all

 users.size

=> 52

 users.size

=> 52

 User.size

undefined method 'size' for #
```

=> Vậy là size sẽ vào database gọi đến `SQL SELECT COUNT(*)` khi chưa có cache (tức là khi chưa có các bản ghi users trên bộ nhớ tạm), khi đã có rồi thì đếm luôn số bản ghi trên bộ nhớ tạm.

Vậy là size là giải pháp thông minh nhất và nên dùng nhất.

Tuy nhiên, không thể hoàn toàn thay thế count bằng size được.

Vì kết quả của length mang ý nghĩa là số phần tử trên mảng, của count mang ý nghĩa là số phần tử trong database.

Size sẽ chọn cách nào tốt nhất để lấy ra kết quả nhanh nhất.

## 2. Trong trường hợp có sử dụng counter cache:

### a. count

```ruby
User.first.comments.count
(0.2ms)  SELECT COUNT(*) FROM "comments" WHERE "comments"."user_id" = ?  [["user_id", 1]]
```

=> Không có gì bất ngờ, count vẫn sẽ sử dụng query database để lấy số comments của user 1

### b. length

```ruby
User.first.comments.length
SELECT "comments".* FROM "comments" WHERE "comments"."user_id" = ? [["user_id", 1]]
```

=> vẫn ngu như ngày nào, length vẫn query vào database để load các comments về và đếm độ dài mảng.
Và sẽ đếm trực tiếp nếu như đã có cache.

### c. size

```ruby
User.first.comments.size
SELECT "users".* FROM "users" ORDER BY "users"."id" ASC LIMIT 1
```

=> lạ nhỉ, tưởng sẽ phải là `SELECT COUNT(*)` chứ?, mà sao lại SELECT từ users.

=> à, cùng nhớ lại, chúng ta đang sử dụng, counter cache, số comments của user đã được lưu ở cột
comments_count nên không cần phải động chạm gì đến bảng comments nữa.

=> thông minh lần 2! size sẽ tự chọn giữa COUNT(*), đếm trên cache, hoặc đọc trong counter_cache.

Vì vậy trong hầu hết các trường hợp nên dùng size.

## 3. Kết luận

Vì vậy trong hầu hết các trường hợp nên dùng size !.

Hi vọng các bạn thấy bài viết này hữu ích.

Tham khảo: [https://www.reddit.com/r/rails/comments/3oh61p/counter_cache_sizecountlength/](https://www.reddit.com/r/rails/comments/3oh61p/counter_cache_sizecountlength/)