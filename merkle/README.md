# Merkle tree exercise

* Generate a Merkle tree given the following body of data, using SHA-256 as your hashing algorithm
  - Your data is the following blocks:
  - "We", "hold", "these", "truths", "to", "be, "self-evident", "that"
* In the internal stages, concatenate blocks as `hash("#{blockA}||#{blockB}")`
* Do not use any special padding for leaf vs internal nodes
* The merkle root in hex output should be equal to `4a359c93d6b6c9beaa3fe8d8e68935aa5b5081bd2603549af88dee298fbfdd0a`

Bonus: create a padding scheme so that arbitrary numbers of blocks can be Merkleized.

Bonus 2: add different padding to the leaves as opposed to internal nodes, so that preimage attacks are impossible.

Bonus 3: implement an interface for Merkle proofs. Have a `generate_proof(block)` method and a `verify_inclusion(block, proof)` method.


## Javascript implementation:
```
var crypto = require('crypto');
var hash = crypto.createHash('sha256').update('your message goes here').digest('hex');
```

## Python implementation: 
```
import hashlib
hashlib.sha256("hello").hexdigest()
```

## Ruby implementation:
```
require 'digest'
Digest::SHA2.hexdigest('your string here')
```
