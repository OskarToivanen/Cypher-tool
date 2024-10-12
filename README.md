# Cypher Tool - Command Line Encryption/Decryption

**Cypher Tool** is a simple command-line utility for encrypting and decrypting messages using algorithms like **ROT13**, **Reverse Alphabet**, **Shift6**. It features a user-friendly interface using the terminal with input validation and preserves non-alphabet characters during processing.

## Features

- **Encrypt/Decrypt**: Choose to encrypt or decrypt a message.
- **Supported Ciphers**:
  - **ROT13**: Shifts letters by 13 positions.
  - **Reverse Alphabet**: Maps letters to their reverse positions in the alphabet.
  - **Shift6**: Shifts letters by 6 positions.
- **Input Validation**: Ensures proper input format.

## Usage

1. Run the tool:
   ```bash
   $ ./cyphertool

2. Select Operation:
   ```bash
   Select operation (1/2):
   1. Encrypt.
   2. Decrypt.
   > 
3. Select cypher method: 
    ```bash
    Choose cypher:
    1. ROT13.
    2. Reverse.
    3. Shift6
    >
4. Enter the message to decrypt or encrypt:
   ```bash
   Enter the message:
   > 

## Example

1. Encrypt:
   ```bash
    ./cyphertool 
    Welcome to the Cypher Tool!

    Select operation (1/2):
    1. Encrypt.
    1. Decrypt.
    > 1
    Choose cypher:
    1. ROT13.
    2. Reverse.
    3. Shift6
    > 1
    Enter the message:
    > asd
    Encrypted message using ROT13: nfq

    Do you want to close the program? (1. Yes / 2. No): 1