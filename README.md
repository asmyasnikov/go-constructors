# Как в Go жить без перегрузок и значений аргументов по умолчанию?

## План доклада

### Вводное слово 

кто я - много практики, живу в OpenSource, много разбираю багов, возникающих при работе с моим же кодом, много думаю про то, как не дать людям написать баг. 
зачем хотеть перегрузки функций и значения аргументов по умолчанию? Тут нытье про то, что в других ЯП есть такое, а в Go нет

### Чуть-чуть теории (вики, словари) про понятия

Перегрузка функций нужна для следующих целей:

* Избежание дублирования имён функций. Если нужно выполнить сходные действия, но с различной программной логикой, можно определить несколько функций с одинаковым именем, но различными параметрами. [1](http://cppstudio.com/post/406/)
* Предоставление различной семантики для функции. [5](https://www.geeksforgeeks.org/function-overloading-c/) В зависимости от сигнатуры функций (списка параметров) перегруженные функции могут по-разному выполняться и возвращать значения разных типов. [3](https://purecodecpp.com/archives/1391)
* Повышение читаемости программы. Если нужно выполнить только одну операцию, и у функций одинаковое название, это упрощает понимание поведения программы. Например, если нужно сложить заданные числа, но аргументов может быть любое количество. [5](https://www.geeksforgeeks.org/function-overloading-c/)
* Перегрузка функций — это возможность в языках программирования использовать одноимённые подпрограммы с разным числом аргументов или их типами. [4](https://ru.ruwiki.ru/wiki/Перегрузка_процедур_и_функций)

### Практическая задача

Описать сложный тип, имеющий состояние, и предложить пути его использования. Например, драйвер базы данных (позже может перепридумаю пример)

#### Пути решения

Решая каждый раз одну и ту же задачу разными подходами - оцениваем плюсы и минусы подхода

##### Публичный тип с публичными полями. 

```go
type Driver struct {
  Connection *grpc.ClientConn
  Timeout time.Duration
  Mutex sync.Mutex
  CallCounter uint64
}
```

Плюсы:
1) Легко объявлять и использовать в своем коде 
   ```go
   db := &Driver{
     Connection: conn,
     Timeout: time.Second,
     Mutex: sync.Mutex{},
     CallCounter: 0,
   }
   ```
2) Прозрачность

Минусы:
1) Легко накосячить с значениями полей
2) Не позволяет автоматически стартовать фоновые процессы

##### Публичный тип с приватным init()

Пример - sync.Pool, container/list.List

Плюсы:
- по прежнему удобно создавать переменную типа

Минусы:
- на каждом вызове публичного метода следует проверять в каком состоянии объект и вызывать init()

##### Публичный тип с приватными полями + публичный конструктор типа

Пример - database/sql.DB

##### Публичный тип с приватными полями + публичный конструктор типа
