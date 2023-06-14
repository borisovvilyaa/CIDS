from random import randint
import hashlib

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

    def sign_message(self, message):
        k = randint(1, self.p - 1)
        r = pow(self.g, k, self.p)

        h = self.hash_message(message)
        inverse_k = self.mod_inverse(k, self.p - 1)
        s = ((h - self.a * r) * inverse_k) % (self.p - 1)

        return r, s

    def hash_message(self, message):
        hashed_message = hashlib.sha256(message.encode()).digest()
        h = int.from_bytes(hashed_message, byteorder='big')
        return h

    def mod_inverse(self, a, m):
        g, x, y = self.extended_gcd(a, m)
        if g != 1:
            raise ValueError("Неможливо знайти обернений елемент")
        return x % m

    def extended_gcd(self, a, b):
        if a == 0:
            return b, 0, 1
        else:
            g, x, y = self.extended_gcd(b % a, a)
            return g, y - (b // a) * x, x

    def verify_signature(self, message, r, s):
        h = self.hash_message(message)
        inverse_s = self.mod_inverse(s, self.p - 1)

        u1 = (h * inverse_s) % (self.p - 1)
        u2 = (r * inverse_s) % (self.p - 1)

        y = self.mod_inverse(self.b, self.p)
        v = (pow(self.g, u1, self.p) * pow(y, u2, self.p)) % self.p

        return v == r

el = ElGamal()
p, g, a, b = el.getParameters()
print("p: %s" % p)
print("g: %s" % g)
print("a: %s" % a)
print("b: %s" % b)

message = "Hello, world!"
r, s = el.sign_message(message)
print("Підпис повідомлення:")
print("r: %s" % r)
print("s: %s" % s)

valid_signature = el.verify_signature(message, r, s)
print("Перевірка підпису: %s" % valid_signature)
