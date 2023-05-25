import random

def count_series_length(sequence) -> list:
    """
    Count the series lengths of consecutive 0s and 1s in a given sequence.
    
    @param sequence: The input sequence of bits.
    @type sequence: str
    
    @return: A dictionary containing the count of series for each length.
    @rtype: dict
    """
    series_lengths = {'0': [0]*7, '1': [0]*7}  # Stores the count of series for each length
    
    current_char = sequence[0]  # Initial character of the sequence
    current_length = 1  # Initial length of the series

    for i in range(1, len(sequence)):
        if sequence[i] == current_char:  # If the current character matches the previous one
            current_length += 1
        else:  # If the current character is different from the previous one
            series_lengths[current_char][min(current_length, 6)] += 1
            current_char = sequence[i]
            current_length = 1

    # Update the count of series for the last character in the sequence
    series_lengths[current_char][min(current_length, 6)] += 1
    
    return series_lengths
    

def test_random_sequence(sequence) -> bool:
    """
    Test if a given sequence meets the randomness criteria based on the count of series.
    
    @param sequence: The input sequence of bits.
    @type sequence: str
    
    @return: True if the sequence is considered random, False otherwise.
    @rtype: bool
    """
    series_lengths = count_series_length(sequence)
    # Check if the count of series falls within the appropriate ranges
    if (
        2267 <= series_lengths['0'][1] <= 2733 and
        1079 <= series_lengths['0'][2] <= 1421 and
        502 <= series_lengths['0'][3] <= 748 and
        223 <= series_lengths['0'][4] <= 402 and
        90 <= series_lengths['0'][5] <= 223 and
        90 <= series_lengths['0'][6] <= 223 and
        2267 <= series_lengths['1'][1] <= 2733 and
        1079 <= series_lengths['1'][2] <= 1421 and
        502 <= series_lengths['1'][3] <= 748 and
        223 <= series_lengths['1'][4] <= 402 and
        90 <= series_lengths['1'][5] <= 223 and
        90 <= series_lengths['1'][6] <= 223
    ):
        return True  
    else:
        return False


# Generating a sequence of bits with a length of 20,000
sequence = ''.join(random.choices('01', k=20000))

# Performing the maximum series length test
result = test_random_sequence(sequence)

if result:
    print("The sequence is considered random.")
else:
    print("The sequence does not meet the randomness criteria.")
