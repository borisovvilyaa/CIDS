import random

def poker_test(sequence) -> bool:
    """
    Performs a poker test on the given binary sequence.

    @param sequence: The binary sequence to test.
    @return: True if the sequence passes the test, False otherwise.
    """
    m = 4  # Poker block length
    k = len(sequence) // m  # Number of poker blocks in the sequence
    block_counts = [0] * (2 ** m)  # Array to store the count of each block

    # Count the occurrences of each block
    for i in range(k):
        block = sequence[i * m: (i + 1) * m]  # Select a block of length m
        block_num = int(''.join(map(str, block)), 2)  # Convert the block to an integer
        block_counts[block_num] += 1  # Increase the counter for the corresponding block

    # Calculate the X3 parameter
    x3 = (2 ** m / k) * sum([n_i ** 2 for n_i in block_counts]) - k

    # Compare with the constant and determine the result
    if 1.03 < x3 < 57.4:
        return True
    else:
        return False

sequence = ''.join(random.choices('01', k=20000))

# Performing the poker test
result = poker_test(sequence)

if result:
    print("The sequence is considered random.")
else:
    print("The sequence does not meet the randomness criteria.")
