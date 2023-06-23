import numpy as np
import matplotlib.pyplot as plt

# Значения параметров a и b
a = 0
b = 7

# Создание массива значений x
x = np.linspace(-10, 10, 400)

# Вычисление соответствующих значений y
y = np.sqrt(x**3 + a*x + b)

# Построение эллиптической кривой
plt.plot(x, y, label='Y^2 = x^3 + {}x + {}'.format(a, b))
plt.plot(x, -y, label='Y^2 = x^3 + {}x + {}'.format(a, b))
plt.xlabel('x')
plt.ylabel('y')
plt.title('Эллиптическая кривая')
plt.legend()
plt.grid(True)
plt.show()
