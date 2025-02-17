package questions

/*
https://leetcode.com/problems/employee-importance/description/
Given an integer id that represents an employee's ID, return the total importance value of this employee and all their direct and indirect subordinates.

Example 1:
Input: employees = [[1,5,[2,3]],[2,3,[]],[3,3,[]]], id = 1
Output: 11
Explanation: Employee 1 has an importance value of 5 and has two direct subordinates: employee 2 and employee 3.
They both have an importance value of 3.
Thus, the total importance value of employee 1 is 5 + 3 + 3 = 11.
*/

// Definition for Employee.
type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance1(employees []*Employee, id int) int {
	var empIDMap = map[int]*Employee{}
	var keyEmpNode *Employee
	for _, employee := range employees {
		if employee.Id == id {
			keyEmpNode = employee
		}
		empIDMap[employee.Id] = employee
	}

	var dfs func(*Employee)
	var totalImportance int

	dfs = func(emp *Employee) {
		if emp == nil {
			return
		}
		totalImportance += emp.Importance
		for _, empID := range emp.Subordinates {
			if employee, isExist := empIDMap[empID]; isExist {
				dfs(employee)
			}
		}
	}

	dfs(keyEmpNode)
	return totalImportance
}

func getImportance2(employees []*Employee, id int) int {
	// Create a map for quick lookup of employees by ID
	empIDMap := make(map[int]*Employee, len(employees))
	for _, emp := range employees {
		empIDMap[emp.Id] = emp
	}

	// DFS function to calculate total importance
	var dfs func(int)
	var totalImportance int
	dfs = func(empID int) {
		emp := empIDMap[empID]
		totalImportance += emp.Importance
		for _, subID := range emp.Subordinates {
			dfs(subID) // Recursively add subordinates' importance
		}
	}

	dfs(id) // Start DFS from the given employee ID
	return totalImportance
}
