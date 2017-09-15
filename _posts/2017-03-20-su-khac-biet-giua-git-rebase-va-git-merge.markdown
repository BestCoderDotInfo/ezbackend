---
title: Sự Khác Biệt Giữa Git Rebase Và Git Merge
date: 2017-03-20 19:00
comments: true
external-url: 
categories: Git
keywords: git, github, git rebase, git merge
excerpt: Vấn đề tưởng chừng như đơn giản nhưng đôi lúc làm bạn cảm thấy rất confuse.
---
>Vấn đề tưởng chừng như đơn giản nhưng đôi lúc làm bạn cảm thấy rất confuse. Mình sẽ làm rõ vấn đề này thông qua ví dụ như sau:

1. Mình tạo một project có tên là: [https://github.com/vinhnglx/git-merge-rebase](https://github.com/vinhnglx/git-merge-rebase)

2. Clone và từ push README file lên master branch:

```bash
git clone git@github.com:vinhnglx/git-merge-rebase.git 
touch README.md
echo "Just a simple to clear git rebase and git merge" >> README.md
git add README.md
git commit -m 'Create README'
git push origin master
```

3. Tạo 2 branch mới từ master branch

```bash
git checkout master -b rebase-ast
git checkout master -b merge-ast
```

4. Checkout đến mỗi branch and tạo một vài thay đổi

```bash
git checkout rebase-ast
touch sample_key.txt
echo "Sample keys for project" >> sample_key.txt
git add sample_key.txt
git commit -m 'Create sample key file'
git push origin rebase-ast
```

```bash
git checkout merge-ast
touch hello.rb
echo "# Say hi to guys" >> hello.rb
git add hello.rb
git commit -m 'Create hello.rb file'
git push origin merge-ast
```

5. Lại checkout về master branch, update và push lên GitHub

```bash
git checkout master
echo "Hope to help you guys clear about git rebase and git merge" >> README.md
git add README.md
git commit -m 'Update README'
git push origin master
```

6. Master branch có sự thay đổi mới. Chúng ta cần phải sync sự thay đổi này về 2 branches: `merge-ast` và `rebase-ast`.

```bash
git checkout rebase-ast
git rebase master
```


```bash
git checkout merge-ast
git merge master
```

7. Yah, một khi đã đến bước này, các bạn sẽ thấy được vấn đề. Hãy compare 2 source-tree của 2 branches này

**merge-ast branch**

```bash
vinhnguyen@Vinh-Nguyen ~/Documents/projects/GitHub/git-merge-rebase (rebase-ast)$ git checkout merge-ast                                                             [ruby-2.2.0]
Switched to branch 'merge-ast'

vinhnguyen@Vinh-Nguyen ~/Documents/projects/GitHub/git-merge-rebase (merge-ast)$ git log --graph --pretty=oneline --abbrev-commit                                    [ruby-2.2.0]
*   4589613 Merge branch 'master' into merge-ast
|\
| * 56f0582 Update README
* | 20647a7 Create hello.rb file
|/
* dd67d6b Create README
```

**rebase-ast branch**

```bash
vinhnguyen@Vinh-Nguyen ~/Documents/projects/GitHub/git-merge-rebase (merge-ast)$ git checkout rebase-ast                                                             [ruby-2.2.0]
Switched to branch 'rebase-ast'

vinhnguyen@Vinh-Nguyen ~/Documents/projects/GitHub/git-merge-rebase (rebase-ast)$ git log --graph --pretty=oneline --abbrev-commit                                   [ruby-2.2.0]
* 680bacd Create sample key file
* 56f0582 Update README
* dd67d6b Create README
```

## Kết Luận:

- Commit trên cùng ở mỗi branch là commit mới nhất.
- Chú ý vào rebase-ast, mọi người sẽ thấy commit của rebase-ast nằm phía trên commit mới nhất của master. Còn ở merge-ast, mọi người sẽ thấy commit của master nằm phía trên commit mới nhất của merge-ast, ngoài ra một commit Merge branch cũng được tạo ra.
- Okay, lúc này nhiều bạn sẽ ồ, lạ vậy? Tại sao lại có 2 cách để sync các thay đổi mới nhất từmaster branch về. Thật phức tạp. Yeah, câu trả lời ở đây rất đơn giản. Nó phụ thuộc vào cách các bạn rewrite history như thế nào?.
- Bạn sử dụng git rebase nếu như bạn muốn các sự thay đổi thuộc về branch của bạn luôn luôn là mới nhất.
- Bạn sử dụng git merge nếu bạn muốn sắp xếp các commit theo mặc định.

Ở hầu hết các trường hợp, mình khuyến khích sử dụng git merge bởi vì git rebase có 1 số vấn đề sau:

- Không thể push các commit sau khi đã rebase ở local lên GitHub vì lịch sử của local và GitHub không giống nhau. Và chỉ có 1 cách duy nhất để push là sử dụng git push --force origin . Và đây có thể là nguyên nhân gây ra nhiều vấn đề, ví dụ như: Vì rebase nên một số change của master branch có thể sẽ không work fine trên branch của bạn.
- Conflict sẽ kinh khủng hơn. Ví dụ: Nếu như master branch có time line hơn branch của bạn 1 tháng. Hehe :trollface:. Lúc đó hãy rebase branch của bạn và sẽ thấy conflict đến mức nào 🌝 . Mình đã gặp rồi.

Khuyết điểm duy nhất của git merge là làm cho git commit list dài ra. Khó trace log. Nhất là trong 1 dự án dài hơi, việc nhìn lại log của vài tháng trước có thể sẽ là vấn đề với bạn.
Chính vì vậy, khi làm việc các bạn hãy cẩn thận trước khi quyết định là rebase
hay merge.
Bài viết được dịch từ [http://ruby-journal.com/what-is-the-difference-between-git-rebase-and-git-merge/]( http://ruby-journal.com/what-is-the-difference-between-git-rebase-and-git-merge/) và một số kinh nghiệm của mình khi làm việc với Git.
Rất mong mọi người đóng góp ý kiến.

Ví dụ: [https://github.com/vinhnglx/git-merge-rebase](https://github.com/vinhnglx/git-merge-rebase)

**Bài viết được lấy từ github : https://github.com/AsianTechInc/AST-ruby-code-review/issues/9**