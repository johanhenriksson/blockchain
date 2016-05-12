/**
 * Go Example Blockchain
 * Copyright (C) 2016 Johan Henriksson
 * johanhenriksson@live.com
 */
package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "strconv"
    "encoding/hex"
    "crypto/sha256"
)

// block difficulty
const prefix = "000000"

// block structure
type Block struct {
    Hash        string
    Previous    string
    Content     string
    Random      int
}

// print block
func (b *Block) Print() {
    fmt.Println("------------------------------------------")
    fmt.Println("Block found:   ", b.Hash)
    fmt.Println("Previous block:", b.Previous)
    fmt.Println("Random:        ", b.Random)
    fmt.Println("Content:       ", b.Content)
    fmt.Println("------------------------------------------")
}

// entry point
func main() {
    reader := bufio.NewReader(os.Stdin)
    block  := &Block { Hash: "genesis", }

    fmt.Println("Example text blockchain. Difficulty prefix:", prefix)
    fmt.Println("Enter some data:")
    for {
        // read line from standard input
        fmt.Print("> ")
        line, _ := reader.ReadString('\n')
        line     = line[:len(line)-1]

        // compute a block with this content
        block = Compute(line, block)
        block.Print()
    }
}

// mine the next block on the chain
func Compute(content string, previous *Block) *Block {
    i    := 0
    sha  := sha256.New()
    hash := "none"

    // add padding data until a hash matching the desired difficulty is found
    for !strings.HasPrefix(hash, prefix) {
        sha.Reset()
        sha.Write([]byte(previous.Hash))   // previous hash
        sha.Write([]byte(content))         // block content
        sha.Write([]byte(strconv.Itoa(i))) // padding

        hash = hex.EncodeToString(sha.Sum(nil))
        i += 1
    }

    // return the mined block
    return &Block {
        Hash:     hash,
        Previous: previous.Hash,
        Content:  content,
        Random:   i,
    }
}
