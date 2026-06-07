# Challenge: Describable interface + generic vs interface slice

## Goal
Understand the difference between `[]T` (concrete, homogeneous) and `[]Describable` (interface slice, heterogeneous), and when generics help vs. when they don't.

## Tasks

### 1. Define the interface
```go
type Describable interface {
    Describe() string
}
```

### 2. Implement it on existing shapes
Add a `Describe() string` method to both `Rectangle` and `Circle` that returns a human-readable string, e.g.:
- `"Rectangle 3.0 x 5.0"`
- `"Circle with radius 1.5"`

### 3. Generic logger (homogeneous slice)
```go
func LogAll[T Describable](items []T)
```
Call it with `[]Rectangle{...}` and `[]Circle{...}`.

### 4. Non-generic logger (mixed slice)
```go
func LogMixed(items []Describable)
```
Call it with `[]Describable{rect, circle}` — mixing types in one slice.

### 5. The twist
Try passing a `[]Rectangle` to `LogMixed` and a `[]Describable` to `LogAll`.
Neither will compile. Figure out why, and write a short comment in the code explaining the difference.

## Key insight
Generics constrain to one concrete type per call. Interface slices allow mixing but require explicit wrapping. Neither replaces the other.
