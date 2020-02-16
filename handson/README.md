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
```

コンテナ作成
```
$docker-compose up -d
```

コンテナに入る
```
$docker exec -it chat-app sh
```

CloudSQL接続環境設定
```
$wget https://dl.google.com/cloudsql/cloud_sql_proxy.linux.amd64 -O cloud_sql_proxy
$chmod +x cloud_sql_proxy

$ $($GOPATH/src/api/cloud_sql_proxy -dir=/cloudsql -instances={}:{リージョン}:{DB} -credential_file=$GOPATH/src/api/{クレデンシャル}) &
```

Cloud Run
```
docker build -t gcr.io/sendaigo/goapp:6 -f step3/Dockerfile .
docker push gcr.io/sendaigo/goapp:5
```

GAE
```
go mod init main
gcloud app deploy app.yaml 
gcloud app browse
```
いろいろやってみよう

