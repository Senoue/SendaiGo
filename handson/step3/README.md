
docker exec -it 98456b871b8b sh
wget https://dl.google.com/cloudsql/cloud_sql_proxy.linux.amd64 -O cloud_sql_proxy
chmod +x cloud_sql_proxy
      $($GOPATH/src/api/cloud_sql_proxy -dir=/cloudsql -instances=sendaigo:us-central1:sendaigo -credential_file=$GOPATH/src/api/sendaigo-sql.json) &


      docker build -t gcr.io/sendaigo/goapp:6 -f step3/Dockerfile .


      docker push gcr.io/sendaigo/goapp:5