package utils

type UtilsInterface interface{}

type Utils struct{}

func Init() *Utils {
	return &Utils{}
}
