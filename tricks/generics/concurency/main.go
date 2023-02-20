package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Пример использования дженериков при конкурентном доступе из горутин
// https://medium.com/@jon_43067/go-generics-and-concurrency-d0dccab73a73
// https://github.com/jonbodner/gcon/blob/main/gcon.go

func ThingToDoConcurrently(ctx context.Context, a int, b string) (string, error) {
	// do some stuff in here, like I/O or a long computation
	return "", nil
}

func OtherThingToDoConcurrently(ctx context.Context, a int, b string) (string, error) {
	// do some stuff in here, like I/O or a long computation
	return "", nil
}

type Input struct {
	A int
	B string
}

type Output struct {
	Out1 string
	Out2 string
}

// Инфоиация функции дженерика с констрэйнами any
type Func[T, V any] func(context.Context, T) (V, error)

// Promise Функция промис для асинхронного вызова функций
// Обещание имеет возвращаемое значение и ошибку, а также готовый канал. Обратите внимание,
// что поля val и err не экспортируются; мы хотим заблокировать доступ к ним, пока значения не будут заполнены.
type Promise[V any] struct {
	val  V
	err  error
	done <-chan struct{}
}

// Run выполняет функцию Func в конструкции промиса
func Run[T, V any](ctx context.Context, t T, f Func[T, V]) *Promise[V] {
	done := make(chan struct{})
	p := Promise[V]{
		done: done,
	}

	go func() {
		defer close(done)
		p.val, p.err = f(ctx, t)
	}()

	return &p
}

// Get Run и Get работают вместе, чтобы гарантировать, что данные будут записаны в Promise до того, как они станут доступными.
// Мы никогда не пишем в готовый канал, хранящийся в Promise. Вместо этого он закрывается, чтобы сигнализировать о завершении
// обработки с использованием шаблона готового канала (done channel pattern.).
func (p *Promise[V]) Get() (V, error) {
	<-p.done // Флаг закрытия канала (завершения наполнение данными структуры Promise

	return p.val, p.err
}

type Waiter interface {
	Wait() error
}

func (p *Promise[V]) Wait() error {
	<-p.done
	return p.err
}

// WaitTakeOne Функция, ожидающая завершения одного или нескольких экземпляров Promise, но возвращающаяся немедленно,
//если какой-либо из них возвращает ошибку. Вот первая попытка реализации:
// Поскольку Promise является универсальным типом, вы должны указывать параметры типа всякий раз, когда используете его.
// поэтому будет ошибка компиляции
//func WaitTakeOne(ps ...*Promise) error { // cannot use generic type Promise[V any] without instantiation
//	var wg sync.WaitGroup
//	wg.Add(len(ps))
//	errChan := make(chan error, len(ps))
//	done := make(chan struct{})
//
//	for _, p := range ps {
//		go func(p *Promise) {
//			defer wg.Done()
//			_, err := p.Get()
//			if err != nil {
//				errChan <- err
//			}
//		}(p)
//	}
//	go func() {
//		defer close(done)
//		wg.Wait()
//	}()
//	select {
//	case err := <-errChan:
//		return err
//	case <-done:
//	}
//	return nil
//}

// WaitTakeTwo Верное решение
func WaitTakeTwo[V any](ps ...*Promise[V]) error {
	var wg sync.WaitGroup
	wg.Add(len(ps))
	errChan := make(chan error, len(ps))
	done := make(chan struct{})
	for _, p := range ps {
		go func(p *Promise[V]) {
			defer wg.Done()
			_, err := p.Get()
			if err != nil {
				errChan <- err
			}
		}(p)
	}
	go func() {
		defer close(done)
		wg.Wait()
	}()
	select {
	case err := <-errChan:
		return err
	case <-done:
	}
	return nil
}

func Wait(ws ...Waiter) error {
	var wg sync.WaitGroup
	wg.Add(len(ws))
	errChan := make(chan error, len(ws))
	done := make(chan struct{})
	for _, w := range ws {
		go func(w Waiter) {
			defer wg.Done()
			err := w.Wait()
			if err != nil {
				errChan <- err
			}
		}(w)
	}
	go func() {
		defer close(done)
		wg.Wait()
	}()
	select {
	case err := <-errChan:
		return err
	case <-done:
	}
	return nil
}

// RunProcess базовая версия работы с горутинами
func RunProcess(ctx context.Context, data Input) (Output, error) {
	var out Output

	type result struct {
		out string
		err error
	}

	// Буферизированный канал разером 1
	ch1 := make(chan result, 1)

	go func() {
		out, err := ThingToDoConcurrently(ctx, data.A, data.B)
		ch1 <- result{out, err}
	}()

	ch2 := make(chan result, 1)
	go func() {
		out, err := OtherThingToDoConcurrently(ctx, data.A, data.B)
		ch2 <- result{out, err}
	}()

	count := 0
	for count < 2 {
		select {
		case r := <-ch1:
			if r.err != nil {
				return Output{}, r.err
			}
			out.Out1 = r.out
			count++
		case r := <-ch2:
			if r.err != nil {
				return Output{}, r.err
			}
			out.Out2 = r.out
			count++
		}
	}

	return out, nil

}

// RunProcess2 - работа с Get Run
// Эта версия RunProcess скрывает запуск всех горутин и управление каналами. Читабельность значительно улучшилась,
// но это не совсем то, что было у нас изначально. Раньше, если вызов OtherThingToDoConcurrently завершался раньше ThingToDoConcurrently
// и завершался с ошибкой, код не ждал; RunProcess немедленно завершает работу, возвращая ошибку.
func RunProcess2(ctx context.Context, data Input) (Output, error) {
	var out Output

	// Инициализация функций замыканий
	var f1 Func[Input, string] = func(ctx context.Context, in Input) (string, error) {
		return ThingToDoConcurrently(ctx, in.A, in.B)
	}
	var f2 Func[Input, string] = func(ctx context.Context, in Input) (string, error) {
		return OtherThingToDoConcurrently(ctx, in.A, in.B)
	}

	p1 := Run(ctx, data, f1)
	p2 := Run(ctx, data, f2)

	v1, err := p1.Get()
	if err != nil {
		return Output{}, err
	}
	out.Out1 = v1

	v2, err := p2.Get()
	if err != nil {
		return Output{}, err
	}
	out.Out2 = v2

	return out, nil
}

// RunProcess3 using Promise, Run, and Wait:
func RunProcess3(ctx context.Context, data Input) (Output, error) {
	var out Output

	var f1 Func[Input, string] = func(ctx context.Context, in Input) (string, error) {
		return ThingToDoConcurrently(ctx, in.A, in.B)
	}

	var f2 Func[Input, string] = func(ctx context.Context, in Input) (string, error) {
		return OtherThingToDoConcurrently(ctx, in.A, in.B)
	}

	p1 := Run(ctx, data, f1)
	p2 := Run(ctx, data, f2)

	if err := Wait(p1, p2); err != nil {
		return Output{}, err
	}

	out.Out1, _ = p1.Get()
	out.Out2, _ = p2.Get()

	return out, nil
}

// WithCancellation использует преимущества нашего старого друга, шаблона канала готовности.
// Фактически мы используем два готовых канала. В каждом контексте есть метод Done, который
// возвращает канал. Этот канал закрывается, когда контекст отменяется либо по тайм-ауту,
// либо по вызову функции отмены, связанной с контекстом. Мы также создаем свой собственный
// сделанный канал. Функция, возвращаемая WithCancellation, вызывает переданную функцию
// Func в горутине. Когда горутина завершается, наш канал done закрывается.
// Мы используем оператор select, чтобы дождаться закрытия канала done или закрытия канала,
// возвращаемого контекстным методом Done. Если наш закрывается первым, мы возвращаем
// результаты переданного Func. Если канал done контекста закрывается первым, мы возвращаем
// нулевое значение и ошибку из контекста.
func WithCancellation[T, V any](f Func[T, V]) Func[T, V] {
	return func(ctx context.Context, t T) (V, error) {
		done := make(chan struct{})
		var val V
		var err error

		go func() {
			defer close(done)
			val, err = f(ctx, t)
		}()

		select {
		case <-ctx.Done():
			var zero V
			return zero, ctx.Err()
		case <-done:
		}

		return val, err
	}
}

// ImplementSolution
// Мы отправляем данные в два сервиса, а затем берем результаты этих двух вызовов и
//отправляем их в третий, возвращая результат. Весь процесс должен занимать менее 50 миллисекунд,
//иначе должна быть возвращена ошибка.

func ImplementSolution(ctx context.Context, in Input) (Output, error) {
	ws1 := WithCancellation(func(ctx context.Context, in Input) (string, error) {
		return ThingToDoConcurrently(ctx, in.A, in.B)
	})

	ws2 := WithCancellation(func(ctx context.Context, in Input) (string, error) {
		return OtherThingToDoConcurrently(ctx, in.A, in.B)
	})
	ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()

	p1 := Run(ctx, in, ws1)
	p2 := Run(ctx, in, ws2)
	if err := Wait(p1, p2); err != nil {
		return Output{}, err
	}
	type intermediateResults struct {
		result1 string
		result2 string
	}
	ws3 := WithCancellation(func(ctx context.Context, ir intermediateResults) (Output, error) {
		//return WebService3(ctx, ir.result1, ir.result2)
		return Output{}, nil
	})
	var ir intermediateResults
	ir.result1, _ = p1.Get()
	ir.result2, _ = p2.Get()
	p3 := Run(ctx, ir, ws3)

	return p3.Get()
}

// Then Есть еще одна операция, которую люди любят делать с обещаниями: связывать их вместе.
// Давайте посмотрим на реализацию Then:
func Then[T, V any](ctx context.Context, p *Promise[T], f Func[T, V]) *Promise[V] {
	done := make(chan struct{})
	out := Promise[V]{
		done: done,
	}
	go func() {
		defer close(done)
		val, err := p.Get()
		if err != nil {
			out.err = err
			return
		}
		out.val, out.err = f(ctx, val)
	}()
	return &out
}

func main() {
	// 1. Проверка работы Run
	ctx := context.Background()
	p1 := Run(ctx, 10, func(ctx context.Context, i int) (int, error) {
		return i * 2, nil
	})

	p2 := Run(ctx, 10, func(ctx context.Context, i int) (int, error) {
		return i * 3, nil
	})
	err := WaitTakeTwo(p1, p2) // Тут ждет пока не получим два значения из Run
	fmt.Println(err)           // nil

	fmt.Println(p1.Get()) // 20 <nil>
	fmt.Println(p2.Get()) // 30 <nil>

	// Но если нужно разные типы возврата, ваш код не скомпилируется. Вот краткий пример:
	// type *Promise[string] of p2 does not match inferred type *Promise[int] for *Promise[V]

	//p1 = Run(ctx, 10, func(ctx context.Context, i int) (int, error) {
	//	return i * 2, nil
	//})
	//p2 = Run(ctx, 10, func(ctx context.Context, i int) (string, error) {
	//	return strconv.Itoa(i), nil
	//})
	//err := WaitTakeTwo(p1, p2)}

	// Чтобы создать функцию ожидания, которая работает, когда не все экземпляры
	//Promise возвращают значения одного и того же типа, нам нужно определить новый,
	//не универсальный интерфейс, Waiter, и нам нужно, чтобы Promise реализовал этот интерфейс

	// Использование Then
	ctx = context.Background()
	p1 = Run(ctx, 10, func(ctx context.Context, i int) (int, error) {
		return i * 2, nil
	})

	p2 = Then(ctx, p1, func(ctx context.Context, i int) (int, error) {
		return i * 4, nil
	})
	val, err := p2.Get()
	fmt.Println(val, err) // 80 <nil>

}
