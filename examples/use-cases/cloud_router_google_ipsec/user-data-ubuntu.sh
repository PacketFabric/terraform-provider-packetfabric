#!/usr/bin/env bash

user="ubuntu"

# install nginx
sudo apt-get update
sudo apt-get -y install nginx iperf3 docker.io docker-compose
# make sure nginx is started
sudo service nginx start
# disable firewall
sudo ufw disable
# custo docker
sudo systemctl unmask docker.service
sudo systemctl unmask docker.socket
sudo systemctl start docker.service
sudo gpasswd -a $user docker

## Locust/Traffic Generator: https://locust.io
mkdir /home/$user/locust
cat <<EOT > /home/$user/locust/locustfile.py
import time
from locust import HttpUser, task, between
# https://docs.locust.io/en/stable/quickstart.html
class QuickstartUser(HttpUser):
    wait_time = between(5, 10)
    @task
    def index_page(self):
        self.client.get("/", verify=False)
EOT

sudo docker run --restart=unless-stopped --name=locust -dit -p 8089:8089 -v /home/$user/locust:/mnt/locust locustio/locust:latest -f /mnt/locust/locustfile.py --host http://10.2.1.x --web-auth demo:packetfabric
