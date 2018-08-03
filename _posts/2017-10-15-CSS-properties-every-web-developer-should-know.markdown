---
layout: post
author: derek
image: assets/images/feature_default.jpg
featured: false
hidden: false
title: CSS properties every web developers should know
date: '2017-10-15 10:00'
comments: true
categories: Web Development
keywords: programming, developer should know, css, css properties, css tricks
excerpt: There are CSS properties, such as background images, border images, masking, and clipping properties, with which you can directly add images to web pages and control their behavior. However, there are also less frequently mentioned image-related CSS properties that work on images added with the <img> HTML tag, which the preferred way of adding images to web pages.
---

There are CSS properties, such as background images, border images, masking, and clipping properties, with which you can directly add images to web pages and control their behavior. However, there are also less frequently mentioned image-related CSS properties that work on images added with the <img> HTML tag, which the preferred way of adding images to web pages.

## 1. Sharpen images with image-rendering

The `image-rendering` property defines how the browser should render an image if it is scaled up or down from its original dimensions. By default, each browser will attempt to apply aliasing to this scaled image in order to prevent distortion, but this can sometimes be a problem if we want the image to preserve its original pixelated form.

```css
img {
  image-rendering: auto;
  image-rendering: crisp-edges;
  image-rendering: pixelated;
}
```

- auto: default value that uses the browser’s standard algorithm to maximize the appearance of an image.
- crisp-edges: the contrast, colors and edges of the image will be preserved without any smoothing or blurring. According to the spec this was specifically intended for pixel art. This value applies to images scaled up or down.
- pixelated: as the image changes size the browser will preserve its pixelated style by using nearest-neighbour scaling. This value only applies to images that are scaled up.


## 2. Stretch images with object-fit

The `object-fit` property defines how an element responds to the height and width of its content box. It's intended for images, videos and other embeddable media formats in conjunction with the object-position property. Used by itself, `object-fit` lets us crop an inline image by giving us fine-grained control over how it squishes and stretches inside its box.

- fill: this is the default value which stretches the image to fit the content box, regardless of its aspect-ratio.
- contain: increases or decreases the size of the image to fill the box whilst preserving its aspect-ratio.
- cover: the image will fill the height and width of its box, once again maintaining its aspect ratio but often cropping the image in the process.
- none: image will ignore the height and width of the parent and retain its original size.
- scale-down: the image will compare the difference between none and contain in order to find the smallest concrete object size.

```css
img {
  height: 120px;
}

.cover {
  width: 260px;
  object-fit: cover;
}
```

## 3. Shift images with object-position

The `object-position` property is used in conjunction with object-fit property and defines how an element such as a video or image is positioned with X/Y coordinates inside its content-box. This property takes two numerical values, such as 0 10% or 0 10px. In those examples, the first number indicates that the image should be placed at the left of the content box (0), the second that it should be positioned 10% or 10px from the top. It is also possible to use negative values.

The default value for object-position is 50% 50% when using the object-fit property on an image, so that by default all images are positioned in the center of their content box, like so:

```css
img {
  object-fit: none;
  object-position: 50% 50%; /* default value: image is centered*/
  object-position: 0 0; /* positioned top left of the content box */
}
```


## 4. Situate images with

The `vertical-align` property in CSS controls how elements set next to each other on a line are lined up.

```css
img {
  vertical-align: middle;
}
```
In order for this to work, the elements need to be set along a baseline. As in, inline (e.g. <span>, <img>) or inline-block (e.g. as set by the display property) elements.

- baseline - This is the default value.
- top - Align the top of the element and its descendants with the top of the entire line.
- bottom - Align the bottom of the element and its descendants with the bottom of the entire line.
- middle - Aligns the middle of the element with the middle of lowercase letters in the parent.
- text-top - Aligns the top of the element with the top of the parent element's font
- text-bottom - Aligns the bottom of the element with the bottom of the parent element's font.
- sub - Aligns the baseline of the element with the subscript-baseline of its parent. Like where a <sub> would sit.
- super - Aligns the baseline of the element with the superscript-baseline of its parent. Like where a <sup> would sit.
- length - Aligns the baseline of the element at the given length above the baseline of its parent. (e.g. px, %, em, rem, etc.)

## 5. Shadow images with filter: drop-shadow()

CSS Filters are a powerful tool that authors can use to achieve varying visual effects (sort of like Photoshop filters for the browser). The CSS filter property provides access to effects like blur or color shifting on an element’s rendering before the element is displayed. Filters are commonly used to adjust the rendering of an image, a background, or a border.

The syntax is:

```css
.filter-me {
  filter: <filter-function> [<filter-function>]* | none
}
```

Where is one of:

- blur()
- brightness()
- contrast()
- drop-shadow()
- grayscale()
- hue-rotate()
- invert()
- opacity()
- saturate()
- sepia()
- url() - for applying SVG filters

Multiple functions can be used, space separated.

### drop-shadow()

```css
filter: drop-shadow(<length>{2,3} <color>?)
```

Applies a drop shadow effect to the input image. A drop shadow is effectively a blurred, offset version of the input image's alpha mask drawn in a particular color, composited below the image. The function accepts a parameter of type (defined in CSS3 Backgrounds), with the exception that the ‘inset’ keyword is not allowed.

This function is similar to the more established box-shadow property; the difference is that with filters, some browsers provide hardware acceleration for better performance.

See the Pen Contrast Filter by GRAY GHOST (@grayghostvisuals) on CodePen.

Drop-shadow also mimics the intended objects outline naturally unlike box-shadow that treats only the box as its path.

See the Pen Drop-shadow vs box-shadow (2) by Kseso (@Kseso) on CodePen.

Just like with text-shadow, there is no ‘spread’ parameter to create a solid shadow larger than the object.

> That's all. Happy codeing!

Resources:

- https://css-tricks.com/almanac/properties/i/image-rendering/
- https://css-tricks.com/almanac/properties/o/object-fit/
- https://css-tricks.com/almanac/properties/o/object-position/
- https://css-tricks.com/almanac/properties/v/vertical-align/
- https://css-tricks.com/almanac/properties/f/filter/
