import random
from sympy import isprime

def generate_prime(bits):
    while True:
        p = random.getrandbits(bits)
        if p.bit_length() == bits and isprime(p):
            return p

# Використання функції generate_prime для генерації простого числа p
p = generate_prime(2048)  # Виберіть бажану довжину в бітах (2048-4096)
print(p)
