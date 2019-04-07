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

