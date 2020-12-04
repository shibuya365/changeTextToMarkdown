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
 
"hoge"のセールスポイントや差別化などを説明する
 
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
 
必ずサブコマンド`table`の後に`youfile`を入れてコマンドを実行して下さい
 
# Author
 
* shibuya365
* shibuya365days@gmail.com
 
# License
ライセンスを明示する
 
"ChangeTextToMarkdown" is under [MIT license](https://en.wikipedia.org/wiki/MIT_License).