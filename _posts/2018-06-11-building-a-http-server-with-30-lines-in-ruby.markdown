---
title: Building A HTTP Server With 30 Lines In Ruby
date: '2018-06-11 15:00:00'
comments: true
external-url: null
categories: Ruby
keywords: programing, developer, ruby, http, server
excerpt: Web servers, and HTTP in general, might seem difficult to understand. How does the browser format a request, and how does the response get sent to the user? In this Ruby Magic episode we’ll learn how a to build a simple Ruby HTTP server in 30 lines of code. When we’re done, our server will handle HTTP GET requests and we’ll use it to serve a simple Rack app.
---
Web servers, and HTTP in general, might seem difficult to understand. How does the browser format a request, and how does the response get sent to the user? In this Ruby Magic episode we’ll learn how a to build a simple Ruby HTTP server in 30 lines of code. When we’re done, our server will handle HTTP GET requests and we’ll use it to serve a simple Rack app.

# How HTTP and TCP work together?

**TCP** is a transport protocol that describes how a server and a client exchange data.

**HTTP** is a request-response protocol that specifically describes how web servers exchange data with HTTP clients or web browsers. HTTP commonly uses TCP as its transport protocol. In essence, an HTTP server is a TCP server that “speaks” HTTP.

```ruby
# tcp_server.rb
require 'socket'
server = TCPServer.new 5678

while session = server.accept
  session.puts "Hello world! The time is #{Time.now}"
  session.close
end
```

In this example of a TCP server, the server binds to port `5678` and waits for a client to connect. When that happens, it sends a message to the client, and then closes the connection. After it’s done talking to the first client, the server waits for another client to connect to send its message to again.

```ruby
# tcp_client.rb
require 'socket'
server = TCPSocket.new 'localhost', 5678

while line = server.gets
  puts line
end

server.close
```

To connect to our server, we’ll need a simple TCP client. This example client connects to the same port (`5678`) and uses `server.gets` to receive data from the server, which is then printed. When it stops receiving data, it closes the connection to the server and the program will exit.

When you start the server server is running (`$ ruby tcp_server.rb`), you can start the client in a separate tab to receive the server’s message.

```bash
$ ruby tcp_client.rb
Hello world! The time is 2016-11-23 15:17:11 +0100
$
```

With a bit of imagination, our TCP server and client work somewhat like a web server and a browser. The client sends a request, the server responds, and the connection is closed. That’s how the request-response pattern works, which is exactly what we need to build an HTTP server.

Before we get to the good part, let’s look at what HTTP requests and responses look like.

## The simplest HTTP GET request

The simplest HTTP GET request is a request-line without any additional headers or a request body.

```bash
GET / HTTP/1.1\r\n
```

The Request-Line consists of four parts:

- A method token (GET, in this example)
- The Request-URI (/)
- The protocol version (HTTP/1.1)
- A CRLF (a carriage return: \r, followed by line feed: \n) to indicate the end of the line

The server will respond with an HTTP response, which may look like this:

```bash
HTTP/1.1 200\r\nContent-Type: text/html\r\n\r\n\Hello world!
```

This response consists of:

- A status line: the protocol version (“HTTP/1.1”), followed by a space, the response’s status code (“200”), and terminated with a CRLF (\r\n)
- Optional header lines. In this case, there’s only one header line (“Content-Type: text/html”), but there could be multiple (separated with with a CRLF: \r\n)
- A newline (or a double CRLF) to separate the status line and header from the body: (\r\n\r\n)
- The body: “Hello world!”

## A simple Ruby HTTP server

Enough talk. Now that we know how to create a TCP server in Ruby and what some HTTP requests and responses look like, we can build a really simple HTTP server. You’ll notice that the web server looks mostly the same as the TCP server we discussed earlier. The general idea is the same, we’re just using the HTTP protocol to format our message. Also, because we’ll use a browser to send requests and parse responses, we won’t have to implement a client this time.

```ruby
# http_server.rb
require 'socket'
server = TCPServer.new 5678

while session = server.accept
  request = session.gets
  puts request

  session.print "HTTP/1.1 200\r\n" # 1
  session.print "Content-Type: text/html\r\n" # 2
  session.print "\r\n" # 3
  session.print "Hello world! The time is #{Time.now}" #4

  session.close
end
```

After the server receives a request, like before, it uses session.print to send a message back to the client: Instead of just our message, it prefixes the response with a status line, a header and a newline:

- 1 .The status line (HTTP 1.1 200\r\n) to tell the browser that the HTTP version is 1.1 and the response code is “200”
- 2 .A header to indicate that the response has a text/html content type (Content-Type: text/html\r\n)
- 3 .The newline (\r\n)
- 4 . The body: “Hello world! …”

Like before, it closes the connection after sending the message. We’re not reading the request yet, so it just prints it to the console for now.

If you start the server and open http://localhost:5678 in your browser, you should see the “Hello world! …”-line with the current time, like we received from our TCP client earlier. 🎉

## Serving a Rack app

Until now, our server has been returning a single response for each request. To make it a little more useful, we could add more responses to our server. Instead of adding these to the server directly, we’ll use a Rack app. Our server will parse HTTP requests and pass them to the Rack app, which will then return a response for the server to send back to the client.

Rack is an interface between web servers that support Ruby and most Ruby web frameworks like Rails and Sinatra. In its simplest form, a Rack app is an object that responds to `call` and returns a “tiplet”, an array with three items: an HTTP response code, a hash of HTTP headers and a body.

```ruby
app = Proc.new do |env|
  ['200', {'Content-Type' => 'text/html'}, ["Hello world! The time is #{Time.now}"]]
end
```

In this example, the response code is “200”, we’re passing “text/html” as the content type through the headers, and the body is an array with a string.

To allow our server to serve responses from this app, we’ll need to turn the returned triplet into a HTTP response string. Instead of always returning a static response, like we did before, we’ll now have to build the response from the triplet returned by the Rack app.

```ruby
# http_server.rb
require 'socket'

app = Proc.new do
  ['200', {'Content-Type' => 'text/html'}, ["Hello world! The time is #{Time.now}"]]
end

server = TCPServer.new 5678

while session = server.accept
  request = session.gets
  puts request

  # 1
  status, headers, body = app.call({})

  # 2
  session.print "HTTP/1.1 #{status}\r\n"

  # 3
  headers.each do |key, value|
    session.print "#{key}: #{value}\r\n"
  end

  # 4
  session.print "\r\n"

  # 5
  body.each do |part|
    session.print part
  end
  session.close
end
```

To serve the response we’ve received from the Rack app, there’s some changes we’ll make to our server:

- 1.Get the status code, headers, and body from the triplet returned by app.call.
- 2.Use the status code to build the status line
- 3.Loop over the headers and add a header line for each key-value pair in the hash
- 4.Print a newline to separate the status line and headers from the body
- 5.Loop over the body and print each part. Since there’s only one part in our body array, it’ll simply print our “Hello world”-message to the session before closing it.

## Reading requests

Until now, our server has been ignoring the `request` variable. We didn’t need to as our Rack app always returned the same response.

`Rack::Lobster` is an example app that ships with Rack and uses request URL parameters in order to function. Instead of the Proc we used as an app before, we’ll use that as our testing app from now on.

```ruby
# http_server.rb
require 'socket'
require 'rack'
require 'rack/lobster'

app = Rack::Lobster.new
server = TCPServer.new 5678

while session = server.accept
# ...
```

Opening the browser will now show a lobster instead of the boring string it printed before. Lobstericious!

The “flip!” and “crash!” links link to `/?flip=left` and `/?flip=crash` respectively. However, when following the links, the lobster doesn’t flip and nothing crashes just yet. That’s because our server doesn’t handle query strings right now. Remember the `request` variable we ignored before? If we look at our server’s logs, we’ll see the request strings for each of the pages.

```bash
GET / HTTP/1.1
GET /?flip=left HTTP/1.1
GET /?flip=crash HTTP/1.1
```

The HTTP request strings include the request method (“GET”), the request path (`/, /?flip=left` and `/?flip=crash`), and the HTTP version. We can use this information to determine what we need to serve.

```ruby
# http_server.rb
require 'socket'
require 'rack'
require 'rack/lobster'

app = Rack::Lobster.new
server = TCPServer.new 5678

while session = server.accept
  request = session.gets
  puts request

  # 1
  method, full_path = request.split(' ')
  # 2
  path, query = full_path.split('?')

  # 3
  status, headers, body = app.call({
    'REQUEST_METHOD' => method,
    'PATH_INFO' => path,
    'QUERY_STRING' => query
  })

  session.print "HTTP/1.1 #{status}\r\n"
  headers.each do |key, value|
    session.print "#{key}: #{value}\r\n"
  end
  session.print "\r\n"
  body.each do |part|
    session.print part
  end
  session.close
end
```

To parse the request and send the request parameters to the Rack app, we’ll split the request string up and send it to the Rack app:

- 1.Split the request string into a method and a full path
- 2.Split the full path into a path and a query
- 3.Pass those to our app in a Rack environment hash.

For example, a request like GET `/?flip=left HTTP/1.1\r\n` will be passed to the app like this:

```bash
{
  'REQUEST_METHOD' => 'GET',
  'PATH_INFO' => '/',
  'QUERY_STRING' => '?flip=left'
}
```

Restarting our server, visiting `http://localhost:5678`, and clicking the “flip!”-link will now flip the lobster, and clicking the “crash!” link will crash our web server.

We’ve just scratched the surface of implementing a HTTP server, and ours is only 30 lines of code, but it explains the basic idea. It accepts GET requests, passes the request’s attributes to a Rack app, and sends back responses to the browser. Although it doesn’t handle things like request streaming and POST requests, our server could theoretically be used to serve other Rack apps too.

This concludes our quick look into building an HTTP server in Ruby. If you want to play around with our server, here’s the [code](https://gist.github.com/jeffkreeftmeijer/7f08d1f7e381b9c552666750914925eb).

### Resources:

- https://blog.appsignal.com/2016/11/23/ruby-magic-building-a-30-line-http-server-in-ruby.html