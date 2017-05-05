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
