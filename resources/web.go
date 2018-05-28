package resources

const NGINX = `server {
    listen 80;
    root ${dir};
    index index.php;
    server_name ${url};
    client_max_body_size 20m;
    client_body_timeout 120s;
    location / {
        try_files /public/router.php =404;
        fastcgi_split_path_info ^(.+?\.php)(/.*)$;
        fastcgi_pass ${php};
        fastcgi_index router.php;
        fastcgi_param SCRIPT_FILENAME \$document_root\$fastcgi_script_name;
        include /etc/nginx/fastcgi_params;
    }
    location /assets {
        try_files /app/\$uri =404;
    }
}
#server {
#    listen 443;
#    root ${dir};
#    index index.php;
#
#    server_name ${url};
#
#    ssl on;
#    ssl_certificate     /etc/nginx/ssl/${url}.crt;
#    ssl_certificate_key /etc/nginx/ssl/${url}.key;
#
#    location / {
#        try_files /public/router.php =404;
#        fastcgi_split_path_info ^(.+?\.php)(/.*)$;
#        fastcgi_pass ${php};
#        fastcgi_index router.php;
#        fastcgi_param SCRIPT_FILENAME \$document_root\$fastcgi_script_name;
#        include /etc/nginx/fastcgi_params;
#    }
#
#    location /assets {
#        try_files /app/\$uri =404;
#    }
#}
`

const APACHE = `<VirtualHost *:80>
        ServerName ${url}
        DocumentRoot ${dir}/public
</VirtualHost>
<Directory '/srv/pufferpanel/public'>
        Require all granted
        AllowOverride All
</Directory>
`
