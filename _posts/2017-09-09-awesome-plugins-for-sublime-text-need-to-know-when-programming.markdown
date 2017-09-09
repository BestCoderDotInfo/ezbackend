---
title: Awesome packages for SublimeText 3 need to know when programming
layout: post
date: '2017-09-09 09:00:00'
comments: true
categories: programming
keywords: programming, sublimetext, packages, awesome packages for sublimetext, sublimetext awesome packages
excerpt: Sublime Text is one of the most popular code editors available right now. It is adored by many programmers for it's speed, simplicity, and rich plugin ecosystem.
---
Sublime Text is one of the most popular code editors available right now. It is adored by many programmers for it's speed, simplicity, and rich plugin ecosystem.

To help developers get the most our of Sublime, I decided to make a list containing some of the extensions that we use and love.

## Package Control

We cannot start off our list without mentioning Package Control. It's the plugin manager for Sublime, and without it installing and removing packages would be a huge pain. If you do not have Package Control installed, make sure you do that first, as it will allow you to quickly try out the other plugins in this article.

![](/assets/images/posts/09/1-package-control.png){:height="100%" width="100%"}

## Terminal

Shortcuts and menu entries for opening a terminal at the current file, or the current root project folder in Sublime Text.

[Install & Usage](https://github.com/wbond/sublime_terminal)

![](/assets/images/posts/09/terminal.png){:height="100%" width="100%"}

## Sublimetext automatic backups

When you edit text files (scripts, prose, whatever) you often find yourself wishing for an older version. Ever accidentally deleted a chunk from an important configuration file, or wished you could roll back a document a few hours? This plugin takes a copy of every file you save and copies it into a backup directory structure, ensuring that you never lose an old version of a file.

[Install & Usage](https://github.com/joelpt/sublimetext-automatic-backups)

## JavaScript & NodeJS Snippets

A collection of shorthand snippets for writing common JavaScript expressions much faster. Why write document.querySelector('selector'); when you can simply type qs, press tab, and let Sublime do the rest of the work.

![](/assets/images/posts/09/2-js-snippets.gif){:height="100%" width="100%"}

## Emmet

Like the previous plugin in the list, this one let's you use snippets to write code faster. The difference here is that instead of JS expresions, Emmet works for HTML and CSS, letting you write long tags, nested elements, or whole page templates in one go.

Emmet is a bit complex, so if want a simpler alternative you could try a similar plugin called [HTML Snippets](https://packagecontrol.io/packages/HTML%20Page%20Snippets). It has less features, but is way easier to use, and has a great straightforward documentation.

![](/assets/images/posts/09/3-emmet.gif){:height="100%" width="100%"}

## Advanced New File

This awesome package makes it possible to create new files blazingly fast. Instead of browsing folders, and using the menus, you simply open a prompt with super+alt+n and write the path to your new file. The plugin will also add any non-existing directories from the path, and even supports auto completion for folder names.

![](/assets/images/posts/09/5-new-file.gif){:height="100%" width="100%"}

## Git

A Git integration that works directly from the Sublime Text command palette. The package provides quick access to a number of commonly used Git commands, allowing developers to add files, make commits, or open the Git logs, without ever leaving Sublime.

![](/assets/images/posts/09/4-git-.png){:height="100%" width="100%"}

## GitGutter

Very useful extension that marks each line in your source code, telling you its Git status and giving you an overview of the changes that have occurred. GitGutter can be used to compare your files to the git HEAD, origin, a branch of your choice, or even certain commits.

![](/assets/images/posts/09/12-git.png){:height="100%" width="100%"}

## Side Bar Enhancement

In Sublime Text the project you are working on is overviewed in the left side panel. Although it gives you some options for working with your files, the available default actions are quite limited. This plugin changes that by adding over 20 options to the right-click menu, including Open in browser, duplicate, and lots of other useful stuff.

![](/assets/images/posts/09/7-sidebar.png){:height="100%" width="100%"}


## ColorPicker

A tiny, useful color picker that is very simple to use and great for quickly grabbing color hex values. The plugin opens in a separate window and allows you to choose a color from a palette or use an eye dropper to extract color from anywhere on your screen.

![](/assets/images/posts/09/8-colorpicker-.png){:height="100%" width="100%"}

## Placeholders

Sublime Text 3 has a built-in Lorem Ipsum generator that you can use for creating dummy text. The Placeholders plugin extends that functionality and allows ST to quickly generate for you placeholder images, forms, lists, and tables.

![](/assets/images/posts/09/9-placeholder.gif){:height="100%" width="100%"}

## DocBlockr

This is an extension for those of you who like to add detailed comments to function definitions. DocBlockr allows you to effortlessly generate descriptions for your functions including the parameters they take, the returned value, and variable types.

![](/assets/images/posts/09/10-docblockrgif.gif){:height="100%" width="100%"}

## SublimeCodeIntel

Code intelligence plugin that indexes your source files and enables you to find function definitions and jump to them. This extension works for a plethora of popular and not-so-popular programming languages.

![](/assets/images/posts/09/11-codeintel.gif){:height="100%" width="100%"}

## Minify

A code minifer and beautifier in one. Minify takes your current opened file, and creates a new .min or .pretty version of it in the same directory. Works with CSS, HTML, JavaScript, JSONs, and SVGs.

This package relies on external node.js libraries for minifying and beautifying, so you will need to install them separately:

```
npm install -g clean-css uglifycss js-beautify html-minifier uglify-js minjson svgo
```

![](/assets/images/posts/09/14-minifier.png){:height="100%" width="100%"}

## Sublime Linter

This package enables the code editor to check for syntax errors, bad practices, or other mistakes that the developer may have made. SublimeLinter itself just acts as a base framework for linting, so you also need to install separate plugins for each language you code in.

![](/assets/images/posts/09/6-linter.png){:height="100%" width="100%"}

## Color Highlighter

A feature you can see in many other IDEs and text editors, but is missing from Sublime, is color previews. Using the Color Highlighter extension you can enable it in ST, allowing you to see how all hex and RGBA values are translated to color, directly in your style sheets.

![](/assets/images/posts/09/14-highlighter.gif){:height="100%" width="100%"}

## Language Packs

Sublime Text has code highlighting for over 50 languages but there are some frameworks or niche web dev languages that are not supported yet. Thanks to the plugin nature of the editor, the community can create and distribute packs for any imaginable programming language:

- AngularJS
- TypeScript
- Babel (React)

![](/assets/images/posts/09/15-lang-packs.png){:height="100%" width="100%"}

# Bonus

## How to backup Sublime Text 3 settings

For PackageControl and Settings sync, follow these [instructions](https://packagecontrol.io/docs/syncing).

However, on Windows, symlinking the User folder breaks auto-reload for Settings changes.

To bypass this, either symlink the parent folder - Packages or the main Sublime Text 3 folder.

Make sure to gitignore/exclude everything but the User dir, as they are installed packages and caches.

Code:

```bash
cd "$env:AppData\Sublime Text 3\"
mkdir $env:UserProfile\Sublime
mv Packages $env:UserProfile\Sublime\
New-Item -ItemType "SymbolicLink" -Path "Packages" -Target "$env:UserProfile\Sublime"
```

Gitignore:

```bash
sublime/*
!sublime/User/
sublime/User/*
!*.sublime-settings
!*.sublime-keymap
```

Source: [https://github.com/arvindch](https://gist.github.com/arvindch/1322cf387d96633121c170e17890369b)