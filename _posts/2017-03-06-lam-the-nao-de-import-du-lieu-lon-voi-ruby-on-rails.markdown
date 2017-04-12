---
layout: post
title: Làm  Thế Nào Để Import Dữ Liệu Lớn Với Ruby on Rails
date: 2017-03-05 18:00
comments: true
external-url: 
categories: ReactNative
keywords: ruby, rails, ruby on rails, import data, import data with rails, import du lieu voi rails
excerpt: Sử dụng ActiveRecord để insert thêm dữ liệu vào database là việc làm hết sức thường xuyên và quá quen thuộc với mọi người. Để thực hiện import dữ liệu vào database, thông thường ta sẽ chạy một vòng lặp để duyệt qua tất cả các dòng dữ liệu trong file, tạo ra đống các ActiveRecord object để thực hiện insert vào database (DB)
image:
  feature: ../../assets/rails-best-coder-info.png
  credit: 
  creditlink:
---
Sử dụng ActiveRecord để insert thêm dữ liệu vào database là việc làm hết sức thường xuyên và quá quen thuộc với mọi người. Để thực hiện import dữ liệu vào database, thông thường ta sẽ chạy một vòng lặp để duyệt qua tất cả các dòng dữ liệu trong file, tạo ra đống các ActiveRecord object để thực hiện insert vào database (DB). Đại loại sẽ kiểu như thế này :

```ruby
data_file = ".../test.csv"
CSV.foreach(data_file) do |attrs|
    Model.create attr1: attrs[0], ....
end
```

Ừ thì đoạn code trên hết sức bình thường, và có thể được sử dụng thường xuyên, nhưng đó là với lượng dữ liệu ít, tuy nhiên với dữ liệu lớn, con số lên tới từ mấy chục nghìn thì sao ? Làm như thế này sẽ tốn rất nhiều thời gian xử lý cũng như hao tốn tài nguyên hệ thống. Việc hao tốn thế nào bạn sẽ thấy ngay sau mỗi lần tối ưu lại code.

## Bối cảnh

Ta sẽ có một file csv chứa tọa độ địa điểm các cửa hàng ở Hà Nội gồm 1000 dòng, tương ứng với 1000 cửa hàng cần thêm vào database. Mình sử dụng MySQL và cài đặt ngay trên máy local, sử dụng HDD có tốc độ quay 5400 rpm.

### 1. Sử dụng cách "nông dân"

Nông dân chính là đoạn code đầu tiên ở bên trên, với bối cảnh ở trên, thời gian mình chạy được mất khoảng `48.864907147s`, chưa biết nhanh chậm thế nào nhỉ ? Vụ này để lát nữa tính tiếp, giờ nhìn vào log khi chạy, ta sẽ thấy dạng :

```
....
(0.1ms)  SAVEPOINT active_record_1
SQL (0.2ms)  INSERT INTO ...
(0.1ms)  RELEASE SAVEPOINT active_record_1
...
(0.1ms)  SAVEPOINT active_record_1
SQL (0.2ms)  INSERT INTO ...
(0.1ms)  RELEASE SAVEPOINT active_record_1
```

Rất rất nhiều cặp log như bên trên. Trông bối cảnh trên ta sẽ có 1000 cặp log như vậy. Nếu bạn thắc mắc ý nghĩa của đoạn log trên thì có thể đọc thêm về `Active Record Transactions`.
Phân tích một chút về cặp log gồm 3 dòng bên trên. Mỗi dòng sẽ tương ứng với 1 lần kết nối vào database để xử lý, vậy ta sẽ mất 3 lần kết nối vào db cho việc insert thành công một cửa hàng vào db, tổng cộng ta sẽ mất 3000 lần kết nối tới db cho việc import 1000 cửa hàng. Đây là kết quả của việc mỗi lần insert một cửa hàng ta lại tạo 1 db transaction.

### 2. Gói gọn trong 1 transaction

Thay vì mỗi lần insert một cửa hàng ta tạo 1 db transaction, ta sẽ chỉ sử dụng một transaction duy nhất để commit cả 1000 cửa hàng. Rất đơn giản, chỉ việc bao đoạn code của chúng ta với `ActiveRecord::Base.transaction`

```ruby
data_file = ".../test.csv"
ActiveRecord::Base.transaction do
    CSV.foreach(data_file) do |attrs|
        Model.create attr1: attrs[0], ....
    end
end
```

Và giờ, ta sẽ mất `1.510126093s` để import 1000 cửa hàng. Rồi, nhìn coi, từ `48.864907147s` giản xuống còn `1.510126093s`. Không chỉ giảm thời gian chạy xuống, mà còn giảm gánh nặng cho db rất rất nhiều. Nhìn vào log ta sẽ thấy dạng :

```
(0.1ms)  SAVEPOINT active_record_1
SQL (0.2ms)  INSERT INTO...
SQL (0.2ms)  INSERT INTO...
...
SQL (0.2ms)  INSERT INTO...
(0.1ms)  RELEASE SAVEPOINT active_record_1
```

Sẽ có 1000 dòng `INSERT INTO...`, vậy nên ta sẽ tốn 1002 lần kết nối tới db thay vì 3000 lần như trước. Vẫn có vẻ nhiều kết nối quá, tiếp tục giảm nó xuống thôi.

### 3. Insert hàng loạt

Như ta đã biết, MySQL nói riêng và các hệ quản trị dữ liệu nói chung đều hỗ trợ việc insert dữ liệu hàng loạt với một lần insert duy nhất.

```
INSERT INTO shops (name, lat, long) VALUES ('A', '123', '321'), ('B', '113', '121'), ('C', '122', '324') ...
```

Vậy nên tội gì chúng ta không tận dụng để giảm số lần kết nối tới db đi. Sử dụng đặc tính này, ta sẽ chỉ cần tới 3 kết nối tới db thôi.
Để làm việc này bạn có thể viết câu SQL thuần, rồi execute nó, hoặc có thể sử dụng gem để nhìn nó sạch sẽ sáng sủa hơn. Ví dụ ở đây mình dùng gem `activerecord-import`. Đoạn code mới của ta sẽ có dạng :

```ruby
data_file = ".../test.csv"
shops = CSV.read datafile
attributes = [:name, :lat, :long]
Shop.import attributes, shops
```

Thời gian import 1000 cửa hàng bây giờ sẽ chỉ còn `0.247711571s`. Hay chưa ? Sau 2 lần tối ưu, ta đã giảm được thời gian từ `48.864907147s` xuống còn `0.247711571s`.

### 4. Chia nhỏ file dữ liệu

Đến đây ta đã rút ngắn được thời gian một cách rõ ràng, tuy nhiên với lượng dữ liệu lớn hơn nữa, hàng trăm triệu thì sao ? Rõ ràng tới bước thứ 3, ta vẫn có thể import một cách trơn tru với thời gian ngắn, nhưng vẫn còn cách nữa để nó còn nhanh và mượt hơn nữa. Cách này mình chưa kiểm chứng vì chưa cần thiết cũng như chưa có bối cảnh cụ thể, vậy nên sẽ chỉ có lý thuyết, bạn có thể tự thực hành :)
Garbage collector là một vấn đề cần quan tâm (Google để hiểu tìm hiểu về khái niệm này nhé). Với file chứa hàng trăm triệu dòng dữ liệu thì file này khá là lớn và việc quản lý bộ nhớ ta cần quan tâm đến. Ý tưởng ở đây là chia nhỏ file ra và import từng phần nhỏ và chạy, sau mỗi lần import phần nhỏ, memory sẽ được giải phóng. Việc này có thể không rút ngắn hơn thời gian import, nhưng chắc chắn rằng nó sẽ giúp cho hệ thống chạy mượt mà và trơn tru hơn.

# Kết luận

Sau vài bước, ta đã rút ngắn thời gian import 1000 cửa hàng vào db từ `48.864907147s` xuống còn `0.247711571s`. Các bước trên hoàn toàn là kiến thức cơ bản, tuy nhiên lại ít được các new dev để ý và hiểu rõ vấn đề. Bài viết này mình đã chỉ ra các vấn đề, cách thức hoạt động bên trong và hướng xử lý vấn đề, hi vọng sẽ giúp ích cho các bạn.
Do kiến thức còn hạn hẹp nên bài viết tất sẽ không tránh khỏi những sai sót, rất mong nhận được nhận xét, góp ý từ các bạn để mình hoàn thiện lại bài viết cũng như củng cố lại kiến thức.

CẢM ƠN TÁC GIẢ ĐÃ CHO CHÚNG TA MỘT BÀI VIẾT HỮU ÍCH

Nguồn : [https://viblo.asia/nguyentrunghieu/posts/PdbknoxJvyA](https://viblo.asia/nguyentrunghieu/posts/PdbknoxJvyA)