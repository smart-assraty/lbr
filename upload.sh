#!bin/bash
cd ~/go/mariadb/
sshpass -p "1q2w3e\$R" scp -r mariadb scripts/ templates/ root@192.168.1.17:/home/user
sshpass -p "1q2w3e\$R" ssh root@192.168.1.17 "cd /home/user && docker cp mariadb api:/root && docker cp templates/ api:/root/ && docker cp scripts/ apache2-container:/var/www/html/ && rm -rf mariadb templates/ scripts/"
exit 0
