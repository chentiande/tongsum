## 简单汇总程序

tongsum.exe   windows运行程序

tongsum    linux运行程序

db.conf   mysql数据库配置文件

*.cfg  每个报表汇总程序的配置文件，名字自定义



##  cfg配置说明

### 配置文件为xml格式，确保xml规范要求

1、可配置多个sqls，sqlname非必须配置，但为了后续维护进行描述该步骤执行含义

2、配置文件文件中可以配置参数，#starttime#为开始时间，#endtime#为结束时间，执行时会替换为程序参数

3、对于特殊符号的sql放在<![CDATA[sql]]>



## 程序参数

Usage of tongsum:
  -conf string
        汇总配置文件路径，默认为link.cfg (default "link.cfg")
  -endtime string
        结束时间，sql中可使用#endtime#代替
  -starttime string
        开始时间，sql中可使用#starttime#代替
  -sumlevel string
        汇总维度 0-小时，1-天，2-周，3-月，sql中可以使用#sumlevel#代替 (default "0")
  -v    show build version





