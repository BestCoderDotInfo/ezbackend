---
layout: post
author: derek
image: assets/trending-web-tech-2019-freecodecamp.jpeg
featured: false
hidden: false
title: "How to control large YAML file in Ruby or Rails"
date: 2019-09-22 16:00
comments: true
external-url:
categories: Ruby
keywords: Ruby, Rails, Web Development, TechTechnology
excerpt: Large Rails apps have large locale .yml files. Some have files so large that it is not feasible to simply open them in editor and work with them. Sure, you can edit them just fine, but you can’t efficiently search them.
---

# What is YAML file?

YAML ("YAML Ain't Markup Language") is a human-readable data-serialization language. It is commonly used for configuration files and in applications where data is being stored or transmitted. 

Anyway, YAML is a data serialization language designed for human interaction. It’s a strict superset of JSON, another data serialization language. But because it’s a strict superset, it can do everything that JSON can and more. One major difference is that newlines and indentation actually mean something in YAML, as opposed to JSON, which uses brackets and braces.

# How To Write YAML?

The basic structure of a YAML file is a map. You might call this a dictionary, hash or object, depending on your programming language or mood.

Very generally, it’s keys and values all the way down:

```bash
key: value
```

# How to Work With Large YAML Files and Not Go Crazy

Large Rails apps have large locale .yml files. Some have files so large that it is not feasible to simply open them in editor and work with them. Sure, you can edit them just fine, but you can’t efficiently search them.

```ruby
#! /usr/bin/env ruby

require 'yaml'
require 'colorize'

filename = ARGV[0]
pattern_text = ARGV[1]

unless filename && pattern_text
  puts "Usage: grep_yaml.rb filename pattern"
  exit(1)
end

pattern = Regexp.new(pattern_text, :nocase)
p pattern

hash = YAML.load_file(filename)

def recurse(obj, pattern, current_path = [], &block)
  case obj
  when String
    path = current_path.join('.')
    if obj =~ pattern || path =~ pattern
      yield [path, obj]
    end
  when Hash
    obj.each do |k, v|
      recurse(v, pattern, current_path + [k], &block)
    end
  end
end

recurse(hash, pattern) do |path, value|
  line = "#{path}:\t#{value}"
  line = line.gsub(pattern) {|match| match.green }
  puts line
end
```

#### That's all! Hope you find it useful.

Resources:
- [Wiki YAML](https://en.wikipedia.org/wiki/YAML)
- [What is YAML? A Beginner's Guide](https://circleci.com/blog/what-is-yaml-a-beginner-s-guide/)
- [How to Work With Large YAML Files and Not Go Crazy](http://tech.tulentsev.com/2014/04/work-with-large-yaml-files-and-not-go-crazy/)

