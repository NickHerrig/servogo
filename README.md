# servogo
a thin cli/api for controlling servo motors over serial protocol written in go

## Usage
- Pull code and build with:
```shell
go build
```

- Sent env vars
```shell
SERVO_DRIVE_ID={motor-id}
SERVO_USB_PORT={serial-port}
```

```shell
./servogo --command {command} --data {optional-data}
```

## Implemented Commands (functions.go)


| Command         |  Data Range              |  Description                              |
|:---------------:|:------------------------:|:-----------------------------------------:|
|    stop         |  None                    | stops the motor                           | 
|    forwards     |  None                    | sends the motor fowards                   | 
|    backwards    |  None                    | sends the motor backwards                 | 
|    send-to      | -134217728 ~ 1342177270  | sends the motor to specific position      |
|    position     |  None                    | returns the motors specific position      | 
|    set-speed    |  0 ~ 127                 | sets the motors speed                     | 
|    read-speed   |  None                    | reads the motors speed                    |


## Testing 
To run tests in the current directory type
```shell
go test .
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
