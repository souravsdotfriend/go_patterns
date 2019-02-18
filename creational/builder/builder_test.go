package builder

import (
	"testing"
	"fmt"
)

func TestBuilder1(t *testing.T) {
	assembly := NewBuilder().Paint(RED)

	familyCar := assembly.TopSpeed(50 * KPH).Build()
	res := familyCar.Drive()
	fmt.Println(res)

	sportsCar := assembly.TopSpeed(150 * KPH).Build()
	res = sportsCar.Drive()
	fmt.Println(res)
}
