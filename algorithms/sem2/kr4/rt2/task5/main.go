package main

import (
	"fmt"
	"math/rand"
	"time"
)

type StoreA struct {
	grades   map[int]int
	counters [4]int
}

func NewStoreA() *StoreA {
	return &StoreA{grades: make(map[int]int)}
}

func (s *StoreA) Add(id, grade int) {
	s.grades[id] = grade
	s.counters[grade-2]++
}

func (s *StoreA) FindGrade(id int) (int, bool) {
	g, ok := s.grades[id]
	return g, ok
}

func (s *StoreA) CountByGrade(grade int) int {
	return s.counters[grade-2]
}

type StoreB struct {
	ids      []int
	grades   []int8
	counters [4]int
	sorted   bool
}

func (s *StoreB) Add(id, grade int) {
	s.ids = append(s.ids, id)
	s.grades = append(s.grades, int8(grade))
	s.counters[grade-2]++
	s.sorted = false
}

func (s *StoreB) Build() {
	n := len(s.ids)
	for i := 1; i < n; i++ {
		keyID, keyGrade := s.ids[i], s.grades[i]
		j := i - 1
		for j >= 0 && s.ids[j] > keyID {
			s.ids[j+1] = s.ids[j]
			s.grades[j+1] = s.grades[j]
			j--
		}
		s.ids[j+1] = keyID
		s.grades[j+1] = keyGrade
	}
	s.sorted = true
}

func (s *StoreB) FindGrade(id int) (int, bool) {
	lo, hi := 0, len(s.ids)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if s.ids[mid] == id {
			return int(s.grades[mid]), true
		} else if s.ids[mid] < id {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return 0, false
}

func (s *StoreB) CountByGrade(grade int) int {
	return s.counters[grade-2]
}

func benchmark(n int) {
	rng := rand.New(rand.NewSource(42))

	fmt.Printf("Генерация %d записей...\n", n)

	storeA := NewStoreA()
	startA := time.Now()
	for i := 0; i < n; i++ {
		storeA.Add(i, rng.Intn(4)+2)
	}
	fmt.Printf("\n[A] map + счётчик:\n")
	fmt.Printf("  Вставка %d записей: %v\n", n, time.Since(startA))

	rng2 := rand.New(rand.NewSource(99))
	start := time.Now()
	for i := 0; i < 1000; i++ {
		storeA.FindGrade(rng2.Intn(n))
	}
	fmt.Printf("  1000 поисков по ID: %v\n", time.Since(start))

	for g := 2; g <= 5; g++ {
		fmt.Printf("  Оценка %d: %d студентов\n", g, storeA.CountByGrade(g))
	}

	rng3 := rand.New(rand.NewSource(42))
	storeB := &StoreB{}
	startB := time.Now()
	for i := 0; i < n; i++ {
		storeB.Add(i, rng3.Intn(4)+2)
	}
	storeB.Build()
	fmt.Printf("\n[B] отсортированный массив + бинарный поиск:\n")
	fmt.Printf("  Вставка + сортировка %d записей: %v\n", n, time.Since(startB))

	rng4 := rand.New(rand.NewSource(99))
	start = time.Now()
	for i := 0; i < 1000; i++ {
		storeB.FindGrade(rng4.Intn(n))
	}
	fmt.Printf("  1000 поисков по ID: %v\n", time.Since(start))

	for g := 2; g <= 5; g++ {
		fmt.Printf("  Оценка %d: %d студентов\n", g, storeB.CountByGrade(g))
	}
}

func main() {
	fmt.Println("=== Анализ структур хранения архива оценок ===\n")
	fmt.Println("Подход A — map[int]int + [4]int счётчик:")
	fmt.Println("  FindGrade(id)    → O(1)")
	fmt.Println("  CountByGrade(g)  → O(1)")
	fmt.Println("  Память           → ~480 МБ на 10M записей (map overhead)")
	fmt.Println("  Вставка          → O(1) амортизированно")

	fmt.Println("\nПодход B — sorted []int ids + []int8 grades + [4]int счётчик:")
	fmt.Println("  FindGrade(id)    → O(log n)")
	fmt.Println("  CountByGrade(g)  → O(1)")
	fmt.Println("  Память           → ~90 МБ на 10M записей (плотное хранение)")
	fmt.Println("  Вставка          → O(n) (нужна пересортировка) — только для статичных данных")

	fmt.Println("\nВывод: при частых вставках — Подход A.")
	fmt.Println("При ограниченной памяти и статичных данных — Подход B.")

	fmt.Println("\n--- Бенчмарк на 100 000 записей ---")
	benchmark(100_000)
}
