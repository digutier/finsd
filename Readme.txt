En caso de que las dependencias de GO no estén en orden (lo cual pasaba mucho en las VM)
Probar con los siguientes comandos

export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN