#!/usr/bin/env bash
# Environment
opsys=windows
if [[ "$OSTYPE" == linux* ]]; then
  opsys=linux
elif [[ "$OSTYPE" == darwin* ]]; then
  opsys=darwin
fi
echo "OSTYPE: $OSTYPE"

if [ "$opsys" = 'darwin' ]; then
    # 获取硬件架构, arm64 对应 Mac M1，x86_64 对应 intel 架构
    UNAME_MACHINE=$(uname -m)
    echo "UNAME_MACHINE: $UNAME_MACHINE"
    if [ "$UNAME_MACHINE" = 'arm64' ]; then
        link='http://TODO/cli/kusionup/darwin-arm64/bin/kusionup'
    else
        link='http://TODO/cli/kusionup/darwin/bin/kusionup'
    fi
elif [ "$opsys" = 'windows' ]; then
    link='http://TODO/cli/kusionup/windows/bin/kusionup.exe'
else
    link='http://TODO/cli/kusionup/linux/bin/kusionup'
fi

echo "Installing latest kusionup..."

install_dir=~/.kusionup/bin
install_path=$install_dir/kusionup
if [ $opsys = 'windows' ]; then
    install_path=$PWD/kusionup
fi

rm $install_path &> /dev/null
mkdir -p $install_dir &> /dev/null
curl -sL $link -o $install_path
chmod +x $install_path

echo "Successfully installed latest kusionup!"
echo "kusionup installed to ${install_path}"

echo "Start execute kusionup init!"
$install_path init --skip-prompt $1
