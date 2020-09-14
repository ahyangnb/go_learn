查看开放端口
netstat -nupl (UDP类型的端口)
netstat -ntpl (TCP类型的端口)

a 表示所有
n 表示不查询dns
t 表示tcp协议
u 表示udp协议
p 表示查询占用的程序
l 表示查询正在监听的程序

netstat -ntpl | grep 3306 //这个表示查找处于监听状态的，端口号为3306的进程

[root@home ~]# netstat -ntpl | grep 3306
tcp        0      0 0.0.0.0:3306                0.0.0.0:*                   LISTEN      25302/mysqld   
1
2
[root@home ~]# netstat -nupl
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address               Foreign Address             State       PID/Program name   
udp        0      0 172.18.1.143:123            0.0.0.0:*                               1526/ntpd           
udp        0      0 127.0.0.1:123               0.0.0.0:*                               1526/ntpd           
udp        0      0 0.0.0.0:123                 0.0.0.0:*                               1526/ntpd     
1
2
3
4
5
6
[root@home ~]# netstat -ntpl
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address               Foreign Address             State       PID/Program name   
tcp        0      0 0.0.0.0:6369                0.0.0.0:*                   LISTEN      3284/beam.smp       
tcp        0      0 0.0.0.0:18083               0.0.0.0:*                   LISTEN      3284/beam.smp       
tcp        0      0 127.0.0.1:8743              0.0.0.0:*                   LISTEN      26488/java          
tcp        0      0 127.0.0.1:25                0.0.0.0:*                   LISTEN      19651/sendmail      
tcp        0      0 0.0.0.0:443                 0.0.0.0:*                   LISTEN      24538/nginx       


# Linux 如何开放端口和关闭端口
一、查看哪些端口被打开 netstat -anp
二、关闭端口号:

iptables -A OUTPUT -p tcp --dport 端口号-j DROP
1
三、打开端口号：

iptables -A INPUT -ptcp --dport  端口号-j ACCEPT
1
四、保存设置

service iptables save
1
五、以下是linux打开端口命令的使用方法。
　　nc -lp 23 &(打开23端口，即telnet)
　　netstat -an | grep 23 (查看是否打开23端口)
六、linux打开端口命令每一个打开的端口，都需要有相应的监听程序才可以

设置权限：chmod 777 main