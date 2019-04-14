# mydocker

> 准备一个目录 比如:/nicktming 目录下准备一个busybox.tar
```
root@nicktming:/nicktming# pwd
/nicktming
root@nicktming:/nicktming# ls
busybox.tar
```

# build local
> -r 表示哪个目录下有busybox.tar (表示根据busybox镜像启动容器)

> ./mydocker run -it -r /nicktming /bin/sh

### code-4.2
```
go build .
// 准备镜像busybox.tar
-------------------------------terminal 01----------------------------------
root@nicktming:/nicktming# pwd
/nicktming
root@nicktming:/nicktming# ls
busybox.tar

// 创建容器
-------------------------------terminal 02----------------------------------
root@nicktming:~/go/src/github.com/nicktming/mydocker# pwd
/root/go/src/github.com/nicktming/mydocker
root@nicktming:~/go/src/github.com/nicktming/mydocker# go build .
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker run -it /bin/sh
2019/04/07 15:37:30 rootPath:
2019/04/07 15:37:30 rootPath is empaty, set cmd.Dir by default: /nicktming/busybox
2019/04/07 15:37:30 current path: /nicktming/mnt.
/ # ls
bin   dev   etc   home  proc  root  sys   tmp   usr   var
/ # mkdir nicktming01 && echo "testing01\n" > nicktming01/test01.txt
/ # ls
bin          dev          etc          home         nicktming01  proc         root         sys          tmp          usr          var
/ # cat nicktming01/test01.txt 
testing01\n


// 查看宿主机内容
-------------------------------terminal 01----------------------------------
root@nicktming:/nicktming# ls
busybox  busybox.tar  mnt  writerLayer
root@nicktming:/nicktming# cat mnt/nicktming01/test01.txt 
testing01\n
root@nicktming:/nicktming# cat writerLayer/nicktming01/test01.txt 
testing01\n
root@nicktming:/nicktming# df -h
Filesystem      Size  Used Avail Use% Mounted on
...
none             50G  2.7G   45G   6% /nicktming/mnt


// 退出容器
-------------------------------terminal 02----------------------------------
/ # exit
root@nicktming:~/go/src/github.com/nicktming/mydocker# 

// 查看宿主机内容 
-------------------------------terminal 01----------------------------------
root@nicktming:/nicktming# ls
busybox  busybox.tar
root@nicktming:/nicktming# df -h
Filesystem      Size  Used Avail Use% Mounted on
...
```

### code-4.3
```
------------------------------terminal 01---------------------------------
root@nicktming:/nicktming# pwd
/nicktming
root@nicktming:/nicktming# ls
busybox  busybox.tar  volume
root@nicktming:/nicktming# 

// 创建容器
------------------------------terminal 02---------------------------------
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker run -it -v /nicktming/volume01:/containerVolume01 -v /nicktming/volume02:/containerVolume02 /bin/sh
2019/04/07 23:18:14 volume:[/nicktming/volume01:/containerVolume01 /nicktming/volume02:/containerVolume02]
2019/04/07 23:18:14 rootPath:
2019/04/07 23:18:14 rootPath is empaty, set cmd.Dir by default: /nicktming/mnt
2019/04/07 23:18:14 current path: /nicktming/mnt.
/ # ls -l
total 52
drwxr-xr-x    2 root     root         12288 Feb 14 18:58 bin
drwxr-xr-x    4 root     root          4096 Apr  7 15:18 containerVolume01
drwxr-xr-x    4 root     root          4096 Apr  7 15:18 containerVolume02
drwxr-xr-x    4 root     root          4096 Mar 17 16:05 dev
drwxr-xr-x    3 root     root          4096 Mar 17 16:05 etc
drwxr-xr-x    2 nobody   nogroup       4096 Feb 14 18:58 home
dr-xr-xr-x   98 root     root             0 Apr  7 15:18 proc
drwx------    2 root     root          4096 Apr  7 15:18 root
drwxr-xr-x    2 root     root          4096 Mar 17 16:05 sys
drwxrwxrwt    2 root     root          4096 Feb 14 18:58 tmp
drwxr-xr-x    3 root     root          4096 Feb 14 18:58 usr
drwxr-xr-x    4 root     root          4096 Feb 14 18:58 var
/ # echo "containerVolume01" > containerVolume01/test001.txt
/ # echo "containerVolume02" > containerVolume02/test002.txt
/ # exit
root@nicktming:~/go/src/github.com/nicktming/mydocker# 

// 查看宿主机的内容
------------------------------terminal 01---------------------------------
root@nicktming:/nicktming# pwd
/nicktming
root@nicktming:/nicktming# ls
busybox  busybox.tar  volume  volume01  volume02
root@nicktming:/nicktming# tree volume01
volume01
`-- test001.txt

0 directories, 1 file
root@nicktming:/nicktming# cat volume01/test001.txt 
containerVolume01
root@nicktming:/nicktming# tree volume02/
volume02/
`-- test002.txt

0 directories, 1 file
root@nicktming:/nicktming# cat volume02/test002.txt 
containerVolume02
root@nicktming:/nicktming# 
```

### code-4.4
```
-------------------------------terminal 01--------------------------------
root@nicktming:/nicktming# pwd
/nicktming
root@nicktming:/nicktming# ls
busybox.tar  volume
root@nicktming:/nicktming# tree volume/
volume/
`-- test01.txt

0 directories, 1 file
root@nicktming:/nicktming# cat volume/test01.txt 
volume
volume again
root@nicktming:/nicktming# 

-------------------------------terminal 02--------------------------------
root@nicktming:~/go/src/github.com/nicktming/mydocker# git clone https://github.com/nicktming/mydocker.git
root@nicktming:~/go/src/github.com/nicktming/mydocker# git checkout code-4.4
root@nicktming:~/go/src/github.com/nicktming/mydocker# go build .
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker run -it -v /nicktming/volume:/containerVolume /bin/sh
2019/04/08 01:13:51 volume:[/nicktming/volume:/containerVolume]
2019/04/08 01:13:51 rootPath:
2019/04/08 01:13:51 rootPath is empaty, set cmd.Dir by default: /nicktming/mnt
2019/04/08 01:13:51 current path: /nicktming/mnt.
/ # ls -l
total 48
drwxr-xr-x    2 root     root         12288 Feb 14 18:58 bin
drwxr-xr-x    4 root     root          4096 Apr  7 17:13 containerVolume
drwxr-xr-x    4 root     root          4096 Mar 17 16:05 dev
drwxr-xr-x    3 root     root          4096 Mar 17 16:05 etc
drwxr-xr-x    2 nobody   nogroup       4096 Feb 14 18:58 home
dr-xr-xr-x  104 root     root             0 Apr  7 17:13 proc
drwx------    2 root     root          4096 Apr  7 17:13 root
drwxr-xr-x    2 root     root          4096 Mar 17 16:05 sys
drwxrwxrwt    2 root     root          4096 Feb 14 18:58 tmp
drwxr-xr-x    3 root     root          4096 Feb 14 18:58 usr
drwxr-xr-x    4 root     root          4096 Feb 14 18:58 var
/ # mkdir nicktming && echo "nicktming" > nicktming/test02.txt
/ #

-------------------------------terminal 03--------------------------------
root@nicktming:~/go/src/github.com/nicktming/mydocker# pwd
/root/go/src/github.com/nicktming/mydocker
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker commit image01
2019/04/08 01:15:01 imageTar:/nicktming/image01.tar
root@nicktming:~/go/src/github.com/nicktming/mydocker# 

// 退出容器
-------------------------------terminal 02--------------------------------
/ # exit
root@nicktming:~/go/src/github.com/nicktming/mydocker# 

// 查看生成的image
-------------------------------terminal 01--------------------------------
root@nicktming:/nicktming# pwd
/nicktming
root@nicktming:/nicktming# ls
busybox  busybox.tar  image01.tar  volume
root@nicktming:/nicktming# ls
busybox  busybox.tar  image01.tar  volume
root@nicktming:/nicktming# mkdir image01 && tar -xvf image01.tar -C image01
root@nicktming:/nicktming# cd image01/
root@nicktming:/nicktming/image01# ls -l
total 56
drwxr-xr-x 2 root   root    12288 Feb 15 02:58 bin
drwxr-xr-x 2 root   root     4096 Apr  8 01:13 containerVolume
drwxr-xr-x 4 root   root     4096 Mar 18 00:05 dev
drwxr-xr-x 3 root   root     4096 Mar 18 00:05 etc
drwxr-xr-x 2 nobody nogroup  4096 Feb 15 02:58 home
drwxr-xr-x 2 root   root     4096 Apr  8 01:14 nicktming
drwxr-xr-x 2 root   root     4096 Mar 18 00:05 proc
drwx------ 2 root   root     4096 Apr  8 01:13 root
drwxr-xr-x 2 root   root     4096 Mar 18 00:05 sys
drwxrwxrwt 2 root   root     4096 Feb 15 02:58 tmp
drwxr-xr-x 3 root   root     4096 Feb 15 02:58 usr
drwxr-xr-x 4 root   root     4096 Feb 15 02:58 var
root@nicktming:/nicktming/image01# 
root@nicktming:/nicktming/image01# cat containerVolume/test01.txt 
volume
volume again
root@nicktming:/nicktming/image01# cat nicktming/test02.txt 
nicktming
root@nicktming:/nicktming/image01# 
```

### code-5.2
```
// 前提条件
root@nicktming:/nicktming# pwd
/nicktming
root@nicktming:/nicktming# ls
busybox.tar

// 运行
root@nicktming:~/go/src/github.com/nicktming/mydocker# go build .
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker run -d /bin/top
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker run -d -name test /bin/top
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker ps
ID                     NAME                   PID         STATUS      COMMAND     CREATED
15549958821549242021   15549958821549242021   28396       running     /bin/top    2019-04-11 23:18:02
15549959221141642621   test                   28451       running     /bin/top    2019-04-11 23:18:42

```

### code-5.3
```
// 前提条件
root@nicktming:/nicktming# pwd
/nicktming
root@nicktming:/nicktming# ls
busybox.tar

// 运行
root@nicktming:~/go/src/github.com/nicktming/mydocker# go build .
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker run -d /bin/top
2019/04/12 21:35:51 rootPath:
2019/04/12 21:35:51 rootPath is empaty, set cmd.Dir by default: /nicktming/mnt
2019/04/12 21:35:51 mkdir /nicktming/writerLayer err:mkdir /nicktming/writerLayer: file exists
2019/04/12 21:35:51 containerId:15550761518017354161
2019/04/12 21:35:51 jsonInfo:{"pid":"11447","id":"15550761518017354161","name":"15550761518017354161","command":"/bin/top","createTime":"2019-04-12 21:35:51","status":"running"}
2019/04/12 21:35:51 mount -f /nicktming/mnt, err:exit status 1
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker ps
ID                     NAME                   PID         STATUS      COMMAND     CREATED
15550761518017354161   15550761518017354161   11447       running     /bin/top    2019-04-12 21:35:51
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker logs 15550761518017354161

Mem: 881648K used, 136164K free, 400K shrd, 110080K buff, 558300K cached
CPU:  0.0% usr  0.0% sys  0.0% nic  100% idle  0.0% io  0.0% irq  0.0% sirq
Load average: 0.02 0.04 0.05 1/102 3
  PID  PPID USER     STAT   VSZ %VSZ CPU %CPU COMMAND
    1     0 root     R     1280  0.1   0  0.0 /bin/top
```

### code-5.4
```
// 前提条件
root@nicktming:/nicktming# pwd
/nicktming
root@nicktming:/nicktming# ls
busybox.tar

// 运行
root@nicktming:~/go/src/github.com/nicktming/mydocker# git clone https://github.com/nicktming/mydocker.git
root@nicktming:~/go/src/github.com/nicktming/mydocker# git checkout code-5.4
root@nicktming:~/go/src/github.com/nicktming/mydocker# go build .
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker run -d /bin/top
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker ps
missing mydocker_pid env skip nsenter
ID                     NAME                   PID         STATUS      COMMAND     CREATED
15552033304408860601   15552033304408860601   21338       running     /bin/top    2019-04-14 08:55:30
root@nicktming:~/go/src/github.com/nicktming/mydocker# ps -ef | grep 21338
root     21338     1  0 08:55 pts/3    00:00:00 /bin/top
root     31541 29996  0 10:43 pts/4    00:00:00 grep --color=auto 21338
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker exec 15552033304408860601 /bin/sh
missing mydocker_pid env skip nsenter
2019/04/14 10:41:44 containerName:15552033304408860601,command:/bin/sh
got mydocker_pid=21338
got mydocker_cmd=/bin/sh
setns on ipc namespace succeeded
setns on uts namespace succeeded
setns on net namespace succeeded
setns on pid namespace succeeded
setns on mnt namespace succeeded
/ # ps -l
PID   USER     TIME  COMMAND
    1 root      0:00 /bin/top
    7 root      0:00 /bin/sh
    8 root      0:00 ps -l
/ # ps -ef
PID   USER     TIME  COMMAND
    1 root      0:00 /bin/top
    7 root      0:00 /bin/sh
    9 root      0:00 ps -ef
/ # exit
root@nicktming:~/go/src/github.com/nicktming/mydocker# 
```

# code-5.5
```
// 前提条件
root@nicktming:/nicktming# pwd
/nicktming
root@nicktming:/nicktming# ls
busybox.tar

// 运行
root@nicktming:~/go/src/github.com/nicktming/mydocker# git clone https://github.com/nicktming/mydocker.git
root@nicktming:~/go/src/github.com/nicktming/mydocker# git checkout code-5.5
root@nicktming:~/go/src/github.com/nicktming/mydocker# go build .
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker ps
ID          NAME        PID         STATUS      COMMAND     CREATED
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker run -d /bin/top
2019/04/14 16:27:07 rootPath:
2019/04/14 16:27:07 rootPath is empaty, set cmd.Dir by default: /nicktming/mnt
2019/04/14 16:27:07 mkdir /nicktming/writerLayer err:mkdir /nicktming/writerLayer: file exists
2019/04/14 16:27:07 containerId:15552304271404236701
2019/04/14 16:27:07 jsonInfo:{"pid":"30089","id":"15552304271404236701","name":"15552304271404236701","command":"/bin/top","createTime":"2019-04-14 16:27:07","status":"running"}
2019/04/14 16:27:07 mount -f /nicktming/mnt, err:exit status 1
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker ps
ID                     NAME                   PID         STATUS      COMMAND     CREATED
15552304271404236701   15552304271404236701   30089       running     /bin/top    2019-04-14 16:27:07
root@nicktming:~/go/src/github.com/nicktming/mydocker# ps -ef | grep top
root     30089     1  0 16:27 pts/3    00:00:00 /bin/top
root     30112 26623  0 16:27 pts/3    00:00:00 grep --color=auto top
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker stop 15552304271404236701
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker ps
ID                     NAME                   PID         STATUS      COMMAND     CREATED
15552304271404236701   15552304271404236701               stopped     /bin/top    2019-04-14 16:27:07
root@nicktming:~/go/src/github.com/nicktming/mydocker# 
```