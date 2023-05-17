import unittest
from sblock import *
from pblock import *


class TestPblock(unittest.TestCase):
    def setUp(self):
        self.p_block_cipher = PBlock()
        self.inverse_p_block_cipher = InversePBlock()
        self.input_data = 0b11011010
    def test_encrypt(self):

        encrypted_data = self.p_block_cipher.encrypt(self.input_data)
        print(f".encrypted_data: {encrypted_data}")

        expected = 0b1111001
        self.assertEqual(encrypted_data, expected)
    def test_inverse(self):
        encrypted_data = self.p_block_cipher.encrypt(self.input_data)
        decrypted_data = self.inverse_p_block_cipher.decrypt(encrypted_data)
        print(f"decrypted_data | input_data: {decrypted_data}")
        expected = self.input_data
        self.assertEqual(decrypted_data, expected)
if __name__ == "__main__":
    unittest.main()
