# Опис коду 
Цей код виконує S-блок трансформацію для заданого вхідного числа та його інверсну трансформацію. S-блок є частиною алгоритму шифрування і використовується для заміни блоків даних на основі заданої таблиці заміни.


## Інструкція щодо запуску коду
1. Завантажте код у своє середовище розробки Python.
2. Виконайте код у середовищі Python.
3. Запуск S-блоку
3.1 linux: python3 s-block.py 
3.2 mac: python3 s-block.py
3.3 windows: python s-block.py



### Приклад виклику програми та запуску 

```python
input_data = 0b11011010
direct_transformed_data = s_block_direct(input_data)
inverse_transformed_data = s_block_inverse(direct_transformed_data)

print(f"Input Data: {bin(input_data)}")
print(f"Direct Transformed Data: {bin(direct_transformed_data)}")
print(f"Inverse Transformed Data: {bin(inverse_transformed_data)}")
```

*Приклад результату виконання програми*