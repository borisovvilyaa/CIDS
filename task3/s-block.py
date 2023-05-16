def s_block_direct(input_data):
    t1 = input_data >> 4  # Get the first tetrad
    t2 = input_data & 0x0F  # Get the second tetrad

    # Define the S-box transformation for each tetrad
    sbox = {
        0x0: 0x7,
        0x1: 0x3,
        0x2: 0x8,
        0x3: 0xA,
        0x4: 0x1,
        0x5: 0xC,
        0x6: 0xE,
        0x7: 0x5,
        0x8: 0x9,
        0x9: 0x2,
        0xA: 0xB,
        0xB: 0x4,
        0xC: 0x6,
        0xD: 0xD,
        0xE: 0x0,
        0xF: 0xF,
    }

    # Apply S-box transformation to each tetrad
    transformed_t1 = sbox[t1]
    transformed_t2 = sbox[t2]

    # Combine the transformed tetrads
    output_data = (transformed_t1 << 4) | transformed_t2

    return output_data


def s_block_inverse(output_data):
    t1 = output_data >> 4  # Get the first tetrad
    t2 = output_data & 0x0F  # Get the second tetrad

    # Define the inverse S-box transformation for each tetrad
    inverse_sbox = {
        0x0: 0xE,
        0x1: 0x4,
        0x2: 0x9,
        0x3: 0x1,
        0x4: 0xB,
        0x5: 0x7,
        0x6: 0xC,
        0x7: 0x0,
        0x8: 0x2,
        0x9: 0x8,
        0xA: 0x3,
        0xB: 0xA,
        0xC: 0x6,
        0xD: 0xD,
        0xE: 0x5,
        0xF: 0xF,
    }

    # Apply inverse S-box transformation to each tetrad
    inverse_t1 = inverse_sbox[t1]
    inverse_t2 = inverse_sbox[t2]

    # Combine the inverse transformed tetrads
    input_data = (inverse_t1 << 4) | inverse_t2

    return input_data


# Example usage
input_data = 0b11011010
direct_transformed_data = s_block_direct(input_data)
inverse_transformed_data = s_block_inverse(direct_transformed_data)

print(f"Input Data: {bin(input_data)}")
print(f"Direct Transformed Data: {bin(direct_transformed_data)}")
print(f"Inverse Transformed Data: {bin(inverse_transformed_data)}")
