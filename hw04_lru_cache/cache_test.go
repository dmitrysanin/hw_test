package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(3)
		// ccc(300)->bbb(200)->aaa(100):
		c.Set("aaa", 100)
		c.Set("bbb", 200)
		c.Set("ccc", 300)

		// ddd(400)->ccc(300)->bbb(200):
		c.Set("ddd", 400)

		val, ok := c.Get("aaa")
		require.False(t, ok)
		require.Nil(t, val)

		val, ok = c.Get("ddd")
		require.True(t, ok)
		require.Equal(t, 400, val)

		c.Clear()

		val, ok = c.Get("ddd")
		require.False(t, ok)
		require.Nil(t, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)

		// ccc(300)->bbb(200)->aaa(100):
		c.Set("aaa", 100)
		c.Set("bbb", 200)
		c.Set("ccc", 300)
		val, ok = c.Get("ccc")
		require.True(t, ok)
		require.Equal(t, 300, val)

		// bbb(400)->ccc(300)->aaa(100):
		c.Set("bbb", 400)
		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 400, val)

		// aaa(100)->bbb(400)->ccc(300):
		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		// ddd(500)->aaa(100)->bbb(400):
		c.Set("ddd", 500)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)

		val, ok = c.Get("ddd")
		require.True(t, ok)
		require.Equal(t, 500, val)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 400, val)

		// Кэш из одного элемента:
		c1 := NewCache(1)

		c1.Set("aaa", 100)
		val, ok = c1.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		c1.Set("bbb", 200)

		val, ok = c1.Get("aaa")
		require.False(t, ok)
		require.Nil(t, val)

		val, ok = c1.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
