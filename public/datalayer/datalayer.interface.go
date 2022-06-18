// data layer abstract
package datalayer

type IDatalayer[T any] interface {
	Get() (T, error)
	Update(T) error
	Delete(int64) error
}
