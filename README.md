# geo-make
自定义 geosite.dat 和 geoip.dat 文件，修改自 https://github.com/gamesofts/v2ray-custom-geo
## 安装
```shell
go install -v github.com/lihuu/geo-make@main
```
## 使用
1. 在当前目录下运行 `geo-make`
 - 修改/添加 sites 文件夹下的域名文件，运行 `geo-make geosite` 生成 geosite.dat
 - 修改/添加 ips 文件夹下的地址文件，运行 `geo-make geoip` 生成 geoip.dat
2. 指定目录
 指定文件输入目录和文件输出目录。默认文件输入目录为当前目录，文件输出目录为当前目录。如果不指定输出目录，则目录和输入目录一致
 `geo-make geosite --src xxx --out xxx`

## 推荐
- 国内域名
  https://raw.githubusercontent.com/felixonmars/dnsmasq-china-list/master/accelerated-domains.china.conf

- 国内ip
  https://github.com/17mon/china_ip_list/blob/master/china_ip_list.txt

- 其他
 https://github.com/Loyalsoldier/v2ray-rules-dat
 https://github.com/StevenBlack/hosts



