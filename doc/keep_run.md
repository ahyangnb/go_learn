### 安装tmux

yum -y install tmux


### 创建一个后台会话

tmux new -s api -d


### 进入会话

tmux a -t api 


### 退出会话不是关闭

登到某一个会话后，依次按键ctrl-b + d，这样就会退化该会话，但不会关闭会话。
如果直接ctrl + d，就会在退出会话的通话也关闭了该会话！


### 关闭会话（销毁会话）

tmux kill-session -t api