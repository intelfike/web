# イメージ作成
sudo docker build -t mp_web .

# コンテナを再作成
sudo docker rm -f mp_web
sudo docker run --name mp_web -d -p 443:443 mp_web ./web
sudo docker exec -it mp_web certbot certonly
