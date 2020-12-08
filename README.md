# go_learn
From the git repository of q1 learning go.

# go
test

# mysql
/usr/local/var/mysql

ERROR 2002 (HY000): 
Can't connect to local MySQL server through socket '/tmp/mysql.sock' (2)

### 解决
命令：

mysql.server start

然后在执行

mysql -uroot -p

密码直接使用回车，则可以进行登录，如果你想进行密码的设置，则执行

mysql_secure_installation

### 错误
```
ERROR 1045 (28000): Access denied for user 'root'@'localhost' (using password: NO)
```
这个是密码不正确，请输入正确的密码。