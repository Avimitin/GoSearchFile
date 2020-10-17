# SearchFile 

## Info

This is a pratise project for familiar with go program.

## Todo

- [x] Search

- [x] Search with goroutine

## Idea

- Use a goroutine pool to limit to number of goroutine, each time the program want to make new goroutine, check the pool if the amount of current runnning goroutine is more than maximum limit or not. 

- Only if the amount of current running goroutines are not more than maximum, it can make a new goroutine for searching, else it will make recursion.

- Each time the program found target it will pass a boolean value to channel for add up the count.

- When the for loop done, if the method is the main method of the goroutine, it will send a boolean value to let the goroutine pool know one goroutine is over, and recover the pool.

## Result

In a file tree like below, target is to find all program named "test":

```text
.
├── 13131
│   ├── 34567
│   │   └── test
│   │       ├── 213
│   │       │   └── test
│   │       └── test
│   └── test
│       └── 131313
│           └── 312
│               └── 331
│                   └── test
└── go-test
    └── example
        ├── cmd
        │   └── exm
        │       └── main.go
        ├── go.mod
        └── internal
            ├── exm
            │   └── test
            ├── other
            │   ├── other1.go
            │   └── other2.go
            └── utils
                └── xxx.go

18 directories, 7 files
```

This program will perform result like below:

```text
===Using one thread===
22 matches
During 305.35099ms
===Using multi thread===
22 matches
During 158.611682ms
```

## Contribute

If you have interest in this program, you are welcome to make contribute.

Or if you just want to test this program, you can go to release pages and download the program.

> If you are not the linux-amd64 user, you can clone the project and install go to build this program. Exp like below.

- Exp

```bash
git clone 
cd SearchFile/cmd/Search
go build
```
