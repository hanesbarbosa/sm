# Slot Machine: A Compression Transform For DNA Files

This is an expanded implementation of our paper [Slot Machine: A Compression Transform for Multi-FASTA Files](https://ieeexplore.ieee.org/abstract/document/10826889/?casa_token=8NkFezt277UAAAAA:WlgEPKEht3x_wkZ6aus6C4FKKFsxRusI-9L6aU11WP1m-NsEAWTZM_wmPVZJR_rCEln1oxW_WWQUR_k), presented at the 15th International Conference on Information and Communication Technology Convergence (ICTC) in South Korea. The "expanded" nature of this development lies on the attempt to improve the compression scheme for the encoding of other Deoxyribonucleic Acid (DNA) file formats.

Originally, Slot Machine is a reference-free compression transform designed specifically for multi-FASTA genome files. This transform leverages the distribution of characters in multi-FASTA files to decrease their size by approximately 66% while allowing a user to search and decompress only the parts of the file needed at any given time.

## Key Features

- Selective decompression: Unlike gzip which requires full file decompression, Slot Machine supports compressed pattern matching and decompression of regions of interest.
- Adaptive compression performance: Achieves compression rates of 66% (i.e., first stage) while keeping pattern recognition or up to 80.96% compression rate (i.e., second stage) for archival purposes.
- Compatible with existing tools: The output of the first stage can be further compressed by pairing with existing general-purpose compression algorithms.
- Reference-free: Doesn't require a reference genome for compression.
- Suitable for homomorphic encryption: The result of the first stage procedure can be used for homomorphic encryption purposes, allowing computation under secrecy.
- Allows the parallel processing of encoding, decoding and the distributed access of the same file.

## How Does It Work?

1. Funnel Codec
2. Slot Machine

### Funnel Codec

The Funnel Codec will receive lines from the original file, converting them into a set of new 8-bit characters called **composites**. A composite can represent one or more consecutive original characters. It is comprised by a 2-bit **instruction** and a 6-bit **code**, in a total of 1 byte. For instance, for the 7-byte input 'GCTAN-G', the Funnel Encoding could output (depending on the statistical model) a 3-byte string comprised by 1 byte for the unprintable escape character, 1 byte for the latin capital letter C with cedilla (i.e., 'Ç') and 1 byte for the latin capital letter O with tilde (i.e., 'Õ').

Each resulting composite is based on a piecewise function that analyzes if the given block or line is part of a **comment** or **nucleotide** section. If the line belongs to a comment section (e.g., in a Multi-FASTA format it starts with '>' and ends in a '\n' character), all characters but the delimiters '>' and '\n' are returned without change. When encoding, the delimiters will receive a special code to allow the identification of which compressed section is originally from a nucleotide or comment block. If the block is a nucleotide section (i.e., only comprised by DNA bases), then substrings are processed according to their membership on set Λ³. The alphabet for the procedure is divided as follows:

- Λ (Lambda) is the set of the 4 more frequent bases.
- Ω (Omega) is the set of the ambiguous bases, errors, etc.
- Σ (Sigma) is the set of all allowed bases (i.e., Σ = Λ U Ω).

We know that 6-bit codes only allow us to represent 64 strings, however we need to guarantee that every unique character from Σ is represented, besides the codes for special group of characters (e.g., 1 comment's initial character and 1 line feed character). Therefore, the number of special aggregated strings (i.e., 2 or 3 characters not in Λ³) will have 64 - (|Σ| + 2) possible codes. In general, the possible combination of special codes from the alphabet can be described as one of the following possibilities:

1. 3 characters not in Λ³: Ω³, Λ¹Ω², Ω²Λ¹, Λ²Ω¹, Ω¹Λ², Ω¹Λ¹Ω¹, Λ¹Ω¹Λ¹;
2. 2 characters: Ω², Λ², Λ¹Ω¹, Ω¹Λ¹;
3. 1 character: Ω¹, Λ¹.

### Slot Machine