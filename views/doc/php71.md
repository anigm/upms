# mac安装PHP71

[安装brew工具](https://brew.sh/index_zh-cn.html)

``` shell
$ brew install nginx
$ brew install php71
$ brew services start php71
```

    修改php-fpm.conf将listen的地址端口改为9001
    可以与php其他版本共存

nginx配置:
``` nginx
server {
    listen       8083;
    server_name  dev.local;
    root   /html;
    index  index.html index.htm index.php;

    autoindex_localtime on;
    access_log off;
     
    location ~ .*\.(php)?$ {
        fastcgi_pass 127.0.0.1:9001;
        fastcgi_index index.php;
        include fastcgi.conf;
    }

    location / {
        if (!-e $request_filename){
            rewrite ^(.*)$ /index.php?s=$1 last;
            break;
        }
    } 
}
```