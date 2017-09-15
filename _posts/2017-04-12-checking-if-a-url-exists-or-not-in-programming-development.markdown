---
title: Checking if a URL exists or not in programming development ? 
date: 2017-04-12 16:00
comments: true
external-url: 
categories: Tricks
keywords: Ruby, PHP, Python, Shell, Pure Bash
excerpt: This is a trick help you checking if a URL exist or not in Ruby, PHP, Python, Shell, Bash,...
---
## Ruby

```ruby
  require 'net/http'
  require 'open-uri'

  def working_url?(url_str)
    url = URI.parse(url_str)
    Net::HTTP.start(url.host, url.port) do |http|
      http.head(url.request_uri).code == '200'
    end
  rescue
    false
  end
```

## PHP
**1. Using get_headers Function :**

```php
<?php
$url = "http://www.domain.com/demo.jpg";
$headers = @get_headers($url);
if(strpos($headers[0],'404') === false)
{
  echo "URL Exists";
}
else
{
  echo "URL Not Exists";
}
?>
```

**2. Using CURL**

```php
<?php
$url = "http://www.domain.com/demo.jpg";
$curl = curl_init($url);
curl_setopt($curl, CURLOPT_NOBODY, true);
$result = curl_exec($curl);
if ($result !== false)
{
  $statusCode = curl_getinfo($curl, CURLINFO_HTTP_CODE);
  if ($statusCode == 404)
  {
    echo "URL Not Exists"
  }
  else
  {
     echo "URL Exists";
  }
}
else
{
  echo "URL not Exists";
}
?>
```

## Python

```python
from urllib2 import urlopen
code = urlopen("https://kipalog.com").code
if code == 200:
   print "Exists!"
# via http://stackoverflow.com/questions/1966086/how-can-i-determine-if-anything-at-the-given-url-does-exist
```

or 

```python
import urllib2
ret = urllib2.urlopen('https://kipalog.com')
if ret.code == 200:
    print "Exists!"
# via http://stackoverflow.com/questions/7347888/how-do-i-check-if-a-file-on-http-exists-using-python-django
```

## Shell

```shell
#!/bin/bash

http_code=$(curl -I -s -o /dev/null -w "%{http_code}" "https://kipalog.com/")

if [ "$http_code" == "200" ]; then
  echo "Exist!!!"
fi
```


## Quick test

```shell
$ curl -I "https://kipalog.com" | grep Status
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0 20255    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
Status: 200 OK
```


## Pure Bash 2.0.4 and later

```bash
exec 3<>/dev/tcp/www.domain.com/80
printf "GET /demo.jpg HTTP/1.1\n" >&3
IFS= read -r line <&3
case $line in *200*);;*) false;;esac && echo exist || echo not exist
```
