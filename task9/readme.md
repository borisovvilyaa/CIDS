# ECDSA Signature

### Опис функцій

#### GenerateKey()

Генерує нову пару ключів (`KeyPair`) - приватний та публічний ключі.

#### PrintKeyPair(key KeyPair)

Виводить на екран публічний та приватний ключі з пари ключів `key`.

#### CreateSignature(key KeyPair, message *big.Int) Signature

Створює підпис (`Signature`) для повідомлення `message` з використанням пари ключів `key`.

#### PrintSignature(signature Signature)

Виводить на екран компоненти підпису (`Signature`) - `r` та `s`.

#### VerifySignature(key KeyPair, message *big.Int, signature Signature) bool

Перевіряє автентичність підпису (`signature`) для повідомлення `message` з використанням публічного ключа (`key.PublicKey`) з пари ключів `key`. Повертає `true`, якщо підпис є вірним, та `false` у протилежному випадку.

#### SerializePrivateKey(key *big.Int) string

Перетворює приватний ключ (`key`) в рядкове представлення.

#### DeserializePrivateKey(data string) (*big.Int, error)

Перетворює рядкове представлення приватного ключа (`data`) в `*big.Int`. Повертає вказівник на `big.Int` та помилку, якщо виникають проблеми з десеріалізацією.

#### SerializePublicKey(key elliptic.ECPoint) string

Перетворює публічний ключ (`key`) в рядкове представлення.

#### DeserializePublicKey(data string) (elliptic.ECPoint, error)

Перетворює рядкове представлення публічного ключа (`data`) в `elliptic.ECPoint`. Повертає `elliptic.ECPoint` та помилку, якщо виникають проблеми з десеріалізацією.


#### По коду
1. У мене не вийшло зробити повністю цей алгоритм, бо 
- Не зовсім зрозуміло, що таке k^-1 (з wiki)
- Можливо, у мене не правильно працюють функція ScalarMult, типу, на wiki зовсім інші значення, ніж у мене при тих самих константах. Як вони множать точку на число - я так і не зрозумів. Буду дуже вдячний, якщо ви напишете у чому проблема)