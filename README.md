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

# code-5.6
```
// 前提条件
root@nicktming:/nicktming# pwd
/nicktming
root@nicktming:/nicktming# ls
busybox.tar

// 运行
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker run -d -name test /bin/top
2019/04/14 20:34:47 rootPath:
2019/04/14 20:34:47 rootPath is empaty, set cmd.Dir by default: /nicktming/mnt
2019/04/14 20:34:47 mkdir /nicktming/writerLayer err:mkdir /nicktming/writerLayer: file exists
2019/04/14 20:34:47 containerId:15552452873180321921
2019/04/14 20:34:47 jsonInfo:{"pid":"23197","id":"15552452873180321921","name":"test","command":"/bin/top","createTime":"2019-04-14 20:34:47","status":"running"}
2019/04/14 20:34:47 mount -f /nicktming/mnt, err:exit status 1
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker ps
ID                     NAME        PID         STATUS      COMMAND     CREATED
15552452873180321921   test        23197       running     /bin/top    2019-04-14 20:34:47
root@nicktming:~/go/src/github.com/nicktming/mydocker# ps -ef | grep top
root     23197     1  0 20:34 pts/1    00:00:00 /bin/top
root     23231 21810  0 20:34 pts/1    00:00:00 grep --color=auto top
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker ps
ID                     NAME        PID         STATUS      COMMAND     CREATED
15552452873180321921   test        23197       running     /bin/top    2019-04-14 20:34:47
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker stop test
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker ps
ID                     NAME        PID         STATUS      COMMAND     CREATED
15552452873180321921   test                    stopped     /bin/top    2019-04-14 20:34:47
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker rm test
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker ps
ID          NAME        PID         STATUS      COMMAND     CREATED
root@nicktming:~/go/src/github.com/nicktming/mydocker# 
```

# code-5.7
```
// 前提条件
-----------------------------------------------terminal 01---------------------------------------------------------
root@nicktming:/nicktming# pwd
/nicktming
root@nicktming:/nicktming# ls
busybox.tar
root@nicktming:/nicktming# 

-----------------------------------------------terminal 02---------------------------------------------------------
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker run -it -name container01 -v /nicktming/volume1:/containerVolume busybox /bin/sh
2019/04/16 23:14:37 rootPath is empaty, set rootPath: /nicktming
2019/04/16 23:14:37 current path: /nicktming/mnt/container01.
/ # ls -l
total 48
drwxr-xr-x    2 root     root         12288 Feb 14 18:58 bin
drwxr-xr-x    4 root     root          4096 Apr 16 15:14 containerVolume
drwxr-xr-x    4 root     root          4096 Mar 17 16:05 dev
drwxr-xr-x    3 root     root          4096 Mar 17 16:05 etc
drwxr-xr-x    2 nobody   nogroup       4096 Feb 14 18:58 home
dr-xr-xr-x  103 root     root             0 Apr 16 15:14 proc
drwx------    2 root     root          4096 Apr 16 15:14 root
drwxr-xr-x    2 root     root          4096 Mar 17 16:05 sys
drwxrwxrwt    2 root     root          4096 Feb 14 18:58 tmp
drwxr-xr-x    3 root     root          4096 Feb 14 18:58 usr
drwxr-xr-x    4 root     root          4096 Feb 14 18:58 var
/ # echo "container01:test01" > containerVolume/test01.txt
/ # exit
root@nicktming:~/go/src/github.com/nicktming/mydocker# 

-----------------------------------------------terminal 03---------------------------------------------------------
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker run -it -name container02 -v /nicktming/volume2:/containerVolume busybox /bin/sh
2019/04/16 23:15:41 rootPath is empaty, set rootPath: /nicktming
2019/04/16 23:15:41 current path: /nicktming/mnt/container02.
/ # ls -l
total 48
drwxr-xr-x    2 root     root         12288 Feb 14 18:58 bin
drwxr-xr-x    4 root     root          4096 Apr 16 15:15 containerVolume
drwxr-xr-x    4 root     root          4096 Mar 17 16:05 dev
drwxr-xr-x    3 root     root          4096 Mar 17 16:05 etc
drwxr-xr-x    2 nobody   nogroup       4096 Feb 14 18:58 home
dr-xr-xr-x  105 root     root             0 Apr 16 15:15 proc
drwx------    2 root     root          4096 Apr 16 15:15 root
drwxr-xr-x    2 root     root          4096 Mar 17 16:05 sys
drwxrwxrwt    2 root     root          4096 Feb 14 18:58 tmp
drwxr-xr-x    3 root     root          4096 Feb 14 18:58 usr
drwxr-xr-x    4 root     root          4096 Feb 14 18:58 var
/ # echo "container02:test01" > containerVolume/test01.txt
/ # exit
root@nicktming:~/go/src/github.com/nicktming/mydocker# 

-----------------------------------------------terminal 01---------------------------------------------------------
root@nicktming:/nicktming# pwd
/nicktming
root@nicktming:/nicktming# cat volume1/test01.txt 
container01:test01
root@nicktming:/nicktming# cat volume2/test01.txt 
container02:test01
root@nicktming:/nicktming#

```

# code-5.7.1
```
---------------------------------------terminal 01--------------------------------------------
root@nicktming:/nicktming# pwd
/nicktming
root@nicktming:/nicktming# ls
busybox.tar

// 启动两个容器 container01 container02
---------------------------------------terminal 02--------------------------------------------
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker run -d -name container01 -v /nicktming/from1:/to1 busybox /bin/top
2019/04/18 22:25:24 rootPath is empaty, set rootPath: /nicktming
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker run -d -name container02 -v /nicktming/from2:/to2 busybox /bin/top
2019/04/18 22:25:56 rootPath is empaty, set rootPath: /nicktming
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker ps
ID                     NAME          PID         STATUS      COMMAND     CREATED
15555975245549425111   container01   14158       running     /bin/top    2019-04-18 22:25:24
15555975563445863921   container02   14218       running     /bin/top    2019-04-18 22:25:56
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker exec container01 /bin/sh
2019/04/18 22:26:27 containerName:container01,command:/bin/sh
/ # echo -e "hello container1" >> /to1/test1.txt
/ # mkdir to1-1
/ # echo -e "hello cotainer1,to-1,test1" >> /to1-1/test1.txt
/ # exit
root@nicktming:~/go/src/github.com/nicktming/mydocker# 
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker commit container01 image1

// 查看宿主机内容
---------------------------------------terminal 01--------------------------------------------
root@nicktming:/nicktming# cat mnt/container01/to1-1/test1.txt 
hello cotainer1,to-1,test1
root@nicktming:/nicktming# cat mnt/container01/to1/test1.txt 
hello container1
root@nicktming:/nicktming# ls
busybox  busybox.tar  from1  from2  image1.tar  mnt  writerLayer
root@nicktming:/nicktming# 

// 删除容器container01 根据image1镜像启动容器container03 查看是否有to1,to1-1文件夹
---------------------------------------terminal 02--------------------------------------------
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker stop container01
2019/04/18 22:35:10 rootPath:/nicktming
2019/04/18 22:35:10 [/nicktming/from1:/to1]
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker rm container01
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker ps
ID                     NAME          PID         STATUS      COMMAND     CREATED
15555975563445863921   container02   14218       running     /bin/top    2019-04-18 22:25:56
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker run -d -name container03 image1 /bin/top
2019/04/18 22:37:50 rootPath is empaty, set rootPath: /nicktming
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker ps
ID                     NAME          PID         STATUS      COMMAND     CREATED
15555975563445863921   container02   14218       running     /bin/top    2019-04-18 22:25:56
15555982709688329991   container03   15433       running     /bin/top    2019-04-18 22:37:50
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker exec container03 /bin/sh
2019/04/18 22:38:08 containerName:container03,command:/bin/sh
/ # ls -l
total 52
drwxr-xr-x    2 root     root         12288 Feb 14 18:58 bin
drwxr-xr-x    4 root     root          4096 Mar 17 16:05 dev
drwxr-xr-x    3 root     root          4096 Mar 17 16:05 etc
drwxr-xr-x    2 nobody   nogroup       4096 Feb 14 18:58 home
dr-xr-xr-x   97 root     root             0 Apr 18 14:37 proc
drwx------    2 root     root          4096 Apr 18 14:38 root
drwxr-xr-x    2 root     root          4096 Mar 17 16:05 sys
drwxrwxrwt    2 root     root          4096 Feb 14 18:58 tmp
drwxr-xr-x    2 root     root          4096 Apr 18 14:26 to1
drwxr-xr-x    2 root     root          4096 Apr 18 14:27 to1-1
drwxr-xr-x    3 root     root          4096 Feb 14 18:58 usr
drwxr-xr-x    4 root     root          4096 Feb 14 18:58 var
// 文件夹存在 文件内容也存在 
/ # cat to1/test1.txt 
hello container1
/ # cat to1-1/test1.txt 
hello cotainer1,to-1,test1
/ # exit
root@nicktming:~/go/src/github.com/nicktming/mydocker# 

// 再次根据镜像image1启动 并且用宿主机中的from5映射到容器的/to1 
//根据aufs原理可知容器层的内容会覆盖镜像层的内容, 因此/to1/test1.txt的内容为hello container05
root@nicktming:~/go/src/github.com/nicktming/mydocker# mkdir -p /nicktming/from5 && echo "hello container05" > /nicktming/from5/test1.txt
root@nicktming:~/go/src/github.com/nicktming/mydocker# cat /nicktming/from5/test1.txt 
hello container05
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker run -it -name container05 -v /nicktming/from5:/to1 image1 /bin/sh
2019/04/18 22:45:06 rootPath is empaty, set rootPath: /nicktming
2019/04/18 22:45:06 current path: /nicktming/mnt/container05.
/ # ls -l 
total 52
drwxr-xr-x    2 root     root         12288 Feb 14 18:58 bin
drwxr-xr-x    4 root     root          4096 Mar 17 16:05 dev
drwxr-xr-x    3 root     root          4096 Mar 17 16:05 etc
drwxr-xr-x    2 nobody   nogroup       4096 Feb 14 18:58 home
dr-xr-xr-x   97 root     root             0 Apr 18 14:45 proc
drwx------    2 root     root          4096 Apr 18 14:45 root
drwxr-xr-x    2 root     root          4096 Mar 17 16:05 sys
drwxrwxrwt    2 root     root          4096 Feb 14 18:58 tmp
drwxr-xr-x    4 root     root          4096 Apr 18 14:45 to1
drwxr-xr-x    2 root     root          4096 Apr 18 14:27 to1-1
drwxr-xr-x    3 root     root          4096 Feb 14 18:58 usr
drwxr-xr-x    4 root     root          4096 Feb 14 18:58 var
/ # cat to1
to1-1/  to1/
/ # cat to1/test1.txt 
hello container05
```

### code-5.7.2

```
-------------------------------terminal 01----------------------------
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker run -d -name container01 busybox /bin/top
2019/04/19 15:19:56 rootPath is empaty, set rootPath: /nicktming
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker ps
ID                     NAME          PID         STATUS      COMMAND     CREATED
15556583962876437411   container01   18416       running     /bin/top    2019-04-19 15:19:56
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker exec container01 /bin/sh
2019/04/19 15:20:21 containerName:container01,command:/bin/sh
/ # ls -l
total 44
drwxr-xr-x    2 root     root         12288 Feb 14 18:58 bin
drwxr-xr-x    4 root     root          4096 Mar 17 16:05 dev
drwxr-xr-x    3 root     root          4096 Mar 17 16:05 etc
drwxr-xr-x    2 nobody   nogroup       4096 Feb 14 18:58 home
dr-xr-xr-x  103 root     root             0 Apr 19 07:19 proc
drwx------    2 root     root          4096 Apr 19 07:20 root
drwxr-xr-x    2 root     root          4096 Mar 17 16:05 sys
drwxrwxrwt    2 root     root          4096 Feb 14 18:58 tmp
drwxr-xr-x    3 root     root          4096 Feb 14 18:58 usr
drwxr-xr-x    4 root     root          4096 Feb 14 18:58 var

-------------------------------terminal 02----------------------------
root@nicktming:/nicktming# mkdir copy && echo "copy files" > copy/test01.txt
// 从宿主机copy文件到容器中
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker cp /nicktming/copy/test01.txt container01:/
2019/04/19 15:49:50 source:/nicktming/copy/test01.txt, destination:container01:/
2019/04/19 15:49:50 containerUrl:container01:/, hostUrl:/nicktming/copy/test01.txt, conatinerName:container01, containerPath:/
2019/04/19 15:49:50 containerPath:/nicktming/mnt/container01/, hostPath:/nicktming/copy/test01.txt
2019/04/19 15:49:50 from_container_to_host:false
2019/04/19 15:49:50 from /nicktming/copy/test01.txt to /nicktming/mnt/container01/
root@nicktming:~/go/src/github.com/nicktming/mydocker#
 
// 从容器中copy文件到宿主机
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker cp container01:/bin/top /root/go/src/github.com/nicktming/mydocker
2019/04/19 15:51:00 source:container01:/bin/top, destination:/root/go/src/github.com/nicktming/mydocker
2019/04/19 15:51:00 containerUrl:container01:/bin/top, hostUrl:/root/go/src/github.com/nicktming/mydocker, conatinerName:container01, containerPath:/bin/top
2019/04/19 15:51:00 containerPath:/nicktming/mnt/container01/bin/top, hostPath:/root/go/src/github.com/nicktming/mydocker
2019/04/19 15:51:00 from_container_to_host:true
2019/04/19 15:51:00 from /nicktming/mnt/container01/bin/top to /root/go/src/github.com/nicktming/mydocker
// 验证top命令是否copy到当前位置
root@nicktming:~/go/src/github.com/nicktming/mydocker# ls
cgroups  command  main.go  memory  mydocker  nsenter  pictures  README.md  test  top  urfave-cli-examples
root@nicktming:~/go/src/github.com/nicktming/mydocker# 

// 查看容器中是否有test1.txt文件
-------------------------------terminal 01----------------------------
/ # ls -l
total 48
drwxr-xr-x    2 root     root         12288 Feb 14 18:58 bin
drwxr-xr-x    4 root     root          4096 Mar 17 16:05 dev
drwxr-xr-x    3 root     root          4096 Mar 17 16:05 etc
drwxr-xr-x    2 nobody   nogroup       4096 Feb 14 18:58 home
dr-xr-xr-x  102 root     root             0 Apr 19 07:19 proc
drwx------    2 root     root          4096 Apr 19 07:20 root
drwxr-xr-x    2 root     root          4096 Mar 17 16:05 sys
-rw-r--r--    1 root     root            11 Apr 19 07:49 test01.txt
drwxrwxrwt    2 root     root          4096 Feb 14 18:58 tmp
drwxr-xr-x    3 root     root          4096 Feb 14 18:58 usr
drwxr-xr-x    4 root     root          4096 Feb 14 18:58 var
/ # cat test01.txt 
copy files
/ # exit
root@nicktming:~/go/src/github.com/nicktming/mydocker# 
```

### code-5.8
```
root@nicktming:/nicktming# pwd
/nicktming
root@nicktming:/nicktming# ls
busybox.tar


// -it 的形式没有问题
root@nicktming:~/go/src/github.com/nicktming/mydocker# go build .
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker run -it -e name=nicktming busybox /bin/sh
2019/04/19 22:44:30 rootPath is empaty, set rootPath: /nicktming
2019/04/19 22:44:30 current path: /nicktming/mnt/15556850702382348471.
/ # env | grep name
name=nicktming
/ # exit
root@nicktming:~/go/src/github.com/nicktming/mydocker# 

root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker run -d -name container01 -e name=nicktming busybox /bin/top
2019/04/19 23:03:28 rootPath is empaty, set rootPath: /nicktming
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker exec container01 /bin/sh
2019/04/19 23:03:43 containerName:container01,command:/bin/sh
/ # env | grep name
name=nicktming
/ # exit
root@nicktming:~/go/src/github.com/nicktming/mydocker# 
```

### code-6.2
```
root@nicktming:~/go/src/github.com/nicktming/mydocker# ip link
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN mode DEFAULT group default 
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP mode DEFAULT group default qlen 1000
    link/ether 52:54:00:58:9a:c3 brd ff:ff:ff:ff:ff:ff
3: docker0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN mode DEFAULT group default 
    link/ether 56:84:7a:fe:97:99 brd ff:ff:ff:ff:ff:ff
root@nicktming:~/go/src/github.com/nicktming/mydocker# 
root@nicktming:~/go/src/github.com/nicktming/mydocker# iptables -t nat -vnL POSTROUTING --line
Chain POSTROUTING (policy ACCEPT 2 packets, 120 bytes)
num   pkts bytes target     prot opt in     out     source               destination         
1       14   998 MASQUERADE  all  --  *      !docker0  172.17.0.0/16        0.0.0.0/0           
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker network create --driver bridge --subnet 192.168.10.1/24 testbridge01
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker network list
NAME           IpRange           Driver
testbridge01   192.168.10.1/24   bridge
root@nicktming:~/go/src/github.com/nicktming/mydocker# ifconfig testbridge01
testbridge01 Link encap:Ethernet  HWaddr 02:f4:c8:97:df:39  
          inet addr:192.168.10.1  Bcast:0.0.0.0  Mask:255.255.255.0
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:0 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0 
          RX bytes:0 (0.0 B)  TX bytes:0 (0.0 B)
root@nicktming:~/go/src/github.com/nicktming/mydocker# iptables -t nat -vnL POSTROUTING --line
Chain POSTROUTING (policy ACCEPT 2 packets, 135 bytes)
num   pkts bytes target     prot opt in     out     source               destination         
1       14   998 MASQUERADE  all  --  *      !docker0  172.17.0.0/16        0.0.0.0/0           
2        0     0 MASQUERADE  all  --  *      !testbridge01  192.168.10.0/24      0.0.0.0/0           
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker network remove testbridge01
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker network list
NAME        IpRange     Driver
```
> network implementation
```
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker run -it -p 80:80 -net testbridge01 busybox /bin/sh
2019/05/09 22:47:17 rootPath is empaty, set rootPath: /nicktming
2019/05/09 22:47:17 current path: /nicktming/mnt/15574132372127349531.
/ # ifconfig
cif-15574 Link encap:Ethernet  HWaddr 0A:A9:1E:91:0B:BD  
          inet addr:192.168.10.3  Bcast:0.0.0.0  Mask:255.255.255.0
          inet6 addr: fe80::8a9:1eff:fe91:bbd/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:0 errors:0 dropped:0 overruns:0 frame:0
          TX packets:6 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000 
          RX bytes:0 (0.0 B)  TX bytes:508 (508.0 B)

lo        Link encap:Local Loopback  
          inet addr:127.0.0.1  Mask:255.0.0.0
          inet6 addr: ::1/128 Scope:Host
          UP LOOPBACK RUNNING  MTU:65536  Metric:1
          RX packets:0 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0 
          RX bytes:0 (0.0 B)  TX bytes:0 (0.0 B)

/ # route -n
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         192.168.10.1    0.0.0.0         UG    0      0        0 cif-15574
192.168.10.0    0.0.0.0         255.255.255.0   U     0      0        0 cif-15574
/ # ping -c 1 192.168.10.3
PING 192.168.10.3 (192.168.10.3): 56 data bytes
64 bytes from 192.168.10.3: seq=0 ttl=64 time=0.042 ms

--- 192.168.10.3 ping statistics ---
1 packets transmitted, 1 packets received, 0% packet loss
round-trip min/avg/max = 0.042/0.042/0.042 ms
/ # ping -c 1 127.0.0.1
PING 127.0.0.1 (127.0.0.1): 56 data bytes
64 bytes from 127.0.0.1: seq=0 ttl=64 time=0.039 ms

--- 127.0.0.1 ping statistics ---
1 packets transmitted, 1 packets received, 0% packet loss
round-trip min/avg/max = 0.039/0.039/0.039 ms
/ # ping -c 1 192.168.10.1
PING 192.168.10.1 (192.168.10.1): 56 data bytes
64 bytes from 192.168.10.1: seq=0 ttl=64 time=0.067 ms

--- 192.168.10.1 ping statistics ---
1 packets transmitted, 1 packets received, 0% packet loss
round-trip min/avg/max = 0.067/0.067/0.067 ms
/ # ping -c 1 172.19.16.7
PING 172.19.16.7 (172.19.16.7): 56 data bytes
64 bytes from 172.19.16.7: seq=0 ttl=64 time=0.055 ms

--- 172.19.16.7 ping statistics ---
1 packets transmitted, 1 packets received, 0% packet loss
round-trip min/avg/max = 0.055/0.055/0.055 ms
/ # echo "nameserver 183.60.83.19" >> /etc/resolv.conf 
/ # echo "nameserver 183.60.82.98" >> /etc/resolv.conf 
/ # 
/ # cat /etc/resolv.conf 
nameserver 183.60.83.19
nameserver 183.60.82.98
/ # ping -c 1 www.baidu.com
PING www.baidu.com (119.63.197.139): 56 data bytes
64 bytes from 119.63.197.139: seq=0 ttl=51 time=53.213 ms

--- www.baidu.com ping statistics ---
1 packets transmitted, 1 packets received, 0% packet loss
round-trip min/avg/max = 53.213/53.213/53.213 ms
/ # nc -lp 80

```
> 在另外一台机器访问该容器的服务
```
telnet 150.109.72.88 80
Trying 150.109.72.88...
Connected to 150.109.72.88.
Escape character is '^]'.
hello world!
```
> 容器中可以看到
```
/ # nc -lp 80
hello world!
```
> 容器之间访问
```
root@nicktming:~/go/src/github.com/nicktming/mydocker# ./mydocker run -it -net testbridge01 busybox /bin/sh
2019/05/09 23:06:18 rootPath is empaty, set rootPath: /nicktming
2019/05/09 23:06:18 current path: /nicktming/mnt/7484077863.
/ # ifconfig
cif-74840 Link encap:Ethernet  HWaddr 0E:05:F7:07:54:73  
          inet addr:192.168.10.7  Bcast:0.0.0.0  Mask:255.255.255.0
          inet6 addr: fe80::c05:f7ff:fe07:5473/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:0 errors:0 dropped:0 overruns:0 frame:0
          TX packets:5 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000 
          RX bytes:0 (0.0 B)  TX bytes:418 (418.0 B)

lo        Link encap:Local Loopback  
          inet addr:127.0.0.1  Mask:255.0.0.0
          inet6 addr: ::1/128 Scope:Host
          UP LOOPBACK RUNNING  MTU:65536  Metric:1
          RX packets:0 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0 
          RX bytes:0 (0.0 B)  TX bytes:0 (0.0 B)

/ # ping -c 1 192.168.10.3
PING 192.168.10.3 (192.168.10.3): 56 data bytes
64 bytes from 192.168.10.3: seq=0 ttl=64 time=0.080 ms

--- 192.168.10.3 ping statistics ---
1 packets transmitted, 1 packets received, 0% packet loss
round-trip min/avg/max = 0.080/0.080/0.080 ms
```
> 宿主机访问容器
```
root@nicktming:~/go/src/github.com/nicktming/mydocker# ping -c 1 192.168.10.3
PING 192.168.10.3 (192.168.10.3) 56(84) bytes of data.
64 bytes from 192.168.10.3: icmp_seq=1 ttl=64 time=0.051 ms

--- 192.168.10.3 ping statistics ---
1 packets transmitted, 1 received, 0% packet loss, time 0ms
rtt min/avg/max/mdev = 0.051/0.051/0.051/0.000 ms
```
> Iptables查看
```
root@nicktming:~/go/src/github.com/nicktming/mydocker# iptables -t nat -vnL PREROUTING
Chain PREROUTING (policy ACCEPT 425 packets, 18063 bytes)
 pkts bytes target     prot opt in     out     source               destination         
1614K   62M DOCKER     all  --  *      *       0.0.0.0/0            0.0.0.0/0            ADDRTYPE match dst-type LOCAL
    2   104 DNAT       tcp  --  *      *       0.0.0.0/0            0.0.0.0/0            tcp dpt:80 to:192.168.10.3:80
root@nicktming:~/go/src/github.com/nicktming/mydocker# 
```