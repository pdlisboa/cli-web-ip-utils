Cli web ip utils - Use this to get info about Ip and servers in web

USAGE:
   main [global options] command [command options] [arguments...]

COMMANDS:
   ip       
   server   
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show hel


EXAMPLES:

Get servers by host
```bash
go run main.go server --host amazon.com.br 
```


Get servers ip`s by host
```bash
go run main.go ip --host amazon.com.br 
```


More:

You can build the project to user without go literaly

build:
```bash
go build
```

use:
```bash
./cli-web-ip-utils server --host amazon.com.br 
```