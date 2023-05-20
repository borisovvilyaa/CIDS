from random import randint
def max_series_length_test(sequence):
    
    count = sequence.count(0)
    print(count)

    if count > 9654 and  count < 10346:
        return True
    else:
        return False
    
sequence = [randint(0,1) for i in range(20001)]
is_random = max_series_length_test(sequence)

if is_random:
    print("The sequence is random.")
else:
    print("The sequence is not random.")