
. .\wash-host.ps1


geo-make geosite 

# 获取ip地址
# scoop install wget
wget https://raw.githubusercontent.com/StevenBlack/hosts/master/alternates/fakenews-gambling-porn-only/hosts --no-check-certificate -O porn.txt
Clean-File -filePath "porn.txt" -extraStrings @("t66y.com", "wnacg.com", "hmba.top")


if (Test-Path ./sites/porn){
    rm ./sites/porn
}

mv ./porn.txt ./sites/porn
geo-make geosite 
