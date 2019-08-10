package helper

import(
	"fmt"
)

func CheckErr(err error){
	if err != nil {
		fmt.Println("-------err = " ,err.Error() )
		panic(err)
	}
}

