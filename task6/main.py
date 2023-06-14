from random import randint
from sha1 import sha1
class ElGamal:
    def __init__(self) -> None:
        #parameters
        self.p = self.random_number()
        self.g = self.find_primitive_root()
        
        #First Person
        self.a = randint(1, self.p-1)
        self.b = (self.g**self.a) % self.p
        
        #Signing the message
        self.k = randint(1, self.p-1)
        self.r = (self.g**self.k)%self.p
        self.m = 10 # условно, це буде геш, бо в мене геш - це стрінг, а повинен бути для self.s інт
        self.s = ((self.m - self.a*self.r)*self.k**(-1))%(self.p-1)
        
        #check signature 
        self.y = (self.b**(-1))%self.p
        self.u1 = (self.m*self.s**(-1))%(self.p-1)
        self.u2 = (self.r*self.s**(-1))%(self.p-1)
        self.v = (self.g**self.u1 * self.y ** self.u2) % self.p
        
    def get_parameter(self):
        return self.p, self.g
    def firstPerson(self):
        return self.a, self.b
    def get_signature(self):
        return self.k, self.r, self.s
    
    def check(self):
        return self.v == self.r, self.v, self.r
    # def secondPerson(self):
    #     return self.m, self.k, self.x
    def random_number(self) -> int:
        while True:
            p = randint(2048, 4096)
            if self.is_prime(p):
                return p
            
    def find_primitive_root(self):
        # Calculate Euler's totient function value (phi)
        phi = self.p - 1
        
        # Factorize phi into prime factors
        factors = self.factorize(phi)
        
        # Iterate over numbers from 1 to p-1
        for g in range(1, self.p):
            # Assume g is a primitive root
            is_primitive_root = True
            
            # Check if g raised to the power of (phi // factor) modulo p is equal to 1
            # for any factor in the factors list
            for factor in factors:
                if pow(g, phi // factor, self.p) == 1:
                    # If the condition is met, g is not a primitive root
                    is_primitive_root = False
                    break
            
            # If g is a primitive root, return it
            if is_primitive_root:
                return g
        
        # If no primitive root is found, raise a ValueError
        raise ValueError("The first-order root for the number p could not be found.")

    
    def factorize(self, n):
        factors = []
        divisor = 2
        while divisor * divisor <= n:
            if n % divisor == 0:
                factors.append(divisor)
                n //= divisor
            else:
                divisor += 1
        if n > 1:
            factors.append(n)
        return factors

    def is_prime(self, x) -> bool:
        if x < 2:
            return False
        for i in range(2, (x//2)+1):
            if x % i == 0:
                return False
        return True
o = ElGamal()
print("p %s\ng %s\n" % o.get_parameter())
print("Private Key %s\nPublic Key %s\n" %o.firstPerson())
print("k %s (only for debug)\nr %s\ns %s\n" % o.get_signature())
print("Check %s\nv %s\ns %s" % o.check())
# print("'m' is %s\n'k' is %s\n'x' is %s"%o.secondPerson())