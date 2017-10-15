---
title: How to setup virtual host on MacOS
date: '2017-10-15 09:00'
comments: true
categories: Programming
keywords: programming, developer should know, virtual host, fake domain on locally
excerpt: This is a simple guide on how to setup virtual host on Mac OS. Same as many developers, I spend a lot of time with localhost and I feeling boring right now.
---

This is a simple guide on how to setup virtual host on Mac OS. Same as many developers, I spend a lot of time with localhost and I feeling boring right now.

Before getting started, we should know Virtual Host and why use it?

## What are virtual hosts used for?

In the case of this tutorial, we will create a virtual host to let us access a subfolder on our local web server at a custom address like `http://application` or `http://application.loc` instead of having to type in a long URL like `http://localhost/~username/application`

## Preparation

Of course, one MacBook with Mac OS already installed. 

This articel, I using setup from [https://github.com/virtualhost/virtualhost.sh](https://github.com/virtualhost/virtualhost.sh). Go to [this link](https://raw.githubusercontent.com/virtualhost/virtualhost.sh/master/virtualhost.sh) and download or coppy `virtualhost.sh` file to your locally.

You need run this command after download or coppy file:

```
chmod a+x virtualhost.sh
```

Then, type on terminal with sudo:

```
sudo virtualhost.sh your-domain-name-you-want
```

and answer `Yes` to finish.

Now you can access your localhost with `your-domain`. Default port is `80`, if you have another port: `your-domain:3000`.

Read more document at [here](https://github.com/virtualhost/virtualhost.sh/wiki).

That's all.

Thank for your reading. Happy coding!

Resources:

- [virtualhost.sh](https://github.com/virtualhost/virtualhost.sh)

