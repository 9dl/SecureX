# SecureX

SecureX is a powerful Golang library designed to enhance security measures and provide features to detect and prevent unauthorized access or malicious activities. SecureX is released under the GNU General Public License v3.0.

## License

SecureX is licensed under the GNU General Public License v3.0 (GPL-3.0). This means that you are free to use, modify, and distribute the library under the terms of this license. However, it is important to note the following:

- Skidding (repackaging and claiming as one's own) is strictly prohibited.
- Selling or monetizing SecureX as proprietary software is not allowed.
- Any modifications or derivative works based on SecureX must also be released under the GPL-3.0 license.

Please read the [LICENSE](LICENSE) file for the full text of the GNU General Public License v3.0.

## Installation

To install SecureX, you can use the following command:

```shell
go get github.com/StreamlineX/SecureX
```

## Implementation

```go
func main() {
	go AntiDebugRun()

	// Your main program logic goes here

	// To keep the program running indefinitely
	select {}
}
```

## Current Features

SecureX offers the following features:
- Searches by keywords: Detects specific process keywords.
- Finds renamed processes: Identifies altered process names.
- Dynamic library loading: Loads kernel32.dll dynamically.
- Retrieves function addresses: Gets IsDebuggerPresent function address.
- Handles termination errors: Manages process termination errors.
Stay tuned for more updates and enhancements to SecureX.
