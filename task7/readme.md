# Task7

---

## Етапи виконання практичного завдання
1. Написати обгортки до бібліотечних функцій

2. Перевірити коректність роботи перетворень

### 1 Этап


Цей код є пакетом main, який містить декілька функцій для роботи з еліптичними кривими _(Elliptic Curve Cryptography)_.

Основні функції включають:

***BasePointGGet***: Ця функція повертає базову точку G, що використовується для генерації еліптичної кривої.

***ECPointGen***: Ця функція створює нову точку на еліптичній криві з заданими координатами (x, y).

***IsOnCurveCheck***: Ця функція перевіряє, чи належить задана точка (a) до еліптичної кривої. Вона обчислює значення за формулою `y^2 = x^3 + 1`, де a = 0, b = 1, і порівнює його з значенням y точки a. (Змінив криву для простої перевірки чисел)

***AddECPoints***: Ця функція додає дві точки на еліптичній криві (a і b) і повертає результат (c). Вона використовує формули для обчислення координат нової точки на основі координат точок a і b.

***DoubleECPoints***: Ця функція подвоює задану точку на еліптичній криві (a) і повертає результат (c). Вона використовує формули для обчислення координат подвоєної точки на основі координат точки a.

***ScalarMult***: Ця функція множить задану точку на еліптичній криві (a) на скаляр (k) і повертає результат (c). Вона обчислює нову точку, яка є сумою точки a з собою k разів.

***IsEqual***: Ця функція порівнює дві точки на еліптичній криві (pointFirst і pointSecond) і повертає true, якщо вони мають однакові координати, і false - в іншому випадку.

Ці функції дозволяють виконувати операції з еліптичними кривими, такі як генерація точок, додавання точок, подвоєння точок, множення точок на скаляри та порівняння точок на еліптичній криві.






### 2 Этап

Запуск проекту 

```Bash
--> go run task7/main.go
```

Code: 

``` GO
G := ECPoint{}
G = BasePointGGet()

fmt.Printf("Base point is: (%f, %f)\n", G.X, G.Y)

d := 3
k := 5

H1 := ScalarMult(d, G)
H2 := ScalarMult(k, H1)
H3 := ScalarMult(k, G)
H4 := ScalarMult(d, H3)

fmt.Println(IsEqual(H2, H4))

```

Ouptut: 
```Bash
--> go run .\task7\main.go
Base point is: (3.000000, 10.000000)
true
```

And other output 


 ```GO

//create base point
G := ECPoint{}
G = BasePointGGet()
fmt.Printf("Base point is: (%f, %f)\n", G.X, G.Y)
P := ECPoint{}
P = ECPointGen(1, math.Sqrt(8)) // I found this Y when I wassolving on paper
fmt.Printf("Gen point: (%f, %f)\n", P.X, P.Y)
fmt.Println("Point P is on curve:", IsOnCurveCheck(P))
fmt.Println("Add point", AddECPoints(P, G))
fmt.Println("Duble point", DoubleECPoints(P))
fmt.Println("Scalar Mult", ScalarMult(5, P))
PrintECPoint(G)
```

```BASH
Base point is: (3.000000, 10.000000)
Gen point: (1.000000, 2.828427)
Point P is on curve: true
Add point {8.85786437626905 31.005050633883354}
Duble point {-1.71875 -1.3865922037329959}
Scalar Mult {5 14.142135623730951}
Point (3.000000, 10.000000)
```


# UPDATED CODE!!!!
```go
func main() {
	// k*(d*G) = d*(k*G)

	G := elliptic.BasePointGGet()
	k := *big.NewInt(64)
	d := *big.NewInt(256)

	H1 := elliptic.ScalarMult(d, G)
	H2 := elliptic.ScalarMult(k, H1)

	H3 := elliptic.ScalarMult(k, G)
	H4 := elliptic.ScalarMult(d, H3)

	fmt.Println(elliptic.IsEqual(H2, H4))

}
```

OUTPUT:
```bash
true
```
### По коду:
1. Може я не прав, але не розумію, чому у прикладі використовується Int значення точок. Тому було перероблено у float 
2. Створено ще функція isEquel. Було у прикладі, чому б її і сюди не написати
3. Це перша (майже) програма на Go, тому не судіть сильно за кодстайл:) Якщо можете написати фідбек по коду, буду дууууууже вдячний

upd: я зрозумів, чому саме big.Int. У ECDSA float дуже важко використовувати 
