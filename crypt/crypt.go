package crypt

import (
    "crypto/aes"
    "crypto/cipher"
    "fmt"
)

func Encode(plain string) string {
    key := []byte("examplekey32byteslong")
    plaintext := []byte(plain)

    // Create a new AES cipher block
    block, err := aes.NewCipher(key)
    if err != nil {
        fmt.Println(err)
        return ""
    }

    // Get the block size
    blockSize := block.BlockSize()

    // Create a new initialization vector
    iv := make([]byte, blockSize)

    // Encrypt the plaintext
    ciphertext := make([]byte, len(plaintext))
    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext, plaintext)
    fmt.Printf("Ciphertext: %x\n", ciphertext)

    return string(ciphertext)
}

func Decode(ciphertext string) string {
    key := []byte("examplekey32byteslong")

    // Create a new AES cipher block
    block, err := aes.NewCipher(key)
    if err != nil {
        fmt.Println(err)
        return ""
    }

    // Get the block size
    blockSize := block.BlockSize()

    // Create a new initialization vector
    iv := make([]byte, blockSize)

    decrypted := []byte{}
    stream := cipher.NewCFBDecrypter(block, iv)
    stream.XORKeyStream(decrypted, []byte(ciphertext))
    fmt.Printf("Decrypted: %s\n", decrypted)

    return string(decrypted)
}
