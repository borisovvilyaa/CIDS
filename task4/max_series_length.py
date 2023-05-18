def max_series_length_test(sequence):
    zero_series_length = 0
    one_series_length = 0
    max_zero_series_length = 0
    max_one_series_length = 0

    for bit in sequence:
        if bit == 0:
            zero_series_length += 1
            if zero_series_length > max_zero_series_length:
                max_zero_series_length = zero_series_length
            one_series_length = 0
        else:
            one_series_length += 1
            if one_series_length > max_one_series_length:
                max_one_series_length = one_series_length
            zero_series_length = 0

    if max_zero_series_length > 36 or max_one_series_length > 36:
        return False
    else:
        return True
    
sequence = [0]
is_random = max_series_length_test(sequence)

if is_random:
    print("The sequence is random.")
else:
    print("The sequence is not random.")