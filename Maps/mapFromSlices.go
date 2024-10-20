package main

import "fmt"

type customer struct {
	name string
	age  int
}

func getCustomerMap(names []string, ages []int) (map[string]customer, error) {
	if len(names) != len(ages) {
		return nil, fmt.Errorf("names and ages have different length")
	}
	customerMap := make(map[string]customer)
	for i, name := range names {
		customerMap[name] = customer{
			name: name,
			age:  ages[i],
		}
	}
	return customerMap, nil
}

func test(names []string, ages []int) {
	fmt.Println("Creating Map ======================")
	defer fmt.Println("======================Map created") // it will execute when the function execution finished.
	customers, err := getCustomerMap(names, ages)
	if err != nil {
		fmt.Println(err)
	}
	for key, value := range customers {
		fmt.Printf("Key => %v and Value =>\n", key)
		fmt.Printf("Name : %v\n", value.name)
		fmt.Printf("Age : %v\n", value.age)
	}
}

func main() {
	test(
		[]string{"Afid", "Md", "Azwad"},
		[]int{23, 24, 25},
	)
	test(
		[]string{"Karim", "Rahim", "Joshim"},
		[]int{22, 24, 26},
	)
}
