#!/bin/bash
# vegecli 安装脚本
# 用法: ./install_vegecli.sh [--tags=TAGS]
# 示例: ./install_vegecli.sh --tags=wl
#       TAGS=debug ./install_vegecli.sh

set -e

# 获取脚本所在目录，进入 vegecli 包
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR/vegecli"

# 使用 vegecli 包中的 make install
CGO=1 TAGS="$TAGS" BUILD_DIR="$GOPATH/bin" OUT_NAME=vege make install

echo "vege 已安装到 $GOPATH/bin/vege"
