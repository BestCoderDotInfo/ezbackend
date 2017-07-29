# Introduce Bestcoder.info
---

- This is a opensource for my blog: [bestcoder.info](https://www.bestcoder.info).

- The **Bestcoder.info** hosted on Github with Github Page.

- This project has been using [Jekyll](http://jekyllrb.com/) and theme by [dbyll](https://github.com/dbtek/dbyll). Builded by:

[Ruby lang](https://www.ruby-lang.org) |  [Go lang](https://golang.org)
:-------------------------:|:-------------------------:
![](assets/readme/ruby.png)|![](assets/readme/go-lang.png) 

# Setup

- First, we need install **Ruby** for locally development.

- Then, we need install Ruby gems. Open terminal, run:

```
gem install bundler && bundle install
```

# Usage

Remember run `chmod -R 777 bin/file-name` if get permission denied.

- `./bin/setup` to setup project.

- `./bin/start` to starting your web server. The site will run at [http://127.0.0.1:4000](http://127.0.0.1:4000)

- `./bin/update` to pull new code from remote repo.

- `./bin/open` to open website on browser.

- `./bin/push` to push new code to new remote branch.

---

# Install Go (No need if you don't want commands)

If you want use commands to start|open|push|update for project.

- First, install Go lang: [https://golang.org/doc/install](https://golang.org/doc/install)

- Open Terminal, run:

```
chmod -R 777 bestcoder
```

Then, run:

```
cd your-project && pwd
```

you will get your project rooth path. Next, run:

```
rootDir=$(pwd) && echo "alias bestcoder='cd $rootDir && ./bestcoder'" >> ~/.bash_profile && source ~/.bash_profile
```

Now, you can use `bestcoder` command:

```
bestcoder 
  -task string
        Task required! Please choose one: start|open|push|update
```

- Setup:  `bestcoder -task setup` .
- Start:  `bestcoder -task start` .
- Open:   `bestcoder -task open`  .
- Push:   `bestcoder -task push`  .
- Update: `bestcoder -task update`.
---
