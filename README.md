# mydocker

> 准备一个目录 比如:/nicktming 目录下准备一个busybox.tar
```
root@nicktming:/nicktming# pwd
/nicktming
root@nicktming:/nicktming# ls
busybox.tar
```

### build local
> -r 表示哪个目录下有busybox.tar (表示根据busybox镜像启动容器)

> ./mydocker run -it -r /nicktming /bin/sh
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




