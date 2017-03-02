---
layout: post
title: Synthesis Of Somethings That Need To Know When Programming 
date: 2017-02-21 18:00
comments: true
external-url:
categories: Programming
---
>Synthesis Of Somethings That Need To Know When Programming 

## 1. Sublime Tex Settings – User
`References -> Settings – User`

```
{
  "auto_complete": true,
  "bold_folder_labels": true,
  "caret_style": "phase",
  "close_windows_when_empty": true,
  "draw_white_space": "all",
  "find_selected_text": true,
  "fold_buttons": false,
  "font_size": 13,
  "highlight_line": true,
  "highlight_modified_tabs": true,
  "ignored_packages":
  [
  ],
  "line_padding_bottom": 1,
  "line_padding_top": 1,
  "scroll_past_end": true,
  "tab_completion": false,
  "tab_size": 2,
  "translate_tabs_to_spaces": true,
  "trim_automatic_white_space": true,
  "trim_trailing_white_space_on_save": true,
  "update_check": false,
  "vintage_start_in_command_mode": false,
  "word_wrap": true
}
```

`References -> Key bindings - User`

```
[
  { "keys": ["f12"], "command": "reindent"},
  { "keys": ["ctrl+shift+t"], "command": "delete_trailing_spaces" },
  { "keys": ["ctrl+shift+a"], "command": "toggle_side_bar" },
  { "keys": ["ctrl+1"], "command": "toggle_setting", "args": {"setting": "gutter"} }
]

```

- Youn can open Sublime via Terminal (OSX). Run : 

`ln -s "/Applications/Sublime Text.app/Contents/SharedSupport/bin/subl" /usr/local/bin/sublime`

and use `sublime ~/Documents` or `cd ~/Documents && sublime .`.

## 2. Git workflow

### Branches Overview

**master** branch
- `master` is a main branch of project. It must always be working-code, production-deployable branch.
- Every commit in `master` considered as a release version.
NO direct edit, coding on `master`
- `master` will be used when deploy code with a tag information.

**develop** branch

- `develop` is a main branch for developing the project. All feature branches must be create from develop and merge back to develop after finish.
- `develop` contains latest code of the project but not yet merge to master for releasing
- `Only Lead` can merge feature branches in to develop
- `Lead will merge develop into maste` when it comes to release.


**feature** branch

- May branch of from develop . Must be merge back to develop
- All development must be made on feature branch
- Naming of feature branch can be varied. However, you should follow the following naming convention: devname_featurename_issuenumber
- After finish developing on feature branch, Dev must create a pull request on Github to askLead to merge back into develop branch.

**hotfix** branch

- May branch of from master . Must be merge back to develop and master
hotfix is created from master branch for a purpose of critical bug fix for production.
- Do not use this branch for feature developing
Code developed in this branch, also need merge request in order to merge into masterbranch
- Naming convention: `hotfix-X.X.X`

### Git branching work-flow in details

When PJ start, there’re already **master** branch
Before start developing **Lead** will create a **develop** branch

```
$ git clone git@github.com:GoldenOwlAsia/myproject.git
$ cd myproject
$ git checkout -b develop
$ git push origin develop
```

- **Dev** will create a feature branch from develop when develop a new feature

```
$ git checkout -b devname_feat_xxx develop
# Develop some thing
$ git add ChangeFiles
$ git commit -m "Develop feature xxx"
$ git push origin devname_feat_xxx
```

- After finishing feature, Dev will ask Lead to merge back feature branch into develop by open `New pull Request` on Github. After receiving merge request, **Lead will do code review and merge feature branch into develop**. Do merging with no fast-forward (no-ff).

```
$ git checkout develop
$ git merge --no-ff devname_feat_xxx
# If have any conflict, please resolve it before moving on
$ git push origin develop
```


```
# Delete feature branch after merging
$ git branch -d devname_feat_xxx
$ git push origin devname_feat_xxx
```

```
$ git checkout master
$ git merge --no-ff develop
# If have any conflict, please resolve it before moving on
# Create a tag for further reference
$ git tag -a 1.0.0
$ git push origin master
```

- If critical bugs found on production and must be fixed ASAP, then Lead will create a hotfixbranch from `master`

```
$ git checkout -b hotfix-1.0.1 develop
# Do bugfixes, change release metadata, version
$ git add ChangeFiles
$ git commit -m "Fix critical bugs"
$ git push origin hotfix-1.0.0
```

- Lead will do code review and merge hotfix branch into develop and master. Do merging with no fast-forward (no-ff).

```
$ git checkout develop
$ git merge --no-ff hotfix-1.0.1
# If have any conflict, please resolve it before moving on
$ git push origin develop
```


```
$ git checkout master
$ git merge --no-ff hotfix-1.0.1
# If have any conflict, please resolve it before moving on
# Create a tag for further reference
$ git tag -a 1.0.1
$ git push origin master
```

```
# Delete hotfix branch
$ git branch -d hotfix-1.0.1
$ git push origin hotfix-1.0.1
```

>Happy Coding !