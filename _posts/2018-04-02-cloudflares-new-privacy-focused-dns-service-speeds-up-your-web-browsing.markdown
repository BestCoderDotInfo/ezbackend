---
layout: post
author: derek
image: assets/images/feature_default.jpg
featured: false
hidden: false
title: Cloudflare’s new ‘privacy-focused’ DNS service speeds up your web browsing
date: '2018-04-02 08:00'
comments: true
categories: News
keywords: technology, news, cloudflare, tips, dns, 1.1.1.1
---

[Cloudflare](https://www.cloudflare.com) is launching its own consumer DNS service yesterday, on April Fools’ Day, that promises to speed up your internet connection and help keep it private.

![](/assets/02-04-2018/1.1.1.1.gif.mp4){:height="100%" width="100%"}

The service is using [https://1.1.1.1](https://1.1.1.1), and it’s not a joke but an actual DNS resolver that anyone can use. Cloudflare claims it will be “the Internet’s fastest, privacy-first consumer DNS service.” While OpenDNS and Google DNS both exist, Cloudflare is focusing heavily on the privacy aspect of its own DNS service with a promise to wipe all logs of DNS queries within 24 hours.

DNS services are typically provided by internet service providers to resolve a domain name like Google.com into a real IP address that routers and switches understand. It’s an essential part of the internet, but DNS servers provided by ISPs are often slow and unreliable. ISPs or any Wi-Fi network you connect to can also use DNS servers to identify all sites that are visited, which presents privacy problems. DNS also played an important role in helping Turkish citizens avoid a Twitter ban.

CLOUDFLARE WORKED WITH APNIC TO GET 1.1.1.1 WORKING
Cloudflare has worked with APNIC to offer its DNS service through 1.1.1.1 and 1.0.0.1. Lots of people have used 1.1.1.1 as a dummy address, and APNIC have tried in the past to analyze the flood of traffic to the IP address and been overwhelmed. “We talked to the APNIC team about how we wanted to create a privacy-first, extremely fast DNS system,” explains Cloudflare CEO Matthew Prince. “We offered Cloudflare’s network to receive and study the garbage traffic in exchange for being able to offer a DNS resolver on the memorable IPs. And, with that, 1.1.1.1 was born.”

Cloudflare’s DNS will offer support for both DNS-over-TLS and DNS-over-HTTPS, and the company is hoping that its HTTPS support will see more browsers and operating systems support the protocol. Cloudflare’s DNS is currently sitting at a global response time of 14ms, compared to 20ms for OpenDNS and 34ms for Google’s DNS, so it’s the fastest DNS resolver for consumers.

This isn’t the first time Cloudflare has helped the web with services. The web optimization network deployed its Universal SSL feature a few years ago to provide free SSL encryption to millions of websites, and the company is well known for offering DDoS protection to prevent sites from being overwhelmed by malicious traffic. If you’re interested in enabling Cloudflare’s new DNS you can find all the information over at [https://1.1.1.1](https://1.1.1.1).


