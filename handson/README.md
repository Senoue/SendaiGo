# Sendai.go x GDG Cloud Sendai ハンズオン 

初歩的なコード入門〜

```
$git clone https://github.com/Senoue/SendaiGo.git
```

## Step 1
簡単Webサービスをつくろう
- Goの基本

```
$cd handson/step1/
$go run main.go
```

## Step 2
Go CRUD
- データベースからデータを取得
- packageの管理
- レシーバ

```
$cd handson/step2
$go get github.com/go-sql-driver/mysql
$go run mmain.go
```

## Step 3
組み合わせて、GCPにデプロイしてみる
- 組み合わせれば、チャットぽい

```
$cd handson/step3
$gcloud app deploy app.yaml
```

いろいろやってみよう

