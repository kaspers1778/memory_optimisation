# memory_optimisation

My student ID is IP8102,so by formula ```2 mod 16 + 1``` i got 3rd varinat.
So origin code of my programm looks like : 
```
   static void Main(string[] args)
        {
            int[,,] a = new int[10,10];
            int res = 0;

            for (int i = 0; i < 10; i++)
            {
                for (int j = 0; j < 10; j++)
                {
                        a[j,i]++;
                }
            }

        }
        
```
So far I working in golang,I convert C++ code in GO,so it turns to:
```
func main(){
	var a [10][10]int
	res:=0
	for i :=0; i<10; i++{
		for j:=0; j<10; j++{
			a[j][i]++
		}
	}
}
```

Lets see what we can do to optimize that block of code.
1. Firstable,we need to remove *res* variable,cause it's don't uses anywhere else but implementation.
2. We need to swap *i* and *j* variables in brackets,so we will reach better space location.

After this changes the code will look like :
```
func main(){
	var a [10][10]int
	for i :=0; i<10; i++{
		for j:=0; j<10; j++{
			a[i][j]++
		}
	}
}
```
Lets test this out and see if  our changes we made it work faster.

| Time          | 10    | 100  |  1000 | 10000 |
|---------------|-------|------|-------|-------|
| Origin code   | 456ns | 52ps | 3.8ms | 1.2s  |
| Modified code | 384ns | 48ps | 2.7ms | 289ms |

Conclusion: As we manage to see from the results in table,difference between origin code and modified performance does not seems so big in small ranges of array size, but ot grows so much with increasing of size and at 10000x10000 number of matrix modified code runs faster that origin one  more then 4 times.So tiime and space locality realy matters on perfomance of code.
