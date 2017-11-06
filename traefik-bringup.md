root@elastic-stack:/vagrant/traefik# cat README 
/vagrant/traefik
run 'docker-compose up -d'
cd test
docker-compose up -d
docker-compose scale whoami=2

To bring down:
Do in reverse order
cd test
docker-compose down
cd ..
docker-compose-down


root@elastic-stack:/vagrant/traefik# curl -H Host:whoami.docker.localhost http://127.0.0.1:81
Hostname: bbc586de5601
IP: 127.0.0.1
IP: 172.18.0.3
GET / HTTP/1.1
Host: whoami.docker.localhost
User-Agent: curl/7.52.1
Accept: */*
Accept-Encoding: gzip
X-Forwarded-For: 172.18.0.1
X-Forwarded-Host: whoami.docker.localhost
X-Forwarded-Port: 80
X-Forwarded-Proto: http
X-Forwarded-Server: 80370780a17f
X-Real-Ip: 172.18.0.1


Docker Commands
docker ps
docker logs --follow traefik_proxy_1

root@elastic-stack:/vagrant/traefik# docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                                        NAMES
ceb200e55c77        emilevauge/whoami   "/whoamI"                38 minutes ago      Up 38 minutes       80/tcp                                       test_whoami_2
bbc586de5601        emilevauge/whoami   "/whoamI"                38 minutes ago      Up 38 minutes       80/tcp                                       test_whoami_1
80370780a17f        traefik             "/traefik --web --..."   41 minutes ago      Up 40 minutes       0.0.0.0:81->80/tcp, 0.0.0.0:8081->8080/tcp   traefik_proxy_1
950d4dd59e4b        redis               "docker-entrypoint..."   2 weeks ago         Up 4 days           0.0.0.0:6379->6379/tcp                       myredis




