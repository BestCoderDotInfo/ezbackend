---
layout: post
author: derek
image: assets/action-text.jpg
featured: false
hidden: false
title: "How I resolve problem with Rails 6 Trix Editor (Action Text)"
date: 2019-10-08 21:00
comments: true
external-url:
categories: Rails
keywords: Ruby, Rails, Web Development, TechTechnology, Rails 6, Trix Editor
excerpt: Action Text is a brand new framework coming to Rails 6 that’s going to make creating, editing, and displaying rich text content in your applications super easy. It’s an integration between the Trix editor, Active Storage-backed file and image processing, and a text-processing flow that ties it all together. With Action Text, you really shouldn’t ever have to impoverish your users with a vanilla textarea ever again!
---
Action Text is a brand new framework coming to Rails 6 that’s going to make creating, editing, and displaying rich text content in your applications super easy. It’s an integration between the Trix editor, Active Storage-backed file and image processing, and a text-processing flow that ties it all together. With Action Text, you really shouldn’t ever have to impoverish your users with a vanilla textarea ever again!

As you see, Action Text now available on Rails 6.

## What's my problem?

When Rails 6 released, I am getting setup. I try using Action Text (Trix Editor). Everything still ok, but when I write post with body contain image upload from Trix editor, error is coming. Looks:

![](/assets/trix-error-1.png)

So, I try research Rails [Active Storage documentation](https://edgeguides.rubyonrails.org/active_storage_overview.html#transforming-images) and I see.

I add the image_processing gem to my Gemfile. And it work fine.

Otherwise, I don't adding `image_processing`. I replace some code at `app/views/active_storage/blobs/_blob.html.erb`. And it still work fine.

```ruby
<%= image_tag blob.variant(resize: local_assigns[:in_gallery] ? "800x600" : "1024x768") %>
```

![](/assets/trix-success.png)

## Conclusion

You must read  and understand document when getting start anything. This is best way for you.