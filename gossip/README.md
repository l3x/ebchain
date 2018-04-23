# Building a gossip protocol

## Goal
We're going to build a gossip protocol that gossips with a network of peers about their current favorite books. Each node initially will connect to only one other node to bootstrap its peer list. Then all nodes should know the favorite books of all other nodes in the system. We'll have nodes run as separate processes listening on different ports on your local machine, and they'll pass JSON-encoded messages via HTTP (to make our lives easier).

## State
We'll use a gossip protocol to keep track of each node's current favorite book. You can find a list of books in `books.txt`.

Each node's book should be randomly re-sampled from the pool of all books once every ~10 seconds. Once it chooses a new favorite book, it should flood its peers with this message.

You need to have each node keep track of their own incrementing version number, so we can keep track of their state and order messages. In a gossip protocol we will often receive messages out of order, so we need to know which one is most recent.

The node should also keep a cache of the recent messages it's received. Normally we'd want to cull this, but for now we can just let it grow in memory.

## Endpoints
When you launch the process, you will need to give it the port of another node to boostrap its peers from. That port number should be an argument passed via the command line.

Each node needs the following endpoints:

* GET /peers/ (for bootstrapping into the network)
* POST /gossip/ (for trading gossip between nodes)
  - You can decide whether gossip is bi-directional (I tell you my state, you tell me your state, and we both update)
  - Or if you want you can make gossip uni-directional. All books will eventually get propagated through the system, so new nodes will eventually get up to speed. (Not strictly true in Bitcoin.)

## Message format
Your messages will need the following:

* UUID (for deduplication)
* Originating port (your identity)
* Version number
* TTL
* Payload

Suggested strategy:
* Start by defining your message format (write out some example messages)
* Figure out the state that each node needs to hold
* Write up your message update logic—upon receiving a gossip message, how do you update your view of the world? You should be able to run this on your example message and get a correct state transition.
* Then write a little UI code so you can easily inspect what each node is doing.
* Then write your networking and gossip logic, and test across multiple nodes!

NOTE: Observing and debugging distributed systems can be hard. I recommend investing in your UI: try adding some colors and an organized table to display your state for each node. It'll make debugging a lot easier.

## Extra credit (if you have time):

Can you make the system fault tolerant? Try killing nodes and ensuring your system stays up and active.

Have your nodes randomly gossip to a subset of peers, rather than flooding. You need a large network and a bounded peer count to see this working, but try it out and see what performance you can get.

If you really want to go all-out, write a simulator that creates 100 node processes, seeds them, and monitors the status of the network. You can then use this to try out different parameters (like number of peers gossiped to, TTLs, etc.) and see how long it takes to propagate messages, what the total message count is, etc.


## Build gossip binary

```
~/GOPATHs/ebchain/src/github.com/l3x/ebchain/gossip $ go install && gossip -p 7001 -b 7000
```


## Configure Goland to run main package:

Run > Run... > Edit Configurations... > Go Build > main - gossip

**Run kind:** package
**Package path:** github.com/l3x/ebchain/gossip 


## Run App

```
forego start
```