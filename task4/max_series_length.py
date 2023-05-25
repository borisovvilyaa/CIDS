import random

def max_series_length_test(sequence) -> bool:
    """
    Performs the maximum series length test on a given sequence of bits.

    @param sequence: The sequence of bits to test.
    @return: True if the sequence meets the randomness criteria, False otherwise.
    """
    max_zero_series = 0  # Maximum length of zero series
    max_one_series = 0  # Maximum length of one series

    zero_series = 0  # Current length of zero series
    one_series = 0  # Current length of one series

    for bit in sequence:
        if bit == '0':
            # Increase the length of zero series
            zero_series += 1

            # Update the maximum length of zero series if the current length is greater
            if zero_series > max_zero_series:
                max_zero_series = zero_series

            # Reset the counter for one series
            one_series = 0
        else:
            # Increase the length of one series
            one_series += 1

            # Update the maximum length of one series if the current length is greater
            if one_series > max_one_series:
                max_one_series = one_series

            # Reset the counter for zero series
            zero_series = 0

    # Compare with the maximum allowed series length (36 bits)
    if max_zero_series > 36 or max_one_series > 36:
        return False  # The sequence does not meet the randomness criteria
    else:
        return True  # The sequence meets the randomness criteria

    
# Generating a sequence of bits with a length of 20,000
sequence = ''.join(random.choices('01', k=20000))

# Performing the maximum series length test
result = max_series_length_test(sequence)

if result:
    print("The sequence is considered random.")
else:
    print("The sequence does not meet the randomness criteria.")
