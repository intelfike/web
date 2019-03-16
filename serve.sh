# イメージ作成
sudo docker build -t mp_web .

# コンテナを再作成
sudo docker rm -f mp_web
sudo docker run --name mp_web -d -p 8888:80 mp_web ./web -http :80