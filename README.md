# go-cli-memo
## 概要
わからない単語をメモ。

わかったら意味と一緒にメモ。

をするCLI。

## 使い方

1. CLでわからない単語を入力すると、unknown(table)に保存
> unknown graphql
2. メモしておいたわからない単語を取得
> unknown
3. わかったら単語と意味を入力
> know GraphQL hogehoge 
4. わかっていることを取得
> know
5. わかっている単語の意味を取得
> know GraphQL

## 不明点

.envを参照できないため、以下のようにalias設定して、パスによらずCLIを実行できるようにしている。.envを参照する方法が知りたい。
> alias unknown='(cd /Users/xxx/go/bin/; ./go-cli-memo unknown)'



