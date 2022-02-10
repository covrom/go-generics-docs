# Инсталляция gotip и gopls

```
go install golang.org/dl/gotip@latest
gotip download
mkdir /tmp/gopls && cd "$_"
gotip mod init gopls-unstable
gotip get golang.org/x/tools/gopls@master golang.org/x/tools@master
gotip install golang.org/x/tools/gopls
```

Configure VSCode: set "go.goroot" in settigs.json to the path of installed gotip (in my machine this path is ~/sdk/gotip).

Если лень устанавливать, то есть и онлайн:
https://go2goplay.golang.org/

UPD:
https://go.dev/blog/go1.18beta2