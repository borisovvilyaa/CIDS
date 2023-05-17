def p_block(input_data):
    """
    Implements the P-Block encryption.
    @param input_data: Input data (8 bits)
    @return: Encrypted data (8 bits)
    """
    p_table = [
        0x1, 0x5, 0x2, 0x0, 0x3, 0x7, 0x4, 0x6
    ]
    output_data = 0

    # Passing through the P-Block
    for i in range(8):
        bit = (input_data >> i) & 0x1  # Obtaining the i-th bit of the input data
        output_data |= (bit << p_table[i])  # Shifting the bit according to the P-Block table

    return output_data

def inverse_p_block(input_data):
    """
    Implements the inverse transformation of the P-Block for decryption.

    @param input_data: Encrypted data (8 bits)
    @return: Decrypted data (8 bits)
    """
    p_table = [0x3, 0x0, 0x2, 0x4, 0x6, 0x1, 0x7, 0x5]  # Predefined inverse P-Block table
    output_data = 0

    # Passing through the inverse P-Block table
    for i in range(8):
        bit = (input_data >> i) & 0x1
        output_data |= (bit << p_table[i])
    return output_data

input_data = 0b11011010 # Input data (8 bits)
encrypted_data = p_block(input_data)
decrypted_data = inverse_p_block(encrypted_data)

print("Input Data:", bin(input_data))
print("Encrypted Data:", bin(encrypted_data))
print("Decrypted Data:", bin(decrypted_data))