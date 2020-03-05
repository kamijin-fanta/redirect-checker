# redirect-checker

一括で沢山のURLのリダイレクト設定状況を確認するツール

## install

リリースから適当に持ってくる。****の部分をリリースページのそれっぽいバイナリのURLに置き換える。

```
sudo curl -L ****** -o /usr/local/bin/redirect-checker
sudo chmod +x /usr/local/bin/redirect-checker
```

## usage

```
# 1行につき1つのURLを書いたファイルを用意する
$ echo https://google.com/ >> query.txt
$ echo https://about.twitter.com/ja/company/ >> query.txt
$ echo https://golang.org/ >> query.txt

# パイプでコマンドに渡すと、結果が出力される
$ cat query.txt |  ./redirect-checker
https://google.com/	301	https://www.google.com/	
https://about.twitter.com/ja/company/	301	https://about.twitter.com/ja/company.html	
https://golang.org/	200		
```
