from random import randint

class ElGamal:
    def __init__(self) -> None:
        self.p = self.generate_p()
        self.g = self.find_primitive_root(self.p)
        self.a = self.generate_private_key()
        self.b = pow(self.g, self.a, self.p)

    def getParameters(self):
        return self.p, self.g, self.a, self.b

    def is_prime(self, number):
        if number < 2:
            return False
        for i in range(2, int(number ** 0.5) + 1):
            if number % i == 0:
                return False
        return True

    def find_primitive_root(self, p):
        for g in range(2, p):
            if self.is_primitive_root(g, p):
                return g

    def is_primitive_root(self, g, p):
        elements = set()
        for i in range(1, p):
            elements.add(pow(g, i, p))
        for i in range(1, p):
            if i not in elements:
                return False
        return True

    def generate_p(self):
        p = randint(2048, 4096)
        while True:
            if self.is_prime(p):
                return p
            else:
                p = randint(2048, 4096)

    def generate_private_key(self):
        a = randint(1, self.p - 1)
        return a

    def encrypt(self, message):
        blocks = self.split_message_into_blocks(message)
        encrypted_blocks = []
        for block in blocks:
            k = randint(1, self.p - 1)
            x = pow(self.g, k, self.p)
            y = (pow(self.b, k, self.p) * block) % self.p
            encrypted_blocks.append((x, y))
        return encrypted_blocks

    def decrypt(self, encrypted_blocks):
        decrypted_blocks = []
        for block in encrypted_blocks:
            x, y = block
            s = pow(x, self.a, self.p)
            s_inverse = self.inverse_modulo(s, self.p)
            decrypted_block = (y * s_inverse) % self.p
            decrypted_blocks.append(decrypted_block)
        decrypted_message = self.merge_blocks_into_message(decrypted_blocks)
        return decrypted_message

    def inverse_modulo(self, a, m):
        g, x, y = self.extended_gcd(a, m)
        if g != 1:
            raise ValueError("Modular inverse does not exist.")
        return x % m

    def extended_gcd(self, a, b):
        if a == 0:
            return b, 0, 1
        else:
            g, x, y = self.extended_gcd(b % a, a)
            return g, y - (b // a) * x, x

    def split_message_into_blocks(self, message):
        block_size = len(str(self.p)) - 1  # Size of each block is determined based on the number of digits in p
        blocks = []
        while message > 0:
            blocks.append(message % (10 ** block_size))
            message //= 10 ** block_size
        return blocks[::-1]

    def merge_blocks_into_message(self, blocks):
        message = 0
        for block in blocks:
            message *= 10 ** len(str(block))
            message += block
        return message

el = ElGamal()
p, g, a, b = el.getParameters()
print("p: %s" % p)
print("g: %s" % g)
print("a: %s" % a)
print("b: %s" % b)

message = 12345678901234567890
print("Original message: %s" % message)

encrypted_blocks = el.encrypt(message)
print("Encrypted blocks: %s" % encrypted_blocks)

decrypted_message = el.decrypt(encrypted_blocks)
print("Decrypted message: %s" % decrypted_message)
