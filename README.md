# vege9
* vege tools code helper pkg
using
```go 
    go get -u github.com/Godyu97/vege9@latest
```

* vegecli
```
# Wayland 剪贴板
CGO_ENABLED=1 go install -tags wl github.com/Godyu97/vege9/vegecli@latest

# 其他平台
CGO_ENABLED=1 go install github.com/Godyu97/vege9/vegecli@latest
//或
git clone https://github.com/Godyu97/vege9.git
cd vege9/ && 
TAGS=wl bash install_vegecli.sh //wayland环境 tags
bash install_vegecli.sh //其他平台
```