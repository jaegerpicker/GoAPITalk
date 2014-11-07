#!/bin/bash
source activate_go_vm.sh
if [ -n $GOROOT ]; then
  ./.install_go_into_vm.sh --64
fi
go get github.com/go-sql-driver/mysql
go get -u github.com/jinzhu/gorm
go get github.com/go-martini/martini
go get github.com/martini-contrib/auth
go get github.com/codegangsta/gin
go get github.com/davecgh/go-spew/spew
go get gopkg.in/mgo.v2
go get github.com/codegangsta/martini-contrib/binding
