import random

def monobit_test(sequence):
    count_ones = sequence.count('1')
    
    if 9654 < count_ones < 10346:
        return True
    else:
        return False

# Generating a sequence of bits with a length of 20,000
sequence = ''.join(random.choices('01', k=20000))

# Performing the monobit test
result = monobit_test(sequence)

if result:
    print("The sequence is considered random.")
else:
    print("The sequence does not meet the randomness criteria.")
