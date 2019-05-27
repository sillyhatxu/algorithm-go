# shapeArea

Below we will define an n-interesting polygon. Your task is to find the area of a polygon for a given n.

A 1-interesting polygon is just a square with a side of length 1. An n-interesting polygon is obtained by taking the n - 1-interesting polygon and appending 1-interesting polygons to its rim, side by side. You can see the 1-, 2-, 3- and 4-interesting polygons in the picture below.

![](https://lh3.googleusercontent.com/ti8XxeRUx8Cgq1s0rHlBNnqDVw53jHkip8NNBVtHxpx_kccpJPz5tvDoP-rv20qXabOiBiNFcMaP_EeVtKF3sdQaMrceW9_sPX3KObfFloYuEdfGn4CKsUO58OAQdzR2hPAXdoA3zol9N3gygtdlVYmAMAzDxOMYQjpINbH1e_SriEDk1Z7M6kxpZMQeBz97NI5Ndzrf6Ywng289d9ukD97IexZ0llJinOYexsZLgf4EwshNws3I3IAt-hGcNtlgdDVS7w1vEe-H-MLb4cy9bsZDM9CFQKa9aZHIlfslw89LPYxeerg9U9A62-uT7KX9wRPmxHlUYj5SnqeoVVrmG5fC71Ilg9qgabwrShVLnl3AIN8iHOQREXxR78ikAIWtcMByviwNtuMYDtCt7cG4Z4pL2RUmmsMP_pRIl6DKdM8XwvCokx_Jw6_KBNLyv5oa5MDfHHP-E-Dxjvoz6KnXCJVx-_lqcqe4MspIBv_Jcv0OGtG18W0pJzxSCBCD32y67k31mHhxIX5__r04Zr5GUWmfyKoyg1f8zWrgDxWAN_TSdnF2SeV4o9K8Z4QMndJdKiQANrIDJmizZNRaOZgB6KgqZAI4EvaZeR7FBVtK46Gs2igxzL4mxoJg83lT0ze9s_w1jVZ0eTypx_Z9yjp4zHcIiGTlurM=w1498-h729-no)

Example

For n = 2, the output should be
shapeArea(n) = 5;
For n = 3, the output should be
shapeArea(n) = 13.
Input/Output

[execution time limit] 4 seconds (go)

[input] integer n

Guaranteed constraints:
1 ≤ n < 104.

[output] integer

The area of the n-interesting polygon.
[Go] Syntax Tips

// Prints help message to the console
// Returns a string
func helloWorld(name string) string {
    fmt.Printf("This prints to the console when you Run Tests");
    return "Hello, " + name;
}

