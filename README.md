# userservice [golang - gRPC]

### Build-Steps:

Install the dependencies and devDependencies and start the server.

```sh
$ make run
$ chmod +x scripts/client_app_test.sh
$ # script to execute some api calls using client
$ ./client_app_test
```

### Command to run test cases:

```sh
$ make test
```

### Command to run test client:

```sh
$ # change in scripts/client_app_test.sh:4
$ # export USER_IDS='1,2,3,4,5,6,7,8,9,10'
$ # for test with diffrent users
$ # for test client you must have go installed 
$ make testclient
```

### Use package in your project:

```sh
make
docker
go 1.13.5
```