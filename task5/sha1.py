import struct
import hashlib
import timeit

def sha1(data):
    # Initialization of hash function values
    h0 = 0x67452301
    h1 = 0xEFCDAB89
    h2 = 0x98BADCFE
    h3 = 0x10325476
    h4 = 0xC3D2E1F0

    # Convert the message to a bytearray
    message = bytearray(data)

    # Padding the message
    original_length = (8 * len(message)) & 0xffffffffffffffff
    message.append(0x80)
    while len(message) % 64 != 56:
        message.append(0x00)
    message += struct.pack('>Q', original_length)

    # Function to perform left rotation
    def rotate_left(n, b):
        return ((n << b) | (n >> (32 - b))) & 0xffffffff

    # SHA-1 round function
    def sha1_round(t, a, b, c, d, e, w):
        if t <= 19:
            f = (b & c) | ((~b) & d)
            k = 0x5A827999
        elif t <= 39:
            f = b ^ c ^ d
            k = 0x6ED9EBA1
        elif t <= 59:
            f = (b & c) | (b & d) | (c & d)
            k = 0x8F1BBCDC
        else:
            f = b ^ c ^ d
            k = 0xCA62C1D6

        temp = (rotate_left(a, 5) + f + e + k + w) & 0xffffffff
        e = d
        d = c
        c = rotate_left(b, 30)
        b = a
        a = temp

        return a, b, c, d, e

    # Main hash computation loop
    for i in range(0, len(message), 64):
        chunk = message[i:i + 64]

        # Split chunk into 16 words of 32-bits each
        words = list(struct.unpack('>16L', chunk))

        # Extend the 16 words to 80
        for j in range(16, 80):
            words.append(rotate_left(words[j - 3] ^ words[j - 8] ^ words[j - 14] ^ words[j - 16], 1))

        # Initialize state variables
        a = h0
        b = h1
        c = h2
        d = h3
        e = h4

        # Perform the hash computation for the current chunk
        for t in range(80):
            a, b, c, d, e = sha1_round(t, a, b, c, d, e, words[t])

        # Update the hash values
        h0 = (h0 + a) & 0xffffffff
        h1 = (h1 + b) & 0xffffffff
        h2 = (h2 + c) & 0xffffffff
        h3 = (h3 + d) & 0xffffffff
        h4 = (h4 + e) & 0xffffffff

    # Return the hash value as a hexadecimal string
    return '{:08x}{:08x}{:08x}{:08x}{:08x}'.format(h0, h1, h2, h3, h4)


# Test the sha1 function
data = b'Hello, world!'

execution_time_sha1 = timeit.timeit(lambda: sha1(data), number=1) * 1000

print("Custom SHA-1 Hash:", sha1(data))
print("Execution Time (Custom SHA-1):", execution_time_sha1, "ms")

execution_time_hashlib = timeit.timeit(lambda: hashlib.sha1(b'Hello, world!').hexdigest(), number=1) * 1000

print("hashlib SHA-1 Hash:", hashlib.sha1(b'Hello, world!').hexdigest())
print("Execution Time (hashlib SHA-1):", execution_time_hashlib, "ms")

print("Custom SHA-1 equle library:", sha1(data) == hashlib.sha1(b'Hello, world!').hexdigest())