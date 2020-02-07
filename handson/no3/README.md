# MIRAIL 外部サービス接続用API

MIRAILAPIから外部サービスへ接続する際に接続サービス先でIP制限をかけている場合、GlobalIPの指定が必要となるがGAEは動的になるため
本API経由で接続する

### ローカル環境 ###

* Go 12以上をインスールしておく

`go run main.go`

### ローカルからのデプロイ ###

 1 gcloudでプロジェクト設定
  
`gcloud config set project プロジェクトID `

 2 クラスターに接続

`gcloud container clusters get-credentials standard-cluster-1 --zone us-central1-a`

 3 ビルドとイメージ作成

`gcloud builds submit --tag=gcr.io/goco-sendai2020/redis:v0.1`

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