import random
import math

def is_prime(n, k=5):
    # Використовуємо тест Міллера-Рабіна для перевірки простоти
    if n <= 3:
        return n == 2 or n == 3

    r, s = 0, n - 1
    while s % 2 == 0:
        r += 1
        s //= 2

    for _ in range(k):
        a = random.randint(2, n - 2)
        x = pow(a, s, n)
        if x == 1 or x == n - 1:
            continue

        for _ in range(r - 1):
            x = pow(x, 2, n)
            if x == n - 1:
                break
        else:
            return False

    return True

def generate_prime(length):
    while True:
        # Генеруємо випадкове число з бажаною довжиною
        p = random.getrandbits(length)
        # Забезпечуємо, щоб найстарший біт був встановлений, щоб довжина була точно відповідною
        p |= (1 << length - 1) | 1

        if is_prime(p):
            return p

# Задаємо бажану довжину числа у бітах
length = random.randint(2048, 4096)
p = generate_prime(length)

print(p)
