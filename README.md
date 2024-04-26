# Creating GO CLI and Performing TestCases

### To initialize and Run GO files

```md
$ go mod init hellotesting
```
```md
$ mkdir hellotesting

$ cd .\hellotesting\

```
```md
$ go mod init 'Users/FullStack/source/repos/hellotesting'

```
## Installing Cobra-Cli

```md
$ go install github.com/spf13/cobra-cli@latest

```
#### Check if cobra is installed properly or not
```md
$ cobra-cli

```
### Create a new Cobra-based CLI application

for help '-h'

```md
$  cobra-cli init -h

```

```md
$ cobra-cli init hellotesting

```
Replace ***hellotesting*** with the name of your Go package.

#### Add commands to your CLI:
```md
$ cobra-cli add <command-name>

```
Replace **_<command-name>_**  with the name of the command you want to add.

#### Build and install your CLI:
```md
$ go install

```

### For named cli command parameters 

```md
$ hellotesting hello --name=zuhaib

```
```md
$ hellotesting hello

```

### Running GO program

```md
$ go run .\main.go

```

### To run test cases written, Use below command
```md
$ go test
$ go test ./cmd/...
```