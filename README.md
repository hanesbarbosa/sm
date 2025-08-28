# Slot Machine: A Compression Transform For DNA Files

This is an expanded implementation of our paper [Slot Machine: A Compression Transform for Multi-FASTA Files](https://ieeexplore.ieee.org/abstract/document/10826889/?casa_token=8NkFezt277UAAAAA:WlgEPKEht3x_wkZ6aus6C4FKKFsxRusI-9L6aU11WP1m-NsEAWTZM_wmPVZJR_rCEln1oxW_WWQUR_k), presented at the 15th International Conference on Information and Communication Technology Convergence (ICTC) in South Korea. The "expanded" nature of this development lies on the attempt to improve the compression scheme for the encoding of other Deoxyribonucleic Acid (DNA) file formats.

Originally, Slot Machine is a reference-free compression transform designed specifically for multi-FASTA genome files. This transform leverages the distribution of characters in multi-FASTA files to decrease their size by approximately 66% while allowing a user to search and decompress only the parts of the file needed at any given time.

## Key Features

- Selective decompression: Unlike gzip which requires full file decompression, Slot Machine supports compressed pattern matching and decompression of regions of interest.
- Adaptive compression performance: Achieves compression rates of 66% (i.e., first stage) while keeping pattern recognition or up to 80.96% compression rate (i.e., second stage) for archival purposes.
- Compatible with existing tools: The output of the first stage can be further compressed by pairing with existing general-purpose compression algorithms.
- Reference-free: Doesn't require a reference genome for compression.
- Suitable for homomorphic encryption: The result of the first stage procedure can be used for homomorphic encryption purposes, allowing computation under secrecy.
