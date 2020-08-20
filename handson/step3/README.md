# Websocket API

### ローカル環境 ###

* Go 13以上をインスール

`go run app/router.go`

### ローカルからのデプロイ ###

 1 gcloudでプロジェクト設定
  
`gcloud config set project プロジェクトID `

 2 クラスターに接続

`gcloud container clusters get-credentials mirail --zone asia-northeast1-a`

 3 ビルドとイメージ作成

`gcloud builds submit --tag=gcr.io/jcgp-develop/chat-api:v0.2`
`gcloud builds submit --tag=gcr.io/jcgp-20181019/external-api:v1.3`

 4 デプロイ

 `kubectl apply -f manifests/deployment.yaml`

 参考：
 ・ service作成
 `kubectl apply -f manifests/service.yaml`

     


### Contribution guidelines ###

* Writing tests
* Code review
* Other guidelines

### Who do I talk to? ###

* Repo owner or admin
* Other community or team contact