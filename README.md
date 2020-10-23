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

## Commands


| Command         |  Data Range              |  Description                      |
|:---------------:|:------------------------:|:---------------------------------:|
|    stop         |  None                    | stops the motor                   | 

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
