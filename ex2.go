package main
import
   (
    "fmt"
    "math"
 )
func main() {

var summaVklada float64 = 4823334
procentVklada := 16.0
numberofyears := 3

 fmt.Printf("Входная сумма вклада: %.0f рублей\n", summaVklada);
 fmt.Printf("Годовой процент: %.1f %s\n", procentVklada, "%");
 fmt.Printf("Срок вклада: %d года (лет)\n", numberofyears);

 SummaEnd := summaVklada
 for year := 1;  year <=  numberofyears;  year++ {
    SummaEnd = SummaEnd +  SummaEnd * procentVklada /100
    SummaEnd = SummaEnd * 100
    SummaEnd = math.Round(SummaEnd)
    SummaEnd = SummaEnd / 100
 }
 fmt.Printf("Вы получите %.2f рублей\n", SummaEnd);
}
