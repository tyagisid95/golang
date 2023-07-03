package main

import "fmt"

func main()  {
	fmt.Println("Loops in golangs")

	days := []string{"Saturday","Sunday","Monday","Tuesday","Wednesday","Thursday","Friday"}
	fmt.Println("weekds days are : ",days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println("Day :",days[i])
	// }

	for index,v := range days {
		fmt.Println("Day :",index,v)
	}

	rogueVlaue :=1
	//while loop
	for rogueVlaue<10{

		//break keyword
		// if rogueVlaue ==5{
		// 	break
		// }

		//continiue keyword
		// if rogueVlaue ==5{
		// 	rogueVlaue++
		// 	continue
		// }

		//Go to keyword using jump to statment
		if rogueVlaue ==5{
			goto err
		}
		fmt.Println("Value is : ",rogueVlaue)
		rogueVlaue++
	}
	err:
	fmt.Println("Something went wrong")
}