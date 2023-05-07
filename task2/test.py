import unittest
from Numbers import *


class TestBigInteger(unittest.TestCase):
    def setUp(self):
        self.number_big_a = BigInteger()
        self.number_big_a.setHex(
            "1182d8299c0ec40ca8bf3f49362e95e4ecedaf82bfd167988972412095b13db8"
        )

        self.number_big_b = BigInteger()
        self.number_big_b.setHex(
            "a78865c13b14ae4e25e90771b54963ee2d68c0a64d4a8ba7c6f45ee0e9daa65b"
        )

    def test_XOR(self):
        number_xor = XOR(self.number_big_a, self.number_big_b)
        print(f"Result test_XOR: {number_xor.getHex()}")
        expected = "0xB60ABDE8A71A6A428D5638388367F60AC1856F24F29BEC3F4F861FC07C6B9BE3"
        self.assertEqual(number_xor.getHex(), expected)

    def test_OR(self):
        number_or = OR(self.number_big_a, self.number_big_b)
        print(f"Result test_OR: {number_or.getHex()}")
        expected = "0xB78AFDE9BF1EEE4EADFF3F79B76FF7EEEDEDEFA6FFDBEFBFCFF65FE0FDFBBFFB"
        self.assertEqual(number_or.getHex(), expected)

    def test_AND(self):
        number_and = AND(self.number_big_a, self.number_big_b)
        print(f"Result test_AND: {number_and.getHex()}")
        expected = "0x18040011804840C20A90741340801E42C6880820D4003808070402081902418"
        self.assertEqual(number_and.getHex(), expected)

    def test_INV(self):
        number_inv = INV(self.number_big_a)
        print(f"Result test_INV: {number_inv.getHex()}")
        expected = "0xEE7D27D763F13BF35740C0B6C9D16A1B1312507D402E9867768DBEDF6A4EC247"
        self.assertEqual(number_inv.getHex(), expected)

    def test_ShiftLeft(self):
        shift_l = ShiftLeft(self.number_big_a, 8)
        print(f"Result test_ShiftLeft: {shift_l.getHex()}")
        expected = "0x82D8299C0EC40CA8BF3F49362E95E4ECEDAF82BFD167988972412095B13DB800"
        self.assertEqual(shift_l.getHex(), expected)

    def test_ShiftRight(self):
        shift_r = ShiftRight(self.number_big_a, 8)
        print(f"Result test_ShiftRight: {shift_r.getHex()}")
        expected = "0x00118374009C0F6C00A8BF7500362F8100ECEE6E00BFD1F0008972D60095B13D"
        self.assertEqual(shift_r.getHex(), expected)

    def test_ADD(self):
        add = ADD(self.number_big_a, self.number_big_b)
        print(f"Result test_ADD: {add.getHex()}")
        expected = "0xB90B3DEAD723725ACEA846BAEB77F9D31A5670290D1BF3405066A0017F8BE413"
        self.assertEqual(add.getHex(), expected)

    def test_SUB(self):
        sub = SUB(self.number_big_a, self.number_big_b)
        print(f"Result test_SUB: {sub.getHex()}")
        expected = (
            "0x96058D979F05EA427D29C8277F1ACE09407B11248D79240E3D821DBF542968A3FFFFFFFF"
        )
        self.assertEqual(sub.getHex(), expected)

    def test_MOD(self):
        mod = MOD(self.number_big_a, self.number_big_b)
        print(f"Result test_MOD: {mod.getHex()}")
        expected = "0x1182D8299C0EC40CA8BF3F49362E95E4ECEDAF82BFD167988972412095B13DB8"
        self.assertEqual(mod.getHex(), expected)

    def test_MUL(self):
        mul = MUL(self.number_big_a, self.number_big_b)
        print(f"Result test_MUL: {mul.getHex()}")
        expected = "0xB75AF7BD1E07DA3A4E2CAEDFB839F91EC4387BFC24245EE5F043E74A992A6443A0E63312210E6723D4D396D87A0DBC5FE7DEA5296A5A61842F5A2C12AB64068"
        self.assertEqual(mul.getHex(), expected)


if __name__ == "__main__":
    unittest.main()
