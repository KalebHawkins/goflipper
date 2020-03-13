# Flipper

Flipper is a text based coin flip simulator written in Go. 

# Git it (to play with)

```
git clone https://github.com/kalebhawkins/goflipper.git
```

# Go get it (to code with)
```go
go get github.com/kalebhawkins/goflipper/flipper
```

# Use it

## Building 
```go
go build flip.go
```

## Running
```go
./flip

// Example output:
// #➤ [flipper] git:(master) ✗ go run flip.go 
// The coin has been flipped 100 times.
// Heads: true
// Tails: false
// Heads Count: 52
// TailsCount: 49


// Other things you can do
./flip -count 15

./flip -count 15 -verbose
```

## Getting Help
```
./flip -h

// Example output: 
// -count int
//        number of times to flip the coin (default 100)
// -verbose
//        displays how many times the coin rotates during each flip
```

# Code with it

```go
coin := Coin{true, false, 0, 0} // heads, tails, heads count, tails count

_ = coin.Flip()                // Flip() returns the number of coin rotations during coin flip.

fmt.Println(coin)               // Print a coins stats.

// Example output:
// #➤ [flipper] git:(master) ✗ go run flip.go 
// The coin has been flipped 100 times.
// Heads: true
// Tails: false
// Heads Count: 52
// TailsCount: 49
```