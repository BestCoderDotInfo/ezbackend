---
layout: post
title:  Tất tần tật cài đặt GitHub Pages với Cloudflare
date: 2017-05-05 19:00
comments: true
external-url:
categories: website
keywords: github pages, github, git, website, cloudflare, https
excerpt: Nếu bạn muốn xây dựng một website mà chỉ bao gồm HTML, CSS và Javascript. Bạn có thể xây dựng 1 static web hoàn toàn miễn phí phục vụ cho việc học tập hoặc thậm chí cho mục đích kinh doanh với GitHub Pages(https://pages.github.com).
---
>Nếu bạn muốn xây dựng một website mà chỉ bao gồm HTML, CSS và Javascript. Bạn có thể xây dựng 1 static web hoàn toàn miễn phí phục vụ cho việc học tập hoặc thậm chí cho mục đích kinh doanh với GitHub Pages (https://pages.github.com).

### Bạn nên đọc những điều này:

- Bạn muốn tự cấu hình chuyển hướng hoặc máy chủ khác miễn phí đối với GitHub page của bạn.
- Bạn muốn URL trang web của bạn sẽ là https://trangwebcuaban.domain

## Tại sao lại là Github ?

- Dễ dàng cài đặt và bắt dầu với GitHub Pages.
- Dễ dàng thay đổi với các commit mới.

## Tại sao lại là Cloudflare ?

- Nó miễn phí
- Hỗ trợ SSL (HTTPS)
- Quản lý DNS cực kì đơn giản
- Hỗ trợ cache 
- Minify CSS, Javascript,...
- Điều chỉnh chuyển hướng, luôn là HTTPS với page rule.
- HTTP2/SPDY
- Cho phép cấu hình HSTS (HTTP Strict Transport Security)

## Trước khi bắt đầu chúng ta cần một số thứ :

- Github account
- Cloudflare account
- Một domain (Bạn có thể mua ở GoDady).

Đó là tất cả những gì chúng ta cần. Bạn đã sẵn sàng chưa ? 

## Bước 1: Tạo một Github repo

Ở đây mình dùng Jekyll theme, có rất nhiều blog template có sẵn để bạn chọn: [dbyll](https://github.com/dbtek/dbyll-ghost), [jekyll-bootstrap](https://github.com/plusjade/jekyll-bootstrap/) việc bạn cần chỉ là clone theme đó, tùy chỉnh các cài đặt ở file `_config.yml` cho site của bạn như tên, address,... hoặc bạn tự xây dựng riêng một blog template cho mình. Sau đó bạn push toàn bộ code đó lên repo vừa mới tạo. 

Tiếp theo, vào phần setting của repo và làm theo như hình:

![](/assets/github-1.png){:height="100%" width="100%"}

Chọn branch sẽ dùng làm content cho trang web của bạn.

![](/assets/github-2.png){:height="100%" width="100%"}

Truy cập https://<yourgithubusername>.github.io/repository và bạn sẽ thấy thành quả của mình. 

## Bước 2: Thêm domain cho trang web

![](/assets/github-3.png){:height="100%" width="100%"}

Save lại và bây giờ chúng ta đã có một trang web host github với domain của riêng mình.

Tiếp theo chúng ta sẽ cài đặt Cloudflare cho nó.

## Bước 3: Cài đặt domain trên Cloudflare

![](/assets/github-4.png){:height="100%" width="100%"}

Đăng nhập vào Cloudflare và Add Site. Sau đó Begin Scan.

## Bước 4: Cài đặt DNS Records cho domain của bạn

![](/assets/github-5.png){:height="100%" width="50%"}![](/assets/github-6.png){:height="100%" width="50%"}

Ở bước này chúng ta biết được GitHub Pages dùng 2 A Record DNS là: 
- 192.30.252.153
- 192.30.252.154

Khi  bạn cài đặt xong, tất cả request đến domain của bạn sẽ được chuyển đến website của bạn trên github mà bạn đã thực hiện ở bước 2.

![](/assets/github-7.png){:height="100%" width="100%"}

Continue

## Bước 5: Cập nhật Nameservers cho domain của bạn

Ở đây mình dùng GoDaddy :D 

![](/assets/github-8.png){:height="100%" width="100%"}

nên sẽ như thế này: 

![](/assets/github-9.png){:height="100%" width="100%"}

Chúng ta vào **Overview** và đã thấy có kết quả: 


![](/assets/github-10.png){:height="100%" width="50%"}![](/assets/github-11.png){:height="100%" width="50%"}

## Bước 6: Cài đặt Minify, Browser Cache Expiration, SSL

![](/assets/github-12.png){:height="100%" width="100%"}

![](/assets/github-13.png){:height="100%" width="100%"}

![](/assets/github-14.png){:height="100%" width="50%"}![](/assets/github-15.png){:height="100%" width="50%"}

## Bước 7: Cài đặt Page Rules

Chúng ta muốn 2 thứ: 

- Chuyển hướng tất cả các request từ www.yourcustomdomain.com thành yourcustomdomain.com
- Chuyển hướng tất cả các request từ http://yourcustomdomain.com thành https://yourcustomdomain.comments 

Vào page rules và tạo mới 1 rule: 

![](/assets/github-16.png){:height="100%" width="100%"}

http://www.yourcustomdomain.com/ thành chuyển hướng thành yourcustomdomain.com

![](/assets/github-17.png){:height="100%" width="100%"}

Save and Deploy.

http://yourcustomdomain.com/ chuyển hướng thành https://yourcustomdomain.com/

![](/assets/github-18.png){:height="100%" width="100%"}
Save and Deploy.

## Bước 8: Cài đặt HSTS

![](/assets/github-19.png){:height="100%" width="100%"}

Vào Crypto và kéo xuống đến phần HTTP Strict Transport Security (HSTS). Click **Enable HSTS**

![](/assets/github-20.png){:height="100%" width="100%"}


Đó là tất cả, và bây giờ chúng ta có một website tĩnh host trên GitHub có thể dễ dàng cập nhật,và đặc biệt là có SSL :smile: Chúc các bạn thành công.

>Happy Coding 

Tham khảo : [An illustrated guide to setting up your website using GitHub and Cloudflare](https://medium.freecodecamp.com/an-illustrated-guide-for-setting-up-your-website-using-github-cloudflare-5a7a11ca9465)





