class BigInteger:
    def __init__(self):
        # Initialize an empty list to store the digits of the BigInteger
        self.digits = []

    def setHex(self, hex_str: str) -> None:
        """
        Sets the value of the BigInteger using a hexadecimal string.

        @param hex_str: The hexadecimal string representing the value of the BigInteger.
        @raise ValueError: If the hex string is empty.
        """
        self.digits = []

        if hex_str == "":
            raise ValueError(
                '[setHex -> Error]: hex string must not be empty \nExample: "e035c6cfa42609b998b883bc1699df885cef74e2b2cc372eb8fa7e7"'
            )

        for i in range(len(hex_str) // 8 + 1):
            start = max(len(hex_str) - (i + 1) * 8, 0)
            end = len(hex_str) - i * 8
            chunk = hex_str[start:end]
            if chunk:
                self.digits.insert(0, int(chunk, 16))

        # print("[setHex -> Done]", self.digits)

    def getHex(self) -> str:
        """
        Returns the hexadecimal string representation of the BigInteger.

        @return: The hexadecimal string representation of the BigInteger.
        """
        hex_str = "".join("{:08x}".format(d) for d in self.digits).lstrip("0")
        if not hex_str:
            hex_str = "0"
        return "0x" + hex_str.upper()

    def __ge__(self, other):
        if len(self.digits) > len(other.digits):
            return True
        elif len(self.digits) < len(other.digits):
            return False
        else:
            for i in range(len(self.digits) - 1, -1, -1):
                if self.digits[i] > other.digits[i]:
                    return True
                elif self.digits[i] < other.digits[i]:
                    return False
            return True


class XOR:
    def __init__(self, value1: BigInteger, value2: BigInteger):
        """
        Constructor of the class, takes two BigInteger objects as parameters and XORs them.

        @param value1 (BigInteger): First BigInteger object to XOR.
        @param value2 (BigInteger): Second BigInteger object to XOR.
        """
        self.value = value1
        self.xor(value2)

    def getHex(self) -> str:
        """
        Returns the hexadecimal representation of the current BigInteger value.

        @return (str): The hexadecimal representation of the BigInteger value.
        """
        return self.value.getHex()

    def xor(self, other: BigInteger) -> None:
        """
        Performs XOR operation on the current BigInteger value and the provided BigInteger object.

        @param other (BigInteger): The BigInteger object to XOR with the current value.
        """
        if len(self.value.digits) < len(other.digits):
            self.value.digits, other.digits = other.digits, self.value.digits

        for i in range(len(other.digits)):
            self.value.digits[i] ^= other.digits[i]


class OR:
    def __init__(self, value1: BigInteger, value2: BigInteger):
        """
        Constructor of the class, takes two BigInteger objects as parameters and performs the bitwise OR operation on them.

        @param value1 (BigInteger): First BigInteger object to perform OR operation.
        @param value2 (BigInteger): Second BigInteger object to perform OR operation.
        """
        self.value = value1
        self.or_func(value2)

    def getHex(self) -> str:
        """
        Returns the hexadecimal representation of the current BigInteger value.

        @return (str): The hexadecimal representation of the BigInteger value.
        """
        return self.value.getHex()

    def or_func(self, other: BigInteger) -> None:
        """
        Performs the bitwise OR operation on the current BigInteger value and the provided BigInteger object.

        @param other (BigInteger): The BigInteger object to perform OR operation with the current value.
        """
        if len(self.value.digits) < len(other.digits):
            self.value.digits, other.digits = other.digits, self.value.digits

        for i in range(len(other.digits)):
            self.value.digits[i] |= other.digits[i]


class AND:
    def __init__(self, value1: BigInteger, value2: BigInteger):
        """
        Constructor of the class, takes two BigInteger objects as parameters and performs the bitwise AND operation on them.

        @param value1 (BigInteger): First BigInteger object to perform AND operation.
        @param value2 (BigInteger): Second BigInteger object to perform AND operation.
        """
        self.value = value1
        self.and_func(value2)

    def getHex(self) -> str:
        """
        Returns the hexadecimal representation of the current BigInteger value.

        @return (str): The hexadecimal representation of the BigInteger value.
        """
        return self.value.getHex()

    def and_func(self, other: BigInteger) -> None:
        """
        Performs the bitwise AND operation on the current BigInteger value and the provided BigInteger object.

        @param other (BigInteger): The BigInteger object to perform AND operation with the current value.
        """
        if len(self.value.digits) < len(other.digits):
            self.value.digits, other.digits = other.digits, self.value.digits

        for i in range(len(other.digits)):
            self.value.digits[i] &= other.digits[i]


class INV:
    def __init__(self, value: BigInteger):
        """
        Constructor of the class, takes a BigInteger object as a parameter and performs the bitwise NOT operation on it.

        @param value (BigInteger): The BigInteger object to perform NOT operation.
        """
        self.value = value
        self.do_not()

    def do_not(self) -> None:
        """
        Performs the bitwise NOT operation on the current BigInteger value.

        """
        for i in range(len(self.value.digits)):
            self.value.digits[i] = ~self.value.digits[i] & 0xFFFFFFFF

        carry = 1
        for i in range(len(self.value.digits)):
            self.value.digits[i] += carry
            carry = self.value.digits[i] >> 32
            self.value.digits[i] &= 0xFFFFFFFF

        if carry != 0:
            self.value.digits.append(carry)

        while len(self.value.digits) > 1 and self.value.digits[-1] == 0:
            self.value.digits.pop()

    def getHex(self) -> str:
        """
        Returns the hexadecimal representation of the current BigInteger value.

        @return (str): The hexadecimal representation of the BigInteger value.
        """
        return self.value.getHex()


class ShiftLeft:
    def __init__(self, value: BigInteger, shift):
        """
        Constructor for the ShiftLeft class.

        @param value: an instance of the BigInteger class representing the value to be shifted
        @param shift: an integer representing the number of bits to shift the value to the left
        """
        self.value = value
        self.shift = shift
        self.shift_left()

    def getHex(self) -> str:
        """
        Returns the hexadecimal representation of the shifted value.

        @return: a string representing the hexadecimal representation of the shifted value.
        """
        return "0x" + self.value.getHex()[4:]

    def shift_left(self) -> BigInteger:
        """
        Shifts the value to the left by the specified number of bits.

        @return: an instance of the BigInteger class representing the shifted value.
        """
        num_bits = self.shift
        num_digits = num_bits // 32
        shift_bits = num_bits % 32

        if num_digits >= len(self.value.digits):
            self.value.digits = [0]
            return self.value

        if shift_bits == 0:
            self.value.digits = [0] * num_digits + self.value.digits
            return self.value

        carry = 0
        for i in range(len(self.value.digits) - 1, num_digits - 1, -1):
            new_digit = (self.value.digits[i] << shift_bits) & 0xFFFFFFFF
            new_digit += carry
            carry = self.value.digits[i] >> (32 - shift_bits)
            self.value.digits[i] = new_digit

        if carry > 0:
            self.value.digits = [carry] + self.value.digits[num_digits:]
        else:
            self.value.digits = self.value.digits[num_digits:]

        self.value.digits = [0] * num_digits + self.value.digits

        return self.value


class ShiftRight:
    def __init__(self, value: BigInteger, shift):
        """
        Returns the hexadecimal representation of the shifted value.

        @return: a string representing the hexadecimal representation of the shifted value.
        """
        self.value = value
        self.shift = shift
        self.shift_left()

    def getHex(self) -> str:
        return "0x" + self.value.getHex()[4:]

    def shift_left(self) -> BigInteger:
        """
        Shifts the value to the left by the specified number of bits.

        @return: an instance of the BigInteger class representing the shifted value.
        """
        num_bits = self.shift
        num_digits = num_bits // 32
        shift_bits = num_bits % 32

        if num_digits >= len(self.value.digits):
            self.value.digits = [0]
            return self.value

        if shift_bits == 0:
            self.value.digits = [0] * num_digits + self.value.digits
            return self.value

        carry = 0
        for i in range(len(self.value.digits) - 1, num_digits - 1, -1):
            new_digit = (self.value.digits[i] >> shift_bits) & 0xFFFFFFFF
            new_digit += carry
            carry = self.value.digits[i] >> (32 - shift_bits)
            self.value.digits[i] = new_digit

        if carry > 0:
            self.value.digits = [carry] + self.value.digits[num_digits:]
        else:
            self.value.digits = self.value.digits[num_digits:]

        self.value.digits = [0] * num_digits + self.value.digits

        return self.value


class ADD:
    def __init__(self, num1: BigInteger, num2: BigInteger):
        """
        Constructor for the ADD class. Adds two instances of the BigInteger class.

        @param num1: an instance of the BigInteger class representing the first number to be added
        @param num2: an instance of the BigInteger class representing the second number to be added
        """
        self.result = BigInteger()
        carry = 0
        i, j = len(num1.digits) - 1, len(num2.digits) - 1

        while i >= 0 or j >= 0 or carry > 0:
            x = num1.digits[i] if i >= 0 else 0
            y = num2.digits[j] if j >= 0 else 0
            s = x + y + carry
            self.result.digits.insert(0, s % (2**32))
            carry = s // (2**32)
            i, j = i - 1, j - 1

    def getHex(self) -> str:
        """
        Returns the hexadecimal representation of the sum.

        @return: a string representing the hexadecimal representation of the sum.
        """
        return self.result.getHex()


class SUB:
    def __init__(self, num1: BigInteger, num2: BigInteger):
        """
        Constructor for the SUB class.

        @param num1: The BigInteger to subtract from.
        @param num2: The BigInteger to subtract.
        """
        self.num1 = num1
        self.num2 = num2
        self.result = BigInteger()
        self.subtract()

    def subtract(self):
        """
        Subtracts num2 from num1 and stores the result in result.
        """
        if len(self.num1.digits) < len(self.num2.digits):
            self.num1, self.num2 = self.num2, self.num1

        borrow = 0
        for i in range(len(self.num2.digits)):
            diff = self.num1.digits[i] - self.num2.digits[i] - borrow
            if diff < 0:
                diff += 2**32
                borrow = 1
            else:
                borrow = 0
            self.result.digits.append(diff)

        for i in range(len(self.num2.digits), len(self.num1.digits)):
            diff = self.num1.digits[i] - borrow
            if diff < 0:
                diff += 2**32
                borrow = 1
            else:
                borrow = 0
            self.result.digits.append(diff)

        while len(self.result.digits) > 1 and self.result.digits[-1] == 0:
            self.result.digits.pop()

        if borrow == 1:
            for i in range(len(self.result.digits)):
                self.result.digits[i] = (2**32 - 1) - self.result.digits[i]
            self.result.digits.append(2**32 - 1)

    def getHex(self) -> str:
        """
        Returns the hexadecimal representation of the result.

        @return: The hexadecimal representation of the result.
        """
        return self.result.getHex()


class MOD:
    def __init__(self, num: BigInteger, modulus: BigInteger):
        """
        Constructor for the MOD class.

        @param num: A BigInteger object representing the number to be divided.
        @param modulus: A BigInteger object representing the divisor.
        """
        self.num = num
        self.modulus = modulus
        self.result = BigInteger()

        self.mod()

    def mod(self):
        """
        Performs division of two BigInteger objects and stores the remainder as the result.
        """
        quotient = BigInteger()
        remainder = BigInteger()

        # Perform long division algorithm digit by digit
        for i in range(len(self.num.digits) - 1, -1, -1):
            remainder.digits.insert(0, self.num.digits[i])
            x = 0
            while remainder >= self.modulus:
                remainder -= self.modulus
                x += 1
            quotient.digits.insert(0, x)

        # Store the remainder as the result
        self.result = remainder

    def getHex(self) -> str:
        """
        Returns the hexadecimal representation of the result.

        @return: A string containing the hexadecimal representation of the result.
        """
        return self.result.getHex()


class MUL:
    def __init__(self, num1: BigInteger, num2: BigInteger):
        """
        Initializes an instance of MUL class with two BigInteger arguments

        @param num1: The first BigInteger object
        @param num2: The second BigInteger object
        """
        self.num1 = num1
        self.num2 = num2
        self.result = BigInteger()
        self.multiply()

    def multiply(self):
        """
        Multiplies two BigInteger objects by performing elementary multiplication and then adding the partial results
        """
        intermediate_results = []

        # Elementary multiplication
        for j in range(len(self.num2.digits) - 1, -1, -1):
            carry = 0
            temp_result = BigInteger()

            for i in range(len(self.num1.digits) - 1, -1, -1):
                product = self.num2.digits[j] * self.num1.digits[i] + carry
                temp_result.digits.insert(0, product % (2**32))
                carry = product // (2**32)

            if carry > 0:
                temp_result.digits.insert(0, carry)

            intermediate_results.append(temp_result)

        # Adding the partial results
        for i in range(len(intermediate_results)):
            for j in range(i):
                intermediate_results[i].digits.append(0)
            add = ADD(self.result, intermediate_results[i])
            self.result = add.result

    def getHex(self) -> str:
        """
        Returns the hexadecimal string representation of the result

        @return: A string representing the result in hexadecimal format
        """
        return self.result.getHex()
