# servogo ![Go](https://github.com/NickHerrig/servogo/workflows/Go/badge.svg)
a thin cli/api for controlling servo motors over serial protocol written in go

## Usage
- Pull code and build with:
```shell
go build
```

- Sent env vars
```shell
SERVO_USB_PORT={serial-port}
```

```shell
./servogo --id {servo-id} --command {command} --data {optional-data}
```

## Implemented DMM Commands (functions.go)


| Command         |  Data Range              |  Description                              |
|:---------------:|:------------------------:|:-----------------------------------------:|
|    stop         |  None                    | stops the motor                           | 
|    forwards     |  None                    | sends the motor fowards                   | 
|    backwards    |  None                    | sends the motor backwards                 | 
|    send-to      | -134217728 ~ 1342177270  | sends the motor to specific position      |
|    position     |  None                    | returns the motors specific position      | 
|    set-speed    |  0 ~ 127                 | sets the motors speed                     | 
|    read-speed   |  None                    | reads the motors speed                    |
|    status       |  None                    | reads the motors status (alarm, motion)   |


## Testing 
Testing runs on every commit to master via github actions

To run tests in the current directory type
```shell
go test -v ./...
```

To run test coverage 
```shell
go test -cover ./...
```

To run test coveraged in CLI witho
```shell
go test -coverprofile=/tmp/profile.out ./...
go tool cover -func=/tmp/profile.out
```
