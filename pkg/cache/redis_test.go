package cache

import (
	"log"
	"testing"
	"time"
)

type User struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}

var cache *Redis

func TestInit(t *testing.T) {
	cache = NewRedis()
}

func TestRedis_Set(t *testing.T) {
	err := cache.Set(Item{"test", "value", time.Second * 1})
	if err != nil {
		t.Fatal(err)
	}

	value := &User{}
	if err := cache.Get("test", value); err != nil {
		if err != nil {
			t.Fatal(err)
			return
		}
		t.Fatalf("value must be 'value' instead is `%s`", value)
	}
}

func TestRedis_Forget(t *testing.T) {
	value := "value"
	key := "test"
	if err := cache.Set(Item{key, value, time.Second * 1}); err != nil {
		t.Fatal(err)
	}
	if err := cache.Forget(key); err != nil {
		t.Fatal(err)
	}

	var result string
	if err := cache.Get(key, &result); err != nil {
		if err != nil {
			t.Fatal(err)
		}
	}
	if result != "" {
		t.Fatal("Try to forget but still exists")
	}
}

func TestStructCache(t *testing.T) {
	key := "struct_cache"
	value := User{Name: "Abolfazl", Lastname: "Alizadeh"}
	if err := cache.Set(Item{key, value, time.Second * 1}); err != nil {
		log.Fatal(err)
	}

	var u User
	err := cache.Get(key, &u)

	if err != nil {
		t.Fatal(err)
	}

	if u.Lastname != value.Lastname && u.Name != value.Name {
		log.Fatal("Name and Lastname User in cache does not match")
	}

}
