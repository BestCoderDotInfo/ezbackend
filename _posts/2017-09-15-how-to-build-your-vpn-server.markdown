---
layout: post
author: derek
image: assets/images/feature_default.jpg
featured: false
hidden: false
title: How to build your VPN server?
date: '2017-09-15 16:00:00'
comments: true
categories: Server
keywords: programming, server, vpn, digital ocean, IPsec VPN, Ubuntu, free vpn
excerpt: An IPsec VPN encrypts your network traffic, so that nobody between you and the VPN server can eavesdrop on your data as it travels via the Internet. This is especially useful when using unsecured networks, e.g. at coffee shops, airports or hotel rooms.
---

Set up your own IPsec VPN server in just a few minutes, with both IPsec/L2TP and Cisco IPsec on Ubuntu.

An IPsec VPN encrypts your network traffic, so that nobody between you and the VPN server can eavesdrop on your data as it travels via the Internet. This is especially useful when using unsecured networks, e.g. at coffee shops, airports or hotel rooms.

Follow: [https://github.com/hwdsl2/setup-ipsec-vpn](https://github.com/hwdsl2/setup-ipsec-vpn)

## 1. Create your Cloud computing

I recommend you use Digital Ocean with minimum price each month is $5. Only charge when you turn on your Cloud computing.

Get **free $10** for first time register via [https://m.do.co/c/78b373914eb8](https://m.do.co/c/78b373914eb8).

After you have been successfully registered, you can create new Droplet in dashboard.

![](/assets/15-09-2017/1.png)

![](/assets/15-09-2017/2.png)

![](/assets/15-09-2017/3.png)

![](/assets/15-09-2017/4.png)

After your click button **Create**, wait a moment. Now, we have a Droplet, click it's name to access your **Droplet**. You will look like:

![](/assets/15-09-2017/5.png)

You can turn **ON/OFF** your Droplet at here.

![](/assets/15-09-2017/6.png)

You will receive one email with your Droplet information:
- ip
- username
- password

Remember it.


## 2. Config your Droplet

First, connect your Droplet via Terminal:

```
ssh your-username@droplet-ip-address
```

Then type your password.

Note: In first time login, you need change your password.


### Bonus:

Add new user:

```
sudo adduser USERNAME
```

If user has been exists:

```
sudo passwd USERNAME
```

Enable password authentication by editing `/etc/ssh/sshd_config`: change `PasswordAuthentication no` to `PasswordAuthentication yes`.

Use command as **root** user

By default, a new user is only in their own group, which is created at the time of account creation, and shares a name with the user. In order to add the user to a new group, we can use the usermod command:

```
usermod -aG sudo USERNAME
```

Open:

```
sudo vi /etc/sudoers
```

and add:

```
USERNAME ALL=(ALL:ALL) ALL
```

Then restart:

```
sudo /etc/init.d/ssh restart
```

## 3. Setup VPN

First, update your system with ``apt-get update && apt-get dist-upgrade`` and reboot. This is optional, but recommended.

To install the VPN, please choose one of the following options:

**Option 1**: Have the script generate random VPN credentials for you (will be displayed when finished):

```
wget https://git.io/vpnsetup -O vpnsetup.sh && sudo sh vpnsetup.sh
```

**Option 2:** Edit the script and provide your own VPN credentials:

```
wget https://git.io/vpnsetup -O vpnsetup.sh
nano -w vpnsetup.sh
[Replace with your own values: YOUR_IPSEC_PSK, YOUR_USERNAME and YOUR_PASSWORD]
sudo sh vpnsetup.sh
```

**Option 3:** Define your VPN credentials as environment variables:

```
# All values MUST be placed inside 'single quotes'
# DO NOT use these characters within values:  \ " '
wget https://git.io/vpnsetup -O vpnsetup.sh && sudo \
VPN_IPSEC_PSK='your_ipsec_pre_shared_key' \
VPN_USER='your_vpn_username' \
VPN_PASSWORD='your_vpn_password' sh vpnsetup.sh
```

**Note:** If unable to download via wget, you may also open [vpnsetup.sh](https://github.com/hwdsl2/setup-ipsec-vpn/blob/master/vpnsetup.sh) (or [vpnsetup_centos.sh](https://github.com/hwdsl2/setup-ipsec-vpn/blob/master/vpnsetup_centos.sh)) and click the Raw button. Press `Ctrl-A` to select all, `Ctrl-C` to copy, then paste into your favorite editor.

## 4. Setup Firewall on Digital Ocean

Create **Firewall** and config look like:

![](/assets/15-09-2017/7.png)

![](/assets/15-09-2017/8.png)

## 5. Configure IPsec/L2TP VPN Clients

Visit: [Configure IPsec/L2TP VPN Clients](https://github.com/hwdsl2/setup-ipsec-vpn/blob/master/docs/clients.md).

This is all! Thanks for your reading 🙇


