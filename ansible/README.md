# gosetup

#### 介绍
利用ansible自动安装配置golang开发环境

#### 安装
以 Ubuntu18.04 为例

- 安装ansible

```
apt install python3-pip
pip3 install ansible==2.6.12
```

- 准备

```
ssh-keygen -t rsa -b 2048 -N '' -f /root/.ssh/id_rsa > /dev/null
cat /root/.ssh/id_rsa.pub >> /root/.ssh/authorized_keys
host_if=$(ip route|grep default|cut -d' ' -f5)
host_ip=$(ip a|grep "$host_if$"|awk '{print $2}'|cut -d'/' -f1)
ssh-keyscan -t ecdsa -H "$host_ip" >> /root/.ssh/known_hosts
```

- 安装go

```
ansible-playbook roles/gosetup/gosetup.yml
```
