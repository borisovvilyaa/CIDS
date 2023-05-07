# Task 2

## Робота з великими числами. Логічні та арифметичні операції

---

### Директорія

- `Numbers.py` - файл з класами, завдяки яким є можливість працювати з великими числами
- `test.py` - тести класів та методів

### Task 2.1

#### Задача

Реалізація власного типу даних великого числа з методами `setHex` і `getHex`

#### Код

```python 
def setHex(self, hex_str: str) -> None:
    """
    Sets the value of the BigInteger using a hexadecimal string.

    @param hex_str: The hexadecimal string representing the value of the BigInteger.
    @raise ValueError: If the hex string is empty.
    """
    self.digits = []

    if hex_str == "":
        raise ValueError("[setHex -> Error]: hex string must not be empty \nExample: \"e035c6cfa42609b998b883bc1699df885cef74e2b2cc372eb8fa7e7\"")

    for i in range(len(hex_str) // 8 + 1):
        start = max(len(hex_str) - (i+1)*8, 0)
        end = len(hex_str) - i*8
        chunk = hex_str[start:end]
        if chunk:
            self.digits.insert(0, int(chunk, 16))

def getHex(self) -> str:
    """
    Returns the hexadecimal string representation of the BigInteger.

    @return: The hexadecimal string representation of the BigInteger.
    """
    hex_str = ''.join('{:08x}'.format(d) for d in self.digits).lstrip('0')
    if not hex_str:
        hex_str = '0'
    return "0x" + hex_str.upper()
```

***Особливості:***
1. Обробка помилок при вводі. Якщо ввести неправильно, виведе повідомлення 
    [setHex -> Error]: hex string must not be empty
    Example: "e035c6cfa42609b998b883bc1699df885cef74e2b2cc372eb8fa7e7"


### Task 2.2

***Задача*** - Реалізація побітових операцій для власного типу даних




***Демонстрація ТЕСТІВ***
``` 
.Result test_INV: 0xEE7D27D763F13BF35740C0B6C9D16A1B1312507D402E9867768DBEDF6A4EC247
.Result test_XOR: 0xB60ABDE8A71A6A428D5638388367F60AC1856F24F29BEC3F4F861FC07C6B9BE3
.Result test_OR: 0xB78AFDE9BF1EEE4EADFF3F79B76FF7EEEDEDEFA6FFDBEFBFCFF65FE0FDFBBFFB
.Result test_AND: 0x18040011804840C20A90741340801E42C6880820D4003808070402081902418
.Result test_ShiftLeft: 0x82D8299C0EC40CA8BF3F49362E95E4ECEDAF82BFD167988972412095B13DB800
.Result test_ShiftRight: 0x00118374009C0F6C00A8BF7500362F8100ECEE6E00BFD1F0008972D60095B13D
```

*Взято при значеннях 0х1182d8299c0ec40ca8bf3f49362e95e4ecedaf82bfd167988972412095b13db8 та 0хa78865c13b14ae4e25e90771b54963ee2d68c0a64d4a8ba7c6f45ee0e9daa65b*


### Task 2.3

***Задача*** - Реалізація арифметичних операцій


***Демонстрація ТЕСТІВ***
``` 
Result test_ADD: 0xB90B3DEAD723725ACEA846BAEB77F9D31A5670290D1BF3405066A0017F8BE413
.Result test_SUB: 0x96058D979F05EA427D29C8277F1ACE09407B11248D79240E3D821DBF542968A3FFFFFFFF
.Result test_MOD: 0x1182D8299C0EC40CA8BF3F49362E95E4ECEDAF82BFD167988972412095B13DB8
.Result test_MUL: 0xB75AF7BD1E07DA3A4E2CAEDFB839F91EC4387BFC24245EE5F043E74A992A6443A0E63312210E6723D4D396D87A0DBC5FE7DEA5296A5A61842F5A2C12AB64068
```



*Взято при значеннях 0х1182d8299c0ec40ca8bf3f49362e95e4ecedaf82bfd167988972412095b13db8 та 0хa78865c13b14ae4e25e90771b54963ee2d68c0a64d4a8ba7c6f45ee0e9daa65b*


***Особливості:***
1. Було додатково зроблено метод MUL
---

#### Загальний вивід тестів
```

Result test_ADD: 0xB90B3DEAD723725ACEA846BAEB77F9D31A5670290D1BF3405066A0017F8BE413
.Result test_AND: 0x18040011804840C20A90741340801E42C6880820D4003808070402081902418
.Result test_INV: 0xEE7D27D763F13BF35740C0B6C9D16A1B1312507D402E9867768DBEDF6A4EC247
.Result test_MOD: 0x1182D8299C0EC40CA8BF3F49362E95E4ECEDAF82BFD167988972412095B13DB8
.Result test_MUL: 0xB75AF7BD1E07DA3A4E2CAEDFB839F91EC4387BFC24245EE5F043E74A992A6443A0E63312210E6723D4D396D87A0DBC5FE7DEA5296A5A61842F5A2C12AB64068
.Result test_OR: 0xB78AFDE9BF1EEE4EADFF3F79B76FF7EEEDEDEFA6FFDBEFBFCFF65FE0FDFBBFFB
.Result test_SUB: 0x96058D979F05EA427D29C8277F1ACE09407B11248D79240E3D821DBF542968A3FFFFFFFF
.Result test_ShiftLeft: 0x82D8299C0EC40CA8BF3F49362E95E4ECEDAF82BFD167988972412095B13DB800
.Result test_ShiftRight: 0x00118374009C0F6C00A8BF7500362F8100ECEE6E00BFD1F0008972D60095B13D
.Result test_XOR: 0xB60ABDE8A71A6A428D5638388367F60AC1856F24F29BEC3F4F861FC07C6B9BE3
.
----------------------------------------------------------------------
Ran 10 tests in 0.005s

OK
```

