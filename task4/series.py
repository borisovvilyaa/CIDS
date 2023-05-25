import random

def count_zero_series(sequence) -> list:
    """
    Count the number of series of consecutive zeros of different lengths in the given sequence.

    @param sequence: The input sequence of binary digits.
    @return: A list containing the counts of series of different lengths.
    """
    series_counts = [0, 0, 0, 0, 0, 0]  # Initialize counters for each series length
    current_series_length = 0  # Variable to track the current series length

    for number in sequence:
        if number == '1':  # If a '1' is encountered, reset the current series length
            current_series_length = 0
        else:
            current_series_length += 1  # If a '0' is encountered, increase the current series length
            if current_series_length >= 6:  # If the current series length is >= 6, increment the counter for series of length 6 or more
                series_counts[5] += 1
            else:
                series_counts[current_series_length - 1] += 1  # Increment the counter for the corresponding series length

    return series_counts


def test_sequence(sequence) -> bool:
    """
    Test the given sequence for randomness based on the counts of zero series.

    @param sequence: The input sequence of binary digits.
    @return: True if the sequence meets the randomness criteria, False otherwise.
    """
    series_counts = count_zero_series(sequence)
    print(series_counts)
    # Check if all series counts fall within the specified ranges
    return (2267 <= series_counts[0] <= 2733 and
            1079 <= series_counts[1] <= 1421 and
            502 <= series_counts[2] <= 748 and
            223 <= series_counts[3] <= 402 and
            90 <= series_counts[4] <= 223 and
            90 <= series_counts[5] <= 223)


# Generate a random sequence
sequence = ''.join(random.choices('01', k=20000))

# Test the sequence
result = test_sequence(sequence)

# Output the result
if result:
    print("The sequence is considered random.")
else:
    print("The sequence does not meet the randomness criteria.")
