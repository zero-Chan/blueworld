## 安装lvs
> sudo apt-get install ipvsadm

## 部署规划
```
LVS Server 192.168.26.133  (VIP)：192.168.26.200

LVS负载均衡服务器
TOMCAT Server1 192.168.26.134
TOMCAT Server2 192.168.26.135

在192.168.26.133 上执行 

sudo ifconfig eth0:0 192.168.26.200 netmask 255.255.255.255 broadcast 192.168.26.200
sudo route add -host 192.168.26.200 dev eth0:0

ipvsadm -C 
ipvsadm -A -t 192.168.26.200:8080  -s wrr
ipvsadm -a -t 192.168.26.200:8080 -r  192.168.26.134:8080  -m  -w  1
ipvsadm -a -t 192.168.26.200:8080 -r  192.168.26.135:8080  -m  -w 1
service ipvsadm save
cat    /proc/sys/net/ipv4/ip_forward
echo 1 >  /proc/sys/net/ipv4/ip_forward 
service iptables stop

在134 135 web服务器上分别执行

sudo ipvsadm -a -t 192.168.26.200:8080 -r 192.168.26.134:8080 -g
sudo ipvsadm -a -t 192.168.26.200:8080 -r 192.168.26.135:8080 -g
```


## 测试：

在133上执行

curl  http://192.168.26.200:8080/examples

## keepalived
> https://blog.51cto.com/yangrong/1575909
