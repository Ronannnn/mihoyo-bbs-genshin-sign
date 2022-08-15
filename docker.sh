docker build -t ronannnn/mihoyo-bbs-genshin-sign:0.1.0 -f Dockerfile .
docker push ronannnn/mihoyo-bbs-genshin-sign:0.1.0
# docker save -o image-go.tar ronannnn/jrs-go:1.0
# scp image-go.tar liu_c5@10.70.9.111:~/image-tar/
# scp image-go.tar liu_c5@10.70.9.112:~/image-tar/

# mysql -h 127.0.0.1 -P 3306 -u root -p****** --default-character-set=utf8 jrs < /Users/ronan/Downloads/backup_2022-04-11T20:00:21Z.sql
# mysqldump -h 127.0.0.1 -P 3306 -u root -proot --ssl-mode=DISABLED --skip-add-locks --column-statistics=0 jrs >/Users/ronan/Desktop/jrs_2022-04-12_092007.sql
