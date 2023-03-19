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

		c.Set("first", 10)
		c.Set("second", 20)
		c.Set("third", 30)
		c.Set("fourth", 40)

		val, ok := c.Get("first")
		require.Nil(t, val)
		require.False(t, ok)

		c = NewCache(3)
		c.Set("first", 10)  // [10]
		c.Set("second", 20) // [20 10]
		c.Set("third", 30)  // [30 20 10]

		c.Set("second", 90) // [90 30 10]
		c.Get("third")      // [30 90 10]
		c.Set("fourth", 20) // [20 30 90]

		val, ok = c.Get("first")
		require.Nil(t, val)
		require.False(t, ok)
	})

	t.Run("test clear cache", func(t *testing.T) {
		c := NewCache(2)
		c.Set("first", 10)
		c.Set("second", 20)

		c.Clear()

		val, ok := c.Get("first")
		require.Nil(t, val)
		require.False(t, ok)

		val, ok = c.Get("second")
		require.Nil(t, val)
		require.False(t, ok)
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
