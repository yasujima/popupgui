# popupgui

windows gui のサンプル

# intro

windows guiをGoで開発すること そのノウハウを貯めるためにこのProjectを立ち上げる。
GoでGUIを開発する場合のライブラリはいくつかあるようだが、Windowsを前提とした場合、Walkが簡単のためこれを使う。
TCLやGTK、GLなどを使う場合、環境構築のハードルが高い。環境自体はCUI環境、具体的にはUbuntu（WSL）で開発を行いたいため事前準備が簡単なものを採用する。

機能としては、

* 起動時に引数でいくつかのパラメータを渡し、これをGUI上に表示する。
* HTTPにてパラメータを渡し、返却された値もGUI上に表記する。
* Textエリアに記入した文字列を、HTTPにて外部にPostする。（サーバー側ではそのText情報をDBに登録）
* 文字列やサイズ、また上記HTTPの宛先などは、iniファイルにて外部設定化する

まずはこれらをターゲットする。
また、可能ならインストーラー作成までをターゲットとしたい。

golangについては、まだ学習中のため、そこも試行錯誤を行いながらすすめる。

# 参考情報

https://hake.hatenablog.com/entry/20150831/p1

https://blog.n-z.jp/blog/2017-02-19-golang-windows-gui.html

# 環境

ほとんど上記サイトを採用すればよいので、以下はメモ

```
$ go version
go version go1.13.4 linux/amd64

$ cat /etc/os-release
NAME="Ubuntu"
VERSION="18.04.2 LTS (Bionic Beaver)"
ID=ubuntu
ID_LIKE=debian
PRETTY_NAME="Ubuntu 18.04.2 LTS"
VERSION_ID="18.04"
HOME_URL="https://www.ubuntu.com/"
SUPPORT_URL="https://help.ubuntu.com/"
BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"
VERSION_CODENAME=bionic
UBUNTU_CODENAME=bionic
```
## 以下 Go関係

```
% GOOS=windows GOARCH=amd64 go get github.com/lxn/walk ＃こっちかも GOOS=linux GOARCH=amd64 go get github.com/lxn/walk
% go get github.com/akavel/rsrc
% GOOS=windows GOARCH=amd64 rsrc -manifest walk.manifest -o rsrc.syso
% GOOS=windows GOARCH=amd64 go build -ldflags="-H windowsgui"
```

