#!/bin/bash
# vegecli 安装脚本
# 用法: ./install_vegecli.sh [--tags=TAGS]
# 示例: ./install_vegecli.sh --tags=wl
#       TAGS=wl ./install_vegecli.sh

set -e

# 获取当前 OS：linux / darwin / windows
case "$(uname -s)" in
  Linux*)   OS=linux ;;
  Darwin*)  OS=darwin ;;
  MINGW*|MSYS*|CYGWIN*) OS=windows ;;
  *)        OS=linux ;;
esac

# 获取脚本所在目录，进入 vegecli 包
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR/vegecli"

# $GOPATH 兼容win路径格式
if [ "$OS" = "windows" ]; then
    GOPATH=$(cygpath -u "$GOPATH")
fi
# 使用 vegecli 包中的 make install
CGO=1 OS="$OS" TAGS="$TAGS" BUILD_DIR=$GOPATH/bin OUT_NAME=vege make install