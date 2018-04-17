# Blockchain 

## Transaction 

Create a `Transaction` class, which includes:
 * `from` (public key)
 * `to` (public key)
 * `amount` (int) 
 * `signature` (string). 
 
Each block contain one transaction for simplicity.

## Actions

* automatically create blocks

* compute proof of work for each block

* chain them all together in a blockchain
    * append new blocks to this blockchain.
    
### Validation    

Make your blockchain validate that all of the transactions within the blockchain are valid—i.e., each transaction does not put any party into a negative balance. You'll want to start all blockchains with a genesis block that gives someone seed money to begin the chain.

### Fork choice rule
Implement a fork choice rule. Given your current blockchain and a new blockchain, write a function that replaces the old blockchain with a new blockchain if it's longer than your current blockchain (and valid!)

### Gossip
Go back to your gossip protocol and adapt it to gossip blockchains rather than random messages. You may want to use a good marshalling library for your language to make it easier to serialize and deserialize your objects. Apply the fork choice rule to any messages you receive.

### PKI
Give each node a public key and private key. Give it an endpoint so that another node can query it for its public key. Make it automatically sign each transaction it creates.

### Transfer endpoint
Give each node a `transfer` endpoint, which will make it query another node for its public key, and then transfer that node some of its money.

