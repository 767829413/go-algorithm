user nginx;
worker_processes auto;

error_log /var/log/nginx/error.log warn;
pid /var/run/nginx.pid;

events {
  worker_connections 1024;
}

http {
  include /etc/nginx/mime.types;
  default_type application/octet-stream;

  log_format main '$remote_addr - $remote_user [$time_local] "$request" '
  '$status $body_bytes_sent "$http_referer" '
  '"$http_user_agent" "$http_x_forwarded_for"';

  access_log /var/log/nginx/access.log main;

  sendfile on;
  #tcp_nopush     on;

  keepalive_timeout 65;

  #gzip  on;

  #include /etc/nginx/conf.d/*.conf;
  server {
    listen 80;
    server_name xn--20xa163vfzd.top; #记得改成自己的域名

    location / {
      root /var/www; #要和刚刚的Githooks目录对应上
      index index.html index.htm;
    }

    error_page 500 502 503 504 /50x.html;
    location = /50x.html {
      root /usr/share/nginx/html;
    }
  }
}