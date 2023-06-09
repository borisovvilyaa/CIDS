from random import randint

class ElGamal:
    def __init__(self) -> None:
        self.number = self.random_number()
        
    def getNumber(self):
        return self.number
    
    def makeGroup(self):
        #make group
        pass
    
    def random_number(self) -> int:
        while True:
            p = randint(2048, 4096)
            if self.isPrime(p):
                return p
             
    def isPrime(self, x) -> bool:
        for i in range(2, (x//2)+1):
            if x % i == 0:
                return False
        return True

print(ElGamal().getNumber())