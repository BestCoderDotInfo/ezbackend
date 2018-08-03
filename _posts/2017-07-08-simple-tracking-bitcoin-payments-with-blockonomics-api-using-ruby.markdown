---
layout: post
author: derek
image: assets/images/feature_default.jpg
featured: false
hidden: false
title:  Simple Tracking Bitcoin Payments With Blockonomics API Using Ruby (Rails)
date: 2017-07-08 19:00
comments: true
external-url:
categories: programming
keywords: programming, ruby, rails, api, blockomomics, bitcoin, tracking
excerpt: Blockonomics is a decentralized and permissionless bitcoin payment solution. The API support you tracking bitcoin payments...
---
Blockonomics is a decentralized and permissionless bitcoin payment solution. The API support you tracking bitcoin payments wiht **Free Plan** such as: balance, history, transaction details.

Now, I will write some code with ruby using this API. Let's begin !

I call this function is BlockonomicsApiService, so we will have `blockomomics_api_service.rb` in `app/services` folder (Rails) or not:

```ruby
class BlockonomicsApiService

    require 'uri'
    require 'json'
    require 'net/http'
    require 'net/https'

    API_ENDPOINT = 'https://www.blockonomics.co/api/'

end
```

With libraries, it will help send request **GET** or **POST** with API endpoint and receiving the json response.

Define `API_ENDPOINT`.

Next, we need resolve send request with API endpoint. At [API Document](https://www.blockonomics.co/views/api.html), we can see it using 2 method GET & POST for Free plan.

```ruby
class BlockonomicsApiService
...
    def init_http(method, endpoint, params, body)
      url_string = method == 'GET' ? "#{endpoint}?#{params}" : endpoint
      uri = URI.parse(url_string)
      https = Net::HTTP.new(uri.host,uri.port)
      https.use_ssl = true
      case method
      when 'POST', 'post'
        req = Net::HTTP::Post.new(uri.request_uri, 'Content-Type' => 'application/json')
      when 'GET', 'get'
        req = Net::HTTP::Get.new(uri.request_uri, 'Content-Type' => 'application/json')
      end
      req.body = body.to_json
      res = https.request(req)
      res.is_a?(Net::HTTPSuccess) ? JSON.parse(res.body) : {}
    end
...
end
```

with POST, at balance example:

```javascript
curl -d '{"addr":"1dice8EMZmqKvrGE4Qc9bUFf9PX3xaYDp 1dice97ECuByXAvqXpaYzSaQuPVvrtmz6"}' https://www.blockonomics.co/api/balance
```

we see it have body is json format:

```json
{"addr":"1dice8EMZmqKvrGE4Qc9bUFf9PX3xaYDp 1dice97ECuByXAvqXpaYzSaQuPVvrtmz6"}
```

it is `req.body`

with GET, at Transaction Detail example:

```javascript
curl https://www.blockonomics.co/api/tx_detail?txid=5e4e03748327a22288623b02dab1721ac9f8082c7294aaa7f9581be49dced2c5
```

this `url` contain params `txid` with value is `5e4e03748327a22288623b02dab1721ac9f8082c7294aaa7f9581be49dced2c5`

Also, we have api endpoint: `https://www.blockonomics.co/api/tx_detail` but we need merge it with params.

```ruby
url_string = method == 'GET' ? "#{endpoint}?#{params}" : endpoint
```

Finally, we have new url api endpoint:

```
https://www.blockonomics.co/api/tx_detail?txid=5e4e03748327a22288623b02dab1721ac9f8082c7294aaa7f9581be49dced2c5
```

The, excute request with:

```ruby
res = https.request(req)
```
 and we have the response `res` contain: `res.body`, we need convert to json: `JSON.parse(res.body)`. But, we need check the response status is successfully or not with: `res.is_a?(Net::HTTPSuccess)`.

Look good! Next, we will implement functions with api endpoint.

1. Balance

Returns balance and unconfirmed amount(Amount waiting 2 confirmations) of multiple addresses. Balance units are in satoshis.

Definition

```javascript
POST https://www.blockonomics.co/api/balance
Request body: {"addr": <Whitespace seperated list of bitcoin addresses/xpubs>}
```

Example Request

```javascript
curl -d '{"addr":"1dice8EMZmqKvrGE4Qc9bUFf9PX3xaYDp 1dice97ECuByXAvqXpaYzSaQuPVvrtmz6"}' https://www.blockonomics.co/api/balance
```

Example Response

```javascript
{"response": [{"confirmed": 189412205, "addr": "1dice8EMZmqKvrGE4Qc9bUFf9PX3xaYDp", "unconfirmed": 012211 }, {"confirmed": 746599881, "addr": "1dice97ECuByXAvqXpaYzSaQuPVvrtmz6", "unconfirmed": 0}]}
```

And ruby function:

```ruby
def find_balance_by_address(*btc_address)
    endpoint = "#{API_ENDPOINT}balance".freeze
    body = {"addr":btc_address.join(" ")}
    json_response = init_http('POST', endpoint, nil, body)
end
```

2. History

See more at [API Document](https://www.blockonomics.co/views/api.html#history) :

```ruby
def find_history_by_address(*btc_address)
    endpoint = "#{API_ENDPOINT}searchhistory".freeze
    body = {"addr":btc_address.join(" ")}
    json_response = init_http('POST', endpoint, nil, body)
end
```

3. Transaction Detail

See more at [API Document](https://www.blockonomics.co/views/api.html#txdetail) :

```ruby
def find_transdt_by_txid(txid)
    endpoint = "#{API_ENDPOINT}tx_detail".freeze
    params = {txid: txid}
    json_response = init_http('GET', endpoint, params, {})
end
```


Finally,

```ruby
class BlockonomicsApiService

    require 'uri'
    require 'json'
    require 'net/http'
    require 'net/https'

    API_ENDPOINT = 'https://www.blockonomics.co/api/'

    class << self

        def init_http(method, endpoint, params, body)
          url_string = method == 'GET' ? "#{endpoint}?#{params}" : endpoint
          uri = URI.parse(url_string)
          https = Net::HTTP.new(uri.host,uri.port)
          https.use_ssl = true
          case method
          when 'POST', 'post'
            req = Net::HTTP::Post.new(uri.request_uri, 'Content-Type' => 'application/json')
          when 'GET', 'get'
            req = Net::HTTP::Get.new(uri.request_uri, 'Content-Type' => 'application/json')
          end
          req.body = body.to_json
          res = https.request(req)
          res.is_a?(Net::HTTPSuccess) ? JSON.parse(res.body) : {}
        end

        def find_balance_by_address(*btc_address)
            endpoint = "#{API_ENDPOINT}balance".freeze
            body = {"addr":btc_address.join(" ")}
            json_response = init_http('POST', endpoint, nil, body)
        end

        def find_history_by_address(*btc_address)
            endpoint = "#{API_ENDPOINT}searchhistory".freeze
            body = {"addr":btc_address.join(" ")}
            json_response = init_http('POST', endpoint, nil, body)
        end

        def find_transdt_by_txid(txid)
            endpoint = "#{API_ENDPOINT}tx_detail".freeze
            params = {txid: txid}
            json_response = init_http('GET', endpoint, params, {})
        end

    end

end

```

Using:

```ruby

BlockonomicsApiService.find_balance_by_address('1dice8EMZmqKvrGE4Qc9bUFf9PX3xaYDp', '1dice97ECuByXAvqXpaYzSaQuPVvrtmz6')

BlockonomicsApiService.find_history_by_address('1dice8EMZmqKvrGE4Qc9bUFf9PX3xaYDp', '1dice97ECuByXAvqXpaYzSaQuPVvrtmz6')

BlockonomicsApiService.find_transdt_by_txid('5e4e03748327a22288623b02dab1721ac9f8082c7294aaa7f9581be49dced2c5')

```

>Happy Coding!