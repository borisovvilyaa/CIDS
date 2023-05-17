# Опис коду 
Цей код виконує S-блок трансформацію для заданого вхідного числа та його інверсну трансформацію. S-блок є частиною алгоритму шифрування і використовується для заміни блоків даних на основі заданої таблиці заміни.

## Задача

**Власна програмна реалізація алгоритмів S-блоку та P-блоку (пряме та зворотне перетворення)



## Інструкція щодо запуску коду
1. Завантажте код у своє середовище розробки Python.
2. Виконайте код у середовищі Python.
3. Запуск S-блоку
3.1 linux: python3 s-block.py 
3.2 mac: python3 s-block.py
3.3 windows: python s-block.py



### Приклад виклику програми та запуску S-block та P-block
1. S-block 

```python
#S-block
s_block = SBlock()
input_data = 0b11011010
direct_transformed_data = s_block.encrypt(input_data)
inverse_transformed_data = s_block.decrypt(direct_transformed_data)

print(f"Input Data: {bin(input_data)}")
print(f"Direct Transformed Data: {bin(direct_transformed_data)}")
print(f"Inverse Transformed Data: {bin(inverse_transformed_data)}")

```

*Приклад результату виконання програми*

```
Input Data: 0b11011010
Direct Transformed Data: 0b11011011
Inverse Transformed Data: 0b11011010
```


2. P-block 

```python
#P-block
input_data = 0b11011010  # Input data (8 bits)

p_block_cipher = PBlock()
encrypted_data = p_block_cipher.encrypt(input_data)

inverse_p_block_cipher = InversePBlock()
decrypted_data = inverse_p_block_cipher.decrypt(encrypted_data)

print("Input Data:", bin(input_data))
print("Encrypted Data:", bin(encrypted_data))
print("Decrypted Data:", bin(decrypted_data))


```

*Приклад результату виконання програми*

```
Input Data: 0b11011010
Encrypted Data: 0b1111001
Decrypted Data: 0b11011010
```

## OUTPUT Tests

```
.encrypted_data p-block: 121
.decrypted_data | input_data p-block: 218
.encrypted_data s-block: 219
.decrypted_data | input_data s-block: 218
.
----------------------------------------------------------------------
Ran 4 tests in 0.001s

OK
```