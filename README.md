# backoff

A simple exponential backoff counter.

The algorithm is based on [google](https://cloud.google.com/storage/docs/exponential-backoff#example_algorithm).


[![Github issues](https://img.shields.io/github/issues/kskumgk63/dpfile)](https://github.com/kskumgk63/backoff/issues)
[![Github forks](https://img.shields.io/github/forks/kskumgk63/dpfile)](https://github.com/kskumgk63/backoff/network/members)
[![Github stars](https://img.shields.io/github/stars/kskumgk63/dpfile)](https://github.com/kskumgk63/backoff/stargazers)

![LOGO](https://raw.githubusercontent.com/egonelbre/gophers/master/sketch/superhero/flying.png)

## install

```
go get github.com/kskumgk63/backoff
```

## Examples

### no options

Repeat `alwaysErr()` at intervals of backoff time until the timeout occurs.

```go
func alwaysErr() error {
	return errors.New("internal server error")
}

func main() {
	cmd := backoff.NewCommander()
	if timeoutErr := cmd.Exec(alwaysErr); timeoutErr != nil {
		fmt.Println(timeoutErr)
	}
}
```

After 65s, this message is printed.
No messages during exponential backoff loop.

```
Ends the exponential backoff because of timeout
```

### change timeout

Change the timeout value to 10 seconds.

```go
func alwaysErr() error {
	return errors.New("internal server error")
}

func main() {
	cmd := backoff.NewCommander(
		backoff.Timeout(5 * time.Second),
	)
	if timeoutErr := cmd.Exec(alwaysErr); timeoutErr != nil {
		fmt.Println(timeoutErr)
	}
}

```

After 5s, this message is printed.
No messages during exponential backoff loop.

```
Ends the exponential backoff because of timeout
```

### debug mode on

if debug mode is on, prints errors when repeating `alwaysErr()`

```go
func alwaysErr() error {
	return errors.New("internal server error")
}

func main() {
	cmd := backoff.NewCommander(
		backoff.DebugModeOn(),
	)
	cmd.Exec(alwaysErr)
}

```

```
internal server error
waiting 2.020000s...
internal server error
waiting 4.564000s...
internal server error
waiting 8.586000s...
internal server error
waiting 16.869000s...
.
.
.
```

### change debug printer

```go
func alwaysErr() error {
	return errors.New("internal server error")
}

func main() {
	cmd := backoff.NewCommander(
		backoff.Timeout(10*time.Second),
		backoff.DebugModeOn(),
		backoff.DebugPrint(func(err error) {
			fmt.Printf("[ERR] %+v\n", err)
		}),
	)
	if timeoutErr := cmd.Exec(alwaysErr); timeoutErr != nil {
		fmt.Println(timeoutErr)
	}
}
```

You can change a debug printer!

```
[ERR] internal server error
waiting 2.088000s...
[ERR] internal server error
waiting 4.904000s...
[ERR] internal server error
waiting 8.212000s...
Ends the exponential backoff because of timeout
```

See more [options](https://github.com/kskumgk63/backoff/blob/main/option.go).
