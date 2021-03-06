# デシジョンテーブルを利用したテスト手法

## 2_1.1杯目のビールの価格

| |1|2|3|
|:----|:----|:----|:----|
|クーポン利用|Y|N|N|
|ハッピーアワー|-|Y|N|
|100円|X|-|-|
|290円|-|X|-|
|490円|-|-|X|

## 2_2.手数料

| |1|2|3|4|5|6|7|8|
|:----|:----|:----|:----|:----|:----|:----|:----|:----|
|平日|Y|Y|Y|Y|N|N|N|N|
|8:45~17:59|Y|Y|N|N|Y|N|Y|N|
|特別会員|Y|N|Y|N|Y|Y|N|N|
|手数料0円|X|X|X|-|X|X|-|-|
|手数料110円|-|-|-|X|-|-|X|X|

## 2_3.割引率

### 1 ロジックがわかっていない時のテストケース

| |1|2|3|4|5|6|7|8|
|:----|:----|:----|:----|:----|:----|:----|:----|:----|
|ワイシャツ含まれる|Y|Y|Y|Y|N|N|N|N|
|ネクタイ含まれる|Y|Y|N|N|Y|Y|N|N|
|7点以上買い上げ|Y|N|Y|N|Y|N|Y|N|
|12%割引|X|-|-|-|-|-|-|-|
|7%割引|-|-|X|-|X|-|X|-|
|5%割引|-|X|-|-|-|-|-|-|
|割引なし|-|-|-|X|-|X|-|X|

### 2 ロジックがわかった時のテストケース

| |1|2|3|4|5|6|
|:----|:----|:----|:----|:----|:----|:----|
|7点以上買い上げ|Y|Y|Y|N|N|N|
|ワイシャツ含まれる|Y|Y|N|Y|Y|N|
|ネクタイ含まれる|Y|N|-|Y|N|-|
|12%割引|X|-|-|-|-|-|
|7%割引|-|X|X|-|-|-|
|5%割引|-|-|-|X|-|-|
|0%割引|-|-|-|-|X|X|

## 2_4.魔王を倒すまでの道のり

### 1.簡単化前のデシジョンテーブル

| |1|2|3|4|5|6|7|8|9|10|11|12|13|14|15|16|
|:----|:----|:----|:----|:----|:----|:----|:----|:----|:----|:----|:----|:----|:----|:----|:----|:----|
|まほうのつえ|Y|Y|Y|Y|Y|Y|Y|Y|N|N|N|N|N|N|N|N|
|賢者が仲間|Y|Y|Y|Y|N|N|N|N|Y|Y|Y|Y|N|N|N|N|
|くらやみのかぎ|Y|Y|N|N|Y|Y|N|N|Y|Y|N|N|Y|Y|N|N|
|ひかりのつるぎ|Y|N|Y|N|Y|N|Y|N|Y|N|Y|N|Y|N|Y|N|
|部屋を発見できない|-|-|-|-|-|-|-|-|-|-|-|-|X|X|X|X|
|部屋を発見できる|X|X|X|X|X|X|X|X|X|X|X|X|-|-|-|-|
|部屋に入れる|X|X|-|-|X|X|-|-|X|X|-|-|-|-|-|-|
|魔王を倒せる|X|X|-|-|X|-|-|-|X|-|-|-|-|-|-|-|

### 2.簡単化後のデシジョンテーブル

まほうのつえ、賢者が仲間はどういうロジックかわからないので簡単化できない。

| |1|2|3|4|5|6|7|8|9|10|
|:----|:----|:----|:----|:----|:----|:----|:----|:----|:----|:----|
|まほうのつえ|Y|Y|Y|Y|Y|Y|N|N|N|N|
|賢者が仲間|Y|Y|Y|N|N|N|Y|Y|Y|N|
|くらやみのかぎ|Y|Y|N|Y|Y|N|Y|Y|N|-|
|ひかりのつるぎ|Y|N|-|Y|N|-|Y|N|-|-|
|部屋を発見できない|-|-|-|-|-|-|-|-|-|X|
|部屋を発見できる|X|X|X|X|X|X|X|X|X|-|
|部屋に入れる|X|X|-|X|X|X|X|X|-|-|
|魔王を倒せる|X|-|-|X|-|-|X|-|-|-|

### 3.テストケース

| |まほうのつえ|賢者が仲間|くやまのかぎ|ひかりのつるぎ|期待結果|
|:----|:----|:----|:----|:----|:----|
|1|あり|あり|あり|あり|魔王を倒す|
|2|あり|あり|あり|なし|部屋に入る|
|3|あり|あり|なし|-|部屋を発見できる|
|4|あり|なし|あり|あり|魔王を倒す|
|5|あり|なし|あり|なし|部屋に入る|
|6|あり|なし|なし|-|部屋を発見できる|
|7|なし|あり|あり|あり|魔王を倒す|
|8|なし|あり|あり|なし|部屋に入る|
|9|なし|あり|なし|-|部屋を発見できる|
|10|なし|なし|-|-|部屋を発見できない|

## 2_5.カレンダーの日付の色判定

### 制限指定

Y/Nで条件部分を指定すること。
列幅が狭くて済むが行数は大きくなる。

| |1|2|3|4|5|6|
|:----|:----|:----|:----|:----|:----|:----|
|祝日|Y|Y|Y|N|N|N|
|日曜日|Y|N|N|Y|N|N|
|土曜日|-|Y|N|-|Y|N|
|平日|-|-|Y|-|-|Y|
|赤|X|X|X|X|-|-|
|青|-|-|-|-|X|-|
|黒|-|-|-|-|-|X|

### 拡張指定

こういう書き方もあるとのこと

| |1|2|3|4|5|6|
|:----|:----|:----|:----|:----|:----|:----|
|祝日|Y|Y|Y|N|N|N|
|曜日|日曜日|土曜日|平日|日曜日|土曜日|平日|
|赤|X|X|X|X|-|-|
|青|-|-|-|-|X|-|
|黒|-|-|-|-|-|X|

## 2_6.プリペイドカードのチャージ特典

| |1|2|3|4|5|6|7|8|9|10|11|12|13|14|15|
|:----|:----|:----|:----|:----|:----|:----|:----|:----|:----|:----|:----|:----|:----|:----|:----|
|ランク|シルバー|シルバー|シルバー|シルバー|シルバー|ゴールド|ゴールド|ゴールド|ゴールド|ゴールド|ブラック|ブラック|ブラック|ブラック|ブラック|
|チャージ金額|3000|5000|5000|10000|10000|3000|5000|5000|10000|10000|3000|5000|5000|10000|10000|
|当選|-|Y|N|Y|N|-|Y|N|Y|N|-|Y|N|Y|N|
|ボーナス|1%|2%|2%|4%|4%|3%|5%|5%|7%|7%|5%|7%|7%|15%|15%|
|クーポン|-|X|-|X|-|-|X|-|X|-|-|X|-|X|-|

## 2_7.宅配ピザの特典

| |1|2|3|4|5|6|
|:----|:----|:----|:----|:----|:----|:----|
|合計金額1,500円以上|Y|Y|Y|N|N|N|
|宅配|Y|Y|N|Y|Y|N|
|クーポン|Y|N|-|Y|N|-|
|ポテトサービス|X|X|X|-|-|-|
|2枚目無料|-|-|X|-|-|X|
|20%割引|X|-|-|X|X|-|
|通常料金|-|X|-|-|-|X|

## 2_8.ネットスーパーのポイント倍率

### 15日がYを分割

| |1|2|3|4|5|6|7|8|
|:----|:----|:----|:----|:----|:----|:----|:----|:----|
|特別会員|Y|Y|Y|Y|N|N|N|N|
|カード会員|Y|Y|N|N|Y|Y|N|N|
|モバイル会員|Y|N|Y|N|Y|N|Y|N|
|15日|Y|Y|Y|Y|Y|Y|Y|Y|
|0倍|-|-|-|-|-|-|-|-|
|2倍|-|-|-|-|-|-|-|-|
|3倍|-|-|-|-|-|-|-|X|
|5倍|-|-|-|-|-|X|-|-|
|6倍|-|-|-|X|-|-|-|-|
|7倍|-|-|-|-|-|-|-|-|
|8倍|-|X|-|-|-|-|X|-|
|10倍|-|-|-|-|X|-|-|-|
|11倍|-|-|X|-|-|-|-|-|
|13倍|X|-|-|-|-|-|-|-|

### 15日がNを分割

| |1|2|3|4|5|6|7|8|
|:----|:----|:----|:----|:----|:----|:----|:----|:----|
|特別会員|Y|Y|Y|Y|N|N|N|N|
|カード会員|Y|Y|N|N|Y|Y|N|N|
|モバイル会員|Y|N|Y|N|Y|N|Y|N|
|15日|N|N|N|N|N|N|N|N|
|0倍|-|-|-|-|-|-|-|X|
|2倍|-|-|-|-|-|X|-|-|
|3倍|-|-|-|X|-|-|-|-|
|5倍|-|X|-|-|-|-|X|-|
|6倍|-|-|-|-|-|-|-|-|
|7倍|-|-|-|-|X|-|-|-|
|8倍|-|-|X|-|-|-|-|-|
|10倍|X|-|-|-|-|-|-|-|
|11倍|-|-|-|-|-|-|-|-|
|13倍|-|-|-|-|-|-|-|-|

## 2_9.駐車料金の無料時間

### 1.「最低でも60分の駐車料金無料時間を与える」の解釈の場合

| |1|2|3|4|5|6|
|:----|:----|:----|:----|:----|:----|:----|
|映画を見る|Y|Y|Y|N|N|N|
|お買い物合計2000円以上(5000円未満)|Y|N|N|Y|N|N|
|お買い物合計5000円以上|N|Y|N|N|Y|N|
|駐車料金無料時間0分|-|-|-|-|-|-|
|駐車料金無料時間60分|-|-|-|X|-|X|
|駐車料金無料時間120分|-|-|-|-|X|-|
|駐車料金無料時間180分|X|-|X|-|-|-|
|駐車料金無料時間240分|-|-|-|-|-|-|
|駐車料金無料時間300分|-|X|-|-|-|-|

### 2.「2,000円未満の買い物の場合駐車料金無料時間を60分追加」と解釈する場合

| |1|2|3|4|5|6|
|:----|:----|:----|:----|:----|:----|:----|
|映画を見る|Y|Y|Y|N|N|N|
|お買い物合計2000円以上(5000円未満)|N|Y|N|N|Y|N|
|お買い物合計5000円以上|N|N|Y|N|N|Y|
|駐車料金無料時間0分|-|-|-|-|-|-|
|駐車料金無料時間60分|-|-|-|X|X|-|
|駐車料金無料時間120分|-|-|-|-|-|X|
|駐車料金無料時間180分|-|-|-|-|-|-|
|駐車料金無料時間240分|X|X|-|-|-|-|
|駐車料金無料時間300分|-|-|X|-|-|-|

### 1と2の違い

1だと「最低でも60分は与える」となるため、映画をみて2,000円未満の買い物をした場合180分の無料時間となる

2だと「2,000円未満の買い物の場合60分の無料時間を追加」と判断するため映画を見て2,000円未満の買い物をした場合240分の無料時間となる

