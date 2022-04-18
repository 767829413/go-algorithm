#基于Nginx镜像
FROM nginx:latest

#换国内的源，这里用网易源
RUN sed -i 's#http://deb.debian.org#https://mirrors.163.com#g' /etc/apt/sources.list && sed -i 's#http://security.debian.org#https://mirrors.163.com#g' /etc/apt/sources.list && rm -Rf /var/lib/apt/lists/* && apt-get update

#安装Git，挂载Githooks钩子，并将产物放置在var/www/hexo中
RUN apt-get install -y git && git init --bare ~/blogs.git && mkdir -p /var/www && echo "git --work-tree=/var/www --git-dir=/root/blogs.git checkout -f" >~/blogs.git/hooks/post-receive && chmod a+x ~/blogs.git/hooks/post-receive

#SSH初始化
RUN apt-get install -y openssh-server && ssh-keygen -t rsa -f /etc/ssh/ssh_host_rsa_key -y && ssh-keygen -t ecdsa -f /etc/ssh/ssh_host_ecdsa_key -y && ssh-keygen -t ed25519 -f /etc/ssh/ssh_host_ed25519_key -y

#安装Certbot
RUN apt-get install -y certbot && apt-get install -y python3-certbot-nginx

#初始化Nginx配置挂载，这里选择直接覆盖所有配置
ADD nginx.conf /etc/nginx/nginx.conf

#暴露22 SSH
#暴露80 HTTP
#暴露443 HTTPS
EXPOSE 22
EXPOSE 80
EXPOSE 443