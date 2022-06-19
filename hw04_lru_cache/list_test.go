package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)

		// Дополнительные тесты (граничные случаи):
		l1 := NewList()

		// Указатели на начало и конец списка при добавлении и удалении единственного элемента
		l1.PushFront(10)
		require.Equal(t, 10, l1.Front().Value)
		require.Equal(t, 10, l1.Back().Value)

		l1.Remove(l1.Front())
		require.Equal(t, 0, l1.Len())
		require.Nil(t, l1.Front())
		require.Nil(t, l1.Back())

		l1.PushFront(10)
		l1.Remove(l1.Back())
		require.Equal(t, 0, l1.Len())
		require.Nil(t, l1.Front())
		require.Nil(t, l1.Back())

		l1.PushBack(10)
		require.Equal(t, 10, l1.Front().Value)
		require.Equal(t, 10, l1.Back().Value)

		l1.Remove(l1.Front())
		require.Equal(t, 0, l1.Len())
		require.Nil(t, l1.Front())
		require.Nil(t, l1.Back())

		l1.PushBack(10)
		l1.Remove(l1.Back())
		require.Equal(t, 0, l1.Len())
		require.Nil(t, l1.Front())
		require.Nil(t, l1.Back())

		// "Перемещение" единственного элемента в начало списка:
		l1.PushFront(10)
		l1.MoveToFront(l1.Front())
		require.Equal(t, 10, l1.Front().Value)
		require.Equal(t, 10, l1.Back().Value)

		// Добавляем в начало списка элемент и переносим ставший последним в начало списка:
		l1.PushFront(20)
		l1.MoveToFront(l1.Back())
		require.Equal(t, 10, l1.Front().Value)
		require.Equal(t, 20, l1.Back().Value)

		// Меняем местами:
		l1.MoveToFront(l1.Back())
		require.Equal(t, 20, l1.Front().Value)
		require.Equal(t, 10, l1.Back().Value)

		// Удаляем последний, а оставшийся единственный "перемещаем" в начало списка
		l1.Remove(l1.Back())
		l1.MoveToFront(l1.Back())
		require.Equal(t, 20, l1.Front().Value)
		require.Equal(t, 20, l1.Back().Value)
	})
}
