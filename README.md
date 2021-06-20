# Task for one of the online go-school

## 1. Compiling project
There is MakeFile, so all of you need is just type `make`

## 2. Running tests
There is MakeFile, so all of you need is just type `make test`

## 3. Configuration application
You can change the port of the server and level of debug

To do that run application with  -config-path argument
where you should show path to the toml config file
```
bind_addr = ":8080"
log_level = "debug"
```
## 4. Logs
The application add into the log all your actions. 

## 4. Aditional dependencies 

1. To compile project on Windows you need install _**make**_ `choco install make`
2. To run tests on Windows you should install gcc: for example, [TDM](https://jmeubank.github.io/tdm-gcc/)
