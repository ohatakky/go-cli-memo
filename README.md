# go-cli-memo
## Overview
我流のメモをCLIで。

わからない単語をメモ。

後日わかったら意味と一緒にメモ。

## Usage

1. わからない単語をメモ
> $ unknown りんご
2. メモしておいたわからない単語を取得
> $ unknown

> \> りんご
3. わかったら単語と意味をメモ
> $ known りんご 赤い果物 
4. メモしておいたわかっている単語と意味を取得
> $ known

> \> りんご 赤い果物

5. 更新
> $ unknown --update りんご みかん -u

> $ known --update りんご 赤くて甘い果物 -u

6. 削除
> $ unknown --delete りんご -d

> $ known --delete りんご -d
