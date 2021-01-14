# compare-big-files
This program checks two big files and compare them, avoiding to signal as error if two lines are presents in both files but in different order. 

## Installation
This program requires [jwangsadinata's go-multimap library](https://github.com/jwangsadinata/go-multimap), so run

```
go get -u github.com/jwangsadinata/go-multimap
```

## Usage

```
go run "compare_big_files.go" "file1" "file2"
```

Tested with two 10gb files, taked half an hour to confront
