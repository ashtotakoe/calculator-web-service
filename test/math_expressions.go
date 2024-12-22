package test

var validTestCases = []struct {
	expression string
	expected   float64
}{
	{
		expression: "1+1",
		expected:   1 + 1,
	},
	{
		expression: "3+3*6",
		expected:   3 + 3*6,
	},
	{
		expression: "1+8/2*4",
		expected:   1 + 8/2*4,
	},
	{
		expression: "(1+1) *2",
		expected:   (1 + 1) * 2,
	},
	{
		expression: "((1+4) * (1+2) +10) *4",
		expected:   ((1+4)*(1+2) + 10) * 4,
	},
	{
		expression: "(4+3+2)/(1+2) * 10 / 3",
		expected:   (4 + 3 + 2) / (1 + 2) * 10 / 3,
	},
	{
		expression: "(70/7) * 10 /((3+2) * (3+7)) -2",
		expected:   (70/7)*10/((3+2)*(3+7)) - 2,
	},
	{
		expression: "((7+1) / (2+2) * 4) / 8 * (32 - ((4+12)*2)) -1",
		expected:   ((7+1)/(2+2)*4)/8*(32-((4+12)*2)) - 1,
	},
	{
		expression: "-1",
		expected:   -1,
	},
	{
		expression: "+5",
		expected:   5,
	},
	{
		expression: "5+5+5+5+5",
		expected:   5 + 5 + 5 + 5 + 5,
	},

	{
		expression: "(1)",
		expected:   1,
	},
	{
		expression: "(1+2*(10) + 10)",
		expected:   (1 + 2*(10) + 10),
	},
	{
		expression: "((1+2)*(5*(7+3) - 70 / (3+4) * (1+2)) - (8-1)) + (10 * (5-1 * (2+3)))",
		expected:   ((1+2)*(5*(7+3)-70/(3+4)*(1+2)) - (8 - 1)) + (10 * (5 - 1*(2+3))),
	},
	{
		expression: "-1+2",
		expected:   -1 + 2,
	},
	{
		expression: "5+ -1",
		expected:   5 + -1,
	},
	{
		expression: "5+ -5 + 7 - -6",
		expected:   5 + -5 + 7 - -6,
	},
	{
		expression: "-(5+5)",
		expected:   -(5 + 5),
	},
	{
		expression: "-90+90",
		expected:   -90 + 90,
	},
	{
		expression: "9*-1",
		expected:   9 * -1,
	},
	{
		expression: "10*(10/10*-10)",
		expected:   10 * 10 / 10 * -10,
	},
	{
		expression: "10*-10",
		expected:   10 * -10,
	},
}

var failTestCases = []struct {
	expression string
}{
	{
		expression: "10/0",
	},
	{
		expression: "2*(10+9",
	},
	{
		expression: "not numbs",
	},
	{
		expression: "2r+10b",
	},
	{
		expression: "10*(10+2*(10+2*(3+4) + 3 * (1+3) + 8 )",
	},
	{
		expression: "10**2",
	},
	{
		expression: "67^21",
	},
	{
		expression: "((((((((((1)))))))))",
	},
	{
		expression: "",
	},
	{
		expression: "()",
	},
	{
		expression: "*10",
	},
	{
		expression: "-+",
	},
	{
		expression: "-",
	},
	{
		expression: "'10",
	},
}

func formatExpression(expression string) string {
	return "expression='" + expression + "'"
}
