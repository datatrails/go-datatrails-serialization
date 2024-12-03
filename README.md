# Datatrails Serialization

## Overview

Repository for go modules that serialize data, that will be added to the immutable log.


## Merkle Log

For the merkle immutable log, once the data has been serialized, it can be hashed into the `Event Hash`, which when combined with the `MMR Salt` can
be used to find the corresponding `MMR Entry`

