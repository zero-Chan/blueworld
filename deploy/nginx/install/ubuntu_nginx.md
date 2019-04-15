# 安装参考
[如何在Ubuntu 16.04上安装Nginx](https://www.digitalocean.com/community/tutorials/how-to-install-nginx-on-ubuntu-16-04)

# 运行
1. 启动nginx(绝对路径)
> nginx -c /home/cza/workspace/cza-private/zero-Chan/blueworld/deploy/nginx/config/nginx.conf
2. 检查配置文件是否正确
> nginx -t -c /home/cza/workspace/cza-private/zero-Chan/blueworld/deploy/nginx/config/nginx.conf
3. 平滑重启nginx
> nginx -s reload
