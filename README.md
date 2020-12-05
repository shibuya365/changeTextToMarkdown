# ChangeTextToMarkdown
通常のテキストファイルをMarkdown形式のファイルへ変更する手助けをする

# DEMO
このような表形式のテキストファイルを、
```
タイトル 内容 備考
1 22 333
4444 55555 666666
あいう7 あいうえ アイウエ9
```
このようなMarkdownファイルへ変更する
---
|タイトル |内容     |備考     |
|---------|---------|---------|
|1        |22       |333      |
|4444     |55555    |666666   |
|あいう7  |あいうえ |アイウエ9|
---
 
# Features
 多彩なサブコマンドで必要な部分のみ変更可能
 
# Requirement 
* go
* go-runewidth
 
# Installation
goのインストール
```bash
go get github.com/mattn/go-runewidth
```

# Usage 
```bash
git clone https://github.com/shibuya365/changeTextToMarkdown.git
/changeTextToMarkdown
cd examples
go get github.com/mattn/go-runewidth
go build .
./main table your_file
```
 
# Note
 
必ずサブコマンド`table`の後に変更したいファイル`yourfile`を入れてコマンドを実行して下さい
## サブコマンド一覧

|コマンド        |内容              |
|:--------------:|:----------------:|
|`table`         |表形式に変換      |
|`check`         |□を`- [ ] `へ変換|
|`checked`       |✅を`- [x] `へ変換|
(↑もこのプログラムで作りました )
# Author
 
* shibuya365
* shibuya365days@gmail.com
 
# License
ライセンスを明示する
 
"ChangeTextToMarkdown" is under [MIT license](https://en.wikipedia.org/wiki/MIT_License).