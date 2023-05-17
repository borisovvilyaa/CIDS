class SBlock:
    def __init__(self):
        self.sbox = {
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
        self.inverse_sbox = {
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

    def s_block_direct(self, input_data):
        """
        Apply the S-box transformation to the input data.

        @param input_data: The input data to be transformed.
        @return: The transformed output data.
        """
        t1 = input_data >> 4  # Get the first tetrad
        t2 = input_data & 0x0F  # Get the second tetrad

        # Apply S-box transformation to each tetrad
        transformed_t1 = self.sbox[t1]
        transformed_t2 = self.sbox[t2]

        # Combine the transformed tetrads
        output_data = (transformed_t1 << 4) | transformed_t2

        return output_data

    def s_block_inverse(self, output_data):
        """
        Apply the inverse S-box transformation to the output data.

        @param output_data: The output data to be transformed inversely.
        @return: The inverse transformed input data.
        """
        t1 = output_data >> 4  # Get the first tetrad
        t2 = output_data & 0x0F  # Get the second tetrad

        # Apply inverse S-box transformation to each tetrad
        inverse_t1 = self.inverse_sbox[t1]
        inverse_t2 = self.inverse_sbox[t2]

        # Combine the inverse transformed tetrads
        input_data = (inverse_t1 << 4) | inverse_t2

        return input_data


# Example usage
# s_block = SBlock()
# input_data = 0b11011010
# direct_transformed_data = s_block.s_block_direct(input_data)
# inverse_transformed_data = s_block.s_block_inverse(direct_transformed_data)

# print(f"Input Data: {bin(input_data)}")
# print(f"Direct Transformed Data: {bin(direct_transformed_data)}")
# print(f"Inverse Transformed Data: {bin(inverse_transformed_data)}")
